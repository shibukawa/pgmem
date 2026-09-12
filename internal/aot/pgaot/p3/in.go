package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetupLockInTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v21 = F_hash_search_with_hash_value(m, v17, l2, l3, int32(3), v14+int32(23))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		if v21 == int32(0) {
			v224 = int32(0)
			m.G0 = v14 + int32(32)
			return v224
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
			if v27 == int32(0) {
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = int64(0)
				v39 = v21 + int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v39
				v43 = v21 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v43
				v46 = int32(40)
				v49 = v21 + int32(44)
				if v49&int32(3) == v30 {
					v55 = v21 + int32(84)
					v57 = v21 + int32(48)
					if base.Ui32(v57) < base.Ui32(v55) {
						v59 = v55
					} else {
						v59 = v57
					}
					v67 = (v59-v21-int32(45))&int32(-4) + int32(4)
				} else {
					v67 = v46
				}
				v71 = F__emscripten_memset_bulkmem(m, v49, base.I32_extend8_s(int32(0)), v67)
				mBase = m.M
				v73 = v21 + int32(88)
				if v73&int32(3) == int32(0) {
					v79 = v21 + int32(128)
					v81 = v21 + int32(92)
					if base.Ui32(v81) < base.Ui32(v79) {
						v83 = v79
					} else {
						v83 = v81
					}
					v91 = (v83-v21-int32(89))&int32(-4) + int32(4)
				} else {
					v91 = v46
				}
				v95 = F__emscripten_memset_bulkmem(m, v73, base.I32_extend8_s(int32(0)), v91)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l1
			v103 = *(*int32)(unsafe.Add(mBase, _consts[847]))
			v112 = F_hash_search_with_hash_value(m, v103, v14+int32(24), l1<<(uint(int32(4))%32)^l3, int32(3), v14+int32(23))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				if v112 == int32(0) {
					v116 = int32(0)
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
					if v117 != 0 {
						v224 = v116
						m.G0 = v14 + int32(32)
						return v224
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, _consts[43]))
						v122 = F_hash_search_with_hash_value(m, v119, v21, l3, int32(2), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							if v122 != 0 {
								v224 = v116
								m.G0 = v14 + int32(32)
								return v224
							} else {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(429716), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(480813), int32(1358), int32(383151))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
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
				} else {
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
					if v137 == int32(0) {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+616))
						*(*int64)(unsafe.Add(mBase, uint32(v112)+12)) = int64(0)
						if v140 != 0 {
							v143 = v140
						} else {
							v143 = l1
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v143
						v146 = v112 + int32(20)
						v148 = v21 + int32(24)
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
						if v151 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v148
							*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v148
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = v148
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v157
						*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v148))) = v146
						v162 = v112 + int32(28)
						v165 = l1 + l3&int32(15)<<(uint(int32(3))%32)
						v167 = v165 + int32(148)
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+152))
						if v168 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v167
							*(*int32)(unsafe.Add(mBase, uint32(v167))) = v167
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+32)) = v167
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
						*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v162
						*(*int32)(unsafe.Add(mBase, uint32(v167))) = v162
					} else {
					}
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
					v184 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v183 + v184
					v188 = l4 << (uint(int32(2)) % 32)
					v191 = v21 + v188 + int32(44)
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
					*(*int32)(unsafe.Add(mBase, uint32(v191))) = v192 + v184
					v196 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
					if int32(base.Ui32(v196)>>(uint(l4)%32))&v184 == int32(0) {
						v224 = v112
						m.G0 = v14 + int32(32)
						return v224
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int32(0)
						} else {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v206+v188)))
							v209 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v210
							*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = v209
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v208
							F_errmsg_internal(m, int32(416919), v14)
							mBase = m.M
							v216 = m.ExcPending
							if v216 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(480813), int32(1449), int32(383151))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
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
		}
	}
}
func F_find_in_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L56
	} else {
		goto L127
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L56
	} else {
		goto L123
	}
L3:
	;
	m.G0 = v11 - int32(-64)
	return v391
