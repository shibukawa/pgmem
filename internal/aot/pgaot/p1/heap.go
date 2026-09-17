package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	v4 = int32(1)
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v5&int32(2176) != 0 {
		v147 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v147
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == int32(0) {
		v147 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v5&int32(_a_F_HeapTupleHeaderIsOnlyLocked_0) == v11 {
		v147 = v11
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_HeapTupleGetUpdateXid(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if base.Ui32(v16) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v139 != 0 {
		v147 = v11
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v139 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[0]))
	if v30 == v16 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v139 = int32(1)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[1]))
	if v34 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v139 = v131
	goto L7
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[2]))
	if v38 == int32(0) {
		v131 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[3]))
	v102 = int32(0)
	v104 = v34 - int32(1)
	goto L37
L18:
	;
	v43 = v38
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v48 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v131 = int32(0)
	goto L14
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v95 != 0 {
		v43 = v95
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v51 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v54 = int32(1)
	if v16 == v51 {
		v131 = v54
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v58 = v56 - int32(1)
	if v58 < int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v63 = int32(0)
	v65 = v58
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v71 = int32(2)
	v72 = base.I32_div_s(v65-v63, v71)
	v73 = v72 + v63
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69+v73<<(uint(v71)%32))))
	if v77 == v16 {
		v131 = v54
		goto L14
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v81 = F_TransactionIdPrecedes(m, v77, v16)
	mBase = m.M
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = v73 + int32(1)
	goto L31
L30:
	;
	v82 = v63
	goto L31
L31:
	;
	if v81 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = v65
	goto L34
L33:
	;
	v85 = v73 - int32(1)
	goto L34
L34:
	;
	if v82 <= v85 {
		v63 = v82
		v65 = v85
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L20
L37:
	;
	v109 = int32(2)
	v110 = base.I32_div_s(v104-v102, v109)
	v111 = v110 + v102
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100+v111<<(uint(v109)%32))))
	v116 = base.B2i32(v115 == v16)
	if v115 == v16 {
		v131 = v116
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v131 = v116
	goto L14
L39:
	;
	v119 = base.B2i32(base.Ui32(v115) < base.Ui32(v16))
	if base.Ui32(v115) < base.Ui32(v16) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = v111 + int32(1)
	goto L42
L41:
	;
	v120 = v102
	goto L42
L42:
	;
	if base.Ui32(v115) < base.Ui32(v16) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v123 = v104
	goto L45
L44:
	;
	v123 = v111 - int32(1)
	goto L45
L45:
	;
	if v120 <= v123 {
		v102 = v120
		v104 = v123
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v140 = F_TransactionIdIsInProgress(m, v16)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if v140 != 0 {
		v147 = v11
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v142 = F_TransactionIdDidCommit(m, v16)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v147 = v142 ^ int32(1)
	goto L1
}
func F_HeapTupleSatisfiesVacuumHorizon(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
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
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
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
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v9&int32(256) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v679 = v676 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v679)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L53
	} else {
		goto L255
	}
L2:
	;
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L53
	} else {
		goto L254
	}
L3:
	;
	v666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v668 = v666 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v668)
	goto L2
L4:
	;
	return v664
L5:
	;
	v582 = int32(1)
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v583&int32(2048) != 0 {
		v664 = v582
		goto L4
	} else {
		goto L214
	}
L6:
	;
	v12 = base.I32_extend16_s(v9)
	if v12&int32(512) != 0 {
		v664 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v12&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_0) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v575 = v145 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v575)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L53
	} else {
		goto L213
	}
L9:
	;
	v17 = int32(4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if base.Ui32(v18) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	if v12 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	if v138 != 0 {
		v664 = v17
		goto L4
	} else {
		goto L52
	}
L13:
	;
	v138 = int32(0)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v29 == v18 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v138 = int32(1)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v33 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v138 = v130
	goto L12
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v37 == int32(0) {
		v130 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v101 = int32(0)
	v103 = v33 - int32(1)
	goto L42
L23:
	;
	v42 = v37
	goto L24
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v47 == int32(4) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v130 = int32(0)
	goto L19
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+80))
	if v94 != 0 {
		v42 = v94
		goto L24
	} else {
		goto L41
	}
L27:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v50 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v53 = int32(1)
	if v18 == v50 {
		v130 = v53
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	v57 = v55 - int32(1)
	if v57 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v62 = int32(0)
	v64 = v57
	goto L31
L31:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v70 = int32(2)
	v71 = base.I32_div_s(v64-v62, v70)
	v72 = v71 + v62
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68+v72<<(uint(v70)%32))))
	if v76 == v18 {
		v130 = v53
		goto L19
	} else {
		goto L33
	}
L32:
	;
	goto L26
L33:
	;
	v80 = F_TransactionIdPrecedes(m, v76, v18)
	mBase = m.M
	if v80 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v81 = v72 + int32(1)
	goto L36
L35:
	;
	v81 = v62
	goto L36
L36:
	;
	if v80 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v84 = v64
	goto L39
L38:
	;
	v84 = v72 - int32(1)
	goto L39
L39:
	;
	if v81 <= v84 {
		v62 = v81
		v64 = v84
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
	v108 = int32(2)
	v109 = base.I32_div_s(v103-v101, v108)
	v110 = v109 + v101
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99+v110<<(uint(v108)%32))))
	v115 = base.B2i32(v114 == v18)
	if v114 == v18 {
		v130 = v115
		goto L19
	} else {
		goto L44
	}
L43:
	;
	v130 = v115
	goto L19
L44:
	;
	v118 = base.B2i32(base.Ui32(v114) < base.Ui32(v18))
	if base.Ui32(v114) < base.Ui32(v18) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v119 = v110 + int32(1)
	goto L47
L46:
	;
	v119 = v101
	goto L47
L47:
	;
	if base.Ui32(v114) < base.Ui32(v18) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v122 = v103
	goto L50
L49:
	;
	v122 = v110 - int32(1)
	goto L50
L50:
	;
	if v119 <= v122 {
		v101 = v119
		v103 = v122
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v139 = F_TransactionIdIsInProgress(m, v18)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	if v139 != 0 {
		v664 = v17
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v143 = F_TransactionIdDidCommit(m, v18)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v143 == int32(0) {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v676 = v145
	goto L1
L58:
	;
	v150 = int32(3)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if base.Ui32(v151) < base.Ui32(v150) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui32(v285) < base.Ui32(int32(3)) {
		goto L110
	} else {
		goto L111
	}
L61:
	;
	if v271 != 0 {
		v664 = v150
		goto L4
	} else {
		goto L101
	}
L62:
	;
	v271 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v162 == v151 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v271 = int32(1)
	goto L61
L66:
	;
	goto L67
L67:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v166 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v271 = v263
	goto L61
L69:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v170 == int32(0) {
		v263 = int32(0)
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v234 = int32(0)
	v236 = v166 - int32(1)
	goto L91
L72:
	;
	v175 = v170
	goto L73
L73:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v180 == int32(4) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v263 = int32(0)
	goto L68
L75:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v175)+80))
	if v227 != 0 {
		v175 = v227
		goto L73
	} else {
		goto L90
	}
L76:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v183 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v186 = int32(1)
	if v151 == v183 {
		v263 = v186
		goto L68
	} else {
		goto L78
	}
L78:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+52))
	v190 = v188 - int32(1)
	if v190 < int32(0) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v195 = int32(0)
	v197 = v190
	goto L80
L80:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v175)+48))
	v203 = int32(2)
	v204 = base.I32_div_s(v197-v195, v203)
	v205 = v204 + v195
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v201+v205<<(uint(v203)%32))))
	if v209 == v151 {
		v263 = v186
		goto L68
	} else {
		goto L82
	}
L81:
	;
	goto L75
L82:
	;
	v213 = F_TransactionIdPrecedes(m, v209, v151)
	mBase = m.M
	if v213 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v214 = v205 + int32(1)
	goto L85
L84:
	;
	v214 = v195
	goto L85
L85:
	;
	if v213 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v217 = v197
	goto L88
L87:
	;
	v217 = v205 - int32(1)
	goto L88
L88:
	;
	if v214 <= v217 {
		v195 = v214
		v197 = v217
		goto L80
	} else {
		goto L89
	}
L89:
	;
	goto L81
L90:
	;
	goto L74
L91:
	;
	v241 = int32(2)
	v242 = base.I32_div_s(v236-v234, v241)
	v243 = v242 + v234
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v232+v243<<(uint(v241)%32))))
	v248 = base.B2i32(v247 == v151)
	if v247 == v151 {
		v263 = v248
		goto L68
	} else {
		goto L93
	}
L92:
	;
	v263 = v248
	goto L68
L93:
	;
	v251 = base.B2i32(base.Ui32(v247) < base.Ui32(v151))
	if base.Ui32(v247) < base.Ui32(v151) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v252 = v243 + int32(1)
	goto L96
L95:
	;
	v252 = v234
	goto L96
L96:
	;
	if base.Ui32(v247) < base.Ui32(v151) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v255 = v236
	goto L99
L98:
	;
	v255 = v243 - int32(1)
	goto L99
L99:
	;
	if v252 <= v255 {
		v234 = v252
		v236 = v255
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L92
L101:
	;
	v272 = F_TransactionIdIsInProgress(m, v151)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L53
	} else {
		goto L102
	}
L102:
	;
	if v272 != 0 {
		v664 = v150
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v274 = F_TransactionIdDidCommit(m, v151)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L53
	} else {
		goto L104
	}
L104:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v274 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v676 = v276
	goto L1
L106:
	;
	goto L107
L107:
	;
	v280 = v276 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v280)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L53
	} else {
		goto L108
	}
L108:
	;
	goto L5
L109:
	;
	if v405 != 0 {
		goto L149
	} else {
		goto L150
	}
L110:
	;
	v405 = int32(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v296 == v285 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v405 = int32(1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v300 <= int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v405 = v397
	goto L109
L117:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v304 == int32(0) {
		v397 = int32(0)
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v368 = int32(0)
	v370 = v300 - int32(1)
	goto L139
L120:
	;
	v309 = v304
	goto L121
L121:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	if v314 == int32(4) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v397 = int32(0)
	goto L116
L123:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v309)+80))
	if v361 != 0 {
		v309 = v361
		goto L121
	} else {
		goto L138
	}
L124:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if v317 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v320 = int32(1)
	if v285 == v317 {
		v397 = v320
		goto L116
	} else {
		goto L126
	}
L126:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v309)+52))
	v324 = v322 - int32(1)
	if v324 < int32(0) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v329 = int32(0)
	v331 = v324
	goto L128
L128:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v309)+48))
	v337 = int32(2)
	v338 = base.I32_div_s(v331-v329, v337)
	v339 = v338 + v329
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v335+v339<<(uint(v337)%32))))
	if v343 == v285 {
		v397 = v320
		goto L116
	} else {
		goto L130
	}
L129:
	;
	goto L123
L130:
	;
	v347 = F_TransactionIdPrecedes(m, v343, v285)
	mBase = m.M
	if v347 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v348 = v339 + int32(1)
	goto L133
L132:
	;
	v348 = v329
	goto L133
L133:
	;
	if v347 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v351 = v331
	goto L136
L135:
	;
	v351 = v339 - int32(1)
	goto L136
L136:
	;
	if v348 <= v351 {
		v329 = v348
		v331 = v351
		goto L128
	} else {
		goto L137
	}
L137:
	;
	goto L129
L138:
	;
	goto L122
L139:
	;
	v375 = int32(2)
	v376 = base.I32_div_s(v370-v368, v375)
	v377 = v376 + v368
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v366+v377<<(uint(v375)%32))))
	v382 = base.B2i32(v381 == v285)
	if v381 == v285 {
		v397 = v382
		goto L116
	} else {
		goto L141
	}
L140:
	;
	v397 = v382
	goto L116
L141:
	;
	v385 = base.B2i32(base.Ui32(v381) < base.Ui32(v285))
	if base.Ui32(v381) < base.Ui32(v285) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v386 = v377 + int32(1)
	goto L144
L143:
	;
	v386 = v368
	goto L144
L144:
	;
	if base.Ui32(v381) < base.Ui32(v285) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v389 = v370
	goto L147
L146:
	;
	v389 = v377 - int32(1)
	goto L147
L147:
	;
	if v386 <= v389 {
		v368 = v386
		v370 = v389
		goto L139
	} else {
		goto L148
	}
L148:
	;
	goto L140
L149:
	;
	v406 = int32(3)
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v407&int32(2048)|v407&int32(128)|base.B2i32(v407&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_1) == int32(64)) != 0 {
		v664 = v406
		goto L4
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v554 = F_TransactionIdIsInProgress(m, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L53
	} else {
		goto L203
	}
L152:
	;
	v418 = F_HeapTupleHeaderIsOnlyLocked(m, v6)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L53
	} else {
		goto L153
	}
L153:
	;
	if v418 != 0 {
		v664 = v406
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v422&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_2) == int32(_a_F_HeapTupleSatisfiesVacuumHorizon_3) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if base.Ui32(v430) < base.Ui32(int32(3)) {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	v427 = F_HeapTupleGetUpdateXid(m, v6)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L53
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v430 = v429
	goto L155
L159:
	;
	v430 = v427
	goto L155
L160:
	;
	if v550 != 0 {
		goto L200
	} else {
		goto L201
	}
L161:
	;
	v550 = int32(0)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v441 == v430 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v550 = int32(1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v445 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v550 = v542
	goto L160
L168:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v449 == int32(0) {
		v542 = int32(0)
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v513 = int32(0)
	v515 = v445 - int32(1)
	goto L190
L171:
	;
	v454 = v449
	goto L172
L172:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v454)+20))
	if v459 == int32(4) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v542 = int32(0)
	goto L167
L174:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v454)+80))
	if v506 != 0 {
		v454 = v506
		goto L172
	} else {
		goto L189
	}
L175:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v462 == int32(0) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v465 = int32(1)
	if v430 == v462 {
		v542 = v465
		goto L167
	} else {
		goto L177
	}
L177:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v454)+52))
	v469 = v467 - int32(1)
	if v469 < int32(0) {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v474 = int32(0)
	v476 = v469
	goto L179
L179:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v454)+48))
	v482 = int32(2)
	v483 = base.I32_div_s(v476-v474, v482)
	v484 = v483 + v474
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v480+v484<<(uint(v482)%32))))
	if v488 == v430 {
		v542 = v465
		goto L167
	} else {
		goto L181
	}
L180:
	;
	goto L174
L181:
	;
	v492 = F_TransactionIdPrecedes(m, v488, v430)
	mBase = m.M
	if v492 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v493 = v484 + int32(1)
	goto L184
L183:
	;
	v493 = v474
	goto L184
L184:
	;
	if v492 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v496 = v476
	goto L187
L186:
	;
	v496 = v484 - int32(1)
	goto L187
L187:
	;
	if v493 <= v496 {
		v474 = v493
		v476 = v496
		goto L179
	} else {
		goto L188
	}
L188:
	;
	goto L180
L189:
	;
	goto L173
L190:
	;
	v520 = int32(2)
	v521 = base.I32_div_s(v515-v513, v520)
	v522 = v521 + v513
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v511+v522<<(uint(v520)%32))))
	v527 = base.B2i32(v526 == v430)
	if v526 == v430 {
		v542 = v527
		goto L167
	} else {
		goto L192
	}
L191:
	;
	v542 = v527
	goto L167
L192:
	;
	v530 = base.B2i32(base.Ui32(v526) < base.Ui32(v430))
	if base.Ui32(v526) < base.Ui32(v430) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v531 = v522 + int32(1)
	goto L195
L194:
	;
	v531 = v513
	goto L195
L195:
	;
	if base.Ui32(v526) < base.Ui32(v430) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v534 = v515
	goto L198
L197:
	;
	v534 = v522 - int32(1)
	goto L198
L198:
	;
	if v531 <= v534 {
		v513 = v531
		v515 = v534
		goto L190
	} else {
		goto L199
	}
L199:
	;
	goto L191
L200:
	;
	v551 = int32(4)
	goto L202
L201:
	;
	v551 = int32(3)
	goto L202
L202:
	;
	return v551
L203:
	;
	if v554 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	return int32(3)
L205:
	;
	goto L206
L206:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v559 = F_TransactionIdDidCommit(m, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L53
	} else {
		goto L207
	}
L207:
	;
	if v559 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_HeapTupleSetHintBits(m, v6, l1, int32(256), v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L53
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v567 = v565 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v567)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L53
	} else {
		goto L212
	}
L211:
	;
	goto L5
L212:
	;
	return int32(0)
L213:
	;
	goto L5
L214:
	;
	v588 = int32(0)
	if base.B2i32(v583&int32(128) == v588)&base.B2i32(v583&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_1) != int32(64)) == v588 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	if v583&int32(1024) != 0 {
		v664 = v582
		goto L4
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	if v583&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_3) != 0 {
		goto L230
	} else {
		goto L231
	}
L218:
	;
	if v583&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_3) != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v583&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_4) != int32(_a_F_HeapTupleSatisfiesVacuumHorizon_5) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	goto L221
L221:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v615 = F_TransactionIdIsInProgress(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L53
	} else {
		goto L227
	}
L222:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v607 = F_MultiXactIdIsRunning(m, v605, int32(1))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L53
	} else {
		goto L225
	}
L223:
	;
	v610 = v583
	goto L224
L224:
	;
	v612 = v610 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v612)
	goto L2
L225:
	;
	if v607 != 0 {
		v664 = v582
		goto L4
	} else {
		goto L226
	}
L226:
	;
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v610 = v609
	goto L224
L227:
	;
	if v615 != 0 {
		v664 = v582
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v619 = v617 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v619)
	goto L2
L229:
	;
	v664 = int32(2)
	goto L4
L230:
	;
	v623 = F_HeapTupleGetUpdateXid(m, v6)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L53
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	if v583&int32(1024) == int32(0) {
		goto L244
	} else {
		goto L245
	}
L233:
	;
	v625 = F_TransactionIdIsInProgress(m, v623)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L53
	} else {
		goto L234
	}
L234:
	;
	if v625 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	return int32(4)
L236:
	;
	goto L237
L237:
	;
	v629 = F_TransactionIdDidCommit(m, v623)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L53
	} else {
		goto L238
	}
L238:
	;
	if v629 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v623
	goto L229
L240:
	;
	goto L241
L241:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v634 = F_MultiXactIdIsRunning(m, v632, int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L53
	} else {
		goto L242
	}
L242:
	;
	if v634 != 0 {
		v664 = v582
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v638 = v636 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v638)
	goto L2
L244:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v645 = F_TransactionIdIsInProgress(m, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L53
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v658
	goto L229
L247:
	;
	if v645 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	return int32(4)
L249:
	;
	goto L250
L250:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v650 = F_TransactionIdDidCommit(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L53
	} else {
		goto L251
	}
L251:
	;
	if v650 == int32(0) {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	F_HeapTupleSetHintBits(m, v6, l1, int32(1024), v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L53
	} else {
		goto L253
	}
L253:
	;
	goto L246
L254:
	;
	return int32(1)
L255:
	;
	return int32(0)
}
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
	if v10&int32(2047) < l1 {
		if l2 == int32(0) {
			v64 = int32(1)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+l1<<(uint(int32(4))%32))+12)))
			if v19 == int32(0) {
				v64 = int32(1)
			} else {
				v64 = int32(0)
			}
		}
		m.G0 = v7 + int32(16)
		return v64
	} else {
		if int32(0) < l1 {
			v25 = int32(0)
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
			if v26&int32(1) == v25 {
				v64 = v25
			} else {
				v31 = int32(1)
				v32 = l1 - v31
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(base.Ui32(v32)>>(uint(int32(3))%32)))+23)))
				v64 = base.B2i32(int32(base.Ui32(v36)>>(uint(v32&int32(7))%32))&v31 == int32(0))
			}
			m.G0 = v7 + int32(16)
			return v64
		} else {
			if base.Ui32(int32(-7)) < base.Ui32(l1) {
				v64 = int32(0)
				m.G0 = v7 + int32(16)
				return v64
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(_a_F_heap_attisnull_0), v7)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_heap_attisnull_1), int32(491), int32(_a_F_heap_attisnull_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
func F_heap_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v483 int32
	_ = v483
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int64
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int64
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	v8 = l7
	v9 = l8
	v10 = l9
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	if l11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L82
	} else {
		goto L178
	}
L2:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v51
	if v8 == int32(105) {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	goto L4
L4:
	;
	if l1 == int32(11) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = base.B2i32(v8 != int32(105))
	goto L7
L6:
	;
	v33 = int32(0)
	goto L7
L7:
	;
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 != int32(99) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v48 == int32(2) {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v40 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L12:
	;
	v38 = F_isTempToastNamespace(m, l1)
	mBase = m.M
	v40 = v38
	goto L14
L13:
	;
	v40 = int32(1)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v44 != int32(2) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	goto L2
L18:
	;
	v58 = l2
	goto L20
L19:
	;
	v58 = v51
	goto L20
L20:
	;
	if v8 == int32(114) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = l2
	goto L23
L22:
	;
	v61 = v58
	goto L23
L23:
	;
	if v8 == int32(83) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = l2
	goto L26
L25:
	;
	v64 = v61
	goto L26
L26:
	;
	if v8 == int32(116) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v67 = l2
	goto L29
L28:
	;
	v67 = v64
	goto L29
L29:
	;
	if v8 == int32(109) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v70 = l2
	goto L32
L31:
	;
	v70 = v67
	goto L32
L32:
	;
	if v8 == int32(73) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v73 = l2
	goto L35
L34:
	;
	v73 = v70
	goto L35
L35:
	;
	if v8 == int32(112) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v76 = l2
	goto L38
L37:
	;
	v76 = v73
	goto L38
L38:
	;
	if v8 != int32(83) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v80 = v76
	goto L41
L40:
	;
	v80 = int32(0)
	goto L41
L41:
	;
	switch v8 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L43
	default:
		v85 = l4
		v86 = int32(0)
		goto L42
	}
L42:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[1]))
	if v80 != v89 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if l4 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v84 = l4
	goto L46
L45:
	;
	v84 = l3
	goto L46
L46:
	;
	v85 = v84
	v86 = l14
	goto L42
L47:
	;
	if v86 != 0 {
		goto L165
	} else {
		goto L166
	}
L48:
	;
	v91 = v80
	goto L50
L49:
	;
	v91 = int32(0)
	goto L50
L50:
	;
	v92 = m.G0
	v94 = v92 - int32(48)
	m.G0 = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v99 = int32(1)
	if l3 <= int32(3591) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v170 == v10 {
		goto L76
	} else {
		goto L77
	}
L52:
	;
	goto L51
L53:
	;
	v170 = int32(0)
	goto L52
L54:
	;
	if base.B2i32(base.Ui32(l3-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l3-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v170 = v99
		goto L52
	} else {
		goto L75
	}
L55:
	;
	if l3 <= int32(2670) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if l3 <= int32(_a_F_heap_create_0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	switch l3 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v170 = v99
		goto L52
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L53
	default:
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v111 = l3 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v111))|base.B2i32(int32(1)<<(uint(v111)%32)&int32(226492515) == int32(0)) != 0 {
		goto L54
	} else {
		goto L63
	}
L61:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l3-int32(2396)) {
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v170 = v99
	goto L52
L63:
	;
	v170 = v99
	goto L52
L64:
	;
	if base.Ui32(l3-int32(3592)) < base.Ui32(int32(2)) {
		v170 = v99
		goto L52
	} else {
		goto L73
	}
L65:
	;
	v124 = l3 - int32(_a_F_heap_create_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v124))|base.B2i32(int32(1)<<(uint(v124)%32)&int32(963) == int32(0)) != 0 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	switch l3 - int32(_a_F_heap_create_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v170 = v99
		goto L52
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L53
	default:
		goto L69
	}
L68:
	;
	v170 = v99
	goto L52
L69:
	;
	if base.Ui32(l3-int32(_a_F_heap_create_3)) < base.Ui32(int32(3)) {
		v170 = v99
		goto L52
	} else {
		goto L70
	}
L70:
	;
	v141 = l3 - int32(_a_F_heap_create_4)
	if base.Ui32(int32(15)) < base.Ui32(v141) {
		goto L53
	} else {
		goto L71
	}
L71:
	;
	if int32(1)<<(uint(v141)%32)&int32(_a_F_heap_create_5) != 0 {
		v170 = v99
		goto L52
	} else {
		goto L72
	}
L72:
	;
	goto L53
L73:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l3-int32(4060)) {
		goto L53
	} else {
		goto L74
	}
L74:
	;
	v170 = v99
	goto L52
L75:
	;
	goto L53
L76:
	;
	v173 = l3 - int32(1247)
	v178 = base.B2i32(base.Ui32(v173) < base.Ui32(int32(16))) & int32(base.Ui32(int32(_a_F_heap_create_6))>>(uint(v173)%32))
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[2]))
	if v180 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L82
	} else {
		goto L161
	}
L79:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v189 = v180
	goto L81
L81:
	;
	v190 = int32(_a_F_heap_create_7)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[3])) = v189
	v195 = F_palloc0(m, int32(276))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L82
	} else {
		goto L84
	}
L82:
	;
	return int32(0)
L83:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[2]))
	v189 = v188
	goto L81
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+25)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v178
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[4]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+32)) = v203
	v209 = F_CreateTupleDescCopy(m, l6)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+52)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = int32(1)
	if v96 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v316 = F_palloc0(m, int32(144))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L82
	} else {
		goto L98
	}
L88:
	;
	v231 = int32(0)
	v239 = int32(0)
	goto L89
L89:
	;
	v241 = v231 * int32(100)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = int32(4)
	v247 = v241 + (v242 + v243<<(uint(v244)%32))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v252 = l6 + v248<<(uint(v244)%32) + v241
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+109)))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+109)) = uint8(v253)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+110)))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+110)) = uint8(v255)
	v258 = v252 + int32(106)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+106)) = uint8(v259)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	F_populate_compact_attribute(m, v261, v231)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L82
	} else {
		goto L91
	}
L90:
	;
	if v264&int32(255) == int32(0) {
		goto L87
	} else {
		goto L96
	}
L91:
	;
	v264 = v259 | v239
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v265 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v269 = v231 << (uint(int32(4)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v269)+31)))
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v270)+31)) = uint8(v273)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v276 = int32(1)
	v279 = v231 + v276
	if v279 != v96 {
		v231 = v279
		v239 = v264 & v276
		goto L89
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v286 = F_palloc0(m, int32(20))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L82
	} else {
		goto L97
	}
