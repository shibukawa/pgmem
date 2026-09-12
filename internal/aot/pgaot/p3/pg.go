package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidatePgVersion(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v7 = m.G0
	v9 = v7 - int32(1232)
	m.G0 = v9
	v16 = F_strtox_2(m, int32(560458), v9+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = l0
	v25 = F_pg_snprintf(m, v9+int32(208), int32(1024), int32(530303), v9+int32(112))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v30 = F_AllocateFile(m, v9+int32(208), int32(230977))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			if v30 == int32(0) {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[40]))
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					if v35 == int32(44) {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
							F_errmsg(m, int32(13412), v9+int32(16))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(208)
								F_errdetail(m, int32(626903), v9)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									F_errfinish(m, int32(493597), int32(1793), int32(271517))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
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
						F_errcode_for_file_access(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(208)
							F_errmsg(m, int32(298919), v9+int32(32))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(493597), int32(1797), int32(271517))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
				v57 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+128)) = uint8(v57)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v9 + int32(128)
				v65 = F_fscanf(m, v30, int32(175248), v9+int32(96))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					v73 = F_strtox_2(m, v9+int32(128), v9+int32(204), int32(10), int64(2147483648))
					mBase = m.M
					if v65 != int32(1) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
								F_errmsg(m, int32(13412), v9-int32(-64))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v9 + int32(208)
									F_errdetail(m, int32(655515), v9+int32(48))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return
									} else {
										F_errhint(m, int32(655114), int32(0))
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return
										} else {
											F_errfinish(m, int32(493597), int32(1811), int32(271517))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
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
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
						if v77 == v9+int32(128) {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
									F_errmsg(m, int32(13412), v9-int32(-64))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v9 + int32(208)
										F_errdetail(m, int32(655515), v9+int32(48))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return
										} else {
											F_errhint(m, int32(655114), int32(0))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												F_errfinish(m, int32(493597), int32(1811), int32(271517))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
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
							v81 = F_FreeFile(m, v30)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								if base.I32_wrap_i64(v16) != base.I32_wrap_i64(v73) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return
										} else {
											F_errmsg(m, int32(214602), int32(0))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = int32(560458)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v9 + int32(128)
												F_errdetail(m, int32(604951), v9+int32(80))
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													F_errfinish(m, int32(493597), int32(1821), int32(271517))
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
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
									m.G0 = v9 + int32(1232)
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
func F__PG_init_plpgsql(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1274])))
	if v4 == int32(0) {
		F_DefineCustomEnumVariable(m, int32(109339), int32(599064), int32(0), int32(4608956), int32(4176480))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_DefineCustomBoolVariable(m, int32(151002), int32(598702), int32(4608960), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_DefineCustomBoolVariable(m, int32(117988), int32(586664), int32(4176472), int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_DefineCustomStringVariable(m, int32(156416), int32(627284), int32(4608972), int32(373019), int32(6825), int32(6826))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_DefineCustomStringVariable(m, int32(131682), int32(607610), int32(4608976), int32(373019), int32(6825), int32(6827))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_MarkGUCPrefixReserved(m, int32(300561))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, _consts[12]))
								v48 = F_MemoryContextAlloc(m, v46, int32(12))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = int32(4794)
									v54 = int32(4410840)
									v55 = *(*int32)(unsafe.Add(mBase, _consts[1275]))
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v55
									*(*int32)(unsafe.Add(mBase, _consts[1275])) = v48
									v60 = *(*int32)(unsafe.Add(mBase, _consts[12]))
									v62 = F_MemoryContextAlloc(m, v60, int32(12))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v64
										*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = int32(4790)
										v68 = int32(4410844)
										v69 = *(*int32)(unsafe.Add(mBase, _consts[174]))
										*(*int32)(unsafe.Add(mBase, uint32(v62))) = v69
										*(*int32)(unsafe.Add(mBase, _consts[174])) = v62
										v73 = m.G0
										v75 = v73 - int32(48)
										m.G0 = v75
										v78 = *(*int32)(unsafe.Add(mBase, _consts[1276]))
										if v78 == v64 {
											*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = int64(292057776192)
											v87 = F_hash_create(m, int32(323584), int32(16), v75, int32(24))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[1276])) = v87
												v90 = v87
												v93 = F_hash_search(m, v90, int32(276574), int32(1), v75)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
													if v95 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = int32(0)
													} else {
													}
													m.G0 = v75 + int32(48)
													v106 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[1274])) = uint8(v106)
													*(*int32)(unsafe.Add(mBase, _consts[1277])) = v93 - int32(-64)
													return
												}
											}
										} else {
											v90 = v78
											v93 = F_hash_search(m, v90, int32(276574), int32(1), v75)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
												if v95 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = int32(0)
												} else {
												}
												m.G0 = v75 + int32(48)
												v106 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[1274])) = uint8(v106)
												*(*int32)(unsafe.Add(mBase, _consts[1277])) = v93 - int32(-64)
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
		return
	}
}
func F_create_pg_locale_icu(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(431523), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492931), int32(215), int32(38403))
				v18 = m.ExcPending
				if v18 != 0 {
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
func F_do_pg_backup_stop(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v395 int64
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int64
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v422 int64
	_ = v422
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	v11 = m.G0
	v13 = v11 - int32(2272)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v16 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v24)
	v26 = v24
	goto L3
L2:
	;
	v26 = int32(0)
	goto L3
L3:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L13
	} else {
		goto L165
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L13
	} else {
		goto L161
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L13
	} else {
		goto L156
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L13
	} else {
		goto L151
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L13
	} else {
		goto L146
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v30 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+164)) = v37 - int32(1)
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[359])) = uint8(v42)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
	if (v26|(v46^int32(-1)))&int32(1) == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if v26 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l1 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+440)) = int32(1)
	if v56 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L28
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_s_lock(m, v60+int32(440), int32(498539), int32(9275), int32(234271))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+440)) = int32(0)
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v69)+432))
	if base.Ui64(v72) <= base.Ui64(v73) {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v80 = F_LWLockAcquire(m, v76+int32(1152), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)+136))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+144))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v86
	v89 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v89+int32(1152))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	F_XLogRegisterData(m, l0+int32(1032), int32(8))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v103 = F_XLogInsert(m, int32(0), int32(80))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v103
	v107 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+308))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v108
	F_XLogBeginInsert(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v114 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v116 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1104)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v119
	v123 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = base.I32_wrap_i64(v118) & (v123 - int32(1))
	v128 = base.I64_extend_i32_s(v123)
	v129 = base.I64_div_u_s(v118, v128)
	v131 = base.I64_div_u_s(int64(4294967296), v128)
	v132 = base.I64_div_u_s(v129, v131)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+132)) = uint32(v132)
	v135 = v129 - v131*v132
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+136)) = uint32(v135)
	v143 = F_pg_snprintf(m, v13+int32(208), int32(1024), int32(233288), v13+int32(128))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v148 = F_AllocateFile(m, v13+int32(208), int32(32600))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	if v148 == int32(0) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v153 = F_build_backup_content(m, l0, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v153
	v159 = F_pg_fprintf(m, v148, int32(206200), v13+int32(112))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v153)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v163 = F_fflush(m, v148)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	if v163 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v148)+76))
	if v165 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if int32(base.Ui32(v170)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L4
	} else {
		goto L46
	}
L42:
	;
	goto L41
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v170 = v168
	goto L42
L44:
	;
	goto L45
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v170 = v169
	goto L42
L46:
	;
	v175 = F_FreeFile(m, v148)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v175 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v178 = F_AllocateDir(m, int32(308818))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v181 = F_ReadDir(m, v178, int32(308818))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	if v181 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v186 = v181
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_FreeDir(m, v178)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L13
	} else {
		goto L101
	}
L54:
	;
	v194 = v186 + int32(19)
	v195 = F_strlen(m, v194)
	mBase = m.M
	if base.Ui32(v195) < base.Ui32(int32(25)) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v347 = F_ReadDir(m, v178, int32(308818))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L99
	}
L57:
	;
	v198 = int32(537590)
	v202 = m.G0
	v204 = v202 - int32(32)
	v205 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v204)+24)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v204)+16)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v204)+8)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v205
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _consts[271])))
	if v213 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v281 != int32(24) {
		goto L56
	} else {
		goto L79
	}
L59:
	;
	v281 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[272])))
	if v217 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v221 = v194
	goto L65
L63:
	;
	goto L64
L64:
	;
	v231 = v198
	v232 = v213
	goto L68
L65:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v227 == v213 {
		v221 = v221 + int32(1)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v281 = v221 - v194
	goto L58
L67:
	;
	goto L66
L68:
	;
	v239 = v204 + int32(base.Ui32(v232)>>(uint(int32(3))%32))&int32(28)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v241 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v240 | v241<<(uint(v232)%32)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v245 != 0 {
		v231 = v231 + v241
		v232 = v245
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v248 == int32(0) {
		v273 = v194
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v281 = v273 - v194
	goto L58
L72:
	;
	v252 = v194
	v253 = v248
	goto L73
L73:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(base.Ui32(v253)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v261)>>(uint(v253)%32))&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v273 = v269
	goto L71
L75:
	;
	v273 = v252
	goto L71
L76:
	;
	goto L77
L77:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	v269 = v252 + int32(1)
	if v267 != 0 {
		v252 = v269
		v253 = v267
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v286 = v194 + v195 - int32(7)
	v287 = int32(233312)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _consts[360])))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v291 == int32(0) {
		v310 = v290
		v311 = v291
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v311-v310 != 0 {
		goto L56
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	if v290 != v291 {
		v310 = v290
		v311 = v291
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v295 = v286
	v296 = v287
	goto L84
L84:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if v300 == int32(0) {
		v310 = v299
		v311 = v300
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v310 = v299
	v311 = v300
	goto L81
L86:
	;
	v303 = int32(1)
	if v299 == v300 {
		v295 = v295 + v303
		v296 = v296 + v303
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v313 = F_XLogArchiveCheckDone(m, v194)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	if v313 == int32(0) {
		goto L56
	} else {
		goto L90
	}
L90:
	;
	v319 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	if v319 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v194
	F_errmsg_internal(m, int32(715833), v13+int32(80))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v194
	v339 = F_pg_snprintf(m, v13+int32(1232), int32(1031), int32(177132), v13-int32(-64))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L97
	}
L95:
	;
	F_errfinish(m, int32(498539), int32(4174), int32(13135))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v343 = F_unlink(m, v13+int32(1232))
	mBase = m.M
	F_XLogArchiveCleanup(m, v194)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	goto L56
L99:
	;
	if v347 != 0 {
		v186 = v347
		goto L54
	} else {
		goto L100
	}
L100:
	;
	goto L55
L101:
	;
	goto L17
L102:
	;
	m.G0 = v13 + int32(2272)
	return
L103:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	if v26 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	F_errmsg(m, v538, int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L13
	} else {
		goto L144
	}
L105:
	;
	v379 = base.B2i32(v374 == int32(2))
	goto L107
L106:
	;
	v379 = base.B2i32(int32(0) < v374)
	goto L107
L107:
	;
	if v379 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v383
	v388 = int64(*(*int32)(unsafe.Add(mBase, _consts[189])))
	v389 = base.I64_div_u_s(v382-int64(1), v388)
	v391 = base.I64_div_u_s(int64(4294967296), v388)
	v392 = base.I64_div_u_s(v389, v391)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v392)
	v395 = v389 - v391*v392
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v395)
	v403 = F_pg_snprintf(m, v13+int32(1232), int32(64), int32(510531), v13+int32(32))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L13
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v528 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L13
	} else {
		goto L142
	}
L111:
	;
	v405 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v406
	v410 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = base.I32_wrap_i64(v405) & (v410 - int32(1))
	v415 = base.I64_extend_i32_s(v410)
	v416 = base.I64_div_u_s(v405, v415)
	v418 = base.I64_div_u_s(int64(4294967296), v415)
	v419 = base.I64_div_u_s(v416, v418)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v419)
	v422 = v416 - v418*v419
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v422)
	v430 = F_pg_snprintf(m, v13+int32(144), int32(64), int32(233295), v13+int32(16))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v433 = int32(0)
	v438 = v433
	v439 = v433
	v440 = int32(60)
	goto L113
L113:
	;
	v447 = F_XLogArchiveIsBusy(m, v13+int32(1232))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L13
	} else {
		goto L116
	}
L114:
	;
	v520 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L13
	} else {
		goto L140
	}
L115:
	;
	goto L114
L116:
	;
	if v447 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v453 = F_XLogArchiveIsBusy(m, v13+int32(144))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v458 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v453 == int32(0) {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if (base.B2i32(v438 < int32(6))|v439)&int32(1) != 0 {
		v482 = v439
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v488 = F_WaitLatch(m, v484, int32(41), int32(1000), int32(134217732))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L13
	} else {
		goto L132
	}
L127:
	;
	v466 = int32(1)
	v469 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L13
	} else {
		goto L128
	}
L128:
	;
	if v469 == int32(0) {
		v482 = v466
		goto L126
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(440972), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(498539), int32(9398), int32(234271))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	v482 = v466
	goto L126
L132:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = int32(0)
	goto L133
L133:
	;
	v495 = v438 + int32(1)
	if v495 < v440 {
		v438 = v495
		v439 = v482
		goto L113
	} else {
		goto L134
	}
L134:
	;
	v498 = v440 << (uint(int32(1)) % 32)
	v501 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	if v501 == int32(0) {
		v438 = v495
		v439 = v482
		v440 = v498
		goto L113
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v495
	F_errmsg(m, int32(678347), v13)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L13
	} else {
		goto L137
	}
L137:
	;
	F_errhint(m, int32(586082), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(498539), int32(9416), int32(234271))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	v438 = v495
	v439 = v482
	v440 = v498
	goto L113
L140:
	;
	if v520 == int32(0) {
		goto L102
	} else {
		goto L141
	}
L141:
	;
	v538 = int32(440927)
	v544 = int32(9421)
	goto L104
L142:
	;
	if v528 == int32(0) {
		goto L102
	} else {
		goto L143
	}
L143:
	;
	v538 = int32(233831)
	v544 = int32(9425)
	goto L104
L144:
	;
	F_errfinish(m, int32(498539), v544, int32(234271))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	goto L102
L146:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(233619), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	F_errhint(m, int32(579408), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L13
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(498539), int32(9196), int32(234271))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(233748), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	F_errhint(m, int32(612088), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(498539), int32(9237), int32(234271))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(233672), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	F_errhint(m, int32(620098), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(498539), int32(9287), int32(234271))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v13 + int32(208)
	F_errmsg(m, int32(299598), v13+int32(48))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(498539), int32(9332), int32(234271))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v13 + int32(208)
	F_errmsg(m, int32(299504), v13+int32(96))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L13
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(498539), int32(9343), int32(234271))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L13
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_b64_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
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
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	v5 = int32(0)
	v12 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v12) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v181 = F__emscripten_memset_bulkmem(m, l2, base.I32_extend8_s(int32(0)), l3)
	mBase = m.M
	goto L44
L2:
	;
	v14 = l0
	v18 = v5
	v19 = l2
	v20 = v5
	v22 = v5
	goto L5
L3:
	;
	v163 = l2
	goto L4
L4:
	;
	return v163 - l2
L5:
	;
	v26 = v14 + int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v27 != int32(61) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if v151 != 0 {
		goto L1
	} else {
		goto L43
	}
L7:
	;
	if v149 != v12 {
		v14 = v149
		v18 = v151
		v19 = v152
		v20 = v153
		v22 = v155
		goto L5
	} else {
		goto L42
	}
L8:
	;
	if l3 < v19-l2+int32(1) {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	v105 = int32(2)
	v106 = v26
	v109 = v22 << (uint(int32(6)) % 32)
	goto L8
L10:
	;
	v94 = v91 + v90<<(uint(int32(6))%32)
	v96 = v87 + int32(1)
	if v96 == int32(4) {
		v105 = v88
		v106 = v89
		v109 = v94
		goto L8
	} else {
		goto L28
	}
L11:
	;
	v87 = int32(3)
	v88 = int32(1)
	v89 = v14 + int32(2)
	v90 = v45
	v91 = v40
	goto L10
L12:
	;
	if base.Ui32(int32(125)) < base.Ui32((v64-int32(1))&int32(255)) {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v31 = v27 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v31) {
		v64 = v27
		v65 = v18
		v66 = v20
		v67 = v26
		v68 = v22
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v40 = int32(0)
	if v20 != 0 {
		v87 = v18
		v88 = v20
		v89 = v26
		v90 = v22
		v91 = v40
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if int32(1)<<(uint(v31)%32)&int32(8388627) == int32(0) {
		v64 = v27
		v65 = v18
		v66 = v20
		v67 = v26
		v68 = v22
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L1
L18:
	;
	switch v18 - int32(2) {
	case 0:
		goto L19
	case 1:
		goto L9
	default:
		goto L1
	}
L19:
	;
	if v26 == v12 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v45 = v22 << (uint(int32(6)) % 32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v46 == int32(61) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v50 = v46 - int32(9)
	if int32(1)<<(uint(v50)%32)&int32(8388627) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = base.B2i32(base.Ui32(v50) <= base.Ui32(int32(23)))
	goto L24
L23:
	;
	v58 = int32(0)
	goto L24
L24:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v64 = v46
	v65 = int32(3)
	v66 = int32(1)
	v67 = v14 + int32(2)
	v68 = v45
	goto L12
L26:
	;
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+uint32(_consts[572]))))
	if v78 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v87 = v65
	v88 = v66
	v89 = v67
	v90 = v68
	v91 = v78
	goto L10
L28:
	;
	v149 = v89
	v151 = v96
	v152 = v19
	v153 = v88
	v155 = v94
	goto L7
L29:
	;
	v115 = int32(base.Ui32(v109) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v115)
	v118 = v19 + int32(1)
	if base.Ui32(v105) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = v105
	goto L32
L31:
	;
	v122 = int32(0)
	goto L32
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l3 < v118-l2+int32(1) {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v134 = v118
	goto L35
L35:
	;
	v135 = int32(0)
	if v105 == v135 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v130 = int32(base.Ui32(v109) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)) = uint8(v130)
	v134 = v19 + int32(2)
	goto L35
L37:
	;
	v149 = v106
	v151 = int32(0)
	v152 = v147
	v153 = v105
	v155 = v135
	goto L7
L38:
	;
	if l3 < v134-l2+int32(1) {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v105) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v147 = v134
	goto L37
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v109)
	v147 = v134 + int32(1)
	goto L37
L42:
	;
	goto L6
L43:
	;
	v163 = v152
	goto L4
L44:
	;
	return int32(-1)
}
func F_pg_backup_stop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[359])))
	v19 = F_get_call_result_type(m, l0, v2, v7+int32(28))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(1) {
			if v15 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(128053), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(546234), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495280), int32(142), int32(234274))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
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
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[364]))
				F_do_pg_backup_stop(m, v28, base.B2i32(v13 != int32(0)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					v36 = F_build_backup_content(m, v34, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[364]))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+1088))
						v41 = F_Int64GetDatum(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v41
							v44 = F_cstring_to_text(m, v36)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _consts[365]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = F_cstring_to_text(m, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v50
									F_pfree(m, v36)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v56 = int32(0)
										*(*int32)(unsafe.Add(mBase, _consts[365])) = v56
										*(*int32)(unsafe.Add(mBase, _consts[364])) = v56
										v62 = *(*int32)(unsafe.Add(mBase, _consts[366]))
										F_MemoryContextDelete(m, v62)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[366])) = int32(0)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
											v73 = F_heap_form_tuple(m, v68, v7+int32(16), v7+int32(12))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
												v76 = F_HeapTupleHeaderGetDatum(m, v75)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(32)
													return v76
												}
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
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(367739), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495280), int32(136), int32(234274))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
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
func F_pg_basetype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_SearchSysCache1(m, int32(82), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_ReleaseCatCache(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L13
	}
L2:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
	return int32(0)
L3:
	;
	return int32(0)
L4:
	;
	if v7 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v15 = v13 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	if v16 != int32(100) {
		v47 = v7
		v48 = v6
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = v15
	v21 = v7
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v47 = v27
	v48 = v23
	goto L1
L9:
	;
	v27 = F_SearchSysCache1(m, int32(82), v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v27 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+79)))
	if v34 == int32(100) {
		v20 = v33
		v21 = v27
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	return v48
}
func F_pg_big5_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = int32(-1)
	v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v9 < int32(0) {
		v12 = int32(2)
	} else {
		v12 = int32(1)
	}
	if l1 < v12 {
		v26 = v6
	} else {
		if v9 == int32(-115) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v16 != int32(32) {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v24 != 0 {
					v25 = v12
				} else {
					v25 = int32(-1)
				}
				v26 = v25
			} else {
				v26 = v6
			}
		} else {
			if int32(0) <= v9 {
				v26 = int32(1)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v24 != 0 {
					v25 = v12
				} else {
					v25 = int32(-1)
				}
				v26 = v25
			}
		}
	}
	return v26
}
func F_pg_cancel_backend(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_signal_backend(m, v7, int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		switch v9 - int32(2) {
		case 0:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(16672), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(427360)
						F_errdetail(m, int32(573770), v5+int32(32))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495233), int32(159), int32(427342))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
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
		case 1:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(16672), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(525854)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
						F_errdetail(m, int32(630016), v5)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495233), int32(145), int32(427342))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
		case 2:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(16672), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(220269)
						F_errdetail(m, int32(588706), v5+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495233), int32(152), int32(427342))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		default:
			m.G0 = v5 + int32(48)
			return base.B2i32(v9 == int32(0))
		}
	}
}
func F_pg_checksum_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v5-int32(2)) {
		if v5 != int32(1) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = m.Env.Pgmem_crc32c(m, v12, l1, l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
		}
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = F_pg_cryptohash_update(m, v15, l1, l2)
		mBase = m.M
		if int32(0) <= v16 {
			return int32(0)
		} else {
			return int32(-1)
		}
	}
}
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_class_aclmask_ext(m, l0, l1, l2, int32(1), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_pg_clean_ascii(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_strlen(m, l0)
	mBase = m.M
	v16 = v12<<(uint(int32(2))%32) | int32(1)
	v17 = F_palloc_extended(m, v16, l1)
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
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v17
L6:
	;
	v23 = l0
	v24 = v21
	v25 = v22
	goto L9
L7:
	;
	v51 = v21
	goto L8
L8:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v17))) = uint8(v58)
	goto L5
