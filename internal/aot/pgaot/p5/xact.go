package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XactLogAbortRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	v11 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v11
	if l7&int32(2) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v22
	v25 = v22
	goto L3
L2:
	;
	v25 = v11
	goto L3
L3:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
	v30 = v25 | int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v30
	v32 = v30
	goto L6
L5:
	;
	v32 = v25
	goto L6
L6:
	;
	if l8 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = int32(64)
	goto L9
L8:
	;
	v35 = int32(32)
	goto L9
L9:
	;
	if int32(0) < l3 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
	v40 = v32 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v40
	v44 = v35 | int32(1)
	v45 = v40
	goto L12
L11:
	;
	v44 = v35
	v45 = v32
	goto L12
L12:
	;
	if int32(0) < l5 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l5
	v50 = v45 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v50
	v52 = v50
	goto L15
L14:
	;
	v52 = v45
	goto L15
L15:
	;
	if l8 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[3])))
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v72 = v52
	goto L16
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l8
	v57 = v52 | int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[0]))
	if v60 < int32(2) {
		v72 = v57
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v64 = v52 | int32(145)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v70
	v72 = v64
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v72 | int32(32)
	v79 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v79
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v82
	v85 = int32(1)
	goto L23
L22:
	;
	v85 = v72
	goto L23
L23:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int64(0)
L25:
	;
	F_XLogRegisterData(m, v13+int32(-8), int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v183 = int32(_a_F_XactLogAbortRecord_0)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[6])))
	v186 = v185 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[6])) = uint8(v186)
	goto L56
L28:
	;
	F_XLogRegisterData(m, v13+int32(-12), int32(4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v102&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_XLogRegisterData(m, v13+int32(-36), int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L33
	}
L31:
	;
	v111 = v102
	goto L32
L32:
	;
	if v111&int32(2) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v111 = v110
	goto L32
L34:
	;
	F_XLogRegisterData(m, v13+int32(-16), int32(4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L37
	}
L35:
	;
	v124 = v111
	goto L36
L36:
	;
	if v124&int32(4) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	F_XLogRegisterData(m, l2, l1<<(uint(int32(2))%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v124 = v123
	goto L36
L39:
	;
	F_XLogRegisterData(m, v13+int32(-20), int32(4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	v137 = v124
	goto L41
L41:
	;
	if v137&int32(256) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_XLogRegisterData(m, l4, l3*int32(12))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v137 = v136
	goto L41
L44:
	;
	F_XLogRegisterData(m, v13+int32(-24), int32(4))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L24
	} else {
		goto L47
	}
L45:
	;
	v150 = v137
	goto L46
L46:
	;
	if v150&int32(16) == int32(0) {
		v171 = v150
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F_XLogRegisterData(m, l6, l5<<(uint(int32(4))%32))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L48
	}
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v150 = v149
	goto L46
L49:
	;
	if v171&int32(32) == int32(0) {
		goto L27
	} else {
		goto L54
	}
L50:
	;
	F_XLogRegisterData(m, v13+int32(-28), int32(4))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v160&int32(128) == int32(0) {
		v171 = v160
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v165 = F_strlen(m, l9)
	mBase = m.M
	F_XLogRegisterData(m, l9, v165+int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v171 = v170
	goto L49
L54:
	;
	F_XLogRegisterData(m, v13+int32(-56), int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L24
	} else {
		goto L55
	}
L55:
	;
	goto L27
L56:
	;
	if v85 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v191 = v44 | int32(128)
	goto L59
L58:
	;
	v191 = v44
	goto L59
L59:
	;
	v192 = F_XLogInsert(m, int32(1), v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v15 - int32(-64)
	return v192
}
func F_xact_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
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
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int64
	_ = v711
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v758 int32
	_ = v758
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
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int64
	_ = v821
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int64
	_ = v965
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int64
	_ = v1050
	var v1051 int64
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1069 int32
	_ = v1069
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int64
	_ = v1098
	var v1099 int64
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1135 int64
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1137 int64
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int64
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1159 int32
	_ = v1159
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int64
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1221 int32
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1283 int64
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1368 int64
	_ = v1368
	var v1369 int64
	_ = v1369
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int64
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1409 int64
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int64
	_ = v1415
	var v1416 int64
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int64
	_ = v1432
	var v1433 int64
	_ = v1433
	var v1434 int64
	_ = v1434
	var v1436 int64
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int64
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1537 int32
	_ = v1537
	var v1538 int64
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1630 int64
	_ = v1630
	var v1632 int64
	_ = v1632
	var v1634 int64
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1826 int64
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int64
	_ = v1832
	var v1833 int64
	_ = v1833
	var v1834 int64
	_ = v1834
	var v1835 int64
	_ = v1835
	var v1838 int64
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int64
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1884 int32
	_ = v1884
	var v1890 int32
	_ = v1890
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int64
	_ = v1905
	var v1906 int64
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1923 int32
	_ = v1923
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int64
	_ = v1953
	var v1954 int64
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2028 int64
	_ = v2028
	var v2029 int64
	_ = v2029
	var v2030 int64
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int64
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int64
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int64
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int64
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2301 int32
	_ = v2301
	v3 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(304)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+48)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 <= v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v29 + int32(304)
	return
L2:
	;
	v40 = v34 & int32(112)
	switch int32(base.Ui32(v40)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		goto L5
	case 1, 3:
		goto L7
	default:
		goto L8
	case 4:
		goto L1
	case 5:
		goto L6
	case 6:
		goto L4
	}
L3:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v2195 {
		goto L528
	} else {
		goto L529
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L38
	} else {
		goto L525
	}
L5:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+64))
	v1626 = v29 + int32(16)
	base.MemoryFill(m, v1626, int32(0), int32(288))
	v1630 = *(*int64)(unsafe.Add(mBase, uint32(v1624)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1626))) = v1630
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v1624)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1626)+272)) = v1632
	v1634 = *(*int64)(unsafe.Add(mBase, uint32(v1624)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1626)+280)) = v1634
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+52)) = v1636
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+12)) = v1638
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+20)) = v1640
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+28)) = v1642
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+256)) = v1644
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+36)) = v1646
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+264)) = v1648
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+44)) = v1650
	v2301 = int32(72)
	v1653 = v29 + v2301
	v1655 = v1624 + v2301
	v1656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1624)+54)))
	if (v1655^v1653)&int32(3) != 0 {
		v1727 = v1655
		v1728 = v1656
		v1729 = v1653
		goto L413
	} else {
		goto L414
	}
