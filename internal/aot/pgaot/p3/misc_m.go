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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
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
			v22 = int32(_a_F_MJFillOuter_0)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v25
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) int32)(m, v11, v12, v9+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v23
				if v30 == int32(0) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v64 == int32(0) {
						v73 = v2
					} else {
						v67 = *(*float64)(unsafe.Add(mBase, uint32(v64)+248))
						*(*float64)(unsafe.Add(mBase, uint32(v64)+248)) = base.F64_add(v67, float64(1))
						v73 = v2
					}
					m.G0 = v9 + int32(16)
					return v73
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
					m.T0[v41].(func(*base.Module, int32))(m, v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(_a_F_MJFillOuter_0)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v47
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
						v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v37+int32(4), v38, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v45
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
							v59 = v57 & int32(_a_F_MJFillOuter_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v59)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							*(*uint16)(unsafe.Add(mBase, uint32(v39)+6)) = uint16(v62)
							v73 = v39
							m.G0 = v9 + int32(16)
							return v73
						}
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
			m.T0[v41].(func(*base.Module, int32))(m, v39)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = int32(_a_F_MJFillOuter_0)
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v47
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
				v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v37+int32(4), v38, int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_MJFillOuter[0])) = v45
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
					v59 = v57 & int32(_a_F_MJFillOuter_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v59)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					*(*uint16)(unsafe.Add(mBase, uint32(v39)+6)) = uint16(v62)
					v73 = v39
					m.G0 = v9 + int32(16)
					return v73
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
		v8 = int32(_a_F_MakePerTupleExprContext_0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_MakePerTupleExprContext[0]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_MakePerTupleExprContext[0])) = v11
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
			v28 = F_AllocSetContextCreateInternal(m, v22, int32(_a_F_MakePerTupleExprContext_1), int32(0), int32(_a_F_MakePerTupleExprContext_2), int32(_a_F_MakePerTupleExprContext_3))
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
					*(*int32)(unsafe.Add(mBase, _c_F_MakePerTupleExprContext[0])) = v9
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v97 int64
	_ = v97
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v136 int64
	_ = v136
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v184 int64
	_ = v184
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v262 int64
	_ = v262
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v279 int64
	_ = v279
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v303 int64
	_ = v303
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v317 int64
	_ = v317
	var v322 int64
	_ = v322
	var v326 int64
	_ = v326
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v336 int64
	_ = v336
	var v339 int64
	_ = v339
	var v342 int64
	_ = v342
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v370 int32
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
	var v506 int64
	_ = v506
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
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = int64(281474976710655)
	v25 = l4 & v24
	v28 = (l2 ^ l4) & int64(-9223372036854775807-1)
	v30 = l2 & v24
	v32 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	v33 = int64(48)
	v36 = int32(_a_F___multf3_0)
	v37 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v33)%64))) & v36
	v46 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v33)%64))) & v36
	if base.B2i32(base.Ui32(int32(-32767)) < base.Ui32(v37-v36))&base.B2i32(base.Ui32(int32(-32766)) <= base.Ui32(v46-v36)) != 0 {
		v202 = l1
		v204 = l3
		v205 = v6
		v207 = v25
		v208 = v30
		v210 = v32
		v214 = v37 + v46 + v205 - int32(_a_F___multf3_1)
		v215 = int64(15)
		v216 = v207 << (uint(v215) % 64)
		v217 = int64(32)
		v220 = int64(base.Ui64(v216)>>(uint(v217)%64)) | int64(2147483648)
		v222 = int64(base.Ui64(v202) >> (uint(v217) % 64))
		v223 = v220 * v222
		v225 = v204 << (uint(v215) % 64)
		v227 = int64(base.Ui64(v225) >> (uint(v217) % 64))
		v229 = v210 | int64(65536)
		v231 = v223 + v227*v229
		v237 = int64(4294967295)
		v238 = (int64(base.Ui64(v204)>>(uint(int64(49))%64)) | v216) & v237
		v240 = v208 & v237
		v242 = v231 + v238*v240
		v249 = v225 & int64(4294934528)
		v250 = v249 * v240
		v252 = v250 + v222*v227
		v256 = v202 & v237
		v258 = v252 + v238*v256
		v262 = v242 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v252) < base.Ui64(v250))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v258) < base.Ui64(v252))))
		v266 = v238 * v229
		v268 = v266 + v220*v240
		v279 = v262 + v268<<(uint(v217)%64)
		v283 = v229 * v249
		v285 = v283 + v227*v240
		v287 = v285 + v256*v220
		v289 = v287 + v238*v222
		v303 = v279 + (int64(base.Ui64(v289)>>(uint(v217)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v289) < base.Ui64(v287)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v285) < base.Ui64(v283)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v287) < base.Ui64(v285)))))<<(uint(v217)%64))
		v307 = v222 * v249
		v309 = v307 + v256*v227
		v317 = v258 + (int64(base.Ui64(v309)>>(uint(v217)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v309) < base.Ui64(v307)))<<(uint(v217)%64))
		v322 = v317 + v289<<(uint(v217)%64)
		v326 = v303 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v317) < base.Ui64(v258))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v322) < base.Ui64(v317))))
		v331 = v309 << (uint(v217) % 64)
		v333 = v331 + v256*v249
		v336 = v322 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v333) < base.Ui64(v331)))
		v339 = v326 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v336) < base.Ui64(v322)))
		v342 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v231) < base.Ui64(v223))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v242) < base.Ui64(v231))) + v220*v229 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v262) < base.Ui64(v242))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v268) < base.Ui64(v266)))<<(uint(v217)%64) | int64(base.Ui64(v268)>>(uint(v217)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v279) < base.Ui64(v262))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v303) < base.Ui64(v279))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v326) < base.Ui64(v303))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v339) < base.Ui64(v326)))
		if v342&int64(281474976710656) != int64(0) {
			v366 = v333
			v367 = v336
			v368 = v342
			v369 = v339
			v370 = v214 + int32(1)
		} else {
			v349 = int64(63)
			v351 = int64(1)
			v366 = v333 << (uint(v351) % 64)
			v367 = int64(base.Ui64(v333)>>(uint(v349)%64)) | v336<<(uint(v351)%64)
			v368 = v342<<(uint(v351)%64) | int64(base.Ui64(v339)>>(uint(v349)%64))
			v369 = v339<<(uint(v351)%64) | int64(base.Ui64(v336)>>(uint(v349)%64))
			v370 = v214
		}
		if int32(_a_F___multf3_0) <= v370 {
			v537 = int64(0)
			v545 = v28 | int64(9223090561878065152)
		} else {
			if v370 <= int32(0) {
				v379 = int32(1) - v370
				if base.Ui32(v379) <= base.Ui32(int32(127)) {
					v383 = v22 + int32(48)
					v385 = v370 + int32(127)
					if v385&int32(64) != 0 {
						v404 = int64(0)
						v405 = v366 << (uint(base.I64_extend_i32_u(v370+int32(63))) % 64)
					} else {
						if v385 == int32(0) {
							v404 = v366
							v405 = v367
						} else {
							v396 = base.I64_extend_i32_u(v385)
							v404 = v366 << (uint(v396) % 64)
							v405 = v367<<(uint(v396)%64) | int64(base.Ui64(v366)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v383))) = v404
					*(*int64)(unsafe.Add(mBase, uint32(v383)+8)) = v405
					v410 = v22 + int32(32)
					if v385&int32(64) != 0 {
						v429 = int64(0)
						v430 = v369 << (uint(base.I64_extend_i32_u(v370+int32(63))) % 64)
					} else {
						if v385 == int32(0) {
							v429 = v369
							v430 = v368
						} else {
							v421 = base.I64_extend_i32_u(v385)
							v429 = v369 << (uint(v421) % 64)
							v430 = v368<<(uint(v421)%64) | int64(base.Ui64(v369)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v410))) = v429
					*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v430
					v435 = v22 + int32(16)
					if v379&int32(64) != 0 {
						v454 = int64(base.Ui64(v367) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
						v455 = int64(0)
					} else {
						if v379 == int32(0) {
							v454 = v366
							v455 = v367
						} else {
							v450 = base.I64_extend_i32_u(v379)
							v454 = v367<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v366)>>(uint(v450)%64))
							v455 = int64(base.Ui64(v367) >> (uint(v450) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v435))) = v454
					*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v455
					if v379&int32(64) != 0 {
						v477 = int64(base.Ui64(v368) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
						v478 = int64(0)
					} else {
						if v379 == int32(0) {
							v477 = v369
							v478 = v368
						} else {
							v473 = base.I64_extend_i32_u(v379)
							v477 = v368<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v369)>>(uint(v473)%64))
							v478 = int64(base.Ui64(v368) >> (uint(v473) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v22))) = v477
					*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v478
					v482 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
					v483 = *(*int64)(unsafe.Add(mBase, uint32(v22)+56))
					v488 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
					v489 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
					v492 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
					v493 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
					v495 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					v496 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
					v504 = base.I64_extend_i32_u(base.B2i32(v482|v483 != int64(0))) | (v488 | v489)
					v505 = v492 | v493
					v506 = v495
					v509 = v496
					v510 = v509 | v28
					v511 = int64(0)
					if v505 == int64(-9223372036854775807-1) {
						v517 = base.B2i32(v504 == v511)
					} else {
						v517 = base.B2i32(v511 <= v505)
					}
					if v517 == int32(0) {
						v521 = v506 + int64(1)
						v537 = v521
						v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
					} else {
						if v504|(v505^int64(-9223372036854775807-1)) != int64(0) {
							v537 = v506
							v545 = v510
						} else {
							v533 = v506 + v506&int64(1)
							v537 = v533
							v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v506)))
						}
					}
				} else {
					v537 = int64(0)
					v545 = v28
				}
			} else {
				v504 = v366
				v505 = v367
				v506 = v369
				v509 = v368&int64(281474976710655) | base.I64_extend_i32_u(v370)<<(uint(int64(48))%64)
				v510 = v509 | v28
				v511 = int64(0)
				if v505 == int64(-9223372036854775807-1) {
					v517 = base.B2i32(v504 == v511)
				} else {
					v517 = base.B2i32(v511 <= v505)
				}
				if v517 == int32(0) {
					v521 = v506 + int64(1)
					v537 = v521
					v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
				} else {
					if v504|(v505^int64(-9223372036854775807-1)) != int64(0) {
						v537 = v506
						v545 = v510
					} else {
						v533 = v506 + v506&int64(1)
						v537 = v533
						v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v506)))
					}
				}
			}
		}
	} else {
		v55 = l2 & int64(9223372036854775807)
		v56 = int64(9223090561878065152)
		if v55 == v56 {
			v60 = base.B2i32(l1 == int64(0))
		} else {
			v60 = base.B2i32(base.Ui64(v55) < base.Ui64(v56))
		}
		if v60 == int32(0) {
			v537 = l1
			v545 = l2 | int64(140737488355328)
		} else {
			v68 = l4 & int64(9223372036854775807)
			v69 = int64(9223090561878065152)
			if v68 == v69 {
				v73 = base.B2i32(l3 == int64(0))
			} else {
				v73 = base.B2i32(base.Ui64(v68) < base.Ui64(v69))
			}
			if v73 == int32(0) {
				v537 = l3
				v545 = l4 | int64(140737488355328)
			} else {
				if l1|(v55^int64(9223090561878065152)) == int64(0) {
					if v68|l3 == int64(0) {
						v537 = int64(0)
						v545 = int64(9223231299366420480)
					} else {
						v537 = int64(0)
						v545 = v28 | int64(9223090561878065152)
					}
				} else {
					if l3|(v68^int64(9223090561878065152)) == int64(0) {
						v97 = int64(0)
						if l1|v55 == v97 {
							v537 = v97
							v545 = int64(9223231299366420480)
						} else {
							v537 = v97
							v545 = v28 | int64(9223090561878065152)
						}
					} else {
						if l1|v55 == int64(0) {
							v537 = int64(0)
							v545 = v28
						} else {
							if v68|l3 == int64(0) {
								v537 = int64(0)
								v545 = v28
							} else {
								if base.Ui64(v55) <= base.Ui64(int64(281474976710655)) {
									v114 = v22 + int32(80)
									v116 = base.B2i32(v30 == int64(0))
									if v30 == int64(0) {
										v117 = l1
									} else {
										v117 = v30
									}
									if v30 == int64(0) {
										v121 = int64(64)
									} else {
										v121 = int64(0)
									}
									v123 = base.I32_wrap_i64(base.I64_clz(v117) + v121)
									v125 = v123 - int32(15)
									if v125&int32(64) != 0 {
										v144 = int64(0)
										v145 = l1 << (uint(base.I64_extend_i32_u(v125+int32(-64))) % 64)
									} else {
										if v125 == int32(0) {
											v144 = l1
											v145 = v30
										} else {
											v136 = base.I64_extend_i32_u(v125)
											v144 = l1 << (uint(v136) % 64)
											v145 = v30<<(uint(v136)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v125))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v114))) = v144
									*(*int64)(unsafe.Add(mBase, uint32(v114)+8)) = v145
									v151 = *(*int64)(unsafe.Add(mBase, uint32(v22)+88))
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v22)+80))
									v155 = v154
									v156 = int32(16) - v123
									v157 = v151
									v158 = int64(base.Ui64(v151) >> (uint(int64(32)) % 64))
								} else {
									v155 = l1
									v156 = v6
									v157 = v30
									v158 = v32
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v68) {
									v202 = v155
									v204 = l3
									v205 = v156
									v207 = v25
									v208 = v157
									v210 = v158
								} else {
									v162 = v22 - int32(-64)
									v164 = base.B2i32(v25 == int64(0))
									if v25 == int64(0) {
										v165 = l3
									} else {
										v165 = v25
									}
									if v25 == int64(0) {
										v169 = int64(64)
									} else {
										v169 = int64(0)
									}
									v171 = base.I32_wrap_i64(base.I64_clz(v165) + v169)
									v173 = v171 - int32(15)
									if v173&int32(64) != 0 {
										v192 = int64(0)
										v193 = l3 << (uint(base.I64_extend_i32_u(v173+int32(-64))) % 64)
									} else {
										if v173 == int32(0) {
											v192 = l3
											v193 = v25
										} else {
											v184 = base.I64_extend_i32_u(v173)
											v192 = l3 << (uint(v184) % 64)
											v193 = v25<<(uint(v184)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v173))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v162))) = v192
									*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v193
									v200 = *(*int64)(unsafe.Add(mBase, uint32(v22)+72))
									v201 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
									v202 = v155
									v204 = v201
									v205 = v156 - v171 + int32(16)
									v207 = v200
									v208 = v157
									v210 = v158
								}
								v214 = v37 + v46 + v205 - int32(_a_F___multf3_1)
								v215 = int64(15)
								v216 = v207 << (uint(v215) % 64)
								v217 = int64(32)
								v220 = int64(base.Ui64(v216)>>(uint(v217)%64)) | int64(2147483648)
								v222 = int64(base.Ui64(v202) >> (uint(v217) % 64))
								v223 = v220 * v222
								v225 = v204 << (uint(v215) % 64)
								v227 = int64(base.Ui64(v225) >> (uint(v217) % 64))
								v229 = v210 | int64(65536)
								v231 = v223 + v227*v229
								v237 = int64(4294967295)
								v238 = (int64(base.Ui64(v204)>>(uint(int64(49))%64)) | v216) & v237
								v240 = v208 & v237
								v242 = v231 + v238*v240
								v249 = v225 & int64(4294934528)
								v250 = v249 * v240
								v252 = v250 + v222*v227
								v256 = v202 & v237
								v258 = v252 + v238*v256
								v262 = v242 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v252) < base.Ui64(v250))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v258) < base.Ui64(v252))))
								v266 = v238 * v229
								v268 = v266 + v220*v240
								v279 = v262 + v268<<(uint(v217)%64)
								v283 = v229 * v249
								v285 = v283 + v227*v240
								v287 = v285 + v256*v220
								v289 = v287 + v238*v222
								v303 = v279 + (int64(base.Ui64(v289)>>(uint(v217)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v289) < base.Ui64(v287)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v285) < base.Ui64(v283)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v287) < base.Ui64(v285)))))<<(uint(v217)%64))
								v307 = v222 * v249
								v309 = v307 + v256*v227
								v317 = v258 + (int64(base.Ui64(v309)>>(uint(v217)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v309) < base.Ui64(v307)))<<(uint(v217)%64))
								v322 = v317 + v289<<(uint(v217)%64)
								v326 = v303 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v317) < base.Ui64(v258))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v322) < base.Ui64(v317))))
								v331 = v309 << (uint(v217) % 64)
								v333 = v331 + v256*v249
								v336 = v322 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v333) < base.Ui64(v331)))
								v339 = v326 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v336) < base.Ui64(v322)))
								v342 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v231) < base.Ui64(v223))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v242) < base.Ui64(v231))) + v220*v229 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v262) < base.Ui64(v242))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v268) < base.Ui64(v266)))<<(uint(v217)%64) | int64(base.Ui64(v268)>>(uint(v217)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v279) < base.Ui64(v262))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v303) < base.Ui64(v279))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v326) < base.Ui64(v303))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v339) < base.Ui64(v326)))
								if v342&int64(281474976710656) != int64(0) {
									v366 = v333
									v367 = v336
									v368 = v342
									v369 = v339
									v370 = v214 + int32(1)
								} else {
									v349 = int64(63)
									v351 = int64(1)
									v366 = v333 << (uint(v351) % 64)
									v367 = int64(base.Ui64(v333)>>(uint(v349)%64)) | v336<<(uint(v351)%64)
									v368 = v342<<(uint(v351)%64) | int64(base.Ui64(v339)>>(uint(v349)%64))
									v369 = v339<<(uint(v351)%64) | int64(base.Ui64(v336)>>(uint(v349)%64))
									v370 = v214
								}
								if int32(_a_F___multf3_0) <= v370 {
									v537 = int64(0)
									v545 = v28 | int64(9223090561878065152)
								} else {
									if v370 <= int32(0) {
										v379 = int32(1) - v370
										if base.Ui32(v379) <= base.Ui32(int32(127)) {
											v383 = v22 + int32(48)
											v385 = v370 + int32(127)
											if v385&int32(64) != 0 {
												v404 = int64(0)
												v405 = v366 << (uint(base.I64_extend_i32_u(v370+int32(63))) % 64)
											} else {
												if v385 == int32(0) {
													v404 = v366
													v405 = v367
												} else {
													v396 = base.I64_extend_i32_u(v385)
													v404 = v366 << (uint(v396) % 64)
													v405 = v367<<(uint(v396)%64) | int64(base.Ui64(v366)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v383))) = v404
											*(*int64)(unsafe.Add(mBase, uint32(v383)+8)) = v405
											v410 = v22 + int32(32)
											if v385&int32(64) != 0 {
												v429 = int64(0)
												v430 = v369 << (uint(base.I64_extend_i32_u(v370+int32(63))) % 64)
											} else {
												if v385 == int32(0) {
													v429 = v369
													v430 = v368
												} else {
													v421 = base.I64_extend_i32_u(v385)
													v429 = v369 << (uint(v421) % 64)
													v430 = v368<<(uint(v421)%64) | int64(base.Ui64(v369)>>(uint(base.I64_extend_i32_u(int32(64)-v385))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v410))) = v429
											*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v430
											v435 = v22 + int32(16)
											if v379&int32(64) != 0 {
												v454 = int64(base.Ui64(v367) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
												v455 = int64(0)
											} else {
												if v379 == int32(0) {
													v454 = v366
													v455 = v367
												} else {
													v450 = base.I64_extend_i32_u(v379)
													v454 = v367<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v366)>>(uint(v450)%64))
													v455 = int64(base.Ui64(v367) >> (uint(v450) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v435))) = v454
											*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v455
											if v379&int32(64) != 0 {
												v477 = int64(base.Ui64(v368) >> (uint(base.I64_extend_i32_u(v379+int32(-64))) % 64))
												v478 = int64(0)
											} else {
												if v379 == int32(0) {
													v477 = v369
													v478 = v368
												} else {
													v473 = base.I64_extend_i32_u(v379)
													v477 = v368<<(uint(base.I64_extend_i32_u(int32(64)-v379))%64) | int64(base.Ui64(v369)>>(uint(v473)%64))
													v478 = int64(base.Ui64(v368) >> (uint(v473) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v22))) = v477
											*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v478
											v482 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
											v483 = *(*int64)(unsafe.Add(mBase, uint32(v22)+56))
											v488 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
											v489 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
											v492 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
											v493 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
											v495 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
											v496 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
											v504 = base.I64_extend_i32_u(base.B2i32(v482|v483 != int64(0))) | (v488 | v489)
											v505 = v492 | v493
											v506 = v495
											v509 = v496
											v510 = v509 | v28
											v511 = int64(0)
											if v505 == int64(-9223372036854775807-1) {
												v517 = base.B2i32(v504 == v511)
											} else {
												v517 = base.B2i32(v511 <= v505)
											}
											if v517 == int32(0) {
												v521 = v506 + int64(1)
												v537 = v521
												v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
											} else {
												if v504|(v505^int64(-9223372036854775807-1)) != int64(0) {
													v537 = v506
													v545 = v510
												} else {
													v533 = v506 + v506&int64(1)
													v537 = v533
													v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v506)))
												}
											}
										} else {
											v537 = int64(0)
											v545 = v28
										}
									} else {
										v504 = v366
										v505 = v367
										v506 = v369
										v509 = v368&int64(281474976710655) | base.I64_extend_i32_u(v370)<<(uint(int64(48))%64)
										v510 = v509 | v28
										v511 = int64(0)
										if v505 == int64(-9223372036854775807-1) {
											v517 = base.B2i32(v504 == v511)
										} else {
											v517 = base.B2i32(v511 <= v505)
										}
										if v517 == int32(0) {
											v521 = v506 + int64(1)
											v537 = v521
											v545 = v510 + base.I64_extend_i32_u(base.B2i32(v521 == int64(0)))
										} else {
											if v504|(v505^int64(-9223372036854775807-1)) != int64(0) {
												v537 = v506
												v545 = v510
											} else {
												v533 = v506 + v506&int64(1)
												v537 = v533
												v545 = v510 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v533) < base.Ui64(v506)))
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
	m.G0 = v22 + int32(96)
	return
}
func F__mdfd_segpath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = v9 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_GetRelationPath(m, v12, v13, v14, v15, v16, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	m.G0 = v9 + int32(80)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	v22 = F_pg_sprintf(m, l0, int32(_a_F__mdfd_segpath_0), v9)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v25 = v9 + int32(8)
	if (v25^l0)&int32(3) != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v79)
	if v79&int32(255) == int32(0) {
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v78 = v25
	v79 = v31
	v80 = l0
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v25&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = v25
	v37 = l0
	goto L17
L15:
	;
	v49 = v25
	v51 = l0
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 != v56 {
		v78 = v49
		v79 = v53
		v80 = v51
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
	if v38 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v49 = v45
	v51 = v43
	goto L16
L19:
	;
	v42 = int32(1)
	v43 = v37 + v42
	v45 = v35 + v42
	if v45&int32(3) != 0 {
		v35 = v45
		v37 = v43
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v61 = v49
	v62 = v53
	v63 = v51
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v62
	v65 = int32(4)
	v66 = v63 + v65
	v68 = v61 + v65
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = int32(-2139062144)
	if (int32(16843008)-v70|v70)&v73 == v73 {
		v61 = v68
		v62 = v70
		v63 = v66
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v78 = v68
	v79 = v70
	v80 = v66
	goto L10
L24:
	;
	goto L23
L25:
	;
	v87 = v78
	v89 = v80
	goto L26
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)) = uint8(v90)
	v92 = int32(1)
	if v90 != 0 {
		v87 = v87 + v92
		v89 = v89 + v92
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
		v15 = int32(_a_F_macaddrtomacaddr8_0)
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
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_palloc0(m, int32(108))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+17)) = uint8(v13)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967457)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
		v28 = F_list_make1_impl(m, int32(1), v6+int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)) = uint8(v30)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)) = uint16(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v28
			m.G0 = v6 + int32(16)
			return v9
		}
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
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
	v33 = int32(_a_F_make_execsql_stmt_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[0])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	if l0 != int32(275) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v72 = int32(-1)
	v78 = l0
	v80 = int32(0)
	v86 = v7
	v87 = v72
	v88 = int32(1)
	v89 = v7
	v90 = v72
	v95 = v7
	goto L14
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v43 = int32(_a_F_make_execsql_stmt_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_execsql_stmt[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v67-v68 != 0 {
		goto L3
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v52 = v42
	v53 = v43
	goto L8
L8:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v67 = v57
	v68 = v56
	goto L6
L10:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v70 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v70)
	goto L3
L13:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(_a_F_make_execsql_stmt_2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L117
	}
L14:
	;
	v96 = F_plpgsql_yylex(m, l3, l4, l5)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[0])) = v34
	v289 = v21 + int32(12)
	if v89 != 0 {
		goto L83
	} else {
		goto L84
	}
