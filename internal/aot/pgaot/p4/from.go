package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyFromBinaryStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v15-v16 <= int32(10) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L14
	} else {
		goto L117
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L14
	} else {
		goto L113
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L14
	} else {
		goto L109
	}
L5:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v13)+3))
	if v73^int64(-69144642808756400)|(v76^int64(2830138808553551)) != int64(0) {
		goto L4
	} else {
		goto L25
	}
L6:
	;
	v21 = v13
	v22 = v16
	v24 = int32(0)
	v25 = v15
	goto L9
L7:
	;
	goto L8
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v55 = v54 + v16
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+7)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v58
	v61 = v16 + int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v61
	v65 = v61
	v68 = v15
	v70 = v54
	goto L5
L9:
	;
	if v22 == v25 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v49 != int32(11) {
		goto L4
	} else {
		goto L24
	}
L11:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v36 = v22
	v37 = v25
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v40 = int32(11) - v24
	v41 = v37 - v36
	if v40 < v41 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	return
L15:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v33 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v36 = v35
	v37 = v34
	goto L13
L17:
	;
	v43 = v40
	goto L19
L18:
	;
	v43 = v41
	goto L19
L19:
	;
	if v43 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v21, v36+v38, v43)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v46 = v43 + v36
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v46
	v49 = v43 + v24
	if v49 < int32(11) {
		v21 = v43 + v21
		v22 = v46
		v24 = v49
		v25 = v37
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	v65 = v46
	v68 = v37
	v70 = v38
	goto L5
L25:
	;
	if v68-v65 <= int32(3) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v164 = int32(16711935)
	v170 = base.I32_rotr(v152, int32(24))&v164 | base.I32_rotr(v152&v164, int32(8))
	if v170&int32(_a_F_CopyFromBinaryStart_0) != 0 {
		goto L3
	} else {
		goto L50
	}
L27:
	;
	v89 = v13 + int32(12)
	v90 = v65
	v92 = int32(0)
	v93 = v68
	v95 = v70
	goto L31
L28:
	;
	goto L29
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v65+v70)))
	v150 = v65 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v150
	v152 = v148
	v154 = v150
	v157 = v68
	v159 = v70
	goto L26
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L14
	} else {
		goto L46
	}
L31:
	;
	if v90 == v93 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v118 != int32(4) {
		goto L30
	} else {
		goto L45
	}
L33:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L36
	}
L34:
	;
	v105 = v90
	v106 = v93
	v107 = v95
	goto L35
L35:
	;
	v109 = int32(4) - v92
	v110 = v106 - v105
	if v109 < v110 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v101 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v105 = v104
	v106 = v102
	v107 = v103
	goto L35
L38:
	;
	v112 = v109
	goto L40
L39:
	;
	v112 = v110
	goto L40
L40:
	;
	if v112 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	base.MemoryCopy(m, v89, v105+v107, v112)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v115 = v112 + v105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v115
	v118 = v112 + v92
	if v118 < int32(4) {
		v89 = v112 + v89
		v90 = v115
		v92 = v118
		v93 = v106
		v95 = v107
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v152 = v123
	v154 = v115
	v157 = v106
	v159 = v107
	goto L26
L46:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(205), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if base.Ui32(int32(_a_F_CopyFromBinaryStart_4)) <= base.Ui32(v170) {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v157-v154 <= int32(3) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v269 = v240
	v271 = v224
	v274 = v227
	v276 = v229
	goto L78
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L74
	}
L54:
	;
	v234 = int32(16711935)
	v240 = base.I32_rotr(v222, int32(24))&v234 | base.I32_rotr(v222&v234, int32(8))
	if v240 < int32(0) {
		goto L53
	} else {
		goto L73
	}
L55:
	;
	v182 = v13 + int32(12)
	v183 = v154
	v185 = int32(0)
	v186 = v157
	v188 = v159
	goto L58
L56:
	;
	goto L57
L57:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v154+v159)))
	v220 = v154 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v220
	v222 = v218
	v224 = v220
	v227 = v157
	v229 = v159
	goto L54
L58:
	;
	if v183 == v186 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v211 != int32(4) {
		goto L53
	} else {
		goto L72
	}
L60:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L63
	}
L61:
	;
	v198 = v183
	v199 = v186
	v200 = v188
	goto L62
L62:
	;
	v202 = int32(4) - v185
	v203 = v199 - v198
	if v202 < v203 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v194 != 0 {
		goto L53
	} else {
		goto L64
	}
L64:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v198 = v197
	v199 = v195
	v200 = v196
	goto L62
L65:
	;
	v205 = v202
	goto L67
L66:
	;
	v205 = v203
	goto L67
