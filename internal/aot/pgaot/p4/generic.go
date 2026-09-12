package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenericXLogFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
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
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v456 int64
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[147]))))
	if v12 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v698 = int32(4543684)
	v700 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v700 - int32(1)
	F_pfree(m, l0)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L222
	}
L2:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v557 = int32(4543684)
	v559 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v559 + int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[143])))
	if v563 != 0 {
		goto L176
	} else {
		goto L177
	}
L5:
	;
	return
L6:
	;
	v17 = int32(4543684)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v19 + int32(1)
	v27 = int32(0)
	goto L7
L7:
	;
	v38 = l0 + int32(32768) + v27*int32(8216)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v456 = F_XLogInsert(m, int32(20), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L149
	}
L9:
	;
	v451 = v27 + int32(1)
	if v451 != int32(4) {
		v27 = v451
		goto L7
	} else {
		goto L148
	}
L10:
	;
	if v39 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
	if v61&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+(v39^int32(-1))<<(uint(int32(2))%32))))
	v59 = v51
	goto L11
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v59 = v53 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	v407 = v60
	goto L17
L16:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+14)))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v68
	v77 = int32(-1)
	goto L20
L17:
	;
	v408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	if v408 != 0 {
		goto L133
	} else {
		goto L134
	}
L18:
	;
	v236 = int32(8192)
	v243 = int32(-1)
	v245 = base.B2i32(base.Ui32(v65) < base.Ui32(v64))
	if base.Ui32(v65) < base.Ui32(v64) {
		goto L76
	} else {
		goto L77
	}
L20:
	;
	goto L21
L21:
	;
	goto L24
L22:
	;
	if v200 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v67) < base.Ui32(v66) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v83 = v67
	goto L28
L27:
	;
	v83 = v66
	goto L28
L28:
	;
	if base.Ui32(v83) <= base.Ui32(v68) {
		v200 = v77
		v202 = v77
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v86 = v38 + int32(16)
	v90 = v68
	v94 = v77
	goto L30
L30:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v90))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v90))))
	if v100 != v102 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v200 = v190
	v202 = v191
	goto L22
L32:
	;
	if v94 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v131 = v90
	v135 = v94
	goto L34
L34:
	;
	v141 = v131 + int32(1)
	if v141 < v83 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v106 = v90
	goto L37
L36:
	;
	v106 = v94
	goto L37
L37:
	;
	v110 = v90
	goto L38
L38:
	;
	v120 = v110 + int32(1)
	if v83 <= v120 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v131 = v120
	v135 = v106
	goto L34
L40:
	;
	v200 = v106
	v202 = int32(-1)
	goto L22
L41:
	;
	goto L42
L42:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v120))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v120))))
	if v124 != v126 {
		v110 = v120
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v143 = v83
	goto L46
L45:
	;
	v143 = v141
	goto L46
L46:
	;
	v149 = v131
	goto L47
L47:
	;
	if v149 == v143-int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v135 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	v166 = v143
	goto L49
L51:
	;
	goto L52
L52:
	;
	v160 = v149 + int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v160))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v160))))
	if v162 == v164 {
		v149 = v160
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v166 = v160
	goto L49
L54:
	;
	if v166 < v83 {
		v90 = v166
		v94 = v190
		goto L30
	} else {
		goto L59
	}
L55:
	;
	v190 = int32(-1)
	v191 = v131
	goto L54
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v166-v131) < base.Ui32(int32(5)) {
		v190 = v135
		v191 = v131
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v174 = v86 + v173
	*(*uint16)(unsafe.Add(mBase, uint32(v174))) = uint16(v135)
	v176 = v131 - v135
	*(*uint16)(unsafe.Add(mBase, uint32(v174)+2)) = uint16(v176)
	v182 = v176 & int32(65535)
	v183 = F___memcpy(m, v174+int32(4), v60+v135, v182)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v183 + v182 - v86
	v187 = int32(-1)
	v190 = v187
	v191 = v187
	goto L54
L59:
	;
	goto L31
L60:
	;
	v207 = v83
	goto L62
L61:
	;
	v207 = v200
	goto L62
L62:
	;
	v208 = base.B2i32(base.Ui32(v66) < base.Ui32(v67))
	if base.Ui32(v66) < base.Ui32(v67) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v209 = v207
	goto L65
L64:
	;
	v209 = v200
	goto L65
L65:
	;
	if int32(0) <= v209 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v213 = v38 + int32(16)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v215 = v213 + v214
	*(*uint16)(unsafe.Add(mBase, uint32(v215))) = uint16(v209)
	if base.Ui32(v66) < base.Ui32(v67) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	goto L18
L69:
	;
	v217 = v67
	goto L71
L70:
	;
	v217 = v202
	goto L71
L71:
	;
	if v217 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v220 = v67
	goto L74
L73:
	;
	v220 = v217
	goto L74
L74:
	;
	v221 = v220 - v209
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+2)) = uint16(v221)
	v227 = v221 & int32(65535)
	v228 = F___memcpy(m, v215+int32(4), v60+v209, v227)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v228 + v227 - v213
	goto L68
