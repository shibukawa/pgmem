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
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 float32
	_ = v204
	var v206 int32
	_ = v206
	var v213 float32
	_ = v213
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
			v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v16))))
			v116 = v111 & int32(1023)
			v120 = v111 << (uint(int32(16)) % 32) & int32(-2147483648)
			v123 = int32(31)
			v124 = int32(base.Ui32(v111)>>(uint(int32(10))%32)) & v123
			if v124 != v123 {
				if v124 != 0 {
					v196 = v116
					v197 = v124<<(uint(int32(23))%32) + v120 + int32(939524096)
				} else {
					if v116 != 0 {
						if v111&int32(512) != 0 {
							v184 = v116 << (uint(int32(1)) % 32)
							v186 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v116) {
								v184 = v116 << (uint(int32(2)) % 32)
								v186 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v116) {
									v184 = v116 << (uint(int32(3)) % 32)
									v186 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v116) {
										v184 = v116 << (uint(int32(4)) % 32)
										v186 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v116) {
											v184 = v116 << (uint(int32(5)) % 32)
											v186 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v116) {
												v184 = v116 << (uint(int32(6)) % 32)
												v186 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v116) {
													v184 = v116 << (uint(int32(7)) % 32)
													v186 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v116) {
														v184 = v116 << (uint(int32(8)) % 32)
														v186 = int32(880803840)
													} else {
														v179 = base.B2i32(v116 == int32(1))
														if v116 == int32(1) {
															v180 = int32(1024)
														} else {
															v180 = v116 << (uint(int32(9)) % 32)
														}
														if v116 == int32(1) {
															v183 = int32(864026624)
														} else {
															v183 = int32(872415232)
														}
														v184 = v180
														v186 = v183
													}
												}
											}
										}
									}
								}
							}
						}
						v196 = v184 & int32(1022)
						v197 = v186 | v120
					} else {
						v196 = int32(0)
						v197 = v120
					}
				}
			} else {
				if v116 == int32(0) {
					v196 = int32(0)
					v197 = v120 | int32(2139095040)
				} else {
					v196 = v116
					v197 = v120 | int32(2143289344)
				}
			}
			v204 = base.F32_add(base.F32_mul(base.F32_reinterpret_i32(v104|v103<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v197|v196<<(uint(int32(13))%32))), v14)
			v206 = v12 + int32(1)
			if v206 != l0 {
				v12 = v206
				v14 = v204
				continue
			} else {
				break
			}
			break
		}
		v213 = v204
	} else {
		v213 = v6
	}
	return v213
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
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 float32
	_ = v112
	var v120 int32
	_ = v120
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
					v107 = v34<<(uint(int32(23))%32) + v28 + int32(939524096)
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
						v107 = v96 | v28
					} else {
						v105 = int32(0)
						v107 = v28
					}
				}
			} else {
				if v30 == int32(0) {
					v105 = int32(0)
					v107 = v28 | int32(2139095040)
				} else {
					v105 = v30
					v107 = v28 | int32(2143289344)
				}
			}
			v111 = l1 + v15<<(uint(int32(2))%32)
			v112 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
			*(*float32)(unsafe.Add(mBase, uint32(v111))) = base.F32_add(v112, base.F32_reinterpret_i32(v107|v105<<(uint(int32(13))%32)))
			v120 = v15 + int32(1)
			if v120 != v9 {
				v15 = v120
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
				v29 = *(*int32)(unsafe.Add(mBase, _consts[1381]))
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
						F_errmsg(m, int32(500747), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525467), int32(80), int32(160136))
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
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
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
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
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
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
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
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
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
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	v13 = m.G0
	v15 = v13 - int32(32128)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = v18
	goto L1
L1:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19))))
	goto L3
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v43 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v33 == int32(32))|base.B2i32(base.Ui32((v33-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v19 = v19 + int32(1)
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
	v473 = m.ExcPending
	if v473 != 0 {
		goto L25
	} else {
		goto L114
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L25
	} else {
		goto L110
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L25
	} else {
		goto L105
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L25
	} else {
		goto L101
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L25
	} else {
		goto L97
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L25
	} else {
		goto L93
	}
L11:
	;
	v48 = v19
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L25
	} else {
		goto L88
	}
L14:
	;
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+1)))
	v60 = v48 + int32(1)
	goto L16