L97:
	;
	v288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+16)) = uint8(v288)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+16)) = v286
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+48)) = v316
	v322 = F_strncpy(m, v316+int32(4), l0, int32(64))
	mBase = m.M
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+63)) = uint8(v323)
	goto L99
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+68)) = l1
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v327)+119)) = uint8(v8)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+120)) = uint16(v96)
	v331 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+72)) = v331
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+80)) = int32(10)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+118)) = uint8(v9)
	switch v9 - int32(112) {
	case 0, 5:
		v366 = int32(-1)
		v367 = v331
		goto L100
	default:
		goto L102
	case 4:
		goto L101
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+24)) = uint8(v367)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = v366
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v370)+129)) = uint8(base.B2i32(v8 != int32(109)))
	v374 = int32(110)
	goto L110
L101:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[5]))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[6]))
	if v361 == int32(-1) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L82
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v9
	F_errmsg_internal(m, int32(_a_F_heap_create_8), v94+int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L82
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3665), int32(_a_F_heap_create_10))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L82
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v364 = v359
	goto L108
L107:
	;
	v364 = v361
	goto L108
L108:
	;
	v366 = v364
	v367 = int32(1)
	goto L100
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+130)) = uint8(v392)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+117)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+56)) = l3
	if v96 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if l1 == int32(11) {
		v392 = v374
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v378 = v8 - int32(109)
	if base.Ui32(int32(5)) < base.Ui32(v378&int32(255)) {
		v392 = v374
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v392 = base.I32_wrap_i64(int64(base.Ui64(int64(110425294138980)) >> (uint(base.I64_extend_i32_u(v378<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L109
L113:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v552)+92)) = v91
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	if l10 != 0 {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v401 = v96 & int32(3)
	v402 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v96) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v415 = int32(0)
	v423 = v402
	goto L118
L116:
	;
	v483 = v402
	goto L117
L117:
	;
	v506 = v483
	v508 = v402
	goto L122
L118:
	;
	v433 = v423 * int32(100)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v436 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v433+(v434+v435<<(uint(v436)%32)))+20)) = l3
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	*(*int32)(unsafe.Add(mBase, uint32(v441+v442<<(uint(v436)%32)+v433)+120)) = l3
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*int32)(unsafe.Add(mBase, uint32(v448+v449<<(uint(v436)%32)+v433)+220)) = l3
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v433+(v455+v456<<(uint(v436)%32)))+320)) = l3
	v463 = v423 + v436
	v465 = v415 + v436
	if v465 != v96&int32(2147483644) {
		v415 = v465
		v423 = v463
		goto L118
	} else {
		goto L120
	}
L119:
	;
	if v401 == int32(0) {
		goto L113
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v483 = v463
	goto L117
L122:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v515+v516<<(uint(int32(4))%32)+v506*int32(100))+20)) = l3
	v524 = int32(1)
	v527 = v508 + v524
	if v527 != v401 {
		v506 = v506 + v524
		v508 = v527
		goto L122
	} else {
		goto L124
	}
L123:
	;
	goto L113
L124:
	;
	goto L123
L125:
	;
	v562 = v195 + int32(56)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v195)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+60)) = v563
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[7]))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+117)))
	if v569 != 0 {
		goto L131
	} else {
		goto L132
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v554)+88)) = int32(0)
	F_RelationMapUpdateMap(m, l3, v85, v10, int32(1))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L82
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v554)+88)) = v85
	goto L125
L129:
	;
	goto L125
L130:
	;
	F_RelationInitPhysicalAddr(m, v195)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L82
	} else {
		goto L134
	}
L131:
	;
	v570 = int32(0)
	goto L133
L132:
	;
	v570 = v567
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+64)) = v570
	goto L130
L134:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+84)) = l5
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[3])) = v191
	switch v8 - int32(83) {
	case 0, 26, 31, 33:
		goto L136
	default:
		goto L135
	}
L135:
	;
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[8]))
	v587 = F_hash_search(m, v583, v562, int32(1), v94+int32(47))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L82
	} else {
		goto L138
	}
L136:
	;
	F_RelationInitTableAccessMethod(m, v195)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L82
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+47)))
	if v589 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[9]))
	if v627 <= int32(31) {
		goto L153
	} else {
		goto L154
	}
L140:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v587)+4)) = v195
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v592)+16))
	if v594 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587)+4)) = v195
	goto L139
L143:
	;
	F_RelationDestroyRelation(m, v592, int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L82
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v601 == int32(0) {
		goto L139
	} else {
		goto L147
	}
L146:
	;
	goto L139
L147:
	;
	v606 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L82
	} else {
		goto L148
	}
L148:
	;
	if v606 == int32(0) {
		goto L139
	} else {
		goto L149
	}
L149:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v592)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v610 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_create_11), v94+int32(32))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L82
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3738), int32(_a_F_heap_create_10))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L82
	} else {
		goto L151
	}
L151:
	;
	goto L139
L152:
	;
	v644 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+26)) = uint8(v644)
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[10]))
	F_ResourceOwnerEnlarge(m, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L82
	} else {
		goto L156
	}
L153:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[9])) = v627 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v627<<(uint(int32(2))%32))+uint32(_c_F_heap_create[11]))) = v630
	goto L152
L154:
	;
	goto L155
L155:
	;
	v641 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_create[12])) = uint8(v641)
	goto L152
L156:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v650 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v655 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[10]))
	F_ResourceOwnerRemember(m, v657, v195, int32(_a_F_heap_create_12))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L82
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	m.G0 = v94 + int32(48)
	goto L47
L160:
	;
	goto L159
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = l0
	F_errmsg_internal(m, int32(_a_F_heap_create_13), v94)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L82
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3566), int32(_a_F_heap_create_10))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L82
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[7]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724)+117)))
	if v725 != 0 {
		goto L174
	} else {
		goto L175
	}
L165:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+119)))
	switch v679 - int32(83) {
	case 0, 22:
		goto L168
	default:
		goto L164
	case 26, 31, 33:
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	if v91 == int32(0) {
		goto L164
	} else {
		goto L172
	}
L168:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v686
	v688 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v688
	v691 = F_RelationCreateStorage(m, v26, v9, int32(1))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L82
	} else {
		goto L171
	}
L169:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v195)+188))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+112))
	m.T0[v683].(func(*base.Module, int32, int32, int32, int32, int32))(m, v195, v195, v9, l12, l13)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L82
	} else {
		goto L170
	}
L170:
	;
	goto L164
L171:
	;
	goto L164
L172:
	;
	v695 = m.G0
	v697 = v695 - int32(32)
	m.G0 = v697
	v699 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v697)+28)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v697)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v697)+20)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v697)+16)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v697)+12)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v697)+8)) = int32(1213)
	F_recordSharedDependencyOn(m, v697+int32(20), v697+int32(8), int32(116))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L82
	} else {
		goto L173
	}
L173:
	;
	m.G0 = v697 + int32(32)
	goto L164
L174:
	;
	v726 = int32(0)
	goto L176
L175:
	;
	v726 = v723
	goto L176
L176:
	;
	v727 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v195)+56)))
	F_pgstat_create_transactional(m, int32(2), v726, v727)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L82
	} else {
		goto L177
	}
L177:
	;
	m.G0 = v26 + int32(32)
	return v195
L178:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L82
	} else {
		goto L179
	}
L179:
	;
	v741 = F_get_namespace_name(m, l1)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L82
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v741
	F_errmsg(m, int32(_a_F_heap_create_14), v26+int32(16))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L82
	} else {
		goto L181
	}
L181:
	;
	F_errdetail(m, int32(_a_F_heap_create_15), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L82
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_heap_create_16), int32(323), int32(_a_F_heap_create_17))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L82
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v316 int32
	_ = v316
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int64
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int64
	_ = v432
	var v435 int32
	_ = v435
	var v445 int64
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v502 int32
	_ = v502
	var v506 int64
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
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
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int64
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int64
	_ = v550
	var v553 int32
	_ = v553
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v22, v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v27 <= int32(0) {
			m.G0 = v16 + int32(16)
			return
		} else {
			switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(7) - int32(1) {
			case 0:
				v339 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v340 = F_SnapBuildProcessChange(m, v21, v23, v339)
				mBase = m.M
				v341 = m.ExcPending
				if v341 != 0 {
					return
				} else {
					if v340 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v344 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v345 = m.G0
							v347 = v345 - int32(16)
							m.G0 = v347
							v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
							v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+64))
							v352 = int32(0)
							F_XLogRecGetBlockTag(m, v349, v352, v347+int32(4), v352, v352)
							mBase = m.M
							v358 = m.ExcPending
							if v358 != 0 {
								return
							} else {
								v359 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
								v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+88))
								if v359 != v361 {
									m.G0 = v347 + int32(16)
									m.G0 = v16 + int32(16)
									return
								} else {
									v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v363 != 0 {
										v364 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
										v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364)+56)))
										v366 = F_filter_by_origin_cb_wrapper(m, l0, v365)
										mBase = m.M
										v367 = m.ExcPending
										if v367 != 0 {
											return
										} else {
											if v366 != 0 {
												m.G0 = v347 + int32(16)
												m.G0 = v16 + int32(16)
												return
											} else {
												v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v369 = F_ReorderBufferAllocChange(m, v368)
												mBase = m.M
												v370 = m.ExcPending
												if v370 != 0 {
													return
												} else {
													v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+7)))
													if v373&int32(8) != 0 {
														v376 = int32(10)
													} else {
														v376 = int32(2)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v376
													v378 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
													v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v369)+16)) = uint16(v379)
													v381 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v369)+28)) = v381
													v383 = *(*int64)(unsafe.Add(mBase, uint32(v347)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v369)+20)) = v383
													v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+7)))
													if v385&int32(6) != 0 {
														v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v389 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
														v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+68))
														v392 = v390 - int32(13)
														v393 = F_ReorderBufferAllocTupleBuf(m, v388, v392)
														mBase = m.M
														v394 = m.ExcPending
														if v394 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v369)+36)) = v393
															v396 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v393)+12)) = v396
															*(*uint16)(unsafe.Add(mBase, uint32(v393)+8)) = uint16(v396)
															*(*int32)(unsafe.Add(mBase, uint32(v393)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v393))) = v390 + int32(10)
															v405 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
															v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+12)))
															v407 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
															v408 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = v408
															*(*int64)(unsafe.Add(mBase, uint32(v407)+15)) = v408
															*(*int64)(unsafe.Add(mBase, uint32(v407))) = v408
															if v392 != 0 {
																v414 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
																base.MemoryCopy(m, v414+int32(23), v351+int32(13), v392)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v407)+22)) = uint8(v406)
															*(*int32)(unsafe.Add(mBase, uint32(v407)+18)) = v405
															v427 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v369)+32)) = uint8(v427)
															v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v430 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
															v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+36))
															v432 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v429, v431, v432, v369, int32(0))
															mBase = m.M
															v435 = m.ExcPending
															if v435 != 0 {
																return
															} else {
																m.G0 = v347 + int32(16)
																m.G0 = v16 + int32(16)
																return
															}
														}
													} else {
														v427 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v369)+32)) = uint8(v427)
														v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v430 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
														v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+36))
														v432 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v429, v431, v432, v369, int32(0))
														mBase = m.M
														v435 = m.ExcPending
														if v435 != 0 {
															return
														} else {
															m.G0 = v347 + int32(16)
															m.G0 = v16 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v369 = F_ReorderBufferAllocChange(m, v368)
										mBase = m.M
										v370 = m.ExcPending
										if v370 != 0 {
											return
										} else {
											v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+7)))
											if v373&int32(8) != 0 {
												v376 = int32(10)
											} else {
												v376 = int32(2)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v376
											v378 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
											v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v369)+16)) = uint16(v379)
											v381 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v369)+28)) = v381
											v383 = *(*int64)(unsafe.Add(mBase, uint32(v347)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v369)+20)) = v383
											v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+7)))
											if v385&int32(6) != 0 {
												v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v389 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
												v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+68))
												v392 = v390 - int32(13)
												v393 = F_ReorderBufferAllocTupleBuf(m, v388, v392)
												mBase = m.M
												v394 = m.ExcPending
												if v394 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v369)+36)) = v393
													v396 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v393)+12)) = v396
													*(*uint16)(unsafe.Add(mBase, uint32(v393)+8)) = uint16(v396)
													*(*int32)(unsafe.Add(mBase, uint32(v393)+4)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v393))) = v390 + int32(10)
													v405 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
													v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+12)))
													v407 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
													v408 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = v408
													*(*int64)(unsafe.Add(mBase, uint32(v407)+15)) = v408
													*(*int64)(unsafe.Add(mBase, uint32(v407))) = v408
													if v392 != 0 {
														v414 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
														base.MemoryCopy(m, v414+int32(23), v351+int32(13), v392)
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v407)+22)) = uint8(v406)
													*(*int32)(unsafe.Add(mBase, uint32(v407)+18)) = v405
													v427 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v369)+32)) = uint8(v427)
													v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v430 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
													v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+36))
													v432 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v429, v431, v432, v369, int32(0))
													mBase = m.M
													v435 = m.ExcPending
													if v435 != 0 {
														return
													} else {
														m.G0 = v347 + int32(16)
														m.G0 = v16 + int32(16)
														return
													}
												}
											} else {
												v427 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v369)+32)) = uint8(v427)
												v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v430 = *(*int32)(unsafe.Add(mBase, uint32(v349)+96))
												v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+36))
												v432 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												F_ReorderBufferQueueChange(m, v429, v431, v432, v369, int32(0))
												mBase = m.M
												v435 = m.ExcPending
												if v435 != 0 {
													return
												} else {
													m.G0 = v347 + int32(16)
													m.G0 = v16 + int32(16)
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
			case 1, 3:
				v174 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v175 = F_SnapBuildProcessChange(m, v21, v23, v174)
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return
				} else {
					if v175 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v179 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+64))
							v183 = int32(0)
							F_XLogRecGetBlockTag(m, v180, v183, v16+int32(4), v183, v183)
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return
							} else {
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
								v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+88))
								if v190 != v192 {
									m.G0 = v16 + int32(16)
									return
								} else {
									v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v194 != 0 {
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
										v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+56)))
										v197 = F_filter_by_origin_cb_wrapper(m, l0, v196)
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return
										} else {
											if v197 != 0 {
												m.G0 = v16 + int32(16)
												return
											} else {
												v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v200 = F_ReorderBufferAllocChange(m, v199)
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = int32(1)
													v204 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
													v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v200)+16)) = uint16(v205)
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v200)+28)) = v207
													v209 = *(*int64)(unsafe.Add(mBase, uint32(v16)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v200)+20)) = v209
													v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
													if v211&int32(16) != 0 {
														v214 = int32(0)
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
														if v217 < v214 {
															v239 = v214
															v242 = v239
														} else {
															v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+int32(0))+76)))
															if v222 != int32(1) {
																v239 = v214
																v242 = v239
															} else {
																v226 = v216 + int32(76)
																v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+43)))
																if v227 == int32(0) {
																	if v16 == int32(0) {
																		v239 = v214
																		v242 = v239
																	} else {
																		v232 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v232
																		v242 = v232
																	}
																} else {
																	if v16 != 0 {
																		v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+48)))
																		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v235
																	} else {
																	}
																	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+44))
																	v239 = v237
																	v242 = v239
																}
															}
														}
														v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v244 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
														v247 = F_ReorderBufferAllocTupleBuf(m, v243, v244-int32(5))
														mBase = m.M
														v248 = m.ExcPending
														if v248 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v200)+40)) = v247
															v250 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
															v251 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v247)+12)) = v251
															*(*uint16)(unsafe.Add(mBase, uint32(v247)+8)) = uint16(v251)
															*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v247))) = v250 + int32(18)
															v260 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
															v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)))
															v262 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
															v263 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v262)+8)) = v263
															*(*int64)(unsafe.Add(mBase, uint32(v262)+15)) = v263
															*(*int64)(unsafe.Add(mBase, uint32(v262))) = v263
															v270 = v250 - int32(5)
															if v270 != 0 {
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
																base.MemoryCopy(m, v271+int32(23), v242+int32(5), v270)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v262)+22)) = uint8(v261)
															*(*int32)(unsafe.Add(mBase, uint32(v262)+18)) = v260
															v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
															v286 = v279
															if v286&int32(12) != 0 {
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
																v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+64))
																v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
																v294 = v292 - int32(19)
																v295 = F_ReorderBufferAllocTupleBuf(m, v291, v294)
																mBase = m.M
																v296 = m.ExcPending
																if v296 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v200)+36)) = v295
																	v298 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v298
																	*(*uint16)(unsafe.Add(mBase, uint32(v295)+8)) = uint16(v298)
																	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(-1)
																	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v292 + int32(4)
																	v307 = *(*int32)(unsafe.Add(mBase, uint32(v290)+14))
																	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+18)))
																	v309 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
																	v310 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v310
																	*(*int64)(unsafe.Add(mBase, uint32(v309)+15)) = v310
																	*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310
																	if v294 != 0 {
																		v316 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
																		base.MemoryCopy(m, v316+int32(23), v290+int32(19), v294)
																	} else {
																	}
																	*(*uint8)(unsafe.Add(mBase, uint32(v309)+22)) = uint8(v308)
																	*(*int32)(unsafe.Add(mBase, uint32(v309)+18)) = v307
																	v330 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
																	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
																	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
																	v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																	F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
																	mBase = m.M
																	v338 = m.ExcPending
																	if v338 != 0 {
																		return
																	} else {
																		m.G0 = v16 + int32(16)
																		return
																	}
																}
															} else {
																v330 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
																v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
																v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
																v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
																mBase = m.M
																v338 = m.ExcPending
																if v338 != 0 {
																	return
																} else {
																	m.G0 = v16 + int32(16)
																	return
																}
															}
														}
													} else {
														v286 = v211
														if v286&int32(12) != 0 {
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
															v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+64))
															v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
															v294 = v292 - int32(19)
															v295 = F_ReorderBufferAllocTupleBuf(m, v291, v294)
															mBase = m.M
															v296 = m.ExcPending
															if v296 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v200)+36)) = v295
																v298 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v298
																*(*uint16)(unsafe.Add(mBase, uint32(v295)+8)) = uint16(v298)
																*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v295))) = v292 + int32(4)
																v307 = *(*int32)(unsafe.Add(mBase, uint32(v290)+14))
																v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+18)))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
																v310 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v310
																*(*int64)(unsafe.Add(mBase, uint32(v309)+15)) = v310
																*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310
																if v294 != 0 {
																	v316 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
																	base.MemoryCopy(m, v316+int32(23), v290+int32(19), v294)
																} else {
																}
																*(*uint8)(unsafe.Add(mBase, uint32(v309)+22)) = uint8(v308)
																*(*int32)(unsafe.Add(mBase, uint32(v309)+18)) = v307
																v330 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
																v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
																v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
																v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
																mBase = m.M
																v338 = m.ExcPending
																if v338 != 0 {
																	return
																} else {
																	m.G0 = v16 + int32(16)
																	return
																}
															}
														} else {
															v330 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
															v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
															v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
															v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
															mBase = m.M
															v338 = m.ExcPending
															if v338 != 0 {
																return
															} else {
																m.G0 = v16 + int32(16)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v200 = F_ReorderBufferAllocChange(m, v199)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = int32(1)
											v204 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
											v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v200)+16)) = uint16(v205)
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v200)+28)) = v207
											v209 = *(*int64)(unsafe.Add(mBase, uint32(v16)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v200)+20)) = v209
											v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
											if v211&int32(16) != 0 {
												v214 = int32(0)
												v216 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
												v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
												if v217 < v214 {
													v239 = v214
													v242 = v239
												} else {
													v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+int32(0))+76)))
													if v222 != int32(1) {
														v239 = v214
														v242 = v239
													} else {
														v226 = v216 + int32(76)
														v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+43)))
														if v227 == int32(0) {
															if v16 == int32(0) {
																v239 = v214
																v242 = v239
															} else {
																v232 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v16))) = v232
																v242 = v232
															}
														} else {
															if v16 != 0 {
																v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+48)))
																*(*int32)(unsafe.Add(mBase, uint32(v16))) = v235
															} else {
															}
															v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+44))
															v239 = v237
															v242 = v239
														}
													}
												}
												v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v244 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
												v247 = F_ReorderBufferAllocTupleBuf(m, v243, v244-int32(5))
												mBase = m.M
												v248 = m.ExcPending
												if v248 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+40)) = v247
													v250 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
													v251 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v247)+12)) = v251
													*(*uint16)(unsafe.Add(mBase, uint32(v247)+8)) = uint16(v251)
													*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v247))) = v250 + int32(18)
													v260 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
													v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)))
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
													v263 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v262)+8)) = v263
													*(*int64)(unsafe.Add(mBase, uint32(v262)+15)) = v263
													*(*int64)(unsafe.Add(mBase, uint32(v262))) = v263
													v270 = v250 - int32(5)
													if v270 != 0 {
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
														base.MemoryCopy(m, v271+int32(23), v242+int32(5), v270)
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v262)+22)) = uint8(v261)
													*(*int32)(unsafe.Add(mBase, uint32(v262)+18)) = v260
													v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
													v286 = v279
													if v286&int32(12) != 0 {
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
														v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+64))
														v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
														v294 = v292 - int32(19)
														v295 = F_ReorderBufferAllocTupleBuf(m, v291, v294)
														mBase = m.M
														v296 = m.ExcPending
														if v296 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v200)+36)) = v295
															v298 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v298
															*(*uint16)(unsafe.Add(mBase, uint32(v295)+8)) = uint16(v298)
															*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v295))) = v292 + int32(4)
															v307 = *(*int32)(unsafe.Add(mBase, uint32(v290)+14))
															v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+18)))
															v309 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
															v310 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v310
															*(*int64)(unsafe.Add(mBase, uint32(v309)+15)) = v310
															*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310
															if v294 != 0 {
																v316 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
																base.MemoryCopy(m, v316+int32(23), v290+int32(19), v294)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v309)+22)) = uint8(v308)
															*(*int32)(unsafe.Add(mBase, uint32(v309)+18)) = v307
															v330 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
															v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
															v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
															v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
															mBase = m.M
															v338 = m.ExcPending
															if v338 != 0 {
																return
															} else {
																m.G0 = v16 + int32(16)
																return
															}
														}
													} else {
														v330 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
														v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
														v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
														mBase = m.M
														v338 = m.ExcPending
														if v338 != 0 {
															return
														} else {
															m.G0 = v16 + int32(16)
															return
														}
													}
												}
											} else {
												v286 = v211
												if v286&int32(12) != 0 {
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
													v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+64))
													v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
													v294 = v292 - int32(19)
													v295 = F_ReorderBufferAllocTupleBuf(m, v291, v294)
													mBase = m.M
													v296 = m.ExcPending
													if v296 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v200)+36)) = v295
														v298 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v298
														*(*uint16)(unsafe.Add(mBase, uint32(v295)+8)) = uint16(v298)
														*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v295))) = v292 + int32(4)
														v307 = *(*int32)(unsafe.Add(mBase, uint32(v290)+14))
														v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+18)))
														v309 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
														v310 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v310
														*(*int64)(unsafe.Add(mBase, uint32(v309)+15)) = v310
														*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310
														if v294 != 0 {
															v316 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
															base.MemoryCopy(m, v316+int32(23), v290+int32(19), v294)
														} else {
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v309)+22)) = uint8(v308)
														*(*int32)(unsafe.Add(mBase, uint32(v309)+18)) = v307
														v330 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
														v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
														v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
														mBase = m.M
														v338 = m.ExcPending
														if v338 != 0 {
															return
														} else {
															m.G0 = v16 + int32(16)
															return
														}
													}
												} else {
													v330 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v330)
													v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v333 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
													v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
													v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v332, v334, v335, v200, int32(0))
													mBase = m.M
													v338 = m.ExcPending
													if v338 != 0 {
														return
													} else {
														m.G0 = v16 + int32(16)
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
			case 2:
				v445 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v446 = F_SnapBuildProcessChange(m, v21, v23, v445)
				mBase = m.M
				v447 = m.ExcPending
				if v447 != 0 {
					return
				} else {
					if v446 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v450 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v451 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
							v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+64))
							v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
							v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+88))
							if v454 != v456 {
								m.G0 = v16 + int32(16)
								return
							} else {
								v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								if v458 != 0 {
									v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+56)))
									v460 = F_filter_by_origin_cb_wrapper(m, l0, v459)
									mBase = m.M
									v461 = m.ExcPending
									if v461 != 0 {
										return
									} else {
										if v460 != 0 {
											m.G0 = v16 + int32(16)
											return
										} else {
											v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v463 = F_ReorderBufferAllocChange(m, v462)
											mBase = m.M
											v464 = m.ExcPending
											if v464 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v463)+8)) = int32(11)
												v467 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
												v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+56)))
												*(*uint16)(unsafe.Add(mBase, uint32(v463)+16)) = uint16(v468)
												v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+8)))
												if v470&int32(1) != 0 {
													v473 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v463)+24)) = uint8(v473)
													v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+8)))
													v476 = v475
												} else {
													v476 = v470
												}
												if v476&int32(2) != 0 {
													v479 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v463)+25)) = uint8(v479)
												} else {
												}
												v481 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v463)+20)) = v481
												v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+120))
												v487 = F_MemoryContextAlloc(m, v484, v481<<(uint(int32(2))%32))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v463)+28)) = v487
													v490 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
													v492 = v490 << (uint(int32(2)) % 32)
													if v492 != 0 {
														base.MemoryCopy(m, v487, v453+int32(12), v492)
													} else {
													}
													v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v497 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
													v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+36))
													v499 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v496, v498, v499, v463, int32(0))
													mBase = m.M
													v502 = m.ExcPending
													if v502 != 0 {
														return
													} else {
														m.G0 = v16 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v463 = F_ReorderBufferAllocChange(m, v462)
									mBase = m.M
									v464 = m.ExcPending
									if v464 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v463)+8)) = int32(11)
										v467 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
										v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(v463)+16)) = uint16(v468)
										v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+8)))
										if v470&int32(1) != 0 {
											v473 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v463)+24)) = uint8(v473)
											v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+8)))
											v476 = v475
										} else {
											v476 = v470
										}
										if v476&int32(2) != 0 {
											v479 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v463)+25)) = uint8(v479)
										} else {
										}
										v481 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v463)+20)) = v481
										v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+120))
										v487 = F_MemoryContextAlloc(m, v484, v481<<(uint(int32(2))%32))
										mBase = m.M
										v488 = m.ExcPending
										if v488 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v463)+28)) = v487
											v490 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
											v492 = v490 << (uint(int32(2)) % 32)
											if v492 != 0 {
												base.MemoryCopy(m, v487, v453+int32(12), v492)
											} else {
											}
											v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v497 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
											v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+36))
											v499 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											F_ReorderBufferQueueChange(m, v496, v498, v499, v463, int32(0))
											mBase = m.M
											v502 = m.ExcPending
											if v502 != 0 {
												return
											} else {
												m.G0 = v16 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			case 4:
				v506 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v507 = F_SnapBuildProcessChange(m, v21, v23, v506)
				mBase = m.M
				v508 = m.ExcPending
				if v508 != 0 {
					return
				} else {
					if v507 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v511 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v512 = m.G0
							v514 = v512 - int32(16)
							m.G0 = v514
							v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v517 = int32(0)
							F_XLogRecGetBlockTag(m, v516, v517, v514+int32(4), v517, v517)
							mBase = m.M
							v523 = m.ExcPending
							if v523 != 0 {
								return
							} else {
								v524 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
								v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+88))
								if v524 != v526 {
									m.G0 = v514 + int32(16)
									m.G0 = v16 + int32(16)
									return
								} else {
									v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v528 != 0 {
										v529 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
										v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v529)+56)))
										v531 = F_filter_by_origin_cb_wrapper(m, l0, v530)
										mBase = m.M
										v532 = m.ExcPending
										if v532 != 0 {
											return
										} else {
											if v531 != 0 {
												m.G0 = v514 + int32(16)
												m.G0 = v16 + int32(16)
												return
											} else {
												v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v534 = F_ReorderBufferAllocChange(m, v533)
												mBase = m.M
												v535 = m.ExcPending
												if v535 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = int32(9)
													v538 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
													v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v538)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v534)+16)) = uint16(v539)
													v541 = *(*int64)(unsafe.Add(mBase, uint32(v514)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v534)+20)) = v541
													v543 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v534)+28)) = v543
													v545 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v534)+32)) = uint8(v545)
													v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v548 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
													v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+36))
													v550 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v547, v549, v550, v534, int32(0))
													mBase = m.M
													v553 = m.ExcPending
													if v553 != 0 {
														return
													} else {
														m.G0 = v514 + int32(16)
														m.G0 = v16 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v534 = F_ReorderBufferAllocChange(m, v533)
										mBase = m.M
										v535 = m.ExcPending
										if v535 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = int32(9)
											v538 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
											v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v538)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v534)+16)) = uint16(v539)
											v541 = *(*int64)(unsafe.Add(mBase, uint32(v514)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v534)+20)) = v541
											v543 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v534)+28)) = v543
											v545 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v534)+32)) = uint8(v545)
											v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v548 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
											v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+36))
											v550 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											F_ReorderBufferQueueChange(m, v547, v549, v550, v534, int32(0))
											mBase = m.M
											v553 = m.ExcPending
											if v553 != 0 {
												return
											} else {
												m.G0 = v514 + int32(16)
												m.G0 = v16 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			case 5, 6:
				m.G0 = v16 + int32(16)
				return
			default:
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v37 = F_SnapBuildProcessChange(m, v21, v23, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v37 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v41 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v42 = m.G0
							v44 = v42 - int32(16)
							m.G0 = v44
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
							if v49&int32(8) == int32(0) {
								m.G0 = v44 + int32(16)
								m.G0 = v16 + int32(16)
								return
							} else {
								v54 = int32(0)
								F_XLogRecGetBlockTag(m, v46, v54, v44, v54, v54)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
									if v59 != v61 {
										m.G0 = v44 + int32(16)
										m.G0 = v16 + int32(16)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										if v63 != 0 {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
											v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+56)))
											v66 = F_filter_by_origin_cb_wrapper(m, l0, v65)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												if v66 != 0 {
													m.G0 = v44 + int32(16)
													m.G0 = v16 + int32(16)
													return
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v69 = F_ReorderBufferAllocChange(m, v68)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
														*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v71 << (uint(int32(1)) % 32) & int32(8)
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
														v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+56)))
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)) = uint16(v78)
														v80 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
														*(*int64)(unsafe.Add(mBase, uint32(v69)+20)) = v80
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v82
														v84 = int32(0)
														v86 = v44 + int32(12)
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+72))
														if v89 < v84 {
															v111 = v84
															v114 = v111
														} else {
															v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(0))+76)))
															if v94 != int32(1) {
																v111 = v84
																v114 = v111
															} else {
																v98 = v88 + int32(76)
																v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+43)))
																if v99 == int32(0) {
																	if v86 == int32(0) {
																		v111 = v84
																		v114 = v111
																	} else {
																		v104 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v86))) = v104
																		v114 = v104
																	}
																} else {
																	if v86 != 0 {
																		v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+48)))
																		*(*int32)(unsafe.Add(mBase, uint32(v86))) = v107
																	} else {
																	}
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)+44))
																	v111 = v109
																	v114 = v111
																}
															}
														}
														v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
														v119 = F_ReorderBufferAllocTupleBuf(m, v115, v116-int32(5))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v69)+40)) = v119
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
															v123 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v123
															*(*uint16)(unsafe.Add(mBase, uint32(v119)+8)) = uint16(v123)
															*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v119))) = v122 + int32(18)
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
															v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
															v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															v135 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v135
															*(*int64)(unsafe.Add(mBase, uint32(v134)+15)) = v135
															*(*int64)(unsafe.Add(mBase, uint32(v134))) = v135
															v142 = v122 - int32(5)
															if v142 != 0 {
																v143 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
																base.MemoryCopy(m, v143+int32(23), v114+int32(5), v142)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v134)+22)) = uint8(v133)
															*(*int32)(unsafe.Add(mBase, uint32(v134)+18)) = v132
															v151 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v69)+32)) = uint8(v151)
															v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v154 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+36))
															v156 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
															F_ReorderBufferQueueChange(m, v153, v155, v156, v69, int32(base.Ui32(v157&int32(16))>>(uint(int32(4))%32)))
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return
															} else {
																m.G0 = v44 + int32(16)
																m.G0 = v16 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v69 = F_ReorderBufferAllocChange(m, v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
												*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v71 << (uint(int32(1)) % 32) & int32(8)
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
												v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+56)))
												*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)) = uint16(v78)
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
												*(*int64)(unsafe.Add(mBase, uint32(v69)+20)) = v80
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v82
												v84 = int32(0)
												v86 = v44 + int32(12)
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+72))
												if v89 < v84 {
													v111 = v84
													v114 = v111
												} else {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(0))+76)))
													if v94 != int32(1) {
														v111 = v84
														v114 = v111
													} else {
														v98 = v88 + int32(76)
														v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+43)))
														if v99 == int32(0) {
															if v86 == int32(0) {
																v111 = v84
																v114 = v111
															} else {
																v104 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v86))) = v104
																v114 = v104
															}
														} else {
															if v86 != 0 {
																v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+48)))
																*(*int32)(unsafe.Add(mBase, uint32(v86))) = v107
															} else {
															}
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)+44))
															v111 = v109
															v114 = v111
														}
													}
												}
												v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
												v119 = F_ReorderBufferAllocTupleBuf(m, v115, v116-int32(5))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v69)+40)) = v119
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
													v123 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v123
													*(*uint16)(unsafe.Add(mBase, uint32(v119)+8)) = uint16(v123)
													*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v119))) = v122 + int32(18)
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
													v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
													v135 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v135
													*(*int64)(unsafe.Add(mBase, uint32(v134)+15)) = v135
													*(*int64)(unsafe.Add(mBase, uint32(v134))) = v135
													v142 = v122 - int32(5)
													if v142 != 0 {
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														base.MemoryCopy(m, v143+int32(23), v114+int32(5), v142)
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v134)+22)) = uint8(v133)
													*(*int32)(unsafe.Add(mBase, uint32(v134)+18)) = v132
													v151 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v69)+32)) = uint8(v151)
													v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v154 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+36))
													v156 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
													F_ReorderBufferQueueChange(m, v153, v155, v156, v69, int32(base.Ui32(v157&int32(16))>>(uint(int32(4))%32)))
													mBase = m.M
													v163 = m.ExcPending
													if v163 != 0 {
														return
													} else {
														m.G0 = v44 + int32(16)
														m.G0 = v16 + int32(16)
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
func F_heap_fill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v7 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(128)
	goto L3
L2:
	;
	v18 = v7
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = l5 - int32(1)
	goto L6
L5:
	;
	v23 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v27 = v25 & int32(_a_F_heap_fill_tuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v27)
	if int32(0) < v15 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v11 + int32(16)
	return
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v48 = v11 + int32(8)
	goto L14
L13:
	;
	v48 = int32(0)
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1+v36<<(uint(int32(2))%32))))
	v58 = v56
	goto L17
