package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateAuxProcessResourceOwner(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[0]))
	v6 = F_MemoryContextAllocZero(m, v4, int32(616))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_CreateAuxProcessResourceOwner_0)
		v11 = v6 + int32(608)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+612)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v6)+608)) = v11
		*(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[1])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[2])) = v6
		F_on_shmem_exit(m, int32(2041), int64(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
func F_CreateInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
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
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
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
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	v20 = m.G0
	v22 = v20 - int32(400)
	m.G0 = v22
	v26 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = v22 + int32(232)
	v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v29, int32(1), int32(3), int32(184), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = int32(1)
	v41 = F_systable_beginscan(m, v26, int32(2680), v36, int32(0), v36, v29)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L235
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L231
	}
L6:
	;
	v43 = F_systable_getnext(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v43
	v50 = int32(0)
	goto L11
L9:
	;
	v91 = v36
	goto L10
L10:
	;
	F_systable_endscan(m, v41)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v67 == v68 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v91 = v72 + int32(1)
	goto L10
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v50 < v70 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = v70
	goto L16
L15:
	;
	v72 = v50
	goto L16
L16:
	;
	v73 = F_systable_getnext(m, v41)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v73 != 0 {
		v49 = v73
		v50 = v72
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	v100 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if int32(0) < v103 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L227
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L223
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L219
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L215
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L211
	}
L26:
	;
	v106 = int32(1)
	v112 = v106
	v113 = v103
	v120 = v106
	goto L29
L27:
	;
	goto L28
L28:
	;
	F_relation_close(m, v100, int32(3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L79
	}
L29:
	;
	v132 = v102 + v113<<(uint(int32(3))%32) + v112*int32(100)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+19)))
	if v133 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v138 = v132 - int32(68)
	v139 = F_SearchSysCacheCopyAttName(m, v136, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v256 = v113
	goto L33
L33:
	;
	v262 = v120 + int32(1)
	v263 = base.I32_extend16_s(v262)
	if v263 <= v256 {
		v112 = v263
		v113 = v256
		v120 = v262
		goto L29
	} else {
		goto L78
	}
L34:
	;
	if v139 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v144 = v132 - int32(72)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+68))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v148 = v146 + v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+68))
	if v145 != v149 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+76))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+76))
	if v151 != v152 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+96))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+96))
	if v154 != v155 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+86)))
	if v157 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+90)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+90)))
	if v175 != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+86)))
	if v160 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v144)+74)))
	v163 = F_findNotNullConstraintAttnum(m, v161, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v163 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+22)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v168)+106)))
	if v170 == int32(0) {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	goto L39
L45:
	;
	if l2 != 0 {
		goto L69
	} else {
		goto L70
	}
L46:
	;
	if v174 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v174 != 0 {
		goto L4
	} else {
		goto L68
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v174 == v175 {
		goto L45
	} else {
		goto L56
	}
L52:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_0), v22+int32(128))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_2), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_4), v22+int32(160))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+90)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+90)))
	if v213 == int32(115) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = int32(_a_F_CreateInheritance_5)
	goto L62
L61:
	;
	v216 = int32(_a_F_CreateInheritance_6)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v216
	if v210 == int32(115) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v222 = int32(_a_F_CreateInheritance_5)
	goto L65
L64:
	;
	v222 = int32(_a_F_CreateInheritance_6)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v222
	v227 = F_errdetail(m, int32(_a_F_CreateInheritance_7), v22+int32(144))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_8), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	goto L45
L69:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+89)) = uint8(v234)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+94)))
	v238 = v236 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v148)+94)) = uint16(v238)
	if base.I32_extend16_s(v238) != v238 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+119)))
	if v243 == int32(112) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+92)) = uint8(v246)
	goto L75
L74:
	;
	goto L75
L75:
	;
	F_CatalogTupleUpdate(m, v100, v139+int32(4), v139)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v139)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v256 = v254
	goto L33
L78:
	;
	goto L30
L79:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v290 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v293 = v22 + int32(344)
	F_ScanKeyInit(m, v293, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v287))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v301 = int32(1)
	v304 = F_systable_beginscan(m, v290, int32(2665), v301, int32(0), v301, v293)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v309 = F_build_attrmap_by_name(m, v306, v307, int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v311 = F_systable_getnext(m, v304)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L89
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L207
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L203
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L199
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L195
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L191
	}
L89:
	;
	if v311 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v320 = v311
	goto L93
L91:
	;
	goto L92
L92:
	;
	F_systable_endscan(m, v304)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L178
	}
L93:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+22)))
	v334 = v332 + v333
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+72)))
	switch v335 - int32(99) {
	case 0, 11:
		goto L96
	default:
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	v647 = F_systable_getnext(m, v304)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L176
	}
L96:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+106)))
	if v338 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v335 == int32(110) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v342 = F_extractNotNullColumn(m, v320)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v344 = int32(0)
	goto L100
L100:
	;
	v346 = v22 + int32(288)
	v350 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v346, int32(9), int32(3), int32(184), v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	v344 = v342
	goto L100
L102:
	;
	v354 = int32(1)
	v357 = F_systable_beginscan(m, v290, int32(2665), v354, int32(0), v354, v346)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L107
	}
L103:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+106)))
	if v585 == int32(1) {
		goto L87
	} else {
		goto L159
	}
L104:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+22)))
	v542 = v540 + v541
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+73)))
	v544 = v537 + v539
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+73)))
	if v543 != v545 {
		goto L88
	} else {
		goto L147
	}
L105:
	;
	v519 = F_extractNotNullColumn(m, v320)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L143
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L140
	}
L107:
	;
	v359 = F_systable_getnext(m, v357)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v359 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v362 = v334 + int32(4)
	v368 = v359
	goto L112
L110:
	;
	goto L111
L111:
	;
	F_systable_endscan(m, v357)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L134
	}
L112:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+22)))
	v384 = v382 + v383
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+72)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+72)))
	if v385 != v386 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L111
L114:
	;
	v459 = F_systable_getnext(m, v357)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L132
	}
L115:
	;
	switch v385 - int32(99) {
	case 0:
		goto L116
	default:
		v420 = v385
		goto L117
	case 11:
		goto L118
	}
L116:
	;
	v427 = v384 + int32(4)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if base.B2i32(v430 == int32(0))|base.B2i32(v430 != v433) != 0 {
		v451 = v430
		v452 = v433
		goto L125
	} else {
		goto L126
	}
L117:
	;
	if v420 != int32(99) {
		goto L103
	} else {
		goto L123
	}
L118:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v392 = F_extractNotNullColumn(m, v368)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v395 = int32(1)
	v396 = v392 - v395
	v400 = int32(*(*int16)(unsafe.Add(mBase, uint32(v394+v396<<(uint(v395)%32)))))
	if v344 != v400 {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v391<<(uint(int32(3))%32)+v344*int32(100))+19)))
	if v408 != 0 {
		goto L106
	} else {
		goto L121
	}
L121:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409+v410<<(uint(int32(3))%32)+v396*int32(100))+119)))
	if v417 != 0 {
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+72)))
	v420 = v418
	goto L117
L123:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+22)))
	v537 = v424
	v539 = v425
	goto L104
L124:
	;
	if v451-v452 == int32(0) {
		v537 = v382
		v539 = v383
		goto L104
	} else {
		goto L131
	}
L125:
	;
	goto L124
L126:
	;
	v436 = v362
	v437 = v427
	goto L127
L127:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+1)))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	if v441 == int32(0) {
		v451 = v441
		v452 = v440
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v451 = v441
	v452 = v440
	goto L125
L129:
	;
	v444 = int32(1)
	if v441 == v440 {
		v436 = v436 + v444
		v437 = v437 + v444
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L114
L132:
	;
	if v459 != 0 {
		v368 = v459
		goto L112
	} else {
		goto L133
	}
L133:
	;
	goto L113
L134:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+72)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v482 == int32(110) {
		goto L105
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v334 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_9), v22+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_10), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInheritance_12), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_13), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v522 = F_get_attname(m, v287, v519, int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v524 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_14), v22)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_15), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+74)))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+74)))
	if v547 != v548 {
		goto L88
	} else {
		goto L148
	}
L148:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v290)+52))
	v551 = F_decompile_conbin(m, v320, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v553 = F_decompile_conbin(m, v368, v550)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
	if base.B2i32(v557 == int32(0))|base.B2i32(v557 != v560) != 0 {
		v578 = v557
		v579 = v560
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v578-v579 != 0 {
		goto L88
	} else {
		goto L158
	}
L152:
	;
	goto L151
L153:
	;
	v563 = v551
	v564 = v553
	goto L154
L154:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+1)))
	if v568 == int32(0) {
		v578 = v568
		v579 = v567
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v578 = v568
	v579 = v567
	goto L152
L156:
	;
	v571 = int32(1)
	if v568 == v567 {
		v563 = v563 + v571
		v564 = v564 + v571
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	goto L103
L159:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+76)))
	if v588 != int32(1) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+75)))
	if v597 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+75)))
	if v591 != int32(1) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+76)))
	if v594 == int32(0) {
		goto L86
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+75)))
	if v600 == int32(0) {
		goto L85
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v603 = F_heap_copytuple(m, v368)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v603)+16))
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605)+22)))
	v607 = v605 + v606
	v608 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607)+104)))
	v610 = v608 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v607)+104)) = uint16(v610)
	if base.I32_extend16_s(v610) != v610 {
		goto L84
	} else {
		goto L169
	}
L169:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+119)))
	if v615 == int32(112) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v618 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v607)+103)) = uint8(v618)
	goto L172
L171:
	;
	goto L172
L172:
	;
	F_CatalogTupleUpdate(m, v290, v603+int32(4), v603)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_pfree(m, v603)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_systable_endscan(m, v357)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L95
L176:
	;
	if v647 != 0 {
		v320 = v647
		goto L93
	} else {
		goto L177
	}
L177:
	;
	goto L94
L178:
	;
	F_relation_close(m, v290, int32(3))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673)+119)))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_StoreSingleInheritance(m, v675, v676, v91)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v679 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+296)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v22)+292)) = v676
	v682 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = v682
	*(*int32)(unsafe.Add(mBase, uint32(v22)+352)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v22)+348)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v22)+344)) = v682
	if v674 == int32(112) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v697 = int32(97)
	goto L183
L182:
	;
	v697 = int32(110)
	goto L183
L183:
	;
	F_recordDependencyOn(m, v22+int32(344), v22+int32(288), v697)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInheritance[0]))
	if v701 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v703 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2611), v675, v703, v676, v703)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	F_SetRelationHasSubclass(m, v676, int32(1))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	F_relation_close(m, v26, int32(3))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	m.G0 = v22 + int32(400)
	return
L191:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v724 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_16), v22+int32(80))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_17), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v747 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v384 + v747
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v746 + v747
	F_errmsg(m, int32(_a_F_CreateInheritance_18), v22+int32(32))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_19), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v771 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v384 + v771
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v770 + v771
	F_errmsg(m, int32(_a_F_CreateInheritance_20), v22-int32(-64))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_21), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v795 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v384 + v795
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v794 + v795
	F_errmsg(m, int32(_a_F_CreateInheritance_22), v22+int32(48))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_23), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errmsg(m, int32(_a_F_CreateInheritance_24), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_25), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_26), v22+int32(96))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_27), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L215:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errmsg(m, int32(_a_F_CreateInheritance_24), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_28), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v868 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_14), v22+int32(176))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_29), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v890 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_30), v22+int32(192))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_31), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v912 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_32), v22+int32(208))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_33), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = v934 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_34), v22+int32(224))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_35), int32(_a_F_CreateInheritance_36))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_37), v22+int32(112))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_38), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateIntoRelDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(52))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(7)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(581)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(582)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(583)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(584)
		return v4
	}
}
func F_cr_circle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		*(*float64)(unsafe.Add(mBase, uint32(v8))) = v12
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v5
		*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v14
		return base.I64_extend_i32_u(v8)
	}
}
func F_createPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = F_palloc(m, int32(_a_F_createPostingTree_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(131)
	F_PageInit(m, v20, int32(_a_F_createPostingTree_0), int32(8))
	mBase = m.M
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v29 = v20 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+6)) = uint16(v24)
	goto L3
L3:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v33))) = int32(-1)
	if l2 == int32(0) {
		v89 = v6
		v92 = v6
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v96 = v92 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v96)
	v98 = F_GinNewBuffer(m, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v47 = v6
	v48 = v20 + int32(32)
	v49 = v6
	goto L6
L6:
	;
	v62 = F_ginCompressPostingList(m, l1+v49*int32(6), l2-v49, int32(384), v17+int32(24))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v89 = v79
	v92 = v71
	goto L4
L8:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v70 = (v64+int32(1))&int32(_a_F_createPostingTree_1) + int32(8)
	v71 = v70 + v47
	if base.Ui32(int32(_a_F_createPostingTree_2)) <= base.Ui32(v71) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v89 = v49
	v92 = v47
	goto L4
L10:
	;
	goto L11
L11:
	;
	if v70 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	base.MemoryCopy(m, v48, v62, v70)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_pfree(m, v62)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v79 = v49 + v75
	if base.Ui32(v79) < base.Ui32(l2) {
		v47 = v71
		v48 = v48 + v70
		v49 = v79
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	if v98 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	if v98 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+(v98^int32(-1))<<(uint(int32(2))%32))))
	v117 = v109
	goto L17
L20:
	;
	goto L21
L21:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[1]))
	v117 = v111 + v98<<(uint(int32(13))%32) + int32(-8192)
	goto L17
L22:
	;
	if l4 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[2]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v98^int32(-1))*int32(56))+16))
	v136 = v127
	goto L22
L24:
	;
	goto L25
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[3]))
	v130 = int32(56)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+v98*v130-v130)+16))
	v136 = v135
	goto L22
L26:
	;
	F_PredicateLockPageSplit(m, l0, v155, v136)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[2]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140+(l4^int32(-1))*int32(56))+16))
	v155 = v146
	goto L26
L28:
	;
	goto L29
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[3]))
	v149 = int32(56)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+l4*v149-v149)+16))
	v155 = v154
	goto L26
L30:
	;
	v158 = int32(_a_F_createPostingTree_3)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v160 + int32(1)
	F_PageRestoreTempPage(m, v20, v117)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_MarkBufferDirty(m, v98)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+118)))
	if v169 != int32(112) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v226 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L53
	}
L34:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v220 + int32(1)
	goto L33
L35:
	;
	v212 = int32(_a_F_createPostingTree_3)
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v214 - int32(1)
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L52
	}
L36:
	;
	v204 = int32(_a_F_createPostingTree_3)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v206 - int32(1)
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L50
	}
