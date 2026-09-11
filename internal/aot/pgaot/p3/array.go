package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayCastAndSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v438 int32
	_ = v438
	v1 = l0
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if int32(0) < l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v438
L2:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	if l1 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	switch l3 - int32(99) {
	case 0:
		v438 = l1
		goto L1
	case 1:
		goto L64
	default:
		goto L63
	case 6:
		goto L65
	}
L6:
	;
	switch l1 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L11
	default:
		goto L9
	case 3:
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l4 == v1 {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1
	goto L5
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v1)
	goto L5
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1)
	goto L5
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(_a_F_ArrayCastAndSet_0), v9)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_ArrayCastAndSet_1), int32(230), int32(_a_F_ArrayCastAndSet_2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L5
L18:
	;
	goto L17
L19:
	;
	v36 = l4 + l1
	if base.Ui32(v1-v36) <= base.Ui32(int32(0)-l1<<(uint(int32(1))%32)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v43 = F___memcpy(m, l4, v1, l1)
	mBase = m.M
	goto L17
L21:
	;
	goto L22
L22:
	;
	v46 = (l4 ^ v1) & int32(3)
	if base.Ui32(l4) < base.Ui32(v1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v148 == int32(0) {
		goto L18
	} else {
		goto L59
	}
L24:
	;
	if base.Ui32(v126) <= base.Ui32(int32(3)) {
		v147 = v125
		v148 = v126
		v149 = v127
		goto L23
	} else {
		goto L55
	}
L25:
	;
	if v46 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v46 != 0 {
		v108 = l1
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v147 = v1
	v148 = l1
	v149 = l4
	goto L23
L29:
	;
	goto L30
L30:
	;
	if l4&int32(3) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v125 = v1
	v126 = l1
	v127 = l4
	goto L24
L32:
	;
	goto L33
L33:
	;
	v53 = v1
	v54 = l1
	v55 = l4
	goto L34
L34:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v125 = v62
	v126 = v64
	v127 = v66
	goto L24
L36:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v61 = int32(1)
	v62 = v53 + v61
	v64 = v54 - v61
	v66 = v55 + v61
	if v66&int32(3) != 0 {
		v53 = v62
		v54 = v64
		v55 = v66
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == int32(0) {
		goto L18
	} else {
		goto L51
	}
L39:
	;
	if v36&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v73 = l1
	goto L43
L41:
	;
	v88 = l1
	goto L42
L42:
	;
	if base.Ui32(v88) <= base.Ui32(int32(3)) {
		v108 = v88
		goto L38
	} else {
		goto L47
	}
L43:
	;
	if v73 == int32(0) {
		goto L18
	} else {
		goto L45
	}
L44:
	;
	v88 = v79
	goto L42
L45:
	;
	v79 = v73 - int32(1)
	v80 = l4 + v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1+v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v82)
	if v80&int32(3) != 0 {
		v73 = v79
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v95 = v88
	goto L48
L48:
	;
	v99 = v95 - int32(4)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v1+v99)))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v99))) = v102
	if base.Ui32(int32(3)) < base.Ui32(v99) {
		v95 = v99
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v108 = v99
	goto L38
L50:
	;
	goto L49
L51:
	;
	v115 = v108
	goto L52
L52:
	;
	v119 = v115 - int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1+v119))))
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v119))) = uint8(v122)
	if v119 != 0 {
		v115 = v119
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L18
L54:
	;
	goto L53
L55:
	;
	v132 = v125
	v133 = v126
	v134 = v127
	goto L56
L56:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v136
	v138 = int32(4)
	v139 = v132 + v138
	v141 = v134 + v138
	v143 = v133 - v138
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		v132 = v139
		v133 = v143
		v134 = v141
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v147 = v139
	v148 = v143
	v149 = v141
	goto L23
L58:
	;
	goto L57
L59:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L60
L60:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v158)
	v160 = int32(1)
	v165 = v155 - v160
	if v165 != 0 {
		v154 = v154 + v160
		v155 = v165
		v156 = v156 + v160
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L18
L62:
	;
	goto L61
L63:
	;
	v438 = (l1 + int32(1)) & int32(-2)
	goto L1
L64:
	;
	v438 = (l1 + int32(7)) & int32(-8)
	goto L1
L65:
	;
	v438 = (l1 + int32(3)) & int32(-4)
	goto L1
L66:
	;
	if l4 == v1 {
		goto L96
	} else {
		goto L97
	}
L67:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1))))
	if v193 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	if v1&int32(3) == int32(0) {
		v241 = v1
		goto L80
	} else {
		goto L81
	}
L70:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1)+1)))
	if base.Ui32((v197-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v279 = int32(6)
		goto L66
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v211 = int32(1)
	if v193&v211 != 0 {
		v279 = int32(base.Ui32(v193) >> (uint(v211) % 32))
		goto L66
	} else {
		goto L77
	}
L73:
	;
	v204 = int32(18)
	if v197&int32(255) == v204 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v210 = v204
	goto L76
L75:
	;
	v210 = int32(2)
	goto L76
L76:
	;
	v279 = v210
	goto L66
L77:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v1)))
	v279 = int32(base.Ui32(v215) >> (uint(int32(2)) % 32))
	goto L66
L78:
	;
	v279 = v274 + int32(1)
	goto L66
L79:
	;
	v274 = v266 - v1
	goto L78
L80:
	;
	v245 = v241
	goto L89
L81:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1))))
	if v225 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v274 = int32(0)
	goto L78
L83:
	;
	goto L84
L84:
	;
	v230 = v1
	goto L85
L85:
	;
	v234 = v230 + int32(1)
	if v234&int32(3) == int32(0) {
		v241 = v234
		goto L80
	} else {
		goto L87
	}
L86:
	;
	v266 = v234
	goto L79
L87:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v239 != 0 {
		v230 = v234
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v254 = int32(-2139062144)
	if (int32(16843008)-v251|v251)&v254 == v254 {
		v245 = v245 + int32(4)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v260 = v245
	goto L92
L91:
	;
	goto L90
L92:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v264 != 0 {
		v260 = v260 + int32(1)
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v266 = v260
	goto L79
L94:
	;
	goto L93
L95:
	;
	switch l3 - int32(99) {
	case 0:
		v438 = v279
		goto L1
	case 1:
		goto L142
	default:
		goto L141
	case 6:
		goto L143
	}
L96:
	;
	goto L95
L97:
	;
	v283 = l4 + v279
	if base.Ui32(v1-v283) <= base.Ui32(int32(0)-v279<<(uint(int32(1))%32)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v290 = F___memcpy(m, l4, v1, v279)
	mBase = m.M
	goto L95
L99:
	;
	goto L100
L100:
	;
	v293 = (l4 ^ v1) & int32(3)
	if base.Ui32(l4) < base.Ui32(v1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	if v395 == int32(0) {
		goto L96
	} else {
		goto L137
	}
L102:
	;
	if base.Ui32(v373) <= base.Ui32(int32(3)) {
		v394 = v372
		v395 = v373
		v396 = v374
		goto L101
	} else {
		goto L133
	}
L103:
	;
	if v293 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if v293 != 0 {
		v355 = v279
		goto L116
	} else {
		goto L117
	}
L106:
	;
	v394 = v1
	v395 = v279
	v396 = l4
	goto L101
L107:
	;
	goto L108
L108:
	;
	if l4&int32(3) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v372 = v1
	v373 = v279
	v374 = l4
	goto L102
L110:
	;
	goto L111
L111:
	;
	v300 = v1
	v301 = v279
	v302 = l4
	goto L112
L112:
	;
	if v301 == int32(0) {
		goto L96
	} else {
		goto L114
	}
L113:
	;
	v372 = v309
	v373 = v311
	v374 = v313
	goto L102
L114:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	*(*uint8)(unsafe.Add(mBase, uint32(v302))) = uint8(v306)
	v308 = int32(1)
	v309 = v300 + v308
	v311 = v301 - v308
	v313 = v302 + v308
	if v313&int32(3) != 0 {
		v300 = v309
		v301 = v311
		v302 = v313
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	if v355 == int32(0) {
		goto L96
	} else {
		goto L129
	}
L117:
	;
	if v283&int32(3) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v320 = v279
	goto L121
L119:
	;
	v335 = v279
	goto L120
L120:
	;
	if base.Ui32(v335) <= base.Ui32(int32(3)) {
		v355 = v335
		goto L116
	} else {
		goto L125
	}
L121:
	;
	if v320 == int32(0) {
		goto L96
	} else {
		goto L123
	}
L122:
	;
	v335 = v326
	goto L120
L123:
	;
	v326 = v320 - int32(1)
	v327 = l4 + v326
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1+v326))))
	*(*uint8)(unsafe.Add(mBase, uint32(v327))) = uint8(v329)
	if v327&int32(3) != 0 {
		v320 = v326
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v342 = v335
	goto L126
L126:
	;
	v346 = v342 - int32(4)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v1+v346)))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v346))) = v349
	if base.Ui32(int32(3)) < base.Ui32(v346) {
		v342 = v346
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v355 = v346
	goto L116
L128:
	;
	goto L127
L129:
	;
	v362 = v355
	goto L130
L130:
	;
	v366 = v362 - int32(1)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1+v366))))
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v366))) = uint8(v369)
	if v366 != 0 {
		v362 = v366
		goto L130
	} else {
		goto L132
	}
L131:
	;
	goto L96
L132:
	;
	goto L131
L133:
	;
	v379 = v372
	v380 = v373
	v381 = v374
	goto L134
L134:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v383
	v385 = int32(4)
	v386 = v379 + v385
	v388 = v381 + v385
	v390 = v380 - v385
	if base.Ui32(int32(3)) < base.Ui32(v390) {
		v379 = v386
		v380 = v390
		v381 = v388
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v394 = v386
	v395 = v390
	v396 = v388
	goto L101
L136:
	;
	goto L135
L137:
	;
	v401 = v394
	v402 = v395
	v403 = v396
	goto L138
L138:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v405)
	v407 = int32(1)
	v412 = v402 - v407
	if v412 != 0 {
		v401 = v401 + v407
		v402 = v412
		v403 = v403 + v407
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L96
L140:
	;
	goto L139
L141:
	;
	v438 = (v279 + int32(1)) & int32(-2)
	goto L1
L142:
	;
	v438 = (v279 + int32(7)) & int32(-8)
	goto L1
L143:
	;
	v438 = (v279 + int32(3)) & int32(-4)
	goto L1
}
func F_array_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	v7 = int32(0)
	if l3 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_0))
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_1))
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	if int32(0) < l3 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L10:
	;
	m.T0[l4].(func(*base.Module, int32, int32, int32))(m, l0, l1+l2*v27, l5)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if v27 < l3-int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_2))
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v38 = v27 + int32(1)
	if v38 != l3 {
		v27 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L11
L18:
	;
	return
}
func F_array_dims(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_DatumGetAnyArrayP(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(224)
	return v151
L2:
	;
	return int32(0)
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v20 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = int32(28)
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14+v23)))
	if base.Ui32(v25-int32(7)) <= base.Ui32(int32(-7)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
	v151 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	if v20 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v48 = v11 + int32(16)
	v51 = int32(0)
	goto L14
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v43 = v35
	v44 = v36
	goto L10
L12:
	;
	goto L13
L13:
	;
	v38 = v14 + int32(16)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v43 = v38
	v44 = v38 + v39<<(uint(int32(2))%32)
	goto L10
L14:
	;
	v57 = v51 << (uint(int32(2)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43+v57)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v61 + v59 - int32(1)
	v68 = F_pg_sprintf(m, v48, int32(_a_F_array_dims_0), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v141 = F_cstring_to_text(m, v11+int32(16))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L38
	}
L16:
	;
	if v48&int32(3) == int32(0) {
		v93 = v48
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v129 = v51 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v132 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L18:
	;
	v126 = v118 - v48
	goto L17
L19:
	;
	v97 = v93
	goto L28
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v77 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v126 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v82 = v48
	goto L24
L24:
	;
	v86 = v82 + int32(1)
	if v86&int32(3) == int32(0) {
		v93 = v86
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v118 = v86
	goto L18
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v91 != 0 {
		v82 = v86
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v106 = int32(-2139062144)
	if (int32(16843008)-v103|v103)&v106 == v106 {
		v97 = v97 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v112 = v97
	goto L31
L30:
	;
	goto L29
L31:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 != 0 {
		v112 = v112 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v118 = v112
	goto L18
L33:
	;
	goto L32
L34:
	;
	v135 = int32(28)
	goto L36
L35:
	;
	v135 = int32(4)
	goto L36
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v14+v135)))
	if v129 < v137 {
		v48 = v126 + v48
		v51 = v129
		goto L14
	} else {
		goto L37
	}
L37:
	;
	goto L15
L38:
	;
	v151 = v141
	goto L1
}
func F_array_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v91 int64
	_ = v91
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
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
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v846 int32
	_ = v846
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v978 int32
	_ = v978
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v1001 int32
	_ = v1001
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1252 int32
	_ = v1252
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1383 int32
	_ = v1383
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1419 int32
	_ = v1419
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1613 int32
	_ = v1613
	var v1623 int32
	_ = v1623
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1684 int32
	_ = v1684
	v2 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(576)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v35 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+8)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	v85 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+512)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v28)+504)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v28)+496)) = v85
	v91 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+480)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v28)+472)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v28)+464)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v28)+528)) = v33
	v101 = v33
	v105 = v2
	goto L11
