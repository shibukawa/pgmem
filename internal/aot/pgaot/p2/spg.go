package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_spgFormNodeTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 != 0 {
		v61 = int32(8)
		v63 = F_palloc0(m, v61)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			v67 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
			*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
			if l2 != 0 {
				v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
			} else {
				v73 = v61
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
			if l2 != 0 {
			} else {
				v76 = v63 + int32(8)
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
				if v77 == int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
				} else {
					v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
					if int32(0) < v81 {
						v109 = v81
					} else {
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						if v84 == int32(1) {
							v88 = int32(18)
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if v90 == v88 {
								v93 = v88
							} else {
								v93 = int32(2)
							}
							if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v100 = int32(6)
							} else {
								v100 = v93
							}
							v109 = v100
						} else {
							if v84&int32(1) != 0 {
								v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
							} else {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
							}
						}
					}
					if v109 == int32(0) {
					} else {
						base.MemoryCopy(m, v76, l1, v109)
					}
				}
			}
			m.G0 = v9 + int32(16)
			return v63
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
		if v12 != 0 {
			v61 = int32(16)
			v63 = F_palloc0(m, v61)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v67 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
				*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
				if l2 != 0 {
					v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
				} else {
					v73 = v61
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
				if l2 != 0 {
				} else {
					v76 = v63 + int32(8)
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
					if v77 == int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
					} else {
						v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
						if int32(0) < v81 {
							v109 = v81
						} else {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							if v84 == int32(1) {
								v88 = int32(18)
								v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								if v90 == v88 {
									v93 = v88
								} else {
									v93 = int32(2)
								}
								if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v100 = int32(6)
								} else {
									v100 = v93
								}
								v109 = v100
							} else {
								if v84&int32(1) != 0 {
									v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
								}
							}
						}
						if v109 == int32(0) {
						} else {
							base.MemoryCopy(m, v76, l1, v109)
						}
					}
				}
				m.G0 = v9 + int32(16)
				return v63
			}
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
			if v14 <= int32(0) {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v17 == int32(1) {
					v21 = int32(18)
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					if v23 == v21 {
						v26 = v21
					} else {
						v26 = int32(2)
					}
					if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v33 = int32(6)
					} else {
						v33 = v26
					}
					v41 = v33
					v61 = (v41+int32(7))&int32(248) + int32(8)
					v63 = F_palloc0(m, v61)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
						*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
						if l2 != 0 {
							v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
						} else {
							v73 = v61
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
						if l2 != 0 {
						} else {
							v76 = v63 + int32(8)
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v77 == int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
							} else {
								v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v81 {
									v109 = v81
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
									if v84 == int32(1) {
										v88 = int32(18)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
										if v90 == v88 {
											v93 = v88
										} else {
											v93 = int32(2)
										}
										if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v100 = int32(6)
										} else {
											v100 = v93
										}
										v109 = v100
									} else {
										if v84&int32(1) != 0 {
											v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
										} else {
											v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v109 == int32(0) {
								} else {
									base.MemoryCopy(m, v76, l1, v109)
								}
							}
						}
						m.G0 = v9 + int32(16)
						return v63
					}
				} else {
					if v17&int32(1) == int32(0) {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v52 = int32(base.Ui32(v48) >> (uint(int32(2)) % 32))
						v58 = (v52+int32(7))&int32(2147483640) + int32(8)
						if base.Ui32(int32(_a_F_spgFormNodeTuple_1)) <= base.Ui32(v52) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgFormNodeTuple_2)
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v58
									F_errmsg(m, int32(_a_F_spgFormNodeTuple_3), v9)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_spgFormNodeTuple_4), int32(979), int32(_a_F_spgFormNodeTuple_5))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
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
							v61 = v58
							v63 = F_palloc0(m, v61)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
								*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
								if l2 != 0 {
									v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
								} else {
									v73 = v61
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
								if l2 != 0 {
								} else {
									v76 = v63 + int32(8)
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
									if v77 == int32(1) {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
									} else {
										v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
										if int32(0) < v81 {
											v109 = v81
										} else {
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
											if v84 == int32(1) {
												v88 = int32(18)
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
												if v90 == v88 {
													v93 = v88
												} else {
													v93 = int32(2)
												}
												if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
													v100 = int32(6)
												} else {
													v100 = v93
												}
												v109 = v100
											} else {
												if v84&int32(1) != 0 {
													v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
												} else {
													v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
													v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
												}
											}
										}
										if v109 == int32(0) {
										} else {
											base.MemoryCopy(m, v76, l1, v109)
										}
									}
								}
								m.G0 = v9 + int32(16)
								return v63
							}
						}
					} else {
						v41 = int32(base.Ui32(v17) >> (uint(int32(1)) % 32))
						v61 = (v41+int32(7))&int32(248) + int32(8)
						v63 = F_palloc0(m, v61)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
							*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
							if l2 != 0 {
								v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
							} else {
								v73 = v61
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
							if l2 != 0 {
							} else {
								v76 = v63 + int32(8)
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
								if v77 == int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
								} else {
									v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
									if int32(0) < v81 {
										v109 = v81
									} else {
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
										if v84 == int32(1) {
											v88 = int32(18)
											v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
											if v90 == v88 {
												v93 = v88
											} else {
												v93 = int32(2)
											}
											if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v100 = int32(6)
											} else {
												v100 = v93
											}
											v109 = v100
										} else {
											if v84&int32(1) != 0 {
												v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
											} else {
												v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
											}
										}
									}
									if v109 == int32(0) {
									} else {
										base.MemoryCopy(m, v76, l1, v109)
									}
								}
							}
							m.G0 = v9 + int32(16)
							return v63
						}
					}
				}
			} else {
				v52 = v14
				v58 = (v52+int32(7))&int32(2147483640) + int32(8)
				if base.Ui32(int32(_a_F_spgFormNodeTuple_1)) <= base.Ui32(v52) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgFormNodeTuple_2)
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v58
							F_errmsg(m, int32(_a_F_spgFormNodeTuple_3), v9)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_spgFormNodeTuple_4), int32(979), int32(_a_F_spgFormNodeTuple_5))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
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
					v61 = v58
					v63 = F_palloc0(m, v61)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)) = uint16(v67)
						*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
						if l2 != 0 {
							v73 = v61 | int32(_a_F_spgFormNodeTuple_0)
						} else {
							v73 = v61
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v73)
						if l2 != 0 {
						} else {
							v76 = v63 + int32(8)
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v77 == int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v76))) = l1
							} else {
								v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v81 {
									v109 = v81
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
									if v84 == int32(1) {
										v88 = int32(18)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
										if v90 == v88 {
											v93 = v88
										} else {
											v93 = int32(2)
										}
										if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v100 = int32(6)
										} else {
											v100 = v93
										}
										v109 = v100
									} else {
										if v84&int32(1) != 0 {
											v109 = int32(base.Ui32(v84) >> (uint(int32(1)) % 32))
										} else {
											v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v109 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v109 == int32(0) {
								} else {
									base.MemoryCopy(m, v76, l1, v109)
								}
							}
						}
						m.G0 = v9 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func F_spg_box_quad_inner_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v130 float64
	_ = v130
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v384 float64
	_ = v384
	var v393 int32
	_ = v393
	var v395 float64
	_ = v395
	var v401 int32
	_ = v401
	var v403 float64
	_ = v403
	var v409 int32
	_ = v409
	var v411 float64
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 float64
	_ = v444
	var v445 float64
	_ = v445
	var v451 float64
	_ = v451
	var v452 float64
	_ = v452
	var v458 float64
	_ = v458
	var v459 float64
	_ = v459
	var v465 float64
	_ = v465
	var v466 float64
	_ = v466
	var v475 int32
	_ = v475
	var v476 float64
	_ = v476
	var v477 float64
	_ = v477
	var v483 float64
	_ = v483
	var v484 float64
	_ = v484
	var v490 float64
	_ = v490
	var v491 float64
	_ = v491
	var v497 float64
	_ = v497
	var v498 float64
	_ = v498
	var v507 int32
	_ = v507
	var v508 float64
	_ = v508
	var v510 float64
	_ = v510
	var v511 float64
	_ = v511
	var v515 float64
	_ = v515
	var v516 float64
	_ = v516
	var v522 float64
	_ = v522
	var v526 float64
	_ = v526
	var v532 float64
	_ = v532
	var v534 float64
	_ = v534
	var v535 float64
	_ = v535
	var v539 float64
	_ = v539
	var v540 float64
	_ = v540
	var v546 float64
	_ = v546
	var v550 float64
	_ = v550
	var v559 int32
	_ = v559
	var v560 float64
	_ = v560
	var v561 float64
	_ = v561
	var v567 float64
	_ = v567
	var v576 int32
	_ = v576
	var v577 float64
	_ = v577
	var v579 float64
	_ = v579
	var v580 float64
	_ = v580
	var v584 float64
	_ = v584
	var v591 int32
	_ = v591
	var v592 float64
	_ = v592
	var v594 float64
	_ = v594
	var v595 float64
	_ = v595
	var v599 float64
	_ = v599
	var v606 int32
	_ = v606
	var v607 float64
	_ = v607
	var v608 float64
	_ = v608
	var v614 float64
	_ = v614
	var v623 int32
	_ = v623
	var v624 float64
	_ = v624
	var v626 float64
	_ = v626
	var v627 float64
	_ = v627
	var v631 float64
	_ = v631
	var v638 int32
	_ = v638
	var v639 float64
	_ = v639
	var v640 float64
	_ = v640
	var v646 float64
	_ = v646
	var v655 int32
	_ = v655
	var v656 float64
	_ = v656
	var v657 float64
	_ = v657
	var v663 float64
	_ = v663
	var v670 int32
	_ = v670
	var v671 float64
	_ = v671
	var v673 float64
	_ = v673
	var v674 float64
	_ = v674
	var v678 float64
	_ = v678
	var v686 int32
	_ = v686
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 float64
	_ = v757
	var v758 float64
	_ = v758
	var v762 float64
	_ = v762
	var v768 float64
	_ = v768
	var v769 float64
	_ = v769
	var v770 float64
	_ = v770
	var v774 float64
	_ = v774
	var v780 float64
	_ = v780
	var v781 float64
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v23 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_palloc(m, int32(64))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v47 = v23
	goto L3
L3:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+33)))
	if v48 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	v31 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = v31
	v33 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v33
	v47 = v27
	goto L3
