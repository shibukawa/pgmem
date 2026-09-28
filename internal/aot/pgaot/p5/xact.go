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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
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
	v80 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[4])))
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v78 = v52
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
	if v60 <= int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[1])))
	if v64&int32(1) == int32(0) {
		v78 = v57
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v70 = v52 | int32(145)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v76
	v78 = v70
	goto L16
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v78 | int32(32)
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v85
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v88
	v91 = int32(1)
	goto L26
L25:
	;
	v91 = v78
	goto L26
L26:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int64(0)
L28:
	;
	F_XLogRegisterData(m, v13+int32(-8), int32(8))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v91 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v189 = int32(_a_F_XactLogAbortRecord_0)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[7])))
	v192 = v191 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XactLogAbortRecord[7])) = uint8(v192)
	goto L59
L31:
	;
	F_XLogRegisterData(m, v13+int32(-12), int32(4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v108&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_XLogRegisterData(m, v13+int32(-36), int32(8))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	v117 = v108
	goto L35
L35:
	;
	if v117&int32(2) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v117 = v116
	goto L35
L37:
	;
	F_XLogRegisterData(m, v13+int32(-16), int32(4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L27
	} else {
		goto L40
	}
L38:
	;
	v130 = v117
	goto L39
L39:
	;
	if v130&int32(4) != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	F_XLogRegisterData(m, l2, l1<<(uint(int32(2))%32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v130 = v129
	goto L39
L42:
	;
	F_XLogRegisterData(m, v13+int32(-20), int32(4))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L27
	} else {
		goto L45
	}
L43:
	;
	v143 = v130
	goto L44
L44:
	;
	if v143&int32(256) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	F_XLogRegisterData(m, l4, l3*int32(12))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v143 = v142
	goto L44
L47:
	;
	F_XLogRegisterData(m, v13+int32(-24), int32(4))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L27
	} else {
		goto L50
	}
L48:
	;
	v156 = v143
	goto L49
L49:
	;
	if v156&int32(16) == int32(0) {
		v177 = v156
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_XLogRegisterData(m, l6, l5<<(uint(int32(4))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v156 = v155
	goto L49
L52:
	;
	if v177&int32(32) == int32(0) {
		goto L30
	} else {
		goto L57
	}
L53:
	;
	F_XLogRegisterData(m, v13+int32(-28), int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L27
	} else {
		goto L54
	}
L54:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v166&int32(128) == int32(0) {
		v177 = v166
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v171 = F_strlen(m, l9)
	mBase = m.M
	F_XLogRegisterData(m, l9, v171+int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v177 = v176
	goto L52
L57:
	;
	F_XLogRegisterData(m, v13+int32(-56), int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L27
	} else {
		goto L58
	}
L58:
	;
	goto L30
L59:
	;
	if v91 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v197 = v44 | int32(128)
	goto L62
L61:
	;
	v197 = v44
	goto L62
L62:
	;
	v198 = F_XLogInsert(m, int32(1), v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L27
	} else {
		goto L63
	}
L63:
	;
	m.G0 = v15 - int32(-64)
	return v198
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
	var v216 int32
	_ = v216
	var v227 int64
	_ = v227
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
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
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int64
	_ = v704
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
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
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v751 int32
	_ = v751
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int64
	_ = v814
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
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
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int64
	_ = v958
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v998 int32
	_ = v998
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int64
	_ = v1044
	var v1045 int64
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int64
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1129 int64
	_ = v1129
	var v1130 int64
	_ = v1130
	var v1131 int64
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int64
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1215 int32
	_ = v1215
	var v1216 int64
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1277 int64
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1362 int64
	_ = v1362
	var v1363 int64
	_ = v1363
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int64
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1403 int64
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int64
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int64
	_ = v1426
	var v1427 int64
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1430 int64
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int64
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int64
	_ = v1496
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1532 int32
	_ = v1532
	var v1533 int64
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1618 int32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1633 int64
	_ = v1633
	var v1635 int64
	_ = v1635
	var v1637 int64
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int64
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int64
	_ = v1835
	var v1836 int64
	_ = v1836
	var v1837 int64
	_ = v1837
	var v1838 int64
	_ = v1838
	var v1841 int64
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int64
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1887 int32
	_ = v1887
	var v1893 int32
	_ = v1893
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int64
	_ = v1908
	var v1909 int64
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1956 int64
	_ = v1956
	var v1957 int64
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int64
	_ = v2031
	var v2032 int64
	_ = v2032
	var v2033 int64
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int64
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2044 int64
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int64
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int64
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2315 int32
	_ = v2315
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
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v2209 {
		goto L528
	} else {
		goto L529
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L38
	} else {
		goto L525
	}
L5:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+64))
	v1629 = v29 + int32(16)
	base.MemoryFill(m, v1629, int32(0), int32(288))
	v1633 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1629))) = v1633
	v1635 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1629)+272)) = v1635
	v1637 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1629)+280)) = v1637
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+52)) = v1639
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+12)) = v1641
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+20)) = v1643
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+28)) = v1645
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+256)) = v1647
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+36)) = v1649
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+264)) = v1651
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+44)) = v1653
	v2315 = int32(72)
	v1656 = v29 + v2315
	v1658 = v1627 + v2315
	v1659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627)+54)))
	if (v1658^v1656)&int32(3) != 0 {
		v1730 = v1658
		v1731 = v1659
		v1732 = v1656
		goto L412
	} else {
		goto L413
	}
L6:
	;
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+64))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+36))
	if v1441 != 0 {
		goto L363
	} else {
		goto L364
	}
L7:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+96))
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1265)+48)))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+64))
	v1270 = v29 + int32(16)
	v1271 = int32(0)
	base.MemoryFill(m, v1270, v1271, int32(264))
	v1277 = *(*int64)(unsafe.Add(mBase, uint32(v1268)))
	*(*int64)(unsafe.Add(mBase, uint32(v1270))) = v1277
	if v1271 <= base.I32_extend8_s(v1266) {
		goto L317
	} else {
		goto L318
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
		goto L46
	case 1:
		goto L47
	default:
		goto L45
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
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1044 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1045 = *(*int64)(unsafe.Add(mBase, uint32(v1043)+16))
	goto L268
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L38
	} else {
		goto L263
	}
