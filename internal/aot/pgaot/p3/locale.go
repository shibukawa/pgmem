package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v223 int64
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v268 int64
	_ = v268
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v372 int64
	_ = v372
	var v375 int64
	_ = v375
	var v378 int64
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v423 int64
	_ = v423
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = l1
	goto L3
L1:
	;
	m.G0 = v9 + int32(48)
	return v564
L2:
	;
	if base.B2i32(v14 == int32(0)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	if int32(0) < v14 {
		v12 = v12 + int32(1)
		goto L3
	} else {
		goto L5
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	v25 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	return int32(0)
L10:
	;
	if v25 == int32(0) {
		v564 = v4
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errmsg(m, int32(133392), v9+int32(32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(498951), int32(312), int32(397552))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v564 = v4
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v47 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(48)
	m.G0 = v56
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v188 = v47
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v188 == int32(0) {
		v564 = v4
		goto L1
	} else {
		goto L54
	}
L19:
	;
	m.G0 = v56 + int32(48)
	goto L18
L20:
	;
	if l0 == int32(6) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v147 = int32(0)
	v148 = int32(4680576)
	v153 = v47
	goto L45
L22:
	;
	goto L21
L23:
	;
	goto L24
L24:
	;
	goto L39
L37:
	;
	if v131 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[989])))
	goto L37
L42:
	;
	v136 = v131 + int32(8)
	goto L44
L43:
	;
	v136 = int32(544456)
	goto L44
L44:
	;
	v188 = v136
	goto L19
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[989]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147<<(uint(int32(2))%32))+uint32(_consts[989])))
	if v161 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v179)
	if v174 != int32(6) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v165 = v161 + int32(8)
	goto L49
L48:
	;
	v165 = int32(544456)
	goto L49
L49:
	;
	v166 = F_strlen(m, v165)
	mBase = m.M
	v167 = F___memcpy(m, v148, v165, v166)
	mBase = m.M
	v168 = v148 + v166
	v169 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v169)
	v171 = int32(1)
	v174 = v153 + base.B2i32(v161 == v156)
	v176 = v147 + v171
	if v176 != int32(6) {
		v147 = v176
		v148 = v168 + v171
		v153 = v174
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v184 = int32(4680576)
	goto L53
L52:
	;
	v184 = v165
	goto L53
L53:
	;
	v188 = v184
	goto L19
L54:
	;
	v198 = F_pstrdup(m, v188)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v200 = int32(0)
	v206 = m.G0
	v208 = v206 - int32(48)
	m.G0 = v208
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v340 = v200
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if l2 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L57:
	;
	m.G0 = v208 + int32(48)
	goto L56
L58:
	;
	if l0 == int32(6) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v299 = int32(0)
	v300 = int32(4680576)
	v305 = v200
	goto L83
L60:
	;
	if l1 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if l1 != 0 {
		goto L76
	} else {
		goto L77
	}
L63:
	;
	v217 = *(*int64)(unsafe.Add(mBase, _consts[990]))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+16)) = v217
	v220 = *(*int64)(unsafe.Add(mBase, _consts[991]))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+8)) = v220
	v223 = *(*int64)(unsafe.Add(mBase, _consts[992]))
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = v223
	v226 = int32(0)
	v227 = l1
	goto L65
L64:
	;
	v340 = int32(0)
	goto L57
L65:
	;
	v235 = F___strchrnul(m, v227, int32(59))
	mBase = m.M
	v236 = v235 - v227
	if v236 <= int32(23) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v208)+24))
	*(*int64)(unsafe.Add(mBase, _consts[989])) = v262
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v208)+40))
	*(*int64)(unsafe.Add(mBase, _consts[993])) = v265
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v208)+32))
	*(*int64)(unsafe.Add(mBase, _consts[994])) = v268
	goto L59
L67:
	;
	v239 = F___memcpy(m, v208, v227, v236)
	mBase = m.M
	v241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v236))) = uint8(v241)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v245 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v247 = v227
	goto L69
L69:
	;
	v248 = F___get_locale(m, v226, v208)
	mBase = m.M
	if v248 == int32(-1) {
		goto L64
	} else {
		goto L73
	}
L70:
	;
	v246 = v235 + int32(1)
	goto L72
L71:
	;
	v246 = v227
	goto L72
L72:
	;
	v247 = v246
	goto L69
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208+int32(24)+v226<<(uint(int32(2))%32)))) = v248
	v258 = v226 + int32(1)
	if v258 != int32(6) {
		v226 = v258
		v227 = v247
		goto L65
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	if v284 != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v271 = F___get_locale(m, l0, l1)
	mBase = m.M
	if v271 == int32(-1) {
		v340 = v200
		goto L57
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[989])))
	v284 = v283
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[989]))) = v271
	v284 = v271
	goto L75
L80:
	;
	v288 = v284 + int32(8)
	goto L82
L81:
	;
	v288 = int32(544456)
	goto L82
L82:
	;
	v340 = v288
	goto L57
L83:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[989]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v299<<(uint(int32(2))%32))+uint32(_consts[989])))
	if v313 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v331)
	if v326 != int32(6) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v317 = v313 + int32(8)
	goto L87
L86:
	;
	v317 = int32(544456)
	goto L87