L6:
	;
	m.G0 = v19 + int32(32)
	return int32(0)
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v51
	v55 = F_palloc(m, v51<<(uint(int32(2))%32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v221 = F_palloc(m, int32(32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L44
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v58 <= int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v62 = int32(0)
	goto L12
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v78+v62<<(uint(int32(2))%32)))) = v62
	v84 = v62 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v84 < v85 {
		v62 = v84
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v85 <= int32(0) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v89 <= int32(0) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v94 = F_palloc(m, v89<<(uint(int32(3))%32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v97 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v101 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v175 = F_palloc(m, v172<<(uint(int32(2))%32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L35
	}
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v101*int32(48))+44))
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v47)))
	if base.F64_lt(v125, v126) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v124)+8))
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v47)+32))
	if base.F64_lt(v137, v138) != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v136 = base.F64_sub(v126, v125)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v47)+24))
	if base.F64_gt(v125, v130) == int32(0) {
		v136 = float64(0)
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v136 = base.F64_sub(v125, v130)
	goto L23
L28:
	;
	v149 = F_pg_hypot(m, v136, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L33
	}
L29:
	;
	v148 = base.F64_sub(v138, v137)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v47)+56))
	if base.F64_gt(v137, v142) == int32(0) {
		v148 = float64(0)
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v148 = base.F64_sub(v137, v142)
	goto L28
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v94+v101<<(uint(int32(3))%32)))) = v149
	v153 = v101 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v153 < v154 {
		v101 = v153
		goto L21
	} else {
		goto L34
	}
