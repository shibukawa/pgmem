package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == int32(_a_F_ExecStoreHeapTuple_0) {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v8&int32(4) != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			F_pfree(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v16
				v18 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v18)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v28 = v17 & int32(_a_F_ExecStoreHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v28)
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v32
				if l2 != 0 {
					v35 = v28 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v35)
				} else {
				}
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v37
				return l1
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v18)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v28 = v17 & int32(_a_F_ExecStoreHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v28)
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v32
			if l2 != 0 {
				v35 = v28 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v35)
			} else {
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v37
			return l1
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_ExecStoreHeapTuple_2), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ExecStoreHeapTuple_3), int32(1553), int32(_a_F_ExecStoreHeapTuple_4))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
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
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
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
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[1]))
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
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v88))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v79)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31&int32(_a_F_HeapCheckForSerializableConflictOut_0) != int32(_a_F_HeapCheckForSerializableConflictOut_1) {
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
	F_errmsg_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_2), v13)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_3), int32(_a_F_HeapCheckForSerializableConflictOut_4), int32(_a_F_HeapCheckForSerializableConflictOut_5))
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
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
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
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
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
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L176
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L2
	} else {
		goto L170
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
	if base.Ui32(v170) < base.Ui32(int32(_a_F_HeapCheckForSerializableConflictOut_6)) {
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
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[1]))
	if v138 == v180 {
		goto L50
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+28)) = v138
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
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
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[4]))
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
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	v205 = F_LWLockAcquire(m, v201+int32(_a_F_HeapCheckForSerializableConflictOut_7), int32(1))
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
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	if v337 != v339 {
		goto L105
	} else {
		goto L106
	}
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[5]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v212+int32(_a_F_HeapCheckForSerializableConflictOut_7))
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
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v332+int32(3584))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
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
	v247 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_HeapCheckForSerializableConflictOut_8), base.I64_extend_i32_u(v245), v138)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[6]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v247<<(uint(int32(2))%32))))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v255+v138<<(uint(int32(3))%32)&int32(_a_F_HeapCheckForSerializableConflictOut_9))))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[7])))
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
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
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
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
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
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
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
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_11), v156)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4072), int32(_a_F_HeapCheckForSerializableConflictOut_14))
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
	if v341&int32(1024) != 0 {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+108))
	if v341&int32(8) == int32(0) {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v348+int32(3584))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	if v341&int32(2) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v341&int32(1) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v337)+108)) = v341 | int32(8)
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v363+int32(3584))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v369+int32(3584))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_15), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(_a_F_HeapCheckForSerializableConflictOut_16), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	v418 = int32(0)
	v420 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L133
	}
L125:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v339)+108))
	if v402&int32(32) == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v341&int32(16) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v339)+24))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v337)+24))
	if base.Ui64(v410) <= base.Ui64(v409) {
		goto L124
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v413+int32(3584))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
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
	if v492 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L133:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v422))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v434 != 0 {
		v492 = v418
		goto L132
	} else {
		goto L138
	}
L135:
	;
	v434 = base.B2i32(base.Ui32(v138) < base.Ui32(v422))
	goto L134
L136:
	;
	goto L137
L137:
	;
	v434 = int32(base.Ui32(v138-v422) >> (uint(int32(31)) % 32))
	goto L134
L138:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v436))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v138)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v448 != 0 {
		v492 = int32(1)
		goto L132
	} else {
		goto L143
	}
L140:
	;
	v448 = base.B2i32(base.Ui32(v436) <= base.Ui32(v138))
	goto L139
L141:
	;
	goto L142
L142:
	;
	v448 = base.B2i32(int32(0) <= v138-v436)
	goto L139
L143:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v420)+16))
	if v449 == int32(0) {
		v473 = v418
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v492 = v473
	goto L132
L145:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v454 = int32(0)
	goto L146
L146:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v452+v454<<(uint(int32(2))%32))))
	v468 = base.B2i32(v138 == v467)
	if v138 == v467 {
		v473 = v468
		goto L144
	} else {
		goto L148
	}
L147:
	;
	v473 = v468
	goto L144
L148:
	;
	v470 = v454 + int32(1)
	if v470 != v449 {
		v454 = v470
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v496+int32(3584))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L2
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v501 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503)+108)))
	if v504&int32(8) != 0 {
		v537 = v501
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L50
L154:
	;
	if v537 != 0 {
		goto L164
	} else {
		goto L165
	}
L155:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+108)))
	if v507&int32(8) != 0 {
		v537 = v501
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v503)+36))
	if v510 == int32(0) {
		v537 = v501
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v514 = v503 + int32(32)
	if v510 == v514 {
		v537 = v501
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v337)+44))
	if base.B2i32(v516 == int32(0))|base.B2i32(v516 == v337+int32(40)) != 0 {
		v537 = v501
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v524 = v510
	goto L160
L160:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v524)+20))
	v534 = base.B2i32(v533 == v337)
	if v533 == v337 {
		v537 = v534
		goto L154
	} else {
		goto L162
	}
L161:
	;
	v537 = v534
	goto L154
L162:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if v535 != v514 {
		v524 = v535
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v548+int32(3584))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	F_FlagRWConflict(m, v503, v337)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L168
	}
L167:
	;
	goto L50
L168:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v556+int32(3584))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	goto L50
L170:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L172
	}
L172:
	;
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_17), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4039), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v138
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_18), v156+int32(16))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4080), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	if v14&int32(_a_F_HeapTupleHeaderAdvanceConflictHorizon_0) != int32(_a_F_HeapTupleHeaderAdvanceConflictHorizon_1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if base.Ui32(v69&int32(_a_F_HeapTupleHeaderAdvanceConflictHorizon_2)) < base.Ui32(int32(_a_F_HeapTupleHeaderAdvanceConflictHorizon_3)) {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	v69 = v14
	v71 = v21
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
	v62 = int32(0)
	goto L12
L12:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v69 = v66
	v71 = v62
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
	v62 = v55
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
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v110))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v71)) == int32(0) {
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
	if v71 == v20 {
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
	if v71 != v20 {
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
	v122 = base.B2i32(base.Ui32(v110) < base.Ui32(v71))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v122 = base.B2i32(int32(0) < v71-v110)
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71
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
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v879 int32
	_ = v879
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v8&int32(256) != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	v1151 = v1037 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v1151)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L55
	} else {
		goto L427
	}
L2:
	;
	v1125 = int32(4)
	v1128 = l0 + v1125
	v1130 = v7 + int32(12)
	v1131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1128)+2)))
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1128))))
	v1133 = int32(16)
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1130)+2)))
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1130))))
	if v1131|v1132<<(uint(v1133)%32) == v1136|v1137<<(uint(v1133)%32) {
		goto L420
	} else {
		goto L421
	}
L3:
	;
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v1116 = v1114 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v1116)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L55
	} else {
		goto L417
	}
L4:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1098&int32(32) != 0 {
		goto L411
	} else {
		goto L412
	}
L5:
	;
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L55
	} else {
		goto L409
	}
L6:
	;
	return v1082
L7:
	;
	v731 = int32(0)
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v732&int32(2048) != 0 {
		v1082 = v731
		goto L6
	} else {
		goto L276
	}
L8:
	;
	v11 = int32(1)
	v12 = base.I32_extend16_s(v8)
	if v12&int32(512) != 0 {
		v1082 = v11
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v12&int32(_a_F_HeapTupleSatisfiesUpdate_0) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v724 = v144 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v724)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L55
	} else {
		goto L275
	}
L11:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if base.Ui32(v17) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	if v12 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	if v137 != 0 {
		v1082 = v11
		goto L6
	} else {
		goto L54
	}
L15:
	;
	v137 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v28 == v17 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v137 = int32(1)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v32 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v137 = v129
	goto L14
L22:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v36 == int32(0) {
		v129 = int32(0)
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v100 = int32(0)
	v102 = v32 - int32(1)
	goto L44
L25:
	;
	v41 = v36
	goto L26
L26:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v46 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v129 = int32(0)
	goto L21
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	if v93 != 0 {
		v41 = v93
		goto L26
	} else {
		goto L43
	}
L29:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v49 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v52 = int32(1)
	if v17 == v49 {
		v129 = v52
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v56 = v54 - int32(1)
	if v56 < int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v61 = int32(0)
	v63 = v56
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v69 = int32(2)
	v70 = base.I32_div_s(v63-v61, v69)
	v71 = v70 + v61
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v67+v71<<(uint(v69)%32))))
	if v75 == v17 {
		v129 = v52
		goto L21
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v79 = F_TransactionIdPrecedes(m, v75, v17)
	mBase = m.M
	if v79 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v80 = v71 + int32(1)
	goto L38
L37:
	;
	v80 = v61
	goto L38
L38:
	;
	if v79 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v83 = v63
	goto L41
L40:
	;
	v83 = v71 - int32(1)
	goto L41
L41:
	;
	if v80 <= v83 {
		v61 = v80
		v63 = v83
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	goto L27
L44:
	;
	v107 = int32(2)
	v108 = base.I32_div_s(v102-v100, v107)
	v109 = v108 + v100
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v98+v109<<(uint(v107)%32))))
	v114 = base.B2i32(v113 == v17)
	if v113 == v17 {
		v129 = v114
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v129 = v114
	goto L21
L46:
	;
	v117 = base.B2i32(base.Ui32(v113) < base.Ui32(v17))
	if base.Ui32(v113) < base.Ui32(v17) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v118 = v109 + int32(1)
	goto L49
L48:
	;
	v118 = v100
	goto L49
L49:
	;
	if base.Ui32(v113) < base.Ui32(v17) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v121 = v102
	goto L52
L51:
	;
	v121 = v109 - int32(1)
	goto L52
L52:
	;
	if v118 <= v121 {
		v100 = v118
		v102 = v121
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v138 = F_TransactionIdIsInProgress(m, v17)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	if v138 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v142 = F_TransactionIdDidCommit(m, v17)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v142 == int32(0) {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v148 = v144 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v148)
	goto L5
L60:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if base.Ui32(v152) < base.Ui32(int32(3)) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if base.Ui32(v287) < base.Ui32(int32(3)) {
		goto L112
	} else {
		goto L113
	}
L63:
	;
	if v272 != 0 {
		goto L7
	} else {
		goto L103
	}
L64:
	;
	v272 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v163 == v152 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v272 = int32(1)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v167 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v272 = v264
	goto L63
L71:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v171 == int32(0) {
		v264 = int32(0)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v235 = int32(0)
	v237 = v167 - int32(1)
	goto L93
L74:
	;
	v176 = v171
	goto L75
L75:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	if v181 == int32(4) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v264 = int32(0)
	goto L70
L77:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v176)+80))
	if v228 != 0 {
		v176 = v228
		goto L75
	} else {
		goto L92
	}
