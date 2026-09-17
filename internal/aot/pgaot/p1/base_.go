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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BaseInit[0])))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v106 = F_emscripten_builtin_malloc(m, int32(48))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[1])) = v106
	if v106 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L23
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L19
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v19 = F_open(m, int32(_a_F_BaseInit_0), int32(1089), v5+int32(-16))
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
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[2]))
	v28 = F_freopen(m, int32(_a_F_BaseInit_0), int32(_a_F_BaseInit_1), v27)
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BaseInit[3])))
	if v35&int32(1) == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[4]))
	v44 = F_freopen(m, int32(_a_F_BaseInit_0), int32(_a_F_BaseInit_1), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v44 == int32(0) {
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
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_BaseInit_0)
	F_errmsg(m, int32(_a_F_BaseInit_2), v7)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_BaseInit_3), int32(2132), int32(_a_F_BaseInit_4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_BaseInit_0)
	F_errmsg(m, int32(_a_F_BaseInit_5), v5+int32(-48))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_BaseInit_3), int32(2143), int32(_a_F_BaseInit_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
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
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_BaseInit_0)
	F_errmsg(m, int32(_a_F_BaseInit_6), v5+int32(-32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_BaseInit_3), int32(2156), int32(_a_F_BaseInit_4))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v126
	v128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v106)+36)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v106)+28)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v106)+20)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v106)+12)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v106)+4)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[5])) = int32(1)
	v143 = int32(_a_F_BaseInit_7)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[6]))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[6])) = v147
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[8]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v154 = F_dsa_attach_in_place(m, v152, v126)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L34
	}
L30:
	;
	F_errcode(m, int32(_a_F_BaseInit_8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_BaseInit_9), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_BaseInit_10), int32(917), int32(_a_F_BaseInit_11))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
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
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[9])) = v154
	F_dsa_pin_mapping(m, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[9]))
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[8]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v166 = F_dshash_attach(m, v160, int32(_a_F_BaseInit_12), v164, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[6])) = v144
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[10])) = v166
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[11]))
	v175 = v173
	v177 = int32(24)
	goto L37
L37:
	;
	v179 = int32(0)
	if v175 == v179 {
		v207 = v179
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v214 = v208
	goto L45
L39:
	;
	v208 = int32(1)
	v210 = v177 + v208
	if v210 != int32(33) {
		v175 = v207
		v177 = v210
		goto L37
	} else {
		goto L44
	}
L40:
	;
	v183 = v177 << (uint(int32(2)) % 32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v175+v183-int32(96))))
	if v187 == int32(0) {
		v207 = v175
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v190&int32(1) == int32(0) {
		v207 = v175
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[7]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	v200 = F_MemoryContextAlloc(m, v198, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+uint32(_c_F_BaseInit[12]))) = v200
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[11]))
	v207 = v204
	goto L39
L44:
	;
	goto L38
L45:
	;
	if base.Ui32(v214) <= base.Ui32(int32(12)) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	F_before_shmem_exit(m, int32(1197), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L58
	}
L47:
	;
	v247 = v214 + int32(1)
	if v247 != int32(33) {
		v214 = v247
		goto L45
	} else {
		goto L57
	}
L48:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+28))
	if v240 == int32(0) {
		goto L47
	} else {
		goto L55
	}
L49:
	;
	v239 = v214*int32(72) + int32(_a_F_BaseInit_13)
	goto L48
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(int32(8)) < base.Ui32(v214-int32(24)) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[11]))
	if v228 == int32(0) {
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v228+v214<<(uint(int32(2))%32)-int32(96))))
	if v236 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v239 = v236
	goto L48
