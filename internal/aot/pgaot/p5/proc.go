package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcArrayRemove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	v11 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v17 = F_LWLockAcquire(m, v13+int32(512), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v24 = F_LWLockAcquire(m, v20+int32(384), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
	v30 = base.I32_wrap_i64(v29)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v30)) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82+v26))) = uint8(v84)
	v87 = v11 + int32(36)
	v88 = int32(2)
	v89 = v26 << (uint(v88) % 32)
	v90 = v87 + v89
	v92 = v26 + int32(1)
	v94 = v92 << (uint(v88) % 32)
	v95 = v87 + v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v99 = v96 + (v26 ^ int32(-1))
	v101 = v99 << (uint(v88) % 32)
	if v90 == v95 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v42 = base.B2i32(base.Ui32(v30) < base.Ui32(l1))
	goto L7
L9:
	;
	goto L10
L10:
	;
	v42 = int32(base.Ui32(v30-l1) >> (uint(int32(31)) % 32))
	goto L7
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = v29 + base.I64_extend_i32_s(l1-v30)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v49 + int64(1)
	v53 = int32(4444000)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55+v26<<(uint(int32(2))%32)))) = v59
	v62 = v26 << (uint(int32(1)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v65)+1)) = uint8(v59)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v62))) = uint8(v59)
	goto L6
L14:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v249 = v89 + v248
	v250 = v248 + v94
	if v249 == v250 {
		goto L61
	} else {
		goto L62
	}
L15:
	;
	goto L14
