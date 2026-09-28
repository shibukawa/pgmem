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
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int64
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
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int64
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v586 int32
	_ = v586
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
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
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
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int64
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v662 int64
	_ = v662
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(784)
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
	v23 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
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
	v27 = v18 + int32(512)
	F_ScanKeyInit(m, v27, int32(2), int32(3), int32(184), v23)
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
	v41 = F_systable_beginscan(m, v35, int32(2701), v38, int32(0), v38, v27)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v43 = F_systable_getnext(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v43
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L24
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v63 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v90 = F_systable_getnext(m, v41)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L22
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+84))
	if v66 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(2620)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v71 = F_deleteDependencyRecordsForClass(m, v67, v68, v67, int32(80))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v77 = F_deleteDependencyRecordsForClass(m, int32(2620), v74, int32(1259), int32(83))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+240)) = int32(2620)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+244)) = v81
	F_add_exact_object_address(m, v18+int32(240), v24)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L16
L22:
	;
	if v90 != 0 {
		v47 = v90
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
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_free_object_addresses(m, v24)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_systable_endscan(m, v41)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_relation_close(m, v35, int32(3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v120 = F_RelationGetFKeyList(m, l1)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v494 = F_GetParentedForeignKeyRefs(m, l1)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L100
	}
L30:
	;
	v122 = F_copyObjectImpl(m, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v122 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_list_free_deep(m, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v130 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if int32(0) < v132 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v138 = int32(0)
	v148 = v5
	goto L40
L38:
	;
	v175 = v5
	goto L39
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if int32(0) < v178 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v138<<(uint(int32(2))%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v157 = F_lappend_oid(m, v148, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	v175 = v157
	goto L39
L42:
	;
	v160 = v138 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v160 < v161 {
		v138 = v160
		v148 = v157
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v194 = v5
	goto L47
L45:
	;
	goto L46
L46:
	;
	F_list_free_deep(m, v122)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L96
	}
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v194<<(uint(int32(2))%32))))
	v202 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v201)+4)))
	v203 = F_SearchSysCache1(m, int32(19), v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L52
	}
L48:
	;
	goto L46
L49:
	;
	F_ReleaseCatCache(m, v203)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L94
	}
L50:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v410 = F_table_open(m, v408, int32(6))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L91
	}
L51:
	;
	v355 = v317
	v362 = v317
	goto L86
L52:
	;
	if v203 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+22)))
	v207 = v205 + v206
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+72)))
	if v208 != int32(102) {
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
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L83
	}
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+92))
	if v211 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v214 = int32(0)
	if v175 == v214 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v252 != 0 {
		goto L49
	} else {
		goto L71
	}
L59:
	;
	v252 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v220 <= int32(0) {
		v246 = v214
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v252 = v246
	goto L58
L63:
	;
	v223 = int32(0)
	if v223 < v220 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v226 = v220
	goto L66
L65:
	;
	v226 = v223
	goto L66
L66:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v229 = int32(0)
	goto L67
L67:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227+v229<<(uint(int32(2))%32))))
	v238 = base.B2i32(v237 == v211)
	if v237 == v211 {
		v246 = v238
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v246 = v238
	goto L62
L69:
	;
	v240 = v229 + int32(1)
	if v240 != v226 {
		v229 = v240
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v254 = int32(0)
	F_ConstraintSetParentConstraint(m, v253, v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+20)))
	if v258 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	F_GetForeignKeyCheckTriggers(m, v130, v261, v262, v263, v18+int32(512), v18+int32(240))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_DeconstructFkConstraintRow(m, v203, v18+int32(508), v18+int32(432), v18+int32(368), v18+int32(512), v18+int32(240), v18+int32(112), v18+int32(108), v18+int32(32))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L79
	}
L76:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v18)+512))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v130, v270, int32(0), v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v18)+240))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v130, v275, int32(0), v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v299 = F_palloc0(m, int32(108))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = int64(438086664353)
	v305 = F_pstrdup(m, v207+int32(4))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+8)) = v305
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+12)) = uint8(v308)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+13)) = uint8(v310)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+75)))
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+15)) = uint8(v313)
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+14)) = uint8(v312)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+76)))
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v299)+80)) = v317
	v320 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v299)+72)) = v320
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+16)) = uint8(v316)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+86)) = uint8(v323)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+87)) = uint8(v325)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+101)))
	*(*int64)(unsafe.Add(mBase, uint32(v299)+100)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v299)+92)) = v320
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+88)) = uint8(v327)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	if v317 < v334 {
		goto L51
	} else {
		goto L82
	}
L82:
	;
	goto L50
L83:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v341
	F_errmsg_internal(m, int32(_a_F_DetachPartitionFinalize_0), v18+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_DetachPartitionFinalize_1), int32(_a_F_DetachPartitionFinalize_2), int32(_a_F_DetachPartitionFinalize_3))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v378 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18+int32(432)+v355<<(uint(int32(1))%32)))))
	v384 = F_makeString(m, v368+v369<<(uint(int32(3))%32)+v378*int32(100)-int32(68))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L50
