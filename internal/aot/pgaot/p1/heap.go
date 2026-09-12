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
	if v5&int32(4096) == v11 {
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[69]))
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[70]))
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
	v38 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v100 = *(*int32)(unsafe.Add(mBase, _consts[71]))
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
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v9&int32(256) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L52
	} else {
		goto L257
	}
L2:
	;
	v680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v682 = v680 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v682)
	goto L1
L3:
	;
	return v678
L4:
	;
	v596 = int32(1)
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v597&int32(2048) != 0 {
		v678 = v596
		goto L3
	} else {
		goto L217
	}
L5:
	;
	v12 = base.I32_extend16_s(v9)
	if v12&int32(512) != 0 {
		v678 = v4
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v12&int32(16384) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v589 = v145 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v589)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L52
	} else {
		goto L216
	}
L8:
	;
	v17 = int32(4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if base.Ui32(v18) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	if v12 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L11:
	;
	if v138 != 0 {
		v678 = v17
		goto L3
	} else {
		goto L51
	}
L12:
	;
	v138 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v29 == v18 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v138 = int32(1)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v33 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v138 = v130
	goto L11
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v37 == int32(0) {
		v130 = int32(0)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v101 = int32(0)
	v103 = v33 - int32(1)
	goto L41
L22:
	;
	v42 = v37
	goto L23
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v47 == int32(4) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v130 = int32(0)
	goto L18
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+80))
	if v94 != 0 {
		v42 = v94
		goto L23
	} else {
		goto L40
	}
L26:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v50 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v53 = int32(1)
	if v18 == v50 {
		v130 = v53
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	v57 = v55 - int32(1)
	if v57 < int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v62 = int32(0)
	v64 = v57
	goto L30
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v70 = int32(2)
	v71 = base.I32_div_s(v64-v62, v70)
	v72 = v71 + v62
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68+v72<<(uint(v70)%32))))
	if v76 == v18 {
		v130 = v53
		goto L18
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	v80 = F_TransactionIdPrecedes(m, v76, v18)
	mBase = m.M
	if v80 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v81 = v72 + int32(1)
	goto L35
L34:
	;
	v81 = v62
	goto L35
L35:
	;
	if v80 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v84 = v64
	goto L38
L37:
	;
	v84 = v72 - int32(1)
	goto L38
L38:
	;
	if v81 <= v84 {
		v62 = v81
		v64 = v84
		goto L30
	} else {
		goto L39
	}
L39:
	;
	goto L31
L40:
	;
	goto L24
L41:
	;
	v108 = int32(2)
	v109 = base.I32_div_s(v103-v101, v108)
	v110 = v109 + v101
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99+v110<<(uint(v108)%32))))
	v115 = base.B2i32(v114 == v18)
	if v114 == v18 {
		v130 = v115
		goto L18
	} else {
		goto L43
	}
L42:
	;
	v130 = v115
	goto L18
L43:
	;
	v118 = base.B2i32(base.Ui32(v114) < base.Ui32(v18))
	if base.Ui32(v114) < base.Ui32(v18) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v119 = v110 + int32(1)
	goto L46
L45:
	;
	v119 = v101
	goto L46
L46:
	;
	if base.Ui32(v114) < base.Ui32(v18) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v122 = v103
	goto L49
L48:
	;
	v122 = v110 - int32(1)
	goto L49
L49:
	;
	if v119 <= v122 {
		v101 = v119
		v103 = v122
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v139 = F_TransactionIdIsInProgress(m, v18)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	if v139 != 0 {
		v678 = v17
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v143 = F_TransactionIdDidCommit(m, v18)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v143 == int32(0) {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v149 = v145 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v149)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	return int32(0)
L58:
	;
	v158 = int32(3)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if base.Ui32(v159) < base.Ui32(v158) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui32(v301) < base.Ui32(int32(3)) {
		goto L111
	} else {
		goto L112
	}
L61:
	;
	if v279 != 0 {
		v678 = v158
		goto L3
	} else {
		goto L101
	}
L62:
	;
	v279 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v170 == v159 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v279 = int32(1)
	goto L61
L66:
	;
	goto L67
L67:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v174 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v279 = v271
	goto L61
L69:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v178 == int32(0) {
		v271 = int32(0)
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v242 = int32(0)
	v244 = v174 - int32(1)
	goto L91
L72:
	;
	v183 = v178
	goto L73
L73:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	if v188 == int32(4) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v271 = int32(0)
	goto L68
L75:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v183)+80))
	if v235 != 0 {
		v183 = v235
		goto L73
	} else {
		goto L90
	}
L76:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v191 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v194 = int32(1)
	if v159 == v191 {
		v271 = v194
		goto L68
	} else {
		goto L78
	}
L78:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v183)+52))
	v198 = v196 - int32(1)
	if v198 < int32(0) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v203 = int32(0)
	v205 = v198
	goto L80
L80:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+48))
	v211 = int32(2)
	v212 = base.I32_div_s(v205-v203, v211)
	v213 = v212 + v203
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v209+v213<<(uint(v211)%32))))
	if v217 == v159 {
		v271 = v194
		goto L68
	} else {
		goto L82
	}
L81:
	;
	goto L75
L82:
	;
	v221 = F_TransactionIdPrecedes(m, v217, v159)
	mBase = m.M
	if v221 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v222 = v213 + int32(1)
	goto L85
L84:
	;
	v222 = v203
	goto L85
L85:
	;
	if v221 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v225 = v205
	goto L88
L87:
	;
	v225 = v213 - int32(1)
	goto L88
L88:
	;
	if v222 <= v225 {
		v203 = v222
		v205 = v225
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
	v249 = int32(2)
	v250 = base.I32_div_s(v244-v242, v249)
	v251 = v250 + v242
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v240+v251<<(uint(v249)%32))))
	v256 = base.B2i32(v255 == v159)
	if v255 == v159 {
		v271 = v256
		goto L68
	} else {
		goto L93
	}
L92:
	;
	v271 = v256
	goto L68
L93:
	;
	v259 = base.B2i32(base.Ui32(v255) < base.Ui32(v159))
	if base.Ui32(v255) < base.Ui32(v159) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v260 = v251 + int32(1)
	goto L96
L95:
	;
	v260 = v242
	goto L96
L96:
	;
	if base.Ui32(v255) < base.Ui32(v159) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v263 = v244
	goto L99
L98:
	;
	v263 = v251 - int32(1)
	goto L99
L99:
	;
	if v260 <= v263 {
		v242 = v260
		v244 = v263
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L92
L101:
	;
	v280 = F_TransactionIdIsInProgress(m, v159)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L52
	} else {
		goto L102
	}
L102:
	;
	if v280 != 0 {
		v678 = v158
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v282 = F_TransactionIdDidCommit(m, v159)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L52
	} else {
		goto L104
	}
L104:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v282 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v288 = v284 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v288)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L52
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v296 = v284 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v296)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L52
	} else {
		goto L109
	}
L108:
	;
	return int32(0)
L109:
	;
	goto L4
L110:
	;
	if v421 != 0 {
		goto L150
	} else {
		goto L151
	}
L111:
	;
	v421 = int32(0)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v312 == v301 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v421 = int32(1)
	goto L110
L115:
	;
	goto L116
L116:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v316 <= int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v421 = v413
	goto L110
L118:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v320 == int32(0) {
		v413 = int32(0)
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v384 = int32(0)
	v386 = v316 - int32(1)
	goto L140
L121:
	;
	v325 = v320
	goto L122
L122:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	if v330 == int32(4) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v413 = int32(0)
	goto L117
L124:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v325)+80))
	if v377 != 0 {
		v325 = v377
		goto L122
	} else {
		goto L139
	}
L125:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	if v333 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v336 = int32(1)
	if v301 == v333 {
		v413 = v336
		goto L117
	} else {
		goto L127
	}
L127:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v325)+52))
	v340 = v338 - int32(1)
	if v340 < int32(0) {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v345 = int32(0)
	v347 = v340
	goto L129
L129:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v325)+48))
	v353 = int32(2)
	v354 = base.I32_div_s(v347-v345, v353)
	v355 = v354 + v345
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v351+v355<<(uint(v353)%32))))
	if v359 == v301 {
		v413 = v336
		goto L117
	} else {
		goto L131
	}
L130:
	;
	goto L124
L131:
	;
	v363 = F_TransactionIdPrecedes(m, v359, v301)
	mBase = m.M
	if v363 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v364 = v355 + int32(1)
	goto L134
L133:
	;
	v364 = v345
	goto L134
L134:
	;
	if v363 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v367 = v347
	goto L137
L136:
	;
	v367 = v355 - int32(1)
	goto L137
L137:
	;
	if v364 <= v367 {
		v345 = v364
		v347 = v367
		goto L129
	} else {
		goto L138
	}
L138:
	;
	goto L130
L139:
	;
	goto L123
L140:
	;
	v391 = int32(2)
	v392 = base.I32_div_s(v386-v384, v391)
	v393 = v392 + v384
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v382+v393<<(uint(v391)%32))))
	v398 = base.B2i32(v397 == v301)
	if v397 == v301 {
		v413 = v398
		goto L117
	} else {
		goto L142
	}
L141:
	;
	v413 = v398
	goto L117
L142:
	;
	v401 = base.B2i32(base.Ui32(v397) < base.Ui32(v301))
	if base.Ui32(v397) < base.Ui32(v301) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v402 = v393 + int32(1)
	goto L145
L144:
	;
	v402 = v384
	goto L145
L145:
	;
	if base.Ui32(v397) < base.Ui32(v301) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v405 = v386
	goto L148
L147:
	;
	v405 = v393 - int32(1)
	goto L148
L148:
	;
	if v402 <= v405 {
		v384 = v402
		v386 = v405
		goto L140
	} else {
		goto L149
	}
L149:
	;
	goto L141
L150:
	;
	v422 = int32(3)
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v423&int32(2048) != 0 {
		v678 = v422
		goto L3
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v568 = F_TransactionIdIsInProgress(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L52
	} else {
		goto L206
	}
L153:
	;
	if v423&int32(128) != 0 {
		v678 = v422
		goto L3
	} else {
		goto L154
	}
L154:
	;
	if v423&int32(4176) == int32(64) {
		v678 = v422
		goto L3
	} else {
		goto L155
	}
L155:
	;
	v432 = F_HeapTupleHeaderIsOnlyLocked(m, v6)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L52
	} else {
		goto L156
	}
L156:
	;
	if v432 != 0 {
		v678 = v422
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	if v436&int32(6272) == int32(4096) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if base.Ui32(v444) < base.Ui32(int32(3)) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v441 = F_HeapTupleGetUpdateXid(m, v6)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L52
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v444 = v443
	goto L158
L162:
	;
	v444 = v441
	goto L158
L163:
	;
	if v564 != 0 {
		goto L203
	} else {
		goto L204
	}
L164:
	;
	v564 = int32(0)
	goto L163
L165:
	;
	goto L166
L166:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v455 == v444 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v564 = int32(1)
	goto L163
L168:
	;
	goto L169
L169:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v459 <= int32(0) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v564 = v556
	goto L163
L171:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v463 == int32(0) {
		v556 = int32(0)
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v527 = int32(0)
	v529 = v459 - int32(1)
	goto L193
L174:
	;
	v468 = v463
	goto L175
L175:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v468)+20))
	if v473 == int32(4) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v556 = int32(0)
	goto L170
L177:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v468)+80))
	if v520 != 0 {
		v468 = v520
		goto L175
	} else {
		goto L192
	}
L178:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	if v476 == int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v479 = int32(1)
	if v444 == v476 {
		v556 = v479
		goto L170
	} else {
		goto L180
	}
L180:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v468)+52))
	v483 = v481 - int32(1)
	if v483 < int32(0) {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v488 = int32(0)
	v490 = v483
	goto L182
L182:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v468)+48))
	v496 = int32(2)
	v497 = base.I32_div_s(v490-v488, v496)
	v498 = v497 + v488
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v494+v498<<(uint(v496)%32))))
	if v502 == v444 {
		v556 = v479
		goto L170
	} else {
		goto L184
	}
L183:
	;
	goto L177
L184:
	;
	v506 = F_TransactionIdPrecedes(m, v502, v444)
	mBase = m.M
	if v506 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v507 = v498 + int32(1)
	goto L187
L186:
	;
	v507 = v488
	goto L187
L187:
	;
	if v506 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v510 = v490
	goto L190
L189:
	;
	v510 = v498 - int32(1)
	goto L190
L190:
	;
	if v507 <= v510 {
		v488 = v507
		v490 = v510
		goto L182
	} else {
		goto L191
	}
L191:
	;
	goto L183
L192:
	;
	goto L176
L193:
	;
	v534 = int32(2)
	v535 = base.I32_div_s(v529-v527, v534)
	v536 = v535 + v527
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v525+v536<<(uint(v534)%32))))
	v541 = base.B2i32(v540 == v444)
	if v540 == v444 {
		v556 = v541
		goto L170
	} else {
		goto L195
	}
L194:
	;
	v556 = v541
	goto L170
L195:
	;
	v544 = base.B2i32(base.Ui32(v540) < base.Ui32(v444))
	if base.Ui32(v540) < base.Ui32(v444) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v545 = v536 + int32(1)
	goto L198
L197:
	;
	v545 = v527
	goto L198
L198:
	;
	if base.Ui32(v540) < base.Ui32(v444) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v548 = v529
	goto L201
L200:
	;
	v548 = v536 - int32(1)
	goto L201
L201:
	;
	if v545 <= v548 {
		v527 = v545
		v529 = v548
		goto L193
	} else {
		goto L202
	}
L202:
	;
	goto L194
L203:
	;
	v565 = int32(4)
	goto L205
L204:
	;
	v565 = int32(3)
	goto L205
L205:
	;
	return v565
L206:
	;
	if v568 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	return int32(3)
L208:
	;
	goto L209
L209:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v573 = F_TransactionIdDidCommit(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L52
	} else {
		goto L210
	}
L210:
	;
	if v573 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_HeapTupleSetHintBits(m, v6, l1, int32(256), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L52
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v581 = v579 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v581)
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L52
	} else {
		goto L215
	}
L214:
	;
	goto L4
L215:
	;
	return int32(0)
L216:
	;
	goto L4
L217:
	;
	v602 = int32(0)
	if base.B2i32(v597&int32(128) == v602)&base.B2i32(v597&int32(4176) != int32(64)) == v602 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if v597&int32(1024) != 0 {
		v678 = v596
		goto L3
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if v597&int32(4096) != 0 {
		goto L233
	} else {
		goto L234
	}
L221:
	;
	if v597&int32(4096) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if v597&int32(4304) != int32(4224) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v629 = F_TransactionIdIsInProgress(m, v628)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L52
	} else {
		goto L230
	}
L225:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v621 = F_MultiXactIdIsRunning(m, v619, int32(1))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L52
	} else {
		goto L228
	}
L226:
	;
	v624 = v597
	goto L227
L227:
	;
	v626 = v624 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v626)
	goto L1
L228:
	;
	if v621 != 0 {
		v678 = v596
		goto L3
	} else {
		goto L229
	}
L229:
	;
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v624 = v623
	goto L227
L230:
	;
	if v629 != 0 {
		v678 = v596
		goto L3
	} else {
		goto L231
	}
L231:
	;
	v631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v633 = v631 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v633)
	goto L1
L232:
	;
	v678 = int32(2)
	goto L3
L233:
	;
	v637 = F_HeapTupleGetUpdateXid(m, v6)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L52
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if v597&int32(1024) == int32(0) {
		goto L247
	} else {
		goto L248
	}
L236:
	;
	v639 = F_TransactionIdIsInProgress(m, v637)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L52
	} else {
		goto L237
	}
L237:
	;
	if v639 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	return int32(4)
L239:
	;
	goto L240
L240:
	;
	v643 = F_TransactionIdDidCommit(m, v637)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L52
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
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v637
	goto L232
L243:
	;
	goto L244
L244:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v648 = F_MultiXactIdIsRunning(m, v646, int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L52
	} else {
		goto L245
	}
L245:
	;
	if v648 != 0 {
		v678 = v596
		goto L3
	} else {
		goto L246
	}
L246:
	;
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)))
	v652 = v650 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+20)) = uint16(v652)
	goto L1
L247:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v659 = F_TransactionIdIsInProgress(m, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L52
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v672
	goto L232
L250:
	;
	if v659 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	return int32(4)
L252:
	;
	goto L253
L253:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v664 = F_TransactionIdDidCommit(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L52
	} else {
		goto L254
	}
L254:
	;
	if v664 == int32(0) {
		goto L2
	} else {
		goto L255
	}
L255:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	F_HeapTupleSetHintBits(m, v6, l1, int32(1024), v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L52
	} else {
		goto L256
	}
L256:
	;
	goto L249
L257:
	;
	return int32(1)
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
					F_errmsg_internal(m, int32(482642), v7)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498320), int32(491), int32(301570))
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
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
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v492 int32
	_ = v492
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int64
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v731 int32
	_ = v731
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int64
	_ = v743
	var v745 int32
	_ = v745
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	v8 = l7
	v9 = l8
	v10 = l9
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	if l11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L85
	} else {
		goto L181
	}
L2:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v54
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v54
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
	v35 = base.B2i32(v8 != int32(105))
	goto L7
L6:
	;
	v35 = int32(0)
	goto L7
L7:
	;
	if v35 == int32(0) {
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
	v51 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v51 == int32(2) {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v43 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L12:
	;
	v42 = F_isTempToastNamespace(m, l1)
	mBase = m.M
	v43 = v42
	goto L14
L13:
	;
	v43 = int32(1)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v47 != int32(2) {
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
	v61 = l2
	goto L20
L19:
	;
	v61 = v54
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
	v64 = l2
	goto L23
L22:
	;
	v64 = v61
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
	v67 = l2
	goto L26
L25:
	;
	v67 = v64
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
	v70 = l2
	goto L29
L28:
	;
	v70 = v67
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
	v73 = l2
	goto L32
L31:
	;
	v73 = v70
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
	v76 = l2
	goto L35
L34:
	;
	v76 = v73
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
	v79 = l2
	goto L38
L37:
	;
	v79 = v76
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
	v83 = v79
	goto L41
L40:
	;
	v83 = int32(0)
	goto L41
L41:
	;
	switch v8 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L43
	default:
		v88 = l4
		v89 = int32(0)
		goto L42
	}
L42:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	if v83 != v92 {
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
	v87 = l4
	goto L46
L45:
	;
	v87 = l3
	goto L46
L46:
	;
	v88 = v87
	v89 = l14
	goto L42
L47:
	;
	if v89 != 0 {
		goto L168
	} else {
		goto L169
	}
L48:
	;
	v94 = v83
	goto L50
L49:
	;
	v94 = int32(0)
	goto L50
L50:
	;
	v95 = m.G0
	v97 = v95 - int32(48)
	m.G0 = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v102 = int32(1)
	if l3 <= int32(3591) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v170 == v10 {
		goto L79
	} else {
		goto L80
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
	if base.Ui32(l3-int32(2964)) < base.Ui32(int32(4)) {
		v170 = v102
		goto L52
	} else {
		goto L77
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
	if l3 <= int32(5999) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	switch l3 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v170 = v102
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
	v114 = l3 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v114) {
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
	v170 = v102
	goto L52
L63:
	;
	if int32(1)<<(uint(v114)%32)&int32(226492515) == int32(0) {
		goto L54
	} else {
		goto L64
	}
L64:
	;
	v170 = v102
	goto L52
L65:
	;
	if base.Ui32(l3-int32(3592)) < base.Ui32(int32(2)) {
		v170 = v102
		goto L52
	} else {
		goto L75
	}
L66:
	;
	v126 = l3 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v126) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	switch l3 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v170 = v102
		goto L52
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L53
	default:
		goto L71
	}
L69:
	;
	if int32(1)<<(uint(v126)%32)&int32(963) == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v170 = v102
	goto L52
L71:
	;
	if base.Ui32(l3-int32(6000)) < base.Ui32(int32(3)) {
		v170 = v102
		goto L52
	} else {
		goto L72
	}
L72:
	;
	v142 = l3 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v142) {
		goto L53
	} else {
		goto L73
	}
L73:
	;
	if int32(1)<<(uint(v142)%32)&int32(49153) != 0 {
		v170 = v102
		goto L52
	} else {
		goto L74
	}
L74:
	;
	goto L53
L75:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l3-int32(4060)) {
		goto L53
	} else {
		goto L76
	}
L76:
	;
	v170 = v102
	goto L52
L77:
	;
	if base.Ui32(l3-int32(2846)) < base.Ui32(int32(2)) {
		v170 = v102
		goto L52
	} else {
		goto L78
	}
L78:
	;
	goto L53
L79:
	;
	v173 = l3 - int32(1247)
	v178 = base.B2i32(base.Ui32(v173) < base.Ui32(int32(16))) & int32(base.Ui32(int32(61701))>>(uint(v173)%32))
	v180 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if v180 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L85
	} else {
		goto L164
	}
L82:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v189 = v180
	goto L84
L84:
	;
	v190 = int32(4489440)
	v191 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v189
	v195 = F_palloc0(m, int32(276))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L85
	} else {
		goto L87
	}
L85:
	;
	return int32(0)
L86:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v189 = v188
	goto L84
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+25)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v178
	v202 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+32)) = v203
	v209 = F_CreateTupleDescCopy(m, l6)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+52)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = int32(1)
	if v99 <= int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v324 = F_palloc0(m, int32(144))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L85
	} else {
		goto L101
	}
L91:
	;
	v236 = int32(0)
	v242 = int32(0)
	goto L92
L92:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v248 = int32(4)
	v252 = v236 * int32(100)
	v253 = v246 + v247<<(uint(v248)%32) + v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v258 = l6 + int32(20) + v254<<(uint(v248)%32) + v252
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+109)) = uint8(v259)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+110)) = uint8(v261)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+86)))
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+106)) = uint8(v263)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	F_populate_compact_attribute(m, v265, v236)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L85
	} else {
		goto L94
	}
L93:
	;
	if v270&int32(255) == int32(0) {
		goto L90
	} else {
		goto L99
	}
L94:
	;
	v268 = int32(1)
	v270 = v263 | v242&v268
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+86)))
	if v271 == v268 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v275 = v236 << (uint(int32(4)) % 32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275+(l6+int32(31))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v275+v276)+31)) = uint8(v279)
	goto L97
L96:
	;
	goto L97
L97:
	;
	v282 = int32(1)
	v285 = v236 + v282
	if v285 != v99 {
		v236 = v285
		v242 = v270 & v282
		goto L92
	} else {
		goto L98
	}
L98:
	;
	goto L93
L99:
	;
	v292 = F_palloc0(m, int32(20))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L85
	} else {
		goto L100
	}
L100:
	;
	v294 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v292)+16)) = uint8(v294)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+16)) = v292
	goto L90
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+48)) = v324
	v330 = F_strncpy(m, v324+int32(4), l0, int32(64))
	mBase = m.M
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v330)+63)) = uint8(v331)
	goto L102
