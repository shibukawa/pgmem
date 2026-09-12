package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommandIsReadOnly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v9-int32(2)) < base.Ui32(int32(5)) {
		v38 = v2
		m.G0 = v7 + int32(16)
		return v38 & int32(1)
	} else {
		if v9 == int32(1) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			if v16 != 0 {
				v38 = v2
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
				v38 = v17 ^ int32(1)
			}
			m.G0 = v7 + int32(16)
			return v38 & int32(1)
		} else {
			v22 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v38 = v2
					m.G0 = v7 + int32(16)
					return v38 & int32(1)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
					F_errmsg_internal(m, int32(504617), v7)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(511117), int32(116), int32(19860))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = v2
							m.G0 = v7 + int32(16)
							return v38 & int32(1)
						}
					}
				}
			}
		}
	}
}
func F_CreateCommandTag(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = l0
	goto L1
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	switch v21 - int32(67) {
	case 0:
		goto L15
	default:
		goto L14
	case 69:
		v404 = int32(4)
		goto L3
	case 70:
		goto L13
	case 71:
		v359 = int32(103)
		goto L12
	case 72:
		v362 = int32(191)
		goto L11
	case 73:
		goto L10
	case 74, 77:
		goto L5
	case 78:
		goto L111
	case 79:
		goto L85
	case 83:
		goto L22
	case 84:
		goto L84
	case 85:
		goto L82
	case 88:
		goto L81
	case 89:
		goto L80
	case 90:
		goto L91
	case 91:
		goto L55
	case 92:
		goto L54
	case 93:
		goto L110
	case 95:
		goto L109
	case 96:
		goto L108
	case 97:
		goto L107
	case 98:
		goto L86
	case 99:
		goto L106
	case 100, 101:
		goto L105
	case 102:
		goto L104
	case 103:
		goto L103
	case 104:
		goto L102
	case 105:
		goto L101
	case 106:
		goto L97
	case 107:
		goto L100
	case 108:
		goto L99
	case 109:
		goto L98
	case 110:
		goto L96
	case 111:
		goto L30
	case 112:
		goto L29
	case 113:
		goto L28
	case 114:
		goto L51
	case 115:
		goto L50
	case 116:
		goto L49
	case 117:
		goto L48
	case 118:
		goto L47
	case 119, 120:
		goto L46
	case 121:
		goto L45
	case 122:
		goto L72
	case 123:
		goto L71
	case 124:
		goto L79
	case 125:
		goto L112
	case 126:
		goto L36
	case 128:
		goto L35
	case 129:
		goto L34
	case 130:
		goto L95
	case 131:
		goto L94
	case 132:
		goto L93
	case 133:
		goto L92
	case 134:
		v396 = int32(102)
		goto L4
	case 135:
		goto L114
	case 136:
		goto L113
	case 137:
		goto L74
	case 138:
		goto L19
	case 140:
		goto L18
	case 141:
		goto L75
	case 143:
		goto L83
	case 144:
		goto L70
	case 146:
		goto L62
	case 148:
		goto L90
	case 149:
		goto L89
	case 150:
		goto L88
	case 151:
		goto L87
	case 152:
		goto L33
	case 153, 162:
		goto L77
	case 154:
		goto L73
	case 155:
		goto L66
	case 156:
		goto L65
	case 157:
		goto L64
	case 158:
		goto L115
	case 159, 160, 161:
		goto L78
	case 163:
		goto L76
	case 164:
		goto L63
	case 165:
		goto L69
	case 166, 167, 168:
		goto L68
	case 169:
		goto L67
	case 170:
		goto L56
	case 171:
		goto L61
	case 172:
		goto L60
	case 174:
		goto L59
	case 175:
		goto L58
	case 176:
		goto L57
	case 177:
		goto L40
	case 178:
		goto L53
	case 179:
		goto L42
	case 180:
		goto L41
	case 181:
		goto L39
	case 182:
		goto L38
	case 183:
		goto L37
	case 184:
		goto L52
	case 185:
		goto L21
	case 186:
		goto L20
	case 187:
		goto L17
	case 188:
		goto L44
	case 189:
		goto L43
	case 190:
		goto L32
	case 191:
		goto L31
	case 194:
		goto L27
	case 195:
		goto L26
	case 196:
		goto L25
	case 197:
		goto L24
	case 198:
		goto L23
	case 263:
		goto L16
	}
L3:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v11+v404)))
	v11 = v406
	goto L1
L4:
	;
	m.G0 = v9 + int32(48)
	return v396
L5:
	;
	v396 = int32(179)
	goto L4
L6:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v315<<(uint(int32(2))%32))+uint32(_consts[909])))
	v396 = v388
	goto L4
L7:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v202<<(uint(int32(2))%32))+uint32(_consts[910])))
	v396 = v383
	goto L4