L88:
	;
	v386 = F_lappend(m, v362, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+76)) = v386
	v390 = v355 + int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	if v390 < v391 {
		v355 = v390
		v362 = v386
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v207)+88))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v18)+508))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	v428 = int32(0)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+107)))
	F_addFkRecurseReferenced(m, v299, l1, v410, v412, v413, v414, v18+int32(368), v18+int32(432), v18+int32(512), v18+int32(240), v18+int32(112), v425, v18+int32(32), v428, v428, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_relation_close(m, v410, int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v454 = v194 + int32(1)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v454 < v455 {
		v194 = v454
		goto L47
	} else {
		goto L95
	}
L95:
	;
	goto L48
L96:
	;
	if v130 == int32(0) {
		goto L29
	} else {
		goto L97
	}
L97:
	;
	F_relation_close(m, v130, int32(3))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	goto L29
L99:
	;
	v563 = F_RelationGetIndexList(m, l1)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L111
	}
L100:
	;
	if v494 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v498 <= int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v504 = int32(0)
	goto L103
L103:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v504<<(uint(int32(2))%32))))
	v522 = int32(0)
	F_ConstraintSetParentConstraint(m, v521, v522, v522)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L105
	}
L104:
	;
	goto L99
L105:
	;
	v526 = int32(2606)
	v529 = F_deleteDependencyRecordsForClass(m, v526, v521, v526, int32(105))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v533 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+520)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v18)+516)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v18)+512)) = int32(2606)
	F_performDeletion(m, v18+int32(512), v533, v533)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v545 = v504 + int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v545 < v546 {
		v504 = v545
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	v646 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L4
	} else {
		goto L131
	}
L111:
	;
	if v563 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v567 <= int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v573 = int32(0)
	goto L114
L114:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v586+v573<<(uint(int32(2))%32))))
	v591 = F_has_superclass(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	goto L110
L116:
	;
	if v591 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v594 = F_get_partition_parent(m, v590, int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v626 = v573 + int32(1)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v626 < v627 {
		v573 = v626
		goto L114
	} else {
		goto L130
	}
L120:
	;
	v597 = F_index_open(m, v590, int32(8))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	F_IndexSetParentIndex(m, v597, int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v603 = F_get_relation_idx_constraint_oid(m, v602, v590)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v606 = F_get_relation_idx_constraint_oid(m, v605, v594)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v608 = int32(0)
	if base.B2i32(v606 == v608)|base.B2i32(v603 == v608) == v608 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v615 = int32(0)
	F_ConstraintSetParentConstraint(m, v603, v615, v615)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	F_relation_close(m, v597, int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	goto L119
L130:
	;
	goto L115
L131:
	;
	v649 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
	v651 = F_SearchSysCacheCopy(m, int32(57), v649, int64(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v651 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v654 = v18 + int32(512)
	v655 = int32(0)
	base.MemoryFill(m, v654, v655, int32(272))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+144)) = uint16(v655)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+272)) = uint16(v655)
	v662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+264)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+256)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+248)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+240)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v662
	v678 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v678)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+273)) = uint8(v678)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v646)+52))
	v687 = F_heap_modify_tuple(m, v651, v682, v654, v18+int32(240), v18+int32(112))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L4
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L167
	}
L136:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v687)+16))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689)+22)))
	v692 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v689+v690)+131)) = uint8(v692)
	F_CatalogTupleUpdate(m, v646, v687+int32(4), v687)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v687)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_relation_close(m, v646, int32(3))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v704 = int32(*(*int16)(unsafe.Add(mBase, uint32(v703)+120)))
	if int32(0) < v704 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v710 = int32(0)
	v713 = v703
	goto L143
L141:
	;
	goto L142
L142:
	;
	if l3 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)))
	v730 = v723 + v724<<(uint(int32(3))%32) + v710*int32(100)
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+119)))
	if v731 != 0 {
		v748 = v713
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	v751 = v710 + int32(1)
	v752 = int32(*(*int16)(unsafe.Add(mBase, uint32(v748)+120)))
	if v751 < v752 {
		v710 = v751
		v713 = v748
		goto L143
	} else {
		goto L149
	}
L146:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+int32(28))+89)))
	if v734 == int32(0) {
		v748 = v713
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v743 = int32(1)
	F_ATExecDropIdentity(m, v18+int32(432), l1, v730+int32(32), int32(0), int32(8), v743, v743)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v748 = v747
	goto L145
L149:
	;
	goto L144
L150:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L157
	}
L151:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l3 == v771 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_update_default_partition_oid(m, v773, int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	F_CacheInvalidateRelcacheByRelid(m, l3)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L4
	} else {
		goto L156
	}
L155:
	;
	goto L150
L156:
	;
	goto L150
L157:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+119)))
	if v782 != int32(112) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	m.G0 = v18 + int32(784)
	return
L159:
	;
	v785 = int32(0)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v789 = F_find_all_inheritors(m, v786, int32(8), v785)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	if v789 == int32(0) {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	if v793 <= int32(0) {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	v798 = v785
	goto L163
L163:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v789)+12))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v811+v798<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L4
	} else {
		goto L165
	}
L164:
	;
	goto L158
