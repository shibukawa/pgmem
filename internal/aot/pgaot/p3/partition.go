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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
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
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
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
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int64
	_ = v647
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(656)
	m.G0 = v17
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_RemoveInheritance(m, l1, l0, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v23 = F_new_object_addresses(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v26 = v17 + int32(512)
	F_ScanKeyInit(m, v26, int32(2), int32(3), int32(184), v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v34 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(1)
	v40 = F_systable_beginscan(m, v34, int32(2701), v37, int32(0), v37, v26)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v42 = F_systable_getnext(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = v42
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L24
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v88 = F_systable_getnext(m, v40)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L22
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+84))
	if v64 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = int32(2620)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = F_deleteDependencyRecordsForClass(m, v65, v66, v65, int32(80))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v75 = F_deleteDependencyRecordsForClass(m, int32(2620), v72, int32(1259), int32(83))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+240)) = int32(2620)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+244)) = v79
	F_add_exact_object_address(m, v17+int32(240), v23)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L16
L22:
	;
	if v88 != 0 {
		v46 = v88
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	F_performMultipleDeletions(m, v23, int32(0), int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_free_object_addresses(m, v23)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_systable_endscan(m, v40)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v117 = F_RelationGetFKeyList(m, l1)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v483 = F_GetParentedForeignKeyRefs(m, l1)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L100
	}
L30:
	;
	v119 = F_copyObjectImpl(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v119 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_list_free_deep(m, v119)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v127 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if int32(0) < v129 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = int32(0)
	v145 = v5
	goto L40
L38:
	;
	v171 = v5
	goto L39
L39:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if int32(0) < v173 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v135<<(uint(int32(2))%32))))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v153 = F_lappend_oid(m, v145, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	v171 = v153
	goto L39
L42:
	;
	v156 = v135 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v156 < v157 {
		v135 = v156
		v145 = v153
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v189 = v5
	goto L47
L45:
	;
	goto L46
L46:
	;
	F_list_free_deep(m, v119)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L96
	}
L47:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v189<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v197 = F_SearchSysCache1(m, int32(19), v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L52
	}
L48:
	;
	goto L46
L49:
	;
	F_ReleaseCatCache(m, v197)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L94
	}
L50:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v402 = F_table_open(m, v400, int32(6))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L91
	}
L51:
	;
	v349 = v311
	v356 = v311
	goto L86
L52:
	;
	if v197 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+22)))
	v201 = v199 + v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+72)))
	if v202 != int32(102) {
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
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L83
	}
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+92))
	if v205 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v208 = int32(0)
	if v171 == v208 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v246 != 0 {
		goto L49
	} else {
		goto L71
	}
L59:
	;
	v246 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v214 <= int32(0) {
		v240 = v208
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v246 = v240
	goto L58
L63:
	;
	v217 = int32(0)
	if v217 < v214 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v220 = v214
	goto L66
L65:
	;
	v220 = v217
	goto L66
L66:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v223 = int32(0)
	goto L67
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221+v223<<(uint(int32(2))%32))))
	v232 = base.B2i32(v231 == v205)
	if v231 == v205 {
		v240 = v232
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v240 = v232
	goto L62
L69:
	;
	v234 = v223 + int32(1)
	if v234 != v220 {
		v223 = v234
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v248 = int32(0)
	F_ConstraintSetParentConstraint(m, v247, v248, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+20)))
	if v252 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	F_GetForeignKeyCheckTriggers(m, v127, v255, v256, v257, v17+int32(512), v17+int32(240))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_DeconstructFkConstraintRow(m, v197, v17+int32(508), v17+int32(432), v17+int32(368), v17+int32(512), v17+int32(240), v17+int32(112), v17+int32(108), v17+int32(32))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L79
	}
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v17)+512))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v127, v264, int32(0), v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v17)+240))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_TriggerSetParentTrigger(m, v127, v269, int32(0), v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v293 = F_palloc0(m, int32(108))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = int64(438086664353)
	v299 = F_pstrdup(m, v201+int32(4))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v299
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+12)) = uint8(v302)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+13)) = uint8(v304)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+75)))
	v307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+15)) = uint8(v307)
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+14)) = uint8(v306)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+76)))
	v311 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v293)+80)) = v311
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v293)+72)) = v314
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+16)) = uint8(v310)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+86)) = uint8(v317)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+87)) = uint8(v319)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+101)))
	*(*int64)(unsafe.Add(mBase, uint32(v293)+100)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v293)+92)) = v314
	*(*uint8)(unsafe.Add(mBase, uint32(v293)+88)) = uint8(v321)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+508))
	if v311 < v328 {
		goto L51
	} else {
		goto L82
	}
