package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateSchemaCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int64
	_ = v234
	var v237 int32
	_ = v237
	var v239 int64
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v336 int64
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v762 int32
	_ = v762
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(128)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[0]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(124)))) = v40
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(120)))) = v43
	goto L1
L1:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v45 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v47 = F_get_rolespec_oid(m, v45, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+124))
	v50 = v49
	goto L2
L6:
	;
	return
L7:
	;
	v50 = v47
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L6
	} else {
		goto L165
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L162
	}
L10:
	;
	v55 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(v50))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v69 = v32
	goto L12
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[3]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v30)+124))
	v75 = F_object_aclcheck(m, int32(1262), v72, v73, int64(512))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L17
	}
L13:
	;
	if v55 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v64 = F_pstrdup(m, v59+v60+int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	F_ReleaseCatCache(m, v55)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v69 = v64
	goto L12
L17:
	;
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[3]))
	v80 = F_get_database_name(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+124))
	F_check_can_set_role(m, v84, v50)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	F_aclcheck_error(m, v75, int32(9), v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[4])))
	if v88 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = int32(0)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v92 != int32(112) {
		v101 = v91
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v102 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if v101 != 0 {
		goto L8
	} else {
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v95 != int32(103) {
		v101 = v91
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+2)))
	v101 = base.B2i32(v98 == int32(95))
	goto L28
L31:
	;
	goto L26
L32:
	;
	m.G0 = v30 + int32(128)
	return
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v30)+124))
	if v50 != v140 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v106 = F_get_namespace_oid(m, v69, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if v106 == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+116)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = int32(2615)
	F_checkMembershipInCurrentExtension(m, v30+int32(108))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v121 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v121 == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(100794500))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v69
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_0), v30-int32(-64))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_1), int32(135), int32(_a_F_CreateSchemaCommand_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2])) = v142 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1])) = v50
	goto L46
L44:
	;
	goto L45
L45:
	;
	v150 = F_NamespaceCreate(m, v69, v50, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v155 = int32(_a_F_CreateSchemaCommand_3)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[5]))
	v159 = v157 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[5])) = v159
	goto L49
L49:
	;
	v162 = v30 + int32(92)
	F_initStringInfo(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v165 = F_quote_identifier(m, v69)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_appendStringInfoString(m, v162, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v180 = v34
	goto L53
L53:
	;
	v198 = int32(*(*int8)(unsafe.Add(mBase, uint32(v180))))
	goto L55
L54:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v208 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if base.B2i32(v198 == int32(32))|base.B2i32(base.Ui32((v198-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v180 = v180 + int32(1)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v180
	F_appendStringInfo(m, v30+int32(92), int32(_a_F_CreateSchemaCommand_4), v30+int32(48))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v30)+92))
	F_set_config_option(m, int32(_a_F_CreateSchemaCommand_5), v219, int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = int32(2615)
	v229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+116)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v229
	v234 = *(*int64)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v234
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v30)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v239
	F_EventTriggerCollectSimpleCommand(m, v30+int32(32), v30+int32(16), l1)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v248 = m.G0
	v250 = v248 + int32(-64)
	m.G0 = v250
	v252 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+24)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v250)+20)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v250)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v250)+32)) = v252
	*(*int64)(unsafe.Add(mBase, uint32(v250)+40)) = v252
	v260 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v250)+48)) = v260
	if v247 == v260 {
		v716 = v5
		v720 = v5
		v722 = v5
		v723 = v5
		v724 = v5
		v725 = v5
		v726 = v5
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v729 = F_list_concat(m, int32(0), v723)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L144
	}
L64:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v264 <= int32(0) {
		v716 = v5
		v720 = v5
		v722 = v5
		v723 = v5
		v724 = v5
		v725 = v5
		v726 = v5
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v282 = v5
	v285 = v5
	v286 = v5
	v288 = v5
	v289 = v5
	v290 = v5
	v291 = v5
	v292 = v5
	goto L66
L66:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v285<<(uint(int32(2))%32))))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	switch v299 - int32(152) {
	case 0:
		goto L71
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 30, 31, 32, 33, 34, 35, 36:
		goto L72
	case 8:
		goto L74
	case 29:
		goto L70
	case 37:
		goto L73
	default:
		goto L75
	}
