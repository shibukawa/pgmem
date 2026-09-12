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
		v8 = int32(4486928)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
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
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v323 int64
	_ = v323
	var v327 int64
	_ = v327
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v352 int32
	_ = v352
	var v357 int64
	_ = v357
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v432 int32
	_ = v432
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v592 int32
	_ = v592
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v831 int32
	_ = v831
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	if base.Ui32(int32(65533)) < base.Ui32((l0-int32(3))&int32(65535)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[804])))
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
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L125
	}
L4:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v86 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_hash_seq_init(m, v28+int32(12), v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L12
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v48 = F_LWLockAcquire(m, v44+int32(584), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+612))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+612)) = v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+608)))
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+608)) = uint8(v53)
	F_LWLockRelease(m, v51+int32(584))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v55 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v52 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v52
	v68 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_LockRefindAndRelease(m, int32(1609880), v72, v28+int32(12), int32(7), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v91 = F_hash_seq_search(m, v28+int32(12))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v522 = int32(2)
	v524 = v82 + int32(1)
	if v524 <= v522 {
		goto L83
	} else {
		goto L84
	}
L14:
	;
	if v91 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v97 = v91
	v104 = v3
	goto L16
L16:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v97)+32))
	if v120 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v471&int32(1) == int32(0) {
		goto L13
	} else {
		goto L81
	}
L18:
	;
	goto L17
L19:
	;
	v460 = F_hash_seq_search(m, v28+int32(12))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L79
	}
L20:
	;
	F_RemoveLocalLock(m, v97)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L78
	}
L21:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+15)))
	if l0 != v123 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L75
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+40)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v97)+32)) = v176
	goto L19
L25:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v97)+40))
	if v127 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	if v231 != 0 {
		goto L42
	} else {
		goto L43
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+40)) = int32(0)
	goto L27
L29:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v97)+48))
	v135 = int32(0)
	goto L30
L30:
	;
	v159 = v130 + v135<<(uint(int32(4))%32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v160 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v171 <= int32(0) {
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v170 = v135 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v97)+40))
	if v170 < v171 {
		v135 = v170
		goto L30
	} else {
		goto L37
	}
L33:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v130)+8)) = v165
	goto L32
L34:
	;
	goto L35
L35:
	;
	F_ResourceOwnerForgetLock(m, v160, v97)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	goto L31
L38:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v175 != 0 {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
	if int64(0) < v176 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L28
L41:
	;
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v97)+32))
	if v379 <= int64(0) {
		goto L20
	} else {
		goto L74
	}
L42:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	if v232 != 0 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+15)))
	if v233 != int32(1) {
		goto L23
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+14)))
	if v236 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v238 != v239 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	if v238 == int32(0) {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if int32(4) <= v243 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	if v104&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v255 = F_LWLockAcquire(m, v251+int32(584), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v257 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, _consts[201]))
	v260 = int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v265 = (v259 - v260) & (v262 * int32(49157))
	v267 = v265 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v267)+uint32(_consts[805]))) = v257
	v277 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v282 = v265 & int32(268435455) << (uint(int32(3)) % 32)
	v286 = v277
	v292 = v257
	v304 = int64(0)
	goto L55
L54:
	;
	goto L53
L55:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v286)+604))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308+base.I32_wrap_i64(v304)<<(uint(int32(2))%32)+v265<<(uint(int32(6))%32))))
	if v262 != v314 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v340&int32(1) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+600))
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v344+v282)))
	if int64(base.Ui64(v346)>>(uint(v341)%64))&int64(7) != int64(0) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v338 = v286
	v340 = v292
	v341 = v304 * int64(3)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v286)+600))
	v319 = v318 + v282
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v319)))
	v323 = v304 * int64(3)
	v327 = int64(1) << (uint(base.I64_extend_i32_u(v243-v260+base.I32_wrap_i64(v323))) % 64)
	if v320&v327 == int64(0) {
		v338 = v286
		v340 = v292
		v341 = v323
		goto L57
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v320 & (v327 ^ int64(-1))
	v336 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v338 = v336
	v340 = int32(1)
	v341 = v323
	goto L57
L62:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v267)+uint32(_consts[805])))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+uint32(_consts[805]))) = v352 + int32(1)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v357 = v304 + int64(1)
	if v357 != int64(16) {
		v286 = v338
		v292 = v340
		v304 = v357
		goto L55
	} else {
		goto L65
	}