L34:
	;
	goto L22
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v94
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v179 < int32(2) {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v185 = int32(1)
	goto L37
L37:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v201 = F_palloc(m, v198<<(uint(int32(3))%32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	goto L6
L39:
	;
	v204 = v185 << (uint(int32(2)) % 32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v205))) = v201
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v210 = v208 << (uint(int32(3)) % 32)
	if v210 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211+v204)))
	base.MemoryCopy(m, v213, v94, v210)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v216 = v185 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v216 < v217 {
		v185 = v216
		goto L37
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	v223 = *(*float64)(unsafe.Add(mBase, uint32(v219)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v221))) = v223
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v219)))
	*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = v225
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v219)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v221)+16)) = v227
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v219)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v221)+24)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v234 = F_palloc(m, v231<<(uint(int32(2))%32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v236 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v245 = v4
	goto L49
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v323 = F_palloc(m, v320<<(uint(int32(2))%32))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L61
	}
L49:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v258 = v255 + v245*int32(48)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	switch v259 - int32(603) {
	case 0:
		goto L52
	case 1:
		goto L54
	default:
		goto L53
	}
L50:
	;
	goto L48
L51:
	;
	v284 = F_palloc(m, int32(32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L59
	}
L52:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v258)+44))
	v282 = v281
	goto L51
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)+44))
	v263 = F_pg_detoast_datum(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v282 = v263 + int32(8)
	goto L51
L56:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v271
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_inner_consistent_0), v19)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_inner_consistent_1), int32(544), int32(_a_F_spg_box_quad_inner_consistent_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v282)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v284))) = v286
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v282)))
	*(*float64)(unsafe.Add(mBase, uint32(v284)+8)) = v288
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v282)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v284)+16)) = v290
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v282)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v284)+24)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v234+v245<<(uint(int32(2))%32)))) = v284
	v299 = v245 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v299 < v300 {
		v245 = v299
		goto L49
	} else {
		goto L60
	}
L60:
	;
	goto L50
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v329 = F_palloc(m, v326<<(uint(int32(2))%32))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v329
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v332 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v338 = F_palloc(m, v335<<(uint(int32(2))%32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v341 = int32(_a_F_spg_box_quad_inner_consistent_3)
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0])) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if int32(0) < v346 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v338
	goto L65
L67:
	;
	v359 = v4
	v364 = v4
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0])) = v342
	goto L6
L70:
	;
	v366 = F_palloc(m, int32(64))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v47)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+56)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v47)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+48)) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v47)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+40)) = v372
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+32)) = v374
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v47)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+24)) = v376
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+16)) = v378
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+8)) = v380
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v366))) = v382
	v384 = *(*float64)(unsafe.Add(mBase, uint32(v221)))
	if v359&int32(8) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v359&int32(4) != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v366))) = v384
	goto L73
L75:
	;
	goto L76
L76:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v366)+8)) = v384
	goto L73
L77:
	;
	v393 = int32(16)
	goto L79
L78:
	;
	v393 = int32(24)
	goto L79
L79:
	;
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v221)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v366+v393))) = v395
	if v359&int32(2) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v401 = int32(32)
	goto L82
L81:
	;
	v401 = int32(40)
	goto L82
L82:
	;
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v221)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v366+v401))) = v403
	if v359&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v409 = int32(48)
	goto L85
L84:
	;
	v409 = int32(56)
	goto L85
L85:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v221)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v366+v409))) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v413 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v845 = v364 + int32(1)
	v847 = v845 & int32(255)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v847 < v848 {
		v359 = v847
		v364 = v845
		goto L70
	} else {
		goto L161
	}
L87:
	;
	F_pfree(m, v366)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L160
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L157
	}
L89:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v421 = int32(0)
	goto L92
L90:
	;
	goto L91
L91:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v706 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v704+v705<<(uint(v706)%32)))) = v366
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v710+v711<<(uint(v706)%32)))) = v359
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v716 <= int32(0) {
		goto L139
	} else {
		goto L140
	}
L92:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416+v421*int32(48))+6)))
	switch v437 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L101
	case 2:
		goto L105
	case 3:
		goto L99
	case 4:
		goto L100
	case 5, 7:
		goto L103
	case 6:
		goto L104
	case 8:
		goto L95
	case 9:
		goto L96
	case 10:
		goto L98
	case 11:
		goto L97
	default:
		goto L88
	}
L93:
	;
	goto L91
L94:
	;
	v686 = v421 + int32(1)
	if v686 != v413 {
		v421 = v686
		goto L92
	} else {
		goto L138
	}
L95:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v671 = *(*float64)(unsafe.Add(mBase, uint32(v670)+24))
	v673 = base.F64_add(v671, float64(1e-06))
	v674 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	if base.F64_ge(v673, v674) == int32(0) {
		goto L87
	} else {
		goto L136
	}
L96:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v656 = *(*float64)(unsafe.Add(mBase, uint32(v655)+16))
	v657 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	if base.F64_gt(v656, base.F64_add(v657, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L134
	}
L97:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v639 = *(*float64)(unsafe.Add(mBase, uint32(v638)+16))
	v640 = *(*float64)(unsafe.Add(mBase, uint32(v366)+40))
	if base.F64_le(v639, base.F64_add(v640, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L132
	}
L98:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v624 = *(*float64)(unsafe.Add(mBase, uint32(v623)+24))
	v626 = base.F64_add(v624, float64(1e-06))
	v627 = *(*float64)(unsafe.Add(mBase, uint32(v366)+40))
	if base.F64_lt(v626, v627) == int32(0) {
		goto L87
	} else {
		goto L130
	}
L99:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v607 = *(*float64)(unsafe.Add(mBase, uint32(v606)))
	v608 = *(*float64)(unsafe.Add(mBase, uint32(v366)+8))
	if base.F64_le(v607, base.F64_add(v608, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L128
	}
L100:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v592 = *(*float64)(unsafe.Add(mBase, uint32(v591)+8))
	v594 = base.F64_add(v592, float64(1e-06))
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v366)+8))
	if base.F64_lt(v594, v595) == int32(0) {
		goto L87
	} else {
		goto L126
	}
L101:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v576)+8))
	v579 = base.F64_add(v577, float64(1e-06))
	v580 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	if base.F64_ge(v579, v580) == int32(0) {
		goto L87
	} else {
		goto L124
	}
