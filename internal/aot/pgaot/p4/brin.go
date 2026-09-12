package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__brin_end_parallel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		if int32(0) < v8 {
			v12 = int32(0)
			for {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v17 = v14 + v12<<(uint(int32(7))%32)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v21 = v18 + v12<<(uint(int32(5))%32)
				v22 = int32(4413512)
				v24 = *(*int64)(unsafe.Add(mBase, _consts[5]))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, _consts[5])) = v24 + v25
				v28 = int32(4413520)
				v30 = *(*int64)(unsafe.Add(mBase, _consts[6]))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, _consts[6])) = v30 + v31
				v34 = int32(4413528)
				v36 = *(*int64)(unsafe.Add(mBase, _consts[7]))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
				*(*int64)(unsafe.Add(mBase, _consts[7])) = v36 + v37
				v40 = int32(4413536)
				v42 = *(*int64)(unsafe.Add(mBase, _consts[8]))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int64)(unsafe.Add(mBase, _consts[8])) = v42 + v43
				v46 = int32(4413544)
				v48 = *(*int64)(unsafe.Add(mBase, _consts[9]))
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
				*(*int64)(unsafe.Add(mBase, _consts[9])) = v48 + v49
				v52 = int32(4413552)
				v54 = *(*int64)(unsafe.Add(mBase, _consts[10]))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
				*(*int64)(unsafe.Add(mBase, _consts[10])) = v54 + v55
				v58 = int32(4413560)
				v60 = *(*int64)(unsafe.Add(mBase, _consts[11]))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
				*(*int64)(unsafe.Add(mBase, _consts[11])) = v60 + v61
				v64 = int32(4413568)
				v66 = *(*int64)(unsafe.Add(mBase, _consts[12]))
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
				*(*int64)(unsafe.Add(mBase, _consts[12])) = v66 + v67
				v70 = int32(4413576)
				v72 = *(*int64)(unsafe.Add(mBase, _consts[13]))
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
				*(*int64)(unsafe.Add(mBase, _consts[13])) = v72 + v73
				v76 = int32(4413584)
				v78 = *(*int64)(unsafe.Add(mBase, _consts[14]))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
				*(*int64)(unsafe.Add(mBase, _consts[14])) = v78 + v79
				v82 = int32(4413592)
				v84 = *(*int64)(unsafe.Add(mBase, _consts[15]))
				v85 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
				*(*int64)(unsafe.Add(mBase, _consts[15])) = v84 + v85
				v88 = int32(4413600)
				v90 = *(*int64)(unsafe.Add(mBase, _consts[16]))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
				*(*int64)(unsafe.Add(mBase, _consts[16])) = v90 + v91
				v94 = int32(4413608)
				v96 = *(*int64)(unsafe.Add(mBase, _consts[17]))
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
				*(*int64)(unsafe.Add(mBase, _consts[17])) = v96 + v97
				v100 = int32(4413616)
				v102 = *(*int64)(unsafe.Add(mBase, _consts[18]))
				v103 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
				*(*int64)(unsafe.Add(mBase, _consts[18])) = v102 + v103
				v106 = int32(4413624)
				v108 = *(*int64)(unsafe.Add(mBase, _consts[19]))
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
				*(*int64)(unsafe.Add(mBase, _consts[19])) = v108 + v109
				v112 = int32(4413632)
				v114 = *(*int64)(unsafe.Add(mBase, _consts[20]))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
				*(*int64)(unsafe.Add(mBase, _consts[20])) = v114 + v115
				v118 = int32(4413656)
				v120 = *(*int64)(unsafe.Add(mBase, _consts[21]))
				v121 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
				*(*int64)(unsafe.Add(mBase, _consts[21])) = v120 + v121
				v124 = int32(4413640)
				v126 = *(*int64)(unsafe.Add(mBase, _consts[22]))
				v127 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, _consts[22])) = v126 + v127
				v130 = int32(4413648)
				v132 = *(*int64)(unsafe.Add(mBase, _consts[23]))
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				*(*int64)(unsafe.Add(mBase, _consts[23])) = v132 + v133
				v136 = int32(4413664)
				v138 = *(*int64)(unsafe.Add(mBase, _consts[24]))
				v139 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
				*(*int64)(unsafe.Add(mBase, _consts[24])) = v138 + v139
				v143 = v12 + int32(1)
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
				if v143 < v145 {
					v12 = v143
					continue
				} else {
					break
				}
				break
			}
			v149 = v144
		} else {
			v149 = v7
		}
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
		switch v151 {
		case 0, 5:
			F_UnregisterSnapshot(m, v150)
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return
			} else {
				v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v155 = v154
				F_DestroyParallelContext(m, v155)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return
				} else {
					v160 = *(*int32)(unsafe.Add(mBase, _consts[25]))
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v160)+72)) = v161 - int32(1)
					return
				}
			}
		default:
			v155 = v149
			F_DestroyParallelContext(m, v155)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return
			} else {
				v160 = *(*int32)(unsafe.Add(mBase, _consts[25]))
				v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v160)+72)) = v161 - int32(1)
				return
			}
		}
	}
}
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
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
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
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
	var v653 int32
	_ = v653
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v703 int32
	_ = v703
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v127&int32(64) != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_MemoryContextReset(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v107 = F_brin_new_memtuple(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L13
	}
L5:
	;
	return int32(0)
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if int32(0) < v29 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = int32(20)
	v46 = l2 + (v29*v32+int32(31))&int32(-8)
	v48 = v4
	goto L10
L8:
	;
	goto L9
L9:
	;
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)) = uint8(v105)
	v111 = l2
	goto L1
