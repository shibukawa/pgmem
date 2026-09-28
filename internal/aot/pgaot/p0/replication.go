package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_EndReplicationCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = F_strlen(m, l0)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_EndReplicationCommand[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), l0, v3+int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v4 = l3
	v6 = l5
	v7 = l6
	v18 = m.G0
	v20 = v18 - int32(2240)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0])))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = int32(1)
	goto L3
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[1]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[2]))
	goto L4
L3:
	;
	v33 = F_ReplicationSlotValidateName(m, l0, v31, int32(21))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v31 = base.B2i32(v27 == v29)
	goto L3
L5:
	;
	return
L6:
	;
	if v6 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L117
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L5
	} else {
		goto L113
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L109
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L105
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L101
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L97
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	v65 = F_LWLockAcquire(m, v61+int32(_a_F_ReplicationSlotCreate_0), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L25
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[4])))
	if v39 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[5]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+308))
	v47 = base.B2i32(v45 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[4])) = uint8(v47)
	v49 = v47
	goto L18
L17:
	;
	v49 = int32(0)
	goto L18
L18:
	;
	goto L15
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[6])))
	if v51 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l2 != int32(2) {
		goto L13
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[6])))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	v72 = F_LWLockAcquire(m, v68+int32(_a_F_ReplicationSlotCreate_1), int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[7]))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[8]))
	v78 = v75 + v77
	if v78 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	F_LWLockRelease(m, v82+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l4 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L7
L31:
	;
	v88 = v75
	goto L33
L32:
	;
	v88 = int32(0)
	goto L33
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[9]))
	v102 = int32(0)
	v104 = int32(0)
	goto L34
L34:
	;
	v114 = v93 + v104*int32(296)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
	if v115 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	F_LWLockRelease(m, v159+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L63
	}
L36:
	;
	v119 = v114 + int32(24)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if base.B2i32(v122 == int32(0))|base.B2i32(v122 != v125) != 0 {
		v143 = v122
		v144 = v125
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	if v102 != 0 {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	if v143-v144 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	v128 = l0
	v129 = v119
	goto L42
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	if v133 == int32(0) {
		v143 = v133
		v144 = v132
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v143 = v133
	v144 = v132
	goto L40
L44:
	;
	v136 = int32(1)
	if v133 == v132 {
		v128 = v128 + v136
		v129 = v129 + v136
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L38
L47:
	;
	v148 = v102
	goto L49
L48:
	;
	v148 = v114
	goto L49
L49:
	;
	if v115 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v149 = v102
	goto L52
L51:
	;
	v149 = v148
	goto L52
L52:
	;
	if l4^int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v150 = v149
	goto L55
L54:
	;
	v150 = v102
	goto L55
L55:
	;
	if v77 <= v104 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = v149
	goto L58
L57:
	;
	v152 = v150
	goto L58
L58:
	;
	if v104 < v88+v77 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v154 = v152
	goto L61
L60:
	;
	v154 = v102
	goto L61
L61:
	;
	v156 = v104 + int32(1)
	if v156 != v78 {
		v102 = v154
		v104 = v156
		goto L34
	} else {
		goto L62
	}
L62:
	;
	goto L35
L63:
	;
	if v154 == int32(0) {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v167 = v154 + int32(24)
	v168 = int32(0)
	base.MemoryFill(m, v167, v168, int32(184))
	v172 = F_strncpy(m, v167, l0, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+63)) = uint8(v168)
	goto L65
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[10]))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+136)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+92)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+202)) = uint8(v6)
	v180 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v154)+128)) = v180
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+201)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v154)+236)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+16)) = v180
	v187 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+12)) = uint16(v187)
	*(*int64)(unsafe.Add(mBase, uint32(v154)+244)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+252)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+260)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+268)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+276)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v154)+284)) = v180
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v202 = v176
	goto L68
L67:
	;
	v202 = v187
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+88)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(_a_F_ReplicationSlotCreate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v167
	v212 = F_pg_sprintf(m, v20+int32(192), int32(_a_F_ReplicationSlotCreate_3), v20+int32(80))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(_a_F_ReplicationSlotCreate_2)
	v218 = v20 + int32(1216)
	v222 = F_pg_sprintf(m, v218, int32(_a_F_ReplicationSlotCreate_4), v20-int32(-64))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v228 = F___fstatat(m, int32(-100), v218, v20+int32(96), int32(0))
	mBase = m.M
	goto L72
L71:
	;
	v237 = v20 + int32(1216)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[11]))
	v240 = F_mkdir(m, v237, v239)
	mBase = m.M
	goto L76
L72:
	;
	if v228 != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	if v229&int32(_a_F_ReplicationSlotCreate_5) != int32(_a_F_ReplicationSlotCreate_6) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v234 = F_rmtree(m, v218)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	if v240 < int32(0) {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_fsync_fname(m, v237, int32(1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+13)) = uint8(v246)
	F_SaveSlotToPath(m, v154, v237, int32(21))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v252 = v20 + int32(192)
	v253 = F_rename(m, v237, v252)
	mBase = m.M
	if v253 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v254 = int32(_a_F_ReplicationSlotCreate_7)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[12]))
	v257 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[12])) = v256 + v257
	F_fsync_fname(m, v252, v257)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_fsync_fname(m, int32(_a_F_ReplicationSlotCreate_2), int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v267 = int32(_a_F_ReplicationSlotCreate_7)
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[12])) = v269 - int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	v278 = F_LWLockAcquire(m, v274+int32(_a_F_ReplicationSlotCreate_1), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v280)
	v284 = base.AtomicRmwXchg32(m, v154, int32(0), v280)
	if v284 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_s_lock(m, v154, int32(_a_F_ReplicationSlotCreate_8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v289
	v291 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v154))), uint32(v291))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[14])) = v154
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	F_LWLockRelease(m, v297+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v154)+88))
	if v302 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[9]))
	v309 = base.I32_div_s(v154-v306, int32(296))
	goto L92
L90:
	;
	goto L91
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3]))
	F_LWLockRelease(m, v324+int32(_a_F_ReplicationSlotCreate_0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L95
	}
L92:
	;
	v312 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v309), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	base.MemoryFill(m, v314+int32(24), int32(0), int32(96))
	F_pgstat_unlock_entry(m, v312)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	F_ConditionVariableBroadcast(m, v154+int32(224))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	m.G0 = v20 + int32(2240)
	return
L97:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_9), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(408), int32(_a_F_ReplicationSlotCreate_11))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_12), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(420), int32(_a_F_ReplicationSlotCreate_11))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(_a_F_ReplicationSlotCreate_13))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_14), v20)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(451), int32(_a_F_ReplicationSlotCreate_11))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(1216)
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_15), v20+int32(32))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(2491), int32(_a_F_ReplicationSlotCreate_16))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v20 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v20 + int32(1216)
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_17), v20+int32(48))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(2503), int32(_a_F_ReplicationSlotCreate_16))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
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
	F_errcode(m, int32(_a_F_ReplicationSlotCreate_18))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_19), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	if l4 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v455 = int32(_a_F_ReplicationSlotCreate_20)
	goto L122
L121:
	;
	v455 = int32(_a_F_ReplicationSlotCreate_21)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v455
	F_errhint(m, int32(_a_F_ReplicationSlotCreate_22), v20+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_10), int32(465), int32(_a_F_ReplicationSlotCreate_11))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