L67:
	;
	if v205 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v182, v198+v200, v205)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v208 = v205 + v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v208
	v211 = v205 + v185
	if v211 < int32(4) {
		v182 = v205 + v182
		v183 = v208
		v185 = v211
		v186 = v199
		v188 = v200
		goto L58
	} else {
		goto L71
	}
L71:
	;
	goto L59
L72:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v222 = v216
	v224 = v208
	v227 = v199
	v229 = v200
	goto L54
L73:
	;
	goto L52
L74:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_5), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(220), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	if int32(0) < v269 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	m.G0 = v13 + int32(16)
	goto L1
L80:
	;
	v281 = v269
	v283 = v271
	goto L83
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v292 = v281 - int32(1)
	if v274-v283 <= int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L82
L85:
	;
	v298 = v13
	v299 = v283
	v301 = int32(0)
	v302 = v274
	v304 = v276
	goto L89
L86:
	;
	goto L87
L87:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v276))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v356)
	v358 = int32(1)
	v359 = v283 + v358
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v359
	if v358 < v281 {
		v281 = v292
		v283 = v359
		goto L83
	} else {
		goto L108
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L14
	} else {
		goto L104
	}
L89:
	;
	if v299 == v302 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v327 == int32(1) {
		v269 = v292
		v271 = v324
		v274 = v315
		v276 = v316
		goto L78
	} else {
		goto L103
	}
L91:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L14
	} else {
		goto L94
	}
L92:
	;
	v314 = v299
	v315 = v302
	v316 = v304
	goto L93