L10:
	;
	v63 = l2 + int32(24) + v48*int32(20)
	v65 = v48 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v65)
	v67 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v46
	v72 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)) = uint16(v72)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v74
	v76 = int32(2)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0+v32+v48<<(uint(v76)%32))))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v65 < v85 {
		v46 = v46 + v80<<(uint(v76)%32)
		v48 = v65
		goto L10
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v111 = v107
	goto L1
L14:
	;
	v130 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v130)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v133 = v132
	goto L16
L15:
	;
	v133 = v127
	goto L16
L16:
	;
	if v133&int32(32) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)) = uint8(v138)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if int32(0) < v147 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v151 = l1 + int32(5)
	v152 = int32(0)
	v160 = v152
	goto L23
L21:
	;
	v215 = v147
	v221 = v146
	goto L22
L22:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v230 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	if base.B2i32(v152 <= base.I32_extend8_s(v142)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v215 = v210
	v221 = v209
	goto L22
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v143))) = uint8(v205)
	v208 = v160 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v208 < v210 {
		v160 = v208
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v178 = int32(3)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(base.Ui32(v160)>>(uint(v178)%32))))))
	v182 = int32(7)
	v185 = int32(1)
	v186 = int32(base.Ui32(v181)>>(uint(v160&v182)%32)) & v185
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v144))) = uint8(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = v189 + v160
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v190>>(uint(v178)%32)))))
	v205 = int32(base.Ui32(v194)>>(uint(v190&v182)%32)) & v185
	goto L25
L27:
	;
	goto L28
L28:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v144))) = uint8(v201)
	v205 = v201
	goto L25
L29:
	;
	goto L24
L30:
	;
	v233 = int32(4515248)
	v234 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v239 = F_CreateTemplateTupleDesc(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v356 = v215
	v360 = v230
	v362 = v221
	goto L32
L32:
	;
	if int32(0) < v356 {
		goto L47
	} else {
		goto L48
	}
L33:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if int32(0) < v242 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v253 = int32(1)
	v256 = v4
	v257 = v241
	v259 = v242
	goto L37
L35:
	;
	v340 = v241
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v356 = v352
	v360 = v239
	v362 = v340
	goto L32
L37:
	;
	v269 = l0 + int32(20) + v256<<(uint(int32(2))%32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270))))
	if v271 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v340 = v319
	goto L36
L39:
	;
	v275 = v270
	v276 = int32(0)
	v277 = v253
	goto L42
L40:
	;
	v315 = v253
	v319 = v257
	v321 = v259
	goto L41
L41:
	;
	v329 = v256 + int32(1)
	if v329 < v321 {
		v253 = v315
		v256 = v329
		v257 = v319
		v259 = v321
		goto L37
	} else {
		goto L46
	}
L42:
	;
	v291 = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v275+v276<<(uint(int32(2))%32))+8))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	F_TupleDescInitEntry(m, v239, base.I32_extend16_s(v277), v291, v296, int32(-1), v291)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v315 = v302
	v319 = v308
	v321 = v309
	goto L41
