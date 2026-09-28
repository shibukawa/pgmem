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
func F_halfvec_cosine_distance(m *base.Module, l0 int32) int64 {
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
					return int64(0)
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
					m.G0 = v9 + int32(16)
					return base.I64_reinterpret_f64(base.F64_sub(float64(1), v39))
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v54
						F_errmsg(m, int32(_a_F_halfvec_cosine_distance_0), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_cosine_distance_1), int32(80), int32(_a_F_halfvec_cosine_distance_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
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
func F_halfvec_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v110 float32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 float32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	v14 = m.G0
	v16 = v14 - int32(_a_F_halfvec_in_0)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = v19
	goto L1
L1:
	;
	v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	goto L3
L2:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v45 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v35 == int32(32))|base.B2i32(base.Ui32((v35-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v23 = v23 + int32(1)
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
	v479 = m.ExcPending
	if v479 != 0 {
		goto L25
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L25
	} else {
		goto L106
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L25
	} else {
		goto L101
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L25
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L25
	} else {
		goto L93
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L25
	} else {
		goto L89
	}
L11:
	;
	v48 = base.I32_wrap_i64(v18)
	v49 = v23
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L25
	} else {
		goto L84
	}
L14:
	;
	v63 = v49 + int32(1)
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49)+1)))
	goto L16
L15:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v74 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v64 == int32(32))|base.B2i32(base.Ui32((v64-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v49 = v63
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v77 = v63
	v84 = int32(0)
	goto L20
L19:
	;
	v300 = v243
	goto L71
L20:
	;
	v92 = int32(*(*int8)(unsafe.Add(mBase, uint32(v77))))
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L25
	} else {
		goto L67
	}
L22:
	;
	if base.B2i32(v92 == int32(32))|base.B2i32(base.Ui32((v92-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v77 = v77 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v102 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_halfvec_in[0])) = int32(0)
	v110 = F_strtof(m, v77, v16+int32(124))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int64(0)
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v16)+124))
	if v114 == v77 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v116 = base.I32_reinterpret_f32(v110)
	v118 = int32(base.Ui32(v116) >> (uint(int32(16)) % 32))
	v119 = base.F32_abs(v110)
	if base.F32_ne(v119, math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_CheckElement_1(m, v220&int32(_a_F_halfvec_in_1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L25
	} else {
		goto L53
	}
L29:
	;
	v128 = v118 & int32(_a_F_halfvec_in_2)
	v130 = v116 & int32(_a_F_halfvec_in_3)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v119)) {
		v202 = v128 | int32(base.Ui32(v130)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_in_4)
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v214 = v118 & int32(_a_F_halfvec_in_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(128)+v84<<(uint(int32(1))%32)))) = uint16(v214)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_in[0]))
	if v217 == int32(68) {
		goto L5
	} else {
		goto L52
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(128)+v84<<(uint(int32(1))%32)))) = uint16(v202)
	if v202&int32(_a_F_halfvec_in_6) != int32(_a_F_halfvec_in_7) {
		v220 = v202
		goto L28
	} else {
		goto L51
	}
