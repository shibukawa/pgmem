package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MJFillOuter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 float64
	_ = v68
	var v74 int32
	_ = v74
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_MemoryContextReset(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v20
		if v11 != 0 {
			v22 = int32(4549024)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) int32)(m, v11, v12, v9+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
				if v30 == int32(0) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v65 == int32(0) {
						v74 = v2
					} else {
						v68 = *(*float64)(unsafe.Add(mBase, uint32(v65)+248))
						*(*float64)(unsafe.Add(mBase, uint32(v65)+248)) = base.F64_add(v68, float64(1))
						v74 = v2
					}
					m.G0 = v9 + int32(16)
					return v74
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
					m.T0[v42].(func(*base.Module, int32))(m, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(4549024)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v48
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
						v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v38+int32(4), v39, int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
							v60 = v58 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v60)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v63)
							v74 = v40
							m.G0 = v9 + int32(16)
							return v74
						}
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
			m.T0[v42].(func(*base.Module, int32))(m, v40)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = int32(4549024)
				v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v48
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
				v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v38+int32(4), v39, int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
					v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
					v60 = v58 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v60)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
					*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v63)
					v74 = v40
					m.G0 = v9 + int32(16)
					return v74
				}
			}
		}
	}
}
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v5 == int32(0) {
		v8 = int32(4549024)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
		v14 = F_palloc0(m, int32(72))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(382)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v22
			v28 = F_AllocSetContextCreateInternal(m, v22, int32(67362), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				v34 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l0
				v37 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v37)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v34
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v37)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v33
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				v49 = F_lcons(m, v14, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v49
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
					*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v14
					v55 = v14
					return v55
				}
			}
		}
	} else {
		v55 = v5
		return v55
	}
}
func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v102 int64
	_ = v102
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v141 int64
	_ = v141
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v189 int64
	_ = v189
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v284 int64
	_ = v284
	var v295 int64
	_ = v295
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v312 int64
	_ = v312
	var v329 int64
	_ = v329
	var v338 int64
	_ = v338
	var v341 int64
	_ = v341
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v396 int64
	_ = v396
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v410 int32
	_ = v410
	var v421 int64
	_ = v421
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v435 int32
	_ = v435
	var v450 int64
	_ = v450
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v473 int64
	_ = v473
	var v477 int64
	_ = v477
	var v478 int64
	_ = v478
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v492 int64
	_ = v492
	var v493 int64
	_ = v493
	var v495 int64
	_ = v495
	var v496 int64
	_ = v496
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v508 int64
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v517 int32
	_ = v517
	var v521 int64
	_ = v521
	var v533 int64
	_ = v533
	var v537 int64
	_ = v537
	var v545 int64
	_ = v545
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = int64(281474976710655)
	v31 = l4 & v30
	v34 = (l2 ^ l4) & int64(-9223372036854775807-1)
	v36 = l2 & v30
	v38 = int64(base.Ui64(v36) >> (uint(int64(32)) % 64))
	v39 = int64(48)
	v42 = int32(32767)
	v43 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v39)%64))) & v42
	v48 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v39)%64))) & v42
	if base.Ui32(int32(-32766)) <= base.Ui32(v48-v42) {
		if base.Ui32(int32(-32767)) < base.Ui32(v43-int32(32767)) {
			v207 = l1
			v209 = l3
			v210 = v6
			v213 = v31
			v214 = v36
			v215 = v38
			v216 = int64(15)
			v217 = v209 << (uint(v216) % 64)
			v219 = v217 & int64(4294934528)
			v220 = int64(32)
			v221 = int64(base.Ui64(v207) >> (uint(v220) % 64))
			v222 = v219 * v221
			v224 = int64(base.Ui64(v217) >> (uint(v220) % 64))
			v225 = int64(4294967295)
			v226 = v207 & v225
			v228 = v222 + v224*v226
			v230 = v228 << (uint(v220) % 64)
			v232 = v230 + v226*v219
			v236 = v214 & v225
			v237 = v219 * v236
			v239 = v237 + v221*v224
			v241 = v213 << (uint(v216) % 64)
			v246 = (v241 | int64(base.Ui64(v209)>>(uint(int64(49))%64))) & v225
			v248 = v239 + v246*v226
			v256 = v248 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v222)))<<(uint(v220)%64) | int64(base.Ui64(v228)>>(uint(v220)%64)))
			v258 = v215 | int64(65536)
			v259 = v219 * v258
			v261 = v259 + v224*v236
			v265 = int64(base.Ui64(v241)>>(uint(v220)%64)) | int64(2147483648)
			v267 = v261 + v265*v226
			v269 = v267 + v246*v221
			v272 = v256 + v269<<(uint(v220)%64)
			v273 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v232) < base.Ui64(v230))) + v272
			v277 = v43 + v48 + v210 - int32(16383)
			v278 = v265 * v221
			v280 = v278 + v224*v258
			v284 = v280 + v246*v236
			v295 = v284 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v237))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v248) < base.Ui64(v239))))
			v299 = v246 * v258
			v301 = v299 + v265*v236
			v312 = v295 + v301<<(uint(v220)%64)
			v329 = v312 + ((base.I64_extend_i32_u(base.B2i32(base.Ui64(v269) < base.Ui64(v267)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v261) < base.Ui64(v259)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v261)))))<<(uint(v220)%64) | int64(base.Ui64(v269)>>(uint(v220)%64)))
			v338 = v329 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v256) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v256))))
			v341 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v284) < base.Ui64(v280))) + v265*v258 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v284))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v301) < base.Ui64(v299)))<<(uint(v220)%64) | int64(base.Ui64(v301)>>(uint(v220)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v295))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v329) < base.Ui64(v312))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v338) < base.Ui64(v329)))
			if v341&int64(281474976710656) != int64(0) {
				v365 = v273
				v366 = v338
				v368 = v341
				v369 = v277 + int32(1)
				v370 = v232
			} else {
				v348 = int64(63)
				v350 = int64(1)
				v365 = int64(base.Ui64(v232)>>(uint(v348)%64)) | v273<<(uint(v350)%64)
				v366 = v338<<(uint(v350)%64) | int64(base.Ui64(v273)>>(uint(v348)%64))
				v368 = v341<<(uint(v350)%64) | int64(base.Ui64(v338)>>(uint(v348)%64))
				v369 = v277
				v370 = v232 << (uint(v350) % 64)
			}
			if int32(32767) <= v369 {
				v537 = int64(0)
				v545 = v34 | int64(9223090561878065152)
			} else {
				if v369 <= int32(0) {
					v379 = int32(1) - v369
					if base.Ui32(v379) <= base.Ui32(int32(127)) {
						v383 = v28 + int32(48)
						v385 = v369 + int32(127)
						if v385&int32(64) != 0 {
							v404 = int64(0)
							v405 = v370 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
						} else {
							if v385 == int32(0) {
								v404 = v370
								v405 = v365
							} else {
								v396 = base.I64_extend_i32_u(v385)
								v404 = v370 << (uint(v396) % 64)
								v405 = v365<<(uint(v396)%64) | int64(base.Ui64(v370)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v383))) = v404
						*(*int64)(unsafe.Add(mBase, uint32(v383)+8)) = v405
						v410 = v28 + int32(32)
						if v385&int32(64) != 0 {
							v429 = int64(0)
							v430 = v366 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
						} else {
							if v385 == int32(0) {
								v429 = v366
								v430 = v368
							} else {
								v421 = base.I64_extend_i32_u(v385)
								v429 = v366 << (uint(v421) % 64)
								v430 = v368<<(uint(v421)%64) | int64(base.Ui64(v366)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v410))) = v429
						*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v430
						v435 = v28 + int32(16)
						if v379&int32(64) != 0 {
							v454 = int64(base.Ui64(v365) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
							v455 = int64(0)
						} else {
							if v379 == int32(0) {
								v454 = v370
								v455 = v365
							} else {
								v450 = base.I64_extend_i32_u(v379)
								v454 = v365<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v370)>>(uint(v450)%64))
								v455 = int64(base.Ui64(v365) >> (uint(v450) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v435))) = v454
						*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v455
						if v379&int32(64) != 0 {
							v477 = int64(base.Ui64(v368) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
							v478 = int64(0)
						} else {
							if v379 == int32(0) {
								v477 = v366
								v478 = v368
							} else {
								v473 = base.I64_extend_i32_u(v379)
								v477 = v368<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v366)>>(uint(v473)%64))
								v478 = int64(base.Ui64(v368) >> (uint(v473) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v28))) = v477
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v478
						v482 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
						v483 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
						v488 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
						v489 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
						v492 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
						v493 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
						v495 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						v496 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
						v504 = v492 | v493
						v505 = v495
						v508 = base.I64_extend_i32_u(base.B2i32(v482|v483 != int64(0))) | (v488 | v489)
						v509 = v496
						v510 = v509 | v34
						v511 = int64(0)
						if v504 == int64(-9223372036854775807-1) {
							v517 = base.B2i32(v508 == v511)
						} else {
							v517 = base.B2i32(v511 <= v504)
						}
						if v517 == int32(0) {
							v521 = v505 + int64(1)
							v537 = v521
							v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
						} else {
							if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
								v537 = v505
								v545 = v510
							} else {
								v533 = v505 + v505&int64(1)
								v537 = v533
								v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
							}
						}
					} else {
						v537 = int64(0)
						v545 = v34
					}
				} else {
					v504 = v365
					v505 = v366
					v508 = v370
					v509 = v368&int64(281474976710655) | base.I64_extend_i32_u(v369)<<(uint(int64(48))%64)
					v510 = v509 | v34
					v511 = int64(0)
					if v504 == int64(-9223372036854775807-1) {
						v517 = base.B2i32(v508 == v511)
					} else {
						v517 = base.B2i32(v511 <= v504)
					}
					if v517 == int32(0) {
						v521 = v505 + int64(1)
						v537 = v521
						v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
					} else {
						if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
							v537 = v505
							v545 = v510
						} else {
							v533 = v505 + v505&int64(1)
							v537 = v533
							v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
						}
					}
				}
			}
		} else {
			v60 = l2 & int64(9223372036854775807)
			v61 = int64(9223090561878065152)
			if v60 == v61 {
				v65 = base.B2i32(l1 == int64(0))
			} else {
				v65 = base.B2i32(base.Ui64(v60) < base.Ui64(v61))
			}
			if v65 == int32(0) {
				v537 = l1
				v545 = l2 | int64(140737488355328)
			} else {
				v73 = l4 & int64(9223372036854775807)
				v74 = int64(9223090561878065152)
				if v73 == v74 {
					v78 = base.B2i32(l3 == int64(0))
				} else {
					v78 = base.B2i32(base.Ui64(v73) < base.Ui64(v74))
				}
				if v78 == int32(0) {
					v537 = l3
					v545 = l4 | int64(140737488355328)
				} else {
					if l1|(v60^int64(9223090561878065152)) == int64(0) {
						if v73|l3 == int64(0) {
							v537 = int64(0)
							v545 = int64(9223231299366420480)
						} else {
							v537 = int64(0)
							v545 = v34 | int64(9223090561878065152)
						}
					} else {
						if l3|(v73^int64(9223090561878065152)) == int64(0) {
							v102 = int64(0)
							if l1|v60 == v102 {
								v537 = v102
								v545 = int64(9223231299366420480)
							} else {
								v537 = v102
								v545 = v34 | int64(9223090561878065152)
							}
						} else {
							if l1|v60 == int64(0) {
								v537 = int64(0)
								v545 = v34
							} else {
								if v73|l3 == int64(0) {
									v537 = int64(0)
									v545 = v34
								} else {
									if base.Ui64(v60) <= base.Ui64(int64(281474976710655)) {
										v119 = v28 + int32(80)
										v121 = base.B2i32(v36 == int64(0))
										if v36 == int64(0) {
											v122 = l1
										} else {
											v122 = v36
										}
										v128 = base.I32_wrap_i64(base.I64_clz(v122) + base.I64_extend_i32_u(v121<<(uint(int32(6))%32)))
										v130 = v128 - int32(15)
										if v130&int32(64) != 0 {
											v149 = int64(0)
											v150 = l1 << (uint(base.I64_extend_i32_u(v130+int32(-64))) % 64)
										} else {
											if v130 == int32(0) {
												v149 = l1
												v150 = v36
											} else {
												v141 = base.I64_extend_i32_u(v130)
												v149 = l1 << (uint(v141) % 64)
												v150 = v36<<(uint(v141)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v130))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v119))) = v149
										*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v150
										v156 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
										v159 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
										v160 = v159
										v161 = int32(16) - v128
										v162 = v156
										v163 = int64(base.Ui64(v156) >> (uint(int64(32)) % 64))
									} else {
										v160 = l1
										v161 = v6
										v162 = v36
										v163 = v38
									}
									if base.Ui64(int64(281474976710655)) < base.Ui64(v73) {
										v207 = v160
										v209 = l3
										v210 = v161
										v213 = v31
										v214 = v162
										v215 = v163
									} else {
										v167 = v28 - int32(-64)
										v169 = base.B2i32(v31 == int64(0))
										if v31 == int64(0) {
											v170 = l3
										} else {
											v170 = v31
										}
										v176 = base.I32_wrap_i64(base.I64_clz(v170) + base.I64_extend_i32_u(v169<<(uint(int32(6))%32)))
										v178 = v176 - int32(15)
										if v178&int32(64) != 0 {
											v197 = int64(0)
											v198 = l3 << (uint(base.I64_extend_i32_u(v178+int32(-64))) % 64)
										} else {
											if v178 == int32(0) {
												v197 = l3
												v198 = v31
											} else {
												v189 = base.I64_extend_i32_u(v178)
												v197 = l3 << (uint(v189) % 64)
												v198 = v31<<(uint(v189)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v178))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v167))) = v197
										*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = v198
										v205 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
										v206 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
										v207 = v160
										v209 = v206
										v210 = v161 - v176 + int32(16)
										v213 = v205
										v214 = v162
										v215 = v163
									}
									v216 = int64(15)
									v217 = v209 << (uint(v216) % 64)
									v219 = v217 & int64(4294934528)
									v220 = int64(32)
									v221 = int64(base.Ui64(v207) >> (uint(v220) % 64))
									v222 = v219 * v221
									v224 = int64(base.Ui64(v217) >> (uint(v220) % 64))
									v225 = int64(4294967295)
									v226 = v207 & v225
									v228 = v222 + v224*v226
									v230 = v228 << (uint(v220) % 64)
									v232 = v230 + v226*v219
									v236 = v214 & v225
									v237 = v219 * v236
									v239 = v237 + v221*v224
									v241 = v213 << (uint(v216) % 64)
									v246 = (v241 | int64(base.Ui64(v209)>>(uint(int64(49))%64))) & v225
									v248 = v239 + v246*v226
									v256 = v248 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v222)))<<(uint(v220)%64) | int64(base.Ui64(v228)>>(uint(v220)%64)))
									v258 = v215 | int64(65536)
									v259 = v219 * v258
									v261 = v259 + v224*v236
									v265 = int64(base.Ui64(v241)>>(uint(v220)%64)) | int64(2147483648)
									v267 = v261 + v265*v226
									v269 = v267 + v246*v221
									v272 = v256 + v269<<(uint(v220)%64)
									v273 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v232) < base.Ui64(v230))) + v272
									v277 = v43 + v48 + v210 - int32(16383)
									v278 = v265 * v221
									v280 = v278 + v224*v258
									v284 = v280 + v246*v236
									v295 = v284 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v237))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v248) < base.Ui64(v239))))
									v299 = v246 * v258
									v301 = v299 + v265*v236
									v312 = v295 + v301<<(uint(v220)%64)
									v329 = v312 + ((base.I64_extend_i32_u(base.B2i32(base.Ui64(v269) < base.Ui64(v267)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v261) < base.Ui64(v259)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v261)))))<<(uint(v220)%64) | int64(base.Ui64(v269)>>(uint(v220)%64)))
									v338 = v329 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v256) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v256))))
									v341 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v284) < base.Ui64(v280))) + v265*v258 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v284))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v301) < base.Ui64(v299)))<<(uint(v220)%64) | int64(base.Ui64(v301)>>(uint(v220)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v295))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v329) < base.Ui64(v312))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v338) < base.Ui64(v329)))
									if v341&int64(281474976710656) != int64(0) {
										v365 = v273
										v366 = v338
										v368 = v341
										v369 = v277 + int32(1)
										v370 = v232
									} else {
										v348 = int64(63)
										v350 = int64(1)
										v365 = int64(base.Ui64(v232)>>(uint(v348)%64)) | v273<<(uint(v350)%64)
										v366 = v338<<(uint(v350)%64) | int64(base.Ui64(v273)>>(uint(v348)%64))
										v368 = v341<<(uint(v350)%64) | int64(base.Ui64(v338)>>(uint(v348)%64))
										v369 = v277
										v370 = v232 << (uint(v350) % 64)
									}
									if int32(32767) <= v369 {
										v537 = int64(0)
										v545 = v34 | int64(9223090561878065152)
									} else {
										if v369 <= int32(0) {
											v379 = int32(1) - v369
											if base.Ui32(v379) <= base.Ui32(int32(127)) {
												v383 = v28 + int32(48)
												v385 = v369 + int32(127)
												if v385&int32(64) != 0 {
													v404 = int64(0)
													v405 = v370 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
												} else {
													if v385 == int32(0) {
														v404 = v370
														v405 = v365
													} else {
														v396 = base.I64_extend_i32_u(v385)
														v404 = v370 << (uint(v396) % 64)
														v405 = v365<<(uint(v396)%64) | int64(base.Ui64(v370)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v383))) = v404
												*(*int64)(unsafe.Add(mBase, uint32(v383)+8)) = v405
												v410 = v28 + int32(32)
												if v385&int32(64) != 0 {
													v429 = int64(0)
													v430 = v366 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
												} else {
													if v385 == int32(0) {
														v429 = v366
														v430 = v368
													} else {
														v421 = base.I64_extend_i32_u(v385)
														v429 = v366 << (uint(v421) % 64)
														v430 = v368<<(uint(v421)%64) | int64(base.Ui64(v366)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v410))) = v429
												*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v430
												v435 = v28 + int32(16)
												if v379&int32(64) != 0 {
													v454 = int64(base.Ui64(v365) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
													v455 = int64(0)
												} else {
													if v379 == int32(0) {
														v454 = v370
														v455 = v365
													} else {
														v450 = base.I64_extend_i32_u(v379)
														v454 = v365<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v370)>>(uint(v450)%64))
														v455 = int64(base.Ui64(v365) >> (uint(v450) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v435))) = v454
												*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v455
												if v379&int32(64) != 0 {
													v477 = int64(base.Ui64(v368) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
													v478 = int64(0)
												} else {
													if v379 == int32(0) {
														v477 = v366
														v478 = v368
													} else {
														v473 = base.I64_extend_i32_u(v379)
														v477 = v368<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v366)>>(uint(v473)%64))
														v478 = int64(base.Ui64(v368) >> (uint(v473) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v28))) = v477
												*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v478
												v482 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
												v483 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
												v488 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
												v489 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
												v492 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
												v493 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
												v495 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
												v496 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
												v504 = v492 | v493
												v505 = v495
												v508 = base.I64_extend_i32_u(base.B2i32(v482|v483 != int64(0))) | (v488 | v489)
												v509 = v496
												v510 = v509 | v34
												v511 = int64(0)
												if v504 == int64(-9223372036854775807-1) {
													v517 = base.B2i32(v508 == v511)
												} else {
													v517 = base.B2i32(v511 <= v504)
												}
												if v517 == int32(0) {
													v521 = v505 + int64(1)
													v537 = v521
													v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
												} else {
													if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
														v537 = v505
														v545 = v510
													} else {
														v533 = v505 + v505&int64(1)
														v537 = v533
														v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
													}
												}
											} else {
												v537 = int64(0)
												v545 = v34
											}
										} else {
											v504 = v365
											v505 = v366
											v508 = v370
											v509 = v368&int64(281474976710655) | base.I64_extend_i32_u(v369)<<(uint(int64(48))%64)
											v510 = v509 | v34
											v511 = int64(0)
											if v504 == int64(-9223372036854775807-1) {
												v517 = base.B2i32(v508 == v511)
											} else {
												v517 = base.B2i32(v511 <= v504)
											}
											if v517 == int32(0) {
												v521 = v505 + int64(1)
												v537 = v521
												v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
											} else {
												if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
													v537 = v505
													v545 = v510
												} else {
													v533 = v505 + v505&int64(1)
													v537 = v533
													v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
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
		v60 = l2 & int64(9223372036854775807)
		v61 = int64(9223090561878065152)
		if v60 == v61 {
			v65 = base.B2i32(l1 == int64(0))
		} else {
			v65 = base.B2i32(base.Ui64(v60) < base.Ui64(v61))
		}
		if v65 == int32(0) {
			v537 = l1
			v545 = l2 | int64(140737488355328)
		} else {
			v73 = l4 & int64(9223372036854775807)
			v74 = int64(9223090561878065152)
			if v73 == v74 {
				v78 = base.B2i32(l3 == int64(0))
			} else {
				v78 = base.B2i32(base.Ui64(v73) < base.Ui64(v74))
			}
			if v78 == int32(0) {
				v537 = l3
				v545 = l4 | int64(140737488355328)
			} else {
				if l1|(v60^int64(9223090561878065152)) == int64(0) {
					if v73|l3 == int64(0) {
						v537 = int64(0)
						v545 = int64(9223231299366420480)
					} else {
						v537 = int64(0)
						v545 = v34 | int64(9223090561878065152)
					}
				} else {
					if l3|(v73^int64(9223090561878065152)) == int64(0) {
						v102 = int64(0)
						if l1|v60 == v102 {
							v537 = v102
							v545 = int64(9223231299366420480)
						} else {
							v537 = v102
							v545 = v34 | int64(9223090561878065152)
						}
					} else {
						if l1|v60 == int64(0) {
							v537 = int64(0)
							v545 = v34
						} else {
							if v73|l3 == int64(0) {
								v537 = int64(0)
								v545 = v34
							} else {
								if base.Ui64(v60) <= base.Ui64(int64(281474976710655)) {
									v119 = v28 + int32(80)
									v121 = base.B2i32(v36 == int64(0))
									if v36 == int64(0) {
										v122 = l1
									} else {
										v122 = v36
									}
									v128 = base.I32_wrap_i64(base.I64_clz(v122) + base.I64_extend_i32_u(v121<<(uint(int32(6))%32)))
									v130 = v128 - int32(15)
									if v130&int32(64) != 0 {
										v149 = int64(0)
										v150 = l1 << (uint(base.I64_extend_i32_u(v130+int32(-64))) % 64)
									} else {
										if v130 == int32(0) {
											v149 = l1
											v150 = v36
										} else {
											v141 = base.I64_extend_i32_u(v130)
											v149 = l1 << (uint(v141) % 64)
											v150 = v36<<(uint(v141)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v130))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v119))) = v149
									*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v150
									v156 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
									v159 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
									v160 = v159
									v161 = int32(16) - v128
									v162 = v156
									v163 = int64(base.Ui64(v156) >> (uint(int64(32)) % 64))
								} else {
									v160 = l1
									v161 = v6
									v162 = v36
									v163 = v38
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v73) {
									v207 = v160
									v209 = l3
									v210 = v161
									v213 = v31
									v214 = v162
									v215 = v163
								} else {
									v167 = v28 - int32(-64)
									v169 = base.B2i32(v31 == int64(0))
									if v31 == int64(0) {
										v170 = l3
									} else {
										v170 = v31
									}
									v176 = base.I32_wrap_i64(base.I64_clz(v170) + base.I64_extend_i32_u(v169<<(uint(int32(6))%32)))
									v178 = v176 - int32(15)
									if v178&int32(64) != 0 {
										v197 = int64(0)
										v198 = l3 << (uint(base.I64_extend_i32_u(v178+int32(-64))) % 64)
									} else {
										if v178 == int32(0) {
											v197 = l3
											v198 = v31
										} else {
											v189 = base.I64_extend_i32_u(v178)
											v197 = l3 << (uint(v189) % 64)
											v198 = v31<<(uint(v189)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v178))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v167))) = v197
									*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = v198
									v205 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
									v206 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
									v207 = v160
									v209 = v206
									v210 = v161 - v176 + int32(16)
									v213 = v205
									v214 = v162
									v215 = v163
								}
								v216 = int64(15)
								v217 = v209 << (uint(v216) % 64)
								v219 = v217 & int64(4294934528)
								v220 = int64(32)
								v221 = int64(base.Ui64(v207) >> (uint(v220) % 64))
								v222 = v219 * v221
								v224 = int64(base.Ui64(v217) >> (uint(v220) % 64))
								v225 = int64(4294967295)
								v226 = v207 & v225
								v228 = v222 + v224*v226
								v230 = v228 << (uint(v220) % 64)
								v232 = v230 + v226*v219
								v236 = v214 & v225
								v237 = v219 * v236
								v239 = v237 + v221*v224
								v241 = v213 << (uint(v216) % 64)
								v246 = (v241 | int64(base.Ui64(v209)>>(uint(int64(49))%64))) & v225
								v248 = v239 + v246*v226
								v256 = v248 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v222)))<<(uint(v220)%64) | int64(base.Ui64(v228)>>(uint(v220)%64)))
								v258 = v215 | int64(65536)
								v259 = v219 * v258
								v261 = v259 + v224*v236
								v265 = int64(base.Ui64(v241)>>(uint(v220)%64)) | int64(2147483648)
								v267 = v261 + v265*v226
								v269 = v267 + v246*v221
								v272 = v256 + v269<<(uint(v220)%64)
								v273 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v232) < base.Ui64(v230))) + v272
								v277 = v43 + v48 + v210 - int32(16383)
								v278 = v265 * v221
								v280 = v278 + v224*v258
								v284 = v280 + v246*v236
								v295 = v284 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v237))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v248) < base.Ui64(v239))))
								v299 = v246 * v258
								v301 = v299 + v265*v236
								v312 = v295 + v301<<(uint(v220)%64)
								v329 = v312 + ((base.I64_extend_i32_u(base.B2i32(base.Ui64(v269) < base.Ui64(v267)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v261) < base.Ui64(v259)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v261)))))<<(uint(v220)%64) | int64(base.Ui64(v269)>>(uint(v220)%64)))
								v338 = v329 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v256) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v256))))
								v341 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v284) < base.Ui64(v280))) + v265*v258 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v284))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v301) < base.Ui64(v299)))<<(uint(v220)%64) | int64(base.Ui64(v301)>>(uint(v220)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v295))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v329) < base.Ui64(v312))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v338) < base.Ui64(v329)))
								if v341&int64(281474976710656) != int64(0) {
									v365 = v273
									v366 = v338
									v368 = v341
									v369 = v277 + int32(1)
									v370 = v232
								} else {
									v348 = int64(63)
									v350 = int64(1)
									v365 = int64(base.Ui64(v232)>>(uint(v348)%64)) | v273<<(uint(v350)%64)
									v366 = v338<<(uint(v350)%64) | int64(base.Ui64(v273)>>(uint(v348)%64))
									v368 = v341<<(uint(v350)%64) | int64(base.Ui64(v338)>>(uint(v348)%64))
									v369 = v277
									v370 = v232 << (uint(v350) % 64)
								}
								if int32(32767) <= v369 {
									v537 = int64(0)
									v545 = v34 | int64(9223090561878065152)
								} else {
									if v369 <= int32(0) {
										v379 = int32(1) - v369
										if base.Ui32(v379) <= base.Ui32(int32(127)) {
											v383 = v28 + int32(48)
											v385 = v369 + int32(127)
											if v385&int32(64) != 0 {
												v404 = int64(0)
												v405 = v370 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
											} else {
												if v385 == int32(0) {
													v404 = v370
													v405 = v365
												} else {
													v396 = base.I64_extend_i32_u(v385)
													v404 = v370 << (uint(v396) % 64)
													v405 = v365<<(uint(v396)%64) | int64(base.Ui64(v370)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v383))) = v404
											*(*int64)(unsafe.Add(mBase, uint32(v383)+8)) = v405
											v410 = v28 + int32(32)
											if v385&int32(64) != 0 {
												v429 = int64(0)
												v430 = v366 << (uint(base.I64_extend_i32_u(v369+int32(63))) % 64)
											} else {
												if v385 == int32(0) {
													v429 = v366
													v430 = v368
												} else {
													v421 = base.I64_extend_i32_u(v385)
													v429 = v366 << (uint(v421) % 64)
													v430 = v368<<(uint(v421)%64) | int64(base.Ui64(v366)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v410))) = v429
											*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v430
											v435 = v28 + int32(16)
											if v379&int32(64) != 0 {
												v454 = int64(base.Ui64(v365) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
												v455 = int64(0)
											} else {
												if v379 == int32(0) {
													v454 = v370
													v455 = v365
												} else {
													v450 = base.I64_extend_i32_u(v379)
													v454 = v365<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v370)>>(uint(v450)%64))
													v455 = int64(base.Ui64(v365) >> (uint(v450) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v435))) = v454
											*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v455
											if v379&int32(64) != 0 {
												v477 = int64(base.Ui64(v368) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
												v478 = int64(0)
											} else {
												if v379 == int32(0) {
													v477 = v366
													v478 = v368
												} else {
													v473 = base.I64_extend_i32_u(v379)
													v477 = v368<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v366)>>(uint(v473)%64))
													v478 = int64(base.Ui64(v368) >> (uint(v473) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v477
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v478
											v482 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
											v483 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
											v488 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
											v489 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v492 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
											v493 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
											v495 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v496 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
											v504 = v492 | v493
											v505 = v495
											v508 = base.I64_extend_i32_u(base.B2i32(v482|v483 != int64(0))) | (v488 | v489)
											v509 = v496
											v510 = v509 | v34
											v511 = int64(0)
											if v504 == int64(-9223372036854775807-1) {
												v517 = base.B2i32(v508 == v511)
											} else {
												v517 = base.B2i32(v511 <= v504)
											}
											if v517 == int32(0) {
												v521 = v505 + int64(1)
												v537 = v521
												v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
											} else {
												if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
													v537 = v505
													v545 = v510
												} else {
													v533 = v505 + v505&int64(1)
													v537 = v533
													v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
												}
											}
										} else {
											v537 = int64(0)
											v545 = v34
										}
									} else {
										v504 = v365
										v505 = v366
										v508 = v370
										v509 = v368&int64(281474976710655) | base.I64_extend_i32_u(v369)<<(uint(int64(48))%64)
										v510 = v509 | v34
										v511 = int64(0)
										if v504 == int64(-9223372036854775807-1) {
											v517 = base.B2i32(v508 == v511)
										} else {
											v517 = base.B2i32(v511 <= v504)
										}
										if v517 == int32(0) {
											v521 = v505 + int64(1)
											v537 = v521
											v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
										} else {
											if v508|(v504^int64(-9223372036854775807-1)) != int64(0) {
												v537 = v505
												v545 = v510
											} else {
												v533 = v505 + v505&int64(1)
												v537 = v533
												v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v505)))
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
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v537
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v545
	m.G0 = v28 + int32(96)
	return
}
func F__mdfd_segpath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_GetRelationPath(m, v8+int32(8), v12, v13, v14, v15, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v8 + int32(80)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(8)
	v23 = F_pg_sprintf(m, l0, int32(41608), v8)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v26 = v8 + int32(8)
	if (v26^l0)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L3
L8:
	;
	goto L3
L9:
	;
	goto L8
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v80)
	if v80&int32(255) == int32(0) {
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v79 = v26
	v80 = v32
	v81 = l0
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v26&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = v26
	v38 = l0
	goto L17
L15:
	;
	v50 = v26
	v52 = l0
	goto L16
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 != v57 {
		v79 = v50
		v80 = v54
		v81 = v52
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v39)
	if v39 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v50 = v46
	v52 = v44
	goto L16
L19:
	;
	v43 = int32(1)
	v44 = v38 + v43
	v46 = v36 + v43
	if v46&int32(3) != 0 {
		v36 = v46
		v38 = v44
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v62 = v50
	v63 = v54
	v64 = v52
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63
	v66 = int32(4)
	v67 = v64 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v70 = v62 + v66
	v74 = int32(-2139062144)
	if (v68|(int32(16843008)-v68))&v74 == v74 {
		v62 = v70
		v63 = v68
		v64 = v67
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v79 = v70
	v80 = v68
	v81 = v67
	goto L10
L24:
	;
	goto L23
L25:
	;
	v88 = v79
	v90 = v81
	goto L26
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)) = uint8(v91)
	v93 = int32(1)
	if v91 != 0 {
		v88 = v88 + v93
		v90 = v90 + v93
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	goto L27
}
func F_macaddrtomacaddr8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc0(m, int32(8))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v10)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v12)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
		v15 = int32(65279)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+3)) = uint16(v15)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v14)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v18)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v20)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v22)
		return v6
	}
}
func F_make_execsql_stmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v7)
	F_initStringInfo(m, v21+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = int32(4642300)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1276]))
	*(*int32)(unsafe.Add(mBase, _consts[1276])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	if l0 != int32(275) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v71 = int32(-1)
	v77 = l0
	v79 = int32(0)
	v85 = int32(1)
	v86 = v71
	v87 = v7
	v88 = v7
	v89 = v71
	v94 = v7
	goto L15
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v43 = int32(370454)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1277])))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 == int32(0) {
		v66 = v46
		v67 = v47
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v67-v66 != 0 {
		goto L3
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v46 != v47 {
		v66 = v46
		v67 = v47
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v51 = v42
	v52 = v43
	goto L9
L9:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v55
		v67 = v56
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v66 = v55
	v67 = v56
	goto L6
L11:
	;
	v59 = int32(1)
	if v55 == v56 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v69 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v69)
	goto L3
L14:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(431666))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L123
	}
