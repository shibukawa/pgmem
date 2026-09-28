package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetXLogReceiptTime(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, _c_F_GetXLogReceiptTime[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReceiptTime[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v7 == int32(3)))
	return
}
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
	v9 = base.AtomicRmwXchg32(m, v6, int32(96), int32(1))
	if v9 != 0 {
		F_s_lock(m, v6+int32(96), int32(_a_F_GetXLogReplayRecPtr_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
			v21 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v21))
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
			} else {
			}
			return v20
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
		v21 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v21))
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
		} else {
		}
		return v20
	}
}
func F_WaitXLogInsertionsToFinish(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v123 int64
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int64
	_ = v300
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v338 int64
	_ = v338
	var v343 int64
	_ = v343
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	v1 = l0
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[0]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[1]))
	v25 = base.AtomicRmwOr64(m, v22, int32(256), int64(0))
	if base.Ui64(v1) <= base.Ui64(v25) {
		v343 = v25
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L76
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v343
L5:
	;
	v29 = base.AtomicRmwXchg32(m, v22, int32(0), int32(1))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_s_lock(m, v22, int32(_a_F_WaitXLogInsertionsToFinish_0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v36 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22))), uint32(v36))
	v40 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[2])))
	v41 = base.I64_div_u_s(v35, v40)
	v43 = v35 - v41*v40
	if base.Ui64(v43) <= base.Ui64(int64(8151)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	return int64(0)
L10:
	;
	goto L8
L11:
	;
	v73 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[3])))
	v77 = v41*v73 + v71&int64(4294967295)
	if base.Ui64(v1) <= base.Ui64(v77) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v46 = int64(0)
	if v43 == v46 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v53 = v43 - int64(8152)
	v54 = int64(8168)
	v55 = base.I64_div_u_s(v53, v54)
	v57 = v55 << (uint(int64(13)) % 64)
	v62 = v53 - v55*v54
	if v62 == int64(0) {
		v71 = v57 - int64(-8192)
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v51 = v46
	goto L17
L16:
	;
	v51 = v43 + int64(40)
	goto L17
L17:
	;
	v71 = v51
	goto L11
L18:
	;
	v71 = v62 + v57 + int64(8216)
	goto L11
L19:
	;
	v105 = v77
	v112 = int32(0)
	goto L27
L20:
	;
	v101 = v1
	goto L19
L21:
	;
	goto L22
L22:
	;
	v81 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if v81 == int32(0) {
		v101 = v77
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+12)) = uint32(v77)
	v86 = int64(32)
	v87 = int64(base.Ui64(v77) >> (uint(v86) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+8)) = uint32(v87)
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)) = uint32(v1)
	v91 = int64(base.Ui64(v1) >> (uint(v86) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17))) = uint32(v91)
	F_errmsg(m, int32(_a_F_WaitXLogInsertionsToFinish_1), v17)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_2), int32(1586), int32(_a_F_WaitXLogInsertionsToFinish_3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v101 = v77
	goto L19
L27:
	;
	v116 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v116
	v123 = v116
	goto L30
L28:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[1]))
	v314 = int64(0)
	v317 = base.AtomicRmwCmpxchg64(m, v313, int32(256), v314, v314)
	if base.Ui64(v307) <= base.Ui64(v317) {
		goto L67
	} else {
		goto L68
	}
L29:
	;
	v309 = v112 + int32(1)
	if v309 != int32(8) {
		v105 = v307
		v112 = v309
		goto L27
	} else {
		goto L66
	}
L30:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[4]))
	v137 = v136 + v112<<(uint(int32(7))%32)
	v139 = v137 + int32(16)
	v141 = v17 + int32(24)
	v142 = int32(0)
	v143 = int32(_a_F_WaitXLogInsertionsToFinish_4)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v145 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[0]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v151&int32(_a_F_WaitXLogInsertionsToFinish_5) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v300 == int64(0) {
		v307 = v105
		goto L29
	} else {
		goto L62
	}
L32:
	;
	if int32(0) < v248 {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v159 = v142
	goto L36
L34:
	;
	v233 = v142
	goto L35
L35:
	;
	v248 = v233
	v251 = int32(1)
	goto L32
L36:
	;
	v168 = int64(0)
	v171 = base.AtomicRmwCmpxchg64(m, v139, int32(0), v168, v168)
	if v123 != v171 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v233 = v209
	goto L35
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v141))) = v171
	v248 = v159
	v251 = int32(0)
	goto L32
L39:
	;
	goto L40
L40:
	;
	F_LWLockQueueSelf(m, v137, int32(2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v180 = base.AtomicRmwAnd32(m, v137, int32(4), int32(-1073741825))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v183 = v181 & int32(_a_F_WaitXLogInsertionsToFinish_5)
	if v183 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137))))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v197 | int32(16777216)
	v209 = v159
	goto L48
L43:
	;
	v184 = int64(0)
	v187 = base.AtomicRmwCmpxchg64(m, v139, int32(0), v184, v184)
	if v187 == v123 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_LWLockDequeueSelf(m, v137)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L47
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v141))) = v187
	goto L45