L16:
	;
	v58 = int32(0)
	goto L17
L17:
	;
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v36))))
	v62 = v60
	goto L20
L19:
	;
	v62 = int32(1)
	goto L20
L20:
	;
	F_fill_val(m, l0+int32(20)+v36<<(uint(int32(4))%32), v48, v11+int32(4), v11+int32(12), l4, v58, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v66 = v36 + int32(1)
	if v66 != v15 {
		v36 = v66
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
}
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
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
			v87 = v17
			m.G0 = v10 + int32(16)
			return v87
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
				v80 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					v87 = v80
					m.G0 = v10 + int32(16)
					return v87
				}
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v39 = v23 + v37 + v34
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+6)))
				if v40 != int32(1) {
					v87 = v39
					m.G0 = v10 + int32(16)
					return v87
				} else {
					v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
					switch v43&int32(_a_F_heap_getattr_6_0) - int32(1) {
					case 0:
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39))))
						v87 = v48
						m.G0 = v10 + int32(16)
						return v87
					case 1:
						v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39))))
						v87 = v49
						m.G0 = v10 + int32(16)
						return v87
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
							F_errmsg_internal(m, int32(_a_F_heap_getattr_6_1), v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_6_2), int32(70), int32(_a_F_heap_getattr_6_3))
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
						v87 = v50
						m.G0 = v10 + int32(16)
						return v87
					}
				}
			}
		} else {
			v64 = int32(1)
			v65 = l1 - v64
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(base.Ui32(v65)>>(uint(int32(3))%32)))+23)))
			if int32(base.Ui32(v69)>>(uint(v65&int32(7))%32))&v64 != 0 {
				v80 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					v87 = v80
					m.G0 = v10 + int32(16)
					return v87
				}
			} else {
				v75 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v75)
				v87 = int32(0)
				m.G0 = v10 + int32(16)
				return v87
			}
		}
	}
}
func F_heap_index_delete_tuples(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v202 int32
	_ = v202
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v487 int32
	_ = v487
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v542 int32
	_ = v542
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v566 int64
	_ = v566
	var v567 int64
	_ = v567
	var v568 int64
	_ = v568
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1238 int32
	_ = v1238
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1521 int32
	_ = v1521
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1565 int32
	_ = v1565
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(192)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+188)) = v3
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = int32(6)
	v37 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_heap_index_delete_tuples[0])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+152)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v56 = v3
	goto L2
L2:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56<<(uint(int32(2))%32))+uint32(_c_F_heap_index_delete_tuples[1])))
	if v73 < v42 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v236 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	v83 = v73
	goto L7
L5:
	;
	goto L6
L6:
	;
	v232 = v56 + int32(1)
	if v232 != int32(9) {
		v56 = v232
		goto L2
	} else {
		goto L25
	}
L7:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v43+v83<<(uint(int32(3))%32))))
	if v83 < v73 {
		v173 = v83
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43+v173<<(uint(int32(3))%32)))) = v105
	v202 = v83 + int32(1)
	if v202 != v42 {
		v83 = v202
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v109 = base.I32_rotl(base.I32_wrap_i64(v105), int32(16))
	v114 = base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(int64(32))%64))) & int32(_a_F_heap_index_delete_tuples_0)
	v120 = v83
	goto L11
L11:
	;
	v142 = v120 - v73
	v145 = v43 + v142<<(uint(int32(3))%32)
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145))))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+2)))
	v150 = v146<<(uint(int32(16))%32) | v149
	if v109 != v150 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v173 = v142
	goto L9
L13:
	;
	if v161 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if base.Ui32(v150) < base.Ui32(v109) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
	v161 = base.B2i32(base.Ui32(v114) < base.Ui32(v156)) - base.B2i32(base.Ui32(v156) < base.Ui32(v114))
	goto L13
L17:
	;
	v155 = int32(-1)
	goto L19
L18:
	;
	v155 = int32(1)
	goto L19
L19:
	;
	v161 = v155
	goto L13
L20:
	;
	v173 = v120
	goto L9
L21:
	;
	goto L22
L22:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v43+v120<<(uint(int32(3))%32)))) = v167
	if v73 <= v142 {
		v120 = v142
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L12
L24:
	;
	goto L8
L25:
	;
	goto L3
L26:
	;
	v241 = F_palloc(m, v235*int32(6))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v773 = v235
	v778 = v3
	goto L28
L28:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L109
L29:
	;
	return int32(0)
L30:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v245 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_pg_qsort(m, v241, v487, int32(6), int32(141))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L29
	} else {
		goto L72
	}
L32:
	;
	v248 = int32(0)
	v487 = v248
	v510 = v248
	goto L31
L33:
	;
	goto L34
L34:
	;
	v250 = int32(0)
	v255 = v250
	v257 = v250
	v258 = int32(-1)
	goto L35
L35:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v284 = v281 + v255<<(uint(int32(3))%32)
	v285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v284)+6)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v285*int32(6))+3)))
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+2)))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	v294 = v290 | v291<<(uint(int32(16))%32)
	if v294 != v258 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v332 = int32(0)
	if v314 <= v332 {
		v487 = v314
		v510 = v332
		goto L31
	} else {
		goto L45
	}
L37:
	;
	if v289&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v298 = v241 + v257*int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v298)+4)) = uint16(v255)
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = int32(_a_F_heap_index_delete_tuples_1)
	v314 = v257 + int32(1)
	v315 = v294
	goto L37
L39:
	;
	goto L40
L40:
	;
	v308 = v241 + v257*int32(6) - int32(4)
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308))))
	v311 = v309 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v308))) = uint16(v311)
	v314 = v257
	v315 = v258
	goto L37
L41:
	;
	v318 = int32(6)
	v322 = v241 + v314*v318 - v318
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322))))
	v325 = v323 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v325)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v329 = v255 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v329 < v330 {
		v255 = v329
		v257 = v314
		v258 = v315
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v335 = int32(0)
	if v314 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v487 = v314
	v510 = int32(1)
	goto L31
L47:
	;
	v346 = v335
	v351 = int32(0)
	goto L50
L48:
	;
	v413 = v335
	goto L49
L49:
	;
	v440 = v241 + v413*int32(6)
	v441 = int32(*(*int16)(unsafe.Add(mBase, uint32(v440))))
	if int32(5) <= v441 {
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v373 = v241 + v346*int32(6)
	v374 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373))))
	if int32(5) <= v374 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v314&int32(1) == int32(0) {
		goto L46
	} else {
		goto L65
	}
L52:
	;
	v377 = int32(1)
	if v374&(v374-v377) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v386 = int32(4)
	goto L54
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v373))) = uint16(v386)
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373)+6)))
	if int32(5) <= v388 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v385 = v377 << (uint(int32(32)-base.I32_clz(v374)) % 32)
	goto L57
L56:
	;
	v385 = v374
	goto L57
L57:
	;
	v386 = v385
	goto L54
L58:
	;
	v391 = int32(1)
	if v388&(v388-v391) != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v401 = int32(4)
	goto L60
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+6)) = uint16(v401)
	v403 = int32(2)
	v404 = v346 + v403
	v406 = v351 + v403
	if v406 != v314&int32(2147483646) {
		v346 = v404
		v351 = v406
		goto L50
	} else {
		goto L64
	}
L61:
	;
	v399 = v391 << (uint(int32(32)-base.I32_clz(v388)) % 32)
	goto L63
L62:
	;
	v399 = v388
	goto L63
L63:
	;
	v401 = v399
	goto L60
L64:
	;
	goto L51
L65:
	;
	v413 = v404
	goto L49
L66:
	;
	v444 = int32(1)
	if v441&(v441-v444) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v453 = int32(4)
	goto L68
L68:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v440))) = uint16(v453)
	goto L46
L69:
	;
	v452 = v444 << (uint(int32(32)-base.I32_clz(v441)) % 32)
	goto L71
L70:
	;
	v452 = v441
	goto L71
L71:
	;
	v453 = v452
	goto L68
L72:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v518 = F_palloc(m, v515<<(uint(int32(3))%32))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L29
	} else {
		goto L73
	}
L73:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v510 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v757 = v737 << (uint(int32(3)) % 32)
	if v757 != 0 {
		goto L103
	} else {
		goto L104
	}
L75:
	;
	v732 = v520
	v737 = int32(0)
	v742 = v3
	goto L74
L76:
	;
	goto L77
L77:
	;
	v524 = int32(6)
	if v524 <= v487 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v527 = v524
	goto L80
L79:
	;
	v527 = v487
	goto L80
L80:
	;
	v542 = v3
	v554 = int64(-1)
	goto L82
L81:
	;
	v583 = int32(0)
	if v487 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v559 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241+v542*int32(6))+4)))
	v562 = v520 + v559<<(uint(int32(3))%32)
	v563 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v562))))
	v566 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v562)+2)))
	v567 = v563<<(uint(int64(16))%64) | v566
	v568 = int64(3)
	if (base.B2i32(v567 < v554-v568)|base.B2i32(v554+v568 < v567))&base.B2i32(v554 != int64(-1)) != 0 {
		v581 = v542
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v581 = v527
	goto L81
L84:
	;
	v579 = v542 + int32(1)
	if v579 != v527 {
		v542 = v579
		v554 = v567
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v732 = v728
	v737 = v709
	v742 = v581
	goto L74
L87:
	;
	v595 = v583
	v597 = int32(0)
	v600 = v583
	goto L90
L88:
	;
	v660 = v583
	v665 = v583
	goto L89
L89:
	;
	v686 = v241 + v660*int32(6)
	v687 = int32(*(*int16)(unsafe.Add(mBase, uint32(v686)+2)))
	v689 = v687 << (uint(int32(3)) % 32)
	if v689 != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	v621 = v241 + v595*int32(6)
	v622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+2)))
	v624 = v622 << (uint(int32(3)) % 32)
	if v624 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v527&int32(1) == int32(0) {
		v709 = v649
		goto L86
	} else {
		goto L99
	}
L92:
	;
	v625 = int32(3)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v629 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+4)))
	base.MemoryCopy(m, v518+v600<<(uint(v625)%32), v628+v629<<(uint(v625)%32), v624)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v634 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+2)))
	v635 = v600 + v634
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+8)))
	v638 = v636 << (uint(int32(3)) % 32)
	if v638 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v639 = int32(3)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+10)))
	base.MemoryCopy(m, v518+v635<<(uint(v639)%32), v642+v643<<(uint(v639)%32), v638)
	goto L97
L96:
	;
	goto L97
L97:
	;
	v648 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621)+8)))
	v649 = v635 + v648
	v650 = int32(2)
	v651 = v595 + v650
	v653 = v597 + v650
	if v653 != v527&int32(-2) {
		v595 = v651
		v597 = v653
		v600 = v649
		goto L90
	} else {
		goto L98
	}
L98:
	;
	goto L91
L99:
	;
	v660 = v651
	v665 = v649
	goto L89
L100:
	;
	v690 = int32(3)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v694 = int32(*(*int16)(unsafe.Add(mBase, uint32(v686)+4)))
	base.MemoryCopy(m, v518+v665<<(uint(v690)%32), v693+v694<<(uint(v690)%32), v689)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v699 = int32(*(*int16)(unsafe.Add(mBase, uint32(v686)+2)))
	v709 = v665 + v699
	goto L86
L103:
	;
	base.MemoryCopy(m, v732, v518, v757)
	goto L105
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v737
	F_pfree(m, v518)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L29
	} else {
		goto L106
	}
L106:
	;
	F_pfree(m, v241)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v773 = v764
	v778 = v742
	goto L28
L108:
	;
	v803 = int32(0)
	if v773 <= v803 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	if base.Ui32(v793) < base.Ui32(int32(_a_F_heap_index_delete_tuples_2)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[2]))
	v802 = v797
	goto L108
L111:
	;
	goto L112
L112:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)+92))
	v800 = F_get_tablespace_maintenance_io_concurrency(m, v799)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v802 = v800
	goto L108
L114:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v901 {
		goto L135
	} else {
		goto L136
	}
L115:
	;
	v877 = int32(0)
	v878 = int32(-1)
	goto L114
L116:
	;
	if v802 < v778 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v808 = v802
	goto L119
L118:
	;
	v808 = v778
	goto L119
L119:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v809 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v810 = v808
	goto L122
L121:
	;
	v810 = v802
	goto L122
L122:
	;
	if v810 <= int32(0) {
		goto L115
	} else {
		goto L123
	}
L123:
	;
	v813 = int32(0)
	v817 = int32(-1)
	v818 = v813
	v820 = v813
	goto L124
L124:
	;
	v844 = v792 + v818<<(uint(int32(3))%32)
	v845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844))))
	if v817 == int32(-1) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v877 = v868
	v878 = v865
	goto L114
L126:
	;
	v868 = v818 + int32(1)
	if v773 <= v868 {
		v877 = v868
		v878 = v865
		goto L114
	} else {
		goto L133
	}
L127:
	;
	F_PrefetchBuffer(m, v30+int32(84), l0, int32(0), v857)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L29
	} else {
		goto L132
	}
L128:
	;
	v848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844)+2)))
	v857 = v848 | v845<<(uint(int32(16))%32)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844)+2)))
	v855 = v852 | v845<<(uint(int32(16))%32)
	if v855 != v817 {
		v857 = v855
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v865 = v817
	v866 = v820
	goto L126
L132:
	;
	v865 = v857
	v866 = v820 + int32(1)
	goto L126
L133:
	;
	if v866 < v810 {
		v817 = v865
		v818 = v868
		v820 = v866
		goto L124
	} else {
		goto L134
	}
L134:
	;
	goto L125
L135:
	;
	v908 = v877
	v909 = v878
	v914 = int32(-1)
	v918 = v778
	v919 = v3
	v921 = v3
	v922 = v34
	v923 = v3
	v924 = v3
	v925 = v3
	v926 = v3
	v927 = v3
	v928 = v3
	v929 = v3
	goto L138
L136:
	;
	v1557 = v803
	v1565 = v3
	v1572 = v3
	goto L137
L137:
	;
	F_UnlockReleaseBuffer(m, v1565)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L29
	} else {
		goto L244
	}
L138:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v935 = v932 + v921<<(uint(int32(3))%32)
	v936 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935)+6)))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v914 != int32(-1) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v30)+188))
	v1557 = v1550
	v1565 = v1537
	v1572 = v1544
	goto L137
L140:
	;
	goto L139
L141:
	;
	v1145 = v937 + v936*int32(6)
	v1146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+4)))
	v1148 = v1140 & int32(_a_F_heap_index_delete_tuples_0)
	if base.Ui32(v1146) <= base.Ui32(v1148) {
		goto L185
	} else {
		goto L186
	}
L142:
	;
	v940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+2)))
	v941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935))))
	if v940|v941<<(uint(int32(16))%32) == v914 {
		v1119 = v908
		v1120 = v909
		v1125 = v914
		v1129 = v918
		v1130 = v919
		v1133 = v922
		v1136 = v925
		v1138 = v927
		v1139 = v928
		v1140 = v929
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v946 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	if v923&int32(1)|base.B2i32(v924 == v927)&base.B2i32(int32(0) < v928) != 0 {
		v1537 = v919
		v1544 = v926
		goto L140
	} else {
		goto L149
	}
L147:
	;
	v964 = v918
	v965 = v922
	v966 = v927
	goto L148
L148:
	;
	if v919 != 0 {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	if int32(0) < v918 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v964 = v962
	v965 = v963
	v966 = v924
	goto L148
L151:
	;
	v962 = v918 - int32(1)
	v963 = v922
	goto L150
L152:
	;
	goto L153
L153:
	;
	v961 = base.I32_div_s(v922, int32(2))
	v962 = v918
	v963 = v961
	goto L150
L154:
	;
	F_UnlockReleaseBuffer(m, v919)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L29
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+2)))
	v970 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935))))
	v973 = v969 | v970<<(uint(int32(16))%32)
	v974 = F_ReadBuffer(m, l0, v973)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L29
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	if v908 < v773 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	F_LockBuffer(m, v974, int32(1))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L29
	} else {
		goto L173
	}
L160:
	;
	v980 = v908
	goto L163
L161:
	;
	v1033 = v908
	goto L162
L162:
	;
	v1059 = v909
	v1060 = v1033
	goto L159
L163:
	;
	v1006 = v792 + v980<<(uint(int32(3))%32)
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006))))
	if v909 == int32(-1) {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v1033 = v1028
	goto L162
L165:
	;
	v1028 = v980 + int32(1)
	if v1028 < v773 {
		v980 = v1028
		goto L163
	} else {
		goto L172
	}
L166:
	;
	F_PrefetchBuffer(m, v30+int32(84), l0, int32(0), v1019)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L29
	} else {
		goto L171
	}
L167:
	;
	v1010 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006)+2)))
	v1019 = v1010 | v1007<<(uint(int32(16))%32)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006)+2)))
	v1017 = v1014 | v1007<<(uint(int32(16))%32)
	if v1017 == v909 {
		goto L165
	} else {
		goto L170
	}
L170:
	;
	v1019 = v1017
	goto L166
L171:
	;
	v1059 = v1019
	v1060 = v980 + int32(1)
	goto L159
L172:
	;
	goto L164
L173:
	;
	if v974 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1107) {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[3]))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1092+(v974^int32(-1))<<(uint(int32(2))%32))))
	v1106 = v1098
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[4]))
	v1106 = v1100 + v974<<(uint(int32(13))%32) + int32(-8192)
	goto L174
L178:
	;
	v1115 = int32(base.Ui32(v1107+int32(_a_F_heap_index_delete_tuples_3)) >> (uint(int32(2)) % 32))
	goto L180
L179:
	;
	v1115 = int32(0)
	goto L180
L180:
	;
	v1119 = v1060
	v1120 = v1059
	v1125 = v973
	v1129 = v964
	v1130 = v974
	v1133 = v965
	v1136 = v1106
	v1138 = v966
	v1139 = v928 + int32(1)
	v1140 = v1115
	goto L141
L181:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1510 < v1521 {
		v908 = v1119
		v909 = v1120
		v914 = v1125
		v918 = v1129
		v919 = v1130
		v921 = v1510
		v922 = v1133
		v923 = v1512
		v924 = v1513
		v925 = v1136
		v926 = v1515
		v927 = v1138
		v928 = v1139
		v929 = v1140
		goto L138
	} else {
		goto L243
	}
L182:
	;
	v1510 = v921 + int32(1)
	v1512 = v923
	v1513 = v924
	v1515 = v926
	goto L181
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L29
	} else {
		goto L239
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L29
	} else {
		goto L235
	}
L185:
	;
	v1151 = v1136 + int32(20)
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1151+v1146<<(uint(int32(2))%32))))
	if v1155&int32(_a_F_heap_index_delete_tuples_4) == int32(0) {
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L29
	} else {
		goto L231
	}
L188:
	;
	if base.Ui32(int32(_a_F_heap_index_delete_tuples_5)) <= base.Ui32(v1155) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1136+v1155&int32(_a_F_heap_index_delete_tuples_6))+18)))
	if v1165 < int32(0) {
		goto L183
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+2)))
	if v1168 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L191
