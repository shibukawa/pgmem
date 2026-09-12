package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_CacheInvalidateHeapTupleCommon(m, l0, l1, l2, int32(1597))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_PlanCacheObjectCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1162]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	if v254 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	if v13 == int32(4089216) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v13
	goto L4
L4:
	;
	v30 = v22 - int32(5)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v31 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v239 != int32(4089216) {
		v22 = v239
		goto L4
	} else {
		goto L64
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22-int32(96))))
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v22-int32(32))))
	if v74 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	switch v40 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v44 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22-int32(92))))
	if v47 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v44 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v44 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v51 != int32(6) {
		v66 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v66&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v58 = v56 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v58) {
		v66 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v58)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v129 = v22 - int32(12)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v130 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v77 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v80 = int32(0)
	if v80 < v77 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v83 = v77
	goto L27
L26:
	;
	v83 = v80
	goto L27
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v86 = int32(0)
	goto L28
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84+v86<<(uint(int32(2))%32))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v101 != l1 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L22
L30:
	;
	v115 = v86 + int32(1)
	if v115 != v83 {
		v86 = v115
		goto L28
	} else {
		goto L37
	}
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v103 != l2 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v105)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v22-int32(12))))
	if v109 == v105 {
		goto L22
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)) = uint8(v112)
	goto L22
L37:
	;
	goto L29
L38:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+10)))
	if v133 != int32(1) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v136 == int32(0) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v140 <= v139 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v148 = v139
	v151 = v130
	goto L42
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v148<<(uint(int32(2))%32))))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v159 != int32(6) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L6
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+80))
	if v162 == int32(0) {
		v207 = v151
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v221 = v151
	goto L46
L46:
	;
	v225 = v148 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v225 < v226 {
		v148 = v225
		v151 = v221
		goto L42
	} else {
		goto L63
	}
L47:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+10)))
	if v210 != int32(1) {
		goto L6
	} else {
		goto L62
	}
L48:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v165 <= int32(0) {
		v207 = v151
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v168 = int32(0)
	if v168 < v165 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v171 = v165
	goto L52
L51:
	;
	v171 = v168
	goto L52
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v174 = int32(0)
	goto L53
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172+v174<<(uint(int32(2))%32))))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v189 != l1 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v207 = v151
	goto L47
L55:
	;
	v197 = v174 + int32(1)
	if v197 != v171 {
		v174 = v197
		goto L53
	} else {
		goto L61
	}
L56:
	;
	if l2 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v191 != l2 {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v193 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+10)) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v207 = v195
	goto L47
L60:
	;
	goto L59
L61:
	;
	goto L54
L62:
	;
	v221 = v207
	goto L46
L63:
	;
	goto L43
L64:
	;
	goto L5
L65:
	;
	return
L66:
	;
	if v254 == int32(4089224) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v262 = v254
	goto L68
L68:
	;
	v271 = v262 - int32(16)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v272 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L65
L70:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v324 != int32(4089224) {
		v262 = v324
		goto L68
	} else {
		goto L86
	}
L71:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v262-int32(8))))
	if v277 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v280 <= int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v283 = int32(0)
	if v283 < v280 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v286 = v280
	goto L76
L75:
	;
	v286 = v283
	goto L76
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v289 = int32(0)
	goto L77
L77:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v287+v289<<(uint(int32(2))%32))))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v304 != l1 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L70
L79:
	;
	v311 = v289 + int32(1)
	if v311 != v286 {
		v289 = v311
		goto L77
	} else {
		goto L85
	}
L80:
	;
	if l2 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+8))
	if v306 != l2 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v308)
	goto L70
L84:
	;
	goto L83
L85:
	;
	goto L78