L82:
	;
	goto L50
L83:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v335
	F_errmsg_internal(m, int32(_a_F_DetachPartitionFinalize_0), v17+int32(16))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_DetachPartitionFinalize_1), int32(_a_F_DetachPartitionFinalize_2), int32(_a_F_DetachPartitionFinalize_3))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
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
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17+int32(432)+v349<<(uint(int32(1))%32)))))
	v377 = F_makeString(m, v361+v362<<(uint(int32(4))%32)+v371*int32(100)-int32(76))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L50
L88:
	;
	v379 = F_lappend(m, v356, v377)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+76)) = v379
	v383 = v349 + int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v17)+508))
	if v383 < v384 {
		v349 = v383
		v356 = v379
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v201)+88))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v17)+508))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	v420 = int32(0)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+107)))
	F_addFkRecurseReferenced(m, v293, l1, v402, v404, v405, v406, v17+int32(368), v17+int32(432), v17+int32(512), v17+int32(240), v17+int32(112), v417, v17+int32(32), v420, v420, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_relation_close(m, v402, int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v445 = v189 + int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v445 < v446 {
		v189 = v445
		goto L47
	} else {
		goto L95
	}
L95:
	;
	goto L48
L96:
	;
	if v127 == int32(0) {
		goto L29
	} else {
		goto L97
	}
L97:
	;
	F_relation_close(m, v127, int32(3))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	goto L29
L99:
	;
	v550 = F_RelationGetIndexList(m, l1)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L111
	}
L100:
	;
	if v483 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v487 <= int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v493 = int32(0)
	goto L103
L103:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v505+v493<<(uint(int32(2))%32))))
	v510 = int32(0)
	F_ConstraintSetParentConstraint(m, v509, v510, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L105
	}
L104:
	;
	goto L99
L105:
	;
	v514 = int32(2606)
	v517 = F_deleteDependencyRecordsForClass(m, v514, v509, v514, int32(105))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v521 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+520)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v17)+516)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v17)+512)) = int32(2606)
	F_performDeletion(m, v17+int32(512), v521, v521)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v533 = v493 + int32(1)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v533 < v534 {
		v493 = v533
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	v631 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L131
	}
L111:
	;
	if v550 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	if v554 <= int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v560 = int32(0)
	goto L114
L114:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v550)+12))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v560<<(uint(int32(2))%32))))
	v577 = F_has_superclass(m, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	goto L110
L116:
	;
	if v577 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v580 = F_get_partition_parent(m, v576, int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v612 = v560 + int32(1)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	if v612 < v613 {
		v560 = v612
		goto L114
	} else {
		goto L130
	}
L120:
	;
	v583 = F_index_open(m, v576, int32(8))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	F_IndexSetParentIndex(m, v583, int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v589 = F_get_relation_idx_constraint_oid(m, v588, v576)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v592 = F_get_relation_idx_constraint_oid(m, v591, v580)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v594 = int32(0)
	if base.B2i32(v592 == v594)|base.B2i32(v589 == v594) == v594 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v601 = int32(0)
	F_ConstraintSetParentConstraint(m, v589, v601, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	F_relation_close(m, v583, int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
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
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v636 = F_SearchSysCacheCopy(m, int32(57), v634, int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v636 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v639 = v17 + int32(512)
	v640 = int32(0)
	base.MemoryFill(m, v639, v640, int32(136))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+144)) = uint16(v640)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+272)) = uint16(v640)
	v647 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+264)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+256)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+248)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+240)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v17)+136)) = v647
	v663 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+145)) = uint8(v663)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+273)) = uint8(v663)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v631)+52))
	v672 = F_heap_modify_tuple(m, v636, v667, v639, v17+int32(240), v17+int32(112))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
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
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L167
	}