L9:
	;
	v30 = v24 + v17
	if base.Ui32((v25-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v51 = v46
	goto L8
L11:
	;
	v46 = v45 + v24
	v48 = v23 + int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 != 0 {
		v23 = v48
		v24 = v46
		v25 = v49
		goto L9
	} else {
		goto L16
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v25
	v40 = F_pg_snprintf(m, v30, v16-v24, int32(29644), v10)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v25)
	v45 = int32(1)
	goto L11
L15:
	;
	v45 = int32(4)
	goto L11
L16:
	;
	goto L10
}
func F_pg_client_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_control_system(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[253]))
				v28 = F_get_controlfile(m, v25, v6+int32(7))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(388029), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(500532), int32(50), int32(289556))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v39
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v43
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
							v48 = F_Int64GetDatum(m, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v50)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v48
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
								v58 = F_Int64GetDatum(m, v53*int64(1000000)-int64(946684800000000))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v60)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v58
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									v68 = F_heap_form_tuple(m, v63, v6+int32(16), v6+int32(12))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										v71 = F_HeapTupleHeaderGetDatum(m, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											m.G0 = v6 + int32(32)
											return v71
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
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(367739), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500532), int32(42), int32(289556))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
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
func F_pg_convert_to(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_DirectFunctionCall1Coll(m, int32(500), v2, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_DirectFunctionCall3Coll(m, int32(1652), v2, v6, v12, v3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v16
		}
	}
}
func F_pg_database_encoding_max_length(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(28))+uint32(_consts[979])))
	return v8
}
func F_pg_ddl_command_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(428032)
			F_errmsg(m, int32(192503), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494551), int32(358), int32(67542))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_dependencies_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(168536)
			F_errmsg(m, int32(192537), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494608), int32(661), int32(279479))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_describe_object(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9|v10 == int32(0) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v36 = int32(0)
		m.G0 = v7 + int32(16)
		return v36
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v9
		v24 = F_getObjectDescription(m, v7+int32(4), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v30 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
				v36 = int32(0)
				m.G0 = v7 + int32(16)
				return v36
			} else {
				v33 = F_cstring_to_text(m, v24)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = v33
					m.G0 = v7 + int32(16)
					return v36
				}
			}
		}
	}
}
func F_pg_detoast_datum_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
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
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v298 int32
	_ = v298
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	if v4 <= l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v458
L2:
	;
	v22 = l2
	v25 = l0
	goto L8
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L37
	} else {
		goto L151
	}
L5:
	;
	m.G0 = v16 - int32(-64)
	goto L1
L6:
	;
	v415 = int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	if v416&v415 != 0 {
		goto L132
	} else {
		goto L133
	}
L7:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v102&int32(3) != int32(2) {
		goto L43
	} else {
		goto L44
	}
L8:
	;
	v33 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v49&int32(254) != int32(2) {
		v101 = v25
		goto L7
	} else {
		goto L41
	}
L10:
	;
	v42 = v22
	v45 = int32(-1)
	goto L12
L11:
	;
	v36 = l1 + v22
	v38 = v33 ^ base.B2i32(v36 < l1)
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v46 != int32(1) {
		v101 = v25
		goto L7
	} else {
		goto L19
	}
L13:
	;
	v39 = int32(-1)
	goto L15
L14:
	;
	v39 = v22
	goto L15
L15:
	;
	if v38 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(-1)
	goto L18
L17:
	;
	v41 = v36
	goto L18
L18:
	;
	v42 = v39
	v45 = v41
	goto L12
L19:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v49 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L9
L21:
	;
	if v49 != int32(18) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2))
	v22 = v42
	v25 = v90
	goto L8
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+6))
	v56 = v54 & int32(1073741823)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2))
	if base.Ui32(v56) < base.Ui32(v57-int32(4)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if int32(0) <= v45 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v88 = F_toast_fetch_datum_slice(m, v25, l1, v42)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L37
	} else {
		goto L40
	}
L28:
	;
	if base.Ui32(v54) <= base.Ui32(int32(1073741823)) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v86 = F_toast_fetch_datum(m, v25)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L37
	} else {
		goto L39
	}
L31:
	;
	v72 = base.I64_div_s(base.I64_extend_i32_s(v45)*int64(9)+int64(7), int64(8))
	v74 = v72 + int64(2)
	v75 = base.I64_extend_i32_s(v56)
	if v74 < v75 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v81 = v56
	goto L33
L33:
	;
	v82 = F_toast_fetch_datum_slice(m, v25, int32(0), v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v77 = v74
	goto L36
L35:
	;
	v77 = v75
	goto L36
L36:
	;
	v81 = base.I32_wrap_i64(v77)
	goto L33
L37:
	;
	return int32(0)
L38:
	;
	v101 = v82
	goto L7
L39:
	;
	v101 = v86
	goto L7
L40:
	;
	v458 = v88
	goto L5
L41:
	;
	v95 = F_detoast_external_attr(m, v25)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v101 = v95
	goto L7
L43:
	;
	v413 = v101
	goto L6
L44:
	;
	goto L45
L45:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if int32(0) <= v45 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v101 == v25 {
		v413 = v409
		goto L6
	} else {
		goto L129
	}
L47:
	;
	v405 = F_pglz_decompress_datum(m, v101)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L37
	} else {
		goto L128
	}
L48:
	;
	v111 = int32(base.Ui32(v107) >> (uint(int32(30)) % 32))
	if base.Ui32(v107&int32(1073741823)) <= base.Ui32(v45) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v385 = int32(base.Ui32(v107) >> (uint(int32(30)) % 32))
	switch v385 {
	case 0:
		goto L47
	case 1:
		goto L123
	default:
		goto L122
	}
L51:
	;
	switch v111 {
	case 0:
		goto L47
	case 1:
		goto L55
	default:
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	switch v111 {
	case 0:
		goto L62
	case 1:
		goto L61
	default:
		goto L60
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L37
	} else {
		goto L57
	}
L55:
	;
	v115 = F_lz4_decompress_datum(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	v409 = v115
	goto L46
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v111
	F_errmsg_internal(m, int32(479352), v14+int32(-32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L37
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(493188), int32(489), int32(286509))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L37
	} else {
		goto L119
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L37
	} else {
		goto L114
	}
L62:
	;
	v134 = F_palloc(m, v45+int32(4))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L37
	} else {
		goto L63
	}
L63:
	;
	v136 = int32(8)
	v137 = v101 + v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v144 = v134 + int32(4)
	v153 = v144 + v45
	v154 = v137 + (int32(base.Ui32(v138)>>(uint(int32(2))%32)) - v136)
	if base.Ui32(v154) <= base.Ui32(v137) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	if v325 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L65:
	;
	goto L64
L66:
	;
	goto L103
L67:
	;
	v298 = v144
	goto L66
L68:
	;
	if base.Ui32(v153) <= base.Ui32(v144) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v157 = v137
	v158 = v144
	goto L70
L70:
	;
	v170 = v157 + int32(1)
	if base.Ui32(v154) <= base.Ui32(v170) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v298 = v284
	goto L66
L72:
	;
	if base.Ui32(v154) <= base.Ui32(v283) {
		v298 = v284
		goto L66
	} else {
		goto L100
	}
L73:
	;
	v283 = v170
	v284 = v158
	goto L72
L74:
	;
	goto L75
L75:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v175 = v158
	v177 = v170
	v183 = v172
	v184 = int32(0)
	goto L76
L76:
	;
	if v183&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v283 = v260
	v284 = v272
	goto L72
L78:
	;
	if base.Ui32(int32(6)) < base.Ui32(v184) {
		v283 = v260
		v284 = v272
		goto L72
	} else {
		goto L97
	}
L79:
	;
	v188 = int32(-1)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v193 = v189&int32(15) + int32(3)
	if v193 != int32(18) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v254)
	v256 = int32(1)
	v260 = v177 + v256
	v272 = v175 + v256
	goto L78
L82:
	;
	v203 = v193
	v204 = v177 + int32(2)
	goto L84
L83:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)))
	v203 = v198 + int32(18)
	v204 = v177 + int32(3)
	goto L84
L84:
	;
	if base.Ui32(v154) < base.Ui32(v204) {
		v325 = v188
		goto L65
	} else {
		goto L85
	}
L85:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	v211 = v206 | v189<<(uint(int32(4))%32)&int32(3840)
	if v211 == int32(0) {
		v325 = v188
		goto L65
	} else {
		goto L86
	}
L86:
	;
	if v175-v144 < v211 {
		v325 = v188
		goto L65
	} else {
		goto L87
	}
L87:
	;
	v216 = v153 - v175
	if v203 < v216 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v218 = v203
	goto L90
L89:
	;
	v218 = v216
	goto L90
L90:
	;
	if v211 < v218 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v221 = v175
	v223 = v211
	v225 = v218
	goto L94
L92:
	;
	v240 = v175
	v242 = v211
	v244 = v218
	goto L93
L93:
	;
	v252 = F___memcpy(m, v240, v240-v242, v244)
	mBase = m.M
	v260 = v204
	v272 = v252 + v244
	goto L78
L94:
	;
	v232 = v225 - v223
	v234 = F___memcpy(m, v221, v221-v223, v223)
	mBase = m.M
	v235 = v234 + v223
	v237 = v223 << (uint(int32(1)) % 32)
	if v237 < v232 {
		v221 = v235
		v223 = v237
		v225 = v232
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v240 = v235
	v242 = v237
	v244 = v232
	goto L93
L96:
	;
	goto L95
L97:
	;
	if base.Ui32(v154) <= base.Ui32(v260) {
		v283 = v260
		v284 = v272
		goto L72
	} else {
		goto L98
	}
L98:
	;
	v276 = int32(1)
	if base.Ui32(v272) < base.Ui32(v153) {
		v175 = v272
		v177 = v260
		v183 = int32(base.Ui32(v183&int32(254)) >> (uint(v276) % 32))
		v184 = v184 + v276
		goto L76
	} else {
		goto L99
	}
L99:
	;
	goto L77
L100:
	;
	if base.Ui32(v284) < base.Ui32(v153) {
		v157 = v283
		v158 = v284
		goto L70
	} else {
		goto L101
	}
L101:
	;
	goto L71
L103:
	;
	goto L104
L104:
	;
	v325 = v298 - v144
	goto L65
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L37
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v325<<(uint(int32(2))%32) + int32(16)
	v409 = v134
	goto L46
L110:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L37
	} else {
		goto L111
	}
L111:
	;
	F_errmsg_internal(m, int32(83592), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L37
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(496836), int32(126), int32(418603))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L37
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L37
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(444347), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L37
	} else {
		goto L116
	}
L116:
	;
	F_errdetail(m, int32(579339), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L37
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(496836), int32(218), int32(418660))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L37
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v111
	F_errmsg_internal(m, int32(479352), v14+int32(-48))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L37
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(493188), int32(532), int32(418631))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L37
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L37
	} else {
		goto L125
	}
L123:
	;
	v386 = F_lz4_decompress_datum(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L37
	} else {
		goto L124
	}
L124:
	;
	v409 = v386
	goto L46
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v385
	F_errmsg_internal(m, int32(479352), v14+int32(-16))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L37
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(493188), int32(489), int32(286509))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L37
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	v409 = v405
	goto L46
L129:
	;
	F_pfree(m, v101)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L37
	} else {
		goto L130
	}
L130:
	;
	v413 = v409
	goto L6
L131:
	;
	if l1 < v430 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v419 = int32(1)
	v429 = v415
	v430 = int32(base.Ui32(v416)>>(uint(v419)%32)) - v419
	goto L131
L133:
	;
	goto L134
L134:
	;
	v423 = int32(4)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v429 = v423
	v430 = int32(base.Ui32(v424)>>(uint(int32(2))%32)) - v423
	goto L131
L135:
	;
	v433 = v430 - l1
	if v430 < v45 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v439 = v4
	v440 = int32(0)
	goto L137
L137:
	;
	v442 = v440 + int32(4)
	v443 = F_palloc(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L37
	} else {
		goto L144
	}
L138:
	;
	v435 = v433
	goto L140
L139:
	;
	v435 = v42
	goto L140
L140:
	;
	if v42 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v438 = v433
	goto L143
L142:
	;
	v438 = v435
	goto L143
L143:
	;
	v439 = l1
	v440 = v438
	goto L137
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v442 << (uint(int32(2)) % 32)
	if v440 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v413 == v25 {
		v458 = v443
		goto L5
	} else {
		goto L149
	}
L146:
	;
	v452 = F__emscripten_memcpy_bulkmem(m, v443+int32(4), v413+v429+v439, v440)
	mBase = m.M
	goto L148
L147:
	;
	goto L148
L148:
	;
	goto L145
L149:
	;
	F_pfree(m, v413)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L37
	} else {
		goto L150
	}