L15:
	;
	v95 = F_plpgsql_yylex(m, l3, l4, l5)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1276])) = v34
	if v88 != 0 {
		goto L88
	} else {
		goto L89
	}
L17:
	;
	if base.B2i32(v86 < int32(0))&v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v101 = v100
	goto L20
L19:
	;
	v101 = v86
	goto L20
L20:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v102 != int32(99) {
		v219 = v85
		v221 = v94
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v223 = int32(0)
	if v95 == int32(40) {
		goto L66
	} else {
		goto L67
	}
L22:
	;
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v219 = v85
		v221 = v94
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v95 == int32(348) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v203 != int32(102) {
		goto L60
	} else {
		goto L61
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v85))) = uint8(v197)
	goto L24
L26:
	;
	v197 = int32(111)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if v95 != int32(275) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v113 = int32(436572)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1278])))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v117 == int32(0) {
		v136 = v116
		v137 = v117
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v137-v136 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	goto L30
L32:
	;
	if v116 != v117 {
		v136 = v116
		v137 = v117
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v121 = v112
	v122 = v113
	goto L34
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	if v126 == int32(0) {
		v136 = v125
		v137 = v126
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v136 = v125
	v137 = v126
	goto L31
L36:
	;
	v129 = int32(1)
	if v125 == v126 {
		v121 = v121 + v129
		v122 = v122 + v129
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v197 = int32(114)
	goto L25
L39:
	;
	goto L40
L40:
	;
	v142 = int32(102)
	v143 = int32(264981)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1279])))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v147 == int32(0) {
		v166 = v146
		v167 = v147
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v167-v166 == int32(0) {
		v197 = v142
		goto L25
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	if v146 != v147 {
		v166 = v146
		v167 = v147
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v151 = v112
	v152 = v143
	goto L45
L45:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v156 == int32(0) {
		v166 = v155
		v167 = v156
		goto L42
	} else {
		goto L47
	}
L46:
	;
	v166 = v155
	v167 = v156
	goto L42
L47:
	;
	v159 = int32(1)
	if v155 == v156 {
		v151 = v151 + v159
		v152 = v152 + v159
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v171 = int32(379136)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1280])))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v175 == int32(0) {
		v194 = v174
		v195 = v175
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v195-v194 != 0 {
		goto L24
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	if v174 != v175 {
		v194 = v174
		v195 = v175
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v179 = v112
	v180 = v171
	goto L54
L54:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	if v184 == int32(0) {
		v194 = v183
		v195 = v184
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v194 = v183
	v195 = v184
	goto L51
L56:
	;
	v187 = int32(1)
	if v183 == v184 {
		v179 = v179 + v187
		v180 = v180 + v187
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v197 = v142
	goto L25
L59:
	;
	v219 = v85 + int32(1)
	v221 = v215
	goto L21
L60:
	;
	if v203 != int32(111) {
		v215 = v94
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v215 = int32(1)
	goto L59
L63:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
	if v208 != int32(114) {
		v215 = v94
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
	if v211 != int32(102) {
		v215 = v94
		goto L59
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v232 = int32(1)
	goto L68
L67:
	;
	v232 = v223 - base.B2i32(v95 == int32(41))&base.B2i32(v223 < v79)
	goto L68
L68:
	;
	v233 = v79 + v232
	if v221 == int32(0) {
		v246 = v87
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L16
L70:
	;
	if v95 != int32(332) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	if v95 != int32(59) {
		v251 = v246
		goto L70
	} else {
		goto L76
	}
L72:
	;
	if v233 != 0 {
		v246 = v87
		goto L71
	} else {
		goto L73
	}
L73:
	;
	switch v95 - int32(287) {
	case 0, 3:
		goto L75
	default:
		goto L74
	}
L74:
	;
	v246 = v87 - base.B2i32(v95 == int32(313))&base.B2i32(int32(0) < v87)
	goto L71
L75:
	;
	v251 = v87 + int32(1)
	goto L70
L76:
	;
	if v233 != 0 {
		v251 = v246
		goto L70
	} else {
		goto L77
	}
L77:
	;
	if v246 == int32(0) {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v251 = v246
	goto L70
L79:
	;
	if v95 != 0 {
		v77 = v95
		v79 = v233
		v85 = v219
		v86 = v101
		v87 = v251
		v94 = v221
		goto L15
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v258 = int32(332)
	if l0 == int32(328) {
		v77 = v258
		v79 = v233
		v85 = v219
		v86 = v101
		v87 = v251
		v94 = v221
		goto L15
	} else {
		goto L84
	}
L82:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(260947))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	switch v77 - int32(331) {
	case 0, 6:
		v77 = v258
		v79 = v233
		v85 = v219
		v86 = v101
		v87 = v251
		v94 = v221
		goto L15
	default:
		goto L85
	}
L85:
	;
	if v88 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, _consts[1276])) = int32(0)
	F_read_into_target(m, v21+int32(8), v21+int32(7), l3, l4, l5)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1276])) = int32(2)
	v77 = v95
	v79 = v233
	v85 = v219
	v86 = v101
	v87 = v251
	v88 = int32(1)
	v89 = v261
	v94 = v221
	goto L15
L88:
	;
	F_plpgsql_append_source_text(m, v21+int32(12), l1, v89, l5)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v288 = l1
	goto L90
L90:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_plpgsql_append_source_text(m, v21+int32(12), v288, v289, l5)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	F_appendStringInfoSpaces(m, v21+int32(12), v101-v89)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v288 = v101
	goto L90
L93:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v292 <= int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v360 = F_palloc0(m, int32(80))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L101
	}
L95:
	;
	v302 = v292
	goto L96
L96:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v313+v302-int32(1)))))
	goto L98