L47:
	;
	v248 = v159
	v251 = base.B2i32(v183 == int32(0))
	goto L32
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v150)+332))
	F_PGSemaphoreLock(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L50
	}
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v225&int32(_a_F_WaitXLogInsertionsToFinish_5) != 0 {
		v159 = v209
		goto L36
	} else {
		goto L52
	}
L50:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+344)))
	if v220 != 0 {
		v209 = v209 + int32(1)
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L37
L53:
	;
	v264 = v248
	goto L56
L54:
	;
	goto L55
L55:
	;
	v294 = int32(_a_F_WaitXLogInsertionsToFinish_4)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v296 - int32(1)
	if v251 != 0 {
		v307 = v105
		goto L29
	} else {
		goto L60
	}
L56:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v150)+332))
	F_PGSemaphoreUnlock(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	v276 = int32(1)
	if base.Ui32(v276) < base.Ui32(v264) {
		v264 = v264 - v276
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	if base.Ui64(v300) < base.Ui64(v101) {
		v123 = v300
		goto L30
	} else {
		goto L61
	}
L61:
	;
	goto L31
L62:
	;
	if base.Ui64(v300) < base.Ui64(v105) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v305 = v300
	goto L65
L64:
	;
	v305 = v105
	goto L65
L65:
	;
	v307 = v305
	goto L29
L66:
	;
	goto L28
L67:
	;
	v319 = int32(0)
	v322 = base.AtomicRmwOr32(m, v319, int32(_a_F_WaitXLogInsertionsToFinish_6), v319)
	v343 = v317
	goto L4
L68:
	;
	goto L69
L69:
	;
	v324 = v317
	goto L70
L70:
	;
	v338 = base.AtomicRmwCmpxchg64(m, v313, int32(256), v324, v307)
	if v324 == v338 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v343 = v338
	goto L4
L72:
	;
	v343 = v307
	goto L4
L73:
	;
	goto L74
L74:
	;
	if base.Ui64(v338) < base.Ui64(v307) {
		v324 = v338
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	F_errmsg_internal(m, int32(_a_F_WaitXLogInsertionsToFinish_7), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_2), int32(1558), int32(_a_F_WaitXLogInsertionsToFinish_3))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogArchiveForceDone(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v5 = m.G0
	v7 = v5 - int32(2208)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = int32(_a_F_XLogArchiveForceDone_0)
	v13 = v7 + int32(160)
	v18 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_XLogArchiveForceDone_1), v7+int32(48))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = v7 - int32(-64)
		v24 = F___fstatat(m, int32(-100), v13, v21, int32(0))
		mBase = m.M
		if v24 == int32(0) {
			m.G0 = v7 + int32(2208)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(_a_F_XLogArchiveForceDone_2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
			v31 = v7 + int32(1184)
			v36 = F_pg_snprintf(m, v31, int32(1024), int32(_a_F_XLogArchiveForceDone_1), v7+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v40 = F___fstatat(m, int32(-100), v31, v21, int32(0))
				mBase = m.M
				if v40 == int32(0) {
					v44 = F_durable_rename(m, v31, v13, int32(19))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(2208)
						return
					}
				} else {
					v47 = v7 + int32(160)
					v49 = F_AllocateFile(m, v47, int32(_a_F_XLogArchiveForceDone_3))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						if v49 == int32(0) {
							v55 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								if v55 == int32(0) {
									m.G0 = v7 + int32(2208)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v47
										F_errmsg(m, int32(_a_F_XLogArchiveForceDone_4), v7)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(538), int32(_a_F_XLogArchiveForceDone_6))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												m.G0 = v7 + int32(2208)
												return
											}
										}
									}
								}
							}
						} else {
							v70 = F_FreeFile(m, v49)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								if v70 == int32(0) {
									m.G0 = v7 + int32(2208)
									return
								} else {
									v76 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										if v76 == int32(0) {
											m.G0 = v7 + int32(2208)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(160)
												F_errmsg(m, int32(_a_F_XLogArchiveForceDone_7), v7+int32(16))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(546), int32(_a_F_XLogArchiveForceDone_6))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														m.G0 = v7 + int32(2208)
														return
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
		}
	}
}
func F_XLogArchiveIsBusy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(1184)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a_F_XLogArchiveIsBusy_0)
	v14 = v8 + int32(160)
	v19 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(48))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = v8 - int32(-64)
		v27 = F___fstatat(m, int32(-100), v14, v24, int32(0))
		mBase = m.M
		if v27 == int32(0) {
			v73 = v2
			m.G0 = v8 + int32(1184)
			return v73
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_XLogArchiveIsBusy_2)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
			v37 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(32))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v41 = F___fstatat(m, int32(-100), v14, v24, int32(0))
				mBase = m.M
				if v41 == int32(0) {
					v73 = int32(1)
					m.G0 = v8 + int32(1184)
					return v73
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_XLogArchiveIsBusy_0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v51 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(16))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = F___fstatat(m, int32(-100), v14, v24, int32(0))
						mBase = m.M
						if v55 == int32(0) {
							v73 = v2
							m.G0 = v8 + int32(1184)
							return v73
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v61 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_3), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v65 = F___fstatat(m, int32(-100), v14, v24, int32(0))
								mBase = m.M
								if v65 == int32(0) {
									v73 = int32(1)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveIsBusy[0]))
									if v69 == int32(44) {
										v73 = v2
									} else {
										v73 = int32(1)
									}
								}
								m.G0 = v8 + int32(1184)
								return v73
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogBeginRead(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v23
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v27
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v25)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v25)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v23
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v12 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_pfree(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v18 = v10
	goto L8
L8:
	;
	if v18 != 0 {
		v8 = v18
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v18 = v17
	goto L8
L11:
	;
	goto L5
}
func F_XLogBytePosToRecPtr(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	v4 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogBytePosToRecPtr[0])))
	v5 = base.I64_div_u_s(l0, v4)
	v7 = l0 - v5*v4
	if base.Ui64(v7) <= base.Ui64(int64(8151)) {
		v25 = v7 + int64(40)
	} else {
		v13 = v7 - int64(8152)
		v14 = int64(8168)
		v15 = base.I64_div_u_s(v13, v14)
		v25 = v13 - v15*v14 + v15<<(uint(int64(13))%64) + int64(8216)
	}
	v27 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogBytePosToRecPtr[1])))
	return v5*v27 + v25&int64(4294967295)
}
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	if l0 < int32(33) {
		v11 = int32(4)
		if l0 <= v11 {
			v14 = v11
		} else {
			v14 = l0
		}
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0]))
		if v16 <= v14 {
			v18 = int32(_a_F_XLogEnsureRecordSpace_0)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[1]))
			v22 = v14 + int32(1)
			v25 = F_repalloc(m, v20, v22*int32(_a_F_XLogEnsureRecordSpace_1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[1])) = v25
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0]))
				v31 = int32(_a_F_XLogEnsureRecordSpace_1)
				v32 = (v22 - v29) * v31
				v34 = v29 * v31
				v35 = v25 + v34
				if v35&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v32)) == int32(0) {
					if v29 == v22 {
					} else {
						v46 = int32(_a_F_XLogEnsureRecordSpace_1)
						v50 = v14*v46 + v25 + v46
						v53 = v25 + v34 + int32(4)
						if base.Ui32(v53) < base.Ui32(v50) {
							v55 = v50
						} else {
							v55 = v53
						}
						v61 = (v25^int32(-1)+v55-v34)&int32(-4) + int32(4)
						if v61 == int32(0) {
						} else {
							base.MemoryFill(m, v35, int32(0), v61)
						}
					}
				} else {
					if v32 == int32(0) {
					} else {
						base.MemoryFill(m, v35, int32(0), v32)
					}
				}
				*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0])) = v22
				v81 = int32(20)
				if l1 <= v81 {
					v84 = v81
				} else {
					v84 = l1
				}
				v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2]))
				if v86 < v84 {
					v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3]))
					v92 = F_repalloc(m, v89, v84*int32(12))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2])) = v84
						*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3])) = v92
						return
					}
				} else {
					return
				}
			}
		} else {
			v81 = int32(20)
			if l1 <= v81 {
				v84 = v81
			} else {
				v84 = l1
			}
			v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2]))
			if v86 < v84 {
				v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3]))
				v92 = F_repalloc(m, v89, v84*int32(12))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2])) = v84
					*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3])) = v92
					return
				}
			} else {
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_XLogEnsureRecordSpace_2), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_XLogEnsureRecordSpace_3), int32(198), int32(_a_F_XLogEnsureRecordSpace_4))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
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
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	v15 = base.I64_div_u_s(int64(4294967296), base.I64_extend_i32_s(l3))
	v16 = base.I64_div_u_s(l2, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v16)
	v19 = l2 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+8)) = uint32(v19)
	v23 = F_pg_snprintf(m, l0, int32(64), int32(_a_F_XLogFileName_0), v10)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		m.G0 = v10 + int32(16)
		return
	}
}
func F_XLogGetOldestSegno(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	v5 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_AllocateDir(m, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = F_ReadDir(m, v12, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v17
	v24 = v5
	goto L7
L5:
	;
	v152 = v5
	goto L6
L6:
	;
	F_FreeDir(m, v12)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v26 = v21 + int32(19)
	v27 = F_strlen(m, v26)
	mBase = m.M
	if v27 != int32(24) {
		v143 = v24
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v143
	goto L6
L9:
	;
	v145 = F_ReadDir(m, v12, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L36
	}
L10:
	;
	v30 = int32(_a_F_XLogGetOldestSegno_1)
	v34 = m.G0
	v36 = v34 - int32(32)
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[0])))
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v113 != int32(24) {
		v143 = v24
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v113 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[1])))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = v26
	goto L18