L75:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v407 = v402
	goto L17
L76:
	;
	v246 = v65
	goto L78
L77:
	;
	v246 = v243
	goto L78
L78:
	;
	if base.Ui32(v65) < base.Ui32(v64) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L117
L80:
	;
	v247 = v64
	goto L82
L81:
	;
	v247 = v65
	goto L82
L82:
	;
	goto L84
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(v236) <= base.Ui32(v247) {
		v366 = v246
		v368 = v243
		goto L79
	} else {
		goto L86
	}
L86:
	;
	v252 = v38 + int32(16)
	v256 = v247
	v260 = v246
	goto L87
L87:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v256))))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v256))))
	if v266 != v268 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v366 = v356
	v368 = v357
	goto L79
L89:
	;
	if v260 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v297 = v256
	v301 = v260
	goto L91
L91:
	;
	v307 = v297 + int32(1)
	if v307 < v236 {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	v272 = v256
	goto L94
L93:
	;
	v272 = v260
	goto L94
L94:
	;
	v276 = v256
	goto L95
L95:
	;
	v286 = v276 + int32(1)
	if v236 <= v286 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v297 = v286
	v301 = v272
	goto L91
L97:
	;
	v366 = v272
	v368 = int32(-1)
	goto L79
L98:
	;
	goto L99
L99:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v286))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v286))))
	if v290 != v292 {
		v276 = v286
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	v309 = v236
	goto L103
L102:
	;
	v309 = v307
	goto L103
L103:
	;
	v315 = v297
	goto L104
L104:
	;
	if v315 == v309-int32(1) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if v301 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	v332 = v309
	goto L106
L108:
	;
	goto L109
L109:
	;
	v326 = v315 + int32(1)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v326))))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v326))))
	if v328 == v330 {
		v315 = v326
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v332 = v326
	goto L106
L111:
	;
	if v332 < v236 {
		v256 = v332
		v260 = v356
		goto L87
	} else {
		goto L116
	}
L112:
	;
	v356 = int32(-1)
	v357 = v297
	goto L111
L113:
	;
	goto L114
L114:
	;
	if base.Ui32(v332-v297) < base.Ui32(int32(5)) {
		v356 = v301
		v357 = v297
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v340 = v252 + v339
	*(*uint16)(unsafe.Add(mBase, uint32(v340))) = uint16(v301)
	v342 = v297 - v301
	*(*uint16)(unsafe.Add(mBase, uint32(v340)+2)) = uint16(v342)
	v348 = v342 & int32(65535)
	v349 = F___memcpy(m, v340+int32(4), v60+v301, v348)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v349 + v348 - v252
	v353 = int32(-1)
	v356 = v353
	v357 = v353
	goto L111
L116:
	;
	goto L88
L117:
	;
	goto L119
L119:
	;
	goto L121
L121:
	;
	goto L122
L122:
	;
	if int32(0) <= v366 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v379 = v38 + int32(16)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v381 = v379 + v380
	*(*uint16)(unsafe.Add(mBase, uint32(v381))) = uint16(v366)
	goto L127
L124:
	;
	goto L125
L125:
	;
	goto L75
L127:
	;
	goto L128
L128:
	;
	if v368 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v386 = v236
	goto L131
L130:
	;
	v386 = v368
	goto L131
L131:
	;
	v387 = v386 - v366
	*(*uint16)(unsafe.Add(mBase, uint32(v381)+2)) = uint16(v387)
	v393 = v387 & int32(65535)
	v394 = F___memcpy(m, v381+int32(4), v60+v366, v393)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v394 + v393 - v379
	goto L125
L132:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v417 = F__emscripten_memset_bulkmem(m, v410+v411, base.I32_extend8_s(int32(0)), v414-v411)
	mBase = m.M
	goto L136
L133:
	;
	v409 = F__emscripten_memcpy_bulkmem(m, v59, v407, v408)
	mBase = m.M
	v410 = v409
	goto L135
L134:
	;
	v410 = v59
	goto L135
L135:
	;
	goto L132
L136:
	;
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v423 = int32(8192) - v418
	if v423 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_MarkBufferDirty(m, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L141
	}
L138:
	;
	v424 = F__emscripten_memcpy_bulkmem(m, v410+v418, v420+v418, v423)
	mBase = m.M
	goto L140
L139:
	;
	goto L140
L140:
	;
	goto L137
L141:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
	if v430&int32(1) != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_XLogRegisterBuffer(m, v27, v429, int32(9))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	F_XLogRegisterBuffer(m, v27, v429, int32(8))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L5
	} else {
		goto L146
	}
L145:
	;
	goto L9
L146:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	F_XLogRegisterBufData(m, v27, v38+int32(16), v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L5
	} else {
		goto L147
	}
L147:
	;
	goto L9
L148:
	;
	goto L8
L149:
	;
	v458 = base.I32_wrap_i64(v456)
	v461 = base.I32_wrap_i64(int64(base.Ui64(v456) >> (uint(int64(32)) % 64)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[143])))
	if v462 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v462 < int32(0) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	goto L152