L78:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v184 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v187 = int32(1)
	if v152 == v184 {
		v264 = v187
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	v191 = v189 - int32(1)
	if v191 < int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v196 = int32(0)
	v198 = v191
	goto L82
L82:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	v204 = int32(2)
	v205 = base.I32_div_s(v198-v196, v204)
	v206 = v205 + v196
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(v204)%32))))
	if v210 == v152 {
		v264 = v187
		goto L70
	} else {
		goto L84
	}
L83:
	;
	goto L77
L84:
	;
	v214 = F_TransactionIdPrecedes(m, v210, v152)
	mBase = m.M
	if v214 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v215 = v206 + int32(1)
	goto L87
L86:
	;
	v215 = v196
	goto L87
L87:
	;
	if v214 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v218 = v198
	goto L90
L89:
	;
	v218 = v206 - int32(1)
	goto L90
L90:
	;
	if v215 <= v218 {
		v196 = v215
		v198 = v218
		goto L82
	} else {
		goto L91
	}
L91:
	;
	goto L83
L92:
	;
	goto L76
L93:
	;
	v242 = int32(2)
	v243 = base.I32_div_s(v237-v235, v242)
	v244 = v243 + v235
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(v242)%32))))
	v249 = base.B2i32(v248 == v152)
	if v248 == v152 {
		v264 = v249
		goto L70
	} else {
		goto L95
	}
L94:
	;
	v264 = v249
	goto L70
L95:
	;
	v252 = base.B2i32(base.Ui32(v248) < base.Ui32(v152))
	if base.Ui32(v248) < base.Ui32(v152) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v253 = v244 + int32(1)
	goto L98
L97:
	;
	v253 = v235
	goto L98
L98:
	;
	if base.Ui32(v248) < base.Ui32(v152) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v256 = v237
	goto L101
L100:
	;
	v256 = v244 - int32(1)
	goto L101
L101:
	;
	if v253 <= v256 {
		v235 = v253
		v237 = v256
		goto L93
	} else {
		goto L102
	}
L102:
	;
	goto L94
L103:
	;
	v273 = F_TransactionIdIsInProgress(m, v152)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L55
	} else {
		goto L104
	}
L104:
	;
	if v273 != 0 {
		v1082 = v11
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v275 = F_TransactionIdDidCommit(m, v152)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L55
	} else {
		goto L106
	}
L106:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v275 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v279 = v277 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v279)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L55
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v285 = v277 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v285)
	goto L5
L110:
	;
	goto L7
L111:
	;
	if v407 != 0 {
		goto L151
	} else {
		goto L152
	}
L112:
	;
	v407 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v298 == v287 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v407 = int32(1)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v302 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v407 = v399
	goto L111
L119:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v306 == int32(0) {
		v399 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v370 = int32(0)
	v372 = v302 - int32(1)
	goto L141
L122:
	;
	v311 = v306
	goto L123
L123:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	if v316 == int32(4) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v399 = int32(0)
	goto L118
L125:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v311)+80))
	if v363 != 0 {
		v311 = v363
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v319 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v322 = int32(1)
	if v287 == v319 {
		v399 = v322
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v311)+52))
	v326 = v324 - int32(1)
	if v326 < int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v331 = int32(0)
	v333 = v326
	goto L130
L130:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v311)+48))
	v339 = int32(2)
	v340 = base.I32_div_s(v333-v331, v339)
	v341 = v340 + v331
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337+v341<<(uint(v339)%32))))
	if v345 == v287 {
		v399 = v322
		goto L118
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v349 = F_TransactionIdPrecedes(m, v345, v287)
	mBase = m.M
	if v349 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v350 = v341 + int32(1)
	goto L135
L134:
	;
	v350 = v331
	goto L135
L135:
	;
	if v349 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v353 = v333
	goto L138
L137:
	;
	v353 = v341 - int32(1)
	goto L138
L138:
	;
	if v350 <= v353 {
		v331 = v350
		v333 = v353
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	goto L124
L141:
	;
	v377 = int32(2)
	v378 = base.I32_div_s(v372-v370, v377)
	v379 = v378 + v370
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v368+v379<<(uint(v377)%32))))
	v384 = base.B2i32(v383 == v287)
	if v383 == v287 {
		v399 = v384
		goto L118
	} else {
		goto L143
	}
L142:
	;
	v399 = v384
	goto L118
L143:
	;
	v387 = base.B2i32(base.Ui32(v383) < base.Ui32(v287))
	if base.Ui32(v383) < base.Ui32(v287) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v388 = v379 + int32(1)
	goto L146
L145:
	;
	v388 = v370
	goto L146
L146:
	;
	if base.Ui32(v383) < base.Ui32(v287) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v391 = v372
	goto L149
L148:
	;
	v391 = v379 - int32(1)
	goto L149
L149:
	;
	if v388 <= v391 {
		v370 = v388
		v372 = v391
		goto L141
	} else {
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v410&int32(32) != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	goto L153
L153:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v710 = F_TransactionIdIsInProgress(m, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L55
	} else {
		goto L268
	}
L154:
	;
	if base.Ui32(l1) <= base.Ui32(v419) {
		v1082 = v11
		goto L6
	} else {
		goto L158
	}
L155:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[4]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v409<<(uint(int32(3))%32))))
	v419 = v418
	goto L157
L156:
	;
	v419 = v409
	goto L157
L157:
	;
	goto L154
L158:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v422&int32(2048) != 0 {
		v1082 = int32(0)
		goto L6
	} else {
		goto L159
	}
L159:
	;
	v427 = int32(0)
	if base.B2i32(v422&int32(128) == v427)&base.B2i32(v422&int32(_a_F_HeapTupleSatisfiesUpdate_1) != int32(64)) == v427 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v422&int32(_a_F_HeapTupleSatisfiesUpdate_2) != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	if v422&int32(_a_F_HeapTupleSatisfiesUpdate_2) != 0 {
		goto L174
	} else {
		goto L175
	}
L163:
	;
	v442 = F_MultiXactIdIsRunning(m, v436, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L55
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v448 = F_TransactionIdIsInProgress(m, v436)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L55
	} else {
		goto L170
	}
L166:
	;
	if v442 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v444 = int32(5)
	goto L169
L168:
	;
	v444 = int32(0)
	goto L169
L169:
	;
	return v444
L170:
	;
	if v448 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v450 = int32(5)
	goto L173
L172:
	;
	v450 = int32(0)
	goto L173
L173:
	;
	return v450
L174:
	;
	v454 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L55
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v586) < base.Ui32(int32(3)) {
		goto L226
	} else {
		goto L227
	}
L177:
	;
	if base.Ui32(v454) < base.Ui32(int32(3)) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v575 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L179:
	;
	v575 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v466 == v454 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v575 = int32(1)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v470 <= int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v575 = v567
	goto L178
L186:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v474 == int32(0) {
		v567 = int32(0)
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v538 = int32(0)
	v540 = v470 - int32(1)
	goto L208
L189:
	;
	v479 = v474
	goto L190
L190:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	if v484 == int32(4) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v567 = int32(0)
	goto L185
L192:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v479)+80))
	if v531 != 0 {
		v479 = v531
		goto L190
	} else {
		goto L207
	}
L193:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v487 == int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v490 = int32(1)
	if v454 == v487 {
		v567 = v490
		goto L185
	} else {
		goto L195
	}
L195:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v479)+52))
	v494 = v492 - int32(1)
	if v494 < int32(0) {
		goto L192
	} else {
		goto L196
	}
L196:
	;
	v499 = int32(0)
	v501 = v494
	goto L197
L197:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v507 = int32(2)
	v508 = base.I32_div_s(v501-v499, v507)
	v509 = v508 + v499
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v505+v509<<(uint(v507)%32))))
	if v513 == v454 {
		v567 = v490
		goto L185
	} else {
		goto L199
	}
L198:
	;
	goto L192
L199:
	;
	v517 = F_TransactionIdPrecedes(m, v513, v454)
	mBase = m.M
	if v517 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v518 = v509 + int32(1)
	goto L202
L201:
	;
	v518 = v499
	goto L202
L202:
	;
	if v517 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v521 = v501
	goto L205
L204:
	;
	v521 = v509 - int32(1)
	goto L205
L205:
	;
	if v518 <= v521 {
		v499 = v518
		v501 = v521
		goto L197
	} else {
		goto L206
	}
L206:
	;
	goto L198
L207:
	;
	goto L191
L208:
	;
	v545 = int32(2)
	v546 = base.I32_div_s(v540-v538, v545)
	v547 = v546 + v538
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v536+v547<<(uint(v545)%32))))
	v552 = base.B2i32(v551 == v454)
	if v551 == v454 {
		v567 = v552
		goto L185
	} else {
		goto L210
	}
