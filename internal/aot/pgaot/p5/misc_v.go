package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValuesNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 == v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	m.T0[v30].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v24 + v13
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v13 < v19 {
		v24 = v12
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v13 < int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L1
L7:
	;
	v24 = int32(-1)
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v35 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return v14
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v38 <= v35 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v41 = v35 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41)))
	F_ReScanExprContext(m, v28)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v50 = int32(0)
	v51 = int32(_a_F_ValuesNext_0)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ValuesNext[0]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ValuesNext[0])) = v54
	if v44 == v50 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ValuesNext[0])) = v52
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v130 = v128 & int32(_a_F_ValuesNext_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)) = uint16(v133)
	goto L32
L15:
	;
	v59 = F_ExecInitExprList(m, v47, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v63 = v44
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v64 <= int32(0) {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	if v59 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v63 = v59
	goto L17
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v69 = v50
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v82 = v69 << (uint(int32(2)) % 32)
	v83 = v68 + v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v82)))
	v87 = v69 + v67
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v89 = m.T0[v88].(func(*base.Module, int32, int32, int32) int32)(m, v86, v28, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L14
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v89
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v92 != 0 {
		v109 = v89
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v109
	v112 = v69 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v112 < v113 {
		v69 = v112
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v69<<(uint(int32(4))%32))+24)))
	if v96 != int32(_a_F_ValuesNext_2) {
		v109 = v89
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v99 != int32(1) {
		v108 = v89
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = v108
	goto L24
L28:
	;
	goto L27
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v102 != int32(3) {
		v108 = v89
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+2))
	v108 = v105 + int32(18)
	goto L28
L31:
	;
	goto L22
L32:
	;
	goto L10
}
func F_VirtualXactLockTableInsert(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLockTableInsert[0]))
	v8 = F_LWLockAcquire(m, v4+int32(584), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLockTableInsert[0]))
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+608)) = uint8(v12)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+612)) = v14
		F_LWLockRelease(m, v11+int32(584))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_varchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = v17 - int32(4)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v22 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.B2i32(v19 < int32(0))|base.B2i32(v51 <= v19) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v28 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v39 = int32(1)
	if v22&v39 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v31 = int32(16)
	goto L9
L8:
	;
	v31 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(4)
	goto L12
L11:
	;
	v38 = v31
	goto L12
L12:
	;
	v51 = v38
	goto L3
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v57 = int32(1)
	if v22&v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v107 = v13
	goto L16
L16:
	;
	m.G0 = v10 + int32(16)
	return v107
L17:
	;
	v104 = F_cstring_to_text_with_len(m, v62, v63)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L18:
	;
	v61 = v57
	goto L20
L19:
	;
	v61 = int32(4)
	goto L20
L20:
	;
	v62 = v13 + v61
	v63 = F_pg_mbcharcliplen(m, v62, v51, v19)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v56|base.B2i32(v51 <= v63) != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v68 = v63
	goto L23
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v62))))
	if v75 == int32(32) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L29
	}
L25:
	;
	v79 = v68 + int32(1)
	if v51 != v79 {
		v68 = v79
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L17
L29:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	F_errmsg(m, int32(_a_F_varchar_0), v10)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_varchar_1), int32(640), int32(_a_F_varchar_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v107 = v104
	goto L16
}
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F___vfprintf_internal(m, l0, l1, l2, int32(_a_F_vfprintf_0), int32(_a_F_vfprintf_1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_vfscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v104 int32
	_ = v104
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v189 int32
	_ = v189
	var v195 int64
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
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
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v502 int64
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int64
	_ = v574
	var v585 int32
	_ = v585
	var v600 int64
	_ = v600
	var v602 int64
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int64
	_ = v631
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v668 int64
	_ = v668
	var v669 int64
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v684 int32
	_ = v684
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int64
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int64
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v960 int32
	_ = v960
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1072 int32
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int64
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1200 int64
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int64
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1233 int64
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1320 int32
	_ = v1320
	var v1332 int64
	_ = v1332
	var v1361 int64
	_ = v1361
	var v1363 int64
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1380 int64
	_ = v1380
	var v1389 int64
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1417 int32
	_ = v1417
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int64
	_ = v1464
	var v1466 int64
	_ = v1466
	var v1467 int64
	_ = v1467
	var v1478 int32
	_ = v1478
	var v1490 int64
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int64
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1560 int64
	_ = v1560
	var v1561 int64
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1575 int32
	_ = v1575
	var v1583 int64
	_ = v1583
	var v1587 int64
	_ = v1587
	var v1590 int64
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int64
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1609 int64
	_ = v1609
	var v1615 int64
	_ = v1615
	var v1616 int64
	_ = v1616
	var v1618 int64
	_ = v1618
	var v1621 int64
	_ = v1621
	var v1622 int64
	_ = v1622
	var v1624 int64
	_ = v1624
	var v1625 int64
	_ = v1625
	var v1629 int64
	_ = v1629
	var v1636 int64
	_ = v1636
	var v1647 int64
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1671 int64
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1696 int64
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1766 int32
	_ = v1766
	var v1769 int64
	_ = v1769
	var v1773 int64
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1782 int64
	_ = v1782
	var v1805 int64
	_ = v1805
	var v1813 int64
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1821 int32
	_ = v1821
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1849 int64
	_ = v1849
	var v1853 int64
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1861 int32
	_ = v1861
	var v1865 int64
	_ = v1865
	var v1866 int64
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1921 int64
	_ = v1921
	var v1929 int64
	_ = v1929
	var v1930 int64
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1936 int64
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1940 int64
	_ = v1940
	var v1942 int64
	_ = v1942
	var v1946 int64
	_ = v1946
	var v1947 int64
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2002 int64
	_ = v2002
	var v2006 int64
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2016 int64
	_ = v2016
	var v2021 int64
	_ = v2021
	var v2031 int64
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2035 int64
	_ = v2035
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int64
	_ = v2066
	var v2070 int64
	_ = v2070
	var v2078 int64
	_ = v2078
	var v2079 int64
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2085 int64
	_ = v2085
	var v2088 int64
	_ = v2088
	var v2091 int64
	_ = v2091
	var v2095 int64
	_ = v2095
	var v2105 int64
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int64
	_ = v2110
	var v2112 int64
	_ = v2112
	var v2119 int64
	_ = v2119
	var v2137 int32
	_ = v2137
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2201 int32
	_ = v2201
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2316 int32
	_ = v2316
	var v2323 int32
	_ = v2323
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2394 int32
	_ = v2394
	var v2404 int32
	_ = v2404
	var v2447 int32
	_ = v2447
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2494 int32
	_ = v2494
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2537 int32
	_ = v2537
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2590 int32
	_ = v2590
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2678 int32
	_ = v2678
	var v2684 int32
	_ = v2684
	var v2697 int32
	_ = v2697
	var v2698 int64
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2705 int64
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2709 int64
	_ = v2709
	var v2729 int32
	_ = v2729
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2743 int32
	_ = v2743
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2760 int64
	_ = v2760
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2778 int32
	_ = v2778
	var v2785 int32
	_ = v2785
	var v2789 int64
	_ = v2789
	var v2793 int32
	_ = v2793
	var v2797 int32
	_ = v2797
	var v2805 int32
	_ = v2805
	var v2811 int32
	_ = v2811
	var v2815 int32
	_ = v2815
	var v2825 int32
	_ = v2825
	var v2832 int32
	_ = v2832
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2862 int32
	_ = v2862
	var v2888 int32
	_ = v2888
	var v2900 int32
	_ = v2900
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(304)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v2888 + int32(304)
	return v2900
L2:
	;
	v2888 = v2862
	v2900 = int32(-1)
	goto L1
L3:
	;
	v33 = F___toread(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v40 == int32(0) {
		v2888 = v28
		v2900 = v4
		goto L1
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v37 == int32(0) {
		v2862 = v28
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v43 = l0
	v44 = l1
	v45 = l2
	v47 = v40
	v49 = v4
	v50 = v28
	v55 = v4
	v62 = v4
	v66 = int64(0)
	goto L12
L10:
	;
	if v2842 == int32(0) {
		v2888 = v50
		v2900 = v2845
		goto L1
	} else {
		goto L610
	}
L11:
	;
	if v62 != 0 {
		goto L607
	} else {
		goto L608
	}
L12:
	;
	v69 = v47 & int32(255)
	goto L16
L13:
	;
	v2797 = int32(0)
	v2805 = v2797
	v2811 = v2797
	v2815 = int32(1)
	goto L11
L14:
	;
	goto L13
L15:
	;
	v2793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2770)+1)))
	if v2793 != 0 {
		v43 = v2766
		v44 = v2770 + int32(1)
		v45 = v2768
		v47 = v2793
		v49 = v2772
		v50 = v2773
		v55 = v2778
		v62 = v2785
		v66 = v2789
		goto L12
	} else {
		goto L606
	}
L16:
	;
	if base.B2i32(v69 == int32(32))|base.B2i32(base.Ui32(v69-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = v44
	goto L20
L18:
	;
	goto L19
L19:
	;
	if v69 == int32(37) {
		goto L43
	} else {
		goto L44
	}
L20:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	goto L22
L21:
	;
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v115 - v116)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v122-v116) <= v112) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if base.B2i32(v104 == int32(32))|base.B2i32(base.Ui32(v104-int32(9)) < base.Ui32(int32(5))) != 0 {
		v78 = v78 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L28
L25:
	;
	v129 = v122
	goto L27
L26:
	;
	v129 = v116 + base.I32_wrap_i64(v112)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v129
	goto L24
L28:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v156 != v157 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v174 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	goto L35
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v156 + int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v165 = v162
	goto L30
L32:
	;
	goto L33
L33:
	;
	v163 = F___shgetc(m, v43)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v165 = v163
	goto L30
L35:
	;
	if base.B2i32(v165 == int32(32))|base.B2i32(base.Ui32(v165-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v178 = v173 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v178
	v180 = v178
	goto L39
L38:
	;
	v180 = v173
	goto L39
L39:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v2766 = v43
	v2768 = v45
	v2770 = v78
	v2772 = v49
	v2773 = v50
	v2778 = v55
	v2785 = v62
	v2789 = base.I64_extend_i32_s(v180-v181) + (v184 + v66)
	goto L15
L40:
	;
	v352 = int32(0)
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if base.Ui32((v354-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L82
	} else {
		goto L83
	}
L41:
	;
	v318 = v189 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v318) {
		goto L76
	} else {
		goto L77
	}
L42:
	;
	v348 = v45
	v350 = v44 + int32(2)
	v351 = int32(0)
	goto L40
L43:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v189 == int32(42) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v195 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v198 - v199)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v205-v199) <= v195) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if v189 != int32(37) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v214 == int32(37) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v212 = v205
	goto L51