L152:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[144])))
	if v486 != 0 {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v480)+4)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = v461
	goto L152
L154:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v466+(v462^int32(-1))<<(uint(int32(2))%32))))
	v480 = v472
	goto L153
L155:
	;
	goto L156
L156:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v480 = v474 + v462<<(uint(int32(13))%32) + int32(-8192)
	goto L153
L157:
	;
	if int32(0) <= v486 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[145])))
	if v510 != 0 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+4)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v461
	goto L159
L161:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v504 = v490 + v486<<(uint(int32(13))%32) + int32(-8192)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v497+(v486^int32(-1))<<(uint(int32(2))%32))))
	v504 = v503
	goto L160
L164:
	;
	if int32(0) <= v510 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[146])))
	if v534 == int32(0) {
		goto L1
	} else {
		goto L171
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528)+4)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v461
	goto L166
L168:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v528 = v514 + v510<<(uint(int32(13))%32) + int32(-8192)
	goto L167
L169:
	;
	goto L170
L170:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v521+(v510^int32(-1))<<(uint(int32(2))%32))))
	v528 = v527
	goto L167
L171:
	;
	if int32(0) <= v534 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v554)+4)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v554))) = v461
	goto L1
L173:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v554 = v540 + v534<<(uint(int32(13))%32) + int32(-8192)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v547+(v534^int32(-1))<<(uint(int32(2))%32))))
	v554 = v553
	goto L172
L176:
	;
	if v563 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L178
L178:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[144])))
	if v591 != 0 {
		goto L188
	} else {
		goto L189
	}
