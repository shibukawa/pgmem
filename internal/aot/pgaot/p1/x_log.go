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
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
	v9 = base.AtomicRmwXchg32(m, v6, int32(96), int32(1))
	if v9 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
		F_s_lock(m, v11+int32(96), int32(_a_F_GetXLogReplayRecPtr_0), int32(_a_F_GetXLogReplayRecPtr_1), int32(_a_F_GetXLogReplayRecPtr_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
			v25 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23
			} else {
			}
			return v24
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
		v25 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23
		} else {
		}
		return v24
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
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v118 int64
	_ = v118
	var v125 int64
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int64
	_ = v302
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v336 int64
	_ = v336
	var v341 int64
	_ = v341
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
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
	v25 = base.AtomicRmwOr64(m, v22, int32(264), int64(0))
	if base.Ui64(v1) <= base.Ui64(v25) {
		v341 = v25
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L76
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v341
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
	F_s_lock(m, v22, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1528), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v38 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22))), uint32(v38))
	v42 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[2])))
	v43 = base.I64_div_u_s(v37, v42)
	v45 = v37 - v43*v42
	if base.Ui64(v45) <= base.Ui64(int64(8151)) {
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
	v75 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[3])))
	v79 = v43*v75 + v73&int64(4294967295)
	if base.Ui64(v1) <= base.Ui64(v79) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v48 = int64(0)
	if v45 == v48 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v55 = v45 - int64(8152)
	v56 = int64(8168)
	v57 = base.I64_div_u_s(v55, v56)
	v59 = v57 << (uint(int64(13)) % 64)
	v64 = v55 - v57*v56
	if v64 == int64(0) {
		v73 = v59 - int64(-8192)
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v53 = v48
	goto L17
L16:
	;
	v53 = v45 + int64(40)
	goto L17
L17:
	;
	v73 = v53
	goto L11
L18:
	;
	v73 = v64 + v59 + int64(8216)
	goto L11
L19:
	;
	v107 = v79
	v114 = int32(0)
	goto L27
L20:
	;
	v103 = v1
	goto L19
L21:
	;
	goto L22
L22:
	;
	v83 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if v83 == int32(0) {
		v103 = v79
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+12)) = uint32(v79)
	v88 = int64(32)
	v89 = int64(base.Ui64(v79) >> (uint(v88) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+8)) = uint32(v89)
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)) = uint32(v1)
	v93 = int64(base.Ui64(v1) >> (uint(v88) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17))) = uint32(v93)
	F_errmsg(m, int32(_a_F_WaitXLogInsertionsToFinish_2), v17)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1545), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v103 = v79
	goto L19
L27:
	;
	v118 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v118
	v125 = v118
	goto L30
L28:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[1]))
	v316 = int64(0)
	v319 = base.AtomicRmwCmpxchg64(m, v315, int32(264), v316, v316)
	if base.Ui64(v309) <= base.Ui64(v319) {
		goto L67
	} else {
		goto L68
	}
L29:
	;
	v311 = v114 + int32(1)
	if v311 != int32(8) {
		v107 = v309
		v114 = v311
		goto L27
	} else {
		goto L66
	}
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[4]))
	v139 = v138 + v114<<(uint(int32(7))%32)
	v141 = v139 + int32(16)
	v143 = v17 + int32(24)
	v144 = int32(0)
	v145 = int32(_a_F_WaitXLogInsertionsToFinish_3)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v147 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[0]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v153&int32(_a_F_WaitXLogInsertionsToFinish_4) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v302 == int64(0) {
		v309 = v107
		goto L29
	} else {
		goto L62
	}
L32:
	;
	if int32(0) < v250 {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v161 = v144
	goto L36
L34:
	;
	v235 = v144
	goto L35
L35:
	;
	v250 = v235
	v253 = int32(1)
	goto L32
L36:
	;
	v170 = int64(0)
	v173 = base.AtomicRmwCmpxchg64(m, v141, int32(0), v170, v170)
	if v125 != v173 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v235 = v211
	goto L35
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v143))) = v173
	v250 = v161
	v253 = int32(0)
	goto L32
L39:
	;
	goto L40
L40:
	;
	F_LWLockQueueSelf(m, v139, int32(2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v182 = base.AtomicRmwOr32(m, v139, int32(4), int32(1073741824))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v185 = v183 & int32(_a_F_WaitXLogInsertionsToFinish_4)
	if v185 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139))))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v199 | int32(16777216)
	v211 = v161
	goto L48