L150:
	;
	v458 = v443
	goto L5
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg_internal(m, int32(481974), v16)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L37
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(493188), int32(215), int32(418584))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L37
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_encrypt_iv(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
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
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L113
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v22
	goto L6
L5:
	;
	v28 = v17 + int32(4)
	goto L6
L6:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = F_downcase_truncate_identifier(m, v28, v56, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v31 = int32(4)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v33&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v46 = int32(1)
	if v27 != 0 {
		v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v42 = v31
	goto L13
L12:
	;
	v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
	goto L13
L13:
	;
	if v33 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v31
	goto L16
L15:
	;
	v45 = v42
	goto L16
L16:
	;
	v56 = v45
	goto L7
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v62 = F_px_find_combo(m, v58, v14+int32(28))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v62 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v58)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L95
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = F_pg_detoast_datum_packed(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v73 = F_pg_detoast_datum_packed(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v76 = F_pg_detoast_datum_packed(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v78 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v109 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v81 = int32(4)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v83&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v96 = int32(1)
	if v78&v96 != 0 {
		v108 = int32(base.Ui32(v78)>>(uint(v96)%32)) - v96
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v92 = v81
	goto L33
L32:
	;
	v92 = base.B2i32(v83 == int32(18)) << (uint(v81) % 32)
	goto L33
L33:
	;
	if v83 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v95 = v81
	goto L36
L35:
	;
	v95 = v92
	goto L36
L36:
	;
	v108 = v95
	goto L27
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v140 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v112 = int32(4)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v114&int32(254) == int32(2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v127 = int32(1)
	if v109&v127 != 0 {
		v139 = int32(base.Ui32(v109)>>(uint(v127)%32)) - v127
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v123 = v112
	goto L44
L43:
	;
	v123 = base.B2i32(v114 == int32(18)) << (uint(v112) % 32)
	goto L44
L44:
	;
	if v114 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v112
	goto L47
L46:
	;
	v126 = v123
	goto L47
L47:
	;
	v139 = v126
	goto L38
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v139 = int32(base.Ui32(v133)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v172 = m.T0[v171].(func(*base.Module, int32, int32) int32)(m, v68, v108)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L60
	}
L50:
	;
	v143 = int32(4)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v145&int32(254) == int32(2) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v158 = int32(1)
	if v140&v158 != 0 {
		v170 = int32(base.Ui32(v140)>>(uint(v158)%32)) - v158
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v154 = v143
	goto L55
L54:
	;
	v154 = base.B2i32(v145 == int32(18)) << (uint(v143) % 32)
	goto L55
L55:
	;
	if v145 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v157 = v143
	goto L58
L57:
	;
	v157 = v154
	goto L58
L58:
	;
	v170 = v157
	goto L49
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v170 = int32(base.Ui32(v164)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v172
	v177 = F_palloc(m, v172+int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v179 = int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v181&v179 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v184 = v179
	goto L64
L63:
	;
	v184 = int32(4)
	goto L64
L64:
	;
	v186 = int32(1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v188&v186 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v191 = v186
	goto L67
L66:
	;
	v191 = int32(4)
	goto L67
L67:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v194 = m.T0[v193].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v68, v73+v184, v139, v76+v191, v170)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v194 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	m.T0[v196].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v199 = int32(1)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v201&v199 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v299 = v194
	goto L1
L73:
	;
	v204 = v199
	goto L75
L74:
	;
	v204 = int32(4)
	goto L75
L75:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v211 = m.T0[v210].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v68, v70+v204, v108, v177+int32(4), v14+int32(28))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	m.T0[v213].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v211 != 0 {
		v299 = v211
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v216<<(uint(int32(2))%32) + int32(16)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v222 != v70 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v70)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v226 != v73 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v73)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v230 != v76 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v76)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v234 != v17 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v17)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v14 + int32(32)
	return v177
L94:
	;
	goto L93
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	if v62 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v58
	F_errmsg(m, int32(205804), v14+int32(16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L111
	}
L98:
	;
	v282 = int32(315695)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v259 = int32(4395680)
	goto L102
L101:
	;
	v282 = v277
	goto L97
L102:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	if v62 != v262 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v277 = v274
	goto L101
L104:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	if v265 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	v282 = int32(414290)
	goto L97
L108:
	;
	goto L109
L109:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v62 != v270 {
		v259 = v259 + int32(16)
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v277 = v265
	goto L101
L111:
	;
	F_errfinish(m, int32(496479), int32(513), int32(227614))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	if v299 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v340
	F_errmsg(m, int32(200837), v14)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L129
	}
L116:
	;
	v340 = int32(315695)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v317 = int32(4395680)
	goto L120
L119:
	;
	v340 = v335
	goto L115
L120:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	if v299 != v320 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v335 = v332
	goto L119
L122:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)+20))
	if v323 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	v340 = int32(414290)
	goto L115
L126:
	;
	goto L127
L127:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317)+16))
	if v299 != v328 {
		v317 = v317 + int32(16)
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v335 = v323
	goto L119
L129:
	;
	F_errfinish(m, int32(496479), int32(387), int32(35638))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_euccn_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v5 {
		v8 = int32(1)
	} else {
		v8 = int32(2)
	}
	if v5&int32(254) == int32(142) {
		v13 = int32(3)
	} else {
		v13 = v8
	}
	return v13
}
func F_pg_eucjp2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(0)
	return v77
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v61
	v66 = v18 + int32(1)
	v68 = v14 + int32(4)
	v69 = v15 + v62
	if int32(0) < v69 {
		v13 = v63
		v14 = v68
		v15 = v69
		v18 = v66
		goto L4
	} else {
		goto L18
	}
L8:
	;
	if v19 == int32(0) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L13
	}
L9:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if v15 == int32(1) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v61 = v24 | int32(36352)
	v62 = int32(-2)
	v63 = v13 + int32(2)
	goto L7
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v36 = v32<<(uint(int32(8))%32) | int32(9371648)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v61 = v36 | v38
	v62 = int32(-3)
	v63 = v13 + int32(3)
	goto L7
L13:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v15 == int32(1) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v61 = v19
	v62 = int32(-1)
	v63 = v13 + int32(1)
	goto L7
L17:
	;
	v51 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v61 = v51 | v53
	v62 = int32(-2)
	v63 = v13 + int32(2)
	goto L7
L18:
	;
	v73 = v68
	v77 = v66
	goto L6
}
func F_pg_euctw_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v15 = int32(4)
		return v15
	case 1:
		return int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v15 = v14
		return v15
	}
}
func F_pg_euctw_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	if l1 <= int32(0) {
		v63 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v14 = base.I32_extend8_s(v13)
	if int32(0) <= v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v63 = v57
	goto L1
L5:
	;
	v57 = v56 + v10
	v58 = v9 - v56
	if int32(0) < v58 {
		v9 = v58
		v10 = v57
		goto L3
	} else {
		goto L18
	}
L6:
	;
	if v14 == int32(0) {
		v63 = v10
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	switch v13 - int32(142) {
	case 0:
		goto L11
	case 1:
		v63 = v10
		goto L1
	default:
		goto L10
	}
L9:
	;
	v56 = int32(1)
	goto L5
L10:
	;
	if v9 == int32(1) {
		v63 = v10
		goto L1
	} else {
		goto L16
	}
L11:
	;
	if base.Ui32(v9) < base.Ui32(int32(4)) {
		v63 = v10
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32((v24+int32(88))&int32(255)) < base.Ui32(int32(249)) {
		v63 = v10
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(93)) < base.Ui32((v31+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)))
	if base.Ui32(int32(94)) <= base.Ui32((v38+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v56 = int32(4)
	goto L5
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v48+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v56 = int32(2)
	goto L5
L18:
	;
	goto L4
}
func F_pg_filenode_relation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v5 = F_RelidByRelfilenumber(m, v4, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			if v5 != 0 {
				v13 = v5
			} else {
				v10 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
				v13 = int32(0)
			}
			return v13
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		v13 = int32(0)
		return v13
	}
}
func F_pg_fsync(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v4 != int32(1) {
		v18 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v18
L2:
	;
	goto L3
L3:
	;
	v9 = F_fsync(m, l0)
	mBase = m.M
	if v9 != int32(-1) {
		v18 = v9
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v18 = int32(-1)
	goto L1
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v13 == int32(27) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
}
func F_pg_gb18030_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v4 {
		return int32(1)
	} else {
		v10 = v4 & int32(255)
		if int32(4) <= l1 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
				v46 = int32(-1)
				if v10 == int32(128) {
					v66 = v46
					return v66
				} else {
					if v10 == int32(255) {
						v66 = v46
						return v66
					} else {
						v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						if base.Ui32((v51+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
							return int32(2)
						} else {
							if int32(-2) < v51 {
								v64 = int32(-1)
							} else {
								v64 = int32(2)
							}
							v66 = v64
							return v66
						}
					}
				}
			} else {
				if v10 == int32(128) {
					return int32(-1)
				} else {
					if v10 == int32(255) {
						return int32(-1)
					} else {
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
						if base.Ui32((v24+int32(1))&int32(255)) < base.Ui32(int32(130)) {
							return int32(-1)
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
							if base.Ui32(int32(10)) <= base.Ui32((v31-int32(48))&int32(255)) {
								return int32(-1)
							} else {
								return int32(4)
							}
						}
					}
				}
			}
		} else {
			if int32(2) <= l1 {
				v46 = int32(-1)
				if v10 == int32(128) {
					v66 = v46
					return v66
				} else {
					if v10 == int32(255) {
						v66 = v46
						return v66
					} else {
						v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						if base.Ui32((v51+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
							return int32(2)
						} else {
							if int32(-2) < v51 {
								v64 = int32(-1)
							} else {
								v64 = int32(2)
							}
							v66 = v64
							return v66
						}
					}
				}
			} else {
				return int32(-1)
			}
		}
	}
}
func F_pg_get_object_address(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
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
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
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
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(272)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_text_to_cstring(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = v2
	goto L15
L5:
	;
	switch v70 - int32(2) {
	case 0, 1:
		goto L97
	default:
		goto L95
	case 3, 9, 11, 30, 41:
		goto L99
	case 22, 24:
		goto L98
	case 23:
		goto L96
	case 29, 48:
		goto L100
	}
L6:
	;
	v322 = F_textarray_to_strvaluelist(m, v24)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L85
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L81
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L77
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L73
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L69
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L61
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L15:
	;
	v37 = v31 << (uint(int32(3)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[443])))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v44 == int32(0) {
		v63 = v43
		v64 = v44
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[444])))
	if v70 < int32(0) {
		goto L13
	} else {
		goto L29
	}
L17:
	;
	if v64-v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	goto L17
L19:
	;
	if v43 != v44 {
		v63 = v43
		v64 = v44
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v48 = v40
	v49 = v16
	goto L21
L21:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v53 == int32(0) {
		v63 = v52
		v64 = v53
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v63 = v52
	v64 = v53
	goto L18
L23:
	;
	v56 = int32(1)
	if v52 == v53 {
		v48 = v48 + v56
		v49 = v49 + v56
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v67 = v31 + int32(1)
	if v67 != int32(59) {
		v31 = v67
		goto L15
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L16
L28:
	;
	goto L14
L29:
	;
	switch v70 - int32(5) {
	case 0, 7, 8, 38, 44:
		goto L33
	default:
		goto L31
	case 17:
		goto L32
	}
L30:
	;
	if v70&int32(2147483646) == int32(2) {
		goto L46
	} else {
		goto L47
	}
L31:
	;
	v120 = F_textarray_to_strvaluelist(m, v21)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L44
	}
L32:
	;
	F_deconstruct_array_builtin(m, v21, int32(25), v13+int32(256), v13+int32(268), v13+int32(248))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L39
	}
L33:
	;
	F_deconstruct_array_builtin(m, v21, int32(25), v13+int32(256), v13+int32(268), v13+int32(248))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+248))
	if v84 != int32(1) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 == int32(1) {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v13)+256))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = F_text_to_cstring(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v96 = F_typeStringToTypeName(m, v93, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v124 = v2
	v125 = v96
	goto L30
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+248))
	if v107 != int32(1) {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 == int32(1) {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v13)+256))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = F_text_to_cstring(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v118 = F_makeFloat(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v319 = v118
	v320 = v2
	v321 = v2
	goto L6
L44:
	;
	if v120 == int32(0) {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v124 = v120
	v125 = v2
	goto L30
L46:
	;
	F_deconstruct_array_builtin(m, v24, int32(25), v13+int32(256), v13+int32(268), v13+int32(248))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	switch v70 - int32(1) {
	case 0, 4, 18, 24, 28, 33:
		goto L46
	default:
		v319 = int32(0)
		v320 = v124
		v321 = v125
		goto L6
	}
L48:
	;
	v143 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+248))
	if v145 <= v143 {
		v326 = v143
		v327 = v143
		v328 = v124
		v332 = v125
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v151 = v143
	v154 = int32(0)
	goto L50
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v154))))
	if v161 == int32(1) {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	v326 = v174
	v327 = v143
	v328 = v124
	v332 = v125
	goto L5
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v13)+256))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164+v154<<(uint(int32(2))%32))))
	v169 = F_text_to_cstring(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v172 = F_typeStringToTypeName(m, v169, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v174 = F_lappend(m, v151, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v177 = v154 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+248))
	if v177 < v178 {
		v151 = v174
		v154 = v177
		goto L50
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+208)) = v16
	F_errmsg(m, int32(714293), v13+int32(208))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(494120), int32(2620), int32(330232))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16
	F_errmsg(m, int32(714323), v13)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(494120), int32(2132), int32(128384))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+176)) = int32(1)
	F_errmsg(m, int32(467630), v13+int32(176))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(494120), int32(2151), int32(128384))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(152625), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(494120), int32(2155), int32(128384))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = int32(1)
	F_errmsg(m, int32(467630), v13+int32(192))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(494120), int32(2168), int32(128384))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(302966), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(494120), int32(2172), int32(128384))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(1)
	F_errmsg(m, int32(467981), v13+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(494120), int32(2181), int32(128384))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(152625), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(494120), int32(2210), int32(128384))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v326 = v322
	v327 = v319
	v328 = v320
	v332 = v321
	goto L5
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L171
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L167
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L163
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L159
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L155
	}
L95:
	;
	switch v70 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 33, 36, 38, 42:
		goto L133
	case 1, 19, 25, 29, 34:
		goto L124
	case 2, 3:
		goto L127
	case 4, 6, 7, 8, 10, 18, 20, 23, 24, 26, 28, 35, 37, 39, 40, 41, 44, 45, 46, 47, 48, 51:
		goto L126
	case 5, 13, 43:
		goto L131
	case 11:
		goto L128
	case 12, 49:
		goto L132
	default:
		v455 = v327
		goto L125
	case 31, 50:
		goto L129
	case 32:
		goto L130
	}
L96:
	;
	if v326 == int32(0) {
		goto L92
	} else {
		goto L121
	}
L97:
	;
	if v328 == int32(0) {
		goto L93
	} else {
		goto L119
	}
L98:
	;
	if v328 != 0 {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	if v326 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	if v328 == int32(0) {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v338 != int32(1) {
		goto L94
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v341 == int32(1) {
		goto L95
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = int32(1)
	F_errmsg(m, int32(467590), v13+int32(112))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(494120), int32(2244), int32(128384))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if int32(1) < v363 {
		goto L95
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = int32(2)
	F_errmsg(m, int32(467981), v13+int32(128))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(494120), int32(2251), int32(128384))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v387 <= int32(2) {
		goto L93
	} else {
		goto L120
	}
L120:
	;
	goto L96
L121:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v392 != int32(2) {
		goto L92
	} else {
		goto L122
	}
L122:
	;
	goto L95
L123:
	;
	F_get_object_address(m, v13+int32(256), v70, v480, v13+int32(248), int32(1), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L146
	}
L124:
	;
	v473 = F_palloc0(m, int32(20))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L145
	}
L125:
	;
	if v455 != 0 {
		v480 = v455
		goto L123
	} else {
		goto L141
	}
L126:
	;
	v455 = v328
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+216)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v13)+220)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v326
	v452 = F_list_make2_impl(m, v13+int32(92), v13+int32(88))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L140
	}
L128:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v442 = F_lcons(m, v441, v328)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L139
	}
L129:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+228)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+224)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v430
	v438 = F_list_make2_impl(m, v13+int32(84), v13+int32(80))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L138
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+236)) = v328
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+232)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v416
	v424 = F_list_make2_impl(m, v13+int32(76), v13+int32(72))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L137
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+244)) = v332
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+240)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v404
	v412 = F_list_make2_impl(m, v13+int32(68), v13-int32(-64))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L136
	}
L132:
	;
	v455 = v332
	goto L125
L133:
	;
	if v328 == int32(0) {
		goto L91
	} else {
		goto L134
	}
L134:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v397 != int32(1) {
		goto L91
	} else {
		goto L135
	}
L135:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v455 = v401
	goto L125
L136:
	;
	v455 = v412
	goto L125
L137:
	;
	v455 = v424
	goto L125
L138:
	;
	v455 = v438
	goto L125
L139:
	;
	v455 = v442
	goto L125
L140:
	;
	v455 = v452
	goto L125
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
	F_errmsg_internal(m, int32(485201), v13+int32(32))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(494120), int32(2363), int32(128384))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+8)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = int32(153)
	v480 = v473
	goto L123
L146:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v13)+264))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v13)+260))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v13)+256))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v13)+248))
	if v493 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_relation_close(m, v493, int32(1))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v500 = F_get_call_result_type(m, l0, int32(0), v13+int32(268))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	if v500 != int32(1) {
		goto L90
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+264)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v13)+260)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v13)+256)) = v492
	v507 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+255)) = uint8(v507)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+253)) = uint16(v507)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	v516 = F_heap_form_tuple(m, v511, v13+int32(256), v13+int32(253))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v519 = F_HeapTupleHeaderGetDatum(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	m.G0 = v13 + int32(272)
	return v519
L155:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = int32(1)
	F_errmsg(m, int32(467630), v13+int32(96))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(494120), int32(2233), int32(128384))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = int32(3)
	F_errmsg(m, int32(467981), v13+int32(144))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(494120), int32(2258), int32(128384))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+160)) = int32(2)
	F_errmsg(m, int32(467590), v13+int32(160))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(494120), int32(2265), int32(128384))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(1)
	F_errmsg(m, int32(467630), v13+int32(48))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(494120), int32(2317), int32(128384))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errmsg_internal(m, int32(367739), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(494120), int32(2373), int32(128384))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_replica_identity_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_table_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_RelationGetReplicaIndex(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_sequence_close(m, v6, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v10 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				} else {
				}
				return v10
			}
		}
	}
}
func F_pg_get_triggerdef_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
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
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	v13 = m.G0
	v15 = v13 - int32(320)
	m.G0 = v15
	v19 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v15+int32(256), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(1)
	v36 = F_systable_beginscan(m, v19, int32(2702), v31, int32(0), v31, v15+int32(256))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v15 + int32(320)
	return v668
L5:
	;
	v38 = F_systable_getnext(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
	F_initStringInfo(m, v15+int32(304))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_sequence_close(m, v19, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v668 = int32(0)
	goto L4
L12:
	;
	v53 = v47 + v48
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	v57 = F_quote_identifier(m, v53+int32(12))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v57
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = int32(744429)
	goto L16
L15:
	;
	v62 = int32(757756)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v62
	F_appendStringInfo(m, v15+int32(304), int32(738629), v15+int32(112))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+80)))
	v74 = v72 & int32(66)
	switch v74 {
	case 0:
		goto L19
	case 1:
		goto L21
	case 2:
		v95 = int32(539944)
		goto L18
	default:
		goto L22
	}
L18:
	;
	F_appendStringInfoString(m, v15+int32(304), v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v95 = int32(525792)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	switch v74 - int32(65) {
	case 0:
		goto L21
	case 1:
		goto L20
	default:
		goto L23
	}
L23:
	;
	v95 = int32(537449)
	goto L18
L24:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+80)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v82
	F_errmsg_internal(m, int32(484137), v15+int32(96))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(494353), int32(957), int32(220423))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v101 = v15 + int32(304)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+80)))
	if v106&int32(4) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if l1 != 0 {
		goto L64
	} else {
		goto L65
	}
L29:
	;
	F_appendStringInfoString(m, v101, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L62
	}
L30:
	;
	if v106&int32(32) == int32(0) {
		goto L28
	} else {
		goto L61
	}
L31:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+80)))
	if v219&int32(32) == int32(0) {
		goto L28
	} else {
		goto L60
	}
L32:
	;
	F_appendStringInfoString(m, v101, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	if v106&int32(16) == int32(0) {
		goto L30
	} else {
		goto L44
	}
L34:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+80)))
	if v130&int32(16) == int32(0) {
		goto L31
	} else {
		goto L43
	}
L35:
	;
	F_appendStringInfoString(m, v101, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L36:
	;
	if v106&int32(8) == int32(0) {
		goto L33
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(517985))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v127 = int32(538652)
	goto L35
L40:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+80)))
	if v121&int32(8) == int32(0) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v127 = int32(538590)
	goto L35
L42:
	;
	goto L34
L43:
	;
	v141 = int32(539283)
	goto L32
L44:
	;
	v141 = int32(539318)
	goto L32
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v144 <= int32(0) {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(745177))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v152 <= int32(0) {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+124)))
	v160 = F_get_attname(m, v157, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v162 = F_quote_identifier(m, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_appendStringInfoString(m, v15+int32(304), v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v166 < int32(2) {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v175 = int32(1)
	goto L53
L53:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(746514))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L31
L55:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(124)+v175<<(uint(int32(1))%32)))))
	v197 = F_get_attname(m, v191, v195, int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v199 = F_quote_identifier(m, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v15+int32(304), v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v204 = v175 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v204 < v205 {
		v175 = v204
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v242 = int32(539337)
	goto L29
L61:
	;
	v242 = int32(539340)
	goto L29
L62:
	;
	goto L28
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v263
	F_appendStringInfo(m, v15+int32(304), int32(738650), v15+int32(80))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L69
	}
L64:
	;
	v259 = F_generate_relation_name(m, v257, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v261 = F_generate_qualified_relation_name(m, v257)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v263 = v259
	goto L63
L68:
	;
	v263 = v261
	goto L63
L69:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	if v272 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v53)+84))
	if v273 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v312 = F_fastgetattr_3(m, v38, int32(18), v309, v15+int32(255))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L88
	}
L73:
	;
	v275 = F_generate_relation_name(m, v273, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+96)))
	if v285 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v275
	F_appendStringInfo(m, v15+int32(304), int32(738658), v15-int32(-64))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(744424))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(744186))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+97)))
	if v302 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v303 = int32(745479)
	goto L85
L84:
	;
	v303 = int32(745332)
	goto L85
L85:
	;
	F_appendStringInfoString(m, v15+int32(304), v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L72
L87:
	;
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+80)))
	if v359&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L88:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)))
	if v314 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v315 = int32(0)
	goto L91
L90:
	;
	v315 = v312
	goto L91
L91:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v321 = F_fastgetattr_3(m, v38, int32(19), v318, v15+int32(255))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)))
	if v323 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v324 = int32(0)
	goto L95
L94:
	;
	v324 = v321
	goto L95
L95:
	;
	if v315|v324 == int32(0) {
		goto L87
	} else {
		goto L96
	}
L96:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(745147))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v315 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v333 = F_quote_identifier(m, v315)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v324 == int32(0) {
		goto L87
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v333
	F_appendStringInfo(m, v15+int32(304), int32(738612), v15+int32(48))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v345 = F_quote_identifier(m, v324)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v345
	F_appendStringInfo(m, v15+int32(304), int32(738595), v15+int32(32))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L87
L106:
	;
	v362 = int32(744325)
	goto L108
L107:
	;
	v362 = int32(744460)
	goto L108
L108:
	;
	F_appendStringInfoString(m, v15+int32(304), v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v369 = F_fastgetattr_3(m, v38, int32(17), v366, v15+int32(255))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)))
	if v371 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(687115))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	v510 = int32(0)
	v516 = F_generate_function_name(m, v509, v510, v510, v510, v510, v510, v510)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L132
	}
L114:
	;
	v379 = F_text_to_cstring(m, v369)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v381 = F_stringToNode(m, v379)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v384 = F_get_rel_relkind(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v387 = F_palloc0(m, int32(136))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+12)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = int32(101)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+21)) = uint8(v384)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+16)) = v393
	v400 = F_makeAlias(m, int32(430255), v389)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v400
	v404 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v387)+124)) = uint16(v404)
	v406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+20)) = uint8(v406)
	v409 = F_palloc0(m, int32(136))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v411 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+12)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = int32(101)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v409)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v409)+21)) = uint8(v384)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+16)) = v415
	v422 = F_makeAlias(m, int32(32193), v411)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+8)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v409)+4)) = v422
	v426 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v409)+124)) = uint16(v426)
	v428 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409)+20)) = uint8(v428)
	v435 = F__emscripten_memset_bulkmem(m, v15+int32(136), base.I32_extend8_s(v428), int32(76))
	mBase = m.M
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v409
	v444 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v444
	F_set_rtable_names(m, v15+int32(132), v446, v446)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_set_simple_column_names(m, v15+int32(132))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v462 = v15 + int32(132)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v15 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v462
	v473 = F_list_make1_impl(m, int32(1), v15+int32(20))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v475 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+244)) = uint8(v475)
	v477 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+228)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v15)+220)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+216)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v477
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+247)) = uint8(v477)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+245)) = uint16(v475)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+236)) = int64(34359738368)
	if l1 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v492 = int32(7)
	goto L129