L193:
	;
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+108)) = uint16(v1171)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v935)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v1173
	v1183 = F_heap_hot_search_buffer(m, v30+int32(104), l0, v1130, v30+int32(112), v30+int32(84), int32(0), int32(1))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L29
	} else {
		goto L196
	}
L194:
	;
	v1197 = v1146
	v1198 = v923
	v1199 = v924
	goto L195
L195:
	;
	if base.Ui32(v1148) <= base.Ui32((v1197-int32(1))&int32(_a_F_heap_index_delete_tuples_0)) {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	if v1183 != 0 {
		goto L182
	} else {
		goto L197
	}
L197:
	;
	v1185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1145)+2)) = uint8(v1185)
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v1187 == v1185 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1145)+4)))
	v1191 = v924 + v1190
	v1194 = base.B2i32(v1133 <= v1191) | v923
	v1195 = v1191
	goto L200
L199:
	;
	v1194 = v923
	v1195 = v924
	goto L200
L200:
	;
	v1196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+4)))
	v1197 = v1196
	v1198 = v1194
	v1199 = v1195
	goto L195
L201:
	;
	v1394 = v921 + int32(1)
	v1510 = v1394
	v1512 = v1198
	v1513 = v1199
	v1515 = v1394
	goto L181
L202:
	;
	v1211 = v1197
	v1212 = int32(0)
	goto L203
L203:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1151+v1211&int32(_a_F_heap_index_delete_tuples_0)<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v1238)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L206
	case 1:
		goto L207
	default:
		goto L201
	}
L204:
	;
	goto L201
L205:
	;
	if base.Ui32((v1339-int32(1))&int32(_a_F_heap_index_delete_tuples_0)) < base.Ui32(v1148) {
		v1211 = v1339
		v1212 = v1340
		goto L203
	} else {
		goto L230
	}
L206:
	;
	v1249 = v1136 + v1238&int32(_a_F_heap_index_delete_tuples_6)
	if v1212 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1339 = v1238 & int32(_a_F_heap_index_delete_tuples_6)
	v1340 = v1212
	goto L205
L208:
	;
	v1250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1249)+20)))
	v1251 = int32(768)
	if v1250&v1251 != v1251 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1249, v30+int32(188))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L29
	} else {
		goto L215
	}
L211:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1249)))
	v1257 = v1255
	goto L213
L212:
	;
	v1257 = int32(2)
	goto L213
L213:
	;
	if v1257 != v1212 {
		goto L201
	} else {
		goto L214
	}
L214:
	;
	goto L210
L215:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249)+19)))
	if v1263&int32(64) == int32(0) {
		goto L201
	} else {
		goto L216
	}
L216:
	;
	v1268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1249)+20)))
	if v1268&int32(2048)|base.B2i32(v1268&int32(768) == int32(512)) != 0 {
		goto L201
	} else {
		goto L217
	}
L217:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+4))
	v1277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1249)+16)))
	if v1268&int32(_a_F_heap_index_delete_tuples_7) != int32(_a_F_heap_index_delete_tuples_8) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1339 = v1277
	v1340 = v1276
	goto L205
L219:
	;
	goto L220
L220:
	;
	v1282 = int32(0)
	v1286 = F_GetMultiXactIdMembers(m, v1276, v30+int32(84), v1282)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L29
	} else {
		goto L221
	}
L221:
	;
	if v1286 <= int32(0) {
		v1339 = v1277
		v1340 = v1282
		goto L205
	} else {
		goto L222
	}
L222:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v1294 = int32(0)
	goto L225
L223:
	;
	F_pfree(m, v1291)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L29
	} else {
		goto L229
	}
L224:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1331 = v1329
	goto L223
L225:
	;
	v1321 = v1291 + v1294<<(uint(int32(3))%32)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1322) {
		goto L224
	} else {
		goto L227
	}
L226:
	;
	v1331 = int32(0)
	goto L223
L227:
	;
	v1326 = v1294 + int32(1)
	if v1326 != v1286 {
		v1294 = v1326
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1339 = v1277
	v1340 = v1331
	goto L205
L230:
	;
	goto L204
L231:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L29
	} else {
		goto L232
	}
L232:
	;
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+2)))
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935))))
	v1404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145))))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v1407 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1402 | v1403<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_9), v30)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L29
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_11), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L29
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L29
	} else {
		goto L236
	}
L236:
	;
	v1433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+2)))
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935))))
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145))))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v1438 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v1436
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v1435
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1433 | v1434<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_13), v30+int32(32))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L29
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_14), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L29
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L29
	} else {
		goto L240
	}
L240:
	;
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+2)))
	v1467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935))))
	v1468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145))))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v1471 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+76)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1466 | v1467<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_15), v30-int32(-64))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L29
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_16), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L29
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	v1537 = v1130
	v1544 = v1515
	goto L140
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v1572
	m.G0 = v30 + int32(192)
	return v1557
}
func F_heap_lock_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
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
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v567 int32
	_ = v567
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v734 int32
	_ = v734
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1372 int64
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1497 int32
	_ = v1497
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1537 int32
	_ = v1537
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	v9 = int32(0)
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v9
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v37 = F_ReadBuffer(m, l0, v32|v33<<(uint(int32(16))%32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v37
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v46 = v42 | v43<<(uint(int32(16))%32)
	if v37 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+10)))
	if v65&int32(4) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v37^int32(-1))<<(uint(int32(2))%32))))
	v64 = v56
	goto L3
L5:
	;
	goto L6
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[1]))
	v64 = v58 + v37<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	F_visibilitymap_pin(m, l0, v46, v26+int32(-4))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v73 = v37
	goto L9
L9:
	;
	F_LockBuffer(m, v73, int32(2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v73 = v72
	goto L9
L11:
	;
	v78 = l1 + int32(4)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v79 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v103 = v97 + v98<<(uint(int32(2))%32) + int32(20)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v104&int32(_a_F_heap_lock_tuple_0) + v97
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v109) >> (uint(int32(17)) % 32))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v117 = F_HeapTupleSatisfiesUpdate(m, l1, l2, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[0]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v79^int32(-1))<<(uint(int32(2))%32))))
	v97 = v89
	goto L12
L14:
	;
	goto L15
L15:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[1]))
	v97 = v91 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+60))
	if v1628 != 0 {
		goto L362
	} else {
		goto L363
	}
L17:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1580)))
	F_LockBuffer(m, v1599, int32(0))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L361
	}
L18:
	;
	if v117 == int32(1) {
		v1574 = l0
		v1577 = l3
		v1580 = l6
		v1582 = int32(1)
		v1583 = v28
		v1588 = v78
		v1589 = v9
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v121 = int32(1)
	v128 = l0
	v129 = l1
	v130 = l2
	v131 = l3
	v132 = l4
	v133 = l5
	v134 = l6
	v135 = l7
	v137 = v28
	v139 = v117
	v141 = v121
	v142 = v78
	v143 = v9
	v144 = v97
	v145 = l3*int32(12) + int32(_a_F_heap_lock_tuple_1)
	v146 = v46
	v147 = v9
	v149 = l5 ^ v121
	goto L20
L20:
	;
	v154 = v139 - int32(3)
	if base.Ui32(v154) <= base.Ui32(int32(2)) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1468)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)) = uint16(v1469)
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1468)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v1471
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1473)+4))
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+20)))
	if v1475&int32(_a_F_heap_lock_tuple_2) != int32(_a_F_heap_lock_tuple_3) {
		goto L342
	} else {
		goto L343
	}
L22:
	;
	goto L21
L23:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v1439 = F_HeapTupleSatisfiesUpdate(m, v129, v130, v1438)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L339
	}
L24:
	;
	v1425 = int32(0)
	v1427 = v1401
	v1431 = v1405
	goto L23
L25:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+18)))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+20)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+48)) = uint16(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+44)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v165, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v1203 = v139
	v1205 = v141
	v1207 = v143
	v1211 = v147
	goto L27
L27:
	;
	if v1203 != 0 {
		v1454 = v1203
		v1458 = v1207
		goto L22
	} else {
		goto L309
	}
L28:
	;
	if v141&int32(1) == int32(0) {
		v567 = v147
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v1203 = v1177
	v1205 = int32(0)
	v1207 = v1181
	v1211 = v1185
	goto L27
L30:
	;
	v782 = v159 & int32(_a_F_heap_lock_tuple_3)
	if v782 != 0 {
		goto L176
	} else {
		goto L177
	}
L31:
	;
	if v159&int32(80) != int32(16) {
		v775 = v734
		goto L30
	} else {
		goto L172
	}
L32:
	;
	if v159&int32(_a_F_heap_lock_tuple_3) == int32(0) {
		v734 = v567
		goto L31
	} else {
		goto L166
	}
L33:
	;
	v668 = int32(64)
	if base.B2i32(v159&int32(128) == int32(0))&base.B2i32(v159&int32(_a_F_heap_lock_tuple_4) != v668)|base.B2i32(v645 == v668) != 0 {
		v775 = v656
		goto L30
	} else {
		goto L163
	}
L34:
	;
	switch v131 {
	case 0:
		goto L144
	case 1:
		goto L143
	case 2:
		goto L32
	default:
		v775 = v567
		goto L30
	}
L35:
	;
	if v159&int32(_a_F_heap_lock_tuple_3) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v1603 = v128
	v1606 = v131
	v1611 = int32(0)
	v1612 = v137
	v1617 = v142
	v1618 = v143
	goto L16
L37:
	;
	F_pfree(m, v337)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L142
	}
L38:
	;
	v182 = F_GetMultiXactIdMembers(m, v160, v137+int32(56), int32(base.Ui32(v159&int32(128))>>(uint(int32(7))%32)))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(v160) < base.Ui32(int32(3)) {
		goto L95
	} else {
		goto L96
	}
L41:
	;
	if int32(0) < v182 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v194 = int32(0)
	v205 = v147
	goto L45
L43:
	;
	v368 = v147
	goto L44
L44:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v137)+56))
	if v374 == int32(0) {
		v567 = v368
		goto L34
	} else {
		goto L92
	}
L45:
	;
	v211 = int32(3)
	v212 = v194 << (uint(v211) % 32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v137)+56))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212+v213)))
	if base.Ui32(v215) < base.Ui32(v211) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v368 = v344
	goto L44
L47:
	;
	if v335 != 0 {
		goto L87
	} else {
		goto L88
	}
L48:
	;
	v335 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v226 == v215 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v335 = int32(1)
	goto L47
L52:
	;
	goto L53
L53:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v230 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v335 = v327
	goto L47
L55:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v234 == int32(0) {
		v327 = int32(0)
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v298 = int32(0)
	v300 = v230 - int32(1)
	goto L77
L58:
	;
	v239 = v234
	goto L59
L59:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	if v244 == int32(4) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v327 = int32(0)
	goto L54
L61:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v239)+80))
	if v291 != 0 {
		v239 = v291
		goto L59
	} else {
		goto L76
	}
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v247 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v250 = int32(1)
	if v215 == v247 {
		v327 = v250
		goto L54
	} else {
		goto L64
	}
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v254 = v252 - int32(1)
	if v254 < int32(0) {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v259 = int32(0)
	v261 = v254
	goto L66
L66:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v239)+48))
	v267 = int32(2)
	v268 = base.I32_div_s(v261-v259, v267)
	v269 = v268 + v259
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v265+v269<<(uint(v267)%32))))
	if v273 == v215 {
		v327 = v250
		goto L54
	} else {
		goto L68
	}
L67:
	;
	goto L61
L68:
	;
	v277 = F_TransactionIdPrecedes(m, v273, v215)
	mBase = m.M
	if v277 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v278 = v269 + int32(1)
	goto L71
L70:
	;
	v278 = v259
	goto L71
L71:
	;
	if v277 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v281 = v261
	goto L74
L73:
	;
	v281 = v269 - int32(1)
	goto L74
L74:
	;
	if v278 <= v281 {
		v259 = v278
		v261 = v281
		goto L66
	} else {
		goto L75
	}
L75:
	;
	goto L67
L76:
	;
	goto L60
L77:
	;
	v305 = int32(2)
	v306 = base.I32_div_s(v300-v298, v305)
	v307 = v306 + v298
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v296+v307<<(uint(v305)%32))))
	v312 = base.B2i32(v311 == v215)
	if v311 == v215 {
		v327 = v312
		goto L54
	} else {
		goto L79
	}
L78:
	;
	v327 = v312
	goto L54
L79:
	;
	v315 = base.B2i32(base.Ui32(v311) < base.Ui32(v215))
	if base.Ui32(v311) < base.Ui32(v215) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v316 = v307 + int32(1)
	goto L82
L81:
	;
	v316 = v298
	goto L82
L82:
	;
	if base.Ui32(v311) < base.Ui32(v215) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v319 = v300
	goto L85
L84:
	;
	v319 = v307 - int32(1)
	goto L85
L85:
	;
	if v316 <= v319 {
		v298 = v316
		v300 = v319
		goto L77
	} else {
		goto L86
	}
L86:
	;
	goto L78
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v137)+56))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v212+v337)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v339<<(uint(int32(2))%32))+uint32(_c_F_heap_lock_tuple[6])))
	if base.Ui32(v131) <= base.Ui32(v342) {
		goto L37
	} else {
		goto L90
	}
L88:
	;
	v344 = v205
	goto L89
L89:
	;
	v347 = v194 + int32(1)
	if v347 != v182 {
		v194 = v347
		v205 = v344
		goto L45
	} else {
		goto L91
	}
L90:
	;
	v344 = int32(1)
	goto L89
L91:
	;
	goto L46
L92:
	;
	F_pfree(m, v374)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v567 = v368
	goto L34
L94:
	;
	if v498 == int32(0) {
		v567 = v147
		goto L34
	} else {
		goto L134
	}
L95:
	;
	v498 = int32(0)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v389 == v160 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v498 = int32(1)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v393 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v498 = v490
	goto L94
L102:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v397 == int32(0) {
		v490 = int32(0)
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v461 = int32(0)
	v463 = v393 - int32(1)
	goto L124
L105:
	;
	v402 = v397
	goto L106
L106:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+20))
	if v407 == int32(4) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v490 = int32(0)
	goto L101
L108:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v402)+80))
	if v454 != 0 {
		v402 = v454
		goto L106
	} else {
		goto L123
	}
L109:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if v410 == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v413 = int32(1)
	if v160 == v410 {
		v490 = v413
		goto L101
	} else {
		goto L111
	}
L111:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v402)+52))
	v417 = v415 - int32(1)
	if v417 < int32(0) {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v422 = int32(0)
	v424 = v417
	goto L113
L113:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v402)+48))
	v430 = int32(2)
	v431 = base.I32_div_s(v424-v422, v430)
	v432 = v431 + v422
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v428+v432<<(uint(v430)%32))))
	if v436 == v160 {
		v490 = v413
		goto L101
	} else {
		goto L115
	}
L114:
	;
	goto L108
L115:
	;
	v440 = F_TransactionIdPrecedes(m, v436, v160)
	mBase = m.M
	if v440 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v441 = v432 + int32(1)
	goto L118
L117:
	;
	v441 = v422
	goto L118
L118:
	;
	if v440 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v444 = v424
	goto L121
L120:
	;
	v444 = v432 - int32(1)
	goto L121
L121:
	;
	if v441 <= v444 {
		v422 = v441
		v424 = v444
		goto L113
	} else {
		goto L122
	}
L122:
	;
	goto L114
L123:
	;
	goto L107
L124:
	;
	v468 = int32(2)
	v469 = base.I32_div_s(v463-v461, v468)
	v470 = v469 + v461
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v459+v470<<(uint(v468)%32))))
	v475 = base.B2i32(v474 == v160)
	if v474 == v160 {
		v490 = v475
		goto L101
	} else {
		goto L126
	}
L125:
	;
	v490 = v475
	goto L101
L126:
	;
	v478 = base.B2i32(base.Ui32(v474) < base.Ui32(v160))
	if base.Ui32(v474) < base.Ui32(v160) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v479 = v470 + int32(1)
	goto L129
L128:
	;
	v479 = v461
	goto L129
L129:
	;
	if base.Ui32(v474) < base.Ui32(v160) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v482 = v463
	goto L132
L131:
	;
	v482 = v470 - int32(1)
	goto L132
L132:
	;
	if v479 <= v482 {
		v461 = v479
		v463 = v482
		goto L124
	} else {
		goto L133
	}
L133:
	;
	goto L125
L134:
	;
	switch v131 {
	case 0:
		goto L36
	case 1:
		goto L137
	case 2:
		goto L135
	case 3:
		goto L136
	default:
		v567 = v147
		goto L34
	}
L135:
	;
	if v159&int32(80) == int32(64) {
		goto L36
	} else {
		goto L141
	}
L136:
	;
	if v159&int32(80) != int32(64) {
		v775 = v147
		goto L30
	} else {
		goto L139
	}
L137:
	;
	v502 = v159 & int32(80)
	v504 = v502 + int32(-64)
	if base.B2i32(v504 == int32(0))|base.B2i32(v504 == int32(16)) != 0 {
		goto L36
	} else {
		goto L138
	}
L138:
	;
	v645 = v502
	v656 = v147
	goto L33
L139:
	;
	if v158&int32(_a_F_heap_lock_tuple_5) != 0 {
		goto L36
	} else {
		goto L140
	}
L140:
	;
	v775 = v147
	goto L30
L141:
	;
	v734 = v147
	goto L31
L142:
	;
	goto L36
L143:
	;
	v645 = v159 & int32(80)
	v656 = v567
	goto L33
L144:
	;
	if v158&int32(_a_F_heap_lock_tuple_5) != 0 {
		v775 = v567
		goto L30
	} else {
		goto L145
	}
L145:
	;
	v583 = int32(base.Ui32(v159&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v159&int32(_a_F_heap_lock_tuple_4) == int32(64))
	if (v149|v583)&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v618, int32(2))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L159
	}
L147:
	;
	v588 = v137 + int32(44)
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+2)))
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142))))
	v591 = int32(16)
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588)+2)))
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588))))
	if v589|v590<<(uint(v591)%32) == v594|v595<<(uint(v591)%32) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	if v605 != 0 {
		goto L146
	} else {
		goto L154
	}
L149:
	;
	goto L148
L150:
	;
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588)+4)))
	if v601 == v602 {
		v605 = int32(1)
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v605 = int32(0)
	goto L149
L153:
	;
	goto L152
L154:
	;
	v606 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v609 = F_heap_lock_updated_tuple(m, v128, v159, v160, v588, v606, int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v609 == int32(0) {
		goto L146
	} else {
		goto L157
	}
L157:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v613, int32(2))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v1454 = v609
	v1458 = v143
	goto L22
L159:
	;
	v622 = int32(0)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v624 = F_HeapTupleHeaderIsOnlyLocked(m, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if v624 != 0 {
		v1177 = v622
		v1181 = v143
		v1185 = v567
		goto L29
	} else {
		goto L161
	}
L161:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+19)))
	if (int32(base.Ui32(v627)>>(uint(int32(5))%32))|v583)&int32(1) == int32(0) {
		v1177 = v622
		v1181 = v143
		v1185 = v567
		goto L29
	} else {
		goto L162
	}
L162:
	;
	v1401 = v143
	v1405 = v567
	goto L24
L163:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v674, int32(2))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+20)))
	v682 = int32(64)
	if base.B2i32(v679&int32(80) == v682)|base.B2i32(v679&int32(128) == int32(0))&base.B2i32(v679&int32(_a_F_heap_lock_tuple_4) != v682) != 0 {
		v1401 = v143
		v1405 = v656
		goto L24
	} else {
		goto L165
	}
L165:
	;
	v1177 = int32(0)
	v1181 = v143
	v1185 = v656
	goto L29
L166:
	;
	v701 = F_DoesMultiXactIdConflict(m, v160, v159, int32(2), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	if v701 != 0 {
		v775 = v567
		goto L30
	} else {
		goto L168
	}
L168:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v703, int32(2))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v707)+20)))
	if (v708^v159)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1401 = v143
		v1405 = v567
		goto L24
	} else {
		goto L170
	}
L170:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	if v712 != v160 {
		v1401 = v143
		v1405 = v567
		goto L24
	} else {
		goto L171
	}
L171:
	;
	v1177 = int32(0)
	v1181 = v143
	v1185 = v567
	goto L29
L172:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v744, int32(2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v748)+20)))
	if (v749^v159)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1401 = v143
		v1405 = v734
		goto L24
	} else {
		goto L174
	}
L174:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	if v753 != v160 {
		v1401 = v143
		v1405 = v734
		goto L24
	} else {
		goto L175
	}
L175:
	;
	v1177 = int32(0)
	v1181 = v143
	v1185 = v734
	goto L29
L176:
	;
	if v154 != int32(2) {
		goto L222
	} else {
		goto L223
	}
L177:
	;
	if base.Ui32(v160) < base.Ui32(int32(3)) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v902 == int32(0) {
		goto L176
	} else {
		goto L218
	}
L179:
	;
	v902 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v793 == v160 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v902 = int32(1)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v797 <= int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v902 = v894
	goto L178
L186:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v801 == int32(0) {
		v894 = int32(0)
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v865 = int32(0)
	v867 = v797 - int32(1)
	goto L208
L189:
	;
	v806 = v801
	goto L190
L190:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v806)+20))
	if v811 == int32(4) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v894 = int32(0)
	goto L185
L192:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v806)+80))
	if v858 != 0 {
		v806 = v858
		goto L190
	} else {
		goto L207
	}
L193:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	if v814 == int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v817 = int32(1)
	if v160 == v814 {
		v894 = v817
		goto L185
	} else {
		goto L195
	}
L195:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v806)+52))
	v821 = v819 - int32(1)
	if v821 < int32(0) {
		goto L192
	} else {
		goto L196
	}
L196:
	;
	v826 = int32(0)
	v828 = v821
	goto L197
L197:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v806)+48))
	v834 = int32(2)
	v835 = base.I32_div_s(v828-v826, v834)
	v836 = v835 + v826
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v832+v836<<(uint(v834)%32))))
	if v840 == v160 {
		v894 = v817
		goto L185
	} else {
		goto L199
	}
L198:
	;
	goto L192
L199:
	;
	v844 = F_TransactionIdPrecedes(m, v840, v160)
	mBase = m.M
	if v844 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v845 = v836 + int32(1)
	goto L202
L201:
	;
	v845 = v826
	goto L202
L202:
	;
	if v844 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v848 = v828
	goto L205
L204:
	;
	v848 = v836 - int32(1)
	goto L205
L205:
	;
	if v845 <= v848 {
		v826 = v845
		v828 = v848
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
	v872 = int32(2)
	v873 = base.I32_div_s(v867-v865, v872)
	v874 = v873 + v865
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v863+v874<<(uint(v872)%32))))
	v879 = base.B2i32(v878 == v160)
	if v878 == v160 {
		v894 = v879
		goto L185
	} else {
		goto L210
	}
L209:
	;
	v894 = v879
	goto L185
L210:
	;
	v882 = base.B2i32(base.Ui32(v878) < base.Ui32(v160))
	if base.Ui32(v878) < base.Ui32(v160) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v883 = v874 + int32(1)
	goto L213
L212:
	;
	v883 = v865
	goto L213
L213:
	;
	if base.Ui32(v878) < base.Ui32(v160) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v886 = v867
	goto L216
L215:
	;
	v886 = v874 - int32(1)
	goto L216
L216:
	;
	if v883 <= v886 {
		v865 = v883
		v867 = v886
		goto L208
	} else {
		goto L217
	}
L217:
	;
	goto L209
L218:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v905, int32(2))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v910 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909)+20)))
	if (v910^v159)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1401 = v143
		v1405 = v775
		goto L24
	} else {
		goto L220
	}
L220:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	if v914 != v160 {
		v1401 = v143
		v1405 = v775
		goto L24
	} else {
		goto L221
	}
L221:
	;
	v1177 = int32(0)
	v1181 = v143
	v1185 = v775
	goto L29
L222:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v919, int32(2))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if (v143|v775)&int32(1) != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v1177 = v139
	v1181 = v143
	v1185 = v775
	goto L29
L226:
	;
	if v782 != 0 {
		goto L245
	} else {
		goto L246
	}
L227:
	;
	v970 = base.B2i32(v775 == int32(0)) | v143
	goto L226
L228:
	;
	goto L229
L229:
	;
	v929 = int32(1)
	switch v132 {
	case 0:
		goto L230
	case 1:
		goto L231
	case 2:
		goto L232
	default:
		v970 = v929
		goto L226
	}
L230:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	F_LockTuple(m, v128, v142, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L242
	}
L231:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v958 = F_ConditionalLockTuple(m, v128, v142, v956, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L239
	}
L232:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v933 = F_ConditionalLockTuple(m, v128, v142, v930, v932)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	if v933 != 0 {
		v970 = v929
		goto L226
	} else {
		goto L234
	}
L234:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+32)) = v942 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v137+int32(32))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_9), int32(_a_F_heap_lock_tuple_10))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	if v958 != 0 {
		v970 = v929
		goto L226
	} else {
		goto L240
	}
L240:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v960, int32(2))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v965 = int32(0)
	v1177 = int32(6)
	v1181 = v965
	v1185 = v965
	goto L29
L242:
	;
	v970 = v929
	goto L226
L243:
	;
	if base.B2i32(v133 == int32(0))|v159&int32(128)|base.B2i32(v159&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		goto L273
	} else {
		goto L274
	}
L244:
	;
	v1047 = int32(0)
	v1051 = F_Do_MultiXactIdWait(m, v160, v971, v159, v1047, v128, v142, int32(3), v1047, v1047)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L272
	}
L245:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	switch v132 {
	case 0:
		goto L244
	case 1:
		goto L248
	case 2:
		goto L249
	default:
		goto L243
	}
L246:
	;
	goto L247
L247:
	;
	switch v132 {
	case 0:
		goto L261
	case 1:
		goto L260
	case 2:
		goto L259
	default:
		goto L243
	}
L248:
	;
	v1002 = int32(0)
	v1006 = F_Do_MultiXactIdWait(m, v160, v971, v159, int32(1), v128, v1002, v1002, v1002, v1002)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L256
	}
L249:
	;
	v973 = int32(0)
	v977 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v978 = F_Do_MultiXactIdWait(m, v160, v971, v159, int32(1), v128, v973, v973, v973, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	if v978 != 0 {
		goto L243
	} else {
		goto L251
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v987 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v137+int32(16))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_11), int32(_a_F_heap_lock_tuple_12))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	if v1006 != 0 {
		goto L243
	} else {
		goto L257
	}
L257:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1008, int32(2))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1177 = int32(6)
	v1181 = v970
	v1185 = v775
	goto L29
L259:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v1026 = F_ConditionalXactLockTableWait(m, v160, v1025)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L266
	}