L2:
	;
	F_get_type_io_data(m, v31, int32(0), v54+int32(4), v54+int32(6), v54+int32(7), v54+int32(8), v54+int32(12), v54+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L9
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v40 = F_MemoryContextAlloc(m, v38, int32(48))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v51 == v31 {
		v79 = v35
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v40
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v31 ^ int32(-1)
	v54 = v47
	goto L2
L8:
	;
	v54 = v35
	goto L2
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_fmgr_info_cxt(m, v70, v54+int32(20), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v31
	v79 = v54
	goto L1
L11:
	;
	v125 = v101 + int32(1)
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101))))
	goto L13
L12:
	;
	m.G0 = v28 + int32(576)
	return v1684
L13:
	;
	if base.B2i32(v126 == int32(32))|base.B2i32(base.Ui32((v126-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v101 = v125
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+528)) = v101
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v137 == int32(91) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(496)+v105<<(uint(int32(2))%32)))) = v307
	v101 = v266
	v105 = v105 + int32(1)
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+528)) = v125
	if v105 == int32(6) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L19
L19:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v105 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L20:
	;
	v1684 = int32(0)
	goto L15
L21:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), v328, int32(_a_F_array_in_1))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L83
	}
L22:
	;
	v144 = F_errsave_start(m, v32)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v161 = F_ReadDimensionInt(m, v28+int32(528), v28+int32(524), v32)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	if v144 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_2), v28)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v328 = int32(432)
	goto L21
L29:
	;
	if v161 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v28)+528))
	if v125 == v165 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v167 = F_errsave_start(m, v32)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v185 == int32(58) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	if v167 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errdetail(m, int32(_a_F_array_in_4), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v328 = int32(441)
	goto L21
L39:
	;
	if v238&int32(255) != int32(93) {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v28)+524))
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(464)+v105<<(uint(int32(2))%32)))) = v193
	v196 = v165 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+528)) = v196
	v202 = F_ReadDimensionInt(m, v28+int32(528), v28+int32(560), v32)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v227 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(464)+v105<<(uint(int32(2))%32)))) = v227
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v28)+524))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+560)) = v235
	v237 = v165
	v238 = v185
	v239 = v227
	goto L39
L43:
	;
	if v202 == int32(0) {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v28)+528))
	if v196 != v206 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v237 = v206
	v238 = v208
	v239 = v193
	goto L39
L46:
	;
	goto L47
L47:
	;
	v209 = F_errsave_start(m, v32)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v209 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errdetail(m, int32(_a_F_array_in_5), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v328 = int32(455)
	goto L21
L53:
	;
	v244 = F_errsave_start(m, v32)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v266 = v237 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+528)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v28)+560))
	if v268 < v239 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	if v244 == int32(0) {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(96))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = int32(_a_F_array_in_6)
	F_errdetail(m, int32(_a_F_array_in_7), v28+int32(80))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v328 = int32(468)
	goto L21
L61:
	;
	v270 = F_errsave_start(m, v32)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v268 == int32(2147483647) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v270 == int32(0) {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_array_in_8), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v328 = int32(481)
	goto L21
L68:
	;
	v284 = F_errsave_start(m, v32)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v299 = int32(0)
	v301 = v268 - v239
	if base.B2i32(v299 < v239)^base.B2i32(v301 < v268) == v299 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	if v284 == int32(0) {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = int32(2147483647)
	F_errmsg(m, int32(_a_F_array_in_9), v28+int32(48))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v328 = int32(487)
	goto L21
L75:
	;
	v307 = v301 + int32(1)
	if v301 <= v307 {
		goto L16
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v310 = F_errsave_start(m, v32)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v310 == int32(0) {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_in_10), v28-int32(-64))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v328 = int32(495)
	goto L21
L83:
	;
	goto L20
L84:
	;
	v490 = F_palloc(m, int32(64))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L6
	} else {
		goto L115
	}
L85:
	;
	if v336&int32(255) == int32(123) {
		v459 = v101
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v336&int32(255) == int32(61) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v343 = int32(0)
	v344 = F_errsave_start(m, v32)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	if v344 == int32(0) {
		v1684 = v343
		goto L15
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+400)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(400))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	F_errdetail(m, int32(_a_F_array_in_11), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(265), int32(_a_F_array_in_12))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v1684 = v343
	goto L15
L95:
	;
	v370 = v101
	goto L98
L96:
	;
	goto L97
L97:
	;
	v433 = int32(0)
	v434 = F_errsave_start(m, v32)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L109
	}
L98:
	;
	v396 = v370 + int32(1)
	v397 = int32(*(*int8)(unsafe.Add(mBase, uint32(v396))))
	goto L100
L99:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	if v407 == int32(123) {
		v459 = v396
		goto L84
	} else {
		goto L102
	}
L100:
	;
	if base.B2i32(v397 == int32(32))|base.B2i32(base.Ui32((v397-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v370 = v396
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v410 = int32(0)
	v411 = F_errsave_start(m, v32)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	if v411 == int32(0) {
		v1684 = v410
		goto L15
	} else {
		goto L104
	}
L104:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+416)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(416))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	F_errdetail(m, int32(_a_F_array_in_13), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(285), int32(_a_F_array_in_12))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v1684 = v410
	goto L15
L109:
	;
	if v434 == int32(0) {
		v1684 = v433
		goto L15
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+448)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(448))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+432)) = int32(_a_F_array_in_14)
	F_errdetail(m, int32(_a_F_array_in_7), v28+int32(432))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(275), int32(_a_F_array_in_12))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v1684 = v433
	goto L15
L115:
	;
	v493 = F_palloc(m, int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_initStringInfo(m, v28+int32(560))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v500 = v459
	v504 = int32(0)
	v505 = base.B2i32(v105 != int32(0))
	v508 = v2
	v509 = v2
	v511 = v105
	v513 = v490
	v514 = v493
	v516 = int32(16)
	goto L121
L118:
	;
	if v1206 != 0 {
		goto L327
	} else {
		goto L328
	}
L119:
	;
	v1684 = int32(0)
	goto L15
L120:
	;
	v1314 = F_errsave_start(m, v32)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L6
	} else {
		goto L313
	}
L121:
	;
	v526 = v28 + int32(560)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v528 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v527))) = uint8(v528)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+12)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v528
	goto L123
L122:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v28)+560))
	F_pfree(m, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L6
	} else {
		goto L301
	}
L123:
	;
	v534 = v500
	goto L129
L124:
	;
	if int32(0) < v1201 {
		v500 = v1197
		v504 = v1201
		v505 = v1202
		v508 = v1205
		v509 = v1206
		v511 = v1208
		v513 = v1210
		v514 = v1211
		v516 = v1213
		goto L121
	} else {
		goto L300
	}
L125:
	;
	if v508&int32(1) != 0 {
		goto L269
	} else {
		goto L270
	}
L126:
	;
	v1082 = v829
	v1087 = int32(0)
	goto L125
L127:
	;
	if v508&int32(1) != 0 {
		goto L260
	} else {
		goto L261
	}
L128:
	;
	v1028 = F_errsave_start(m, v32)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L6
	} else {
		goto L254
	}
L129:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	switch v559 - int32(123) {
	case 0:
		goto L133
	case 1:
		goto L131
	case 2:
		goto L132
	default:
		goto L134
	}
L130:
	;
	v829 = v534
	v834 = int32(0)
	v846 = int32(1)
	goto L205
L131:
	;
	v812 = base.I32_extend8_s(v559)
	if v812 == v81 {
		goto L127
	} else {
		goto L199
	}
L132:
	;
	v748 = int32(1)
	v749 = v504 - v748
	v751 = v749 << (uint(int32(2)) % 32)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v751+(v28+int32(528)))))
	v756 = int32(0)
	if (v508|base.B2i32(v755 <= v756))&v748 == v756 {
		goto L182
	} else {
		goto L183
	}
L133:
	;
	if v508&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L134:
	;
	if v559 == int32(0) {
		goto L128
	} else {
		goto L135
	}
L135:
	;
	if v559 != int32(34) {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v568 = v534 + int32(1)
	goto L137
L137:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v593 != int32(92) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v613 = v568
	goto L148
L139:
	;
	goto L138
L140:
	;
	F_appendStringInfoChar(m, v28+int32(560), base.I32_extend8_s(v605))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L147
	}
L141:
	;
	if v593 == int32(0) {
		goto L128
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+1)))
	if v601 == int32(0) {
		goto L128
	} else {
		goto L146
	}
L144:
	;
	if v593 == int32(34) {
		goto L139
	} else {
		goto L145
	}
L145:
	;
	v605 = v593
	v606 = int32(1)
	goto L140
L146:
	;
	v605 = v601
	v606 = int32(2)
	goto L140
L147:
	;
	v568 = v568 + v606
	goto L137
L148:
	;
	v639 = v613 + int32(1)
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	if v640 == int32(0) {
		goto L128
	} else {
		goto L150
	}
L149:
	;
	v659 = F_errsave_start(m, v32)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L155
	}
L150:
	;
	v643 = int32(0)
	v644 = base.I32_extend8_s(v640)
	if v644 == v81 {
		v1082 = v639
		v1087 = v643
		goto L125
	} else {
		goto L151
	}
L151:
	;
	switch v644&int32(255) - int32(123) {
	case 0, 2:
		v1082 = v639
		v1087 = v643
		goto L125
	default:
		goto L152
	}
L152:
	;
	goto L153
L153:
	;
	if base.B2i32(v644 == int32(32))|base.B2i32(base.Ui32((v644-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v613 = v639
		goto L148
	} else {
		goto L154
	}
L154:
	;
	goto L149
L155:
	;
	if v659 == int32(0) {
		goto L119
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+384)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(384))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	F_errdetail(m, int32(_a_F_array_in_15), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(871), int32(_a_F_array_in_16))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v1684 = int32(0)
	goto L15
L161:
	;
	v684 = F_errsave_start(m, v32)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(int32(6)) <= base.Ui32(v504) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	if v684 == int32(0) {
		goto L119
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(320))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = int32(123)
	F_errdetail(m, int32(_a_F_array_in_17), v28+int32(304))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(637), int32(_a_F_array_in_18))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	v1684 = int32(0)
	goto L15
L170:
	;
	v712 = F_errsave_start(m, v32)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L6
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v732 = int32(1)
	v733 = v534 + v732
	v734 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(528)+v504<<(uint(int32(2))%32)))) = v734
	v743 = v504 + v732
	if v504 < v511 {
		goto L178
	} else {
		goto L179
	}
L173:
	;
	if v712 == int32(0) {
		goto L119
	} else {
		goto L174
	}
L174:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_2), v28+int32(336))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(644), int32(_a_F_array_in_18))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v1684 = int32(0)
	goto L15
L178:
	;
	v1197 = v733
	v1201 = v743
	v1202 = v505
	v1205 = v734
	v1206 = v509
	v1208 = v511
	v1210 = v513
	v1211 = v514
	v1213 = v516
	goto L124
L179:
	;
	goto L180
L180:
	;
	if v505&int32(1) != 0 {
		goto L120
	} else {
		goto L181
	}
L181:
	;
	v1197 = v733
	v1201 = v743
	v1202 = int32(0)
	v1205 = v734
	v1206 = v509
	v1208 = v743
	v1210 = v513
	v1211 = v514
	v1213 = v516
	goto L124
L182:
	;
	v763 = F_errsave_start(m, v32)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L6
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v504) {
		goto L191
	} else {
		goto L192
	}
L185:
	;
	if v763 == int32(0) {
		goto L119
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+368)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(368))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+352)) = int32(125)
	F_errdetail(m, int32(_a_F_array_in_17), v28+int32(352))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(670), int32(_a_F_array_in_18))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v1684 = int32(0)
	goto L15
L191:
	;
	v795 = v504<<(uint(int32(2))%32) + v28 + int32(520)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	*(*int32)(unsafe.Add(mBase, uint32(v795))) = v796 + int32(1)
	goto L193
L192:
	;
	goto L193
L193:
	;
	v805 = v28 + int32(496) + v751
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	if v806 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1197 = v534 + int32(1)
	v1201 = v749
	v1202 = v505
	v1205 = int32(1)
	v1206 = v509
	v1208 = v511
	v1210 = v513
	v1211 = v514
	v1213 = v516
	goto L124
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v805))) = v755
	goto L194
L196:
	;
	goto L197
L197:
	;
	if v755 != v806 {
		goto L120
	} else {
		goto L198
	}
L198:
	;
	goto L194
L199:
	;
	goto L201
L200:
	;
	goto L130
L201:
	;
	if base.B2i32(v812 == int32(32))|base.B2i32(base.Ui32((v812-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	goto L200
L203:
	;
	goto L204
L204:
	;
	v534 = v534 + int32(1)
	goto L129
L205:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	if base.Ui32(v854) <= base.Ui32(int32(91)) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v28)+560))
	v949 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v947+v834))) = uint8(v949)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+564)) = v834
	v953 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_array_in[0])))
	if v846&v953&int32(1) == v949 {
		goto L126
	} else {
		goto L239
	}
L207:
	;
	v925 = base.I32_extend8_s(v854)
	if v925 == v81 {
		goto L231
	} else {
		goto L232
	}
L208:
	;
	if v854 == int32(0) {
		goto L128
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if v854 != int32(92) {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	if v854 != int32(34) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v861 = F_errsave_start(m, v32)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L6
	} else {
		goto L213
	}
L213:
	;
	if v861 == int32(0) {
		goto L119
	} else {
		goto L214
	}
L214:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+288)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(288))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	F_errdetail(m, int32(_a_F_array_in_15), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(905), int32(_a_F_array_in_16))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v1684 = int32(0)
	goto L15
L219:
	;
	if v854 != int32(123) {
		goto L207
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v914 = int32(*(*int8)(unsafe.Add(mBase, uint32(v829)+1)))
	if v914 == int32(0) {
		goto L128
	} else {
		goto L229
	}
L222:
	;
	v888 = F_errsave_start(m, v32)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L223
	}