L128:
	;
	v492 = int32(2)
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+232)) = v492
	F_get_rule_expr(m, v381, v15+int32(212), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(746580))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L113
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v516
	F_appendStringInfo(m, v15+int32(304), int32(685825), v15+int32(16))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v526 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+98)))
	if v526 <= int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_appendStringInfoChar(m, v15+int32(304), int32(41))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L174
	}
L135:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v533 = F_fastgetattr_3(m, v38, int32(16), v530, v15+int32(255))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)))
	if v535 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v538 = F_pg_detoast_datum_packed(m, v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L171
	}
L140:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	v541 = F_pg_detoast_datum_packed(m, v533)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v543 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+98)))
	if v543 <= int32(0) {
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v546 = int32(1)
	if v540&v546 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v550 = v546
	goto L145
L144:
	;
	v550 = int32(4)
	goto L145
L145:
	;
	v553 = v541 + v550
	v554 = int32(0)
	goto L146
L146:
	;
	if v554 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_appendStringInfoString(m, v15+int32(304), int32(746514))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_appendStringInfoChar(m, v15+int32(304), int32(39))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v581 = v553
	goto L153
L153:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	v588 = base.I32_extend8_s(v587)
	if v587 != int32(39) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	F_appendStringInfoChar(m, v15+int32(304), v588)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L170
	}
L156:
	;
	if v587 != int32(92) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	F_appendStringInfoChar(m, v15+int32(304), v588)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L169
	}
L159:
	;
	if v587 != 0 {
		goto L155
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1040])))
	if v618 != 0 {
		goto L155
	} else {
		goto L168
	}
L162:
	;
	F_appendStringInfoChar(m, v15+int32(304), int32(39))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v601 = v553
	goto L164
L164:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v612 = v601 + int32(1)
	if v610 != 0 {
		v601 = v612
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v614 = v554 + int32(1)
	v615 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+98)))
	if v614 < v615 {
		v553 = v612
		v554 = v614
		goto L146
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	goto L134
L168:
	;
	goto L158
L169:
	;
	goto L155
L170:
	;
	v581 = v581 + int32(1)
	goto L153
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(43878), v15)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(494353), int32(1140), int32(220423))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_sequence_close(m, v19, int32(1))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v15)+304))
	v668 = v664
	goto L4
}
func F_pg_get_viewdef_name_ext(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_textToQualifiedNameList(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_makeRangeVarFromNameList(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = int32(0)
				v18 = F_RangeVarGetRelidExtended(m, v12, v14, v14, v14, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v9 != 0 {
						v22 = int32(7)
					} else {
						v22 = int32(2)
					}
					v24 = F_pg_get_viewdef_worker(m, v18, v22, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							return int32(0)
						} else {
							v32 = F_cstring_to_text(m, v24)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v24)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									return v32
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_gmtime(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var __phi90 int32
	_ = __phi90
	var v95 int64
	_ = v95
	var __phi95 int64
	_ = __phi95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v7 = F_emscripten_builtin_malloc(m, int32(23440))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1217])) = v7
	if v7 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v15 = F_tzload(m, int32(1831936), int32(0), v7)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = F_tzparse(m, int32(1831936), v7, int32(1))
	mBase = m.M
	goto L1
L9:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	*(*int32)(unsafe.Add(mBase, _consts[1218])) = v526 + int32(22120)
	return v523
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v38 = v37
	goto L12
L11:
	;
	v38 = v25
	goto L12
L12:
	;
	v44 = v38
	goto L15
L13:
	;
	v81 = int64(86400)
	v82 = base.I64_div_s(v78, v81)
	v85 = v78 - v82*v81
	__phi90 = int32(1970)
	__phi95 = v82
	v90 = __phi90
	v95 = __phi95
	goto L22
L14:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v78 = v74
	v79 = int64(0)
	v80 = int32(0)
	goto L13
L15:
	;
	v54 = v44 - int32(1)
	if v54 < int32(0) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 != v57 {
		v78 = v57
		v79 = v63
		v80 = int32(0)
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v60 = v27 + int32(22632) + v54<<(uint(int32(4))%32)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	if v57 < v61 {
		v44 = v54
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v54 == int32(0) {
		v78 = v57
		v79 = v63
		v80 = base.B2i32(int64(0) < v63)
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v60-int32(8))))
	v78 = v57
	v79 = v63
	v80 = base.B2i32(v72 < v63)
	goto L13
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(61)
	v523 = int32(0)
	goto L9
L22:
	;
	v100 = base.B2i32(v95 < int64(0))
	if v100 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v205 = base.I32_wrap_i64(v95)
	v206 = base.I64_extend_i32_s(v25)
	v208 = v206 - v79 + v85
	if v208 < int64(0) {
		goto L53
	} else {
		goto L54
	}
L24:
	;
	goto L23
L25:
	;
	if v90&int32(3) != 0 {
		v113 = int32(0)
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v95+int64(785979015533)) {
		goto L21
	} else {
		goto L32
	}
L28:
	;
	v118 = int64(*(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[1212]))))
	if v95 < v118 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v108 = base.I32_rem_s(v90, int32(100))
	if v108 != 0 {
		v113 = int32(1)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v110 = base.I32_rem_s(v90, int32(400))
	v113 = base.B2i32(v110 == int32(0))
	goto L28
L31:
	;
	goto L27
L32:
	;
	if v95 < int64(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(-1)
	goto L35
L34:
	;
	v126 = int32(1)
	goto L35
L35:
	;
	v128 = base.I64_div_s(v95, int64(366))
	if base.Ui64(v95+int64(365)) < base.Ui64(int64(731)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v134 = v126
	goto L38
L37:
	;
	v134 = base.I32_wrap_i64(v128)
	goto L38
L38:
	;
	if int32(0) <= v90 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v143 = v134 + v90
	v145 = v143 - int32(1)
	if v145 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	if v134 <= v90^int32(2147483647) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v134 < int32(-2147483648)-v90 {
		goto L21
	} else {
		goto L44
	}
L43:
	;
	goto L21
L44:
	;
	goto L39
L45:
	;
	v177 = v90 - int32(1)
	if v177 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v149 = int32(0) - v143
	v153 = base.I32_div_u_s(v149, int32(100))
	v156 = base.I32_div_u_s(v149, int32(400))
	v169 = int32(base.Ui32(v149)>>(uint(int32(2))%32)) - v153 + v156 ^ int32(-1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v163 = base.I32_div_u_s(v145, int32(100))
	v166 = base.I32_div_u_s(v145, int32(400))
	v169 = int32(base.Ui32(v145)>>(uint(int32(2))%32)) - v163 + v166
	goto L45
L49:
	;
	__phi90 = v143
	__phi95 = (base.I64_extend_i32_s(v143)-base.I64_extend_i32_s(v90))*int64(-365) + v95 - base.I64_extend_i32_s(v169-v201)
	v90 = __phi90
	v95 = __phi95
	goto L22
L50:
	;
	v181 = int32(0) - v90
	v185 = base.I32_div_u_s(v181, int32(100))
	v188 = base.I32_div_u_s(v181, int32(400))
	v201 = int32(base.Ui32(v181)>>(uint(int32(2))%32)) - v185 + v188 ^ int32(-1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v195 = base.I32_div_u_s(v177, int32(100))
	v198 = base.I32_div_u_s(v177, int32(400))
	v201 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - v195 + v198
	goto L49
L53:
	;
	v211 = int64(-86400)
	if base.Ui64(v208) <= base.Ui64(v211) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v235 = v205
	v236 = v208
	goto L55
L55:
	;
	if int64(86400) <= v236 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v214 = v211
	goto L58
L57:
	;
	v214 = v208
	goto L58
L58:
	;
	v216 = v79 + v214 - v85
	v218 = base.I64_extend_i32_u(base.B2i32(v216 != v206))
	v221 = int64(86400)
	v222 = base.I64_div_u_s(v216-(v218+v206), v221)
	v223 = v222 + v218
	v235 = base.I32_wrap_i64(v223) ^ int32(-1) + v205
	v236 = v85 + v223*v221 + v206 - v79 + v221
	goto L55
L59:
	;
	v240 = v236 - int64(172799)
	if base.Ui64(v240) <= base.Ui64(v236) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v257 = v235
	v258 = v236
	goto L61
L61:
	;
	if v257 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v243 = v240
	goto L64
L63:
	;
	v243 = int64(0)
	goto L64
L64:
	;
	v246 = int64(86400)
	v247 = base.I64_div_u_s(v243+int64(86399), v246)
	v257 = v235 + base.I32_wrap_i64(v247) + int32(1)
	v258 = v236 + v247*int64(-86400) - v246
	goto L61
L65:
	;
	v262 = v257
	v265 = v90
	goto L68
L66:
	;
	v297 = v257
	v300 = v90
	goto L67
L67:
	;
	v309 = v297
	v312 = v300
	goto L75
L68:
	;
	if v265 == int32(-2147483648) {
		goto L21
	} else {
		goto L70
	}
L69:
	;
	v297 = v294
	v300 = v278
	goto L67
L70:
	;
	v278 = v265 - int32(1)
	if v278&int32(3) != 0 {
		v288 = int32(0)
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v288<<(uint(int32(2))%32))+uint32(_consts[1212])))
	v294 = v293 + v262
	if v294 < int32(0) {
		v262 = v294
		v265 = v278
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v283 = base.I32_rem_s(v278, int32(100))
	if v283 != 0 {
		v288 = int32(1)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v285 = base.I32_rem_s(v278, int32(400))
	v288 = base.B2i32(v285 == int32(0))
	goto L71
L74:
	;
	goto L69
L75:
	;
	v322 = v312 & int32(3)
	if v322 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1219])) = v312
	if v312 < int32(-2147481748) {
		goto L21
	} else {
		goto L89
	}
L77:
	;
	goto L76
L78:
	;
	if v312 == int32(2147483647) {
		goto L21
	} else {
		goto L88
	}
L79:
	;
	v326 = base.I32_rem_s(v312, int32(100))
	if v326 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	if v309 < int32(365) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v330 = base.I32_rem_s(v312, int32(400))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v330 == int32(0))<<(uint(int32(2))%32))+uint32(_consts[1212])))
	if v309 < v337 {
		goto L77
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v309 < int32(366) {
		goto L77
	} else {
		goto L86
	}
L85:
	;
	v340 = base.I32_rem_s(v312, int32(400))
	v349 = base.B2i32(v340 == int32(0))
	goto L78
L86:
	;
	v349 = int32(1)
	goto L78
L87:
	;
	v349 = int32(0)
	goto L78
L88:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v349<<(uint(int32(2))%32))+uint32(_consts[1212])))
	v309 = v309 - v358
	v312 = v312 + int32(1)
	goto L75
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1220])) = v309
	*(*int32)(unsafe.Add(mBase, _consts[1219])) = v312 - int32(1900)
	v373 = base.I32_rem_s(v312-int32(1970), int32(7))
	if v312 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v400 = int32(0)
	v403 = base.I64_div_u_s(v258, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _consts[1221])) = uint32(v403)
	v410 = int32(7)
	v411 = base.I32_rem_s(v309+v373+v399-int32(473), v410)
	if v411 < v400 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v377 = int32(0) - v312
	v381 = base.I32_div_u_s(v377, int32(100))
	v384 = base.I32_div_u_s(v377, int32(400))
	v399 = int32(base.Ui32(v377)>>(uint(int32(2))%32)) - v381 + v384 ^ int32(-1)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v389 = v312 - int32(1)
	v393 = base.I32_div_u_s(v389, int32(100))
	v396 = base.I32_div_u_s(v389, int32(400))
	v399 = int32(base.Ui32(v389)>>(uint(int32(2))%32)) - v393 + v396
	goto L90
L94:
	;
	v416 = v411 + v410
	goto L96
L95:
	;
	v416 = v411
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1222])) = v416
	v422 = base.I32_wrap_i64(v258 - v403*int64(3600))
	v423 = int32(65535)
	v425 = int32(60)
	v426 = base.I32_div_u_s(v422&v423, v425)
	*(*int32)(unsafe.Add(mBase, _consts[1223])) = v426
	*(*int32)(unsafe.Add(mBase, _consts[1224])) = v80 + (v422-v426*v425)&v423
	if v322 != 0 {
		v444 = int32(0)
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v446 = v444 * int32(48)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446)+uint32(_consts[1213])))
	if v449 <= v309 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v439 = base.I32_rem_s(v312, int32(100))
	if v439 != 0 {
		v444 = int32(1)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v441 = base.I32_rem_s(v312, int32(400))
	v444 = base.B2i32(v441 == int32(0))
	goto L97
L100:
	;
	v451 = v309
	v454 = v449
	v455 = v400
	goto L103
L101:
	;
	v471 = v309
	v475 = v400
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1225])) = v25
	*(*int32)(unsafe.Add(mBase, _consts[1226])) = v475
	*(*int32)(unsafe.Add(mBase, _consts[1227])) = v471 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1228])) = int32(0)
	v523 = int32(4516056)
	goto L9
L103:
	;
	v463 = v451 - v454
	v465 = v455 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v446+int32(1831840)+v465<<(uint(int32(2))%32))))
	if v469 <= v463 {
		v451 = v463
		v454 = v469
		v455 = v465
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v471 = v463
	v475 = v465
	goto L102
L105:
	;
	goto L104
}
func F_pg_indexam_has_property(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(0)
			v13 = F_indexam_property(m, l0, v9, v3, v11, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_indexam_progress_phasename(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_GetIndexAmRoutineByAmId(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+80))
			if v11 != 0 {
				v17 = m.T0[v11].(func(*base.Module, int64) int32)(m, v4)
				mBase = m.M
				if v17 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int32(0)
				} else {
					v24 = F_cstring_to_text(m, v17)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v24
					}
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_itoa(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	if int32(0) <= l0 {
		v12 = l0
		v13 = int32(0)
	} else {
		v7 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
		v12 = int32(0) - l0
		v13 = int32(1)
	}
	v14 = l1 + v13
	if v12 == int32(0) {
		v24 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v24)
		v145 = int32(1)
	} else {
		v30 = int32(1233)
		v35 = int32(base.Ui32((base.I32_clz(v12)^int32(31))*v30+v30) >> (uint(int32(12)) % 32))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32))+uint32(_consts[987])))
		v42 = v35 + base.B2i32(base.Ui32(v40) <= base.Ui32(v12))
		v43 = int32(0)
		if base.Ui32(v12) < base.Ui32(int32(10000)) {
			v89 = v43
			v90 = v12
		} else {
			v47 = v12
			v49 = v43
			for {
				v56 = v14 + v42 - v49
				v57 = int32(4)
				v60 = base.I32_div_u_s(v47, int32(10000))
				v63 = v60*int32(-10000) + v47
				v64 = int32(100)
				v65 = base.I32_div_u_s(v63, v64)
				v66 = int32(1)
				v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65<<(uint(v66)%32))+uint32(_consts[988]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-v57))) = uint16(v70)
				v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v63-v65*v64)<<(uint(v66)%32))+uint32(_consts[988]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-int32(2)))) = uint16(v81)
				v84 = v49 + v57
				if base.Ui32(int32(99999999)) < base.Ui32(v47) {
					v47 = v60
					v49 = v84
					continue
				} else {
					break
				}
				break
			}
			v89 = v84
			v90 = v60
		}
		if base.Ui32(v90) < base.Ui32(int32(100)) {
			v119 = v90
			v120 = v89
		} else {
			v100 = int32(2)
			v102 = int32(65535)
			v104 = int32(100)
			v105 = base.I32_div_u_s(v90&v102, v104)
			v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v90-v105*v104)&v102<<(uint(int32(1))%32))+uint32(_consts[988]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v42-v89-v100))) = uint16(v115)
			v119 = v105
			v120 = v89 | v100
		}
		if base.Ui32(int32(10)) <= base.Ui32(v119) {
			v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119<<(uint(int32(1))%32))+uint32(_consts[988]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v42-v120-int32(2)))) = uint16(v131)
			v145 = v42
		} else {
			v134 = v119 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v134)
			v145 = v42
		}
	}
	v146 = v145 + v13
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v146))) = uint8(v148)
	return v146
}
func F_pg_last_wal_receive_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v5 = F_GetWalRcvFlushRecPtr(m, v3, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int64(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = F_Int64GetDatum(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_log_standby_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
		v17 = base.B2i32(v15 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v17)
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v19 == int32(0) {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if v23 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(730366), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495280), int32(216), int32(86370))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
			v26 = F_LogStandbySnapshot(m)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_Int64GetDatum(m, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v30
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(128133), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(683573)
					F_errhint(m, int32(573539), v5)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495280), int32(211), int32(86370))
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
	}
}
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_consts[1128])))
	v13 = m.T0[v12].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_report_invalid_encoding_db(m, l0, v13, v20)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L11
	}
L2:
	;
	return int32(0)
L3:
	;
	if int32(1) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v13
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v20))))
	if v23 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v27 = v20 + int32(1)
	if v27 != v13 {
		v20 = v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mblen_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v5 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_consts[1128])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l1 < v12 {
			F_report_invalid_encoding_db(m, l0, v12, l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return v12
		}
	}
}
func F_pg_ndistinct_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(109156)
			F_errmsg(m, int32(192537), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493792), int32(396), int32(36360))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v117 int64
	_ = v117
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v142 int64
	_ = v142
	var v150 int32
	_ = v150
	var v157 float64
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int64
	_ = v480
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l0 != int32(100) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(112)
	return v797
L2:
	;
	if l0 == int32(950) {
		v797 = int32(4122324)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _consts[1019]))
	v797 = v795
	goto L1
L5:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[1020]))
	if l0 == v26 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L16
	} else {
		goto L179
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	v797 = v29
	goto L1
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1022]))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v66 = int32(16)
	v70 = (int32(base.Ui32(l0)>>(uint(v66)%32)) ^ l0) * int32(-2048144789)
	v75 = (int32(base.Ui32(v70)>>(uint(int32(13))%32)) ^ v70) * int32(-1028477387)
	v78 = int32(base.Ui32(v75)>>(uint(v66)%32)) ^ v75
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v81 = v65
	v85 = v79
	goto L20
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v64 = v31
	v65 = v32
	goto L12
L14:
	;
	goto L15
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v40 = F_AllocSetContextCreateInternal(m, v35, int32(399815), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1023])) = v40
	v46 = F_MemoryContextAllocZero(m, v40, int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v40
	v53 = F_MemoryContextAllocExtended(m, v40, int32(512), int32(5))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
	*(*int32)(unsafe.Add(mBase, _consts[1022])) = v46
	v64 = v46
	v65 = int32(28)
	goto L12
L20:
	;
	if base.Ui32(v81) <= base.Ui32(v85) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v775 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v775
	v81 = v775
	v85 = v765
	goto L20
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L16
	} else {
		goto L176
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L16
	} else {
		goto L173
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1021])) = v711
	*(*int32)(unsafe.Add(mBase, _consts[1020])) = l0
	v797 = v711
	goto L1
L26:
	;
	v561 = *(*int32)(unsafe.Add(mBase, _consts[1023]))
	v563 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L16
	} else {
		goto L124
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = int32(0)
	v553 = v534
	goto L26
L28:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v521 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v520 + v521
	*(*uint8)(unsafe.Add(mBase, uint32(v513)+12)) = uint8(v521)
	v534 = v513
	goto L27
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L16
	} else {
		goto L121
	}
L30:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	if v96 == int64(4294967296) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v332 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v335 = v334 & v78
	v338 = v333 + v335<<(uint(int32(4))%32)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+12)))
	if v339 == v332 {
		v513 = v338
		goto L28
	} else {
		goto L85
	}
L33:
	;
	v99 = int32(0)
	v101 = int64(2)
	v103 = v96 << (uint(int64(1)) % 64)
	if base.Ui64(v103) <= base.Ui64(v101) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L32
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L16
	} else {
		goto L82
	}
L36:
	;
	v106 = v101
	goto L38
L37:
	;
	v106 = v103
	goto L38
