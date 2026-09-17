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
	var v72 int32
	_ = v72
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
	var v160 int32
	_ = v160
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[0]))
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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[2]))
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
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
	v20 = v19
	goto L8
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[3]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
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
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[4]))
	v38 = v28
	v39 = v37
	v40 = v30
	goto L15
L13:
	;
	v72 = v33
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v74 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if v38 < v39 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v72 = v65
	goto L14
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[5]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[3])) = v63
	v65 = v60 + v32
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
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[4]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[3]))
	v59 = v56
	v60 = v54
	v61 = v58
	goto L19
L21:
	;
	goto L16
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
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
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v72)+4))
	*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[6])) = v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[7])) = v97
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[8]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[6]))
	if v84 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[8])) = int64(0)
	goto L22
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[6])) = v86
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
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[8]))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[7]))
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
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[8]))
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
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[9]))
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
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[6]))
	if v152 == v149 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[9])) = v138
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
	v147 = v146
	goto L41
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndSubXact[1]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156+v136)+20))
	v160 = v152
	goto L44
L44:
	;
	v166 = v160 + int32(16)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
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
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v209 != 0 {
		v160 = v209
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
		v199 = int32(24)
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
	v200 = v199 + v169
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if base.Ui32(v200) < base.Ui32(v201) {
		v169 = v200
		goto L49
	} else {
		goto L60
	}
L55:
	;
	if v189 != int32(805306368) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v189 == int32(268435456) {
		v199 = int32(12)
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v199 = int32(16)
	goto L54
L59:
	;
	v199 = int32(4)
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
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
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
	var v381 int32
	_ = v381
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
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
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
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
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
	var v709 int32
	_ = v709
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
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	v29 = m.G0
	v31 = v29 + int32(-64)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[0]))
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
		goto L95
	} else {
		goto L96
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v361
	v365 = v360
	v366 = int32(0)
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
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[1]))
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
	v301 = m.ExcPending
	if v301 != 0 {
		goto L18
	} else {
		goto L78
	}
L9:
	;
	v148 = int32(0)
	if base.B2i32(l5 == v148)|base.B2i32(l10 == v148) == v148 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[1])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2])) = v70
	if v68 <= v42 {
		goto L9
	} else {
		goto L24
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[3]))
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
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2]))
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
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2]))
	v108 = v105 + v90*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v111
	v116 = v90 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[1]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_AfterTriggerSaveEvent_0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v344 = int32(32)
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
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+19)))
	if v266 == int32(0) {
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
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v241 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L33:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if l6 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
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
L36:
	;
	if l7 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v158&int32(2) != 0 {
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
	F_TransitionTableAddTuple(m, l4, l10, l1, l6, int32(0), v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L18
	} else {
		goto L46
	}
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179+l10)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v183 = v182
	goto L39
L41:
	;
	v169 = int32(0)
	if l4 != int32(2) {
		v183 = v169
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10))))
	if v163&int32(1) == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v179 = int32(16)
	goto L40
L44:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+1)))
	if v172&int32(1) == int32(0) {
		v183 = v169
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v179 = int32(12)
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
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7)+4)))
	if v190&int32(2) != 0 {
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
	F_TransitionTableAddTuple(m, l4, l10, l1, l7, v155, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L18
	} else {
		goto L57
	}
L51:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209+l10)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	v213 = v212
	goto L50
L52:
	;
	v199 = int32(0)
	if l4 != int32(2) {
		v213 = v199
		goto L50
	} else {
		goto L55
	}
L53:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+3)))
	if v193&int32(1) == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v209 = int32(8)
	goto L51
L55:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+2)))
	if v202&int32(1) == int32(0) {
		v213 = v199
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v209 = int32(12)
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
	v221 = int32(1)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+14)))
	if v222 != v221 {
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
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
	v226 = int32(1)
	v230 = int32(base.Ui32(v225)>>(uint(v226)%32)) & v226
	goto L63
L62:
	;
	v230 = v221
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
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)))
	v232 = int32(1)
	v237 = int32(base.Ui32(v231)>>(uint(v232)%32)) & v232
	goto L66