L260:
	;
	v1017 = F_ConditionalXactLockTableWait(m, v160, int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L263
	}
L261:
	;
	F_XactLockTableWait(m, v160, v128, v142, int32(3))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	goto L243
L263:
	;
	if v1017 != 0 {
		goto L243
	} else {
		goto L264
	}
L264:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1019, int32(2))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1177 = int32(6)
	v1181 = v970
	v1185 = v775
	goto L29
L266:
	;
	if v1026 != 0 {
		goto L243
	} else {
		goto L267
	}
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v1035 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v137)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_13), int32(_a_F_heap_lock_tuple_12))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	goto L243
L273:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1094, int32(2))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L286
	}
L274:
	;
	v1065 = v137 + int32(44)
	v1066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+2)))
	v1067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142))))
	v1068 = int32(16)
	v1071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1065)+2)))
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1065))))
	if v1066|v1067<<(uint(v1068)%32) == v1071|v1072<<(uint(v1068)%32) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if v1082 != 0 {
		goto L273
	} else {
		goto L281
	}
L276:
	;
	goto L275
L277:
	;
	v1078 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	v1079 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1065)+4)))
	if v1078 == v1079 {
		v1082 = int32(1)
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1082 = int32(0)
	goto L276
L280:
	;
	goto L279
L281:
	;
	v1083 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v1085 = F_heap_lock_updated_tuple(m, v128, v159, v160, v1065, v1083, v131)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	if v1085 == int32(0) {
		goto L273
	} else {
		goto L284
	}
L284:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1089, int32(2))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1177 = v1085
	v1181 = v970
	v1185 = v775
	goto L29
L286:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1099 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1098)+20)))
	if (v1099^v159)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1401 = v970
		v1405 = v775
		goto L24
	} else {
		goto L287
	}
L287:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	if v1103 != v160 {
		v1401 = v970
		v1405 = v775
		goto L24
	} else {
		goto L288
	}
L288:
	;
	if v782|v1099&int32(3072) != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1128 = int32(0)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1129)+20)))
	if v1130&int32(2048)|v1130&int32(128)|base.B2i32(v1130&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		v1177 = v1128
		v1181 = v970
		v1185 = v775
		goto L29
	} else {
		goto L297
	}
L290:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v1099&int32(128)|base.B2i32(v1099&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	F_HeapTupleSetHintBits(m, v1098, v1108, int32(2048), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L296
	}
L292:
	;
	v1116 = F_TransactionIdDidCommit(m, v160)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	if v1116 == int32(0) {
		goto L291
	} else {
		goto L294
	}
L294:
	;
	F_HeapTupleSetHintBits(m, v1098, v1108, int32(1024), v160)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	goto L289
L296:
	;
	goto L289
L297:
	;
	v1141 = F_HeapTupleHeaderIsOnlyLocked(m, v1129)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	if v1141 != 0 {
		v1177 = v1128
		v1181 = v970
		v1185 = v775
		goto L29
	} else {
		goto L299
	}
L299:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1147 = v1145 + int32(12)
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+2)))
	v1149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142))))
	v1150 = int32(16)
	v1153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1147)+2)))
	v1154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1147))))
	if v1148|v1149<<(uint(v1150)%32) == v1153|v1154<<(uint(v1150)%32) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	if v1164 != 0 {
		goto L306
	} else {
		goto L307
	}
L301:
	;
	goto L300
L302:
	;
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	v1161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1147)+4)))
	if v1160 == v1161 {
		v1164 = int32(1)
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1164 = int32(0)
	goto L301
L305:
	;
	goto L304
L306:
	;
	v1165 = int32(4)
	goto L308
L307:
	;
	v1165 = int32(3)
	goto L308
L308:
	;
	v1177 = v1165
	v1181 = v970
	v1185 = v775
	goto L29
L309:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v137)+60))
	if v1217 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1235)+20)))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1235)+4))
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L316
	}
L311:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+10)))
	if v1218&int32(4) == int32(0) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1223, int32(0))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_visibilitymap_pin(m, v128, v146, v137+int32(60))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_LockBuffer(m, v1231, int32(2))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1425 = v1205
	v1427 = v1207
	v1431 = v1211
	goto L23
L316:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1240)+18)))
	v1242 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_compute_new_xmax_infomask(m, v1237, v1236, v1241, v1242, v131, int32(0), v137+int32(56), v137+int32(54), v137+int32(52))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v1253 = int32(_a_F_heap_lock_tuple_14)
	v1255 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8])) = v1255 + int32(1)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1259)+20)))
	v1262 = v1260 & int32(_a_F_heap_lock_tuple_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v1259)+20)) = uint16(v1262)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1264)+18)))
	v1267 = v1265 & int32(_a_F_heap_lock_tuple_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v1264)+18)) = uint16(v1267)
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+54)))
	v1271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1269)+20)))
	v1272 = v1270 | v1271
	*(*uint16)(unsafe.Add(mBase, uint32(v1269)+20)) = uint16(v1272)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1274)+18)))
	v1276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+52)))
	v1277 = v1275 | v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v1274)+18)) = uint16(v1277)
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	if v1270&int32(128)|base.B2i32(v1270&int32(_a_F_heap_lock_tuple_4) == int32(64)) == int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+10)))
	if v1306&int32(4) != 0 {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v137)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1279)+4)) = v1289
	v1304 = v1289
	goto L319
L321:
	;
	goto L322
L322:
	;
	v1291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1279)+18)))
	v1293 = v1291 & int32(_a_F_heap_lock_tuple_17)
	*(*uint16)(unsafe.Add(mBase, uint32(v1279)+18)) = uint16(v1293)
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v137)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+4)) = v1296
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1298)+16)) = uint16(v1299)
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v1298)+12)) = v1301
	v1304 = v1296
	goto L319
L323:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v137)+60))
	v1311 = F_visibilitymap_clear(m, v146, v1309, int32(2))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L1
	} else {
		goto L326
	}
L324:
	;
	v1313 = int32(0)
	goto L325
L325:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_MarkBufferDirty(m, v1314)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L327
	}
L326:
	;
	v1313 = v1311
	goto L325
L327:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317)+118)))
	if v1318 != int32(112) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1380 = int32(_a_F_heap_lock_tuple_14)
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8])) = v1382 - int32(1)
	v1574 = v128
	v1577 = v131
	v1580 = v134
	v1582 = int32(0)
	v1583 = v137
	v1588 = v142
	v1589 = v1207
	goto L17
L329:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[9]))
	if v1322 <= int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
	if v1325 != 0 {
		goto L328
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L335
	}
L333:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v128)+40))
	if v1326 != 0 {
		goto L328
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_XLogRegisterBuffer(m, int32(0), v1330, int32(8))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v1334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+44)) = v1304
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+48)) = uint16(v1334)
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1337)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+51)) = uint8(v1313)
	v1344 = int32(1)
	v1346 = int32(8)
	v1348 = int32(4)
	v1363 = int32(base.Ui32(v1338)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v1270)>>(uint(v1344)%32))&v1346 | (int32(base.Ui32(v1270)>>(uint(v1348)%32))&v1348 | (int32(base.Ui32(v1270)>>(uint(int32(12))%32))&v1344 | int32(base.Ui32(v1270)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+50)) = uint8(v1363)
	F_XLogRegisterData(m, v137+int32(44), v1346)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v1372 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v144))) = base.I64_rotr(v1372, int64(32))
	goto L328
L339:
	;
	if v1439 != int32(1) {
		v139 = v1439
		v141 = v1425
		v143 = v1427
		v147 = v1431
		goto L20
	} else {
		goto L340
	}
L340:
	;
	v1574 = v128
	v1577 = v131
	v1580 = v134
	v1582 = int32(1)
	v1583 = v137
	v1588 = v142
	v1589 = v1427
	goto L17
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v1537
	v1555 = int32(2)
	if v1454 == v1555 {
		goto L354
	} else {
		goto L355
	}
L342:
	;
	v1537 = v1474
	goto L341
L343:
	;
	goto L344
L344:
	;
	v1480 = int32(0)
	v1484 = F_GetMultiXactIdMembers(m, v1474, v137+int32(44), v1480)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	if v1484 <= int32(0) {
		v1537 = v1480
		goto L341
	} else {
		goto L346
	}
L346:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v1497 = v1480
	goto L349
L347:
	;
	F_pfree(m, v1488)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L353
	}
L348:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1516)))
	v1526 = v1524
	goto L347
L349:
	;
	v1516 = v1488 + v1497<<(uint(int32(3))%32)
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1517) {
		goto L348
	} else {
		goto L351
	}
L350:
	;
	v1526 = int32(0)
	goto L347
L351:
	;
	v1521 = v1497 + int32(1)
	if v1521 != v1484 {
		v1497 = v1521
		goto L349
	} else {
		goto L352
	}
L352:
	;
	goto L350
L353:
	;
	v1537 = v1526
	goto L341
L354:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+8))
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1558)+20)))
	if v1561&int32(32) != 0 {
		goto L358
	} else {
		goto L359
	}
L355:
	;
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(-1)
	v1574 = v128
	v1577 = v131
	v1580 = v134
	v1582 = v1454
	v1583 = v137
	v1588 = v142
	v1589 = v1458
	goto L17
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v1570
	v1574 = v128
	v1577 = v131
	v1580 = v134
	v1582 = v1555
	v1583 = v137
	v1588 = v142
	v1589 = v1458
	goto L17
L358:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[10]))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1565+v1560<<(uint(int32(3))%32))+4))
	v1570 = v1569
	goto L360
L359:
	;
	v1570 = v1560
	goto L360
L360:
	;
	goto L357
L361:
	;
	v1603 = v1574
	v1606 = v1577
	v1611 = v1582
	v1612 = v1583
	v1617 = v1588
	v1618 = v1589
	goto L16
L362:
	;
	F_ReleaseBuffer(m, v1628)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	if v1618&int32(1) != 0 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	goto L364
L366:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1606*int32(12))+uint32(_c_F_heap_lock_tuple[11])))
	F_UnlockTuple(m, v1603, v1617, v1635)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	m.G0 = v1612 - int32(-64)
	return v1611
L369:
	;
	goto L368
}
func F_heap_page_items(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v74 int64
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int64
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = F_superuser(m)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v26 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v32 = int32(base.Ui32(v28)>>(uint(int32(2))%32)) - int32(4)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				if v34 == int32(0) {
					if base.Ui32(v32) <= base.Ui32(int32(23)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v428 = m.ExcPending
						if v428 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v431 = m.ExcPending
							if v431 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
								F_errmsg(m, int32(_a_F_heap_page_items_0), v19)
								mBase = m.M
								v435 = m.ExcPending
								if v435 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_page_items_1), int32(152), int32(_a_F_heap_page_items_2))
									mBase = m.M
									v440 = m.ExcPending
									if v440 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v39 = F_init_MultiFuncCall(m, l0)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(_a_F_heap_page_items_3)
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0]))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
							*(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0])) = v44
							v47 = F_palloc(m, int32(12))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v52 = F_get_call_result_type(m, l0, int32(0), v19+int32(32))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v444 = m.ExcPending
										if v444 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_heap_page_items_4), int32(0))
											mBase = m.M
											v448 = m.ExcPending
											if v448 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_heap_page_items_1), int32(161), int32(_a_F_heap_page_items_2))
												mBase = m.M
												v453 = m.ExcPending
												if v453 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
										v57 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v57)
										*(*int32)(unsafe.Add(mBase, uint32(v47))) = v56
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v22 + int32(4)
										v63 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
										*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v47
										if base.Ui64(int64(25)) <= base.Ui64(v63) {
											v74 = int64(base.Ui64(v63+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
										} else {
											v74 = int64(0)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v74
										*(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0])) = v42
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
										v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
										v86 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
										if base.Ui64(v85) < base.Ui64(v86) {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
											v90 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v19)+22)) = v90
											*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v90
											v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v89+v94<<(uint(int32(2))%32))+20))
											*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v94
											v101 = int32(base.Ui32(v98) >> (uint(int32(17)) % 32))
											*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v101
											v104 = v98 & int32(_a_F_heap_page_items_5)
											*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v104
											*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(base.Ui32(v98)>>(uint(int32(15))%32)) & int32(3)
											if base.B2i32(v104 != (v104+int32(7))&int32(_a_F_heap_page_items_6))|base.B2i32(base.Ui32(v98) < base.Ui32(int32(_a_F_heap_page_items_7)))|base.B2i32(v32 < v104+v101) == int32(0) {
												v124 = v104 + v89
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v125
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v127
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v124 + int32(12)
												*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v129
												v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+18)))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v134
												v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v136
												v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
												*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v138
												if base.B2i32(base.Ui32(v138) < base.Ui32(int32(23)))|base.B2i32(base.Ui32(v101) < base.Ui32(v138))|base.B2i32((v138+int32(7))&int32(504) != v138) == int32(0) {
													if v136&int32(1) != 0 {
														v159 = int32(base.Ui32(v134&int32(2047)+int32(7)) >> (uint(int32(3)) % 32))
														if base.Ui32(v159) <= base.Ui32(v138-int32(23)) {
															v164 = v124 + int32(23)
															v165 = int32(0)
															v167 = v159 << (uint(int32(3)) % 32)
															v170 = F_palloc(m, v167+int32(1))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																if v167 == int32(0) {
																} else {
																	if v167 != int32(1) {
																		v181 = v165
																		v194 = int32(0)
																		for {
																			v201 = v164 + int32(base.Ui32(v181)>>(uint(int32(3))%32))
																			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
																			if int32(base.Ui32(v202)>>(uint(v181&int32(6))%32))&int32(1) != 0 {
																				v208 = int32(49)
																			} else {
																				v208 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v181+v170))) = uint8(v208)
																			v210 = int32(1)
																			v211 = v181 | v210
																			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
																			if int32(base.Ui32(v215)>>(uint(v211&int32(7))%32))&v210 != 0 {
																				v221 = int32(49)
																			} else {
																				v221 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v170+v211))) = uint8(v221)
																			v223 = int32(2)
																			v224 = v181 + v223
																			v226 = v194 + v223
																			if v226 != v167&int32(4094) {
																				v181 = v224
																				v194 = v226
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v167&int32(1) == int32(0) {
																		} else {
																			v231 = v224
																			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(base.Ui32(v231)>>(uint(int32(3))%32))))))
																			if int32(base.Ui32(v252)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
																				v258 = int32(49)
																			} else {
																				v258 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v231+v170))) = uint8(v258)
																		}
																	} else {
																		v231 = v165
																		v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(base.Ui32(v231)>>(uint(int32(3))%32))))))
																		if int32(base.Ui32(v252)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
																			v258 = int32(49)
																		} else {
																			v258 = int32(48)
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v231+v170))) = uint8(v258)
																	}
																}
																v277 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v167+v170))) = uint8(v277)
																v279 = F_cstring_to_text(m, v170)
																mBase = m.M
																v280 = m.ExcPending
																if v280 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v279
																	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
																	v288 = v282
																	if v288&int32(8) != 0 {
																		v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																		v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
																		*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
																		v314 = v305
																	} else {
																		v311 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
																		v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																		v314 = v313
																	}
																	v315 = v101 - v314
																	v317 = v315 + int32(4)
																	v318 = F_palloc(m, v317)
																	mBase = m.M
																	v319 = m.ExcPending
																	if v319 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
																		v323 = int32(0)
																		if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
																			v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																			base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
																		} else {
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
																		v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
																		mBase = m.M
																		v366 = m.ExcPending
																		if v366 != 0 {
																			return int32(0)
																		} else {
																			v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
																			v368 = F_HeapTupleHeaderGetDatum(m, v367)
																			mBase = m.M
																			v369 = m.ExcPending
																			if v369 != 0 {
																				return int32(0)
																			} else {
																				v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
																				v371 = int32(1)
																				v372 = v370 + v371
																				*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
																				v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
																				*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
																				v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
																				v392 = v368
																				m.G0 = v19 + int32(96)
																				return v392
																			}
																		}
																	}
																}
															}
														} else {
															v283 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)) = uint8(v283)
															v288 = v136
															if v288&int32(8) != 0 {
																v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
																*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
																v314 = v305
															} else {
																v311 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
																v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																v314 = v313
															}
															v315 = v101 - v314
															v317 = v315 + int32(4)
															v318 = F_palloc(m, v317)
															mBase = m.M
															v319 = m.ExcPending
															if v319 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
																v323 = int32(0)
																if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
																	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																	base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
																} else {
																}
																*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
																v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
																v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
																mBase = m.M
																v366 = m.ExcPending
																if v366 != 0 {
																	return int32(0)
																} else {
																	v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
																	v368 = F_HeapTupleHeaderGetDatum(m, v367)
																	mBase = m.M
																	v369 = m.ExcPending
																	if v369 != 0 {
																		return int32(0)
																	} else {
																		v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
																		v371 = int32(1)
																		v372 = v370 + v371
																		*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
																		v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
																		*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
																		v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
																		v392 = v368
																		m.G0 = v19 + int32(96)
																		return v392
																	}
																}
															}
														}
													} else {
														v285 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)) = uint8(v285)
														v288 = v136
														if v288&int32(8) != 0 {
															v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
															v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
															*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
															v314 = v305
														} else {
															v311 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
															v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
															v314 = v313
														}
														v315 = v101 - v314
														v317 = v315 + int32(4)
														v318 = F_palloc(m, v317)
														mBase = m.M
														v319 = m.ExcPending
														if v319 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
															v323 = int32(0)
															if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
																v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
																base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
															v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
															v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
															mBase = m.M
															v366 = m.ExcPending
															if v366 != 0 {
																return int32(0)
															} else {
																v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
																v368 = F_HeapTupleHeaderGetDatum(m, v367)
																mBase = m.M
																v369 = m.ExcPending
																if v369 != 0 {
																	return int32(0)
																} else {
																	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
																	v371 = int32(1)
																	v372 = v370 + v371
																	*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
																	v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
																	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
																	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
																	v392 = v368
																	m.G0 = v19 + int32(96)
																	return v392
																}
															}
														}
													}
												} else {
													v336 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v19)+29)) = uint8(v336)
													v338 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v19)+27)) = uint16(v338)
													v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
													mBase = m.M
													v366 = m.ExcPending
													if v366 != 0 {
														return int32(0)
													} else {
														v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
														v368 = F_HeapTupleHeaderGetDatum(m, v367)
														mBase = m.M
														v369 = m.ExcPending
														if v369 != 0 {
															return int32(0)
														} else {
															v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
															v371 = int32(1)
															v372 = v370 + v371
															*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
															v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
															*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
															v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
															v392 = v368
															m.G0 = v19 + int32(96)
															return v392
														}
													}
												}
											} else {
												v340 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v19)+28)) = uint16(v340)
												*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = int64(72340172838076673)
												v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
												v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
												mBase = m.M
												v366 = m.ExcPending
												if v366 != 0 {
													return int32(0)
												} else {
													v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
													v368 = F_HeapTupleHeaderGetDatum(m, v367)
													mBase = m.M
													v369 = m.ExcPending
													if v369 != 0 {
														return int32(0)
													} else {
														v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
														v371 = int32(1)
														v372 = v370 + v371
														*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
														v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
														*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
														v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
														v392 = v368
														m.G0 = v19 + int32(96)
														return v392
													}
												}
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v382 = m.ExcPending
											if v382 != 0 {
												return int32(0)
											} else {
												v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = int32(2)
												v386 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v386)
												v392 = int32(0)
												m.G0 = v19 + int32(96)
												return v392
											}
										}
									}
								}
							}
						}
					}
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
					v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
					if base.Ui64(v85) < base.Ui64(v86) {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
						v90 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+22)) = v90
						*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v90
						v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89+v94<<(uint(int32(2))%32))+20))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v94
						v101 = int32(base.Ui32(v98) >> (uint(int32(17)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v101
						v104 = v98 & int32(_a_F_heap_page_items_5)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v104
						*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(base.Ui32(v98)>>(uint(int32(15))%32)) & int32(3)
						if base.B2i32(v104 != (v104+int32(7))&int32(_a_F_heap_page_items_6))|base.B2i32(base.Ui32(v98) < base.Ui32(int32(_a_F_heap_page_items_7)))|base.B2i32(v32 < v104+v101) == int32(0) {
							v124 = v104 + v89
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v125
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v127
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v124 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v129
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+18)))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v134
							v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v136
							v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v138
							if base.B2i32(base.Ui32(v138) < base.Ui32(int32(23)))|base.B2i32(base.Ui32(v101) < base.Ui32(v138))|base.B2i32((v138+int32(7))&int32(504) != v138) == int32(0) {
								if v136&int32(1) != 0 {
									v159 = int32(base.Ui32(v134&int32(2047)+int32(7)) >> (uint(int32(3)) % 32))
									if base.Ui32(v159) <= base.Ui32(v138-int32(23)) {
										v164 = v124 + int32(23)
										v165 = int32(0)
										v167 = v159 << (uint(int32(3)) % 32)
										v170 = F_palloc(m, v167+int32(1))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return int32(0)
										} else {
											if v167 == int32(0) {
											} else {
												if v167 != int32(1) {
													v181 = v165
													v194 = int32(0)
													for {
														v201 = v164 + int32(base.Ui32(v181)>>(uint(int32(3))%32))
														v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
														if int32(base.Ui32(v202)>>(uint(v181&int32(6))%32))&int32(1) != 0 {
															v208 = int32(49)
														} else {
															v208 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v181+v170))) = uint8(v208)
														v210 = int32(1)
														v211 = v181 | v210
														v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
														if int32(base.Ui32(v215)>>(uint(v211&int32(7))%32))&v210 != 0 {
															v221 = int32(49)
														} else {
															v221 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v170+v211))) = uint8(v221)
														v223 = int32(2)
														v224 = v181 + v223
														v226 = v194 + v223
														if v226 != v167&int32(4094) {
															v181 = v224
															v194 = v226
															continue
														} else {
															break
														}
														break
													}
													if v167&int32(1) == int32(0) {
													} else {
														v231 = v224
														v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(base.Ui32(v231)>>(uint(int32(3))%32))))))
														if int32(base.Ui32(v252)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
															v258 = int32(49)
														} else {
															v258 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v231+v170))) = uint8(v258)
													}
												} else {
													v231 = v165
													v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(base.Ui32(v231)>>(uint(int32(3))%32))))))
													if int32(base.Ui32(v252)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
														v258 = int32(49)
													} else {
														v258 = int32(48)
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v231+v170))) = uint8(v258)
												}
											}
											v277 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v167+v170))) = uint8(v277)
											v279 = F_cstring_to_text(m, v170)
											mBase = m.M
											v280 = m.ExcPending
											if v280 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v279
												v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
												v288 = v282
												if v288&int32(8) != 0 {
													v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
													v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
													*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
													v314 = v305
												} else {
													v311 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
													v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
													v314 = v313
												}
												v315 = v101 - v314
												v317 = v315 + int32(4)
												v318 = F_palloc(m, v317)
												mBase = m.M
												v319 = m.ExcPending
												if v319 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
													v323 = int32(0)
													if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
														v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
														base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
													v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
													mBase = m.M
													v366 = m.ExcPending
													if v366 != 0 {
														return int32(0)
													} else {
														v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
														v368 = F_HeapTupleHeaderGetDatum(m, v367)
														mBase = m.M
														v369 = m.ExcPending
														if v369 != 0 {
															return int32(0)
														} else {
															v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
															v371 = int32(1)
															v372 = v370 + v371
															*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
															v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
															*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
															v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
															v392 = v368
															m.G0 = v19 + int32(96)
															return v392
														}
													}
												}
											}
										}
									} else {
										v283 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)) = uint8(v283)
										v288 = v136
										if v288&int32(8) != 0 {
											v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
											v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
											*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
											v314 = v305
										} else {
											v311 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
											v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
											v314 = v313
										}
										v315 = v101 - v314
										v317 = v315 + int32(4)
										v318 = F_palloc(m, v317)
										mBase = m.M
										v319 = m.ExcPending
										if v319 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
											v323 = int32(0)
											if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
												v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
												base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
											v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
											v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
											mBase = m.M
											v366 = m.ExcPending
											if v366 != 0 {
												return int32(0)
											} else {
												v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
												v368 = F_HeapTupleHeaderGetDatum(m, v367)
												mBase = m.M
												v369 = m.ExcPending
												if v369 != 0 {
													return int32(0)
												} else {
													v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
													v371 = int32(1)
													v372 = v370 + v371
													*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
													v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
													*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
													v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
													v392 = v368
													m.G0 = v19 + int32(96)
													return v392
												}
											}
										}
									}
								} else {
									v285 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)) = uint8(v285)
									v288 = v136
									if v288&int32(8) != 0 {
										v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
										v309 = *(*int32)(unsafe.Add(mBase, uint32(v124+v305-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v309
										v314 = v305
									} else {
										v311 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)) = uint8(v311)
										v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
										v314 = v313
									}
									v315 = v101 - v314
									v317 = v315 + int32(4)
									v318 = F_palloc(m, v317)
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
										v323 = int32(0)
										if base.B2i32(v315 == v323)|base.B2i32(v315 <= v323) == v323 {
											v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
											base.MemoryCopy(m, v318+int32(4), v124+v332, v315)
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v318
										v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
										v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
										mBase = m.M
										v366 = m.ExcPending
										if v366 != 0 {
											return int32(0)
										} else {
											v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
											v368 = F_HeapTupleHeaderGetDatum(m, v367)
											mBase = m.M
											v369 = m.ExcPending
											if v369 != 0 {
												return int32(0)
											} else {
												v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
												v371 = int32(1)
												v372 = v370 + v371
												*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
												v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
												*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
												v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
												v392 = v368
												m.G0 = v19 + int32(96)
												return v392
											}
										}
									}
								}
							} else {
								v336 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+29)) = uint8(v336)
								v338 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v19)+27)) = uint16(v338)
								v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
								v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
								mBase = m.M
								v366 = m.ExcPending
								if v366 != 0 {
									return int32(0)
								} else {
									v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
									v368 = F_HeapTupleHeaderGetDatum(m, v367)
									mBase = m.M
									v369 = m.ExcPending
									if v369 != 0 {
										return int32(0)
									} else {
										v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
										v371 = int32(1)
										v372 = v370 + v371
										*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
										v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
										*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
										v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
										v392 = v368
										m.G0 = v19 + int32(96)
										return v392
									}
								}
							}
						} else {
							v340 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+28)) = uint16(v340)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = int64(72340172838076673)
							v360 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v365 = F_heap_form_tuple(m, v360, v19+int32(32), v19+int32(16))
							mBase = m.M
							v366 = m.ExcPending
							if v366 != 0 {
								return int32(0)
							} else {
								v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
								v368 = F_HeapTupleHeaderGetDatum(m, v367)
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return int32(0)
								} else {
									v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
									v371 = int32(1)
									v372 = v370 + v371
									*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v372)
									v374 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
									*(*int64)(unsafe.Add(mBase, uint32(v84))) = v374 + int64(1)
									v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v371
									v392 = v368
									m.G0 = v19 + int32(96)
									return v392
								}
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v382 = m.ExcPending
						if v382 != 0 {
							return int32(0)
						} else {
							v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = int32(2)
							v386 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v386)
							v392 = int32(0)
							m.G0 = v19 + int32(96)
							return v392
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v412 = m.ExcPending
				if v412 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v415 = m.ExcPending
					if v415 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_heap_page_items_8), int32(0))
						mBase = m.M
						v419 = m.ExcPending
						if v419 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_heap_page_items_1), int32(140), int32(_a_F_heap_page_items_2))
							mBase = m.M
							v424 = m.ExcPending
							if v424 != 0 {
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
}
func F_heap_page_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int64
	_ = v105
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v359 int32
	_ = v359
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v550 int32
	_ = v550
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v994 int32
	_ = v994
	var v1002 int32
	_ = v1002
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1079 int32
	_ = v1079
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1180 int32
	_ = v1180
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1208 int32
	_ = v1208
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int64
	_ = v1430
	var v1431 int64
	_ = v1431
	var v1432 int64
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1492 int64
	_ = v1492
	var v1493 int64
	_ = v1493
	var v1494 int64
	_ = v1494
	var v1500 int32
	_ = v1500
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1531 int32
	_ = v1531
	var v1539 int32
	_ = v1539
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1961 int32
	_ = v1961
	var v1967 int32
	_ = v1967
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1984 int32
	_ = v1984
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2054 int32
	_ = v2054
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2168 int32
	_ = v2168
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2238 int32
	_ = v2238
	var v2273 int32
	_ = v2273
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2325 int32
	_ = v2325
	var v2336 int32
	_ = v2336
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2489 int32
	_ = v2489
	var v2518 int32
	_ = v2518
	var v2525 int32
	_ = v2525
	var v2553 int32
	_ = v2553
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	v11 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(_a_F_heap_page_prune_and_freeze_0)
	m.G0 = v35
	if l1 < v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(l1^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L1
L3:
	;
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v54 = v48 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[2]))) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[3]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = l4
	v81 = int32(1)
	v85 = int32(base.Ui32(l3)>>(uint(v81)%32)) & v81
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)) = uint8(v85)
	v88 = l3 & v81
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l2
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v91
	if v85 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v73 = v64
	goto L5
