package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddFileToBackupManifest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	v8 = m.G0
	v10 = v8 - int32(1184)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1176)) = l4
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L104
	}
L2:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(1184)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(466108)
	v24 = F_pg_snprintf(m, v10+int32(144), int32(1024), int32(166798), v10+int32(48))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v28 = l2
	goto L7
L7:
	;
	F_initStringInfo(m, v10+int32(128))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	v28 = v10 + int32(144)
	goto L7
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v33 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v28&int32(3) == int32(0) {
		v71 = v28
		goto L19
	} else {
		goto L20
	}
L12:
	;
	F_appendStringInfoChar(m, v10+int32(128), int32(10))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(713550))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v41)
	goto L11
L16:
	;
	goto L11
L17:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v107 != 0 {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v104 = v96 - v28
	goto L17
L19:
	;
	v75 = v71
	goto L28
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v55 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v60 = v28
	goto L24
L24:
	;
	v64 = v60 + int32(1)
	if v64&int32(3) == int32(0) {
		v71 = v64
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v96 = v64
	goto L18
L26:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v69 != 0 {
		v60 = v64
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v84 = int32(-2139062144)
	if (int32(16843008)-v81|v81)&v84 == v84 {
		v75 = v75 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v90 = v75
	goto L31
L30:
	;
	goto L29
L31:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 != 0 {
		v90 = v90 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v96 = v90
	goto L18
L33:
	;
	goto L32
L34:
	;
	F_appendStringInfoString(m, v10+int32(128), v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L56
	}
L35:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(688699))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L41
	}
L36:
	;
	v110 = F_pg_verify_mbstr(m, int32(6), v28, v104, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	if v110 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(703709))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_escape_json_with_len(m, v10+int32(128), v28, v104)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v229 = int32(703976)
	goto L34
L41:
	;
	F_enlargeStringInfo(m, v10+int32(128), v104<<(uint(int32(1))%32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v137 = v135 + v136
	v141 = v28 + v104
	if base.Ui32(v141) <= base.Ui32(v28) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v223 + base.I32_wrap_i64(base.I64_extend_i32_u(v104)<<(uint(int64(1))%64))
	v229 = int32(703975)
	goto L34
L44:
	;
	goto L43
L45:
	;
	v144 = v104 & int32(3)
	if v144 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v145 = v28
	v147 = v137
	v148 = int32(0)
	goto L49
L47:
	;
	v165 = v28
	v167 = v137
	goto L48
L48:
	;
	if base.Ui32(v104-int32(1)) < base.Ui32(int32(3)) {
		goto L44
	} else {
		goto L52
	}
L49:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v152 = int32(1)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151<<(uint(v152)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v147))) = uint16(v156)
	v159 = v147 + int32(2)
	v161 = v145 + v152
	v163 = v148 + v152
	if v163 != v144 {
		v145 = v161
		v147 = v159
		v148 = v163
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v165 = v161
	v167 = v159
	goto L48
L51:
	;
	goto L50
L52:
	;
	v175 = v165
	v177 = v167
	goto L53
L53:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v182 = int32(1)
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181<<(uint(v182)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v177))) = uint16(v186)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188<<(uint(v182)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+2)) = uint16(v193)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195<<(uint(v182)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+4)) = uint16(v200)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202<<(uint(v182)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+6)) = uint16(v207)
	v212 = v175 + int32(4)
	if v212 != v141 {
		v175 = v212
		v177 = v177 + int32(8)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L44
L55:
	;
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l3
	F_appendStringInfo(m, v10+int32(128), int32(703745), v10+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(688719))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v245 = int32(128)
	F_enlargeStringInfo(m, v10+v245, v245)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v257 = F_pg_gmtime(m, v10+int32(1176))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	v259 = F_pg_strftime(m, v250+v251, int32(128), int32(484424), v257)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v259 + v261
	F_appendStringInfoChar(m, v10+int32(128), int32(34))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v269 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v271 = v10 - int32(-64)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	switch v273 - int32(1) {
	case 0:
		goto L72
	case 1:
		goto L71
	case 2:
		goto L70
	case 3:
		goto L69
	case 4:
		goto L68
	default:
		v322 = int32(0)
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(6925))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L8
	} else {
		goto L101
	}
