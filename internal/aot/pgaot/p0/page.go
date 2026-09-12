package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PageGetTempPage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v5 = F_palloc(m, v2<<(uint(int32(8))%32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v10 = v8 << (uint(int32(8)) % 32)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if base.Ui32(int32(1024)) < base.Ui32(v10) {
			v34 = v10
			v38 = F__emscripten_memset_bulkmem(m, v11, base.I32_extend8_s(int32(0)), v34)
			mBase = m.M
		} else {
			if v11&int32(3) != 0 {
				v34 = v10
				v38 = F__emscripten_memset_bulkmem(m, v11, base.I32_extend8_s(int32(0)), v34)
				mBase = m.M
			} else {
				v21 = v11 + v10
				if base.Ui32(v21) <= base.Ui32(v11) {
				} else {
					v26 = v11 + int32(4)
					if base.Ui32(v26) < base.Ui32(v21) {
						v28 = v21
					} else {
						v28 = v26
					}
					v34 = (v11^int32(-1)+v28)&int32(-4) + int32(4)
					v38 = F__emscripten_memset_bulkmem(m, v11, base.I32_extend8_s(int32(0)), v34)
					mBase = m.M
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+10)) = int32(1572864)
		v44 = v10 | int32(4)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)) = uint16(v44)
		v53 = v10 - (v16&int32(-256)-v15+int32(7))&int32(-8)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v53)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v53)
		v56 = int32(65535)
		v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
		v66 = (v61<<(uint(int32(8))%32) - v59) & v56
		if v66 != 0 {
			v67 = F__emscripten_memcpy_bulkmem(m, v11+v53&v56, l0+v59, v66)
			mBase = m.M
		} else {
		}
		return v11
	}
}
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v17) < base.Ui32(int32(24)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L131
	} else {
		goto L139
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L131
	} else {
		goto L136
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L131
	} else {
		goto L132
	}
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v20) < base.Ui32(v17) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.Ui32(v22) < base.Ui32(v20) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v22) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if (v22+int32(7))&int32(32760) != v22 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v17 != int32(24) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(base.Ui32(v17+int32(262120)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v38 = int32(0)
	goto L11
L11:
	;
	v39 = int32(65535)
	v40 = v38 & v39
	if base.Ui32(v40) <= base.Ui32((l1-int32(1))&v39) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v47 = l0 + int32(24)
	v52 = v47 + l1<<(uint(int32(2))%32) - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = int32(base.Ui32(v53) >> (uint(int32(17)) % 32))
	v57 = v53 & int32(32767)
	if base.Ui32(v57) < base.Ui32(v20) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v22) < base.Ui32(v57+v55) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v57 != (v57+int32(7))&int32(65528) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v70 = v47 + l1<<(uint(int32(2))%32)
	v72 = l0 - v70 + v17
	if int32(0) < v72 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v52 == v70 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v220 = v20
	goto L18
L18:
	;
	v222 = (v55 + int32(7)) & int32(65528)
	if base.Ui32(v220) < base.Ui32(v57) {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v220 = v219
	goto L18
L20:
	;
	goto L19
L21:
	;
	v78 = v52 + v72
	if base.Ui32(v70-v78) <= base.Ui32(int32(0)-v72<<(uint(int32(1))%32)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = F___memcpy(m, v52, v70, v72)
	mBase = m.M
	goto L19
L23:
	;
	goto L24
L24:
	;
	v88 = (v52 ^ v70) & int32(3)
	if base.Ui32(v52) < base.Ui32(v70) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v190 == int32(0) {
		goto L20
	} else {
		goto L61
	}
L26:
	;
	if base.Ui32(v168) <= base.Ui32(int32(3)) {
		v189 = v167
		v190 = v168
		v191 = v169
		goto L25
	} else {
		goto L57
	}
L27:
	;
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v88 != 0 {
		v150 = v72
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v189 = v70
	v190 = v72
	v191 = v52
	goto L25
L31:
	;
	goto L32
L32:
	;
	if v52&int32(3) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v167 = v70
	v168 = v72
	v169 = v52
	goto L26
L34:
	;
	goto L35
L35:
	;
	v95 = v70
	v96 = v72
	v97 = v52
	goto L36
L36:
	;
	if v96 == int32(0) {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v167 = v104
	v168 = v106
	v169 = v108
	goto L26
L38:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v101)
	v103 = int32(1)
	v104 = v95 + v103
	v106 = v96 - v103
	v108 = v97 + v103
	if v108&int32(3) != 0 {
		v95 = v104
		v96 = v106
		v97 = v108
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v150 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L41:
	;
	if v78&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v115 = v72
	goto L45
L43:
	;
	v130 = v72
	goto L44
L44:
	;
	if base.Ui32(v130) <= base.Ui32(int32(3)) {
		v150 = v130
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if v115 == int32(0) {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	v130 = v121
	goto L44
L47:
	;
	v121 = v115 - int32(1)
	v122 = v52 + v121
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v121))))
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v124)
	if v122&int32(3) != 0 {
		v115 = v121
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v137 = v130
	goto L50
L50:
	;
	v141 = v137 - int32(4)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v70+v141)))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v141))) = v144
	if base.Ui32(int32(3)) < base.Ui32(v141) {
		v137 = v141
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v150 = v141
	goto L40
L52:
	;
	goto L51
L53:
	;
	v157 = v150
	goto L54
L54:
	;
	v161 = v157 - int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52+v161))) = uint8(v164)
	if v161 != 0 {
		v157 = v161
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L20
L56:
	;
	goto L55
L57:
	;
	v174 = v167
	v175 = v168
	v176 = v169
	goto L58
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v178
	v180 = int32(4)
	v181 = v174 + v180
	v183 = v176 + v180
	v185 = v175 - v180
	if base.Ui32(int32(3)) < base.Ui32(v185) {
		v174 = v181
		v175 = v185
		v176 = v183
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v189 = v181
	v190 = v185
	v191 = v183
	goto L25
L60:
	;
	goto L59
L61:
	;
	v196 = v189
	v197 = v190
	v198 = v191
	goto L62
L62:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v200)
	v202 = int32(1)
	v207 = v197 - v202
	if v207 != 0 {
		v196 = v196 + v202
		v197 = v207
		v198 = v198 + v202
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L20
L64:
	;
	goto L63
L65:
	;
	v224 = l0 + v220
	v225 = v224 + v222
	v226 = v57 - v220
	if v225 == v224 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v373 = v220
	goto L67
L67:
	;
	v374 = v373 + v222
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v374)
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v378 = v376 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v378)
	if base.Ui32(v378&int32(65535)) < base.Ui32(int32(25)) {
		goto L114
	} else {
		goto L115
	}