L102:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+68)) = l1
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+119)) = uint8(v8)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v337)+120)) = uint16(v99)
	v339 = int32(0)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+72)) = v339
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+80)) = int32(10)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+118)) = uint8(v9)
	switch v9 - int32(112) {
	case 0, 5:
		v374 = v339
		v375 = int32(-1)
		goto L103
	default:
		goto L105
	case 4:
		goto L104
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+24)) = uint8(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = v375
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v378)+129)) = uint8(base.B2i32(v8 != int32(109)))
	v382 = int32(110)
	goto L113
L104:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v369 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v369 == int32(-1) {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v9
	F_errmsg_internal(m, int32(501860), v97+int32(16))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L85
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(498618), int32(3665), int32(263927))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L85
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v372 = v367
	goto L111
L110:
	;
	v372 = v369
	goto L111
L111:
	;
	v374 = int32(1)
	v375 = v372
	goto L103
L112:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+130)) = uint8(v399)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v403)+117)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+56)) = l3
	if v99 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if l1 == int32(11) {
		v399 = v382
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v386 = v8 - int32(109)
	if base.Ui32(int32(5)) < base.Ui32(v386&int32(255)) {
		v399 = v382
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v399 = base.I32_wrap_i64(int64(base.Ui64(int64(110425294138980)) >> (uint(base.I64_extend_i32_u(v386<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L112
L116:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+92)) = v94
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	if l10 != 0 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v409 = v99 & int32(3)
	v410 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v99) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v432 = v410
	v438 = int32(0)
	goto L121
L119:
	;
	v492 = v410
	goto L120
L120:
	;
	if v409 == int32(0) {
		goto L116
	} else {
		goto L124
	}
L121:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = int32(4)
	v448 = v432 * int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v442+v443<<(uint(v444)%32)+v448)+20)) = l3
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v451+v452<<(uint(v444)%32)+v448)+120)) = l3
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v458+v459<<(uint(v444)%32)+v448)+220)) = l3
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	*(*int32)(unsafe.Add(mBase, uint32(v465+v466<<(uint(v444)%32)+v448)+320)) = l3
	v473 = v432 + v444
	v475 = v438 + v444
	if v475 != v99&int32(2147483644) {
		v432 = v473
		v438 = v475
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v492 = v473
	goto L120
L123:
	;
	goto L122
L124:
	;
	v519 = v492
	v521 = v410
	goto L125
L125:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	*(*int32)(unsafe.Add(mBase, uint32(v529+v530<<(uint(int32(4))%32)+v519*int32(100))+20)) = l3
	v538 = int32(1)
	v541 = v521 + v538
	if v541 != v409 {
		v519 = v519 + v538
		v521 = v541
		goto L125
	} else {
		goto L127
	}
L126:
	;
	goto L116
L127:
	;
	goto L126
L128:
	;
	v578 = v195 + int32(56)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v195)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+60)) = v579
	v583 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+117)))
	if v585 != 0 {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+88)) = int32(0)
	F_RelationMapUpdateMap(m, l3, v88, v10, int32(1))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L85
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+88)) = v88
	goto L128
L132:
	;
	goto L128
L133:
	;
	F_RelationInitPhysicalAddr(m, v195)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L85
	} else {
		goto L137
	}
L134:
	;
	v586 = int32(0)
	goto L136
L135:
	;
	v586 = v583
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+64)) = v586
	goto L133
L137:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+84)) = l5
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v191
	switch v8 - int32(83) {
	case 0, 26, 31, 33:
		goto L139
	default:
		goto L138
	}
L138:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v603 = F_hash_search(m, v599, v578, int32(1), v97+int32(47))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L85
	} else {
		goto L141
	}
L139:
	;
	F_RelationInitTableAccessMethod(m, v195)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L85
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+47)))
	if v605 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	if v643 <= int32(31) {
		goto L156
	} else {
		goto L157
	}
L143:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v195
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v610 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v195
	goto L142
L146:
	;
	F_RelationDestroyRelation(m, v608, int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L85
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v617 == int32(0) {
		goto L142
	} else {
		goto L150
	}
L149:
	;
	goto L142
L150:
	;
	v622 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L85
	} else {
		goto L151
	}
L151:
	;
	if v622 == int32(0) {
		goto L142
	} else {
		goto L152
	}
L152:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v608)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v626 + int32(4)
	F_errmsg_internal(m, int32(684371), v97+int32(32))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L85
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(498618), int32(3738), int32(263927))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L85
	} else {
		goto L154
	}
L154:
	;
	goto L142
L155:
	;
	v660 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+26)) = uint8(v660)
	v663 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L85
	} else {
		goto L159
	}
L156:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	*(*int32)(unsafe.Add(mBase, _consts[302])) = v643 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v643<<(uint(int32(2))%32))+uint32(_consts[303]))) = v646
	goto L155
L157:
	;
	goto L158
L158:
	;
	v657 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[304])) = uint8(v657)
	goto L155
L159:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v666 + int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v671 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v673 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v673, v195, int32(1722816))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L85
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	m.G0 = v97 + int32(48)
	goto L47
L163:
	;
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = l0
	F_errmsg_internal(m, int32(655459), v97)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L85
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(498618), int32(3566), int32(263927))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L85
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+117)))
	if v741 != 0 {
		goto L177
	} else {
		goto L178
	}
L168:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+119)))
	switch v695 - int32(83) {
	case 0, 22:
		goto L171
	default:
		goto L167
	case 26, 31, 33:
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	if v94 == int32(0) {
		goto L167
	} else {
		goto L175
	}
L171:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v702
	v704 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v704
	v707 = F_RelationCreateStorage(m, v28, v9, int32(1))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L85
	} else {
		goto L174
	}
L172:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v195)+188))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+112))
	m.T0[v699].(func(*base.Module, int32, int32, int32, int32, int32))(m, v195, v195, v9, l12, l13)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L85
	} else {
		goto L173
	}
L173:
	;
	goto L167
L174:
	;
	goto L167
L175:
	;
	v711 = m.G0
	v713 = v711 - int32(32)
	m.G0 = v713
	v715 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v713)+28)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v713)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v713)+20)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v713)+16)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v713)+12)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = int32(1213)
	F_recordSharedDependencyOn(m, v713+int32(20), v713+int32(8), int32(116))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L85
	} else {
		goto L176
	}
L176:
	;
	m.G0 = v713 + int32(32)
	goto L167
L177:
	;
	v742 = int32(0)
	goto L179
L178:
	;
	v742 = v739
	goto L179
L179:
	;
	v743 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v195)+56)))
	F_pgstat_create_transactional(m, int32(2), v742, v743)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L85
	} else {
		goto L180
	}
L180:
	;
	m.G0 = v28 + int32(32)
	return v195
L181:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L85
	} else {
		goto L182
	}
L182:
	;
	v757 = F_get_namespace_name(m, l1)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L85
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v757
	F_errmsg(m, int32(675974), v28+int32(16))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L85
	} else {
		goto L184
	}
L184:
	;
	F_errdetail(m, int32(630874), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L85
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(495443), int32(323), int32(354814))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L85
	} else {
		goto L186
	}
L186:
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
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
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
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
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v507 int32
	_ = v507
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int64
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v557 int32
	_ = v557
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
				v342 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v343 = F_SnapBuildProcessChange(m, v21, v23, v342)
				mBase = m.M
				v344 = m.ExcPending
				if v344 != 0 {
					return
				} else {
					if v343 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v347 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v348 = m.G0
							v350 = v348 - int32(16)
							m.G0 = v350
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
							v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+64))
							v355 = int32(0)
							F_XLogRecGetBlockTag(m, v352, v355, v350+int32(4), v355, v355)
							mBase = m.M
							v361 = m.ExcPending
							if v361 != 0 {
								return
							} else {
								v362 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
								v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+88))
								if v362 != v364 {
									m.G0 = v350 + int32(16)
									m.G0 = v16 + int32(16)
									return
								} else {
									v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v366 != 0 {
										v367 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
										v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v367)+56)))
										v369 = F_filter_by_origin_cb_wrapper(m, l0, v368)
										mBase = m.M
										v370 = m.ExcPending
										if v370 != 0 {
											return
										} else {
											if v369 != 0 {
												m.G0 = v350 + int32(16)
												m.G0 = v16 + int32(16)
												return
											} else {
												v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v372 = F_ReorderBufferAllocChange(m, v371)
												mBase = m.M
												v373 = m.ExcPending
												if v373 != 0 {
													return
												} else {
													v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+7)))
													if v376&int32(8) != 0 {
														v379 = int32(10)
													} else {
														v379 = int32(2)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = v379
													v381 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
													v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v381)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v372)+16)) = uint16(v382)
													v384 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v372)+28)) = v384
													v386 = *(*int64)(unsafe.Add(mBase, uint32(v350)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v372)+20)) = v386
													v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+7)))
													if v388&int32(6) != 0 {
														v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v392 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
														v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
														v395 = v393 - int32(13)
														v396 = F_ReorderBufferAllocTupleBuf(m, v391, v395)
														mBase = m.M
														v397 = m.ExcPending
														if v397 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v372)+36)) = v396
															v399 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = v399
															*(*uint16)(unsafe.Add(mBase, uint32(v396)+8)) = uint16(v399)
															*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v396))) = v393 + int32(10)
															v408 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
															v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+12)))
															v410 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
															v411 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v411
															*(*int64)(unsafe.Add(mBase, uint32(v410)+15)) = v411
															*(*int64)(unsafe.Add(mBase, uint32(v410))) = v411
															v417 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
															if v395 != 0 {
																v422 = F__emscripten_memcpy_bulkmem(m, v417+int32(23), v354+int32(13), v395)
																mBase = m.M
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v410)+22)) = uint8(v409)
															*(*int32)(unsafe.Add(mBase, uint32(v410)+18)) = v408
															v431 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v372)+32)) = uint8(v431)
															v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v434 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
															v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+36))
															v436 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v433, v435, v436, v372, int32(0))
															mBase = m.M
															v439 = m.ExcPending
															if v439 != 0 {
																return
															} else {
																m.G0 = v350 + int32(16)
																m.G0 = v16 + int32(16)
																return
															}
														}
													} else {
														v431 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v372)+32)) = uint8(v431)
														v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v434 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
														v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+36))
														v436 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v433, v435, v436, v372, int32(0))
														mBase = m.M
														v439 = m.ExcPending
														if v439 != 0 {
															return
														} else {
															m.G0 = v350 + int32(16)
															m.G0 = v16 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v372 = F_ReorderBufferAllocChange(m, v371)
										mBase = m.M
										v373 = m.ExcPending
										if v373 != 0 {
											return
										} else {
											v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+7)))
											if v376&int32(8) != 0 {
												v379 = int32(10)
											} else {
												v379 = int32(2)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = v379
											v381 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
											v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v381)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v372)+16)) = uint16(v382)
											v384 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v372)+28)) = v384
											v386 = *(*int64)(unsafe.Add(mBase, uint32(v350)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v372)+20)) = v386
											v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+7)))
											if v388&int32(6) != 0 {
												v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v392 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
												v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
												v395 = v393 - int32(13)
												v396 = F_ReorderBufferAllocTupleBuf(m, v391, v395)
												mBase = m.M
												v397 = m.ExcPending
												if v397 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v372)+36)) = v396
													v399 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = v399
													*(*uint16)(unsafe.Add(mBase, uint32(v396)+8)) = uint16(v399)
													*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v396))) = v393 + int32(10)
													v408 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
													v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+12)))
													v410 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
													v411 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v410)+8)) = v411
													*(*int64)(unsafe.Add(mBase, uint32(v410)+15)) = v411
													*(*int64)(unsafe.Add(mBase, uint32(v410))) = v411
													v417 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
													if v395 != 0 {
														v422 = F__emscripten_memcpy_bulkmem(m, v417+int32(23), v354+int32(13), v395)
														mBase = m.M
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v410)+22)) = uint8(v409)
													*(*int32)(unsafe.Add(mBase, uint32(v410)+18)) = v408
													v431 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v372)+32)) = uint8(v431)
													v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v434 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
													v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+36))
													v436 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v433, v435, v436, v372, int32(0))
													mBase = m.M
													v439 = m.ExcPending
													if v439 != 0 {
														return
													} else {
														m.G0 = v350 + int32(16)
														m.G0 = v16 + int32(16)
														return
													}
												}
											} else {
												v431 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v372)+32)) = uint8(v431)
												v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v434 = *(*int32)(unsafe.Add(mBase, uint32(v352)+96))
												v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+36))
												v436 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												F_ReorderBufferQueueChange(m, v433, v435, v436, v372, int32(0))
												mBase = m.M
												v439 = m.ExcPending
												if v439 != 0 {
													return
												} else {
													m.G0 = v350 + int32(16)
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
				v175 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v176 = F_SnapBuildProcessChange(m, v21, v23, v175)
				mBase = m.M
				v177 = m.ExcPending
				if v177 != 0 {
					return
				} else {
					if v176 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v180 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+64))
							v184 = int32(0)
							F_XLogRecGetBlockTag(m, v181, v184, v16+int32(4), v184, v184)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
								v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+88))
								if v191 != v193 {
									m.G0 = v16 + int32(16)
									return
								} else {
									v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v195 != 0 {
										v196 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
										v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+56)))
										v198 = F_filter_by_origin_cb_wrapper(m, l0, v197)
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return
										} else {
											if v198 != 0 {
												m.G0 = v16 + int32(16)
												return
											} else {
												v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v201 = F_ReorderBufferAllocChange(m, v200)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = int32(1)
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
													v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v201)+16)) = uint16(v206)
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v208
													v210 = *(*int64)(unsafe.Add(mBase, uint32(v16)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v201)+20)) = v210
													v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
													if v212&int32(16) != 0 {
														v215 = int32(0)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
														v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
														if v218 < v215 {
															v240 = v215
															v243 = v240
														} else {
															v224 = v217 + int32(76)
															v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
															if v225 != int32(1) {
																v240 = v215
																v243 = v240
															} else {
																v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+43)))
																if v228 == int32(0) {
																	if v16 == int32(0) {
																		v240 = v215
																		v243 = v240
																	} else {
																		v233 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v233
																		v243 = v233
																	}
																} else {
																	if v16 != 0 {
																		v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+48)))
																		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v236
																	} else {
																	}
																	v238 = *(*int32)(unsafe.Add(mBase, uint32(v224)+44))
																	v240 = v238
																	v243 = v240
																}
															}
														}
														v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
														v248 = F_ReorderBufferAllocTupleBuf(m, v244, v245-int32(5))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v201)+40)) = v248
															v251 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
															v252 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v252
															*(*uint16)(unsafe.Add(mBase, uint32(v248)+8)) = uint16(v252)
															*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v248))) = v251 + int32(18)
															v261 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
															v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+4)))
															v263 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
															v264 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v264
															*(*int64)(unsafe.Add(mBase, uint32(v263)+15)) = v264
															*(*int64)(unsafe.Add(mBase, uint32(v263))) = v264
															v270 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
															v273 = int32(5)
															v276 = v251 - v273
															if v276 != 0 {
																v277 = F__emscripten_memcpy_bulkmem(m, v270+int32(23), v243+v273, v276)
																mBase = m.M
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v263)+22)) = uint8(v262)
															*(*int32)(unsafe.Add(mBase, uint32(v263)+18)) = v261
															v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
															v288 = v281
															if v288&int32(12) != 0 {
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+64))
																v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+68))
																v296 = v294 - int32(19)
																v297 = F_ReorderBufferAllocTupleBuf(m, v293, v296)
																mBase = m.M
																v298 = m.ExcPending
																if v298 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v201)+36)) = v297
																	v300 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v300
																	*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)) = uint16(v300)
																	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = int32(-1)
																	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v294 + int32(4)
																	v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+14))
																	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+18)))
																	v311 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
																	v312 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v312
																	*(*int64)(unsafe.Add(mBase, uint32(v311)+15)) = v312
																	*(*int64)(unsafe.Add(mBase, uint32(v311))) = v312
																	v318 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
																	if v296 != 0 {
																		v323 = F__emscripten_memcpy_bulkmem(m, v318+int32(23), v292+int32(19), v296)
																		mBase = m.M
																	} else {
																	}
																	*(*uint8)(unsafe.Add(mBase, uint32(v311)+22)) = uint8(v310)
																	*(*int32)(unsafe.Add(mBase, uint32(v311)+18)) = v309
																	v333 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
																	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
																	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
																	v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																	F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
																	mBase = m.M
																	v341 = m.ExcPending
																	if v341 != 0 {
																		return
																	} else {
																		m.G0 = v16 + int32(16)
																		return
																	}
																}
															} else {
																v333 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
																v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
																v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
																v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	m.G0 = v16 + int32(16)
																	return
																}
															}
														}
													} else {
														v288 = v212
														if v288&int32(12) != 0 {
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+64))
															v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+68))
															v296 = v294 - int32(19)
															v297 = F_ReorderBufferAllocTupleBuf(m, v293, v296)
															mBase = m.M
															v298 = m.ExcPending
															if v298 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v201)+36)) = v297
																v300 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v300
																*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)) = uint16(v300)
																*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v297))) = v294 + int32(4)
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+14))
																v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+18)))
																v311 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
																v312 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v312
																*(*int64)(unsafe.Add(mBase, uint32(v311)+15)) = v312
																*(*int64)(unsafe.Add(mBase, uint32(v311))) = v312
																v318 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
																if v296 != 0 {
																	v323 = F__emscripten_memcpy_bulkmem(m, v318+int32(23), v292+int32(19), v296)
																	mBase = m.M
																} else {
																}
																*(*uint8)(unsafe.Add(mBase, uint32(v311)+22)) = uint8(v310)
																*(*int32)(unsafe.Add(mBase, uint32(v311)+18)) = v309
																v333 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
																v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
																v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
																v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	m.G0 = v16 + int32(16)
																	return
																}
															}
														} else {
															v333 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
															v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
															v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
															v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
															mBase = m.M
															v341 = m.ExcPending
															if v341 != 0 {
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
										v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v201 = F_ReorderBufferAllocChange(m, v200)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = int32(1)
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
											v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v201)+16)) = uint16(v206)
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v208
											v210 = *(*int64)(unsafe.Add(mBase, uint32(v16)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v201)+20)) = v210
											v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
											if v212&int32(16) != 0 {
												v215 = int32(0)
												v217 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
												v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
												if v218 < v215 {
													v240 = v215
													v243 = v240
												} else {
													v224 = v217 + int32(76)
													v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
													if v225 != int32(1) {
														v240 = v215
														v243 = v240
													} else {
														v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+43)))
														if v228 == int32(0) {
															if v16 == int32(0) {
																v240 = v215
																v243 = v240
															} else {
																v233 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v16))) = v233
																v243 = v233
															}
														} else {
															if v16 != 0 {
																v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+48)))
																*(*int32)(unsafe.Add(mBase, uint32(v16))) = v236
															} else {
															}
															v238 = *(*int32)(unsafe.Add(mBase, uint32(v224)+44))
															v240 = v238
															v243 = v240
														}
													}
												}
												v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
												v248 = F_ReorderBufferAllocTupleBuf(m, v244, v245-int32(5))
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v201)+40)) = v248
													v251 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
													v252 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v252
													*(*uint16)(unsafe.Add(mBase, uint32(v248)+8)) = uint16(v252)
													*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v248))) = v251 + int32(18)
													v261 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
													v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+4)))
													v263 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
													v264 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v264
													*(*int64)(unsafe.Add(mBase, uint32(v263)+15)) = v264
													*(*int64)(unsafe.Add(mBase, uint32(v263))) = v264
													v270 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
													v273 = int32(5)
													v276 = v251 - v273
													if v276 != 0 {
														v277 = F__emscripten_memcpy_bulkmem(m, v270+int32(23), v243+v273, v276)
														mBase = m.M
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v263)+22)) = uint8(v262)
													*(*int32)(unsafe.Add(mBase, uint32(v263)+18)) = v261
													v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
													v288 = v281
													if v288&int32(12) != 0 {
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+64))
														v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+68))
														v296 = v294 - int32(19)
														v297 = F_ReorderBufferAllocTupleBuf(m, v293, v296)
														mBase = m.M
														v298 = m.ExcPending
														if v298 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v201)+36)) = v297
															v300 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v300
															*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)) = uint16(v300)
															*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v297))) = v294 + int32(4)
															v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+14))
															v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+18)))
															v311 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
															v312 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v312
															*(*int64)(unsafe.Add(mBase, uint32(v311)+15)) = v312
															*(*int64)(unsafe.Add(mBase, uint32(v311))) = v312
															v318 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
															if v296 != 0 {
																v323 = F__emscripten_memcpy_bulkmem(m, v318+int32(23), v292+int32(19), v296)
																mBase = m.M
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v311)+22)) = uint8(v310)
															*(*int32)(unsafe.Add(mBase, uint32(v311)+18)) = v309
															v333 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
															v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
															v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
															v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
															mBase = m.M
															v341 = m.ExcPending
															if v341 != 0 {
																return
															} else {
																m.G0 = v16 + int32(16)
																return
															}
														}
													} else {
														v333 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
														v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
														v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
														v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
														mBase = m.M
														v341 = m.ExcPending
														if v341 != 0 {
															return
														} else {
															m.G0 = v16 + int32(16)
															return
														}
													}
												}
											} else {
												v288 = v212
												if v288&int32(12) != 0 {
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+64))
													v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+68))
													v296 = v294 - int32(19)
													v297 = F_ReorderBufferAllocTupleBuf(m, v293, v296)
													mBase = m.M
													v298 = m.ExcPending
													if v298 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v201)+36)) = v297
														v300 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v300
														*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)) = uint16(v300)
														*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v297))) = v294 + int32(4)
														v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+14))
														v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+18)))
														v311 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
														v312 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v312
														*(*int64)(unsafe.Add(mBase, uint32(v311)+15)) = v312
														*(*int64)(unsafe.Add(mBase, uint32(v311))) = v312
														v318 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
														if v296 != 0 {
															v323 = F__emscripten_memcpy_bulkmem(m, v318+int32(23), v292+int32(19), v296)
															mBase = m.M
														} else {
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v311)+22)) = uint8(v310)
														*(*int32)(unsafe.Add(mBase, uint32(v311)+18)) = v309
														v333 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
														v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
														v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
														v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
														mBase = m.M
														v341 = m.ExcPending
														if v341 != 0 {
															return
														} else {
															m.G0 = v16 + int32(16)
															return
														}
													}
												} else {
													v333 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v333)
													v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v336 = *(*int32)(unsafe.Add(mBase, uint32(v181)+96))
													v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+36))
													v338 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v335, v337, v338, v201, int32(0))
													mBase = m.M
													v341 = m.ExcPending
													if v341 != 0 {
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
				v449 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v450 = F_SnapBuildProcessChange(m, v21, v23, v449)
				mBase = m.M
				v451 = m.ExcPending
				if v451 != 0 {
					return
				} else {
					if v450 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v454 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+96))
							v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+64))
							v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
							v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+88))
							if v458 != v460 {
								m.G0 = v16 + int32(16)
								return
							} else {
								v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								if v462 != 0 {
									v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v456)+56)))
									v464 = F_filter_by_origin_cb_wrapper(m, l0, v463)
									mBase = m.M
									v465 = m.ExcPending
									if v465 != 0 {
										return
									} else {
										if v464 != 0 {
											m.G0 = v16 + int32(16)
											return
										} else {
											v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v467 = F_ReorderBufferAllocChange(m, v466)
											mBase = m.M
											v468 = m.ExcPending
											if v468 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v467)+8)) = int32(11)
												v471 = *(*int32)(unsafe.Add(mBase, uint32(v455)+96))
												v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v471)+56)))
												*(*uint16)(unsafe.Add(mBase, uint32(v467)+16)) = uint16(v472)
												v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+8)))
												if v474&int32(1) != 0 {
													v477 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v467)+24)) = uint8(v477)
													v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+8)))
													v480 = v479
												} else {
													v480 = v474
												}
												if v480&int32(2) != 0 {
													v483 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v467)+25)) = uint8(v483)
												} else {
												}
												v485 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v467)+20)) = v485
												v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+120))
												v491 = F_MemoryContextAlloc(m, v488, v485<<(uint(int32(2))%32))
												mBase = m.M
												v492 = m.ExcPending
												if v492 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v467)+28)) = v491
													v496 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
													v498 = v496 << (uint(int32(2)) % 32)
													if v498 != 0 {
														v499 = F__emscripten_memcpy_bulkmem(m, v491, v457+int32(12), v498)
														mBase = m.M
													} else {
													}
													v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v502 = *(*int32)(unsafe.Add(mBase, uint32(v455)+96))
													v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+36))
													v504 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v501, v503, v504, v467, int32(0))
													mBase = m.M
													v507 = m.ExcPending
													if v507 != 0 {
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
									v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v467 = F_ReorderBufferAllocChange(m, v466)
									mBase = m.M
									v468 = m.ExcPending
									if v468 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v467)+8)) = int32(11)
										v471 = *(*int32)(unsafe.Add(mBase, uint32(v455)+96))
										v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v471)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(v467)+16)) = uint16(v472)
										v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+8)))
										if v474&int32(1) != 0 {
											v477 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v467)+24)) = uint8(v477)
											v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+8)))
											v480 = v479
										} else {
											v480 = v474
										}
										if v480&int32(2) != 0 {
											v483 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v467)+25)) = uint8(v483)
										} else {
										}
										v485 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v467)+20)) = v485
										v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+120))
										v491 = F_MemoryContextAlloc(m, v488, v485<<(uint(int32(2))%32))
										mBase = m.M
										v492 = m.ExcPending
										if v492 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v467)+28)) = v491
											v496 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
											v498 = v496 << (uint(int32(2)) % 32)
											if v498 != 0 {
												v499 = F__emscripten_memcpy_bulkmem(m, v491, v457+int32(12), v498)
												mBase = m.M
											} else {
											}
											v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v502 = *(*int32)(unsafe.Add(mBase, uint32(v455)+96))
											v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+36))
											v504 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											F_ReorderBufferQueueChange(m, v501, v503, v504, v467, int32(0))
											mBase = m.M
											v507 = m.ExcPending
											if v507 != 0 {
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
				v510 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v511 = F_SnapBuildProcessChange(m, v21, v23, v510)
				mBase = m.M
				v512 = m.ExcPending
				if v512 != 0 {
					return
				} else {
					if v511 == int32(0) {
						m.G0 = v16 + int32(16)
						return
					} else {
						v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v515 != 0 {
							m.G0 = v16 + int32(16)
							return
						} else {
							v516 = m.G0
							v518 = v516 - int32(16)
							m.G0 = v518
							v520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v521 = int32(0)
							F_XLogRecGetBlockTag(m, v520, v521, v518+int32(4), v521, v521)
							mBase = m.M
							v527 = m.ExcPending
							if v527 != 0 {
								return
							} else {
								v528 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
								v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+88))
								if v528 != v530 {
									m.G0 = v518 + int32(16)
									m.G0 = v16 + int32(16)
									return
								} else {
									v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v532 != 0 {
										v533 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
										v534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v533)+56)))
										v535 = F_filter_by_origin_cb_wrapper(m, l0, v534)
										mBase = m.M
										v536 = m.ExcPending
										if v536 != 0 {
											return
										} else {
											if v535 != 0 {
												m.G0 = v518 + int32(16)
												m.G0 = v16 + int32(16)
												return
											} else {
												v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v538 = F_ReorderBufferAllocChange(m, v537)
												mBase = m.M
												v539 = m.ExcPending
												if v539 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v538)+8)) = int32(9)
													v542 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
													v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v538)+16)) = uint16(v543)
													v545 = *(*int64)(unsafe.Add(mBase, uint32(v518)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v538)+20)) = v545
													v547 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v538)+28)) = v547
													v549 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v538)+32)) = uint8(v549)
													v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v552 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
													v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+36))
													v554 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													F_ReorderBufferQueueChange(m, v551, v553, v554, v538, int32(0))
													mBase = m.M
													v557 = m.ExcPending
													if v557 != 0 {
														return
													} else {
														m.G0 = v518 + int32(16)
														m.G0 = v16 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v538 = F_ReorderBufferAllocChange(m, v537)
										mBase = m.M
										v539 = m.ExcPending
										if v539 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v538)+8)) = int32(9)
											v542 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
											v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v538)+16)) = uint16(v543)
											v545 = *(*int64)(unsafe.Add(mBase, uint32(v518)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v538)+20)) = v545
											v547 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v538)+28)) = v547
											v549 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v538)+32)) = uint8(v549)
											v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v552 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
											v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+36))
											v554 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											F_ReorderBufferQueueChange(m, v551, v553, v554, v538, int32(0))
											mBase = m.M
											v557 = m.ExcPending
											if v557 != 0 {
												return
											} else {
												m.G0 = v518 + int32(16)
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
															v95 = v88 + int32(76)
															v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
															if v96 != int32(1) {
																v111 = v84
																v114 = v111
															} else {
																v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+43)))
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
																		v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+48)))
																		*(*int32)(unsafe.Add(mBase, uint32(v86))) = v107
																	} else {
																	}
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+44))
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
															v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															v144 = int32(5)
															v147 = v122 - v144
															if v147 != 0 {
																v148 = F__emscripten_memcpy_bulkmem(m, v141+int32(23), v114+v144, v147)
																mBase = m.M
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v134)+22)) = uint8(v133)
															*(*int32)(unsafe.Add(mBase, uint32(v134)+18)) = v132
															v152 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v69)+32)) = uint8(v152)
															v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
															v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
															v157 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
															F_ReorderBufferQueueChange(m, v154, v156, v157, v69, int32(base.Ui32(v158&int32(16))>>(uint(int32(4))%32)))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
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
													v95 = v88 + int32(76)
													v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
													if v96 != int32(1) {
														v111 = v84
														v114 = v111
													} else {
														v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+43)))
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
																v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+48)))
																*(*int32)(unsafe.Add(mBase, uint32(v86))) = v107
															} else {
															}
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+44))
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
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
													v144 = int32(5)
													v147 = v122 - v144
													if v147 != 0 {
														v148 = F__emscripten_memcpy_bulkmem(m, v141+int32(23), v114+v144, v147)
														mBase = m.M
													} else {
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v134)+22)) = uint8(v133)
													*(*int32)(unsafe.Add(mBase, uint32(v134)+18)) = v132
													v152 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v69)+32)) = uint8(v152)
													v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
													v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
													v157 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
													F_ReorderBufferQueueChange(m, v154, v156, v157, v69, int32(base.Ui32(v158&int32(16))>>(uint(int32(4))%32)))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v7 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = base.B2i32(l5 != v7) << (uint(int32(7)) % 32)
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = l5 - int32(1)
	goto L3