L55:
	;
	m.T0[v240].(func(*base.Module))(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	goto L47
L57:
	;
	goto L46
L58:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[13]))
	if v255 != int32(12) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v301 = m.G0
	v303 = v301 - int32(48)
	m.G0 = v303
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BaseInit[3])))
	if v306 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L8
	} else {
		goto L71
	}
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[14]))
	if v259 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L59
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[15]))
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[16]))
	if base.Ui32(v265+int32(38)) <= base.Ui32(v263) {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[17]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[18])) = v272 + v263*int32(164)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[19]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	if v279 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	m.T0[v279].(func(*base.Module))(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_before_shmem_exit(m, int32(1066), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	goto L63
L71:
	;
	F_errmsg_internal(m, int32(_a_F_BaseInit_14), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_BaseInit_15), int32(227), int32(_a_F_BaseInit_16))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	m.G0 = v303 + int32(48)
	v344 = int32(_a_F_BaseInit_17)
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[20]))
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[20])) = v346 + int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[7]))
	v357 = F_AllocSetContextCreateInternal(m, v352, int32(_a_F_BaseInit_18), int32(0), int32(_a_F_BaseInit_19), int32(_a_F_BaseInit_20))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L81
	}
L75:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[13]))
	if v310 != int32(11) {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[7]))
	v320 = F_AllocSetContextCreateInternal(m, v315, int32(_a_F_BaseInit_21), int32(0), int32(_a_F_BaseInit_19), int32(_a_F_BaseInit_20))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L8
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[21])) = v320
	v323 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v320)+5)) = uint8(v323)
	*(*int64)(unsafe.Add(mBase, uint32(v303)+16)) = int64(137438953496)
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+40)) = v328
	v334 = F_hash_create(m, int32(_a_F_BaseInit_22), int32(100), v303, int32(1064))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[22])) = v334
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[23])) = int32(0)
	goto L74
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[24])) = v357
	v360 = int32(_a_F_BaseInit_17)
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[20]))
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[20])) = v362 - int32(1)
	F_on_proc_exit(m, int32(1122))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v369 = m.G0
	v371 = v369 - int32(48)
	m.G0 = v371
	v374 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[25])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[26])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[27])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[28])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[29])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[30])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[31])) = v374
	*(*int64)(unsafe.Add(mBase, _c_F_BaseInit[32])) = v374
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[33]))
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[16]))
	v404 = base.I32_div_s(v399, v401+int32(38))
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[34])) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v371)+16)) = int64(34359738372)
	v412 = F_hash_create(m, int32(_a_F_BaseInit_23), int32(100), v371, int32(40))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[35])) = v412
	F_on_shmem_exit(m, int32(1076), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v371 + int32(48)
	F_before_shmem_exit(m, int32(1090), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[36]))
	if v427 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[7]))
	v437 = F_AllocSetContextCreateInternal(m, v432, int32(_a_F_BaseInit_24), int32(0), int32(_a_F_BaseInit_19), int32(_a_F_BaseInit_20))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L89
	}
L87:
	;
	v440 = v427
	goto L88
L88:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[37]))
	if v442 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[36])) = v437
	v440 = v437
	goto L88
L90:
	;
	v446 = F_MemoryContextAllocZero(m, v440, int32(_a_F_BaseInit_25))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[38]))
	if v455 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[39])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[37])) = v446
	goto L92
L94:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[36]))
	v461 = F_MemoryContextAlloc(m, v459, int32(240))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L8
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[40]))
	if v470 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[41])) = int32(20)
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[38])) = v461
	goto L96
L98:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_BaseInit[36]))
	v477 = F_MemoryContextAllocZero(m, v475, int32(928))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v480 = m.G0
	v482 = v480 - int32(48)
	m.G0 = v482
	*(*int64)(unsafe.Add(mBase, uint32(v482)+16)) = int64(240518168596)
	v490 = F_hash_create(m, int32(_a_F_BaseInit_26), int32(16), v482, int32(40))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[40])) = v477
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BaseInit[42])) = v490
	m.G0 = v482 + int32(48)
	F_before_shmem_exit(m, int32(1026), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	return
}
