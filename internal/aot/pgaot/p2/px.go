package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_THROW_ERROR(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != int32(-17) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	if l0 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
	F_errmsg(m, int32(_a_F_px_THROW_ERROR_0), v8)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v41 = int32(_a_F_px_THROW_ERROR_1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v24 = int32(_a_F_px_THROW_ERROR_2)
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if l0 != v28 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v41 = v38
	goto L7
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v30 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v41 = int32(_a_F_px_THROW_ERROR_3)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if l0 != v34 {
		v24 = v24 + int32(16)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v41 = v30
	goto L7
L20:
	;
	F_errfinish(m, int32(_a_F_px_THROW_ERROR_4), int32(107), int32(_a_F_px_THROW_ERROR_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	F_errmsg(m, int32(_a_F_px_THROW_ERROR_6), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_px_THROW_ERROR_4), int32(100), int32(_a_F_px_THROW_ERROR_5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_px_find_cipher(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	v6 = int32(_a_F_px_find_cipher_0)
	v7 = int32(_a_F_px_find_cipher_1)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_cipher[0]))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v428
L2:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_cipher[1]))
	F_ResourceOwnerEnlarge(m, v388)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L109
	} else {
		goto L110
	}
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[2])))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v30 == int32(0))|base.B2i32(v30 != v33) != 0 {
		v51 = v30
		v52 = v33
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v10 = v7
	goto L7
L5:
	;
	goto L6
L6:
	;
	v27 = l0
	goto L3
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = F_pg_strcasecmp(m, v13, l0)
	mBase = m.M
	if v14 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v27 = v17
	goto L3
L10:
	;
	goto L11
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v18 != 0 {
		v10 = v10 + int32(8)
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	if v51-v52 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_2)
		goto L2
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v36 = v6
	v37 = v27
	goto L16
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v41
		v52 = v40
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v51 = v41
	v52 = v40
	goto L14
L18:
	;
	v44 = int32(1)
	if v41 == v40 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v57 = int32(_a_F_px_find_cipher_3)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[3])))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v81-v82 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_4)
		goto L2
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v66 = v57
	v67 = v27
	goto L24
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v81 = v71
	v82 = v70
	goto L22
L26:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v87 = int32(_a_F_px_find_cipher_5)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[4])))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v90 == int32(0))|base.B2i32(v90 != v93) != 0 {
		v111 = v90
		v112 = v93
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v111-v112 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_6)
		goto L2
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v96 = v87
	v97 = v27
	goto L32
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v101 == int32(0) {
		v111 = v101
		v112 = v100
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v111 = v101
	v112 = v100
	goto L30
L34:
	;
	v104 = int32(1)
	if v101 == v100 {
		v96 = v96 + v104
		v97 = v97 + v104
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v117 = int32(_a_F_px_find_cipher_7)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[5])))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v120 == int32(0))|base.B2i32(v120 != v123) != 0 {
		v141 = v120
		v142 = v123
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v141-v142 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_8)
		goto L2
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v126 = v117
	v127 = v27
	goto L40
L40:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v131 == int32(0) {
		v141 = v131
		v142 = v130
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v141 = v131
	v142 = v130
	goto L38
L42:
	;
	v134 = int32(1)
	if v131 == v130 {
		v126 = v126 + v134
		v127 = v127 + v134
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v147 = int32(_a_F_px_find_cipher_9)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[6])))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v150 == int32(0))|base.B2i32(v150 != v153) != 0 {
		v171 = v150
		v172 = v153
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v171-v172 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_10)
		goto L2
	} else {
		goto L52
	}
L46:
	;
	goto L45
L47:
	;
	v156 = v147
	v157 = v27
	goto L48
L48:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v161 == int32(0) {
		v171 = v161
		v172 = v160
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v171 = v161
	v172 = v160
	goto L46
L50:
	;
	v164 = int32(1)
	if v161 == v160 {
		v156 = v156 + v164
		v157 = v157 + v164
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v177 = int32(_a_F_px_find_cipher_11)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[7])))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v180 == int32(0))|base.B2i32(v180 != v183) != 0 {
		v201 = v180
		v202 = v183
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v201-v202 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_12)
		goto L2
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v186 = v177
	v187 = v27
	goto L56
L56:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if v191 == int32(0) {
		v201 = v191
		v202 = v190
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v201 = v191
	v202 = v190
	goto L54
L58:
	;
	v194 = int32(1)
	if v191 == v190 {
		v186 = v186 + v194
		v187 = v187 + v194
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v207 = int32(_a_F_px_find_cipher_13)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[8])))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v210 == int32(0))|base.B2i32(v210 != v213) != 0 {
		v231 = v210
		v232 = v213
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v231-v232 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_14)
		goto L2
	} else {
		goto L68
	}
L62:
	;
	goto L61
L63:
	;
	v216 = v207
	v217 = v27
	goto L64
L64:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v221 == int32(0) {
		v231 = v221
		v232 = v220
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v231 = v221
	v232 = v220
	goto L62
L66:
	;
	v224 = int32(1)
	if v221 == v220 {
		v216 = v216 + v224
		v217 = v217 + v224
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v237 = int32(_a_F_px_find_cipher_15)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[9])))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v240 == int32(0))|base.B2i32(v240 != v243) != 0 {
		v261 = v240
		v262 = v243
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v261-v262 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_16)
		goto L2
	} else {
		goto L76
	}
