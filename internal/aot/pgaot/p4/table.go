package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenTableList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
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
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_list_free(m, v504)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L14
	} else {
		goto L164
	}
L2:
	;
	v501 = v2
	v504 = v2
	v507 = v2
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = v2
	v30 = v2
	v31 = v2
	v33 = v2
	v35 = v2
	goto L7
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L14
	} else {
		goto L160
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L14
	} else {
		goto L156
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v35<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+16)))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L14
	} else {
		goto L152
	}
L9:
	;
	goto L8
L10:
	;
	v434 = v35 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v434 < v435 {
		v27 = v422
		v30 = v425
		v31 = v426
		v33 = v428
		v35 = v434
		goto L7
	} else {
		goto L151
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v52 = v43
	goto L13
L13:
	;
	v54 = F_table_openrv(m, v52, int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v52 = v51
	goto L13
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	v57 = int32(0)
	if v30 == v57 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v95 != 0 {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v95 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v63 <= int32(0) {
		v88 = v57
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v95 = v88
	goto L17
L22:
	;
	v66 = int32(0)
	if v66 < v63 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v69 = v63
	goto L25
L24:
	;
	v69 = v66
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v72 = int32(0)
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70+v72<<(uint(int32(2))%32))))
	v81 = base.B2i32(v80 == v56)
	if v80 == v56 {
		v88 = v81
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v88 = v81
	goto L21
L28:
	;
	v83 = v72 + int32(1)
	if v83 != v69 {
		v72 = v83
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v96 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v180 = F_palloc(m, int32(12))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L64
	}
L33:
	;
	v97 = int32(0)
	if v27 == v97 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v135 != 0 {
		goto L9
	} else {
		goto L47
	}
L35:
	;
	v135 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v103 <= int32(0) {
		v128 = v97
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = v128
	goto L34
L39:
	;
	v106 = int32(0)
	if v106 < v103 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v109 = v103
	goto L42
L41:
	;
	v109 = v106
	goto L42
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v112 = int32(0)
	goto L43
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110+v112<<(uint(int32(2))%32))))
	v121 = base.B2i32(v120 == v56)
	if v120 == v56 {
		v128 = v121
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v128 = v121
	goto L38
L45:
	;
	v123 = v112 + int32(1)
	if v123 != v109 {
		v112 = v123
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v136 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v137 = int32(0)
	if v31 == v137 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v175 != 0 {
		goto L6
	} else {
		goto L62
	}
L50:
	;
	v175 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v143 <= int32(0) {
		v168 = v137
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v175 = v168
	goto L49
L54:
	;
	v146 = int32(0)
	if v146 < v143 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v149 = v143
	goto L57
L56:
	;
	v149 = v146
	goto L57
L57:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v152 = int32(0)
	goto L58
L58:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150+v152<<(uint(int32(2))%32))))
	v161 = base.B2i32(v160 == v56)
	if v160 == v56 {
		v168 = v161
		goto L53
	} else {
		goto L60
	}
L59:
	;
	v168 = v161
	goto L53
L60:
	;
	v163 = v152 + int32(1)
	if v163 != v149 {
		v152 = v163
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_sequence_close(m, v54, int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	v422 = v27
	v425 = v30
	v426 = v31
	v428 = v33
	goto L10
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v54
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v185
	v187 = F_lappend(m, v33, v180)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v189 = F_lappend_oid(m, v30, v56)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v191 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v192 = F_lappend_oid(m, v27, v56)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L70
	}
L68:
	;
	v194 = v27
	goto L69
L69:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v195 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v194 = v192
	goto L69
L71:
	;
	v196 = F_lappend_oid(m, v31, v56)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L14
	} else {
		goto L74
	}
L72:
	;
	v198 = v31
	goto L73
L73:
	;
	if v44&int32(1) == int32(0) {
		v422 = v194
		v425 = v189
		v426 = v198
		v428 = v187
		goto L10
	} else {
		goto L75
	}
L74:
	;
	v198 = v196
	goto L73
L75:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+119)))
	if v204 == int32(112) {
		v422 = v194
		v425 = v189
		v426 = v198
		v428 = v187
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v209 = F_find_all_inheritors(m, v56, int32(4), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	if v209 == int32(0) {
		v422 = v194
		v425 = v189
		v426 = v198
		v428 = v187
		goto L10
	} else {
		goto L78
	}
L78:
	;
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v214 <= v213 {
		v422 = v194
		v425 = v189
		v426 = v198
		v428 = v187
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v220 = v194
	v222 = v54
	v223 = v189
	v224 = v198
	v226 = v187
	v227 = v213
	goto L80
L80:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v227<<(uint(int32(2))%32))))
	v237 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v237 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v422 = v409
	v425 = v411
	v426 = v412
	v428 = v413
	goto L10