L65:
	;
	v237 = int32(1)
	goto L66
L66:
	;
	if v237 != v230 {
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
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v246)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v248
	v250 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v250)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = int32(-1)
	v344 = int32(4)
	goto L4
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_AfterTriggerSaveEvent_0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	F_cancel_prior_stmt_triggers(m, v259, int32(3), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	v360 = int32(4)
	v361 = int32(268435456)
	goto L3
L73:
	;
	goto L29
L74:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v271)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v273
	v275 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v275)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = int32(-1)
	v344 = int32(8)
	goto L4
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_AfterTriggerSaveEvent_0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	F_cancel_prior_stmt_triggers(m, v284, int32(4), int32(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	v360 = int32(8)
	v361 = int32(268435456)
	goto L3
L78:
	;
	F_errmsg_internal(m, int32(_a_F_AfterTriggerSaveEvent_1), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_AfterTriggerSaveEvent_2), int32(_a_F_AfterTriggerSaveEvent_3), int32(_a_F_AfterTriggerSaveEvent_4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v315
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+54)) = uint16(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+50)) = v319
	v321 = int32(16)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+119)))
	if v323 != int32(112) {
		v344 = v321
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_AfterTriggerSaveEvent_0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+44)) = int64(-281470681743361)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	v337 = int32(2)
	F_cancel_prior_stmt_triggers(m, v336, v337, v337)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L18
	} else {
		goto L85
	}
L84:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v330
	v344 = v321
	goto L4
L85:
	;
	v360 = int32(16)
	v361 = int32(268435456)
	goto L3
L86:
	;
	if v40 == int32(102) {
		v365 = v344
		v366 = int32(1)
		goto L2
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(l4 != int32(2)) != 0 {
		v360 = v344
		v361 = int32(268435456)
		goto L3
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	if v40 == int32(112) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v358 = int32(134217728)
	goto L93
L92:
	;
	v358 = int32(805306368)
	goto L93
L93:
	;
	v360 = v344
	v361 = v358
	goto L3
L94:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v399 <= int32(0) {
		goto L1
	} else {
		goto L113
	}
L95:
	;
	v396 = l7
	v398 = l6
	goto L94
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+119)))
	if v370 != int32(112) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v373 = F_ExecGetTriggerOldSlot(m, l0, l1)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L18
	} else {
		goto L98
	}
L98:
	;
	v375 = F_ExecGetChildToRootMap(m, l2)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L100
	}
L99:
	;
	v385 = F_ExecGetTriggerNewSlot(m, l0, l1)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L106
	}
L100:
	;
	if v375 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v378 = F_execute_attr_map_slot(m, v377, l6, v373)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L18
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+32))
	m.T0[v381].(func(*base.Module, int32, int32))(m, v373, l6)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L18
	} else {
		goto L105
	}
L104:
	;
	v384 = v378
	goto L99
L105:
	;
	v384 = v373
	goto L99
L106:
	;
	v387 = F_ExecGetChildToRootMap(m, l3)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L18
	} else {
		goto L107
	}
L107:
	;
	if v387 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	v390 = F_execute_attr_map_slot(m, v389, l7, v385)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L18
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v385)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+32))
	m.T0[v393].(func(*base.Module, int32, int32))(m, v385, l7)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L18
	} else {
		goto L112
	}
L111:
	;
	v396 = v390
	v398 = v384
	goto L94
L112:
	;
	v396 = v385
	v398 = v384
	goto L94
L113:
	;
	if l5 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v404 = int32(4)
	goto L116
L115:
	;
	v404 = int32(0)
	goto L116
