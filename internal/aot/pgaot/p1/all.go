package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetAllNullSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v5 == int32(0) {
		v8 = int32(_a_F_ExecGetAllNullSlot_0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetAllNullSlot[0]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecGetAllNullSlot[0])) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
		v15 = F_table_slot_callbacks(m, v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_ExecInitExtraTupleSlot(m, l0, v14, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_ExecStoreAllNullTuple(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v19
					*(*int32)(unsafe.Add(mBase, _c_F_ExecGetAllNullSlot[0])) = v9
					v26 = v19
					return v26
				}
			}
		}
	} else {
		v26 = v5
		return v26
	}
}
func F_LockReleaseAll(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v318 int64
	_ = v318
	var v322 int64
	_ = v322
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int64
	_ = v336
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v347 int32
	_ = v347
	var v352 int64
	_ = v352
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v546 int32
	_ = v546
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v795 int32
	_ = v795
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	if base.Ui32(int32(_a_F_LockReleaseAll_0)) < base.Ui32((l0-int32(3))&int32(_a_F_LockReleaseAll_1)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_LockReleaseAll[0])))
	if l0 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L6
	} else {
		goto L120
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v82 = v28 + int32(12)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[1]))
	F_hash_seq_init(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L11
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v46 = F_LWLockAcquire(m, v42+int32(584), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+612))
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+612)) = v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+608)))
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+608)) = uint8(v51)
	F_LWLockRelease(m, v49+int32(584))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v53|base.B2i32(v50 == int32(0)) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v50
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v67
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	F_LockRefindAndRelease(m, int32(_a_F_LockReleaseAll_2), v71, v28+int32(12), int32(7), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v87 = F_hash_seq_search(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v483 = int32(2)
	v485 = v80 + int32(1)
	if v485 <= v483 {
		goto L78
	} else {
		goto L79
	}
L13:
	;
	if v87 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v93 = v87
	v96 = v3
	goto L15
L15:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v93)+32))
	if v116 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v426 == int32(0) {
		goto L12
	} else {
		goto L76
	}
L17:
	;
	v448 = F_hash_seq_search(m, v28+int32(12))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L74
	}
L18:
	;
	F_RemoveLocalLock(m, v93)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L73
	}
L19:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+15)))
	if l0 != v119 {
		v426 = v96
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L70
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+40)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v172
	v426 = v96
	goto L17
L23:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v93)+40))
	if v123 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	if v227 != 0 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+40)) = int32(0)
	goto L25
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v131 = int32(0)
	goto L28
L28:
	;
	v155 = v126 + v131<<(uint(int32(4))%32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v156 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v167 <= int32(0) {
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v166 = v131 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v93)+40))
	if v166 < v167 {
		v131 = v166
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v155)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v159
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v161
	goto L30
L32:
	;
	goto L33
L33:
	;
	F_ResourceOwnerForgetLock(m, v156, v93)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	goto L29
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v171 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v126)+8))
	if int64(0) < v172 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	goto L26
L39:
	;
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v93)+32))
	if v368 <= int64(0) {
		goto L18
	} else {
		goto L69
	}
L40:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v93)+24))
	if v228 != 0 {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+15)))
	if v229 != int32(1) {
		goto L21
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+14)))
	if v232 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[4]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if base.B2i32(v234 != v235)|base.B2i32(v234 == int32(0)) != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	if int32(4) <= v240 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	if v96 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v250 = F_LWLockAcquire(m, v246+int32(584), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v252 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[5]))
	v255 = int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v260 = (v254 - v255) & (v257 * int32(_a_F_LockReleaseAll_3))
	v262 = v260 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+uint32(_c_F_LockReleaseAll[6]))) = v252
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v277 = v260 & int32(268435455) << (uint(int32(3)) % 32)
	v281 = v272
	v283 = v252
	v299 = int64(0)
	goto L52
L51:
	;
	goto L50
L52:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v281)+604))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303+v260<<(uint(int32(6))%32)+base.I32_wrap_i64(v299)<<(uint(int32(2))%32))))
	if v257 != v309 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v334 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+600))
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v339+v277)))
	if int64(base.Ui64(v341)>>(uint(v336)%64))&int64(7) != int64(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v333 = v281
	v334 = v283
	v336 = v299 * int64(3)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v281)+600))
	v314 = v313 + v277
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v314)))
	v318 = v299 * int64(3)
	v322 = int64(1) << (uint(base.I64_extend_i32_u(v240-v255+base.I32_wrap_i64(v318))) % 64)
	if v315&v322 == int64(0) {
		v333 = v281
		v334 = v283
		v336 = v318
		goto L54
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v314))) = v315 & (v322 ^ int64(-1))
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v333 = v331
	v334 = int32(1)
	v336 = v318
	goto L54
