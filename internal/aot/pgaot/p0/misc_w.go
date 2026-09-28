package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_WaitLSNCleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[0]))
	if v6 == int32(0) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[1]))
		v13 = v6 + v10<<(uint(int32(5))%32)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+96)))
		if v14 != int32(1) {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[2]))
			v23 = F_LWLockAcquire(m, v19+int32(_a_F_WaitLSNCleanup_0), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = v13 + int32(80)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)))
				if v27 == int32(1) {
					v31 = v17 * int32(12)
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[0]))
					F_pairingheap_remove(m, v31+v33+int32(32), v13+int32(100))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v41)
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[0]))
						v45 = v44 + v31
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
						if v46 != 0 {
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[0]))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(32))+8))
							v54 = *(*int64)(unsafe.Add(mBase, uint32(v51-int32(20))))
							v56 = v48
							v57 = v54
						} else {
							v56 = v44
							v57 = int64(-1)
						}
						v62 = base.AtomicRmwXchg64(m, v56+v17<<(uint(int32(3))%32), int32(0), v57)
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[2]))
						F_LWLockRelease(m, v67+int32(_a_F_WaitLSNCleanup_0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNCleanup[2]))
					F_LWLockRelease(m, v67+int32(_a_F_WaitLSNCleanup_0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_WaitLSNShmemExit(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	F_WaitLSNCleanup(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_WaitLSNShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(-1)
	v7 = v3 + int32(32)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(444)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(-1)
	v19 = v15 + int32(44)
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(444)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = int64(-1)
	v31 = v27 + int32(56)
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(444)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = int64(-1)
	v43 = v39 + int32(68)
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(444)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[1]))
	v55 = v51<<(uint(int32(5))%32) + int32(1216)
	if v55 != 0 {
		v57 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLSNShmemInit[0]))
		base.MemoryFill(m, v57+int32(80), int32(0), v55)
	} else {
	}
	return
}
func F_WakeupRecovery(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_WakeupRecovery[0]))
	v4 = v2 + int32(4)
	v5 = int32(0)
	v8 = base.AtomicRmwOr32(m, v5, int32(_a_F_WakeupRecovery_0), v5)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
	v12 = int32(0)
	v15 = base.AtomicRmwOr32(m, v12, int32(_a_F_WakeupRecovery_0), v12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v16 == v12 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_WakeupRecovery[1]))
	if v23 == v19 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_WakeupRecovery[2]))
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v53 = F_pgmem_kill(m, v19, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v27 + int32(16)
	goto L1
L10:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)) = uint8(v33)
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_WakeupRecovery[3]))
	v41 = F_write(m, v37, v27+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v41 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_WakeupRecovery[4]))
	if v45 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_wcslen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v5 = l0
	for {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v9 != 0 {
			v5 = v5 + int32(4)
			continue
		} else {
			break
		}
		break
	}
	return (v5 - l0) >> (uint(int32(2)) % 32)
}
func F_wcstombs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	v4 = int32(0)
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v16 = v12 + int32(12)
	v17 = m.G0
	v19 = v17 - v11
	m.G0 = v19
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v527 = int32(16)
	m.G0 = v19 + v527
	m.G0 = v12 + v527
	return v523
L2:
	;
	v523 = l2 - v508
	goto L1
L3:
	;
	if v269 != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v140 = l2
	v141 = l0
	v145 = v139
	goto L37
L5:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 == int32(0) {
		v523 = v4
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v269 = l2
	v270 = l0
	goto L3
L9:
	;
	v27 = v24
	v28 = v23
	v30 = v4
	goto L10
L10:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v27) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v523 = v138
	goto L1
L12:
	;
	v41 = v19 + int32(12)
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v133 = int32(1)
	goto L14
L14:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v138 = v133 + v30
	if v135 != 0 {
		v27 = v135
		v28 = v28 + int32(4)
		v30 = v138
		goto L10
	} else {
		goto L36
	}