L16:
	;
	v105 = v90 + v101
	if base.Ui32(v95-v105) <= base.Ui32(int32(0)-v101<<(uint(int32(1))%32)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v112 = F___memcpy(m, v90, v95, v101)
	mBase = m.M
	goto L14
L18:
	;
	goto L19
L19:
	;
	v115 = (v90 ^ v95) & int32(3)
	if base.Ui32(v90) < base.Ui32(v95) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v217 == int32(0) {
		goto L15
	} else {
		goto L56
	}
L21:
	;
	if base.Ui32(v195) <= base.Ui32(int32(3)) {
		v216 = v194
		v217 = v195
		v218 = v196
		goto L20
	} else {
		goto L52
	}
L22:
	;
	if v115 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v115 != 0 {
		v177 = v101
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v216 = v95
	v217 = v101
	v218 = v90
	goto L20
L26:
	;
	goto L27
L27:
	;
	if v90&int32(3) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v194 = v95
	v195 = v101
	v196 = v90
	goto L21
L29:
	;
	goto L30
L30:
	;
	v122 = v95
	v123 = v101
	v124 = v90
	goto L31
L31:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v194 = v131
	v195 = v133
	v196 = v135
	goto L21
L33:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v128)
	v130 = int32(1)
	v131 = v122 + v130
	v133 = v123 - v130
	v135 = v124 + v130
	if v135&int32(3) != 0 {
		v122 = v131
		v123 = v133
		v124 = v135
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v177 == int32(0) {
		goto L15
	} else {
		goto L48
	}
L36:
	;
	if v105&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = v101
	goto L40
L38:
	;
	v157 = v101
	goto L39
L39:
	;
	if base.Ui32(v157) <= base.Ui32(int32(3)) {
		v177 = v157
		goto L35
	} else {
		goto L44
	}
L40:
	;
	if v142 == int32(0) {
		goto L15
	} else {
		goto L42
	}
L41:
	;
	v157 = v148
	goto L39
L42:
	;
	v148 = v142 - int32(1)
	v149 = v90 + v148
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v148))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v151)
	if v149&int32(3) != 0 {
		v142 = v148
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v164 = v157
	goto L45
L45:
	;
	v168 = v164 - int32(4)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v95+v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v168))) = v171
	if base.Ui32(int32(3)) < base.Ui32(v168) {
		v164 = v168
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v177 = v168
	goto L35
L47:
	;
	goto L46
L48:
	;
	v184 = v177
	goto L49
L49:
	;
	v188 = v184 - int32(1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v188))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v188))) = uint8(v191)
	if v188 != 0 {
		v184 = v188
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L15
L51:
	;
	goto L50
L52:
	;
	v201 = v194
	v202 = v195
	v203 = v196
	goto L53
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v205
	v207 = int32(4)
	v208 = v201 + v207
	v210 = v203 + v207
	v212 = v202 - v207
	if base.Ui32(int32(3)) < base.Ui32(v212) {
		v201 = v208
		v202 = v212
		v203 = v210
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v216 = v208
	v217 = v212
	v218 = v210
	goto L20
L55:
	;
	goto L54
L56:
	;
	v223 = v216
	v224 = v217
	v225 = v218
	goto L57
L57:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v227)
	v229 = int32(1)
	v234 = v224 - v229
	if v234 != 0 {
		v223 = v223 + v229
		v224 = v234
		v225 = v225 + v229
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L15
L59:
	;
	goto L58
L60:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	v398 = int32(1)
	v400 = v397 + v26<<(uint(v398)%32)
	v403 = v397 + v92<<(uint(v398)%32)
	v405 = v99 << (uint(v398) % 32)
	if v400 == v403 {
		goto L107
	} else {
		goto L108
	}
L61:
	;
	goto L60
L62:
	;
	v254 = v249 + v101
	if base.Ui32(v250-v254) <= base.Ui32(int32(0)-v101<<(uint(int32(1))%32)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v261 = F___memcpy(m, v249, v250, v101)
	mBase = m.M
	goto L60
L64:
	;
	goto L65
L65:
	;
	v264 = (v249 ^ v250) & int32(3)
	if base.Ui32(v249) < base.Ui32(v250) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v366 == int32(0) {
		goto L61
	} else {
		goto L102
	}
L67:
	;
	if base.Ui32(v344) <= base.Ui32(int32(3)) {
		v365 = v343
		v366 = v344
		v367 = v345
		goto L66
	} else {
		goto L98
	}
L68:
	;
	if v264 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if v264 != 0 {
		v326 = v101
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v365 = v250
	v366 = v101
	v367 = v249
	goto L66
L72:
	;
	goto L73
L73:
	;
	if v249&int32(3) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v343 = v250
	v344 = v101
	v345 = v249
	goto L67
L75:
	;
	goto L76
L76:
	;
	v271 = v250
	v272 = v101
	v273 = v249
	goto L77
L77:
	;
	if v272 == int32(0) {
		goto L61
	} else {
		goto L79
	}
L78:
	;
	v343 = v280
	v344 = v282
	v345 = v284
	goto L67
L79:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v277)
	v279 = int32(1)
	v280 = v271 + v279
	v282 = v272 - v279
	v284 = v273 + v279
	if v284&int32(3) != 0 {
		v271 = v280
		v272 = v282
		v273 = v284
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	if v326 == int32(0) {
		goto L61
	} else {
		goto L94
	}
L82:
	;
	if v254&int32(3) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v291 = v101
	goto L86
L84:
	;
	v306 = v101
	goto L85
L85:
	;
	if base.Ui32(v306) <= base.Ui32(int32(3)) {
		v326 = v306
		goto L81
	} else {
		goto L90
	}
L86:
	;
	if v291 == int32(0) {
		goto L61
	} else {
		goto L88
	}
L87:
	;
	v306 = v297
	goto L85
L88:
	;
	v297 = v291 - int32(1)
	v298 = v249 + v297
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v297))))
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v300)
	if v298&int32(3) != 0 {
		v291 = v297
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v313 = v306
	goto L91
L91:
	;
	v317 = v313 - int32(4)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v250+v317)))
	*(*int32)(unsafe.Add(mBase, uint32(v249+v317))) = v320
	if base.Ui32(int32(3)) < base.Ui32(v317) {
		v313 = v317
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v326 = v317
	goto L81
L93:
	;
	goto L92
L94:
	;
	v333 = v326
	goto L95
L95:
	;
	v337 = v333 - int32(1)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v337))))
	*(*uint8)(unsafe.Add(mBase, uint32(v249+v337))) = uint8(v340)
	if v337 != 0 {
		v333 = v337
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L61
L97:
	;
	goto L96
L98:
	;
	v350 = v343
	v351 = v344
	v352 = v345
	goto L99
L99:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v354
	v356 = int32(4)
	v357 = v350 + v356
	v359 = v352 + v356
	v361 = v351 - v356
	if base.Ui32(int32(3)) < base.Ui32(v361) {
		v350 = v357
		v351 = v361
		v352 = v359
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v365 = v357
	v366 = v361
	v367 = v359
	goto L66
L101:
	;
	goto L100
L102:
	;
	v372 = v365
	v373 = v366
	v374 = v367
	goto L103
L103:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	*(*uint8)(unsafe.Add(mBase, uint32(v374))) = uint8(v376)
	v378 = int32(1)
	v383 = v373 - v378
	if v383 != 0 {
		v372 = v372 + v378
		v373 = v383
		v374 = v374 + v378
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L61
L105:
	;
	goto L104
L106:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	v553 = v552 + v26
	v554 = v92 + v552
	if v553 == v554 {
		goto L153
	} else {
		goto L154
	}
L107:
	;
	goto L106
L108:
	;
	v409 = v400 + v405
	if base.Ui32(v403-v409) <= base.Ui32(int32(0)-v405<<(uint(int32(1))%32)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v416 = F___memcpy(m, v400, v403, v405)
	mBase = m.M
	goto L106
L110:
	;
	goto L111
L111:
	;
	v419 = (v400 ^ v403) & int32(3)
	if base.Ui32(v400) < base.Ui32(v403) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v521 == int32(0) {
		goto L107
	} else {
		goto L148
	}
L113:
	;
	if base.Ui32(v499) <= base.Ui32(int32(3)) {
		v520 = v498
		v521 = v499
		v522 = v500
		goto L112
	} else {
		goto L144
	}
L114:
	;
	if v419 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	if v419 != 0 {
		v481 = v405
		goto L127
	} else {
		goto L128
	}
L117:
	;
	v520 = v403
	v521 = v405
	v522 = v400
	goto L112
L118:
	;
	goto L119
L119:
	;
	if v400&int32(3) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v498 = v403
	v499 = v405
	v500 = v400
	goto L113
L121:
	;
	goto L122
L122:
	;
	v426 = v403
	v427 = v405
	v428 = v400
	goto L123
L123:
	;
	if v427 == int32(0) {
		goto L107
	} else {
		goto L125
	}
L124:
	;
	v498 = v435
	v499 = v437
	v500 = v439
	goto L113
L125:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v432)
	v434 = int32(1)
	v435 = v426 + v434
	v437 = v427 - v434
	v439 = v428 + v434
	if v439&int32(3) != 0 {
		v426 = v435
		v427 = v437
		v428 = v439
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	if v481 == int32(0) {
		goto L107
	} else {
		goto L140
	}
L128:
	;
	if v409&int32(3) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v446 = v405
	goto L132
L130:
	;
	v461 = v405
	goto L131
L131:
	;
	if base.Ui32(v461) <= base.Ui32(int32(3)) {
		v481 = v461
		goto L127
	} else {
		goto L136
	}
L132:
	;
	if v446 == int32(0) {
		goto L107
	} else {
		goto L134
	}
L133:
	;
	v461 = v452
	goto L131
L134:
	;
	v452 = v446 - int32(1)
	v453 = v400 + v452
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v452))))
	*(*uint8)(unsafe.Add(mBase, uint32(v453))) = uint8(v455)
	if v453&int32(3) != 0 {
		v446 = v452
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v468 = v461
	goto L137
L137:
	;
	v472 = v468 - int32(4)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v403+v472)))
	*(*int32)(unsafe.Add(mBase, uint32(v400+v472))) = v475
	if base.Ui32(int32(3)) < base.Ui32(v472) {
		v468 = v472
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v481 = v472
	goto L127
L139:
	;
	goto L138
L140:
	;
	v488 = v481
	goto L141
L141:
	;
	v492 = v488 - int32(1)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v492))))
	*(*uint8)(unsafe.Add(mBase, uint32(v400+v492))) = uint8(v495)
	if v492 != 0 {
		v488 = v492
		goto L141
	} else {
		goto L143
	}