L223:
	;
	if v888 == int32(0) {
		goto L119
	} else {
		goto L224
	}
L224:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(272))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = int32(123)
	F_errdetail(m, int32(_a_F_array_in_17), v28+int32(256))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(899), int32(_a_F_array_in_16))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	v1684 = int32(0)
	goto L15
L229:
	;
	F_appendStringInfoChar(m, v28+int32(560), v914)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v28)+564))
	v829 = v829 + int32(2)
	v834 = v924
	v846 = int32(0)
	goto L205
L231:
	;
	goto L206
L232:
	;
	if v925 == int32(125) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	F_appendStringInfoChar(m, v28+int32(560), v925)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	v933 = int32(*(*int8)(unsafe.Add(mBase, uint32(v829))))
	goto L235
L235:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v28)+564))
	if base.B2i32(v933 == int32(32))|base.B2i32(base.Ui32((v933-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v944 = v834
	goto L238
L237:
	;
	v944 = v943
	goto L238
L238:
	;
	v829 = v829 + int32(1)
	v834 = v944
	goto L205
L239:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v28)+560))
	v963 = v959
	v964 = int32(_a_F_array_in_19)
	goto L241
L240:
	;
	if v1001 != 0 {
		goto L126
	} else {
		goto L253
	}
L241:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963))))
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964))))
	if v967 == v968 {
		v990 = v967
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1001 = int32(0)
	goto L240
L243:
	;
	v992 = int32(1)
	if v990 != 0 {
		v963 = v963 + v992
		v964 = v964 + v992
		goto L241
	} else {
		goto L252
	}
L244:
	;
	if base.Ui32((v967-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v978 = v967 | int32(32)
	goto L247
L246:
	;
	v978 = v967
	goto L247
L247:
	;
	if base.Ui32((v968-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v987 = v968 | int32(32)
	goto L250
L249:
	;
	v987 = v968
	goto L250
L250:
	;
	if v978 == v987 {
		v990 = v978
		goto L243
	} else {
		goto L251
	}
L251:
	;
	v1001 = v978 - v987
	goto L240
L252:
	;
	goto L242
L253:
	;
	v1082 = v829
	v1087 = int32(1)
	goto L125
L254:
	;
	if v1028 == int32(0) {
		goto L119
	} else {
		goto L255
	}
L255:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L6
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+240)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(240))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L6
	} else {
		goto L257
	}
L257:
	;
	F_errdetail(m, int32(_a_F_array_in_20), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L6
	} else {
		goto L258
	}
L258:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(942), int32(_a_F_array_in_16))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L6
	} else {
		goto L259
	}
L259:
	;
	v1684 = int32(0)
	goto L15
L260:
	;
	v1197 = v534 + int32(1)
	v1201 = v504
	v1202 = v505
	v1205 = int32(0)
	v1206 = v509
	v1208 = v511
	v1210 = v513
	v1211 = v514
	v1213 = v516
	goto L124
L261:
	;
	goto L262
L262:
	;
	v1056 = F_errsave_start(m, v32)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	if v1056 == int32(0) {
		goto L119
	} else {
		goto L264
	}
L264:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L6
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(160))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L6
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = v81
	F_errdetail(m, int32(_a_F_array_in_17), v28+int32(144))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(705), int32(_a_F_array_in_18))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L6
	} else {
		goto L268
	}
L268:
	;
	v1684 = int32(0)
	goto L15
L269:
	;
	v1109 = F_errsave_start(m, v32)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L6
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	if v516 <= v509 {
		goto L278
	} else {
		goto L279
	}
L272:
	;
	if v1109 == int32(0) {
		goto L119
	} else {
		goto L273
	}
L273:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L6
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(176))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L6
	} else {
		goto L275
	}
L275:
	;
	F_errdetail(m, int32(_a_F_array_in_21), int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L6
	} else {
		goto L276
	}
L276:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(719), int32(_a_F_array_in_18))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L6
	} else {
		goto L277
	}
L277:
	;
	v1684 = int32(0)
	goto L15
L278:
	;
	if base.Ui32(int32(268435455)) <= base.Ui32(v516) {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1168 = v513
	v1169 = v514
	v1170 = v516
	goto L280
L280:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v28)+560))
	if v1087 != 0 {
		goto L294
	} else {
		goto L295
	}
L281:
	;
	v1135 = F_errsave_start(m, v32)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L6
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1155 = int32(268435455)
	v1157 = v516 << (uint(int32(1)) % 32)
	if base.Ui32(v1155) <= base.Ui32(v1157) {
		goto L289
	} else {
		goto L290
	}
L284:
	;
	if v1135 == int32(0) {
		goto L119
	} else {
		goto L285
	}
L285:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L6
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+224)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_in_10), v28+int32(224))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L6
	} else {
		goto L287
	}
L287:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(728), int32(_a_F_array_in_18))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L6
	} else {
		goto L288
	}
L288:
	;
	v1684 = int32(0)
	goto L15
L289:
	;
	v1160 = v1155
	goto L291
L290:
	;
	v1160 = v1157
	goto L291
L291:
	;
	v1163 = F_repalloc(m, v513, v1160<<(uint(int32(2))%32))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	v1165 = F_repalloc(m, v514, v1160)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L6
	} else {
		goto L293
	}
L293:
	;
	v1168 = v1163
	v1169 = v1165
	v1170 = v1160
	goto L280
L294:
	;
	v1173 = int32(0)
	goto L296
L295:
	;
	v1173 = v1172
	goto L296
L296:
	;
	v1177 = F_InputFunctionCallSafe(m, v79+int32(20), v1173, v80, v30, v32, v1168+v509<<(uint(int32(2))%32))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L6
	} else {
		goto L297
	}
L297:
	;
	if v1177 == int32(0) {
		goto L119
	} else {
		goto L298
	}
L298:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v509+v1169))) = uint8(v1087)
	if v504 != v511 {
		goto L120
	} else {
		goto L299
	}
L299:
	;
	v1184 = int32(1)
	v1189 = v504<<(uint(int32(2))%32) + v28 + int32(524)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	*(*int32)(unsafe.Add(mBase, uint32(v1189))) = v1190 + v1184
	v1197 = v1082
	v1201 = v504
	v1202 = v1184
	v1205 = v1184
	v1206 = v509 + v1184
	v1208 = v504
	v1210 = v1168
	v1211 = v1169
	v1213 = v1170
	goto L124
L300:
	;
	goto L122
L301:
	;
	v1227 = v1197
	goto L302
L302:
	;
	v1252 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1227))))
	if v1252 == int32(0) {
		goto L118
	} else {
		goto L304
	}
L303:
	;
	v1266 = int32(0)
	v1267 = F_errsave_start(m, v32)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L6
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	if base.B2i32(v1252 == int32(32))|base.B2i32(base.Ui32((v1252-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v1227 = v1227 + int32(1)
		goto L302
	} else {
		goto L306
	}
L306:
	;
	goto L303
L307:
	;
	if v1267 == int32(0) {
		v1684 = v1266
		goto L15
	} else {
		goto L308
	}
L308:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L6
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(128))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L6
	} else {
		goto L310
	}
L310:
	;
	F_errdetail(m, int32(_a_F_array_in_22), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L6
	} else {
		goto L311
	}
L311:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(308), int32(_a_F_array_in_12))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L6
	} else {
		goto L312
	}
L312:
	;
	v1684 = v1266
	goto L15
L313:
	;
	if v105 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	if v1314 == int32(0) {
		goto L119
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	if v1314 == int32(0) {
		goto L119
	} else {
		goto L322
	}
L317:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L6
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(192))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L6
	} else {
		goto L319
	}
L319:
	;
	F_errdetail(m, int32(_a_F_array_in_23), int32(0))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L6
	} else {
		goto L320
	}
L320:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(778), int32(_a_F_array_in_18))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L6
	} else {
		goto L321
	}
L321:
	;
	v1684 = int32(0)
	goto L15
L322:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L6
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v33
	F_errmsg(m, int32(_a_F_array_in_3), v28+int32(208))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L6
	} else {
		goto L324
	}
L324:
	;
	F_errdetail(m, int32(_a_F_array_in_24), int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L6
	} else {
		goto L325
	}
L325:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(783), int32(_a_F_array_in_18))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L6
	} else {
		goto L326
	}
L326:
	;
	goto L119
L327:
	;
	v1383 = int32(0)
	if v1206 <= v1383 {
		v1582 = v1383
		goto L331
	} else {
		goto L332
	}
L328:
	;
	goto L329
L329:
	;
	v1669 = F_palloc0(m, int32(16))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L6
	} else {
		goto L396
	}
L330:
	;
	v1637 = v1636 + v1613
	v1638 = F_palloc0(m, v1637)
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L6
	} else {
		goto L384
	}
L331:
	;
	v1613 = v1582
	v1623 = v1383
	v1636 = (v1208<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L330
L332:
	;
	v1391 = int32(0)
	v1393 = v1391
	v1395 = v1383
	v1397 = v1391
	goto L333
L333:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393+v1211))))
	if v1419 != 0 {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	if v1559&int32(1) == int32(0) {
		v1582 = v1558
		goto L331
	} else {
		goto L383
	}
L335:
	;
	v1563 = v1393 + int32(1)
	if v1563 != v1206 {
		v1393 = v1563
		v1395 = v1558
		v1397 = v1559
		goto L333
	} else {
		goto L382
	}
L336:
	;
	v1558 = v1395
	v1559 = int32(1)
	goto L335
L337:
	;
	goto L338
L338:
	;
	if base.B2i32(v83 == int32(-1)) == int32(0) {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v1522 = v1395 + v1520
	switch v84 - int32(99) {
	case 0:
		v1535 = v1522
		goto L372
	case 1:
		goto L374
	default:
		goto L373
	case 6:
		goto L375
	}
L340:
	;
	if int32(0) < v83 {
		v1520 = v83
		goto L339
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v1490 = v1210 + v1393<<(uint(int32(2))%32)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)))
	v1492 = F_pg_detoast_datum(m, v1491)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L6
	} else {
		goto L361
	}
L343:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1210+v1393<<(uint(int32(2))%32))))
	if v1428&int32(3) == int32(0) {
		v1452 = v1428
		goto L346
	} else {
		goto L347
	}
L344:
	;
	v1520 = v1485 + int32(1)
	goto L339
L345:
	;
	v1485 = v1477 - v1428
	goto L344
L346:
	;
	v1456 = v1452
	goto L355
L347:
	;
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1428))))
	if v1436 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1485 = int32(0)
	goto L344
L349:
	;
	goto L350
L350:
	;
	v1441 = v1428
	goto L351
L351:
	;
	v1445 = v1441 + int32(1)
	if v1445&int32(3) == int32(0) {
		v1452 = v1445
		goto L346
	} else {
		goto L353
	}
L352:
	;
	v1477 = v1445
	goto L345
L353:
	;
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1445))))
	if v1450 != 0 {
		v1441 = v1445
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	v1465 = int32(-2139062144)
	if (int32(16843008)-v1462|v1462)&v1465 == v1465 {
		v1456 = v1456 + int32(4)
		goto L355
	} else {
		goto L357
	}
L356:
	;
	v1471 = v1456
	goto L358
L357:
	;
	goto L356
L358:
	;
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471))))
	if v1475 != 0 {
		v1471 = v1471 + int32(1)
		goto L358
	} else {
		goto L360
	}
L359:
	;
	v1477 = v1471
	goto L345
L360:
	;
	goto L359
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490))) = v1492
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492))))
	if v1495 == int32(1) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492)+1)))
	if base.Ui32((v1499-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v1520 = int32(6)
		goto L339
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	if v1495&int32(1) != 0 {
		goto L369
	} else {
		goto L370
	}
L365:
	;
	v1506 = int32(18)
	if v1499&int32(255) == v1506 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1512 = v1506
	goto L368
L367:
	;
	v1512 = int32(2)
	goto L368
L368:
	;
	v1520 = v1512
	goto L339
L369:
	;
	v1520 = int32(base.Ui32(v1495) >> (uint(int32(1)) % 32))
	goto L339
L370:
	;
	goto L371
L371:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1492)))
	v1520 = int32(base.Ui32(v1517) >> (uint(int32(2)) % 32))
	goto L339
L372:
	;
	if base.Ui32(v1535) < base.Ui32(int32(1073741824)) {
		v1558 = v1535
		v1559 = v1397
		goto L335
	} else {
		goto L376
	}
L373:
	;
	v1535 = (v1522 + int32(1)) & int32(-2)
	goto L372
L374:
	;
	v1535 = (v1522 + int32(7)) & int32(-8)
	goto L372
L375:
	;
	v1535 = (v1522 + int32(3)) & int32(-4)
	goto L372
L376:
	;
	v1538 = int32(0)
	v1539 = F_errsave_start(m, v32)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L6
	} else {
		goto L377
	}
L377:
	;
	if v1539 == int32(0) {
		v1684 = v1538
		goto L15
	} else {
		goto L378
	}
L378:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L6
	} else {
		goto L379
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_in_10), v28+int32(112))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L6
	} else {
		goto L380
	}
L380:
	;
	F_errsave_finish(m, v32, int32(_a_F_array_in_0), int32(336), int32(_a_F_array_in_12))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L6
	} else {
		goto L381
	}
L381:
	;
	v1684 = v1538
	goto L15
L382:
	;
	goto L334