L42:
	;
	m.G0 = v209 + int32(160)
	goto L40
L43:
	;
	if v205 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui64(v236) <= base.Ui64(v204) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	if int32(1) < v211 {
		v243 = int32(0)
		goto L43
	} else {
		goto L54
	}
L46:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui64(v204) < base.Ui64(v227) {
		goto L42
	} else {
		goto L53
	}
L47:
	;
	v214 = int32(3)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v203)+60))
	if base.B2i32(base.Ui32(v171) < base.Ui32(v214))|base.B2i32(base.Ui32(v216) < base.Ui32(v214)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v171-v216 < int32(0) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v216) <= base.Ui32(v171) {
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L44
L52:
	;
	goto L46
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v204 + int64(1)
	goto L42
L54:
	;
	goto L44
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v204 + int64(1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+36)))
	v243 = v241
	goto L43
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v171
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v446 = F_ReorderBufferXidHasCatalogChanges(m, v445, v171)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L38
	} else {
		goto L115
	}
L59:
	;
	v420 = v171
	v429 = v3
	goto L58
L60:
	;
	goto L61
L61:
	;
	v250 = v171
	v257 = v3
	v259 = v3
	goto L62
L62:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v206+v257<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v280 = F_ReorderBufferXidHasCatalogChanges(m, v279, v277)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L38
	} else {
		goto L66
	}
L63:
	;
	v420 = v411
	v429 = v413
	goto L58
L64:
	;
	v416 = v257 + int32(1)
	if v416 != v205 {
		v250 = v411
		v257 = v416
		v259 = v413
		goto L62
	} else {
		goto L109
	}
L65:
	;
	if v243&int32(1) == int32(0) {
		v411 = v250
		v413 = v259
		goto L64
	} else {
		goto L94
	}
L66:
	;
	if v280 == int32(0) {
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
	v301 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L38
	} else {
		goto L74
	}
L70:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v286 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v203)+84))
	v294 = F_bsearch(m, v209+int32(156), v291, v286, int32(4), int32(187))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L38
	} else {
		goto L72
	}
L72:
	;
	if v294 == int32(0) {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	if v301 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+132)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v209)+128)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_0), v209+int32(128))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L38
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v315 != v316 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(999), int32(_a_F_xact_decode_2))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v349 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v347 + v349
	*(*int32)(unsafe.Add(mBase, uint32(v348+v347<<(uint(int32(2))%32)))) = v277
	if int32(0) < v277-v250 {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v347 = v315
	v348 = v318
	goto L80
L82:
	;
	goto L83
L83:
	;
	v319 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v315<<(uint(v319)%32) | v319
	v326 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L38
	} else {
		goto L84
	}
L84:
	;
	if v326 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+112)) = v328
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(112))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L38
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v343 = F_repalloc_mul(m, v340, int32(4), v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L38
	} else {
		goto L90
	}
L88:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(841), int32(_a_F_xact_decode_4))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L38
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v343
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v347 = v346
	v348 = v343
	goto L80
L91:
	;
	v360 = v277
	goto L93
L92:
	;
	v360 = v250
	goto L93
L93:
	;
	v411 = v360
	v413 = v349
	goto L64
L94:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v366 != v367 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v398 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v399+v398<<(uint(int32(2))%32)))) = v277
	if int32(0) < v277-v250 {
		goto L106
	} else {
		goto L107
	}
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v398 = v366
	v399 = v369
	goto L95
L97:
	;
	goto L98
L98:
	;
	v370 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v366<<(uint(v370)%32) | v370
	v377 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L38
	} else {
		goto L99
	}
L99:
	;
	if v377 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+144)) = v379
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(144))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L38
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v394 = F_repalloc_mul(m, v391, int32(4), v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L38
	} else {
		goto L105
	}
L103:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(841), int32(_a_F_xact_decode_4))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L38
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v394
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v398 = v397
	v399 = v394
	goto L95
L106:
	;
	v410 = v277
	goto L108
L107:
	;
	v410 = v250
	goto L108
L108:
	;
	v411 = v410
	v413 = v259
	goto L64
L109:
	;
	goto L63
L110:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v661 <= int32(0) {
		goto L42
	} else {
		goto L191
	}
L111:
	;
	if v622 == int32(0) {
		goto L42
	} else {
		goto L190
	}
L112:
	;
	v653 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+72)) = uint8(v653)
	if v429 != 0 {
		goto L110
	} else {
		goto L189
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+64)) = v620 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v621+v620<<(uint(int32(2))%32)))) = v171
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v630 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L114:
	;
	if v429 != 0 {
		goto L140
	} else {
		goto L141
	}
L115:
	;
	if v446 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v190&int32(8) == int32(0) {
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v469 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L38
	} else {
		goto L123
	}
L119:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v454 == int32(0) {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v203)+84))
	v462 = F_bsearch(m, v209+int32(156), v459, v454, int32(4), int32(187))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L38
	} else {
		goto L121
	}
L121:
	;
	if v462 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	if v469 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+32)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_5), v209+int32(32))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L38
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v482 != v483 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1025), int32(_a_F_xact_decode_2))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L38
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v620 = v514
	v621 = v515
	v622 = int32(1)
	goto L113
L130:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v514 = v482
	v515 = v485
	goto L129
L131:
	;
	goto L132
L132:
	;
	v486 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v482<<(uint(v486)%32) | v486
	v493 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L38
	} else {
		goto L133
	}
L133:
	;
	if v493 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+16)) = v495
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(16))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L38
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v510 = F_repalloc_mul(m, v507, int32(4), v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L38
	} else {
		goto L139
	}
L137:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(841), int32(_a_F_xact_decode_4))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L38
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v510
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v514 = v513
	v515 = v510
	goto L129
