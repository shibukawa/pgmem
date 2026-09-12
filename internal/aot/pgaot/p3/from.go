package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AllocSetAllocFromNewBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = v13 - v14
	if base.Ui32(int32(16)) <= base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v15
	goto L4
L2:
	;
	goto L3
L3:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v87 = v85 << (uint(int32(1)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if base.Ui32(v87) < base.Ui32(v88) {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v39 = int32(8)
	v40 = v24 - v39
	if base.Ui32(v39) < base.Ui32(v40) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v43 = int32(29) - base.I32_clz(v24-int32(9))
	goto L8
L7:
	;
	v43 = int32(0)
	goto L8
L8:
	;
	v45 = int32(8)
	v49 = base.B2i32(v40 != v45<<(uint(v43)%32))
	if v40 != v45<<(uint(v43)%32) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v50 = int32(4)<<(uint(v43)%32) + v45
	goto L11
L10:
	;
	v50 = v24
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v31 + v50
	v53 = v43 - v49
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = base.I64_extend_i32_u(v53)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v31-v12)<<(uint(int64(34))%64) | int64(3)
	v67 = l0 + int32(48) + v53<<(uint(int32(2))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v31
	v71 = v24 - v50
	if base.Ui32(int32(15)) < base.Ui32(v71) {
		v24 = v71
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	v90 = v87
	goto L15
L14:
	;
	v90 = v88
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v90
	v93 = int32(8) << (uint(l3) % 32)
	v95 = v93 + int32(32)
	v102 = v85
	goto L16
L16:
	;
	if base.Ui32(v102) < base.Ui32(v95) {
		v102 = v102 << (uint(int32(1)) % 32)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v110 = F_emscripten_builtin_malloc(m, v102)
	mBase = m.M
	v112 = base.B2i32(v110 == int32(0))
	if base.Ui32(v102) < base.Ui32(int32(1048577)) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L17
L19:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v165 + v142
	*(*int32)(unsafe.Add(mBase, uint32(v141)+16)) = v141 + v142
	v171 = v141 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v176
	if v176 != 0 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	v160 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	if v144 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L22:
	;
	v141 = v110
	v142 = v102
	v144 = v112
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v110 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v141 = v110
	v142 = v102
	v144 = v112
	goto L21
L26:
	;
	goto L27
L27:
	;
	v119 = v102
	goto L28
L28:
	;
	v127 = int32(base.Ui32(v119) >> (uint(int32(1)) % 32))
	if base.Ui32(v127) < base.Ui32(v95) {
		goto L20
	} else {
		goto L30
	}
L29:
	;
	v141 = v129
	v142 = v127
	v144 = v131
	goto L21
L30:
	;
	v129 = F_emscripten_builtin_malloc(m, v127)
	mBase = m.M
	v131 = base.B2i32(v129 == int32(0))
	if base.Ui32(v119) < base.Ui32(int32(2097154)) {
		v141 = v129
		v142 = v127
		v144 = v131
		goto L21
	} else {
		goto L31
	}
L31:
	;
	if v129 == int32(0) {
		v119 = v127
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L20
L34:
	;
	return int32(0)
L35:
	;
	return v160
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v141
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v180 = v179
	goto L38
L37:
	;
	v180 = v171
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v141
	v183 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v180 + v93 + v183
	*(*int64)(unsafe.Add(mBase, uint32(v180))) = base.I64_extend_i32_u(l3<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v180-v141)<<(uint(int64(34))%64) | int64(3)
	return v180 + v183
}
func F_CopyFromBinaryInFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_getTypeBinaryInputInfo(m, l1, v7+int32(12), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		F_fmgr_info(m, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_CopyFromCSVOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v163 int64
	_ = v163
	var v183 int64
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int64
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	v29 = v27
	goto L3
L2:
	;
	v29 = int32(0)
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v34 != int64(0) {
		v183 = v34
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v514 = int32(0)
	F_errstart_cold(m, int32(21), v514)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L11
	} else {
		goto L119
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L115
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L11
	} else {
		goto L111
	}
L7:
	;
	m.G0 = v23 + int32(96)
	return v470
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v183 + int64(1)
	v188 = F_CopyReadLine(m, l0, int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L11
	} else {
		goto L42
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v38 == int32(0) {
		v183 = int64(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(1)
	v44 = F_CopyReadLine(m, l0, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v48 != int32(2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v44 != 0 {
		v470 = int32(0)
		goto L7
	} else {
		goto L40
	}
L14:
	;
	v51 = F_CopyReadAttributesCSV(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v51 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v51 != v58 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L4
L20:
	;
	v67 = int32(0)
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v83 <= v67 {
		goto L13
	} else {
		goto L23
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L36
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v90 = v67 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90+v91)))
	v96 = v33 - int32(80) + v85<<(uint(int32(4))%32) + v93*int32(100)
	v98 = v67 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90+v99)))
	if v101 == int32(0) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v105 = v96 + int32(4)
	if v105|v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v119 == int32(0) {
		v67 = v98
		goto L21
	} else {
		goto L35
	}
L26:
	;
	v111 = int32(-1)
	goto L28
L27:
	;
	v111 = int32(0)
	goto L28
L28:
	;
	if v105 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(1)
	goto L31
L30:
	;
	v112 = v111
	goto L31
L31:
	;
	if v105 == int32(0) {
		v119 = v112
		goto L32
	} else {
		goto L33
	}
L32:
	;
	goto L25
L33:
	;
	if v101 == int32(0) {
		v119 = v112
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v118 = F_strncmp(m, v105, v101, int32(64))
	mBase = m.M
	v119 = v118
	goto L32
L35:
	;
	goto L22
L36:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v98
	F_errmsg(m, int32(704563), v23+int32(80))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(495484), int32(826), int32(310765))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
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
	v163 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v183 = v163
	goto L8
L41:
	;
	v194 = F_CopyReadAttributesCSV(m, l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L45
	}
L42:
	;
	if v188 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v192 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v470 = int32(0)
	goto L7
L45:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if base.B2i32(int32(0) < v29)&base.B2i32(v29 < v194) != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v201 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v470 = int32(1)
	goto L7
L48:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v205 <= int32(0) {
		v470 = int32(1)
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v210 = int32(0)
	if v210 < v194 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v214 = v194
	goto L52
L51:
	;
	v214 = v210
	goto L52
L52:
	;
	v222 = v210
	goto L53
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v240 = v222 << (uint(int32(2)) % 32)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240+v241)))
	v245 = v243 - int32(1)
	v248 = v33 + int32(20) + v235<<(uint(int32(4))%32) + v245*int32(100)
	if v222 != v214 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L47
L55:
	;
	v426 = v222 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v426 < v427 {
		v222 = v426
		goto L53
	} else {
		goto L110
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = int64(0)
	goto L55
L57:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v240+v196)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v252 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L11
	} else {
		goto L106
	}
L60:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v252))))
	if v254 != int32(1) {
		goto L55
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v251 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	goto L62
L64:
	;
	v317 = v245 << (uint(int32(2)) % 32)
	v318 = l2 + v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v245))))
	if v321 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L65:
	;
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v245))) = uint8(v313)
	v315 = v311
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v248 + int32(4)
	v311 = v251
	goto L65