L2:
	;
	v24 = v7
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v24
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v28 = v26 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v28)
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = v7
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v11 + int32(16)
	return
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v49 = v11 + int32(8)
	goto L11
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1+v37<<(uint(int32(2))%32))))
	v59 = v57
	goto L14
L13:
	;
	v59 = int32(0)
	goto L14
L14:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v37))))
	v63 = v61
	goto L17
L16:
	;
	v63 = int32(1)
	goto L17
L17:
	;
	F_fill_val(m, l0+int32(20)+v37<<(uint(int32(4))%32), v49, v11+int32(4), v11+int32(12), l4, v59, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v67 = v37 + int32(1)
	if v67 != v15 {
		v37 = v67
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L8
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
		v23 = int32(1)
		v24 = l1 - v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
		if v26&v23 == v21 {
			v35 = l2 + v24<<(uint(int32(4))%32) + int32(20)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 < int32(0) {
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
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v41 = v25 + v39 + v36
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+6)))
				if v42 != int32(1) {
					v87 = v41
					m.G0 = v10 + int32(16)
					return v87
				} else {
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+4)))
					switch v45&int32(65535) - int32(1) {
					case 0:
						v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41))))
						v87 = v50
						m.G0 = v10 + int32(16)
						return v87
					case 1:
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
						v87 = v51
						m.G0 = v10 + int32(16)
						return v87
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v45
							F_errmsg_internal(m, int32(482718), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(326160), int32(70), int32(67716))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						v87 = v52
						m.G0 = v10 + int32(16)
						return v87
					}
				}
			}
		} else {
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(base.Ui32(v24)>>(uint(int32(3))%32)))+23)))
			if int32(base.Ui32(v69)>>(uint(v24&int32(7))%32))&int32(1) != 0 {
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v112 int64
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int64
	_ = v175
	var v180 int32
	_ = v180
	var v211 int32
	_ = v211
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v532 int32
	_ = v532
	var v546 int64
	_ = v546
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v558 int64
	_ = v558
	var v559 int64
	_ = v559
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
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
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
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
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
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
	var v846 int32
	_ = v846
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1202 int32
	_ = v1202
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
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
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1519 int32
	_ = v1519
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1541 int32
	_ = v1541
	var v1549 int32
	_ = v1549
	v3 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(192)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+188)) = v3
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = int32(6)
	v38 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38<<(uint(int32(2))%32))+uint32(_consts[77])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+152)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v63 = v3
	goto L2
L2:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63<<(uint(int32(2))%32))+uint32(_consts[78])))
	if v79 < v45 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v246 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	v91 = v79
	goto L7
L5:
	;
	goto L6
L6:
	;
	v242 = v63 + int32(1)
	if v242 != int32(9) {
		v63 = v242
		goto L2
	} else {
		goto L25
	}
L7:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v46+v91<<(uint(int32(3))%32))))
	if v91 < v79 {
		v180 = v91
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46+v180<<(uint(int32(3))%32)))) = v112
	v211 = v91 + int32(1)
	if v211 != v45 {
		v91 = v211
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v116 = base.I32_rotl(base.I32_wrap_i64(v112), int32(16))
	v121 = base.I32_wrap_i64(int64(base.Ui64(v112)>>(uint(int64(32))%64))) & int32(65535)
	v125 = v91
	goto L11
L11:
	;
	v150 = v125 - v79
	v153 = v46 + v150<<(uint(int32(3))%32)
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+2)))
	v158 = v154<<(uint(int32(16))%32) | v157
	if v116 != v158 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v180 = v150
	goto L9
L13:
	;
	if v169 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if base.Ui32(v158) < base.Ui32(v116) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	v169 = base.B2i32(base.Ui32(v121) < base.Ui32(v164)) - base.B2i32(base.Ui32(v164) < base.Ui32(v121))
	goto L13
L17:
	;
	v163 = int32(-1)
	goto L19
L18:
	;
	v163 = int32(1)
	goto L19
L19:
	;
	v169 = v163
	goto L13
L20:
	;
	v180 = v125
	goto L9
L21:
	;
	goto L22
L22:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	*(*int64)(unsafe.Add(mBase, uint32(v46+v125<<(uint(int32(3))%32)))) = v175
	if v79 <= v150 {
		v125 = v150
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
	v251 = F_palloc(m, v245*int32(6))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v750 = v245
	v754 = v3
	goto L28
L28:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L120
L29:
	;
	return int32(0)
L30:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v255 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_pg_qsort(m, v251, v477, int32(6), int32(141))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L29
	} else {
		goto L73
	}
L32:
	;
	v258 = int32(0)
	v476 = v258
	v477 = v258
	goto L31
L33:
	;
	goto L34
L34:
	;
	v264 = int32(0)
	v270 = v264
	v271 = v264
	v275 = int32(-1)
	goto L35
L35:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v299 = v296 + v271<<(uint(int32(3))%32)
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v299)+6)))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v300*int32(6))+3)))
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+2)))
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299))))
	v309 = v305 | v306<<(uint(int32(16))%32)
	if v309 != v275 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v327 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L37:
	;
	if v304&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v313 = v251 + v270*int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v313)+4)) = uint16(v271)
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = int32(65536)
	v327 = v270 + int32(1)
	v328 = v309
	goto L37
L39:
	;
	goto L40
L40:
	;
	v321 = v251 - int32(4) + v270*int32(6)
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321))))
	v324 = v322 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v321))) = uint16(v324)
	v327 = v270
	v328 = v275
	goto L37
L41:
	;
	v333 = v251 - int32(6) + v327*int32(6)
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333))))
	v336 = v334 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v333))) = uint16(v336)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v340 = v271 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v340 < v341 {
		v270 = v327
		v271 = v340
		v275 = v328
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v476 = int32(0)
	v477 = v327
	goto L31
L46:
	;
	goto L47
L47:
	;
	v346 = int32(1)
	v348 = int32(0)
	if v327 != v346 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v358 = v348
	v365 = int32(0)
	goto L51
L49:
	;
	v429 = v348
	goto L50
L50:
	;
	v453 = int32(1)
	if v327&v346 == int32(0) {
		v476 = v453
		v477 = v327
		goto L31
	} else {
		goto L66
	}
L51:
	;
	v382 = int32(4)
	v386 = v251 + v358*int32(6)
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v386))))
	if int32(5) <= v387 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v429 = v421
	goto L50
L53:
	;
	v390 = int32(1)
	if v387&(v387-v390) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v399 = v382
	goto L55
L55:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v386))) = uint16(v399)
	v405 = v251 + (v358|int32(1))*int32(6)
	v406 = int32(*(*int16)(unsafe.Add(mBase, uint32(v405))))
	if int32(5) <= v406 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v398 = v390 << (uint(int32(32)-base.I32_clz(v387)) % 32)
	goto L58
L57:
	;
	v398 = v387
	goto L58
L58:
	;
	v399 = v398
	goto L55
L59:
	;
	v409 = int32(1)
	if v406&(v406-v409) != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v418 = v382
	goto L61
L61:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v405))) = uint16(v418)
	v420 = int32(2)
	v421 = v358 + v420
	v423 = v365 + v420
	if v423 != v327&int32(2147483646) {
		v358 = v421
		v365 = v423
		goto L51
	} else {
		goto L65
	}
L62:
	;
	v417 = v409 << (uint(int32(32)-base.I32_clz(v406)) % 32)
	goto L64
L63:
	;
	v417 = v406
	goto L64
L64:
	;
	v418 = v417
	goto L61
L65:
	;
	goto L52
L66:
	;
	v459 = v251 + v429*int32(6)
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v459))))
	if int32(5) <= v460 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v463 = int32(1)
	if v460&(v460-v463) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v472 = int32(4)
	goto L69
L69:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v459))) = uint16(v472)
	v476 = v453
	v477 = v327
	goto L31
L70:
	;
	v471 = v463 << (uint(int32(32)-base.I32_clz(v460)) % 32)
	goto L72
L71:
	;
	v471 = v460
	goto L72
L72:
	;
	v472 = v471
	goto L69
L73:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v509 = F_palloc(m, v506<<(uint(int32(3))%32))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L29
	} else {
		goto L74
	}
L74:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v476 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v733 = v706 << (uint(int32(3)) % 32)
	if v733 != 0 {
		goto L114
	} else {
		goto L115
	}
L76:
	;
	v706 = int32(0)
	v712 = v511
	v716 = v3
	goto L75
L77:
	;
	goto L78
L78:
	;
	v515 = int32(6)
	if v515 <= v477 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v518 = v515
	goto L81
L80:
	;
	v518 = v477
	goto L81
L81:
	;
	v532 = v3
	v546 = int64(-1)
	goto L83
L82:
	;
	v573 = int32(1)
	if v477 == v573 {
		goto L92
	} else {
		goto L93
	}
L83:
	;
	v551 = int32(*(*int16)(unsafe.Add(mBase, uint32(v251+v532*int32(6))+4)))
	v554 = v511 + v551<<(uint(int32(3))%32)
	v555 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v554))))
	v558 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v554)+2)))
	v559 = v555<<(uint(int64(16))%64) | v558
	if v546 != int64(-1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v571 = v518
	goto L82
L85:
	;
	if v559 < v546-int64(3) {
		v571 = v532
		goto L82
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v569 = v532 + int32(1)
	if v569 != v518 {
		v532 = v569
		v546 = v559
		goto L83
	} else {
		goto L90
	}
L88:
	;
	if v546+int64(3) < v559 {
		v571 = v532
		goto L82
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	goto L84
L91:
	;
	if v518&v573 != 0 {
		goto L106
	} else {
		goto L107
	}
L92:
	;
	v577 = int32(0)
	v657 = v577
	v659 = v577
	goto L91
L93:
	;
	goto L94
L94:
	;
	v581 = int32(0)
	v586 = v581
	v588 = v581
	v592 = v581
	goto L95
L95:
	;
	v612 = int32(3)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v618 = v251 + v588*int32(6)
	v619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v618)+4)))
	v623 = int32(*(*int16)(unsafe.Add(mBase, uint32(v618)+2)))
	v625 = v623 << (uint(v612) % 32)
	if v625 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v657 = v649
	v659 = v651
	goto L91
L97:
	;
	v628 = int32(*(*int16)(unsafe.Add(mBase, uint32(v618)+2)))
	v629 = v586 + v628
	v630 = int32(3)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v638 = v251 + (v588|int32(1))*int32(6)
	v639 = int32(*(*int16)(unsafe.Add(mBase, uint32(v638)+4)))
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v638)+2)))
	v645 = v643 << (uint(v630) % 32)
	if v645 != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v626 = F__emscripten_memcpy_bulkmem(m, v509+v586<<(uint(v612)%32), v615+v619<<(uint(v612)%32), v625)
	mBase = m.M
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L97
L101:
	;
	v648 = int32(*(*int16)(unsafe.Add(mBase, uint32(v638)+2)))
	v649 = v629 + v648
	v650 = int32(2)
	v651 = v588 + v650
	v653 = v592 + v650
	if v653 != v518&int32(-2) {
		v586 = v649
		v588 = v651
		v592 = v653
		goto L95
	} else {
		goto L105
	}
L102:
	;
	v646 = F__emscripten_memcpy_bulkmem(m, v509+v629<<(uint(v630)%32), v633+v639<<(uint(v630)%32), v645)
	mBase = m.M
	goto L104
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	goto L96
L106:
	;
	v683 = int32(3)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v689 = v251 + v659*int32(6)
	v690 = int32(*(*int16)(unsafe.Add(mBase, uint32(v689)+4)))
	v694 = int32(*(*int16)(unsafe.Add(mBase, uint32(v689)+2)))
	v696 = v694 << (uint(v683) % 32)
	if v696 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v701 = v657
	goto L108
L108:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v706 = v701
	v712 = v703
	v716 = v571
	goto L75
L109:
	;
	v699 = int32(*(*int16)(unsafe.Add(mBase, uint32(v689)+2)))
	v701 = v657 + v699
	goto L108
L110:
	;
	v697 = F__emscripten_memcpy_bulkmem(m, v509+v657<<(uint(v683)%32), v686+v690<<(uint(v683)%32), v696)
	mBase = m.M
	goto L112
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v706
	F_pfree(m, v509)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L29
	} else {
		goto L117
	}
L114:
	;
	v734 = F__emscripten_memcpy_bulkmem(m, v712, v509, v733)
	mBase = m.M
	goto L116
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	F_pfree(m, v251)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L29
	} else {
		goto L118
	}
L118:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v750 = v741
	v754 = v716
	goto L28
L119:
	;
	v781 = int32(0)
	if v750 <= v781 {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	if base.Ui32(v771) < base.Ui32(int32(12000)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v780 = v775
	goto L119
L122:
	;
	goto L123
L123:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)+92))
	v778 = F_get_tablespace_maintenance_io_concurrency(m, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L29
	} else {
		goto L124
	}
L124:
	;
	v780 = v778
	goto L119
L125:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v880 {
		goto L146
	} else {
		goto L147
	}
L126:
	;
	v854 = int32(-1)
	v856 = int32(0)
	goto L125
L127:
	;
	if v780 < v754 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v786 = v780
	goto L130
L129:
	;
	v786 = v754
	goto L130
L130:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v787 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v788 = v786
	goto L133
L132:
	;
	v788 = v780
	goto L133
L133:
	;
	if v788 <= int32(0) {
		goto L126
	} else {
		goto L134
	}
L134:
	;
	v791 = int32(0)
	v796 = int32(-1)
	v797 = v791
	v799 = v791
	goto L135
L135:
	;
	v823 = v770 + v797<<(uint(int32(3))%32)
	v824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823))))
	if v796 == int32(-1) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v854 = v843
	v856 = v846
	goto L125
L137:
	;
	v846 = v797 + int32(1)
	if v750 <= v846 {
		v854 = v843
		v856 = v846
		goto L125
	} else {
		goto L144
	}
L138:
	;
	F_PrefetchBuffer(m, v31+int32(84), l0, v836)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L29
	} else {
		goto L143
	}
L139:
	;
	v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+2)))
	v836 = v827 | v824<<(uint(int32(16))%32)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+2)))
	v834 = v831 | v824<<(uint(int32(16))%32)
	if v834 != v796 {
		v836 = v834
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v843 = v796
	v844 = v799
	goto L137
L143:
	;
	v843 = v836
	v844 = v799 + int32(1)
	goto L137
L144:
	;
	if v844 < v788 {
		v796 = v843
		v797 = v846
		v799 = v844
		goto L135
	} else {
		goto L145
	}
L145:
	;
	goto L136
L146:
	;
	v886 = v854
	v888 = v856
	v896 = v754
	v897 = v3
	v899 = v3
	v901 = v3
	v902 = v3
	v903 = v3
	v904 = v35
	v905 = v3
	v906 = int32(-1)
	v907 = v3
	v908 = v3
	v909 = v3
	goto L149
L147:
	;
	v1530 = v781
	v1533 = v3
	v1541 = v3
	goto L148
L148:
	;
	F_UnlockReleaseBuffer(m, v1533)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L29
	} else {
		goto L256
	}
L149:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v915 = v912 + v899<<(uint(int32(3))%32)
	v916 = int32(*(*int16)(unsafe.Add(mBase, uint32(v915)+6)))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v906 != int32(-1) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v31)+188))
	v1530 = v1519
	v1533 = v1504
	v1541 = v1512
	goto L148
L151:
	;
	goto L150
L152:
	;
	v1104 = v917 + v916*int32(6)
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+4)))
	v1107 = v1099 & int32(65535)
	if base.Ui32(v1105) <= base.Ui32(v1107) {
		goto L196
	} else {
		goto L197
	}