L142:
	;
	goto L107
L143:
	;
	goto L142
L144:
	;
	v505 = v498
	v506 = v499
	v507 = v500
	goto L145
L145:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v509
	v511 = int32(4)
	v512 = v505 + v511
	v514 = v507 + v511
	v516 = v506 - v511
	if base.Ui32(int32(3)) < base.Ui32(v516) {
		v505 = v512
		v506 = v516
		v507 = v514
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v520 = v512
	v521 = v516
	v522 = v514
	goto L112
L147:
	;
	goto L146
L148:
	;
	v527 = v520
	v528 = v521
	v529 = v522
	goto L149
L149:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527))))
	*(*uint8)(unsafe.Add(mBase, uint32(v529))) = uint8(v531)
	v533 = int32(1)
	v538 = v528 - v533
	if v538 != 0 {
		v527 = v527 + v533
		v528 = v538
		v529 = v529 + v533
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L107
L151:
	;
	goto L150
L152:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v699<<(uint(int32(2))%32)+v87-int32(4)))) = int32(-1)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v709 = v707 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v709
	if v26 < v709 {
		goto L198
	} else {
		goto L199
	}
L153:
	;
	goto L152
L154:
	;
	v558 = v553 + v99
	if base.Ui32(v554-v558) <= base.Ui32(int32(0)-v99<<(uint(int32(1))%32)) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v565 = F___memcpy(m, v553, v554, v99)
	mBase = m.M
	goto L152
L156:
	;
	goto L157
L157:
	;
	v568 = (v553 ^ v554) & int32(3)
	if base.Ui32(v553) < base.Ui32(v554) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	if v670 == int32(0) {
		goto L153
	} else {
		goto L194
	}
L159:
	;
	if base.Ui32(v648) <= base.Ui32(int32(3)) {
		v669 = v647
		v670 = v648
		v671 = v649
		goto L158
	} else {
		goto L190
	}