L67:
	;
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v248 + int32(4)
	v315 = v301
	goto L64
L68:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+v245))))
	if v261 == int32(0) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v245))))
	if v272 != int32(1) {
		goto L66
	} else {
		goto L73
	}
L71:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v248 + int32(4)
	if v264 != 0 {
		v311 = v264
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v315 = int32(0)
	goto L64
L73:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v279 == int32(0) {
		v298 = v278
		v299 = v279
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v299-v298 != 0 {
		goto L66
	} else {
		goto L82
	}
L75:
	;
	goto L74
L76:
	;
	if v278 != v279 {
		v298 = v278
		v299 = v279
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v283 = v251
	v284 = v275
	goto L78
L78:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v288 == int32(0) {
		v298 = v287
		v299 = v288
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v298 = v287
	v299 = v288
	goto L75
L80:
	;
	v291 = int32(1)
	if v287 == v288 {
		v283 = v283 + v291
		v284 = v284 + v291
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L67
L83:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v317+v30)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v328 = m.T0[v327].(func(*base.Module, int32, int32, int32) int32)(m, v325, l1, l3+v245)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v317+v31)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v248)+76))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v338 = F_InputFunctionCallSafe(m, v32+v245*int32(28), v315, v335, v336, v337, v318)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v328
	goto L56