L50:
	;
	v212 = v199 + base.I32_wrap_i64(v195)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v212
	goto L48
L52:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v295 != v274 {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	goto L56
L54:
	;
	goto L55
L55:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v261 != v262 {
		goto L65
	} else {
		goto L66
	}
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v242 != v243 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v271 = v44 + int32(1)
	v274 = v251
	goto L52
L58:
	;
	goto L63
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v242 + int32(1)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v251 = v248
	goto L58
L60:
	;
	goto L61
L61:
	;
	v249 = F___shgetc(m, v43)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v251 = v249
	goto L58
L63:
	;
	if base.B2i32(v251 == int32(32))|base.B2i32(base.Ui32(v251-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v261 + int32(1)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v271 = v44
	v274 = v267
	goto L52
L66:
	;
	goto L67
L67:
	;
	v268 = F___shgetc(m, v43)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v271 = v44
	v274 = v268
	goto L52
L69:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v297 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v2766 = v43
	v2768 = v45
	v2770 = v271
	v2772 = v49
	v2773 = v50
	v2778 = v55
	v2785 = v62
	v2789 = base.I64_extend_i32_s(v307-v308) + (v311 + v66)
	goto L15
L72:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v300 - int32(1)
	goto L74
L73:
	;
	goto L74
L74:
	;
	if base.B2i32(int32(0) <= v274)|v62 != 0 {
		v2888 = v50
		v2900 = v62
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v2862 = v50
	goto L2
L76:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v348 = v45 + int32(4)
	v350 = v44 + int32(1)
	v351 = v344
	goto L40
L77:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+2)))
	if v321 != int32(36) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v326 = m.G0
	v328 = v326 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v328)+12)) = v45
	if base.Ui32(int32(1)) < base.Ui32(v318) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v337 = v45 + v318<<(uint(int32(2))%32) - int32(4)
	goto L81
L80:
	;
	v337 = v45
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+8)) = v337 + int32(4)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v348 = v45
	v350 = v44 + int32(3)
	v351 = v341
	goto L40
L82:
	;
	v362 = v354
	v365 = v350
	v369 = v352
	goto L85
L83:
	;
	v403 = v354
	v406 = v350
	v410 = v352
	goto L84
L84:
	;
	if v403&int32(255) != int32(109) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v386 = int32(10)
	v388 = int32(255)
	v391 = int32(48)
	v392 = v369*v386 + v362&v388 - v391
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+1)))
	v395 = v365 + int32(1)
	if base.Ui32((v393-v391)&v388) < base.Ui32(v386) {
		v362 = v393
		v365 = v395
		v369 = v392
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v403 = v393
	v406 = v395
	v410 = v392
	goto L84
L87:
	;
	goto L86
L88:
	;
	v438 = v403
	v439 = v49
	v440 = v55
	v441 = v352
	v442 = v406
	goto L90
L89:
	;
	v431 = int32(0)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	v438 = v434
	v439 = v431
	v440 = v431
	v441 = base.B2i32(v351 != v431)
	v442 = v406 + int32(1)
	goto L90
L90:
	;
	v444 = v442 + int32(1)
	switch v438&int32(255) - int32(65) {
	case 0, 2, 4, 5, 6, 18, 23, 26, 32, 34, 35, 36, 37, 38, 40, 45, 46, 47, 50, 52, 55:
		goto L92
	default:
		v2805 = v439
		v2811 = v440
		v2815 = v441
		goto L11
	case 11:
		goto L93
	case 39:
		goto L96
	case 41:
		v472 = int32(3)
		v473 = v444
		goto L91
	case 43:
		goto L95
	case 51, 57:
		goto L94
	}
L91:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v479 = base.B2i32(v475&int32(47) == int32(3))
	if v475&int32(47) == int32(3) {
		goto L109
	} else {
		goto L110
	}
L92:
	;
	v472 = int32(0)
	v473 = v442
	goto L91
L93:
	;
	v472 = int32(2)
	v473 = v444
	goto L91
L94:
	;
	v472 = int32(1)
	v473 = v444
	goto L91
L95:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	v463 = base.B2i32(v461 == int32(108))
	if v461 == int32(108) {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	v454 = base.B2i32(v452 == int32(104))
	if v452 == int32(104) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v455 = v442 + int32(2)
	goto L99
L98:
	;
	v455 = v444
	goto L99
L99:
	;
	if v452 == int32(104) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v458 = int32(-2)
	goto L102
L101:
	;
	v458 = int32(-1)
	goto L102
L102:
	;
	v472 = v458
	v473 = v455
	goto L91
L103:
	;
	v464 = v442 + int32(2)
	goto L105
L104:
	;
	v464 = v444
	goto L105
L105:
	;
	if v461 == int32(108) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v467 = int32(3)
	goto L108
L107:
	;
	v467 = int32(1)
	goto L108
L108:
	;
	v472 = v467
	v473 = v464
	goto L91
L109:
	;
	v480 = int32(1)
	goto L111
L110:
	;
	v480 = v472
	goto L111
L111:
	;
	if v475&int32(47) == int32(3) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v602 = base.I64_extend_i32_s(v585)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v602
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v605 - v606)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if base.B2i32(v602 == int64(0))|base.B2i32(base.I64_extend_i32_s(v612-v606) <= v602) != 0 {
		goto L149
	} else {
		goto L150
	}
L113:
	;
	v483 = v475 | int32(32)
	goto L115
L114:
	;
	v483 = v475
	goto L115
L115:
	;
	if v483 == int32(91) {
		v585 = v410
		v600 = v66
		goto L112
	} else {
		goto L116
	}
L116:
	;
	if v483 != int32(110) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v502 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v502
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v505 - v506)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v512-v506) <= v502) != 0 {
		goto L133
	} else {
		goto L134
	}
L118:
	;
	if v483 != int32(99) {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v351 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v490 = int32(1)
	if v410 <= v490 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v493 = v490
	goto L124
L123:
	;
	v493 = v410
	goto L124
L124:
	;
	v585 = v493
	v600 = v66
	goto L112
L125:
	;
	v2766 = v43
	v2768 = v348
	v2770 = v473
	v2772 = v439
	v2773 = v50
	v2778 = v440
	v2785 = v62
	v2789 = v66
	goto L15
L126:
	;
	goto L125
L127:
	;
	switch v480 + int32(2) {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2, 3:
		goto L129
	default:
		goto L126
	case 5:
		goto L128
	}
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v66
	goto L126
L129:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v351))) = uint32(v66)
	goto L125
