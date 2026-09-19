package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HalfvecInnerProductDefault(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 float32
	_ = v6
	var v12 int32
	_ = v12
	var v14 float32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 float32
	_ = v202
	var v204 int32
	_ = v204
	var v211 float32
	_ = v211
	v4 = int32(0)
	v6 = float32(0)
	if v4 < l0 {
		v12 = v4
		v14 = v6
		for {
			v16 = v12 << (uint(int32(1)) % 32)
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v16))))
			v23 = v18 & int32(1023)
			v27 = v18 << (uint(int32(16)) % 32) & int32(-2147483648)
			v30 = int32(31)
			v31 = int32(base.Ui32(v18)>>(uint(int32(10))%32)) & v30
			if v31 != v30 {
				if v31 != 0 {
					v103 = v23
					v104 = v31<<(uint(int32(23))%32) + v27 + int32(939524096)
				} else {
					if v23 != 0 {
						if v18&int32(512) != 0 {
							v91 = v23 << (uint(int32(1)) % 32)
							v93 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v23) {
								v91 = v23 << (uint(int32(2)) % 32)
								v93 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v23) {
									v91 = v23 << (uint(int32(3)) % 32)
									v93 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v23) {
										v91 = v23 << (uint(int32(4)) % 32)
										v93 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v23) {
											v91 = v23 << (uint(int32(5)) % 32)
											v93 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v23) {
												v91 = v23 << (uint(int32(6)) % 32)
												v93 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v23) {
													v91 = v23 << (uint(int32(7)) % 32)
													v93 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v23) {
														v91 = v23 << (uint(int32(8)) % 32)
														v93 = int32(880803840)
													} else {
														v86 = base.B2i32(v23 == int32(1))
														if v23 == int32(1) {
															v87 = int32(1024)
														} else {
															v87 = v23 << (uint(int32(9)) % 32)
														}
														if v23 == int32(1) {
															v90 = int32(864026624)
														} else {
															v90 = int32(872415232)
														}
														v91 = v87
														v93 = v90
													}
												}
											}
										}
									}
								}
							}
						}
						v103 = v91 & int32(1022)
						v104 = v93 | v27
					} else {
						v103 = int32(0)
						v104 = v27
					}
				}
			} else {
				if v23 == int32(0) {
					v103 = int32(0)
					v104 = v27 | int32(2139095040)
				} else {
					v103 = v23
					v104 = v27 | int32(2143289344)
				}
			}
			v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v16))))
			v115 = v110 & int32(1023)
			v119 = v110 << (uint(int32(16)) % 32) & int32(-2147483648)
			v122 = int32(31)
			v123 = int32(base.Ui32(v110)>>(uint(int32(10))%32)) & v122
			if v123 != v122 {
				if v123 != 0 {
					v195 = v115
					v196 = v123<<(uint(int32(23))%32) + v119 + int32(939524096)
				} else {
					if v115 != 0 {
						if v110&int32(512) != 0 {
							v183 = v115 << (uint(int32(1)) % 32)
							v185 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v115) {
								v183 = v115 << (uint(int32(2)) % 32)
								v185 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v115) {
									v183 = v115 << (uint(int32(3)) % 32)
									v185 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v115) {
										v183 = v115 << (uint(int32(4)) % 32)
										v185 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v115) {
											v183 = v115 << (uint(int32(5)) % 32)
											v185 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v115) {
												v183 = v115 << (uint(int32(6)) % 32)
												v185 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v115) {
													v183 = v115 << (uint(int32(7)) % 32)
													v185 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v115) {
														v183 = v115 << (uint(int32(8)) % 32)
														v185 = int32(880803840)
													} else {
														v178 = base.B2i32(v115 == int32(1))
														if v115 == int32(1) {
															v179 = int32(1024)
														} else {
															v179 = v115 << (uint(int32(9)) % 32)
														}
														if v115 == int32(1) {
															v182 = int32(864026624)
														} else {
															v182 = int32(872415232)
														}
														v183 = v179
														v185 = v182
													}
												}
											}
										}
									}
								}
							}
						}
						v195 = v183 & int32(1022)
						v196 = v185 | v119
					} else {
						v195 = int32(0)
						v196 = v119
					}
				}
			} else {
				if v115 == int32(0) {
					v195 = int32(0)
					v196 = v119 | int32(2139095040)
				} else {
					v195 = v115
					v196 = v119 | int32(2143289344)
				}
			}
			v202 = base.F32_add(base.F32_mul(base.F32_reinterpret_i32(v104|v103<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v196|v195<<(uint(int32(13))%32))), v14)
			v204 = v12 + int32(1)
			if v204 != l0 {
				v12 = v204
				v14 = v202
				continue
			} else {
				break
			}
			break
		}
		v211 = v202
	} else {
		v211 = v6
	}
	return v211
}
func F_HalfvecSumCenter(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 float32
	_ = v111
	var v119 int32
	_ = v119
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if int32(0) < v9 {
		v15 = int32(0)
		for {
			v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+int32(8)+v15<<(uint(int32(1))%32)))))
			v28 = v26 & int32(-2147483648)
			v30 = v26 & int32(1023)
			v33 = int32(31)
			v34 = int32(base.Ui32(v26)>>(uint(int32(10))%32)) & v33
			if v34 != v33 {
				if v34 != 0 {
					v105 = v30
					v106 = v34<<(uint(int32(23))%32) + v28 + int32(939524096)
				} else {
					if v30 != 0 {
						if v26&int32(512) != 0 {
							v94 = v30 << (uint(int32(1)) % 32)
							v96 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v30) {
								v94 = v30 << (uint(int32(2)) % 32)
								v96 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v30) {
									v94 = v30 << (uint(int32(3)) % 32)
									v96 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v30) {
										v94 = v30 << (uint(int32(4)) % 32)
										v96 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v30) {
											v94 = v30 << (uint(int32(5)) % 32)
											v96 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v30) {
												v94 = v30 << (uint(int32(6)) % 32)
												v96 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v30) {
													v94 = v30 << (uint(int32(7)) % 32)
													v96 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v30) {
														v94 = v30 << (uint(int32(8)) % 32)
														v96 = int32(880803840)
													} else {
														v89 = base.B2i32(v30 == int32(1))
														if v30 == int32(1) {
															v90 = int32(1024)
														} else {
															v90 = v30 << (uint(int32(9)) % 32)
														}
														if v30 == int32(1) {
															v93 = int32(864026624)
														} else {
															v93 = int32(872415232)
														}
														v94 = v90
														v96 = v93
													}
												}
											}
										}
									}
								}
							}
						}
						v105 = v94 & int32(1022)
						v106 = v96 | v28
					} else {
						v105 = int32(0)
						v106 = v28
					}
				}
			} else {
				if v30 == int32(0) {
					v105 = int32(0)
					v106 = v28 | int32(2139095040)
				} else {
					v105 = v30
					v106 = v28 | int32(2143289344)
				}
			}
			v110 = l1 + v15<<(uint(int32(2))%32)
			v111 = *(*float32)(unsafe.Add(mBase, uint32(v110)))
			*(*float32)(unsafe.Add(mBase, uint32(v110))) = base.F32_add(v111, base.F32_reinterpret_i32(v106|v105<<(uint(int32(13))%32)))
			v119 = v15 + int32(1)
			if v119 != v9 {
				v15 = v119
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_halfvec_cosine_distance(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v39 float64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
			if v19 == v20 {
				v24 = int32(8)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_cosine_distance[0]))
				v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) float64)(m, base.I32_extend16_s(v19), v12+v24, v17+v24)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.F64_gt(v30, float64(1)) != 0 {
						v39 = float64(1)
					} else {
						if base.F64_lt(v30, float64(-1)) == int32(0) {
							v39 = v30
						} else {
							v39 = float64(-1)
						}
					}
					v42 = F_Float8GetDatum(m, base.F64_sub(float64(1), v39))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v42
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v55
						F_errmsg(m, int32(_a_F_halfvec_cosine_distance_0), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_cosine_distance_1), int32(80), int32(_a_F_halfvec_cosine_distance_2))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
