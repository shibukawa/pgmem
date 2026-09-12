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
	var v36 int32
	_ = v36
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v259 float64
	_ = v259
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v310 int64
	_ = v310
	var v317 int64
	_ = v317
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v335 float64
	_ = v335
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v350 int64
	_ = v350
	var v353 int64
	_ = v353
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if int32(base.Ui32(l0)>>(uint(int32(31))%32)) != 0 {
		if int32(0) < v36 {
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1) - v36
			v148 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v148-int32(13)) {
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v158-int32(32)) {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					if v170&int32(3) != 0 {
						v180 = int32(0)
					} else {
						v175 = base.I32_rem_s(v170, int32(100))
						if v175 != 0 {
							v180 = int32(1)
						} else {
							v177 = base.I32_rem_s(v170, int32(400))
							v180 = base.B2i32(v177 == int32(0))
						}
					}
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v180*int32(52)+v183<<(uint(int32(2))%32))+uint32(_consts[940])))
					if v168 <= v189 {
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
		if int32(0) < v36 {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v148-int32(13)) {
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v158-int32(32)) {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					if v170&int32(3) != 0 {
						v180 = int32(0)
					} else {
						v175 = base.I32_rem_s(v170, int32(100))
						if v175 != 0 {
							v180 = int32(1)
						} else {
							v177 = base.I32_rem_s(v170, int32(400))
							v180 = base.B2i32(v177 == int32(0))
						}
					}
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v180*int32(52)+v183<<(uint(int32(2))%32))+uint32(_consts[940])))
					if v168 <= v189 {
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
				v457 = m.ExcPending
				if v457 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v460 = m.ExcPending
					if v460 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
						F_errmsg(m, int32(457620), v15+int32(96))
						mBase = m.M
						v468 = m.ExcPending
						if v468 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(486977), int32(607), int32(305929))
							mBase = m.M
							v473 = m.ExcPending
							if v473 != 0 {
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
					v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
					v221 = base.B2i32(int32(2) < v202)
					if int32(2) < v202 {
						v222 = int32(4800)
					} else {
						v222 = int32(4799)
					}
					v223 = v222 + v203
					v228 = base.I32_div_s(v223, int32(4))
					v231 = base.I32_div_s(v223, int32(-100))
					v234 = base.I32_div_s(v223, int32(400))
					if int32(2) < v202 {
						v238 = int32(1)
					} else {
						v238 = int32(13)
					}
					v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
					v247 = int32(1)
					if base.Ui32(int32(24)) < base.Ui32(l3) {
						v281 = v247
					} else {
						if base.Ui32(int32(59)) < base.Ui32(l4) {
							v281 = v247
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807)) {
								v281 = v247
							} else {
								v259 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
								if base.F64_lt(v259, float64(0)) != 0 {
									v281 = v247
								} else {
									if base.F64_gt(v259, float64(6e+07)) != 0 {
										v281 = v247
									} else {
										if base.F64_lt(base.F64_abs(v259), float64(9.223372036854776e+18)) != 0 {
											v267 = base.I64_trunc_f64_s(v259)
											v269 = v267
										} else {
											v269 = int64(-9223372036854775807 - 1)
										}
										v270 = int32(60)
										v281 = base.B2i32(int64(86400000000) < v269+base.I64_extend_i32_u((l3*v270+l4)*v270)*int64(1000000))
									}
								}
							}
						}
					}
					if v281 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v389 = m.ExcPending
						if v389 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v392 = m.ExcPending
							if v392 != 0 {
								return int64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
								F_errmsg(m, int32(332099), v15)
								mBase = m.M
								v398 = m.ExcPending
								if v398 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(486977), int32(616), int32(305929))
									mBase = m.M
									v403 = m.ExcPending
									if v403 != 0 {
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
						v284 = v15 + int32(80)
						v287 = base.I64_extend_i32_s(v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167) - int32(2451545))
						v296 = int64(32)
						v297 = int64(20)
						v299 = int64(base.Ui64(v287) >> (uint(v296) % 64))
						v302 = int64(4294967295)
						v303 = int64(500654080)
						v305 = v287 & v302
						v306 = v303 * v305
						v310 = int64(base.Ui64(v306)>>(uint(v296)%64)) + v303*v299
						v317 = v305*v297 + v310&v302
						*(*int64)(unsafe.Add(mBase, uint32(v284)+8)) = v287*int64(0) + v287>>(uint(int64(63))%64)*int64(86400000000) + v297*v299 + int64(base.Ui64(v310)>>(uint(v296)%64)) + int64(base.Ui64(v317)>>(uint(v296)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v284))) = v306&v302 | v317<<(uint(v296)%64)
						v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
						v329 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
						if v328 != v329>>(uint(int64(63))%64) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v409 = m.ExcPending
							if v409 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v412 = m.ExcPending
								if v412 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
									F_errmsg(m, int32(332049), v15+int32(16))
									mBase = m.M
									v423 = m.ExcPending
									if v423 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(486977), int32(628), int32(305929))
										mBase = m.M
										v428 = m.ExcPending
										if v428 != 0 {
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
							v335 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
							if base.F64_lt(base.F64_abs(v335), float64(9.223372036854776e+18)) != 0 {
								v339 = base.I64_trunc_f64_s(v335)
								v341 = v339
							} else {
								v341 = int64(-9223372036854775807 - 1)
							}
							v342 = int32(60)
							v350 = v341 + base.I64_extend_i32_s((l3*v342+l4)*v342)*int64(1000000)
							v353 = v329 + v350
							if base.B2i32(v350 < int64(0)) != base.B2i32(v353 < v329) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v409 = m.ExcPending
								if v409 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v412 = m.ExcPending
									if v412 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
										F_errmsg(m, int32(332049), v15+int32(16))
										mBase = m.M
										v423 = m.ExcPending
										if v423 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(486977), int32(628), int32(305929))
											mBase = m.M
											v428 = m.ExcPending
											if v428 != 0 {
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
								if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v353+int64(211813488000000000)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v432 = m.ExcPending
									if v432 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v435 = m.ExcPending
										if v435 != 0 {
											return int64(0)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
											F_errmsg(m, int32(332049), v15+int32(48))
											mBase = m.M
											v448 = m.ExcPending
											if v448 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(486977), int32(636), int32(305929))
												mBase = m.M
												v453 = m.ExcPending
												if v453 != 0 {
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
									return v353
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v457 = m.ExcPending
					if v457 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v460 = m.ExcPending
						if v460 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
							F_errmsg(m, int32(457620), v15+int32(96))
							mBase = m.M
							v468 = m.ExcPending
							if v468 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(486977), int32(607), int32(305929))
								mBase = m.M
								v473 = m.ExcPending
								if v473 != 0 {
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
			if v203 < int32(5874898) {
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
				v221 = base.B2i32(int32(2) < v202)
				if int32(2) < v202 {
					v222 = int32(4800)
				} else {
					v222 = int32(4799)
				}
				v223 = v222 + v203
				v228 = base.I32_div_s(v223, int32(4))
				v231 = base.I32_div_s(v223, int32(-100))
				v234 = base.I32_div_s(v223, int32(400))
				if int32(2) < v202 {
					v238 = int32(1)
				} else {
					v238 = int32(13)
				}
				v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
				v247 = int32(1)
				if base.Ui32(int32(24)) < base.Ui32(l3) {
					v281 = v247
				} else {
					if base.Ui32(int32(59)) < base.Ui32(l4) {
						v281 = v247
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807)) {
							v281 = v247
						} else {
							v259 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
							if base.F64_lt(v259, float64(0)) != 0 {
								v281 = v247
							} else {
								if base.F64_gt(v259, float64(6e+07)) != 0 {
									v281 = v247
								} else {
									if base.F64_lt(base.F64_abs(v259), float64(9.223372036854776e+18)) != 0 {
										v267 = base.I64_trunc_f64_s(v259)
										v269 = v267
									} else {
										v269 = int64(-9223372036854775807 - 1)
									}
									v270 = int32(60)
									v281 = base.B2i32(int64(86400000000) < v269+base.I64_extend_i32_u((l3*v270+l4)*v270)*int64(1000000))
								}
							}
						}
					}
				}
				if v281 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v389 = m.ExcPending
					if v389 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v392 = m.ExcPending
						if v392 != 0 {
							return int64(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
							F_errmsg(m, int32(332099), v15)
							mBase = m.M
							v398 = m.ExcPending
							if v398 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(486977), int32(616), int32(305929))
								mBase = m.M
								v403 = m.ExcPending
								if v403 != 0 {
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
					v284 = v15 + int32(80)
					v287 = base.I64_extend_i32_s(v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167) - int32(2451545))
					v296 = int64(32)
					v297 = int64(20)
					v299 = int64(base.Ui64(v287) >> (uint(v296) % 64))
					v302 = int64(4294967295)
					v303 = int64(500654080)
					v305 = v287 & v302
					v306 = v303 * v305
					v310 = int64(base.Ui64(v306)>>(uint(v296)%64)) + v303*v299
					v317 = v305*v297 + v310&v302
					*(*int64)(unsafe.Add(mBase, uint32(v284)+8)) = v287*int64(0) + v287>>(uint(int64(63))%64)*int64(86400000000) + v297*v299 + int64(base.Ui64(v310)>>(uint(v296)%64)) + int64(base.Ui64(v317)>>(uint(v296)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v284))) = v306&v302 | v317<<(uint(v296)%64)
					v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
					v329 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
					if v328 != v329>>(uint(int64(63))%64) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v409 = m.ExcPending
						if v409 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v412 = m.ExcPending
							if v412 != 0 {
								return int64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
								F_errmsg(m, int32(332049), v15+int32(16))
								mBase = m.M
								v423 = m.ExcPending
								if v423 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(486977), int32(628), int32(305929))
									mBase = m.M
									v428 = m.ExcPending
									if v428 != 0 {
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
						v335 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
						if base.F64_lt(base.F64_abs(v335), float64(9.223372036854776e+18)) != 0 {
							v339 = base.I64_trunc_f64_s(v335)
							v341 = v339
						} else {
							v341 = int64(-9223372036854775807 - 1)
						}
						v342 = int32(60)
						v350 = v341 + base.I64_extend_i32_s((l3*v342+l4)*v342)*int64(1000000)
						v353 = v329 + v350
						if base.B2i32(v350 < int64(0)) != base.B2i32(v353 < v329) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v409 = m.ExcPending
							if v409 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v412 = m.ExcPending
								if v412 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
									F_errmsg(m, int32(332049), v15+int32(16))
									mBase = m.M
									v423 = m.ExcPending
									if v423 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(486977), int32(628), int32(305929))
										mBase = m.M
										v428 = m.ExcPending
										if v428 != 0 {
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
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v353+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v432 = m.ExcPending
								if v432 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v435 = m.ExcPending
									if v435 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
										F_errmsg(m, int32(332049), v15+int32(48))
										mBase = m.M
										v448 = m.ExcPending
										if v448 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(486977), int32(636), int32(305929))
											mBase = m.M
											v453 = m.ExcPending
											if v453 != 0 {
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
								return v353
							}
						}
					}
				}
			} else {
				if v203 != int32(5874898) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v457 = m.ExcPending
					if v457 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v460 = m.ExcPending
						if v460 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
							F_errmsg(m, int32(457620), v15+int32(96))
							mBase = m.M
							v468 = m.ExcPending
							if v468 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(486977), int32(607), int32(305929))
								mBase = m.M
								v473 = m.ExcPending
								if v473 != 0 {
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
					if int32(6) <= v202 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v457 = m.ExcPending
						if v457 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v460 = m.ExcPending
							if v460 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l0
								F_errmsg(m, int32(457620), v15+int32(96))
								mBase = m.M
								v468 = m.ExcPending
								if v468 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(486977), int32(607), int32(305929))
									mBase = m.M
									v473 = m.ExcPending
									if v473 != 0 {
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
						v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
						v221 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v222 = int32(4800)
						} else {
							v222 = int32(4799)
						}
						v223 = v222 + v203
						v228 = base.I32_div_s(v223, int32(4))
						v231 = base.I32_div_s(v223, int32(-100))
						v234 = base.I32_div_s(v223, int32(400))
						if int32(2) < v202 {
							v238 = int32(1)
						} else {
							v238 = int32(13)
						}
						v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
						v247 = int32(1)
						if base.Ui32(int32(24)) < base.Ui32(l3) {
							v281 = v247
						} else {
							if base.Ui32(int32(59)) < base.Ui32(l4) {
								v281 = v247
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l5)&int64(9223372036854775807)) {
									v281 = v247
								} else {
									v259 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
									if base.F64_lt(v259, float64(0)) != 0 {
										v281 = v247
									} else {
										if base.F64_gt(v259, float64(6e+07)) != 0 {
											v281 = v247
										} else {
											if base.F64_lt(base.F64_abs(v259), float64(9.223372036854776e+18)) != 0 {
												v267 = base.I64_trunc_f64_s(v259)
												v269 = v267
											} else {
												v269 = int64(-9223372036854775807 - 1)
											}
											v270 = int32(60)
											v281 = base.B2i32(int64(86400000000) < v269+base.I64_extend_i32_u((l3*v270+l4)*v270)*int64(1000000))
										}
									}
								}
							}
						}
						if v281 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v389 = m.ExcPending
							if v389 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v392 = m.ExcPending
								if v392 != 0 {
									return int64(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
									F_errmsg(m, int32(332099), v15)
									mBase = m.M
									v398 = m.ExcPending
									if v398 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(486977), int32(616), int32(305929))
										mBase = m.M
										v403 = m.ExcPending
										if v403 != 0 {
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
							v284 = v15 + int32(80)
							v287 = base.I64_extend_i32_s(v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167) - int32(2451545))
							v296 = int64(32)
							v297 = int64(20)
							v299 = int64(base.Ui64(v287) >> (uint(v296) % 64))
							v302 = int64(4294967295)
							v303 = int64(500654080)
							v305 = v287 & v302
							v306 = v303 * v305
							v310 = int64(base.Ui64(v306)>>(uint(v296)%64)) + v303*v299
							v317 = v305*v297 + v310&v302
							*(*int64)(unsafe.Add(mBase, uint32(v284)+8)) = v287*int64(0) + v287>>(uint(int64(63))%64)*int64(86400000000) + v297*v299 + int64(base.Ui64(v310)>>(uint(v296)%64)) + int64(base.Ui64(v317)>>(uint(v296)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v284))) = v306&v302 | v317<<(uint(v296)%64)
							v328 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
							v329 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
							if v328 != v329>>(uint(int64(63))%64) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v409 = m.ExcPending
								if v409 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v412 = m.ExcPending
									if v412 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
										F_errmsg(m, int32(332049), v15+int32(16))
										mBase = m.M
										v423 = m.ExcPending
										if v423 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(486977), int32(628), int32(305929))
											mBase = m.M
											v428 = m.ExcPending
											if v428 != 0 {
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
								v335 = base.F64_nearest(base.F64_mul(l5, float64(1e+06)))
								if base.F64_lt(base.F64_abs(v335), float64(9.223372036854776e+18)) != 0 {
									v339 = base.I64_trunc_f64_s(v335)
									v341 = v339
								} else {
									v341 = int64(-9223372036854775807 - 1)
								}
								v342 = int32(60)
								v350 = v341 + base.I64_extend_i32_s((l3*v342+l4)*v342)*int64(1000000)
								v353 = v329 + v350
								if base.B2i32(v350 < int64(0)) != base.B2i32(v353 < v329) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v412 = m.ExcPending
										if v412 != 0 {
											return int64(0)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
											F_errmsg(m, int32(332049), v15+int32(16))
											mBase = m.M
											v423 = m.ExcPending
											if v423 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(486977), int32(628), int32(305929))
												mBase = m.M
												v428 = m.ExcPending
												if v428 != 0 {
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
									if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v353+int64(211813488000000000)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v432 = m.ExcPending
										if v432 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v435 = m.ExcPending
											if v435 != 0 {
												return int64(0)
											} else {
												*(*float64)(unsafe.Add(mBase, uint32(v15)+72)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
												*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
												F_errmsg(m, int32(332049), v15+int32(48))
												mBase = m.M
												v448 = m.ExcPending
												if v448 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(486977), int32(636), int32(305929))
													mBase = m.M
													v453 = m.ExcPending
													if v453 != 0 {
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
										return v353
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
		v369 = m.ExcPending
		if v369 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v372 = m.ExcPending
			if v372 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = l0
				F_errmsg(m, int32(457576), v15+int32(112))
				mBase = m.M
				v380 = m.ExcPending
				if v380 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(486977), int32(601), int32(305929))
					mBase = m.M
					v385 = m.ExcPending
					if v385 != 0 {
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
							F_errmsg(m, int32(395441), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(489677), int32(1379), int32(350731))
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
						v33 = int32(4800)
					} else {
						v33 = int32(4799)
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
					v54 = base.I32_div_s((v49+v26)*int32(7834), int32(256))
					v60 = v27 + v34*int32(365) + v39 + v42 + v45 + v54 - int32(32167) - int32(2451545)
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
	var v55 int64
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
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
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(395504), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486977), int32(2866), int32(314305))
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
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(395504), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486977), int32(2875), int32(314305))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
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
			v55 = v8 - v6
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v55
			if base.B2i32(int64(0) < v6) != base.B2i32(v55 < v8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(395504), int32(0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(486977), int32(2890), int32(314305))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
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
				v65 = F_DirectFunctionCall1Coll(m, int32(1514), int32(0), v10)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					return v65
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
			v19 = F_DirectFunctionCall2Coll(m, int32(1282), int32(0), v17, v6)
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
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(1510)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(1511)
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
