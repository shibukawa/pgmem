package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 == int32(1655112) {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(4) != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v18 = v17
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
				v27 = l1 + int32(32)
				*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v21)
				v31 = v18 & int32(65529)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v33)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v35
				if l2 != 0 {
					v38 = v31 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v38)
				} else {
				}
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v40
				return l1
			}
		} else {
			v18 = v9
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
			v21 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
			v27 = l1 + int32(32)
			*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v21)
			v31 = v18 & int32(65529)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v33)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v35
			if l2 != 0 {
				v38 = v31 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v38)
			} else {
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v40
			return l1
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(92238), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519730), int32(1553), int32(405008))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
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
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int64
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_CheckForSerializableConflictOutNeeded(m, l1, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v21 = F_HeapTupleSatisfiesVacuum(m, l2, v20, l3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L11
	}
L5:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v127 == v136 {
		goto L1
	} else {
		goto L40
	}
L6:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v127 = v124
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L37
	}
L8:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+20)))
	v105 = int32(768)
	if v104&v105 != v105 {
		v123 = v103
		goto L6
	} else {
		goto L36
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)))
	if l0 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if l0 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	switch v21 {
	case 0:
		goto L1
	case 1:
		goto L10
	case 2, 4:
		goto L9
	case 3:
		goto L8
	default:
		goto L7
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+20)))
	v25 = int32(768)
	if v24&v25 != v25 {
		v123 = v23
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v127 = int32(2)
	goto L5
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v88))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v79)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31&int32(6272) != int32(4096) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v72 = int32(768)
	if v31&v72 == v72 {
		v79 = int32(2)
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v79 = v32
	goto L14
L19:
	;
	goto L20
L20:
	;
	v37 = int32(0)
	v41 = F_GetMultiXactIdMembers(m, v32, v13+int32(12), v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v41 <= int32(0) {
		v79 = v37
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v48 = v37
	goto L24
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	F_pfree(m, v45)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L29
	}
L24:
	;
	v58 = v45 + v48<<(uint(int32(3))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	F_pfree(m, v45)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	v63 = v48 + int32(1)
	if v63 != v41 {
		v48 = v63
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v79 = int32(0)
	goto L14
L29:
	;
	v79 = v68
	goto L14
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v79 = v76
	goto L14
L31:
	;
	if v100 == int32(0) {
		v127 = v79
		goto L5
	} else {
		goto L35
	}
L32:
	;
	v100 = base.B2i32(base.Ui32(v79) < base.Ui32(v88))
	goto L31
L33:
	;
	goto L34
L34:
	;
	v100 = int32(base.Ui32(v79-v88) >> (uint(int32(31)) % 32))
	goto L31
L35:
	;
	goto L1
L36:
	;
	v127 = int32(2)
	goto L5
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
	F_errmsg_internal(m, int32(63650), v13)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(523527), int32(9310), int32(73777))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v138 = F_SubTransGetTopmostTransaction(m, v127)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v141))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v153 != 0 {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v153 = base.B2i32(base.Ui32(v138) < base.Ui32(v141))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v153 = int32(base.Ui32(v138-v141) >> (uint(int32(31)) % 32))
	goto L42
L46:
	;
	v154 = m.G0
	v156 = v154 - int32(32)
	m.G0 = v156
	v159 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v159 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L1
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L2
	} else {
		goto L177
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L2
	} else {
		goto L171
	}
L50:
	;
	m.G0 = v156 + int32(32)
	goto L47
L51:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	switch v162 {
	case 0, 5:
		goto L52
	default:
		goto L50
	}
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)+108))
	if v163&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui32(v170) < base.Ui32(int32(12000)) {
		goto L50
	} else {
		goto L57
	}
L56:
	;
	goto L50
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+118)))
	if v174 == int32(116) {
		goto L50
	} else {
		goto L58
	}
L58:
	;
	if v163&int32(8) != 0 {
		goto L49
	} else {
		goto L59
	}
L59:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v138 == v180 {
		goto L50
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+28)) = v138
	v184 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v188 = F_LWLockAcquire(m, v184+int32(3584), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v194 = int32(0)
	v196 = F_hash_search(m, v191, v156+int32(28), v194, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v196 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v205 = F_LWLockAcquire(m, v201+int32(6656), int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v338 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v336 != v338 {
		goto L105
	} else {
		goto L106
	}
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v212 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v212+int32(6656))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	if v210 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v331+int32(3584))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L103
	}
L69:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v209))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v230 != 0 {
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v230 = base.B2i32(base.Ui32(v138) < base.Ui32(v209))
	goto L70
L72:
	;
	goto L73
L73:
	;
	v230 = int32(base.Ui32(v138-v209) >> (uint(int32(31)) % 32))
	goto L70
L74:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v210))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v242 != 0 {
		goto L68
	} else {
		goto L79
	}
L76:
	;
	v242 = base.B2i32(base.Ui32(v210) < base.Ui32(v138))
	goto L75
L77:
	;
	goto L78
L78:
	;
	v242 = base.B2i32(int32(0) < v138-v210)
	goto L75
L79:
	;
	v245 = int32(base.Ui32(v138) >> (uint(int32(10)) % 32))
	v247 = F_SimpleLruReadPage_ReadOnly(m, int32(4485364), base.I64_extend_i32_u(v245), v138)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v247<<(uint(int32(2))%32))))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v255+v138<<(uint(int32(3))%32)&int32(8184))))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, _consts[90])))
	v265 = base.I32_rem_u_s(v245, v264)
	F_LWLockRelease(m, v262+v265<<(uint(int32(7))%32))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v271 = int64(1)
	v272 = v261 + v271
	if base.Ui64(v272) <= base.Ui64(v271) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v312&int32(512) != 0 {
		goto L48
	} else {
		goto L98
	}
L83:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+108))
	v312 = v311
	v313 = v310
	goto L82
L84:
	;
	if base.I32_wrap_i64(v272)-int32(1) != 0 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+108))
	if v280&int32(32) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L68
L88:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v279)+24))
	if base.Ui64(v283) < base.Ui64(v261) {
		v312 = v280
		v313 = v279
		goto L82
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(151408), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v138
	F_errdetail_internal(m, int32(607136), v156)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errhint(m, int32(682086), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(524956), int32(4072), int32(73781))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	if v316 != v313+int32(40) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v321 = v316
	goto L101
L100:
	;
	v321 = int32(0)
	goto L101
L101:
	;
	if v321 != 0 {
		goto L48
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+108)) = v312 | int32(1024)
	goto L68
L103:
	;
	goto L50
L104:
	;
	if v340&int32(1024) != 0 {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336)+108))
	if v340&int32(8) == int32(0) {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v347+int32(3584))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	goto L50
L110:
	;
	if v340&int32(2) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v340&int32(1) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336)+108)) = v340 | int32(8)
	v362 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v362+int32(3584))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v368+int32(3584))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L117
	}
L116:
	;
	goto L50
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(151408), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errdetail_internal(m, int32(609181), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errhint(m, int32(682086), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(524956), int32(4119), int32(73781))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	v417 = int32(0)
	v419 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L2
	} else {
		goto L133
	}
L125:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v338)+108))
	if v401&int32(32) == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v340&int32(16) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v338)+24))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v336)+24))
	if base.Ui64(v409) <= base.Ui64(v408) {
		goto L124
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v412+int32(3584))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	goto L50
L132:
	;
	if v491 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L133:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v421))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v433 != 0 {
		v491 = v417
		goto L132
	} else {
		goto L138
	}
L135:
	;
	v433 = base.B2i32(base.Ui32(v138) < base.Ui32(v421))
	goto L134
L136:
	;
	goto L137
L137:
	;
	v433 = int32(base.Ui32(v138-v421) >> (uint(int32(31)) % 32))
	goto L134
L138:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v435))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v447 != 0 {
		v491 = int32(1)
		goto L132
	} else {
		goto L143
	}
L140:
	;
	v447 = base.B2i32(base.Ui32(v435) <= base.Ui32(v138))
	goto L139
L141:
	;
	goto L142
L142:
	;
	v447 = base.B2i32(int32(0) <= v138-v435)
	goto L139
L143:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	if v448 == int32(0) {
		v472 = v417
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v491 = v472
	goto L132
L145:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v453 = int32(0)
	goto L146
L146:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v451+v453<<(uint(int32(2))%32))))
	v467 = base.B2i32(v138 == v466)
	if v138 == v466 {
		v472 = v467
		goto L144
	} else {
		goto L148
	}
L147:
	;
	v472 = v467
	goto L144
L148:
	;
	v469 = v453 + int32(1)
	if v469 != v448 {
		v453 = v469
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v495+int32(3584))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L2
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v500 = int32(0)
	v502 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+108)))
	if v503&int32(8) != 0 {
		v540 = v500
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L50
L154:
	;
	if v540 != 0 {
		goto L165
	} else {
		goto L166
	}
L155:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+108)))
	if v506&int32(8) != 0 {
		v540 = v500
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v502)+36))
	if v509 == int32(0) {
		v540 = v500
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v513 = v502 + int32(32)
	if v509 == v513 {
		v540 = v500
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v336)+44))
	if v515 == int32(0) {
		v540 = v500
		goto L154
	} else {
		goto L159
	}
L159:
	;
	if v515 == v336+int32(40) {
		v540 = v500
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v522 = v509
	goto L161
L161:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v522)+20))
	v532 = base.B2i32(v531 == v336)
	if v531 == v336 {
		v540 = v532
		goto L154
	} else {
		goto L163
	}
L162:
	;
	v540 = v532
	goto L154
L163:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v533 != v513 {
		v522 = v533
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v546+int32(3584))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L2
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	F_FlagRWConflict(m, v502, v336)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L169
	}
L168:
	;
	goto L50
L169:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v554+int32(3584))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	goto L50
L171:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L2
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(151408), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	F_errdetail_internal(m, int32(658541), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	F_errhint(m, int32(682086), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(524956), int32(4039), int32(73781))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L2
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	F_errmsg(m, int32(151408), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v138
	F_errdetail_internal(m, int32(607373), v156+int32(16))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	F_errhint(m, int32(682086), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(524956), int32(4080), int32(73781))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_HeapTupleHeaderAdvanceConflictHorizon(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v15 = int32(768)
	if v14&v15 != v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v19
	goto L3
L2:
	;
	v20 = int32(2)
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14&int32(6272) != int32(4096) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if base.Ui32(v69&int32(65535)) < base.Ui32(int32(16384)) {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	v69 = v14
	v70 = v21
	goto L4
L6:
	;
	goto L7
L7:
	;
	v29 = F_GetMultiXactIdMembers(m, v21, v11+int32(12), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if int32(0) < v29 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v37 = int32(0)
	goto L15
L11:
	;
	v61 = int32(0)
	goto L12
L12:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v69 = v66
	v70 = v61
	goto L4
L13:
	;
	F_pfree(m, v34)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L19
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v55 = v53
	goto L13
L15:
	;
	v45 = v34 + v37<<(uint(int32(3))%32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v55 = int32(0)
	goto L13
L17:
	;
	v50 = v37 + int32(1)
	if v50 != v29 {
		v37 = v50
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v61 = v55
	goto L12
L20:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v97&int32(256) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v80))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v79)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v92 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v92 = base.B2i32(base.Ui32(v79) < base.Ui32(v80))
	goto L22
L24:
	;
	goto L25
L25:
	;
	v92 = int32(base.Ui32(v79-v80) >> (uint(int32(31)) % 32))
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	goto L20
L27:
	;
	m.G0 = v11 + int32(16)
	return
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v110))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v70)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	if v97&int32(512) != 0 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v70 == v20 {
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v104 = F_TransactionIdDidCommit(m, v20)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	if v104 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v70 != v20 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L28
L37:
	;
	if v122 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L38:
	;
	v122 = base.B2i32(base.Ui32(v110) < base.Ui32(v70))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v122 = base.B2i32(int32(0) < v70-v110)
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v70
	goto L27
}
func F_HeapTupleSatisfiesUpdate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v819 int32
	_ = v819
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v8&int32(256) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L51
	} else {
		goto L462
	}
