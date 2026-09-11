package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AfterTriggerEndSubXact(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	v8 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	goto L1
L1:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v13 = v9 * int32(24)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v13)))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	if v25 <= v9 {
		goto L2
	} else {
		goto L11
	}
L6:
	;
	F_pfree(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v20 = v11
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v13))) = int32(0)
	goto L2
L9:
	;
	return
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v20 = v19
	goto L8
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v32 = v9 * int32(24)
	v33 = v30 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 < v28 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v38 = v28
	v39 = v30
	v40 = v37
	goto L15
L13:
	;
	v71 = v33
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v74 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if v38 < v40 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v71 = v65
	goto L14
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	F_AfterTriggerFreeQuery(m, v47+v38*int32(20))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	v59 = v39
	v60 = v40
	v61 = v38
	goto L19
L19:
	;
	v63 = v61 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[284])) = v63
	v65 = v59 + v32
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v66 < v63 {
		v38 = v63
		v39 = v59
		v40 = v60
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v59 = v54
	v60 = v56
	v61 = v58
	goto L19
L21:
	;
	goto L16
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v136 = v9 * int32(24)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v136)))
	if v138 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	goto L26
L24:
	;
	goto L25
L25:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int64)(unsafe.Add(mBase, _consts[325])) = v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	*(*int32)(unsafe.Add(mBase, _consts[326])) = v97
	v100 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v84 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _consts[327])) = int64(0)
	goto L22
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v86
	F_pfree(m, v84)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L26
L32:
	;
	v102 = v101
	goto L35
L33:
	;
	v119 = v100
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	v125 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v125
	goto L22
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	F_pfree(m, v102)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L37
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	v119 = v112
	goto L34
L37:
	;
	if v108 != 0 {
		v102 = v108
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[328]))
	F_pfree(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	v147 = v134
	goto L41
L41:
	;
	v149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v136))) = v149
	v152 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v152 == v149 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[328])) = v138
	v146 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v147 = v146
	goto L41
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156+v136)+20))
	v161 = v152
	goto L44
L44:
	;
	v166 = v161 + int32(16)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if base.Ui32(v166) < base.Ui32(v167) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L2
L46:
	;
	v169 = v166
	goto L49
L47:
	;
	goto L48
L48:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v210 != 0 {
		v161 = v210
		goto L44
	} else {
		goto L61
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if base.Ui32(v175) < base.Ui32(int32(1073741824)) {
		v186 = v175
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	v189 = v186 & int32(939524096)
	if v189 == int32(134217728) {
		v200 = int32(24)
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v169+v175&int32(134217727))+16))
	if base.Ui32(v181) < base.Ui32(v158) {
		v186 = v175
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v184 = v175 & int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v184
	v186 = v184
	goto L51
L54:
	;
	v201 = v200 + v169
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if base.Ui32(v201) < base.Ui32(v202) {
		v169 = v201
		goto L49
	} else {
		goto L60
	}
L55:
	;
	if v189 == int32(268435456) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v196 = int32(12)
	goto L58
L57:
	;
	v196 = int32(4)
	goto L58
L58:
	;
	if v189 != int32(805306368) {
		v200 = v196
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v200 = int32(16)
	goto L54
L60:
	;
	goto L50
L61:
	;
	goto L45
}
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v90 int32
	_ = v90
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v430 int32
	_ = v430
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
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
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
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
	var v522 int32
	_ = v522
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
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
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
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
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v719 int32
	_ = v719
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	v29 = m.G0
	v31 = v29 + int32(-64)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	if int32(0) <= v34 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v31 - int32(-64)
	return
L2:
	;
	if l5 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v359
	v364 = v360
	v365 = int32(1)
	goto L2
L4:
	;
	if l5 != 0 {
		goto L86
	} else {
		goto L87
	}
L5:
	;
	if l5 != 0 {
		goto L81
	} else {
		goto L82
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
	v42 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v34 < v42 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L18
	} else {
		goto L78
	}
L9:
	;
	if l5 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	v45 = v34 + int32(1)
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[316])) = v68
	*(*int32)(unsafe.Add(mBase, _consts[317])) = v70
	if v68 <= v42 {
		goto L9
	} else {
		goto L24
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v50 = int32(8)
	if v45 <= v50 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v61 = v42 << (uint(int32(1)) % 32)
	if v61 < v45 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v53 = v50
	goto L17
L16:
	;
	v53 = v45
	goto L17
L17:
	;
	v56 = F_MemoryContextAlloc(m, v49, v53*int32(20))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v68 = v53
	v70 = v56
	goto L11
L20:
	;
	v63 = v45
	goto L22
L21:
	;
	v63 = v61
	goto L22
L22:
	;
	v66 = F_repalloc(m, v59, v63*int32(20))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v68 = v63
	v70 = v66
	goto L11
L24:
	;
	v90 = v42
	goto L25
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v108 = v105 + v90*int32(20)
	v109 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v109
	v116 = v90 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v116 < v118 {
		v90 = v116
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L9
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(65535)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v341 = int32(32)
	goto L4
L29:
	;
	if l5 != 0 {
		goto L74
	} else {
		goto L75
	}
L30:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+19)))
	if v263 == int32(0) {
		goto L1
	} else {
		goto L73
	}