L86:
	;
	goto L69
}
func F_PlanCacheSysCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1162]))
	if v5 == int32(0) {
	} else {
		if v5 == int32(4089216) {
		} else {
			v10 = v5
			for {
				v14 = v10 - int32(5)
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v15 != int32(1) {
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(96))))
					if v20 != 0 {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 - int32(137) {
						case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
							v28 = int32(1)
						default:
							v28 = int32(0)
						}
						if v28 != 0 {
							v56 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v56)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(12))))
							if v60 == v56 {
							} else {
								v63 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)) = uint8(v63)
							}
						} else {
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(92))))
						if v31 == int32(0) {
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v35 != int32(6) {
								v50 = int32(1)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
								v42 = v40 - int32(201)
								if base.Ui32(int32(41)) < base.Ui32(v42) {
									v50 = int32(0)
								} else {
									v50 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v42)) % 64)))
								}
							}
							if v50&int32(1) == int32(0) {
							} else {
								v56 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v56)
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(12))))
								if v60 == v56 {
								} else {
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)) = uint8(v63)
								}
							}
						}
					}
				}
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v67 != int32(4089216) {
					v10 = v67
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v74 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	if v74 == int32(0) {
	} else {
		if v74 == int32(4089224) {
		} else {
			v79 = v74
			for {
				v84 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v79-int32(16)))) = uint8(v84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
				if v86 != int32(4089224) {
					v79 = v86
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
func F_cache_locale_time(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	v1 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(3104)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1102])))
	if v23 == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L186
	}
L2:
	;
	v28 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v20 + int32(3104)
	return
L5:
	;
	return
L6:
	;
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	F_errmsg_internal(m, int32(701136), v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(44)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	v47 = int32(0)
	v51 = m.G0
	v53 = v51 - int32(32)
	m.G0 = v53
	v58 = v47
	goto L15
L10:
	;
	F_errfinish(m, int32(490141), int32(744), int32(368901))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v181 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L13:
	;
	m.G0 = v53 + int32(32)
	goto L12
L14:
	;
	v181 = int32(0)
	goto L13
L15:
	;
	goto L18
L16:
	;
	v86 = F___loc_is_allocated(m, v47)
	mBase = m.M
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+int32(8)+v58<<(uint(int32(2))%32)))) = v77
	if v77 == int32(-1) {
		goto L14
	} else {
		goto L24
	}
L18:
	;
	if int32(1)<<(uint(v58)%32)&int32(2147483647) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v76 = v46
	goto L23
L22:
	;
	v76 = int32(731167)
	goto L23
L23:
	;
	v77 = F___get_locale(m, v58, v76)
	mBase = m.M
	goto L17
L24:
	;
	v83 = v58 + int32(1)
	if v83 != int32(6) {
		v58 = v83
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	v89 = int32(4064040)
	v94 = F_memcmp(m, v53+int32(8), v89, int32(24))
	mBase = m.M
	if v94 == int32(0) {
		v181 = v89
		goto L13
	} else {
		goto L29
	}
L27:
	;
	v162 = v47
	goto L28
L28:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v170
	v181 = v162
	goto L13
L29:
	;
	v97 = int32(4064064)
	v102 = F_memcmp(m, v53+int32(8), v97, int32(24))
	mBase = m.M
	if v102 == int32(0) {
		v181 = v97
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v105 = int32(0)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1104])))
	if v107 == v105 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = v105
	goto L34
L32:
	;
	goto L33
L33:
	;
	v140 = int32(4634628)
	v145 = F_memcmp(m, v53+int32(8), v140, int32(24))
	mBase = m.M
	if v145 == int32(0) {
		v181 = v140
		goto L13
	} else {
		goto L37
	}
L34:
	;
	v121 = F___get_locale(m, v113, int32(731167))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[1105]))) = v121
	v124 = v113 + int32(1)
	if v124 != int32(6) {
		v113 = v124
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1104])) = uint8(v128)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[1105]))
	*(*int32)(unsafe.Add(mBase, _consts[1106])) = v132
	goto L33
L36:
	;
	goto L35