L97:
	;
	goto L94
L98:
	;
	if base.B2i32(v317 == int32(32))|base.B2i32(base.Ui32((v317-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v331 = v329 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v331
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v333+v331))) = uint8(v335)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v335 < v337 {
		v302 = v337
		goto L96
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	v362 = F_pstrdup(m, v358)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v364 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v362
	v368 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = v368
	v371 = *(*int32)(unsafe.Add(mBase, _consts[1262]))
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+20)) = uint8(v364)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+12)) = v371
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_pfree(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1255])))
	if v381 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v386 = int32(4549024)
	v387 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v390 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(7243)
	v396 = int32(4541928)
	v397 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v21 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(40)
	v406 = F_raw_parser(m, v385, v384)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v418 = F_palloc0(m, int32(24))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v387
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v411
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = int32(16)
	v422 = int32(0)
	if l1 < v422 {
		v465 = v422
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v465
	v470 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+520))
	v473 = v471 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v470)+520)) = v473
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+18)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v473
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+19)) = uint8(v478)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+20)) = v480
	m.G0 = v21 + int32(48)
	return v418
L110:
	;
	goto L109
L111:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+60))
	if v428 == int32(0) {
		v465 = v422
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v431 = l1 + v428
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v427)+188))
	if base.Ui32(v432) <= base.Ui32(v431) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v427)+196))
	if v441 == int32(0) {
		v465 = v442
		goto L110
	} else {
		goto L117
	}