L136:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v672)+16))
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+22)))
	v677 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v674+v675)+131)) = uint8(v677)
	F_CatalogTupleUpdate(m, v631, v672+int32(4), v672)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v672)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_relation_close(m, v631, int32(3))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v689 = int32(*(*int16)(unsafe.Add(mBase, uint32(v688)+120)))
	if int32(0) < v689 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v695 = int32(0)
	v698 = v688
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
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v714 = v707 + v708<<(uint(int32(4))%32) + v695*int32(100)
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+111)))
	if v715 != 0 {
		v732 = v698
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	v735 = v695 + int32(1)
	v736 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732)+120)))
	if v735 < v736 {
		v695 = v735
		v698 = v732
		goto L143
	} else {
		goto L149
	}
L146:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+int32(20))+89)))
	if v718 == int32(0) {
		v732 = v698
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v727 = int32(1)
	F_ATExecDropIdentity(m, v17+int32(432), l1, v714+int32(24), int32(0), int32(8), v727, v727)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v732 = v731
	goto L145
L149:
	;
	goto L144
L150:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L157
	}
L151:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l3 == v754 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_update_default_partition_oid(m, v756, int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
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
	v761 = m.ExcPending
	if v761 != 0 {
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
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764)+119)))
	if v765 != int32(112) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	m.G0 = v17 + int32(656)
	return
L159:
	;
	v768 = int32(0)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v772 = F_find_all_inheritors(m, v769, int32(8), v768)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	if v772 == int32(0) {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v772)+4))
	if v776 <= int32(0) {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	v781 = v768
	goto L163
L163:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v772)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793+v781<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v797)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L165
	}
L164:
	;
	goto L158
L165:
	;
	v801 = v781 + int32(1)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v772)+4))
	if v801 < v802 {
		v781 = v801
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v825
	F_errmsg_internal(m, int32(_a_F_DetachPartitionFinalize_4), v17)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_DetachPartitionFinalize_1), int32(_a_F_DetachPartitionFinalize_5), int32(_a_F_DetachPartitionFinalize_3))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
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
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_InitPartitionPruneContext[0]))
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
	v70 = int32(0)
	if base.B2i32(v67 != int32(377))|base.B2i32(v16 <= v70) == v70 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = v66
	v79 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v149 = v56 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v149 < v150 {
		v56 = v149
		goto L10
	} else {
		goto L37
	}
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v88 = F_bms_is_member(m, v79, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v90 = int32(0)
	if v88|base.B2i32(v78 == v90) == v90 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v96 != int32(7) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v131 = v78
	goto L23
L23:
	;
	v135 = v79 + int32(1)
	if v135 != v16 {
		v78 = v131
		v79 = v135
		goto L18
	} else {
		goto L36
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
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
	v121 = v78 + int32(4)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if base.Ui32(v121) < base.Ui32(v124+v125<<(uint(int32(2))%32)) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v111 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v99*v16<<(uint(v111)%32)+v79<<(uint(v111)%32)))) = v109
	goto L26
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v104 = F_ExecInitExprWithParams(m, v95, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v107 = F_ExecInitExpr(m, v95, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v109 = v104
	goto L27
L32:
	;
	v109 = v107
	goto L27
L33:
	;
	v130 = v121
	goto L35
L34:
	;
	v130 = int32(0)
	goto L35
L35:
	;
	v131 = v130
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
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
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v220 int32
	_ = v220
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
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
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
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
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v531 int32
	_ = v531
	var v560 int32
	_ = v560
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
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
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1171 int32
	_ = v1171
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1272 int32
	_ = v1272
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1353 int32
	_ = v1353
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1420 int32
	_ = v1420
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1454 int32
	_ = v1454
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1502 int32
	_ = v1502
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
	m.G0 = v1502 + int32(112)
	return
L5:
	;
	if v37 == int32(0) {
		v1502 = v30
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
		v1502 = v30
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
		v1502 = v30
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
	F_errmsg(m, int32(_a_F_check_new_partition_bound_0), v30)
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
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2922), int32(_a_F_check_new_partition_bound_2))
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
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L204
	}
L17:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v1286 <= int32(0) {
		v1502 = v30
		goto L4
	} else {
		goto L179
	}
L18:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v453 = F_make_one_partition_rbound(m, v32, int32(-1), v451, int32(1))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L80
	}
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v77 <= int32(0) {
		v1502 = v30
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v84 = v82 - int32(1)
	if int32(0) <= v84 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L74
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L68
	}
L23:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v349 <= v80 {
		goto L59
	} else {
		goto L60
	}