L87:
	;
	if v338 != 0 {
		goto L56
	} else {
		goto L88
	}
L88:
	;
	v340 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v340 + int64(1)
	v344 = int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v345 != v344 {
		v470 = v344
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v348 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v348)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v350 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v399 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v399)
	goto L47
L91:
	;
	v351 = F_CopyLimitPrintoutLength(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v378 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L102
	}
L94:
	;
	v355 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	if v355 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v358
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v357
	F_errmsg(m, int32(709329), v23+int32(32))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L11
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_pfree(m, v351)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L11
	} else {
		goto L101
	}
L99:
	;
	F_errfinish(m, int32(495484), int32(1059), int32(31894))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	goto L90
L102:
	;
	if v378 == int32(0) {
		goto L90
	} else {
		goto L103
	}
L103:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v382
	F_errmsg(m, int32(64596), v23+int32(16))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(495484), int32(1066), int32(31894))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v248 + int32(4)
	F_errmsg(m, int32(691929), v23)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(495484), int32(977), int32(31894))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	goto L54
L111:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(272769), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(495484), int32(962), int32(31894))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v96 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v98
	F_errmsg(m, int32(704480), v23-int32(-64))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(495484), int32(819), int32(310765))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v522 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v524 = v523
	goto L123
L122:
	;
	v524 = v514
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v51
	F_errmsg(m, int32(475992), v23+int32(48))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(495484), int32(803), int32(310765))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromTextLikeStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v8 = F_palloc(m, int32(65537))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = int64(0)
			v13 = v8
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v14)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v13
			F_initStringInfo(m, l0+int32(288))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v21 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
					v25 = v22 << (uint(int32(16)) % 32)
				} else {
					v25 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v25 >> (uint(int32(16)) % 32)
				v31 = F_palloc(m, v25>>(uint(int32(14))%32))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v31
					return
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
		v13 = v12
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v14)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v13
		F_initStringInfo(m, l0+int32(288))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v25 = v22 << (uint(int32(16)) % 32)
			} else {
				v25 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v25 >> (uint(int32(16)) % 32)
			v31 = F_palloc(m, v25>>(uint(int32(14))%32))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v31
				return
			}
		}
	}
}
func F_get_from_clause_item(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v222 int32
	_ = v222
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
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
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v21 - int32(63) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L8
	} else {
		goto L191
	}
L2:
	;
	m.G0 = v15 + int32(112)
	return
L3:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v442+v443<<(uint(int32(2))%32)-int32(4))))
	v450 = int32(0)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v451&int32(1) != 0 {
		goto L119
	} else {
		goto L120
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v24<<(uint(int32(2))%32) - int32(4)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v28)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+124)))
	if v37 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_appendStringInfoString(m, v17, int32(725962))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	switch v43 {
	case 0:
		goto L13
	case 1:
		goto L19
	default:
		goto L14
	case 3:
		goto L18
	case 4:
		goto L17
	case 5:
		goto L16
	case 6:
		goto L15
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v348 != 0 {
		goto L2
	} else {
		goto L99
	}
L11:
	;
	F_get_column_alias_list(m, v32, l2)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L8
	} else {
		goto L98
	}
L12:
	;
	F_get_rte_alias(m, v36, v24, int32(0), l2)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L8
	} else {
		goto L97
	}
L13:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v305 = F_generate_relation_name(m, v303, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L8
	} else {
		goto L92
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L8
	} else {
		goto L89
	}
L15:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	v282 = F_quote_identifier(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L87
	}
L16:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L84
	}
L17:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	switch v265 {
	case 0:
		goto L81
	case 1:
		goto L80
	default:
		goto L12
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v61 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_get_query_def(m, v47, v17, v48, int32(0), int32(1), v51, v52, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)))
	if v248 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L24:
	;
	v237 = int32(0)
	goto L23
L25:
	;
	F_appendStringInfoString(m, v17, int32(667221))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L50
	}
L26:
	;
	F_appendStringInfoString(m, v17, int32(667123))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L47
	}