L37:
	;
	v148 = int32(4634652)
	v153 = F_memcmp(m, v53+int32(8), v148, int32(24))
	mBase = m.M
	if v153 == int32(0) {
		v181 = v148
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v157 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v157 == int32(0) {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v162 = v157
	goto L28
L40:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	F_report_newlocale_failure(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v194 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v194
	v200 = F___gmtime_r(m, v20+int32(56), v20+int32(12))
	mBase = m.M
	v202 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = v202
	v210 = F___strftime_l(m, v20-int32(-64), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v216 = F___strftime_l(m, v20+int32(144), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(1)
	v224 = F___strftime_l(m, v20+int32(224), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v230 = F___strftime_l(m, v20+int32(304), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(2)
	v238 = F___strftime_l(m, v20+int32(384), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v244 = F___strftime_l(m, v20+int32(464), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(3)
	v252 = F___strftime_l(m, v20+int32(544), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v258 = F___strftime_l(m, v20+int32(624), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(4)
	v266 = F___strftime_l(m, v20+int32(704), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v272 = F___strftime_l(m, v20+int32(784), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(5)
	v280 = F___strftime_l(m, v20+int32(864), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v286 = F___strftime_l(m, v20+int32(944), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = int32(6)
	v294 = F___strftime_l(m, v20+int32(1024), int32(80), int32(498227), v200, v181)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v296 = int32(0)
	v324 = F___strftime_l(m, v20+int32(1104), int32(80), int32(535732), v200, v181)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v326 = int32(0)
	v348 = v1
	v349 = base.B2i32(v210 == v296) | (base.B2i32(v216 == v296) | (base.B2i32(v224 == v296) | (base.B2i32(v230 == v296) | (base.B2i32(v238 == v296) | (base.B2i32(v244 == v296) | (base.B2i32(v252 == v296) | (base.B2i32(v258 == v296) | (base.B2i32(v266 == v296) | (base.B2i32(v272 == v296) | (base.B2i32(v280 == v296) | (base.B2i32(v286 == v296) | (base.B2i32(v324 == v326) | base.B2i32(v294 == v326)))))))))))))
	v351 = v20 + int32(1184)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v348
	v367 = F___strftime_l(m, v351, int32(80), int32(494804), v200, v181)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	v387 = F___loc_is_allocated(m, v181)
	mBase = m.M
	if v387 != 0 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v369 = int32(80)
	v373 = F___strftime_l(m, v351+v369, v369, int32(535305), v200, v181)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v375 = int32(0)
	v380 = base.B2i32(v373 == v375) | base.B2i32(v367 == v375) | v349
	v384 = v348 + int32(1)
	if v384 != int32(12) {
		v348 = v384
		v349 = v380
		v351 = v351 + int32(160)
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	if v380&int32(1) != 0 {
		goto L1
	} else {
		goto L67
	}
L64:
	;
	F_emscripten_builtin_free(m, v181)
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	v394 = F_pg_get_encoding_from_locale(m, v392, int32(1))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v396 = int32(0)
	if v396 < v394 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v399 = v394
	goto L71
L70:
	;
	v399 = v396
	goto L71
L71:
	;
	v403 = v20 - int32(-64)
	v407 = int32(0)
	goto L72
L72:
	;
	if v403&int32(3) == int32(0) {
		v443 = v403
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v575 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1107])) = v575
	*(*int32)(unsafe.Add(mBase, _consts[1108])) = v575
	v581 = v569
	v585 = v575
	goto L129
L74:
	;
	v477 = F_pg_any_to_server(m, v403, v476, v399)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L5
	} else {
		goto L91
	}
L75:
	;
	v476 = v468 - v403
	goto L74
L76:
	;
	v447 = v443
	goto L85
L77:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v427 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v476 = int32(0)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v432 = v403
	goto L81
L81:
	;
	v436 = v432 + int32(1)
	if v436&int32(3) == int32(0) {
		v443 = v436
		goto L76
	} else {
		goto L83
	}
L82:
	;
	v468 = v436
	goto L75
L83:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v441 != 0 {
		v432 = v436
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v456 = int32(-2139062144)
	if (int32(16843008)-v453|v453)&v456 == v456 {
		v447 = v447 + int32(4)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v462 = v447
	goto L88
L87:
	;
	goto L86
L88:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v466 != 0 {
		v462 = v462 + int32(1)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v468 = v462
	goto L75
L90:
	;
	goto L89
L91:
	;
	v480 = v407 << (uint(int32(2)) % 32)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v480)+uint32(_consts[1109])))
	v485 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v486 = F_MemoryContextStrdup(m, v485, v477)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v480)+uint32(_consts[1109]))) = v486
	if v483 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_pfree(m, v483)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L5
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v403 != v477 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	F_pfree(m, v477)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v495 = v403 + int32(80)
	if v495&int32(3) == int32(0) {
		v519 = v495
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L99
L101:
	;
	v553 = F_pg_any_to_server(m, v495, v552, v399)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L118
	}
L102:
	;
	v552 = v544 - v495
	goto L101
L103:
	;
	v523 = v519
	goto L112
L104:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	if v503 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v552 = int32(0)
	goto L101
L106:
	;
	goto L107
L107:
	;
	v508 = v495
	goto L108
L108:
	;
	v512 = v508 + int32(1)
	if v512&int32(3) == int32(0) {
		v519 = v512
		goto L103
	} else {
		goto L110
	}
L109:
	;
	v544 = v512
	goto L102
L110:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	if v517 != 0 {
		v508 = v512
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v532 = int32(-2139062144)
	if (int32(16843008)-v529|v529)&v532 == v532 {
		v523 = v523 + int32(4)
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v538 = v523
	goto L115
L114:
	;
	goto L113
L115:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	if v542 != 0 {
		v538 = v538 + int32(1)
		goto L115
	} else {
		goto L117
	}
L116:
	;
	v544 = v538
	goto L102
L117:
	;
	goto L116
L118:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v480)+uint32(_consts[1110])))
	v559 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v560 = F_MemoryContextStrdup(m, v559, v553)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v480)+uint32(_consts[1110]))) = v560
	if v557 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_pfree(m, v557)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if v553 != v495 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	F_pfree(m, v553)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L5
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v569 = v403 + int32(160)
	v571 = v407 + int32(1)
	if v571 != int32(7) {
		v403 = v569
		v407 = v571
		goto L72
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	goto L73
L129:
	;
	if v581&int32(3) == int32(0) {
		v621 = v581
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v753 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1102])) = uint8(v753)
	v756 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1111])) = v756
	*(*int32)(unsafe.Add(mBase, _consts[1112])) = v756
	goto L4
L131:
	;
	v655 = F_pg_any_to_server(m, v581, v654, v399)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L5
	} else {
		goto L148
	}
L132:
	;
	v654 = v646 - v581
	goto L131
L133:
	;
	v625 = v621
	goto L142
L134:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	if v605 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v654 = int32(0)
	goto L131
L136:
	;
	goto L137
L137:
	;
	v610 = v581
	goto L138
L138:
	;
	v614 = v610 + int32(1)
	if v614&int32(3) == int32(0) {
		v621 = v614
		goto L133
	} else {
		goto L140
	}
L139:
	;
	v646 = v614
	goto L132
L140:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v619 != 0 {
		v610 = v614
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625)))
	v634 = int32(-2139062144)
	if (int32(16843008)-v631|v631)&v634 == v634 {
		v625 = v625 + int32(4)
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v640 = v625
	goto L145
L144:
	;
	goto L143
L145:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v644 != 0 {
		v640 = v640 + int32(1)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v646 = v640
	goto L132
L147:
	;
	goto L146
L148:
	;
	v658 = v585 << (uint(int32(2)) % 32)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v658)+uint32(_consts[1113])))
	v663 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v664 = F_MemoryContextStrdup(m, v663, v655)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L5
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+uint32(_consts[1113]))) = v664
	if v661 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_pfree(m, v661)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	if v581 != v655 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	F_pfree(m, v655)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v673 = v581 + int32(80)
	if v673&int32(3) == int32(0) {
		v697 = v673
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L156
L158:
	;
	v731 = F_pg_any_to_server(m, v673, v730, v399)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L175
	}
