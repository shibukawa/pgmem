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
	var v56 int32
	_ = v56
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_client_addr[0]))
	if v11 == v2 {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v56 = v2
		m.G0 = v8 + int32(256)
		return v56
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
					v56 = v2
					m.G0 = v8 + int32(256)
					return v56
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
						v56 = v54
						m.G0 = v8 + int32(256)
						return v56
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v56 = v2
			m.G0 = v8 + int32(256)
			return v56
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
					v37 = F_DirectFunctionCall1Coll(m, int32(1427), int32(0), v7)
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
	var v61 int32
	_ = v61
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 float64
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v511 float64
	_ = v511
	var v514 int32
	_ = v514
	var v515 float64
	_ = v515
	var v526 float64
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
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
	v112 = int32(base.Ui32(l1-int32(2))>>(uint(int32(10))%32)) + int32(1)
	if base.Ui32(l1) <= base.Ui32(v112) {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	v106 = v100
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
	v100 = v42 - v50
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
		v100 = v60
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v61 = int32(0)
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
	if base.B2i32(v79 < v78)&base.B2i32(v79 <= l3) != 0 {
		v100 = v61
		goto L8
	} else {
		goto L28
	}
L28:
	;
	if base.B2i32(v69 == v77)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v100 = v61
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v90 = int32(0)
	if v90 <= v78 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = l3
	goto L32
L31:
	;
	v93 = v90
	goto L32
L32:
	;
	if l3 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v93
	goto L35
L34:
	;
	v96 = l3
	goto L35
L35:
	;
	v106 = v96
	goto L7
L36:
	;
	return math.Float64frombits(uint64(0x7ff8000000000000))
L37:
	;
	goto L38
L38:
	;
	v117 = l3 + int32(1)
	v120 = v112
	v130 = v29
	v131 = v106
	v132 = int32(0)
	v136 = float64(0)
	goto L39
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0+v120<<(uint(int32(2))%32))))
	v141 = F_pg_detoast_datum_packed(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	return base.F64_div(v526, base.F64_convert_i32_s(v528))
L41:
	;
	v528 = v132 + int32(1)
	v529 = v120 + v112
	if v529 < l1 {
		v120 = v529
		v130 = v141
		v131 = v218
		v132 = v528
		v136 = v526
		goto L39
	} else {
		goto L179
	}
L42:
	;
	v147 = int32(1)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v149&v147 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v131|v218 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L44:
	;
	v218 = v212
	goto L43
L45:
	;
	v152 = v147
	goto L47
L46:
	;
	v152 = int32(4)
	goto L47
L47:
	;
	v153 = v141 + v152
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v155 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v157&v155 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v160 = v155
	goto L50
L49:
	;
	v160 = int32(4)
	goto L50
L50:
	;
	v161 = v24 + v160
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v154 == v162 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v164 = int32(2)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if base.Ui32(v168) < base.Ui32(v169) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v212 = v154 - v162
	goto L44
L54:
	;
	v171 = v168
	goto L56
L55:
	;
	v171 = v169
	goto L56
L56:
	;
	v172 = F_bitncmp(m, v153+v164, v161+v164, v171)
	mBase = m.M
	if v172 != 0 {
		v212 = v172
		goto L44
	} else {
		goto L57
	}
L57:
	;
	v173 = int32(0)
	v174 = int32(1)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v176&v174 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v179 = v174
	goto L60
L59:
	;
	v179 = int32(4)
	goto L60
L60:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v179)+1)))
	v182 = int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v184&v182 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v187 = v182
	goto L63
L62:
	;
	v187 = int32(4)
	goto L63
L63:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v187)+1)))
	v190 = v181 - v189
	v191 = int32(0)
	if base.B2i32(v191 < v190)&base.B2i32(v191 <= l3) != 0 {
		v212 = v173
		goto L44
	} else {
		goto L64
	}
L64:
	;
	if base.B2i32(v181 == v189)&base.B2i32(base.Ui32(l3+int32(1)) <= base.Ui32(int32(2))) != 0 {
		v212 = v173
		goto L44
	} else {
		goto L65
	}
L65:
	;
	v202 = int32(0)
	if v202 <= v190 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v205 = l3
	goto L68
L67:
	;
	v205 = v202
	goto L68
L68:
	;
	if l3 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v208 = v205
	goto L71
L70:
	;
	v208 = l3
	goto L71