L87:
	;
	v318 = F_strlen(m, v317)
	mBase = m.M
	v319 = F___memcpy(m, v300, v317, v318)
	mBase = m.M
	v320 = v300 + v318
	v321 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v321)
	v323 = int32(1)
	v326 = v305 + base.B2i32(v313 == v308)
	v328 = v299 + v323
	if v328 != int32(6) {
		v299 = v328
		v300 = v320 + v323
		v305 = v326
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	v336 = int32(4680576)
	goto L91
L90:
	;
	v336 = v317
	goto L91
L91:
	;
	v340 = v336
	goto L57
L92:
	;
	v355 = int32(0)
	v361 = m.G0
	v363 = v361 - int32(48)
	m.G0 = v363
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v495 = v355
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if v340 == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v352 = F_pstrdup(m, v340)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v352
	goto L92
L96:
	;
	v521 = base.B2i32(v340 != int32(0))
	F_pfree(m, v198)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L9
	} else {
		goto L138
	}
L97:
	;
	if v495 != 0 {
		goto L96
	} else {
		goto L133
	}
L98:
	;
	m.G0 = v363 + int32(48)
	goto L97
L99:
	;
	if l0 == int32(6) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v454 = int32(0)
	v455 = int32(4680576)
	v460 = v355
	goto L124
L101:
	;
	if v198 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v198 != 0 {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	v372 = *(*int64)(unsafe.Add(mBase, _consts[990]))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+16)) = v372
	v375 = *(*int64)(unsafe.Add(mBase, _consts[991]))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+8)) = v375
	v378 = *(*int64)(unsafe.Add(mBase, _consts[992]))
	*(*int64)(unsafe.Add(mBase, uint32(v363))) = v378
	v381 = int32(0)
	v382 = v198
	goto L106
L105:
	;
	v495 = int32(0)
	goto L98
L106:
	;
	v390 = F___strchrnul(m, v382, int32(59))
	mBase = m.M
	v391 = v390 - v382
	if v391 <= int32(23) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v363)+24))
	*(*int64)(unsafe.Add(mBase, _consts[989])) = v417
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v363)+40))
	*(*int64)(unsafe.Add(mBase, _consts[993])) = v420
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v363)+32))
	*(*int64)(unsafe.Add(mBase, _consts[994])) = v423
	goto L100
L108:
	;
	v394 = F___memcpy(m, v363, v382, v391)
	mBase = m.M
	v396 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v363+v391))) = uint8(v396)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	if v400 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v402 = v382
	goto L110
L110:
	;
	v403 = F___get_locale(m, v381, v363)
	mBase = m.M
	if v403 == int32(-1) {
		goto L105
	} else {
		goto L114
	}
L111:
	;
	v401 = v390 + int32(1)
	goto L113
L112:
	;
	v401 = v382
	goto L113
L113:
	;
	v402 = v401
	goto L110
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363+int32(24)+v381<<(uint(int32(2))%32)))) = v403
	v413 = v381 + int32(1)
	if v413 != int32(6) {
		v381 = v413
		v382 = v402
		goto L106
	} else {
		goto L115
	}
L115:
	;
	goto L107
L116:
	;
	if v439 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v426 = F___get_locale(m, l0, v198)
	mBase = m.M
	if v426 == int32(-1) {
		v495 = v355
		goto L98
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[989])))
	v439 = v438
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[989]))) = v426
	v439 = v426
	goto L116
L121:
	;
	v443 = v439 + int32(8)
	goto L123
L122:
	;
	v443 = int32(544456)
	goto L123
L123:
	;
	v495 = v443
	goto L98
L124:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _consts[989]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v454<<(uint(int32(2))%32))+uint32(_consts[989])))
	if v468 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v475))) = uint8(v486)
	if v481 != int32(6) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v472 = v468 + int32(8)
	goto L128
L127:
	;
	v472 = int32(544456)
	goto L128
L128:
	;
	v473 = F_strlen(m, v472)
	mBase = m.M
	v474 = F___memcpy(m, v455, v472, v473)
	mBase = m.M
	v475 = v455 + v473
	v476 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v475))) = uint8(v476)
	v478 = int32(1)
	v481 = v460 + base.B2i32(v468 == v463)
	v483 = v454 + v478
	if v483 != int32(6) {
		v454 = v483
		v455 = v475 + v478
		v460 = v481
		goto L124
	} else {
		goto L129
	}
L129:
	;
	goto L125
L130:
	;
	v491 = int32(4680576)
	goto L132
L131:
	;
	v491 = v472
	goto L132
L132:
	;
	v495 = v491
	goto L98
L133:
	;
	v505 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	if v505 == int32(0) {
		goto L96
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v198
	F_errmsg_internal(m, int32(721261), v9+int32(16))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(498951), int32(335), int32(397552))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	goto L96
L138:
	;
	if l2 == int32(0) {
		v564 = v521
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v526 == int32(0) {
		v564 = v521
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v530 = v526
	goto L142
L141:
	;
	if v532 == int32(0) {
		v564 = v521
		goto L1
	} else {
		goto L145
	}
L142:
	;
	v532 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530))))
	if int32(0) < v532 {
		v530 = v530 + int32(1)
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	goto L143
L145:
	;
	v539 = int32(0)
	v542 = F_errstart(m, int32(19), v539)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	if v542 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L9
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_pfree(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L9
	} else {
		goto L153
	}
L150:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v547
	F_errmsg(m, int32(133392), v9)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(498951), int32(344), int32(397552))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v564 = v539
	goto L1
}
func F_check_locale_time(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_check_locale(m, int32(2), v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