L165:
	;
	v819 = v798 + int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	if v819 < v820 {
		v798 = v819
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v844
	F_errmsg_internal(m, int32(_a_F_DetachPartitionFinalize_4), v18)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_DetachPartitionFinalize_1), int32(_a_F_DetachPartitionFinalize_5), int32(_a_F_DetachPartitionFinalize_3))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	v5 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v9 == v5 {
		v12 = int32(_a_F_ExecPartitionCheck_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = F_RelationGetPartitionQual(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(_a_F_ExecPartitionCheck_0)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
			*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v25
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
							v35 = int32(_a_F_ExecPartitionCheck_0)
							*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v23
							*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v34
							*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v13
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
										v54 = int32(0)
										if v52|base.B2i32(l3 == v54) == v54 {
											F_ExecPartitionCheckEmitError(m, l0, l1, l2)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											return v52
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
									v54 = int32(0)
									if v52|base.B2i32(l3 == v54) == v54 {
										F_ExecPartitionCheckEmitError(m, l0, l1, l2)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										return v52
									}
								}
							}
						}
					}
				} else {
					v34 = v5
					v35 = int32(_a_F_ExecPartitionCheck_0)
					*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v23
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v34
					*(*int32)(unsafe.Add(mBase, _c_F_ExecPartitionCheck[0])) = v13
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
								v54 = int32(0)
								if v52|base.B2i32(l3 == v54) == v54 {
									F_ExecPartitionCheckEmitError(m, l0, l1, l2)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									return v52
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
							v54 = int32(0)
							if v52|base.B2i32(l3 == v54) == v54 {
								F_ExecPartitionCheckEmitError(m, l0, l1, l2)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								return v52
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
					v54 = int32(0)
					if v52|base.B2i32(l3 == v54) == v54 {
						F_ExecPartitionCheckEmitError(m, l0, l1, l2)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return v52
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
				v54 = int32(0)
				if v52|base.B2i32(l3 == v54) == v54 {
					F_ExecPartitionCheckEmitError(m, l0, l1, l2)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return v52
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	v27 = v13 * v16
	v28 = F_palloc0_mul(m, int32(28), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v28
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_InitPartitionPruneContext[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
	v37 = F_palloc0_mul(m, int32(4), v27)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v37
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v42 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = v7
	goto L10
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v54<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = v63
	goto L14
L13:
	;
	v64 = int32(0)
	goto L14
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v68 = int32(0)
	if base.B2i32(v65 != int32(381))|base.B2i32(v16 <= v68) == v68 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = v64
	v77 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v147 = v54 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v147 < v148 {
		v54 = v147
		goto L10
	} else {
		goto L37
	}
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v86 = F_bms_is_member(m, v77, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v88 = int32(0)
	if v86|base.B2i32(v76 == v88) == v88 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != int32(7) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v129 = v76
	goto L23
L23:
	;
	v133 = v77 + int32(1)
	if v133 != v16 {
		v76 = v129
		v77 = v133
		goto L18
	} else {
		goto L36
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
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
	v119 = v76 + int32(4)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if base.Ui32(v119) < base.Ui32(v122+v123<<(uint(int32(2))%32)) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v109 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v108+v97*v16<<(uint(v109)%32)+v77<<(uint(v109)%32)))) = v107
	goto L26
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v102 = F_ExecInitExprWithParams(m, v93, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v105 = F_ExecInitExpr(m, v93, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v107 = v102
	goto L27
L32:
	;
	v107 = v105
	goto L27
L33:
	;
	v128 = v119
	goto L35
L34:
	;
	v128 = int32(0)
	goto L35
L35:
	;
	v129 = v128
	goto L23
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
	var v25 int32
	_ = v25
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
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	v5 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+2)))
	if v5 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v194)
	return v196
L2:
	;
	v18 = base.B2i32(l2 == int32(-1))
	v25 = v5
	v30 = v5
	goto L5
L3:
	;
	v183 = v5
	goto L4
L4:
	;
	v194 = int32(0)
	v196 = v183
	goto L1
L5:
	;
	v32 = v30 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
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
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v183 = v167
	goto L4
L7:
	;
	v174 = v30 + int32(1)
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+2)))
	if v174 < v175 {
		v25 = v167
		v30 = v174
		goto L5
	} else {
		goto L37
	}
L8:
	;
	v159 = F_lappend(m, v25, v50)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L36
	}
L9:
	;
	return int32(0)
L10:
	;
	if v50 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+40)))
	if v55 != 0 {
		v167 = v25
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v84 = int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86+v32)))
	if base.B2i32(v88 != int32(2222))&base.B2i32(v88 != int32(424)) != 0 {
		v194 = v84
		v196 = v25
		goto L1
	} else {
		goto L21
	}
L14:
	;
	if v25 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v58 <= int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v65 = int32(0)
	goto L17
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v61+v65<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v54 == v79 {
		v167 = v25
		goto L7
	} else {
		goto L19
	}
L18:
	;
	goto L8