L6:
	;
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+64))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+36))
	if v1447 != 0 {
		goto L364
	} else {
		goto L365
	}
L7:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+96))
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+48)))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+64))
	v1276 = v29 + int32(16)
	v1277 = int32(0)
	base.MemoryFill(m, v1276, v1277, int32(264))
	v1283 = *(*int64)(unsafe.Add(mBase, uint32(v1274)))
	*(*int64)(unsafe.Add(mBase, uint32(v1276))) = v1283
	if v1277 <= base.I32_extend8_s(v1272) {
		goto L318
	} else {
		goto L319
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+48)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
	v51 = v29 + int32(16)
	v52 = int32(0)
	base.MemoryFill(m, v51, v52, int32(288))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v58
	if v52 <= base.I32_extend8_s(v47) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v166 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	goto L9
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v63
	if v63&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v69
	v75 = v49 + int32(20)
	goto L14
L13:
	;
	v75 = v49 + int32(12)
	goto L14
L14:
	;
	if v63&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v80 = v75 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+24)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v78
	v86 = v80 + v78<<(uint(int32(2))%32)
	goto L17
L16:
	;
	v86 = v75
	goto L17
L17:
	;
	if v63&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v92 = v86 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v90
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v99 = v92 + v95*int32(12)
	goto L20
L19:
	;
	v99 = v86
	goto L20
L20:
	;
	if v63&int32(256) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = int32(4)
	v106 = v99 + v105
	*(*int32)(unsafe.Add(mBase, uint32(v51)+40)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v104
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v113 = v106 + v109<<(uint(v105)%32)
	goto L23
L22:
	;
	v113 = v99
	goto L23
L23:
	;
	if v63&int32(8) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v119 = int32(4)
	v120 = v113 + v119
	*(*int32)(unsafe.Add(mBase, uint32(v51)+48)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v51)+44)) = v118
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v127 = v120 + v123<<(uint(v119)%32)
	goto L26
L25:
	;
	v127 = v113
	goto L26
L26:
	;
	if v63&int32(16) == int32(0) {
		v151 = v63
		v152 = v127
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v151&int32(32) == int32(0) {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+52)) = v134
	v137 = v127 + int32(4)
	if v63&int32(128) == int32(0) {
		v151 = v63
		v152 = v137
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v145 = F_strlcpy(m, v29+int32(72), v137, int32(200))
	mBase = m.M
	v146 = F_strlen(m, v137)
	mBase = m.M
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v151 = v150
	v152 = v146 + v137 + int32(1)
	goto L27
L30:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+280)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v51)+272)) = v157
	goto L10
L31:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+36))
	v171 = v170
	goto L33
L32:
	;
	v171 = v166
	goto L33
L33:
	;
	if v40 != int32(48) {
		v187 = v3
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v29+int32(16)+v190<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(280))))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+96))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+56)))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v29)+288))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v207 = m.G0
	v209 = v207 - int32(160)
	m.G0 = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	switch v211 + int32(1) {
	case 0:
		goto L45
	case 1:
		goto L46
	default:
		v235 = v211
		goto L43
	}
L35:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v174 != int32(1) {
		v187 = v3
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v178 == int32(0) {
		v187 = int32(1)
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v183 = F_filter_prepare_cb_wrapper(m, l0, v171, v29+int32(72))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	v187 = v183 ^ int32(1)
	goto L34
L40:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1050 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1051 = *(*int64)(unsafe.Add(mBase, uint32(v1049)+16))
	goto L269
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L38
	} else {
		goto L264
	}
L42:
	;
	m.G0 = v209 + int32(160)
	goto L40
L43:
	;
	if v235 <= int32(1) {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v235 = v234
	goto L43
L45:
	;
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui64(v204) < base.Ui64(v229) {
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v203)+60))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v214))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v171)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v226 == int32(0) {
		goto L44
	} else {
		goto L51
	}
L48:
	;
	v226 = base.B2i32(base.Ui32(v171) < base.Ui32(v214))
	goto L47
L49:
	;
	goto L50
L50:
	;
	v226 = int32(base.Ui32(v171-v214) >> (uint(int32(31)) % 32))
	goto L47
L51:
	;
	goto L45
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v204 + int64(1)
	goto L42
L53:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui64(v238) <= base.Ui64(v204) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v244 = v3
	goto L55
L55:
	;
	if int32(0) < v205 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v204 + int64(1)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+36)))
	v244 = v243
	goto L55
L59:
	;
	v251 = v171
	v259 = v3
	v260 = v3
	goto L62
L60:
	;
	v423 = v171
	v432 = v3
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v171
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v449 = F_ReorderBufferXidHasCatalogChanges(m, v448, v171)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L38
	} else {
		goto L114
	}
L62:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v206+v259<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v281 = F_ReorderBufferXidHasCatalogChanges(m, v280, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L66
	}
L63:
	;
	v423 = v414
	v432 = v416
	goto L61
L64:
	;
	v419 = v259 + int32(1)
	if v419 != v205 {
		v251 = v414
		v259 = v419
		v260 = v416
		goto L62
	} else {
		goto L109
	}
L65:
	;
	if v244&int32(1) == int32(0) {
		v414 = v251
		v416 = v260
		goto L64
	} else {
		goto L94
	}
L66:
	;
	if v281 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v190&int32(8) == int32(0) {
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v302 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L38
	} else {
		goto L74
	}
L70:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v287 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v203)+84))
	v295 = F_bsearch(m, v209+int32(156), v292, v287, int32(4), int32(185))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L38
	} else {
		goto L72
	}
L72:
	;
	if v295 == int32(0) {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	if v302 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+132)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v209)+128)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_0), v209+int32(128))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L38
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v316 != v317 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(995), int32(_a_F_xact_decode_2))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v351 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v349 + v351
	*(*int32)(unsafe.Add(mBase, uint32(v350+v349<<(uint(int32(2))%32)))) = v278
	if int32(0) < v278-v251 {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v349 = v316
	v350 = v319
	goto L80
L82:
	;
	goto L83
L83:
	;
	v320 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v316<<(uint(v320)%32) | v320
	v327 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L38
	} else {
		goto L84
	}
L84:
	;
	if v327 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+112)) = v329
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(112))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L38
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v345 = F_repalloc(m, v341, v342<<(uint(int32(2))%32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L38
	} else {
		goto L90
	}
L88:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(838), int32(_a_F_xact_decode_4))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L38
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v345
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v349 = v348
	v350 = v345
	goto L80
L91:
	;
	v362 = v278
	goto L93
