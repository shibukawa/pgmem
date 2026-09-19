package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parse_tsquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int64
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
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
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v448 int32
	_ = v448
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v21 = l3 & int32(2)
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(1508)
	goto L3
L2:
	;
	v22 = int32(1509)
	goto L3
L3:
	;
	v24 = l3 & int32(1)
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(1507)
	goto L6
L5:
	;
	v25 = v22
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v25
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v30 = base.B2i32(v27 != int32(447))
	goto L9
L8:
	;
	v30 = int32(1)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(12884901888)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	v38 = int32(3)
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = int32(7)
	goto L12
L11:
	;
	v41 = v38
	goto L12
L12:
	;
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = v38
	goto L15
L14:
	;
	v42 = v41
	goto L15
L15:
	;
	v43 = F_init_tsvector_parser(m, l0, v42, l4)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v14)+60)) = int64(64)
	v51 = F_palloc(m, int32(64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v51
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v55)
	F_makepol_1(m, v14+int32(28), l1, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	F_close_tsvector_parser(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	v465 = m.ExcPending
	if v465 != 0 {
		goto L16
	} else {
		goto L117
	}
L22:
	;
	m.G0 = v14 + int32(80)
	return v448
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v72 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v66 != int32(447) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v70 != 0 {
		v448 = int32(0)
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v30 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v103 = base.I32_div_u_s(int32(1073741815)-v100, int32(12))
	if base.Ui32(v103) < base.Ui32(v98) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v94 = F_palloc(m, int32(8))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L36
	}
L31:
	;
	v79 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	if v79 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v83
	F_errmsg(m, int32(_a_F_parse_tsquery_0), v14)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(880), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = int64(32)
	v448 = v94
	goto L22
L37:
	;
	v105 = int32(0)
	v106 = F_errsave_start(m, l4)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v126 = v100 + v98*int32(12) + int32(8)
	v127 = F_palloc0(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L16
	} else {
		goto L45
	}
L40:
	;
	if v106 == int32(0) {
		v448 = v105
		goto L22
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_parse_tsquery_3), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, l4, int32(_a_F_parse_tsquery_1), int32(890), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v448 = v105
	goto L22
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v126 << (uint(int32(2)) % 32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v132 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v135 = v133
	goto L48
L47:
	;
	v135 = int32(0)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v135
	v138 = v127 + int32(8)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v139 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if int32(0) < v140 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v217 = v135
	goto L51
L51:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if v224 != 0 {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	v148 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v217 = v212
	goto L51
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v148<<(uint(int32(2))%32))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	switch v160 - int32(1) {
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
	v198 = v148 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v198 < v199 {
		v148 = v198
		goto L55
	} else {
		goto L65
	}
L58:
	;
	v191 = v138 + v148*int32(12)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v194
	goto L57
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L62
	}
L60:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, uint32(v138+v148*int32(12)))) = v171
	goto L57
L61:
	;
	v166 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v138+v148*int32(12)))) = uint8(v166)
	goto L57
L62:
	;
	v177 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v177
	F_errmsg_internal(m, int32(_a_F_parse_tsquery_4), v14+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(917), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
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
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	base.MemoryCopy(m, v138+v217*int32(12), v228, v224)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_pfree(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v234
	F_findoprnd_recurse(m, v138, v14+int32(76), v233, v14+int32(27))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v233 != v244 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	if v246 == int32(0) {
		v448 = v127
		goto L22
	} else {
		goto L72
	}
L72:
	;
	v249 = m.G0
	v251 = v249 - int32(32)
	m.G0 = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v253 == int32(0) {
		v434 = v127
		goto L73
	} else {
		goto L74
	}
L73:
	;
	m.G0 = v251 + int32(32)
	v448 = v434
	goto L22
L74:
	;
	v257 = v127 + int32(8)
	v258 = F_maketree(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	v264 = F_clean_stopword_intree(m, v258, v251+int32(16), v251+int32(12))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	if v264 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v30 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v290 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v294 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L80:
	;
	v286 = F_palloc(m, int32(8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L86
	}
L81:
	;
	v272 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
	;
	if v272 == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_parse_tsquery_5), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L16
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_6), int32(409), int32(_a_F_parse_tsquery_7))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	goto L80
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v286))) = int64(32)
	v434 = v286
	goto L73
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v251)+24)) = int64(16)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	v330 = int32(1)
	if base.Ui32((v329-v330)&int32(255)) <= base.Ui32(v330) {
		goto L96
	} else {
		goto L97
	}
