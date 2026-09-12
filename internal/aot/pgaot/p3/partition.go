package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DetachPartitionFinalize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
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
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
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
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int64
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
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
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int64
	_ = v662
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
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
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(656)
	m.G0 = v18
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_RemoveInheritance(m, l1, l0, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v24 = F_new_object_addresses(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_ScanKeyInit(m, v18+int32(512), int32(2), int32(3), int32(184), v23)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v38 = int32(1)
	v43 = F_systable_beginscan(m, v35, int32(2701), v38, int32(0), v38, v18+int32(512))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = v45
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L24
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v64 = v62 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v65 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v92 = F_systable_getnext(m, v43)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L22
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
	if v68 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v69 = int32(2620)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v73 = F_deleteDependencyRecordsForClass(m, v69, v70, v69, int32(80))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v79 = F_deleteDependencyRecordsForClass(m, int32(2620), v76, int32(1259), int32(83))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+240)) = int32(2620)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+244)) = v83
	F_add_exact_object_address(m, v18+int32(240), v24)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L16
L22:
	;
	if v92 != 0 {
		v49 = v92
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	F_performMultipleDeletions(m, v24, int32(0), int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_free_object_addresses(m, v24)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_sequence_close(m, v35, int32(3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v122 = F_RelationGetFKeyList(m, l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v496 = F_GetParentedForeignKeyRefs(m, l1)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L100
	}
L30:
	;
	v124 = F_copyObjectImpl(m, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v124 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_list_free_deep(m, v124)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v132 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if int32(0) < v134 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = int32(0)
	v151 = v5
	goto L40
L38:
	;
	v178 = v5
	goto L39
L39:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if int32(0) < v180 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v140<<(uint(int32(2))%32))))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v159 = F_lappend_oid(m, v151, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	v178 = v159
	goto L39
L42:
	;
	v162 = v140 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v162 < v163 {
		v140 = v162
		v151 = v159
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v197 = v5
	goto L47
L45:
	;
	goto L46
L46:
	;
	F_list_free_deep(m, v124)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L96
	}
L47:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(int32(2))%32))))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v205 = F_SearchSysCache1(m, int32(19), v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L52
	}
L48:
	;
	goto L46
L49:
	;
	F_ReleaseCatCache(m, v205)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L94
	}
L50:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v412 = F_table_open(m, v410, int32(6))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L91
	}
L51:
	;
	v357 = v319
	v360 = v319
	goto L86
L52:
	;
	if v205 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+22)))
	v209 = v207 + v208
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+72)))
	if v210 != int32(102) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L83
	}
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+92))
	if v213 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v216 = int32(0)
	if v178 == v216 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v254 != 0 {
		goto L49
	} else {
		goto L71
	}
L59:
	;
	v254 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v222 <= int32(0) {
		v247 = v216
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v254 = v247
	goto L58
L63:
	;
	v225 = int32(0)
	if v225 < v222 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v228 = v222
	goto L66
L65:
	;
	v228 = v225
	goto L66
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v231 = int32(0)
	goto L67
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229+v231<<(uint(int32(2))%32))))
	v240 = base.B2i32(v239 == v213)
	if v239 == v213 {
		v247 = v240
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v247 = v240
	goto L62
L69:
	;
	v242 = v231 + int32(1)
	if v242 != v228 {
		v231 = v242
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v256 = int32(0)
	F_ConstraintSetParentConstraint(m, v255, v256, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+20)))
	if v260 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	F_GetForeignKeyCheckTriggers(m, v132, v263, v264, v265, v18+int32(512), v18+int32(240))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_DeconstructFkConstraintRow(m, v205, v18+int32(508), v18+int32(432), v18+int32(368), v18+int32(512), v18+int32(240), v18+int32(112), v18+int32(108), v18+int32(32))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L79
	}
L76:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v18)+512))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v132, v272, int32(0), v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v18)+240))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v132, v277, int32(0), v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v301 = F_palloc0(m, int32(108))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v301))) = int64(438086664353)
	v307 = F_pstrdup(m, v209+int32(4))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+8)) = v307
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+12)) = uint8(v310)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+13)) = uint8(v312)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+75)))
	v315 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+15)) = uint8(v315)
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+14)) = uint8(v314)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+76)))
	v319 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+80)) = v319
	v322 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v301)+72)) = v322
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+16)) = uint8(v318)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+86)) = uint8(v325)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+87)) = uint8(v327)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+101)))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+100)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v301)+92)) = v322
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+88)) = uint8(v329)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	if v319 < v336 {
		goto L51
	} else {
		goto L82
	}
L82:
	;
	goto L50
L83:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v343
	F_errmsg_internal(m, int32(40426), v18+int32(16))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(485619), int32(21154), int32(335494))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v380 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18+int32(432)+v357<<(uint(int32(1))%32)))))
	v386 = F_makeString(m, v370+v371<<(uint(int32(4))%32)+v380*int32(100)-int32(76))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L50