L92:
	;
	v362 = v251
	goto L93
L93:
	;
	v414 = v362
	v416 = v351
	goto L64
L94:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v368 != v369 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v401 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v402+v401<<(uint(int32(2))%32)))) = v278
	if int32(0) < v278-v251 {
		goto L106
	} else {
		goto L107
	}
L96:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v401 = v368
	v402 = v371
	goto L95
L97:
	;
	goto L98
L98:
	;
	v372 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v368<<(uint(v372)%32) | v372
	v379 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L38
	} else {
		goto L99
	}
L99:
	;
	if v379 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+144)) = v381
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(144))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L38
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v397 = F_repalloc(m, v393, v394<<(uint(int32(2))%32))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L38
	} else {
		goto L105
	}
L103:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(838), int32(_a_F_xact_decode_4))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L38
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v397
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v401 = v400
	v402 = v397
	goto L95
L106:
	;
	v413 = v278
	goto L108
L107:
	;
	v413 = v251
	goto L108
L108:
	;
	v414 = v413
	v416 = v260
	goto L64
L109:
	;
	goto L63
L110:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v668 <= int32(0) {
		goto L42
	} else {
		goto L192
	}
L111:
	;
	v660 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+72)) = uint8(v660)
	if v432 == v660 {
		goto L42
	} else {
		goto L191
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v626 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v627+v626<<(uint(int32(2))%32)))) = v171
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v636 != 0 {
		goto L178
	} else {
		goto L179
	}
L113:
	;
	if v432 != 0 {
		goto L139
	} else {
		goto L140
	}
L114:
	;
	if v449 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v190&int32(8) == int32(0) {
		goto L113
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v472 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L38
	} else {
		goto L122
	}
L118:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v457 == int32(0) {
		goto L113
	} else {
		goto L119
	}
L119:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v203)+84))
	v465 = F_bsearch(m, v209+int32(156), v462, v457, int32(4), int32(185))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L38
	} else {
		goto L120
	}
L120:
	;
	if v465 == int32(0) {
		goto L113
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	if v472 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+32)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_5), v209+int32(32))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L38
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v485 != v486 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1021), int32(_a_F_xact_decode_2))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L38
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v626 = v518
	v627 = v519
	v628 = int32(1)
	goto L112
L129:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v518 = v485
	v519 = v488
	goto L128
L130:
	;
	goto L131
L131:
	;
	v489 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v485<<(uint(v489)%32) | v489
	v496 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L38
	} else {
		goto L132
	}
L132:
	;
	if v496 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+16)) = v498
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(16))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L38
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v514 = F_repalloc(m, v510, v511<<(uint(int32(2))%32))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L38
	} else {
		goto L138
	}
L136:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(838), int32(_a_F_xact_decode_4))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L38
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v514
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v518 = v517
	v519 = v514
	goto L128
L139:
	;
	v524 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L38
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v244&int32(1) == int32(0) {
		goto L111
	} else {
		goto L159
	}
L142:
	;
	if v524 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+64)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_6), v209-int32(-64))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L38
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v537 != v538 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1030), int32(_a_F_xact_decode_2))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L38
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v626 = v570
	v627 = v571
	v628 = v432
	goto L112
L149:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v570 = v537
	v571 = v540
	goto L148
L150:
	;
	goto L151
L151:
	;
	v541 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v537<<(uint(v541)%32) | v541
	v548 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L38
	} else {
		goto L152
	}
L152:
	;
	if v548 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+48)) = v550
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(48))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L38
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v566 = F_repalloc(m, v562, v563<<(uint(int32(2))%32))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L38
	} else {
		goto L158
	}
L156:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(838), int32(_a_F_xact_decode_4))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L38
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v566
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v570 = v569
	v571 = v566
	goto L148
L159:
	;
	v578 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L38
	} else {
		goto L160
	}
L160:
	;
	if v578 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+96)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_7), v209+int32(96))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L38
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v591 != v592 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1036), int32(_a_F_xact_decode_2))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L38
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v626 = v624
	v627 = v625
	v628 = v432
	goto L112
L167:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v624 = v591
	v625 = v594
	goto L166
L168:
	;
	goto L169
L169:
	;
	v595 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v591<<(uint(v595)%32) | v595
	v602 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L38
	} else {
		goto L170
	}
L170:
	;
	if v602 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+80)) = v604
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(80))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L38
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v620 = F_repalloc(m, v616, v617<<(uint(int32(2))%32))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L38
	} else {
		goto L176
	}
L174:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(838), int32(_a_F_xact_decode_4))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L38
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v620
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v624 = v623
	v625 = v620
	goto L166
L177:
	;
	if v628 == int32(0) {
		goto L42
	} else {
		goto L190
	}
L178:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v636))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v423)) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	goto L180
L180:
	;
	v651 = int32(3)
	v653 = v423 + int32(1)
	if base.Ui32(v653) <= base.Ui32(v651) {
		goto L186
	} else {
		goto L187
	}
L181:
	;
	if v648 == int32(0) {
		goto L177
	} else {
		goto L185
	}
L182:
	;
	v648 = base.B2i32(base.Ui32(v636) <= base.Ui32(v423))
	goto L181
L183:
	;
	goto L184
L184:
	;
	v648 = base.B2i32(int32(0) <= v423-v636)
	goto L181
L185:
	;
	goto L180
L186:
	;
	v656 = v651
	goto L188
L187:
	;
	v656 = v653
	goto L188
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v656
	if v628 != 0 {
		goto L110
	} else {
		goto L189
	}
L189:
	;
	goto L42
L190:
	;
	goto L110
L191:
	;
	goto L110
L192:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	if v671 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v690 = F_MemoryContextAllocZero(m, v684, v685<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L38
	} else {
		goto L198
	}
L194:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+30)))
	if v674 == int32(1) {
		goto L41
	} else {
		goto L195
	}
L195:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v671)+44))
	v679 = v677 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v671)+44)) = v679
	if v679 != 0 {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	F_pfree(m, v671)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L38
	} else {
		goto L197
	}
L197:
	;
	goto L193
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v690))) = int32(5)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v690)+4)) = v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v698 = v690 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v690)+12)) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v690)+8)) = v696
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v690)+16)) = v701
	v704 = v701 << (uint(int32(2)) % 32)
	if v704 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	base.MemoryCopy(m, v698, v705, v704)
	goto L201
L200:
	;
	goto L201
L201:
	;
	F_pg_qsort(m, v698, v701, int32(4), int32(185))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L38
	} else {
		goto L202
	}