L130:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v66)
	goto L125
L131:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v66)
	goto L125
L132:
	;
	goto L136
L133:
	;
	v519 = v512
	goto L135
L134:
	;
	v519 = v506 + base.I32_wrap_i64(v502)
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v519
	goto L132
L136:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v546 != v547 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v564 {
		goto L145
	} else {
		goto L146
	}
L138:
	;
	goto L143
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v546 + int32(1)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	v555 = v552
	goto L138
L140:
	;
	goto L141
L141:
	;
	v553 = F___shgetc(m, v43)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	v555 = v553
	goto L138
L143:
	;
	if base.B2i32(v555 == int32(32))|base.B2i32(base.Ui32(v555-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L136
	} else {
		goto L144
	}
L144:
	;
	goto L137
L145:
	;
	v568 = v563 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v568
	v570 = v568
	goto L147
L146:
	;
	v570 = v563
	goto L147
L147:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v585 = v410
	v600 = base.I64_extend_i32_s(v570-v571) + (v574 + v66)
	goto L112
L148:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v621 != v622 {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v619 = v612
	goto L151
L150:
	;
	v619 = v606 + base.I32_wrap_i64(v602)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v619
	goto L148
L152:
	;
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v631 {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v621 + int32(1)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v627 = F___shgetc(m, v43)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	if v627 < int32(0) {
		v2805 = v439
		v2811 = v440
		v2815 = v441
		goto L11
	} else {
		goto L157
	}
L157:
	;
	goto L152
L158:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v634 - int32(1)
	goto L160
L159:
	;
	goto L160
L160:
	;
	switch v483 - int32(88) {
	case 0, 24, 32:
		v1072 = int32(16)
		goto L166
	case 1, 2, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 22, 25, 26, 28, 30, 31:
		v2735 = v473
		v2737 = v439
		v2743 = v440
		goto L161
	case 3, 11, 27:
		goto L170
	case 9, 13, 14, 15:
		goto L171
	case 12, 29:
		goto L168
	case 17:
		goto L167
	case 23:
		goto L169
	default:
		goto L172
	}
L161:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v2760 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v2766 = v43
	v2768 = v348
	v2770 = v2735
	v2772 = v2737
	v2773 = v50
	v2778 = v2743
	v2785 = v62 + base.B2i32(v351 != int32(0))
	v2789 = base.I64_extend_i32_s(v2756-v2757) + (v2760 + v600)
	goto L15
L162:
	;
	v2162 = base.B2i32(v483 != int32(99))
	if v483 != int32(99) {
		goto L480
	} else {
		goto L481
	}
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v669
	*(*int64)(unsafe.Add(mBase, uint32(v351)+8)) = v668
	v2735 = v473
	v2737 = v439
	v2743 = v440
	goto L161
L164:
	;
	v1997 = m.G0
	v1999 = v1997 - int32(32)
	m.G0 = v1999
	v2002 = v668 & int64(281474976710655)
	v2006 = int64(base.Ui64(v668)>>(uint(int64(48))%64)) & int64(32767)
	v2007 = base.I32_wrap_i64(v2006)
	if base.Ui32(v2007-int32(_a_F_vfscanf_0)) <= base.Ui32(int32(2045)) {
		goto L443
	} else {
		goto L444
	}
L165:
	;
	v1844 = m.G0
	v1846 = v1844 - int32(32)
	m.G0 = v1846
	v1849 = v668 & int64(281474976710655)
	v1853 = int64(base.Ui64(v668)>>(uint(int64(48))%64)) & int64(32767)
	v1854 = base.I32_wrap_i64(v1853)
	if base.Ui32(v1854-int32(_a_F_vfscanf_1)) <= base.Ui32(int32(253)) {
		goto L395
	} else {
		goto L396
	}
L166:
	;
	v1073 = int64(0)
	v1074 = int32(0)
	v1078 = m.G0
	v1080 = v1078 - int32(16)
	m.G0 = v1080
	if base.B2i32(v1072 != int32(1))&base.B2i32(base.Ui32(v1072) <= base.Ui32(int32(36))) == v1074 {
		goto L235
	} else {
		goto L236
	}
L167:
	;
	v1072 = int32(0)
	goto L166
L168:
	;
	v1072 = int32(10)
	goto L166
L169:
	;
	v1072 = int32(8)
	goto L166
L170:
	;
	if v483|int32(16) == int32(115) {
		goto L177
	} else {
		goto L178
	}
L171:
	;
	F___floatscan(m, v50+int32(8), v43, v480, int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L6
	} else {
		goto L174
	}
L172:
	;
	v642 = v483 - int32(65)
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v642))|base.B2i32(int32(1)<<(uint(v642)%32)&int32(113) == int32(0)) != 0 {
		v2735 = v473
		v2737 = v439
		v2743 = v440
		goto L161
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	if v658 == int64(0)-base.I64_extend_i32_s(v660-v661) {
		v2832 = v439
		v2838 = v440
		v2842 = v441
		v2845 = v62
		goto L10
	} else {
		goto L175
	}
L175:
	;
	if v351 == int32(0) {
		v2735 = v473
		v2737 = v439
		v2743 = v440
		goto L161
	} else {
		goto L176
	}
L176:
	;
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v50)+16))
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	switch v480 {
	case 0:
		goto L165
	case 1:
		goto L164
	case 2:
		goto L163
	default:
		v2735 = v473
		v2737 = v439
		v2743 = v440
		goto L161
	}
L177:
	;
	v675 = v50 + int32(32)
	v676 = int32(-1)
	goto L182
L178:
	;
	goto L179
L179:
	;
	v795 = v50 + int32(32)
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+1)))
	v798 = base.B2i32(v796 == int32(94))
	goto L195
L180:
	;
	v784 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+32)) = uint8(v784)
	if v483 != int32(115) {
		v2137 = v473
		goto L162
	} else {
		goto L192
	}
L181:
	;
	goto L180
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v675))) = uint8(v676)
	v684 = v50 + int32(289)
	*(*uint8)(unsafe.Add(mBase, uint32(v684-int32(1)))) = uint8(v676)
	goto L183
L183:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v675)+2)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v675)+1)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v684-int32(3)))) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v684-int32(2)))) = uint8(v676)
	goto L184
L184:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v675)+3)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v684-int32(4)))) = uint8(v676)
	goto L185
L185:
	;
	v709 = (int32(0) - v675) & int32(3)
	v710 = v675 + v709
	v714 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v710))) = v714
	v718 = (int32(257) - v709) & int32(-4)
	v719 = v710 + v718
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(4)))) = v714
	if base.Ui32(v718) < base.Ui32(int32(9)) {
		goto L181
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+8)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(8)))) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(12)))) = v714
	if base.Ui32(v718) < base.Ui32(int32(25)) {
		goto L181
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+24)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v710)+20)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v710)+16)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v710)+12)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(16)))) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(20)))) = v714
	v745 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v719-v745))) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v719-int32(28)))) = v714
	v754 = v710&int32(4) | v745
	v755 = v718 - v754
	if base.Ui32(v755) < base.Ui32(int32(32)) {
		goto L181
	} else {
		goto L188
	}
L188:
	;
	v760 = base.I64_extend_i32_u(v714) * int64(4294967297)
	v763 = v754 + v710
	v764 = v755
	goto L189
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v763)+24)) = v760
	*(*int64)(unsafe.Add(mBase, uint32(v763)+16)) = v760
	*(*int64)(unsafe.Add(mBase, uint32(v763)+8)) = v760
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v760
	v772 = int32(32)
	v775 = v764 - v772
	if base.Ui32(int32(31)) < base.Ui32(v775) {
		v763 = v763 + v772
		v764 = v775
		goto L189
	} else {
		goto L191
	}
L190:
	;
	goto L181
L191:
	;
	goto L190
L192:
	;
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+65)) = uint8(v788)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+46)) = uint8(v788)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+42)) = v788
	v2137 = v473
	goto L162
L193:
	;
	v906 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+32)) = uint8(v906)
	if v796 == int32(94) {
		goto L205
	} else {
		goto L206
	}
L194:
	;
	goto L193
