package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v105 int32
	_ = v105
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v6
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(112)
	return v459
L2:
	;
	v26 = l0
	v31 = v25
	v33 = v6
	goto L5
L3:
	;
	v437 = l1
	goto L4
L4:
	;
	if v437&int32(-32801) != int32(14) {
		goto L91
	} else {
		goto L92
	}
L5:
	;
	v43 = v31 & int32(255)
	goto L7
L6:
	;
	v170 = l1
	v179 = v6
	v184 = v6
	goto L29
L7:
	;
	if base.B2i32(base.Ui32(v43-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v43|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v57 != 0 {
		v26 = v26 + int32(1)
		v31 = v57
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+v33<<(uint(int32(2))%32)))) = v26
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.Ui32((v65-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v459 = int32(-1)
	goto L1
L12:
	;
	if v137&int32(255) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v72 = v26
	goto L16
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(int32(25)) < base.Ui32((v65|int32(32)-int32(97))&int32(255)) {
		v137 = v65
		v138 = v26
		goto L12
	} else {
		goto L19
	}
L16:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32((v90-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v72 = v72 + int32(1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v137 = v90
	v138 = v72
	goto L12
L19:
	;
	v105 = v26
	goto L20
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if base.Ui32((v123|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v105 = v105 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v137 = v123
	v138 = v105
	goto L12
L22:
	;
	goto L21
L23:
	;
	goto L6
L24:
	;
	v168 = v33 + int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v154)
	v157 = v33 + int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v158 == v154 {
		v168 = v157
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v33) < base.Ui32(int32(24)) {
		v26 = v138 + int32(1)
		v31 = v158
		v33 = v157
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v168 = v157
	goto L23
L29:
	;
	v186 = v179 << (uint(int32(2)) % 32)
	v187 = v19 + v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v188))))
	if base.Ui32(int32(25)) < base.Ui32((v189|int32(32)-int32(97))&int32(255)) {
		v371 = v170
		v385 = v184
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v391 = v371
	v395 = int32(0)
	goto L80
L31:
	;
	v387 = v179 + int32(1)
	if v387 != v168 {
		v170 = v371
		v179 = v387
		v184 = v385
		goto L29
	} else {
		goto L79
	}
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v186)+uint32(_c_F_DecodeDate[0])))
	if v198 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+uint32(_c_F_DecodeDate[0]))) = v339
	v353 = int32(-1)
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+11)))
	switch v354 - int32(1) {
	case 0:
		goto L77
	default:
		v459 = v353
		goto L1
	case 7:
		v371 = v170
		v385 = v184
		goto L31
	}
L34:
	;
	goto L39
L35:
	;
	goto L36
L36:
	;
	v249 = int32(_a_F_DecodeDate_0)
	v255 = int32(_a_F_DecodeDate_1)
	goto L51
L37:
	;
	if v236-v237 == int32(0) {
		v339 = v198
		goto L33
	} else {
		goto L50
	}
L39:
	;
	goto L40
L40:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v205 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v206 = v188
	v207 = v198
	v208 = int32(10)
	v209 = v205
	goto L45
L42:
	;
	v232 = v198
	v236 = int32(0)
	goto L43
L43:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	goto L37
L44:
	;
	v232 = v227
	v236 = v229
	goto L43
