package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileCreateTemp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v1 = l0
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_OpenTemporaryFile(m, v1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc(m, int32(_a_F_BufFileCreateTemp_0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v14)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileCreateTemp[0]))
				v20 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v19
				*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v20
				*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v20
				v30 = F_palloc(m, int32(4))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v30
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v9
					*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(0)
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v1)
					return v12
				}
			}
		}
	}
}
func F_BufFileOpenFileSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int64
	_ = v462
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	v15 = m.G0
	v17 = v15 - int32(1056)
	m.G0 = v17
	v19 = int32(16)
	v22 = F_palloc_mul(m, int32(4), v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = v19
	v35 = v22
	v36 = int32(0)
	goto L3
L3:
	;
	v41 = v36 + int32(1)
	if base.Ui32(v30) < base.Ui32(v41) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v36 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L5:
	;
	v45 = v30 << (uint(int32(1)) % 32)
	v46 = F_repalloc_mul(m, v35, int32(4), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v48 = v30
	v49 = v35
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v36
	v53 = v17 + int32(32)
	v58 = F_pg_snprintf(m, v53, int32(1024), int32(_a_F_BufFileOpenFileSet_0), v17+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v48 = v45
	v49 = v46
	goto L7
L9:
	;
	v60 = m.G0
	v62 = v60 - int32(2080)
	m.G0 = v62
	v65 = v62 + int32(1056)
	v66 = F_strlen(m, v53)
	mBase = m.M
	v72 = v66 - int32(1636608432)
	if v53&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v332 = base.I32_rem_u_s(v326^v318-base.I32_rotl(v326, int32(24)), v331)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0+v332<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v65, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L50
	}
L11:
	;
	v304 = int32(14)
	v306 = v300 ^ v301 - base.I32_rotl(v300, v304)
	v310 = v306 ^ v299 - base.I32_rotl(v306, int32(11))
	v314 = v310 ^ v300 - base.I32_rotl(v310, int32(25))
	v318 = v314 ^ v306 - base.I32_rotl(v314, int32(16))
	v322 = v318 ^ v310 - base.I32_rotl(v318, int32(4))
	v326 = v322 ^ v314 - base.I32_rotl(v322, v304)
	goto L10
L12:
	;
	switch v230 - int32(1) {
	case 0:
		v292 = v231
		v293 = v232
		v294 = v233
		goto L39
	case 1:
		v285 = v231
		v286 = v232
		v287 = v233
		goto L40
	case 2:
		v278 = v231
		v279 = v232
		v280 = v233
		goto L41
	case 3:
		v272 = v232
		v273 = v233
		goto L42
	case 4:
		v268 = v232
		v269 = v233
		goto L43
	case 5:
		v262 = v232
		v263 = v233
		goto L44
	case 6:
		v256 = v232
		v257 = v233
		goto L45
	case 7:
		v251 = v233
		goto L46
	case 8:
		v246 = v233
		goto L47
	case 9:
		v241 = v233
		goto L48
	case 10:
		goto L49
	default:
		v299 = v231
		v300 = v232
		v301 = v233
		goto L11
	}
L13:
	;
	v181 = v53
	v182 = v66
	v183 = v72
	v184 = v72
	v185 = v72
	goto L36
L14:
	;
	if base.Ui32(int32(11)) < base.Ui32(v66) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v66) < base.Ui32(int32(12)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v229 = v53
	v230 = v66
	v231 = v72
	v232 = v72
	v233 = v72
	goto L12
L18:
	;
	switch v128 - int32(1) {
	case 0:
		v178 = v129
		goto L25
	case 1:
		v173 = v129
		goto L26
	case 2:
		goto L27
	case 3:
		v166 = v130
		goto L28
	case 4:
		v163 = v130
		goto L29
	case 5:
		v158 = v130
		goto L30
	case 6:
		goto L31
	case 7:
		v149 = v131
		goto L32
	case 8:
		v144 = v131
		goto L33
	case 9:
		v139 = v131
		goto L34
	case 10:
		goto L35
	default:
		v299 = v129
		v300 = v130
		v301 = v131
		goto L11
	}
L19:
	;
	v127 = v53
	v128 = v66
	v129 = v72
	v130 = v72
	v131 = v72
	goto L18
L20:
	;
	goto L21
L21:
	;
	v79 = v53
	v80 = v66
	v81 = v72
	v82 = v72
	v83 = v72
	goto L22
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v86 = v85 + v82
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v90 = v89 + v83
	v92 = int32(4)
	v94 = v87 + v81 - v90 ^ base.I32_rotl(v90, v92)
	v98 = v86 - v94 ^ base.I32_rotl(v94, int32(6))
	v99 = v90 + v86
	v100 = v94 + v99
	v101 = v98 + v100
	v105 = v99 - v98 ^ base.I32_rotl(v98, int32(8))
	v109 = v100 - v105 ^ base.I32_rotl(v105, int32(16))
	v113 = v101 - v109 ^ base.I32_rotl(v109, int32(19))
	v114 = v105 + v101
	v115 = v109 + v114
	v116 = v113 + v115
	v120 = v114 - v113 ^ base.I32_rotl(v113, v92)
	v121 = int32(12)
	v122 = v79 + v121
	v124 = v80 - v121
	if base.Ui32(int32(11)) < base.Ui32(v124) {
		v79 = v122
		v80 = v124
		v81 = v115
		v82 = v116
		v83 = v120
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v127 = v122
	v128 = v124
	v129 = v115
	v130 = v116
	v131 = v120
	goto L18
L24:
	;
	goto L23
L25:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v299 = v178 + v179
	v300 = v130
	v301 = v131
	goto L11
L26:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v178 = v174<<(uint(int32(8))%32) + v173
	goto L25
L27:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	v173 = v169<<(uint(int32(16))%32) + v129
	goto L26
L28:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v299 = v167 + v129
	v300 = v166
	v301 = v131
	goto L11
L29:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	v166 = v163 + v164
	goto L28
L30:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
	v163 = v159<<(uint(int32(8))%32) + v158
	goto L29
L31:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
	v158 = v154<<(uint(int32(16))%32) + v130
	goto L30
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v299 = v150 + v129
	v300 = v152 + v130
	v301 = v149
	goto L11
L33:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
	v149 = v145<<(uint(int32(8))%32) + v144
	goto L32
L34:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
	v144 = v140<<(uint(int32(16))%32) + v139
	goto L33
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
	v139 = v135<<(uint(int32(24))%32) + v131
	goto L34
L36:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v188 = v187 + v184
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v192 = v191 + v185
	v194 = int32(4)
	v196 = v189 + v183 - v192 ^ base.I32_rotl(v192, v194)
	v200 = v188 - v196 ^ base.I32_rotl(v196, int32(6))
	v201 = v192 + v188
	v202 = v196 + v201
	v203 = v200 + v202
	v207 = v201 - v200 ^ base.I32_rotl(v200, int32(8))
	v211 = v202 - v207 ^ base.I32_rotl(v207, int32(16))
	v215 = v203 - v211 ^ base.I32_rotl(v211, int32(19))
	v216 = v207 + v203
	v217 = v211 + v216
	v218 = v215 + v217
	v222 = v216 - v215 ^ base.I32_rotl(v215, v194)
	v223 = int32(12)
	v224 = v181 + v223
	v226 = v182 - v223
	if base.Ui32(int32(11)) < base.Ui32(v226) {
		v181 = v224
		v182 = v226
		v183 = v217
		v184 = v218
		v185 = v222
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v229 = v224
	v230 = v226
	v231 = v217
	v232 = v218
	v233 = v222
	goto L12
L38:
	;
	goto L37
L39:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v299 = v292 + v295
	v300 = v293
	v301 = v294
	goto L11
L40:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v292 = v288<<(uint(int32(8))%32) + v285
	v293 = v286
	v294 = v287
	goto L39
L41:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+2)))
	v285 = v281<<(uint(int32(16))%32) + v278
	v286 = v279
	v287 = v280
	goto L40
L42:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+3)))
	v278 = v274<<(uint(int32(24))%32) + v231
	v279 = v272
	v280 = v273
	goto L41