L31:
	;
	if l5 != 0 {
		goto L69
	} else {
		goto L70
	}
L32:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v238 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L33:
	;
	switch l4 - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L5
	case 2:
		goto L28
	default:
		goto L31
	}
L34:
	;
	if l10 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if l6 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if l7 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v155&int32(2) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if l4 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	F_TransitionTableAddTuple(m, l4, l10, l1, l6, int32(0), v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L18
	} else {
		goto L46
	}
L40:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176+l10)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	v180 = v179
	goto L39
L41:
	;
	v166 = int32(0)
	if l4 != int32(2) {
		v180 = v166
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10))))
	if v160&int32(1) == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v176 = int32(16)
	goto L40
L44:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+1)))
	if v169&int32(1) == int32(0) {
		v180 = v166
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v176 = int32(12)
	goto L40
L46:
	;
	goto L36
L47:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L48:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7)+4)))
	if v187&int32(2) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if l4 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_TransitionTableAddTuple(m, l4, l10, l1, l7, v152, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L18
	} else {
		goto L57
	}
L51:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206+l10)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+28))
	v210 = v209
	goto L50
L52:
	;
	v196 = int32(0)
	if l4 != int32(2) {
		v210 = v196
		goto L50
	} else {
		goto L55
	}
L53:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+3)))
	if v190&int32(1) == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v206 = int32(8)
	goto L51
L55:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+2)))
	if v199&int32(1) == int32(0) {
		v210 = v196
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v206 = int32(12)
	goto L51
L57:
	;
	goto L47
L58:
	;
	switch l4 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L59
	case 2:
		goto L28
	default:
		goto L32
	}
L59:
	;
	v218 = int32(1)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+14)))
	if v219 != v218 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if l6 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
	v223 = int32(1)
	v227 = int32(base.Ui32(v222)>>(uint(v223)%32)) & v223
	goto L63
L62:
	;
	v227 = v218
	goto L63
L63:
	;
	if l7 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)))
	v229 = int32(1)
	v234 = int32(base.Ui32(v228)>>(uint(v229)%32)) & v229
	goto L66
L65:
	;
	v234 = int32(1)
	goto L66
L66:
	;
	if v234 != v227 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L5
L68:
	;
	goto L31