L82:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L14
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v240 = int32(0)
	if v223 == v240 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L84
L86:
	;
	v416 = v227 + int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v416 < v417 {
		v220 = v409
		v222 = v410
		v223 = v411
		v224 = v412
		v226 = v413
		v227 = v416
		goto L80
	} else {
		goto L150
	}
L87:
	;
	if v278 != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v278 = int32(0)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v246 <= int32(0) {
		v271 = v240
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v278 = v271
	goto L87
L92:
	;
	v249 = int32(0)
	if v249 < v246 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v252 = v246
	goto L95
L94:
	;
	v252 = v249
	goto L95
L95:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v255 = int32(0)
	goto L96
L96:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v253+v255<<(uint(int32(2))%32))))
	v264 = base.B2i32(v263 == v235)
	if v263 == v235 {
		v271 = v264
		goto L91
	} else {
		goto L98
	}
L97:
	;
	v271 = v264
	goto L91
L98:
	;
	v266 = v255 + int32(1)
	if v266 != v252 {
		v255 = v266
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	if v235 == v56 {
		v409 = v220
		v410 = v222
		v411 = v223
		v412 = v224
		v413 = v226
		goto L86
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v386 = F_table_open(m, v235, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L140
	}
L103:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v280 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v281 = int32(0)
	if v220 == v281 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v319 != 0 {
		goto L5
	} else {
		goto L118
	}
L106:
	;
	v319 = int32(0)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v287 <= int32(0) {
		v312 = v281
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v319 = v312
	goto L105
L110:
	;
	v290 = int32(0)
	if v290 < v287 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v293 = v287
	goto L113
L112:
	;
	v293 = v290
	goto L113
L113:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v296 = int32(0)
	goto L114
L114:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v294+v296<<(uint(int32(2))%32))))
	v305 = base.B2i32(v304 == v235)
	if v304 == v235 {
		v312 = v305
		goto L109
	} else {
		goto L116
	}
L115:
	;
	v312 = v305
	goto L109
L116:
	;
	v307 = v296 + int32(1)
	if v307 != v293 {
		v296 = v307
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v320 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v323 = int32(0)
	if v224 == v323 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L121
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L14
	} else {
		goto L136
	}
L122:
	;
	if v361 == int32(0) {
		v409 = v220
		v410 = v222
		v411 = v223
		v412 = v224
		v413 = v226
		goto L86
	} else {
		goto L135
	}
L123:
	;
	v361 = int32(0)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v329 <= int32(0) {
		v354 = v323
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v361 = v354
	goto L122
L127:
	;
	v332 = int32(0)
	if v332 < v329 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v335 = v329
	goto L130
L129:
	;
	v335 = v332
	goto L130
L130:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v338 = int32(0)
	goto L131
L131:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v336+v338<<(uint(int32(2))%32))))
	v347 = base.B2i32(v346 == v235)
	if v346 == v235 {
		v354 = v347
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v354 = v347
	goto L126
L133:
	;
	v349 = v338 + int32(1)
	if v349 != v335 {
		v338 = v349
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	goto L121
L136:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v371 + int32(4)
	F_errmsg(m, int32(667460), v15+int32(-16))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L14
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(464032), int32(1775), int32(72095))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L14
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v389 = F_palloc(m, int32(12))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L14
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v386
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+8)) = v394
	v396 = F_lappend(m, v226, v389)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	v398 = F_lappend_oid(m, v223, v235)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v400 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v401 = F_lappend_oid(m, v220, v235)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L147
	}
L145:
	;
	v403 = v220
	goto L146
L146:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v404 == int32(0) {
		v409 = v403
		v410 = v386
		v411 = v398
		v412 = v224
		v413 = v396
		goto L86
	} else {
		goto L148
	}
L147:
	;
	v403 = v401
	goto L146
L148:
	;
	v407 = F_lappend_oid(m, v224, v235)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L14
	} else {
		goto L149
	}
L149:
	;
	v409 = v403
	v410 = v386
	v411 = v398
	v412 = v407
	v413 = v396
	goto L86
L150:
	;
	goto L81
L151:
	;
	v501 = v422
	v504 = v425
	v507 = v428
	goto L1
L152:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L14
	} else {
		goto L153
	}
L153:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v444 + int32(4)
	F_errmsg(m, int32(667513), v17)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L14
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(464032), int32(1700), int32(72095))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L14
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L14
	} else {
		goto L157
	}
L157:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v463 + int32(4)
	F_errmsg(m, int32(667460), v15+int32(-48))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L14
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(464032), int32(1707), int32(72095))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L14
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L14
	} else {
		goto L161
	}
L161:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v484 + int32(4)
	F_errmsg(m, int32(667513), v15+int32(-32))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L14
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(464032), int32(1763), int32(72095))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L14
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_list_free(m, v501)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L14
	} else {
		goto L165
	}
