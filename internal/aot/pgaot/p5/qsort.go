package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_arg_entries(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v360 int64
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int64
	_ = v482
	var v483 int32
	_ = v483
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int64
	_ = v490
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v499 int32
	_ = v499
	var v500 int64
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v559 int64
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v653 int32
	_ = v653
	var v662 int64
	_ = v662
	var v668 int32
	_ = v668
	var v680 int64
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int64
	_ = v685
	var v686 int64
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int64
	_ = v695
	var v696 int64
	_ = v696
	var v719 int32
	_ = v719
	if base.Ui32(int32(7)) <= base.Ui32(l1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(v624) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L95
	}
L3:
	;
	v21 = l0
	v22 = l1
	goto L6
L4:
	;
	v600 = l0
	v601 = l1
	goto L5
L5:
	;
	v618 = v600
	v624 = v601
	goto L2
L6:
	;
	v40 = v21 + int32(8)
	v42 = v22
	goto L8
L7:
	;
	v600 = v21
	v601 = v296
	goto L5
L8:
	;
	if base.Ui32(v42) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v63 = v21 + v42<<(uint(int32(3))%32)
	v68 = v40
	goto L11
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v68-int32(8))))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	v88 = F_FunctionCall2Coll(m, v82, v83, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v104 = v21 + v42<<(uint(int32(2))%32)&int32(-8)
	if v42 != int32(7) {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	goto L12
L14:
	;
	v98 = v68 + int32(8)
	if base.Ui32(v98) < base.Ui32(v63) {
		v68 = v98
		goto L11
	} else {
		goto L21
	}
L15:
	;
	return
L16:
	;
	v90 = base.I32_wrap_i64(v88)
	if v90 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v93)
	goto L14
L18:
	;
	goto L19
L19:
	;
	if int32(0) < v90 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	goto L1
L22:
	;
	v108 = v63 - int32(8)
	if base.Ui32(v42) < base.Ui32(int32(41)) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v138 = v104
	goto L24
L24:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v141
	v146 = v63 - int32(8)
	v148 = v40
	v151 = v40
	v152 = v146
	v156 = v146
	goto L33
L25:
	;
	v134 = F_qsort_arg_entries_med3(m, v133, v131, v132, l2)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L15
	} else {
		goto L32
	}
L26:
	;
	v131 = v104
	v132 = v108
	v133 = v21
	goto L25
L27:
	;
	goto L28
L28:
	;
	v112 = v42 & int32(-8)
	v117 = v42 << (uint(int32(1)) % 32) & int32(-16)
	v119 = F_qsort_arg_entries_med3(m, v21, v21+v112, v21+v117, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v123 = F_qsort_arg_entries_med3(m, v104-v112, v104, v112+v104, l2)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v127 = F_qsort_arg_entries_med3(m, v108-v117, v108-v112, v108, l2)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v131 = v123
	v132 = v127
	v133 = v119
	goto L25
L32:
	;
	v138 = v134
	goto L24
L33:
	;
	if base.Ui32(v152) < base.Ui32(v151) {
		v207 = v148
		v210 = v151
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v292 = int32(3)
	v293 = (v207 - v21) >> (uint(v292) % 32)
	v296 = (v210 - v207) >> (uint(v292) % 32)
	if v293 < v296 {
		goto L59
	} else {
		goto L60
	}
L35:
	;
	if base.Ui32(v210) <= base.Ui32(v152) {
		goto L46
	} else {
		goto L47
	}
L36:
	;
	v167 = v148
	v170 = v151
	goto L37
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v188 = F_FunctionCall2Coll(m, v184, v185, v186, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L15
	} else {
		goto L40
	}
L38:
	;
	v207 = v201
	v210 = v204
	goto L35
L39:
	;
	v204 = v170 + int32(8)
	if base.Ui32(v204) <= base.Ui32(v152) {
		v167 = v201
		v170 = v204
		goto L37
	} else {
		goto L45
	}
L40:
	;
	v190 = base.I32_wrap_i64(v188)
	if v190 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v190 <= int32(0) {
		v201 = v167
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v193)
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v195
	v201 = v167 + int32(8)
	goto L39
L44:
	;
	v207 = v167
	v210 = v170
	goto L35
L45:
	;
	goto L38
L46:
	;
	v230 = v152
	v234 = v156
	goto L49
L47:
	;
	v278 = v152
	v282 = v156
	goto L48
L48:
	;
	goto L34
L49:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v247 = F_FunctionCall2Coll(m, v243, v244, v245, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L52
	}
L50:
	;
	v278 = v271
	v282 = v268
	goto L48
L51:
	;
	v271 = v230 - int32(8)
	if base.Ui32(v210) <= base.Ui32(v271) {
		v230 = v271
		v234 = v268
		goto L49
	} else {
		goto L57
	}
L52:
	;
	v249 = base.I32_wrap_i64(v247)
	if v249 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if int32(0) <= v249 {
		v268 = v234
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v260 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v260)
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v263
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = v262
	v268 = v234 - int32(8)
	goto L51
L56:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v210))) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v252
	v256 = int32(8)
	v148 = v207
	v151 = v210 + v256
	v152 = v230 - v256
	v156 = v234
	goto L33