L67:
	;
	v716 = v685
	v720 = v689
	v722 = v691
	v723 = v692
	v724 = v693
	v725 = v694
	v726 = v695
	goto L63
L68:
	;
	v698 = v285 + int32(1)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v698 < v699 {
		v282 = v685
		v285 = v698
		v286 = v689
		v288 = v691
		v289 = v692
		v290 = v693
		v291 = v694
		v292 = v695
		goto L66
	} else {
		goto L143
	}
L69:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
	F_checkSchemaNameRV(m, v248+int32(-48), v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L6
	} else {
		goto L141
	}
L70:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	F_checkSchemaNameRV(m, v248+int32(-48), v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L139
	}
L71:
	;
	v651 = F_lappend(m, v290, v298)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L6
	} else {
		goto L138
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L135
	}
L73:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	F_checkSchemaNameRV(m, v248+int32(-48), v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L133
	}
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	F_checkSchemaNameRV(m, v248+int32(-48), v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L80
	}
L75:
	;
	if v299 == int32(204) {
		goto L69
	} else {
		goto L76
	}
L76:
	;
	if v299 != int32(230) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	F_checkSchemaNameRV(m, v248+int32(-48), v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v311 = F_lappend(m, v286, v298)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+32)) = v311
	v685 = v282
	v689 = v311
	v691 = v288
	v692 = v289
	v693 = v290
	v694 = v291
	v695 = v292
	goto L68
L80:
	;
	v320 = F_palloc0(m, int32(56))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = int32(160)
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v298)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+48)) = v324
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v298)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+40)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v298)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+32)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v298)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+24)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v298)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v332
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v298)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v334
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
	if v338 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v610
	v626 = F_lappend(m, v288, v320)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L6
	} else {
		goto L132
	}
L83:
	;
	v610 = int32(0)
	v613 = v282
	goto L82
L84:
	;
	goto L85
L85:
	;
	v342 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v344 <= v342 {
		v610 = v342
		v613 = v282
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v359 = v342
	v362 = v282
	v364 = v342
	goto L87
L87:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374+v364<<(uint(int32(2))%32))))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if v379 != int32(90) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v610 = v579
	v613 = v582
	goto L82
L89:
	;
	v595 = v364 + int32(1)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v595 < v596 {
		v359 = v579
		v362 = v582
		v364 = v595
		goto L87
	} else {
		goto L131
	}
L90:
	;
	v558 = F_palloc0(m, int32(68))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L129
	}
L91:
	;
	v539 = int32(0)
	v545 = v362
	goto L90
L92:
	;
	v527 = F_lappend(m, v359, v378)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L128
	}
L93:
	;
	if v379 != int32(161) {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v378)+56))
	v420 = F_copyObjectImpl(m, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L6
	} else {
		goto L104
	}
L96:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v384 != int32(9) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v388 = F_palloc0(m, int32(32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = int64(68719476883)
	v394 = F_copyObjectImpl(m, v378)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+20)) = v394
	v398 = F_palloc0(m, int32(20))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398))) = int32(146)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v403 = F_copyObjectImpl(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398)+4)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v250)+60)) = v388
	v411 = F_list_make1_impl(m, int32(1), v248+int32(-60))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398)+12)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = v411
	v416 = F_lappend(m, v362, v398)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+48)) = v416
	v579 = v359
	v582 = v416
	goto L89
L104:
	;
	F_transformConstraintAttrs(m, l0, v420)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v424 = int32(0)
	if v420 == v424 {
		goto L91
	} else {
		goto L106
	}
L106:
	;
	v429 = v424
	v437 = v420
	v438 = int32(0)
	v443 = v362
	goto L107
L107:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v455 <= v429 {
		v539 = v437
		v545 = v443
		goto L90
	} else {
		goto L109
	}
L108:
	;
	v539 = v522
	v545 = v524
	goto L90