L16:
	;
	goto L17
L17:
	;
	v63 = v30
	v64 = v45
	goto L21
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == v45 {
		v53 = v53 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v113 = v53 - v26
	goto L11
L20:
	;
	goto L19
L21:
	;
	v71 = v36 + int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v73<<(uint(v64)%32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v77 != 0 {
		v63 = v63 + v73
		v64 = v77
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v80 == int32(0) {
		v103 = v26
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v113 = v103 - v26
	goto L11
L25:
	;
	v84 = v26
	v85 = v80
	goto L26
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(base.Ui32(v85)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v93)>>(uint(v85)%32))&int32(1) == int32(0) {
		v103 = v84
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v103 = v101
	goto L24
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v101 = v84 + int32(1)
	if v99 != 0 {
		v84 = v101
		v85 = v99
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v117 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(24)
	v128 = F_sscanf(m, v26, int32(_a_F_XLogGetOldestSegno_2), v9)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if l0 != v130 {
		v143 = v24
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)))
	v135 = base.I64_div_u_s(int64(4294967296), v117)
	v137 = v132 + v133*v135
	if base.Ui64(v24-int64(1)) < base.Ui64(v137) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = v24
	goto L35
L34:
	;
	v141 = v137
	goto L35
L35:
	;
	v143 = v141
	goto L9
L36:
	;
	if v145 != 0 {
		v21 = v145
		v24 = v143
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	m.G0 = v9 + int32(32)
	return v152
}
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = F_XLogReadBufferForRedoExtended(m, l0, l1, int32(1), int32(0), v6+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		m.G0 = v6 + int32(16)
		return v16
	}
}
func F_XLogNeedsFlush(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0]))
	if v5 < int32(0) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])))
		if v9 == int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+308))
			v17 = base.B2i32(v15 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])) = uint8(v17)
			if v15 != int32(2) {
				v24 = int32(0)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
				if v26 != 0 {
					v110 = v24
					return v110
				} else {
					v28 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
					if base.Ui64(l0) <= base.Ui64(v28) {
						v110 = v24
						return v110
					} else {
						if v28 != int64(0) {
							v43 = int32(1)
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
							v49 = F_LWLockConditionalAcquire(m, v45+int32(1152), v43)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v49 == int32(0) {
									v110 = v43
									return v110
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
									v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+144))
									*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v58
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
									F_LWLockRelease(m, v61+int32(1152))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v67 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
										if v67 != int64(0) {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
											v76 = v71
										} else {
											v73 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v73)
											v76 = v73
										}
										v110 = (v76 ^ int32(1)) & base.B2i32(base.Ui64(v67) < base.Ui64(l0))
										return v110
									}
								}
							}
						} else {
							v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7])))
							if v33&int32(1) == int32(0) {
								v43 = int32(1)
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
								v49 = F_LWLockConditionalAcquire(m, v45+int32(1152), v43)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v49 == int32(0) {
										v110 = v43
										return v110
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
										v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+144))
										*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v58
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
										F_LWLockRelease(m, v61+int32(1152))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
											if v67 != int64(0) {
												v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
												v76 = v71
											} else {
												v73 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v73)
												v76 = v73
											}
											v110 = (v76 ^ int32(1)) & base.B2i32(base.Ui64(v67) < base.Ui64(l0))
											return v110
										}
									}
								}
							} else {
								v39 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v39)
								return int32(0)
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])) = int32(1)
				v84 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
				if base.Ui64(l0) <= base.Ui64(v84) {
					v110 = int32(0)
					return v110
				} else {
					v86 = int32(_a_F_XLogNeedsFlush_0)
					v87 = int32(_a_F_XLogNeedsFlush_1)
					v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
					v89 = int64(0)
					v92 = base.AtomicRmwCmpxchg64(m, v88, int32(272), v89, v89)
					*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8])) = v92
					v94 = int32(0)
					v97 = base.AtomicRmwOr32(m, v94, int32(_a_F_XLogNeedsFlush_2), v94)
					v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
					v104 = base.AtomicRmwCmpxchg64(m, v100, int32(264), v89, v89)
					*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[9])) = v104
					v107 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
					return base.B2i32(base.Ui64(v107) < base.Ui64(l0))
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])) = int32(1)
			v84 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
			if base.Ui64(l0) <= base.Ui64(v84) {
				v110 = int32(0)
				return v110
			} else {
				v86 = int32(_a_F_XLogNeedsFlush_0)
				v87 = int32(_a_F_XLogNeedsFlush_1)
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v89 = int64(0)
				v92 = base.AtomicRmwCmpxchg64(m, v88, int32(272), v89, v89)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8])) = v92
				v94 = int32(0)
				v97 = base.AtomicRmwOr32(m, v94, int32(_a_F_XLogNeedsFlush_2), v94)
				v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v104 = base.AtomicRmwCmpxchg64(m, v100, int32(264), v89, v89)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[9])) = v104
				v107 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
				return base.B2i32(base.Ui64(v107) < base.Ui64(l0))
			}
		}
	} else {
		if v5 != 0 {
			v84 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
			if base.Ui64(l0) <= base.Ui64(v84) {
				v110 = int32(0)
				return v110
			} else {
				v86 = int32(_a_F_XLogNeedsFlush_0)
				v87 = int32(_a_F_XLogNeedsFlush_1)
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v89 = int64(0)
				v92 = base.AtomicRmwCmpxchg64(m, v88, int32(272), v89, v89)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8])) = v92
				v94 = int32(0)
				v97 = base.AtomicRmwOr32(m, v94, int32(_a_F_XLogNeedsFlush_2), v94)
				v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v104 = base.AtomicRmwCmpxchg64(m, v100, int32(264), v89, v89)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[9])) = v104
				v107 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8]))
				return base.B2i32(base.Ui64(v107) < base.Ui64(l0))
			}
		} else {
			v24 = int32(0)
			v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
			if v26 != 0 {
				v110 = v24
				return v110
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
				if base.Ui64(l0) <= base.Ui64(v28) {
					v110 = v24
					return v110
				} else {
					if v28 != int64(0) {
						v43 = int32(1)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
						v49 = F_LWLockConditionalAcquire(m, v45+int32(1152), v43)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								v110 = v43
								return v110
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+144))
								*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v58
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
								F_LWLockRelease(m, v61+int32(1152))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v67 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
									if v67 != int64(0) {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
										v76 = v71
									} else {
										v73 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v73)
										v76 = v73
									}
									v110 = (v76 ^ int32(1)) & base.B2i32(base.Ui64(v67) < base.Ui64(l0))
									return v110
								}
							}
						}
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7])))
						if v33&int32(1) == int32(0) {
							v43 = int32(1)
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
							v49 = F_LWLockConditionalAcquire(m, v45+int32(1152), v43)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v49 == int32(0) {
									v110 = v43
									return v110
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
									v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+144))
									*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v58
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5]))
									F_LWLockRelease(m, v61+int32(1152))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v67 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
										if v67 != int64(0) {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])))
											v76 = v71
										} else {
											v73 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v73)
											v76 = v73
										}
										v110 = (v76 ^ int32(1)) & base.B2i32(base.Ui64(v67) < base.Ui64(l0))
										return v110
									}
								}
							}
						} else {
							v39 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = uint8(v39)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_XLogPrefetchShmemInit(m *base.Module, l0 int32) {
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
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchShmemInit[0]))
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	F_gettimeofday(m, v9)
	mBase = m.M
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
	m.G0 = v9 + v8
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = v13 + v12*int64(1000000) - int64(946684800000000)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchShmemInit[0]))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v25
	return
}
func F_XLogReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l3|base.B2i32(l4 == v6) == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(112)
	return v264