L19:
	;
	v82 = v65 + int32(1)
	if v82 != v58 {
		v65 = v82
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v94 == int32(0) {
		v194 = v84
		v196 = v25
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v98 <= v97 {
		v194 = v84
		v196 = v25
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v103 = v97
	goto L24
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v103<<(uint(int32(2))%32))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v118 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v194 = v142
	v196 = v25
	goto L1
L26:
	;
	v142 = int32(1)
	v144 = v103 + v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v144 < v145 {
		v103 = v144
		goto L24
	} else {
		goto L35
	}
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119+v32)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v125 = F_equal(m, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if v125 != 0 {
		v167 = v25
		goto L7
	} else {
		goto L29
	}
L29:
	;
	if v124 == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v129 != int32(21) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v132 != int32(2) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v138 = F_equal(m, v123, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if v138 != 0 {
		v167 = v25
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	goto L25
L36:
	;
	v167 = v159
	goto L7
L37:
	;
	goto L6
}
func F_check_new_partition_bound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v223 int32
	_ = v223
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
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
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
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
	var v480 int32
	_ = v480
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int64
	_ = v527
	var v529 int64
	_ = v529
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v572 int32
	_ = v572
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
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
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v781 int64
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v881 int32
	_ = v881
	var v891 int32
	_ = v891
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1058 int32
	_ = v1058
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int64
	_ = v1102
	var v1104 int64
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1206 int32
	_ = v1206
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1268 int32
	_ = v1268
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1309 int32
	_ = v1309
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1376 int64
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1391 int32
	_ = v1391
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int64
	_ = v1418
	var v1419 int64
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
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
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1460 int32
	_ = v1460
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1545 int32
	_ = v1545
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(112)
	m.G0 = v31
	v33 = F_RelationGetPartitionKey(m, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = F_RelationGetPartitionDesc(m, l1, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v39 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v1545 + int32(112)
	return
L5:
	;
	if v38 == int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	switch v75 - int32(104) {
	case 0:
		goto L19
	default:
		v1545 = v31
		goto L4
	case 4:
		goto L17
	case 10:
		goto L18
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v44 == int32(-1) {
		v1545 = v31
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55<<(uint(int32(2))%32))))
	v60 = F_get_rel_name(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l0
	F_errmsg(m, int32(_a_F_check_new_partition_bound_0), v31)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_parser_errposition(m, l3, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2914), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L204
	}
L17:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v1324 <= int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L179
	}
L18:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v461 = F_make_one_partition_rbound(m, v33, int32(-1), v459, int32(1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L80
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v78 <= int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v85 = v83 - int32(1)
	if int32(0) <= v85 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L74
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L68
	}
L23:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v356 <= v81 {
		goto L59
	} else {
		goto L60
	}
L24:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v314 = v223 << (uint(int32(2)) % 32)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312+v314)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v318 = base.I32_rem_s(v82, v317)
	if v318 != 0 {
		goto L22
	} else {
		goto L56
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v94 = int32(-1)
	v96 = v85
	v103 = v5
	goto L28
L26:
	;
	goto L27
L27:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v280 = base.I32_rem_s(v279, v82)
	if v280 == int32(0) {
		goto L23
	} else {
		goto L49
	}
L28:
	;
	v119 = int32(2)
	v120 = base.I32_div_s(v96+v103, v119)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v88+v120<<(uint(v119)%32))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v125 < v82 {
		v192 = v120
		v194 = v96
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if int32(0) <= v223 {
		goto L24
	} else {
		goto L48
	}
L30:
	;
	goto L29
L31:
	;
	if v192 < v194 {
		v94 = v192
		v96 = v194
		v103 = v192 + int32(1)
		goto L28
	} else {
		goto L47
	}
L32:
	;
	v131 = v120
	v133 = v96
	v135 = v125
	v138 = v124
	goto L33
L33:
	;
	if v135 <= v82 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v192 = v181
	v194 = v175
	goto L31
L35:
	;
	v180 = int32(2)
	v181 = base.I32_div_s(v175+v103, v180)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v88+v181<<(uint(v180)%32))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v82 <= v186 {
		v131 = v181
		v133 = v175
		v135 = v186
		v138 = v185
		goto L33
	} else {
		goto L46
	}
L36:
	;
	v223 = v94
	goto L30
L37:
	;
	v156 = base.B2i32(v135 != v82)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v159 = v156 | base.B2i32(v81 == v157)
	if v156|base.B2i32(v157 <= v81) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v169 = v131 - int32(1)
	if v94 < v169 {
		v175 = v169
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v165 = v131 - int32(1)
	if v159|base.B2i32(v165 <= v94) != 0 {
		goto L36
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v159 != 0 {
		v223 = v131
		goto L30
	} else {
		goto L44
	}
L43:
	;
	v175 = v165
	goto L35
L44:
	;
	v192 = v131
	v194 = v133
	goto L31
L45:
	;
	goto L36
L46:
	;
	goto L34
L47:
	;
	v223 = v192
	goto L30
L48:
	;
	goto L27
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v297 = F_get_rel_name(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v294
	v305 = F_errdetail(m, int32(_a_F_check_new_partition_bound_4), v31+int32(32))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2967), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
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
	v320 = v223 + int32(1)
	if v83 <= v320 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	v323 = v320 << (uint(int32(2)) % 32)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v312+v323)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = base.I32_rem_s(v326, v82)
	if v327 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	goto L23
L59:
	;
	v358 = base.I32_rem_s(v81, v356)
	v359 = v358
	goto L61
L60:
	;
	v359 = v81
	goto L61
L61:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v362 = v359
	goto L62
L62:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v360+v362<<(uint(int32(2))%32))))
	if v392 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v1545 = v31
	goto L4
L64:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1480 = l0
	v1483 = l3
	v1484 = v392
	v1489 = v31
	v1490 = v395
	v1496 = v36
	goto L16
L65:
	;
	goto L66
L66:
	;
	v396 = v362 + v82
	if v396 < v356 {
		v362 = v396
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v410+v314)))
	v413 = F_get_rel_name(m, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v31)+68)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v409
	v421 = F_errdetail(m, int32(_a_F_check_new_partition_bound_5), v31-int32(-64))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2987), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v440+v323)))
	v443 = F_get_rel_name(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v439
	v451 = F_errdetail(m, int32(_a_F_check_new_partition_bound_4), v31+int32(48))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3008), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v466 = F_make_one_partition_rbound(m, v33, int32(-1), v464, int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+12)))
	v469 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v469 <= int32(0) {
		v542 = v5
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v666 <= int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L109
	}