L383:
	;
	v1572 = base.I32_div_s(v1206+int32(7), int32(8))
	v1579 = (v1572 + v1208<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v1613 = v1558
	v1623 = v1579
	v1636 = v1579
	goto L330
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1638)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1638)+8)) = v1623
	*(*int32)(unsafe.Add(mBase, uint32(v1638)+4)) = v1208
	v1643 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1638))) = v1637 << (uint(v1643) % 32)
	v1647 = v1638 + int32(16)
	v1651 = v1208 << (uint(v1643) % 32)
	if v1651 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	if v1651 != 0 {
		goto L390
	} else {
		goto L391
	}
L386:
	;
	v1652 = F__emscripten_memcpy_bulkmem(m, v1647, v28+int32(496), v1651)
	mBase = m.M
	v1653 = v1652
	goto L388
L387:
	;
	v1653 = v1647
	goto L388
L388:
	;
	goto L385
L389:
	;
	v1659 = int32(1)
	F_CopyArrayEls(m, v1638, v1210, v1211, v1206, v83, v82&v1659, base.I32_extend8_s(v84), v1659)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L6
	} else {
		goto L393
	}
L390:
	;
	v1657 = F__emscripten_memcpy_bulkmem(m, v1653+v1651, v28+int32(464), v1651)
	mBase = m.M
	goto L392
L391:
	;
	goto L392
L392:
	;
	goto L389
L393:
	;
	F_pfree(m, v1210)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L6
	} else {
		goto L394
	}
L394:
	;
	F_pfree(m, v1211)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L6
	} else {
		goto L395
	}
L395:
	;
	v1684 = v1638
	goto L15
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1669))) = int64(64)
	v1684 = v1669
	goto L15
}
func F_array_ndims(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_DatumGetAnyArrayP(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v10 == int32(-1) {
			v13 = int32(28)
		} else {
			v13 = int32(4)
		}
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4+v13)))
		if base.Ui32(v15-int32(7)) <= base.Ui32(int32(-7)) {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v23 = int32(0)
		} else {
			v23 = v15
		}
		return v23
	}
}
func F_array_ne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2 ^ int32(1)
	}
}
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v428 int32
	_ = v428
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1268 int32
	_ = v1268
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1309 int32
	_ = v1309
	var v1316 int32
	_ = v1316
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1659 int32
	_ = v1659
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1816 int32
	_ = v1816
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1962 int32
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2104 int32
	_ = v2104
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2192 int32
	_ = v2192
	var v2202 int32
	_ = v2202
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2232 int32
	_ = v2232
	var v2239 int32
	_ = v2239
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2277 int32
	_ = v2277
	var v2286 int32
	_ = v2286
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2352 int32
	_ = v2352
	var v2361 int32
	_ = v2361
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2386 int32
	_ = v2386
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2461 int32
	_ = v2461
	var v2488 int32
	_ = v2488
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2523 int32
	_ = v2523
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2565 int32
	_ = v2565
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2597 int32
	_ = v2597
	var v2602 int32
	_ = v2602
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2618 int32
	_ = v2618
	var v2623 int32
	_ = v2623
	v5 = l4
	v10 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(128)
	m.G0 = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+59)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = l3
	if v10 < l5 {
		goto L17
	} else {
		goto L18
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2608 = m.ExcPending
	if v2608 != 0 {
		goto L24
	} else {
		goto L605
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L24
	} else {
		goto L601
	}
L3:
	;
	m.G0 = v28 + int32(128)
	return v2565
L4:
	;
	v1726 = v28 + int32(96)
	v1728 = v28 - int32(-64)
	v1729 = int32(0)
	v1738 = l1 - int32(1)
	if v1738 < v1729 {
		v1816 = v1729
		goto L421
	} else {
		goto L422
	}
L5:
	;
	v1690 = F_ArrayGetNItems(m, l1, v28+int32(96))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L24
	} else {
		goto L418
	}
L6:
	;
	v1659 = int32(0)
	if v1647 == v1659 {
		v1706 = v1659
		v1709 = v1644
		v1713 = v1648
		v1724 = v1659
		goto L4
	} else {
		goto L417
	}
L7:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v1590 <= v1589 {
		goto L408
	} else {
		goto L409
	}
L8:
	;
	v648 = F_ArrayGetNItems(m, l1, v28+int32(96))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L24
	} else {
		goto L165
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L24
	} else {
		goto L161
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L24
	} else {
		goto L157
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L24
	} else {
		goto L153
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L24
	} else {
		goto L149
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L24
	} else {
		goto L145
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L24
	} else {
		goto L141
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L24
	} else {
		goto L137
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L24
	} else {
		goto L133
	}
L17:
	;
	if l1 != int32(1) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(l1-int32(7)) <= base.Ui32(int32(-7)) {
		goto L13
	} else {
		goto L31
	}
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v36 < int32(0) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v39 = base.I32_div_s(l5, l6)
	if v39 <= v36 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if v5 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v41 = F_palloc(m, l5)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	if l5 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v50 = F_ArrayCastAndSet(m, l3, l6, l7, l8, v46+v47*l6)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v45 = F__emscripten_memcpy_bulkmem(m, v41, l0, l5)
	mBase = m.M
	v46 = v45
	goto L29
L28:
	;
	v46 = v41
	goto L29
L29:
	;
	goto L26
L30:
	;
	v2565 = v41
	goto L3
L31:
	;
	if v5 != 0 {
		v61 = l3
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v62 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if l6 != int32(-1) {
		v61 = l3
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v58 = F_pg_detoast_datum(m, l3)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v58
	v61 = v58
	goto L32
L36:
	;
	v196 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L24
	} else {
		goto L78
	}
L37:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v65&int32(254) != int32(2) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v72 = F_DatumGetExpandedArray(m, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v77 = v75 << (uint(int32(2)) % 32)
	if v77 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72)+36))
	if v77 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v78 = F__emscripten_memcpy_bulkmem(m, v28+int32(96), v74, v77)
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	if v75 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, v28-int32(-64), v82, v77)
	mBase = m.M
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	F_deconstruct_expanded_array(m, v72)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L60
	}
L49:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v89 = l1 << (uint(int32(2)) % 32)
	v90 = F_MemoryContextAllocZero(m, v87, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L24
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v75 != l1 {
		goto L12
	} else {
		goto L59
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = v90
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v94 = F_MemoryContextAllocZero(m, v93, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+36)) = v94
	v101 = F__emscripten_memset_bulkmem(m, v28+int32(96), base.I32_extend8_s(int32(0)), v89)
	mBase = m.M
	goto L54
L54:
	;
	if v89 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L48
L56:
	;
	v104 = F__emscripten_memcpy_bulkmem(m, v28-int32(-64), l2, v89)
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L48
L60:
	;
	if v5 != 0 {
		v124 = v61
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v126 = int32(0)
	v127 = base.B2i32(v75 == v126)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	v132 = base.B2i32(v129 != v126) | v5
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v72)+48))
	if l1 == int32(1) {
		goto L7
	} else {
		goto L65
	}
L62:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+46)))
	if v110&int32(1) != 0 {
		v124 = v61
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v113 = int32(_a_F_array_set_element_0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v116
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+44)))
	v120 = F_datumCopy(m, v61, int32(0), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L24
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v114
	v124 = v120
	goto L61
L65:
	;
	v146 = v126
	goto L66
L66:
	;
	v162 = v146 << (uint(int32(2)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2+v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v28-int32(-64)+v162)))
	if v168 <= v164 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v1644 = int32(0)
	v1647 = v127
	v1648 = v132
	goto L6
L68:
	;
	v193 = v146 + int32(1)
	if v193 != l1 {
		v146 = v193
		goto L66
	} else {
		goto L77
	}
L69:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(96)+v162)))
	if v164 < v173+v168 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L24
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2653), int32(_a_F_array_set_element_3))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L24
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	goto L67
L78:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v198 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v205 = l1 << (uint(int32(2)) % 32)
	if v205 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L81
L81:
	;
	if l1 != v198 {
		goto L11
	} else {
		goto L99
	}
L82:
	;
	v208 = int32(0)
	if base.Ui32(int32(7)) <= base.Ui32(l1-int32(1)) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v206 = F__emscripten_memcpy_bulkmem(m, v28-int32(-64), l2, v205)
	mBase = m.M
	goto L85
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	v219 = v208
	v224 = v10
	goto L89
L87:
	;
	v261 = v208
	goto L88
L88:
	;
	if l1 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v243 = v28 + int32(96) + v219<<(uint(int32(2))%32)
	v244 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v243))) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v243)+8)) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v243)+16)) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v243)+24)) = v244
	v252 = int32(8)
	v253 = v219 + v252
	v255 = v224 + v252
	if v255 != 0 {
		v219 = v253
		v224 = v255
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v261 = v253
	goto L88
L91:
	;
	goto L90
L92:
	;
	v283 = v208
	v286 = v261
	goto L95
L93:
	;
	goto L94
L94:
	;
	v351 = F_construct_md_array(m, v28+int32(60), v28+int32(59), l1, v28+int32(96), v28-int32(-64), v201, l6, l7, l8)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L24
	} else {
		goto L98
	}
L95:
	;
	v311 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(96)+v286<<(uint(int32(2))%32)))) = v311
	v316 = v283 + v311
	if v316 != l1 {
		v283 = v316
		v286 = v286 + v311
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	goto L96
L98:
	;
	v2565 = v351
	goto L3
L99:
	;
	v357 = v196 + int32(16)
	v359 = l1 << (uint(int32(2)) % 32)
	if v359 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v359 != 0 {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v360 = F__emscripten_memcpy_bulkmem(m, v28+int32(96), v357, v359)
	mBase = m.M
	goto L103
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	v370 = int32(0)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v374 = base.B2i32(v371 != v370) | v5
	if l1 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v368 = F__emscripten_memcpy_bulkmem(m, v28-int32(-64), v357+v364<<(uint(int32(2))%32), v359)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v378 <= v377 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	v428 = v370
	goto L124
L111:
	;
	v401 = v397 + v398
	if v377 < v401 {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v397 = v378
	v398 = v380
	v399 = v10
	v400 = v374
	goto L111
L113:
	;
	goto L114
L114:
	;
	v381 = v378 - v377
	if base.B2i32(v381 < v378)^base.B2i32(int32(0) < v377) != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v387 = v386 + v381
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v387
	if base.B2i32(v381 < int32(0)) != base.B2i32(v387 < v386) {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v377
	v397 = v377
	v398 = v387
	v399 = v381
	v400 = base.B2i32(int32(1) < v381) | v374
	goto L111
L117:
	;
	v621 = int32(1)
	v639 = v399
	v640 = v400
	goto L8
L118:
	;
	goto L119
L119:
	;
	v406 = v377 - v401
	if base.B2i32(int32(0) < v401)^base.B2i32(v406 < v377) != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	v410 = v406 + int32(1)
	if v410 < v406 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	v412 = v410 + v398
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v412
	if base.B2i32(v410 < int32(0)) != base.B2i32(v412 < v398) {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	v621 = base.B2i32(v410 == int32(0))
	v639 = v399
	v640 = base.B2i32(int32(1) < v410) | v400
	goto L8
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L24
	} else {
		goto L129
	}
L124:
	;
	v449 = v428 << (uint(int32(2)) % 32)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l2+v449)))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v28-int32(-64)+v449)))
	if v451 < v455 {
		goto L123
	} else {
		goto L126
	}
L125:
	;
	v621 = v463
	v639 = v10
	v640 = v374
	goto L8
L126:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(96)+v449)))
	if v460+v455 <= v451 {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v463 = int32(1)
	v465 = v428 + v463
	if v465 != l1 {
		v428 = v465
		goto L124
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L24
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L24
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2374), int32(_a_F_array_set_element_4))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L24
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L24
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L24
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2245), int32(_a_F_array_set_element_4))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L24
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L24
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2250), int32(_a_F_array_set_element_4))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	F_errmsg(m, int32(_a_F_array_set_element_6), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L24
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2255), int32(_a_F_array_set_element_4))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L24
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L24
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L24
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2267), int32(_a_F_array_set_element_4))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L24
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L24
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2570), int32(_a_F_array_set_element_3))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L24
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L24
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L24
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2316), int32(_a_F_array_set_element_4))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L24
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L24
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v28+int32(32))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L24
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2342), int32(_a_F_array_set_element_4))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L24
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L24
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v28+int32(48))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L24
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2357), int32(_a_F_array_set_element_4))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L24
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_ArrayCheckBounds(m, l1, v28+int32(96), v28-int32(-64))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L24
	} else {
		goto L166
	}
L166:
	;
	v657 = l1 << (uint(int32(3)) % 32)
	if v640&int32(1) != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v675 = F_ArrayGetNItems(m, l1, v357)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L24
	} else {
		goto L171
	}
L168:
	;
	v663 = base.I32_div_s(v648+int32(7), int32(8))
	v668 = (v657 + v663 + int32(23)) & int32(-8)
	v673 = v668
	v674 = v668
	goto L167
L169:
	;
	goto L170
L170:
	;
	v673 = v10
	v674 = (v657 + int32(23)) & int32(120)
	goto L167
L171:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v683 = v681 << (uint(int32(3)) % 32)
	if v680 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v688 = v680
	goto L174
L173:
	;
	v688 = (v683 + int32(23)) & int32(-8)
	goto L174
L174:
	;
	v689 = int32(base.Ui32(v677)>>(uint(int32(2))%32)) - v688
	if v680 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v692 = v683 + v357
	goto L177
L176:
	;
	v692 = int32(0)
	goto L177
L177:
	;
	if v639 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v5 != 0 {
		v1042 = v10
		goto L242
	} else {
		goto L243
	}
L179:
	;
	v932 = int32(0)
	v934 = v10
	v935 = v10
	v936 = v689
	goto L178
L180:
	;
	goto L181
