package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_inet_client_addr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_client_addr[0]))
	if v11 == int32(0) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v62 = int64(0)
		m.G0 = v8 + int32(256)
		return v62
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)))
		switch v17 - int32(2) {
		case 0, 8:
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v23)
			v26 = v11 + int32(144)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+272))
			v32 = F_pg_getnameinfo_all(m, v26, v27, v8, int32(255), v23, v23, int32(3))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				if v32 != 0 {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
					v62 = int64(0)
					m.G0 = v8 + int32(256)
					return v62
				} else {
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
					if v39 != int32(10) {
					} else {
						v42 = int32(37)
						v43 = F___strchrnul(m, v8, v42)
						mBase = m.M
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
						if v45 == v42 {
							v49 = v43
						} else {
							v49 = int32(0)
						}
						if v49 == int32(0) {
						} else {
							v52 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v52)
						}
					}
					v55 = int32(0)
					v57 = F_network_in(m, v8, v55, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						v62 = base.I64_extend_i32_u(v57)
						m.G0 = v8 + int32(256)
						return v62
					}
				}
			}
		default:
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v62 = int64(0)
			m.G0 = v8 + int32(256)
			return v62
		}
	}
}
func F_inet_client_port(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_inet_client_port[0]))
	if v10 == int32(0) {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v43 = int64(0)
		m.G0 = v7 + int32(32)
		return v43
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)))
		switch v16 - int32(2) {
		case 0, 8:
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v22)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+272))
			v31 = F_pg_getnameinfo_all(m, v10+int32(144), v26, v22, v22, v7, int32(32), int32(3))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				if v31 != 0 {
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
					v43 = int64(0)
					m.G0 = v7 + int32(32)
					return v43
				} else {
					v41 = F_DirectFunctionCall1Coll(m, int32(1142), int32(0), base.I64_extend_i32_u(v7))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = v41
						m.G0 = v7 + int32(32)
						return v43
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v43 = int64(0)
			m.G0 = v7 + int32(32)
			return v43
		}
	}
}
func F_inet_hist_value_sel(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
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
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
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
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v523 float64
	_ = v523
	var v526 int32
	_ = v526
	var v530 float64
	_ = v530
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 float64
	_ = v546
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 float64
	_ = v559
	var v560 int32
	_ = v560
	var v577 float64
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
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
	v25 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return float64(0)
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_pg_detoast_datum_packed(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v38&v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v115 = int32(base.Ui32(l1-int32(2))>>(uint(int32(10))%32)) + int32(1)
	if base.Ui32(l1) <= base.Ui32(v115) {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	v109 = v103
	goto L7
L9:
	;
	v41 = v36
	goto L11
L10:
	;
	v41 = int32(4)
	goto L11
L11:
	;
	v42 = v30 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v44 = int32(1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v46&v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = v44
	goto L14
L13:
	;
	v49 = int32(4)
	goto L14
L14:
	;
	v50 = v25 + v49
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v43 == v51 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = int32(2)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if base.Ui32(v57) < base.Ui32(v58) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v103 = v43 - v51
	goto L8
L18:
	;
	v60 = v42
	goto L20
L19:
	;
	v60 = v50
	goto L20
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v62 = F_bitncmp(m, v42+v53, v50+v53, v61)
	mBase = m.M
	if v62 != 0 {
		v103 = v62
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v64 = int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v66&v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v64
	goto L24
L23:
	;
	v69 = int32(4)
	goto L24
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v69)+1)))
	v72 = int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v74&v72 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v77 = v72
	goto L27
L26:
	;
	v77 = int32(4)
	goto L27
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v77)+1)))
	v80 = v71 - v79
	v81 = int32(0)
	if base.B2i32(v81 < v80)&base.B2i32(v81 <= l3)|base.B2i32(v71 == v79)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v103 = int32(0)
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v93 = int32(0)
	if v93 <= v80 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v96 = l3
	goto L31
L30:
	;
	v96 = v93
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
	v99 = v96
	goto L34
L33:
	;
	v99 = l3
	goto L34
L34:
	;
	v109 = v99
	goto L7
L35:
	;
	return math.Float64frombits(uint64(0x7ff8000000000000))