L59:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v262)+uint32(_c_F_LockReleaseAll[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+uint32(_c_F_LockReleaseAll[6]))) = v347 + int32(1)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v352 = v299 + int64(1)
	if v352 != int64(16) {
		v281 = v333
		v283 = v334
		v299 = v352
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
L63:
	;
	F_LWLockRelease(m, v333+int32(584))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_RemoveLocalLock(m, v93)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L68
	}
L66:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	F_LockRefindAndRelease(m, v38, v362, v93, v240, int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v426 = v334
	goto L17
L69:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = v371 | int32(1)<<(uint(v373)%32)
	goto L18
L70:
	;
	F_errmsg_internal(m, int32(_a_F_LockReleaseAll_4), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_LockReleaseAll_5), int32(2385), int32(_a_F_LockReleaseAll_6))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v426 = v96
	goto L17
L74:
	;
	if v448 != 0 {
		v93 = v448
		v96 = v426
		goto L15
	} else {
		goto L75
	}
L75:
	;
	goto L16
L76:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	F_LWLockRelease(m, v453+int32(584))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	goto L12
L78:
	;
	v488 = v483
	goto L80
L79:
	;
	v488 = v485
	goto L80
L80:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[7]))
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v495 = v490
	v496 = v492
	v508 = v3
	goto L81
L81:
	;
	v520 = v496 + v508<<(uint(int32(3))%32)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+152))
	if v521 == int32(0) {
		v771 = v495
		v772 = v496
		goto L83
	} else {
		goto L84
	}
L82:
	;
	m.G0 = v28 + int32(32)
	return
L83:
	;
	v795 = v508 + int32(1)
	if v795 != int32(16) {
		v495 = v771
		v496 = v772
		v508 = v795
		goto L81
	} else {
		goto L119
	}
L84:
	;
	v525 = v520 + int32(148)
	if v521 == v525 {
		v771 = v495
		v772 = v496
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v531 = v495 + v508<<(uint(int32(7))%32) + int32(_a_F_LockReleaseAll_7)
	v533 = F_LWLockAcquire(m, v531, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v536 = int32(0)
	if base.B2i32(v535 == v536)|base.B2i32(v535 == v525) == v536 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v546 = v535
	goto L90
L88:
	;
	goto L89
L89:
	;
	F_LWLockRelease(m, v531)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L6
	} else {
		goto L118
	}
L90:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v569 = v546 - int32(28)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+15)))
	if l0 != v571 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L89
L92:
	;
	if v567 != v525 {
		v546 = v567
		goto L90
	} else {
		goto L117
	}
L93:
	;
	if l1 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v584 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v546-int32(12))))
	v584 = v577
	goto L94
L96:
	;
	goto L97
L97:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v546-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v546-int32(12)))) = v582
	v584 = v582
	goto L94
L98:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v546-int32(16))))
	if v589 != 0 {
		goto L92
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v591 = v546 - int32(12)
	if v80 <= int32(0) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = int32(0)
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[8]))
	v706 = F_get_hash_value(m, v705, v570)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L115
	}
L103:
	;
	v679 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v596 = v546 - int32(16)
	v605 = int32(0)
	v606 = int32(1)
	goto L106
L106:
	;
	v629 = int32(1) << (uint(v606) % 32)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	if v629&v630 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v679 = v670
	goto L102
L108:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v570)+84))
	v633 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v570)+84)) = v632 - v633
	v637 = v606 << (uint(int32(2)) % 32)
	v638 = v570 + int32(44) + v637
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v639 - v633
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v570)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v570)+128)) = v643 - v633
	v647 = v637 + (v570 + int32(88))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	v650 = v648 - v633
	*(*int32)(unsafe.Add(mBase, uint32(v647))) = v650
	v653 = v629 ^ int32(-1)
	if v650 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v670 = v605
	goto L110
L110:
	;
	v675 = v606 + int32(1)
	if v675 != v488 {
		v605 = v670
		v606 = v675
		goto L106
	} else {
		goto L114
	}
L111:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v570)+16)) = v656 & v653
	goto L113
L112:
	;
	goto L113
L113:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v659+v637)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v570)+20))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = v663 & v653
	v670 = v605 | base.B2i32(v662&v661 != int32(0))
	goto L110
L114:
	;
	goto L107
L115:
	;
	F_CleanUpLock(m, v570, v569, v38, v706, v679&int32(1))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	goto L92
L117:
	;
	goto L91
L118:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[7]))
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_LockReleaseAll[2]))
	v771 = v766
	v772 = v768
	goto L83
L119:
	;
	goto L82
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l0
	F_errmsg_internal(m, int32(_a_F_LockReleaseAll_8), v28)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_LockReleaseAll_5), int32(2287), int32(_a_F_LockReleaseAll_6))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strip_all_phvs_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v3 = int32(0)
	if l0 == v3 {
		v20 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v20
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 == int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v14 = F_expression_tree_mutator_impl(m, v6, int32(1489), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v12 != 0 {
		v6 = v12
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	v20 = v3
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v20 = v14
	goto L1
}