L88:
	;
	goto L87
L89:
	;
	v297 = v264
	v298 = v293
	v300 = v290
	goto L92
L90:
	;
	v313 = v293
	v315 = v290
	goto L91
L91:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	v323 = v316&int32(4095) + int32(1)
	v324 = v315
	goto L88
L92:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v302 = F_calcstrlen(m, v301)
	mBase = m.M
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	if v303 == int32(1) {
		v323 = v302
		v324 = v300
		goto L88
	} else {
		goto L94
	}
L93:
	;
	v313 = v308
	v315 = v306
	goto L91
L94:
	;
	v306 = v302 + v300
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v309 != int32(1) {
		v297 = v307
		v298 = v308
		v300 = v306
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v337 = F_palloc(m, int32(192))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L16
	} else {
		goto L99
	}
L97:
	;
	v346 = v6
	v347 = v6
	goto L98
L98:
	;
	v349 = v347 * int32(12)
	v352 = v323 + v324 + v349 + int32(8)
	v353 = F_palloc(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L16
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+20)) = v337
	F_plainnode(m, v251+int32(20), v264)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v251)+20))
	v346 = v345
	v347 = v344
	goto L98
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v352 << (uint(int32(2)) % 32)
	v360 = v353 + int32(8)
	if v349 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	base.MemoryCopy(m, v360, v346, v349)
	goto L104
L103:
	;
	goto L104
L104:
	;
	if int32(0) < v347 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v372 = v360 + v349
	v373 = int32(0)
	v374 = v347
	goto L108
L106:
	;
	goto L107
L107:
	;
	v434 = v353
	goto L73
L108:
	;
	v379 = v360 + v373*int32(12)
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v380 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L107
L110:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	v385 = v383 & int32(4095)
	if v385 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v416 = v372
	v417 = v374
	goto L112
L112:
	;
	v420 = v373 + int32(1)
	if v420 < v417 {
		v372 = v416
		v373 = v420
		v374 = v417
		goto L108
	} else {
		goto L116
	}