func F_halfvec_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v105 float32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 float32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
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
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	v13 = m.G0
	v15 = v13 - int32(_a_F_halfvec_in_0)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = v18
	goto L1
L1:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
	goto L3
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v43 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v33 == int32(32))|base.B2i32(base.Ui32((v33-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v22 = v22 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L25
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L25
	} else {
		goto L106
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L25
	} else {
		goto L101
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L25
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L25
	} else {
		goto L93
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L25
	} else {
		goto L89
	}
L11:
	;
	v46 = v22
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L25
	} else {
		goto L84
	}
L14:
	;
	v59 = v46 + int32(1)
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46)+1)))
	goto L16
L15:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v70 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v60 == int32(32))|base.B2i32(base.Ui32((v60-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v46 = v59
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v73 = v59
	v80 = int32(0)
	goto L20
L19:
	;
	v294 = v237
	goto L71
L20:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L67
	}
L22:
	;
	if base.B2i32(v87 == int32(32))|base.B2i32(base.Ui32((v87-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v73 = v73 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v97 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_halfvec_in[0])) = int32(0)
	v105 = F_strtof(m, v73, v15+int32(124))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	if v109 == v73 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v111 = base.I32_reinterpret_f32(v105)
	v113 = int32(base.Ui32(v111) >> (uint(int32(16)) % 32))
	v114 = base.F32_abs(v105)
	if base.F32_ne(v114, math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_CheckElement_1(m, v215&int32(_a_F_halfvec_in_1))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L25
	} else {
		goto L53
	}
L29:
	;
	v123 = v113 & int32(_a_F_halfvec_in_2)
	v125 = v111 & int32(_a_F_halfvec_in_3)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v114)) {
		v197 = v123 | int32(base.Ui32(v125)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_in_4)
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v209 = v113 & int32(_a_F_halfvec_in_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(128)+v80<<(uint(int32(1))%32)))) = uint16(v209)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_in[0]))
	if v212 == int32(68) {
		goto L5
	} else {
		goto L52
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(128)+v80<<(uint(int32(1))%32)))) = uint16(v197)
	if v197&int32(_a_F_halfvec_in_6) != int32(_a_F_halfvec_in_7) {
		v215 = v197
		goto L28
	} else {
		goto L51
	}
