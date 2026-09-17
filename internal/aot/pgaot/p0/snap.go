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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
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
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
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
	v434 = m.ExcPending
	if v434 != 0 {
		goto L9
	} else {
		goto L103
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L9
	} else {
		goto L99
	}
L3:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	v389 = F_CloseTransientFile(m, v241)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L9
	} else {
		goto L94
	}
L4:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	v362 = F_CloseTransientFile(m, v241)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L9
	} else {
		goto L86
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L82
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
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
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1546), int32(_a_F_SnapBuildSerialize_4))
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
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1573), int32(_a_F_SnapBuildSerialize_4))
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
	v117 = F_palloc(m, v114<<(uint(int32(2))%32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L37
	}
L35:
	;
	v175 = int32(0)
	goto L36
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v183 = (v177+v178)<<(uint(int32(2))%32) + int32(104)
	v184 = F_palloc0(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L45
	}
L37:
	;
	v119 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	if v120 == v119 {
		v151 = v119
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_pg_qsort(m, v117, v151, int32(4), int32(185))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L44
	}
L39:
	;
	v124 = v113 + int32(20)
	if v120 == v124 {
		v151 = v119
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v129 = v119
	v130 = v120
	goto L41
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130-int32(192))))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v129<<(uint(int32(2))%32)))) = v142
	v145 = v129 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v146 != v124 {
		v129 = v145
		v130 = v146
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v151 = v145
	goto L38
L43:
	;
	goto L42
L44:
	;
	v175 = v117
	goto L36
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = int64(-2925404159)
	v192 = int32(8)
	v195 = m.Env.Pgmem_crc32c(m, int32(-1), v184+v192, v192)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v195
	v198 = v184 + int32(16)
	base.MemoryCopy(m, v198, l0, int32(80))
	v201 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+100)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v184)+92)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v184)+72)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v184)+56)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v184)+96)) = v177
	v213 = m.Env.Pgmem_crc32c(m, v195, v198, int32(88))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v213
	v216 = v184 + int32(104)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v217 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v219 = v217 << (uint(int32(2)) % 32)
	if v219 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v225 = v216
	v226 = v213
	goto L48
L48:
	;
	if v177 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	base.MemoryCopy(m, v216, v220, v219)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v222 = m.Env.Pgmem_crc32c(m, v213, v216, v219)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v222
	v225 = v216 + v219
	v226 = v222
	goto L48
L52:
	;
	v229 = v177 << (uint(int32(2)) % 32)
	if v229 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v234 = v226
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v234 ^ int32(-1)
	v239 = v14 + int32(1280)
	v241 = F_OpenTransientFile(m, v239, int32(193))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L58
	}
L55:
	;
	base.MemoryCopy(m, v225, v175, v229)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v232 = m.Env.Pgmem_crc32c(m, v231, v225, v229)
	mBase = m.M
	v234 = v232
	goto L54
L58:
	;
	if v241 < int32(0) {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = int32(167772216)
	v252 = F_write(m, v241, v184, v183)
	mBase = m.M
	if v252 != v183 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v254 = int32(_a_F_SnapBuildSerialize_8)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	v256 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v256
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = int32(167772215)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[4])))
	if v264 != int32(1) {
		v278 = v256
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v278 != 0 {
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
	v269 = F_fsync(m, v241)
	mBase = m.M
	if v269 != int32(-1) {
		v278 = v269
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v278 = int32(-1)
	goto L62
L66:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0]))
	if v273 == int32(27) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(0)
	v283 = F_CloseTransientFile(m, v241)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v283 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildSerialize_0), int32(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v290 = v14 + int32(256)
	v291 = F_rename(m, v239, v290)
	mBase = m.M
	if v291 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_fsync_fname(m, v290, int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildSerialize_0), int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[2])) = v109
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v302)+136)) = l1
	F_pfree(m, v184)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	if v175 == int32(0) {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v175)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
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
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_9), v14+int32(80))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1589), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_10), v14)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1651), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
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
	if v361 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v366 = v361
	goto L89
L88:
	;
	v366 = int32(51)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = v366
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_11), v14-int32(-64))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1665), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildSerialize[0])) = v388
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_12), v14+int32(48))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1689), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(1280)
	F_errmsg(m, int32(_a_F_SnapBuildSerialize_13), v14+int32(32))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1696), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
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
	v447 = m.ExcPending
	if v447 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_SnapBuildSerialize_3), int32(1709), int32(_a_F_SnapBuildSerialize_4))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
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
