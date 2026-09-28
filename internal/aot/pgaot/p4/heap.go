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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			F_pfree(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v16
				v18 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v18)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v28 = v17 & int32(_a_F_ExecStoreHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v28)
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v32
				if l2 != 0 {
					v35 = v28 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v35)
				} else {
				}
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v37
				return l1
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v18)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v28 = v17 & int32(_a_F_ExecStoreHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v28)
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v32
			if l2 != 0 {
				v35 = v28 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v35)
			} else {
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v37
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
				F_errfinish(m, int32(_a_F_ExecStoreHeapTuple_3), int32(1646), int32(_a_F_ExecStoreHeapTuple_4))
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int64
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
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
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
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
		goto L12
	}
L5:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[1]))
	if v134 == v143 {
		goto L1
	} else {
		goto L41
	}
L6:
	;
	if base.Ui32(v123) < base.Ui32(v124) {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v134 = v120
	goto L5
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L37
	}
L9:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+20)))
	v101 = int32(768)
	if v100&v101 != v101 {
		v119 = v99
		goto L7
	} else {
		goto L36
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)))
	if l0 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	if l0 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	switch v21 {
	case 0:
		goto L1
	case 1:
		goto L11
	case 2, 4:
		goto L10
	case 3:
		goto L9
	default:
		goto L8
	}
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+20)))
	v25 = int32(768)
	if v24&v25 != v25 {
		v119 = v23
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v134 = int32(2)
	goto L5
L15:
	;
	v89 = int32(3)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
	if base.B2i32(base.Ui32(v81) < base.Ui32(v89))|base.B2i32(base.Ui32(v92) < base.Ui32(v89)) != 0 {
		v123 = v81
		v124 = v92
		goto L6
	} else {
		goto L34
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31&int32(_a_F_HeapCheckForSerializableConflictOut_0) != int32(_a_F_HeapCheckForSerializableConflictOut_1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v71 = int32(768)
	if v31&v71 == v71 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v81 = v32
	goto L15
L20:
	;
	goto L21
L21:
	;
	v37 = int32(0)
	v41 = F_GetMultiXactIdMembers(m, v32, v13+int32(12), v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v41 <= int32(0) {
		v81 = v37
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v48 = v37
	goto L25
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	F_pfree(m, v45)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L30
	}
L25:
	;
	v58 = v45 + v48<<(uint(int32(3))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L24
	} else {
		goto L27
	}
L26:
	;
	F_pfree(m, v45)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	v63 = v48 + int32(1)
	if v63 != v41 {
		v48 = v63
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v81 = int32(0)
	goto L15
L30:
	;
	v81 = v68
	goto L15
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
	v123 = int32(2)
	v124 = v76
	goto L6
L32:
	;
	goto L33
L33:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v81 = v78
	goto L15
L34:
	;
	if int32(0) <= v81-v92 {
		v134 = v81
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v134 = int32(2)
	goto L5
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
	F_errmsg_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_2), v13)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_3), int32(_a_F_HeapCheckForSerializableConflictOut_4), int32(_a_F_HeapCheckForSerializableConflictOut_5))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
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
	v134 = v123
	goto L5
L41:
	;
	v145 = F_SubTransGetTopmostTransaction(m, v134)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	v160 = m.G0
	v162 = v160 - int32(32)
	m.G0 = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	if v165 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v147 = int32(3)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[0]))
	if base.B2i32(base.Ui32(v145) < base.Ui32(v147))|base.B2i32(base.Ui32(v150) < base.Ui32(v147)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if int32(0) <= v145-v150 {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v145) < base.Ui32(v150) {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L1
L48:
	;
	goto L42
L49:
	;
	goto L1
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L2
	} else {
		goto L182
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L2
	} else {
		goto L176
	}
L52:
	;
	m.G0 = v162 + int32(32)
	goto L49
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v168 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	if v169&int32(128) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui32(v176) < base.Ui32(int32(_a_F_HeapCheckForSerializableConflictOut_6)) {
		goto L52
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+118)))
	if v180 == int32(116) {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	if v169&int32(8) != 0 {
		goto L51
	} else {
		goto L61
	}
L61:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[1]))
	if v145 == v186 {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+24)) = v145
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	v194 = F_LWLockAcquire(m, v190+int32(3584), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[4]))
	v200 = int32(0)
	v202 = F_hash_search(m, v197, v162+int32(24), v200, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	if v202 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+28)) = v145
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	v212 = F_LWLockAcquire(m, v208+int32(_a_F_HeapCheckForSerializableConflictOut_7), int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	if v346 != v348 {
		goto L109
	} else {
		goto L110
	}
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[5]))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v219+int32(_a_F_HeapCheckForSerializableConflictOut_7))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	if v217 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v341+int32(3584))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L107
	}
L71:
	;
	v226 = int32(3)
	if base.B2i32(base.Ui32(v145) < base.Ui32(v226))|base.B2i32(base.Ui32(v216) < base.Ui32(v226)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v237 = int32(3)
	if base.B2i32(base.Ui32(v145) < base.Ui32(v237))|base.B2i32(base.Ui32(v217) < base.Ui32(v237)) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	if int32(0) <= v145-v216 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v145) < base.Ui32(v216) {
		goto L70
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	goto L72
L78:
	;
	v254 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_HeapCheckForSerializableConflictOut_8), base.I64_extend_i32_u(int32(base.Ui32(v145)>>(uint(int32(10))%32))), v162+int32(28))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L84
	}
L79:
	;
	if v145-v217 <= int32(0) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(v217) < base.Ui32(v145) {
		goto L70
	} else {
		goto L83
	}
L82:
	;
	goto L70
L83:
	;
	goto L78
L84:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[6]))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258+v254<<(uint(int32(2))%32))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v162)+28))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v262+v263<<(uint(int32(3))%32)&int32(_a_F_HeapCheckForSerializableConflictOut_9))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v257)+28))
	v274 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[7])))
	v275 = base.I32_rem_u_s(int32(base.Ui32(v263)>>(uint(int32(10))%32)), v274)
	F_LWLockRelease(m, v270+v275<<(uint(int32(7))%32))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v281 = int64(1)
	v282 = v269 + v281
	if base.Ui64(v282) <= base.Ui64(v281) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if v322&int32(512) != 0 {
		goto L50
	} else {
		goto L102
	}
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+108))
	v322 = v321
	v323 = v320
	goto L86
L88:
	;
	if base.I32_wrap_i64(v282)-int32(1) != 0 {
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+108))
	if v290&int32(32) != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L70
L92:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v289)+24))
	if base.Ui64(v293) < base.Ui64(v269) {
		v322 = v290
		v323 = v289
		goto L86
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v145
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_11), v162)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4001), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v323)+44))
	if v326 != v323+int32(40) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v331 = v326
	goto L105
L104:
	;
	v331 = int32(0)
	goto L105
L105:
	;
	if v331 != 0 {
		goto L50
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+108)) = v322 | int32(1024)
	goto L70
L107:
	;
	goto L52
L108:
	;
	if v350&int32(1024) != 0 {
		goto L114
	} else {
		goto L115
	}
L109:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346)+108))
	if v350&int32(8) == int32(0) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v357+int32(3584))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	goto L52
L114:
	;
	if v350&int32(2) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	if v350&int32(1) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+108)) = v350 | int32(8)
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v372+int32(3584))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v378+int32(3584))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L2
	} else {
		goto L121
	}
L120:
	;
	goto L52
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_15), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4048), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	v427 = int32(0)
	v430 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L2
	} else {
		goto L139
	}
L129:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v348)+108))
	if v411&int32(32) == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	if v350&int32(16) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v348)+24))
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v346)+24))
	if base.Ui64(v419) <= base.Ui64(v418) {
		goto L128
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v422+int32(3584))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L2
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	goto L52
L136:
	;
	if v500 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L137:
	;
	v500 = v485
	goto L136
L138:
	;
	v443 = int32(3)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	if base.B2i32(base.Ui32(v145) < base.Ui32(v443))|base.B2i32(base.Ui32(v445) < base.Ui32(v443)) == int32(0) {
		goto L146
	} else {
		goto L147
	}
L139:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if base.B2i32(base.Ui32(v145) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v432) < base.Ui32(int32(3))) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v145-v432 < int32(0) {
		v485 = v427
		goto L137
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if base.Ui32(v432) <= base.Ui32(v145) {
		goto L138
	} else {
		goto L144
	}
L143:
	;
	goto L138
L144:
	;
	v500 = int32(0)
	goto L136
L145:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v430)+16))
	if v457 == int32(0) {
		v485 = v427
		goto L137
	} else {
		goto L151
	}
L146:
	;
	if v145-v445 < int32(0) {
		goto L145
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if base.Ui32(v145) < base.Ui32(v445) {
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v500 = int32(1)
	goto L136
L150:
	;
	v500 = int32(1)
	goto L136
L151:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v462 = int32(0)
	goto L152
L152:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v460+v462<<(uint(int32(2))%32))))
	v476 = base.B2i32(v145 == v475)
	if v145 == v475 {
		v485 = v476
		goto L137
	} else {
		goto L154
	}
L153:
	;
	v485 = v476
	goto L137
L154:
	;
	v478 = v462 + int32(1)
	if v478 != v457 {
		v462 = v478
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v504+int32(3584))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L2
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v509 = int32(0)
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[2]))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+108)))
	if v512&int32(8) != 0 {
		v550 = v509
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L52
L160:
	;
	if v550 != 0 {
		goto L170
	} else {
		goto L171
	}
L161:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+108)))
	if v515&int32(8) != 0 {
		v550 = v509
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v511)+36))
	if v518 == int32(0) {
		v550 = v509
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v522 = v511 + int32(32)
	if v518 == v522 {
		v550 = v509
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v346)+44))
	if base.B2i32(v524 == int32(0))|base.B2i32(v524 == v346+int32(40)) != 0 {
		v550 = v509
		goto L160
	} else {
		goto L165
	}
L165:
	;
	v532 = v518
	goto L166
L166:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	v542 = base.B2i32(v541 == v346)
	if v541 == v346 {
		v550 = v542
		goto L160
	} else {
		goto L168
	}
L167:
	;
	v550 = v542
	goto L160
L168:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v543 != v522 {
		v532 = v543
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v556+int32(3584))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_FlagRWConflict(m, v511, v346)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L174
	}
L173:
	;
	goto L52
L174:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_HeapCheckForSerializableConflictOut[3]))
	F_LWLockRelease(m, v564+int32(3584))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	goto L52
L176:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_16), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(3968), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L2
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_HeapCheckForSerializableConflictOut_10), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v145
	F_errdetail_internal(m, int32(_a_F_HeapCheckForSerializableConflictOut_17), v162+int32(16))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	F_errhint(m, int32(_a_F_HeapCheckForSerializableConflictOut_12), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_HeapCheckForSerializableConflictOut_13), int32(4009), int32(_a_F_HeapCheckForSerializableConflictOut_14))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
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
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
		v94 = v69
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
	if v94&int32(256) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = int32(3)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(v79) < base.Ui32(v80))|base.B2i32(base.Ui32(v82) < base.Ui32(v80)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v79
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v94 = v93
	goto L20
L23:
	;
	if v82-v79 < int32(0) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v79) <= base.Ui32(v82) {
		v94 = v69
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v94 = v69
	goto L20
L27:
	;
	goto L22
L28:
	;
	m.G0 = v11 + int32(16)
	return
L29:
	;
	v109 = int32(3)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(v71) < base.Ui32(v109))|base.B2i32(base.Ui32(v111) < base.Ui32(v109)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if v94&int32(512) != 0 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v71 == v20 {
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v103 = F_TransactionIdDidCommit(m, v20)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(0) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v71 != v20 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	goto L29
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71
	goto L28
L39:
	;
	if int32(0) < v71-v111 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v71) <= base.Ui32(v111) {
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L28
L43:
	;
	goto L38
}
func F_HeapTupleSatisfiesMVCC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v758 int32
	_ = v758
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	if v9&int32(256) == v5 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_SetHintBitsExt(m, v8, l2, int32(2048), int32(0), l3)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L9
	} else {
		goto L267
	}
L2:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	if v795&int32(32) != 0 {
		goto L264
	} else {
		goto L265
	}
L3:
	;
	return v786
L4:
	;
	v477 = int32(1)
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	if v478&int32(2048)|v478&int32(128)|base.B2i32(v478&int32(_a_F_HeapTupleSatisfiesMVCC_0) == int32(64)) != 0 {
		v786 = v477
		goto L3
	} else {
		goto L160
	}