L27:
	;
	v75 = int32(0)
	goto L37
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v65 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if int32(0) < v61 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)))
	if v66 != 0 {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	F_get_rule_expr_funccall(m, v67, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v237 = v64
	goto L23
L36:
	;
	v135 = int32(0)
	goto L26
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60+v75<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(15) {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	v101 = int32(0)
	v104 = v101
	v110 = v101
	goto L43
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v94 != int32(2331) {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v97 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v99 = v75 + int32(1)
	if v99 != v61 {
		v75 = v99
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v104<<(uint(int32(2))%32))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	v122 = F_list_concat(m, v110, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L45
	}
L44:
	;
	v135 = v122
	goto L26
L45:
	;
	v125 = v104 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v125 < v126 {
		v104 = v125
		v110 = v122
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_get_rule_expr(m, v135, l2, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L24
L50:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v152 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L71
	}
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v155 <= int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	F_get_rule_expr_funccall(m, v160, l2)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v163 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_appendStringInfoString(m, v17, int32(725657))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v170 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v171 <= v170 {
		goto L51
	} else {
		goto L60
	}
L58:
	;
	F_get_from_clause_coldeflist(m, v159, int32(0), l2)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v182 = v170
	goto L61
L61:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v182<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, int32(727439))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L63
	}
L62:
	;
	goto L51
L63:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	F_get_rule_expr_funccall(m, v194, l2)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	if v197 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_appendStringInfoString(m, v17, int32(725657))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v205 = v182 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v205 < v206 {
		v182 = v205
		goto L61
	} else {
		goto L70
	}
L68:
	;
	F_get_from_clause_coldeflist(m, v190, int32(0), l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L62
L71:
	;
	goto L24
L72:
	;
	F_appendStringInfoString(m, v17, int32(505206))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_get_rte_alias(m, v36, v24, int32(0), l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if v237 == int32(0) {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	if v259 == int32(0) {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_get_from_clause_coldeflist(m, v237, v32, l2)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	goto L10
L80:
	;
	F_get_json_table(m, v264, l2, int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L83
	}
L81:
	;
	F_get_xmltable(m, v264, l2, int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	goto L12
L83:
	;
	goto L12
L84:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	F_get_values_def(m, v275, l2)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	goto L12
L87:
	;
	F_appendStringInfoString(m, v17, v282)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	goto L12
L89:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v290
	F_errmsg_internal(m, int32(484305), v15+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(490777), int32(12506), int32(289085))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v305
	if v302 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v310 = int32(738681)
	goto L95
L94:
	;
	v310 = int32(725105)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v310
	F_appendStringInfo(m, v17, int32(174963), v15+int32(48))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	goto L12
L97:
	;
	goto L11
L98:
	;
	goto L10
L99:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v349 == int32(0) {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = int32(2281)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v357 = int32(0)
	v363 = F_generate_function_name(m, v355, int32(1), v357, v15+int32(108), v357, v357, v357)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v363
	F_appendStringInfo(m, v352, int32(667471), v15+int32(32))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v371 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_appendStringInfoChar(m, v352, int32(41))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L113
	}
L104:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v374 <= int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	F_get_rule_expr(m, v378, l2, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	v382 = int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v383 <= v382 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v389 = v382
	goto L108
L108:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	F_appendStringInfoString(m, v352, int32(727439))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L8
	} else {
		goto L110
	}
L109:
	;
	goto L103
L110:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v398+v389<<(uint(int32(2))%32))))
	F_get_rule_expr(m, v405, l2, int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	v410 = v389 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v410 < v411 {
		v389 = v410
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	if v428 == int32(0) {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	F_appendStringInfoString(m, v352, int32(668305))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	F_get_rule_expr(m, v434, l2, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	F_appendStringInfoChar(m, v352, int32(41))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	goto L2
L118:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_get_from_clause_item(m, v473, l1, l2)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L127
	}
L119:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	switch v455 - int32(63) {
	case 0:
		v462 = v450
		goto L122
	case 1:
		goto L124
	default:
		goto L123
	}
L120:
	;
	v466 = v450
	goto L121
L121:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L126
	}
L122:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v463 == int32(0) {
		v471 = v462
		goto L118
	} else {
		goto L125
	}
L123:
	;
	v462 = int32(1)
	goto L122
L124:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454)+32))
	v462 = base.B2i32(v458 == int32(0))
	goto L122
L125:
	;
	v466 = v462
	goto L121
L126:
	;
	v471 = v466
	goto L118
L127:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v477 {
	case 0:
		goto L129
	case 1:
		v500 = int32(725781)
		goto L128
	case 2:
		goto L132
	case 3:
		goto L131
	default:
		goto L130
	}
L128:
	;
	F_appendContextKeyword(m, l2, v500, int32(-8), int32(8), int32(4))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L8
	} else {
		goto L139
	}
L129:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v498 != 0 {
		goto L136
	} else {
		goto L137
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L8
	} else {
		goto L133
	}
L131:
	;
	v500 = int32(725768)
	goto L128
L132:
	;
	v500 = int32(725806)
	goto L128
L133:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v484
	F_errmsg_internal(m, int32(482077), v15-int32(-64))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(490777), int32(12578), int32(289085))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v499 = int32(725811)
	goto L138
L137:
	;
	v499 = int32(725793)
	goto L138
L138:
	;
	v500 = v499
	goto L128
L139:
	;
	if v471 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v518 != 0 {
		goto L149
	} else {
		goto L150
	}
L141:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_from_clause_item(m, v515, l1, l2)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
	} else {
		goto L147
	}
L144:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_from_clause_item(m, v509, l1, l2)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L8
	} else {
		goto L146
	}