L116:
	;
	v408 = int32(1)
	v415 = int32(0)
	v420 = v415
	v431 = v415
	goto L117
L117:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v448 = v445 + v431*int32(60)
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v448)+12)))
	if (v365|int32(67))&v449 != l5|v365 {
		v897 = v420
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v897 == int32(0) {
		goto L1
	} else {
		goto L219
	}
L119:
	;
	v923 = v431 + int32(1)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v923 < v924 {
		v420 = v897
		v431 = v923
		goto L117
	} else {
		goto L218
	}
L120:
	;
	v452 = F_TriggerEnabled(m, l0, l1, v448, l4, l9, v398, v396)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	if v452 == int32(0) {
		v897 = v420
		goto L119
	} else {
		goto L122
	}
L122:
	;
	if v366 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if v420 != 0 {
		v496 = v420
		v498 = int32(0)
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v501 = v420
	goto L125
L125:
	;
	if base.Ui32(v408) < base.Ui32(l4-v408) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v498
	v501 = v496
	goto L125
L127:
	;
	v457 = int32(536870912)
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2]))
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[0]))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v459+v461*int32(20))+12))
	if v465 != 0 {
		v496 = v465
		v498 = v457
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v466 = int32(_a_F_AfterTriggerSaveEvent_5)
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[4]))
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[4])) = v470
	v472 = int32(_a_F_AfterTriggerSaveEvent_6)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[6]))
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[6])) = v476
	v478 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[8]))
	v482 = F_tuplestore_begin_heap(m, v478, v478, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L18
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[4])) = v467
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[6])) = v473
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2]))
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v489+v491*int32(20))+12)) = v482
	v496 = v482
	v498 = v457
	goto L126
L130:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v793 == int32(1250) {
		goto L189
	} else {
		goto L190
	}
L131:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v507 = v505 - int32(1644)
	if base.Ui32(v507) <= base.Ui32(int32(11)) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	if l5 == int32(0) {
		goto L130
	} else {
		goto L187
	}
L133:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+119)))
	if v616 == int32(112) {
		v897 = v501
		goto L119
	} else {
		goto L162
	}
L134:
	;
	if base.B2i32(l4 == v408)&l11 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	switch v514 {
	case 0:
		goto L132
	case 1:
		goto L134
	case 2:
		goto L133
	default:
		goto L130
	}
L136:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507<<(uint(int32(2))%32))+uint32(_c_F_AfterTriggerSaveEvent[9])))
	v514 = v512
	goto L138
L137:
	;
	v514 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+16)))
	if v515 != 0 {
		v897 = v501
		goto L119
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v517 = F_ri_FetchConstraintInfo(m, v448, v38, int32(1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L18
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v517)+168))
	if v519 <= int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v614 = int32(0)
	goto L146
L145:
	;
	v538 = int32(0)
	v542 = v519
	v544 = int32(1)
	goto L147
L146:
	;
	if v614 != 0 {
		goto L130
	} else {
		goto L161
	}
L147:
	;
	v558 = int32(*(*int16)(unsafe.Add(mBase, uint32(v517+int32(172)+v538<<(uint(int32(1))%32)))))
	v559 = int32(*(*int16)(unsafe.Add(mBase, uint32(v398)+6)))
	if v559 < v558 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v576 = int32(0)
	if v572&int32(1) == v576 {
		v585 = v576
		goto L154
	} else {
		goto L155
	}
L149:
	;
	F_slot_getsomeattrs_int(m, v398, v558)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L18
	} else {
		goto L152
	}
L150:
	;
	v564 = v542
	goto L151
L151:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v398)+20))
	v567 = int32(1)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v558-v567))))
	v572 = (v569 ^ v567) & v544
	v574 = v538 + v567
	if v574 < v564 {
		v538 = v574
		v542 = v564
		v544 = v572
		goto L147
	} else {
		goto L153
	}
L152:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v517)+168))
	v564 = v563
	goto L151