L109:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457+v429<<(uint(int32(2))%32))))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if base.Ui32(int32(6)) <= base.Ui32(v462-int32(10)) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	if v522 != 0 {
		v429 = v521 + int32(1)
		v437 = v522
		v438 = v523
		v443 = v524
		goto L107
	} else {
		goto L127
	}
L111:
	;
	v521 = v429
	v522 = v437
	v523 = int32(0)
	v524 = v443
	goto L110
L112:
	;
	v515 = int32(1)
	v518 = F_list_delete_nth_cell(m, v437, v429)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L126
	}
L113:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	v472 = F_makeString(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L6
	} else {
		goto L119
	}
L114:
	;
	if v462 == int32(9) {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if v438 == int32(0) {
		goto L111
	} else {
		goto L118
	}
L117:
	;
	goto L111
L118:
	;
	v514 = v443
	goto L112
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+12)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v250)+56)) = v472
	v479 = F_list_make1_impl(m, int32(1), v248+int32(-52))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461)+76)) = v479
	v483 = F_palloc0(m, int32(32))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+20)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v483)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v483))) = int64(68719476883)
	v491 = F_palloc0(m, int32(20))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = int32(146)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v496 = F_copyObjectImpl(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v250)+8)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v250)+52)) = v483
	v504 = F_list_make1_impl(m, int32(1), v248+int32(-56))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+12)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v491)+8)) = v504
	v509 = F_lappend(m, v443, v491)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+48)) = v509
	v514 = v509
	goto L112
L126:
	;
	v521 = v429 - v515
	v522 = v518
	v523 = v515
	v524 = v514
	goto L110
L127:
	;
	goto L108
L128:
	;
	v579 = v527
	v582 = v362
	goto L89
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = int32(90)
	base.MemoryCopy(m, v558, v378, int32(68))
	*(*int32)(unsafe.Add(mBase, uint32(v558)+56)) = v539
	v565 = F_lappend(m, v359, v558)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v579 = v565
	v582 = v545
	goto L89
L131:
	;
	goto L88
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+28)) = v626
	v685 = v613
	v689 = v286
	v691 = v626
	v692 = v289
	v693 = v290
	v694 = v291
	v695 = v292
	goto L68
L133:
	;
	v634 = F_lappend(m, v289, v298)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+24)) = v634
	v685 = v282
	v689 = v286
	v691 = v288
	v692 = v634
	v693 = v290
	v694 = v291
	v695 = v292
	goto L68
L135:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v641
	F_errmsg_internal(m, int32(_a_F_CreateSchemaCommand_6), v250)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_7), int32(_a_F_CreateSchemaCommand_8), int32(_a_F_CreateSchemaCommand_9))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+44)) = v651
	v685 = v282
	v689 = v286
	v691 = v288
	v692 = v289
	v693 = v651
	v694 = v291
	v695 = v292
	goto L68
L139:
	;
	v659 = F_lappend(m, v291, v298)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+40)) = v659
	v685 = v282
	v689 = v286
	v691 = v288
	v692 = v289
	v693 = v290
	v694 = v659
	v695 = v292
	goto L68
L141:
	;
	v667 = F_lappend(m, v292, v298)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+36)) = v667
	v685 = v282
	v689 = v286
	v691 = v288
	v692 = v289
	v693 = v290
	v694 = v291
	v695 = v667
	goto L68
L143:
	;
	goto L67
L144:
	;
	v731 = F_list_concat(m, v729, v722)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v733 = F_list_concat(m, v731, v720)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	v735 = F_list_concat(m, v733, v726)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	v737 = F_list_concat(m, v735, v725)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v739 = F_list_concat(m, v737, v724)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	v741 = F_list_concat(m, v739, v716)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	m.G0 = v250 - int32(-64)
	if v741 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	F_AtEOXact_GUC(m, int32(1), v159)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L160
	}
L152:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	if v748 <= int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v762 = int32(0)
	goto L154
L154:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v741)+12))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v778+v762<<(uint(int32(2))%32))))
	v784 = F_palloc0(m, int32(120))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L156
	}