L24:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v308 = v220 << (uint(int32(2)) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+v308)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = base.I32_rem_s(v81, v311)
	if v312 != 0 {
		goto L22
	} else {
		goto L56
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v94 = int32(-1)
	v99 = v84
	v102 = v5
	goto L28
L26:
	;
	goto L27
L27:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v274 = base.I32_rem_s(v273, v81)
	if v274 == int32(0) {
		goto L23
	} else {
		goto L49
	}
L28:
	;
	v117 = int32(2)
	v118 = base.I32_div_s(v99+v102, v117)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v87+v118<<(uint(v117)%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123 < v81 {
		v190 = v118
		v195 = v99
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if int32(0) <= v220 {
		goto L24
	} else {
		goto L48
	}
L30:
	;
	goto L29
L31:
	;
	if v190 < v195 {
		v94 = v190
		v99 = v195
		v102 = v190 + int32(1)
		goto L28
	} else {
		goto L47
	}
L32:
	;
	v130 = v118
	v131 = v123
	v135 = v99
	v137 = v122
	goto L33
L33:
	;
	if v131 <= v81 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v190 = v178
	v195 = v173
	goto L31
L35:
	;
	v177 = int32(2)
	v178 = base.I32_div_s(v173+v102, v177)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v87+v178<<(uint(v177)%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v81 <= v183 {
		v130 = v178
		v131 = v183
		v135 = v173
		v137 = v182
		goto L33
	} else {
		goto L46
	}
L36:
	;
	v220 = v94
	goto L30
L37:
	;
	v153 = base.B2i32(v131 != v81)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v156 = v153 | base.B2i32(v154 == v80)
	if v153|base.B2i32(v154 <= v80) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v166 = v130 - int32(1)
	if v94 < v166 {
		v173 = v166
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v162 = v130 - int32(1)
	if v156|base.B2i32(v162 <= v94) != 0 {
		goto L36
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v156 != 0 {
		v220 = v130
		goto L30
	} else {
		goto L44
	}
L43:
	;
	v173 = v162
	goto L35
L44:
	;
	v190 = v130
	v195 = v135
	goto L31
L45:
	;
	goto L36
L46:
	;
	goto L34
L47:
	;
	v220 = v190
	goto L30
L48:
	;
	goto L27
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v291 = F_get_rel_name(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v288
	F_errdetail(m, int32(_a_F_check_new_partition_bound_4), v30+int32(32))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2975), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	v314 = v220 + int32(1)
	if v82 <= v314 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	v317 = v314 << (uint(int32(2)) % 32)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v306+v317)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v321 = base.I32_rem_s(v320, v81)
	if v321 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	goto L23
L59:
	;
	v351 = base.I32_rem_s(v80, v349)
	v352 = v351
	goto L61
L60:
	;
	v352 = v80
	goto L61
L61:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v355 = v352
	goto L62
L62:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v353+v355<<(uint(int32(2))%32))))
	if v384 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v1502 = v30
	goto L4
L64:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1438 = l0
	v1441 = l3
	v1443 = v387
	v1444 = v384
	v1447 = v30
	v1454 = v35
	goto L16
L65:
	;
	goto L66
L66:
	;
	v388 = v355 + v81
	if v388 < v349 {
		v355 = v388
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
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v402+v308)))
	v405 = F_get_rel_name(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v401
	F_errdetail(m, int32(_a_F_check_new_partition_bound_5), v30-int32(-64))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(2995), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
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
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_check_new_partition_bound_3), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v432+v317)))
	v435 = F_get_rel_name(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v431
	F_errdetail(m, int32(_a_F_check_new_partition_bound_4), v30+int32(48))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3016), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
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
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v458 = F_make_one_partition_rbound(m, v32, int32(-1), v456, int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+12)))
	v461 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v461 <= int32(0) {
		v531 = v5
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v650 <= int32(0) {
		v1502 = v30
		goto L4
	} else {
		goto L109
	}
L83:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582+v560<<(uint(int32(2))%32)-int32(4))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L101
	}
L84:
	;
	if v460 != 0 {
		v628 = v5
		goto L82
	} else {
		goto L99
	}
L85:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v458)+8))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v474 = v5
	goto L86