L153:
	;
	v920 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+2)))
	v921 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915))))
	if v920|v921<<(uint(int32(16))%32) == v906 {
		v1076 = v886
		v1078 = v888
		v1086 = v896
		v1087 = v897
		v1092 = v902
		v1094 = v904
		v1096 = v906
		v1097 = v907
		v1098 = v908
		v1099 = v909
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v926 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	if v901&int32(1) != 0 {
		v1504 = v897
		v1512 = v905
		goto L151
	} else {
		goto L160
	}
L158:
	;
	v943 = v896
	v944 = v904
	v945 = v907
	goto L159
L159:
	;
	if v897 != 0 {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	if base.B2i32(v903 == v907)&base.B2i32(int32(0) < v908) != 0 {
		v1504 = v897
		v1512 = v905
		goto L151
	} else {
		goto L161
	}
L161:
	;
	if int32(0) < v896 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v943 = v941
	v944 = v942
	v945 = v903
	goto L159
L163:
	;
	v941 = v896 - int32(1)
	v942 = v904
	goto L162
L164:
	;
	goto L165
L165:
	;
	v940 = base.I32_div_s(v904, int32(2))
	v941 = v896
	v942 = v940
	goto L162
L166:
	;
	F_UnlockReleaseBuffer(m, v897)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L29
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v949 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+2)))
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915))))
	v953 = v949 | v950<<(uint(int32(16))%32)
	v954 = F_ReadBuffer(m, l0, v953)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L29
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	if v750 <= v888 {
		v1017 = v886
		v1018 = v888
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_LockBuffer(m, v954, int32(1))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L29
	} else {
		goto L184
	}
L172:
	;
	v959 = v886
	v961 = v888
	v963 = int32(0)
	goto L173
L173:
	;
	v987 = v770 + v961<<(uint(int32(3))%32)
	v988 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v987))))
	if v959 == int32(-1) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v1017 = v1007
	v1018 = v1010
	goto L171
L175:
	;
	v1010 = v961 + int32(1)
	if v750 <= v1010 {
		v1017 = v1007
		v1018 = v1010
		goto L171
	} else {
		goto L182
	}
L176:
	;
	F_PrefetchBuffer(m, v31+int32(84), l0, v1000)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L29
	} else {
		goto L181
	}
L177:
	;
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v987)+2)))
	v1000 = v991 | v988<<(uint(int32(16))%32)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v987)+2)))
	v998 = v995 | v988<<(uint(int32(16))%32)
	if v998 != v959 {
		v1000 = v998
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v1007 = v959
	v1008 = v963
	goto L175
L181:
	;
	v1007 = v1000
	v1008 = v963 + int32(1)
	goto L175
L182:
	;
	if v1008 <= int32(0) {
		v959 = v1007
		v961 = v1010
		v963 = v1008
		goto L173
	} else {
		goto L183
	}
L183:
	;
	goto L174
L184:
	;
	if v954 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1065) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1050+(v954^int32(-1))<<(uint(int32(2))%32))))
	v1064 = v1056
	goto L185
L187:
	;
	goto L188
L188:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1064 = v1058 + v954<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	v1073 = int32(base.Ui32(v1065+int32(262120)) >> (uint(int32(2)) % 32))
	goto L191
L190:
	;
	v1073 = int32(0)
	goto L191
L191:
	;
	v1076 = v1017
	v1078 = v1018
	v1086 = v943
	v1087 = v954
	v1092 = v1064
	v1094 = v944
	v1096 = v953
	v1097 = v945
	v1098 = v908 + int32(1)
	v1099 = v1073
	goto L152
L192:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1476 < v1489 {
		v886 = v1076
		v888 = v1078
		v896 = v1086
		v897 = v1087
		v899 = v1476
		v901 = v1478
		v902 = v1092
		v903 = v1480
		v904 = v1094
		v905 = v1482
		v906 = v1096
		v907 = v1097
		v908 = v1098
		v909 = v1099
		goto L149
	} else {
		goto L255
	}
L193:
	;
	v1476 = v899 + int32(1)
	v1478 = v901
	v1480 = v903
	v1482 = v905
	goto L192
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L29
	} else {
		goto L251
	}
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L29
	} else {
		goto L247
	}
L196:
	;
	v1110 = v1092 + int32(24)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1110+v1105<<(uint(int32(2))%32)-int32(4))))
	if v1116&int32(98304) == int32(0) {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L29
	} else {
		goto L243
	}
L199:
	;
	if base.Ui32(int32(131072)) <= base.Ui32(v1116) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1092+v1116&int32(32767))+18)))
	if v1126 < int32(0) {
		goto L194
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104)+2)))
	if v1129 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L202
L204:
	;
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+108)) = uint16(v1132)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1134
	v1144 = F_heap_hot_search_buffer(m, v31+int32(104), l0, v1087, v31+int32(112), v31+int32(84), int32(0), int32(1))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L29
	} else {
		goto L207
	}
L205:
	;
	v1158 = v1105
	v1159 = v901
	v1160 = v903
	goto L206
L206:
	;
	if base.Ui32(v1107) <= base.Ui32((v1158-int32(1))&int32(65535)) {
		goto L212
	} else {
		goto L213
	}
L207:
	;
	if v1144 != 0 {
		goto L193
	} else {
		goto L208
	}
L208:
	;
	v1146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1104)+2)) = uint8(v1146)
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+4)))
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v1149 == v1146 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1104)+4)))
	v1153 = v903 + v1152
	v1156 = v1153
	v1157 = base.B2i32(v1094 <= v1153) | v901
	goto L211
L210:
	;
	v1156 = v903
	v1157 = v901
	goto L211
L211:
	;
	v1158 = v1148
	v1159 = v1157
	v1160 = v1156
	goto L206
L212:
	;
	v1361 = v899 + int32(1)
	v1476 = v1361
	v1478 = v1159
	v1480 = v1160
	v1482 = v1361
	goto L192
L213:
	;
	v1173 = v1158
	v1177 = int32(0)
	goto L214
L214:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1173&int32(65535)<<(uint(int32(2))%32)+v1110-int32(4))))
	switch int32(base.Ui32(v1202)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L217
	case 1:
		goto L218
	default:
		goto L212
	}
L215:
	;
	goto L212
L216:
	;
	if base.Ui32((v1305-int32(1))&int32(65535)) < base.Ui32(v1107) {
		v1173 = v1305
		v1177 = v1309
		goto L214
	} else {
		goto L242
	}
L217:
	;
	v1213 = v1092 + v1202&int32(32767)
	if v1177 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1305 = v1202 & int32(32767)
	v1309 = v1177
	goto L216
L219:
	;
	v1215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213)+20)))
	v1216 = int32(768)
	if v1215&v1216 != v1216 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	goto L221
L221:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1213, v31+int32(188))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L29
	} else {
		goto L226
	}
L222:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	v1221 = v1220
	goto L224
L223:
	;
	v1221 = int32(2)
	goto L224
L224:
	;
	if v1221 != v1177 {
		goto L212
	} else {
		goto L225
	}
L225:
	;
	goto L221
L226:
	;
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1213)+19)))
	if v1228&int32(64) == int32(0) {
		goto L212
	} else {
		goto L227
	}
L227:
	;
	v1233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213)+20)))
	if v1233&int32(2048) != 0 {
		goto L212
	} else {
		goto L228
	}
L228:
	;
	if v1233&int32(768) == int32(512) {
		goto L212
	} else {
		goto L229
	}
L229:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+4))
	v1241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213)+16)))
	if v1233&int32(4224) != int32(4096) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1305 = v1241
	v1309 = v1240
	goto L216
L231:
	;
	goto L232
L232:
	;
	v1246 = int32(0)
	v1250 = F_GetMultiXactIdMembers(m, v1240, v31+int32(84), v1246)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L29
	} else {
		goto L233
	}
L233:
	;
	if v1250 <= int32(0) {
		v1305 = v1241
		v1309 = v1246
		goto L216
	} else {
		goto L234
	}
L234:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
	v1259 = int32(0)
	goto L237
L235:
	;
	F_pfree(m, v1255)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L29
	} else {
		goto L241
	}
L236:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1286)))
	v1296 = v1294
	goto L235
L237:
	;
	v1286 = v1255 + v1259<<(uint(int32(3))%32)
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1287) {
		goto L236
	} else {
		goto L239
	}
L238:
	;
	v1296 = int32(0)
	goto L235
L239:
	;
	v1291 = v1259 + int32(1)
	if v1291 != v1250 {
		v1259 = v1291
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	v1305 = v1241
	v1309 = v1296
	goto L216
L242:
	;
	goto L215
L243:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L29
	} else {
		goto L244
	}
L244:
	;
	v1369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+2)))
	v1370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915))))
	v1371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1104))))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1374 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v1372
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1369 | v1370<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(677701), v31)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L29
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(496630), int32(8083), int32(432446))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L29
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L29
	} else {
		goto L248
	}
L248:
	;
	v1400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+2)))
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915))))
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1104))))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1405 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1400 | v1401<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(677822), v31+int32(32))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L29
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(496630), int32(8093), int32(432446))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L29
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L29
	} else {
		goto L252
	}
L252:
	;
	v1433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915)+2)))
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915))))
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1104))))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1438 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+76)) = v1436
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v1435
	*(*int32)(unsafe.Add(mBase, uint32(v31)+68)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1433 | v1434<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(677927), v31-int32(-64))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L29
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(496630), int32(8109), int32(432446))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L29
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v1504 = v1087
	v1512 = v1482
	goto L151
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v1541
	m.G0 = v31 + int32(192)
	return v1530
}
func F_heap_lock_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v534 int32
	_ = v534
	var v582 int32
	_ = v582
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v664 int32
	_ = v664
	var v675 int32
	_ = v675
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v920 int32
	_ = v920
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
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
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1335 int32
	_ = v1335
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1569 int64
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	v9 = int32(0)
	v28 = m.G0
	v30 = v28 + int32(-64)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v9
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v39 = F_ReadBuffer(m, l0, v34|v35<<(uint(int32(16))%32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v39
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v48 = v44 | v45<<(uint(int32(16))%32)
	if v39 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+10)))
	if v67&int32(4) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+(v39^int32(-1))<<(uint(int32(2))%32))))
	v66 = v58
	goto L3
L5:
	;
	goto L6
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v66 = v60 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	F_visibilitymap_pin(m, l0, v48, v28+int32(-4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v75 = v39
	goto L9
L9:
	;
	F_LockBuffer(m, v75, int32(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v75 = v74
	goto L9
L11:
	;
	v80 = l1 + int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v81 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v105 = v99 + v100<<(uint(int32(2))%32) + int32(20)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v106&int32(32767) + v99
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v111) >> (uint(int32(17)) % 32))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v115
	v117 = int32(1)
	v120 = l3 * int32(12)
	v124 = v120 + int32(742144)
	v137 = v117
	v144 = v9
	v147 = v9
	goto L18
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v81^int32(-1))<<(uint(int32(2))%32))))
	v99 = v91
	goto L12
L14:
	;
	goto L15
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v99 = v93 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+60))
	if v1641 != 0 {
		goto L368
	} else {
		goto L369
	}
L17:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1610, int32(0))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L367
	}
L18:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v156 = F_HeapTupleSatisfiesUpdate(m, l1, l2, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1434)+20)))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+4))
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L344
	}
L20:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	if v1416 != 0 {
		goto L338
	} else {
		goto L339
	}
L21:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1307
	v1309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1306)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)) = uint16(v1309)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	v1313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+20)))
	if v1313&int32(6272) != int32(4096) {
		goto L319
	} else {
		goto L320
	}
L22:
	;
	if v1263 == int32(0) {
		goto L20
	} else {
		goto L317
	}
L23:
	;
	v158 = int32(3)
	v159 = v156 - v158
	if base.Ui32(v158) <= base.Ui32(v159) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v156 == int32(1) {
		v1583 = l0
		v1593 = v30
		v1596 = v156
		v1597 = v80
		v1599 = v144
		v1601 = v124
		goto L17
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+18)))
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+20)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-16)))) = uint16(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v172, int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v1259 = v137
	v1263 = v156
	v1266 = v144
	v1269 = v147
	goto L22
L28:
	;
	if v137&int32(1) == int32(0) {
		v582 = v147
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v137 = int32(0)
	v144 = v1238
	v147 = v1241
	goto L18
L30:
	;
	v1259 = int32(0)
	v1263 = v1207
	v1266 = v1210
	v1269 = v1213
	goto L22
L31:
	;
	v808 = v166 & int32(4096)
	if v808 != 0 {
		goto L178
	} else {
		goto L179
	}
L32:
	;
	if v166&int32(80) != int32(16) {
		v799 = v755
		goto L31
	} else {
		goto L174
	}
L33:
	;
	if v166&int32(4096) == int32(0) {
		v755 = v582
		goto L32
	} else {
		goto L168
	}
L34:
	;
	if base.B2i32(v166&int32(128) == int32(0))&base.B2i32(v166&int32(4176) != int32(64)) != 0 {
		v799 = v675
		goto L31
	} else {
		goto L163
	}
L35:
	;
	switch l3 {
	case 0:
		goto L144
	case 1:
		goto L143
	case 2:
		goto L33
	default:
		v799 = v582
		goto L31
	}
L36:
	;
	if v166&int32(4096) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v1614 = l0
	v1624 = v30
	v1627 = int32(0)
	v1628 = v80
	v1630 = v144
	v1632 = v124
	goto L16
L38:
	;
	F_pfree(m, v351)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L142
	}
L39:
	;
	v194 = F_GetMultiXactIdMembers(m, v167, v30+int32(56), int32(base.Ui32(v166&int32(128))>>(uint(int32(7))%32))|base.B2i32(v166&int32(4176) == int32(64)))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v167) < base.Ui32(int32(3)) {
		goto L96
	} else {
		goto L97
	}
L42:
	;
	if int32(0) < v194 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v206 = int32(0)
	v217 = v147
	goto L46
L44:
	;
	v384 = v147
	goto L45
L45:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	if v392 == int32(0) {
		v582 = v384
		goto L35
	} else {
		goto L93
	}
L46:
	;
	v225 = int32(3)
	v226 = v206 << (uint(v225) % 32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v226+v227)))
	if base.Ui32(v229) < base.Ui32(v225) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v384 = v361
	goto L45
L48:
	;
	if v349 != 0 {
		goto L88
	} else {
		goto L89
	}
L49:
	;
	v349 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v240 == v229 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v349 = int32(1)
	goto L48
L53:
	;
	goto L54
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v244 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v349 = v341
	goto L48
L56:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v248 == int32(0) {
		v341 = int32(0)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v312 = int32(0)
	v314 = v244 - int32(1)
	goto L78
L59:
	;
	v253 = v248
	goto L60
L60:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v258 == int32(4) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v341 = int32(0)
	goto L55
L62:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v253)+80))
	if v305 != 0 {
		v253 = v305
		goto L60
	} else {
		goto L77
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v261 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v264 = int32(1)
	if v229 == v261 {
		v341 = v264
		goto L55
	} else {
		goto L65
	}
L65:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v253)+52))
	v268 = v266 - int32(1)
	if v268 < int32(0) {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v273 = int32(0)
	v275 = v268
	goto L67
L67:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v253)+48))
	v281 = int32(2)
	v282 = base.I32_div_s(v275-v273, v281)
	v283 = v282 + v273
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v279+v283<<(uint(v281)%32))))
	if v287 == v229 {
		v341 = v264
		goto L55
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v291 = F_TransactionIdPrecedes(m, v287, v229)
	mBase = m.M
	if v291 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v292 = v283 + int32(1)
	goto L72
L71:
	;
	v292 = v273
	goto L72
L72:
	;
	if v291 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v295 = v275
	goto L75
L74:
	;
	v295 = v283 - int32(1)
	goto L75
L75:
	;
	if v292 <= v295 {
		v273 = v292
		v275 = v295
		goto L67
	} else {
		goto L76
	}
L76:
	;
	goto L68
L77:
	;
	goto L61
L78:
	;
	v319 = int32(2)
	v320 = base.I32_div_s(v314-v312, v319)
	v321 = v320 + v312
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v310+v321<<(uint(v319)%32))))
	v326 = base.B2i32(v325 == v229)
	if v325 == v229 {
		v341 = v326
		goto L55
	} else {
		goto L80
	}
L79:
	;
	v341 = v326
	goto L55
L80:
	;
	v329 = base.B2i32(base.Ui32(v325) < base.Ui32(v229))
	if base.Ui32(v325) < base.Ui32(v229) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v330 = v321 + int32(1)
	goto L83
L82:
	;
	v330 = v312
	goto L83
L83:
	;
	if base.Ui32(v325) < base.Ui32(v229) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v333 = v314
	goto L86
L85:
	;
	v333 = v321 - int32(1)
	goto L86
L86:
	;
	if v330 <= v333 {
		v312 = v330
		v314 = v333
		goto L78
	} else {
		goto L87
	}
L87:
	;
	goto L79
L88:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v226+v351)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353<<(uint(int32(2))%32))+uint32(_consts[72])))
	if base.Ui32(l3) <= base.Ui32(v358) {
		goto L38
	} else {
		goto L91
	}
L89:
	;
	v361 = v217
	goto L90
L90:
	;
	v363 = v206 + int32(1)
	if v363 != v194 {
		v206 = v363
		v217 = v361
		goto L46
	} else {
		goto L92
	}
L91:
	;
	v361 = int32(1)
	goto L90
L92:
	;
	goto L47
L93:
	;
	F_pfree(m, v392)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v582 = v384
	goto L35
L95:
	;
	if v516 == int32(0) {
		v582 = v147
		goto L35
	} else {
		goto L135
	}
L96:
	;
	v516 = int32(0)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v407 == v167 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v516 = int32(1)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v411 <= int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v516 = v508
	goto L95
L103:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v415 == int32(0) {
		v508 = int32(0)
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v479 = int32(0)
	v481 = v411 - int32(1)
	goto L125
L106:
	;
	v420 = v415
	goto L107
L107:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420)+20))
	if v425 == int32(4) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v508 = int32(0)
	goto L102
L109:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v420)+80))
	if v472 != 0 {
		v420 = v472
		goto L107
	} else {
		goto L124
	}
L110:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	if v428 == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v431 = int32(1)
	if v167 == v428 {
		v508 = v431
		goto L102
	} else {
		goto L112
	}
L112:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v420)+52))
	v435 = v433 - int32(1)
	if v435 < int32(0) {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v440 = int32(0)
	v442 = v435
	goto L114
L114:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v448 = int32(2)
	v449 = base.I32_div_s(v442-v440, v448)
	v450 = v449 + v440
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v446+v450<<(uint(v448)%32))))
	if v454 == v167 {
		v508 = v431
		goto L102
	} else {
		goto L116
	}
L115:
	;
	goto L109
L116:
	;
	v458 = F_TransactionIdPrecedes(m, v454, v167)
	mBase = m.M
	if v458 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v459 = v450 + int32(1)
	goto L119
L118:
	;
	v459 = v440
	goto L119
L119:
	;
	if v458 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v462 = v442
	goto L122
L121:
	;
	v462 = v450 - int32(1)
	goto L122
L122:
	;
	if v459 <= v462 {
		v440 = v459
		v442 = v462
		goto L114
	} else {
		goto L123
	}
L123:
	;
	goto L115
L124:
	;
	goto L108
L125:
	;
	v486 = int32(2)
	v487 = base.I32_div_s(v481-v479, v486)
	v488 = v487 + v479
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v477+v488<<(uint(v486)%32))))
	v493 = base.B2i32(v492 == v167)
	if v492 == v167 {
		v508 = v493
		goto L102
	} else {
		goto L127
	}
L126:
	;
	v508 = v493
	goto L102
L127:
	;
	v496 = base.B2i32(base.Ui32(v492) < base.Ui32(v167))
	if base.Ui32(v492) < base.Ui32(v167) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v497 = v488 + int32(1)
	goto L130
L129:
	;
	v497 = v479
	goto L130
L130:
	;
	if base.Ui32(v492) < base.Ui32(v167) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v500 = v481
	goto L133
L132:
	;
	v500 = v488 - int32(1)
	goto L133
L133:
	;
	if v497 <= v500 {
		v479 = v497
		v481 = v500
		goto L125
	} else {
		goto L134
	}
L134:
	;
	goto L126
L135:
	;
	switch l3 {
	case 0:
		goto L37
	case 1:
		goto L138
	case 2:
		goto L136
	case 3:
		goto L137
	default:
		v582 = v147
		goto L35
	}
L136:
	;
	if v166&int32(80) == int32(64) {
		goto L37
	} else {
		goto L141
	}
L137:
	;
	if v166&int32(80) != int32(64) {
		v799 = v147
		goto L31
	} else {
		goto L139
	}
L138:
	;
	v520 = v166 & int32(80)
	switch v520 + int32(-64) {
	case 0, 16:
		goto L37
	default:
		v664 = v520
		v675 = v147
		goto L34
	}
L139:
	;
	if v165&int32(8192) != 0 {
		goto L37
	} else {
		goto L140
	}
L140:
	;
	v799 = v147
	goto L31
L141:
	;
	v755 = v147
	goto L32
L142:
	;
	goto L37
L143:
	;
	v664 = v166 & int32(80)
	v675 = v582
	goto L34
L144:
	;
	if v165&int32(8192) != 0 {
		v799 = v582
		goto L31
	} else {
		goto L145
	}
L145:
	;
	v600 = int32(base.Ui32(v166&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v166&int32(4176) == int32(64))
	if (l5^v117|v600)&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v637, int32(2))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L159
	}
L147:
	;
	v605 = v30 + int32(44)
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+2)))
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v608 = int32(16)
	v611 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v605)+2)))
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v605))))
	if v606|v607<<(uint(v608)%32) == v611|v612<<(uint(v608)%32) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	if v622 != 0 {
		goto L146
	} else {
		goto L154
	}
L149:
	;
	goto L148
L150:
	;
	v618 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v605)+4)))
	if v618 == v619 {
		v622 = int32(1)
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v622 = int32(0)
	goto L149
L153:
	;
	goto L152
L154:
	;
	v625 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v628 = F_heap_lock_updated_tuple(m, l0, v166, v167, v30+int32(44), v625, int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v628 == int32(0) {
		goto L146
	} else {
		goto L157
	}
L157:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v632, int32(2))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v1291 = v628
	v1295 = v144
	goto L21
L159:
	;
	v641 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v643 = F_HeapTupleHeaderIsOnlyLocked(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if v643 != 0 {
		v1207 = v641
		v1210 = v144
		v1213 = v582
		goto L30
	} else {
		goto L161
	}
L161:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+19)))
	if (int32(base.Ui32(v646)>>(uint(int32(5))%32))|v600)&int32(1) == int32(0) {
		v1207 = v641
		v1210 = v144
		v1213 = v582
		goto L30
	} else {
		goto L162
	}
L162:
	;
	v1238 = v144
	v1241 = v582
	goto L29
L163:
	;
	if v664 == int32(64) {
		v799 = v675
		goto L31
	} else {
		goto L164
	}
L164:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v694, int32(2))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v698)+20)))
	if v699&int32(80) == int32(64) {
		v1238 = v144
		v1241 = v675
		goto L29
	} else {
		goto L166
	}
L166:
	;
	if base.B2i32(v699&int32(128) == int32(0))&base.B2i32(v699&int32(4176) != int32(64)) != 0 {
		v1238 = v144
		v1241 = v675
		goto L29
	} else {
		goto L167
	}
L167:
	;
	v713 = int32(0)
	v1259 = v713
	v1263 = v713
	v1266 = v144
	v1269 = v675
	goto L22