L68:
	;
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v373 = v371
	goto L67
L69:
	;
	goto L68
L70:
	;
	v230 = v225 + v226
	if base.Ui32(v224-v230) <= base.Ui32(int32(0)-v226<<(uint(int32(1))%32)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v237 = F___memcpy(m, v225, v224, v226)
	mBase = m.M
	goto L68
L72:
	;
	goto L73
L73:
	;
	v240 = (v225 ^ v224) & int32(3)
	if base.Ui32(v225) < base.Ui32(v224) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	if v342 == int32(0) {
		goto L69
	} else {
		goto L110
	}
L75:
	;
	if base.Ui32(v320) <= base.Ui32(int32(3)) {
		v341 = v319
		v342 = v320
		v343 = v321
		goto L74
	} else {
		goto L106
	}
L76:
	;
	if v240 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	if v240 != 0 {
		v302 = v226
		goto L89
	} else {
		goto L90
	}
L79:
	;
	v341 = v224
	v342 = v226
	v343 = v225
	goto L74
L80:
	;
	goto L81
L81:
	;
	if v225&int32(3) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v319 = v224
	v320 = v226
	v321 = v225
	goto L75
L83:
	;
	goto L84
L84:
	;
	v247 = v224
	v248 = v226
	v249 = v225
	goto L85
L85:
	;
	if v248 == int32(0) {
		goto L69
	} else {
		goto L87
	}
L86:
	;
	v319 = v256
	v320 = v258
	v321 = v260
	goto L75
L87:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	*(*uint8)(unsafe.Add(mBase, uint32(v249))) = uint8(v253)
	v255 = int32(1)
	v256 = v247 + v255
	v258 = v248 - v255
	v260 = v249 + v255
	if v260&int32(3) != 0 {
		v247 = v256
		v248 = v258
		v249 = v260
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v302 == int32(0) {
		goto L69
	} else {
		goto L102
	}
L90:
	;
	if v230&int32(3) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v267 = v226
	goto L94
L92:
	;
	v282 = v226
	goto L93
L93:
	;
	if base.Ui32(v282) <= base.Ui32(int32(3)) {
		v302 = v282
		goto L89
	} else {
		goto L98
	}
L94:
	;
	if v267 == int32(0) {
		goto L69
	} else {
		goto L96
	}
L95:
	;
	v282 = v273
	goto L93
L96:
	;
	v273 = v267 - int32(1)
	v274 = v225 + v273
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v273))))
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v276)
	if v274&int32(3) != 0 {
		v267 = v273
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v289 = v282
	goto L99
L99:
	;
	v293 = v289 - int32(4)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v224+v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v225+v293))) = v296
	if base.Ui32(int32(3)) < base.Ui32(v293) {
		v289 = v293
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v302 = v293
	goto L89
L101:
	;
	goto L100
L102:
	;
	v309 = v302
	goto L103
L103:
	;
	v313 = v309 - int32(1)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v313))))
	*(*uint8)(unsafe.Add(mBase, uint32(v225+v313))) = uint8(v316)
	if v313 != 0 {
		v309 = v313
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L69
L105:
	;
	goto L104
L106:
	;
	v326 = v319
	v327 = v320
	v328 = v321
	goto L107
L107:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = v330
	v332 = int32(4)
	v333 = v326 + v332
	v335 = v328 + v332
	v337 = v327 - v332
	if base.Ui32(int32(3)) < base.Ui32(v337) {
		v326 = v333
		v327 = v337
		v328 = v335
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v341 = v333
	v342 = v337
	v343 = v335
	goto L74
L109:
	;
	goto L108
L110:
	;
	v348 = v341
	v349 = v342
	v350 = v343
	goto L111
L111:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v352)
	v354 = int32(1)
	v359 = v349 - v354
	if v359 != 0 {
		v348 = v348 + v354
		v349 = v359
		v350 = v350 + v354
		goto L111
	} else {
		goto L113
	}
L112:
	;
	goto L69
L113:
	;
	goto L112
L114:
	;
	m.G0 = v15 + int32(48)
	return
L115:
	;
	if base.Ui32(v38&int32(65535)) < base.Ui32(int32(2)) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v388 = int32(1)
	v390 = v40 - v388
	v393 = int32(0)
	if v38&int32(65535) != int32(2) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v400 = v388
	v402 = v393
	goto L120
L118:
	;
	v448 = v393
	goto L119
L119:
	;
	if v390&v388 == int32(0) {
		goto L114
	} else {
		goto L129
	}
L120:
	;
	v414 = v400<<(uint(int32(2))%32) + v47
	v416 = v414 - int32(4)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if base.Ui32(v417&int32(32767)) <= base.Ui32(v57) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v448 = v400 + int32(1)
	goto L119
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = (v417+v222)&int32(32767) | v417&int32(-32768)
	goto L124
L123:
	;
	goto L124
L124:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if base.Ui32(v428&int32(32767)) <= base.Ui32(v57) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = (v428+v222)&int32(32767) | v428&int32(-32768)
	goto L127
L126:
	;
	goto L127
L127:
	;
	v439 = int32(2)
	v442 = v402 + v439
	if v442 != v390&int32(-2) {
		v400 = v400 + v439
		v402 = v442
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L121
L129:
	;
	v462 = v47 + v448<<(uint(int32(2))%32)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	if base.Ui32(v57) < base.Ui32(v463&int32(32767)) {
		goto L114
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = (v463+v222)&int32(32767) | v463&int32(-32768)
	goto L114
L131:
	;
	return
L132:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v498
	F_errmsg(m, int32(58080), v15)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(511364), int32(1073), int32(359480))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L131
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg_internal(m, int32(59267), v15+int32(32))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L131
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(511364), int32(1077), int32(359480))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L131
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v57
	F_errmsg(m, int32(58142), v15+int32(16))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L131
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(511364), int32(1092), int32(359480))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L131
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v20) < base.Ui32(int32(24)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L78
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L78
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L78
	} else {
		goto L79
	}