L4:
	;
	if l0&int32(3) == int32(0) {
		v37 = l0
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v391 = int32(0)
	goto L3
L7:
	;
	v74 = l1
	goto L24
L8:
	;
	v70 = v62 - l0
	goto L7
L9:
	;
	v41 = v37
	goto L18
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v21 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v70 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v26 = l0
	goto L14
L14:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v62 = v30
	goto L8
L16:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v56 = v41
	goto L21
L20:
	;
	goto L19
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v62 = v56
	goto L8
L23:
	;
	goto L22
L24:
	;
	v82 = v74
	goto L27
L25:
	;
	goto L6
L26:
	;
	if v92 == v74 {
		goto L2
	} else {
		goto L34
	}
L27:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v84 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v92 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v84 == int32(58) {
		v92 = v82
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v82 = v82 + int32(1)
	goto L27
L34:
	;
	if v92 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v156 = v154 + int32(1)
	v157 = F_palloc(m, v156)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L56
	} else {
		goto L57
	}
L36:
	;
	if v74&int32(3) == int32(0) {
		v119 = v74
		goto L41
	} else {
		goto L42
	}
L37:
	;
	goto L38
L38:
	;
	v154 = v92 - v74
	goto L35
L39:
	;
	v154 = v152
	goto L35
L40:
	;
	v152 = v144 - v74
	goto L39
L41:
	;
	v123 = v119
	goto L50
L42:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v103 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = int32(0)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v108 = v74
	goto L46
L46:
	;
	v112 = v108 + int32(1)
	if v112&int32(3) == int32(0) {
		v119 = v112
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v144 = v112
	goto L40
L48:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v117 != 0 {
		v108 = v112
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v132 = int32(-2139062144)
	if (int32(16843008)-v129|v129)&v132 == v132 {
		v123 = v123 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v138 = v123
	goto L53
L52:
	;
	goto L51
L53:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v142 != 0 {
		v138 = v138 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v144 = v138
	goto L40
L55:
	;
	goto L54
L56:
	;
	return int32(0)
L57:
	;
	if v156 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v277 = F_substitute_path_macro(m, v157, int32(205606), int32(4452096))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L56
	} else {
		goto L90
	}
L59:
	;
	v272 = F_strlen(m, v268)
	mBase = m.M
	goto L58
L60:
	;
	v268 = v74
	goto L59
L61:
	;
	goto L62
L62:
	;
	v166 = v156 - int32(1)
	if (v157^v74)&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v265)
	v268 = v261
	goto L59
L64:
	;
	v246 = v241
	v247 = v242
	v248 = v243
	goto L86
L65:
	;
	if v236 == int32(0) {
		v261 = v234
		v262 = v235
		goto L63
	} else {
		goto L85
	}
L66:
	;
	v234 = v74
	v235 = v157
	v236 = v166
	goto L65
L67:
	;
	goto L68
L68:
	;
	v170 = int32(0)
	if v74&int32(3) == v170 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v203 == int32(0) {
		v261 = v200
		v262 = v201
		goto L63
	} else {
		goto L78
	}
L70:
	;
	v200 = v74
	v201 = v157
	v202 = v166
	v203 = base.B2i32(v166 != v170)
	goto L69
L71:
	;
	if v166 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v179 = v74
	v180 = v157
	v181 = v166
	goto L73
L73:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v183)
	if v183 == int32(0) {
		v241 = v179
		v242 = v180
		v243 = v181
		goto L64
	} else {
		goto L75
	}
L74:
	;
	v200 = v194
	v201 = v188
	v202 = v190
	v203 = v192
	goto L69
L75:
	;
	v187 = int32(1)
	v188 = v180 + v187
	v190 = v181 - v187
	v191 = int32(0)
	v192 = base.B2i32(v190 != v191)
	v194 = v179 + v187
	if v194&int32(3) == v191 {
		v200 = v194
		v201 = v188
		v202 = v190
		v203 = v192
		goto L69
	} else {
		goto L76
	}
L76:
	;
	if v190 != 0 {
		v179 = v194
		v180 = v188
		v181 = v190
		goto L73
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v206 == int32(0) {
		v234 = v200
		v235 = v201
		v236 = v202
		goto L65
	} else {
		goto L79
	}
L79:
	;
	if base.Ui32(v202) < base.Ui32(int32(4)) {
		v234 = v200
		v235 = v201
		v236 = v202
		goto L65
	} else {
		goto L80
	}
L80:
	;
	v212 = v200
	v213 = v201
	v214 = v202
	goto L81
L81:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v220 = int32(-2139062144)
	if (int32(16843008)-v217|v217)&v220 != v220 {
		v241 = v212
		v242 = v213
		v243 = v214
		goto L64
	} else {
		goto L83
	}
L82:
	;
	v234 = v228
	v235 = v226
	v236 = v230
	goto L65
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v217
	v225 = int32(4)
	v226 = v213 + v225
	v228 = v212 + v225
	v230 = v214 - v225
	if base.Ui32(int32(3)) < base.Ui32(v230) {
		v212 = v228
		v213 = v226
		v214 = v230
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v241 = v234
	v242 = v235
	v243 = v236
	goto L64
L86:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v250)
	if v250 == int32(0) {
		v261 = v246
		v262 = v247
		goto L63
	} else {
		goto L88
	}
L87:
	;
	v261 = v257
	v262 = v255
	goto L63
L88:
	;
	v254 = int32(1)
	v255 = v247 + v254
	v257 = v246 + v254
	v259 = v248 - v254
	if v259 != 0 {
		v246 = v257
		v247 = v255
		v248 = v259
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	F_pfree(m, v157)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L56
	} else {
		goto L91
	}
L91:
	;
	F_canonicalize_path_enc(m, v277)
	mBase = m.M
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v282 != int32(47) {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v277&int32(3) == int32(0) {
		v308 = v277
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v343 = F_palloc(m, v341+(v70+int32(2)))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L56
	} else {
		goto L110
	}
L94:
	;
	v341 = v333 - v277
	goto L93
L95:
	;
	v312 = v308
	goto L104
L96:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v292 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v341 = int32(0)
	goto L93
L98:
	;
	goto L99
L99:
	;
	v297 = v277
	goto L100
L100:
	;
	v301 = v297 + int32(1)
	if v301&int32(3) == int32(0) {
		v308 = v301
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v333 = v301
	goto L94
L102:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v306 != 0 {
		v297 = v301
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v321 = int32(-2139062144)
	if (int32(16843008)-v318|v318)&v321 == v321 {
		v312 = v312 + int32(4)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v327 = v312
	goto L107
L106:
	;
	goto L105
L107:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v331 != 0 {
		v327 = v327 + int32(1)
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v333 = v327
	goto L94
L109:
	;
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v277
	v350 = F_pg_sprintf(m, v343, int32(169664), v9+int32(-32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L56
	} else {
		goto L111
	}
L111:
	;
	F_pfree(m, v277)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L56
	} else {
		goto L112
	}
L112:
	;
	v356 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L56
	} else {
		goto L113
	}
L113:
	;
	if v356 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(310103)
	F_errmsg_internal(m, int32(677989), v9+int32(-48))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L56
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v371 = F_pg_file_exists(m, v343)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L56
	} else {
		goto L119
	}
L117:
	;
	F_errfinish(m, int32(478373), int32(631), int32(310103))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L56
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	if v371 != 0 {
		v391 = v343
		goto L3
	} else {
		goto L120
	}
L120:
	;
	F_pfree(m, v343)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L56
	} else {
		goto L121
	}
L121:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v154))))
	if v377 != 0 {
		v74 = v74 + v156
		goto L24
	} else {
		goto L122
	}
L122:
	;
	goto L25
L123:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L56
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(309981)
	F_errmsg(m, int32(666247), v11)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L56
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(478373), int32(606), int32(310103))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L56
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L56
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(309981)
	F_errmsg(m, int32(310370), v9+int32(-16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L56
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(478373), int32(625), int32(310103))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L56
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_in_grouping_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v13
	goto L2
L1:
	;
	return v113
L2:
	;
	if v14 <= v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v113 = int32(0)
	goto L1
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	if base.Ui32(v33) < base.Ui32(int32(192)) {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l3 < v90 {
		v113 = v91
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v37 = v23 + int32(1)
	if v37 == v14 {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v15))))
	v42 = v40 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v33) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v15))))
	v58 = v56 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v33) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v46 = v23 + int32(2)
	if v46 != v14 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v90 = v33<<(uint(int32(6))%32)&int32(1984) | v42
	v91 = int32(2)
	goto L7
