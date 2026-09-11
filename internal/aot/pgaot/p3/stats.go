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
	v12 = l0 + l1<<(uint(int32(3))%32)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	if v13 != 0 {
		v63 = v9
		m.G0 = v7 + int32(16)
		return v63
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
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
						v44 = int32(75)
						v45 = int32(23432)
						F_errcode(m, int32(50856066))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[766])))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v53
							F_errmsg(m, v45, v7)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(463644), v44, int32(22382))
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
								v44 = int32(84)
								v45 = int32(147028)
								F_errcode(m, int32(50856066))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[766])))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v53
									F_errmsg(m, v45, v7)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(463644), v44, int32(22382))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v54 = int32(1)
	v61 = F_extract_variadic_args(m, l0, v14+int32(76), v14+int32(68), v14+int32(72))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v30 = int32(3)
	v32 = l1 + int32(20) + v23<<(uint(v30)%32)
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(0)
	v38 = v23 + v33
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2+v38<<(uint(v30)%32))))
	if v42 != 0 {
		v23 = v38
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
	if v61&int32(1) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(0) < v61 {
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
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L83
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L78
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L75
	}
L14:
	;
	v79 = v54
	v80 = int32(0)
	goto L17
L15:
	;
	v315 = v54
	goto L16
L16:
	;
	m.G0 = v14 + int32(80)
	return v315 & int32(1)
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v80))))
	if v87 == int32(1) {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v315 = v301
	goto L16
L19:
	;
	v91 = v80 << (uint(int32(2)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)))
	if v94 != int32(25) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v98 = v80 | int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v98))))
	if v100 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v308 = v80 + int32(2)
	if v308 < v61 {
		v79 = v301
		v80 = v308
		goto L17
	} else {
		goto L74
	}
L22:
	;
	v301 = v79
	goto L21
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101+v91)))
	v104 = F_text_to_cstring(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v109 = v104
	v110 = int32(255602)
	goto L26
L25:
	;
	if v147 == int32(0) {
		goto L22
	} else {
		goto L38
	}
L26:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v113 == v114 {
		v136 = v113
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v147 = int32(0)
	goto L25
L28:
	;
	v138 = int32(1)
	if v136 != 0 {
		v109 = v109 + v138
		v110 = v110 + v138
		goto L26
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v113-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v124 = v113 | int32(32)
	goto L32
L31:
	;
	v124 = v113
	goto L32
L32:
	;
	if base.Ui32((v114-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = v114 | int32(32)
	goto L35
L34:
	;
	v133 = v114
	goto L35
L35:
	;
	if v124 == v133 {
		v136 = v124
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v147 = v124 - v133
	goto L25
L37:
	;
	goto L27
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v151 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v241 = int32(0)
	if v156 < v241 {
		v301 = v241
		goto L21
	} else {
		goto L64
	}
L40:
	;
	v156 = int32(0)
	v157 = v151
	goto L43
L41:
	;
	goto L42
L42:
	;
	v223 = int32(0)
	v226 = F_errstart(m, int32(19), v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L60
	}
L43:
	;
	v165 = v104
	v166 = v157
	goto L46
L44:
	;
	goto L42
L45:
	;
	if v203 == int32(0) {
		goto L39
	} else {
		goto L58
	}
L46:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v169 == v170 {
		v192 = v169
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v203 = int32(0)
	goto L45
L48:
	;
	v194 = int32(1)
	if v192 != 0 {
		v165 = v165 + v194
		v166 = v166 + v194
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v169-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v180 = v169 | int32(32)
	goto L52
L51:
	;
	v180 = v169
	goto L52
L52:
	;
	if base.Ui32((v170-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v189 = v170 | int32(32)
	goto L55
L54:
	;
	v189 = v170
	goto L55
L55:
	;
	if v180 == v189 {
		v192 = v180
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v203 = v180 - v189
	goto L45
L57:
	;
	goto L47
L58:
	;
	v207 = v156 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2+v207<<(uint(int32(3))%32))))
	if v211 != 0 {
		v156 = v207
		v157 = v211
		goto L43
	} else {
		goto L59
	}
L59:
	;
	goto L44
L60:
	;
	if v226 == int32(0) {
		v301 = v223
		goto L21
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v104
	F_errmsg(m, int32(676272), v14+int32(16))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(463644), int32(260), int32(354880))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v301 = v223
	goto L21
L64:
	;
	v245 = v98 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245+v246)))
	v250 = v156 << (uint(int32(3)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l2+v250)+4))
	if v248 != v252 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v254 = int32(0)
	v257 = F_errstart(m, int32(19), v254)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v278+v245)))
	v281 = v250 + (l1 + int32(20))
	v282 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+4)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v280
	goto L22
L68:
	;
	if v257 == int32(0) {
		v301 = v254
		goto L21
	} else {
		goto L69
	}
L69:
	;
	v261 = F_format_type_be(m, v248)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	v263 = F_format_type_be(m, v252)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v104
	F_errmsg(m, int32(181232), v14+int32(32))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(463644), int32(276), int32(344331))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v301 = v254
	goto L21
L74:
	;
	goto L18
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v80 | int32(1)
	F_errmsg(m, int32(283691), v14)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(463644), int32(330), int32(121170))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v80<<(uint(int32(2))%32))))
	v351 = F_format_type_be(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	v354 = F_format_type_be(m, int32(25))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v80 | int32(1)
	F_errmsg(m, int32(181173), v14+int32(48))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(463644), int32(336), int32(121170))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
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
	F_errmsg(m, int32(121203), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(539275), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(463644), int32(316), int32(121170))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
