package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LWLockDequeueSelf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(536870912)
	v14 = base.AtomicRmwOr32(m, l0, int32(4), v12)
	if v14&v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v14
	goto L4
L2:
	;
	goto L3
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+344)))
	if v94 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(_a_F_LWLockDequeueSelf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(860)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_LWLockDequeueSelf_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	if v18&int32(536870912) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v60 = int32(_a_F_LWLockDequeueSelf_2)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v63 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v80 = int32(536870912)
	v82 = base.AtomicRmwOr32(m, l0, int32(4), v80)
	if v82&v80 != 0 {
		v18 = v82
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1])) = v78
	goto L15
L17:
	;
	if int32(999) < v61 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v61 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v68 = int32(900)
	if v68 <= v61 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = v68
	goto L23
L22:
	;
	v71 = v61
	goto L23
L23:
	;
	v78 = v71 + int32(100)
	goto L16
L24:
	;
	v78 = v61 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[3]))
	v104 = v99 + v101*int32(768)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+348))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+352))
	if v106 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v132 != int32(-1) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v105 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)+352))
	v115 = v110
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v106*int32(768))+348)) = v105
	v115 = v106
	goto L29
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+348)) = int64(0)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v115
	goto L33
L35:
	;
	goto L36
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v105*int32(768))+352)) = v115
	goto L33
L37:
	;
	v143 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	if v94 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) <= v135 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v140 = base.AtomicRmwAnd32(m, l0, int32(4), int32(2147483647))
	goto L37
L40:
	;
	m.G0 = v10 + int32(32)
	return
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+344)) = uint8(v148)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v152 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-1073741825))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v157 = v154
	v158 = int32(0)
	goto L44
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+332))
	F_PGSemaphoreLock(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	if v158 <= int32(0) {
		goto L40
	} else {
		goto L48
	}
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+344)))
	if v170 != 0 {
		v157 = v169
		v158 = v158 + int32(1)
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v173 = v158
	goto L49
L49:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+332))
	F_PGSemaphoreUnlock(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L51
	}
L50:
	;
	goto L40
L51:
	;
	v185 = int32(1)
	if base.Ui32(v185) < base.Ui32(v173) {
		v173 = v173 - v185
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
}
func F_LWLockInitialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v2 = l1
	v5 = F_GetLWTrancheName(m, v2&int32(_a_F_LWLockInitialize_0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
		return
	}
}
func F_LWLockShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v132 int32
	_ = v132
	var v187 int32
	_ = v187
	var v212 int32
	_ = v212
	var v281 int32
	_ = v281
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_LWLockShmemInit[1]))) = v2
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_LWLockShmemInit[2]))), uint32(v2))
	v17 = v2
	goto L1
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v27 = F_GetLWTrancheName(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v43 = v37
	goto L6
L3:
	;
	return
L4:
	;
	v31 = v26 + v17<<(uint(int32(7))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v31))) = uint16(v17)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(-1)
	v37 = int32(58)
	v39 = v17 + int32(1)
	if v39 != v37 {
		v17 = v39
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
L6:
	;
	v51 = v43 << (uint(int32(7)) % 32)
	v52 = int32(_a_F_LWLockShmemInit_0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v54 = v51 + v53
	v55 = int32(70)
	*(*uint16)(unsafe.Add(mBase, uint32(v54))) = uint16(v55)
	v57 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v57
	v59 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v63 = v62 + v51
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+128)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+132)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v63)+136)) = v59
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v72 = v71 + v51
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+256)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+260)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v72)+264)) = v59
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v81 = v80 + v51
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+384)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+388)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v81)+392)) = v59
	v89 = v43 + int32(4)
	if v89 != int32(186) {
		v43 = v89
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v92 = int32(_a_F_LWLockShmemInit_0)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[4]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[5]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[6]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[7]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[8]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[9]))) = v94
	v106 = int32(71)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[10]))) = uint16(v106)
	v108 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[11]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[12]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[13]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[14]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[15]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[16]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[17]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[18]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[19]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[20]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[21]))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[22]))) = v94
	v132 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[23]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[24]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[25]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[26]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[27]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[28]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[29]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[30]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[31]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[32]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[33]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[34]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[35]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[36]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[37]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[38]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[39]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[40]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[41]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[42]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[43]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[44]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[45]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[46]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[47]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[48]))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_LWLockShmemInit[49]))) = v132
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[50]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[51]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[52]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[53]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[54]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[55]))) = v94
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[56]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[57]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[58]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[59]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[60]))) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[61]))) = v108
	v212 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[62]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[63]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[64]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[65]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[66]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[67]))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[68]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[69]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[70]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[71]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[72]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[73]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[74]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[75]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[76]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[77]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[78]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[79]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[80]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[81]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[82]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[83]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[84]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[85]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[86]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[87]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[88]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[89]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[90]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[91]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[92]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[93]))) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[94]))) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_LWLockShmemInit[95]))) = v132
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[96]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[97]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[98]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[99]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[100]))) = v94
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[101]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[102]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[103]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[104]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[105]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[106]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[107]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[108]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[109]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[110]))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[111]))) = uint16(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[112]))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_LWLockShmemInit[113]))) = v94
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[114]))
	if v319 == v94 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	return
