package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_inet_client_addr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_client_addr[0]))
	if v11 == v2 {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v58 = v2
		m.G0 = v8 + int32(256)
		return v58
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)))
		switch v16 - int32(2) {
		case 0, 8:
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v21)
			v24 = v11 + int32(144)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+272))
			v30 = F_pg_getnameinfo_all(m, v24, v25, v8, int32(255), v21, v21, int32(3))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v30 != 0 {
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					v58 = v2
					m.G0 = v8 + int32(256)
					return v58
				} else {
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
					if v36 != int32(10) {
					} else {
						v39 = int32(37)
						v40 = F___strchrnul(m, v8, v39)
						mBase = m.M
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
						if v42 == v39 {
							v46 = v40
						} else {
							v46 = int32(0)
						}
						if v46 == int32(0) {
						} else {
							v49 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v49)
						}
					}
					v52 = int32(0)
					v54 = F_network_in(m, v8, v52, v52)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = v54
						m.G0 = v8 + int32(256)
						return v58
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v58 = v2
			m.G0 = v8 + int32(256)
			return v58
		}
	}
}
func F_inet_client_port(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_inet_client_port[0]))
	if v10 == v2 {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v39 = v2
		m.G0 = v7 + int32(32)
		return v39
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)))
		switch v15 - int32(2) {
		case 0, 8:
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v20)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+272))
			v29 = F_pg_getnameinfo_all(m, v10+int32(144), v24, v20, v20, v7, int32(32), int32(3))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 != 0 {
					v33 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
					v39 = v2
					m.G0 = v7 + int32(32)
					return v39
				} else {
					v37 = F_DirectFunctionCall1Coll(m, int32(1408), int32(0), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v37
						m.G0 = v7 + int32(32)
						return v39
					}
				}
			}
		default:
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v39 = v2
			m.G0 = v7 + int32(32)
			return v39
		}
	}
}
func F_inet_hist_value_sel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v527 float64
	_ = v527
	var v530 int32
	_ = v530
	var v534 float64
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v550 float64
	_ = v550
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 float64
	_ = v563
	var v564 int32
	_ = v564
	var v582 float64
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	if l1 < int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	goto L3
L3:
	;
	v24 = F_pg_detoast_datum_packed(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return float64(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v37&v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v113 = int32(base.Ui32(l1-int32(2))>>(uint(int32(10))%32)) + int32(1)
	if base.Ui32(l1) <= base.Ui32(v113) {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	v107 = v101
	goto L7
L9:
	;
	v40 = v35
	goto L11
L10:
	;
	v40 = int32(4)
	goto L11
L11:
	;
	v41 = v29 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v43 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v45&v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v43
	goto L14
L13:
	;
	v48 = int32(4)
	goto L14
L14:
	;
	v49 = v24 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v42 == v50 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = int32(2)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if base.Ui32(v56) < base.Ui32(v57) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v101 = v42 - v50
	goto L8
L18:
	;
	v59 = v56
	goto L20
L19:
	;
	v59 = v57
	goto L20
L20:
	;
	v60 = F_bitncmp(m, v41+v52, v49+v52, v59)
	mBase = m.M
	if v60 != 0 {
		v101 = v60
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v62 = int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v64&v62 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = v62
	goto L24
L23:
	;
	v67 = int32(4)
	goto L24
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v67)+1)))
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v72&v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = v70
	goto L27
L26:
	;
	v75 = int32(4)
	goto L27
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v75)+1)))
	v78 = v69 - v77
	v79 = int32(0)
	if base.B2i32(v79 < v78)&base.B2i32(v79 <= l3)|base.B2i32(v69 == v77)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v101 = int32(0)
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v91 = int32(0)
	if v91 <= v78 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = l3
	goto L31
L30:
	;
	v94 = v91
	goto L31
L31:
	;
	if l3 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v94
	goto L34
L33:
	;
	v97 = l3
	goto L34
L34:
	;
	v107 = v97
	goto L7
L35:
	;
	return math.Float64frombits(uint64(0x7ff8000000000000))
L36:
	;
	goto L37