L153:
	;
	goto L148
L154:
	;
	v614 = v585
	goto L146
L155:
	;
	if v396 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v582 = F_ri_KeysEqual(m, v38, v398, v396, v517, int32(1))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L18
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v585 = int32(1)
	goto L154
L159:
	;
	if v582 != 0 {
		v585 = v576
		goto L154
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v897 = v501
	goto L119
L162:
	;
	v619 = int32(0)
	v621 = F_ri_FetchConstraintInfo(m, v448, v38, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L18
	} else {
		goto L165
	}
L163:
	;
	if v758 != 0 {
		goto L130
	} else {
		goto L186
	}
L164:
	;
	v758 = v709
	goto L163
L165:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v621)+168))
	if v623 <= int32(0) {
		v709 = v619
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v628 = int32(1)
	v637 = v619
	v645 = v628
	v647 = v628
	v650 = v623
	goto L167
L167:
	;
	v661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621+int32(236)+v637<<(uint(int32(1))%32)))))
	v662 = int32(*(*int16)(unsafe.Add(mBase, uint32(v396)+6)))
	if v662 < v661 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v680 = int32(0)
	if v673&int32(1) != 0 {
		v709 = v680
		goto L164
	} else {
		goto L174
	}
L169:
	;
	F_slot_getsomeattrs_int(m, v396, v661)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L18
	} else {
		goto L172
	}
L170:
	;
	v667 = v650
	goto L171
L171:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	v670 = int32(1)
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668+v661-v670))))
	v673 = v672 & v647
	v676 = (v672 ^ v670) & v645
	v678 = v637 + v670
	if v678 < v667 {
		v637 = v678
		v645 = v676
		v647 = v673
		v650 = v667
		goto L167
	} else {
		goto L173
	}
L172:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v621)+168))
	v667 = v666
	goto L171
L173:
	;
	goto L168
L174:
	;
	if v676&int32(1) != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+24))
	v695 = m.T0[v694].(func(*base.Module, int32) int32)(m, v398)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L183
	}
L176:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+164)))
	v687 = v685 - int32(102)
	if v687 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if v687 == int32(13) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L179
L179:
	;
	v758 = int32(1)
	goto L163
L180:
	;
	v709 = v680
	goto L164
L181:
	;
	goto L175
L183:
	;
	if v695 != 0 {
		v709 = int32(1)
		goto L164
	} else {
		goto L184
	}
L184:
	;
	v698 = F_ri_KeysEqual(m, v38, v398, v396, v621, int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L18
	} else {
		goto L185
	}
L185:
	;
	v709 = v698 ^ int32(1)
	goto L164
L186:
	;
	v897 = v501
	goto L119
L187:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+119)))
	if v762 == int32(112) {
		v897 = v501
		goto L119
	} else {
		goto L188
	}
L188:
	;
	goto L130
L189:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
	v797 = int32(0)
	if l8 == v797 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+33)))
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v838<<(uint(int32(6))%32)&int32(64) | (v404 | v843<<(uint(int32(5))%32)&int32(32)) | l4
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v852
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v854
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[10]))
	v858 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v858
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v857
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v448)+52))
	if v861 == v858 {
		goto L208
	} else {
		goto L209
	}
L192:
	;
	if v835 == int32(0) {
		v897 = v501
		goto L119
	} else {
		goto L205
	}
L193:
	;
	v835 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v803 <= int32(0) {
		v829 = v797
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v835 = v829
	goto L192
L197:
	;
	v806 = int32(0)
	if v806 < v803 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v809 = v803
	goto L200
L199:
	;
	v809 = v806
	goto L200
L200:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v812 = int32(0)
	goto L201
L201:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v810+v812<<(uint(int32(2))%32))))
	v821 = base.B2i32(v820 == v796)
	if v820 == v796 {
		v829 = v821
		goto L196
	} else {
		goto L203
	}