L16:
	;
	if base.B2i32(v87 < int32(0))&v89 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v102 = v101
	goto L19
L18:
	;
	v102 = v87
	goto L19
L19:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.B2i32(v103 != int32(99))|base.B2i32(base.Ui32(int32(3)) < base.Ui32(v88)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v96 == int32(348) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v226 = v88
	v228 = v95
	goto L22
L22:
	;
	v230 = int32(0)
	if v96 == int32(40) {
		goto L64
	} else {
		goto L65
	}
L23:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v210 != int32(102) {
		goto L56
	} else {
		goto L57
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v88))) = uint8(v204)
	goto L23
L25:
	;
	v204 = int32(111)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v96 != int32(275) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v117 = int32(_a_F_make_execsql_stmt_3)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_execsql_stmt[2])))
	if base.B2i32(v120 == int32(0))|base.B2i32(v120 != v123) != 0 {
		v141 = v120
		v142 = v123
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v141-v142 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v126 = v116
	v127 = v117
	goto L32
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v131 == int32(0) {
		v141 = v131
		v142 = v130
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v141 = v131
	v142 = v130
	goto L30
L34:
	;
	v134 = int32(1)
	if v131 == v130 {
		v126 = v126 + v134
		v127 = v127 + v134
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v204 = int32(114)
	goto L24
L37:
	;
	goto L38
L38:
	;
	v147 = int32(102)
	v148 = int32(_a_F_make_execsql_stmt_4)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_execsql_stmt[3])))
	if base.B2i32(v151 == int32(0))|base.B2i32(v151 != v154) != 0 {
		v172 = v151
		v173 = v154
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v172-v173 == int32(0) {
		v204 = v147
		goto L24
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	v157 = v116
	v158 = v148
	goto L42
L42:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v162 == int32(0) {
		v172 = v162
		v173 = v161
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v172 = v162
	v173 = v161
	goto L40
L44:
	;
	v165 = int32(1)
	if v162 == v161 {
		v157 = v157 + v165
		v158 = v158 + v165
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v177 = int32(_a_F_make_execsql_stmt_5)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_execsql_stmt[4])))
	if base.B2i32(v180 == int32(0))|base.B2i32(v180 != v183) != 0 {
		v201 = v180
		v202 = v183
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v201-v202 != 0 {
		goto L23
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v186 = v116
	v187 = v177
	goto L50
L50:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if v191 == int32(0) {
		v201 = v191
		v202 = v190
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v201 = v191
	v202 = v190
	goto L48
L52:
	;
	v194 = int32(1)
	if v191 == v190 {
		v186 = v186 + v194
		v187 = v187 + v194
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v204 = v147
	goto L24
L55:
	;
	v226 = v88 + int32(1)
	v228 = v222
	goto L22
L56:
	;
	if v210 != int32(111) {
		v222 = v95
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v222 = int32(1)
	goto L55
L59:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
	if v215 != int32(114) {
		v222 = v95
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
	if v218 != int32(102) {
		v222 = v95
		goto L55
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	goto L15
L63:
	;
	if v96 != int32(332) {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v239 = int32(1)
	goto L66
L65:
	;
	v239 = v230 - base.B2i32(v96 == int32(41))&base.B2i32(v230 < v80)
	goto L66
L66:
	;
	v240 = v80 + v239
	v241 = int32(0)
	if v240|base.B2i32(v228 == v241) == v241 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	switch v96 - int32(287) {
	case 0, 3:
		goto L71
	default:
		goto L70
	}
L68:
	;
	v256 = v86
	goto L69
L69:
	;
	if base.B2i32(v96 != int32(59))|v240 != 0 {
		v262 = v256
		goto L63
	} else {
		goto L72
	}
L70:
	;
	v256 = v86 - base.B2i32(v96 == int32(313))&base.B2i32(int32(0) < v86)
	goto L69
L71:
	;
	v262 = v86 + int32(1)
	goto L63
L72:
	;
	if v256 == int32(0) {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	v262 = v256
	goto L63
L74:
	;
	if v96 != 0 {
		v78 = v96
		v80 = v240
		v86 = v262
		v87 = v102
		v88 = v226
		v95 = v228
		goto L14
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v269 = int32(332)
	if l0 == int32(328) {
		v78 = v269
		v80 = v240
		v86 = v262
		v87 = v102
		v88 = v226
		v95 = v228
		goto L14
	} else {
		goto L79
	}
L77:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(_a_F_make_execsql_stmt_6))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
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
	switch v78 - int32(331) {
	case 0, 6:
		v78 = v269
		v80 = v240
		v86 = v262
		v87 = v102
		v88 = v226
		v95 = v228
		goto L14
	default:
		goto L80
	}
L80:
	;
	if v89 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[0])) = int32(0)
	F_read_into_target(m, v21+int32(8), v21+int32(7), l3, l4, l5)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[0])) = int32(2)
	v78 = v96
	v80 = v240
	v86 = v262
	v87 = v102
	v88 = v226
	v89 = int32(1)
	v90 = v272
	v95 = v228
	goto L14
L83:
	;
	F_plpgsql_append_source_text(m, v289, l1, v90, l5)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	v295 = l1
	goto L85
L85:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_plpgsql_append_source_text(m, v289, v295, v296, l5)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	F_appendStringInfoSpaces(m, v289, v102-v90)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v295 = v102
	goto L85
L88:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v299 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v367 = F_palloc0(m, int32(80))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v302 = v299
	goto L91
L91:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v324 = int32(*(*int8)(unsafe.Add(mBase, uint32(v320+v302-int32(1)))))
	goto L93
L92:
	;
	goto L89
L93:
	;
	if base.B2i32(v324 == int32(32))|base.B2i32(base.Ui32((v324-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v338 = v336 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v340+v338))) = uint8(v342)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v342 < v344 {
		v302 = v344
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	v369 = F_pstrdup(m, v365)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v367)+4)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v367))) = v369
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+8)) = v375
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[6]))
	*(*uint8)(unsafe.Add(mBase, uint32(v367)+20)) = uint8(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v367)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v367)+12)) = v378
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_pfree(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_execsql_stmt[7])))
	if v388 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v393 = int32(_a_F_make_execsql_stmt_7)
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[8]))
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[8])) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(_a_F_make_execsql_stmt_8)
	v403 = int32(_a_F_make_execsql_stmt_9)
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[10])) = v21 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(40)
	v413 = F_raw_parser(m, v392, v391)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v424 = F_palloc0(m, int32(24))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[8])) = v394
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[10])) = v418
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(16)
	v428 = int32(0)
	if l1 < v428 {
		v473 = v428
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v473
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_make_execsql_stmt[5]))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+520))
	v480 = v478 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v477)+520)) = v480
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+18)) = uint8(v89)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+12)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = v480
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+19)) = uint8(v485)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+20)) = v487
	m.G0 = v21 + int32(48)
	return v424