L37:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[5]))
	if v173 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v92
	F_XLogBeginInsert(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v176|l3 != 0 {
		goto L36
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l3 != 0 {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v178 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	goto L38
L45:
	;
	F_XLogRegisterData(m, v17+int32(24), int32(4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_XLogRegisterData(m, v117+int32(32), v92)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_XLogRegisterBuffer(m, int32(0), v98, int32(6))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v199 = F_XLogInsert(m, int32(13), int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = base.I64_rotl(v199, int64(32))
	goto L36
L50:
	;
	if l3 != 0 {
		goto L34
	} else {
		goto L51
	}
L51:
	;
	goto L33
L52:
	;
	goto L34
L53:
	;
	if v226 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v89
	F_errmsg_internal(m, int32(_a_F_createPostingTree_4), v17)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v89) < base.Ui32(l2) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_errfinish(m, int32(_a_F_createPostingTree_5), int32(1865), int32(_a_F_createPostingTree_6))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v238 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(40)
	v254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(45)
	v266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+60)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v136
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+77)) = uint8(base.B2i32(l3 != v254))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l2 - v89
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l1 + v89*int32(6)
	v282 = v17 + int32(90)
	v285 = v254
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v17 + int32(96)
	return v136
L62:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v301 = v298 + v285*int32(6)
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+4)) = uint16(v302)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v304
	v307 = v17 + int32(24)
	v310 = F_ginFindLeafPage(m, v307, int32(0), int32(1))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	F_ginInsertValue(m, v307, v310, v17+int32(12), l3)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if base.Ui32(v316) < base.Ui32(v317) {
		v285 = v316
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
}
func F_create_ctas_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_palloc0(m, int32(56))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(160)
		v19 = *(*int64)(unsafe.Add(mBase, _c_F_create_ctas_internal[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v19
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = int64(0)
		v25 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v22
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v29
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v31
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v33
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v37
		if v21 != 0 {
			v41 = int32(109)
		} else {
			v41 = int32(114)
		}
		v42 = int32(0)
		F_DefineRelation(m, l0, v14, v41, v42, v42, v42)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			F_CommandCounterIncrement(m)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
				v57 = F_transformRelOptions(m, int64(0), v51, int32(_a_F_create_ctas_internal_0), v11+int32(8), int32(1), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_heap_reloptions(m, int32(116), v57)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_NewRelationCreateToastTable(m, v61, v57)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v21 != 0 {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
								v65 = F_copyObjectImpl(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = m.G0
									v69 = v67 - int32(32)
									m.G0 = v69
									v72 = F_pstrdup(m, int32(_a_F_create_ctas_internal_1))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v65
										*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v65
										v78 = int32(0)
										v79 = int32(1)
										v85 = F_list_make1_impl(m, v79, v69+int32(12))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_DefineQueryRewrite(m, v69+int32(16), v72, v61, v78, v79, v79, v78, v85)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												m.G0 = v69 + int32(32)
												F_CommandCounterIncrement(m)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_create_final_unique_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v168 float64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 float64
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l4)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v22 = F_estimate_num_groups(m, l0, v17, v19, v7, v7)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+16)) = v22
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+45)))
	if v25 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+46)))
	if v192 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v28 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v36 = int32(0)
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v36<<(uint(int32(2))%32))))
	v52 = base.B2i32(v51 == v18)
	if v52 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L3
L9:
	;
	v177 = v36 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v177 < v178 {
		v36 = v177
		goto L7
	} else {
		goto L66
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v55 != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+64))
	v58 = v15 + int32(12)
	if l2 == v56 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L12
L14:
	;
	if v52|v136 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L15:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v124
	v136 = int32(1)
	goto L14
L16:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
	v136 = int32(1)
	goto L14
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
	v136 = int32(1)
	goto L14
L21:
	;
	goto L22
L22:
	;
	if v56 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v76
	v136 = v76
	goto L14
L24:
	;
	goto L25
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v80 = int32(0)
	if v80 < v79 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v83 = v79
	goto L28
L27:
	;
	v83 = v80
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v89 = int32(0)
	goto L29
L29:
	;
	if v89 < v84 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v100 = v96 + v89<<(uint(int32(2))%32)
	goto L33
L32:
	;
	v100 = int32(0)
	goto L33
L33:
	;
	if v89 == v83 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v83
	v136 = base.B2i32(v100 == int32(0))
	goto L14
L35:
	;
	goto L36
L36:
	;
	v106 = base.B2i32(v100 == int32(0))
	if v100 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v89
	v136 = v106
	goto L14
L38:
	;
	goto L39
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v110 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v89
	v136 = v106
	goto L14
L41:
	;
	goto L42
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110+v89<<(uint(int32(2))%32))))
	if v114 != v118 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v89
	v136 = int32(0)
	goto L14
L44:
	;
	v89 = v89 + int32(1)
	goto L29
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v140 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v150 = F_create_projection_path(m, l0, l5, v51, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_unique_paths[0])))
	if v144&int32(1) == int32(0) {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	if l2 != 0 {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	if v136 != 0 {
		v164 = v150
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v152 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v161 = F_create_incremental_sort_path(m, l0, l5, v150, l2, v152, float64(-1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_unique_paths[0])))
	if v154&int32(1) != 0 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v158 = F_create_sort_path(m, l5, v150, l2, float64(-1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v164 = v158
	goto L51
L60:
	;
	v164 = v161
	goto L51
L61:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v167 = v165
	goto L63
L62:
	;
	v167 = int32(0)
	goto L63
L63:
	;
	v168 = *(*float64)(unsafe.Add(mBase, uint32(l5)+16))
	v169 = F_create_unique_path(m, l5, v164, v167, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_add_path(m, l5, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L9
L66:
	;
	goto L8
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v196 = F_create_projection_path(m, l0, l5, v18, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	m.G0 = v15 + int32(16)
	return
L70:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v200 = int32(0)
	v203 = *(*float64)(unsafe.Add(mBase, uint32(l5)+16))
	v204 = F_create_agg_path(m, l0, l5, v196, v198, int32(2), v200, l3, v200, v200, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_add_path(m, l5, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L69
}
func F_create_material_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v50 float64
	_ = v50
	var v58 float64
	_ = v58
	var v64 float64
	_ = v64
	var v70 float64
	_ = v70
	v6 = F_palloc0(m, int32(80))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(1563368096040)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v15
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v19 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v24 = v22
		} else {
			v24 = int32(0)
		}
		v26 = v24 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v30
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_create_material_path[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+32)) = v37
		v46 = *(*float64)(unsafe.Add(mBase, _c_F_create_material_path[1]))
		v50 = base.F64_add(base.F64_mul(base.F64_add(v46, v46), v37), base.F64_sub(v36, v35))
		v58 = base.F64_mul(v37, base.F64_convert_i32_u((v39+int32(7))&int32(-8)+int32(24)))
		if base.F64_gt(v58, base.F64_convert_i32_u(v43<<(uint(int32(10))%32))) != 0 {
			v64 = *(*float64)(unsafe.Add(mBase, _c_F_create_material_path[2]))
			v70 = base.F64_add(base.F64_mul(v64, base.F64_ceil(base.F64_mul(v58, float64(0.0001220703125)))), v50)
		} else {
			v70 = v50
		}
		*(*float64)(unsafe.Add(mBase, uint32(v6)+48)) = v35
		*(*float64)(unsafe.Add(mBase, uint32(v6)+56)) = base.F64_add(v35, v70)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v34 + int32(0)
		return v6
	}
}
func F_create_ordinary_grouping_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v104 int32
	_ = v104
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v274 int32
	_ = v274
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int64
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 float64
	_ = v702
	var v703 int32
	_ = v703
	var v704 float64
	_ = v704
	var v705 int32
	_ = v705
	var v706 float64
	_ = v706
	var v707 float64
	_ = v707
	var v708 int32
	_ = v708
	var v709 float64
	_ = v709
	var v710 int32
	_ = v710
	var v711 float64
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v744 int32
	_ = v744
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1038 int32
	_ = v1038
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1072 int32
	_ = v1072
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1299 int32
	_ = v1299
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1350 int32
	_ = v1350
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1434 int32
	_ = v1434
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1515 int32
	_ = v1515
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1597 int32
	_ = v1597
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
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
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1728 int64
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1826 int32
	_ = v1826
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1847 int32
	_ = v1847
	var v1859 int32
	_ = v1859
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 float64
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 float64
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 float64
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1938 float64
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 float64
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 float64
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1976 int32
	_ = v1976
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2010 int32
	_ = v2010
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2086 int32
	_ = v2086
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2104 int32
	_ = v2104
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2272 int32
	_ = v2272
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2306 int32
	_ = v2306
	var v2326 int32
	_ = v2326
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2400 int32
	_ = v2400
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2419 int32
	_ = v2419
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2581 int32
	_ = v2581
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2634 int32
	_ = v2634
	v8 = int32(0)
	v29 = float64(0)
	v31 = m.G0
	v33 = v31 - int32(128)
	m.G0 = v33
	v35 = int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v36 == v8 {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v345&int32(4) == int32(0) {
		v1515 = v8
		goto L50
	} else {
		goto L51
	}
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v39 == int32(0) {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if v42 == int32(0) {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if v45 <= int32(0) {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	if v48 == int32(0) {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v53 == v51 {
		v74 = v51
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v74 != 0 {
		v325 = v8
		v335 = v35
		v338 = v8
		goto L1
	} else {
		goto L17
	}
L8:
	;
	goto L7
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v57 = v56
	goto L10
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if base.Ui32(int32(2)) <= base.Ui32(v61-int32(303)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v74 = int32(1)
	goto L8
L12:
	;
	if v61 != int32(293) {
		v74 = v51
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v57 = v60 + int32(72)
	goto L10
L14:
	;
	goto L11
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	if v68 != 0 {
		v74 = v51
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v75 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v308 = v306 & int32(4)
	v325 = int32(base.Ui32(v308) >> (uint(int32(2)) % 32))
	v335 = base.B2i32(v308 == int32(0))
	v338 = int32(base.Ui32(v308) >> (uint(int32(1)) % 32))
	goto L1
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+100))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v81 = F_get_sortgrouplist_exprs(m, v79, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v86 = int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+2)))
	if v88 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v325 = v8
	v335 = int32(0)
	v338 = v86
	goto L1
L24:
	;
	goto L25
L25:
	;
	v104 = v8
	goto L26
L26:
	;
	v123 = v104 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123+v124)))
	if v126 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L27:
	;
	goto L18
L28:
	;
	v129 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v130 <= v129 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v140 = v129
	v143 = v130
	goto L30
L30:
	;
	if v81 == int32(0) {
		v253 = v143
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	v274 = v140 + int32(1)
	if v274 < v253 {
		v140 = v274
		v143 = v253
		goto L30
	} else {
		goto L49
	}
L33:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v166 <= v165 {
		v253 = v143
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169+v140<<(uint(int32(2))%32))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175+v123)))
	v187 = v165
	goto L36
L35:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v253 = v242
	goto L32
L36:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v187<<(uint(int32(2))%32))))
	v213 = F_exprCollation(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v228 = int32(0)
	if base.B2i32(v177 == v228)|base.B2i32(v213 == v228)|base.B2i32(v177 == v213) == v228 {
		goto L18
	} else {
		goto L47
	}
L38:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v215 == int32(27) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v219 = v218
	goto L41
L40:
	;
	v219 = v212
	goto L41
L41:
	;
	v220 = F_equal(m, v219, v173)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	if v220 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v225 = v187 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v226 <= v225 {
		goto L35
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L37
L46:
	;
	v187 = v225
	goto L36
L47:
	;
	v237 = int32(0)
	v240 = v104 + int32(1)
	if v240 == v88 {
		v325 = v237
		v335 = v237
		v338 = v86
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v104 = v240
	goto L26
L49:
	;
	goto L31
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1515
	if v335 != 0 {
		goto L323
	} else {
		goto L324
	}
L51:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v351 == int32(0) {
		v383 = v8
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v384 == int32(0) {
		v391 = v8
		goto L68
	} else {
		goto L69
	}
L53:
	;
	v354 = int32(0)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v351)+44))
	if v356 == v354 {
		v377 = v354
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v377 != 0 {
		v383 = v8
		goto L52
	} else {
		goto L64
	}
L55:
	;
	goto L54
L56:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v360 = v359
	goto L57
L57:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	if base.Ui32(int32(2)) <= base.Ui32(v364-int32(303)) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v377 = int32(1)
	goto L55
L59:
	;
	if v364 != int32(293) {
		v377 = v354
		goto L55
	} else {
		goto L62
	}
L60:
	;
	v360 = v363 + int32(72)
	goto L57
L61:
	;
	goto L58
L62:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v363)+72))
	if v371 != 0 {
		v377 = v354
		goto L55
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v378)+44))
	if v380 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v381 = v378
	goto L67
L66:
	;
	v381 = int32(0)
	goto L67
L67:
	;
	v383 = v381
	goto L52
L68:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	if v393 != int32(1) {
		v403 = int32(0)
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v387 != int32(2) {
		v391 = v8
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v391 = v390
	goto L68
L71:
	;
	if v325|base.B2i32(v383|(v391|v403) != int32(0)) != int32(1) {
		v1515 = v8
		goto L50
	} else {
		goto L74
	}
L72:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v397 == v396 {
		v403 = v396
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v403 = v401
	goto L71
L74:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v413 = F_fetch_upper_rel(m, l0, int32(1), v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v413)+26)) = uint8(v415)
	v417 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v413)+32)) = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+4)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l2)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+164)) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+168)) = v423
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v413)+172)) = uint8(v425)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l2)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+176)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v431 = F_create_empty_pathtarget(m)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L20
	} else {
		goto L76
	}
L76:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v433 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v429 != 0 {
		goto L106
	} else {
		goto L107
	}
L78:
	;
	v546 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v437 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v438 <= v437 {
		v546 = v437
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v449 = v437
	v452 = int32(0)
	goto L82
L82:
	;
	v473 = v452 << (uint(int32(2)) % 32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v473+v474)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	if v477 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v546 = v532
	goto L77
L84:
	;
	v536 = v452 + int32(1)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v536 < v537 {
		v449 = v532
		v452 = v536
		goto L82
	} else {
		goto L105
	}
L85:
	;
	v530 = F_lappend(m, v449, v476)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L20
	} else {
		goto L104
	}
L86:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v473+v477)))
	if v481 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v484 == int32(0) {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	if v484 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	if v522 == int32(0) {
		goto L85
	} else {
		goto L102
	}
L90:
	;
	goto L89
L91:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if v490 <= int32(0) {
		v522 = int32(0)
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v522 = int32(0)
	goto L90
L94:
	;
	v493 = int32(0)
	if v493 < v490 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v496 = v490
	goto L97
L96:
	;
	v496 = v493
	goto L97
L97:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v500 = int32(0)
	goto L98
L98:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v497+v500<<(uint(int32(2))%32))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if v508 == v481 {
		v522 = v507
		goto L90
	} else {
		goto L100
	}
L99:
	;
	goto L93
L100:
	;
	v511 = v500 + int32(1)
	if v511 != v496 {
		v500 = v511
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_add_column_to_pathtarget(m, v431, v476, v481)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	v532 = v449
	goto L84
L104:
	;
	v532 = v530
	goto L84
L105:
	;
	goto L83
L106:
	;
	v569 = F_lappend(m, v546, v429)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L20
	} else {
		goto L109
	}
L107:
	;
	v571 = v546
	goto L108
L108:
	;
	v573 = F_pull_var_clause(m, v571, int32(25))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L20
	} else {
		goto L110
	}
L109:
	;
	v571 = v569
	goto L108
L110:
	;
	F_add_new_columns_to_pathtarget(m, v431, v573)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v577 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v675 = l5 + int32(8)
	F_list_free(m, v573)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L20
	} else {
		goto L125
	}
L113:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v580 <= int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v593 = v580
	v594 = int32(0)
	goto L115
L115:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v577)+12))
	v617 = v614 + v594<<(uint(int32(2))%32)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v619 == int32(9) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L112
L117:
	;
	v623 = F_palloc0(m, int32(72))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L20
	} else {
		goto L120
	}
L118:
	;
	v639 = v593
	goto L119
L119:
	;
	v642 = v594 + int32(1)
	if v642 < v639 {
		v593 = v639
		v594 = v642
		goto L115
	} else {
		goto L124
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = int32(9)
	base.MemoryCopy(m, v623, v618, int32(72))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+56)) = int32(6)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v623)+20))
	if v632 == int32(2281) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v635 = int32(17)
	goto L123
L122:
	;
	v635 = v632
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+8)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v623
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	v639 = v638
	goto L119
L124:
	;
	goto L116
L125:
	;
	F_list_free(m, v571)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	v680 = F_set_pathtarget_cost_width(m, l0, v431)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L20
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+40)) = v680
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
	if v683 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	base.MemoryFill(m, v675, int32(0), int32(80))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+36)))
	if v689 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	if v391 != 0 {
		goto L136
	} else {
		goto L137
	}
L131:
	;
	F_get_agg_clause_costs(m, l0, int32(6), v675)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L20
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v700 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)) = uint8(v700)
	goto L130
L134:
	;
	F_get_agg_clause_costs(m, l0, int32(9), l5+int32(48))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v702 = *(*float64)(unsafe.Add(mBase, uint32(v391)+32))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v704 = F_get_number_of_groups(m, l0, v702, l4, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L20
	} else {
		goto L139
	}
L137:
	;
	v706 = v29
	goto L138
L138:
	;
	if v403 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v706 = v704
	goto L138
