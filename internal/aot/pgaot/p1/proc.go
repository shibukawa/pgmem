package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CleanupProcSignalState(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_CleanupProcSignalState_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[0])) = int32(0)
	v15 = int32(80)
	v16 = v11 + v15
	v19 = base.AtomicRmwXchg32(m, v11, v15, int32(1))
	if v19 != 0 {
		F_s_lock(m, v16, int32(_a_F_CleanupProcSignalState_1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[1]))
			if v23 != v25 {
				v27 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v27))
				v32 = F_errstart(m, int32(15), v27)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[2]))
						v46 = base.I32_div_s(v11-v41-int32(8), int32(112))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v46
						F_errmsg_internal(m, int32(_a_F_CleanupProcSignalState_2), v8)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CleanupProcSignalState_3), int32(264), int32(_a_F_CleanupProcSignalState_4))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v56 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v56
				v62 = base.AtomicRmwXchg64(m, v11, int32(88), int64(-1))
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11)+80)), uint32(v56))
				F_ConditionVariableBroadcast(m, v11+int32(100))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[1]))
		if v23 != v25 {
			v27 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v27))
			v32 = F_errstart(m, int32(15), v27)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupProcSignalState[2]))
					v46 = base.I32_div_s(v11-v41-int32(8), int32(112))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v46
					F_errmsg_internal(m, int32(_a_F_CleanupProcSignalState_2), v8)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CleanupProcSignalState_3), int32(264), int32(_a_F_CleanupProcSignalState_4))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v56 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v56
			v62 = base.AtomicRmwXchg64(m, v11, int32(88), int64(-1))
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11)+80)), uint32(v56))
			F_ConditionVariableBroadcast(m, v11+int32(100))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_ProcArrayEndTransactionInternal(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v76 int64
	_ = v76
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32)))) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+336)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v21&int32(14) != 0 {
		v25 = v21 & int32(241)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v25)
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint8)(unsafe.Add(mBase, uint32(v29+v30))) = uint8(v25)
	} else {
	}
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v34 == int32(0) {
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
		if v37 != int32(1) {
		} else {
			v41 = v9 << (uint(int32(1)) % 32)
			v42 = int32(_a_F_ProcArrayEndTransactionInternal_0)
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			v46 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v41+v44))) = uint8(v46)
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
			*(*uint8)(unsafe.Add(mBase, uint32(v50+v41)+1)) = uint8(v46)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v46)
		}
	} else {
		v41 = v9 << (uint(int32(1)) % 32)
		v42 = int32(_a_F_ProcArrayEndTransactionInternal_0)
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
		v46 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v41+v44))) = uint8(v46)
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[0]))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
		*(*uint8)(unsafe.Add(mBase, uint32(v50+v41)+1)) = uint8(v46)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v46)
	}
	v57 = int32(3)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransactionInternal[1]))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)+48))
	v62 = base.I32_wrap_i64(v61)
	if base.B2i32(base.Ui32(l1) < base.Ui32(v57))|base.B2i32(base.Ui32(v62) < base.Ui32(v57)) == int32(0) {
		if v62-l1 < int32(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v61 + base.I64_extend_i32_s(l1-v62)
		} else {
		}
	} else {
		if base.Ui32(l1) <= base.Ui32(v62) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v61 + base.I64_extend_i32_s(l1-v62)
		}
	}
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v60)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+56)) = v76 + int64(1)
	return
}
func F_ProcArrayShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(_a_F_ProcArrayShmemInit_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[1]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[2]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[0]))
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v13)+20)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v14
	v20 = v9 + v11
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v20 * int32(65)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = int64(1)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[4]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemInit[5])) = v32
	return
}
func F_ProcGlobalShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(100)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10)+20)), uint32(v2))
	v16 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v16
	v21 = v10 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v21
	v25 = v10 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v25
	v29 = v10 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v29
	v33 = v10 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v16
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[1]))
	v42 = int32(3)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[2]))
	if v41&v42|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v45))|v45&v42 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[3]))
	v81 = int32(_a_F_ProcGlobalShmemInit_0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v41
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[4]))
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	v90 = v41 + v80*int32(768)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v90
	v94 = v90 + v80<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v85 + int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v94 + v80<<(uint(int32(1))%32)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[5]))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[6]))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[7]))
	if v108 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v58 = v45 + v41
	v60 = v41 + int32(4)
	if base.Ui32(v60) < base.Ui32(v58) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v62 = v58
	goto L8