L2:
	;
	return v1249
L3:
	;
	v774 = int32(0)
	v775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v775&int32(2048) != 0 {
		v1249 = v774
		goto L2
	} else {
		goto L287
	}
L4:
	;
	v11 = int32(1)
	v12 = base.I32_extend16_s(v8)
	if v12&int32(512) != 0 {
		v1249 = v11
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v12&int32(16384) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v767 = v144 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v767)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L51
	} else {
		goto L286
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if base.Ui32(v17) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	if v12 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L10:
	;
	if v137 != 0 {
		v1249 = v11
		goto L2
	} else {
		goto L50
	}
L11:
	;
	v137 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v28 == v17 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v137 = int32(1)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v32 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v137 = v129
	goto L10
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v36 == int32(0) {
		v129 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v100 = int32(0)
	v102 = v32 - int32(1)
	goto L40
L21:
	;
	v41 = v36
	goto L22
L22:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v46 == int32(4) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v129 = int32(0)
	goto L17
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	if v93 != 0 {
		v41 = v93
		goto L22
	} else {
		goto L39
	}
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v49 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v52 = int32(1)
	if v17 == v49 {
		v129 = v52
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v56 = v54 - int32(1)
	if v56 < int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v61 = int32(0)
	v63 = v56
	goto L29
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v69 = int32(2)
	v70 = base.I32_div_s(v63-v61, v69)
	v71 = v70 + v61
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v67+v71<<(uint(v69)%32))))
	if v75 == v17 {
		v129 = v52
		goto L17
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v79 = F_TransactionIdPrecedes(m, v75, v17)
	mBase = m.M
	if v79 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = v71 + int32(1)
	goto L34
L33:
	;
	v80 = v61
	goto L34
L34:
	;
	if v79 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v83 = v63
	goto L37
L36:
	;
	v83 = v71 - int32(1)
	goto L37
L37:
	;
	if v80 <= v83 {
		v61 = v80
		v63 = v83
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	goto L23
L40:
	;
	v107 = int32(2)
	v108 = base.I32_div_s(v102-v100, v107)
	v109 = v108 + v100
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v98+v109<<(uint(v107)%32))))
	v114 = base.B2i32(v113 == v17)
	if v113 == v17 {
		v129 = v114
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v129 = v114
	goto L17
L42:
	;
	v117 = base.B2i32(base.Ui32(v113) < base.Ui32(v17))
	if base.Ui32(v113) < base.Ui32(v17) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v118 = v109 + int32(1)
	goto L45
L44:
	;
	v118 = v100
	goto L45
L45:
	;
	if base.Ui32(v113) < base.Ui32(v17) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v121 = v102
	goto L48
L47:
	;
	v121 = v109 - int32(1)
	goto L48
L48:
	;
	if v118 <= v121 {
		v100 = v118
		v102 = v121
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v138 = F_TransactionIdIsInProgress(m, v17)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return int32(0)
L52:
	;
	if v138 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v142 = F_TransactionIdDidCommit(m, v17)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v142 == int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v148 = v144 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v148)
	goto L1
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if base.Ui32(v152) < base.Ui32(int32(3)) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if base.Ui32(v287) < base.Ui32(int32(3)) {
		goto L108
	} else {
		goto L109
	}
L59:
	;
	if v272 != 0 {
		goto L3
	} else {
		goto L99
	}
L60:
	;
	v272 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v163 == v152 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v272 = int32(1)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v167 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v272 = v264
	goto L59
L67:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v171 == int32(0) {
		v264 = int32(0)
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v235 = int32(0)
	v237 = v167 - int32(1)
	goto L89
L70:
	;
	v176 = v171
	goto L71
L71:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	if v181 == int32(4) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v264 = int32(0)
	goto L66
L73:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v176)+80))
	if v228 != 0 {
		v176 = v228
		goto L71
	} else {
		goto L88
	}
L74:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v184 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v187 = int32(1)
	if v152 == v184 {
		v264 = v187
		goto L66
	} else {
		goto L76
	}
L76:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	v191 = v189 - int32(1)
	if v191 < int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v196 = int32(0)
	v198 = v191
	goto L78
L78:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	v204 = int32(2)
	v205 = base.I32_div_s(v198-v196, v204)
	v206 = v205 + v196
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(v204)%32))))
	if v210 == v152 {
		v264 = v187
		goto L66
	} else {
		goto L80
	}
L79:
	;
	goto L73
L80:
	;
	v214 = F_TransactionIdPrecedes(m, v210, v152)
	mBase = m.M
	if v214 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v215 = v206 + int32(1)
	goto L83
L82:
	;
	v215 = v196
	goto L83
L83:
	;
	if v214 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v218 = v198
	goto L86
L85:
	;
	v218 = v206 - int32(1)
	goto L86
L86:
	;
	if v215 <= v218 {
		v196 = v215
		v198 = v218
		goto L78
	} else {
		goto L87
	}
L87:
	;
	goto L79
L88:
	;
	goto L72
L89:
	;
	v242 = int32(2)
	v243 = base.I32_div_s(v237-v235, v242)
	v244 = v243 + v235
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(v242)%32))))
	v249 = base.B2i32(v248 == v152)
	if v248 == v152 {
		v264 = v249
		goto L66
	} else {
		goto L91
	}
L90:
	;
	v264 = v249
	goto L66
L91:
	;
	v252 = base.B2i32(base.Ui32(v248) < base.Ui32(v152))
	if base.Ui32(v248) < base.Ui32(v152) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v253 = v244 + int32(1)
	goto L94
L93:
	;
	v253 = v235
	goto L94
L94:
	;
	if base.Ui32(v248) < base.Ui32(v152) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v256 = v237
	goto L97
L96:
	;
	v256 = v244 - int32(1)
	goto L97
L97:
	;
	if v253 <= v256 {
		v235 = v253
		v237 = v256
		goto L89
	} else {
		goto L98
	}
L98:
	;
	goto L90
L99:
	;
	v273 = F_TransactionIdIsInProgress(m, v152)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L51
	} else {
		goto L100
	}
L100:
	;
	if v273 != 0 {
		v1249 = v11
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v275 = F_TransactionIdDidCommit(m, v152)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L51
	} else {
		goto L102
	}
L102:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v275 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v279 = v277 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v279)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L51
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v285 = v277 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v285)
	goto L1
L106:
	;
	goto L3
L107:
	;
	if v407 != 0 {
		goto L147
	} else {
		goto L148
	}
L108:
	;
	v407 = int32(0)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v298 == v287 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v407 = int32(1)
	goto L107
L112:
	;
	goto L113
L113:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v302 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v407 = v399
	goto L107
L115:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v306 == int32(0) {
		v399 = int32(0)
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v370 = int32(0)
	v372 = v302 - int32(1)
	goto L137
L118:
	;
	v311 = v306
	goto L119
L119:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	if v316 == int32(4) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v399 = int32(0)
	goto L114
L121:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v311)+80))
	if v363 != 0 {
		v311 = v363
		goto L119
	} else {
		goto L136
	}
L122:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v319 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v322 = int32(1)
	if v287 == v319 {
		v399 = v322
		goto L114
	} else {
		goto L124
	}
L124:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v311)+52))
	v326 = v324 - int32(1)
	if v326 < int32(0) {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v331 = int32(0)
	v333 = v326
	goto L126
L126:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v311)+48))
	v339 = int32(2)
	v340 = base.I32_div_s(v333-v331, v339)
	v341 = v340 + v331
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337+v341<<(uint(v339)%32))))
	if v345 == v287 {
		v399 = v322
		goto L114
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v349 = F_TransactionIdPrecedes(m, v345, v287)
	mBase = m.M
	if v349 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v350 = v341 + int32(1)
	goto L131
L130:
	;
	v350 = v331
	goto L131
L131:
	;
	if v349 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v353 = v333
	goto L134
L133:
	;
	v353 = v341 - int32(1)
	goto L134
L134:
	;
	if v350 <= v353 {
		v331 = v350
		v333 = v353
		goto L126
	} else {
		goto L135
	}
L135:
	;
	goto L127
L136:
	;
	goto L120
L137:
	;
	v377 = int32(2)
	v378 = base.I32_div_s(v372-v370, v377)
	v379 = v378 + v370
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v368+v379<<(uint(v377)%32))))
	v384 = base.B2i32(v383 == v287)
	if v383 == v287 {
		v399 = v384
		goto L114
	} else {
		goto L139
	}
L138:
	;
	v399 = v384
	goto L114
L139:
	;
	v387 = base.B2i32(base.Ui32(v383) < base.Ui32(v287))
	if base.Ui32(v383) < base.Ui32(v287) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v388 = v379 + int32(1)
	goto L142
L141:
	;
	v388 = v370
	goto L142
L142:
	;
	if base.Ui32(v383) < base.Ui32(v287) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v391 = v372
	goto L145
L144:
	;
	v391 = v379 - int32(1)
	goto L145
L145:
	;
	if v388 <= v391 {
		v370 = v388
		v372 = v391
		goto L137
	} else {
		goto L146
	}
L146:
	;
	goto L138
L147:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v410&int32(32) != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	goto L149
L149:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v753 = F_TransactionIdIsInProgress(m, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L51
	} else {
		goto L279
	}
L150:
	;
	if base.Ui32(l1) <= base.Ui32(v419) {
		v1249 = v11
		goto L2
	} else {
		goto L154
	}
L151:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v409<<(uint(int32(3))%32))))
	v419 = v418
	goto L153
L152:
	;
	v419 = v409
	goto L153
L153:
	;
	goto L150
L154:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v422&int32(2048) != 0 {
		v1249 = int32(0)
		goto L2
	} else {
		goto L155
	}
L155:
	;
	v427 = int32(0)
	if base.B2i32(v422&int32(128) == v427)&base.B2i32(v422&int32(4176) != int32(64)) == v427 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v422&int32(4096) != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if v422&int32(4096) != 0 {
		goto L170
	} else {
		goto L171
	}
L159:
	;
	v442 = F_MultiXactIdIsRunning(m, v436, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L51
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v448 = F_TransactionIdIsInProgress(m, v436)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L51
	} else {
		goto L166
	}
L162:
	;
	if v442 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v444 = int32(5)
	goto L165
L164:
	;
	v444 = int32(0)
	goto L165
L165:
	;
	return v444
L166:
	;
	if v448 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v450 = int32(5)
	goto L169
L168:
	;
	v450 = int32(0)
	goto L169
L169:
	;
	return v450
L170:
	;
	v454 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L51
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v603) < base.Ui32(int32(3)) {
		goto L229
	} else {
		goto L230
	}
L173:
	;
	if base.Ui32(v454) < base.Ui32(int32(3)) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v575 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L175:
	;
	v575 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v466 == v454 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v575 = int32(1)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v470 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v575 = v567
	goto L174
L182:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v474 == int32(0) {
		v567 = int32(0)
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v538 = int32(0)
	v540 = v470 - int32(1)
	goto L204
L185:
	;
	v479 = v474
	goto L186
L186:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	if v484 == int32(4) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v567 = int32(0)
	goto L181
L188:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v479)+80))
	if v531 != 0 {
		v479 = v531
		goto L186
	} else {
		goto L203
	}
L189:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v487 == int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v490 = int32(1)
	if v454 == v487 {
		v567 = v490
		goto L181
	} else {
		goto L191
	}
L191:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v479)+52))
	v494 = v492 - int32(1)
	if v494 < int32(0) {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v499 = int32(0)
	v501 = v494
	goto L193
L193:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v507 = int32(2)
	v508 = base.I32_div_s(v501-v499, v507)
	v509 = v508 + v499
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v505+v509<<(uint(v507)%32))))
	if v513 == v454 {
		v567 = v490
		goto L181
	} else {
		goto L195
	}