L140:
	;
	v707 = *(*float64)(unsafe.Add(mBase, uint32(v403)+32))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v709 = F_get_number_of_groups(m, l0, v707, l4, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L20
	} else {
		goto L143
	}
L141:
	;
	v711 = v29
	goto L142
L142:
	;
	v713 = v345 & int32(1)
	v714 = int32(0)
	if base.B2i32(v713 == v714)|base.B2i32(v391 == v714) != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v711 = v709
	goto L142
L144:
	;
	v1006 = v345 & int32(2)
	v1007 = int32(0)
	if base.B2i32(v713 == v1007)|base.B2i32(v403 == v1007) != 0 {
		goto L218
	} else {
		goto L219
	}
L145:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v719 == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v722 <= int32(0) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v744 = v8
	goto L148
L148:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v719)+12))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v755+v744<<(uint(int32(2))%32))))
	v760 = F_get_useful_group_keys_orderings(m, l0, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L20
	} else {
		goto L151
	}
L149:
	;
	goto L144
L150:
	;
	v972 = v744 + int32(1)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v972 < v973 {
		v744 = v972
		goto L148
	} else {
		goto L217
	}
L151:
	;
	if v760 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v764 = int32(0)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v765 <= v764 {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	v778 = v764
	goto L154
L154:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v798+v778<<(uint(int32(2))%32))))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v759)+64))
	v806 = v33 + int32(16)
	if v803 == v804 {
		goto L160
	} else {
		goto L161
	}
L155:
	;
	goto L150
L156:
	;
	v938 = v778 + int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v938 < v939 {
		v778 = v938
		goto L154
	} else {
		goto L216
	}
L157:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+36)))
	if v914 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L158:
	;
	if v884 != 0 {
		goto L190
	} else {
		goto L191
	}
L159:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v872
	v884 = int32(1)
	goto L158
L160:
	;
	if v803 != 0 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v803 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = int32(0)
	v884 = int32(1)
	goto L158
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = int32(0)
	v884 = int32(1)
	goto L158
L165:
	;
	goto L166
L166:
	;
	if v804 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v824 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v824
	v884 = v824
	goto L158
L168:
	;
	goto L169
L169:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	v828 = int32(0)
	if v828 < v827 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v831 = v827
	goto L172
L171:
	;
	v831 = v828
	goto L172
L172:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	v837 = int32(0)
	goto L173
L173:
	;
	if v837 < v832 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v803)+12))
	v848 = v844 + v837<<(uint(int32(2))%32)
	goto L177
L176:
	;
	v848 = int32(0)
	goto L177
L177:
	;
	if v837 == v831 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v831
	v884 = base.B2i32(v848 == int32(0))
	goto L158
L179:
	;
	goto L180
L180:
	;
	v854 = base.B2i32(v848 == int32(0))
	if v848 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v837
	v884 = v854
	goto L158
L182:
	;
	goto L183
L183:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v804)+12))
	if v858 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v837
	v884 = v854
	goto L158
L185:
	;
	goto L186
L186:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v858+v837<<(uint(int32(2))%32))))
	if v862 != v866 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v837
	v884 = int32(0)
	goto L158
L188:
	;
	v837 = v837 + int32(1)
	goto L173
L190:
	;
	v911 = v759
	goto L157
L191:
	;
	goto L192
L192:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v759 == v391 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v896&int32(1) != 0 {
		goto L199
	} else {
		goto L200
	}
L194:
	;
	v896 = v886
	goto L193
L195:
	;
	goto L196
L196:
	;
	if v887 == int32(0) {
		goto L156
	} else {
		goto L197
	}
L197:
	;
	v891 = int32(1)
	if v886&v891 == int32(0) {
		goto L156
	} else {
		goto L198
	}
L198:
	;
	v896 = v891
	goto L193
L199:
	;
	v900 = v887
	goto L201
L200:
	;
	v900 = int32(0)
	goto L201
L201:
	;
	if v900 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v904 = F_create_sort_path(m, v413, v759, v803, float64(-1))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L20
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v907 = F_create_incremental_sort_path(m, l0, v413, v759, v803, v887, float64(-1))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L20
	} else {
		goto L207
	}
L205:
	;
	if v904 != 0 {
		v911 = v904
		goto L157
	} else {
		goto L206
	}
L206:
	;
	goto L156
L207:
	;
	if v907 == int32(0) {
		goto L156
	} else {
		goto L208
	}
L208:
	;
	v911 = v907
	goto L157
L209:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v350)+100))
	v919 = int32(0)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v802)+8))
	v924 = F_create_agg_path(m, l0, v413, v911, v917, base.B2i32(v918 != v919), int32(6), v922, v919, v675, v706)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L20
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v802)+8))
	v930 = F_create_group_path(m, l0, v413, v911, v928, int32(0), v706)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L20
	} else {
		goto L214
	}
L212:
	;
	F_add_path(m, v413, v924)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L20
	} else {
		goto L213
	}
L213:
	;
	goto L156
L214:
	;
	F_add_path(m, v413, v930)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L20
	} else {
		goto L215
	}
L215:
	;
	goto L156
L216:
	;
	goto L155
L217:
	;
	goto L149
L218:
	;
	v1299 = int32(0)
	if base.B2i32(v1006 == v1299)|base.B2i32(v391 == v1299) == v1299 {
		goto L292
	} else {
		goto L293
	}
L219:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v1012 == int32(0) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+4))
	if v1015 <= int32(0) {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	v1038 = int32(0)
	goto L222
L222:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+12))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1038<<(uint(int32(2))%32))))
	v1054 = F_get_useful_group_keys_orderings(m, l0, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L20
	} else {
		goto L225
	}
L223:
	;
	goto L218
L224:
	;
	v1266 = v1038 + int32(1)
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+4))
	if v1266 < v1267 {
		v1038 = v1266
		goto L222
	} else {
		goto L291
	}
L225:
	;
	if v1054 == int32(0) {
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v1058 = int32(0)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1059 <= v1058 {
		goto L224
	} else {
		goto L227
	}
L227:
	;
	v1072 = v1058
	goto L228
L228:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1092+v1072<<(uint(int32(2))%32))))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+64))
	v1100 = v33 + int32(16)
	if v1097 == v1098 {
		goto L234
	} else {
		goto L235
	}
L229:
	;
	goto L224
L230:
	;
	v1232 = v1072 + int32(1)
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1232 < v1233 {
		v1072 = v1232
		goto L228
	} else {
		goto L290
	}
L231:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+36)))
	if v1208 == int32(1) {
		goto L283
	} else {
		goto L284
	}
L232:
	;
	if v1178 != 0 {
		goto L264
	} else {
		goto L265
	}
L233:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1166
	v1178 = int32(1)
	goto L232
L234:
	;
	if v1097 != 0 {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if v1097 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = int32(0)
	v1178 = int32(1)
	goto L232
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = int32(0)
	v1178 = int32(1)
	goto L232
L239:
	;
	goto L240
L240:
	;
	if v1098 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1118
	v1178 = v1118
	goto L232
L242:
	;
	goto L243
L243:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	v1122 = int32(0)
	if v1122 < v1121 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1125 = v1121
	goto L246
L245:
	;
	v1125 = v1122
	goto L246
L246:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	v1131 = int32(0)
	goto L247
L247:
	;
	if v1131 < v1126 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	v1142 = v1138 + v1131<<(uint(int32(2))%32)
	goto L251
L250:
	;
	v1142 = int32(0)
	goto L251
L251:
	;
	if v1131 == v1125 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1125
	v1178 = base.B2i32(v1142 == int32(0))
	goto L232
L253:
	;
	goto L254
L254:
	;
	v1148 = base.B2i32(v1142 == int32(0))
	if v1142 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1131
	v1178 = v1148
	goto L232
L256:
	;
	goto L257
L257:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	if v1152 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1131
	v1178 = v1148
	goto L232
L259:
	;
	goto L260
L260:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1131<<(uint(int32(2))%32))))
	if v1156 != v1160 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100))) = v1131
	v1178 = int32(0)
	goto L232
L262:
	;
	v1131 = v1131 + int32(1)
	goto L247
L264:
	;
	v1205 = v1053
	goto L231
L265:
	;
	goto L266
L266:
	;
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v1053 == v403 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	if v1190&int32(1) != 0 {
		goto L273
	} else {
		goto L274
	}
L268:
	;
	v1190 = v1180
	goto L267
L269:
	;
	goto L270
L270:
	;
	if v1181 == int32(0) {
		goto L230
	} else {
		goto L271
	}
L271:
	;
	v1185 = int32(1)
	if v1180&v1185 == int32(0) {
		goto L230
	} else {
		goto L272
	}
L272:
	;
	v1190 = v1185
	goto L267
L273:
	;
	v1194 = v1181
	goto L275
L274:
	;
	v1194 = int32(0)
	goto L275
L275:
	;
	if v1194 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1198 = F_create_sort_path(m, v413, v1053, v1097, float64(-1))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L20
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1201 = F_create_incremental_sort_path(m, l0, v413, v1053, v1097, v1181, float64(-1))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L20
	} else {
		goto L281
	}
L279:
	;
	if v1198 != 0 {
		v1205 = v1198
		goto L231
	} else {
		goto L280
	}
L280:
	;
	goto L230
L281:
	;
	if v1201 == int32(0) {
		goto L230
	} else {
		goto L282
	}
L282:
	;
	v1205 = v1201
	goto L231
L283:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v350)+100))
	v1213 = int32(0)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+8))
	v1218 = F_create_agg_path(m, l0, v413, v1205, v1211, base.B2i32(v1212 != v1213), int32(6), v1216, v1213, v675, v711)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L20
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+8))
	v1224 = F_create_group_path(m, l0, v413, v1205, v1222, int32(0), v711)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L20
	} else {
		goto L288
	}
L286:
	;
	F_add_partial_path(m, v413, v1218)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L20
	} else {
		goto L287
	}
L287:
	;
	goto L230
L288:
	;
	F_add_partial_path(m, v413, v1224)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L20
	} else {
		goto L289
	}
L289:
	;
	goto L230
L290:
	;
	goto L229
L291:
	;
	goto L223
L292:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v1311 = F_create_agg_path(m, l0, v413, v391, v1306, int32(2), int32(6), v1309, int32(0), v675, v706)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L20
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1315 = int32(0)
	if base.B2i32(v1006 == v1315)|base.B2i32(v403 == v1315) == v1315 {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	F_add_path(m, v413, v1311)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L20
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v1327 = F_create_agg_path(m, l0, v413, v403, v1322, int32(2), int32(6), v1325, int32(0), v675, v711)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L20
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	if v383 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	F_add_partial_path(m, v413, v1327)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L20
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v413)+176))
	if v1498 == int32(0) {
		v1515 = v413
		goto L50
	} else {
		goto L320
	}
L303:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v383)+44))
	if v1333 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+26)))
	if v1414 != int32(1) {
		goto L302
	} else {
		goto L312
	}
L305:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	if v1336 <= int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1350 = int32(0)
	goto L307
L307:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+12))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1370+v1350<<(uint(int32(2))%32))))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v1376 = F_create_projection_path(m, l0, v413, v1374, v1375)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L20
	} else {
		goto L309
	}
L308:
	;
	goto L304
L309:
	;
	F_add_path(m, v413, v1376)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L20
	} else {
		goto L310
	}
L310:
	;
	v1381 = v1350 + int32(1)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	if v1381 < v1382 {
		v1350 = v1381
		goto L307
	} else {
		goto L311
	}
L311:
	;
	goto L308
L312:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v383)+52))
	if v1417 == int32(0) {
		goto L302
	} else {
		goto L313
	}
L313:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1420 <= int32(0) {
		goto L302
	} else {
		goto L314
	}
L314:
	;
	v1434 = int32(0)
	goto L315
L315:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+12))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1454+v1434<<(uint(int32(2))%32))))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v413)+40))
	v1460 = F_create_projection_path(m, l0, v413, v1458, v1459)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L20
	} else {
		goto L317
	}
L316:
	;
	goto L302
L317:
	;
	F_add_partial_path(m, v413, v1460)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L20
	} else {
		goto L318
	}
L318:
	;
	v1465 = v1434 + int32(1)
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1465 < v1466 {
		v1434 = v1465
		goto L315
	} else {
		goto L319
	}
L319:
	;
	goto L316
L320:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1498)+36))
	if v1501 == int32(0) {
		v1515 = v413
		goto L50
	} else {
		goto L321
	}
L321:
	;
	m.T0[v1501].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(1), l1, v413, l5)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L20
	} else {
		goto L322
	}
L322:
	;
	v1515 = v413
	goto L50
L323:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v1902 == int32(2) {
		goto L402
	} else {
		goto L403
	}
L324:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v1539 = int32(0)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	if v1540 == v1539 {
		goto L328
	} else {
		goto L329
	}
L325:
	;
	v1859 = int32(0)
	if base.B2i32(v1515 == v1859)|base.B2i32(v1847 == v1859) == v1859 {
		goto L394
	} else {
		goto L395
	}
L326:
	;
	if v1597 < int32(0) {
		goto L337
	} else {
		goto L338
	}
L327:
	;
	v1597 = base.I32_ctz(v1583) | v1584<<(uint(int32(5))%32)
	goto L326
L328:
	;
	v1597 = int32(-2)
	goto L326
L329:
	;
	v1548 = int32(0)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	if v1551 <= v1548 {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1554 = v1540 + int32(8)
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1554)))
	v1561 = v1558 & int32(-1)
	if v1561 != 0 {
		v1583 = v1561
		v1584 = v1548
		goto L327
	} else {
		goto L331
	}
L331:
	;
	v1562 = int32(1)
	if v1562 == v1551 {
		goto L328
	} else {
		goto L332
	}
L332:
	;
	v1566 = v1562
	goto L333
L333:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1554+v1566<<(uint(int32(2))%32))))
	if v1573 != 0 {
		v1583 = v1573
		v1584 = v1566
		goto L327
	} else {
		goto L335
	}
L334:
	;
	goto L328
L335:
	;
	v1575 = v1566 + int32(1)
	if v1575 != v1551 {
		v1566 = v1575
		goto L333
	} else {
		goto L336
	}
L336:
	;
	goto L334
L337:
	;
	v1836 = int32(0)
	v1841 = v1539
	v1847 = int32(1)
	goto L325
L338:
	;
	goto L339
L339:
	;
	v1611 = int32(0)
	v1616 = v1539
	v1622 = int32(1)
	v1623 = v1597
	goto L340
L340:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1634+v1623<<(uint(int32(2))%32))))
	v1639 = int32(0)
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+44))
	if v1641 == v1639 {
		v1662 = v1639
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v1836 = v1762
	v1841 = v1766
	v1847 = v1768
	goto L325
L342:
	;
	if v1662 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L343:
	;
	goto L342
L344:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1641)+12))
	v1645 = v1644
	goto L345
L345:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1645)))
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1648)))
	if base.Ui32(int32(2)) <= base.Ui32(v1649-int32(303)) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1662 = int32(1)
	goto L343
L347:
	;
	if v1649 != int32(293) {
		v1662 = v1639
		goto L343
	} else {
		goto L350
	}
L348:
	;
	v1645 = v1648 + int32(72)
	goto L345
L349:
	;
	goto L346
L350:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+72))
	if v1656 != 0 {
		v1662 = v1639
		goto L343
	} else {
		goto L351
	}
L351:
	;
	goto L349
L352:
	;
	v1665 = F_copy_pathtarget(m, v1538)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L20
	} else {
		goto L355
	}
L353:
	;
	v1762 = v1611
	v1766 = v1616
	v1768 = v1622
	goto L354
L354:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	if v1770 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L355:
	;
	base.MemoryCopy(m, v33+int32(16), l5, int32(104))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+8))
	v1674 = F_find_appinfos_by_relids(m, l0, v1671, v33+int32(124))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L20
	} else {
		goto L356
	}
