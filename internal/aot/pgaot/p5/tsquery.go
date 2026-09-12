package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parse_tsquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int64
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v20 = l3 & int32(2)
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(1524)
	goto L3
L2:
	;
	v21 = int32(1525)
	goto L3
L3:
	;
	v23 = l3 & int32(1)
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(1523)
	goto L6
L5:
	;
	v24 = v21
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v24
	v26 = int32(3)
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = int32(7)
	goto L9
L8:
	;
	v29 = v26
	goto L9
L9:
	;
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = v26
	goto L12
L11:
	;
	v30 = v29
	goto L12
L12:
	;
	if l4 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v34 = base.B2i32(v31 != int32(447))
	goto L15
L14:
	;
	v34 = int32(1)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = int64(12884901888)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l0
	v42 = F_init_tsvector_parser(m, l0, v30, l4)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v13)+60)) = int64(64)
	v50 = F_palloc(m, int32(64))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v50
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v54)
	F_makepol_1(m, v13+int32(28), l1, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	F_close_tsvector_parser(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if l4 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L16
	} else {
		goto L120
	}
L22:
	;
	m.G0 = v13 + int32(80)
	return v445
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v71 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v65 != int32(447) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v69 != 0 {
		v445 = int32(0)
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v34 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v102 = base.I32_div_u_s(int32(1073741815)-v99, int32(12))
	if base.Ui32(v102) < base.Ui32(v97) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v93 = F_palloc(m, int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L36
	}
L31:
	;
	v78 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	if v78 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v82
	F_errmsg(m, int32(724575), v13)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(492089), int32(880), int32(15299))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = int64(32)
	v445 = v93
	goto L22
L37:
	;
	v104 = int32(0)
	v105 = F_errsave_start(m, l4)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v125 = v99 + v97*int32(12) + int32(8)
	v126 = F_palloc0(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L16
	} else {
		goto L45
	}
L40:
	;
	if v105 == int32(0) {
		v445 = v104
		goto L22
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(400000), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, l4, int32(492089), int32(890), int32(15299))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v445 = v104
	goto L22
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v125 << (uint(int32(2)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v131 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v134 = v132
	goto L48
L47:
	;
	v134 = int32(0)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v134
	v137 = v126 + int32(8)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v138 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v139 < v140 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v220 = v134
	goto L51
L51:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	if v225 != 0 {
		goto L67
	} else {
		goto L68
	}
L52:
	;
	v147 = v139
	goto L55
L53:
	;
	goto L54
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v220 = v209
	goto L51
L55:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v147<<(uint(int32(2))%32))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	switch v158 - int32(1) {
	case 0:
		goto L58
	case 1:
		goto L60
	case 2:
		goto L61
	default:
		goto L59
	}
L56:
	;
	goto L54
L57:
	;
	v196 = v147 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v196 < v197 {
		v147 = v196
		goto L55
	} else {
		goto L65
	}
L58:
	;
	v189 = v137 + v147*int32(12)
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v189))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v192
	goto L57
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L62
	}
L60:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v137+v147*int32(12)))) = v169
	goto L57
L61:
	;
	v164 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v147*int32(12)))) = uint8(v164)
	goto L57
L62:
	;
	v175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v157))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v175
	F_errmsg_internal(m, int32(485084), v13+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(492089), int32(917), int32(15299))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	goto L56
L66:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	F_pfree(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L70
	}