L66:
	;
	if v324 < int32(0) {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v324 = v322
	goto L66
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v315 = F_pg_cryptohash_final(m, v313, v271, int32(64))
	mBase = m.M
	if v315 < int32(0) {
		v322 = int32(-1)
		goto L67
	} else {
		goto L79
	}
L69:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v305 = F_pg_cryptohash_final(m, v303, v271, int32(48))
	mBase = m.M
	if v305 < int32(0) {
		v322 = int32(-1)
		goto L67
	} else {
		goto L77
	}
L70:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v295 = F_pg_cryptohash_final(m, v293, v271, int32(32))
	mBase = m.M
	if v295 < int32(0) {
		v322 = int32(-1)
		goto L67
	} else {
		goto L75
	}
L71:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v285 = F_pg_cryptohash_final(m, v283, v271, int32(28))
	mBase = m.M
	if v285 < int32(0) {
		v322 = int32(-1)
		goto L67
	} else {
		goto L73
	}
L72:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v278 = v276 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v278
	v324 = int32(4)
	goto L66
L73:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v324 = int32(28)
	goto L66
L75:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v324 = int32(32)
	goto L66
L77:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v324 = int32(48)
	goto L66
L79:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v322 = int32(64)
	goto L67
L81:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v327) <= base.Ui32(int32(5)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v327<<(uint(int32(2))%32))+uint32(_consts[400])))
	v336 = v334
	goto L84
L83:
	;
	v336 = int32(521570)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v336
	F_appendStringInfo(m, v10+int32(128), int32(688655), v10+int32(16))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	F_enlargeStringInfo(m, v10+int32(128), v324<<(uint(int32(1))%32))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v352 = v10 - int32(-64)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v355 = v353 + v354
	v359 = v352 + v324
	if base.Ui32(v359) <= base.Ui32(v352) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v441 + base.I32_wrap_i64(base.I64_extend_i32_u(v324)<<(uint(int64(1))%64))
	F_appendStringInfoChar(m, v10+int32(128), int32(34))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L8
	} else {
		goto L100
	}
L88:
	;
	goto L87
L89:
	;
	v362 = v324 & int32(3)
	if v362 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v363 = v352
	v365 = v355
	v366 = int32(0)
	goto L93
L91:
	;
	v383 = v352
	v385 = v355
	goto L92
L92:
	;
	if base.Ui32(v324-int32(1)) < base.Ui32(int32(3)) {
		goto L88
	} else {
		goto L96
	}
L93:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	v370 = int32(1)
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369<<(uint(v370)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v365))) = uint16(v374)
	v377 = v365 + int32(2)
	v379 = v363 + v370
	v381 = v366 + v370
	if v381 != v362 {
		v363 = v379
		v365 = v377
		v366 = v381
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v383 = v379
	v385 = v377
	goto L92
L95:
	;
	goto L94
L96:
	;
	v393 = v383
	v395 = v385
	goto L97
L97:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v400 = int32(1)
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399<<(uint(v400)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v395))) = uint16(v404)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406<<(uint(v400)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+2)) = uint16(v411)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+2)))
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413<<(uint(v400)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+4)) = uint16(v418)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+3)))
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420<<(uint(v400)%32))+uint32(_consts[399]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+6)) = uint16(v425)
	v430 = v393 + int32(4)
	if v430 != v359 {
		v393 = v430
		v395 = v395 + int32(8)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L88
L99:
	;
	goto L98
L100:
	;
	goto L65
L101:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	F_AppendStringToManifest(m, l0, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	F_pfree(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	goto L4
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
	F_errmsg_internal(m, int32(675172), v10)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(469715), int32(187), int32(74096))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateLockFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
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
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
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
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int64
	_ = v487
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(2608)
	m.G0 = v14
	v16 = int32(518894)
	v22 = F___strchrnul(m, v16, int32(61))
	mBase = m.M
	if v16 == v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v64 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = v22 - v16
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1104]))))
	if v27 != 0 {
		v57 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = v57
	goto L1
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1105]))
	if v29 == int32(0) {
		v57 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 == int32(0) {
		v57 = v6
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v36 = v29
	v37 = v32
	goto L9
L9:
	;
	v40 = F_strncmp(m, v16, v37, v25)
	mBase = m.M
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v57 = v44 + int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v44 = v43 + v25
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 == int32(61) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v49 != 0 {
		v36 = v36 + int32(4)
		v37 = v49
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v57 = v6
	goto L5
L17:
	;
	v68 = v64
	goto L21
L18:
	;
	v113 = v6
	goto L19
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v115
	v120 = F_open(m, l0, int32(194), v14+int32(288))
	mBase = m.M
	if v120 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v113 = v112
	goto L19
L21:
	;
	v73 = v68 + int32(1)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68))))
	v75 = F___isspace(m, v74)
	mBase = m.M
	if v75 != 0 {
		v68 = v73
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v76 = int32(1)
	switch v74&int32(255) - int32(43) {
	case 0:
		v82 = v76
		goto L25
	default:
		v84 = v74
		v85 = v68
		v86 = v76
		goto L24
	case 2:
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v87 = int32(0)
	v89 = v84 - int32(48)
	if base.Ui32(v89) <= base.Ui32(int32(9)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v84 = v83
	v85 = v73
	v86 = v82
	goto L24
L26:
	;
	v82 = int32(0)
	goto L25
L27:
	;
	v92 = v87
	v93 = v89
	v94 = v85
	goto L30
L28:
	;
	v106 = v87
	goto L29
L29:
	;
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v96 = int32(10)
	v98 = v92*v96 - v93
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94)+1)))
	v103 = v99 - int32(48)
	if base.Ui32(v103) < base.Ui32(v96) {
		v92 = v98
		v93 = v103
		v94 = v94 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v106 = v98
	goto L29
L32:
	;
	goto L31
L33:
	;
	v112 = int32(0) - v106
	goto L35
L34:
	;
	v112 = v106
	goto L35
L35:
	;
	goto L20
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L62
	} else {
		goto L261
	}