L356:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+4))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1678 = F_adjust_appendrel_attrs(m, l0, v1676, v1677, v1674)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L20
	} else {
		goto L357
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+4)) = v1678
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1683 = F_adjust_appendrel_attrs(m, l0, v1681, v1682, v1674)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L20
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1683
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1688 = F_adjust_appendrel_attrs(m, l0, v1686, v1687, v1674)
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L20
	} else {
		goto L359
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1688
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+88)))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v33)+108))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	v1701 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v1694))|base.B2i32(int32(1)<<(uint(v1694)%32)&int32(44) == v1701) == v1701 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1716)+40)) = v1665
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+26)))
	if v1692&v1718 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L361:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+8))
	v1708 = F_fetch_upper_rel(m, l0, int32(2), v1707)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L20
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1714 = F_fetch_upper_rel(m, l0, int32(2), int32(0))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L20
	} else {
		goto L365
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1708)+4)) = int32(5)
	v1716 = v1708
	goto L360
L365:
	;
	v1716 = v1714
	goto L360
L366:
	;
	v1728 = *(*int64)(unsafe.Add(mBase, uint32(v1638)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1716)+32)) = v1728
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1716)+164)) = v1730
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v1716)+168)) = v1732
	v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1716)+172)) = uint8(v1734)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v1716)+176)) = v1736
	F_create_ordinary_grouping_paths(m, l0, v1638, v1716, l3, l4, v33+int32(16), v33+int32(12))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L20
	} else {
		goto L370
	}
L367:
	;
	v1722 = F_is_parallel_safe(m, l0, v1693)
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L20
	} else {
		goto L368
	}
L368:
	;
	if v1722 == int32(0) {
		goto L366
	} else {
		goto L369
	}
L369:
	;
	v1726 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1716)+26)) = uint8(v1726)
	goto L366
L370:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v1744 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	if v338 == int32(1) {
		goto L376
	} else {
		goto L377
	}
L372:
	;
	v1750 = v1616
	v1751 = int32(0)
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1748 = F_lappend(m, v1616, v1744)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L20
	} else {
		goto L375
	}
L375:
	;
	v1750 = v1748
	v1751 = v1622
	goto L371
L376:
	;
	F_set_cheapest(m, v1716)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L20
	} else {
		goto L379
	}
L377:
	;
	v1758 = v1611
	goto L378
L378:
	;
	F_pfree(m, v1674)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L20
	} else {
		goto L381
	}
L379:
	;
	v1756 = F_lappend(m, v1611, v1716)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L20
	} else {
		goto L380
	}
L380:
	;
	v1758 = v1756
	goto L378
L381:
	;
	v1762 = v1758
	v1766 = v1750
	v1768 = v1751
	goto L354
L382:
	;
	if int32(0) <= v1826 {
		v1611 = v1762
		v1616 = v1766
		v1622 = v1768
		v1623 = v1826
		goto L340
	} else {
		goto L393
	}
L383:
	;
	v1826 = base.I32_ctz(v1812) | v1813<<(uint(int32(5))%32)
	goto L382
L384:
	;
	v1826 = int32(-2)
	goto L382
L385:
	;
	v1777 = v1623 + int32(1)
	v1779 = int32(base.Ui32(v1777) >> (uint(int32(5)) % 32))
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1770)+4))
	if v1780 <= v1779 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1783 = v1770 + int32(8)
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1783+v1779<<(uint(int32(2))%32))))
	v1790 = v1787 & (int32(-1) << (uint(v1777) % 32))
	if v1790 != 0 {
		v1812 = v1790
		v1813 = v1779
		goto L383
	} else {
		goto L387
	}
L387:
	;
	v1792 = v1779 + int32(1)
	if v1792 == v1780 {
		goto L384
	} else {
		goto L388
	}
L388:
	;
	v1795 = v1792
	goto L389
L389:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1783+v1795<<(uint(int32(2))%32))))
	if v1802 != 0 {
		v1812 = v1802
		v1813 = v1795
		goto L383
	} else {
		goto L391
	}
L390:
	;
	goto L384
L391:
	;
	v1804 = v1795 + int32(1)
	if v1804 != v1780 {
		v1795 = v1804
		goto L389
	} else {
		goto L392
	}
L392:
	;
	goto L390
L393:
	;
	goto L341
L394:
	;
	F_add_paths_to_append_rel(m, l0, v1515, v1841)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L20
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	if v338 != int32(1) {
		goto L323
	} else {
		goto L398
	}
L397:
	;
	goto L396
L398:
	;
	F_add_paths_to_append_rel(m, l0, l2, v1836)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L20
	} else {
		goto L399
	}
L399:
	;
	goto L323
L400:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L20
	} else {
		goto L597
	}
L401:
	;
	m.G0 = v33 + int32(128)
	return
L402:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+44))
	if v1905 == int32(0) {
		goto L401
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	if v1515 == int32(0) {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	F_set_cheapest(m, v1515)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L20
	} else {
		goto L406
	}
L406:
	;
	goto L401
L407:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1925 = int32(0)
	v1926 = float64(0)
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1928 = *(*float64)(unsafe.Add(mBase, uint32(v1927)+32))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1930 = F_get_number_of_groups(m, l0, v1928, l4, v1929)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L20
	} else {
		goto L415
	}
L408:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+52))
	if v1912 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	F_gather_grouping_paths(m, l0, v1515)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L20
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+44))
	if v1915 == int32(0) {
		goto L407
	} else {
		goto L413
	}
L412:
	;
	goto L411
L413:
	;
	F_set_cheapest(m, v1515)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L20
	} else {
		goto L414
	}
L414:
	;
	goto L407
L415:
	;
	if v1515 == int32(0) {
		v1942 = v1925
		v1943 = v1926
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1945 = l5 + int32(48)
	v1947 = v1920 & int32(2)
	if v1920&int32(1) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L417:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+44))
	if v1934 == int32(0) {
		v1942 = v1925
		v1943 = v1926
		goto L416
	} else {
		goto L418
	}
L418:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+60))
	v1938 = *(*float64)(unsafe.Add(mBase, uint32(v1937)+32))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1940 = F_get_number_of_groups(m, l0, v1938, l4, v1939)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L20
	} else {
		goto L419
	}
L419:
	;
	v1942 = v1937
	v1943 = v1940
	goto L416
L420:
	;
	if v1947 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L421:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v1950 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if v1515 == int32(0) {
		goto L420
	} else {
		goto L500
	}
L423:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+4))
	if v1953 <= int32(0) {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v1976 = int32(0)
	goto L425
L425:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+12))
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1987+v1976<<(uint(int32(2))%32))))
	v1992 = F_get_useful_group_keys_orderings(m, l0, v1991)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L20
	} else {
		goto L428
	}
L426:
	;
	goto L422
L427:
	;
	v2211 = v1976 + int32(1)
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+4))
	if v2211 < v2212 {
		v1976 = v2211
		goto L425
	} else {
		goto L499
	}
L428:
	;
	if v1992 == int32(0) {
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v1996 = int32(0)
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	if v1997 <= v1996 {
		goto L427
	} else {
		goto L430
	}
L430:
	;
	v2010 = v1996
	goto L431
L431:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+12))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2030+v2010<<(uint(int32(2))%32))))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+4))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v1991)+64))
	v2038 = v33 + int32(16)
	if v2035 == v2036 {
		goto L437
	} else {
		goto L438
	}
L432:
	;
	goto L427
L433:
	;
	v2177 = v2010 + int32(1)
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	if v2177 < v2178 {
		v2010 = v2177
		goto L431
	} else {
		goto L498
	}
L434:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+108))
	if v2146 != 0 {
		goto L486
	} else {
		goto L487
	}
L435:
	;
	if v2116 != 0 {
		goto L467
	} else {
		goto L468
	}
L436:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2035)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2104
	v2116 = int32(1)
	goto L435
L437:
	;
	if v2035 != 0 {
		goto L436
	} else {
		goto L440
	}
L438:
	;
	goto L439
L439:
	;
	if v2035 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = int32(0)
	v2116 = int32(1)
	goto L435
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = int32(0)
	v2116 = int32(1)
	goto L435
L442:
	;
	goto L443
L443:
	;
	if v2036 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v2056 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2056
	v2116 = v2056
	goto L435
L445:
	;
	goto L446
L446:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2036)+4))
	v2060 = int32(0)
	if v2060 < v2059 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2063 = v2059
	goto L449
L448:
	;
	v2063 = v2060
	goto L449
L449:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2035)+4))
	v2069 = int32(0)
	goto L450
L450:
	;
	if v2069 < v2064 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2035)+12))
	v2080 = v2076 + v2069<<(uint(int32(2))%32)
	goto L454
L453:
	;
	v2080 = int32(0)
	goto L454
L454:
	;
	if v2069 == v2063 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2063
	v2116 = base.B2i32(v2080 == int32(0))
	goto L435
L456:
	;
	goto L457
L457:
	;
	v2086 = base.B2i32(v2080 == int32(0))
	if v2080 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2069
	v2116 = v2086
	goto L435
L459:
	;
	goto L460
L460:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2036)+12))
	if v2090 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2069
	v2116 = v2086
	goto L435
L462:
	;
	goto L463
L463:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2080)))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2090+v2069<<(uint(int32(2))%32))))
	if v2094 != v2098 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2069
	v2116 = int32(0)
	goto L435
L465:
	;
	v2069 = v2069 + int32(1)
	goto L450
L467:
	;
	v2143 = v1991
	goto L434
L468:
	;
	goto L469
L469:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v1991 == v1927 {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	if v2128&int32(1) != 0 {
		goto L476
	} else {
		goto L477
	}
L471:
	;
	v2128 = v2118
	goto L470
L472:
	;
	goto L473
L473:
	;
	if v2119 == int32(0) {
		goto L433
	} else {
		goto L474
	}
L474:
	;
	v2123 = int32(1)
	if v2118&v2123 == int32(0) {
		goto L433
	} else {
		goto L475
	}
L475:
	;
	v2128 = v2123
	goto L470
L476:
	;
	v2132 = v2119
	goto L478
L477:
	;
	v2132 = int32(0)
	goto L478
L478:
	;
	if v2132 == int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v2136 = F_create_sort_path(m, l2, v1991, v2035, float64(-1))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L20
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2139 = F_create_incremental_sort_path(m, l0, l2, v1991, v2035, v2119, float64(-1))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L20
	} else {
		goto L484
	}
L482:
	;
	if v2136 != 0 {
		v2143 = v2136
		goto L434
	} else {
		goto L483
	}
L483:
	;
	goto L433
L484:
	;
	if v2139 == int32(0) {
		goto L433
	} else {
		goto L485
	}
L485:
	;
	v2143 = v2139
	goto L434
L486:
	;
	F_consider_groupingsets_paths(m, l0, l2, v2143, int32(1), base.B2i32(v1947 != int32(0)), l4, l3, v1930)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L20
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	v2152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1924)+36)))
	if v2152 == int32(1) {
		goto L490
	} else {
		goto L491
	}
L489:
	;
	goto L433
L490:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+100))
	v2157 = int32(0)
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+8))
	v2161 = F_create_agg_path(m, l0, l2, v2143, v2155, base.B2i32(v2156 != v2157), v2157, v2160, v1923, l3, v1930)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L20
	} else {
		goto L493
	}
L491:
	;
	goto L492
L492:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+100))
	if v2165 == int32(0) {
		goto L433
	} else {
		goto L495
	}
L493:
	;
	F_add_path(m, l2, v2161)
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L20
	} else {
		goto L494
	}
L494:
	;
	goto L433
L495:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+8))
	v2169 = F_create_group_path(m, l0, l2, v2143, v2168, v1923, v1930)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L20
	} else {
		goto L496
	}
L496:
	;
	F_add_path(m, l2, v2169)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L20
	} else {
		goto L497
	}
L497:
	;
	goto L433
L498:
	;
	goto L432
L499:
	;
	goto L426
L500:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+44))
	if v2246 == int32(0) {
		goto L420
	} else {
		goto L501
	}
L501:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2246)+4))
	if v2249 <= int32(0) {
		goto L420
	} else {
		goto L502
	}
L502:
	;
	v2272 = int32(0)
	goto L503
L503:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2246)+12))
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2283+v2272<<(uint(int32(2))%32))))
	v2288 = F_get_useful_group_keys_orderings(m, l0, v2287)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L20
	} else {
		goto L506
	}
L504:
	;
	goto L420
L505:
	;
	v2498 = v2272 + int32(1)
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2246)+4))
	if v2498 < v2499 {
		v2272 = v2498
		goto L503
	} else {
		goto L572
	}
L506:
	;
	if v2288 == int32(0) {
		goto L505
	} else {
		goto L507
	}
L507:
	;
	v2292 = int32(0)
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2288)+4))
	if v2293 <= v2292 {
		goto L505
	} else {
		goto L508
	}
L508:
	;
	v2306 = v2292
	goto L509
L509:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2288)+12))
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2326+v2306<<(uint(int32(2))%32))))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2330)+4))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+64))
	v2334 = v33 + int32(16)
	if v2331 == v2332 {
		goto L515
	} else {
		goto L516
	}
L510:
	;
	goto L505
L511:
	;
	v2464 = v2306 + int32(1)
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2288)+4))
	if v2464 < v2465 {
		v2306 = v2464
		goto L509
	} else {
		goto L571
	}
L512:
	;
	v2442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1924)+36)))
	if v2442 == int32(1) {
		goto L564
	} else {
		goto L565
	}
L513:
	;
	if v2412 != 0 {
		goto L545
	} else {
		goto L546
	}
L514:
	;
	v2400 = *(*int32)(unsafe.Add(mBase, uint32(v2331)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2400
	v2412 = int32(1)
	goto L513
L515:
	;
	if v2331 != 0 {
		goto L514
	} else {
		goto L518
	}
L516:
	;
	goto L517
L517:
	;
	if v2331 == int32(0) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = int32(0)
	v2412 = int32(1)
	goto L513
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = int32(0)
	v2412 = int32(1)
	goto L513
L520:
	;
	goto L521
L521:
	;
	if v2332 == int32(0) {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v2352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2352
	v2412 = v2352
	goto L513
L523:
	;
	goto L524
L524:
	;
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+4))
	v2356 = int32(0)
	if v2356 < v2355 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2359 = v2355
	goto L527
L526:
	;
	v2359 = v2356
	goto L527
L527:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2331)+4))
	v2365 = int32(0)
	goto L528
L528:
	;
	if v2365 < v2360 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2331)+12))
	v2376 = v2372 + v2365<<(uint(int32(2))%32)
	goto L532
L531:
	;
	v2376 = int32(0)
	goto L532
L532:
	;
	if v2365 == v2359 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2359
	v2412 = base.B2i32(v2376 == int32(0))
	goto L513
L534:
	;
	goto L535
L535:
	;
	v2382 = base.B2i32(v2376 == int32(0))
	if v2376 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2365
	v2412 = v2382
	goto L513
L537:
	;
	goto L538
L538:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	if v2386 == int32(0) {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2365
	v2412 = v2382
	goto L513
L540:
	;
	goto L541
L541:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2376)))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2386+v2365<<(uint(int32(2))%32))))
	if v2390 != v2394 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334))) = v2365
	v2412 = int32(0)
	goto L513
L543:
	;
	v2365 = v2365 + int32(1)
	goto L528
L545:
	;
	v2439 = v2287
	goto L512
L546:
	;
	goto L547
L547:
	;
	v2414 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v2415 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v2287 == v1942 {
		goto L549
	} else {
		goto L550
	}
L548:
	;
	if v2424&int32(1) != 0 {
		goto L554
	} else {
		goto L555
	}
L549:
	;
	v2424 = v2414
	goto L548
L550:
	;
	goto L551
L551:
	;
	if v2415 == int32(0) {
		goto L511
	} else {
		goto L552
	}
L552:
	;
	v2419 = int32(1)
	if v2414&v2419 == int32(0) {
		goto L511
	} else {
		goto L553
	}
L553:
	;
	v2424 = v2419
	goto L548
L554:
	;
	v2428 = v2415
	goto L556
L555:
	;
	v2428 = int32(0)
	goto L556
L556:
	;
	if v2428 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v2432 = F_create_sort_path(m, l2, v2287, v2331, float64(-1))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L20
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	v2435 = F_create_incremental_sort_path(m, l0, l2, v2287, v2331, v2415, float64(-1))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L20
	} else {
		goto L562
	}