L88:
	;
	v388 = F_lappend(m, v360, v386)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+76)) = v388
	v392 = v357 + int32(1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	if v392 < v393 {
		v357 = v392
		v360 = v388
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v209)+88))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	v430 = int32(0)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+107)))
	F_addFkRecurseReferenced(m, v301, l1, v412, v414, v415, v416, v18+int32(368), v18+int32(432), v18+int32(512), v18+int32(240), v18+int32(112), v427, v18+int32(32), v430, v430, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_sequence_close(m, v412, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v456 = v197 + int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v456 < v457 {
		v197 = v456
		goto L47
	} else {
		goto L95
	}
L95:
	;
	goto L48
L96:
	;
	if v132 == int32(0) {
		goto L29
	} else {
		goto L97
	}
L97:
	;
	F_sequence_close(m, v132, int32(3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	goto L29
L99:
	;
	v565 = F_RelationGetIndexList(m, l1)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L111
	}
L100:
	;
	if v496 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v500 = int32(0)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v501 <= v500 {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v506 = v500
	goto L103
L103:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v496)+12))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v519+v506<<(uint(int32(2))%32))))
	v524 = int32(0)
	F_ConstraintSetParentConstraint(m, v523, v524, v524)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L105
	}
L104:
	;
	goto L99
L105:
	;
	v528 = int32(2606)
	v531 = F_deleteDependencyRecordsForClass(m, v528, v523, v528, int32(105))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v535 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+520)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v18)+516)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v18)+512)) = int32(2606)
	F_performDeletion(m, v18+int32(512), v535, v535)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v547 = v506 + int32(1)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v547 < v548 {
		v506 = v547
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	v645 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L131
	}
L111:
	;
	if v565 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v569 = int32(0)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	if v570 <= v569 {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v575 = v569
	goto L114
L114:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v565)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v588+v575<<(uint(int32(2))%32))))
	v593 = F_has_superclass(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	goto L110
L116:
	;
	if v593 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v596 = F_get_partition_parent(m, v592, int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v625 = v575 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	if v625 < v626 {
		v575 = v625
		goto L114
	} else {
		goto L130
	}
L120:
	;
	v599 = F_index_open(m, v592, int32(8))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	F_IndexSetParentIndex(m, v599, int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v605 = F_get_relation_idx_constraint_oid(m, v604, v592)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v608 = F_get_relation_idx_constraint_oid(m, v607, v596)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L4
	} else {
		goto L125
	}
L124:
	;
	F_relation_close(m, v599, int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L4
	} else {
		goto L129
	}
L125:
	;
	if v608 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v605 == int32(0) {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v614 = int32(0)
	F_ConstraintSetParentConstraint(m, v605, v614, v614)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	goto L124
L129:
	;
	goto L119
L130:
	;
	goto L115
L131:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v650 = F_SearchSysCacheCopy(m, int32(57), v648, int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v650 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v657 = F__emscripten_memset_bulkmem(m, v18+int32(512), base.I32_extend8_s(int32(0)), int32(136))
	mBase = m.M
	goto L136
L134:
	;
	goto L135
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L4
	} else {
		goto L168
	}
L136:
	;
	v658 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+144)) = uint16(v658)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+272)) = uint16(v658)
	v662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+264)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+256)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+248)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+240)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v662
	v678 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v678)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+273)) = uint8(v678)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v645)+52))
	v689 = F_heap_modify_tuple(m, v650, v682, v18+int32(512), v18+int32(240), v18+int32(112))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v689)+16))
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+22)))
	v694 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v691+v692)+131)) = uint8(v694)
	F_CatalogTupleUpdate(m, v645, v689+int32(4), v689)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_pfree(m, v689)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_sequence_close(m, v645, int32(3))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(v705)+120)))
	if int32(0) < v706 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v712 = int32(0)
	v717 = v705
	goto L144
L142:
	;
	goto L143
L143:
	;
	if l3 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L144:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	v732 = v725 + v726<<(uint(int32(4))%32) + v712*int32(100)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732)+111)))
	if v733 != 0 {
		v751 = v717
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L143
L146:
	;
	v753 = v712 + int32(1)
	v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v751)+120)))
	if v753 < v754 {
		v712 = v753
		v717 = v751
		goto L144
	} else {
		goto L150
	}
L147:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+int32(20))+89)))
	if v736 == int32(0) {
		v751 = v717
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v745 = int32(1)
	F_ATExecDropIdentity(m, v18+int32(432), l1, v732+int32(24), int32(0), int32(8), v745, v745)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v751 = v749
	goto L146
L150:
	;
	goto L145
L151:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L158
	}
L152:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l3 == v773 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_update_default_partition_oid(m, v775, int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L4
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	F_CacheInvalidateRelcacheByRelid(m, l3)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L157
	}
L156:
	;
	goto L151
L157:
	;
	goto L151
L158:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+119)))
	if v784 != int32(112) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	m.G0 = v18 + int32(656)
	return
L160:
	;
	v787 = int32(0)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v791 = F_find_all_inheritors(m, v788, int32(8), v787)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v791 == int32(0) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	if v795 <= int32(0) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v800 = v787
	goto L164
