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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v16 != 0 {
				v38 = v2
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
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
						F_errfinish(m, int32(_a_F_CommandIsReadOnly_1), int32(117), int32(_a_F_CommandIsReadOnly_2))
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
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
		v392 = int32(4)
		goto L3
	case 70:
		goto L13
	case 71:
		v355 = int32(103)
		goto L12
	case 72:
		v358 = int32(192)
		goto L11
	case 73:
		goto L10
	case 74, 77:
		goto L5
	case 78:
		goto L112
	case 79:
		goto L86
	case 83:
		goto L23
	case 84:
		goto L85
	case 85:
		goto L83
	case 88:
		goto L82
	case 89:
		goto L81
	case 90:
		goto L92
	case 91:
		goto L56
	case 92:
		goto L55
	case 93:
		goto L111
	case 95:
		goto L110
	case 96:
		goto L109
	case 97:
		goto L108
	case 98:
		goto L87
	case 99:
		goto L107
	case 100, 101:
		goto L106
	case 102:
		goto L105
	case 103:
		goto L104
	case 104:
		goto L103
	case 105:
		goto L102
	case 106:
		goto L98
	case 107:
		goto L101
	case 108:
		goto L100
	case 109:
		goto L99
	case 110:
		goto L97
	case 111:
		goto L31
	case 112:
		goto L30
	case 113:
		goto L29
	case 114:
		goto L52
	case 115:
		goto L51
	case 116:
		goto L50
	case 117:
		goto L49
	case 118:
		goto L48
	case 119, 120:
		goto L47
	case 121:
		goto L46
	case 122:
		goto L73
	case 123:
		goto L72
	case 124:
		goto L80
	case 125:
		goto L113
	case 126:
		goto L37
	case 128:
		goto L36
	case 129:
		goto L35
	case 130:
		goto L96
	case 131:
		goto L95
	case 132:
		goto L94
	case 133:
		goto L93
	case 134:
		v384 = int32(102)
		goto L4
	case 135:
		goto L115
	case 136:
		goto L114
	case 137:
		goto L75
	case 138:
		goto L20
	case 140:
		goto L19
	case 141:
		goto L76
	case 143:
		goto L84
	case 144:
		goto L71
	case 146:
		goto L63
	case 148:
		goto L91
	case 149:
		goto L90
	case 150:
		goto L89
	case 151:
		goto L88
	case 152:
		goto L34
	case 153, 162:
		goto L78
	case 154:
		goto L74
	case 155:
		goto L67
	case 156:
		goto L66
	case 157:
		goto L65
	case 158:
		goto L116
	case 159, 160, 161:
		goto L79
	case 163:
		goto L77
	case 164:
		goto L64
	case 165:
		goto L70
	case 166, 167, 168:
		goto L69
	case 169:
		goto L68
	case 170:
		goto L57
	case 171:
		goto L62
	case 173:
		goto L61
	case 174:
		goto L60
	case 175:
		goto L59
	case 176:
		goto L58
	case 177:
		goto L41
	case 178:
		goto L54
	case 179:
		goto L43
	case 180:
		goto L42
	case 181:
		goto L40
	case 182:
		goto L39
	case 183:
		goto L38
	case 184:
		goto L53
	case 185:
		goto L22
	case 186:
		goto L21
	case 187:
		goto L18
	case 188:
		goto L45
	case 189:
		goto L44
	case 190:
		goto L33
	case 191:
		goto L32
	case 195:
		goto L28
	case 196:
		goto L27
	case 197:
		goto L26
	case 198:
		goto L25
	case 199:
		goto L24
	case 200:
		goto L17
	case 267:
		goto L16
	}
L3:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v11+v392)))
	v11 = v394
	goto L1
L4:
	;
	m.G0 = v9 + int32(48)
	return v384
L5:
	;
	v384 = int32(180)
	goto L4
L6:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v311<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[0])))
	v384 = v376
	goto L4
