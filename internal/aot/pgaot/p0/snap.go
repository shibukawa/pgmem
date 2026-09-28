package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildSerialize(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	v12 = m.G0
	v14 = v12 - int32(2304)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L9
	} else {
		goto L103
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L9
	} else {
		goto L99
	}
L3:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	v388 = F_CloseTransientFile(m, v240)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L94
	}
L4:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	v361 = F_CloseTransientFile(m, v240)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L86
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L9
	} else {
		goto L82
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L78
	}
L7:
	;
	m.G0 = v14 + int32(2304)
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = int32(_a_F_SnapBuildSerialize_0)
	v21 = base.I32_wrap_i64(l1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v21
	v25 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = v25
	v28 = v14 + int32(256)
	v32 = F_pg_sprintf(m, v28, int32(_a_F_SnapBuildSerialize_1), v14+int32(144))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v38 = F___fstatat(m, int32(-100), v28, v14+int32(160), int32(0))
	mBase = m.M
	goto L12
L11:
	;
	v74 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L23
	}
L12:
	;
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	if v40 == int32(44) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_fsync_fname(m, v14+int32(256), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v28
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_2), v14+int32(128))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1550), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
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
	F_fsync_fname(m, int32(_a_F_SnapBuildSerialize_0), int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+136)) = l1
	goto L7
L23:
	;
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v14 + int32(256)
	F_errmsg_internal(m, int32(_a_F_SnapBuildSerialize_5), v14+int32(112))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = int32(_a_F_SnapBuildSerialize_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v21
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v94
	v97 = v14 + int32(1280)
	v101 = F_pg_sprintf(m, v97, int32(_a_F_SnapBuildSerialize_6), v14+int32(96))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L29
	}
L27:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1577), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v103 = F_unlink(m, v97)
	mBase = m.M
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	if v105 != int32(44) {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v108 = int32(_a_F_SnapBuildSerialize_7)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[2]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[2])) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v116 = F_palloc_mul(m, int32(4), v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L37
	}
L35:
	;
	v174 = int32(0)
	goto L36
L36:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v182 = (v176+v177)<<(uint(int32(2))%32) + int32(104)
	v183 = F_palloc0(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L45
	}
L37:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	if v119 == v118 {
		v150 = v118
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_pg_qsort(m, v116, v150, int32(4), int32(187))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L44
	}
L39:
	;
	v123 = v113 + int32(20)
	if v119 == v123 {
		v150 = v118
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v128 = v118
	v129 = v119
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129-int32(192))))
	*(*int32)(unsafe.Add(mBase, uint32(v116+v128<<(uint(int32(2))%32)))) = v141
	v144 = v128 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v145 != v123 {
		v128 = v144
		v129 = v145
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v150 = v144
	goto L38
L43:
	;
	goto L42
L44:
	;
	v174 = v116
	goto L36
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = int64(-2925404159)
	v191 = int32(8)
	v194 = m.Env.Pgmem_crc32c(m, int32(-1), v183+v191, v191)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v194
	v197 = v183 + int32(16)
	base.MemoryCopy(m, v197, l0, int32(80))
	v200 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+100)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v183)+92)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v183)+72)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v183)+20)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v183)+96)) = v176
	v212 = m.Env.Pgmem_crc32c(m, v194, v197, int32(88))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v212
	v215 = v183 + int32(104)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v216 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v218 = v216 << (uint(int32(2)) % 32)
	if v218 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v224 = v215
	v225 = v212
	goto L48
L48:
	;
	if v176 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	base.MemoryCopy(m, v215, v219, v218)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v221 = m.Env.Pgmem_crc32c(m, v212, v215, v218)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v221
	v224 = v215 + v218
	v225 = v221
	goto L48
L52:
	;
	v228 = v176 << (uint(int32(2)) % 32)
	if v228 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v233 = v225
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v233 ^ int32(-1)
	v238 = v14 + int32(1280)
	v240 = F_OpenTransientFile(m, v238, int32(193))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L9
	} else {
		goto L58
	}
L55:
	;
	base.MemoryCopy(m, v224, v174, v228)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v231 = m.Env.Pgmem_crc32c(m, v230, v224, v228)
	mBase = m.M
	v233 = v231
	goto L54
L58:
	;
	if v240 < int32(0) {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = int32(167772218)
	v251 = F_write(m, v240, v183, v182)
	mBase = m.M
	if v251 != v182 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v253 = int32(_a_F_SnapBuildSerialize_8)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v255
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = int32(167772217)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[4])))
	if v263 != int32(1) {
		v277 = v255
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v277 != 0 {
		goto L3
	} else {
		goto L68
	}
L62:
	;
	goto L61
L63:
	;
	goto L64
L64:
	;
	v268 = F_fsync(m, v240)
	mBase = m.M
	if v268 != int32(-1) {
		v277 = v268
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v277 = int32(-1)
	goto L62
L66:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	if v272 == int32(27) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = int32(0)
	v282 = F_CloseTransientFile(m, v240)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v282 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildSerialize_0), int32(1))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v289 = v14 + int32(256)
	v290 = F_rename(m, v238, v289)
	mBase = m.M
	if v290 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_fsync_fname(m, v289, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildSerialize_0), int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[2])) = v109
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+136)) = l1
	F_pfree(m, v183)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	if v174 == int32(0) {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v174)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	goto L7
L78:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_9), v14+int32(80))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1593), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_10), v14)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1655), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	if v360 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v365 = v360
	goto L89
L88:
	;
	v365 = int32(51)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = v365
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_11), v14-int32(-64))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1669), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = v387
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_12), v14+int32(48))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1693), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_13), v14+int32(32))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1700), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v14 + int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_14), v14+int32(16))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1713), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SnapBuildXactNeedsSkip(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	return base.B2i32(base.Ui64(l1) < base.Ui64(v3))
}
