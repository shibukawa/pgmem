package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = l2
	v20 = F__emscripten_memset_bulkmem(m, v12+int32(160), base.I32_extend8_s(v6), int32(40))
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v21
	v30 = F_printf_core(m, int32(0), l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		if v30 < int32(0) {
			v122 = int32(-1)
			m.G0 = v12 + int32(208)
			return v122
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v38 = int32(0)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40 & int32(-33)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v44 == v38 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(80)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v12
				v56 = v53
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v59 - int32(1) | v59
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v64&int32(8) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 | int32(32)
					v81 = int32(-1)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v73
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + v76
					v81 = int32(0)
				}
				if v81 != 0 {
					v92 = int32(-1)
					v93 = v56
					if v93 != 0 {
						v96 = int32(0)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v101 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							if v106 != 0 {
								v110 = v92
							} else {
								v110 = int32(-1)
							}
							v111 = v110
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					} else {
						v111 = v92
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
						if v113&int32(32) != 0 {
							v119 = int32(-1)
						} else {
							v119 = v111
						}
						if v37 < v38 {
							v122 = v119
						} else {
							v122 = v119
						}
						m.G0 = v12 + int32(208)
						return v122
					}
				} else {
					v83 = v56
					v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = v90
						v93 = v83
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					}
				}
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v55 != 0 {
					v83 = v6
					v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = v90
						v93 = v83
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					}
				} else {
					v56 = v6
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v59 - int32(1) | v59
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v64&int32(8) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 | int32(32)
						v81 = int32(-1)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v73
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + v76
						v81 = int32(0)
					}
					if v81 != 0 {
						v92 = int32(-1)
						v93 = v56
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					} else {
						v83 = v56
						v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = v90
							v93 = v83
							if v93 != 0 {
								v96 = int32(0)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
									if v106 != 0 {
										v110 = v92
									} else {
										v110 = int32(-1)
									}
									v111 = v110
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
									if v113&int32(32) != 0 {
										v119 = int32(-1)
									} else {
										v119 = v111
									}
									if v37 < v38 {
										v122 = v119
									} else {
										v122 = v119
									}
									m.G0 = v12 + int32(208)
									return v122
								}
							} else {
								v111 = v92
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						}
					}
				}
			}
		}
	}
}
func F_vac_update_datfrozenxid(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
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
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var __phi569 int32
	_ = __phi569
	var v571 int32
	_ = v571
	var __phi571 int32
	_ = __phi571
	var v575 int32
	_ = v575
	var __phi575 int32
	_ = __phi575
	var v578 int32
	_ = v578
	var __phi578 int32
	_ = __phi578
	var v584 int64
	_ = v584
	var __phi584 int64
	_ = __phi584
	var v585 int64
	_ = v585
	var __phi585 int64
	_ = __phi585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int64
	_ = v617
	var v618 int64
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v677 int64
	_ = v677
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int64
	_ = v697
	var v698 int64
	_ = v698
	var v704 int32
	_ = v704
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int64
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int64
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int64
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v838 int64
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int64
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int64
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1106 int64
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1131 int64
	_ = v1131
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	v1 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(16908288)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = int64(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
	v38 = F_LockAcquire(m, v26, int32(7), v1, v1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	m.G0 = v26 + int32(16)
	v44 = F_GetOldestNonRemovableTransactionId(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v48 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = base.I32_wrap_i64(v48)
	v51 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v55 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v62)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	v57 = int32(0)
	v62 = F_systable_beginscan(m, v55, v57, v57, v57, v57, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v64 = F_systable_getnext(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v64 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v68 = v64
	v70 = v44
	v72 = v46
	goto L14
L12:
	;
	v144 = v44
	v146 = v46
	goto L13
L13:
	;
	v164 = v144
	v166 = v146
	v179 = int32(1)
	goto L7
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
	v87 = v85 + v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+136))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+140))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v90 == int32(114) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v144 = v136
	v146 = v137
	goto L13
L16:
	;
	v138 = F_systable_getnext(m, v62)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L17:
	;
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v93 == int32(109) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v96 != int32(116) {
		v136 = v70
		v137 = v72
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v99 = int32(0)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v88))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == v99 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v125 = v70
	goto L23
L23:
	;
	if v89 == int32(0) {
		v136 = v125
		v137 = v72
		goto L16
	} else {
		goto L36
	}
L24:
	;
	if v111 != 0 {
		v164 = v70
		v166 = v72
		v179 = v99
		goto L7
	} else {
		goto L28
	}
L25:
	;
	v111 = base.B2i32(base.Ui32(v50) < base.Ui32(v88))
	goto L24
L26:
	;
	goto L27
L27:
	;
	v111 = int32(base.Ui32(v50-v88) >> (uint(int32(31)) % 32))
	goto L24
L28:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v70))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v123 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v123 = base.B2i32(base.Ui32(v88) < base.Ui32(v70))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v123 = int32(base.Ui32(v88-v70) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	v124 = v88
	goto L35
L34:
	;
	v124 = v70
	goto L35
L35:
	;
	v125 = v124
	goto L23
L36:
	;
	goto L37
L37:
	;
	if int32(base.Ui32(v51-v89)>>(uint(int32(31))%32)) != 0 {
		v164 = v125
		v166 = v72
		v179 = int32(0)
		goto L7
	} else {
		goto L38
	}
L38:
	;
	goto L39
L39:
	;
	if int32(base.Ui32(v89-v72)>>(uint(int32(31))%32)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = v89
	goto L42
L41:
	;
	v135 = v72
	goto L42
L42:
	;
	v136 = v125
	v137 = v135
	goto L16
L43:
	;
	if v138 != 0 {
		v68 = v138
		v70 = v136
		v72 = v137
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L15
L45:
	;
	F_sequence_close(m, v55, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v179 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	m.G0 = v22 + int32(96)
	return
L48:
	;
	v187 = int32(0)
	v190 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	F_ScanKeyInit(m, v22+int32(32), int32(1), int32(3), int32(184), v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_systable_inplace_update_begin(m, v190, int32(2672), v22+int32(32), v22+int32(92), v22+int32(28))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	if v210 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_pfree(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L84
	}
L53:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_systable_inplace_update_finish(m, v278, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L83
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+88)) = v166
	v277 = v166
	goto L53
L55:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+22)))
	v213 = v211 + v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	if v214 == v164 {
		v244 = v164
		v245 = v187
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L80
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	if v166 != v246 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v164))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v214)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+84)) = v164
	v244 = v164
	v245 = int32(1)
	goto L58
L61:
	;
	if v227 != 0 {
		goto L60
	} else {
		goto L65
	}
L62:
	;
	v227 = base.B2i32(base.Ui32(v214) < base.Ui32(v164))
	goto L61
L63:
	;
	goto L64
L64:
	;
	v227 = int32(base.Ui32(v214-v164) >> (uint(int32(31)) % 32))
	goto L61
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v228))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v240 != 0 {
		goto L60
	} else {
		goto L70
	}
L67:
	;
	v240 = base.B2i32(base.Ui32(v50) < base.Ui32(v228))
	goto L66
L68:
	;
	goto L69
L69:
	;
	v240 = int32(base.Ui32(v50-v228) >> (uint(int32(31)) % 32))
	goto L66
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	v244 = v241
	v245 = v187
	goto L58
L71:
	;
	goto L74
L72:
	;
	v256 = v166
	goto L73
L73:
	;
	if v245 != 0 {
		v277 = v256
		goto L53
	} else {
		goto L78
	}
L74:
	;
	if int32(base.Ui32(v246-v166)>>(uint(int32(31))%32)) != 0 {
		goto L54
	} else {
		goto L75
	}
L75:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	goto L76
L76:
	;
	if int32(base.Ui32(v51-v251)>>(uint(int32(31))%32)) != 0 {
		goto L54
	} else {
		goto L77
	}
L77:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	v256 = v255
	goto L73
L78:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_systable_inplace_update_cancel(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v283 = v256
	v284 = int32(0)
	goto L52
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v266
	F_errmsg_internal(m, int32(49295), v22)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(495260), int32(1776), int32(431219))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v283 = v277
	v284 = int32(1)
	goto L52