L140:
	;
	v520 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L38
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v243&int32(1) == int32(0) {
		goto L112
	} else {
		goto L160
	}
L143:
	;
	if v520 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+64)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_6), v209-int32(-64))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L38
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v533 != v534 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1034), int32(_a_F_xact_decode_2))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L38
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v620 = v565
	v621 = v566
	v622 = v429
	goto L113
L150:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v565 = v533
	v566 = v536
	goto L149
L151:
	;
	goto L152
L152:
	;
	v537 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v533<<(uint(v537)%32) | v537
	v544 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L38
	} else {
		goto L153
	}
L153:
	;
	if v544 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+48)) = v546
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(48))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L38
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v561 = F_repalloc_mul(m, v558, int32(4), v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L38
	} else {
		goto L159
	}
L157:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(841), int32(_a_F_xact_decode_4))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L38
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v561
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v565 = v564
	v566 = v561
	goto L149
L160:
	;
	v573 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L38
	} else {
		goto L161
	}
L161:
	;
	if v573 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+96)) = v171
	F_errmsg_internal(m, int32(_a_F_xact_decode_7), v209+int32(96))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L38
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	if v586 != v587 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(1040), int32(_a_F_xact_decode_2))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L38
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v620 = v618
	v621 = v619
	v622 = v429
	goto L113
L168:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v618 = v586
	v619 = v589
	goto L167
L169:
	;
	goto L170
L170:
	;
	v590 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v586<<(uint(v590)%32) | v590
	v597 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L38
	} else {
		goto L171
	}
L171:
	;
	if v597 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+80)) = v599
	F_errmsg_internal(m, int32(_a_F_xact_decode_3), v209+int32(80))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L38
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v614 = F_repalloc_mul(m, v611, int32(4), v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L38
	} else {
		goto L177
	}
L175:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(841), int32(_a_F_xact_decode_4))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L38
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+76)) = v614
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v618 = v617
	v619 = v614
	goto L167
L178:
	;
	v644 = int32(3)
	v646 = v420 + int32(1)
	if base.Ui32(v646) <= base.Ui32(v644) {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v633 = int32(3)
	if base.B2i32(base.Ui32(v420) < base.Ui32(v633))|base.B2i32(base.Ui32(v630) < base.Ui32(v633)) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(0) <= v420-v630 {
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	if base.Ui32(v420) < base.Ui32(v630) {
		goto L111
	} else {
		goto L184
	}
L183:
	;
	goto L111
L184:
	;
	goto L178
L185:
	;
	v649 = v644
	goto L187
L186:
	;
	v649 = v646
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v649
	if v622 == int32(0) {
		goto L42
	} else {
		goto L188
	}
L188:
	;
	goto L110
L189:
	;
	goto L42
L190:
	;
	goto L110
L191:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	if v664 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	v683 = F_MemoryContextAllocZero(m, v677, v678<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L38
	} else {
		goto L197
	}
L193:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+30)))
	if v667 == int32(1) {
		goto L41
	} else {
		goto L194
	}
L194:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v664)+44))
	v672 = v670 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v664)+44)) = v672
	if v672 != 0 {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	F_pfree(m, v664)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L38
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = int32(5)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v683)+4)) = v687
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v691 = v683 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v683)+12)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v683)+8)) = v689
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v683)+16)) = v694
	v697 = v694 << (uint(int32(2)) % 32)
	if v697 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	base.MemoryCopy(m, v691, v698, v697)
	goto L200
L199:
	;
	goto L200
L200:
	;
	F_pg_qsort(m, v691, v694, int32(4), int32(187))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L38
	} else {
		goto L201
	}
L201:
	;
	v704 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v683)+64)) = v704
	*(*int64)(unsafe.Add(mBase, uint32(v683)+44)) = v704
	v708 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v683)+32)) = v708
	*(*int64)(unsafe.Add(mBase, uint32(v683)+20)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v683)+27)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v203)+40)) = v683
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v716 = F_ReorderBufferXidHasBaseSnapshot(m, v715, v171)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L38
	} else {
		goto L202
	}
L202:
	;
	if v716 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v720)+44)) = v721 + int32(1)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	F_ReorderBufferSetBaseSnapshot(m, v725, v171, v204, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L38
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+44)) = v731 + int32(1)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+8))
	if v736 == int32(0) {
		goto L42
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v740 = v735 + int32(4)
	if v736 == v740 {
		goto L42
	} else {
		goto L208
	}
L208:
	;
	v751 = v736
	goto L209
L209:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v774 = v751 - int32(184)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v776 = F_ReorderBufferXidHasBaseSnapshot(m, v772, v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L38
	} else {
		goto L212
	}
L210:
	;
	goto L42
L211:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v751)+4))
	if v998 != v740 {
		v751 = v998
		goto L209
	} else {
		goto L262
	}
L212:
	;
	if v776 == int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751-int32(188)))))
	if v782&int32(64) != 0 {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	v787 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L38
	} else {
		goto L215
	}
L215:
	;
	if v787 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = base.I32_wrap_i64(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = base.I32_wrap_i64(int64(base.Ui64(v204) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v789
	F_errmsg_internal(m, int32(_a_F_xact_decode_8), v209)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L38
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+44)) = v803 + int32(1)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+124))
	v812 = F_MemoryContextAlloc(m, v810, int32(64))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L38
	} else {
		goto L221
	}
L219:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(781), int32(_a_F_xact_decode_9))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L38
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v814 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v812)+16)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+8)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+56)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+48)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+40)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+32)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812)+24)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v812))) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v812)+20)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v812)+8)) = int32(5)
	F_ReorderBufferQueueChange(m, v809, v807, v204, v812, int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L38
	} else {
		goto L222
	}
L222:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	if v836 == v171 {
		goto L211
	} else {
		goto L223
	}
L223:
	;
	v838 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+156)) = v838
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v843 = m.G0
	v845 = v843 - int32(16)
	m.G0 = v845
	*(*int32)(unsafe.Add(mBase, uint32(v845)+12)) = v171
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v840)+32))
	if base.B2i32(v848 == v838)|base.B2i32(v171 != v848) == v838 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	m.G0 = v845 + int32(16)
	if v885 == int32(0) {
		goto L211
	} else {
		goto L235
	}