L33:
	;
	v137 = int32(base.Ui32(v111)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v137) < base.Ui32(int32(99)) {
		v197 = v123
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v137) <= base.Ui32(int32(112)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v149 = int32(1)<<(uint(v137-int32(90))%32) + int32(base.Ui32(v125)>>(uint(int32(113)-v137)%32))
	v151 = v149 | v111
	v152 = v149
	goto L37
L36:
	;
	v151 = v111
	v152 = v125
	goto L37
L37:
	;
	v154 = int32(base.Ui32(v152) >> (uint(int32(13)) % 32))
	v159 = int32(1)
	v163 = int32(3)
	v164 = int32(base.Ui32(v152)>>(uint(int32(12))%32)) & v163
	if base.B2i32(v164 != v163)&(base.B2i32(v151&int32(4095) == int32(0))|base.B2i32(v164 != v159)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v175 = v154
	goto L40
L39:
	;
	v175 = v154 + v159
	goto L40
L40:
	;
	v177 = base.B2i32(v175 == int32(1024))
	if v175 == int32(1024) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v178 = int32(-126)
	goto L43
L42:
	;
	v178 = int32(-127)
	goto L43
L43:
	;
	v179 = v178 + v137
	if int32(16) <= v179 {
		v197 = v123 | int32(_a_F_halfvec_in_7)
		goto L32
	} else {
		goto L44
	}
L44:
	;
	if int32(-15) < v179 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v189 = v179<<(uint(int32(10))%32) + int32(_a_F_halfvec_in_8) | v123
	goto L47
L46:
	;
	v189 = v123
	goto L47
L47:
	;
	if v175 == int32(1024) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v191 = int32(0)
	goto L50
L49:
	;
	v191 = v175
	goto L50
L50:
	;
	v197 = v189 | v191
	goto L32
L51:
	;
	goto L5
L52:
	;
	v215 = v209
	goto L28
L53:
	;
	v224 = v109
	goto L54
L54:
	;
	v237 = v224 + int32(1)
	v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v224))))
	goto L56