L102:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v560 = *(*float64)(unsafe.Add(mBase, uint32(v559)))
	v561 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	if base.F64_gt(v560, base.F64_add(v561, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L122
	}
L103:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v508 = *(*float64)(unsafe.Add(mBase, uint32(v507)+8))
	v510 = base.F64_add(v508, float64(1e-06))
	v511 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	if base.F64_ge(v510, v511) == int32(0) {
		goto L87
	} else {
		goto L114
	}
L104:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v476 = *(*float64)(unsafe.Add(mBase, uint32(v475)+8))
	v477 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_le(v476, base.F64_add(v477, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L110
	}
L105:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v234+v421<<(uint(int32(2))%32))))
	v444 = *(*float64)(unsafe.Add(mBase, uint32(v443)))
	v445 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_le(v444, base.F64_add(v445, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L106
	}
L106:
	;
	v451 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v443)+8))
	if base.F64_le(v451, base.F64_add(v452, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L107
	}
L107:
	;
	v458 = *(*float64)(unsafe.Add(mBase, uint32(v443)+16))
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_le(v458, base.F64_add(v459, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L108
	}
L108:
	;
	v465 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v443)+24))
	if base.F64_le(v465, base.F64_add(v466, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L109
	}
L109:
	;
	goto L94
L110:
	;
	v483 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	v484 = *(*float64)(unsafe.Add(mBase, uint32(v475)))
	if base.F64_le(v483, base.F64_add(v484, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L111
	}
L111:
	;
	v490 = *(*float64)(unsafe.Add(mBase, uint32(v475)+24))
	v491 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_le(v490, base.F64_add(v491, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L112
	}
L112:
	;
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	v498 = *(*float64)(unsafe.Add(mBase, uint32(v475)+16))
	if base.F64_le(v497, base.F64_add(v498, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L113
	}
L113:
	;
	goto L94
L114:
	;
	v515 = *(*float64)(unsafe.Add(mBase, uint32(v507)))
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v366)+8))
	if base.F64_le(v515, base.F64_add(v516, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L115
	}
L115:
	;
	v522 = *(*float64)(unsafe.Add(mBase, uint32(v366)+16))
	if base.F64_le(v522, v510) == int32(0) {
		goto L87
	} else {
		goto L116
	}
L116:
	;
	v526 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_ge(base.F64_add(v526, float64(1e-06)), v515) == int32(0) {
		goto L87
	} else {
		goto L117
	}
L117:
	;
	v532 = *(*float64)(unsafe.Add(mBase, uint32(v507)+24))
	v534 = base.F64_add(v532, float64(1e-06))
	v535 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	if base.F64_ge(v534, v535) == int32(0) {
		goto L87
	} else {
		goto L118
	}
L118:
	;
	v539 = *(*float64)(unsafe.Add(mBase, uint32(v507)+16))
	v540 = *(*float64)(unsafe.Add(mBase, uint32(v366)+40))
	if base.F64_le(v539, base.F64_add(v540, float64(1e-06))) == int32(0) {
		goto L87
	} else {
		goto L119
	}
L119:
	;
	v546 = *(*float64)(unsafe.Add(mBase, uint32(v366)+48))
	if base.F64_le(v546, v534) == int32(0) {
		goto L87
	} else {
		goto L120
	}
L120:
	;
	v550 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_ge(base.F64_add(v550, float64(1e-06)), v539) == int32(0) {
		goto L87
	} else {
		goto L121
	}
L121:
	;
	goto L94
L122:
	;
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v366)+16))
	if base.F64_lt(base.F64_add(v567, float64(1e-06)), v560) == int32(0) {
		goto L87
	} else {
		goto L123
	}
L123:
	;
	goto L94
L124:
	;
	v584 = *(*float64)(unsafe.Add(mBase, uint32(v366)+16))
	if base.F64_le(v584, v579) == int32(0) {
		goto L87
	} else {
		goto L125
	}
L125:
	;
	goto L94
L126:
	;
	v599 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_gt(v599, v594) == int32(0) {
		goto L87
	} else {
		goto L127
	}
L127:
	;
	goto L94
L128:
	;
	v614 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_ge(base.F64_add(v614, float64(1e-06)), v607) == int32(0) {
		goto L87
	} else {
		goto L129
	}
L129:
	;
	goto L94
L130:
	;
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_gt(v631, v626) == int32(0) {
		goto L87
	} else {
		goto L131
	}
L131:
	;
	goto L94
L132:
	;
	v646 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_ge(base.F64_add(v646, float64(1e-06)), v639) == int32(0) {
		goto L87
	} else {
		goto L133
	}
L133:
	;
	goto L94
L134:
	;
	v663 = *(*float64)(unsafe.Add(mBase, uint32(v366)+48))
	if base.F64_lt(base.F64_add(v663, float64(1e-06)), v656) != 0 {
		goto L94
	} else {
		goto L135
	}
L135:
	;
	goto L87
L136:
	;
	v678 = *(*float64)(unsafe.Add(mBase, uint32(v366)+48))
	if base.F64_le(v678, v673) == int32(0) {
		goto L87
	} else {
		goto L137
	}
L137:
	;
	goto L94
L138:
	;
	goto L93
L139:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v804 + int32(1)
	goto L86
L140:
	;
	v721 = F_palloc(m, v716<<(uint(int32(3))%32))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v723+v724<<(uint(int32(2))%32)))) = v721
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v729 <= int32(0) {
		goto L139
	} else {
		goto L142
	}
L142:
	;
	v736 = int32(0)
	goto L143
L143:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v752+v736*int32(48))+44))
	v757 = *(*float64)(unsafe.Add(mBase, uint32(v756)))
	v758 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	if base.F64_lt(v757, v758) != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L139
L145:
	;
	v769 = *(*float64)(unsafe.Add(mBase, uint32(v756)+8))
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	if base.F64_lt(v769, v770) != 0 {
		goto L151
	} else {
		goto L152
	}
L146:
	;
	v768 = base.F64_sub(v758, v757)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v762 = *(*float64)(unsafe.Add(mBase, uint32(v366)+24))
	if base.F64_gt(v757, v762) == int32(0) {
		v768 = float64(0)
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v768 = base.F64_sub(v757, v762)
	goto L145
L150:
	;
	v781 = F_pg_hypot(m, v768, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L155
	}
L151:
	;
	v780 = base.F64_sub(v770, v769)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v774 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_gt(v769, v774) == int32(0) {
		v780 = float64(0)
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v780 = base.F64_sub(v769, v774)
	goto L150
L155:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v721+v736<<(uint(int32(3))%32)))) = v781
	v785 = v736 + int32(1)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v785 < v786 {
		v736 = v785
		goto L143
	} else {
		goto L156
	}
L156:
	;
	goto L144
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v437
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_inner_consistent_4), v19+int32(16))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_inner_consistent_1), int32(691), int32(_a_F_spg_box_quad_inner_consistent_5))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	goto L86
L161:
	;
	goto L71
}
func F_spg_box_quad_leaf_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v2)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)))
	if v18 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
	goto L3
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if int32(0) < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(32)
	return v181