L560:
	;
	if v2432 != 0 {
		v2439 = v2432
		goto L512
	} else {
		goto L561
	}
L561:
	;
	goto L511
L562:
	;
	if v2435 == int32(0) {
		goto L511
	} else {
		goto L563
	}
L563:
	;
	v2439 = v2435
	goto L512
L564:
	;
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+100))
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2330)+8))
	v2451 = F_create_agg_path(m, l0, l2, v2439, v2445, base.B2i32(v2446 != int32(0)), int32(9), v2450, v1923, v1945, v1943)
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L20
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2330)+8))
	v2456 = F_create_group_path(m, l0, l2, v2439, v2455, v1923, v1943)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L20
	} else {
		goto L569
	}
L567:
	;
	F_add_path(m, l2, v2451)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L20
	} else {
		goto L568
	}
L568:
	;
	goto L511
L569:
	;
	F_add_path(m, l2, v2456)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L20
	} else {
		goto L570
	}
L570:
	;
	goto L511
L571:
	;
	goto L510
L572:
	;
	goto L504
L573:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v2559 != 0 {
		goto L586
	} else {
		goto L587
	}
L574:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+108))
	if v2533 != 0 {
		goto L576
	} else {
		goto L577
	}
L575:
	;
	if v1515 == int32(0) {
		goto L573
	} else {
		goto L582
	}
L576:
	;
	F_consider_groupingsets_paths(m, l0, l2, v1927, int32(0), int32(1), l4, l3, v1930)
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L20
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v2542 = F_create_agg_path(m, l0, l2, v1927, v2538, int32(2), int32(0), v2541, v1923, l3, v1930)
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L20
	} else {
		goto L580
	}
L579:
	;
	goto L575
L580:
	;
	F_add_path(m, l2, v2542)
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L20
	} else {
		goto L581
	}
L581:
	;
	goto L575
L582:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+44))
	if v2548 == int32(0) {
		goto L573
	} else {
		goto L583
	}
L583:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v2555 = F_create_agg_path(m, l0, l2, v1942, v2551, int32(2), int32(9), v2554, v1923, v1945, v1943)
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L20
	} else {
		goto L584
	}
L584:
	;
	F_add_path(m, l2, v2555)
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L20
	} else {
		goto L585
	}
L585:
	;
	goto L573
L586:
	;
	F_gather_grouping_paths(m, l0, l2)
	mBase = m.M
	v2561 = m.ExcPending
	if v2561 != 0 {
		goto L20
	} else {
		goto L589
	}
L587:
	;
	goto L588
L588:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v2562 == int32(0) {
		goto L400
	} else {
		goto L590
	}
L589:
	;
	goto L588
L590:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(l2)+176))
	if v2565 == int32(0) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[1]))
	if v2576 == int32(0) {
		goto L401
	} else {
		goto L595
	}
L592:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2565)+36))
	if v2568 == int32(0) {
		goto L591
	} else {
		goto L593
	}
L593:
	;
	m.T0[v2568].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L20
	} else {
		goto L594
	}
L594:
	;
	goto L591
L595:
	;
	m.T0[v2576].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L20
	} else {
		goto L596
	}
L596:
	;
	goto L401
L597:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L20
	} else {
		goto L598
	}
L598:
	;
	F_errmsg(m, int32(_a_F_create_ordinary_grouping_paths_0), int32(0))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L20
	} else {
		goto L599
	}
L599:
	;
	v2628 = F_errdetail(m, int32(_a_F_create_ordinary_grouping_paths_1), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L20
	} else {
		goto L600
	}
L600:
	;
	F_errfinish(m, int32(_a_F_create_ordinary_grouping_paths_2), int32(_a_F_create_ordinary_grouping_paths_3), int32(_a_F_create_ordinary_grouping_paths_4))
	mBase = m.M
	v2634 = m.ExcPending
	if v2634 != 0 {
		goto L20
	} else {
		goto L601
	}
L601:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_scan_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
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
	var v489 int32
	_ = v489
	var v491 float64
	_ = v491
	var v493 float64
	_ = v493
	var v495 float64
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
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
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
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
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v546 float64
	_ = v546
	var v548 float64
	_ = v548
	var v550 float64
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
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
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v780 float64
	_ = v780
	var v782 float64
	_ = v782
	var v784 float64
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
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
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1009 int32
	_ = v1009
	var v1011 float64
	_ = v1011
	var v1013 float64
	_ = v1013
	var v1015 float64
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1144 int32
	_ = v1144
	var v1146 float64
	_ = v1146
	var v1148 float64
	_ = v1148
	var v1150 float64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1185 int32
	_ = v1185
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1243 int32
	_ = v1243
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1341 int32
	_ = v1341
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1483 int32
	_ = v1483
	var v1485 float64
	_ = v1485
	var v1487 float64
	_ = v1487
	var v1489 float64
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1540 int32
	_ = v1540
	var v1542 float64
	_ = v1542
	var v1544 float64
	_ = v1544
	var v1546 float64
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1595 int32
	_ = v1595
	var v1597 float64
	_ = v1597
	var v1599 float64
	_ = v1599
	var v1601 float64
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1650 int32
	_ = v1650
	var v1652 float64
	_ = v1652
	var v1654 float64
	_ = v1654
	var v1656 float64
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1882 int32
	_ = v1882
	var v1884 float64
	_ = v1884
	var v1886 float64
	_ = v1886
	var v1888 float64
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1934 int32
	_ = v1934
	var v1936 float64
	_ = v1936
	var v1938 float64
	_ = v1938
	var v1940 float64
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1984 float64
	_ = v1984
	var v1986 float64
	_ = v1986
	var v1988 float64
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2077 int32
	_ = v2077
	var v2079 float64
	_ = v2079
	var v2081 float64
	_ = v2081
	var v2083 float64
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2126 float64
	_ = v2126
	var v2128 float64
	_ = v2128
	var v2130 float64
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2198 int32
	_ = v2198
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2316 int32
	_ = v2316
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2338 float64
	_ = v2338
	var v2340 float64
	_ = v2340
	var v2342 float64
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2373 int32
	_ = v2373
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2431 int32
	_ = v2431
	var v2436 int32
	_ = v2436
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2463 int32
	_ = v2463
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2479 int32
	_ = v2479
	var v2484 int32
	_ = v2484
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(176)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v22-int32(345)) <= base.Ui32(int32(1)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v32 = v27 + int32(96)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = v21 + int32(204)
	goto L1
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v36 = F_list_concat_copy(m, v33, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v40 = v33
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+335)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	switch v42 - int32(1) {
	case 0, 2:
		goto L14
	default:
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	v40 = v36
	goto L7
L10:
	;
	v66 = int32(0)
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v60 = F_order_qual_clauses(m, l0, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L20
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v53 == int32(358) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	if v41&int32(1) != 0 {
		v59 = v40
		goto L11
	} else {
		goto L16
	}
L14:
	;
	if v41&int32(1) != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(0)
	goto L10
L16:
	;
	v65 = int32(0)
	goto L10
L17:
	;
	v56 = int32(76)
	goto L19
L18:
	;
	v56 = int32(80)
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1+v56)))
	v59 = v58
	goto L11
L20:
	;
	v63 = F_extract_actual_clauses(m, v60, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v65 = v63
	goto L10
L22:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v465 - int32(335) {
	case 0:
		goto L137
	default:
		goto L133
	case 8:
		goto L149
	case 9:
		goto L148
	case 10:
		goto L132
	case 11:
		goto L147
	case 13:
		goto L146
	case 14:
		goto L145
	case 15:
		goto L144
	case 16:
		goto L143
	case 17:
		goto L142
	case 18:
		goto L140
	case 19:
		goto L141
	case 20:
		goto L139
	case 21:
		goto L138
	case 22:
		goto L136
	case 23:
		goto L135
	case 24:
		goto L134
	}
L23:
	;
	v68 = v66
	goto L25
L24:
	;
	v68 = l2
	goto L25
L25:
	;
	if v68 == int32(8) {
		v458 = v66
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v71 = F_use_physical_tlist(m, l0, l1, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v439, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L8
	} else {
		goto L125
	}
L28:
	;
	if v68&int32(4) == int32(0) {
		v458 = v287
		goto L22
	} else {
		goto L124
	}
L29:
	;
	if v71 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v73 == int32(346) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	if v373 == int32(0) {
		v458 = v66
		goto L22
	} else {
		goto L110
	}
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+92))
	v78 = F_copyObjectImpl(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v82 = m.G0
	v84 = v82 - int32(16)
	m.G0 = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v87 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if v68&int32(4) != 0 {
		v439 = v78
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v458 = v78
	goto L22
L38:
	;
	if v287 != 0 {
		goto L28
	} else {
		goto L95
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L8
	} else {
		goto L92
	}
L40:
	;
	m.G0 = v84 + int32(16)
	goto L38
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if base.Ui32(int32(6)) <= base.Ui32(v101-int32(3)) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v99 = v87 + v86<<(uint(int32(2))%32)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v99 = v93 + v86<<(uint(int32(2))%32) - int32(4)
	goto L41
L45:
	;
	switch v101 {
	case 0:
		goto L49
	case 1:
		goto L48
	default:
		goto L39
	}
L46:
	;
	goto L47
L47:
	;
	v232 = int32(0)
	F_expandRTE(m, v100, v86, v232, v232, int32(-1), int32(1), v232, v84+int32(12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L81
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+76))
	if v188 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v109 = F_table_open(m, v107, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+120)))
	if int32(0) < v112 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v118 = v106
	v122 = int32(1)
	goto L54
L52:
	;
	v170 = v106
	goto L53
L53:
	;
	F_relation_close(m, v109, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L68
	}
L54:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v139 = v132 + v133<<(uint(int32(3))%32) + v122*int32(100)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+19)))
	if v140 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v170 = v164
	goto L53
L56:
	;
	v141 = int32(0)
	F_relation_close(m, v109, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v146 = v139 - int32(72)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+88)))
	if v147 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v287 = v141
	goto L40
L60:
	;
	v148 = int32(0)
	F_relation_close(m, v109, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v153 = base.I32_extend16_s(v122)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146)+68))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+76))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)+96))
	v158 = F_makeVar(m, v86, v153, v154, v155, v156, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L64
	}
L63:
	;
	v287 = v148
	goto L40
L64:
	;
	v160 = int32(0)
	v162 = F_makeTargetEntry(m, v158, v153, v160, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	v164 = F_lappend(m, v118, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	if v122 != v112 {
		v118 = v164
		v122 = v122 + int32(1)
		goto L54
	} else {
		goto L67
	}
L67:
	;
	goto L55
L68:
	;
	v287 = v170
	goto L40
L69:
	;
	v287 = int32(0)
	goto L40
L70:
	;
	goto L71
L71:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v193 <= v192 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v287 = int32(0)
	goto L40
L73:
	;
	goto L74
L74:
	;
	v200 = int32(0)
	v204 = v192
	goto L75
L75:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v204<<(uint(int32(2))%32))))
	v219 = F_makeVarFromTargetEntry(m, v86, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L77
	}
L76:
	;
	v287 = v226
	goto L40
L77:
	;
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v218)+8)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+26)))
	v224 = F_makeTargetEntry(m, v219, v221, int32(0), v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v226 = F_lappend(m, v200, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v229 = v204 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v229 < v230 {
		v200 = v226
		v204 = v229
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v242 == int32(0) {
		v287 = v232
		goto L40
	} else {
		goto L82
	}
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v245 <= int32(0) {
		v287 = v232
		goto L40
	} else {
		goto L83
	}
L83:
	;
	v251 = v232
	v252 = int32(0)
	goto L84
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v252<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v270 != int32(6) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v287 = v279
	goto L40
L86:
	;
	v287 = int32(0)
	goto L40
L87:
	;
	goto L88
L88:
	;
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v269)+8)))
	v275 = int32(0)
	v277 = F_makeTargetEntry(m, v269, v274, v275, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v279 = F_lappend(m, v251, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v282 = v252 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v282 < v283 {
		v251 = v279
		v252 = v282
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v308
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_0), v84)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_1), int32(2122), int32(_a_F_create_scan_plan_2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v319 == int32(0) {
		v458 = v66
		goto L22
	} else {
		goto L96
	}
L96:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v323 <= int32(0) {
		v458 = v66
		goto L22
	} else {
		goto L97
	}
L97:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v330 = int32(1)
	v332 = v4
	v336 = v66
	goto L98
L98:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v332<<(uint(int32(2))%32))))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v348 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v458 = v366
	goto L22
L100:
	;
	v349 = F_replace_nestloop_params_mutator(m, v347, l0)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	v351 = v347
	goto L102
L102:
	;
	v353 = int32(0)
	v355 = F_makeTargetEntry(m, v351, base.I32_extend16_s(v330), v353, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v351 = v349
	goto L102
L104:
	;
	if v326 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v326+v330<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+16)) = v362
	goto L107
L106:
	;
	goto L107
L107:
	;
	v366 = F_lappend(m, v336, v355)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v369 = v332 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v369 < v370 {
		v330 = v330 + int32(1)
		v332 = v369
		v336 = v366
		goto L98
	} else {
		goto L109
	}
L109:
	;
	goto L99
L110:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v377 <= int32(0) {
		v458 = v66
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v384 = int32(1)
	v386 = v4
	v390 = v66
	goto L112
L112:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397+v386<<(uint(int32(2))%32))))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v402 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v458 = v420
	goto L22
L114:
	;
	v403 = F_replace_nestloop_params_mutator(m, v401, l0)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L117
	}
L115:
	;
	v405 = v401
	goto L116
L116:
	;
	v407 = int32(0)
	v409 = F_makeTargetEntry(m, v405, base.I32_extend16_s(v384), v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L8
	} else {
		goto L118
	}
L117:
	;
	v405 = v403
	goto L116
L118:
	;
	if v380 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v380+v384<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v409)+16)) = v416
	goto L121
L120:
	;
	goto L121
L121:
	;
	v420 = F_lappend(m, v390, v409)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v423 = v386 + int32(1)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v423 < v424 {
		v384 = v384 + int32(1)
		v386 = v423
		v390 = v420
		goto L112
	} else {
		goto L123
	}
L123:
	;
	goto L113
L124:
	;
	v439 = v287
	goto L27
L125:
	;
	v458 = v439
	goto L22
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L8
	} else {
		goto L600
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L8
	} else {
		goto L597
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L8
	} else {
		goto L594
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L8
	} else {
		goto L591
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L8
	} else {
		goto L588
	}
L131:
	;
	if v65 != 0 {
		goto L584
	} else {
		goto L585
	}
L132:
	;
	v2380 = F_create_indexscan_plan(m, l0, l1, v458, v40, int32(0))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L8
	} else {
		goto L583
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L8
	} else {
		goto L580
	}
L134:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v2274 == int32(0) {
		goto L564
	} else {
		goto L565
	}
L135:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+76))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v2095 != 0 {
		goto L513
	} else {
		goto L514
	}
L136:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+76))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1999 != 0 {
		goto L493
	} else {
		goto L494
	}
L137:
	;
	v1949 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L8
	} else {
		goto L482
	}
L138:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+76))
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1899 != 0 {
		goto L472
	} else {
		goto L473
	}
L139:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+76))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1667 != 0 {
		goto L424
	} else {
		goto L425
	}
L140:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+76))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1612 != 0 {
		goto L412
	} else {
		goto L413
	}
L141:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1555)+76))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1557 != 0 {
		goto L400
	} else {
		goto L401
	}
L142:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1498)+76))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1500 != 0 {
		goto L388
	} else {
		goto L389
	}
L143:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+76))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+148))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1163 = F_create_plan(m, v1161, v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L8
	} else {
		goto L316
	}
L144:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+76))
	v1027 = int32(0)
	if v40 == v1027 {
		v1107 = v1027
		goto L284
	} else {
		goto L285
	}
L145:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+76))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v795 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L146:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+76))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v571 = F_create_bitmap_subplan(m, l0, v564, v19+int32(172), v19+int32(168), v19+int32(164))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L170
	}