L43:
	;
	v186 = int64(0)
	v189 = base.AtomicRmwCmpxchg64(m, v141, int32(0), v186, v186)
	if v189 == v125 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_LWLockDequeueSelf(m, v139)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L47
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v143))) = v189
	goto L45
L47:
	;
	v250 = v161
	v253 = base.B2i32(v185 == int32(0))
	goto L32
L48:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	F_PGSemaphoreLock(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L50
	}
L49:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v227&int32(_a_F_WaitXLogInsertionsToFinish_4) != 0 {
		v161 = v211
		goto L36
	} else {
		goto L52
	}
L50:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+74)))
	if v222 != 0 {
		v211 = v211 + int32(1)
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
	v266 = v250
	goto L56
L54:
	;
	goto L55
L55:
	;
	v296 = int32(_a_F_WaitXLogInsertionsToFinish_3)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v298 - int32(1)
	if v253 != 0 {
		v309 = v107
		goto L29
	} else {
		goto L60
	}
L56:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	F_PGSemaphoreUnlock(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L9
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	v278 = int32(1)
	if base.Ui32(v278) < base.Ui32(v266) {
		v266 = v266 - v278
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	if base.Ui64(v302) < base.Ui64(v103) {
		v125 = v302
		goto L30
	} else {
		goto L61
	}
L61:
	;
	goto L31
L62:
	;
	if base.Ui64(v302) < base.Ui64(v107) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v307 = v302
	goto L65
L64:
	;
	v307 = v107
	goto L65
L65:
	;
	v309 = v307
	goto L29
L66:
	;
	goto L28
L67:
	;
	v341 = v319
	goto L4
L68:
	;
	goto L69
L69:
	;
	v322 = v319
	goto L70
L70:
	;
	v336 = base.AtomicRmwCmpxchg64(m, v315, int32(264), v322, v309)
	if v322 == v336 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v341 = v336
	goto L4
L72:
	;
	v341 = v309
	goto L4
L73:
	;
	goto L74
L74:
	;
	if base.Ui64(v336) < base.Ui64(v309) {
		v322 = v336
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	F_errmsg_internal(m, int32(_a_F_WaitXLogInsertionsToFinish_5), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1517), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
											F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(537), int32(_a_F_XLogArchiveForceDone_6))
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
													F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(545), int32(_a_F_XLogArchiveForceDone_6))
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
				F_errfinish(m, int32(_a_F_XLogEnsureRecordSpace_3), int32(194), int32(_a_F_XLogEnsureRecordSpace_4))
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])))
	if v6 != int32(1) {
		v79 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
		if base.Ui64(l0) <= base.Ui64(v79) {
			v101 = v3
		} else {
			v81 = int32(_a_F_XLogNeedsFlush_0)
			v82 = int32(_a_F_XLogNeedsFlush_1)
			v83 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
			v84 = int64(0)
			v87 = base.AtomicRmwCmpxchg64(m, v83, int32(280), v84, v84)
			*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])) = v87
			v91 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
			v95 = base.AtomicRmwCmpxchg64(m, v91, int32(272), v84, v84)
			*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = v95
			v98 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
			v101 = base.B2i32(base.Ui64(v98) < base.Ui64(l0))
		}
		return v101
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+316))
		v13 = int32(2)
		*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])) = uint8(base.B2i32(v12 != v13))
		if v12 == v13 {
			v79 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
			if base.Ui64(l0) <= base.Ui64(v79) {
				v101 = v3
			} else {
				v81 = int32(_a_F_XLogNeedsFlush_0)
				v82 = int32(_a_F_XLogNeedsFlush_1)
				v83 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v84 = int64(0)
				v87 = base.AtomicRmwCmpxchg64(m, v83, int32(280), v84, v84)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])) = v87
				v91 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v95 = base.AtomicRmwCmpxchg64(m, v91, int32(272), v84, v84)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = v95
				v98 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
				v101 = base.B2i32(base.Ui64(v98) < base.Ui64(l0))
			}
			return v101
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
			if v19 != int64(0) {
				if base.Ui64(l0) <= base.Ui64(v19) {
					v101 = v3
					return v101
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
					if v35&int32(1) != 0 {
						v101 = v3
						return v101
					} else {
						v38 = int32(1)
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
						v44 = F_LWLockConditionalAcquire(m, v40+int32(1152), v38)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v44 == int32(0) {
								v101 = v38
								return v101
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7]))
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+136))
								*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v53
								v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
								F_LWLockRelease(m, v56+int32(1152))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v62 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
									if v62 != int64(0) {
										v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
										v71 = v66
									} else {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v68)
										v71 = v68
									}
									return (v71 ^ int32(1)) & base.B2i32(base.Ui64(v62) < base.Ui64(l0))
								}
							}
						}
					}
				}
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8])))
				if v23&int32(1) == int32(0) {
					if base.Ui64(l0) <= base.Ui64(v19) {
						v101 = v3
						return v101
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
						if v35&int32(1) != 0 {
							v101 = v3
							return v101
						} else {
							v38 = int32(1)
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
							v44 = F_LWLockConditionalAcquire(m, v40+int32(1152), v38)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v44 == int32(0) {
									v101 = v38
									return v101
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7]))
									v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+136))
									*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v53
									v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
									F_LWLockRelease(m, v56+int32(1152))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v62 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
										if v62 != int64(0) {
											v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
											v71 = v66
										} else {
											v68 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v68)
											v71 = v68
										}
										return (v71 ^ int32(1)) & base.B2i32(base.Ui64(v62) < base.Ui64(l0))
									}
								}
							}
						}
					}
				} else {
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v29)
					return int32(0)
				}
			}
		}
	}
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v320 int32
	_ = v320
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int64
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int64
	_ = v437
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
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
	return v449