L44:
	;
	v301 = int32(1)
	v302 = v277 + v301
	v304 = v276 + v301
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305))))
	if base.Ui32(v304) < base.Ui32(v306) {
		v275 = v305
		v276 = v304
		v277 = v302
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L38
L47:
	;
	v375 = l1 + v142&int32(31)
	v376 = int32(20)
	v380 = int32(0)
	v386 = v356
	v387 = v380
	v388 = v380
	v392 = v362
	v393 = v380
	goto L50
L48:
	;
	v553 = v362
	goto L49
L49:
	;
	v562 = int32(4515248)
	v563 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	if int32(0) < v567 {
		goto L97
	} else {
		goto L98
	}
L50:
	;
	v403 = l0 + v376 + v393<<(uint(int32(2))%32)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404))))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v144))))
	if v407 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v553 = v532
	goto L49
L52:
	;
	v542 = v393 + int32(1)
	if v542 < v526 {
		v386 = v526
		v387 = v527
		v388 = v528
		v392 = v532
		v393 = v542
		goto L50
	} else {
		goto L96
	}
L53:
	;
	v410 = int32(0)
	if v405 == v410 {
		v526 = v386
		v527 = v387
		v528 = v388
		v532 = v392
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v526 = v386
	v527 = v387
	v528 = v388 + v405
	v532 = v392
	goto L52
L56:
	;
	v417 = v387
	v418 = v388
	v420 = v410
	goto L57
L57:
	;
	v433 = v360 + v376 + v418<<(uint(int32(4))%32)
	v434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v433)+4)))
	v435 = int32(65535)
	v436 = v434 & v435
	if v436 == v435 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v526 = v521
	v527 = v514
	v528 = v513
	v532 = v520
	goto L52
L59:
	;
	v450 = v448 + v375
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+6)))
	if v454 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v375))))
	if v440 != 0 {
		v448 = v417
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+12)))
	v448 = (v417 + v441 - int32(1)) & (int32(0) - v441)
	goto L59
L63:
	;
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v418<<(uint(int32(2))%32)))) = v475
	v477 = int32(*(*int16)(unsafe.Add(mBase, uint32(v433)+4)))
	if int32(0) < v477 {
		v511 = v477
		goto L75
	} else {
		goto L76
	}
L65:
	;
	switch v436 - int32(1) {
	case 0:
		goto L71
	case 1:
		goto L70
	default:
		goto L68
	case 3:
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v475 = v450
	goto L64
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L72
	}
L69:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v475 = v461
	goto L64
L70:
	;
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v450))))
	v475 = v460
	goto L64
L71:
	;
	v459 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	v475 = v459
	goto L64
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v434
	F_errmsg_internal(m, int32(483136), v21)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(326451), int32(70), int32(67779))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v512 = int32(1)
	v513 = v418 + v512
	v514 = v511 + v448
	v516 = v420 + v512
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517))))
	if base.Ui32(v516) < base.Ui32(v518) {
		v417 = v514
		v418 = v513
		v420 = v516
		goto L57
	} else {
		goto L95
	}
L76:
	;
	if v477 == int32(-1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	if v482 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v508 = F_strlen(m, v450)
	mBase = m.M
	v511 = v508 + int32(1)
	goto L75
L80:
	;
	v485 = int32(6)
	v487 = int32(18)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+1)))
	if v489 == v487 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if v482&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L83:
	;
	v492 = v487
	goto L85
L84:
	;
	v492 = int32(2)
	goto L85
L85:
	;
	if v489&int32(254) == int32(2) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v497 = v485
	goto L88
L87:
	;
	v497 = v492
	goto L88
L88:
	;
	if v489 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v500 = v485
	goto L91
L90:
	;
	v500 = v497
	goto L91
L91:
	;
	v511 = v500
	goto L75
L92:
	;
	v511 = int32(base.Ui32(v482) >> (uint(int32(1)) % 32))
	goto L75
L93:
	;
	goto L94
L94:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v511 = int32(base.Ui32(v505) >> (uint(int32(2)) % 32))
	goto L75
L95:
	;
	goto L58
L96:
	;
	goto L51
L97:
	;
	v576 = int32(0)
	v582 = v576
	v584 = v567
	v586 = v576
	goto L100
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v563
	m.G0 = v21 + int32(16)
	return v111
L100:
	;
	v598 = l0 + int32(20) + v586<<(uint(int32(2))%32)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v599))))
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586+v144))))
	if v602 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L99
L102:
	;
	v703 = v586 + int32(1)
	if v703 < v690 {
		v582 = v688
		v584 = v690
		v586 = v703
		goto L100
	} else {
		goto L113
	}
L103:
	;
	if v600 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v688 = v582 + v600
	v690 = v584
	goto L102
L106:
	;
	v612 = v599
	v613 = v582
	v614 = int32(0)
	goto L109
L107:
	;
	v653 = v582
	goto L108
