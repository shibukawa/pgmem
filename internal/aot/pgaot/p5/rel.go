package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecRelGenVirtualNotNull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v11 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = int32(_a_F_ExecRelGenVirtualNotNull_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0])) = v18
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v95 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0])) = v16
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v22 = F_palloc0_mul(m, int32(4), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v70 = F_palloc0_mul(m, int32(4), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v22
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v27 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v35 = v5
	goto L11
L11:
	;
	v41 = v35 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v46 = F_palloc0(m, int32(20))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L4
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(52)
	v50 = F_build_generation_expression(m, v14, v44)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = int32(-1)
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v50
	v59 = F_ExecPrepareExpr(m, v46, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v61+v41))) = v59
	v65 = v35 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v65 < v66 {
		v35 = v65
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v70
	goto L4
L18:
	;
	v98 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v100 = v95
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = l1
	if l3 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v100 = v98
	goto L20
L22:
	;
	return int32(0)
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v106 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	goto L27
L27:
	;
	v113 = int32(0)
	goto L29
L28:
	;
	return base.I32_extend16_s(v126)
L29:
	;
	v123 = v113 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123+v124)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127+v123)))
	v130 = F_ExecCheck(m, v129, v100)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	if v130 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v135 = v113 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v135 < v136 {
		v113 = v135
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
}
func F_rel_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v548 int32
	_ = v548
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
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
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
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v839 int32
	_ = v839
	var v868 int32
	_ = v868
	v5 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 != 0 {
		v868 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v868
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	switch v15 {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		v868 = v5
		goto L1
	}
L3:
	;
	v868 = int32(1)
	goto L1
L4:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v395+v396<<(uint(int32(2))%32))))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+36))
	if l2 == int32(0) {
		v481 = v5
		goto L109
	} else {
		goto L110
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v16 == int32(0) {
		v375 = v5
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v394 != 0 {
		goto L3
	} else {
		goto L108
	}
L7:
	;
	v394 = v375
	goto L6
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v19 == int32(0) {
		v63 = l2
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v63 == int32(0) {
		v375 = v5
		goto L7
	} else {
		goto L23
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		v63 = l2
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = l2
	v31 = v5
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v31<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	if v43 == int32(0) {
		v56 = v27
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v63 = v56
	goto L9
L14:
	;
	v58 = v31 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v58 < v59 {
		v27 = v56
		v31 = v58
		goto L12
	} else {
		goto L22
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v47 != 0 {
		v56 = v27
		goto L14
	} else {
		goto L19
	}
L17:
	;
	v50 = int32(1)
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+120)) = uint8(v50)
	v52 = F_lappend(m, v27, v42)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v50 = int32(0)
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	v56 = v52
	goto L14
L22:
	;
	goto L13
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v76 == int32(0) {
		v375 = v5
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		v375 = v5
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v89 = v5
	goto L26
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v89<<(uint(int32(2))%32))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+101)))
	if v100 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v375 = int32(0)
	goto L7
L28:
	;
	v364 = v89 + int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v364 < v365 {
		v89 = v364
		goto L26
	} else {
		goto L107
	}
L29:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+103)))
	if v103 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v99)+88))
	if v106 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+40))
	if v107 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v335 != v337 {
		goto L28
	} else {
		goto L105
	}
L33:
	;
	v110 = int32(0)
	v335 = v107
	v337 = v110
	v341 = v110
	goto L32
L34:
	;
	goto L35
L35:
	;
	v112 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v114 <= v112 {
		v335 = v107
		v337 = v112
		v341 = v112
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v123 = v112
	v127 = v112
	goto L37
L37:
	;
	v135 = int32(0)
	goto L39
L38:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v99)+40))
	v335 = v330
	v337 = v123
	v341 = v127
	goto L32
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v145 = int32(2)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v135<<(uint(v145)%32))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+96))
	v151 = v123 << (uint(v145) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v99)+52))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151+v152)))
	v155 = int32(0)
	if v149 == v155 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L38
L41:
	;
	v327 = v135 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v327 < v328 {
		v135 = v327
		goto L39
	} else {
		goto L104
	}
