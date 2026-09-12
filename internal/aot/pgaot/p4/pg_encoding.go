package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PG_encoding_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v4) <= base.Ui32(int32(41)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(3))%32))+uint32(_consts[494])))
		v14 = v13
	} else {
		v14 = int32(757461)
	}
	v15 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return v15
	}
}
func F_pg_do_encoding_conversion_buf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v12 = int32(base.Ui32(l6-int32(1)) >> (uint(int32(2)) % 32))
	if base.Ui32(l4) < base.Ui32(v12) {
		v14 = l4
	} else {
		v14 = v12
	}
	v15 = F_OidFunctionCall6Coll(m, l0, l1, l2, l3, l5, v14, l7)
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return v15
	}
}
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[1462])))
		v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l1, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = v13
			return v17
		}
	} else {
		v17 = int32(1)
		return v17
	}
}
func F_pg_get_encoding_from_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v390 int32
	_ = v390
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
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
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	goto L6
L2:
	;
	v165 = l0
	goto L3
L3:
	;
	v169 = v165
	v170 = int32(544775)
	goto L42
L4:
	;
	v165 = v105
	goto L3
L5:
	;
	m.G0 = v25 + int32(48)
	goto L4
L6:
	;
	goto L9
L9:
	;
	goto L10
L10:
	;
	goto L25
L23:
	;
	if v100 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[1463]))
	goto L23
L28:
	;
	v105 = v100 + int32(8)
	goto L30
L29:
	;
	v105 = int32(544775)
	goto L30
L30:
	;
	goto L5
L40:
	;
	m.G0 = v11 + int32(16)
	return v617
L41:
	;
	if v207 == int32(0) {
		v617 = v3
		goto L40
	} else {
		goto L54
	}
L42:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v173 == v174 {
		v196 = v173
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v207 = int32(0)
	goto L41
L44:
	;
	v198 = int32(1)
	if v196 != 0 {
		v169 = v169 + v198
		v170 = v170 + v198
		goto L42
	} else {
		goto L53
	}
L45:
	;
	if base.Ui32((v173-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v184 = v173 | int32(32)
	goto L48
L47:
	;
	v184 = v173
	goto L48
L48:
	;
	if base.Ui32((v174-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v193 = v174 | int32(32)
	goto L51
L50:
	;
	v193 = v174
	goto L51
L51:
	;
	if v184 == v193 {
		v196 = v184
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v207 = v184 - v193
	goto L41
L53:
	;
	goto L43
L54:
	;
	v213 = v165
	v214 = int32(510090)
	goto L56
L55:
	;
	if v251 == int32(0) {
		v617 = v3
		goto L40
	} else {
		goto L68
	}
L56:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v217 == v218 {
		v240 = v217
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v251 = int32(0)
	goto L55
L58:
	;
	v242 = int32(1)
	if v240 != 0 {
		v213 = v213 + v242
		v214 = v214 + v242
		goto L56
	} else {
		goto L67
	}
L59:
	;
	if base.Ui32((v217-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v228 = v217 | int32(32)
	goto L62
L61:
	;
	v228 = v217
	goto L62
L62:
	;
	if base.Ui32((v218-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v237 = v218 | int32(32)
	goto L65
L64:
	;
	v237 = v218
	goto L65
L65:
	;
	if v228 == v237 {
		v240 = v228
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v251 = v228 - v237
	goto L55
L67:
	;
	goto L57
L68:
	;
	v254 = int32(-1)
	v256 = int32(0)
	v260 = m.G0
	v262 = v260 - int32(32)
	m.G0 = v262
	v267 = v256
	goto L72
L69:
	;
	if v390 == int32(0) {
		v617 = v254
		goto L40
	} else {
		goto L97
	}
L70:
	;
	m.G0 = v262 + int32(32)
	goto L69
L71:
	;
	v390 = int32(0)
	goto L70
L72:
	;
	goto L75
L73:
	;
	v295 = F___loc_is_allocated(m, v256)
	mBase = m.M
	if v295 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262+int32(8)+v267<<(uint(int32(2))%32)))) = v286
	if v286 == int32(-1) {
		goto L71
	} else {
		goto L81
	}
L75:
	;
	if int32(1)<<(uint(v267)%32)&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v285 = v165
	goto L80
L79:
	;
	v285 = int32(757461)
	goto L80
L80:
	;
	v286 = F___get_locale(m, v267, v285)
	mBase = m.M
	goto L74
L81:
	;
	v292 = v267 + int32(1)
	if v292 != int32(6) {
		v267 = v292
		goto L72
	} else {
		goto L82
	}
L82:
	;
	goto L73
L83:
	;
	v298 = int32(4097064)
	v303 = F_memcmp(m, v262+int32(8), v298, int32(24))
	mBase = m.M
	if v303 == int32(0) {
		v390 = v298
		goto L70
	} else {
		goto L86
	}
L84:
	;
	v371 = v256
	goto L85
L85:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v262)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v371))) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v262)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v371)+16)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v262)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v371)+8)) = v379
	v390 = v371
	goto L70
L86:
	;
	v306 = int32(4097088)
	v311 = F_memcmp(m, v262+int32(8), v306, int32(24))
	mBase = m.M
	if v311 == int32(0) {
		v390 = v306
		goto L70
	} else {
		goto L87
	}
