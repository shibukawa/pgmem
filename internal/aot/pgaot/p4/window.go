package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_window_frame_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	if l0&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l0&int32(2) != 0 {
		v21 = int32(703143)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v25 = l0 & int32(16)
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_appendStringInfoString(m, v9, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if l0&int32(4) != 0 {
		v21 = int32(702351)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l0&int32(8) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v21 = int32(702365)
	goto L5
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	F_appendStringInfoString(m, v9, int32(702665))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l0&int32(32) != 0 {
		v50 = int32(702856)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	if v25 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_appendStringInfoString(m, v9, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L23
	}
L17:
	;
	if l0&int32(512) != 0 {
		v50 = int32(702062)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if l0&int32(10240) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_get_rule_expr(m, l1, l3, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if l0&int32(2048) != 0 {
		v50 = int32(702865)
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if l0&int32(8192) == int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v50 = int32(702813)
	goto L16
L23:
	;
	goto L15
L24:
	;
	if l0&int32(32768) != 0 {
		v93 = int32(702054)
		goto L36
	} else {
		goto L37
	}
L25:
	;
	F_appendStringInfoString(m, v9, int32(703186))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if l0&int32(256) != 0 {
		v79 = int32(702804)
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_appendStringInfoString(m, v9, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L34
	}
L28:
	;
	if l0&int32(1024) != 0 {
		v79 = int32(702062)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if l0&int32(20480) == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_get_rule_expr(m, l2, l3, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	if l0&int32(4096) != 0 {
		v79 = int32(702865)
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if l0&int32(16384) == int32(0) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v79 = int32(702813)
	goto L27
L34:
	;
	goto L24
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v98 = v96 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100+v98))) = uint8(v102)
	goto L3
L36:
	;
	F_appendStringInfoString(m, v9, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L40
	}
L37:
	;
	if l0&int32(65536) != 0 {
		v93 = int32(702516)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if l0&int32(131072) == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v93 = int32(702446)
	goto L36
L40:
	;
	goto L35
}
func F_transformWindowDefinitions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(112)
	return v579
L2:
	;
	v29 = v4
	v32 = v4
	goto L11
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v18 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v579 = v4
	goto L1
L6:
	;
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L57
	} else {
		goto L142
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L57
	} else {
		goto L139
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L57
	} else {
		goto L134
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L57
	} else {
		goto L129
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32<<(uint(int32(2))%32))))
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v43 == v39 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L57
	} else {
		goto L124
	}
L13:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v259 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L14:
	;
	v256 = int32(0)
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L57
	} else {
		goto L63
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L57
	} else {
		goto L58
	}
L17:
	;
	if v29 == int32(0) {
		goto L16
	} else {
		goto L40
	}
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v115 == int32(0) {
		goto L14
	} else {
		goto L39
	}
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v114 != 0 {
		goto L17
	} else {
		goto L38
	}
L20:
	;
	if v29 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v48 <= int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v56 = int32(0)
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v51+v56<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L19
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v99 = v56 + int32(1)
	if v48 != v99 {
		v56 = v99
		goto L23
	} else {
		goto L37
	}
L28:
	;
	if v94-v93 == int32(0) {
		goto L15
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v78 = v70
	v79 = v43
	goto L32
L32:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v93 = v82
	v94 = v83
	goto L29
L34:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L27
L37:
	;
	goto L24
L38:
	;
	goto L14
L39:
	;
	goto L16
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v120 <= int32(0) {
		goto L16
	} else {
		goto L41
	}
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v128 = int32(0)
	goto L42
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v123+v128<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v142 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L16
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 == int32(0) {
		v165 = v145
		v166 = v146
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	v171 = v128 + int32(1)
	if v120 != v171 {
		v128 = v171
		goto L42
	} else {
		goto L56
	}
L47:
	;
	if v166-v165 == int32(0) {
		v256 = v141
		goto L13
	} else {
		goto L55
	}
L48:
	;
	goto L47
L49:
	;
	if v145 != v146 {
		v165 = v145
		v166 = v146
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v150 = v142
	v151 = v114
	goto L51
L51:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v155 == int32(0) {
		v165 = v154
		v166 = v155
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v165 = v154
	v166 = v155
	goto L48
L53:
	;
	v158 = int32(1)
	if v154 == v155 {
		v150 = v150 + v158
		v151 = v151 + v158
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L46
L56:
	;
	goto L43
L57:
	;
	return int32(0)
L58:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v195
	F_errmsg(m, int32(67259), v16+int32(80))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(475560), int32(2806), int32(129718))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L57
	} else {
		goto L64
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v217
	F_errmsg(m, int32(431991), v16+int32(96))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(475560), int32(2793), int32(129718))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L57
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v314 = F_transformGroupClause(m, l0, v310, int32(0), l2, v302, int32(9), int32(1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L57
	} else {
		goto L78
	}
L69:
	;
	v302 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v263 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v265 <= v263 {
		v302 = v263
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v271 = v263
	v273 = v263
	goto L73
L73:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281+v271<<(uint(int32(2))%32))))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v288 = F_findTargetlistEntrySQL99(m, l0, v286, l2, int32(10))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L57
	} else {
		goto L75
	}
L74:
	;
	v302 = v291
	goto L68
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v291 = F_addTargetToSortList(m, l0, v288, v273, v290, v285)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L57
	} else {
		goto L76
	}
L76:
	;
	v294 = v271 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v294 < v295 {
		v271 = v294
		v273 = v291
		goto L73
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v317 = F_palloc0(m, int32(56))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = int32(108)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+4)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+8)) = v323
	if v256 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L12
L81:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+20)) = v403
	if v403&int32(2) == int32(0) {
		v444 = v403
		goto L107
	} else {
		goto L108
	}