L55:
	;
	v249 = v80 + int32(1)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v250 != int32(44) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if base.B2i32(v238 == int32(32))|base.B2i32(base.Ui32((v238-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v224 = v237
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v250 == int32(93) {
		goto L19
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v249 != int32(_a_F_halfvec_in_9) {
		v73 = v237
		v80 = v249
		goto L20
	} else {
		goto L66
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v18
	F_errmsg(m, int32(_a_F_halfvec_in_10), v15+int32(48))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(265), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L21
L67:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(_a_F_halfvec_in_9)
	F_errmsg(m, int32(_a_F_halfvec_in_13), v15-int32(-64))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(218), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v308 = int32(*(*int8)(unsafe.Add(mBase, uint32(v294))))
	goto L73
L72:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v318 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	if base.B2i32(v308 == int32(32))|base.B2i32(base.Ui32((v308-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v294 = v294 + int32(1)
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	F_CheckDim_1(m, v249)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	if base.B2i32(v17 != int32(-1))&base.B2i32(v249 != v17) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v327 = F_mul_size(m, int32(2), v249)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v329 = F_add_size(m, int32(8), v327)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	v331 = F_palloc0(m, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v331)+4)) = uint16(v249)
	v334 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v329 << (uint(v334) % 32)
	v340 = v80<<(uint(int32(1))%32) + v334
	if v340 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	base.MemoryCopy(m, v331+int32(8), v15+int32(128), v340)
	goto L83
L82:
	;
	goto L83
L83:
	;
	m.G0 = v15 + int32(_a_F_halfvec_in_0)
	return v331
L84:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v18
	F_errmsg(m, int32(_a_F_halfvec_in_10), v15+int32(112))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	F_errdetail(m, int32(_a_F_halfvec_in_14), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L25
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(198), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L25
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
	F_errcode(m, int32(130))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_halfvec_in_15), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(208), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v18
	F_errmsg(m, int32(_a_F_halfvec_in_10), v15)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(227), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L25
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L25
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	F_errmsg(m, int32(_a_F_halfvec_in_10), v15+int32(16))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L25
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(237), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L25
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L25
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v18
	F_errmsg(m, int32(_a_F_halfvec_in_10), v15+int32(96))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L25
	} else {
		goto L103
	}
L103:
	;
	F_errdetail(m, int32(_a_F_halfvec_in_16), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L25
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(276), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L25
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L25
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v17
	F_errmsg(m, int32(_a_F_halfvec_in_17), v15+int32(80))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L25
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(92), int32(_a_F_halfvec_in_18))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L25
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L25
	} else {
		goto L111
	}
L111:
	;
	v476 = F_pnstrdup(m, v73, v109-v73)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v476
	F_errmsg(m, int32(_a_F_halfvec_in_19), v15+int32(32))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(245), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L25
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_l1_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13933(m, l0, int32(_a_F_halfvec_l1_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_halfvec_l2_norm(m *base.Module, l0 int32) int32 {
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
	var v19 float64
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v119 int32
	_ = v119
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+4)))
		if int32(0) < v11 {
			v17 = int32(0)
			v19 = float64(0)
			for {
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(8)+v17<<(uint(int32(1))%32)))))
				v29 = v24 & int32(1023)
				v33 = v24 << (uint(int32(16)) % 32) & int32(-2147483648)
				v36 = int32(31)
				v37 = int32(base.Ui32(v24)>>(uint(int32(10))%32)) & v36
				if v37 != v36 {
					if v37 != 0 {
						v109 = v29
						v110 = v37<<(uint(int32(23))%32) + v33 + int32(939524096)
					} else {
						if v29 != 0 {
							if v24&int32(512) != 0 {
								v97 = v29 << (uint(int32(1)) % 32)
								v99 = int32(939524096)
							} else {
								if base.Ui32(int32(255)) < base.Ui32(v29) {
									v97 = v29 << (uint(int32(2)) % 32)
									v99 = int32(931135488)
								} else {
									if base.Ui32(int32(127)) < base.Ui32(v29) {
										v97 = v29 << (uint(int32(3)) % 32)
										v99 = int32(922746880)
									} else {
										if base.Ui32(int32(63)) < base.Ui32(v29) {
											v97 = v29 << (uint(int32(4)) % 32)
											v99 = int32(914358272)
										} else {
											if base.Ui32(int32(31)) < base.Ui32(v29) {
												v97 = v29 << (uint(int32(5)) % 32)
												v99 = int32(905969664)
											} else {
												if base.Ui32(int32(15)) < base.Ui32(v29) {
													v97 = v29 << (uint(int32(6)) % 32)
													v99 = int32(897581056)
												} else {
													if base.Ui32(int32(7)) < base.Ui32(v29) {
														v97 = v29 << (uint(int32(7)) % 32)
														v99 = int32(889192448)
													} else {
														if base.Ui32(int32(3)) < base.Ui32(v29) {
															v97 = v29 << (uint(int32(8)) % 32)
															v99 = int32(880803840)
														} else {
															v92 = base.B2i32(v29 == int32(1))
															if v29 == int32(1) {
																v93 = int32(1024)
															} else {
																v93 = v29 << (uint(int32(9)) % 32)
															}
															if v29 == int32(1) {
																v96 = int32(864026624)
															} else {
																v96 = int32(872415232)
															}
															v97 = v93
															v99 = v96
														}
													}
												}
											}
										}
									}
								}
							}
							v109 = v97 & int32(1022)
							v110 = v99 | v33
						} else {
							v109 = int32(0)
							v110 = v33
						}
					}
				} else {
					if v29 == int32(0) {
						v109 = int32(0)
						v110 = v33 | int32(2139095040)
					} else {
						v109 = v29
						v110 = v33 | int32(2143289344)
					}
				}
				v115 = base.F64_promote_f32(base.F32_reinterpret_i32(v110 | v109<<(uint(int32(13))%32)))
				v117 = base.F64_add(base.F64_mul(v115, v115), v19)
				v119 = v17 + int32(1)
				if v119 != v11 {
					v17 = v119
					v19 = v117
					continue
				} else {
					break
				}
				break
			}
			v128 = base.F64_sqrt(v117)
		} else {
			v128 = float64(0)
		}
		v129 = F_Float8GetDatum(m, v128)
		mBase = m.M
		v130 = m.ExcPending
		if v130 != 0 {
			return int32(0)
		} else {
			return v129
		}
	}
}
func F_halfvec_subvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if int32(0) < v16 {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if int32(0) < v21 {
				if v19 < v21 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_halfvec_subvector_0), int32(0))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_halfvec_subvector_1), int32(971), int32(_a_F_halfvec_subvector_2))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
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
					v25 = v21
					if v19-v16 < v21 {
						v31 = v19 + int32(1)
					} else {
						v31 = v21 + v16
					}
					v32 = v31 - v25
					F_CheckDim_1(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v37 = F_mul_size(m, int32(2), v32)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = F_add_size(m, int32(8), v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_palloc0(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v32)
									*(*int32)(unsafe.Add(mBase, uint32(v41))) = v39 << (uint(int32(2)) % 32)
									if v32 <= int32(0) {
									} else {
										v50 = v32 & int32(3)
										v52 = v41 + int32(8)
										v57 = v12 + v25<<(uint(int32(1))%32) + int32(6)
										v58 = int32(0)
										if base.Ui32(v25-v31) <= base.Ui32(int32(-4)) {
											v65 = v58
											v67 = int32(0)
											for {
												v76 = v65 << (uint(int32(1)) % 32)
												v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v76))) = uint16(v79)
												v82 = v76 | int32(2)
												v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v82))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v82))) = uint16(v85)
												v87 = int32(4)
												v88 = v76 | v87
												v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v88))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v88))) = uint16(v91)
												v94 = v76 | int32(6)
												v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v94))) = uint16(v97)
												v100 = v65 + v87
												v102 = v67 + v87
												if v102 != v32&int32(2147483644) {
													v65 = v100
													v67 = v102
													continue
												} else {
													break
												}
												break
											}
											if v50 == int32(0) {
											} else {
												v106 = v100
												v116 = v106
												v125 = int32(0)
												for {
													v126 = int32(1)
													v127 = v116 << (uint(v126) % 32)
													v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v57))))
													*(*uint16)(unsafe.Add(mBase, uint32(v52+v127))) = uint16(v130)
													v135 = v125 + v126
													if v135 != v50 {
														v116 = v116 + v126
														v125 = v135
														continue
													} else {
														break
													}
													break
												}
											}
										} else {
											v106 = v58
											v116 = v106
											v125 = int32(0)
											for {
												v126 = int32(1)
												v127 = v116 << (uint(v126) % 32)
												v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v127))) = uint16(v130)
												v135 = v125 + v126
												if v135 != v50 {
													v116 = v116 + v126
													v125 = v135
													continue
												} else {
													break
												}
												break
											}
										}
									}
									return v41
								}
							}
						}
					}
				}
			} else {
				v25 = int32(1)
				if v19-v16 < v21 {
					v31 = v19 + int32(1)
				} else {
					v31 = v21 + v16
				}
				v32 = v31 - v25
				F_CheckDim_1(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v37 = F_mul_size(m, int32(2), v32)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_add_size(m, int32(8), v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = F_palloc0(m, v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v32)
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = v39 << (uint(int32(2)) % 32)
								if v32 <= int32(0) {
								} else {
									v50 = v32 & int32(3)
									v52 = v41 + int32(8)
									v57 = v12 + v25<<(uint(int32(1))%32) + int32(6)
									v58 = int32(0)
									if base.Ui32(v25-v31) <= base.Ui32(int32(-4)) {
										v65 = v58
										v67 = int32(0)
										for {
											v76 = v65 << (uint(int32(1)) % 32)
											v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v76))) = uint16(v79)
											v82 = v76 | int32(2)
											v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v82))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v82))) = uint16(v85)
											v87 = int32(4)
											v88 = v76 | v87
											v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v88))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v88))) = uint16(v91)
											v94 = v76 | int32(6)
											v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v94))) = uint16(v97)
											v100 = v65 + v87
											v102 = v67 + v87
											if v102 != v32&int32(2147483644) {
												v65 = v100
												v67 = v102
												continue
											} else {
												break
											}
											break
										}
										if v50 == int32(0) {
										} else {
											v106 = v100
											v116 = v106
											v125 = int32(0)
											for {
												v126 = int32(1)
												v127 = v116 << (uint(v126) % 32)
												v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v127))) = uint16(v130)
												v135 = v125 + v126
												if v135 != v50 {
													v116 = v116 + v126
													v125 = v135
													continue
												} else {
													break
												}
												break
											}
										}
									} else {
										v106 = v58
										v116 = v106
										v125 = int32(0)
										for {
											v126 = int32(1)
											v127 = v116 << (uint(v126) % 32)
											v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v127))) = uint16(v130)
											v135 = v125 + v126
											if v135 != v50 {
												v116 = v116 + v126
												v125 = v135
												continue
											} else {
												break
											}
											break
										}
									}
								}
								return v41
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_halfvec_subvector_0), int32(0))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_halfvec_subvector_1), int32(954), int32(_a_F_halfvec_subvector_2))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
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
