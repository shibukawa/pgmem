package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BaseInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v114 int64
	_ = v114
	var v126 int64
	_ = v126
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int64
	_ = v383
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1192])))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v105 = F_emscripten_builtin_malloc(m, int32(48))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[749])) = v105
	if v105 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L23
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L19
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L15
	}
L5:
	;
	m.G0 = v7 - int32(-64)
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(438)
	v19 = F_open(m, int32(4515392), int32(1089), v5+int32(-16))
	mBase = m.M
	if v19 < int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v22 = F_isatty(m, v19)
	mBase = m.M
	v23 = F_close(m, v19)
	mBase = m.M
	v27 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	v28 = F_freopen(m, int32(4515392), int32(508942), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v28 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v22 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v35 != int32(1) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	v42 = F_freopen(m, int32(4515392), int32(508942), v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L5
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(4515392)
	F_errmsg(m, int32(299484), v7)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(499610), int32(2132), int32(282592))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(4515392)
	F_errmsg(m, int32(294505), v5+int32(-48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(499610), int32(2143), int32(282592))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(4515392)
	F_errmsg(m, int32(293353), v5+int32(-32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(499610), int32(2156), int32(282592))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v159 = int32(4520560)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v163 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v163
	v167 = *(*int32)(unsafe.Add(mBase, _consts[938]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v170 = F_dsa_attach_in_place(m, v168, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L40
	}
L28:
	;
	if v105&int32(3) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[750])) = int32(1)
	goto L27
L32:
	;
	if base.Ui32(v105+int32(48)) <= base.Ui32(v105) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v126 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+4)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v105)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+36)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v105)+28)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v105)+20)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = v126
	goto L31
L35:
	;
	v114 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+4)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v105)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+36)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v105)+28)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v105)+20)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = v114
	goto L31
L36:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(501242), int32(917), int32(130892))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[932])) = v170
	F_dsa_pin_mapping(m, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[932]))
	v179 = *(*int32)(unsafe.Add(mBase, _consts[938]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v182 = F_dshash_attach(m, v176, int32(1656104), v180, int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v160
	*(*int32)(unsafe.Add(mBase, _consts[931])) = v182
	v189 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v191 = v189
	v193 = int32(24)
	goto L43
L43:
	;
	v195 = int32(0)
	if v191 == v195 {
		v223 = v195
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v230 = v224
	goto L51
L45:
	;
	v224 = int32(1)
	v226 = v193 + v224
	if v226 != int32(33) {
		v191 = v223
		v193 = v226
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v199 = v193 << (uint(int32(2)) % 32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v191+v199-int32(96))))
	if v203 == int32(0) {
		v223 = v191
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v206&int32(1) == int32(0) {
		v223 = v191
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	v216 = F_MemoryContextAlloc(m, v214, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+uint32(_consts[1193]))) = v216
	v220 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v223 = v220
	goto L45
L50:
	;
	goto L44
L51:
	;
	if base.Ui32(v230) <= base.Ui32(int32(12)) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	F_before_shmem_exit(m, int32(1213), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L63
	}
L53:
	;
	v264 = v230 + int32(1)
	if v264 != int32(33) {
		v230 = v264
		goto L51
	} else {
		goto L62
	}
L54:
	;
	v254 = v230*int32(72) + int32(1655136)
	goto L56
L55:
	;
	if base.Ui32(int32(8)) < base.Ui32(v230-int32(24)) {
		goto L53
	} else {
		goto L57
	}
L56:
	;
	if v254 == int32(0) {
		goto L53
	} else {
		goto L59
	}
L57:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	if v244 == int32(0) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v244+v230<<(uint(int32(2))%32)-int32(96))))
	v254 = v252
	goto L56
L59:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+28))
	if v257 == int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	m.T0[v257].(func(*base.Module))(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	goto L53
L62:
	;
	goto L52
L63:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v272 != int32(12) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v318 = m.G0
	v320 = v318 - int32(48)
	m.G0 = v320
	v323 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v323 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L8
	} else {
		goto L76
	}
L66:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v276 == int32(0) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L64
L69:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v282 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if base.Ui32(v282+int32(38)) <= base.Ui32(v280) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	*(*int32)(unsafe.Add(mBase, _consts[720])) = v289 + v280*int32(164)
	v295 = *(*int32)(unsafe.Add(mBase, _consts[721]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
	if v296 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	m.T0[v296].(func(*base.Module))(m)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_before_shmem_exit(m, int32(1066), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L8
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L68
L76:
	;
	F_errmsg_internal(m, int32(545723), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(494646), int32(227), int32(428183))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	m.G0 = v320 + int32(48)
	v361 = int32(4515212)
	v363 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v363 + int32(1)
	F_mdinit(m)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L86
	}
L80:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v327 != int32(11) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v337 = F_AllocSetContextCreateInternal(m, v332, int32(60565), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1194])) = v337
	v340 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+5)) = uint8(v340)
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = int64(137438953496)
	v345 = *(*int32)(unsafe.Add(mBase, _consts[1194]))
	*(*int32)(unsafe.Add(mBase, uint32(v320)+40)) = v345
	v351 = F_hash_create(m, int32(398523), int32(100), v320, int32(1064))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1195])) = v351
	*(*int32)(unsafe.Add(mBase, _consts[1196])) = int32(0)
	goto L79
L86:
	;
	v369 = int32(4515212)
	v371 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v371 - int32(1)
	F_on_proc_exit(m, int32(1119))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v378 = m.G0
	v380 = v378 - int32(48)
	m.G0 = v380
	v383 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[265])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[267])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[268])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[269])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[270])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[271])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[272])) = v383
	*(*int64)(unsafe.Add(mBase, _consts[273])) = v383
	v408 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v410 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v413 = base.I32_div_s(v408, v410+int32(38))
	*(*int32)(unsafe.Add(mBase, _consts[1197])) = v413
	*(*int64)(unsafe.Add(mBase, uint32(v380)+16)) = int64(34359738372)
	v421 = F_hash_create(m, int32(88217), int32(100), v380, int32(40))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v421
	F_on_shmem_exit(m, int32(1076), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	m.G0 = v380 + int32(48)
	F_before_shmem_exit(m, int32(1090), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
	if v436 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v446 = F_AllocSetContextCreateInternal(m, v441, int32(251722), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
	} else {
		goto L94
	}
L92:
	;
	v449 = v436
	goto L93
L93:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	if v451 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1198])) = v446
	v449 = v446
	goto L93
L95:
	;
	v455 = F_MemoryContextAllocZero(m, v449, int32(41300))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L8
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[257]))
	if v464 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[254])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[255])) = v455
	goto L97
L99:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
	v470 = F_MemoryContextAlloc(m, v468, int32(240))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L8
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _consts[1199]))
	if v479 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[256])) = int32(20)
	*(*int32)(unsafe.Add(mBase, _consts[257])) = v470
	goto L101
L103:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
	v486 = F_MemoryContextAllocZero(m, v484, int32(928))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L8
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v489 = m.G0
	v491 = v489 - int32(48)
	m.G0 = v491
	*(*int64)(unsafe.Add(mBase, uint32(v491)+16)) = int64(240518168596)
	v499 = F_hash_create(m, int32(324398), int32(16), v491, int32(40))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L8
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1199])) = v486
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[199])) = v499
	m.G0 = v491 + int32(48)
	F_before_shmem_exit(m, int32(1026), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	return
}