L86:
	;
	v498 = v474 << (uint(int32(2)) % 32)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v466+v498)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v498+v464)))
	if v500 < v502 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v531 = v461
	goto L84
L88:
	;
	v628 = v474 ^ int32(-1)
	goto L82
L89:
	;
	goto L90
L90:
	;
	v507 = v474 + int32(1)
	if v502 < v500 {
		v560 = v507
		goto L83
	} else {
		goto L91
	}
L91:
	;
	if v500 != 0 {
		v531 = v507
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v498+v468)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v498+v467)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v498+v465)))
	v518 = F_FunctionCall2Coll(m, v469+v474*int32(28), v513, v515, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v518 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if int32(0) <= v518 {
		v560 = v507
		goto L83
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v507 != v461 {
		v474 = v507
		goto L86
	} else {
		goto L98
	}
L97:
	;
	v628 = v474 ^ int32(-1)
	goto L82
L98:
	;
	goto L87
L99:
	;
	if v531 <= int32(0) {
		v628 = v531
		goto L82
	} else {
		goto L100
	}
L100:
	;
	v560 = v531
	goto L83
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = l0
	F_errmsg(m, int32(_a_F_check_new_partition_bound_6), v30+int32(96))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v603 = F_get_range_partbound_string(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v606 = F_get_range_partbound_string(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v603
	F_errdetail(m, int32(_a_F_check_new_partition_bound_7), v30+int32(80))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v588)+12))
	F_parser_errposition(m, l3, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3132), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
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
	v653 = int32(-1)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v656 = v654 - int32(1)
	if int32(0) <= v656 {
		goto L117
	} else {
		goto L118
	}
L110:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1283)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	v1438 = v1256
	v1441 = v1259
	v1443 = v1285
	v1444 = v1262
	v1447 = v1265
	v1454 = v1272
	goto L16
L111:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+12))
	v1248 = v924 >> (uint(int32(31)) % 32)
	v1256 = l0
	v1259 = l3
	v1262 = v952
	v1265 = v30
	v1272 = v35
	v1283 = v1246 + (v924^v1248-v1248)<<(uint(int32(2))%32) - int32(4)
	goto L110
L112:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+12))
	v1256 = v1216
	v1259 = v1219
	v1262 = v1222
	v1265 = v1225
	v1272 = v1232
	v1283 = v1244
	goto L110
L113:
	;
	if v924 != 0 {
		goto L111
	} else {
		goto L178
	}
L114:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v1004 <= v978 {
		v1502 = v986
		goto L4
	} else {
		goto L151
	}
L115:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v970 = v695 + int32(1)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v968+v970<<(uint(int32(2))%32))))
	if int32(0) <= v974 {
		v1216 = l0
		v1219 = l3
		v1222 = v974
		v1225 = v30
		v1232 = v35
		goto L112
	} else {
		goto L150
	}
L116:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v959 = v695 + int32(1)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v957+v959<<(uint(int32(2))%32))))
	if int32(0) <= v963 {
		v1216 = l0
		v1219 = l3
		v1222 = v963
		v1225 = v30
		v1232 = v35
		goto L112
	} else {
		goto L149
	}
L117:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	v663 = v37 + int32(24)
	v670 = v656
	v674 = v653
	goto L120
L118:
	;
	v924 = v628
	v929 = v653
	goto L119
L119:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v948 = v929 + int32(1)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v946+v948<<(uint(int32(2))%32))))
	if int32(0) <= v952 {
		goto L113
	} else {
		goto L148
	}
L120:
	;
	v694 = int32(2)
	v695 = base.I32_div_s(v670+v674+int32(1), v694)
	v697 = v695 << (uint(v694) % 32)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v697+v698)))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+12)))
	v702 = int32(0)
	if v661 <= v702 {
		v791 = v702
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v924 = v896
	v929 = v901
	goto L119
L122:
	;
	if v901 < v897 {
		v670 = v897
		v674 = v901
		goto L120
	} else {
		goto L147
	}
L123:
	;
	v896 = v867
	v897 = v695 - int32(1)
	v901 = v674
	goto L122
L124:
	;
	if int32(0) < v857 {
		v867 = v857
		goto L123
	} else {
		goto L145
	}
L125:
	;
	v825 = int32(0)
	if v799 < v825 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	v793 = base.B2i32(v700 == int32(-1))
	if v701 == v793 {
		goto L115
	} else {
		goto L138
	}
L127:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v705+v697)))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v708+v697)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v717 = v702
	goto L128
