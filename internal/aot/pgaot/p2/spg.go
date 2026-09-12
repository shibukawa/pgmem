package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_spgFormNodeTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 != 0 {
		v64 = int32(8)
		v66 = int32(32768)
		v67 = F_palloc0(m, v64)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			v71 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
			*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
			v75 = v64 | v66
			*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
			if l2 != 0 {
			} else {
				v78 = v67 + int32(8)
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
				if v79 == int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
				} else {
					v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
					if int32(0) < v83 {
						v112 = v83
					} else {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						if v86 == int32(1) {
							v89 = int32(6)
							v91 = int32(18)
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if v93 == v91 {
								v96 = v91
							} else {
								v96 = int32(2)
							}
							if v93&int32(254) == int32(2) {
								v101 = v89
							} else {
								v101 = v96
							}
							if v93 == int32(1) {
								v104 = v89
							} else {
								v104 = v101
							}
							v112 = v104
						} else {
							if v86&int32(1) != 0 {
								v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
							} else {
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
							}
						}
					}
					if v112 != 0 {
						v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
						mBase = m.M
					} else {
					}
				}
			}
			m.G0 = v10 + int32(16)
			return v67
		}
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
		if v14 != 0 {
			v64 = int32(16)
			v66 = v4
			v67 = F_palloc0(m, v64)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
				*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
				v75 = v64 | v66
				*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
				if l2 != 0 {
				} else {
					v78 = v67 + int32(8)
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
					if v79 == int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
					} else {
						v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
						if int32(0) < v83 {
							v112 = v83
						} else {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							if v86 == int32(1) {
								v89 = int32(6)
								v91 = int32(18)
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								if v93 == v91 {
									v96 = v91
								} else {
									v96 = int32(2)
								}
								if v93&int32(254) == int32(2) {
									v101 = v89
								} else {
									v101 = v96
								}
								if v93 == int32(1) {
									v104 = v89
								} else {
									v104 = v101
								}
								v112 = v104
							} else {
								if v86&int32(1) != 0 {
									v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
								}
							}
						}
						if v112 != 0 {
							v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
							mBase = m.M
						} else {
						}
					}
				}
				m.G0 = v10 + int32(16)
				return v67
			}
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
			if v16 <= int32(0) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v19 == int32(1) {
					v22 = int32(6)
					v24 = int32(18)
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					if v26 == v24 {
						v29 = v24
					} else {
						v29 = int32(2)
					}
					if v26&int32(254) == int32(2) {
						v34 = v22
					} else {
						v34 = v29
					}
					if v26 == int32(1) {
						v37 = v22
					} else {
						v37 = v34
					}
					v45 = v37
					v64 = (v45+int32(7))&int32(248) + int32(8)
					v66 = v4
					v67 = F_palloc0(m, v64)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
						v75 = v64 | v66
						*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
						if l2 != 0 {
						} else {
							v78 = v67 + int32(8)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v79 == int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
							} else {
								v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v83 {
									v112 = v83
								} else {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
									if v86 == int32(1) {
										v89 = int32(6)
										v91 = int32(18)
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
										if v93 == v91 {
											v96 = v91
										} else {
											v96 = int32(2)
										}
										if v93&int32(254) == int32(2) {
											v101 = v89
										} else {
											v101 = v96
										}
										if v93 == int32(1) {
											v104 = v89
										} else {
											v104 = v101
										}
										v112 = v104
									} else {
										if v86&int32(1) != 0 {
											v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
										} else {
											v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v112 != 0 {
									v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
									mBase = m.M
								} else {
								}
							}
						}
						m.G0 = v10 + int32(16)
						return v67
					}
				} else {
					if v19&int32(1) == int32(0) {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v55 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
						v61 = (v55+int32(7))&int32(2147483640) + int32(8)
						if base.Ui32(int32(8177)) <= base.Ui32(v55) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8191)
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
									F_errmsg(m, int32(36281), v10)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(485051), int32(979), int32(378215))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
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
							v64 = v61
							v66 = v4
							v67 = F_palloc0(m, v64)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
								*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
								v75 = v64 | v66
								*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
								if l2 != 0 {
								} else {
									v78 = v67 + int32(8)
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
									if v79 == int32(1) {
										*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
									} else {
										v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
										if int32(0) < v83 {
											v112 = v83
										} else {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
											if v86 == int32(1) {
												v89 = int32(6)
												v91 = int32(18)
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
												if v93 == v91 {
													v96 = v91
												} else {
													v96 = int32(2)
												}
												if v93&int32(254) == int32(2) {
													v101 = v89
												} else {
													v101 = v96
												}
												if v93 == int32(1) {
													v104 = v89
												} else {
													v104 = v101
												}
												v112 = v104
											} else {
												if v86&int32(1) != 0 {
													v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
												} else {
													v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
													v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
												}
											}
										}
										if v112 != 0 {
											v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
											mBase = m.M
										} else {
										}
									}
								}
								m.G0 = v10 + int32(16)
								return v67
							}
						}
					} else {
						v45 = int32(base.Ui32(v19) >> (uint(int32(1)) % 32))
						v64 = (v45+int32(7))&int32(248) + int32(8)
						v66 = v4
						v67 = F_palloc0(m, v64)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
							v75 = v64 | v66
							*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
							if l2 != 0 {
							} else {
								v78 = v67 + int32(8)
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
								if v79 == int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
								} else {
									v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
									if int32(0) < v83 {
										v112 = v83
									} else {
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
										if v86 == int32(1) {
											v89 = int32(6)
											v91 = int32(18)
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
											if v93 == v91 {
												v96 = v91
											} else {
												v96 = int32(2)
											}
											if v93&int32(254) == int32(2) {
												v101 = v89
											} else {
												v101 = v96
											}
											if v93 == int32(1) {
												v104 = v89
											} else {
												v104 = v101
											}
											v112 = v104
										} else {
											if v86&int32(1) != 0 {
												v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
											}
										}
									}
									if v112 != 0 {
										v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
										mBase = m.M
									} else {
									}
								}
							}
							m.G0 = v10 + int32(16)
							return v67
						}
					}
				}
			} else {
				v55 = v16
				v61 = (v55+int32(7))&int32(2147483640) + int32(8)
				if base.Ui32(int32(8177)) <= base.Ui32(v55) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8191)
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
							F_errmsg(m, int32(36281), v10)
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485051), int32(979), int32(378215))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
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
					v64 = v61
					v66 = v4
					v67 = F_palloc0(m, v64)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v71)
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
						v75 = v64 | v66
						*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v75)
						if l2 != 0 {
						} else {
							v78 = v67 + int32(8)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v79 == int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v78))) = l1
							} else {
								v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v83 {
									v112 = v83
								} else {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
									if v86 == int32(1) {
										v89 = int32(6)
										v91 = int32(18)
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
										if v93 == v91 {
											v96 = v91
										} else {
											v96 = int32(2)
										}
										if v93&int32(254) == int32(2) {
											v101 = v89
										} else {
											v101 = v96
										}
										if v93 == int32(1) {
											v104 = v89
										} else {
											v104 = v101
										}
										v112 = v104
									} else {
										if v86&int32(1) != 0 {
											v112 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
										} else {
											v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v112 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v112 != 0 {
									v113 = F__emscripten_memcpy_bulkmem(m, v78, l1, v112)
									mBase = m.M
								} else {
								}
							}
						}
						m.G0 = v10 + int32(16)
						return v67
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
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v230 float64
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v291 float64
	_ = v291
	var v293 float64
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v321 int32
	_ = v321
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
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v385 float64
	_ = v385
	var v394 int32
	_ = v394
	var v396 float64
	_ = v396
	var v402 int32
	_ = v402
	var v404 float64
	_ = v404
	var v410 int32
	_ = v410
	var v412 float64
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 float64
	_ = v445
	var v446 float64
	_ = v446
	var v452 float64
	_ = v452
	var v453 float64
	_ = v453
	var v459 float64
	_ = v459
	var v460 float64
	_ = v460
	var v466 float64
	_ = v466
	var v467 float64
	_ = v467
	var v474 int32
	_ = v474
	var v475 float64
	_ = v475
	var v476 float64
	_ = v476
	var v482 float64
	_ = v482
	var v483 float64
	_ = v483
	var v489 float64
	_ = v489
	var v490 float64
	_ = v490
	var v496 float64
	_ = v496
	var v497 float64
	_ = v497
	var v504 int32
	_ = v504
	var v505 float64
	_ = v505
	var v507 float64
	_ = v507
	var v508 float64
	_ = v508
	var v512 float64
	_ = v512
	var v513 float64
	_ = v513
	var v519 float64
	_ = v519
	var v523 float64
	_ = v523
	var v529 float64
	_ = v529
	var v531 float64
	_ = v531
	var v532 float64
	_ = v532
	var v536 float64
	_ = v536
	var v537 float64
	_ = v537
	var v543 float64
	_ = v543
	var v547 float64
	_ = v547
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v556 float64
	_ = v556
	var v562 float64
	_ = v562
	var v569 int32
	_ = v569
	var v570 float64
	_ = v570
	var v572 float64
	_ = v572
	var v573 float64
	_ = v573
	var v577 float64
	_ = v577
	var v582 int32
	_ = v582
	var v583 float64
	_ = v583
	var v585 float64
	_ = v585
	var v586 float64
	_ = v586
	var v590 float64
	_ = v590
	var v595 int32
	_ = v595
	var v596 float64
	_ = v596
	var v597 float64
	_ = v597
	var v603 float64
	_ = v603
	var v610 int32
	_ = v610
	var v611 float64
	_ = v611
	var v613 float64
	_ = v613
	var v614 float64
	_ = v614
	var v618 float64
	_ = v618
	var v623 int32
	_ = v623
	var v624 float64
	_ = v624
	var v625 float64
	_ = v625
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
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 float64
	_ = v669
	var v671 float64
	_ = v671
	var v672 float64
	_ = v672
	var v676 float64
	_ = v676
	var v684 int32
	_ = v684
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 float64
	_ = v755
	var v756 float64
	_ = v756
	var v760 float64
	_ = v760
	var v766 float64
	_ = v766
	var v767 float64
	_ = v767
	var v768 float64
	_ = v768
	var v772 float64
	_ = v772
	var v778 float64
	_ = v778
	var v779 float64
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
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
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v222 = F_palloc(m, int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L45
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
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208+v204)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v213 = v211 << (uint(int32(3)) % 32)
	if v213 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v217 = v185 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v217 < v218 {
		v185 = v217
		goto L37
	} else {
		goto L44
	}
L41:
	;
	v214 = F__emscripten_memcpy_bulkmem(m, v210, v94, v213)
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L38
L45:
	;
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v220)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v222))) = v224
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	*(*float64)(unsafe.Add(mBase, uint32(v222)+8)) = v226
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v220)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v222)+16)) = v228
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v220)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v222)+24)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v235 = F_palloc(m, v232<<(uint(int32(2))%32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v237 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = v4
	goto L50
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v324 = F_palloc(m, v321<<(uint(int32(2))%32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L62
	}
L50:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v259 = v256 + v249*int32(48)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	switch v260 - int32(603) {
	case 0:
		goto L53
	case 1:
		goto L55
	default:
		goto L54
	}
L51:
	;
	goto L49
L52:
	;
	v285 = F_palloc(m, int32(32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L60
	}
L53:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v259)+44))
	v283 = v282
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+44))
	v264 = F_pg_detoast_datum(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v283 = v264 + int32(8)
	goto L52
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v272
	F_errmsg_internal(m, int32(475651), v19)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(483830), int32(544), int32(26690))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
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
	v287 = *(*float64)(unsafe.Add(mBase, uint32(v283)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v285))) = v287
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v283)))
	*(*float64)(unsafe.Add(mBase, uint32(v285)+8)) = v289
	v291 = *(*float64)(unsafe.Add(mBase, uint32(v283)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v285)+16)) = v291
	v293 = *(*float64)(unsafe.Add(mBase, uint32(v283)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v285)+24)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v235+v249<<(uint(int32(2))%32)))) = v285
	v300 = v249 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v300 < v301 {
		v249 = v300
		goto L50
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v324
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v330 = F_palloc(m, v327<<(uint(int32(2))%32))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v333 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v339 = F_palloc(m, v336<<(uint(int32(2))%32))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v342 = int32(4470400)
	v343 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v345
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if int32(0) < v347 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v339
	goto L66
L68:
	;
	v357 = v4
	v363 = v4
	goto L71
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v343
	goto L6
L71:
	;
	v367 = F_palloc(m, int32(64))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L73
	}
