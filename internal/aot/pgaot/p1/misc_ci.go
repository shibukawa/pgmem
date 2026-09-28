package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cidr_abbrev(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v634 int32
	_ = v634
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v27&v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v25
	goto L5
L4:
	;
	v30 = int32(4)
	goto L5
L5:
	;
	v31 = v21 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v33 = int32(2)
	v34 = v31 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v36 = int32(50)
	v37 = m.G0
	v39 = v37 - int32(240)
	m.G0 = v39
	switch v32 - v33 {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L6:
	;
	m.G0 = v39 + int32(240)
	if v634 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L7:
	;
	v634 = int32(0)
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L9:
	;
	if base.Ui32(v585) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L158
	}
L10:
	;
	v550 = v35 & int32(7)
	if v550 == int32(0) {
		v582 = v537
		v585 = v540
		goto L9
	} else {
		goto L152
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(5)
	goto L7
L12:
	;
	if base.Ui32(int32(129)) <= base.Ui32(v35) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	if base.Ui32(int32(33)) <= base.Ui32(v35) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L15:
	;
	goto L16
L16:
	;
	if v35 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v50 = int32(48)
	*(*uint16)(unsafe.Add(mBase, uint32(v18))) = uint16(v50)
	v582 = v16 + int32(-63)
	v585 = int32(49)
	goto L9
L18:
	;
	goto L19
L19:
	;
	v56 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	if v56 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v537 = v18
	v539 = v34
	v540 = v36
	goto L10
L21:
	;
	goto L22
L22:
	;
	v60 = v18
	v63 = v56
	v64 = v34
	v65 = v36
	goto L23
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v74
	v77 = v64 + int32(1)
	v81 = F_pg_sprintf(m, v60, int32(_a_F_cidr_abbrev_0), v39+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L8
L25:
	;
	v83 = v81 + v60
	if v63 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v537 = v83
	v539 = v77
	v540 = v60 + v65 - v83
	goto L10
L27:
	;
	goto L28
L28:
	;
	v88 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v88)
	v90 = int32(1)
	v94 = v83 + v90
	v95 = v60 + v65 - v94
	if base.Ui32(int32(6)) <= base.Ui32(v95) {
		v60 = v94
		v63 = v63 - v90
		v64 = v77
		v65 = v95
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L31:
	;
	goto L32
L32:
	;
	if v35 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v35
	v445 = F_pg_sprintf(m, v440, int32(_a_F_cidr_abbrev_1), v39+int32(48))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L127
	}
L34:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+162)) = uint8(v105)
	v107 = int32(_a_F_cidr_abbrev_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+160)) = uint16(v107)
	v440 = v39 + int32(160) | int32(2)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v116 = int32(base.Ui32(v35+int32(7)) >> (uint(int32(3)) % 32))
	if v116 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	base.MemoryCopy(m, v39+int32(224), v34, v116)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v122 = v39 + int32(224) + v116
	v124 = int32(16) - v116
	if v124 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryFill(m, v122, int32(0), v124)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v128 = v35 & int32(7)
	if v128 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v122 - int32(1)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v136 = v131 & (int32(-1) << (uint(int32(8)-v128) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v136)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v143 = int32(base.Ui32(v35+int32(15)) >> (uint(int32(4)) % 32))
	if v143 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v146 = int32(2)
	goto L48
L47:
	;
	v146 = v143
	goto L48
L48:
	;
	v149 = int32(0)
	v152 = v149
	v154 = v2
	v156 = v149
	v159 = v2
	v160 = v2
	goto L49
L49:
	;
	v168 = v39 + int32(224) + v156
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169|v170 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v190 = int32(0)
	v194 = base.B2i32(v183 != v190) & base.B2i32(v185 < v183)
	if v194 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v188 = v156 + int32(2)
	if base.Ui32(v188) < base.Ui32(v146<<(uint(int32(1))%32)) {
		v152 = v183
		v154 = v184
		v156 = v188
		v159 = v185
		v160 = v186
		goto L49
	} else {
		goto L62
	}
L52:
	;
	if v152 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v152 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v176 = v154
	goto L57
L56:
	;
	v176 = int32(base.Ui32(v156) >> (uint(int32(1)) % 32))
	goto L57
L57:
	;
	v183 = v152 + int32(1)
	v184 = v176
	v185 = v159
	v186 = v160
	goto L51
L58:
	;
	if v152 <= v159 {
		v183 = v152
		v184 = v154
		v185 = v159
		v186 = v160
		goto L51
	} else {
		goto L61
	}
L59:
	;
	v180 = v159
	v181 = v160
	goto L60
L60:
	;
	v183 = int32(0)
	v184 = v154
	v185 = v180
	v186 = v181
	goto L51
L61:
	;
	v180 = v152
	v181 = v154
	goto L60
L62:
	;
	goto L50
L63:
	;
	v215 = v196 + v195
	v217 = v146 - int32(1)
	if v217 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	v195 = v184
	goto L66
L65:
	;
	v195 = v186
	goto L66
L66:
	;
	if v194 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v196 = v183
	goto L69
L68:
	;
	v196 = v185
	goto L69
L69:
	;
	if v195|base.B2i32(v196 == v146) != 0 {
		v214 = v2
		goto L63
	} else {
		goto L70
	}
L70:
	;
	switch v196 - int32(5) {
	case 0:
		goto L73
	case 1:
		goto L71
	case 2:
		goto L72
	default:
		v214 = v2
		goto L63
	}
L71:
	;
	v214 = int32(1)
	goto L63
L72:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+238)))
	if v207 == int32(0) {
		v214 = v2
		goto L63
	} else {
		goto L76
	}
