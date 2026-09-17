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
					F_errmsg_internal(m, int32(_a_F_CommandIsReadOnly_0), v7)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_CommandIsReadOnly_1), int32(116), int32(_a_F_CommandIsReadOnly_2))
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
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
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
		v386 = int32(4)
		goto L3
	case 70:
		goto L13
	case 71:
		v349 = int32(103)
		goto L12
	case 72:
		v352 = int32(191)
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
		v378 = int32(102)
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
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v11+v386)))
	v11 = v388
	goto L1
L4:
	;
	m.G0 = v9 + int32(48)
	return v378
L5:
	;
	v378 = int32(179)
	goto L4
L6:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v305<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[0])))
	v378 = v370
	goto L4
L7:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v194<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[1])))
	v378 = v367
	goto L4
L8:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[2])))
	v378 = v364
	goto L4
L9:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[3])))
	v378 = v361
	goto L4
L10:
	;
	v378 = int32(163)
	goto L4
L11:
	;
	v378 = v352
	goto L4
L12:
	;
	v378 = v349
	goto L4
L13:
	;
	v378 = int32(158)
	goto L4
L14:
	;
	v328 = int32(0)
	v331 = F_errstart(m, int32(19), v328)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L196
	} else {
		goto L212
	}
L15:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v289 == int32(6) {
		goto L201
	} else {
		goto L202
	}
L16:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v245 == int32(6) {
		goto L189
	} else {
		goto L190
	}
L17:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v243 != 0 {
		goto L186
	} else {
		goto L187
	}
L18:
	;
	v378 = int32(31)
	goto L4
L19:
	;
	v378 = int32(86)
	goto L4
L20:
	;
	v378 = int32(152)
	goto L4
L21:
	;
	v378 = int32(166)
	goto L4
L22:
	;
	v378 = int32(4)
	goto L4
L23:
	;
	v378 = int32(140)
	goto L4
L24:
	;
	v378 = int32(32)
	goto L4
L25:
	;
	v378 = int32(87)
	goto L4
L26:
	;
	v378 = int32(24)
	goto L4
L27:
	;
	v378 = int32(79)
	goto L4
L28:
	;
	v378 = int32(58)
	goto L4
L29:
	;
	v378 = int32(22)
	goto L4
L30:
	;
	v378 = int32(77)
	goto L4
L31:
	;
	v378 = int32(36)
	goto L4
L32:
	;
	v378 = int32(37)
	goto L4
L33:
	;
	v378 = int32(19)
	goto L4
L34:
	;
	v378 = int32(21)
	goto L4
L35:
	;
	v378 = int32(76)
	goto L4
L36:
	;
	v378 = int32(75)
	goto L4
L37:
	;
	v378 = int32(60)
	goto L4
L38:
	;
	v378 = int32(63)
	goto L4
L39:
	;
	v378 = int32(170)
	goto L4
L40:
	;
	v378 = int32(48)
	goto L4
L41:
	;
	v378 = int32(186)
	goto L4
L42:
	;
	v378 = int32(161)
	goto L4
L43:
	;
	v378 = int32(168)
	goto L4
L44:
	;
	v378 = int32(129)
	goto L4
L45:
	;
	v378 = int32(133)
	goto L4
L46:
	;
	v378 = int32(25)
	goto L4
L47:
	;
	v378 = int32(80)
	goto L4
L48:
	;
	v378 = int32(72)
	goto L4
L49:
	;
	v378 = int32(10)
	goto L4
L50:
	;
	v378 = int32(66)
	goto L4
L51:
	;
	v378 = int32(96)
	goto L4
L52:
	;
	v378 = int32(95)
	goto L4
L53:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v199) < base.Ui32(int32(4)) {
		goto L183
	} else {
		goto L184
	}
L54:
	;
	v378 = int32(187)
	goto L4
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v194) < base.Ui32(int32(6)) {
		goto L7
	} else {
		goto L182
	}
L56:
	;
	v378 = int32(33)
	goto L4
L57:
	;
	v378 = int32(169)
	goto L4
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v182 = v180 - int32(23)
	if v182 == int32(0) {
		v378 = int32(73)
		goto L4
	} else {
		goto L175
	}
L59:
	;
	v378 = int32(153)
	goto L4
L60:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v176 != 0 {
		goto L172
	} else {
		goto L173
	}
L61:
	;
	v378 = int32(52)
	goto L4
L62:
	;
	v378 = int32(47)
	goto L4
L63:
	;
	v378 = int32(160)
	goto L4