L7:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v199<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[1])))
	v384 = v373
	goto L4
L8:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[2])))
	v384 = v370
	goto L4
L9:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[3])))
	v384 = v367
	goto L4
L10:
	;
	v384 = int32(163)
	goto L4
L11:
	;
	v384 = v358
	goto L4
L12:
	;
	v384 = v355
	goto L4
L13:
	;
	v384 = int32(158)
	goto L4
L14:
	;
	v334 = int32(0)
	v337 = F_errstart(m, int32(19), v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L200
	} else {
		goto L216
	}
L15:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v295 == int32(6) {
		goto L205
	} else {
		goto L206
	}
L16:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v251 == int32(6) {
		goto L193
	} else {
		goto L194
	}
L17:
	;
	v384 = int32(194)
	goto L4
L18:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v248 != 0 {
		goto L190
	} else {
		goto L191
	}
L19:
	;
	v384 = int32(31)
	goto L4
L20:
	;
	v384 = int32(86)
	goto L4
L21:
	;
	v384 = int32(152)
	goto L4
L22:
	;
	v384 = int32(166)
	goto L4
L23:
	;
	v384 = int32(4)
	goto L4
L24:
	;
	v384 = int32(140)
	goto L4
L25:
	;
	v384 = int32(32)
	goto L4
L26:
	;
	v384 = int32(87)
	goto L4
L27:
	;
	v384 = int32(24)
	goto L4
L28:
	;
	v384 = int32(79)
	goto L4
L29:
	;
	v384 = int32(58)
	goto L4
L30:
	;
	v384 = int32(22)
	goto L4
L31:
	;
	v384 = int32(77)
	goto L4
L32:
	;
	v384 = int32(36)
	goto L4
L33:
	;
	v384 = int32(37)
	goto L4
L34:
	;
	v384 = int32(19)
	goto L4
L35:
	;
	v384 = int32(21)
	goto L4
L36:
	;
	v384 = int32(76)
	goto L4
L37:
	;
	v384 = int32(75)
	goto L4
L38:
	;
	v384 = int32(60)
	goto L4
L39:
	;
	v384 = int32(63)
	goto L4
L40:
	;
	v384 = int32(170)
	goto L4
L41:
	;
	v384 = int32(48)
	goto L4
L42:
	;
	v384 = int32(187)
	goto L4
L43:
	;
	v384 = int32(161)
	goto L4
L44:
	;
	v384 = int32(168)
	goto L4
L45:
	;
	v384 = int32(129)
	goto L4
L46:
	;
	v384 = int32(133)
	goto L4
L47:
	;
	v384 = int32(25)
	goto L4
L48:
	;
	v384 = int32(80)
	goto L4
L49:
	;
	v384 = int32(72)
	goto L4
L50:
	;
	v384 = int32(10)
	goto L4
L51:
	;
	v384 = int32(66)
	goto L4
L52:
	;
	v384 = int32(96)
	goto L4
L53:
	;
	v384 = int32(95)
	goto L4
L54:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v204) < base.Ui32(int32(4)) {
		goto L187
	} else {
		goto L188
	}
L55:
	;
	v384 = int32(188)
	goto L4
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v199) < base.Ui32(int32(6)) {
		goto L7
	} else {
		goto L186
	}
L57:
	;
	v384 = int32(33)
	goto L4
L58:
	;
	v384 = int32(169)
	goto L4
L59:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v187 = v185 - int32(23)
	if v187 == int32(0) {
		v384 = int32(73)
		goto L4
	} else {
		goto L179
	}
L60:
	;
	v384 = int32(153)
	goto L4
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v179 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L62:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v175 != 0 {
		goto L173
	} else {
		goto L174
	}
L63:
	;
	v384 = int32(47)
	goto L4
L64:
	;
	v384 = int32(160)
	goto L4
L65:
	;
	v384 = int32(191)
	goto L4
L66:
	;
	v384 = int32(159)
	goto L4