L69:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v245
	v247 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = int32(-1)
	v341 = int32(4)
	goto L4
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(65535)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	F_cancel_prior_stmt_triggers(m, v256, int32(3), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	v359 = int32(268435456)
	v360 = int32(4)
	goto L3
L73:
	;
	goto L29
L74:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v271
	v273 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v273)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = int32(-1)
	v341 = int32(8)
	goto L4
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(65535)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	F_cancel_prior_stmt_triggers(m, v281, int32(4), int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	v359 = int32(268435456)
	v360 = int32(8)
	goto L3
L78:
	;
	F_errmsg_internal(m, int32(15450), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(465125), int32(6194), int32(85898))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v310)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v312
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v314)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = v316
	v318 = int32(16)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+119)))
	if v320 != int32(112) {
		v341 = v318
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(65535)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	v334 = int32(2)
	F_cancel_prior_stmt_triggers(m, v333, v334, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L18
	} else {
		goto L85
	}
L84:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v327
	v341 = v318
	goto L4
L85:
	;
	v359 = int32(268435456)
	v360 = int32(16)
	goto L3
L86:
	;
	if v40&int32(255) == int32(102) {
		v364 = v341
		v365 = int32(0)
		goto L2
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v347 = int32(268435456)
	if l5 == int32(0) {
		v359 = v347
		v360 = v341
		goto L3
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	if l4 != int32(2) {
		v359 = v347
		v360 = v341
		goto L3
	} else {
		goto L91
	}
L91:
	;
	if v40&int32(255) == int32(112) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v358 = int32(134217728)
	goto L94
L93:
	;
	v358 = int32(805306368)
	goto L94
L94:
	;
	v359 = v358
	v360 = v341
	goto L3
L95:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v398 <= int32(0) {
		goto L1
	} else {
		goto L114
	}
L96:
	;
	v396 = l6
	v397 = l7
	goto L95
L97:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+119)))
	if v369 != int32(112) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v372 = F_ExecGetTriggerOldSlot(m, l0, l1)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L18
	} else {
		goto L99
	}
L99:
	;
	v374 = F_ExecGetChildToRootMap(m, l2)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L18
	} else {
		goto L101
	}
L100:
	;
	v384 = F_ExecGetTriggerNewSlot(m, l0, l1)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L107
	}
L101:
	;
	if v374 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	v377 = F_execute_attr_map_slot(m, v376, l6, v372)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L18
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+32))
	m.T0[v380].(func(*base.Module, int32, int32))(m, v372, l6)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L18
	} else {
		goto L106
	}
L105:
	;
	v383 = v377
	goto L100
L106:
	;
	v383 = v372
	goto L100
L107:
	;
	v386 = F_ExecGetChildToRootMap(m, l3)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L18
	} else {
		goto L108
	}
L108:
	;
	if v386 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v386)+8))
	v389 = F_execute_attr_map_slot(m, v388, l7, v384)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L18
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	m.T0[v392].(func(*base.Module, int32, int32))(m, v384, l7)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L18
	} else {
		goto L113
	}
L112:
	;
	v396 = v383
	v397 = v389
	goto L95
L113:
	;
	v396 = v383
	v397 = v384
	goto L95
L114:
	;
	if l5 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v403 = int32(4)
	goto L117
L116:
	;
	v403 = int32(0)
	goto L117
L117:
	;
	v407 = int32(1)
	v414 = int32(0)
	v419 = v414
	v430 = v414
	goto L118
L118:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v447 = v444 + v430*int32(60)
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447)+12)))
	if (v364|int32(67))&v448 != l5|v364 {
		v895 = v419
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v895 == int32(0) {
		goto L1
	} else {
		goto L221
	}
L120:
	;
	v921 = v430 + int32(1)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v921 < v922 {
		v419 = v895
		v430 = v921
		goto L118
	} else {
		goto L220
	}
L121:
	;
	v451 = F_TriggerEnabled(m, l0, l1, v447, l4, l9, v396, v397)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	if v451 == int32(0) {
		v895 = v419
		goto L120
	} else {
		goto L123
	}
L123:
	;
	if v365 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if v419 != 0 {
		v497 = v419
		v498 = int32(0)
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v502 = v419
	goto L126
L126:
	;
	if base.Ui32(v407) < base.Ui32(l4-v407) {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v498
	v502 = v497
	goto L126
L128:
	;
	v458 = int32(536870912)
	v460 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v462 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460+v462*int32(20))+12))
	if v466 != 0 {
		v497 = v466
		v498 = v458
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v467 = int32(4425280)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v471 = *(*int32)(unsafe.Add(mBase, _consts[319]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v471
	v473 = int32(4425332)
	v474 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v477 = *(*int32)(unsafe.Add(mBase, _consts[320]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v477
	v479 = int32(0)
	v482 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v483 = F_tuplestore_begin_heap(m, v479, v479, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v468
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v474
	v490 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v492 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	*(*int32)(unsafe.Add(mBase, uint32(v490+v492*int32(20))+12)) = v483
	v497 = v483
	v498 = v458
	goto L127
L131:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	if v793 == int32(1250) {
		goto L185
	} else {
		goto L186
	}
L132:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v509 = v506 - int32(1644)
	if base.Ui32(v509) <= base.Ui32(int32(11)) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	if l5 == int32(0) {
		goto L131
	} else {
		goto L183
	}
L134:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+119)))
	if v619 == int32(112) {
		v895 = v502
		goto L120
	} else {
		goto L163
	}
L135:
	;
	if base.B2i32(l4 == v407)&l11 != 0 {
		goto L140
	} else {
		goto L141
	}
L136:
	;
	switch v517 {
	case 0:
		goto L133
	case 1:
		goto L135
	case 2:
		goto L134
	default:
		goto L131
	}
L137:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v509<<(uint(int32(2))%32))+uint32(_consts[321])))
	v517 = v516
	goto L139