L195:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v795))) = uint8(v798)
	v806 = v50 + int32(289)
	*(*uint8)(unsafe.Add(mBase, uint32(v806-int32(1)))) = uint8(v798)
	goto L196
L196:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+2)) = uint8(v798)
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+1)) = uint8(v798)
	*(*uint8)(unsafe.Add(mBase, uint32(v806-int32(3)))) = uint8(v798)
	*(*uint8)(unsafe.Add(mBase, uint32(v806-int32(2)))) = uint8(v798)
	goto L197
L197:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+3)) = uint8(v798)
	*(*uint8)(unsafe.Add(mBase, uint32(v806-int32(4)))) = uint8(v798)
	goto L198
L198:
	;
	v831 = (int32(0) - v795) & int32(3)
	v832 = v795 + v831
	v836 = v798 & int32(255) * int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = v836
	v840 = (int32(257) - v831) & int32(-4)
	v841 = v832 + v840
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(4)))) = v836
	if base.Ui32(v840) < base.Ui32(int32(9)) {
		goto L194
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(8)))) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(12)))) = v836
	if base.Ui32(v840) < base.Ui32(int32(25)) {
		goto L194
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+24)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v832)+20)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v832)+16)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v832)+12)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(16)))) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(20)))) = v836
	v867 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v841-v867))) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v841-int32(28)))) = v836
	v876 = v832&int32(4) | v867
	v877 = v840 - v876
	if base.Ui32(v877) < base.Ui32(int32(32)) {
		goto L194
	} else {
		goto L201
	}
L201:
	;
	v882 = base.I64_extend_i32_u(v836) * int64(4294967297)
	v885 = v876 + v832
	v886 = v877
	goto L202
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v885)+24)) = v882
	*(*int64)(unsafe.Add(mBase, uint32(v885)+16)) = v882
	*(*int64)(unsafe.Add(mBase, uint32(v885)+8)) = v882
	*(*int64)(unsafe.Add(mBase, uint32(v885))) = v882
	v894 = int32(32)
	v897 = v886 - v894
	if base.Ui32(int32(31)) < base.Ui32(v897) {
		v885 = v885 + v894
		v886 = v897
		goto L202
	} else {
		goto L204
	}
L203:
	;
	goto L194
L204:
	;
	goto L203
L205:
	;
	v912 = v473 + int32(2)
	goto L207
L206:
	;
	v912 = v473 + int32(1)
	goto L207
L207:
	;
	if v796 == int32(94) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v939 = v934
	goto L218
L209:
	;
	v933 = v930
	v934 = v912 + int32(1)
	goto L208
L210:
	;
	v928 = base.B2i32(v796 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+126)) = uint8(v928)
	v930 = v928
	goto L209
L211:
	;
	v915 = int32(2)
	goto L213
L212:
	;
	v915 = int32(1)
	goto L213
L213:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473+v915))))
	if v917 != int32(45) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	if v917 == int32(93) {
		goto L210
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v925 = base.B2i32(v796 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+78)) = uint8(v925)
	v930 = v925
	goto L209
L217:
	;
	v933 = base.B2i32(v796 != int32(94))
	v934 = v912
	goto L208
L218:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939))))
	if v960 != int32(45) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(32)+v1041)+1)) = uint8(v933)
	v939 = v1042 + int32(1)
	goto L218
L221:
	;
	if v960 == int32(0) {
		v2805 = v439
		v2811 = v440
		v2815 = v441
		goto L11
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+1)))
	if base.B2i32(v968 == int32(0))|base.B2i32(v968 == int32(93)) != 0 {
		v1041 = int32(45)
		v1042 = v939
		goto L220
	} else {
		goto L226
	}
L224:
	;
	if v960 == int32(93) {
		v2137 = v939
		goto L162
	} else {
		goto L225
	}
L225:
	;
	v1041 = v960
	v1042 = v939
	goto L220
L226:
	;
	v974 = int32(1)
	v975 = v939 + v974
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939-v974))))
	if base.Ui32(v968) <= base.Ui32(v978) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1041 = v1016
	v1042 = v975
	goto L220
L228:
	;
	v1016 = v968
	goto L227
L229:
	;
	goto L230
L230:
	;
	v981 = v978
	goto L231
L231:
	;
	v1006 = v981 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006+(v50+int32(32))))) = uint8(v933)
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975))))
	if base.Ui32(v1006) < base.Ui32(v1011) {
		v981 = v1006
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v1016 = v1011
	goto L227
L233:
	;
	goto L232
L234:
	;
	m.G0 = v1080 + int32(16)
	v1813 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	if v1813 == int64(0)-base.I64_extend_i32_s(v1815-v1816) {
		v2832 = v439
		v2838 = v440
		v2842 = v441
		v2845 = v62
		goto L10
	} else {
		goto L382
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vfscanf[0])) = int32(28)
	v1805 = v1073
	goto L234
L236:
	;
	goto L237
L237:
	;
	goto L238
L238:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1117 != v1118 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	switch v1126 - int32(43) {
	case 0, 2:
		goto L248
	default:
		v1151 = v1126
		v1152 = v1074
		goto L247
	}
L240:
	;
	goto L245
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1117 + int32(1)
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117))))
	v1126 = v1123
	goto L240
L242:
	;
	goto L243
L243:
	;
	v1124 = F___shgetc(m, v43)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L6
	} else {
		goto L244
	}
L244:
	;
	v1126 = v1124
	goto L240
L245:
	;
	if base.B2i32(v1126 == int32(32))|base.B2i32(base.Ui32(v1126-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L238
	} else {
		goto L246
	}
L246:
	;
	goto L239
L247:
	;
	v1153 = int32(0)
	if base.B2i32(v1072 != v1153)&base.B2i32(v1072 != int32(16))|base.B2i32(v1151 != int32(48)) == v1153 {
		goto L260
	} else {
		goto L261
	}
L248:
	;
	if v1126 == int32(45) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1140 = int32(-1)
	goto L251
L250:
	;
	v1140 = int32(0)
	goto L251
L251:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1141 != v1142 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1141 + int32(1)
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1141))))
	v1151 = v1147
	v1152 = v1140
	goto L247
L253:
	;
	goto L254
L254:
	;
	v1148 = F___shgetc(m, v43)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L6
	} else {
		goto L255
	}
L255:
	;
	v1151 = v1148
	v1152 = v1140
	goto L247
L256:
	;
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v1773 {
		goto L377
	} else {
		goto L378
	}
L257:
	;
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1676) <= base.Ui32(v1702) {
		v1766 = v1152
		v1769 = v1696
		goto L256
	} else {
		goto L368
	}
L258:
	;
	if v1397&(v1397-int32(1)) != 0 {
		goto L324
	} else {
		goto L325
	}
L259:
	;
	if v1255 != int32(10) {
		v1397 = v1255
		v1399 = v1257
		goto L258
	} else {
		goto L296
	}
L260:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1163 != v1164 {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	goto L262
L262:
	;
	if v1072 != 0 {
		goto L285
	} else {
		goto L286
	}
L263:
	;
	if v1172&int32(-33) == int32(88) {
		goto L268
	} else {
		goto L269
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1163 + int32(1)
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	v1172 = v1169
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1170 = F___shgetc(m, v43)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	v1172 = v1170
	goto L263
L268:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1177 != v1178 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	goto L270
L270:
	;
	if v1072 != 0 {
		v1255 = v1072
		v1257 = v1172
		goto L259
	} else {
		goto L284
	}
L271:
	;
	v1187 = int32(16)
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1190) < base.Ui32(v1187) {
		v1397 = v1187
		v1399 = v1186
		goto L258
	} else {
		goto L276
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1177 + int32(1)
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1177))))
	v1186 = v1183
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1184 = F___shgetc(m, v43)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L6
	} else {
		goto L275
	}
L275:
	;
	v1186 = v1184
	goto L271
L276:
	;
	v1193 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v1193 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1196 - int32(1)
	goto L279
L278:
	;
	goto L279
L279:
	;
	v1200 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v1200
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v1203 - v1204)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v1210-v1204) <= v1200) != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1805 = v1073
	goto L234
L281:
	;
	v1217 = v1210
	goto L283
L282:
	;
	v1217 = v1204 + base.I32_wrap_i64(v1200)
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v1217
	goto L280
L284:
	;
	v1397 = int32(8)
	v1399 = v1172
	goto L258
L285:
	;
	v1221 = v1072
	goto L287
L286:
	;
	v1221 = int32(10)
	goto L287
L287:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1224) < base.Ui32(v1221) {
		v1255 = v1221
		v1257 = v1151
		goto L259
	} else {
		goto L288
	}