L159:
	;
	v730 = v722 - v673
	goto L158
L160:
	;
	v701 = v697
	goto L169
L161:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	if v681 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v730 = int32(0)
	goto L158
L163:
	;
	goto L164
L164:
	;
	v686 = v673
	goto L165
L165:
	;
	v690 = v686 + int32(1)
	if v690&int32(3) == int32(0) {
		v697 = v690
		goto L160
	} else {
		goto L167
	}
L166:
	;
	v722 = v690
	goto L159
L167:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if v695 != 0 {
		v686 = v690
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v710 = int32(-2139062144)
	if (int32(16843008)-v707|v707)&v710 == v710 {
		v701 = v701 + int32(4)
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v716 = v701
	goto L172
L171:
	;
	goto L170
L172:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	if v720 != 0 {
		v716 = v716 + int32(1)
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v722 = v716
	goto L159
L174:
	;
	goto L173
L175:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v658)+uint32(_consts[1114])))
	v737 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v738 = F_MemoryContextStrdup(m, v737, v731)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+uint32(_consts[1114]))) = v738
	if v735 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	F_pfree(m, v735)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L5
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v731 != v673 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L179
L181:
	;
	F_pfree(m, v731)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L5
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v749 = v585 + int32(1)
	if v749 != int32(12) {
		v581 = v581 + int32(160)
		v585 = v749
		goto L129
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	goto L130
L186:
	;
	F_errmsg_internal(m, int32(446767), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(490141), int32(810), int32(368901))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cache_reduce_memory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var __phi178 int32
	_ = __phi178
	var v180 int32
	_ = v180
	var __phi180 int32
	_ = __phi180
	var v182 int32
	_ = v182
	var __phi182 int32
	_ = __phi182
	var v183 int32
	_ = v183
	var __phi183 int32
	_ = __phi183
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v247 int32
	_ = v247
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if base.Ui64(v17) < base.Ui64(v16) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v16
	goto L3
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v22 = l0 + int32(180)
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v22
	goto L6
L6:
	;
	v35 = int32(1)
	v36 = v23
	v39 = int64(0)
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L13
	} else {
		goto L46
	}