L138:
	;
	v517 = int32(0)
	goto L139
L139:
	;
	goto L136
L140:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+16)))
	if v518 != 0 {
		v895 = v502
		goto L120
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v520 = F_ri_FetchConstraintInfo(m, v447, v38, int32(1))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L18
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v520)+168))
	if v522 <= int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v617 = int32(0)
	goto L147
L146:
	;
	v541 = int32(0)
	v549 = int32(1)
	v551 = v522
	goto L148
L147:
	;
	if v617 != 0 {
		goto L131
	} else {
		goto L162
	}
L148:
	;
	v561 = int32(*(*int16)(unsafe.Add(mBase, uint32(v520+int32(172)+v541<<(uint(int32(1))%32)))))
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v396)+6)))
	if v562 < v561 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v579 = int32(0)
	if v575&int32(1) == v579 {
		v588 = v579
		goto L155
	} else {
		goto L156
	}
L150:
	;
	F_slot_getsomeattrs_int(m, v396, v561)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L18
	} else {
		goto L153
	}
L151:
	;
	v567 = v551
	goto L152
L152:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	v570 = int32(1)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v561-v570))))
	v575 = (v572 ^ v570) & v549
	v577 = v541 + v570
	if v577 < v567 {
		v541 = v577
		v549 = v575
		v551 = v567
		goto L148
	} else {
		goto L154
	}
L153:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v520)+168))
	v567 = v566
	goto L152
L154:
	;
	goto L149
L155:
	;
	v617 = v588
	goto L147
L156:
	;
	if v397 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v585 = F_ri_KeysEqual(m, v38, v396, v397, v520, int32(1))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L18
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v588 = int32(1)
	goto L155
L160:
	;
	if v585 != 0 {
		v588 = v579
		goto L155
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v895 = v502
	goto L120
L163:
	;
	v622 = int32(0)
	v624 = F_ri_FetchConstraintInfo(m, v447, v38, v622)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L18
	} else {
		goto L166
	}
L164:
	;
	if v758 != 0 {
		goto L131
	} else {
		goto L182
	}
L165:
	;
	v758 = v719
	goto L164
L166:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v624)+168))
	if v626 <= int32(0) {
		v719 = v622
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v631 = int32(1)
	v650 = v622
	v654 = v631
	v655 = v631
	v656 = v626
	goto L168
L168:
	;
	v664 = int32(*(*int16)(unsafe.Add(mBase, uint32(v624+int32(236)+v650<<(uint(int32(1))%32)))))
	v665 = int32(*(*int16)(unsafe.Add(mBase, uint32(v397)+6)))
	if v665 < v664 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v683 = int32(0)
	if v676&int32(1) != 0 {
		v719 = v683
		goto L165
	} else {
		goto L175
	}
L170:
	;
	F_slot_getsomeattrs_int(m, v397, v664)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L18
	} else {
		goto L173
	}
L171:
	;
	v670 = v656
	goto L172
L172:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v397)+20))
	v673 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+v664-v673))))
	v676 = v675 & v655
	v679 = (v675 ^ v673) & v654
	v681 = v650 + v673
	if v681 < v670 {
		v650 = v681
		v654 = v679
		v655 = v676
		v656 = v670
		goto L168
	} else {
		goto L174
	}
L173:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v624)+168))
	v670 = v669
	goto L172