L5:
	;
	if v9&int32(512) != 0 {
		v786 = v5
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v470 = int32(768)
	if v9&v470 == v470 {
		goto L4
	} else {
		goto L157
	}
L8:
	;
	v16 = F_HeapTupleCleanMoved(m, v8, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v16 == int32(0) {
		v786 = v5
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui32(v22) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v154 != 0 {
		goto L52
	} else {
		goto L53
	}
L13:
	;
	v154 = int32(0)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[0]))
	if v34 == v22 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v154 = int32(1)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[1]))
	if v38 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v154 = v144
	goto L12
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[2]))
	if v42 == int32(0) {
		v144 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[3]))
	v114 = int32(0)
	v117 = v38 - int32(1)
	goto L42
L23:
	;
	v47 = v42
	goto L24
L24:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	if v53 == int32(4) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v144 = int32(0)
	goto L19
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	if v107 != 0 {
		v47 = v107
		goto L24
	} else {
		goto L41
	}
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v56 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v59 = int32(1)
	if v22 == v56 {
		v144 = v59
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v69 = int32(0)
	v72 = v63
	goto L31
L31:
	;
	v77 = int32(2)
	v78 = base.I32_div_s(v72-v69, v77)
	v79 = v78 + v69
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v66+v79<<(uint(v77)%32))))
	if v83 == v22 {
		v144 = v59
		goto L19
	} else {
		goto L33
	}
L32:
	;
	goto L26
L33:
	;
	v92 = base.B2i32(v83-v22 < int32(0)) | base.B2i32(base.Ui32(v83) < base.Ui32(int32(3)))
	if v92 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = v79 + int32(1)
	goto L36
L35:
	;
	v93 = v69
	goto L36
L36:
	;
	if v92 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v96 = v72
	goto L39
L38:
	;
	v96 = v79 - int32(1)
	goto L39
L39:
	;
	if v93 <= v96 {
		v69 = v93
		v72 = v96
		goto L31
	} else {
		goto L40
	}
L40:
	;
	goto L32
L41:
	;
	goto L25
L42:
	;
	v122 = int32(2)
	v123 = base.I32_div_s(v117-v114, v122)
	v124 = v123 + v114
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v112+v124<<(uint(v122)%32))))
	v129 = base.B2i32(v128 == v22)
	if v128 == v22 {
		v144 = v129
		goto L19
	} else {
		goto L44
	}
L43:
	;
	v144 = v129
	goto L19
L44:
	;
	v132 = base.B2i32(base.Ui32(v128) < base.Ui32(v22))
	if base.Ui32(v128) < base.Ui32(v22) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = v124 + int32(1)
	goto L47
L46:
	;
	v133 = v114
	goto L47
L47:
	;
	if base.Ui32(v128) < base.Ui32(v22) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v136 = v117
	goto L50
L49:
	;
	v136 = v124 - int32(1)
	goto L50
L50:
	;
	if v133 <= v136 {
		v114 = v133
		v117 = v136
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	if v157&int32(32) != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v455 = F_XidInMVCCSnapshot(m, v454, l1)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L149
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v167) <= base.Ui32(v166) {
		v786 = v5
		goto L3
	} else {
		goto L59
	}
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[4]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v156<<(uint(int32(3))%32))))
	v166 = v165
	goto L58
L57:
	;
	v166 = v156
	goto L58
L58:
	;
	goto L55
L59:
	;
	v169 = int32(1)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	if v170&int32(2048)|v170&int32(128)|base.B2i32(v170&int32(_a_F_HeapTupleSatisfiesMVCC_0) == int32(64)) != 0 {
		v786 = v169
		goto L3
	} else {
		goto L60
	}
L60:
	;
	if v170&int32(_a_F_HeapTupleSatisfiesMVCC_1) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v183 = F_HeapTupleGetUpdateXid(m, v8)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.Ui32(v319) < base.Ui32(int32(3)) {
		goto L107
	} else {
		goto L108
	}
L64:
	;
	if base.Ui32(v183) < base.Ui32(int32(3)) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v316 == int32(0) {
		v786 = v169
		goto L3
	} else {
		goto L105
	}
L66:
	;
	v316 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[0]))
	if v196 == v183 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v316 = int32(1)
	goto L65
L70:
	;
	goto L71
L71:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[1]))
	if v200 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v316 = v306
	goto L65
L73:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[2]))
	if v204 == int32(0) {
		v306 = int32(0)
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[3]))
	v276 = int32(0)
	v279 = v200 - int32(1)
	goto L95
L76:
	;
	v209 = v204
	goto L77
L77:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	if v215 == int32(4) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v306 = int32(0)
	goto L72
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v209)+80))
	if v269 != 0 {
		v209 = v269
		goto L77
	} else {
		goto L94
	}
L80:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v218 == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v221 = int32(1)
	if v183 == v218 {
		v306 = v221
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v209)+52))
	v225 = v223 - int32(1)
	if v225 < int32(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	v231 = int32(0)
	v234 = v225
	goto L84
L84:
	;
	v239 = int32(2)
	v240 = base.I32_div_s(v234-v231, v239)
	v241 = v240 + v231
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v228+v241<<(uint(v239)%32))))
	if v245 == v183 {
		v306 = v221
		goto L72
	} else {
		goto L86
	}
L85:
	;
	goto L79
L86:
	;
	v254 = base.B2i32(v245-v183 < int32(0)) | base.B2i32(base.Ui32(v245) < base.Ui32(int32(3)))
	if v254 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v255 = v241 + int32(1)
	goto L89
L88:
	;
	v255 = v231
	goto L89
L89:
	;
	if v254 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v258 = v234
	goto L92
L91:
	;
	v258 = v241 - int32(1)
	goto L92
L92:
	;
	if v255 <= v258 {
		v231 = v255
		v234 = v258
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
	v284 = int32(2)
	v285 = base.I32_div_s(v279-v276, v284)
	v286 = v285 + v276
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v274+v286<<(uint(v284)%32))))
	v291 = base.B2i32(v290 == v183)
	if v290 == v183 {
		v306 = v291
		goto L72
	} else {
		goto L97
	}
L96:
	;
	v306 = v291
	goto L72
L97:
	;
	v294 = base.B2i32(base.Ui32(v290) < base.Ui32(v183))
	if base.Ui32(v290) < base.Ui32(v183) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v295 = v286 + int32(1)
	goto L100
L99:
	;
	v295 = v276
	goto L100
L100:
	;
	if base.Ui32(v290) < base.Ui32(v183) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v298 = v279
	goto L103
L102:
	;
	v298 = v286 - int32(1)
	goto L103
L103:
	;
	if v295 <= v298 {
		v276 = v295
		v279 = v298
		goto L95
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	goto L2
L106:
	;
	if v451 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L107:
	;
	v451 = int32(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[0]))
	if v331 == v319 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v451 = int32(1)
	goto L106
L111:
	;
	goto L112
L112:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[1]))
	if v335 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v451 = v441
	goto L106
L114:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[2]))
	if v339 == int32(0) {
		v441 = int32(0)
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[3]))
	v411 = int32(0)
	v414 = v335 - int32(1)
	goto L136
L117:
	;
	v344 = v339
	goto L118
L118:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	if v350 == int32(4) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v441 = int32(0)
	goto L113
L120:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v344)+80))
	if v404 != 0 {
		v344 = v404
		goto L118
	} else {
		goto L135
	}
L121:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v353 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v356 = int32(1)
	if v319 == v353 {
		v441 = v356
		goto L113
	} else {
		goto L123
	}
L123:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
	v360 = v358 - int32(1)
	if v360 < int32(0) {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v344)+48))
	v366 = int32(0)
	v369 = v360
	goto L125
L125:
	;
	v374 = int32(2)
	v375 = base.I32_div_s(v369-v366, v374)
	v376 = v375 + v366
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v363+v376<<(uint(v374)%32))))
	if v380 == v319 {
		v441 = v356
		goto L113
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v389 = base.B2i32(v380-v319 < int32(0)) | base.B2i32(base.Ui32(v380) < base.Ui32(int32(3)))
	if v389 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v390 = v376 + int32(1)
	goto L130
L129:
	;
	v390 = v366
	goto L130
L130:
	;
	if v389 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v393 = v369
	goto L133
L132:
	;
	v393 = v376 - int32(1)
	goto L133
L133:
	;
	if v390 <= v393 {
		v366 = v390
		v369 = v393
		goto L125
	} else {
		goto L134
	}
L134:
	;
	goto L126
L135:
	;
	goto L119
L136:
	;
	v419 = int32(2)
	v420 = base.I32_div_s(v414-v411, v419)
	v421 = v420 + v411
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v409+v421<<(uint(v419)%32))))
	v426 = base.B2i32(v425 == v319)
	if v425 == v319 {
		v441 = v426
		goto L113
	} else {
		goto L138
	}
L137:
	;
	v441 = v426
	goto L113
L138:
	;
	v429 = base.B2i32(base.Ui32(v425) < base.Ui32(v319))
	if base.Ui32(v425) < base.Ui32(v319) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v430 = v421 + int32(1)
	goto L141
L140:
	;
	v430 = v411
	goto L141
L141:
	;
	if base.Ui32(v425) < base.Ui32(v319) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v433 = v414
	goto L144
L143:
	;
	v433 = v421 - int32(1)
	goto L144
L144:
	;
	if v430 <= v433 {
		v411 = v430
		v414 = v433
		goto L136
	} else {
		goto L145
	}
L145:
	;
	goto L137
L146:
	;
	goto L1
L147:
	;
	goto L148
L148:
	;
	goto L2
L149:
	;
	if v455 != 0 {
		v786 = v5
		goto L3
	} else {
		goto L150
	}
L150:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v458 = F_TransactionIdDidCommit(m, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	if v458 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	F_SetHintBitsExt(m, v8, l2, int32(256), v461, l3)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	F_SetHintBitsExt(m, v8, l2, int32(512), int32(0), l3)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L9
	} else {
		goto L156
	}
L155:
	;
	goto L4
L156:
	;
	return int32(0)
L157:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v475 = F_XidInMVCCSnapshot(m, v474, l1)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	if v475 != 0 {
		v786 = v5
		goto L3
	} else {
		goto L159
	}
L159:
	;
	goto L4
L160:
	;
	if v478&int32(_a_F_HeapTupleSatisfiesMVCC_1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v491 = F_HeapTupleGetUpdateXid(m, v8)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L9
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v478&int32(1024) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L164:
	;
	if base.Ui32(v491) < base.Ui32(int32(3)) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if v624 != 0 {
		goto L2
	} else {
		goto L205
	}
L166:
	;
	v624 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[0]))
	if v504 == v491 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v624 = int32(1)
	goto L165
L170:
	;
	goto L171
L171:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[1]))
	if v508 <= int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v624 = v614
	goto L165
L173:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[2]))
	if v512 == int32(0) {
		v614 = int32(0)
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[3]))
	v584 = int32(0)
	v587 = v508 - int32(1)
	goto L195
L176:
	;
	v517 = v512
	goto L177
L177:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v517)+20))
	if v523 == int32(4) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v614 = int32(0)
	goto L172
L179:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v517)+80))
	if v577 != 0 {
		v517 = v577
		goto L177
	} else {
		goto L194
	}
L180:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	if v526 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v529 = int32(1)
	if v491 == v526 {
		v614 = v529
		goto L172
	} else {
		goto L182
	}
L182:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v517)+52))
	v533 = v531 - int32(1)
	if v533 < int32(0) {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v517)+48))
	v539 = int32(0)
	v542 = v533
	goto L184
L184:
	;
	v547 = int32(2)
	v548 = base.I32_div_s(v542-v539, v547)
	v549 = v548 + v539
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v536+v549<<(uint(v547)%32))))
	if v553 == v491 {
		v614 = v529
		goto L172
	} else {
		goto L186
	}
L185:
	;
	goto L179
L186:
	;
	v562 = base.B2i32(v553-v491 < int32(0)) | base.B2i32(base.Ui32(v553) < base.Ui32(int32(3)))
	if v562 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v563 = v549 + int32(1)
	goto L189
L188:
	;
	v563 = v539
	goto L189
L189:
	;
	if v562 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v566 = v542
	goto L192
L191:
	;
	v566 = v549 - int32(1)
	goto L192
L192:
	;
	if v563 <= v566 {
		v539 = v563
		v542 = v566
		goto L184
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	goto L178
L195:
	;
	v592 = int32(2)
	v593 = base.I32_div_s(v587-v584, v592)
	v594 = v593 + v584
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v582+v594<<(uint(v592)%32))))
	v599 = base.B2i32(v598 == v491)
	if v598 == v491 {
		v614 = v599
		goto L172
	} else {
		goto L197
	}
L196:
	;
	v614 = v599
	goto L172
L197:
	;
	v602 = base.B2i32(base.Ui32(v598) < base.Ui32(v491))
	if base.Ui32(v598) < base.Ui32(v491) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v603 = v594 + int32(1)
	goto L200
L199:
	;
	v603 = v584
	goto L200
L200:
	;
	if base.Ui32(v598) < base.Ui32(v491) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v606 = v587
	goto L203
L202:
	;
	v606 = v594 - int32(1)
	goto L203