L83:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v597+v572<<(uint(int32(2))%32)-int32(4))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L101
	}
L84:
	;
	if v468 != 0 {
		v643 = v5
		goto L82
	} else {
		goto L99
	}
L85:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v480 = int32(0)
	goto L86
L86:
	;
	v508 = v480 << (uint(int32(2)) % 32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v474+v508)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508+v472)))
	if v510 < v512 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v542 = v469
	goto L84
L88:
	;
	v643 = v480 ^ int32(-1)
	goto L82
L89:
	;
	goto L90
L90:
	;
	v517 = v480 + int32(1)
	if v512 < v510 {
		v572 = v517
		goto L83
	} else {
		goto L91
	}
L91:
	;
	if v510 != 0 {
		v542 = v517
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v508+v476)))
	v525 = v480 << (uint(int32(3)) % 32)
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v475+v525)))
	v529 = *(*int64)(unsafe.Add(mBase, uint32(v525+v473)))
	v530 = F_FunctionCall2Coll(m, v477+v480*int32(28), v523, v527, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v532 = base.I32_wrap_i64(v530)
	if v532 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if int32(0) <= v532 {
		v572 = v517
		goto L83
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v517 != v469 {
		v480 = v517
		goto L86
	} else {
		goto L98
	}
L97:
	;
	v643 = v480 ^ int32(-1)
	goto L82
L98:
	;
	goto L87
L99:
	;
	if v542 <= int32(0) {
		v643 = v542
		goto L82
	} else {
		goto L100
	}
L100:
	;
	v572 = v542
	goto L83
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = l0
	F_errmsg(m, int32(_a_F_check_new_partition_bound_6), v31+int32(96))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v618 = F_get_range_partbound_string(m, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v621 = F_get_range_partbound_string(m, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v618
	v628 = F_errdetail(m, int32(_a_F_check_new_partition_bound_7), v31+int32(80))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	F_parser_errposition(m, l3, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3124), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
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
	v669 = int32(-1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v672 = v670 - int32(1)
	if int32(0) <= v672 {
		goto L117
	} else {
		goto L118
	}
L110:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+12))
	v1480 = v1293
	v1483 = v1296
	v1484 = v1297
	v1489 = v1302
	v1490 = v1323
	v1496 = v1309
	goto L16
L111:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+12))
	v1285 = v950 >> (uint(int32(31)) % 32)
	v1293 = l0
	v1296 = l3
	v1297 = v979
	v1302 = v31
	v1309 = v36
	v1321 = v1283 + (v950^v1285-v1285)<<(uint(int32(2))%32) - int32(4)
	goto L110
L112:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+12))
	v1293 = v1252
	v1296 = v1255
	v1297 = v1256
	v1302 = v1261
	v1309 = v1268
	v1321 = v1281
	goto L110
L113:
	;
	if v950 != 0 {
		goto L111
	} else {
		goto L178
	}
L114:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v1032 <= v1005 {
		v1545 = v1013
		goto L4
	} else {
		goto L151
	}
L115:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v997 = v712 + int32(1)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v995+v997<<(uint(int32(2))%32))))
	if int32(0) <= v1001 {
		v1252 = l0
		v1255 = l3
		v1256 = v1001
		v1261 = v31
		v1268 = v36
		goto L112
	} else {
		goto L150
	}
L116:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v986 = v712 + int32(1)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v984+v986<<(uint(int32(2))%32))))
	if int32(0) <= v990 {
		v1252 = l0
		v1255 = l3
		v1256 = v990
		v1261 = v31
		v1268 = v36
		goto L112
	} else {
		goto L149
	}
L117:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v677 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	v679 = v38 + int32(24)
	v686 = v669
	v694 = v672
	goto L120
L118:
	;
	v950 = v643
	v951 = v669
	goto L119
L119:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v975 = v951 + int32(1)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v973+v975<<(uint(int32(2))%32))))
	if int32(0) <= v979 {
		goto L113
	} else {
		goto L148
	}
L120:
	;
	v711 = int32(2)
	v712 = base.I32_div_s(v686+v694+int32(1), v711)
	v714 = v712 << (uint(v711) % 32)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v714+v715)))
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+12)))
	v719 = int32(0)
	if v677 <= v719 {
		v813 = v719
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v950 = v921
	v951 = v922
	goto L119
L122:
	;
	if v922 < v930 {
		v686 = v922
		v694 = v930
		goto L120
	} else {
		goto L147
	}
L123:
	;
	v921 = v891
	v922 = v686
	v930 = v712 - int32(1)
	goto L122
L124:
	;
	if int32(0) < v881 {
		v891 = v881
		goto L123
	} else {
		goto L145
	}
L125:
	;
	v848 = int32(0)
	if v824 < v848 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	v815 = base.B2i32(v717 == int32(-1))
	if v718 == v815 {
		goto L115
	} else {
		goto L138
	}
L127:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v722+v714)))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v725+v714)))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v731 = v719
	goto L128