L7:
	;
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v73 = v72
	goto L5
L9:
	;
	v129 = v35 + int32(44)
	v131 = *(*int64)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[8]))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[9]))) = uint8(v127)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[10]))) = uint8(v127)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v137
	if base.Ui32(v136) < base.Ui32(int32(25)) {
		v1180 = v11
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[11]))) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[12]))) = v97
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[14]))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v100
	v105 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v105
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[17]))) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[18]))) = l5 + int32(36)
	v127 = v81
	goto L9
L11:
	;
	goto L12
L12:
	;
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[12]))) = v112
	v114 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[11]))) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[17]))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[18]))) = l5 + int32(36)
	v127 = v114
	goto L9
L13:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[3])))
	v1192 = v1190 - int32(1)
	if int32(0) <= v1192 {
		goto L132
	} else {
		goto L133
	}
L14:
	;
	v144 = int32(base.Ui32(v136+int32(_a_F_heap_page_prune_and_freeze_1)) >> (uint(int32(2)) % 32))
	v146 = v144 & int32(_a_F_heap_page_prune_and_freeze_2)
	if v146 == int32(0) {
		v1180 = v11
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v150 = int32(base.Ui32(v73) >> (uint(int32(16)) % 32))
	v158 = v35 + int32(_a_F_heap_page_prune_and_freeze_3)
	v166 = v144
	v167 = v146
	goto L16
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v167)
	v196 = v167 + (v35 + int32(_a_F_heap_page_prune_and_freeze_4))
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v197)
	v199 = v167 + (v35 + int32(_a_F_heap_page_prune_and_freeze_5))
	v200 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v200)
	v204 = v54 + int32(20) + v167<<(uint(int32(2))%32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	switch int32(base.Ui32(v205)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L20
	case 2:
		goto L21
	default:
		goto L22
	}
L17:
	;
	v319 = *(*int64)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[7]))
	v320 = base.B2i32(v131 != v319)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	v323 = v321 - int32(1)
	if v323 < int32(0) {
		v1180 = v320
		goto L13
	} else {
		goto L44
	}
L18:
	;
	v312 = int32(1)
	v315 = v166 - v312
	if v315&int32(_a_F_heap_page_prune_and_freeze_2) != 0 {
		v166 = v315
		v167 = v167 - v312
		goto L16
	} else {
		goto L43
	}
L19:
	;
	v247 = v54 + v205&int32(_a_F_heap_page_prune_and_freeze_6)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v167)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+14)) = uint16(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v150)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(base.Ui32(v249) >> (uint(int32(17)) % 32))
	v260 = F_HeapTupleSatisfiesVacuumHorizon(m, v35+int32(8), l1, v35+int32(_a_F_heap_page_prune_and_freeze_7))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	v238 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4]))) = v237 + v238
	*(*uint16)(unsafe.Add(mBase, uint32(v158+v237<<(uint(v238)%32)))) = uint16(v167)
	goto L18
L21:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v215)
	if v214 == v215 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v212)
	goto L18
L23:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v220 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(1810)+v219<<(uint(v220)%32)))) = uint16(v167)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v224 + v220
	goto L18
L24:
	;
	goto L25
L25:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	v229 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v228 + v229
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[18])))
	*(*uint16)(unsafe.Add(mBase, uint32(v232+v228<<(uint(v229)%32)))) = uint16(v167)
	goto L18
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v285)
	v289 = int32(*(*int16)(unsafe.Add(mBase, uint32(v247)+18)))
	if int32(0) <= v289 {
		goto L40
	} else {
		goto L41
	}
L27:
	;
	return
L28:
	;
	if v260 != int32(2) {
		v285 = v260
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	if v264 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v282 = F_GlobalVisTestIsRemovableXid(m, v281, v277)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[19])))
	v277 = v267
	goto L30
L32:
	;
	goto L33
L33:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[19])))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	if v269 == int32(0) {
		v277 = v268
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v272 = int32(0)
	if v268-v269 < v272 {
		v285 = v272
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v277 = v268
	goto L30
L36:
	;
	if v282 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v284 = int32(0)
	goto L39
L38:
	;
	v284 = int32(2)
	goto L39
L39:
	;
	v285 = v284
	goto L26
L40:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	v293 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[4]))) = v292 + v293
	*(*uint16)(unsafe.Add(mBase, uint32(v158+v292<<(uint(v293)%32)))) = uint16(v167)
	goto L18
L41:
	;
	goto L42
L42:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[3])))
	v301 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[3]))) = v300 + v301
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_8)+v300<<(uint(v301)%32)))) = uint16(v167)
	goto L18
L43:
	;
	goto L17
L44:
	;
	v327 = v54 + int32(20)
	v329 = v35 + int32(1228)
	v331 = v35 + int32(1810)
	v333 = v35 - int32(-64)
	v337 = v35 + int32(_a_F_heap_page_prune_and_freeze_4)
	v359 = v323
	goto L45
L45:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_3)+v359<<(uint(int32(1))%32)))))
	v376 = v337 + v375
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v377 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v1180 = v320
	goto L13
L47:
	;
	if int32(0) < v359 {
		v359 = v359 - int32(1)
		goto L45
	} else {
		goto L131
	}
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v375)
	v381 = v327 + v375<<(uint(int32(2))%32)
	v382 = int32(0)
	v385 = int32(_a_F_heap_page_prune_and_freeze_2)
	v386 = v144 & v385
	if base.Ui32(v386) <= base.Ui32((v375-int32(1))&v385) {
		v520 = v382
		v522 = v382
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v592 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L50:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if base.B2i32(v550&int32(_a_F_heap_page_prune_and_freeze_9) != int32(_a_F_heap_page_prune_and_freeze_10))|base.B2i32(int32(1) < v522) != 0 {
		v592 = v520
		v594 = v522
		goto L49
	} else {
		goto L81
	}
L51:
	;
	v394 = v382
	v395 = v375
	v396 = v382
	v410 = v382
	goto L52
L52:
	;
	v425 = v395 & int32(_a_F_heap_page_prune_and_freeze_2)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+v425))))
	if v427 != 0 {
		v520 = v394
		v522 = v396
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v520 = v508
	v522 = v510
	goto L50
L54:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v327+v425<<(uint(int32(2))%32))))
	if v431&int32(_a_F_heap_page_prune_and_freeze_9) == int32(_a_F_heap_page_prune_and_freeze_10) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if base.Ui32((v509-int32(1))&int32(_a_F_heap_page_prune_and_freeze_2)) < base.Ui32(v386) {
		v394 = v508
		v395 = v509
		v396 = v510
		v410 = v512
		goto L52
	} else {
		goto L80
	}
L56:
	;
	if int32(0) < v396 {
		v520 = v394
		v522 = v396
		goto L50
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v451 = v54 + v431&int32(_a_F_heap_page_prune_and_freeze_6)
	if v410 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v440 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v396<<(uint(v440)%32)))) = uint16(v395)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	v508 = v394
	v509 = v446 & int32(_a_F_heap_page_prune_and_freeze_6)
	v510 = v396 + v440
	v512 = v410
	goto L55
L60:
	;
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451)+20)))
	v453 = int32(768)
	if v452&v453 != v453 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v463 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v396<<(uint(v463)%32)))) = uint16(v395)
	v468 = v396 + v463
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+(v35+int32(_a_F_heap_page_prune_and_freeze_5))))))
	switch v470 {
	case 0:
		goto L68
	case 1, 3, 4:
		v592 = v394
		v594 = v468
		goto L49
	case 2:
		v486 = v394
		goto L67
	default:
		goto L69
	}
L63:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v459 = v457
	goto L65
L64:
	;
	v459 = int32(2)
	goto L65
L65:
	;
	if v459 != v410 {
		v520 = v394
		v522 = v396
		goto L50
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+19)))
	if v487&int32(64) == int32(0) {
		v592 = v486
		v594 = v468
		goto L49
	} else {
		goto L74
	}
L68:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v451, v129)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L27
	} else {
		goto L73
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_11), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_12), int32(1117), int32(_a_F_heap_page_prune_and_freeze_13))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L27
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
	v486 = v468
	goto L67
L74:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451)+20)))
	if v492&int32(2048)|base.B2i32(v492&int32(768) == int32(512)) != 0 {
		v592 = v486
		v594 = v468
		goto L49
	} else {
		goto L75
	}
L75:
	;
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451)+16)))
	if v492&int32(_a_F_heap_page_prune_and_freeze_14) == int32(_a_F_heap_page_prune_and_freeze_15) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v505 = F_HeapTupleGetUpdateXid(m, v451)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L27
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v508 = v486
	v509 = v500
	v510 = v468
	v512 = v507
	goto L55
L79:
	;
	v508 = v486
	v509 = v500
	v510 = v468
	v512 = v505
	goto L55
L80:
	;
	goto L53
L81:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v559 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v559)
	if v558 == v559 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v564 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v563<<(uint(v564)%32)))) = uint16(v375)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v568 + v564
	goto L47
L83:
	;
	goto L84
L84:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v573 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v329+v572<<(uint(v573)%32)))) = uint16(v375)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v577 + v573
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v581 + v573
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[18])))
	*(*uint16)(unsafe.Add(mBase, uint32(v585+v581<<(uint(v573)%32)))) = uint16(v375)
	goto L47
L85:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v625&int32(_a_F_heap_page_prune_and_freeze_9) == int32(_a_F_heap_page_prune_and_freeze_10) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v592 == v594 {
		goto L96
	} else {
		goto L97
	}
L88:
	;
	v630 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v630)
	v633 = v630
	goto L90
L89:
	;
	v633 = int32(0)
	goto L90
L90:
	;
	if v594 <= v633 {
		goto L47
	} else {
		goto L91
	}
L91:
	;
	v638 = v633
	goto L92
L92:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v638<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L27
	} else {
		goto L94
	}
L93:
	;
	goto L47
L94:
	;
	v678 = v638 + int32(1)
	if v678 != v594 {
		v638 = v678
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v683 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v683)
	v686 = v681 & int32(_a_F_heap_page_prune_and_freeze_9)
	if v682 == v683 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	goto L98
L98:
	;
	v870 = int32(1)
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v592<<(uint(v870)%32)))))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v870)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v878 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v333+v877<<(uint(v878)%32)))) = uint16(v375)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v333+v882<<(uint(v878)%32))+2)) = uint16(v873)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v887 + v870
	if v874&int32(_a_F_heap_page_prune_and_freeze_9) == int32(_a_F_heap_page_prune_and_freeze_16) {
		goto L114
	} else {
		goto L115
	}
L99:
	;
	if v594 < int32(2) {
		goto L47
	} else {
		goto L106
	}
L100:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v721 + int32(1)
	goto L99
L101:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v690 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v689<<(uint(v690)%32)))) = uint16(v375)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v694 + v690
	if v686 == int32(_a_F_heap_page_prune_and_freeze_16) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v701 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v329+v700<<(uint(v701)%32)))) = uint16(v375)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v705 + v701
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v709 + v701
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[18])))
	*(*uint16)(unsafe.Add(mBase, uint32(v713+v709<<(uint(v701)%32)))) = uint16(v375)
	if v686 != int32(_a_F_heap_page_prune_and_freeze_16) {
		goto L99
	} else {
		goto L105
	}
L104:
	;
	goto L99
L105:
	;
	goto L100
L106:
	;
	v728 = int32(1)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v594 != int32(2) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v732 = int32(1)
	v733 = v594 - v732
	v741 = v729
	v742 = v728
	v743 = int32(0)
	goto L110
L108:
	;
	v817 = v729
	v818 = v728
	goto L109
L109:
	;
	v849 = int32(1)
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v818<<(uint(v849)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v852))) = uint8(v849)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v817<<(uint(v849)%32)))) = uint16(v852)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v860 + v849
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v864 + v849
	goto L47
L110:
	;
	v773 = int32(1)
	v775 = v35 + int32(_a_F_heap_page_prune_and_freeze_7) + v742<<(uint(v773)%32)
	v776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775))))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v776))) = uint8(v773)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v741<<(uint(v773)%32)))) = uint16(v776)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v784 + v773
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v790 = v788 + v773
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v790
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v792))) = uint8(v773)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v790<<(uint(v773)%32)))) = uint16(v792)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v802 = v800 + v773
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v802
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v804 + v773
	v808 = int32(2)
	v809 = v742 + v808
	v811 = v743 + v808
	if v811 != v733&int32(-2) {
		v741 = v802
		v742 = v809
		v743 = v811
		goto L110
	} else {
		goto L112
	}
L111:
	;
	if v733&v732 == int32(0) {
		goto L47
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v817 = v802
	v818 = v809
	goto L109
L114:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v895 + int32(1)
	goto L116
L115:
	;
	goto L116
L116:
	;
	v899 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[20]))) = uint8(v899)
	if v592 < int32(2) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v594 <= v592 {
		goto L47
	} else {
		goto L126
	}
L118:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v592 != int32(2) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v907 = int32(1)
	v908 = v592 - v907
	v918 = v907
	v926 = v904
	v929 = int32(0)
	goto L122
L120:
	;
	v994 = int32(1)
	v1002 = v904
	goto L121
L121:
	;
	v1025 = int32(1)
	v1028 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v994<<(uint(v1025)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v1028))) = uint8(v1025)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v1002<<(uint(v1025)%32)))) = uint16(v1028)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v1036 + v1025
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v1040 + v1025
	goto L117
L122:
	;
	v949 = int32(1)
	v951 = v35 + int32(_a_F_heap_page_prune_and_freeze_7) + v918<<(uint(v949)%32)
	v952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v951))))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v952))) = uint8(v949)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v926<<(uint(v949)%32)))) = uint16(v952)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v960 + v949
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v966 = v964 + v949
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v966
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v951)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v968))) = uint8(v949)
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v966<<(uint(v949)%32)))) = uint16(v968)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v978 = v976 + v949
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v978
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v980 + v949
	v984 = int32(2)
	v985 = v918 + v984
	v987 = v929 + v984
	if v987 != v908&int32(-2) {
		v918 = v985
		v926 = v978
		v929 = v987
		goto L122
	} else {
		goto L124
	}
L123:
	;
	if v908&v907 == int32(0) {
		goto L117
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v994 = v985
	v1002 = v978
	goto L121
L126:
	;
	v1079 = v592
	goto L127
L127:
	;
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_7)+v1079<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L27
	} else {
		goto L129
	}
L128:
	;
	goto L47
L129:
	;
	v1120 = v1079 + int32(1)
	if v1120 != v594 {
		v1079 = v1120
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L46
L132:
	;
	v1208 = v1192
	goto L135
L133:
	;
	goto L134
L134:
	;
	v1341 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1341)
	v1343 = int32(1)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	if v1341 < v1345 {
		v1355 = v1343
		goto L153
	} else {
		goto L154
	}
L135:
	;
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(_a_F_heap_page_prune_and_freeze_8)+v1208<<(uint(int32(1))%32)))))
	v1241 = v35 + int32(_a_F_heap_page_prune_and_freeze_4) + v1240
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241))))
	if v1242 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L134
L137:
	;
	if int32(0) < v1208 {
		v1208 = v1208 - int32(1)
		goto L135
	} else {
		goto L152
	}
L138:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1240)
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240+(v35+int32(_a_F_heap_page_prune_and_freeze_5))))))
	if v1245 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(20)+v1240<<(uint(int32(2))%32))))
	v1254 = v54 + v1251&int32(_a_F_heap_page_prune_and_freeze_6)
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+19)))
	if v1255&int32(64) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v1240)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L27
	} else {
		goto L151
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L27
	} else {
		goto L148
	}
L143:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1254, v129)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L27
	} else {
		goto L147
	}
L144:
	;
	v1260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1254)+20)))
	if v1260&int32(2048) != 0 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	if v1260&int32(768) != int32(512) {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v1270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1241))) = uint8(v1270)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(1810)+v1272<<(uint(v1270)%32)))) = uint16(v1240)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v1277 + v1270
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v1281 + v1270
	goto L137
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v1240
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v73
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_17), v35)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L27
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_12), int32(635), int32(_a_F_heap_page_prune_and_freeze_18))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L27
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	goto L137
L152:
	;
	goto L136
L153:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if v1356 == v1357 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	if int32(0) < v1349 {
		v1355 = int32(1)
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v1355 = base.B2i32(int32(0) < v1352)
	goto L153
L156:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)))
	v1364 = int32(base.Ui32(v1359&int32(2)) >> (uint(int32(1)) % 32))
	goto L158
L157:
	;
	v1364 = v1343
	goto L158
L158:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)))
	if v1365 != int32(1) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v1714 = int32(_a_F_heap_page_prune_and_freeze_19)
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[21])) = v1716 + int32(1)
	if v1364 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L160:
	;
	v1673 = int32(0)
	if v1672 <= v1673 {
		v1713 = v1673
		goto L159
	} else {
		goto L238
	}
L161:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1672 = v1671
	goto L160
L162:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[2]))))
	if v1368 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1510 = m.G0
	v1512 = v1510 - int32(32)
	m.G0 = v1512
	if l1 < int32(0) {
		goto L208
	} else {
		goto L209
	}
L164:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[10]))))
	if v1370 != int32(1) {
		v1672 = v1369
		goto L160
	} else {
		goto L165
	}
L165:
	;
	v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[9]))))
	v1376 = int32(0)
	if base.B2i32(v1373&int32(1) == v1376)|base.B2i32(v1369 <= v1376) != 0 {
		v1672 = v1369
		goto L160
	} else {
		goto L166
	}
L166:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+118)))
	if v1382 != int32(112) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[22]))
	if v1386 <= int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v1355 != 0 {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1389 != 0 {
		goto L161
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if v1180 != 0 {
		goto L163
	} else {
		goto L175
	}
L172:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1391 = int32(0)
	if v1180|base.B2i32(v1390 != v1391) == v1391 {
		goto L168
	} else {
		goto L173
	}
L173:
	;
	if v1390 != 0 {
		goto L161
	} else {
		goto L174
	}
L174:
	;
	goto L163
L175:
	;
	goto L168
L176:
	;
	v1399 = m.G0
	v1401 = v1399 - int32(16)
	m.G0 = v1401
	F_GetFullPageWriteInfo(m, v1401+int32(8), v1401+int32(7))
	mBase = m.M
	if l1 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L177:
	;
	goto L178
L178:
	;
	if v1364 == int32(0) {
		goto L161
	} else {
		goto L190
	}
L179:
	;
	if v1438 == int32(0) {
		goto L161
	} else {
		goto L189
	}
L180:
	;
	v1426 = int32(1)
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1401)+7)))
	if v1427 == v1426 {
		goto L185
	} else {
		goto L186
	}
L181:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1411+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1425 = v1417
	goto L180
L182:
	;
	goto L183
L183:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v1425 = v1419 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L180
L184:
	;
	m.G0 = v1401 + int32(16)
	goto L179
L185:
	;
	v1430 = *(*int64)(unsafe.Add(mBase, uint32(v1401)+8))
	v1431 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1425)+4)))
	v1432 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1425))))
	if base.Ui64(v1431|v1432<<(uint(int64(32))%64)) <= base.Ui64(v1430) {
		v1438 = v1426
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1438 = int32(0)
	goto L184
L188:
	;
	goto L187
L189:
	;
	goto L163
L190:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[23]))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+252))
	goto L191
L191:
	;
	if base.B2i32(v1448 != int32(0)) == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[24])))
	if v1454&int32(1) == int32(0) {
		goto L161
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1461 = m.G0
	v1463 = v1461 - int32(16)
	m.G0 = v1463
	F_GetFullPageWriteInfo(m, v1463+int32(8), v1463+int32(7))
	mBase = m.M
	if l1 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	goto L194
L196:
	;
	if v1500 == int32(0) {
		goto L161
	} else {
		goto L206
	}
L197:
	;
	v1488 = int32(1)
	v1489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463)+7)))
	if v1489 == v1488 {
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1473+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1487 = v1479
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v1487 = v1481 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L197
L201:
	;
	m.G0 = v1463 + int32(16)
	goto L196
L202:
	;
	v1492 = *(*int64)(unsafe.Add(mBase, uint32(v1463)+8))
	v1493 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1487)+4)))
	v1494 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1487))))
	if base.Ui64(v1493|v1494<<(uint(int64(32))%64)) <= base.Ui64(v1492) {
		v1500 = v1488
		goto L201
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1500 = int32(0)
	goto L201
L205:
	;
	goto L204
L206:
	;
	goto L163
L207:
	;
	if int32(0) < v1509 {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1517+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1531 = v1523
	goto L207
L209:
	;
	goto L210
L210:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v1531 = v1525 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L207
L211:
	;
	v1713 = int32(1)
	goto L159
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L27
	} else {
		goto L234
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L27
	} else {
		goto L230
	}
L214:
	;
	v1539 = int32(0)
	goto L217
L215:
	;
	goto L216
L216:
	;
	m.G0 = v1512 + int32(32)
	goto L211
L217:
	;
	v1571 = v35 + int32(2392) + v1539*int32(12)
	v1572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1571)+10)))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1531+int32(20)+v1572<<(uint(int32(2))%32))))
	v1579 = v1531 + v1576&int32(_a_F_heap_page_prune_and_freeze_6)
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1571)+9)))
	if v1580&int32(1) != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L216
L219:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1579)))
	v1584 = F_TransactionIdDidCommit(m, v1583)
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L27
	} else {
		goto L222
	}
L220:
	;
	v1590 = v1580
	goto L221
L221:
	;
	if v1590&int32(2) != 0 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v1584 == int32(0) {
		goto L213
	} else {
		goto L223
	}
L223:
	;
	v1588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1571)+9)))
	v1590 = v1588
	goto L221
L224:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+4))
	v1594 = F_TransactionIdDidCommit(m, v1593)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L27
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1598 = v1539 + int32(1)
	if v1598 != v1509 {
		v1539 = v1598
		goto L217
	} else {
		goto L229
	}
L227:
	;
	if v1594 != 0 {
		goto L212
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	goto L218
L230:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L27
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1512)+16)) = v1583
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_20), v1512+int32(16))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L27
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_21), int32(_a_F_heap_page_prune_and_freeze_22), int32(_a_F_heap_page_prune_and_freeze_23))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L27
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L27
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v1593
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_24), v1512)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L27
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_21), int32(_a_F_heap_page_prune_and_freeze_25), int32(_a_F_heap_page_prune_and_freeze_23))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L27
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	v1676 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+60)) = v1676
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[9]))) = uint8(v1676)
	v1713 = v1676
	goto L159
L239:
	;
	if v1355|v1713 != int32(1) {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v1722
	v1724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)))
	v1726 = v1724 & int32(_a_F_heap_page_prune_and_freeze_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)) = uint16(v1726)
	if (v1355|v1713)&int32(1) != 0 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L27
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	v2615 = int32(_a_F_heap_page_prune_and_freeze_19)
	v2617 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[21]))
	v2618 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[21])) = v2617 - v2618
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[16])))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v2621
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v2623
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v2625
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[25])))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v2627
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[26])))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v2629
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	v2632 = int32(0)
	v2634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[10]))))
	if base.B2i32(v2631 == v2632)&base.B2i32(v2634 == v2618) == v2632 {
		goto L330
	} else {
		goto L331
	}
L244:
	;
	if v1355 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	if l1 < int32(0) {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	goto L247
L247:
	;
	if v1713 != 0 {
		goto L286
	} else {
		goto L287
	}
L248:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	if v1737 <= int32(0) {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1741+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1755 = v1747
	goto L248
L250:
	;
	goto L251
L251:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v1755 = v1749 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L248
L252:
	;
	if v1757 <= int32(0) {
		goto L261
	} else {
		goto L262
	}
L253:
	;
	v1761 = v35 - int32(-64)
	v1763 = v1755 + int32(20)
	if v1737 != int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1773 = int32(0)
	v1775 = v1761
	goto L257
L255:
	;
	v1834 = v1761
	goto L256
L256:
	;
	v1862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1834))))
	v1866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1834)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1763+v1862<<(uint(int32(2))%32)))) = v1866&int32(_a_F_heap_page_prune_and_freeze_6) | int32(_a_F_heap_page_prune_and_freeze_10)
	goto L252
L257:
	;
	v1803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1775))))
	v1804 = int32(2)
	v1807 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1775)+2)))
	v1808 = int32(_a_F_heap_page_prune_and_freeze_6)
	v1810 = int32(_a_F_heap_page_prune_and_freeze_10)
	*(*int32)(unsafe.Add(mBase, uint32(v1763+v1803<<(uint(v1804)%32)))) = v1807&v1808 | v1810
	v1813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1775)+4)))
	v1817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1775)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1763+v1813<<(uint(v1804)%32)))) = v1817&v1808 | v1810
	v1824 = v1775 + int32(8)
	v1826 = v1773 + v1804
	if v1826 != v1737&int32(2147483646) {
		v1773 = v1826
		v1775 = v1824
		goto L257
	} else {
		goto L259
	}