L203:
	;
	if v603 <= v606 {
		v584 = v603
		v587 = v606
		goto L195
	} else {
		goto L204
	}
L204:
	;
	goto L196
L205:
	;
	v625 = F_XidInMVCCSnapshot(m, v491, l1)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L9
	} else {
		goto L206
	}
L206:
	;
	if v625 != 0 {
		v786 = v477
		goto L3
	} else {
		goto L207
	}
L207:
	;
	v627 = F_TransactionIdDidCommit(m, v491)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L9
	} else {
		goto L208
	}
L208:
	;
	return v627 ^ int32(1)
L209:
	;
	v786 = int32(0)
	goto L3
L210:
	;
	if base.Ui32(v632) < base.Ui32(int32(3)) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	v781 = F_XidInMVCCSnapshot(m, v632, l1)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L9
	} else {
		goto L261
	}
L213:
	;
	if v768 != 0 {
		goto L2
	} else {
		goto L253
	}
L214:
	;
	v768 = int32(0)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v648 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[0]))
	if v648 == v632 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v768 = int32(1)
	goto L213
L218:
	;
	goto L219
L219:
	;
	v652 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[1]))
	if v652 <= int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v768 = v758
	goto L213
L221:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[2]))
	if v656 == int32(0) {
		v758 = int32(0)
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[3]))
	v728 = int32(0)
	v731 = v652 - int32(1)
	goto L243
L224:
	;
	v661 = v656
	goto L225
L225:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661)+20))
	if v667 == int32(4) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v758 = int32(0)
	goto L220
L227:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v661)+80))
	if v721 != 0 {
		v661 = v721
		goto L225
	} else {
		goto L242
	}
L228:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	if v670 == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v673 = int32(1)
	if v632 == v670 {
		v758 = v673
		goto L220
	} else {
		goto L230
	}
L230:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v661)+52))
	v677 = v675 - int32(1)
	if v677 < int32(0) {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v661)+48))
	v683 = int32(0)
	v686 = v677
	goto L232
L232:
	;
	v691 = int32(2)
	v692 = base.I32_div_s(v686-v683, v691)
	v693 = v692 + v683
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v680+v693<<(uint(v691)%32))))
	if v697 == v632 {
		v758 = v673
		goto L220
	} else {
		goto L234
	}
L233:
	;
	goto L227
L234:
	;
	v706 = base.B2i32(v697-v632 < int32(0)) | base.B2i32(base.Ui32(v697) < base.Ui32(int32(3)))
	if v706 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v707 = v693 + int32(1)
	goto L237
L236:
	;
	v707 = v683
	goto L237
L237:
	;
	if v706 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v710 = v686
	goto L240
L239:
	;
	v710 = v693 - int32(1)
	goto L240
L240:
	;
	if v707 <= v710 {
		v683 = v707
		v686 = v710
		goto L232
	} else {
		goto L241
	}
L241:
	;
	goto L233
L242:
	;
	goto L226
L243:
	;
	v736 = int32(2)
	v737 = base.I32_div_s(v731-v728, v736)
	v738 = v737 + v728
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v726+v738<<(uint(v736)%32))))
	v743 = base.B2i32(v742 == v632)
	if v742 == v632 {
		v758 = v743
		goto L220
	} else {
		goto L245
	}
L244:
	;
	v758 = v743
	goto L220
L245:
	;
	v746 = base.B2i32(base.Ui32(v742) < base.Ui32(v632))
	if base.Ui32(v742) < base.Ui32(v632) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v747 = v738 + int32(1)
	goto L248
L247:
	;
	v747 = v728
	goto L248
L248:
	;
	if base.Ui32(v742) < base.Ui32(v632) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v750 = v731
	goto L251
L250:
	;
	v750 = v738 - int32(1)
	goto L251
L251:
	;
	if v747 <= v750 {
		v728 = v747
		v731 = v750
		goto L243
	} else {
		goto L252
	}
L252:
	;
	goto L244
L253:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v770 = F_XidInMVCCSnapshot(m, v769, l1)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	if v770 != 0 {
		v786 = v477
		goto L3
	} else {
		goto L255
	}
L255:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v773 = F_TransactionIdDidCommit(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	if v773 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	goto L1
L258:
	;
	goto L259
L259:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F_SetHintBitsExt(m, v8, l2, int32(1024), v778, l3)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L9
	} else {
		goto L260
	}
L260:
	;
	goto L209
L261:
	;
	if v781 != 0 {
		v786 = v477
		goto L3
	} else {
		goto L262
	}
L262:
	;
	goto L209
L263:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	return base.B2i32(base.Ui32(v805) <= base.Ui32(v804))
L264:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesMVCC[4]))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v799+v794<<(uint(int32(3))%32))+4))
	v804 = v803
	goto L266
L265:
	;
	v804 = v794
	goto L266
L266:
	;
	goto L263
L267:
	;
	return int32(1)
}
func F_HeapTupleSatisfiesUpdate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
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
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = v8 + int32(20)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	if v11&int32(256) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_BufferSetHintBits16(m, v10, v827|int32(2048), l2)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L11
	} else {
		goto L327
	}
L2:
	;
	v913 = int32(4)
	v916 = l0 + v913
	v918 = v8 + int32(12)
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916)+2)))
	v920 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916))))
	v921 = int32(16)
	v924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+2)))
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918))))
	if v919|v920<<(uint(v921)%32) == v924|v925<<(uint(v921)%32) {
		goto L320
	} else {
		goto L321
	}
L3:
	;
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	F_BufferSetHintBits16(m, v10, v904|int32(2048), l2)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L11
	} else {
		goto L317
	}
L4:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	if v888&int32(32) != 0 {
		goto L311
	} else {
		goto L312
	}
L5:
	;
	v874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	F_BufferSetHintBits16(m, v10, v874|int32(512), l2)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L11
	} else {
		goto L309
	}
L6:
	;
	return v872
L7:
	;
	v16 = int32(1)
	if v11&int32(512) != 0 {
		v872 = v16
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v496 = v11
	goto L9
L9:
	;
	v498 = int32(0)
	if v496&int32(2048) != 0 {
		v872 = v498
		goto L6
	} else {
		goto L176
	}
L10:
	;
	v19 = F_HeapTupleCleanMoved(m, v8, l2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v19 == int32(0) {
		v872 = v16
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui32(v25) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v157 != 0 {
		goto L54
	} else {
		goto L55
	}
L15:
	;
	v157 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v37 == v25 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v157 = int32(1)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v41 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v157 = v147
	goto L14
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v45 == int32(0) {
		v147 = int32(0)
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v117 = int32(0)
	v120 = v41 - int32(1)
	goto L44
L25:
	;
	v50 = v45
	goto L26
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v56 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v147 = int32(0)
	goto L21
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)+80))
	if v110 != 0 {
		v50 = v110
		goto L26
	} else {
		goto L43
	}
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v59 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v62 = int32(1)
	if v25 == v59 {
		v147 = v62
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v66 = v64 - int32(1)
	if v66 < int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v72 = int32(0)
	v75 = v66
	goto L33
L33:
	;
	v80 = int32(2)
	v81 = base.I32_div_s(v75-v72, v80)
	v82 = v81 + v72
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69+v82<<(uint(v80)%32))))
	if v86 == v25 {
		v147 = v62
		goto L21
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v95 = base.B2i32(v86-v25 < int32(0)) | base.B2i32(base.Ui32(v86) < base.Ui32(int32(3)))
	if v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v96 = v82 + int32(1)
	goto L38
L37:
	;
	v96 = v72
	goto L38
L38:
	;
	if v95 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v99 = v75
	goto L41
L40:
	;
	v99 = v82 - int32(1)
	goto L41
L41:
	;
	if v96 <= v99 {
		v72 = v96
		v75 = v99
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
	v125 = int32(2)
	v126 = base.I32_div_s(v120-v117, v125)
	v127 = v126 + v117
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115+v127<<(uint(v125)%32))))
	v132 = base.B2i32(v131 == v25)
	if v131 == v25 {
		v147 = v132
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v147 = v132
	goto L21
L46:
	;
	v135 = base.B2i32(base.Ui32(v131) < base.Ui32(v25))
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v136 = v127 + int32(1)
	goto L49
L48:
	;
	v136 = v117
	goto L49
L49:
	;
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v139 = v120
	goto L52
L51:
	;
	v139 = v127 - int32(1)
	goto L52
L52:
	;
	if v136 <= v139 {
		v117 = v136
		v120 = v139
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	if v160&int32(32) != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v484 = F_TransactionIdIsInProgress(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L171
	}
L57:
	;
	if base.Ui32(l1) <= base.Ui32(v169) {
		v872 = v16
		goto L6
	} else {
		goto L61
	}
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[4]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164+v159<<(uint(int32(3))%32))))
	v169 = v168
	goto L60
L59:
	;
	v169 = v159
	goto L60
L60:
	;
	goto L57
L61:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	if v172&int32(2048) != 0 {
		v872 = int32(0)
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v177 = int32(0)
	if base.B2i32(v172&int32(128) == v177)&base.B2i32(v172&int32(_a_F_HeapTupleSatisfiesUpdate_0) != int32(64)) == v177 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v172&int32(_a_F_HeapTupleSatisfiesUpdate_1) != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if v172&int32(_a_F_HeapTupleSatisfiesUpdate_1) != 0 {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v192 = F_MultiXactIdIsRunning(m, v186, int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v198 = F_TransactionIdIsInProgress(m, v186)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L11
	} else {
		goto L73
	}
L69:
	;
	if v192 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v194 = int32(5)
	goto L72
L71:
	;
	v194 = int32(0)
	goto L72
L72:
	;
	return v194
L73:
	;
	if v198 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v200 = int32(5)
	goto L76
L75:
	;
	v200 = int32(0)
	goto L76
L76:
	;
	return v200
L77:
	;
	v204 = F_HeapTupleGetUpdateXid(m, v8)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.Ui32(v348) < base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L80:
	;
	if base.Ui32(v204) < base.Ui32(int32(3)) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v337 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L82:
	;
	v337 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v217 == v204 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v337 = int32(1)
	goto L81
L86:
	;
	goto L87
L87:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v221 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v337 = v327
	goto L81
L89:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v225 == int32(0) {
		v327 = int32(0)
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v297 = int32(0)
	v300 = v221 - int32(1)
	goto L111
L92:
	;
	v230 = v225
	goto L93
L93:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v236 == int32(4) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v327 = int32(0)
	goto L88
L95:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v230)+80))
	if v290 != 0 {
		v230 = v290
		goto L93
	} else {
		goto L110
	}
L96:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v239 == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v242 = int32(1)
	if v204 == v239 {
		v327 = v242
		goto L88
	} else {
		goto L98
	}
L98:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v230)+52))
	v246 = v244 - int32(1)
	if v246 < int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v230)+48))
	v252 = int32(0)
	v255 = v246
	goto L100
L100:
	;
	v260 = int32(2)
	v261 = base.I32_div_s(v255-v252, v260)
	v262 = v261 + v252
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v249+v262<<(uint(v260)%32))))
	if v266 == v204 {
		v327 = v242
		goto L88
	} else {
		goto L102
	}
L101:
	;
	goto L95
L102:
	;
	v275 = base.B2i32(v266-v204 < int32(0)) | base.B2i32(base.Ui32(v266) < base.Ui32(int32(3)))
	if v275 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v276 = v262 + int32(1)
	goto L105
L104:
	;
	v276 = v252
	goto L105
L105:
	;
	if v275 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v279 = v255
	goto L108
L107:
	;
	v279 = v262 - int32(1)
	goto L108
L108:
	;
	if v276 <= v279 {
		v252 = v276
		v255 = v279
		goto L100
	} else {
		goto L109
	}
L109:
	;
	goto L101
L110:
	;
	goto L94
L111:
	;
	v305 = int32(2)
	v306 = base.I32_div_s(v300-v297, v305)
	v307 = v306 + v297
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v295+v307<<(uint(v305)%32))))
	v312 = base.B2i32(v311 == v204)
	if v311 == v204 {
		v327 = v312
		goto L88
	} else {
		goto L113
	}
L112:
	;
	v327 = v312
	goto L88
L113:
	;
	v315 = base.B2i32(base.Ui32(v311) < base.Ui32(v204))
	if base.Ui32(v311) < base.Ui32(v204) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v316 = v307 + int32(1)
	goto L116
L115:
	;
	v316 = v297
	goto L116
L116:
	;
	if base.Ui32(v311) < base.Ui32(v204) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v319 = v300
	goto L119
L118:
	;
	v319 = v307 - int32(1)
	goto L119