L15:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v70 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v58 == int32(32))|base.B2i32(base.Ui32((v58-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v48 = v60
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v73 = v60
	v80 = int32(0)
	goto L20
L19:
	;
	v295 = v238
	goto L74
L20:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L25
	} else {
		goto L70
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
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
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
	F_CheckElement_1(m, v216&int32(65535))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L25
	} else {
		goto L56
	}
L29:
	;
	v125 = v111 & int32(8388607)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v114)) {
		v198 = v113&int32(32768) | int32(base.Ui32(v125)>>(uint(int32(13))%32)) | int32(32256)
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v210 = v113 & int32(64512)
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(128)+v80<<(uint(int32(1))%32)))) = uint16(v210)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v213 == int32(68) {
		goto L5
	} else {
		goto L55
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(128)+v80<<(uint(int32(1))%32)))) = uint16(v198)
	if v198&int32(32767) != int32(31744) {
		v216 = v198
		goto L28
	} else {
		goto L54
	}
L33:
	;
	v135 = v113 & int32(32768)
	v139 = int32(base.Ui32(v111)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v139) < base.Ui32(int32(99)) {
		v198 = v135
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v139) <= base.Ui32(int32(112)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v151 = int32(1)<<(uint(v139-int32(90))%32) + int32(base.Ui32(v125)>>(uint(int32(113)-v139)%32))
	v153 = v151 | v111
	v154 = v151
	goto L37
L36:
	;
	v153 = v111
	v154 = v125
	goto L37
L37:
	;
	v156 = int32(base.Ui32(v154) >> (uint(int32(13)) % 32))
	v159 = int32(3)
	v160 = int32(base.Ui32(v154)>>(uint(int32(12))%32)) & v159
	if v160 != v159 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v177 = base.B2i32(v171 == int32(1024))
	if v171 == int32(1024) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	if v160 != int32(1) {
		v171 = v156
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v171 = v156 + int32(1)
	goto L38
L42:
	;
	if v153&int32(4095) == int32(0) {
		v171 = v156
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v178 = int32(-126)
	goto L46
L45:
	;
	v178 = int32(-127)
	goto L46
L46:
	;
	v179 = v178 + v139
	if int32(16) <= v179 {
		v198 = v135 | int32(31744)
		goto L32
	} else {
		goto L47
	}
L47:
	;
	if int32(-15) < v179 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v189 = v179<<(uint(int32(10))%32) + int32(15360) | v135
	goto L50
L49:
	;
	v189 = v135
	goto L50
L50:
	;
	if v171 == int32(1024) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v191 = int32(0)
	goto L53
L52:
	;
	v191 = v171
	goto L53
L53:
	;
	v198 = v189 | v191
	goto L32
L54:
	;
	goto L5
L55:
	;
	v216 = v210
	goto L28
L56:
	;
	v225 = v109
	goto L57
L57:
	;
	v238 = v225 + int32(1)
	v239 = int32(*(*int8)(unsafe.Add(mBase, uint32(v225))))
	goto L59
L58:
	;
	v250 = v80 + int32(1)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v251 != int32(44) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if base.B2i32(v239 == int32(32))|base.B2i32(base.Ui32((v239-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v225 = v238
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v251 == int32(93) {
		goto L19
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v250 != int32(16000) {
		v73 = v238
		v80 = v250
		goto L20
	} else {
		goto L69
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v18
	F_errmsg(m, int32(762138), v15+int32(48))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(525467), int32(265), int32(293538))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L25
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
	goto L21
L70:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(16000)
	F_errmsg(m, int32(156250), v15-int32(-64))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L25
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(525467), int32(218), int32(293538))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L25
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v295))))
	goto L76
L75:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v319 != 0 {
		goto L7
	} else {
		goto L78
	}
L76:
	;
	if base.B2i32(v309 == int32(32))|base.B2i32(base.Ui32((v309-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v295 = v295 + int32(1)
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_CheckDim_1(m, v250)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	if base.B2i32(v17 != int32(-1))&base.B2i32(v250 != v17) != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	v328 = F_mul_size(m, int32(2), v250)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v330 = F_add_size(m, int32(8), v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L25
	} else {
		goto L82
	}
L82:
	;
	v332 = F_palloc0(m, v330)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v332)+4)) = uint16(v250)
	v335 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v330 << (uint(v335) % 32)
	v345 = v80<<(uint(int32(1))%32) + v335
	if v345 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	m.G0 = v15 + int32(32128)
	return v332
L85:
	;
	v346 = F__emscripten_memcpy_bulkmem(m, v332+int32(8), v15+int32(128), v345)
	mBase = m.M
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L25
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v18
	F_errmsg(m, int32(762138), v15+int32(112))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	F_errdetail(m, int32(699953), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(525467), int32(198), int32(293538))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
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
	F_errcode(m, int32(130))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(285025), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(525467), int32(208), int32(293538))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
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
	v396 = m.ExcPending
	if v396 != 0 {
		goto L25
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v18
	F_errmsg(m, int32(762138), v15)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L25
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(525467), int32(227), int32(293538))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v412 = m.ExcPending
	if v412 != 0 {
		goto L25
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	F_errmsg(m, int32(762138), v15+int32(16))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L25
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(525467), int32(237), int32(293538))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L25
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L25
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v18
	F_errmsg(m, int32(762138), v15+int32(96))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L25
	} else {
		goto L107
	}