L160:
	;
	if v568 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	if v568 != 0 {
		v630 = v99
		goto L173
	} else {
		goto L174
	}
L163:
	;
	v669 = v554
	v670 = v99
	v671 = v553
	goto L158
L164:
	;
	goto L165
L165:
	;
	if v553&int32(3) == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v647 = v554
	v648 = v99
	v649 = v553
	goto L159
L167:
	;
	goto L168
L168:
	;
	v575 = v554
	v576 = v99
	v577 = v553
	goto L169
L169:
	;
	if v576 == int32(0) {
		goto L153
	} else {
		goto L171
	}
L170:
	;
	v647 = v584
	v648 = v586
	v649 = v588
	goto L159
L171:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
	*(*uint8)(unsafe.Add(mBase, uint32(v577))) = uint8(v581)
	v583 = int32(1)
	v584 = v575 + v583
	v586 = v576 - v583
	v588 = v577 + v583
	if v588&int32(3) != 0 {
		v575 = v584
		v576 = v586
		v577 = v588
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	if v630 == int32(0) {
		goto L153
	} else {
		goto L186
	}
L174:
	;
	if v558&int32(3) != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v595 = v99
	goto L178
L176:
	;
	v610 = v99
	goto L177
L177:
	;
	if base.Ui32(v610) <= base.Ui32(int32(3)) {
		v630 = v610
		goto L173
	} else {
		goto L182
	}
L178:
	;
	if v595 == int32(0) {
		goto L153
	} else {
		goto L180
	}
L179:
	;
	v610 = v601
	goto L177
L180:
	;
	v601 = v595 - int32(1)
	v602 = v553 + v601
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554+v601))))
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v604)
	if v602&int32(3) != 0 {
		v595 = v601
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v617 = v610
	goto L183
L183:
	;
	v621 = v617 - int32(4)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v554+v621)))
	*(*int32)(unsafe.Add(mBase, uint32(v553+v621))) = v624
	if base.Ui32(int32(3)) < base.Ui32(v621) {
		v617 = v621
		goto L183
	} else {
		goto L185
	}
L184:
	;
	v630 = v621
	goto L173
L185:
	;
	goto L184
L186:
	;
	v637 = v630
	goto L187
L187:
	;
	v641 = v637 - int32(1)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554+v641))))
	*(*uint8)(unsafe.Add(mBase, uint32(v553+v641))) = uint8(v644)
	if v641 != 0 {
		v637 = v641
		goto L187
	} else {
		goto L189
	}
L188:
	;
	goto L153
L189:
	;
	goto L188
L190:
	;
	v654 = v647
	v655 = v648
	v656 = v649
	goto L191
L191:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	*(*int32)(unsafe.Add(mBase, uint32(v656))) = v658
	v660 = int32(4)
	v661 = v654 + v660
	v663 = v656 + v660
	v665 = v655 - v660
	if base.Ui32(int32(3)) < base.Ui32(v665) {
		v654 = v661
		v655 = v665
		v656 = v663
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v669 = v661
	v670 = v665
	v671 = v663
	goto L158
L193:
	;
	goto L192
L194:
	;
	v676 = v669
	v677 = v670
	v678 = v671
	goto L195
L195:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	*(*uint8)(unsafe.Add(mBase, uint32(v678))) = uint8(v680)
	v682 = int32(1)
	v687 = v677 - v682
	if v687 != 0 {
		v676 = v676 + v682
		v677 = v687
		v678 = v678 + v682
		goto L195
	} else {
		goto L197
	}
L196:
	;
	goto L153
L197:
	;
	goto L196
L198:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v714 = v26
	goto L201
L199:
	;
	goto L200
L200:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v745+int32(384))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L204
	}
L201:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v87+v714<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v713+v726*int32(640))+48)) = v714
	v732 = v714 + int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v732 < v733 {
		v714 = v732
		goto L201
	} else {
		goto L203
	}
L202:
	;
	goto L200
L203:
	;
	goto L202
L204:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v751+int32(512))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	return
}
func F_ProcNumberGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v2 = int32(0)
	if l0 < v2 {
		v20 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[101]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		if base.Ui32(v9) <= base.Ui32(l0) {
			v20 = int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v14 = v11 + l0*int32(640)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
			if v16 != 0 {
				v17 = v14
			} else {
				v17 = int32(0)
			}
			v20 = v17
		}
	}
	return v20
}
func F_proc_exit(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[451]))
	if v8 == int32(42) {
		F_proc_exit_prepare(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v15 = F_errstart(m, int32(12), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
					F_errmsg_internal(m, int32(680490), v5)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						F_errfinish(m, int32(501305), int32(155), int32(99843))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pgl_exit(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					F_pgl_exit(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(130477), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_errfinish(m, int32(501305), int32(109), int32(99843))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
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