L258:
	;
	if v1737&int32(1) == int32(0) {
		goto L252
	} else {
		goto L260
	}
L259:
	;
	goto L258
L260:
	;
	v1834 = v1824
	goto L256
L261:
	;
	if v1756 <= int32(0) {
		goto L273
	} else {
		goto L274
	}
L262:
	;
	v1907 = v35 + int32(1228)
	v1909 = v1757 & int32(3)
	v1911 = v1755 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v1757) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1919 = int32(0)
	v1921 = v1907
	goto L266
L264:
	;
	v1984 = v1907
	goto L265
L265:
	;
	v2015 = int32(0)
	v2017 = v1984
	goto L270
L266:
	;
	v1949 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1921))))
	v1950 = int32(2)
	v1953 = int32(_a_F_heap_page_prune_and_freeze_9)
	*(*int32)(unsafe.Add(mBase, uint32(v1911+v1949<<(uint(v1950)%32)))) = v1953
	v1955 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1921)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1911+v1955<<(uint(v1950)%32)))) = v1953
	v1961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1921)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1911+v1961<<(uint(v1950)%32)))) = v1953
	v1967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1921)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1911+v1967<<(uint(v1950)%32)))) = v1953
	v1974 = v1921 + int32(8)
	v1976 = v1919 + int32(4)
	if v1976 != v1757&int32(2147483644) {
		v1919 = v1976
		v1921 = v1974
		goto L266
	} else {
		goto L268
	}
L267:
	;
	if v1909 == int32(0) {
		goto L261
	} else {
		goto L269
	}
L268:
	;
	goto L267
L269:
	;
	v1984 = v1974
	goto L265
L270:
	;
	v2045 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2017))))
	v2046 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1911+v2045<<(uint(v2046)%32)))) = int32(_a_F_heap_page_prune_and_freeze_9)
	v2054 = v2015 + int32(1)
	if v2054 != v1909 {
		v2015 = v2054
		v2017 = v2017 + v2046
		goto L270
	} else {
		goto L272
	}
L271:
	;
	goto L261
L272:
	;
	goto L271
L273:
	;
	F_PageRepairFragmentation(m, v1755)
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L27
	} else {
		goto L285
	}
L274:
	;
	v2091 = v35 + int32(1810)
	v2093 = v1756 & int32(3)
	v2095 = v1755 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v1756) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v2103 = int32(0)
	v2105 = v2091
	goto L278
L276:
	;
	v2168 = v2091
	goto L277
L277:
	;
	v2199 = int32(0)
	v2201 = v2168
	goto L282
L278:
	;
	v2133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2105))))
	v2134 = int32(2)
	v2137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2095+v2133<<(uint(v2134)%32)))) = v2137
	v2139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2105)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v2095+v2139<<(uint(v2134)%32)))) = v2137
	v2145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2105)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2095+v2145<<(uint(v2134)%32)))) = v2137
	v2151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2105)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v2095+v2151<<(uint(v2134)%32)))) = v2137
	v2158 = v2105 + int32(8)
	v2160 = v2103 + int32(4)
	if v2160 != v1756&int32(2147483644) {
		v2103 = v2160
		v2105 = v2158
		goto L278
	} else {
		goto L280
	}
L279:
	;
	if v2093 == int32(0) {
		goto L273
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	v2168 = v2158
	goto L277
L282:
	;
	v2229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2201))))
	v2230 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2095+v2229<<(uint(v2230)%32)))) = int32(0)
	v2238 = v2199 + int32(1)
	if v2238 != v2093 {
		v2199 = v2238
		v2201 = v2201 + v2230
		goto L282
	} else {
		goto L284
	}
L283:
	;
	goto L273
L284:
	;
	goto L283
L285:
	;
	goto L247
L286:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v2307 = int32(0)
	if l1 < v2307 {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	goto L288
L288:
	;
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L27
	} else {
		goto L305
	}
L289:
	;
	if int32(0) < v2306 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	v2311 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2311+(l1^int32(-1))<<(uint(int32(2))%32))))
	v2325 = v2317
	goto L289
L291:
	;
	goto L292
L292:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v2325 = v2319 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L289
L293:
	;
	v2336 = v2307
	goto L296
L294:
	;
	goto L295
L295:
	;
	goto L288
L296:
	;
	v2366 = v35 + int32(2392) + v2336*int32(12)
	v2367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2366)+10)))
	v2368 = int32(2)
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2325+int32(20)+v2367<<(uint(v2368)%32))))
	v2374 = v2325 + v2371&int32(_a_F_heap_page_prune_and_freeze_6)
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2366)))
	*(*int32)(unsafe.Add(mBase, uint32(v2374)+4)) = v2375
	v2377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2366)+8)))
	if v2377&v2368 != 0 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	goto L295
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2374)+8)) = int32(2)
	v2382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2366)+8)))
	v2383 = v2382
	goto L300
L299:
	;
	v2383 = v2377
	goto L300
L300:
	;
	if v2383&int32(4) != 0 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2374)+8)) = int32(0)
	goto L303
L302:
	;
	goto L303
L303:
	;
	v2388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2366)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2374)+20)) = uint16(v2388)
	v2390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2366)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2374)+18)) = uint16(v2390)
	v2393 = v2336 + int32(1)
	if v2393 != v2306 {
		v2336 = v2393
		goto L296
	} else {
		goto L304
	}
L304:
	;
	goto L297
L305:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2461)+118)))
	if v2462 != int32(112) {
		goto L243
	} else {
		goto L306
	}
L306:
	;
	v2465 = int32(0)
	v2467 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[22]))
	if v2467 <= v2465 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2470 != 0 {
		goto L243
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	if v1713 == int32(0) {
		v2525 = v2465
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2471 != 0 {
		goto L243
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2553))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2525)) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L313:
	;
	v2474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[10]))))
	if v2474 != int32(1) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2483)+8))
	v2489 = v2484
	goto L317
L315:
	;
	v2477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[9]))))
	if v2477&int32(1) == int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[8])))
	v2525 = v2482
	goto L312
L317:
	;
	v2518 = v2489 - int32(1)
	if base.Ui32(v2518) < base.Ui32(int32(3)) {
		v2489 = v2518
		goto L317
	} else {
		goto L319
	}
L318:
	;
	v2525 = v2518
	goto L312
L319:
	;
	goto L318
L320:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	if v2565 != 0 {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	v2565 = base.B2i32(base.Ui32(v2553) < base.Ui32(v2525))
	goto L320
L322:
	;
	goto L323
L323:
	;
	v2565 = base.B2i32(int32(0) < v2525-v2553)
	goto L320
L324:
	;
	v2567 = v2525
	goto L326
L325:
	;
	v2567 = v2566
	goto L326
L326:
	;
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	F_log_heap_prune_and_freeze(m, l0, l1, v2567, int32(1), l6, v35+int32(2392), v2571, v35-int32(-64), v2574, v35+int32(1228), v2577, v35+int32(1810), v2580)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L27
	} else {
		goto L327
	}
L327:
	;
	goto L243
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v2631
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v2654
	v2657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)))
	if v2657 == int32(1) {
		goto L334
	} else {
		goto L335
	}
L329:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[8])))
	v2654 = v2652
	goto L328
L330:
	;
	v2640 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+20)) = uint16(v2640)
	v2642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v2642)
	goto L329
L331:
	;
	goto L332
L332:
	;
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[9]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)) = uint8(v2644)
	v2646 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+20)) = uint8(v2646)
	v2648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v2648)
	if v2644 != 0 {
		v2654 = int32(0)
		goto L328
	} else {
		goto L333
	}
L333:
	;
	goto L329
L334:
	;
	v2661 = v35 + int32(28)
	v2665 = base.B2i32(int32(0) < v2625)
	if int32(0) < v2625 {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	goto L336
L336:
	;
	m.G0 = v35 + int32(_a_F_heap_page_prune_and_freeze_0)
	return
L337:
	;
	v2666 = int32(_a_F_heap_page_prune_and_freeze_27)
	goto L339
L338:
	;
	v2666 = int32(_a_F_heap_page_prune_and_freeze_28)
	goto L339
L339:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2661+v2666)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v2668
	if int32(0) < v2625 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2672 = int32(_a_F_heap_page_prune_and_freeze_29)
	goto L342
L341:
	;
	v2672 = int32(_a_F_heap_page_prune_and_freeze_30)
	goto L342
L342:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2672+v2661)))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v2674
	goto L336
}
func F_heap_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int64
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int64
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int64
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int64
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v845 int64
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int64
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(_a_F_heap_redo_0)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L8
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	case 6:
		goto L9
	default:
		goto L15
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L16
	} else {
		goto L265
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L16
	} else {
		goto L262
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L16
	} else {
		goto L259
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L16
	} else {
		goto L256
	}
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L16
	} else {
		goto L253
	}
L6:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L16
	} else {
		goto L250
	}
L7:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L16
	} else {
		goto L247
	}
L8:
	;
	m.G0 = v17 + int32(_a_F_heap_redo_0)
	return
L9:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v863 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v867 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L16
	} else {
		goto L213
	}
L10:
	;
	v666 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+7)))
	if v668&int32(1) != 0 {
		goto L175
	} else {
		goto L176
	}
L11:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v575 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v579 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L16
	} else {
		goto L155
	}
L12:
	;
	F_heap_xlog_update(m, l0, int32(1))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L16
	} else {
		goto L154
	}
L13:
	;
	F_heap_xlog_update(m, l0, int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L16
	} else {
		goto L153
	}
L14:
	;
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v362 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v362, v17+int32(92), v362, v17+int32(_a_F_heap_redo_1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L102
	}
L15:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v29 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v29, v17+int32(80), v29, v17+int32(76))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	if v39&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v44
	v48 = F_CreateFakeRelcacheEntry(m, v17+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+48)))
	if v69 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	F_visibilitymap_pin(m, v48, v52, v17+int32(92))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v60 = F_visibilitymap_clear(m, v57, v58, int32(3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_ReleaseBuffer(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v48)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	if v343 != 0 {
		goto L96
	} else {
		goto L97
	}
L27:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	if v150 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L28:
	;
	v73 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v147 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(_a_F_heap_redo_1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L16
	} else {
		goto L47
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0]))) = v73
	if v73 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v94 = int32(_a_F_heap_redo_2)
	v95 = int32(0)
	if v95|(v93&int32(3)|int32(1)) == v95 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v73^int32(-1))<<(uint(int32(2))%32))))
	v93 = v85
	goto L32
L34:
	;
	goto L35
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v93 = v87 + v73<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v150 = v73
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+10)) = int32(_a_F_heap_redo_3)
	v135 = int32(_a_F_heap_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+18)) = uint16(v135)
	v141 = int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+16)) = uint16(v141)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+14)) = uint16(v141)
	goto L36
L38:
	;
	goto L41
L39:
	;
	goto L40
L40:
	;
	goto L46
L41:
	;
	v112 = v93 + v94
	v114 = v93 + int32(4)
	if base.Ui32(v114) < base.Ui32(v112) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = v112
	goto L44
L43:
	;
	v116 = v114
	goto L44
L44:
	;
	v121 = (v93^int32(-1)+v116)&int32(-4) + int32(4)
	if v121 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	base.MemoryFill(m, v93, int32(0), v121)
	goto L37
L46:
	;
	base.MemoryFill(m, v93, int32(0), v94)
	goto L37
L47:
	;
	if v147 != 0 {
		v335 = v2
		v337 = v2
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v150 = v149
	goto L27
L49:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+12)))
	if base.Ui32(v171) < base.Ui32(int32(25)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156+(v150^int32(-1))<<(uint(int32(2))%32))))
	v170 = v162
	goto L49
L51:
	;
	goto L52
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v170 = v164 + v150<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v180 = int32(1)
	goto L55
L54:
	;
	v180 = int32(base.Ui32(v171+int32(_a_F_heap_redo_5))>>(uint(int32(2))%32)) + int32(1)
	goto L55
L55:
	;
	if base.Ui32(v180&int32(_a_F_heap_redo_6)) < base.Ui32(v151) {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v185 = int32(base.Ui32(v37) >> (uint(int32(16)) % 32))
	v186 = int32(0)
	v188 = v17 + int32(72)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
	if v191 < v186 {
		v213 = v186
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(0)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v224 = v222 - int32(5)
	if v224 != 0 {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	v216 = v213
	goto L57
L59:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+int32(0))+76)))
	if v196 != int32(1) {
		v213 = v186
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v200 = v190 + int32(76)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+43)))
	if v201 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v188 == int32(0) {
		v213 = v186
		goto L58
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v188 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v206
	v216 = v206
	goto L57
L65:
	;
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v209
	goto L67
L66:
	;
	goto L67
L67:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+44))
	v213 = v211
	goto L58
L68:
	;
	base.MemoryCopy(m, v17+int32(115), v216+int32(5), v224)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+114)) = uint8(v219)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+110)) = uint16(v218)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v235 = v217 & int32(_a_F_heap_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+112)) = uint16(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v233
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+108)) = uint16(v38)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+106)) = uint16(v37)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+104)) = uint16(v185)
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v249 = F_PageAddItemExtended(m, v170, v17+int32(92), v222+int32(18), v247, int32(3))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	if v249 == int32(0) {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v256 = int32(4)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+14)))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+12)))
	v259 = v257 - v258
	if v259 <= v256 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = base.I64_rotr(v27, int64(32))
	v323 = int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	if v324&v323 != 0 {
		goto L92
	} else {
		goto L93
	}
L74:
	;
	v262 = v256
	goto L76
L75:
	;
	v262 = v259
	goto L76
L76:
	;
	v264 = v262 - int32(4)
	if v264 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v319 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v258) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v319 = v264
	goto L73
L81:
	;
	v275 = int32(base.Ui32(v258+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L83
L82:
	;
	v275 = int32(0)
	goto L83
L83:
	;
	if base.Ui32(v275&int32(_a_F_heap_redo_6)) < base.Ui32(int32(291)) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+10)))
	if v280&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v319 = int32(0)
	goto L73
L86:
	;
	goto L87
L87:
	;
	v289 = int32(1)
	goto L88
L88:
	;
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170+int32(20)+v289&int32(_a_F_heap_redo_6)<<(uint(int32(2))%32))+1)))
	if v298&int32(384) == int32(0) {
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v319 = int32(0)
	goto L73
L90:
	;
	v304 = v289 + int32(1)
	v305 = int32(_a_F_heap_redo_6)
	if base.Ui32(v304&v305) <= base.Ui32(v275&v305) {
		v289 = v304
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+10)))
	v329 = v327 & int32(_a_F_heap_redo_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v170)+10)) = uint16(v329)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	F_MarkBufferDirty(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	v335 = v323
	v337 = v319
	goto L26
L96:
	;
	F_UnlockReleaseBuffer(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L16
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v335&base.B2i32(base.Ui32(v337) < base.Ui32(int32(1638))) == int32(0) {
		goto L8
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v353
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	F_XLogRecordPageWithFreeSpace(m, v17+int32(8), v357, v337)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	goto L8
L102:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v361)+4)))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	if v372&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v377
	v381 = F_CreateFakeRelcacheEntry(m, v17+int32(40))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v404 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(80))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L16
	} else {
		goto L111
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = int32(0)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	F_visibilitymap_pin(m, v381, v385, v17+int32(80))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	v393 = F_visibilitymap_clear(m, v390, v391, int32(3))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	F_ReleaseBuffer(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	F_pfree(m, v381)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L16
	} else {
		goto L110
	}
L110:
	;
	goto L105
L111:
	;
	if v404 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v361)+4)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v409 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L114
L114:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v563 == int32(0) {
		goto L8
	} else {
		goto L151
	}
L115:
	;
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v428) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v413+(v409^int32(-1))<<(uint(int32(2))%32))))
	v427 = v419
	goto L115
L117:
	;
	goto L118
L118:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v427 = v421 + v409<<(uint(int32(13))%32) + int32(-8192)
	goto L115
L119:
	;
	v436 = int32(base.Ui32(v428+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L121
L120:
	;
	v436 = int32(0)
	goto L121
L121:
	;
	if base.Ui32(v436&int32(_a_F_heap_redo_6)) < base.Ui32(v408) {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v441 = v427 + int32(20)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441+v408<<(uint(int32(2))%32))))
	if v445&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v452 = v427 + v445&int32(_a_F_heap_redo_11)
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)))
	v455 = v453 & int32(_a_F_heap_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)) = uint16(v455)
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)))
	v459 = v457 & int32(-24577)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)) = uint16(v459)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)) = uint16(v459)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)) = uint16(v455)
	v464 = int32(1)
	v483 = v461<<(uint(v464)%32)&int32(16) | (v461<<(uint(int32(4))%32)&int32(64) | (v461<<(uint(int32(6))%32)&int32(128) | v461&v464<<(uint(int32(12))%32))) | v455
	if v461&int32(15) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)) = uint16(v483)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v461&int32(16) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v490 = v459 | int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)) = uint16(v490)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	if v492&int32(8) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452)+8)) = int32(0)
	v504 = v483 & int32(_a_F_heap_redo_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)) = uint16(v504)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v506 != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v452)+4)) = v497
	goto L130
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = int32(0)
	goto L130
L134:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	if v526&int32(1) != 0 {
		goto L144
	} else {
		goto L145
	}
L135:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v506))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v508)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = v524
	goto L134
L138:
	;
	if v520 == int32(0) {
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v520 = base.B2i32(base.Ui32(v508) < base.Ui32(v506))
	goto L138
L140:
	;
	goto L141
L141:
	;
	v520 = int32(base.Ui32(v508-v506) >> (uint(int32(31)) % 32))
	goto L138
L142:
	;
	goto L137
L143:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+14)) = uint16(v546)
	*(*uint32)(unsafe.Add(mBase, uint32(v427)+4)) = uint32(v360)
	v550 = int64(base.Ui64(v360) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v427))) = uint32(v550)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	F_MarkBufferDirty(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L16
	} else {
		goto L150
	}
L144:
	;
	v529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+10)))
	v531 = v529 & int32(_a_F_heap_redo_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v427)+10)) = uint16(v531)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	v534 = v533
	goto L146
L145:
	;
	v534 = v526
	goto L146
L146:
	;
	if v534&int32(16) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v537 = int32(_a_F_heap_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+16)) = uint16(v537)
	v539 = int32(_a_F_heap_redo_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+12)) = uint16(v539)
	v546 = v539
	goto L143
L148:
	;
	goto L149
L149:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+16)) = uint16(v371)
	v544 = int32(base.Ui32(v370) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+12)) = uint16(v544)
	v546 = v370
	goto L143
L150:
	;
	goto L114
L151:
	;
	F_UnlockReleaseBuffer(m, v563)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L16
	} else {
		goto L152
	}
L152:
	;
	goto L8
L153:
	;
	goto L8
L154:
	;
	goto L8
L155:
	;
	if v579 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v584 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v661 == int32(0) {
		goto L8
	} else {
		goto L173
	}
L159:
	;
	v603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v603) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v588+(v584^int32(-1))<<(uint(int32(2))%32))))
	v602 = v594
	goto L159
L161:
	;
	goto L162
L162:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v602 = v596 + v584<<(uint(int32(13))%32) + int32(-8192)
	goto L159
L163:
	;
	v611 = int32(base.Ui32(v603+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L165
L164:
	;
	v611 = int32(0)
	goto L165
L165:
	;
	if base.Ui32(v611&int32(_a_F_heap_redo_6)) < base.Ui32(v583) {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v602+v583<<(uint(int32(2))%32))+20))
	if v618&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	if v584 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v644 = v602 + v618&int32(_a_F_heap_redo_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v644)+16)) = uint16(v583)
	*(*uint16)(unsafe.Add(mBase, uint32(v644)+14)) = uint16(v641)
	v648 = int32(base.Ui32(v641) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v644)+12)) = uint16(v648)
	*(*uint32)(unsafe.Add(mBase, uint32(v602)+4)) = uint32(v575)
	v652 = int64(base.Ui64(v575) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v602))) = uint32(v652)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L16
	} else {
		goto L172
	}
L169:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[3]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v626+(v584^int32(-1))<<(uint(int32(6))%32))+16))
	v641 = v632
	goto L168
L170:
	;
	goto L171
L171:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[4]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v634+v584<<(uint(int32(6))%32)+int32(-64))+16))
	v641 = v640
	goto L168
L172:
	;
	goto L158
L173:
	;
	F_UnlockReleaseBuffer(m, v661)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	goto L8
L175:
	;
	v671 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v671
	F_XLogRecGetBlockTag(m, l0, v671, v17+int32(92), v671, v17+int32(_a_F_heap_redo_1))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L16
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v708 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L16
	} else {
		goto L184
	}
L178:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v681
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v683
	v687 = F_CreateFakeRelcacheEntry(m, v17+int32(56))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L16
	} else {
		goto L179
	}
L179:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	F_visibilitymap_pin(m, v687, v689, v17+int32(80))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L16
	} else {
		goto L180
	}
L180:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	v697 = F_visibilitymap_clear(m, v694, v695, int32(2))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L16
	} else {
		goto L181
	}
L181:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	F_ReleaseBuffer(m, v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L16
	} else {
		goto L182
	}
L182:
	;
	F_pfree(m, v687)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L16
	} else {
		goto L183
	}
L183:
	;
	goto L177
L184:
	;
	if v708 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v667)+4)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v713 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v857 == int32(0) {
		goto L8
	} else {
		goto L211
	}
L188:
	;
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v732) {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v717+(v713^int32(-1))<<(uint(int32(2))%32))))
	v731 = v723
	goto L188
L190:
	;
	goto L191
L191:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v731 = v725 + v713<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L192:
	;
	v740 = int32(base.Ui32(v732+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L194
L193:
	;
	v740 = int32(0)
	goto L194
L194:
	;
	if base.Ui32(v740&int32(_a_F_heap_redo_6)) < base.Ui32(v712) {
		goto L3
	} else {
		goto L195
	}
L195:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v731+v712<<(uint(int32(2))%32))+20))
	if v747&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L3
	} else {
		goto L196
	}
L196:
	;
	v754 = v731 + v747&int32(_a_F_heap_redo_11)
	v755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)))
	v757 = v755 & int32(_a_F_heap_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)) = uint16(v757)
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v754)+18)))
	v761 = v759 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+18)) = uint16(v761)
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+18)) = uint16(v761)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)) = uint16(v757)
	v766 = int32(1)
	v785 = v763<<(uint(v766)%32)&int32(16) | (v763<<(uint(int32(4))%32)&int32(64) | (v763<<(uint(int32(6))%32)&int32(128) | v763&v766<<(uint(int32(12))%32))) | v757
	if v763&int32(15) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)) = uint16(v785)
	goto L199
L198:
	;
	goto L199
L199:
	;
	if v763&int32(16) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v792 = v759 | int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+18)) = uint16(v792)
	v794 = v792
	goto L202
L201:
	;
	v794 = v761
	goto L202
L202:
	;
	v797 = int32(0)
	if base.B2i32(v785&int32(128) == v797)&base.B2i32(v785&int32(_a_F_heap_redo_15) != int32(64)) == v797 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v807 = v794 & int32(_a_F_heap_redo_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+18)) = uint16(v807)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v809 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v835 = v785
	goto L205
L205:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v838 = v835 & int32(_a_F_heap_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)) = uint16(v838)
	*(*int32)(unsafe.Add(mBase, uint32(v754)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v754)+4)) = v836
	*(*uint32)(unsafe.Add(mBase, uint32(v731)+4)) = uint32(v666)
	v845 = int64(base.Ui64(v666) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v731))) = uint32(v845)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L16
	} else {
		goto L210
	}
L206:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+16)) = uint16(v712)
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+14)) = uint16(v828)
	v832 = int32(base.Ui32(v828) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v754)+12)) = uint16(v832)
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v754)+20)))
	v835 = v834
	goto L205
L207:
	;
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[3]))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v813+(v809^int32(-1))<<(uint(int32(6))%32))+16))
	v828 = v819
	goto L206
L208:
	;
	goto L209
L209:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[4]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v821+v809<<(uint(int32(6))%32)+int32(-64))+16))
	v828 = v827
	goto L206
L210:
	;
	goto L187
L211:
	;
	F_UnlockReleaseBuffer(m, v857)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L16
	} else {
		goto L212
	}
L212:
	;
	goto L8
L213:
	;
	if v867 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v871 = int32(0)
	v873 = v17 + int32(80)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)+72))
	if v876 < v871 {
		v898 = v871
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v964 != 0 {
		goto L242
	} else {
		goto L243
	}
L217:
	;
	v902 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v862))))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v903 < int32(0) {
		goto L229
	} else {
		goto L230
	}
L218:
	;
	v901 = v898
	goto L217
L219:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875+int32(0))+76)))
	if v881 != int32(1) {
		v898 = v871
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v885 = v875 + int32(76)
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+43)))
	if v886 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v873 == int32(0) {
		v898 = v871
		goto L218
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v873 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v891 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v873))) = v891
	v901 = v891
	goto L217
L225:
	;
	v894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v885)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v873))) = v894
	goto L227
L226:
	;
	goto L227
L227:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v885)+44))
	v898 = v896
	goto L218
L228:
	;
	v922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v922) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v907+(v903^int32(-1))<<(uint(int32(2))%32))))
	v921 = v913
	goto L228
L230:
	;
	goto L231
L231:
	;
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v921 = v915 + v903<<(uint(int32(13))%32) + int32(-8192)
	goto L228
L232:
	;
	v930 = int32(base.Ui32(v922+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L234
L233:
	;
	v930 = int32(0)
	goto L234
L234:
	;
	if base.Ui32(v930&int32(_a_F_heap_redo_6)) < base.Ui32(v902) {
		goto L2
	} else {
		goto L235
	}
L235:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v921+v902<<(uint(int32(2))%32))+20))
	if v937&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	v946 = v921 + v937&int32(_a_F_heap_redo_11)
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+22)))
	v948 = int32(base.Ui32(v937)>>(uint(int32(17))%32)) - v947
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v948 != v949 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	if v948 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	base.MemoryCopy(m, v947+v946, v901, v948)
	goto L240
L239:
	;
	goto L240
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v921))) = base.I64_rotr(v863, int64(32))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L16
	} else {
		goto L241
	}
L241:
	;
	goto L216