L202:
	;
	v711 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v690)+64)) = v711
	*(*int64)(unsafe.Add(mBase, uint32(v690)+44)) = v711
	v715 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v690)+32)) = v715
	*(*int64)(unsafe.Add(mBase, uint32(v690)+20)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v690)+27)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v203)+40)) = v690
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v723 = F_ReorderBufferXidHasBaseSnapshot(m, v722, v171)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L38
	} else {
		goto L203
	}
L203:
	;
	if v723 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v727)+44)) = v728 + int32(1)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	F_ReorderBufferSetBaseSnapshot(m, v732, v171, v204, v733)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L38
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+44)) = v738 + int32(1)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+8))
	if v743 == int32(0) {
		goto L42
	} else {
		goto L208
	}
L207:
	;
	goto L206
L208:
	;
	v747 = v742 + int32(4)
	if v743 == v747 {
		goto L42
	} else {
		goto L209
	}
L209:
	;
	v758 = v743
	goto L210
L210:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v781 = v758 - int32(184)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v783 = F_ReorderBufferXidHasBaseSnapshot(m, v779, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L38
	} else {
		goto L213
	}
L211:
	;
	goto L42
L212:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	if v1004 != v747 {
		v758 = v1004
		goto L210
	} else {
		goto L263
	}
L213:
	;
	if v783 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758-int32(188)))))
	if v789&int32(64) != 0 {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v794 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L38
	} else {
		goto L216
	}
L216:
	;
	if v794 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = base.I32_wrap_i64(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = base.I32_wrap_i64(int64(base.Ui64(v204) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v796
	F_errmsg_internal(m, int32(_a_F_xact_decode_8), v209)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L38
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v809)+44)) = v810 + int32(1)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+124))
	v819 = F_MemoryContextAlloc(m, v817, int32(64))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L38
	} else {
		goto L222
	}
L220:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(778), int32(_a_F_xact_decode_9))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L38
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v821 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v819)+16)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+8)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+56)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+48)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+40)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+32)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819)+24)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v819))) = v821
	*(*int32)(unsafe.Add(mBase, uint32(v819)+20)) = v815
	*(*int32)(unsafe.Add(mBase, uint32(v819)+8)) = int32(5)
	F_ReorderBufferQueueChange(m, v816, v814, v204, v819, int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L38
	} else {
		goto L223
	}
L223:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	if v843 == v171 {
		goto L212
	} else {
		goto L224
	}
L224:
	;
	v845 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v845
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v850 = m.G0
	v852 = v850 - int32(16)
	m.G0 = v852
	*(*int32)(unsafe.Add(mBase, uint32(v852)+12)) = v171
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v847)+32))
	if base.B2i32(v855 == v845)|base.B2i32(v171 != v855) == v845 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	m.G0 = v852 + int32(16)
	if v892 == int32(0) {
		goto L212
	} else {
		goto L236
	}
L226:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v887)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v209+int32(156)))) = v889
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v887)+172))
	v892 = v891
	goto L225
L227:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v847)+36))
	if v862 != 0 {
		v887 = v862
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v864 = int32(0)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	v871 = F_hash_search(m, v865, v852+int32(12), v864, v852+int32(11))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L38
	} else {
		goto L231
	}
L230:
	;
	v892 = int32(0)
	goto L225
L231:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+11)))
	if v873 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v852)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v847)+32)) = v876
	v892 = v864
	goto L225
L233:
	;
	goto L234
L234:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v852)+12))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+36)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(v847)+32)) = v880
	if v881 == int32(0) {
		v892 = v864
		goto L225
	} else {
		goto L235
	}
L235:
	;
	v887 = v881
	goto L226
L236:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v209)+156))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v904 = F_ReorderBufferTXNByXid(m, v901, v902, int32(0), v204)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L38
	} else {
		goto L237
	}
L237:
	;
	v906 = int32(_a_F_xact_decode_10)
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v901)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v909
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v904)+40))
	if v911 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v901)+124))
	v963 = F_MemoryContextAlloc(m, v961, int32(64))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L38
	} else {
		goto L257
	}
L239:
	;
	v912 = v911
	goto L241
L240:
	;
	v912 = v904
	goto L241
L241:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	if v913&int32(_a_F_xact_decode_11) != 0 {
		goto L238
	} else {
		goto L242
	}
L242:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v912)+180))
	v917 = v916 + v892
	if base.Ui32(int32(_a_F_xact_decode_12)) <= base.Ui32(v917) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912))) = v913 | int32(_a_F_xact_decode_11)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v912)+184))
	if v923 == int32(0) {
		goto L238
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if v916 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	F_pfree(m, v923)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L38
	} else {
		goto L247
	}
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v912)+180)) = int64(0)
	goto L238
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+180)) = v892
	v934 = v892 << (uint(int32(4)) % 32)
	v935 = F_palloc(m, v934)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L38
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v912)+184))
	v944 = F_repalloc(m, v941, v917<<(uint(int32(4))%32))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L38
	} else {
		goto L253
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+184)) = v935
	if v934 == int32(0) {
		goto L238
	} else {
		goto L252
	}
L252:
	;
	base.MemoryCopy(m, v935, v900, v934)
	goto L238
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+184)) = v944
	v948 = v892 << (uint(int32(4)) % 32)
	if v948 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v912)+180))
	base.MemoryCopy(m, v944+v949<<(uint(int32(4))%32), v900, v948)
	goto L256
L255:
	;
	goto L256
L256:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v912)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v912)+180)) = v954 + v892
	goto L238
L257:
	;
	v965 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v963)+16)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+8)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+56)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+48)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+40)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+32)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963)+24)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v963))) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v963)+20)) = v892
	v982 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v963)+8)) = v982
	v985 = v892 << (uint(v982) % 32)
	v986 = F_palloc(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L38
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v963)+24)) = v986
	if v985 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	base.MemoryCopy(m, v986, v900, v985)
	goto L261
L260:
	;
	goto L261
L261:
	;
	F_ReorderBufferQueueChange(m, v901, v902, v204, v963, int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L38
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v907
	goto L212
L263:
	;
	goto L211
L264:
	;
	F_errmsg_internal(m, int32(_a_F_xact_decode_13), int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L38
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(344), int32(_a_F_xact_decode_14))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L38
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	v1225 = m.G0
	v1227 = v1225 - int32(16)
	m.G0 = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v1227)+12)) = v171
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+32))
	v1231 = int32(0)
	if base.B2i32(v1230 == v1231)|base.B2i32(v1230 != v171) == v1231 {
		goto L306
	} else {
		goto L307
	}