L194:
	;
	goto L188
L195:
	;
	v517 = F_TransactionIdPrecedes(m, v513, v454)
	mBase = m.M
	if v517 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v518 = v509 + int32(1)
	goto L198
L197:
	;
	v518 = v499
	goto L198
L198:
	;
	if v517 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v521 = v501
	goto L201
L200:
	;
	v521 = v509 - int32(1)
	goto L201
L201:
	;
	if v518 <= v521 {
		v499 = v518
		v501 = v521
		goto L193
	} else {
		goto L202
	}
L202:
	;
	goto L194
L203:
	;
	goto L187
L204:
	;
	v545 = int32(2)
	v546 = base.I32_div_s(v540-v538, v545)
	v547 = v546 + v538
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v536+v547<<(uint(v545)%32))))
	v552 = base.B2i32(v551 == v454)
	if v551 == v454 {
		v567 = v552
		goto L181
	} else {
		goto L206
	}
L205:
	;
	v567 = v552
	goto L181
L206:
	;
	v555 = base.B2i32(base.Ui32(v551) < base.Ui32(v454))
	if base.Ui32(v551) < base.Ui32(v454) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v556 = v547 + int32(1)
	goto L209
L208:
	;
	v556 = v538
	goto L209
L209:
	;
	if base.Ui32(v551) < base.Ui32(v454) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v559 = v540
	goto L212
L211:
	;
	v559 = v547 - int32(1)
	goto L212
L212:
	;
	if v556 <= v559 {
		v538 = v556
		v540 = v559
		goto L204
	} else {
		goto L213
	}
L213:
	;
	goto L205
L214:
	;
	v579 = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v582 = F_MultiXactIdIsRunning(m, v580, v579)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L51
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v590&int32(32) != 0 {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	if v582 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v584 = int32(5)
	goto L220
L219:
	;
	v584 = v579
	goto L220
L220:
	;
	return v584
L221:
	;
	if base.Ui32(v599) < base.Ui32(l1) {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594+v589<<(uint(int32(3))%32))+4))
	v599 = v598
	goto L224
L223:
	;
	v599 = v589
	goto L224
L224:
	;
	goto L221
L225:
	;
	v601 = int32(1)
	goto L227
L226:
	;
	v601 = int32(2)
	goto L227
L227:
	;
	return v601
L228:
	;
	if v723 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L229:
	;
	v723 = int32(0)
	goto L228
L230:
	;
	goto L231
L231:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v614 == v603 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v723 = int32(1)
	goto L228
L233:
	;
	goto L234
L234:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v618 <= int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v723 = v715
	goto L228
L236:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v622 == int32(0) {
		v715 = int32(0)
		goto L235
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v684 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v686 = int32(0)
	v688 = v618 - int32(1)
	goto L258
L239:
	;
	v627 = v622
	goto L240
L240:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v627)+20))
	if v632 == int32(4) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v715 = int32(0)
	goto L235
L242:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v627)+80))
	if v679 != 0 {
		v627 = v679
		goto L240
	} else {
		goto L257
	}
L243:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	if v635 == int32(0) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v638 = int32(1)
	if v603 == v635 {
		v715 = v638
		goto L235
	} else {
		goto L245
	}
L245:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v627)+52))
	v642 = v640 - int32(1)
	if v642 < int32(0) {
		goto L242
	} else {
		goto L246
	}
L246:
	;
	v647 = int32(0)
	v649 = v642
	goto L247
L247:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v627)+48))
	v655 = int32(2)
	v656 = base.I32_div_s(v649-v647, v655)
	v657 = v656 + v647
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v653+v657<<(uint(v655)%32))))
	if v661 == v603 {
		v715 = v638
		goto L235
	} else {
		goto L249
	}
L248:
	;
	goto L242
L249:
	;
	v665 = F_TransactionIdPrecedes(m, v661, v603)
	mBase = m.M
	if v665 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v666 = v657 + int32(1)
	goto L252
L251:
	;
	v666 = v647
	goto L252
L252:
	;
	if v665 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v669 = v649
	goto L255
L254:
	;
	v669 = v657 - int32(1)
	goto L255
L255:
	;
	if v666 <= v669 {
		v647 = v666
		v649 = v669
		goto L247
	} else {
		goto L256
	}
L256:
	;
	goto L248
L257:
	;
	goto L241
L258:
	;
	v693 = int32(2)
	v694 = base.I32_div_s(v688-v686, v693)
	v695 = v694 + v686
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v684+v695<<(uint(v693)%32))))
	v700 = base.B2i32(v699 == v603)
	if v699 == v603 {
		v715 = v700
		goto L235
	} else {
		goto L260
	}
L259:
	;
	v715 = v700
	goto L235
L260:
	;
	v703 = base.B2i32(base.Ui32(v699) < base.Ui32(v603))
	if base.Ui32(v699) < base.Ui32(v603) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v704 = v695 + int32(1)
	goto L263
L262:
	;
	v704 = v686
	goto L263
L263:
	;
	if base.Ui32(v699) < base.Ui32(v603) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v707 = v688
	goto L266
L265:
	;
	v707 = v695 - int32(1)
	goto L266
L266:
	;
	if v704 <= v707 {
		v686 = v704
		v688 = v707
		goto L258
	} else {
		goto L267
	}
L267:
	;
	goto L259
L268:
	;
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v728 = v726 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v728)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L51
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v739&int32(32) != 0 {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	return int32(0)
L272:
	;
	if base.Ui32(v748) < base.Ui32(l1) {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743+v738<<(uint(int32(3))%32))+4))
	v748 = v747
	goto L275
L274:
	;
	v748 = v738
	goto L275
L275:
	;
	goto L272
L276:
	;
	v750 = int32(1)
	goto L278
L277:
	;
	v750 = int32(2)
	goto L278
L278:
	;
	return v750
L279:
	;
	if v753 != 0 {
		v1249 = v11
		goto L2
	} else {
		goto L280
	}
L280:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v756 = F_TransactionIdDidCommit(m, v755)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L51
	} else {
		goto L281
	}
L281:
	;
	if v756 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	F_HeapTupleSetHintBits(m, v7, l2, int32(256), v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L51
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v764 = v762 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v764)
	goto L1
L285:
	;
	goto L3
L286:
	;
	goto L3
L287:
	;
	if v775&int32(1024) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	if v775&int32(128) != 0 {
		v1249 = v774
		goto L2
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	if v775&int32(4096) != 0 {
		goto L302
	} else {
		goto L303
	}
L291:
	;
	if v775&int32(4176) == int32(64) {
		v1249 = v774
		goto L2
	} else {
		goto L292
	}
L292:
	;
	v786 = int32(4)
	v789 = l0 + v786
	v791 = v7 + int32(12)
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v789)+2)))
	v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v789))))
	v794 = int32(16)
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791)+2)))
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791))))
	if v792|v793<<(uint(v794)%32) == v797|v798<<(uint(v794)%32) {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	if v808 != 0 {
		goto L299
	} else {
		goto L300
	}
L294:
	;
	goto L293
L295:
	;
	v804 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v789)+4)))
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791)+4)))
	if v804 == v805 {
		v808 = int32(1)
		goto L294
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v808 = int32(0)
	goto L294
L298:
	;
	goto L297
L299:
	;
	v809 = v786
	goto L301
L300:
	;
	v809 = int32(3)
	goto L301
L301:
	;
	return v809
L302:
	;
	if v775&int32(4304) == int32(4224) {
		v1249 = v774
		goto L2
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v1035) < base.Ui32(int32(3)) {
		goto L388
	} else {
		goto L389
	}
L305:
	;
	v819 = int32(0)
	if base.B2i32(v775&int32(128) == v819)&base.B2i32(v775&int32(4176) != int32(64)) == v819 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v830 = F_MultiXactIdIsRunning(m, v828, int32(1))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L51
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v843 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L51
	} else {
		goto L315
	}
L309:
	;
	if v830 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	return int32(5)
L311:
	;
	goto L312
L312:
	;
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v836 = v834 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v836)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L51
	} else {
		goto L313
	}
L313:
	;
	return int32(0)
L314:
	;
	if base.Ui32(v843) < base.Ui32(int32(3)) {
		goto L320
	} else {
		goto L321
	}
L315:
	;
	if v843 != 0 {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v847 = F_MultiXactIdIsRunning(m, v845, int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L51
	} else {
		goto L317
	}
L317:
	;
	if v847 == int32(0) {
		goto L314
	} else {
		goto L318
	}
L318:
	;
	return int32(5)
L319:
	;
	if v972 != 0 {
		goto L359
	} else {
		goto L360
	}
L320:
	;
	v972 = int32(0)
	goto L319
L321:
	;
	goto L322
L322:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v863 == v843 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v972 = int32(1)
	goto L319
L324:
	;
	goto L325
L325:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v867 <= int32(0) {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	v972 = v964
	goto L319
L327:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v871 == int32(0) {
		v964 = int32(0)
		goto L326
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v935 = int32(0)
	v937 = v867 - int32(1)
	goto L349
L330:
	;
	v876 = v871
	goto L331
L331:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v876)+20))
	if v881 == int32(4) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v964 = int32(0)
	goto L326
L333:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v876)+80))
	if v928 != 0 {
		v876 = v928
		goto L331
	} else {
		goto L348
	}
L334:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	if v884 == int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v887 = int32(1)
	if v843 == v884 {
		v964 = v887
		goto L326
	} else {
		goto L336
	}
L336:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v876)+52))
	v891 = v889 - int32(1)
	if v891 < int32(0) {
		goto L333
	} else {
		goto L337
	}
L337:
	;
	v896 = int32(0)
	v898 = v891
	goto L338
L338:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v876)+48))
	v904 = int32(2)
	v905 = base.I32_div_s(v898-v896, v904)
	v906 = v905 + v896
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v902+v906<<(uint(v904)%32))))
	if v910 == v843 {
		v964 = v887
		goto L326
	} else {
		goto L340
	}
L339:
	;
	goto L333
L340:
	;
	v914 = F_TransactionIdPrecedes(m, v910, v843)
	mBase = m.M
	if v914 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v915 = v906 + int32(1)
	goto L343
L342:
	;
	v915 = v896
	goto L343
L343:
	;
	if v914 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v918 = v898
	goto L346
L345:
	;
	v918 = v906 - int32(1)
	goto L346
L346:
	;
	if v915 <= v918 {
		v896 = v915
		v898 = v918
		goto L338
	} else {
		goto L347
	}
L347:
	;
	goto L339
L348:
	;
	goto L332
L349:
	;
	v942 = int32(2)
	v943 = base.I32_div_s(v937-v935, v942)
	v944 = v943 + v935
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v933+v944<<(uint(v942)%32))))
	v949 = base.B2i32(v948 == v843)
	if v948 == v843 {
		v964 = v949
		goto L326
	} else {
		goto L351
	}
L350:
	;
	v964 = v949
	goto L326
L351:
	;
	v952 = base.B2i32(base.Ui32(v948) < base.Ui32(v843))
	if base.Ui32(v948) < base.Ui32(v843) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v953 = v944 + int32(1)
	goto L354
L353:
	;
	v953 = v935
	goto L354
L354:
	;
	if base.Ui32(v948) < base.Ui32(v843) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v956 = v937
	goto L357
L356:
	;
	v956 = v944 - int32(1)
	goto L357
L357:
	;
	if v953 <= v956 {
		v935 = v953
		v937 = v956
		goto L349
	} else {
		goto L358
	}
L358:
	;
	goto L350
L359:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v977&int32(32) != 0 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	goto L361
L361:
	;
	v990 = int32(5)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v993 = F_MultiXactIdIsRunning(m, v991, int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L51
	} else {
		goto L369
	}
L362:
	;
	if base.Ui32(v986) < base.Ui32(l1) {
		goto L366
	} else {
		goto L367
	}
L363:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v981+v976<<(uint(int32(3))%32))+4))
	v986 = v985
	goto L365