L164:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v791)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813+v800<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	goto L159
L166:
	;
	v821 = v800 + int32(1)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	if v821 < v822 {
		v800 = v821
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v846
	F_errmsg_internal(m, int32(45662), v18)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(485619), int32(21350), int32(335494))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v5 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v9 == v5 {
		v12 = int32(4470400)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = F_RelationGetPartitionQual(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(4470400)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
			v27 = F_expression_planner(m, v18)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					v29 = F_make_ands_explicit(m, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v32 = F_ExecInitExpr(m, v29, int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(4470400)
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
							*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v34
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
							if v44 == int32(0) {
								v47 = F_MakePerTupleExprContext(m, l2)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = v47
									*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									v52 = F_ExecCheck(m, v51, v49)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										if v52 != 0 {
											return v52
										} else {
											if l3 == int32(0) {
												return v52
											} else {
												F_ExecPartitionCheckEmitError(m, l0, l1, l2)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
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
							} else {
								v49 = v44
								*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								v52 = F_ExecCheck(m, v51, v49)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return v52
									} else {
										if l3 == int32(0) {
											return v52
										} else {
											F_ExecPartitionCheckEmitError(m, l0, l1, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
				} else {
					v34 = v5
					v35 = int32(4470400)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v34
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
					if v44 == int32(0) {
						v47 = F_MakePerTupleExprContext(m, l2)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = v47
							*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
							v52 = F_ExecCheck(m, v51, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return v52
								} else {
									if l3 == int32(0) {
										return v52
									} else {
										F_ExecPartitionCheckEmitError(m, l0, l1, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
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
					} else {
						v49 = v44
						*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
						v52 = F_ExecCheck(m, v51, v49)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return v52
							} else {
								if l3 == int32(0) {
									return v52
								} else {
									F_ExecPartitionCheckEmitError(m, l0, l1, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
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
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
		if v44 == int32(0) {
			v47 = F_MakePerTupleExprContext(m, l2)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = v47
				*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				v52 = F_ExecCheck(m, v51, v49)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return v52
					} else {
						if l3 == int32(0) {
							return v52
						} else {
							F_ExecPartitionCheckEmitError(m, l0, l1, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
		} else {
			v49 = v44
			*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
			v52 = F_ExecCheck(m, v51, v49)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v52 != 0 {
					return v52
				} else {
					if l3 == int32(0) {
						return v52
					} else {
						F_ExecPartitionCheckEmitError(m, l0, l1, l2)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
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
func F_InitPartitionPruneContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	v7 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = v12
	goto L3
L2:
	;
	v13 = v7
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v14)
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v24
	v26 = v13 * v16
	v29 = F_palloc0(m, v26*int32(28))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v29
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v33
	v39 = F_palloc0(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v39
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v44 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v56 = v7
	goto L10
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v56<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v64 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = v65
	goto L14
L13:
	;
	v66 = int32(0)
	goto L14
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v67 != int32(377) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v143 = v56 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v143 < v144 {
		v56 = v143
		goto L10
	} else {
		goto L37
	}
L16:
	;
	if v16 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v75 = v66
	v76 = int32(0)
	goto L18
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v85 = F_bms_is_member(m, v76, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L15
L20:
	;
	v129 = v76 + int32(1)
	if v129 != v16 {
		v75 = v125
		v76 = v129
		goto L18
	} else {
		goto L36
	}
L21:
	;
	if v85 != 0 {
		v125 = v75
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v75 == int32(0) {
		v125 = v75
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 != int32(7) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if l4 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v115 = v75 + int32(4)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if base.Ui32(v115) < base.Ui32(v118+v119<<(uint(int32(2))%32)) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v105 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v104+v93*v16<<(uint(v105)%32)+v76<<(uint(v105)%32)))) = v103
	goto L26
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v98 = F_ExecInitExprWithParams(m, v89, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v101 = F_ExecInitExpr(m, v89, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v103 = v98
	goto L27
L32:
	;
	v103 = v101
	goto L27
L33:
	;
	v124 = v115
	goto L35
L34:
	;
	v124 = int32(0)
	goto L35
L35:
	;
	v125 = v124
	goto L20
L36:
	;
	goto L19
L37:
	;
	goto L11
}
func F_build_partition_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	v5 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+2)))
	if v14 <= v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v194)
	return v195
L2:
	;
	v194 = int32(0)
	v195 = v182
	goto L1
L3:
	;
	v182 = v5
	goto L2
L4:
	;
	goto L5
L5:
	;
	v18 = base.B2i32(l2 == int32(-1))
	v24 = v5
	v29 = v5
	goto L6
L6:
	;
	v32 = v29 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41+v32)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44+v32)))
	v47 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v50 = F_make_pathkey_from_sortinfo(m, l0, v37, v40, v43, v46, v18, v18, v47, v48, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v182 = v166
	goto L2
L8:
	;
	v174 = v29 + int32(1)
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+2)))
	if v174 < v175 {
		v24 = v166
		v29 = v174
		goto L6
	} else {
		goto L38
	}
L9:
	;
	v159 = F_lappend(m, v24, v50)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L37
	}
L10:
	;
	return int32(0)
L11:
	;
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+40)))
	if v55 != 0 {
		v166 = v24
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v84 = int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86+v32)))
	if base.B2i32(v88 != int32(2222))&base.B2i32(v88 != int32(424)) != 0 {
		v194 = v84
		v195 = v24
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v24 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v58 <= int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v65 = int32(0)
	goto L18
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v61+v65<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v54 == v79 {
		v166 = v24
		goto L8
	} else {
		goto L20
	}
L19:
	;
	goto L9
L20:
	;
	v82 = v65 + int32(1)
	if v82 != v58 {
		v65 = v82
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v94 == int32(0) {
		v194 = v84
		v195 = v24
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v98 <= v97 {
		v194 = v84
		v195 = v24
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v103 = v97
	goto L25
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v103<<(uint(int32(2))%32))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v118 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v194 = v142
	v195 = v24
	goto L1
L27:
	;
	v142 = int32(1)
	v144 = v103 + v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v144 < v145 {
		v103 = v144
		goto L25
	} else {
		goto L36
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119+v32)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v125 = F_equal(m, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v125 != 0 {
		v166 = v24
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if v124 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v129 != int32(21) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v132 != int32(2) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v138 = F_equal(m, v123, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v138 != 0 {
		v166 = v24
		goto L8
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L26
L37:
	;
	v166 = v159
	goto L8
L38:
	;
	goto L7
}
func F_check_new_partition_bound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v427 int32
	_ = v427
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
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
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
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
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v763 int32
	_ = v763
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v791 int32
	_ = v791
	var v804 int32
	_ = v804
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v988 int32
	_ = v988
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1038 int32
	_ = v1038
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1142 int32
	_ = v1142
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1370 int32
	_ = v1370
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(112)
	m.G0 = v30
	v32 = F_RelationGetPartitionKey(m, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = F_RelationGetPartitionDesc(m, l1, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v38 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v1370 + int32(112)
	return
L5:
	;
	if v37 == int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	switch v74 - int32(104) {
	case 0:
		goto L19
	default:
		v1370 = v30
		goto L4
	case 4:
		goto L17
	case 10:
		goto L18
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	if v43 == int32(-1) {
		v1370 = v30
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54<<(uint(int32(2))%32))))
	v59 = F_get_rel_name(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = l0
	F_errmsg(m, int32(677215), v30)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_parser_errposition(m, l3, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(485356), int32(2922), int32(417024))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L200
	}
L17:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v1153 <= int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L174
	}
L18:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v318 = F_make_one_partition_rbound(m, v32, int32(-1), v316, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L75
	}
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v77 <= int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v91 = int32(-1)
	v92 = v80 - int32(1)
	goto L21
L21:
	;
	if v91 < v92 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v151 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	v114 = int32(1)
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v120 = int32(2)
	v121 = base.I32_div_s(v91+v92+v114, v120)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116+v121<<(uint(v120)%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126 < v84 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v151 = v91
	goto L25
L25:
	;
	goto L22
L26:
	;
	if v144 == int32(0) {
		v91 = v141
		v92 = v142
		goto L21
	} else {
		goto L36
	}
L27:
	;
	v141 = v121
	v142 = v92
	v144 = v114
	v145 = v115
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v84 < v126 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v141 = v91
	v142 = v121 - int32(1)
	v144 = int32(0)
	v145 = v115
	goto L26
L31:
	;
	goto L32
L32:
	;
	v132 = base.B2i32(v84 != v126)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v135 = v132 | base.B2i32(v133 == v83)
	if v132|base.B2i32(v133 <= v83) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = v121
	v142 = v92
	v144 = v114
	v145 = v135
	goto L26
L34:
	;
	goto L35
L35:
	;
	v141 = v91
	v142 = v121 - int32(1)
	v144 = int32(0)
	v145 = v135
	goto L26
L36:
	;
	if v145 == int32(0) {
		v91 = v141
		v92 = v142
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v151 = v141
	goto L25
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L69
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L63
	}
L40:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v214 <= v83 {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v163 = base.I32_rem_s(v162, v84)
	if v163 == int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v196 = v151 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v158+v196)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = base.I32_rem_s(v84, v199)
	if v200 != 0 {
		goto L39
	} else {
		goto L51
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(112926), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = F_get_rel_name(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v177
	F_errdetail(m, int32(640323), v30+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(485356), int32(2975), int32(417024))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v202 = v151 + int32(1)
	if v80 <= v202 {
		goto L40
	} else {
		goto L52
	}
L52:
	;
	v205 = v202 << (uint(int32(2)) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v158+v205)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = base.I32_rem_s(v208, v84)
	if v209 != 0 {
		goto L38
	} else {
		goto L53
	}
L53:
	;
	goto L40
L54:
	;
	v216 = base.I32_rem_s(v83, v214)
	v217 = v216
	goto L56
L55:
	;
	v217 = v83
	goto L56
L56:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v220 = v217
	goto L57
L57:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v218+v220<<(uint(int32(2))%32))))
	if v249 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v1370 = v30
	goto L4
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1307 = l0
	v1310 = l3
	v1313 = v249
	v1315 = v30
	v1322 = v252
	v1326 = v35
	goto L16
L60:
	;
	goto L61
L61:
	;
	v253 = v220 + v84
	if v253 < v214 {
		v220 = v253
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(112926), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267+v196)))
	v270 = F_get_rel_name(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v266
	F_errdetail(m, int32(640240), v30-int32(-64))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(485356), int32(2995), int32(417024))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(112926), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v297+v205)))
	v300 = F_get_rel_name(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v296
	F_errdetail(m, int32(640323), v30+int32(48))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(485356), int32(3016), int32(417024))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
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
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v323 = F_make_one_partition_rbound(m, v32, int32(-1), v321, int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+12)))
	v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v326 <= int32(0) {
		v396 = v5
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v517 <= int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L104
	}
L78:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v449+v427<<(uint(int32(2))%32)-int32(4))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L96
	}
L79:
	;
	if v325&int32(255) != 0 {
		v495 = v5
		goto L77
	} else {
		goto L94
	}
L80:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v339 = v5
	goto L81
L81:
	;
	v363 = v339 << (uint(int32(2)) % 32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v331+v363)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v329)))
	if v365 < v367 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v396 = v326
	goto L79