L242:
	;
	F_UnlockReleaseBuffer(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L16
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+12)))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v862)+8))
	F_ProcessCommittedInvalidationMessages(m, v862+int32(20), v969, v970, v971, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L16
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	goto L8
L247:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_17), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L16
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(480), int32(_a_F_heap_redo_19))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L16
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_20), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L16
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(505), int32(_a_F_heap_redo_19))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L16
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L16
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(380), int32(_a_F_heap_redo_22))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L16
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L16
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(977), int32(_a_F_heap_redo_23))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L16
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L16
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1037), int32(_a_F_heap_redo_24))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L16
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L16
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1157), int32(_a_F_heap_redo_25))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L16
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_26), int32(0))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L16
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1163), int32(_a_F_heap_redo_25))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L16
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_scan_stream_read_next_serial(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
		v10 = int32(-1)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v11 == int32(0) {
			v35 = v10
			v39 = v35
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			if v14 == int32(0) {
				v35 = v10
				v39 = v35
			} else {
				if v9 == int32(1) {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v39 = v19
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v20 & int32(-129)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v14 != int32(-1) {
						v30 = base.I32_rem_u_s(v24+v14-int32(1), v11)
						v39 = v30
					} else {
						if v24 != 0 {
							v39 = v24 - int32(1)
						} else {
							v35 = v11 - int32(1)
							v39 = v35
						}
					}
				}
			}
		}
		v40 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v40)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v39
		return v39
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
		if v45 == int32(1) {
			v49 = v44 + int32(1)
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			if base.Ui32(v49) < base.Ui32(v51) {
				v53 = v49
			} else {
				v53 = int32(0)
			}
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
			if v54&int32(128) != 0 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ss_report_location(m, v57, v53)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v62 == v53 {
						v64 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v64
						return v64
					} else {
						v68 = int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
						if v69 == v68 {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v53
							return v53
						} else {
							v75 = v69 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v75
							if v75 == int32(0) {
								v101 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v101
								return v101
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v53
								return v53
							}
						}
					}
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				if v62 == v53 {
					v64 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v64
					return v64
				} else {
					v68 = int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					if v69 == v68 {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v53
						return v53
					} else {
						v75 = v69 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v75
						if v75 == int32(0) {
							v101 = v68
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v101
							return v101
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v53
							return v53
						}
					}
				}
			}
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if v81 == v44 {
				v83 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v83
				return v83
			} else {
				v87 = int32(-1)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v88 != v87 {
					v92 = v88 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v92
					if v92 == int32(0) {
						v101 = v87
					} else {
						if v44 != 0 {
							v98 = v44
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v98 = v97
						}
						v101 = v98 - int32(1)
					}
				} else {
					if v44 != 0 {
						v98 = v44
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v98 = v97
					}
					v101 = v98 - int32(1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v101
				return v101
			}
		}
	}
}
func F_heap_truncate_check_FKs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = v3
	v22 = v3
	goto L6
L4:
	;
	v49 = v3
	goto L5
L5:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v22<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+125)))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v49 = v41
	goto L5
L8:
	;
	v43 = v22 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v21 = v41
		v22 = v43
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
	if v35 != int32(112) {
		v41 = v21
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v39 = F_lappend_oid(m, v21, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return
L14:
	;
	v41 = v39
	goto L8
L15:
	;
	goto L7
L16:
	;
	v56 = F_heap_truncate_find_FKs(m, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v60 <= int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v64 = int32(0)
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v64<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v76
	v82 = F_list_make1_impl(m, int32(472), v11+int32(40))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L1
L22:
	;
	v206 = v64 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v206 < v207 {
		v64 = v206
		goto L20
	} else {
		goto L60
	}
L23:
	;
	v84 = F_heap_truncate_find_FKs(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	if v84 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v89 <= v88 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v96 = v88
	goto L27
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v96<<(uint(int32(2))%32))))
	v105 = int32(0)
	if v49 == v105 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v148 = F_get_rel_name(m, v76)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L46
	}
L29:
	;
	if v143 != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v143 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v111 <= int32(0) {
		v137 = v105
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = v137
	goto L29
L34:
	;
	v114 = int32(0)
	if v114 < v111 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = v111
	goto L37
L36:
	;
	v117 = v114
	goto L37
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v120 = int32(0)
	goto L38
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118+v120<<(uint(int32(2))%32))))
	v129 = base.B2i32(v128 == v104)
	if v128 == v104 {
		v137 = v129
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v137 = v129
	goto L33
L40:
	;
	v131 = v120 + int32(1)
	if v131 != v117 {
		v120 = v131
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v145 = v96 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v145 < v146 {
		v96 = v145
		goto L27
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L28
L45:
	;
	goto L22
L46:
	;
	v150 = F_get_rel_name(m, v104)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	if l1 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errmsg(m, int32(_a_F_heap_truncate_check_FKs_0), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_errmsg(m, int32(_a_F_heap_truncate_check_FKs_1), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L57
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v150
	F_errdetail(m, int32(_a_F_heap_truncate_check_FKs_2), v11+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v150
	F_errhint(m, int32(_a_F_heap_truncate_check_FKs_3), v11+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_heap_truncate_check_FKs_4), int32(3749), int32(_a_F_heap_truncate_check_FKs_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v150
	F_errdetail(m, int32(_a_F_heap_truncate_check_FKs_6), v11)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_heap_truncate_check_FKs_4), int32(3740), int32(_a_F_heap_truncate_check_FKs_5))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	goto L21
}
func F_heap_tuple_should_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v15 = int32(768)
	if v14&v15 == v15 {
		v51 = v5
		v52 = v14
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v19) < base.Ui32(int32(3)) {
			v51 = v5
			v52 = v14
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v22))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v19)) == int32(0) {
				v34 = base.B2i32(base.Ui32(v19) < base.Ui32(v22))
			} else {
				v34 = int32(base.Ui32(v19-v22) >> (uint(int32(31)) % 32))
			}
			if v34 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
			} else {
			}
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v36))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v19)) == int32(0) {
				v48 = base.B2i32(base.Ui32(v19) < base.Ui32(v36))
			} else {
				v48 = int32(base.Ui32(v19-v36) >> (uint(int32(31)) % 32))
			}
			v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			v51 = v48
			v52 = v49
		}
	}
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52&int32(_a_F_heap_tuple_should_freeze_0) != 0 {
		v57 = int32(0)
	} else {
		v57 = v54
	}
	if base.Ui32(int32(3)) <= base.Ui32(v57) {
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v60))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v57)) == int32(0) {
			v72 = base.B2i32(base.Ui32(v57) < base.Ui32(v60))
		} else {
			v72 = int32(base.Ui32(v57-v60) >> (uint(int32(31)) % 32))
		}
		if v72 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57
		} else {
		}
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v74))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v57)) == int32(0) {
			v86 = base.B2i32(base.Ui32(v57) < base.Ui32(v74))
		} else {
			v86 = int32(base.Ui32(v57-v74) >> (uint(int32(31)) % 32))
		}
		v189 = v86 | v51
		v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
		if base.Ui32(v193) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
			v217 = v189
		} else {
			v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v196) < base.Ui32(int32(3)) {
				v217 = v189
			} else {
				v199 = int32(1)
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v200))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
					v212 = base.B2i32(base.Ui32(v196) < base.Ui32(v200))
				} else {
					v212 = int32(base.Ui32(v196-v200) >> (uint(int32(31)) % 32))
				}
				if v212 == int32(0) {
					v217 = v199
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v196
					v217 = v199
				}
			}
		}
		m.G0 = v12 + int32(16)
		return v217 & int32(1)
	} else {
		v92 = v52 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & v54
		if v92 == int32(0) {
			v189 = v51
			v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			if base.Ui32(v193) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
				v217 = v189
			} else {
				v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v196) < base.Ui32(int32(3)) {
					v217 = v189
				} else {
					v199 = int32(1)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v200))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
						v212 = base.B2i32(base.Ui32(v196) < base.Ui32(v200))
					} else {
						v212 = int32(base.Ui32(v196-v200) >> (uint(int32(31)) % 32))
					}
					if v212 == int32(0) {
						v217 = v199
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v196
						v217 = v199
					}
				}
			}
			m.G0 = v12 + int32(16)
			return v217 & int32(1)
		} else {
			if v52&int32(_a_F_heap_tuple_should_freeze_2) == int32(_a_F_heap_tuple_should_freeze_3) {
				v99 = int32(1)
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if int32(base.Ui32(v92-v100)>>(uint(int32(31))%32)) == int32(0) {
					v189 = v99
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
					v189 = v99
				}
				v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				if base.Ui32(v193) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
					v217 = v189
				} else {
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v196) < base.Ui32(int32(3)) {
						v217 = v189
					} else {
						v199 = int32(1)
						v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v200))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
							v212 = base.B2i32(base.Ui32(v196) < base.Ui32(v200))
						} else {
							v212 = int32(base.Ui32(v196-v200) >> (uint(int32(31)) % 32))
						}
						if v212 == int32(0) {
							v217 = v199
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v196
							v217 = v199
						}
					}
				}
				m.G0 = v12 + int32(16)
				return v217 & int32(1)
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if int32(base.Ui32(v92-v107)>>(uint(int32(31))%32)) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
				} else {
				}
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v117 = int32(base.Ui32(v92-v113)>>(uint(int32(31))%32)) | v51
				v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				v130 = F_GetMultiXactIdMembers(m, v92, v12+int32(12), int32(base.Ui32(v120&int32(128))>>(uint(int32(7))%32))|base.B2i32(v120&int32(_a_F_heap_tuple_should_freeze_4) == int32(64)))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					if v130 <= int32(0) {
						v189 = v117
						v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						if base.Ui32(v193) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
							v217 = v189
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v196) < base.Ui32(int32(3)) {
								v217 = v189
							} else {
								v199 = int32(1)
								v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v200))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
									v212 = base.B2i32(base.Ui32(v196) < base.Ui32(v200))
								} else {
									v212 = int32(base.Ui32(v196-v200) >> (uint(int32(31)) % 32))
								}
								if v212 == int32(0) {
									v217 = v199
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v196
									v217 = v199
								}
							}
						}
						m.G0 = v12 + int32(16)
						return v217 & int32(1)
					} else {
						v141 = v117
						v142 = int32(0)
						for {
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v146 = int32(3)
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v142<<(uint(v146)%32))))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v150))&base.B2i32(base.Ui32(v146) <= base.Ui32(v149)) == int32(0) {
								v162 = base.B2i32(base.Ui32(v149) < base.Ui32(v150))
							} else {
								v162 = int32(base.Ui32(v149-v150) >> (uint(int32(31)) % 32))
							}
							if v162 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v149
							} else {
							}
							v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v164))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v149)) == int32(0) {
								v176 = base.B2i32(base.Ui32(v149) < base.Ui32(v164))
							} else {
								v176 = int32(base.Ui32(v149-v164) >> (uint(int32(31)) % 32))
							}
							v177 = v176 | v141
							v179 = v142 + int32(1)
							if v179 != v130 {
								v141 = v177
								v142 = v179
								continue
							} else {
								break
							}
							break
						}
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						F_pfree(m, v181)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return int32(0)
						} else {
							v189 = v177
							v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							if base.Ui32(v193) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
								v217 = v189
							} else {
								v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if base.Ui32(v196) < base.Ui32(int32(3)) {
									v217 = v189
								} else {
									v199 = int32(1)
									v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v200))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
										v212 = base.B2i32(base.Ui32(v196) < base.Ui32(v200))
									} else {
										v212 = int32(base.Ui32(v196-v200) >> (uint(int32(31)) % 32))
									}
									if v212 == int32(0) {
										v217 = v199
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v196
										v217 = v199
									}
								}
							}
							m.G0 = v12 + int32(16)
							return v217 & int32(1)
						}
					}
				}
			}
		}
	}
}
func F_heap_vac_scan_next_block(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v19 = v17 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if base.Ui32(v20) <= base.Ui32(v19) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v166
L2:
	;
	v22 = int32(-1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v23 == int32(0) {
		v166 = v22
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if base.B2i32(base.Ui32(v19) <= base.Ui32(v32))&base.B2i32(v32 != int32(-1)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_ReleaseBuffer(m, v23)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = int32(0)
	v166 = v22
	goto L1
L8:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v166 = v165
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+224)) = v134
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	v151 = v147 | v148<<(uint(int32(1))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v151)
	goto L8
L10:
	;
	v132 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v132)
	v134 = v70
	goto L9
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v39
	v44 = v32
	v47 = int32(0)
	goto L15
L12:
	;
	v116 = v32
	v120 = v19
	goto L13
L13:
	;
	if base.Ui32(v116) <= base.Ui32(v120) {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)) = uint8(v105)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+228)) = v70
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = v109
	v113 = base.B2i32(base.Ui32(int32(31)) < base.Ui32(v56-v17))
	if v47&v113 != 0 {
		goto L10
	} else {
		goto L30
	}
L15:
	;
	v56 = v44
	goto L17
L16:
	;
	v105 = v94 ^ int32(1)
	goto L14
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v70 = v56 + int32(1)
	v73 = F_visibilitymap_get_status(m, v68, v70, v15+int32(12))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v94 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if base.Ui32(v75) <= base.Ui32(v70) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+252)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(l1)+240)) = v75 + int32(_a_F_heap_vac_scan_next_block_0)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v82 = int32(0)
	v84 = v73 & int32(1)
	if base.B2i32(v84 == v82)|base.B2i32(v56 == v20-int32(2)) != 0 {
		v105 = v82
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v89 != int32(1) {
		v105 = v82
		goto L14
	} else {
		goto L24
	}
L24:
	;
	if v73&int32(2) != 0 {
		v56 = v70
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v98 == int32(0) {
		v44 = v70
		v47 = int32(1)
		goto L15
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L16
L29:
	;
	goto L28
L30:
	;
	if base.Ui32(int32(31)) < base.Ui32(v56-v17) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = v70
	goto L33
L32:
	;
	v115 = v19
	goto L33
L33:
	;
	v116 = v70
	v120 = v115
	goto L13
L34:
	;
	v134 = v120
	goto L9
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+224)) = v120
	v130 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v130)
	goto L8
}
func F_heap_xlog_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v637 int64
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(_a_F_heap_xlog_update_0)
	m.G0 = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	F_XLogRecGetBlockTag(m, l0, v3, v26+int32(_a_F_heap_xlog_update_1), v3, v26+int32(_a_F_heap_xlog_update_2))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v40 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v47 < int32(1) {
		v71 = v40
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	goto L3
L5:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(52))+76)))
	if v52 != int32(1) {
		v71 = v40
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L8
L8:
	;
	goto L9
L9:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	if v26+int32(_a_F_heap_xlog_update_3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(128))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1]))) = v63
	goto L15
L14:
	;
	goto L15
L15:
	;
	v71 = int32(1)
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1]))) = v72
	goto L19
L18:
	;
	goto L19
L19:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v77&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v82
	v86 = F_CreateFakeRelcacheEntry(m, v26+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v107 = v72
	goto L22
L22:
	;
	v109 = int32(base.Ui32(v72) >> (uint(int32(16)) % 32))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1])))
	v114 = F_XLogReadBufferForRedo(m, l0, base.B2i32(v110 != v107), v26+int32(_a_F_heap_xlog_update_4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1])))
	F_visibilitymap_pin(m, v86, v90, v26+int32(48))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1])))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	F_ReleaseBuffer(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v86)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	v107 = v105
	goto L22
L28:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L191
	}
L29:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L188
	}
L30:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L185
	}
L31:
	;
	if v114 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[4])))
	if v119 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v251 = v3
	v257 = v3
	goto L34
L34:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[1])))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	if v258 == v259 {
		goto L67
	} else {
		goto L68
	}
L35:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v138) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+(v119^int32(-1))<<(uint(int32(2))%32))))
	v137 = v129
	goto L35
L37:
	;
	goto L38
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[6]))
	v137 = v131 + v119<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v146 = int32(base.Ui32(v138+int32(_a_F_heap_xlog_update_5)) >> (uint(int32(2)) % 32))
	goto L41
L40:
	;
	v146 = int32(0)
	goto L41
L41:
	;
	if base.Ui32(v146&int32(_a_F_heap_xlog_update_6)) < base.Ui32(v118) {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v151 = v137 + int32(20)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v118<<(uint(int32(2))%32))))
	if v155&int32(_a_F_heap_xlog_update_7) != int32(_a_F_heap_xlog_update_8) {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v162 = v137 + v155&int32(_a_F_heap_xlog_update_9)
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162)+20)))
	v165 = v163 & int32(_a_F_heap_xlog_update_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+20)) = uint16(v165)
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162)+18)))
	v169 = v167 & int32(-24577)
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v172 = v169 | int32(_a_F_heap_xlog_update_11)
	goto L46
L45:
	;
	v172 = v169
	goto L46
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+18)) = uint16(v172)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+18)) = uint16(v172)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+20)) = uint16(v165)
	v177 = int32(1)
	v196 = v174<<(uint(v177)%32)&int32(16) | (v174<<(uint(int32(4))%32)&int32(64) | (v174<<(uint(int32(6))%32)&int32(128) | v174&v177<<(uint(int32(12))%32))) | v165
	if v174&int32(15) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+20)) = uint16(v196)
	goto L49
L48:
	;
	goto L49
L49:
	;
	if v174&int32(16) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v203 = v172 | int32(_a_F_heap_xlog_update_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+18)) = uint16(v203)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v207 = v196 & int32(_a_F_heap_xlog_update_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+20)) = uint16(v207)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v205
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+16)) = uint16(v76)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+14)) = uint16(v72)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+12)) = uint16(v109)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v215 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v235&int32(1) != 0 {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v215))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v217)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v233
	goto L53
L57:
	;
	if v229 == int32(0) {
		goto L53
	} else {
		goto L61
	}
L58:
	;
	v229 = base.B2i32(base.Ui32(v217) < base.Ui32(v215))
	goto L57
L59:
	;
	goto L60
L60:
	;
	v229 = int32(base.Ui32(v217-v215) >> (uint(int32(31)) % 32))
	goto L57
L61:
	;
	goto L56
L62:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+10)))
	v240 = v238 & int32(_a_F_heap_xlog_update_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+10)) = uint16(v240)
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v137))) = base.I64_rotr(v28, int64(32))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[4])))
	F_MarkBufferDirty(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v251 = v162
	v257 = int32(base.Ui32(v155) >> (uint(int32(17)) % 32))
	goto L34
L66:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v347&int32(2) != 0 {
		goto L90
	} else {
		goto L91
	}
L67:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[7]))) = v261
	v346 = v114
	goto L66
L68:
	;
	goto L69
L69:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(v263)+48)))
	if v264 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v268 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v343 = F_XLogReadBufferForRedo(m, l0, int32(0), v26+int32(_a_F_heap_xlog_update_15))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L89
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[7]))) = v268
	v271 = int32(0)
	if v268 < v271 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v290 = int32(_a_F_heap_xlog_update_12)
	v291 = int32(0)
	if v291|(v289&int32(3)|int32(1)) == v291 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+(v268^int32(-1))<<(uint(int32(2))%32))))
	v289 = v281
	goto L74
L76:
	;
	goto L77
L77:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[6]))
	v289 = v283 + v268<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	v346 = v271
	goto L66
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+10)) = int32(_a_F_heap_xlog_update_16)
	v331 = int32(_a_F_heap_xlog_update_17)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+18)) = uint16(v331)
	v337 = int32(_a_F_heap_xlog_update_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+16)) = uint16(v337)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+14)) = uint16(v337)
	goto L78
L80:
	;
	goto L83
L81:
	;
	goto L82
L82:
	;
	goto L88
L83:
	;
	v308 = v289 + v290
	v310 = v289 + int32(4)
	if base.Ui32(v310) < base.Ui32(v308) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v312 = v308
	goto L86
L85:
	;
	v312 = v310
	goto L86
L86:
	;
	v317 = (v289^int32(-1)+v312)&int32(-4) + int32(4)
	if v317 == int32(0) {
		goto L79
	} else {
		goto L87
	}
L87:
	;
	base.MemoryFill(m, v289, int32(0), v317)
	goto L79
L88:
	;
	base.MemoryFill(m, v289, int32(0), v290)
	goto L79
L89:
	;
	v346 = v343
	goto L66
L90:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v350
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v352
	v356 = F_CreateFakeRelcacheEntry(m, v26+int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v376 = int32(0)
	if v346 == v376 {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	F_visibilitymap_pin(m, v356, v360, v26+int32(48))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v368 = F_visibilitymap_clear(m, v365, v366, int32(3))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	F_ReleaseBuffer(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_pfree(m, v356)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L92
L98:
	;
	v379 = int32(0)
	v381 = v26 + int32(44)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+72))
	if v384 < v379 {
		v406 = v379
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v607 = v376
	goto L100
L100:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[4])))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[7])))
	if v619 != 0 {
		goto L172
	} else {
		goto L173
	}
L101:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[7])))
	if v412 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L102:
	;
	v409 = v406
	goto L101
L103:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+int32(0))+76)))
	if v389 != int32(1) {
		v406 = v379
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v393 = v383 + int32(76)
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+43)))
	if v394 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v381 == int32(0) {
		v406 = v379
		goto L102
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v381 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v399
	v409 = v399
	goto L101
L109:
	;
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v402
	goto L111
L110:
	;
	goto L111
L111:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v393)+44))
	v406 = v404
	goto L102
L112:
	;
	v431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430)+12)))
	if base.Ui32(v431) < base.Ui32(int32(25)) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416+(v412^int32(-1))<<(uint(int32(2))%32))))
	v430 = v422
	goto L112
L114:
	;
	goto L115
L115:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[6]))
	v430 = v424 + v412<<(uint(int32(13))%32) + int32(-8192)
	goto L112
L116:
	;
	v440 = int32(1)
	goto L118
L117:
	;
	v440 = int32(base.Ui32(v431+int32(_a_F_heap_xlog_update_5))>>(uint(int32(2))%32)) + int32(1)
	goto L118
L118:
	;
	if base.Ui32(v440&int32(_a_F_heap_xlog_update_6)) < base.Ui32(v410) {
		goto L29
	} else {
		goto L119
	}
L119:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v446&int32(32) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409))))
	v452 = v409 + int32(2)
	v453 = v449
	goto L122
L121:
	;
	v452 = v409
	v453 = v376
	goto L122
L122:
	;
	if v446&int32(64) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	v460 = v452 + int32(2)
	v461 = v457
	goto L125
L124:
	;
	v460 = v452
	v461 = int32(0)
	goto L125
L125:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+4)))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460)+2)))
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460))))
	v465 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+63)) = v465
	*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = v465
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v465
	v472 = v460 + int32(5)
	v473 = v409 + v444 - v472
	v475 = v26 + int32(71)
	if v453 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v461 != 0 {
		goto L142
	} else {
		goto L143
	}
L127:
	;
	v477 = v462 - int32(23)
	if v477 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	if v473 != 0 {
		goto L139
	} else {
		goto L140
	}
L130:
	;
	base.MemoryCopy(m, v475, v472, v477)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v481 = v26 + int32(48) + v462
	if v453 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+22)))
	base.MemoryCopy(m, v481, v251+v482, v453)
	goto L135
L134:
	;
	goto L135
L135:
	;
	v485 = v453 + v481
	v486 = v473 - v477
	if v486 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	base.MemoryCopy(m, v485, v472+v477, v486)
	goto L138
L137:
	;
	goto L138
L138:
	;
	v495 = v485 + v486
	goto L126
L139:
	;
	base.MemoryCopy(m, v475, v472, v473)
	goto L141
L140:
	;
	goto L141
L141:
	;
	v495 = v473 + v475
	goto L126
L142:
	;
	base.MemoryCopy(m, v495, v251+v257-v461, v461)
	goto L144
L143:
	;
	goto L144
L144:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+70)) = uint8(v462)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+66)) = uint16(v464)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+36))
	v504 = v463 & int32(_a_F_heap_xlog_update_18)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+68)) = uint16(v504)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v502
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+64)) = uint16(v76)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+62)) = uint16(v72)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+60)) = uint16(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v509
	v521 = F_PageAddItemExtended(m, v430, v26+int32(48), v453+v461+v473+int32(23), v410, int32(3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	if v521 == int32(0) {
		goto L28
	} else {
		goto L146
	}
L146:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v525&int32(2) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430)+10)))
	v530 = v528 & int32(_a_F_heap_xlog_update_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v430)+10)) = uint16(v530)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v535 = int32(4)
	v536 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430)+14)))
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430)+12)))
	v538 = v536 - v537
	if v538 <= v535 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v430))) = base.I64_rotr(v28, int64(32))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[7])))
	F_MarkBufferDirty(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L169
	}
L151:
	;
	v541 = v535
	goto L153
L152:
	;
	v541 = v538
	goto L153
L153:
	;
	v543 = v541 - int32(4)
	if v543 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v598 = int32(0)
	goto L150
L155:
	;
	goto L156
L156:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v537) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v598 = v543
	goto L150
L158:
	;
	v554 = int32(base.Ui32(v537+int32(_a_F_heap_xlog_update_5)) >> (uint(int32(2)) % 32))
	goto L160
L159:
	;
	v554 = int32(0)
	goto L160
L160:
	;
	if base.Ui32(v554&int32(_a_F_heap_xlog_update_6)) < base.Ui32(int32(291)) {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+10)))
	if v559&int32(1) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v598 = int32(0)
	goto L150
L163:
	;
	goto L164
L164:
	;
	v568 = int32(1)
	goto L165
L165:
	;
	v577 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430+int32(20)+v568&int32(_a_F_heap_xlog_update_6)<<(uint(int32(2))%32))+1)))
	if v577&int32(384) == int32(0) {
		goto L157
	} else {
		goto L167
	}
L166:
	;
	v598 = int32(0)
	goto L150
L167:
	;
	v583 = v568 + int32(1)
	v584 = int32(_a_F_heap_xlog_update_6)
	if base.Ui32(v583&v584) <= base.Ui32(v554&v584) {
		v568 = v583
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v607 = v598
	goto L100
L170:
	;
	if l1|v346|base.B2i32(base.Ui32(int32(1637)) < base.Ui32(v607)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L171:
	;
	F_UnlockReleaseBuffer(m, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L180
	}
L172:
	;
	if v618 == v619 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v624 = v618
	goto L174
L174:
	;
	if v624 == int32(0) {
		goto L170
	} else {
		goto L179
	}
L175:
	;
	v627 = v619
	goto L171
L176:
	;
	goto L177
L177:
	;
	F_UnlockReleaseBuffer(m, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[4])))
	v624 = v623
	goto L174
L179:
	;
	v627 = v624
	goto L171
L180:
	;
	goto L170
L181:
	;
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v637
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_heap_xlog_update[0])))
	F_XLogRecordPageWithFreeSpace(m, v26, v641, v607)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v26 + int32(_a_F_heap_xlog_update_0)
	return
L184:
	;
	goto L183
L185:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_19), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_20), int32(763), int32(_a_F_heap_xlog_update_21))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_22), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_20), int32(841), int32(_a_F_heap_xlog_update_21))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_23), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_20), int32(919), int32(_a_F_heap_xlog_update_21))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