L72:
	;
	goto L70
L73:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v47)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+56)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v47)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+48)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v47)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+40)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+32)) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v47)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+24)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+16)) = v379
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+8)) = v381
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = v383
	v385 = *(*float64)(unsafe.Add(mBase, uint32(v222)))
	if v357&int32(8) != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v357&int32(4) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v367))) = v385
	goto L74
L76:
	;
	goto L77
L77:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v367)+8)) = v385
	goto L74
L78:
	;
	v394 = int32(16)
	goto L80
L79:
	;
	v394 = int32(24)
	goto L80
L80:
	;
	v396 = *(*float64)(unsafe.Add(mBase, uint32(v222)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v367+v394))) = v396
	if v357&int32(2) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v402 = int32(32)
	goto L83
L82:
	;
	v402 = int32(40)
	goto L83
L83:
	;
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v222)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v367+v402))) = v404
	if v357&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v410 = int32(48)
	goto L86
L85:
	;
	v410 = int32(56)
	goto L86
L86:
	;
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v222)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v367+v410))) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v414 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v828 = v363 + int32(1)
	v830 = v828 & int32(255)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v830 < v831 {
		v357 = v830
		v363 = v828
		goto L71
	} else {
		goto L162
	}
L88:
	;
	F_pfree(m, v367)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L4
	} else {
		goto L161
	}