L42:
	;
	if v193 == int32(0) {
		goto L41
	} else {
		goto L55
	}
L43:
	;
	v193 = int32(0)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v161 <= int32(0) {
		v187 = v155
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v193 = v187
	goto L42
L47:
	;
	v164 = int32(0)
	if v164 < v161 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v167 = v161
	goto L50
L49:
	;
	v167 = v164
	goto L50
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v170 = int32(0)
	goto L51
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168+v170<<(uint(int32(2))%32))))
	v179 = base.B2i32(v178 == v154)
	if v178 == v154 {
		v187 = v179
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v187 = v179
	goto L46
L53:
	;
	v181 = v170 + int32(1)
	if v181 != v167 {
		v170 = v181
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v99)+48))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196+v151)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v200 = int32(0)
	if v199 == v200 {
		v228 = v200
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v229 = F_collations_agree_on_equality(m, v198, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L66
	}
L57:
	;
	goto L56
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v207 = v205 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v207) {
		v228 = v200
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v211 = int32(1) << (uint(v207) % 32)
	if v211&int32(3904) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223+v199)))
	v228 = v225
	goto L57
L61:
	;
	if v211&int32(5) != 0 {
		v223 = int32(16)
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v223 = int32(24)
	goto L60
L64:
	;
	if v207 != int32(30) {
		v228 = v200
		goto L57
	} else {
		goto L65
	}
L65:
	;
	v223 = int32(12)
	goto L60
L66:
	;
	if v229 == int32(0) {
		goto L41
	} else {
		goto L67
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+28))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+120)))
	if v235 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v252 = F_match_index_to_operand(m, v251, v123, v99)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L20
	} else {
		goto L77
	}
L69:
	;
	v238 = int32(0)
	if v234 == v238 {
		v251 = v238
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v234 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v241 < int32(2) {
		v251 = v238
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v251 = v245
	goto L68
L74:
	;
	v251 = int32(0)
	goto L68
L75:
	;
	goto L76
L76:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v251 = v250
	goto L68
L77:
	;
	if v252 == int32(0) {
		goto L41
	} else {
		goto L78
	}
L78:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	v257 = int32(0)
	if v256 == v257 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v302 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L80:
	;
	v302 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v265 = int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v266 <= v265 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v269 = v265
	goto L85
L84:
	;
	v269 = v266
	goto L85
L85:
	;
	v273 = int32(0)
	v275 = v257
	goto L86
L86:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v256+int32(8)+v273<<(uint(int32(2))%32))))
	if v282 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v302 = v294
	goto L79
L88:
	;
	goto L87
L89:
	;
	v283 = int32(2)
	if v275 != 0 {
		v294 = v283
		goto L88
	} else {
		goto L92
	}
L90:
	;
	v289 = v275
	goto L91
L91:
	;
	v291 = v273 + int32(1)
	if v291 != v269 {
		v273 = v291
		v275 = v289
		goto L86
	} else {
		goto L94
	}
L92:
	;
	v284 = int32(1)
	if base.Ui32(v284) < base.Ui32(base.I32_popcnt(v282)) {
		v294 = v283
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v289 = v284
	goto L91
L94:
	;
	v294 = v289
	goto L88
L95:
	;
	v305 = int32(_a_F_rel_is_distinct_for_0)
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_rel_is_distinct_for[0]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, _c_F_rel_is_distinct_for[0])) = v308
	if l3 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v316 = v127
	goto L97
L97:
	;
	v318 = v123 + int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v99)+40))
	if v319 <= v318 {
		v335 = v319
		v337 = v318
		v341 = v316
		goto L32
	} else {
		goto L102
	}
L98:
	;
	v310 = F_lappend(m, v127, v148)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L20
	} else {
		goto L101
	}
L99:
	;
	v312 = v127
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rel_is_distinct_for[0])) = v306
	v316 = v312
	goto L97
L101:
	;
	v312 = v310
	goto L100
L102:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v321 <= int32(0) {
		v335 = v319
		v337 = v318
		v341 = v316
		goto L32
	} else {
		goto L103
	}
L103:
	;
	v123 = v318
	v127 = v316
	goto L37
