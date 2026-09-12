package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildRestore(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
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
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int64
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	v2 = l1
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == int32(2) {
		v379 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L95
	}
L2:
	;
	m.G0 = v14 + int32(112)
	return v379
L3:
	;
	v20 = v14 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = m.G0
	v24 = v22 - int32(1120)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = int32(119409)
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+88)) = uint32(v2)
	v30 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+84)) = uint32(v30)
	v37 = F_pg_sprintf(m, v24+int32(96), int32(238612), v24+int32(80))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(base.Ui32(v44^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		v379 = v3
		goto L2
	} else {
		goto L56
	}
L5:
	;
	return int32(0)
L6:
	;
	v44 = F_OpenTransientFile(m, v24+int32(96), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L52
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L48
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L44
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L40
	}
L11:
	;
	m.G0 = v24 + int32(1120)
	goto L4
L12:
	;
	if v44 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v49 == int32(44) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_fsync_fname(m, v24+int32(96), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v24 + int32(96)
	F_errmsg(m, int32(299339), v24-int32(-64))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(500890), int32(1763), int32(87424))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_fsync_fname(m, int32(119409), int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_SnapBuildRestoreContents(m, v44, v20, int32(16), v24+int32(96))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v85 != int32(1369563137) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v89 = v14 + int32(16)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v90 != int32(6) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v95 = m.Env.Pgmem_crc32c(m, int32(-1), v89, int32(8))
	mBase = m.M
	v97 = v14 + int32(24)
	F_SnapBuildRestoreContents(m, v44, v97, int32(88), v24+int32(96))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v104 = m.Env.Pgmem_crc32c(m, v95, v97, int32(88))
	mBase = m.M
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = v105 << (uint(int32(2)) % 32)
	v108 = F_MemoryContextAllocZero(m, v21, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	v117 = v104
	goto L29
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	if v120 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v108
	F_SnapBuildRestoreContents(m, v44, v108, v107, v24+int32(96))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v116 = m.Env.Pgmem_crc32c(m, v104, v115, v107)
	mBase = m.M
	v117 = v116
	goto L29
L32:
	;
	v122 = v120 << (uint(int32(2)) % 32)
	v123 = F_MemoryContextAllocZero(m, v21, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	v132 = v117
	goto L34
L34:
	;
	v135 = F_CloseTransientFile(m, v44)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v123
	F_SnapBuildRestoreContents(m, v44, v123, v122, v24+int32(96))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	v131 = m.Env.Pgmem_crc32c(m, v117, v130, v122)
	mBase = m.M
	v132 = v131
	goto L34
L37:
	;
	if v135 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v138 = v132 ^ int32(-1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v138 != v139 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L11
L40:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = int32(1369563137)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v24 + int32(96)
	F_errmsg(m, int32(48846), v24+int32(48))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(500890), int32(1784), int32(87424))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v24 + int32(96)
	F_errmsg(m, int32(48913), v24+int32(32))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(500890), int32(1790), int32(87424))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v24 + int32(96)
	F_errmsg(m, int32(300082), v24+int32(16))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(500890), int32(1822), int32(87424))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v24 + int32(96)
	F_errmsg(m, int32(53303), v24)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(500890), int32(1831), int32(87424))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v242 < int32(2) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	if v368 != 0 {
		goto L89
	} else {
		goto L90
	}
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v246))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v245)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v258 != 0 {
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v258 = base.B2i32(base.Ui32(v245) < base.Ui32(v246))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v258 = int32(base.Ui32(v245-v246) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v267
	if v267 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = int32(0)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v278 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v274
	goto L66
L68:
	;
	F_pfree(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v283
	v285 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v287 == v285 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v306 = F_MemoryContextAllocZero(m, v300, v301<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L77
	}
L73:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+30)))
	if v290 == int32(1) {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287)+44))
	v295 = v293 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+44)) = v295
	if v295 != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	F_pfree(m, v287)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = int32(5)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v314 = v306 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+12)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = v312
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+16)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v321 = v317 << (uint(int32(2)) % 32)
	if v321 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	F_pg_qsort(m, v323, v317, int32(4), int32(185))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L82
	}
L79:
	;
	v322 = F__emscripten_memcpy_bulkmem(m, v314, v319, v321)
	mBase = m.M
	v323 = v322
	goto L81
L80:
	;
	v323 = v314
	goto L81
L81:
	;
	goto L78
L82:
	;
	v328 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+64)) = v328
	*(*int64)(unsafe.Add(mBase, uint32(v306)+44)) = v328
	v332 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+32)) = v332
	*(*int64)(unsafe.Add(mBase, uint32(v306)+20)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v306)+27)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v306
	v339 = int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v306)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+44)) = v340 + v339
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v344)+136)) = v2
	goto L83
L83:
	;
	v348 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	if v348 == int32(0) {
		v379 = v339
		goto L2
	} else {
		goto L85
	}
L85:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v2)
	v354 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v354)
	F_errmsg(m, int32(513261), v14)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	F_errdetail(m, int32(581005), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(500890), int32(1918), int32(365260))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	v379 = v339
	goto L2
L89:
	;
	F_pfree(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	if v371 == int32(0) {
		v379 = v3
		goto L2
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	F_pfree(m, v371)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v379 = v3
	goto L2
L95:
	;
	F_errmsg_internal(m, int32(87149), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(500890), int32(344), int32(87711))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SnapBuildRestoreContents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4126764)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(167772214)
	v15 = F_read(m, l0, l1, l2)
	mBase = m.M
	v17 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
	if v15 != l2 {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		v23 = F_CloseTransientFile(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v15 < int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[140])) = v22
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
						F_errmsg(m, int32(300258), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(500890), int32(1951), int32(120969))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
						F_errmsg(m, int32(37473), v9+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errfinish(m, int32(500890), int32(1957), int32(120969))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
	} else {
		m.G0 = v9 + int32(32)
		return
	}
}
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v3 != int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v8 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v8
		if v8 == int32(0) {
			F_pfree(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(87149), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errfinish(m, int32(500890), int32(344), int32(87711))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
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