L268:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1152 {
		goto L296
	} else {
		goto L297
	}
L269:
	;
	if base.Ui64(v1050) < base.Ui64(v1051) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	if v1048 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+88))
	if v1048 != v1054 {
		goto L268
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1056 != 0 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L273
L275:
	;
	v1057 = F_filter_by_origin_cb_wrapper(m, l0, v201)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L38
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1059 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	if v1057 != 0 {
		goto L268
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1062 {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	goto L282
L282:
	;
	v1150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v1150)
	goto L268
L283:
	;
	v1069 = int32(0)
	goto L286
L284:
	;
	goto L285
L285:
	;
	if v190&int32(32) != 0 {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1093+v1069<<(uint(int32(2))%32))))
	v1098 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ReorderBufferCommitChild(m, v1092, v171, v1097, v1098, v1099)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L38
	} else {
		goto L288
	}
L287:
	;
	goto L285
L288:
	;
	v1103 = v1069 + int32(1)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1103 < v1104 {
		v1069 = v1103
		goto L286
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	v1135 = v202
	goto L292
L291:
	;
	v1135 = int64(0)
	goto L292
L292:
	;
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1137 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v187 == int32(0) {
		goto L267
	} else {
		goto L293
	}
L293:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1142 = *(*int64)(unsafe.Add(mBase, uint32(v1141)+24))
	F_ReorderBufferFinishPrepared(m, v1138, v171, v1137, v1136, v1142, v198, v201, v1135, v29+int32(72), int32(1))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L38
	} else {
		goto L294
	}
L294:
	;
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L38
	} else {
		goto L295
	}
L295:
	;
	goto L1
L296:
	;
	v1159 = int32(0)
	goto L299
L297:
	;
	goto L298
L298:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferForget(m, v1221, v171, v1222)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L38
	} else {
		goto L303
	}
L299:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1183+v1159<<(uint(int32(2))%32))))
	v1188 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferForget(m, v1182, v1187, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L38
	} else {
		goto L301
	}
L300:
	;
	goto L298
L301:
	;
	v1192 = v1159 + int32(1)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1192 < v1193 {
		v1159 = v1192
		goto L299
	} else {
		goto L302
	}
L302:
	;
	goto L300
L303:
	;
	goto L1
L304:
	;
	m.G0 = v1227 + int32(16)
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L38
	} else {
		goto L316
	}
L305:
	;
	F_ReorderBufferReplay(m, v1260, v1138, v1137, v1136, v198, v201, v1135)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L38
	} else {
		goto L315
	}
L306:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+36))
	if v1237 != 0 {
		v1260 = v1237
		goto L305
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1138)))
	v1244 = F_hash_search(m, v1238, v1227+int32(12), int32(0), v1227+int32(11))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L38
	} else {
		goto L310
	}
L309:
	;
	goto L304
L310:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+11)))
	if v1246 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+32)) = v1249
	goto L304
L312:
	;
	goto L313
L313:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+12))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+36)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+32)) = v1253
	if v1254 == int32(0) {
		goto L304
	} else {
		goto L314
	}
L314:
	;
	v1260 = v1254
	goto L305
L315:
	;
	goto L304
L316:
	;
	goto L1
L317:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if v1377 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L318:
	;
	goto L317
L319:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+8)) = v1288
	if v1288&int32(1) != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+12)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+16)) = v1294
	v1300 = v1274 + int32(20)
	goto L322
L321:
	;
	v1300 = v1274 + int32(12)
	goto L322
L322:
	;
	if v1288&int32(2) != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	v1305 = v1300 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+24)) = v1305
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+20)) = v1303
	v1311 = v1305 + v1303<<(uint(int32(2))%32)
	goto L325
L324:
	;
	v1311 = v1300
	goto L325
L325:
	;
	if v1288&int32(4) != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1311)))
	v1317 = v1311 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+32)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+28)) = v1315
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1311)))
	v1324 = v1317 + v1320*int32(12)
	goto L328
L327:
	;
	v1324 = v1311
	goto L328
L328:
	;
	if v1288&int32(256) != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	v1330 = int32(4)
	v1331 = v1324 + v1330
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+40)) = v1331
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+36)) = v1329
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	v1338 = v1331 + v1334<<(uint(v1330)%32)
	goto L331
L330:
	;
	v1338 = v1324
	goto L331
L331:
	;
	if v1288&int32(16) == int32(0) {
		v1362 = v1288
		v1363 = v1338
		goto L332
	} else {
		goto L333
	}
L332:
	;
	if v1362&int32(32) == int32(0) {
		goto L318
	} else {
		goto L335
	}
L333:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+44)) = v1345
	v1348 = v1338 + int32(4)
	if v1288&int32(128) == int32(0) {
		v1362 = v1288
		v1363 = v1348
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1356 = F_strlcpy(m, v29+int32(64), v1348, int32(200))
	mBase = m.M
	v1357 = F_strlen(m, v1348)
	mBase = m.M
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+8))
	v1362 = v1361
	v1363 = v1357 + v1348 + int32(1)
	goto L332
L335:
	;
	v1368 = *(*int64)(unsafe.Add(mBase, uint32(v1363)))
	v1369 = *(*int64)(unsafe.Add(mBase, uint32(v1363)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1276)+256)) = v1369
	*(*int64)(unsafe.Add(mBase, uint32(v1276)+248)) = v1368
	goto L318
L336:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+36))
	v1382 = v1381
	goto L338
L337:
	;
	v1382 = v1377
	goto L338
L338:
	;
	v1383 = int32(0)
	if v40 != int32(64) {
		v1399 = v1383
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1400 = *(*int64)(unsafe.Add(mBase, uint32(v29)+264))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v1405 = v1403 & int32(32)
	v1409 = *(*int64)(unsafe.Add(mBase, uint32(v29+int32(16)+v1405<<(uint(int32(3))%32))))
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+96))
	v1412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1411)+56)))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1415 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1416 = *(*int64)(unsafe.Add(mBase, uint32(v1414)+16))
	goto L344
L340:
	;
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v1386 != int32(1) {
		v1399 = v1383
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1390 == int32(0) {
		v1399 = int32(1)
		goto L339
	} else {
		goto L342
	}
L342:
	;
	v1395 = F_filter_prepare_cb_wrapper(m, l0, v1382, v29-int32(-64))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L38
	} else {
		goto L343
	}
L343:
	;
	v1399 = v1395 ^ int32(1)
	goto L339