L181:
	;
	if v621 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v932 = v689
	v934 = v675
	v935 = v10
	v936 = int32(0)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v698 = v28 + int32(96)
	v700 = v28 - int32(-64)
	v701 = int32(0)
	v710 = l1 - int32(1)
	if v710 < v701 {
		v788 = v701
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v796 != 0 {
		goto L196
	} else {
		goto L197
	}
L186:
	;
	goto L185
L187:
	;
	v713 = int32(1)
	if v710 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if l1&v713 == int32(0) {
		v788 = v765
		goto L186
	} else {
		goto L195
	}
L189:
	;
	v765 = v701
	v766 = v710
	v767 = v713
	goto L188
L190:
	;
	goto L191
L191:
	;
	v724 = v701
	v725 = v710
	v726 = v713
	v727 = v701
	goto L192
L192:
	;
	v732 = int32(2)
	v733 = v725 << (uint(v732) % 32)
	v735 = v733 - int32(4)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l2+v735)))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v700+v735)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v733+v698)))
	v743 = v742 * v726
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v733+l2)))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v733+v700)))
	v752 = (v737-v739)*v743 + ((v746-v748)*v726 + v724)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v698+v735)))
	v755 = v754 * v743
	v757 = v725 - v732
	v759 = v727 + v732
	if v759 != l1&int32(-2) {
		v724 = v752
		v725 = v757
		v726 = v755
		v727 = v759
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v765 = v752
	v766 = v757
	v767 = v755
	goto L188
L194:
	;
	goto L193
L195:
	;
	v776 = v766 << (uint(int32(2)) % 32)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l2+v776)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v776+v700)))
	v788 = (v778-v780)*v767 + v765
	goto L186
L196:
	;
	v804 = v796
	goto L198
L197:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v804 = (v797<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L198
L198:
	;
	v805 = v804 + v196
	v807 = F_array_seek(m, v805, int32(0), v692, v788, l6, l8)
	mBase = m.M
	v808 = v807 - v805
	if v692 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v932 = v808
	v934 = v788
	v935 = v928
	v936 = v689 - (v808 + v928)
	goto L178
L200:
	;
	v810 = base.I32_div_s(v788, int32(8))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v810))))
	if int32(base.Ui32(v812)>>(uint(v788&int32(7))%32))&int32(1) == int32(0) {
		v928 = v10
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if int32(0) < l6 {
		v911 = l6
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L202
L204:
	;
	switch l8 - int32(99) {
	case 0:
		v928 = v911
		goto L199
	case 1:
		goto L240
	default:
		goto L239
	case 6:
		goto L241
	}
L205:
	;
	if l6 == int32(-1) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v824 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L208
L208:
	;
	if v807&int32(3) == int32(0) {
		v873 = v807
		goto L224
	} else {
		goto L225
	}
L209:
	;
	v827 = int32(6)
	v829 = int32(18)
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+1)))
	if v831 == v829 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	v843 = int32(1)
	if v824&v843 != 0 {
		v911 = int32(base.Ui32(v824) >> (uint(v843) % 32))
		goto L204
	} else {
		goto L221
	}
L212:
	;
	v834 = v829
	goto L214
L213:
	;
	v834 = int32(2)
	goto L214
L214:
	;
	if v831&int32(254) == int32(2) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v839 = v827
	goto L217
L216:
	;
	v839 = v834
	goto L217
L217:
	;
	if v831 == int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v842 = v827
	goto L220
L219:
	;
	v842 = v839
	goto L220
L220:
	;
	v911 = v842
	goto L204
L221:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	v911 = int32(base.Ui32(v847) >> (uint(int32(2)) % 32))
	goto L204
L222:
	;
	v911 = v906 + int32(1)
	goto L204
L223:
	;
	v906 = v898 - v807
	goto L222
L224:
	;
	v877 = v873
	goto L233
L225:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v857 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v906 = int32(0)
	goto L222
L227:
	;
	goto L228
L228:
	;
	v862 = v807
	goto L229
L229:
	;
	v866 = v862 + int32(1)
	if v866&int32(3) == int32(0) {
		v873 = v866
		goto L224
	} else {
		goto L231
	}
L230:
	;
	v898 = v866
	goto L223
L231:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	if v871 != 0 {
		v862 = v866
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	v886 = int32(-2139062144)
	if (int32(16843008)-v883|v883)&v886 == v886 {
		v877 = v877 + int32(4)
		goto L233
	} else {
		goto L235
	}
L234:
	;
	v892 = v877
	goto L236
L235:
	;
	goto L234
L236:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	if v896 != 0 {
		v892 = v892 + int32(1)
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v898 = v892
	goto L223
L238:
	;
	goto L237
L239:
	;
	v928 = (v911 + int32(1)) & int32(-2)
	goto L199
L240:
	;
	v928 = (v911 + int32(7)) & int32(-8)
	goto L199
L241:
	;
	v928 = (v911 + int32(3)) & int32(-4)
	goto L199
L242:
	;
	v1045 = v932 + v674 + v936 + v1042
	v1046 = F_palloc0(m, v1045)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L24
	} else {
		goto L277
	}
L243:
	;
	if int32(0) < l6 {
		v1026 = l6
		goto L244
	} else {
		goto L245
	}
L244:
	;
	switch l8 - int32(99) {
	case 0:
		v1042 = v1026
		goto L242
	case 1:
		goto L275
	default:
		goto L274
	case 6:
		goto L276
	}
L245:
	;
	if l6 == int32(-1) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v941 == int32(1) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	if v61&int32(3) == int32(0) {
		v989 = v61
		goto L259
	} else {
		goto L260
	}
L249:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if base.Ui32((v945-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v1026 = int32(6)
		goto L244
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v959 = int32(1)
	if v941&v959 != 0 {
		v1026 = int32(base.Ui32(v941) >> (uint(v959) % 32))
		goto L244
	} else {
		goto L256
	}
L252:
	;
	v952 = int32(18)
	if v945&int32(255) == v952 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v958 = v952
	goto L255
L254:
	;
	v958 = int32(2)
	goto L255
L255:
	;
	v1026 = v958
	goto L244
L256:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v1026 = int32(base.Ui32(v963) >> (uint(int32(2)) % 32))
	goto L244
L257:
	;
	v1026 = v1022 + int32(1)
	goto L244
L258:
	;
	v1022 = v1014 - v61
	goto L257
L259:
	;
	v993 = v989
	goto L268
L260:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v973 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1022 = int32(0)
	goto L257
L262:
	;
	goto L263
L263:
	;
	v978 = v61
	goto L264
L264:
	;
	v982 = v978 + int32(1)
	if v982&int32(3) == int32(0) {
		v989 = v982
		goto L259
	} else {
		goto L266
	}
L265:
	;
	v1014 = v982
	goto L258
L266:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	if v987 != 0 {
		v978 = v982
		goto L264
	} else {
		goto L267
	}
L267:
	;
	goto L265
L268:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	v1002 = int32(-2139062144)
	if (int32(16843008)-v999|v999)&v1002 == v1002 {
		v993 = v993 + int32(4)
		goto L268
	} else {
		goto L270
	}
L269:
	;
	v1008 = v993
	goto L271
L270:
	;
	goto L269
L271:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	if v1012 != 0 {
		v1008 = v1008 + int32(1)
		goto L271
	} else {
		goto L273
	}
L272:
	;
	v1014 = v1008
	goto L258
L273:
	;
	goto L272
L274:
	;
	v1042 = (v1026 + int32(1)) & int32(-2)
	goto L242
L275:
	;
	v1042 = (v1026 + int32(7)) & int32(-8)
	goto L242
L276:
	;
	v1042 = (v1026 + int32(3)) & int32(-4)
	goto L242
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1046)+8)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v1046)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1046))) = v1045 << (uint(int32(2)) % 32)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1046)+12)) = v1053
	v1056 = v1046 + int32(16)
	if v359 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	if v359 != 0 {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	v1059 = F__emscripten_memcpy_bulkmem(m, v1056, v28+int32(96), v359)
	mBase = m.M
	v1060 = v1059
	goto L281
L280:
	;
	v1060 = v1056
	goto L281
L281:
	;
	goto L278
L282:
	;
	v1066 = v1046 + v674
	v1067 = v196 + v688
	if v932 != 0 {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	v1064 = F__emscripten_memcpy_bulkmem(m, v1060+v359, v28-int32(-64), v359)
	mBase = m.M
	goto L285
L284:
	;
	goto L285
L285:
	;
	goto L282
L286:
	;
	if v5 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	v1068 = F__emscripten_memcpy_bulkmem(m, v1066, v1067, v932)
	mBase = m.M
	v1069 = v1068
	goto L289
L288:
	;
	v1069 = v1066
	goto L289
L289:
	;
	goto L286
L290:
	;
	v1073 = F_ArrayCastAndSet(m, v61, l6, l7, l8, v1069+v932)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L24
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	if v936 != 0 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	goto L292
L294:
	;
	if v640&int32(1) == int32(0) {
		v2565 = v1046
		goto L3
	} else {
		goto L298
	}
L295:
	;
	v1079 = F__emscripten_memcpy_bulkmem(m, v1069+v932+v1042, v1067+v932+v935, v936)
	mBase = m.M
	goto L297
L296:
	;
	goto L297
L297:
	;
	goto L294
L298:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+4))
	v1088 = v1060 + v1085<<(uint(int32(3))%32)
	if v621 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1091 = v934
	goto L301
L300:
	;
	v1091 = v648 - int32(1)
	goto L301
L301:
	;
	v1093 = base.I32_div_s(v1091, int32(8))
	v1094 = v1088 + v1093
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	v1098 = v1091 & int32(7)
	if v5 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1104 = v1095 & base.I32_rotl(int32(-2), v1098)
	goto L304
L303:
	;
	v1104 = v1095 | int32(1)<<(uint(v1098)%32)
	goto L304
L304:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1094))) = uint8(v1104)
	if v639 != 0 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1567))) = uint8(v1564)
	v2565 = v1046
	goto L3
L306:
	;
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
	v1494 = int32(1)
	v1495 = v1115
	v1497 = v1492
	v1498 = v1114
	v1499 = v1111
	v1505 = v675
	v1509 = v692
	goto L391
L307:
	;
	if v675 <= int32(0) {
		v2565 = v1046
		goto L3
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	if v934 <= int32(0) {
		goto L319
	} else {
		goto L320
	}
L310:
	;
	v1111 = int32(1) << (uint(v639&int32(7)) % 32)
	v1113 = base.I32_div_s(v639, int32(8))
	v1114 = v1088 + v1113
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114))))
	if v692 != 0 {
		goto L306
	} else {
		goto L311
	}
L311:
	;
	v1117 = v1115
	v1120 = v1114
	v1121 = v1111
	v1127 = v675
	goto L312
L312:
	;
	v1141 = v1117 | v1121
	v1142 = int32(1)
	v1143 = v1127 - v1142
	v1145 = v1121 << (uint(v1142) % 32)
	if v1145 == int32(256) {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1564 = v1141
	v1567 = v1120
	goto L305
L314:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1120))) = uint8(v1141)
	if v1143 == int32(0) {
		v2565 = v1046
		goto L3
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1127) {
		v1117 = v1141
		v1121 = v1145
		v1127 = v1143
		goto L312
	} else {
		goto L318
	}
L317:
	;
	v1151 = int32(1)
	v1153 = v1120 + v1151
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153))))
	v1117 = v1154
	v1120 = v1153
	v1121 = v1151
	v1127 = v1143
	goto L312
L318:
	;
	goto L313
L319:
	;
	if v621 == int32(0) {
		v2565 = v1046
		goto L3
	} else {
		goto L359
	}
L320:
	;
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1088))))
	if v692 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1316))) = uint8(v1309)
	goto L319
L322:
	;
	v1239 = v1159
	v1243 = v934
	v1246 = v1088
	goto L349
L323:
	;
	if v934 != int32(1) {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
	v1167 = int32(1)
	v1170 = v1159
	v1171 = v1167
	v1174 = v1167
	v1175 = v934
	v1176 = v692
	v1177 = v1088
	v1181 = v1166
	goto L327
L326:
	;
	v1309 = v1159 | int32(1)
	v1316 = v1088
	goto L321
L327:
	;
	if v1171&v1181 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1214 != int32(1) {
		v1309 = v1213
		v1316 = v1215
		goto L321
	} else {
		goto L342
	}
L329:
	;
	v1199 = v1170 | v1174
	goto L331
L330:
	;
	v1199 = v1170 & (v1174 ^ int32(-1))
	goto L331
L331:
	;
	v1200 = int32(1)
	v1201 = v1175 - v1200
	v1203 = v1174 << (uint(v1200) % 32)
	if v1203 == int32(256) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1177))) = uint8(v1199)
	if v1201 == int32(0) {
		goto L319
	} else {
		goto L335
	}
L333:
	;
	v1213 = v1199
	v1214 = v1203
	v1215 = v1177
	goto L334
L334:
	;
	v1217 = v1171 << (uint(int32(1)) % 32)
	if v1217 == int32(256) {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	v1209 = int32(1)
	v1211 = v1177 + v1209
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	v1213 = v1212
	v1214 = v1209
	v1215 = v1211
	goto L334
L336:
	;
	goto L328
L337:
	;
	if v1201 == int32(0) {
		goto L336
	} else {
		goto L340
	}
L338:
	;
	v1226 = v1217
	v1227 = v1176
	v1228 = v1181
	goto L339
L339:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1175) {
		v1170 = v1213
		v1171 = v1226
		v1174 = v1214
		v1175 = v1201
		v1176 = v1227
		v1177 = v1215
		v1181 = v1228
		goto L327
	} else {
		goto L341
	}
L340:
	;
	v1222 = int32(1)
	v1223 = v1176 + v1222
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223))))
	v1226 = v1222
	v1227 = v1223
	v1228 = v1224
	goto L339