L119:
	;
	if v316 <= v319 {
		v297 = v316
		v300 = v319
		goto L111
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v341 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v344 = F_MultiXactIdIsRunning(m, v342, v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L11
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L4
L124:
	;
	if v344 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v346 = int32(5)
	goto L127
L126:
	;
	v346 = v341
	goto L127
L127:
	;
	return v346
L128:
	;
	if v480 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L129:
	;
	v480 = int32(0)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v360 == v348 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v480 = int32(1)
	goto L128
L133:
	;
	goto L134
L134:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v364 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v480 = v470
	goto L128
L136:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v368 == int32(0) {
		v470 = int32(0)
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v440 = int32(0)
	v443 = v364 - int32(1)
	goto L158
L139:
	;
	v373 = v368
	goto L140
L140:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	if v379 == int32(4) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v470 = int32(0)
	goto L135
L142:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v373)+80))
	if v433 != 0 {
		v373 = v433
		goto L140
	} else {
		goto L157
	}
L143:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	if v382 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v385 = int32(1)
	if v348 == v382 {
		v470 = v385
		goto L135
	} else {
		goto L145
	}
L145:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)+52))
	v389 = v387 - int32(1)
	if v389 < int32(0) {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v373)+48))
	v395 = int32(0)
	v398 = v389
	goto L147
L147:
	;
	v403 = int32(2)
	v404 = base.I32_div_s(v398-v395, v403)
	v405 = v404 + v395
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v392+v405<<(uint(v403)%32))))
	if v409 == v348 {
		v470 = v385
		goto L135
	} else {
		goto L149
	}
L148:
	;
	goto L142
L149:
	;
	v418 = base.B2i32(v409-v348 < int32(0)) | base.B2i32(base.Ui32(v409) < base.Ui32(int32(3)))
	if v418 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v419 = v405 + int32(1)
	goto L152
L151:
	;
	v419 = v395
	goto L152
L152:
	;
	if v418 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v422 = v398
	goto L155
L154:
	;
	v422 = v405 - int32(1)
	goto L155
L155:
	;
	if v419 <= v422 {
		v395 = v419
		v398 = v422
		goto L147
	} else {
		goto L156
	}
L156:
	;
	goto L148
L157:
	;
	goto L141
L158:
	;
	v448 = int32(2)
	v449 = base.I32_div_s(v443-v440, v448)
	v450 = v449 + v440
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v438+v450<<(uint(v448)%32))))
	v455 = base.B2i32(v454 == v348)
	if v454 == v348 {
		v470 = v455
		goto L135
	} else {
		goto L160
	}
L159:
	;
	v470 = v455
	goto L135
L160:
	;
	v458 = base.B2i32(base.Ui32(v454) < base.Ui32(v348))
	if base.Ui32(v454) < base.Ui32(v348) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v459 = v450 + int32(1)
	goto L163
L162:
	;
	v459 = v440
	goto L163
L163:
	;
	if base.Ui32(v454) < base.Ui32(v348) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v462 = v443
	goto L166
L165:
	;
	v462 = v450 - int32(1)
	goto L166
L166:
	;
	if v459 <= v462 {
		v440 = v459
		v443 = v462
		goto L158
	} else {
		goto L167
	}
L167:
	;
	goto L159
L168:
	;
	goto L3
L169:
	;
	goto L170
L170:
	;
	goto L4
L171:
	;
	if v484 != 0 {
		v872 = v16
		goto L6
	} else {
		goto L172
	}
L172:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v487 = F_TransactionIdDidCommit(m, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L173
	}
L173:
	;
	if v487 == int32(0) {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	F_HeapTupleSetHintBits(m, v8, l2, int32(256), v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L175
	}
L175:
	;
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	v496 = v495
	goto L9
L176:
	;
	if v496&int32(1024) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if v496&int32(128)|base.B2i32(v496&int32(_a_F_HeapTupleSatisfiesUpdate_0) == int32(64)) != 0 {
		v872 = v498
		goto L6
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v496&int32(_a_F_HeapTupleSatisfiesUpdate_1) != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L2
L181:
	;
	if v496&int32(_a_F_HeapTupleSatisfiesUpdate_2) == int32(_a_F_HeapTupleSatisfiesUpdate_3) {
		v872 = v498
		goto L6
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.Ui32(v677) < base.Ui32(int32(3)) {
		goto L245
	} else {
		goto L246
	}
L184:
	;
	if v496&int32(128) != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v520 = F_MultiXactIdIsRunning(m, v518, int32(1))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v524 = F_HeapTupleGetUpdateXid(m, v8)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L193
	}
L188:
	;
	if v520 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	return int32(5)
L190:
	;
	goto L191
L191:
	;
	goto L3
L192:
	;
	if base.Ui32(v524) < base.Ui32(int32(3)) {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if v524 != 0 {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v528 = F_MultiXactIdIsRunning(m, v526, int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	if v528 == int32(0) {
		goto L192
	} else {
		goto L196
	}
L196:
	;
	return int32(5)
L197:
	;
	if v665 != 0 {
		goto L4
	} else {
		goto L237
	}
L198:
	;
	v665 = int32(0)
	goto L197
L199:
	;
	goto L200
L200:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v545 == v524 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v665 = int32(1)
	goto L197
L202:
	;
	goto L203
L203:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v549 <= int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v665 = v655
	goto L197
L205:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v553 == int32(0) {
		v655 = int32(0)
		goto L204
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v625 = int32(0)
	v628 = v549 - int32(1)
	goto L227
L208:
	;
	v558 = v553
	goto L209
L209:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	if v564 == int32(4) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v655 = int32(0)
	goto L204
L211:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v558)+80))
	if v618 != 0 {
		v558 = v618
		goto L209
	} else {
		goto L226
	}
L212:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if v567 == int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v570 = int32(1)
	if v524 == v567 {
		v655 = v570
		goto L204
	} else {
		goto L214
	}
L214:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v558)+52))
	v574 = v572 - int32(1)
	if v574 < int32(0) {
		goto L211
	} else {
		goto L215
	}
L215:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v558)+48))
	v580 = int32(0)
	v583 = v574
	goto L216
L216:
	;
	v588 = int32(2)
	v589 = base.I32_div_s(v583-v580, v588)
	v590 = v589 + v580
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v577+v590<<(uint(v588)%32))))
	if v594 == v524 {
		v655 = v570
		goto L204
	} else {
		goto L218
	}
L217:
	;
	goto L211
L218:
	;
	v603 = base.B2i32(v594-v524 < int32(0)) | base.B2i32(base.Ui32(v594) < base.Ui32(int32(3)))
	if v603 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v604 = v590 + int32(1)
	goto L221
L220:
	;
	v604 = v580
	goto L221
L221:
	;
	if v603 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v607 = v583
	goto L224
L223:
	;
	v607 = v590 - int32(1)
	goto L224
L224:
	;
	if v604 <= v607 {
		v580 = v604
		v583 = v607
		goto L216
	} else {
		goto L225
	}
L225:
	;
	goto L217
L226:
	;
	goto L210
L227:
	;
	v633 = int32(2)
	v634 = base.I32_div_s(v628-v625, v633)
	v635 = v634 + v625
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v623+v635<<(uint(v633)%32))))
	v640 = base.B2i32(v639 == v524)
	if v639 == v524 {
		v655 = v640
		goto L204
	} else {
		goto L229
	}
L228:
	;
	v655 = v640
	goto L204
L229:
	;
	v643 = base.B2i32(base.Ui32(v639) < base.Ui32(v524))
	if base.Ui32(v639) < base.Ui32(v524) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v644 = v635 + int32(1)
	goto L232
L231:
	;
	v644 = v625
	goto L232
L232:
	;
	if base.Ui32(v639) < base.Ui32(v524) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v647 = v628
	goto L235
L234:
	;
	v647 = v635 - int32(1)
	goto L235
L235:
	;
	if v644 <= v647 {
		v625 = v644
		v628 = v647
		goto L227
	} else {
		goto L236
	}
L236:
	;
	goto L228
L237:
	;
	v666 = int32(5)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v669 = F_MultiXactIdIsRunning(m, v667, int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L11
	} else {
		goto L238
	}
L238:
	;
	if v669 != 0 {
		v872 = v666
		goto L6
	} else {
		goto L239
	}
L239:
	;
	v671 = F_TransactionIdDidCommit(m, v524)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L11
	} else {
		goto L240
	}
L240:
	;
	if v671 != 0 {
		goto L2
	} else {
		goto L241
	}
L241:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v675 = F_MultiXactIdIsRunning(m, v673, int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L11
	} else {
		goto L242
	}
L242:
	;
	if v675 != 0 {
		v872 = v666
		goto L6
	} else {
		goto L243
	}
L243:
	;
	goto L3
L244:
	;
	if v809 != 0 {
		goto L284
	} else {
		goto L285
	}
L245:
	;
	v809 = int32(0)
	goto L244
L246:
	;
	goto L247
L247:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[0]))
	if v689 == v677 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v809 = int32(1)
	goto L244
L249:
	;
	goto L250
L250:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[1]))
	if v693 <= int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v809 = v799
	goto L244
L252:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[2]))
	if v697 == int32(0) {
		v799 = int32(0)
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[3]))
	v769 = int32(0)
	v772 = v693 - int32(1)
	goto L274
L255:
	;
	v702 = v697
	goto L256
L256:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v702)+20))
	if v708 == int32(4) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v799 = int32(0)
	goto L251
L258:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v702)+80))
	if v762 != 0 {
		v702 = v762
		goto L256
	} else {
		goto L273
	}
L259:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	if v711 == int32(0) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v714 = int32(1)
	if v677 == v711 {
		v799 = v714
		goto L251
	} else {
		goto L261
	}
L261:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v702)+52))
	v718 = v716 - int32(1)
	if v718 < int32(0) {
		goto L258
	} else {
		goto L262
	}
L262:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v724 = int32(0)
	v727 = v718
	goto L263
L263:
	;
	v732 = int32(2)
	v733 = base.I32_div_s(v727-v724, v732)
	v734 = v733 + v724
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v721+v734<<(uint(v732)%32))))
	if v738 == v677 {
		v799 = v714
		goto L251
	} else {
		goto L265
	}
L264:
	;
	goto L258
L265:
	;
	v747 = base.B2i32(v738-v677 < int32(0)) | base.B2i32(base.Ui32(v738) < base.Ui32(int32(3)))
	if v747 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v748 = v734 + int32(1)
	goto L268
L267:
	;
	v748 = v724
	goto L268
L268:
	;
	if v747 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v751 = v727
	goto L271
L270:
	;
	v751 = v734 - int32(1)
	goto L271
L271:
	;
	if v748 <= v751 {
		v724 = v748
		v727 = v751
		goto L263
	} else {
		goto L272
	}
L272:
	;
	goto L264
L273:
	;
	goto L257
L274:
	;
	v777 = int32(2)
	v778 = base.I32_div_s(v772-v769, v777)
	v779 = v778 + v769
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v767+v779<<(uint(v777)%32))))
	v784 = base.B2i32(v783 == v677)
	if v783 == v677 {
		v799 = v784
		goto L251
	} else {
		goto L276
	}
L275:
	;
	v799 = v784
	goto L251
L276:
	;
	v787 = base.B2i32(base.Ui32(v783) < base.Ui32(v677))
	if base.Ui32(v783) < base.Ui32(v677) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v788 = v779 + int32(1)
	goto L279
L278:
	;
	v788 = v769
	goto L279
L279:
	;
	if base.Ui32(v783) < base.Ui32(v677) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v791 = v772
	goto L282
L281:
	;
	v791 = v779 - int32(1)
	goto L282
L282:
	;
	if v788 <= v791 {
		v769 = v788
		v772 = v791
		goto L274
	} else {
		goto L283
	}
L283:
	;
	goto L275
L284:
	;
	v811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	if v811&int32(128)|base.B2i32(v811&int32(_a_F_HeapTupleSatisfiesUpdate_0) == int32(64)) != 0 {
		v872 = int32(5)
		goto L6
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v820 = F_TransactionIdIsInProgress(m, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L11
	} else {
		goto L288
	}
L287:
	;
	goto L4
L288:
	;
	if v820 != 0 {
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
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v825 = F_TransactionIdDidCommit(m, v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L11
	} else {
		goto L292
	}
L292:
	;
	v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)))
	if v825 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	goto L1
L294:
	;
	goto L295
L295:
	;
	v832 = int32(0)
	if base.B2i32(v827&int32(128) == v832)&base.B2i32(v827&int32(_a_F_HeapTupleSatisfiesUpdate_0) != int32(64)) == v832 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	goto L1
L297:
	;
	goto L298
L298:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F_HeapTupleSetHintBits(m, v8, l2, int32(1024), v842)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L11
	} else {
		goto L299
	}
L299:
	;
	v845 = int32(4)
	v848 = l0 + v845
	v850 = v8 + int32(12)
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v848)+2)))
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v848))))
	v853 = int32(16)
	v856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850)+2)))
	v857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850))))
	if v851|v852<<(uint(v853)%32) == v856|v857<<(uint(v853)%32) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	if v867 != 0 {
		goto L306
	} else {
		goto L307
	}
