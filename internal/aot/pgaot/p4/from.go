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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v55 = v54 + v16
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+7)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v58
	v61 = v16 + int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v61
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v33 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v46
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v150
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
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v101 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v115
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(209), int32(_a_F_CopyFromBinaryStart_3))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v220
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
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v194 != 0 {
		goto L53
	} else {
		goto L64
	}
L64:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v208
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(224), int32(_a_F_CopyFromBinaryStart_3))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v359
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
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v310 != 0 {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v324
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(231), int32(_a_F_CopyFromBinaryStart_3))
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(204), int32(_a_F_CopyFromBinaryStart_3))
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(213), int32(_a_F_CopyFromBinaryStart_3))
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
	F_errfinish(m, int32(_a_F_CopyFromBinaryStart_2), int32(218), int32(_a_F_CopyFromBinaryStart_3))
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
func F_EndCopyFrom(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.T0[v9].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
		if v12 == int32(1) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v16 = F_ClosePipeStream(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				switch v16 + int32(1) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_EndCopyFrom_0), int32(0))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_EndCopyFrom_1), int32(1977), int32(_a_F_EndCopyFrom_2))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				case 1:
					v117 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[0]))
					if v117 == int32(0) {
					} else {
						v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EndCopyFrom[1])))
						if v121&int32(1) == int32(0) {
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+220))
							if v126 == int32(0) {
							} else {
								v129 = int32(_a_F_EndCopyFrom_3)
								v131 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
								v132 = int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v131 + v132
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135 + v132
								v139 = int32(0)
								v141 = int32(_a_F_EndCopyFrom_4)
								v142 = base.AtomicRmwOr32(m, v139, v141, v139)
								*(*int32)(unsafe.Add(mBase, uint32(v117)+220)) = v139
								*(*int32)(unsafe.Add(mBase, uint32(v117)+224)) = v139
								v150 = base.AtomicRmwOr32(m, v139, v141, v139)
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								*(*int32)(unsafe.Add(mBase, uint32(v117))) = v151 + v132
								v157 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v157 - v132
							}
						}
					}
					v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
					F_MemoryContextDelete(m, v161)
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return
						} else {
							m.G0 = v6 + int32(48)
							return
						}
					}
				default:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
					if v35 == int32(0) {
						v43 = v16 & int32(127)
						v49 = int32(255)
						if base.B2i32(int32(13) == v43)&base.B2i32(base.Ui32(v16&int32(_a_F_EndCopyFrom_5)-int32(1)) < base.Ui32(v49))|base.B2i32(v43 == int32(0))&base.B2i32(int32(141) == int32(base.Ui32(v16)>>(uint(int32(8))%32))&v49) != 0 {
							v117 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[0]))
							if v117 == int32(0) {
							} else {
								v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EndCopyFrom[1])))
								if v121&int32(1) == int32(0) {
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+220))
									if v126 == int32(0) {
									} else {
										v129 = int32(_a_F_EndCopyFrom_3)
										v131 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
										v132 = int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v131 + v132
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
										*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135 + v132
										v139 = int32(0)
										v141 = int32(_a_F_EndCopyFrom_4)
										v142 = base.AtomicRmwOr32(m, v139, v141, v139)
										*(*int32)(unsafe.Add(mBase, uint32(v117)+220)) = v139
										*(*int32)(unsafe.Add(mBase, uint32(v117)+224)) = v139
										v150 = base.AtomicRmwOr32(m, v139, v141, v139)
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
										*(*int32)(unsafe.Add(mBase, uint32(v117))) = v151 + v132
										v157 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v157 - v132
									}
								}
							}
							v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
							F_MemoryContextDelete(m, v161)
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return
								} else {
									m.G0 = v6 + int32(48)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errcode(m, int32(515))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v70
									F_errmsg(m, int32(_a_F_EndCopyFrom_6), v6+int32(16))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v77 = F_wait_result_to_str(m, v16)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v77
											F_errdetail_internal(m, int32(_a_F_EndCopyFrom_7), v6)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_EndCopyFrom_1), int32(1994), int32(_a_F_EndCopyFrom_2))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errcode(m, int32(515))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v70
								F_errmsg(m, int32(_a_F_EndCopyFrom_6), v6+int32(16))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									v77 = F_wait_result_to_str(m, v16)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v77
										F_errdetail_internal(m, int32(_a_F_EndCopyFrom_7), v6)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_EndCopyFrom_1), int32(1994), int32(_a_F_EndCopyFrom_2))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
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
				}
			}
		} else {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v88 == int32(0) {
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[0]))
				if v117 == int32(0) {
				} else {
					v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EndCopyFrom[1])))
					if v121&int32(1) == int32(0) {
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+220))
						if v126 == int32(0) {
						} else {
							v129 = int32(_a_F_EndCopyFrom_3)
							v131 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
							v132 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v131 + v132
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135 + v132
							v139 = int32(0)
							v141 = int32(_a_F_EndCopyFrom_4)
							v142 = base.AtomicRmwOr32(m, v139, v141, v139)
							*(*int32)(unsafe.Add(mBase, uint32(v117)+220)) = v139
							*(*int32)(unsafe.Add(mBase, uint32(v117)+224)) = v139
							v150 = base.AtomicRmwOr32(m, v139, v141, v139)
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							*(*int32)(unsafe.Add(mBase, uint32(v117))) = v151 + v132
							v157 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v157 - v132
						}
					}
				}
				v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
				F_MemoryContextDelete(m, v161)
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return
					} else {
						m.G0 = v6 + int32(48)
						return
					}
				}
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v92 = F_FreeFile(m, v91)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					if v92 == int32(0) {
						v117 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[0]))
						if v117 == int32(0) {
						} else {
							v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EndCopyFrom[1])))
							if v121&int32(1) == int32(0) {
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+220))
								if v126 == int32(0) {
								} else {
									v129 = int32(_a_F_EndCopyFrom_3)
									v131 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
									v132 = int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v131 + v132
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
									*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135 + v132
									v139 = int32(0)
									v141 = int32(_a_F_EndCopyFrom_4)
									v142 = base.AtomicRmwOr32(m, v139, v141, v139)
									*(*int32)(unsafe.Add(mBase, uint32(v117)+220)) = v139
									*(*int32)(unsafe.Add(mBase, uint32(v117)+224)) = v139
									v150 = base.AtomicRmwOr32(m, v139, v141, v139)
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
									*(*int32)(unsafe.Add(mBase, uint32(v117))) = v151 + v132
									v157 = *(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_EndCopyFrom[2])) = v157 - v132
								}
							}
						}
						v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
						F_MemoryContextDelete(m, v161)
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								m.G0 = v6 + int32(48)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v102
								F_errmsg(m, int32(_a_F_EndCopyFrom_8), v6+int32(32))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_EndCopyFrom_1), int32(1954), int32(_a_F_EndCopyFrom_9))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
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
		}
	}
}
func F_from_char_seq_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v520 int32
	_ = v520
	var v538 int32
	_ = v538
	var v564 int32
	_ = v564
	var v574 int32
	_ = v574
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v720 int32
	_ = v720
	v17 = m.G0
	v19 = v17 - int32(336)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v19 + int32(336)
	return v720
