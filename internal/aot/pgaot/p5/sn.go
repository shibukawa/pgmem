package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sn_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	if l1 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v15 <= int32(0) {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v19 = v18 + v15
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
			if v22 == int32(91) {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v25 <= v15+int32(1) {
					F_appendStringInfoChar(m, v14, int32(44))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v36 = int32(44)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v41 = v39 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v43+v41))) = uint8(v45)
					return int32(0)
				}
			}
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		if v7 != int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v15 <= int32(0) {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v19 = v18 + v15
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
				if v22 == int32(91) {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					if v25 <= v15+int32(1) {
						F_appendStringInfoChar(m, v14, int32(44))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v36 = int32(44)
						*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v41 = v39 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v45 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v43+v41))) = uint8(v45)
						return int32(0)
					}
				}
			}
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v10)
			return int32(0)
		}
	}
}
func F_sn_object_end(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14015(m, l0, int32(125))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_sn_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 - v13
	if base.Ui32(v11) < base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v11
	goto L3
L2:
	;
	v16 = v14
	goto L3
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v16) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v192 = v11
	v193 = v10
	goto L6
L6:
	;
	if base.Ui32(v192) < base.Ui32(l2) {
		goto L54
	} else {
		goto L55
	}
L7:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v187 = v186 + v16
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v190 = v189 - v16
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v190
	v192 = v190
	v193 = v187
	goto L6
L8:
	;
	if v16 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v23 = v10 + v16
	if (v10^v13)&int32(3) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	base.MemoryCopy(m, v10, v13, v16)
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L7
L14:
	;
	if base.Ui32(v155) < base.Ui32(v23) {
		goto L48
	} else {
		goto L49
	}
L15:
	;
	if v10&int32(3) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(v23) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L18:
	;
	v59 = v23 & int32(-4)
	if base.Ui32(v23) < base.Ui32(int32(64)) {
		v109 = v53
		v110 = v54
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v53 = v13
	v54 = v10
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v16 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v53 = v13
	v54 = v10
	goto L18
L23:
	;
	goto L24
L24:
	;
	v36 = v13
	v37 = v10
	goto L25
L25:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v41)
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v37 + v43
	if v46&int32(3) == int32(0) {
		v53 = v44
		v54 = v46
		goto L18
	} else {
		goto L27
	}
L26:
	;
	v53 = v44
	v54 = v46
	goto L18
L27:
	;
	if base.Ui32(v46) < base.Ui32(v23) {
		v36 = v44
		v37 = v46
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if base.Ui32(v59) <= base.Ui32(v110) {
		v154 = v109
		v155 = v110
		goto L14
	} else {
		goto L35
	}
L30:
	;
	v63 = v59 + int32(-64)
	if base.Ui32(v63) < base.Ui32(v54) {
		v109 = v53
		v110 = v54
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v66 = v53
	v67 = v54
	goto L32
L32:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+24)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+32)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+36)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v66)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+40)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v66)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+48)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v66)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+52)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+56)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v66)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+60)) = v101
	v103 = int32(-64)
	v104 = v66 - v103
	v106 = v67 - v103
	if base.Ui32(v106) <= base.Ui32(v63) {
		v66 = v104
		v67 = v106
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v104
	v110 = v106
	goto L29
L34:
	;
	goto L33
L35:
	;
	v116 = v109
	v117 = v110
	goto L36
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v121
	v123 = int32(4)
	v124 = v116 + v123
	v126 = v117 + v123
	if base.Ui32(v126) < base.Ui32(v59) {
		v116 = v124
		v117 = v126
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v154 = v124
	v155 = v126
	goto L14
L38:
	;
	goto L37
L39:
	;
	v154 = v13
	v155 = v10
	goto L14
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v16) < base.Ui32(int32(4)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v154 = v13
	v155 = v10
	goto L14
L43:
	;
	goto L44
L44:
	;
	v135 = v13
	v136 = v10
	goto L45
L45:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v140)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)) = uint8(v142)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+2)) = uint8(v144)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+3)) = uint8(v146)
	v148 = int32(4)
	v149 = v135 + v148
	v151 = v136 + v148
	if base.Ui32(v151) <= base.Ui32(v23-int32(4)) {
		v135 = v149
		v136 = v151
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v154 = v149
	v155 = v151
	goto L14
L47:
	;
	goto L46
L48:
	;
	v161 = v154
	v162 = v155
	goto L51
L49:
	;
	goto L50
L50:
	;
	goto L7
L51:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
	v168 = int32(1)
	v171 = v162 + v168
	if v171 != v23 {
		v161 = v161 + v168
		v162 = v171
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	v195 = v192
	goto L56
L55:
	;
	v195 = l2
	goto L56
L56:
	;
	if v195 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v195) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v371 = v193
	goto L59
L59:
	;
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v372)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v374
	return l2