L128:
	;
	v761 = v731 << (uint(int32(2)) % 32)
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v724+v761)))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v761+v728)))
	if v763 < v765 {
		v881 = v731 ^ int32(-1)
		goto L124
	} else {
		goto L130
	}
L129:
	;
	v813 = v677
	goto L126
L130:
	;
	v768 = v731 + int32(1)
	if v765 < v763 {
		v891 = v768
		goto L123
	} else {
		goto L131
	}
L131:
	;
	if v763 != 0 {
		v813 = v768
		goto L126
	} else {
		goto L132
	}
L132:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v761+v675)))
	v776 = v731 << (uint(int32(3)) % 32)
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v727+v776)))
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v776+v729)))
	v781 = F_FunctionCall2Coll(m, v676+v731*int32(28), v774, v778, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v783 = base.I32_wrap_i64(v781)
	if v783 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v821 = v768
	v824 = v783
	goto L125
L135:
	;
	goto L136
L136:
	;
	if v677 != v768 {
		v731 = v768
		goto L128
	} else {
		goto L137
	}
L137:
	;
	goto L129
L138:
	;
	if v717 == int32(-1) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v819 = int32(1)
	goto L141
L140:
	;
	v819 = int32(-1)
	goto L141
L141:
	;
	v821 = v813
	v824 = v819
	goto L125
L142:
	;
	v852 = v848 - v821
	goto L144
L143:
	;
	v852 = v821
	goto L144
L144:
	;
	v881 = v852
	goto L124
L145:
	;
	if v881 == int32(0) {
		goto L116
	} else {
		goto L146
	}
L146:
	;
	v921 = v881
	v922 = v712
	v930 = v694
	goto L122
L147:
	;
	goto L121
L148:
	;
	v1004 = l0
	v1005 = v975
	v1007 = l3
	v1008 = v979
	v1013 = v31
	v1014 = v951
	v1015 = v38 + int32(24)
	v1020 = v36
	goto L114
L149:
	;
	v1004 = l0
	v1005 = v986
	v1007 = l3
	v1008 = v990
	v1013 = v31
	v1014 = v712
	v1015 = v38 + int32(24)
	v1020 = v36
	goto L114
L150:
	;
	v1004 = l0
	v1005 = v997
	v1007 = l3
	v1008 = v1001
	v1013 = v31
	v1014 = v712
	v1015 = v679
	v1020 = v36
	goto L114
L151:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v1037 = v1005 << (uint(int32(2)) % 32)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1037+v1038)))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1041+v1037)))
	v1045 = base.B2i32(v1008 == int32(-1))
	v1046 = int32(0)
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+12)))
	v1048 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v1048 <= v1046 {
		v1110 = v1046
		goto L155
	} else {
		goto L156
	}
L152:
	;
	if int32(0) <= v1235 {
		v1545 = v1013
		goto L4
	} else {
		goto L177
	}
L153:
	;
	v1235 = v1206
	goto L152
L154:
	;
	v1171 = int32(0)
	if v1147 < v1171 {
		goto L171
	} else {
		goto L172
	}
L155:
	;
	if v1008 == int32(-1) {
		goto L165
	} else {
		goto L166
	}
L156:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v1058 = int32(0)
	goto L157
L157:
	;
	v1085 = v1058 << (uint(int32(2)) % 32)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1043+v1085)))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1085+v1051)))
	if v1087 < v1089 {
		v1235 = v1058 ^ int32(-1)
		goto L152
	} else {
		goto L159
	}
L158:
	;
	v1110 = v1048
	goto L155
L159:
	;
	v1092 = v1058 + int32(1)
	if v1089 < v1087 {
		v1206 = v1092
		goto L153
	} else {
		goto L160
	}
L160:
	;
	if v1087 != 0 {
		v1110 = v1092
		goto L155
	} else {
		goto L161
	}
L161:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1085+v1035)))
	v1100 = v1058 << (uint(int32(3)) % 32)
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v1040+v1100)))
	v1104 = *(*int64)(unsafe.Add(mBase, uint32(v1100+v1052)))
	v1105 = F_FunctionCall2Coll(m, v1034+v1058*int32(28), v1098, v1102, v1104)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v1107 = base.I32_wrap_i64(v1105)
	if v1107 != 0 {
		v1144 = v1092
		v1147 = v1107
		goto L154
	} else {
		goto L163
	}
L163:
	;
	if v1092 != v1048 {
		v1058 = v1092
		goto L157
	} else {
		goto L164
	}
L164:
	;
	goto L158
L165:
	;
	v1139 = int32(1)
	goto L167
L166:
	;
	v1139 = int32(-1)
	goto L167
L167:
	;
	if v1045 != v1047 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v1142 = v1139
	goto L170
L169:
	;
	v1142 = int32(0)
	goto L170
L170:
	;
	v1144 = v1110
	v1147 = v1142
	goto L154
L171:
	;
	v1175 = v1171 - v1144
	goto L173
L172:
	;
	v1175 = v1144
	goto L173
L173:
	;
	if v1147 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1177 = v1175
	goto L176
L175:
	;
	v1177 = int32(0)
	goto L176
L176:
	;
	v1206 = v1177
	goto L153
L177:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1015)))
	v1239 = int32(2)
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1238+v1014<<(uint(v1239)%32))+8))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+12))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1244+(v1235^int32(-1))<<(uint(v1239)%32))))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+12))
	v1480 = v1004
	v1483 = v1007
	v1484 = v1242
	v1489 = v1013
	v1490 = v1251
	v1496 = v1020
	goto L16