L2:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v709 + v700
	v720 = int32(1)
	goto L1
L3:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v617 = F_pstrdup(m, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L56
	} else {
		goto L140
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = (v588 - l3) >> (uint(int32(2)) % 32)
	if v587 != 0 {
		v700 = v587
		goto L2
	} else {
		goto L138
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = (v46 - l2) >> (uint(int32(2)) % 32)
	v700 = v85 - v21
	goto L2
L6:
	;
	F_pfree(m, v227)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L56
	} else {
		goto L137
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	goto L3
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v132 = F_strlen(m, v21)
	mBase = m.M
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v133 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32((v24-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = v24 | int32(32)
	goto L15
L14:
	;
	v38 = v24
	goto L15
L15:
	;
	v46 = l2
	v47 = v27
	goto L16
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if base.Ui32((v55-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L7
L18:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v131 != 0 {
		v46 = v46 + int32(4)
		v47 = v131
		goto L16
	} else {
		goto L34
	}
L19:
	;
	v64 = v55 | int32(32)
	goto L21
L20:
	;
	v64 = v55
	goto L21
L21:
	;
	if v64&int32(255) != v38 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v72 = v21
	v76 = v47
	goto L23
L23:
	;
	v85 = v72 + int32(1)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v86 == int32(0) {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L18
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v89 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32((v86-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v102 = v86 | int32(32)
	goto L29
L28:
	;
	v102 = v86
	goto L29
L29:
	;
	if base.Ui32((v89-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v111 = v89 | int32(32)
	goto L32
L31:
	;
	v111 = v89
	goto L32
L32:
	;
	if v102 == v111 {
		v72 = v85
		v76 = v76 + int32(1)
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L24
L34:
	;
	goto L17
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v136 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v139 = v136
	v145 = l3
	goto L39
L37:
	;
	goto L38
L38:
	;
	v220 = F_pg_newlocale_from_collation(m, l4)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v153 = F_strlen(m, v139)
	mBase = m.M
	if v153 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	if v198 == int32(0) {
		v587 = v153
		v588 = v145
		goto L4
	} else {
		goto L54
	}
L42:
	;
	v198 = int32(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v159 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v160 = v21
	v161 = v139
	v162 = v153
	v163 = v159
	goto L49
L46:
	;
	v186 = v139
	v190 = int32(0)
	goto L47
L47:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v198 = v190 - v191
	goto L41
L48:
	;
	v186 = v181
	v190 = v183
	goto L47
L49:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if base.B2i32(v163 != v165)|base.B2i32(v165 == int32(0)) != 0 {
		v181 = v161
		v183 = v163
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v181 = v175
	v183 = int32(0)
	goto L48
L51:
	;
	v171 = v162 - int32(1)
	if v171 == int32(0) {
		v181 = v161
		v183 = v163
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v174 = int32(1)
	v175 = v161 + v174
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v176 != 0 {
		v160 = v160 + v174
		v161 = v175
		v162 = v171
		v163 = v176
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v201 != 0 {
		v139 = v201
		v145 = v145 + int32(4)
		goto L39
	} else {
		goto L55
	}
L55:
	;
	goto L40
L56:
	;
	return int32(0)
L57:
	;
	v224 = F_str_toupper(m, v21, v132, l4)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v226 = F_strlen(m, v224)
	mBase = m.M
	v227 = F_str_tolower(m, v224, v226, l4)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v224)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v231 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v232 = v21 + v132
	v235 = l3
	v240 = v231
	goto L64
L62:
	;
	goto L63
L63:
	;
	F_pfree(m, v227)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L56
	} else {
		goto L136
	}
L64:
	;
	v250 = v19 + int32(96)
	v252 = F_strlen(m, v240)
	mBase = m.M
	v253 = F_pg_strupper(m, v250, int32(80), v240, v252, v220)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L56
	} else {
		goto L67
	}
L65:
	;
	goto L63
L66:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v520 != 0 {
		v235 = v235 + int32(4)
		v240 = v520
		goto L64
	} else {
		goto L135
	}
L67:
	;
	if base.Ui32(int32(79)) < base.Ui32(v253) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v258 = v19 + int32(16)
	v260 = F_pg_strlower(m, v258, int32(80), v250, v253, v220)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L56
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(int32(79)) < base.Ui32(v260) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	if v260 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v308 != 0 {
		goto L66
	} else {
		goto L84
	}
L72:
	;
	v308 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v269 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v270 = v227
	v271 = v258
	v272 = v260
	v273 = v269
	goto L79
L76:
	;
	v296 = v258
	v300 = int32(0)
	goto L77
L77:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v308 = v300 - v301
	goto L71
L78:
	;
	v296 = v291
	v300 = v293
	goto L77
L79:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if base.B2i32(v273 != v275)|base.B2i32(v275 == int32(0)) != 0 {
		v291 = v271
		v293 = v273
		goto L78
	} else {
		goto L81
	}
L80:
	;
	v291 = v285
	v293 = int32(0)
	goto L78
L81:
	;
	v281 = v272 - int32(1)
	if v281 == int32(0) {
		v291 = v271
		v293 = v273
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v284 = int32(1)
	v285 = v271 + v284
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	if v286 != 0 {
		v270 = v270 + v284
		v271 = v285
		v272 = v281
		v273 = v286
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v309 = F_strlen(m, v227)
	mBase = m.M
	if v309 == v260 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_pfree(m, v227)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L56
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v313 = int32(0)
	if v260 == v313 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v587 = v132
	v588 = v235
	goto L4
L89:
	;
	v389 = v19 + int32(256)
	v391 = F_pg_strupper(m, v389, int32(80), v21, v379, v220)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L56
	} else {
		goto L105
	}
L90:
	;
	v379 = int32(0)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v318 = v19 + int32(16)
	v319 = v318 + v260
	v322 = v313
	v327 = v318
	goto L93
L93:
	;
	v337 = v322 + int32(1)
	v338 = F_pg_mblen_range(m, v327, v319)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L56
	} else {
		goto L95
	}
L94:
	;
	v342 = int32(0)
	if v337 == v342 {
		v379 = v342
		goto L89
	} else {
		goto L97
	}
L95:
	;
	v340 = v338 + v327
	if base.Ui32(v340) < base.Ui32(v319) {
		v322 = v337
		v327 = v340
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v345 = int32(0)
	if v132 == v345 {
		v379 = v342
		goto L89
	} else {
		goto L98
	}
L98:
	;
	v352 = v345
	v355 = v342
	goto L99
L99:
	;
	v365 = F_pg_mblen_range(m, v355+v21, v232)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L56
	} else {
		goto L101
	}
L100:
	;
	v379 = v367
	goto L89
L101:
	;
	v367 = v365 + v355
	v369 = v352 + int32(1)
	if base.Ui32(v337) <= base.Ui32(v369) {
		v379 = v367
		goto L89
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32(v367) < base.Ui32(v132) {
		v352 = v369
		v355 = v367
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v433 = int32(0)
	if v132 == v433 {
		goto L66
	} else {
		goto L117
	}
L105:
	;
	if base.Ui32(int32(79)) < base.Ui32(v391) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v396 = v19 + int32(176)
	v398 = F_pg_strlower(m, v396, int32(80), v389, v391, v220)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L56
	} else {
		goto L107
	}
L107:
	;
	if base.Ui32(int32(79)) < base.Ui32(v398) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v403 = v19 + int32(16)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if base.B2i32(v406 == int32(0))|base.B2i32(v406 != v409) != 0 {
		v427 = v406
		v428 = v409
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v427-v428 == int32(0) {
		v564 = v379
		goto L6
	} else {
		goto L116
	}
L110:
	;
	goto L109
L111:
	;
	v412 = v396
	v413 = v403
	goto L112
L112:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+1)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+1)))
	if v417 == int32(0) {
		v427 = v417
		v428 = v416
		goto L110
	} else {
		goto L114
	}
L113:
	;
	v427 = v417
	v428 = v416
	goto L110
L114:
	;
	v420 = int32(1)
	if v417 == v416 {
		v412 = v412 + v420
		v413 = v413 + v420
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	goto L104
L117:
	;
	v443 = v433
	goto L118
L118:
	;
	v453 = v19 + int32(256)
	v456 = F_pg_mblen_range(m, v443+v21, v232)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L56
	} else {
		goto L121
	}
L119:
	;
	goto L66
L120:
	;
	if base.Ui32(v458) < base.Ui32(v132) {
		v443 = v458
		goto L118
	} else {
		goto L134
	}
L121:
	;
	v458 = v456 + v443
	v459 = F_pg_strupper(m, v453, int32(80), v21, v458, v220)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L56
	} else {
		goto L122
	}
L122:
	;
	if base.Ui32(int32(79)) < base.Ui32(v459) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	v464 = v19 + int32(176)
	v466 = F_pg_strlower(m, v464, int32(80), v453, v459, v220)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L56
	} else {
		goto L124
	}
L124:
	;
	if base.Ui32(int32(79)) < base.Ui32(v466) {
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v471 = v19 + int32(16)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if base.B2i32(v474 == int32(0))|base.B2i32(v474 != v477) != 0 {
		v495 = v474
		v496 = v477
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v495-v496 == int32(0) {
		v564 = v458
		goto L6
	} else {
		goto L133
	}
L127:
	;
	goto L126
L128:
	;
	v480 = v464
	v481 = v471
	goto L129
L129:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	if v485 == int32(0) {
		v495 = v485
		v496 = v484
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v495 = v485
	v496 = v484
	goto L127
L131:
	;
	v488 = int32(1)
	if v485 == v484 {
		v480 = v480 + v488
		v481 = v481 + v488
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	goto L120
L134:
	;
	goto L119
L135:
	;
	goto L65
L136:
	;
	goto L7
L137:
	;
	v587 = v564
	v588 = v235
	goto L4
L138:
	;
	goto L3
L139:
	;
	v669 = int32(0)
	v670 = F_errsave_start(m, l6)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L56
	} else {
		goto L149
	}
L140:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617))))
	if v619 == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v622 = v617
	v629 = v619
	goto L142
L142:
	;
	v638 = base.I32_extend8_s(v629)
	goto L144
L143:
	;
	goto L139
L144:
	;
	if base.B2i32(v638 == int32(32))|base.B2i32(base.Ui32((v638-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v648 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v622))) = uint8(v648)
	goto L139
L146:
	;
	goto L147
L147:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+1)))
	if v650 != 0 {
		v622 = v622 + int32(1)
		v629 = v650
		goto L142
	} else {
		goto L148
	}
L148:
	;
	goto L143
L149:
	;
	if v670 == int32(0) {
		v720 = v669
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L56
	} else {
		goto L151
	}
L151:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v617
	F_errmsg(m, int32(_a_F_from_char_seq_search_0), v19)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L56
	} else {
		goto L152
	}
L152:
	;
	v686 = F_errdetail(m, int32(_a_F_from_char_seq_search_1), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L56
	} else {
		goto L153
	}
L153:
	;
	F_errsave_finish(m, l6, int32(_a_F_from_char_seq_search_2), int32(2570), int32(_a_F_from_char_seq_search_3))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L56
	} else {
		goto L154
	}
L154:
	;
	v720 = v669
	goto L1
}