L364:
	;
	v986 = v976
	goto L365
L365:
	;
	goto L362
L366:
	;
	v988 = int32(1)
	goto L368
L367:
	;
	v988 = int32(2)
	goto L368
L368:
	;
	return v988
L369:
	;
	if v993 != 0 {
		v1249 = v990
		goto L2
	} else {
		goto L370
	}
L370:
	;
	v995 = F_TransactionIdDidCommit(m, v843)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L51
	} else {
		goto L371
	}
L371:
	;
	if v995 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v997 = int32(4)
	v1000 = l0 + v997
	v1002 = v7 + int32(12)
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1000)+2)))
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1000))))
	v1005 = int32(16)
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002)+2)))
	v1009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002))))
	if v1003|v1004<<(uint(v1005)%32) == v1008|v1009<<(uint(v1005)%32) {
		goto L377
	} else {
		goto L378
	}
L373:
	;
	goto L374
L374:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v1024 = F_MultiXactIdIsRunning(m, v1022, int32(0))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L51
	} else {
		goto L384
	}
L375:
	;
	if v1019 != 0 {
		goto L381
	} else {
		goto L382
	}
L376:
	;
	goto L375
L377:
	;
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1000)+4)))
	v1016 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002)+4)))
	if v1015 == v1016 {
		v1019 = int32(1)
		goto L376
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1019 = int32(0)
	goto L376
L380:
	;
	goto L379
L381:
	;
	v1020 = v997
	goto L383
L382:
	;
	v1020 = int32(3)
	goto L383
L383:
	;
	return v1020
L384:
	;
	if v1024 != 0 {
		v1249 = v990
		goto L2
	} else {
		goto L385
	}
L385:
	;
	v1026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v1028 = v1026 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v1028)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L51
	} else {
		goto L386
	}
L386:
	;
	return int32(0)
L387:
	;
	if v1155 != 0 {
		goto L427
	} else {
		goto L428
	}
L388:
	;
	v1155 = int32(0)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v1046 == v1035 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1155 = int32(1)
	goto L387
L392:
	;
	goto L393
L393:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v1050 <= int32(0) {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	v1155 = v1147
	goto L387
L395:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v1054 == int32(0) {
		v1147 = int32(0)
		goto L394
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v1118 = int32(0)
	v1120 = v1050 - int32(1)
	goto L417
L398:
	;
	v1059 = v1054
	goto L399
L399:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+20))
	if v1064 == int32(4) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1147 = int32(0)
	goto L394
L401:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+80))
	if v1111 != 0 {
		v1059 = v1111
		goto L399
	} else {
		goto L416
	}
L402:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1059)))
	if v1067 == int32(0) {
		goto L401
	} else {
		goto L403
	}
L403:
	;
	v1070 = int32(1)
	if v1035 == v1067 {
		v1147 = v1070
		goto L394
	} else {
		goto L404
	}
L404:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+52))
	v1074 = v1072 - int32(1)
	if v1074 < int32(0) {
		goto L401
	} else {
		goto L405
	}
L405:
	;
	v1079 = int32(0)
	v1081 = v1074
	goto L406
L406:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+48))
	v1087 = int32(2)
	v1088 = base.I32_div_s(v1081-v1079, v1087)
	v1089 = v1088 + v1079
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1085+v1089<<(uint(v1087)%32))))
	if v1093 == v1035 {
		v1147 = v1070
		goto L394
	} else {
		goto L408
	}
L407:
	;
	goto L401
L408:
	;
	v1097 = F_TransactionIdPrecedes(m, v1093, v1035)
	mBase = m.M
	if v1097 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1098 = v1089 + int32(1)
	goto L411
L410:
	;
	v1098 = v1079
	goto L411
L411:
	;
	if v1097 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1101 = v1081
	goto L414
L413:
	;
	v1101 = v1089 - int32(1)
	goto L414
L414:
	;
	if v1098 <= v1101 {
		v1079 = v1098
		v1081 = v1101
		goto L406
	} else {
		goto L415
	}
L415:
	;
	goto L407
L416:
	;
	goto L400
L417:
	;
	v1125 = int32(2)
	v1126 = base.I32_div_s(v1120-v1118, v1125)
	v1127 = v1126 + v1118
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1116+v1127<<(uint(v1125)%32))))
	v1132 = base.B2i32(v1131 == v1035)
	if v1131 == v1035 {
		v1147 = v1132
		goto L394
	} else {
		goto L419
	}
L418:
	;
	v1147 = v1132
	goto L394
L419:
	;
	v1135 = base.B2i32(base.Ui32(v1131) < base.Ui32(v1035))
	if base.Ui32(v1131) < base.Ui32(v1035) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1136 = v1127 + int32(1)
	goto L422
L421:
	;
	v1136 = v1118
	goto L422
L422:
	;
	if base.Ui32(v1131) < base.Ui32(v1035) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1139 = v1120
	goto L425
L424:
	;
	v1139 = v1127 - int32(1)
	goto L425
L425:
	;
	if v1136 <= v1139 {
		v1118 = v1136
		v1120 = v1139
		goto L417
	} else {
		goto L426
	}
L426:
	;
	goto L418
L427:
	;
	v1156 = int32(5)
	v1157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1157&int32(128) != 0 {
		v1249 = v1156
		goto L2
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v1182 = F_TransactionIdIsInProgress(m, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L51
	} else {
		goto L439
	}
L430:
	;
	if v1157&int32(4176) == int32(64) {
		v1249 = v1156
		goto L2
	} else {
		goto L431
	}
L431:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1168&int32(32) != 0 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	if base.Ui32(v1177) < base.Ui32(l1) {
		goto L436
	} else {
		goto L437
	}
L433:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1172+v1167<<(uint(int32(3))%32))+4))
	v1177 = v1176
	goto L435
L434:
	;
	v1177 = v1167
	goto L435
L435:
	;
	goto L432
L436:
	;
	v1179 = int32(1)
	goto L438
L437:
	;
	v1179 = int32(2)
	goto L438
L438:
	;
	return v1179
L439:
	;
	if v1182 != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	return int32(5)
L441:
	;
	goto L442
L442:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v1187 = F_TransactionIdDidCommit(m, v1186)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L51
	} else {
		goto L443
	}
L443:
	;
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1187 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1193 = v1189 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v1193)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L51
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	v1202 = int32(0)
	if base.B2i32(v1189&int32(128) == v1202)&base.B2i32(v1189&int32(4176) != int32(64)) == v1202 {
		goto L448
	} else {
		goto L449
	}
L447:
	;
	return int32(0)
L448:
	;
	v1212 = v1189 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v1212)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L51
	} else {
		goto L451
	}
L449:
	;
	goto L450
L450:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	F_HeapTupleSetHintBits(m, v7, l2, int32(1024), v1220)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L51
	} else {
		goto L452
	}
L451:
	;
	return int32(0)
L452:
	;
	v1223 = int32(4)
	v1226 = l0 + v1223
	v1228 = v7 + int32(12)
	v1229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1226)+2)))
	v1230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1226))))
	v1231 = int32(16)
	v1234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228)+2)))
	v1235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228))))
	if v1229|v1230<<(uint(v1231)%32) == v1234|v1235<<(uint(v1231)%32) {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	if v1245 != 0 {
		goto L459
	} else {
		goto L460
	}
L454:
	;
	goto L453
L455:
	;
	v1241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1226)+4)))
	v1242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228)+4)))
	if v1241 == v1242 {
		v1245 = int32(1)
		goto L454
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	v1245 = int32(0)
	goto L454
L458:
	;
	goto L457
L459:
	;
	v1246 = v1223
	goto L461
L460:
	;
	v1246 = int32(3)
	goto L461
L461:
	;
	v1249 = v1246
	goto L2
L462:
	;
	return int32(1)
}
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	v18 = m.G0
	v20 = v18 - int32(160)
	m.G0 = v20
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = int64(0)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v28 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+136)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v20)+128)) = v59
	F_swap_relation_files(m, l0, l1, base.B2i32(l0 == int32(1259)), l3, l5, l6, l7, v20+int32(128))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v32 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v35 = int32(4556756)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v37 + v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v41 + v38
	*(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))+232)) = int64(5)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v49 + v38
	v55 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v55 - v38
	goto L2
L5:
	;
	return
L6:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = m.G0
	v71 = v69 - int32(16)
	m.G0 = v71
	v75 = int32(1)
	if l0 <= int32(3591) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	goto L9
L9:
	;
	if l4 != 0 {
		goto L51
	} else {
		goto L52
	}
L10:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v146 = F_PrepareInvalidationState(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L38
	}
L11:
	;
	goto L10
L12:
	;
	v143 = int32(0)
	goto L11
L13:
	;
	if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
		v143 = v75
		goto L11
	} else {
		goto L36
	}
L14:
	;
	if l0 <= int32(2670) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if l0 <= int32(5999) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v143 = v75
		goto L11
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L12
	default:
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v87 = l0 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v87) {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v143 = v75
	goto L11
L22:
	;
	if int32(1)<<(uint(v87)%32)&int32(226492515) == int32(0) {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v143 = v75
	goto L11
L24:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v143 = v75
		goto L11
	} else {
		goto L34
	}
L25:
	;
	v99 = l0 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v99) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	switch l0 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v143 = v75
		goto L11
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L12
	default:
		goto L30
	}
L28:
	;
	if int32(1)<<(uint(v99)%32)&int32(963) == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v143 = v75
	goto L11
L30:
	;
	if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
		v143 = v75
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v115 = l0 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v115) {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	if int32(1)<<(uint(v115)%32)&int32(49153) != 0 {
		v143 = v75
		goto L11
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v143 = v75
	goto L11
L36:
	;
	if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
		v143 = v75
		goto L11
	} else {
		goto L37
	}
L37:
	;
	goto L12
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v152 = *(*int32)(unsafe.Add(mBase, _consts[396]))
	if v152 <= v150 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v149 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v174 = v149
	goto L41
L41:
	;
	v178 = v174 + v150<<(uint(int32(4))%32)
	v179 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v179)
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+13)))
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+1)) = uint16(v181)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)) = uint8(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = l0
	if v143 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[396])) = v168
	*(*int32)(unsafe.Add(mBase, _consts[121])) = v169
	v174 = v169
	goto L41
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v160 = F_MemoryContextAlloc(m, v158, int32(512))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v166 = F_repalloc(m, v149, v152<<(uint(int32(5))%32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	v168 = int32(32)
	v169 = v160
	goto L42
L47:
	;
	v168 = v152 << (uint(int32(1)) % 32)
	v169 = v166
	goto L42
L48:
	;
	v187 = int32(0)
	goto L50
L49:
	;
	v187 = v145
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v189 + int32(1)
	m.G0 = v71 + int32(16)
	goto L9
L51:
	;
	v206 = int32(6)
	goto L53
L52:
	;
	v206 = int32(2)
	goto L53
L53:
	;
	switch l8 - int32(112) {
	case 0:
		goto L55
	default:
		v213 = v206
		goto L54
	case 5:
		goto L56
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v218 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v213 = v206 | int32(16)
	goto L54
L56:
	;
	v213 = v206 | int32(8)
	goto L54
L57:
	;
	v249 = int32(0)
	v253 = F_reindex_relation(m, v249, l0, v213, v20+int32(120))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L61
	}
L58:
	;
	goto L57
L59:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v222 != int32(1) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v225 = int32(4556756)
	v227 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v228 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v227 + v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v231 + v228
	*(*int64)(unsafe.Add(mBase, uint32(v218+int32(8))+232)) = int64(6)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v239 + v228
	v245 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v245 - v228
	goto L58
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v259 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if l0 == int32(1259) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	goto L62
L64:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v263 != int32(1) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v266 = int32(4556756)
	v268 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v269 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v268 + v269
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v272 + v269
	*(*int64)(unsafe.Add(mBase, uint32(v259+int32(8))+232)) = int64(7)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v280 + v269
	v286 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v286 - v269
	goto L63
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L5
	} else {
		goto L125
	}
L67:
	;
	v294 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v318 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+148)) = int32(1259)
	F_performDeletion(m, v20+int32(148), v318, int32(1))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L75
	}