L37:
	;
	v118 = l3 + int32(1)
	v121 = v113
	v130 = v29
	v131 = v107
	v136 = int32(0)
	v137 = float64(0)
	goto L38
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0+v121<<(uint(int32(2))%32))))
	v142 = F_pg_detoast_datum_packed(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	return base.F64_div(v582, base.F64_convert_i32_s(v584))
L40:
	;
	v584 = v136 + int32(1)
	v585 = v121 + v113
	if v585 < l1 {
		v121 = v585
		v130 = v142
		v131 = v220
		v136 = v584
		v137 = v582
		goto L38
	} else {
		goto L186
	}
L41:
	;
	v148 = int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v150&v148 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v131|v220 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L43:
	;
	v220 = v214
	goto L42
L44:
	;
	v153 = v148
	goto L46
L45:
	;
	v153 = int32(4)
	goto L46
L46:
	;
	v154 = v142 + v153
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v156 = int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v158&v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v161 = v156
	goto L49
L48:
	;
	v161 = int32(4)
	goto L49
L49:
	;
	v162 = v24 + v161
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v155 == v163 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v165 = int32(2)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if base.Ui32(v169) < base.Ui32(v170) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v214 = v155 - v163
	goto L43
L53:
	;
	v172 = v169
	goto L55
L54:
	;
	v172 = v170
	goto L55
L55:
	;
	v173 = F_bitncmp(m, v154+v165, v162+v165, v172)
	mBase = m.M
	if v173 != 0 {
		v214 = v173
		goto L43
	} else {
		goto L56
	}
L56:
	;
	v175 = int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v177&v175 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v180 = v175
	goto L59
L58:
	;
	v180 = int32(4)
	goto L59
L59:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v180)+1)))
	v183 = int32(1)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v185&v183 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v188 = v183
	goto L62
L61:
	;
	v188 = int32(4)
	goto L62
L62:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v188)+1)))
	v191 = v182 - v190
	v192 = int32(0)
	if base.B2i32(v192 < v191)&base.B2i32(v192 <= l3)|base.B2i32(v182 == v190)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v214 = int32(0)
		goto L43
	} else {
		goto L63
	}
L63:
	;
	v204 = int32(0)
	if v204 <= v191 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v207 = l3
	goto L66
L65:
	;
	v207 = v204
	goto L66
L66:
	;
	if l3 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v210 = v207
	goto L69
L68:
	;
	v210 = l3
	goto L69
L69:
	;
	v220 = v210
	goto L42
L70:
	;
	v582 = base.F64_add(v137, float64(1))
	goto L40
L71:
	;
	goto L72
L72:
	;
	v226 = int32(0)
	if base.B2i32(base.B2i32(v131 <= v226)&base.B2i32(v226 <= v220) == v226)&(base.B2i32(v131 < v226)|base.B2i32(v226 < v220)) != 0 {
		v582 = v137
		goto L40
	} else {
		goto L73
	}
L73:
	;
	v239 = int32(-1)
	v240 = int32(1)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v242&v240 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v383 = int32(1)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v385&v383 != 0 {
		goto L121
	} else {
		goto L122
	}
L75:
	;
	v245 = v240
	goto L77
L76:
	;
	v245 = int32(4)
	goto L77
L77:
	;
	v246 = v130 + v245
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v248 = int32(1)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v252 = v250 & v248
	if v252 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v253 = v248
	goto L80
L79:
	;
	v253 = int32(4)
	goto L80
L80:
	;
	v254 = v24 + v253
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v247 != v255 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v378 = v252
	v379 = int32(-1)
	goto L74
L82:
	;
	goto L83
L83:
	;
	v258 = int32(0)
	v259 = base.B2i32(l3 < v258)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	v262 = v260 - v261
	if v259|base.B2i32(v262 <= v258) == v258 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v291 == int32(0) {
		v378 = v252
		v379 = v290
		goto L74
	} else {
		goto L102
	}
L85:
	;
	if l3 != 0 {
		goto L99
	} else {
		goto L100
	}
L86:
	;
	if base.Ui32(v260) < base.Ui32(v261) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v270 = int32(0)
	if base.B2i32(v262 < v270)&base.B2i32(l3 <= v270)|(base.B2i32(l3 == v270)|base.B2i32(v260 == v261)&base.B2i32(base.Ui32(v118) <= base.Ui32(int32(2)))) == v270 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v269 = v260
	goto L91
L90:
	;
	v269 = v261
	goto L91
L91:
	;
	v288 = v269
	goto L85
L92:
	;
	v378 = v252
	v379 = int32(-1)
	goto L74
L93:
	;
	goto L94
L94:
	;
	if base.Ui32(v260) < base.Ui32(v261) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v287 = v260
	goto L97
L96:
	;
	v287 = v261
	goto L97