L114:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v427)+192))
	v441 = v434
	goto L113
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v427)+188)) = v428
	v439 = F_strchr(m, v428, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v427)+192)) = v439
	v441 = v439
	goto L113
L117:
	;
	if base.Ui32(v431) <= base.Ui32(v441) {
		v465 = v442
		goto L110
	} else {
		goto L118
	}
L118:
	;
	v446 = v441
	v448 = v442
	goto L119
L119:
	;
	v451 = int32(1)
	v452 = v448 + v451
	*(*int32)(unsafe.Add(mBase, uint32(v427)+196)) = v452
	v455 = v446 + v451
	*(*int32)(unsafe.Add(mBase, uint32(v427)+188)) = v455
	v458 = F_strchr(m, v455, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v427)+192)) = v458
	if v458 == int32(0) {
		v465 = v452
		goto L110
	} else {
		goto L121
	}
L120:
	;
	v465 = v452
	goto L110
L121:
	;
	if base.Ui32(v458) < base.Ui32(v431) {
		v446 = v458
		v448 = v452
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_one_partition_rbound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v4 = l3
	v8 = F_palloc0(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v16 = F_palloc0(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v16
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v22 = F_palloc0(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	return v8
L7:
	;
	v28 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= v28 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = v28
	goto L9
L9:
	;
	v39 = v32 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39+v40))) = v45
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+24)))
	if v50 == int32(1) {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v59 = v32 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v59 < v60 {
		v32 = v59
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v39))) = v55
	goto L13
L15:
	;
	goto L10
L16:
	;
	F_errmsg_internal(m, int32(297829), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(514067), int32(3456), int32(441270))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_palloc0(m, int32(36))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l3
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(17)
		if l2 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
			v36 = F_list_make2_impl(m, v9+int32(16), v9+int32(12))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v45 = v36
				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v45
				m.G0 = v9 + int32(32)
				return v12
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
			v43 = F_list_make1_impl(m, int32(1), v9+int32(8))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = v43
				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v45
				m.G0 = v9 + int32(32)
				return v12
			}
		}
	}
}
func F_make_orclause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967317)
		return v4
	}
}
func F_make_pathkeys_for_sortclauses(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v17 = F_make_pathkeys_for_sortclauses_extended(m, l0, v7+int32(12), l2, v4, v4, v7+int32(11), v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v17
	}
}
func F_make_tlist_from_pathtarget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v43
L2:
	;
	v13 = v2
	v14 = v2
	goto L7
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if int32(0) < v9 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v43 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v20 = v13 << (uint(int32(2)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21)))
	v25 = v13 + int32(1)
	v27 = int32(0)
	v29 = F_makeTargetEntry(m, v23, base.I32_extend16_s(v25), v27, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v43 = v37
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20+v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v35
	goto L13
L12:
	;
	goto L13
L13:
	;
	v37 = F_lappend(m, v14, v29)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v25 < v39 {
		v13 = v25
		v14 = v37
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
}
func F_makeaclitem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int64
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v16 = F_convert_any_priv_string(m, v10, int32(1684784))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = F_palloc(m, int32(16))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v8
				if v14 != 0 {
					v28 = v16 << (uint(int64(32)) % 64)
				} else {
					v28 = int64(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v16&int64(4294967295) | v28
				return v19
			}
		}
	}
}
func F_makepol_1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
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
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v4
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+14)) = uint16(v4)
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v33 = v4
	goto L3
L3:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = m.T0[v51].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v15+int32(159), v15+int32(152), v15+int32(148), v15+int32(14), v15+int32(13))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	m.G0 = v15 + int32(160)
	return
L5:
	;
	goto L4
L6:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v270 == int32(0) {
		v33 = v262
		goto L3
	} else {
		goto L66
	}
L7:
	;
	v251 = v15 + int32(16) + v243<<(uint(int32(2))%32)
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+2)) = uint16(v252)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	*(*uint8)(unsafe.Add(mBase, uint32(v251))) = uint8(v254)
	v262 = v243 + int32(1)
	goto L6
L8:
	;
	if v67 != int32(32) {
		v243 = v67
		goto L7
	} else {
		goto L62
	}
L9:
	;
	if v33 == int32(0) {
		goto L5
	} else {
		goto L53
	}
L10:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v15)+148))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+14)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+13)))
	m.T0[l1].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l2, l0, v176, v177, v178, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L11:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v152 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L12:
	;
	if v33 == int32(0) {
		goto L5
	} else {
		goto L34
	}
