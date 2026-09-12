package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddToDataDirLockFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16512)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v3
	v22 = F_open(m, int32(418275), int32(2), v14+int32(112))
	mBase = m.M
	if v22 < v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16512)
	return
L2:
	;
	v27 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v43 = int32(4074956)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(167772187)
	v50 = F_read(m, v22, v14+int32(8320), int32(8191))
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	if v50 < v53 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return
L6:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(418275)
	F_errmsg(m, int32(287912), v14)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(476454), int32(1587), int32(375835))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	v59 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v77 = v14 + int32(8320)
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v50))) = uint8(v79)
	v83 = int32(1)
	if l0 < int32(2) {
		v129 = v77
		v130 = v83
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v75 = F_close(m, v22)
	mBase = m.M
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(418275)
	F_errmsg(m, int32(287978), v14+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(476454), int32(1598), int32(375835))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = l1
	v242 = v14 + int32(8320)
	v247 = F_pg_snprintf(m, v231, v242-v231, int32(713373), v14+int32(96))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L67
	}
L22:
	;
	v138 = v14 + int32(8320)
	v141 = v129 - v138
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	v91 = v77
	v92 = v83
	goto L24
L24:
	;
	v97 = int32(10)
	v98 = F___strchrnul(m, v91, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == v97 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v115 = v14 + int32(8320)
	v118 = v108 - v115
	if v118 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v104 == int32(0) {
		v129 = v91
		v130 = v92
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v104 = v98
	goto L29
L28:
	;
	v104 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	v107 = int32(1)
	v108 = v104 + v107
	v110 = v92 + v107
	if v110 != l0 {
		v91 = v108
		v92 = v110
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	v231 = v14 + int32(128) + v118
	v234 = v108
	goto L21
L33:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v14+int32(128), v115, v118)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v146 = v14 + int32(128) + v141
	if l0 <= v130 {
		v231 = v146
		v234 = v129
		goto L21
	} else {
		goto L40
	}
L37:
	;
	v142 = F__emscripten_memcpy_bulkmem(m, v14+int32(128), v138, v141)
	mBase = m.M
	goto L39
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	v149 = v14 + int32(8320)
	v152 = (l0 - v130) & int32(3)
	if v152 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v130-l0) {
		v231 = v179
		v234 = v129
		goto L21
	} else {
		goto L51
	}
L42:
	;
	v179 = v146
	v184 = v130
	goto L41
L43:
	;
	goto L44
L44:
	;
	v157 = v146
	v162 = v130
	v165 = v3
	goto L45
L45:
	;
	if base.Ui32(v157) < base.Ui32(v149) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v179 = v171
	v184 = v173
	goto L41
L47:
	;
	v167 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v167)
	v171 = v157 + int32(1)
	goto L49
L48:
	;
	v171 = v157
	goto L49
L49:
	;
	v172 = int32(1)
	v173 = v162 + v172
	v175 = v165 + v172
	if v175 != v152 {
		v157 = v171
		v162 = v173
		v165 = v175
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v193 = v179
	v198 = v184
	goto L52
L52:
	;
	if base.Ui32(v193) < base.Ui32(v149) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v231 = v225
	v234 = v129
	goto L21
L54:
	;
	v203 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v203)
	v207 = v193 + int32(1)
	goto L56
L55:
	;
	v207 = v193
	goto L56
L56:
	;
	if base.Ui32(v207) < base.Ui32(v149) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v209 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v209)
	v213 = v207 + int32(1)
	goto L59
L58:
	;
	v213 = v207
	goto L59
L59:
	;
	if base.Ui32(v213) < base.Ui32(v149) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v215 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v215)
	v219 = v213 + int32(1)
	goto L62
L61:
	;
	v219 = v213
	goto L62
L62:
	;
	if base.Ui32(v219) < base.Ui32(v149) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v221 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v221)
	v225 = v219 + int32(1)
	goto L65
L64:
	;
	v225 = v219
	goto L65