L60:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v366 = v365 + v195
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v368 - v195
	v371 = v366
	goto L59
L61:
	;
	if v195 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v202 = v193 + v195
	if (v193^l1)&int32(3) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	base.MemoryCopy(m, v193, l1, v195)
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L60
L67:
	;
	if base.Ui32(v334) < base.Ui32(v202) {
		goto L101
	} else {
		goto L102
	}
L68:
	;
	if v193&int32(3) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v202) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L71:
	;
	v238 = v202 & int32(-4)
	if base.Ui32(v202) < base.Ui32(int32(64)) {
		v288 = v232
		v289 = v233
		goto L82
	} else {
		goto L83
	}
L72:
	;
	v232 = l1
	v233 = v193
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v195 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v232 = l1
	v233 = v193
	goto L71
L76:
	;
	goto L77
L77:
	;
	v215 = l1
	v216 = v193
	goto L78
L78:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v220)
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v216 + v222
	if v225&int32(3) == int32(0) {
		v232 = v223
		v233 = v225
		goto L71
	} else {
		goto L80
	}
L79:
	;
	v232 = v223
	v233 = v225
	goto L71
L80:
	;
	if base.Ui32(v225) < base.Ui32(v202) {
		v215 = v223
		v216 = v225
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if base.Ui32(v238) <= base.Ui32(v289) {
		v333 = v288
		v334 = v289
		goto L67
	} else {
		goto L88
	}
L83:
	;
	v242 = v238 + int32(-64)
	if base.Ui32(v242) < base.Ui32(v233) {
		v288 = v232
		v289 = v233
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v245 = v232
	v246 = v233
	goto L85
L85:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+20)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v245)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v245)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+32)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v245)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v245)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+40)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v245)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+44)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v245)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+48)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v245)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+52)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v245)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+56)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v245)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+60)) = v280
	v282 = int32(-64)
	v283 = v245 - v282
	v285 = v246 - v282
	if base.Ui32(v285) <= base.Ui32(v242) {
		v245 = v283
		v246 = v285
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v288 = v283
	v289 = v285
	goto L82
L87:
	;
	goto L86
L88:
	;
	v295 = v288
	v296 = v289
	goto L89
L89:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v300
	v302 = int32(4)
	v303 = v295 + v302
	v305 = v296 + v302
	if base.Ui32(v305) < base.Ui32(v238) {
		v295 = v303
		v296 = v305
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v333 = v303
	v334 = v305
	goto L67
L91:
	;
	goto L90
L92:
	;
	v333 = l1
	v334 = v193
	goto L67
L93:
	;
	goto L94
L94:
	;
	if base.Ui32(v195) < base.Ui32(int32(4)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v333 = l1
	v334 = v193
	goto L67
L96:
	;
	goto L97
L97:
	;
	v314 = l1
	v315 = v193
	goto L98
L98:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	*(*uint8)(unsafe.Add(mBase, uint32(v315))) = uint8(v319)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)) = uint8(v321)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+2)) = uint8(v323)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+3)) = uint8(v325)
	v327 = int32(4)
	v328 = v314 + v327
	v330 = v315 + v327
	if base.Ui32(v330) <= base.Ui32(v202-int32(4)) {
		v314 = v328
		v315 = v330
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v333 = v328
	v334 = v330
	goto L67
L100:
	;
	goto L99
L101:
	;
	v340 = v333
	v341 = v334
	goto L104
L102:
	;
	goto L103
L103:
	;
	goto L60
L104:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v345)
	v347 = int32(1)
	v350 = v341 + v347
	if v350 != v202 {
		v340 = v340 + v347
		v341 = v350
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
}