L13:
	;
	F_makepol_1(m, l0, l1, l2)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v52 {
	case 0:
		goto L9
	default:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L12
	}
L16:
	;
	v243 = int32(0)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+159)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_consts[1065])))
	v67 = v33
	goto L19
L19:
	;
	v75 = int32(2)
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+int32(12)+v67<<(uint(v75)%32)))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(v75)%32))+uint32(_consts[1065])))
	if v57 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v243 = int32(0)
	goto L7
L21:
	;
	v91 = v67 - int32(1)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(16)+v91<<(uint(int32(2))%32))+2)))
	v98 = F_palloc0(m, int32(8))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	if v62 <= v83 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v83 <= v62 {
		goto L8
	} else {
		goto L26
	}
L25:
	;
	goto L8
L26:
	;
	goto L21
L27:
	;
	if v78 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v95
	goto L30
L29:
	;
	v103 = int32(0)
	goto L30
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+2)) = uint16(v103)
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)) = uint8(v78)
	v106 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v109 = F_lcons(m, v98, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v109
	if v91 != 0 {
		v67 = v91
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	v262 = v33
	goto L6
L34:
	;
	v120 = v33
	goto L35
L35:
	;
	v131 = v120 - int32(1)
	v134 = v15 + int32(16) + v131<<(uint(int32(2))%32)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v138 = F_palloc0(m, int32(8))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L5
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)) = uint8(v136)
	v141 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v141)
	if v136 == int32(4) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v146 = v135
	goto L40
L39:
	;
	v146 = int32(0)
	goto L40
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)) = uint16(v146)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v149 = F_lcons(m, v138, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v149
	if v131 != 0 {
		v120 = v131
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	v159 = F_errsave_start(m, v152)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L47
	}
L44:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v155 != int32(447) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+4)))
	if v158 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if v159 == int32(0) {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v166
	F_errmsg(m, int32(752750), v15)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, v152, int32(511492), int32(714), int32(313572))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L5
L52:
	;
	v262 = v33
	goto L6
L53:
	;
	v188 = v33
	goto L54
L54:
	;
	v199 = v188 - int32(1)
	v202 = v15 + int32(16) + v199<<(uint(int32(2))%32)
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+2)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v206 = F_palloc0(m, int32(8))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L5
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)) = uint8(v204)
	v209 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v209)
	if v204 == int32(4) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v214 = v203
	goto L59
L58:
	;
	v214 = int32(0)
	goto L59
L59:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v206)+2)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = F_lcons(m, v206, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v217
	if v199 != 0 {
		v188 = v199
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg_internal(m, int32(316811), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(511492), int32(639), int32(331537))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
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
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v273 != int32(447) {
		v33 = v262
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+4)))
	if v276 == int32(0) {
		v33 = v262
		goto L3
	} else {
		goto L68
	}
L68:
	;
	goto L5
}
func F_manifest_process_system_identifier(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v12 != l1 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = l1
		m.T0[v14].(func(*base.Module, int32, int32, int32))(m, l0, int32(39197), v8)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_mask_page_hint_bits(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v6 = v4 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v6)
	return
}
func F_match_boolean_index_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	v3 = l2
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = F_match_index_to_operand(m, v16, v3, l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v50 = v16
			v51 = v15
			v55 = F_makeBoolConst(m, v51, int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v58 = F_make_opclause(m, int32(91), v50, v55, int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					if v58 == int32(0) {
						v92 = v5
						m.G0 = v13 + int32(16)
						return v92
					} else {
						v63 = F_palloc0(m, int32(20))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(281)
							v69 = int32(0)
							v76 = F_make_restrictinfo(m, l0, v58, int32(1), v69, v69, v69, v69, v69, v69, v69)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
								v83 = F_list_make1_impl(m, int32(1), v13+int32(8))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v85
									*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v3)
									*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v85)
									*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v83
									v92 = v63
									m.G0 = v13 + int32(16)
									return v92
								}
							}
						}
					}
				}
			}
		} else {
			if v16 == int32(0) {
				v92 = v5
				m.G0 = v13 + int32(16)
				return v92
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if v23 != int32(53) {
					if v23 != int32(21) {
						v92 = v5
						m.G0 = v13 + int32(16)
						return v92
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
						if v28 != int32(2) {
							v92 = v5
							m.G0 = v13 + int32(16)
							return v92
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v35 = F_match_index_to_operand(m, v34, v3, l3)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 != 0 {
									v50 = v34
									v51 = int32(0)
									v55 = F_makeBoolConst(m, v51, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v58 = F_make_opclause(m, int32(91), v50, v55, int32(0))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											if v58 == int32(0) {
												v92 = v5
												m.G0 = v13 + int32(16)
												return v92
											} else {
												v63 = F_palloc0(m, int32(20))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(281)
													v69 = int32(0)
													v76 = F_make_restrictinfo(m, l0, v58, int32(1), v69, v69, v69, v69, v69, v69, v69)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
														*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
														v83 = F_list_make1_impl(m, int32(1), v13+int32(8))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v85 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v85
															*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v3)
															*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v85)
															*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v83
															v92 = v63
															m.G0 = v13 + int32(16)
															return v92
														}
													}
												}
											}
										}
									}
								} else {
									v92 = v5
									m.G0 = v13 + int32(16)
									return v92
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					if v38 != 0 {
						v42 = v38
						if v42 != int32(2) {
							v92 = v5
							m.G0 = v13 + int32(16)
							return v92
						} else {
							v46 = F_match_index_to_operand(m, v37, v3, l3)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v46 == int32(0) {
									v92 = v5
									m.G0 = v13 + int32(16)
									return v92
								} else {
									v50 = v37
									v51 = int32(0)
									v55 = F_makeBoolConst(m, v51, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v58 = F_make_opclause(m, int32(91), v50, v55, int32(0))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											if v58 == int32(0) {
												v92 = v5
												m.G0 = v13 + int32(16)
												return v92
											} else {
												v63 = F_palloc0(m, int32(20))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(281)
													v69 = int32(0)
													v76 = F_make_restrictinfo(m, l0, v58, int32(1), v69, v69, v69, v69, v69, v69, v69)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
														*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
														v83 = F_list_make1_impl(m, int32(1), v13+int32(8))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v85 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v85
															*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v3)
															*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v85)
															*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v83
															v92 = v63
															m.G0 = v13 + int32(16)
															return v92
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
						v39 = F_match_index_to_operand(m, v37, v3, l3)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								v50 = v37
								v51 = v15
								v55 = F_makeBoolConst(m, v51, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v58 = F_make_opclause(m, int32(91), v50, v55, int32(0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										if v58 == int32(0) {
											v92 = v5
											m.G0 = v13 + int32(16)
											return v92
										} else {
											v63 = F_palloc0(m, int32(20))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(281)
												v69 = int32(0)
												v76 = F_make_restrictinfo(m, l0, v58, int32(1), v69, v69, v69, v69, v69, v69, v69)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
													*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
													v83 = F_list_make1_impl(m, int32(1), v13+int32(8))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v85 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v85
														*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v3)
														*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v85)
														*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v83
														v92 = v63
														m.G0 = v13 + int32(16)
														return v92
													}
												}
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
								v42 = v41
								if v42 != int32(2) {
									v92 = v5
									m.G0 = v13 + int32(16)
									return v92
								} else {
									v46 = F_match_index_to_operand(m, v37, v3, l3)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										if v46 == int32(0) {
											v92 = v5
											m.G0 = v13 + int32(16)
											return v92
										} else {
											v50 = v37
											v51 = int32(0)
											v55 = F_makeBoolConst(m, v51, int32(0))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												v58 = F_make_opclause(m, int32(91), v50, v55, int32(0))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													if v58 == int32(0) {
														v92 = v5
														m.G0 = v13 + int32(16)
														return v92
													} else {
														v63 = F_palloc0(m, int32(20))
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(281)
															v69 = int32(0)
															v76 = F_make_restrictinfo(m, l0, v58, int32(1), v69, v69, v69, v69, v69, v69, v69)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
																*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
																v83 = F_list_make1_impl(m, int32(1), v13+int32(8))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	v85 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v85
																	*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v3)
																	*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v85)
																	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v83
																	v92 = v63
																	m.G0 = v13 + int32(16)
																	return v92
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
func F_matchingsel(m *base.Module, l0 int32) int32 {
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
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_generic_restriction_selectivity(m, v2, v3, v4, v5, v6, float64(0.01))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_Float8GetDatum(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_mcelem_array_contain_overlap_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v28 float32
	_ = v28
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v103 float64
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
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
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v242 float32
	_ = v242
	var v250 int32
	_ = v250
	var v263 float64
	_ = v263
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v283 int32
	_ = v283
	var v286 float64
	_ = v286
	var v297 int32
	_ = v297
	var v308 float64
	_ = v308
	v12 = int32(0)
	if l3 == l1+int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = l2
	goto L3
L2:
	;
	v24 = v12
	goto L3
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*float32)(unsafe.Add(mBase, uint32(v24+l1<<(uint(int32(2))%32))))
	v33 = base.F64_promote_f32(base.F32_mul(v28, float32(0.5)))
	goto L6
L5:
	;
	v33 = float64(0.004999999888241291)
	goto L6
L6:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = base.B2i32(base.Ui32(int32(65535)) < base.Ui32(l1))
	v37 = v35 << (uint(int32(4)) % 32)
	if base.Ui32(int32(65535)) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v72 = int32(-1)
	goto L9
L9:
	;
	if l6 == int32(2751) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	v42 = int32(base.Ui32(l1) >> (uint(int32(16)) % 32))
	goto L12
L11:
	;
	v42 = l1
	goto L12
L12:
	;
	v44 = base.B2i32(base.Ui32(int32(255)) < base.Ui32(v42))
	if base.Ui32(int32(255)) < base.Ui32(v42) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = v37 | int32(8)
	goto L15
L14:
	;
	v45 = v37
	goto L15
L15:
	;
	if base.Ui32(int32(255)) < base.Ui32(v42) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = int32(base.Ui32(v42) >> (uint(int32(8)) % 32))
	goto L18
L17:
	;
	v50 = v42
	goto L18
L18:
	;
	v52 = base.B2i32(base.Ui32(int32(15)) < base.Ui32(v50))
	if base.Ui32(int32(15)) < base.Ui32(v50) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v53 = v45 | int32(4)
	goto L21
L20:
	;
	v53 = v45
	goto L21
L21:
	;
	if base.Ui32(int32(15)) < base.Ui32(v50) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(base.Ui32(v50) >> (uint(int32(4)) % 32))
	goto L24
L23:
	;
	v58 = v50
	goto L24
L24:
	;
	v60 = base.B2i32(base.Ui32(int32(3)) < base.Ui32(v58))
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = v53 | int32(2)
	goto L27
L26:
	;
	v61 = v53
	goto L27
L27:
	;
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v64 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	goto L30
L29:
	;
	v64 = v58
	goto L30
L30:
	;
	v72 = v61 + base.B2i32(base.Ui32(int32(1)) < base.Ui32(v64))
	goto L9
L31:
	;
	v77 = float64(1)
	goto L33
L32:
	;
	v77 = float64(0)
	goto L33
L33:
	;
	if int32(0) < l5 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v80 = float64(0.005)
	if base.F64_gt(v33, v80) != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v308 = v77
	goto L36
L36:
	;
	return v308
L37:
	;
	v83 = v80
	goto L39
L38:
	;
	v83 = v33
	goto L39
L39:
	;
	v85 = l1 - int32(1)
	v87 = l7 + int32(104)
	v100 = int32(0)
	v103 = v77
	v106 = v12
	goto L40
L40:
	;
	if v106 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v308 = v286
	goto L36
L42:
	;
	v297 = v106 + int32(1)
	if v297 != l5 {
		v100 = v283
		v103 = v286
		v106 = v297
		goto L40
	} else {
		goto L87
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v116 = l4 + v106<<(uint(int32(2))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v121 = F_FunctionCall2Coll(m, v87, v113, v119, v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if base.B2i32(v72*l5 < l1+l5) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	return float64(0)
L47:
	;
	if v121 == int32(0) {
		v283 = v100
		v286 = v103
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	if l6 != int32(2751) {
		goto L80
	} else {
		goto L81
	}
L50:
	;
	if v24 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	v186 = v100
	goto L70
L52:
	;
	if v100 < l1 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v85 < v100 {
		v250 = v100
		v263 = v83
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v250 = v100
	v263 = v83
	goto L49
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l4+v106<<(uint(int32(2))%32))))
	v138 = v85
	v142 = v100
	goto L57
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v157 = int32(2)
	v158 = base.I32_div_s(v138+v142, v157)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0+v158<<(uint(v157)%32))))
	v163 = F_FunctionCall2Coll(m, v87, v155, v162, v135)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L46
	} else {
		goto L59
	}
L58:
	;
	v250 = v172
	v263 = v83
	goto L49
L59:
	;
	if v163 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = int32(1)
	v217 = v158
	goto L50
L61:
	;
	goto L62
L62:
	;
	v171 = base.B2i32(v163 < int32(0))
	if v163 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v172 = v158 + int32(1)
	goto L65
L64:
	;
	v172 = v142
	goto L65
L65:
	;
	if v163 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v175 = v138
	goto L68
L67:
	;
	v175 = v158 - int32(1)
	goto L68
L68:
	;
	if v172 <= v175 {
		v138 = v175
		v142 = v172
		goto L57
	} else {
		goto L69
	}
L69:
	;
	goto L58
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0+v186<<(uint(int32(2))%32))))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l4+v106<<(uint(int32(2))%32))))
	v205 = F_FunctionCall2Coll(m, v87, v199, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L46
	} else {
		goto L72
	}