L70:
	;
	v299 = F_SearchSysCacheCopy(m, int32(57), int32(1259), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	if v299 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299)+16))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+22)))
	v305 = v303 + v304
	*(*int32)(unsafe.Add(mBase, uint32(v305)+140)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v305)+136)) = l6
	F_CatalogTupleUpdate(m, v294, v299+int32(4), v299)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_sequence_close(m, v294, int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v20)+128))
	if v329 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v334 = v329
	v335 = v249
	goto L79
L77:
	;
	goto L78
L78:
	;
	if l3 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v347 = int32(0)
	v348 = m.G0
	v350 = v348 - int32(16)
	m.G0 = v350
	v353 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if v353 <= v347 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L78
L81:
	;
	v430 = v335 + int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v430<<(uint(int32(2))%32))))
	if v434 != 0 {
		v334 = v434
		v335 = v430
		goto L79
	} else {
		goto L93
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L90
	}
L83:
	;
	v357 = v347
	goto L84
L84:
	;
	v374 = v357 << (uint(int32(3)) % 32)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v374)+uint32(_consts[398])))
	if v377 != v334 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v353<<(uint(int32(3))%32))+uint32(_consts[399])))
	*(*int64)(unsafe.Add(mBase, uint32(v374)+uint32(_consts[398]))) = v386
	v388 = int32(4551420)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	*(*int32)(unsafe.Add(mBase, _consts[397])) = v390 - int32(1)
	m.G0 = v350 + int32(16)
	goto L81
L86:
	;
	v380 = v357 + int32(1)
	if v353 != v380 {
		v357 = v380
		goto L84
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	goto L82
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v334
	F_errmsg_internal(m, int32(49976), v350)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(521070), int32(454), int32(352474))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
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
	goto L80
L94:
	;
	v455 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if l2 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L97:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v455)+48))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+112))
	if v458 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v460 = F_toast_get_valid_index(m, v458, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_relation_close(m, v455, int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L5
	} else {
		goto L118
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l0
	v469 = F_pg_snprintf(m, v20+int32(48), int32(64), int32(40242), v20+int32(32))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v455)+48))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+112))
	F_RenameRelationInternal(m, v472, v20+int32(48), int32(1), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	v486 = F_pg_snprintf(m, v20+int32(48), int32(64), int32(28796), v20+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v490 = int32(1)
	F_RenameRelationInternal(m, v460, v20+int32(48), v490, v490)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v455)+48))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+112))
	v498 = m.G0
	v500 = v498 - int32(16)
	m.G0 = v500
	v504 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v508 = F_SearchSysCacheCopy(m, int32(57), v497, int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	if v508 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L5
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v525+v526)+132)) = int32(0)
	F_CatalogTupleUpdate(m, v504, v508+int32(4), v508)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L5
	} else {
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v497
	F_errmsg_internal(m, int32(50136), v500)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(520068), int32(4376), int32(368368))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_pfree(m, v508)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	F_sequence_close(m, v504, int32(3))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	m.G0 = v500 + int32(16)
	goto L100
L118:
	;
	goto L96
L119:
	;
	v557 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	m.G0 = v20 + int32(160)
	return
L122:
	;
	F_RelationClearMissing(m, v557)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_relation_close(m, v557, int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1259)
	F_errmsg_internal(m, int32(50136), v20)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(520989), int32(1543), int32(250634))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_compare_slots_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	if v11 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+144))
	v17 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+l1<<(uint(v17)%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16+l0<<(uint(v17)%32))))
	v31 = v4
	goto L6
L4:
	;
	return int32(-1)
L5:
	;
	v99 = int32(0)
	if v90 < v99 {
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	v38 = v35 + v31*int32(36)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+10)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+6)))
	if v40 < v39 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	F_slot_getsomeattrs_int(m, v24, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v47 = v39 - int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v48))))
	v52 = v47 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+6)))
	if v56 < v39 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	F_slot_getsomeattrs_int(m, v20, v39)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v47))))
	if v50&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v93 = v31 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	if v93 < v94 {
		v31 = v93
		goto L6
	} else {
		goto L33
	}
L18:
	;
	if v62&int32(1) != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v62&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v67 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	return int32(1)
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v74 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77+v52)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int32)(m, v55, v79, v38)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	return int32(1)
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v83 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v81 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v90 = v81
	goto L30
L30:
	;
	if v90 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v90 = int32(0) - v81
	goto L30
L32:
	;
	goto L17
L33:
	;
	goto L7
L34:
	;
	v103 = int32(1)
	goto L36
L35:
	;
	v103 = v99 - v90
	goto L36
L36:
	;
	return v103
}
func F_heap_copy_tuple_as_datum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
	if v7&int32(4) != 0 {
		v10 = F_toast_flatten_tuple_to_datum(m, v6, v5, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	} else {
		v15 = F_palloc(m, v5)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v18 != 0 {
				v19 = F__emscripten_memcpy_bulkmem(m, v15, v17, v18)
				mBase = m.M
				v20 = v19
			} else {
				v20 = v15
			}
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21 << (uint(int32(2)) % 32)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v27
			return v20
		}
	}
}
func F_heap_copytuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v2 = int32(0)
	if l0 == v2 {
		v32 = v2
		return v32
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v7 == int32(0) {
			v32 = v2
			return v32
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = F_palloc(m, v10+int32(24))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v19
				v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v21)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v25 = v13 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v23
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v29 != 0 {
					v30 = F__emscripten_memcpy_bulkmem(m, v25, v28, v29)
					mBase = m.M
				} else {
				}
				v32 = v13
				return v32
			}
		}
	}
}
func F_heap_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v901 int32
	_ = v901
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = l2
	v23 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v27
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)) = uint8(v27)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v38&int32(1) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v38 = int32(1)
	goto L6
L5:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+76)))
	v38 = v37
	goto L6
L6:
	;
	goto L3
L7:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v47 = v43 | v44<<(uint(int32(16))%32)
	v48 = F_ReadBuffer(m, l0, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L263
	}
L10:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+10)))
	if v68&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v48 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v48^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L10
L13:
	;
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v67 = v61 + v48<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L15:
	;
	F_visibilitymap_pin(m, l0, v47, v18+int32(-28))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_LockBuffer(m, v48, int32(2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v79 = l1 + int32(4)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v81
	v87 = v80<<(uint(int32(2))%32) + v67 + int32(20)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v67 + v88&int32(32767)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(base.Ui32(v93) >> (uint(int32(17)) % 32))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v99
	v102 = v18 + int32(-20)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v114 = v103
	v118 = int32(0)
	goto L25
L20:
	;
	m.G0 = v20 - int32(-64)
	return v901
L21:
	;
	if v48 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L22:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v452
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)) = uint16(v454)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451)+20)))
	if v457&int32(6272) != int32(4096) {
		goto L148
	} else {
		goto L149
	}
L23:
	;
	if l3 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L24:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+20)))
	if v398&int32(2048) != 0 {
		v432 = v396
		v436 = v395
		goto L23
	} else {
		goto L127
	}
L25:
	;
	if v114 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v367&int32(3072) != 0 {
		v395 = v365
		goto L24
	} else {
		goto L119
	}
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v139 = F_HeapTupleSatisfiesUpdate(m, v18+int32(-24), v138, v48)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+10)))
	if v121&int32(4) == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_LockBuffer(m, v48, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_visibilitymap_pin(m, l0, v47, v18+int32(-28))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_LockBuffer(m, v48, int32(2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	if v139 != int32(5) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v139 != int32(1) {
		v432 = v139
		v436 = v118
		goto L23
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if l4 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_UnlockReleaseBuffer(m, v48)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(404501), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(523527), int32(2853), int32(369904))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v446 = int32(5)
	v450 = int32(0)
	goto L22
L44:
	;
	goto L45
L45:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+20)))
	if v169&int32(4096) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)) = uint8(v172)
	v177 = F_DoesMultiXactIdConflict(m, v168, v169, int32(3), v18+int32(-48))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v168) < base.Ui32(int32(3)) {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	if v177 == int32(0) {
		v395 = v118
		goto L24
	} else {
		goto L50
	}
L50:
	;
	F_LockBuffer(m, v48, int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)))
	if (v184|v118)&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v197 = int32(0)
	v201 = F_Do_MultiXactIdWait(m, v168, int32(5), v169, v197, l0, v102, int32(2), v197, v197)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L57
	}
L53:
	;
	v195 = v184 ^ int32(1) | v118
	goto L52
L54:
	;
	goto L55
L55:
	;
	F_LockTuple(m, l0, v102, int32(8))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v195 = int32(1)
	goto L52
L57:
	;
	F_LockBuffer(m, v48, int32(2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v206 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+10)))
	if v209&int32(4) != 0 {
		v114 = v206
		v118 = v195
		goto L25
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+20)))
	if (v213^v169)&int32(4304) != 0 {
		v114 = v206
		v118 = v195
		goto L25
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v217 != v168 {
		v114 = v206
		v118 = v195
		goto L25
	} else {
		goto L64
	}
L64:
	;
	v395 = v195
	goto L24
L65:
	;
	if v338 != 0 {
		v395 = v118
		goto L24
	} else {
		goto L105
	}
L66:
	;
	v338 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v229 == v168 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v338 = int32(1)
	goto L65
L70:
	;
	goto L71
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v233 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v338 = v330
	goto L65
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v237 == int32(0) {
		v330 = int32(0)
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v301 = int32(0)
	v303 = v233 - int32(1)
	goto L95
L76:
	;
	v242 = v237
	goto L77
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)+20))
	if v247 == int32(4) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v330 = int32(0)
	goto L72
L79:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v242)+80))
	if v294 != 0 {
		v242 = v294
		goto L77
	} else {
		goto L94
	}
L80:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v250 == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v253 = int32(1)
	if v168 == v250 {
		v330 = v253
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v242)+52))
	v257 = v255 - int32(1)
	if v257 < int32(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v262 = int32(0)
	v264 = v257
	goto L84
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	v270 = int32(2)
	v271 = base.I32_div_s(v264-v262, v270)
	v272 = v271 + v262
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v268+v272<<(uint(v270)%32))))
	if v276 == v168 {
		v330 = v253
		goto L72
	} else {
		goto L86
	}
L85:
	;
	goto L79
L86:
	;
	v280 = F_TransactionIdPrecedes(m, v276, v168)
	mBase = m.M
	if v280 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v281 = v272 + int32(1)
	goto L89
L88:
	;
	v281 = v262
	goto L89
L89:
	;
	if v280 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v284 = v264
	goto L92
L91:
	;
	v284 = v272 - int32(1)
	goto L92
L92:
	;
	if v281 <= v284 {
		v262 = v281
		v264 = v284
		goto L84
	} else {
		goto L93
	}
L93:
	;
	goto L85
L94:
	;
	goto L78
L95:
	;
	v308 = int32(2)
	v309 = base.I32_div_s(v303-v301, v308)
	v310 = v309 + v301
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v299+v310<<(uint(v308)%32))))
	v315 = base.B2i32(v314 == v168)
	if v314 == v168 {
		v330 = v315
		goto L72
	} else {
		goto L97
	}
L96:
	;
	v330 = v315
	goto L72
L97:
	;
	v318 = base.B2i32(base.Ui32(v314) < base.Ui32(v168))
	if base.Ui32(v314) < base.Ui32(v168) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v319 = v310 + int32(1)
	goto L100
L99:
	;
	v319 = v301
	goto L100
L100:
	;
	if base.Ui32(v314) < base.Ui32(v168) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v322 = v303
	goto L103
L102:
	;
	v322 = v310 - int32(1)
	goto L103