L341:
	;
	goto L336
L342:
	;
	goto L319
L343:
	;
	v1309 = v1239 | int32(3)
	v1316 = v1246
	goto L321
L344:
	;
	v1309 = v1239 | int32(7)
	v1316 = v1246
	goto L321
L345:
	;
	v1309 = v1239 | int32(15)
	v1316 = v1246
	goto L321
L346:
	;
	v1309 = v1239 | int32(31)
	v1316 = v1246
	goto L321
L347:
	;
	v1309 = v1239 | int32(63)
	v1316 = v1246
	goto L321
L348:
	;
	v1309 = v1239 | int32(127)
	v1316 = v1246
	goto L321
L349:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1243-int32(3)) {
		goto L343
	} else {
		goto L351
	}
L350:
	;
	v1309 = v1291 | int32(1)
	v1316 = v1290
	goto L321
L351:
	;
	v1268 = v1243 & int32(-2)
	if v1268 == int32(2) {
		goto L344
	} else {
		goto L352
	}
L352:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1243-int32(5)) {
		goto L345
	} else {
		goto L353
	}
L353:
	;
	if v1268 == int32(4) {
		goto L346
	} else {
		goto L354
	}
L354:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1243-int32(7)) {
		goto L347
	} else {
		goto L355
	}
L355:
	;
	if v1268 == int32(6) {
		goto L348
	} else {
		goto L356
	}
L356:
	;
	v1283 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1246))) = uint8(v1283)
	v1286 = v1243 - int32(8)
	if v1286 == int32(0) {
		goto L319
	} else {
		goto L357
	}
L357:
	;
	v1289 = int32(1)
	v1290 = v1246 + v1289
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290))))
	if v1286 != v1289 {
		v1239 = v1291
		v1243 = v1286
		v1246 = v1290
		goto L349
	} else {
		goto L358
	}
L358:
	;
	goto L350
L359:
	;
	v1362 = v934 + int32(1)
	v1365 = v675 + (v934 ^ int32(-1))
	if v1365 <= int32(0) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v2565 = v1046
	goto L3
L361:
	;
	goto L360
L362:
	;
	v1375 = int32(1) << (uint(v1362&int32(7)) % 32)
	v1377 = base.I32_div_s(v1362, int32(8))
	v1378 = v1088 + v1377
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	if v692 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1479))) = uint8(v1474)
	goto L361
L364:
	;
	v1383 = v1379
	v1386 = v1365
	v1387 = v1375
	v1388 = v1378
	goto L367
L365:
	;
	goto L366
L366:
	;
	v1417 = base.I32_div_s(v1362, int32(8))
	v1418 = v692 + v1417
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418))))
	v1420 = int32(1) << (uint(v1362&int32(7)) % 32)
	v1421 = v1379
	v1424 = v1365
	v1425 = v1375
	v1426 = v1378
	v1427 = v1418
	v1428 = v1419
	goto L375
L367:
	;
	v1391 = v1383 | v1387
	v1392 = int32(1)
	v1393 = v1386 - v1392
	v1395 = v1387 << (uint(v1392) % 32)
	if v1395 == int32(256) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	if v1406 != int32(1) {
		v1474 = v1405
		v1479 = v1407
		goto L363
	} else {
		goto L374
	}
L369:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1388))) = uint8(v1391)
	if v1393 == int32(0) {
		goto L361
	} else {
		goto L372
	}
L370:
	;
	v1405 = v1391
	v1406 = v1395
	v1407 = v1388
	goto L371
L371:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1386) {
		v1383 = v1405
		v1386 = v1393
		v1387 = v1406
		v1388 = v1407
		goto L367
	} else {
		goto L373
	}
L372:
	;
	v1401 = int32(1)
	v1403 = v1388 + v1401
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403))))
	v1405 = v1404
	v1406 = v1401
	v1407 = v1403
	goto L371
L373:
	;
	goto L368
L374:
	;
	goto L361
L375:
	;
	if v1420&v1428 != 0 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	if v1449 == int32(1) {
		goto L361
	} else {
		goto L390
	}
L377:
	;
	v1434 = v1421 | v1425
	goto L379
L378:
	;
	v1434 = v1421 & (v1425 ^ int32(-1))
	goto L379
L379:
	;
	v1435 = int32(1)
	v1436 = v1424 - v1435
	v1438 = v1425 << (uint(v1435) % 32)
	if v1438 == int32(256) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1426))) = uint8(v1434)
	if v1436 == int32(0) {
		goto L361
	} else {
		goto L383
	}
L381:
	;
	v1448 = v1434
	v1449 = v1438
	v1450 = v1426
	goto L382
L382:
	;
	v1452 = v1420 << (uint(int32(1)) % 32)
	if v1452 == int32(256) {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	v1444 = int32(1)
	v1446 = v1426 + v1444
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446))))
	v1448 = v1447
	v1449 = v1444
	v1450 = v1446
	goto L382
L384:
	;
	goto L376
L385:
	;
	if v1436 == int32(0) {
		goto L384
	} else {
		goto L388
	}
L386:
	;
	v1461 = v1452
	v1462 = v1427
	v1463 = v1428
	goto L387
L387:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1424) {
		v1420 = v1461
		v1421 = v1448
		v1424 = v1436
		v1425 = v1449
		v1426 = v1450
		v1427 = v1462
		v1428 = v1463
		goto L375
	} else {
		goto L389
	}
L388:
	;
	v1457 = int32(1)
	v1458 = v1427 + v1457
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458))))
	v1461 = v1457
	v1462 = v1458
	v1463 = v1459
	goto L387
L389:
	;
	goto L384
L390:
	;
	v1474 = v1448
	v1479 = v1450
	goto L363
L391:
	;
	if v1494&v1497 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	if v1540 == int32(1) {
		v2565 = v1046
		goto L3
	} else {
		goto L406
	}
L393:
	;
	v1524 = v1495 | v1499
	goto L395
L394:
	;
	v1524 = v1495 & (v1499 ^ int32(-1))
	goto L395
L395:
	;
	v1525 = int32(1)
	v1526 = v1505 - v1525
	v1528 = v1499 << (uint(v1525) % 32)
	if v1528 == int32(256) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1498))) = uint8(v1524)
	if v1526 == int32(0) {
		v2565 = v1046
		goto L3
	} else {
		goto L399
	}
L397:
	;
	v1538 = v1524
	v1539 = v1498
	v1540 = v1528
	goto L398
L398:
	;
	v1542 = v1494 << (uint(int32(1)) % 32)
	if v1542 == int32(256) {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	v1534 = int32(1)
	v1536 = v1498 + v1534
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536))))
	v1538 = v1537
	v1539 = v1536
	v1540 = v1534
	goto L398
L400:
	;
	goto L392
L401:
	;
	if v1526 == int32(0) {
		goto L400
	} else {
		goto L404
	}
L402:
	;
	v1551 = v1542
	v1552 = v1497
	v1553 = v1509
	goto L403
L403:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1505) {
		v1494 = v1551
		v1495 = v1538
		v1497 = v1552
		v1498 = v1539
		v1499 = v1540
		v1505 = v1526
		v1509 = v1553
		goto L391
	} else {
		goto L405
	}
L404:
	;
	v1547 = int32(1)
	v1548 = v1509 + v1547
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548))))
	v1551 = v1547
	v1552 = v1549
	v1553 = v1548
	goto L403
L405:
	;
	goto L400
L406:
	;
	v1564 = v1538
	v1567 = v1539
	goto L305
L407:
	;
	v1615 = v1610 + v1612
	if v1589 < v1615 {
		v1644 = v1611
		v1647 = v1613
		v1648 = v1614
		goto L6
	} else {
		goto L413
	}
L408:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v1610 = v1590
	v1611 = v126
	v1612 = v1592
	v1613 = v127
	v1614 = v132
	goto L407
L409:
	;
	goto L410
L410:
	;
	v1593 = v1590 - v1589
	if base.B2i32(v1593 < v1590)^base.B2i32(int32(0) < v1589) != 0 {
		goto L2
	} else {
		goto L411
	}
L411:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v1599 = v1598 + v1593
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1599
	if base.B2i32(v1593 < int32(0)) != base.B2i32(v1599 < v1598) {
		goto L2
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v1589
	v1606 = int32(1)
	v1610 = v1589
	v1611 = v1593
	v1612 = v1599
	v1613 = v1606
	v1614 = base.B2i32(v1606 < v1593) | v132
	goto L407
L413:
	;
	v1619 = v1589 - v1615
	if base.B2i32(int32(0) < v1615)^base.B2i32(v1619 < v1589) != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v1623 = v1619 + int32(1)
	if v1623 < v1619 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v1625 = v1623 + v1612
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1625
	if base.B2i32(v1623 < int32(0)) != base.B2i32(v1625 < v1612) {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1673 = v1611
	v1674 = v1623
	v1677 = base.B2i32(int32(1) < v1623) | v1614
	goto L5
L417:
	;
	v1673 = v1644
	v1674 = v1659
	v1677 = v1648
	goto L5
L418:
	;
	F_ArrayCheckBounds(m, l1, v28+int32(96), v28-int32(-64))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L24
	} else {
		goto L419
	}
L419:
	;
	v1706 = v1674
	v1709 = v1673
	v1713 = v1677
	v1724 = int32(1)
	goto L4
L420:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v72)+56))
	if v1825 < v1824 {
		goto L431
	} else {
		goto L432
	}
L421:
	;
	goto L420
L422:
	;
	v1741 = int32(1)
	if v1738 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	if l1&v1741 == int32(0) {
		v1816 = v1793
		goto L421
	} else {
		goto L430
	}
L424:
	;
	v1793 = v1729
	v1794 = v1738
	v1795 = v1741
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1752 = v1729
	v1753 = v1738
	v1754 = v1741
	v1755 = v1729
	goto L427
L427:
	;
	v1760 = int32(2)
	v1761 = v1753 << (uint(v1760) % 32)
	v1763 = v1761 - int32(4)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1763)))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1728+v1763)))
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1761+v1726)))
	v1771 = v1770 * v1754
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1761+l2)))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1761+v1728)))
	v1780 = (v1765-v1767)*v1771 + ((v1774-v1776)*v1754 + v1752)
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1726+v1763)))
	v1783 = v1782 * v1771
	v1785 = v1753 - v1760
	v1787 = v1755 + v1760
	if v1787 != l1&int32(-2) {
		v1752 = v1780
		v1753 = v1785
		v1754 = v1783
		v1755 = v1787
		goto L427
	} else {
		goto L429
	}
L428:
	;
	v1793 = v1780
	v1794 = v1785
	v1795 = v1783
	goto L423
L429:
	;
	goto L428
L430:
	;
	v1804 = v1794 << (uint(int32(2)) % 32)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1804)))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1804+v1728)))
	v1816 = (v1806-v1808)*v1795 + v1793
	goto L421
L431:
	;
	v1828 = base.I32_div_s(v1824, int32(8))
	v1829 = v1828 + v1824
	if v1824 < v1829 {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	v1845 = v133
	v1846 = v129
	v1847 = v1825
	goto L433
L433:
	;
	if base.B2i32(v1846 == int32(0))&v1713 != 0 {
		goto L443
	} else {
		goto L444
	}
L434:
	;
	v1831 = v1829
	goto L436
L435:
	;
	v1831 = v1824
	goto L436
L436:
	;
	v1834 = F_repalloc(m, v133, v1831<<(uint(int32(2))%32))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L24
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+48)) = v1834
	if v129 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+56)) = v1831
	v1845 = v1834
	v1846 = v1843
	v1847 = v1831
	goto L433
L439:
	;
	v1843 = int32(0)
	goto L438
L440:
	;
	goto L441
L441:
	;
	v1840 = F_repalloc(m, v129, v1831)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L24
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+52)) = v1840
	v1843 = v1840
	goto L438
L443:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v1852 = F_MemoryContextAllocZero(m, v1851, v1847)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L24
	} else {
		goto L446
	}
L444:
	;
	v1855 = v1846
	goto L445
L445:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v72)+64)) = int64(0)
	if v1724 != 0 {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+52)) = v1852
	v1855 = v1852
	goto L445
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+28)) = l1
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
	v1863 = l1 << (uint(int32(2)) % 32)
	if v1863 != 0 {
		goto L451
	} else {
		goto L452
	}
L448:
	;
	goto L449
L449:
	;
	if int32(0) < v1709 {
		goto L458
	} else {
		goto L459
	}
L450:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v72)+36))
	if v1863 != 0 {
		goto L455
	} else {
		goto L456
	}
L451:
	;
	v1864 = F__emscripten_memcpy_bulkmem(m, v1859, v28+int32(96), v1863)
	mBase = m.M
	goto L453
L452:
	;
	goto L453
L453:
	;
	goto L450
L454:
	;
	goto L449
L455:
	;
	v1869 = F__emscripten_memcpy_bulkmem(m, v1866, v28-int32(-64), v1863)
	mBase = m.M
	goto L457
L456:
	;
	goto L457
L457:
	;
	goto L454
L458:
	;
	v1874 = int32(2)
	v1875 = v1709 << (uint(v1874) % 32)
	v1876 = v1845 + v1875
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	v1879 = v1877 << (uint(v1874) % 32)
	if v1876 == v1845 {
		goto L462
	} else {
		goto L463
	}
L459:
	;
	goto L460
L460:
	;
	if int32(0) < v1706 {
		goto L558
	} else {
		goto L559
	}
L461:
	;
	v2026 = F__emscripten_memset_bulkmem(m, v1845, base.I32_extend8_s(int32(0)), v1875)
	mBase = m.M
	goto L507
L462:
	;
	goto L461