L82:
	;
	if v314 != 0 {
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v396 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+52)) = uint8(v396)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v317)+12)) = v314
	v400 = v302
	goto L81
L85:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v326 = F_copyObjectImpl(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L57
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+12)) = v326
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	if v302 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+52)) = uint8(v358)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	if v362 == int32(1058) {
		v400 = v359
		goto L81
	} else {
		goto L98
	}
L88:
	;
	v330 = int32(0)
	if v329 == v330 {
		v358 = v330
		v359 = v302
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v356 = F_copyObjectImpl(m, v329)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L57
	} else {
		goto L97
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L57
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L57
	} else {
		goto L93
	}
L93:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v340
	F_errmsg(m, int32(655311), v16+int32(48))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L57
	} else {
		goto L94
	}
L94:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L57
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(475560), int32(2867), int32(129718))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L57
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v358 = int32(1)
	v359 = v356
	goto L87
L98:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v302|v365 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v367 != int32(1058) {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L57
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L57
	} else {
		goto L102
	}
L102:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v377
	F_errmsg(m, int32(340953), v16+int32(32))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L57
	} else {
		goto L103
	}
L103:
	;
	F_errhint(m, int32(590894), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L57
	} else {
		goto L104
	}
L104:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L57
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(475560), int32(2904), int32(129718))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L57
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	if v444&int32(8) != 0 {
		goto L116
	} else {
		goto L117
	}
L108:
	;
	if v403&int32(30720) == int32(0) {
		v444 = v403
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if v400 == int32(0) {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v415 != int32(1) {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v421 = F_get_sortgroupclause_expr(m, v419, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L57
	} else {
		goto L112
	}
L112:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v430 = F_get_ordering_op_properties(m, v423, v16+int32(108), v16+int32(104), v16+int32(100))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L57
	} else {
		goto L113
	}
L113:
	;
	if v430 == int32(0) {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	v434 = F_exprCollation(m, v421)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L57
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+40)) = v434
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+16)))
	v439 = v437 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+44)) = uint8(v439)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+45)) = uint8(v441)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v317)+20))
	v444 = v443
	goto L107
L116:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v317)+16))
	if v448 == int32(0) {
		goto L7
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v456 = F_transformFrameOffset(m, l0, v444, v451, v452, v317+int32(32), v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L57
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+24)) = v456
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v317)+20))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v465 = F_transformFrameOffset(m, l0, v459, v460, v461, v317+int32(36), v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L57
	} else {
		goto L121
	}