L57:
	;
	goto L50
L58:
	;
	v438 = int32(3)
	v439 = (v282 - v278) >> (uint(v438) % 32)
	v444 = (v63-v282)>>(uint(v438)%32) - int32(1)
	if v439 < v444 {
		goto L74
	} else {
		goto L75
	}
L59:
	;
	v298 = v293
	goto L61
L60:
	;
	v298 = v296
	goto L61
L61:
	;
	if v298 == int32(0) {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v301 = int32(3)
	v303 = v210 - v298<<(uint(v301)%32)
	v305 = v298 & v301
	v306 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v298) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v319 = v306
	v328 = int32(0)
	goto L66
L64:
	;
	v376 = v306
	goto L65
L65:
	;
	v394 = v376
	v402 = v306
	goto L70
L66:
	;
	v332 = v319 << (uint(int32(3)) % 32)
	v333 = v21 + v332
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v333)))
	v335 = v332 + v303
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)))
	*(*int64)(unsafe.Add(mBase, uint32(v333))) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v335))) = v334
	v340 = v332 | int32(8)
	v341 = v21 + v340
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v341)))
	v343 = v340 + v303
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v342
	v348 = v332 | int32(16)
	v349 = v21 + v348
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
	v351 = v348 + v303
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v351)))
	*(*int64)(unsafe.Add(mBase, uint32(v349))) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v350
	v356 = v332 | int32(24)
	v357 = v21 + v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v357)))
	v359 = v356 + v303
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v359)))
	*(*int64)(unsafe.Add(mBase, uint32(v357))) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v359))) = v358
	v363 = int32(4)
	v364 = v319 + v363
	v366 = v328 + v363
	if v366 != v298&int32(-4) {
		v319 = v364
		v328 = v366
		goto L66
	} else {
		goto L68
	}
L67:
	;
	if v305 == int32(0) {
		goto L58
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v376 = v364
	goto L65
L70:
	;
	v407 = v394 << (uint(int32(3)) % 32)
	v408 = v21 + v407
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v408)))
	v410 = v407 + v303
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v410)))
	*(*int64)(unsafe.Add(mBase, uint32(v408))) = v411
	*(*int64)(unsafe.Add(mBase, uint32(v410))) = v409
	v414 = int32(1)
	v417 = v402 + v414
	if v417 != v305 {
		v394 = v394 + v414
		v402 = v417
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L58
L72:
	;
	goto L71
L73:
	;
	if base.Ui32(v296) <= base.Ui32(v439) {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v446 = v439
	goto L76
L75:
	;
	v446 = v444
	goto L76
L76:
	;
	if v446 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v449 = int32(3)
	v451 = v63 - v446<<(uint(v449)%32)
	v453 = v446 & v449
	v454 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v446) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v466 = v454
	v470 = int32(0)
	goto L81
L79:
	;
	v523 = v454
	goto L80
L80:
	;
	v541 = v523
	v551 = v454
	goto L85
L81:
	;
	v480 = v466 << (uint(int32(3)) % 32)
	v481 = v210 + v480
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v481)))
	v483 = v451 + v480
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v483)))
	*(*int64)(unsafe.Add(mBase, uint32(v481))) = v484
	*(*int64)(unsafe.Add(mBase, uint32(v483))) = v482
	v488 = v480 | int32(8)
	v489 = v210 + v488
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v489)))
	v491 = v488 + v451
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v491)))
	*(*int64)(unsafe.Add(mBase, uint32(v489))) = v492
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v490
	v496 = v480 | int32(16)
	v497 = v210 + v496
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v497)))
	v499 = v496 + v451
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v499)))
	*(*int64)(unsafe.Add(mBase, uint32(v497))) = v500
	*(*int64)(unsafe.Add(mBase, uint32(v499))) = v498
	v504 = v480 | int32(24)
	v505 = v210 + v504
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v505)))
	v507 = v504 + v451
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v507)))
	*(*int64)(unsafe.Add(mBase, uint32(v505))) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v507))) = v506
	v511 = int32(4)
	v512 = v466 + v511
	v514 = v470 + v511
	if v514 != v446&int32(-4) {
		v466 = v512
		v470 = v514
		goto L81
	} else {
		goto L83
	}