L65:
	;
	v227 = v198 + int32(4)
	if v227 != l0 {
		v193 = v225
		v198 = v227
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v249 = int32(10)
	v250 = F___strchrnul(m, v234, v249)
	mBase = m.M
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v252 == v249 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v256 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v256 = v250
	goto L71
L70:
	;
	v256 = int32(0)
	goto L71
L71:
	;
	goto L68
L72:
	;
	if v231&int32(3) == int32(0) {
		v280 = v231
		goto L77
	} else {
		goto L78
	}
L73:
	;
	goto L74
L74:
	;
	v327 = v14 + int32(128)
	if v327&int32(3) == int32(0) {
		v351 = v327
		goto L95
	} else {
		goto L96
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v256 + int32(1)
	v317 = v231 + v313
	v322 = F_pg_snprintf(m, v317, v242-v317, int32(198531), v14+int32(80))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L92
	}
L76:
	;
	v313 = v305 - v231
	goto L75
L77:
	;
	v284 = v280
	goto L86
L78:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v264 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v313 = int32(0)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v269 = v231
	goto L82
L82:
	;
	v273 = v269 + int32(1)
	if v273&int32(3) == int32(0) {
		v280 = v273
		goto L77
	} else {
		goto L84
	}
L83:
	;
	v305 = v273
	goto L76
L84:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v278 != 0 {
		v269 = v273
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v293 = int32(-2139062144)
	if (int32(16843008)-v290|v290)&v293 == v293 {
		v284 = v284 + int32(4)
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v299 = v284
	goto L89
L88:
	;
	goto L87
L89:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v303 != 0 {
		v299 = v299 + int32(1)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v305 = v299
	goto L76
L91:
	;
	goto L90
L92:
	;
	goto L74
L93:
	;
	v386 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[166])) = v386
	v388 = int32(4074956)
	v389 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = int32(167772189)
	v395 = F_pwrite(m, v22, v14+int32(128), v384, int64(0))
	mBase = m.M
	v397 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = v386
	if v384 != v395 {
		goto L110
	} else {
		goto L111
	}
L94:
	;
	v384 = v376 - v327
	goto L93
L95:
	;
	v355 = v351
	goto L104
L96:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v335 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v384 = int32(0)
	goto L93
L98:
	;
	goto L99
L99:
	;
	v340 = v327
	goto L100
L100:
	;
	v344 = v340 + int32(1)
	if v344&int32(3) == int32(0) {
		v351 = v344
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v376 = v344
	goto L94
L102:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v349 != 0 {
		v340 = v344
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v364 = int32(-2139062144)
	if (int32(16843008)-v361|v361)&v364 == v364 {
		v355 = v355 + int32(4)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v370 = v355
	goto L107
L106:
	;
	goto L105
L107:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if v374 != 0 {
		v370 = v370 + int32(1)
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v376 = v370
	goto L94
L109:
	;
	goto L108
L110:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v402 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = int32(167772188)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
	if v433 != int32(1) {
		v447 = int32(0)
		goto L125
	} else {
		goto L126
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(51)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v410 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	if v410 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v426 = F_close(m, v22)
	mBase = m.M
	goto L1
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(418275)
	F_errmsg(m, int32(287595), v14-int32(-64))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(476454), int32(1662), int32(375835))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v472 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v472
	v474 = F_close(m, v22)
	mBase = m.M
	if v474 == v472 {
		goto L1
	} else {
		goto L137
	}
L124:
	;
	if v447 == int32(0) {
		goto L123
	} else {
		goto L131
	}
L125:
	;
	goto L124
L126:
	;
	goto L127
L127:
	;
	v438 = F_fsync(m, v22)
	mBase = m.M
	if v438 != int32(-1) {
		v447 = v438
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v447 = int32(-1)
	goto L125
L129:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v442 == int32(27) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v452 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	if v452 == int32(0) {
		goto L123
	} else {
		goto L133
	}
L133:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(418275)
	F_errmsg(m, int32(287595), v14+int32(48))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(476454), int32(1673), int32(375835))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	goto L123
L137:
	;
	v479 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	if v479 == int32(0) {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(418275)
	F_errmsg(m, int32(287595), v14+int32(32))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(476454), int32(1681), int32(375835))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	goto L1
}
func F_CopyToBinaryEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v9, v5+int32(14), int32(2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_CopySendEndOfRow(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_CopyToCSVOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L8
	} else {
		goto L62
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v105 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v106 <= v105 {
		goto L1
	} else {
		goto L29
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = F_OutputFunctionCall(m, v18+v22*int32(28), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v44&int32(3) == int32(0) {
		v68 = v44
		goto L13
	} else {
		goto L14
	}
L8:
	;
	return
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v22))))
	F_CopyAttributeOutCSV(m, l0, v36, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	F_appendBinaryStringInfo(m, v43, v44, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L28
	}
L12:
	;
	v101 = v93 - v44
	goto L11
L13:
	;
	v72 = v68
	goto L22
L14:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v52 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v101 = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v57 = v44
	goto L18
L18:
	;
	v61 = v57 + int32(1)
	if v61&int32(3) == int32(0) {
		v68 = v61
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v93 = v61
	goto L12
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v66 != 0 {
		v57 = v61
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 == v81 {
		v72 = v72 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v87 = v72
	goto L25
L24:
	;
	goto L23
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		v87 = v87 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v93 = v87
	goto L12
L27:
	;
	goto L26
L28:
	;
	goto L4
L29:
	;
	v111 = v105
	goto L30
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v121 = int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v111<<(uint(v121)%32))))
	v125 = int32(1)
	v126 = v124 - v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v127))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v121)%32))))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v141 <= v138+v125 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L1
L32:
	;
	if v129&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	F_appendStringInfoChar(m, v137, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*uint8)(unsafe.Add(mBase, uint32(v145+v138))) = uint8(v136)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v151 = v149 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153+v151))) = uint8(v155)
	goto L32
L36:
	;
	goto L32
L37:
	;
	v234 = v111 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v234 < v235 {
		v111 = v234
		goto L30
	} else {
		goto L61
	}
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v162&int32(3) == int32(0) {
		v186 = v162
		goto L43
	} else {
		goto L44
	}