L105:
	;
	goto L104
L106:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+60))
	if v434 == int32(0) {
		v473 = v428
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v437 = l1 + v434
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)+188))
	if base.Ui32(v438) <= base.Ui32(v437) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v433)+196))
	if base.B2i32(v447 == int32(0))|base.B2i32(base.Ui32(v437) <= base.Ui32(v447)) != 0 {
		v473 = v448
		goto L105
	} else {
		goto L112
	}
L109:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v433)+192))
	v447 = v440
	goto L108
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+188)) = v434
	v445 = F_strchr(m, v434, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v433)+192)) = v445
	v447 = v445
	goto L108
L112:
	;
	v453 = v447
	v456 = v448
	goto L113
L113:
	;
	v458 = int32(1)
	v459 = v456 + v458
	*(*int32)(unsafe.Add(mBase, uint32(v433)+196)) = v459
	v462 = v453 + v458
	*(*int32)(unsafe.Add(mBase, uint32(v433)+188)) = v462
	v465 = F_strchr(m, v462, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v433)+192)) = v465
	if v465 == int32(0) {
		v473 = v459
		goto L105
	} else {
		goto L115
	}
L114:
	;
	v473 = v459
	goto L105
L115:
	;
	if base.Ui32(v465) < base.Ui32(v437) {
		v453 = v465
		v456 = v459
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(0)
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
	F_errmsg_internal(m, int32(_a_F_make_one_partition_rbound_0), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_make_one_partition_rbound_1), int32(3456), int32(_a_F_make_one_partition_rbound_2))
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
func F_make_scalar_list1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	F_check_assignable(m, l1, l3, l4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = F_palloc0(m, int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(_a_F_make_scalar_list1_0)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
			v21 = F_palloc(m, int32(4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v21
				v25 = F_palloc(m, int32(4))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v25
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = l0
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = v30
					F_plpgsql_adddatum(m, v11)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return v11
					}
				}
			}
		}
	}
}
func F_make_tlist_from_pathtarget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == v2 {
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v2
	v20 = v2
	goto L7
L5:
	;
	v49 = v2
	goto L6
L6:
	;
	return v49
L7:
	;
	v24 = v18 << (uint(int32(2)) % 32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24+v25)))
	v29 = v18 + int32(1)
	v31 = int32(0)
	v33 = F_makeTargetEntry(m, v27, base.I32_extend16_s(v29), v31, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v49 = v41
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24+v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v39
	goto L13
L12:
	;
	goto L13
L13:
	;
	v41 = F_lappend(m, v20, v33)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v29 < v43 {
		v18 = v29
		v20 = v41
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int64
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v15 = F_convert_any_priv_string(m, v9, int32(_a_F_makeaclitem_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v6
				if v13 != 0 {
					v27 = v15 << (uint(int64(32)) % 64)
				} else {
					v27 = int64(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v15&int64(4294967295) | v27
				return v18
			}
		}
	}
}
func F_makepol_1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+159)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = v4
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+14)) = uint16(v4)
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = v4
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v14+int32(159), v14+int32(152), v14+int32(148), v14+int32(14), v14+int32(13))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	m.G0 = v14 + int32(160)
	return
