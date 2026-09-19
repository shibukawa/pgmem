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
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
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
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	v4 = l3
	v5 = l4
	v6 = l5
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(2224)
	m.G0 = v16
	v19 = F_ReplicationSlotValidateName(m, l0, int32(21))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v5 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L101
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L97
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L93
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L85
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L81
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	v51 = F_LWLockAcquire(m, v47+int32(_a_F_ReplicationSlotCreate_0), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[1])))
	if v25 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[2]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+316))
	v33 = base.B2i32(v31 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[1])) = uint8(v33)
	v35 = v33
	goto L14
L13:
	;
	v35 = int32(0)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3])))
	if v37 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l2 != int32(2) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[3])))
	if v43 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	v58 = F_LWLockAcquire(m, v54+int32(_a_F_ReplicationSlotCreate_1), int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[4]))
	if v61 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	F_LWLockRelease(m, v65+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[5]))
	v78 = v7
	v80 = v7
	goto L27
L26:
	;
	goto L3
L27:
	;
	v87 = v71 + v80*int32(288)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	F_LWLockRelease(m, v127+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L47
	}
L29:
	;
	v92 = v87 + int32(24)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	if v78 != 0 {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	if v116-v117 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v101 = l0
	v102 = v92
	goto L35
L35:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v116 = v106
	v117 = v105
	goto L33
L37:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L31
L40:
	;
	v121 = v78
	goto L42
L41:
	;
	v121 = v87
	goto L42
L42:
	;
	if v88 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = v78
	goto L45
L44:
	;
	v122 = v121
	goto L45
L45:
	;
	v124 = v80 + int32(1)
	if v124 != v61 {
		v78 = v122
		v80 = v124
		goto L27
	} else {
		goto L46
	}
L46:
	;
	goto L28
L47:
	;
	if v122 == int32(0) {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v135 = v122 + int32(24)
	v136 = int32(0)
	base.MemoryFill(m, v135, v136, int32(184))
	v140 = F_strncpy(m, v135, l0, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+63)) = uint8(v136)
	goto L49
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[6]))
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+136)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+92)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+202)) = uint8(v5)
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+128)) = v148
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+201)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+236)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v122)+16)) = v148
	v155 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+12)) = uint16(v155)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+244)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v122)+252)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v122)+260)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v122)+268)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v122)+276)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v122)+284)) = v155
	if l1 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = v144
	goto L52
L51:
	;
	v170 = v155
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+88)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = int32(_a_F_ReplicationSlotCreate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v135
	v180 = F_pg_sprintf(m, v16+int32(176), int32(_a_F_ReplicationSlotCreate_3), v16-int32(-64))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = int32(_a_F_ReplicationSlotCreate_2)
	v186 = v16 + int32(1200)
	v190 = F_pg_sprintf(m, v186, int32(_a_F_ReplicationSlotCreate_4), v16+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v196 = F___fstatat(m, int32(-100), v186, v16+int32(80), int32(0))
	mBase = m.M
	goto L56
L55:
	;
	v205 = v16 + int32(1200)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[7]))
	v208 = F_mkdir(m, v205, v207)
	mBase = m.M
	goto L60
L56:
	;
	if v196 != 0 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	if v197&int32(_a_F_ReplicationSlotCreate_5) != int32(_a_F_ReplicationSlotCreate_6) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v202 = F_rmtree(m, v186)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	if v208 < int32(0) {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_fsync_fname(m, v205, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+13)) = uint8(v214)
	F_SaveSlotToPath(m, v122, v205, int32(21))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v220 = v16 + int32(176)
	v221 = F_rename(m, v205, v220)
	mBase = m.M
	if v221 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v222 = int32(_a_F_ReplicationSlotCreate_7)
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[8]))
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[8])) = v224 + v225
	F_fsync_fname(m, v220, v225)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_fsync_fname(m, int32(_a_F_ReplicationSlotCreate_2), int32(1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v235 = int32(_a_F_ReplicationSlotCreate_7)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[8])) = v237 - int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	v246 = F_LWLockAcquire(m, v242+int32(_a_F_ReplicationSlotCreate_1), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)) = uint8(v248)
	v252 = base.AtomicRmwXchg32(m, v122, int32(0), v248)
	if v252 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_s_lock(m, v122, int32(_a_F_ReplicationSlotCreate_8), int32(476), int32(_a_F_ReplicationSlotCreate_9))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v259
	v261 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v122))), uint32(v261))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[10])) = v122
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	F_LWLockRelease(m, v267+int32(_a_F_ReplicationSlotCreate_1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v122)+88))
	if v272 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[5]))
	v279 = base.I32_div_s(v122-v276, int32(288))
	goto L76
L74:
	;
	goto L75
L75:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCreate[0]))
	F_LWLockRelease(m, v294+int32(_a_F_ReplicationSlotCreate_0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L79
	}
L76:
	;
	v282 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v279), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	base.MemoryFill(m, v284+int32(24), int32(0), int32(72))
	F_pgstat_unlock_entry(m, v282)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	F_ConditionVariableBroadcast(m, v122+int32(224))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	m.G0 = v16 + int32(2224)
	return
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_10), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(377), int32(_a_F_ReplicationSlotCreate_9))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_11), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(389), int32(_a_F_ReplicationSlotCreate_9))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(_a_F_ReplicationSlotCreate_12))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_13), v16)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(414), int32(_a_F_ReplicationSlotCreate_9))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(1200)
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_14), v16+int32(16))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(2288), int32(_a_F_ReplicationSlotCreate_15))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v16 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v16 + int32(1200)
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_16), v16+int32(32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(2300), int32(_a_F_ReplicationSlotCreate_15))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
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
	F_errcode(m, int32(_a_F_ReplicationSlotCreate_17))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_ReplicationSlotCreate_18), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errhint(m, int32(_a_F_ReplicationSlotCreate_19), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotCreate_8), int32(425), int32(_a_F_ReplicationSlotCreate_9))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