L147:
	;
	v560 = F_create_indexscan_plan(m, l0, l1, v458, v40, int32(1))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L8
	} else {
		goto L169
	}
L148:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+76))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v506 != 0 {
		goto L158
	} else {
		goto L159
	}
L149:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+76))
	v470 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L8
	} else {
		goto L150
	}
L150:
	;
	v473 = F_extract_actual_clauses(m, v470, int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v475 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v476 = F_replace_nestloop_params_mutator(m, v473, l0)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L155
	}
L153:
	;
	v478 = v473
	goto L154
L154:
	;
	v480 = F_palloc0(m, int32(80))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L8
	} else {
		goto L156
	}
L155:
	;
	v478 = v476
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v480)+72)) = v469
	*(*int64)(unsafe.Add(mBase, uint32(v480)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v480)+48)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v480)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = int32(343)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v480)+4)) = v489
	v491 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+8)) = v491
	v493 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+16)) = v493
	v495 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+24)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v480)+32)) = v498
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v480)+36)) = uint8(v500)
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v480)+37)) = uint8(v502)
	v2385 = v480
	goto L131
L157:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+32))
	v521 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L161
	}
L158:
	;
	v518 = v506 + v505<<(uint(int32(2))%32)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+52))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+12))
	v518 = v512 + v505<<(uint(int32(2))%32) - int32(4)
	goto L157
L161:
	;
	v524 = F_extract_actual_clauses(m, v521, int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v526 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v527 = F_replace_nestloop_params_mutator(m, v524, l0)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L8
	} else {
		goto L166
	}
L164:
	;
	v531 = v520
	v532 = v524
	goto L165
L165:
	;
	v534 = F_palloc0(m, int32(88))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L8
	} else {
		goto L168
	}
L166:
	;
	v529 = F_replace_nestloop_params_mutator(m, v520, l0)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	v531 = v529
	v532 = v527
	goto L165
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+80)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v534)+72)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v534)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+48)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v534)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(344)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = v544
	v546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+8)) = v546
	v548 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+16)) = v548
	v550 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+24)) = v550
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+32)) = v553
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+36)) = uint8(v555)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+37)) = uint8(v557)
	v2385 = v534
	goto L131
L169:
	;
	v2385 = v560
	goto L131
L170:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v573 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v579 = v571
	goto L175
L172:
	;
	goto L173
L173:
	;
	v638 = int32(0)
	if v40 == v638 {
		v741 = v638
		goto L183
	} else {
		goto L184
	}
L174:
	;
	v620 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v579)+84)) = uint8(v620)
	goto L173
L175:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	switch v593 - int32(341) {
	case 0:
		v599 = int32(72)
		goto L178
	case 1:
		goto L179
	default:
		goto L177
	case 6:
		goto L174
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L8
	} else {
		goto L180
	}
L177:
	;
	goto L176
L178:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v599+v579)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v579 = v603
	goto L175
L179:
	;
	v596 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v579)+72)) = uint8(v596)
	v599 = int32(76)
	goto L178
L180:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v608
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_3), v19+int32(16))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(_a_F_create_scan_plan_5), int32(_a_F_create_scan_plan_6))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L8
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v751 = F_order_qual_clauses(m, l0, v741)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L8
	} else {
		goto L218
	}
L184:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v641 <= int32(0) {
		v741 = v638
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	v650 = int32(0)
	v653 = v638
	goto L186
L186:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v650<<(uint(int32(2))%32))))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+10)))
	if v668 != 0 {
		v729 = v653
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v741 = v729
	goto L183
L188:
	;
	v732 = v650 + int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v732 < v733 {
		v650 = v732
		v653 = v729
		goto L186
	} else {
		goto L217
	}
L189:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	v670 = F_list_member(m, v645, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L8
	} else {
		goto L190
	}
L190:
	;
	if v670 != 0 {
		v729 = v653
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v667)+60))
	if v672 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v673 = int32(0)
	if v644 == v673 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v712 = F_contain_mutable_functions(m, v669)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L8
	} else {
		goto L209
	}
L195:
	;
	if v711 != 0 {
		v729 = v653
		goto L188
	} else {
		goto L208
	}
L196:
	;
	v711 = int32(0)
	goto L195
L197:
	;
	goto L198
L198:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	if v679 <= int32(0) {
		v705 = v673
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v711 = v705
	goto L195
L200:
	;
	v682 = int32(0)
	if v682 < v679 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v685 = v679
	goto L203
L202:
	;
	v685 = v682
	goto L203
L203:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v644)+12))
	v688 = int32(0)
	goto L204
L204:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v686+v688<<(uint(int32(2))%32))))
	v697 = base.B2i32(v696 == v672)
	if v696 == v672 {
		v705 = v697
		goto L199
	} else {
		goto L206
	}
L205:
	;
	v705 = v697
	goto L199
L206:
	;
	v699 = v688 + int32(1)
	if v699 != v685 {
		v688 = v699
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	goto L194
L209:
	;
	if v712 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v669
	v721 = F_list_make1_impl(m, int32(1), v19+int32(24))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L8
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v726 = F_lappend(m, v653, v667)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L8
	} else {
		goto L216
	}
L213:
	;
	v724 = F_predicate_implied_by(m, v721, v645, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L8
	} else {
		goto L214
	}
L214:
	;
	if v724 != 0 {
		v729 = v653
		goto L188
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	v729 = v726
	goto L188
L217:
	;
	goto L187
L218:
	;
	v754 = F_extract_actual_clauses(m, v751, int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L8
	} else {
		goto L219
	}
L219:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v757 = F_list_difference_ptr(m, v756, v754)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L8
	} else {
		goto L220
	}
L220:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v759 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v760 = F_replace_nestloop_params_mutator(m, v754, l0)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L8
	} else {
		goto L224
	}
L222:
	;
	v764 = v754
	v765 = v757
	goto L223
L223:
	;
	v767 = F_palloc0(m, int32(88))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L8
	} else {
		goto L226
	}
L224:
	;
	v762 = F_replace_nestloop_params_mutator(m, v757, l0)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L8
	} else {
		goto L225
	}
L225:
	;
	v764 = v760
	v765 = v762
	goto L223
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+80)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v767)+72)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v767)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v767)+52)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v767)+48)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v767)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = int32(348)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v767)+4)) = v778
	v780 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+8)) = v780
	v782 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+16)) = v782
	v784 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+24)) = v784
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v767)+32)) = v787
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+36)) = uint8(v789)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+37)) = uint8(v791)
	v2385 = v767
	goto L131
L227:
	;
	v965 = F_order_qual_clauses(m, l0, v951)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L8
	} else {
		goto L269
	}
L228:
	;
	v951 = v40
	goto L227
L229:
	;
	goto L230
L230:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v798 != int32(1) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v951 = v40
	goto L227
L232:
	;
	goto L233
L233:
	;
	if v40 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v951 = int32(0)
	goto L227
L235:
	;
	goto L236
L236:
	;
	v804 = int32(0)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v805 <= v804 {
		v951 = v804
		goto L227
	} else {
		goto L237
	}
L237:
	;
	v811 = v804
	v812 = int32(0)
	goto L238
L238:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v825+v812<<(uint(int32(2))%32))))
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+10)))
	if v830 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v951 = v931
	goto L227
L240:
	;
	v946 = v812 + int32(1)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v946 < v947 {
		v811 = v931
		v812 = v946
		goto L238
	} else {
		goto L268
	}
L241:
	;
	v831 = int32(0)
	if v795 == v831 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v869 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L255
	}
L243:
	;
	v869 = int32(0)
	goto L242
L244:
	;
	goto L245
L245:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v837 <= int32(0) {
		v863 = v831
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v869 = v863
	goto L242
L247:
	;
	v840 = int32(0)
	if v840 < v837 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v843 = v837
	goto L250
L249:
	;
	v843 = v840
	goto L250
L250:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v846 = int32(0)
	goto L251
L251:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v844+v846<<(uint(int32(2))%32))))
	v855 = base.B2i32(v854 == v829)
	if v854 == v829 {
		v863 = v855
		goto L246
	} else {
		goto L253
	}
L252:
	;
	v863 = v855
	goto L246
L253:
	;
	v857 = v846 + int32(1)
	if v857 != v843 {
		v846 = v857
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v870 = int32(0)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v829)+60))
	if base.B2i32(v871 == v870)|base.B2i32(v795 == v870) != 0 {
		v921 = v870
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if v921 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L266
	}
L257:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v877 <= int32(0) {
		v921 = v870
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v880 = int32(0)
	if v880 < v877 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v883 = v877
	goto L261
L260:
	;
	v883 = v880
	goto L261
L261:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v892 = int32(0)
	goto L262
L262:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v884+v892<<(uint(int32(2))%32))))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+60))
	v907 = base.B2i32(v906 == v871)
	if v906 == v871 {
		v921 = v907
		goto L256
	} else {
		goto L264
	}
L263:
	;
	v921 = v907
	goto L256
L264:
	;
	v909 = v892 + int32(1)
	if v909 != v883 {
		v892 = v909
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v927 = F_lappend(m, v811, v829)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L8
	} else {
		goto L267
	}
L267:
	;
	v931 = v927
	goto L240
L268:
	;
	goto L239
L269:
	;
	v968 = F_extract_actual_clauses(m, v795, int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L8
	} else {
		goto L270
	}
L270:
	;
	v971 = F_extract_actual_clauses(m, v965, int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L8
	} else {
		goto L271
	}
L271:
	;
	if v968 == int32(0) {
		v989 = v971
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v991 != 0 {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v968)+4))
	if v975 < int32(2) {
		v989 = v971
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v978 = F_make_orclause(m, v968)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v978
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = v978
	v985 = F_list_make1_impl(m, int32(1), v19+int32(28))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L8
	} else {
		goto L276
	}
L276:
	;
	v987 = F_list_difference(m, v971, v985)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	v989 = v987
	goto L272
L278:
	;
	v992 = F_replace_nestloop_params_mutator(m, v968, l0)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L8
	} else {
		goto L281
	}
L279:
	;
	v996 = v968
	v997 = v989
	goto L280
L280:
	;
	v999 = F_palloc0(m, int32(88))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L8
	} else {
		goto L283
	}
L281:
	;
	v994 = F_replace_nestloop_params_mutator(m, v989, l0)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L8
	} else {
		goto L282
	}
L282:
	;
	v996 = v992
	v997 = v994
	goto L280
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v999)+80)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v999)+72)) = v794
	*(*int64)(unsafe.Add(mBase, uint32(v999)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v999)+48)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v999)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = int32(349)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v1009
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+8)) = v1011
	v1013 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+16)) = v1013
	v1015 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+24)) = v1015
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+32)) = v1018
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+36)) = uint8(v1020)
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+37)) = uint8(v1022)
	v2385 = v999
	goto L131
L284:
	;
	v1118 = F_order_qual_clauses(m, l0, v1107)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L8
	} else {
		goto L307
	}
L285:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v1030 <= int32(0) {
		v1107 = v1027
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1037 = int32(0)
	v1039 = v1027
	goto L287
L287:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1050+v1037<<(uint(int32(2))%32))))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054)+10)))
	if v1055 != 0 {
		v1097 = v1039
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1107 = v1097
	goto L284
L289:
	;
	v1099 = v1037 + int32(1)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v1099 < v1100 {
		v1037 = v1099
		v1039 = v1097
		goto L287
	} else {
		goto L306
	}
L290:
	;
	v1056 = int32(0)
	if v1024 == v1056 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	if v1094 != 0 {
		v1097 = v1039
		goto L289
	} else {
		goto L304
	}
L292:
	;
	v1094 = int32(0)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	if v1062 <= int32(0) {
		v1088 = v1056
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1094 = v1088
	goto L291
L296:
	;
	v1065 = int32(0)
	if v1065 < v1062 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1068 = v1062
	goto L299
L298:
	;
	v1068 = v1065
	goto L299
L299:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+12))
	v1071 = int32(0)
	goto L300
L300:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v1071<<(uint(int32(2))%32))))
	v1080 = base.B2i32(v1079 == v1054)
	if v1079 == v1054 {
		v1088 = v1080
		goto L295
	} else {
		goto L302
	}
L301:
	;
	v1088 = v1080
	goto L295
L302:
	;
	v1082 = v1071 + int32(1)
	if v1082 != v1068 {
		v1071 = v1082
		goto L300
	} else {
		goto L303
	}
L303:
	;
	goto L301
L304:
	;
	v1095 = F_lappend(m, v1039, v1054)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L8
	} else {
		goto L305
	}
L305:
	;
	v1097 = v1095
	goto L289
L306:
	;
	goto L288
L307:
	;
	v1121 = F_extract_actual_clauses(m, v1024, int32(0))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L8
	} else {
		goto L308
	}
L308:
	;
	v1124 = F_extract_actual_clauses(m, v1118, int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L8
	} else {
		goto L309
	}
L309:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1126 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1127 = F_replace_nestloop_params_mutator(m, v1121, l0)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L8
	} else {
		goto L313
	}
L311:
	;
	v1131 = v1121
	v1132 = v1124
	goto L312
L312:
	;
	v1134 = F_palloc0(m, int32(88))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L8
	} else {
		goto L315
	}
L313:
	;
	v1129 = F_replace_nestloop_params_mutator(m, v1124, l0)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L8
	} else {
		goto L314
	}
L314:
	;
	v1131 = v1127
	v1132 = v1129
	goto L312
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+80)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+72)) = v1026
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+48)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1134))) = int32(350)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+4)) = v1144
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+8)) = v1146
	v1148 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+16)) = v1148
	v1150 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+24)) = v1150
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+32)) = v1153
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1134)+36)) = uint8(v1155)
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1134)+37)) = uint8(v1157)
	v2385 = v1134
	goto L131
L316:
	;
	v1165 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L8
	} else {
		goto L317
	}
L317:
	;
	v1168 = F_extract_actual_clauses(m, v1165, int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L8
	} else {
		goto L318
	}
L318:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1170 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1171 = int32(0)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+152))
	if v1172 == v1171 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1456 = v1168
	goto L321
L321:
	;
	v1471 = F_palloc0(m, int32(88))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L8
	} else {
		goto L386
	}
L322:
	;
	v1452 = F_replace_nestloop_params_mutator(m, v1168, l0)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L8
	} else {
		goto L385
	}
L323:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1175 <= int32(0) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1185 = v1171
	goto L327
L325:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L8
	} else {
		goto L382
	}
L326:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L8
	} else {
		goto L379
	}
L327:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+12))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1194+v1185<<(uint(int32(2))%32))))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+4))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	if v1200 != int32(6) {
		goto L332
	} else {
		goto L333
	}
L328:
	;
	goto L322
L329:
	;
	v1407 = v1185 + int32(1)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1407 < v1408 {
		v1185 = v1407
		goto L327
	} else {
		goto L378
	}
L330:
	;
	v1377 = F_palloc0(m, int32(12))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L8
	} else {
		goto L375
	}
L331:
	;
	v1262 = F_find_placeholder_info(m, l0, v1199)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L8
	} else {
		goto L350
	}
L332:
	;
	if v1200 == int32(321) {
		goto L331
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v1220 = F_bms_is_member(m, v1218, v1219)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L8
	} else {
		goto L339
	}
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L8
	} else {
		goto L336
	}
L336:
	;
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_7), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L8
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(597), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L8
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	if v1220 == int32(0) {
		goto L326
	} else {
		goto L340
	}
L340:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v1224 == int32(0) {
		goto L330
	} else {
		goto L341
	}
L341:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+4))
	if v1227 <= int32(0) {
		goto L330
	} else {
		goto L342
	}
L342:
	;
	v1230 = int32(0)
	if v1230 < v1227 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1233 = v1227
	goto L345
L344:
	;
	v1233 = v1230
	goto L345
L345:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+12))
	v1243 = int32(0)
	goto L346
L346:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1243<<(uint(int32(2))%32))))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+4))
	if v1257 == v1234 {
		goto L329
	} else {
		goto L348
	}
L347:
	;
	goto L330
