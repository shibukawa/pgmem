package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomInitMetapage(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = int32(0)
	v6 = F_ReadBufferExtended(m, l0, l1, int32(-1), v4, v4)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_LockBufferInternal(m, v6, int32(3))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = F_GenericXLogStart(m, l0)
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = F_GenericXLogRegisterBuffer(m, v11, v6, int32(1))
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_BloomFillMetapage(m, l0, v14)
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						F_GenericXLogFinish(m, v11)
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v6)
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_bloom_lacks_element(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
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
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
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
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
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
	var v333 int64
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v432 int32
	_ = v432
	var __phi432 int32
	_ = __phi432
	var v433 int32
	_ = v433
	var __phi433 int32
	_ = __phi433
	var v434 int32
	_ = v434
	var __phi434 int32
	_ = __phi434
	var v437 int32
	_ = v437
	var __phi437 int32
	_ = __phi437
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v477 int32
	_ = v477
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = l2 - int32(1636608432)
	if v18 == int64(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v336 = v334 - int32(1)
	v338 = v336 & base.I32_wrap_i64(v333)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(2) <= v340 {
		goto L46
	} else {
		goto L47
	}
L2:
	;
	if l1&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	v61 = v24
	v63 = v24
	v65 = v24
	goto L2
L4:
	;
	goto L5
L5:
	;
	v28 = v24 + base.I32_wrap_i64(v18)
	v29 = v28 + v24
	v33 = int32(4)
	v35 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ base.I32_rotl(v24, v33)
	v39 = v28 - v35 ^ base.I32_rotl(v35, int32(6))
	v43 = v29 - v39 ^ base.I32_rotl(v39, int32(8))
	v44 = v29 + v35
	v45 = v39 + v44
	v46 = v43 + v45
	v50 = v44 - v43 ^ base.I32_rotl(v43, int32(16))
	v54 = v45 - v50 ^ base.I32_rotl(v50, int32(19))
	v59 = v46 + v50
	v61 = v59
	v63 = v46 - v54 ^ base.I32_rotl(v54, v33)
	v65 = v54 + v59
	goto L2
L6:
	;
	v302 = int32(14)
	v304 = v298 ^ v299 - base.I32_rotl(v298, v302)
	v308 = v304 ^ v297 - base.I32_rotl(v304, int32(11))
	v312 = v308 ^ v298 - base.I32_rotl(v308, int32(25))
	v316 = v312 ^ v304 - base.I32_rotl(v312, int32(16))
	v320 = v316 ^ v308 - base.I32_rotl(v316, int32(4))
	v324 = v320 ^ v312 - base.I32_rotl(v320, v302)
	v333 = base.I64_extend_i32_u(v324)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v324^v316-base.I32_rotl(v324, int32(24)))
	goto L1
L7:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v297 = v289 + v292
	v298 = v290
	v299 = v291
	goto L6
L8:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v289 = v285<<(uint(int32(8))%32) + v282
	v290 = v283
	v291 = v284
	goto L7
L9:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
	v282 = v278<<(uint(int32(16))%32) + v275
	v283 = v276
	v284 = v277
	goto L8
L10:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
	v275 = v271<<(uint(int32(24))%32) + v122
	v276 = v269
	v277 = v270
	goto L9
L11:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
	v269 = v265 + v267
	v270 = v266
	goto L10
L12:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
	v265 = v261<<(uint(int32(8))%32) + v259
	v266 = v260
	goto L11
L13:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
	v259 = v255<<(uint(int32(16))%32) + v253
	v260 = v254
	goto L12
L14:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
	v253 = v249<<(uint(int32(24))%32) + v123
	v254 = v248
	goto L13
L15:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
	v248 = v244<<(uint(int32(8))%32) + v243
	goto L14
L16:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+9)))
	v243 = v239<<(uint(int32(16))%32) + v238
	goto L15
L17:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+10)))
	v238 = v234<<(uint(int32(24))%32) + v124
	goto L16
L18:
	;
	if base.Ui32(int32(11)) < base.Ui32(l2) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(int32(12)) <= base.Ui32(l2) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v70 = l1
	v71 = l2
	v73 = v61
	v74 = v65
	v75 = v63
	goto L24
L22:
	;
	v119 = l1
	v120 = l2
	v122 = v61
	v123 = v65
	v124 = v63
	goto L23
