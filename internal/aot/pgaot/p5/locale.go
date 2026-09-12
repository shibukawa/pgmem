package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___get_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v302 int64
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v6 != 0 {
		v161 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v164 = int32(0)
	goto L64
L2:
	;
	v7 = int32(546262)
	v8 = int32(0)
	v13 = F___strchrnul(m, v7, int32(61))
	mBase = m.M
	if v7 == v13 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v55 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = v13 - v7
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[1496]))))
	if v18 != 0 {
		v48 = v8
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = v48
	goto L3
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v20 == int32(0) {
		v48 = v8
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v23 == int32(0) {
		v48 = v8
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v27 = v20
	v28 = v23
	goto L11
L11:
	;
	v31 = F_strncmp(m, v7, v28, v16)
	mBase = m.M
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v48 = v35 + int32(1)
	goto L7
L13:
	;
	goto L12
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v35 = v34 + v16
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 == int32(61) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v40 != 0 {
		v27 = v27 + int32(4)
		v28 = v40
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v48 = v8
	goto L7
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != 0 {
		v161 = v55
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v60 = l0*int32(12) + int32(4117648)
	v61 = int32(0)
	v66 = F___strchrnul(m, v60, int32(61))
	mBase = m.M
	if v60 == v66 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	if v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L24:
	;
	v108 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v69 = v66 - v60
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v69))))
	if v71 != 0 {
		v101 = v61
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v108 = v101
	goto L23
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v73 == int32(0) {
		v101 = v61
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v76 == int32(0) {
		v101 = v61
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v80 = v73
	v81 = v76
	goto L31
L31:
	;
	v84 = F_strncmp(m, v60, v81, v69)
	mBase = m.M
	if v84 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v101 = v88 + int32(1)
	goto L27
L33:
	;
	goto L32
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v88 = v87 + v69
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 == int32(61) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v93 != 0 {
		v80 = v80 + int32(4)
		v81 = v93
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v101 = v61
	goto L27
L39:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 != 0 {
		v161 = v108
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v110 = int32(549443)
	v111 = int32(0)
	v116 = F___strchrnul(m, v110, int32(61))
	mBase = m.M
	if v110 == v116 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L41
L43:
	;
	if v158 != 0 {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	v158 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v119 = v116 - v110
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+uint32(_consts[1497]))))
	if v121 != 0 {
		v151 = v111
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v158 = v151
	goto L43
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v123 == int32(0) {
		v151 = v111
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v126 == int32(0) {
		v151 = v111
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v130 = v123
	v131 = v126
	goto L51
L51:
	;
	v134 = F_strncmp(m, v110, v131, v119)
	mBase = m.M
	if v134 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v151 = v138 + int32(1)
	goto L47
L53:
	;
	goto L52
L54:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v138 = v137 + v119
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v139 == int32(61) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v143 != 0 {
		v130 = v130 + int32(4)
		v131 = v143
		goto L51
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v151 = v111
	goto L47
L59:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != 0 {
		v161 = v158
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v161 = int32(569231)
	goto L1
L62:
	;
	goto L61
L63:
	;
	v180 = int32(569231)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v181 == int32(46) {
		v188 = v180
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v164))))
	if v168 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v179 = v164
	goto L63
L66:
	;
	goto L65
L67:
	;
	if v168 == int32(47) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v173 = int32(23)
	v175 = v164 + int32(1)
	if v175 != v173 {
		v164 = v175
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v179 = v173
	goto L63
L70:
	;
	return v322
L71:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[1498]))
	if v258 != 0 {
		goto L101
	} else {
		goto L102
	}
L72:
	;
	if l0 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L73:
	;
	v193 = int32(569231)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1499])))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v197 == int32(0) {
		v216 = v196
		v217 = v197
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v189 == int32(0) {
		v247 = v188
		goto L72
	} else {
		goto L78
	}
L75:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v179))))
	if v185 != 0 {
		v188 = v180
		goto L74
	} else {
		goto L76
	}
L76:
	;
	if v181 != int32(67) {
		v192 = v161
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v188 = v161
	goto L74
L78:
	;
	v192 = v188
	goto L73
L79:
	;
	if v217-v216 == int32(0) {
		v247 = v192
		goto L72
	} else {
		goto L87
	}
L80:
	;
	goto L79
L81:
	;
	if v196 != v197 {
		v216 = v196
		v217 = v197
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v201 = v192
	v202 = v193
	goto L83
L83:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	if v206 == int32(0) {
		v216 = v205
		v217 = v206
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v216 = v205
	v217 = v206
	goto L80
L85:
	;
	v209 = int32(1)
	if v205 == v206 {
		v201 = v201 + v209
		v202 = v202 + v209
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v221 = int32(522700)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1263])))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v225 == int32(0) {
		v244 = v224
		v245 = v225
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v245-v244 != 0 {
		goto L71
	} else {
		goto L96
	}
L89:
	;
	goto L88
L90:
	;
	if v224 != v225 {
		v244 = v224
		v245 = v225
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v229 = v192
	v230 = v221
	goto L92
L92:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	if v234 == int32(0) {
		v244 = v233
		v245 = v234
		goto L89
	} else {
		goto L94
	}
L93:
	;
	v244 = v233
	v245 = v234
	goto L89
L94:
	;
	v237 = int32(1)
	if v233 == v234 {
		v229 = v229 + v237
		v230 = v230 + v237
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v247 = v192
	goto L72
L97:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	if v251 == int32(46) {
		v322 = int32(4117556)
		goto L70
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	return int32(0)
L100:
	;
	goto L99
L101:
	;
	v261 = v258
	goto L104
L102:
	;
	goto L103
L103:
	;
	v300 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v300 != 0 {
		goto L116
	} else {
		goto L117
	}
L104:
	;
	v265 = v261 + int32(8)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v269 == int32(0) {
		v288 = v268
		v289 = v269
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L103
L106:
	;
	if v289-v288 == int32(0) {
		v322 = v261
		goto L70
	} else {
		goto L114
	}
L107:
	;
	goto L106
L108:
	;
	if v268 != v269 {
		v288 = v268
		v289 = v269
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v273 = v192
	v274 = v265
	goto L110
L110:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v278 == int32(0) {
		v288 = v277
		v289 = v278
		goto L107
	} else {
		goto L112
	}
L111:
	;
	v288 = v277
	v289 = v278
	goto L107
L112:
	;
	v281 = int32(1)
	if v277 == v278 {
		v273 = v273 + v281
		v274 = v274 + v281
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v261)+32))
	if v293 != 0 {
		v261 = v293
		goto L104
	} else {
		goto L115
	}
L115:
	;
	goto L105
L116:
	;
	v302 = *(*int64)(unsafe.Add(mBase, _consts[1500]))
	*(*int64)(unsafe.Add(mBase, uint32(v300))) = v302
	v305 = v300 + int32(8)
	if v179 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L118
L118:
	;
	if l0|v300 != 0 {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305+v179))) = uint8(v309)
	v311 = int32(4701152)
	v312 = *(*int32)(unsafe.Add(mBase, _consts[1498]))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+32)) = v312
	*(*int32)(unsafe.Add(mBase, _consts[1498])) = v300
	goto L118
L120:
	;
	v306 = F__emscripten_memcpy_bulkmem(m, v305, v192, v179)
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	v319 = v300
	goto L125
L124:
	;
	v319 = int32(4117556)
	goto L125
L125:
	;
	v322 = v319
	goto L70
}
func F_check_locale_numeric(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_check_locale(m, int32(1), v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