L67:
	;
	v226 = F__emscripten_memcpy_bulkmem(m, v137+v220*int32(12), v224, v225)
	mBase = m.M
	goto L69
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v232
	F_findoprnd_recurse(m, v137, v13+int32(76), v231, v13+int32(27))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if v231 != v242 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	if v244 == int32(0) {
		v445 = v126
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v247 = m.G0
	v249 = v247 - int32(32)
	m.G0 = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v251 == int32(0) {
		v432 = v126
		goto L74
	} else {
		goto L75
	}
L74:
	;
	m.G0 = v249 + int32(32)
	v445 = v432
	goto L22
L75:
	;
	v255 = v126 + int32(8)
	v256 = F_maketree(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	v262 = F_clean_stopword_intree(m, v256, v249+int32(16), v249+int32(12))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	if v262 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v34 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v288 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v292 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L81:
	;
	v284 = F_palloc(m, int32(8))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L16
	} else {
		goto L87
	}
L82:
	;
	v270 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	if v270 == int32(0) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(449980), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(495411), int32(409), int32(171741))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	goto L81
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v284))) = int64(32)
	v432 = v284
	goto L74
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v249)+24)) = int64(16)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v329 = int32(1)
	if base.Ui32((v328-v329)&int32(255)) <= base.Ui32(v329) {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	goto L88
L90:
	;
	v295 = v262
	v296 = v288
	v297 = v291
	goto L93
L91:
	;
	v311 = v288
	v312 = v291
	goto L92
L92:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	v320 = v311
	v322 = v314&int32(4095) + int32(1)
	goto L89
L93:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v300 = F_calcstrlen(m, v299)
	mBase = m.M
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v301 == int32(1) {
		v320 = v296
		v322 = v300
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v311 = v304
	v312 = v306
	goto L92
L95:
	;
	v304 = v296 + v300
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v307 != int32(1) {
		v295 = v305
		v296 = v304
		v297 = v306
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v336 = F_palloc(m, int32(192))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L100
	}
L98:
	;
	v345 = int32(0)
	v346 = int32(0)
	goto L99
L99:
	;
	v348 = v346 * int32(12)
	v351 = v320 + v322 + v348 + int32(8)
	v352 = F_palloc(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+20)) = v336
	F_plainnode(m, v249+int32(20), v262)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v249)+28))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	v345 = v344
	v346 = v343
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v351 << (uint(int32(2)) % 32)
	v359 = v352 + int32(8)
	if v348 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if int32(0) < v346 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v360 = F__emscripten_memcpy_bulkmem(m, v359, v345, v348)
	mBase = m.M
	v361 = v360
	goto L106
L105:
	;
	v361 = v359
	goto L106
L106:
	;
	goto L103
L107:
	;
	v369 = v361 + v348
	v374 = v346
	v375 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	v432 = v352
	goto L74
L110:
	;
	v378 = v361 + v375*int32(12)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v379 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L109
L112:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v383 = int32(12)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v391 = v386 & int32(4095)
	if v391 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v416 = v369
	v417 = v374
	goto L114
L114:
	;
	v419 = v375 + int32(1)
	if v419 < v417 {
		v369 = v416
		v374 = v417
		v375 = v419
		goto L110
	} else {
		goto L119
	}
L115:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v395 = int32(4095)
	v398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v393+v394&v395))) = uint8(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v402 = v400 & v395
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v404 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+8)) = v402 | (v393-(v361+v403*v404))<<(uint(v404)%32)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v416 = v402 + v393 + int32(1)
	v417 = v415
	goto L114
L116:
	;
	v392 = F__emscripten_memcpy_bulkmem(m, v369, v255+v382*v383+int32(base.Ui32(v386)>>(uint(v383)%32)), v391)
	mBase = m.M
	v393 = v392
	goto L118