L168:
	;
	v721 = F_DoesMultiXactIdConflict(m, v167, v166, int32(2), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	if v721 != 0 {
		v799 = v582
		goto L31
	} else {
		goto L170
	}
L170:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v723, int32(2))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+20)))
	if (v728^v166)&int32(4304) != 0 {
		v1238 = v144
		v1241 = v582
		goto L29
	} else {
		goto L172
	}
L172:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	if v732 != v167 {
		v1238 = v144
		v1241 = v582
		goto L29
	} else {
		goto L173
	}
L173:
	;
	v734 = int32(0)
	v1259 = v734
	v1263 = v734
	v1266 = v144
	v1269 = v582
	goto L22
L174:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v767, int32(2))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v771)+20)))
	if (v772^v166)&int32(4304) != 0 {
		v1238 = v144
		v1241 = v755
		goto L29
	} else {
		goto L176
	}
L176:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	if v776 != v167 {
		v1238 = v144
		v1241 = v755
		goto L29
	} else {
		goto L177
	}
L177:
	;
	v778 = int32(0)
	v1259 = v778
	v1263 = v778
	v1266 = v144
	v1269 = v755
	goto L22
L178:
	;
	if base.Ui32(v159) <= base.Ui32(int32(1)) {
		goto L224
	} else {
		goto L225
	}
L179:
	;
	if base.Ui32(v167) < base.Ui32(int32(3)) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v928 == int32(0) {
		goto L178
	} else {
		goto L220
	}
L181:
	;
	v928 = int32(0)
	goto L180
L182:
	;
	goto L183
L183:
	;
	v819 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v819 == v167 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v928 = int32(1)
	goto L180
L185:
	;
	goto L186
L186:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v823 <= int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v928 = v920
	goto L180
L188:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v827 == int32(0) {
		v920 = int32(0)
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v891 = int32(0)
	v893 = v823 - int32(1)
	goto L210
L191:
	;
	v832 = v827
	goto L192
L192:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v832)+20))
	if v837 == int32(4) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v920 = int32(0)
	goto L187
L194:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v832)+80))
	if v884 != 0 {
		v832 = v884
		goto L192
	} else {
		goto L209
	}
L195:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	if v840 == int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v843 = int32(1)
	if v167 == v840 {
		v920 = v843
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v832)+52))
	v847 = v845 - int32(1)
	if v847 < int32(0) {
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v852 = int32(0)
	v854 = v847
	goto L199
L199:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v832)+48))
	v860 = int32(2)
	v861 = base.I32_div_s(v854-v852, v860)
	v862 = v861 + v852
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v858+v862<<(uint(v860)%32))))
	if v866 == v167 {
		v920 = v843
		goto L187
	} else {
		goto L201
	}
L200:
	;
	goto L194
L201:
	;
	v870 = F_TransactionIdPrecedes(m, v866, v167)
	mBase = m.M
	if v870 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v871 = v862 + int32(1)
	goto L204
L203:
	;
	v871 = v852
	goto L204
L204:
	;
	if v870 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v874 = v854
	goto L207
L206:
	;
	v874 = v862 - int32(1)
	goto L207
L207:
	;
	if v871 <= v874 {
		v852 = v871
		v854 = v874
		goto L199
	} else {
		goto L208
	}
L208:
	;
	goto L200
L209:
	;
	goto L193
L210:
	;
	v898 = int32(2)
	v899 = base.I32_div_s(v893-v891, v898)
	v900 = v899 + v891
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v889+v900<<(uint(v898)%32))))
	v905 = base.B2i32(v904 == v167)
	if v904 == v167 {
		v920 = v905
		goto L187
	} else {
		goto L212
	}
L211:
	;
	v920 = v905
	goto L187
L212:
	;
	v908 = base.B2i32(base.Ui32(v904) < base.Ui32(v167))
	if base.Ui32(v904) < base.Ui32(v167) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v909 = v900 + int32(1)
	goto L215
L214:
	;
	v909 = v891
	goto L215
L215:
	;
	if base.Ui32(v904) < base.Ui32(v167) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v912 = v893
	goto L218
L217:
	;
	v912 = v900 - int32(1)
	goto L218
L218:
	;
	if v909 <= v912 {
		v891 = v909
		v893 = v912
		goto L210
	} else {
		goto L219
	}
L219:
	;
	goto L211
L220:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v931, int32(2))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+20)))
	if (v936^v166)&int32(4304) != 0 {
		v1238 = v144
		v1241 = v799
		goto L29
	} else {
		goto L222
	}
L222:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v935)+4))
	if v940 != v167 {
		v1238 = v144
		v1241 = v799
		goto L29
	} else {
		goto L223
	}
L223:
	;
	v942 = int32(0)
	v1259 = v942
	v1263 = v942
	v1266 = v144
	v1269 = v799
	goto L22
L224:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v946, int32(2))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	if (v144|v799)&int32(1) != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v1259 = int32(0)
	v1263 = v156
	v1266 = v144
	v1269 = v799
	goto L22
L228:
	;
	if v808 != 0 {
		goto L247
	} else {
		goto L248
	}
L229:
	;
	v999 = base.B2i32(v799 == int32(0)) | v144
	goto L228
L230:
	;
	goto L231
L231:
	;
	v957 = int32(1)
	switch l4 {
	case 0:
		goto L232
	case 1:
		goto L233
	case 2:
		goto L234
	default:
		v999 = v957
		goto L228
	}
L232:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[73])))
	F_LockTuple(m, l0, v80, v996)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L244
	}
L233:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[73])))
	v986 = F_ConditionalLockTuple(m, l0, v80, v984, int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L241
	}
L234:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[73])))
	v960 = int32(*(*uint8)(unsafe.Add(mBase, _consts[75])))
	v961 = F_ConditionalLockTuple(m, l0, v80, v958, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	if v961 != 0 {
		v999 = v957
		goto L228
	} else {
		goto L236
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v970 + int32(4)
	F_errmsg(m, int32(690369), v30+int32(32))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(496630), int32(5299), int32(315630))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	if v986 != 0 {
		v999 = v957
		goto L228
	} else {
		goto L242
	}
L242:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v988, int32(2))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v993 = int32(0)
	v1259 = v993
	v1263 = int32(6)
	v1266 = v993
	v1269 = v993
	goto L22
L244:
	;
	v999 = v957
	goto L228
L245:
	;
	if l5 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L246:
	;
	v1078 = int32(0)
	v1082 = F_Do_MultiXactIdWait(m, v167, v1000, v166, v1078, l0, v80, int32(3), v1078, v1078)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L274
	}
L247:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[76])))
	switch l4 {
	case 0:
		goto L246
	case 1:
		goto L250
	case 2:
		goto L251
	default:
		goto L245
	}
L248:
	;
	goto L249
L249:
	;
	switch l4 {
	case 0:
		goto L263
	case 1:
		goto L262
	case 2:
		goto L261
	default:
		goto L245
	}
L250:
	;
	v1031 = int32(0)
	v1035 = F_Do_MultiXactIdWait(m, v167, v1000, v166, int32(1), l0, v1031, v1031, v1031, v1031)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L258
	}
L251:
	;
	v1002 = int32(0)
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, _consts[75])))
	v1007 = F_Do_MultiXactIdWait(m, v167, v1000, v166, int32(1), l0, v1002, v1002, v1002, v1006)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	if v1007 != 0 {
		goto L245
	} else {
		goto L253
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v1016 + int32(4)
	F_errmsg(m, int32(690369), v30+int32(16))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(496630), int32(4994), int32(383134))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	if v1035 != 0 {
		goto L245
	} else {
		goto L259
	}
L259:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1037, int32(2))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v1259 = int32(0)
	v1263 = int32(6)
	v1266 = v999
	v1269 = v799
	goto L22
L261:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, _consts[75])))
	v1057 = F_ConditionalXactLockTableWait(m, v167, v1056)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L268
	}
L262:
	;
	v1047 = F_ConditionalXactLockTableWait(m, v167, int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	F_XactLockTableWait(m, v167, l0, v80, int32(3))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	goto L245
L265:
	;
	if v1047 != 0 {
		goto L245
	} else {
		goto L266
	}
L266:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1049, int32(2))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1259 = int32(0)
	v1263 = int32(6)
	v1266 = v999
	v1269 = v799
	goto L22
L268:
	;
	if v1057 != 0 {
		goto L245
	} else {
		goto L269
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1066 + int32(4)
	F_errmsg(m, int32(690369), v30)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(496630), int32(5032), int32(383134))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	goto L245
L275:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1126, int32(2))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L290
	}
L276:
	;
	if v166&int32(128) != 0 {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	if v166&int32(4176) == int32(64) {
		goto L275
	} else {
		goto L278
	}
L278:
	;
	v1094 = v30 + int32(44)
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+2)))
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v1097 = int32(16)
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1094)+2)))
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1094))))
	if v1095|v1096<<(uint(v1097)%32) == v1100|v1101<<(uint(v1097)%32) {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	if v1111 != 0 {
		goto L275
	} else {
		goto L285
	}
L280:
	;
	goto L279
L281:
	;
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1094)+4)))
	if v1107 == v1108 {
		v1111 = int32(1)
		goto L280
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1111 = int32(0)
	goto L280
L284:
	;
	goto L283
L285:
	;
	v1114 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v1116 = F_heap_lock_updated_tuple(m, l0, v166, v167, v30+int32(44), v1114, l3)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	if v1116 == int32(0) {
		goto L275
	} else {
		goto L288
	}
L288:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1120, int32(2))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1259 = int32(0)
	v1263 = v1116
	v1266 = v999
	v1269 = v799
	goto L22
L290:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1130)+20)))
	if (v1131^v166)&int32(4304) != 0 {
		v1238 = v999
		v1241 = v799
		goto L29
	} else {
		goto L291
	}
L291:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1135 != v167 {
		v1238 = v999
		v1241 = v799
		goto L29
	} else {
		goto L292
	}
L292:
	;
	if v808 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1158 = int32(0)
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1159)+20)))
	if v1160&int32(2048) != 0 {
		v1207 = v1158
		v1210 = v999
		v1213 = v799
		goto L30
	} else {
		goto L303
	}
L294:
	;
	if v1131&int32(3072) != 0 {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v1131&int32(128) != 0 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	F_HeapTupleSetHintBits(m, v1130, v1139, int32(2048), int32(0))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L302
	}
L297:
	;
	if v1131&int32(4176) == int32(64) {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1146 = F_TransactionIdDidCommit(m, v167)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	if v1146 == int32(0) {
		goto L296
	} else {
		goto L300
	}
L300:
	;
	F_HeapTupleSetHintBits(m, v1130, v1139, int32(1024), v167)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L293
L302:
	;
	goto L293
L303:
	;
	if v1160&int32(128) != 0 {
		v1207 = v1158
		v1210 = v999
		v1213 = v799
		goto L30
	} else {
		goto L304
	}
L304:
	;
	if v1160&int32(4176) == int32(64) {
		v1207 = v1158
		v1210 = v999
		v1213 = v799
		goto L30
	} else {
		goto L305
	}
L305:
	;
	v1169 = F_HeapTupleHeaderIsOnlyLocked(m, v1159)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	if v1169 != 0 {
		v1207 = v1158
		v1210 = v999
		v1213 = v799
		goto L30
	} else {
		goto L307
	}
L307:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1175 = v1173 + int32(12)
	v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+2)))
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v1178 = int32(16)
	v1181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+2)))
	v1182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175))))
	if v1176|v1177<<(uint(v1178)%32) == v1181|v1182<<(uint(v1178)%32) {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	if v1192 != 0 {
		goto L314
	} else {
		goto L315
	}
L309:
	;
	goto L308
L310:
	;
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	if v1188 == v1189 {
		v1192 = int32(1)
		goto L309
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1192 = int32(0)
	goto L309
L313:
	;
	goto L312
L314:
	;
	v1193 = int32(4)
	goto L316
L315:
	;
	v1193 = int32(3)
	goto L316
L316:
	;
	v1207 = v1193
	v1210 = v999
	v1213 = v799
	goto L30
L317:
	;
	v1291 = v1263
	v1295 = v1266
	goto L21
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+8)) = v1377
	v1397 = int32(2)
	if v1291 == v1397 {
		goto L331
	} else {
		goto L332
	}
L319:
	;
	v1377 = v1312
	goto L318
L320:
	;
	goto L321
L321:
	;
	v1318 = int32(0)
	v1322 = F_GetMultiXactIdMembers(m, v1312, v30+int32(44), v1318)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	if v1322 <= int32(0) {
		v1377 = v1318
		goto L318
	} else {
		goto L323
	}
L323:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
	v1335 = v1318
	goto L326
L324:
	;
	F_pfree(m, v1326)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L1
	} else {
		goto L330
	}
L325:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1356)))
	v1366 = v1364
	goto L324
L326:
	;
	v1356 = v1326 + v1335<<(uint(int32(3))%32)
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1357) {
		goto L325
	} else {
		goto L328
	}
L327:
	;
	v1366 = int32(0)
	goto L324
L328:
	;
	v1361 = v1335 + int32(1)
	if v1361 != v1322 {
		v1335 = v1361
		goto L326
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	v1377 = v1366
	goto L318
L331:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+8))
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400)+20)))
	if v1403&int32(32) != 0 {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+12)) = int32(-1)
	v1583 = l0
	v1593 = v30
	v1596 = v1291
	v1597 = v80
	v1599 = v1295
	v1601 = v124
	goto L17
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+12)) = v1412
	v1583 = l0
	v1593 = v30
	v1596 = v1397
	v1597 = v80
	v1599 = v1295
	v1601 = v124
	goto L17
L335:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1407+v1402<<(uint(int32(3))%32))+4))
	v1412 = v1411
	goto L337
L336:
	;
	v1412 = v1402
	goto L337
L337:
	;
	goto L334
L338:
	;
	goto L19
L339:
	;
	v1417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+10)))
	if v1417&int32(4) == int32(0) {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1422, int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	F_visibilitymap_pin(m, l0, v48, v30+int32(60))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_LockBuffer(m, v1430, int32(2))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v137 = v1259
	v144 = v1266
	v147 = v1269
	goto L18
L344:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1439)+18)))
	v1441 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	F_compute_new_xmax_infomask(m, v1436, v1435, v1440, v1441, l3, int32(0), v30+int32(56), v30+int32(54), v30+int32(52))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1452 = int32(4484100)
	v1454 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1454 + int32(1)
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1458)+20)))
	v1461 = v1459 & int32(58159)
	*(*uint16)(unsafe.Add(mBase, uint32(v1458)+20)) = uint16(v1461)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1463)+18)))
	v1466 = v1464 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(v1463)+18)) = uint16(v1466)
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1468)+20)))
	v1470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+54)))
	v1471 = v1469 | v1470
	*(*uint16)(unsafe.Add(mBase, uint32(v1468)+20)) = uint16(v1471)
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+18)))
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+52)))
	v1476 = v1474 | v1475
	*(*uint16)(unsafe.Add(mBase, uint32(v1473)+18)) = uint16(v1476)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1470&int32(128) != 0 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+10)))
	if v1502&int32(4) != 0 {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+18)))
	v1489 = v1487 & int32(49151)
	*(*uint16)(unsafe.Add(mBase, uint32(v1478)+18)) = uint16(v1489)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+4)) = v1492
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+12)) = v1495
	v1497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1494)+16)) = uint16(v1497)
	v1500 = v1492
	goto L347
L349:
	;
	if v1470&int32(4176) == int32(64) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1478)+4)) = v1485
	v1500 = v1485
	goto L347
L351:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v1507 = F_visibilitymap_clear(m, v48, v1505, int32(2))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L354
	}
L352:
	;
	v1509 = int32(0)
	goto L353
L353:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_MarkBufferDirty(m, v1511)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L355
	}
L354:
	;
	v1509 = v1507
	goto L353
L355:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1514)+118)))
	if v1515 != int32(112) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1577 = int32(4484100)
	v1579 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1579 - int32(1)
	v1583 = l0
	v1593 = v30
	v1596 = int32(0)
	v1597 = v80
	v1599 = v1266
	v1601 = v124
	goto L17
L357:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v1519 <= int32(0) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1522 != 0 {
		goto L356
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L363
	}
L361:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1523 != 0 {
		goto L356
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_XLogRegisterBuffer(m, int32(0), v1527, int32(8))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v1531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v1500
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+48)) = uint16(v1531)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1534)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+51)) = uint8(v1509)
	v1541 = int32(1)
	v1543 = int32(8)
	v1545 = int32(4)
	v1560 = int32(base.Ui32(v1535)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v1470)>>(uint(v1541)%32))&v1543 | (int32(base.Ui32(v1470)>>(uint(v1545)%32))&v1545 | (int32(base.Ui32(v1470)>>(uint(int32(12))%32))&v1541 | int32(base.Ui32(v1470)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+50)) = uint8(v1560)
	F_XLogRegisterData(m, v30+int32(44), v1543)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1569 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = base.I64_rotr(v1569, int64(32))
	goto L356
L367:
	;
	v1614 = v1583
	v1624 = v1593
	v1627 = v1596
	v1628 = v1597
	v1630 = v1599
	v1632 = v1601
	goto L16
L368:
	;
	F_ReleaseBuffer(m, v1641)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L1
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	if v1630&int32(1) != 0 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L370
L372:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1632)))
	F_UnlockTuple(m, v1614, v1628, v1646)
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L1
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	m.G0 = v1624 - int32(-64)
	return v1627
L375:
	;
	goto L374
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
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v116 int64
	_ = v116
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v328 int64
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v365 int32
	_ = v365
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v421 int32
	_ = v421
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1090 int32
	_ = v1090
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1193 int32
	_ = v1193
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1421 int32
	_ = v1421
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int64
	_ = v1440
	var v1441 int64
	_ = v1441
	var v1442 int64
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1481 int32
	_ = v1481
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1500 int64
	_ = v1500
	var v1501 int64
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
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
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1831 int32
	_ = v1831
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1981 int32
	_ = v1981
	var v1989 int32
	_ = v1989
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2006 int32
	_ = v2006
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2080 int32
	_ = v2080
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2175 int32
	_ = v2175
	var v2183 int32
	_ = v2183
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2274 int32
	_ = v2274
	var v2309 int32
	_ = v2309
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2363 int32
	_ = v2363
	var v2372 int32
	_ = v2372
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2525 int32
	_ = v2525
	var v2554 int32
	_ = v2554
	var v2561 int32
	_ = v2561
	var v2589 int32
	_ = v2589
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2607 int32
	_ = v2607
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	v11 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(8288)
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(l1^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L1
L3:
	;
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v54 = v48 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v74 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v74
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[93]))) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[94]))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95]))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = l4
	v85 = int32(1)
	v89 = int32(base.Ui32(l3)>>(uint(v85)%32)) & v85
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)) = uint8(v89)
	v92 = l3 & v85
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v74
	if v89 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v73 = v64
	goto L5
L7:
	;
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v73 = v72
	goto L5
L9:
	;
	v137 = v35 + int32(44)
	v139 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[96]))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[97]))) = uint8(v135)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[98]))) = uint8(v135)
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v145
	if base.Ui32(v144) < base.Ui32(int32(25)) {
		v1193 = v11
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[99]))) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[100]))) = v97
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v103 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[101]))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[103]))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[104]))) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[106]))) = l5 + int32(36)
	v135 = v85
	goto L9
L11:
	;
	goto L12
L12:
	;
	v116 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[99]))) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[101]))) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[100]))) = v116
	v128 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102]))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[106]))) = l5 + int32(36)
	v135 = v128
	goto L9
L13:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[94])))
	v1203 = v1201 - int32(1)
	if int32(0) <= v1203 {
		goto L134
	} else {
		goto L135
	}
L14:
	;
	v152 = int32(base.Ui32(v144+int32(262120)) >> (uint(int32(2)) % 32))
	v154 = v152 & int32(65535)
	if v154 == int32(0) {
		v1193 = v11
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v158 = int32(base.Ui32(v73) >> (uint(int32(16)) % 32))
	v166 = v35 + int32(5888)
	v173 = v152
	v174 = v154
	goto L16
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v174)
	v204 = v174 + (v35 + int32(7058))
	v205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v205)
	v207 = v174 + (v35 + int32(7350))
	v208 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v208)
	v210 = int32(1)
	v211 = v174 - v210
	v214 = v54 + int32(24) + v211<<(uint(int32(2))%32)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	switch int32(base.Ui32(v215)>>(uint(int32(15))%32))&int32(3) - v210 {
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
	v328 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v329 = base.B2i32(v139 != v328)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95])))
	v332 = v330 - int32(1)
	if v332 < int32(0) {
		v1193 = v329
		goto L13
	} else {
		goto L44
	}
L18:
	;
	v324 = v173 - int32(1)
	if v324&int32(65535) != 0 {
		v173 = v324
		v174 = v211
		goto L16
	} else {
		goto L43
	}
L19:
	;
	v257 = v54 + v215&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v174)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+14)) = uint16(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(base.Ui32(v259) >> (uint(int32(17)) % 32))
	v270 = F_HeapTupleSatisfiesVacuumHorizon(m, v35+int32(8), l1, v35+int32(7696))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95])))
	v248 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95]))) = v247 + v248
	*(*uint16)(unsafe.Add(mBase, uint32(v166+v247<<(uint(v248)%32)))) = uint16(v174)
	goto L18
L21:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v225)
	if v224 == v225 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v222)
	goto L18
L23:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v230 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(1810)+v229<<(uint(v230)%32)))) = uint16(v174)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v234 + v230
	goto L18
L24:
	;
	goto L25
L25:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102])))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102]))) = v238 + v239
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[106])))
	*(*uint16)(unsafe.Add(mBase, uint32(v242+v238<<(uint(v239)%32)))) = uint16(v174)
	goto L18
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v295)
	v299 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+18)))
	if int32(0) <= v299 {
		goto L40
	} else {
		goto L41
	}
L27:
	;
	return
L28:
	;
	if v270 != int32(2) {
		v295 = v270
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	if v274 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v292 = F_GlobalVisTestIsRemovableXid(m, v291, v287)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[107])))
	v287 = v277
	goto L30
L32:
	;
	goto L33
L33:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[107])))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	if v279 == int32(0) {
		v287 = v278
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v282 = int32(0)
	if v278-v279 < v282 {
		v295 = v282
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v287 = v278
	goto L30
L36:
	;
	if v292 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v294 = int32(0)
	goto L39
L38:
	;
	v294 = int32(2)
	goto L39
L39:
	;
	v295 = v294
	goto L26
L40:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95])))
	v303 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[95]))) = v302 + v303
	*(*uint16)(unsafe.Add(mBase, uint32(v166+v302<<(uint(v303)%32)))) = uint16(v174)
	goto L18
L41:
	;
	goto L42
L42:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[94])))
	v311 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[94]))) = v310 + v311
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(6476)+v310<<(uint(v311)%32)))) = uint16(v174)
	goto L18
L43:
	;
	goto L17
L44:
	;
	v336 = v54 + int32(24)
	v338 = v35 + int32(1228)
	v340 = v35 + int32(1810)
	v342 = v35 - int32(-64)
	v346 = v35 + int32(7058)
	v365 = v332
	goto L45
L45:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(5888)+v365<<(uint(int32(1))%32)))))
	v385 = v346 + v384
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v386 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v1193 = v329
	goto L13