L89:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v422 = int32(0)
	goto L92
L90:
	;
	goto L91
L91:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v704 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v702+v703<<(uint(v704)%32)))) = v367
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v708+v709<<(uint(v704)%32)))) = v357
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v714 <= int32(0) {
		goto L143
	} else {
		goto L144
	}
L92:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417+v422*int32(48))+6)))
	switch v438 - int32(1) {
	case 0:
		goto L103
	case 1:
		goto L102
	case 2:
		goto L106
	case 3:
		goto L100
	case 4:
		goto L101
	case 5, 7:
		goto L104
	case 6:
		goto L105
	case 8:
		goto L95
	case 9:
		goto L97
	case 10:
		goto L99
	case 11:
		goto L98
	default:
		goto L96
	}
L93:
	;
	goto L91
L94:
	;
	v684 = v422 + int32(1)
	if v684 != v414 {
		v422 = v684
		goto L92
	} else {
		goto L142
	}
L95:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v669 = *(*float64)(unsafe.Add(mBase, uint32(v668)+24))
	v671 = base.F64_add(v669, float64(1e-06))
	v672 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	if base.F64_ge(v671, v672) == int32(0) {
		goto L88
	} else {
		goto L140
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L137
	}
L97:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v639 = *(*float64)(unsafe.Add(mBase, uint32(v638)+16))
	v640 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	if base.F64_gt(v639, base.F64_add(v640, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L135
	}
L98:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v624 = *(*float64)(unsafe.Add(mBase, uint32(v623)+16))
	v625 = *(*float64)(unsafe.Add(mBase, uint32(v367)+40))
	if base.F64_le(v624, base.F64_add(v625, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L133
	}
L99:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v611 = *(*float64)(unsafe.Add(mBase, uint32(v610)+24))
	v613 = base.F64_add(v611, float64(1e-06))
	v614 = *(*float64)(unsafe.Add(mBase, uint32(v367)+40))
	if base.F64_lt(v613, v614) == int32(0) {
		goto L88
	} else {
		goto L131
	}
L100:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v596 = *(*float64)(unsafe.Add(mBase, uint32(v595)))
	v597 = *(*float64)(unsafe.Add(mBase, uint32(v367)+8))
	if base.F64_le(v596, base.F64_add(v597, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L129
	}
L101:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v583 = *(*float64)(unsafe.Add(mBase, uint32(v582)+8))
	v585 = base.F64_add(v583, float64(1e-06))
	v586 = *(*float64)(unsafe.Add(mBase, uint32(v367)+8))
	if base.F64_lt(v585, v586) == int32(0) {
		goto L88
	} else {
		goto L127
	}
L102:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v570 = *(*float64)(unsafe.Add(mBase, uint32(v569)+8))
	v572 = base.F64_add(v570, float64(1e-06))
	v573 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	if base.F64_ge(v572, v573) == int32(0) {
		goto L88
	} else {
		goto L125
	}
L103:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v555 = *(*float64)(unsafe.Add(mBase, uint32(v554)))
	v556 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	if base.F64_gt(v555, base.F64_add(v556, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L123
	}
L104:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v505 = *(*float64)(unsafe.Add(mBase, uint32(v504)+8))
	v507 = base.F64_add(v505, float64(1e-06))
	v508 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	if base.F64_ge(v507, v508) == int32(0) {
		goto L88
	} else {
		goto L115
	}
L105:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v475 = *(*float64)(unsafe.Add(mBase, uint32(v474)+8))
	v476 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_le(v475, base.F64_add(v476, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L111
	}
L106:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v235+v422<<(uint(int32(2))%32))))
	v445 = *(*float64)(unsafe.Add(mBase, uint32(v444)))
	v446 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_le(v445, base.F64_add(v446, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L107
	}
L107:
	;
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	v453 = *(*float64)(unsafe.Add(mBase, uint32(v444)+8))
	if base.F64_le(v452, base.F64_add(v453, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L108
	}
L108:
	;
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v444)+16))
	v460 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_le(v459, base.F64_add(v460, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L109
	}
L109:
	;
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	v467 = *(*float64)(unsafe.Add(mBase, uint32(v444)+24))
	if base.F64_le(v466, base.F64_add(v467, float64(1e-06))) != 0 {
		goto L94
	} else {
		goto L110
	}
L110:
	;
	goto L88
L111:
	;
	v482 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	v483 = *(*float64)(unsafe.Add(mBase, uint32(v474)))
	if base.F64_le(v482, base.F64_add(v483, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L112
	}
L112:
	;
	v489 = *(*float64)(unsafe.Add(mBase, uint32(v474)+24))
	v490 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_le(v489, base.F64_add(v490, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L113
	}
L113:
	;
	v496 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v474)+16))
	if base.F64_le(v496, base.F64_add(v497, float64(1e-06))) != 0 {
		goto L94
	} else {
		goto L114
	}
L114:
	;
	goto L88
L115:
	;
	v512 = *(*float64)(unsafe.Add(mBase, uint32(v504)))
	v513 = *(*float64)(unsafe.Add(mBase, uint32(v367)+8))
	if base.F64_le(v512, base.F64_add(v513, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L116
	}
L116:
	;
	v519 = *(*float64)(unsafe.Add(mBase, uint32(v367)+16))
	if base.F64_le(v519, v507) == int32(0) {
		goto L88
	} else {
		goto L117
	}
L117:
	;
	v523 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_ge(base.F64_add(v523, float64(1e-06)), v512) == int32(0) {
		goto L88
	} else {
		goto L118
	}
L118:
	;
	v529 = *(*float64)(unsafe.Add(mBase, uint32(v504)+24))
	v531 = base.F64_add(v529, float64(1e-06))
	v532 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	if base.F64_ge(v531, v532) == int32(0) {
		goto L88
	} else {
		goto L119
	}
L119:
	;
	v536 = *(*float64)(unsafe.Add(mBase, uint32(v504)+16))
	v537 = *(*float64)(unsafe.Add(mBase, uint32(v367)+40))
	if base.F64_le(v536, base.F64_add(v537, float64(1e-06))) == int32(0) {
		goto L88
	} else {
		goto L120
	}
L120:
	;
	v543 = *(*float64)(unsafe.Add(mBase, uint32(v367)+48))
	if base.F64_le(v543, v531) == int32(0) {
		goto L88
	} else {
		goto L121
	}
L121:
	;
	v547 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_ge(base.F64_add(v547, float64(1e-06)), v536) != 0 {
		goto L94
	} else {
		goto L122
	}
L122:
	;
	goto L88
L123:
	;
	v562 = *(*float64)(unsafe.Add(mBase, uint32(v367)+16))
	if base.F64_lt(base.F64_add(v562, float64(1e-06)), v555) != 0 {
		goto L94
	} else {
		goto L124
	}
L124:
	;
	goto L88
L125:
	;
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v367)+16))
	if base.F64_le(v577, v572) != 0 {
		goto L94
	} else {
		goto L126
	}