L209:
	;
	v567 = v552
	goto L185
L210:
	;
	v555 = base.B2i32(base.Ui32(v551) < base.Ui32(v454))
	if base.Ui32(v551) < base.Ui32(v454) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v556 = v547 + int32(1)
	goto L213
L212:
	;
	v556 = v538
	goto L213
L213:
	;
	if base.Ui32(v551) < base.Ui32(v454) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v559 = v540
	goto L216
L215:
	;
	v559 = v547 - int32(1)
	goto L216
L216:
	;
	if v556 <= v559 {
		v538 = v556
		v540 = v559
		goto L208
	} else {
		goto L217
	}
L217:
	;
	goto L209
L218:
	;
	v579 = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v582 = F_MultiXactIdIsRunning(m, v580, v579)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L55
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	goto L4
L221:
	;
	if v582 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v584 = int32(5)
	goto L224
L223:
	;
	v584 = v579
	goto L224
L224:
	;
	return v584
L225:
	;
	if v706 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L226:
	;
	v706 = int32(0)
	goto L225
L227:
	;
	goto L228
L228:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v597 == v586 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v706 = int32(1)
	goto L225
L230:
	;
	goto L231
L231:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v601 <= int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v706 = v698
	goto L225
L233:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v605 == int32(0) {
		v698 = int32(0)
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v669 = int32(0)
	v671 = v601 - int32(1)
	goto L255
L236:
	;
	v610 = v605
	goto L237
L237:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v610)+20))
	if v615 == int32(4) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v698 = int32(0)
	goto L232
L239:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v610)+80))
	if v662 != 0 {
		v610 = v662
		goto L237
	} else {
		goto L254
	}
L240:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	if v618 == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v621 = int32(1)
	if v586 == v618 {
		v698 = v621
		goto L232
	} else {
		goto L242
	}
L242:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v610)+52))
	v625 = v623 - int32(1)
	if v625 < int32(0) {
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v630 = int32(0)
	v632 = v625
	goto L244
L244:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v610)+48))
	v638 = int32(2)
	v639 = base.I32_div_s(v632-v630, v638)
	v640 = v639 + v630
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v636+v640<<(uint(v638)%32))))
	if v644 == v586 {
		v698 = v621
		goto L232
	} else {
		goto L246
	}
L245:
	;
	goto L239
L246:
	;
	v648 = F_TransactionIdPrecedes(m, v644, v586)
	mBase = m.M
	if v648 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v649 = v640 + int32(1)
	goto L249
L248:
	;
	v649 = v630
	goto L249
L249:
	;
	if v648 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v652 = v632
	goto L252
L251:
	;
	v652 = v640 - int32(1)
	goto L252
L252:
	;
	if v649 <= v652 {
		v630 = v649
		v632 = v652
		goto L244
	} else {
		goto L253
	}
L253:
	;
	goto L245
L254:
	;
	goto L238
L255:
	;
	v676 = int32(2)
	v677 = base.I32_div_s(v671-v669, v676)
	v678 = v677 + v669
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v667+v678<<(uint(v676)%32))))
	v683 = base.B2i32(v682 == v586)
	if v682 == v586 {
		v698 = v683
		goto L232
	} else {
		goto L257
	}
L256:
	;
	v698 = v683
	goto L232
L257:
	;
	v686 = base.B2i32(base.Ui32(v682) < base.Ui32(v586))
	if base.Ui32(v682) < base.Ui32(v586) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v687 = v678 + int32(1)
	goto L260
L259:
	;
	v687 = v669
	goto L260
L260:
	;
	if base.Ui32(v682) < base.Ui32(v586) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v690 = v671
	goto L263
L262:
	;
	v690 = v678 - int32(1)
	goto L263
L263:
	;
	if v687 <= v690 {
		v669 = v687
		v671 = v690
		goto L255
	} else {
		goto L264
	}
L264:
	;
	goto L256
L265:
	;
	goto L3
L266:
	;
	goto L267
L267:
	;
	goto L4
L268:
	;
	if v710 != 0 {
		v1082 = v11
		goto L6
	} else {
		goto L269
	}
L269:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v713 = F_TransactionIdDidCommit(m, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L55
	} else {
		goto L270
	}
L270:
	;
	if v713 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	F_HeapTupleSetHintBits(m, v7, l2, int32(256), v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L55
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v721 = v719 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)) = uint16(v721)
	goto L5
L274:
	;
	goto L7
L275:
	;
	goto L7
L276:
	;
	if v732&int32(1024) != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	if v732&int32(128)|base.B2i32(v732&int32(_a_F_HeapTupleSatisfiesUpdate_1) == int32(64)) != 0 {
		v1082 = v731
		goto L6
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	if v732&int32(_a_F_HeapTupleSatisfiesUpdate_2) != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	goto L2
L281:
	;
	if v732&int32(_a_F_HeapTupleSatisfiesUpdate_3) == int32(_a_F_HeapTupleSatisfiesUpdate_4) {
		v1082 = v731
		goto L6
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v899) < base.Ui32(int32(3)) {
		goto L345
	} else {
		goto L346
	}
L284:
	;
	if v732&int32(128) != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v754 = F_MultiXactIdIsRunning(m, v752, int32(1))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L55
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v758 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L55
	} else {
		goto L293
	}
L288:
	;
	if v754 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	return int32(5)
L290:
	;
	goto L291
L291:
	;
	goto L3
L292:
	;
	if base.Ui32(v758) < base.Ui32(int32(3)) {
		goto L298
	} else {
		goto L299
	}
L293:
	;
	if v758 != 0 {
		goto L292
	} else {
		goto L294
	}
L294:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v762 = F_MultiXactIdIsRunning(m, v760, int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L55
	} else {
		goto L295
	}
L295:
	;
	if v762 == int32(0) {
		goto L292
	} else {
		goto L296
	}
L296:
	;
	return int32(5)
L297:
	;
	if v887 != 0 {
		goto L4
	} else {
		goto L337
	}
L298:
	;
	v887 = int32(0)
	goto L297
L299:
	;
	goto L300
L300:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v778 == v758 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v887 = int32(1)
	goto L297
L302:
	;
	goto L303
L303:
	;
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v782 <= int32(0) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v887 = v879
	goto L297
L305:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v786 == int32(0) {
		v879 = int32(0)
		goto L304
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v848 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v850 = int32(0)
	v852 = v782 - int32(1)
	goto L327
L308:
	;
	v791 = v786
	goto L309
L309:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v791)+20))
	if v796 == int32(4) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v879 = int32(0)
	goto L304
L311:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v791)+80))
	if v843 != 0 {
		v791 = v843
		goto L309
	} else {
		goto L326
	}
L312:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v791)))
	if v799 == int32(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v802 = int32(1)
	if v758 == v799 {
		v879 = v802
		goto L304
	} else {
		goto L314
	}
L314:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v791)+52))
	v806 = v804 - int32(1)
	if v806 < int32(0) {
		goto L311
	} else {
		goto L315
	}
L315:
	;
	v811 = int32(0)
	v813 = v806
	goto L316
L316:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v791)+48))
	v819 = int32(2)
	v820 = base.I32_div_s(v813-v811, v819)
	v821 = v820 + v811
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v817+v821<<(uint(v819)%32))))
	if v825 == v758 {
		v879 = v802
		goto L304
	} else {
		goto L318
	}
L317:
	;
	goto L311
L318:
	;
	v829 = F_TransactionIdPrecedes(m, v825, v758)
	mBase = m.M
	if v829 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v830 = v821 + int32(1)
	goto L321
L320:
	;
	v830 = v811
	goto L321
L321:
	;
	if v829 != 0 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v833 = v813
	goto L324
L323:
	;
	v833 = v821 - int32(1)
	goto L324
L324:
	;
	if v830 <= v833 {
		v811 = v830
		v813 = v833
		goto L316
	} else {
		goto L325
	}
L325:
	;
	goto L317
L326:
	;
	goto L310
L327:
	;
	v857 = int32(2)
	v858 = base.I32_div_s(v852-v850, v857)
	v859 = v858 + v850
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v848+v859<<(uint(v857)%32))))
	v864 = base.B2i32(v863 == v758)
	if v863 == v758 {
		v879 = v864
		goto L304
	} else {
		goto L329
	}
L328:
	;
	v879 = v864
	goto L304
L329:
	;
	v867 = base.B2i32(base.Ui32(v863) < base.Ui32(v758))
	if base.Ui32(v863) < base.Ui32(v758) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v868 = v859 + int32(1)
	goto L332
L331:
	;
	v868 = v850
	goto L332
L332:
	;
	if base.Ui32(v863) < base.Ui32(v758) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v871 = v852
	goto L335
L334:
	;
	v871 = v859 - int32(1)
	goto L335
L335:
	;
	if v868 <= v871 {
		v850 = v868
		v852 = v871
		goto L327
	} else {
		goto L336
	}
L336:
	;
	goto L328
L337:
	;
	v888 = int32(5)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v891 = F_MultiXactIdIsRunning(m, v889, int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L55
	} else {
		goto L338
	}
L338:
	;
	if v891 != 0 {
		v1082 = v888
		goto L6
	} else {
		goto L339
	}
L339:
	;
	v893 = F_TransactionIdDidCommit(m, v758)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L55
	} else {
		goto L340
	}
L340:
	;
	if v893 != 0 {
		goto L2
	} else {
		goto L341
	}
L341:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v897 = F_MultiXactIdIsRunning(m, v895, int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L55
	} else {
		goto L342
	}
L342:
	;
	if v897 != 0 {
		v1082 = v888
		goto L6
	} else {
		goto L343
	}