L8:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(2))%32))+uint32(_consts[911])))
	v396 = v378
	goto L4
L9:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[912])))
	v396 = v373
	goto L4
L10:
	;
	v396 = int32(163)
	goto L4
L11:
	;
	v396 = v362
	goto L4
L12:
	;
	v396 = v359
	goto L4
L13:
	;
	v396 = int32(158)
	goto L4
L14:
	;
	v338 = int32(0)
	v341 = F_errstart(m, int32(19), v338)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L194
	} else {
		goto L210
	}
L15:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v299 == int32(6) {
		goto L199
	} else {
		goto L200
	}
L16:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v253 == int32(6) {
		goto L187
	} else {
		goto L188
	}
L17:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v251 != 0 {
		goto L184
	} else {
		goto L185
	}
L18:
	;
	v396 = int32(31)
	goto L4
L19:
	;
	v396 = int32(86)
	goto L4
L20:
	;
	v396 = int32(152)
	goto L4
L21:
	;
	v396 = int32(166)
	goto L4
L22:
	;
	v396 = int32(4)
	goto L4
L23:
	;
	v396 = int32(140)
	goto L4
L24:
	;
	v396 = int32(32)
	goto L4
L25:
	;
	v396 = int32(87)
	goto L4
L26:
	;
	v396 = int32(24)
	goto L4
L27:
	;
	v396 = int32(79)
	goto L4
L28:
	;
	v396 = int32(58)
	goto L4
L29:
	;
	v396 = int32(22)
	goto L4
L30:
	;
	v396 = int32(77)
	goto L4
L31:
	;
	v396 = int32(36)
	goto L4
L32:
	;
	v396 = int32(37)
	goto L4
L33:
	;
	v396 = int32(19)
	goto L4
L34:
	;
	v396 = int32(21)
	goto L4
L35:
	;
	v396 = int32(76)
	goto L4
L36:
	;
	v396 = int32(75)
	goto L4
L37:
	;
	v396 = int32(60)
	goto L4
L38:
	;
	v396 = int32(63)
	goto L4
L39:
	;
	v396 = int32(170)
	goto L4
L40:
	;
	v396 = int32(48)
	goto L4
L41:
	;
	v396 = int32(186)
	goto L4
L42:
	;
	v396 = int32(161)
	goto L4
L43:
	;
	v396 = int32(168)
	goto L4
L44:
	;
	v396 = int32(129)
	goto L4
L45:
	;
	v396 = int32(133)
	goto L4
L46:
	;
	v396 = int32(25)
	goto L4
L47:
	;
	v396 = int32(80)
	goto L4
L48:
	;
	v396 = int32(72)
	goto L4
L49:
	;
	v396 = int32(10)
	goto L4
L50:
	;
	v396 = int32(66)
	goto L4
L51:
	;
	v396 = int32(96)
	goto L4
L52:
	;
	v396 = int32(95)
	goto L4
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v207) < base.Ui32(int32(4)) {
		goto L181
	} else {
		goto L182
	}
L54:
	;
	v396 = int32(187)
	goto L4
L55:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v202) < base.Ui32(int32(6)) {
		goto L7
	} else {
		goto L180
	}
L56:
	;
	v396 = int32(33)
	goto L4
L57:
	;
	v396 = int32(169)
	goto L4
L58:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	switch v192 - int32(23) {
	case 0:
		v396 = int32(73)
		goto L4
	default:
		goto L175
	case 18:
		goto L176
	}
L59:
	;
	v396 = int32(153)
	goto L4
L60:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v188 != 0 {
		goto L172
	} else {
		goto L173
	}
L61:
	;
	v396 = int32(52)
	goto L4
L62:
	;
	v396 = int32(47)
	goto L4
L63:
	;
	v396 = int32(160)
	goto L4
L64:
	;
	v396 = int32(190)
	goto L4
L65:
	;
	v396 = int32(159)
	goto L4
L66:
	;
	v396 = int32(165)
	goto L4
L67:
	;
	v396 = int32(116)
	goto L4
L68:
	;
	v396 = int32(7)
	goto L4
L69:
	;
	v396 = int32(64)
	goto L4
L70:
	;
	v396 = int32(109)
	goto L4
L71:
	;
	v396 = int32(29)
	goto L4
L72:
	;
	v396 = int32(84)
	goto L4
L73:
	;
	v396 = int32(82)
	goto L4
L74:
	;
	v396 = int32(71)
	goto L4
L75:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v170 != 0 {
		goto L169
	} else {
		goto L170
	}
L76:
	;
	v396 = int32(99)
	goto L4
L77:
	;
	v396 = int32(42)
	goto L4
L78:
	;
	v396 = int32(97)
	goto L4