L2:
	;
	if v403 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v27
	v30 = v18 + int32(88)
	v31 = m.G0
	v33 = v31 - int32(32)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	F_ResourceOwnerEnlarge(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v339
	v341 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v341
	v346 = F_smgropen(m, v18+int32(72), int32(-1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L84
	}
L6:
	;
	return int32(0)
L7:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if l4 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v33 + int32(32)
	if v320 != 0 {
		v403 = l4
		goto L2
	} else {
		goto L83
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[1]))
	v54 = v49 + (l4^int32(-1))<<(uint(int32(6))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v55&int32(16777216) == int32(0) {
		v320 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = l4
	v85 = v81 + l4<<(uint(int32(6))%32)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[3]))
	if l4 == v87 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v45 != v60 {
		v320 = v6
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v44 != v62 {
		v320 = v6
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if v43 != v64 {
		v320 = v6
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if l2 != v66 {
		v320 = v6
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if l1 != v68 {
		v320 = v6
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v70 = int32(1)
	F_PinLocalBuffer(m, v54, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v74 = int32(_a_F_XLogReadBufferExtended_0)
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4])) = v76 + int64(1)
	v320 = v70
	goto L9
L20:
	;
	if v227&int32(16777216) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = int32(_a_F_XLogReadBufferExtended_1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = int32(_a_F_XLogReadBufferExtended_2)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(_a_F_XLogReadBufferExtended_3)
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(0)
	v154 = v85 - int32(40)
	v155 = int32(_a_F_XLogReadBufferExtended_4)
	v157 = base.AtomicRmwOr32(m, v154, v148, v155)
	if v157&v155 != 0 {
		goto L51
	} else {
		goto L52
	}
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v134 <= int32(0) {
		goto L21
	} else {
		goto L50
	}
L23:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_5)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[5]))
	if l4 == v91 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_6)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[6]))
	if l4 == v95 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_7)
	goto L22
L30:
	;
	goto L31
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[7]))
	if l4 == v99 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_8)
	goto L22
L33:
	;
	goto L34
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[8]))
	if l4 == v103 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_9)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[9]))
	if l4 == v107 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_10)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[10]))
	if l4 == v111 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_11)
	goto L22
L42:
	;
	goto L43
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[11]))
	if l4 == v115 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_12)
	goto L22
L45:
	;
	goto L46
L46:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[12]))
	if v119 == int32(0) {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[13]))
	v126 = int32(0)
	v128 = F_hash_search(m, v123, v33+int32(8), v126, v126)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v128 == int32(0) {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v132 = v128
	goto L22
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(40))))
	v227 = v139
	v231 = int32(1)
	goto L20
L51:
	;
	goto L54
