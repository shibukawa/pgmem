package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamp(m *base.Module, l0 int32) int32 {
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
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = F_make_timestamp_internal(m, v2, v3, v4, v5, v6, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
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
							F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(607), int32(_a_F_make_timestamp_internal_2))
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
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(616), int32(_a_F_make_timestamp_internal_2))
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
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
												F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(636), int32(_a_F_make_timestamp_internal_2))
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
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(607), int32(_a_F_make_timestamp_internal_2))
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
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(616), int32(_a_F_make_timestamp_internal_2))
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
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(636), int32(_a_F_make_timestamp_internal_2))
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
								F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(607), int32(_a_F_make_timestamp_internal_2))
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
									F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(616), int32(_a_F_make_timestamp_internal_2))
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
										F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
											F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(628), int32(_a_F_make_timestamp_internal_2))
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
												F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(636), int32(_a_F_make_timestamp_internal_2))
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
					F_errfinish(m, int32(_a_F_make_timestamp_internal_1), int32(601), int32(_a_F_make_timestamp_internal_2))
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
func F_timestamp_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if v10 == int64(-9223372036854775807-1) {
		v60 = int32(-2147483648)
		m.G0 = v6 + int32(48)
		return v60
	} else {
		if v10 == int64(9223372036854775807) {
			v60 = int32(2147483647)
			m.G0 = v6 + int32(48)
			return v60
		} else {
			v16 = int32(0)
			v21 = F_timestamp2tm(m, v10, v16, v6+int32(4), v6, v16, v16)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_date_0), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_date_1), int32(1379), int32(_a_F_timestamp_date_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
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
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					v32 = base.B2i32(int32(2) < v26)
					if int32(2) < v26 {
						v33 = int32(_a_F_timestamp_date_3)
					} else {
						v33 = int32(_a_F_timestamp_date_4)
					}
					v34 = v33 + v25
					v39 = base.I32_div_s(v34, int32(4))
					v42 = base.I32_div_s(v34, int32(-100))
					v45 = base.I32_div_s(v34, int32(400))
					if int32(2) < v26 {
						v49 = int32(1)
					} else {
						v49 = int32(13)
					}
					v54 = base.I32_div_s((v49+v26)*int32(_a_F_timestamp_date_5), int32(256))
					v60 = v27 + v34*int32(365) + v39 + v42 + v45 + v54 - int32(_a_F_timestamp_date_6) - int32(_a_F_timestamp_date_7)
					m.G0 = v6 + int32(48)
					return v60
				}
			}
		}
	}
}
func F_timestamp_gt_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
			}
		}
	}
}
func F_timestamp_le_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 == int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 != int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v15 <= v10)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamp_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v35 int64
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v10 = F_palloc(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int64(9223372036854775807)
		if base.B2i32(base.Ui64(int64(1)) < base.Ui64(v6-v14))&base.B2i32(base.Ui64(int64(2)) <= base.Ui64(v8-v14)) == int32(0) {
			if v8 != int64(9223372036854775807) {
				if v8 != int64(-9223372036854775807-1) {
					if v6 == int64(-9223372036854775807-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(9223372034707292159)
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(9223372036854775807)
						return v10
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(-9223372034707292160)
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(-9223372036854775807 - 1)
						return v10
					}
				} else {
					if v6 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2866), int32(_a_F_timestamp_mi_2))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
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
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(-9223372034707292160)
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(-9223372036854775807 - 1)
						return v10
					}
				}
			} else {
				if v6 == int64(9223372036854775807) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2875), int32(_a_F_timestamp_mi_2))
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
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(9223372034707292159)
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(9223372036854775807)
					return v10
				}
			}
		} else {
			v35 = v8 - v6
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v35
			if base.B2i32(int64(0) < v6) != base.B2i32(v35 < v8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp_mi_0), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_mi_1), int32(2890), int32(_a_F_timestamp_mi_2))
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
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
				v45 = F_DirectFunctionCall1Coll(m, int32(1499), int32(0), v10)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					return v45
				}
			}
		}
	}
}
func F_timestamp_mi_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_interval_um_internal(m, v10, v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v17 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v17, v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v19
			}
		}
	}
}
func F_timestamp_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(1495)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(1496)
	v8 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v2))) = v8
		v14 = F_Int64GetDatum(m, int64(9223372036854775807))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v2)+4)) = v14
			return int32(0)
		}
	}
}
