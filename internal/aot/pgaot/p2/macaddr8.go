package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = v23
	goto L2
L1:
	;
	m.G0 = v20 + int32(16)
	return v389
L2:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.Ui32(v41-int32(9)) < base.Ui32(int32(5)) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v370 = int32(0)
	v371 = F_errsave_start(m, v22)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L79
	} else {
		goto L81
	}
L4:
	;
	goto L3
L5:
	;
	v24 = v24 + int32(1)
	goto L2
L6:
	;
	if v41 == int32(32) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v41 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v50 = v24
	v51 = v41
	v52 = v2
	v55 = v2
	v56 = v2
	v57 = v2
	v59 = v2
	v60 = v2
	v61 = v2
	v62 = v2
	v63 = v2
	v64 = v2
	goto L9
L9:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	switch v314 - int32(6) {
	case 0:
		v333 = int32(254)
		v334 = int32(255)
		v335 = v317
		v336 = v318
		v337 = v319
		goto L77
	default:
		goto L4
	case 2:
		goto L78
	}
L11:
	;
	goto L10
L12:
	;
	v314 = v52
	v317 = v55
	v318 = v56
	v319 = v57
	v322 = v60
	v323 = v61
	v324 = v62
	v325 = v63
	v326 = v64
	goto L11
L13:
	;
	goto L14
L14:
	;
	switch v52 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	case 5:
		goto L18
	case 6:
		goto L17
	case 7:
		goto L16
	default:
		goto L4
	}
L15:
	;
	v241 = v50 + int32(2)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v244 = v242 - int32(45)
	if base.Ui32(int32(13)) < base.Ui32(v244) {
		goto L57
	} else {
		goto L58
	}
L16:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L52
	}
L17:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L48
	}
L18:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L44
	}
L19:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L40
	}
L20:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L36
	}
L21:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L32
	}
L22:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L28
	}
L23:
	;
	if base.I32_extend8_s(v51) < int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v77 < int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v84 < int32(0) {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v232 = v55
	v233 = v56
	v234 = v57
	v235 = v60
	v236 = v61
	v237 = v62
	v238 = v63
	v239 = v84 + v77<<(uint(int32(4))%32)
	goto L15
L28:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v97 < int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v104 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v232 = v55
	v233 = v56
	v234 = v57
	v235 = v60
	v236 = v61
	v237 = v62
	v238 = v104 + v97<<(uint(int32(4))%32)
	v239 = v64
	goto L15
L32:
	;
	v117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v117 < int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v124 < int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v232 = v55
	v233 = v56
	v234 = v57
	v235 = v60
	v236 = v61
	v237 = v124 + v117<<(uint(int32(4))%32)
	v238 = v63
	v239 = v64
	goto L15
L36:
	;
	v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v137 < int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v144 < int32(0) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v232 = v55
	v233 = v56
	v234 = v144 + v137<<(uint(int32(4))%32)
	v235 = v60
	v236 = v61
	v237 = v62
	v238 = v63
	v239 = v64
	goto L15
L40:
	;
	v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v157 < int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v164 < int32(0) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v232 = v55
	v233 = v164 + v157<<(uint(int32(4))%32)
	v234 = v57
	v235 = v60
	v236 = v61
	v237 = v62
	v238 = v63
	v239 = v64
	goto L15
L44:
	;
	v177 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v177 < int32(0) {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v184 < int32(0) {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v232 = v184 + v177<<(uint(int32(4))%32)
	v233 = v56
	v234 = v57
	v235 = v60
	v236 = v61
	v237 = v62
	v238 = v63
	v239 = v64
	goto L15
L48:
	;
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v197 < int32(0) {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v204 < int32(0) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v232 = v55
	v233 = v56
	v234 = v57
	v235 = v60
	v236 = v204 + v197<<(uint(int32(4))%32)
	v237 = v62
	v238 = v63
	v239 = v64
	goto L15
L52:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51&int32(255))+uint32(_consts[1097]))))
	if v217 < int32(0) {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v67 < int32(0) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1097]))))
	if v224 < int32(0) {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v232 = v55
	v233 = v56
	v234 = v57
	v235 = v224 + v217<<(uint(int32(4))%32)
	v236 = v61
	v237 = v62
	v238 = v63
	v239 = v64
	goto L15
L56:
	;
	v265 = v52 + int32(1)
	switch v52 - int32(5) {
	case 0, 2:
		goto L68
	default:
		goto L67
	}
L57:
	;
	v262 = v59
	v263 = v241
	goto L56
L58:
	;
	goto L59
L59:
	;
	if int32(1)<<(uint(v244)%32)&int32(8195) == int32(0) {
		v262 = v59
		v263 = v241
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v254 = v59 & int32(255)
	if v254 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v262 = v258
	v263 = v50 + int32(3)
	goto L56
L62:
	;
	v258 = v242
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v254 != v242 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v258 = v59
	goto L61
L66:
	;
	if v308&int32(255) != 0 {
		v50 = v263
		v51 = v308
		v52 = v265
		v55 = v232
		v56 = v233
		v57 = v234
		v59 = v262
		v60 = v235
		v61 = v236
		v62 = v237
		v63 = v238
		v64 = v239
		goto L9
	} else {
		goto L76
	}
L67:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v308 = v307
	goto L66
L68:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v270 = v268 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v270) {
		v308 = v268
		goto L66
	} else {
		goto L69
	}
L69:
	;
	if int32(1)<<(uint(v270)%32)&int32(8388639) == int32(0) {
		v308 = v268
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v279 = v263
	goto L71
L71:
	;
	v297 = v279 + int32(1)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if base.Ui32(v298-int32(9)) < base.Ui32(int32(5)) {
		v279 = v297
		goto L71
	} else {
		goto L73
	}
L72:
	;
	if v298 == int32(0) {
		v314 = v265
		v317 = v232
		v318 = v233
		v319 = v234
		v322 = v235
		v323 = v236
		v324 = v237
		v325 = v238
		v326 = v239
		goto L11
	} else {
		goto L75
	}
L73:
	;
	if v298 == int32(32) {
		v279 = v297
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L4
L76:
	;
	v314 = v265
	v317 = v232
	v318 = v233
	v319 = v234
	v322 = v235
	v323 = v236
	v324 = v237
	v325 = v238
	v326 = v239
	goto L11
L77:
	;
	v339 = F_palloc0(m, int32(8))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v333 = v318
	v334 = v319
	v335 = v322
	v336 = v323
	v337 = v317
	goto L77
L79:
	;
	return int32(0)
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+7)) = uint8(v335)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+6)) = uint8(v336)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+5)) = uint8(v337)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+4)) = uint8(v333)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+3)) = uint8(v334)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+2)) = uint8(v324)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)) = uint8(v325)
	*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v326)
	v389 = v339
	goto L1
L81:
	;
	if v371 == int32(0) {
		v389 = v370
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(543633)
	F_errmsg(m, int32(704449), v20)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	F_errsave_finish(m, v22, int32(495165), int32(227), int32(277273))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	v389 = v370
	goto L1
}