L2:
	;
	if v223 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v27
	v30 = v18 + int32(88)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	F_ResourceOwnerEnlarge(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v165
	v170 = F_smgropen(m, v18+int32(72), int32(-1))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L37
	}
L6:
	;
	if v153 != 0 {
		v223 = l4
		goto L2
	} else {
		goto L36
	}
L7:
	;
	return int32(0)
L8:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if l4 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v153 = int32(0)
	goto L6
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[1]))
	v50 = v45 + (l4^int32(-1))*int32(56)
	v51 = int64(0)
	v54 = base.AtomicRmwCmpxchg64(m, v50, int32(24), v51, v51)
	if v54&int64(16777216) == v51 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[5]))
	v81 = int32(56)
	v83 = v80 + l4*v81
	v85 = v83 - v81
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v41 != v86 {
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v41 != v59 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v40 != v61 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v39 != v63 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if l2 != v65 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if l1 != v67 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	F_PinLocalBuffer(m, v50, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(_a_F_XLogReadBufferExtended_0)
	v74 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4])) = v74 + int64(1)
	v153 = int32(1)
	goto L6
L21:
	;
	v89 = v83 - int32(52)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v40 != v90 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v93 = v83 - int32(48)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v39 != v94 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v97 = v83 - int32(40)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if l2 != v98 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v101 = v83 - int32(44)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if l1 != v102 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v106 = F_PinBuffer(m, v85, int32(0), int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v106 == int32(0) {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v41 != v110 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v83-int32(36))))
	F_ResourceOwnerForget(m, v128, base.I64_extend_i32_s(v131+int32(1)), int32(_a_F_XLogReadBufferExtended_1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L34
	}
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v40 != v112 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v39 != v114 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if l2 != v116 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if l1 != v118 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v120 = int32(_a_F_XLogReadBufferExtended_2)
	v122 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[6])) = v122 + int64(1)
	v153 = int32(1)
	goto L6