L344:
	;
	if base.Ui64(v1415) < base.Ui64(v1416) {
		goto L3
	} else {
		goto L345
	}
L345:
	;
	if v1413 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1418)+88))
	if v1413 != v1419 {
		goto L3
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1421 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L348
L350:
	;
	v1422 = F_filter_by_origin_cb_wrapper(m, l0, v1412)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L38
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1424 == int32(1) {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	if v1422 != 0 {
		goto L3
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1427 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v1427)
	goto L3
L356:
	;
	goto L357
L357:
	;
	if v1399 == int32(0) {
		goto L3
	} else {
		goto L358
	}
L358:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1432 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1433 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1434 = int64(0)
	if v1405 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1436 = v1400
	goto L361
L360:
	;
	v1436 = v1434
	goto L361
L361:
	;
	F_ReorderBufferFinishPrepared(m, v1431, v1382, v1432, v1433, v1434, v1409, v1412, v1436, v29-int32(-64), int32(0))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L38
	} else {
		goto L362
	}
L362:
	;
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L38
	} else {
		goto L363
	}
L363:
	;
	goto L1
L364:
	;
	if v1444&int32(1) == int32(0) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	goto L366
L366:
	;
	if v1444&int32(1) != 0 {
		goto L1
	} else {
		goto L391
	}
L367:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	v1454 = v1446 + int32(4)
	v1456 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1457 = F_ReorderBufferTXNByXid(m, v31, v1447, int32(0), v1456)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L38
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1538 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferXidSetCatalogChanges(m, v1537, v1447, v1538)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L38
	} else {
		goto L390
	}
L370:
	;
	v1459 = int32(_a_F_xact_decode_10)
	v1460 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v1462
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+40))
	if v1464 != 0 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1500 = F_MemoryContextAlloc(m, v1498, int32(64))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L38
	} else {
		goto L384
	}
L372:
	;
	v1465 = v1464
	goto L374
L373:
	;
	v1465 = v1457
	goto L374
L374:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+172))
	if v1466 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+172)) = v1452
	v1471 = v1452 << (uint(int32(4)) % 32)
	v1472 = F_palloc(m, v1471)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L38
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+176))
	v1482 = F_repalloc(m, v1478, (v1466+v1452)<<(uint(int32(4))%32))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L38
	} else {
		goto L380
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+176)) = v1472
	if v1471 == int32(0) {
		v1496 = v1471
		goto L371
	} else {
		goto L379
	}
L379:
	;
	base.MemoryCopy(m, v1472, v1454, v1471)
	v1496 = v1471
	goto L371
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+176)) = v1482
	v1486 = v1452 << (uint(int32(4)) % 32)
	if v1486 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+172))
	base.MemoryCopy(m, v1482+v1487<<(uint(int32(4))%32), v1454, v1486)
	goto L383
L382:
	;
	goto L383
L383:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+172)) = v1492 + v1452
	v1496 = v1486
	goto L371
L384:
	;
	v1502 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+16)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+8)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+56)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+48)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+40)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+32)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500)+24)) = v1502
	*(*int64)(unsafe.Add(mBase, uint32(v1500))) = v1502
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+20)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+8)) = int32(4)
	v1521 = F_palloc(m, v1496)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L38
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+24)) = v1521
	if v1496 != 0 {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	base.MemoryCopy(m, v1521, v1454, v1496)
	goto L388
L387:
	;
	goto L388
L388:
	;
	F_ReorderBufferQueueChange(m, v31, v1447, v1456, v1500, int32(0))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L38
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v1460
	goto L369
L390:
	;
	goto L1
L391:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	v1545 = int32(0)
	v1547 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[1]))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+24))
	v1550 = base.B2i32(v1548 != v1545)
	goto L392
L392:
	;
	if v1548 != v1545 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_xact_decode_15))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L38
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	if v1544 != 0 {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L38
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	v1558 = v1545
	goto L401
L399:
	;
	goto L400
L400:
	;
	if v1548 != v1545 {
		goto L405
	} else {
		goto L406
	}
L401:
	;
	F_LocalExecuteInvalidationMessage(m, v1446+int32(4)+v1558<<(uint(int32(4))%32))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L38
	} else {
		goto L403
	}
L402:
	;
	goto L400
L403:
	;
	v1590 = v1558 + int32(1)
	if v1590 != v1544 {
		v1558 = v1590
		goto L401
	} else {
		goto L404
	}
L404:
	;
	goto L402
L405:
	;
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L38
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	goto L1
L408:
	;
	goto L407
L409:
	;
	v1766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1624)+54)))
	v1767 = int32(7)
	v1771 = v1655 + (v1766+v1767)&int32(_a_F_xact_decode_16)
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+24)) = v1771
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+28))
	v1778 = int32(-8)
	v1780 = v1771 + (v1773<<(uint(int32(2))%32)+v1767)&v1778
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+32)) = v1780
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+32))
	v1783 = int32(12)
	v1789 = v1780 + (v1782*v1783+v1767)&v1778
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+260)) = v1789
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+36))
	v1798 = v1789 + (v1791*v1783+v1767)&v1778
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+40)) = v1798
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+40))
	v1801 = int32(4)
	v1803 = v1798 + v1800<<(uint(v1801)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+268)) = v1803
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+48)) = v1803 + v1805<<(uint(v1801)%32)
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v1811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v1811 == int32(1) {
		goto L435
	} else {
		goto L436
	}
L410:
	;
	F___memset(m, v1762, int32(0), v1761)
	mBase = m.M
	goto L409
L411:
	;
	v1761 = int32(0)
	v1762 = v1756
	goto L410
L412:
	;
	v1739 = v1734
	v1740 = v1735
	v1741 = v1736
	goto L430
L413:
	;
	if v1728 == int32(0) {
		v1756 = v1729
		goto L411
	} else {
		goto L429
	}
L414:
	;
	v1662 = int32(0)
	if base.B2i32(v1655&int32(3) == v1662)|base.B2i32(v1656 == v1662) != 0 {
		v1693 = v1655
		v1694 = v1656
		v1695 = v1653
		v1696 = base.B2i32(v1656 != v1662)
		goto L415
	} else {
		goto L416
	}
L415:
	;
	if v1696 == int32(0) {
		v1756 = v1695
		goto L411
	} else {
		goto L422
	}
L416:
	;
	v1672 = v1655
	v1673 = v1656
	v1674 = v1653
	goto L417
