package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyReadLine(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
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
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int64
	_ = v629
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v693 int32
	_ = v693
	v3 = int32(0)
	v18 = l0 + int32(304)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v3
	goto L1
L1:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v26)
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v693 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v693)
	return v678
L3:
	;
	v649 = int32(0)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v650 - int32(1) {
	case 0:
		goto L176
	case 1:
		goto L175
	case 2:
		goto L174
	default:
		v678 = v649
		goto L2
	}
L4:
	;
	v603 = int32(1)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v604 != v603 {
		v678 = v603
		goto L2
	} else {
		goto L169
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v31 != v34 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v285 = v283
	v288 = v283
	v289 = v3
	v291 = v3
	v292 = v281
	goto L86
L8:
	;
	v36 = v31
	goto L10
L9:
	;
	v36 = int32(0)
	goto L10
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v41 = v39
	v43 = v39
	v45 = v28
	v46 = v3
	v47 = v3
	v51 = v3
	v52 = v3
	goto L11
L11:
	;
	if base.B2i32(v45 <= v43)|v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v41 < v43 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v75 = v41
	v76 = v45
	v77 = v46
	v78 = v43
	goto L15
L15:
	;
	v79 = int32(1)
	v80 = v78 + v79
	v82 = base.B2i32(v80 < v76) | v77
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v29))))
	if v82&v79|base.B2i32(v86 != int32(13)) != 0 {
		v134 = v75
		v135 = v78
		v136 = v80
		v137 = v86
		v138 = v76
		v139 = v77
		v141 = v82
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_appendBinaryStringInfo(m, v18, v59+v41, v43-v41)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_CopyLoadInputBuf(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v43
	goto L18
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v69-v70 <= int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	v75 = v70
	v76 = v69
	v77 = v74
	v78 = v70
	goto L15
L23:
	;
	v150 = v137 & int32(255)
	v152 = base.B2i32(v150 == v36&int32(255))
	v154 = v52 ^ v51&v152
	v156 = v51 ^ (base.B2i32(v34 != v150) | v154)
	if v156&int32(1) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	v91 = v75
	v92 = v78
	goto L25
L25:
	;
	if v91 < v92 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v134 = v116
	v135 = v116
	v136 = v122
	v137 = v126
	v138 = v115
	v139 = v120
	v141 = v124
	goto L23
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_appendBinaryStringInfo(m, v18, v107+v91, v92-v91)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_CopyLoadInputBuf(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v92
	goto L29
L31:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v115-v116 <= int32(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	v122 = v116 + int32(1)
	v124 = v120 | base.B2i32(v122 < v115)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v29))))
	if v126 != int32(13) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v134 = v116
	v135 = v116
	v136 = v122
	v137 = v126
	v138 = v115
	v139 = v120
	v141 = v124
	goto L23
L34:
	;
	goto L35
L35:
	;
	if v124&int32(1) == int32(0) {
		v91 = v116
		v92 = v116
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v41 = v134
	v43 = v275
	v45 = v138
	v46 = v139
	v47 = v277
	v51 = v156 ^ int32(1)
	v52 = v154 & v152
	goto L11
L38:
	;
	v161 = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v165 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	switch v137 - int32(10) {
	case 0:
		goto L48
	default:
		v275 = v136
		v277 = int32(0)
		goto L37
	case 3:
		goto L47
	}
L41:
	;
	v168 = int32(10)
	goto L43
L42:
	;
	v168 = int32(13)
	goto L43
L43:
	;
	if base.I32_extend8_s(v137) != v168 {
		v275 = v136
		v277 = v161
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v170 + int64(1)
	v275 = v136
	v277 = v161
	goto L37
L45:
	;
	if v265 <= v134 {
		goto L3
	} else {
		goto L78
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v259
	v265 = v260
	goto L45
L47:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v203 {
	case 0, 3:
		goto L58
	case 1:
		goto L57
	default:
		v265 = v136
		goto L45
	}
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v177&int32(-2) != int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v259 = int32(1)
	v260 = v136
	goto L46
L50:
	;
	goto L51
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_0), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_1), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1721), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L19
	} else {
		goto L73
	}
L58:
	;
	v204 = int32(1)
	if v141&v204 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v275 = v135
	v277 = v204
	goto L37
L60:
	;
	goto L61
L61:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v29))))
	if v210 == int32(10) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v259 = int32(3)
	v260 = v135 + int32(2)
	goto L46
L63:
	;
	goto L64
L64:
	;
	if v203 != int32(3) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v259 = int32(2)
	v260 = v136
	goto L46