L15:
	;
	if v130 == int32(-1) {
		v523 = int32(-1)
		goto L1
	} else {
		goto L35
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v27)
	v130 = int32(1)
	goto L15
L17:
	;
	if base.Ui32(v27) <= base.Ui32(int32(127)) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v127 = int32(1)
	goto L19
L19:
	;
	v130 = v127
	goto L15
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_wcstombs[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_wcstombs[1])) = int32(25)
	v127 = int32(-1)
	goto L19
L22:
	;
	if v27&int32(-128) == int32(_a_F_wcstombs_0) {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v27) <= base.Ui32(int32(2047)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L21
L26:
	;
	v58 = v27&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v58)
	v63 = int32(base.Ui32(v27)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v63)
	v130 = int32(2)
	goto L15
L27:
	;
	goto L28
L28:
	;
	if base.B2i32(v27&int32(-8192) != int32(_a_F_wcstombs_1))&base.B2i32(base.Ui32(int32(_a_F_wcstombs_2)) <= base.Ui32(v27)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v75 = int32(63)
	v77 = int32(128)
	v78 = v27&v75 | v77
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)) = uint8(v78)
	v83 = int32(base.Ui32(v27)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v83)
	v90 = int32(base.Ui32(v27)>>(uint(int32(6))%32))&v75 | v77
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v90)
	v130 = int32(3)
	goto L15
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(v27-int32(_a_F_wcstombs_3)) <= base.Ui32(int32(_a_F_wcstombs_4)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = int32(63)
	v99 = int32(128)
	v100 = v27&v97 | v99
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)) = uint8(v100)
	v105 = int32(base.Ui32(v27)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v105)
	v112 = int32(base.Ui32(v27)>>(uint(int32(6))%32))&v97 | v99
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)) = uint8(v112)
	v119 = int32(base.Ui32(v27)>>(uint(int32(12))%32))&v97 | v99
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v119)
	v130 = int32(4)
	goto L15
L33:
	;
	goto L34
L34:
	;
	goto L21
L35:
	;
	v133 = v130
	goto L14
L36:
	;
	goto L11
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if base.Ui32(v149-int32(128)) <= base.Ui32(int32(-128)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v269 = v259
	v270 = v262
	goto L3
L39:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v265 = v263 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v265
	if base.Ui32(int32(3)) < base.Ui32(v259) {
		v140 = v259
		v141 = v262
		v145 = v265
		goto L37
	} else {
		goto L67
	}
L40:
	;
	if v149 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v149)
	v255 = int32(1)
	v259 = v140 - v255
	v262 = v141 + v255
	goto L39
L43:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v156
	v508 = v140
	goto L2
L44:
	;
	goto L45
L45:
	;
	if v141 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v249 == int32(-1) {
		v523 = int32(-1)
		goto L1
	} else {
		goto L66
	}
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v149)
	v249 = int32(1)
	goto L46
L48:
	;
	if base.Ui32(v149) <= base.Ui32(int32(127)) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v246 = int32(1)
	goto L50
L50:
	;
	v249 = v246
	goto L46
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_wcstombs[0]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v165 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_wcstombs[1])) = int32(25)
	v246 = int32(-1)
	goto L50
L53:
	;
	if v149&int32(-128) == int32(_a_F_wcstombs_0) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v149) <= base.Ui32(int32(2047)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L52
L57:
	;
	v177 = v149&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)) = uint8(v177)
	v182 = int32(base.Ui32(v149)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v182)
	v249 = int32(2)
	goto L46
L58:
	;
	goto L59