L43:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+4)))
	v272 = v268 + v270
	v273 = v269
	goto L42
L44:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+5)))
	v268 = v264<<(uint(int32(8))%32) + v262
	v269 = v263
	goto L43
L45:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+6)))
	v262 = v258<<(uint(int32(16))%32) + v256
	v263 = v257
	goto L44
L46:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+7)))
	v256 = v252<<(uint(int32(24))%32) + v232
	v257 = v251
	goto L45
L47:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+8)))
	v251 = v247<<(uint(int32(8))%32) + v246
	goto L46
L48:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+9)))
	v246 = v242<<(uint(int32(16))%32) + v241
	goto L47
L49:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+10)))
	v241 = v237<<(uint(int32(24))%32) + v233
	goto L48
L50:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = int32(_a_F_BufFileOpenFileSet_1)
	*(*int64)(unsafe.Add(mBase, uint32(v62)+24)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v65
	v345 = v62 + int32(32)
	v350 = F_pg_snprintf(m, v345, int32(1024), int32(_a_F_BufFileOpenFileSet_2), v62+int32(16))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v345
	v356 = F_pg_snprintf(m, v65, int32(1024), int32(_a_F_BufFileOpenFileSet_3), v62)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v358 = m.G0
	v360 = v358 - int32(16)
	m.G0 = v360
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	F_ResourceOwnerEnlarge(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[1]))
	v368 = F_PathNameOpenFilePerm(m, v65, l2, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	m.G0 = v360 + int32(16)
	m.G0 = v62 + int32(2080)
	*(*int32)(unsafe.Add(mBase, uint32(v49+v36<<(uint(int32(2))%32)))) = v368
	if int32(0) < v368 {
		goto L65
	} else {
		goto L66
	}
L55:
	;
	if v368 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[2]))
	if v373 == int32(44) {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	F_ResourceOwnerRemember(m, v392, base.I64_extend_i32_u(v368), int32(_a_F_BufFileOpenFileSet_4))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L64
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v65
	F_errmsg(m, int32(_a_F_BufFileOpenFileSet_5), v360)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_BufFileOpenFileSet_6), int32(1909), int32(_a_F_BufFileOpenFileSet_7))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[3]))
	v401 = v398 + v368*int32(48)
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = v403
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+4)))
	v407 = v405 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v401)+4)) = uint16(v407)
	v410 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[4])) = uint8(v410)
	goto L54