L47:
	;
	if int32(0) < v365 {
		v365 = v365 - int32(1)
		goto L45
	} else {
		goto L133
	}
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v384)
	v392 = v384<<(uint(int32(2))%32) + v336 - int32(4)
	v393 = int32(0)
	v396 = int32(65535)
	v397 = v152 & v396
	if base.Ui32(v397) <= base.Ui32((v384-int32(1))&v396) {
		v533 = v393
		v535 = v393
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v604 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L50:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v563&int32(98304) != int32(65536) {
		v604 = v533
		v606 = v535
		goto L49
	} else {
		goto L82
	}
L51:
	;
	v405 = v393
	v406 = v384
	v407 = v393
	v421 = v393
	goto L52
L52:
	;
	v436 = v406 & int32(65535)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v436))))
	if v438 != 0 {
		v533 = v405
		v535 = v407
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v533 = v520
	v535 = v522
	goto L50
L54:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v436<<(uint(int32(2))%32)+v336-int32(4))))
	if v444&int32(98304) == int32(65536) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if base.Ui32((v521-int32(1))&int32(65535)) < base.Ui32(v397) {
		v405 = v520
		v406 = v521
		v407 = v522
		v421 = v525
		goto L52
	} else {
		goto L81
	}
L56:
	;
	if int32(0) < v407 {
		v533 = v405
		v535 = v407
		goto L50
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v464 = v54 + v444&int32(32767)
	if v421 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v453 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v407<<(uint(v453)%32)))) = uint16(v406)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v520 = v405
	v521 = v459 & int32(32767)
	v522 = v407 + v453
	v525 = v421
	goto L55
L60:
	;
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v464)+20)))
	v466 = int32(768)
	if v465&v466 != v466 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v476 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v407<<(uint(v476)%32)))) = uint16(v406)
	v481 = v407 + v476
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436+(v35+int32(7350))))))
	switch v483 {
	case 0:
		goto L68
	case 1, 3, 4:
		v604 = v405
		v606 = v481
		goto L49
	case 2:
		v499 = v405
		goto L67
	default:
		goto L69
	}
L63:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	v472 = v470
	goto L65
L64:
	;
	v472 = int32(2)
	goto L65
L65:
	;
	if v472 != v421 {
		v533 = v405
		v535 = v407
		goto L50
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+19)))
	if v500&int32(64) == int32(0) {
		v604 = v499
		v606 = v481
		goto L49
	} else {
		goto L74
	}
L68:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v464, v137)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L27
	} else {
		goto L73
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	F_errmsg_internal(m, int32(97822), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(495438), int32(1117), int32(278001))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
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
	v499 = v481
	goto L67
L74:
	;
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v464)+20)))
	if v505&int32(2048) != 0 {
		v604 = v499
		v606 = v481
		goto L49
	} else {
		goto L75
	}
L75:
	;
	if v505&int32(768) == int32(512) {
		v604 = v499
		v606 = v481
		goto L49
	} else {
		goto L76
	}
L76:
	;
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v464)+16)))
	if v505&int32(4224) == int32(4096) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v517 = F_HeapTupleGetUpdateXid(m, v464)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L27
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v520 = v499
	v521 = v512
	v522 = v481
	v525 = v519
	goto L55
L80:
	;
	v520 = v499
	v521 = v512
	v522 = v481
	v525 = v517
	goto L55
L81:
	;
	goto L53
L82:
	;
	if int32(1) < v535 {
		v604 = v533
		v606 = v535
		goto L49
	} else {
		goto L83
	}
L83:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v571)
	if v570 == v571 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v576 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v575<<(uint(v576)%32)))) = uint16(v384)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v580 + v576
	goto L47
L85:
	;
	goto L86
L86:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v585 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v338+v584<<(uint(v585)%32)))) = uint16(v384)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v589 + v585
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102]))) = v593 + v585
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[106])))
	*(*uint16)(unsafe.Add(mBase, uint32(v597+v593<<(uint(v585)%32)))) = uint16(v384)
	goto L47
L87:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v637&int32(98304) == int32(65536) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if v604 == v606 {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	v642 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v642)
	v645 = v642
	goto L92
L91:
	;
	v645 = int32(0)
	goto L92
L92:
	;
	if v606 <= v645 {
		goto L47
	} else {
		goto L93
	}
L93:
	;
	v650 = v645
	goto L94
L94:
	;
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v650<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L27
	} else {
		goto L96
	}
L95:
	;
	goto L47
L96:
	;
	v690 = v650 + int32(1)
	if v690 != v606 {
		v650 = v690
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v695 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v695)
	v698 = v693 & int32(98304)
	if v694 == v695 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	goto L100
L100:
	;
	v882 = int32(1)
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v604<<(uint(v882)%32)))))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v882)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v890 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v342+v889<<(uint(v890)%32)))) = uint16(v384)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v342+v894<<(uint(v890)%32))+2)) = uint16(v885)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v899 + v882
	if v886&int32(98304) == int32(32768) {
		goto L116
	} else {
		goto L117
	}
L101:
	;
	if v606 < int32(2) {
		goto L47
	} else {
		goto L108
	}
L102:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v733 + int32(1)
	goto L101
L103:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v702 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v701<<(uint(v702)%32)))) = uint16(v384)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v706 + v702
	if v698 == int32(32768) {
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v713 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v338+v712<<(uint(v713)%32)))) = uint16(v384)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v717 + v713
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102]))) = v721 + v713
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[106])))
	*(*uint16)(unsafe.Add(mBase, uint32(v725+v721<<(uint(v713)%32)))) = uint16(v384)
	if v698 != int32(32768) {
		goto L101
	} else {
		goto L107
	}
L106:
	;
	goto L101
L107:
	;
	goto L102
L108:
	;
	v740 = int32(1)
	v742 = v606 - v740
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v606 != int32(2) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v753 = v745
	v754 = v740
	v755 = int32(0)
	goto L112
L110:
	;
	v827 = v745
	v828 = v740
	goto L111
L111:
	;
	if v742&v740 == int32(0) {
		goto L47
	} else {
		goto L115
	}
L112:
	;
	v785 = int32(1)
	v787 = v35 + int32(7696) + v754<<(uint(v785)%32)
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v787))))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v788))) = uint8(v785)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v753<<(uint(v785)%32)))) = uint16(v788)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v796 + v785
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v802 = v800 + v785
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v802
	v804 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v787)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v804))) = uint8(v785)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v802<<(uint(v785)%32)))) = uint16(v804)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v814 = v812 + v785
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v816 + v785
	v820 = int32(2)
	v821 = v754 + v820
	v823 = v755 + v820
	if v823 != v742&int32(-2) {
		v753 = v814
		v754 = v821
		v755 = v823
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v827 = v814
	v828 = v821
	goto L111
L114:
	;
	goto L113
L115:
	;
	v861 = int32(1)
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v828<<(uint(v861)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v864))) = uint8(v861)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v827<<(uint(v861)%32)))) = uint16(v864)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v872 + v861
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v876 + v861
	goto L47
L116:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v907 + int32(1)
	goto L118
L117:
	;
	goto L118
L118:
	;
	v911 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[108]))) = uint8(v911)
	if v604 < int32(2) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if v606 <= v604 {
		goto L47
	} else {
		goto L128
	}
L120:
	;
	v915 = int32(1)
	v917 = v604 - v915
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v604 != int32(2) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v929 = v915
	v937 = v920
	v940 = int32(0)
	goto L124
L122:
	;
	v1003 = v915
	v1011 = v920
	goto L123
L123:
	;
	if v917&v915 == int32(0) {
		goto L119
	} else {
		goto L127
	}
L124:
	;
	v960 = int32(1)
	v962 = v35 + int32(7696) + v929<<(uint(v960)%32)
	v963 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962))))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v963))) = uint8(v960)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v937<<(uint(v960)%32)))) = uint16(v963)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v971 + v960
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v977 = v975 + v960
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v977
	v979 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v979))) = uint8(v960)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v977<<(uint(v960)%32)))) = uint16(v979)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v989 = v987 + v960
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v989
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v991 + v960
	v995 = int32(2)
	v996 = v929 + v995
	v998 = v940 + v995
	if v998 != v917&int32(-2) {
		v929 = v996
		v937 = v989
		v940 = v998
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v1003 = v996
	v1011 = v989
	goto L123
L126:
	;
	goto L125
L127:
	;
	v1036 = int32(1)
	v1039 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v1003<<(uint(v1036)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v1039))) = uint8(v1036)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v1011<<(uint(v1036)%32)))) = uint16(v1039)
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v1047 + v1036
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v1051 + v1036
	goto L119
L128:
	;
	v1090 = v604
	goto L129
L129:
	;
	v1127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(7696)+v1090<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v1127)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L27
	} else {
		goto L131
	}
L130:
	;
	goto L47
L131:
	;
	v1131 = v1090 + int32(1)
	if v1131 != v606 {
		v1090 = v1131
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	goto L46
L134:
	;
	v1219 = v1203
	goto L137
L135:
	;
	goto L136
L136:
	;
	v1354 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1354)
	v1356 = int32(1)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	if v1354 < v1358 {
		v1368 = v1356
		goto L155
	} else {
		goto L156
	}
L137:
	;
	v1251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(6476)+v1219<<(uint(int32(1))%32)))))
	v1252 = v35 + int32(7058) + v1251
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252))))
	if v1253 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L136
L139:
	;
	if int32(0) < v1219 {
		v1219 = v1219 - int32(1)
		goto L137
	} else {
		goto L154
	}
L140:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1251)
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251+(v35+int32(7350))))))
	if v1256 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1251<<(uint(int32(2))%32)+(v54+int32(24))-int32(4))))
	v1267 = v54 + v1264&int32(32767)
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267)+19)))
	if v1268&int32(64) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	F_heap_prune_record_unchanged_lp_normal(m, v54, v35+int32(28), v1251)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L27
	} else {
		goto L153
	}
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L27
	} else {
		goto L150
	}
L145:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1267, v137)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L27
	} else {
		goto L149
	}
L146:
	;
	v1273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1267)+20)))
	if v1273&int32(2048) != 0 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	if v1273&int32(768) != int32(512) {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	v1283 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1252))) = uint8(v1283)
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(1810)+v1285<<(uint(v1283)%32)))) = uint16(v1251)
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v1290 + v1283
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105]))) = v1294 + v1283
	goto L139
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v1251
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v73
	F_errmsg_internal(m, int32(278018), v35)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L27
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(495438), int32(635), int32(341922))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L27
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	goto L139
L154:
	;
	goto L138
L155:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if v1369 == v1370 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	if int32(0) < v1362 {
		v1368 = int32(1)
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v1368 = base.B2i32(int32(0) < v1365)
	goto L155
L158:
	;
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)))
	v1377 = int32(base.Ui32(v1372&int32(2)) >> (uint(int32(1)) % 32))
	goto L160
L159:
	;
	v1377 = v1356
	goto L160
L160:
	;
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)))
	if v1378 != int32(1) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v1724 = int32(4484100)
	v1726 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1726 + int32(1)
	if v1377 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L162:
	;
	v1683 = int32(0)
	if v1682 <= v1683 {
		v1723 = v1683
		goto L161
	} else {
		goto L241
	}
L163:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1682 = v1681
	goto L162
L164:
	;
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[93]))))
	if v1381 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1518 = m.G0
	v1520 = v1518 - int32(32)
	m.G0 = v1520
	if l1 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L166:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[98]))))
	if v1382 != int32(1) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[97]))))
	if v1386 != int32(1) {
		v1682 = v1385
		goto L162
	} else {
		goto L168
	}
L168:
	;
	if v1385 <= int32(0) {
		v1682 = v1385
		goto L162
	} else {
		goto L169
	}
L169:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391)+118)))
	if v1392 != int32(112) {
		goto L163
	} else {
		goto L170
	}
L170:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v1396 <= int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v1368 != 0 {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1399 != 0 {
		goto L163
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	if v1193 != 0 {
		goto L165
	} else {
		goto L178
	}
L175:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1401 = int32(0)
	if v1193|base.B2i32(v1400 != v1401) == v1401 {
		goto L171
	} else {
		goto L176
	}
L176:
	;
	if v1400 != 0 {
		goto L163
	} else {
		goto L177
	}
L177:
	;
	goto L165
L178:
	;
	goto L171
L179:
	;
	v1409 = m.G0
	v1411 = v1409 - int32(16)
	m.G0 = v1411
	F_GetFullPageWriteInfo(m, v1411+int32(8), v1411+int32(7))
	mBase = m.M
	if l1 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	goto L181
L181:
	;
	if v1377 == int32(0) {
		goto L163
	} else {
		goto L193
	}
L182:
	;
	if v1448 == int32(0) {
		goto L163
	} else {
		goto L192
	}
L183:
	;
	v1436 = int32(1)
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1411)+7)))
	if v1437 == v1436 {
		goto L188
	} else {
		goto L189
	}
L184:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1421+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1435 = v1427
	goto L183
L185:
	;
	goto L186
L186:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1435 = v1429 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L183
L187:
	;
	m.G0 = v1411 + int32(16)
	goto L182
L188:
	;
	v1440 = *(*int64)(unsafe.Add(mBase, uint32(v1411)+8))
	v1441 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1435)+4)))
	v1442 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1435))))
	if base.Ui64(v1441|v1442<<(uint(int64(32))%64)) <= base.Ui64(v1440) {
		v1448 = v1436
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1448 = int32(0)
	goto L187
L191:
	;
	goto L190
L192:
	;
	goto L165
L193:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+252))
	goto L194
L194:
	;
	if base.B2i32(v1458 != int32(0)) == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1464 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
	if v1464 != int32(1) {
		goto L163
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1469 = m.G0
	v1471 = v1469 - int32(16)
	m.G0 = v1471
	F_GetFullPageWriteInfo(m, v1471+int32(8), v1471+int32(7))
	mBase = m.M
	if l1 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	goto L197
L199:
	;
	if v1508 == int32(0) {
		goto L163
	} else {
		goto L209
	}
L200:
	;
	v1496 = int32(1)
	v1497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471)+7)))
	if v1497 == v1496 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1481+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1495 = v1487
	goto L200
L202:
	;
	goto L203
L203:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1495 = v1489 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L200
L204:
	;
	m.G0 = v1471 + int32(16)
	goto L199
L205:
	;
	v1500 = *(*int64)(unsafe.Add(mBase, uint32(v1471)+8))
	v1501 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1495)+4)))
	v1502 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1495))))
	if base.Ui64(v1501|v1502<<(uint(int64(32))%64)) <= base.Ui64(v1500) {
		v1508 = v1496
		goto L204
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v1508 = int32(0)
	goto L204
L208:
	;
	goto L207
L209:
	;
	goto L165
L210:
	;
	if int32(0) < v1517 {
		goto L217
	} else {
		goto L218
	}
L211:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1525+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1539 = v1531
	goto L210
L212:
	;
	goto L213
L213:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1539 = v1533 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L210
L214:
	;
	v1723 = int32(1)
	goto L161
L215:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L27
	} else {
		goto L237
	}
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L27
	} else {
		goto L233
	}
L217:
	;
	v1547 = int32(0)
	goto L220
L218:
	;
	goto L219
L219:
	;
	m.G0 = v1520 + int32(32)
	goto L214
L220:
	;
	v1579 = v35 + int32(2392) + v1547*int32(12)
	v1580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1579)+10)))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1580<<(uint(int32(2))%32)+(v1539+int32(24))-int32(4))))
	v1589 = v1539 + v1586&int32(32767)
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579)+9)))
	if v1590&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L219
L222:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1589)))
	v1594 = F_TransactionIdDidCommit(m, v1593)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L27
	} else {
		goto L225
	}
L223:
	;
	v1600 = v1590
	goto L224
L224:
	;
	if v1600&int32(2) != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	if v1594 == int32(0) {
		goto L216
	} else {
		goto L226
	}
L226:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579)+9)))
	v1600 = v1598
	goto L224
L227:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1589)+4))
	v1604 = F_TransactionIdDidCommit(m, v1603)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L27
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1608 = v1547 + int32(1)
	if v1608 != v1517 {
		v1547 = v1608
		goto L220
	} else {
		goto L232
	}
L230:
	;
	if v1604 != 0 {
		goto L215
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	goto L221
L233:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L27
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1520)+16)) = v1593
	F_errmsg_internal(m, int32(280693), v1520+int32(16))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L27
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(496630), int32(7359), int32(154341))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L27
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L27
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1520))) = v1603
	F_errmsg_internal(m, int32(40293), v1520)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L27
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(496630), int32(7376), int32(154341))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L27
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	v1686 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+60)) = v1686
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[97]))) = uint8(v1686)
	v1723 = v1686
	goto L161
L242:
	;
	if v1368|v1723 != int32(1) {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v1732
	v1734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)))
	v1736 = v1734 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)) = uint16(v1736)
	if (v1368|v1723)&int32(1) != 0 {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	F_MarkBufferDirtyHint(m, l1, int32(1))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L27
	} else {
		goto L245
	}
L245:
	;
	goto L242
L246:
	;
	v2651 = int32(4484100)
	v2653 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2654 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2653 - v2654
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[105])))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v2657
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v2659
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v2661
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[111])))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v2663
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[112])))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v2665
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[102])))
	v2668 = int32(0)
	v2670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[98]))))
	if base.B2i32(v2667 == v2668)&base.B2i32(v2670 == v2654) == v2668 {
		goto L333
	} else {
		goto L334
	}
L247:
	;
	if v1368 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	if l1 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	goto L250
L250:
	;
	if v1723 != 0 {
		goto L289
	} else {
		goto L290
	}
L251:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	if v1747 <= int32(0) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1751+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1765 = v1757
	goto L251
L253:
	;
	goto L254
L254:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1765 = v1759 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L251
L255:
	;
	if v1767 <= int32(0) {
		goto L264
	} else {
		goto L265
	}
L256:
	;
	v1771 = v35 - int32(-64)
	v1772 = int32(1)
	v1775 = v1765 + int32(24)
	if v1747 != v1772 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1783 = int32(0)
	v1785 = v1771
	goto L260
L258:
	;
	v1846 = v1771
	goto L259
L259:
	;
	if v1747&v1772 == int32(0) {
		goto L255
	} else {
		goto L263
	}
L260:
	;
	v1813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1785))))
	v1814 = int32(2)
	v1817 = int32(4)
	v1819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1785)+2)))
	v1820 = int32(32767)
	v1822 = int32(65536)
	*(*int32)(unsafe.Add(mBase, uint32(v1813<<(uint(v1814)%32)+v1775-v1817))) = v1819&v1820 | v1822
	v1825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1785)+4)))
	v1831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1785)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1825<<(uint(v1814)%32)+v1775-v1817))) = v1831&v1820 | v1822
	v1838 = v1785 + int32(8)
	v1840 = v1783 + v1814
	if v1840 != v1747&int32(2147483646) {
		v1783 = v1840
		v1785 = v1838
		goto L260
	} else {
		goto L262
	}
L261:
	;
	v1846 = v1838
	goto L259
L262:
	;
	goto L261
L263:
	;
	v1876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1846))))
	v1882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1846)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1876<<(uint(int32(2))%32)+v1775-int32(4)))) = v1882&int32(32767) | int32(65536)
	goto L255
L264:
	;
	if v1766 <= int32(0) {
		goto L276
	} else {
		goto L277
	}
L265:
	;
	v1923 = v35 + int32(1228)
	v1925 = v1767 & int32(3)
	v1927 = v1765 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v1767) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1935 = int32(0)
	v1937 = v1923
	goto L269
L267:
	;
	v2006 = v1923
	goto L268
L268:
	;
	if v1925 == int32(0) {
		goto L264
	} else {
		goto L272
	}
L269:
	;
	v1965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1937))))
	v1966 = int32(2)
	v1969 = int32(4)
	v1971 = int32(98304)
	*(*int32)(unsafe.Add(mBase, uint32(v1965<<(uint(v1966)%32)+v1927-v1969))) = v1971
	v1973 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1937)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1973<<(uint(v1966)%32)+v1927-v1969))) = v1971
	v1981 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1937)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1981<<(uint(v1966)%32)+v1927-v1969))) = v1971
	v1989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1937)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1989<<(uint(v1966)%32)+v1927-v1969))) = v1971
	v1998 = v1937 + int32(8)
	v2000 = v1935 + v1969
	if v2000 != v1767&int32(2147483644) {
		v1935 = v2000
		v1937 = v1998
		goto L269
	} else {
		goto L271
	}
L270:
	;
	v2006 = v1998
	goto L268
L271:
	;
	goto L270
L272:
	;
	v2039 = int32(0)
	v2041 = v2006
	goto L273
L273:
	;
	v2069 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2041))))
	v2070 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2069<<(uint(v2070)%32)+v1927-int32(4)))) = int32(98304)
	v2080 = v2039 + int32(1)
	if v2080 != v1925 {
		v2039 = v2080
		v2041 = v2041 + v2070
		goto L273
	} else {
		goto L275
	}
L274:
	;
	goto L264
L275:
	;
	goto L274
L276:
	;
	F_PageRepairFragmentation(m, v1765)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L27
	} else {
		goto L288
	}
L277:
	;
	v2117 = v35 + int32(1810)
	v2119 = v1766 & int32(3)
	v2121 = v1765 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v1766) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v2129 = int32(0)
	v2131 = v2117
	goto L281
L279:
	;
	v2200 = v2117
	goto L280
L280:
	;
	if v2119 == int32(0) {
		goto L276
	} else {
		goto L284
	}
L281:
	;
	v2159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131))))
	v2160 = int32(2)
	v2163 = int32(4)
	v2165 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2159<<(uint(v2160)%32)+v2121-v2163))) = v2165
	v2167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v2167<<(uint(v2160)%32)+v2121-v2163))) = v2165
	v2175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2175<<(uint(v2160)%32)+v2121-v2163))) = v2165
	v2183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v2183<<(uint(v2160)%32)+v2121-v2163))) = v2165
	v2192 = v2131 + int32(8)
	v2194 = v2129 + v2163
	if v2194 != v1766&int32(2147483644) {
		v2129 = v2194
		v2131 = v2192
		goto L281
	} else {
		goto L283
	}
L282:
	;
	v2200 = v2192
	goto L280
L283:
	;
	goto L282
L284:
	;
	v2233 = int32(0)
	v2235 = v2200
	goto L285
L285:
	;
	v2263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2235))))
	v2264 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2263<<(uint(v2264)%32)+v2121-int32(4)))) = int32(0)
	v2274 = v2233 + int32(1)
	if v2274 != v2119 {
		v2233 = v2274
		v2235 = v2235 + v2264
		goto L285
	} else {
		goto L287
	}
L286:
	;
	goto L276
L287:
	;
	goto L286
L288:
	;
	goto L250
L289:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v2345 = int32(0)
	if l1 < v2345 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	goto L291
L291:
	;
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L27
	} else {
		goto L308
	}
L292:
	;
	if int32(0) < v2344 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2349+(l1^int32(-1))<<(uint(int32(2))%32))))
	v2363 = v2355
	goto L292
L294:
	;
	goto L295
L295:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v2363 = v2357 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L292
L296:
	;
	v2372 = v2345
	goto L299
L297:
	;
	goto L298
L298:
	;
	goto L291