L108:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586+v143))))
	v671 = v111 + int32(24) + v586*int32(20)
	v672 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v671)+8)) = v672
	*(*uint8)(unsafe.Add(mBase, uint32(v671)+3)) = uint8(v672)
	*(*uint8)(unsafe.Add(mBase, uint32(v671)+2)) = uint8(v668)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+12)) = v679
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v688 = v653
	v690 = v682
	goto L102
L109:
	;
	v627 = int32(2)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v145+v613<<(uint(v627)%32))))
	v632 = v614 << (uint(v627) % 32)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v612+v632)+8))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+10)))
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v634)+8)))
	v637 = F_datumCopy(m, v630, v635, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L111
	}
L110:
	;
	v653 = v643
	goto L108
L111:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(28)+v586*int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v639+v632))) = v637
	v642 = int32(1)
	v643 = v613 + v642
	v645 = v614 + v642
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646))))
	if base.Ui32(v645) < base.Ui32(v647) {
		v612 = v646
		v613 = v643
		v614 = v645
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L101
}
func F_brin_initialize_empty_new_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	v3 = int32(0)
	v6 = int32(4509908)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v8 + int32(1)
	if l1 < v3 {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(l1^int32(-1))<<(uint(int32(2))%32))))
		v29 = v21
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v29 = v23 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	if v29&int32(3) != 0 {
	} else {
	}
	v56 = F___memset(m, v29, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v29)+10)) = int32(1572864)
	v62 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+18)) = uint16(v62)
	v68 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v73 = int32(61587)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+v71)+6)) = uint16(v73)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		return
	} else {
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+118)))
		if v78 != int32(112) {
			v90 = int32(4509908)
			v92 = *(*int32)(unsafe.Add(mBase, _consts[26]))
			*(*int32)(unsafe.Add(mBase, _consts[26])) = v92 - int32(1)
			if l1 < int32(0) {
				v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(l1^int32(-1))<<(uint(int32(6))%32))+16))
				v114 = v105
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, _consts[4]))
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+l1<<(uint(int32(6))%32)+int32(-64))+16))
				v114 = v113
			}
			v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
			v116 = v29 + v115
			v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
			if v117 != int32(61587) {
				v132 = v3
			} else {
				v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
				if v120&int32(1) != 0 {
					v132 = v3
				} else {
					v123 = int32(4)
					v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
					v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
					v126 = v124 - v125
					if v126 <= v123 {
						v129 = v123
					} else {
						v129 = v126
					}
					v132 = v129 - int32(4)
				}
			}
			F_RecordPageWithFreeSpace(m, l0, v114, v132)
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return
			} else {
				return
			}
		} else {
			v82 = *(*int32)(unsafe.Add(mBase, _consts[27]))
			if v82 <= int32(0) {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v85 != 0 {
					v90 = int32(4509908)
					v92 = *(*int32)(unsafe.Add(mBase, _consts[26]))
					*(*int32)(unsafe.Add(mBase, _consts[26])) = v92 - int32(1)
					if l1 < int32(0) {
						v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v114 = v105
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v114 = v113
					}
					v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
					v116 = v29 + v115
					v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
					if v117 != int32(61587) {
						v132 = v3
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
						if v120&int32(1) != 0 {
							v132 = v3
						} else {
							v123 = int32(4)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
							v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
							v126 = v124 - v125
							if v126 <= v123 {
								v129 = v123
							} else {
								v129 = v126
							}
							v132 = v129 - int32(4)
						}
					}
					F_RecordPageWithFreeSpace(m, l0, v114, v132)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						return
					}
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v86 != 0 {
						v90 = int32(4509908)
						v92 = *(*int32)(unsafe.Add(mBase, _consts[26]))
						*(*int32)(unsafe.Add(mBase, _consts[26])) = v92 - int32(1)
						if l1 < int32(0) {
							v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(l1^int32(-1))<<(uint(int32(6))%32))+16))
							v114 = v105
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+l1<<(uint(int32(6))%32)+int32(-64))+16))
							v114 = v113
						}
						v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
						v116 = v29 + v115
						v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
						if v117 != int32(61587) {
							v132 = v3
						} else {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
							if v120&int32(1) != 0 {
								v132 = v3
							} else {
								v123 = int32(4)
								v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
								v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
								v126 = v124 - v125
								if v126 <= v123 {
									v129 = v123
								} else {
									v129 = v126
								}
								v132 = v129 - int32(4)
							}
						}
						F_RecordPageWithFreeSpace(m, l0, v114, v132)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							return
						}
					} else {
						F_log_newpage_buffer(m, l1, int32(1))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							v90 = int32(4509908)
							v92 = *(*int32)(unsafe.Add(mBase, _consts[26]))
							*(*int32)(unsafe.Add(mBase, _consts[26])) = v92 - int32(1)
							if l1 < int32(0) {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(l1^int32(-1))<<(uint(int32(6))%32))+16))
								v114 = v105
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+l1<<(uint(int32(6))%32)+int32(-64))+16))
								v114 = v113
							}
							v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
							v116 = v29 + v115
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
							if v117 != int32(61587) {
								v132 = v3
							} else {
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
								if v120&int32(1) != 0 {
									v132 = v3
								} else {
									v123 = int32(4)
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
									v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
									v126 = v124 - v125
									if v126 <= v123 {
										v129 = v123
									} else {
										v129 = v126
									}
									v132 = v129 - int32(4)
								}
							}
							F_RecordPageWithFreeSpace(m, l0, v114, v132)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_log_newpage_buffer(m, l1, int32(1))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					v90 = int32(4509908)
					v92 = *(*int32)(unsafe.Add(mBase, _consts[26]))
					*(*int32)(unsafe.Add(mBase, _consts[26])) = v92 - int32(1)
					if l1 < int32(0) {
						v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v114 = v105
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v114 = v113
					}
					v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
					v116 = v29 + v115
					v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
					if v117 != int32(61587) {
						v132 = v3
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
						if v120&int32(1) != 0 {
							v132 = v3
						} else {
							v123 = int32(4)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
							v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
							v126 = v124 - v125
							if v126 <= v123 {
								v129 = v123
							} else {
								v129 = v126
							}
							v132 = v129 - int32(4)
						}
					}
					F_RecordPageWithFreeSpace(m, l0, v114, v132)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_brin_minmax_multi_distance_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v7 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_minmax_multi_distance_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_DirectFunctionCall2Coll(m, int32(18), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(17), v3, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_brin_minmax_multi_distance_pg_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = F_Float8GetDatum(m, base.F64_convert_i64_u(v3-v5))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_brin_minmax_multi_distance_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = F_Float8GetDatum(m, base.F64_convert_i64_s(v3-v5))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_brin_minmax_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v24 = v14 + v15<<(uint(int32(4))%32) + base.I32_extend16_s(v13)*int32(100) - int32(80)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v27 = F_minmax_get_strategy_procinfo(m, v11, v13, v25, v10)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
		v35 = F_FunctionCall2Coll(m, v27, v9, v32, v34)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 != 0 {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
				if v37 == int32(0) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					F_pfree(m, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
						v45 = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
						v51 = F_datumCopy(m, v47, v45&int32(1), v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v53))) = v51
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
							v58 = F_minmax_get_strategy_procinfo(m, v11, v13, v56, int32(5))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
								v64 = F_FunctionCall2Coll(m, v58, v9, v61, v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									if v64 != 0 {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
										if v67 == int32(0) {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
											F_pfree(m, v71)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
												v75 = v74
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
												v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
												v81 = F_datumCopy(m, v77, v75&int32(1), v80)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
													return int32(0)
												}
											}
										} else {
											v75 = int32(1)
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
											v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
											v81 = F_datumCopy(m, v77, v75&int32(1), v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						}
					}
				} else {
					v45 = v10
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
					v51 = F_datumCopy(m, v47, v45&int32(1), v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = v51
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
						v58 = F_minmax_get_strategy_procinfo(m, v11, v13, v56, int32(5))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
							v64 = F_FunctionCall2Coll(m, v58, v9, v61, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								if v64 != 0 {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
									if v67 == int32(0) {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
										F_pfree(m, v71)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
											v75 = v74
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
											v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
											v81 = F_datumCopy(m, v77, v75&int32(1), v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
												return int32(0)
											}
										}
									} else {
										v75 = int32(1)
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
										v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
										v81 = F_datumCopy(m, v77, v75&int32(1), v80)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
				v58 = F_minmax_get_strategy_procinfo(m, v11, v13, v56, int32(5))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					v64 = F_FunctionCall2Coll(m, v58, v9, v61, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						if v64 != 0 {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
							if v67 == int32(0) {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
								F_pfree(m, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)))
									v75 = v74
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
									v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
									v81 = F_datumCopy(m, v77, v75&int32(1), v80)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
										return int32(0)
									}
								}
							} else {
								v75 = int32(1)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+72)))
								v81 = F_datumCopy(m, v77, v75&int32(1), v80)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_brin_redo(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int64
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	v14 = v12 & int32(240)
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 < int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v22^int32(-1))<<(uint(int32(2))%32))))
				v41 = v33
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				v41 = v35 + v22<<(uint(int32(13))%32) + int32(-8192)
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
			F_PageInit(m, v41, int32(8192), int32(8))
			mBase = m.M
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
			v49 = int32(61585)
			*(*uint16)(unsafe.Add(mBase, uint32(v41+v47)+6)) = uint16(v49)
			*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v42
			*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(-1475306246)
			v57 = int32(40)
			*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)) = uint16(v57)
			*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v20, int64(32))
			F_MarkBufferDirty(m, v22)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_UnlockReleaseBuffer(m, v22)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			}
		}
	case 1:
		v421 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		F_brin_xlog_insert_update(m, l0, v421)
		mBase = m.M
		v423 = m.ExcPending
		if v423 != 0 {
			return
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	case 2:
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v71 = F_XLogReadBufferForRedo(m, l0, int32(2), v9+int32(20))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			if v71 == int32(0) {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v75 < int32(0) {
					v79 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v75^int32(-1))<<(uint(int32(2))%32))))
					v93 = v85
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v93 = v87 + v75<<(uint(int32(13))%32) + int32(-8192)
				}
				v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
				F_PageIndexTupleDeleteNoCompact(m, v93, v94)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93))) = base.I64_rotr(v67, int64(32))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					F_MarkBufferDirty(m, v100)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						F_brin_xlog_insert_update(m, l0, v66+int32(4))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							if v108 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v108)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				F_brin_xlog_insert_update(m, l0, v66+int32(4))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					if v108 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						F_UnlockReleaseBuffer(m, v108)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			}
		}
	case 3:
		v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v118 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(20))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			if v118 == int32(0) {
				v122 = int32(0)
				v124 = v9 + int32(28)
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+72))
				if v127 < v122 {
					v149 = v122
					v152 = v149
				} else {
					v133 = v126 + int32(76)
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
					if v134 != int32(1) {
						v149 = v122
						v152 = v149
					} else {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+43)))
						if v137 == int32(0) {
							if v124 == int32(0) {
								v149 = v122
								v152 = v149
							} else {
								v142 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v124))) = v142
								v152 = v142
							}
						} else {
							if v124 != 0 {
								v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+48)))
								*(*int32)(unsafe.Add(mBase, uint32(v124))) = v145
							} else {
							}
							v147 = *(*int32)(unsafe.Add(mBase, uint32(v133)+44))
							v149 = v147
							v152 = v149
						}
					}
				}
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v153 < int32(0) {
					v157 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v153^int32(-1))<<(uint(int32(2))%32))))
					v171 = v163
				} else {
					v165 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v171 = v165 + v153<<(uint(int32(13))%32) + int32(-8192)
				}
				v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
				v173 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v174 = F_PageIndexTupleOverwrite(m, v171, v172, v152, v173)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return
				} else {
					if v174 == int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v435 = m.ExcPending
						if v435 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(384414), int32(0))
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								F_errfinish(m, int32(497897), int32(193), int32(355871))
								mBase = m.M
								v444 = m.ExcPending
								if v444 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v171))) = base.I64_rotr(v113, int64(32))
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
						F_MarkBufferDirty(m, v181)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return
						} else {
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							if v186 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v186)
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v186 == int32(0) {
					m.G0 = v9 + int32(32)
					return
				} else {
					F_UnlockReleaseBuffer(m, v186)
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	case 4:
		v191 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v194 = int32(0)
		F_XLogRecGetBlockTag(m, l0, int32(1), v194, v194, v9+int32(28))
		mBase = m.M
		v199 = m.ExcPending
		if v199 != 0 {
			return
		} else {
			v203 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(20))
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return
			} else {
				if v203 == int32(0) {
					v207 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
					v208 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					if v208 < int32(0) {
						v212 = *(*int32)(unsafe.Add(mBase, _consts[1]))
						v218 = *(*int32)(unsafe.Add(mBase, uint32(v212+(v208^int32(-1))<<(uint(int32(2))%32))))
						v226 = v218
					} else {
						v220 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v226 = v220 + v208<<(uint(int32(13))%32) + int32(-8192)
					}
					v227 = int32(40)
					*(*uint16)(unsafe.Add(mBase, uint32(v226)+12)) = uint16(v227)
					*(*uint32)(unsafe.Add(mBase, uint32(v226)+4)) = uint32(v191)
					v231 = int64(base.Ui64(v191) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v226))) = uint32(v231)
					*(*int32)(unsafe.Add(mBase, uint32(v226)+36)) = v207
					v234 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					F_MarkBufferDirty(m, v234)
					mBase = m.M
					v236 = m.ExcPending
					if v236 != 0 {
						return
					} else {
						v240 = F_XLogInitBufferForRedo(m, l0, int32(1))
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return
						} else {
							if v240 < int32(0) {
								v245 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+(v240^int32(-1))<<(uint(int32(2))%32))))
								v259 = v251
							} else {
								v253 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v259 = v253 + v240<<(uint(int32(13))%32) + int32(-8192)
							}
							v260 = int32(61586)
							F_PageInit(m, v259, int32(8192), int32(8))
							mBase = m.M
							v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v259+v264)+6)) = uint16(v260)
							*(*int64)(unsafe.Add(mBase, uint32(v259))) = base.I64_rotr(v191, int64(32))
							F_MarkBufferDirty(m, v240)
							mBase = m.M
							v271 = m.ExcPending
							if v271 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v240)
								mBase = m.M
								v273 = m.ExcPending
								if v273 != 0 {
									return
								} else {
									v274 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									if v274 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v274)
										mBase = m.M
										v278 = m.ExcPending
										if v278 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v240 = F_XLogInitBufferForRedo(m, l0, int32(1))
					mBase = m.M
					v241 = m.ExcPending
					if v241 != 0 {
						return
					} else {
						if v240 < int32(0) {
							v245 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+(v240^int32(-1))<<(uint(int32(2))%32))))
							v259 = v251
						} else {
							v253 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v259 = v253 + v240<<(uint(int32(13))%32) + int32(-8192)
						}
						v260 = int32(61586)
						F_PageInit(m, v259, int32(8192), int32(8))
						mBase = m.M
						v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
						*(*uint16)(unsafe.Add(mBase, uint32(v259+v264)+6)) = uint16(v260)
						*(*int64)(unsafe.Add(mBase, uint32(v259))) = base.I64_rotr(v191, int64(32))
						F_MarkBufferDirty(m, v240)
						mBase = m.M
						v271 = m.ExcPending
						if v271 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v240)
							mBase = m.M
							v273 = m.ExcPending
							if v273 != 0 {
								return
							} else {
								v274 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								if v274 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v274)
									mBase = m.M
									v278 = m.ExcPending
									if v278 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	case 5:
		v279 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v280 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v284 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(28))
		mBase = m.M
		v285 = m.ExcPending
		if v285 != 0 {
			return
		} else {
			if v284 == int32(0) {
				v288 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+24)) = uint16(v288)
				v290 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v290
				v292 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
				v293 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)) = uint16(v288)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v290
				v298 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v300 = v9 + int32(12)
				v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300))))
				v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+2)))
				if v298 < v288 {
					v308 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308+(v298^int32(-1))<<(uint(int32(2))%32))))
					v322 = v314
				} else {
					v316 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v322 = v316 + v298<<(uint(int32(13))%32) + int32(-8192)
				}
				v323 = base.I32_div_u_s(v292, v293)
				v325 = base.I32_rem_u_s(v323, int32(1360))
				v328 = v322 + v325*int32(6)
				v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v328)+28)) = uint16(v329)
				if v329 != 0 {
					v332 = v304
				} else {
					v332 = int32(-1)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v328)+26)) = uint16(v332)
				if v329 != 0 {
					v335 = v303
				} else {
					v335 = int32(-1)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v328)+24)) = uint16(v335)
				v337 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				if v337 < int32(0) {
					v341 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v347 = *(*int32)(unsafe.Add(mBase, uint32(v341+(v337^int32(-1))<<(uint(int32(2))%32))))
					v355 = v347
				} else {
					v349 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v355 = v349 + v337<<(uint(int32(13))%32) + int32(-8192)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v355))) = base.I64_rotr(v279, int64(32))
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				F_MarkBufferDirty(m, v359)
				mBase = m.M
				v361 = m.ExcPending
				if v361 != 0 {
					return
				} else {
					v364 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
					if v364 != 0 {
						F_UnlockReleaseBuffer(m, v364)
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return
						} else {
							v370 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
							mBase = m.M
							v371 = m.ExcPending
							if v371 != 0 {
								return
							} else {
								if v370 == int32(0) {
									v374 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v374 < int32(0) {
										v378 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+(v374^int32(-1))<<(uint(int32(2))%32))))
										v392 = v384
									} else {
										v386 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v392 = v386 + v374<<(uint(int32(13))%32) + int32(-8192)
									}
									v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+8)))
									F_PageIndexTupleDeleteNoCompact(m, v392, v393)
									mBase = m.M
									v395 = m.ExcPending
									if v395 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v392))) = base.I64_rotr(v279, int64(32))
										v399 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										F_MarkBufferDirty(m, v399)
										mBase = m.M
										v401 = m.ExcPending
										if v401 != 0 {
											return
										} else {
											v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
											if v403 == int32(0) {
												m.G0 = v9 + int32(32)
												return
											} else {
												F_UnlockReleaseBuffer(m, v403)
												mBase = m.M
												v407 = m.ExcPending
												if v407 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v403 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v403)
										mBase = m.M
										v407 = m.ExcPending
										if v407 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					} else {
						v370 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
						mBase = m.M
						v371 = m.ExcPending
						if v371 != 0 {
							return
						} else {
							if v370 == int32(0) {
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v374 < int32(0) {
									v378 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+(v374^int32(-1))<<(uint(int32(2))%32))))
									v392 = v384
								} else {
									v386 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v392 = v386 + v374<<(uint(int32(13))%32) + int32(-8192)
								}
								v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+8)))
								F_PageIndexTupleDeleteNoCompact(m, v392, v393)
								mBase = m.M
								v395 = m.ExcPending
								if v395 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v392))) = base.I64_rotr(v279, int64(32))
									v399 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									F_MarkBufferDirty(m, v399)
									mBase = m.M
									v401 = m.ExcPending
									if v401 != 0 {
										return
									} else {
										v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										if v403 == int32(0) {
											m.G0 = v9 + int32(32)
											return
										} else {
											F_UnlockReleaseBuffer(m, v403)
											mBase = m.M
											v407 = m.ExcPending
											if v407 != 0 {
												return
											} else {
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v403 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v403)
									mBase = m.M
									v407 = m.ExcPending
									if v407 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v364 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				if v364 != 0 {
					F_UnlockReleaseBuffer(m, v364)
					mBase = m.M
					v366 = m.ExcPending
					if v366 != 0 {
						return
					} else {
						v370 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
						mBase = m.M
						v371 = m.ExcPending
						if v371 != 0 {
							return
						} else {
							if v370 == int32(0) {
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v374 < int32(0) {
									v378 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+(v374^int32(-1))<<(uint(int32(2))%32))))
									v392 = v384
								} else {
									v386 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v392 = v386 + v374<<(uint(int32(13))%32) + int32(-8192)
								}
								v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+8)))
								F_PageIndexTupleDeleteNoCompact(m, v392, v393)
								mBase = m.M
								v395 = m.ExcPending
								if v395 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v392))) = base.I64_rotr(v279, int64(32))
									v399 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									F_MarkBufferDirty(m, v399)
									mBase = m.M
									v401 = m.ExcPending
									if v401 != 0 {
										return
									} else {
										v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										if v403 == int32(0) {
											m.G0 = v9 + int32(32)
											return
										} else {
											F_UnlockReleaseBuffer(m, v403)
											mBase = m.M
											v407 = m.ExcPending
											if v407 != 0 {
												return
											} else {
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v403 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v403)
									mBase = m.M
									v407 = m.ExcPending
									if v407 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v370 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
					mBase = m.M
					v371 = m.ExcPending
					if v371 != 0 {
						return
					} else {
						if v370 == int32(0) {
							v374 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
							if v374 < int32(0) {
								v378 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+(v374^int32(-1))<<(uint(int32(2))%32))))
								v392 = v384
							} else {
								v386 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v392 = v386 + v374<<(uint(int32(13))%32) + int32(-8192)
							}
							v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+8)))
							F_PageIndexTupleDeleteNoCompact(m, v392, v393)
							mBase = m.M
							v395 = m.ExcPending
							if v395 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v392))) = base.I64_rotr(v279, int64(32))
								v399 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								F_MarkBufferDirty(m, v399)
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return
								} else {
									v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v403 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v403)
										mBase = m.M
										v407 = m.ExcPending
										if v407 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						} else {
							v403 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
							if v403 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v403)
								mBase = m.M
								v407 = m.ExcPending
								if v407 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v411 = m.ExcPending
		if v411 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
			F_errmsg_internal(m, int32(52443), v9)
			mBase = m.M
			v415 = m.ExcPending
			if v415 != 0 {
				return
			} else {
				F_errfinish(m, int32(497897), int32(334), int32(242509))
				mBase = m.M
				v420 = m.ExcPending
				if v420 != 0 {
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