L343:
	;
	goto L3
L344:
	;
	if v1019 != 0 {
		goto L384
	} else {
		goto L385
	}
L345:
	;
	v1019 = int32(0)
	goto L344
L346:
	;
	goto L347
L347:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v910 == v899 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1019 = int32(1)
	goto L344
L349:
	;
	goto L350
L350:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v914 <= int32(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1019 = v1011
	goto L344
L352:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v918 == int32(0) {
		v1011 = int32(0)
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v980 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v982 = int32(0)
	v984 = v914 - int32(1)
	goto L374
L355:
	;
	v923 = v918
	goto L356
L356:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v923)+20))
	if v928 == int32(4) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1011 = int32(0)
	goto L351
L358:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v923)+80))
	if v975 != 0 {
		v923 = v975
		goto L356
	} else {
		goto L373
	}
L359:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v923)))
	if v931 == int32(0) {
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v934 = int32(1)
	if v899 == v931 {
		v1011 = v934
		goto L351
	} else {
		goto L361
	}
L361:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v923)+52))
	v938 = v936 - int32(1)
	if v938 < int32(0) {
		goto L358
	} else {
		goto L362
	}
L362:
	;
	v943 = int32(0)
	v945 = v938
	goto L363
L363:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v923)+48))
	v951 = int32(2)
	v952 = base.I32_div_s(v945-v943, v951)
	v953 = v952 + v943
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v949+v953<<(uint(v951)%32))))
	if v957 == v899 {
		v1011 = v934
		goto L351
	} else {
		goto L365
	}
L364:
	;
	goto L358
L365:
	;
	v961 = F_TransactionIdPrecedes(m, v957, v899)
	mBase = m.M
	if v961 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v962 = v953 + int32(1)
	goto L368
L367:
	;
	v962 = v943
	goto L368
L368:
	;
	if v961 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v965 = v945
	goto L371
L370:
	;
	v965 = v953 - int32(1)
	goto L371
L371:
	;
	if v962 <= v965 {
		v943 = v962
		v945 = v965
		goto L363
	} else {
		goto L372
	}
L372:
	;
	goto L364
L373:
	;
	goto L357
L374:
	;
	v989 = int32(2)
	v990 = base.I32_div_s(v984-v982, v989)
	v991 = v990 + v982
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v980+v991<<(uint(v989)%32))))
	v996 = base.B2i32(v995 == v899)
	if v995 == v899 {
		v1011 = v996
		goto L351
	} else {
		goto L376
	}
L375:
	;
	v1011 = v996
	goto L351
L376:
	;
	v999 = base.B2i32(base.Ui32(v995) < base.Ui32(v899))
	if base.Ui32(v995) < base.Ui32(v899) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1000 = v991 + int32(1)
	goto L379
L378:
	;
	v1000 = v982
	goto L379
L379:
	;
	if base.Ui32(v995) < base.Ui32(v899) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1003 = v984
	goto L382
L381:
	;
	v1003 = v991 - int32(1)
	goto L382
L382:
	;
	if v1000 <= v1003 {
		v982 = v1000
		v984 = v1003
		goto L374
	} else {
		goto L383
	}
L383:
	;
	goto L375
L384:
	;
	v1021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1021&int32(128)|base.B2i32(v1021&int32(_a_F_HeapTupleSatisfiesUpdate_1) == int32(64)) != 0 {
		v1082 = int32(5)
		goto L6
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v1030 = F_TransactionIdIsInProgress(m, v1029)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L55
	} else {
		goto L388
	}
L387:
	;
	goto L4
L388:
	;
	if v1030 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	return int32(5)
L390:
	;
	goto L391
L391:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v1035 = F_TransactionIdDidCommit(m, v1034)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L55
	} else {
		goto L392
	}
L392:
	;
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v1035 == int32(0) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	goto L1
L394:
	;
	goto L395
L395:
	;
	v1042 = int32(0)
	if base.B2i32(v1037&int32(128) == v1042)&base.B2i32(v1037&int32(_a_F_HeapTupleSatisfiesUpdate_1) != int32(64)) == v1042 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	goto L1
L397:
	;
	goto L398
L398:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	F_HeapTupleSetHintBits(m, v7, l2, int32(1024), v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L55
	} else {
		goto L399
	}
L399:
	;
	v1055 = int32(4)
	v1058 = l0 + v1055
	v1060 = v7 + int32(12)
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+2)))
	v1062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058))))
	v1063 = int32(16)
	v1066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1060)+2)))
	v1067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1060))))
	if v1061|v1062<<(uint(v1063)%32) == v1066|v1067<<(uint(v1063)%32) {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	if v1077 != 0 {
		goto L406
	} else {
		goto L407
	}
L401:
	;
	goto L400
L402:
	;
	v1073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+4)))
	v1074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1060)+4)))
	if v1073 == v1074 {
		v1077 = int32(1)
		goto L401
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	v1077 = int32(0)
	goto L401
L405:
	;
	goto L404
L406:
	;
	v1078 = v1055
	goto L408
L407:
	;
	v1078 = int32(3)
	goto L408
L408:
	;
	v1082 = v1078
	goto L6
L409:
	;
	return int32(1)
L410:
	;
	if base.Ui32(v1107) < base.Ui32(l1) {
		goto L414
	} else {
		goto L415
	}
L411:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[4]))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1102+v1097<<(uint(int32(3))%32))+4))
	v1107 = v1106
	goto L413
L412:
	;
	v1107 = v1097
	goto L413
L413:
	;
	goto L410
L414:
	;
	v1109 = int32(1)
	goto L416
L415:
	;
	v1109 = int32(2)
	goto L416
L416:
	;
	return v1109
L417:
	;
	return int32(0)
L418:
	;
	if v1147 != 0 {
		goto L424
	} else {
		goto L425
	}
L419:
	;
	goto L418
L420:
	;
	v1143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1128)+4)))
	v1144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1130)+4)))
	if v1143 == v1144 {
		v1147 = int32(1)
		goto L419
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v1147 = int32(0)
	goto L419
L423:
	;
	goto L422
L424:
	;
	v1148 = v1125
	goto L426
L425:
	;
	v1148 = int32(3)
	goto L426
L426:
	;
	return v1148
L427:
	;
	return int32(0)
}
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v114 int32
	_ = v114
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
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v380 int64
	_ = v380
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	v17 = m.G0
	v19 = v17 - int32(160)
	m.G0 = v19
	*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = int64(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v27 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+136)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v60
	F_swap_relation_files(m, l0, l1, base.B2i32(l0 == int32(1259)), l3, l5, l6, l7, v19+int32(128))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v31&int32(1) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v36 = int32(_a_F_finish_heap_swap_0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v39 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v38 + v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v42 + v39
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(8))+232)) = int64(5)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v50 + v39
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v56 - v39
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
	v72 = int32(1)
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
		goto L48
	} else {
		goto L49
	}
L10:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[3]))
	v146 = F_PrepareInvalidationState(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L35
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
	if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v143 = v72
		goto L11
	} else {
		goto L34
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
	if l0 <= int32(_a_F_finish_heap_swap_1) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v143 = v72
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
	v84 = l0 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v84))|base.B2i32(int32(1)<<(uint(v84)%32)&int32(226492515) == int32(0)) != 0 {
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
	v143 = v72
	goto L11
L22:
	;
	v143 = v72
	goto L11
L23:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v143 = v72
		goto L11
	} else {
		goto L32
	}
L24:
	;
	v97 = l0 - int32(_a_F_finish_heap_swap_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v97))|base.B2i32(int32(1)<<(uint(v97)%32)&int32(963) == int32(0)) != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	switch l0 - int32(_a_F_finish_heap_swap_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v143 = v72
		goto L11
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L12
	default:
		goto L28
	}
L27:
	;
	v143 = v72
	goto L11
L28:
	;
	if base.Ui32(l0-int32(_a_F_finish_heap_swap_4)) < base.Ui32(int32(3)) {
		v143 = v72
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v114 = l0 - int32(_a_F_finish_heap_swap_5)
	if base.Ui32(int32(15)) < base.Ui32(v114) {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	if int32(1)<<(uint(v114)%32)&int32(_a_F_finish_heap_swap_6) != 0 {
		v143 = v72
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v143 = v72
	goto L11
L34:
	;
	goto L12
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[4]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[5]))
	if v152 <= v150 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v149 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v174 = v149
	goto L38
L38:
	;
	v178 = v174 + v150<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = l0
	if v143 != 0 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[5])) = v168
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[4])) = v169
	v174 = v169
	goto L38
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[6]))
	v160 = F_MemoryContextAlloc(m, v158, int32(512))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v166 = F_repalloc(m, v149, v152<<(uint(int32(5))%32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v168 = int32(32)
	v169 = v160
	goto L39
L44:
	;
	v168 = v152 << (uint(int32(1)) % 32)
	v169 = v166
	goto L39
L45:
	;
	v181 = int32(0)
	goto L47
L46:
	;
	v181 = v145
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v181
	v183 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v183)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v185 + int32(1)
	goto L9
L48:
	;
	v198 = int32(6)
	goto L50
L49:
	;
	v198 = int32(2)
	goto L50
L50:
	;
	switch l8 - int32(112) {
	case 0:
		goto L52
	default:
		v205 = v198
		goto L51
	case 5:
		goto L53
	}
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v210 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v205 = v198 | int32(16)
	goto L51
L53:
	;
	v205 = v198 | int32(8)
	goto L51
L54:
	;
	v243 = int32(0)
	v247 = F_reindex_relation(m, v243, l0, v205, v19+int32(120))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v214&int32(1) == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v219 = int32(_a_F_finish_heap_swap_0)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v222 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v221 + v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v225 + v222
	*(*int64)(unsafe.Add(mBase, uint32(v210+int32(8))+232)) = int64(6)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v233 + v222
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v239 - v222
	goto L55
L58:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v253 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if l0 == int32(1259) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	goto L59
L61:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v257&int32(1) == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v262 = int32(_a_F_finish_heap_swap_0)
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v265 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v264 + v265
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v268 + v265
	*(*int64)(unsafe.Add(mBase, uint32(v253+int32(8))+232)) = int64(7)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v276 + v265
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v282 - v265
	goto L60
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L122
	}