L33:
	;
	v142 = int32(base.Ui32(v116)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v142) < base.Ui32(int32(99)) {
		v202 = v128
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v142) <= base.Ui32(int32(112)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v154 = int32(1)<<(uint(v142-int32(90))%32) + int32(base.Ui32(v130)>>(uint(int32(113)-v142)%32))
	v156 = v154 | v116
	v157 = v154
	goto L37
L36:
	;
	v156 = v116
	v157 = v130
	goto L37
L37:
	;
	v159 = int32(base.Ui32(v157) >> (uint(int32(13)) % 32))
	v164 = int32(1)
	v168 = int32(3)
	v169 = int32(base.Ui32(v157)>>(uint(int32(12))%32)) & v168
	if base.B2i32(v169 != v168)&(base.B2i32(v156&int32(4095) == int32(0))|base.B2i32(v169 != v164)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v180 = v159
	goto L40
L39:
	;
	v180 = v159 + v164
	goto L40
L40:
	;
	v182 = base.B2i32(v180 == int32(1024))
	if v180 == int32(1024) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v183 = int32(-126)
	goto L43
L42:
	;
	v183 = int32(-127)
	goto L43
L43:
	;
	v184 = v183 + v142
	if int32(16) <= v184 {
		v202 = v128 | int32(_a_F_halfvec_in_7)
		goto L32
	} else {
		goto L44
	}
L44:
	;
	if int32(-15) < v184 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v194 = v184<<(uint(int32(10))%32) + int32(_a_F_halfvec_in_8) | v128
	goto L47
L46:
	;
	v194 = v128
	goto L47
L47:
	;
	if v180 == int32(1024) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v196 = int32(0)
	goto L50
L49:
	;
	v196 = v180
	goto L50
L50:
	;
	v202 = v194 | v196
	goto L32
L51:
	;
	goto L5
L52:
	;
	v220 = v214
	goto L28
L53:
	;
	v229 = v114
	goto L54
L54:
	;
	v243 = v229 + int32(1)
	v244 = int32(*(*int8)(unsafe.Add(mBase, uint32(v229))))
	goto L56
L55:
	;
	v255 = v84 + int32(1)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v256 != int32(44) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if base.B2i32(v244 == int32(32))|base.B2i32(base.Ui32((v244-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v229 = v243
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v256 == int32(93) {
		goto L19
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v255 != int32(_a_F_halfvec_in_9) {
		v77 = v243
		v84 = v255
		goto L20
	} else {
		goto L66
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v19
	F_errmsg(m, int32(_a_F_halfvec_in_10), v16+int32(48))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(265), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
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
	v287 = m.ExcPending
	if v287 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = int32(_a_F_halfvec_in_9)
	F_errmsg(m, int32(_a_F_halfvec_in_13), v16-int32(-64))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(218), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	v315 = int32(*(*int8)(unsafe.Add(mBase, uint32(v300))))
	goto L73
L72:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v325 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	if base.B2i32(v315 == int32(32))|base.B2i32(base.Ui32((v315-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v300 = v300 + int32(1)
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	F_CheckDim_1(m, v255)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	if base.B2i32(v48 != int32(-1))&base.B2i32(v255 != v48) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v334 = F_mul_size(m, int32(2), v255)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v336 = F_add_size(m, int32(8), v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	v338 = F_palloc0(m, v336)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v338)+4)) = uint16(v255)
	v341 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = v336 << (uint(v341) % 32)
	v347 = v84<<(uint(int32(1))%32) + v341
	if v347 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	base.MemoryCopy(m, v338+int32(8), v16+int32(128), v347)
	goto L83
L82:
	;
	goto L83
L83:
	;
	m.G0 = v16 + int32(_a_F_halfvec_in_0)
	return base.I64_extend_i32_u(v338)
L84:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v19
	F_errmsg(m, int32(_a_F_halfvec_in_10), v16+int32(112))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	v373 = F_errdetail(m, int32(_a_F_halfvec_in_14), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L25
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(198), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
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
	v386 = m.ExcPending
	if v386 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_halfvec_in_15), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(208), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
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
	v402 = m.ExcPending
	if v402 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
	F_errmsg(m, int32(_a_F_halfvec_in_10), v16)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(227), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	v418 = m.ExcPending
	if v418 != 0 {
		goto L25
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	F_errmsg(m, int32(_a_F_halfvec_in_10), v16+int32(16))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L25
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(237), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L25
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v19
	F_errmsg(m, int32(_a_F_halfvec_in_10), v16+int32(96))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L25
	} else {
		goto L103
	}
L103:
	;
	v445 = F_errdetail(m, int32(_a_F_halfvec_in_16), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L25
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(276), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
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
	v458 = m.ExcPending
	if v458 != 0 {
		goto L25
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v48
	F_errmsg(m, int32(_a_F_halfvec_in_17), v16+int32(80))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L25
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(92), int32(_a_F_halfvec_in_18))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
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
	v482 = m.ExcPending
	if v482 != 0 {
		goto L25
	} else {
		goto L111
	}
L111:
	;
	v484 = F_pnstrdup(m, v77, v114-v77)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v484
	F_errmsg(m, int32(_a_F_halfvec_in_19), v16+int32(32))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_halfvec_in_11), int32(245), int32(_a_F_halfvec_in_12))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
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
func F_halfvec_l1_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14300(m, l0, int32(_a_F_halfvec_l1_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_halfvec_l2_norm(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v21 float64
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v121 int32
	_ = v121
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+4)))
		if v11 <= int32(0) {
			return int64(0)
		} else {
			v19 = int32(0)
			v21 = float64(0)
			for {
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(8)+v19<<(uint(int32(1))%32)))))
				v31 = v26 & int32(1023)
				v35 = v26 << (uint(int32(16)) % 32) & int32(-2147483648)
				v38 = int32(31)
				v39 = int32(base.Ui32(v26)>>(uint(int32(10))%32)) & v38
				if v39 != v38 {
					if v39 != 0 {
						v111 = v31
						v112 = v39<<(uint(int32(23))%32) + v35 + int32(939524096)
					} else {
						if v31 != 0 {
							if v26&int32(512) != 0 {
								v99 = v31 << (uint(int32(1)) % 32)
								v101 = int32(939524096)
							} else {
								if base.Ui32(int32(255)) < base.Ui32(v31) {
									v99 = v31 << (uint(int32(2)) % 32)
									v101 = int32(931135488)
								} else {
									if base.Ui32(int32(127)) < base.Ui32(v31) {
										v99 = v31 << (uint(int32(3)) % 32)
										v101 = int32(922746880)
									} else {
										if base.Ui32(int32(63)) < base.Ui32(v31) {
											v99 = v31 << (uint(int32(4)) % 32)
											v101 = int32(914358272)
										} else {
											if base.Ui32(int32(31)) < base.Ui32(v31) {
												v99 = v31 << (uint(int32(5)) % 32)
												v101 = int32(905969664)
											} else {
												if base.Ui32(int32(15)) < base.Ui32(v31) {
													v99 = v31 << (uint(int32(6)) % 32)
													v101 = int32(897581056)
												} else {
													if base.Ui32(int32(7)) < base.Ui32(v31) {
														v99 = v31 << (uint(int32(7)) % 32)
														v101 = int32(889192448)
													} else {
														if base.Ui32(int32(3)) < base.Ui32(v31) {
															v99 = v31 << (uint(int32(8)) % 32)
															v101 = int32(880803840)
														} else {
															v94 = base.B2i32(v31 == int32(1))
															if v31 == int32(1) {
																v95 = int32(1024)
															} else {
																v95 = v31 << (uint(int32(9)) % 32)
															}
															if v31 == int32(1) {
																v98 = int32(864026624)
															} else {
																v98 = int32(872415232)
															}
															v99 = v95
															v101 = v98
														}
													}
												}
											}
										}
									}
								}
							}
							v111 = v99 & int32(1022)
							v112 = v101 | v35
						} else {
							v111 = int32(0)
							v112 = v35
						}
					}
				} else {
					if v31 == int32(0) {
						v111 = int32(0)
						v112 = v35 | int32(2139095040)
					} else {
						v111 = v31
						v112 = v35 | int32(2143289344)
					}
				}
				v117 = base.F64_promote_f32(base.F32_reinterpret_i32(v112 | v111<<(uint(int32(13))%32)))
				v119 = base.F64_add(base.F64_mul(v117, v117), v21)
				v121 = v19 + int32(1)
				if v121 != v11 {
					v19 = v121
					v21 = v119
					continue
				} else {
					break
				}
				break
			}
			return base.I64_reinterpret_f64(base.F64_sqrt(v119))
		}
	}
}
func F_halfvec_subvector(m *base.Module, l0 int32) int64 {
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
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if int32(0) < v16 {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if int32(0) < v21 {
				if v19 < v21 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_halfvec_subvector_0), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_halfvec_subvector_1), int32(971), int32(_a_F_halfvec_subvector_2))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
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
						return int64(0)
					} else {
						v37 = F_mul_size(m, int32(2), v32)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = F_add_size(m, int32(8), v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = F_palloc0(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
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
									return base.I64_extend_i32_u(v41)
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
					return int64(0)
				} else {
					v37 = F_mul_size(m, int32(2), v32)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = F_add_size(m, int32(8), v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v41 = F_palloc0(m, v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
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
								return base.I64_extend_i32_u(v41)
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_halfvec_subvector_0), int32(0))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_halfvec_subvector_1), int32(954), int32(_a_F_halfvec_subvector_2))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
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