L299:
	;
	v2402 = v35 + int32(2392) + v2372*int32(12)
	v2403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2402)+10)))
	v2404 = int32(2)
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v2403<<(uint(v2404)%32)+(v2363+int32(24))-int32(4))))
	v2412 = v2363 + v2409&int32(32767)
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2402)))
	*(*int32)(unsafe.Add(mBase, uint32(v2412)+4)) = v2413
	v2415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2402)+8)))
	if v2415&v2404 != 0 {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	goto L298
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2412)+8)) = int32(2)
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2402)+8)))
	v2421 = v2420
	goto L303
L302:
	;
	v2421 = v2415
	goto L303
L303:
	;
	if v2421&int32(4) != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2412)+8)) = int32(0)
	goto L306
L305:
	;
	goto L306
L306:
	;
	v2426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2402)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2412)+20)) = uint16(v2426)
	v2428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2402)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2412)+18)) = uint16(v2428)
	v2431 = v2372 + int32(1)
	if v2431 != v2344 {
		v2372 = v2431
		goto L299
	} else {
		goto L307
	}
L307:
	;
	goto L300
L308:
	;
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2499)+118)))
	if v2500 != int32(112) {
		goto L246
	} else {
		goto L309
	}
L309:
	;
	v2503 = int32(0)
	v2505 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v2505 <= v2503 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2508 != 0 {
		goto L246
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	if v1723 == int32(0) {
		v2561 = v2503
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2509 != 0 {
		goto L246
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2589))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2561)) == int32(0) {
		goto L324
	} else {
		goto L325
	}
L316:
	;
	v2512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[98]))))
	if v2512 != int32(1) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v2519)+8))
	v2525 = v2520
	goto L320
L318:
	;
	v2515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[97]))))
	if v2515 != int32(1) {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[96])))
	v2561 = v2518
	goto L315
L320:
	;
	v2554 = v2525 - int32(1)
	if base.Ui32(v2554) < base.Ui32(int32(3)) {
		v2525 = v2554
		goto L320
	} else {
		goto L322
	}
L321:
	;
	v2561 = v2554
	goto L315
L322:
	;
	goto L321
L323:
	;
	v2602 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	if v2601 != 0 {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	v2601 = base.B2i32(base.Ui32(v2589) < base.Ui32(v2561))
	goto L323
L325:
	;
	goto L326
L326:
	;
	v2601 = base.B2i32(int32(0) < v2561-v2589)
	goto L323
L327:
	;
	v2603 = v2561
	goto L329
L328:
	;
	v2603 = v2602
	goto L329
L329:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	F_log_heap_prune_and_freeze(m, l0, l1, v2603, int32(1), l6, v35+int32(2392), v2607, v35-int32(-64), v2610, v35+int32(1228), v2613, v35+int32(1810), v2616)
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L27
	} else {
		goto L330
	}
L330:
	;
	goto L246
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v2667
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v2690
	v2693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+33)))
	if v2693 == int32(1) {
		goto L337
	} else {
		goto L338
	}
L332:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[96])))
	v2690 = v2688
	goto L331
L333:
	;
	v2676 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+20)) = uint16(v2676)
	v2678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[108]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v2678)
	goto L332
L334:
	;
	goto L335
L335:
	;
	v2680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[97]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)) = uint8(v2680)
	v2682 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+20)) = uint8(v2682)
	v2684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[108]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v2684)
	if v2680 != 0 {
		v2690 = int32(0)
		goto L331
	} else {
		goto L336
	}
L336:
	;
	goto L332
L337:
	;
	v2701 = base.B2i32(int32(0) < v2661)
	if int32(0) < v2661 {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	goto L339
L339:
	;
	m.G0 = v35 + int32(8288)
	return
L340:
	;
	v2702 = int32(7620)
	goto L342
L341:
	;
	v2702 = int32(7628)
	goto L342
L342:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(28)+v2702)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v2704
	if int32(0) < v2661 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2710 = int32(7624)
	goto L345
L344:
	;
	v2710 = int32(7632)
	goto L345
L345:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(28)+v2710)))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v2712
	goto L339
}
func F_heap_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
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
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
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
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int64
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int64
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int64
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
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
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v833 int64
	_ = v833
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int64
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(8256)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+48)))
	switch int32(base.Ui32(v17)>>(uint(int32(4))%32))&int32(7) - int32(1) {
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
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L16
	} else {
		goto L266
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L16
	} else {
		goto L263
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L16
	} else {
		goto L260
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L16
	} else {
		goto L257
	}
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L16
	} else {
		goto L254
	}
L6:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L16
	} else {
		goto L251
	}
L7:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L16
	} else {
		goto L248
	}
L8:
	;
	m.G0 = v14 + int32(8256)
	return
L9:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v851 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v855 = F_XLogReadBufferForRedo(m, l0, int32(0), v14+int32(92))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L16
	} else {
		goto L213
	}
L10:
	;
	v651 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+7)))
	if v653&int32(1) != 0 {
		goto L175
	} else {
		goto L176
	}
L11:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v560 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v564 = F_XLogReadBufferForRedo(m, l0, int32(0), v14+int32(92))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L16
	} else {
		goto L155
	}
L12:
	;
	F_heap_xlog_update(m, l0, int32(1))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L16
	} else {
		goto L154
	}
L13:
	;
	F_heap_xlog_update(m, l0, int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L16
	} else {
		goto L153
	}
L14:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v350 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v350, v14+int32(92), v350, v14+int32(8252))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L16
	} else {
		goto L102
	}
L15:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v26 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v26, v14+int32(80), v26, v14+int32(76))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	if v36&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v41
	v45 = F_CreateFakeRelcacheEntry(m, v14+int32(24))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+48)))
	if v66 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	F_visibilitymap_pin(m, v45, v49, v14+int32(92))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v57 = F_visibilitymap_clear(m, v54, v55, int32(3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_ReleaseBuffer(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v45)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	if v331 != 0 {
		goto L96
	} else {
		goto L97
	}
L27:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v138 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L28:
	;
	v70 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v135 = F_XLogReadBufferForRedo(m, l0, int32(0), v14+int32(8252))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L16
	} else {
		goto L46
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81]))) = v70
	if v70 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v90&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v70^int32(-1))<<(uint(int32(2))%32))))
	v90 = v82
	goto L32
L34:
	;
	goto L35
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v90 = v84 + v70<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v138 = v70
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+10)) = int32(1572864)
	v123 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+18)) = uint16(v123)
	v129 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)) = uint16(v129)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+14)) = uint16(v129)
	goto L36
L38:
	;
	v117 = F___memset(m, v90, int32(0), int32(8192))
	mBase = m.M
	goto L37
L39:
	;
	goto L38
L46:
	;
	if v135 != 0 {
		v327 = v2
		v328 = v2
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	v138 = v137
	goto L27
L48:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+12)))
	if base.Ui32(v159) < base.Ui32(int32(25)) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144+(v138^int32(-1))<<(uint(int32(2))%32))))
	v158 = v150
	goto L48
L50:
	;
	goto L51
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v158 = v152 + v138<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	v168 = int32(1)
	goto L54
L53:
	;
	v168 = int32(base.Ui32(v159+int32(262120))>>(uint(int32(2))%32)) + int32(1)
	goto L54
L54:
	;
	if base.Ui32(v168&int32(65535)) < base.Ui32(v139) {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v172 = int32(0)
	v174 = v14 + int32(72)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+72))
	if v177 < v172 {
		v199 = v172
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+2)))
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202))))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = int32(0)
	v210 = int32(5)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v214 = v212 - v210
	if v214 != 0 {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	v202 = v199
	goto L56
L58:
	;
	v183 = v176 + int32(76)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v184 != int32(1) {
		v199 = v172
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+43)))
	if v187 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v174 == int32(0) {
		v199 = v172
		goto L57
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v174 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v192
	v202 = v192
	goto L56
L64:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v195
	goto L66
L65:
	;
	goto L66
L66:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v199 = v197
	goto L57
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+114)) = uint8(v205)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+110)) = uint16(v204)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+36))
	v222 = v203 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+112)) = uint16(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = v220
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+108)) = uint16(v35)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+106)) = uint16(v34)
	v230 = int32(base.Ui32(v34) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+104)) = uint16(v230)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v238 = F_PageAddItemExtended(m, v158, v14+int32(92), v212+int32(18), v236, int32(3))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L16
	} else {
		goto L71
	}
L68:
	;
	v215 = F__emscripten_memcpy_bulkmem(m, v14+int32(115), v202+v210, v214)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	if v238 == int32(0) {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v245 = int32(4)
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+14)))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+12)))
	v248 = v246 - v247
	if v248 <= v245 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = base.I64_rotr(v24, int64(32))
	v314 = int32(1)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	if v315&v314 != 0 {
		goto L92
	} else {
		goto L93
	}
L74:
	;
	v251 = v245
	goto L76
L75:
	;
	v251 = v248
	goto L76
L76:
	;
	v253 = v251 - int32(4)
	if v253 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v310 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v247) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v310 = v253
	goto L73
L81:
	;
	v264 = int32(base.Ui32(v247+int32(262120)) >> (uint(int32(2)) % 32))
	goto L83
L82:
	;
	v264 = int32(0)
	goto L83
L83:
	;
	if base.Ui32(v264&int32(65535)) < base.Ui32(int32(291)) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+10)))
	if v269&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v310 = int32(0)
	goto L73
L86:
	;
	goto L87
L87:
	;
	v278 = int32(1)
	goto L88
L88:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278&int32(65535)<<(uint(int32(2))%32)+(v158+int32(24))-int32(3)))))
	if v289&int32(384) == int32(0) {
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v310 = int32(0)
	goto L73
L90:
	;
	v295 = v278 + int32(1)
	v296 = int32(65535)
	if base.Ui32(v295&v296) <= base.Ui32(v264&v296) {
		v278 = v295
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+10)))
	v320 = v318 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+10)) = uint16(v320)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	F_MarkBufferDirty(m, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	v327 = v310
	v328 = v314
	goto L26
L96:
	;
	F_UnlockReleaseBuffer(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v328&base.B2i32(base.Ui32(v327) < base.Ui32(int32(1638))) == int32(0) {
		goto L8
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v339
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v341
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	F_XLogRecordPageWithFreeSpace(m, v14+int32(8), v345, v327)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	goto L8
L102:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+4)))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+7)))
	if v360&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v14)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v365
	v369 = F_CreateFakeRelcacheEntry(m, v14+int32(40))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L16
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v392 = F_XLogReadBufferForRedo(m, l0, int32(0), v14+int32(80))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L16
	} else {
		goto L111
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	F_visibilitymap_pin(m, v369, v373, v14+int32(80))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v381 = F_visibilitymap_clear(m, v378, v379, int32(3))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	F_ReleaseBuffer(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	F_pfree(m, v369)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L16
	} else {
		goto L110
	}
L110:
	;
	goto L105
L111:
	;
	if v392 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	if v396 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L114
L114:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	if v548 == int32(0) {
		goto L8
	} else {
		goto L151
	}
L115:
	;
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v415) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+(v396^int32(-1))<<(uint(int32(2))%32))))
	v414 = v406
	goto L115
L117:
	;
	goto L118
L118:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v414 = v408 + v396<<(uint(int32(13))%32) + int32(-8192)
	goto L115
L119:
	;
	v423 = int32(base.Ui32(v415+int32(262120)) >> (uint(int32(2)) % 32))
	goto L121
L120:
	;
	v423 = int32(0)
	goto L121
L121:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+4)))
	if base.Ui32(v423&int32(65535)) < base.Ui32(v426) {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426<<(uint(int32(2))%32)+v414)+20))
	if v431&int32(98304) != int32(32768) {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v438 = v414 + v431&int32(32767)
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v438)+20)))
	v441 = v439 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+20)) = uint16(v441)
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v438)+18)))
	v445 = v443 & int32(-24577)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+18)) = uint16(v445)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+18)) = uint16(v445)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+20)) = uint16(v441)
	v450 = int32(1)
	v469 = v447<<(uint(v450)%32)&int32(16) | (v447<<(uint(int32(4))%32)&int32(64) | (v447<<(uint(int32(6))%32)&int32(128) | v447&v450<<(uint(int32(12))%32))) | v441
	if v447&int32(15) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+20)) = uint16(v469)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v447&int32(16) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v476 = v445 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+18)) = uint16(v476)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+7)))
	if v478&int32(8) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+8)) = int32(0)
	v490 = v469 & int32(13279)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+20)) = uint16(v490)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v414)+20))
	if v492 != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v483
	goto L130
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = int32(0)
	goto L130
L134:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+7)))
	if v512&int32(1) != 0 {
		goto L144
	} else {
		goto L145
	}
L135:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v492))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v494)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v414)+20)) = v510
	goto L134
L138:
	;
	if v506 == int32(0) {
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v506 = base.B2i32(base.Ui32(v494) < base.Ui32(v492))
	goto L138
L140:
	;
	goto L141
L141:
	;
	v506 = int32(base.Ui32(v494-v492) >> (uint(int32(31)) % 32))
	goto L138
L142:
	;
	goto L137
L143:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+14)) = uint16(v532)
	*(*uint32)(unsafe.Add(mBase, uint32(v414)+4)) = uint32(v348)
	v536 = int64(base.Ui64(v348) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v414))) = uint32(v536)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	F_MarkBufferDirty(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L16
	} else {
		goto L150
	}
L144:
	;
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+10)))
	v517 = v515 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v414)+10)) = uint16(v517)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+7)))
	v520 = v519
	goto L146
L145:
	;
	v520 = v512
	goto L146
L146:
	;
	if v520&int32(16) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v523 = int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+16)) = uint16(v523)
	v525 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+12)) = uint16(v525)
	v532 = v525
	goto L143
L148:
	;
	goto L149
L149:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+16)) = uint16(v359)
	v530 = int32(base.Ui32(v358) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+12)) = uint16(v530)
	v532 = v358
	goto L143
L150:
	;
	goto L114
L151:
	;
	F_UnlockReleaseBuffer(m, v548)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
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
	if v564 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v559))))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v569 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v646 == int32(0) {
		goto L8
	} else {
		goto L173
	}
L159:
	;
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v588) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573+(v569^int32(-1))<<(uint(int32(2))%32))))
	v587 = v579
	goto L159
L161:
	;
	goto L162
L162:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v587 = v581 + v569<<(uint(int32(13))%32) + int32(-8192)
	goto L159
L163:
	;
	v596 = int32(base.Ui32(v588+int32(262120)) >> (uint(int32(2)) % 32))
	goto L165
L164:
	;
	v596 = int32(0)
	goto L165
L165:
	;
	if base.Ui32(v596&int32(65535)) < base.Ui32(v568) {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v568<<(uint(int32(2))%32)+v587)+20))
	if v603&int32(98304) != int32(32768) {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	if v569 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v629 = v587 + v603&int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v629)+16)) = uint16(v568)
	*(*uint16)(unsafe.Add(mBase, uint32(v629)+14)) = uint16(v626)
	v633 = int32(base.Ui32(v626) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v629)+12)) = uint16(v633)
	*(*uint32)(unsafe.Add(mBase, uint32(v587)+4)) = uint32(v560)
	v637 = int64(base.Ui64(v560) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v587))) = uint32(v637)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_MarkBufferDirty(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L16
	} else {
		goto L172
	}
L169:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v611+(v569^int32(-1))<<(uint(int32(6))%32))+16))
	v626 = v617
	goto L168
L170:
	;
	goto L171
L171:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619+v569<<(uint(int32(6))%32)+int32(-64))+16))
	v626 = v625
	goto L168
L172:
	;
	goto L158
L173:
	;
	F_UnlockReleaseBuffer(m, v646)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	goto L8
L175:
	;
	v656 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v656
	F_XLogRecGetBlockTag(m, l0, v656, v14+int32(92), v656, v14+int32(8252))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L16
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v695 = F_XLogReadBufferForRedo(m, l0, int32(0), v14+int32(92))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L16
	} else {
		goto L184
	}
L178:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v668
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v14)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v670
	v674 = F_CreateFakeRelcacheEntry(m, v14+int32(56))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L16
	} else {
		goto L179
	}
L179:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	F_visibilitymap_pin(m, v674, v676, v14+int32(80))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L16
	} else {
		goto L180
	}
L180:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[81])))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v684 = F_visibilitymap_clear(m, v681, v682, int32(2))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L16
	} else {
		goto L181
	}
L181:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	F_ReleaseBuffer(m, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L16
	} else {
		goto L182
	}
L182:
	;
	F_pfree(m, v674)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L16
	} else {
		goto L183
	}
L183:
	;
	goto L177
L184:
	;
	if v695 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v652)+4)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v700 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v845 == int32(0) {
		goto L8
	} else {
		goto L211
	}
L188:
	;
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v719) {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v704+(v700^int32(-1))<<(uint(int32(2))%32))))
	v718 = v710
	goto L188
L190:
	;
	goto L191
L191:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v718 = v712 + v700<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L192:
	;
	v727 = int32(base.Ui32(v719+int32(262120)) >> (uint(int32(2)) % 32))
	goto L194
L193:
	;
	v727 = int32(0)
	goto L194
L194:
	;
	if base.Ui32(v727&int32(65535)) < base.Ui32(v699) {
		goto L3
	} else {
		goto L195
	}
L195:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v699<<(uint(int32(2))%32)+v718)+20))
	if v734&int32(98304) != int32(32768) {
		goto L3
	} else {
		goto L196
	}
L196:
	;
	v741 = v718 + v734&int32(32767)
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)))
	v744 = v742 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)) = uint16(v744)
	v746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v741)+18)))
	v748 = v746 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+18)) = uint16(v748)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+18)) = uint16(v748)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)) = uint16(v744)
	v753 = int32(1)
	v772 = v750<<(uint(v753)%32)&int32(16) | (v750<<(uint(int32(4))%32)&int32(64) | (v750<<(uint(int32(6))%32)&int32(128) | v750&v753<<(uint(int32(12))%32))) | v744
	if v750&int32(15) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)) = uint16(v772)
	goto L199
L198:
	;
	goto L199
L199:
	;
	if v750&int32(16) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v779 = v746 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+18)) = uint16(v779)
	v781 = v779
	goto L202
L201:
	;
	v781 = v748
	goto L202
L202:
	;
	v784 = int32(0)
	if base.B2i32(v772&int32(128) == v784)&base.B2i32(v772&int32(4176) != int32(64)) == v784 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v794 = v781 & int32(49151)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+18)) = uint16(v794)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v796 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v823 = v772
	goto L205
L205:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	v826 = v823 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)) = uint16(v826)
	*(*int32)(unsafe.Add(mBase, uint32(v741)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v741)+4)) = v824
	*(*uint32)(unsafe.Add(mBase, uint32(v718)+4)) = uint32(v651)
	v833 = int64(base.Ui64(v651) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v718))) = uint32(v833)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_MarkBufferDirty(m, v835)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L16
	} else {
		goto L210
	}
L206:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+16)) = uint16(v699)
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+14)) = uint16(v815)
	v819 = int32(base.Ui32(v815) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+12)) = uint16(v819)
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v741)+20)))
	v823 = v821
	goto L205
L207:
	;
	v800 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v800+(v796^int32(-1))<<(uint(int32(6))%32))+16))
	v815 = v806
	goto L206
L208:
	;
	goto L209
L209:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v808+v796<<(uint(int32(6))%32)+int32(-64))+16))
	v815 = v814
	goto L206
L210:
	;
	goto L187
L211:
	;
	F_UnlockReleaseBuffer(m, v845)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L16
	} else {
		goto L212
	}
L212:
	;
	goto L8
L213:
	;
	if v855 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v859 = int32(0)
	v861 = v14 + int32(80)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+72))
	if v864 < v859 {
		v886 = v859
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v953 != 0 {
		goto L243
	} else {
		goto L244
	}
L217:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v890 < int32(0) {
		goto L229
	} else {
		goto L230
	}
L218:
	;
	v889 = v886
	goto L217
L219:
	;
	v870 = v863 + int32(76)
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870))))
	if v871 != int32(1) {
		v886 = v859
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+43)))
	if v874 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v861 == int32(0) {
		v886 = v859
		goto L218
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v861 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v879 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v861))) = v879
	v889 = v879
	goto L217
L225:
	;
	v882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v870)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v861))) = v882
	goto L227
L226:
	;
	goto L227
L227:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v870)+44))
	v886 = v884
	goto L218
L228:
	;
	v909 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v908)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v909) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v894+(v890^int32(-1))<<(uint(int32(2))%32))))
	v908 = v900
	goto L228
L230:
	;
	goto L231
L231:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v908 = v902 + v890<<(uint(int32(13))%32) + int32(-8192)
	goto L228
L232:
	;
	v917 = int32(base.Ui32(v909+int32(262120)) >> (uint(int32(2)) % 32))
	goto L234
L233:
	;
	v917 = int32(0)
	goto L234
L234:
	;
	v920 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850))))
	if base.Ui32(v917&int32(65535)) < base.Ui32(v920) {
		goto L2
	} else {
		goto L235
	}
L235:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v920<<(uint(int32(2))%32)+v908)+20))
	if v925&int32(98304) != int32(32768) {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	v934 = v908 + v925&int32(32767)
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+22)))
	v936 = int32(base.Ui32(v925)>>(uint(int32(17))%32)) - v935
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	if v936 != v937 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	if v936 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v908))) = base.I64_rotr(v851, int64(32))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_MarkBufferDirty(m, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L16
	} else {
		goto L242
	}
L239:
	;
	v940 = F__emscripten_memcpy_bulkmem(m, v934+v935, v889, v936)
	mBase = m.M
	goto L241
L240:
	;
	goto L241
L241:
	;
	goto L238
L242:
	;
	goto L216
L243:
	;
	F_UnlockReleaseBuffer(m, v953)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L16
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v850)+16))
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+12)))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v850)+8))
	F_ProcessCommittedInvalidationMessages(m, v850+int32(20), v958, v959, v960, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L16
	} else {
		goto L247
	}
L246:
	;
	goto L245
L247:
	;
	goto L8