L64:
	;
	v290 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v314 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+156)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = int32(1259)
	F_performDeletion(m, v19+int32(148), v314, int32(1))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L72
	}
L67:
	;
	v295 = F_SearchSysCacheCopy(m, int32(57), int32(1259), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	if v295 == int32(0) {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+22)))
	v301 = v299 + v300
	*(*int32)(unsafe.Add(mBase, uint32(v301)+140)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v301)+136)) = l6
	F_CatalogTupleUpdate(m, v290, v295+int32(4), v295)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	F_relation_close(m, v290, int32(3))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	goto L66
L72:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v19)+128))
	if v325 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v330 = v325
	v331 = v243
	goto L76
L74:
	;
	goto L75
L75:
	;
	if l3 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	v342 = int32(0)
	v343 = m.G0
	v345 = v343 - int32(16)
	m.G0 = v345
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[7]))
	if v348 <= v342 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L75
L78:
	;
	v421 = v331 + int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(128)+v421<<(uint(int32(2))%32))))
	if v425 != 0 {
		v330 = v425
		v331 = v421
		goto L76
	} else {
		goto L90
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L87
	}
L80:
	;
	v352 = v342
	goto L81
L81:
	;
	v368 = v352 << (uint(int32(3)) % 32)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_finish_heap_swap[8])))
	if v369 != v330 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v348<<(uint(int32(3))%32))+uint32(_c_F_finish_heap_swap[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_finish_heap_swap[8]))) = v380
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[7])) = v348 - int32(1)
	m.G0 = v345 + int32(16)
	goto L78
L83:
	;
	v372 = v352 + int32(1)
	if v348 != v372 {
		v352 = v372
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L79
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v330
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_7), v345)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_8), int32(454), int32(_a_F_finish_heap_swap_9))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	goto L77
L91:
	;
	v445 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if l2 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L94:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v445)+48))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+112))
	if v448 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v450 = F_toast_get_valid_index(m, v448, int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	F_relation_close(m, v445, int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L5
	} else {
		goto L115
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l0
	v454 = v19 + int32(48)
	v459 = F_pg_snprintf(m, v454, int32(64), int32(_a_F_finish_heap_swap_10), v19+int32(32))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v445)+48))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+112))
	F_RenameRelationInternal(m, v462, v454, int32(1), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	v472 = F_pg_snprintf(m, v454, int32(64), int32(_a_F_finish_heap_swap_11), v19+int32(16))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v474 = int32(1)
	F_RenameRelationInternal(m, v450, v454, v474, v474)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v445)+48))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+112))
	v482 = m.G0
	v484 = v482 - int32(16)
	m.G0 = v484
	v488 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v492 = F_SearchSysCacheCopy(m, int32(57), v481, int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v492 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v509+v510)+132)) = int32(0)
	F_CatalogTupleUpdate(m, v488, v492+int32(4), v492)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L5
	} else {
		goto L112
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v481
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_12), v484)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_13), int32(_a_F_finish_heap_swap_14), int32(_a_F_finish_heap_swap_15))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_pfree(m, v492)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	F_relation_close(m, v488, int32(3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	m.G0 = v484 + int32(16)
	goto L97
L115:
	;
	goto L93
L116:
	;
	v541 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	m.G0 = v19 + int32(160)
	return
L119:
	;
	F_RelationClearMissing(m, v541)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	F_relation_close(m, v541, int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(1259)
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_12), v19)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_16), int32(1543), int32(_a_F_finish_heap_swap_17))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
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
	var v30 int32
	_ = v30
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
	v30 = v4
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
	v38 = v35 + v30*int32(36)
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
	v93 = v30 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	if v93 < v94 {
		v30 = v93
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v17 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				base.MemoryCopy(m, v15, v18, v17)
			} else {
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20 << (uint(int32(2)) % 32)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v26
			return v15
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v2 = int32(0)
	if l0 == v2 {
		v33 = v2
		return v33
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v7 == int32(0) {
			v33 = v2
			return v33
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v28 == int32(0) {
					v33 = v13
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					base.MemoryCopy(m, v25, v31, v28)
					v33 = v13
				}
				return v33
			}
		}
	}
}
func F_heap_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int64
	_ = v870
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v938 int32
	_ = v938
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = l2
	v25 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v29
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)) = uint8(v29)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[0]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v22 - int32(-64)
	return v938
L4:
	;
	if v49 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L5:
	;
	F_UnlockReleaseBuffer(m, v49)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L166
	}
L6:
	;
	if v39&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v39 = int32(1)
	goto L9
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+76)))
	v39 = v38
	goto L9
L9:
	;
	goto L6
L10:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v48 = v44 | v45<<(uint(int32(16))%32)
	v49 = F_ReadBuffer(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L162
	}
L13:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
	if v69&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v49 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v49^int32(-1))<<(uint(int32(2))%32))))
	v68 = v60
	goto L13
L16:
	;
	goto L17
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[2]))
	v68 = v62 + v49<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L18:
	;
	F_visibilitymap_pin(m, l0, v48, v20+int32(-28))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_LockBuffer(m, v49, int32(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v80
	v83 = v68 + int32(20)
	v86 = v83 + v79<<(uint(int32(2))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v68 + v87&int32(_a_F_heap_delete_0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(base.Ui32(v92) >> (uint(int32(17)) % 32))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v96
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+48)) = uint16(v98)
	v101 = v20 + int32(-20)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v105 = v102
	v116 = int32(0)
	goto L26
L23:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)) = uint16(v446)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v445)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v448
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445)+20)))
	if v451&int32(_a_F_heap_delete_1) != int32(_a_F_heap_delete_2) {
		goto L136
	} else {
		goto L137
	}
L24:
	;
	v424 = int32(0)
	if base.B2i32(l3 == v424)|v419 == v424 {
		goto L129
	} else {
		goto L130
	}
L25:
	;
	v379 = int32(0)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v380)+20)))
	if v381&int32(2048)|v381&int32(128)|base.B2i32(v381&int32(_a_F_heap_delete_3) == int32(64)) != 0 {
		v419 = v379
		v421 = v377
		goto L24
	} else {
		goto L117
	}
L26:
	;
	if v105 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v349&int32(3072) != 0 {
		v377 = v347
		goto L25
	} else {
		goto L110
	}
L28:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v140 = F_HeapTupleSatisfiesUpdate(m, v20+int32(-24), v139, v49)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
	if v122&int32(4) == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_visibilitymap_pin(m, l0, v48, v20+int32(-28))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_LockBuffer(m, v49, int32(2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	if v140 == int32(1) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(l4 == int32(0))|base.B2i32(v140 != int32(5)) != 0 {
		v419 = v140
		v421 = v116
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+20)))
	if v151&int32(_a_F_heap_delete_2) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)) = uint8(v154)
	v159 = F_DoesMultiXactIdConflict(m, v150, v151, int32(3), v20+int32(-48))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v150) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	if v159 == int32(0) {
		v377 = v116
		goto L25
	} else {
		goto L41
	}
L41:
	;
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)))
	if (v166|v116)&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v179 = int32(0)
	v183 = F_Do_MultiXactIdWait(m, v150, int32(5), v151, v179, l0, v101, int32(2), v179, v179)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	v177 = v166 ^ int32(1) | v116
	goto L43
L45:
	;
	goto L46
L46:
	;
	F_LockTuple(m, l0, v101, int32(8))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v177 = int32(1)
	goto L43
L48:
	;
	F_LockBuffer(m, v49, int32(2))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v188 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
	if v191&int32(4) != 0 {
		v105 = v188
		v116 = v177
		goto L26
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+20)))
	if (v195^v151)&int32(_a_F_heap_delete_4) != 0 {
		v105 = v188
		v116 = v177
		goto L26
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v199 != v150 {
		v105 = v188
		v116 = v177
		goto L26
	} else {
		goto L55
	}
L55:
	;
	v377 = v177
	goto L25
L56:
	;
	if v320 != 0 {
		v377 = v116
		goto L25
	} else {
		goto L96
	}
L57:
	;
	v320 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[3]))
	if v211 == v150 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v320 = int32(1)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[4]))
	if v215 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v320 = v312
	goto L56
L64:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[0]))
	if v219 == int32(0) {
		v312 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[5]))
	v283 = int32(0)
	v285 = v215 - int32(1)
	goto L86
L67:
	;
	v224 = v219
	goto L68
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	if v229 == int32(4) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v312 = int32(0)
	goto L63
L70:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v224)+80))
	if v276 != 0 {
		v224 = v276
		goto L68
	} else {
		goto L85
	}
L71:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v232 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v235 = int32(1)
	if v150 == v232 {
		v312 = v235
		goto L63
	} else {
		goto L73
	}
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
	v239 = v237 - int32(1)
	if v239 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v244 = int32(0)
	v246 = v239
	goto L75