L36:
	;
	goto L37
L37:
	;
	v120 = l3 + int32(1)
	v132 = v30
	v133 = v109
	v134 = v115
	v138 = int32(0)
	v139 = float64(0)
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0+v134<<(uint(int32(3))%32))))
	v144 = F_pg_detoast_datum_packed(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	return base.F64_div(v577, base.F64_convert_i32_s(v579))
L40:
	;
	v579 = v138 + int32(1)
	v580 = v134 + v115
	if v580 < l1 {
		v132 = v144
		v133 = v223
		v134 = v580
		v138 = v579
		v139 = v577
		goto L38
	} else {
		goto L172
	}
L41:
	;
	v150 = int32(1)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v152&v150 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v133|v223 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L43:
	;
	v223 = v217
	goto L42
L44:
	;
	v155 = v150
	goto L46
L45:
	;
	v155 = int32(4)
	goto L46
L46:
	;
	v156 = v144 + v155
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v158 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v160&v158 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v163 = v158
	goto L49
L48:
	;
	v163 = int32(4)
	goto L49
L49:
	;
	v164 = v25 + v163
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v157 == v165 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v167 = int32(2)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	if base.Ui32(v171) < base.Ui32(v172) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v217 = v157 - v165
	goto L43
L53:
	;
	v174 = v156
	goto L55
L54:
	;
	v174 = v164
	goto L55
L55:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	v176 = F_bitncmp(m, v156+v167, v164+v167, v175)
	mBase = m.M
	if v176 != 0 {
		v217 = v176
		goto L43
	} else {
		goto L56
	}
L56:
	;
	v178 = int32(1)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v180&v178 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v183 = v178
	goto L59
L58:
	;
	v183 = int32(4)
	goto L59
L59:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v183)+1)))
	v186 = int32(1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v188&v186 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v191 = v186
	goto L62
L61:
	;
	v191 = int32(4)
	goto L62
L62:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v191)+1)))
	v194 = v185 - v193
	v195 = int32(0)
	if base.B2i32(v195 < v194)&base.B2i32(v195 <= l3)|base.B2i32(v185 == v193)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v217 = int32(0)
		goto L43
	} else {
		goto L63
	}
L63:
	;
	v207 = int32(0)
	if v207 <= v194 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v210 = l3
	goto L66
L65:
	;
	v210 = v207
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
	v213 = v210
	goto L69
L68:
	;
	v213 = l3
	goto L69
L69:
	;
	v223 = v213
	goto L42
L70:
	;
	v577 = base.F64_add(v139, float64(1))
	goto L40
L71:
	;
	goto L72
L72:
	;
	v229 = int32(0)
	if base.B2i32(base.B2i32(v133 <= v229)&base.B2i32(v229 <= v223) == v229)&(base.B2i32(v133 < v229)|base.B2i32(v229 < v223)) != 0 {
		v577 = v139
		goto L40
	} else {
		goto L73
	}
L73:
	;
	v242 = int32(-1)
	v243 = int32(1)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v245&v243 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v383 = int32(1)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v385&v383 != 0 {
		goto L115
	} else {
		goto L116
	}
L75:
	;
	v248 = v243
	goto L77
L76:
	;
	v248 = int32(4)
	goto L77
L77:
	;
	v249 = v132 + v248
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v251 = int32(1)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v255 = v253 & v251
	if v255 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v256 = v251
	goto L80
L79:
	;
	v256 = int32(4)
	goto L80
L80:
	;
	v257 = v25 + v256
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if v250 != v258 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v379 = int32(-1)
	v381 = v255
	goto L74
L82:
	;
	goto L83
L83:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	v263 = v261 - v262
	v264 = int32(0)
	if base.B2i32(v264 < v263)&base.B2i32(v264 <= l3)|base.B2i32(v261 == v262)&base.B2i32(base.Ui32(v120) < base.Ui32(int32(3)))|(base.B2i32(l3 == v264)|base.B2i32(v263 < v264)&base.B2i32(l3 <= v264)) == v264 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v379 = int32(-1)
	v381 = v255
	goto L74
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v261) < base.Ui32(v262) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v287 = v249
	goto L89