L97:
	;
	if l3 < v258 {
		v290 = v260
		v291 = v287
		goto L84
	} else {
		goto L98
	}
L98:
	;
	v288 = v287
	goto L85
L99:
	;
	v289 = v261
	goto L101
L100:
	;
	v289 = v288
	goto L101
L101:
	;
	v290 = v289
	v291 = v288
	goto L84
L102:
	;
	v294 = int32(2)
	v295 = v246 + v294
	v297 = v254 + v294
	v298 = int32(0)
	v303 = int32(8)
	v304 = base.I32_div_s(v291, v303)
	if v303 <= v291 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v378 = v375 & int32(1)
	v379 = v290 - (v370 + v366<<(uint(int32(3))%32))
	goto L74
L104:
	;
	goto L103
L105:
	;
	v352 = v343
	goto L116
L106:
	;
	v310 = v298
	goto L109
L107:
	;
	v327 = v298
	goto L108
L108:
	;
	v334 = v291 - v304<<(uint(int32(3))%32)
	if v334 == int32(0) {
		v366 = v327
		v370 = v298
		goto L104
	} else {
		goto L115
	}
L109:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v310))))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v310))))
	if v316 != v318 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v327 = v304
	goto L108
L111:
	;
	v343 = int32(7)
	v344 = v310
	v346 = v316
	v347 = v318
	goto L105
L112:
	;
	goto L113
L113:
	;
	v322 = v310 + int32(1)
	if v322 != v304 {
		v310 = v322
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v327))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v327))))
	v343 = v334
	v344 = v327
	v346 = v340
	v347 = v338
	goto L105
L116:
	;
	if int32(base.Ui32(v346^v347)>>(uint(int32(8)-v352)%32)) != 0 {
		v352 = v352 - int32(1)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v366 = v344
	v370 = v352
	goto L104
L118:
	;
	goto L117
L119:
	;
	v527 = float64(1)
	if v523 < v379 {
		goto L165
	} else {
		goto L166
	}
L120:
	;
	if int32(0) <= v379 {
		v523 = v514
		goto L119
	} else {
		goto L163
	}
L121:
	;
	v388 = v383
	goto L123
L122:
	;
	v388 = int32(4)
	goto L123
L123:
	;
	v389 = v142 + v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v378 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v393 = int32(1)
	goto L126
L125:
	;
	v393 = int32(4)
	goto L126
L126:
	;
	v394 = v24 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v390 != v395 {
		v514 = v239
		goto L120
	} else {
		goto L127
	}
L127:
	;
	v397 = int32(0)
	v398 = base.B2i32(l3 < v397)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	v401 = v399 - v400
	if v398|base.B2i32(v401 <= v397) == v397 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	if v429 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L129:
	;
	if l3 != 0 {
		goto L141
	} else {
		goto L142
	}
L130:
	;
	if base.Ui32(v399) < base.Ui32(v400) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v409 = int32(0)
	if base.B2i32(l3 == v409)|base.B2i32(v399 == v400)&base.B2i32(base.Ui32(v118) <= base.Ui32(int32(2)))|base.B2i32(v401 < v409)&base.B2i32(l3 <= v409) == v409 {
		v514 = v239
		goto L120
	} else {
		goto L136
	}
L133:
	;
	v408 = v399
	goto L135
L134:
	;
	v408 = v400
	goto L135
L135:
	;
	v426 = v408
	goto L129
L136:
	;
	if base.Ui32(v399) < base.Ui32(v400) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v425 = v399
	goto L139
L138:
	;
	v425 = v400
	goto L139
L139:
	;
	if l3 < v397 {
		v428 = v399
		v429 = v425
		goto L128
	} else {
		goto L140
	}
L140:
	;
	v426 = v425
	goto L129
L141:
	;
	v427 = v400
	goto L143
L142:
	;
	v427 = v426
	goto L143
L143:
	;
	v428 = v427
	v429 = v426
	goto L128
L144:
	;
	v523 = v428
	goto L119
L145:
	;
	goto L146
L146:
	;
	v432 = int32(2)
	v433 = v389 + v432
	v435 = v394 + v432
	v436 = int32(0)
	v441 = int32(8)
	v442 = base.I32_div_s(v429, v441)
	if v441 <= v429 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v514 = v428 - (v508 + v504<<(uint(int32(3))%32))
	goto L120
L148:
	;
	goto L147
L149:
	;
	v490 = v481
	goto L160
L150:
	;
	v448 = v436
	goto L153
L151:
	;
	v465 = v436
	goto L152
L152:
	;
	v472 = v429 - v442<<(uint(int32(3))%32)
	if v472 == int32(0) {
		v504 = v465
		v508 = v436
		goto L148
	} else {
		goto L159
	}
L153:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v448))))
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v448))))
	if v454 != v456 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v465 = v442
	goto L152