L52:
	;
	v189 = v157
	goto L53
L53:
	;
	v202 = int32(_a_F_XLogReadBufferExtended_13)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[14]))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(8))+8))
	if v205 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	F_perform_spin_delay(m, v33+int32(8))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L56
	}
L55:
	;
	v189 = v181
	goto L53
L56:
	;
	v179 = int32(_a_F_XLogReadBufferExtended_4)
	v181 = base.AtomicRmwOr32(m, v154, int32(0), v179)
	if v181&v179 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v227 = v189
	v231 = v148
	goto L20
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[14])) = v220
	goto L59
L61:
	;
	if int32(999) < v203 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v203 < int32(11) {
		goto L59
	} else {
		goto L68
	}
L64:
	;
	v210 = int32(900)
	if v210 <= v203 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v213 = v210
	goto L67
L66:
	;
	v213 = v203
	goto L67
L67:
	;
	v220 = v213 + int32(100)
	goto L60
L68:
	;
	v220 = v203 - int32(1)
	goto L60
L69:
	;
	if v231 != 0 {
		v320 = v6
		goto L9
	} else {
		goto L82
	}
L70:
	;
	v242 = v85 + int32(-64)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v45 != v243 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(60))))
	if v44 != v247 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(56))))
	if v43 != v251 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(48))))
	if l2 != v255 {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(52))))
	if l1 != v259 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	if v231 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v293 = int32(_a_F_XLogReadBufferExtended_14)
	v295 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[15])) = v295 + int64(1)
	v320 = int32(1)
	goto L9
L77:
	;
	v262 = F_PinBuffer(m, v242, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v265 = v85 - int32(40)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v267 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = (v266 + v267) & int32(-4194305)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(44))))
	v275 = int32(_a_F_XLogReadBufferExtended_15)
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[16])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = v267
	v283 = v274 + v267
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v283
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	F_ResourceOwnerRemember(m, v286, v283, int32(_a_F_XLogReadBufferExtended_16))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	goto L76
L81:
	;
	goto L76
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85-int32(40)))) = v227 & int32(-4194305)
	v320 = v6
	goto L9
L83:
	;
	goto L5
L84:
	;
	F_smgrcreate(m, v346, l1, int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	v351 = F_smgrnblocks(m, v346, l1)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(v351) <= base.Ui32(l2) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if l3 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v389
	v391 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v391
	v397 = F_ReadBufferWithoutRelcache(m, v18+int32(24), l1, l2, l3, int32(0), int32(1))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L6
	} else {
		goto L96
	}
L90:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v358
	v360 = int32(0)
	F_log_invalid_page(m, v18+int32(40), l1, l2, v360)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if l3 == int32(4) {
		v449 = int32(0)
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v449 = v360
	goto L1
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v346
	v370 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v370
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+109)) = uint16(v370)
	v374 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+108)) = uint8(v374)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v370)
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v378
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v380
	v387 = F_ExtendBufferedRelTo(m, v18+int32(56), l1, int32(3), l2+int32(1), l3)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v449 = v387
	goto L1
L96:
	;
	if l3 != 0 {
		v449 = v397
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v403 = v397
	goto L2
L98:
	;
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+14)))
	if v432 != 0 {
		v449 = v403
		goto L1
	} else {
		goto L102
	}
L99:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[17]))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417+(v403^int32(-1))<<(uint(int32(2))%32))))
	v431 = v423
	goto L98
L100:
	;
	goto L101
L101:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[18]))
	v431 = v425 + v403<<(uint(int32(13))%32) + int32(-8192)
	goto L98
L102:
	;
	F_ReleaseBuffer(m, v403)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v435
	v437 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v437
	F_log_invalid_page(m, v18+int32(8), l1, l2, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v449 = int32(0)
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
					F_errfinish(m, int32(_a_F_XLogRegisterBuffer_1), int32(267), int32(_a_F_XLogRegisterBuffer_2))
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
				v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
				v42 = v36 + l1<<(uint(int32(6))%32) + int32(-64)
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
			v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
			v42 = v36 + l1<<(uint(int32(6))%32) + int32(-64)
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
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
	F_s_lock(m, v6, int32(_a_F_XLogShutdownWalRcv_0), int32(190), int32(_a_F_XLogShutdownWalRcv_1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	switch v15 - int32(1) {
	case 0:
		goto L9
	case 1, 2, 3:
		goto L8
	case 4:
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
	v42 = v4 + int32(12)
	F_ConditionVariablePrepareToSleep(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v33 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v33))
	if v32 == v33 {
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(5)
	goto L7
L9:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v21
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v21))
	F_ConditionVariableBroadcast(m, v4+int32(12))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v18 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v18))
	goto L6