L38:
	;
	v107 = int64(1)
	if v106&(v106-v107) == int64(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v117 = v106
	goto L41
L40:
	;
	v117 = v107 << (uint(int64(64)-base.I64_clz(v106)) % 64)
	goto L41
L41:
	;
	if base.Ui64(v117<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v129 = F_MemoryContextAllocExtended(m, v124, base.I32_wrap_i64(v117)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L16
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L16
	} else {
		goto L79
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v129
	v132 = int64(1)
	if v117&(v117-v132) == int64(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = v117
	goto L48
L47:
	;
	v142 = v132 << (uint(int64(64)-base.I64_clz(v117)) % 64)
	goto L48
L48:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v142<<(uint(int64(4))%64)) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64))) = v142
	v150 = base.I32_wrap_i64(v142) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v150
	v157 = base.F64_mul(base.F64_convert_i64_u(v142), float64(0.9))
	if base.F64_lt(v157, float64(4.294967296e+09))&base.F64_ge(v157, float64(0)) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v142 == int64(4294967296) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v163 = base.I32_trunc_f64_u(v157)
	v165 = v163
	goto L50
L52:
	;
	goto L53
L53:
	;
	v165 = int32(0)
	goto L50
L54:
	;
	v166 = int32(-85899346)
	goto L56
L55:
	;
	v166 = v165
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v166
	if v123 != int64(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v171 = v99
	goto L61
L58:
	;
	goto L59
L59:
	;
	F_pfree(m, v122)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L16
	} else {
		goto L78
	}
L60:
	;
	v201 = v199
	v208 = v99
	goto L66
L61:
	;
	v187 = v122 + v171<<(uint(int32(4))%32)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+12)))
	if v188 != int32(1) {
		v199 = v171
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v199 = int32(0)
	goto L60
L63:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v191&v150 == v171 {
		v199 = v171
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v195 = v171 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v195)) < base.Ui64(v123) {
		v171 = v195
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v217 = v122 + v201<<(uint(int32(4))%32)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+12)))
	if v218 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L59
L68:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v228 = v222
	goto L71
L69:
	;
	goto L70
L70:
	;
	v265 = v201 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v265)) < base.Ui64(v123) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v238 = v228 & v221
	v243 = v129 + v238<<(uint(int32(4))%32)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+12)))
	if v244 != 0 {
		v228 = v238 + int32(1)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
	*(*int64)(unsafe.Add(mBase, uint32(v243))) = v245
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v243)+8)) = v247
	goto L70
L73:
	;
	goto L72
L74:
	;
	v269 = v265
	goto L76
L75:
	;
	v269 = int32(0)
	goto L76
L76:
	;
	v271 = v208 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v271)) < base.Ui64(v123) {
		v201 = v269
		v208 = v271
		goto L66
	} else {
		goto L77
	}
L77:
	;
	goto L67
L78:
	;
	goto L34
L79:
	;
	F_errmsg_internal(m, int32(401023), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(326927), int32(327), int32(341574))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L16
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
	F_errmsg_internal(m, int32(401023), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(326927), int32(327), int32(341574))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L16
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
	v343 = v335
	v345 = v332
	v350 = v338
	goto L86
L86:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	if v357 == v78 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v491 != 0 {
		v711 = v491
		goto L25
	} else {
		goto L120
	}
L88:
	;
	goto L87
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v359 == l0 {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v362 = v343 + int32(1)
	v363 = v334 & v357
	if base.Ui32(v343) < base.Ui32(v363) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v367 = v343 + v365
	goto L95
L94:
	;
	v367 = v343
	goto L95
L95:
	;
	if base.Ui32(v367-v363) < base.Ui32(v345) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v371 = v334 & v362
	v374 = v333 + v371<<(uint(int32(4))%32)
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+12)))
	if v375 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v475 = v345 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v475) {
		goto L115
	} else {
		goto L116
	}
L99:
	;
	v379 = v371
	v385 = int32(0)
	goto L102
L100:
	;
	v413 = v371
	v415 = v374
	goto L101
L101:
	;
	if v343 != v413 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v392 = v385 + int32(1)
	if int32(151) <= v392 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v413 = v405
	v415 = v408
	goto L101
L104:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v395), base.F64_convert_i64_u(v397)), float64(0.1)) != 0 {
		v765 = v395
		goto L22
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v405 = (v379 + int32(1)) & v334
	v408 = v333 + v405<<(uint(int32(4))%32)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)))
	if v409 != 0 {
		v379 = v405
		v385 = v392
		goto L102
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	goto L103
L109:
	;
	v429 = v413
	v431 = v415
	goto L112
L110:
	;
	goto L111
L111:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v469 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v468 + v469
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+12)) = uint8(v469)
	v534 = v350
	goto L27
L112:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v444 = v441 & (v429 - int32(1))
	v447 = v333 + v444<<(uint(int32(4))%32)
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v447)))
	*(*int64)(unsafe.Add(mBase, uint32(v431))) = v448
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v447)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v431)+8)) = v450
	if v343 != v444 {
		v429 = v444
		v431 = v447
		goto L112
	} else {
		goto L114
	}
L113:
	;
	goto L111
L114:
	;
	goto L113
L115:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v478), base.F64_convert_i64_u(v480)), float64(0.1)) != 0 {
		v765 = v478
		goto L22
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v486 = v334 & v362
	v489 = v333 + v486<<(uint(int32(4))%32)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+12)))
	if v490 != 0 {
		v343 = v486
		v345 = v475
		v350 = v489
		goto L86
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v513 = v489
	goto L28
L120:
	;
	v553 = v350
	goto L26
L121:
	;
	F_errmsg_internal(m, int32(462923), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(326927), int32(630), int32(312229))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L16
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
	if v563 == int32(0) {
		goto L24
	} else {
		goto L125
	}
L125:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)+16))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+22)))
	v569 = v567 + v568
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+76)))
	switch v570 - int32(98) {
	case 0:
		goto L127
	case 1:
		goto L129
	default:
		goto L128
	case 7:
		goto L130
	}
L126:
	;
	v598 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v597)+4)) = uint8(v598)
	v604 = F_SysCacheGetAttr(m, int32(16), v563, int32(12), v18+int32(111))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L16
	} else {
		goto L137
	}
L127:
	;
	v595 = F_create_pg_locale_builtin(m, l0, v561)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L16
	} else {
		goto L136
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L16
	} else {
		goto L133
	}
L129:
	;
	v575 = F_create_pg_locale_libc(m, l0, v561)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L16
	} else {
		goto L132
	}
L130:
	;
	v573 = F_create_pg_locale_icu(m)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	v597 = v573
	goto L126
L132:
	;
	v597 = v575
	goto L126
L133:
	;
	v581 = int32(*(*int8)(unsafe.Add(mBase, uint32(v569)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(398107)
	F_errmsg_internal(m, int32(502595), v18+int32(16))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L16
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(499565), int32(1096), int32(398107))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L16
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v597 = v595
	goto L126
L137:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)))
	if v606 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_ReleaseCatCache(m, v563)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L16
	} else {
		goto L172
	}
L139:
	;
	v607 = F_text_to_cstring(m, v604)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L16
	} else {
		goto L140
	}
L140:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+76)))
	if v612 == int32(99) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v615 = int32(8)
	goto L143
L142:
	;
	v615 = int32(10)
	goto L143
L143:
	;
	v616 = F_SysCacheGetAttrNotNull(m, int32(16), v563, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L16
	} else {
		goto L144
	}
L144:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+76)))
	v619 = F_text_to_cstring(m, v616)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L16
	} else {
		goto L145
	}
L145:
	;
	switch v618 - int32(98) {
	case 0:
		goto L148
	case 1:
		goto L147
	default:
		goto L23
	}
L146:
	;
	if v637 == int32(0) {
		goto L23
	} else {
		goto L154
	}
L147:
	;
	v626 = F_pg_strcasecmp(m, int32(545070), v619)
	mBase = m.M
	if v626 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v623 = F_get_collation_actual_version_builtin(m, v619)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L16
	} else {
		goto L149
	}
L149:
	;
	v637 = v623
	goto L146
L150:
	;
	v637 = int32(0)
	goto L146
L151:
	;
	goto L150
L152:
	;
	v631 = F_pg_strncasecmp(m, int32(660007), v619, int32(2))
	mBase = m.M
	if v631 == int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v635 = F_pg_strcasecmp(m, int32(510385), v619)
	mBase = m.M
	goto L151
L154:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v643 == int32(0) {
		v662 = v642
		v663 = v643
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v663-v662 == int32(0) {
		goto L138
	} else {
		goto L163
	}
L156:
	;
	goto L155
L157:
	;
	if v642 != v643 {
		v662 = v642
		v663 = v643
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v647 = v637
	v648 = v607
	goto L159
L159:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648)+1)))
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647)+1)))
	if v652 == int32(0) {
		v662 = v651
		v663 = v652
		goto L156
	} else {
		goto L161
	}
L160:
	;
	v662 = v651
	v663 = v652
	goto L156
L161:
	;
	v655 = int32(1)
	if v651 == v652 {
		v647 = v647 + v655
		v648 = v648 + v655
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v669 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L16
	} else {
		goto L164
	}
L164:
	;
	if v669 == int32(0) {
		goto L138
	} else {
		goto L165
	}
L165:
	;
	v674 = v569 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v674
	F_errmsg(m, int32(324934), v18+int32(80))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L16
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v607
	F_errdetail(m, int32(605310), v18-int32(-64))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L16
	} else {
		goto L167
	}
L167:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v569)+68))
	v689 = F_get_namespace_name(m, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L16
	} else {
		goto L168
	}
L168:
	;
	v691 = F_quote_qualified_identifier(m, v689, v674)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L16
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v691
	F_errhint(m, int32(618743), v18+int32(48))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L16
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(499565), int32(1142), int32(398107))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L16
	} else {
		goto L171
	}
L171:
	;
	goto L138
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v597
	v711 = v597
	goto L25
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(45963), v18)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(499565), int32(1085), int32(398107))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L16
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v569 + int32(4)
	F_errmsg(m, int32(460605), v18+int32(32))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L16
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(499565), int32(1128), int32(398107))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L16
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = int32(0)
	F_errmsg_internal(m, int32(45963), v18+int32(96))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L16
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(499565), int32(1212), int32(262718))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L16
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_notification_queue_usage(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		v14 = F_LWLockAcquire(m, v10+int32(3456), int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[458]))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != v19 {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[459]))
				v27 = base.F64_div(base.F64_convert_i64_s(v18-v19), base.F64_convert_i32_s(v24))
			} else {
				v27 = float64(0)
			}
			v29 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			F_LWLockRelease(m, v29+int32(3456))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_Float8GetDatum(m, v27)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_pg_numa_available(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_pg_opfamily_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OpfamilyIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_options_to_table(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_untransformRelOptions(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v10 + int32(16)
	return int32(0)
L5:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v24 <= v23 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = v23
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v27<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = F_cstring_to_text(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+6)) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v41
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = F_cstring_to_text(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v52 = int32(1)
	v53 = int32(0)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)) = uint8(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_tuplestore_putvalues(m, v56, v57, v10+int32(8), v10+int32(6))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v52 = int32(0)
	v53 = v50
	goto L12
L14:
	;
	v65 = v27 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v65 < v66 {
		v27 = v65
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
}
func F_pg_partition_ancestors(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = int32(0)
			v19 = F_SearchSysCacheExists(m, int32(57), v10, v16, v16, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = int32(2)
						v85 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
						return int32(0)
					}
				} else {
					v23 = F_get_rel_relkind(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = F_get_rel_relispartition(m, v10)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							if v25 != 0 {
								v33 = int32(4515712)
								v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
								v38 = F_get_partition_ancestors(m, v10)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = F_lcons_oid(m, v10, v38)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v40
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
										if v50 == int32(0) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
												v75 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
												return int32(0)
											}
										} else {
											v53 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
											v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+4)))
											if base.Ui64(v54) <= base.Ui64(v53) {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
													v75 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
													return int32(0)
												}
											} else {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
												v61 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v53)<<(uint(int32(2))%32))))
												*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53 + int64(1)
												v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(1)
												return v61
											}
										}
									}
								}
							} else {
								if v23 == int32(112) {
									v33 = int32(4515712)
									v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
									v38 = F_get_partition_ancestors(m, v10)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v40 = F_lcons_oid(m, v10, v38)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v40
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
											if v50 == int32(0) {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
													v75 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
													return int32(0)
												}
											} else {
												v53 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
												v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+4)))
												if base.Ui64(v54) <= base.Ui64(v53) {
													F_end_MultiFuncCall(m, l0)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int32(0)
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
														v75 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
														return int32(0)
													}
												} else {
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v53)<<(uint(int32(2))%32))))
													*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53 + int64(1)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(1)
													return v61
												}
											}
										}
									}
								} else {
									if v23&int32(255) != int32(73) {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = int32(2)
											v85 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
											return int32(0)
										}
									} else {
										v33 = int32(4515712)
										v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
										v38 = F_get_partition_ancestors(m, v10)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											v40 = F_lcons_oid(m, v10, v38)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v40
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
												v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
												if v50 == int32(0) {
													F_end_MultiFuncCall(m, l0)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int32(0)
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
														v75 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
														return int32(0)
													}
												} else {
													v53 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
													v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+4)))
													if base.Ui64(v54) <= base.Ui64(v53) {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
															v75 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
															return int32(0)
														}
													} else {
														v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
														v61 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v53)<<(uint(int32(2))%32))))
														*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53 + int64(1)
														v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(1)
														return v61
													}
												}
											}
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
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
		if v50 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
				v75 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
				return int32(0)
			}
		} else {
			v53 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
			v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+4)))
			if base.Ui64(v54) <= base.Ui64(v53) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = int32(2)
					v75 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
					return int32(0)
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v53)<<(uint(int32(2))%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53 + int64(1)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(1)
				return v61
			}
		}
	}
}
func F_pg_partition_tree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v96 int32
	_ = v96
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v198
L2:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L51
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L48
	}
L4:
	;
	v20 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	goto L21
L7:
	;
	return int32(0)
L8:
	;
	v25 = int32(0)
	v28 = F_SearchSysCacheExists(m, int32(57), v15, v25, v25, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v28 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v32 = F_get_rel_relkind(m, v15)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = F_get_rel_relispartition(m, v15)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v42 = int32(4515712)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
	v49 = F_find_all_inheritors(m, v15, int32(1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L17
	}
L13:
	;
	if v34 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v32 == int32(112) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v32&int32(255) != int32(73) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v54 = F_get_call_result_type(m, l0, int32(0), v13+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	if v54 != int32(1) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v58
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v43
	goto L6
L20:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L47
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v69 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+4)))
	if base.Ui64(v73) <= base.Ui64(v72) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v75
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82+base.I32_wrap_i64(v72)<<(uint(int32(2))%32))))
	v88 = F_get_rel_relkind(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v90 = F_get_partition_ancestors(m, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v87
	if v90 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v145
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v155 = F_heap_form_tuple(m, v150, v13+int32(16), v13+int32(12))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L7
	} else {
		goto L45
	}
L27:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v114 < v115 {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = base.B2i32(v88 != int32(112)) & base.B2i32(v88 != int32(73))
	if v87 == v15 {
		v145 = v79
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = base.B2i32(v88 != int32(112)) & base.B2i32(v88 != int32(73))
	if v87 == v15 {
		v145 = v79
		goto L26
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	if v90 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v145 = v79
	goto L26
L35:
	;
	goto L27
L36:
	;
	v119 = v115
	goto L38
L37:
	;
	v119 = v114
	goto L38
L38:
	;
	v121 = v114
	goto L39
L39:
	;
	if v121 == v119 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v145 = v134
	goto L26
L41:
	;
	v145 = v119
	goto L26
L42:
	;
	goto L43
L43:
	;
	v134 = v121 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32)+v135)))
	if v137 != v15 {
		v121 = v134
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	v158 = F_HeapTupleHeaderGetDatum(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v160 + int64(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+20)) = int32(1)
	v198 = v158
	goto L1
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = int32(2)
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v173)
	v198 = int32(0)
	goto L1
L48:
	;
	F_errmsg_internal(m, int32(367739), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(495090), int32(91), int32(410554))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = int32(2)
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
	v198 = v2
	goto L1
}
func F_pg_postmaster_start_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _consts[1050]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_prepared_statement(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
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
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
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
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	if v26 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(80)
	return int32(0)
L4:
	;
	F_hash_seq_init(m, v17+int32(60), v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = F_hash_seq_search(m, v17+int32(60))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v35 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v46 = v35
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
	v57 = F_cstring_to_text(m, v46)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v62 = F_cstring_to_text(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v62
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v46)+72))
	v66 = F_Int64GetDatum(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	v74 = F_palloc(m, v71<<(uint(int32(2))%32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v71 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v186 = F_construct_array_builtin(m, v74, v71, int32(2206))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	v79 = v71 & int32(3)
	v80 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v71) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v87 = v80
	v97 = int32(0)
	goto L19
L17:
	;
	v130 = v80
	goto L18
L18:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L19:
	;
	v102 = v87 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v102))) = v105
	v107 = int32(4)
	v108 = v102 | v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70+v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v108))) = v111
	v114 = v102 | int32(8)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v70+v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v114))) = v117
	v120 = v102 | int32(12)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v120))) = v123
	v126 = v87 + v107
	v128 = v97 + v107
	if v128 != v71&int32(2147483644) {
		v87 = v126
		v97 = v128
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v130 = v126
	goto L18
L21:
	;
	goto L20
L22:
	;
	v146 = v130
	v151 = v80
	goto L23
L23:
	;
	v161 = v146 << (uint(int32(2)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v161))) = v164
	v166 = int32(1)
	v169 = v151 + v166
	if v169 != v79 {
		v146 = v146 + v166
		v151 = v169
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v186
	if v56 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+68)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v368 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v368 + int32(136)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	F_tuplestore_putvalues(m, v375, v376, v17+int32(16), v17+int32(8))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L52
	}
L28:
	;
	v349 = F_construct_array_builtin(m, v335, v337, int32(2206))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L51
	}
L29:
	;
	v332 = F_palloc(m, v195<<(uint(int32(2))%32))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L50
	}
L30:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v192 = F_palloc(m, v189<<(uint(int32(2))%32))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v328)
	goto L27
L33:
	;
	v194 = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v195 <= v194 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v200 = v194
	v203 = v195
	goto L35
L35:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(88)+v203<<(uint(int32(4))%32)+v200*int32(100))))
	*(*int32)(unsafe.Add(mBase, uint32(v192+v200<<(uint(int32(2))%32)))) = v223
	v226 = v200 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v226 < v227 {
		v200 = v226
		v203 = v227
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v231 = F_palloc(m, v227<<(uint(int32(2))%32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	if v227 <= int32(0) {
		v335 = v231
		v337 = v227
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v235 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v227) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v242 = v235
	v250 = int32(0)
	goto L43
L41:
	;
	v285 = v235
	goto L42
L42:
	;
	v300 = v227 & int32(3)
	if v300 == int32(0) {
		v335 = v231
		v337 = v227
		goto L28
	} else {
		goto L46
	}
L43:
	;
	v257 = v242 << (uint(int32(2)) % 32)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v192+v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v257))) = v260
	v262 = int32(4)
	v263 = v257 | v262
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v192+v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v263))) = v266
	v269 = v257 | int32(8)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v192+v269)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v269))) = v272
	v275 = v257 | int32(12)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v192+v275)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v275))) = v278
	v281 = v242 + v262
	v283 = v250 + v262
	if v283 != v227&int32(2147483644) {
		v242 = v281
		v250 = v283
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v285 = v281
	goto L42
L45:
	;
	goto L44
L46:
	;
	v303 = v285
	v313 = v235
	goto L47
L47:
	;
	v318 = v303 << (uint(int32(2)) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v192+v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v318))) = v321
	v323 = int32(1)
	v326 = v313 + v323
	if v326 != v300 {
		v303 = v303 + v323
		v313 = v326
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v335 = v231
	v337 = v227
	goto L28
L49:
	;
	goto L48
L50:
	;
	v335 = v332
	v337 = v195
	goto L28
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v349
	goto L27
L52:
	;
	v385 = F_hash_seq_search(m, v17+int32(60))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v385 != 0 {
		v46 = v385
		goto L8
	} else {
		goto L54
	}
L54:
	;
	goto L9
}
func F_pg_printf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, _consts[718]))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
		m.G0 = v7 + int32(1072)
		return
	} else {
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1064)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1060)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1056)) = v7 + int32(1040)
		v26 = v7 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1052)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1048)) = v26
		F_dopr(m, v7+int32(1048), l0, l1)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)))
			if v35 == int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1048))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1052))
				if v38 != v39 {
					v45 = v38 - v39
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1060))
					v47 = F_fwrite(m, v39, int32(1), v45, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 != v45 {
						} else {
						}
						m.G0 = v7 + int32(1072)
						return
					}
				} else {
					if v35 != 0 {
					} else {
					}
					m.G0 = v7 + int32(1072)
					return
				}
			} else {
				if v35 != 0 {
				} else {
				}
				m.G0 = v7 + int32(1072)
				return
			}
		}
	}
}
func F_pg_qsort_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v7 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l1, l2)
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v7 < int32(0) {
				if v11 < int32(0) {
					v30 = l1
					return v30
				} else {
					v17 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						if v17 < int32(0) {
							v21 = l2
						} else {
							v21 = l0
						}
						return v21
					}
				}
			} else {
				if int32(0) < v11 {
					v30 = l1
					return v30
				} else {
					v25 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v29 = l0
						} else {
							v29 = l2
						}
						v30 = v29
						return v30
					}
				}
			}
		}
	}
}
func F_pg_random_bytes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(-1025)) < base.Ui32(v4-int32(1025)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_px_THROW_ERROR(m, int32(-17))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L25
	}