L71:
	;
	v218 = v208
	goto L43
L72:
	;
	v526 = base.F64_add(v136, float64(1))
	goto L41
L73:
	;
	goto L74
L74:
	;
	v224 = int32(0)
	if base.B2i32(v131 <= v224)&base.B2i32(v224 <= v218) == v224 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v131 < int32(0) {
		v526 = v136
		goto L41
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v235 = int32(-1)
	v236 = int32(1)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v238&v236 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	if int32(0) < v218 {
		v526 = v136
		goto L41
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v372 = int32(1)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v374&v372 != 0 {
		goto L128
	} else {
		goto L129
	}
L81:
	;
	v241 = v236
	goto L83
L82:
	;
	v241 = int32(4)
	goto L83
L83:
	;
	v242 = v130 + v241
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v244 = int32(1)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v248 = v246 & v244
	if v248 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v249 = v244
	goto L86
L85:
	;
	v249 = int32(4)
	goto L86
L86:
	;
	v250 = v24 + v249
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v243 != v251 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v367 = int32(-1)
	v369 = v248
	goto L80
L88:
	;
	goto L89
L89:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v256 = v254 - v255
	v258 = base.B2i32(l3 < int32(0))
	if l3 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	if v280 == int32(0) {
		v367 = v279
		v369 = v248
		goto L80
	} else {
		goto L109
	}
L91:
	;
	if l3 != 0 {
		goto L106
	} else {
		goto L107
	}
L92:
	;
	if base.B2i32(v254 == v255)&base.B2i32(base.Ui32(v117) <= base.Ui32(int32(2))) != 0 {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if v256 <= int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(v254) < base.Ui32(v255) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v262 = v254
	goto L97
L96:
	;
	v262 = v255
	goto L97
L97:
	;
	v277 = v262
	goto L91
L98:
	;
	if base.Ui32(v254) < base.Ui32(v255) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if l3 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v269 = int32(0)
	if base.B2i32(v256 < v269)&base.B2i32(l3 <= v269) != 0 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v367 = int32(-1)
	v369 = v248
	goto L80
L102:
	;
	v276 = v254
	goto L104
L103:
	;
	v276 = v255
	goto L104
L104:
	;
	if l3 < int32(0) {
		v279 = v254
		v280 = v276
		goto L90
	} else {
		goto L105
	}
L105:
	;
	v277 = v276
	goto L91
L106:
	;
	v278 = v255
	goto L108
L107:
	;
	v278 = v277
	goto L108
L108:
	;
	v279 = v278
	v280 = v277
	goto L90
L109:
	;
	v283 = int32(2)
	v284 = v242 + v283
	v286 = v250 + v283
	v287 = int32(0)
	v292 = int32(8)
	v293 = base.I32_div_s(v280, v292)
	if v292 <= v280 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v367 = v279 - (v357 + v355<<(uint(int32(3))%32))
	v369 = v364 & int32(1)
	goto L80
L111:
	;
	goto L110
L112:
	;
	v343 = v334
	goto L123
L113:
	;
	v299 = v287
	goto L116
L114:
	;
	v316 = v287
	goto L115
L115:
	;
	v323 = v280 - v293<<(uint(int32(3))%32)
	if v323 == int32(0) {
		v355 = v316
		v357 = v287
		goto L111
	} else {
		goto L122
	}
L116:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v299))))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v299))))
	if v305 != v307 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v316 = v293
	goto L115
L118:
	;
	v333 = v299
	v334 = int32(7)
	v336 = v305
	v337 = v307
	goto L112
L119:
	;
	goto L120
L120:
	;
	v311 = v299 + int32(1)
	if v311 != v293 {
		v299 = v311
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v316))))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v316))))
	v333 = v316
	v334 = v323
	v336 = v329
	v337 = v327
	goto L112