L288:
	;
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v1226 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1229 - int32(1)
	goto L291
L290:
	;
	goto L291
L291:
	;
	v1233 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v1233
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+120)) = base.I64_extend_i32_s(v1236 - v1237)
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v1243-v1237) <= v1233) != 0 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vfscanf[0])) = int32(28)
	v1805 = v1073
	goto L234
L293:
	;
	v1250 = v1243
	goto L295
L294:
	;
	v1250 = v1237 + base.I32_wrap_i64(v1233)
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v1250
	goto L292
L296:
	;
	v1261 = v1257 - int32(48)
	if base.Ui32(v1261) <= base.Ui32(int32(9)) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1270 = int32(0)
	v1274 = v1261
	goto L300
L298:
	;
	v1320 = v1261
	v1332 = v1073
	goto L299
L299:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1320) {
		v1766 = v1152
		v1769 = v1332
		goto L256
	} else {
		goto L308
	}
L300:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1290 != v1291 {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	v1320 = v1306
	v1332 = base.I64_extend_i32_u(v1302)
	goto L299
L302:
	;
	v1302 = v1270*int32(10) + v1274
	v1306 = v1299 - int32(48)
	if base.B2i32(base.Ui32(v1302) < base.Ui32(int32(429496729)))&base.B2i32(base.Ui32(v1306) <= base.Ui32(int32(9))) != 0 {
		v1270 = v1302
		v1274 = v1306
		goto L300
	} else {
		goto L307
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1290 + int32(1)
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290))))
	v1299 = v1296
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1297 = F___shgetc(m, v43)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L6
	} else {
		goto L306
	}
L306:
	;
	v1299 = v1297
	goto L302
L307:
	;
	goto L301
L308:
	;
	v1361 = v1332 * int64(10)
	v1363 = base.I64_extend_i32_u(v1320)
	goto L309
L309:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1366 != v1367 {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	v1676 = int32(10)
	v1680 = v1375
	v1696 = v1380
	goto L257
L311:
	;
	goto L310
L312:
	;
	v1377 = v1375 - int32(48)
	v1380 = v1361 + v1363
	if base.B2i32(base.Ui32(v1377) <= base.Ui32(int32(9)))&base.B2i32(base.Ui64(v1380) < base.Ui64(int64(1844674407370955162))) == int32(0) {
		goto L317
	} else {
		goto L318
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1366 + int32(1)
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366))))
	v1375 = v1372
	goto L312
L314:
	;
	goto L315
L315:
	;
	v1373 = F___shgetc(m, v43)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L6
	} else {
		goto L316
	}
L316:
	;
	v1375 = v1373
	goto L312
L317:
	;
	if base.Ui32(v1377) <= base.Ui32(int32(9)) {
		goto L311
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	v1389 = v1380 * int64(10)
	v1390 = base.I64_extend_i32_u(v1377)
	if base.Ui64(v1389) <= base.Ui64(v1390^int64(-1)) {
		v1361 = v1389
		v1363 = v1390
		goto L309
	} else {
		goto L321
	}
L320:
	;
	v1766 = v1152
	v1769 = v1380
	goto L256
L321:
	;
	goto L311
L322:
	;
	v1676 = v1397
	v1680 = v1655
	v1696 = v1671
	goto L257
L323:
	;
	v1523 = v1074
	v1527 = v1405
	goto L348
L324:
	;
	v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1405) < base.Ui32(v1397) {
		goto L323
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1397) <= base.Ui32(v1409) {
		v1655 = v1399
		v1671 = v1073
		goto L322
	} else {
		goto L328
	}
L327:
	;
	v1655 = v1399
	v1671 = v1073
	goto L322
L328:
	;
	v1417 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1397*int32(23))>>(uint(int32(5))%32))&int32(7))+uint32(_c_F_vfscanf[2]))))
	v1427 = v1409
	v1431 = v1074
	goto L329
L329:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1443 != v1444 {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	v1464 = base.I64_extend_i32_u(v1454)
	if base.Ui32(v1397) <= base.Ui32(v1457) {
		v1676 = v1397
		v1680 = v1452
		v1696 = v1464
		goto L257
	} else {
		goto L337
	}
L331:
	;
	v1453 = v1431 << (uint(v1417) % 32)
	v1454 = v1427 | v1453
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1452)+uint32(_c_F_vfscanf[1]))))
	v1458 = base.B2i32(base.Ui32(v1397) <= base.Ui32(v1457))
	if base.B2i32(v1458 == int32(0))&base.B2i32(base.Ui32(v1453) < base.Ui32(int32(134217728))) != 0 {
		v1427 = v1457
		v1431 = v1454
		goto L329
	} else {
		goto L336
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1443 + int32(1)
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443))))
	v1452 = v1449
	goto L331
L333:
	;
	goto L334
L334:
	;
	v1450 = F___shgetc(m, v43)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L6
	} else {
		goto L335
	}
L335:
	;
	v1452 = v1450
	goto L331
L336:
	;
	goto L330
L337:
	;
	v1466 = base.I64_extend_i32_u(v1417)
	v1467 = int64(base.Ui64(int64(-1)) >> (uint(v1466) % 64))
	if base.Ui64(v1467) < base.Ui64(v1464) {
		v1676 = v1397
		v1680 = v1452
		v1696 = v1464
		goto L257
	} else {
		goto L338
	}
L338:
	;
	v1478 = v1457
	v1490 = v1464
	goto L339
L339:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1497 != v1498 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	v1676 = v1397
	v1680 = v1506
	v1696 = v1508
	goto L257
L341:
	;
	v1508 = v1490<<(uint(v1466)%64) | base.I64_extend_i32_u(v1478)&int64(255)
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1397) <= base.Ui32(v1511) {
		v1676 = v1397
		v1680 = v1506
		v1696 = v1508
		goto L257
	} else {
		goto L346
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1497 + int32(1)
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497))))
	v1506 = v1503
	goto L341
L343:
	;
	goto L344
L344:
	;
	v1504 = F___shgetc(m, v43)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L6
	} else {
		goto L345
	}
L345:
	;
	v1506 = v1504
	goto L341
L346:
	;
	if base.Ui64(v1508) <= base.Ui64(v1467) {
		v1478 = v1511
		v1490 = v1508
		goto L339
	} else {
		goto L347
	}
L347:
	;
	goto L340
L348:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1539 != v1540 {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	v1560 = base.I64_extend_i32_u(v1550)
	if base.Ui32(v1397) <= base.Ui32(v1553) {
		v1676 = v1397
		v1680 = v1548
		v1696 = v1560
		goto L257
	} else {
		goto L356
	}
L350:
	;
	v1550 = v1527 + v1397*v1523
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+uint32(_c_F_vfscanf[1]))))
	v1554 = base.B2i32(base.Ui32(v1397) <= base.Ui32(v1553))
	if base.B2i32(v1554 == int32(0))&base.B2i32(base.Ui32(v1550) < base.Ui32(int32(119304647))) != 0 {
		v1523 = v1550
		v1527 = v1553
		goto L348
	} else {
		goto L355
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1539 + int32(1)
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539))))
	v1548 = v1545
	goto L350
L352:
	;
	goto L353
L353:
	;
	v1546 = F___shgetc(m, v43)
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L6
	} else {
		goto L354
	}
L354:
	;
	v1548 = v1546
	goto L350
L355:
	;
	goto L349
L356:
	;
	v1561 = base.I64_extend_i32_u(v1397)
	v1567 = v1548
	v1575 = v1553
	v1583 = v1560
	goto L357
L357:
	;
	v1587 = v1583 * v1561
	v1590 = base.I64_extend_i32_u(v1575) & int64(255)
	if base.Ui64(v1590^int64(-1)) < base.Ui64(v1587) {
		v1676 = v1397
		v1680 = v1567
		v1696 = v1583
		goto L257
	} else {
		goto L359
	}
L358:
	;
	v1655 = v1603
	v1671 = v1604
	goto L322
L359:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1594 != v1595 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v1604 = v1590 + v1587
	v1607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1603)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1397) <= base.Ui32(v1607) {
		v1676 = v1397
		v1680 = v1603
		v1696 = v1604
		goto L257
	} else {
		goto L365
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1594 + int32(1)
	v1600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1594))))
	v1603 = v1600
	goto L360
L362:
	;
	goto L363
L363:
	;
	v1601 = F___shgetc(m, v43)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L6
	} else {
		goto L364
	}
L364:
	;
	v1603 = v1601
	goto L360