L87:
	;
	v314 = int32(0)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	if v316 == v314 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v322 = v314
	goto L91
L89:
	;
	goto L90
L90:
	;
	v349 = int32(4680276)
	v354 = F_memcmp(m, v262+int32(8), v349, int32(24))
	mBase = m.M
	if v354 == int32(0) {
		v390 = v349
		goto L70
	} else {
		goto L94
	}
L91:
	;
	v330 = F___get_locale(m, v322, int32(757461))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v322<<(uint(int32(2))%32))+uint32(_consts[1318]))) = v330
	v333 = v322 + int32(1)
	if v333 != int32(6) {
		v322 = v333
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1317])) = uint8(v337)
	v341 = *(*int32)(unsafe.Add(mBase, _consts[1318]))
	*(*int32)(unsafe.Add(mBase, _consts[1319])) = v341
	goto L90
L93:
	;
	goto L92
L94:
	;
	v357 = int32(4680300)
	v362 = F_memcmp(m, v262+int32(8), v357, int32(24))
	mBase = m.M
	if v362 == int32(0) {
		v390 = v357
		goto L70
	} else {
		goto L95
	}
L95:
	;
	v366 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v366 == int32(0) {
		goto L71
	} else {
		goto L96
	}
L96:
	;
	v371 = v366
	goto L85
L97:
	;
	goto L99
L98:
	;
	if v407 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L99:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	if v406 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v407 = int32(556391)
	goto L104
L103:
	;
	v407 = int32(534714)
	goto L104
L104:
	;
	goto L98
L126:
	;
	v459 = F___loc_is_allocated(m, v390)
	mBase = m.M
	if v459 != 0 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v463 = F_strlen(m, v407)
	mBase = m.M
	v465 = v463 + int32(1)
	v466 = F_emscripten_builtin_malloc(m, v465)
	mBase = m.M
	if v466 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v617 = v254
	goto L40
L130:
	;
	F_emscripten_builtin_free(m, v390)
	mBase = m.M
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L129
L133:
	;
	v472 = F___loc_is_allocated(m, v390)
	mBase = m.M
	if v472 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v471 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v470 = F___memcpy(m, v466, v407, v465)
	mBase = m.M
	v471 = v470
	goto L133
L137:
	;
	if v471 == int32(0) {
		v617 = v254
		goto L40
	} else {
		goto L141
	}
L138:
	;
	F_emscripten_builtin_free(m, v390)
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
	v479 = v471
	v480 = int32(526975)
	goto L144
L142:
	;
	if l1 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L143:
	;
	if v517 != 0 {
		goto L156
	} else {
		goto L157
	}
L144:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v483 == v484 {
		v506 = v483
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v517 = int32(0)
	goto L143
L146:
	;
	v508 = int32(1)
	if v506 != 0 {
		v479 = v479 + v508
		v480 = v480 + v508
		goto L144
	} else {
		goto L155
	}
L147:
	;
	if base.Ui32((v483-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v494 = v483 | int32(32)
	goto L150
L149:
	;
	v494 = v483
	goto L150
L150:
	;
	if base.Ui32((v484-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v503 = v484 | int32(32)
	goto L153
L152:
	;
	v503 = v484
	goto L153
L153:
	;
	if v494 == v503 {
		v506 = v494
		goto L146
	} else {
		goto L154
	}
L154:
	;
	v517 = v494 - v503
	goto L143
L155:
	;
	goto L145
L156:
	;
	v521 = int32(0)
	goto L159
L157:
	;
	v588 = int32(2144960)
	goto L158
L158:
	;
	F_emscripten_builtin_free(m, v471)
	mBase = m.M
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	v617 = v590
	goto L40
L159:
	;
	v528 = v521 + int32(1)
	v530 = v528 << (uint(int32(3)) % 32)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v530)+uint32(_consts[1464])))
	if v533 == int32(0) {
		goto L142
	} else {
		goto L161
	}
L160:
	;
	v588 = v530 + int32(2144960)
	goto L158
L161:
	;
	v538 = v471
	v539 = v533
	goto L163
L162:
	;
	if v576 != 0 {
		v521 = v528
		goto L159
	} else {
		goto L175
	}
L163:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	if v542 == v543 {
		v565 = v542
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v576 = int32(0)
	goto L162
L165:
	;
	v567 = int32(1)
	if v565 != 0 {
		v538 = v538 + v567
		v539 = v539 + v567
		goto L163
	} else {
		goto L174
	}
L166:
	;
	if base.Ui32((v542-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v553 = v542 | int32(32)
	goto L169
L168:
	;
	v553 = v542
	goto L169
L169:
	;
	if base.Ui32((v543-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v562 = v543 | int32(32)
	goto L172
L171:
	;
	v562 = v543
	goto L172
L172:
	;
	if v553 == v562 {
		v565 = v553
		goto L165
	} else {
		goto L173
	}
L173:
	;
	v576 = v553 - v562
	goto L162
L174:
	;
	goto L164
L175:
	;
	goto L160
L176:
	;
	F_emscripten_builtin_free(m, v471)
	mBase = m.M
	v617 = v254
	goto L40
L177:
	;
	v595 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	return int32(0)
L179:
	;
	if v595 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v165
	F_errmsg(m, int32(699838), v11)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(499251), int32(377), int32(397807))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L178
	} else {
		goto L182
	}
L182:
	;
	goto L176
}