L88:
	;
	v287 = v257
	goto L89
L89:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	if l3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v289 = v262
	goto L92
L91:
	;
	v289 = v288
	goto L92
L92:
	;
	if l3 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v292 = v261
	goto L95
L94:
	;
	v292 = v289
	goto L95
L95:
	;
	if v288 == int32(0) {
		v379 = v292
		v381 = v255
		goto L74
	} else {
		goto L96
	}
L96:
	;
	v295 = int32(2)
	v296 = v249 + v295
	v298 = v257 + v295
	v299 = int32(0)
	v304 = int32(8)
	v305 = base.I32_div_s(v288, v304)
	if v304 <= v288 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v379 = v292 - (v371 + v367<<(uint(int32(3))%32))
	v381 = v376 & int32(1)
	goto L74
L98:
	;
	goto L97
L99:
	;
	v353 = v344
	goto L110
L100:
	;
	v311 = v299
	goto L103
L101:
	;
	v328 = v299
	goto L102
L102:
	;
	v335 = v288 - v305<<(uint(int32(3))%32)
	if v335 == int32(0) {
		v367 = v328
		v371 = v299
		goto L98
	} else {
		goto L109
	}
L103:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v311))))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v311))))
	if v317 != v319 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v328 = v305
	goto L102
L105:
	;
	v344 = int32(7)
	v345 = v311
	v347 = v317
	v348 = v319
	goto L99
L106:
	;
	goto L107
L107:
	;
	v323 = v311 + int32(1)
	if v323 != v305 {
		v311 = v323
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v328))))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v328))))
	v344 = v335
	v345 = v328
	v347 = v341
	v348 = v339
	goto L99
L110:
	;
	if int32(base.Ui32(v347^v348)>>(uint(int32(8)-v353)%32)) != 0 {
		v353 = v353 - int32(1)
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v367 = v345
	v371 = v353
	goto L98
L112:
	;
	goto L111
L113:
	;
	v523 = float64(1)
	if v520 < v379 {
		goto L151
	} else {
		goto L152
	}
L114:
	;
	if int32(0) <= v379 {
		v520 = v512
		goto L113
	} else {
		goto L149
	}
L115:
	;
	v388 = v383
	goto L117
L116:
	;
	v388 = int32(4)
	goto L117
L117:
	;
	v389 = v144 + v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v381 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v393 = int32(1)
	goto L120
L119:
	;
	v393 = int32(4)
	goto L120
L120:
	;
	v394 = v25 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v390 != v395 {
		v512 = v242
		goto L114
	} else {
		goto L121
	}
L121:
	;
	v397 = int32(0)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	v401 = v399 - v400
	if base.B2i32(l3 == v397)|(base.B2i32(v397 < v401)&base.B2i32(v397 <= l3)|base.B2i32(v399 == v400)&base.B2i32(base.Ui32(v120) < base.Ui32(int32(3))))|base.B2i32(v401 < v397)&base.B2i32(l3 <= v397) == v397 {
		v512 = v242
		goto L114
	} else {
		goto L122
	}
L122:
	;
	if base.Ui32(v399) < base.Ui32(v400) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v422 = v389
	goto L125
L124:
	;
	v422 = v394
	goto L125
L125:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	if l3 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v424 = v400
	goto L128
L127:
	;
	v424 = v423
	goto L128
L128:
	;
	if l3 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v427 = v399
	goto L131
L130:
	;
	v427 = v424
	goto L131
L131:
	;
	if v423 == int32(0) {
		v520 = v427
		goto L113
	} else {
		goto L132
	}
L132:
	;
	v430 = int32(2)
	v431 = v389 + v430
	v433 = v394 + v430
	v434 = int32(0)
	v439 = int32(8)
	v440 = base.I32_div_s(v423, v439)
	if v439 <= v423 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v512 = v427 - (v506 + v502<<(uint(int32(3))%32))
	goto L114
L134:
	;
	goto L133
L135:
	;
	v488 = v479
	goto L146
L136:
	;
	v446 = v434
	goto L139
L137:
	;
	v463 = v434
	goto L138
L138:
	;
	v470 = v423 - v440<<(uint(int32(3))%32)
	if v470 == int32(0) {
		v502 = v463
		v506 = v434
		goto L134
	} else {
		goto L145
	}
L139:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431+v446))))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v446))))
	if v452 != v454 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v463 = v440
	goto L138
