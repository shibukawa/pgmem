package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_stats_check_arg_array(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v12 = l0 + l1<<(uint(int32(4))%32)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
	if v13 != 0 {
		v63 = v9
		m.G0 = v7 + int32(16)
		return v63
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v19 != int32(1) {
				v22 = int32(0)
				v25 = F_errstart(m, int32(19), v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v63 = v22
						m.G0 = v7 + int32(16)
						return v63
					} else {
						v44 = int32(88)
						v45 = int32(_a_F_stats_check_arg_array_0)
						F_errcode(m, int32(50856066))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_array[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v53
							F_errmsg(m, v45, v7)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_stats_check_arg_array_1), v44, int32(_a_F_stats_check_arg_array_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v63 = int32(0)
									m.G0 = v7 + int32(16)
									return v63
								}
							}
						}
					}
				}
			} else {
				v31 = F_array_contains_nulls(m, v15)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v63 = v9
						m.G0 = v7 + int32(16)
						return v63
					} else {
						v35 = int32(0)
						v38 = F_errstart(m, int32(19), v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == int32(0) {
								v63 = v35
								m.G0 = v7 + int32(16)
								return v63
							} else {
								v44 = int32(97)
								v45 = int32(_a_F_stats_check_arg_array_3)
								F_errcode(m, int32(50856066))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_array[0])))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v53
									F_errmsg(m, v45, v7)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_stats_check_arg_array_1), v44, int32(_a_F_stats_check_arg_array_2))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v63 = int32(0)
											m.G0 = v7 + int32(16)
											return v63
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
func F_stats_fill_fcinfo_from_arg_pairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int64
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v60 = int32(1)
	v67 = F_extract_variadic_args(m, l0, v16+int32(76), v16+int32(68), v16+int32(72))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v36 = l1 + int32(24) + v26<<(uint(int32(4))%32)
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)) = uint8(v37)
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = int64(0)
	v42 = v26 + v37
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2+v42<<(uint(int32(3))%32))))
	if v46 != 0 {
		v26 = v42
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	return int32(0)
L8:
	;
	if v67&int32(1) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(0) < v67 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L85
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L79
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L75
	}
L14:
	;
	v81 = int32(0)
	v84 = v60
	goto L17
L15:
	;
	v335 = v60
	goto L16
L16:
	;
	m.G0 = v16 + int32(80)
	return v335 & int32(1)
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v81))))
	if v95 == int32(1) {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v335 = v319
	goto L16
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v81<<(uint(int32(2))%32))))
	if v102 != int32(25) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v106 = v81 | int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v106))))
	if v108 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v329 = v81 + int32(2)
	if v329 < v67 {
		v81 = v329
		v84 = v319
		goto L17
	} else {
		goto L74
	}
L22:
	;
	v319 = v84
	goto L21
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v81<<(uint(int32(3))%32))))
	v114 = F_text_to_cstring(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v119 = v114
	v120 = int32(_a_F_stats_fill_fcinfo_from_arg_pairs_0)
	goto L26
L25:
	;
	if v157 == int32(0) {
		goto L22
	} else {
		goto L38
	}
L26:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v123 == v124 {
		v146 = v123
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v157 = int32(0)
	goto L25
L28:
	;
	v148 = int32(1)
	if v146 != 0 {
		v119 = v119 + v148
		v120 = v120 + v148
		goto L26
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v123-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v134 = v123 | int32(32)
	goto L32
L31:
	;
	v134 = v123
	goto L32
L32:
	;
	if base.Ui32((v124-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = v124 | int32(32)
	goto L35
L34:
	;
	v143 = v124
	goto L35
L35:
	;
	if v134 == v143 {
		v146 = v134
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v157 = v134 - v143
	goto L25
L37:
	;
	goto L27
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v161 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v255 = int32(0)
	if v167 < v255 {
		v319 = v255
		goto L21
	} else {
		goto L64
	}
L40:
	;
	v166 = v161
	v167 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v237 = int32(0)
	v240 = F_errstart(m, int32(19), v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L60
	}
L43:
	;
	v177 = v114
	v178 = v166
	goto L46
L44:
	;
	goto L42
L45:
	;
	if v215 == int32(0) {
		goto L39
	} else {
		goto L58
	}
L46:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v181 == v182 {
		v204 = v181
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v215 = int32(0)
	goto L45
L48:
	;
	v206 = int32(1)
	if v204 != 0 {
		v177 = v177 + v206
		v178 = v178 + v206
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v181-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v192 = v181 | int32(32)
	goto L52
L51:
	;
	v192 = v181
	goto L52
L52:
	;
	if base.Ui32((v182-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v201 = v182 | int32(32)
	goto L55
L54:
	;
	v201 = v182
	goto L55
L55:
	;
	if v192 == v201 {
		v204 = v192
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v215 = v192 - v201
	goto L45
L57:
	;
	goto L47
L58:
	;
	v219 = v167 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l2+v219<<(uint(int32(3))%32))))
	if v223 != 0 {
		v166 = v223
		v167 = v219
		goto L43
	} else {
		goto L59
	}
L59:
	;
	goto L44
L60:
	;
	if v240 == int32(0) {
		v319 = v237
		goto L21
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v114
	F_errmsg(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_1), v16+int32(16))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_2), int32(273), int32(_a_F_stats_fill_fcinfo_from_arg_pairs_3))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v319 = v237
	goto L21
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258+v106<<(uint(int32(2))%32))))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l2+v167<<(uint(int32(3))%32))+4))
	if v262 != v266 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v270 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v291+v106<<(uint(int32(3))%32))))
	v298 = l1 + int32(24) + v167<<(uint(int32(4))%32)
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+8)) = uint8(v299)
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v295
	goto L22
L68:
	;
	if v270 == int32(0) {
		v319 = v255
		goto L21
	} else {
		goto L69
	}
L69:
	;
	v274 = F_format_type_be(m, v262)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	v276 = F_format_type_be(m, v266)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v114
	F_errmsg(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_4), v16+int32(32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_2), int32(289), int32(_a_F_stats_fill_fcinfo_from_arg_pairs_5))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v319 = v255
	goto L21
L74:
	;
	goto L18
L75:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v81 | int32(1)
	F_errmsg(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_6), v16)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_2), int32(389), int32(_a_F_stats_fill_fcinfo_from_arg_pairs_7))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L7
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375+v81<<(uint(int32(2))%32))))
	v380 = F_format_type_be(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v383 = F_format_type_be(m, int32(25))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v81 | int32(1)
	F_errmsg(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_8), v16+int32(48))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_2), int32(396), int32(_a_F_stats_fill_fcinfo_from_arg_pairs_7))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_9), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	F_errhint(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_10), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_stats_fill_fcinfo_from_arg_pairs_2), int32(374), int32(_a_F_stats_fill_fcinfo_from_arg_pairs_7))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