L5:
	;
	v32 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v168 = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v169 <= int32(0) {
		v181 = v168
		goto L4
	} else {
		goto L67
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v36 = v33 + v32*int32(48)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	switch v38 - int32(603) {
	case 0:
		goto L11
	case 1:
		goto L13
	default:
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	switch v37 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L38
	case 3:
		goto L32
	case 4:
		goto L33
	case 5:
		goto L36
	case 6:
		goto L26
	case 7:
		goto L37
	case 8:
		goto L28
	case 9:
		goto L29
	case 10:
		goto L31
	case 11:
		goto L30
	default:
		goto L27
	}
L11:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v36)+44))
	v75 = v74
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L22
	}
L13:
	;
	if int32(1)<<(uint(v37)%32)&int32(_a_F_spg_box_quad_leaf_consistent_0) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = base.B2i32(base.Ui32(v37) <= base.Ui32(int32(12)))
	goto L16
L15:
	;
	v48 = int32(0)
	goto L16
L16:
	;
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v51)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v36)+44))
	v54 = F_pg_detoast_datum(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v75 = v54 + int32(8)
	goto L10
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_leaf_consistent_1), v11)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_leaf_consistent_2), int32(544), int32(_a_F_spg_box_quad_leaf_consistent_3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L20
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
	v157 = v32 + int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v157 < v158 {
		v32 = v157
		goto L8
	} else {
		goto L66
	}
L26:
	;
	v148 = int32(0)
	v151 = F_DirectFunctionCall2Coll(m, int32(99), v148, v14, v75)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L64
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L20
	} else {
		goto L61
	}
L28:
	;
	v128 = int32(0)
	v131 = F_DirectFunctionCall2Coll(m, int32(102), v128, v14, v75)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L59
	}
L29:
	;
	v123 = int32(0)
	v126 = F_DirectFunctionCall2Coll(m, int32(103), v123, v14, v75)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L57
	}
L30:
	;
	v118 = int32(0)
	v121 = F_DirectFunctionCall2Coll(m, int32(101), v118, v14, v75)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L55
	}
L31:
	;
	v113 = int32(0)
	v116 = F_DirectFunctionCall2Coll(m, int32(100), v113, v14, v75)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L53
	}
L32:
	;
	v108 = int32(0)
	v111 = F_DirectFunctionCall2Coll(m, int32(104), v108, v14, v75)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L20
	} else {
		goto L51
	}
L33:
	;
	v103 = int32(0)
	v106 = F_DirectFunctionCall2Coll(m, int32(95), v103, v14, v75)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L49
	}
L34:
	;
	v98 = int32(0)
	v101 = F_DirectFunctionCall2Coll(m, int32(98), v98, v14, v75)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L47
	}
L35:
	;
	v93 = int32(0)
	v96 = F_DirectFunctionCall2Coll(m, int32(97), v93, v14, v75)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L45
	}
L36:
	;
	v88 = int32(0)
	v91 = F_DirectFunctionCall2Coll(m, int32(117), v88, v14, v75)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L43
	}
L37:
	;
	v83 = int32(0)
	v86 = F_DirectFunctionCall2Coll(m, int32(118), v83, v14, v75)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L41
	}
L38:
	;
	v78 = int32(0)
	v81 = F_DirectFunctionCall2Coll(m, int32(96), v78, v14, v75)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v81 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v181 = v78
	goto L4
L41:
	;
	if v86 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v181 = v83
	goto L4
L43:
	;
	if v91 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v181 = v88
	goto L4
L45:
	;
	if v96 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v181 = v93
	goto L4
L47:
	;
	if v101 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v181 = v98
	goto L4
L49:
	;
	if v106 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v181 = v103
	goto L4
L51:
	;
	if v111 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	v181 = v108
	goto L4
L53:
	;
	if v116 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v181 = v113
	goto L4
L55:
	;
	if v121 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v181 = v118
	goto L4
L57:
	;
	if v126 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	v181 = v123
	goto L4
L59:
	;
	if v131 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	v181 = v128
	goto L4
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v37
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_leaf_consistent_4), v11+int32(16))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_leaf_consistent_2), int32(831), int32(_a_F_spg_box_quad_leaf_consistent_5))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
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
	if v151 == int32(0) {
		v181 = v148
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L25
L66:
	;
	goto L9
L67:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	v175 = F_spg_key_orderbys_distances(m, v14, int32(0), v172, v169)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(base.B2i32(v173 == int32(3292)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v175
	v181 = v168
	goto L4
}
func F_spg_kd_choose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 float64
	_ = v39
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
	if v7 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_spg_kd_choose_0), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_spg_kd_choose_1), int32(62), int32(_a_F_spg_kd_choose_2))
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
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v29 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v39 = *(*float64)(unsafe.Add(mBase, uint32(v27+(v31^int32(-1))<<(uint(int32(3))%32)&int32(8))))
		*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v29
		v44 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = base.B2i32(base.F64_gt(v26, v39) == v44)
		return v44
	}
}
func F_spg_kd_inner_consistent(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 float64
	_ = v72
	var v80 int32
	_ = v80
	var v85 float64
	_ = v85
	var v87 int32
	_ = v87
	var v92 float64
	_ = v92
	var v101 float64
	_ = v101
	var v108 int32
	_ = v108
	var v111 float64
	_ = v111
	var v119 int32
	_ = v119
	var v122 float64
	_ = v122
	var v126 int32
	_ = v126
	var v131 float64
	_ = v131
	var v135 float64
	_ = v135
	var v141 float64
	_ = v141
	var v145 float64
	_ = v145
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 float64
	_ = v248
	var v249 float64
	_ = v249
	var v252 float64
	_ = v252
	var v254 float64
	_ = v254
	var v258 float64
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
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
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+33)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v194 = F_palloc(m, int32(8))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L51
	}