L225:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v880)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v209+int32(156)))) = v882
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v880)+172))
	v885 = v884
	goto L224
L226:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v840)+36))
	if v855 != 0 {
		v880 = v855
		goto L225
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v857 = int32(0)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	v864 = F_hash_search(m, v858, v845+int32(12), v857, v845+int32(11))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L38
	} else {
		goto L230
	}
L229:
	;
	v885 = int32(0)
	goto L224
L230:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845)+11)))
	if v866 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v845)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v840)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v840)+32)) = v869
	v885 = v857
	goto L224
L232:
	;
	goto L233
L233:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v845)+12))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v840)+36)) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v840)+32)) = v873
	if v874 == int32(0) {
		v885 = v857
		goto L224
	} else {
		goto L234
	}
L234:
	;
	v880 = v874
	goto L225
L235:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v209)+156))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v203)+56))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v897 = F_ReorderBufferTXNByXid(m, v894, v895, int32(0), v204)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L38
	} else {
		goto L236
	}
L236:
	;
	v899 = int32(_a_F_xact_decode_10)
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v894)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v902
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v897)+40))
	if v904 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v894)+124))
	v956 = F_MemoryContextAlloc(m, v954, int32(64))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L38
	} else {
		goto L256
	}
L238:
	;
	v905 = v904
	goto L240
L239:
	;
	v905 = v897
	goto L240
L240:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	if v906&int32(_a_F_xact_decode_11) != 0 {
		goto L237
	} else {
		goto L241
	}
L241:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v905)+180))
	v910 = v909 + v885
	if base.Ui32(int32(_a_F_xact_decode_12)) <= base.Ui32(v910) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v905))) = v906 | int32(_a_F_xact_decode_11)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v905)+184))
	if v916 == int32(0) {
		goto L237
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if v909 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	F_pfree(m, v916)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L38
	} else {
		goto L246
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v905)+180)) = int64(0)
	goto L237
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v905)+180)) = v885
	v927 = F_palloc_mul(m, int32(16), v885)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L38
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v905)+184))
	v937 = F_repalloc_mul(m, v935, int32(16), v910)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L38
	} else {
		goto L252
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v905)+184)) = v927
	v931 = v885 << (uint(int32(4)) % 32)
	if v931 == int32(0) {
		goto L237
	} else {
		goto L251
	}
L251:
	;
	base.MemoryCopy(m, v927, v893, v931)
	goto L237
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v905)+184)) = v937
	v941 = v885 << (uint(int32(4)) % 32)
	if v941 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v905)+180))
	base.MemoryCopy(m, v937+v942<<(uint(int32(4))%32), v893, v941)
	goto L255
L254:
	;
	goto L255
L255:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v905)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v905)+180)) = v947 + v885
	goto L237
L256:
	;
	v958 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v956)+16)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+8)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+56)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+48)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+40)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+32)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956)+24)) = v958
	*(*int64)(unsafe.Add(mBase, uint32(v956))) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v956)+20)) = v885
	*(*int32)(unsafe.Add(mBase, uint32(v956)+8)) = int32(4)
	v978 = F_palloc_mul(m, int32(16), v885)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L38
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v956)+24)) = v978
	v982 = v885 << (uint(int32(4)) % 32)
	if v982 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	base.MemoryCopy(m, v978, v893, v982)
	goto L260
L259:
	;
	goto L260
L260:
	;
	F_ReorderBufferQueueChange(m, v894, v895, v204, v956, int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L38
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v900
	goto L211
L262:
	;
	goto L210
L263:
	;
	F_errmsg_internal(m, int32(_a_F_xact_decode_13), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L38
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_xact_decode_1), int32(348), int32(_a_F_xact_decode_14))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L38
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
	v1219 = m.G0
	v1221 = v1219 - int32(16)
	m.G0 = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+12)) = v171
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+32))
	v1225 = int32(0)
	if base.B2i32(v1224 == v1225)|base.B2i32(v1224 != v171) == v1225 {
		goto L305
	} else {
		goto L306
	}
L267:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1146 {
		goto L295
	} else {
		goto L296
	}
L268:
	;
	if base.Ui64(v1044) < base.Ui64(v1045) {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	if v1042 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+88))
	if v1042 != v1048 {
		goto L267
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1050 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L272
L274:
	;
	v1051 = F_filter_by_origin_cb_wrapper(m, l0, v201)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L38
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1053 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	if v1051 != 0 {
		goto L267
	} else {
		goto L278
	}
L278:
	;
	goto L276
L279:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1056 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	v1144 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v1144)
	goto L267
L282:
	;
	v1063 = int32(0)
	goto L285
L283:
	;
	goto L284
L284:
	;
	if v190&int32(32) != 0 {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1087+v1063<<(uint(int32(2))%32))))
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1093 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ReorderBufferCommitChild(m, v1086, v171, v1091, v1092, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L38
	} else {
		goto L287
	}
L286:
	;
	goto L284
L287:
	;
	v1097 = v1063 + int32(1)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1097 < v1098 {
		v1063 = v1097
		goto L285
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	v1129 = v202
	goto L291
L290:
	;
	v1129 = int64(0)
	goto L291
L291:
	;
	v1130 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1131 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v187 == int32(0) {
		goto L266
	} else {
		goto L292
	}
L292:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1135)+24))
	F_ReorderBufferFinishPrepared(m, v1132, v171, v1131, v1130, v1136, v198, v201, v1129, v29+int32(72), int32(1))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L38
	} else {
		goto L293
	}
L293:
	;
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L38
	} else {
		goto L294
	}
L294:
	;
	goto L1
L295:
	;
	v1153 = int32(0)
	goto L298
L296:
	;
	goto L297
L297:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1216 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferForget(m, v1215, v171, v1216)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L38
	} else {
		goto L302
	}
