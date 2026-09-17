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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v160 int64
	_ = v160
	var v179 int64
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v387 int32
	_ = v387
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
	v28 = v26
	goto L3
L2:
	;
	v28 = int32(0)
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v33 != int64(0) {
		v179 = v33
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L11
	} else {
		goto L104
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L11
	} else {
		goto L100
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L11
	} else {
		goto L96
	}
L7:
	;
	m.G0 = v22 + int32(96)
	return v387
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v179 + int64(1)
	v184 = F_CopyReadLine(m, l0, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L41
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v37 == int32(0) {
		v179 = int64(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(1)
	v43 = F_CopyReadLine(m, l0, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v47 != int32(2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v43 != 0 {
		v387 = v5
		goto L7
	} else {
		goto L40
	}
L14:
	;
	v50 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v52 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v50 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v50 != v57 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L4
L20:
	;
	v64 = int32(0)
	goto L21
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v79 <= v64 {
		goto L13
	} else {
		goto L23
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L36
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v86 = v64 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87)))
	v94 = v32 + v81<<(uint(int32(4))%32) + v89*int32(100) - int32(80)
	v96 = v64 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v86+v97)))
	if v99 == int32(0) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v103 = v94 + int32(4)
	if v103|v99 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v118 == int32(0) {
		v64 = v96
		goto L21
	} else {
		goto L35
	}
L26:
	;
	v109 = int32(-1)
	goto L28
L27:
	;
	v109 = int32(0)
	goto L28
L28:
	;
	if v103 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v110 = int32(1)
	goto L31
L30:
	;
	v110 = v109
	goto L31
L31:
	;
	v111 = int32(0)
	if base.B2i32(v103 == v111)|base.B2i32(v99 == v111) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = v110
	goto L34
L33:
	;
	v117 = F_strncmp(m, v103, v99, int32(64))
	mBase = m.M
	v118 = v117
	goto L34
L34:
	;
	goto L25
L35:
	;
	goto L22
L36:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v96
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_0), v22+int32(80))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(826), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v179 = v160
	goto L8
L41:
	;
	if v184 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v186 == int32(0) {
		v387 = v5
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v189 = int32(1)
	v190 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if base.B2i32(v28 < v190)&base.B2i32(int32(0) < v28) != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v197 == int32(0) {
		v387 = v189
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v200 <= int32(0) {
		v387 = v189
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v203 = int32(0)
	if v203 < v190 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v207 = v190
	goto L52
L51:
	;
	v207 = v203
	goto L52
L52:
	;
	v217 = v203
	goto L53
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v232 = v217 << (uint(int32(2)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232+v233)))
	v240 = v32 + v227<<(uint(int32(4))%32) + v235*int32(100) - int32(80)
	if v217 != v207 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v387 = int32(1)
	goto L7
L55:
	;
	goto L54
L56:
	;
	v368 = v217 + int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v368 < v369 {
		v217 = v368
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
	v243 = v235 - int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v232+v192)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v246 != 0 {
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
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L91
	}
L61:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v246))))
	if v248 != int32(1) {
		goto L56
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v240 + int32(4)
	if v245 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v243))) = uint8(v256)
	goto L67
L66:
	;
	goto L67
L67:
	;
	v259 = v243 << (uint(int32(2)) % 32)
	v260 = l2 + v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v243))))
	if v263 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v259+v29)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	v270 = m.T0[v269].(func(*base.Module, int32, int32, int32) int32)(m, v267, l1, l3+v243)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v259+v30)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v280 = F_InputFunctionCallSafe(m, v31+v243*int32(28), v245, v277, v278, v279, v260)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v270
	goto L57
L72:
	;
	if v280 != 0 {
		goto L57
	} else {
		goto L73
	}
L73:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v282 + int64(1)
	v286 = int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v287 != v286 {
		v387 = v286
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v290)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v292 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v341)
	goto L55
L76:
	;
	v293 = F_CopyLimitPrintoutLength(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v320 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L87
	}
L79:
	;
	v297 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v297 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v300
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v299
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_3), v22+int32(32))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_pfree(m, v293)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1059), int32(_a_F_CopyFromTextOneRow_4))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
	if v320 == int32(0) {
		goto L75
	} else {
		goto L88
	}
L88:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v324
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_5), v22+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1066), int32(_a_F_CopyFromTextOneRow_4))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
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
	v349 = m.ExcPending
	if v349 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v240 + int32(4)
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_6), v22)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(977), int32(_a_F_CopyFromTextOneRow_4))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
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
	v408 = m.ExcPending
	if v408 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_7), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(962), int32(_a_F_CopyFromTextOneRow_4))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
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
	v424 = m.ExcPending
	if v424 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v94 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v96
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_8), v22-int32(-64))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(819), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
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
	v447 = m.ExcPending
	if v447 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v448 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v451 = v449
	goto L108
L107:
	;
	v451 = int32(0)
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v50
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_9), v22+int32(48))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(803), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
	var v57 int64
	_ = v57
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v11 = F_MakeTupleTableSlot(m, v9, int32(_a_F_RunFromStore_0))
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
		v67 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	m.T0[v68].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	v31 = v8
	goto L6
L6:
	;
	v32 = int32(_a_F_RunFromStore_1)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0])) = v35
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
	v67 = l2
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0])) = v33
	if v39 == int32(0) {
		v67 = v31
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
		v67 = v31
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
	v57 = v31 + int64(1)
	if base.B2i32(l2 == int64(0))|base.B2i32(v57 != l2) != 0 {
		v31 = v57
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	F_ExecDropSingleTupleTableSlot(m, v11)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return v67
}
