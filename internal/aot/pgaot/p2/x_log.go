package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetXLogBuffer(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v59 int32
	_ = v59
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	v1 = l0
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int64(base.Ui64(v1) >> (uint(int64(13)) % 64))
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_GetXLogBuffer[0]))
	if v16 == v18 {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[1]))
		v153 = v21 + base.I32_wrap_i64(v1)&int32(_a_F_GetXLogBuffer_0)
		m.G0 = v13 + int32(16)
		return v153
	} else {
		v27 = v1 & int64(8191)
		v28 = int64(-8192)
		v31 = v1&v28 - v28
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+300))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+304))
		v39 = base.I64_rem_u_s(v16, base.I64_extend_i32_s(v35+int32(1)))
		v40 = base.I32_wrap_i64(v39)
		v44 = int64(0)
		v47 = base.AtomicRmwCmpxchg64(m, v34+v40<<(uint(int32(3))%32), int32(0), v44, v44)
		if v31 != v47 {
			v50 = v27 - int64(24)
			if base.Ui64(int64(16)) < base.Ui64(v50) {
				v78 = v1
			} else {
				switch base.I32_wrap_i64(v50) - int32(1) {
				case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
					v78 = v1
				case 15:
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[3]))
					if base.Ui64(v1&base.I64_extend_i32_s(v70-int32(1))) < base.Ui64(int64(8192)) {
						v77 = v1 - int64(40)
					} else {
						v77 = v1
					}
					v78 = v77
				default:
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[3]))
					if base.Ui64(int64(8192)) < base.Ui64(v1&base.I64_extend_i32_s(v59-int32(1))) {
						v66 = v1 - int64(24)
					} else {
						v66 = v1
					}
					v78 = v66
				}
			}
			v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[4]))
			v82 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetXLogBuffer[5])))
			if v82 != 0 {
				F_LWLockUpdateVar(m, v80+int32(896), v80+int32(912), v78)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					F_AdvanceXLInsertBuffer(m, v1, l1, int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+300))
						v110 = int64(0)
						v113 = base.AtomicRmwCmpxchg64(m, v106+v40<<(uint(int32(3))%32), int32(0), v110, v110)
						if v31 == v113 {
							*(*int64)(unsafe.Add(mBase, _c_F_GetXLogBuffer[0])) = v16
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+296))
							v142 = v139 + v40<<(uint(int32(13))%32)
							*(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[1])) = v142
							v153 = v142 + base.I32_wrap_i64(v27)
							m.G0 = v13 + int32(16)
							return v153
						} else {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v1)
								v121 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v121)
								F_errmsg_internal(m, int32(_a_F_GetXLogBuffer_1), v13)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetXLogBuffer_2), int32(1712), int32(_a_F_GetXLogBuffer_3))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
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
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[6]))
				v95 = v80 + v92<<(uint(int32(7))%32)
				F_LWLockUpdateVar(m, v95, v95+int32(16), v78)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_AdvanceXLInsertBuffer(m, v1, l1, int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+300))
						v110 = int64(0)
						v113 = base.AtomicRmwCmpxchg64(m, v106+v40<<(uint(int32(3))%32), int32(0), v110, v110)
						if v31 == v113 {
							*(*int64)(unsafe.Add(mBase, _c_F_GetXLogBuffer[0])) = v16
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+296))
							v142 = v139 + v40<<(uint(int32(13))%32)
							*(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[1])) = v142
							v153 = v142 + base.I32_wrap_i64(v27)
							m.G0 = v13 + int32(16)
							return v153
						} else {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v1)
								v121 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v121)
								F_errmsg_internal(m, int32(_a_F_GetXLogBuffer_1), v13)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetXLogBuffer_2), int32(1712), int32(_a_F_GetXLogBuffer_3))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
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
			}
		} else {
			*(*int64)(unsafe.Add(mBase, _c_F_GetXLogBuffer[0])) = v16
			v138 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[2]))
			v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+296))
			v142 = v139 + v40<<(uint(int32(13))%32)
			*(*int32)(unsafe.Add(mBase, _c_F_GetXLogBuffer[1])) = v142
			v153 = v142 + base.I32_wrap_i64(v27)
			m.G0 = v13 + int32(16)
			return v153
		}
	}
}
func F_XLogArchiveCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(1056)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_XLogArchiveCleanup_0)
	v12 = v6 + int32(32)
	v17 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_XLogArchiveCleanup_1), v6+int32(16))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = F_unlink(m, v12)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(_a_F_XLogArchiveCleanup_2)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v25 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_XLogArchiveCleanup_1), v6)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = F_unlink(m, v12)
			mBase = m.M
			m.G0 = v6 + int32(1056)
			return
		}
	}
}
func F_XLogFlush(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v321 int32
	_ = v321
	var v325 int64
	_ = v325
	var v328 int64
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v439 int64
	_ = v439
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v523 int32
	_ = v523
	var v525 int64
	_ = v525
	var v528 int64
	_ = v528
	var v529 int64
	_ = v529
	var v532 int64
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	v1 = l0
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+308))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[1]))
	if v22 < v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return