L298:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1177+v1153<<(uint(int32(2))%32))))
	v1182 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferForget(m, v1176, v1181, v1182)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L38
	} else {
		goto L300
	}
L299:
	;
	goto L297
L300:
	;
	v1186 = v1153 + int32(1)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1186 < v1187 {
		v1153 = v1186
		goto L298
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	goto L1
L303:
	;
	m.G0 = v1221 + int32(16)
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L38
	} else {
		goto L315
	}
L304:
	;
	F_ReorderBufferReplay(m, v1254, v1132, v1131, v1130, v198, v201, v1129)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L38
	} else {
		goto L314
	}
L305:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+36))
	if v1231 != 0 {
		v1254 = v1231
		goto L304
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	v1238 = F_hash_search(m, v1232, v1221+int32(12), int32(0), v1221+int32(11))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L38
	} else {
		goto L309
	}
L308:
	;
	goto L303
L309:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221)+11)))
	if v1240 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+32)) = v1243
	goto L303
L311:
	;
	goto L312
L312:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+12))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+36)) = v1248
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+32)) = v1247
	if v1248 == int32(0) {
		goto L303
	} else {
		goto L313
	}
L313:
	;
	v1254 = v1248
	goto L304
L314:
	;
	goto L303
L315:
	;
	goto L1
L316:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if v1371 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L317:
	;
	goto L316
L318:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+8)) = v1282
	if v1282&int32(1) != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+12)) = v1286
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+16)) = v1288
	v1294 = v1268 + int32(20)
	goto L321
L320:
	;
	v1294 = v1268 + int32(12)
	goto L321
L321:
	;
	if v1282&int32(2) != 0 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	v1299 = v1294 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+24)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+20)) = v1297
	v1305 = v1299 + v1297<<(uint(int32(2))%32)
	goto L324
L323:
	;
	v1305 = v1294
	goto L324
L324:
	;
	if v1282&int32(4) != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v1311 = v1305 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+32)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+28)) = v1309
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v1318 = v1311 + v1314*int32(12)
	goto L327
L326:
	;
	v1318 = v1305
	goto L327
L327:
	;
	if v1282&int32(256) != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1318)))
	v1324 = int32(4)
	v1325 = v1318 + v1324
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+40)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+36)) = v1323
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1318)))
	v1332 = v1325 + v1328<<(uint(v1324)%32)
	goto L330
L329:
	;
	v1332 = v1318
	goto L330
L330:
	;
	if v1282&int32(16) == int32(0) {
		v1356 = v1282
		v1357 = v1332
		goto L331
	} else {
		goto L332
	}
L331:
	;
	if v1356&int32(32) == int32(0) {
		goto L317
	} else {
		goto L334
	}
L332:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+44)) = v1339
	v1342 = v1332 + int32(4)
	if v1282&int32(128) == int32(0) {
		v1356 = v1282
		v1357 = v1342
		goto L331
	} else {
		goto L333
	}
L333:
	;
	v1350 = F_strlcpy(m, v29+int32(64), v1342, int32(200))
	mBase = m.M
	v1351 = F_strlen(m, v1342)
	mBase = m.M
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+8))
	v1356 = v1355
	v1357 = v1351 + v1342 + int32(1)
	goto L331
L334:
	;
	v1362 = *(*int64)(unsafe.Add(mBase, uint32(v1357)))
	v1363 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1270)+256)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v1270)+248)) = v1362
	goto L317
L335:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+36))
	v1376 = v1375
	goto L337
L336:
	;
	v1376 = v1371
	goto L337
L337:
	;
	v1377 = int32(0)
	if v40 != int32(64) {
		v1393 = v1377
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1394 = *(*int64)(unsafe.Add(mBase, uint32(v29)+264))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v1399 = v1397 & int32(32)
	v1403 = *(*int64)(unsafe.Add(mBase, uint32(v29+int32(16)+v1399<<(uint(int32(3))%32))))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+96))
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1405)+56)))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1409 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(v1408)+16))
	goto L343
L339:
	;
	v1380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v1380 != int32(1) {
		v1393 = v1377
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1384 == int32(0) {
		v1393 = int32(1)
		goto L338
	} else {
		goto L341
	}
L341:
	;
	v1389 = F_filter_prepare_cb_wrapper(m, l0, v1376, v29-int32(-64))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L38
	} else {
		goto L342
	}
L342:
	;
	v1393 = v1389 ^ int32(1)
	goto L338
L343:
	;
	if base.Ui64(v1409) < base.Ui64(v1410) {
		goto L3
	} else {
		goto L344
	}
L344:
	;
	if v1407 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+88))
	if v1407 != v1413 {
		goto L3
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1415 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L347
L349:
	;
	v1416 = F_filter_by_origin_cb_wrapper(m, l0, v1406)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L38
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1418 == int32(1) {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	if v1416 != 0 {
		goto L3
	} else {
		goto L353
	}
L353:
	;
	goto L351
L354:
	;
	v1421 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v1421)
	goto L3
L355:
	;
	goto L356
L356:
	;
	if v1393 == int32(0) {
		goto L3
	} else {
		goto L357
	}
L357:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1426 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1427 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1428 = int64(0)
	if v1399 != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1430 = v1394
	goto L360
L359:
	;
	v1430 = v1428
	goto L360
L360:
	;
	F_ReorderBufferFinishPrepared(m, v1425, v1376, v1426, v1427, v1428, v1403, v1406, v1430, v29-int32(-64), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L38
	} else {
		goto L361
	}
L361:
	;
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L38
	} else {
		goto L362
	}
L362:
	;
	goto L1
L363:
	;
	if v1438&int32(1) == int32(0) {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	goto L365
L365:
	;
	if v1438&int32(1) != 0 {
		goto L1
	} else {
		goto L390
	}
L366:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1448 = v1440 + int32(4)
	v1450 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1451 = F_ReorderBufferTXNByXid(m, v31, v1441, int32(0), v1450)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L38
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1533 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferXidSetCatalogChanges(m, v1532, v1441, v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L38
	} else {
		goto L389
	}
L369:
	;
	v1453 = int32(_a_F_xact_decode_10)
	v1454 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v1456
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+40))
	if v1458 != 0 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1494 = F_MemoryContextAlloc(m, v1492, int32(64))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L38
	} else {
		goto L383
	}