L126:
	;
	goto L88
L127:
	;
	v590 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_gt(v590, v585) != 0 {
		goto L94
	} else {
		goto L128
	}
L128:
	;
	goto L88
L129:
	;
	v603 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_ge(base.F64_add(v603, float64(1e-06)), v596) != 0 {
		goto L94
	} else {
		goto L130
	}
L130:
	;
	goto L88
L131:
	;
	v618 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_gt(v618, v613) != 0 {
		goto L94
	} else {
		goto L132
	}
L132:
	;
	goto L88
L133:
	;
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_ge(base.F64_add(v631, float64(1e-06)), v624) != 0 {
		goto L94
	} else {
		goto L134
	}
L134:
	;
	goto L88
L135:
	;
	v646 = *(*float64)(unsafe.Add(mBase, uint32(v367)+48))
	if base.F64_lt(base.F64_add(v646, float64(1e-06)), v639) != 0 {
		goto L94
	} else {
		goto L136
	}
L136:
	;
	goto L88
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v438
	F_errmsg_internal(m, int32(472156), v19+int32(16))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(483830), int32(691), int32(90762))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v676 = *(*float64)(unsafe.Add(mBase, uint32(v367)+48))
	if base.F64_le(v676, v671) == int32(0) {
		goto L88
	} else {
		goto L141
	}