L34:
	;
	F_UnpinBufferNoOwner(m, v85)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L10
L36:
	;
	goto L5
L37:
	;
	F_smgrcreate(m, v170, l1, int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v175 = F_smgrnblocks(m, v170, l1)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(v175) <= base.Ui32(l2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v215
	v221 = F_ReadBufferWithoutRelcache(m, v18+int32(24), l1, l2, l3, int32(0), int32(1))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L49
	}
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v180
	v182 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v182
	v184 = int32(0)
	F_log_invalid_page(m, v18+int32(40), l1, l2, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if l3 == int32(4) {
		v264 = int32(0)
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v264 = v184
	goto L1
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v170
	v194 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v194
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+109)) = uint16(v194)
	v198 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+108)) = uint8(v198)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v194)
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v204
	v211 = F_ExtendBufferedRelTo(m, v18+int32(56), l1, int32(3), l2+int32(1), l3)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v264 = v211
	goto L1
L49:
	;
	if l3 != 0 {
		v264 = v221
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v223 = v221
	goto L2
L51:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250)+14)))
	if v251 != 0 {
		v264 = v223
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[2]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v223^int32(-1))<<(uint(int32(2))%32))))
	v250 = v242
	goto L51
L53:
	;
	goto L54
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[3]))
	v250 = v244 + v223<<(uint(int32(13))%32) + int32(-8192)
	goto L51