L155:
	;
	goto L151
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v784)+116)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v784)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v784)+100)) = v782
	v789 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+30)) = uint8(v789)
	*(*int64)(unsafe.Add(mBase, uint32(v784))) = int64(25769804110)
	*(*int32)(unsafe.Add(mBase, uint32(v784)+24)) = int32(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[8]))
	F_ProcessUtility(m, v784, v795, v789, int32(3), v789, v789, v801, v789)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	v808 = v762 + int32(1)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	if v808 < v809 {
		v762 = v808
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v30)+124))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2])) = v842
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1])) = v841
	goto L161
L161:
	;
	goto L32
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v50
	F_errmsg_internal(m, int32(_a_F_CreateSchemaCommand_10), v30)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_1), int32(85), int32(_a_F_CreateSchemaCommand_2))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v69
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_11), v30+int32(80))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v905 = F_errdetail(m, int32(_a_F_CreateSchemaCommand_12), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_1), int32(110), int32(_a_F_CreateSchemaCommand_2))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSchemaPublications(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v8 = int64(0)
	v10 = F_SearchSysCacheList(m, int32(50), int32(1), base.I64_extend_i32_u(l0), v8, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ReleaseCatCacheList(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v24 = int32(0)
	v26 = int32(0)
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)+v24<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33)+4))
	v36 = F_lappend_oid(m, v26, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_ReleaseCatCacheList(m, v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v39 = v24 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	if v39 < v40 {
		v24 = v39
		v26 = v36
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	return v36
}
func F_has_schema_privilege_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14302(m, l0, int32(_a_F_has_schema_privilege_id_0), int32(2615))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_schema_privilege_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_has_schema_privilege_name[0]))
			v15 = F_text_to_cstring(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v18 = F_get_namespace_oid(m, v15, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v21 = F_convert_any_priv_string(m, v10, int32(_a_F_has_schema_privilege_name_0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = F_object_aclcheck(m, int32(2615), v18, v13, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v23 == int32(0)))
						}
					}
				}
			}
		}
	}
}
func F_has_schema_privilege_name_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14304(m, l0, int32(_a_F_has_schema_privilege_name_id_0), int32(2615))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_schema_privilege_name_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v16 = F_text_to_cstring(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					v19 = F_get_namespace_oid(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						v22 = F_convert_any_priv_string(m, v11, int32(_a_F_has_schema_privilege_name_name_0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							v24 = F_object_aclcheck(m, int32(2615), v19, v13, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v24 == int32(0)))
							}
						}
					}
				}
			}
		}
	}
}
func F_schema_to_xml_and_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v12 = F_LookupExplicitNamespace(m, v3, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_schema_to_xmlschema_internal(m, v3, v9)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_schema_to_xmlschema_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = F_makeStringInfo(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_LookupExplicitNamespace(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_schema_to_xmlschema_internal_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_appendStringInfo(m, v11, int32(_a_F_schema_to_xmlschema_internal_1), v9+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_schema_to_xmlschema_internal_2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v33 = v9 + int32(32)
	F_initStringInfo(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_appendStringInfo(m, v33, int32(_a_F_schema_to_xmlschema_internal_3), v9)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v41 = F_query_to_oid_list(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v81 = F_map_sql_typecoll_to_xmlschema_types(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	if v41 == int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v45 <= int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v50 = int32(0)
	v54 = v3
	goto L17
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v50<<(uint(int32(2))%32))))
	v61 = F_table_open(m, v59, int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v80 = v66
	goto L13
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v64 = F_CreateTupleDescCopy(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v66 = F_lappend(m, v54, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_relation_close(m, v61, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v72 = v50 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v72 < v73 {
		v50 = v72
		v54 = v66
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	F_appendStringInfoString(m, v11, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v85 = m.G0
	v87 = v85 - int32(16)
	m.G0 = v87
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_schema_to_xmlschema_internal[0]))
	v91 = F_get_database_name(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v93 = F_get_namespace_name(m, v14)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_initStringInfo(m, v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v97 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