L117:
	;
	v393 = v369
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L111
L120:
	;
	F_errmsg_internal(m, int32(171120), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(492089), int32(793), int32(424764))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsquery_and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v19 != v7 {
					v78 = v7
					v79 = v14
					F_pfree(m, v78)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v83 = v79
						return v83
					}
				} else {
					v83 = v14
					return v83
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v21 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v14 != v24 {
						v78 = v14
						v79 = v7
						F_pfree(m, v78)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v83 = v79
							return v83
						}
					} else {
						v83 = v7
						return v83
					}
				} else {
					v27 = F_palloc0(m, int32(24))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v29 | int32(1)
						v34 = F_palloc0(m, int32(12))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27))) = v34
							v37 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v37)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)) = uint8(v37)
							v43 = F_palloc0(m, int32(8))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v43
								v47 = v14 + int32(8)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
								v52 = F_QT2QTN(m, v47, v47+v48*int32(12))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v54))) = v52
									v57 = v7 + int32(8)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
									v62 = F_QT2QTN(m, v57, v57+v58*int32(12))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v62
										*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
										v68 = F_QTN2QT(m, v27)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											F_QTNFree(m, v27)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v72 != v7 {
													F_pfree(m, v7)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v14 == v76 {
															v83 = v68
															return v83
														} else {
															v78 = v14
															v79 = v68
															F_pfree(m, v78)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																v83 = v79
																return v83
															}
														}
													}
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v14 == v76 {
														v83 = v68
														return v83
													} else {
														v78 = v14
														v79 = v68
														F_pfree(m, v78)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v83 = v79
															return v83
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
func F_tsquery_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v16)
							}
						} else {
							return base.B2i32(int32(0) < v16)
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v16)
						}
					} else {
						return base.B2i32(int32(0) < v16)
					}
				}
			}
		}
	}
}
func F_tsquery_rewrite_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v98 int32
	_ = v98
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_copy(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = l0 + int32(28)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_pg_detoast_datum_packed(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v29 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L84
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L81
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L78
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return v265
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v27 == v32 {
		v265 = v20
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v39 = v20 + int32(8)
	v43 = F_QT2QTN(m, v39, v39+v29*int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_pfree(m, v27)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v265 = v20
	goto L7
L13:
	;
	F_QTNTernary(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_QTNSort(m, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = F_text_to_cstring(m, v27)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(0)
	v56 = F_SPI_prepare(m, v49, v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v60 = F_SPI_cursor_open(m, v56)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v60 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_SPI_cursor_fetch(m, v60, int32(100))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v68 == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(2) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v76 = F_SPI_gettypeid(m, v71, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v76 != int32(3615) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v84 = F_SPI_gettypeid(m, v82, int32(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v84 != int32(3615) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v89 = base.B2i32(v43 != int32(0))
	v91 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	if v91 == int64(0) {
		v219 = v89
		v220 = v43
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	F_SPI_freetuptable(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L60
	}
L30:
	;
	if v43 == int32(0) {
		v219 = v89
		v220 = v43
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v98 = v43
	v109 = int64(0)
	goto L32
L32:
	;
	v112 = base.I32_wrap_i64(v109) << (uint(int32(2)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112+v115)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v122 = F_SPI_getbinval(m, v117, v118, int32(1), v17+int32(30))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v219 = v212
	v220 = v199
	goto L29
L34:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)))
	if v124 != 0 {
		v189 = v98
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	F_SPI_freetuptable(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L56
	}
L36:
	;
	v194 = v109 + int64(1)
	v196 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	if base.Ui64(v194) < base.Ui64(v196) {
		v98 = v189
		v109 = v194
		goto L32
	} else {
		goto L55
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127+v112)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v134 = F_SPI_getbinval(m, v129, v130, int32(2), v17+int32(30))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)))
	if v136 != 0 {
		v189 = v98
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v137 == int32(0) {
		v189 = v98
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v141 = v122 + int32(8)
	v145 = F_QT2QTN(m, v141, v141+v137*int32(12))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_QTNTernary(m, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_QTNSort(m, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v152 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v154 = v134 + int32(8)
	v158 = F_QT2QTN(m, v154, v154+v152*int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v161 = int32(0)
	goto L46
L46:
	;
	v162 = int32(4515120)
	v163 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v37
	v166 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+31)) = uint8(v166)
	v170 = F_dofindsubquery(m, v98, v145, v161, v17+int32(31))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v161 = v158
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v163
	F_QTNFree(m, v145)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_QTNFree(m, v161)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v178 = int32(0)
	if v170 == v178 {
		v199 = v178
		goto L35
	} else {
		goto L51
	}
L51:
	;
	F_QTNClearFlags(m, v170, int32(2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_QTNTernary(m, v170)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_QTNSort(m, v170)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v189 = v170
	goto L36
L55:
	;
	v199 = v189
	goto L35
L56:
	;
	F_SPI_cursor_fetch(m, v60, int32(100))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v212 = base.B2i32(v199 != int32(0))
	v214 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	if v214 == int64(0) {
		v219 = v212
		v220 = v199
		goto L29
	} else {
		goto L58
	}
L58:
	;
	if v199 != 0 {
		v98 = v199
		v109 = int64(0)
		goto L32
	} else {
		goto L59
	}
L59:
	;
	goto L33
L60:
	;
	F_SPI_cursor_close(m, v60)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_SPI_freeplan(m, v56)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v240 = F_SPI_finish(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v219 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_pfree(m, v49)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L73
	}
L65:
	;
	F_QTNBinary(m, v220)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(32)
	v254 = v20
	goto L64
L68:
	;
	v244 = F_QTN2QT(m, v220)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_QTNFree(m, v220)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20 == v248 {
		v254 = v244
		goto L64
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v20)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v254 = v244
	goto L64
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v257 != v27 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v27)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v265 = v254
	goto L7
L77:
	;
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v49
	F_errmsg_internal(m, int32(454392), v17)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(498183), int32(308), int32(15925))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v49
	F_errmsg_internal(m, int32(454363), v17+int32(16))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(498183), int32(311), int32(15925))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(147335), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(498183), int32(321), int32(15925))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