L75:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v224)+48))
	v252 = int32(2)
	v253 = base.I32_div_s(v246-v244, v252)
	v254 = v253 + v244
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v250+v254<<(uint(v252)%32))))
	if v258 == v150 {
		v312 = v235
		goto L63
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v262 = F_TransactionIdPrecedes(m, v258, v150)
	mBase = m.M
	if v262 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v263 = v254 + int32(1)
	goto L80
L79:
	;
	v263 = v244
	goto L80
L80:
	;
	if v262 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v266 = v246
	goto L83
L82:
	;
	v266 = v254 - int32(1)
	goto L83
L83:
	;
	if v263 <= v266 {
		v244 = v263
		v246 = v266
		goto L75
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	goto L69
L86:
	;
	v290 = int32(2)
	v291 = base.I32_div_s(v285-v283, v290)
	v292 = v291 + v283
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v281+v292<<(uint(v290)%32))))
	v297 = base.B2i32(v296 == v150)
	if v296 == v150 {
		v312 = v297
		goto L63
	} else {
		goto L88
	}
L87:
	;
	v312 = v297
	goto L63
L88:
	;
	v300 = base.B2i32(base.Ui32(v296) < base.Ui32(v150))
	if base.Ui32(v296) < base.Ui32(v150) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v301 = v292 + int32(1)
	goto L91
L90:
	;
	v301 = v283
	goto L91
L91:
	;
	if base.Ui32(v296) < base.Ui32(v150) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v304 = v285
	goto L94
L93:
	;
	v304 = v292 - int32(1)
	goto L94
L94:
	;
	if v301 <= v304 {
		v283 = v301
		v285 = v304
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L87
L96:
	;
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v116&int32(1) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_LockTuple(m, l0, v101, int32(8))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_XactLockTableWait(m, v150, l0, v101, int32(2))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	F_LockBuffer(m, v49, int32(2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v337 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
	if v342&int32(4) != 0 {
		v105 = int32(0)
		v116 = int32(1)
		goto L26
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v347 = int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+20)))
	if (v151^v349)&int32(_a_F_heap_delete_4) != 0 {
		v105 = v337
		v116 = v347
		goto L26
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v353 != v150 {
		v105 = v337
		v116 = v347
		goto L26
	} else {
		goto L109
	}
L109:
	;
	goto L27
L110:
	;
	if v349&int32(128)|base.B2i32(v349&int32(_a_F_heap_delete_3) == int32(64)) != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_HeapTupleSetHintBits(m, v348, v49, int32(2048), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L116
	}
L112:
	;
	v364 = F_TransactionIdDidCommit(m, v150)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v364 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	F_HeapTupleSetHintBits(m, v348, v49, int32(1024), v150)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v377 = v347
	goto L25
L116:
	;
	v377 = v347
	goto L25
L117:
	;
	v392 = F_HeapTupleHeaderIsOnlyLocked(m, v380)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	if v392 != 0 {
		v419 = v379
		v421 = v377
		goto L24
	} else {
		goto L119
	}
L119:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v398 = v396 + int32(12)
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+2)))
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	v401 = int32(16)
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+2)))
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398))))
	if v399|v400<<(uint(v401)%32) == v404|v405<<(uint(v401)%32) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v415 != 0 {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	goto L120
L122:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+4)))
	if v411 == v412 {
		v415 = int32(1)
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v415 = int32(0)
	goto L121
L125:
	;
	goto L124
L126:
	;
	v416 = int32(4)
	goto L128
L127:
	;
	v416 = int32(3)
	goto L128
L128:
	;
	v440 = v416
	v442 = v377
	goto L23
L129:
	;
	v432 = F_HeapTupleSatisfiesVisibility(m, v20+int32(-24), l3, v49)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v419 == int32(0) {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	if v432 == int32(0) {
		v440 = int32(3)
		v442 = v421
		goto L23
	} else {
		goto L133
	}
L133:
	;
	goto L4
L134:
	;
	v440 = v419
	v442 = v421
	goto L23
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v500
	if v440 == int32(2) {
		goto L148
	} else {
		goto L149
	}
L136:
	;
	v500 = v450
	goto L135
L137:
	;
	goto L138
L138:
	;
	v456 = int32(0)
	v460 = F_GetMultiXactIdMembers(m, v450, v20+int32(-48), v456)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v460 <= int32(0) {
		v500 = v456
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v466 = v456
	goto L143
L141:
	;
	F_pfree(m, v464)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L147
	}
L142:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v496 = v494
	goto L141
L143:
	;
	v486 = v464 + v466<<(uint(int32(3))%32)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v487) {
		goto L142
	} else {
		goto L145
	}
L144:
	;
	v496 = int32(0)
	goto L141
L145:
	;
	v491 = v466 + int32(1)
	if v491 != v460 {
		v466 = v491
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v500 = v496
	goto L135
L148:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v521)+8))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+20)))
	if v524&int32(32) != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v535 = int32(-1)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v535
	F_UnlockReleaseBuffer(m, v49)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L155
	}
L151:
	;
	v535 = v533
	goto L150
L152:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[6]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v523<<(uint(int32(3))%32))+4))
	v533 = v532
	goto L154
L153:
	;
	v533 = v523
	goto L154
L154:
	;
	goto L151
L155:
	;
	if v442&int32(1) != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_UnlockTuple(m, l0, v101, int32(8))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v544 == int32(0) {
		v938 = v440
		goto L3
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	F_ReleaseBuffer(m, v544)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v938 = v440
	goto L3
L162:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(_a_F_heap_delete_5), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_heap_delete_6), int32(2806), int32(_a_F_heap_delete_7))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_heap_delete_8), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_heap_delete_6), int32(2853), int32(_a_F_heap_delete_7))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_CheckForSerializableConflictIn(m, l0, l1, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L175
	}
L172:
	;
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[7]))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587+(v49^int32(-1))<<(uint(int32(6))%32))+16))
	v602 = v593
	goto L171
L173:
	;
	goto L174
L174:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[8]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v595+v49<<(uint(int32(6))%32)+int32(-64))+16))
	v602 = v601
	goto L171
L175:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	F_HeapTupleHeaderAdjustCmax(m, v605, v20+int32(-4), v20+int32(-37))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v617 = F_ExtractReplicaIdentity(m, l0, v20+int32(-24), int32(1), v20+int32(-38))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+4))
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+20)))
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+18)))
	F_compute_new_xmax_infomask(m, v622, v623, v624, v25, int32(3), int32(1), v20+int32(-32), v20+int32(-34), v20+int32(-36))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v635 = int32(_a_F_heap_delete_9)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9])) = v637 + int32(1)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v641 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+10)))
	v659 = v657 & int32(4)
	if v659 != 0 {
		goto L189
	} else {
		goto L190
	}
L181:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v641))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v25)) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v25
	goto L180
L184:
	;
	if v653 == int32(0) {
		goto L180
	} else {
		goto L188
	}
L185:
	;
	v653 = base.B2i32(base.Ui32(v25) < base.Ui32(v641))
	goto L184
L186:
	;
	goto L187
L187:
	;
	v653 = int32(base.Ui32(v25-v641) >> (uint(int32(31)) % 32))
	goto L184
L188:
	;
	goto L183
L189:
	;
	v661 = v657 & int32(_a_F_heap_delete_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v68)+10)) = uint16(v661)
	if v49 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686)+20)))
	v689 = v687 & int32(_a_F_heap_delete_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v686)+20)) = uint16(v689)
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686)+18)))
	v693 = v691 & int32(_a_F_heap_delete_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v686)+18)) = uint16(v693)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v695)+20)))
	v697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+30)))
	v698 = v696 | v697
	*(*uint16)(unsafe.Add(mBase, uint32(v695)+20)) = uint16(v698)
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v695)+18)))
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+28)))
	v702 = v700 | v701
	*(*uint16)(unsafe.Add(mBase, uint32(v695)+18)) = uint16(v702)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+18)))
	v707 = v705 & int32(_a_F_heap_delete_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v704)+18)) = uint16(v707)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v709)+4)) = v710
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+27)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = v714
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v713)+20)))
	v723 = v716&int32(_a_F_heap_delete_14) | v712<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v713)+20)) = uint16(v723)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v725)+16)) = uint16(v726)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v725)+12)) = v728
	if l6 != 0 {
		goto L197
	} else {
		goto L198
	}
L192:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v684 = F_visibilitymap_clear(m, v681, v682, int32(3))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L196
	}
L193:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[7]))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v666+(v49^int32(-1))<<(uint(int32(6))%32))+16))
	v681 = v672
	goto L192
L194:
	;
	goto L195
L195:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[8]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v674+v49<<(uint(int32(6))%32)+int32(-64))+16))
	v681 = v680
	goto L192
L196:
	;
	goto L191
L197:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v731 = int32(_a_F_heap_delete_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v730)+16)) = uint16(v731)
	*(*int32)(unsafe.Add(mBase, uint32(v730)+12)) = int32(-1)
	goto L199
L198:
	;
	goto L199
L199:
	;
	F_MarkBufferDirty(m, v49)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+118)))
	if v739 != int32(112) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v878 = int32(_a_F_heap_delete_9)
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9])) = v880 - int32(1)
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L238
	}
L202:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[10]))
	if v743 <= int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v774 = int32(base.Ui32(v659) >> (uint(int32(2)) % 32))
	if l6 != 0 {
		goto L218
	} else {
		goto L219
	}
L204:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v746 != 0 {
		goto L201
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	if v743 == int32(1) {
		goto L203
	} else {
		goto L209
	}
L207:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v747 == int32(0) {
		goto L203
	} else {
		goto L208
	}
L208:
	;
	goto L201
L209:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L210
L210:
	;
	if base.B2i32(base.Ui32(v752) < base.Ui32(int32(_a_F_heap_delete_16))) == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v757 == int32(0) {
		goto L203
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	F_log_heap_new_cid(m, l0, v20+int32(-24))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L217
	}
L214:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760)+119)))
	switch v761 - int32(109) {
	case 0, 5:
		goto L215
	default:
		goto L203
	}