L365:
	;
	v1609 = int64(0)
	v1615 = int64(32)
	v1616 = int64(base.Ui64(v1604) >> (uint(v1615) % 64))
	v1618 = int64(base.Ui64(v1561) >> (uint(v1615) % 64))
	v1621 = int64(4294967295)
	v1622 = v1604 & v1621
	v1624 = v1561 & v1621
	v1625 = v1622 * v1624
	v1629 = int64(base.Ui64(v1625)>>(uint(v1615)%64)) + v1622*v1618
	v1636 = v1624*v1616 + v1629&v1621
	*(*int64)(unsafe.Add(mBase, uint32(v1080)+8)) = v1561*v1609 + v1609*v1604 + v1616*v1618 + int64(base.Ui64(v1629)>>(uint(v1615)%64)) + int64(base.Ui64(v1636)>>(uint(v1615)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1080))) = v1625&v1621 | v1636<<(uint(v1615)%64)
	goto L366
L366:
	;
	v1647 = *(*int64)(unsafe.Add(mBase, uint32(v1080)+8))
	if v1647 == int64(0) {
		v1567 = v1603
		v1575 = v1607
		v1583 = v1604
		goto L357
	} else {
		goto L367
	}
L367:
	;
	goto L358
L368:
	;
	goto L369
L369:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v1729 != v1730 {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vfscanf[0])) = int32(68)
	v1766 = int32(0)
	v1769 = int64(-1)
	goto L256
L371:
	;
	v1741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1738)+uint32(_c_F_vfscanf[1]))))
	if base.Ui32(v1741) < base.Ui32(v1676) {
		goto L369
	} else {
		goto L376
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1729 + int32(1)
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1729))))
	v1738 = v1735
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1736 = F___shgetc(m, v43)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L6
	} else {
		goto L375
	}
L375:
	;
	v1738 = v1736
	goto L371
L376:
	;
	goto L370
L377:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v1776 - int32(1)
	goto L379
L378:
	;
	goto L379
L379:
	;
	goto L380
L380:
	;
	v1782 = base.I64_extend_i32_s(v1766)
	v1805 = v1769 ^ v1782 - v1782
	goto L234
L382:
	;
	v1821 = int32(0)
	if base.B2i32(v351 == v1821)|base.B2i32(v483 != int32(112)) == v1821 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v351))) = uint32(v1805)
	v2735 = v473
	v2737 = v439
	v2743 = v440
	goto L161
L384:
	;
	goto L385
L385:
	;
	if v351 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	v2735 = v473
	v2737 = v439
	v2743 = v440
	goto L161
L387:
	;
	goto L386
L388:
	;
	switch v480 + int32(2) {
	case 0:
		goto L392
	case 1:
		goto L391
	case 2, 3:
		goto L390
	default:
		goto L387
	case 5:
		goto L389
	}
L389:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v1805
	goto L387
L390:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v351))) = uint32(v1805)
	goto L386
L391:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v1805)
	goto L386
L392:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v1805)
	goto L386
L393:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v351))) = base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(v668)>>(uint(int64(32))%64)))&int32(-2147483648) | v1972<<(uint(int32(23))%32) | v1971)
	v2735 = v473
	v2737 = v439
	v2743 = v440
	goto L161
L394:
	;
	m.G0 = v1846 + int32(32)
	goto L393
L395:
	;
	v1861 = base.I32_wrap_i64(int64(base.Ui64(v1849) >> (uint(int64(25)) % 64)))
	v1865 = v668 & int64(33554431)
	v1866 = int64(16777216)
	if v1865 == v1866 {
		goto L399
	} else {
		goto L400
	}
L396:
	;
	goto L397
L397:
	;
	if base.B2i32(v669|v1849 == int64(0))|base.B2i32(v1853 != int64(32767)) == int32(0) {
		goto L412
	} else {
		goto L413
	}
L398:
	;
	v1886 = base.B2i32(base.Ui32(int32(_a_F_vfscanf_2)) < base.Ui32(v1883))
	if base.Ui32(int32(_a_F_vfscanf_2)) < base.Ui32(v1883) {
		goto L406
	} else {
		goto L407
	}
L399:
	;
	v1870 = base.B2i32(v669 == int64(0))
	goto L401
L400:
	;
	v1870 = base.B2i32(base.Ui64(v1865) < base.Ui64(v1866))
	goto L401
L401:
	;
	if v1870 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1883 = v1861 + int32(1)
	goto L398
L403:
	;
	goto L404
L404:
	;
	if v669|(v1865^int64(16777216)) != int64(0) {
		v1883 = v1861
		goto L398
	} else {
		goto L405
	}
L405:
	;
	v1883 = v1861&int32(1) + v1861
	goto L398
L406:
	;
	v1887 = int32(0)
	goto L408
L407:
	;
	v1887 = v1883
	goto L408
L408:
	;
	if base.Ui32(int32(_a_F_vfscanf_2)) < base.Ui32(v1883) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1890 = int32(-16255)
	goto L411
L410:
	;
	v1890 = int32(-16256)
	goto L411
L411:
	;
	v1971 = v1887
	v1972 = v1890 + v1854
	goto L394
L412:
	;
	v1971 = base.I32_wrap_i64(int64(base.Ui64(v1849)>>(uint(int64(25))%64))) | int32(_a_F_vfscanf_3)
	v1972 = int32(255)
	goto L394
L413:
	;
	goto L414
L414:
	;
	if base.Ui32(int32(_a_F_vfscanf_4)) < base.Ui32(v1854) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1971 = int32(0)
	v1972 = int32(255)
	goto L394
L416:
	;
	goto L417
L417:
	;
	v1912 = base.B2i32(v1853 == int64(0))
	if v1853 == int64(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1913 = int32(_a_F_vfscanf_5)
	goto L420
L419:
	;
	v1913 = int32(_a_F_vfscanf_1)
	goto L420
L420:
	;
	v1914 = v1913 - v1854
	if int32(112) < v1914 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1917 = int32(0)
	v1971 = v1917
	v1972 = v1917
	goto L394
L422:
	;
	goto L423
L423:
	;
	if v1853 == int64(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1921 = v1849
	goto L426
L425:
	;
	v1921 = v1849 | int64(281474976710656)
	goto L426
L426:
	;
	if v1854 != v1913 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	F___ashlti3(m, v1846+int32(16), v669, v1921, int32(128)-v1914)
	mBase = m.M
	v1929 = *(*int64)(unsafe.Add(mBase, uint32(v1846)+16))
	v1930 = *(*int64)(unsafe.Add(mBase, uint32(v1846)+24))
	v1934 = base.B2i32(v1929|v1930 != int64(0))
	goto L429
L428:
	;
	v1934 = int32(0)
	goto L429
L429:
	;
	F___lshrti3(m, v1846, v669, v1921, v1914)
	mBase = m.M
	v1936 = *(*int64)(unsafe.Add(mBase, uint32(v1846)+8))
	v1939 = base.I32_wrap_i64(int64(base.Ui64(v1936) >> (uint(int64(25)) % 64)))
	v1940 = *(*int64)(unsafe.Add(mBase, uint32(v1846)))
	v1942 = v1940 | base.I64_extend_i32_u(v1934)
	v1946 = v1936 & int64(33554431)
	v1947 = int64(16777216)
	if v1946 == v1947 {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	v1968 = base.B2i32(base.Ui32(int32(_a_F_vfscanf_2)) < base.Ui32(v1964))
	if base.Ui32(int32(_a_F_vfscanf_2)) < base.Ui32(v1964) {
		goto L438
	} else {
		goto L439
	}
L431:
	;
	v1951 = base.B2i32(v1942 == int64(0))
	goto L433
L432:
	;
	v1951 = base.B2i32(base.Ui64(v1946) < base.Ui64(v1947))
	goto L433
L433:
	;
	if v1951 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1964 = v1939 + int32(1)
	goto L430
L435:
	;
	goto L436
L436:
	;
	if v1942|(v1946^int64(16777216)) != int64(0) {
		v1964 = v1939
		goto L430
	} else {
		goto L437
	}
L437:
	;
	v1964 = v1939&int32(1) + v1939
	goto L430
L438:
	;
	v1969 = v1964 ^ int32(_a_F_vfscanf_6)
	goto L440
L439:
	;
	v1969 = v1964
	goto L440
L440:
	;
	v1971 = v1969
	v1972 = v1968
	goto L394
L441:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v351))) = base.F64_reinterpret_i64(v668&int64(-9223372036854775807-1) | v2119<<(uint(int64(52))%64) | v2112)
	v2735 = v473
	v2737 = v439
	v2743 = v440
	goto L161