L59:
	;
	if base.B2i32(v149&int32(-8192) != int32(_a_F_wcstombs_1))&base.B2i32(base.Ui32(int32(_a_F_wcstombs_2)) <= base.Ui32(v149)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v194 = int32(63)
	v196 = int32(128)
	v197 = v149&v194 | v196
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+2)) = uint8(v197)
	v202 = int32(base.Ui32(v149)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v202)
	v209 = int32(base.Ui32(v149)>>(uint(int32(6))%32))&v194 | v196
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)) = uint8(v209)
	v249 = int32(3)
	goto L46
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v149-int32(_a_F_wcstombs_3)) <= base.Ui32(int32(_a_F_wcstombs_4)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v216 = int32(63)
	v218 = int32(128)
	v219 = v149&v216 | v218
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+3)) = uint8(v219)
	v224 = int32(base.Ui32(v149)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v224)
	v231 = int32(base.Ui32(v149)>>(uint(int32(6))%32))&v216 | v218
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+2)) = uint8(v231)
	v238 = int32(base.Ui32(v149)>>(uint(int32(12))%32))&v216 | v218
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)) = uint8(v238)
	v249 = int32(4)
	goto L46
L64:
	;
	goto L65
L65:
	;
	goto L52
L66:
	;
	v259 = v140 - v249
	v262 = v141 + v249
	goto L39
L67:
	;
	goto L38
L68:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v279 = v269
	v280 = v270
	v282 = v278
	goto L71
L69:
	;
	goto L70
L70:
	;
	v523 = l2
	goto L1
L71:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if base.Ui32(v288-int32(128)) <= base.Ui32(int32(-128)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L70
L73:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v497 = v495 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v497
	if v491 != 0 {
		v279 = v491
		v280 = v494
		v282 = v497
		goto L71
	} else {
		goto L122
	}
L74:
	;
	if v288 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v288)
	v487 = int32(1)
	v491 = v279 - v487
	v494 = v280 + v487
	goto L73
L77:
	;
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v295
	v508 = v279
	goto L2
L78:
	;
	goto L79
L79:
	;
	v301 = v19 + int32(12)
	if v301 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	if v390 == int32(-1) {
		v523 = int32(-1)
		goto L1
	} else {
		goto L100
	}
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v288)
	v390 = int32(1)
	goto L80
L82:
	;
	if base.Ui32(v288) <= base.Ui32(int32(127)) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	v387 = int32(1)
	goto L84
L84:
	;
	v390 = v387
	goto L80
L85:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_wcstombs[0]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v306 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_wcstombs[1])) = int32(25)
	v387 = int32(-1)
	goto L84
L87:
	;
	if v288&int32(-128) == int32(_a_F_wcstombs_0) {
		goto L81
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v288) <= base.Ui32(int32(2047)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L86
L91:
	;
	v318 = v288&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)) = uint8(v318)
	v323 = int32(base.Ui32(v288)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v323)
	v390 = int32(2)
	goto L80
L92:
	;
	goto L93
L93:
	;
	if base.B2i32(v288&int32(-8192) != int32(_a_F_wcstombs_1))&base.B2i32(base.Ui32(int32(_a_F_wcstombs_2)) <= base.Ui32(v288)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v335 = int32(63)
	v337 = int32(128)
	v338 = v288&v335 | v337
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+2)) = uint8(v338)
	v343 = int32(base.Ui32(v288)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v343)
	v350 = int32(base.Ui32(v288)>>(uint(int32(6))%32))&v335 | v337
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)) = uint8(v350)
	v390 = int32(3)
	goto L80
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(v288-int32(_a_F_wcstombs_3)) <= base.Ui32(int32(_a_F_wcstombs_4)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v357 = int32(63)
	v359 = int32(128)
	v360 = v288&v357 | v359
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+3)) = uint8(v360)
	v365 = int32(base.Ui32(v288)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v365)
	v372 = int32(base.Ui32(v288)>>(uint(int32(6))%32))&v357 | v359
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+2)) = uint8(v372)
	v379 = int32(base.Ui32(v288)>>(uint(int32(12))%32))&v357 | v359
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)) = uint8(v379)
	v390 = int32(4)
	goto L80