L2:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2]))
	if base.Ui64(v1) <= base.Ui64(v44) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	F_UpdateMinRecoveryPoint(m, v1, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[3])))
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if v22 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	v32 = base.B2i32(v30 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[3])) = uint8(v32)
	if v30 != int32(2) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[1])) = int32(1)
	goto L2
L10:
	;
	goto L9
L11:
	;
	goto L3
L12:
	;
	return
L13:
	;
	goto L1
L14:
	;
	v46 = int32(_a_F_XLogFlush_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[4])) = v48 + int32(1)
	v64 = v1
	goto L16
L15:
	;
	v482 = int32(_a_F_XLogFlush_0)
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[4]))
	v485 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[4])) = v484 - v485
	v490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[3])))
	if v490 == v485 {
		goto L95
	} else {
		goto L96
	}
L16:
	;
	v65 = int32(_a_F_XLogFlush_1)
	v66 = int32(_a_F_XLogFlush_2)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v68 = int64(0)
	v71 = base.AtomicRmwCmpxchg64(m, v67, int32(280), v68, v68)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v79 = base.AtomicRmwCmpxchg64(m, v75, int32(272), v68, v68)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[5])) = v79
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2]))
	if base.Ui64(v1) <= base.Ui64(v82) {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v311 = int32(_a_F_XLogFlush_1)
	v312 = int32(_a_F_XLogFlush_2)
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v314 = int64(0)
	v317 = base.AtomicRmwCmpxchg64(m, v313, int32(280), v314, v314)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2])) = v317
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v325 = base.AtomicRmwCmpxchg64(m, v321, int32(272), v314, v314)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[5])) = v325
	v328 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2]))
	if base.Ui64(v328) < base.Ui64(v1) {
		goto L70
	} else {
		goto L71
	}
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v88 = base.AtomicRmwXchg32(m, v85, int32(440), int32(1))
	if v88 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	F_s_lock(m, v90+int32(440), int32(_a_F_XLogFlush_3), int32(2843), int32(_a_F_XLogFlush_4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)+184))
	v101 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v99)+440)), uint32(v101))
	if base.Ui64(v100) < base.Ui64(v64) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v105 = v64
	goto L25
L24:
	;
	v105 = v100
	goto L25
L25:
	;
	v106 = F_WaitXLogInsertionsToFinish(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[6]))
	v111 = v109 + int32(1024)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[7]))
	if v113 < int32(200) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v308 == int32(0) {
		v64 = v105
		goto L16
	} else {
		goto L69
	}
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[8]))
	v118 = int32(_a_F_XLogFlush_5)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[9])) = v120 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v126 = v124
	goto L31
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L66
	}