L82:
	;
	if v453 == int32(0) {
		goto L73
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v523 = v512
	goto L80
L85:
	;
	v555 = v541 << (uint(int32(3)) % 32)
	v556 = v210 + v555
	v557 = *(*int64)(unsafe.Add(mBase, uint32(v556)))
	v558 = v555 + v451
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v558)))
	*(*int64)(unsafe.Add(mBase, uint32(v556))) = v559
	*(*int64)(unsafe.Add(mBase, uint32(v558))) = v557
	v562 = int32(1)
	v565 = v551 + v562
	if v565 != v453 {
		v541 = v541 + v562
		v551 = v565
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L73
L87:
	;
	goto L86
L88:
	;
	F_qsort_arg_entries(m, v21, v296, l2)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L15
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_qsort_arg_entries(m, v63-v439<<(uint(int32(3))%32), v439, l2)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L15
	} else {
		goto L93
	}
L91:
	;
	v590 = v63 - v439<<(uint(int32(3))%32)
	if base.Ui32(v439) < base.Ui32(int32(7)) {
		v618 = v590
		v624 = v439
		goto L2
	} else {
		goto L92
	}
L92:
	;
	v21 = v590
	v22 = v439
	goto L6
L93:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v296) {
		v42 = v296
		goto L8
	} else {
		goto L94
	}
L94:
	;
	goto L9
L95:
	;
	v653 = v618 + int32(8)
	goto L96
L96:
	;
	if base.Ui32(v653) <= base.Ui32(v618) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L1
L98:
	;
	v719 = v653 + int32(8)
	if base.Ui32(v719) < base.Ui32(v618+v624<<(uint(int32(3))%32)) {
		v653 = v719
		goto L96
	} else {
		goto L108
	}
L99:
	;
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v653)))
	v668 = v653
	v680 = v662
	goto L100
L100:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v684 = v668 - int32(8)
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v684)))
	v686 = F_FunctionCall2Coll(m, v681, v682, v685, v680)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L15
	} else {
		goto L102
	}
L101:
	;
	goto L98
L102:
	;
	v688 = base.I32_wrap_i64(v686)
	if v688 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v691 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v691)
	goto L98
L104:
	;
	goto L105
L105:
	;
	if v688 <= int32(0) {
		goto L98
	} else {
		goto L106
	}
L106:
	;
	v695 = *(*int64)(unsafe.Add(mBase, uint32(v668)))
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v684)))
	*(*int64)(unsafe.Add(mBase, uint32(v668))) = v696
	*(*int64)(unsafe.Add(mBase, uint32(v684))) = v695
	if base.Ui32(v618) < base.Ui32(v684) {
		v668 = v684
		v680 = v695
		goto L100
	} else {
		goto L107
	}
L107:
	;
	goto L101
L108:
	;
	goto L97
}
func F_qsort_interruptible_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v8 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l4)
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v8 < int32(0) {
				if v12 < int32(0) {
					v31 = l1
					return v31
				} else {
					v18 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l4)
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 < int32(0) {
							v22 = l2
						} else {
							v22 = l0
						}
						return v22
					}
				}
			} else {
				if int32(0) < v12 {
					v31 = l1
					return v31
				} else {
					v26 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l4)
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 < int32(0) {
							v30 = l0
						} else {
							v30 = l2
						}
						v31 = v30
						return v31
					}
				}
			}
		}
	}
}