L83:
	;
	v495 = v339 ^ int32(-1)
	goto L77
L84:
	;
	goto L85
L85:
	;
	v372 = v339 + int32(1)
	if v367 < v365 {
		v427 = v372
		goto L78
	} else {
		goto L86
	}
L86:
	;
	if v365 != 0 {
		v396 = v372
		goto L79
	} else {
		goto L87
	}
L87:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v363+v333)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v363+v332)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+v330)))
	v383 = F_FunctionCall2Coll(m, v334+v339*int32(28), v378, v380, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v383 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if int32(0) <= v383 {
		v427 = v372
		goto L78
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v326 != v372 {
		v339 = v372
		goto L81
	} else {
		goto L93
	}
L92:
	;
	v495 = v339 ^ int32(-1)
	goto L77
L93:
	;
	goto L82
L94:
	;
	if v396 <= int32(0) {
		v495 = v396
		goto L77
	} else {
		goto L95
	}
L95:
	;
	v427 = v396
	goto L78
L96:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = l0
	F_errmsg(m, int32(677360), v30+int32(96))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v470 = F_get_range_partbound_string(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v473 = F_get_range_partbound_string(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v470
	F_errdetail(m, int32(580597), v30+int32(80))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	F_parser_errposition(m, l3, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(485356), int32(3132), int32(417024))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v520 = int32(-1)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v523 = v521 - int32(1)
	if int32(0) <= v523 {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1307 = v1123
	v1310 = v1126
	v1313 = v1129
	v1315 = v1131
	v1322 = v1152
	v1326 = v1142
	goto L16
L106:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+12))
	v1115 = v791 >> (uint(int32(31)) % 32)
	v1123 = l0
	v1126 = l3
	v1129 = v819
	v1131 = v30
	v1142 = v35
	v1150 = v1113 + (v791^v1115-v1115)<<(uint(int32(2))%32) - int32(4)
	goto L105
L107:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+12))
	v1123 = v1083
	v1126 = v1086
	v1129 = v1089
	v1131 = v1091
	v1142 = v1102
	v1150 = v1111
	goto L105
L108:
	;
	if v791 != 0 {
		goto L106
	} else {
		goto L173
	}
L109:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v871 <= v845 {
		v1370 = v852
		goto L4
	} else {
		goto L146
	}
L110:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v837 = v562 + int32(1)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v835+v837<<(uint(int32(2))%32))))
	if int32(0) <= v841 {
		v1083 = l0
		v1086 = l3
		v1089 = v841
		v1091 = v30
		v1102 = v35
		goto L107
	} else {
		goto L145
	}