L442:
	;
	m.G0 = v1999 + int32(32)
	goto L441
L443:
	;
	v2016 = v2002<<(uint(int64(4))%64) | int64(base.Ui64(v669)>>(uint(int64(60))%64))
	v2021 = v669 & int64(1152921504606846975)
	if base.Ui64(int64(576460752303423489)) <= base.Ui64(v2021) {
		goto L447
	} else {
		goto L448
	}
L444:
	;
	goto L445
L445:
	;
	if base.B2i32(v669|v2002 == int64(0))|base.B2i32(v2006 != int64(32767)) == int32(0) {
		goto L454
	} else {
		goto L455
	}
L446:
	;
	v2034 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v2031))
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v2031) {
		goto L451
	} else {
		goto L452
	}
L447:
	;
	v2031 = v2016 + int64(1)
	goto L446
L448:
	;
	goto L449
L449:
	;
	if v2021 != int64(576460752303423488) {
		v2031 = v2016
		goto L446
	} else {
		goto L450
	}
L450:
	;
	v2031 = v2016&int64(1) + v2016
	goto L446
L451:
	;
	v2035 = int64(0)
	goto L453
L452:
	;
	v2035 = v2031
	goto L453
L453:
	;
	v2112 = v2035
	v2119 = base.I64_extend_i32_u(v2034) + base.I64_extend_i32_u(v2007-int32(_a_F_vfscanf_7))
	goto L442
L454:
	;
	v2112 = v2002<<(uint(int64(4))%64) | int64(base.Ui64(v669)>>(uint(int64(60))%64)) | int64(2251799813685248)
	v2119 = int64(2047)
	goto L442
L455:
	;
	goto L456
L456:
	;
	if base.Ui32(int32(_a_F_vfscanf_8)) < base.Ui32(v2007) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2112 = int64(0)
	v2119 = int64(2047)
	goto L442
L458:
	;
	goto L459
L459:
	;
	v2061 = base.B2i32(v2006 == int64(0))
	if v2006 == int64(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2062 = int32(_a_F_vfscanf_7)
	goto L462
L461:
	;
	v2062 = int32(_a_F_vfscanf_0)
	goto L462
L462:
	;
	v2063 = v2062 - v2007
	if int32(112) < v2063 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2066 = int64(0)
	v2112 = v2066
	v2119 = v2066
	goto L442
L464:
	;
	goto L465
L465:
	;
	if v2006 == int64(0) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2070 = v2002
	goto L468
L467:
	;
	v2070 = v2002 | int64(281474976710656)
	goto L468
L468:
	;
	if v2007 != v2062 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	F___ashlti3(m, v1999+int32(16), v669, v2070, int32(128)-v2063)
	mBase = m.M
	v2078 = *(*int64)(unsafe.Add(mBase, uint32(v1999)+16))
	v2079 = *(*int64)(unsafe.Add(mBase, uint32(v1999)+24))
	v2083 = base.B2i32(v2078|v2079 != int64(0))
	goto L471
L470:
	;
	v2083 = int32(0)
	goto L471
L471:
	;
	F___lshrti3(m, v1999, v669, v2070, v2063)
	mBase = m.M
	v2085 = *(*int64)(unsafe.Add(mBase, uint32(v1999)+8))
	v2088 = *(*int64)(unsafe.Add(mBase, uint32(v1999)))
	v2091 = v2085<<(uint(int64(4))%64) | int64(base.Ui64(v2088)>>(uint(int64(60))%64))
	v2095 = base.I64_extend_i32_u(v2083) | v2088&int64(1152921504606846975)
	if base.Ui64(int64(576460752303423489)) <= base.Ui64(v2095) {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	v2109 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v2105))
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v2105) {
		goto L477
	} else {
		goto L478
	}
L473:
	;
	v2105 = v2091 + int64(1)
	goto L472
L474:
	;
	goto L475
L475:
	;
	if v2095 != int64(576460752303423488) {
		v2105 = v2091
		goto L472
	} else {
		goto L476
	}
L476:
	;
	v2105 = v2091&int64(1) + v2091
	goto L472
L477:
	;
	v2110 = v2105 ^ int64(4503599627370496)
	goto L479
L478:
	;
	v2110 = v2105
	goto L479
L479:
	;
	v2112 = v2110
	v2119 = base.I64_extend_i32_u(v2109)
	goto L442
L480:
	;
	v2163 = int32(31)
	goto L482
L481:
	;
	v2163 = v585 + int32(1)
	goto L482
L482:
	;
	if v480 == int32(1) {
		goto L484
	} else {
		goto L485
	}
L483:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2698 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	if int64(0) <= v2698 {
		goto L592
	} else {
		goto L593
	}
L484:
	;
	if v441 != 0 {
		goto L487
	} else {
		goto L488
	}
L485:
	;
	goto L486
L486:
	;
	if v441 != 0 {
		goto L553
	} else {
		goto L554
	}
L487:
	;
	v2168 = F_emscripten_builtin_malloc(m, v2163<<(uint(int32(2))%32))
	mBase = m.M
	if v2168 == int32(0) {
		goto L14
	} else {
		goto L490
	}
L488:
	;
	v2171 = v351
	goto L489
L489:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50)+296)) = int64(0)
	v2176 = int32(0)
	v2180 = v2163
	v2183 = v2171
	goto L493
L490:
	;
	v2171 = v2168
	goto L489
L491:
	;
	v2805 = v2183
	v2811 = v2494
	v2815 = v441
	goto L11
L492:
	;
	v2474 = int32(0)
	v2476 = v50 + int32(296)
	if v2476 != 0 {
		goto L549
	} else {
		goto L550
	}
L493:
	;
	v2201 = v2176
	goto L495
L494:
	;
	v2805 = v2183
	v2811 = int32(0)
	v2815 = int32(1)
	goto L11
L495:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v2225 != v2226 {
		goto L498
	} else {
		goto L499
	}
L496:
	;
	v2465 = int32(1)
	v2468 = v2180<<(uint(v2465)%32) | v2465
	v2471 = F_emscripten_builtin_realloc(m, v2183, v2468<<(uint(int32(2))%32))
	mBase = m.M
	if v2471 != 0 {
		v2176 = v2460
		v2180 = v2468
		v2183 = v2471
		goto L493
	} else {
		goto L548
	}
L497:
	;
	v2236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v2234)+33)))
	if v2236 == int32(0) {
		goto L492
	} else {
		goto L502
	}
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v2225 + int32(1)
	v2231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2225))))
	v2234 = v2231
	goto L497
L499:
	;
	goto L500
L500:
	;
	v2232 = F___shgetc(m, v43)
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L6
	} else {
		goto L501
	}
L501:
	;
	v2234 = v2232
	goto L497
L502:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+27)) = uint8(v2234)
	v2241 = v50 + int32(28)
	v2243 = v50 + int32(296)
	if v2243 != 0 {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	if v2447 == int32(-2) {
		goto L495
	} else {
		goto L540
	}
L504:
	;
	v2245 = v2243
	goto L506
L505:
	;
	v2245 = int32(_a_F_vfscanf_9)
	goto L506
L506:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2245)))
	v2248 = v50 + int32(27)
	if v2248 == int32(0) {
		goto L510
	} else {
		goto L511
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245))) = v2404
	v2447 = int32(-2)
	goto L503
L508:
	;
	v2447 = v2394
	goto L503
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_vfscanf[0])) = int32(25)
	v2394 = int32(-1)
	goto L508
L510:
	;
	if v2246 != 0 {
		goto L509
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	if v2246 != 0 {
		goto L515
	} else {
		goto L516
	}
L513:
	;
	v2447 = int32(0)
	goto L503
L514:
	;
	v2279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2248))))
	v2281 = int32(base.Ui32(v2279) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v2281-int32(16)|(v2246>>(uint(int32(26))%32)+v2281)) {
		goto L509
	} else {
		goto L529
	}
L515:
	;
	goto L514
L516:
	;
	goto L517
L517:
	;
	v2253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2248))))
	v2254 = base.I32_extend8_s(v2253)
	if int32(0) <= v2254 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	if v2241 != 0 {
		goto L521
	} else {
		goto L522
	}
L519:
	;
	goto L520
L520:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, _c_F_vfscanf[3]))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2261)))
	if v2262 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241))) = v2253
	goto L523
L522:
	;
	goto L523
