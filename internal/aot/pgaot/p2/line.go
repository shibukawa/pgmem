package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v44 float64
	_ = v44
	var v54 float64
	_ = v54
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 float64
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v121 float64
	_ = v121
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = base.F64_abs(v13)
	if base.F64_le(v14, float64(1e-06)) == int32(0) {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		v20 = base.F64_abs(v19)
		if base.F64_le(v20, float64(1e-06)) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
			v54 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v54
			v100 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return float64(0)
			} else {
				if v100 == int32(0) {
					v106 = math.Float64frombits(uint64(0x7ff8000000000000))
					if l0 == int32(0) {
						v121 = v106
					} else {
						v109 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v109
						v111 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
						v121 = v106
					}
					m.G0 = v11 + int32(48)
					return v121
				} else {
					if l0 != 0 {
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v115
					} else {
					}
					v119 = F_point_dt(m, v11+int32(32), l2)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return float64(0)
					} else {
						v121 = v119
						m.G0 = v11 + int32(48)
						return v121
					}
				}
			}
		} else {
			v23 = base.F64_div(v19, v13)
			v24 = base.F64_abs(v23)
			v25 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(v24, v25)&base.F64_ne(v20, v25) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v23, float64(0))&base.F64_ne(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if base.F64_ne(v24, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if base.F64_ne(v23, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v23
							v59 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							v60 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
							v61 = base.F64_mul(v23, v60)
							v62 = base.F64_abs(v61)
							v63 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(v62, v63)&base.F64_ne(base.F64_abs(v60), v63) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v69 = float64(0)
								if base.F64_eq(v61, v69)&base.F64_ne(v60, v69) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v75 = math.Float64frombits(uint64(0x7ff0000000000000))
									v77 = base.F64_sub(v59, v61)
									if base.B2i32(base.F64_eq(base.F64_abs(v59), v75)|base.F64_ne(base.F64_abs(v77), v75) == int32(0))&base.F64_ne(v62, v75) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v77
										if base.F64_ne(v77, float64(0)) != 0 {
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
										}
										v100 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return float64(0)
										} else {
											if v100 == int32(0) {
												v106 = math.Float64frombits(uint64(0x7ff8000000000000))
												if l0 == int32(0) {
													v121 = v106
												} else {
													v109 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v109
													v111 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
													v121 = v106
												}
												m.G0 = v11 + int32(48)
												return v121
											} else {
												if l0 != 0 {
													v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113
													v115 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v115
												} else {
												}
												v119 = F_point_dt(m, v11+int32(32), l2)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return float64(0)
												} else {
													v121 = v119
													m.G0 = v11 + int32(48)
													return v121
												}
											}
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
							v54 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v54
							v100 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return float64(0)
							} else {
								if v100 == int32(0) {
									v106 = math.Float64frombits(uint64(0x7ff8000000000000))
									if l0 == int32(0) {
										v121 = v106
									} else {
										v109 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v109
										v111 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
										v121 = v106
									}
									m.G0 = v11 + int32(48)
									return v121
								} else {
									if l0 != 0 {
										v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113
										v115 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v115
									} else {
									}
									v119 = F_point_dt(m, v11+int32(32), l2)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return float64(0)
									} else {
										v121 = v119
										m.G0 = v11 + int32(48)
										return v121
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(-4616189618054758400)
						v44 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v44
						v100 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return float64(0)
						} else {
							if v100 == int32(0) {
								v106 = math.Float64frombits(uint64(0x7ff8000000000000))
								if l0 == int32(0) {
									v121 = v106
								} else {
									v109 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v109
									v111 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
									v121 = v106
								}
								m.G0 = v11 + int32(48)
								return v121
							} else {
								if l0 != 0 {
									v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113
									v115 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v115
								} else {
								}
								v119 = F_point_dt(m, v11+int32(32), l2)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return float64(0)
								} else {
									v121 = v119
									m.G0 = v11 + int32(48)
									return v121
								}
							}
						}
					}
				}
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(-4616189618054758400)
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v44
		v100 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return float64(0)
		} else {
			if v100 == int32(0) {
				v106 = math.Float64frombits(uint64(0x7ff8000000000000))
				if l0 == int32(0) {
					v121 = v106
				} else {
					v109 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v109
					v111 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
					v121 = v106
				}
				m.G0 = v11 + int32(48)
				return v121
			} else {
				if l0 != 0 {
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113
					v115 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v115
				} else {
				}
				v119 = F_point_dt(m, v11+int32(32), l2)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return float64(0)
				} else {
					v121 = v119
					m.G0 = v11 + int32(48)
					return v121
				}
			}
		}
	}
}
func F_line_construct(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v22 float64
	_ = v22
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4616189618054758400)
		v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v14
		return
	} else {
		if base.F64_eq(l2, float64(0)) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v22
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*float64)(unsafe.Add(mBase, uint32(l0))) = l2
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			v29 = base.F64_mul(l2, v28)
			v30 = base.F64_abs(v29)
			v31 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(v30, v31)&base.F64_ne(base.F64_abs(v28), v31) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v37 = float64(0)
				if base.F64_eq(v29, v37)&base.F64_ne(v28, v37) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v43 = math.Float64frombits(uint64(0x7ff0000000000000))
					v45 = base.F64_sub(v27, v29)
					if base.B2i32(base.F64_eq(base.F64_abs(v27), v43)|base.F64_ne(base.F64_abs(v45), v43) == int32(0))&base.F64_ne(v30, v43) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v45
						if base.F64_eq(v45, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
						} else {
						}
						return
					}
				}
			}
		}
	}
}
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v37 int32
	_ = v37
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v88 float64
	_ = v88
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v108 float64
	_ = v108
	var v118 float64
	_ = v118
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v155 float64
	_ = v155
	var v160 float64
	_ = v160
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v178 float64
	_ = v178
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v187 float64
	_ = v187
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v205 float64
	_ = v205
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v226 float64
	_ = v226
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v239 float64
	_ = v239
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v247 float64
	_ = v247
	var v248 float64
	_ = v248
	var v249 float64
	_ = v249
	var v250 int32
	_ = v250
	var v260 float64
	_ = v260
	var v261 int32
	_ = v261
	var v262 float64
	_ = v262
	var v267 float64
	_ = v267
	var v268 int32
	_ = v268
	var v283 float64
	_ = v283
	var v284 int32
	_ = v284
	var v285 float64
	_ = v285
	var v286 float64
	_ = v286
	var v287 float64
	_ = v287
	var v297 float64
	_ = v297
	var v302 float64
	_ = v302
	var v303 float64
	_ = v303
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v311 float64
	_ = v311
	var v319 float64
	_ = v319
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v342 float64
	_ = v342
	var v343 float64
	_ = v343
	var v344 float64
	_ = v344
	var v353 float64
	_ = v353
	var v358 float64
	_ = v358
	var v360 float64
	_ = v360
	var v361 float64
	_ = v361
	var v376 float64
	_ = v376
	var v377 float64
	_ = v377
	var v378 float64
	_ = v378
	var v385 float64
	_ = v385
	var v394 float64
	_ = v394
	var v396 float64
	_ = v396
	var v397 float64
	_ = v397
	var v406 float64
	_ = v406
	var v416 float64
	_ = v416
	var v421 float64
	_ = v421
	var v429 float64
	_ = v429
	var v433 float64
	_ = v433
	var v441 float64
	_ = v441
	var v442 float64
	_ = v442
	var v458 int32
	_ = v458
	var v461 float64
	_ = v461
	var v464 float64
	_ = v464
	var v466 float64
	_ = v466
	var v469 float64
	_ = v469
	var v488 int32
	_ = v488
	var v510 int32
	_ = v510
	var v529 int32
	_ = v529
	var v548 int32
	_ = v548
	v15 = int32(0)
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v23 = base.F64_abs(v22)
	if base.F64_le(v23, float64(1e-06)) == v15 {
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		v30 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
		v31 = base.F64_div(v30, v22)
		v32 = base.F64_abs(v31)
		v33 = math.Float64frombits(uint64(0x7ff0000000000000))
		v37 = base.F64_ne(base.F64_abs(v30), v33)
		if base.F64_eq(v32, v33)&v37 != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v510 = m.ExcPending
			if v510 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v39 = float64(0)
			v40 = base.F64_ne(v30, v39)
			v42 = base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000)))
			v44 = base.F64_ne(v31, v39)
			if v40&base.B2i32(v42|v44 == int32(0)) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v548 = m.ExcPending
				if v548 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v50 = math.Float64frombits(uint64(0x7ff0000000000000))
				v51 = base.F64_eq(base.F64_abs(v28), v50)
				v52 = base.F64_mul(v28, v31)
				if base.B2i32(v51|base.F64_ne(base.F64_abs(v52), v50) == int32(0))&base.F64_ne(v32, v50) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v510 = m.ExcPending
					if v510 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v62 = float64(0)
					v63 = base.F64_eq(v28, v62)
					if v44&base.B2i32(v63|base.F64_ne(v52, v62) == int32(0)) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v548 = m.ExcPending
						if v548 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_eq(v52, v29)|base.F64_le(base.F64_abs(base.F64_sub(v29, v52)), float64(1e-06)) != 0 {
							v488 = v15
							return v488
						} else {
							v76 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
							v77 = base.F64_mul(v22, v76)
							v78 = base.F64_abs(v77)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.B2i32(base.F64_ne(v78, v79)|v42 == int32(0))&base.F64_ne(base.F64_abs(v76), v79) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v510 = m.ExcPending
								if v510 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v88 = float64(0)
								if base.F64_eq(v77, v88)&base.F64_ne(v76, v88) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v548 = m.ExcPending
									if v548 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
									v97 = base.F64_mul(v30, v96)
									v98 = base.F64_abs(v97)
									if base.B2i32(base.F64_eq(base.F64_abs(v30), v94)|base.F64_ne(v98, v94) == int32(0))&base.F64_ne(base.F64_abs(v96), v94) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v510 = m.ExcPending
										if v510 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v108 = float64(0)
										if base.B2i32(base.F64_eq(v30, v108)|base.F64_ne(v97, v108) == int32(0))&base.F64_ne(v96, v108) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v548 = m.ExcPending
											if v548 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v118 = math.Float64frombits(uint64(0x7ff0000000000000))
											v120 = base.F64_sub(v77, v97)
											v121 = base.F64_abs(v120)
											if base.B2i32(base.F64_eq(v78, v118)|base.F64_ne(v121, v118) == int32(0))&base.F64_ne(v98, v118) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v510 = m.ExcPending
												if v510 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v130 = base.F64_mul(v28, v30)
												v131 = base.F64_abs(v130)
												if v37&base.B2i32(base.F64_ne(v131, math.Float64frombits(uint64(0x7ff0000000000000)))|v51 == int32(0)) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v510 = m.ExcPending
													if v510 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if v40&base.B2i32(base.F64_ne(v130, float64(0))|v63 == int32(0)) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v548 = m.ExcPending
														if v548 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v144 = base.F64_mul(v22, v29)
														v145 = base.F64_abs(v144)
														v146 = math.Float64frombits(uint64(0x7ff0000000000000))
														if base.B2i32(base.F64_ne(v145, v146)|v42 == int32(0))&base.F64_ne(base.F64_abs(v29), v146) != 0 {
															F_float_overflow_error(m)
															mBase = m.M
															v510 = m.ExcPending
															if v510 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v155 = float64(0)
															if base.F64_eq(v144, v155)&base.F64_ne(v29, v155) != 0 {
																F_float_underflow_error(m)
																mBase = m.M
																v548 = m.ExcPending
																if v548 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v160 = math.Float64frombits(uint64(0x7ff0000000000000))
																v162 = base.F64_sub(v130, v144)
																v163 = base.F64_abs(v162)
																if base.B2i32(base.F64_eq(v145, v160)|base.F64_ne(v163, v160) == int32(0))&base.F64_ne(v131, v160) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v510 = m.ExcPending
																	if v510 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v121)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v162, float64(0)) != 0 {
																		F_float_zero_divide_error(m)
																		mBase = m.M
																		v529 = m.ExcPending
																		if v529 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v178 = base.F64_div(v120, v162)
																		v179 = base.F64_abs(v178)
																		v180 = math.Float64frombits(uint64(0x7ff0000000000000))
																		if base.F64_eq(v179, v180)&base.F64_ne(v121, v180) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v510 = m.ExcPending
																			if v510 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v187 = float64(0)
																			if base.B2i32(base.F64_eq(v163, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v178, v187) == int32(0))&base.F64_ne(v120, v187) != 0 {
																				F_float_underflow_error(m)
																				mBase = m.M
																				v548 = m.ExcPending
																				if v548 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			} else {
																				v195 = base.F64_mul(v28, v178)
																				v196 = base.F64_abs(v195)
																				v197 = math.Float64frombits(uint64(0x7ff0000000000000))
																				if base.B2i32(base.F64_ne(v196, v197)|v51 == int32(0))&base.F64_ne(v179, v197) != 0 {
																					F_float_overflow_error(m)
																					mBase = m.M
																					v510 = m.ExcPending
																					if v510 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				} else {
																					v205 = float64(0)
																					if base.B2i32(base.F64_ne(v195, v205)|v63 == int32(0))&base.F64_ne(v178, v205) != 0 {
																						F_float_underflow_error(m)
																						mBase = m.M
																						v548 = m.ExcPending
																						if v548 != 0 {
																							return int32(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					} else {
																						v213 = base.F64_add(v96, v195)
																						if base.F64_eq(base.F64_abs(v213), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																							v218 = math.Float64frombits(uint64(0x7ff0000000000000))
																							if base.F64_ne(base.F64_abs(v96), v218)&base.F64_ne(v196, v218) != 0 {
																								F_float_overflow_error(m)
																								mBase = m.M
																								v510 = m.ExcPending
																								if v510 != 0 {
																									return int32(0)
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							} else {
																								v230 = base.F64_div(base.F64_neg(v213), v22)
																								v233 = float64(0)
																								if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v230, v233)|base.F64_eq(v213, v233) != 0 {
																									v441 = v178
																									v442 = v230
																									v458 = int32(1)
																									if l0 == int32(0) {
																										v488 = v458
																									} else {
																										v461 = float64(0)
																										if base.F64_eq(v442, v461) != 0 {
																											v464 = v461
																										} else {
																											v464 = v442
																										}
																										*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																										v466 = float64(0)
																										if base.F64_eq(v441, v466) != 0 {
																											v469 = v466
																										} else {
																											v469 = v441
																										}
																										*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																										v488 = v458
																									}
																									return v488
																								} else {
																									F_float_underflow_error(m)
																									mBase = m.M
																									v548 = m.ExcPending
																									if v548 != 0 {
																										return int32(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v226 = base.F64_div(base.F64_neg(v213), v22)
																							if base.F64_eq(base.F64_abs(v226), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																								F_float_overflow_error(m)
																								mBase = m.M
																								v510 = m.ExcPending
																								if v510 != 0 {
																									return int32(0)
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							} else {
																								v230 = v226
																								v233 = float64(0)
																								if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v230, v233)|base.F64_eq(v213, v233) != 0 {
																									v441 = v178
																									v442 = v230
																									v458 = int32(1)
																									if l0 == int32(0) {
																										v488 = v458
																									} else {
																										v461 = float64(0)
																										if base.F64_eq(v442, v461) != 0 {
																											v464 = v461
																										} else {
																											v464 = v442
																										}
																										*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																										v466 = float64(0)
																										if base.F64_eq(v441, v466) != 0 {
																											v469 = v466
																										} else {
																											v469 = v441
																										}
																										*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																										v488 = v458
																									}
																									return v488
																								} else {
																									F_float_underflow_error(m)
																									mBase = m.M
																									v548 = m.ExcPending
																									if v548 != 0 {
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
		v239 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
		v240 = base.F64_abs(v239)
		if base.F64_le(v240, float64(1e-06)) != 0 {
			v488 = v15
			return v488
		} else {
			v243 = base.F64_div(v22, v239)
			if base.F64_eq(base.F64_abs(v243), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v510 = m.ExcPending
				if v510 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v247 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v248 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				v249 = float64(0)
				v250 = base.F64_eq(v22, v249)
				if base.B2i32(v250|base.F64_ne(v243, v249) == int32(0))&base.F64_ne(v240, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v548 = m.ExcPending
					if v548 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v260 = math.Float64frombits(uint64(0x7ff0000000000000))
					v261 = base.F64_ne(base.F64_abs(v247), v260)
					v262 = base.F64_mul(v243, v247)
					if v261&base.F64_eq(base.F64_abs(v262), v260) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v510 = m.ExcPending
						if v510 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v267 = float64(0)
						v268 = base.F64_ne(v247, v267)
						if v268&base.B2i32(base.F64_eq(v243, v267)|base.F64_ne(v262, v267) == int32(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v548 = m.ExcPending
							if v548 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if base.F64_eq(v248, v262)|base.F64_le(base.F64_abs(base.F64_sub(v248, v262)), float64(1e-06)) != 0 {
								v488 = v15
								return v488
							} else {
								v283 = math.Float64frombits(uint64(0x7ff0000000000000))
								v284 = base.F64_eq(v240, v283)
								v285 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								v286 = base.F64_mul(v239, v285)
								v287 = base.F64_abs(v286)
								if base.B2i32(v284|base.F64_ne(v287, v283) == int32(0))&base.F64_ne(base.F64_abs(v285), v283) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v510 = m.ExcPending
									if v510 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v297 = float64(0)
									if base.F64_eq(v286, v297)&base.F64_ne(v285, v297) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v548 = m.ExcPending
										if v548 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v302 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
										v303 = base.F64_mul(v22, v302)
										v304 = base.F64_abs(v303)
										v305 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(v304, v305)&base.F64_ne(base.F64_abs(v302), v305) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v510 = m.ExcPending
											if v510 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v311 = float64(0)
											if base.B2i32(base.F64_ne(v303, v311)|v250 == int32(0))&base.F64_ne(v302, v311) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v548 = m.ExcPending
												if v548 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v319 = math.Float64frombits(uint64(0x7ff0000000000000))
												v321 = base.F64_sub(v286, v303)
												v322 = base.F64_abs(v321)
												if base.B2i32(base.F64_eq(v287, v319)|base.F64_ne(v322, v319) == int32(0))&base.F64_ne(v304, v319) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v510 = m.ExcPending
													if v510 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v331 = base.F64_mul(v22, v247)
													v332 = base.F64_abs(v331)
													if v261&base.F64_eq(v332, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_overflow_error(m)
														mBase = m.M
														v510 = m.ExcPending
														if v510 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														if v268&base.B2i32(base.F64_ne(v331, float64(0))|v250 == int32(0)) != 0 {
															F_float_underflow_error(m)
															mBase = m.M
															v548 = m.ExcPending
															if v548 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v342 = base.F64_mul(v239, v248)
															v343 = base.F64_abs(v342)
															v344 = math.Float64frombits(uint64(0x7ff0000000000000))
															if base.B2i32(base.F64_ne(v343, v344)|v284 == int32(0))&base.F64_ne(base.F64_abs(v248), v344) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v510 = m.ExcPending
																if v510 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v353 = float64(0)
																if base.F64_eq(v342, v353)&base.F64_ne(v248, v353) != 0 {
																	F_float_underflow_error(m)
																	mBase = m.M
																	v548 = m.ExcPending
																	if v548 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v358 = math.Float64frombits(uint64(0x7ff0000000000000))
																	v360 = base.F64_sub(v331, v342)
																	v361 = base.F64_abs(v360)
																	if base.B2i32(base.F64_eq(v343, v358)|base.F64_ne(v361, v358) == int32(0))&base.F64_ne(v332, v358) != 0 {
																		F_float_overflow_error(m)
																		mBase = m.M
																		v510 = m.ExcPending
																		if v510 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v322)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v360, float64(0)) != 0 {
																			F_float_zero_divide_error(m)
																			mBase = m.M
																			v529 = m.ExcPending
																			if v529 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v376 = base.F64_div(v321, v360)
																			v377 = base.F64_abs(v376)
																			v378 = math.Float64frombits(uint64(0x7ff0000000000000))
																			if base.F64_eq(v377, v378)&base.F64_ne(v322, v378) != 0 {
																				F_float_overflow_error(m)
																				mBase = m.M
																				v510 = m.ExcPending
																				if v510 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			} else {
																				v385 = float64(0)
																				if base.B2i32(base.F64_eq(v361, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v376, v385) == int32(0))&base.F64_ne(v321, v385) != 0 {
																					F_float_underflow_error(m)
																					mBase = m.M
																					v548 = m.ExcPending
																					if v548 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				} else {
																					v394 = math.Float64frombits(uint64(0x7ff0000000000000))
																					v396 = base.F64_mul(v247, v376)
																					v397 = base.F64_abs(v396)
																					if base.B2i32(base.F64_eq(base.F64_abs(v247), v394)|base.F64_ne(v397, v394) == int32(0))&base.F64_ne(v377, v394) != 0 {
																						F_float_overflow_error(m)
																						mBase = m.M
																						v510 = m.ExcPending
																						if v510 != 0 {
																							return int32(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					} else {
																						v406 = float64(0)
																						if base.B2i32(base.F64_eq(v247, v406)|base.F64_ne(v396, v406) == int32(0))&base.F64_ne(v376, v406) != 0 {
																							F_float_underflow_error(m)
																							mBase = m.M
																							v548 = m.ExcPending
																							if v548 != 0 {
																								return int32(0)
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						} else {
																							v416 = base.F64_add(v302, v396)
																							if base.F64_eq(base.F64_abs(v416), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																								v421 = math.Float64frombits(uint64(0x7ff0000000000000))
																								if base.F64_ne(base.F64_abs(v302), v421)&base.F64_ne(v397, v421) != 0 {
																									F_float_overflow_error(m)
																									mBase = m.M
																									v510 = m.ExcPending
																									if v510 != 0 {
																										return int32(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								} else {
																									v433 = base.F64_div(base.F64_neg(v416), v239)
																									if base.F64_eq(v240, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v433, float64(0)) != 0 {
																										v441 = v376
																										v442 = v433
																										v458 = int32(1)
																										if l0 == int32(0) {
																											v488 = v458
																										} else {
																											v461 = float64(0)
																											if base.F64_eq(v442, v461) != 0 {
																												v464 = v461
																											} else {
																												v464 = v442
																											}
																											*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																											v466 = float64(0)
																											if base.F64_eq(v441, v466) != 0 {
																												v469 = v466
																											} else {
																												v469 = v441
																											}
																											*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																											v488 = v458
																										}
																										return v488
																									} else {
																										if base.F64_ne(v416, float64(0)) != 0 {
																											F_float_underflow_error(m)
																											mBase = m.M
																											v548 = m.ExcPending
																											if v548 != 0 {
																												return int32(0)
																											} else {
																												base.Wasm_trap_unreachable()
																												for {
																												}
																											}
																										} else {
																											v441 = v376
																											v442 = v433
																											v458 = int32(1)
																											if l0 == int32(0) {
																												v488 = v458
																											} else {
																												v461 = float64(0)
																												if base.F64_eq(v442, v461) != 0 {
																													v464 = v461
																												} else {
																													v464 = v442
																												}
																												*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																												v466 = float64(0)
																												if base.F64_eq(v441, v466) != 0 {
																													v469 = v466
																												} else {
																													v469 = v441
																												}
																												*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																												v488 = v458
																											}
																											return v488
																										}
																									}
																								}
																							} else {
																								v429 = base.F64_div(base.F64_neg(v416), v239)
																								if base.F64_eq(base.F64_abs(v429), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																									F_float_overflow_error(m)
																									mBase = m.M
																									v510 = m.ExcPending
																									if v510 != 0 {
																										return int32(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								} else {
																									v433 = v429
																									if base.F64_eq(v240, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v433, float64(0)) != 0 {
																										v441 = v376
																										v442 = v433
																										v458 = int32(1)
																										if l0 == int32(0) {
																											v488 = v458
																										} else {
																											v461 = float64(0)
																											if base.F64_eq(v442, v461) != 0 {
																												v464 = v461
																											} else {
																												v464 = v442
																											}
																											*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																											v466 = float64(0)
																											if base.F64_eq(v441, v466) != 0 {
																												v469 = v466
																											} else {
																												v469 = v441
																											}
																											*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																											v488 = v458
																										}
																										return v488
																									} else {
																										if base.F64_ne(v416, float64(0)) != 0 {
																											F_float_underflow_error(m)
																											mBase = m.M
																											v548 = m.ExcPending
																											if v548 != 0 {
																												return int32(0)
																											} else {
																												base.Wasm_trap_unreachable()
																												for {
																												}
																											}
																										} else {
																											v441 = v376
																											v442 = v433
																											v458 = int32(1)
																											if l0 == int32(0) {
																												v488 = v458
																											} else {
																												v461 = float64(0)
																												if base.F64_eq(v442, v461) != 0 {
																													v464 = v461
																												} else {
																													v464 = v442
																												}
																												*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v464
																												v466 = float64(0)
																												if base.F64_eq(v441, v466) != 0 {
																													v469 = v466
																												} else {
																													v469 = v441
																												}
																												*(*float64)(unsafe.Add(mBase, uint32(l0))) = v469
																												v488 = v458
																											}
																											return v488
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
			}
		}
	}
}
func F_line_intersect(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_line_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_line_perp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v17 float64
	_ = v17
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v31 int32
	_ = v31
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = base.F64_abs(v13)
	if base.F64_le(v14, float64(1e-06)) != 0 {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		return base.F64_le(base.F64_abs(v17), float64(1e-06))
	} else {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		v23 = base.F64_abs(v22)
		v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v25 = base.F64_abs(v24)
		if base.F64_le(v25, float64(1e-06)) != 0 {
			return base.F64_le(v23, float64(1e-06))
		} else {
			v31 = int32(0)
			if base.F64_le(v23, float64(1e-06)) != 0 {
				v86 = v31
				return v86
			} else {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v35 = base.F64_abs(v34)
				if base.F64_le(v35, float64(1e-06)) != 0 {
					v86 = v31
					return v86
				} else {
					v38 = math.Float64frombits(uint64(0x7ff0000000000000))
					v40 = base.F64_mul(v13, v24)
					v41 = base.F64_abs(v40)
					if base.B2i32(base.F64_eq(v14, v38)|base.F64_ne(v41, v38) == int32(0))&base.F64_ne(v25, v38) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_eq(v40, float64(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v52 = math.Float64frombits(uint64(0x7ff0000000000000))
							v54 = base.F64_mul(v22, v34)
							v55 = base.F64_abs(v54)
							if base.B2i32(base.F64_eq(v23, v52)|base.F64_ne(v55, v52) == int32(0))&base.F64_ne(v35, v52) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_eq(v54, float64(0)) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v66 = base.F64_div(v40, v54)
									v68 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v66), v68)&base.F64_ne(v41, v68) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_eq(v66, float64(0))&base.F64_ne(v55, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v86 = base.F64_eq(v66, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v66, float64(1))), float64(1e-06))
											return v86
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