L146:
	;
	goto L140
L147:
	;
	goto L140
L148:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v633&int32(1) != 0 {
		goto L182
	} else {
		goto L183
	}
L149:
	;
	F_appendStringInfoString(m, v17, int32(668274))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L8
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v593 != 0 {
		goto L169
	} else {
		goto L170
	}
L152:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v449)+44))
	if v522 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L8
	} else {
		goto L165
	}
L154:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v526 <= int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v531 = F_quote_identifier(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	F_appendStringInfoString(m, v17, v531)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v535 <= int32(1) {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v539 = int32(1)
	goto L159
L159:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550+v539<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, int32(727439))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L8
	} else {
		goto L161
	}
L160:
	;
	goto L153
L161:
	;
	v558 = F_quote_identifier(m, v554)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	F_appendStringInfoString(m, v17, v558)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v563 = v539 + int32(1)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v563 < v564 {
		v539 = v563
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v581 == int32(0) {
		goto L148
	} else {
		goto L166
	}
L166:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v585 = F_quote_identifier(m, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v585
	F_appendStringInfo(m, v17, int32(197111), v15+int32(96))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	goto L148
L169:
	;
	F_appendStringInfoString(m, v17, int32(725763))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L8
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v615 == int32(0) {
		goto L148
	} else {
		goto L180
	}
L172:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v597&int32(1) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L8
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_get_rule_expr(m, v605, l2, int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L8
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v609&int32(1) != 0 {
		goto L148
	} else {
		goto L178
	}
L178:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L8
	} else {
		goto L179
	}
L179:
	;
	goto L148
L180:
	;
	F_appendStringInfoString(m, v17, int32(534158))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	goto L148
L182:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v636 == int32(0) {
		goto L2
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L8
	} else {
		goto L186
	}
L185:
	;
	goto L184
L186:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v642 == int32(0) {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)+12))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+4))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+12))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v649+v650<<(uint(int32(2))%32)-int32(4))))
	v657 = F_quote_identifier(m, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L8
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v657
	F_appendStringInfo(m, v17, int32(205217), v15+int32(80))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	F_get_column_alias_list(m, v449, l2)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L8
	} else {
		goto L190
	}
L190:
	;
	goto L2
L191:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v686
	F_errmsg_internal(m, int32(482882), v15)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(490777), int32(12646), int32(289085))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L8
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prepare_sort_from_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v257 float64
	_ = v257
	var v259 float64
	_ = v259
	var v261 float64
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v11 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = v29
	goto L3
L2:
	;
	v30 = v11
	goto L3