L174:
	;
	goto L169
L175:
	;
	if v679&int32(1) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+24))
	v695 = m.T0[v694].(func(*base.Module, int32) int32)(m, v396)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L179
	}
L177:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+164)))
	switch v688 - int32(102) {
	case 0:
		goto L178
	default:
		goto L176
	case 13:
		v719 = v683
		goto L165
	}
L178:
	;
	v758 = int32(1)
	goto L164
L179:
	;
	if v695 != 0 {
		v719 = int32(1)
		goto L165
	} else {
		goto L180
	}
L180:
	;
	v698 = F_ri_KeysEqual(m, v38, v396, v397, v624, int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L18
	} else {
		goto L181
	}
L181:
	;
	v719 = v698 ^ int32(1)
	goto L165
L182:
	;
	v895 = v502
	goto L120
L183:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+119)))
	if v762 == int32(112) {
		v895 = v502
		goto L120
	} else {
		goto L184
	}
L184:
	;
	goto L131
L185:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v447)+24))
	v797 = int32(0)
	if l8 == v797 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+32)))
	if v840 != 0 {
		goto L202
	} else {
		goto L203
	}
L188:
	;
	if v835 == int32(0) {
		v895 = v502
		goto L120
	} else {
		goto L201
	}
L189:
	;
	v835 = int32(0)
	goto L188
L190:
	;
	goto L191
L191:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v803 <= int32(0) {
		v828 = v797
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v835 = v828
	goto L188
L193:
	;
	v806 = int32(0)
	if v806 < v803 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v809 = v803
	goto L196
L195:
	;
	v809 = v806
	goto L196
L196:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v812 = int32(0)
	goto L197
L197:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v810+v812<<(uint(int32(2))%32))))
	v821 = base.B2i32(v820 == v796)
	if v820 == v796 {
		v828 = v821
		goto L192
	} else {
		goto L199
	}
L198:
	;
	v828 = v821
	goto L192
L199:
	;
	v823 = v812 + int32(1)
	if v823 != v809 {
		v812 = v823
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	goto L187
L202:
	;
	v841 = int32(32)
	goto L204
L203:
	;
	v841 = int32(0)
	goto L204
L204:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+33)))
	if v845 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v846 = int32(64)
	goto L207
L206:
	;
	v846 = int32(0)
	goto L207
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v403 | v841 | v846 | l4
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v850
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v852
	v855 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v856 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v856
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v855
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v447)+52))
	if v859 == v856 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v876
	v880 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v882 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	F_afterTriggerAddEvent(m, v880+v882*int32(20), v29+int32(-24), v29+int32(-52))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L18
	} else {
		goto L219
	}
L209:
	;
	switch l4 - int32(1) {
	case 0:
		goto L216
	case 1:
		goto L217
	case 2:
		v876 = int32(0)
		goto L208
	default:
		goto L218
	}
L210:
	;
	v862 = int32(0)
	if l10 == v862 {
		v876 = v862
		goto L208
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	if l10 != 0 {
		goto L209
	} else {
		goto L215
	}
L213:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v447)+56))
	if v865 == int32(0) {
		v876 = v862
		goto L208
	} else {
		goto L214
	}
L214:
	;
	goto L209
L215:
	;
	v876 = int32(0)
	goto L208
L216:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v876 = v875
	goto L208
L217:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v876 = v874
	goto L208
L218:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v876 = v873
	goto L208
L219:
	;
	v895 = v502
	goto L120
L220:
	;
	goto L119
L221:
	;
	if v396 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_tuplestore_puttupleslot(m, v895, v396)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L18
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v397 == int32(0) {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	F_tuplestore_puttupleslot(m, v895, v397)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L18
	} else {
		goto L227
	}
L227:
	;
	goto L1
}
func F_afterTriggerInvokeEvents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int64
	_ = v218
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v762 int32
	_ = v762
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	v5 = int32(0)
	v36 = m.G0
	v38 = v36 + int32(-64)
	m.G0 = v38
	if l2 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = F_CreateExecutorState(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = l2
	goto L3
L3:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v53 = F_AllocSetContextCreateInternal(m, v48, int32(59255), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v46 = v42
	goto L3
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L202
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L199
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L196
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L193
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L190
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L187
	}
L13:
	;
	F_MemoryContextDelete(m, v53)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L4
	} else {
		goto L180
	}