L98:
	;
	goto L99
L99:
	;
	goto L86
L100:
	;
	if base.Ui32(v279) < base.Ui32(v390) {
		v508 = v279
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v280 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v491 = v279 - v390
	v494 = v280 + v390
	goto L73
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v394)
	goto L102
L104:
	;
	if base.Ui32(v394) <= base.Ui32(int32(127)) {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L102
L107:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_wcstombs[0]))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if v399 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_wcstombs[1])) = int32(25)
	goto L106
L109:
	;
	if v394&int32(-128) == int32(_a_F_wcstombs_0) {
		goto L103
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if base.Ui32(v394) <= base.Ui32(int32(2047)) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L108
L113:
	;
	v411 = v394&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)) = uint8(v411)
	v416 = int32(base.Ui32(v394)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v416)
	goto L102
L114:
	;
	goto L115
L115:
	;
	if base.B2i32(v394&int32(-8192) != int32(_a_F_wcstombs_1))&base.B2i32(base.Ui32(int32(_a_F_wcstombs_2)) <= base.Ui32(v394)) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v428 = int32(63)
	v430 = int32(128)
	v431 = v394&v428 | v430
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+2)) = uint8(v431)
	v436 = int32(base.Ui32(v394)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v436)
	v443 = int32(base.Ui32(v394)>>(uint(int32(6))%32))&v428 | v430
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)) = uint8(v443)
	goto L102
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v394-int32(_a_F_wcstombs_3)) <= base.Ui32(int32(_a_F_wcstombs_4)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v450 = int32(63)
	v452 = int32(128)
	v453 = v394&v450 | v452
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+3)) = uint8(v453)
	v458 = int32(base.Ui32(v394)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v458)
	v465 = int32(base.Ui32(v394)>>(uint(int32(6))%32))&v450 | v452
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+2)) = uint8(v465)
	v472 = int32(base.Ui32(v394)>>(uint(int32(12))%32))&v450 | v452
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)) = uint8(v472)
	goto L102
L120:
	;
	goto L121
L121:
	;
	goto L108