L301:
	;
	goto L300
L302:
	;
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v848)+4)))
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850)+4)))
	if v863 == v864 {
		v867 = int32(1)
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v867 = int32(0)
	goto L301
L305:
	;
	goto L304
L306:
	;
	v868 = v845
	goto L308
L307:
	;
	v868 = int32(3)
	goto L308
L308:
	;
	v872 = v868
	goto L6
L309:
	;
	return int32(1)
L310:
	;
	if base.Ui32(v897) < base.Ui32(l1) {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesUpdate[4]))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v892+v887<<(uint(int32(3))%32))+4))
	v897 = v896
	goto L313
L312:
	;
	v897 = v887
	goto L313
L313:
	;
	goto L310
L314:
	;
	v899 = int32(1)
	goto L316
L315:
	;
	v899 = int32(2)
	goto L316
L316:
	;
	return v899
L317:
	;
	return int32(0)
L318:
	;
	if v935 != 0 {
		goto L324
	} else {
		goto L325
	}
L319:
	;
	goto L318
L320:
	;
	v931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916)+4)))
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)))
	if v931 == v932 {
		v935 = int32(1)
		goto L319
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v935 = int32(0)
	goto L319
L323:
	;
	goto L322
L324:
	;
	v936 = v913
	goto L326
L325:
	;
	v936 = int32(3)
	goto L326
L326:
	;
	return v936
L327:
	;
	return int32(0)
}
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v121 int32
	_ = v121
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v412 int64
	_ = v412
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v26 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v67
	F_swap_relation_files(m, l0, l1, base.B2i32(l0 == int32(1259)), l3, l5, l7, l8, v20+int32(112))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v30&int32(1) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v35 = int32(_a_F_finish_heap_swap_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v37 + v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v41 + v38
	v45 = int32(0)
	v47 = int32(_a_F_finish_heap_swap_1)
	v48 = base.AtomicRmwOr32(m, v45, v47, v45)
	*(*int64)(unsafe.Add(mBase, uint32(v26+int32(8))+232)) = int64(6)
	v56 = base.AtomicRmwOr32(m, v45, v47, v45)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v57 + v38
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v63 - v38
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
	v79 = int32(1)
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
	if l6 != 0 {
		goto L48
	} else {
		goto L49
	}
L10:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[3]))
	v153 = F_PrepareInvalidationState(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L35
	}
L11:
	;
	goto L10
L12:
	;
	v150 = int32(0)
	goto L11
L13:
	;
	if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v150 = v79
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
	if l0 <= int32(_a_F_finish_heap_swap_2) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v150 = v79
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
	v91 = l0 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v91))|base.B2i32(int32(1)<<(uint(v91)%32)&int32(226492515) == int32(0)) != 0 {
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
	v150 = v79
	goto L11
L22:
	;
	v150 = v79
	goto L11
L23:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v150 = v79
		goto L11
	} else {
		goto L32
	}
L24:
	;
	v104 = l0 - int32(_a_F_finish_heap_swap_3)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v104))|base.B2i32(int32(1)<<(uint(v104)%32)&int32(963) == int32(0)) != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	switch l0 - int32(_a_F_finish_heap_swap_4) {
	case 0, 1, 2, 3, 4, 59, 60:
		v150 = v79
		goto L11
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L12
	default:
		goto L28
	}
L27:
	;
	v150 = v79
	goto L11
L28:
	;
	if base.Ui32(l0-int32(_a_F_finish_heap_swap_5)) < base.Ui32(int32(3)) {
		v150 = v79
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v121 = l0 - int32(_a_F_finish_heap_swap_6)
	if base.Ui32(int32(15)) < base.Ui32(v121) {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	if int32(1)<<(uint(v121)%32)&int32(_a_F_finish_heap_swap_7) != 0 {
		v150 = v79
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
	v150 = v79
	goto L11
L34:
	;
	goto L12
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[4]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[5]))
	if v159 <= v157 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v156 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v181 = v156
	goto L38
L38:
	;
	v185 = v181 + v157<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = l0
	if v150 != 0 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[5])) = v175
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[4])) = v176
	v181 = v176
	goto L38
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[6]))
	v167 = F_MemoryContextAlloc(m, v165, int32(512))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v173 = F_repalloc(m, v156, v159<<(uint(int32(5))%32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v175 = int32(32)
	v176 = v167
	goto L39
L44:
	;
	v175 = v159 << (uint(int32(1)) % 32)
	v176 = v173
	goto L39
L45:
	;
	v188 = int32(0)
	goto L47
L46:
	;
	v188 = v152
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v188
	v190 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v190)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v192 + int32(1)
	goto L9
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(0)
	if l4 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v270 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	v207 = int32(6)
	goto L53
L52:
	;
	v207 = int32(2)
	goto L53
L53:
	;
	switch l9 - int32(112) {
	case 0:
		goto L55
	default:
		v214 = v207
		goto L54
	case 5:
		goto L56
	}
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[0]))
	if v219 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v214 = v207 | int32(16)
	goto L54
L56:
	;
	v214 = v207 | int32(8)
	goto L54
L57:
	;
	v263 = F_reindex_relation(m, int32(0), l0, v214, v20+int32(48))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L61
	}
L58:
	;
	goto L57
L59:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v223&int32(1) == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v228 = int32(_a_F_finish_heap_swap_0)
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v231 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v230 + v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v234 + v231
	v238 = int32(0)
	v240 = int32(_a_F_finish_heap_swap_1)
	v241 = base.AtomicRmwOr32(m, v238, v240, v238)
	*(*int64)(unsafe.Add(mBase, uint32(v219+int32(8))+232)) = int64(7)
	v249 = base.AtomicRmwOr32(m, v238, v240, v238)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v250 + v231
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v256 - v231
	goto L58
L61:
	;
	goto L50
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
	v274 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_finish_heap_swap[1])))
	if v274&int32(1) == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v279 = int32(_a_F_finish_heap_swap_0)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	v282 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v281 + v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v285 + v282
	v289 = int32(0)
	v291 = int32(_a_F_finish_heap_swap_1)
	v292 = base.AtomicRmwOr32(m, v289, v291, v289)
	*(*int64)(unsafe.Add(mBase, uint32(v270+int32(8))+232)) = int64(8)
	v300 = base.AtomicRmwOr32(m, v289, v291, v289)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v301 + v282
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[2])) = v307 - v282
	goto L63
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L5
	} else {
		goto L129
	}
L67:
	;
	v315 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v339 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+140)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = int32(1259)
	if l6 == v339 {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v320 = F_SearchSysCacheCopy(m, int32(57), int64(1259), int64(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	if v320 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+22)))
	v326 = v324 + v325
	*(*int32)(unsafe.Add(mBase, uint32(v326)+140)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v326)+136)) = l7
	F_CatalogTupleUpdate(m, v315, v320+int32(4), v320)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_relation_close(m, v315, int32(3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_performDeletion(m, v20+int32(132), int32(0), int32(1))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
	if v355 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v357 = v355
	v364 = v339
	goto L83
L81:
	;
	goto L82
L82:
	;
	if l3 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L83:
	;
	v373 = int32(0)
	v374 = m.G0
	v376 = v374 - int32(16)
	m.G0 = v376
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[7]))
	if v379 <= v373 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L82
L85:
	;
	v454 = v364 + int32(1)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(112)+v454<<(uint(int32(2))%32))))
	if v458 != 0 {
		v357 = v458
		v364 = v454
		goto L83
	} else {
		goto L97
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L94
	}
L87:
	;
	v386 = v373
	goto L88
L88:
	;
	v400 = v386 << (uint(int32(3)) % 32)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+uint32(_c_F_finish_heap_swap[8])))
	if v401 != v357 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v379<<(uint(int32(3))%32))+uint32(_c_F_finish_heap_swap[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_c_F_finish_heap_swap[8]))) = v412
	*(*int32)(unsafe.Add(mBase, _c_F_finish_heap_swap[7])) = v379 - int32(1)
	m.G0 = v376 + int32(16)
	goto L85
L90:
	;
	v404 = v386 + int32(1)
	if v379 != v404 {
		v386 = v404
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L86
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v357
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_8), v376)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_9), int32(455), int32(_a_F_finish_heap_swap_10))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
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
	goto L84
L98:
	;
	v479 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if l2 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L101:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+112))
	if v482 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v483 = F_toast_get_valid_index(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_relation_close(m, v479, int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L5
	} else {
		goto L122
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l0
	v487 = v20 + int32(48)
	v492 = F_pg_snprintf(m, v487, int32(64), int32(_a_F_finish_heap_swap_11), v20+int32(32))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+112))
	F_RenameRelationInternal(m, v495, v487, int32(1), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	v505 = F_pg_snprintf(m, v487, int32(64), int32(_a_F_finish_heap_swap_12), v20+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v507 = int32(1)
	F_RenameRelationInternal(m, v483, v487, v507, v507)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+112))
	v515 = m.G0
	v517 = v515 - int32(16)
	m.G0 = v517
	v521 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v526 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v514), int64(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	if v526 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L5
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v543+v544)+132)) = int32(0)
	F_CatalogTupleUpdate(m, v521, v526+int32(4), v526)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L119
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v514
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_13), v517)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_14), int32(_a_F_finish_heap_swap_15), int32(_a_F_finish_heap_swap_16))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_pfree(m, v526)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	F_relation_close(m, v521, int32(3))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	m.G0 = v517 + int32(16)
	goto L104
L122:
	;
	goto L100
L123:
	;
	v575 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L5
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	m.G0 = v20 + int32(144)
	return
L126:
	;
	F_RelationClearMissing(m, v575)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	F_relation_close(m, v575, int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1259)
	F_errmsg_internal(m, int32(_a_F_finish_heap_swap_13), v20)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_finish_heap_swap_17), int32(2113), int32(_a_F_finish_heap_swap_18))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_compare_slots_1(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	if v12 <= v4 {
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+144))
	v19 = int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17+base.I32_wrap_i64(l1)<<(uint(v19)%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17+base.I32_wrap_i64(l0)<<(uint(v19)%32))))
	v35 = v4
	goto L6
L4:
	;
	return int32(-1)
L5:
	;
	v107 = int32(0)
	if v98 < v107 {
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	v42 = v39 + v35*int32(36)
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+10)))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+6)))
	if v44 < v43 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	m.T0[v47].(func(*base.Module, int32, int32))(m, v27, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v53 = v43 - int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v54))))
	v58 = v53 << (uint(int32(3)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)))
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	if v62 < v43 {
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	m.T0[v65].(func(*base.Module, int32, int32))(m, v22, v43)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v53))))
	if v56&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v101 = v35 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	if v101 < v102 {
		v35 = v101
		goto L6
	} else {
		goto L33
	}
L18:
	;
	if v70&int32(1) != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v70&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+9)))
	if v75 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	return int32(1)
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+9)))
	if v82 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v85+v58)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v89 = m.T0[v88].(func(*base.Module, int64, int64, int32) int32)(m, v61, v87, v42)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	return int32(1)
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+8)))
	if v91 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v89 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v98 = v89
	goto L30
L30:
	;
	if v98 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(0) - v89
	goto L30
L32:
	;
	goto L17
L33:
	;
	goto L7
L34:
	;
	v111 = int32(1)
	goto L36
L35:
	;
	v111 = v107 - v98
	goto L36
L36:
	;
	return v111
}
func F_heap_copy_tuple_as_datum(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
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
			return int64(0)
		} else {
			return v10
		}
	} else {
		v15 = F_palloc(m, v5)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
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
			return base.I64_extend_i32_u(v15)
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
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
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int64
	_ = v907
	var v908 int32
	_ = v908
	var v910 int64
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v934 int64
	_ = v934
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	v22 = m.G0
	v24 = v22 + int32(-64)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = l2
	v27 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+26)) = uint8(v31)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v24 - int32(-64)
	return v999
L4:
	;
	if v51 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L5:
	;
	F_UnlockReleaseBuffer(m, v51)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L166
	}
L6:
	;
	if v41&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v41 = int32(1)
	goto L9
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+76)))
	v41 = v40
	goto L9
L9:
	;
	goto L6
L10:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v50 = v46 | v47<<(uint(int32(16))%32)
	v51 = F_ReadBuffer(m, l0, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L162
	}
L13:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	if v71&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v51 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[1]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v51^int32(-1))<<(uint(int32(2))%32))))
	v70 = v62
	goto L13