L123:
	;
	if int32(base.Ui32(v336^v337)>>(uint(int32(8)-v343)%32)) != 0 {
		v343 = v343 - int32(1)
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v355 = v333
	v357 = v343
	goto L111
L125:
	;
	goto L124
L126:
	;
	v511 = float64(1)
	if v506 < v367 {
		goto L175
	} else {
		goto L176
	}
L127:
	;
	if int32(0) <= v367 {
		v506 = v497
		goto L126
	} else {
		goto L173
	}
L128:
	;
	v377 = v372
	goto L130
L129:
	;
	v377 = int32(4)
	goto L130
L130:
	;
	v378 = v141 + v377
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v369 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v382 = int32(1)
	goto L133
L132:
	;
	v382 = int32(4)
	goto L133
L133:
	;
	v383 = v24 + v382
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v379 != v384 {
		v497 = v235
		goto L127
	} else {
		goto L134
	}
L134:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+1)))
	v388 = v386 - v387
	v390 = base.B2i32(l3 < int32(0))
	if l3 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	if v412 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L136:
	;
	if l3 != 0 {
		goto L151
	} else {
		goto L152
	}
L137:
	;
	if base.B2i32(v387 == v386)&base.B2i32(base.Ui32(v117) <= base.Ui32(int32(2))) != 0 {
		goto L143
	} else {
		goto L144
	}
L138:
	;
	if v388 <= int32(0) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	if base.Ui32(v386) < base.Ui32(v387) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v394 = v386
	goto L142
L141:
	;
	v394 = v387
	goto L142
L142:
	;
	v410 = v394
	goto L136
L143:
	;
	if base.Ui32(v386) < base.Ui32(v387) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	if l3 == int32(0) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v401 = int32(0)
	if base.B2i32(v388 < v401)&base.B2i32(l3 <= v401) == v401 {
		v497 = v235
		goto L127
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v409 = v386
	goto L149
L148:
	;
	v409 = v387
	goto L149
L149:
	;
	if l3 < int32(0) {
		v412 = v409
		v413 = v386
		goto L135
	} else {
		goto L150
	}
L150:
	;
	v410 = v409
	goto L136
L151:
	;
	v411 = v387
	goto L153
L152:
	;
	v411 = v410
	goto L153
L153:
	;
	v412 = v410
	v413 = v411
	goto L135
L154:
	;
	v506 = v413
	goto L126
L155:
	;
	goto L156
L156:
	;
	v416 = int32(2)
	v417 = v378 + v416
	v419 = v383 + v416
	v420 = int32(0)
	v425 = int32(8)
	v426 = base.I32_div_s(v412, v425)
	if v425 <= v412 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v497 = v413 - (v490 + v488<<(uint(int32(3))%32))
	goto L127
L158:
	;
	goto L157
L159:
	;
	v476 = v467
	goto L170
L160:
	;
	v432 = v420
	goto L163
L161:
	;
	v449 = v420
	goto L162
L162:
	;
	v456 = v412 - v426<<(uint(int32(3))%32)
	if v456 == int32(0) {
		v488 = v449
		v490 = v420
		goto L158
	} else {
		goto L169
	}
L163:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v432))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+v432))))
	if v438 != v440 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v449 = v426
	goto L162
L165:
	;
	v466 = v432
	v467 = int32(7)
	v469 = v438
	v470 = v440
	goto L159
L166:
	;
	goto L167
L167:
	;
	v444 = v432 + int32(1)
	if v444 != v426 {
		v432 = v444
		goto L163
	} else {
		goto L168
	}
L168:
	;
	goto L164
L169:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+v449))))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v449))))
	v466 = v449
	v467 = v456
	v469 = v462
	v470 = v460
	goto L159
L170:
	;
	if int32(base.Ui32(v469^v470)>>(uint(int32(8)-v476)%32)) != 0 {
		v476 = v476 - int32(1)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v488 = v466
	v490 = v476
	goto L158
L172:
	;
	goto L171
L173:
	;
	if v497 < int32(0) {
		v526 = v136
		goto L41
	} else {
		goto L174
	}
L174:
	;
	v506 = v497
	goto L126
L175:
	;
	v514 = v367
	goto L177
L176:
	;
	v514 = v506
	goto L177
L177:
	;
	v515 = F_scalbn(m, v511, v514)
	mBase = m.M
	goto L178
L178:
	;
	v526 = base.F64_add(v136, base.F64_div(v511, v515))
	goto L41
L179:
	;
	goto L40
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
	var v56 int32
	_ = v56
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_inet_server_addr[0]))
	if v11 == v2 {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v56 = v2
		m.G0 = v8 + int32(256)
		return v56
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
					v56 = v2
					m.G0 = v8 + int32(256)
					return v56
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
						v56 = v54
						m.G0 = v8 + int32(256)
						return v56
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v56 = v2
			m.G0 = v8 + int32(256)
			return v56
		}
	}
}