L122:
	;
	goto L72
}
func F_weak_input_status(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = int64(*(*uint8)(unsafe.Add(mBase, _c_F_weak_input_status[0])))
	return v3
}
func F_websearch_to_tsquery(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14391(m, l0, int32(1277))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_width_bucket_float8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int64
	_ = v18
	var v23 float64
	_ = v23
	var v27 float64
	_ = v27
	var v31 int64
	_ = v31
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v84 float64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 float64
	_ = v126
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if int32(0) < v10 {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v13&int64(9223372036854775807)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v178 = m.ExcPending
			if v178 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(386138242))
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_width_bucket_float8_0), int32(0))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_2), int32(_a_F_width_bucket_float8_3))
						mBase = m.M
						v190 = m.ExcPending
						if v190 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v18&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(386138242))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_width_bucket_float8_0), int32(0))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_2), int32(_a_F_width_bucket_float8_3))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v23 = base.F64_reinterpret_i64(v13)
				if base.F64_eq(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(386138242))
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_width_bucket_float8_4), int32(0))
							mBase = m.M
							v202 = m.ExcPending
							if v202 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_5), int32(_a_F_width_bucket_float8_3))
								mBase = m.M
								v207 = m.ExcPending
								if v207 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v27 = base.F64_reinterpret_i64(v18)
					if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(386138242))
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_width_bucket_float8_4), int32(0))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_5), int32(_a_F_width_bucket_float8_3))
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						v32 = base.F64_reinterpret_i64(v31)
						if base.F64_lt(v23, v27) != 0 {
							v35 = int32(0)
							if base.B2i32(base.F64_le(v27, v32) == v35)&base.B2i32(base.Ui64(v31&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) == v35 {
								v45 = v10 + int32(1)
								if v10 <= v45 {
									v151 = v45
									return base.I64_extend_i32_s(v151)
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50331778))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_float8_6), int32(0))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_7), int32(_a_F_width_bucket_float8_3))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								if base.F64_gt(v23, v32) != 0 {
									v151 = int32(0)
									return base.I64_extend_i32_s(v151)
								} else {
									v68 = base.F64_sub(v27, v23)
									if base.F64_ne(base.F64_abs(v68), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v84 = base.F64_div(base.F64_sub(v32, v23), v68)
									} else {
										v74 = float64(0.5)
										v77 = base.F64_mul(v23, v74)
										v84 = base.F64_div(base.F64_sub(base.F64_mul(v32, v74), v77), base.F64_sub(base.F64_mul(v27, v74), v77))
									}
									v86 = base.I32_trunc_sat_f64_s(base.F64_mul(v84, base.F64_convert_i32_u(v10)))
									v88 = v10 - int32(1)
									if v86 < v88 {
										v90 = v86
									} else {
										v90 = v88
									}
									return base.I64_extend_i32_s(v90 + int32(1))
								}
							}
						} else {
							if base.F64_gt(v23, v27) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v211 = m.ExcPending
								if v211 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v214 = m.ExcPending
									if v214 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_width_bucket_float8_8), int32(0))
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_9), int32(_a_F_width_bucket_float8_3))
											mBase = m.M
											v223 = m.ExcPending
											if v223 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v31&int64(9223372036854775807)))|base.F64_lt(v23, v32) != 0 {
									v151 = int32(0)
									return base.I64_extend_i32_s(v151)
								} else {
									if base.F64_ge(v27, v32) != 0 {
										v107 = v10 + int32(1)
										if v10 <= v107 {
											v151 = v107
											return base.I64_extend_i32_s(v151)
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_width_bucket_float8_6), int32(0))
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_10), int32(_a_F_width_bucket_float8_3))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									} else {
										v126 = base.F64_sub(v23, v27)
										if base.F64_ne(base.F64_abs(v126), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v142 = base.F64_div(base.F64_sub(v23, v32), v126)
										} else {
											v132 = float64(0.5)
											v133 = base.F64_mul(v23, v132)
											v142 = base.F64_div(base.F64_sub(v133, base.F64_mul(v32, v132)), base.F64_sub(v133, base.F64_mul(v27, v132)))
										}
										v144 = base.I32_trunc_sat_f64_s(base.F64_mul(v142, base.F64_convert_i32_u(v10)))
										v146 = v10 - int32(1)
										if v144 < v146 {
											v148 = v144
										} else {
											v148 = v146
										}
										v151 = v148 + int32(1)
										return base.I64_extend_i32_s(v151)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v161 = m.ExcPending
		if v161 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(386138242))
			mBase = m.M
			v164 = m.ExcPending
			if v164 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_width_bucket_float8_11), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_width_bucket_float8_1), int32(_a_F_width_bucket_float8_12), int32(_a_F_width_bucket_float8_3))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int64(0)
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
func F_win866_to_koi8r(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_win866_to_koi8r_0), int32(22), int32(20))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_write_nondefault_variables(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v10 = m.G0
	v12 = v10 - int32(144)
	m.G0 = v12
	if l0 == int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(15)
	goto L3
L2:
	;
	v18 = int32(21)
	goto L3
L3:
	;
	v21 = F_AllocateFile(m, int32(_a_F_write_nondefault_variables_0), int32(_a_F_write_nondefault_variables_1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L68
	}
L5:
	;
	m.G0 = v12 + int32(144)
	return
L6:
	;
	return
L7:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = F_errstart(m, v18, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_write_nondefault_variables[0]))
	v44 = int32(0)
	if base.B2i32(v43 == v44)|base.B2i32(v43 == int32(_a_F_write_nondefault_variables_2)) == v44 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_write_nondefault_variables_0)
	F_errmsg(m, int32(_a_F_write_nondefault_variables_3), v12)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_write_nondefault_variables_4), int32(_a_F_write_nondefault_variables_5), int32(_a_F_write_nondefault_variables_6))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L5