L5:
	;
	goto L4
L6:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v281 == int32(0) {
		v30 = v274
		goto L3
	} else {
		goto L68
	}
L7:
	;
	v263 = v14 + int32(16) + v253<<(uint(int32(2))%32)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v263)+2)) = uint16(v264)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+159)))
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v266)
	v274 = v253 + int32(1)
	goto L6
L8:
	;
	if v63 != int32(32) {
		v253 = v63
		goto L7
	} else {
		goto L64
	}
L9:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L54
	}
L10:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+148))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+14)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	m.T0[l1].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l2, l0, v180, v181, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v156 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L12:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L34
	}
L13:
	;
	F_makepol_1(m, l0, l1, l2)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v30 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v48 {
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
	v253 = int32(0)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+159)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(2))%32))+uint32(_c_F_makepol_1[0])))
	v63 = v30
	goto L19
L19:
	;
	v70 = int32(2)
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63<<(uint(v70)%32)+v14)+12)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(v70)%32))+uint32(_c_F_makepol_1[0])))
	if v53 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v253 = int32(0)
	goto L7
L21:
	;
	v86 = v63 - int32(1)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(16)+v86<<(uint(int32(2))%32))+2)))
	v93 = F_palloc0(m, int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	if v58 <= v78 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v78 <= v58 {
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
	if v73 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v90
	goto L30
L29:
	;
	v98 = int32(0)
	goto L30
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)) = uint16(v98)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)) = uint8(v73)
	v101 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v104 = F_lcons(m, v93, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v104
	if v86 != 0 {
		v63 = v86
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	v274 = v30
	goto L6
L34:
	;
	v115 = v30
	goto L35
L35:
	;
	v122 = int32(2)
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115<<(uint(v122)%32)+v14)+12)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125<<(uint(v122)%32))+uint32(_c_F_makepol_1[0])))
	if v130 <= int32(0) {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	goto L5
L37:
	;
	v136 = v115 - int32(1)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(16)+v136<<(uint(int32(2))%32))+2)))
	v142 = F_palloc0(m, int32(8))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v125 == int32(4) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v147 = v140
	goto L41