L113:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v387 = int32(12)
	base.MemoryCopy(m, v372, v257+v386*v387+int32(base.Ui32(v383)>>(uint(v387)%32)), v385)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	v395 = int32(4095)
	v398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v372+v394&v395))) = uint8(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	v402 = v400 & v395
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v404 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v402 | (v372-(v360+v403*v404))<<(uint(v404)%32)
	v416 = v402 + v372 + int32(1)
	v417 = v403
	goto L112
L116:
	;
	goto L109
L117:
	;
	F_errmsg_internal(m, int32(_a_F_parse_tsquery_8), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L16
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(793), int32(_a_F_parse_tsquery_9))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsquery_and(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14030(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_tsquery_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v13)
							}
						} else {
							return base.B2i32(int32(0) < v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v13)
						}
					} else {
						return base.B2i32(int32(0) < v13)
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L83
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L80
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L77
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return v262
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v25 == v30 {
		v262 = v20
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0]))
	v37 = v20 + int32(8)
	v41 = F_QT2QTN(m, v37, v37+v27*int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_pfree(m, v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v262 = v20
	goto L7
L13:
	;
	F_QTNTernary(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_QTNSort(m, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v47 = F_text_to_cstring(m, v25)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(0)
	v54 = F_SPI_prepare(m, v47, v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v54 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v58 = F_SPI_cursor_open(m, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v58 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_SPI_cursor_fetch(m, v58, int32(100))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	if v66 == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != int32(2) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v74 = F_SPI_gettypeid(m, v69, int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v74 != int32(3615) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v82 = F_SPI_gettypeid(m, v80, int32(2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v82 != int32(3615) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v86 = int32(0)
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if base.B2i32(v41 == v86)|base.B2i32(v91 == int64(0)) != 0 {
		v215 = base.B2i32(v41 != v86)
		v216 = v41
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	F_SPI_freetuptable(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L59
	}
L30:
	;
	v97 = v41
	v108 = int64(0)
	goto L31
L31:
	;
	v111 = base.I32_wrap_i64(v108) << (uint(int32(2)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v114)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v120 = v17 + int32(30)
	v121 = F_SPI_getbinval(m, v116, v117, int32(1), v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v215 = v208
	v216 = v195
	goto L29
L33:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)))
	if v123 != 0 {
		v185 = v97
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	F_SPI_freetuptable(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L55
	}
L35:
	;
	v190 = v108 + int64(1)
	v192 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if base.Ui64(v190) < base.Ui64(v192) {
		v97 = v185
		v108 = v190
		goto L31
	} else {
		goto L54
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v111+v126)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v131 = F_SPI_getbinval(m, v128, v129, int32(2), v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)))
	if v133 != 0 {
		v185 = v97
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v134 == int32(0) {
		v185 = v97
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v138 = v121 + int32(8)
	v142 = F_QT2QTN(m, v138, v138+v134*int32(12))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_QTNTernary(m, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_QTNSort(m, v142)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v149 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v151 = v131 + int32(8)
	v155 = F_QT2QTN(m, v151, v151+v149*int32(12))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v157 = int32(0)
	goto L45
L45:
	;
	v158 = int32(_a_F_tsquery_rewrite_query_0)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0])) = v35
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+31)) = uint8(v162)
	v166 = F_dofindsubquery(m, v97, v142, v157, v17+int32(31))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v157 = v155
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0])) = v159
	F_QTNFree(m, v142)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_QTNFree(m, v157)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v174 = int32(0)
	if v166 == v174 {
		v195 = v174
		goto L34
	} else {
		goto L50
	}
L50:
	;
	F_QTNClearFlags(m, v166, int32(2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_QTNTernary(m, v166)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_QTNSort(m, v166)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v185 = v166
	goto L35
L54:
	;
	v195 = v185
	goto L34
L55:
	;
	F_SPI_cursor_fetch(m, v58, int32(100))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v208 = base.B2i32(v195 != int32(0))
	v210 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if v210 == int64(0) {
		v215 = v208
		v216 = v195
		goto L29
	} else {
		goto L57
	}
L57:
	;
	if v195 != 0 {
		v97 = v195
		v108 = int64(0)
		goto L31
	} else {
		goto L58
	}
L58:
	;
	goto L32
L59:
	;
	F_SPI_cursor_close(m, v58)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_SPI_freeplan(m, v54)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v236 = F_SPI_finish(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v215 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	F_pfree(m, v47)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L72
	}
L64:
	;
	F_QTNBinary(m, v216)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(32)
	v250 = v20
	goto L63
L67:
	;
	v240 = F_QTN2QT(m, v216)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_QTNFree(m, v216)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20 == v244 {
		v250 = v240
		goto L63
	} else {
		goto L70
	}
L70:
	;
	F_pfree(m, v20)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v250 = v240
	goto L63
L72:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v253 != v25 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_pfree(m, v25)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v262 = v250
	goto L7
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v47
	F_errmsg_internal(m, int32(_a_F_tsquery_rewrite_query_1), v17)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(308), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v47
	F_errmsg_internal(m, int32(_a_F_tsquery_rewrite_query_4), v17+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(311), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_tsquery_rewrite_query_5), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(321), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