L121:
	;
	v468 = v32 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+48)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v317)+28)) = v465
	v471 = F_lappend(m, v29, v317)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L57
	} else {
		goto L122
	}
L122:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v473 <= v468 {
		v579 = v471
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v29 = v471
	v32 = v468
	goto L11
L124:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L57
	} else {
		goto L125
	}
L125:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v482
	F_errmsg(m, int32(655358), v16-int32(-64))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L57
	} else {
		goto L126
	}
L126:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L57
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(475560), int32(2855), int32(129718))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L57
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L57
	} else {
		goto L130
	}
L130:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v504
	F_errmsg(m, int32(340953), v16+int32(16))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L57
	} else {
		goto L131
	}
L131:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L57
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(475560), int32(2897), int32(129718))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L57
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L57
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(261386), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L57
	} else {
		goto L136
	}
L136:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L57
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(475560), int32(2924), int32(129718))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L57
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
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v542
	F_errmsg_internal(m, int32(198838), v16)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L57
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(475560), int32(2933), int32(129718))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L57
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L57
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(341119), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L57
	} else {
		goto L144
	}
L144:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_parser_errposition(m, l0, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L57
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(475560), int32(2947), int32(129718))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L57
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_window_cume_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_WinGetPartitionRowCount(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+176))
	goto L3
L3:
	;
	v15 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_WinSetMarkPosition(m, v7, v13)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	if v17 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1)
	v29 = int32(0)
	goto L4
L7:
	;
	goto L8
L8:
	;
	v25 = F_WinRowsArePeers(m, v7, v13-int64(1), v13)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v29 = v25 ^ int32(1)
	goto L4
L10:
	;
	v33 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v29 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v75 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v68), base.F64_convert_i64_s(v8)))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	if v37 != int64(1) {
		v68 = v37
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)+176))
	goto L17
L16:
	;
	goto L15
L17:
	;
	v44 = v42 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v44
	if v8 <= v44 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v44
	goto L12
L19:
	;
	goto L20
L20:
	;
	v48 = v44
	goto L21
L21:
	;
	v55 = F_WinRowsArePeers(m, v7, v48-int64(1), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v68 = v61
	goto L12
L23:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	if v55 == int32(0) {
		v68 = v57
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v60 = int64(1)
	v61 = v57 + v60
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v61
	v64 = v48 + v60
	if v64 != v8 {
		v48 = v64
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	return v75
}
func F_window_lag_with_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_WinGetFuncArgCurrent(m, v9, int32(1), v7+int32(15))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v17 == int32(0) {
			v20 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v22 == v20 {
				v69 = v20
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
				if v28 == int32(0) {
					v69 = v20
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v33 = v31 - int32(11)
					if base.Ui32(int32(9)) < base.Ui32(v33) {
						v69 = v20
					} else {
						if int32(base.Ui32(int32(977))>>(uint(v33)%32))&int32(1) == int32(0) {
							v69 = v20
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32))+uint32(_consts[1352])))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v28+v48)))
							if v50 == int32(0) {
								v69 = v20
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								if v53 <= int32(1) {
									v69 = v20
								} else {
									v55 = int32(1)
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(4))))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									switch v61 - int32(7) {
									case 0:
										v69 = v55
									case 1:
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
										if v64 == int32(0) {
											v69 = v55
										} else {
											v69 = int32(0)
										}
									default:
										v69 = int32(0)
									}
								}
							}
						}
					}
				}
			}
			v75 = F_WinGetFuncArgInPartition(m, v9, v20-v13, v69, v7+int32(15), v7+int32(14))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v77 != int32(1) {
					v84 = v75
				} else {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
					v84 = int32(0)
				}
				m.G0 = v7 + int32(16)
				return v84
			}
		} else {
			v81 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
			v84 = int32(0)
			m.G0 = v7 + int32(16)
			return v84
		}
	}
}