L84:
	;
	F_sequence_close(m, v190, int32(3))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v284 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v293 = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v299 = F_LWLockAcquire(m, v295+int32(384), v293)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v346 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L101
	}
L89:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+36))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v302)+20))
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v302)+8))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)+16))
	v308 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v308+int32(384))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v306) < base.Ui32(int32(3)) {
		v338 = v293
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v338 == int32(0) {
		goto L47
	} else {
		goto L100
	}
L92:
	;
	if v304 == int32(0) {
		v338 = v293
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v317 = base.I32_wrap_i64(v305)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v304))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v317)) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v329 != 0 {
		v338 = v293
		goto L91
	} else {
		goto L98
	}
L95:
	;
	v329 = base.B2i32(base.Ui32(v304) <= base.Ui32(v317))
	goto L94
L96:
	;
	goto L97
L97:
	;
	v329 = base.B2i32(int32(0) <= v317-v304)
	goto L94
L98:
	;
	v331 = int32(0)
	v334 = F_SearchSysCacheExists(m, int32(21), v303, v331, v331, v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v338 = v334 ^ int32(1)
	goto L91
L100:
	;
	goto L88
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v353 = F_LWLockAcquire(m, v349+int32(5888), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v359 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+188))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+12))
	m.T0[v504].(func(*base.Module, int32))(m, v363)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L152
	}
L104:
	;
	v361 = int32(0)
	v363 = F_table_beginscan_catalog(m, v359, v361, v361)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v365 = F_heap_getnext(m, v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v365 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v487 = v244
	v489 = v283
	v491 = v1
	v493 = v1
	v495 = v356
	v496 = v356
	goto L103
L108:
	;
	goto L109
L109:
	;
	v369 = base.I32_wrap_i64(v346)
	v372 = v365
	v374 = v244
	v376 = v283
	v378 = v1
	v380 = v1
	v382 = v356
	v383 = v356
	goto L110
L110:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+22)))
	v391 = v389 + v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+84))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+88))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+80))
	goto L113
L111:
	;
	v487 = v475
	v489 = v476
	v491 = v477
	v493 = v478
	v495 = v479
	v496 = v480
	goto L103
L112:
	;
	v481 = F_heap_getnext(m, v363)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L150
	}
L113:
	;
	if v394 == int32(-2) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v399 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v392))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	if v399 == int32(0) {
		v475 = v374
		v476 = v376
		v477 = v378
		v478 = v380
		v479 = v382
		v480 = v383
		goto L112
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v391 + int32(4)
	F_errmsg_internal(m, int32(431243), v22+int32(16))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(495260), int32(1905), int32(326252))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v475 = v374
	v476 = v376
	v477 = v378
	v478 = v380
	v479 = v382
	v480 = v383
	goto L112
L121:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v392))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v369)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	if v427 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v427 = base.B2i32(base.Ui32(v50) < base.Ui32(v392))
	goto L122
L124:
	;
	goto L125
L125:
	;
	v427 = int32(base.Ui32(v50-v392) >> (uint(int32(31)) % 32))
	goto L122
L126:
	;
	goto L129
L127:
	;
	goto L128
L128:
	;
	v436 = int32(1)
	goto L121
L129:
	;
	if int32(base.Ui32(v51-v393)>>(uint(int32(31))%32)) == int32(0) {
		v436 = v378
		goto L121
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L146
L132:
	;
	if v448 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v448 = base.B2i32(base.Ui32(v369) < base.Ui32(v392))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v448 = int32(base.Ui32(v369-v392) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v465 = v374
	v466 = int32(1)
	v467 = v383
	goto L131
L137:
	;
	goto L138
L138:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v374))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v392)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v461 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v461 = base.B2i32(base.Ui32(v392) < base.Ui32(v374))
	goto L139
L141:
	;
	goto L142
L142:
	;
	v461 = int32(base.Ui32(v392-v374) >> (uint(int32(31)) % 32))
	goto L139
L143:
	;
	v465 = v374
	v466 = v380
	v467 = v383
	goto L131
L144:
	;
	goto L145
L145:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v465 = v392
	v466 = v380
	v467 = v464
	goto L131
L146:
	;
	if int32(base.Ui32(v393-v376)>>(uint(int32(31))%32)) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v472 = v393
	v473 = v471
	goto L149
L148:
	;
	v472 = v376
	v473 = v382
	goto L149
L149:
	;
	v475 = v465
	v476 = v472
	v477 = v436
	v478 = v466
	v479 = v473
	v480 = v467
	goto L112
L150:
	;
	if v481 != 0 {
		v372 = v481
		v374 = v475
		v376 = v476
		v378 = v477
		v380 = v478
		v382 = v479
		v383 = v480
		goto L110
	} else {
		goto L151
	}
L151:
	;
	goto L111
L152:
	;
	F_sequence_close(m, v359, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v493 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v1229+int32(5888))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L321
	}
L155:
	;
	v512 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if v491 != 0 {
		goto L154
	} else {
		goto L163
	}
L158:
	;
	if v512 == int32(0) {
		goto L154
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(141720), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errdetail(m, int32(569850), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(495260), int32(1951), int32(326252))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	goto L154
L163:
	;
	v529 = int32(0)
	v531 = m.G0
	v533 = v531 - int32(16)
	m.G0 = v533
	v536 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v540 = F_LWLockAcquire(m, v536+int32(6016), int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v547 = F_LWLockAcquire(m, v543+int32(3456), int32(1))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _consts[449]))
	v551 = *(*int64)(unsafe.Add(mBase, uint32(v550)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v550)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v553
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v550)))
	v558 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v558+int32(3456))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	if base.B2i32(v553 == v555)&base.B2i32(v551 == v556) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v725+int32(6016))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L204
	}
L168:
	;
	__phi569 = v529
	__phi571 = v553
	__phi575 = int32(-1)
	__phi578 = v529
	__phi584 = int64(-1)
	__phi585 = v551
	v569 = __phi569
	v571 = __phi571
	v575 = __phi575
	v578 = __phi578
	v584 = __phi584
	v585 = __phi585
	goto L169
L169:
	;
	if v584 != v585 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v640 < int32(0) {
		goto L167
	} else {
		goto L199
	}
L171:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if int32(0) <= v575 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v639 = v569
	v640 = v575
	v641 = v578
	goto L173
L173:
	;
	v642 = v639 + v571
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	if base.Ui32(v643) < base.Ui32(int32(3)) {
		v667 = v641
		goto L183
	} else {
		goto L184
	}
L174:
	;
	if v578 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v613 = v578
	v614 = v589
	goto L176
L176:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+28))
	v617 = int64(*(*uint16)(unsafe.Add(mBase, _consts[512])))
	v618 = base.I64_rem_s(v585, v617)
	v624 = F_LWLockAcquire(m, v615+base.I32_wrap_i64(v618)<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L181
	}
L177:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v594 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v592+v575))) = uint8(v594)
	v597 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v598 = v597
	goto L179
L178:
	;
	v598 = v589
	goto L179
L179:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v598)+28))
	v602 = int64(*(*uint16)(unsafe.Add(mBase, _consts[512])))
	v603 = base.I64_rem_s(v584, v602)
	F_LWLockRelease(m, v600+base.I32_wrap_i64(v603)<<(uint(int32(7))%32))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v613 = int32(0)
	v614 = v611
	goto L176
L181:
	;
	v629 = F_SimpleLruReadPage(m, int32(4386024), v585, int32(1), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633+v629<<(uint(int32(2))%32))))
	v639 = v637
	v640 = v629
	v641 = v613
	goto L173
L183:
	;
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v533)+8))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
	v671 = v669 + v670
	v675 = base.B2i32(base.Ui32(v671-int32(8173)) < base.Ui32(int32(-8193)))
	v677 = v668 + base.I64_extend_i32_u(v675)
	*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v677
	if base.Ui32(v671-int32(8173)) < base.Ui32(int32(-8193)) {
		goto L194
	} else {
		goto L195
	}