L103:
	;
	if v319 <= v322 {
		v301 = v319
		v303 = v322
		goto L95
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	F_LockBuffer(m, v48, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v118&int32(1) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_LockTuple(m, l0, v102, int32(8))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_XactLockTableWait(m, v168, l0, v102, int32(2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	F_LockBuffer(m, v48, int32(2))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v355 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+10)))
	if v360&int32(4) != 0 {
		v114 = int32(0)
		v118 = int32(1)
		goto L25
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v365 = int32(1)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v366)+20)))
	if (v169^v367)&int32(4304) != 0 {
		v114 = v355
		v118 = v365
		goto L25
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v371 != v168 {
		v114 = v355
		v118 = v365
		goto L25
	} else {
		goto L118
	}
L118:
	;
	goto L26
L119:
	;
	if v367&int32(128) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_HeapTupleSetHintBits(m, v366, v48, int32(2048), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L126
	}
L121:
	;
	if v367&int32(4176) == int32(64) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v381 = F_TransactionIdDidCommit(m, v168)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v381 == int32(0) {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	F_HeapTupleSetHintBits(m, v366, v48, int32(1024), v168)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v395 = v365
	goto L24
L126:
	;
	v395 = v365
	goto L24
L127:
	;
	if v398&int32(128) != 0 {
		v432 = v396
		v436 = v395
		goto L23
	} else {
		goto L128
	}
L128:
	;
	if v398&int32(4176) == int32(64) {
		v432 = v396
		v436 = v395
		goto L23
	} else {
		goto L129
	}
L129:
	;
	v407 = F_HeapTupleHeaderIsOnlyLocked(m, v397)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v407 != 0 {
		v432 = v396
		v436 = v395
		goto L23
	} else {
		goto L131
	}
L131:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v413 = v411 + int32(12)
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+2)))
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v416 = int32(16)
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413)+2)))
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413))))
	if v414|v415<<(uint(v416)%32) == v419|v420<<(uint(v416)%32) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	if v430 != 0 {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	goto L132
L134:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413)+4)))
	if v426 == v427 {
		v430 = int32(1)
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v430 = int32(0)
	goto L133
L137:
	;
	goto L136
L138:
	;
	v431 = int32(4)
	goto L140
L139:
	;
	v431 = int32(3)
	goto L140
L140:
	;
	v446 = v431
	v450 = v395
	goto L22
L141:
	;
	if v432 == int32(0) {
		goto L21
	} else {
		goto L146
	}
L142:
	;
	if v432 != 0 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v441 = F_HeapTupleSatisfiesVisibility(m, v18+int32(-24), l3, v48)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	if v441 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	v446 = int32(3)
	v450 = v436
	goto L22
L146:
	;
	v446 = v432
	v450 = v436
	goto L22
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v512
	if v446 == int32(2) {
		goto L160
	} else {
		goto L161
	}
L148:
	;
	v512 = v456
	goto L147
L149:
	;
	goto L150
L150:
	;
	v462 = int32(0)
	v466 = F_GetMultiXactIdMembers(m, v456, v18+int32(-48), v462)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v466 <= int32(0) {
		v512 = v462
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v480 = v462
	goto L155
L153:
	;
	F_pfree(m, v470)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L159
	}
L154:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v500 = v498
	goto L153
L155:
	;
	v490 = v470 + v480<<(uint(int32(3))%32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v491) {
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v500 = int32(0)
	goto L153
L157:
	;
	v495 = v480 + int32(1)
	if v495 != v466 {
		v480 = v495
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v512 = v500
	goto L147
L160:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+20)))
	if v526&int32(32) != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v537 = int32(-1)
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v537
	F_UnlockReleaseBuffer(m, v48)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L167
	}
L163:
	;
	v537 = v535
	goto L162
L164:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530+v525<<(uint(int32(3))%32))+4))
	v535 = v534
	goto L166
L165:
	;
	v535 = v525
	goto L166
L166:
	;
	goto L163
L167:
	;
	if v450&int32(1) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_UnlockTuple(m, l0, v102, int32(8))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v546 == int32(0) {
		v901 = v446
		goto L20
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	F_ReleaseBuffer(m, v546)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v901 = v446
	goto L20
L174:
	;
	F_CheckForSerializableConflictIn(m, l0, l1, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L178
	}
L175:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v554+(v48^int32(-1))<<(uint(int32(6))%32))+16))
	v569 = v560
	goto L174
L176:
	;
	goto L177
L177:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v562+v48<<(uint(int32(6))%32)+int32(-64))+16))
	v569 = v568
	goto L174
L178:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	F_HeapTupleHeaderAdjustCmax(m, v572, v18+int32(-4), v18+int32(-37))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v584 = F_ExtractReplicaIdentity(m, l0, v18+int32(-24), int32(1), v18+int32(-38))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588)+20)))
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588)+18)))
	F_compute_new_xmax_infomask(m, v589, v590, v591, v23, int32(3), int32(1), v18+int32(-32), v18+int32(-34), v18+int32(-36))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v602 = int32(4556756)
	v604 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v604 + int32(1)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v608 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+10)))
	v626 = v624 & int32(4)
	if v626 != 0 {
		goto L192
	} else {
		goto L193
	}
L184:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v608))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v23)) == int32(0) {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = v23
	goto L183
L187:
	;
	if v620 == int32(0) {
		goto L183
	} else {
		goto L191
	}
L188:
	;
	v620 = base.B2i32(base.Ui32(v23) < base.Ui32(v608))
	goto L187
L189:
	;
	goto L190
L190:
	;
	v620 = int32(base.Ui32(v23-v608) >> (uint(int32(31)) % 32))
	goto L187
L191:
	;
	goto L186
L192:
	;
	v628 = v624 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+10)) = uint16(v628)
	if v48 < int32(0) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v653)+20)))
	v656 = v654 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v653)+20)) = uint16(v656)
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v653)+18)))
	v660 = v658 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(v653)+18)) = uint16(v660)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662)+20)))
	v664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)))
	v665 = v663 | v664
	*(*uint16)(unsafe.Add(mBase, uint32(v662)+20)) = uint16(v665)
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662)+18)))
	v668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	v669 = v667 | v668
	*(*uint16)(unsafe.Add(mBase, uint32(v662)+18)) = uint16(v669)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v671)+18)))
	v674 = v672 & int32(49151)
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+18)) = uint16(v674)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v676)+4)) = v677
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+27)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+8)) = v681
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v680)+20)))
	if v679 != 0 {
		goto L200
	} else {
		goto L201
	}
L195:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v651 = F_visibilitymap_clear(m, v648, v649, int32(3))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L199
	}
L196:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633+(v48^int32(-1))<<(uint(int32(6))%32))+16))
	v648 = v639
	goto L195
L197:
	;
	goto L198
L198:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v641+v48<<(uint(int32(6))%32)+int32(-64))+16))
	v648 = v647
	goto L195
L199:
	;
	goto L194
L200:
	;
	v688 = int32(32)
	goto L202
L201:
	;
	v688 = int32(0)
	goto L202
L202:
	;
	v689 = v683&int32(65503) | v688
	*(*uint16)(unsafe.Add(mBase, uint32(v680)+20)) = uint16(v689)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v691)+12)) = v692
	v694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v691)+16)) = uint16(v694)
	if l6 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v697 = int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v696)+16)) = uint16(v697)
	*(*int32)(unsafe.Add(mBase, uint32(v696)+12)) = int32(-1)
	goto L205
L204:
	;
	goto L205
L205:
	;
	F_MarkBufferDirty(m, v48)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+118)))
	if v705 != int32(112) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v844 = int32(4556756)
	v846 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v846 - int32(1)
	F_LockBuffer(m, v48, int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L244
	}
L208:
	;
	v709 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v709 <= int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v740 = int32(base.Ui32(v626) >> (uint(int32(2)) % 32))
	if l6 != 0 {
		goto L224
	} else {
		goto L225
	}
L210:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v712 != 0 {
		goto L207
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	if v709 == int32(1) {
		goto L209
	} else {
		goto L215
	}
L213:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v713 == int32(0) {
		goto L209
	} else {
		goto L214
	}
L214:
	;
	goto L207
L215:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L216
L216:
	;
	if base.B2i32(base.Ui32(v718) < base.Ui32(int32(12000))) == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v723 == int32(0) {
		goto L209
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	F_log_heap_new_cid(m, l0, v18+int32(-24))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L223
	}
L220:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+119)))
	switch v727 - int32(109) {
	case 0, 5:
		goto L221
	default:
		goto L209
	}
L221:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+104)))
	if v730 != int32(1) {
		goto L209
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	goto L209
L224:
	;
	v743 = v740 | int32(16)
	goto L226
L225:
	;
	v743 = v740
	goto L226
L226:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+23)) = uint8(v743)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v745)+18)))
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v745)+20)))
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+20)) = uint16(v748)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v677
	v755 = int32(1)
	v759 = int32(4)
	v774 = int32(base.Ui32(v746)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v747)>>(uint(v755)%32))&int32(8) | (int32(base.Ui32(v747)>>(uint(v759)%32))&v759 | (int32(base.Ui32(v747)>>(uint(int32(12))%32))&v755 | int32(base.Ui32(v747)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)) = uint8(v774)
	if v584 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v829 = int32(4457908)
	v831 = int32(*(*uint8)(unsafe.Add(mBase, _consts[94])))
	v832 = v831 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v832)
	goto L242
L228:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+130)))
	if v779 == int32(102) {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L239
	}
L231:
	;
	v782 = int32(2)
	goto L233
L232:
	;
	v782 = int32(4)
	goto L233
L233:
	;
	v783 = v743 | v782
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+23)) = uint8(v783)
	F_XLogBeginInsert(m)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_XLogRegisterData(m, v18+int32(-48), int32(8))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_XLogRegisterBuffer(m, int32(0), v48, int32(8))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v584)+16))
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+10)) = uint16(v797)
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v799)
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)) = uint8(v801)
	F_XLogRegisterData(m, v18+int32(-54), int32(5))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v584)+16))
	v809 = int32(23)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	F_XLogRegisterData(m, v808+v809, v811-v809)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	goto L227
L239:
	;
	F_XLogRegisterData(m, v18+int32(-48), int32(8))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_XLogRegisterBuffer(m, int32(0), v48, int32(8))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	goto L227
L242:
	;
	v836 = F_XLogInsert(m, int32(10), int32(16))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotr(v836, int64(32))
	goto L207
L244:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v853 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	F_ReleaseBuffer(m, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+119)))
	switch v857 - int32(109) {
	case 0, 5:
		goto L250
	default:
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	v871 = int32(0)
	F_CacheInvalidateHeapTuple(m, l0, v18+int32(-24), v871)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L253
	}
L250:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+20)))
	if v861&int32(4) == int32(0) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	F_heap_toast_delete(m, l0, v18+int32(-24), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	goto L249
L253:
	;
	F_ReleaseBuffer(m, v48)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	if v436&int32(1) != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_UnlockTuple(m, l0, v102, int32(8))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	F_pgstat_count_heap_delete(m, l0)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L259
	}
L258:
	;
	goto L257
L259:
	;
	if v584 == int32(0) {
		v901 = v871
		goto L20
	} else {
		goto L260
	}
L260:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v888 != int32(1) {
		v901 = v871
		goto L20
	} else {
		goto L261
	}
L261:
	;
	F_pfree(m, v584)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v901 = v871
	goto L20