L64:
	;
	v378 = int32(190)
	goto L4
L65:
	;
	v378 = int32(159)
	goto L4
L66:
	;
	v378 = int32(165)
	goto L4
L67:
	;
	v378 = int32(116)
	goto L4
L68:
	;
	v378 = int32(7)
	goto L4
L69:
	;
	v378 = int32(64)
	goto L4
L70:
	;
	v378 = int32(109)
	goto L4
L71:
	;
	v378 = int32(29)
	goto L4
L72:
	;
	v378 = int32(84)
	goto L4
L73:
	;
	v378 = int32(82)
	goto L4
L74:
	;
	v378 = int32(71)
	goto L4
L75:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v158 != 0 {
		goto L169
	} else {
		goto L170
	}
L76:
	;
	v378 = int32(99)
	goto L4
L77:
	;
	v378 = int32(42)
	goto L4
L78:
	;
	v378 = int32(97)
	goto L4
L79:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v143 {
	case 0:
		goto L161
	case 1:
		v378 = int32(59)
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
	v378 = int32(8)
	goto L4
L81:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v139 != 0 {
		goto L157
	} else {
		goto L158
	}
L82:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v135 != 0 {
		goto L154
	} else {
		goto L155
	}
L83:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v127 - int32(19) {
	case 0:
		v378 = int32(14)
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
	v378 = int32(9)
	goto L4
L85:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v117 = v115 - int32(1)
	if base.Ui32(v117) <= base.Ui32(int32(50)) {
		goto L148
	} else {
		goto L149
	}
L86:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v107 = v105 - int32(1)
	if base.Ui32(v107) <= base.Ui32(int32(50)) {
		goto L144
	} else {
		goto L145
	}
L87:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v97 = v95 - int32(1)
	if base.Ui32(v97) <= base.Ui32(int32(50)) {
		goto L140
	} else {
		goto L141
	}
L88:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v87 = v85 - int32(1)
	if base.Ui32(v87) <= base.Ui32(int32(50)) {
		goto L136
	} else {
		goto L137
	}
L89:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v77 = v75 - int32(1)
	if base.Ui32(v77) <= base.Ui32(int32(50)) {
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
	v378 = int32(56)
	goto L4
L92:
	;
	v378 = int32(178)
	goto L4
L93:
	;
	v378 = int32(53)
	goto L4
L94:
	;
	v378 = int32(189)
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
	v378 = int32(157)
	goto L4
L97:
	;
	v378 = int32(69)
	goto L4
L98:
	;
	v378 = int32(150)
	goto L4
L99:
	;
	v378 = int32(43)
	goto L4
L100:
	;
	v378 = int32(98)
	goto L4
L101:
	;
	v378 = int32(30)
	goto L4
L102:
	;
	v378 = int32(85)
	goto L4
L103:
	;
	v378 = int32(12)
	goto L4
L104:
	;
	v378 = int32(68)
	goto L4
L105:
	;
	v378 = int32(11)
	goto L4
L106:
	;
	v378 = int32(67)
	goto L4
L107:
	;
	v378 = int32(35)
	goto L4
L108:
	;
	v378 = int32(142)
	goto L4
L109:
	;
	v378 = int32(90)
	goto L4
L110:
	;
	v378 = int32(88)
	goto L4
L111:
	;
	v378 = int32(83)
	goto L4
L112:
	;
	v378 = int32(65)
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
	v378 = int32(0)
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
	v378 = v31
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
	v378 = v35
	goto L4
L123:
	;
	v378 = int32(0)
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
	v67 = v65 - int32(1)
	if base.Ui32(v67) <= base.Ui32(int32(50)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v378 = v74
	goto L4
L128:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v74 = v72
	goto L130
L129:
	;
	v74 = int32(0)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v378 = v84
	goto L4
L132:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v84 = v82
	goto L134
L133:
	;
	v84 = int32(0)
	goto L134
L134:
	;
	goto L131
L135:
	;
	v378 = v94
	goto L4
L136:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v94 = v92
	goto L138
L137:
	;
	v94 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	v378 = v104
	goto L4
L140:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v104 = v102
	goto L142
L141:
	;
	v104 = int32(0)
	goto L142
L142:
	;
	goto L139
L143:
	;
	v378 = v114
	goto L4
L144:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v114 = v112
	goto L146
L145:
	;
	v114 = int32(0)
	goto L146
L146:
	;
	goto L143
L147:
	;
	v378 = v124
	goto L4
L148:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v124 = v122
	goto L150
L149:
	;
	v124 = int32(0)
	goto L150
L150:
	;
	goto L147
L151:
	;
	v378 = int32(0)
	goto L4
L152:
	;
	v378 = int32(26)
	goto L4
L153:
	;
	v378 = int32(23)
	goto L4
L154:
	;
	v136 = int32(155)
	goto L156
L155:
	;
	v136 = int32(173)
	goto L156
L156:
	;
	v378 = v136
	goto L4
L157:
	;
	v140 = int32(156)
	goto L159
L158:
	;
	v140 = int32(174)
	goto L159
L159:
	;
	v378 = v140
	goto L4
L160:
	;
	v378 = int32(0)
	goto L4
L161:
	;
	v378 = int32(58)
	goto L4
L162:
	;
	v378 = int32(61)
	goto L4
L163:
	;
	v378 = int32(91)
	goto L4
L164:
	;
	v378 = int32(94)
	goto L4
L165:
	;
	v378 = int32(92)
	goto L4
L166:
	;
	v378 = int32(93)
	goto L4
L167:
	;
	v378 = int32(97)
	goto L4
L168:
	;
	v378 = int32(74)
	goto L4
L169:
	;
	v159 = int32(78)
	goto L171
L170:
	;
	v159 = int32(70)
	goto L171
L171:
	;
	v378 = v159
	goto L4
L172:
	;
	v177 = int32(192)
	goto L174
L173:
	;
	v177 = int32(45)
	goto L174
L174:
	;
	v378 = v177
	goto L4
L175:
	;
	if v182 == int32(18) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v189 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	v378 = int32(0)
	goto L4
L179:
	;
	v190 = int32(184)
	goto L181
L180:
	;
	v190 = int32(89)
	goto L181
L181:
	;
	v378 = v190
	goto L4
L182:
	;
	v378 = int32(0)
	goto L4
L183:
	;
	v205 = v199 + int32(105)
	goto L185
L184:
	;
	v205 = int32(0)
	goto L185
L185:
	;
	v378 = v205
	goto L4
L186:
	;
	v244 = int32(100)
	goto L188
L187:
	;
	v244 = int32(101)
	goto L188
L188:
	;
	v378 = v244
	goto L4
L189:
	;
	v386 = int32(88)
	goto L3
L190:
	;
	goto L191
L191:
	;
	switch v245 - int32(1) {
	case 0:
		goto L193
	case 1:
		v378 = int32(191)
		goto L4
	case 2:
		v349 = int32(158)
		goto L12
	case 3:
		v352 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L192
	}
L192:
	;
	v268 = int32(0)
	v271 = F_errstart(m, int32(19), v268)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v254 = int32(179)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v255 == int32(0) {
		v378 = v254
		goto L4
	} else {
		goto L194
	}
L194:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	v262 = v260 - int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v262) {
		v378 = v254
		goto L4
	} else {
		goto L195
	}
L195:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v262<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[0])))
	v378 = v267
	goto L4
