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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
				v28 = F_palloc(m, int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v9
					*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(0)
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v34)
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	v15 = m.G0
	v17 = v15 - int32(1056)
	m.G0 = v17
	v21 = F_palloc(m, int32(64))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = int32(16)
	v34 = int32(0)
	v35 = v21
	goto L3
L3:
	;
	v40 = v34 + int32(1)
	if base.Ui32(v29) < base.Ui32(v40) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v34 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L5:
	;
	v44 = F_repalloc(m, v35, v29<<(uint(int32(3))%32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v48 = v29
	v49 = v35
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v34
	v58 = F_pg_snprintf(m, v17+int32(32), int32(1024), int32(_a_F_BufFileOpenFileSet_0), v17+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v48 = v29 << (uint(int32(1)) % 32)
	v49 = v44
	goto L7
L9:
	;
	v60 = m.G0
	v62 = v60 - int32(2080)
	m.G0 = v62
	v67 = v17 + int32(32)
	v68 = F_strlen(m, v67)
	mBase = m.M
	v74 = v68 - int32(1636608432)
	if v67&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v334 = base.I32_rem_u_s(v328^v320-base.I32_rotl(v328, int32(24)), v333)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0+v334<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v62+int32(1056), v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L50
	}
L11:
	;
	v306 = int32(14)
	v308 = v302 ^ v303 - base.I32_rotl(v302, v306)
	v312 = v308 ^ v301 - base.I32_rotl(v308, int32(11))
	v316 = v312 ^ v302 - base.I32_rotl(v312, int32(25))
	v320 = v316 ^ v308 - base.I32_rotl(v316, int32(16))
	v324 = v320 ^ v312 - base.I32_rotl(v320, int32(4))
	v328 = v324 ^ v316 - base.I32_rotl(v324, v306)
	goto L10
L12:
	;
	switch v232 - int32(1) {
	case 0:
		v294 = v233
		v295 = v234
		v296 = v235
		goto L39
	case 1:
		v287 = v233
		v288 = v234
		v289 = v235
		goto L40
	case 2:
		v280 = v233
		v281 = v234
		v282 = v235
		goto L41
	case 3:
		v274 = v234
		v275 = v235
		goto L42
	case 4:
		v270 = v234
		v271 = v235
		goto L43
	case 5:
		v264 = v234
		v265 = v235
		goto L44
	case 6:
		v258 = v234
		v259 = v235
		goto L45
	case 7:
		v253 = v235
		goto L46
	case 8:
		v248 = v235
		goto L47
	case 9:
		v243 = v235
		goto L48
	case 10:
		goto L49
	default:
		v301 = v233
		v302 = v234
		v303 = v235
		goto L11
	}
L13:
	;
	v183 = v67
	v184 = v68
	v185 = v74
	v186 = v74
	v187 = v74
	goto L36
L14:
	;
	if base.Ui32(int32(11)) < base.Ui32(v68) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v68) < base.Ui32(int32(12)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v231 = v67
	v232 = v68
	v233 = v74
	v234 = v74
	v235 = v74
	goto L12
L18:
	;
	switch v130 - int32(1) {
	case 0:
		v180 = v131
		goto L25
	case 1:
		v175 = v131
		goto L26
	case 2:
		goto L27
	case 3:
		v168 = v132
		goto L28
	case 4:
		v165 = v132
		goto L29
	case 5:
		v160 = v132
		goto L30
	case 6:
		goto L31
	case 7:
		v151 = v133
		goto L32
	case 8:
		v146 = v133
		goto L33
	case 9:
		v141 = v133
		goto L34
	case 10:
		goto L35
	default:
		v301 = v131
		v302 = v132
		v303 = v133
		goto L11
	}
L19:
	;
	v129 = v67
	v130 = v68
	v131 = v74
	v132 = v74
	v133 = v74
	goto L18
L20:
	;
	goto L21
L21:
	;
	v81 = v67
	v82 = v68
	v83 = v74
	v84 = v74
	v85 = v74
	goto L22
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v87 + v84
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v92 = v91 + v85
	v94 = int32(4)
	v96 = v89 + v83 - v92 ^ base.I32_rotl(v92, v94)
	v100 = v88 - v96 ^ base.I32_rotl(v96, int32(6))
	v101 = v92 + v88
	v102 = v96 + v101
	v103 = v100 + v102
	v107 = v101 - v100 ^ base.I32_rotl(v100, int32(8))
	v111 = v102 - v107 ^ base.I32_rotl(v107, int32(16))
	v115 = v103 - v111 ^ base.I32_rotl(v111, int32(19))
	v116 = v107 + v103
	v117 = v111 + v116
	v118 = v115 + v117
	v122 = v116 - v115 ^ base.I32_rotl(v115, v94)
	v123 = int32(12)
	v124 = v81 + v123
	v126 = v82 - v123
	if base.Ui32(int32(11)) < base.Ui32(v126) {
		v81 = v124
		v82 = v126
		v83 = v117
		v84 = v118
		v85 = v122
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v129 = v124
	v130 = v126
	v131 = v117
	v132 = v118
	v133 = v122
	goto L18
L24:
	;
	goto L23
L25:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v301 = v180 + v181
	v302 = v132
	v303 = v133
	goto L11
L26:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v180 = v176<<(uint(int32(8))%32) + v175
	goto L25
L27:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
	v175 = v171<<(uint(int32(16))%32) + v131
	goto L26
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v301 = v169 + v131
	v302 = v168
	v303 = v133
	goto L11
L29:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
	v168 = v165 + v166
	goto L28
L30:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
	v165 = v161<<(uint(int32(8))%32) + v160
	goto L29
L31:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
	v160 = v156<<(uint(int32(16))%32) + v132
	goto L30
L32:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v301 = v152 + v131
	v302 = v154 + v132
	v303 = v151
	goto L11
L33:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
	v151 = v147<<(uint(int32(8))%32) + v146
	goto L32
L34:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
	v146 = v142<<(uint(int32(16))%32) + v141
	goto L33
L35:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+10)))
	v141 = v137<<(uint(int32(24))%32) + v133
	goto L34
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v190 = v189 + v186
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v194 = v193 + v187
	v196 = int32(4)
	v198 = v191 + v185 - v194 ^ base.I32_rotl(v194, v196)
	v202 = v190 - v198 ^ base.I32_rotl(v198, int32(6))
	v203 = v194 + v190
	v204 = v198 + v203
	v205 = v202 + v204
	v209 = v203 - v202 ^ base.I32_rotl(v202, int32(8))
	v213 = v204 - v209 ^ base.I32_rotl(v209, int32(16))
	v217 = v205 - v213 ^ base.I32_rotl(v213, int32(19))
	v218 = v209 + v205
	v219 = v213 + v218
	v220 = v217 + v219
	v224 = v218 - v217 ^ base.I32_rotl(v217, v196)
	v225 = int32(12)
	v226 = v183 + v225
	v228 = v184 - v225
	if base.Ui32(int32(11)) < base.Ui32(v228) {
		v183 = v226
		v184 = v228
		v185 = v219
		v186 = v220
		v187 = v224
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v231 = v226
	v232 = v228
	v233 = v219
	v234 = v220
	v235 = v224
	goto L12
L38:
	;
	goto L37
L39:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v301 = v294 + v297
	v302 = v295
	v303 = v296
	goto L11
L40:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	v294 = v290<<(uint(int32(8))%32) + v287
	v295 = v288
	v296 = v289
	goto L39
L41:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+2)))
	v287 = v283<<(uint(int32(16))%32) + v280
	v288 = v281
	v289 = v282
	goto L40