L73:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+234)))
	if v201 != int32(255) {
		v214 = v2
		goto L63
	} else {
		goto L74
	}
L74:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+235)))
	if v204 == int32(255) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v214 = v2
	goto L63
L76:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+239)))
	if v210 == int32(1) {
		v214 = v2
		goto L63
	} else {
		goto L77
	}
L77:
	;
	goto L71
L78:
	;
	v346 = int32(0)
	if base.B2i32(base.B2i32(v196 == v346)|base.B2i32(v335 < v195) == v346)&base.B2i32(v335 < v215) == v346 {
		goto L107
	} else {
		goto L108
	}
L79:
	;
	v332 = v39 + int32(160)
	v334 = v39 + int32(224)
	v335 = v190
	goto L78
L80:
	;
	goto L81
L81:
	;
	v233 = v39 + int32(160)
	v235 = v39 + int32(224)
	v236 = v190
	goto L82
L82:
	;
	v247 = int32(0)
	if base.B2i32(v196 == v247)|base.B2i32(v236 < v195)|base.B2i32(v215 <= v236) == v247 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v332 = v325
	v334 = v326
	v335 = v329
	goto L78
L84:
	;
	v329 = v236 + int32(1)
	if v146-int32(2) != v236 {
		v233 = v325
		v235 = v326
		v236 = v329
		goto L82
	} else {
		goto L106
	}
L85:
	;
	if v236 == v195 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v236))&v214 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v256 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v256)
	v260 = v233 + int32(1)
	goto L90
L89:
	;
	v260 = v233
	goto L90
L90:
	;
	v325 = v260
	v326 = v235 + int32(2)
	goto L84
L91:
	;
	if v236 == int32(6) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v304 = v39 + int32(160)
	if v304 == v233 {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	v270 = int32(58)
	goto L96
L95:
	;
	v270 = int32(46)
	goto L96
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v270)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v272
	v275 = v233 + int32(1)
	v279 = F_pg_sprintf(m, v275, int32(_a_F_cidr_abbrev_0), v39+int32(128))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v281 = v279 + v275
	if base.B2i32(base.Ui32(int32(120)) < base.Ui32(v35))|base.B2i32(v236 != int32(7)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v325 = v281
	v326 = v235 + int32(1)
	goto L84
L99:
	;
	goto L100
L100:
	;
	v289 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v289)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v291
	v294 = v281 + int32(1)
	v298 = F_pg_sprintf(m, v294, int32(_a_F_cidr_abbrev_0), v39+int32(112))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v325 = v298 + v294
	v326 = v235 + int32(2)
	goto L84
L102:
	;
	v310 = v304
	goto L104
L103:
	;
	v306 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v306)
	v310 = v233 + int32(1)
	goto L104
L104:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v311 | v312<<(uint(int32(8))%32)
	v322 = F_pg_sprintf(m, v310, int32(_a_F_cidr_abbrev_3), v39+int32(144))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v325 = v322 + v310
	v326 = v235 + int32(2)
	goto L84
L106:
	;
	goto L83
L107:
	;
	if v214&base.B2i32(base.Ui32(int32(5)) < base.Ui32(v335)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	if v335 == v195 {
		goto L123
	} else {
		goto L124
	}
L110:
	;
	v362 = v39 + int32(160)
	if v362 == v332 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v335 == int32(6) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v368 = v362
	goto L115
L114:
	;
	v364 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v364)
	v368 = v332 + int32(1)
	goto L115
L115:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v369 | v370<<(uint(int32(8))%32)
	v378 = F_pg_sprintf(m, v368, int32(_a_F_cidr_abbrev_3), v39+int32(96))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v440 = v378 + v368
	goto L33
L117:
	;
	v385 = int32(58)
	goto L119