L23:
	;
	switch v120 - int32(1) {
	case 0:
		v289 = v122
		v290 = v123
		v291 = v124
		goto L7
	case 1:
		v282 = v122
		v283 = v123
		v284 = v124
		goto L8
	case 2:
		v275 = v122
		v276 = v123
		v277 = v124
		goto L9
	case 3:
		v269 = v123
		v270 = v124
		goto L10
	case 4:
		v265 = v123
		v266 = v124
		goto L11
	case 5:
		v259 = v123
		v260 = v124
		goto L12
	case 6:
		v253 = v123
		v254 = v124
		goto L13
	case 7:
		v248 = v124
		goto L14
	case 8:
		v243 = v124
		goto L15
	case 9:
		v238 = v124
		goto L16
	case 10:
		goto L17
	default:
		v297 = v122
		v298 = v123
		v299 = v124
		goto L6
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v78 = v77 + v74
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v82 = v81 + v75
	v84 = int32(4)
	v86 = v79 + v73 - v82 ^ base.I32_rotl(v82, v84)
	v90 = v78 - v86 ^ base.I32_rotl(v86, int32(6))
	v91 = v82 + v78
	v92 = v86 + v91
	v93 = v90 + v92
	v97 = v91 - v90 ^ base.I32_rotl(v90, int32(8))
	v101 = v92 - v97 ^ base.I32_rotl(v97, int32(16))
	v105 = v93 - v101 ^ base.I32_rotl(v101, int32(19))
	v106 = v97 + v93
	v107 = v101 + v106
	v108 = v105 + v107
	v112 = v106 - v105 ^ base.I32_rotl(v105, v84)
	v113 = int32(12)
	v114 = v70 + v113
	v116 = v71 - v113
	if base.Ui32(int32(11)) < base.Ui32(v116) {
		v70 = v114
		v71 = v116
		v73 = v107
		v74 = v108
		v75 = v112
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v119 = v114
	v120 = v116
	v122 = v107
	v123 = v108
	v124 = v112
	goto L23
L26:
	;
	goto L25
L27:
	;
	v130 = l1
	v131 = l2
	v133 = v61
	v134 = v65
	v135 = v63
	goto L30
L28:
	;
	v179 = l1
	v180 = l2
	v182 = v61
	v183 = v65
	v184 = v63
	goto L29
L29:
	;
	switch v180 - int32(1) {
	case 0:
		v231 = v182
		goto L33
	case 1:
		v226 = v182
		goto L34
	case 2:
		goto L35
	case 3:
		v219 = v183
		goto L36
	case 4:
		v216 = v183
		goto L37
	case 5:
		v211 = v183
		goto L38
	case 6:
		goto L39
	case 7:
		v202 = v184
		goto L40
	case 8:
		v197 = v184
		goto L41
	case 9:
		v192 = v184
		goto L42
	case 10:
		goto L43
	default:
		v297 = v182
		v298 = v183
		v299 = v184
		goto L6
	}
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v138 = v137 + v134
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v142 = v141 + v135
	v144 = int32(4)
	v146 = v139 + v133 - v142 ^ base.I32_rotl(v142, v144)
	v150 = v138 - v146 ^ base.I32_rotl(v146, int32(6))
	v151 = v142 + v138
	v152 = v146 + v151
	v153 = v150 + v152
	v157 = v151 - v150 ^ base.I32_rotl(v150, int32(8))
	v161 = v152 - v157 ^ base.I32_rotl(v157, int32(16))
	v165 = v153 - v161 ^ base.I32_rotl(v161, int32(19))
	v166 = v157 + v153
	v167 = v161 + v166
	v168 = v165 + v167
	v172 = v166 - v165 ^ base.I32_rotl(v165, v144)
	v173 = int32(12)
	v174 = v130 + v173
	v176 = v131 - v173
	if base.Ui32(int32(11)) < base.Ui32(v176) {
		v130 = v174
		v131 = v176
		v133 = v167
		v134 = v168
		v135 = v172
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v179 = v174
	v180 = v176
	v182 = v167
	v183 = v168
	v184 = v172
	goto L29
L32:
	;
	goto L31
L33:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v297 = v231 + v232
	v298 = v183
	v299 = v184
	goto L6
L34:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v231 = v227<<(uint(int32(8))%32) + v226
	goto L33
L35:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)))
	v226 = v222<<(uint(int32(16))%32) + v182
	goto L34