L42:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+3)))
	v280 = v276<<(uint(int32(24))%32) + v233
	v281 = v274
	v282 = v275
	goto L41
L43:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
	v274 = v270 + v272
	v275 = v271
	goto L42
L44:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+5)))
	v270 = v266<<(uint(int32(8))%32) + v264
	v271 = v265
	goto L43
L45:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+6)))
	v264 = v260<<(uint(int32(16))%32) + v258
	v265 = v259
	goto L44
L46:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+7)))
	v258 = v254<<(uint(int32(24))%32) + v234
	v259 = v253
	goto L45
L47:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+8)))
	v253 = v249<<(uint(int32(8))%32) + v248
	goto L46
L48:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+9)))
	v248 = v244<<(uint(int32(16))%32) + v243
	goto L47
L49:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+10)))
	v243 = v239<<(uint(int32(24))%32) + v235
	goto L48
L50:
	;
	v341 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = int32(_a_F_BufFileOpenFileSet_1)
	*(*int64)(unsafe.Add(mBase, uint32(v62)+24)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v62 + int32(1056)
	v354 = F_pg_snprintf(m, v62+int32(32), int32(1024), int32(_a_F_BufFileOpenFileSet_2), v62+int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v62 + int32(32)
	v364 = F_pg_snprintf(m, v62+int32(1056), int32(1024), int32(_a_F_BufFileOpenFileSet_3), v62)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v367 = v62 + int32(1056)
	v368 = m.G0
	v370 = v368 - int32(16)
	m.G0 = v370
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	F_ResourceOwnerEnlarge(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[1]))
	v378 = F_PathNameOpenFilePerm(m, v367, l2, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	m.G0 = v370 + int32(16)
	m.G0 = v62 + int32(2080)
	*(*int32)(unsafe.Add(mBase, uint32(v49+v34<<(uint(int32(2))%32)))) = v378
	if int32(0) < v378 {
		goto L65
	} else {
		goto L66
	}
L55:
	;
	if v378 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[2]))
	if v383 == int32(44) {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	F_ResourceOwnerRemember(m, v402, v378, int32(_a_F_BufFileOpenFileSet_4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L64
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = v367
	F_errmsg(m, int32(_a_F_BufFileOpenFileSet_5), v370)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_BufFileOpenFileSet_6), int32(1925), int32(_a_F_BufFileOpenFileSet_7))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
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
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[3]))
	v410 = v407 + v378*int32(48)
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+8)) = v412
	v415 = v410 + int32(4)
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415))))
	v418 = v416 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v415))) = uint16(v418)
	v421 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[4])) = uint8(v421)
	goto L54
L65:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[5]))
	if v437 == int32(0) {
		v29 = v48
		v34 = v40
		v35 = v49
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
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v29 = v48
	v34 = v40
	v35 = v49
	goto L3
L70:
	;
	m.G0 = v17 + int32(1056)
	return v488
L71:
	;
	F_pfree(m, v49)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v466 = F_palloc(m, int32(_a_F_BufFileOpenFileSet_8))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	if l3 != 0 {
		v488 = int32(0)
		goto L70
	} else {
		goto L75
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
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
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_BufFileOpenFileSet_10), int32(339), int32(_a_F_BufFileOpenFileSet_11))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
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
	v468 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v466)+8)) = uint16(v468)
	*(*int32)(unsafe.Add(mBase, uint32(v466))) = v34
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileOpenFileSet[0]))
	v473 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v466)+32)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v466)+24)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v466)+20)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v466)+40)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v466)+12)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+10)) = uint8(base.B2i32(l2 == v468))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+4)) = v49
	v485 = F_pstrdup(m, l1)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466)+16)) = v485
	v488 = v466
	goto L70
}