L3:
	;
	v44 = base.F64_add(v23, float64(1e-06))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v48 = int32(0)
	v53 = int32(6)
	goto L12
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(0) < v24 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v184 = int32(6)
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	F_errmsg_internal(m, int32(_a_F_spg_kd_inner_consistent_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_spg_kd_inner_consistent_1), int32(173), int32(_a_F_spg_kd_inner_consistent_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
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
	v62 = v45 + v48*int32(48)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	switch v64 - int32(1) {
	case 0:
		goto L23
	default:
		goto L15
	case 4:
		goto L22
	case 5:
		goto L21
	case 7:
		goto L18
	case 9, 28:
		goto L20
	case 10, 29:
		goto L19
	}
L13:
	;
	v184 = v174
	goto L2
L14:
	;
	v177 = v48 + int32(1)
	if v177 != v24 {
		v48 = v177
		v53 = v174
		goto L12
	} else {
		goto L50
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L47
	}
L16:
	;
	if v153 != 0 {
		v174 = v153
		goto L14
	} else {
		goto L46
	}
L17:
	;
	v153 = v53 & int32(4)
	goto L16
L18:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v126&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v119&int32(1) != 0 {
		v174 = v53
		goto L14
	} else {
		goto L37
	}
L20:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v108&int32(1) != 0 {
		v174 = v53
		goto L14
	} else {
		goto L35
	}
L21:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v87&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v80&int32(1) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v67&int32(1) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_lt(base.F64_add(v72, float64(1e-06)), v23) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v153 = v53 & int32(2)
	goto L16
L26:
	;
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(v85, v44) != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v174 = v53
	goto L14
L28:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(v23, base.F64_add(v92, float64(1e-06))) != 0 {
		v153 = v53 & int32(2)
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	if base.F64_gt(v23, base.F64_add(v101, float64(1e-06))) != 0 {
		v153 = v53 & int32(2)
		goto L16
	} else {
		goto L33
	}
L31:
	;
	if base.F64_lt(v44, v92) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L32
	}
L32:
	;
	goto L17
L33:
	;
	if base.F64_lt(v44, v101) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L34
	}
L34:
	;
	goto L17
L35:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	if base.F64_lt(base.F64_add(v111, float64(1e-06)), v23) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v153 = v53 & int32(2)
	goto L16
L37:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	if base.F64_gt(v122, v44) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L17
L39:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(v23, base.F64_add(v131, float64(1e-06))) != 0 {
		v153 = v53 & int32(2)
		goto L16
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	if base.F64_gt(v23, base.F64_add(v141, float64(1e-06))) != 0 {
		v153 = v53 & int32(2)
		goto L16
	} else {
		goto L44
	}
L42:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v63)+16))
	if base.F64_gt(v135, v44) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L43
	}
L43:
	;
	goto L17
L44:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v63)+24))
	if base.F64_gt(v145, v44) == int32(0) {
		v174 = v53
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L17
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	goto L1
L47:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160+v48*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v164
	F_errmsg_internal(m, int32(_a_F_spg_kd_inner_consistent_3), v15)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_spg_kd_inner_consistent_1), int32(247), int32(_a_F_spg_kd_inner_consistent_2))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	goto L13
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v197 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v203 = F_palloc(m, v200<<(uint(int32(2))%32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v184&int32(2) != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v203
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v209 = F_palloc(m, v206<<(uint(int32(2))%32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v212 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v258
	goto L54
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+80)) = v23
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v231)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v252
	v254 = *(*float64)(unsafe.Add(mBase, uint32(v231)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = v254
	v258 = v23
	goto L57
L59:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v23
	*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = v23
	v248 = *(*float64)(unsafe.Add(mBase, uint32(v244)))
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v245)))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+80)) = v249
	v258 = v248
	goto L57
L60:
	;
	v215 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v215
	v217 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v215
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v215
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v217
	v244 = v15 + int32(24)
	v245 = v15 + int32(16)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v231)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v231)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v238
	if v212&int32(1) != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v244 = v231
	v245 = v231 + int32(16)
	goto L59
L64:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v265+v266<<(uint(int32(2))%32)))) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v270 < v272 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v184&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v275 = int32(_a_F_spg_kd_inner_consistent_4)
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v278
	v282 = F_box_copy(m, v15+int32(32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v305 + int32(1)
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v276
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v286+v287<<(uint(int32(2))%32)))) = v282
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v295 = F_spg_key_orderbys_distances(m, v282, int32(0), v293, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v297+v298<<(uint(int32(2))%32)))) = v295
	goto L69
L72:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v313+v314<<(uint(int32(2))%32)))) = int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v320 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v360 = F_palloc(m, int32(8))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L80
	}
L75:
	;
	v323 = int32(_a_F_spg_kd_inner_consistent_4)
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0]))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v326
	v330 = F_box_copy(m, v15-int32(-64))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L8
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v353 + int32(1)
	goto L74
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v324
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v334+v335<<(uint(int32(2))%32)))) = v330
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v343 = F_spg_key_orderbys_distances(m, v330, int32(0), v341, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(2))%32)))) = v343
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v360
	v363 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v363
	goto L1
}
func F_spg_key_orderbys_distances(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v48 float64
	_ = v48
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	var v69 float64
	_ = v69
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v89 float64
	_ = v89
	var v99 int32
	_ = v99
	v16 = F_palloc(m, l3<<(uint(int32(3))%32))
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
	if int32(0) < l3 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = l2
	v31 = v16
	v34 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v16
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v31))) = v89
	v99 = v34 + int32(1)
	if v99 != l3 {
		v24 = v24 + int32(48)
		v31 = v31 + int32(8)
		v34 = v99
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v38 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), v35, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v41 = math.Float64frombits(uint64(0x7ff8000000000000))
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v35)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v42)&int64(9223372036854775807)) {
		v89 = v41
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v89 = v40
	goto L8
L13:
	;
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) {
		v89 = v41
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v35)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) {
		v89 = v41
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
		v89 = v41
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if base.F64_lt(v42, v48) != 0 {
		v75 = base.F64_sub(v48, v42)
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if base.F64_lt(v54, v60) != 0 {
		v85 = base.F64_sub(v60, v54)
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_gt(v42, v69) == int32(0) {
		v75 = float64(0)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v75 = base.F64_sub(v42, v69)
	goto L17
L20:
	;
	v86 = F_pg_hypot(m, v75, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F64_gt(v54, v79) == int32(0) {
		v85 = float64(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v85 = base.F64_sub(v54, v79)
	goto L20
L23:
	;
	v89 = v86
	goto L8
L24:
	;
	goto L7
}
func F_spg_range_quad_leaf_consistent(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v15 = F_pg_detoast_datum(m, v14)
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
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v19)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v24 = F_range_get_typcache(m, l0, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v27 <= int32(0) {
		v116 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v10 + int32(16)
	return v116
L5:
	;
	v35 = int32(0)
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v40 = v37 + v35*int32(48)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+44))
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)))
	switch v42 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L9
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	case 5:
		goto L15
	case 6:
		goto L14
	case 7:
		goto L13
	default:
		goto L10
	case 15:
		goto L12
	case 17:
		goto L11
	}
L7:
	;
	v116 = v111
	goto L4
L8:
	;
	v111 = int32(1)
	v113 = v35 + v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v113 < v114 {
		v35 = v113
		goto L6
	} else {
		goto L52
	}
L9:
	;
	v106 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L49
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L46
	}
L11:
	;
	v83 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L43
	}
