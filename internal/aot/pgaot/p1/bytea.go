package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaGetByte(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v19 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v48 = v35
		} else {
			v36 = int32(1)
			if v19&v36 != 0 {
				v48 = int32(base.Ui32(v19)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.B2i32(int32(0) <= v16)&base.B2i32(v16 < v48) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
					F_errmsg(m, int32(_a_F_byteaGetByte_0), v9)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaGetByte_1), int32(652), int32(_a_F_byteaGetByte_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
			v72 = int32(1)
			if v19&v72 != 0 {
				v76 = v72
			} else {
				v76 = int32(4)
			}
			v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v12+v76+v16))))
			m.G0 = v9 + int32(16)
			return v79
		}
	}
}
func F_byteaSetByte(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v17 = int32(base.Ui32(v15) >> (uint(int32(2)) % 32))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = int32(0)
		if base.B2i32(v19 <= v18)&base.B2i32(v18 < v17-int32(4)) == v19 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17 - int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
					F_errmsg(m, int32(_a_F_byteaSetByte_0), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaSetByte_1), int32(719), int32(_a_F_byteaSetByte_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v11)+4)) = uint8(v47)
			m.G0 = v8 + int32(16)
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_bytea_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 float64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 float64
	_ = v186
	var v188 float64
	_ = v188
	var v190 float64
	_ = v190
	var v203 float64
	_ = v203
	var v213 float64
	_ = v213
	var v224 float64
	_ = v224
	var v236 float64
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 float64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 float64
	_ = v265
	var v267 float64
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 float64
	_ = v270
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v279 float64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v300 int32
	_ = v300
	var v301 float64
	_ = v301
	var v312 float64
	_ = v312
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
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
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v411 float64
	_ = v411
	var v413 float64
	_ = v413
	var v415 float64
	_ = v415
	var v428 float64
	_ = v428
	var v438 float64
	_ = v438
	var v449 float64
	_ = v449
	var v461 float64
	_ = v461
	var v464 float64
	_ = v464
	var v465 float64
	_ = v465
	var v468 float64
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v483 float64
	_ = v483
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 float64
	_ = v501
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 float64
	_ = v520
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l0 < int32(100) {
		v534 = v3
		m.G0 = v10 + int32(80)
		return v534
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v15 = int32(0)
		v22 = float64(0)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v24 != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			if v24 != int32(1) {
				v34 = v15
				v35 = v15
				v40 = v22
				for {
					v42 = float64(1)
					v43 = v34 + v25
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
					v45 = F_scalbn(m, v42, v44)
					mBase = m.M
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
					v49 = F_scalbn(m, v42, v48)
					mBase = m.M
					v54 = base.F64_add(base.F64_add(v40, base.F64_div(v42, v49)), base.F64_div(v42, v45))
					v55 = int32(2)
					v56 = v34 + v55
					v58 = v35 + v55
					if v58 != v24&int32(-2) {
						v34 = v56
						v35 = v58
						v40 = v54
						continue
					} else {
						break
					}
					break
				}
				if v24&int32(1) == int32(0) {
					v87 = v54
				} else {
					v64 = v56
					v70 = v54
					v72 = float64(1)
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v25))))
					v76 = F_scalbn(m, v72, v75)
					mBase = m.M
					v87 = base.F64_add(v70, base.F64_div(v72, v76))
				}
			} else {
				v64 = v15
				v70 = v22
				v72 = float64(1)
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v25))))
				v76 = F_scalbn(m, v72, v75)
				mBase = m.M
				v87 = base.F64_add(v70, base.F64_div(v72, v76))
			}
			v89 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			v90 = base.F64_div(v89, v87)
			v91 = base.F64_convert_i32_u(v24)
			if base.F64_le(v90, base.F64_mul(v91, float64(2.5))) == int32(0) {
				v203 = v90
				if base.F64_gt(v203, float64(1.4316557653333333e+08)) == int32(0) {
					v224 = v203
				} else {
					v213 = F_log(m, base.F64_add(base.F64_mul(v203, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v224 = base.F64_mul(v213, float64(-4.294967296e+09))
				}
				v236 = v224
			} else {
				v98 = v24 & int32(3)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				v100 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v24) {
					v109 = int32(0)
					v110 = v100
					v111 = v100
					for {
						v118 = v110 + v99
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
						v120 = int32(0)
						v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
						v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
						v134 = v111 + base.B2i32(v119 == v120) + base.B2i32(v123 == v120) + base.B2i32(v127 == v120) + base.B2i32(v131 == v120)
						v135 = int32(4)
						v136 = v110 + v135
						v138 = v109 + v135
						if v138 != v24&int32(-4) {
							v109 = v138
							v110 = v136
							v111 = v134
							continue
						} else {
							break
						}
						break
					}
					if v98 == int32(0) {
						v175 = v134
					} else {
						v144 = v136
						v145 = v134
						v154 = v144
						v155 = v145
						v158 = v100
						for {
							v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v99))))
							v166 = v155 + base.B2i32(v163 == int32(0))
							v167 = int32(1)
							v170 = v158 + v167
							if v170 != v98 {
								v154 = v154 + v167
								v155 = v166
								v158 = v170
								continue
							} else {
								break
							}
							break
						}
						v175 = v166
					}
				} else {
					v144 = v100
					v145 = v100
					v154 = v144
					v155 = v145
					v158 = v100
					for {
						v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v99))))
						v166 = v155 + base.B2i32(v163 == int32(0))
						v167 = int32(1)
						v170 = v158 + v167
						if v170 != v98 {
							v154 = v154 + v167
							v155 = v166
							v158 = v170
							continue
						} else {
							break
						}
						break
					}
					v175 = v166
				}
				if v175 == int32(0) {
					v224 = v90
					v236 = v224
				} else {
					v186 = F_log(m, base.F64_div(v91, base.F64_convert_i32_s(v175)))
					mBase = m.M
					v236 = base.F64_mul(v186, v91)
				}
			}
		} else {
			v188 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			v190 = base.F64_div(v188, float64(0))
			if base.F64_le(v190, base.F64_mul(base.F64_convert_i32_u(v24), float64(2.5))) != 0 {
				v224 = v190
			} else {
				v203 = v190
				if base.F64_gt(v203, float64(1.4316557653333333e+08)) == int32(0) {
					v224 = v203
				} else {
					v213 = F_log(m, base.F64_add(base.F64_mul(v203, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v224 = base.F64_mul(v213, float64(-4.294967296e+09))
				}
			}
			v236 = v224
		}
		v239 = v14 + int32(24)
		v240 = int32(0)
		v247 = float64(0)
		v249 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
		if v249 != 0 {
			v250 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
			if v249 != int32(1) {
				v259 = v240
				v260 = v240
				v265 = v247
				for {
					v267 = float64(1)
					v268 = v259 + v250
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
					v270 = F_scalbn(m, v267, v269)
					mBase = m.M
					v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
					v274 = F_scalbn(m, v267, v273)
					mBase = m.M
					v279 = base.F64_add(base.F64_add(v265, base.F64_div(v267, v274)), base.F64_div(v267, v270))
					v280 = int32(2)
					v281 = v259 + v280
					v283 = v260 + v280
					if v283 != v249&int32(-2) {
						v259 = v281
						v260 = v283
						v265 = v279
						continue
					} else {
						break
					}
					break
				}
				if v249&int32(1) == int32(0) {
					v312 = v279
				} else {
					v289 = v281
					v295 = v279
					v297 = float64(1)
					v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v250))))
					v301 = F_scalbn(m, v297, v300)
					mBase = m.M
					v312 = base.F64_add(v295, base.F64_div(v297, v301))
				}
			} else {
				v289 = v240
				v295 = v247
				v297 = float64(1)
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v250))))
				v301 = F_scalbn(m, v297, v300)
				mBase = m.M
				v312 = base.F64_add(v295, base.F64_div(v297, v301))
			}
			v314 = *(*float64)(unsafe.Add(mBase, uint32(v239)+8))
			v315 = base.F64_div(v314, v312)
			v316 = base.F64_convert_i32_u(v249)
			if base.F64_le(v315, base.F64_mul(v316, float64(2.5))) == int32(0) {
				v428 = v315
				if base.F64_gt(v428, float64(1.4316557653333333e+08)) == int32(0) {
					v449 = v428
				} else {
					v438 = F_log(m, base.F64_add(base.F64_mul(v428, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v449 = base.F64_mul(v438, float64(-4.294967296e+09))
				}
				v461 = v449
			} else {
				v323 = v249 & int32(3)
				v324 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
				v325 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v249) {
					v334 = int32(0)
					v335 = v325
					v336 = v325
					for {
						v343 = v335 + v324
						v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
						v345 = int32(0)
						v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
						v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+2)))
						v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+3)))
						v359 = v336 + base.B2i32(v344 == v345) + base.B2i32(v348 == v345) + base.B2i32(v352 == v345) + base.B2i32(v356 == v345)
						v360 = int32(4)
						v361 = v335 + v360
						v363 = v334 + v360
						if v363 != v249&int32(-4) {
							v334 = v363
							v335 = v361
							v336 = v359
							continue
						} else {
							break
						}
						break
					}
					if v323 == int32(0) {
						v400 = v359
					} else {
						v369 = v361
						v370 = v359
						v379 = v369
						v380 = v370
						v383 = v325
						for {
							v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+v324))))
							v391 = v380 + base.B2i32(v388 == int32(0))
							v392 = int32(1)
							v395 = v383 + v392
							if v395 != v323 {
								v379 = v379 + v392
								v380 = v391
								v383 = v395
								continue
							} else {
								break
							}
							break
						}
						v400 = v391
					}
				} else {
					v369 = v325
					v370 = v325
					v379 = v369
					v380 = v370
					v383 = v325
					for {
						v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+v324))))
						v391 = v380 + base.B2i32(v388 == int32(0))
						v392 = int32(1)
						v395 = v383 + v392
						if v395 != v323 {
							v379 = v379 + v392
							v380 = v391
							v383 = v395
							continue
						} else {
							break
						}
						break
					}
					v400 = v391
				}
				if v400 == int32(0) {
					v449 = v315
					v461 = v449
				} else {
					v411 = F_log(m, base.F64_div(v316, base.F64_convert_i32_s(v400)))
					mBase = m.M
					v461 = base.F64_mul(v411, v316)
				}
			}
		} else {
			v413 = *(*float64)(unsafe.Add(mBase, uint32(v239)+8))
			v415 = base.F64_div(v413, float64(0))
			if base.F64_le(v415, base.F64_mul(base.F64_convert_i32_u(v249), float64(2.5))) != 0 {
				v449 = v415
			} else {
				v428 = v415
				if base.F64_gt(v428, float64(1.4316557653333333e+08)) == int32(0) {
					v449 = v428
				} else {
					v438 = F_log(m, base.F64_add(base.F64_mul(v428, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v449 = base.F64_mul(v438, float64(-4.294967296e+09))
				}
			}
			v461 = v449
		}
		if base.F64_lt(v461, float64(1)) != 0 {
			v464 = float64(1)
		} else {
			v464 = v461
		}
		v465 = float64(1)
		if base.F64_lt(v236, v465) != 0 {
			v468 = v465
		} else {
			v468 = v236
		}
		v470 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bytea_abbrev_abort[0])))
		if v470 != int32(1) {
			v501 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
			if base.F64_lt(base.F64_mul(v464, v501), v468) != 0 {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_bytea_abbrev_abort_0)) {
					v534 = v3
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = base.F64_mul(v501, float64(0.65))
					v534 = v3
				}
				m.G0 = v10 + int32(80)
				return v534
			} else {
				v509 = int32(1)
				v511 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bytea_abbrev_abort[0])))
				if v511 != v509 {
					v534 = v509
					m.G0 = v10 + int32(80)
					return v534
				} else {
					v516 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v517 = m.ExcPending
					if v517 != 0 {
						return int32(0)
					} else {
						if v516 == int32(0) {
							v534 = v509
							m.G0 = v10 + int32(80)
							return v534
						} else {
							v520 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
							*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v520
							*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v464
							*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v468
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
							F_errmsg_internal(m, int32(_a_F_bytea_abbrev_abort_1), v10)
							mBase = m.M
							v527 = m.ExcPending
							if v527 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_bytea_abbrev_abort_2), int32(1208), int32(_a_F_bytea_abbrev_abort_3))
								mBase = m.M
								v532 = m.ExcPending
								if v532 != 0 {
									return int32(0)
								} else {
									v534 = v509
									m.G0 = v10 + int32(80)
									return v534
								}
							}
						}
					}
				}
			}
		} else {
			v475 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v478 = m.ExcPending
			if v478 != 0 {
				return int32(0)
			} else {
				if v475 == int32(0) {
					v501 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
					if base.F64_lt(base.F64_mul(v464, v501), v468) != 0 {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_bytea_abbrev_abort_0)) {
							v534 = v3
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = base.F64_mul(v501, float64(0.65))
							v534 = v3
						}
						m.G0 = v10 + int32(80)
						return v534
					} else {
						v509 = int32(1)
						v511 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bytea_abbrev_abort[0])))
						if v511 != v509 {
							v534 = v509
							m.G0 = v10 + int32(80)
							return v534
						} else {
							v516 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v517 = m.ExcPending
							if v517 != 0 {
								return int32(0)
							} else {
								if v516 == int32(0) {
									v534 = v509
									m.G0 = v10 + int32(80)
									return v534
								} else {
									v520 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
									*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v520
									*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v464
									*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v468
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
									F_errmsg_internal(m, int32(_a_F_bytea_abbrev_abort_1), v10)
									mBase = m.M
									v527 = m.ExcPending
									if v527 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_bytea_abbrev_abort_2), int32(1208), int32(_a_F_bytea_abbrev_abort_3))
										mBase = m.M
										v532 = m.ExcPending
										if v532 != 0 {
											return int32(0)
										} else {
											v534 = v509
											m.G0 = v10 + int32(80)
											return v534
										}
									}
								}
							}
						}
					}
				} else {
					v483 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
					*(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v483
					*(*float64)(unsafe.Add(mBase, uint32(v10)+48)) = v464
					*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = v468
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v10)+56)) = base.F64_div(v468, base.F64_convert_i32_u(l0))
					F_errmsg_internal(m, int32(_a_F_bytea_abbrev_abort_4), v10+int32(32))
					mBase = m.M
					v495 = m.ExcPending
					if v495 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bytea_abbrev_abort_2), int32(1184), int32(_a_F_bytea_abbrev_abort_3))
						mBase = m.M
						v500 = m.ExcPending
						if v500 != 0 {
							return int32(0)
						} else {
							v501 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
							if base.F64_lt(base.F64_mul(v464, v501), v468) != 0 {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_bytea_abbrev_abort_0)) {
									v534 = v3
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = base.F64_mul(v501, float64(0.65))
									v534 = v3
								}
								m.G0 = v10 + int32(80)
								return v534
							} else {
								v509 = int32(1)
								v511 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bytea_abbrev_abort[0])))
								if v511 != v509 {
									v534 = v509
									m.G0 = v10 + int32(80)
									return v534
								} else {
									v516 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v517 = m.ExcPending
									if v517 != 0 {
										return int32(0)
									} else {
										if v516 == int32(0) {
											v534 = v509
											m.G0 = v10 + int32(80)
											return v534
										} else {
											v520 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
											*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v520
											*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v464
											*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v468
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
											F_errmsg_internal(m, int32(_a_F_bytea_abbrev_abort_1), v10)
											mBase = m.M
											v527 = m.ExcPending
											if v527 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_bytea_abbrev_abort_2), int32(1208), int32(_a_F_bytea_abbrev_abort_3))
												mBase = m.M
												v532 = m.ExcPending
												if v532 != 0 {
													return int32(0)
												} else {
													v534 = v509
													m.G0 = v10 + int32(80)
													return v534
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
func F_bytea_catenate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
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
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(1) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v15 == int32(18) {
			v18 = int32(16)
		} else {
			v18 = int32(0)
		}
		if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v25 = int32(4)
		} else {
			v25 = v18
		}
		v38 = v25
	} else {
		v26 = int32(1)
		if v9&v26 != 0 {
			v38 = int32(base.Ui32(v9)>>(uint(v26)%32)) - v26
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v39 = int32(0)
	if v39 < v38 {
		v42 = v38
	} else {
		v42 = v39
	}
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43 == int32(1) {
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v49 == int32(18) {
			v52 = int32(16)
		} else {
			v52 = int32(0)
		}
		if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v59 = int32(4)
		} else {
			v59 = v52
		}
		v72 = v59
	} else {
		v60 = int32(1)
		if v43&v60 != 0 {
			v72 = int32(base.Ui32(v43)>>(uint(v60)%32)) - v60
		} else {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v72 = int32(base.Ui32(v66)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v73 = int32(0)
	if v73 < v72 {
		v76 = v72
	} else {
		v76 = v73
	}
	v79 = v42 + v76 + int32(4)
	v80 = F_palloc(m, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v80))) = v79 << (uint(int32(2)) % 32)
		v88 = v80 + int32(4)
		v89 = int32(0)
		if base.B2i32(v42 == v89)|base.B2i32(v38 <= v89) == v89 {
			v96 = int32(1)
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v98&v96 != 0 {
				v101 = v96
			} else {
				v101 = int32(4)
			}
			base.MemoryCopy(m, v88, l0+v101, v42)
		} else {
		}
		v104 = int32(0)
		if base.B2i32(v76 == v104)|base.B2i32(v72 <= v104) == v104 {
			v112 = int32(1)
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v114&v112 != 0 {
				v117 = v112
			} else {
				v117 = int32(4)
			}
			base.MemoryCopy(m, v42+v88, l1+v117, v76)
		} else {
		}
		return v80
	}
}
func F_bytea_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v10 == int32(1) {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if base.Ui32((v13-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v63 = F_errsave_start(m, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					if v63 != 0 {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v62, int32(_a_F_bytea_int2_1), int32(1259), int32(_a_F_bytea_int2_2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									return int64(0)
								}
							}
						}
					} else {
						return int64(0)
					}
				}
			} else {
				if v13 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				v37 = v24
				if base.Ui32(int32(2)) < base.Ui32(v37) {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v63 = F_errsave_start(m, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						if v63 != 0 {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v62, int32(_a_F_bytea_int2_1), int32(1259), int32(_a_F_bytea_int2_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										return int64(0)
									}
								}
							}
						} else {
							return int64(0)
						}
					}
				} else {
					if v37 == int32(0) {
						return int64(0)
					} else {
						v44 = int32(1)
						if v10&v44 != 0 {
							v48 = v44
						} else {
							v48 = int32(4)
						}
						v49 = v6 + v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						if v37 != int32(1) {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
							v57 = v53 | v50<<(uint(int32(8))%32)
						} else {
							v57 = v50
						}
						return base.I64_extend16_s(base.I64_extend_i32_u(v57))
					}
				}
			}
		} else {
			v25 = int32(1)
			if v10&v25 != 0 {
				v37 = int32(base.Ui32(v10)>>(uint(v25)%32)) - v25
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v37 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(2)) < base.Ui32(v37) {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v63 = F_errsave_start(m, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					if v63 != 0 {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v62, int32(_a_F_bytea_int2_1), int32(1259), int32(_a_F_bytea_int2_2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									return int64(0)
								}
							}
						}
					} else {
						return int64(0)
					}
				}
			} else {
				if v37 == int32(0) {
					return int64(0)
				} else {
					v44 = int32(1)
					if v10&v44 != 0 {
						v48 = v44
					} else {
						v48 = int32(4)
					}
					v49 = v6 + v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					if v37 != int32(1) {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
						v57 = v53 | v50<<(uint(int32(8))%32)
					} else {
						v57 = v50
					}
					return base.I64_extend16_s(base.I64_extend_i32_u(v57))
				}
			}
		}
	}
}
func F_bytea_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	var v135 int32
	_ = v135
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v13 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v47 = int32(4)
				v50 = v47 & int32(3)
				v51 = int32(1)
				if v13&v51 != 0 {
					v55 = v51
				} else {
					v55 = int32(4)
				}
				v56 = v9 + v55
				v57 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v47) {
					v63 = v57
					v64 = v57
					for {
						v70 = int32(4)
						v71 = v63 + v70
						v73 = v64 + v70
						if v73 != v47&int32(4) {
							v63 = v71
							v64 = v73
							continue
						} else {
							break
						}
						break
					}
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
					v77 = int32(16711935)
					v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
					if v50 == int32(0) {
						v113 = v85
					} else {
						v88 = v71
						v89 = v85
						v95 = v88
						v96 = v89
						v100 = int32(0)
						for {
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
							v106 = v103 | v96<<(uint(int32(8))%32)
							v107 = int32(1)
							v110 = v100 + v107
							if v110 != v50 {
								v95 = v95 + v107
								v96 = v106
								v100 = v110
								continue
							} else {
								break
							}
							break
						}
						v113 = v106
					}
				} else {
					v88 = v57
					v89 = v57
					v95 = v88
					v96 = v89
					v100 = int32(0)
					for {
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
						v106 = v103 | v96<<(uint(int32(8))%32)
						v107 = int32(1)
						v110 = v100 + v107
						if v110 != v50 {
							v95 = v95 + v107
							v96 = v106
							v100 = v110
							continue
						} else {
							break
						}
						break
					}
					v113 = v106
				}
				return base.I64_extend_i32_s(v113)
			} else {
				if v17 == int32(18) {
					v28 = int32(16)
				} else {
					v28 = int32(0)
				}
				v42 = v28
				if base.Ui32(int32(4)) < base.Ui32(v42) {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v122 = F_errsave_start(m, v121)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						if v122 != 0 {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_bytea_int4_0), int32(0))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v121, int32(_a_F_bytea_int4_1), int32(1284), int32(_a_F_bytea_int4_2))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int64(0)
									} else {
										return int64(0)
									}
								}
							}
						} else {
							return int64(0)
						}
					}
				} else {
					if v42 != 0 {
						v47 = v42
						v50 = v47 & int32(3)
						v51 = int32(1)
						if v13&v51 != 0 {
							v55 = v51
						} else {
							v55 = int32(4)
						}
						v56 = v9 + v55
						v57 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v47) {
							v63 = v57
							v64 = v57
							for {
								v70 = int32(4)
								v71 = v63 + v70
								v73 = v64 + v70
								if v73 != v47&int32(4) {
									v63 = v71
									v64 = v73
									continue
								} else {
									break
								}
								break
							}
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
							v77 = int32(16711935)
							v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
							if v50 == int32(0) {
								v113 = v85
							} else {
								v88 = v71
								v89 = v85
								v95 = v88
								v96 = v89
								v100 = int32(0)
								for {
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
									v106 = v103 | v96<<(uint(int32(8))%32)
									v107 = int32(1)
									v110 = v100 + v107
									if v110 != v50 {
										v95 = v95 + v107
										v96 = v106
										v100 = v110
										continue
									} else {
										break
									}
									break
								}
								v113 = v106
							}
						} else {
							v88 = v57
							v89 = v57
							v95 = v88
							v96 = v89
							v100 = int32(0)
							for {
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
								v106 = v103 | v96<<(uint(int32(8))%32)
								v107 = int32(1)
								v110 = v100 + v107
								if v110 != v50 {
									v95 = v95 + v107
									v96 = v106
									v100 = v110
									continue
								} else {
									break
								}
								break
							}
							v113 = v106
						}
						return base.I64_extend_i32_s(v113)
					} else {
						return int64(0)
					}
				}
			}
		} else {
			v29 = int32(1)
			if v13&v29 != 0 {
				v42 = int32(base.Ui32(v13)>>(uint(v29)%32)) - v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(4)) < base.Ui32(v42) {
				v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v122 = F_errsave_start(m, v121)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int64(0)
				} else {
					if v122 != 0 {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int4_0), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v121, int32(_a_F_bytea_int4_1), int32(1284), int32(_a_F_bytea_int4_2))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int64(0)
								} else {
									return int64(0)
								}
							}
						}
					} else {
						return int64(0)
					}
				}
			} else {
				if v42 != 0 {
					v47 = v42
					v50 = v47 & int32(3)
					v51 = int32(1)
					if v13&v51 != 0 {
						v55 = v51
					} else {
						v55 = int32(4)
					}
					v56 = v9 + v55
					v57 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v47) {
						v63 = v57
						v64 = v57
						for {
							v70 = int32(4)
							v71 = v63 + v70
							v73 = v64 + v70
							if v73 != v47&int32(4) {
								v63 = v71
								v64 = v73
								continue
							} else {
								break
							}
							break
						}
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
						v77 = int32(16711935)
						v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
						if v50 == int32(0) {
							v113 = v85
						} else {
							v88 = v71
							v89 = v85
							v95 = v88
							v96 = v89
							v100 = int32(0)
							for {
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
								v106 = v103 | v96<<(uint(int32(8))%32)
								v107 = int32(1)
								v110 = v100 + v107
								if v110 != v50 {
									v95 = v95 + v107
									v96 = v106
									v100 = v110
									continue
								} else {
									break
								}
								break
							}
							v113 = v106
						}
					} else {
						v88 = v57
						v89 = v57
						v95 = v88
						v96 = v89
						v100 = int32(0)
						for {
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
							v106 = v103 | v96<<(uint(int32(8))%32)
							v107 = int32(1)
							v110 = v100 + v107
							if v110 != v50 {
								v95 = v95 + v107
								v96 = v106
								v100 = v110
								continue
							} else {
								break
							}
							break
						}
						v113 = v106
					}
					return base.I64_extend_i32_s(v113)
				} else {
					return int64(0)
				}
			}
		}
	}
}
func F_bytea_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v146 int64
	_ = v146
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v14 == int32(1) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v48 = int32(4)
				v51 = v48 & int32(3)
				v52 = int32(1)
				if v14&v52 != 0 {
					v56 = v52
				} else {
					v56 = int32(4)
				}
				v57 = v10 + v56
				v58 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v48) {
					v64 = v58
					v69 = int32(0)
					v71 = v8
					for {
						v72 = int64(16)
						v74 = v64 + v57
						v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
						v76 = int64(8)
						v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
						v83 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
						v87 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+3)))
						v88 = (v71<<(uint(v72)%64)|v75<<(uint(v76)%64)|v79)<<(uint(v72)%64) | v83<<(uint(v76)%64) | v87
						v89 = int32(4)
						v90 = v64 + v89
						v92 = v69 + v89
						if v92 != v48&int32(12) {
							v64 = v90
							v69 = v92
							v71 = v88
							continue
						} else {
							break
						}
						break
					}
					if v51 == int32(0) {
						v146 = v88
					} else {
						v96 = v90
						v103 = v88
						v104 = v96
						v107 = v58
						v111 = v103
						for {
							v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
							v116 = v113 | v111<<(uint(int64(8))%64)
							v117 = int32(1)
							v120 = v107 + v117
							if v120 != v51 {
								v104 = v104 + v117
								v107 = v120
								v111 = v116
								continue
							} else {
								break
							}
							break
						}
						v146 = v116
					}
				} else {
					v96 = v58
					v103 = v8
					v104 = v96
					v107 = v58
					v111 = v103
					for {
						v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
						v116 = v113 | v111<<(uint(int64(8))%64)
						v117 = int32(1)
						v120 = v107 + v117
						if v120 != v51 {
							v104 = v104 + v117
							v107 = v120
							v111 = v116
							continue
						} else {
							break
						}
						break
					}
					v146 = v116
				}
				return v146
			} else {
				if v18 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				v43 = v29
				if base.Ui32(int32(8)) < base.Ui32(v43) {
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v123 = F_errsave_start(m, v122)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int64(0)
					} else {
						if v123 == int32(0) {
							v146 = v8
							return v146
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_bytea_int8_0), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v122, int32(_a_F_bytea_int8_1), int32(1309), int32(_a_F_bytea_int8_2))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int64(0)
									} else {
										v146 = v8
										return v146
									}
								}
							}
						}
					}
				} else {
					if v43 != 0 {
						v48 = v43
						v51 = v48 & int32(3)
						v52 = int32(1)
						if v14&v52 != 0 {
							v56 = v52
						} else {
							v56 = int32(4)
						}
						v57 = v10 + v56
						v58 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v48) {
							v64 = v58
							v69 = int32(0)
							v71 = v8
							for {
								v72 = int64(16)
								v74 = v64 + v57
								v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								v76 = int64(8)
								v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
								v83 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
								v87 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+3)))
								v88 = (v71<<(uint(v72)%64)|v75<<(uint(v76)%64)|v79)<<(uint(v72)%64) | v83<<(uint(v76)%64) | v87
								v89 = int32(4)
								v90 = v64 + v89
								v92 = v69 + v89
								if v92 != v48&int32(12) {
									v64 = v90
									v69 = v92
									v71 = v88
									continue
								} else {
									break
								}
								break
							}
							if v51 == int32(0) {
								v146 = v88
							} else {
								v96 = v90
								v103 = v88
								v104 = v96
								v107 = v58
								v111 = v103
								for {
									v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
									v116 = v113 | v111<<(uint(int64(8))%64)
									v117 = int32(1)
									v120 = v107 + v117
									if v120 != v51 {
										v104 = v104 + v117
										v107 = v120
										v111 = v116
										continue
									} else {
										break
									}
									break
								}
								v146 = v116
							}
						} else {
							v96 = v58
							v103 = v8
							v104 = v96
							v107 = v58
							v111 = v103
							for {
								v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
								v116 = v113 | v111<<(uint(int64(8))%64)
								v117 = int32(1)
								v120 = v107 + v117
								if v120 != v51 {
									v104 = v104 + v117
									v107 = v120
									v111 = v116
									continue
								} else {
									break
								}
								break
							}
							v146 = v116
						}
						return v146
					} else {
						return int64(0)
					}
				}
			}
		} else {
			v30 = int32(1)
			if v14&v30 != 0 {
				v43 = int32(base.Ui32(v14)>>(uint(v30)%32)) - v30
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v43 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(8)) < base.Ui32(v43) {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v123 = F_errsave_start(m, v122)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int64(0)
				} else {
					if v123 == int32(0) {
						v146 = v8
						return v146
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int8_0), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v122, int32(_a_F_bytea_int8_1), int32(1309), int32(_a_F_bytea_int8_2))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int64(0)
								} else {
									v146 = v8
									return v146
								}
							}
						}
					}
				}
			} else {
				if v43 != 0 {
					v48 = v43
					v51 = v48 & int32(3)
					v52 = int32(1)
					if v14&v52 != 0 {
						v56 = v52
					} else {
						v56 = int32(4)
					}
					v57 = v10 + v56
					v58 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v48) {
						v64 = v58
						v69 = int32(0)
						v71 = v8
						for {
							v72 = int64(16)
							v74 = v64 + v57
							v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
							v76 = int64(8)
							v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
							v83 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
							v87 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+3)))
							v88 = (v71<<(uint(v72)%64)|v75<<(uint(v76)%64)|v79)<<(uint(v72)%64) | v83<<(uint(v76)%64) | v87
							v89 = int32(4)
							v90 = v64 + v89
							v92 = v69 + v89
							if v92 != v48&int32(12) {
								v64 = v90
								v69 = v92
								v71 = v88
								continue
							} else {
								break
							}
							break
						}
						if v51 == int32(0) {
							v146 = v88
						} else {
							v96 = v90
							v103 = v88
							v104 = v96
							v107 = v58
							v111 = v103
							for {
								v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
								v116 = v113 | v111<<(uint(int64(8))%64)
								v117 = int32(1)
								v120 = v107 + v117
								if v120 != v51 {
									v104 = v104 + v117
									v107 = v120
									v111 = v116
									continue
								} else {
									break
								}
								break
							}
							v146 = v116
						}
					} else {
						v96 = v58
						v103 = v8
						v104 = v96
						v107 = v58
						v111 = v103
						for {
							v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104+v57))))
							v116 = v113 | v111<<(uint(int64(8))%64)
							v117 = int32(1)
							v120 = v107 + v117
							if v120 != v51 {
								v104 = v104 + v117
								v107 = v120
								v111 = v116
								continue
							} else {
								break
							}
							break
						}
						v146 = v116
					}
					return v146
				} else {
					return int64(0)
				}
			}
		}
	}
}
func F_bytea_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v15 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v45 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v21 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v24 = int32(16)
	goto L10