L37:
	;
	v845 = int32(4606604)
	v846 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v847 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v846
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L62
	} else {
		goto L257
	}
L38:
	;
	v822 = int32(4606604)
	v823 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v824 = F_close(m, v472)
	mBase = m.M
	v825 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v823
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L62
	} else {
		goto L253
	}
L39:
	;
	v797 = int32(4606604)
	v798 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v799 = F_close(m, v472)
	mBase = m.M
	v800 = F_unlink(m, l0)
	mBase = m.M
	if v798 != 0 {
		goto L246
	} else {
		goto L247
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L62
	} else {
		goto L241
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L62
	} else {
		goto L236
	}
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L62
	} else {
		goto L233
	}
L43:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L62
	} else {
		goto L228
	}
L44:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L62
	} else {
		goto L224
	}
L45:
	;
	v132 = int32(0)
	goto L48
L46:
	;
	v472 = v120
	goto L47
L47:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
	*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = l2
	v484 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v484
	v487 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v487
	if l1 != 0 {
		goto L162
	} else {
		goto L163
	}
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v136 != int32(20) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v472 = v463
	goto L47
L50:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v146
	v148 = int32(0)
	v151 = F_open(m, l0, v148, v14+int32(272))
	mBase = m.M
	if v151 < v148 {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	if v136 != int32(2) {
		goto L36
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v132) {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(v132) <= base.Ui32(int32(100)) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	goto L50
L57:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v456
	v463 = F_open(m, l0, int32(194), v14+int32(96))
	mBase = m.M
	if v463 < int32(0) {
		v132 = v132 + int32(1)
		goto L48
	} else {
		goto L161
	}
L58:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v155 == int32(44) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(167772190)
	v182 = F_read(m, v151, v14+int32(304), int32(2303))
	mBase = m.M
	if v182 < int32(0) {
		goto L44
	} else {
		goto L67
	}
L61:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = l0
	F_errmsg(m, int32(284040), v14+int32(112))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(470121), int32(1300), int32(370848))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v187
	v189 = F_close(m, v151)
	mBase = m.M
	if v182 == v187 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v193 = v14 + int32(304)
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193+v182))) = uint8(v195)
	v202 = v193
	goto L70
L69:
	;
	v248 = v246 >> (uint(int32(31)) % 32)
	v250 = v246 ^ v248 - v248
	if v250 <= int32(0) {
		goto L42
	} else {
		goto L85
	}