L31:
	;
	v141 = v126 & int32(_a_F_XLogFlush_6)
	if v141 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(v141 == int32(0)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v142 = v126
	goto L35
L34:
	;
	v142 = v126 | int32(_a_F_XLogFlush_7)
	goto L35
L35:
	;
	v146 = base.AtomicRmwCmpxchg32(m, v111, int32(4), v126, v142)
	if v126 != v146 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v126 = v146
	goto L31
L37:
	;
	goto L38
L38:
	;
	goto L32
L39:
	;
	F_LWLockQueueSelf(m, v111, int32(2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v266 = int32(_a_F_XLogFlush_8)
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[7]))
	v269 = v267 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v269)+uint32(_c_F_XLogFlush[10]))) = v111
	v274 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[7])) = v267 + v274
	*(*int32)(unsafe.Add(mBase, uint32(v269)+uint32(_c_F_XLogFlush[11]))) = int32(0)
	v308 = v274
	goto L27
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v155 = v153
	goto L43
L43:
	;
	v170 = v155 & int32(_a_F_XLogFlush_6)
	if v170 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if base.B2i32(v170 == int32(0)) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	v171 = v155
	goto L47
L46:
	;
	v171 = v155 | int32(_a_F_XLogFlush_7)
	goto L47
L47:
	;
	v175 = base.AtomicRmwCmpxchg32(m, v111, int32(4), v155, v171)
	if v155 != v175 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v155 = v175
	goto L43
L49:
	;
	goto L50
L50:
	;
	goto L44
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[12]))
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111))))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v182 | int32(16777216)
	v188 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_LWLockDequeueSelf(m, v111)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L65
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	F_PGSemaphoreLock(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L56
	}
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[12]))
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v207
	if v207 < v188 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+74)))
	if v204 != 0 {
		v188 = v188 + int32(1)
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v212 = v188
	goto L61
L59:
	;
	goto L60
L60:
	;
	v244 = int32(_a_F_XLogFlush_5)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[9])) = v246 - int32(1)
	v308 = int32(0)
	goto L27
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	F_PGSemaphoreUnlock(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v227 = int32(1)
	if base.Ui32(v227) < base.Ui32(v212) {
		v212 = v212 - v227
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L41
L66:
	;
	F_errmsg_internal(m, int32(_a_F_XLogFlush_9), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_XLogFlush_10), int32(1425), int32(_a_F_XLogFlush_11))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	goto L17
L70:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[13]))
	if v331 <= int32(0) {
		v439 = v106
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[6]))
	F_LWLockRelease(m, v464+int32(1024))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L12
	} else {
		goto L94
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v439
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v439
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v439
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v439
	F_XLogWrite(m, v16+int32(16), v20, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L12
	} else {
		goto L93
	}
L74:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[14])))
	if v335&int32(1) == int32(0) {
		v439 = v106
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[15]))
	if v341 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[16]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v344 <= int32(0) {
		v400 = v2
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v420 = int32(1)
	goto L78
L78:
	;
	if v420 == int32(0) {
		v439 = v106
		goto L73
	} else {
		goto L91
	}
L79:
	;
	v420 = base.B2i32(v341 <= v400)
	goto L78
L80:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[8]))
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[17]))
	v355 = int32(0)
	v362 = v2
	goto L81
L81:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v343+int32(36)+v355<<(uint(int32(2))%32))))
	if v370 == int32(-1) {
		v388 = v362
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v400 = v388
	goto L79
L83:
	;
	v390 = v355 + int32(1)
	if v390 != v344 {
		v355 = v390
		v362 = v388
		goto L81
	} else {
		goto L90
	}
L84:
	;
	v375 = v352 + v370*int32(640)
	if v375 == v350 {
		v388 = v362
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)+36))
	if v377 == int32(0) {
		v388 = v362
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v375)+44))
	if v380 == int32(0) {
		v388 = v362
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)+92))
	if v383 != 0 {
		v388 = v362
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v385 = v362 + int32(1)
	if v341 <= v385 {
		v400 = v385
		goto L79
	} else {
		goto L89
	}
L89:
	;
	v388 = v385
	goto L83
L90:
	;
	goto L82
L91:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[13]))
	F_pg_usleep(m, v424)
	mBase = m.M
	v426 = F_WaitXLogInsertionsToFinish(m, v106)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	v439 = v426
	goto L73