L8:
	;
	if v22 != v36 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v252 + v251
	return v247 & int32(1)
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v43 = v36 - int32(4)
	F_prepare_probe_slot(m, l0, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v247 = v35
	v251 = v39
	goto L12
L12:
	;
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v49 = F_MemoizeHash_hash(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v53 = v49 & v52
	v56 = v51 + v53<<(uint(int32(4))%32)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+12)))
	if v57 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v62 = v56
	v63 = v53
	v65 = v51
	v66 = v52
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v75 == v49 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v91 != v43 {
		goto L7
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v78 = F_MemoizeHash_equal(m, v48, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	v82 = v65
	v83 = v66
	goto L22
L22:
	;
	v86 = v83 & (v63 + int32(1))
	v89 = v82 + v86<<(uint(int32(4))%32)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v90 != 0 {
		v62 = v89
		v63 = v86
		v65 = v82
		v66 = v83
		goto L17
	} else {
		goto L25
	}
L23:
	;
	if v78 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v82 = v81
	v83 = v80
	goto L22
L25:
	;
	goto L7
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v96
	v98 = int64(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v99 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v103 = v99
	v113 = v98
	goto L30
L28:
	;
	v139 = v98
	goto L29
L29:
	;
	v142 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v142
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+13)) = uint8(v142)
	v146 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v147 = v146 - v139
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v147 - base.I64_extend_i32_u(v151+int32(28))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v158 - v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v165 = int32(4)
	v169 = v163 & ((v62-v162)>>(uint(v165)%32) + v159)
	v172 = v162 + v169<<(uint(v165)%32)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+12)))
	if v173 != v159 {
		v212 = v62
		goto L35
	} else {
		goto L36
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_pfree(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L13
	} else {
		goto L32
	}
L31:
	;
	v139 = v125
	goto L29
L32:
	;
	F_pfree(m, v103)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v125 = v113 + base.I64_extend_i32_u(v117+int32(8))
	if v115 != 0 {
		v103 = v115
		v113 = v125
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v224 = base.B2i32(l1 != v43) & v35
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+12)) = uint8(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	F_pfree(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L43
	}
L36:
	;
	__phi178 = v62
	__phi180 = v169
	__phi182 = v163
	__phi183 = v172
	v178 = __phi178
	v180 = __phi180
	v182 = __phi182
	v183 = __phi183
	goto L37
L37:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	if v180 == v191&v182 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v212 = v183
	goto L35
L39:
	;
	v212 = v178
	goto L35
L40:
	;
	goto L41
L41:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, uint32(v178))) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v178)+8)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v200 = int32(1)
	v202 = v199 & (v180 + v200)
	v205 = v198 + v202<<(uint(int32(4))%32)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+12)))
	if v206 == v200 {
		__phi178 = v183
		__phi180 = v202
		__phi182 = v199
		__phi183 = v205
		v178 = __phi178
		v180 = __phi180
		v182 = __phi182
		v183 = __phi183
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	F_pfree(m, v91)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v233 = v39 + int64(1)
	v234 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v235) < base.Ui64(v234) {
		v35 = v224
		v36 = v41
		v39 = v233
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v247 = v224
	v251 = v233
	goto L12
L46:
	;
	F_errmsg_internal(m, int32(12392), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(489424), int32(484), int32(13612))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