L141:
	;
	goto L94
L142:
	;
	goto L93
L143:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v802 + int32(1)
	goto L87
L144:
	;
	v719 = F_palloc(m, v714<<(uint(int32(3))%32))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v721+v722<<(uint(int32(2))%32)))) = v719
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v727 <= int32(0) {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v734 = int32(0)
	goto L147
L147:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v750+v734*int32(48))+44))
	v755 = *(*float64)(unsafe.Add(mBase, uint32(v754)))
	v756 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	if base.F64_lt(v755, v756) != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	goto L143
L149:
	;
	v767 = *(*float64)(unsafe.Add(mBase, uint32(v754)+8))
	v768 = *(*float64)(unsafe.Add(mBase, uint32(v367)+32))
	if base.F64_lt(v767, v768) != 0 {
		goto L155
	} else {
		goto L156
	}
L150:
	;
	v766 = base.F64_sub(v756, v755)
	goto L149
L151:
	;
	goto L152
L152:
	;
	v760 = *(*float64)(unsafe.Add(mBase, uint32(v367)+24))
	if base.F64_gt(v755, v760) == int32(0) {
		v766 = float64(0)
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v766 = base.F64_sub(v755, v760)
	goto L149
L154:
	;
	v779 = F_pg_hypot(m, v766, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L159
	}
L155:
	;
	v778 = base.F64_sub(v768, v767)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v772 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	if base.F64_gt(v767, v772) == int32(0) {
		v778 = float64(0)
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v778 = base.F64_sub(v767, v772)
	goto L154
L159:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v719+v734<<(uint(int32(3))%32)))) = v779
	v783 = v734 + int32(1)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v783 < v784 {
		v734 = v783
		goto L147
	} else {
		goto L160
	}
L160:
	;
	goto L148
L161:
	;
	goto L87
L162:
	;
	goto L72
}
func F_spg_box_quad_leaf_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v2)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
	if v17 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v13
	goto L3
L2:
	;
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if int32(0) < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return v179
L5:
	;
	v28 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v165 = int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v166 <= int32(0) {
		v179 = v165
		goto L4
	} else {
		goto L67
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v34 = v31 + v28*int32(48)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	switch v36 - int32(603) {
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
	switch v35 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L26
	case 3:
		goto L32
	case 4:
		goto L33
	case 5:
		goto L36
	case 6:
		goto L38
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v73 = v72
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L22
	}
L13:
	;
	if int32(1)<<(uint(v35)%32)&int32(7734) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = base.B2i32(base.Ui32(v35) <= base.Ui32(int32(12)))
	goto L16
L15:
	;
	v46 = int32(0)
	goto L16
L16:
	;
	if v46 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v49)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v73 = v52 + int32(8)
	goto L10
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v62
	F_errmsg_internal(m, int32(475651), v10)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(483830), int32(544), int32(26690))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	v155 = v28 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v155 < v156 {
		v28 = v155
		goto L8
	} else {
		goto L66
	}
L26:
	;
	v146 = int32(0)
	v149 = F_DirectFunctionCall2Coll(m, int32(96), v146, v13, v73)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L20
	} else {
		goto L64
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L61
	}
L28:
	;
	v126 = int32(0)
	v129 = F_DirectFunctionCall2Coll(m, int32(102), v126, v13, v73)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L59
	}