L16:
	;
	v54 = v43
	goto L19
L17:
	;
	goto L18
L18:
	;
	v212 = F_FreeFile(m, v21)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L59
	}
L19:
	;
	v61 = v54 - int32(68)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v62
	v67 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_7), v12+int32(128))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	F_do_putc(m, int32(0), v21)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(44))))
	switch v74 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	default:
		goto L23
	}
L23:
	;
	F_do_putc(m, int32(0), v21)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L48
	}
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v54)+36))
	if v115 == int32(0) {
		goto L4
	} else {
		goto L38
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v104 == int32(0) {
		goto L23
	} else {
		goto L36
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v95)))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+64)) = v96
	v101 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_8), v12-int32(-64))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L35
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v88
	v93 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_9), v12+int32(48))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L34
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_10), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v85 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_11), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	goto L23
L33:
	;
	goto L23
L34:
	;
	goto L23
L35:
	;
	goto L23
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v104
	v111 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_7), v12+int32(80))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L23
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v118 == int32(0) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v114 != v121 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v123 = v115
	goto L43
L41:
	;
	v145 = v118
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v145
	v152 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_7), v12+int32(112))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L47
	}
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if v132 == int32(0) {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	v145 = v132
	goto L42
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if v114 != v135 {
		v123 = v123 + int32(12)
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L23
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v166 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v166
	v171 = F_pg_fprintf(m, v21, int32(_a_F_write_nondefault_variables_7), v12+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_do_putc(m, int32(0), v21)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v180 = F_fwrite(m, v54+int32(24), int32(1), int32(4), v21)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v186 = F_fwrite(m, v54-int32(36), int32(1), int32(4), v21)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v192 = F_fwrite(m, v54-int32(28), int32(1), int32(4), v21)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v198 = F_fwrite(m, v54-int32(20), int32(1), int32(4), v21)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v200 != int32(_a_F_write_nondefault_variables_2) {
		v54 = v200
		goto L19
	} else {
		goto L58
	}
L58:
	;
	goto L20
L59:
	;
	if v212 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v215 = F_errstart(m, v18, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v235 = F_rename(m, int32(_a_F_write_nondefault_variables_0), int32(_a_F_write_nondefault_variables_12))
	mBase = m.M
	goto L5
L63:
	;
	if v215 == int32(0) {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_write_nondefault_variables_0)
	F_errmsg(m, int32(_a_F_write_nondefault_variables_3), v12+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_write_nondefault_variables_4), int32(_a_F_write_nondefault_variables_13), int32(_a_F_write_nondefault_variables_6))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L5
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v114
	F_errmsg_internal(m, int32(_a_F_write_nondefault_variables_14), v12+int32(96))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_write_nondefault_variables_4), int32(2938), int32(_a_F_write_nondefault_variables_15))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_write_stderr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_write_stderr[0]))
	v14 = m.G0
	v16 = v14 - int32(1056)
	m.G0 = v16
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_write_stderr[1])) = int32(28)
		m.G0 = v16 + int32(1056)
		v63 = F_fflush(m, v13)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	} else {
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+1052)) = uint8(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1048)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1044)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1040)) = v16 + int32(1024)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1036)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1032)) = v16
		F_dopr(m, v16+int32(1032), l0, l1)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1052)))
			if v38 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1032))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1036))
				if v41 != v42 {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1044))
					v52 = F_fwrite(m, v42, int32(1), v41-v42, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v16 + int32(1056)
						v63 = F_fflush(m, v13)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v16 + int32(1056)
					v63 = F_fflush(m, v13)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			} else {
				m.G0 = v16 + int32(1056)
				v63 = F_fflush(m, v13)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	}
}