L55:
	;
	F_ReleaseBuffer(m, v223)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v256
	F_log_invalid_page(m, v18+int32(8), l1, l2, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v264 = int32(0)
	goto L1
}
func F_XLogReadRecord(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v40 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v4 == v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
	if v15 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v21
	goto L1
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, v4)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v21 = v18
	goto L13
L11:
	;
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v29
	goto L1
L13:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v22 != int32(1) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v25 != 0 {
		v21 = v25
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	goto L1
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v45 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v41 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v43 = F_XLogReadAhead(m, l0, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v79 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v45 == v52 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v56 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v62
	goto L23
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v59 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v45)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L39
	}
L32:
	;
	v62 = v59
	goto L35
L33:
	;
	goto L34
L34:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v70
	goto L23
L35:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)))
	if v63 != int32(1) {
		goto L28
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v66 != 0 {
		v62 = v66
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L23
L40:
	;
	if v103 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v82
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v85 != int32(1) {
		v103 = v82
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v79
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v103 = v101
	goto L40
L44:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v88
	goto L47
L46:
	;
	goto L47
L47:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v91)
	v103 = v91
	goto L40
L48:
	;
	v107 = v103 + int32(32)
	goto L50
L49:
	;
	v107 = int32(0)
	goto L50
L50:
	;
	return v107
}
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v3 = l2
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[0]))
	if v5 <= l0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[1]))
		if v8 <= l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_XLogRegisterBuffer_0), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XLogRegisterBuffer_1), int32(275), int32(_a_F_XLogRegisterBuffer_2))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[0])) = l0 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[2]))
			v18 = v15 + l0*int32(_a_F_XLogRegisterBuffer_3)
			v20 = v18 + int32(4)
			if l1 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[3]))
				v42 = v29 + (l1^int32(-1))*int32(56)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
				v37 = int32(56)
				v42 = v36 + l1*v37 - v37
			}
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v44
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = v43
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v18+int32(20)))) = v49
			if l1 < int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[5]))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
				v68 = v60
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[6]))
				v68 = v62 + l1<<(uint(int32(13))%32) + int32(-8192)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v68
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(0)
			v73 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v73)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(32)
			return
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[2]))
		v18 = v15 + l0*int32(_a_F_XLogRegisterBuffer_3)
		v20 = v18 + int32(4)
		if l1 < int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[3]))
			v42 = v29 + (l1^int32(-1))*int32(56)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
			v37 = int32(56)
			v42 = v36 + l1*v37 - v37
		}
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v44
		*(*int64)(unsafe.Add(mBase, uint32(v20))) = v43
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v47
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v18+int32(20)))) = v49
		if l1 < int32(0) {
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[5]))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
			v68 = v60
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[6]))
			v68 = v62 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v68
		*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(0)
		v73 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v73)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(32)
		return
	}
}
func F_XLogShutdownWalRcv(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[0]))
	v5 = int32(1456)
	v6 = v4 + v5
	v9 = base.AtomicRmwXchg32(m, v4, v5, int32(1))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v6, int32(_a_F_XLogShutdownWalRcv_0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	switch v13 - int32(1) {
	case 0:
		goto L9
	case 1, 2, 3, 4:
		goto L8
	case 5:
		goto L7
	default:
		goto L10
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v40 = v4 + int32(12)
	F_ConditionVariablePrepareToSleep(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v31 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v31))
	if v30 == v31 {
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(6)
	goto L7
L9:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v19
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v19))
	F_ConditionVariableBroadcast(m, v4+int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v16 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v16))
	goto L6
L11:
	;
	goto L6
L12:
	;
	v37 = F_pgmem_kill(m, v30, int32(15))
	mBase = m.M
	goto L6
L13:
	;
	v43 = F_WalRcvRunning(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L23
	}