L179:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[148])))
	goto L184
L180:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v567+(v563^int32(-1))<<(uint(int32(2))%32))))
	v581 = v573
	goto L179
L181:
	;
	goto L182
L182:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v581 = v575 + v563<<(uint(int32(13))%32) + int32(-8192)
	goto L179
L183:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[143])))
	F_MarkBufferDirty(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L187
	}
L184:
	;
	v584 = F__emscripten_memcpy_bulkmem(m, v581, v582, int32(8192))
	mBase = m.M
	goto L186
L186:
	;
	goto L183
L187:
	;
	goto L178
L188:
	;
	if int32(0) <= v591 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L190
L190:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[145])))
	if v623 != 0 {
		goto L200
	} else {
		goto L201
	}
L191:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[149])))
	goto L196
L192:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v609 = v595 + v591<<(uint(int32(13))%32) + int32(-8192)
	goto L191
L193:
	;
	goto L194
L194:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v602+(v591^int32(-1))<<(uint(int32(2))%32))))
	v609 = v608
	goto L191
L195:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[144])))
	F_MarkBufferDirty(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L199
	}
L196:
	;
	v614 = F__emscripten_memcpy_bulkmem(m, v609, v612, int32(8192))
	mBase = m.M
	goto L198
L198:
	;
	goto L195
L199:
	;
	goto L190
L200:
	;
	if int32(0) <= v623 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	goto L202
L202:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[146])))
	if v655 == int32(0) {
		goto L1
	} else {
		goto L212
	}
L203:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[150])))
	goto L208
L204:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v641 = v627 + v623<<(uint(int32(13))%32) + int32(-8192)
	goto L203
L205:
	;
	goto L206
L206:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v634+(v623^int32(-1))<<(uint(int32(2))%32))))
	v641 = v640
	goto L203
L207:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[145])))
	F_MarkBufferDirty(m, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L5
	} else {
		goto L211
	}
L208:
	;
	v646 = F__emscripten_memcpy_bulkmem(m, v641, v644, int32(8192))
	mBase = m.M
	goto L210
L210:
	;
	goto L207
L211:
	;
	goto L202
L212:
	;
	if int32(0) <= v655 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[151])))
	goto L218
L214:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v675 = v661 + v655<<(uint(int32(13))%32) + int32(-8192)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v668+(v655^int32(-1))<<(uint(int32(2))%32))))
	v675 = v674
	goto L213
L217:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[146])))
	F_MarkBufferDirty(m, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L221
	}
L218:
	;
	v680 = F__emscripten_memcpy_bulkmem(m, v675, v678, int32(8192))
	mBase = m.M
	goto L220
L220:
	;
	goto L217
L221:
	;
	goto L1