L263:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(273887), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(523527), int32(2806), int32(369904))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_fetch_next_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v4 != 0 {
		F_ReleaseBuffer(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			if v10 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if l1 != v13 {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v15
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						F_read_stream_reset(m, v17)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							v23 = F_read_stream_next_buffer(m, v21, int32(0))
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
								if v23 != 0 {
									if v23 < int32(0) {
										v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
										v44 = v35
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
										v44 = v43
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
								} else {
								}
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
							} else {
							}
							return
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				if l1 != v13 {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v15
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					F_read_stream_reset(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
							} else {
							}
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
						} else {
						}
						return
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		if v10 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				if l1 != v13 {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v15
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					F_read_stream_reset(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
							} else {
							}
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
						} else {
						}
						return
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			if l1 != v13 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v15
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				F_read_stream_reset(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
						} else {
						}
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l1
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				v23 = F_read_stream_next_buffer(m, v21, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v23
					if v23 != 0 {
						if v23 < int32(0) {
							v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
							v44 = v35
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23<<(uint(int32(6))%32)+int32(-64))+16))
							v44 = v43
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v44
					} else {
					}
					return
				}
			}
		}
	}
}
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 <= int32(1664) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < v15 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L14
	} else {
		goto L40
	}
L4:
	;
	v70 = F_heap_compute_data_size(m, l0, l1, l2)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v25 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v66 = int32(16)
	v67 = v5
	v69 = int32(0)
	goto L4
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v25))))
	if v32 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v38 = base.I32_div_s(v15+int32(7), int32(8))
	v66 = (v38 + int32(22)) & int32(-8)
	v67 = int32(128)
	v69 = int32(1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	v46 = v25 + int32(1)
	if v46 != v15 {
		v25 = v46
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	v74 = v70 + v66
	v76 = F_palloc0(m, v74+l3)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v80 = F__emscripten_memset_bulkmem(m, v76, base.I32_extend8_s(int32(0)), l3)
	mBase = m.M
	goto L17
L17:
	;
	v81 = v80 + l3
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v74
	v84 = v66 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)) = uint8(v84)
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+10)))
	v89 = v86&int32(63488) | v15
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+10)) = uint16(v89)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v81 + v66
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v67
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v98 = v81 + int32(15)
	goto L20
L19:
	;
	v98 = int32(0)
	goto L20
L20:
	;
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v102 = v98 - int32(1)
	goto L23
L22:
	;
	v102 = int32(0)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v102
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+12)))
	v106 = v104 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+12)) = uint16(v106)
	if int32(0) < v93 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	m.G0 = v13 + int32(32)
	return v81
L27:
	;
	v125 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v132 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v133 = v13 + int32(24)
	goto L31
L30:
	;
	v133 = v125
	goto L31
L31:
	;
	if l1 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1+v119<<(uint(int32(2))%32))))
	v142 = v141
	goto L34
L33:
	;
	v142 = v125
	goto L34
L34:
	;
	if l2 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v119))))
	v146 = v144
	goto L37
L36:
	;
	v146 = int32(1)
	goto L37
L37:
	;
	F_fill_val(m, l0+int32(20)+v119<<(uint(int32(4))%32), v133, v13+int32(20), v13+int32(28), v81+int32(12), v142, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v150 = v119 + int32(1)
	if v150 != v93 {
		v119 = v150
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(1664)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
	F_errmsg(m, int32(713007), v13)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(525249), int32(1473), int32(403386))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_get_root_tuples(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v13 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), int32(582))
	mBase = m.M
	goto L1
L1:
	;
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v14) < base.Ui32(int32(25)) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v22 = int32(base.Ui32(v14+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v22 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = v13 - int32(2)
	v28 = l0 + int32(24)
	v29 = int32(1)
	v32 = v29
	v35 = v29
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+v28-int32(4))))
	switch int32(base.Ui32(v45)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v168 = v35 + int32(1)
	v170 = v168 & int32(65535)
	if base.Ui32(v170) <= base.Ui32(v22) {
		v32 = v170
		v35 = v168
		goto L5
	} else {
		goto L39
	}
L8:
	;
	v90 = v85 & int32(65535)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32)+v28-int32(4))))
	if v96&int32(98304) != int32(32768) {
		goto L7
	} else {
		goto L20
	}
L9:
	;
	v85 = v45 & int32(32767)
	v88 = int32(0)
	goto L8
L10:
	;
	v54 = l0 + v45&int32(32767)
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+18)))
	if v55 < int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(1))%32)))) = uint16(v35)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+19)))
	if v62&int32(64) == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+20)))
	if v67&int32(2048) != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v67&int32(768) == int32(512) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
	if v67&int32(4224) == int32(4096) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = F_HeapTupleGetUpdateXid(m, v54)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v85 = v74
	v88 = v81
	goto L8
L18:
	;
	return
L19:
	;
	v85 = v74
	v88 = v79
	goto L8
L20:
	;
	v102 = v96
	v103 = v90
	v104 = v88
	goto L21
L21:
	;
	v112 = l0 + v102&int32(32767)
	if v104 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L7
L23:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+20)))
	v115 = int32(768)
	if v114&v115 != v115 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v103<<(uint(int32(1))%32)))) = uint16(v35)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+19)))
	if v127&int32(64) == int32(0) {
		goto L7
	} else {
		goto L30
	}
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v120 = v119
	goto L28
L27:
	;
	v120 = int32(2)
	goto L28
L28:
	;
	if v120 != v104 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+20)))
	if v132&int32(2048) != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v132&int32(768) == int32(512) {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	if v132&int32(4224) == int32(4096) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v139<<(uint(int32(2))%32)+v28-int32(4))))
	if v153&int32(98304) == int32(32768) {
		v102 = v153
		v103 = v139
		v104 = v147
		goto L21
	} else {
		goto L38
	}
L34:
	;
	v144 = F_HeapTupleGetUpdateXid(m, v112)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v147 = v146
	goto L33
L37:
	;
	v147 = v144
	goto L33
L38:
	;
	goto L22
L39:
	;
	goto L6
}
func F_heap_getattr_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+18)))
	if base.Ui32(v11&int32(2047)) <= base.Ui32(int32(20)) {
		v17 = F_getmissingattr(m, l1, int32(21), l2)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v77 = v17
			m.G0 = v8 + int32(16)
			return v77
		}
	} else {
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
		if v24&int32(1) == v21 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+340))
			if int32(0) <= v29 {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v34 = v23 + v32 + v29
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+346)))
				if v35 != int32(1) {
					v77 = v34
					m.G0 = v8 + int32(16)
					return v77
				} else {
					v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+344)))
					switch v38&int32(65535) - int32(1) {
					case 0:
						v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
						v77 = v43
						m.G0 = v8 + int32(16)
						return v77
					case 1:
						v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34))))
						v77 = v44
						m.G0 = v8 + int32(16)
						return v77
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
							F_errmsg_internal(m, int32(507614), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(344204), int32(70), int32(73868))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v77 = v45
						m.G0 = v8 + int32(16)
						return v77
					}
				}
			} else {
				v60 = F_nocachegetattr(m, l0, int32(21), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v77 = v60
					m.G0 = v8 + int32(16)
					return v77
				}
			}
		} else {
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)))
			if v62&int32(16) == int32(0) {
				v67 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v67)
				v77 = int32(0)
				m.G0 = v8 + int32(16)
				return v77
			} else {
				v71 = F_nocachegetattr(m, l0, int32(21), l1)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v77 = v71
					m.G0 = v8 + int32(16)
					return v77
				}
			}
		}
	}
}
func F_heap_getattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+18)))
	if base.Ui32(v13&int32(2047)) < base.Ui32(l1) {
		v17 = F_getmissingattr(m, l2, l1, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v84 = v17
			m.G0 = v10 + int32(16)
			return v84
		}
	} else {
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
		if v24&int32(1) == v21 {
			v29 = int32(4)
			v33 = l1<<(uint(v29)%32) + l2 + v29
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
			if v34 < int32(0) {
				v77 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v84 = v77
					m.G0 = v10 + int32(16)
					return v84
				}
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v39 = v23 + v37 + v34
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+6)))
				if v40 != int32(1) {
					v84 = v39
					m.G0 = v10 + int32(16)
					return v84
				} else {
					v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
					switch v43&int32(65535) - int32(1) {
					case 0:
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39))))
						v84 = v48
						m.G0 = v10 + int32(16)
						return v84
					case 1:
						v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39))))
						v84 = v49
						m.G0 = v10 + int32(16)
						return v84
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
							F_errmsg_internal(m, int32(507614), v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(344204), int32(70), int32(73868))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						v84 = v50
						m.G0 = v10 + int32(16)
						return v84
					}
				}
			}
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+26)))
			v65 = int32(1)
			if int32(base.Ui32(v64)>>(uint((l1-v65)&int32(7))%32))&v65 != 0 {
				v77 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v84 = v77
					m.G0 = v10 + int32(16)
					return v84
				}
			} else {
				v72 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v72)
				v84 = int32(0)
				m.G0 = v10 + int32(16)
				return v84
			}
		}
	}
}
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v6 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = F_palloc(m, v14)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_heap_deform_tuple(m, l0, l1, v17, v21)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v14 <= int32(0) {
				} else {
					v27 = int32(1)
					if v14 != v27 {
						v38 = v6
						v42 = v6
						for {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v38))))
							if v47 == int32(1) {
								v51 = v38 << (uint(int32(2)) % 32)
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v51)))
								*(*int32)(unsafe.Add(mBase, uint32(v17+v51))) = v54
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v38))))
								*(*uint8)(unsafe.Add(mBase, uint32(v38+v21))) = uint8(v58)
							} else {
							}
							v61 = int32(1)
							v62 = v38 | v61
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v62))))
							if v64 == v61 {
								v68 = v62 << (uint(int32(2)) % 32)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l2+v68)))
								*(*int32)(unsafe.Add(mBase, uint32(v17+v68))) = v71
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v62))))
								*(*uint8)(unsafe.Add(mBase, uint32(v62+v21))) = uint8(v75)
							} else {
							}
							v78 = int32(2)
							v79 = v38 + v78
							v81 = v42 + v78
							if v81 != v14&int32(2147483646) {
								v38 = v79
								v42 = v81
								continue
							} else {
								break
							}
							break
						}
						v88 = v79
					} else {
						v88 = v6
					}
					if v14&v27 == int32(0) {
					} else {
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v88))))
						if v99 != int32(1) {
						} else {
							v103 = v88 << (uint(int32(2)) % 32)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l2+v103)))
							*(*int32)(unsafe.Add(mBase, uint32(v17+v103))) = v106
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v88))))
							*(*uint8)(unsafe.Add(mBase, uint32(v88+v21))) = uint8(v110)
						}
					}
				}
				v125 = F_heap_form_tuple(m, l1, v17, v21)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v17)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v21)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v133
							v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)) = uint16(v135)
							v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
							*(*uint16)(unsafe.Add(mBase, uint32(v125)+8)) = uint16(v137)
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v139
							v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = v141
							return v125
						}
					}
				}
			}
		}
	}
}
func F_heap_modify_tuple_by_cols(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = F_palloc(m, v18<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = F_palloc(m, v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_heap_deform_tuple(m, l0, l1, v21, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(0) < l2 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v37 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v79 = F_heap_form_tuple(m, l1, v21, v25)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v45 = v37 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3+v45)))
	if v47 <= int32(0) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	if v18 < v47 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v51 = int32(1)
	v52 = v47 - v51
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l4+v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v52<<(uint(int32(2))%32)))) = v57
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52+v25))) = uint8(v61)
	v64 = v37 + v51
	if v64 != l2 {
		v37 = v64
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	F_pfree(m, v21)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v25)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v87
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)) = uint16(v89)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)) = uint16(v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v95
	m.G0 = v16 + int32(16)
	return v79
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v47
	F_errmsg_internal(m, int32(495256), v16)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(525249), int32(1305), int32(161947))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_set_tidrange(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v10 == int32(0) {
	} else {
		v13 = int32(2048)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v13)
		v15 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
		v20 = v10 - v15
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v20)
		v22 = int32(16)
		v23 = int32(base.Ui32(v20) >> (uint(v22) % 32))
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)) = uint16(v23)
		v26 = v8 + int32(8)
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
		v34 = v30 | v31<<(uint(v22)%32)
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
		v39 = v35 | v36<<(uint(v22)%32)
		if base.Ui32(v34) < base.Ui32(v39) {
			v50 = int32(-1)
		} else {
			if base.Ui32(v39) < base.Ui32(v34) {
				v50 = int32(1)
			} else {
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
				if base.Ui32(v44) < base.Ui32(v45) {
					v50 = int32(-1)
				} else {
					v50 = base.B2i32(base.Ui32(v45) < base.Ui32(v44))
				}
			}
		}
		if v50 < int32(0) {
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v53)
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v55
		} else {
		}
		v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
		v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		v62 = int32(16)
		v64 = v60 | v61<<(uint(v62)%32)
		v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+2)))
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
		v69 = v65 | v66<<(uint(v62)%32)
		if base.Ui32(v64) < base.Ui32(v69) {
			v80 = int32(-1)
		} else {
			if base.Ui32(v69) < base.Ui32(v64) {
				v80 = int32(1)
			} else {
				v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
				if base.Ui32(v74) < base.Ui32(v75) {
					v80 = int32(-1)
				} else {
					v80 = base.B2i32(base.Ui32(v75) < base.Ui32(v74))
				}
			}
		}
		if int32(0) < v80 {
			v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)) = uint16(v83)
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v85
		} else {
		}
		v88 = v8 + int32(8)
		v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
		v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
		v94 = int32(16)
		v96 = v92 | v93<<(uint(v94)%32)
		v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+2)))
		v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
		v101 = v97 | v98<<(uint(v94)%32)
		if base.Ui32(v96) < base.Ui32(v101) {
			v112 = int32(-1)
		} else {
			if base.Ui32(v101) < base.Ui32(v96) {
				v112 = int32(1)
			} else {
				v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)))
				v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
				if base.Ui32(v106) < base.Ui32(v107) {
					v112 = int32(-1)
				} else {
					v112 = base.B2i32(base.Ui32(v107) < base.Ui32(v106))
				}
			}
		}
		if v112 < int32(0) {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
		} else {
			v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)))
			v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+2)))
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
			v121 = int32(16)
			v123 = v119 | v120<<(uint(v121)%32)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v123
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v117 | v118<<(uint(v121)%32) - v123 + int32(1)
			v132 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v132
			v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v134)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+22)) = v136
			v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v138)
		}
	}
	m.G0 = v8 + int32(16)
	return
}
func F_heap_tuple_infomask_flags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)) = uint16(v2)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_superuser(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v406
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v417 = F_heap_form_tuple(m, v412, v12+int32(8), v12+int32(6))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L125
	}
