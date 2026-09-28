package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
	v8 = F_make_timestamp_internal(m, v2, v3, v4, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return v8
	}
}
func F_make_timestamp_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v262 float64
	_ = v262
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v309 int64
	_ = v309
	var v316 int64
	_ = v316
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v336 int32
	_ = v336
	var v344 int64
	_ = v344
	var v347 int64
	_ = v347
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	v13 = m.G0
	v15 = v13 - int32(176)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = l2
	if l0 < int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = int32(0) - l0
	} else {
	}
	v31 = v15 + int32(132)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if int32(base.Ui32(l0)>>(uint(int32(31))%32)) != 0 {
		if int32(0) < v37 {
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1) - v37
			v149 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v149-int32(13)) {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v159-int32(32)) {
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					if v171&int32(3) != 0 {
						v181 = int32(0)
					} else {
						v176 = base.I32_rem_s(v171, int32(100))
						if v176 != 0 {
							v181 = int32(1)
						} else {
							v178 = base.I32_rem_s(v171, int32(400))
							v181 = base.B2i32(v178 == int32(0))
						}
					}
					v184 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v181*int32(52)+v184<<(uint(int32(2))%32))+uint32(_c_F_make_timestamp_internal[0])))
					if v169 <= v190 {
						v199 = int32(0)
					} else {
						v199 = int32(-2)
					}
				} else {
					v199 = int32(-3)
				}
			} else {
				v199 = int32(-3)
			}
		} else {
			v199 = int32(-2)
		}
	} else {
		if int32(0) < v37 {
			v149 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v149-int32(13)) {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v159-int32(32)) {
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					if v171&int32(3) != 0 {
						v181 = int32(0)
					} else {
						v176 = base.I32_rem_s(v171, int32(100))
						if v176 != 0 {
							v181 = int32(1)
						} else {
							v178 = base.I32_rem_s(v171, int32(400))
							v181 = base.B2i32(v178 == int32(0))
						}
					}
					v184 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v181*int32(52)+v184<<(uint(int32(2))%32))+uint32(_c_F_make_timestamp_internal[0])))
					if v169 <= v190 {
						v199 = int32(0)
					} else {
						v199 = int32(-2)
					}
				} else {
					v199 = int32(-3)
				}
			} else {
				v199 = int32(-3)
			}
		} else {
			v199 = int32(-2)
		}
	}
	if v199 == int32(0) {
		v202 = *(*int32)(unsafe.Add(mBase, uint32(v15)+148))
		v203 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
		if v203 <= int32(-4713) {
			if v203 != int32(-4713) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v451 = m.ExcPending
				if v451 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v454 = m.ExcPending
					if v454 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
						F_errmsg(m, int32(_a_F_make_timestamp_internal_0), v15+int32(96))
						mBase = m.M
						v462 = m.ExcPending
						if v462 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(603), int32(_a_F_make_timestamp_internal_2))
							mBase = m.M
							v467 = m.ExcPending
							if v467 != 0 {
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
				if int32(10) < v202 {
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
					v222 = base.B2i32(int32(2) < v202)
					if int32(2) < v202 {
						v223 = int32(_a_F_make_timestamp_internal_3)
					} else {
						v223 = int32(_a_F_make_timestamp_internal_4)
					}
					v224 = v223 + v203
					v229 = base.I32_div_s(v224, int32(4))
					v232 = base.I32_div_s(v224, int32(-100))
					v235 = base.I32_div_s(v224, int32(400))
					if int32(2) < v202 {
						v239 = int32(1)
					} else {
						v239 = int32(13)
					}
					v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_timestamp_internal_5), int32(256))
					v248 = int32(1)
					if base.B2i32(base.Ui32(int32(24)) < base.Ui32(l3))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(l4))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807))) != 0 {
						v280 = v248
					} else {
						v262 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
						if base.F64_lt(v262, float64(0))|base.F64_gt(v262, float64(6e+07)) != 0 {
							v280 = v248
						} else {
							v269 = int32(60)
							v280 = base.B2i32(int64(86400000000) < base.I64_trunc_sat_f64_s(v262)+base.I64_extend_i32_u((l3*v269+l4)*v269)*int64(1000000))
						}
					}
					if v280 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v383 = m.ExcPending
						if v383 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v386 = m.ExcPending
							if v386 != 0 {
								return int64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
								F_errmsg(m, int32(_a_F_make_timestamp_internal_6), v15)
								mBase = m.M
								v392 = m.ExcPending
								if v392 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(612), int32(_a_F_make_timestamp_internal_2))
									mBase = m.M
									v397 = m.ExcPending
									if v397 != 0 {
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
						v283 = v15 + int32(80)
						v286 = base.I64_extend_i32_s(v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_timestamp_internal_7) - int32(_a_F_make_timestamp_internal_8))
						v295 = int64(32)
						v296 = int64(20)
						v298 = int64(base.Ui64(v286) >> (uint(v295) % 64))
						v301 = int64(4294967295)
						v302 = int64(500654080)
						v304 = v286 & v301
						v305 = v302 * v304
						v309 = int64(base.Ui64(v305)>>(uint(v295)%64)) + v302*v298
						v316 = v304*v296 + v309&v301
						*(*int64)(unsafe.Add(mBase, uint32(v283)+8)) = v286*int64(0) + v286>>(uint(int64(63))%64)*int64(86400000000) + v296*v298 + int64(base.Ui64(v309)>>(uint(v295)%64)) + int64(base.Ui64(v316)>>(uint(v295)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v283))) = v305&v301 | v316<<(uint(v295)%64)
						v327 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
						v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
						if v327 != v328>>(uint(int64(63))%64) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v403 = m.ExcPending
							if v403 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
									F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
									mBase = m.M
									v417 = m.ExcPending
									if v417 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
										mBase = m.M
										v422 = m.ExcPending
										if v422 != 0 {
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
							v336 = int32(60)
							v344 = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(l5, float64(1e+06)))) + base.I64_extend_i32_s((l3*v336+l4)*v336)*int64(1000000)
							v347 = v328 + v344
							if base.B2i32(v344 < int64(0)) != base.B2i32(v347 < v328) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v403 = m.ExcPending
								if v403 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v406 = m.ExcPending
									if v406 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
										F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
										mBase = m.M
										v417 = m.ExcPending
										if v417 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
											mBase = m.M
											v422 = m.ExcPending
											if v422 != 0 {
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
								if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v347+int64(211813488000000000)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v426 = m.ExcPending
									if v426 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v429 = m.ExcPending
										if v429 != 0 {
											return int64(0)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
											F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(48))
											mBase = m.M
											v442 = m.ExcPending
											if v442 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(632), int32(_a_F_make_timestamp_internal_2))
												mBase = m.M
												v447 = m.ExcPending
												if v447 != 0 {
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
									m.G0 = v15 + int32(176)
									return v347
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v451 = m.ExcPending
					if v451 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v454 = m.ExcPending
						if v454 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
							F_errmsg(m, int32(_a_F_make_timestamp_internal_0), v15+int32(96))
							mBase = m.M
							v462 = m.ExcPending
							if v462 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(603), int32(_a_F_make_timestamp_internal_2))
								mBase = m.M
								v467 = m.ExcPending
								if v467 != 0 {
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
		} else {
			if v203 < int32(_a_F_make_timestamp_internal_10) {
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
				v222 = base.B2i32(int32(2) < v202)
				if int32(2) < v202 {
					v223 = int32(_a_F_make_timestamp_internal_3)
				} else {
					v223 = int32(_a_F_make_timestamp_internal_4)
				}
				v224 = v223 + v203
				v229 = base.I32_div_s(v224, int32(4))
				v232 = base.I32_div_s(v224, int32(-100))
				v235 = base.I32_div_s(v224, int32(400))
				if int32(2) < v202 {
					v239 = int32(1)
				} else {
					v239 = int32(13)
				}
				v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_timestamp_internal_5), int32(256))
				v248 = int32(1)
				if base.B2i32(base.Ui32(int32(24)) < base.Ui32(l3))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(l4))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807))) != 0 {
					v280 = v248
				} else {
					v262 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
					if base.F64_lt(v262, float64(0))|base.F64_gt(v262, float64(6e+07)) != 0 {
						v280 = v248
					} else {
						v269 = int32(60)
						v280 = base.B2i32(int64(86400000000) < base.I64_trunc_sat_f64_s(v262)+base.I64_extend_i32_u((l3*v269+l4)*v269)*int64(1000000))
					}
				}
				if v280 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v383 = m.ExcPending
					if v383 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v386 = m.ExcPending
						if v386 != 0 {
							return int64(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
							F_errmsg(m, int32(_a_F_make_timestamp_internal_6), v15)
							mBase = m.M
							v392 = m.ExcPending
							if v392 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(612), int32(_a_F_make_timestamp_internal_2))
								mBase = m.M
								v397 = m.ExcPending
								if v397 != 0 {
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
					v283 = v15 + int32(80)
					v286 = base.I64_extend_i32_s(v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_timestamp_internal_7) - int32(_a_F_make_timestamp_internal_8))
					v295 = int64(32)
					v296 = int64(20)
					v298 = int64(base.Ui64(v286) >> (uint(v295) % 64))
					v301 = int64(4294967295)
					v302 = int64(500654080)
					v304 = v286 & v301
					v305 = v302 * v304
					v309 = int64(base.Ui64(v305)>>(uint(v295)%64)) + v302*v298
					v316 = v304*v296 + v309&v301
					*(*int64)(unsafe.Add(mBase, uint32(v283)+8)) = v286*int64(0) + v286>>(uint(int64(63))%64)*int64(86400000000) + v296*v298 + int64(base.Ui64(v309)>>(uint(v295)%64)) + int64(base.Ui64(v316)>>(uint(v295)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v283))) = v305&v301 | v316<<(uint(v295)%64)
					v327 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
					v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
					if v327 != v328>>(uint(int64(63))%64) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v403 = m.ExcPending
						if v403 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v406 = m.ExcPending
							if v406 != 0 {
								return int64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
								F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
									mBase = m.M
									v422 = m.ExcPending
									if v422 != 0 {
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
						v336 = int32(60)
						v344 = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(l5, float64(1e+06)))) + base.I64_extend_i32_s((l3*v336+l4)*v336)*int64(1000000)
						v347 = v328 + v344
						if base.B2i32(v344 < int64(0)) != base.B2i32(v347 < v328) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v403 = m.ExcPending
							if v403 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
									F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
									mBase = m.M
									v417 = m.ExcPending
									if v417 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
										mBase = m.M
										v422 = m.ExcPending
										if v422 != 0 {
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
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v347+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v426 = m.ExcPending
								if v426 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
										F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(48))
										mBase = m.M
										v442 = m.ExcPending
										if v442 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(632), int32(_a_F_make_timestamp_internal_2))
											mBase = m.M
											v447 = m.ExcPending
											if v447 != 0 {
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
								m.G0 = v15 + int32(176)
								return v347
							}
						}
					}
				}
			} else {
				if base.B2i32(v203 != int32(_a_F_make_timestamp_internal_10))|base.B2i32(int32(6) <= v202) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v451 = m.ExcPending
					if v451 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v454 = m.ExcPending
						if v454 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
							F_errmsg(m, int32(_a_F_make_timestamp_internal_0), v15+int32(96))
							mBase = m.M
							v462 = m.ExcPending
							if v462 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(603), int32(_a_F_make_timestamp_internal_2))
								mBase = m.M
								v467 = m.ExcPending
								if v467 != 0 {
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
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
					v222 = base.B2i32(int32(2) < v202)
					if int32(2) < v202 {
						v223 = int32(_a_F_make_timestamp_internal_3)
					} else {
						v223 = int32(_a_F_make_timestamp_internal_4)
					}
					v224 = v223 + v203
					v229 = base.I32_div_s(v224, int32(4))
					v232 = base.I32_div_s(v224, int32(-100))
					v235 = base.I32_div_s(v224, int32(400))
					if int32(2) < v202 {
						v239 = int32(1)
					} else {
						v239 = int32(13)
					}
					v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_timestamp_internal_5), int32(256))
					v248 = int32(1)
					if base.B2i32(base.Ui32(int32(24)) < base.Ui32(l3))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(l4))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807))) != 0 {
						v280 = v248
					} else {
						v262 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
						if base.F64_lt(v262, float64(0))|base.F64_gt(v262, float64(6e+07)) != 0 {
							v280 = v248
						} else {
							v269 = int32(60)
							v280 = base.B2i32(int64(86400000000) < base.I64_trunc_sat_f64_s(v262)+base.I64_extend_i32_u((l3*v269+l4)*v269)*int64(1000000))
						}
					}
					if v280 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v383 = m.ExcPending
						if v383 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v386 = m.ExcPending
							if v386 != 0 {
								return int64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
								F_errmsg(m, int32(_a_F_make_timestamp_internal_6), v15)
								mBase = m.M
								v392 = m.ExcPending
								if v392 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(612), int32(_a_F_make_timestamp_internal_2))
									mBase = m.M
									v397 = m.ExcPending
									if v397 != 0 {
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
						v283 = v15 + int32(80)
						v286 = base.I64_extend_i32_s(v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_timestamp_internal_7) - int32(_a_F_make_timestamp_internal_8))
						v295 = int64(32)
						v296 = int64(20)
						v298 = int64(base.Ui64(v286) >> (uint(v295) % 64))
						v301 = int64(4294967295)
						v302 = int64(500654080)
						v304 = v286 & v301
						v305 = v302 * v304
						v309 = int64(base.Ui64(v305)>>(uint(v295)%64)) + v302*v298
						v316 = v304*v296 + v309&v301
						*(*int64)(unsafe.Add(mBase, uint32(v283)+8)) = v286*int64(0) + v286>>(uint(int64(63))%64)*int64(86400000000) + v296*v298 + int64(base.Ui64(v309)>>(uint(v295)%64)) + int64(base.Ui64(v316)>>(uint(v295)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v283))) = v305&v301 | v316<<(uint(v295)%64)
						v327 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
						v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
						if v327 != v328>>(uint(int64(63))%64) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v403 = m.ExcPending
							if v403 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
									F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
									mBase = m.M
									v417 = m.ExcPending
									if v417 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
										mBase = m.M
										v422 = m.ExcPending
										if v422 != 0 {
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
							v336 = int32(60)
							v344 = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(l5, float64(1e+06)))) + base.I64_extend_i32_s((l3*v336+l4)*v336)*int64(1000000)
							v347 = v328 + v344
							if base.B2i32(v344 < int64(0)) != base.B2i32(v347 < v328) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v403 = m.ExcPending
								if v403 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v406 = m.ExcPending
									if v406 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
										F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(16))
										mBase = m.M
										v417 = m.ExcPending
										if v417 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(624), int32(_a_F_make_timestamp_internal_2))
											mBase = m.M
											v422 = m.ExcPending
											if v422 != 0 {
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
								if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v347+int64(211813488000000000)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v426 = m.ExcPending
									if v426 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v429 = m.ExcPending
										if v429 != 0 {
											return int64(0)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
											F_errmsg(m, int32(_a_F_make_timestamp_internal_9), v15+int32(48))
											mBase = m.M
											v442 = m.ExcPending
											if v442 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(632), int32(_a_F_make_timestamp_internal_2))
												mBase = m.M
												v447 = m.ExcPending
												if v447 != 0 {
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
									m.G0 = v15 + int32(176)
									return v347
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
		v363 = m.ExcPending
		if v363 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v366 = m.ExcPending
			if v366 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = l0
				F_errmsg(m, int32(_a_F_make_timestamp_internal_11), v15+int32(112))
				mBase = m.M
				v374 = m.ExcPending
				if v374 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(597), int32(_a_F_make_timestamp_internal_2))
					mBase = m.M
					v379 = m.ExcPending
					if v379 != 0 {
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
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32 {
	return base.B2i32(l1 < l0) - base.B2i32(l0 < l1)
}
func F_timestamp_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_timestamp2date_safe(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			return base.I64_extend_i32_s(v6)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(453) {
				return base.I64_extend_i32_s(v6)
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
				if v16 != int32(1) {
					return base.I64_extend_i32_s(v6)
				} else {
					v19 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
					return int64(0)
				}
			}
		}
	}
}
func F_timestamp_gt_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(int32(base.Ui32(v4) >> (uint(int32(31)) % 32)))
	}
}
func F_timestamp_le_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_le_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_le_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v17 <= v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v17 <= v9)
				} else {
					v33 = base.B2i32(v9 == int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 != int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamp_mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v33 int64
	_ = v33
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int64(9223372036854775807)
		if base.B2i32(base.Ui64(int64(1)) < base.Ui64(v5-v12))&base.B2i32(base.Ui64(int64(2)) <= base.Ui64(v6-v12)) == int32(0) {
			if v6 != int64(9223372036854775807) {
				if v6 != int64(-9223372036854775807-1) {
					if v5 == int64(-9223372036854775807-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
						return base.I64_extend_i32_u(v8)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
						return base.I64_extend_i32_u(v8)
					}
				} else {
					if v5 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2861), int32(_a_F_timestamp_mi_2))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
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
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
						return base.I64_extend_i32_u(v8)
					}
				}
			} else {
				if v5 == int64(9223372036854775807) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2870), int32(_a_F_timestamp_mi_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
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
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
					return base.I64_extend_i32_u(v8)
				}
			}
		} else {
			v33 = v6 - v5
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v33
			if base.B2i32(int64(0) < v5) != base.B2i32(v33 < v6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2885), int32(_a_F_timestamp_mi_2))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
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
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
				v44 = F_DirectFunctionCall1Coll(m, int32(1712), int32(0), base.I64_extend_i32_u(v8))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					return v44 & int64(4294967295)
				}
			}
		}
	}
}
func F_timestamp_mi_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_interval_um_internal(m, v9, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v17 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v8, base.I64_extend_i32_u(v6))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func F_timestamp_skipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(1708)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(1709)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-9223372036854775807 - 1)
	return int64(0)
}