L463:
	;
	v1883 = v1876 + v1879
	if base.Ui32(v1845-v1883) <= base.Ui32(int32(0)-v1879<<(uint(int32(1))%32)) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1890 = F___memcpy(m, v1876, v1845, v1879)
	mBase = m.M
	goto L461
L465:
	;
	goto L466
L466:
	;
	v1893 = (v1876 ^ v1845) & int32(3)
	if base.Ui32(v1876) < base.Ui32(v1845) {
		goto L469
	} else {
		goto L470
	}
L467:
	;
	if v1995 == int32(0) {
		goto L462
	} else {
		goto L503
	}
L468:
	;
	if base.Ui32(v1973) <= base.Ui32(int32(3)) {
		v1994 = v1972
		v1995 = v1973
		v1996 = v1974
		goto L467
	} else {
		goto L499
	}
L469:
	;
	if v1893 != 0 {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	goto L471
L471:
	;
	if v1893 != 0 {
		v1955 = v1879
		goto L482
	} else {
		goto L483
	}
L472:
	;
	v1994 = v1845
	v1995 = v1879
	v1996 = v1876
	goto L467
L473:
	;
	goto L474
L474:
	;
	if v1876&int32(3) == int32(0) {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1972 = v1845
	v1973 = v1879
	v1974 = v1876
	goto L468
L476:
	;
	goto L477
L477:
	;
	v1900 = v1845
	v1901 = v1879
	v1902 = v1876
	goto L478
L478:
	;
	if v1901 == int32(0) {
		goto L462
	} else {
		goto L480
	}
L479:
	;
	v1972 = v1909
	v1973 = v1911
	v1974 = v1913
	goto L468
L480:
	;
	v1906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1900))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1902))) = uint8(v1906)
	v1908 = int32(1)
	v1909 = v1900 + v1908
	v1911 = v1901 - v1908
	v1913 = v1902 + v1908
	if v1913&int32(3) != 0 {
		v1900 = v1909
		v1901 = v1911
		v1902 = v1913
		goto L478
	} else {
		goto L481
	}
L481:
	;
	goto L479
L482:
	;
	if v1955 == int32(0) {
		goto L462
	} else {
		goto L495
	}
L483:
	;
	if v1883&int32(3) != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1920 = v1879
	goto L487
L485:
	;
	v1935 = v1879
	goto L486
L486:
	;
	if base.Ui32(v1935) <= base.Ui32(int32(3)) {
		v1955 = v1935
		goto L482
	} else {
		goto L491
	}
L487:
	;
	if v1920 == int32(0) {
		goto L462
	} else {
		goto L489
	}
L488:
	;
	v1935 = v1926
	goto L486
L489:
	;
	v1926 = v1920 - int32(1)
	v1927 = v1876 + v1926
	v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1845+v1926))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1927))) = uint8(v1929)
	if v1927&int32(3) != 0 {
		v1920 = v1926
		goto L487
	} else {
		goto L490
	}
L490:
	;
	goto L488
L491:
	;
	v1942 = v1935
	goto L492
L492:
	;
	v1946 = v1942 - int32(4)
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1845+v1946)))
	*(*int32)(unsafe.Add(mBase, uint32(v1876+v1946))) = v1949
	if base.Ui32(int32(3)) < base.Ui32(v1946) {
		v1942 = v1946
		goto L492
	} else {
		goto L494
	}
L493:
	;
	v1955 = v1946
	goto L482
L494:
	;
	goto L493
L495:
	;
	v1962 = v1955
	goto L496
L496:
	;
	v1966 = v1962 - int32(1)
	v1969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1845+v1966))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1876+v1966))) = uint8(v1969)
	if v1966 != 0 {
		v1962 = v1966
		goto L496
	} else {
		goto L498
	}
L497:
	;
	goto L462
L498:
	;
	goto L497
L499:
	;
	v1979 = v1972
	v1980 = v1973
	v1981 = v1974
	goto L500
L500:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1979)))
	*(*int32)(unsafe.Add(mBase, uint32(v1981))) = v1983
	v1985 = int32(4)
	v1986 = v1979 + v1985
	v1988 = v1981 + v1985
	v1990 = v1980 - v1985
	if base.Ui32(int32(3)) < base.Ui32(v1990) {
		v1979 = v1986
		v1980 = v1990
		v1981 = v1988
		goto L500
	} else {
		goto L502
	}
L501:
	;
	v1994 = v1986
	v1995 = v1990
	v1996 = v1988
	goto L467
L502:
	;
	goto L501
L503:
	;
	v2001 = v1994
	v2002 = v1995
	v2003 = v1996
	goto L504
L504:
	;
	v2005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2001))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2003))) = uint8(v2005)
	v2007 = int32(1)
	v2012 = v2002 - v2007
	if v2012 != 0 {
		v2001 = v2001 + v2007
		v2002 = v2012
		v2003 = v2003 + v2007
		goto L504
	} else {
		goto L506
	}
L505:
	;
	goto L462
L506:
	;
	goto L505
L507:
	;
	if v1855 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2027 = v1855 + v1709
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	if v2027 == v1855 {
		goto L512
	} else {
		goto L513
	}
L509:
	;
	goto L510
L510:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+60)) = v2176 + v1709
	goto L460
L511:
	;
	v2175 = F__emscripten_memset_bulkmem(m, v1855, base.I32_extend8_s(int32(1)), v1709)
	mBase = m.M
	goto L557
L512:
	;
	goto L511
L513:
	;
	v2032 = v2027 + v2028
	if base.Ui32(v1855-v2032) <= base.Ui32(int32(0)-v2028<<(uint(int32(1))%32)) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2039 = F___memcpy(m, v2027, v1855, v2028)
	mBase = m.M
	goto L511
L515:
	;
	goto L516
L516:
	;
	v2042 = (v2027 ^ v1855) & int32(3)
	if base.Ui32(v2027) < base.Ui32(v1855) {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	if v2144 == int32(0) {
		goto L512
	} else {
		goto L553
	}
L518:
	;
	if base.Ui32(v2122) <= base.Ui32(int32(3)) {
		v2143 = v2121
		v2144 = v2122
		v2145 = v2123
		goto L517
	} else {
		goto L549
	}
L519:
	;
	if v2042 != 0 {
		goto L522
	} else {
		goto L523
	}
L520:
	;
	goto L521
L521:
	;
	if v2042 != 0 {
		v2104 = v2028
		goto L532
	} else {
		goto L533
	}
L522:
	;
	v2143 = v1855
	v2144 = v2028
	v2145 = v2027
	goto L517
L523:
	;
	goto L524
L524:
	;
	if v2027&int32(3) == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2121 = v1855
	v2122 = v2028
	v2123 = v2027
	goto L518
L526:
	;
	goto L527
L527:
	;
	v2049 = v1855
	v2050 = v2028
	v2051 = v2027
	goto L528
L528:
	;
	if v2050 == int32(0) {
		goto L512
	} else {
		goto L530
	}
L529:
	;
	v2121 = v2058
	v2122 = v2060
	v2123 = v2062
	goto L518
L530:
	;
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2049))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2051))) = uint8(v2055)
	v2057 = int32(1)
	v2058 = v2049 + v2057
	v2060 = v2050 - v2057
	v2062 = v2051 + v2057
	if v2062&int32(3) != 0 {
		v2049 = v2058
		v2050 = v2060
		v2051 = v2062
		goto L528
	} else {
		goto L531
	}
L531:
	;
	goto L529
L532:
	;
	if v2104 == int32(0) {
		goto L512
	} else {
		goto L545
	}
L533:
	;
	if v2032&int32(3) != 0 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v2069 = v2028
	goto L537
L535:
	;
	v2084 = v2028
	goto L536
L536:
	;
	if base.Ui32(v2084) <= base.Ui32(int32(3)) {
		v2104 = v2084
		goto L532
	} else {
		goto L541
	}
L537:
	;
	if v2069 == int32(0) {
		goto L512
	} else {
		goto L539
	}
L538:
	;
	v2084 = v2075
	goto L536
L539:
	;
	v2075 = v2069 - int32(1)
	v2076 = v2027 + v2075
	v2078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2075))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2076))) = uint8(v2078)
	if v2076&int32(3) != 0 {
		v2069 = v2075
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	v2091 = v2084
	goto L542
L542:
	;
	v2095 = v2091 - int32(4)
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1855+v2095)))
	*(*int32)(unsafe.Add(mBase, uint32(v2027+v2095))) = v2098
	if base.Ui32(int32(3)) < base.Ui32(v2095) {
		v2091 = v2095
		goto L542
	} else {
		goto L544
	}
L543:
	;
	v2104 = v2095
	goto L532
L544:
	;
	goto L543
L545:
	;
	v2111 = v2104
	goto L546
L546:
	;
	v2115 = v2111 - int32(1)
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2115))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2027+v2115))) = uint8(v2118)
	if v2115 != 0 {
		v2111 = v2115
		goto L546
	} else {
		goto L548
	}
L547:
	;
	goto L512
L548:
	;
	goto L547
L549:
	;
	v2128 = v2121
	v2129 = v2122
	v2130 = v2123
	goto L550
L550:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	*(*int32)(unsafe.Add(mBase, uint32(v2130))) = v2132
	v2134 = int32(4)
	v2135 = v2128 + v2134
	v2137 = v2130 + v2134
	v2139 = v2129 - v2134
	if base.Ui32(int32(3)) < base.Ui32(v2139) {
		v2128 = v2135
		v2129 = v2139
		v2130 = v2137
		goto L550
	} else {
		goto L552
	}
L551:
	;
	v2143 = v2135
	v2144 = v2139
	v2145 = v2137
	goto L517
L552:
	;
	goto L551
L553:
	;
	v2150 = v2143
	v2151 = v2144
	v2152 = v2145
	goto L554
L554:
	;
	v2154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2150))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2152))) = uint8(v2154)
	v2156 = int32(1)
	v2161 = v2151 - v2156
	if v2161 != 0 {
		v2150 = v2150 + v2156
		v2151 = v2161
		v2152 = v2152 + v2156
		goto L554
	} else {
		goto L556
	}
L555:
	;
	goto L512
L556:
	;
	goto L555
L557:
	;
	goto L510
L558:
	;
	v2183 = v1706 & int32(3)
	v2184 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1706) {
		goto L561
	} else {
		goto L562
	}
L559:
	;
	goto L560
L560:
	;
	v2516 = int32(0)
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+46)))
	if v2517 == v2516 {
		goto L587
	} else {
		goto L588
	}
L561:
	;
	v2192 = v2184
	v2202 = int32(0)
	goto L564
L562:
	;
	v2252 = v2184
	goto L563
L563:
	;
	if v2183 != 0 {
		goto L567
	} else {
		goto L568
	}
L564:
	;
	v2216 = int32(2)
	v2217 = v2192 << (uint(v2216) % 32)
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	v2223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2217+(v1845+v2218<<(uint(v2216)%32))))) = v2223
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v2225<<(uint(v2216)%32)+v2217)+4)) = v2223
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v2232<<(uint(v2216)%32)+v2217)+8)) = v2223
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v2239<<(uint(v2216)%32)+v2217)+12)) = v2223
	v2246 = int32(4)
	v2247 = v2192 + v2246
	v2249 = v2202 + v2246
	if v2249 != v1706&int32(2147483644) {
		v2192 = v2247
		v2202 = v2249
		goto L564
	} else {
		goto L566
	}
L565:
	;
	v2252 = v2247
	goto L563
L566:
	;
	goto L565
L567:
	;
	v2277 = v2252
	v2286 = v2184
	goto L570
L568:
	;
	goto L569
L569:
	;
	if v1855 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L570:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	v2302 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v2301<<(uint(v2302)%32)+v2277<<(uint(v2302)%32)))) = int32(0)
	v2310 = int32(1)
	v2313 = v2286 + v2310
	if v2313 != v2183 {
		v2277 = v2277 + v2310
		v2286 = v2313
		goto L570
	} else {
		goto L572
	}
L571:
	;
	goto L569
L572:
	;
	goto L571
L573:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+60)) = v2488 + v1706
	goto L560
L574:
	;
	v2343 = v1706 & int32(3)
	v2344 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1706) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v2352 = v2344
	v2361 = int32(0)
	goto L578
L576:
	;
	v2402 = v2344
	goto L577
L577:
	;
	if v2343 == int32(0) {
		goto L573
	} else {
		goto L581
	}
L578:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	v2379 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2376+v2352))) = uint8(v2379)
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2381+v2352)+1)) = uint8(v2379)
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2386+v2352)+2)) = uint8(v2379)
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2391+v2352)+3)) = uint8(v2379)
	v2396 = int32(4)
	v2397 = v2352 + v2396
	v2399 = v2361 + v2396
	if v2399 != v1706&int32(2147483644) {
		v2352 = v2397
		v2361 = v2399
		goto L578
	} else {
		goto L580
	}
L579:
	;
	v2402 = v2397
	goto L577
L580:
	;
	goto L579
L581:
	;
	v2429 = v2402
	v2430 = v2344
	goto L582
L582:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v72)+60))
	v2456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v2453+v2429))) = uint8(v2456)
	v2461 = v2430 + v2456
	if v2461 != v2343 {
		v2429 = v2429 + v2456
		v2430 = v2461
		goto L582
	} else {
		goto L584
	}
L583:
	;
	goto L573
L584:
	;
	goto L583
L585:
	;
	if v2544 == int32(0) {
		goto L594
	} else {
		goto L595
	}
L586:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1855+v1816))) = uint8(v5)
	v2544 = v2541
	goto L585
L587:
	;
	if v1855 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L588:
	;
	v2534 = v2516
	goto L589
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v1816<<(uint(int32(2))%32)))) = v124
	if v1855 == int32(0) {
		v2544 = v2534
		goto L585
	} else {
		goto L593
	}