L215:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757)+104)))
	if v764 != int32(1) {
		goto L203
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	goto L203
L218:
	;
	v777 = v774 | int32(16)
	goto L220
L219:
	;
	v777 = v774
	goto L220
L220:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)) = uint8(v777)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v779)+18)))
	v781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v779)+20)))
	v782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+48)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)) = uint16(v782)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v710
	v789 = int32(1)
	v793 = int32(4)
	v808 = int32(base.Ui32(v780)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v781)>>(uint(v789)%32))&int32(8) | (int32(base.Ui32(v781)>>(uint(v793)%32))&v793 | (int32(base.Ui32(v781)>>(uint(int32(12))%32))&v789 | int32(base.Ui32(v781)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)) = uint8(v808)
	if v617 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v863 = int32(_a_F_heap_delete_17)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[11])))
	v866 = v865 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[11])) = uint8(v866)
	goto L236
L222:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+130)))
	if v813 == int32(102) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L233
	}
L225:
	;
	v816 = int32(2)
	goto L227
L226:
	;
	v816 = int32(4)
	goto L227
L227:
	;
	v817 = v777 | v816
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)) = uint8(v817)
	F_XLogBeginInsert(m)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_XLogRegisterData(m, v20+int32(-48), int32(8))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_XLogRegisterBuffer(m, int32(0), v49, int32(8))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
	v831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v830)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+10)) = uint16(v831)
	v833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v830)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+12)) = uint16(v833)
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+14)) = uint8(v835)
	F_XLogRegisterData(m, v20+int32(-54), int32(5))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
	v843 = int32(23)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	F_XLogRegisterData(m, v842+v843, v845-v843)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	goto L221
L233:
	;
	F_XLogRegisterData(m, v20+int32(-48), int32(8))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_XLogRegisterBuffer(m, int32(0), v49, int32(8))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	goto L221
L236:
	;
	v870 = F_XLogInsert(m, int32(10), int32(16))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = base.I64_rotr(v870, int64(32))
	goto L201
L238:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v887 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	F_ReleaseBuffer(m, v887)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890)+119)))
	switch v891 - int32(109) {
	case 0, 5:
		goto L244
	default:
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v905 = int32(0)
	F_CacheInvalidateHeapTuple(m, l0, v20+int32(-24), v905)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L247
	}
L244:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v894)+20)))
	if v895&int32(4) == int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	F_heap_toast_delete(m, l0, v20+int32(-24), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	goto L243
L247:
	;
	F_ReleaseBuffer(m, v49)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	if v421&int32(1) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	F_UnlockTuple(m, l0, v101, int32(8))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	F_pgstat_count_heap_delete(m, l0)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	goto L251
L253:
	;
	if v617 == int32(0) {
		v938 = v905
		goto L3
	} else {
		goto L254
	}
L254:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
	if v922&int32(1) == int32(0) {
		v938 = v905
		goto L3
	} else {
		goto L255
	}
L255:
	;
	F_pfree(m, v617)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v938 = v905
	goto L3
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
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[0]))
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
										v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
										v44 = v35
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[0]))
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
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
							v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))<<(uint(int32(6))%32))+16))
							v44 = v35
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 <= int32(1664) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < v16 {
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
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L42
	}
L4:
	;
	v74 = F_heap_compute_data_size(m, l0, l1, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v27 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v69 = int32(16)
	v71 = v5
	v73 = int32(0)
	goto L4
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v27))))
	if v34 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v69 = (int32(base.Ui32(v16+int32(7))>>(uint(int32(3))%32)) + int32(22)) & int32(536870904)
	v71 = int32(128)
	v73 = int32(1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	v48 = v27 + int32(1)
	if v48 != v16 {
		v27 = v48
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
	v78 = v74 + v69
	v80 = F_palloc0(m, v78+l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if l3 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	base.MemoryFill(m, v80, int32(0), l3)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v84 = l3 + v80
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v78
	v87 = v69 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+14)) = uint8(v87)
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+10)))
	v92 = v89&int32(_a_F_heap_form_minimal_tuple_0) | v16
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+10)) = uint16(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v84 + v69
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v71
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v101 = v84 + int32(15)
	goto L22
L21:
	;
	v101 = int32(0)
	goto L22
L22:
	;
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v105 = v101 - int32(1)
	goto L25
L24:
	;
	v105 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v105
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
	v109 = v107 & int32(_a_F_heap_form_minimal_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)) = uint16(v109)
	if int32(0) < v96 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v123 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	m.G0 = v14 + int32(32)
	return v84
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v135 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v136 = v14 + int32(24)
	goto L33
L32:
	;
	v136 = int32(0)
	goto L33
L33:
	;
	if l1 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1+v123<<(uint(int32(2))%32))))
	v146 = v144
	goto L36
L35:
	;
	v146 = int32(0)
	goto L36
L36:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v123))))
	v150 = v148
	goto L39
L38:
	;
	v150 = int32(1)
	goto L39
L39:
	;
	F_fill_val(m, l0+int32(20)+v123<<(uint(int32(4))%32), v136, v14+int32(20), v14+int32(28), v84+int32(12), v146, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v154 = v123 + int32(1)
	if v154 != v96 {
		v123 = v154
		goto L29
	} else {
		goto L41
	}
L41:
	;
	goto L30
L42:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(1664)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg(m, int32(_a_F_heap_form_minimal_tuple_2), v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_heap_form_minimal_tuple_3), int32(1473), int32(_a_F_heap_form_minimal_tuple_4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_get_root_tuples(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	base.MemoryFill(m, l1, int32(0), int32(582))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v12) < base.Ui32(int32(25)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = int32(base.Ui32(v12+int32(_a_F_heap_get_root_tuples_0))>>(uint(int32(2))%32)) & int32(_a_F_heap_get_root_tuples_1)
	if v20 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = l0 + int32(20)
	v25 = int32(1)
	v29 = v25
	v32 = v25
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24+v29<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v38)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v160 = v32 + int32(1)
	v162 = v160 & int32(_a_F_heap_get_root_tuples_1)
	if base.Ui32(v162) <= base.Ui32(v20) {
		v29 = v162
		v32 = v160
		goto L4
	} else {
		goto L36
	}
L7:
	;
	v86 = v81 & int32(_a_F_heap_get_root_tuples_1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v24+v86<<(uint(int32(2))%32))))
	if v90&int32(_a_F_heap_get_root_tuples_2) != int32(_a_F_heap_get_root_tuples_3) {
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v81 = v38 & int32(_a_F_heap_get_root_tuples_4)
	v84 = int32(0)
	goto L7
L9:
	;
	v47 = l0 + v38&int32(_a_F_heap_get_root_tuples_4)
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+18)))
	if v48 < int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v29<<(uint(int32(1))%32)-int32(2)))) = uint16(v32)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+19)))
	if v57&int32(64) == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)))
	if v62&int32(2048)|base.B2i32(v62&int32(768) == int32(512)) != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
	if v62&int32(_a_F_heap_get_root_tuples_5) == int32(_a_F_heap_get_root_tuples_6) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v75 = F_HeapTupleGetUpdateXid(m, v47)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v81 = v70
	v84 = v77
	goto L7
L16:
	;
	return
L17:
	;
	v81 = v70
	v84 = v75
	goto L7
L18:
	;
	v97 = v90
	v98 = v84
	v99 = v86
	goto L19
L19:
	;
	v105 = l0 + v97&int32(_a_F_heap_get_root_tuples_4)
	if v98 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L6
L21:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+20)))
	v107 = int32(768)
	if v106&v107 != v107 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v99<<(uint(int32(1))%32)-int32(2)))) = uint16(v32)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+19)))
	if v121&int32(64) == int32(0) {
		goto L6
	} else {
		goto L28
	}
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v113 = v111
	goto L26
L25:
	;
	v113 = int32(2)
	goto L26
L26:
	;
	if v113 != v98 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+20)))
	if v126&int32(2048)|base.B2i32(v126&int32(768) == int32(512)) != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+16)))
	if v126&int32(_a_F_heap_get_root_tuples_5) == int32(_a_F_heap_get_root_tuples_6) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v24+v134<<(uint(int32(2))%32))))
	if v146&int32(_a_F_heap_get_root_tuples_2) == int32(_a_F_heap_get_root_tuples_3) {
		v97 = v146
		v98 = v142
		v99 = v134
		goto L19
	} else {
		goto L35
	}
L31:
	;
	v139 = F_HeapTupleGetUpdateXid(m, v105)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v142 = v141
	goto L30
L34:
	;
	v142 = v139
	goto L30
L35:
	;
	goto L20