L79:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v155 {
	case 0:
		goto L161
	case 1:
		v396 = int32(59)
		goto L4
	default:
		goto L160
	case 7:
		goto L162
	case 25:
		goto L168
	case 45:
		goto L163
	case 46:
		goto L165
	case 47:
		goto L166
	case 48:
		goto L164
	case 49:
		goto L167
	}
L80:
	;
	v396 = int32(8)
	goto L4
L81:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v151 != 0 {
		goto L157
	} else {
		goto L158
	}
L82:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v147 != 0 {
		goto L154
	} else {
		goto L155
	}
L83:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v139 - int32(19) {
	case 0:
		v396 = int32(14)
		goto L4
	default:
		goto L151
	case 10:
		goto L153
	case 15:
		goto L152
	}
L84:
	;
	v396 = int32(9)
	goto L4
L85:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v128 = v125 - int32(1)
	if base.Ui32(v128) <= base.Ui32(int32(50)) {
		goto L148
	} else {
		goto L149
	}
L86:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v116 = v113 - int32(1)
	if base.Ui32(v116) <= base.Ui32(int32(50)) {
		goto L144
	} else {
		goto L145
	}
L87:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v104 = v101 - int32(1)
	if base.Ui32(v104) <= base.Ui32(int32(50)) {
		goto L140
	} else {
		goto L141
	}
L88:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v92 = v89 - int32(1)
	if base.Ui32(v92) <= base.Ui32(int32(50)) {
		goto L136
	} else {
		goto L137
	}
L89:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v80 = v77 - int32(1)
	if base.Ui32(v80) <= base.Ui32(int32(50)) {
		goto L132
	} else {
		goto L133
	}
L90:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v61 == int32(6) {
		goto L124
	} else {
		goto L125
	}
L91:
	;
	v396 = int32(56)
	goto L4
L92:
	;
	v396 = int32(178)
	goto L4
L93:
	;
	v396 = int32(53)
	goto L4
L94:
	;
	v396 = int32(189)
	goto L4
L95:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if base.Ui32(v53) < base.Ui32(int32(52)) {
		goto L8
	} else {
		goto L123
	}
L96:
	;
	v396 = int32(157)
	goto L4
L97:
	;
	v396 = int32(69)
	goto L4
L98:
	;
	v396 = int32(150)
	goto L4
L99:
	;
	v396 = int32(43)
	goto L4
L100:
	;
	v396 = int32(98)
	goto L4
L101:
	;
	v396 = int32(30)
	goto L4
L102:
	;
	v396 = int32(85)
	goto L4
L103:
	;
	v396 = int32(12)
	goto L4
L104:
	;
	v396 = int32(68)
	goto L4
L105:
	;
	v396 = int32(11)
	goto L4
L106:
	;
	v396 = int32(67)
	goto L4
L107:
	;
	v396 = int32(35)
	goto L4
L108:
	;
	v396 = int32(142)
	goto L4
L109:
	;
	v396 = int32(90)
	goto L4
L110:
	;
	v396 = int32(88)
	goto L4
L111:
	;
	v396 = int32(83)
	goto L4
L112:
	;
	v396 = int32(65)
	goto L4
L113:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v34 != 0 {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v30 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v24) < base.Ui32(int32(10)) {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	v396 = int32(0)
	goto L4
L117:
	;
	v31 = int32(50)
	goto L119
L118:
	;
	v31 = int32(51)
	goto L119
L119:
	;
	v396 = v31
	goto L4
L120:
	;
	v35 = int32(164)
	goto L122
L121:
	;
	v35 = int32(154)
	goto L122
L122:
	;
	v396 = v35
	goto L4
L123:
	;
	v396 = int32(0)
	goto L4
L124:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v65 = v64
	goto L126
L125:
	;
	v65 = v61
	goto L126
L126:
	;
	v68 = v65 - int32(1)
	if base.Ui32(v68) <= base.Ui32(int32(50)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v396 = v76
	goto L4
L128:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68<<(uint(int32(2))%32))+uint32(_consts[913])))
	v76 = v75
	goto L130
L129:
	;
	v76 = int32(0)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v396 = v88
	goto L4
L132:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80<<(uint(int32(2))%32))+uint32(_consts[913])))
	v88 = v87
	goto L134
L133:
	;
	v88 = int32(0)
	goto L134
L134:
	;
	goto L131
L135:
	;
	v396 = v100
	goto L4
L136:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32))+uint32(_consts[913])))
	v100 = v99
	goto L138
L137:
	;
	v100 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	v396 = v112
	goto L4
L140:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104<<(uint(int32(2))%32))+uint32(_consts[913])))
	v112 = v111
	goto L142
L141:
	;
	v112 = int32(0)
	goto L142
L142:
	;
	goto L139
L143:
	;
	v396 = v124
	goto L4
L144:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(int32(2))%32))+uint32(_consts[913])))
	v124 = v123
	goto L146