L18:
	;
	F_ConditionVariableSleep(m, v40, int32(134217785))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v50 = F_WalRcvRunning(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if v50 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	v61 = F_LWLockAcquire(m, v57+int32(1152), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[2]))
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+312)) = uint8(v65)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	F_LWLockRelease(m, v68+int32(1152))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	return
}
func F_XLogWalRcvFlush(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
	if base.Ui64(v15) <= base.Ui64(v13) {
		m.G0 = v10 + int32(80)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[2]))
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[3]))
		v22 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[4]))
		F_issue_xlog_fsync(m, v20, v22, l1)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
			*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0])) = v27
			v29 = int32(1456)
			v30 = v18 + v29
			v33 = base.AtomicRmwXchg32(m, v18, v29, int32(1))
			if v33 != 0 {
				F_s_lock(m, v30, int32(_a_F_XLogWalRcvFlush_0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
					v39 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
					if base.Ui64(v37) < base.Ui64(v39) {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v39
						*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v37
					} else {
					}
					v44 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v30))), uint32(v44))
					F_WaitLSNWakeup(m, int32(2), v39)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5]))
						F_SetLatch(m, v51+int32(4))
						mBase = m.M
						v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
						if v56 != int32(1) {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
							if v68 == int32(1) {
								v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
								v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
								v78 = v10 + int32(16)
								v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v83 = F_strlen(m, v78)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v86 = int32(0)
										F_XLogWalRcvSendReply(m, v86, v86, v86)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v10 + int32(80)
												return
											}
										}
									}
								}
							} else {
								if l0 != 0 {
									m.G0 = v10 + int32(80)
									return
								} else {
									v86 = int32(0)
									F_XLogWalRcvSendReply(m, v86, v86, v86)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v10 + int32(80)
											return
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[8]))
							if v60 <= int32(0) {
								v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
								if v68 == int32(1) {
									v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
									v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
									v78 = v10 + int32(16)
									v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v83 = F_strlen(m, v78)
										mBase = m.M
										if l0 != 0 {
											m.G0 = v10 + int32(80)
											return
										} else {
											v86 = int32(0)
											F_XLogWalRcvSendReply(m, v86, v86, v86)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v10 + int32(80)
													return
												}
											}
										}
									}
								} else {
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v86 = int32(0)
										F_XLogWalRcvSendReply(m, v86, v86, v86)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v10 + int32(80)
												return
											}
										}
									}
								}
							} else {
								F_WalSndWakeup(m, int32(1), int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
									if v68 == int32(1) {
										v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
										*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
										v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
										v78 = v10 + int32(16)
										v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v83 = F_strlen(m, v78)
											mBase = m.M
											if l0 != 0 {
												m.G0 = v10 + int32(80)
												return
											} else {
												v86 = int32(0)
												F_XLogWalRcvSendReply(m, v86, v86, v86)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													F_XLogWalRcvSendHSFeedback(m, int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														m.G0 = v10 + int32(80)
														return
													}
												}
											}
										}
									} else {
										if l0 != 0 {
											m.G0 = v10 + int32(80)
											return
										} else {
											v86 = int32(0)
											F_XLogWalRcvSendReply(m, v86, v86, v86)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v10 + int32(80)
													return
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
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
				v39 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
				if base.Ui64(v37) < base.Ui64(v39) {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v37
				} else {
				}
				v44 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v30))), uint32(v44))
				F_WaitLSNWakeup(m, int32(2), v39)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5]))
					F_SetLatch(m, v51+int32(4))
					mBase = m.M
					v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
					if v56 != int32(1) {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
						if v68 == int32(1) {
							v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
							v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
							v78 = v10 + int32(16)
							v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v83 = F_strlen(m, v78)
								mBase = m.M
								if l0 != 0 {
									m.G0 = v10 + int32(80)
									return
								} else {
									v86 = int32(0)
									F_XLogWalRcvSendReply(m, v86, v86, v86)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v10 + int32(80)
											return
										}
									}
								}
							}
						} else {
							if l0 != 0 {
								m.G0 = v10 + int32(80)
								return
							} else {
								v86 = int32(0)
								F_XLogWalRcvSendReply(m, v86, v86, v86)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v10 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[8]))
						if v60 <= int32(0) {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
							if v68 == int32(1) {
								v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
								v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
								v78 = v10 + int32(16)
								v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v83 = F_strlen(m, v78)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v86 = int32(0)
										F_XLogWalRcvSendReply(m, v86, v86, v86)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v10 + int32(80)
												return
											}
										}
									}
								}
							} else {
								if l0 != 0 {
									m.G0 = v10 + int32(80)
									return
								} else {
									v86 = int32(0)
									F_XLogWalRcvSendReply(m, v86, v86, v86)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v10 + int32(80)
											return
										}
									}
								}
							}
						} else {
							F_WalSndWakeup(m, int32(1), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
								if v68 == int32(1) {
									v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v72)
									v75 = int64(base.Ui64(v72) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v75)
									v78 = v10 + int32(16)
									v81 = F_pg_snprintf(m, v78, int32(50), int32(_a_F_XLogWalRcvFlush_1), v10)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v83 = F_strlen(m, v78)
										mBase = m.M
										if l0 != 0 {
											m.G0 = v10 + int32(80)
											return
										} else {
											v86 = int32(0)
											F_XLogWalRcvSendReply(m, v86, v86, v86)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v10 + int32(80)
													return
												}
											}
										}
									}
								} else {
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v86 = int32(0)
										F_XLogWalRcvSendReply(m, v86, v86, v86)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v10 + int32(80)
												return
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
}
func F_XLogWalRcvSendHSFeedback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[1]))
	if v13&base.B2i32(v2 < v15) == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[2])))
	if v22&int32(1) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v28 = m.G0
	v29 = int32(16)
	v30 = v28 - v29
	m.G0 = v30
	F_gettimeofday(m, v30)
	mBase = m.M
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v34 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30)+8)))
	m.G0 = v30 + v29
	v42 = v34 + v33*int64(1000000) - int64(946684800000000)
	goto L6