L36:
	;
	goto L5
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
					switch v38&int32(_a_F_heap_getattr_3_0) - int32(1) {
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
							F_errmsg_internal(m, int32(_a_F_heap_getattr_3_1), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_3_2), int32(70), int32(_a_F_heap_getattr_3_3))
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
			v33 = l2 + l1<<(uint(v29)%32) + v29
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
					switch v43&int32(_a_F_heap_getattr_5_0) - int32(1) {
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
							F_errmsg_internal(m, int32(_a_F_heap_getattr_5_1), v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_5_2), int32(70), int32(_a_F_heap_getattr_5_3))
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
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
	var v91 int32
	_ = v91
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
					if v14 != int32(1) {
						v39 = v6
						v45 = v6
						for {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v39))))
							if v47 == int32(1) {
								v51 = v39 << (uint(int32(2)) % 32)
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v51)))
								*(*int32)(unsafe.Add(mBase, uint32(v17+v51))) = v54
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v39))))
								*(*uint8)(unsafe.Add(mBase, uint32(v39+v21))) = uint8(v58)
							} else {
							}
							v61 = int32(1)
							v62 = v39 | v61
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
							v79 = v39 + v78
							v81 = v45 + v78
							if v81 != v14&int32(2147483646) {
								v39 = v79
								v45 = v81
								continue
							} else {
								break
							}
							break
						}
						if v14&int32(1) == int32(0) {
						} else {
							v91 = v79
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v91))))
							if v99 != int32(1) {
							} else {
								v103 = v91 << (uint(int32(2)) % 32)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l2+v103)))
								*(*int32)(unsafe.Add(mBase, uint32(v17+v103))) = v106
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v91))))
								*(*uint8)(unsafe.Add(mBase, uint32(v91+v21))) = uint8(v110)
							}
						}
					} else {
						v91 = v6
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v91))))
						if v99 != int32(1) {
						} else {
							v103 = v91 << (uint(int32(2)) % 32)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l2+v103)))
							*(*int32)(unsafe.Add(mBase, uint32(v17+v103))) = v106
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v91))))
							*(*uint8)(unsafe.Add(mBase, uint32(v91+v21))) = uint8(v110)
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
							v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)) = uint16(v133)
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v135
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
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
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v41 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v80 = F_heap_form_tuple(m, l1, v21, v25)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	v45 = v41 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3+v45)))
	if base.B2i32(v47 <= int32(0))|base.B2i32(v18 < v47) != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v52 = int32(1)
	v53 = v47 - v52
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l4+v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v53<<(uint(int32(2))%32)))) = v58
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v41))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53+v25))) = uint8(v62)
	v65 = v41 + v52
	if v65 != l2 {
		v41 = v65
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_pfree(m, v21)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v25)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)) = uint16(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v90
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)) = uint16(v92)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v96
	m.G0 = v16 + int32(16)
	return v80
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v47
	F_errmsg_internal(m, int32(_a_F_heap_modify_tuple_by_cols_0), v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_heap_modify_tuple_by_cols_1), int32(1305), int32(_a_F_heap_modify_tuple_by_cols_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v402
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v412 = F_heap_form_tuple(m, v407, v12+int32(8), v12+int32(6))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L127
	}
L2:
	;
	F_pfree(m, v68)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L126
	}
L3:
	;
	v396 = F_construct_array_builtin(m, v68, v394, int32(25))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L122
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
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L118
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
	v29 = int32(8)
	v30 = int32(base.Ui32(v18) >> (uint(v29) % 32))
	v31 = int32(255)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30&v31)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18&v31)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v42 = int32(base.Ui32(v19) >> (uint(v29) % 32))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42&v31)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v31)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v55 = v35 + (v40 + (v47 + v52))
	if v55 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v67 = v55 << (uint(int32(2)) % 32)
	v68 = F_palloc0(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v59
	v63 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v402 = v63
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
	v73 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v77 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v73
	v77 = int32(1)
	goto L20
L22:
	;
	v84 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	v89 = v77
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v77<<(uint(int32(2))%32)))) = v84
	v89 = v77 + int32(1)
	goto L24
L26:
	;
	v96 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_5))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	v101 = v89
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v89<<(uint(int32(2))%32)))) = v96
	v101 = v89 + int32(1)
	goto L28
L30:
	;
	v108 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_6))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v113 = v101
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v101<<(uint(int32(2))%32)))) = v108
	v113 = v101 + int32(1)
	goto L32
L34:
	;
	v120 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_7))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v125 = v113
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v113<<(uint(int32(2))%32)))) = v120
	v125 = v113 + int32(1)
	goto L36
L38:
	;
	v132 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v137 = v125
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v125<<(uint(int32(2))%32)))) = v132
	v137 = v125 + int32(1)
	goto L40
L42:
	;
	v144 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_9))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v149 = v137
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
	*(*int32)(unsafe.Add(mBase, uint32(v68+v137<<(uint(int32(2))%32)))) = v144
	v149 = v137 + int32(1)
	goto L44
L46:
	;
	v156 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_10))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	v161 = v149
	goto L48
L48:
	;
	if v42&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v149<<(uint(int32(2))%32)))) = v156
	v161 = v149 + int32(1)
	goto L48
L50:
	;
	v168 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_11))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	v173 = v161
	goto L52
L52:
	;
	if v42&int32(2) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v161<<(uint(int32(2))%32)))) = v168
	v173 = v161 + int32(1)
	goto L52
L54:
	;
	v180 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_12))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	v185 = v173
	goto L56
L56:
	;
	if v42&int32(4) != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v173<<(uint(int32(2))%32)))) = v180
	v185 = v173 + int32(1)
	goto L56
L58:
	;
	v192 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_13))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	v197 = v185
	goto L60
L60:
	;
	if v42&int32(8) != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v185<<(uint(int32(2))%32)))) = v192
	v197 = v185 + int32(1)
	goto L60
L62:
	;
	v204 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_14))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L65
	}
L63:
	;
	v209 = v197
	goto L64
L64:
	;
	if v42&int32(16) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v197<<(uint(int32(2))%32)))) = v204
	v209 = v197 + int32(1)
	goto L64
L66:
	;
	v216 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_15))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	v221 = v209
	goto L68
L68:
	;
	if v42&int32(32) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v209<<(uint(int32(2))%32)))) = v216
	v221 = v209 + int32(1)
	goto L68
L70:
	;
	v228 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	v233 = v221
	goto L72
L72:
	;
	if v42&int32(64) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v221<<(uint(int32(2))%32)))) = v228
	v233 = v221 + int32(1)
	goto L72
L74:
	;
	v240 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_17))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	v245 = v233
	goto L76
L76:
	;
	v247 = v42 << (uint(int32(8)) % 32)
	if base.I32_extend16_s(v247) < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v233<<(uint(int32(2))%32)))) = v240
	v245 = v233 + int32(1)
	goto L76
L78:
	;
	v255 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_18))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L81
	}
L79:
	;
	v260 = v245
	goto L80
L80:
	;
	if v30&int32(32) != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v245<<(uint(int32(2))%32)))) = v255
	v260 = v245 + int32(1)
	goto L80
L82:
	;
	v267 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_19))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	v272 = v260
	goto L84
L84:
	;
	if v30&int32(64) != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v260<<(uint(int32(2))%32)))) = v267
	v272 = v260 + int32(1)
	goto L84
L86:
	;
	v279 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_20))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L89
	}
L87:
	;
	v284 = v272
	goto L88
L88:
	;
	if v30&int32(128) != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v272<<(uint(int32(2))%32)))) = v279
	v284 = v272 + int32(1)
	goto L88
L90:
	;
	v291 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_21))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L93
	}
L91:
	;
	v296 = v284
	goto L92
L92:
	;
	v298 = F_construct_array_builtin(m, v68, v296, int32(25))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v284<<(uint(int32(2))%32)))) = v291
	v296 = v284 + int32(1)
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v298
	if v68&int32(3)|base.B2i32(base.Ui32(int32(256)) < base.Ui32(v55)) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v310 = v68 + v67
	v312 = v68 + int32(4)
	if base.Ui32(v312) < base.Ui32(v310) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v321 = v67
	goto L97
L97:
	;
	if v321 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v314 = v310
	goto L100
L99:
	;
	v314 = v312
	goto L100
L100:
	;
	v321 = (v68^int32(-1)+v314)&int32(-4) + int32(4)
	goto L97
L101:
	;
	base.MemoryFill(m, v68, int32(0), v321)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v325 = int32(80)
	if v19&v325 == v325 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v330 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_22))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L107
	}
L105:
	;
	v334 = int32(0)
	goto L106
L106:
	;
	v335 = int32(3)
	if v42&v335 == v335 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v330
	v334 = int32(1)
	goto L106
L108:
	;
	v343 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_23))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	v348 = v334
	goto L110
L110:
	;
	if base.Ui32(int32(_a_F_heap_tuple_infomask_flags_24)) <= base.Ui32(v247&int32(_a_F_heap_tuple_infomask_flags_25)) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v334<<(uint(int32(2))%32)))) = v343
	v348 = v334 + int32(1)
	goto L110
L112:
	;
	v357 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_26))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v348 != 0 {
		v394 = v348
		goto L3
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v348<<(uint(int32(2))%32)))) = v357
	v394 = v348 + int32(1)
	goto L3
L116:
	;
	v363 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v399 = v363
	goto L2
L118:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_heap_tuple_infomask_flags_27), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_heap_tuple_infomask_flags_1), int32(537), int32(_a_F_heap_tuple_infomask_flags_2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errmsg_internal(m, int32(_a_F_heap_tuple_infomask_flags_0), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_heap_tuple_infomask_flags_1), int32(541), int32(_a_F_heap_tuple_infomask_flags_2))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	v399 = v396
	goto L2
L126:
	;
	v402 = v399
	goto L1
L127:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v412)+16))
	v415 = F_HeapTupleHeaderGetDatum(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	m.G0 = v12 + int32(16)
	return v415
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
		if v4&int32(_a_F_heap_tuple_needs_eventual_freeze_0) != 0 {
			if v14 == int32(0) {
				if base.Ui32(v4) < base.Ui32(int32(_a_F_heap_tuple_needs_eventual_freeze_1)) {
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
				if base.Ui32(v4) < base.Ui32(int32(_a_F_heap_tuple_needs_eventual_freeze_1)) {
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
			if v4&int32(_a_F_heap_tuple_needs_eventual_freeze_0) != 0 {
				if v14 == int32(0) {
					if base.Ui32(v4) < base.Ui32(int32(_a_F_heap_tuple_needs_eventual_freeze_1)) {
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
					if base.Ui32(v4) < base.Ui32(int32(_a_F_heap_tuple_needs_eventual_freeze_1)) {
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