L2:
	;
	v10 = v4 + int32(4)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L21
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
	v20 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v20
	v32 = F_open(m, int32(288751), v20, v26)
	mBase = m.M
	if v32 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v65 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v35 = int32(1)
	if v4 == int32(0) {
		v58 = v35
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v65 = v20
	goto L10
L10:
	;
	m.G0 = v26 + int32(16)
	goto L7
L11:
	;
	v60 = F_close(m, v32)
	mBase = m.M
	v65 = v58
	goto L10
L12:
	;
	v38 = v11 + int32(4)
	v39 = v4
	goto L13
L13:
	;
	v44 = F_read(m, v32, v38, v39)
	mBase = m.M
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = v35
	goto L11
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v48 == int32(27) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v53 = v39 - v44
	if v53 != 0 {
		v38 = v38 + v44
		v39 = v53
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(0)
	goto L11
L19:
	;
	goto L14
L20:
	;
	return v11
L21:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(402310), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496479), int32(465), int32(158861))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_random_uuid(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_gen_random_uuid(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_read_binary_file_off_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if int64(0) <= v11 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			v16 = F_convert_and_check_filename(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = F_read_binary_file(m, v16, v15, v11, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v23 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
						return int32(0)
					} else {
						return v19
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(343479), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499423), int32(269), int32(246116))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
func F_pg_server_to_client(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	if l1 <= int32(0) {
		v35 = l0
		return v35
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		if v9 == int32(0) {
			v35 = l0
			return v35
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[460]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v9 == v14 {
				v35 = l0
				return v35
			} else {
				if v14 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_consts[1127])))
					v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if l1 == v23 {
							v35 = l0
							return v35
						} else {
							F_report_invalid_encoding(m, v9, l0+v23, l1-v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = F_perform_default_encoding_conversion(m, l0, l1, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v33
						return v35
					}
				}
			}
		}
	}
}
func F_pg_signal_backend(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_BackendPidGetProc(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v18 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v94 = int32(1)
					m.G0 = v8 + int32(16)
					return v94
				} else {
					v82 = int32(130048)
					v83 = int32(74)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, v82, v8)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495233), v83, int32(427360))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v94 = int32(1)
							m.G0 = v8 + int32(16)
							return v94
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
			if v24 != 0 {
				v25 = F_superuser_arg(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
						v57 = F_has_privs_of_role(m, v55, v56)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v66 = int32(0)
								v69 = F_kill(m, v66-l0, l1)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									if v69 == int32(0) {
										v94 = v66
										m.G0 = v8 + int32(16)
										return v94
									} else {
										v75 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											if v75 == int32(0) {
												v94 = int32(1)
												m.G0 = v8 + int32(16)
												return v94
											} else {
												v82 = int32(296102)
												v83 = int32(123)
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
												F_errmsg(m, v82, v8)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(495233), v83, int32(427360))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														v94 = int32(1)
														m.G0 = v8 + int32(16)
														return v94
													}
												}
											}
										}
									}
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								v62 = F_has_privs_of_role(m, v60, int32(4200))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 != 0 {
										v66 = int32(0)
										v69 = F_kill(m, v66-l0, l1)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											if v69 == int32(0) {
												v94 = v66
												m.G0 = v8 + int32(16)
												return v94
											} else {
												v75 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													if v75 == int32(0) {
														v94 = int32(1)
														m.G0 = v8 + int32(16)
														return v94
													} else {
														v82 = int32(296102)
														v83 = int32(123)
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
														F_errmsg(m, v82, v8)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(495233), v83, int32(427360))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																v94 = int32(1)
																m.G0 = v8 + int32(16)
																return v94
															}
														}
													}
												}
											}
										}
									} else {
										v94 = int32(2)
										m.G0 = v8 + int32(16)
										return v94
									}
								}
							}
						}
					} else {
						v29 = int32(4)
						v31 = *(*int32)(unsafe.Add(mBase, _consts[845]))
						v33 = *(*int32)(unsafe.Add(mBase, _consts[150]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v37 = base.I32_div_s(v10-v34, int32(640))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v37*int32(408))+8))
						if v41 == v29 {
							v45 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v47 = F_has_privs_of_role(m, v45, int32(6392))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									v94 = v29
									m.G0 = v8 + int32(16)
									return v94
								} else {
									v66 = int32(0)
									v69 = F_kill(m, v66-l0, l1)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										if v69 == int32(0) {
											v94 = v66
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v75 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												if v75 == int32(0) {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												} else {
													v82 = int32(296102)
													v83 = int32(123)
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
													F_errmsg(m, v82, v8)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(495233), v83, int32(427360))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															v94 = int32(1)
															m.G0 = v8 + int32(16)
															return v94
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v51 = F_superuser(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 != 0 {
									v66 = int32(0)
									v69 = F_kill(m, v66-l0, l1)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										if v69 == int32(0) {
											v94 = v66
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v75 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												if v75 == int32(0) {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												} else {
													v82 = int32(296102)
													v83 = int32(123)
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
													F_errmsg(m, v82, v8)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(495233), v83, int32(427360))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															v94 = int32(1)
															m.G0 = v8 + int32(16)
															return v94
														}
													}
												}
											}
										}
									}
								} else {
									v94 = int32(3)
									m.G0 = v8 + int32(16)
									return v94
								}
							}
						}
					}
				}
			} else {
				v29 = int32(4)
				v31 = *(*int32)(unsafe.Add(mBase, _consts[845]))
				v33 = *(*int32)(unsafe.Add(mBase, _consts[150]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v37 = base.I32_div_s(v10-v34, int32(640))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v37*int32(408))+8))
				if v41 == v29 {
					v45 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v47 = F_has_privs_of_role(m, v45, int32(6392))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 == int32(0) {
							v94 = v29
							m.G0 = v8 + int32(16)
							return v94
						} else {
							v66 = int32(0)
							v69 = F_kill(m, v66-l0, l1)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								if v69 == int32(0) {
									v94 = v66
									m.G0 = v8 + int32(16)
									return v94
								} else {
									v75 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										if v75 == int32(0) {
											v94 = int32(1)
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v82 = int32(296102)
											v83 = int32(123)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
											F_errmsg(m, v82, v8)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495233), v83, int32(427360))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v51 = F_superuser(m)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v66 = int32(0)
							v69 = F_kill(m, v66-l0, l1)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								if v69 == int32(0) {
									v94 = v66
									m.G0 = v8 + int32(16)
									return v94
								} else {
									v75 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										if v75 == int32(0) {
											v94 = int32(1)
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v82 = int32(296102)
											v83 = int32(123)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
											F_errmsg(m, v82, v8)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495233), v83, int32(427360))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												}
											}
										}
									}
								}
							}
						} else {
							v94 = int32(3)
							m.G0 = v8 + int32(16)
							return v94
						}
					}
				}
			}
		}
	}
}
func F_pg_snapshot_out(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_initStringInfo(m, v7-int32(-64))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v18
	F_appendStringInfo(m, v7-int32(-64), int32(547471), v7+int32(48))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v27
	F_appendStringInfo(m, v7-int32(-64), int32(547471), v7+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+64))
	m.G0 = v7 + int32(80)
	return v81
L7:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v39
	F_appendStringInfo(m, v7-int32(-64), int32(38340), v7+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.Ui32(v48) < base.Ui32(int32(2)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v54 = int32(1)
	goto L10
L10:
	;
	F_appendStringInfoChar(m, v7-int32(-64), int32(44))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)+v54<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v66
	F_appendStringInfo(m, v7-int32(-64), int32(38340), v7)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v74 = v54 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.Ui32(v74) < base.Ui32(v75) {
		v54 = v74
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
}
func F_pg_snapshot_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pq_getmsgint(m, v10, int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L26
	}
L2:
	;
	return int32(0)
L3:
	;
	if base.Ui32(int32(134217724)) < base.Ui32(v12) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if base.I32_wrap_i64(v18) == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v20&int64(4294967295) == int64(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if base.Ui64(v20) < base.Ui64(v18) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = F_palloc(m, v12<<(uint(int32(3))%32)+int32(24))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v18
	if v12 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v72<<(uint(int32(5))%32) + int32(96)
	return v34
L12:
	;
	v72 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = int32(0)
	v45 = v12
	v52 = int64(0)
	goto L15
L15:
	;
	v53 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v72 = v68
	goto L11
L17:
	;
	if base.Ui64(v53) < base.Ui64(v52) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if base.Ui64(v53) < base.Ui64(v18) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if base.Ui64(v20) < base.Ui64(v53) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v53 == v52 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v67 < v68 {
		v44 = v67
		v45 = v68
		v52 = v69
		goto L15
	} else {
		goto L25
	}
L22:
	;
	v67 = v44
	v68 = v45 - int32(1)
	v69 = v52
	goto L21
L23:
	;
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v34+int32(24)+v44<<(uint(int32(3))%32)))) = v53
	v67 = v44 + int32(1)
	v68 = v45
	v69 = v53
	goto L21
L25:
	;
	goto L16
L26:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(504758), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(495370), int32(522), int32(36310))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_sockaddr_cidr_mask(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v142 int32
	_ = v142
	v3 = l2
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l1 == int32(0) {
		if v3 == int32(2) {
			v18 = int32(32)
		} else {
			v18 = int32(128)
		}
		v31 = v18
		v33 = int32(-1)
		switch v3 - int32(2) {
		case 0:
			if base.Ui32(int32(32)) < base.Ui32(v31) {
				v142 = v33
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
				v40 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
				v45 = int32(-1) << (uint(int32(32)-v31) % 32)
				v46 = int32(24)
				v48 = int32(65280)
				v50 = int32(8)
				if v31 != 0 {
					v62 = v45<<(uint(v46)%32) | v45&v48<<(uint(v50)%32) | (int32(base.Ui32(v45)>>(uint(v50)%32))&v48 | int32(base.Ui32(v45)>>(uint(v46)%32)))
				} else {
					v62 = v40
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
				v142 = int32(0)
			}
		default:
			v142 = v33
		case 8:
			if base.Ui32(int32(128)) < base.Ui32(v31) {
				v142 = v33
			} else {
				v66 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
				v69 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v69
				v72 = v10 + int32(8)
				*(*int64)(unsafe.Add(mBase, uint32(v72))) = v69
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v69
				v78 = v66
				v80 = v31
				for {
					v85 = int32(0)
					if v80 <= v85 {
						v95 = v85
					} else {
						if base.Ui32(int32(7)) < base.Ui32(v80) {
							v95 = int32(255)
						} else {
							v95 = int32(255) << (uint(int32(8)-v80) % 32)
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v78+v72))) = uint8(v95)
					v100 = int32(0)
					v102 = v80 - int32(8)
					if v102 <= v100 {
						v112 = v100
					} else {
						if base.Ui32(int32(7)) < base.Ui32(v102) {
							v112 = int32(255)
						} else {
							v112 = int32(255) << (uint(int32(16)-v80) % 32)
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v72+(v78|int32(1))))) = uint8(v112)
					v114 = int32(16)
					v117 = v78 + int32(2)
					if v117 != v114 {
						v78 = v117
						v80 = v80 - v114
						continue
					} else {
						break
					}
					break
				}
				v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v120
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
				v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v124
				v126 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v126
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
				v142 = int32(0)
			}
		}
	} else {
		v23 = F_strtox_2(m, l1, v10+int32(28), int32(10), int64(2147483648))
		mBase = m.M
		v25 = int32(-1)
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v26 == int32(0) {
			v142 = v25
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			if v30 != 0 {
				v142 = v25
			} else {
				v31 = base.I32_wrap_i64(v23)
				v33 = int32(-1)
				switch v3 - int32(2) {
				case 0:
					if base.Ui32(int32(32)) < base.Ui32(v31) {
						v142 = v33
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
						v40 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
						v45 = int32(-1) << (uint(int32(32)-v31) % 32)
						v46 = int32(24)
						v48 = int32(65280)
						v50 = int32(8)
						if v31 != 0 {
							v62 = v45<<(uint(v46)%32) | v45&v48<<(uint(v50)%32) | (int32(base.Ui32(v45)>>(uint(v50)%32))&v48 | int32(base.Ui32(v45)>>(uint(v46)%32)))
						} else {
							v62 = v40
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
						*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
						v142 = int32(0)
					}
				default:
					v142 = v33
				case 8:
					if base.Ui32(int32(128)) < base.Ui32(v31) {
						v142 = v33
					} else {
						v66 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
						v69 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v69
						v72 = v10 + int32(8)
						*(*int64)(unsafe.Add(mBase, uint32(v72))) = v69
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v69
						v78 = v66
						v80 = v31
						for {
							v85 = int32(0)
							if v80 <= v85 {
								v95 = v85
							} else {
								if base.Ui32(int32(7)) < base.Ui32(v80) {
									v95 = int32(255)
								} else {
									v95 = int32(255) << (uint(int32(8)-v80) % 32)
								}
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v78+v72))) = uint8(v95)
							v100 = int32(0)
							v102 = v80 - int32(8)
							if v102 <= v100 {
								v112 = v100
							} else {
								if base.Ui32(int32(7)) < base.Ui32(v102) {
									v112 = int32(255)
								} else {
									v112 = int32(255) << (uint(int32(16)-v80) % 32)
								}
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v72+(v78|int32(1))))) = uint8(v112)
							v114 = int32(16)
							v117 = v78 + int32(2)
							if v117 != v114 {
								v78 = v117
								v80 = v80 - v114
								continue
							} else {
								break
							}
							break
						}
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v120
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v124
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v126
						*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
						v142 = int32(0)
					}
				}
			}
		}
	}
	m.G0 = v10 + int32(32)
	return v142
}
func F_pg_split_walfile_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v149 int64
	_ = v149
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int64
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v10 = m.G0
	v12 = v10 - int32(336)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_text_to_cstring(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+322)) = uint16(v21)
	v23 = F_pstrdup(m, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = v23
	v31 = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v59 = F_strlen(m, v23)
	mBase = m.M
	if v59 != int32(24) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	if base.Ui32((v31-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v45)
	v48 = v28 + int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 != 0 {
		v28 = v48
		v31 = v49
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v43 = v31 - int32(32)
	goto L13
L12:
	;
	v43 = v31
	goto L13
L13:
	;
	v45 = v43 & int32(255)
	goto L10
L14:
	;
	goto L9
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L52
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L48
	}
L17:
	;
	v62 = int32(537590)
	v66 = m.G0
	v68 = v66 - int32(32)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68)+24)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[271])))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v145 != int32(24) {
		goto L16
	} else {
		goto L39
	}
L19:
	;
	v145 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _consts[272])))
	if v81 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = v23
	goto L25
L23:
	;
	goto L24
L24:
	;
	v95 = v62
	v96 = v77
	goto L28
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v91 == v77 {
		v85 = v85 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v145 = v85 - v23
	goto L18
L27:
	;
	goto L26
L28:
	;
	v103 = v68 + int32(base.Ui32(v96)>>(uint(int32(3))%32))&int32(28)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 | v105<<(uint(v96)%32)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v109 != 0 {
		v95 = v95 + v105
		v96 = v109
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v112 == int32(0) {
		v137 = v23
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v145 = v137 - v23
	goto L18
L32:
	;
	v116 = v23
	v117 = v112
	goto L33
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v125)>>(uint(v117)%32))&int32(1) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v137 = v133
	goto L31
L35:
	;
	v137 = v116
	goto L31
L36:
	;
	goto L37
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v133 = v116 + int32(1)
	if v131 != 0 {
		v116 = v133
		v117 = v131
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v149 = int64(*(*int32)(unsafe.Add(mBase, _consts[189])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(332)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v12 + int32(324)
	v162 = F_sscanf(m, v23, int32(510531), v12+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+324)))
	v165 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)))
	v169 = F_get_call_result_type(m, l0, int32(0), v12+int32(316))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v169 != int32(1) {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	v174 = base.I64_div_u_s(int64(4294967296), v149)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v174*v165 + v164
	v182 = F_pg_snprintf(m, v12+int32(48), int32(256), int32(38340), v12)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v185 = int32(0)
	v190 = F_DirectFunctionCall3Coll(m, int32(408), v185, v12+int32(48), v185, int32(-1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+324)) = v190
	v193 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+332)))
	v194 = F_Int64GetDatum(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+328)) = v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+316))
	v202 = F_heap_form_tuple(m, v197, v12+int32(324), v12+int32(322))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
	v205 = F_HeapTupleHeaderGetDatum(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v12 + int32(336)
	return v205
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v19
	F_errmsg(m, int32(715268), v12+int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(495280), int32(487), int32(380346))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(367739), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(495280), int32(492), int32(380346))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_statistics_obj_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_StatisticsObjIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_stats_ext_mcvlist_items(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 float64
	_ = v168
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L45
	}
L2:
	;
	v20 = F_init_MultiFuncCall(m, l0)
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	goto L17
L5:
	;
	return int32(0)
L6:
	;
	v24 = int32(4515712)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = F_statext_mcv_deserialize(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v32
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v37
	goto L11
L10:
	;
	goto L11
L11:
	;
	v40 = F_get_call_result_type(m, l0, int32(0), v14)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v40 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v45 = F_BlessTupleDesc(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v45
	v48 = F_TupleDescGetAttInMetadata(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v48
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	goto L4
L16:
	;
	m.G0 = v14 + int32(48)
	return v202
L17:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui64(v58) < base.Ui64(v59) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v62 = base.I32_wrap_i64(v58)
	v67 = v61 + v62*int32(24) + int32(48)
	v68 = int32(0)
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+12)))
	if v68 < v69 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L44
	}
L21:
	;
	v77 = int32(0)
	v80 = v68
	v81 = v2
	goto L24
L22:
	;
	v146 = v68
	v147 = v2
	v152 = v62
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v152
	v155 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v156 = F_makeArrayResult(m, v147, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L38
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v77))))
	v92 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v93 = F_accumArrayResult(m, v80, v88, int32(0), int32(16), v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v146 = v93
	v147 = v135
	v152 = v140
	goto L23
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v77))))
	if v97 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v137 = v77 + int32(1)
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+12)))
	if v137 < v138 {
		v77 = v137
		v80 = v93
		v81 = v135
		goto L24
	} else {
		goto L37
	}