L128:
	;
	v743 = v717 << (uint(int32(2)) % 32)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v707+v743)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743+v711)))
	if v745 < v747 {
		v857 = v717 ^ int32(-1)
		goto L124
	} else {
		goto L130
	}
L129:
	;
	v791 = v661
	goto L126
L130:
	;
	v750 = v717 + int32(1)
	if v747 < v745 {
		v867 = v750
		goto L123
	} else {
		goto L131
	}
L131:
	;
	if v745 != 0 {
		v791 = v750
		goto L126
	} else {
		goto L132
	}
L132:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v743+v659)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v743+v710)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v743+v712)))
	v761 = F_FunctionCall2Coll(m, v660+v717*int32(28), v756, v758, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	if v761 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v799 = v761
	v802 = v750
	goto L125
L135:
	;
	goto L136
L136:
	;
	if v661 != v750 {
		v717 = v750
		goto L128
	} else {
		goto L137
	}
L137:
	;
	goto L129
L138:
	;
	if v700 == int32(-1) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v797 = int32(1)
	goto L141
L140:
	;
	v797 = int32(-1)
	goto L141
L141:
	;
	v799 = v797
	v802 = v791
	goto L125
L142:
	;
	v829 = v825 - v802
	goto L144
L143:
	;
	v829 = v802
	goto L144
L144:
	;
	v857 = v829
	goto L124
L145:
	;
	if v857 == int32(0) {
		goto L116
	} else {
		goto L146
	}
L146:
	;
	v896 = v857
	v897 = v670
	v901 = v695
	goto L122
L147:
	;
	goto L121
L148:
	;
	v977 = l0
	v978 = v948
	v980 = l3
	v983 = v952
	v985 = v929
	v986 = v30
	v992 = v37 + int32(24)
	v993 = v35
	goto L114
L149:
	;
	v977 = l0
	v978 = v959
	v980 = l3
	v983 = v963
	v985 = v695
	v986 = v30
	v992 = v37 + int32(24)
	v993 = v35
	goto L114
L150:
	;
	v977 = l0
	v978 = v970
	v980 = l3
	v983 = v974
	v985 = v695
	v986 = v30
	v992 = v663
	v993 = v35
	goto L114
L151:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v1009 = v978 << (uint(int32(2)) % 32)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1009+v1010)))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1013+v1009)))
	v1017 = base.B2i32(v983 == int32(-1))
	v1018 = int32(0)
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+12)))
	v1021 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v1021 <= v1018 {
		v1078 = v1018
		goto L155
	} else {
		goto L156
	}
L152:
	;
	if int32(0) <= v1199 {
		v1502 = v986
		goto L4
	} else {
		goto L177
	}
L153:
	;
	v1199 = v1171
	goto L152
L154:
	;
	v1137 = int32(0)
	if v1114 < v1137 {
		goto L171
	} else {
		goto L172
	}
L155:
	;
	if v983 == int32(-1) {
		goto L165
	} else {
		goto L166
	}
L156:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v458)+8))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v1030 = v1018
	goto L157
L157:
	;
	v1056 = v1030 << (uint(int32(2)) % 32)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1015+v1056)))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1024)))
	if v1058 < v1060 {
		v1199 = v1030 ^ int32(-1)
		goto L152
	} else {
		goto L159
	}
L158:
	;
	v1078 = v1021
	goto L155
L159:
	;
	v1063 = v1030 + int32(1)
	if v1060 < v1058 {
		v1171 = v1063
		goto L153
	} else {
		goto L160
	}
L160:
	;
	if v1058 != 0 {
		v1078 = v1063
		goto L155
	} else {
		goto L161
	}
L161:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1007)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1012)))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1025)))
	v1074 = F_FunctionCall2Coll(m, v1006+v1030*int32(28), v1069, v1071, v1073)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v1074 != 0 {
		v1111 = v1063
		v1114 = v1074
		goto L154
	} else {
		goto L163
	}
L163:
	;
	if v1063 != v1021 {
		v1030 = v1063
		goto L157
	} else {
		goto L164
	}
L164:
	;
	goto L158
L165:
	;
	v1106 = int32(1)
	goto L167
L166:
	;
	v1106 = int32(-1)
	goto L167