L93:
	;
	goto L72
L94:
	;
	goto L15
L95:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[0]))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+316))
	v497 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[3])) = uint8(base.B2i32(v496 != v497))
	v502 = base.B2i32(v496 == v497)
	goto L97
L96:
	;
	v502 = v485
	goto L97
L97:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[18])))
	if v504 != int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v518 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2]))
	if base.Ui64(v1) <= base.Ui64(v518) {
		goto L1
	} else {
		goto L102
	}
L99:
	;
	v508 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFlush[18])) = uint8(v508)
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFlush[19]))
	if v511 <= v508 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_WalSndWakeup(m, int32(1), v502)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v525 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFlush[2]))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+12)) = uint32(v525)
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+4)) = uint32(v1)
	v528 = int64(32)
	v529 = int64(base.Ui64(v1) >> (uint(v528) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16))) = uint32(v529)
	v532 = int64(base.Ui64(v525) >> (uint(v528) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+8)) = uint32(v532)
	F_errmsg_internal(m, int32(_a_F_XLogFlush_12), v16)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_XLogFlush_3), int32(2942), int32(_a_F_XLogFlush_4))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogPrefetcherComputeStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+124))
	if v7 != 0 {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+120))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
		v13 = base.I32_wrap_i64(v8 - v10)
	} else {
		v13 = int32(0)
	}
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherComputeStats[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v18 + v15
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v24 - int64(-8192)
	return
}
func F_XLogReadDetermineTimeline(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+1192)))
	v14 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	v17 = v13 + v14*v15
	if v17 != l1 {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
		if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
			m.G0 = v11 + int32(16)
			return
		} else {
			v36 = int32(0)
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
			if base.B2i32(v32 == v36)|(base.B2i32(v38 == int64(0))|base.B2i32(l3 == v32)) == v36 {
				v48 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
				v49 = base.I64_div_u_s(v38, v14)
				if base.Ui64(v48) < base.Ui64(v49) {
					m.G0 = v11 + int32(16)
					return
				} else {
					v51 = F_readTimeLineHistory(m, l3)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
						v54 = base.I64_div_u_s(l1, v53)
						v55 = int64(1)
						v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
							v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
								F_list_free_deep(m, v51)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v72 = F_errstart(m, int32(12), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if v72 == int32(0) {
											m.G0 = v11 + int32(16)
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
											v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
											v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
											F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
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
				v51 = F_readTimeLineHistory(m, l3)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
					v54 = base.I64_div_u_s(l1, v53)
					v55 = int64(1)
					v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
						v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
							F_list_free_deep(m, v51)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v72 = F_errstart(m, int32(12), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									if v72 == int32(0) {
										m.G0 = v11 + int32(16)
										return
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
										v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
										v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
										F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
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
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v19 == int32(0) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
			if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
				m.G0 = v11 + int32(16)
				return
			} else {
				v36 = int32(0)
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
				if base.B2i32(v32 == v36)|(base.B2i32(v38 == int64(0))|base.B2i32(l3 == v32)) == v36 {
					v48 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
					v49 = base.I64_div_u_s(v38, v14)
					if base.Ui64(v48) < base.Ui64(v49) {
						m.G0 = v11 + int32(16)
						return
					} else {
						v51 = F_readTimeLineHistory(m, l3)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
							v54 = base.I64_div_u_s(l1, v53)
							v55 = int64(1)
							v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
								v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
									F_list_free_deep(m, v51)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										v72 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											if v72 == int32(0) {
												m.G0 = v11 + int32(16)
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
												v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
												v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
												F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
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
					v51 = F_readTimeLineHistory(m, l3)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
						v54 = base.I64_div_u_s(l1, v53)
						v55 = int64(1)
						v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
							v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
								F_list_free_deep(m, v51)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v72 = F_errstart(m, int32(12), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if v72 == int32(0) {
											m.G0 = v11 + int32(16)
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
											v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
											v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
											F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
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
		} else {
			v24 = int32(_a_F_XLogReadDetermineTimeline_3)
			if base.Ui32(v24) <= base.Ui32(l2) {
				v27 = v24
			} else {
				v27 = l2
			}
			if base.Ui64(l1+base.I64_extend_i32_u(v27)) <= base.Ui64(l1+base.I64_extend_i32_u(v19)) {
				m.G0 = v11 + int32(16)
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
				if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
					m.G0 = v11 + int32(16)
					return
				} else {
					v36 = int32(0)
					v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
					if base.B2i32(v32 == v36)|(base.B2i32(v38 == int64(0))|base.B2i32(l3 == v32)) == v36 {
						v48 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
						v49 = base.I64_div_u_s(v38, v14)
						if base.Ui64(v48) < base.Ui64(v49) {
							m.G0 = v11 + int32(16)
							return
						} else {
							v51 = F_readTimeLineHistory(m, l3)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
								v54 = base.I64_div_u_s(l1, v53)
								v55 = int64(1)
								v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
									v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
										F_list_free_deep(m, v51)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											v72 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												if v72 == int32(0) {
													m.G0 = v11 + int32(16)
													return
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
													v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
													v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
													F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16)
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
						v51 = F_readTimeLineHistory(m, l3)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
							v54 = base.I64_div_u_s(l1, v53)
							v55 = int64(1)
							v60 = F_tliOfPointInHistory(m, (v54+v55)*v53-v55, v51)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v60
								v65 = F_tliSwitchPoint(m, v60, v51, l0+int32(1240))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v65
									F_list_free_deep(m, v51)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										v72 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											if v72 == int32(0) {
												m.G0 = v11 + int32(16)
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
												v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v77)
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
												v81 = int64(base.Ui64(v77) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v81)
												F_errmsg_internal(m, int32(_a_F_XLogReadDetermineTimeline_0), v11)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_XLogReadDetermineTimeline_1), int32(800), int32(_a_F_XLogReadDetermineTimeline_2))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
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
func F_XLogRecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v13
	v17 = base.I32_div_u_s(l1, int32(4069))
	v19 = base.I32_div_u_s(l1, int32(16556761))
	v25 = F_XLogReadBufferExtended(m, v9, int32(1), v17+v19+int32(2), int32(3), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v25 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+14)))
	if v48 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRecordPageWithFreeSpace[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(v25^int32(-1))<<(uint(int32(2))%32))))
	v47 = v39
	goto L4
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRecordPageWithFreeSpace[1]))
	v47 = v41 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v51 = int32(_a_F_XLogRecordPageWithFreeSpace_0)
	v52 = int32(0)
	if v52|(v47&int32(3)|int32(1)) == v52 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(_a_F_XLogRecordPageWithFreeSpace_1)) < base.Ui32(l2) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+10)) = int32(_a_F_XLogRecordPageWithFreeSpace_2)
	v92 = int32(_a_F_XLogRecordPageWithFreeSpace_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+18)) = uint16(v92)
	v98 = int32(_a_F_XLogRecordPageWithFreeSpace_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)) = uint16(v98)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+14)) = uint16(v98)
	goto L11
