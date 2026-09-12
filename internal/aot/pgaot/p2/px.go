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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
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
	v63 = m.ExcPending
	if v63 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v45
	F_errmsg(m, int32(206200), v8)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v45 = int32(315634)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v26 = int32(4395632)
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if l0 != v30 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v45 = v42
	goto L7
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v33 == int32(0) {
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
	v45 = int32(414155)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if l0 != v38 {
		v26 = v26 + int32(16)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v45 = v33
	goto L7
L20:
	;
	F_errfinish(m, int32(492716), int32(107), int32(525240))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	F_errmsg(m, int32(228386), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(492716), int32(100), int32(525240))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v460 int32
	_ = v460
	v7 = int32(492388)
	v9 = int32(4394160)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1379]))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v460
L2:
	;
	v413 = int32(4515652)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	F_ResourceOwnerEnlarge(m, v414)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L143
	} else {
		goto L144
	}
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1380])))
	if v34 == int32(0) {
		v53 = v33
		v54 = v34
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v13 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	v30 = l0
	goto L3
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = F_pg_strcasecmp(m, v16, l0)
	mBase = m.M
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v30 = v20
	goto L3
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v21 != 0 {
		v13 = v13 + int32(8)
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	if v54-v53 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	goto L13
L15:
	;
	if v33 != v34 {
		v53 = v33
		v54 = v34
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = v7
	v39 = v30
	goto L17
L17:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v42
		v54 = v43
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v53 = v42
	v54 = v43
	goto L14
L19:
	;
	v46 = int32(1)
	if v42 == v43 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v412 = int32(4394288)
	goto L2
L22:
	;
	goto L23
L23:
	;
	v61 = int32(503587)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1381])))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v86-v85 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	goto L24
L26:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v70 = v61
	v71 = v30
	goto L28
L28:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L25
	} else {
		goto L30
	}
L29:
	;
	v85 = v74
	v86 = v75
	goto L25
L30:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v412 = int32(4394296)
	goto L2
L33:
	;
	goto L34
L34:
	;
	v93 = int32(503388)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1382])))
	if v98 == int32(0) {
		v117 = v97
		v118 = v98
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v118-v117 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	goto L35
L37:
	;
	if v97 != v98 {
		v117 = v97
		v118 = v98
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v102 = v93
	v103 = v30
	goto L39
L39:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v106
		v118 = v107
		goto L36
	} else {
		goto L41
	}
L40:
	;
	v117 = v106
	v118 = v107
	goto L36
L41:
	;
	v110 = int32(1)
	if v106 == v107 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v412 = int32(4394304)
	goto L2
L44:
	;
	goto L45
L45:
	;
	v125 = int32(503533)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1383])))
	if v130 == int32(0) {
		v149 = v129
		v150 = v130
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v150-v149 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	goto L46
L48:
	;
	if v129 != v130 {
		v149 = v129
		v150 = v130
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v134 = v125
	v135 = v30
	goto L50
L50:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v139 == int32(0) {
		v149 = v138
		v150 = v139
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v149 = v138
	v150 = v139
	goto L47
L52:
	;
	v142 = int32(1)
	if v138 == v139 {
		v134 = v134 + v142
		v135 = v135 + v142
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v412 = int32(4394312)
	goto L2
L55:
	;
	goto L56
L56:
	;
	v157 = int32(492346)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1384])))
	if v162 == int32(0) {
		v181 = v161
		v182 = v162
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v182-v181 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	goto L57
L59:
	;
	if v161 != v162 {
		v181 = v161
		v182 = v162
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v166 = v157
	v167 = v30
	goto L61
L61:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v171 == int32(0) {
		v181 = v170
		v182 = v171
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v181 = v170
	v182 = v171
	goto L58
L63:
	;
	v174 = int32(1)
	if v170 == v171 {
		v166 = v166 + v174
		v167 = v167 + v174
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v412 = int32(4394320)
	goto L2
L66:
	;
	goto L67
L67:
	;
	v189 = int32(503604)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1385])))
	if v194 == int32(0) {
		v213 = v193
		v214 = v194
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v214-v213 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	goto L68
L70:
	;
	if v193 != v194 {
		v213 = v193
		v214 = v194
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v198 = v189
	v199 = v30
	goto L72
L72:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	if v203 == int32(0) {
		v213 = v202
		v214 = v203
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v213 = v202
	v214 = v203
	goto L69
L74:
	;
	v206 = int32(1)
	if v202 == v203 {
		v198 = v198 + v206
		v199 = v199 + v206
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v412 = int32(4394328)
	goto L2
L77:
	;
	goto L78
L78:
	;
	v221 = int32(492405)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1386])))
	if v226 == int32(0) {
		v245 = v225
		v246 = v226
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v246-v245 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	goto L79
L81:
	;
	if v225 != v226 {
		v245 = v225
		v246 = v226
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v230 = v221
	v231 = v30
	goto L83
L83:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v235 == int32(0) {
		v245 = v234
		v246 = v235
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v245 = v234
	v246 = v235
	goto L80
L85:
	;
	v238 = int32(1)
	if v234 == v235 {
		v230 = v230 + v238
		v231 = v231 + v238
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v412 = int32(4394336)
	goto L2
L88:
	;
	goto L89
L89:
	;
	v253 = int32(503594)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1387])))
	if v258 == int32(0) {
		v277 = v257
		v278 = v258
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v278-v277 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	goto L90
L92:
	;
	if v257 != v258 {
		v277 = v257
		v278 = v258
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v262 = v253
	v263 = v30
	goto L94
L94:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if v267 == int32(0) {
		v277 = v266
		v278 = v267
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v277 = v266
	v278 = v267
	goto L91
L96:
	;
	v270 = int32(1)
	if v266 == v267 {
		v262 = v262 + v270
		v263 = v263 + v270
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v412 = int32(4394344)
	goto L2
L99:
	;
	goto L100
L100:
	;
	v285 = int32(492395)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1388])))
	if v290 == int32(0) {
		v309 = v289
		v310 = v290
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v310-v309 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	goto L101
L103:
	;
	if v289 != v290 {
		v309 = v289
		v310 = v290
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v294 = v285
	v295 = v30
	goto L105
L105:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+1)))
	if v299 == int32(0) {
		v309 = v298
		v310 = v299
		goto L102
	} else {
		goto L107
	}