L70:
	;
	goto L69
L71:
	;
	v246 = v237
	v247 = v27
	goto L72
L72:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	if v251 == int32(0) {
		v261 = v251
		v262 = v250
		goto L70
	} else {
		goto L74
	}
L73:
	;
	v261 = v251
	v262 = v250
	goto L70
L74:
	;
	v254 = int32(1)
	if v251 == v250 {
		v246 = v246 + v254
		v247 = v247 + v254
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v267 = int32(_a_F_px_find_cipher_17)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[10])))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v270 == int32(0))|base.B2i32(v270 != v273) != 0 {
		v291 = v270
		v292 = v273
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v291-v292 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_18)
		goto L2
	} else {
		goto L84
	}
L78:
	;
	goto L77
L79:
	;
	v276 = v267
	v277 = v27
	goto L80
L80:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	if v281 == int32(0) {
		v291 = v281
		v292 = v280
		goto L78
	} else {
		goto L82
	}
L81:
	;
	v291 = v281
	v292 = v280
	goto L78
L82:
	;
	v284 = int32(1)
	if v281 == v280 {
		v276 = v276 + v284
		v277 = v277 + v284
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v297 = int32(_a_F_px_find_cipher_19)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[11])))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v300 == int32(0))|base.B2i32(v300 != v303) != 0 {
		v321 = v300
		v322 = v303
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v321-v322 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_20)
		goto L2
	} else {
		goto L92
	}
L86:
	;
	goto L85
L87:
	;
	v306 = v297
	v307 = v27
	goto L88
L88:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	if v311 == int32(0) {
		v321 = v311
		v322 = v310
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v321 = v311
	v322 = v310
	goto L86
L90:
	;
	v314 = int32(1)
	if v311 == v310 {
		v306 = v306 + v314
		v307 = v307 + v314
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v327 = int32(_a_F_px_find_cipher_21)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[12])))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v330 == int32(0))|base.B2i32(v330 != v333) != 0 {
		v351 = v330
		v352 = v333
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v351-v352 == int32(0) {
		v386 = int32(_a_F_px_find_cipher_22)
		goto L2
	} else {
		goto L100
	}
L94:
	;
	goto L93
L95:
	;
	v336 = v327
	v337 = v27
	goto L96
L96:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	if v341 == int32(0) {
		v351 = v341
		v352 = v340
		goto L94
	} else {
		goto L98
	}
L97:
	;
	v351 = v341
	v352 = v340
	goto L94
L98:
	;
	v344 = int32(1)
	if v341 == v340 {
		v336 = v336 + v344
		v337 = v337 + v344
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v357 = int32(_a_F_px_find_cipher_23)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_cipher[13])))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v360 == int32(0))|base.B2i32(v360 != v363) != 0 {
		v381 = v360
		v382 = v363
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v381-v382 != 0 {
		v428 = int32(-3)
		goto L1
	} else {
		goto L108
	}
L102:
	;
	goto L101
L103:
	;
	v366 = v357
	v367 = v27
	goto L104
L104:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
	if v371 == int32(0) {
		v381 = v371
		v382 = v370
		goto L102
	} else {
		goto L106
	}
L105:
	;
	v381 = v371
	v382 = v370
	goto L102
L106:
	;
	v374 = int32(1)
	if v371 == v370 {
		v366 = v366 + v374
		v367 = v367 + v374
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v386 = int32(_a_F_px_find_cipher_24)
	goto L2
L109:
	;
	return int32(0)
L110:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_cipher[14]))
	v396 = F_MemoryContextAllocZero(m, v394, int32(100))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+92)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_cipher[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+96)) = v401
	F_ResourceOwnerRemember(m, v401, v396, int32(_a_F_px_find_cipher_25))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v407 = F_palloc(m, int32(36))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+24)) = int32(_a_F_px_find_cipher_26)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = int32(_a_F_px_find_cipher_27)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(_a_F_px_find_cipher_28)
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = int32(_a_F_px_find_cipher_29)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v396)+92))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+28)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v407)+20)) = int32(_a_F_px_find_cipher_30)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+16)) = int32(_a_F_px_find_cipher_31)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v407
	v428 = int32(0)
	goto L1
}
func F_px_resolve_alias(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return l1
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v11 = v8
	v12 = l1
	goto L7
L5:
	;
	goto L3
L6:
	;
	if v49 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v15 == v16 {
		v38 = v15
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v49 = int32(0)
	goto L6
L9:
	;
	v40 = int32(1)
	if v38 != 0 {
		v11 = v11 + v40
		v12 = v12 + v40
		goto L7
	} else {
		goto L18
	}
L10:
	;
	if base.Ui32((v15-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = v15 | int32(32)
	goto L13
L12:
	;
	v26 = v15
	goto L13
L13:
	;
	if base.Ui32((v16-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = v16 | int32(32)
	goto L16
L15:
	;
	v35 = v16
	goto L16
L16:
	;
	if v26 == v35 {
		v38 = v26
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v49 = v26 - v35
	goto L6
L18:
	;
	goto L8
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	return v52
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v54 != 0 {
		v5 = v5 + int32(8)
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L5
}