L5:
	;
	goto L4
L6:
	;
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[3]))
	if v42 < v46 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v51 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[1])))
	if v51 <= int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v57 = int64(9223372036854775807)
	goto L13
L12:
	;
	v57 = v42 + v51*int64(1000000)
	goto L13
L13:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v57
	goto L16
L15:
	;
	v61 = int64(9223372036854775807)
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[3])) = v61
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[4])))
	if v65 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	v72 = base.AtomicRmwXchg32(m, v69, int32(96), int32(1))
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v86 = int32(1)
	goto L19
L19:
	;
	if v86&int32(1) == int32(0) {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	F_s_lock(m, v69+int32(96), int32(_a_F_XLogWalRcvSendHSFeedback_0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[4])) = uint8(v81)
	v83 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v80)+96)), uint32(v83))
	v86 = v81
	goto L19
L23:
	;
	return
L24:
	;
	goto L22
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	if v93 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v121 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v100 = m.G0
	v102 = v100 - int32(48)
	m.G0 = v102
	F_ComputeXidHorizons(m, v102+int32(8))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v115
	goto L26
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(28)))) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v110
	m.G0 = v102 + int32(48)
	goto L26
L31:
	;
	v125 = base.I32_wrap_i64(int64(base.Ui64(v121) >> (uint(int64(32)) % 64)))
	v127 = v125 - int32(1)
	v128 = base.I32_wrap_i64(v121)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v128) < base.Ui32(v129) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = v127
	goto L34
L33:
	;
	v131 = v125
	goto L34
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if base.Ui32(v128) < base.Ui32(v132) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v134 = v127
	goto L37
L36:
	;
	v134 = v125
	goto L37
L37:
	;
	v137 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	if v137 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v134
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v131
	F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendHSFeedback_1), v10)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L23
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v153 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[7])) = v155
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v155
	goto L44
L42:
	;
	F_errfinish(m, int32(_a_F_XLogWalRcvSendHSFeedback_2), int32(1361), int32(_a_F_XLogWalRcvSendHSFeedback_3))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_4), int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v167 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v170 = int32(104)
	*(*uint8)(unsafe.Add(mBase, uint32(v166+v168))) = uint8(v170)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v174 + int32(1)
	v181 = m.G0
	v182 = int32(16)
	v183 = v181 - v182
	m.G0 = v183
	F_gettimeofday(m, v183)
	mBase = m.M
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	v187 = int64(*(*int32)(unsafe.Add(mBase, uint32(v183)+8)))
	m.G0 = v183 + v182
	v195 = v187 + v186*int64(1000000) - int64(946684800000000)
	goto L46
L46:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_4), int32(8))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	v200 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v202 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v205 = int64(56)
	v207 = int64(65280)
	v209 = int64(40)
	v212 = int64(16711680)
	v214 = int64(24)
	v216 = int64(4278190080)
	v218 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v201+v203))) = v195<<(uint(v205)%64) | v195&v207<<(uint(v209)%64) | (v195&v212<<(uint(v214)%64) | v195&v216<<(uint(v218)%64)) | (int64(base.Ui64(v195)>>(uint(v218)%64))&v216 | int64(base.Ui64(v195)>>(uint(v214)%64))&v212 | (int64(base.Ui64(v195)>>(uint(v209)%64))&v207 | int64(base.Ui64(v195)>>(uint(v205)%64))))
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v243 + int32(8)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_enlargeStringInfo(m, v200, int32(4))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	v252 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v254 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v259 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v253+v255))) = base.I32_rotr(v247, int32(24))&v259 | base.I32_rotr(v247&v259, int32(8))
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v270 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v269 + v270
	F_enlargeStringInfo(m, v252, v270)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	v277 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v279 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v284 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v278+v280))) = base.I32_rotr(v134, int32(24))&v284 | base.I32_rotr(v134&v284, int32(8))
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v295 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v294 + v295
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_enlargeStringInfo(m, v277, v295)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	v303 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v305 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v310 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v304+v306))) = base.I32_rotr(v298, int32(24))&v310 | base.I32_rotr(v298&v310, int32(8))
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v321 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v320 + v321
	F_enlargeStringInfo(m, v303, v321)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v328 = int32(_a_F_XLogWalRcvSendHSFeedback_4)
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v330 = int32(_a_F_XLogWalRcvSendHSFeedback_5)
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v335 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v329+v331))) = base.I32_rotr(v131, int32(24))&v335 | base.I32_rotr(v131&v335, int32(8))
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v347 = v345 + int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v347
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[9]))
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[10]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+44))
	m.T0[v355].(func(*base.Module, int32, int32, int32))(m, v350, v352, v347)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[2])) = uint8(base.B2i32(v359|v360 == int32(0)))
	goto L1
}
