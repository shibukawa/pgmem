package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(1)
	v17 = v9 + v16
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v22 = v20 & v16
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v17
	goto L6
L5:
	;
	v23 = v9 + int32(4)
	goto L6
L6:
	;
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = F_str_tolower(m, v23, v51, int32(100))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = v26
	goto L13
L12:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L13
L13:
	;
	if v28 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v26
	goto L16
L15:
	;
	v40 = v37
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v55 = int32(1)
	v56 = v14 + v55
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v61 = v59 & v55
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = v56
	goto L21
L20:
	;
	v62 = v14 + int32(4)
	goto L21
L21:
	;
	if v59 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v92 = F_str_tolower(m, v62, v90, int32(100))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v65 = int32(4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v67&int32(254) == int32(2) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v80 = int32(1)
	if v61 != 0 {
		v90 = int32(base.Ui32(v59)>>(uint(v80)%32)) - v80
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v76 = v65
	goto L28
L27:
	;
	v76 = base.B2i32(v67 == int32(18)) << (uint(v65) % 32)
	goto L28
L28:
	;
	if v67 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = v65
	goto L31
L30:
	;
	v79 = v76
	goto L31
L31:
	;
	v90 = v79
	goto L22
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v97 == int32(0) {
		v116 = v96
		v117 = v97
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v53)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	if v96 != v97 {
		v116 = v96
		v117 = v97
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v101 = v53
	v102 = v92
	goto L38
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v105
		v117 = v106
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v116 = v105
	v117 = v106
	goto L35
L40:
	;
	v109 = int32(1)
	if v105 == v106 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_pfree(m, v92)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v123 != v9 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_pfree(m, v9)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v127 != v14 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	F_pfree(m, v14)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	return base.B2i32(v117-v116 == int32(0))
L51:
	;
	goto L50
}
func F_citext_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_citext_hash(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(1)
	v12 = v7 + v11
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v17 = v15 & v11
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = v12
	goto L5
L4:
	;
	v18 = v7 + int32(4)
	goto L5
L5:
	;
	if v15 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v48 = F_str_tolower(m, v18, v46, int32(100))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v21 = int32(4)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v23&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v36 = int32(1)
	if v17 != 0 {
		v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v32 = v21
	goto L12
L11:
	;
	v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
	goto L12
L12:
	;
	if v23 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = v21
	goto L15
L14:
	;
	v35 = v32
	goto L15
L15:
	;
	v46 = v35
	goto L6
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if v48&int32(3) == int32(0) {
		v73 = v48
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v112 = v106 - int32(1636608432)
	if v48&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v106 = v98 - v48
	goto L18
L20:
	;
	v77 = v73
	goto L29
L21:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v57 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v106 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v62 = v48
	goto L25
L25:
	;
	v66 = v62 + int32(1)
	if v66&int32(3) == int32(0) {
		v73 = v66
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v98 = v66
	goto L19
L27:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != 0 {
		v62 = v66
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v86 = int32(-2139062144)
	if (int32(16843008)-v83|v83)&v86 == v86 {
		v77 = v77 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v92 = v77
	goto L32
L31:
	;
	goto L30
L32:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 != 0 {
		v92 = v92 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v98 = v92
	goto L19
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v48)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L75
	}
L36:
	;
	v344 = int32(14)
	v346 = v340 ^ v341 - base.I32_rotl(v340, v344)
	v350 = v346 ^ v339 - base.I32_rotl(v346, int32(11))
	v354 = v350 ^ v340 - base.I32_rotl(v350, int32(25))
	v358 = v354 ^ v346 - base.I32_rotl(v354, int32(16))
	v362 = v358 ^ v350 - base.I32_rotl(v358, int32(4))
	v366 = v362 ^ v354 - base.I32_rotl(v362, v344)
	goto L35
L37:
	;
	switch v270 - int32(1) {
	case 0:
		v332 = v271
		v333 = v272
		v334 = v273
		goto L64
	case 1:
		v325 = v271
		v326 = v272
		v327 = v273
		goto L65
	case 2:
		v318 = v271
		v319 = v272
		v320 = v273
		goto L66
	case 3:
		v312 = v272
		v313 = v273
		goto L67
	case 4:
		v308 = v272
		v309 = v273
		goto L68
	case 5:
		v302 = v272
		v303 = v273
		goto L69
	case 6:
		v296 = v272
		v297 = v273
		goto L70
	case 7:
		v291 = v273
		goto L71
	case 8:
		v286 = v273
		goto L72
	case 9:
		v281 = v273
		goto L73
	case 10:
		goto L74
	default:
		v339 = v271
		v340 = v272
		v341 = v273
		goto L36
	}
L38:
	;
	v221 = v48
	v222 = v106
	v223 = v112
	v224 = v112
	v225 = v112
	goto L61
L39:
	;
	if base.Ui32(int32(11)) < base.Ui32(v106) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v106) < base.Ui32(int32(12)) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v269 = v48
	v270 = v106
	v271 = v112
	v272 = v112
	v273 = v112
	goto L37
L43:
	;
	switch v168 - int32(1) {
	case 0:
		v218 = v169
		goto L50
	case 1:
		v213 = v169
		goto L51
	case 2:
		goto L52
	case 3:
		v206 = v170
		goto L53
	case 4:
		v203 = v170
		goto L54
	case 5:
		v198 = v170
		goto L55
	case 6:
		goto L56
	case 7:
		v189 = v171
		goto L57
	case 8:
		v184 = v171
		goto L58
	case 9:
		v179 = v171
		goto L59
	case 10:
		goto L60
	default:
		v339 = v169
		v340 = v170
		v341 = v171
		goto L36
	}
L44:
	;
	v167 = v48
	v168 = v106
	v169 = v112
	v170 = v112
	v171 = v112
	goto L43
L45:
	;
	goto L46
L46:
	;
	v119 = v48
	v120 = v106
	v121 = v112
	v122 = v112
	v123 = v112
	goto L47
L47:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v126 = v125 + v122
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	v130 = v129 + v123
	v132 = int32(4)
	v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
	v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
	v139 = v130 + v126
	v140 = v134 + v139
	v141 = v138 + v140
	v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
	v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
	v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
	v154 = v145 + v141
	v155 = v149 + v154
	v156 = v153 + v155
	v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
	v161 = int32(12)
	v162 = v119 + v161
	v164 = v120 - v161
	if base.Ui32(int32(11)) < base.Ui32(v164) {
		v119 = v162
		v120 = v164
		v121 = v155
		v122 = v156
		v123 = v160
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v167 = v162
	v168 = v164
	v169 = v155
	v170 = v156
	v171 = v160
	goto L43
L49:
	;
	goto L48
L50:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v339 = v218 + v219
	v340 = v170
	v341 = v171
	goto L36
L51:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v218 = v214<<(uint(int32(8))%32) + v213
	goto L50
L52:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
	v213 = v209<<(uint(int32(16))%32) + v169
	goto L51
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v339 = v207 + v169
	v340 = v206
	v341 = v171
	goto L36
L54:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
	v206 = v203 + v204
	goto L53
L55:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
	v203 = v199<<(uint(int32(8))%32) + v198
	goto L54
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
	v198 = v194<<(uint(int32(16))%32) + v170
	goto L55
L57:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v339 = v190 + v169
	v340 = v192 + v170
	v341 = v189
	goto L36
L58:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
	v189 = v185<<(uint(int32(8))%32) + v184
	goto L57
L59:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
	v184 = v180<<(uint(int32(16))%32) + v179
	goto L58
L60:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+10)))
	v179 = v175<<(uint(int32(24))%32) + v171
	goto L59
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v228 = v227 + v224
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v232 = v231 + v225
	v234 = int32(4)
	v236 = v229 + v223 - v232 ^ base.I32_rotl(v232, v234)
	v240 = v228 - v236 ^ base.I32_rotl(v236, int32(6))
	v241 = v232 + v228
	v242 = v236 + v241
	v243 = v240 + v242
	v247 = v241 - v240 ^ base.I32_rotl(v240, int32(8))
	v251 = v242 - v247 ^ base.I32_rotl(v247, int32(16))
	v255 = v243 - v251 ^ base.I32_rotl(v251, int32(19))
	v256 = v247 + v243
	v257 = v251 + v256
	v258 = v255 + v257
	v262 = v256 - v255 ^ base.I32_rotl(v255, v234)
	v263 = int32(12)
	v264 = v221 + v263
	v266 = v222 - v263
	if base.Ui32(int32(11)) < base.Ui32(v266) {
		v221 = v264
		v222 = v266
		v223 = v257
		v224 = v258
		v225 = v262
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v269 = v264
	v270 = v266
	v271 = v257
	v272 = v258
	v273 = v262
	goto L37
L63:
	;
	goto L62
L64:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v339 = v332 + v335
	v340 = v333
	v341 = v334
	goto L36
L65:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	v332 = v328<<(uint(int32(8))%32) + v325
	v333 = v326
	v334 = v327
	goto L64
L66:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+2)))
	v325 = v321<<(uint(int32(16))%32) + v318
	v326 = v319
	v327 = v320
	goto L65
L67:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+3)))
	v318 = v314<<(uint(int32(24))%32) + v271
	v319 = v312
	v320 = v313
	goto L66
L68:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+4)))
	v312 = v308 + v310
	v313 = v309
	goto L67
L69:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+5)))
	v308 = v304<<(uint(int32(8))%32) + v302
	v309 = v303
	goto L68
L70:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+6)))
	v302 = v298<<(uint(int32(16))%32) + v296
	v303 = v297
	goto L69
L71:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+7)))
	v296 = v292<<(uint(int32(24))%32) + v272
	v297 = v291
	goto L70
L72:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+8)))
	v291 = v287<<(uint(int32(8))%32) + v286
	goto L71
L73:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+9)))
	v286 = v282<<(uint(int32(16))%32) + v281
	goto L72
L74:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+10)))
	v281 = v277<<(uint(int32(24))%32) + v273
	goto L73
L75:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v373 != v7 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v7)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	return v366 ^ v358 - base.I32_rotl(v366, int32(24))
L79:
	;
	goto L78
}