L184:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v643)) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v657 == int32(0) {
		v667 = v641
		goto L183
	} else {
		goto L189
	}
L186:
	;
	v657 = base.B2i32(base.Ui32(v643) < base.Ui32(v487))
	goto L185
L187:
	;
	goto L188
L188:
	;
	v657 = int32(base.Ui32(v643-v487) >> (uint(int32(31)) % 32))
	goto L185
L189:
	;
	v662 = F_TransactionIdDidCommit(m, v643)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	if v662 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v664 = int32(2)
	goto L193
L192:
	;
	v664 = int32(0)
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+8)) = v664
	v667 = int32(1)
	goto L183
L194:
	;
	v680 = int32(0)
	goto L196
L195:
	;
	v680 = v671
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v680
	if v677 != v556 {
		__phi569 = v639
		__phi571 = v680
		__phi575 = v640
		__phi578 = v667
		__phi584 = v585
		__phi585 = v677
		v569 = __phi569
		v571 = __phi571
		v575 = __phi575
		v578 = __phi578
		v584 = __phi584
		v585 = __phi585
		goto L169
	} else {
		goto L197
	}
L197:
	;
	if v680 != v555 {
		__phi569 = v639
		__phi571 = v680
		__phi575 = v640
		__phi578 = v667
		__phi584 = v585
		__phi585 = v677
		v569 = __phi569
		v571 = __phi571
		v575 = __phi575
		v578 = __phi578
		v584 = __phi584
		v585 = __phi585
		goto L169
	} else {
		goto L198
	}
L198:
	;
	goto L170
L199:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if v667 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)+12))
	v690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v688+v640))) = uint8(v690)
	v693 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v694 = v693
	goto L202
L201:
	;
	v694 = v687
	goto L202
L202:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+28))
	v697 = int64(*(*uint16)(unsafe.Add(mBase, _consts[512])))
	v698 = base.I64_rem_s(v585, v697)
	F_LWLockRelease(m, v695+base.I32_wrap_i64(v698)<<(uint(int32(7))%32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	goto L167
L204:
	;
	m.G0 = v533 + int32(16)
	v734 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v738 = F_LWLockAcquire(m, v734+int32(4992), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+40))
	if v742 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v763+int32(4992))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L213
	}
L207:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v742)) == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v756 == int32(0) {
		goto L206
	} else {
		goto L212
	}
L209:
	;
	v756 = base.B2i32(base.Ui32(v742) < base.Ui32(v487))
	goto L208
L210:
	;
	goto L211
L211:
	;
	v756 = int32(base.Ui32(v742-v487) >> (uint(int32(31)) % 32))
	goto L208
L212:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v760)+40)) = v487
	goto L206
L213:
	;
	v768 = m.G0
	v770 = v768 - int32(32)
	m.G0 = v770
	*(*int64)(unsafe.Add(mBase, uint32(v770)+8)) = base.I64_extend_i32_u(int32(base.Ui32(v487) >> (uint(int32(15)) % 32)))
	v780 = F_SlruScanDirectory(m, int32(4383680), int32(288), v770+int32(8))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	if v780 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	F_AdvanceOldestClogXid(m, v487)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v806 = int32(32)
	m.G0 = v770 + v806
	v809 = m.G0
	v811 = v809 - v806
	m.G0 = v811
	v814 = base.I32_div_u_s(v487, int32(819))
	*(*int64)(unsafe.Add(mBase, uint32(v811)+8)) = base.I64_extend_i32_u(v814)
	v821 = F_SlruScanDirectory(m, int32(4383764), int32(288), v811+int32(8))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L224
	}
L218:
	;
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v770)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+28)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v770)+24)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v770)+16)) = v784
	F_XLogBeginInsert(m)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v790 = int32(16)
	F_XLogRegisterData(m, v770+v790, v790)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v797 = F_XLogInsert(m, int32(3), int32(16))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_XLogFlush(m, v797)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v770)+8))
	F_SimpleLruTruncate(m, int32(4383680), v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	goto L217
L224:
	;
	if v821 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v811)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v811)+24)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v811)+16)) = v823
	F_XLogBeginInsert(m)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	m.G0 = v811 + int32(32)
	v845 = m.G0
	v847 = v845 - int32(144)
	m.G0 = v847
	v850 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v854 = F_LWLockAcquire(m, v850+int32(5248), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L232
	}
L228:
	;
	F_XLogRegisterData(m, v811+int32(16), int32(12))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v835 = F_XLogInsert(m, int32(18), int32(16))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v811)+8))
	F_SimpleLruTruncate(m, int32(4383764), v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	goto L227
L232:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v861 = F_LWLockAcquire(m, v857+int32(1664), int32(1))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	v869 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v869+int32(1664))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	if v489-v867 <= int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	m.G0 = v847 + int32(144)
	F_SetTransactionIdLimit(m, v487, v496)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L319
	}
L236:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v878+int32(5248))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v847)+104)) = int64(-1)
	v889 = F_SlruScanDirectory(m, int32(4383860), int32(294), v847+int32(104))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	goto L235
L240:
	;
	v891 = int32(1)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v847)+104))
	v894 = v892 << (uint(int32(11)) % 32)
	if base.Ui32(v894) <= base.Ui32(v891) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v897 = v891
	goto L243
L242:
	;
	v897 = v894
	goto L243
L243:
	;
	if v867-v897 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v902+int32(5248))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v867 == v866 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L235
L248:
	;
	if v489 != v866 {
		goto L261
	} else {
		goto L262
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+116)) = v865
	goto L248
L250:
	;
	goto L251
L251:
	;
	v911 = F_find_multixact_start(m, v867, v847+int32(116))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	if v911 != 0 {
		goto L248
	} else {
		goto L253
	}
L253:
	;
	v915 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	if v915 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+100)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v847)+96)) = v867
	F_errmsg(m, int32(264681), v847+int32(96))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v930+int32(5248))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	F_errfinish(m, int32(491686), int32(3259), int32(111597))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	goto L235
L261:
	;
	v938 = F_find_multixact_start(m, v489, v847+int32(120))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	v964 = v865
	goto L263
L263:
	;
	if v964 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L264:
	;
	if v938 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v944 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v847)+120))
	v964 = v963
	goto L263
L268:
	;
	if v944 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+80)) = v489
	F_errmsg(m, int32(264841), v847+int32(80))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v958+int32(5248))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L274
	}
L272:
	;
	F_errfinish(m, int32(491686), int32(3277), int32(111597))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	goto L235
L275:
	;
	v969 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v988 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L285
	}
L278:
	;
	if v969 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847))) = v489
	F_errmsg(m, int32(264755), v847)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v981+int32(5248))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	F_errfinish(m, int32(491686), int32(3294), int32(111597))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	goto L235
L285:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v847)+116))
	if v988 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1041 = int32(4483812)
	v1043 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v1044 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1043 + v1044
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+120)) = v1049 | v1044
	*(*int32)(unsafe.Add(mBase, uint32(v847)+140)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v847)+136)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v847)+132)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v847)+128)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v847)+124)) = v495
	F_XLogBeginInsert(m)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L292
	}
L287:
	;
	v993 = int32(1636)
	v994 = base.I32_div_u_s(v990, v993)
	v995 = int32(5)
	v998 = base.I32_div_u_s(v964, v993)
	v1000 = int32(base.Ui32(v998) >> (uint(v995) % 32))
	v1038 = v1000
	v1039 = int32(base.Ui32(v994) >> (uint(v995) % 32))
	v1040 = base.I64_extend_i32_u(v1000)
	goto L286
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+60)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v847)+56)) = v990
	v1004 = int32(1636)
	v1005 = base.I32_div_u_s(v964, v1004)
	v1006 = int32(5)
	v1007 = int32(base.Ui32(v1005) >> (uint(v1006) % 32))
	v1008 = base.I64_extend_i32_u(v1007)
	*(*int64)(unsafe.Add(mBase, uint32(v847)+72)) = v1008
	v1013 = base.I32_div_u_s(v990, v1004)
	v1015 = int32(base.Ui32(v1013) >> (uint(v1006) % 32))
	*(*int64)(unsafe.Add(mBase, uint32(v847-int32(-64)))) = base.I64_extend_i32_u(v1015)
	*(*int32)(unsafe.Add(mBase, uint32(v847)+36)) = v489
	v1019 = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v847)+48)) = base.I64_extend_i32_u(int32(base.Ui32(v489) >> (uint(v1019) % 32)))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+32)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v847)+40)) = base.I64_extend_i32_u(int32(base.Ui32(v867) >> (uint(v1019) % 32)))
	F_errmsg_internal(m, int32(653502), v847+int32(32))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(491686), int32(3307), int32(111597))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v1038 = v1007
	v1039 = v1015
	v1040 = v1008
	goto L286