L523:
	;
	v2447 = base.B2i32(v2254 != int32(0))
	goto L503
L524:
	;
	if v2241 == int32(0) {
		v2394 = int32(1)
		goto L508
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v2273 = v2253 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v2273) {
		goto L509
	} else {
		goto L528
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241))) = v2254 & int32(_a_F_vfscanf_10)
	v2447 = int32(1)
	goto L503
L528:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2273<<(uint(int32(2))%32))+uint32(_c_F_vfscanf[4])))
	v2404 = v2278
	goto L507
L529:
	;
	v2299 = v2246
	v2302 = v2279
	v2303 = int32(1)
	v2308 = v2248
	goto L530
L530:
	;
	v2316 = v2303 - int32(1)
	v2323 = v2302&int32(255) - int32(128) | v2299<<(uint(int32(6))%32)
	if int32(0) <= v2323 {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	goto L509
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245))) = int32(0)
	if v2241 != 0 {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	goto L534
L534:
	;
	if v2316 == int32(0) {
		v2404 = v2323
		goto L507
	} else {
		goto L538
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241))) = v2323
	goto L537
L536:
	;
	goto L537
L537:
	;
	v2447 = int32(1) - v2316
	goto L503
L538:
	;
	v2334 = v2308 + int32(1)
	v2335 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2334))))
	if v2335 < int32(-64) {
		v2299 = v2323
		v2302 = v2335
		v2303 = v2316
		v2308 = v2334
		goto L530
	} else {
		goto L539
	}
L539:
	;
	goto L531
L540:
	;
	if v2447 == int32(-1) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2494 = int32(0)
	goto L491
L542:
	;
	goto L543
L543:
	;
	if v2183 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2183+v2201<<(uint(int32(2))%32)))) = v2456
	v2460 = v2201 + int32(1)
	goto L546
L545:
	;
	v2460 = v2201
	goto L546
L546:
	;
	if base.B2i32(v441 == int32(0))|base.B2i32(v2460 != v2180) != 0 {
		v2201 = v2460
		goto L495
	} else {
		goto L547
	}
L547:
	;
	goto L496
L548:
	;
	goto L494
L549:
	;
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v2476)))
	v2479 = v2477
	goto L551
L550:
	;
	v2479 = int32(0)
	goto L551
L551:
	;
	if v2479 == int32(0) {
		v2673 = v2201
		v2675 = v2183
		v2678 = v2183
		v2684 = v2474
		goto L483
	} else {
		goto L552
	}
L552:
	;
	v2494 = v2474
	goto L491
L553:
	;
	v2507 = int32(0)
	v2508 = F_emscripten_builtin_malloc(m, v2163)
	mBase = m.M
	if v2508 == v2507 {
		goto L14
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	if v351 != 0 {
		goto L571
	} else {
		goto L572
	}
L556:
	;
	v2512 = v2507
	v2516 = v2163
	v2519 = v2508
	goto L557
L557:
	;
	v2537 = v2512
	goto L559
L558:
	;
	v2805 = int32(0)
	v2811 = v2519
	v2815 = int32(1)
	goto L11
L559:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v2561 != v2562 {
		goto L562
	} else {
		goto L563
	}
L560:
	;
	v2581 = int32(1)
	v2584 = v2516<<(uint(v2581)%32) | v2581
	v2585 = F_emscripten_builtin_realloc(m, v2519, v2584)
	mBase = m.M
	if v2585 != 0 {
		v2512 = v2579
		v2516 = v2584
		v2519 = v2585
		goto L557
	} else {
		goto L570
	}
L561:
	;
	v2572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v2570)+33)))
	if v2572 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v2561 + int32(1)
	v2567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561))))
	v2570 = v2567
	goto L561
L563:
	;
	goto L564
L564:
	;
	v2568 = F___shgetc(m, v43)
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L6
	} else {
		goto L565
	}
L565:
	;
	v2570 = v2568
	goto L561
L566:
	;
	v2673 = v2537
	v2675 = v2519
	v2678 = int32(0)
	v2684 = v2519
	goto L483
L567:
	;
	goto L568
L568:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2537+v2519))) = uint8(v2570)
	v2579 = v2537 + int32(1)
	if v2579 != v2516 {
		v2537 = v2579
		goto L559
	} else {
		goto L569
	}
L569:
	;
	goto L560
L570:
	;
	goto L558
L571:
	;
	v2590 = int32(0)
	goto L574
L572:
	;
	goto L573
L573:
	;
	goto L584
L574:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v2614 != v2615 {
		goto L577
	} else {
		goto L578
	}
L576:
	;
	v2625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v2623)+33)))
	if v2625 != 0 {
		goto L581
	} else {
		goto L582
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v2614 + int32(1)
	v2620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2614))))
	v2623 = v2620
	goto L576
L578:
	;
	goto L579
L579:
	;
	v2621 = F___shgetc(m, v43)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L6
	} else {
		goto L580
	}
L580:
	;
	v2623 = v2621
	goto L576
L581:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2590+v351))) = uint8(v2623)
	v2590 = v2590 + int32(1)
	goto L574
L582:
	;
	v2673 = v2590
	v2675 = v351
	v2678 = int32(0)
	v2684 = v351
	goto L483
L584:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v2656 != v2657 {
		goto L587
	} else {
		goto L588
	}
L585:
	;
	v2668 = int32(0)
	v2673 = v2668
	v2675 = v2668
	v2678 = v2668
	v2684 = v2668
	goto L483
L586:
	;
	v2667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2665+v50)+33)))
	if v2667 != 0 {
		goto L584
	} else {
		goto L591
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v2656 + int32(1)
	v2662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2656))))
	v2665 = v2662
	goto L586
L588:
	;
	goto L589
L589:
	;
	v2663 = F___shgetc(m, v43)
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L6
	} else {
		goto L590
	}
L590:
	;
	v2665 = v2663
	goto L586
L591:
	;
	goto L585
L592:
	;
	v2702 = v2697 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v2702
	v2704 = v2702
	goto L594
L593:
	;
	v2704 = v2697
	goto L594
L594:
	;
	v2705 = *(*int64)(unsafe.Add(mBase, uint32(v43)+120))
	v2706 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	v2709 = v2705 + base.I64_extend_i32_s(v2704-v2706)
	if base.B2i32(v2709 == int64(0))|base.B2i32(v2162|base.B2i32(v2709 == v602) == int32(0)) != 0 {
		v2832 = v2678
		v2838 = v2684
		v2842 = v441
		v2845 = v62
		goto L10
	} else {
		goto L595
	}
L595:
	;
	if v441 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v2675
	goto L598
L597:
	;
	goto L598
L598:
	;
	if v483 == int32(99) {
		v2735 = v2137
		v2737 = v2678
		v2743 = v2684
		goto L161
	} else {
		goto L599
	}
L599:
	;
	if v2678 != 0 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2678+v2673<<(uint(int32(2))%32)))) = int32(0)
	goto L602
L601:
	;
	goto L602
L602:
	;
	if v2684 == int32(0) {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v2735 = v2137
	v2737 = v2678
	v2743 = int32(0)
	goto L161
L604:
	;
	goto L605
L605:
	;
	v2729 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2673+v2684))) = uint8(v2729)
	v2735 = v2137
	v2737 = v2678
	v2743 = v2684
	goto L161
L606:
	;
	v2888 = v2773
	v2900 = v2785
	goto L1
L607:
	;
	v2825 = v62
	goto L609
L608:
	;
	v2825 = int32(-1)
	goto L609
L609:
	;
	v2832 = v2805
	v2838 = v2811
	v2842 = v2815
	v2845 = v2825
	goto L10
L610:
	;
	F_emscripten_builtin_free(m, v2838)
	mBase = m.M
	F_emscripten_builtin_free(m, v2832)
	mBase = m.M
	v2888 = v50
	v2900 = v2845
	goto L1
}
func F_view_reloptions(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v7 = F_build_reloptions(m, l0, int32(1), int32(512), int32(12), int32(_a_F_view_reloptions_0), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_visibilitymap_pin_ok(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	if l1 == int32(0) {
		return int32(0)
	} else {
		if l1 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_pin_ok[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v25 = v16
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_pin_ok[1]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v25 = v24
		}
		v27 = base.I32_div_u_s(l0, int32(_a_F_visibilitymap_pin_ok_0))
		return base.B2i32(v25 == v27)
	}
}
func F_void_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13 << (uint(int32(2)) % 32)
		m.G0 = v5 + int32(16)
		return v12
	}
}