L165:
	;
	m.G0 = v17 - int32(-64)
	return v507
}
func F_fetch_table_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v14
	v17 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v22 = m.T0[v21].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = F_makeStringInfo(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_GetPublicationsStr(m, l1, v26, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(160000) <= v22 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v56
	F_appendStringInfo(m, v11-int32(-64), v55, v11+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(22)
	v55 = int32(699294)
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(1003)
	F_appendStringInfoString(m, v11-int32(-64), int32(705970))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = int32(634798)
	if v22 < int32(150000) {
		v55 = v47
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(696730))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v55 = v47
	goto L6
L13:
	;
	F_free_attrmap(m, v26)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(149999) < v22 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = int32(3)
	goto L17
L16:
	;
	v72 = int32(2)
	goto L17
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+60))
	v78 = m.T0[v77].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v67, v72, v11+int32(48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v83 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L69
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v88 = F_MakeSingleTupleTableSlot(m, v86, int32(1567052))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L65
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v93 = F_tuplestore_gettupleslot(m, v90, int32(1), int32(0), v88)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = v66
	goto L29
L27:
	;
	v144 = v66
	goto L28
L28:
	;
	F_ExecDropSingleTupleTableSlot(m, v88)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L51
	}
L29:
	;
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
	if v105 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v144 = v132
	goto L28
L31:
	;
	F_slot_getsomeattrs_int(m, v88, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = F_text_to_cstring(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
	if v115 <= int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_slot_getsomeattrs_int(m, v88, int32(2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v123 = F_text_to_cstring(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v126 = F_makeRangeVar(m, v113, v123, int32(-1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if base.B2i32(v22 < int32(150000)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = F_list_member(m, v98, v126)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v132 = F_lappend(m, v98, v126)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	if v130 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	m.T0[v135].(func(*base.Module, int32))(m, v88)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v141 = F_tuplestore_gettupleslot(m, v138, int32(1), int32(0), v88)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v141 != 0 {
		v98 = v132
		goto L29
	} else {
		goto L50
	}
L50:
	;
	goto L30
L51:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v153 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v156 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_tuplestore_end(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v159 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	F_FreeTupleDesc(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v78)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	m.G0 = v11 + int32(80)
	return v144
L65:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v175
	F_errmsg(m, int32(189586), v11+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(463982), int32(2277), int32(70002))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v113
	F_errmsg(m, int32(132578), v11)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(463982), int32(2299), int32(70002))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_table_am_handler_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(206541)
			F_errmsg(m, int32(180554), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(463820), int32(370), int32(262447))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v7
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		switch v11 {
		case 0, 5:
			v12 = l1 + v7
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v24)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v20
			if v23 != 0 {
				v31 = v19
			} else {
				v31 = int32(0)
			}
			if v24 != 0 {
				v32 = v31
			} else {
				v32 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			if v34 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v40 = F___memcpy(m, v12+int32(24), v37, v34<<(uint(int32(2))%32))
				mBase = m.M
			} else {
			}
			if int32(0) < v32 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v44 = int32(2)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v53 = F___memcpy(m, v12+v43<<(uint(v44)%32)+int32(24), v49, v50<<(uint(v44)%32))
				mBase = m.M
			} else {
			}
			v55 = int32(0)
		default:
			v55 = int32(1)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v55)
		return
	}
}
func F_table_slot_callbacks(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v5 = m.T0[v4].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13 == int32(102) {
			v16 = int32(1567000)
		} else {
			v16 = int32(1566948)
		}
		return v16
	}
}
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v6 = m.T0[v5].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = v6
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v19 = F_MakeSingleTupleTableSlot(m, v18, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if l1 != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v22 = F_lappend(m, v21, v19)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
						return v19
					}
				} else {
					return v19
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13 == int32(102) {
			v16 = int32(1567000)
		} else {
			v16 = int32(1566948)
		}
		v17 = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v19 = F_MakeSingleTupleTableSlot(m, v18, v17)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if l1 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v22 = F_lappend(m, v21, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
					return v19
				}
			} else {
				return v19
			}
		}
	}
}
func F_table_to_xml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_text_to_cstring(m, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_initStringInfo(m, v7+int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v24 = F_DirectFunctionCall1Coll(m, int32(1562), int32(0), v10)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
					F_appendStringInfo(m, v7+int32(16), int32(186078), v7)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						v33 = F_get_rel_name(m, v10)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(0)
							v38 = F_query_to_xml_internal(m, v32, v33, v35, base.B2i32(v9 != v35), v16)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
								v42 = F_cstring_to_text_with_len(m, v40, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return v42
								}
							}
						}
					}
				}
			}
		}
	}
}