L348:
	;
	v1260 = v1243 + int32(1)
	if v1233 != v1260 {
		v1243 = v1260
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+12))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v1266 = int32(0)
	if v1264 == v1266 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	if v1319 == int32(0) {
		goto L325
	} else {
		goto L365
	}
L352:
	;
	v1319 = int32(1)
	goto L351
L353:
	;
	goto L354
L354:
	;
	if v1265 == int32(0) {
		v1312 = v1266
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1319 = v1312
	goto L351
L356:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+4))
	if v1276 < v1275 {
		v1312 = v1266
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1278 = int32(1)
	if v1275 <= v1278 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1281 = v1278
	goto L360
L359:
	;
	v1281 = v1275
	goto L360
L360:
	;
	v1282 = int32(8)
	v1287 = int32(0)
	goto L361
L361:
	;
	v1294 = v1287 << (uint(int32(2)) % 32)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1264+v1282+v1294)))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1282+v1294)))
	v1301 = v1296 & (v1298 ^ int32(-1))
	v1303 = base.B2i32(v1301 == int32(0))
	if v1301 != 0 {
		v1312 = v1303
		goto L355
	} else {
		goto L363
	}
L362:
	;
	v1312 = v1303
	goto L355
L363:
	;
	v1305 = v1287 + int32(1)
	if v1305 != v1281 {
		v1287 = v1305
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v1322 == int32(0) {
		goto L330
	} else {
		goto L366
	}
L366:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+4))
	if v1325 <= int32(0) {
		goto L330
	} else {
		goto L367
	}
L367:
	;
	v1328 = int32(0)
	if v1328 < v1325 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1331 = v1325
	goto L370
L369:
	;
	v1331 = v1328
	goto L370
L370:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+12))
	v1341 = int32(0)
	goto L371
L371:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1333+v1341<<(uint(int32(2))%32))))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+4))
	if v1355 == v1332 {
		goto L329
	} else {
		goto L373
	}
L372:
	;
	goto L330
L373:
	;
	v1358 = v1341 + int32(1)
	if v1358 != v1331 {
		v1341 = v1358
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377))) = int32(361)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+4)) = v1381
	v1383 = F_copyObjectImpl(m, v1199)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L8
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+8)) = v1383
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v1387 = F_lappend(m, v1386, v1377)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L8
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v1387
	goto L329
L378:
	;
	goto L328
L379:
	;
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_10), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L8
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(543), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L8
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_10), int32(0))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L8
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(574), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L8
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	v1456 = v1452
	goto L321
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+80)) = v1163
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+72)) = v1160
	*(*int64)(unsafe.Add(mBase, uint32(v1471)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+48)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = int32(351)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+4)) = v1483
	v1485 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+8)) = v1485
	v1487 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+16)) = v1487
	v1489 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+24)) = v1489
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+32)) = v1492
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+36)) = uint8(v1494)
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+37)) = uint8(v1496)
	v2385 = v1471
	goto L131
L387:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1513)+68))
	v1515 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L8
	} else {
		goto L391
	}
L388:
	;
	v1512 = v1500 + v1499<<(uint(int32(2))%32)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+52))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+12))
	v1512 = v1506 + v1499<<(uint(int32(2))%32) - int32(4)
	goto L387
L391:
	;
	v1518 = F_extract_actual_clauses(m, v1515, int32(0))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L8
	} else {
		goto L392
	}
L392:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1520 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1521 = F_replace_nestloop_params_mutator(m, v1518, l0)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L8
	} else {
		goto L396
	}
L394:
	;
	v1525 = v1514
	v1526 = v1518
	goto L395
L395:
	;
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513)+72)))
	v1529 = F_palloc0(m, int32(88))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L8
	} else {
		goto L398
	}
L396:
	;
	v1523 = F_replace_nestloop_params_mutator(m, v1514, l0)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L8
	} else {
		goto L397
	}
L397:
	;
	v1525 = v1523
	v1526 = v1521
	goto L395
L398:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+84)) = uint8(v1527)
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+80)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+72)) = v1499
	*(*int64)(unsafe.Add(mBase, uint32(v1529)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+48)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1529))) = int32(352)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+4)) = v1540
	v1542 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+8)) = v1542
	v1544 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+16)) = v1544
	v1546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+24)) = v1546
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+32)) = v1549
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+36)) = uint8(v1551)
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+37)) = uint8(v1553)
	v2385 = v1529
	goto L131
L399:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1569)))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+76))
	v1572 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L8
	} else {
		goto L403
	}
L400:
	;
	v1569 = v1557 + v1556<<(uint(int32(2))%32)
	goto L399
L401:
	;
	goto L402
L402:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+52))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+12))
	v1569 = v1563 + v1556<<(uint(int32(2))%32) - int32(4)
	goto L399
L403:
	;
	v1575 = F_extract_actual_clauses(m, v1572, int32(0))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L8
	} else {
		goto L404
	}
L404:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1577 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1578 = F_replace_nestloop_params_mutator(m, v1575, l0)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L8
	} else {
		goto L408
	}
L406:
	;
	v1582 = v1571
	v1583 = v1575
	goto L407
L407:
	;
	v1585 = F_palloc0(m, int32(88))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L8
	} else {
		goto L410
	}
L408:
	;
	v1580 = F_replace_nestloop_params_mutator(m, v1571, l0)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L8
	} else {
		goto L409
	}
L409:
	;
	v1582 = v1580
	v1583 = v1578
	goto L407
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+80)) = v1582
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+72)) = v1556
	*(*int64)(unsafe.Add(mBase, uint32(v1585)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+48)) = v1583
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1585))) = int32(354)
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+4)) = v1595
	v1597 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+8)) = v1597
	v1599 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+16)) = v1599
	v1601 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+24)) = v1601
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+32)) = v1604
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1585)+36)) = uint8(v1606)
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1585)+37)) = uint8(v1608)
	v2385 = v1585
	goto L131
L411:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1624)))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+80))
	v1627 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L8
	} else {
		goto L415
	}
L412:
	;
	v1624 = v1612 + v1611<<(uint(int32(2))%32)
	goto L411
L413:
	;
	goto L414
L414:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1616)+52))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+12))
	v1624 = v1618 + v1611<<(uint(int32(2))%32) - int32(4)
	goto L411
L415:
	;
	v1630 = F_extract_actual_clauses(m, v1627, int32(0))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L8
	} else {
		goto L416
	}
L416:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1632 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1633 = F_replace_nestloop_params_mutator(m, v1630, l0)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L8
	} else {
		goto L420
	}
L418:
	;
	v1637 = v1626
	v1638 = v1630
	goto L419
L419:
	;
	v1640 = F_palloc0(m, int32(88))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L8
	} else {
		goto L422
	}
L420:
	;
	v1635 = F_replace_nestloop_params_mutator(m, v1626, l0)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L8
	} else {
		goto L421
	}
L421:
	;
	v1637 = v1635
	v1638 = v1633
	goto L419
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+80)) = v1637
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+72)) = v1611
	*(*int64)(unsafe.Add(mBase, uint32(v1640)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+48)) = v1638
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1640))) = int32(353)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+4)) = v1650
	v1652 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+8)) = v1652
	v1654 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+16)) = v1654
	v1656 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+24)) = v1656
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+32)) = v1659
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1640)+36)) = uint8(v1661)
	v1663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1640)+37)) = uint8(v1663)
	v2385 = v1640
	goto L131
L423:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1679)))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+88))
	v1684 = l0
	v1685 = v1681
	goto L428
L424:
	;
	v1679 = v1667 + v1666<<(uint(int32(2))%32)
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+52))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+12))
	v1679 = v1673 + v1666<<(uint(int32(2))%32) - int32(4)
	goto L423
L427:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+4))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+48))
	if v1720 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L428:
	;
	if v1685 == int32(0) {
		goto L427
	} else {
		goto L430
	}
L429:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L8
	} else {
		goto L432
	}
L430:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+16))
	if v1702 != 0 {
		v1684 = v1702
		v1685 = v1685 - int32(1)
		goto L428
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1707
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(96))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L8
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3757), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L8
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+84))
	if v1813 == int32(0) {
		goto L130
	} else {
		goto L453
	}
L436:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L8
	} else {
		goto L450
	}
L437:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+4))
	if v1723 <= int32(0) {
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+12))
	v1732 = int32(0)
	goto L439
L439:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1727+v1732<<(uint(int32(2))%32))))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+4))
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749))))
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726))))
	if base.B2i32(v1752 == int32(0))|base.B2i32(v1752 != v1755) != 0 {
		v1773 = v1752
		v1774 = v1755
		goto L442
	} else {
		goto L443
	}
L440:
	;
	goto L436
L441:
	;
	if v1773-v1774 == int32(0) {
		goto L435
	} else {
		goto L448
	}
L442:
	;
	goto L441
L443:
	;
	v1758 = v1749
	v1759 = v1726
	goto L444
L444:
	;
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759)+1)))
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+1)))
	if v1763 == int32(0) {
		v1773 = v1763
		v1774 = v1762
		goto L442
	} else {
		goto L446
	}
L445:
	;
	v1773 = v1763
	v1774 = v1762
	goto L442
L446:
	;
	v1766 = int32(1)
	if v1763 == v1762 {
		v1758 = v1758 + v1766
		v1759 = v1759 + v1766
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	v1779 = v1732 + int32(1)
	if v1723 != v1779 {
		v1732 = v1779
		goto L439
	} else {
		goto L449
	}
L449:
	;
	goto L440
L450:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1801
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_13), v19+int32(32))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L8
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3775), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L8
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+4))
	if v1816 <= v1732 {
		goto L130
	} else {
		goto L454
	}
L454:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+12))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1818+v1732<<(uint(int32(2))%32))))
	if v1822 <= int32(0) {
		goto L129
	} else {
		goto L455
	}
L455:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+80))
	if v1825 == int32(0) {
		goto L128
	} else {
		goto L456
	}
L456:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+4))
	if v1828 <= int32(0) {
		goto L128
	} else {
		goto L457
	}
L457:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+12))
	v1836 = int32(0)
	goto L458
L458:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1831+v1836<<(uint(int32(2))%32))))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	if v1822 != v1853 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+40))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+12))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1859)))
	v1861 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L8
	} else {
		goto L464
	}
L460:
	;
	v1856 = v1836 + int32(1)
	if v1856 != v1828 {
		v1836 = v1856
		goto L458
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	goto L459
L463:
	;
	goto L128
L464:
	;
	v1864 = F_extract_actual_clauses(m, v1861, int32(0))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L8
	} else {
		goto L465
	}
L465:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1866 != 0 {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v1867 = F_replace_nestloop_params_mutator(m, v1864, l0)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L8
	} else {
		goto L469
	}
L467:
	;
	v1869 = v1864
	goto L468
L468:
	;
	v1871 = F_palloc0(m, int32(88))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L8
	} else {
		goto L470
	}
L469:
	;
	v1869 = v1867
	goto L468
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+84)) = v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+80)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+72)) = v1666
	*(*int64)(unsafe.Add(mBase, uint32(v1871)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+48)) = v1869
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1871))) = int32(355)
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+4)) = v1882
	v1884 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+8)) = v1884
	v1886 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+16)) = v1886
	v1888 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+24)) = v1888
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1890)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+32)) = v1891
	v1893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1871)+36)) = uint8(v1893)
	v1895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1871)+37)) = uint8(v1895)
	v2385 = v1871
	goto L131
L471:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1911)))
	v1913 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L8
	} else {
		goto L475
	}
L472:
	;
	v1911 = v1899 + v1898<<(uint(int32(2))%32)
	goto L471
L473:
	;
	goto L474
L474:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+52))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1904)+12))
	v1911 = v1905 + v1898<<(uint(int32(2))%32) - int32(4)
	goto L471
L475:
	;
	v1916 = F_extract_actual_clauses(m, v1913, int32(0))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L8
	} else {
		goto L476
	}
L476:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1918 != 0 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1919 = F_replace_nestloop_params_mutator(m, v1916, l0)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L8
	} else {
		goto L480
	}
L478:
	;
	v1921 = v1916
	goto L479
L479:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1912)+108))
	v1924 = F_palloc0(m, int32(88))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L8
	} else {
		goto L481
	}
L480:
	;
	v1921 = v1919
	goto L479
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+80)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+72)) = v1898
	*(*int64)(unsafe.Add(mBase, uint32(v1924)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+48)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1924))) = int32(356)
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+4)) = v1934
	v1936 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+8)) = v1936
	v1938 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+16)) = v1938
	v1940 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+24)) = v1940
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+32)) = v1943
	v1945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1924)+36)) = uint8(v1945)
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1924)+37)) = uint8(v1947)
	v2385 = v1924
	goto L131
L482:
	;
	v1952 = F_extract_actual_clauses(m, v1949, int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L8
	} else {
		goto L483
	}
L483:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1954 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1955 = F_replace_nestloop_params_mutator(m, v1952, l0)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L8
	} else {
		goto L487
	}
L485:
	;
	v1957 = v1952
	goto L486
L486:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1960 = F_palloc0(m, int32(88))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L8
	} else {
		goto L488
	}
L487:
	;
	v1957 = v1955
	goto L486
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1960)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v1960))) = int32(335)
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1958)+4))
	switch v1971 - int32(1) {
	case 0, 2:
		v1975 = int32(2)
		goto L490
	default:
		goto L491
	case 3, 4:
		v1977 = int32(3)
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+76)) = v1957
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+72)) = v1977
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1958)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+80)) = v1980
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+4)) = v1982
	v1984 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1960)+8)) = v1984
	v1986 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1960)+16)) = v1986
	v1988 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1960)+24)) = v1988
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1960)+32)) = v1991
	v1993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1960)+36)) = uint8(v1993)
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1960)+37)) = uint8(v1995)
	v2385 = v1960
	goto L131
L490:
	;
	v1977 = v1975
	goto L489
L491:
	;
	v1975 = int32(1)
	goto L490
L492:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+88))
	if v2013 == int32(0) {
		goto L127
	} else {
		goto L496
	}
L493:
	;
	v2011 = v1999 + v1998<<(uint(int32(2))%32)
	goto L492
L494:
	;
	goto L495
L495:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v2003)+52))
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v2004)+12))
	v2011 = v2005 + v1998<<(uint(int32(2))%32) - int32(4)
	goto L492
L496:
	;
	v2018 = l0
	v2019 = v2013
	goto L498
L497:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+360))
	if v2053 < int32(0) {
		goto L126
	} else {
		goto L505
	}
L498:
	;
	v2033 = v2019 - int32(1)
	if v2033 == int32(0) {
		goto L497
	} else {
		goto L500
	}
L499:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L8
	} else {
		goto L502
	}
L500:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+16))
	if v2036 != 0 {
		v2018 = v2036
		v2019 = v2033
		goto L498
	} else {
		goto L501
	}
L501:
	;
	goto L499
L502:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v2041
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(144))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L8
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3929), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L8
	} else {
		goto L504
	}
L504:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L505:
	;
	v2056 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L8
	} else {
		goto L506
	}
L506:
	;
	v2059 = F_extract_actual_clauses(m, v2056, int32(0))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L8
	} else {
		goto L507
	}
L507:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2061 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2062 = F_replace_nestloop_params_mutator(m, v2059, l0)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L8
	} else {
		goto L511
	}
L509:
	;
	v2064 = v2059
	goto L510
L510:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+360))
	v2067 = F_palloc0(m, int32(88))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L8
	} else {
		goto L512
	}
L511:
	;
	v2064 = v2062
	goto L510
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+80)) = v2065
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+72)) = v1998
	*(*int64)(unsafe.Add(mBase, uint32(v2067)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+48)) = v2064
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+44)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v2067))) = int32(357)
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+4)) = v2077
	v2079 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2067)+8)) = v2079
	v2081 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2067)+16)) = v2081
	v2083 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2067)+24)) = v2083
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2085)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+32)) = v2086
	v2088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2067)+36)) = uint8(v2088)
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2067)+37)) = uint8(v2090)
	v2385 = v2067
	goto L131
L513:
	;
	v2097 = F_create_plan_recurse(m, l0, v2095, int32(1))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L8
	} else {
		goto L516
	}