L13:
	;
	goto L16
L14:
	;
	goto L15
L15:
	;
	goto L21
L16:
	;
	v69 = v47 + v51
	v71 = v47 + int32(4)
	if base.Ui32(v71) < base.Ui32(v69) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v73 = v69
	goto L19
L18:
	;
	v73 = v71
	goto L19
L19:
	;
	v78 = (v47^int32(-1)+v73)&int32(-4) + int32(4)
	if v78 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	base.MemoryFill(m, v47, int32(0), v78)
	goto L12
L21:
	;
	base.MemoryFill(m, v47, int32(0), v51)
	goto L12
L22:
	;
	v109 = int32(-1)
	goto L24
L23:
	;
	v109 = int32(base.Ui32(l2) >> (uint(int32(5)) % 32))
	goto L24
L24:
	;
	v111 = v109 & int32(255)
	v117 = v47 + int32(28)
	v119 = l1 - v17*int32(4069) + int32(4095)
	v120 = v117 + v119
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v121 != v111 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v213 != 0 {
		goto L56
	} else {
		goto L57
	}
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v111)
	v130 = v119
	goto L29
L27:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.Ui32(v123) < base.Ui32(v111) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v213 = int32(0)
	goto L25
L29:
	;
	v134 = int32(1)
	v135 = v130 - v134
	v136 = int32(2)
	v137 = base.I32_div_s(v135, v136)
	v139 = v137 << (uint(v134) % 32)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v139)+1)))
	v143 = v139 + v136
	if base.Ui32(v143) <= base.Ui32(int32(_a_F_XLogRecordPageWithFreeSpace_4)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.Ui32(v162) < base.Ui32(v111) {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v147 = v141 & int32(255)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v143))))
	if base.Ui32(v149) < base.Ui32(v147) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v152 = v141
	goto L33