L178:
	;
	v1252 = l0
	v1255 = l3
	v1256 = v979
	v1261 = v31
	v1268 = v36
	goto L112
L179:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1327 == int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+4))
	if v1330 <= int32(0) {
		v1545 = v31
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v1345 = v5
	v1346 = v5
	goto L182
L182:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+12))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1361+v1346<<(uint(int32(2))%32))))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+36))
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365)+32)))
	if v1367 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v1545 = v31
	goto L4
L184:
	;
	v1477 = v1346 + int32(1)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+4))
	if v1477 < v1478 {
		v1345 = v1460
		v1346 = v1477
		goto L182
	} else {
		goto L203
	}
L185:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v1373 = v1371 - int32(1)
	if v1373 < int32(0) {
		v1460 = v1345
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	if v1445 != int32(-1) {
		v1480 = l0
		v1483 = l3
		v1484 = v1445
		v1489 = v31
		v1490 = v1366
		v1496 = v36
		goto L16
	} else {
		goto L202
	}
L188:
	;
	v1376 = *(*int64)(unsafe.Add(mBase, uint32(v1365)+24))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v1380 = v1373
	v1384 = int32(-1)
	v1391 = v1345
	goto L189
L189:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v1412 = int32(2)
	v1413 = base.I32_div_s(v1380+v1384+int32(1), v1412)
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1408+v1413<<(uint(v1412)%32))))
	v1418 = *(*int64)(unsafe.Add(mBase, uint32(v1417)))
	v1419 = F_FunctionCall2Coll(m, v1378, v1407, v1418, v1376)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L193
	}
L190:
	;
	v1435 = int32(0)
	if base.B2i32(v1434 == v1435)|base.B2i32(v1433 < v1435) != 0 {
		v1460 = v1434
		goto L184
	} else {
		goto L201
	}
L191:
	;
	goto L190
L192:
	;
	if v1429 < v1428 {
		v1380 = v1428
		v1384 = v1429
		v1391 = v1430
		goto L189
	} else {
		goto L200
	}
L193:
	;
	v1421 = base.I32_wrap_i64(v1419)
	if v1421 <= int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if v1421 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v1428 = v1413 - int32(1)
	v1429 = v1384
	v1430 = v1391
	goto L192
L197:
	;
	v1428 = v1380
	v1429 = v1413
	v1430 = int32(0)
	goto L192
L198:
	;
	goto L199
L199:
	;
	v1433 = v1413
	v1434 = int32(1)
	goto L191
L200:
	;
	v1433 = v1429
	v1434 = v1430
	goto L191
L201:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1440+v1433<<(uint(int32(2))%32))))
	v1480 = l0
	v1483 = l3
	v1484 = v1444
	v1489 = v31
	v1490 = v1366
	v1496 = v36
	goto L16
L202:
	;
	v1460 = v1345
	goto L184
L203:
	;
	goto L183
L204:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1496)+8))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1515+v1484<<(uint(int32(2))%32))))
	v1520 = F_get_rel_name(m, v1519)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1489)+20)) = v1520
	*(*int32)(unsafe.Add(mBase, uint32(v1489)+16)) = v1480
	F_errmsg(m, int32(_a_F_check_new_partition_bound_8), v1489+int32(16))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_parser_errposition(m, v1483, v1490)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3231), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
			if v15&int32(1) != 0 {
				m.G0 = v7 + int32(16)
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v19 = F_lappend_oid(m, v18, v11)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
					F_get_partition_ancestors_worker(m, l0, v11, l2)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
	var v19 int32
	_ = v19
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v22 = int32(3)
		F_ScanKeyInit(m, v8+int32(56), v22, v22, int32(65), int64(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v32 = F_systable_beginscan(m, l0, int32(2680), int32(1), int32(0), int32(2), v8)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_systable_getnext(m, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
						v38 = v36 + v37
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
						if v39 == int32(1) {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v42)
						} else {
						}
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v46 = v44
					} else {
						v46 = v4
					}
					F_systable_endscan(m, v32)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(112)
						return v46
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
										*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v52
										m.G0 = v11 + int32(16)
										return v51
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_transformPartitionBoundValue_0), int32(0))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_transformPartitionBoundValue_1), int32(_a_F_transformPartitionBoundValue_2), int32(_a_F_transformPartitionBoundValue_3))
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
						*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v52
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
								F_errmsg(m, int32(_a_F_transformPartitionBoundValue_4), v11)
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
										F_errfinish(m, int32(_a_F_transformPartitionBoundValue_1), int32(_a_F_transformPartitionBoundValue_5), int32(_a_F_transformPartitionBoundValue_3))
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v116 int32
	_ = v116
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
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
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
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
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	v29 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(l3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L3
	} else {
		goto L118
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L3
	} else {
		goto L115
	}
L3:
	;
	return int32(0)
L4:
	;
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = v34 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	if v33 != v37 {
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
	v540 = m.ExcPending
	if v540 != 0 {
		goto L3
	} else {
		goto L112
	}
L8:
	;
	m.G0 = v25 + int32(48)
	return v532
L9:
	;
	v141 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+4)))
	v142 = F_SearchSysCache1(m, int32(19), v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L21
	}