L65:
	;
	goto L56
L66:
	;
	F_LWLockRelease(m, v338+int32(584))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_RemoveLocalLock(m, v97)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_LockRefindAndRelease(m, v40, v369, v97, v243, int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v377 = F_hash_seq_search(m, v28+int32(12))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	if v377 != 0 {
		v97 = v377
		v104 = v340
		goto L16
	} else {
		goto L73
	}
L73:
	;
	v471 = v340
	goto L18
L74:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v382 | int32(1)<<(uint(v384)%32)
	goto L20
L75:
	;
	F_errmsg_internal(m, int32(442284), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(494492), int32(2385), int32(303641))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	goto L19
L79:
	;
	if v460 != 0 {
		v97 = v460
		goto L16
	} else {
		goto L80
	}
L80:
	;
	v471 = v104
	goto L18
L81:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_LWLockRelease(m, v492+int32(584))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L13
L83:
	;
	v527 = v522
	goto L85
L84:
	;
	v527 = v524
	goto L85
L85:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v531 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v536 = v531
	v537 = v529
	v545 = v3
	goto L86
L86:
	;
	v559 = v536 + v545<<(uint(int32(3))%32)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)+152))
	if v560 == int32(0) {
		v809 = v536
		v810 = v537
		goto L88
	} else {
		goto L89
	}
L87:
	;
	m.G0 = v28 + int32(32)
	return
L88:
	;
	v831 = v545 + int32(1)
	if v831 != int32(16) {
		v536 = v809
		v537 = v810
		v545 = v831
		goto L86
	} else {
		goto L124
	}
L89:
	;
	v564 = v559 + int32(148)
	if v560 == v564 {
		v809 = v536
		v810 = v537
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v570 = v537 + v545<<(uint(int32(7))%32) + int32(23296)
	v572 = F_LWLockAcquire(m, v570, int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	if v574 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_LWLockRelease(m, v570)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L123
	}
L93:
	;
	if v574 == v564 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v592 = v574
	goto L95
L95:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	v605 = v592 - int32(28)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+15)))
	if l0 != v607 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L92
L97:
	;
	if v603 != v564 {
		v592 = v603
		goto L95
	} else {
		goto L122
	}
L98:
	;
	if l1 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v620 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v592-int32(12))))
	v620 = v613
	goto L99
L101:
	;
	goto L102
L102:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v592-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v592-int32(12)))) = v618
	v620 = v618
	goto L99
L103:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v592-int32(16))))
	if v625 != 0 {
		goto L97
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v627 = v592 - int32(12)
	if v82 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v627))) = int32(0)
	v741 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	v742 = F_get_hash_value(m, v741, v606)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L6
	} else {
		goto L120
	}
L108:
	;
	v720 = int32(0)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v632 = v592 - int32(16)
	v642 = int32(1)
	v646 = int32(0)
	goto L111
L111:
	;
	v665 = int32(1) << (uint(v642) % 32)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	if v665&v666 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v720 = v708
	goto L107
L113:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v606)+84))
	v669 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v606)+84)) = v668 - v669
	v673 = v642 << (uint(int32(2)) % 32)
	v674 = v606 + int32(44) + v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v675 - v669
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v606)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+128)) = v679 - v669
	v683 = v673 + (v606 + int32(88))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	v686 = v684 - v669
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = v686
	v689 = v665 ^ int32(-1)
	if v686 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v708 = v646
	goto L115
L115:
	;
	v711 = v642 + int32(1)
	if v711 != v527 {
		v642 = v711
		v646 = v708
		goto L111
	} else {
		goto L119
	}
L116:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v606)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+16)) = v692 & v689
	goto L118
L117:
	;
	goto L118
L118:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v695+v673)))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v606)+20))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = v699 & v689
	v708 = v646 | base.B2i32(v697&v698 != int32(0))
	goto L115
L119:
	;
	goto L112
L120:
	;
	F_CleanUpLock(m, v606, v605, v40, v742, v720&int32(1))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	goto L97
L122:
	;
	goto L96
L123:
	;
	v802 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v804 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v809 = v804
	v810 = v802
	goto L88
L124:
	;
	goto L87
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l0
	F_errmsg_internal(m, int32(484004), v28)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(494492), int32(2287), int32(303641))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
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
	v14 = F_expression_tree_mutator_impl(m, v6, int32(1505), l1)
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