L590:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v1845+v1816<<(uint(int32(2))%32))))
	v2534 = v2533
	goto L589
L591:
	;
	v2523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1855+v1816))))
	if v2523 != int32(1) {
		goto L590
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1845+v1816<<(uint(int32(2))%32)))) = v124
	v2541 = v2516
	goto L586
L593:
	;
	v2541 = v2534
	goto L586
L594:
	;
	v2565 = v72 + int32(12)
	goto L3
L595:
	;
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v72)+72))
	if base.Ui32(v2547) <= base.Ui32(v2544) {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v72)+76))
	if base.Ui32(v2544) < base.Ui32(v2549) {
		goto L594
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	F_pfree(m, v2544)
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L24
	} else {
		goto L600
	}
L599:
	;
	goto L598
L600:
	;
	goto L594
L601:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L24
	} else {
		goto L602
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v28)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L24
	} else {
		goto L603
	}
L603:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2619), int32(_a_F_array_set_element_3))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L24
	} else {
		goto L604
	}
L604:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L605:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L24
	} else {
		goto L606
	}
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v28+int32(16))
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L24
	} else {
		goto L607
	}
L607:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2635), int32(_a_F_array_set_element_3))
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L24
	} else {
		goto L608
	}
L608:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_subscript_check_subscripts(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v143
L2:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v136)
	v143 = int32(0)
	goto L1
L3:
	;
	v19 = int32(0)
	v20 = v11
	goto L6
L4:
	;
	goto L5
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v74 <= int32(0) {
		v143 = v8
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v19))))
	if v26 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v19))))
	if v31 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v63 = v20
	goto L10
L10:
	;
	v65 = v19 + int32(1)
	if v65 < v63 {
		v19 = v65
		v20 = v63
		goto L6
	} else {
		goto L20
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v34 != int32(1) {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v56 = v19 << (uint(int32(2)) % 32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)+v56))) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v63 = v62
	goto L10
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_array_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(199), int32(_a_F_array_subscript_check_subscripts_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	goto L7
L21:
	;
	v82 = int32(0)
	v83 = v74
	goto L22
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v82))))
	if v89 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v143 = v8
	goto L1
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v82))))
	if v94 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v124 = v83
	goto L26
L26:
	;
	v126 = v82 + int32(1)
	if v126 < v124 {
		v82 = v126
		v83 = v124
		goto L22
	} else {
		goto L35
	}
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v97 != int32(1) {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v117 = v82 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(36)+v117))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v124 = v123
	goto L26
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_array_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(218), int32(_a_F_array_subscript_check_subscripts_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	goto L23
}
func F_array_to_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_detoast_datum_packed(m, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v16 == int32(1) {
					v19 = int32(4)
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
					if v21&int32(254) == int32(2) {
						v30 = v19
					} else {
						v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
					}
					if v21 == int32(1) {
						v33 = v19
					} else {
						v33 = v30
					}
					v46 = v33
				} else {
					v34 = int32(1)
					if v16&v34 != 0 {
						v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v49 = F_palloc(m, v46+int32(1))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = int32(1)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
					if v53&v51 != 0 {
						v56 = v51
					} else {
						v56 = int32(4)
					}
					if v46 != 0 {
						v58 = F__emscripten_memcpy_bulkmem(m, v49, v14+v56, v46)
						mBase = m.M
						v59 = v58
					} else {
						v59 = v49
					}
					v61 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v46+v59))) = uint8(v61)
					if v14 != v12 {
						F_pfree(m, v14)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v67 = F_array_to_text_internal(m, l0, v7, v59, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								return v67
							}
						}
					} else {
						v67 = F_array_to_text_internal(m, l0, v7, v59, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							return v67
						}
					}
				}
			}
		}
	}
}
func F_array_typanalyze(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_std_typanalyze(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v66 = int32(0)
			m.G0 = v9 + int32(16)
			return v66
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = F_get_base_element_type(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v76
						F_errmsg_internal(m, int32(_a_F_array_typanalyze_0), v9)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_typanalyze_1), int32(118), int32(_a_F_array_typanalyze_2))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = int32(1)
					v25 = F_lookup_type_cache(m, v19, int32(193))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
						if v27 == int32(0) {
							v66 = v23
							m.G0 = v9 + int32(16)
							return v66
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+108))
							if v30 == int32(0) {
								v66 = v23
								m.G0 = v9 + int32(16)
								return v66
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+136))
								if v33 == int32(0) {
									v66 = v23
									m.G0 = v9 + int32(16)
									return v66
								} else {
									v37 = F_palloc(m, int32(36))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
										*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v41
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v43
										v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+10)))
										*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v45)
										v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v37)+14)) = uint16(v47)
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+11)))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v25 + int32(132)
										*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v25 + int32(104)
										*(*uint8)(unsafe.Add(mBase, uint32(v37)+16)) = uint8(v49)
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v57
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v59
										*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v37
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1259)
										v66 = v23
										m.G0 = v9 + int32(16)
										return v66
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
func F_array_unnest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(_a_F_array_unnest_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_DatumGetAnyArrayP(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = F_palloc(m, int32(32))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v27 == int32(-1) {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
						if v30 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v30
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
							v33 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
							v89 = v33
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
							if v40 == int32(0) {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v50 = (v43<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v50 = v40
							}
							*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v39 + v50
							v53 = int32(0)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
							if v55 == v53 {
								v89 = v53
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
								v89 = v54 + v58<<(uint(int32(3))%32) + int32(16)
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
						if v66 != 0 {
							v74 = v66
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v74 + v22
						v77 = int32(0)
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
						if v78 == v77 {
							v89 = v77
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							v89 = v22 + v81<<(uint(int32(3))%32) + int32(16)
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = int64(1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v89
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v95 == int32(-1) {
						v98 = int32(28)
					} else {
						v98 = int32(4)
					}
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v22+v98)))
					if v95 == int32(-1) {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
						v106 = v103
					} else {
						v106 = v22 + int32(16)
					}
					v107 = F_ArrayGetNItems(m, v100, v106)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v107
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						if v110 == int32(-1) {
							v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
							*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)) = uint16(v113)
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+30)) = uint8(v115)
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+31)) = uint8(v117)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v25
							*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v17
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
							if v140 < v141 {
								*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = v140 + int32(1)
								v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+28)))
								v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+30)))
								v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+31)))
								v151 = F_array_iter_next(m, v139, l0+int32(16), v140, v148, v149, v150)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									v153 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
									*(*int64)(unsafe.Add(mBase, uint32(v138))) = v153 + int64(1)
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v157)+20)) = int32(1)
									return v151
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
									v166 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v166)
									return int32(0)
								}
							}
						} else {
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							F_get_typlenbyvalalign(m, v119, v25+int32(28), v25+int32(30), v25+int32(31))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v25
								*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v17
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
								if v140 < v141 {
									*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = v140 + int32(1)
									v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+28)))
									v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+30)))
									v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+31)))
									v151 = F_array_iter_next(m, v139, l0+int32(16), v140, v148, v149, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v153 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v153 + int64(1)
										v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v157)+20)) = int32(1)
										return v151
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return int32(0)
									} else {
										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
										v166 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v166)
										return int32(0)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
		v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
		v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
		if v140 < v141 {
			*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = v140 + int32(1)
			v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+28)))
			v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+30)))
			v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+31)))
			v151 = F_array_iter_next(m, v139, l0+int32(16), v140, v148, v149, v150)
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return int32(0)
			} else {
				v153 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
				*(*int64)(unsafe.Add(mBase, uint32(v138))) = v153 + int64(1)
				v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v157)+20)) = int32(1)
				return v151
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return int32(0)
			} else {
				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
				v166 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v166)
				return int32(0)
			}
		}
	}
}
func F_array_unnest_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(460) {
		v26 = v2
		return v26
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		if v8 == int32(0) {
			v26 = v2
			return v26
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v11 != int32(15) {
				v26 = v2
				return v26
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = F_estimate_expression_value(m, v14, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					v23 = F_estimate_array_length(m, v22, v18)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v4)+16)) = v23
						v26 = v4
						return v26
					}
				}
			}
		}
	}
}
func F_construct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v13
	v21 = F_construct_md_array(m, l0, int32(0), v13, v10+int32(12), v10+int32(8), l2, l3, l4, l5)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(16)
		return v21
	}
}
func F_fetch_array_arg_replace_nulls(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		v16 = F_MemoryContextAlloc(m, v14, int32(48))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v16
			v24 = v16
			v26 = v8 + int32(12)
			v27 = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v28 == v27 {
				v45 = int32(0)
				if v26 == v45 {
					v53 = v45
				} else {
					v48 = v45
					v49 = v27
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
				}
				v56 = v53
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				switch v31 - int32(429) {
				case 0:
					if v26 == int32(0) {
						v56 = int32(1)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
						v48 = v38
						v49 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
						v56 = v53
					}
				case 1:
					if v26 == int32(0) {
						v56 = int32(2)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+368))
						v48 = v43
						v49 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
						v56 = v53
					}
				default:
					v45 = int32(0)
					if v26 == v45 {
						v53 = v45
					} else {
						v48 = v45
						v49 = v27
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
					}
					v56 = v53
				}
			}
			if v56 == int32(0) {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v60
			} else {
			}
			v64 = l0 + l1<<(uint(int32(3))%32)
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
			if v65 == int32(0) {
				v68 = int32(_a_F_fetch_array_arg_replace_nulls_0)
				v69 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v71
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
				if v74 != int32(1) {
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v93 = F_expand_array(m, v73, v92, v24)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
						v97 = v95
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				} else {
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
					if v77 != int32(3) {
						v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
						v93 = F_expand_array(m, v73, v92, v24)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
							v97 = v95
							*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
							v123 = v97
							m.G0 = v8 + int32(16)
							return v123
						}
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2))
						if v24 == int32(0) {
							v97 = v80
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v24))) = v83
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)))
							*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v85)
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v87)
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+47)))
							*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v89)
							v97 = v80
						}
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				}
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v101 = F_get_fn_expr_argtype(m, v100, l1)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					if v101 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(118), int32(_a_F_fetch_array_arg_replace_nulls_3))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
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
						v105 = F_get_element_type(m, v101)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							if v105 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(123), int32(_a_F_fetch_array_arg_replace_nulls_3))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
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
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v111 = F_palloc0(m, int32(16))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v111))) = int64(64)
									v118 = F_expand_array(m, v111, v109, v24)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+2))
											v123 = v122
											m.G0 = v8 + int32(16)
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
	} else {
		v24 = v11
		v26 = v8 + int32(12)
		v27 = int32(0)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v28 == v27 {
			v45 = int32(0)
			if v26 == v45 {
				v53 = v45
			} else {
				v48 = v45
				v49 = v27
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
				v53 = v49
			}
			v56 = v53
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			switch v31 - int32(429) {
			case 0:
				if v26 == int32(0) {
					v56 = int32(1)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
					v48 = v38
					v49 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
					v56 = v53
				}
			case 1:
				if v26 == int32(0) {
					v56 = int32(2)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+368))
					v48 = v43
					v49 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
					v56 = v53
				}
			default:
				v45 = int32(0)
				if v26 == v45 {
					v53 = v45
				} else {
					v48 = v45
					v49 = v27
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
				}
				v56 = v53
			}
		}
		if v56 == int32(0) {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v60
		} else {
		}
		v64 = l0 + l1<<(uint(int32(3))%32)
		v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
		if v65 == int32(0) {
			v68 = int32(_a_F_fetch_array_arg_replace_nulls_0)
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
			if v74 != int32(1) {
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v93 = F_expand_array(m, v73, v92, v24)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
					v97 = v95
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
					v123 = v97
					m.G0 = v8 + int32(16)
					return v123
				}
			} else {
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
				if v77 != int32(3) {
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v93 = F_expand_array(m, v73, v92, v24)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
						v97 = v95
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2))
					if v24 == int32(0) {
						v97 = v80
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v24))) = v83
						v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)))
						*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v85)
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+46)))
						*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v87)
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+47)))
						*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v89)
						v97 = v80
					}
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
					v123 = v97
					m.G0 = v8 + int32(16)
					return v123
				}
			}
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v101 = F_get_fn_expr_argtype(m, v100, l1)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				if v101 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(118), int32(_a_F_fetch_array_arg_replace_nulls_3))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
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
					v105 = F_get_element_type(m, v101)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						if v105 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(123), int32(_a_F_fetch_array_arg_replace_nulls_3))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
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
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							v111 = F_palloc0(m, int32(16))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v105
								*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v111))) = int64(64)
								v118 = F_expand_array(m, v111, v109, v24)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v111)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+2))
										v123 = v122
										m.G0 = v8 + int32(16)
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
func F_initArrayResultWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	v3 = l2
	if v3 != 0 {
		v10 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_initArrayResultWithSize_0), int32(0), int32(_a_F_initArrayResultWithSize_1), int32(_a_F_initArrayResultWithSize_2))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = v10
			v16 = F_MemoryContextAlloc(m, v14, int32(32))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
				v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(2))%32))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v27 = F_MemoryContextAlloc(m, v14, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v27
						F_get_typlenbyvalalign(m, l0, v16+int32(24), v16+int32(26), v16+int32(27))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							return v16
						}
					}
				}
			}
		}
	} else {
		v14 = l1
		v16 = F_MemoryContextAlloc(m, v14, int32(32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
			v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(2))%32))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v27 = F_MemoryContextAlloc(m, v14, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v27
					F_get_typlenbyvalalign(m, l0, v16+int32(24), v16+int32(26), v16+int32(27))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			}
		}
	}
}