L70:
	;
	v207 = v202 + int32(1)
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v202))))
	v209 = F___isspace(m, v208)
	mBase = m.M
	if v209 != 0 {
		v202 = v207
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v210 = int32(1)
	switch v208&int32(255) - int32(43) {
	case 0:
		v216 = v210
		goto L74
	default:
		v218 = v208
		v219 = v202
		v220 = v210
		goto L73
	case 2:
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v221 = int32(0)
	v223 = v218 - int32(48)
	if base.Ui32(v223) <= base.Ui32(int32(9)) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v207))))
	v218 = v217
	v219 = v207
	v220 = v216
	goto L73
L75:
	;
	v216 = int32(0)
	goto L74
L76:
	;
	v226 = v221
	v227 = v223
	v228 = v219
	goto L79
L77:
	;
	v240 = v221
	goto L78
L78:
	;
	if v220 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v230 = int32(10)
	v232 = v226*v230 - v227
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v228)+1)))
	v237 = v233 - int32(48)
	if base.Ui32(v237) < base.Ui32(v230) {
		v226 = v232
		v227 = v237
		v228 = v228 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v240 = v232
	goto L78
L81:
	;
	goto L80
L82:
	;
	v246 = int32(0) - v240
	goto L84
L83:
	;
	v246 = v240
	goto L84
L84:
	;
	goto L69
L85:
	;
	if v250 == int32(42) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if l3 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L87:
	;
	if v250 == int32(1) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if v250 == v113 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v259 = F_kill(m, v250, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L62
	} else {
		goto L91
	}
L90:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L62
	} else {
		goto L93
	}
L91:
	;
	if v259 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	switch v264 - int32(63) {
	case 0, 8:
		goto L86
	default:
		goto L90
	}
L93:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L62
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = l0
	F_errmsg(m, int32(108972), v14+int32(256))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L62
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v250
	if l3 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v284 = int32(521606)
	goto L98
L97:
	;
	v284 = int32(521730)
	goto L98
L98:
	;
	if l3 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v287 = int32(521667)
	goto L101
L100:
	;
	v287 = int32(521783)
	goto L101
L101:
	;
	if v246 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v290 = v284
	goto L104
L103:
	;
	v290 = v287
	goto L104
L104:
	;
	F_errhint(m, v290, v14+int32(240))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L62
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(470121), int32(1372), int32(370848))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L62
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v449 = F_unlink(m, l0)
	mBase = m.M
	if v449 < int32(0) {
		goto L40
	} else {
		goto L160
	}
L108:
	;
	v304 = int32(10)
	v305 = F___strchrnul(m, v14+int32(304), v304)
	mBase = m.M
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v307 == v304 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v311 == int32(0) {
		goto L107
	} else {
		goto L113
	}
L110:
	;
	v311 = v305
	goto L112
L111:
	;
	v311 = int32(0)
	goto L112
L112:
	;
	goto L109
L113:
	;
	v316 = int32(10)
	v317 = F___strchrnul(m, v311+int32(1), v316)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v319 == v316 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v323 == int32(0) {
		goto L107
	} else {
		goto L118
	}
L115:
	;
	v323 = v317
	goto L117
L116:
	;
	v323 = int32(0)
	goto L117
L117:
	;
	goto L114
L118:
	;
	v328 = int32(10)
	v329 = F___strchrnul(m, v323+int32(1), v328)
	mBase = m.M
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	if v331 == v328 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v335 == int32(0) {
		goto L107
	} else {
		goto L123
	}
L120:
	;
	v335 = v329
	goto L122
L121:
	;
	v335 = int32(0)
	goto L122
L122:
	;
	goto L119
L123:
	;
	v340 = int32(10)
	v341 = F___strchrnul(m, v335+int32(1), v340)
	mBase = m.M
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v343 == v340 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v347 == int32(0) {
		goto L107
	} else {
		goto L128
	}
L125:
	;
	v347 = v341
	goto L127
L126:
	;
	v347 = int32(0)
	goto L127
L127:
	;
	goto L124
L128:
	;
	v352 = int32(10)
	v353 = F___strchrnul(m, v347+int32(1), v352)
	mBase = m.M
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v355 == v352 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v359 == int32(0) {
		goto L107
	} else {
		goto L133
	}
L130:
	;
	v359 = v353
	goto L132