L292:
	;
	F_XLogRegisterData(m, v847+int32(124), int32(20))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1067 = F_XLogInsert(m, int32(6), int32(48))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	F_XLogFlush(m, v1067)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v1076 = F_LWLockAcquire(m, v1072+int32(1664), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+16)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+12)) = v489
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v1083+int32(1664))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	if v1038 != v1039 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1106 = base.I64_extend_i32_u(v1039)
	goto L301
L299:
	;
	goto L300
L300:
	;
	v1154 = int32(1)
	if v489 == v1154 {
		goto L314
	} else {
		goto L315
	}
L301:
	;
	v1111 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	goto L300
L303:
	;
	if v1111 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v847)+16)) = v1106
	F_errmsg_internal(m, int32(27146), v847+int32(16))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	F_SlruDeleteSegment(m, v1106)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L309
	}
L307:
	;
	F_errfinish(m, int32(491686), int32(3132), int32(264951))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	if v1106 != int64(82040) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1131 = v1106 + int64(1)
	goto L312
L311:
	;
	v1131 = int64(0)
	goto L312
L312:
	;
	if v1131 != v1040 {
		v1106 = v1131
		goto L301
	} else {
		goto L313
	}
L313:
	;
	goto L302
L314:
	;
	v1160 = int32(2097151)
	goto L316
L315:
	;
	v1160 = int32(base.Ui32(v489-v1154) >> (uint(int32(11)) % 32))
	goto L316
L316:
	;
	F_SimpleLruTruncate(m, int32(4383860), base.I64_extend_i32_u(v1160))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1165)+120)) = v1166 & int32(-2)
	v1170 = int32(4483812)
	v1172 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1172 - int32(1)
	v1177 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v1177+int32(5248))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	goto L235
L319:
	;
	F_SetMultiXactIdLimit(m, v489, v495, int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	goto L154
L321:
	;
	goto L47
}
func F_vac_update_relstats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 float32
	_ = v49
	var v50 float32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v19-int32(-64), int32(1), int32(3), int32(184), v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_systable_inplace_update_begin(m, v24, int32(2662), v19-int32(-64), v19+int32(60), v19+int32(56))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	if l1 != v46 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L85
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+96)) = l1
	goto L10
L9:
	;
	goto L10
L10:
	;
	v49 = base.F32_demote_f64(l2)
	v50 = *(*float32)(unsafe.Add(mBase, uint32(v45)+100))
	if base.F32_eq(v49, v50) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = base.B2i32(l1 != v46)
	goto L13
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v45)+100)) = v49
	v55 = int32(1)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	if l3 != v56 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+104)) = l3
	v60 = int32(1)
	goto L16
L15:
	;
	v60 = v55
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+108))
	if l4 != v61 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+108)) = l4
	v65 = int32(1)
	goto L19
L18:
	;
	v65 = v60
	goto L19
L19:
	;
	if l10 != 0 {
		v90 = v65
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v45)+136))
	if l8 != 0 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	if l5 != 0 {
		v74 = v65
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+124)))
	if v75 != int32(1) {
		v82 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+116)))
	if v66&int32(1) == int32(0) {
		v74 = v65
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+116)) = uint8(v71)
	v74 = int32(1)
	goto L22
L25:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+125)))
	if v83 != int32(1) {
		v90 = v82
		goto L20
	} else {
		goto L28
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v78 != 0 {
		v82 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+124)) = uint8(v79)
	v82 = int32(1)
	goto L25
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v86 != 0 {
		v90 = v82
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+125)) = uint8(v87)
	v90 = int32(1)
	goto L20
L30:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v92)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v94 = int32(0)
	if base.Ui32(l6) < base.Ui32(int32(3)) {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v45)+140))
	if l9 != 0 {
		goto L50
	} else {
		goto L51
	}
L34:
	;
	if v91 == l6 {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v91)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v109 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v109 = base.B2i32(base.Ui32(v91) < base.Ui32(l6))
	goto L36
L38:
	;
	goto L39
L39:
	;
	v109 = int32(base.Ui32(v91-l6) >> (uint(int32(31)) % 32))
	goto L36
L40:
	;
	v112 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+136)) = l6
	v130 = int32(1)
	v132 = v109 ^ v130
	if l8 == int32(0) {
		v137 = v130
		v138 = v132
		goto L33
	} else {
		goto L49
	}
L43:
	;
	v114 = base.I32_wrap_i64(v112)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v114)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v126 == int32(0) {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L48
	}
L45:
	;
	v126 = base.B2i32(base.Ui32(v114) < base.Ui32(v91))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v126 = int32(base.Ui32(v114-v91) >> (uint(int32(31)) % 32))
	goto L44
L48:
	;
	goto L42
L49:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v135)
	v137 = v130
	v138 = v132
	goto L33
L50:
	;
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v141)
	goto L52
L51:
	;
	goto L52
