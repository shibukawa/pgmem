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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v4) <= base.Ui32(int32(41)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(3))%32))+uint32(_c_F_PG_encoding_to_char[0])))
		v11 = v9
	} else {
		v11 = int32(_a_F_PG_encoding_to_char_0)
	}
	v12 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_verifymbchar[0])))
		v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, l1, l2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = v9
			return v14
		}
	} else {
		v14 = int32(1)
		return v14
	}
}
func F_pg_get_encoding_from_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = m.G0
	v24 = v22 - int32(48)
	m.G0 = v24
	goto L6
L2:
	;
	v158 = l0
	goto L3
L3:
	;
	v162 = v158
	v163 = int32(_a_F_pg_get_encoding_from_locale_0)
	goto L42
L4:
	;
	v158 = v100
	goto L3
L5:
	;
	m.G0 = v24 + int32(48)
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
	if v95 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_encoding_from_locale[0]))
	goto L23
L28:
	;
	v100 = v95 + int32(8)
	goto L30
L29:
	;
	v100 = int32(_a_F_pg_get_encoding_from_locale_0)
	goto L30
L30:
	;
	goto L5
L40:
	;
	m.G0 = v10 + int32(16)
	return v560
L41:
	;
	if v200 == int32(0) {
		v560 = v3
		goto L40
	} else {
		goto L54
	}
L42:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v166 == v167 {
		v189 = v166
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v200 = int32(0)
	goto L41
L44:
	;
	v191 = int32(1)
	if v189 != 0 {
		v162 = v162 + v191
		v163 = v163 + v191
		goto L42
	} else {
		goto L53
	}
L45:
	;
	if base.Ui32((v166-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v177 = v166 | int32(32)
	goto L48
L47:
	;
	v177 = v166
	goto L48
L48:
	;
	if base.Ui32((v167-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = v167 | int32(32)
	goto L51
L50:
	;
	v186 = v167
	goto L51
L51:
	;
	if v177 == v186 {
		v189 = v177
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v200 = v177 - v186
	goto L41
L53:
	;
	goto L43
L54:
	;
	v206 = v158
	v207 = int32(_a_F_pg_get_encoding_from_locale_1)
	goto L56
L55:
	;
	if v244 == int32(0) {
		v560 = v3
		goto L40
	} else {
		goto L68
	}
L56:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v210 == v211 {
		v233 = v210
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v244 = int32(0)
	goto L55
L58:
	;
	v235 = int32(1)
	if v233 != 0 {
		v206 = v206 + v235
		v207 = v207 + v235
		goto L56
	} else {
		goto L67
	}
L59:
	;
	if base.Ui32((v210-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v221 = v210 | int32(32)
	goto L62
L61:
	;
	v221 = v210
	goto L62
L62:
	;
	if base.Ui32((v211-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v230 = v211 | int32(32)
	goto L65
L64:
	;
	v230 = v211
	goto L65
L65:
	;
	if v221 == v230 {
		v233 = v221
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v244 = v221 - v230
	goto L55
L67:
	;
	goto L57
L68:
	;
	v247 = int32(-1)
	v249 = int32(0)
	v254 = m.G0
	v256 = v254 - int32(32)
	m.G0 = v256
	v261 = v249
	goto L72
L69:
	;
	if v384 == int32(0) {
		v560 = v247
		goto L40
	} else {
		goto L97
	}
L70:
	;
	m.G0 = v256 + int32(32)
	goto L69
L71:
	;
	v384 = int32(0)
	goto L70
L72:
	;
	v266 = v261 << (uint(int32(2)) % 32)
	v272 = int32(1) << (uint(v261) % 32) & int32(1)
	if v272|int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v291 = F___loc_is_allocated(m, v249)
	mBase = m.M
	if v291 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266+(v256+int32(8))))) = v283
	if v283 == int32(-1) {
		goto L71
	} else {
		goto L81
	}
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v249+v266)))
	v283 = v279
	goto L74
L76:
	;
	goto L77
L77:
	;
	if v272 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v281 = v158
	goto L80
L79:
	;
	v281 = int32(_a_F_pg_get_encoding_from_locale_2)
	goto L80
L80:
	;
	v282 = F___get_locale(m, v261, v281)
	mBase = m.M
	v283 = v282
	goto L74
L81:
	;
	v288 = v261 + int32(1)
	if v288 != int32(6) {
		v261 = v288
		goto L72
	} else {
		goto L82
	}
L82:
	;
	goto L73
L83:
	;
	v294 = int32(_a_F_pg_get_encoding_from_locale_3)
	v296 = v256 + int32(8)
	v299 = F_memcmp(m, v296, v294, int32(24))
	mBase = m.M
	if v299 == int32(0) {
		v384 = v294
		goto L70
	} else {
		goto L86
	}
L84:
	;
	v363 = v249
	goto L85
L85:
	;
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v256)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+16)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v256)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+8)) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v363))) = v372
	v384 = v363
	goto L70