L71:
	;
	v250 = l1
	v263 = v83
	goto L49
L72:
	;
	if int32(0) <= v205 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v216 = base.B2i32(v205 == int32(0))
	v217 = v186
	goto L50
L74:
	;
	goto L75
L75:
	;
	v212 = v186 + int32(1)
	if v212 != l1 {
		v186 = v212
		goto L70
	} else {
		goto L76
	}
L76:
	;
	goto L71
L77:
	;
	v250 = v217
	v263 = v83
	goto L49
L78:
	;
	if v216 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v242 = *(*float32)(unsafe.Add(mBase, uint32(v24+v217<<(uint(int32(2))%32))))
	v250 = v217 + int32(1)
	v263 = base.F64_promote_f32(v242)
	goto L49
L80:
	;
	v268 = base.F64_sub(base.F64_add(v103, v263), base.F64_mul(v103, v263))
	goto L82
L81:
	;
	v268 = base.F64_mul(v103, v263)
	goto L82
L82:
	;
	v269 = float64(0)
	if base.F64_lt(v268, v269) != 0 {
		v283 = v250
		v286 = v269
		goto L42
	} else {
		goto L83
	}
L83:
	;
	if base.F64_gt(v268, float64(1)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v283 = v250
	v286 = v268
	goto L42
L85:
	;
	goto L86
L86:
	;
	v283 = v250
	v286 = float64(1)
	goto L42
L87:
	;
	goto L41
}
func F_md_readv_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	if v18&int32(256) != 0 {
		v21 = v16
	} else {
		v21 = int32(-1)
	}
	F_GetRelationPath(m, v8+int32(56), v12, v13, v14, v21, base.I32_extend8_s(v18))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(512)) <= base.Ui32(v26) {
			*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(base.Ui32(v26) >> (uint(int32(9)) % 32))
			v34 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if v34 == int32(0) {
					m.G0 = v8 + int32(128)
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v40 + v41 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v8 + int32(56)
						F_errmsg(m, int32(310400), v8+int32(32))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v88 = int32(2064)
							F_errfinish(m, int32(520146), v88, int32(86464))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					}
				}
			}
		} else {
			v57 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				if v57 == int32(0) {
					m.G0 = v8 + int32(128)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v67 = int32(13)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v66 << (uint(v67) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v64 << (uint(v67) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v66 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(56)
						F_errmsg(m, int32(167408), v8)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							v88 = int32(2079)
							F_errfinish(m, int32(520146), v88, int32(86464))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_mda_get_prod(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2+l0<<(uint(v7)%32)-int32(4)))) = int32(1)
	v15 = l0 - v7
	if v15 < int32(0) {
	} else {
		if l0&int32(1) == int32(0) {
			v22 = int32(2)
			v28 = l0<<(uint(v22)%32) - int32(4)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1+v28)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2+v28)))
			*(*int32)(unsafe.Add(mBase, uint32(l2+v15<<(uint(v22)%32)))) = v30 * v32
			v37 = l0 - int32(3)
		} else {
			v37 = v15
		}
		if v15 == int32(0) {
		} else {
			v43 = v37
			for {
				v46 = int32(2)
				v47 = v43 << (uint(v46) % 32)
				v50 = v47 + int32(4)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+v50)))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v50)))
				v55 = v52 * v54
				*(*int32)(unsafe.Add(mBase, uint32(l2+v47))) = v55
				v58 = v43 - int32(1)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l2+v58<<(uint(v46)%32)))) = v63 * v55
				if v58 != 0 {
					v43 = v43 - v46
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_mdc_flush(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(5331)
	*(*uint16)(unsafe.Add(mBase, uint32(v6))) = uint16(v8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	m.T0[v11].(func(*base.Module, int32, int32, int32))(m, l1, v6, int32(2))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		m.T0[v18].(func(*base.Module, int32, int32))(m, l1, v6|int32(2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = F_pushf_write(m, l0, v6, int32(22))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v26 = F___memset(m, v6, int32(0), int32(22))
				mBase = m.M
				m.G0 = v6 + int32(32)
				return v22
			}
		}
	}
}
func F_mdc_free_1(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		m.T0[v7].(func(*base.Module, int32))(m, v6)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_mdc_init_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v8 = F_pgp_load_digest(m, int32(2), l1+int32(116))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_mdc_read(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
		if v10 == int32(0) {
			v18 = F_pullf_read(m, l1, l2, l3)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if int32(0) <= v18 {
					if v18 == int32(0) {
						F_px_debug(m, int32(352677), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							return int32(-100)
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
						m.T0[v32].(func(*base.Module, int32, int32, int32))(m, v30, v31, v18)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							return v18
						}
					}
				} else {
					return v18
				}
			}
		} else {
			v13 = F_pullf_read(m, l1, l2, l3)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	} else {
		v13 = F_pullf_read(m, l1, l2, l3)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
func F_mdcreate(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v252
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L15
	} else {
		goto L106
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+40))
	if int32(0) < v16 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = m.G0
	v23 = v21 - int32(160)
	m.G0 = v23
	if v19 != int32(1664) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v11+int32(8), v232, v233, v234, v235, l1)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L15
	} else {
		goto L83
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L15
	} else {
		goto L79
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L15
	} else {
		goto L75
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L71
	}
L11:
	;
	v27 = F_GetDatabasePath(m, v20, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	m.G0 = v23 + int32(160)
	goto L7
L14:
	;
	F_pfree(m, v27)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L15
	} else {
		goto L70
	}
L15:
	;
	return
L16:
	;
	v33 = F___fstatat(m, int32(-100), v27, v23-int32(-64), int32(0))
	mBase = m.M
	goto L17
L17:
	;
	if v33 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v37 == int32(44) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	if v169&int32(61440) != int32(16384) {
		goto L8
	} else {
		goto L69
	}
L21:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v45 = F_LWLockAcquire(m, v41+int32(2432), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L15
	} else {
		goto L65
	}
L24:
	;
	v51 = F___fstatat(m, int32(-100), v27, v23-int32(-64), int32(0))
	mBase = m.M
	goto L26
L25:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v147+int32(2432))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L64
	}
L26:
	;
	if v51 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	if v54&int32(61440) == int32(16384) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	v61 = F_mkdir(m, v27, v60)
	mBase = m.M
	goto L31
L30:
	;
	goto L29
L31:
	;
	if int32(0) <= v61 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if l2 == int32(0) {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v67 != int32(44) {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	v76 = m.G0
	v78 = v76 - int32(96)
	m.G0 = v78
	v81 = F_umask(m, int32(0))
	mBase = m.M
	v84 = F_umask(m, v81&int32(-193))
	mBase = m.M
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v91 = v27 + base.B2i32(v85 == int32(47))
	goto L37
L35:
	;
	if v139 < int32(0) {
		goto L9
	} else {
		goto L63
	}
L36:
	;
	v140 = F_umask(m, v81)
	mBase = m.M
	m.G0 = v78 + int32(96)
	goto L35
L37:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 != int32(47) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v139 = v124 >> (uint(int32(31)) % 32)
	goto L36
L39:
	;
	goto L38
L40:
	;
	v91 = v91 + int32(1)
	goto L37
L41:
	;
	v107 = F_stat(m, v27, v78)
	mBase = m.M
	if v107 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v104 = F_umask(m, v81)
	mBase = m.M
	v106 = int32(0)
	goto L41
L43:
	;
	if v95 != 0 {
		goto L40
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v100)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v103 != 0 {
		v106 = int32(1)
		goto L41
	} else {
		goto L47
	}
L46:
	;
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v98)
	goto L42
L47:
	;
	goto L42
L48:
	;
	v131 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v131)
	goto L40
L49:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v110&int32(61440) != int32(16384) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if v106 != 0 {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	if v106 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v106 != 0 {
		goto L48
	} else {
		goto L58
	}
L55:
	;
	v118 = int32(54)
	goto L57
L56:
	;
	v118 = int32(20)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v118
	v139 = int32(-1)
	goto L36
L58:
	;
	v139 = int32(0)
	goto L36
L59:
	;
	v123 = int32(511)
	goto L61
L60:
	;
	v123 = v71
	goto L61
L61:
	;
	v124 = F_mkdir(m, v27, v123)
	mBase = m.M
	v125 = int32(0)
	if v106&base.B2i32(v125 <= v124) == v125 {
		goto L39
	} else {
		goto L62
	}
L62:
	;
	goto L48
L63:
	;
	goto L25
L64:
	;
	goto L14
L65:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v27
	F_errmsg(m, int32(308100), v23+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(520016), int32(184), int32(436308))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L15
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
	goto L14
L70:
	;
	goto L13
L71:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L15
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v27
	F_errmsg(m, int32(308471), v23+int32(16))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L15
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(520016), int32(158), int32(436308))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v27
	F_errmsg(m, int32(308471), v23)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(520016), int32(174), int32(436308))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v27
	F_errmsg(m, int32(13592), v23+int32(48))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(520016), int32(194), int32(436308))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v243&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v246 = int32(16578)
	goto L86
L85:
	;
	v246 = int32(194)
	goto L86