L2:
	;
	F_pfree(m, v320)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L124
	}
L3:
	;
	v400 = F_construct_array_builtin(m, v320, v398, int32(25))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L123
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L120
	}
L5:
	;
	return int32(0)
L6:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = F_get_call_result_type(m, l0, int32(0), v12)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L116
	}
L10:
	;
	if v25 != int32(1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(8)
	v31 = int32(base.Ui32(v19) >> (uint(v30) % 32))
	v32 = int32(255)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31&v32)+uint32(_consts[811]))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v32)+uint32(_consts[811]))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18&v32)+uint32(_consts[811]))))
	v47 = int32(base.Ui32(v18) >> (uint(v30) % 32))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47&v32)+uint32(_consts[811]))))
	v52 = v35 + v39 + v44 + v51
	if v52 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v64 = v52 << (uint(int32(2)) % 32)
	v65 = F_palloc0(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v56
	v60 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v406 = v60
	goto L1
L17:
	;
	if v19&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = F_cstring_to_text(m, int32(559780))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v74 = int32(0)
	goto L20
L20:
	;
	if v19&int32(2) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
	v74 = int32(1)
	goto L20
L22:
	;
	v81 = F_cstring_to_text(m, int32(563083))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	v86 = v74
	goto L24
L24:
	;
	if v19&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v74<<(uint(int32(2))%32)))) = v81
	v86 = v74 + int32(1)
	goto L24
L26:
	;
	v93 = F_cstring_to_text(m, int32(561425))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	v98 = v86
	goto L28
L28:
	;
	if v19&int32(8) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v86<<(uint(int32(2))%32)))) = v93
	v98 = v86 + int32(1)
	goto L28
L30:
	;
	v105 = F_cstring_to_text(m, int32(570415))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v110 = v98
	goto L32
L32:
	;
	if v19&int32(16) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v98<<(uint(int32(2))%32)))) = v105
	v110 = v98 + int32(1)
	goto L32
L34:
	;
	v117 = F_cstring_to_text(m, int32(561613))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v122 = v110
	goto L36
L36:
	;
	if v19&int32(32) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v110<<(uint(int32(2))%32)))) = v117
	v122 = v110 + int32(1)
	goto L36
L38:
	;
	v129 = F_cstring_to_text(m, int32(571073))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v134 = v122
	goto L40
L40:
	;
	if v19&int32(64) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v122<<(uint(int32(2))%32)))) = v129
	v134 = v122 + int32(1)
	goto L40
L42:
	;
	v141 = F_cstring_to_text(m, int32(561635))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v146 = v134
	goto L44
L44:
	;
	if v19&int32(128) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v134<<(uint(int32(2))%32)))) = v141
	v146 = v134 + int32(1)
	goto L44
L46:
	;
	v153 = F_cstring_to_text(m, int32(536636))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	v158 = v146
	goto L48
L48:
	;
	if v31&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v146<<(uint(int32(2))%32)))) = v153
	v158 = v146 + int32(1)
	goto L48
L50:
	;
	v165 = F_cstring_to_text(m, int32(571584))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	v170 = v158
	goto L52
L52:
	;
	if v31&int32(2) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v158<<(uint(int32(2))%32)))) = v165
	v170 = v158 + int32(1)
	goto L52
L54:
	;
	v177 = F_cstring_to_text(m, int32(570952))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	v182 = v170
	goto L56
L56:
	;
	if v31&int32(4) != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v170<<(uint(int32(2))%32)))) = v177
	v182 = v170 + int32(1)
	goto L56
L58:
	;
	v189 = F_cstring_to_text(m, int32(571564))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	v194 = v182
	goto L60
L60:
	;
	if v31&int32(8) != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v182<<(uint(int32(2))%32)))) = v189
	v194 = v182 + int32(1)
	goto L60
L62:
	;
	v201 = F_cstring_to_text(m, int32(570934))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L65
	}
L63:
	;
	v206 = v194
	goto L64
L64:
	;
	if v31&int32(16) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v194<<(uint(int32(2))%32)))) = v201
	v206 = v194 + int32(1)
	goto L64
L66:
	;
	v213 = F_cstring_to_text(m, int32(561700))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	v218 = v206
	goto L68
L68:
	;
	if v31&int32(32) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v206<<(uint(int32(2))%32)))) = v213
	v218 = v206 + int32(1)
	goto L68
L70:
	;
	v225 = F_cstring_to_text(m, int32(571713))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	v230 = v218
	goto L72
L72:
	;
	if v31&int32(64) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v218<<(uint(int32(2))%32)))) = v225
	v230 = v218 + int32(1)
	goto L72
L74:
	;
	v237 = F_cstring_to_text(m, int32(564630))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	v242 = v230
	goto L76
L76:
	;
	v244 = v31 << (uint(int32(8)) % 32)
	if base.I32_extend16_s(v244) < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v230<<(uint(int32(2))%32)))) = v237
	v242 = v230 + int32(1)
	goto L76
L78:
	;
	v252 = F_cstring_to_text(m, int32(557700))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L81
	}
L79:
	;
	v257 = v242
	goto L80
L80:
	;
	if v47&int32(32) != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v242<<(uint(int32(2))%32)))) = v252
	v257 = v242 + int32(1)
	goto L80
L82:
	;
	v264 = F_cstring_to_text(m, int32(571695))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	v269 = v257
	goto L84
L84:
	;
	if v47&int32(64) != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v257<<(uint(int32(2))%32)))) = v264
	v269 = v257 + int32(1)
	goto L84
L86:
	;
	v276 = F_cstring_to_text(m, int32(571678))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L89
	}
L87:
	;
	v281 = v269
	goto L88
L88:
	;
	if v47&int32(128) != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v269<<(uint(int32(2))%32)))) = v276
	v281 = v269 + int32(1)
	goto L88
L90:
	;
	v288 = F_cstring_to_text(m, int32(567722))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L93
	}
L91:
	;
	v293 = v281
	goto L92
L92:
	;
	v295 = F_construct_array_builtin(m, v65, v293, int32(25))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v281<<(uint(int32(2))%32)))) = v288
	v293 = v281 + int32(1)
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v295
	if base.Ui32(int32(256)) < base.Ui32(v52) {
		v316 = v64
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v317 = int32(0)
	v320 = F__emscripten_memset_bulkmem(m, v65, base.I32_extend8_s(v317), v316)
	mBase = m.M
	goto L101
L96:
	;
	if v65&int32(3) != 0 {
		v316 = v64
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v304 = v65 + v64
	v306 = v65 + int32(4)
	if base.Ui32(v306) < base.Ui32(v304) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v308 = v304
	goto L100
L99:
	;
	v308 = v306
	goto L100
L100:
	;
	v316 = (v65^int32(-1)+v308)&int32(-4) + int32(4)
	goto L95
L101:
	;
	v321 = int32(80)
	if v19&v321 == v321 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v326 = F_cstring_to_text(m, int32(561594))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L105
	}
L103:
	;
	v330 = v317
	goto L104
L104:
	;
	v331 = int32(3)
	if v31&v331 == v331 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v326
	v330 = int32(1)
	goto L104
L106:
	;
	v339 = F_cstring_to_text(m, int32(558105))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L109
	}
L107:
	;
	v344 = v330
	goto L108
L108:
	;
	if base.Ui32(int32(49152)) <= base.Ui32(v244&int32(65535)) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320+v330<<(uint(int32(2))%32)))) = v339
	v344 = v330 + int32(1)
	goto L108
L110:
	;
	v353 = F_cstring_to_text(m, int32(571536))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if v344 != 0 {
		v398 = v344
		goto L3
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320+v344<<(uint(int32(2))%32)))) = v353
	v398 = v344 + int32(1)
	goto L3
L114:
	;
	v359 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v403 = v359
	goto L2
L116:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(150263), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(520260), int32(537), int32(167189))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errmsg_internal(m, int32(386892), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(520260), int32(541), int32(167189))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	v403 = v400
	goto L2
L124:
	;
	v406 = v403
	goto L1
L125:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	v420 = F_HeapTupleHeaderGetDatum(m, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	m.G0 = v12 + int32(16)
	return v420
}
func F_heap_tuple_needs_eventual_freeze(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v5 = int32(768)
	if v4&v5 == v5 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v4&int32(4096) != 0 {
			if v14 == int32(0) {
				if base.Ui32(v4) < base.Ui32(int32(16384)) {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v27) <= base.Ui32(int32(2)) {
						return int32(0)
					} else {
						return int32(1)
					}
				}
			} else {
				return int32(1)
			}
		} else {
			if base.Ui32(v14) <= base.Ui32(int32(2)) {
				if base.Ui32(v4) < base.Ui32(int32(16384)) {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v27) <= base.Ui32(int32(2)) {
						return int32(0)
					} else {
						return int32(1)
					}
				}
			} else {
				return int32(1)
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v9) <= base.Ui32(int32(2)) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v4&int32(4096) != 0 {
				if v14 == int32(0) {
					if base.Ui32(v4) < base.Ui32(int32(16384)) {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v27) <= base.Ui32(int32(2)) {
							return int32(0)
						} else {
							return int32(1)
						}
					}
				} else {
					return int32(1)
				}
			} else {
				if base.Ui32(v14) <= base.Ui32(int32(2)) {
					if base.Ui32(v4) < base.Ui32(int32(16384)) {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v27) <= base.Ui32(int32(2)) {
							return int32(0)
						} else {
							return int32(1)
						}
					}
				} else {
					return int32(1)
				}
			}
		} else {
			return int32(1)
		}
	}
}