L40:
	;
	v147 = int32(0)
	goto L41
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v142)+2)) = uint16(v147)
	*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)) = uint8(v125)
	v150 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v153 = F_lcons(m, v142, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v153
	if v136 != 0 {
		v115 = v136
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	v163 = F_errsave_start(m, v156)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v159 != int32(447) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
	if v162 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	if v163 == int32(0) {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v170
	F_errmsg(m, int32(_a_F_makepol_1_0), v14)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v156, int32(_a_F_makepol_1_1), int32(714), int32(_a_F_makepol_1_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L5
L53:
	;
	v274 = v30
	goto L6
L54:
	;
	v192 = v30
	goto L55
L55:
	;
	v199 = int32(2)
	v202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v192<<(uint(v199)%32)+v14)+12)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202<<(uint(v199)%32))+uint32(_c_F_makepol_1[0])))
	if v207 <= int32(0) {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	goto L5
L57:
	;
	v213 = v192 - int32(1)
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(16)+v213<<(uint(int32(2))%32))+2)))
	v219 = F_palloc0(m, int32(8))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v202 == int32(4) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v224 = v217
	goto L61
L60:
	;
	v224 = int32(0)
	goto L61
L61:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+2)) = uint16(v224)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)) = uint8(v202)
	v227 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v230 = F_lcons(m, v219, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v230
	if v213 != 0 {
		v192 = v213
		goto L55
	} else {
		goto L63
	}
L63:
	;
	goto L56
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_makepol_1_3), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_makepol_1_1), int32(639), int32(_a_F_makepol_1_4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v284 != int32(447) {
		v30 = v274
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+4)))
	if v287 == int32(0) {
		v30 = v274
		goto L3
	} else {
		goto L70
	}