L118:
	;
	v385 = int32(46)
	goto L119
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v385)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v387
	v390 = v332 + int32(1)
	v394 = F_pg_sprintf(m, v390, int32(_a_F_cidr_abbrev_0), v39+int32(80))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v396 = v394 + v390
	if base.B2i32(v335 == int32(7))&base.B2i32(base.Ui32(v35) <= base.Ui32(int32(120))) != 0 {
		v440 = v396
		goto L33
	} else {
		goto L121
	}
L121:
	;
	v402 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v402)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v404
	v407 = v396 + int32(1)
	v411 = F_pg_sprintf(m, v407, int32(_a_F_cidr_abbrev_0), v39-int32(-64))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v440 = v411 + v407
	goto L33
L123:
	;
	v415 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v415)
	v419 = v332 + int32(1)
	goto L125
L124:
	;
	v419 = v332
	goto L125
L125:
	;
	if v335 != v217 {
		v440 = v419
		goto L33
	} else {
		goto L126
	}
L126:
	;
	v421 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v419))) = uint8(v421)
	v440 = v419 + int32(1)
	goto L33
L127:
	;
	v448 = v39 + int32(160)
	v449 = F_strlen(m, v448)
	mBase = m.M
	if base.Ui32(v449+int32(1)) <= base.Ui32(int32(50)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if (v448^v18)&int32(3) != 0 {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L131:
	;
	v634 = v18
	goto L6
L132:
	;
	goto L131
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v507)
	if v507&int32(255) == int32(0) {
		goto L132
	} else {
		goto L148
	}
L134:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v506 = v448
	v507 = v459
	v508 = v18
	goto L133
L135:
	;
	goto L136
L136:
	;
	if v448&int32(3) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v463 = v448
	v465 = v18
	goto L140
L138:
	;
	v477 = v448
	v479 = v18
	goto L139
L139:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v484 = int32(-2139062144)
	if (int32(16843008)-v481|v481)&v484 != v484 {
		v506 = v477
		v507 = v481
		v508 = v479
		goto L133
	} else {
		goto L144
	}
L140:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v466)
	if v466 == int32(0) {
		goto L132
	} else {
		goto L142
	}
L141:
	;
	v477 = v473
	v479 = v471
	goto L139
L142:
	;
	v470 = int32(1)
	v471 = v465 + v470
	v473 = v463 + v470
	if v473&int32(3) != 0 {
		v463 = v473
		v465 = v471
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v489 = v477
	v490 = v481
	v491 = v479
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v490
	v493 = int32(4)
	v494 = v491 + v493
	v496 = v489 + v493
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	v501 = int32(-2139062144)
	if (int32(16843008)-v498|v498)&v501 == v501 {
		v489 = v496
		v490 = v498
		v491 = v494
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v506 = v496
	v507 = v498
	v508 = v494
	goto L133
L147:
	;
	goto L146
L148:
	;
	v515 = v506
	v517 = v508
	goto L149
L149:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)) = uint8(v518)
	v520 = int32(1)
	if v518 != 0 {
		v515 = v515 + v520
		v517 = v517 + v520
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L132
L151:
	;
	goto L150
L152:
	;
	if base.Ui32(v540) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	if v537 != v18 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v556 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v556)
	v560 = v537 + int32(1)
	goto L156
L155:
	;
	v560 = v18
	goto L156
L156:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	v562 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v561 & ((v562<<(uint(v550)%32) ^ v562) << (uint(int32(8)-v550) % 32))
	v575 = F_pg_sprintf(m, v560, int32(_a_F_cidr_abbrev_0), v39+int32(16))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v577 = v575 + v560
	v582 = v577
	v585 = v537 + v540 - v577
	goto L9
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v35
	v598 = F_pg_sprintf(m, v582, int32(_a_F_cidr_abbrev_1), v39)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v634 = v18
	goto L6
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v670 = F_cstring_to_text(m, v18)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L167
	}