L4:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v23) < base.Ui32(v20) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.Ui32(v25) < base.Ui32(v23) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v25) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if (v25+int32(7))&int32(32760) != v25 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v20 != int32(24) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = int32(base.Ui32(v20+int32(262120)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v41 = int32(0)
	goto L11
L11:
	;
	v42 = int32(65535)
	v43 = v41 & v42
	if base.Ui32(v43) <= base.Ui32((l1-int32(1))&v42) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v50 = l0 + int32(24)
	v55 = v50 + l1<<(uint(int32(2))%32) - int32(4)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v58 = int32(base.Ui32(v56) >> (uint(int32(17)) % 32))
	v60 = v56 & int32(32767)
	if base.Ui32(v60) < base.Ui32(v23) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v25) < base.Ui32(v60+v58) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v60 != (v60+int32(7))&int32(65528) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v69 = int32(7)
	v72 = (l3 + v69) & int32(-8)
	v77 = (v58 + v69) & int32(65528)
	v78 = v23 - v20 + v77
	if base.Ui32(v72) <= base.Ui32(v78) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v80 = v77 - v72
	if v77 != v72 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	m.G0 = v18 + int32(48)
	return base.B2i32(base.Ui32(v72) <= base.Ui32(v78))
L19:
	;
	v82 = l0 + v23
	v83 = v82 + v80
	v84 = v60 - v23
	if v83 == v82 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v286 = (v80 + v56) & int32(32767)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v286 | (v287&int32(98304) | l3<<(uint(int32(17))%32))
	if l3 != 0 {
		goto L75
	} else {
		goto L76
	}
L22:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v230 = v229 + v80
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v230)
	v238 = int32(1)
	goto L68