L196:
	;
	return int32(0)
L197:
	;
	if v271 == int32(0) {
		v378 = v268
		goto L4
	} else {
		goto L198
	}
L198:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v277
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_0), v9+int32(16))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L196
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3158), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L196
	} else {
		goto L200
	}
L200:
	;
	v378 = v268
	goto L4
L201:
	;
	v386 = int32(28)
	goto L3
L202:
	;
	goto L203
L203:
	;
	switch v289 - int32(1) {
	case 0:
		goto L205
	case 1:
		v378 = int32(191)
		goto L4
	case 2:
		v349 = int32(158)
		goto L12
	case 3:
		v352 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L204
	}
L204:
	;
	v309 = int32(0)
	v312 = F_errstart(m, int32(19), v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L196
	} else {
		goto L208
	}
L205:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
	if v298 == int32(0) {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+8))
	v305 = v303 - int32(1)
	if base.Ui32(v305) < base.Ui32(int32(4)) {
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v378 = int32(0)
	goto L4
L208:
	;
	if v312 == int32(0) {
		v378 = v309
		goto L4
	} else {
		goto L209
	}
L209:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v316
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_0), v9+int32(32))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L196
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3221), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L196
	} else {
		goto L211
	}
L211:
	;
	v378 = v309
	goto L4
L212:
	;
	if v331 == int32(0) {
		v378 = v328
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v335
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_3), v9)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L196
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3230), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L196
	} else {
		goto L215
	}
L215:
	;
	v378 = v328
	goto L4
}