L45:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if base.B2i32(v209 != v211)|base.B2i32(v211 == int32(0)) != 0 {
		v227 = v207
		v229 = v209
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v227 = v221
	v229 = int32(0)
	goto L44
L47:
	;
	v217 = v208 - int32(1)
	if v217 == int32(0) {
		v227 = v207
		v229 = v209
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v220 = int32(1)
	v221 = v207 + v220
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v222 != 0 {
		v206 = v206 + v220
		v207 = v221
		v208 = v217
		v209 = v222
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L36
L51:
	;
	v270 = v249 + (v255-v249)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v271 = int32(*(*int8)(unsafe.Add(mBase, uint32(v270))))
	v272 = v189 - v271
	if v272 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v459 = int32(-1)
	goto L1
L53:
	;
	goto L58
L54:
	;
	v323 = v272
	goto L55
L55:
	;
	v327 = base.B2i32(v323 < int32(0))
	if v323 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L56:
	;
	if v314 == int32(0) {
		v339 = v270
		goto L33
	} else {
		goto L69
	}
L58:
	;
	goto L59
L59:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v281 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v282 = v188
	v283 = v270
	v284 = int32(10)
	v285 = v281
	goto L64
L61:
	;
	v308 = v270
	v312 = int32(0)
	goto L62
L62:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v314 = v312 - v313
	goto L56
L63:
	;
	v308 = v303
	v312 = v305
	goto L62
L64:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if base.B2i32(v285 != v287)|base.B2i32(v287 == int32(0)) != 0 {
		v303 = v283
		v305 = v285
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v303 = v297
	v305 = int32(0)
	goto L63
L66:
	;
	v293 = v284 - int32(1)
	if v293 == int32(0) {
		v303 = v283
		v305 = v285
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v296 = int32(1)
	v297 = v283 + v296
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v298 != 0 {
		v282 = v282 + v296
		v283 = v297
		v284 = v293
		v285 = v298
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v323 = v314
	goto L55
L70:
	;
	v328 = v270 - int32(16)
	goto L72
L71:
	;
	v328 = v255
	goto L72
L72:
	;
	if v323 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v331 = v249
	goto L75
L74:
	;
	v331 = v270 + int32(16)
	goto L75
L75:
	;
	if base.Ui32(v331) <= base.Ui32(v328) {
		v249 = v331
		v255 = v328
		goto L51
	} else {
		goto L76
	}
L76:
	;
	goto L52
L77:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v357
	if v170&int32(2) != 0 {
		v459 = v353
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v362 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v361 | v362
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = int32(0)
	v371 = v170 | v362
	v385 = int32(1)
	goto L31
L79:
	;
	goto L30
L80:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v19+v395<<(uint(int32(2))%32))))
	if v409 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v437 = v430
	goto L4
L82:
	;
	v411 = F_strlen(m, v409)
	mBase = m.M
	if v411 <= int32(0) {
		v459 = int32(-1)
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v430 = v391
	goto L84
L84:
	;
	v434 = v395 + int32(1)
	if v434 != v168 {
		v391 = v430
		v395 = v434
		goto L80
	} else {
		goto L90
	}
L85:
	;
	v418 = F_DecodeNumber(m, v411, v409, v385, v391, v19+int32(104), l4, v19+int32(108), l3)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	return int32(0)
L87:
	;
	if v418 != 0 {
		v459 = v418
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v423&v391 != 0 {
		v459 = int32(-1)
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v425 | v423
	v430 = v391 | v423
	goto L84
L90:
	;
	goto L81
L91:
	;
	v458 = int32(-1)
	goto L93
L92:
	;
	v458 = int32(0)
	goto L93
L93:
	;
	v459 = v458
	goto L1
}
func F_EncodeDateOnly(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
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
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	switch l1 - int32(1) {
	case 0, 3:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v6 {
			v11 = v6
		} else {
			v11 = int32(1) - v6
		}
		v12 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v11)) == int32(0) {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v23)
			v38 = l2 + int32(2)
		} else {
			v27 = F_pg_ultoa_n(m, v11, l2)
			mBase = m.M
			if v12 <= v27 {
				v38 = l2 + v27
			} else {
				v30 = l2 + v12
				if v27 != 0 {
					base.MemoryCopy(m, v30-v27, l2, v27)
				} else {
				}
				v33 = v12 - v27
				if v33 != 0 {
					base.MemoryFill(m, l2, int32(48), v33)
				} else {
				}
				v38 = v30
			}
		}
		v39 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v39)
		v42 = v38 + int32(1)
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v44 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v43)) == int32(0) {
			v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v42))) = uint16(v55)
			v70 = v38 + int32(3)
		} else {
			v59 = F_pg_ultoa_n(m, v43, v42)
			mBase = m.M
			if v44 <= v59 {
				v70 = v42 + v59
			} else {
				v62 = v38 + int32(3)
				if v59 != 0 {
					base.MemoryCopy(m, v62-v59, v42, v59)
				} else {
				}
				v65 = v44 - v59
				if v65 != 0 {
					base.MemoryFill(m, v42, int32(48), v65)
				} else {
				}
				v70 = v62
			}
		}
		v71 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
		v74 = v70 + int32(1)
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v76 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v75)) == int32(0) {
			v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v74))) = uint16(v87)
			v102 = v70 + int32(3)
		} else {
			v91 = F_pg_ultoa_n(m, v75, v74)
			mBase = m.M
			if v76 <= v91 {
				v102 = v74 + v91
			} else {
				v94 = v70 + int32(3)
				if v91 != 0 {
					base.MemoryCopy(m, v94-v91, v74, v91)
				} else {
				}
				v97 = v76 - v91
				if v97 != 0 {
					base.MemoryFill(m, v74, int32(48), v97)
				} else {
				}
				v102 = v94
			}
		}
		v420 = v102
	case 1:
		v106 = *(*int32)(unsafe.Add(mBase, _c_F_EncodeDateOnly[1]))
		v108 = base.B2i32(v106 == int32(1))
		if v106 == int32(1) {
			v109 = int32(12)
		} else {
			v109 = int32(16)
		}
		v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v109)))
		v112 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v111)) == int32(0) {
			v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v123)
			v138 = l2 + int32(2)
		} else {
			v127 = F_pg_ultoa_n(m, v111, l2)
			mBase = m.M
			if v112 <= v127 {
				v138 = l2 + v127
			} else {
				v130 = l2 + v112
				if v127 != 0 {
					base.MemoryCopy(m, v130-v127, l2, v127)
				} else {
				}
				v133 = v112 - v127
				if v133 != 0 {
					base.MemoryFill(m, l2, int32(48), v133)
				} else {
				}
				v138 = v130
			}
		}
		v139 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v139)
		v142 = v138 + int32(1)
		if v106 == int32(1) {
			v145 = int32(16)
		} else {
			v145 = int32(12)
		}
		v147 = *(*int32)(unsafe.Add(mBase, uint32(l0+v145)))
		v148 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v147)) == int32(0) {
			v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v142))) = uint16(v159)
			v174 = v138 + int32(3)
		} else {
			v163 = F_pg_ultoa_n(m, v147, v142)
			mBase = m.M
			if v148 <= v163 {
				v174 = v142 + v163
			} else {
				v166 = v138 + int32(3)
				if v163 != 0 {
					base.MemoryCopy(m, v166-v163, v142, v163)
				} else {
				}
				v169 = v148 - v163
				if v169 != 0 {
					base.MemoryFill(m, v142, int32(48), v169)
				} else {
				}
				v174 = v166
			}
		}
		v175 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v175)
		v177 = int32(1)
		v178 = v174 + v177
		v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v179 {
			v184 = v179
		} else {
			v184 = v177 - v179
		}
		v185 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v184)) == int32(0) {
			v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v178))) = uint16(v196)
			v211 = v174 + int32(3)
		} else {
			v200 = F_pg_ultoa_n(m, v184, v178)
			mBase = m.M
			if v185 <= v200 {
				v211 = v178 + v200
			} else {
				v203 = v174 + int32(5)
				if v200 != 0 {
					base.MemoryCopy(m, v203-v200, v178, v200)
				} else {
				}
				v206 = v185 - v200
				if v206 != 0 {
					base.MemoryFill(m, v178, int32(48), v206)
				} else {
				}
				v211 = v203
			}
		}
		v420 = v211
	case 2:
		v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v213 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v212)) == int32(0) {
			v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v224)
			v239 = l2 + int32(2)
		} else {
			v228 = F_pg_ultoa_n(m, v212, l2)
			mBase = m.M
			if v213 <= v228 {
				v239 = l2 + v228
			} else {
				v231 = l2 + v213
				if v228 != 0 {
					base.MemoryCopy(m, v231-v228, l2, v228)
				} else {
				}
				v234 = v213 - v228
				if v234 != 0 {
					base.MemoryFill(m, l2, int32(48), v234)
				} else {
				}
				v239 = v231
			}
		}
		v240 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v240)
		v243 = v239 + int32(1)
		v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v245 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v244)) == int32(0) {
			v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v243))) = uint16(v256)
			v271 = v239 + int32(3)
		} else {
			v260 = F_pg_ultoa_n(m, v244, v243)
			mBase = m.M
			if v245 <= v260 {
				v271 = v243 + v260
			} else {
				v263 = v239 + int32(3)
				if v260 != 0 {
					base.MemoryCopy(m, v263-v260, v243, v260)
				} else {
				}
				v266 = v245 - v260
				if v266 != 0 {
					base.MemoryFill(m, v243, int32(48), v266)
				} else {
				}
				v271 = v263
			}
		}
		v272 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v272)
		v274 = int32(1)
		v275 = v271 + v274
		v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v276 {
			v281 = v276
		} else {
			v281 = v274 - v276
		}
		v282 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v281)) == int32(0) {
			v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v275))) = uint16(v293)
			v308 = v271 + int32(3)
		} else {
			v297 = F_pg_ultoa_n(m, v281, v275)
			mBase = m.M
			if v282 <= v297 {
				v308 = v275 + v297
			} else {
				v300 = v271 + int32(5)
				if v297 != 0 {
					base.MemoryCopy(m, v300-v297, v275, v297)
				} else {
				}
				v303 = v282 - v297
				if v303 != 0 {
					base.MemoryFill(m, v275, int32(48), v303)
				} else {
				}
				v308 = v300
			}
		}
		v420 = v308
	default:
		v312 = *(*int32)(unsafe.Add(mBase, _c_F_EncodeDateOnly[1]))
		v314 = base.B2i32(v312 == int32(1))
		if v312 == int32(1) {
			v315 = int32(12)
		} else {
			v315 = int32(16)
		}
		v317 = *(*int32)(unsafe.Add(mBase, uint32(l0+v315)))
		v318 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v317)) == int32(0) {
			v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v317<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v329)
			v344 = l2 + int32(2)
		} else {
			v333 = F_pg_ultoa_n(m, v317, l2)
			mBase = m.M
			if v318 <= v333 {
				v344 = l2 + v333
			} else {
				v336 = l2 + v318
				if v333 != 0 {
					base.MemoryCopy(m, v336-v333, l2, v333)
				} else {
				}
				v339 = v318 - v333
				if v339 != 0 {
					base.MemoryFill(m, l2, int32(48), v339)
				} else {
				}
				v344 = v336
			}
		}
		v345 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v344))) = uint8(v345)
		v348 = v344 + int32(1)
		if v312 == int32(1) {
			v351 = int32(16)
		} else {
			v351 = int32(12)
		}
		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0+v351)))
		v354 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v353)) == int32(0) {
			v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v348))) = uint16(v365)
			v380 = v344 + int32(3)
		} else {
			v369 = F_pg_ultoa_n(m, v353, v348)
			mBase = m.M
			if v354 <= v369 {
				v380 = v348 + v369
			} else {
				v372 = v344 + int32(3)
				if v369 != 0 {
					base.MemoryCopy(m, v372-v369, v348, v369)
				} else {
				}
				v375 = v354 - v369
				if v375 != 0 {
					base.MemoryFill(m, v348, int32(48), v375)
				} else {
				}
				v380 = v372
			}
		}
		v381 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v380))) = uint8(v381)
		v383 = int32(1)
		v384 = v380 + v383
		v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v385 {
			v390 = v385
		} else {
			v390 = v383 - v385
		}
		v391 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v390)) == int32(0) {
			v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateOnly[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v384))) = uint16(v402)
			v417 = v380 + int32(3)
		} else {
			v406 = F_pg_ultoa_n(m, v390, v384)
			mBase = m.M
			if v391 <= v406 {
				v417 = v384 + v406
			} else {
				v409 = v380 + int32(5)
				if v406 != 0 {
					base.MemoryCopy(m, v409-v406, v384, v406)
				} else {
				}
				v412 = v391 - v406
				if v412 != 0 {
					base.MemoryFill(m, v384, int32(48), v412)
				} else {
				}
				v417 = v409
			}
		}
		v420 = v417
	}
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v421 <= int32(0) {
		v425 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateOnly[2])))
		*(*uint8)(unsafe.Add(mBase, uint32(v420)+2)) = uint8(v425)
		v428 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateOnly[3])))
		*(*uint16)(unsafe.Add(mBase, uint32(v420))) = uint16(v428)
		v432 = v420 + int32(3)
	} else {
		v432 = v420
	}
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v433)
	return
}
func F_date_cmp_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_date_cmp_timestamptz_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_date_cmp_timestamptz[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return v65
}
func F_date_eq_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
	} else {
		if v6 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
		} else {
			if int32(106751982) < v6 {
				v24 = int32(0)
			} else {
				v18 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
			}
		}
	}
	return v24
}
func F_date_ge_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_date_ge_timestamptz_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_date_ge_timestamptz[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return int32(base.Ui32(v65^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_date_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	v9 = m.G0
	v11 = v9 - int32(448)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = v11 + int32(288)
	v21 = v11 + int32(176)
	v24 = F_ParseDateTime(m, v14, v11+int32(32), int32(129), v19, v21, v11+int32(388))
	mBase = m.M
	if v24 == int32(0) {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+388))
		v38 = F_DecodeDateTime(m, v19, v21, v27, v11+int32(392), v11+int32(400), v11+int32(444), v11+int32(396), v11+int32(24))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v38 == int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+392))
				switch v54 - int32(2) {
				case 0:
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+420))
					if v120 <= int32(-4713) {
						if v120 != int32(-4713) {
							v137 = int32(0)
							v138 = F_errsave_start(m, v13)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int32(0)
							} else {
								if v138 == int32(0) {
									v210 = v137
									m.G0 = v11 + int32(448)
									return v210
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
										F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												v210 = v137
												m.G0 = v11 + int32(448)
												return v210
											}
										}
									}
								}
							}
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
							if v125 <= int32(10) {
								v137 = int32(0)
								v138 = F_errsave_start(m, v13)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									if v138 == int32(0) {
										v210 = v137
										m.G0 = v11 + int32(448)
										return v210
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
											F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													v210 = v137
													m.G0 = v11 + int32(448)
													return v210
												}
											}
										}
									}
								}
							} else {
								v156 = v125
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
								v162 = base.B2i32(int32(2) < v156)
								if int32(2) < v156 {
									v163 = int32(_a_F_date_in_3)
								} else {
									v163 = int32(_a_F_date_in_4)
								}
								v164 = v163 + v120
								v169 = base.I32_div_s(v164, int32(4))
								v172 = base.I32_div_s(v164, int32(-100))
								v175 = base.I32_div_s(v164, int32(400))
								if int32(2) < v156 {
									v179 = int32(1)
								} else {
									v179 = int32(13)
								}
								v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
								v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
									v190 = int32(0)
									v191 = F_errsave_start(m, v13)
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										if v191 == int32(0) {
											v210 = v190
											m.G0 = v11 + int32(448)
											return v210
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
												F_errmsg(m, int32(_a_F_date_in_0), v11)
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return int32(0)
													} else {
														v210 = v190
														m.G0 = v11 + int32(448)
														return v210
													}
												}
											}
										}
									}
								} else {
									v210 = v187 - int32(_a_F_date_in_7)
									m.G0 = v11 + int32(448)
									return v210
								}
							}
						}
					} else {
						if v120 <= int32(_a_F_date_in_8) {
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
							v156 = v130
							v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
							v162 = base.B2i32(int32(2) < v156)
							if int32(2) < v156 {
								v163 = int32(_a_F_date_in_3)
							} else {
								v163 = int32(_a_F_date_in_4)
							}
							v164 = v163 + v120
							v169 = base.I32_div_s(v164, int32(4))
							v172 = base.I32_div_s(v164, int32(-100))
							v175 = base.I32_div_s(v164, int32(400))
							if int32(2) < v156 {
								v179 = int32(1)
							} else {
								v179 = int32(13)
							}
							v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
							v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
								v190 = int32(0)
								v191 = F_errsave_start(m, v13)
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return int32(0)
								} else {
									if v191 == int32(0) {
										v210 = v190
										m.G0 = v11 + int32(448)
										return v210
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
											F_errmsg(m, int32(_a_F_date_in_0), v11)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return int32(0)
												} else {
													v210 = v190
													m.G0 = v11 + int32(448)
													return v210
												}
											}
										}
									}
								}
							} else {
								v210 = v187 - int32(_a_F_date_in_7)
								m.G0 = v11 + int32(448)
								return v210
							}
						} else {
							if v120 != int32(_a_F_date_in_9) {
								v137 = int32(0)
								v138 = F_errsave_start(m, v13)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									if v138 == int32(0) {
										v210 = v137
										m.G0 = v11 + int32(448)
										return v210
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
											F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													v210 = v137
													m.G0 = v11 + int32(448)
													return v210
												}
											}
										}
									}
								}
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
								if v133 < int32(6) {
									v156 = v133
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
									v162 = base.B2i32(int32(2) < v156)
									if int32(2) < v156 {
										v163 = int32(_a_F_date_in_3)
									} else {
										v163 = int32(_a_F_date_in_4)
									}
									v164 = v163 + v120
									v169 = base.I32_div_s(v164, int32(4))
									v172 = base.I32_div_s(v164, int32(-100))
									v175 = base.I32_div_s(v164, int32(400))
									if int32(2) < v156 {
										v179 = int32(1)
									} else {
										v179 = int32(13)
									}
									v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
									v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
										v190 = int32(0)
										v191 = F_errsave_start(m, v13)
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int32(0)
										} else {
											if v191 == int32(0) {
												v210 = v190
												m.G0 = v11 + int32(448)
												return v210
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													F_errmsg(m, int32(_a_F_date_in_0), v11)
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
														mBase = m.M
														v206 = m.ExcPending
														if v206 != 0 {
															return int32(0)
														} else {
															v210 = v190
															m.G0 = v11 + int32(448)
															return v210
														}
													}
												}
											}
										}
									} else {
										v210 = v187 - int32(_a_F_date_in_7)
										m.G0 = v11 + int32(448)
										return v210
									}
								} else {
									v137 = int32(0)
									v138 = F_errsave_start(m, v13)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										if v138 == int32(0) {
											v210 = v137
											m.G0 = v11 + int32(448)
											return v210
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
												F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														v210 = v137
														m.G0 = v11 + int32(448)
														return v210
													}
												}
											}
										}
									}
								}
							}
						}
					}
				default:
					F_DateTimeParseError(m, int32(-1), v11+int32(24), v14, int32(_a_F_date_in_10), v13)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
						v210 = int32(0)
						m.G0 = v11 + int32(448)
						return v210
					}
				case 7:
					v210 = int32(-2147483648)
					m.G0 = v11 + int32(448)
					return v210
				case 8:
					v210 = int32(2147483647)
					m.G0 = v11 + int32(448)
					return v210
				case 9:
					v67 = m.G0
					v69 = v67 - int32(16)
					m.G0 = v69
					*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = int64(0)
					v75 = F_pg_gmtime(m, v69+int32(8))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						if v75 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_date_in_11), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_date_in_12), int32(2176), int32(_a_F_date_in_13))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v93 = v11 + int32(400)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v94
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v96
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v98
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v100
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v102
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v94 + int32(1900)
							*(*int32)(unsafe.Add(mBase, uint32(v93))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v96 + int32(1)
							m.G0 = v69 + int32(16)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+420))
							if v120 <= int32(-4713) {
								if v120 != int32(-4713) {
									v137 = int32(0)
									v138 = F_errsave_start(m, v13)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										if v138 == int32(0) {
											v210 = v137
											m.G0 = v11 + int32(448)
											return v210
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
												F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														v210 = v137
														m.G0 = v11 + int32(448)
														return v210
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
									if v125 <= int32(10) {
										v137 = int32(0)
										v138 = F_errsave_start(m, v13)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											if v138 == int32(0) {
												v210 = v137
												m.G0 = v11 + int32(448)
												return v210
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v144 = m.ExcPending
												if v144 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
													F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															v210 = v137
															m.G0 = v11 + int32(448)
															return v210
														}
													}
												}
											}
										}
									} else {
										v156 = v125
										v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
										v162 = base.B2i32(int32(2) < v156)
										if int32(2) < v156 {
											v163 = int32(_a_F_date_in_3)
										} else {
											v163 = int32(_a_F_date_in_4)
										}
										v164 = v163 + v120
										v169 = base.I32_div_s(v164, int32(4))
										v172 = base.I32_div_s(v164, int32(-100))
										v175 = base.I32_div_s(v164, int32(400))
										if int32(2) < v156 {
											v179 = int32(1)
										} else {
											v179 = int32(13)
										}
										v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
										v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
										if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
											v190 = int32(0)
											v191 = F_errsave_start(m, v13)
											mBase = m.M
											v192 = m.ExcPending
											if v192 != 0 {
												return int32(0)
											} else {
												if v191 == int32(0) {
													v210 = v190
													m.G0 = v11 + int32(448)
													return v210
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v197 = m.ExcPending
													if v197 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
														F_errmsg(m, int32(_a_F_date_in_0), v11)
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
															mBase = m.M
															v206 = m.ExcPending
															if v206 != 0 {
																return int32(0)
															} else {
																v210 = v190
																m.G0 = v11 + int32(448)
																return v210
															}
														}
													}
												}
											}
										} else {
											v210 = v187 - int32(_a_F_date_in_7)
											m.G0 = v11 + int32(448)
											return v210
										}
									}
								}
							} else {
								if v120 <= int32(_a_F_date_in_8) {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
									v156 = v130
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
									v162 = base.B2i32(int32(2) < v156)
									if int32(2) < v156 {
										v163 = int32(_a_F_date_in_3)
									} else {
										v163 = int32(_a_F_date_in_4)
									}
									v164 = v163 + v120
									v169 = base.I32_div_s(v164, int32(4))
									v172 = base.I32_div_s(v164, int32(-100))
									v175 = base.I32_div_s(v164, int32(400))
									if int32(2) < v156 {
										v179 = int32(1)
									} else {
										v179 = int32(13)
									}
									v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
									v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
										v190 = int32(0)
										v191 = F_errsave_start(m, v13)
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int32(0)
										} else {
											if v191 == int32(0) {
												v210 = v190
												m.G0 = v11 + int32(448)
												return v210
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													F_errmsg(m, int32(_a_F_date_in_0), v11)
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
														mBase = m.M
														v206 = m.ExcPending
														if v206 != 0 {
															return int32(0)
														} else {
															v210 = v190
															m.G0 = v11 + int32(448)
															return v210
														}
													}
												}
											}
										}
									} else {
										v210 = v187 - int32(_a_F_date_in_7)
										m.G0 = v11 + int32(448)
										return v210
									}
								} else {
									if v120 != int32(_a_F_date_in_9) {
										v137 = int32(0)
										v138 = F_errsave_start(m, v13)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											if v138 == int32(0) {
												v210 = v137
												m.G0 = v11 + int32(448)
												return v210
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v144 = m.ExcPending
												if v144 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
													F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															v210 = v137
															m.G0 = v11 + int32(448)
															return v210
														}
													}
												}
											}
										}
									} else {
										v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
										if v133 < int32(6) {
											v156 = v133
											v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
											v162 = base.B2i32(int32(2) < v156)
											if int32(2) < v156 {
												v163 = int32(_a_F_date_in_3)
											} else {
												v163 = int32(_a_F_date_in_4)
											}
											v164 = v163 + v120
											v169 = base.I32_div_s(v164, int32(4))
											v172 = base.I32_div_s(v164, int32(-100))
											v175 = base.I32_div_s(v164, int32(400))
											if int32(2) < v156 {
												v179 = int32(1)
											} else {
												v179 = int32(13)
											}
											v184 = base.I32_div_s((v179+v156)*int32(_a_F_date_in_5), int32(256))
											v187 = v157 + v164*int32(365) + v169 + v172 + v175 + v184 - int32(_a_F_date_in_6)
											if base.Ui32(int32(2147483494)) <= base.Ui32(v187) {
												v190 = int32(0)
												v191 = F_errsave_start(m, v13)
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													if v191 == int32(0) {
														v210 = v190
														m.G0 = v11 + int32(448)
														return v210
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
															F_errmsg(m, int32(_a_F_date_in_0), v11)
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(176), int32(_a_F_date_in_2))
																mBase = m.M
																v206 = m.ExcPending
																if v206 != 0 {
																	return int32(0)
																} else {
																	v210 = v190
																	m.G0 = v11 + int32(448)
																	return v210
																}
															}
														}
													}
												}
											} else {
												v210 = v187 - int32(_a_F_date_in_7)
												m.G0 = v11 + int32(448)
												return v210
											}
										} else {
											v137 = int32(0)
											v138 = F_errsave_start(m, v13)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return int32(0)
											} else {
												if v138 == int32(0) {
													v210 = v137
													m.G0 = v11 + int32(448)
													return v210
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v144 = m.ExcPending
													if v144 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
														F_errmsg(m, int32(_a_F_date_in_0), v11+int32(16))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v13, int32(_a_F_date_in_1), int32(168), int32(_a_F_date_in_2))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																v210 = v137
																m.G0 = v11 + int32(448)
																return v210
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
			} else {
				v44 = v38
				F_DateTimeParseError(m, v44, v11+int32(24), v14, int32(_a_F_date_in_10), v13)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
					v210 = int32(0)
					m.G0 = v11 + int32(448)
					return v210
				}
			}
		}
	} else {
		v44 = v24
		F_DateTimeParseError(m, v44, v11+int32(24), v14, int32(_a_F_date_in_10), v13)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
			v210 = int32(0)
			m.G0 = v11 + int32(448)
			return v210
		}
	}
}
func F_date_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4 < v3 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_date_lt_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
			}
		}
	}
}