L14:
	;
	v762 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v78 = v5
	v79 = v55
	v80 = v5
	v83 = v5
	v85 = v5
	v88 = int32(1)
	v89 = v5
	v94 = v5
	v101 = v5
	goto L17
L17:
	;
	v103 = int32(1)
	v105 = v79 + int32(16)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if base.Ui32(v105) < base.Ui32(v106) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v702 == int32(0) {
		v762 = v710
		goto L13
	} else {
		goto L177
	}
L19:
	;
	v115 = v105
	v118 = v78
	v120 = v80
	v123 = v83
	v125 = v85
	v128 = v88
	v129 = v89
	v132 = v103
	v134 = v94
	v141 = v101
	goto L22
L20:
	;
	v700 = v78
	v702 = v80
	v705 = v83
	v707 = v85
	v710 = v88
	v711 = v89
	v714 = v103
	v716 = v94
	v723 = v101
	goto L21
L21:
	;
	if l3 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L22:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v143&int32(1073741824) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v700 = v648
	v702 = v650
	v705 = v653
	v707 = v655
	v710 = v658
	v711 = v659
	v714 = v662
	v716 = v664
	v723 = v671
	goto L21
L24:
	;
	v675 = v646 & int32(939524096)
	if v675 == int32(134217728) {
		v686 = int32(24)
		goto L165
	} else {
		goto L166
	}
L25:
	;
	v635 = base.B2i32(v143 < int32(0))
	v646 = v143
	v648 = v118
	v650 = v120
	v653 = v123
	v655 = v125
	v658 = v635 & v128
	v659 = v129
	v662 = v635 & v132
	v664 = v134
	v671 = v141
	goto L24
L26:
	;
	v150 = v115 + v143&int32(134217727)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v151 != l1 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v125 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v194&int32(939524096) == int32(134217728) {
		goto L43
	} else {
		goto L44
	}
L29:
	;
	v161 = F_ExecGetTriggerResultRel(m, v46, v159, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v159 = v155
	goto L29
L31:
	;
	goto L32
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v125)+56))
	if v156 == v157 {
		v187 = v118
		v188 = v120
		v189 = v123
		v190 = v125
		v191 = v129
		v192 = v134
		v193 = v141
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v159 = v156
	goto L29
L34:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)+64))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+56))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+52))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	if v120 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_ExecDropSingleTupleTableSlot(m, v120)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v172 = v123
	goto L37
L37:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+119)))
	if v175 != int32(102) {
		v187 = v161
		v188 = int32(0)
		v189 = v172
		v190 = v166
		v191 = v163
		v192 = v165
		v193 = v164
		goto L28
	} else {
		goto L40
	}
L38:
	;
	F_ExecDropSingleTupleTableSlot(m, v123)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v172 = int32(0)
	goto L37
L40:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v166)+52))
	v180 = F_MakeSingleTupleTableSlot(m, v178, int32(1567052))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v166)+52))
	v184 = F_MakeSingleTupleTableSlot(m, v182, int32(1567052))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v187 = v161
	v188 = v180
	v189 = v184
	v190 = v166
	v191 = v163
	v192 = v165
	v193 = v164
	goto L28
L43:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v200 = F_ExecGetTriggerResultRel(m, v46, v199, v187)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	v206 = v194
	v207 = v187
	v208 = v187
	goto L45
L45:
	;
	v211 = v115 + v206&int32(134217727)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(-8)))) = v216
	v218 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36+int32(-16)))) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v36+int32(-24)))) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v36+int32(-40)))) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v218
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)) = uint8(v216)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+6)) = uint8(v216)
	if v192 == v216 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	v203 = F_ExecGetTriggerResultRel(m, v46, v202, v187)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v206 = v205
	v207 = v200
	v208 = v203
	goto L45
L48:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v631 = v627&int32(1073741823) | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v631
	v646 = v631
	v648 = v187
	v650 = v188
	v653 = v189
	v655 = v190
	v658 = v128
	v659 = v191
	v662 = v132
	v664 = v192
	v671 = v193
	goto L24