L111:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v826 = v562 + int32(1)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v824+v826<<(uint(int32(2))%32))))
	if int32(0) <= v830 {
		v1083 = l0
		v1086 = l3
		v1089 = v830
		v1091 = v30
		v1102 = v35
		goto L107
	} else {
		goto L144
	}
L112:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v528 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	v530 = v37 + int32(24)
	v549 = v520
	v553 = v523
	goto L115
L113:
	;
	v791 = v495
	v804 = v520
	goto L114
L114:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v815 = v804 + int32(1)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v813+v815<<(uint(int32(2))%32))))
	if int32(0) <= v819 {
		goto L108
	} else {
		goto L143
	}
L115:
	;
	v561 = int32(2)
	v562 = base.I32_div_s(v549+v553+int32(1), v561)
	v564 = v562 << (uint(v561) % 32)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v564+v565)))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+12)))
	v569 = int32(0)
	if v528 <= v569 {
		v658 = v569
		goto L121
	} else {
		goto L122
	}
L116:
	;
	v791 = v763
	v804 = v776
	goto L114
L117:
	;
	if v776 < v780 {
		v549 = v776
		v553 = v780
		goto L115
	} else {
		goto L142
	}
L118:
	;
	v763 = v734
	v776 = v549
	v780 = v562 - int32(1)
	goto L117
L119:
	;
	if int32(0) < v724 {
		v734 = v724
		goto L118
	} else {
		goto L140
	}
L120:
	;
	v692 = int32(0)
	if v666 < v692 {
		goto L137
	} else {
		goto L138
	}
L121:
	;
	v660 = base.B2i32(v567 == int32(-1))
	if v568 == v660 {
		goto L110
	} else {
		goto L133
	}
L122:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v572+v564)))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v575+v564)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v584 = v569
	goto L123
L123:
	;
	v610 = v584 << (uint(int32(2)) % 32)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v574+v610)))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610+v578)))
	if v612 < v614 {
		v724 = v584 ^ int32(-1)
		goto L119
	} else {
		goto L125
	}
L124:
	;
	v658 = v528
	goto L121
L125:
	;
	v617 = v584 + int32(1)
	if v614 < v612 {
		v734 = v617
		goto L118
	} else {
		goto L126
	}