L12:
	;
	v80 = F_range_contains_elem_internal(m, v24, v15, v41)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L41
	}
L13:
	;
	v75 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L38
	}
L14:
	;
	v70 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L35
	}
L15:
	;
	v65 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v60 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v55 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v50 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v45 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v47 = F_range_before_internal(m, v24, v15, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v47 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v116 = int32(0)
	goto L4
L23:
	;
	v52 = F_range_overlaps_internal(m, v24, v15, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v52 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v116 = int32(0)
	goto L4
L26:
	;
	v57 = F_range_overright_internal(m, v24, v15, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v57 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v116 = int32(0)
	goto L4
L29:
	;
	v62 = F_range_after_internal(m, v24, v15, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v62 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v116 = int32(0)
	goto L4
L32:
	;
	v67 = F_range_adjacent_internal(m, v24, v15, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v67 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v116 = int32(0)
	goto L4
L35:
	;
	v72 = F_range_contains_internal(m, v24, v15, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v72 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v116 = int32(0)
	goto L4
L38:
	;
	v77 = F_range_contained_by_internal(m, v24, v15, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v77 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v116 = int32(0)
	goto L4
L41:
	;
	if v80 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v116 = int32(0)
	goto L4
L43:
	;
	v85 = F_range_eq_internal(m, v24, v15, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v85 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v116 = int32(0)
	goto L4
L46:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+v35*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v96
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_leaf_consistent_0), v10)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_leaf_consistent_1), int32(985), int32(_a_F_spg_range_quad_leaf_consistent_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
	v108 = F_range_overleft_internal(m, v24, v15, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v108 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v116 = int32(0)
	goto L4
L52:
	;
	goto L7
}
func F_spg_range_quad_picksplit(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
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
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = F_pg_detoast_datum(m, v20)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_range_get_typcache(m, l0, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v31 = F_palloc(m, v28<<(uint(int32(3))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v36 = F_palloc(m, v33<<(uint(int32(3))%32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if int32(0) < v38 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v15 + int32(48)
	return int32(0)
L7:
	;
	F_qsort_arg(m, v31, v72, int32(8), int32(1473), v26)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	v42 = int32(0)
	v47 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v89
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v89)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(8589934592)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v99 = F_palloc(m, v96<<(uint(int32(2))%32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v47<<(uint(int32(2))%32))))
	v59 = F_pg_detoast_datum(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	if v72 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v62 = v42 << (uint(int32(3)) % 32)
	F_range_deserialize(m, v26, v59, v31+v62, v62+v36, v15+int32(10))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)))
	v70 = int32(1)
	v72 = v42 + (v69 ^ v70)
	v74 = v47 + v70
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v74 < v75 {
		v42 = v72
		v47 = v74
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v105 = F_palloc(m, v102<<(uint(int32(2))%32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v108 <= int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v111 = v89
	goto L20
L20:
	;
	v124 = v111 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124+v125)))
	v128 = F_pg_detoast_datum(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L6
L22:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v130+v124))) = v128
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v133+v124))) = int32(0)
	v138 = v111 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v138 < v139 {
		v111 = v138
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_qsort_arg(m, v36, v72, int32(8), int32(1473), v26)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v149 = int32(0)
	v151 = base.I32_div_s(v72, int32(2))
	v153 = v151 << (uint(int32(3)) % 32)
	v158 = F_range_serialize(m, v26, v31+v153, v153+v36, v149, v149)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v158
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	if v163 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = int32(4)
	goto L29
L28:
	;
	v168 = int32(5)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v173 = F_palloc(m, v170<<(uint(int32(2))%32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v173
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v179 = F_palloc(m, v176<<(uint(int32(2))%32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v182 <= int32(0) {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v190 = v149
	goto L33
L33:
	;
	v198 = v190 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v198+v199)))
	v202 = F_pg_detoast_datum(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L6
L35:
	;
	v205 = v15 + int32(40)
	v207 = v15 + int32(32)
	F_range_deserialize(m, v26, v158, v205, v207, v15+int32(31))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v213 = v15 + int32(20)
	v215 = v15 + int32(12)
	F_range_deserialize(m, v26, v202, v213, v215, v15+int32(11))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+11)))
	if v221 != 0 {
		v240 = int32(5)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v198))) = v202
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v246 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v244+v198))) = v240 - v246
	v250 = v190 + v246
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v250 < v251 {
		v190 = v250
		goto L33
	} else {
		goto L49
	}
L39:
	;
	v222 = F_range_cmp_bounds(m, v26, v213, v205)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v226 = F_range_cmp_bounds(m, v26, v215, v207)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v226 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v230 = int32(1)
	goto L44
L43:
	;
	v230 = int32(2)
	goto L44