L14:
	;
	goto L13
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v62))))
	v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(1835008) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
	v91 = int32(4)
	goto L7
L16:
	;
	v62 = v23 + int32(3)
	if v62 != v14 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v90 = v33<<(uint(int32(12))%32)&int32(61440) | v42<<(uint(int32(6))%32) | v58
	v91 = int32(3)
	goto L7
L19:
	;
	goto L18
L20:
	;
	v95 = v90 - l2
	if v95 < int32(0) {
		v113 = v91
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v95)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		v113 = v91
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v109 = v91 + v23
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
	if l4 != 0 {
		v23 = v109
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L3
}
func F_in_range_int8_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v13 != 0 {
			v14 = int64(0) - v7
		} else {
			v14 = v7
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v19 = v18 + v14
		if base.B2i32(v14 < int64(0)) != base.B2i32(v19 < v18) {
			v22 = int32(0)
			return base.B2i32(v13 != v22) ^ base.B2i32(v10 != v22)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
			if v10 != 0 {
				return base.B2i32(v29 <= v19)
			} else {
				return base.B2i32(v19 <= v29)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(242417), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(483251), int32(413), int32(530548))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_in_range_numeric_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v29 = F_pg_detoast_datum(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
				if v31 == int32(49152) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v366 = m.ExcPending
					if v366 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50593922))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(242417), int32(0))
							mBase = m.M
							v373 = m.ExcPending
							if v373 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483087), int32(2699), int32(474131))
								mBase = m.M
								v378 = m.ExcPending
								if v378 != 0 {
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
					if v31 == int32(61440) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50593922))
							mBase = m.M
							v369 = m.ExcPending
							if v369 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(242417), int32(0))
								mBase = m.M
								v373 = m.ExcPending
								if v373 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483087), int32(2699), int32(474131))
									mBase = m.M
									v378 = m.ExcPending
									if v378 != 0 {
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
						v36 = int32(49152)
						v37 = v31 & v36
						if v37 != v36 {
							if v37 != int32(32768) {
								v48 = v37
							} else {
								v48 = v31 << (uint(int32(1)) % 32) & int32(16384)
							}
						} else {
							v48 = v31 & int32(61440)
						}
						if v48 == int32(16384) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v366 = m.ExcPending
							if v366 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50593922))
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(242417), int32(0))
									mBase = m.M
									v373 = m.ExcPending
									if v373 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(483087), int32(2699), int32(474131))
										mBase = m.M
										v378 = m.ExcPending
										if v378 != 0 {
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
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
							v53 = base.I32_extend16_s(v52)
							v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
							if v59 == int32(49152) {
								v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v53 == int32(-16384))
								v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v345 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v348 = m.ExcPending
									if v348 != 0 {
										return int32(0)
									} else {
										v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v349 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v352 = m.ExcPending
											if v352 != 0 {
												return int32(0)
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									}
								} else {
									v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v349 != v26 {
										F_pfree(m, v26)
										mBase = m.M
										v352 = m.ExcPending
										if v352 != 0 {
											return int32(0)
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									} else {
										v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v353 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v356 = m.ExcPending
											if v356 != 0 {
												return int32(0)
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										} else {
											m.G0 = v18 + int32(96)
											return v344
										}
									}
								}
							} else {
								if v53 == int32(-16384) {
									v344 = base.B2i32(v51 != int32(0))
									v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v345 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v348 = m.ExcPending
										if v348 != 0 {
											return int32(0)
										} else {
											v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v349 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v352 = m.ExcPending
												if v352 != 0 {
													return int32(0)
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										}
									} else {
										v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v349 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v352 = m.ExcPending
											if v352 != 0 {
												return int32(0)
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v67 = base.I32_extend16_s(v59)
									v68 = base.I32_extend16_s(v31)
									if base.Ui32(int32(-16384)) <= base.Ui32(v68) {
										if v66 != 0 {
											if v53 != int32(-12288) {
												v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v67 == int32(-4096))
											} else {
												v344 = int32(1)
											}
										} else {
											if v53 != int32(-4096) {
												v344 = base.B2i32(v67 == int32(-12288)) | base.B2i32(v51 != int32(0))
											} else {
												v344 = int32(1)
											}
										}
										v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v345 != v21 {
											F_pfree(m, v21)
											mBase = m.M
											v348 = m.ExcPending
											if v348 != 0 {
												return int32(0)
											} else {
												v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if v349 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v352 = m.ExcPending
													if v352 != 0 {
														return int32(0)
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											}
										} else {
											v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v349 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v352 = m.ExcPending
												if v352 != 0 {
													return int32(0)
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										}
									} else {
										if base.Ui32(int32(-16384)) <= base.Ui32(v67) {
											v91 = int32(-12288)
											if v67 == v91 {
												v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v53 == v91)
											} else {
												v344 = base.B2i32(v53 == int32(-4096)) | base.B2i32(v51 != int32(0))
											}
											v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v345 != v21 {
												F_pfree(m, v21)
												mBase = m.M
												v348 = m.ExcPending
												if v348 != 0 {
													return int32(0)
												} else {
													v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v349 != v26 {
														F_pfree(m, v26)
														mBase = m.M
														v352 = m.ExcPending
														if v352 != 0 {
															return int32(0)
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												}
											} else {
												v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if v349 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v352 = m.ExcPending
													if v352 != 0 {
														return int32(0)
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v53) {
												v344 = base.B2i32(v53 == int32(-4096)) ^ base.B2i32(v51 != int32(0))
												v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v345 != v21 {
													F_pfree(m, v21)
													mBase = m.M
													v348 = m.ExcPending
													if v348 != 0 {
														return int32(0)
													} else {
														v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v349 != v26 {
															F_pfree(m, v26)
															mBase = m.M
															v352 = m.ExcPending
															if v352 != 0 {
																return int32(0)
															} else {
																v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																if v353 != v29 {
																	F_pfree(m, v29)
																	mBase = m.M
																	v356 = m.ExcPending
																	if v356 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
																	}
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															}
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													}
												} else {
													v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v349 != v26 {
														F_pfree(m, v26)
														mBase = m.M
														v352 = m.ExcPending
														if v352 != 0 {
															return int32(0)
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												}
											} else {
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
												v114 = base.B2i32(int32(0) <= v67)
												if int32(0) <= v67 {
													v115 = int32(-8)
												} else {
													v115 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(base.Ui32(int32(base.Ui32(v108)>>(uint(int32(2))%32))+v115) >> (uint(int32(1)) % 32))
												if int32(0) <= v67 {
													v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
													v130 = v120
												} else {
													v130 = v59<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v59&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v130
												v132 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v132
												v141 = base.B2i32(v67 < v132)
												if v67 < v132 {
													v142 = int32(base.Ui32(v59)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v142 = v59 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v142
												v149 = v59 & int32(49152)
												if v149 == int32(32768) {
													v152 = v59 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v152 = v149
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v152
												if v67 < v132 {
													v156 = int32(6)
												} else {
													v156 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v21 + v156
												v159 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
												v165 = base.B2i32(int32(0) <= v53)
												if int32(0) <= v53 {
													v166 = int32(-8)
												} else {
													v166 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(base.Ui32(int32(base.Ui32(v159)>>(uint(int32(2))%32))+v166) >> (uint(int32(1)) % 32))
												if int32(0) <= v53 {
													v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
													v181 = v171
												} else {
													v181 = v52<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v52&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v181
												v183 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v183
												v192 = base.B2i32(v53 < v183)
												if v53 < v183 {
													v193 = int32(base.Ui32(v52)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v193 = v52 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v193
												v200 = v52 & int32(49152)
												if v200 == int32(32768) {
													v203 = v52 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v203 = v200
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v203
												if v53 < v183 {
													v207 = int32(6)
												} else {
													v207 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v26 + v207
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
												v216 = base.B2i32(int32(0) <= v68)
												if int32(0) <= v68 {
													v217 = int32(-8)
												} else {
													v217 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(base.Ui32(int32(base.Ui32(v210)>>(uint(int32(2))%32))+v217) >> (uint(int32(1)) % 32))
												if int32(0) <= v68 {
													v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+6)))
													v232 = v222
												} else {
													v232 = v31<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v31&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v232
												if v37 != int32(49152) {
													if v37 != int32(32768) {
														v244 = v37
													} else {
														v244 = v31 << (uint(int32(1)) % 32) & int32(16384)
													}
												} else {
													v244 = v31 & int32(61440)
												}
												v245 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v245
												*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v245
												*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v244
												*(*int64)(unsafe.Add(mBase, uint32(v18))) = v245
												v252 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v252
												v261 = base.B2i32(v68 < v252)
												if v68 < v252 {
													v262 = int32(base.Ui32(v31)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v262 = v31 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v262
												if v68 < v252 {
													v266 = int32(6)
												} else {
													v266 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v29 + v266
												if v66 != 0 {
													F_sub_var(m, v18+int32(48), v18+int32(24), v18)
													mBase = m.M
													v274 = m.ExcPending
													if v274 != 0 {
														return int32(0)
													} else {
														v282 = v18 + int32(72)
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
														v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
														if v291 == int32(0) {
															if v290 == int32(0) {
																v327 = int32(0)
															} else {
																if v289 == int32(16384) {
																	v301 = int32(1)
																} else {
																	v301 = int32(-1)
																}
																v327 = v301
															}
														} else {
															v302 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
															if v290 == int32(0) {
																if v302 != 0 {
																	v307 = int32(-1)
																} else {
																	v307 = int32(1)
																}
																v327 = v307
															} else {
																v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
																v310 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
																v311 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
																if v302 == int32(0) {
																	if v289 == int32(16384) {
																		v327 = int32(1)
																	} else {
																		v317 = F_cmp_abs_common(m, v311, v291, v310, v309, v290, v308)
																		mBase = m.M
																		v327 = v317
																	}
																} else {
																	if v289 == int32(0) {
																		v327 = int32(-1)
																	} else {
																		v321 = F_cmp_abs_common(m, v309, v290, v308, v311, v291, v310)
																		mBase = m.M
																		v327 = v321
																	}
																}
															}
														}
														v328 = int32(0)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
														if v332 != 0 {
															F_pfree(m, v332)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																if v51 != 0 {
																	v335 = base.B2i32(v327 <= v328)
																} else {
																	v335 = base.B2i32(v328 <= v327)
																}
																v344 = v335
																v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																if v345 != v21 {
																	F_pfree(m, v21)
																	mBase = m.M
																	v348 = m.ExcPending
																	if v348 != 0 {
																		return int32(0)
																	} else {
																		v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		if v349 != v26 {
																			F_pfree(m, v26)
																			mBase = m.M
																			v352 = m.ExcPending
																			if v352 != 0 {
																				return int32(0)
																			} else {
																				v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																				if v353 != v29 {
																					F_pfree(m, v29)
																					mBase = m.M
																					v356 = m.ExcPending
																					if v356 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v18 + int32(96)
																						return v344
																					}
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			}
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	}
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															}
														} else {
															if v51 != 0 {
																v335 = base.B2i32(v327 <= v328)
															} else {
																v335 = base.B2i32(v328 <= v327)
															}
															v344 = v335
															v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															if v345 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v348 = m.ExcPending
																if v348 != 0 {
																	return int32(0)
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															} else {
																v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																if v349 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v352 = m.ExcPending
																	if v352 != 0 {
																		return int32(0)
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																} else {
																	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v353 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v356 = m.ExcPending
																		if v356 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
																	}
																}
															}
														}
													}
												} else {
													F_add_var(m, v18+int32(48), v18+int32(24), v18)
													mBase = m.M
													v280 = m.ExcPending
													if v280 != 0 {
														return int32(0)
													} else {
														v282 = v18 + int32(72)
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
														v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
														if v291 == int32(0) {
															if v290 == int32(0) {
																v327 = int32(0)
															} else {
																if v289 == int32(16384) {
																	v301 = int32(1)
																} else {
																	v301 = int32(-1)
																}
																v327 = v301
															}
														} else {
															v302 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
															if v290 == int32(0) {
																if v302 != 0 {
																	v307 = int32(-1)
																} else {
																	v307 = int32(1)
																}
																v327 = v307
															} else {
																v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
																v310 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
																v311 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
																if v302 == int32(0) {
																	if v289 == int32(16384) {
																		v327 = int32(1)
																	} else {
																		v317 = F_cmp_abs_common(m, v311, v291, v310, v309, v290, v308)
																		mBase = m.M
																		v327 = v317
																	}
																} else {
																	if v289 == int32(0) {
																		v327 = int32(-1)
																	} else {
																		v321 = F_cmp_abs_common(m, v309, v290, v308, v311, v291, v310)
																		mBase = m.M
																		v327 = v321
																	}
																}
															}
														}
														v328 = int32(0)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
														if v332 != 0 {
															F_pfree(m, v332)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																if v51 != 0 {
																	v335 = base.B2i32(v327 <= v328)
																} else {
																	v335 = base.B2i32(v328 <= v327)
																}
																v344 = v335
																v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																if v345 != v21 {
																	F_pfree(m, v21)
																	mBase = m.M
																	v348 = m.ExcPending
																	if v348 != 0 {
																		return int32(0)
																	} else {
																		v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		if v349 != v26 {
																			F_pfree(m, v26)
																			mBase = m.M
																			v352 = m.ExcPending
																			if v352 != 0 {
																				return int32(0)
																			} else {
																				v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																				if v353 != v29 {
																					F_pfree(m, v29)
																					mBase = m.M
																					v356 = m.ExcPending
																					if v356 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v18 + int32(96)
																						return v344
																					}
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			}
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	}
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															}
														} else {
															if v51 != 0 {
																v335 = base.B2i32(v327 <= v328)
															} else {
																v335 = base.B2i32(v328 <= v327)
															}
															v344 = v335
															v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															if v345 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v348 = m.ExcPending
																if v348 != 0 {
																	return int32(0)
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															} else {
																v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																if v349 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v352 = m.ExcPending
																	if v352 != 0 {
																		return int32(0)
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																} else {
																	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v353 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v356 = m.ExcPending
																		if v356 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
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
}
func F_in_range_time_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v15 != 0 {
			v45 = v12 - v7
			if v10 != 0 {
				return base.B2i32(v14 <= v45)
			} else {
				return base.B2i32(v45 <= v14)
			}
		} else {
			v19 = v7 + v12
			if base.B2i32(v7 < int64(0))^base.B2i32(v19 < v12) == int32(0) {
				v45 = v19
				if v10 != 0 {
					return base.B2i32(v14 <= v45)
				} else {
					return base.B2i32(v45 <= v14)
				}
			} else {
				return base.B2i32(v10 != int32(0))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(242417), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481749), int32(2180), int32(298227))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
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
func F_in_range_timetz_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	if int64(0) <= v8 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v15 != 0 {
			v45 = v14 - v8
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v48 = int64(1000000)
			v50 = base.I64_extend_i32_s(v46)*v48 + v45
			v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v56 = v51 + base.I64_extend_i32_s(v52)*v48
			if v11 != 0 {
				if v50 < v56 {
					return int32(0)
				} else {
					return base.B2i32(v52 <= v46) | base.B2i32(v56 < v50)
				}
			} else {
				if v50 < v56 {
					return int32(1)
				} else {
					return base.B2i32(v46 <= v52) & base.B2i32(v50 <= v56)
				}
			}
		} else {
			v19 = v8 + v14
			if base.B2i32(v8 < int64(0))^base.B2i32(v19 < v14) == int32(0) {
				v45 = v19
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				v48 = int64(1000000)
				v50 = base.I64_extend_i32_s(v46)*v48 + v45
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v56 = v51 + base.I64_extend_i32_s(v52)*v48
				if v11 != 0 {
					if v50 < v56 {
						return int32(0)
					} else {
						return base.B2i32(v52 <= v46) | base.B2i32(v56 < v50)
					}
				} else {
					if v50 < v56 {
						return int32(1)
					} else {
						return base.B2i32(v46 <= v52) & base.B2i32(v50 <= v56)
					}
				}
			} else {
				return base.B2i32(v11 != int32(0))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(242417), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481749), int32(2732), int32(297739))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
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