L36:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v297 = v220 + v182
	v298 = v219
	v299 = v184
	goto L6
L37:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
	v219 = v216 + v217
	goto L36
L38:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)))
	v216 = v212<<(uint(int32(8))%32) + v211
	goto L37
L39:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+6)))
	v211 = v207<<(uint(int32(16))%32) + v183
	goto L38
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v297 = v203 + v182
	v298 = v205 + v183
	v299 = v202
	goto L6
L41:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+8)))
	v202 = v198<<(uint(int32(8))%32) + v197
	goto L40
L42:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+9)))
	v197 = v193<<(uint(int32(16))%32) + v192
	goto L41
L43:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+10)))
	v192 = v188<<(uint(int32(24))%32) + v184
	goto L42
L44:
	;
	m.G0 = v16 + int32(48)
	return v512
L45:
	;
	v477 = int32(0)
	goto L61
L46:
	;
	v344 = v340 - int32(1)
	v345 = int32(3)
	v346 = v344 & v345
	v349 = base.I32_wrap_i64(int64(base.Ui64(v333) >> (uint(int64(32)) % 64)))
	if base.Ui32(v340-int32(2)) < base.Ui32(v345) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L48
L48:
	;
	if v340 != int32(1) {
		v512 = int32(0)
		goto L44
	} else {
		goto L60
	}
L49:
	;
	__phi432 = v418
	__phi433 = v419
	__phi434 = v420
	__phi437 = int32(0)
	v432 = __phi432
	v433 = __phi433
	v434 = __phi434
	v437 = __phi437
	goto L57
L50:
	;
	v418 = int32(1)
	v419 = v349
	v420 = v338
	goto L49
L51:
	;
	goto L52
L52:
	;
	v360 = int32(1)
	v361 = v349
	v362 = v338
	v365 = int32(0)
	goto L53
L53:
	;
	v371 = int32(2)
	v374 = v336 & v361
	v376 = (v362 + v374) & v336
	*(*int32)(unsafe.Add(mBase, uint32(v16+v360<<(uint(v371)%32)))) = v376
	v379 = v360 + int32(1)
	v384 = (v360 + v374) & v336
	v386 = (v384 + v376) & v336
	*(*int32)(unsafe.Add(mBase, uint32(v16+v379<<(uint(v371)%32)))) = v386
	v389 = v360 + v371
	v394 = (v384 + v379) & v336
	v396 = (v394 + v386) & v336
	*(*int32)(unsafe.Add(mBase, uint32(v16+v389<<(uint(v371)%32)))) = v396
	v399 = v360 + int32(3)
	v404 = (v394 + v389) & v336
	v406 = (v396 + v404) & v336
	*(*int32)(unsafe.Add(mBase, uint32(v16+v399<<(uint(v371)%32)))) = v406
	v408 = v404 + v399
	v409 = int32(4)
	v410 = v360 + v409
	v412 = v365 + v409
	if v412 != v344&int32(-4) {
		v360 = v410
		v361 = v408
		v362 = v406
		v365 = v412
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v346 == int32(0) {
		goto L45
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v418 = v410
	v419 = v408
	v420 = v406
	goto L49
L57:
	;
	v446 = v336 & v433
	v448 = (v446 + v434) & v336
	*(*int32)(unsafe.Add(mBase, uint32(v16+v432<<(uint(int32(2))%32)))) = v448
	v451 = int32(1)
	v454 = v437 + v451
	if v454 != v346 {
		__phi432 = v432 + v451
		__phi433 = v432 + v446
		__phi434 = v448
		__phi437 = v454
		v432 = __phi432
		v433 = __phi433
		v434 = __phi434
		v437 = __phi437
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L45
L59:
	;
	goto L58
L60:
	;
	goto L45
L61:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v16+v477<<(uint(int32(2))%32))))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(24)+int32(base.Ui32(v491)>>(uint(int32(3))%32))))))
	v502 = base.B2i32(int32(base.Ui32(v495)>>(uint(v491&int32(7))%32))&int32(1) == int32(0))
	if int32(base.Ui32(v495)>>(uint(v491&int32(7))%32))&int32(1) == int32(0) {
		v512 = v502
		goto L44
	} else {
		goto L63
	}
L62:
	;
	v512 = v502
	goto L44
L63:
	;
	v506 = v477 + int32(1)
	if v506 != v340 {
		v477 = v506
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
}