L145:
	;
	v124 = int32(0)
	goto L146
L146:
	;
	goto L143
L147:
	;
	v396 = v136
	goto L4
L148:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128<<(uint(int32(2))%32))+uint32(_consts[913])))
	v136 = v135
	goto L150
L149:
	;
	v136 = int32(0)
	goto L150
L150:
	;
	goto L147
L151:
	;
	v396 = int32(0)
	goto L4
L152:
	;
	v396 = int32(26)
	goto L4
L153:
	;
	v396 = int32(23)
	goto L4
L154:
	;
	v148 = int32(155)
	goto L156
L155:
	;
	v148 = int32(173)
	goto L156
L156:
	;
	v396 = v148
	goto L4
L157:
	;
	v152 = int32(156)
	goto L159
L158:
	;
	v152 = int32(174)
	goto L159
L159:
	;
	v396 = v152
	goto L4
L160:
	;
	v396 = int32(0)
	goto L4
L161:
	;
	v396 = int32(58)
	goto L4
L162:
	;
	v396 = int32(61)
	goto L4
L163:
	;
	v396 = int32(91)
	goto L4
L164:
	;
	v396 = int32(94)
	goto L4
L165:
	;
	v396 = int32(92)
	goto L4
L166:
	;
	v396 = int32(93)
	goto L4
L167:
	;
	v396 = int32(97)
	goto L4
L168:
	;
	v396 = int32(74)
	goto L4
L169:
	;
	v171 = int32(78)
	goto L171
L170:
	;
	v171 = int32(70)
	goto L171
L171:
	;
	v396 = v171
	goto L4
L172:
	;
	v189 = int32(192)
	goto L174
L173:
	;
	v189 = int32(45)
	goto L174
L174:
	;
	v396 = v189
	goto L4
L175:
	;
	v396 = int32(0)
	goto L4
L176:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v197 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v198 = int32(184)
	goto L179
L178:
	;
	v198 = int32(89)
	goto L179
L179:
	;
	v396 = v198
	goto L4
L180:
	;
	v396 = int32(0)
	goto L4
L181:
	;
	v213 = v207 + int32(105)
	goto L183
L182:
	;
	v213 = int32(0)
	goto L183
L183:
	;
	v396 = v213
	goto L4
L184:
	;
	v252 = int32(100)
	goto L186
L185:
	;
	v252 = int32(101)
	goto L186
L186:
	;
	v396 = v252
	goto L4
L187:
	;
	v404 = int32(88)
	goto L3
L188:
	;
	goto L189
L189:
	;
	switch v253 - int32(1) {
	case 0:
		goto L191
	case 1:
		v396 = int32(191)
		goto L4
	case 2:
		v359 = int32(158)
		goto L12
	case 3:
		v362 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L190
	}
L190:
	;
	v278 = int32(0)
	v281 = F_errstart(m, int32(19), v278)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v262 = int32(179)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v263 == int32(0) {
		v396 = v262
		goto L4
	} else {
		goto L192
	}
L192:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	v270 = v268 - int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v270) {
		v396 = v262
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v270<<(uint(int32(2))%32))+uint32(_consts[909])))
	v396 = v277
	goto L4
L194:
	;
	return int32(0)
L195:
	;
	if v281 == int32(0) {
		v396 = v278
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v287
	F_errmsg_internal(m, int32(504617), v9+int32(16))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L194
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(511117), int32(3158), int32(351782))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v396 = v278
	goto L4
L199:
	;
	v404 = int32(28)
	goto L3
L200:
	;
	goto L201
L201:
	;
	switch v299 - int32(1) {
	case 0:
		goto L203
	case 1:
		v396 = int32(191)
		goto L4
	case 2:
		v359 = int32(158)
		goto L12
	case 3:
		v362 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L202
	}
L202:
	;
	v319 = int32(0)
	v322 = F_errstart(m, int32(19), v319)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L194
	} else {
		goto L206
	}
L203:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
	if v308 == int32(0) {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	v315 = v313 - int32(1)
	if base.Ui32(v315) < base.Ui32(int32(4)) {
		goto L6
	} else {
		goto L205
	}
L205:
	;
	v396 = int32(0)
	goto L4
L206:
	;
	if v322 == int32(0) {
		v396 = v319
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v326
	F_errmsg_internal(m, int32(504617), v9+int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L194
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(511117), int32(3221), int32(351782))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L194
	} else {
		goto L209
	}
L209:
	;
	v396 = v319
	goto L4
L210:
	;
	if v341 == int32(0) {
		v396 = v338
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v345
	F_errmsg_internal(m, int32(504300), v9)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L194
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(511117), int32(3230), int32(351782))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L194
	} else {
		goto L213
	}
L213:
	;
	v396 = v338
	goto L4
}