L66:
	;
	goto L67
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_4), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_5), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L19
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1688), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_4), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_5), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1705), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L19
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_appendBinaryStringInfo(m, v18, v269+v134, v265-v134)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v265
	goto L3
L80:
	;
	if v319 < v575 {
		goto L164
	} else {
		goto L165
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v568
	v574 = int32(0)
	v575 = v569
	goto L80
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L19
	} else {
		goto L160
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L19
	} else {
		goto L156
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L19
	} else {
		goto L152
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L19
	} else {
		goto L148
	}
L86:
	;
	if v289&int32(1)|base.B2i32(v292 <= v285) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454+v282))))
	switch v463 - int32(10) {
	case 0, 3:
		goto L140
	default:
		goto L141
	}
L88:
	;
	if v288 < v285 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v319 = v288
	v320 = v291
	v321 = v292
	v322 = v285
	goto L90
L90:
	;
	v323 = int32(0)
	v325 = v322 + int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v282))))
	switch v327 - int32(10) {
	case 0:
		goto L98
	case 1, 2:
		v285 = v325
		v288 = v319
		v289 = v323
		v291 = v320
		v292 = v321
		goto L86
	case 3:
		goto L99
	default:
		goto L97
	}
L91:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_appendBinaryStringInfo(m, v18, v305+v288, v285-v288)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L19
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_CopyLoadInputBuf(m, l0)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L19
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v285
	goto L93
L95:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v313-v314 <= int32(0) {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	v319 = v314
	v320 = v318
	v321 = v313
	v322 = v314
	goto L90
L97:
	;
	if v327 != int32(92) {
		v285 = v325
		v288 = v319
		v289 = v323
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L127
	}
L98:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v389&int32(-2) != int32(2) {
		goto L119
	} else {
		goto L120
	}
L99:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v331 {
	case 0, 3:
		goto L101
	case 1:
		goto L100
	default:
		v574 = int32(0)
		v575 = v325
		goto L80
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L19
	} else {
		goto L114
	}
L101:
	;
	v332 = int32(1)
	if (base.B2i32(v325 < v321)|v320)&v332 == int32(0) {
		v285 = v322
		v288 = v319
		v289 = v332
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L102
	}
L102:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+v282))))
	if v340 == int32(10) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v568 = int32(3)
	v569 = v322 + int32(2)
	goto L81
L104:
	;
	goto L105
L105:
	;
	if v331 != int32(3) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v568 = int32(2)
	v569 = v325
	goto L81
L107:
	;
	goto L108
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L19
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_6), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_7), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1688), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_6), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_7), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1705), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L19
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v568 = int32(1)
	v569 = v325
	goto L81
L120:
	;
	goto L121
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L19
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_8), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L19
	} else {
		goto L124
	}
L124:
	;
	F_errhint(m, int32(_a_F_CopyReadLine_9), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L19
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1721), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L19
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	v417 = int32(1)
	if (base.B2i32(v325 < v321)|v320)&v417 == int32(0) {
		v285 = v322
		v288 = v319
		v289 = v417
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L128
	}
L128:
	;
	if v320&base.B2i32(v321 <= v325) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v574 = int32(1)
	v575 = v325
	goto L80
L130:
	;
	goto L131
L131:
	;
	v429 = v322 + int32(2)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+v282))))
	if v431 != int32(46) {
		v285 = v429
		v288 = v319
		v289 = int32(0)
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L132
	}
L132:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v434 == int32(3) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v437 = int32(1)
	if (base.B2i32(v429 < v321)|v320)&v437 == int32(0) {
		v285 = v322
		v288 = v319
		v289 = v437
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L136
	}
L134:
	;
	v454 = v429
	goto L135
L135:
	;
	v455 = int32(1)
	if (base.B2i32(v454 < v321)|v320)&v455 == int32(0) {
		v285 = v322
		v288 = v319
		v289 = v455
		v291 = v320
		v292 = v321
		goto L86
	} else {
		goto L139
	}
L136:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429+v282))))
	if v445 == int32(10) {
		goto L85
	} else {
		goto L137
	}
L137:
	;
	if v445 != int32(13) {
		goto L84
	} else {
		goto L138
	}
L138:
	;
	v454 = v322 + int32(3)
	goto L135
L139:
	;
	goto L87
L140:
	;
	if (base.B2i32(v434 == int32(1))|base.B2i32(v434 == int32(3)))&base.B2i32(v463 != int32(10))|base.B2i32(v434 == int32(2))&base.B2i32(v463 != int32(13)) != 0 {
		goto L83
	} else {
		goto L146
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
	} else {
		goto L142
	}