L70:
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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_manifest_process_system_identifier[0]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v12 != l1 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = l1
		m.T0[v14].(func(*base.Module, int32, int32, int32))(m, l0, int32(_a_F_manifest_process_system_identifier_0), v8)
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
	v6 = v4 & int32(_a_F_mask_page_hint_bits_0)
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
	var v94 int32
	_ = v94
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
						v94 = v5
						m.G0 = v13 + int32(16)
						return v94
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
									v94 = v63
									m.G0 = v13 + int32(16)
									return v94
								}
							}
						}
					}
				}
			}
		} else {
			if v16 == int32(0) {
				v94 = v5
				m.G0 = v13 + int32(16)
				return v94
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if v23 != int32(53) {
					if v23 != int32(21) {
						v94 = v5
						m.G0 = v13 + int32(16)
						return v94
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
						if v28 != int32(2) {
							v94 = v5
							m.G0 = v13 + int32(16)
							return v94
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
												v94 = v5
												m.G0 = v13 + int32(16)
												return v94
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
															v94 = v63
															m.G0 = v13 + int32(16)
															return v94
														}
													}
												}
											}
										}
									}
								} else {
									v94 = v5
									m.G0 = v13 + int32(16)
									return v94
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
							v94 = v5
							m.G0 = v13 + int32(16)
							return v94
						} else {
							v46 = F_match_index_to_operand(m, v37, v3, l3)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v46 == int32(0) {
									v94 = v5
									m.G0 = v13 + int32(16)
									return v94
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
												v94 = v5
												m.G0 = v13 + int32(16)
												return v94
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
															v94 = v63
															m.G0 = v13 + int32(16)
															return v94
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
											v94 = v5
											m.G0 = v13 + int32(16)
											return v94
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
														v94 = v63
														m.G0 = v13 + int32(16)
														return v94
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
									v94 = v5
									m.G0 = v13 + int32(16)
									return v94
								} else {
									v46 = F_match_index_to_operand(m, v37, v3, l3)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										if v46 == int32(0) {
											v94 = v5
											m.G0 = v13 + int32(16)
											return v94
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
														v94 = v5
														m.G0 = v13 + int32(16)
														return v94
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
																	v94 = v63
																	m.G0 = v13 + int32(16)
																	return v94
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
	var v27 int32
	_ = v27
	var v31 float32
	_ = v31
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v106 float64
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v245 float32
	_ = v245
	var v255 int32
	_ = v255
	var v257 float64
	_ = v257
	var v270 float64
	_ = v270
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v287 int32
	_ = v287
	var v290 float64
	_ = v290
	var v302 int32
	_ = v302
	var v313 float64
	_ = v313
	v12 = int32(0)
	v27 = base.B2i32(l3 == l1+int32(3)) & base.B2i32(l2 != v12)
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*float32)(unsafe.Add(mBase, uint32(l2+l1<<(uint(int32(2))%32))))
	v35 = base.F64_promote_f32(base.F32_mul(v31, float32(0.5)))
	goto L3
L2:
	;
	v35 = float64(0.004999999888241291)
	goto L3
L3:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = base.B2i32(base.Ui32(int32(_a_F_mcelem_array_contain_overlap_selec_0)) < base.Ui32(l1))
	if base.Ui32(int32(_a_F_mcelem_array_contain_overlap_selec_0)) < base.Ui32(l1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v75 = int32(-1)
	goto L6
L6:
	;
	if l6 == int32(2751) {
		goto L31
	} else {
		goto L32
	}
L7:
	;
	v40 = int32(16)
	goto L9
L8:
	;
	v40 = int32(0)
	goto L9
L9:
	;
	if base.Ui32(int32(_a_F_mcelem_array_contain_overlap_selec_0)) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v45 = int32(base.Ui32(l1) >> (uint(int32(16)) % 32))
	goto L12
L11:
	;
	v45 = l1
	goto L12
L12:
	;
	v47 = base.B2i32(base.Ui32(int32(255)) < base.Ui32(v45))
	if base.Ui32(int32(255)) < base.Ui32(v45) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = v40 | int32(8)
	goto L15
L14:
	;
	v48 = v40
	goto L15
L15:
	;
	if base.Ui32(int32(255)) < base.Ui32(v45) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = int32(base.Ui32(v45) >> (uint(int32(8)) % 32))
	goto L18
L17:
	;
	v53 = v45
	goto L18
L18:
	;
	v55 = base.B2i32(base.Ui32(int32(15)) < base.Ui32(v53))
	if base.Ui32(int32(15)) < base.Ui32(v53) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v56 = v48 | int32(4)
	goto L21
L20:
	;
	v56 = v48
	goto L21
L21:
	;
	if base.Ui32(int32(15)) < base.Ui32(v53) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(base.Ui32(v53) >> (uint(int32(4)) % 32))
	goto L24
L23:
	;
	v61 = v53
	goto L24
L24:
	;
	v63 = base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61))
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v64 = v56 | int32(2)
	goto L27
L26:
	;
	v64 = v56
	goto L27
L27:
	;
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v67 = int32(base.Ui32(v61) >> (uint(int32(2)) % 32))
	goto L30
L29:
	;
	v67 = v61
	goto L30
L30:
	;
	v75 = v64 + base.B2i32(base.Ui32(int32(1)) < base.Ui32(v67))
	goto L6
L31:
	;
	v80 = float64(1)
	goto L33
L32:
	;
	v80 = float64(0)
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
	v83 = float64(0.005)
	if base.F64_gt(v35, v83) != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v313 = v80
	goto L36
L36:
	;
	return v313
L37:
	;
	v86 = v83
	goto L39
L38:
	;
	v86 = v35
	goto L39
L39:
	;
	v88 = l1 - int32(1)
	v90 = l7 + int32(104)
	v103 = int32(0)
	v106 = v80
	v110 = v12
	goto L40
L40:
	;
	if v110 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v313 = v290
	goto L36
L42:
	;
	v302 = v110 + int32(1)
	if v302 != l5 {
		v103 = v287
		v106 = v290
		v110 = v302
		goto L40
	} else {
		goto L87
	}
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v120 = l4 + v110<<(uint(int32(2))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120-int32(4))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v125 = F_FunctionCall2Coll(m, v90, v117, v123, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if base.B2i32(v75*l5 < l1+l5) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	return float64(0)
L47:
	;
	if v125 == int32(0) {
		v287 = v103
		v290 = v106
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v270 = base.F64_mul(v106, v257)
	if l6 != int32(2751) {
		goto L80
	} else {
		goto L81
	}
L50:
	;
	v245 = *(*float32)(unsafe.Add(mBase, uint32(l2+v228<<(uint(int32(2))%32))))
	v255 = v228 + int32(1)
	v257 = base.F64_promote_f32(v245)
	goto L49
L51:
	;
	if v27 == int32(0) {
		v255 = v163
		v257 = v86
		goto L49
	} else {
		goto L79
	}
L52:
	;
	v190 = v103
	goto L71
L53:
	;
	if v103 < l1 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v88 < v103 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v255 = v103
	v257 = v86
	goto L49
L57:
	;
	v255 = v103
	v257 = v86
	goto L49
L58:
	;
	goto L59
L59:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l4+v110<<(uint(int32(2))%32))))
	v143 = v103
	v152 = v88
	goto L60