L44:
	;
	if int32(0) <= v222 {
		v240 = v230
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if int32(0) <= v226 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v237 = int32(4)
	goto L48
L47:
	;
	v237 = int32(3)
	goto L48
L48:
	;
	v240 = v237
	goto L38
L49:
	;
	goto L34
}
func F_spg_text_inner_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v303 int32
	_ = v303
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = F_pg_newlocale_from_collation(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v29 = int32(1)
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+34)))
	if v34 == v29 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v38 = F_pg_detoast_datum_packed(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v71 = int32(0)
	v72 = v2
	v73 = v30
	goto L5
L5:
	;
	v75 = v73 + int32(4)
	v76 = F_palloc(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v71 = v69
	v72 = v38
	v73 = v30 + v69
	goto L5
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v40 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v46 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v57 = int32(1)
	if v40&v57 != 0 {
		v69 = int32(base.Ui32(v40)>>(uint(v57)%32)) - v57
		goto L6
	} else {
		goto L17
	}
L11:
	;
	v49 = int32(16)
	goto L13
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = int32(4)
	goto L16
L15:
	;
	v56 = v49
	goto L16
L16:
	;
	v69 = v56
	goto L6
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v69 = int32(base.Ui32(v63)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v75 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v82 = int32(0)
	v83 = base.B2i32(v81 == v82)
	if v83|v83 == v82 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = int32(4)
	base.MemoryCopy(m, v76+v89, v31+v89, v81)
	goto L21
L20:
	;
	goto L21
L21:
	;
	if v71 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v96 = int32(4)
	v98 = int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v100&v98 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v109 = F_palloc(m, v106<<(uint(int32(2))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	v103 = v98
	goto L27
L26:
	;
	v103 = v96
	goto L27
L27:
	;
	base.MemoryCopy(m, v76+v94+v96, v72+v103, v71)
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v115 = F_palloc(m, v112<<(uint(int32(2))%32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v121 = F_palloc(m, v118<<(uint(int32(2))%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v121
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v123 < v126 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v130 = v76 + int32(4)
	v131 = int32(1)
	v150 = v2
	goto L34
L32:
	;
	goto L33
L33:
	;
	m.G0 = v19 + int32(16)
	return int32(0)
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v150<<(uint(int32(2))%32))))
	if int32(0) < base.I32_extend16_s(v158) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v73+v76+int32(3)))) = uint8(v158)
	v163 = v73
	goto L38
L37:
	;
	v163 = v73 - v131
	goto L38
L38:
	;
	v164 = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v164 < v165 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v406 = v150 + int32(1)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v406 < v407 {
		v150 = v406
		goto L34
	} else {
		goto L101
	}
L40:
	;
	v168 = v164
	goto L43
L41:
	;
	goto L42
L42:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v358 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v356+v357<<(uint(v358)%32)))) = v150
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v362+v363<<(uint(v358)%32)))) = v163 - v367
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v163<<(uint(v358)%32) + int32(16)
	v377 = F_datumCopy(m, v76, int32(0), int32(-1))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L100
	}
L43:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v187 = v184 + v168*int32(48)
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+6)))
	if base.B2i32(base.Ui32(v188) < base.Ui32(int32(11)))|base.B2i32(v188 == int32(28)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L42
L45:
	;
	v337 = v168 + int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v337 < v338 {
		v168 = v337
		goto L43
	} else {
		goto L99
	}
L46:
	;
	if v32&v131 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v200 = v188
	goto L48
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v187)+44))
	v202 = F_pg_detoast_datum_packed(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v200 = v188 - int32(10)
	goto L48
L50:
	;
	v234 = int32(1)
	if v204&v234 != 0 {
		goto L62
	} else {
		goto L63
	}
L51:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v204 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v210 == int32(18) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v221 = int32(1)
	if v204&v221 != 0 {
		v233 = int32(base.Ui32(v204)>>(uint(v221)%32)) - v221
		goto L50
	} else {
		goto L61
	}
L55:
	;
	v213 = int32(16)
	goto L57
L56:
	;
	v213 = int32(0)
	goto L57
L57:
	;
	if base.Ui32((v210-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v220 = int32(4)
	goto L60
L59:
	;
	v220 = v213
	goto L60
L60:
	;
	v233 = v220
	goto L50
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v233 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(4)
	goto L50
L62:
	;
	v238 = v234
	goto L64
L63:
	;
	v238 = int32(4)
	goto L64
L64:
	;
	v239 = v202 + v238
	v240 = base.B2i32(v233 < v163)
	if v233 < v163 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v241 = v233
	goto L67
L66:
	;
	v241 = v163
	goto L67
L67:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v241) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	switch v200&int32(_a_F_spg_text_inner_consistent_0) - int32(1) {
	case 0, 1:
		goto L90
	case 2:
		goto L89
	case 3, 4:
		goto L88
	default:
		goto L87
	case 27:
		goto L86
	}
L69:
	;
	v303 = int32(0)
	goto L68
L70:
	;
	v277 = v272
	v278 = v273
	v279 = v274
	goto L80
L71:
	;
	if (v130|v239)&int32(3) != 0 {
		v272 = v130
		v273 = v239
		v274 = v241
		goto L70
	} else {
		goto L74
	}
L72:
	;
	v265 = v130
	v266 = v239
	v267 = v241
	goto L73
L73:
	;
	if v267 == int32(0) {
		goto L69
	} else {
		goto L79
	}
L74:
	;
	v249 = v130
	v250 = v239
	v251 = v241
	goto L75
L75:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v254 != v255 {
		v272 = v249
		v273 = v250
		v274 = v251
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v265 = v260
	v266 = v258
	v267 = v262
	goto L73
L77:
	;
	v257 = int32(4)
	v258 = v250 + v257
	v260 = v249 + v257
	v262 = v251 - v257
	if base.Ui32(int32(3)) < base.Ui32(v262) {
		v249 = v260
		v250 = v258
		v251 = v262
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v272 = v265
	v273 = v266
	v274 = v267
	goto L70
L80:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v282 == v283 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v303 = v282 - v283
	goto L68
L82:
	;
	v285 = int32(1)
	v290 = v279 - v285
	if v290 != 0 {
		v277 = v277 + v285
		v278 = v278 + v285
		v279 = v290
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	goto L69
L86:
	;
	if v303 != 0 {
		goto L39
	} else {
		goto L98
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L95
	}
L88:
	;
	if int32(0) <= v303 {
		goto L45
	} else {
		goto L94
	}
L89:
	;
	if v233 < v163 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	if v303 <= int32(0) {
		goto L45
	} else {
		goto L91
	}
L91:
	;
	goto L39
L92:
	;
	if v303 == int32(0) {
		goto L45
	} else {
		goto L93
	}
L93:
	;
	goto L39
L94:
	;
	goto L39
L95:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318+v168*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v322
	F_errmsg_internal(m, int32(_a_F_spg_text_inner_consistent_1), v19)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_spg_text_inner_consistent_2), int32(551), int32(_a_F_spg_text_inner_consistent_3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	goto L45
L99:
	;
	goto L44
L100:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v379+v380<<(uint(int32(2))%32)))) = v377
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v385 + int32(1)
	goto L39
L101:
	;
	goto L35
}
