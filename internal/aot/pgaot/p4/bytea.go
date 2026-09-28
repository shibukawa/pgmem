package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaGetBit(m *base.Module, l0 int32) int64 {
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
	var v16 int64
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
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
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
		v51 = base.I64_extend_i32_s(v48) << (uint(int64(3)) % 64)
		if base.B2i32(int64(0) <= v16)&base.B2i32(v16 < v51) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v51 - int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v16
					F_errmsg(m, int32(_a_F_byteaGetBit_0), v9)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaGetBit_1), int32(683), int32(_a_F_byteaGetBit_2))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
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
			v78 = int32(1)
			if v19&v78 != 0 {
				v82 = v78
			} else {
				v82 = int32(4)
			}
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v16)>>(uint(int64(3))%64)))+(v12+v82)))))
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(int32(base.Ui32(v85)>>(uint(base.I32_wrap_i64(v16)&int32(7))%32)) & int32(1))
		}
	}
}
func F_bytea_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int64
	_ = v472
	var v474 int64
	_ = v474
	var v476 int64
	_ = v476
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v485 int64
	_ = v485
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
		v22 = int32(1)
		v25 = v19 & v22
		if v25 != 0 {
			v26 = v22
		} else {
			v26 = int32(4)
		}
		v27 = v26 + v15
		if v19 == int32(1) {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v34 == int32(18) {
				v37 = int32(16)
			} else {
				v37 = int32(0)
			}
			if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v44 = int32(4)
			} else {
				v44 = v37
			}
			v55 = v44
		} else {
			v45 = int32(1)
			if v25 != 0 {
				v55 = int32(base.Ui32(v19)>>(uint(v45)%32)) - v45
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(8)) <= base.Ui32(v55) {
			v58 = int32(8)
		} else {
			v58 = v55
		}
		if v58 != 0 {
			base.MemoryCopy(m, v11+int32(8), v27, v58)
		} else {
		}
		v62 = int32(128)
		if v62 <= v55 {
			v65 = v62
		} else {
			v65 = v55
		}
		v71 = v65 - int32(1636608432)
		if v27&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v65) {
				v180 = v27
				v181 = v65
				v182 = v71
				v183 = v71
				v184 = v71
				for {
					v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
					v187 = v186 + v183
					v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
					v191 = v190 + v184
					v193 = int32(4)
					v195 = v188 + v182 - v191 ^ base.I32_rotl(v191, v193)
					v199 = v187 - v195 ^ base.I32_rotl(v195, int32(6))
					v200 = v191 + v187
					v201 = v195 + v200
					v202 = v199 + v201
					v206 = v200 - v199 ^ base.I32_rotl(v199, int32(8))
					v210 = v201 - v206 ^ base.I32_rotl(v206, int32(16))
					v214 = v202 - v210 ^ base.I32_rotl(v210, int32(19))
					v215 = v206 + v202
					v216 = v210 + v215
					v217 = v214 + v216
					v221 = v215 - v214 ^ base.I32_rotl(v214, v193)
					v222 = int32(12)
					v223 = v180 + v222
					v225 = v181 - v222
					if base.Ui32(int32(11)) < base.Ui32(v225) {
						v180 = v223
						v181 = v225
						v182 = v216
						v183 = v217
						v184 = v221
						continue
					} else {
						break
					}
					break
				}
				v228 = v223
				v229 = v225
				v230 = v216
				v231 = v217
				v232 = v221
			} else {
				v228 = v27
				v229 = v65
				v230 = v71
				v231 = v71
				v232 = v71
			}
			switch v229 - int32(1) {
			case 0:
				v291 = v230
				v292 = v231
				v293 = v232
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 1:
				v284 = v230
				v285 = v231
				v286 = v232
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 2:
				v277 = v230
				v278 = v231
				v279 = v232
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 3:
				v271 = v231
				v272 = v232
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 4:
				v267 = v231
				v268 = v232
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 5:
				v261 = v231
				v262 = v232
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 6:
				v255 = v231
				v256 = v232
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+6)))
				v261 = v257<<(uint(int32(16))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 7:
				v250 = v232
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+7)))
				v255 = v251<<(uint(int32(24))%32) + v231
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+6)))
				v261 = v257<<(uint(int32(16))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 8:
				v245 = v232
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+8)))
				v250 = v246<<(uint(int32(8))%32) + v245
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+7)))
				v255 = v251<<(uint(int32(24))%32) + v231
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+6)))
				v261 = v257<<(uint(int32(16))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 9:
				v240 = v232
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+9)))
				v245 = v241<<(uint(int32(16))%32) + v240
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+8)))
				v250 = v246<<(uint(int32(8))%32) + v245
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+7)))
				v255 = v251<<(uint(int32(24))%32) + v231
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+6)))
				v261 = v257<<(uint(int32(16))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			case 10:
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+10)))
				v240 = v236<<(uint(int32(24))%32) + v232
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+9)))
				v245 = v241<<(uint(int32(16))%32) + v240
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+8)))
				v250 = v246<<(uint(int32(8))%32) + v245
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+7)))
				v255 = v251<<(uint(int32(24))%32) + v231
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+6)))
				v261 = v257<<(uint(int32(16))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
				v267 = v263<<(uint(int32(8))%32) + v261
				v268 = v262
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
				v271 = v267 + v269
				v272 = v268
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+3)))
				v277 = v273<<(uint(int32(24))%32) + v230
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
				v284 = v280<<(uint(int32(16))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
				v291 = v287<<(uint(int32(8))%32) + v284
				v292 = v285
				v293 = v286
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				v298 = v291 + v294
				v299 = v292
				v300 = v293
			default:
				v298 = v230
				v299 = v231
				v300 = v232
			}
		} else {
			if base.Ui32(v65) < base.Ui32(int32(12)) {
				v126 = v27
				v127 = v65
				v128 = v71
				v129 = v71
				v130 = v71
			} else {
				v78 = v27
				v79 = v65
				v80 = v71
				v81 = v71
				v82 = v71
				for {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
					v85 = v84 + v81
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
					v89 = v88 + v82
					v91 = int32(4)
					v93 = v86 + v80 - v89 ^ base.I32_rotl(v89, v91)
					v97 = v85 - v93 ^ base.I32_rotl(v93, int32(6))
					v98 = v89 + v85
					v99 = v93 + v98
					v100 = v97 + v99
					v104 = v98 - v97 ^ base.I32_rotl(v97, int32(8))
					v108 = v99 - v104 ^ base.I32_rotl(v104, int32(16))
					v112 = v100 - v108 ^ base.I32_rotl(v108, int32(19))
					v113 = v104 + v100
					v114 = v108 + v113
					v115 = v112 + v114
					v119 = v113 - v112 ^ base.I32_rotl(v112, v91)
					v120 = int32(12)
					v121 = v78 + v120
					v123 = v79 - v120
					if base.Ui32(int32(11)) < base.Ui32(v123) {
						v78 = v121
						v79 = v123
						v80 = v114
						v81 = v115
						v82 = v119
						continue
					} else {
						break
					}
					break
				}
				v126 = v121
				v127 = v123
				v128 = v114
				v129 = v115
				v130 = v119
			}
			switch v127 - int32(1) {
			case 0:
				v177 = v128
				v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
				v298 = v177 + v178
				v299 = v129
				v300 = v130
			case 1:
				v172 = v128
				v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
				v177 = v173<<(uint(int32(8))%32) + v172
				v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
				v298 = v177 + v178
				v299 = v129
				v300 = v130
			case 2:
				v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
				v172 = v168<<(uint(int32(16))%32) + v128
				v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
				v177 = v173<<(uint(int32(8))%32) + v172
				v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
				v298 = v177 + v178
				v299 = v129
				v300 = v130
			case 3:
				v165 = v129
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v298 = v166 + v128
				v299 = v165
				v300 = v130
			case 4:
				v162 = v129
				v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+4)))
				v165 = v162 + v163
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v298 = v166 + v128
				v299 = v165
				v300 = v130
			case 5:
				v157 = v129
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+5)))
				v162 = v158<<(uint(int32(8))%32) + v157
				v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+4)))
				v165 = v162 + v163
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v298 = v166 + v128
				v299 = v165
				v300 = v130
			case 6:
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+6)))
				v157 = v153<<(uint(int32(16))%32) + v129
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+5)))
				v162 = v158<<(uint(int32(8))%32) + v157
				v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+4)))
				v165 = v162 + v163
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v298 = v166 + v128
				v299 = v165
				v300 = v130
			case 7:
				v148 = v130
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
				v298 = v149 + v128
				v299 = v151 + v129
				v300 = v148
			case 8:
				v143 = v130
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+8)))
				v148 = v144<<(uint(int32(8))%32) + v143
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
				v298 = v149 + v128
				v299 = v151 + v129
				v300 = v148
			case 9:
				v138 = v130
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+9)))
				v143 = v139<<(uint(int32(16))%32) + v138
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+8)))
				v148 = v144<<(uint(int32(8))%32) + v143
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
				v298 = v149 + v128
				v299 = v151 + v129
				v300 = v148
			case 10:
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+10)))
				v138 = v134<<(uint(int32(24))%32) + v130
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+9)))
				v143 = v139<<(uint(int32(16))%32) + v138
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+8)))
				v148 = v144<<(uint(int32(8))%32) + v143
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
				v298 = v149 + v128
				v299 = v151 + v129
				v300 = v148
			default:
				v298 = v128
				v299 = v129
				v300 = v130
			}
		}
		v303 = int32(14)
		v305 = v299 ^ v300 - base.I32_rotl(v299, v303)
		v309 = v305 ^ v298 - base.I32_rotl(v305, int32(11))
		v313 = v309 ^ v299 - base.I32_rotl(v309, int32(25))
		v317 = v313 ^ v305 - base.I32_rotl(v313, int32(16))
		v321 = v317 ^ v309 - base.I32_rotl(v317, int32(4))
		v325 = v321 ^ v313 - base.I32_rotl(v321, v303)
		v329 = v325 ^ v317 - base.I32_rotl(v325, int32(24))
		v331 = v13 + int32(24)
		if int32(129) <= v55 {
			v338 = int32(711645284)
			v341 = v55 - int32(1636608428) ^ v338 - int32(1455628627)
			v346 = v341 ^ int32(-1636608428) - base.I32_rotl(v341, int32(25))
			v351 = v346 ^ v338 - base.I32_rotl(v346, int32(16))
			v355 = v351 ^ v341 - base.I32_rotl(v351, int32(4))
			v359 = v355 ^ v346 - base.I32_rotl(v355, int32(14))
			v365 = v359 ^ v351 - base.I32_rotl(v359, int32(24)) ^ v329
		} else {
			v365 = v329
		}
		v368 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
		v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
		v371 = int32(32) - v370
		v373 = v368 + int32(base.Ui32(v365)>>(uint(v371)%32))
		v374 = v365 << (uint(v370) % 32)
		if v374 != 0 {
			v381 = int32(32) - (base.I32_clz(v374) ^ int32(31))
			v382 = int32(255)
			if base.Ui32(v371&v382) < base.Ui32(v381&v382) {
				v387 = v371 + int32(1)
			} else {
				v387 = v381
			}
			v391 = v387
		} else {
			v391 = v371 + int32(1)
		}
		v393 = v391 & int32(255)
		v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
		if base.Ui32(v394) < base.Ui32(v393) {
			v396 = v393
		} else {
			v396 = v394
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v373))) = uint8(v396)
		v398 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		v407 = int32(711645284)
		v410 = base.I32_wrap_i64(int64(base.Ui64(v398)>>(uint(int64(32))%64))^v398) - int32(1636608428) ^ v407 - int32(1455628627)
		v415 = v410 ^ int32(-1636608428) - base.I32_rotl(v410, int32(25))
		v420 = v415 ^ v407 - base.I32_rotl(v415, int32(16))
		v424 = v420 ^ v410 - base.I32_rotl(v420, int32(4))
		v428 = v424 ^ v415 - base.I32_rotl(v424, int32(14))
		v432 = v428 ^ v420 - base.I32_rotl(v428, int32(24))
		v435 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v438 = int32(32) - v437
		v440 = v435 + int32(base.Ui32(v432)>>(uint(v438)%32))
		v441 = v432 << (uint(v437) % 32)
		if v441 != 0 {
			v448 = int32(32) - (base.I32_clz(v441) ^ int32(31))
			v449 = int32(255)
			if base.Ui32(v438&v449) < base.Ui32(v448&v449) {
				v454 = v438 + int32(1)
			} else {
				v454 = v448
			}
			v458 = v454
		} else {
			v458 = v438 + int32(1)
		}
		v460 = v458 & int32(255)
		v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
		if base.Ui32(v461) < base.Ui32(v460) {
			v463 = v460
		} else {
			v463 = v461
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v463)
		if base.I64_extend_i32_u(v15) != l0 {
			F_pfree(m, v15)
			mBase = m.M
			v468 = m.ExcPending
			if v468 != 0 {
				return int64(0)
			} else {
				m.G0 = v11 + int32(16)
				v472 = int64(56)
				v474 = int64(65280)
				v476 = int64(40)
				v479 = int64(16711680)
				v481 = int64(24)
				v483 = int64(4278190080)
				v485 = int64(8)
				return v398<<(uint(v472)%64) | v398&v474<<(uint(v476)%64) | (v398&v479<<(uint(v481)%64) | v398&v483<<(uint(v485)%64)) | (int64(base.Ui64(v398)>>(uint(v485)%64))&v483 | int64(base.Ui64(v398)>>(uint(v481)%64))&v479 | (int64(base.Ui64(v398)>>(uint(v476)%64))&v474 | int64(base.Ui64(v398)>>(uint(v472)%64))))
			}
		} else {
			m.G0 = v11 + int32(16)
			v472 = int64(56)
			v474 = int64(65280)
			v476 = int64(40)
			v479 = int64(16711680)
			v481 = int64(24)
			v483 = int64(4278190080)
			v485 = int64(8)
			return v398<<(uint(v472)%64) | v398&v474<<(uint(v476)%64) | (v398&v479<<(uint(v481)%64) | v398&v483<<(uint(v485)%64)) | (int64(base.Ui64(v398)>>(uint(v485)%64))&v483 | int64(base.Ui64(v398)>>(uint(v481)%64))&v479 | (int64(base.Ui64(v398)>>(uint(v476)%64))&v474 | int64(base.Ui64(v398)>>(uint(v472)%64))))
		}
	}
}
func F_bytea_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	if int32(0) < l2 {
		v9 = l2 + l3
		if base.B2i32(l3 < int32(0)) != base.B2i32(v9 < l2) {
			F_errstart_cold(m, int32(21), int32(0))
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bytea_overlay_0), int32(0))
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bytea_overlay_1), int32(177), int32(_a_F_bytea_overlay_2))
						v65 = m.ExcPending
						if v65 != 0 {
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
			v13 = int32(1)
			v16 = F_bytea_substring(m, base.I64_extend_i32_u(l0), v13, l2-v13)
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(1)
				if v9 <= v20 {
					v23 = v20
				} else {
					v23 = v9
				}
				v27 = F_detoast_attr_slice(m, l0, v23-int32(1), int32(-1))
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_bytea_catenate(m, v16, l1)
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = F_bytea_catenate(m, v29, v27)
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							return v31
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_bytea_overlay_3), int32(0))
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bytea_overlay_1), int32(173), int32(_a_F_bytea_overlay_2))
					v49 = m.ExcPending
					if v49 != 0 {
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
func F_bytea_reverse(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v12 = int32(1)
		v13 = v11 & v12
		if v11 == v12 {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if v19 == int32(18) {
				v22 = int32(16)
			} else {
				v22 = int32(0)
			}
			if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v29 = int32(4)
			} else {
				v29 = v22
			}
			v40 = v29
		} else {
			v30 = int32(1)
			if v13 != 0 {
				v40 = int32(base.Ui32(v11)>>(uint(v30)%32)) - v30
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v42 = v40 + int32(4)
		v43 = F_palloc(m, v42)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v43))) = v42 << (uint(int32(2)) % 32)
			if int32(0) < v40 {
				if v13 != 0 {
					v52 = int32(1)
				} else {
					v52 = int32(4)
				}
				v53 = v7 + v52
				v58 = v53
				v59 = v40 + v43 + int32(4)
				for {
					v63 = int32(1)
					v64 = v59 - v63
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
					*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v65)
					v68 = v58 + v63
					if base.Ui32(v68) < base.Ui32(v53+v40) {
						v58 = v68
						v59 = v64
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return base.I64_extend_i32_u(v43)
		}
	}
}
func F_bytea_string_agg_transfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v18 = v17
	} else {
		v18 = v2
	}
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v19 != 0 {
		v185 = v18
		if v185 != 0 {
			v193 = base.I64_extend_i32_u(v185)
		} else {
			v190 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v190)
			v193 = int64(0)
		}
		m.G0 = v12 + int32(16)
		return v193
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v18 == int32(0) {
				v28 = v12 + int32(12)
				v29 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v30 == v29 {
					v47 = int32(0)
					if v28 == v47 {
						v55 = v47
					} else {
						v50 = v47
						v51 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
						v55 = v51
					}
					v58 = v55
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					switch v33 - int32(435) {
					case 0:
						if v28 == int32(0) {
							v58 = int32(1)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
							v50 = v40
							v51 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
							v58 = v55
						}
					case 1:
						if v28 == int32(0) {
							v58 = int32(2)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+376))
							v50 = v45
							v51 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
							v58 = v55
						}
					default:
						v47 = int32(0)
						if v28 == v47 {
							v55 = v47
						} else {
							v50 = v47
							v51 = v29
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
						}
						v58 = v55
					}
				}
				if v58 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_bytea_string_agg_transfn_0), int32(0))
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_string_agg_transfn_1), int32(416), int32(_a_F_bytea_string_agg_transfn_2))
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v61 = int32(_a_F_bytea_string_agg_transfn_3)
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0])) = v64
					v66 = F_makeStringInfo(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0])) = v62
						v71 = v66
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
						if v72 != 0 {
							v147 = int32(1)
							v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v151 = v149 & v147
							if v151 != 0 {
								v152 = v147
							} else {
								v152 = int32(4)
							}
							if v149 == int32(1) {
								v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
								if v159 == int32(18) {
									v162 = int32(16)
								} else {
									v162 = int32(0)
								}
								if base.Ui32((v159-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v169 = int32(4)
								} else {
									v169 = v162
								}
								v180 = v169
							} else {
								v170 = int32(1)
								if v151 != 0 {
									v180 = int32(base.Ui32(v149)>>(uint(v170)%32)) - v170
								} else {
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v180 = int32(base.Ui32(v174)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v71, v21+v152, v180)
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
								return int64(0)
							} else {
								v185 = v71
								if v185 != 0 {
									v193 = base.I64_extend_i32_u(v185)
								} else {
									v190 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v190)
									v193 = int64(0)
								}
								m.G0 = v12 + int32(16)
								return v193
							}
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v74 = F_pg_detoast_datum_packed(m, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
							} else {
								v76 = int32(1)
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								v80 = v78 & v76
								if v80 != 0 {
									v81 = v76
								} else {
									v81 = int32(4)
								}
								if v78 == int32(1) {
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
									if v88 == int32(18) {
										v91 = int32(16)
									} else {
										v91 = int32(0)
									}
									if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v98 = int32(4)
									} else {
										v98 = v91
									}
									v109 = v98
								} else {
									v99 = int32(1)
									if v80 != 0 {
										v109 = int32(base.Ui32(v78)>>(uint(v99)%32)) - v99
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_appendBinaryStringInfo(m, v71, v74+v81, v109)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int64(0)
								} else {
									if v18 != 0 {
									} else {
										v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
										if v112 == int32(1) {
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
											if v118 == int32(18) {
												v121 = int32(16)
											} else {
												v121 = int32(0)
											}
											if base.Ui32((v118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v128 = int32(4)
											} else {
												v128 = v121
											}
											v141 = v128
										} else {
											v129 = int32(1)
											if v112&v129 != 0 {
												v141 = int32(base.Ui32(v112)>>(uint(v129)%32)) - v129
											} else {
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
												v141 = int32(base.Ui32(v135)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v141
									}
									v147 = int32(1)
									v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									v151 = v149 & v147
									if v151 != 0 {
										v152 = v147
									} else {
										v152 = int32(4)
									}
									if v149 == int32(1) {
										v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
										if v159 == int32(18) {
											v162 = int32(16)
										} else {
											v162 = int32(0)
										}
										if base.Ui32((v159-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v169 = int32(4)
										} else {
											v169 = v162
										}
										v180 = v169
									} else {
										v170 = int32(1)
										if v151 != 0 {
											v180 = int32(base.Ui32(v149)>>(uint(v170)%32)) - v170
										} else {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
											v180 = int32(base.Ui32(v174)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									F_appendBinaryStringInfo(m, v71, v21+v152, v180)
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
										return int64(0)
									} else {
										v185 = v71
										if v185 != 0 {
											v193 = base.I64_extend_i32_u(v185)
										} else {
											v190 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v190)
											v193 = int64(0)
										}
										m.G0 = v12 + int32(16)
										return v193
									}
								}
							}
						}
					}
				}
			} else {
				v71 = v18
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
				if v72 != 0 {
					v147 = int32(1)
					v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					v151 = v149 & v147
					if v151 != 0 {
						v152 = v147
					} else {
						v152 = int32(4)
					}
					if v149 == int32(1) {
						v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
						if v159 == int32(18) {
							v162 = int32(16)
						} else {
							v162 = int32(0)
						}
						if base.Ui32((v159-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v169 = int32(4)
						} else {
							v169 = v162
						}
						v180 = v169
					} else {
						v170 = int32(1)
						if v151 != 0 {
							v180 = int32(base.Ui32(v149)>>(uint(v170)%32)) - v170
						} else {
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v180 = int32(base.Ui32(v174)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_appendBinaryStringInfo(m, v71, v21+v152, v180)
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return int64(0)
					} else {
						v185 = v71
						if v185 != 0 {
							v193 = base.I64_extend_i32_u(v185)
						} else {
							v190 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v190)
							v193 = int64(0)
						}
						m.G0 = v12 + int32(16)
						return v193
					}
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v74 = F_pg_detoast_datum_packed(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						v76 = int32(1)
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
						v80 = v78 & v76
						if v80 != 0 {
							v81 = v76
						} else {
							v81 = int32(4)
						}
						if v78 == int32(1) {
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
							if v88 == int32(18) {
								v91 = int32(16)
							} else {
								v91 = int32(0)
							}
							if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v98 = int32(4)
							} else {
								v98 = v91
							}
							v109 = v98
						} else {
							v99 = int32(1)
							if v80 != 0 {
								v109 = int32(base.Ui32(v78)>>(uint(v99)%32)) - v99
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_appendBinaryStringInfo(m, v71, v74+v81, v109)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int64(0)
						} else {
							if v18 != 0 {
							} else {
								v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								if v112 == int32(1) {
									v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
									if v118 == int32(18) {
										v121 = int32(16)
									} else {
										v121 = int32(0)
									}
									if base.Ui32((v118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v128 = int32(4)
									} else {
										v128 = v121
									}
									v141 = v128
								} else {
									v129 = int32(1)
									if v112&v129 != 0 {
										v141 = int32(base.Ui32(v112)>>(uint(v129)%32)) - v129
									} else {
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v141 = int32(base.Ui32(v135)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v141
							}
							v147 = int32(1)
							v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v151 = v149 & v147
							if v151 != 0 {
								v152 = v147
							} else {
								v152 = int32(4)
							}
							if v149 == int32(1) {
								v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
								if v159 == int32(18) {
									v162 = int32(16)
								} else {
									v162 = int32(0)
								}
								if base.Ui32((v159-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v169 = int32(4)
								} else {
									v169 = v162
								}
								v180 = v169
							} else {
								v170 = int32(1)
								if v151 != 0 {
									v180 = int32(base.Ui32(v149)>>(uint(v170)%32)) - v170
								} else {
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v180 = int32(base.Ui32(v174)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v71, v21+v152, v180)
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
								return int64(0)
							} else {
								v185 = v71
								if v185 != 0 {
									v193 = base.I64_extend_i32_u(v185)
								} else {
									v190 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v190)
									v193 = int64(0)
								}
								m.G0 = v12 + int32(16)
								return v193
							}
						}
					}
				}
			}
		}
	}
}