L60:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v162 = int32(2)
	v163 = base.I32_div_s(v143+v152, v162)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0+v163<<(uint(v162)%32))))
	v168 = F_FunctionCall2Coll(m, v90, v160, v167, v139)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L46
	} else {
		goto L62
	}
L61:
	;
	v255 = v176
	v257 = v86
	goto L49
L62:
	;
	if v168 == int32(0) {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v175 = base.B2i32(v168 < int32(0))
	if v168 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v176 = v163 + int32(1)
	goto L66
L65:
	;
	v176 = v143
	goto L66
L66:
	;
	if v168 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v179 = v152
	goto L69
L68:
	;
	v179 = v163 - int32(1)
	goto L69
L69:
	;
	if v176 <= v179 {
		v143 = v176
		v152 = v179
		goto L60
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0+v190<<(uint(int32(2))%32))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l4+v110<<(uint(int32(2))%32))))
	v210 = F_FunctionCall2Coll(m, v90, v204, v208, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L46
	} else {
		goto L73
	}
L72:
	;
	v255 = l1
	v257 = v86
	goto L49
L73:
	;
	if int32(0) <= v210 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	if v27&base.B2i32(v210 == int32(0)) != 0 {
		v228 = v190
		goto L50
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v218 = v190 + int32(1)
	if v218 != l1 {
		v190 = v218
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v255 = v190
	v257 = v86
	goto L49
L78:
	;
	goto L72
L79:
	;
	v228 = v163
	goto L50
L80:
	;
	v272 = base.F64_sub(base.F64_add(v106, v257), v270)
	goto L82
L81:
	;
	v272 = v270
	goto L82
L82:
	;
	v273 = float64(0)
	if base.F64_lt(v272, v273) != 0 {
		v287 = v255
		v290 = v273
		goto L42
	} else {
		goto L83
	}
L83:
	;
	if base.F64_gt(v272, float64(1)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v287 = v255
	v290 = v272
	goto L42
L85:
	;
	goto L86
L86:
	;
	v287 = v255
	v290 = float64(1)
	goto L42
L87:
	;
	goto L41
}
func F_md_readv_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v12 = v9 + int32(56)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_md_readv_report[0]))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	if v19&int32(256) != 0 {
		v22 = v17
	} else {
		v22 = int32(-1)
	}
	F_GetRelationPath(m, v12, v13, v14, v15, v22, base.I32_extend8_s(v19))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = int32(base.Ui32(v27) >> (uint(int32(9)) % 32))
		if v29 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_md_readv_report[1])) = v29
			v33 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				if v33 == int32(0) {
					m.G0 = v9 + int32(128)
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v39 + v40 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v12
						F_errmsg(m, int32(_a_F_md_readv_report_0), v9+int32(32))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v85 = int32(2064)
							F_errfinish(m, int32(_a_F_md_readv_report_1), v85, int32(_a_F_md_readv_report_2))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				}
			}
		} else {
			v54 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				if v54 == int32(0) {
					m.G0 = v9 + int32(128)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v64 = int32(13)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v63 << (uint(v64) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v61 << (uint(v64) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v62 + v63 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(56)
						F_errmsg(m, int32(_a_F_md_readv_report_3), v9)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v85 = int32(2079)
							F_errfinish(m, int32(_a_F_md_readv_report_1), v85, int32(_a_F_md_readv_report_2))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
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
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = int32(2)
	v8 = l0 << (uint(v7) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l2+v8-int32(4)))) = int32(1)
	v15 = l0 - v7
	if v15 < int32(0) {
	} else {
		if l0&int32(1) == int32(0) {
			v26 = v8 - int32(4)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1+v26)))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l2+v26)))
			*(*int32)(unsafe.Add(mBase, uint32(l2+v15<<(uint(int32(2))%32)))) = v28 * v30
			v35 = l0 - int32(3)
		} else {
			v35 = v15
		}
		if v15 == int32(0) {
		} else {
			v41 = v35
			for {
				v44 = int32(2)
				v45 = v41 << (uint(v44) % 32)
				v48 = v45 + int32(4)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1+v48)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+l2)))
				v53 = v50 * v52
				*(*int32)(unsafe.Add(mBase, uint32(l2+v45))) = v53
				v56 = v41 - int32(1)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1+v45)))
				*(*int32)(unsafe.Add(mBase, uint32(l2+v56<<(uint(v44)%32)))) = v61 * v53
				if v56 != 0 {
					v41 = v41 - v44
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
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(_a_F_mdc_flush_0)
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
				base.MemoryFill(m, v6, int32(0), int32(22))
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
						F_px_debug(m, int32(_a_F_mdc_read_0), int32(0))
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
		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(_a_F_message_cb_wrapper_0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(993)
		v23 = int32(_a_F_message_cb_wrapper_1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_message_cb_wrapper[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_message_cb_wrapper[0])) = v13 + int32(4)
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
		*(*int32)(unsafe.Add(mBase, _c_F_message_cb_wrapper[0])) = v44
	} else {
	}
	m.G0 = v13 + int32(32)
	return
}
func F_minimal_tuple_from_heap_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 - int32(8)
	v9 = F_palloc(m, v7+l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			base.MemoryFill(m, v9, int32(0), l1)
		} else {
		}
		v15 = l1 + v9
		if v7 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			base.MemoryCopy(m, v15, v16+int32(8), v7)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v7
		return v15
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != int32(2) {
		v64 = l1
		v66 = l3
		v67 = v10
		v72 = l2
		if int32(0) < v72 {
			v97 = v64
			v99 = v66
			v100 = v67
			for {
				v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(52))))))
				v109 = v106 ^ v108
				*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(84))))) = uint8(v109)
				*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
				v112 = int32(1)
				v117 = v100 + v112
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v119 = v118 + v72
				if v117 < v119 {
					v97 = v97 + v112
					v99 = v99 + v112
					v100 = v117
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
			return v72
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
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
			v79 = l0 + int32(20)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v82 = v80 - int32(2)
			if v82 != 0 {
				base.MemoryCopy(m, v79, l0+int32(86), v82)
			} else {
			}
			v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+84)))
			*(*uint16)(unsafe.Add(mBase, uint32(v82+v79))) = uint16(v87)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			return v17
		} else {
			v64 = v51
			v66 = v53
			v67 = v54
			v72 = l2 - v17
			if int32(0) < v72 {
				v97 = v64
				v99 = v66
				v100 = v67
				for {
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(52))))))
					v109 = v106 ^ v108
					*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(84))))) = uint8(v109)
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = int32(1)
					v117 = v100 + v112
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v119 = v118 + v72
					if v117 < v119 {
						v97 = v97 + v112
						v99 = v99 + v112
						v100 = v117
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
				return v72
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
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
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
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
	v30 = l1
	v31 = v6
	v32 = v6
	goto L4
L4:
	;
	v43 = v23 + v30*int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = int32(base.Ui32(v44)>>(uint(int32(10))%32)) & int32(_a_F_mkANode_0)
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
	v59 = v31
	v60 = v32
	goto L8
L8:
	;
	v62 = v30 + int32(1)
	if v62 != l2 {
		v30 = v62
		v31 = v59
		v32 = v60
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
	v59 = v31 + base.B2i32(v32&int32(255) != v56)
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
	v125 = int32(0)
	v126 = l1
	v128 = v105 + int32(4)
	v135 = v6
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
	v93 = int32(_a_F_mkANode_1)
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
	v139 = v126 * int32(24)
	v140 = v137 + v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v145 = int32(base.Ui32(v141)>>(uint(int32(10))%32)) & int32(_a_F_mkANode_0)
	if v145 <= l3 {
		v249 = v121
		v250 = v125
		v252 = v128
		v254 = v135
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v258 = F_mkANode(m, l0, v249, l2, v114, l4)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L63
	}
L30:
	;
	v256 = v126 + int32(1)
	if v256 != l2 {
		v121 = v249
		v125 = v250
		v126 = v256
		v128 = v252
		v135 = v254
		goto L28
	} else {
		goto L62
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
	v149 = v116 + v145
	goto L34
L33:
	;
	v149 = l3
	goto L34
L34:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v149))))
	if v151 != v135 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v135 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v218 = v121
	v219 = v125
	v220 = v147
	v221 = v128
	v222 = v137
	v223 = v135
	goto L37