L163:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errmsg(m, int32(_a_F_cidr_abbrev_4), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_cidr_abbrev_5), int32(1185), int32(_a_F_cidr_abbrev_6))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	m.G0 = v18 - int32(-64)
	return base.I64_extend_i32_u(v670)
}
func F_cidr_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_network_out(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_cidr_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_network_send(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_cidr_set_masklen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v17 == int32(-1) {
			v22 = int32(1)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v26 = v24 & v22
			if v26 != 0 {
				v27 = v22
			} else {
				v27 = int32(4)
			}
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v27))))
			if v29 == int32(2) {
				v32 = int32(32)
			} else {
				v32 = int32(128)
			}
			v38 = v32
			v39 = v26
			if v39 != 0 {
				v44 = int32(1)
			} else {
				v44 = int32(4)
			}
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
			if v46 == int32(2) {
				v49 = int32(32)
			} else {
				v49 = int32(128)
			}
			if base.Ui32(v49) < base.Ui32(v38) {
				v116 = v38
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(355), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
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
				v52 = F_palloc0(m, int32(22))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					v54 = int32(1)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					if v56&v54 != 0 {
						v59 = v54
					} else {
						v59 = int32(4)
					}
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v59))))
					v62 = int32(1)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					if v64&v62 != 0 {
						v67 = v62
					} else {
						v67 = int32(4)
					}
					v68 = v52 + v67
					*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)) = uint8(v38)
					*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v61)
					if v38 == int32(0) {
					} else {
						v74 = v68 + int32(2)
						v78 = int32(base.Ui32(v38+int32(7)) >> (uint(int32(3)) % 32))
						if v78 != 0 {
							v79 = int32(1)
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v81&v79 != 0 {
								v84 = v79
							} else {
								v84 = int32(4)
							}
							base.MemoryCopy(m, v74, v13+v84+int32(2), v78)
						} else {
						}
						v90 = v38 & int32(7)
						if v90 == int32(0) {
						} else {
							v95 = v74 + int32(base.Ui32(v38)>>(uint(int32(3))%32))
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
							v99 = v96 & (int32(-256) >> (uint(v90) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v99)
						}
					}
					if v61 == int32(2) {
						v109 = int32(40)
					} else {
						v109 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v109
					m.G0 = v10 + int32(16)
					return base.I64_extend_i32_u(v52)
				}
			}
		} else {
			if v17 < int32(0) {
				v116 = v17
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(355), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
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
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v38 = v17
				v39 = v35 & int32(1)
				if v39 != 0 {
					v44 = int32(1)
				} else {
					v44 = int32(4)
				}
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
				if v46 == int32(2) {
					v49 = int32(32)
				} else {
					v49 = int32(128)
				}
				if base.Ui32(v49) < base.Ui32(v38) {
					v116 = v38
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
							F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(355), int32(_a_F_cidr_set_masklen_2))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
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
					v52 = F_palloc0(m, int32(22))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = int32(1)
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v56&v54 != 0 {
							v59 = v54
						} else {
							v59 = int32(4)
						}
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v59))))
						v62 = int32(1)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v64&v62 != 0 {
							v67 = v62
						} else {
							v67 = int32(4)
						}
						v68 = v52 + v67
						*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)) = uint8(v38)
						*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v61)
						if v38 == int32(0) {
						} else {
							v74 = v68 + int32(2)
							v78 = int32(base.Ui32(v38+int32(7)) >> (uint(int32(3)) % 32))
							if v78 != 0 {
								v79 = int32(1)
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
								if v81&v79 != 0 {
									v84 = v79
								} else {
									v84 = int32(4)
								}
								base.MemoryCopy(m, v74, v13+v84+int32(2), v78)
							} else {
							}
							v90 = v38 & int32(7)
							if v90 == int32(0) {
							} else {
								v95 = v74 + int32(base.Ui32(v38)>>(uint(int32(3))%32))
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
								v99 = v96 & (int32(-256) >> (uint(v90) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v99)
							}
						}
						if v61 == int32(2) {
							v109 = int32(40)
						} else {
							v109 = int32(88)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v109
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v52)
					}
				}
			}
		}
	}
}
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v6 = int32(1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v10 = v8 & v6
	if v10 != 0 {
		v11 = v6
	} else {
		v11 = int32(4)
	}
	if v8 == int32(1) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v18 == int32(18) {
			v21 = int32(16)
		} else {
			v21 = int32(0)
		}
		if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v28 = int32(4)
		} else {
			v28 = v21
		}
		v39 = v28
	} else {
		v29 = int32(1)
		if v10 != 0 {
			v39 = int32(base.Ui32(v8)>>(uint(v29)%32)) - v29
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v41 = F_str_tolower(m, l0+v11, v39, int32(100))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return int32(0)
	} else {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v45 == int32(1) {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v51 == int32(18) {
				v54 = int32(16)
			} else {
				v54 = int32(0)
			}
			if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v61 = int32(4)
			} else {
				v61 = v54
			}
			v74 = v61
		} else {
			v62 = int32(1)
			if v45&v62 != 0 {
				v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
			} else {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v75 = int32(1)
		if v45&v75 != 0 {
			v79 = v75
		} else {
			v79 = int32(4)
		}
		v82 = F_str_tolower(m, l1+v79, v74, int32(100))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			v84 = F_strlen(m, v41)
			mBase = m.M
			v85 = F_strlen(m, v82)
			mBase = m.M
			v86 = F_varstr_cmp(m, v41, v84, v82, v85, l2)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v41)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v82)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						return v86
					}
				}
			}
		}
	}
}