L10:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v322 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v326 = int32(218)
	v332 = v2
	goto L12
L12:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334+v332<<(uint(int32(2))%32))))
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[0]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+uint32(_c_F_LWLockShmemInit[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+uint32(_c_F_LWLockShmemInit[1]))) = v341 + int32(1)
	v346 = v341 * int32(68)
	v347 = v340 + v346
	goto L17
L13:
	;
	goto L9
L14:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v468+v346)+64)) = v326
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	if int32(0) < v471 {
		goto L45
	} else {
		goto L46
	}
L15:
	;
	v464 = F_strlen(m, v453)
	mBase = m.M
	goto L14
L17:
	;
	goto L18
L18:
	;
	v354 = int32(63)
	if (v347^v338)&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v457)
	goto L15
L20:
	;
	v438 = v433
	v439 = v434
	v440 = v435
	goto L41
L21:
	;
	if v428 == int32(0) {
		v453 = v426
		v454 = v427
		goto L19
	} else {
		goto L40
	}
L22:
	;
	v426 = v338
	v427 = v347
	v428 = v354
	goto L21
L23:
	;
	goto L24
L24:
	;
	v358 = int32(0)
	if base.B2i32(v338&int32(3) == v358)|int32(0) == v358 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v394 == int32(0) {
		v453 = v391
		v454 = v392
		goto L19
	} else {
		goto L34
	}
L26:
	;
	v370 = v338
	v371 = v347
	v372 = v354
	goto L29
L27:
	;
	goto L28
L28:
	;
	v391 = v338
	v392 = v347
	v393 = v354
	v394 = int32(1)
	goto L25
L29:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v374)
	if v374 == int32(0) {
		v433 = v370
		v434 = v371
		v435 = v372
		goto L20
	} else {
		goto L31
	}
L30:
	;
	v391 = v385
	v392 = v379
	v393 = v381
	v394 = v383
	goto L25
L31:
	;
	v378 = int32(1)
	v379 = v371 + v378
	v381 = v372 - v378
	v382 = int32(0)
	v383 = base.B2i32(v381 != v382)
	v385 = v370 + v378
	if v385&int32(3) == v382 {
		v391 = v385
		v392 = v379
		v393 = v381
		v394 = v383
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v381 != 0 {
		v370 = v385
		v371 = v379
		v372 = v381
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if base.B2i32(v397 == int32(0))|base.B2i32(base.Ui32(v393) < base.Ui32(int32(4))) != 0 {
		v426 = v391
		v427 = v392
		v428 = v393
		goto L21
	} else {
		goto L35
	}
L35:
	;
	v404 = v391
	v405 = v392
	v406 = v393
	goto L36
L36:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v412 = int32(-2139062144)
	if (int32(16843008)-v409|v409)&v412 != v412 {
		v433 = v404
		v434 = v405
		v435 = v406
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v426 = v420
	v427 = v418
	v428 = v422
	goto L21
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405))) = v409
	v417 = int32(4)
	v418 = v405 + v417
	v420 = v404 + v417
	v422 = v406 - v417
	if base.Ui32(int32(3)) < base.Ui32(v422) {
		v404 = v420
		v405 = v418
		v406 = v422
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v433 = v426
	v434 = v427
	v435 = v428
	goto L20
L41:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	*(*uint8)(unsafe.Add(mBase, uint32(v439))) = uint8(v442)
	if v442 == int32(0) {
		v453 = v438
		v454 = v439
		goto L19
	} else {
		goto L43
	}
L42:
	;
	v453 = v449
	v454 = v447
	goto L19
L43:
	;
	v446 = int32(1)
	v447 = v439 + v446
	v449 = v438 + v446
	v451 = v440 - v446
	if v451 != 0 {
		v438 = v449
		v439 = v447
		v440 = v451
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v476 = v341 + int32(100)
	v479 = v326
	v483 = int32(0)
	goto L48
L46:
	;
	v505 = v326
	goto L47
L47:
	;
	v514 = v332 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v514 < v515 {
		v326 = v505
		v332 = v514
		goto L12
	} else {
		goto L52
	}
L48:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemInit[3]))
	v489 = F_GetLWTrancheName(m, v476&int32(_a_F_LWLockShmemInit_1))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L3
	} else {
		goto L50
	}