L23:
	;
	goto L22
L24:
	;
	v88 = v83 + v84
	if base.Ui32(v82-v88) <= base.Ui32(int32(0)-v84<<(uint(int32(1))%32)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v95 = F___memcpy(m, v83, v82, v84)
	mBase = m.M
	goto L22
L26:
	;
	goto L27
L27:
	;
	v98 = (v83 ^ v82) & int32(3)
	if base.Ui32(v83) < base.Ui32(v82) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v200 == int32(0) {
		goto L23
	} else {
		goto L64
	}
L29:
	;
	if base.Ui32(v178) <= base.Ui32(int32(3)) {
		v199 = v177
		v200 = v178
		v201 = v179
		goto L28
	} else {
		goto L60
	}
L30:
	;
	if v98 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v98 != 0 {
		v160 = v84
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v199 = v82
	v200 = v84
	v201 = v83
	goto L28
L34:
	;
	goto L35
L35:
	;
	if v83&int32(3) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v177 = v82
	v178 = v84
	v179 = v83
	goto L29
L37:
	;
	goto L38
L38:
	;
	v105 = v82
	v106 = v84
	v107 = v83
	goto L39
L39:
	;
	if v106 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L40:
	;
	v177 = v114
	v178 = v116
	v179 = v118
	goto L29
L41:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v111)
	v113 = int32(1)
	v114 = v105 + v113
	v116 = v106 - v113
	v118 = v107 + v113
	if v118&int32(3) != 0 {
		v105 = v114
		v106 = v116
		v107 = v118
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v160 == int32(0) {
		goto L23
	} else {
		goto L56
	}
L44:
	;
	if v88&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v125 = v84
	goto L48
L46:
	;
	v140 = v84
	goto L47
L47:
	;
	if base.Ui32(v140) <= base.Ui32(int32(3)) {
		v160 = v140
		goto L43
	} else {
		goto L52
	}
L48:
	;
	if v125 == int32(0) {
		goto L23
	} else {
		goto L50
	}
L49:
	;
	v140 = v131
	goto L47
L50:
	;
	v131 = v125 - int32(1)
	v132 = v83 + v131
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v131))))
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v134)
	if v132&int32(3) != 0 {
		v125 = v131
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v147 = v140
	goto L53
L53:
	;
	v151 = v147 - int32(4)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v82+v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v83+v151))) = v154
	if base.Ui32(int32(3)) < base.Ui32(v151) {
		v147 = v151
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v160 = v151
	goto L43
L55:
	;
	goto L54
L56:
	;
	v167 = v160
	goto L57
L57:
	;
	v171 = v167 - int32(1)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v171))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v171))) = uint8(v174)
	if v171 != 0 {
		v167 = v171
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L23
L59:
	;
	goto L58
L60:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L61
L61:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v188
	v190 = int32(4)
	v191 = v184 + v190
	v193 = v186 + v190
	v195 = v185 - v190
	if base.Ui32(int32(3)) < base.Ui32(v195) {
		v184 = v191
		v185 = v195
		v186 = v193
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v199 = v191
	v200 = v195
	v201 = v193
	goto L28
L63:
	;
	goto L62
L64:
	;
	v206 = v199
	v207 = v200
	v208 = v201
	goto L65
L65:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v210)
	v212 = int32(1)
	v217 = v207 - v212
	if v217 != 0 {
		v206 = v206 + v212
		v207 = v217
		v208 = v208 + v212
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L23
L67:
	;
	goto L66
L68:
	;
	v252 = v238<<(uint(int32(2))%32) + v50 - int32(4)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if base.Ui32(v253) < base.Ui32(int32(131072)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L21
L70:
	;
	if v238 != v43 {
		v238 = v238 + int32(1)
		goto L68
	} else {
		goto L73
	}
L71:
	;
	if base.Ui32(v60) < base.Ui32(v253&int32(32767)) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = (v253+v80)&int32(32767) | v253&int32(-32768)
	goto L70
L73:
	;
	goto L69
L74:
	;
	goto L18
L75:
	;
	v296 = F__emscripten_memcpy_bulkmem(m, l0+v286, l2, l3)
	mBase = m.M
	goto L77
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	return int32(0)
L79:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v329
	F_errmsg(m, int32(58080), v18)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(511364), int32(1426), int32(357262))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L78
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l1
	F_errmsg_internal(m, int32(59267), v18+int32(32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(511364), int32(1430), int32(357262))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v60
	F_errmsg(m, int32(58142), v18+int32(16))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L78
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(511364), int32(1442), int32(357262))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = base.I32_div_u_s(l1, int32(4069))
	v15 = base.I64_extend_i32_u(v12) << (uint(int64(32)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
	v19 = F_fsm_readbuf(m, l0, v9, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v19, int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(8159)) < base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(-1)
	goto L6
L5:
	;
	v29 = int32(base.Ui32(l2) >> (uint(int32(5)) % 32))
	goto L6
L6:
	;
	if v19 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v52 = v29 & int32(255)
	v57 = v50 + int32(28)
	v59 = l1 - v12*int32(4069) + int32(4095)
	v60 = v57 + v59
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v61 != v52 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v19^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L7
L9:
	;
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v50 = v44 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	if v151 != 0 {
		goto L42
	} else {
		goto L43
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v52)
	v71 = v59
	goto L15
L13:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.Ui32(v63) < base.Ui32(v52) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v151 = int32(0)
	goto L11
L15:
	;
	v73 = int32(1)
	v74 = v71 - v73
	v75 = int32(2)
	v76 = base.I32_div_s(v74, v75)
	v78 = v76 << (uint(v73) % 32)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v78)+1)))
	v82 = v78 + v75
	if base.Ui32(v82) <= base.Ui32(int32(8163)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.Ui32(v101) < base.Ui32(v52) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v86 = v80 & int32(255)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v82))))
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v91 = v80
	goto L19
L19:
	;
	v93 = v57 + v76
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v94 != v91&int32(255) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v90 = v86
	goto L22
L21:
	;
	v90 = v88
	goto L22
L22:
	;
	v91 = v90
	goto L19
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v91)
	if int32(1) < v74 {
		v71 = v76
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L16
L26:
	;
	goto L25
L27:
	;
	v107 = int32(4094)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v151 = int32(1)
	goto L11
L30:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v107) {
		v129 = int32(0)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v130 = v57 + v107
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v131 != v129&int32(255) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v114 = v107 << (uint(int32(1)) % 32)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v114)+1)))
	if v107 == int32(4081) {
		v129 = v116
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v120 = v116 & int32(255)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+(v114+int32(2))))))
	if base.Ui32(v124) < base.Ui32(v120) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = v120
	goto L37
L36:
	;
	v126 = v124
	goto L37
L37:
	;
	v129 = v126
	goto L32
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
	goto L40
L39:
	;
	goto L40
L40:
	;
	if v107 != 0 {
		v107 = v107 - int32(1)
		goto L30
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	F_MarkBufferDirtyHint(m, v19, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_UnlockReleaseBuffer(m, v19)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	m.G0 = v9 + int32(16)
	return
}