L37:
	;
	if l4 != 0 {
		goto L58
	} else {
		goto L59
	}
L38:
	;
	v153 = F_mkANode(m, l0, v121, v126, v114, l4)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L41
	}
L39:
	;
	v203 = v121
	v204 = v125
	v205 = v147
	v206 = v128
	v207 = v137
	goto L40
L40:
	;
	if l4 != 0 {
		goto L55
	} else {
		goto L56
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v153
	if v125 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199+v139)+12))
	v203 = v126
	v204 = int32(0)
	v205 = v201
	v206 = v128 + int32(12)
	v207 = v199
	goto L40
L43:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v158 | v125<<(uint(int32(8))%32)
	v164 = v125 << (uint(int32(2)) % 32)
	if base.Ui32(int32(1025)) <= base.Ui32(v164) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v187
	if v164 == int32(0) {
		goto L42
	} else {
		goto L54
	}
L45:
	;
	v167 = F_palloc0(m, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v172 = (v164 + int32(7)) & int32(4088)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v172) <= base.Ui32(v173) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v187 = v167
	goto L44
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v180 - v172
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v181 + v172
	v187 = v181
	goto L44
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v180 = v173
	v181 = v175
	goto L49
L51:
	;
	goto L52
L52:
	;
	v176 = int32(_a_F_mkANode_1)
	v178 = F_palloc0(m, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v180 = v176
	v181 = v178
	goto L49
L54:
	;
	base.MemoryCopy(m, v187, v74, v164)
	goto L42
L55:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v139+v207)+4))
	v215 = int32(base.Ui32(v209)>>(uint(int32(10))%32))&int32(_a_F_mkANode_0) + v116
	goto L57
L56:
	;
	v215 = l3
	goto L57
L57:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+v205))))
	v218 = v203
	v219 = v204
	v220 = v205
	v221 = v206
	v222 = v207
	v223 = v217
	goto L37
L58:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v139+v222)+4))
	v231 = int32(base.Ui32(v225)>>(uint(int32(10))%32))&int32(_a_F_mkANode_0) + v116
	goto L60
L59:
	;
	v231 = l3
	goto L60
L60:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v220))))
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v236 = v235 + v139
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if int32(base.Ui32(v237)>>(uint(int32(10))%32))&int32(_a_F_mkANode_0) != v114 {
		v249 = v218
		v250 = v219
		v252 = v221
		v254 = v223
		goto L30
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74+v219<<(uint(int32(2))%32)))) = v236
	v249 = v218
	v250 = v219 + int32(1)
	v252 = v221
	v254 = v223
	goto L30
L62:
	;
	goto L29
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v258
	if v250 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v74)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L16
	} else {
		goto L77
	}
L65:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v263 | v250<<(uint(int32(8))%32)
	v269 = v250 << (uint(int32(2)) % 32)
	if base.Ui32(int32(1025)) <= base.Ui32(v269) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = v293
	if v269 == int32(0) {
		goto L64
	} else {
		goto L76
	}
L67:
	;
	v272 = F_palloc0(m, v269)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v277 = (v269 + int32(7)) & int32(4088)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v277) <= base.Ui32(v278) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v293 = v272
	goto L66
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v285 - v277
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v277 + v286
	v293 = v286
	goto L66
L72:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v285 = v278
	v286 = v280
	goto L71
L73:
	;
	goto L74
L74:
	;
	v281 = int32(_a_F_mkANode_1)
	v283 = F_palloc0(m, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	v285 = v281
	v286 = v283
	goto L71
L76:
	;
	base.MemoryCopy(m, v293, v74, v269)
	goto L64
L77:
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
		*(*int32)(unsafe.Add(mBase, _c_F_mkdir[0])) = int32(0) - v4
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
	F_px_debug(m, int32(_a_F_mpi_check_0), v10)
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
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[0]))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mq_putmessage[1])))
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
	*(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[0])) = int32(0)
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
	*(*uint8)(unsafe.Add(mBase, _c_F_mq_putmessage[1])) = uint8(v26)
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
	*(*uint8)(unsafe.Add(mBase, _c_F_mq_putmessage[1])) = uint8(v90)
	if v43 != 0 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[2]))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[3]))
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
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[2]))
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
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[4]))
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
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[5]))
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
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(0)
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[6]))
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
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_mq_putmessage[0]))
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