L371:
	;
	v1459 = v1458
	goto L373
L372:
	;
	v1459 = v1451
	goto L373
L373:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+172))
	if v1460 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+172)) = v1446
	v1465 = F_palloc_mul(m, int32(16), v1446)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L38
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+176))
	v1476 = F_repalloc_mul(m, v1473, int32(16), v1446+v1460)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L38
	} else {
		goto L379
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+176)) = v1465
	v1469 = v1446 << (uint(int32(4)) % 32)
	if v1469 == int32(0) {
		v1490 = v1469
		goto L370
	} else {
		goto L378
	}
L378:
	;
	base.MemoryCopy(m, v1465, v1448, v1469)
	v1490 = v1469
	goto L370
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+176)) = v1476
	v1480 = v1446 << (uint(int32(4)) % 32)
	if v1480 != 0 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+172))
	base.MemoryCopy(m, v1476+v1481<<(uint(int32(4))%32), v1448, v1480)
	goto L382
L381:
	;
	goto L382
L382:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+172)) = v1486 + v1446
	v1490 = v1480
	goto L370
L383:
	;
	v1496 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+16)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+8)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+56)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+48)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+40)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+32)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+24)) = v1496
	*(*int64)(unsafe.Add(mBase, uint32(v1494))) = v1496
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+20)) = v1446
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+8)) = int32(4)
	v1516 = F_palloc_mul(m, int32(16), v1446)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L38
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+24)) = v1516
	if v1490 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	base.MemoryCopy(m, v1516, v1448, v1490)
	goto L387
L386:
	;
	goto L387
L387:
	;
	F_ReorderBufferQueueChange(m, v31, v1441, v1450, v1494, int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L38
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v1454
	goto L368
L389:
	;
	goto L1
L390:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1542 = int32(0)
	v1544 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[1]))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+24))
	v1547 = base.B2i32(v1545 != v1542)
	goto L391
L391:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[2]))
	v1551 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	if v1545 != v1542 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_xact_decode_15))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L38
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	if v1539 != 0 {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L38
	} else {
		goto L396
	}
L396:
	;
	goto L394
L397:
	;
	v1557 = v1542
	goto L400
L398:
	;
	goto L399
L399:
	;
	if v1545 != v1542 {
		goto L404
	} else {
		goto L405
	}
L400:
	;
	F_LocalExecuteInvalidationMessage(m, v1440+int32(4)+v1557<<(uint(int32(4))%32))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L38
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v1589 = v1557 + int32(1)
	if v1589 != v1539 {
		v1557 = v1589
		goto L400
	} else {
		goto L403
	}
L403:
	;
	goto L401
L404:
	;
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L38
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	goto L1
L407:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[2])) = v1549
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v1551
	goto L406
L408:
	;
	v1769 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627)+54)))
	v1770 = int32(7)
	v1774 = v1658 + (v1769+v1770)&int32(_a_F_xact_decode_16)
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+24)) = v1774
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+28))
	v1781 = int32(-8)
	v1783 = v1774 + (v1776<<(uint(int32(2))%32)+v1770)&v1781
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+32)) = v1783
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+32))
	v1786 = int32(12)
	v1792 = v1783 + (v1785*v1786+v1770)&v1781
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+260)) = v1792
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+36))
	v1801 = v1792 + (v1794*v1786+v1770)&v1781
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+40)) = v1801
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+40))
	v1804 = int32(4)
	v1806 = v1801 + v1803<<(uint(v1804)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+268)) = v1806
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+48)) = v1806 + v1808<<(uint(v1804)%32)
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v1814 == int32(1) {
		goto L434
	} else {
		goto L435
	}
L409:
	;
	F___memset(m, v1765, int32(0), v1764)
	mBase = m.M
	goto L408
L410:
	;
	v1764 = int32(0)
	v1765 = v1759
	goto L409
L411:
	;
	v1742 = v1737
	v1743 = v1738
	v1744 = v1739
	goto L429
L412:
	;
	if v1731 == int32(0) {
		v1759 = v1732
		goto L410
	} else {
		goto L428
	}
L413:
	;
	v1665 = int32(0)
	if base.B2i32(v1658&int32(3) == v1665)|base.B2i32(v1659 == v1665) != 0 {
		v1696 = v1658
		v1697 = v1659
		v1698 = v1656
		v1699 = base.B2i32(v1659 != v1665)
		goto L414
	} else {
		goto L415
	}
L414:
	;
	if v1699 == int32(0) {
		v1759 = v1698
		goto L410
	} else {
		goto L421
	}
L415:
	;
	v1675 = v1658
	v1676 = v1659
	v1677 = v1656
	goto L416
L416:
	;
	v1679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1677))) = uint8(v1679)
	if v1679 == int32(0) {
		v1764 = v1676
		v1765 = v1677
		goto L409
	} else {
		goto L418
	}
L417:
	;
	v1696 = v1690
	v1697 = v1686
	v1698 = v1684
	v1699 = v1688
	goto L414
L418:
	;
	v1683 = int32(1)
	v1684 = v1677 + v1683
	v1686 = v1676 - v1683
	v1687 = int32(0)
	v1688 = base.B2i32(v1686 != v1687)
	v1690 = v1675 + v1683
	if v1690&int32(3) == v1687 {
		v1696 = v1690
		v1697 = v1686
		v1698 = v1684
		v1699 = v1688
		goto L414
	} else {
		goto L419
	}
L419:
	;
	if v1686 != 0 {
		v1675 = v1690
		v1676 = v1686
		v1677 = v1684
		goto L416
	} else {
		goto L420
	}
L420:
	;
	goto L417
L421:
	;
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1696))))
	if v1702 == int32(0) {
		v1764 = v1697
		v1765 = v1698
		goto L409
	} else {
		goto L422
	}