L52:
	;
	if l7 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	F_sequence_close(m, v24, int32(3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L70
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_systable_inplace_update_cancel(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L69
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_systable_inplace_update_finish(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L68
	}
L56:
	;
	v166 = int32(0)
	if v137 == v166 {
		goto L54
	} else {
		goto L67
	}
L57:
	;
	if l7 == v140 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v148 = int32(base.Ui32(v140-l7) >> (uint(int32(31)) % 32))
	goto L59
L59:
	;
	if v148 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v151 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+140)) = l7
	v160 = v148 ^ int32(1)
	if l9 == int32(0) {
		v170 = v160
		goto L55
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if int32(base.Ui32(v151-v140)>>(uint(int32(31))%32)) == int32(0) {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v163)
	v170 = v160
	goto L55
L67:
	;
	v170 = v166
	goto L55
L68:
	;
	v179 = v170
	goto L53
L69:
	;
	v179 = v166
	goto L53
L70:
	;
	if v138 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v179 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v187 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v187 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v194 + int32(4)
	F_errmsg_internal(m, int32(700674), v19+int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(495260), int32(1595), int32(125352))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L71
L78:
	;
	m.G0 = v19 + int32(112)
	return
L79:
	;
	v215 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v215 == int32(0) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v222 + int32(4)
	F_errmsg_internal(m, int32(700747), v19+int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(495260), int32(1601), int32(125352))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L78
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	F_errmsg_internal(m, int32(334280), v19)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(495260), int32(1474), int32(125352))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varbit_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v19 = F_palloc(m, v16+int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = v16 - int32(8)
			if v22 < int32(0) {
				v89 = v19
				v91 = v15
				v92 = v2
			} else {
				v25 = v19
				v27 = v15
				v28 = v2
				for {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					v36 = int32(48)
					v37 = v33&int32(1) | v36
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)) = uint8(v37)
					if v33&int32(2) != 0 {
						v43 = int32(49)
					} else {
						v43 = v36
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v43)
					if v33&int32(4) != 0 {
						v49 = int32(49)
					} else {
						v49 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v49)
					if v33&int32(8) != 0 {
						v55 = int32(49)
					} else {
						v55 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)) = uint8(v55)
					if v33&int32(16) != 0 {
						v61 = int32(49)
					} else {
						v61 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)) = uint8(v61)
					if v33&int32(32) != 0 {
						v67 = int32(49)
					} else {
						v67 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)) = uint8(v67)
					if v33&int32(64) != 0 {
						v73 = int32(49)
					} else {
						v73 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v73)
					if int32(0) <= base.I32_extend8_s(v33) {
						v80 = int32(48)
					} else {
						v80 = int32(49)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v80)
					v83 = v27 + int32(1)
					v84 = int32(8)
					v85 = v25 + v84
					v87 = v28 + v84
					if v87 <= v22 {
						v25 = v85
						v27 = v83
						v28 = v87
						continue
					} else {
						break
					}
					break
				}
				v89 = v85
				v91 = v83
				v92 = v87
			}
			if v16 <= v92 {
				v180 = v89
			} else {
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v101 = (v16 - v92) & int32(3)
				if v101 == int32(0) {
					v129 = v89
					v130 = v98
					v131 = v92
				} else {
					v105 = v89
					v106 = v98
					v107 = v92
					v110 = int32(0)
					for {
						if int32(0) <= base.I32_extend8_s(v106) {
							v118 = int32(48)
						} else {
							v118 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v118)
						v120 = int32(1)
						v121 = v107 + v120
						v123 = v106 << (uint(v120) % 32)
						v125 = v105 + v120
						v127 = v110 + v120
						if v127 != v101 {
							v105 = v125
							v106 = v123
							v107 = v121
							v110 = v127
							continue
						} else {
							break
						}
						break
					}
					v129 = v125
					v130 = v123
					v131 = v121
				}
				if base.Ui32(int32(-4)) < base.Ui32(v92-v16) {
					v180 = v129
				} else {
					v140 = v129
					v141 = v130
					v142 = v131
					for {
						if v141&int32(16) != 0 {
							v152 = int32(49)
						} else {
							v152 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+3)) = uint8(v152)
						if v141&int32(32) != 0 {
							v158 = int32(49)
						} else {
							v158 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)) = uint8(v158)
						if v141&int32(64) != 0 {
							v164 = int32(49)
						} else {
							v164 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)) = uint8(v164)
						if int32(0) <= base.I32_extend8_s(v141) {
							v171 = int32(48)
						} else {
							v171 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v171)
						v173 = int32(4)
						v176 = v140 + v173
						v178 = v142 + v173
						if v178 != v16 {
							v140 = v176
							v141 = v141 << (uint(v173) % 32)
							v142 = v178
							continue
						} else {
							break
						}
						break
					}
					v180 = v176
				}
			}
			v188 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v188)
			return v19
		}
	}
}
func F_varbit_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if base.Ui32(v14) < base.Ui32(int32(2147483641)) {
			if base.B2i32(v11 < v14)&base.B2i32(int32(0) < v11) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16777346))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg(m, int32(661645), v9)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(491490), int32(662), int32(36253))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
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
				v27 = int32(base.Ui32(v14+int32(7)) >> (uint(int32(3)) % 32))
				v29 = v27 + int32(8)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v29 << (uint(int32(2)) % 32)
					F_pq_copymsgbytes(m, v12, v30+int32(8), v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						v48 = v42<<(uint(int32(3))%32) - v45 + int32(-64)
						if int32(0) < v48 {
							v53 = v30 + v42 - int32(1)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
							v57 = v54 & (int32(255) << (uint(v48) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
						} else {
						}
						m.G0 = v9 + int32(16)
						return v30
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(329337), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491490), int32(652), int32(36253))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
func F_varchar_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v36 = v2
		return v36
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v36 = v2
			return v36
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v36 = v2
				return v36
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if int32(0) <= v22 {
						if v18 < int32(0) {
							v36 = v2
							return v36
						} else {
							v27 = int32(4)
							if v22-v27 < v18-v27 {
								v36 = v2
								return v36
							} else {
								v32 = F_relabel_to_typmod(m, v17, v22)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v36 = v32
									return v36
								}
							}
						}
					} else {
						v32 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v36 = v32
							return v36
						}
					}
				}
			}
		}
	}
}
func F_varstr_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
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
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
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
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
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
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
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
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = F_pg_detoast_datum_packed(m, l0)
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v24 = int32(1)
	v25 = v15 + v24
	v27 = v19 & v24
	if v19 == v24 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v32&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v19)>>(uint(v45)%32)) - v45
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v41 = v30
	goto L9
L8:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L9
L9:
	;
	if v32 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = v30
	goto L12
L11:
	;
	v44 = v41
	goto L12
L12:
	;
	v55 = v44
	goto L3
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v56 = v25
	goto L16
L15:
	;
	v56 = v15 + int32(4)
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v57 == int32(1042) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = int32(-1)
	v64 = v55 - int32(1)
	if v62 <= v64 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v83 = v55
	goto L19
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+29)))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v83 = v82
	goto L19
L21:
	;
	v67 = v62
	goto L23
L22:
	;
	v67 = v64
	goto L23
L23:
	;
	v71 = v55
	goto L24
L24:
	;
	v75 = v71 - int32(1)
	if v75 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v82 = v71
	goto L20
L26:
	;
	v82 = v67 + int32(1)
	goto L20
L27:
	;
	goto L28