L9:
	;
	v24 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(4)
	goto L13
L12:
	;
	v31 = v24
	goto L13
L13:
	;
	v44 = v31
	goto L4
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v74 < v44 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v51 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v62 = int32(1)
	if v45&v62 != 0 {
		v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v54 = int32(16)
	goto L21
L20:
	;
	v54 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(4)
	goto L24
L23:
	;
	v61 = v54
	goto L24
L24:
	;
	v74 = v61
	goto L15
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v76 = v8
	goto L28
L27:
	;
	v76 = v13
	goto L28
L28:
	;
	v77 = int32(1)
	if v15&v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v77
	goto L31
L30:
	;
	v81 = int32(4)
	goto L31
L31:
	;
	v82 = v8 + v81
	v83 = int32(1)
	if v45&v83 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = v83
	goto L34
L33:
	;
	v87 = int32(4)
	goto L34
L34:
	;
	v88 = v13 + v87
	if v44 < v74 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v44
	goto L37
L36:
	;
	v90 = v74
	goto L37
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v152 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v152 = int32(0)
	goto L38
L40:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L50
L41:
	;
	if (v82|v88)&int32(3) != 0 {
		v121 = v82
		v122 = v88
		v123 = v90
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v114 = v82
	v115 = v88
	v116 = v90
	goto L43
L43:
	;
	if v116 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v98 = v82
	v99 = v88
	v100 = v90
	goto L45
L45:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L43
L47:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L40
L50:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v152 = v131 - v132
	goto L38
L52:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L39
L56:
	;
	v153 = v13
	goto L58
L57:
	;
	v153 = v76
	goto L58
L58:
	;
	if int32(0) < v152 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = v8
	goto L61
L60:
	;
	v156 = v153
	goto L61
L61:
	;
	return base.I64_extend_i32_u(v156)
}
func F_bytea_substr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = F_bytea_substring(m, v2, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_bytea_substring(m *base.Module, l0 int64, l1 int32, l2 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	if int32(0) <= l2 {
		v8 = int32(1)
		if l1 <= v8 {
			v11 = v8
		} else {
			v11 = l1
		}
		v16 = l1 + l2
		if base.B2i32(l2 < int32(0))^base.B2i32(v16 < l1) != 0 {
			v34 = int32(-1)
			v35 = F_detoast_attr_slice(m, base.I32_wrap_i64(l0), v11-int32(1), v34)
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		} else {
			if v16 <= int32(0) {
				v25 = F_DirectFunctionCall1Coll(m, int32(624), int32(0), int64(828758))
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v25))
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v30
					}
				}
			} else {
				v34 = v16 - v11
				v35 = F_detoast_attr_slice(m, base.I32_wrap_i64(l0), v11-int32(1), v34)
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_bytea_substring_0), int32(0))
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bytea_substring_1), int32(125), int32(_a_F_bytea_substring_2))
					v53 = m.ExcPending
					if v53 != 0 {
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