L16:
	;
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[2]))
	v70 = v64 + v51<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L18:
	;
	F_visibilitymap_pin(m, l0, v50, v22+int32(-28))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v79 = l3 & int32(2)
	v81 = l3 & int32(1)
	F_LockBufferInternal(m, v51, int32(3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v86
	v89 = v70 + int32(20)
	v92 = v89 + v85<<(uint(int32(2))%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v70 + v93&int32(_a_F_heap_delete_0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(base.Ui32(v98) >> (uint(int32(17)) % 32))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v102
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+48)) = uint16(v104)
	v107 = v22 + int32(-20)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v111 = v108
	v121 = int32(0)
	goto L26
L23:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)) = uint16(v463)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v465
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462)+20)))
	if v468&int32(_a_F_heap_delete_1) != int32(_a_F_heap_delete_2) {
		goto L136
	} else {
		goto L137
	}
L24:
	;
	v441 = int32(0)
	if base.B2i32(l4 == v441)|v435 == v441 {
		goto L129
	} else {
		goto L130
	}
L25:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+20)))
	if v398&int32(2048)|v398&int32(128)|base.B2i32(v398&int32(_a_F_heap_delete_3) == int32(64)) != 0 {
		v435 = v396
		v438 = v394
		goto L24
	} else {
		goto L117
	}
L26:
	;
	if v111 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v366&int32(3072) != 0 {
		v394 = v364
		goto L25
	} else {
		goto L110
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v147 = F_HeapTupleSatisfiesUpdate(m, v22+int32(-24), v146, v51)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	if v130&int32(4) == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_UnlockBuffer(m, v51)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_visibilitymap_pin(m, l0, v50, v22+int32(-28))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_LockBufferInternal(m, v51, int32(3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	if v147 == int32(1) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v147 != int32(5)) != 0 {
		v435 = v147
		v438 = v121
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+20)))
	if v158&int32(_a_F_heap_delete_2) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)) = uint8(v161)
	v166 = F_DoesMultiXactIdConflict(m, v157, v158, int32(3), v22+int32(-48))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v157) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	if v166 == int32(0) {
		v394 = v121
		goto L25
	} else {
		goto L41
	}
L41:
	;
	F_UnlockBuffer(m, v51)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
	if (v172|v121)&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v185 = int32(0)
	v189 = F_Do_MultiXactIdWait(m, v157, int32(5), v158, v185, l0, v107, int32(2), v185, v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	v183 = v172 ^ int32(1) | v121
	goto L43
L45:
	;
	goto L46
L46:
	;
	F_LockTuple(m, l0, v107, int32(8))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v183 = int32(1)
	goto L43
L48:
	;
	F_LockBufferInternal(m, v51, int32(3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v194 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	if v197&int32(4) != 0 {
		v111 = v194
		v121 = v183
		goto L26
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+20)))
	if (v201^v158)&int32(_a_F_heap_delete_4) != 0 {
		v111 = v194
		v121 = v183
		goto L26
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v205 != v157 {
		v111 = v194
		v121 = v183
		goto L26
	} else {
		goto L55
	}
L55:
	;
	v394 = v183
	goto L25
L56:
	;
	if v338 != 0 {
		v394 = v121
		goto L25
	} else {
		goto L96
	}
L57:
	;
	v338 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[3]))
	if v218 == v157 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v338 = int32(1)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[4]))
	if v222 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v338 = v328
	goto L56
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[0]))
	if v226 == int32(0) {
		v328 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[5]))
	v298 = int32(0)
	v301 = v222 - int32(1)
	goto L86
L67:
	;
	v231 = v226
	goto L68
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v237 == int32(4) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v328 = int32(0)
	goto L63
L70:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	if v291 != 0 {
		v231 = v291
		goto L68
	} else {
		goto L85
	}
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v240 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v243 = int32(1)
	if v157 == v240 {
		v328 = v243
		goto L63
	} else {
		goto L73
	}
L73:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v247 = v245 - int32(1)
	if v247 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v253 = int32(0)
	v256 = v247
	goto L75
L75:
	;
	v261 = int32(2)
	v262 = base.I32_div_s(v256-v253, v261)
	v263 = v262 + v253
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v250+v263<<(uint(v261)%32))))
	if v267 == v157 {
		v328 = v243
		goto L63
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v276 = base.B2i32(v267-v157 < int32(0)) | base.B2i32(base.Ui32(v267) < base.Ui32(int32(3)))
	if v276 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v277 = v263 + int32(1)
	goto L80
L79:
	;
	v277 = v253
	goto L80
L80:
	;
	if v276 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v280 = v256
	goto L83
L82:
	;
	v280 = v263 - int32(1)
	goto L83
L83:
	;
	if v277 <= v280 {
		v253 = v277
		v256 = v280
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
	v306 = int32(2)
	v307 = base.I32_div_s(v301-v298, v306)
	v308 = v307 + v298
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v296+v308<<(uint(v306)%32))))
	v313 = base.B2i32(v312 == v157)
	if v312 == v157 {
		v328 = v313
		goto L63
	} else {
		goto L88
	}
L87:
	;
	v328 = v313
	goto L63
L88:
	;
	v316 = base.B2i32(base.Ui32(v312) < base.Ui32(v157))
	if base.Ui32(v312) < base.Ui32(v157) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v317 = v308 + int32(1)
	goto L91
L90:
	;
	v317 = v298
	goto L91
L91:
	;
	if base.Ui32(v312) < base.Ui32(v157) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v320 = v301
	goto L94
L93:
	;
	v320 = v308 - int32(1)
	goto L94
L94:
	;
	if v317 <= v320 {
		v298 = v317
		v301 = v320
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L87
L96:
	;
	F_UnlockBuffer(m, v51)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v121&int32(1) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_LockTuple(m, l0, v107, int32(8))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_XactLockTableWait(m, v157, l0, v107, int32(2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	F_LockBufferInternal(m, v51, int32(3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v354 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	if v359&int32(4) != 0 {
		v111 = int32(0)
		v121 = int32(1)
		goto L26
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v364 = int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+20)))
	if (v158^v366)&int32(_a_F_heap_delete_4) != 0 {
		v111 = v354
		v121 = v364
		goto L26
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v370 != v157 {
		v111 = v354
		v121 = v364
		goto L26
	} else {
		goto L109
	}
L109:
	;
	goto L27
L110:
	;
	if v366&int32(128)|base.B2i32(v366&int32(_a_F_heap_delete_3) == int32(64)) != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_HeapTupleSetHintBits(m, v365, v51, int32(2048), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L116
	}
L112:
	;
	v381 = F_TransactionIdDidCommit(m, v157)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v381 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	F_HeapTupleSetHintBits(m, v365, v51, int32(1024), v157)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v394 = v364
	goto L25
L116:
	;
	v394 = v364
	goto L25
L117:
	;
	v409 = F_HeapTupleHeaderIsOnlyLocked(m, v397)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	if v409 != 0 {
		v435 = v396
		v438 = v394
		goto L24
	} else {
		goto L119
	}
L119:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v415 = v413 + int32(12)
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+2)))
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107))))
	v418 = int32(16)
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+2)))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415))))
	if v416|v417<<(uint(v418)%32) == v421|v422<<(uint(v418)%32) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v432 != 0 {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	goto L120
L122:
	;
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+4)))
	v429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+4)))
	if v428 == v429 {
		v432 = int32(1)
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v432 = int32(0)
	goto L121
L125:
	;
	goto L124
L126:
	;
	v433 = int32(4)
	goto L128
L127:
	;
	v433 = int32(3)
	goto L128
L128:
	;
	v456 = v433
	v459 = v394
	goto L23
L129:
	;
	v449 = F_HeapTupleSatisfiesVisibility(m, v22+int32(-24), l4, v51)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v435 == int32(0) {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	if v449 == int32(0) {
		v456 = int32(3)
		v459 = v438
		goto L23
	} else {
		goto L133
	}
L133:
	;
	goto L4
L134:
	;
	v456 = v435
	v459 = v438
	goto L23
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v519
	if v456 == int32(2) {
		goto L148
	} else {
		goto L149
	}
L136:
	;
	v519 = v467
	goto L135
L137:
	;
	goto L138
L138:
	;
	v473 = int32(0)
	v477 = F_GetMultiXactIdMembers(m, v467, v22+int32(-48), v473)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v477 <= int32(0) {
		v519 = v473
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v483 = v473
	goto L143
L141:
	;
	F_pfree(m, v481)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L147
	}
L142:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v515 = v513
	goto L141
L143:
	;
	v505 = v481 + v483<<(uint(int32(3))%32)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v506) {
		goto L142
	} else {
		goto L145
	}
L144:
	;
	v515 = int32(0)
	goto L141
L145:
	;
	v510 = v483 + int32(1)
	if v510 != v477 {
		v483 = v510
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v519 = v515
	goto L135
L148:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+20)))
	if v545&int32(32) != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v556 = int32(-1)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v556
	F_UnlockReleaseBuffer(m, v51)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L155
	}
L151:
	;
	v556 = v554
	goto L150
L152:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[6]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v544<<(uint(int32(3))%32))+4))
	v554 = v553
	goto L154
L153:
	;
	v554 = v544
	goto L154
L154:
	;
	goto L151
L155:
	;
	if v459&int32(1) != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_UnlockTuple(m, l0, v107, int32(8))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v565 == int32(0) {
		v999 = v456
		goto L3
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	F_ReleaseBuffer(m, v565)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v999 = v456
	goto L3
L162:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(_a_F_heap_delete_5), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_heap_delete_6), int32(2795), int32(_a_F_heap_delete_7))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_heap_delete_8), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_heap_delete_6), int32(2842), int32(_a_F_heap_delete_7))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
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
	F_CheckForSerializableConflictIn(m, l0, l1, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L175
	}
L172:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[7]))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v608+(v51^int32(-1))*int32(56))+16))
	v623 = v614
	goto L171
L173:
	;
	goto L174
L174:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[8]))
	v617 = int32(56)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v616+v51*v617-v617)+16))
	v623 = v622
	goto L171
L175:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	F_HeapTupleHeaderAdjustCmax(m, v626, v22+int32(-4), v22+int32(-37))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v633 = int32(0)
	if v79 == v633 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v641 = F_ExtractReplicaIdentity(m, l0, v22+int32(-24), int32(1), v22+int32(-38))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	v643 = v633
	goto L179
L179:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	v643 = v641
	goto L179
L181:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+4))
	v648 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646)+20)))
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646)+18)))
	F_compute_new_xmax_infomask(m, v647, v648, v649, v27, int32(3), int32(1), v22+int32(-32), v22+int32(-34), v22+int32(-36))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+10)))
	v662 = v660 & int32(4)
	if v662 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	F_LockBufferInternal(m, v663, int32(3))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v667 = int32(_a_F_heap_delete_9)
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9])) = v669 + int32(1)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v673 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L185
L187:
	;
	if v662 != 0 {
		goto L195
	} else {
		goto L196
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v27
	goto L187
L189:
	;
	v676 = int32(3)
	if base.B2i32(base.Ui32(v27) < base.Ui32(v676))|base.B2i32(base.Ui32(v673) < base.Ui32(v676)) == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if v27-v673 < int32(0) {
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if base.Ui32(v673) <= base.Ui32(v27) {
		goto L187
	} else {
		goto L194
	}
L193:
	;
	goto L187
L194:
	;
	goto L188
L195:
	;
	if v51 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	v716 = int32(0)
	goto L197
L197:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v718 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v717)+20)))
	v720 = v718 & int32(_a_F_heap_delete_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v717)+20)) = uint16(v720)
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v717)+18)))
	v724 = v722 & int32(_a_F_heap_delete_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v717)+18)) = uint16(v724)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)))
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+30)))
	v729 = v727 | v728
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)) = uint16(v729)
	v731 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)))
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)))
	v733 = v731 | v732
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)) = uint16(v733)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v735)+18)))
	v738 = v736 & int32(_a_F_heap_delete_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v735)+18)) = uint16(v738)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v740)+4)) = v741
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+27)))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v744)+8)) = v745
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v744)+20)))
	v754 = v747&int32(_a_F_heap_delete_13) | v743<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v744)+20)) = uint16(v754)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v756)+16)) = uint16(v757)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+12)) = v759
	if v81 != 0 {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v710 = F_visibilitymap_clear(m, v707, v708, int32(3))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L202
	}
L199:
	;
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[7]))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v692+(v51^int32(-1))*int32(56))+16))
	v707 = v698
	goto L198
L200:
	;
	goto L201
L201:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[8]))
	v701 = int32(56)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v700+v51*v701-v701)+16))
	v707 = v706
	goto L198
L202:
	;
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+10)))
	v714 = v712 & int32(_a_F_heap_delete_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+10)) = uint16(v714)
	v716 = v710
	goto L197
L203:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v762 = int32(_a_F_heap_delete_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v761)+16)) = uint16(v762)
	*(*int32)(unsafe.Add(mBase, uint32(v761)+12)) = int32(-1)
	goto L205
L204:
	;
	goto L205
L205:
	;
	F_MarkBufferDirty(m, v51)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+118)))
	if v770 != int32(112) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v943 = int32(_a_F_heap_delete_9)
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_delete[9])) = v945 - int32(1)
	if v662 != 0 {
		goto L257
	} else {
		goto L258
	}