L222:
	;
	return
}
func F_GenericXLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l0 + int32(32768)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[143])))
	if v13 == int32(0) {
		v52 = v12
		*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
		if l1 < int32(0) {
			v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l1^int32(-1))<<(uint(int32(2))%32))))
			v76 = v68
		} else {
			v70 = *(*int32)(unsafe.Add(mBase, _consts[2]))
			v76 = v70 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		v78 = F__emscripten_memcpy_bulkmem(m, v56, v76, int32(8192))
		mBase = m.M
		v86 = v52 + int32(12)
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
		m.G0 = v9 + int32(16)
		return v88
	} else {
		if l1 == v13 {
			v81 = v12
			v86 = v81 + int32(12)
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
			m.G0 = v9 + int32(16)
			return v88
		} else {
			v18 = l0 + int32(40984)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[144])))
			if v19 == int32(0) {
				v52 = v18
				*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
				if l1 < int32(0) {
					v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l1^int32(-1))<<(uint(int32(2))%32))))
					v76 = v68
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v76 = v70 + l1<<(uint(int32(13))%32) + int32(-8192)
				}
				v78 = F__emscripten_memcpy_bulkmem(m, v56, v76, int32(8192))
				mBase = m.M
				v86 = v52 + int32(12)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				m.G0 = v9 + int32(16)
				return v88
			} else {
				if l1 == v19 {
					v81 = v18
					v86 = v81 + int32(12)
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					m.G0 = v9 + int32(16)
					return v88
				} else {
					v24 = l0 + int32(49200)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[145])))
					if v25 == int32(0) {
						v52 = v24
						*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
						if l1 < int32(0) {
							v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l1^int32(-1))<<(uint(int32(2))%32))))
							v76 = v68
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v76 = v70 + l1<<(uint(int32(13))%32) + int32(-8192)
						}
						v78 = F__emscripten_memcpy_bulkmem(m, v56, v76, int32(8192))
						mBase = m.M
						v86 = v52 + int32(12)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
						m.G0 = v9 + int32(16)
						return v88
					} else {
						if l1 == v25 {
							v81 = v24
							v86 = v81 + int32(12)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
							m.G0 = v9 + int32(16)
							return v88
						} else {
							v30 = l0 + int32(57416)
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[146])))
							if v31 == int32(0) {
								v52 = v30
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
								if l1 < int32(0) {
									v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l1^int32(-1))<<(uint(int32(2))%32))))
									v76 = v68
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v76 = v70 + l1<<(uint(int32(13))%32) + int32(-8192)
								}
								v78 = F__emscripten_memcpy_bulkmem(m, v56, v76, int32(8192))
								mBase = m.M
								v86 = v52 + int32(12)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
								m.G0 = v9 + int32(16)
								return v88
							} else {
								if v31 == l1 {
									v81 = v30
									v86 = v81 + int32(12)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
									m.G0 = v9 + int32(16)
									return v88
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(4)
										F_errmsg_internal(m, int32(479823), v9)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(518538), int32(327), int32(236028))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
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
						}
					}
				}
			}
		}
	}
}
func F_Generic_Text_IC_like(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	if l2 != 0 {
		v10 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v230 = m.ExcPending
					if v230 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(562456), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519598), int32(202), int32(414639))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[495]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_consts[1299])))
				if v24 <= int32(1) {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
					if v27 != int32(105) {
						v127 = int32(1)
						v128 = l1 + v127
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v131 = v129 & v127
						if v129 == v127 {
							v134 = int32(4)
							v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
							if v136&int32(254) == int32(2) {
								v145 = v134
							} else {
								v145 = base.B2i32(v136 == int32(18)) << (uint(v134) % 32)
							}
							if v136 == int32(1) {
								v148 = v134
							} else {
								v148 = v145
							}
							v159 = v148
						} else {
							v149 = int32(1)
							if v131 != 0 {
								v159 = int32(base.Ui32(v129)>>(uint(v149)%32)) - v149
							} else {
								v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v159 = int32(base.Ui32(v153)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v131 != 0 {
							v160 = v128
						} else {
							v160 = l1 + int32(4)
						}
						v161 = int32(1)
						v162 = l0 + v161
						v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						v167 = v165 & v161
						if v167 != 0 {
							v168 = v162
						} else {
							v168 = l0 + int32(4)
						}
						if v165 == int32(1) {
							v171 = int32(4)
							v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
							if v173&int32(254) == int32(2) {
								v182 = v171
							} else {
								v182 = base.B2i32(v173 == int32(18)) << (uint(v171) % 32)
							}
							if v173 == int32(1) {
								v185 = v171
							} else {
								v185 = v182
							}
							v186 = F_SB_IMatchText(m, v168, v185, v160, v159, v10)
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int32(0)
							} else {
								return v186
							}
						} else {
							if v167 != 0 {
								v189 = int32(1)
								v193 = F_SB_IMatchText(m, v168, int32(base.Ui32(v165)>>(uint(v189)%32))-v189, v160, v159, v10)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									return v193
								}
							} else {
								v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v201 = F_SB_IMatchText(m, v168, int32(base.Ui32(v196)>>(uint(int32(2))%32))-int32(4), v160, v159, v10)
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int32(0)
								} else {
									return v201
								}
							}
						}
					} else {
						v31 = F_DirectFunctionCall1Coll(m, int32(1453), l2, l1)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_pg_detoast_datum_packed(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = int32(1)
								v36 = v33 + v35
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								v39 = v37 & v35
								if v37 == v35 {
									v42 = int32(4)
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
									if v44&int32(254) == int32(2) {
										v53 = v42
									} else {
										v53 = base.B2i32(v44 == int32(18)) << (uint(v42) % 32)
									}
									if v44 == int32(1) {
										v56 = v42
									} else {
										v56 = v53
									}
									v67 = v56
								} else {
									v57 = int32(1)
									if v39 != 0 {
										v67 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v71 = F_DirectFunctionCall1Coll(m, int32(1453), l2, l0)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = F_pg_detoast_datum_packed(m, v71)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v77 = int32(1)
										v78 = v73 + v77
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
										v81 = v79 & v77
										if v79 == v77 {
											v84 = int32(4)
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
											if v86&int32(254) == int32(2) {
												v95 = v84
											} else {
												v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
											}
											if v86 == int32(1) {
												v98 = v84
											} else {
												v98 = v95
											}
											v109 = v98
										} else {
											v99 = int32(1)
											if v81 != 0 {
												v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
												v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										if v39 != 0 {
											v110 = v36
										} else {
											v110 = v33 + int32(4)
										}
										if v81 != 0 {
											v111 = v78
										} else {
											v111 = v73 + int32(4)
										}
										v113 = *(*int32)(unsafe.Add(mBase, _consts[495]))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
										if v114 == int32(6) {
											v118 = F_UTF8_MatchText(m, v111, v109, v110, v67, int32(0))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												return v118
											}
										} else {
											v122 = F_MB_MatchText(m, v111, v109, v110, v67, int32(0))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												return v122
											}
										}
									}
								}
							}
						}
					}
				} else {
					v31 = F_DirectFunctionCall1Coll(m, int32(1453), l2, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_pg_detoast_datum_packed(m, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(1)
							v36 = v33 + v35
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
							v39 = v37 & v35
							if v37 == v35 {
								v42 = int32(4)
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
								if v44&int32(254) == int32(2) {
									v53 = v42
								} else {
									v53 = base.B2i32(v44 == int32(18)) << (uint(v42) % 32)
								}
								if v44 == int32(1) {
									v56 = v42
								} else {
									v56 = v53
								}
								v67 = v56
							} else {
								v57 = int32(1)
								if v39 != 0 {
									v67 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v71 = F_DirectFunctionCall1Coll(m, int32(1453), l2, l0)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = F_pg_detoast_datum_packed(m, v71)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v77 = int32(1)
									v78 = v73 + v77
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
									v81 = v79 & v77
									if v79 == v77 {
										v84 = int32(4)
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
										if v86&int32(254) == int32(2) {
											v95 = v84
										} else {
											v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
										}
										if v86 == int32(1) {
											v98 = v84
										} else {
											v98 = v95
										}
										v109 = v98
									} else {
										v99 = int32(1)
										if v81 != 0 {
											v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
											v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									if v39 != 0 {
										v110 = v36
									} else {
										v110 = v33 + int32(4)
									}
									if v81 != 0 {
										v111 = v78
									} else {
										v111 = v73 + int32(4)
									}
									v113 = *(*int32)(unsafe.Add(mBase, _consts[495]))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
									if v114 == int32(6) {
										v118 = F_UTF8_MatchText(m, v111, v109, v110, v67, int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											return v118
										}
									} else {
										v122 = F_MB_MatchText(m, v111, v109, v110, v67, int32(0))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											return v122
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v207 = m.ExcPending
		if v207 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v210 = m.ExcPending
			if v210 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(562403), int32(0))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(597242), int32(0))
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519598), int32(194), int32(414639))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
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
	}
}
func F_generic_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 float64
	_ = v71
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v88 float64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 float32
	_ = v94
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v123 float64
	_ = v123
	v8 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v20 = F_get_restriction_variable(m, l0, l3, l4, v12-int32(-64), v12+int32(60), v12+int32(59))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return float64(0)
	} else {
		if v20 == int32(0) {
			v123 = l5
			m.G0 = v12 + int32(96)
			return v123
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			if v27 != int32(7) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
				v105 = v30
				v106 = l5
				if v105 != 0 {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
					m.T0[v109].(func(*base.Module, int32))(m, v105)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return float64(0)
					} else {
						v112 = float64(0)
						if base.F64_lt(v106, v112) != 0 {
							v123 = v112
						} else {
							if base.F64_gt(v106, float64(1)) == int32(0) {
								v123 = v106
							} else {
								v123 = float64(1)
							}
						}
						m.G0 = v12 + int32(96)
						return v123
					}
				} else {
					v112 = float64(0)
					if base.F64_lt(v106, v112) != 0 {
						v123 = v112
					} else {
						if base.F64_gt(v106, float64(1)) == int32(0) {
							v123 = v106
						} else {
							v123 = float64(1)
						}
					}
					m.G0 = v12 + int32(96)
					return v123
				}
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
				if v31 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
					if v34 == int32(0) {
						v123 = v8
						m.G0 = v12 + int32(96)
						return v123
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
						m.T0[v37].(func(*base.Module, int32))(m, v34)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return float64(0)
						} else {
							v123 = v8
							m.G0 = v12 + int32(96)
							return v123
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v41 = F_get_opcode(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return float64(0)
					} else {
						F_fmgr_info(m, v41, v12+int32(28))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return float64(0)
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+59)))
							v54 = F_mcv_selectivity(m, v12-int32(-64), v12+int32(28), l2, v40, v51, v12+int32(16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return float64(0)
							} else {
								v62 = F_histogram_selectivity(m, v12-int32(-64), v12+int32(28), l2, v40, v51, v12+int32(12))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return float64(0)
								} else {
									if base.F64_lt(v62, float64(0)) != 0 {
										v78 = l5
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										if int32(99) < v66 {
											v78 = v62
										} else {
											v71 = base.F64_div(base.F64_convert_i32_s(v66), float64(100))
											v78 = base.F64_add(base.F64_mul(v62, v71), base.F64_mul(l5, base.F64_sub(float64(1), v71)))
										}
									}
									v80 = float64(0.0001)
									if base.F64_lt(v78, v80) != 0 {
										v88 = v80
									} else {
										if base.F64_gt(v78, float64(0.9999)) == int32(0) {
											v88 = v78
										} else {
											v88 = float64(0.9999)
										}
									}
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
									if v90 != 0 {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
										v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
										v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
										v98 = base.F64_promote_f32(v94)
									} else {
										v98 = float64(0)
									}
									v100 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									v105 = v90
									v106 = base.F64_add(v54, base.F64_mul(v88, base.F64_sub(base.F64_sub(float64(1), v98), v100)))
									if v105 != 0 {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
										m.T0[v109].(func(*base.Module, int32))(m, v105)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return float64(0)
										} else {
											v112 = float64(0)
											if base.F64_lt(v106, v112) != 0 {
												v123 = v112
											} else {
												if base.F64_gt(v106, float64(1)) == int32(0) {
													v123 = v106
												} else {
													v123 = float64(1)
												}
											}
											m.G0 = v12 + int32(96)
											return v123
										}
									} else {
										v112 = float64(0)
										if base.F64_lt(v106, v112) != 0 {
											v123 = v112
										} else {
											if base.F64_gt(v106, float64(1)) == int32(0) {
												v123 = v106
											} else {
												v123 = float64(1)
											}
										}
										m.G0 = v12 + int32(96)
										return v123
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