L86:
	;
	v247 = F_PathNameOpenFile(m, v11+int32(8), v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	if v247 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if l2 == int32(0) {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v269 = v247
	goto L90
L90:
	;
	v272 = l0 + l1<<(uint(int32(2))%32)
	v274 = v272 + int32(40)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v275 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v260&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v263 = int32(16386)
	goto L94
L93:
	;
	v263 = int32(2)
	goto L94
L94:
	;
	v264 = F_PathNameOpenFile(m, v11+int32(8), v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L15
	} else {
		goto L95
	}
L95:
	;
	if v264 < int32(0) {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v269 = v264
	goto L90
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v269
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v300 != int32(-1) {
		goto L2
	} else {
		goto L104
	}
L98:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[884]))
	v281 = F_MemoryContextAlloc(m, v279, int32(8))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L15
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v285 = v272 + int32(56)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	if int32(0) < v275 {
		v293 = v286
		goto L97
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+56)) = v281
	v293 = v281
	goto L97
L102:
	;
	v290 = F_repalloc(m, v286, int32(8))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v290
	v293 = v290
	goto L97
L104:
	;
	F_register_dirty_segment(m, l0, l1, v293)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L15
	} else {
		goto L105
	}
L105:
	;
	goto L2
L106:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(8)
	F_errmsg(m, int32(311157), v11)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(520146), int32(252), int32(370198))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mdstartreadv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	v7 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v32 = F__mdfd_getseg(m, l1, l2, l3, v7, int32(9))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = l3 & int32(131071)
	v37 = int32(131072) - v36
	if base.Ui32(l5) <= base.Ui32(v37) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L105
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v47 = v43 + v44<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v49
	if base.Ui32(l5) < base.Ui32(v37) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L102
	}
L7:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
	if v208&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v56 = l5
	goto L10
L9:
	;
	v56 = v37
	goto L10
L10:
	;
	if base.Ui32(v56) < base.Ui32(int32(2)) {
		v190 = int32(1)
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v59 = int32(1)
	v61 = v56 - v59
	if v56 == int32(2) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v61&v59 == int32(0) {
		v190 = v148
		goto L7
	} else {
		goto L27
	}
L13:
	;
	v146 = v47
	v147 = v49
	v148 = int32(1)
	v150 = v59
	goto L12
L14:
	;
	goto L15
L15:
	;
	v78 = v47
	v79 = v49
	v80 = int32(1)
	v82 = v59
	v91 = v7
	goto L16
L16:
	;
	v98 = v82 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l4+v98)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v100 == v79+v101 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v146 = v132
	v147 = v133
	v148 = v134
	v150 = v136
	goto L12
L18:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l4+int32(4)+v98)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v118 != v115+v119 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v101 - int32(-8192)
	v114 = v78
	v115 = v79
	v116 = v80
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(8192)
	v110 = v78 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v100
	v114 = v110
	v115 = v100
	v116 = v80 + int32(1)
	goto L18
L22:
	;
	v135 = int32(2)
	v136 = v82 + v135
	v138 = v91 + v135
	if v138 != v61&int32(-2) {
		v78 = v132
		v79 = v133
		v80 = v134
		v82 = v136
		v91 = v138
		goto L16
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = int32(8192)
	v125 = v114 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v118
	v132 = v125
	v133 = v118
	v134 = v116 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v119 - int32(-8192)
	v132 = v114
	v133 = v115
	v134 = v116
	goto L22
L26:
	;
	goto L17
L27:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l4+v150<<(uint(int32(2))%32))))
	if v147+v167 != v172 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v172
	v190 = v148 + int32(1)
	goto L7
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v167 - int32(-8192)
	v190 = v148
	goto L7
L31:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v215 = v214 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v215)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v219 = l0 + int32(104)
	goto L35
L34:
	;
	goto L33
L35:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+8)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v219))) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v219)+16)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v219)+12)) = l3
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+21)))
	v231 = int32(8)
	v233 = l2&int32(255) | v230<<(uint(v231)%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+20)) = uint16(v233)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v242 = v233&int32(65279) | base.B2i32(v237 != int32(-1))<<(uint(v231)%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+20)) = uint16(v242)
	v245 = v242 & int32(65023)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+20)) = uint16(v245)
	F_pgaio_io_register_callbacks(m, l0, v220, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v253 = F_FileAccess(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	if v603 != 0 {
		goto L3
	} else {
		goto L101
	}
L38:
	;
	if v253 < int32(0) {
		v603 = int32(-1)
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258+v252*int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = base.I64_extend_i32_u(v36 << (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v262
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+92)) = uint16(v190)
	v266 = m.G0
	v268 = v266 - int32(32)
	m.G0 = v268
	v270 = int32(1)
	v271 = int32(4543676)
	v273 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v273 + v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v270)
	F_pgaio_io_update_state(m, l0, int32(2))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+16)) = int32(0)
	v288 = m.G0
	v290 = v288 - int32(32)
	m.G0 = v290
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v292 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v307 = v292
	goto L44
L42:
	;
	goto L43
L43:
	;
	m.G0 = v290 + int32(32)
	F_pgaio_io_update_state(m, l0, int32(3))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L69
	}
L44:
	;
	v331 = v307 - int32(1)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(5)+v331))))
	v335 = v333 << (uint(int32(3)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[885])))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	if v339 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+(l0+int32(9))))))
	v344 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(int32(1)) < base.Ui32(v307) {
		v307 = v331
		goto L44
	} else {
		goto L68
	}
L49:
	;
	if v344 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errhidestmt(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	m.T0[v404].(func(*base.Module, int32, int32))(m, l0, v341)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L67
	}
L53:
	;
	F_errhidecontext(m)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+24))
	goto L55
L55:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v357) <= base.Ui32(int32(2)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v366<<(uint(int32(2))%32))+uint32(_consts[783])))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	goto L60
L57:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v357<<(uint(int32(2))%32))+uint32(_consts[784])))
	v365 = v364
	goto L59
L58:
	;
	v365 = int32(0)
	goto L59
L59:
	;
	goto L56
L60:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v374) <= base.Ui32(int32(7)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(28)))) = v341
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[886])))
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(24)))) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(20)))) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(16)))) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v290)+12)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v290)+8)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v290)+4)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = (l0 - v352) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(696214), v290)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L65
	}
L62:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v374<<(uint(int32(2))%32))+uint32(_consts[785])))
	v382 = v381
	goto L64
L63:
	;
	v382 = int32(0)
	goto L64
L64:
	;
	goto L61
L65:
	;
	F_errfinish(m, int32(518105), int32(215), int32(419622))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L52
L67:
	;
	goto L48
L68:
	;
	goto L45
L69:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v445&int32(1) != 0 {
		v457 = v270
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v460 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L74
	}
L71:
	;
	v448 = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+16))
	if v451 == v448 {
		v457 = v448
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v454 = m.T0[v451].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v457 = v454
	goto L70
L74:
	;
	if v460 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_errhidestmt(m)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v457 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L78:
	;
	F_errhidecontext(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+24))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v473) <= base.Ui32(int32(2)) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v482<<(uint(int32(2))%32))+uint32(_consts[783])))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	goto L84
L81:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v473<<(uint(int32(2))%32))+uint32(_consts[784])))
	v481 = v480
	goto L83
L82:
	;
	v481 = int32(0)
	goto L83
L83:
	;
	goto L80
L84:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v490) <= base.Ui32(int32(7)) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v490<<(uint(int32(2))%32))+uint32(_consts[785])))
	v498 = v497
	goto L87
L86:
	;
	v498 = int32(0)
	goto L87
L87:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v268)+16)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = (l0 - v468) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v268)+8)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v268)+12)) = v498
	F_errmsg_internal(m, int32(707069), v268)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(516083), int32(459), int32(419607))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L77
L90:
	;
	v568 = int32(4543676)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v570 - int32(1)
	m.G0 = v268 + int32(32)
	v603 = int32(0)
	goto L37
L91:
	;
	v523 = int32(4459632)
	v524 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v524)+22)))
	v527 = v525 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v524)+22)) = uint16(v527)
	*(*int32)(unsafe.Add(mBase, uint32(v524+v525<<(uint(int32(2))%32))+24)) = l0
	v534 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+20)))
	if v535 != 0 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_pgaio_io_update_state(m, l0, int32(4))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v544 = v542 + int32(152)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v542)+156))
	if v545 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+156)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v542)+152)) = v544
	goto L99
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v544
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v542)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v553
	v556 = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v542)+152)) = v556
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v542)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v542)+160)) = v559 + int32(1)
	F_pgaio_io_perform_synchronously(m, l0)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L90
L101:
	;
	m.G0 = v28 + int32(16)
	return
L102:
	;
	F_errmsg_internal(m, int32(18439), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(520146), int32(1009), int32(36912))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v628 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v626*int32(48))+32))
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l3 + v56 - int32(1)
	F_errmsg(m, int32(310298), v28)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(520146), int32(1037), int32(36912))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mdwriteback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	if l3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = l0 + l1<<(uint(int32(2))%32)
	v18 = l2
	v19 = l3
	goto L3
L3:
	;
	v23 = int32(base.Ui32(v18) >> (uint(int32(17)) % 32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(40))))
	if base.Ui32(v24) <= base.Ui32(v23) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(56))))
	v29 = v26 + v23<<(uint(int32(3))%32)
	if v29 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = int32(131072)
	if base.Ui32(v18+v19-int32(1)^v18) < base.Ui32(v33) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v77 = v19 - v43
	if v77 != 0 {
		v18 = v43 + v18
		v19 = v77
		goto L3
	} else {
		goto L16
	}
L8:
	;
	v43 = v19
	goto L10
L9:
	;
	v43 = v33 - v18&int32(131071)
	goto L10
L10:
	;
	if base.I64_extend_i32_u(v43)<<(uint(int64(13))%64) == int64(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v32*int32(48))+37)))
	if v54&int32(64) != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v57 = F_FileAccess(m, v32)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v57 < int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v61 = int32(4155324)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(167772178)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v32*int32(48))))
	v71 = F_fsync(m, v70)
	mBase = m.M
	v73 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(0)
	goto L7