L131:
	;
	v359 = int32(0)
	goto L132
L132:
	;
	goto L129
L133:
	;
	v364 = int32(10)
	v365 = F___strchrnul(m, v359+int32(1), v364)
	mBase = m.M
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v367 == v364 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v371 == int32(0) {
		goto L107
	} else {
		goto L138
	}
L135:
	;
	v371 = v365
	goto L137
L136:
	;
	v371 = int32(0)
	goto L137
L137:
	;
	goto L134
L138:
	;
	v375 = v371 + int32(1)
	if v375 == int32(0) {
		goto L107
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v14 + int32(296)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v14 + int32(300)
	v387 = F_sscanf(m, v375, int32(35885), v14+int32(224))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L62
	} else {
		goto L140
	}
L140:
	;
	if v387 != int32(2) {
		goto L107
	} else {
		goto L141
	}
L141:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	v393 = m.G0
	v395 = v393 - int32(16)
	m.G0 = v395
	v399 = F_PGSharedMemoryAttach(m, v392, v395+int32(12))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L62
	} else {
		goto L142
	}
L142:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	if v401 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	m.G0 = v395 + int32(16)
	if base.Ui32(v399) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L159
	}
L144:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _consts[1108]))
	if v406 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if int32(0) <= v423 {
		goto L143
	} else {
		goto L154
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	v423 = int32(-1)
	goto L145
L147:
	;
	v410 = v406
	goto L148
L148:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	if v401 != v411 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v423 = int32(0)
	goto L145
L150:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+20))
	if v413 != 0 {
		v410 = v413
		goto L148
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	goto L146
L154:
	;
	v428 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L62
	} else {
		goto L155
	}
L155:
	;
	if v428 == int32(0) {
		goto L143
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v401
	F_errmsg_internal(m, int32(280822), v395)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L62
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(473866), int32(324), int32(345164))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L62
	} else {
		goto L158
	}
L158:
	;
	goto L143
L159:
	;
	goto L107
L160:
	;
	goto L57
L161:
	;
	goto L49
L162:
	;
	v491 = int32(42)
	goto L164
L163:
	;
	v491 = int32(-42)
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v491
	v499 = F_pg_snprintf(m, v14+int32(304), int32(2304), int32(706094), v14+int32(48))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L62
	} else {
		goto L165
	}
L165:
	;
	if l1 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v516 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v516
	v519 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = int32(167772192)
	v523 = v14 + int32(304)
	if v523&int32(3) == v516 {
		v549 = v523
		goto L176
	} else {
		goto L177
	}
L167:
	;
	if l3 == int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v504 = v14 + int32(304)
	v505 = int32(715062)
	v506 = int32(2304)
	v508 = F_pg_ascii_verifystr(m, v504, v506)
	mBase = m.M
	if v508 == v506 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L166
L170:
	;
	goto L169
L171:
	;
	v510 = F_strlen(m, v505)
	mBase = m.M
	goto L170
L172:
	;
	goto L173
L173:
	;
	v513 = F_strlcpy(m, v504+v508, v505, v506-v508)
	mBase = m.M
	goto L170
L174:
	;
	v583 = F_write(m, v472, v523, v582)
	mBase = m.M
	v585 = v14 + int32(304)
	if v585&int32(3) == int32(0) {
		v609 = v585
		goto L193
	} else {
		goto L194
	}
L175:
	;
	v582 = v574 - v523
	goto L174
L176:
	;
	v553 = v549
	goto L185
L177:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v533 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v582 = int32(0)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v538 = v523
	goto L181
L181:
	;
	v542 = v538 + int32(1)
	if v542&int32(3) == int32(0) {
		v549 = v542
		goto L176
	} else {
		goto L183
	}
L182:
	;
	v574 = v542
	goto L175