L141:
	;
	v479 = int32(7)
	v480 = v446
	v482 = v452
	v483 = v454
	goto L135
L142:
	;
	goto L143
L143:
	;
	v458 = v446 + int32(1)
	if v458 != v440 {
		v446 = v458
		goto L139
	} else {
		goto L144
	}
L144:
	;
	goto L140
L145:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v463))))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431+v463))))
	v479 = v470
	v480 = v463
	v482 = v476
	v483 = v474
	goto L135
L146:
	;
	if int32(base.Ui32(v482^v483)>>(uint(int32(8)-v488)%32)) != 0 {
		v488 = v488 - int32(1)
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v502 = v480
	v506 = v488
	goto L134
L148:
	;
	goto L147
L149:
	;
	if v512 < int32(0) {
		v577 = v139
		goto L40
	} else {
		goto L150
	}
L150:
	;
	v520 = v512
	goto L113
L151:
	;
	v526 = v379
	goto L153
L152:
	;
	v526 = v520
	goto L153
L153:
	;
	if int32(1024) <= v526 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v577 = base.F64_add(v139, base.F64_div(v523, base.F64_mul(v559, base.F64_reinterpret_i64(base.I64_extend_i32_u(v560+int32(1023))<<(uint(int64(52))%64)))))
	goto L40
L155:
	;
	goto L154
L156:
	;
	v530 = base.F64_mul(v523, float64(8.98846567431158e+307))
	if base.Ui32(v526) < base.Ui32(int32(2047)) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if int32(-1023) < v526 {
		v559 = v523
		v560 = v526
		goto L155
	} else {
		goto L165
	}
L159:
	;
	v559 = v530
	v560 = v526 - int32(1023)
	goto L155
L160:
	;
	goto L161
L161:
	;
	v537 = int32(3069)
	if base.Ui32(v537) <= base.Ui32(v526) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v540 = v537
	goto L164
L163:
	;
	v540 = v526
	goto L164
L164:
	;
	v559 = base.F64_mul(v530, float64(8.98846567431158e+307))
	v560 = v540 - int32(2046)
	goto L155
L165:
	;
	v546 = base.F64_mul(v523, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v526) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v559 = v546
	v560 = v526 + int32(969)
	goto L155
L167:
	;
	goto L168
L168:
	;
	v553 = int32(-2960)
	if base.Ui32(v526) <= base.Ui32(v553) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v556 = v553
	goto L171
L170:
	;
	v556 = v526
	goto L171
L171:
	;
	v559 = base.F64_mul(v546, float64(2.004168360008973e-292))
	v560 = v556 + int32(1938)
	goto L155
L172:
	;
	goto L39
}
func F_inet_same_family(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
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
			return base.I64_extend_i32_u(base.B2i32(v18 == v26))
		}
	}
}
func F_inet_server_addr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_server_addr[0]))
	if v11 == int32(0) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v62 = int64(0)
		m.G0 = v8 + int32(256)
		return v62
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
		switch v17 - int32(2) {
		case 0, 8:
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v23)
			v26 = v11 + int32(12)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
			v32 = F_pg_getnameinfo_all(m, v26, v27, v8, int32(255), v23, v23, int32(3))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				if v32 != 0 {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
					v62 = int64(0)
					m.G0 = v8 + int32(256)
					return v62
				} else {
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
					if v39 != int32(10) {
					} else {
						v42 = int32(37)
						v43 = F___strchrnul(m, v8, v42)
						mBase = m.M
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
						if v45 == v42 {
							v49 = v43
						} else {
							v49 = int32(0)
						}
						if v49 == int32(0) {
						} else {
							v52 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v52)
						}
					}
					v55 = int32(0)
					v57 = F_network_in(m, v8, v55, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						v62 = base.I64_extend_i32_u(v57)
						m.G0 = v8 + int32(256)
						return v62
					}
				}
			}
		default:
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v62 = int64(0)
			m.G0 = v8 + int32(256)
			return v62
		}
	}
}