L7:
	;
	v62 = v60
	goto L8
L8:
	;
	v67 = (v41^int32(-1)+v62)&int32(-4) + int32(4)
	if v67 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	base.MemoryFill(m, v41, int32(0), v67)
	goto L1
L10:
	;
	base.MemoryFill(m, v41, int32(0), v45)
	goto L1
L11:
	;
	base.MemoryFill(m, v106, int32(0), v108)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[4]))
	v115 = m.G0
	v117 = v115 - int32(112)
	m.G0 = v117
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[8]))
	v125 = F___fstatat(m, int32(-100), v120, v117+int32(16), int32(0))
	mBase = m.M
	goto L14
L14:
	;
	if v125 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[9])) = v112 + int32(38)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[10])) = int32(0)
	F_on_shmem_exit(m, int32(961), int64(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L18
	} else {
		goto L23
	}
L18:
	;
	return
L19:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135
	F_errmsg(m, int32(_a_F_ProcGlobalShmemInit_1), v117)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_ProcGlobalShmemInit_2), int32(212), int32(_a_F_ProcGlobalShmemInit_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
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
	m.G0 = v117 + int32(112)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[3]))
	if v158 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v165 = v106
	v167 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[4]))
	v439 = v41 + v436*int32(768)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[11])) = v439
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[12])) = v439 + int32(_a_F_ProcGlobalShmemInit_4)
	return
L27:
	;
	v173 = v41 + v167*int32(768)
	v174 = v165 + v104<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+568)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v173)+564)) = v165
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[4]))
	if v167 < v178+int32(38) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[10]))
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[9]))
	if v183 < v185 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	goto L31
L31:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[13]))
	if v167 < v254 {
		goto L52
	} else {
		goto L53
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+332)) = v192
	v239 = v173 + int32(316)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = int64(0)
	v244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+8)) = uint8(v244)
	goto L48
L33:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L45
	}
L34:
	;
	v187 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[14]))
	v192 = v189 + v183<<(uint(int32(7))%32)
	v194 = m.Env.Pgmem_sem(m, v187, v192, int32(1))
	mBase = m.M
	if v187 <= v194 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L42
	}
L37:
	;
	if v202 < int32(0) {
		goto L33
	} else {
		goto L41
	}
L38:
	;
	v202 = v194
	goto L37
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[15])) = int32(0) - v194
	v202 = int32(-1)
	goto L37
L41:
	;
	v205 = int32(_a_F_ProcGlobalShmemInit_5)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[10])) = v207 + int32(1)
	goto L32
L42:
	;
	F_errmsg_internal(m, int32(_a_F_ProcGlobalShmemInit_6), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_ProcGlobalShmemInit_2), int32(264), int32(_a_F_ProcGlobalShmemInit_7))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errmsg_internal(m, int32(_a_F_ProcGlobalShmemInit_8), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ProcGlobalShmemInit_2), int32(138), int32(_a_F_ProcGlobalShmemInit_9))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L18
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
	F_LWLockInitialize(m, v173+int32(548), int32(69))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	goto L31