L183:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	if v547 != 0 {
		v538 = v542
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v562 = int32(-2139062144)
	if (int32(16843008)-v559|v559)&v562 == v562 {
		v553 = v553 + int32(4)
		goto L185
	} else {
		goto L187
	}
L186:
	;
	v568 = v553
	goto L188
L187:
	;
	goto L186
L188:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v572 != 0 {
		v568 = v568 + int32(1)
		goto L188
	} else {
		goto L190
	}
L189:
	;
	v574 = v568
	goto L175
L190:
	;
	goto L189
L191:
	;
	if v583 != v642 {
		goto L39
	} else {
		goto L208
	}
L192:
	;
	v642 = v634 - v585
	goto L191
L193:
	;
	v613 = v609
	goto L202
L194:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	if v593 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v642 = int32(0)
	goto L191
L196:
	;
	goto L197
L197:
	;
	v598 = v585
	goto L198
L198:
	;
	v602 = v598 + int32(1)
	if v602&int32(3) == int32(0) {
		v609 = v602
		goto L193
	} else {
		goto L200
	}
L199:
	;
	v634 = v602
	goto L192
L200:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	if v607 != 0 {
		v598 = v602
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v622 = int32(-2139062144)
	if (int32(16843008)-v619|v619)&v622 == v622 {
		v613 = v613 + int32(4)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v628 = v613
	goto L205
L204:
	;
	goto L203
L205:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v632 != 0 {
		v628 = v628 + int32(1)
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v634 = v628
	goto L192
L207:
	;
	goto L206
L208:
	;
	v644 = int32(4062620)
	v645 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v645))) = v646
	v649 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v649))) = int32(167772191)
	v654 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v654 != int32(1) {
		v668 = v646
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v668 != 0 {
		goto L38
	} else {
		goto L216
	}
L210:
	;
	goto L209
L211:
	;
	goto L212
L212:
	;
	v659 = F_fsync(m, v472)
	mBase = m.M
	if v659 != int32(-1) {
		v668 = v659
		goto L210
	} else {
		goto L214
	}
L213:
	;
	v668 = int32(-1)
	goto L210
L214:
	;
	v663 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v663 == int32(27) {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v670))) = int32(0)
	v673 = F_close(m, v472)
	mBase = m.M
	if v673 != 0 {
		goto L37
	} else {
		goto L217
	}
L217:
	;
	v675 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
	if v675 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	F_on_proc_exit(m, int32(1638))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L62
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v682 = F_pstrdup(m, l0)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L62
	} else {
		goto L222
	}
L221:
	;
	goto L220
L222:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
	v686 = F_lcons(m, v682, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L62
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1109])) = v686
	m.G0 = v14 + int32(2608)
	return
L224:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L62
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = l0
	F_errmsg(m, int32(284185), v14+int32(128))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L62
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(470121), int32(1307), int32(370848))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L62
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L62
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = l0
	F_errmsg(m, int32(8559), v14+int32(144))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L62
	} else {
		goto L230
	}
L230:
	;
	F_errhint(m, int32(583517), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L62
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(470121), int32(1316), int32(370848))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L62
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v14 + int32(304)
	F_errmsg_internal(m, int32(686086), v14+int32(160))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L62
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(470121), int32(1327), int32(370848))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L62
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L62
	} else {
		goto L237
	}
L237:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v14)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v756
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v758
	F_errmsg(m, int32(342784), v14+int32(208))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L62
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = l4
	F_errhint(m, int32(623049), v14+int32(192))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L62
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(470121), int32(1410), int32(370848))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L62
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L62
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = l0
	F_errmsg(m, int32(284145), v14+int32(176))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L62
	} else {
		goto L243
	}
L243:
	;
	F_errhint(m, int32(578592), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L62
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(470121), int32(1426), int32(370848))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L62
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	v803 = v798
	goto L248
L247:
	;
	v803 = int32(51)
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v803
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L62
	} else {
		goto L249
	}
L249:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L62
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	F_errmsg(m, int32(284074), v14+int32(32))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L62
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(470121), int32(1461), int32(370848))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L62
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L62
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(284074), v14+int32(16))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L62
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(470121), int32(1475), int32(370848))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L62
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L62
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(284074), v14)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L62
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(470121), int32(1486), int32(370848))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L62
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L62
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l0
	F_errmsg(m, int32(284109), v14+int32(80))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L62
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(470121), int32(1286), int32(370848))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L62
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = F_FileAccess(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v16 != int32(1) {
		v38 = int32(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(0)
	return v38
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0*int32(48))))
	goto L8