L67:
	;
	v384 = int32(165)
	goto L4
L68:
	;
	v384 = int32(116)
	goto L4
L69:
	;
	v384 = int32(7)
	goto L4
L70:
	;
	v384 = int32(64)
	goto L4
L71:
	;
	v384 = int32(109)
	goto L4
L72:
	;
	v384 = int32(29)
	goto L4
L73:
	;
	v384 = int32(84)
	goto L4
L74:
	;
	v384 = int32(82)
	goto L4
L75:
	;
	v384 = int32(71)
	goto L4
L76:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v158 != 0 {
		goto L170
	} else {
		goto L171
	}
L77:
	;
	v384 = int32(99)
	goto L4
L78:
	;
	v384 = int32(42)
	goto L4
L79:
	;
	v384 = int32(97)
	goto L4
L80:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v143 {
	case 0:
		goto L162
	case 1:
		v384 = int32(59)
		goto L4
	default:
		goto L161
	case 7:
		goto L163
	case 25:
		goto L169
	case 46:
		goto L164
	case 47:
		goto L166
	case 48:
		goto L167
	case 49:
		goto L165
	case 50:
		goto L168
	}
L81:
	;
	v384 = int32(8)
	goto L4
L82:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v139 != 0 {
		goto L158
	} else {
		goto L159
	}
L83:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v135 != 0 {
		goto L155
	} else {
		goto L156
	}
L84:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v127 - int32(19) {
	case 0:
		v384 = int32(14)
		goto L4
	default:
		goto L152
	case 10:
		goto L154
	case 16:
		goto L153
	}
L85:
	;
	v384 = int32(9)
	goto L4
L86:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v117 = v115 - int32(1)
	if base.Ui32(v117) <= base.Ui32(int32(51)) {
		goto L149
	} else {
		goto L150
	}
L87:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v107 = v105 - int32(1)
	if base.Ui32(v107) <= base.Ui32(int32(51)) {
		goto L145
	} else {
		goto L146
	}
L88:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v97 = v95 - int32(1)
	if base.Ui32(v97) <= base.Ui32(int32(51)) {
		goto L141
	} else {
		goto L142
	}
L89:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v87 = v85 - int32(1)
	if base.Ui32(v87) <= base.Ui32(int32(51)) {
		goto L137
	} else {
		goto L138
	}
L90:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v77 = v75 - int32(1)
	if base.Ui32(v77) <= base.Ui32(int32(51)) {
		goto L133
	} else {
		goto L134
	}
L91:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v61 == int32(6) {
		goto L125
	} else {
		goto L126
	}
L92:
	;
	v384 = int32(56)
	goto L4
L93:
	;
	v384 = int32(179)
	goto L4
L94:
	;
	v384 = int32(53)
	goto L4
L95:
	;
	v384 = int32(190)
	goto L4
L96:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if base.Ui32(v53) < base.Ui32(int32(53)) {
		goto L8
	} else {
		goto L124
	}
L97:
	;
	v384 = int32(157)
	goto L4
L98:
	;
	v384 = int32(69)
	goto L4
L99:
	;
	v384 = int32(150)
	goto L4
L100:
	;
	v384 = int32(43)
	goto L4
L101:
	;
	v384 = int32(98)
	goto L4
L102:
	;
	v384 = int32(30)
	goto L4
L103:
	;
	v384 = int32(85)
	goto L4
L104:
	;
	v384 = int32(12)
	goto L4
L105:
	;
	v384 = int32(68)
	goto L4
L106:
	;
	v384 = int32(11)
	goto L4
L107:
	;
	v384 = int32(67)
	goto L4
L108:
	;
	v384 = int32(35)
	goto L4
L109:
	;
	v384 = int32(142)
	goto L4
L110:
	;
	v384 = int32(90)
	goto L4
L111:
	;
	v384 = int32(88)
	goto L4
L112:
	;
	v384 = int32(83)
	goto L4