L248:
	;
	F_errmsg_internal(m, int32(227995), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L16
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(497474), int32(480), int32(81574))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L16
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errmsg_internal(m, int32(384140), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L16
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(497474), int32(505), int32(81574))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L16
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F_errmsg_internal(m, int32(237241), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L16
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(497474), int32(380), int32(350672))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L16
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errmsg_internal(m, int32(237241), int32(0))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L16
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(497474), int32(977), int32(287626))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L16
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errmsg_internal(m, int32(237241), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L16
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(497474), int32(1037), int32(316470))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L16
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_errmsg_internal(m, int32(237241), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L16
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(497474), int32(1157), int32(419118))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L16
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errmsg_internal(m, int32(320610), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L16
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(497474), int32(1163), int32(419118))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L16
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_scan_stream_read_next_serial(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
		v11 = int32(-1)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v12 == int32(0) {
			v38 = v11
			v42 = v38
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			if v15 == int32(0) {
				v38 = v11
				v42 = v38
			} else {
				if v10 == int32(1) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v42 = v20
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v21 & int32(-129)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v15 != int32(-1) {
						v31 = base.I32_rem_u_s(v25+v15-int32(1), v12)
						v42 = v31
					} else {
						if v25 != 0 {
							v42 = v25 - int32(1)
						} else {
							v38 = v12 - int32(1)
							v42 = v38
						}
					}
				}
			}
		}
		v43 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v43)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v42
		return v42
	} else {
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
		if v48 == int32(1) {
			v52 = v47 + int32(1)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			if base.Ui32(v52) < base.Ui32(v54) {
				v56 = v52
			} else {
				v56 = int32(0)
			}
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
			if v57&int32(128) != 0 {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ss_report_location(m, v60, v56)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v65 == v56 {
						v67 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v67
						return v67
					} else {
						v71 = int32(-1)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
						if v72 == v71 {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v56
							return v56
						} else {
							v78 = v72 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v78
							if v78 == int32(0) {
								v106 = v71
								*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v106
								return v106
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v56
								return v56
							}
						}
					}
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				if v65 == v56 {
					v67 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v67
					return v67
				} else {
					v71 = int32(-1)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					if v72 == v71 {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v56
						return v56
					} else {
						v78 = v72 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v78
						if v78 == int32(0) {
							v106 = v71
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v106
							return v106
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v56
							return v56
						}
					}
				}
			}
		} else {
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if v84 == v47 {
				v86 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v86
				return v86
			} else {
				v90 = int32(-1)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v91 != v90 {
					v95 = v91 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v95
					if v95 == int32(0) {
						v106 = v90
					} else {
						if v47 != 0 {
							v101 = v47
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v101 = v100
						}
						v106 = v101 - int32(1)
					}
				} else {
					if v47 != 0 {
						v101 = v47
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v101 = v100
					}
					v106 = v101 - int32(1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v106
				return v106
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
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
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
	var v95 int32
	_ = v95
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
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errmsg(m, int32(261175), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L15
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v11 + int32(48)
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v50 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L5:
	;
	v50 = v3
	goto L4
L6:
	;
	goto L7
L7:
	;
	v21 = v3
	v22 = v3
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v21<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+125)))
	if v32 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v50 = v41
	goto L4
L10:
	;
	v43 = v21 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v21 = v43
		v22 = v41
		goto L8
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
	if v35 != int32(112) {
		v41 = v22
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v39 = F_lappend_oid(m, v22, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	return
L16:
	;
	v41 = v39
	goto L10
L17:
	;
	goto L9
L18:
	;
	v56 = F_heap_truncate_find_FKs(m, v50)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	if v56 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v60 <= int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v69 = int32(0)
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v69<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v76
	v82 = F_list_make1_impl(m, int32(472), v11+int32(40))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	goto L2
L24:
	;
	v192 = v69 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v192 < v193 {
		v69 = v192
		goto L22
	} else {
		goto L57
	}
L25:
	;
	v84 = F_heap_truncate_find_FKs(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v84 == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v89 <= v88 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v95 = v88
	goto L29
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v95<<(uint(int32(2))%32))))
	v105 = int32(0)
	if v50 == v105 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L24
L31:
	;
	if v143 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v143 = int32(0)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v111 <= int32(0) {
		v136 = v105
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v143 = v136
	goto L31
L36:
	;
	v114 = int32(0)
	if v114 < v111 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v117 = v111
	goto L39
L38:
	;
	v117 = v114
	goto L39
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v120 = int32(0)
	goto L40
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118+v120<<(uint(int32(2))%32))))
	v129 = base.B2i32(v128 == v104)
	if v128 == v104 {
		v136 = v129
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v136 = v129
	goto L35
L42:
	;
	v131 = v120 + int32(1)
	if v131 != v117 {
		v120 = v131
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v146 = F_get_rel_name(m, v76)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v180 = v95 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v180 < v181 {
		v95 = v180
		goto L29
	} else {
		goto L56
	}
L47:
	;
	v148 = F_get_rel_name(m, v104)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(89418), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v148
	F_errdetail(m, int32(649249), v11+int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v148
	F_errhint(m, int32(643663), v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(495443), int32(3749), int32(174750))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	goto L30
L57:
	;
	goto L23
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v148
	F_errdetail(m, int32(610328), v11)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(495443), int32(3740), int32(174750))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_tuple_should_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v14 = int32(768)
	if v13&v14 == v14 {
		v50 = v5
		v51 = v13
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v18) < base.Ui32(int32(3)) {
			v50 = v5
			v51 = v13
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v21))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v18)) == int32(0) {
				v33 = base.B2i32(base.Ui32(v18) < base.Ui32(v21))
			} else {
				v33 = int32(base.Ui32(v18-v21) >> (uint(int32(31)) % 32))
			}
			if v33 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			} else {
			}
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v35))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v18)) == int32(0) {
				v47 = base.B2i32(base.Ui32(v18) < base.Ui32(v35))
			} else {
				v47 = int32(base.Ui32(v18-v35) >> (uint(int32(31)) % 32))
			}
			v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			v50 = v47
			v51 = v48
		}
	}
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51&int32(4096) == int32(0) {
		if base.Ui32(v52) < base.Ui32(int32(3)) {
			v182 = v50
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v59))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v52)) == int32(0) {
				v71 = base.B2i32(base.Ui32(v52) < base.Ui32(v59))
			} else {
				v71 = int32(base.Ui32(v52-v59) >> (uint(int32(31)) % 32))
			}
			if v71 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
			} else {
			}
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v73))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v52)) == int32(0) {
				v85 = base.B2i32(base.Ui32(v52) < base.Ui32(v73))
			} else {
				v85 = int32(base.Ui32(v52-v73) >> (uint(int32(31)) % 32))
			}
			v182 = v85 | v50
		}
		v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
		if base.Ui32(v185) < base.Ui32(int32(16384)) {
			v208 = v182
		} else {
			v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v188) < base.Ui32(int32(3)) {
				v208 = v182
			} else {
				v191 = int32(1)
				v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
					v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
				} else {
					v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
				}
				if v204 == int32(0) {
					v208 = v191
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
					v208 = v191
				}
			}
		}
		m.G0 = v11 + int32(16)
		return v208 & int32(1)
	} else {
		if v52 == int32(0) {
			v182 = v50
			v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			if base.Ui32(v185) < base.Ui32(int32(16384)) {
				v208 = v182
			} else {
				v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v188) < base.Ui32(int32(3)) {
					v208 = v182
				} else {
					v191 = int32(1)
					v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
						v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
					} else {
						v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
					}
					if v204 == int32(0) {
						v208 = v191
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
						v208 = v191
					}
				}
			}
			m.G0 = v11 + int32(16)
			return v208 & int32(1)
		} else {
			if v51&int32(4304) == int32(4224) {
				v93 = int32(1)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if int32(base.Ui32(v52-v94)>>(uint(int32(31))%32)) == int32(0) {
					v182 = v93
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v52
					v182 = v93
				}
				v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				if base.Ui32(v185) < base.Ui32(int32(16384)) {
					v208 = v182
				} else {
					v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v188) < base.Ui32(int32(3)) {
						v208 = v182
					} else {
						v191 = int32(1)
						v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
							v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
						} else {
							v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
						}
						if v204 == int32(0) {
							v208 = v191
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
							v208 = v191
						}
					}
				}
				m.G0 = v11 + int32(16)
				return v208 & int32(1)
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if int32(base.Ui32(v52-v101)>>(uint(int32(31))%32)) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v52
				} else {
				}
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v111 = int32(base.Ui32(v52-v107)>>(uint(int32(31))%32)) | v50
				v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				v124 = F_GetMultiXactIdMembers(m, v52, v11+int32(12), int32(base.Ui32(v114&int32(128))>>(uint(int32(7))%32))|base.B2i32(v114&int32(4176) == int32(64)))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					if v124 <= int32(0) {
						v182 = v111
						v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						if base.Ui32(v185) < base.Ui32(int32(16384)) {
							v208 = v182
						} else {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v188) < base.Ui32(int32(3)) {
								v208 = v182
							} else {
								v191 = int32(1)
								v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
									v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
								} else {
									v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
								}
								if v204 == int32(0) {
									v208 = v191
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
									v208 = v191
								}
							}
						}
						m.G0 = v11 + int32(16)
						return v208 & int32(1)
					} else {
						v135 = v111
						v136 = int32(0)
						for {
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v139 = int32(3)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v136<<(uint(v139)%32))))
							v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v143))&base.B2i32(base.Ui32(v139) <= base.Ui32(v142)) == int32(0) {
								v155 = base.B2i32(base.Ui32(v142) < base.Ui32(v143))
							} else {
								v155 = int32(base.Ui32(v142-v143) >> (uint(int32(31)) % 32))
							}
							if v155 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v142
							} else {
							}
							v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v157))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v142)) == int32(0) {
								v169 = base.B2i32(base.Ui32(v142) < base.Ui32(v157))
							} else {
								v169 = int32(base.Ui32(v142-v157) >> (uint(int32(31)) % 32))
							}
							v170 = v169 | v135
							v172 = v136 + int32(1)
							if v172 != v124 {
								v135 = v170
								v136 = v172
								continue
							} else {
								break
							}
							break
						}
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						F_pfree(m, v174)
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							v182 = v170
							v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							if base.Ui32(v185) < base.Ui32(int32(16384)) {
								v208 = v182
							} else {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if base.Ui32(v188) < base.Ui32(int32(3)) {
									v208 = v182
								} else {
									v191 = int32(1)
									v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
										v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
									} else {
										v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
									}
									if v204 == int32(0) {
										v208 = v191
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
										v208 = v191
									}
								}
							}
							m.G0 = v11 + int32(16)
							return v208 & int32(1)
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+224)) = v133
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v149 != 0 {
		goto L38
	} else {
		goto L39
	}
L10:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v131)
	v133 = v70
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
	v115 = v32
	v119 = v19
	goto L13
L13:
	;
	if base.Ui32(v115) <= base.Ui32(v119) {
		goto L35
	} else {
		goto L36
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)) = uint8(v104)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+228)) = v70
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = v108
	v112 = base.B2i32(base.Ui32(int32(31)) < base.Ui32(v56-v17))
	if v47&v112 != 0 {
		goto L10
	} else {
		goto L31
	}
L15:
	;
	v56 = v44
	goto L17
L16:
	;
	v104 = v93 ^ int32(1)
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
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v93 == int32(0) {
		goto L27
	} else {
		goto L28
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+240)) = v75 + int32(4096)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v83 = v73 & int32(1)
	v84 = int32(0)
	if v56 == v20-int32(2) {
		v104 = v84
		goto L14
	} else {
		goto L23
	}
L23:
	;
	if v83 == int32(0) {
		v104 = v84
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v88 != int32(1) {
		v104 = v84
		goto L14
	} else {
		goto L25
	}
L25:
	;
	if v73&int32(2) != 0 {
		v56 = v70
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v97 == int32(0) {
		v44 = v70
		v47 = int32(1)
		goto L15
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L16
L30:
	;
	goto L29
L31:
	;
	if base.Ui32(int32(31)) < base.Ui32(v56-v17) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v114 = v70
	goto L34
L33:
	;
	v114 = v19
	goto L34
L34:
	;
	v115 = v70
	v119 = v114
	goto L13
L35:
	;
	v133 = v119
	goto L9
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+224)) = v119
	v129 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v129)
	goto L8
L38:
	;
	v150 = int32(2)
	goto L40
L39:
	;
	v150 = int32(0)
	goto L40
L40:
	;
	v151 = v146 | v150
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v151)
	goto L8
}
func F_heap_xlog_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
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
	var v456 int64
	_ = v456
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v632 int64
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	v3 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(8240)
	m.G0 = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	F_XLogRecGetBlockTag(m, l0, v3, v24+int32(8228), v3, v24+int32(8220))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v38 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if v45 < int32(1) {
		v69 = v38
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	if v69 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	goto L3
L5:
	;
	v51 = v44 + int32(128)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v52 != int32(1) {
		v69 = v38
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
	if v24+int32(8224) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83]))) = v61
	goto L15
L14:
	;
	goto L15
L15:
	;
	v69 = int32(1)
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83]))) = v70
	goto L19
L18:
	;
	goto L19
L19:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	if v75&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[85])))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v82
	v86 = F_CreateFakeRelcacheEntry(m, v24+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v107 = v70
	goto L22
L22:
	;
	v109 = int32(base.Ui32(v70) >> (uint(int32(16)) % 32))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83])))
	v115 = F_XLogReadBufferForRedo(m, l0, base.B2i32(v111 != v107), v24+int32(8216))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L31
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83])))
	F_visibilitymap_pin(m, v86, v90, v24+int32(48))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83])))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
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
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	v107 = v105
	goto L22
L28:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L199
	}
L29:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L196
	}
L30:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L193
	}
L31:
	;
	if v115 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[86])))
	if v119 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v250 = int32(0)
	v254 = v3
	goto L34
L34:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[83])))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	if v255 == v256 {
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
	v123 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+(v119^int32(-1))<<(uint(int32(2))%32))))
	v137 = v129
	goto L35
L37:
	;
	goto L38
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v137 = v131 + v119<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v146 = int32(base.Ui32(v138+int32(262120)) >> (uint(int32(2)) % 32))
	goto L41
L40:
	;
	v146 = int32(0)
	goto L41
L41:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if base.Ui32(v146&int32(65535)) < base.Ui32(v149) {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149<<(uint(int32(2))%32)+v137)+20))
	if v154&int32(98304) != int32(32768) {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v161 = v137 + v154&int32(32767)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+20)))
	v164 = v162 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+20)) = uint16(v164)
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)))
	v168 = v166 & int32(-24577)
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v171 = v168 | int32(16384)
	goto L46
L45:
	;
	v171 = v168
	goto L46
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)) = uint16(v171)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)) = uint16(v171)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+20)) = uint16(v164)
	v176 = int32(1)
	v195 = v173<<(uint(v176)%32)&int32(16) | (v173<<(uint(int32(4))%32)&int32(64) | (v173<<(uint(int32(6))%32)&int32(128) | v173&v176<<(uint(int32(12))%32))) | v164
	if v173&int32(15) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+20)) = uint16(v195)
	goto L49
L48:
	;
	goto L49
L49:
	;
	if v173&int32(16) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v202 = v171 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)) = uint16(v202)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v206 = v195 & int32(13279)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+20)) = uint16(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = v204
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+16)) = uint16(v74)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+14)) = uint16(v70)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+12)) = uint16(v109)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v214 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	if v234&int32(1) != 0 {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v214))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v216)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v232
	goto L53
L57:
	;
	if v228 == int32(0) {
		goto L53
	} else {
		goto L61
	}
L58:
	;
	v228 = base.B2i32(base.Ui32(v216) < base.Ui32(v214))
	goto L57
L59:
	;
	goto L60
L60:
	;
	v228 = int32(base.Ui32(v216-v214) >> (uint(int32(31)) % 32))
	goto L57
L61:
	;
	goto L56
L62:
	;
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+10)))
	v239 = v237 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+10)) = uint16(v239)
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v137))) = base.I64_rotr(v26, int64(32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[86])))
	F_MarkBufferDirty(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v250 = v161
	v254 = int32(base.Ui32(v154) >> (uint(int32(17)) % 32))
	goto L34
L66:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	if v335&int32(2) != 0 {
		goto L89
	} else {
		goto L90
	}
L67:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[86])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[87]))) = v258
	v334 = v115
	goto L66
L68:
	;
	goto L69
L69:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v261 = int32(*(*int8)(unsafe.Add(mBase, uint32(v260)+48)))
	if v261 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v265 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v331 = F_XLogReadBufferForRedo(m, l0, int32(0), v24+int32(8212))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L88
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[87]))) = v265
	if v265 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v285&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v271+(v265^int32(-1))<<(uint(int32(2))%32))))
	v285 = v277
	goto L74
L76:
	;
	goto L77
L77:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v285 = v279 + v265<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	v334 = int32(0)
	goto L66
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+10)) = int32(1572864)
	v319 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v285)+18)) = uint16(v319)
	v325 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v285)+16)) = uint16(v325)
	*(*uint16)(unsafe.Add(mBase, uint32(v285)+14)) = uint16(v325)
	goto L78
L80:
	;
	v313 = F___memset(m, v285, int32(0), int32(8192))
	mBase = m.M
	goto L79
L81:
	;
	goto L80
L88:
	;
	v334 = v331
	goto L66
L89:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v340
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[85])))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v342
	v346 = F_CreateFakeRelcacheEntry(m, v24+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v366 = int32(0)
	if v334 == v366 {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	F_visibilitymap_pin(m, v346, v350, v24+int32(48))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v358 = F_visibilitymap_clear(m, v355, v356, int32(3))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	F_ReleaseBuffer(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_pfree(m, v346)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	v369 = int32(0)
	v371 = v24 + int32(44)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+72))
	if v374 < v369 {
		v396 = v369
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v603 = v366
	goto L99
L99:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[86])))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[87])))
	if v614 != 0 {
		goto L179
	} else {
		goto L180
	}
L100:
	;
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[87])))
	if v402 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v399 = v396
	goto L100
L102:
	;
	v380 = v373 + int32(76)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if v381 != int32(1) {
		v396 = v369
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+43)))
	if v384 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v371 == int32(0) {
		v396 = v369
		goto L101
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v371 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v389
	v399 = v389
	goto L100
L108:
	;
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v380)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v392
	goto L110
L109:
	;
	goto L110
L110:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v380)+44))
	v396 = v394
	goto L101
L111:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	if base.Ui32(v421) < base.Ui32(int32(25)) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406+(v402^int32(-1))<<(uint(int32(2))%32))))
	v420 = v412
	goto L111
L113:
	;
	goto L114
L114:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v420 = v414 + v402<<(uint(int32(13))%32) + int32(-8192)
	goto L111
L115:
	;
	v430 = int32(1)
	goto L117
L116:
	;
	v430 = int32(base.Ui32(v421+int32(262120))>>(uint(int32(2))%32)) + int32(1)
	goto L117
L117:
	;
	if base.Ui32(v430&int32(65535)) < base.Ui32(v400) {
		goto L29
	} else {
		goto L118
	}
L118:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v435 = int32(0)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	if v437&int32(32) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399))))
	v443 = v440
	v444 = v399 + int32(2)
	goto L121
L120:
	;
	v443 = v435
	v444 = v399
	goto L121
L121:
	;
	if v437&int32(64) != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444))))
	v451 = v448
	v452 = v444 + int32(2)
	goto L124
L123:
	;
	v451 = v435
	v452 = v444
	goto L124
L124:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+4)))
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+2)))
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	v456 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+63)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v456
	v463 = v452 + int32(5)
	v464 = v399 + v434 - v463
	v466 = v24 + int32(71)
	if v443 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v451 != 0 {
		goto L145
	} else {
		goto L146
	}
L126:
	;
	v468 = v453 - int32(23)
	if v468 != 0 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	if v464 != 0 {
		goto L142
	} else {
		goto L143
	}
L129:
	;
	v473 = v24 + int32(48) + v453
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+22)))
	if v443 != 0 {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v469 = F__emscripten_memcpy_bulkmem(m, v466, v463, v468)
	mBase = m.M
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L129
L133:
	;
	v478 = v477 + v443
	v480 = v464 - v468
	if v480 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v476 = F__emscripten_memcpy_bulkmem(m, v473, v250+v474, v443)
	mBase = m.M
	v477 = v476
	goto L136
L135:
	;
	v477 = v473
	goto L136
L136:
	;
	goto L133
L137:
	;
	v489 = v482 + v480
	goto L125
L138:
	;
	v481 = F__emscripten_memcpy_bulkmem(m, v478, v463+v468, v480)
	mBase = m.M
	v482 = v481
	goto L140
L139:
	;
	v482 = v478
	goto L140
L140:
	;
	goto L137
L141:
	;
	v489 = v485 + v464
	goto L125
L142:
	;
	v484 = F__emscripten_memcpy_bulkmem(m, v466, v463, v464)
	mBase = m.M
	v485 = v484
	goto L144
L143:
	;
	v485 = v466
	goto L144
L144:
	;
	goto L141
L145:
	;
	if v451 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+70)) = uint8(v453)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+66)) = uint16(v455)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+36))
	v499 = v454 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+68)) = uint16(v499)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v497
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+64)) = uint16(v74)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+62)) = uint16(v70)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+60)) = uint16(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v504
	v516 = F_PageAddItemExtended(m, v420, v24+int32(48), v443+v451+v464+int32(23), v400, int32(3))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L152
	}
L148:
	;
	goto L147
L149:
	;
	v492 = F__emscripten_memcpy_bulkmem(m, v489, v250+v254-v451, v451)
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	if v516 == int32(0) {
		goto L28
	} else {
		goto L153
	}
L153:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	if v520&int32(2) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+10)))
	v525 = v523 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v420)+10)) = uint16(v525)
	goto L156
L155:
	;
	goto L156
L156:
	;
	v530 = int32(4)
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+14)))
	v532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	v533 = v531 - v532
	if v533 <= v530 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v420))) = base.I64_rotr(v26, int64(32))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[87])))
	F_MarkBufferDirty(m, v599)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L176
	}
L158:
	;
	v536 = v530
	goto L160
L159:
	;
	v536 = v533
	goto L160
L160:
	;
	v538 = v536 - int32(4)
	if v538 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v595 = int32(0)
	goto L157
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v532) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v595 = v538
	goto L157
L165:
	;
	v549 = int32(base.Ui32(v532+int32(262120)) >> (uint(int32(2)) % 32))
	goto L167
L166:
	;
	v549 = int32(0)
	goto L167
L167:
	;
	if base.Ui32(v549&int32(65535)) < base.Ui32(int32(291)) {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+10)))
	if v554&int32(1) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v595 = int32(0)
	goto L157
L170:
	;
	goto L171
L171:
	;
	v563 = int32(1)
	goto L172
L172:
	;
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563&int32(65535)<<(uint(int32(2))%32)+(v420+int32(24))-int32(3)))))
	if v574&int32(384) == int32(0) {
		goto L164
	} else {
		goto L174
	}
L173:
	;
	v595 = int32(0)
	goto L157
L174:
	;
	v580 = v563 + int32(1)
	v581 = int32(65535)
	if base.Ui32(v580&v581) <= base.Ui32(v549&v581) {
		v563 = v580
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v603 = v595
	goto L99
L177:
	;
	if l1 != 0 {
		goto L188
	} else {
		goto L189
	}
L178:
	;
	F_UnlockReleaseBuffer(m, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L187
	}
L179:
	;
	if v613 == v614 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v619 = v613
	goto L181
L181:
	;
	if v619 == int32(0) {
		goto L177
	} else {
		goto L186
	}
L182:
	;
	v622 = v614
	goto L178
L183:
	;
	goto L184
L184:
	;
	F_UnlockReleaseBuffer(m, v614)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[86])))
	v619 = v618
	goto L181
L186:
	;
	v622 = v619
	goto L178
L187:
	;
	goto L177
L188:
	;
	m.G0 = v24 + int32(8240)
	return
L189:
	;
	if v334 != 0 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	if base.Ui32(int32(1637)) < base.Ui32(v603) {
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v630
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[85])))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[82])))
	F_XLogRecordPageWithFreeSpace(m, v24, v634, v603)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	goto L188
L193:
	;
	F_errmsg_internal(m, int32(237241), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(497474), int32(763), int32(355487))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errmsg_internal(m, int32(227995), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(497474), int32(841), int32(355487))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errmsg_internal(m, int32(384140), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(497474), int32(919), int32(355487))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