L28:
	;
	v101 = v77 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(16)+v101)))
	F_getTypeOutputInfo(m, v103, v14+int32(40), v14+int32(39))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v132 = F_accumArrayResult(m, v81, int32(0), int32(1), int32(25), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L36
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_fmgr_info(m, v110, v14)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114+v101)))
	v117 = F_FunctionCall1Coll(m, v14, int32(0), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v119 = F_cstring_to_text(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v125 = F_accumArrayResult(m, v81, v119, int32(0), int32(25), v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v135 = v125
	goto L27
L36:
	;
	v135 = v132
	goto L27
L37:
	;
	goto L25
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v156
	v160 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v161 = F_makeArrayResult(m, v146, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v161
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	v165 = F_Float8GetDatum(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v165
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	v169 = F_Float8GetDatum(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v169
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v180 = F_heap_form_tuple(m, v177, v14, v14+int32(40))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v183 = F_HeapTupleHeaderGetDatum(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v185 + int64(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+20)) = int32(1)
	v202 = v183
	goto L16
L44:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+20)) = int32(2)
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v197)
	v202 = int32(0)
	goto L16
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(421636), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(492918), int32(1369), int32(150799))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_strtitle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
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
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v14 - int32(98) {
	case 0:
		goto L2
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v672
L2:
	;
	v640 = m.G0
	v642 = v640 - int32(16)
	m.G0 = v642
	v644 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v642)+8)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v642)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v642))) = l2
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v642)+15)) = uint8(v644)
	*(*uint16)(unsafe.Add(mBase, uint32(v642)+13)) = uint16(v644)
	v653 = int32(1)
	v654 = v648 ^ v653
	*(*uint8)(unsafe.Add(mBase, uint32(v642)+12)) = uint8(v654)
	v658 = F_convert_case(m, l0, l1, l2, l3, v653, v648, int32(1478), v642)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L15
	} else {
		goto L203
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L15
	} else {
		goto L200
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_consts[979])))
	goto L7
L5:
	;
	v672 = v623
	goto L1
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L15
	} else {
		goto L196
	}
L7:
	;
	if int32(2) <= v24 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l3 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L11:
	;
	v30 = F_strlen(m, l2)
	mBase = m.M
	v31 = v30
	goto L13
L12:
	;
	v31 = l3
	goto L13
L13:
	;
	v33 = v31 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v33) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v38 = F_palloc(m, v33<<(uint(int32(2))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_char2wchar(m, v38, v33, l2, v31, l4)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v45
	v50 = int32(0)
	v52 = v44
	goto L21
L19:
	;
	v86 = v44
	goto L20
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92*int32(28))+uint32(_consts[979])))
	goto L34
L21:
	;
	if v50 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v86 = v76
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+v52<<(uint(int32(2))%32)))) = v63
	if base.Ui32(int32(10)) <= base.Ui32(v63-int32(48)) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v60 = F_casemap(m, v49, int32(0))
	mBase = m.M
	goto L27
L25:
	;
	goto L26
L26:
	;
	v62 = F_casemap(m, v49, int32(1))
	mBase = m.M
	goto L28
L27:
	;
	v63 = v60
	goto L23
L28:
	;
	v63 = v62
	goto L23
L29:
	;
	v76 = v52 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v38+v76<<(uint(int32(2))%32))))
	if v80 != 0 {
		v49 = v80
		v50 = v74
		v52 = v76
		goto L21
	} else {
		goto L33
	}
L30:
	;
	v71 = F_iswalpha(m, v63)
	mBase = m.M
	v74 = base.B2i32(v71 != int32(0))
	goto L32
L31:
	;
	v74 = int32(1)
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L22
L34:
	;
	v100 = v97*v86 + int32(1)
	v101 = F_palloc(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v100 == int32(0) {
		v489 = v6
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(v489+int32(1)) <= base.Ui32(l1) {
		goto L147
	} else {
		goto L148
	}
L37:
	;
	if l4 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v107 = int32(0)
	v113 = m.G0
	v114 = int32(16)
	v115 = v113 - v114
	m.G0 = v115
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v38
	v119 = v115 + int32(12)
	v120 = m.G0
	v122 = v120 - v114
	m.G0 = v122
	if v101 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	goto L40
L40:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v287 = *(*int32)(unsafe.Add(mBase, _consts[1012]))
	if v284 != 0 {
		goto L85
	} else {
		goto L86
	}
L41:
	;
	v489 = v274
	goto L36
L42:
	;
	v278 = int32(16)
	m.G0 = v122 + v278
	m.G0 = v115 + v278
	goto L41
L43:
	;
	v274 = v100 - v259
	goto L42
L44:
	;
	if v196 != 0 {
		goto L69
	} else {
		goto L70
	}
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v155 = v100
	v156 = v101
	v160 = v154
	goto L58
L46:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v100) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 == int32(0) {
		v274 = v107
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v196 = v100
	v197 = v101
	goto L44
L50:
	;
	v130 = v127
	v131 = v126
	v133 = v107
	goto L51
L51:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v130) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v274 = v153
	goto L42
L53:
	;
	v142 = int32(-1)
	v145 = F_wcrtomb(m, v122+int32(12), v130)
	mBase = m.M
	if v145 == v142 {
		v274 = v142
		goto L42
	} else {
		goto L56
	}
L54:
	;
	v148 = int32(1)
	goto L55
L55:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v153 = v133 + v148
	if v150 != 0 {
		v130 = v150
		v131 = v131 + int32(4)
		v133 = v153
		goto L51
	} else {
		goto L57
	}
L56:
	;
	v148 = v145
	goto L55
L57:
	;
	goto L52
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if base.Ui32(v164-int32(128)) <= base.Ui32(int32(-128)) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v196 = v186
	v197 = v189
	goto L44
L60:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v192 = v190 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v192
	if base.Ui32(int32(3)) < base.Ui32(v186) {
		v155 = v186
		v156 = v189
		v160 = v192
		goto L58
	} else {
		goto L68
	}
L61:
	;
	if v164 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v164)
	v182 = int32(1)
	v186 = v155 - v182
	v189 = v156 + v182
	goto L60
L64:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v171
	v259 = v155
	goto L43
L65:
	;
	goto L66
L66:
	;
	v175 = int32(-1)
	v176 = F_wcrtomb(m, v156, v164)
	mBase = m.M
	if v176 == v175 {
		v274 = v175
		goto L42
	} else {
		goto L67
	}
L67:
	;
	v186 = v155 - v176
	v189 = v156 + v176
	goto L60
L68:
	;
	goto L59
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v206 = v196
	v207 = v197
	v209 = v205
	goto L72
L70:
	;
	goto L71
L71:
	;
	v274 = v100
	goto L42
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if base.Ui32(v215-int32(128)) <= base.Ui32(int32(-128)) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L71
L74:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v248 = v246 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v248
	if v242 != 0 {
		v206 = v242
		v207 = v245
		v209 = v248
		goto L72
	} else {
		goto L83
	}
L75:
	;
	if v215 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v215)
	v238 = int32(1)
	v242 = v206 - v238
	v245 = v207 + v238
	goto L74
L78:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v222
	v259 = v206
	goto L43
L79:
	;
	goto L80
L80:
	;
	v226 = int32(-1)
	v229 = F_wcrtomb(m, v122+int32(12), v215)
	mBase = m.M
	if v229 == v226 {
		v274 = v226
		goto L42
	} else {
		goto L81
	}
L81:
	;
	if base.Ui32(v206) < base.Ui32(v229) {
		v259 = v206
		goto L43
	} else {
		goto L82
	}
L82:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v234 = F_wcrtomb(m, v207, v233)
	mBase = m.M
	v242 = v206 - v229
	v245 = v207 + v229
	goto L74
L83:
	;
	goto L73
L84:
	;
	v298 = int32(0)
	v304 = m.G0
	v305 = int32(16)
	v306 = v304 - v305
	m.G0 = v306
	*(*int32)(unsafe.Add(mBase, uint32(v306)+12)) = v38
	v310 = v306 + int32(12)
	v311 = m.G0
	v313 = v311 - v305
	m.G0 = v313
	if v101 != 0 {
		goto L99
	} else {
		goto L100
	}
L85:
	;
	if v284 == int32(-1) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v287 == int32(4680360) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v292 = int32(4680360)
	goto L90
L89:
	;
	v292 = v284
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1012])) = v292
	goto L87
L91:
	;
	v297 = int32(-1)
	goto L93
L92:
	;
	v297 = v287
	goto L93
L93:
	;
	goto L84
L94:
	;
	if v297 != 0 {
		goto L138
	} else {
		goto L139
	}
L95:
	;
	v469 = int32(16)
	m.G0 = v313 + v469
	m.G0 = v306 + v469
	goto L94
L96:
	;
	v465 = v100 - v450
	goto L95
L97:
	;
	if v387 != 0 {
		goto L122
	} else {
		goto L123
	}
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v346 = v100
	v347 = v101
	v351 = v345
	goto L111
L99:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v100) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v318 == int32(0) {
		v465 = v298
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v387 = v100
	v388 = v101
	goto L97
L103:
	;
	v321 = v318
	v322 = v317
	v324 = v298
	goto L104
L104:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v321) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v465 = v344
	goto L95
L106:
	;
	v333 = int32(-1)
	v336 = F_wcrtomb(m, v313+int32(12), v321)
	mBase = m.M
	if v336 == v333 {
		v465 = v333
		goto L95
	} else {
		goto L109
	}
L107:
	;
	v339 = int32(1)
	goto L108
L108:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v344 = v324 + v339
	if v341 != 0 {
		v321 = v341
		v322 = v322 + int32(4)
		v324 = v344
		goto L104
	} else {
		goto L110
	}
L109:
	;
	v339 = v336
	goto L108
L110:
	;
	goto L105
L111:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	if base.Ui32(v355-int32(128)) <= base.Ui32(int32(-128)) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v387 = v377
	v388 = v380
	goto L97
L113:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v383 = v381 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v383
	if base.Ui32(int32(3)) < base.Ui32(v377) {
		v346 = v377
		v347 = v380
		v351 = v383
		goto L111
	} else {
		goto L121
	}
L114:
	;
	if v355 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v347))) = uint8(v355)
	v373 = int32(1)
	v377 = v346 - v373
	v380 = v347 + v373
	goto L113
L117:
	;
	v362 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v347))) = uint8(v362)
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v362
	v450 = v346
	goto L96
L118:
	;
	goto L119
L119:
	;
	v366 = int32(-1)
	v367 = F_wcrtomb(m, v347, v355)
	mBase = m.M
	if v367 == v366 {
		v465 = v366
		goto L95
	} else {
		goto L120
	}
L120:
	;
	v377 = v346 - v367
	v380 = v347 + v367
	goto L113
L121:
	;
	goto L112
L122:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v397 = v387
	v398 = v388
	v400 = v396
	goto L125
L123:
	;
	goto L124
L124:
	;
	v465 = v100
	goto L95
L125:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if base.Ui32(v406-int32(128)) <= base.Ui32(int32(-128)) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L124
L127:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v439 = v437 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v439
	if v433 != 0 {
		v397 = v433
		v398 = v436
		v400 = v439
		goto L125
	} else {
		goto L136
	}
L128:
	;
	if v406 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v406)
	v429 = int32(1)
	v433 = v397 - v429
	v436 = v398 + v429
	goto L127
L131:
	;
	v413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v413)
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v413
	v450 = v397
	goto L96
L132:
	;
	goto L133
L133:
	;
	v417 = int32(-1)
	v420 = F_wcrtomb(m, v313+int32(12), v406)
	mBase = m.M
	if v420 == v417 {
		v465 = v417
		goto L95
	} else {
		goto L134
	}
L134:
	;
	if base.Ui32(v397) < base.Ui32(v420) {
		v450 = v397
		goto L96
	} else {
		goto L135
	}
L135:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v425 = F_wcrtomb(m, v398, v424)
	mBase = m.M
	v433 = v397 - v420
	v436 = v398 + v420
	goto L127
L136:
	;
	goto L126
L137:
	;
	v489 = v465
	goto L36
L138:
	;
	if v297 == int32(-1) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	goto L144
L141:
	;
	v482 = int32(4680360)
	goto L143
L142:
	;
	v482 = v297
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1012])) = v482
	goto L140
L144:
	;
	goto L146
L146:
	;
	goto L137
L147:
	;
	if v489 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	goto L149
L149:
	;
	F_pfree(m, v38)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L15
	} else {
		goto L154
	}
L150:
	;
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494+v489))) = uint8(v496)
	goto L149
L151:
	;
	v493 = F__emscripten_memcpy_bulkmem(m, l0, v101, v489)
	mBase = m.M
	v494 = v493
	goto L153
L152:
	;
	v494 = l0
	goto L153
L153:
	;
	goto L150
L154:
	;
	F_pfree(m, v101)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L15
	} else {
		goto L155
	}
L155:
	;
	v623 = v489
	goto L5
L156:
	;
	v504 = F_strlen(m, l2)
	mBase = m.M
	v505 = v504
	goto L158
L157:
	;
	v505 = l3
	goto L158
L158:
	;
	if base.Ui32(l1) < base.Ui32(v505+int32(1)) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v623 = v505
	goto L5
L160:
	;
	if v505 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v513 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v511+v505))) = uint8(v513)
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	if v515 == v513 {
		goto L159
	} else {
		goto L165
	}
L162:
	;
	v510 = F__emscripten_memcpy_bulkmem(m, l0, l2, v505)
	mBase = m.M
	v511 = v510
	goto L164
L163:
	;
	v511 = l0
	goto L164
L164:
	;
	goto L161
L165:
	;
	v518 = l0
	v520 = v515
	v523 = v6
	goto L166
L166:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v527 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L159
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v518))) = uint8(v571)
	v574 = v571 & int32(255)
	goto L194
L169:
	;
	if v523 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v555 = v520 & int32(255)
	if v523 != 0 {
		goto L183
	} else {
		goto L184
	}
L172:
	;
	v530 = int32(255)
	v531 = v520 & v530
	if base.Ui32((v531-int32(65))&v530) < base.Ui32(int32(26)) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	goto L174
L174:
	;
	v541 = int32(255)
	v542 = v520 & v541
	if base.Ui32((v542-int32(97))&v541) < base.Ui32(int32(26)) {
		goto L180
	} else {
		goto L181
	}
L175:
	;
	v571 = v540
	goto L168
L176:
	;
	v540 = v531 | int32(32)
	goto L178
L177:
	;
	v540 = v531
	goto L178
L178:
	;
	goto L175
L179:
	;
	v571 = v551 & int32(255)
	goto L168
L180:
	;
	v551 = v542 - int32(32)
	goto L182
L181:
	;
	v551 = v542
	goto L182
L182:
	;
	goto L179
L183:
	;
	if base.Ui32(v555-int32(65)) < base.Ui32(int32(26)) {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	if base.Ui32(v555-int32(97)) < base.Ui32(int32(26)) {
		goto L191
	} else {
		goto L192
	}
L186:
	;
	v571 = v562
	goto L168
L187:
	;
	v562 = v555 | int32(32)
	goto L189
L188:
	;
	v562 = v555
	goto L189
L189:
	;
	goto L186
L190:
	;
	v571 = v569
	goto L168
L191:
	;
	v569 = v555 & int32(95)
	goto L193
L192:
	;
	v569 = v555
	goto L193
L193:
	;
	goto L190
L194:
	;
	v587 = v518 + int32(1)
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v588 != 0 {
		v518 = v587
		v520 = v588
		v523 = base.B2i32(base.Ui32(v574-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v574|int32(32)-int32(97)) < base.Ui32(int32(26)))
		goto L166
	} else {
		goto L195
	}
L195:
	;
	goto L167
L196:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L15
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L15
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(500494), int32(302), int32(503364))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L15
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	v628 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(383549)
	F_errmsg_internal(m, int32(502595), v12)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L15
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(499565), int32(1303), int32(383549))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L15
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	m.G0 = v642 + int32(16)
	v672 = v658
	goto L1
}
func F_pg_strupper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v13 - int32(98) {
	case 0:
		goto L2
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v603
L2:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	v591 = int32(0)
	v593 = F_convert_case(m, l0, l1, l2, l3, int32(2), v590, v591, v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L15
	} else {
		goto L179
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L15
	} else {
		goto L176
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18*int32(28))+uint32(_consts[979])))
	goto L7
L5:
	;
	v603 = v572
	goto L1
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L15
	} else {
		goto L172
	}
L7:
	;
	if int32(2) <= v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l3 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L11:
	;
	v29 = F_strlen(m, l2)
	mBase = m.M
	v30 = v29
	goto L13
L12:
	;
	v30 = l3
	goto L13
L13:
	;
	v32 = v30 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v32) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v37 = F_palloc(m, v32<<(uint(int32(2))%32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_char2wchar(m, v37, v32, l2, v30, l4)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = v43
	v49 = v45
	goto L21
L19:
	;
	v68 = v43
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76*int32(28))+uint32(_consts[979])))
	goto L25
L21:
	;
	v58 = F_casemap(m, v49, int32(1))
	mBase = m.M
	goto L23
L22:
	;
	v68 = v61
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v48<<(uint(int32(2))%32)))) = v58
	v61 = v48 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v37+v61<<(uint(int32(2))%32))))
	if v65 != 0 {
		v48 = v61
		v49 = v65
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v84 = v81*v68 + int32(1)
	v85 = F_palloc(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v84 == int32(0) {
		v473 = v43
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui32(v473+int32(1)) <= base.Ui32(l1) {
		goto L138
	} else {
		goto L139
	}
L28:
	;
	if l4 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v91 = int32(0)
	v97 = m.G0
	v98 = int32(16)
	v99 = v97 - v98
	m.G0 = v99
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v37
	v103 = v99 + int32(12)
	v104 = m.G0
	v106 = v104 - v98
	m.G0 = v106
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	goto L31
L31:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v271 = *(*int32)(unsafe.Add(mBase, _consts[1012]))
	if v268 != 0 {
		goto L76
	} else {
		goto L77
	}
L32:
	;
	v473 = v258
	goto L27
L33:
	;
	v262 = int32(16)
	m.G0 = v106 + v262
	m.G0 = v99 + v262
	goto L32
L34:
	;
	v258 = v84 - v243
	goto L33
L35:
	;
	if v180 != 0 {
		goto L60
	} else {
		goto L61
	}
L36:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v139 = v84
	v140 = v85
	v144 = v138
	goto L49
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v84) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 == int32(0) {
		v258 = v91
		goto L33
	} else {
		goto L41
	}
L40:
	;
	v180 = v84
	v181 = v85
	goto L35
L41:
	;
	v114 = v111
	v115 = v110
	v117 = v91
	goto L42
L42:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v114) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v258 = v137
	goto L33
L44:
	;
	v126 = int32(-1)
	v129 = F_wcrtomb(m, v106+int32(12), v114)
	mBase = m.M
	if v129 == v126 {
		v258 = v126
		goto L33
	} else {
		goto L47
	}
L45:
	;
	v132 = int32(1)
	goto L46
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v137 = v117 + v132
	if v134 != 0 {
		v114 = v134
		v115 = v115 + int32(4)
		v117 = v137
		goto L42
	} else {
		goto L48
	}
L47:
	;
	v132 = v129
	goto L46
L48:
	;
	goto L43
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if base.Ui32(v148-int32(128)) <= base.Ui32(int32(-128)) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v180 = v170
	v181 = v173
	goto L35
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v176 = v174 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v176
	if base.Ui32(int32(3)) < base.Ui32(v170) {
		v139 = v170
		v140 = v173
		v144 = v176
		goto L49
	} else {
		goto L59
	}
L52:
	;
	if v148 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v148)
	v166 = int32(1)
	v170 = v139 - v166
	v173 = v140 + v166
	goto L51
L55:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v155
	v243 = v139
	goto L34
L56:
	;
	goto L57
L57:
	;
	v159 = int32(-1)
	v160 = F_wcrtomb(m, v140, v148)
	mBase = m.M
	if v160 == v159 {
		v258 = v159
		goto L33
	} else {
		goto L58
	}
L58:
	;
	v170 = v139 - v160
	v173 = v140 + v160
	goto L51
L59:
	;
	goto L50
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v190 = v180
	v191 = v181
	v193 = v189
	goto L63
L61:
	;
	goto L62
L62:
	;
	v258 = v84
	goto L33
L63:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if base.Ui32(v199-int32(128)) <= base.Ui32(int32(-128)) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v232 = v230 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v232
	if v226 != 0 {
		v190 = v226
		v191 = v229
		v193 = v232
		goto L63
	} else {
		goto L74
	}
L66:
	;
	if v199 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v199)
	v222 = int32(1)
	v226 = v190 - v222
	v229 = v191 + v222
	goto L65
L69:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v206
	v243 = v190
	goto L34
L70:
	;
	goto L71
L71:
	;
	v210 = int32(-1)
	v213 = F_wcrtomb(m, v106+int32(12), v199)
	mBase = m.M
	if v213 == v210 {
		v258 = v210
		goto L33
	} else {
		goto L72
	}