L113:
	;
	v384 = int32(65)
	goto L4
L114:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v34 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v30 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v24) < base.Ui32(int32(10)) {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	v384 = int32(0)
	goto L4
L118:
	;
	v31 = int32(50)
	goto L120
L119:
	;
	v31 = int32(51)
	goto L120
L120:
	;
	v384 = v31
	goto L4
L121:
	;
	v35 = int32(164)
	goto L123
L122:
	;
	v35 = int32(154)
	goto L123
L123:
	;
	v384 = v35
	goto L4
L124:
	;
	v384 = int32(0)
	goto L4
L125:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v65 = v64
	goto L127
L126:
	;
	v65 = v61
	goto L127
L127:
	;
	v67 = v65 - int32(1)
	if base.Ui32(v67) <= base.Ui32(int32(51)) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v384 = v74
	goto L4
L129:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v74 = v72
	goto L131
L130:
	;
	v74 = int32(0)
	goto L131
L131:
	;
	goto L128
L132:
	;
	v384 = v84
	goto L4
L133:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v84 = v82
	goto L135
L134:
	;
	v84 = int32(0)
	goto L135
L135:
	;
	goto L132
L136:
	;
	v384 = v94
	goto L4
L137:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v94 = v92
	goto L139
L138:
	;
	v94 = int32(0)
	goto L139
L139:
	;
	goto L136
L140:
	;
	v384 = v104
	goto L4
L141:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v104 = v102
	goto L143
L142:
	;
	v104 = int32(0)
	goto L143
L143:
	;
	goto L140
L144:
	;
	v384 = v114
	goto L4
L145:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v114 = v112
	goto L147
L146:
	;
	v114 = int32(0)
	goto L147
L147:
	;
	goto L144
L148:
	;
	v384 = v124
	goto L4
L149:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[4])))
	v124 = v122
	goto L151
L150:
	;
	v124 = int32(0)
	goto L151
L151:
	;
	goto L148
L152:
	;
	v384 = int32(0)
	goto L4
L153:
	;
	v384 = int32(26)
	goto L4
L154:
	;
	v384 = int32(23)
	goto L4
L155:
	;
	v136 = int32(155)
	goto L157
L156:
	;
	v136 = int32(174)
	goto L157
L157:
	;
	v384 = v136
	goto L4
L158:
	;
	v140 = int32(156)
	goto L160
L159:
	;
	v140 = int32(175)
	goto L160
L160:
	;
	v384 = v140
	goto L4
L161:
	;
	v384 = int32(0)
	goto L4
L162:
	;
	v384 = int32(58)
	goto L4
L163:
	;
	v384 = int32(61)
	goto L4
L164:
	;
	v384 = int32(91)
	goto L4
L165:
	;
	v384 = int32(94)
	goto L4
L166:
	;
	v384 = int32(92)
	goto L4
L167:
	;
	v384 = int32(93)
	goto L4
L168:
	;
	v384 = int32(97)
	goto L4
L169:
	;
	v384 = int32(74)
	goto L4
L170:
	;
	v159 = int32(78)
	goto L172
L171:
	;
	v159 = int32(70)
	goto L172
L172:
	;
	v384 = v159
	goto L4
L173:
	;
	v176 = int32(193)
	goto L175
L174:
	;
	v176 = int32(45)
	goto L175
L175:
	;
	v384 = v176
	goto L4
L176:
	;
	v182 = int32(52)
	goto L178
L177:
	;
	v182 = int32(172)
	goto L178
L178:
	;
	v384 = v182
	goto L4
L179:
	;
	if v187 == int32(19) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v194 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v384 = int32(0)
	goto L4
L183:
	;
	v195 = int32(185)
	goto L185
L184:
	;
	v195 = int32(89)
	goto L185
L185:
	;
	v384 = v195
	goto L4
L186:
	;
	v384 = int32(0)
	goto L4
L187:
	;
	v210 = v204 + int32(105)
	goto L189