L104:
	;
	goto L40
L105:
	;
	if l3 == int32(0) {
		v375 = int32(1)
		goto L7
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v341
	v394 = int32(1)
	goto L6
L107:
	;
	goto L27
L108:
	;
	v868 = v5
	goto L1
L109:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v401)+120))
	if v489 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L110:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v404 <= int32(0) {
		v481 = v5
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v409 = int32(0)
	v413 = v5
	goto L112
L112:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v421+v409<<(uint(int32(2))%32))))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+28))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+120)))
	if v428 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v481 = v471
	goto L109
L114:
	;
	v473 = v409 + int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v473 < v474 {
		v409 = v473
		v413 = v471
		goto L112
	} else {
		goto L132
	}
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v441 == int32(0) {
		v471 = v413
		goto L114
	} else {
		goto L122
	}
L116:
	;
	if v427 == int32(0) {
		v471 = v413
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v427 == int32(0) {
		v471 = v413
		goto L114
	} else {
		goto L121
	}
L119:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v431 < int32(2) {
		v471 = v413
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v440 = v434 + int32(4)
	goto L115
L121:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v440 = v439
	goto L115
L122:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v444 == int32(27) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v447 == int32(0) {
		v471 = v413
		goto L114
	} else {
		goto L126
	}
L124:
	;
	v451 = v441
	v452 = v444
	goto L125
L125:
	;
	if v452 != int32(6) {
		v471 = v413
		goto L114
	} else {
		goto L127
	}
L126:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v451 = v447
	v452 = v450
	goto L125
L127:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v455 != v396 {
		v471 = v413
		goto L114
	} else {
		goto L128
	}
L128:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+28))
	if v457 != 0 {
		v471 = v413
		goto L114
	} else {
		goto L129
	}
L129:
	;
	v459 = F_palloc(m, int32(12))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L20
	} else {
		goto L130
	}
L130:
	;
	v461 = int32(*(*int16)(unsafe.Add(mBase, uint32(v451)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v461
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+8)) = v465
	v467 = F_lappend(m, v413, v459)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L20
	} else {
		goto L131
	}
L131:
	;
	v471 = v467
	goto L114
L132:
	;
	goto L113
L133:
	;
	if v839 == int32(0) {
		v868 = v5
		goto L1
	} else {
		goto L220
	}
L134:
	;
	v839 = v817
	goto L133
L135:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+38)))
	if v586 != 0 {
		v839 = int32(0)
		goto L133
	} else {
		goto L159
	}
L136:
	;
	v492 = int32(1)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+38)))
	if v493 == v492 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+40)))
	if v496 != 0 {
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v497 <= int32(0) {
		v817 = v492
		goto L134
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v506 = v5
	goto L142
L142:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513+v506<<(uint(int32(2))%32))))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v401)+76))
	v519 = F_get_sortgroupclause_tle(m, v517, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L20
	} else {
		goto L144
	}
L143:
	;
	v817 = v567
	goto L134
L144:
	;
	if v481 == int32(0) {
		goto L135
	} else {
		goto L145
	}
L145:
	;
	v523 = int32(*(*int16)(unsafe.Add(mBase, uint32(v519)+8)))
	v524 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v524 < v525 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v529 = v525
	goto L148
L147:
	;
	v529 = v524
	goto L148
L148:
	;
	v534 = v524
	goto L149
L149:
	;
	if v534 == v529 {
		goto L135
	} else {
		goto L151
	}
L150:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	v555 = F_equality_ops_are_compatible(m, v553, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L20
	} else {
		goto L153
	}
L151:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v534<<(uint(int32(2))%32)+v548)))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	if v551 != v523 {
		v534 = v534 + int32(1)
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	if v555 == int32(0) {
		goto L135
	} else {
		goto L154
	}
L154:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v561 = F_exprCollation(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L20
	} else {
		goto L155
	}
L155:
	;
	v563 = F_collations_agree_on_equality(m, v559, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	if v563 == int32(0) {
		goto L135
	} else {
		goto L157
	}
L157:
	;
	v567 = int32(1)
	v569 = v506 + v567
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v569 < v570 {
		v506 = v569
		goto L142
	} else {
		goto L158
	}
L158:
	;
	goto L143
L159:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v401)+108))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v401)+100))
	if v588 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v695 = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v401)+144))
	if v696 == v695 {
		v839 = v695
		goto L133
	} else {
		goto L191
	}