L28:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v75))))
	if v79 == int32(32) {
		v71 = v75
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if l0 != v15 {
		goto L181
	} else {
		goto L182
	}
L31:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v688 = int32(4)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if base.Ui32(v688) <= base.Ui32(v689) {
		goto L174
	} else {
		goto L175
	}
L32:
	;
	v268 = v83
	v275 = v56
	goto L34
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v87 <= v83 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v276 = int32(4)
	if base.Ui32(v276) <= base.Ui32(v268) {
		goto L99
	} else {
		goto L100
	}
L35:
	;
	v89 = int32(1)
	v90 = v83 + v89
	v91 = int32(1073741823)
	v93 = v87 << (uint(v89) % 32)
	if base.Ui32(v91) <= base.Ui32(v93) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v83 != v106 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v96 = v91
	goto L40
L39:
	;
	v96 = v93
	goto L40
L40:
	;
	if base.Ui32(v96) < base.Ui32(v90) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v98 = v90
	goto L43
L42:
	;
	v98 = v96
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v101 = F_repalloc(m, v100, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v101
	goto L37
L45:
	;
	if v83 != 0 {
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v177 = v108
	goto L45
L47:
	;
	goto L48
L48:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	if v110 != int32(1) {
		v177 = v109
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v83) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v174 == int32(0) {
		goto L31
	} else {
		goto L68
	}
L51:
	;
	v174 = int32(0)
	goto L50
L52:
	;
	v148 = v143
	v149 = v144
	v150 = v145
	goto L62
L53:
	;
	if (v109|v56)&int32(3) != 0 {
		v143 = v109
		v144 = v56
		v145 = v83
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v136 = v109
	v137 = v56
	v138 = v83
	goto L55
L55:
	;
	if v138 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v120 = v109
	v121 = v56
	v122 = v83
	goto L57
L57:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != v126 {
		v143 = v120
		v144 = v121
		v145 = v122
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v136 = v131
	v137 = v129
	v138 = v133
	goto L55
L59:
	;
	v128 = int32(4)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v143 = v136
	v144 = v137
	v145 = v138
	goto L52
L62:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 == v154 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v174 = v153 - v154
	goto L50
L64:
	;
	v156 = int32(1)
	v161 = v150 - v156
	if v161 != 0 {
		v148 = v148 + v156
		v149 = v149 + v156
		v150 = v161
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L51
L68:
	;
	v177 = v109
	goto L45
L69:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180+v83))) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v83
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v187 == v182 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v177, v56, v83)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v268 = v258
	v275 = v265
	goto L34
L74:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v194 = F_pg_strxfrm(m, v190, v191, v192, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v231) {
		goto L91
	} else {
		goto L92
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v194) < base.Ui32(v197) {
		v258 = v194
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v201 = v194
	v203 = v197
	goto L79
L79:
	;
	v208 = int32(1)
	v209 = v201 + v208
	v210 = int32(1073741823)
	v212 = v203 << (uint(v208) % 32)
	if base.Ui32(v210) <= base.Ui32(v212) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v258 = v226
	goto L73
L81:
	;
	v215 = v210
	goto L83
L82:
	;
	v215 = v212
	goto L83
L83:
	;
	if base.Ui32(v215) < base.Ui32(v209) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v217 = v209
	goto L86
L85:
	;
	v217 = v215
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v220 = F_repalloc(m, v219, v217)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v220
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v226 = F_pg_strxfrm(m, v220, v223, v224, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v226
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v229) <= base.Ui32(v226) {
		v201 = v226
		v203 = v229
		goto L79
	} else {
		goto L89
	}
L89:
	;
	goto L80
L90:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = m.T0[v252].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v246, int32(4), v248, int32(-1), v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L98
	}
L91:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v246 = v234
	goto L90
L92:
	;
	goto L93
L93:
	;
	v235 = int32(2)
	if base.Ui32(v231) <= base.Ui32(v235) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v238 = v235
	goto L96
L95:
	;
	v238 = v231
	goto L96
L96:
	;
	v240 = v238 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v243 = F_repalloc(m, v242, v240)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v243
	v246 = v243
	goto L90
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v253
	v258 = v253
	goto L73
L99:
	;
	v279 = v276
	goto L101
L100:
	;
	v279 = v268
	goto L101
L101:
	;
	if v279 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v282 = int32(128)
	if v282 <= v83 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v280 = F__emscripten_memcpy_bulkmem(m, v12+int32(12), v275, v279)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v285 = v282
	goto L108
L107:
	;
	v285 = v83
	goto L108
L108:
	;
	v291 = v285 - int32(1636608432)
	if v56&int32(3) != 0 {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v551 = v14 - int32(-64)
	if int32(129) <= v83 {
		goto L149
	} else {
		goto L150
	}
L110:
	;
	v523 = int32(14)
	v525 = v519 ^ v520 - base.I32_rotl(v519, v523)
	v529 = v525 ^ v518 - base.I32_rotl(v525, int32(11))
	v533 = v529 ^ v519 - base.I32_rotl(v529, int32(25))
	v537 = v533 ^ v525 - base.I32_rotl(v533, int32(16))
	v541 = v537 ^ v529 - base.I32_rotl(v537, int32(4))
	v545 = v541 ^ v533 - base.I32_rotl(v541, v523)
	v549 = v545 ^ v537 - base.I32_rotl(v545, int32(24))
	goto L109
L111:
	;
	switch v449 - int32(1) {
	case 0:
		v511 = v450
		v512 = v451
		v513 = v452
		goto L138
	case 1:
		v504 = v450
		v505 = v451
		v506 = v452
		goto L139
	case 2:
		v497 = v450
		v498 = v451
		v499 = v452
		goto L140
	case 3:
		v491 = v451
		v492 = v452
		goto L141
	case 4:
		v487 = v451
		v488 = v452
		goto L142
	case 5:
		v481 = v451
		v482 = v452
		goto L143
	case 6:
		v475 = v451
		v476 = v452
		goto L144
	case 7:
		v470 = v452
		goto L145
	case 8:
		v465 = v452
		goto L146
	case 9:
		v460 = v452
		goto L147
	case 10:
		goto L148
	default:
		v518 = v450
		v519 = v451
		v520 = v452
		goto L110
	}
L112:
	;
	v400 = v56
	v401 = v285
	v402 = v291
	v403 = v291
	v404 = v291
	goto L135
L113:
	;
	if base.Ui32(int32(11)) < base.Ui32(v285) {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v285) < base.Ui32(int32(12)) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v448 = v56
	v449 = v285
	v450 = v291
	v451 = v291
	v452 = v291
	goto L111
L117:
	;
	switch v347 - int32(1) {
	case 0:
		v397 = v348
		goto L124
	case 1:
		v392 = v348
		goto L125
	case 2:
		goto L126
	case 3:
		v385 = v349
		goto L127
	case 4:
		v382 = v349
		goto L128
	case 5:
		v377 = v349
		goto L129
	case 6:
		goto L130
	case 7:
		v368 = v350
		goto L131
	case 8:
		v363 = v350
		goto L132
	case 9:
		v358 = v350
		goto L133
	case 10:
		goto L134
	default:
		v518 = v348
		v519 = v349
		v520 = v350
		goto L110
	}
L118:
	;
	v346 = v56
	v347 = v285
	v348 = v291
	v349 = v291
	v350 = v291
	goto L117
L119:
	;
	goto L120
L120:
	;
	v298 = v56
	v299 = v285
	v300 = v291
	v301 = v291
	v302 = v291
	goto L121
L121:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v305 = v304 + v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
	v309 = v308 + v302
	v311 = int32(4)
	v313 = v306 + v300 - v309 ^ base.I32_rotl(v309, v311)
	v317 = v305 - v313 ^ base.I32_rotl(v313, int32(6))
	v318 = v309 + v305
	v319 = v313 + v318
	v320 = v317 + v319
	v324 = v318 - v317 ^ base.I32_rotl(v317, int32(8))
	v328 = v319 - v324 ^ base.I32_rotl(v324, int32(16))
	v332 = v320 - v328 ^ base.I32_rotl(v328, int32(19))
	v333 = v324 + v320
	v334 = v328 + v333
	v335 = v332 + v334
	v339 = v333 - v332 ^ base.I32_rotl(v332, v311)
	v340 = int32(12)
	v341 = v298 + v340
	v343 = v299 - v340
	if base.Ui32(int32(11)) < base.Ui32(v343) {
		v298 = v341
		v299 = v343
		v300 = v334
		v301 = v335
		v302 = v339
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v346 = v341
	v347 = v343
	v348 = v334
	v349 = v335
	v350 = v339
	goto L117
L123:
	;
	goto L122
L124:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v518 = v397 + v398
	v519 = v349
	v520 = v350
	goto L110
L125:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	v397 = v393<<(uint(int32(8))%32) + v392
	goto L124
L126:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+2)))
	v392 = v388<<(uint(int32(16))%32) + v348
	goto L125
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v518 = v386 + v348
	v519 = v385
	v520 = v350
	goto L110
L128:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+4)))
	v385 = v382 + v383
	goto L127
L129:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+5)))
	v382 = v378<<(uint(int32(8))%32) + v377
	goto L128
L130:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+6)))
	v377 = v373<<(uint(int32(16))%32) + v349
	goto L129
L131:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v518 = v369 + v348
	v519 = v371 + v349
	v520 = v368
	goto L110
L132:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+8)))
	v368 = v364<<(uint(int32(8))%32) + v363
	goto L131
L133:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+9)))
	v363 = v359<<(uint(int32(16))%32) + v358
	goto L132
L134:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+10)))
	v358 = v354<<(uint(int32(24))%32) + v350
	goto L133
L135:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v407 = v406 + v403
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	v411 = v410 + v404
	v413 = int32(4)
	v415 = v408 + v402 - v411 ^ base.I32_rotl(v411, v413)
	v419 = v407 - v415 ^ base.I32_rotl(v415, int32(6))
	v420 = v411 + v407
	v421 = v415 + v420
	v422 = v419 + v421
	v426 = v420 - v419 ^ base.I32_rotl(v419, int32(8))
	v430 = v421 - v426 ^ base.I32_rotl(v426, int32(16))
	v434 = v422 - v430 ^ base.I32_rotl(v430, int32(19))
	v435 = v426 + v422
	v436 = v430 + v435
	v437 = v434 + v436
	v441 = v435 - v434 ^ base.I32_rotl(v434, v413)
	v442 = int32(12)
	v443 = v400 + v442
	v445 = v401 - v442
	if base.Ui32(int32(11)) < base.Ui32(v445) {
		v400 = v443
		v401 = v445
		v402 = v436
		v403 = v437
		v404 = v441
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v448 = v443
	v449 = v445
	v450 = v436
	v451 = v437
	v452 = v441
	goto L111
L137:
	;
	goto L136
L138:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v518 = v511 + v514
	v519 = v512
	v520 = v513
	goto L110
L139:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+1)))
	v511 = v507<<(uint(int32(8))%32) + v504
	v512 = v505
	v513 = v506
	goto L138
L140:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+2)))
	v504 = v500<<(uint(int32(16))%32) + v497
	v505 = v498
	v506 = v499
	goto L139