L188:
	;
	v210 = int32(0)
	goto L189
L189:
	;
	v384 = v210
	goto L4
L190:
	;
	v249 = int32(100)
	goto L192
L191:
	;
	v249 = int32(101)
	goto L192
L192:
	;
	v384 = v249
	goto L4
L193:
	;
	v392 = int32(100)
	goto L3
L194:
	;
	goto L195
L195:
	;
	switch v251 - int32(1) {
	case 0:
		goto L197
	case 1:
		v384 = int32(192)
		goto L4
	case 2:
		v355 = int32(158)
		goto L12
	case 3:
		v358 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L196
	}
L196:
	;
	v274 = int32(0)
	v277 = F_errstart(m, int32(19), v274)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	v260 = int32(180)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v261 == int32(0) {
		v384 = v260
		goto L4
	} else {
		goto L198
	}
L198:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v268 = v266 - int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v268) {
		v384 = v260
		goto L4
	} else {
		goto L199
	}
L199:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268<<(uint(int32(2))%32))+uint32(_c_F_CreateCommandTag[0])))
	v384 = v273
	goto L4
L200:
	;
	return int32(0)
L201:
	;
	if v277 == int32(0) {
		v384 = v274
		goto L4
	} else {
		goto L202
	}
L202:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v283
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_0), v9+int32(16))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3189), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L200
	} else {
		goto L204
	}
L204:
	;
	v384 = v274
	goto L4
L205:
	;
	v392 = int32(28)
	goto L3
L206:
	;
	goto L207
L207:
	;
	switch v295 - int32(1) {
	case 0:
		goto L209
	case 1:
		v384 = int32(192)
		goto L4
	case 2:
		v355 = int32(158)
		goto L12
	case 3:
		v358 = int32(103)
		goto L11
	case 4:
		goto L10
	default:
		goto L208
	}
L208:
	;
	v315 = int32(0)
	v318 = F_errstart(m, int32(19), v315)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L200
	} else {
		goto L212
	}
L209:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
	if v304 == int32(0) {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v311 = v309 - int32(1)
	if base.Ui32(v311) < base.Ui32(int32(4)) {
		goto L6
	} else {
		goto L211
	}
L211:
	;
	v384 = int32(0)
	goto L4
L212:
	;
	if v318 == int32(0) {
		v384 = v315
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v322
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_0), v9+int32(32))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L200
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3252), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L200
	} else {
		goto L215
	}
L215:
	;
	v384 = v315
	goto L4
L216:
	;
	if v337 == int32(0) {
		v384 = v334
		goto L4
	} else {
		goto L217
	}
L217:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v341
	F_errmsg_internal(m, int32(_a_F_CreateCommandTag_3), v9)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L200
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CreateCommandTag_1), int32(3261), int32(_a_F_CreateCommandTag_2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L200
	} else {
		goto L219
	}
L219:
	;
	v384 = v334
	goto L4
}
func F_EndCommandExtended(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = v11 << (uint(int32(3)) % 32)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_EndCommandExtended[0]))))
	if v14 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_EndCommandExtended[1])))
		base.MemoryCopy(m, v5, v15, v14)
	} else {
	}
	v17 = v5 + v14
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_EndCommandExtended[2]))))
	if v18&int32(1) != 0 {
		if v11 == int32(158) {
			v23 = int32(_a_F_EndCommandExtended_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v23)
			v27 = v17 + int32(2)
		} else {
			v27 = v17
		}
		v28 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v28)
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v32 = v27 + int32(1)
		v33 = F_pg_ulltoa_n(m, v30, v32)
		mBase = m.M
		v36 = v33 + v32
	} else {
		v36 = v17
	}
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v37)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_EndCommandExtended[3]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	m.T0[v44].(func(*base.Module, int32, int32, int32))(m, int32(67), v5, v36-v5+int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		return
	} else {
		m.G0 = v5 - int32(-64)
		return
	}
}