L202:
	;
	v829 = v821
	goto L196
L203:
	;
	v823 = v812 + int32(1)
	if v823 != v809 {
		v812 = v823
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	goto L191
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v878
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[2]))
	v884 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerSaveEvent[0]))
	F_afterTriggerAddEvent(m, v882+v884*int32(20), v29+int32(-24), v29+int32(-52))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L18
	} else {
		goto L217
	}
L207:
	;
	switch l4 - int32(1) {
	case 0:
		goto L214
	case 1:
		goto L215
	case 2:
		v878 = int32(0)
		goto L206
	default:
		goto L216
	}
L208:
	;
	v864 = int32(0)
	if l10 == v864 {
		v878 = v864
		goto L206
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if l10 != 0 {
		goto L207
	} else {
		goto L213
	}
L211:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v448)+56))
	if v867 == int32(0) {
		v878 = v864
		goto L206
	} else {
		goto L212
	}
L212:
	;
	goto L207
L213:
	;
	v878 = int32(0)
	goto L206
L214:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v878 = v877
	goto L206
L215:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v878 = v876
	goto L206
L216:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v878 = v875
	goto L206
L217:
	;
	v897 = v501
	goto L119
L218:
	;
	goto L118
L219:
	;
	if v398 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_tuplestore_puttupleslot(m, v897, v398)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L18
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if v396 == int32(0) {
		goto L1
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	F_tuplestore_puttupleslot(m, v897, v396)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	goto L1
}
func F_afterTriggerInvokeEvents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v567 int32
	_ = v567
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v725 int32
	_ = v725
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	v5 = int32(0)
	v32 = m.G0
	v34 = v32 + int32(-64)
	m.G0 = v34
	if l2 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = F_CreateExecutorState(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = l2
	goto L3
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[0]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_afterTriggerInvokeEvents_0), int32(0), int32(_a_F_afterTriggerInvokeEvents_1), int32(_a_F_afterTriggerInvokeEvents_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v42 = v38
	goto L3
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v51 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L4
	} else {
		goto L197
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L194
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L4
	} else {
		goto L191
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L188
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L185
	}
L12:
	;
	F_MemoryContextDelete(m, v49)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L178
	}
L13:
	;
	v725 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v67 = v5
	v68 = v51
	v69 = v5
	v71 = v5
	v73 = v5
	v75 = int32(1)
	v76 = v5
	v81 = v5
	v84 = v5
	goto L16
L16:
	;
	v87 = int32(1)
	v89 = v68 + int32(16)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if base.Ui32(v89) < base.Ui32(v90) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v670 == int32(0) {
		v725 = v676
		goto L12
	} else {
		goto L175
	}
L18:
	;
	v100 = v89
	v103 = v67
	v105 = v69
	v107 = v71
	v109 = v73
	v111 = v75
	v112 = v76
	v115 = v87
	v117 = v81
	v120 = v84
	goto L21
L19:
	;
	v668 = v67
	v670 = v69
	v672 = v71
	v674 = v73
	v676 = v75
	v677 = v76
	v680 = v87
	v682 = v81
	v685 = v84
	goto L20
L20:
	;
	v688 = int32(0)
	if base.B2i32(l3 == v688)|base.B2i32(v680 == v688) != 0 {
		goto L171
	} else {
		goto L172
	}
L21:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v123&int32(1073741824) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v668 = v621
	v670 = v623
	v672 = v625
	v674 = v627
	v676 = v629
	v677 = v630
	v680 = v633
	v682 = v635
	v685 = v638
	goto L20
L23:
	;
	v643 = v619 & int32(939524096)
	if v643 == int32(134217728) {
		v653 = int32(24)
		goto L164
	} else {
		goto L165
	}
L24:
	;
	v607 = base.B2i32(v123 < int32(0))
	v619 = v123
	v621 = v103
	v623 = v105
	v625 = v107
	v627 = v109
	v629 = v607 & v111
	v630 = v112
	v633 = v607 & v115
	v635 = v117
	v638 = v120
	goto L23