L49:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v234 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v247 = int32(0)
	goto L51
L51:
	;
	v276 = v237 + v247*int32(60)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v212 != v277 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v276
	if v191 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v280 = v247 + int32(1)
	if v280 != v234 {
		v247 = v280
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L48
L57:
	;
	F_InstrStartNode(m, v191+v247*int32(416))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	v289 = v206
	goto L59
L59:
	;
	v291 = v289 & int32(939524096)
	if v291 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v289 = v288
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = int64(0)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	if v487 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L62:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+8)))
	if v378 != 0 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v188
	v363 = F_ExecFetchSlotHeapTuple(m, v188, int32(1), v36+int32(-57))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L75
	}
L64:
	;
	if v291 != int32(536870912) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v299 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v297+v299*int32(20))+12))
	if v303 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v306 = int32(4425280)
	v307 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v310 = *(*int32)(unsafe.Add(mBase, _consts[319]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v310
	v312 = int32(4425332)
	v313 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v316 = *(*int32)(unsafe.Add(mBase, _consts[320]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v316
	v318 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v322 = F_tuplestore_begin_heap(m, v318, v318, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	v336 = v303
	goto L68
L68:
	;
	v341 = F_tuplestore_gettupleslot(m, v336, int32(1), int32(0), v188)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v307
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v313
	v329 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v331 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	*(*int32)(unsafe.Add(mBase, uint32(v329+v331*int32(20))+12)) = v322
	v336 = v322
	goto L68
L70:
	;
	if v341 == int32(0) {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v345&int32(3) != int32(2) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v352 = F_tuplestore_gettupleslot(m, v336, int32(1), int32(0), v189)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	if v352 == int32(0) {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	goto L63
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v363
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v367&int32(3) != int32(2) {
		v480 = int32(0)
		goto L61
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+44)) = v189
	v376 = F_ExecFetchSlotHeapTuple(m, v189, int32(1), v36+int32(-58))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v480 = v376
	goto L61
L78:
	;
	v429 = int32(0)
	if base.B2i32(v428&int32(134217728) == v429)&base.B2i32(v426 != int32(805306368)) != 0 {
		v480 = v429
		goto L61
	} else {
		goto L101
	}
L79:
	;
	v379 = F_ExecGetTriggerOldSlot(m, v46, v207)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = int32(0)
	v426 = v291
	v428 = v289
	goto L78
L82:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v382 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v384&int32(1) == int32(0) {
		goto L10
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v214)+188))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+60))
	v394 = m.T0[v393].(func(*base.Module, int32, int32, int32, int32) int32)(m, v214, v115+int32(4), int32(4099040), v379)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	if v394 == int32(0) {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	if v187 != v207 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v417 = F_ExecFetchSlotHeapTuple(m, v412, int32(0), v36+int32(-57))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L100
	}
L90:
	;
	v399 = F_ExecGetChildToRootMap(m, v207)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v379
	v412 = v379
	goto L89
L93:
	;
	v401 = F_ExecGetTriggerOldSlot(m, v46, v187)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v401
	if v399 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	v405 = F_execute_attr_map_slot(m, v404, v379, v401)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+32))
	m.T0[v408].(func(*base.Module, int32, int32))(m, v401, v379)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	v412 = v401
	goto L89
L99:
	;
	v412 = v401
	goto L89
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v417
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v426 = v420 & int32(939524096)
	v428 = v420
	goto L78
L101:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+14)))
	if v437 == int32(0) {
		v480 = v429
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v440 = F_ExecGetTriggerNewSlot(m, v46, v208)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v443 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v445&int32(1) == int32(0) {
		goto L8
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v213)+188))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+60))
	v455 = m.T0[v454].(func(*base.Module, int32, int32, int32, int32) int32)(m, v213, v115+int32(10), int32(4099040), v440)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	if v455 == int32(0) {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	if v187 != v208 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v478 = F_ExecFetchSlotHeapTuple(m, v473, int32(0), v36+int32(-58))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L121
	}
L111:
	;
	v460 = F_ExecGetChildToRootMap(m, v208)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+44)) = v440
	v473 = v440
	goto L110