L72:
	;
	if base.Ui32(v190) < base.Ui32(v213) {
		v243 = v190
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v218 = F_wcrtomb(m, v191, v217)
	mBase = m.M
	v226 = v190 - v213
	v229 = v191 + v213
	goto L65
L74:
	;
	goto L64
L75:
	;
	v282 = int32(0)
	v288 = m.G0
	v289 = int32(16)
	v290 = v288 - v289
	m.G0 = v290
	*(*int32)(unsafe.Add(mBase, uint32(v290)+12)) = v37
	v294 = v290 + int32(12)
	v295 = m.G0
	v297 = v295 - v289
	m.G0 = v297
	if v85 != 0 {
		goto L90
	} else {
		goto L91
	}
L76:
	;
	if v268 == int32(-1) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	if v271 == int32(4680360) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v276 = int32(4680360)
	goto L81
L80:
	;
	v276 = v268
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1012])) = v276
	goto L78
L82:
	;
	v281 = int32(-1)
	goto L84
L83:
	;
	v281 = v271
	goto L84
L84:
	;
	goto L75
L85:
	;
	if v281 != 0 {
		goto L129
	} else {
		goto L130
	}
L86:
	;
	v453 = int32(16)
	m.G0 = v297 + v453
	m.G0 = v290 + v453
	goto L85
L87:
	;
	v449 = v84 - v434
	goto L86
L88:
	;
	if v371 != 0 {
		goto L113
	} else {
		goto L114
	}
L89:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v330 = v84
	v331 = v85
	v335 = v329
	goto L102
L90:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v84) {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	if v302 == int32(0) {
		v449 = v282
		goto L86
	} else {
		goto L94
	}
L93:
	;
	v371 = v84
	v372 = v85
	goto L88
L94:
	;
	v305 = v302
	v306 = v301
	v308 = v282
	goto L95
L95:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v305) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v449 = v328
	goto L86
L97:
	;
	v317 = int32(-1)
	v320 = F_wcrtomb(m, v297+int32(12), v305)
	mBase = m.M
	if v320 == v317 {
		v449 = v317
		goto L86
	} else {
		goto L100
	}
L98:
	;
	v323 = int32(1)
	goto L99
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v328 = v308 + v323
	if v325 != 0 {
		v305 = v325
		v306 = v306 + int32(4)
		v308 = v328
		goto L95
	} else {
		goto L101
	}
L100:
	;
	v323 = v320
	goto L99
L101:
	;
	goto L96
L102:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if base.Ui32(v339-int32(128)) <= base.Ui32(int32(-128)) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v371 = v361
	v372 = v364
	goto L88
L104:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v367 = v365 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v367
	if base.Ui32(int32(3)) < base.Ui32(v361) {
		v330 = v361
		v331 = v364
		v335 = v367
		goto L102
	} else {
		goto L112
	}
L105:
	;
	if v339 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v339)
	v357 = int32(1)
	v361 = v330 - v357
	v364 = v331 + v357
	goto L104
L108:
	;
	v346 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v346)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v346
	v434 = v330
	goto L87
L109:
	;
	goto L110
L110:
	;
	v350 = int32(-1)
	v351 = F_wcrtomb(m, v331, v339)
	mBase = m.M
	if v351 == v350 {
		v449 = v350
		goto L86
	} else {
		goto L111
	}
L111:
	;
	v361 = v330 - v351
	v364 = v331 + v351
	goto L104
L112:
	;
	goto L103
L113:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v381 = v371
	v382 = v372
	v384 = v380
	goto L116
L114:
	;
	goto L115
L115:
	;
	v449 = v84
	goto L86
L116:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if base.Ui32(v390-int32(128)) <= base.Ui32(int32(-128)) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L115
L118:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v423 = v421 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v423
	if v417 != 0 {
		v381 = v417
		v382 = v420
		v384 = v423
		goto L116
	} else {
		goto L127
	}
L119:
	;
	if v390 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v390)
	v413 = int32(1)
	v417 = v381 - v413
	v420 = v382 + v413
	goto L118
L122:
	;
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v397)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v397
	v434 = v381
	goto L87
L123:
	;
	goto L124
L124:
	;
	v401 = int32(-1)
	v404 = F_wcrtomb(m, v297+int32(12), v390)
	mBase = m.M
	if v404 == v401 {
		v449 = v401
		goto L86
	} else {
		goto L125
	}
L125:
	;
	if base.Ui32(v381) < base.Ui32(v404) {
		v434 = v381
		goto L87
	} else {
		goto L126
	}
L126:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v409 = F_wcrtomb(m, v382, v408)
	mBase = m.M
	v417 = v381 - v404
	v420 = v382 + v404
	goto L118
L127:
	;
	goto L117
L128:
	;
	v473 = v449
	goto L27
L129:
	;
	if v281 == int32(-1) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	goto L135
L132:
	;
	v466 = int32(4680360)
	goto L134
L133:
	;
	v466 = v281
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1012])) = v466
	goto L131
L135:
	;
	goto L137
L137:
	;
	goto L128
L138:
	;
	if v473 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	goto L140
L140:
	;
	F_pfree(m, v37)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L15
	} else {
		goto L145
	}
L141:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v478+v473))) = uint8(v480)
	goto L140
L142:
	;
	v477 = F__emscripten_memcpy_bulkmem(m, l0, v85, v473)
	mBase = m.M
	v478 = v477
	goto L144
L143:
	;
	v478 = l0
	goto L144
L144:
	;
	goto L141
L145:
	;
	F_pfree(m, v85)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	v572 = v473
	goto L5
L147:
	;
	v488 = F_strlen(m, l2)
	mBase = m.M
	v489 = v488
	goto L149
L148:
	;
	v489 = l3
	goto L149
L149:
	;
	if base.Ui32(l1) < base.Ui32(v489+int32(1)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v572 = v489
	goto L5
L151:
	;
	if v489 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v497 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v495+v489))) = uint8(v497)
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	if v499 == v497 {
		goto L150
	} else {
		goto L156
	}
L153:
	;
	v494 = F__emscripten_memcpy_bulkmem(m, l0, l2, v489)
	mBase = m.M
	v495 = v494
	goto L155
L154:
	;
	v495 = l0
	goto L155
L155:
	;
	goto L152
L156:
	;
	v502 = l0
	v504 = v499
	goto L157
L157:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v510 == int32(1) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L150
L159:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v502))) = uint8(v535)
	v538 = v502 + int32(1)
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	if v539 != 0 {
		v502 = v538
		v504 = v539
		goto L157
	} else {
		goto L171
	}
L160:
	;
	v513 = int32(255)
	v514 = v504 & v513
	if base.Ui32((v514-int32(97))&v513) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	goto L162
L162:
	;
	v527 = v504 & int32(255)
	if base.Ui32(v527-int32(97)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L163:
	;
	v535 = v523 & int32(255)
	goto L159
L164:
	;
	v523 = v514 - int32(32)
	goto L166
L165:
	;
	v523 = v514
	goto L166
L166:
	;
	goto L163
L167:
	;
	v535 = v534
	goto L159
L168:
	;
	v534 = v527 & int32(95)
	goto L170
L169:
	;
	v534 = v527
	goto L170
L170:
	;
	goto L167
L171:
	;
	goto L158
L172:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L15
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L15
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(500494), int32(390), int32(503347))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L15
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
	v577 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(217864)
	F_errmsg_internal(m, int32(502595), v11)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L15
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(499565), int32(1322), int32(217864))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L15
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v603 = v593
	goto L1
}
func F_pg_timezone_abbrevs_abbrevs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+82)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+80)) = uint16(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v18 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L87
	}
L2:
	;
	v21 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	goto L10
L5:
	;
	return int32(0)
L6:
	;
	v25 = int32(4515712)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v31 = F_palloc(m, int32(4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v31
	v39 = F_get_call_result_type(m, l0, v33, v11+int32(40))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v39 != int32(1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v43
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	goto L4
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[953]))
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	m.G0 = v11 + int32(96)
	return v360
L12:
	;
	v70 = v53 + v55<<(uint(int32(4))%32)
	v72 = v70 + int32(8)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+19)))
	switch v73 - int32(5) {
	case 0:
		goto L19
	case 1:
		goto L22
	case 2:
		goto L21
	default:
		goto L20
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v55 < v56 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = int32(2)
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
	v360 = int32(0)
	goto L11
L18:
	;
	v147 = v11 + int32(69)
	goto L40
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v143 = v140
	v145 = int32(0)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L34
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v79 = v53 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 != 0 {
		v113 = v80
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v143 = v76
	v145 = int32(1)
	goto L18
L23:
	;
	v117 = *(*int64)(unsafe.Add(mBase, _consts[139]))
	v120 = F_DetermineTimeZoneAbbrevOffsetTS(m, v117, v72, v113, v11+int32(40))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L33
	}
L24:
	;
	v82 = v79 + int32(4)
	v83 = F_pg_tzset(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v83
	if v83 != 0 {
		v113 = v83
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(0)
	v88 = F_errsave_start(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v88 == int32(0) {
		v113 = v86
		goto L23
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v82
	F_errmsg(m, int32(438536), v11+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v72
	F_errdetail(m, int32(666762), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, int32(0), int32(499320), int32(4258), int32(212638))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v113 = v86
	goto L23
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v143 = int32(0) - v120
	v145 = base.B2i32(v123 != int32(0))
	goto L18
L34:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72)+11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v130
	F_errmsg_internal(m, int32(476738), v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(499320), int32(5294), int32(114578))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+69)))
	if v263 != 0 {
		goto L69
	} else {
		goto L70
	}
L38:
	;
	v260 = F_strlen(m, v249)
	mBase = m.M
	goto L37
L40:
	;
	goto L41
L41:
	;
	v154 = int32(10)
	if (v147^v72)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v253)
	goto L38
L43:
	;
	v234 = v229
	v235 = v230
	v236 = v231
	goto L65
L44:
	;
	if v224 == int32(0) {
		v249 = v222
		v250 = v223
		goto L42
	} else {
		goto L64
	}
L45:
	;
	v222 = v72
	v223 = v147
	v224 = v154
	goto L44
L46:
	;
	goto L47
L47:
	;
	if v72&int32(3) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v191 == int32(0) {
		v249 = v188
		v250 = v189
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v188 = v72
	v189 = v147
	v190 = v154
	v191 = int32(1)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v167 = v72
	v168 = v147
	v169 = v154
	goto L52
L52:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
	if v171 == int32(0) {
		v229 = v167
		v230 = v168
		v231 = v169
		goto L43
	} else {
		goto L54
	}
L53:
	;
	v188 = v182
	v189 = v176
	v190 = v178
	v191 = v180
	goto L48
L54:
	;
	v175 = int32(1)
	v176 = v168 + v175
	v178 = v169 - v175
	v179 = int32(0)
	v180 = base.B2i32(v178 != v179)
	v182 = v167 + v175
	if v182&int32(3) == v179 {
		v188 = v182
		v189 = v176
		v190 = v178
		v191 = v180
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v178 != 0 {
		v167 = v182
		v168 = v176
		v169 = v178
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v194 == int32(0) {
		v222 = v188
		v223 = v189
		v224 = v190
		goto L44
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v190) < base.Ui32(int32(4)) {
		v222 = v188
		v223 = v189
		v224 = v190
		goto L44
	} else {
		goto L59
	}
L59:
	;
	v200 = v188
	v201 = v189
	v202 = v190
	goto L60
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v208 = int32(-2139062144)
	if (int32(16843008)-v205|v205)&v208 != v208 {
		v229 = v200
		v230 = v201
		v231 = v202
		goto L43
	} else {
		goto L62
	}
L61:
	;
	v222 = v216
	v223 = v214
	v224 = v218
	goto L44
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v205
	v213 = int32(4)
	v214 = v201 + v213
	v216 = v200 + v213
	v218 = v202 - v213
	if base.Ui32(int32(3)) < base.Ui32(v218) {
		v200 = v216
		v201 = v214
		v202 = v218
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v229 = v222
	v230 = v223
	v231 = v224
	goto L43
L65:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v238)
	if v238 == int32(0) {
		v249 = v234
		v250 = v235
		goto L42
	} else {
		goto L67
	}
L66:
	;
	v249 = v245
	v250 = v243
	goto L42
L67:
	;
	v242 = int32(1)
	v243 = v235 + v242
	v245 = v234 + v242
	v247 = v236 - v242
	if v247 != 0 {
		v234 = v245
		v235 = v243
		v236 = v247
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v268 = v11 + int32(69)
	v269 = v263
	goto L72
L70:
	;
	goto L71
L71:
	;
	v301 = F_cstring_to_text(m, v11+int32(69))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L5
	} else {
		goto L79
	}
L72:
	;
	v274 = int32(255)
	v275 = v269 & v274
	if base.Ui32((v275-int32(97))&v274) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L71
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v286)
	v289 = v268 + int32(1)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v290 != 0 {
		v268 = v289
		v269 = v290
		goto L72
	} else {
		goto L78
	}
L75:
	;
	v284 = v275 - int32(32)
	goto L77
L76:
	;
	v284 = v275
	goto L77
L77:
	;
	v286 = v284 & int32(255)
	goto L74
L78:
	;
	goto L73
L79:
	;
	v303 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v303
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = base.I64_extend_i32_s(v143) * int64(1000000)
	v313 = v11 + int32(40)
	v315 = F_palloc(m, int32(16))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v318 = int64(*(*int32)(unsafe.Add(mBase, uint32(v313)+12)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v313)+16)))
	v322 = v318 + v319*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v322-int64(2147483648)) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v315
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v337 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v346 = F_heap_form_tuple(m, v341, v11+int32(84), v11+int32(80))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L85
	}
L82:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v315)+12)) = uint32(v322)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v313)))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v330
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346)+16))
	v349 = F_HeapTupleHeaderGetDatum(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v351 + int64(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+20)) = int32(1)
	v360 = v349
	goto L11
L87:
	;
	F_errmsg_internal(m, int32(367739), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(499320), int32(5247), int32(114578))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_total_relation_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_try_relation_open(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_calculate_table_size(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = F_calculate_indexes_size(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_relation_close(m, v7, int32(1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = F_Int64GetDatum(m, v17+v19)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return v25
						}
					}
				}
			}
		}
	}
}
func F_pg_typeof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_get_fn_expr_argtype(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_ultostr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	if l1 == int32(0) {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v12)
		v133 = int32(1)
	} else {
		v18 = int32(1233)
		v23 = int32(base.Ui32((base.I32_clz(l1)^int32(31))*v18+v18) >> (uint(int32(12)) % 32))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_consts[987])))
		v30 = v23 + base.B2i32(base.Ui32(v28) <= base.Ui32(l1))
		v31 = int32(0)
		if base.Ui32(l1) < base.Ui32(int32(10000)) {
			v77 = v31
			v78 = l1
		} else {
			v35 = l1
			v37 = v31
			for {
				v44 = l0 + v30 - v37
				v45 = int32(4)
				v48 = base.I32_div_u_s(v35, int32(10000))
				v51 = v48*int32(-10000) + v35
				v52 = int32(100)
				v53 = base.I32_div_u_s(v51, v52)
				v54 = int32(1)
				v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53<<(uint(v54)%32))+uint32(_consts[988]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v44-v45))) = uint16(v58)
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v51-v53*v52)<<(uint(v54)%32))+uint32(_consts[988]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v44-int32(2)))) = uint16(v69)
				v72 = v37 + v45
				if base.Ui32(int32(99999999)) < base.Ui32(v35) {
					v35 = v48
					v37 = v72
					continue
				} else {
					break
				}
				break
			}
			v77 = v72
			v78 = v48
		}
		if base.Ui32(v78) < base.Ui32(int32(100)) {
			v107 = v78
			v108 = v77
		} else {
			v88 = int32(2)
			v90 = int32(65535)
			v92 = int32(100)
			v93 = base.I32_div_u_s(v78&v90, v92)
			v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v78-v93*v92)&v90<<(uint(int32(1))%32))+uint32(_consts[988]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0+v30-v77-v88))) = uint16(v103)
			v107 = v93
			v108 = v77 | v88
		}
		if base.Ui32(int32(10)) <= base.Ui32(v107) {
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107<<(uint(int32(1))%32))+uint32(_consts[988]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0+v30-v108-int32(2)))) = uint16(v119)
			v133 = v30
		} else {
			v122 = v107 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v122)
			v133 = v30
		}
	}
	return v133 + l0
}
func F_pg_utf2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v19 == int32(0) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(0)
	return v113
L6:
	;
	goto L5
L7:
	;
	if int32(0) <= base.I32_extend8_s(v19) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v97
	v102 = v18 + int32(1)
	v104 = v14 + int32(4)
	v105 = v15 + v98
	if int32(0) < v105 {
		v13 = v99
		v14 = v104
		v15 = v105
		v18 = v102
		goto L4
	} else {
		goto L21
	}
L9:
	;
	v97 = v19
	v98 = int32(-1)
	v99 = v13 + int32(1)
	goto L8
L10:
	;
	if v19&int32(224) == int32(192) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v15 == int32(1) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v19&int32(240) == int32(224) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v97 = v19<<(uint(int32(6))%32)&int32(1984) | v35&int32(63)
	v98 = int32(-2)
	v99 = v13 + int32(2)
	goto L8
L15:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v19&int32(248) != int32(240) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v49 = int32(63)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v97 = v48&v49 | (v19<<(uint(int32(12))%32)&int32(61440) | v55&v49<<(uint(int32(6))%32))
	v98 = int32(-3)
	v99 = v13 + int32(3)
	goto L8
L19:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v72 = int32(63)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v97 = v71&v72 | (v19<<(uint(int32(18))%32)&int32(1835008) | v78&v72<<(uint(int32(12))%32) | v84&v72<<(uint(int32(6))%32))
	v98 = int32(-4)
	v99 = v13 + int32(4)
	goto L8
L21:
	;
	v109 = v104
	v113 = v102
	goto L6
}
func F_pg_verifymbstr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_consts[1127])))
	v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if l1 != v13 {
			F_report_invalid_encoding(m, v7, l0+v13, l1-v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return
		}
	}
}
func F_pg_wchar2utf_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v17 = v4
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v105)
	return v103
L6:
	;
	if base.Ui32(v19) <= base.Ui32(int32(127)) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v100 = v14
	v103 = v17
	goto L8
L8:
	;
	goto L5
L9:
	;
	v91 = int32(1)
	v95 = v14 + v90
	v96 = v90 + v17
	if v91 < v15 {
		v13 = v13 + int32(4)
		v14 = v95
		v15 = v15 - v91
		v17 = v96
		goto L4
	} else {
		goto L24
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v19)
	v90 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v19) <= base.Ui32(int32(2047)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v75 = v19&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v70+v14))) = uint8(v75)
	if v69&int32(224) == int32(192) {
		v90 = int32(2)
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v29 = int32(base.Ui32(v19)>>(uint(int32(6))%32)) | int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v29)
	v69 = v29
	v70 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v19) <= base.Ui32(int32(65535)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v37 = int32(base.Ui32(v19)>>(uint(int32(12))%32)) | int32(-32)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v37)
	v44 = int32(base.Ui32(v19)>>(uint(int32(6))%32))&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v44)
	v69 = v37
	v70 = int32(2)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v49 = int32(63)
	v51 = int32(128)
	v52 = int32(base.Ui32(v19)>>(uint(int32(6))%32))&v49 | v51
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v52)
	v59 = int32(base.Ui32(v19)>>(uint(int32(12))%32))&v49 | v51
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v59)
	v66 = int32(base.Ui32(v19)>>(uint(int32(18))%32))&int32(7) | int32(-16)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v66)
	v69 = v66
	v70 = int32(3)
	goto L13
L20:
	;
	if v69&int32(240) == int32(224) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = int32(3)
	goto L23
L22:
	;
	v88 = int32(4)
	goto L23
L23:
	;
	v90 = v88
	goto L9
L24:
	;
	v100 = v95
	v103 = v96
	goto L8
}
func F_pg_xact_commit_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_TransactionIdGetCommitTsData(m, v8, v6+int32(8), v2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v23 = v2
			m.G0 = v6 + int32(16)
			return v23
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v21 = F_Int64GetDatum(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = v21
				m.G0 = v6 + int32(16)
				return v23
			}
		}
	}
}