L11:
	;
	goto L6
L12:
	;
	v39 = F_pgmem_kill(m, v32, int32(15))
	mBase = m.M
	goto L6
L13:
	;
	v45 = F_WalRcvRunning(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v45 != 0 {
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
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L23
	}
L18:
	;
	F_ConditionVariableSleep(m, v42, int32(134217781))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v52 = F_WalRcvRunning(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if v52 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	v63 = F_LWLockAcquire(m, v59+int32(1152), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[2]))
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+320)) = uint8(v67)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	F_LWLockRelease(m, v70+int32(1152))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
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
				F_s_lock(m, v30, int32(_a_F_XLogWalRcvFlush_0), int32(998), int32(_a_F_XLogWalRcvFlush_1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
					v41 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
					if base.Ui64(v39) < base.Ui64(v41) {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v39
					} else {
					}
					v46 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v30))), uint32(v46))
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5]))
					F_SetLatch(m, v50+int32(4))
					mBase = m.M
					v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
					if v55 != int32(1) {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
						if v67 == int32(1) {
							v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
							v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
							v77 = v10 + int32(16)
							v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = F_strlen(m, v77)
								mBase = m.M
								if l0 != 0 {
									m.G0 = v10 + int32(80)
									return
								} else {
									v85 = int32(0)
									F_XLogWalRcvSendReply(m, v85, v85)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
								v85 = int32(0)
								F_XLogWalRcvSendReply(m, v85, v85)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										m.G0 = v10 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[8]))
						if v59 <= int32(0) {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
							if v67 == int32(1) {
								v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
								v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
								v77 = v10 + int32(16)
								v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = F_strlen(m, v77)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v85 = int32(0)
										F_XLogWalRcvSendReply(m, v85, v85)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
									v85 = int32(0)
									F_XLogWalRcvSendReply(m, v85, v85)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
								if v67 == int32(1) {
									v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
									v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
									v77 = v10 + int32(16)
									v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = F_strlen(m, v77)
										mBase = m.M
										if l0 != 0 {
											m.G0 = v10 + int32(80)
											return
										} else {
											v85 = int32(0)
											F_XLogWalRcvSendReply(m, v85, v85)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
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
										v85 = int32(0)
										F_XLogWalRcvSendReply(m, v85, v85)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
			} else {
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
				v41 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
				if base.Ui64(v39) < base.Ui64(v41) {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v39
				} else {
				}
				v46 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v30))), uint32(v46))
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5]))
				F_SetLatch(m, v50+int32(4))
				mBase = m.M
				v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
				if v55 != int32(1) {
					v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
					if v67 == int32(1) {
						v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
						*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
						v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
						*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
						v77 = v10 + int32(16)
						v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = F_strlen(m, v77)
							mBase = m.M
							if l0 != 0 {
								m.G0 = v10 + int32(80)
								return
							} else {
								v85 = int32(0)
								F_XLogWalRcvSendReply(m, v85, v85)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
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
							v85 = int32(0)
							F_XLogWalRcvSendReply(m, v85, v85)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_XLogWalRcvSendHSFeedback(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									m.G0 = v10 + int32(80)
									return
								}
							}
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[8]))
					if v59 <= int32(0) {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
						if v67 == int32(1) {
							v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
							v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
							v77 = v10 + int32(16)
							v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = F_strlen(m, v77)
								mBase = m.M
								if l0 != 0 {
									m.G0 = v10 + int32(80)
									return
								} else {
									v85 = int32(0)
									F_XLogWalRcvSendReply(m, v85, v85)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
								v85 = int32(0)
								F_XLogWalRcvSendReply(m, v85, v85)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
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
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7])))
							if v67 == int32(1) {
								v71 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v71)
								v74 = int64(base.Ui64(v71) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v74)
								v77 = v10 + int32(16)
								v80 = F_pg_snprintf(m, v77, int32(50), int32(_a_F_XLogWalRcvFlush_2), v10)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = F_strlen(m, v77)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v10 + int32(80)
										return
									} else {
										v85 = int32(0)
										F_XLogWalRcvSendReply(m, v85, v85)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
									v85 = int32(0)
									F_XLogWalRcvSendReply(m, v85, v85)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v199 int64
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
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
	v90 = int32(1)
	goto L19
L19:
	;
	if v90&int32(1) == int32(0) {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	F_s_lock(m, v74+int32(96), int32(_a_F_XLogWalRcvSendHSFeedback_0), int32(_a_F_XLogWalRcvSendHSFeedback_1), int32(_a_F_XLogWalRcvSendHSFeedback_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[4])) = uint8(v85)
	v87 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v84)+96)), uint32(v87))
	v90 = v85
	goto L19