L25:
	;
	v130 = v100 + v123&int32(134217727)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v131 != l1 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if v109 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v174&int32(939524096) == int32(134217728) {
		goto L42
	} else {
		goto L43
	}
L28:
	;
	v141 = F_ExecGetTriggerResultRel(m, v42, v139, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L33
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v139 = v135
	goto L28
L30:
	;
	goto L31
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v109)+56))
	if v136 == v137 {
		v167 = v103
		v168 = v105
		v169 = v107
		v170 = v109
		v171 = v112
		v172 = v117
		v173 = v120
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v139 = v136
	goto L28
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+64))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+56))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+52))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v105 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ExecDropSingleTupleTableSlot(m, v105)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v152 = v107
	goto L36
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146)+48))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+119)))
	if v155 != int32(102) {
		v167 = v141
		v168 = int32(0)
		v169 = v152
		v170 = v146
		v171 = v143
		v172 = v145
		v173 = v144
		goto L27
	} else {
		goto L39
	}
L37:
	;
	F_ExecDropSingleTupleTableSlot(m, v107)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v152 = int32(0)
	goto L36
L39:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v146)+52))
	v160 = F_MakeTupleTableSlot(m, v158, int32(_a_F_afterTriggerInvokeEvents_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v146)+52))
	v164 = F_MakeTupleTableSlot(m, v162, int32(_a_F_afterTriggerInvokeEvents_3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v167 = v141
	v168 = v160
	v169 = v164
	v170 = v146
	v171 = v143
	v172 = v145
	v173 = v144
	goto L27
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v180 = F_ExecGetTriggerResultRel(m, v42, v179, v167)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	v186 = v167
	v187 = v174
	v188 = v167
	goto L44
L44:
	;
	v191 = v100 + v187&int32(134217727)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v196
	v198 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+48)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v34)+40)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v34)+32)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v198
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+7)) = uint8(v196)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+6)) = uint8(v196)
	if v172 == v196 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v183 = F_ExecGetTriggerResultRel(m, v42, v182, v167)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v186 = v180
	v187 = v185
	v188 = v183
	goto L44
L47:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v603 = v599&int32(1073741823) | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v603
	v619 = v603
	v621 = v167
	v623 = v168
	v625 = v169
	v627 = v170
	v629 = v111
	v630 = v171
	v633 = v115
	v635 = v172
	v638 = v173
	goto L23
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v214 <= int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v228 = int32(0)
	goto L50
L50:
	;
	v252 = v217 + v228*int32(60)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v192 != v253 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v252
	if v171 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v256 = v228 + int32(1)
	if v256 != v214 {
		v228 = v256
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L47
L56:
	;
	F_InstrStartNode(m, v171+v228*int32(416))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	v265 = v187
	goto L58
L58:
	;
	v267 = v265 & int32(939524096)
	if v267 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v265 = v264
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v34)+48)) = int64(0)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v463 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L61:
	;
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
	if v354 != 0 {
		goto L78
	} else {
		goto L79
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v168
	v339 = F_ExecFetchSlotHeapTuple(m, v168, int32(1), v32+int32(-57))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L74
	}
L63:
	;
	if v267 != int32(536870912) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[1]))
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[2]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+v275*int32(20))+12))
	if v279 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v282 = int32(_a_F_afterTriggerInvokeEvents_4)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[0]))
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[0])) = v286
	v288 = int32(_a_F_afterTriggerInvokeEvents_5)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[4]))
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[4])) = v292
	v294 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[6]))
	v298 = F_tuplestore_begin_heap(m, v294, v294, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v313 = v279
	goto L67
L67:
	;
	v317 = F_tuplestore_gettupleslot(m, v313, int32(1), int32(0), v168)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[0])) = v283
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[4])) = v289
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[1]))
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v305+v307*int32(20))+12)) = v298
	v313 = v298
	goto L67