L65:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[5]))
	if v426 == int32(0) {
		v30 = v48
		v35 = v49
		v36 = v41
		goto L3
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L4
L68:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v30 = v48
	v35 = v49
	v36 = v41
	goto L3
L70:
	;
	m.G0 = v17 + int32(1056)
	return v480
L71:
	;
	F_pfree(m, v49)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v455 = F_palloc(m, int32(_a_F_BufFileOpenFileSet_8))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	if l3 != 0 {
		v480 = int32(0)
		goto L70
	} else {
		goto L75
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(32)
	F_errmsg(m, int32(_a_F_BufFileOpenFileSet_9), v17)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_BufFileOpenFileSet_10), int32(340), int32(_a_F_BufFileOpenFileSet_11))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v457 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v455)+8)) = uint16(v457)
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = v36
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	v462 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v455)+32)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v455)+24)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455)+20)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v455)+40)) = v462
	*(*int64)(unsafe.Add(mBase, uint32(v455)+48)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v455)+12)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+10)) = uint8(base.B2i32(l2 == v457))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v49
	v476 = F_pstrdup(m, l1)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+16)) = v476
	v480 = v455
	goto L70
}
func F_WaitBufHdrUnlocked(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int64
	_ = v21
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	v3 = int64(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_WaitBufHdrUnlocked_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(_a_F_WaitBufHdrUnlocked_1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_WaitBufHdrUnlocked_2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v3
	v21 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v3, v3)
	if v21&int64(4194304) != v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	v45 = v21
	goto L3
L3:
	;
	v49 = int32(_a_F_WaitBufHdrUnlocked_3)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_WaitBufHdrUnlocked[0]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v52 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v45 = v38
	goto L3
L6:
	;
	return int64(0)
L7:
	;
	v35 = int64(0)
	v38 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v35, v35)
	if v38&int64(4194304) != v35 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	m.G0 = v6 + int32(32)
	return v45
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitBufHdrUnlocked[0])) = v67
	goto L10
L12:
	;
	if int32(999) < v50 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v50 < int32(11) {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	v57 = int32(900)
	if v57 <= v50 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = v57
	goto L18
L17:
	;
	v60 = v50
	goto L18
L18:
	;
	v67 = v60 + int32(100)
	goto L11
L19:
	;
	v67 = v50 - int32(1)
	goto L11
}