L23:
	;
	return
L24:
	;
	goto L22
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	if v97 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v125 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v104 = m.G0
	v106 = v104 - int32(48)
	m.G0 = v106
	F_ComputeXidHorizons(m, v106+int32(8))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v119
	goto L26
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(28)))) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v114
	m.G0 = v106 + int32(48)
	goto L26
L31:
	;
	v129 = base.I32_wrap_i64(int64(base.Ui64(v125) >> (uint(int64(32)) % 64)))
	v131 = v129 - int32(1)
	v132 = base.I32_wrap_i64(v125)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v132) < base.Ui32(v133) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v135 = v131
	goto L34
L33:
	;
	v135 = v129
	goto L34
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if base.Ui32(v132) < base.Ui32(v136) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v138 = v131
	goto L37
L36:
	;
	v138 = v129
	goto L37
L37:
	;
	v141 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	if v141 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v138
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v135
	F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendHSFeedback_3), v10)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v157 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[7])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v159
	goto L44
L42:
	;
	F_errfinish(m, int32(_a_F_XLogWalRcvSendHSFeedback_4), int32(1233), int32(_a_F_XLogWalRcvSendHSFeedback_5))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_6), int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v171 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v174 = int32(104)
	*(*uint8)(unsafe.Add(mBase, uint32(v170+v172))) = uint8(v174)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v178 + int32(1)
	v185 = m.G0
	v186 = int32(16)
	v187 = v185 - v186
	m.G0 = v187
	F_gettimeofday(m, v187)
	mBase = m.M
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v187)))
	v191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+8)))
	m.G0 = v187 + v186
	v199 = v191 + v190*int64(1000000) - int64(946684800000000)
	goto L46
L46:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_6), int32(8))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	v204 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v206 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v209 = int64(56)
	v211 = int64(65280)
	v213 = int64(40)
	v216 = int64(16711680)
	v218 = int64(24)
	v220 = int64(4278190080)
	v222 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v199<<(uint(v209)%64) | v199&v211<<(uint(v213)%64) | (v199&v216<<(uint(v218)%64) | v199&v220<<(uint(v222)%64)) | (int64(base.Ui64(v199)>>(uint(v222)%64))&v220 | int64(base.Ui64(v199)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v199)>>(uint(v213)%64))&v211 | int64(base.Ui64(v199)>>(uint(v209)%64))))
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v247 + int32(8)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_enlargeStringInfo(m, v204, int32(4))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	v256 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v258 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v263 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v259))) = base.I32_rotr(v251, int32(24))&v263 | base.I32_rotr(v251&v263, int32(8))
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v274 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v273 + v274
	F_enlargeStringInfo(m, v256, v274)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	v281 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v283 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v288 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v282+v284))) = base.I32_rotr(v138, int32(24))&v288 | base.I32_rotr(v138&v288, int32(8))
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v299 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v298 + v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_enlargeStringInfo(m, v281, v299)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	v307 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v309 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v314 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v308+v310))) = base.I32_rotr(v302, int32(24))&v314 | base.I32_rotr(v302&v314, int32(8))
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v325 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v324 + v325
	F_enlargeStringInfo(m, v307, v325)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v332 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v334 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v339 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v333+v335))) = base.I32_rotr(v135, int32(24))&v339 | base.I32_rotr(v135&v339, int32(8))
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v351 = v349 + int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v351
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[9]))
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[10]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+44))
	m.T0[v359].(func(*base.Module, int32, int32, int32))(m, v354, v356, v351)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[2])) = uint8(base.B2i32(v363|v364 == int32(0)))
	goto L1
}
