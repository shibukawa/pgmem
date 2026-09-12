package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyFromTextOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v162 int64
	_ = v162
	var v182 int64
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int64
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v391 int32
	_ = v391
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	v29 = v27
	goto L3
L2:
	;
	v29 = int32(0)
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v34 != int64(0) {
		v182 = v34
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v446 = int32(0)
	F_errstart_cold(m, int32(21), v446)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L11
	} else {
		goto L104
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L11
	} else {
		goto L100
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L11
	} else {
		goto L96
	}
L7:
	;
	m.G0 = v23 + int32(96)
	return v391
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v182 + int64(1)
	v187 = F_CopyReadLine(m, l0, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L41
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v38 == int32(0) {
		v182 = int64(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(1)
	v44 = F_CopyReadLine(m, l0, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v48 != int32(2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v44 != 0 {
		v391 = v5
		goto L7
	} else {
		goto L40
	}
L14:
	;
	v51 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v51 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v51 != v58 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L4
L20:
	;
	v67 = int32(0)
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v83 <= v67 {
		goto L13
	} else {
		goto L23
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L36
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v90 = v67 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90+v91)))
	v96 = v33 - int32(80) + v85<<(uint(int32(4))%32) + v93*int32(100)
	v98 = v67 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90+v99)))
	if v101 == int32(0) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v105 = v96 + int32(4)
	if v105|v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v119 == int32(0) {
		v67 = v98
		goto L21
	} else {
		goto L35
	}
L26:
	;
	v111 = int32(-1)
	goto L28
L27:
	;
	v111 = int32(0)
	goto L28
L28:
	;
	if v105 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(1)
	goto L31
L30:
	;
	v112 = v111
	goto L31
L31:
	;
	if v105 == int32(0) {
		v119 = v112
		goto L32
	} else {
		goto L33
	}
L32:
	;
	goto L25
L33:
	;
	if v101 == int32(0) {
		v119 = v112
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v118 = F_strncmp(m, v105, v101, int32(64))
	mBase = m.M
	v119 = v118
	goto L32
L35:
	;
	goto L22
L36:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v98
	F_errmsg(m, int32(688442), v23+int32(80))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(481783), int32(826), int32(301288))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v182 = v162
	goto L8
L41:
	;
	if v187 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v189 == int32(0) {
		v391 = v5
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v192 = int32(1)
	v193 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if base.B2i32(int32(0) < v29)&base.B2i32(v29 < v193) != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v200 == int32(0) {
		v391 = v192
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v203 <= int32(0) {
		v391 = v192
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v208 = int32(0)
	if v208 < v193 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v212 = v193
	goto L52
L51:
	;
	v212 = v208
	goto L52
L52:
	;
	v219 = v208
	goto L53
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v238 = v219 << (uint(int32(2)) % 32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v238+v239)))
	v243 = v241 - int32(1)
	v246 = v33 + int32(20) + v233<<(uint(int32(4))%32) + v243*int32(100)
	if v219 != v212 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v391 = int32(1)
	goto L7
L55:
	;
	goto L54
L56:
	;
	v372 = v219 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v372 < v373 {
		v219 = v372
		goto L53
	} else {
		goto L95
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = int64(0)
	goto L56
L58:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238+v195)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v250 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L11
	} else {
		goto L91
	}
L61:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v250))))
	if v252 != int32(1) {
		goto L56
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v246 + int32(4)
	if v249 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v260 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v243))) = uint8(v260)
	goto L67
L66:
	;
	goto L67
L67:
	;
	v263 = v243 << (uint(int32(2)) % 32)
	v264 = l2 + v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v243))))
	if v267 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v263+v30)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v274 = m.T0[v273].(func(*base.Module, int32, int32, int32) int32)(m, v271, l1, l3+v243)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v263+v31)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v246)+76))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v284 = F_InputFunctionCallSafe(m, v32+v243*int32(28), v249, v281, v282, v283, v264)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L11
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v274
	goto L57
L72:
	;
	if v284 != 0 {
		goto L57
	} else {
		goto L73
	}
L73:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v286 + int64(1)
	v290 = int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v291 != v290 {
		v391 = v290
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v294 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v294)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v296 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v345 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v345)
	goto L55
L76:
	;
	v297 = F_CopyLimitPrintoutLength(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v324 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L87
	}
L79:
	;
	v301 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v301 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v303
	F_errmsg(m, int32(693170), v23+int32(32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L11
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_pfree(m, v297)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	F_errfinish(m, int32(481783), int32(1059), int32(30508))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L75
L87:
	;
	if v324 == int32(0) {
		goto L75
	} else {
		goto L88
	}
L88:
	;
	v328 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v328
	F_errmsg(m, int32(62955), v23+int32(16))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(481783), int32(1066), int32(30508))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	goto L75
L91:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v246 + int32(4)
	F_errmsg(m, int32(675808), v23)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(481783), int32(977), int32(30508))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	goto L55
L96:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(264244), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(481783), int32(962), int32(30508))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v96 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v98
	F_errmsg(m, int32(688359), v23-int32(-64))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(481783), int32(819), int32(301288))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v454 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v456 = v455
	goto L108
L107:
	;
	v456 = v446
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v51
	F_errmsg(m, int32(462814), v23+int32(48))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(481783), int32(803), int32(301288))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RunFromStore(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v11 = F_MakeSingleTupleTableSlot(m, v9, int32(1583036))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	m.T0[v17].(func(*base.Module, int32, int32, int32))(m, l3, int32(1), v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 == int32(0) {
		v66 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	m.T0[v67].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v31 = v8
	goto L6
L6:
	;
	v32 = int32(4455216)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v39 = F_tuplestore_gettupleslot(m, v37, base.B2i32(l1 == int32(1)), int32(0), v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v66 = l2
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
	if v39 == int32(0) {
		v66 = v31
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v46 = m.T0[v45].(func(*base.Module, int32, int32) int32)(m, v11, l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v46 == int32(0) {
		v66 = v31
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	m.T0[v51].(func(*base.Module, int32))(m, v11)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v55 = v31 + int64(1)
	if l2 == int64(0) {
		v31 = v55
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if l2 != v55 {
		v31 = v55
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	F_ExecDropSingleTupleTableSlot(m, v11)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	return v66
}