L50:
	;
	v348 = v173 + int32(540)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+544)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v173)+540)) = v348
	v352 = v173 + int32(532)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+536)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v173)+532)) = v352
	v356 = v173 + int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+528)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v173)+524)) = v356
	v360 = v173 + int32(516)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+520)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v173)+516)) = v360
	v364 = v173 + int32(508)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+512)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v173)+508)) = v364
	v368 = v173 + int32(500)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+504)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v173)+500)) = v368
	v372 = v173 + int32(492)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+496)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v173)+492)) = v372
	v376 = v173 + int32(484)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+488)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v173)+484)) = v376
	v380 = v173 + int32(476)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+480)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v173)+476)) = v380
	v384 = v173 + int32(468)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+472)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v173)+468)) = v384
	v388 = v173 + int32(460)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+464)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v173)+460)) = v388
	v392 = v173 + int32(452)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+456)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v173)+452)) = v392
	v396 = v173 + int32(444)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+448)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v173)+444)) = v396
	v400 = v173 + int32(436)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+440)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v173)+436)) = v400
	v404 = v173 + int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+432)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v173)+428)) = v404
	v408 = v173 + int32(420)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+424)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v173)+420)) = v408
	v411 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+608)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v173)+620)) = v411
	v416 = v173 + int32(368)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+372)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v173)+368)) = v416
	*(*int64)(unsafe.Add(mBase, uint32(v173)+408)) = int64(0)
	v422 = v167 + int32(1)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[3]))
	if base.Ui32(v422) < base.Ui32(v424) {
		v165 = v104<<(uint(int32(6))%32) + v174
		v167 = v422
		goto L27
	} else {
		goto L74
	}
L51:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v339 + v341
	goto L50
L52:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	v259 = v257 + int32(24)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+28))
	if v260 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[16]))
	v277 = v254 + v274 + int32(2)
	if v167 < v277 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+28)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v257)+24)) = v259
	goto L57
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v259
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v266
	v269 = v173 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v269
	v339 = int32(24)
	goto L51
L58:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	v282 = v280 + int32(32)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)+36))
	if v283 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[17]))
	if v167 < v297+v277 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+36)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v280)+32)) = v282
	goto L63
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v282
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v289
	v292 = v173 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v292
	v339 = int32(32)
	goto L51
L64:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	v303 = v301 + int32(40)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+44))
	if v304 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[4]))
	if v318 <= v167 {
		goto L50
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+44)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v301)+40)) = v303
	goto L69
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v303
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v310
	v313 = v173 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v313
	v339 = int32(40)
	goto L51
L70:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemInit[0]))
	v323 = v321 + int32(48)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v321)+52))
	if v324 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+52)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v321)+48)) = v323
	goto L73
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v323
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v330
	v333 = v173 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v333
	v339 = int32(48)
	goto L51
L74:
	;
	goto L28
}
func F_ProcSignalShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalShmemRequest[0]))
	v12 = F_mul_size(m, v8+int32(38), int32(112))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = F_add_size(m, v12, int32(8))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_ProcSignalShmemRequest_0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ProcSignalShmemRequest_1)
			F_ShmemRequestStructWithOpts(m, v5)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_RemoveProcFromArray(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveProcFromArray[0]))
	F_ProcArrayRemove(m, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_proc_exit_prepare(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[1])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[2])) = v2
	v18 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[3])) = v18
	*(*uint8)(unsafe.Add(mBase, _c_F_proc_exit_prepare[4])) = uint8(v18)
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[5])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[6])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[7])) = v2
	F_shmem_exit(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v40
	F_errmsg_internal(m, int32(_a_F_proc_exit_prepare_0), v6)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v50 = int32(_a_F_proc_exit_prepare_1)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8]))
	v54 = v52 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8])) = v54
	if int32(0) <= v54 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errfinish(m, int32(_a_F_proc_exit_prepare_2), int32(202), int32(_a_F_proc_exit_prepare_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v59 = v54
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8])) = int32(0)
	m.G0 = v6 + int32(16)
	return
L12:
	;
	v62 = v59 << (uint(int32(4)) % 32)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_proc_exit_prepare[9])))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_proc_exit_prepare[10])))
	m.T0[v64].(func(*base.Module, int32, int64))(m, l0, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v67 = int32(_a_F_proc_exit_prepare_1)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8]))
	v71 = v69 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_proc_exit_prepare[8])) = v71
	if int32(0) <= v71 {
		v59 = v71
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