L16:
	;
	goto L4
}
func F_message_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v16 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(421292)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(993)
		v23 = int32(4541928)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[52]))
		*(*int32)(unsafe.Add(mBase, _consts[52])) = v13 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(16)
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+147)) = uint8(v33)
		if l1 != 0 {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v37 = v36
		} else {
			v37 = int32(0)
		}
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+164)) = uint8(v38)
		*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v37
		m.T0[v16].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32))(m, v15, l1, l2, l3, l4, l5, l6)
		mBase = m.M
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		*(*int32)(unsafe.Add(mBase, _consts[52])) = v44
	} else {
	}
	m.G0 = v13 + int32(32)
	return
}
func F_minimal_tuple_from_heap_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 - int32(8)
	v8 = F_palloc(m, v6+l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(int32(0)), l1)
		mBase = m.M
		v15 = v14 + l1
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 != 0 {
			v19 = F__emscripten_memcpy_bulkmem(m, v15, v16+int32(8), v6)
			mBase = m.M
			v20 = v19
		} else {
			v20 = v15
		}
		*(*int32)(unsafe.Add(mBase, uint32(v20))) = v6
		return v20
	}
}
func F_mix_encrypt_resync(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != int32(2) {
		v64 = l1
		v66 = l3
		v67 = v10
		v72 = l2
		if v72 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
			return v72
		} else {
			v83 = v64
			v85 = v66
			v86 = v67
			for {
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(52))))))
				v95 = v92 ^ v94
				*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(84))))) = uint8(v95)
				*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v95)
				v98 = int32(1)
				v103 = v86 + v98
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v105 = v104 + v72
				if v103 < v105 {
					v83 = v83 + v98
					v85 = v85 + v98
					v86 = v103
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
			return v72
		}
	} else {
		v15 = int32(2) - v10
		if l2 < v15 {
			v17 = l2
		} else {
			v17 = v15
		}
		if v17 <= int32(0) {
			v51 = l1
			v53 = l3
			v54 = v10 + v17
		} else {
			v26 = l1
			v28 = l3
			v30 = v10
			for {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(52))))))
				v38 = v35 ^ v37
				*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(84))))) = uint8(v38)
				*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v38)
				v41 = int32(1)
				v42 = v28 + v41
				v44 = v26 + v41
				v46 = v30 + v41
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v48 = v47 + v17
				if v46 < v48 {
					v26 = v44
					v28 = v42
					v30 = v46
					continue
				} else {
					break
				}
				break
			}
			v51 = v44
			v53 = v42
			v54 = v48
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54
		if v54 == int32(2) {
			v110 = l0 + int32(20)
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v115 = v113 - int32(2)
			if v115 != 0 {
				v116 = F__emscripten_memcpy_bulkmem(m, v110, l0+int32(86), v115)
				mBase = m.M
				v117 = v116
			} else {
				v117 = v110
			}
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+84)))
			*(*uint16)(unsafe.Add(mBase, uint32(v117+v115))) = uint16(v119)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			return v17
		} else {
			v64 = v51
			v66 = v53
			v67 = v54
			v72 = l2 - v17
			if v72 <= int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
				return v72
			} else {
				v83 = v64
				v85 = v66
				v86 = v67
				for {
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(52))))))
					v95 = v92 ^ v94
					*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(84))))) = uint8(v95)
					*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v95)
					v98 = int32(1)
					v103 = v86 + v98
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v105 = v104 + v72
					if v103 < v105 {
						v83 = v83 + v98
						v85 = v85 + v98
						v86 = v103
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
				return v72
			}
		}
	}
}
func F_mkANode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
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
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	v6 = int32(0)
	if l2 <= l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = l1
	v30 = v6
	v33 = v6
	goto L4
L4:
	;
	v43 = v23 + v29*int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = int32(base.Ui32(v44)>>(uint(int32(10))%32)) & int32(16383)
	if l3 < v48 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v59 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if l4 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v59 = v30
	v60 = v33
	goto L8
L8:
	;
	v62 = v29 + int32(1)
	if v62 != l2 {
		v29 = v62
		v30 = v59
		v33 = v60
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v54 = l3 ^ int32(-1) + v48
	goto L11
L10:
	;
	v54 = l3
	goto L11
L11:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v54))))
	v59 = v30 + base.B2i32(v33&int32(255) != v56)
	v60 = v56
	goto L8
L12:
	;
	goto L5
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v74 = F_MemoryContextAlloc(m, v68, (l2-l1)<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v79 = v59 * int32(12)
	v81 = v79 + int32(4)
	if base.Ui32(int32(1025)) <= base.Ui32(v81) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v106&v107 | v59<<(uint(v107)%32)
	v114 = l3 + v107
	v116 = l3 ^ int32(-1)
	v121 = l1
	v125 = l1
	v128 = int32(0)
	v129 = v105 + int32(4)
	v134 = v6
	goto L28
L19:
	;
	v84 = F_palloc0(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v89 = (v79 + int32(11)) & int32(4088)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v89) <= base.Ui32(v90) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v105 = v84
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v97 - v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v89 + v98
	v105 = v98
	goto L18
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v97 = v90
	v98 = v92
	goto L23
L25:
	;
	goto L26
L26:
	;
	v93 = int32(8192)
	v95 = F_palloc0(m, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v97 = v93
	v98 = v95
	goto L23
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v139 = v125 * int32(24)
	v140 = v137 + v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v145 = int32(base.Ui32(v141)>>(uint(int32(10))%32)) & int32(16383)
	if v145 <= l3 {
		v248 = v121
		v251 = v128
		v252 = v129
		v255 = v134
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v259 = F_mkANode(m, l0, v248, l2, v114, l4)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L16
	} else {
		goto L67
	}
L30:
	;
	v257 = v125 + int32(1)
	if v257 != l2 {
		v121 = v248
		v125 = v257
		v128 = v251
		v129 = v252
		v134 = v255
		goto L28
	} else {
		goto L66
	}
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	if l4 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v149 = v145 + v116
	goto L34
L33:
	;
	v149 = l3
	goto L34
L34:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v149))))
	if v151 != v134 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v134 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v216 = v121
	v217 = v147
	v218 = v128
	v219 = v129
	v221 = v137
	v222 = v134
	goto L37
L37:
	;
	if l4 != 0 {
		goto L62
	} else {
		goto L63
	}
L38:
	;
	v153 = F_mkANode(m, l0, v121, v125, v114, l4)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L41
	}
L39:
	;
	v200 = v121
	v201 = v147
	v202 = v128
	v203 = v129
	v205 = v137
	goto L40
L40:
	;
	if l4 != 0 {
		goto L59
	} else {
		goto L60
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v153
	if v128 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v156 | v128<<(uint(int32(8))%32)
	v162 = v128 << (uint(int32(2)) % 32)
	if base.Ui32(int32(1025)) <= base.Ui32(v162) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196+v139)+12))
	v200 = v125
	v201 = v198
	v202 = int32(0)
	v203 = v129 + int32(12)
	v205 = v196
	goto L40
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v186
	if v162 != 0 {
		goto L56
	} else {
		goto L57
	}
L46:
	;
	v165 = F_palloc0(m, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v170 = (v162 + int32(7)) & int32(4088)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v170) <= base.Ui32(v171) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v186 = v165
	goto L45
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v178 - v170
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v170 + v179
	v186 = v179
	goto L45
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v178 = v171
	v179 = v173
	goto L50
L52:
	;
	goto L53
L53:
	;
	v174 = int32(8192)
	v176 = F_palloc0(m, v174)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	v178 = v174
	v179 = v176
	goto L50
L55:
	;
	goto L44
L56:
	;
	v188 = F__emscripten_memcpy_bulkmem(m, v186, v74, v162)
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v139+v205)+4))
	v213 = int32(base.Ui32(v207)>>(uint(int32(10))%32))&int32(16383) + v116
	goto L61
L60:
	;
	v213 = l3
	goto L61
L61:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v201))))
	v216 = v200
	v217 = v201
	v218 = v202
	v219 = v203
	v221 = v205
	v222 = v215
	goto L37
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v139+v221)+4))
	v230 = int32(base.Ui32(v224)>>(uint(int32(10))%32))&int32(16383) + v116
	goto L64
L63:
	;
	v230 = l3
	goto L64
L64:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v232)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v235 = v234 + v139
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if int32(base.Ui32(v236)>>(uint(int32(10))%32))&int32(16383) != v114 {
		v248 = v216
		v251 = v218
		v252 = v219
		v255 = v222
		goto L30
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74+v218<<(uint(int32(2))%32)))) = v235
	v248 = v216
	v251 = v218 + int32(1)
	v252 = v219
	v255 = v222
	goto L30
L66:
	;
	goto L29
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v259
	if v251 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v262 | v251<<(uint(int32(8))%32)
	v268 = v251 << (uint(int32(2)) % 32)
	if base.Ui32(int32(1025)) <= base.Ui32(v268) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	goto L70
L70:
	;
	F_pfree(m, v74)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L85
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = v291
	if v268 != 0 {
		goto L82
	} else {
		goto L83
	}
L72:
	;
	v271 = F_palloc0(m, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L16
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v276 = (v268 + int32(7)) & int32(4088)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v276) <= base.Ui32(v277) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v291 = v271
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v284 - v276
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v276 + v285
	v291 = v285
	goto L71
L77:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v284 = v277
	v285 = v279
	goto L76
L78:
	;
	goto L79
L79:
	;
	v280 = int32(8192)
	v282 = F_palloc0(m, v280)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	v284 = v280
	v285 = v282
	goto L76
L81:
	;
	goto L70
L82:
	;
	v294 = F__emscripten_memcpy_bulkmem(m, v291, v74, v268)
	mBase = m.M
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	return v105
}
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_mkdirat(m, int32(-100), l0, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v4) {
		*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0) - v4
		v12 = int32(-1)
	} else {
		v12 = v4
	}
	return v12
}
func F_mpi_check(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 <= v2 {
		v29 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = base.B2i32(v55 == v60)
	if v61 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	if v29 == v13 {
		v55 = v2
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v17 = v2
	goto L4
L4:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v12))))
	if v24 != 0 {
		v29 = v17
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v55 = v2
	goto L1
L6:
	;
	v26 = v17 + int32(1)
	if v26 != v13 {
		v17 = v26
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v40 = (v13 + (v29 ^ int32(-1))) << (uint(int32(3)) % 32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v12))))
	if v42 == int32(0) {
		v55 = v40
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v55 = v40 + (int32(8)-base.I32_clz(v42<<(uint(int32(24))%32)))&int32(255)
	goto L1
L10:
	;
	v64 = int32(0)
	if v13 <= v64 {
		v81 = v64
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	m.G0 = v10 + int32(16)
	return v61
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v60
	F_px_debug(m, int32(484223), v10)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v81 == v13 {
		v107 = v64
		goto L13
	} else {
		goto L20
	}
L15:
	;
	v69 = v64
	goto L16
L16:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v12))))
	if v76 != 0 {
		v81 = v69
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v107 = v64
	goto L13
L18:
	;
	v78 = v69 + int32(1)
	if v78 != v13 {
		v69 = v78
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v92 = (v13 + (v81 ^ int32(-1))) << (uint(int32(3)) % 32)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v12))))
	if v94 == int32(0) {
		v107 = v92
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v107 = v92 + (int32(8)-base.I32_clz(v94<<(uint(int32(24))%32)))&int32(255)
	goto L13
L22:
	;
	return int32(0)
L23:
	;
	goto L12
}
func F_mq_putmessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v1 = l0
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)) = uint8(v1)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[588])))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v99
L2:
	;
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v12 == int32(0) {
		v99 = int32(0)
		goto L1
	} else {
		goto L10
	}
L5:
	;
	F_shm_mq_detach(m, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[587])) = int32(0)
	v99 = int32(-1)
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	goto L7
L10:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[588])) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(31)
	v35 = v12
	goto L11
L11:
	;
	v41 = int32(1)
	v43 = F_shm_mq_sendv(m, v35, v8, int32(2), v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[588])) = uint8(v90)
	if v43 != 0 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	if v49 == v47 {
		v58 = v47
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v43 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	if v58 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
	if v52 != int32(1) {
		v58 = v47
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v58 = base.B2i32(v55 == int32(3))
	goto L17
L20:
	;
	v63 = int32(6)
	goto L22
L21:
	;
	v63 = int32(2)
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[591]))
	v66 = F_SendProcSignal(m, v60, v63, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	v77 = F_WaitLatch(m, v73, int32(33), int32(0), int32(134217762))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L12
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(0)
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v35 = v88
	goto L11
L32:
	;
	goto L31
L33:
	;
	v94 = int32(-1)
	goto L35
L34:
	;
	v94 = v90
	goto L35
L35:
	;
	v99 = v94
	goto L1
}