L69:
	;
	if v317 == int32(0) {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v321&int32(3) != int32(2) {
		goto L62
	} else {
		goto L71
	}
L71:
	;
	v328 = F_tuplestore_gettupleslot(m, v313, int32(1), int32(0), v169)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	if v328 == int32(0) {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	goto L62
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v339
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v343&int32(3) != int32(2) {
		v457 = int32(0)
		goto L60
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v169
	v352 = F_ExecFetchSlotHeapTuple(m, v169, int32(1), v32+int32(-58))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v457 = v352
	goto L60
L77:
	;
	v405 = int32(0)
	if base.B2i32(v404&int32(134217728) == v405)&base.B2i32(v403 != int32(805306368)) != 0 {
		v457 = v405
		goto L60
	} else {
		goto L100
	}
L78:
	;
	v355 = F_ExecGetTriggerOldSlot(m, v42, v186)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = int32(0)
	v403 = v267
	v404 = v265
	goto L77
L81:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[7]))
	if v358 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[8])))
	if v360&int32(1) == int32(0) {
		goto L7
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v194)+188))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+60))
	v370 = m.T0[v369].(func(*base.Module, int32, int32, int32, int32) int32)(m, v194, v100+int32(4), int32(_a_F_afterTriggerInvokeEvents_6), v355)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	if v370 == int32(0) {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	if v186 != v167 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v393 = F_ExecFetchSlotHeapTuple(m, v389, int32(0), v32+int32(-57))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L99
	}
L89:
	;
	v375 = F_ExecGetChildToRootMap(m, v186)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v355
	v389 = v355
	goto L88
L92:
	;
	v377 = F_ExecGetTriggerOldSlot(m, v42, v167)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v377
	if v375 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v381 = F_execute_attr_map_slot(m, v380, v355, v377)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377)+8))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+32))
	m.T0[v384].(func(*base.Module, int32, int32))(m, v377, v355)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L98
	}
L97:
	;
	v389 = v377
	goto L88
L98:
	;
	v389 = v377
	goto L88
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v393
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v403 = v396 & int32(939524096)
	v404 = v396
	goto L77
L100:
	;
	v413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+14)))
	if v413 == int32(0) {
		v457 = v405
		goto L60
	} else {
		goto L101
	}
L101:
	;
	v416 = F_ExecGetTriggerNewSlot(m, v42, v188)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[7]))
	if v419 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[8])))
	if v421&int32(1) == int32(0) {
		goto L7
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v193)+188))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+60))
	v431 = m.T0[v430].(func(*base.Module, int32, int32, int32, int32) int32)(m, v193, v100+int32(10), int32(_a_F_afterTriggerInvokeEvents_6), v416)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	if v431 == int32(0) {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	if v167 != v188 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v454 = F_ExecFetchSlotHeapTuple(m, v450, int32(0), v32+int32(-58))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L120
	}
L110:
	;
	v436 = F_ExecGetChildToRootMap(m, v188)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v416
	v450 = v416
	goto L109
L113:
	;
	v438 = F_ExecGetTriggerNewSlot(m, v42, v167)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v438
	if v436 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	v442 = F_execute_attr_map_slot(m, v441, v416, v438)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v438)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+32))
	m.T0[v445].(func(*base.Module, int32, int32))(m, v438, v416)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L119
	}
L118:
	;
	v450 = v438
	goto L109
L119:
	;
	v450 = v438
	goto L109
L120:
	;
	v457 = v454
	goto L60
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = int32(442)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v482 & int32(7)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+12)))
	if v487&int32(16) != 0 {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v252)+52))
	if v466 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v467
	v469 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v463)+8)) = uint8(v469)
	goto L125
L124:
	;
	goto L125
L125:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v252)+56))
	if v471 == int32(0) {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v475
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)) = uint8(v477)
	goto L121