L126:
	;
	if v612 != 0 {
		v658 = v617
		goto L121
	} else {
		goto L127
	}
L127:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v610+v526)))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v610+v577)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v610+v579)))
	v628 = F_FunctionCall2Coll(m, v527+v584*int32(28), v623, v625, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v628 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v666 = v628
	v669 = v617
	goto L120
L130:
	;
	goto L131
L131:
	;
	if v528 != v617 {
		v584 = v617
		goto L123
	} else {
		goto L132
	}
L132:
	;
	goto L124
L133:
	;
	if v567 == int32(-1) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v664 = int32(1)
	goto L136
L135:
	;
	v664 = int32(-1)
	goto L136
L136:
	;
	v666 = v664
	v669 = v658
	goto L120
L137:
	;
	v696 = v692 - v669
	goto L139
L138:
	;
	v696 = v669
	goto L139
L139:
	;
	v724 = v696
	goto L119
L140:
	;
	if v724 == int32(0) {
		goto L111
	} else {
		goto L141
	}
L141:
	;
	v763 = v724
	v776 = v562
	v780 = v553
	goto L117
L142:
	;
	goto L116
L143:
	;
	v844 = l0
	v845 = v815
	v847 = l3
	v850 = v819
	v852 = v30
	v863 = v35
	v867 = v804
	v868 = v37 + int32(24)
	goto L109
L144:
	;
	v844 = l0
	v845 = v826
	v847 = l3
	v850 = v830
	v852 = v30
	v863 = v35
	v867 = v562
	v868 = v37 + int32(24)
	goto L109
L145:
	;
	v844 = l0
	v845 = v837
	v847 = l3
	v850 = v841
	v852 = v30
	v863 = v35
	v867 = v562
	v868 = v530
	goto L109
L146:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v876 = v845 << (uint(int32(2)) % 32)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v876+v877)))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v880+v876)))
	v884 = base.B2i32(v850 == int32(-1))
	v885 = int32(0)
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+12)))
	v888 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v888 <= v885 {
		v945 = v885
		goto L150
	} else {
		goto L151
	}
L147:
	;
	if int32(0) <= v1066 {
		v1370 = v852
		goto L4
	} else {
		goto L172
	}
L148:
	;
	v1066 = v1038
	goto L147
L149:
	;
	v1004 = int32(0)
	if v988 < v1004 {
		goto L166
	} else {
		goto L167
	}
L150:
	;
	if v850 == int32(-1) {
		goto L160
	} else {
		goto L161
	}
L151:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v897 = v885
	goto L152
L152:
	;
	v923 = v897 << (uint(int32(2)) % 32)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v882+v923)))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v923+v891)))
	if v925 < v927 {
		v1066 = v897 ^ int32(-1)
		goto L147
	} else {
		goto L154
	}
L153:
	;
	v945 = v888
	goto L150
L154:
	;
	v930 = v897 + int32(1)
	if v927 < v925 {
		v1038 = v930
		goto L148
	} else {
		goto L155
	}
L155:
	;
	if v925 != 0 {
		v945 = v930
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v874+v923)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v923+v879)))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v923+v892)))
	v941 = F_FunctionCall2Coll(m, v873+v897*int32(28), v936, v938, v940)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	if v941 != 0 {
		v978 = v930
		v988 = v941
		goto L149
	} else {
		goto L158
	}
L158:
	;
	if v930 != v888 {
		v897 = v930
		goto L152
	} else {
		goto L159
	}
L159:
	;
	goto L153
L160:
	;
	v973 = int32(1)
	goto L162
L161:
	;
	v973 = int32(-1)
	goto L162
L162:
	;
	if v887 != v884 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v976 = v973
	goto L165
L164:
	;
	v976 = int32(0)
	goto L165
L165:
	;
	v978 = v945
	v988 = v976
	goto L149
L166:
	;
	v1008 = v1004 - v978
	goto L168
L167:
	;
	v1008 = v978
	goto L168
L168:
	;
	if v988 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1010 = v1008
	goto L171
L170:
	;
	v1010 = int32(0)
	goto L171
L171:
	;
	v1038 = v1010
	goto L148
L172:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v1070 = int32(2)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v867<<(uint(v1070)%32))+8))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+12))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1075+(v1066^int32(-1))<<(uint(v1070)%32))))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	v1307 = v844
	v1310 = v847
	v1313 = v1073
	v1315 = v852
	v1322 = v1082
	v1326 = v863
	goto L16
L173:
	;
	v1083 = l0
	v1086 = l3
	v1089 = v819
	v1091 = v30
	v1102 = v35
	goto L107
L174:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1156 == int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+4))
	if v1159 <= int32(0) {
		v1370 = v30
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v1178 = v5
	v1183 = v5
	goto L177
L177:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+12))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1178<<(uint(int32(2))%32))))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+28))
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193)+24)))
	if v1195 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v1370 = v30
	goto L4
L179:
	;
	v1304 = v1178 + int32(1)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+4))
	if v1304 < v1305 {
		v1178 = v1304
		v1183 = v1297
		goto L177
	} else {
		goto L199
	}