L161:
	;
	if v587 != 0 {
		v839 = int32(0)
		goto L133
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v587 != 0 {
		goto L183
	} else {
		goto L184
	}
L164:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v591 <= int32(0) {
		v817 = int32(1)
		goto L134
	} else {
		goto L165
	}
L165:
	;
	v601 = int32(0)
	goto L166
L166:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v588)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608+v601<<(uint(int32(2))%32))))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v401)+76))
	v614 = F_get_sortgroupclause_tle(m, v612, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L20
	} else {
		goto L168
	}
L167:
	;
	v817 = v662
	goto L134
L168:
	;
	if v481 == int32(0) {
		goto L160
	} else {
		goto L169
	}
L169:
	;
	v618 = int32(*(*int16)(unsafe.Add(mBase, uint32(v614)+8)))
	v619 = int32(0)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v619 < v620 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v624 = v620
	goto L172
L171:
	;
	v624 = v619
	goto L172
L172:
	;
	v629 = v619
	goto L173
L173:
	;
	if v629 == v624 {
		goto L160
	} else {
		goto L175
	}
L174:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v645)+4))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v612)+8))
	v650 = F_equality_ops_are_compatible(m, v648, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L20
	} else {
		goto L177
	}
L175:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v629<<(uint(int32(2))%32)+v643)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	if v646 != v618 {
		v629 = v629 + int32(1)
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	if v650 == int32(0) {
		goto L160
	} else {
		goto L178
	}
L178:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v645)+8))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v656 = F_exprCollation(m, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L20
	} else {
		goto L179
	}
L179:
	;
	v658 = F_collations_agree_on_equality(m, v654, v656)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L20
	} else {
		goto L180
	}
L180:
	;
	if v658 == int32(0) {
		goto L160
	} else {
		goto L181
	}
L181:
	;
	v662 = int32(1)
	v664 = v601 + v662
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v664 < v665 {
		v601 = v664
		goto L166
	} else {
		goto L182
	}
L182:
	;
	goto L167
L183:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+104)))
	if v668 != 0 {
		v839 = int32(1)
		goto L133
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v679 = int32(1)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+36)))
	if v680 != 0 {
		v817 = v679
		goto L134
	} else {
		goto L189
	}
L186:
	;
	v669 = int32(0)
	v672 = F_expand_grouping_sets(m, v587, v669, int32(-1))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L20
	} else {
		goto L187
	}
L187:
	;
	if v672 == int32(0) {
		v817 = v669
		goto L134
	} else {
		goto L188
	}
L188:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v839 = base.B2i32(v676 == int32(1))
	goto L133
L189:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v401)+112))
	if v681 != 0 {
		v817 = v679
		goto L134
	} else {
		goto L190
	}
L190:
	;
	goto L160
L191:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696)+8)))
	if v700 != 0 {
		v817 = int32(0)
		goto L134
	} else {
		goto L192
	}
L192:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v696)+32))
	if v701 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+12))
	v704 = v702
	goto L195
L194:
	;
	v704 = int32(0)
	goto L195
L195:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v401)+76))
	if v706 == int32(0) {
		v839 = int32(1)
		goto L133
	} else {
		goto L196
	}
L196:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	if v710 <= int32(0) {
		v817 = int32(1)
		goto L134
	} else {
		goto L197
	}
L197:
	;
	v715 = v704
	v721 = int32(0)
	goto L198
L198:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v706)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v721<<(uint(int32(2))%32))))
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+26)))
	if v732 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v817 = v808
	goto L134
L200:
	;
	v735 = int32(0)
	if v481 == v735 {
		v817 = v735
		goto L134
	} else {
		goto L203
	}
L201:
	;
	v796 = v715
	goto L202