L39:
	;
	goto L40
L40:
	;
	v225 = F_OutputFunctionCall(m, v18+v126*int32(28), v134)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L59
	}
L41:
	;
	F_appendBinaryStringInfo(m, v161, v162, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L58
	}
L42:
	;
	v219 = v211 - v162
	goto L41
L43:
	;
	v190 = v186
	goto L52
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v170 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v219 = int32(0)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v175 = v162
	goto L48
L48:
	;
	v179 = v175 + int32(1)
	if v179&int32(3) == int32(0) {
		v186 = v179
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v211 = v179
	goto L42
L50:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v184 != 0 {
		v175 = v179
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v199 = int32(-2139062144)
	if (int32(16843008)-v196|v196)&v199 == v199 {
		v190 = v190 + int32(4)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v205 = v190
	goto L55
L54:
	;
	goto L53
L55:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v209 != 0 {
		v205 = v205 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v211 = v205
	goto L42
L57:
	;
	goto L56
L58:
	;
	goto L37
L59:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+v126))))
	F_CopyAttributeOutCSV(m, l0, v225, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	goto L37
L61:
	;
	goto L31
L62:
	;
	return
}
func F_CopyToTextLikeOutFunc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_getTypeOutputInfo(m, l1, v6+int32(12), v6+int32(11))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		F_fmgr_info(m, v14, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_CopyToTextLikeStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = F_pg_server_to_any(m, v13, v14, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v16
	goto L3
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L33
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = l1 - int32(76)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v36 = v27 + v28<<(uint(int32(4))%32) + v33*int32(100)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v37&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v48 <= v47 {
		goto L9
	} else {
		goto L18
	}
L13:
	;
	F_CopyAttributeOutText(m, l0, v36)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_CopyAttributeOutCSV(m, l0, v36, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	goto L12
L18:
	;
	v54 = v47
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v54<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v71 <= v68+int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L9
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v95 = v27 + v89<<(uint(int32(4))%32) + v64*int32(100)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v96 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_appendStringInfoChar(m, v67, v66)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v68))) = uint8(v66)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = v79 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v81))) = uint8(v85)
	goto L21
L25:
	;
	goto L21
L26:
	;
	v105 = v54 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v105 < v106 {
		v54 = v105
		goto L19
	} else {
		goto L32
	}
L27:
	;
	F_CopyAttributeOutCSV(m, l0, v95, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_CopyAttributeOutText(m, l0, v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	goto L26
L32:
	;
	goto L20
L33:
	;
	goto L8
}
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_exprType(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(16) {
			v26 = l1
			v27 = F_expression_returns_set(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg(m, int32(101437), v9)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = F_exprLocation(m, v26)
								mBase = m.M
								F_parser_errposition(m, l0, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(482586), int32(1190), int32(273516))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
				} else {
					m.G0 = v9 + int32(32)
					return v26
				}
			}
		} else {
			v18 = int32(-1)
			v22 = F_coerce_to_target_type(m, l0, l1, v11, int32(16), v18, int32(1), int32(2), v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_format_type_be(m, v11)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(273762)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
								F_errmsg(m, int32(181246), v9+int32(16))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = F_exprLocation(m, l1)
									mBase = m.M
									F_parser_errposition(m, l0, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(482586), int32(1180), int32(273516))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
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
				} else {
					v26 = v22
					v27 = F_expression_returns_set(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
									F_errmsg(m, int32(101437), v9)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = F_exprLocation(m, v26)
										mBase = m.M
										F_parser_errposition(m, l0, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(482586), int32(1190), int32(273516))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
						} else {
							m.G0 = v9 + int32(32)
							return v26
						}
					}
				}
			}
		}
	}
}
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(160)
	switch l1 - int32(8) {
	case 0:
		v51 = v15
		v52 = int32(7939)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 1:
		v51 = v15
		v52 = int32(544265)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if base.Ui32(l1) <= base.Ui32(int32(41)) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[769])))
					v40 = v39
				} else {
					v40 = int32(722455)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v40
				F_errmsg(m, int32(428669), v12)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(480916), int32(78), int32(308686))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 8:
		v51 = v15
		v52 = int32(7842)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 21:
		v51 = int32(128)
		v52 = int32(696466)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	}
}
func F_to_oct64(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v14 = v9 - v8
	v15 = v14
	v16 = v12
	goto L1
L1:
	;
	v22 = v15 - int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)&int32(7))+uint32(_consts[865]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v28)
	if base.Ui64(v16) < base.Ui64(int64(8)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v14 - v22
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v22) {
		v15 = v22
		v16 = int64(base.Ui64(v16) >> (uint(int64(3)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v39
L9:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v39+int32(4), v22, v36)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_regclass(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[844]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[845]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1496), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regtype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[844]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[845]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1253), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_tsquery_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v17 = F_text_to_cstring(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v22 = int32(0)
			v24 = F_parse_tsquery(m, v17, int32(1173), v6+int32(8), v22, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func F_to_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1171), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