L180:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1201 = v1199 - int32(1)
	if v1201 < int32(0) {
		v1297 = v1183
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v1273 != int32(-1) {
		v1307 = l0
		v1310 = l3
		v1313 = v1273
		v1315 = v30
		v1322 = v1194
		v1326 = v35
		goto L16
	} else {
		goto L198
	}
L183:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+20))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1208 = v1201
	v1211 = int32(-1)
	v1216 = v1183
	goto L184
L184:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v1239 = int32(2)
	v1240 = base.I32_div_s(v1208+v1211+int32(1), v1239)
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1240<<(uint(v1239)%32))))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1244)))
	v1246 = F_FunctionCall2Coll(m, v1206, v1234, v1245, v1204)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v1262 = base.B2i32(v1259 < int32(0))
	v1263 = v1262 & v1260
	if v1259 < int32(0) {
		v1297 = v1263
		goto L179
	} else {
		goto L196
	}
L186:
	;
	goto L185
L187:
	;
	if v1255 < v1254 {
		v1208 = v1254
		v1211 = v1255
		v1216 = v1256
		goto L184
	} else {
		goto L195
	}
L188:
	;
	if v1246 <= int32(0) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if v1246 != 0 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	v1254 = v1240 - int32(1)
	v1255 = v1211
	v1256 = v1216
	goto L187
L192:
	;
	v1254 = v1208
	v1255 = v1240
	v1256 = int32(0)
	goto L187
L193:
	;
	goto L194
L194:
	;
	v1259 = v1240
	v1260 = int32(1)
	goto L186
L195:
	;
	v1259 = v1255
	v1260 = v1256
	goto L186
L196:
	;
	if v1260&int32(1) == int32(0) {
		v1297 = v1263
		goto L179
	} else {
		goto L197
	}
L197:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1268+v1259<<(uint(int32(2))%32))))
	v1307 = l0
	v1310 = l3
	v1313 = v1272
	v1315 = v30
	v1322 = v1194
	v1326 = v35
	goto L16
L198:
	;
	v1297 = v1183
	goto L179
L199:
	;
	goto L178
L200:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+8))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1341+v1313<<(uint(int32(2))%32))))
	v1346 = F_get_rel_name(m, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+20)) = v1346
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+16)) = v1307
	F_errmsg(m, int32(677407), v1315+int32(16))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_parser_errposition(m, v1310, v1322)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(485356), int32(3239), int32(417024))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_partition_ancestors_worker(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_get_partition_parent_worker(m, l0, l1, v7+int32(15))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			if v15 != 0 {
				m.G0 = v7 + int32(16)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v17 = F_lappend_oid(m, v16, v11)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
					F_get_partition_ancestors_worker(m, l0, v11, l2)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_get_partition_parent_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = int32(3)
		F_ScanKeyInit(m, v8+int32(48), v21, v21, int32(65), int32(1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v31 = F_systable_beginscan(m, l0, int32(2680), int32(1), int32(0), int32(2), v8)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_systable_getnext(m, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
						v37 = v35 + v36
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)))
						if v38 == int32(1) {
							v41 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v41)
						} else {
						}
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						v44 = v43
					} else {
						v44 = v4
					}
					F_systable_endscan(m, v31)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(96)
						return v44
					}
				}
			}
		}
	}
}
func F_transformPartitionBoundValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_transformExpr(m, l0, l1, int32(39))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_exprType(m, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v23 = F_coerce_to_target_type(m, l0, v14, v18, l3, l4, int32(1), int32(2), int32(-1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					if v25 != int32(7) {
						F_assign_expr_collations(m, l0, v23)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = F_expression_planner(m, v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = F_evaluate_expr(m, v30, l3, l4, l5)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									if v34 == int32(7) {
										v51 = v32
										v52 = F_exprLocation(m, l1)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v52
										m.G0 = v11 + int32(16)
										return v51
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(265771), int32(0))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(490667), int32(4636), int32(341422))
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
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
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l5
						v51 = v23
						v52 = F_exprLocation(m, l1)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v52
						m.G0 = v11 + int32(16)
						return v51
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v65 = F_format_type_be(m, l3)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v65
								F_errmsg(m, int32(684262), v11)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v72 = F_exprLocation(m, l1)
									mBase = m.M
									F_parser_errposition(m, l0, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490667), int32(4623), int32(341422))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
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
		}
	}
}
func F_tryAttachPartitionForeignKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v112 int32
	_ = v112
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
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
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	v22 = m.G0
	v24 = v22 - int32(48)
	m.G0 = v24
	v27 = F_SearchSysCache1(m, int32(19), l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L118
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L3
	} else {
		goto L115
	}
L3:
	;
	return int32(0)
L4:
	;
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
	if v31 != v35 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L3
	} else {
		goto L112
	}
L8:
	;
	m.G0 = v24 + int32(48)
	return v522
L9:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v137 = F_SearchSysCache1(m, int32(19), v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L21
	}
L10:
	;
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L20
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v37 != l4 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if l4 <= int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(0)
	goto L14
L14:
	;
	v69 = v59 << (uint(int32(1)) % 32)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(22)+v69))))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5+v69))))
	if v71 != v73 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+(l1+int32(86))))))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6+v69))))
	if v76 != v78 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v81 = v59 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(152)+v81)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l7+v81)))
	if v83 != v85 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v88 = v59 + int32(1)
	if l4 != v88 {
		v59 = v88
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v522 = int32(0)
	goto L8
L21:
	;
	if v137 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+22)))
	v143 = v141 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+75)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+75)))
	if v144 != v145 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+92))
	if v147 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L34
	}
