package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimeCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v222 float64
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 float64
	_ = v443
	var v446 float64
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v465 int64
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	v5 = int32(0)
	v8 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(31744)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v5
	v25 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v28 == int32(68) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v493
L3:
	;
	v493 = int32(-2)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v32 = int32(-1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(58) {
		v493 = v32
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v45 = F_strtol(m, v33+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v48 == int32(68) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v493 = int32(-2)
	goto L2
L9:
	;
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	switch v53 - int32(46) {
	case 0:
		goto L14
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v493 = v32
		goto L2
	case 12:
		goto L13
	default:
		goto L15
	}
L11:
	;
	if base.Ui32(int32(59)) < base.Ui32(v472) {
		goto L134
	} else {
		goto L135
	}
L12:
	;
	v465 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if v465 < int64(0) {
		goto L131
	} else {
		goto L132
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v261 = F_strtol(m, v52+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L74
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
	v73 = v52 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v53 != 0 {
		v493 = v32
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(0)
	if l1 != int32(6144) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui64(v60-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v493 = int32(-2)
	goto L2
L19:
	;
	goto L20
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v45
	v69 = base.I32_wrap_i64(v60)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v69
	v472 = v69
	goto L11
L21:
	;
	v75 = int32(539317)
	v79 = m.G0
	v81 = v79 - int32(32)
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[946])))
	if v90 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v230 = v8
	goto L23
L23:
	;
	v233 = base.F64_nearest(base.F64_mul(v230, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v233), float64(2.147483648e+09)) != 0 {
		goto L68
	} else {
		goto L69
	}
L24:
	;
	if v73&int32(3) == int32(0) {
		v182 = v73
		goto L47
	} else {
		goto L48
	}
L25:
	;
	v158 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _consts[947])))
	if v94 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v73
	goto L31
L29:
	;
	goto L30
L30:
	;
	v108 = v75
	v109 = v90
	goto L34
L31:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v104 == v90 {
		v98 = v98 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v158 = v98 - v73
	goto L24
L33:
	;
	goto L32
L34:
	;
	v116 = v81 + int32(base.Ui32(v109)>>(uint(int32(3))%32))&int32(28)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v118 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 | v118<<(uint(v109)%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v122 != 0 {
		v108 = v108 + v118
		v109 = v122
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v125 == int32(0) {
		v150 = v73
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v158 = v150 - v73
	goto L24
L38:
	;
	v129 = v73
	v130 = v125
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(base.Ui32(v130)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v138)>>(uint(v130)%32))&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v150 = v146
	goto L37
L41:
	;
	v150 = v129
	goto L37
L42:
	;
	goto L43
L43:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v146 = v129 + int32(1)
	if v144 != 0 {
		v129 = v146
		v130 = v144
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	if v158 != v215 {
		v493 = v32
		goto L2
	} else {
		goto L62
	}
L46:
	;
	v215 = v207 - v73
	goto L45
L47:
	;
	v186 = v182
	goto L56
L48:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v166 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v215 = int32(0)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v171 = v73
	goto L52
L52:
	;
	v175 = v171 + int32(1)
	if v175&int32(3) == int32(0) {
		v182 = v175
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v207 = v175
	goto L46
L54:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v180 != 0 {
		v171 = v175
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v195 = int32(-2139062144)
	if (int32(16843008)-v192|v192)&v195 == v195 {
		v186 = v186 + int32(4)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v201 = v186
	goto L59
L58:
	;
	goto L57
L59:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v205 != 0 {
		v201 = v201 + int32(1)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v207 = v201
	goto L46
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v222 = F_strtod(m, v52, v12+int32(12))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return int32(0)
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v227 != 0 {
		v493 = v32
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v229 != 0 {
		v493 = v32
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v230 = v222
	goto L23
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui64(v241-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v237 = base.I32_trunc_f64_s(v233)
	v239 = v237
	goto L67
L69:
	;
	goto L70
L70:
	;
	v239 = int32(-2147483648)
	goto L67
L71:
	;
	v493 = int32(-2)
	goto L2
L72:
	;
	goto L73
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = int64(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v250 = base.I32_wrap_i64(v241)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v249
	v472 = v250
	goto L11
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v261
	v264 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v264 == int32(68) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v493 = int32(-2)
	goto L2
L76:
	;
	goto L77
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v269 == int32(0) {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if v269 != int32(46) {
		v493 = v32
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v276 = m.G0
	v278 = v276 - int32(16)
	m.G0 = v278
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v268
	v282 = v268 + int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v283 == int32(0) {
		v443 = v8
		goto L81
	} else {
		goto L82
	}
L80:
	;
	m.G0 = v278 + int32(16)
	if v456 != 0 {
		v493 = v32
		goto L2
	} else {
		goto L130
	}
L81:
	;
	v446 = base.F64_nearest(base.F64_mul(v443, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v446), float64(2.147483648e+09)) != 0 {
		goto L127
	} else {
		goto L128
	}
L82:
	;
	v286 = int32(539317)
	v290 = m.G0
	v292 = v290 - int32(32)
	v293 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v292)+24)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v292)+16)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v292)+8)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = v293
	v301 = int32(*(*uint8)(unsafe.Add(mBase, _consts[946])))
	if v301 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v456 = int32(-1)
	goto L80
L84:
	;
	if v282&int32(3) == int32(0) {
		v393 = v282
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v369 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, _consts[947])))
	if v305 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v309 = v282
	goto L91
L89:
	;
	goto L90
L90:
	;
	v319 = v286
	v320 = v301
	goto L94
L91:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v315 == v301 {
		v309 = v309 + int32(1)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v369 = v309 - v282
	goto L84
L93:
	;
	goto L92
L94:
	;
	v327 = v292 + int32(base.Ui32(v320)>>(uint(int32(3))%32))&int32(28)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v329 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v328 | v329<<(uint(v320)%32)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	if v333 != 0 {
		v319 = v319 + v329
		v320 = v333
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v336 == int32(0) {
		v361 = v282
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v369 = v361 - v282
	goto L84
L98:
	;
	v340 = v282
	v341 = v336
	goto L99
L99:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v292+int32(base.Ui32(v341)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v349)>>(uint(v341)%32))&int32(1) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v361 = v357
	goto L97
L101:
	;
	v361 = v340
	goto L97
L102:
	;
	goto L103
L103:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	v357 = v340 + int32(1)
	if v355 != 0 {
		v340 = v357
		v341 = v355
		goto L99
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	if v369 != v426 {
		goto L83
	} else {
		goto L122
	}
L106:
	;
	v426 = v418 - v282
	goto L105
L107:
	;
	v397 = v393
	goto L116
L108:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v377 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v426 = int32(0)
	goto L105
L110:
	;
	goto L111
L111:
	;
	v382 = v282
	goto L112
L112:
	;
	v386 = v382 + int32(1)
	if v386&int32(3) == int32(0) {
		v393 = v386
		goto L107
	} else {
		goto L114
	}
L113:
	;
	v418 = v386
	goto L106
L114:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v391 != 0 {
		v382 = v386
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v406 = int32(-2139062144)
	if (int32(16843008)-v403|v403)&v406 == v406 {
		v397 = v397 + int32(4)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v412 = v397
	goto L119
L118:
	;
	goto L117
L119:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	if v416 != 0 {
		v412 = v412 + int32(1)
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v418 = v412
	goto L106
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v433 = F_strtod(m, v268, v278+int32(12))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L63
	} else {
		goto L123
	}
L123:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v436 != 0 {
		goto L83
	} else {
		goto L124
	}
L124:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v438 == int32(0) {
		v443 = v433
		goto L81
	} else {
		goto L125
	}
L125:
	;
	goto L83
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)))) = v452
	v456 = int32(0)
	goto L80
L127:
	;
	v450 = base.I32_trunc_f64_s(v446)
	v452 = v450
	goto L126
L128:
	;
	goto L129
L129:
	;
	v452 = int32(-2147483648)
	goto L126
L130:
	;
	goto L12
L131:
	;
	v493 = int32(-2)
	goto L2
L132:
	;
	goto L133
L133:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v472 = v469
	goto L11
L134:
	;
	v493 = int32(-2)
	goto L2
L135:
	;
	goto L136
L136:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v480 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v493 = int32(-2)
	goto L2
L138:
	;
	goto L139
L139:
	;
	if base.Ui32(int32(60)) < base.Ui32(v480) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v493 = int32(-2)
	goto L2
L141:
	;
	goto L142
L142:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(int32(1000000)) < base.Ui32(v488) {
		v493 = int32(-2)
		goto L2
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v488
	v493 = int32(0)
	goto L2
}
func F_time_hash_extended(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hashint8extended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_time_mi_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v7 - v5
		return v9
	}
}
func F_time_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pq_getmsgint64(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if base.Ui64(v6) < base.Ui64(int64(86400000001)) {
			if base.Ui32(v4) <= base.Ui32(int32(6)) {
				v15 = v4 << (uint(int32(3)) % 32)
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[1248])))
				v19 = v18 + v6
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[1249])))
				v23 = base.I64_rem_s(v19, v22)
				v27 = v19 - v23
			} else {
				v27 = v6
			}
			v28 = F_Int64GetDatum(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(396582), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490962), int32(1601), int32(36204))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