L3:
	;
	v33 = F_palloc(m, v30<<(uint(int32(1))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = v30 << (uint(int32(2)) % 32)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v41 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v43 = F_palloc(m, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l1 == int32(0) {
		v335 = l0
		v349 = v11
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L78
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L75
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v349
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v43
	m.G0 = v26 + int32(16)
	return v335
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v47 <= int32(0) {
		v335 = l0
		v349 = v11
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = l0
	v54 = l4
	v61 = v28
	v64 = v11
	goto L14
L14:
	;
	v74 = v64 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+41)))
	if v79 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v335 = v289
	v349 = v332
	goto L11
L16:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v314 = F_get_opfamily_member_for_cmptype(m, v312, v301, v301, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L72
	}
L17:
	;
	v224 = int32(0)
	v226 = F_find_computable_ec_member(m, v224, v78, v61, l2, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L57
	}
L18:
	;
	v193 = F_get_sortgroupref_tle(m, v82, v61)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L55
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	if v82 != 0 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errmsg_internal(m, int32(337072), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(493534), int32(6347), int32(112046))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v167 = v100
	goto L48
L27:
	;
	if v61 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3+v64<<(uint(int32(1))%32)))))
	if v61 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v100 < v101 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	if v145 == int32(0) {
		goto L17
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v111 <= int32(0) {
		v145 = int32(0)
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v145 = int32(0)
	goto L33
L37:
	;
	v114 = int32(0)
	if v114 < v111 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = v111
	goto L40
L39:
	;
	v117 = v114
	goto L40
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v122 = int32(0)
	goto L41
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v118+v122<<(uint(int32(2))%32))))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+8)))
	if v131 == v107&int32(65535) {
		v145 = v130
		goto L33
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v134 = v122 + int32(1)
	if v134 != v117 {
		v122 = v134
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v150 = F_find_ec_member_matching_expr(m, v78, v149, l2)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v150 == int32(0) {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	v289 = v50
	v293 = v54
	v299 = v145
	v300 = v61
	v301 = v154
	goto L16
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v167<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v184 = F_find_ec_member_matching_expr(m, v78, v183, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v289 = v50
	v293 = v54
	v299 = v182
	v300 = v61
	v301 = v192
	goto L16
L50:
	;
	if v184 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v189 = v167 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v189 < v190 {
		v167 = v189
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L49
L54:
	;
	goto L17
L55:
	;
	if v193 == int32(0) {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v289 = v50
	v293 = v54
	v299 = v193
	v300 = v61
	v301 = v200
	goto L16
L57:
	;
	if v226 == int32(0) {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	if v54&int32(1) != 0 {
		v268 = v50
		v271 = v61
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v274 = F_copyObjectImpl(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L66
	}
L60:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v233 - int32(332) {
	case 0, 1, 2, 3, 4, 28, 29, 30, 35, 38, 39, 40, 41:
		goto L61
	default:
		v268 = v50
		v271 = v61
		goto L59
	case 23:
		goto L62
	}
L61:
	;
	v239 = F_copyObjectImpl(m, v61)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+80)))
	if v236&int32(4) != 0 {
		v268 = v50
		v271 = v61
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+37)))
	v243 = F_palloc0(m, int32(80))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v245 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+72)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v243)+56)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v243)+52)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v243)+48)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v243)+44)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = int32(331)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v255
	v257 = *(*float64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v243)+8)) = v257
	v259 = *(*float64)(unsafe.Add(mBase, uint32(v50)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v243)+16)) = v259
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v50)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v243)+24)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+37)) = uint8(v241)
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+36)) = uint8(v245)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+32)) = v263
	v268 = v243
	v271 = v239
	goto L59
L66:
	;
	if v271 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+4)))
	v280 = v277 + int32(1)
	goto L69
L68:
	;
	v280 = int32(1)
	goto L69
L69:
	;
	v284 = F_makeTargetEntry(m, v274, base.I32_extend16_s(v280), int32(0), int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v286 = F_lappend(m, v271, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+44)) = v286
	v289 = v268
	v293 = int32(1)
	v299 = v284
	v300 = v286
	v301 = v230
	goto L16
L72:
	;
	if v314 == int32(0) {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v318 = int32(1)
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v33+v64<<(uint(v318)%32)))) = uint16(v321)
	*(*int32)(unsafe.Add(mBase, uint32(v39+v74))) = v314
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v74))) = v326
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64+v43))) = uint8(v329)
	v332 = v64 + v318
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v332 < v333 {
		v50 = v289
		v54 = v293
		v61 = v300
		v64 = v332
		goto L14
	} else {
		goto L74
	}
L74:
	;
	goto L15
L75:
	;
	F_errmsg_internal(m, int32(78220), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(493534), int32(6413), int32(112046))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
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
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v384
	F_errmsg_internal(m, int32(39557), v26)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(493534), int32(6453), int32(112046))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