L208:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[10]))
	if v774 <= int32(0) {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v813 = int32(base.Ui32(v662) >> (uint(int32(2)) % 32))
	if v81 != 0 {
		goto L227
	} else {
		goto L228
	}
L210:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L219
L211:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v777 != 0 {
		goto L207
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if v774 != int32(1) {
		goto L210
	} else {
		goto L217
	}
L214:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v778 != 0 {
		goto L207
	} else {
		goto L215
	}
L215:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[11])))
	if v780 == int32(1) {
		goto L210
	} else {
		goto L216
	}
L216:
	;
	goto L209
L217:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[11])))
	if v786&int32(1) == int32(0) {
		goto L209
	} else {
		goto L218
	}
L218:
	;
	goto L210
L219:
	;
	if base.B2i32(base.Ui32(v791) < base.Ui32(int32(_a_F_heap_delete_16))) == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v796 == int32(0) {
		goto L209
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	F_log_heap_new_cid(m, l0, v22+int32(-24))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L226
	}
L223:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799)+119)))
	switch v800 - int32(109) {
	case 0, 5:
		goto L224
	default:
		goto L209
	}
L224:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+112)))
	if v803 != int32(1) {
		goto L209
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	goto L209
L227:
	;
	v816 = v813 | int32(16)
	goto L229
L228:
	;
	v816 = v813
	goto L229
L229:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)) = uint8(v816)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v818)+18)))
	v820 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v818)+20)))
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+48)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)) = uint16(v821)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v741
	v828 = int32(1)
	v832 = int32(4)
	v847 = int32(base.Ui32(v819)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v820)>>(uint(v828)%32))&int32(8) | (int32(base.Ui32(v820)>>(uint(v832)%32))&v832 | (int32(base.Ui32(v820)>>(uint(int32(12))%32))&v828 | int32(base.Ui32(v820)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)) = uint8(v847)
	if v643 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851)+130)))
	if v852 == int32(102) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v858 = v816
	goto L232
L232:
	;
	if v79 != 0 {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	v855 = int32(2)
	goto L235
L234:
	;
	v855 = int32(4)
	goto L235
L235:
	;
	v856 = v816 | v855
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)) = uint8(v856)
	v858 = v856
	goto L232
L236:
	;
	v860 = v858 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)) = uint8(v860)
	goto L238
L237:
	;
	goto L238
L238:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_XLogRegisterData(m, v22+int32(-48), int32(8))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_XLogRegisterBuffer(m, int32(0), v51, int32(8))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	if v643 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v643)+16))
	v874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)) = uint16(v874)
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)) = uint16(v876)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+14)) = uint8(v878)
	F_XLogRegisterData(m, v22+int32(-54), int32(5))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v895 = int32(_a_F_heap_delete_17)
	v897 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[12])))
	v898 = v897 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_delete[12])) = uint8(v898)
	goto L247
L245:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v643)+16))
	v886 = int32(23)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	F_XLogRegisterData(m, v885+v886, v888-v886)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	if v716 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	F_XLogRegisterBuffer(m, int32(1), v901, int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v934 = F_XLogInsert(m, int32(10), int32(16))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L256
	}
L251:
	;
	v907 = F_XLogInsert(m, int32(10), int32(16))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v910 = base.I64_rotl(v907, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v910
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v912 < int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[1]))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v916+(v912^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v922))) = v910
	goto L207
L254:
	;
	goto L255
L255:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _c_F_heap_delete[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v925+v912<<(uint(int32(13))%32))+uint32(_c_F_heap_delete[13]))) = v910
	goto L207
L256:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = base.I64_rotl(v934, int64(32))
	goto L207
L257:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	F_UnlockBuffer(m, v949)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	F_UnlockBuffer(m, v51)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v954 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	F_ReleaseBuffer(m, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957)+119)))
	switch v958 - int32(109) {
	case 0, 5:
		goto L267
	default:
		goto L266
	}
L265:
	;
	goto L264
L266:
	;
	v972 = int32(0)
	F_CacheInvalidateHeapTuple(m, l0, v22+int32(-24), v972)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L270
	}
L267:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961)+20)))
	if v962&int32(4) == int32(0) {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	F_heap_toast_delete(m, l0, v22+int32(-24), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	goto L266
L270:
	;
	F_ReleaseBuffer(m, v51)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	if v438&int32(1) != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	F_UnlockTuple(m, l0, v107, int32(8))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	F_pgstat_count_heap_delete(m, l0)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	if v643 == int32(0) {
		v999 = v972
		goto L3
	} else {
		goto L277
	}
L277:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+26)))
	if v989&int32(1) == int32(0) {
		v999 = v972
		goto L3
	} else {
		goto L278
	}
L278:
	;
	F_pfree(m, v643)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v999 = v972
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v4 != 0 {
		F_ReleaseBuffer(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[0]))
			if v10 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					if l1 != v13 {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v15
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						F_read_stream_reset(m, v17)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							v23 = F_read_stream_next_buffer(m, v21, int32(0))
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
								if v23 != 0 {
									if v23 < int32(0) {
										v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
										v44 = v35
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
										v38 = int32(56)
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
										v44 = v43
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
								} else {
								}
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
									v38 = int32(56)
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
							} else {
							}
							return
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				if l1 != v13 {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v15
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					F_read_stream_reset(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
									v38 = int32(56)
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
							} else {
							}
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
								v38 = int32(56)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
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
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				if l1 != v13 {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v15
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					F_read_stream_reset(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v23 = F_read_stream_next_buffer(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
							if v23 != 0 {
								if v23 < int32(0) {
									v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
									v44 = v35
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
									v38 = int32(56)
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
									v44 = v43
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
							} else {
							}
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
								v38 = int32(56)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
						} else {
						}
						return
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			if l1 != v13 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v15
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				F_read_stream_reset(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v23 = F_read_stream_next_buffer(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
						if v23 != 0 {
							if v23 < int32(0) {
								v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
								v44 = v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
								v38 = int32(56)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
								v44 = v43
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
						} else {
						}
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				v23 = F_read_stream_next_buffer(m, v21, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v23
					if v23 != 0 {
						if v23 < int32(0) {
							v29 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[1]))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v23^int32(-1))*int32(56))+16))
							v44 = v35
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch_next_buffer[2]))
							v38 = int32(56)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v23*v38-v38)+16))
							v44 = v43
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
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
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
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
	v164 = m.ExcPending
	if v164 != 0 {
		goto L14
	} else {
		goto L39
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
	v26 = v5
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
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
	v66 = (int32(base.Ui32(v15+int32(7))>>(uint(int32(3))%32)) + int32(22)) & int32(536870904)
	v67 = int32(128)
	v69 = int32(1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	v46 = v26 + int32(1)
	if v46 != v15 {
		v26 = v46
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
	v78 = v76 + l3
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v74
	v81 = v66 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)) = uint8(v81)
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)))
	v86 = v83&int32(_a_F_heap_form_minimal_tuple_0) | v15
	*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)) = uint16(v86)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v78 + v66
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v67
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v95 = v78 + int32(15)
	goto L19
L18:
	;
	v95 = int32(0)
	goto L19
L19:
	;
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v99 = v95 - int32(1)
	goto L22
L21:
	;
	v99 = int32(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v99
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+12)))
	v103 = v101 & int32(_a_F_heap_form_minimal_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v78)+12)) = uint16(v103)
	if int32(0) < v90 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v117 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	m.G0 = v13 + int32(32)
	return v78
L26:
	;
	v123 = v117 << (uint(int32(3)) % 32)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v128 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v129 = v13 + int32(24)
	goto L30
L29:
	;
	v129 = int32(0)
	goto L30
L30:
	;
	if l1 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l1+v123)))
	v137 = v135
	goto L33
L32:
	;
	v137 = int64(0)
	goto L33
L33:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v117))))
	v141 = v139
	goto L36
L35:
	;
	v141 = int32(1)
	goto L36
L36:
	;
	F_fill_val(m, v123+(l0+int32(28)), v129, v13+int32(20), v13+int32(28), v78+int32(12), v137, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v145 = v117 + int32(1)
	if v145 != v90 {
		v117 = v145
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(1664)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
	F_errmsg(m, int32(_a_F_heap_form_minimal_tuple_2), v13)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_heap_form_minimal_tuple_3), int32(1410), int32(_a_F_heap_form_minimal_tuple_4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
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
	var v30 int32
	_ = v30
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	v30 = v25
	v32 = v25
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24+v30<<(uint(int32(2))%32))))
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
		v30 = v162
		v32 = v160
		goto L4
	} else {
		goto L37
	}
L7:
	;
	if base.Ui32(v20) <= base.Ui32((v82-int32(1))&int32(_a_F_heap_get_root_tuples_1)) {
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v82 = v38 & int32(_a_F_heap_get_root_tuples_2)
	v84 = int32(0)
	goto L7
L9:
	;
	v47 = l0 + v38&int32(_a_F_heap_get_root_tuples_2)
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+18)))
	if v48 < int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v30<<(uint(int32(1))%32)-int32(2)))) = uint16(v32)
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
	if v62&int32(_a_F_heap_get_root_tuples_3) == int32(_a_F_heap_get_root_tuples_4) {
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
	v82 = v70
	v84 = v77
	goto L7
L16:
	;
	return
L17:
	;
	v82 = v70
	v84 = v75
	goto L7
L18:
	;
	v93 = v82
	v94 = v84
	goto L19
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24+v93<<(uint(int32(2))%32))))
	if v101&int32(_a_F_heap_get_root_tuples_5) != int32(_a_F_heap_get_root_tuples_6) {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	goto L6
L21:
	;
	v108 = l0 + v101&int32(_a_F_heap_get_root_tuples_2)
	if v94 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+20)))
	v110 = int32(768)
	if v109&v110 != v110 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v93<<(uint(int32(1))%32)-int32(2)))) = uint16(v32)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+19)))
	if v124&int32(64) == int32(0) {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v116 = v114
	goto L27
L26:
	;
	v116 = int32(2)
	goto L27
L27:
	;
	if v116 != v94 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+20)))
	if v129&int32(2048)|base.B2i32(v129&int32(768) == int32(512)) != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+16)))
	if v129&int32(_a_F_heap_get_root_tuples_3) == int32(_a_F_heap_get_root_tuples_4) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if base.Ui32((v137-int32(1))&int32(_a_F_heap_get_root_tuples_1)) < base.Ui32(v20) {
		v93 = v137
		v94 = v145
		goto L19
	} else {
		goto L36
	}
L32:
	;
	v142 = F_HeapTupleGetUpdateXid(m, v108)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L16
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v145 = v144
	goto L31
L35:
	;
	v145 = v142
	goto L31
L36:
	;
	goto L20