L202:
	;
	v808 = int32(1)
	v810 = v721 + v808
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	if v810 < v811 {
		v715 = v796
		v721 = v810
		goto L198
	} else {
		goto L219
	}
L203:
	;
	v739 = v715 + int32(4)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v696)+32))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	if base.Ui32(v739) < base.Ui32(v742+v743<<(uint(int32(2))%32)) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v748 = v739
	goto L206
L205:
	;
	v748 = int32(0)
	goto L206
L206:
	;
	v749 = int32(*(*int16)(unsafe.Add(mBase, uint32(v731)+8)))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v751 = int32(0)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v751 < v752 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v756 = v752
	goto L209
L208:
	;
	v756 = v751
	goto L209
L209:
	;
	v759 = int32(0)
	goto L210
L210:
	;
	if v759 == v756 {
		v817 = v751
		goto L134
	} else {
		goto L212
	}
L211:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v750)+8))
	v783 = F_equality_ops_are_compatible(m, v781, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L20
	} else {
		goto L214
	}
L212:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v759<<(uint(int32(2))%32)+v776)))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	if v779 != v749 {
		v759 = v759 + int32(1)
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	if v783 == int32(0) {
		v817 = v751
		goto L134
	} else {
		goto L215
	}
L215:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v778)+8))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	v789 = F_exprCollation(m, v788)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L20
	} else {
		goto L216
	}
L216:
	;
	v791 = F_collations_agree_on_equality(m, v787, v789)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L20
	} else {
		goto L217
	}
L217:
	;
	if v791 == int32(0) {
		v817 = v751
		goto L134
	} else {
		goto L218
	}
L218:
	;
	v796 = v748
	goto L202
L219:
	;
	goto L199
L220:
	;
	goto L3
}
func F_truncate_check_rel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = l1 + int32(4)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+119)))
	switch v11 - int32(102) {
	case 0:
		v14 = F_GetForeignServerIdByRelId(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = F_GetFdwRoutineByServerId(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+140))
				if v18 != 0 {
					v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[0])))
					if v54 != 0 {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						v56 = int32(1)
						if base.Ui32(l0) < base.Ui32(int32(_a_F_truncate_check_rel_0)) {
							v64 = v56
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
							if v59 == int32(99) {
								v64 = v56
							} else {
								v62 = F_isTempToastNamespace(m, v59)
								mBase = m.M
								v64 = v62
							}
						}
						if v64 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							if v76 != 0 {
								v79 = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
								m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[2])))
							if v68 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
										F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2445), int32(_a_F_truncate_check_rel_3))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
								if l0 == int32(2613) {
									v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
									if v76 != 0 {
										v79 = int32(0)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
										m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									} else {
										m.G0 = v7 + int32(48)
										return
									}
								} else {
									if l0 != int32(2995) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
												F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2445), int32(_a_F_truncate_check_rel_3))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
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
										v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
										if v76 != 0 {
											v79 = int32(0)
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
											m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
											}
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v10
							F_errmsg(m, int32(_a_F_truncate_check_rel_4), v7+int32(16))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2422), int32(_a_F_truncate_check_rel_3))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
				F_errmsg(m, int32(_a_F_truncate_check_rel_5), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2428), int32(_a_F_truncate_check_rel_3))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 10, 12:
		v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[0])))
		if v54 != 0 {
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
			if v76 != 0 {
				v79 = int32(0)
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
				m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		} else {
			v56 = int32(1)
			if base.Ui32(l0) < base.Ui32(int32(_a_F_truncate_check_rel_0)) {
				v64 = v56
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if v59 == int32(99) {
					v64 = v56
				} else {
					v62 = F_isTempToastNamespace(m, v59)
					mBase = m.M
					v64 = v62
				}
			}
			if v64 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
				if v76 != 0 {
					v79 = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
					m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						m.G0 = v7 + int32(48)
						return
					}
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[2])))
				if v68 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
							F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2445), int32(_a_F_truncate_check_rel_3))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
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
					if l0 == int32(2613) {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						if l0 != int32(2995) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
									F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2445), int32(_a_F_truncate_check_rel_3))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
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
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							if v76 != 0 {
								v79 = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
								m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	}
}