L29:
	;
	v121 = int32(0)
	v124 = F_DirectFunctionCall2Coll(m, int32(103), v121, v13, v73)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L57
	}
L30:
	;
	v116 = int32(0)
	v119 = F_DirectFunctionCall2Coll(m, int32(101), v116, v13, v73)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L20
	} else {
		goto L55
	}
L31:
	;
	v111 = int32(0)
	v114 = F_DirectFunctionCall2Coll(m, int32(100), v111, v13, v73)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L53
	}
L32:
	;
	v106 = int32(0)
	v109 = F_DirectFunctionCall2Coll(m, int32(104), v106, v13, v73)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L51
	}
L33:
	;
	v101 = int32(0)
	v104 = F_DirectFunctionCall2Coll(m, int32(95), v101, v13, v73)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L20
	} else {
		goto L49
	}
L34:
	;
	v96 = int32(0)
	v99 = F_DirectFunctionCall2Coll(m, int32(98), v96, v13, v73)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L47
	}
L35:
	;
	v91 = int32(0)
	v94 = F_DirectFunctionCall2Coll(m, int32(97), v91, v13, v73)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L20
	} else {
		goto L45
	}
L36:
	;
	v86 = int32(0)
	v89 = F_DirectFunctionCall2Coll(m, int32(117), v86, v13, v73)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L43
	}
L37:
	;
	v81 = int32(0)
	v84 = F_DirectFunctionCall2Coll(m, int32(118), v81, v13, v73)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L41
	}
L38:
	;
	v76 = int32(0)
	v79 = F_DirectFunctionCall2Coll(m, int32(99), v76, v13, v73)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v79 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v179 = v76
	goto L4
L41:
	;
	if v84 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v179 = v81
	goto L4
L43:
	;
	if v89 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v179 = v86
	goto L4
L45:
	;
	if v94 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v179 = v91
	goto L4
L47:
	;
	if v99 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v179 = v96
	goto L4
L49:
	;
	if v104 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v179 = v101
	goto L4
L51:
	;
	if v109 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	v179 = v106
	goto L4
L53:
	;
	if v114 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v179 = v111
	goto L4
L55:
	;
	if v119 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v179 = v116
	goto L4
L57:
	;
	if v124 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	v179 = v121
	goto L4
L59:
	;
	if v129 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	v179 = v126
	goto L4
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v35
	F_errmsg_internal(m, int32(472156), v10+int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(483830), int32(831), int32(91239))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
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
	if v149 == int32(0) {
		v179 = v146
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
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+20))
	v172 = F_spg_key_orderbys_distances(m, v13, int32(0), v169, v166)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(base.B2i32(v170 == int32(3292)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v172
	v179 = v165
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
			F_errmsg_internal(m, int32(168279), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(490876), int32(62), int32(354705))
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
	var v52 int32
	_ = v52
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
	var v183 int32
	_ = v183
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
	v52 = int32(6)
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
	v183 = int32(6)
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	F_errmsg_internal(m, int32(168279), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(490876), int32(173), int32(90738))
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
	v183 = v174
	goto L2
L14:
	;
	v177 = v48 + int32(1)
	if v177 != v24 {
		v48 = v177
		v52 = v174
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
	v153 = v52 & int32(4)
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
		v174 = v52
		goto L14
	} else {
		goto L37
	}
L20:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v108&int32(1) != 0 {
		v174 = v52
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
		v174 = v52
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v67&int32(1) == int32(0) {
		v174 = v52
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_lt(base.F64_add(v72, float64(1e-06)), v23) == int32(0) {
		v174 = v52
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v153 = v52 & int32(2)
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
	v174 = v52
	goto L14
L28:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(v23, base.F64_add(v92, float64(1e-06))) != 0 {
		v153 = v52 & int32(2)
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
		v153 = v52 & int32(2)
		goto L16
	} else {
		goto L33
	}
L31:
	;
	if base.F64_lt(v44, v92) == int32(0) {
		v174 = v52
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
		v174 = v52
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
		v174 = v52
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v153 = v52 & int32(2)
	goto L16
L37:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	if base.F64_gt(v122, v44) == int32(0) {
		v174 = v52
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
		v153 = v52 & int32(2)
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
		v153 = v52 & int32(2)
		goto L16
	} else {
		goto L44
	}
L42:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v63)+16))
	if base.F64_gt(v135, v44) == int32(0) {
		v174 = v52
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
		v174 = v52
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
	F_errmsg_internal(m, int32(473734), v15)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(490876), int32(247), int32(90738))
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
	if v183&int32(2) != 0 {
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
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v238
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
	if v183&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v275 = int32(4470400)
	v276 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v278
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v276
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
	v323 = int32(4470400)
	v324 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v326
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v324
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
	var v91 float64
	_ = v91
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
	*(*float64)(unsafe.Add(mBase, uint32(v31))) = v91
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
		v91 = v41
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v91 = v40
	goto L8
L13:
	;
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) {
		v91 = v41
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v35)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) {
		v91 = v41
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
		v91 = v41
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if base.F64_gt(v48, v42) != 0 {
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
	v91 = v86
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
		goto L9
	case 1:
		goto L19
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
	v47 = F_range_overleft_internal(m, v24, v15, v45)
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
	F_errmsg_internal(m, int32(472124), v10)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(483810), int32(985), int32(91293))
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
	v108 = F_range_before_internal(m, v24, v15, v106)
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
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
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum(m, v18)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v29 = F_palloc(m, v26<<(uint(int32(3))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v34 = F_palloc(m, v31<<(uint(int32(3))%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v13 + int32(48)
	return int32(0)
L7:
	;
	F_qsort_arg(m, v29, v68, int32(8), int32(1488), v24)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	v40 = int32(0)
	v45 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v83
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v83)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(8589934592)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v93 = F_palloc(m, v90<<(uint(int32(2))%32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v45<<(uint(int32(2))%32))))
	v55 = F_pg_detoast_datum(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	if v68 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v58 = v40 << (uint(int32(3)) % 32)
	F_range_deserialize(m, v24, v55, v29+v58, v34+v58, v13+int32(10))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)))
	v66 = int32(1)
	v68 = v40 + (v65 ^ v66)
	v70 = v45 + v66
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v70 < v71 {
		v40 = v68
		v45 = v70
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v99 = F_palloc(m, v96<<(uint(int32(2))%32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v102 <= int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v105 = v83
	goto L20
L20:
	;
	v116 = v105 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116+v117)))
	v120 = F_pg_detoast_datum(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L6
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v122+v116))) = v120
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v116))) = int32(0)
	v130 = v105 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v130 < v131 {
		v105 = v130
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_qsort_arg(m, v34, v68, int32(8), int32(1488), v24)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v141 = int32(0)
	v143 = base.I32_div_s(v68, int32(2))
	v145 = v143 << (uint(int32(3)) % 32)
	v150 = F_range_serialize(m, v24, v29+v145, v145+v34, v141, v141)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v150
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	if v155 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v160 = int32(4)
	goto L29
L28:
	;
	v160 = int32(5)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v165 = F_palloc(m, v162<<(uint(int32(2))%32))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v171 = F_palloc(m, v168<<(uint(int32(2))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v174 <= int32(0) {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v182 = v141
	goto L33
L33:
	;
	v188 = v182 << (uint(int32(2)) % 32)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188+v189)))
	v192 = F_pg_detoast_datum(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L6
L35:
	;
	F_range_deserialize(m, v24, v150, v13+int32(40), v13+int32(32), v13+int32(31))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_range_deserialize(m, v24, v192, v13+int32(20), v13+int32(12), v13+int32(11))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if v211 != 0 {
		v238 = int32(5)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v239+v188))) = v192
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v244 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v242+v188))) = v238 - v244
	v248 = v182 + v244
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v248 < v249 {
		v182 = v248
		goto L33
	} else {
		goto L49
	}