L141:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+3)))
	v497 = v493<<(uint(int32(24))%32) + v450
	v498 = v491
	v499 = v492
	goto L140
L142:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+4)))
	v491 = v487 + v489
	v492 = v488
	goto L141
L143:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+5)))
	v487 = v483<<(uint(int32(8))%32) + v481
	v488 = v482
	goto L142
L144:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+6)))
	v481 = v477<<(uint(int32(16))%32) + v475
	v482 = v476
	goto L143
L145:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+7)))
	v475 = v471<<(uint(int32(24))%32) + v451
	v476 = v470
	goto L144
L146:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)))
	v470 = v466<<(uint(int32(8))%32) + v465
	goto L145
L147:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+9)))
	v465 = v461<<(uint(int32(16))%32) + v460
	goto L146
L148:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+10)))
	v460 = v456<<(uint(int32(24))%32) + v452
	goto L147
L149:
	;
	v558 = int32(711645284)
	v561 = v83 - int32(1636608428) ^ v558 - int32(1455628627)
	v566 = v561 ^ int32(-1636608428) - base.I32_rotl(v561, int32(25))
	v571 = v566 ^ v558 - base.I32_rotl(v566, int32(16))
	v575 = v571 ^ v561 - base.I32_rotl(v571, int32(4))
	v579 = v575 ^ v566 - base.I32_rotl(v575, int32(14))
	goto L152
L150:
	;
	v585 = v549
	goto L151
L151:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v590 = int32(32) - v589
	v591 = v585 << (uint(v589) % 32)
	if v591 != 0 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v585 = v579 ^ v571 - base.I32_rotl(v579, int32(24)) ^ v549
	goto L151
L153:
	;
	v619 = v14 + int32(40)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v625 = int32(711645284)
	v628 = v620 - int32(1636608428) ^ v625 - int32(1455628627)
	v633 = v628 ^ int32(-1636608428) - base.I32_rotl(v628, int32(25))
	v638 = v633 ^ v625 - base.I32_rotl(v633, int32(16))
	v642 = v638 ^ v628 - base.I32_rotl(v638, int32(4))
	v646 = v642 ^ v633 - base.I32_rotl(v642, int32(14))
	v650 = v646 ^ v638 - base.I32_rotl(v646, int32(24))
	goto L163
L154:
	;
	v598 = int32(32) - (base.I32_clz(v591) ^ int32(31))
	v599 = int32(255)
	if base.Ui32(v590&v599) < base.Ui32(v598&v599) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v608 = v590 + int32(1)
	goto L156
L156:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v551)+16))
	v611 = v609 + int32(base.Ui32(v585)>>(uint(v590)%32))
	v613 = v608 & int32(255)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if base.Ui32(v614) < base.Ui32(v613) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v604 = v590 + int32(1)
	goto L159
L158:
	;
	v604 = v598
	goto L159
L159:
	;
	v608 = v604
	goto L156
L160:
	;
	v616 = v613
	goto L162
L161:
	;
	v616 = v614
	goto L162
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v616)
	goto L153
L163:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v655 = int32(32) - v654
	v656 = v650 << (uint(v654) % 32)
	if v656 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v683 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v683)
	v698 = v620
	goto L30
L165:
	;
	v663 = int32(32) - (base.I32_clz(v656) ^ int32(31))
	v664 = int32(255)
	if base.Ui32(v655&v664) < base.Ui32(v663&v664) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v673 = v655 + int32(1)
	goto L167
L167:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	v676 = v674 + int32(base.Ui32(v650)>>(uint(v655)%32))
	v678 = v673 & int32(255)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if base.Ui32(v679) < base.Ui32(v678) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v669 = v655 + int32(1)
	goto L170
L169:
	;
	v669 = v663
	goto L170
L170:
	;
	v673 = v669
	goto L167
L171:
	;
	v681 = v678
	goto L173
L172:
	;
	v681 = v679
	goto L173
L173:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v681)
	goto L164
L174:
	;
	v692 = v688
	goto L176
L175:
	;
	v692 = v689
	goto L176
L176:
	;
	if v692 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v698 = v695
	goto L30
L178:
	;
	v693 = F__emscripten_memcpy_bulkmem(m, v12+int32(12), v687, v692)
	mBase = m.M
	goto L180
L179:
	;
	goto L180
L180:
	;
	goto L177
L181:
	;
	F_pfree(m, v15)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v12 + int32(16)
	v711 = int32(24)
	v713 = int32(65280)
	v715 = int32(8)
	return v698<<(uint(v711)%32) | v698&v713<<(uint(v715)%32) | (int32(base.Ui32(v698)>>(uint(v715)%32))&v713 | int32(base.Ui32(v698)>>(uint(v711)%32)))
L184:
	;
	goto L183
}
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
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
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = F_pg_newlocale_from_collation(m, l4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L80
	}
L4:
	;
	return v219
L5:
	;
	return int32(0)
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
	if v11 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if l1 != l3 {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v15 = l1
	goto L12
L11:
	;
	v15 = l3
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v15) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v77 != 0 {
		v219 = v77
		goto L4
	} else {
		goto L31
	}
L14:
	;
	v77 = int32(0)
	goto L13
L15:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L25
L16:
	;
	if (l0|l2)&int32(3) != 0 {
		v46 = l0
		v47 = l2
		v48 = v15
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v39 = l0
	v40 = l2
	v41 = v15
	goto L18
L18:
	;
	if v41 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v23 = l0
	v24 = l2
	v25 = v15
	goto L20
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L18
L22:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L15
L25:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v77 = v56 - v57
	goto L13
L27:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	return base.B2i32(l3 < l1) - v14
L32:
	;
	v146 = F_pg_strncoll(m, l0, l1, l2, l3, v7)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L53
	}
L33:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v143 != 0 {
		goto L32
	} else {
		goto L52
	}
L35:
	;
	v143 = int32(0)
	goto L34
L36:
	;
	v117 = v112
	v118 = v113
	v119 = v114
	goto L46