L33:
	;
	v154 = v137 + v117
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v155 != v152&int32(255) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v151 = v147
	goto L36
L35:
	;
	v151 = v149
	goto L36
L36:
	;
	v152 = v151
	goto L33
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v152)
	if int32(1) < v135 {
		v130 = v137
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L30
L40:
	;
	goto L39
L41:
	;
	v168 = int32(4094)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v213 = int32(1)
	goto L25
L44:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v168) {
		v189 = int32(0)
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v190 = v168 + v117
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v191 != v189&int32(255) {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v178 = v168 << (uint(int32(1)) % 32)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(29)+v178))))
	if v168 == int32(4081) {
		v189 = v180
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v117)+2)))
	if base.Ui32(v184) < base.Ui32(v180) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = v180
	goto L51
L50:
	;
	v186 = v184
	goto L51
L51:
	;
	v189 = v186
	goto L46
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v189)
	goto L54
L53:
	;
	goto L54
L54:
	;
	if v168 != 0 {
		v168 = v168 - int32(1)
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L45
L56:
	;
	F_MarkBufferDirtyHint(m, v25, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_UnlockReleaseBuffer(m, v25)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[0]))
	v15 = v12 + l0*int32(_a_F_XLogRegisterBufData_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[1]))
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[2]))
		if v20 <= v18 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_XLogRegisterBufData_1), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v71
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v74
					F_errdetail_internal(m, int32(_a_F_XLogRegisterBufData_2), v9+int32(16))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_XLogRegisterBufData_3), int32(428), int32(_a_F_XLogRegisterBufData_4))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
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
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
			if base.Ui32(int32(_a_F_XLogRegisterBufData_5)) <= base.Ui32(v22+l2|l2) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_XLogRegisterBufData_1), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v94
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_XLogRegisterBufData_6)
						F_errdetail_internal(m, int32(_a_F_XLogRegisterBufData_7), v9)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_XLogRegisterBufData_3), int32(433), int32(_a_F_XLogRegisterBufData_4))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
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
				*(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[1])) = v18 + int32(1)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBufData[3]))
				v35 = v32 + v18*int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v35
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v41 + l2
				m.G0 = v9 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
			F_errmsg_internal(m, int32(_a_F_XLogRegisterBufData_8), v9+int32(32))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_XLogRegisterBufData_3), int32(416), int32(_a_F_XLogRegisterBufData_4))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
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