L167:
	;
	if v1017 != v1020 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v1109 = v1106
	goto L170
L169:
	;
	v1109 = int32(0)
	goto L170
L170:
	;
	v1111 = v1078
	v1114 = v1109
	goto L154
L171:
	;
	v1141 = v1137 - v1111
	goto L173
L172:
	;
	v1141 = v1111
	goto L173
L173:
	;
	if v1114 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1143 = v1141
	goto L176
L175:
	;
	v1143 = int32(0)
	goto L176
L176:
	;
	v1171 = v1143
	goto L153
L177:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v1203 = int32(2)
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1202+v985<<(uint(v1203)%32))+8))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+12))
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1208+(v1199^int32(-1))<<(uint(v1203)%32))))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+12))
	v1438 = v977
	v1441 = v980
	v1443 = v1215
	v1444 = v1206
	v1447 = v986
	v1454 = v993
	goto L16
L178:
	;
	v1216 = l0
	v1219 = l3
	v1222 = v952
	v1225 = v30
	v1232 = v35
	goto L112
L179:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1289 == int32(0) {
		v1502 = v30
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+4))
	if v1292 <= int32(0) {
		v1502 = v30
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v1305 = v5
	v1308 = v5
	goto L182
L182:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+12))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1322+v1305<<(uint(int32(2))%32))))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+28))
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326)+24)))
	if v1328 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v1502 = v30
	goto L4
L184:
	;
	v1435 = v1305 + int32(1)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+4))
	if v1435 < v1436 {
		v1305 = v1435
		v1308 = v1420
		goto L182
	} else {
		goto L203
	}
L185:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1334 = v1332 - int32(1)
	if v1334 < int32(0) {
		v1420 = v1308
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v1404 != int32(-1) {
		v1438 = l0
		v1441 = l3
		v1443 = v1327
		v1444 = v1404
		v1447 = v30
		v1454 = v35
		goto L16
	} else {
		goto L202
	}
L188:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+20))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1341 = v1334
	v1344 = int32(-1)
	v1353 = v1308
	goto L189
L189:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v1372 = int32(2)
	v1373 = base.I32_div_s(v1341+v1344+int32(1), v1372)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1368+v1373<<(uint(v1372)%32))))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	v1379 = F_FunctionCall2Coll(m, v1339, v1367, v1378, v1337)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L193
	}
L190:
	;
	v1394 = int32(0)
	if base.B2i32(v1393 == v1394)|base.B2i32(v1392 < v1394) != 0 {
		v1420 = v1393
		goto L184
	} else {
		goto L201
	}
L191:
	;
	goto L190
L192:
	;
	if v1388 < v1387 {
		v1341 = v1387
		v1344 = v1388
		v1353 = v1389
		goto L189
	} else {
		goto L200
	}
L193:
	;
	if v1379 <= int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if v1379 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v1387 = v1373 - int32(1)
	v1388 = v1344
	v1389 = v1353
	goto L192
L197:
	;
	v1387 = v1341
	v1388 = v1373
	v1389 = int32(0)
	goto L192
L198:
	;
	goto L199
L199:
	;
	v1392 = v1373
	v1393 = int32(1)
	goto L191
L200:
	;
	v1392 = v1388
	v1393 = v1389
	goto L191
L201:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1399+v1392<<(uint(int32(2))%32))))
	v1438 = l0
	v1441 = l3
	v1443 = v1327
	v1444 = v1403
	v1447 = v30
	v1454 = v35
	goto L16
L202:
	;
	v1420 = v1308
	goto L184
L203:
	;
	goto L183