L8:
	;
	v28 = F_fsync(m, v24)
	mBase = m.M
	if v28 != int32(-1) {
		v38 = v28
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(-1)
	goto L6
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v32 == int32(27) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
func F_load_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v90 = F_expand_dynamic_library_name(m, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L34
	}
L2:
	;
	v10 = int32(530488)
	goto L5
L3:
	;
	if v47-v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = l0
	v19 = v10
	v20 = int32(16)
	v21 = v17
	goto L11
L8:
	;
	v43 = v10
	v47 = int32(0)
	goto L9
L9:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	goto L3
L10:
	;
	v43 = v38
	v47 = v40
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 != v23 {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v38 = v32
	v40 = int32(0)
	goto L10
L13:
	;
	if v23 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v28 = v20 - int32(1)
	if v28 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(1)
	v32 = v19 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v33 != 0 {
		v18 = v18 + v31
		v19 = v32
		v20 = v28
		v21 = v33
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v61 = l0 + int32(16)
	goto L21
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	if v71 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v63 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	goto L22
L24:
	;
	v71 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v63 == int32(47) {
		v71 = v61
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v61 = v61 + int32(1)
	goto L21
L28:
	;
	goto L19
L29:
	;
	return
L30:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(418297), v6)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(472010), int32(528), int32(359406))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v92 = F_internal_load_library(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_sendFileWithContent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v18 = F_pg_checksum_init(m, v13+int32(24), v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L58
	}
L2:
	;
	return
L3:
	;
	if int32(0) <= v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l2&int32(3) == int32(0) {
		v45 = l2
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L55
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(123)
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v81
	v83 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = base.I64_extend_i32_s(v78)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v88
	F__tarWriteHeader(m, l0, l1, v81, v13+int32(32), v81)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L24
	}
L8:
	;
	v78 = v70 - l2
	goto L7
L9:
	;
	v49 = v45
	goto L18
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v29 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v34 = l2
	goto L14
L14:
	;
	v38 = v34 + int32(1)
	if v38&int32(3) == int32(0) {
		v45 = v38
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v70 = v38
	goto L8
L16:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v43 != 0 {
		v34 = v38
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 == v58 {
		v49 = v49 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v64 = v49
	goto L21
L20:
	;
	goto L19
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v68 != 0 {
		v64 = v64 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v70 = v64
	goto L8
L23:
	;
	goto L22
L24:
	;
	v98 = F_pg_checksum_update(m, v13+int32(24), l2, v78)
	mBase = m.M
	if v98 < int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if int32(0) < v78 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v105 = l2
	v110 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	v141 = (v78+int32(511))&int32(-512) - v78
	if int32(0) < v141 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = v78 - v110
	if base.Ui32(v114) < base.Ui32(v115) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v117 = v114
	goto L33
L32:
	;
	v117 = v115
	goto L33
L33:
	;
	if v117 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	m.T0[v121].(func(*base.Module, int32, int32))(m, l0, v117)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v118 = F__emscripten_memcpy_bulkmem(m, v113, v105, v117)
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	v125 = v117 + v110
	if v125 < v78 {
		v105 = v105 + v117
		v110 = v125
		goto L29
	} else {
		goto L39
	}
L39:
	;
	goto L30
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v144&int32(3) != 0 {
		v166 = v141
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	F_AddFileToBackupManifest(m, l3, int32(0), l1, v78, v83, v13+int32(24))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L54
	}
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	m.T0[v173].(func(*base.Module, int32, int32))(m, l0, v141)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L53
	}
L44:
	;
	v169 = F__emscripten_memset_bulkmem(m, v144, base.I32_extend8_s(int32(0)), v166)
	mBase = m.M
	goto L52
L45:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v141) {
		v166 = v141
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if v141&int32(3) != 0 {
		v166 = v141
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v151 = v141 + v144
	if base.Ui32(v151) <= base.Ui32(v144) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v156 = v144 + int32(4)
	if base.Ui32(v156) < base.Ui32(v151) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v158 = v151
	goto L51
L50:
	;
	v158 = v156
	goto L51
L51:
	;
	v166 = (v144^int32(-1)+v158)&int32(-4) + int32(4)
	goto L44
L52:
	;
	goto L43
L53:
	;
	goto L42
L54:
	;
	m.G0 = v13 + int32(128)
	return
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg_internal(m, int32(675213), v13)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(472560), int32(1084), int32(87707))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg_internal(m, int32(675256), v13+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(472560), int32(1109), int32(87707))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