L86:
	;
	v302 = int32(_a_F_pg_get_encoding_from_locale_4)
	v305 = F_memcmp(m, v296, v302, int32(24))
	mBase = m.M
	if v305 == int32(0) {
		v384 = v302
		goto L70
	} else {
		goto L87
	}
L87:
	;
	v308 = int32(0)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_encoding_from_locale[1])))
	if v310 == v308 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v316 = v308
	goto L91
L89:
	;
	goto L90
L90:
	;
	v343 = int32(_a_F_pg_get_encoding_from_locale_5)
	v345 = v256 + int32(8)
	v348 = F_memcmp(m, v345, v343, int32(24))
	mBase = m.M
	if v348 == int32(0) {
		v384 = v343
		goto L70
	} else {
		goto L94
	}
L91:
	;
	v323 = F___get_locale(m, v316, int32(_a_F_pg_get_encoding_from_locale_2))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(int32(2))%32))+uint32(_c_F_pg_get_encoding_from_locale[2]))) = v323
	v326 = v316 + int32(1)
	if v326 != int32(6) {
		v316 = v326
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_encoding_from_locale[1])) = uint8(v330)
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_encoding_from_locale[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_encoding_from_locale[3])) = v334
	goto L90
L93:
	;
	goto L92
L94:
	;
	v351 = int32(_a_F_pg_get_encoding_from_locale_6)
	v354 = F_memcmp(m, v345, v351, int32(24))
	mBase = m.M
	if v354 == int32(0) {
		v384 = v351
		goto L70
	} else {
		goto L95
	}
L95:
	;
	v358 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v358 == int32(0) {
		goto L71
	} else {
		goto L96
	}
L96:
	;
	v363 = v358
	goto L85
L97:
	;
	goto L99
L98:
	;
	if v402 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L99:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v401 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v402 = int32(_a_F_pg_get_encoding_from_locale_7)
	goto L104
L103:
	;
	v402 = int32(_a_F_pg_get_encoding_from_locale_8)
	goto L104
L104:
	;
	goto L98
L124:
	;
	v457 = F___loc_is_allocated(m, v384)
	mBase = m.M
	if v457 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	v461 = F_strlen(m, v402)
	mBase = m.M
	v463 = v461 + int32(1)
	v464 = F_emscripten_builtin_malloc(m, v463)
	mBase = m.M
	if v464 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v560 = v247
	goto L40
L128:
	;
	F_emscripten_builtin_free(m, v384)
	mBase = m.M
	goto L130
L129:
	;
	goto L130
L130:
	;
	goto L127
L131:
	;
	v470 = F___loc_is_allocated(m, v384)
	mBase = m.M
	if v470 != 0 {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v469 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v468 = F___memcpy(m, v464, v402, v463)
	mBase = m.M
	v469 = v468
	goto L131
L135:
	;
	if v469 == int32(0) {
		v560 = v247
		goto L40
	} else {
		goto L139
	}
L136:
	;
	F_emscripten_builtin_free(m, v384)
	mBase = m.M
	goto L138
L137:
	;
	goto L138
L138:
	;
	goto L135
L139:
	;
	v477 = int32(0)
	goto L141
L140:
	;
	if l1 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L141:
	;
	v483 = v477 << (uint(int32(3)) % 32)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_pg_get_encoding_from_locale[4])))
	v487 = v469
	v488 = v484
	goto L144
L142:
	;
	F_emscripten_builtin_free(m, v469)
	mBase = m.M
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_pg_get_encoding_from_locale[5])))
	v560 = v533
	goto L40
L143:
	;
	if v525 != 0 {
		goto L156
	} else {
		goto L157
	}
L144:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v491 == v492 {
		v514 = v491
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v525 = int32(0)
	goto L143
L146:
	;
	v516 = int32(1)
	if v514 != 0 {
		v487 = v487 + v516
		v488 = v488 + v516
		goto L144
	} else {
		goto L155
	}
L147:
	;
	if base.Ui32((v491-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v502 = v491 | int32(32)
	goto L150
L149:
	;
	v502 = v491
	goto L150
L150:
	;
	if base.Ui32((v492-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v511 = v492 | int32(32)
	goto L153
L152:
	;
	v511 = v492
	goto L153
L153:
	;
	if v502 == v511 {
		v514 = v502
		goto L146
	} else {
		goto L154
	}
L154:
	;
	v525 = v502 - v511
	goto L143
L155:
	;
	goto L145
L156:
	;
	v527 = v477 + int32(1)
	if v527 != int32(110) {
		v477 = v527
		goto L141
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	goto L142
L159:
	;
	goto L140
L160:
	;
	F_emscripten_builtin_free(m, v469)
	mBase = m.M
	v560 = v247
	goto L40
L161:
	;
	v538 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	return int32(0)
L163:
	;
	if v538 == int32(0) {
		goto L160
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v158
	F_errmsg(m, int32(_a_F_pg_get_encoding_from_locale_9), v10)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L162
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_pg_get_encoding_from_locale_10), int32(377), int32(_a_F_pg_get_encoding_from_locale_11))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L162
	} else {
		goto L166
	}
L166:
	;
	goto L160
}