L142:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_10), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1774), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L19
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if base.B2i32(v319 < v322)|base.B2i32(int32(0) < v497) != 0 {
		goto L82
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v454 + int32(1)
	goto L4
L148:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L19
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_11), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L19
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1759), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L19
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_10), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L19
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1763), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L19
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
	F_errcode(m, int32(67240066))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_11), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1781), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L19
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
	F_errcode(m, int32(67240066))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_10), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1790), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L19
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
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_appendBinaryStringInfo(m, v18, v579+v319, v575-v319)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L19
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v574 == int32(0) {
		goto L3
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v575
	goto L166
L168:
	;
	goto L4
L169:
	;
	goto L170
L170:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v625 = F_CopyGetData(m, l0, v623, int32(_a_F_CopyReadLine_12))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L19
	} else {
		goto L172
	}
L171:
	;
	v629 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v629
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v629
	v678 = v603
	goto L2
L172:
	;
	if int32(0) < v625 {
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v671 = v669 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v671
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v675 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v673+v671))) = uint8(v675)
	v678 = v649
	goto L2
L175:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v663 = v661 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v663
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v667 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v665+v663))) = uint8(v667)
	v678 = v649
	goto L2
L176:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v655 = v653 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v655
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v659 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v657+v655))) = uint8(v659)
	v678 = v649
	goto L2
}
func F_line_construct_pp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 float64
	_ = v28
	var v31 int64
	_ = v31
	var v38 float64
	_ = v38
	var v45 int64
	_ = v45
	var v50 float64
	_ = v50
	var v68 int32
	_ = v68
	var v75 float64
	_ = v75
	var v79 int64
	_ = v79
	var v80 float64
	_ = v80
	var v83 int64
	_ = v83
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_palloc(m, int32(24))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		if base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			v26 = int64(9223372036854775807)
			v27 = base.I64_reinterpret_f64(v24) & v26
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v31 = base.I64_reinterpret_f64(v28) & v26
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v31) {
				v68 = base.B2i32(base.Ui64(v27) < base.Ui64(int64(9218868437227405313)))
				if base.B2i32(v68 == int32(0))|base.F64_ne(v18, v24) != 0 {
					v122 = F_point_sl(m, v12, v11)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						F_line_construct(m, v14, v12, v122)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v14)
						}
					}
				} else {
					v75 = v28
					v79 = v31
					v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
					if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
						if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v14)
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v14)
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v27) {
					v122 = F_point_sl(m, v12, v11)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						F_line_construct(m, v14, v12, v122)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v14)
						}
					}
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					if base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v18, v24)), float64(1e-06)) == int32(0))&base.F64_ne(v18, v24) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v14)
								}
							}
						} else {
							if base.F64_eq(v28, v38)|base.F64_le(base.F64_abs(base.F64_sub(v28, v38)), float64(1e-06)) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v122 = F_point_sl(m, v12, v11)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int64(0)
								} else {
									F_line_construct(m, v14, v12, v122)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v14)
									}
								}
							}
						}
					} else {
						v68 = int32(1)
						if base.B2i32(v68 == int32(0))|base.F64_ne(v18, v24) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v14)
								}
							}
						} else {
							v75 = v28
							v79 = v31
							v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
							v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
							if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
								if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
									v122 = F_point_sl(m, v12, v11)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int64(0)
									} else {
										F_line_construct(m, v14, v12, v122)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v14)
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
									v122 = F_point_sl(m, v12, v11)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int64(0)
									} else {
										F_line_construct(m, v14, v12, v122)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v14)
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
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
		} else {
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(v45&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
				v122 = F_point_sl(m, v12, v11)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int64(0)
				} else {
					F_line_construct(m, v14, v12, v122)
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v14)
					}
				}
			} else {
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
				v75 = v50
				v79 = base.I64_reinterpret_f64(v50) & int64(9223372036854775807)
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
				if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
						v122 = F_point_sl(m, v12, v11)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int64(0)
						} else {
							F_line_construct(m, v14, v12, v122)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v14)
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
						v122 = F_point_sl(m, v12, v11)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int64(0)
						} else {
							F_line_construct(m, v14, v12, v122)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v14)
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1171), int32(_a_F_line_construct_pp_2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
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
func F_line_horizontal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	return base.I64_extend_i32_u(base.F64_le(base.F64_abs(v3), float64(1e-06)))
}
func F_line_vertical(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	return base.I64_extend_i32_u(base.F64_le(base.F64_abs(v3), float64(1e-06)))
}