L514:
	;
	v2099 = int32(0)
	goto L515
L515:
	;
	if v2093 != 0 {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v2099 = v2097
	goto L515
L517:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2100 != 0 {
		goto L521
	} else {
		goto L522
	}
L518:
	;
	v2117 = int32(0)
	goto L519
L519:
	;
	v2118 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L8
	} else {
		goto L524
	}
L520:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2112)))
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2113)+16))
	v2117 = v2114
	goto L519
L521:
	;
	v2112 = v2100 + v2093<<(uint(int32(2))%32)
	goto L520
L522:
	;
	goto L523
L523:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+52))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+12))
	v2112 = v2106 + v2093<<(uint(int32(2))%32) - int32(4)
	goto L520
L524:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+176))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+12))
	v2122 = m.T0[v2121].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2092, v2117, l1, v458, v2118, v2099)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L8
	} else {
		goto L525
	}
L525:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+4)) = v2124
	v2126 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2122)+8)) = v2126
	v2128 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2122)+16)) = v2128
	v2130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2122)+24)) = v2130
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+32)) = v2133
	v2135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2122)+36)) = uint8(v2135)
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2122)+37)) = uint8(v2137)
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+88)) = v2139
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+92)) = v2141
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+4))
	if v2143 == int32(4) {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2151 = l0 + int32(60)
	goto L528
L527:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2151 = v2148 + int32(8)
	goto L528
L528:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2151)))
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+112)) = v2152
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2155 = F_bms_difference(m, v2152, v2154)
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L8
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+116)) = v2155
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+172)))
	if v2158 == int32(1) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2161)+93)) = uint8(v2162)
	goto L532
L531:
	;
	goto L532
L532:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2164 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2122)+48))
	v2166 = F_replace_nestloop_params_mutator(m, v2165, l0)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L8
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	v2177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2122)+120)) = uint8(v2177)
	if v2093 == v2177 {
		v2385 = v2122
		goto L131
	} else {
		goto L539
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+48)) = v2166
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2122)+96))
	v2170 = F_replace_nestloop_params_mutator(m, v2169, l0)
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L8
	} else {
		goto L537
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+96)) = v2170
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2122)+108))
	v2174 = F_replace_nestloop_params_mutator(m, v2173, l0)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L8
	} else {
		goto L538
	}
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2122)+108)) = v2174
	goto L535
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = int32(0)
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+40))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+4))
	F_pull_varattnos(m, v2184, v2093, v19+int32(172))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L8
	} else {
		goto L540
	}
L540:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+204))
	if v2189 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2244 = F_bms_is_member(m, int32(1), v2243)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L8
	} else {
		goto L550
	}
L542:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2189)+4))
	if v2192 <= int32(0) {
		goto L541
	} else {
		goto L543
	}
L543:
	;
	v2198 = int32(0)
	goto L544
L544:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2189)+12))
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2212+v2198<<(uint(int32(2))%32))))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2216)+4))
	F_pull_varattnos(m, v2217, v2093, v19+int32(172))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L8
	} else {
		goto L546
	}
L545:
	;
	goto L541
L546:
	;
	v2223 = v2198 + int32(1)
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2189)+4))
	if v2223 < v2224 {
		v2198 = v2223
		goto L544
	} else {
		goto L547
	}
L547:
	;
	goto L545
L548:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	F_bms_free(m, v2270)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L8
	} else {
		goto L562
	}
L549:
	;
	v2268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2122)+120)) = uint8(v2268)
	goto L548
L550:
	;
	if v2244 != 0 {
		goto L549
	} else {
		goto L551
	}
L551:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2248 = F_bms_is_member(m, int32(2), v2247)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L8
	} else {
		goto L552
	}
L552:
	;
	if v2248 != 0 {
		goto L549
	} else {
		goto L553
	}
L553:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2252 = F_bms_is_member(m, int32(3), v2251)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L8
	} else {
		goto L554
	}
L554:
	;
	if v2252 != 0 {
		goto L549
	} else {
		goto L555
	}
L555:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2256 = F_bms_is_member(m, int32(4), v2255)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L8
	} else {
		goto L556
	}
L556:
	;
	if v2256 != 0 {
		goto L549
	} else {
		goto L557
	}
L557:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2260 = F_bms_is_member(m, int32(5), v2259)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L8
	} else {
		goto L558
	}
L558:
	;
	if v2260 != 0 {
		goto L549
	} else {
		goto L559
	}
L559:
	;
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2264 = F_bms_is_member(m, int32(6), v2263)
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L8
	} else {
		goto L560
	}
L560:
	;
	if v2264 == int32(0) {
		goto L548
	} else {
		goto L561
	}
L561:
	;
	goto L549
L562:
	;
	v2385 = v2122
	goto L131
L563:
	;
	v2330 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L8
	} else {
		goto L575
	}
L564:
	;
	v2316 = int32(0)
	goto L563
L565:
	;
	goto L566
L566:
	;
	v2278 = int32(0)
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+4))
	if v2279 <= v2278 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v2316 = int32(0)
	goto L563
L568:
	;
	goto L569
L569:
	;
	v2286 = int32(0)
	v2287 = v2278
	goto L570
L570:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+12))
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2300+v2287<<(uint(int32(2))%32))))
	v2306 = F_create_plan_recurse(m, l0, v2304, int32(1))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L8
	} else {
		goto L572
	}
L571:
	;
	v2316 = v2308
	goto L563
L572:
	;
	v2308 = F_lappend(m, v2286, v2306)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L8
	} else {
		goto L573
	}
L573:
	;
	v2311 = v2287 + int32(1)
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+4))
	if v2311 < v2312 {
		v2286 = v2308
		v2287 = v2311
		goto L570
	} else {
		goto L574
	}
L574:
	;
	goto L571
L575:
	;
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+4))
	v2334 = m.T0[v2333].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2273, l1, v458, v2330, v2316)
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L8
	} else {
		goto L576
	}
L576:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+4)) = v2336
	v2338 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2334)+8)) = v2338
	v2340 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2334)+16)) = v2340
	v2342 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2334)+24)) = v2342
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+32)) = v2345
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2334)+36)) = uint8(v2347)
	v2349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2334)+37)) = uint8(v2349)
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2351)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+100)) = v2352
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2354 == int32(0) {
		v2385 = v2334
		goto L131
	} else {
		goto L577
	}
L577:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+48))
	v2358 = F_replace_nestloop_params_mutator(m, v2357, l0)
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L8
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+48)) = v2358
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+88))
	v2362 = F_replace_nestloop_params_mutator(m, v2361, l0)
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L8
	} else {
		goto L579
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+88)) = v2362
	v2385 = v2334
	goto L131
L580:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2369
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_3), v19)
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L8
	} else {
		goto L581
	}
L581:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(794), int32(_a_F_create_scan_plan_15))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L8
	} else {
		goto L582
	}
L582:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L583:
	;
	v2385 = v2380
	goto L131
L584:
	;
	v2398 = F_create_gating_plan(m, l0, l1, v2385, v65)
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L8
	} else {
		goto L587
	}
L585:
	;
	v2400 = v2385
	goto L586
L586:
	;
	m.G0 = v19 + int32(176)
	return v2400
L587:
	;
	v2400 = v2398
	goto L586
L588:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v2409
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_16), v19+int32(48))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L8
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3777), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L8
	} else {
		goto L590
	}
L590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L591:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v2425
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_17), v19-int32(-64))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L8
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3780), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L8
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v2457
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_16), v19+int32(80))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L8
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3788), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L8
	} else {
		goto L596
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2473
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(112))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L8
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3922), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L8
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L600:
	;
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2489
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_18), v19+int32(128))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L8
	} else {
		goto L601
	}
L601:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3932), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2500 = m.ExcPending
	if v2500 != 0 {
		goto L8
	} else {
		goto L602
	}
L602:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	v11 = F_palloc0(m, int32(80))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1593432867124)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v20
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v24 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v29 = v27
		} else {
			v29 = int32(0)
		}
		v31 = v29 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v31)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v39
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+48)) = v41
		v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v46 = *(*float64)(unsafe.Add(mBase, _c_F_create_unique_path[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l3
		*(*float64)(unsafe.Add(mBase, uint32(v11)+56)) = base.F64_add(v43, base.F64_mul(base.F64_mul(v46, v44), base.F64_convert_i32_s(l2)))
		return v11
	}
}
func F_crosstab_hash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int64
	_ = v176
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v277 int64
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v470 int32
	_ = v470
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int64
	_ = v531
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	v25 = m.G0
	v27 = v25 - int32(160)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = F_pg_detoast_datum_packed(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v34 = F_text_to_cstring(m, v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v37 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = F_text_to_cstring(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v41 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L153
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L143
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L139
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L136
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L132
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L128
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L123
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L118
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L114
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L110
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v44 != int32(389) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	if v47&int32(2) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v52 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v55 = int32(_a_F_crosstab_hash_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v59
	v61 = F_CreateTupleDescCopy(m, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v63 <= int32(1) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = int64(292057776192)
	v74 = F_hash_create(m, int32(_a_F_crosstab_hash_1), int64(64), v27+int32(48), int32(1048))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v81 = F_SPI_execute(m, v39, int32(1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v202 = F_SPI_finish(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L39
	}
L26:
	;
	if v81 != int32(5) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab_hash[1]))
	if v86 == int64(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[2]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v92 != int32(1) {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v117 = int64(0)
	goto L30
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119+base.I32_wrap_i64(v117)<<(uint(int32(2))%32))))
	v126 = F_SPI_getvalue(m, v124, v91, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	if v126 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v130 = int32(_a_F_crosstab_hash_0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v59
	v135 = F_palloc(m, int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v126
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+152)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+144)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+136)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+128)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+120)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+104)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v140
	v157 = v27 + int32(96)
	v162 = F_pg_snprintf(m, v157, int32(63), int32(_a_F_crosstab_hash_2), v27+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v167 = F_hash_search(m, v74, v157, int32(1), v27+int32(47))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+47)))
	if v169 == int32(1) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v135
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v131
	v176 = v117 + int64(1)
	if v176 != v86 {
		v117 = v176
		goto L30
	} else {
		goto L38
	}
L38:
	;
	goto L31
L39:
	;
	if v202 != int32(2) {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = int32(2)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v210)+8))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v210)+808))
	if v212 != int64(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v278 = F_TupleDescGetAttInMetadata(m, v61)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v210)+752))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v210)+728))
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v210)+704))
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v210)+680))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v210)+656))
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v210)+632))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v210)+608))
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v210)+584))
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v210)+560))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v210)+536))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v210)+512))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v210)+488))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v210)+464))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v210)+440))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v210)+416))
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v210)+392))
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v210)+368))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v210)+344))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v210)+320))
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v210)+296))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v210)+272))
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v210)+248))
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v210)+224))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v210)+200))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v210)+176))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v210)+152))
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v210)+128))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v210)+104))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v210)+80))
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v210)+56))
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v210)+32))
	v277 = v215 + (v216 + (v217 + (v218 + (v219 + (v220 + (v221 + (v222 + (v223 + (v224 + (v225 + (v226 + (v227 + (v228 + (v229 + (v230 + (v231 + (v232 + (v233 + (v234 + (v235 + (v236 + (v237 + (v238 + (v239 + (v240 + (v241 + (v242 + (v243 + (v244 + (v245 + v211))))))))))))))))))))))))))))))
	goto L44
L43:
	;
	v277 = v211
	goto L44
L44:
	;
	goto L41
L45:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[3]))
	v287 = F_tuplestore_begin_heap(m, int32(base.Ui32(v208&int32(4))>>(uint(int32(2))%32)), int32(0), v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v294 = F_SPI_execute(m, v34, int32(1), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v612 = F_SPI_finish(m)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L108
	}
L49:
	;
	if v294 != int32(5) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v299 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab_hash[1]))
	if v299 == int64(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v302 = base.I32_wrap_i64(v277)
	if v302 == int32(0) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[2]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v308 <= int32(2) {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v312 = v308 - int32(2)
	v313 = v312 + v302
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v313 != v314 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v318 = F_palloc0(m, v313<<(uint(int32(2))%32))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v323 = int32(1)
	v329 = v323
	v335 = int32(0)
	v351 = int64(0)
	goto L56
L56:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353+base.I32_wrap_i64(v351)<<(uint(int32(2))%32))))
	v360 = F_SPI_getvalue(m, v358, v307, int32(1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v584 = F_BuildTupleFromCStrings(m, v278, v318)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L106
	}
L58:
	;
	if v329&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v527 = F_SPI_getvalue(m, v358, v307, v308-v323)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L90
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v360
	if v308 == int32(3) {
		goto L59
	} else {
		goto L84
	}
L61:
	;
	if v360|v335 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v367 = int32(0)
	if base.B2i32(v335 == v367)|base.B2i32(v360 == v367) == v367 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if base.B2i32(v376 == int32(0))|base.B2i32(v376 != v379) != 0 {
		v397 = v376
		v398 = v379
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	v402 = F_BuildTupleFromCStrings(m, v278, v318)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	if v397-v398 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	v382 = v335
	v383 = v360
	goto L69
L69:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+1)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+1)))
	if v387 == int32(0) {
		v397 = v387
		v398 = v386
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v397 = v387
	v398 = v386
	goto L67
L71:
	;
	v390 = int32(1)
	if v387 == v386 {
		v382 = v382 + v390
		v383 = v383 + v390
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	goto L65
L74:
	;
	F_tuplestore_puttuple(m, v287, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v406 = int32(0)
	if v313 <= v406 {
		goto L60
	} else {
		goto L76
	}
L76:
	;
	v409 = v406
	goto L77
L77:
	;
	v435 = v318 + v409<<(uint(int32(2))%32)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	if v436 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L60
L79:
	;
	F_pfree(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v442 = v409 + int32(1)
	if v442 != v313 {
		v409 = v442
		goto L77
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = int32(0)
	goto L81
L83:
	;
	goto L78
L84:
	;
	v470 = int32(1)
	goto L85
L85:
	;
	v498 = v470 + int32(1)
	v499 = F_SPI_getvalue(m, v358, v307, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L59
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318+v470<<(uint(int32(2))%32)))) = v499
	if v498 != v312 {
		v470 = v498
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v335 != 0 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	if v527 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v531 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+152)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+144)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+136)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+128)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+120)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+104)) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v527
	v549 = v27 + int32(96)
	v552 = F_pg_snprintf(m, v549, int32(63), int32(_a_F_crosstab_hash_2), v27)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v554 = int32(0)
	v556 = F_hash_search(m, v74, v549, v554, v554)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v556 == int32(0) {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556)+64))
	if v560 == int32(0) {
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v563 = F_SPI_getvalue(m, v358, v307, v308)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v560)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v318+v308<<(uint(int32(2))%32)+v565<<(uint(int32(2))%32)-int32(8)))) = v563
	goto L89
L97:
	;
	F_pfree(m, v335)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v576 = int32(0)
	if v360 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v578 = F_pstrdup(m, v360)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	v580 = v576
	goto L103
L103:
	;
	v582 = v351 + int64(1)
	if v582 != v299 {
		v329 = v576
		v335 = v580
		v351 = v582
		goto L56
	} else {
		goto L105
	}
L104:
	;
	v580 = v578
	goto L103
L105:
	;
	goto L57
L106:
	;
	F_tuplestore_puttuple(m, v287, v584)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L48
L108:
	;
	if v612 != int32(2) {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v287
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v56
	m.G0 = v27 + int32(160)
	return int64(0)
L110:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_3), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(652), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_6), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(657), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_7), int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v671 = F_errdetail(m, int32(_a_F_crosstab_hash_8), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(676), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_9), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v691 = F_errdetail(m, int32(_a_F_crosstab_hash_10), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(751), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_12), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(767), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(_a_F_crosstab_hash_13))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_14), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(776), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(_a_F_crosstab_hash_15), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(784), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_16), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(834), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_18), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v772 = F_errdetail(m, int32(_a_F_crosstab_hash_19), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(854), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_7), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v313
	v796 = F_errdetail(m, int32(_a_F_crosstab_hash_20), v27+int32(16))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(864), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_crosstab_hash_21), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(934), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