L417:
	;
	v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1676)
	if v1676 == int32(0) {
		v1761 = v1673
		v1762 = v1674
		goto L410
	} else {
		goto L419
	}
L418:
	;
	v1693 = v1687
	v1694 = v1683
	v1695 = v1681
	v1696 = v1685
	goto L415
L419:
	;
	v1680 = int32(1)
	v1681 = v1674 + v1680
	v1683 = v1673 - v1680
	v1684 = int32(0)
	v1685 = base.B2i32(v1683 != v1684)
	v1687 = v1672 + v1680
	if v1687&int32(3) == v1684 {
		v1693 = v1687
		v1694 = v1683
		v1695 = v1681
		v1696 = v1685
		goto L415
	} else {
		goto L420
	}
L420:
	;
	if v1683 != 0 {
		v1672 = v1687
		v1673 = v1683
		v1674 = v1681
		goto L417
	} else {
		goto L421
	}
L421:
	;
	goto L418
L422:
	;
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693))))
	if v1699 == int32(0) {
		v1761 = v1694
		v1762 = v1695
		goto L410
	} else {
		goto L423
	}
L423:
	;
	if base.Ui32(v1694) < base.Ui32(int32(4)) {
		v1727 = v1693
		v1728 = v1694
		v1729 = v1695
		goto L413
	} else {
		goto L424
	}
L424:
	;
	v1705 = v1693
	v1706 = v1694
	v1707 = v1695
	goto L425
L425:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1705)))
	v1713 = int32(-2139062144)
	if (int32(16843008)-v1710|v1710)&v1713 != v1713 {
		v1734 = v1705
		v1735 = v1706
		v1736 = v1707
		goto L412
	} else {
		goto L427
	}
L426:
	;
	v1727 = v1721
	v1728 = v1723
	v1729 = v1719
	goto L413
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1707))) = v1710
	v1718 = int32(4)
	v1719 = v1707 + v1718
	v1721 = v1705 + v1718
	v1723 = v1706 - v1718
	if base.Ui32(int32(3)) < base.Ui32(v1723) {
		v1705 = v1721
		v1706 = v1723
		v1707 = v1719
		goto L425
	} else {
		goto L428
	}
L428:
	;
	goto L426
L429:
	;
	v1734 = v1727
	v1735 = v1728
	v1736 = v1729
	goto L412
L430:
	;
	v1743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1739))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1741))) = uint8(v1743)
	if v1743 == int32(0) {
		v1761 = v1740
		v1762 = v1741
		goto L410
	} else {
		goto L432
	}
L431:
	;
	v1756 = v1748
	goto L411
L432:
	;
	v1747 = int32(1)
	v1748 = v1741 + v1747
	v1752 = v1740 - v1747
	if v1752 != 0 {
		v1739 = v1739 + v1747
		v1740 = v1752
		v1741 = v1748
		goto L430
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1832 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1833 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1834 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	v1835 = *(*int64)(unsafe.Add(mBase, uint32(v29)+296))
	if v1835 == int64(0) {
		goto L442
	} else {
		goto L443
	}
L435:
	;
	v1815 = v29 + int32(72)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1816 == int32(0) {
		v1829 = v1810
		goto L434
	} else {
		goto L438
	}
L436:
	;
	v1824 = v1810
	goto L437
L437:
	;
	v1826 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v31, v1824, v1826)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L38
	} else {
		goto L441
	}
L438:
	;
	v1819 = F_filter_prepare_cb_wrapper(m, l0, v1810, v1815)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L38
	} else {
		goto L439
	}
L439:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v1819 == int32(0) {
		v1829 = v1821
		goto L434
	} else {
		goto L440
	}
L440:
	;
	v1824 = v1821
	goto L437
L441:
	;
	goto L1
L442:
	;
	v1838 = v1834
	goto L444
L443:
	;
	v1838 = v1835
	goto L444
L444:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+96))
	v1841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840)+56)))
	v1842 = *(*int64)(unsafe.Add(mBase, uint32(v29)+288))
	v1843 = m.G0
	v1845 = v1843 - int32(16)
	m.G0 = v1845
	*(*int32)(unsafe.Add(mBase, uint32(v1845)+12)) = v1829
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+32))
	v1849 = int32(0)
	if base.B2i32(v1848 == v1849)|base.B2i32(v1829 != v1848) == v1849 {
		goto L447
	} else {
		goto L448
	}
L445:
	;
	m.G0 = v1845 + int32(16)
	if v1890 == int32(0) {
		goto L1
	} else {
		goto L456
	}
L446:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1877)+72)) = v1838
	*(*int64)(unsafe.Add(mBase, uint32(v1877)+32)) = v1833
	*(*int64)(unsafe.Add(mBase, uint32(v1877)+24)) = v1832
	*(*int64)(unsafe.Add(mBase, uint32(v1877)+64)) = v1842
	*(*uint16)(unsafe.Add(mBase, uint32(v1877)+56)) = uint16(v1841)
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1877)))
	*(*int32)(unsafe.Add(mBase, uint32(v1877))) = v1884 | int32(64)
	v1890 = int32(1)
	goto L445
L447:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+36))
	if v1855 != 0 {
		v1877 = v1855
		goto L446
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	v1862 = F_hash_search(m, v1856, v1845+int32(12), int32(0), v1845+int32(11))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L38
	} else {
		goto L451
	}
L450:
	;
	v1890 = v3
	goto L445
L451:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1845)+11)))
	if v1864 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1845)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1831)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1831)+32)) = v1867
	v1890 = v3
	goto L445
L453:
	;
	goto L454
L454:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1845)+12))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1862)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1831)+36)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v1831)+32)) = v1871
	if v1872 == int32(0) {
		v1890 = v3
		goto L445
	} else {
		goto L455
	}
L455:
	;
	v1877 = v1872
	goto L446
L456:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1830)))
	if v1897 <= int32(1) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ReorderBufferSkipPrepare(m, v1900, v1829)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L38
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1905 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1906 = *(*int64)(unsafe.Add(mBase, uint32(v1904)+16))
	goto L462
L460:
	;
	goto L1
L461:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ReorderBufferSkipPrepare(m, v2056, v1829)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L38
	} else {
		goto L499
	}
L462:
	;
	if base.Ui64(v1905) < base.Ui64(v1906) {
		goto L461
	} else {
		goto L463
	}