L106:
	;
	v309 = v298
	v310 = v299
	goto L102
L107:
	;
	v302 = int32(1)
	if v298 == v299 {
		v294 = v294 + v302
		v295 = v295 + v302
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v412 = int32(4394352)
	goto L2
L110:
	;
	goto L111
L111:
	;
	v317 = int32(503541)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1389])))
	if v322 == int32(0) {
		v341 = v321
		v342 = v322
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v342-v341 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L113:
	;
	goto L112
L114:
	;
	if v321 != v322 {
		v341 = v321
		v342 = v322
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v326 = v317
	v327 = v30
	goto L116
L116:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+1)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	if v331 == int32(0) {
		v341 = v330
		v342 = v331
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v341 = v330
	v342 = v331
	goto L113
L118:
	;
	v334 = int32(1)
	if v330 == v331 {
		v326 = v326 + v334
		v327 = v327 + v334
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v412 = int32(4394360)
	goto L2
L121:
	;
	goto L122
L122:
	;
	v349 = int32(492354)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1390])))
	if v354 == int32(0) {
		v373 = v353
		v374 = v354
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if v374-v373 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	goto L123
L125:
	;
	if v353 != v354 {
		v373 = v353
		v374 = v354
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v358 = v349
	v359 = v30
	goto L127
L127:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v363 == int32(0) {
		v373 = v362
		v374 = v363
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v373 = v362
	v374 = v363
	goto L124
L129:
	;
	v366 = int32(1)
	if v362 == v363 {
		v358 = v358 + v366
		v359 = v359 + v366
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v412 = int32(4394368)
	goto L2
L132:
	;
	goto L133
L133:
	;
	v382 = int32(503354)
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1391])))
	if v387 == int32(0) {
		v406 = v386
		v407 = v387
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v407-v406 != 0 {
		v460 = int32(-3)
		goto L1
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	if v386 != v387 {
		v406 = v386
		v407 = v387
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v391 = v382
	v392 = v30
	goto L138
L138:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+1)))
	if v396 == int32(0) {
		v406 = v395
		v407 = v396
		goto L135
	} else {
		goto L140
	}
L139:
	;
	v406 = v395
	v407 = v396
	goto L135
L140:
	;
	v399 = int32(1)
	if v395 == v396 {
		v391 = v391 + v399
		v392 = v392 + v399
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v412 = int32(4394376)
	goto L2
L143:
	;
	return int32(0)
L144:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v422 = F_MemoryContextAllocZero(m, v420, int32(100))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v422)+92)) = v424
	v426 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	*(*int32)(unsafe.Add(mBase, uint32(v422)+96)) = v426
	F_ResourceOwnerRemember(m, v426, v422, int32(4394652))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v433 = F_palloc(m, int32(36))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L143
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+24)) = int32(6852)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+8)) = int32(6853)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = int32(6854)
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = int32(6855)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v422)+92))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+28)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v433)+20)) = int32(6856)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+16)) = int32(6857)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+12)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v433
	v460 = int32(0)
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