L204:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+8))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1472+v1444<<(uint(int32(2))%32))))
	v1477 = F_get_rel_name(m, v1476)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1447)+20)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v1447)+16)) = v1438
	F_errmsg(m, int32(_a_F_check_new_partition_bound_8), v1447+int32(16))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_parser_errposition(m, v1441, v1443)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_check_new_partition_bound_1), int32(3239), int32(_a_F_check_new_partition_bound_2))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
						v45 = v43
					} else {
						v45 = v4
					}
					F_systable_endscan(m, v31)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(96)
						return v45
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
	var v63 int32
	_ = v63
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
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
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
	v555 = m.ExcPending
	if v555 != 0 {
		goto L3
	} else {
		goto L118
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
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
	v526 = m.ExcPending
	if v526 != 0 {
		goto L3
	} else {
		goto L112
	}
L8:
	;
	m.G0 = v24 + int32(48)
	return v518
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
	v63 = int32(0)
	goto L14
L14:
	;
	v69 = v63 << (uint(int32(1)) % 32)
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
	v81 = v63 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(152)+v81)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l7+v81)))
	if v83 != v85 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v88 = v63 + int32(1)
	if l4 != v88 {
		v63 = v88
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v518 = int32(0)
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
	v518 = int32(0)
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
	v518 = int32(1)
	goto L8
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
	} else {
		goto L109
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
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
	v456 = m.ExcPending
	if v456 != 0 {
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
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+76)))
	if v390 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L49:
	;
	v203 = v175 + int32(96)
	F_ScanKeyInit(m, v203, int32(9), int32(3), int32(184), v192)
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
	v213 = F_systable_beginscan(m, v200, int32(2665), v210, int32(0), v210, v203)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v215 = F_new_object_addresses(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v217 = F_systable_getnext(m, v213)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	if v217 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v220 = v217
	goto L57
L55:
	;
	goto L56
L56:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L76
	}
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+22)))
	v242 = v240 + v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+92))
	if v172 == v243 {
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
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+88)) = v247
	F_add_exact_object_address(m, v175+int32(84), v215)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v334 = F_systable_getnext(m, v213)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L74
	}
L62:
	;
	v255 = int32(2606)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	F_deleteDependencyRecordsForSpecific(m, v255, v256, int32(105), v255, v172)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v262 = v175 + int32(36)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	F_ScanKeyInit(m, v262, int32(11), int32(3), int32(184), v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v270 = int32(1)
	v273 = F_systable_beginscan(m, l10, int32(2699), v270, int32(0), v270, v262)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	goto L66
L66:
	;
	v296 = F_systable_getnext(m, v273)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L68
	}
L67:
	;
	F_systable_endscan(m, v273)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L3
	} else {
		goto L73
	}
L68:
	;
	if v296 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+84)) = int32(2620)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+22)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300+v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+88)) = v303
	F_add_exact_object_address(m, v175+int32(84), v215)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	if v334 != 0 {
		v220 = v334
		goto L57
	} else {
		goto L75
	}
L75:
	;
	goto L58
L76:
	;
	F_performMultipleDeletions(m, v215, int32(0), int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_systable_endscan(m, v213)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_relation_close(m, v200, int32(2))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	goto L48
L80:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+76)))
	v396 = v393 ^ int32(1)
	goto L82
L81:
	;
	v396 = int32(0)
	goto L82
L82:
	;
	F_ReleaseCatCache(m, v185)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_ReleaseCatCache(m, v178)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	F_DropForeignKeyConstraintTriggers(m, l10, v172, v193, v192)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ConstraintSetParentConstraint(m, v172, l3, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v413 = m.ExcPending
	if v413 != 0 {
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
	v423 = m.ExcPending
	if v423 != 0 {
		goto L3
	} else {
		goto L93
	}
L90:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v175)+96))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v414, l8, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v175)+36))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_TriggerSetParentTrigger(m, l10, v418, l9, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v396&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v428 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
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
	v431 = F_SearchSysCache1(m, int32(19), v172)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	if v431 == int32(0) {
		goto L37
	} else {
		goto L99
	}
L99:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431)+16))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+22)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+v436)+96))
	F_QueueFKConstraintValidation(m, l0, v428, l2, v438, v431, int32(4))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	F_ReleaseCatCache(m, v431)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	F_relation_close(m, v428, int32(3))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v175)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_2), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v175+int32(16))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_4), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v175+int32(32))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_5), int32(_a_F_tryAttachPartitionForeignKey_3))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v24)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_6), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
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
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v540
	F_errmsg_internal(m, int32(_a_F_tryAttachPartitionForeignKey_0), v24+int32(16))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_8), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
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
	v558 = m.ExcPending
	if v558 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v560 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v143 + v560
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v34 + v560
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v559 + v560
	F_errmsg(m, int32(_a_F_tryAttachPartitionForeignKey_9), v24+int32(32))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_tryAttachPartitionForeignKey_1), int32(_a_F_tryAttachPartitionForeignKey_10), int32(_a_F_tryAttachPartitionForeignKey_7))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
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