L49:
	;
	v505 = v500
	goto L47
L50:
	;
	v493 = v488 + v479<<(uint(int32(7))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v493))) = uint16(v476)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+4)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v493)+8)) = int64(-1)
	v499 = int32(1)
	v500 = v479 + v499
	v502 = v483 + v499
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	if v502 < v503 {
		v479 = v500
		v483 = v502
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L13
}
func F_LWLockUpdateVar(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var __phi133 int32
	_ = __phi133
	var v135 int32
	_ = v135
	var __phi135 int32
	_ = __phi135
	var v136 int32
	_ = v136
	var __phi136 int32
	_ = __phi136
	var v138 int32
	_ = v138
	var __phi138 int32
	_ = __phi138
	var v139 int32
	_ = v139
	var __phi139 int32
	_ = __phi139
	var v142 int32
	_ = v142
	var __phi142 int32
	_ = __phi142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = base.AtomicRmwXchg64(m, l1, int32(0), l2)
	v18 = int32(536870912)
	v20 = base.AtomicRmwOr32(m, l0, int32(4), v18)
	if v20&v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v20
	goto L4
L2:
	;
	goto L3
L3:
	;
	v114 = int32(-1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v115 == v114 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LWLockUpdateVar_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(860)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LWLockUpdateVar_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	if v24&int32(536870912) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v78 = int32(_a_F_LWLockUpdateVar_2)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v81 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v61&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v98 = int32(536870912)
	v100 = base.AtomicRmwOr32(m, l0, int32(4), v98)
	if v100&v98 != 0 {
		v24 = v100
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[0])) = v96
	goto L15
L17:
	;
	if int32(999) < v79 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v79 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v86 = int32(900)
	if v86 <= v79 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v89 = v86
	goto L23
L22:
	;
	v89 = v79
	goto L23
L23:
	;
	v96 = v89 + int32(100)
	goto L16
L24:
	;
	v96 = v79 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	m.G0 = v14 + int32(32)
	return
L27:
	;
	v120 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = v115 * int32(768)
	v126 = v123 + v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+345)))
	if v127 != int32(2) {
		v202 = v114
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v214 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	if v202 == int32(-1) {
		goto L26
	} else {
		goto L48
	}
L31:
	;
	__phi133 = v114
	__phi135 = v125 + v123
	__phi136 = v115
	__phi138 = int32(-1)
	__phi139 = v126
	__phi142 = v123
	v133 = __phi133
	v135 = __phi135
	v136 = __phi136
	v138 = __phi138
	v139 = __phi139
	v142 = __phi142
	goto L32
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+348))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+348))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+352))
	if v145 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v202 = v185
	goto L30
L34:
	;
	if v143 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v143
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+352))
	v154 = v149
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142+v145*int32(768))+348)) = v143
	v154 = v145
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139)+348)) = int64(0)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v172 = v169 + v136*int32(768)
	if v138 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v154
	goto L38
L40:
	;
	goto L41
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v160+v143*int32(768))+352)) = v154
	goto L38
L42:
	;
	v186 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+348)) = v186
	v188 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+344)) = uint8(v188)
	if v144 == v186 {
		v202 = v185
		goto L30
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+352)) = int32(-1)
	v185 = v136
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+352)) = v138
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v180+v138*int32(768))+348)) = v136
	v185 = v133
	goto L42
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v197 = v194 + v144*int32(768)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+345)))
	if v198 == int32(2) {
		__phi133 = v185
		__phi135 = v197
		__phi136 = v144
		__phi138 = v136
		__phi139 = v197
		__phi142 = v194
		v133 = __phi133
		v135 = __phi135
		v136 = __phi136
		v138 = __phi138
		v139 = __phi139
		v142 = __phi142
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v218 = v202
	goto L49
L49:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v233 = v230 + v218*int32(768)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+348))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+352))
	if v235 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L26
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v235*int32(768))+348)) = v234
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v234 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v246+v234*int32(768))+352)) = v235
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v233)+348)) = int64(0)
	v253 = int32(0)
	v256 = base.AtomicRmwOr32(m, v253, int32(_a_F_LWLockUpdateVar_3), v253)
	*(*uint8)(unsafe.Add(mBase, uint32(v233)+344)) = uint8(v253)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v233)+332))
	F_PGSemaphoreUnlock(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	if v234 != int32(-1) {
		v218 = v234
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
}