L155:
	;
	v481 = int32(7)
	v482 = v448
	v484 = v454
	v485 = v456
	goto L149
L156:
	;
	goto L157
L157:
	;
	v460 = v448 + int32(1)
	if v460 != v442 {
		v448 = v460
		goto L153
	} else {
		goto L158
	}
L158:
	;
	goto L154
L159:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v465))))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v465))))
	v481 = v472
	v482 = v465
	v484 = v478
	v485 = v476
	goto L149
L160:
	;
	if int32(base.Ui32(v484^v485)>>(uint(int32(8)-v490)%32)) != 0 {
		v490 = v490 - int32(1)
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v504 = v482
	v508 = v490
	goto L148
L162:
	;
	goto L161
L163:
	;
	if v514 < int32(0) {
		v582 = v137
		goto L40
	} else {
		goto L164
	}
L164:
	;
	v523 = v514
	goto L119
L165:
	;
	v530 = v379
	goto L167
L166:
	;
	v530 = v523
	goto L167
L167:
	;
	if int32(1024) <= v530 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v582 = base.F64_add(v137, base.F64_div(v527, base.F64_mul(v563, base.F64_reinterpret_i64(base.I64_extend_i32_u(v564+int32(1023))<<(uint(int64(52))%64)))))
	goto L40
L169:
	;
	goto L168
L170:
	;
	v534 = base.F64_mul(v527, float64(8.98846567431158e+307))
	if base.Ui32(v530) < base.Ui32(int32(2047)) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	if int32(-1023) < v530 {
		v563 = v527
		v564 = v530
		goto L169
	} else {
		goto L179
	}
L173:
	;
	v563 = v534
	v564 = v530 - int32(1023)
	goto L169
L174:
	;
	goto L175
L175:
	;
	v541 = int32(3069)
	if base.Ui32(v541) <= base.Ui32(v530) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v544 = v541
	goto L178
L177:
	;
	v544 = v530
	goto L178
L178:
	;
	v563 = base.F64_mul(v534, float64(8.98846567431158e+307))
	v564 = v544 - int32(2046)
	goto L169
L179:
	;
	v550 = base.F64_mul(v527, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v530) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v563 = v550
	v564 = v530 + int32(969)
	goto L169
L181:
	;
	goto L182
L182:
	;
	v557 = int32(-2960)
	if base.Ui32(v530) <= base.Ui32(v557) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v560 = v557
	goto L185
L184:
	;
	v560 = v530
	goto L185
L185:
	;
	v563 = base.F64_mul(v550, float64(2.004168360008973e-292))
	v564 = v560 + int32(1938)
	goto L169
L186:
	;
	goto L39
}
func F_inet_same_family(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(1)
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v13&v11 != 0 {
				v16 = v11
			} else {
				v16 = int32(4)
			}
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v16))))
			v19 = int32(1)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v21&v19 != 0 {
				v24 = v19
			} else {
				v24 = int32(4)
			}
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v24))))
			return base.B2i32(v18 == v26)
		}
	}
}
func F_inet_server_addr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_server_addr[0]))
	if v11 == v2 {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v58 = v2
		m.G0 = v8 + int32(256)
		return v58
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
		switch v16 - int32(2) {
		case 0, 8:
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v21)
			v24 = v11 + int32(12)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
			v30 = F_pg_getnameinfo_all(m, v24, v25, v8, int32(255), v21, v21, int32(3))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v30 != 0 {
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					v58 = v2
					m.G0 = v8 + int32(256)
					return v58
				} else {
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
					if v36 != int32(10) {
					} else {
						v39 = int32(37)
						v40 = F___strchrnul(m, v8, v39)
						mBase = m.M
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
						if v42 == v39 {
							v46 = v40
						} else {
							v46 = int32(0)
						}
						if v46 == int32(0) {
						} else {
							v49 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v49)
						}
					}
					v52 = int32(0)
					v54 = F_network_in(m, v8, v52, v52)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = v54
						m.G0 = v8 + int32(256)
						return v58
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v58 = v2
			m.G0 = v8 + int32(256)
			return v58
		}
	}
}