L422:
	;
	if base.Ui32(v1697) < base.Ui32(int32(4)) {
		v1730 = v1696
		v1731 = v1697
		v1732 = v1698
		goto L412
	} else {
		goto L423
	}
L423:
	;
	v1708 = v1696
	v1709 = v1697
	v1710 = v1698
	goto L424
L424:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1708)))
	v1716 = int32(-2139062144)
	if (int32(16843008)-v1713|v1713)&v1716 != v1716 {
		v1737 = v1708
		v1738 = v1709
		v1739 = v1710
		goto L411
	} else {
		goto L426
	}
L425:
	;
	v1730 = v1724
	v1731 = v1726
	v1732 = v1722
	goto L412
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1710))) = v1713
	v1721 = int32(4)
	v1722 = v1710 + v1721
	v1724 = v1708 + v1721
	v1726 = v1709 - v1721
	if base.Ui32(int32(3)) < base.Ui32(v1726) {
		v1708 = v1724
		v1709 = v1726
		v1710 = v1722
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	v1737 = v1730
	v1738 = v1731
	v1739 = v1732
	goto L411
L429:
	;
	v1746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1742))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1744))) = uint8(v1746)
	if v1746 == int32(0) {
		v1764 = v1743
		v1765 = v1744
		goto L409
	} else {
		goto L431
	}
L430:
	;
	v1759 = v1751
	goto L410
L431:
	;
	v1750 = int32(1)
	v1751 = v1744 + v1750
	v1755 = v1743 - v1750
	if v1755 != 0 {
		v1742 = v1742 + v1750
		v1743 = v1755
		v1744 = v1751
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1835 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1836 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1837 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	v1838 = *(*int64)(unsafe.Add(mBase, uint32(v29)+296))
	if v1838 == int64(0) {
		goto L441
	} else {
		goto L442
	}
L434:
	;
	v1818 = v29 + int32(72)
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1819 == int32(0) {
		v1832 = v1813
		goto L433
	} else {
		goto L437
	}
L435:
	;
	v1827 = v1813
	goto L436
L436:
	;
	v1829 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v31, v1827, v1829)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L38
	} else {
		goto L440
	}
L437:
	;
	v1822 = F_filter_prepare_cb_wrapper(m, l0, v1813, v1818)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L38
	} else {
		goto L438
	}
L438:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v1822 == int32(0) {
		v1832 = v1824
		goto L433
	} else {
		goto L439
	}
L439:
	;
	v1827 = v1824
	goto L436
L440:
	;
	goto L1
L441:
	;
	v1841 = v1837
	goto L443
L442:
	;
	v1841 = v1838
	goto L443
L443:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1842)+96))
	v1844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1843)+56)))
	v1845 = *(*int64)(unsafe.Add(mBase, uint32(v29)+288))
	v1846 = m.G0
	v1848 = v1846 - int32(16)
	m.G0 = v1848
	*(*int32)(unsafe.Add(mBase, uint32(v1848)+12)) = v1832
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+32))
	v1852 = int32(0)
	if base.B2i32(v1851 == v1852)|base.B2i32(v1832 != v1851) == v1852 {
		goto L446
	} else {
		goto L447
	}
L444:
	;
	m.G0 = v1848 + int32(16)
	if v1893 == int32(0) {
		goto L1
	} else {
		goto L455
	}
L445:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1880)+72)) = v1841
	*(*int64)(unsafe.Add(mBase, uint32(v1880)+32)) = v1836
	*(*int64)(unsafe.Add(mBase, uint32(v1880)+24)) = v1835
	*(*int64)(unsafe.Add(mBase, uint32(v1880)+64)) = v1845
	*(*uint16)(unsafe.Add(mBase, uint32(v1880)+56)) = uint16(v1844)
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1880)))
	*(*int32)(unsafe.Add(mBase, uint32(v1880))) = v1887 | int32(64)
	v1893 = int32(1)
	goto L444
L446:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+36))
	if v1858 != 0 {
		v1880 = v1858
		goto L445
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	v1865 = F_hash_search(m, v1859, v1848+int32(12), int32(0), v1848+int32(11))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L38
	} else {
		goto L450
	}
L449:
	;
	v1893 = v3
	goto L444
L450:
	;
	v1867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1848)+11)))
	if v1867 == int32(0) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1848)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+32)) = v1870
	v1893 = v3
	goto L444
L452:
	;
	goto L453
L453:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1848)+12))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+36)) = v1875
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+32)) = v1874
	if v1875 == int32(0) {
		v1893 = v3
		goto L444
	} else {
		goto L454
	}
L454:
	;
	v1880 = v1875
	goto L445
L455:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	if v1900 <= int32(1) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ReorderBufferSkipPrepare(m, v1903, v1832)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L38
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1908 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1909 = *(*int64)(unsafe.Add(mBase, uint32(v1907)+16))
	goto L461
L459:
	;
	goto L1
L460:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ReorderBufferSkipPrepare(m, v2062, v1832)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L38
	} else {
		goto L499
	}
L461:
	;
	if base.Ui64(v1908) < base.Ui64(v1909) {
		goto L460
	} else {
		goto L462
	}
L462:
	;
	if v1906 != 0 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+88))
	if v1906 != v1912 {
		goto L460
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1914 != 0 {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	goto L465
L467:
	;
	v1915 = F_filter_by_origin_cb_wrapper(m, l0, v1844)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L38
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v1917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1917 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	if v1915 != 0 {
		goto L460
	} else {
		goto L471
	}
L471:
	;
	goto L469
L472:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if int32(0) < v1920 {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	goto L474
L474:
	;
	v2060 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v2060)
	goto L460
L475:
	;
	v1926 = int32(0)
	goto L478
L476:
	;
	goto L477
L477:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1991 = m.G0
	v1993 = v1991 - int32(16)
	m.G0 = v1993
	*(*int32)(unsafe.Add(mBase, uint32(v1993)+12)) = v1832
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+32))
	v1997 = int32(0)
	if base.B2i32(v1996 == v1997)|base.B2i32(v1996 != v1832) == v1997 {
		goto L484
	} else {
		goto L485
	}