L37:
	;
	goto L5
}
func F_heap_getattr_3(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)))
	if base.Ui32(v12&int32(2047)) <= base.Ui32(int32(20)) {
		v18 = F_getmissingattr(m, l1, int32(21), l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v80 = v18
			m.G0 = v9 + int32(16)
			return v80
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+188)))
			if int32(0) <= v30 {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v35 = v24 + v33 + v30
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+192)))
				if v36 == int32(1) {
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+190)))
					if base.I32_popcnt(v39) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
							F_errmsg_internal(m, int32(_a_F_heap_getattr_3_0), v9)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_3_1), int32(123), int32(_a_F_heap_getattr_3_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v39) {
						case 0:
							v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35))))
							v80 = v44
							m.G0 = v9 + int32(16)
							return v80
						case 1:
							v45 = int64(*(*int16)(unsafe.Add(mBase, uint32(v35))))
							v80 = v45
							m.G0 = v9 + int32(16)
							return v80
						case 2:
							v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35))))
							v80 = v46
							m.G0 = v9 + int32(16)
							return v80
						case 3:
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
							v80 = v47
							m.G0 = v9 + int32(16)
							return v80
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
								F_errmsg_internal(m, int32(_a_F_heap_getattr_3_0), v9)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_3_1), int32(123), int32(_a_F_heap_getattr_3_2))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
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
					v80 = base.I64_extend_i32_u(v35)
					m.G0 = v9 + int32(16)
					return v80
				}
			} else {
				v63 = F_nocachegetattr(m, l0, int32(21), l1)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					v80 = v63
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		} else {
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+25)))
			if v65&int32(16) == int32(0) {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v70)
				v80 = int64(0)
				m.G0 = v9 + int32(16)
				return v80
			} else {
				v74 = F_nocachegetattr(m, l0, int32(21), l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					v80 = v74
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		}
	}
}
func F_heap_getattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v87 int64
	_ = v87
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)))
	if base.Ui32(v14&int32(2047)) < base.Ui32(l1) {
		v18 = F_getmissingattr(m, l2, l1, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v87 = v18
			m.G0 = v11 + int32(16)
			return v87
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v34 = l2 + l1<<(uint(int32(3))%32) + int32(20)
			v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34))))
			if v35 < int32(0) {
				v80 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
				} else {
					v87 = v80
					m.G0 = v11 + int32(16)
					return v87
				}
			} else {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v40 = v24 + v38 + v35
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
				if v41 == int32(1) {
					v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+2)))
					if base.I32_popcnt(v44) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
							F_errmsg_internal(m, int32(_a_F_heap_getattr_5_0), v11)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_5_1), int32(123), int32(_a_F_heap_getattr_5_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v44) {
						case 0:
							v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v40))))
							v87 = v49
							m.G0 = v11 + int32(16)
							return v87
						case 1:
							v50 = int64(*(*int16)(unsafe.Add(mBase, uint32(v40))))
							v87 = v50
							m.G0 = v11 + int32(16)
							return v87
						case 2:
							v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40))))
							v87 = v51
							m.G0 = v11 + int32(16)
							return v87
						case 3:
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
							v87 = v52
							m.G0 = v11 + int32(16)
							return v87
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
								F_errmsg_internal(m, int32(_a_F_heap_getattr_5_0), v11)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_5_1), int32(123), int32(_a_F_heap_getattr_5_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
					v87 = base.I64_extend_i32_u(v40)
					m.G0 = v11 + int32(16)
					return v87
				}
			}
		} else {
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+26)))
			v68 = int32(1)
			if int32(base.Ui32(v67)>>(uint((l1-v68)&int32(7))%32))&v68 != 0 {
				v80 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
				} else {
					v87 = v80
					m.G0 = v11 + int32(16)
					return v87
				}
			} else {
				v75 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v75)
				v87 = int64(0)
				m.G0 = v11 + int32(16)
				return v87
			}
		}
	}
}
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v54 int64
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
	var v71 int64
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
	var v106 int64
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = F_palloc_mul(m, int32(8), v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = F_palloc_mul(m, int32(1), v15)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_heap_deform_tuple(m, l0, l1, v16, v21)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v15 <= int32(0) {
				} else {
					if v15 != int32(1) {
						v39 = v6
						v45 = v6
						for {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v39))))
							if v47 == int32(1) {
								v51 = v39 << (uint(int32(3)) % 32)
								v54 = *(*int64)(unsafe.Add(mBase, uint32(l2+v51)))
								*(*int64)(unsafe.Add(mBase, uint32(v16+v51))) = v54
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v39))))
								*(*uint8)(unsafe.Add(mBase, uint32(v39+v21))) = uint8(v58)
							} else {
							}
							v61 = int32(1)
							v62 = v39 | v61
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v62))))
							if v64 == v61 {
								v68 = v62 << (uint(int32(3)) % 32)
								v71 = *(*int64)(unsafe.Add(mBase, uint32(l2+v68)))
								*(*int64)(unsafe.Add(mBase, uint32(v16+v68))) = v71
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v62))))
								*(*uint8)(unsafe.Add(mBase, uint32(v62+v21))) = uint8(v75)
							} else {
							}
							v78 = int32(2)
							v79 = v39 + v78
							v81 = v45 + v78
							if v81 != v15&int32(2147483646) {
								v39 = v79
								v45 = v81
								continue
							} else {
								break
							}
							break
						}
						if v15&int32(1) == int32(0) {
						} else {
							v91 = v79
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v91))))
							if v99 != int32(1) {
							} else {
								v103 = v91 << (uint(int32(3)) % 32)
								v106 = *(*int64)(unsafe.Add(mBase, uint32(l2+v103)))
								*(*int64)(unsafe.Add(mBase, uint32(v16+v103))) = v106
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v91))))
								*(*uint8)(unsafe.Add(mBase, uint32(v91+v21))) = uint8(v110)
							}
						}
					} else {
						v91 = v6
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v91))))
						if v99 != int32(1) {
						} else {
							v103 = v91 << (uint(int32(3)) % 32)
							v106 = *(*int64)(unsafe.Add(mBase, uint32(l2+v103)))
							*(*int64)(unsafe.Add(mBase, uint32(v16+v103))) = v106
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v91))))
							*(*uint8)(unsafe.Add(mBase, uint32(v91+v21))) = uint8(v110)
						}
					}
				}
				v125 = F_heap_form_tuple(m, l1, v16, v21)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v16)
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_palloc_mul(m, int32(8), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_palloc_mul(m, int32(1), v18)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_heap_deform_tuple(m, l0, l1, v19, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
		goto L16
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
	v79 = F_heap_form_tuple(m, l1, v19, v24)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3+v37<<(uint(int32(2))%32))))
	if base.B2i32(v45 <= int32(0))|base.B2i32(v18 < v45) != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v50 = int32(1)
	v51 = v45 - v50
	v52 = int32(3)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l4+v37<<(uint(v52)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v19+v51<<(uint(v52)%32)))) = v58
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v24))) = uint8(v62)
	v65 = v37 + v50
	if v65 != l2 {
		v37 = v65
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_pfree(m, v19)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v24)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)) = uint16(v87)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v89
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)) = uint16(v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v95
	m.G0 = v15 + int32(16)
	return v79
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v45
	F_errmsg_internal(m, int32(_a_F_heap_modify_tuple_by_cols_0), v15)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_heap_modify_tuple_by_cols_1), int32(1213), int32(_a_F_heap_modify_tuple_by_cols_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
			*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
		} else {
			v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)))
			v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+2)))
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
			v121 = int32(16)
			v123 = v119 | v120<<(uint(v121)%32)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v123
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v117 | v118<<(uint(v121)%32) - v123 + int32(1)
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
func F_heap_tuple_infomask_flags(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int64
	_ = v450
	var v451 int32
	_ = v451
	v2 = int32(0)
	v9 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v9
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)) = uint16(v2)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_superuser(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = base.I64_extend_i32_u(v436)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v447 = F_heap_form_tuple(m, v442, v13+int32(16), v13+int32(14))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L127
	}
L2:
	;
	F_pfree(m, v77)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L126
	}
L3:
	;
	v430 = F_construct_array_builtin(m, v77, v428, int32(25))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L122
	}
L5:
	;
	return int64(0)
L6:
	;
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v30 = F_get_call_result_type(m, l0, int32(0), v13+int32(8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
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
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L118
	}
L10:
	;
	if v30 != int32(1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v34 = int64(8)
	v36 = base.I32_wrap_i64(int64(base.Ui64(v21) >> (uint(v34) % 64)))
	v37 = int32(255)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36&v37)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v21)&v37)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v50 = base.I32_wrap_i64(int64(base.Ui64(v22) >> (uint(v34) % 64)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50&v37)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v56 = base.I32_wrap_i64(v22)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56&v37)+uint32(_c_F_heap_tuple_infomask_flags[0]))))
	v64 = v41 + (v47 + (v55 + v61))
	if v64 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v68 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v77 = F_palloc0_mul(m, int32(8), v64)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v68)
	v73 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v436 = v73
	goto L1
L17:
	;
	if v56&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v87 = int32(0)
	goto L20
L20:
	;
	if v56&int32(2) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_extend_i32_u(v82)
	v87 = int32(1)
	goto L20
L22:
	;
	v94 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	v100 = v87
	goto L24
L24:
	;
	if v56&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v87<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v94)
	v100 = v87 + int32(1)
	goto L24
L26:
	;
	v107 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_5))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	v113 = v100
	goto L28
L28:
	;
	if v56&int32(8) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v100<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v107)
	v113 = v100 + int32(1)
	goto L28
L30:
	;
	v120 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_6))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v126 = v113
	goto L32
L32:
	;
	if v56&int32(16) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v113<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v120)
	v126 = v113 + int32(1)
	goto L32
L34:
	;
	v133 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_7))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v139 = v126
	goto L36
L36:
	;
	if v56&int32(32) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v126<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v133)
	v139 = v126 + int32(1)
	goto L36
L38:
	;
	v146 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_8))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v152 = v139
	goto L40
L40:
	;
	if v56&int32(64) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v139<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v146)
	v152 = v139 + int32(1)
	goto L40
L42:
	;
	v159 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_9))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v165 = v152
	goto L44
L44:
	;
	if v56&int32(128) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v152<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v159)
	v165 = v152 + int32(1)
	goto L44
L46:
	;
	v172 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_10))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	v178 = v165
	goto L48
L48:
	;
	if v50&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v165<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v172)
	v178 = v165 + int32(1)
	goto L48
L50:
	;
	v185 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_11))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	v191 = v178
	goto L52
L52:
	;
	if v50&int32(2) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v178<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v185)
	v191 = v178 + int32(1)
	goto L52
L54:
	;
	v198 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_12))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	v204 = v191
	goto L56
L56:
	;
	if v50&int32(4) != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v191<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v198)
	v204 = v191 + int32(1)
	goto L56
L58:
	;
	v211 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_13))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	v217 = v204
	goto L60
L60:
	;
	if v50&int32(8) != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v204<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v211)
	v217 = v204 + int32(1)
	goto L60
L62:
	;
	v224 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_14))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L65
	}
L63:
	;
	v230 = v217
	goto L64
L64:
	;
	if v50&int32(16) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v217<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v224)
	v230 = v217 + int32(1)
	goto L64
L66:
	;
	v237 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_15))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	v243 = v230
	goto L68
L68:
	;
	if v50&int32(32) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v230<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v237)
	v243 = v230 + int32(1)
	goto L68
L70:
	;
	v250 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_16))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	v256 = v243
	goto L72
L72:
	;
	if v50&int32(64) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v243<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v250)
	v256 = v243 + int32(1)
	goto L72
L74:
	;
	v263 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_17))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	v269 = v256
	goto L76
L76:
	;
	v271 = v50 << (uint(int32(8)) % 32)
	if base.I32_extend16_s(v271) < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v256<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v263)
	v269 = v256 + int32(1)
	goto L76
L78:
	;
	v279 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_18))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L81
	}
L79:
	;
	v285 = v269
	goto L80
L80:
	;
	if v36&int32(32) != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v269<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v279)
	v285 = v269 + int32(1)
	goto L80
L82:
	;
	v292 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_19))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	v298 = v285
	goto L84
L84:
	;
	if v36&int32(64) != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v285<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v292)
	v298 = v285 + int32(1)
	goto L84
L86:
	;
	v305 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_20))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L89
	}
L87:
	;
	v311 = v298
	goto L88
L88:
	;
	if v36&int32(128) != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v298<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v305)
	v311 = v298 + int32(1)
	goto L88
L90:
	;
	v318 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_21))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L93
	}
L91:
	;
	v324 = v311
	goto L92
L92:
	;
	v326 = F_construct_array_builtin(m, v77, v324, int32(25))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L94
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v311<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v318)
	v324 = v311 + int32(1)
	goto L92
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v326)
	v330 = int32(3)
	v331 = v64 << (uint(v330) % 32)
	if v77&v330|base.B2i32(base.Ui32(int32(128)) < base.Ui32(v64)) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v341 = v77 + v331
	v343 = v77 + int32(4)
	if base.Ui32(v343) < base.Ui32(v341) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v352 = v331
	goto L97
L97:
	;
	if v352 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v345 = v341
	goto L100
L99:
	;
	v345 = v343
	goto L100
L100:
	;
	v352 = (v77^int32(-1)+v345)&int32(-4) + int32(4)
	goto L97
L101:
	;
	base.MemoryFill(m, v77, int32(0), v352)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v356 = int32(80)
	if v56&v356 == v356 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v361 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_22))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L107
	}
L105:
	;
	v366 = int32(0)
	goto L106
L106:
	;
	v367 = int32(3)
	if v50&v367 == v367 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_extend_i32_u(v361)
	v366 = int32(1)
	goto L106
L108:
	;
	v375 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_23))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	v381 = v366
	goto L110
L110:
	;
	if base.Ui32(int32(_a_F_heap_tuple_infomask_flags_24)) <= base.Ui32(v271&int32(_a_F_heap_tuple_infomask_flags_25)) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v366<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v375)
	v381 = v366 + int32(1)
	goto L110
L112:
	;
	v390 = F_cstring_to_text(m, int32(_a_F_heap_tuple_infomask_flags_26))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v381 != 0 {
		v428 = v381
		goto L3
	} else {
		goto L116
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v381<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v390)
	v428 = v381 + int32(1)
	goto L3
L116:
	;
	v397 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v433 = v397
	goto L2
L118:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_heap_tuple_infomask_flags_27), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_heap_tuple_infomask_flags_1), int32(530), int32(_a_F_heap_tuple_infomask_flags_2))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
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
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_heap_tuple_infomask_flags_1), int32(534), int32(_a_F_heap_tuple_infomask_flags_2))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
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
	v433 = v430
	goto L2
L126:
	;
	v436 = v433
	goto L1
L127:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	v450 = F_HeapTupleHeaderGetDatum(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	m.G0 = v13 + int32(32)
	return v450
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