L93:
	;
	v318 = int32(1) - v301
	v319 = v315 - v314
	if v318 < v319 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v310 != 0 {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v314 = v313
	v315 = v311
	v316 = v312
	goto L93
L96:
	;
	v321 = v318
	goto L98
L97:
	;
	v321 = v319
	goto L98
L98:
	;
	if v321 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	base.MemoryCopy(m, v298, v314+v316, v321)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v324 = v321 + v314
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v324
	v327 = v321 + v301
	if v327 <= int32(0) {
		v298 = v321 + v298
		v299 = v324
		v301 = v327
		v302 = v315
		v304 = v316
		goto L89
	} else {
		goto L102
	}
L102:
	;
	goto L90
L103:
	;
	goto L88
L104:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_6), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(227), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L14
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	goto L84
L109:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_7), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(200), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L14
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_8), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(209), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L14
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L14
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryStart_9), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L14
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(214), int32(_a_F_CopyFromBinaryStart_3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L14
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_from_char_seq_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
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
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v420
L2:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v408 + v316
	v420 = int32(1)
	goto L1
L3:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v340 = F_pstrdup(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L61
	} else {
		goto L95
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = (v317 - v311) >> (uint(int32(2)) % 32)
	if int32(0) < v316 {
		goto L2
	} else {
		goto L93
	}
L5:
	;
	v311 = l2
	v316 = v72 - v17
	v317 = v38
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	goto L3
L7:
	;
	if v18 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v18 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L10:
	;
	if base.Ui32((v18-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L12:
	;
	v31 = v18 | int32(32)
	goto L14
L13:
	;
	v31 = v18
	goto L14
L14:
	;
	goto L11
L15:
	;
	v38 = l2
	v42 = v32
	goto L16
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if base.Ui32((v47-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L6
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v115 != 0 {
		v38 = v38 + int32(4)
		v42 = v115
		goto L16
	} else {
		goto L37
	}
L19:
	;
	if v56 != v31 {
		goto L18
	} else {
		goto L23
	}
L20:
	;
	v56 = v47 | int32(32)
	goto L22
L21:
	;
	v56 = v47
	goto L22
L22:
	;
	goto L19
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v62 = v58
	v66 = v17
	goto L24
L24:
	;
	v72 = v66 + int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v73 == int32(0) {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	goto L18
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 == int32(0) {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v89 = v73 | int32(32)
	goto L31
L30:
	;
	v89 = v73
	goto L31
L31:
	;
	goto L28
L32:
	;
	if v89 == v99 {
		v62 = v62 + int32(1)
		v66 = v72
		goto L24
	} else {
		goto L36
	}
L33:
	;
	v99 = v90 | int32(32)
	goto L35
L34:
	;
	v99 = v90
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L25
L37:
	;
	goto L17
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v118 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v126 = v118
	v127 = l3
	goto L42
L40:
	;
	goto L41
L41:
	;
	v194 = F_strlen(m, v17)
	mBase = m.M
	v195 = F_str_toupper(m, v17, v194, l4)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L61
	} else {
		goto L62
	}
L42:
	;
	v131 = F_strlen(m, v126)
	mBase = m.M
	if v131 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	if v176 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v176 = int32(0)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v137 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v138 = v17
	v139 = v126
	v140 = v131
	v141 = v137
	goto L52
L49:
	;
	v164 = v126
	v168 = int32(0)
	goto L50
L50:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v176 = v168 - v169
	goto L44
L51:
	;
	v164 = v159
	v168 = v161
	goto L50
L52:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if base.B2i32(v141 != v143)|base.B2i32(v143 == int32(0)) != 0 {
		v159 = v139
		v161 = v141
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v159 = v153
	v161 = int32(0)
	goto L51
L54:
	;
	v149 = v140 - int32(1)
	if v149 == int32(0) {
		v159 = v139
		v161 = v141
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v152 = int32(1)
	v153 = v139 + v152
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v154 != 0 {
		v138 = v138 + v152
		v139 = v153
		v140 = v149
		v141 = v154
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v311 = l3
	v316 = v131
	v317 = v127
	goto L4
L58:
	;
	goto L59
L59:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v179 != 0 {
		v126 = v179
		v127 = v127 + int32(4)
		goto L42
	} else {
		goto L60
	}
L60:
	;
	goto L43
L61:
	;
	return int32(0)
L62:
	;
	v199 = F_strlen(m, v195)
	mBase = m.M
	v200 = F_str_tolower(m, v195, v199, l4)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v195)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v204 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v212 = v204
	v213 = l3
	goto L68
L66:
	;
	goto L67
L67:
	;
	F_pfree(m, v200)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L61
	} else {
		goto L92
	}
L68:
	;
	v217 = F_strlen(m, v212)
	mBase = m.M
	v218 = F_str_toupper(m, v212, v217, l4)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L61
	} else {
		goto L70
	}
L69:
	;
	goto L67
L70:
	;
	v220 = F_strlen(m, v218)
	mBase = m.M
	v221 = F_str_tolower(m, v218, v220, l4)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L61
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v218)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	v225 = F_strlen(m, v221)
	mBase = m.M
	if v225 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	F_pfree(m, v221)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L61
	} else {
		goto L86
	}
L74:
	;
	v270 = int32(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v231 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v232 = v200
	v233 = v221
	v234 = v225
	v235 = v231
	goto L81
L78:
	;
	v258 = v221
	v262 = int32(0)
	goto L79
L79:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v270 = v262 - v263
	goto L73
L80:
	;
	v258 = v253
	v262 = v255
	goto L79
L81:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if base.B2i32(v235 != v237)|base.B2i32(v237 == int32(0)) != 0 {
		v253 = v233
		v255 = v235
		goto L80
	} else {
		goto L83
	}
L82:
	;
	v253 = v247
	v255 = int32(0)
	goto L80
L83:
	;
	v243 = v234 - int32(1)
	if v243 == int32(0) {
		v253 = v233
		v255 = v235
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v246 = int32(1)
	v247 = v233 + v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	if v248 != 0 {
		v232 = v232 + v246
		v233 = v247
		v234 = v243
		v235 = v248
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	if v270 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v200)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L61
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v277 != 0 {
		v212 = v277
		v213 = v213 + int32(4)
		goto L68
	} else {
		goto L91
	}
L90:
	;
	v311 = l3
	v316 = v225
	v317 = v213
	goto L4
L91:
	;
	goto L69
L92:
	;
	goto L6
L93:
	;
	goto L3
L94:
	;
	v384 = int32(0)
	v385 = F_errsave_start(m, l6)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L61
	} else {
		goto L104
	}
L95:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v342 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v352 = v340
	v353 = v342
	goto L97
L97:
	;
	v357 = base.I32_extend8_s(v353)
	goto L99
L98:
	;
	goto L94
L99:
	;
	if base.B2i32(v357 == int32(32))|base.B2i32(base.Ui32((v357-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v367 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v352))) = uint8(v367)
	goto L94
L101:
	;
	goto L102
L102:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v369 != 0 {
		v352 = v352 + int32(1)
		v353 = v369
		goto L97
	} else {
		goto L103
	}
L103:
	;
	goto L98
L104:
	;
	if v385 == int32(0) {
		v420 = v384
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L61
	} else {
		goto L106
	}
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v340
	F_errmsg(m, int32(_a_F_from_char_seq_search_0), v15)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L61
	} else {
		goto L107
	}
L107:
	;
	F_errdetail(m, int32(_a_F_from_char_seq_search_1), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L61
	} else {
		goto L108
	}
L108:
	;
	F_errsave_finish(m, l6, int32(_a_F_from_char_seq_search_2), int32(2493), int32(_a_F_from_char_seq_search_3))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L61
	} else {
		goto L109
	}
L109:
	;
	v420 = v384
	goto L1
}