L114:
	;
	v462 = F_ExecGetTriggerNewSlot(m, v46, v187)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+44)) = v462
	if v460 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v460)+8))
	v466 = F_execute_attr_map_slot(m, v465, v440, v462)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+32))
	m.T0[v469].(func(*base.Module, int32, int32))(m, v462, v440)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L120
	}
L119:
	;
	v473 = v462
	goto L110
L120:
	;
	v473 = v462
	goto L110
L121:
	;
	v480 = v478
	goto L61
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = int32(442)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v506 & int32(7)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+12)))
	if v511&int32(16) != 0 {
		goto L128
	} else {
		goto L129
	}
L123:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	if v490 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v491
	v493 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v487)+8)) = uint8(v493)
	goto L126
L125:
	;
	goto L126
L126:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v276)+56))
	if v495 == int32(0) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v499
	v501 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+8)) = uint8(v501)
	goto L122
L128:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v514
	goto L130
L129:
	;
	goto L130
L130:
	;
	F_MemoryContextReset(m, v53)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(-52)))) = v523
	v526 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(-56)))) = v526
	goto L132
L132:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v528 != v529 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v531 | int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v528
	goto L136
L134:
	;
	goto L135
L135:
	;
	v541 = F_ExecCallTriggerFunc(m, v36+int32(-48), v247, v193, int32(0), v53)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if v551 != v552 {
		goto L143
	} else {
		goto L144
	}
L138:
	;
	if v541 == int32(0) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	if v541 == v545 {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v541 == v547 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	F_pfree(m, v541)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	goto L137
L143:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v554
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v551
	goto L146
L144:
	;
	goto L145
L145:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
	if v559 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	F_pfree(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+6)))
	if v565 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	F_pfree(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v188 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L153
L155:
	;
	if v191 == int32(0) {
		goto L48
	} else {
		goto L163
	}
L156:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v571 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	m.T0[v573].(func(*base.Module, int32))(m, v571)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v38)+44))
	if v576 == int32(0) {
		goto L155
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v576)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+12))
	m.T0[v580].(func(*base.Module, int32))(m, v576)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	goto L155
L163:
	;
	F_InstrStopNode(m, v191+v247*int32(416), float64(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	goto L48
L165:
	;
	v687 = v686 + v115
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if base.Ui32(v687) < base.Ui32(v688) {
		v115 = v687
		v118 = v648
		v120 = v650
		v123 = v653
		v125 = v655
		v128 = v658
		v129 = v659
		v132 = v662
		v134 = v664
		v141 = v671
		goto L22
	} else {
		goto L171
	}
L166:
	;
	if v675 == int32(268435456) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v682 = int32(12)
	goto L169
L168:
	;
	v682 = int32(4)
	goto L169
L169:
	;
	if v675 != int32(805306368) {
		v686 = v682
		goto L165
	} else {
		goto L170
	}
L170:
	;
	v686 = int32(16)
	goto L165
L171:
	;
	goto L23
L172:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v735 != 0 {
		v78 = v700
		v79 = v735
		v80 = v702
		v83 = v705
		v85 = v707
		v88 = v710
		v89 = v711
		v94 = v716
		v101 = v723
		goto L17
	} else {
		goto L176
	}
L173:
	;
	if v714 == int32(0) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v105
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v730
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v79 != v732 {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
	goto L172
L176:
	;
	goto L18
L177:
	;
	F_ExecDropSingleTupleTableSlot(m, v702)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	F_ExecDropSingleTupleTableSlot(m, v705)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v762 = v710
	goto L13
L180:
	;
	if l2 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_ExecCloseResultRelations(m, v46)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v38 - int32(-64)
	return v762 & int32(1)
L184:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	F_ExecResetTupleTable(m, v783, int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	F_FreeExecutorState(m, v46)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	F_errmsg_internal(m, int32(211143), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(465125), int32(4388), int32(327283))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errmsg_internal(m, int32(211102), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(465125), int32(4394), int32(327283))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errmsg_internal(m, int32(315356), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(306339), int32(1264), int32(255277))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(211143), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(465125), int32(4435), int32(327283))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(315356), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(306339), int32(1264), int32(255277))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errmsg_internal(m, int32(211102), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(465125), int32(4477), int32(327283))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