L478:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1951+v1926<<(uint(int32(2))%32))))
	v1956 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1957 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ReorderBufferCommitChild(m, v1950, v1832, v1955, v1956, v1957)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L38
	} else {
		goto L480
	}
L479:
	;
	goto L477
L480:
	;
	v1961 = v1926 + int32(1)
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v1961 < v1962 {
		v1926 = v1961
		goto L478
	} else {
		goto L481
	}
L481:
	;
	goto L479
L482:
	;
	m.G0 = v1993 + int32(16)
	F_UpdateDecodingStats(m, l0)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L38
	} else {
		goto L498
	}
L483:
	;
	v2028 = F_pstrdup(m, v1818)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L38
	} else {
		goto L493
	}
L484:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+36))
	if v2003 != 0 {
		v2026 = v2003
		goto L483
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1990)))
	v2010 = F_hash_search(m, v2004, v1993+int32(12), int32(0), v1993+int32(11))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L38
	} else {
		goto L488
	}
L487:
	;
	goto L482
L488:
	;
	v2012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993)+11)))
	if v2012 == int32(0) {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+32)) = v2015
	goto L482
L490:
	;
	goto L491
L491:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+12))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+36)) = v2020
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+32)) = v2019
	if v2020 == int32(0) {
		goto L482
	} else {
		goto L492
	}
L492:
	;
	v2026 = v2020
	goto L483
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2026)+12)) = v2028
	v2031 = *(*int64)(unsafe.Add(mBase, uint32(v2026)+24))
	v2032 = *(*int64)(unsafe.Add(mBase, uint32(v2026)+32))
	v2033 = *(*int64)(unsafe.Add(mBase, uint32(v2026)+72))
	v2034 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2026)+56)))
	v2035 = *(*int64)(unsafe.Add(mBase, uint32(v2026)+64))
	F_ReorderBufferReplay(m, v2026, v1990, v2031, v2032, v2033, v2034, v2035)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L38
	} else {
		goto L494
	}
L494:
	;
	v2038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026)+1)))
	if v2038&int32(2) != 0 {
		goto L482
	} else {
		goto L495
	}
L495:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+80))
	if v2041 == int32(0) {
		goto L482
	} else {
		goto L496
	}
L496:
	;
	v2044 = *(*int64)(unsafe.Add(mBase, uint32(v2026)+24))
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+64))
	m.T0[v2045].(func(*base.Module, int32, int32, int64))(m, v1990, v2026, v2044)
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L38
	} else {
		goto L497
	}
L497:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2026)))
	*(*int32)(unsafe.Add(mBase, uint32(v2026))) = v2048 | int32(512)
	goto L482
L498:
	;
	goto L1
L499:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2067 = m.G0
	v2069 = v2067 - int32(16)
	m.G0 = v2069
	*(*int32)(unsafe.Add(mBase, uint32(v2069)+12)) = v1832
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2066)+32))
	v2073 = int32(0)
	if base.B2i32(v2072 == v2073)|base.B2i32(v2072 != v1832) == v2073 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	m.G0 = v2069 + int32(16)
	goto L1
L501:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+80))
	if v2103 == int32(0) {
		goto L500
	} else {
		goto L511
	}
L502:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2066)+36))
	if v2079 != 0 {
		v2101 = v2079
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2066)))
	v2086 = F_hash_search(m, v2080, v2069+int32(12), int32(0), v2069+int32(11))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L38
	} else {
		goto L506
	}
L505:
	;
	goto L500
L506:
	;
	v2088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2069)+11)))
	if v2088 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2066)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2066)+32)) = v2091
	goto L500
L508:
	;
	goto L509
L509:
	;
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+12))
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2086)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2066)+36)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v2066)+32)) = v2095
	if v2096 == int32(0) {
		goto L500
	} else {
		goto L510
	}
L510:
	;
	v2101 = v2096
	goto L501
L511:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+172))
	if v2106 == int32(0) {
		goto L500
	} else {
		goto L512
	}
L512:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+176))
	v2110 = int32(0)
	v2112 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[1]))
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+24))
	v2115 = base.B2i32(v2113 != v2110)
	goto L513
L513:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[2]))
	v2119 = *(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0]))
	if v2113 != v2110 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_xact_decode_15))
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L38
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2125 = v2110
	goto L519
L517:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L38
	} else {
		goto L518
	}
L518:
	;
	goto L516
L519:
	;
	F_LocalExecuteInvalidationMessage(m, v2109+v2125<<(uint(int32(4))%32))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L38
	} else {
		goto L521
	}
L520:
	;
	if v2115 == int32(0) {
		goto L500
	} else {
		goto L523
	}
L521:
	;
	v2157 = v2125 + int32(1)
	if v2157 != v2106 {
		v2125 = v2157
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
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L38
	} else {
		goto L524
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[2])) = v2117
	*(*int32)(unsafe.Add(mBase, _c_F_xact_decode[0])) = v2119
	goto L500
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	F_errmsg_internal(m, int32(_a_F_xact_decode_17), v29)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L38
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_xact_decode_18), int32(347), int32(_a_F_xact_decode_19))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
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
	v2216 = int32(0)
	goto L531
L529:
	;
	goto L530
L530:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2281 = *(*int64)(unsafe.Add(mBase, uint32(v2280)+40))
	F_ReorderBufferAbort(m, v2279, v1376, v2281, v1403)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L38
	} else {
		goto L535
	}
L531:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2240+v2216<<(uint(int32(2))%32))))
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2246 = *(*int64)(unsafe.Add(mBase, uint32(v2245)+40))
	F_ReorderBufferAbort(m, v2239, v2244, v2246, v1403)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L38
	} else {
		goto L533
	}
L532:
	;
	goto L530
L533:
	;
	v2250 = v2216 + int32(1)
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v2250 < v2251 {
		v2216 = v2250
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
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L38
	} else {
		goto L536
	}
L536:
	;
	goto L1
}