L39:
	;
	v216 = F_range_cmp_bounds(m, v24, v13+int32(20), v13+int32(40))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v222 = F_range_cmp_bounds(m, v24, v13+int32(12), v13+int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v222 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v228 = int32(1)
	goto L44
L43:
	;
	v228 = int32(2)
	goto L44
L44:
	;
	if int32(0) <= v216 {
		v238 = v228
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if int32(0) <= v222 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v235 = int32(4)
	goto L48
L47:
	;
	v235 = int32(3)
	goto L48
L48:
	;
	v238 = v235
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
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
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v297 int32
	_ = v297
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
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
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
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
	if v34 != v29 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v76 = v74 + int32(4)
	v77 = F_palloc(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	v72 = int32(0)
	v73 = v2
	v74 = v30
	goto L3
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v38 = F_pg_detoast_datum_packed(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v72 = v70
	v73 = v38
	v74 = v30 + v70
	goto L3
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v40 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = int32(4)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v45&int32(254) == int32(2) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v58 = int32(1)
	if v40&v58 != 0 {
		v70 = int32(base.Ui32(v40)>>(uint(v58)%32)) - v58
		goto L7
	} else {
		goto L18
	}
L12:
	;
	v54 = v43
	goto L14
L13:
	;
	v54 = base.B2i32(v45 == int32(18)) << (uint(v43) % 32)
	goto L14
L14:
	;
	if v45 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = v43
	goto L17
L16:
	;
	v57 = v54
	goto L17
L17:
	;
	v70 = v57
	goto L7
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v76 << (uint(int32(2)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v82 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v83 = int32(4)
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	v87 = F__emscripten_memcpy_bulkmem(m, v77+v83, v31+v83, v82)
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v91 = int32(4)
	v93 = int32(1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v95&v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v105 = F_palloc(m, v102<<(uint(int32(2))%32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v98 = v93
	goto L32
L31:
	;
	v98 = v91
	goto L32
L32:
	;
	if v72 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L29
L34:
	;
	v100 = F__emscripten_memcpy_bulkmem(m, v77+v89+v91, v73+v98, v72)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v111 = F_palloc(m, v108<<(uint(int32(2))%32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v117 = F_palloc(m, v114<<(uint(int32(2))%32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v117
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v119 < v122 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v126 = v77 + int32(4)
	v127 = int32(1)
	v144 = v2
	goto L43
L41:
	;
	goto L42
L42:
	;
	m.G0 = v19 + int32(16)
	return int32(0)
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v144<<(uint(int32(2))%32))))
	if int32(0) < base.I32_extend16_s(v154) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v77+int32(3)))) = uint8(v154)
	v159 = v74
	goto L47
L46:
	;
	v159 = v74 - v127
	goto L47
L47:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v160 < v161 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v400 = v144 + int32(1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v400 < v401 {
		v144 = v400
		goto L43
	} else {
		goto L111
	}
L49:
	;
	v164 = v160
	v166 = v161
	goto L52
L50:
	;
	goto L51
L51:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v352 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v350+v351<<(uint(v352)%32)))) = v144
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v357<<(uint(v352)%32)))) = v159 - v361
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v159<<(uint(v352)%32) + int32(16)
	v371 = F_datumCopy(m, v77, int32(0), int32(-1))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L110
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v183 = v180 + v164*int32(48)
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+6)))
	if base.Ui32(v184) < base.Ui32(int32(11)) {
		v193 = v184
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L51
L54:
	;
	v332 = v164 + int32(1)
	if v332 < v327 {
		v164 = v332
		v166 = v327
		goto L52
	} else {
		goto L109
	}
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v195 = F_pg_detoast_datum_packed(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L60
	}
L56:
	;
	if v184 == int32(28) {
		v193 = v184
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if v32&v127 == int32(0) {
		v327 = v166
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v193 = v184 - int32(10)
	goto L55
L59:
	;
	v228 = int32(1)
	if v197&v228 != 0 {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v197 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v200 = int32(4)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	if v202&int32(254) == int32(2) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v215 = int32(1)
	if v197&v215 != 0 {
		v227 = int32(base.Ui32(v197)>>(uint(v215)%32)) - v215
		goto L59
	} else {
		goto L70
	}
L64:
	;
	v211 = v200
	goto L66
L65:
	;
	v211 = base.B2i32(v202 == int32(18)) << (uint(v200) % 32)
	goto L66
L66:
	;
	if v202 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v214 = v200
	goto L69
L68:
	;
	v214 = v211
	goto L69
L69:
	;
	v227 = v214
	goto L59
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v227 = int32(base.Ui32(v221)>>(uint(int32(2))%32)) - int32(4)
	goto L59
L71:
	;
	v232 = v228
	goto L73
L72:
	;
	v232 = int32(4)
	goto L73
L73:
	;
	v233 = v195 + v232
	v234 = base.B2i32(v227 < v159)
	if v227 < v159 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v235 = v227
	goto L76
L75:
	;
	v235 = v159
	goto L76
L76:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v235) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	switch v193&int32(65535) - int32(1) {
	case 0, 1:
		goto L100
	case 2:
		goto L96
	case 3, 4:
		goto L99
	default:
		goto L97
	case 27:
		goto L98
	}
L78:
	;
	v297 = int32(0)
	goto L77
L79:
	;
	v271 = v266
	v272 = v267
	v273 = v268
	goto L89
L80:
	;
	if (v126|v233)&int32(3) != 0 {
		v266 = v126
		v267 = v233
		v268 = v235
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v259 = v126
	v260 = v233
	v261 = v235
	goto L82
L82:
	;
	if v261 == int32(0) {
		goto L78
	} else {
		goto L88
	}
L83:
	;
	v243 = v126
	v244 = v233
	v245 = v235
	goto L84
L84:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	if v248 != v249 {
		v266 = v243
		v267 = v244
		v268 = v245
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v259 = v254
	v260 = v252
	v261 = v256
	goto L82
L86:
	;
	v251 = int32(4)
	v252 = v244 + v251
	v254 = v243 + v251
	v256 = v245 - v251
	if base.Ui32(int32(3)) < base.Ui32(v256) {
		v243 = v254
		v244 = v252
		v245 = v256
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v266 = v259
	v267 = v260
	v268 = v261
	goto L79
L89:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v276 == v277 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v297 = v276 - v277
	goto L77
L91:
	;
	v279 = int32(1)
	v284 = v273 - v279
	if v284 != 0 {
		v271 = v271 + v279
		v272 = v272 + v279
		v273 = v284
		goto L89
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	goto L78
L95:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v327 = v326
	goto L54
L96:
	;
	if v227 < v159 {
		goto L48
	} else {
		goto L107
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L104
	}
L98:
	;
	if v297 == int32(0) {
		goto L95
	} else {
		goto L103
	}
L99:
	;
	if int32(0) <= v297 {
		goto L95
	} else {
		goto L102
	}
L100:
	;
	if v297 <= int32(0) {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	goto L48
L102:
	;
	goto L48
L103:
	;
	goto L48
L104:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312+v164*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v316
	F_errmsg_internal(m, int32(473734), v19)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(490841), int32(551), int32(90686))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	if v297 != 0 {
		goto L48
	} else {
		goto L108
	}
L108:
	;
	goto L95
L109:
	;
	goto L53
L110:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v373+v374<<(uint(int32(2))%32)))) = v371
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v379 + int32(1)
	goto L48
L111:
	;
	goto L44
}