L10:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L20
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v39 != l4 {
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
	v65 = int32(0)
	goto L14
L14:
	;
	v72 = v65 << (uint(int32(1)) % 32)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(22)+v72))))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5+v72))))
	if v74 != v76 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72+(l1+int32(86))))))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6+v72))))
	if v79 != v81 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v84 = v65 << (uint(int32(2)) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(152)+v84)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l7+v84)))
	if v86 != v88 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v91 = v65 + int32(1)
	if l4 != v91 {
		v65 = v91
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v532 = int32(0)
	goto L8
L21:
	;
	if v142 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v148 = v146 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+75)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+75)))
	if v149 != v150 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+92))
	if v152 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L34
	}
L25:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L32
	}
L26:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+73)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+73)))
	if v153 != v154 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+74)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+74)))
	if v156 != v157 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+100)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+100)))
	if v159 != v160 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+101)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+101)))
	if v162 != v163 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+102)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+102)))
	if v165 == v166 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	F_ReleaseCatCache(m, v142)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v532 = int32(0)
	goto L8
L34:
	;
	F_ReleaseCatCache(m, v142)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v178 = m.G0
	v180 = v178 - int32(176)
	m.G0 = v180
	v184 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(l3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L39
	}
L36:
	;
	v532 = int32(1)
	goto L8
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L3
	} else {
		goto L109
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L3
	} else {
		goto L106
	}
L39:
	;
	if v184 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+22)))
	v188 = v186 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+75)))
	v191 = base.I64_extend_i32_u(v177)
	v192 = F_SearchSysCache1(m, int32(19), v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
	v469 = m.ExcPending
	if v469 != 0 {
		goto L3
	} else {
		goto L103
	}
L43:
	;
	if v192 == int32(0) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+22)))
	v198 = v196 + v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+80))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v198)+96))
	v201 = F_get_rel_relkind(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	if v201 == int32(112) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v207 = F_table_open(m, int32(2606), int32(2))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+76)))
	if v403 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L49:
	;
	v210 = v180 + int32(120)
	F_ScanKeyInit(m, v210, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v199))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v218 = int32(1)
	v221 = F_systable_beginscan(m, v207, int32(2665), v218, int32(0), v218, v210)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v223 = F_new_object_addresses(m)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v225 = F_systable_getnext(m, v221)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	if v225 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v228 = v225
	goto L57
L55:
	;
	goto L56
L56:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L76
	}
L57:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+22)))
	v251 = v249 + v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+92))
	if v177 == v252 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+108)) = int32(2606)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+116)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+112)) = v256
	F_add_exact_object_address(m, v180+int32(108), v223)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v345 = F_systable_getnext(m, v221)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L74
	}
L62:
	;
	v264 = int32(2606)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	F_deleteDependencyRecordsForSpecific(m, v264, v265, int32(105), v264, v177)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v271 = v180 + int32(48)
	v275 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v251))))
	F_ScanKeyInit(m, v271, int32(11), int32(3), int32(184), v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v279 = int32(1)
	v282 = F_systable_beginscan(m, l10, int32(2699), v279, int32(0), v279, v271)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	goto L66
L66:
	;
	v306 = F_systable_getnext(m, v282)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L68
	}
L67:
	;
	F_systable_endscan(m, v282)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L73
	}
L68:
	;
	if v306 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+108)) = int32(2620)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306)+16))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+22)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v310+v311)))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+116)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+112)) = v313
	F_add_exact_object_address(m, v180+int32(108), v223)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
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
	if v345 != 0 {
		v228 = v345
		goto L57
	} else {
		goto L75
	}
L75:
	;
	goto L58
L76:
	;
	F_performMultipleDeletions(m, v223, int32(0), int32(1))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_systable_endscan(m, v221)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_relation_close(m, v207, int32(2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	goto L48
L80:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+76)))
	v409 = v406 ^ int32(1)
	goto L82
L81:
	;
	v409 = int32(0)
	goto L82
L82:
	;
	F_ReleaseCatCache(m, v192)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_ReleaseCatCache(m, v184)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	F_DropForeignKeyConstraintTriggers(m, l10, v177, v200, v199)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ConstraintSetParentConstraint(m, v177, l3, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	if v189&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_GetForeignKeyCheckTriggers(m, l10, v177, v200, v199, v180+int32(120), v180+int32(48))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L93
	}
L90:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v180)+120))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v427, l8, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v180)+48))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v431, l9, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v409&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v441 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L3
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	m.G0 = v180 + int32(176)
	goto L36
L97:
	;
	v444 = F_SearchSysCache1(m, int32(19), v191)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	if v444 == int32(0) {
		goto L37
	} else {
		goto L99
	}
L99:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+22)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v448+v449)+96))
	F_QueueFKConstraintValidation(m, l0, v441, l2, v451, v444, int32(4))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	F_ReleaseCatCache(m, v444)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	F_relation_close(m, v441, int32(3))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = l3
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v180)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_2), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v177
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v180+int32(16))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_4), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v180)+32)) = v177
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v180+int32(32))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_5), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l3
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v25)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_6), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
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
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v554
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v25+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_8), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
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
	v572 = m.ExcPending
	if v572 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v574 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v148 + v574
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v36 + v574
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v573 + v574
	F_errmsg(m, int32(_a_F_tryAttachPartitionForeignKey_9), v25+int32(32))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_10), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
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