L107:
	;
	F_errdetail(m, int32(673183), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L25
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(525467), int32(276), int32(293538))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
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
	F_errcode(m, int32(130))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L25
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v17
	F_errmsg(m, int32(489736), v15+int32(80))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(525467), int32(92), int32(302808))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L25
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L25
	} else {
		goto L115
	}
L115:
	;
	v478 = F_pnstrdup(m, v73, v109-v73)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L25
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v478
	F_errmsg(m, int32(514863), v15+int32(32))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(525467), int32(245), int32(293538))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L25
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_l1_distance(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 float32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
			if v17 != v18 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
						v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
						F_errmsg(m, int32(500747), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525467), int32(80), int32(160136))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
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
				v40 = int32(8)
				v45 = *(*int32)(unsafe.Add(mBase, _consts[1382]))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v17), v10+v40, v15+v40)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = F_Float8GetDatum(m, base.F64_promote_f32(v46))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v49
					}
				}
			}
		}
	}
}
func F_halfvec_l2_norm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
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
	var v116 float64
	_ = v116
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v125 float64
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	v4 = float64(0)
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
			v19 = v4
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
				v116 = base.F64_promote_f32(base.F32_reinterpret_i32(v110 | v109<<(uint(int32(13))%32)))
				v118 = base.F64_add(base.F64_mul(v116, v116), v19)
				v120 = v17 + int32(1)
				if v120 != v11 {
					v17 = v120
					v19 = v118
					continue
				} else {
					break
				}
				break
			}
			v125 = v118
		} else {
			v125 = v4
		}
		v128 = F_Float8GetDatum(m, base.F64_sqrt(v125))
		mBase = m.M
		v129 = m.ExcPending
		if v129 != 0 {
			return int32(0)
		} else {
			return v128
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
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
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(285025), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(525467), int32(971), int32(219214))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
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
											v66 = v58
											v70 = int32(0)
											for {
												v77 = v66 << (uint(int32(1)) % 32)
												v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v77))) = uint16(v80)
												v83 = v77 | int32(2)
												v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v83))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v83))) = uint16(v86)
												v88 = int32(4)
												v89 = v77 | v88
												v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v89))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v89))) = uint16(v92)
												v95 = v77 | int32(6)
												v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v95))) = uint16(v98)
												v101 = v66 + v88
												v103 = v70 + v88
												if v103 != v32&int32(2147483644) {
													v66 = v101
													v70 = v103
													continue
												} else {
													break
												}
												break
											}
											v105 = v101
										} else {
											v105 = v58
										}
										if v50 == int32(0) {
										} else {
											v117 = v105
											v122 = v58
											for {
												v127 = int32(1)
												v128 = v117 << (uint(v127) % 32)
												v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+v57))))
												*(*uint16)(unsafe.Add(mBase, uint32(v52+v128))) = uint16(v131)
												v136 = v122 + v127
												if v136 != v50 {
													v117 = v117 + v127
													v122 = v136
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
										v66 = v58
										v70 = int32(0)
										for {
											v77 = v66 << (uint(int32(1)) % 32)
											v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v77))) = uint16(v80)
											v83 = v77 | int32(2)
											v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v83))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v83))) = uint16(v86)
											v88 = int32(4)
											v89 = v77 | v88
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+v89))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v89))) = uint16(v92)
											v95 = v77 | int32(6)
											v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v95))) = uint16(v98)
											v101 = v66 + v88
											v103 = v70 + v88
											if v103 != v32&int32(2147483644) {
												v66 = v101
												v70 = v103
												continue
											} else {
												break
											}
											break
										}
										v105 = v101
									} else {
										v105 = v58
									}
									if v50 == int32(0) {
									} else {
										v117 = v105
										v122 = v58
										for {
											v127 = int32(1)
											v128 = v117 << (uint(v127) % 32)
											v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+v57))))
											*(*uint16)(unsafe.Add(mBase, uint32(v52+v128))) = uint16(v131)
											v136 = v122 + v127
											if v136 != v50 {
												v117 = v117 + v127
												v122 = v136
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
			v152 = m.ExcPending
			if v152 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(285025), int32(0))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525467), int32(954), int32(219214))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
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