L37:
	;
	if (l0|l2)&int32(3) != 0 {
		v112 = l0
		v113 = l2
		v114 = l1
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v105 = l0
	v106 = l2
	v107 = l1
	goto L39
L39:
	;
	if v107 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v89 = l0
	v90 = l2
	v91 = l1
	goto L41
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != v95 {
		v112 = v89
		v113 = v90
		v114 = v91
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v105 = v100
	v106 = v98
	v107 = v102
	goto L39
L43:
	;
	v97 = int32(4)
	v98 = v90 + v97
	v100 = v89 + v97
	v102 = v91 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v89 = v100
		v90 = v98
		v91 = v102
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v112 = v105
	v113 = v106
	v114 = v107
	goto L36
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 == v123 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v143 = v122 - v123
	goto L34
L48:
	;
	v125 = int32(1)
	v130 = v119 - v125
	if v130 != 0 {
		v117 = v117 + v125
		v118 = v118 + v125
		v119 = v130
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	return int32(0)
L53:
	;
	if v146 != 0 {
		v219 = v146
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v148 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	goto L57
L57:
	;
	v153 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = l1
	goto L60
L59:
	;
	v154 = l3
	goto L60
L60:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v154) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v216 != 0 {
		v219 = v216
		goto L4
	} else {
		goto L79
	}
L62:
	;
	v216 = int32(0)
	goto L61
L63:
	;
	v190 = v185
	v191 = v186
	v192 = v187
	goto L73
L64:
	;
	if (l0|l2)&int32(3) != 0 {
		v185 = l0
		v186 = l2
		v187 = v154
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v178 = l0
	v179 = l2
	v180 = v154
	goto L66
L66:
	;
	if v180 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v162 = l0
	v163 = l2
	v164 = v154
	goto L68
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v167 != v168 {
		v185 = v162
		v186 = v163
		v187 = v164
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v178 = v173
	v179 = v171
	v180 = v175
	goto L66
L70:
	;
	v170 = int32(4)
	v171 = v163 + v170
	v173 = v162 + v170
	v175 = v164 - v170
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v162 = v173
		v163 = v171
		v164 = v175
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L63
L73:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 == v196 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v216 = v195 - v196
	goto L61
L75:
	;
	v198 = int32(1)
	v203 = v192 - v198
	if v203 != 0 {
		v190 = v190 + v198
		v191 = v191 + v198
		v192 = v203
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	v219 = base.B2i32(l3 < l1) - v153
	goto L4
L80:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(244851), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(557901), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(498299), int32(1648), int32(105964))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varstrfastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
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
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if l1 != l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v74 == int32(1042) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v71 = int32(0)
	goto L3
L5:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L15
L6:
	;
	if (l0|l2)&int32(3) != 0 {
		v40 = l0
		v41 = l2
		v42 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v33 = l0
	v34 = l2
	v35 = l1
	goto L8
L8:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v17 = l0
	v18 = l2
	v19 = l1
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L8
L12:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L5
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v71 = v50 - v51
	goto L3
L17:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	return int32(0)
L22:
	;
	v79 = int32(-1)
	v81 = l1 - int32(1)
	if v79 <= v81 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v123 = l1
	v124 = l3
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v125 <= v123 {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	v102 = int32(-1)
	v104 = l3 - int32(1)
	if v102 <= v104 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v84 = v79
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v88 = l1
	goto L29
L29:
	;
	v92 = v88 - int32(1)
	if v92 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v99 = v88
	goto L25
L31:
	;
	v99 = v84 + int32(1)
	goto L25
L32:
	;
	goto L33
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92))))
	if v96 == int32(32) {
		v88 = v92
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v123 = v99
	v124 = v122
	goto L24
L36:
	;
	v107 = v102
	goto L38
L37:
	;
	v107 = v104
	goto L38
L38:
	;
	v111 = l3
	goto L39
L39:
	;
	v115 = v111 - int32(1)
	if v115 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v122 = v111
	goto L35
L41:
	;
	v122 = v107 + int32(1)
	goto L35
L42:
	;
	goto L43
L43:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v115))))
	if v119 == int32(32) {
		v111 = v115
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v127 = int32(1)
	v128 = v123 + v127
	v129 = int32(1073741823)
	v131 = v125 << (uint(v127) % 32)
	if base.Ui32(v129) <= base.Ui32(v131) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v146 <= v124 {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v134 = v129
	goto L50
L49:
	;
	v134 = v131
	goto L50
L50:
	;
	if base.Ui32(v134) < base.Ui32(v128) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v136 = v128
	goto L53
L52:
	;
	v136 = v134
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v139 = F_repalloc(m, v138, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v139
	goto L47
L56:
	;
	v148 = int32(1)
	v149 = v124 + v148
	v150 = int32(1073741823)
	v152 = v146 << (uint(v148) % 32)
	if base.Ui32(v150) <= base.Ui32(v152) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v123 != v166 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v155 = v150
	goto L61
L60:
	;
	v155 = v152
	goto L61
L61:
	;
	if base.Ui32(v155) < base.Ui32(v149) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v157 = v149
	goto L64
L63:
	;
	v157 = v155
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v160 = F_repalloc(m, v159, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v160
	goto L58
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v241 == v124 {
		goto L94
	} else {
		goto L95
	}
L67:
	;
	if v123 != 0 {
		goto L89
	} else {
		goto L90
	}
L68:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v123) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v229 != 0 {
		goto L67
	} else {
		goto L87
	}
L70:
	;
	v229 = int32(0)
	goto L69
L71:
	;
	v203 = v198
	v204 = v199
	v205 = v200
	goto L81
L72:
	;
	if (v165|l0)&int32(3) != 0 {
		v198 = v165
		v199 = l0
		v200 = v123
		goto L71
	} else {
		goto L75
	}
L73:
	;
	v191 = v165
	v192 = l0
	v193 = v123
	goto L74
L74:
	;
	if v193 == int32(0) {
		goto L70
	} else {
		goto L80
	}
L75:
	;
	v175 = v165
	v176 = l0
	v177 = v123
	goto L76
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v180 != v181 {
		v198 = v175
		v199 = v176
		v200 = v177
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v191 = v186
	v192 = v184
	v193 = v188
	goto L74
L78:
	;
	v183 = int32(4)
	v184 = v176 + v183
	v186 = v175 + v183
	v188 = v177 - v183
	if base.Ui32(int32(3)) < base.Ui32(v188) {
		v175 = v186
		v176 = v184
		v177 = v188
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v198 = v191
	v199 = v192
	v200 = v193
	goto L71
L81:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v208 == v209 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v229 = v208 - v209
	goto L69
L83:
	;
	v211 = int32(1)
	v216 = v205 - v211
	if v216 != 0 {
		v203 = v203 + v211
		v204 = v204 + v211
		v205 = v216
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L70
L87:
	;
	v239 = int32(1)
	goto L66
L88:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v123))) = uint8(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v123
	v239 = v235
	goto L66
L89:
	;
	v231 = F__emscripten_memcpy_bulkmem(m, v165, l0, v123)
	mBase = m.M
	goto L91
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v322 = int32(-1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = m.T0[v326].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v321, v322, v320, v322, v324)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L54
	} else {
		goto L123
	}
L93:
	;
	if v239 == int32(0) {
		v320 = v240
		goto L92
	} else {
		goto L120
	}
L94:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v124) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	goto L96
L96:
	;
	if v124 != 0 {
		goto L117
	} else {
		goto L118
	}
L97:
	;
	if v304 == int32(0) {
		goto L93
	} else {
		goto L115
	}
L98:
	;
	v304 = int32(0)
	goto L97
L99:
	;
	v278 = v273
	v279 = v274
	v280 = v275
	goto L109
L100:
	;
	if (v240|l2)&int32(3) != 0 {
		v273 = v240
		v274 = l2
		v275 = v124
		goto L99
	} else {
		goto L103
	}
L101:
	;
	v266 = v240
	v267 = l2
	v268 = v124
	goto L102
L102:
	;
	if v268 == int32(0) {
		goto L98
	} else {
		goto L108
	}
L103:
	;
	v250 = v240
	v251 = l2
	v252 = v124
	goto L104
L104:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v255 != v256 {
		v273 = v250
		v274 = v251
		v275 = v252
		goto L99
	} else {
		goto L106
	}
L105:
	;
	v266 = v261
	v267 = v259
	v268 = v263
	goto L102
L106:
	;
	v258 = int32(4)
	v259 = v251 + v258
	v261 = v250 + v258
	v263 = v252 - v258
	if base.Ui32(int32(3)) < base.Ui32(v263) {
		v250 = v261
		v251 = v259
		v252 = v263
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v273 = v266
	v274 = v267
	v275 = v268
	goto L99
L109:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v283 == v284 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v304 = v283 - v284
	goto L97
L111:
	;
	v286 = int32(1)
	v291 = v280 - v286
	if v291 != 0 {
		v278 = v278 + v286
		v279 = v279 + v286
		v280 = v291
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L98
L115:
	;
	goto L96
L116:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309+v124))) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v124
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v320 = v314
	goto L92
L117:
	;
	v307 = F__emscripten_memcpy_bulkmem(m, v240, l2, v124)
	mBase = m.M
	goto L119
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
	if v317 != 0 {
		v320 = v240
		goto L92
	} else {
		goto L121
	}
L121:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	return v318
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v361
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v363)
	return v361
L123:
	;
	if v327 != 0 {
		v361 = v327
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v331 != int32(1) {
		v361 = int32(0)
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v339 == int32(0) {
		v358 = v338
		v359 = v339
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v361 = v359 - v358
	goto L122
L127:
	;
	goto L126
L128:
	;
	if v338 != v339 {
		v358 = v338
		v359 = v339
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v343 = v334
	v344 = v335
	goto L130
L130:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v358 = v347
		v359 = v348
		goto L127
	} else {
		goto L132
	}
L131:
	;
	v358 = v347
	v359 = v348
	goto L127
L132:
	;
	v351 = int32(1)
	if v347 == v348 {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
}
func F_void_out(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pstrdup(m, int32(740129))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