L463:
	;
	if v1903 != 0 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1908)+88))
	if v1903 != v1909 {
		goto L461
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1911 != 0 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	goto L466
L468:
	;
	v1912 = F_filter_by_origin_cb_wrapper(m, l0, v1841)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L38
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	v1914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1914 == int32(0) {
		goto L473
	} else {
		goto L474
	}
L471:
	;
	if v1912 != 0 {
		goto L461
	} else {
		goto L472
	}
L472:
	;
	goto L470
L473:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1917 {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	goto L475
L475:
	;
	v2054 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v2054)
	goto L461
L476:
	;
	v1923 = int32(0)
	goto L479
L477:
	;
	goto L478
L478:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1988 = m.G0
	v1990 = v1988 - int32(16)
	m.G0 = v1990
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+12)) = v1829
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+32))
	v1994 = int32(0)
	if base.B2i32(v1993 == v1994)|base.B2i32(v1993 != v1829) == v1994 {
		goto L485
	} else {
		goto L486
	}
L479:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1948+v1923<<(uint(int32(2))%32))))
	v1953 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1954 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ReorderBufferCommitChild(m, v1947, v1829, v1952, v1953, v1954)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L38
	} else {
		goto L481
	}
L480:
	;
	goto L478
L481:
	;
	v1958 = v1923 + int32(1)
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1958 < v1959 {
		v1923 = v1958
		goto L479
	} else {
		goto L482
	}
L482:
	;
	goto L480
L483:
	;
	m.G0 = v1990 + int32(16)
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L38
	} else {
		goto L498
	}
L484:
	;
	v2025 = F_pstrdup(m, v1815)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L38
	} else {
		goto L494
	}
L485:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+36))
	if v2000 != 0 {
		v2023 = v2000
		goto L484
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1987)))
	v2007 = F_hash_search(m, v2001, v1990+int32(12), int32(0), v1990+int32(11))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L38
	} else {
		goto L489
	}
L488:
	;
	goto L483
L489:
	;
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1990)+11)))
	if v2009 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1987)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1987)+32)) = v2012
	goto L483
L491:
	;
	goto L492
L492:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+12))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1987)+36)) = v2017
	*(*int32)(unsafe.Add(mBase, uint32(v1987)+32)) = v2016
	if v2017 == int32(0) {
		goto L483
	} else {
		goto L493
	}
L493:
	;
	v2023 = v2017
	goto L484
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2023)+12)) = v2025
	v2028 = *(*int64)(unsafe.Add(mBase, uint32(v2023)+24))
	v2029 = *(*int64)(unsafe.Add(mBase, uint32(v2023)+32))
	v2030 = *(*int64)(unsafe.Add(mBase, uint32(v2023)+72))
	v2031 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2023)+56)))
	v2032 = *(*int64)(unsafe.Add(mBase, uint32(v2023)+64))
	F_ReorderBufferReplay(m, v2023, v1987, v2028, v2029, v2030, v2031, v2032)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L38
	} else {
		goto L495
	}
L495:
	;
	v2035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2023)+1)))
	if v2035&int32(2) != 0 {
		goto L483
	} else {
		goto L496
	}
L496:
	;
	v2038 = *(*int64)(unsafe.Add(mBase, uint32(v2023)+24))
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+64))
	m.T0[v2039].(func(*base.Module, int32, int32, int64))(m, v1987, v2023, v2038)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L38
	} else {
		goto L497
	}
L497:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2023)))
	*(*int32)(unsafe.Add(mBase, uint32(v2023))) = v2042 | int32(512)
	goto L483
L498:
	;
	goto L1
L499:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2061 = m.G0
	v2063 = v2061 - int32(16)
	m.G0 = v2063
	*(*int32)(unsafe.Add(mBase, uint32(v2063)+12)) = v1829
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2060)+32))
	v2067 = int32(0)
	if base.B2i32(v2066 == v2067)|base.B2i32(v2066 != v1829) == v2067 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	m.G0 = v2063 + int32(16)
	goto L1
L501:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+80))
	if v2097 == int32(0) {
		goto L500
	} else {
		goto L511
	}
L502:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2060)+36))
	if v2073 != 0 {
		v2095 = v2073
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2060)))
	v2080 = F_hash_search(m, v2074, v2063+int32(12), int32(0), v2063+int32(11))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L38
	} else {
		goto L506
	}
L505:
	;
	goto L500
L506:
	;
	v2082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2063)+11)))
	if v2082 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2060)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2060)+32)) = v2085
	goto L500
L508:
	;
	goto L509
L509:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+12))
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2060)+36)) = v2090
	*(*int32)(unsafe.Add(mBase, uint32(v2060)+32)) = v2089
	if v2090 == int32(0) {
		goto L500
	} else {
		goto L510
	}
L510:
	;
	v2095 = v2090
	goto L501
L511:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+172))
	if v2100 == int32(0) {
		goto L500
	} else {
		goto L512
	}
L512:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+176))
	v2105 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[1]))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+24))
	v2108 = base.B2i32(v2106 != int32(0))
	goto L513
L513:
	;
	if v2106 != int32(0) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_xact_decode_15))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L38
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2115 = int32(0)
	goto L519
L517:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L38
	} else {
		goto L518
	}
L518:
	;
	goto L516
L519:
	;
	F_LocalExecuteInvalidationMessage(m, v2103+v2115<<(uint(int32(4))%32))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L38
	} else {
		goto L521
	}
L520:
	;
	if v2108 == int32(0) {
		goto L500
	} else {
		goto L523
	}
L521:
	;
	v2147 = v2115 + int32(1)
	if v2147 != v2100 {
		v2115 = v2147
		goto L519
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L38
	} else {
		goto L524
	}
L524:
	;
	goto L500
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	F_errmsg_internal(m, int32(_a_F_xact_decode_17), v29)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L38
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_xact_decode_18), int32(351), int32(_a_F_xact_decode_19))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L38
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	v2202 = int32(0)
	goto L531
L529:
	;
	goto L530
L530:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v2266)+40))
	F_ReorderBufferAbort(m, v2265, v1382, v2267, v1409)
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L38
	} else {
		goto L535
	}
L531:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2226+v2202<<(uint(int32(2))%32))))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2232 = *(*int64)(unsafe.Add(mBase, uint32(v2231)+40))
	F_ReorderBufferAbort(m, v2225, v2230, v2232, v1409)
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L38
	} else {
		goto L533
	}
L532:
	;
	goto L530
L533:
	;
	v2236 = v2202 + int32(1)
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v2236 < v2237 {
		v2202 = v2236
		goto L531
	} else {
		goto L534
	}
L534:
	;
	goto L532
L535:
	;
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L38
	} else {
		goto L536
	}
L536:
	;
	goto L1
}