L25:
	;
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L32
	}
L26:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+73)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+73)))
	if v148 != v149 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+74)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+74)))
	if v151 != v152 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+100)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+100)))
	if v154 != v155 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+101)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+101)))
	if v157 != v158 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+102)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+102)))
	if v160 == v161 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	F_ReleaseCatCache(m, v137)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v522 = int32(0)
	goto L8
L34:
	;
	F_ReleaseCatCache(m, v137)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v173 = m.G0
	v175 = v173 - int32(144)
	m.G0 = v175
	v178 = F_SearchSysCache1(m, int32(19), l3)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L39
	}
L36:
	;
	v522 = int32(1)
	goto L8
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L3
	} else {
		goto L109
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L3
	} else {
		goto L106
	}
L39:
	;
	if v178 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
	v182 = v180 + v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+75)))
	v185 = F_SearchSysCache1(m, int32(19), v172)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L103
	}
L43:
	;
	if v185 == int32(0) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v191 = v189 + v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+80))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+96))
	v194 = F_get_rel_relkind(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	if v194 == int32(112) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v200 = F_table_open(m, int32(2606), int32(2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+76)))
	if v394 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L49:
	;
	F_ScanKeyInit(m, v175+int32(96), int32(9), int32(3), int32(184), v192)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v210 = int32(1)
	v215 = F_systable_beginscan(m, v200, int32(2665), v210, int32(0), v210, v175+int32(96))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v217 = F_new_object_addresses(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v219 = F_systable_getnext(m, v215)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	if v219 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v222 = v219
	goto L57
L55:
	;
	goto L56
L56:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L76
	}
L57:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+22)))
	v244 = v242 + v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+92))
	if v172 == v245 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+84)) = int32(2606)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+88)) = v249
	F_add_exact_object_address(m, v175+int32(84), v217)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v338 = F_systable_getnext(m, v215)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L74
	}
L62:
	;
	v257 = int32(2606)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	F_deleteDependencyRecordsForSpecific(m, v257, v258, int32(105), v257, v172)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	F_ScanKeyInit(m, v175+int32(36), int32(11), int32(3), int32(184), v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v272 = int32(1)
	v277 = F_systable_beginscan(m, l10, int32(2699), v272, int32(0), v272, v175+int32(36))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	goto L66
L66:
	;
	v300 = F_systable_getnext(m, v277)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L68
	}
L67:
	;
	F_systable_endscan(m, v277)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L3
	} else {
		goto L73
	}
L68:
	;
	if v300 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+84)) = int32(2620)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+22)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304+v305)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+88)) = v307
	F_add_exact_object_address(m, v175+int32(84), v217)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L67
L72:
	;
	goto L66
L73:
	;
	goto L61
L74:
	;
	if v338 != 0 {
		v222 = v338
		goto L57
	} else {
		goto L75
	}
L75:
	;
	goto L58
L76:
	;
	F_performMultipleDeletions(m, v217, int32(0), int32(1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_systable_endscan(m, v215)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_sequence_close(m, v200, int32(2))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	goto L48
L80:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+76)))
	v400 = v397 ^ int32(1)
	goto L82
L81:
	;
	v400 = int32(0)
	goto L82
L82:
	;
	F_ReleaseCatCache(m, v185)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_ReleaseCatCache(m, v178)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	F_DropForeignKeyConstraintTriggers(m, l10, v172, v193, v192)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ConstraintSetParentConstraint(m, v172, l3, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	if v183&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_GetForeignKeyCheckTriggers(m, l10, v172, v193, v192, v175+int32(96), v175+int32(36))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L3
	} else {
		goto L93
	}
L90:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v175)+96))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v418, l8, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v175)+36))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v422, l9, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v400&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v432 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L3
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	m.G0 = v175 + int32(144)
	goto L36
L97:
	;
	v435 = F_SearchSysCache1(m, int32(19), v172)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	if v435 == int32(0) {
		goto L37
	} else {
		goto L99
	}
L99:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435)+16))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+22)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v439+v440)+96))
	F_QueueFKConstraintValidation(m, l0, v432, l2, v442, v435, int32(4))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	F_ReleaseCatCache(m, v435)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	F_sequence_close(m, v432, int32(3))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = l3
	F_errmsg_internal(m, int32(40426), v175)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(485619), int32(11811), int32(22300))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L3
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
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v172
	F_errmsg_internal(m, int32(40426), v175+int32(16))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(485619), int32(11819), int32(22300))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
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
	*(*int32)(unsafe.Add(mBase, uint32(v175)+32)) = v172
	F_errmsg_internal(m, int32(40426), v175+int32(32))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(485619), int32(11902), int32(22300))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
	F_errmsg_internal(m, int32(40426), v24)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(485619), int32(11714), int32(22297))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v544
	F_errmsg_internal(m, int32(40426), v24+int32(16))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(485619), int32(11740), int32(22297))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L3
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v564 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v143 + v564
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v34 + v564
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v563 + v564
	F_errmsg(m, int32(679999), v24+int32(32))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(485619), int32(11757), int32(22297))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