L127:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v191)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v490
	goto L129
L128:
	;
	goto L129
L129:
	;
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(-52)))) = v499
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(-56)))) = v502
	goto L131
L131:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v504 != v505 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[10])) = v507 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[9])) = v504
	goto L135
L133:
	;
	goto L134
L134:
	;
	v517 = F_ExecCallTriggerFunc(m, v32+int32(-48), v228, v173, int32(0), v49)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if v527 != v528 {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	if v517 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	if v517 == v521 {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if v517 == v523 {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	F_pfree(m, v517)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	goto L136
L142:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[10])) = v530
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerInvokeEvents[9])) = v527
	goto L145
L143:
	;
	goto L144
L144:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+7)))
	if v535 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	F_pfree(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+6)))
	if v541 == int32(1) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	F_pfree(m, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L4
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	if v168 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	if v171 == int32(0) {
		goto L47
	} else {
		goto L162
	}
L155:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	if v547 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+12))
	m.T0[v549].(func(*base.Module, int32))(m, v547)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	if v552 == int32(0) {
		goto L154
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v552)+8))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	m.T0[v556].(func(*base.Module, int32))(m, v552)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	goto L154
L162:
	;
	F_InstrStopNode(m, v171+v228*int32(416), float64(1))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	goto L47
L164:
	;
	v654 = v653 + v100
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if base.Ui32(v654) < base.Ui32(v655) {
		v100 = v654
		v103 = v621
		v105 = v623
		v107 = v625
		v109 = v627
		v111 = v629
		v112 = v630
		v115 = v633
		v117 = v635
		v120 = v638
		goto L21
	} else {
		goto L170
	}
L165:
	;
	if v643 != int32(805306368) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if v643 == int32(268435456) {
		v653 = int32(12)
		goto L164
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v653 = int32(16)
	goto L164
L169:
	;
	v653 = int32(4)
	goto L164
L170:
	;
	goto L22
L171:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v699 != 0 {
		v67 = v668
		v68 = v699
		v69 = v670
		v71 = v672
		v73 = v674
		v75 = v676
		v76 = v677
		v81 = v682
		v84 = v685
		goto L16
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v89
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 != v696 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89
	goto L171
L174:
	;
	goto L17
L175:
	;
	F_ExecDropSingleTupleTableSlot(m, v670)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	F_ExecDropSingleTupleTableSlot(m, v672)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	v725 = v676
	goto L12
L178:
	;
	if l2 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	F_ExecCloseResultRelations(m, v42)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L4
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	m.G0 = v34 - int32(-64)
	return v725 & int32(1)
L182:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	F_ExecResetTupleTable(m, v743, int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_FreeExecutorState(m, v42)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	goto L181
L185:
	;
	F_errmsg_internal(m, int32(_a_F_afterTriggerInvokeEvents_7), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_afterTriggerInvokeEvents_8), int32(_a_F_afterTriggerInvokeEvents_9), int32(_a_F_afterTriggerInvokeEvents_10))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_afterTriggerInvokeEvents_11), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_afterTriggerInvokeEvents_8), int32(_a_F_afterTriggerInvokeEvents_12), int32(_a_F_afterTriggerInvokeEvents_10))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_afterTriggerInvokeEvents_7), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_afterTriggerInvokeEvents_8), int32(_a_F_afterTriggerInvokeEvents_13), int32(_a_F_afterTriggerInvokeEvents_10))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errmsg_internal(m, int32(_a_F_afterTriggerInvokeEvents_11), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_afterTriggerInvokeEvents_8), int32(_a_F_afterTriggerInvokeEvents_14), int32(_a_F_afterTriggerInvokeEvents_10))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_afterTriggerInvokeEvents_15), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_afterTriggerInvokeEvents_16), int32(1264), int32(_a_F_afterTriggerInvokeEvents_17))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
