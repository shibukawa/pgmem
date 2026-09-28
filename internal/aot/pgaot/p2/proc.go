package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_EmitProcSignalBarrier(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	if v2 < v7+int32(38) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v43 = int32(0)
	v44 = base.AtomicRmwAdd64(m, v41, v43, int64(1))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	v48 = v46 + int32(37)
	if v43 <= v48 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v27 = base.AtomicRmwOr32(m, v20+v15*int32(112)+int32(104), int32(0), int32(1)<<(uint(l0)%32))
	v29 = v15 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	if v29 < v31+int32(38) {
		v15 = v29
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	v51 = v48
	goto L10
L8:
	;
	goto L9
L9:
	;
	return v44 + int64(1)
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v60 = v57 + v51*int32(112)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	if int32(0) < v51 {
		v51 = v51 - int32(1)
		goto L10
	} else {
		goto L22
	}
L13:
	;
	v65 = v60 + int32(8)
	v67 = v60 + int32(88)
	v70 = base.AtomicRmwXchg32(m, v65, int32(80), int32(1))
	if v70 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v67, int32(_a_F_EmitProcSignalBarrier_0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v76 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return int64(0)
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+56)) = int32(1)
	v79 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v65)+80)), uint32(v79))
	v83 = F_pgmem_kill(m, v76, int32(10))
	mBase = m.M
	goto L12
L20:
	;
	goto L21
L21:
	;
	v84 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v67))), uint32(v84))
	goto L12
L22:
	;
	goto L11
}
func F_ProcArrayApplyRecoveryInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v586 int32
	_ = v586
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int64
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v24 = F_LWLockAcquire(m, v20+int32(512), v2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = v18
	goto L3
L3:
	;
	v40 = v27 - int32(1)
	if base.Ui32(v40) < base.Ui32(int32(3)) {
		v27 = v40
		goto L3
	} else {
		goto L5
	}
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v49 = int32(0)
	if base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v45))&base.B2i32(v49 <= v45-v40) == v49 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = v54 + base.I64_extend_i32_s(v40-base.I32_wrap_i64(v54))
	goto L8
L7:
	;
	goto L8
L8:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v61 + int64(1)
	v65 = int32(3)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if base.B2i32(base.Ui32(v18) < base.Ui32(v65))|base.B2i32(base.Ui32(v69) < base.Ui32(v65)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v81 = m.G0
	v83 = v81 - int32(16)
	m.G0 = v83
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v89 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
	goto L9
L11:
	;
	if v69-v18 < int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v18) <= base.Ui32(v69) {
		goto L9
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L10
L16:
	;
	if v18 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	m.G0 = v83 + int32(16)
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v329+int32(512))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L64
	}
L18:
	;
	if v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+12)) = int64(0)
	goto L17
L24:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(_a_F_ProcArrayApplyRecoveryInfo_2), int32(_a_F_ProcArrayApplyRecoveryInfo_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v18
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_4), v83)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	if v115 < v116 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(_a_F_ProcArrayApplyRecoveryInfo_5), int32(_a_F_ProcArrayApplyRecoveryInfo_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v213
	v224 = int32(0)
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	if v227-v228 == v230 {
		goto L17
	} else {
		goto L53
	}
L32:
	;
	v121 = v115
	v125 = v2
	goto L35
L33:
	;
	goto L34
L34:
	;
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v207
	v213 = v207
	goto L31
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v121))))
	if v136 != int32(1) {
		v163 = v125
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v171 - v169
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v179 = v115
	goto L49
L37:
	;
	goto L36
L38:
	;
	v166 = v121 + int32(1)
	if v166 != v116 {
		v121 = v166
		v125 = v163
		goto L35
	} else {
		goto L48
	}
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[4]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v121<<(uint(int32(2))%32))))
	if base.B2i32(base.Ui32(v18) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v144) < base.Ui32(int32(3))) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v154 = F_StandbyTransactionIdIsPrepared(m, v144)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	if v144-v18 < int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v18) <= base.Ui32(v144) {
		v169 = v125
		goto L37
	} else {
		goto L45
	}
L44:
	;
	v169 = v125
	goto L37
L45:
	;
	goto L40
L46:
	;
	if v154 != 0 {
		v163 = v125
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157+v121))) = uint8(v159)
	v163 = v125 + int32(1)
	goto L38
L48:
	;
	v169 = v163
	goto L37
L49:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v179))))
	if v190 != 0 {
		v213 = v179
		goto L31
	} else {
		goto L51
	}
L50:
	;
	goto L34
L51:
	;
	v192 = v179 + int32(1)
	if v192 != v116 {
		v179 = v192
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	if v228 < v227 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v236 = v228
	v237 = v234
	v238 = v224
	goto L57
L55:
	;
	v279 = v224
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v226)+16)) = int32(0)
	v296 = m.G0
	v297 = int32(16)
	v298 = v296 - v297
	m.G0 = v298
	F_gettimeofday(m, v298)
	mBase = m.M
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v298)+8)))
	m.G0 = v298 + v297
	goto L63
L57:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v237))))
	if v249 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v279 = v272
	goto L56
L59:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[4]))
	v254 = int32(2)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v253+v236<<(uint(v254)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v253+v238<<(uint(v254)%32)))) = v260
	v262 = int32(_a_F_ProcArrayApplyRecoveryInfo_6)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v263+v238))) = uint8(v265)
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v271 = v270
	v272 = v238 + v265
	goto L61
L60:
	;
	v271 = v237
	v272 = v238
	goto L61
L61:
	;
	v274 = v236 + int32(1)
	if v274 != v227 {
		v236 = v274
		v237 = v271
		v238 = v272
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[5])) = v302 + v301*int64(1000000) - int64(946684800000000)
	goto L17
L64:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v336 = v334
	goto L65
L65:
	;
	v349 = v336 - int32(1)
	if base.Ui32(v349) < base.Ui32(int32(3)) {
		v336 = v349
		goto L65
	} else {
		goto L67
	}
L66:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v349)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v355 = m.G0
	v357 = v355 - int32(32)
	m.G0 = v357
	v360 = v357 + int32(12)
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[6]))
	F_hash_seq_init(m, v360, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v365 = F_hash_seq_search(m, v360)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v365 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v370 = v365
	goto L74
L72:
	;
	goto L73
L73:
	;
	m.G0 = v357 + int32(32)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7]))
	switch v426 - int32(2) {
	case 0:
		goto L93
	case 1:
		goto L89
	default:
		goto L92
	}
L74:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v383 = F_StandbyTransactionIdIsPrepared(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v406 = F_hash_seq_search(m, v357+int32(12))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L87
	}
L77:
	;
	if v383 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	if base.B2i32(base.Ui32(v354) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v385) < base.Ui32(int32(3))) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	F_StandbyReleaseXidEntryLocks(m, v370)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L85
	}
L80:
	;
	if v385-v354 < int32(0) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(v354) <= base.Ui32(v385) {
		goto L76
	} else {
		goto L84
	}
L83:
	;
	goto L76
L84:
	;
	goto L79
L85:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[6]))
	v401 = F_hash_search(m, v398, v370, int32(2), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L76
L87:
	;
	if v406 != 0 {
		v370 = v406
		goto L74
	} else {
		goto L88
	}
L88:
	;
	goto L75
L89:
	;
	m.G0 = v16 + int32(48)
	return
L90:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), v842, int32(_a_F_ProcArrayApplyRecoveryInfo_7))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L192
	}
L91:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	v786 = int32(3)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(base.Ui32(v785) < base.Ui32(v786))|base.B2i32(base.Ui32(v788) < base.Ui32(v786)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L92:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v461 = F_LWLockAcquire(m, v457+int32(512), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L100
	}
L93:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v429 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v432 != 0 {
		goto L91
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v440 = F_LWLockAcquire(m, v436+int32(512), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v434)+12)) = int64(0)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v447+int32(512))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(1)
	goto L92
L100:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v467 = F_palloc_mul(m, int32(4), v464+v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v469+v470 <= int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v766+int32(512))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L175
	}
L103:
	;
	F_pfree(m, v467)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L133
	}
L104:
	;
	v476 = int32(0)
	v483 = v2
	goto L105
L105:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v488+v476<<(uint(int32(2))%32))))
	v493 = F_TransactionIdDidCommit(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	if v503 <= int32(0) {
		goto L103
	} else {
		goto L113
	}
L107:
	;
	v505 = v476 + int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v505 < v506+v507 {
		v476 = v505
		v483 = v503
		goto L105
	} else {
		goto L112
	}
L108:
	;
	if v493 != 0 {
		v503 = v483
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v495 = F_TransactionIdDidAbort(m, v492)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v495 != 0 {
		v503 = v483
		goto L107
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467+v483<<(uint(int32(2))%32)))) = v492
	v503 = v483 + int32(1)
	goto L107
L112:
	;
	goto L106
L113:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	if v514 != 0 {
		goto L102
	} else {
		goto L114
	}
L114:
	;
	F_pg_qsort(m, v467, v503, int32(4), int32(1200))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v519 = int32(1)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	F_KnownAssignedXidsAdd(m, v520, v520, v519)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v503 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v527 = v519
	goto L120
L118:
	;
	goto L119
L119:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L132
	}
L120:
	;
	v541 = v467 + v527<<(uint(int32(2))%32)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v541-int32(4))))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	if v544 == v545 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L119
L122:
	;
	v569 = v527 + int32(1)
	if v569 != v503 {
		v527 = v569
		goto L120
	} else {
		goto L131
	}
L123:
	;
	v549 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_KnownAssignedXidsAdd(m, v545, v545, int32(1))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L130
	}
L126:
	;
	if v549 == int32(0) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v553
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_8), v16+int32(16))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(1208), int32(_a_F_ProcArrayApplyRecoveryInfo_7))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	goto L122
L131:
	;
	goto L121
L132:
	;
	goto L103
L133:
	;
	v602 = int32(_a_F_ProcArrayApplyRecoveryInfo_9)
	v603 = int32(3)
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	v607 = v605 + int32(1)
	if base.Ui32(v607) <= base.Ui32(v603) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v610 = v603
	goto L136
L135:
	;
	v610 = v607
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v610
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(base.Ui32(v612) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v610-v612) != 0 {
		v651 = v610
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v664 = v651
	goto L147
L138:
	;
	v620 = v610
	goto L139
L139:
	;
	F_ExtendSUBTRANS(m, v620)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	v651 = v642
	goto L137
L141:
	;
	v634 = int32(_a_F_ProcArrayApplyRecoveryInfo_9)
	v635 = int32(3)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	v639 = v637 + int32(1)
	if base.Ui32(v639) <= base.Ui32(v635) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v642 = v635
	goto L144
L143:
	;
	v642 = v639
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v644) < base.Ui32(int32(3)) {
		v651 = v642
		goto L137
	} else {
		goto L145
	}
L145:
	;
	if v642-v644 < int32(0) {
		v620 = v642
		goto L139
	} else {
		goto L146
	}
L146:
	;
	goto L140
L147:
	;
	v677 = v664 - int32(1)
	if base.Ui32(v677) < base.Ui32(int32(3)) {
		v664 = v677
		goto L147
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v677
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v682 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L148
L150:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v709)+8))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v709)+48))
	if v711 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8])) = v677
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(2)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v691)+24)) = v677
	goto L150
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8])) = int32(0)
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	if v682 == int32(2) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+24)) = v677
	goto L150
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+24)) = int32(0)
	goto L150
L157:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v731+int32(512))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L165
	}
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v709)+48)) = v710 + base.I64_extend_i32_s(v707-base.I32_wrap_i64(v710))
	goto L157
L159:
	;
	v714 = int32(3)
	if base.B2i32(base.Ui32(v707) < base.Ui32(v714))|base.B2i32(base.Ui32(v711) < base.Ui32(v714)) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v711-v707 < int32(0) {
		goto L158
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if base.Ui32(v707) <= base.Ui32(v711) {
		goto L157
	} else {
		goto L164
	}
L163:
	;
	goto L157
L164:
	;
	goto L158
L165:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7]))
	v743 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	if v740 == int32(3) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v743 == int32(0) {
		goto L89
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	if v743 == int32(0) {
		goto L89
	} else {
		goto L173
	}
L171:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_10), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v842 = int32(1295)
	goto L90
L173:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v758
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v756
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_11), v16)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v842 = int32(1301)
	goto L90
L175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_12), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(1183), int32(_a_F_ProcArrayApplyRecoveryInfo_7))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v814 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L189
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(3)
	v803 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L186
	}
L181:
	;
	if v785-v788 < int32(0) {
		goto L180
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	if base.Ui32(v788) <= base.Ui32(v785) {
		goto L179
	} else {
		goto L185
	}
L184:
	;
	goto L179
L185:
	;
	goto L180
L186:
	;
	if v803 == int32(0) {
		goto L89
	} else {
		goto L187
	}
L187:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_10), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v842 = int32(1116)
	goto L90
L189:
	;
	if v814 == int32(0) {
		goto L89
	} else {
		goto L190
	}
L190:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v818
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_11), v16+int32(32))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v842 = int32(1123)
	goto L90
L192:
	;
	goto L89
}
func F_ProcArrayInstallImportedXmin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[1]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v105 = v3
	goto L3
L3:
	;
	return v105
L4:
	;
	return int32(0)
L5:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v23 <= v22 {
		v95 = v22
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[1]))
	F_LWLockRelease(m, v97+int32(512))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[2]))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[3]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[4]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v40 = v3
	goto L8
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v34))))
	if v48&int32(2) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v95 = int32(0)
	goto L6
L10:
	;
	v82 = v40 + int32(1)
	if v82 != v23 {
		v40 = v82
		goto L8
	} else {
		goto L16
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v40<<(uint(int32(2))%32))))
	v57 = v31 + v54*int32(768)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v58 != v59 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v61 != v62 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v64 != v29 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	if base.B2i32(base.Ui32(l0) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v66) < base.Ui32(int32(3)))|base.B2i32(int32(0) < v66-l0) != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[5])) = l0
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+52)) = l0
	v95 = int32(1)
	goto L6
L16:
	;
	goto L9
L17:
	;
	v105 = v95
	goto L3
}
func F_ProcArrayShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemAttach[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemAttach[1])) = v5
	return
}
func F_ProcArrayShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
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
	var v73 int32
	_ = v73
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[0])))
	if v8 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(_a_F_ProcArrayShmemRequest_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[1]))
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[2]))
		v21 = F_mul_size(m, int32(4), (v15+v17)*int32(65))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = int32(_a_F_ProcArrayShmemRequest_1)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v21
			F_ShmemRequestStructWithOpts(m, v5+int32(32))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(_a_F_ProcArrayShmemRequest_2)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[1]))
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[2]))
				v42 = F_mul_size(m, int32(1), (v36+v38)*int32(65))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(_a_F_ProcArrayShmemRequest_3)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v42
					F_ShmemRequestStructWithOpts(m, v5+int32(16))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ProcArrayShmemRequest_4)
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[1]))
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[2]))
						v63 = F_mul_size(m, int32(4), v59+v61)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = F_add_size(m, int32(36), v63)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_ProcArrayShmemRequest_5)
								*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v65
								F_ShmemRequestStructWithOpts(m, v5)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									m.G0 = v5 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ProcArrayShmemRequest_4)
		v59 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[1]))
		v61 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayShmemRequest[2]))
		v63 = F_mul_size(m, int32(4), v59+v61)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			v65 = F_add_size(m, int32(36), v63)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_ProcArrayShmemRequest_5)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v65
				F_ShmemRequestStructWithOpts(m, v5)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					m.G0 = v5 + int32(48)
					return
				}
			}
		}
	}
}
func F_ProcGlobalShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[1]))
	v14 = F_add_size(m, int32(38), v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = F_add_size(m, v10, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[2])) = v16
			v22 = F_mul_size(m, v16, int32(768))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = F_add_size(m, int32(0), v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[2]))
					v29 = F_mul_size(m, v27, int32(4))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = F_add_size(m, v24, v29)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[2]))
							v36 = F_mul_size(m, v34, int32(2))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = F_add_size(m, v31, v36)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[2]))
									v43 = F_mul_size(m, v41, int32(1))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										v45 = F_add_size(m, v38, v43)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[3])) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_ProcGlobalShmemRequest_0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_ProcGlobalShmemRequest_1)
											F_ShmemRequestStructWithOpts(m, v6+int32(32))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[4])))
												if v61 == int32(0) {
													v67 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[2]))
													v69 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[5]))
													v72 = F_mul_size(m, v67, v69*int32(72))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														v74 = F_add_size(m, int32(0), v72)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[6])) = v74
															v77 = v74
															*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_ProcGlobalShmemRequest_2)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v77
															*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_ProcGlobalShmemRequest_3)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
															F_ShmemRequestStructWithOpts(m, v6+int32(16))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_ProcGlobalShmemRequest_4)
																*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(80)
																*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ProcGlobalShmemRequest_5)
																F_ShmemRequestStructWithOpts(m, v6)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[0]))
																	v101 = m.G0
																	v103 = v101 - int32(16)
																	m.G0 = v103
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(_a_F_ProcGlobalShmemRequest_6)
																	v108 = F_mul_size(m, v98+int32(38), int32(128))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = int32(_a_F_ProcGlobalShmemRequest_7)
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v108
																		F_ShmemRequestStructWithOpts(m, v103)
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			m.G0 = v103 + int32(16)
																			m.G0 = v6 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													}
												} else {
													v77 = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_ProcGlobalShmemRequest_2)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v77
													*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_ProcGlobalShmemRequest_3)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
													F_ShmemRequestStructWithOpts(m, v6+int32(16))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_ProcGlobalShmemRequest_4)
														*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(80)
														*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ProcGlobalShmemRequest_5)
														F_ShmemRequestStructWithOpts(m, v6)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, _c_F_ProcGlobalShmemRequest[0]))
															v101 = m.G0
															v103 = v101 - int32(16)
															m.G0 = v103
															*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(_a_F_ProcGlobalShmemRequest_6)
															v108 = F_mul_size(m, v98+int32(38), int32(128))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = int32(_a_F_ProcGlobalShmemRequest_7)
																*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v108
																F_ShmemRequestStructWithOpts(m, v103)
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v103 + int32(16)
																	m.G0 = v6 + int32(48)
																	return
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
				}
			}
		}
	}
}
func F_ProcSignalInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
	if int32(0) <= v12 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[1]))
		if v16+int32(38) <= v12 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return
			} else {
				v125 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v125
				v128 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v128 + int32(38)
				F_errmsg_internal(m, int32(_a_F_ProcSignalInit_0), v9+int32(16))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(179), int32(_a_F_ProcSignalInit_2))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
			v24 = v21 + v12*int32(112)
			v26 = v24 + int32(8)
			v29 = base.AtomicRmwXchg32(m, v24, int32(88), int32(1))
			if v29 != 0 {
				F_s_lock(m, v24+int32(88), int32(_a_F_ProcSignalInit_3))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v36 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v24)+72)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v24-int32(-64)))) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v36
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
					v50 = int32(0)
					v51 = base.AtomicRmwXchg32(m, v26, v50, v49)
					*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v50
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
					v59 = base.AtomicRmwCmpxchg64(m, v55, v50, v36, v36)
					v61 = base.AtomicRmwXchg64(m, v26, int32(88), v59)
					if base.B2i32(l1 == v50)|base.B2i32(l1 <= v50) == v50 {
						base.MemoryCopy(m, v24+int32(16), l0, l1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l1
					v73 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+80)), uint32(v73))
					if v35 == v73 {
						*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
						F_on_shmem_exit(m, int32(1204), int64(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					} else {
						v80 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							if v80 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
								F_on_shmem_exit(m, int32(1204), int64(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v85
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
								F_errmsg_internal(m, int32(_a_F_ProcSignalInit_4), v9)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(224), int32(_a_F_ProcSignalInit_2))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
										F_on_shmem_exit(m, int32(1204), int64(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
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
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v36 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v24)+72)) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v24-int32(-64)))) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v36
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
				v50 = int32(0)
				v51 = base.AtomicRmwXchg32(m, v26, v50, v49)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v50
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
				v59 = base.AtomicRmwCmpxchg64(m, v55, v50, v36, v36)
				v61 = base.AtomicRmwXchg64(m, v26, int32(88), v59)
				if base.B2i32(l1 == v50)|base.B2i32(l1 <= v50) == v50 {
					base.MemoryCopy(m, v24+int32(16), l0, l1)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l1
				v73 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+80)), uint32(v73))
				if v35 == v73 {
					*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
					F_on_shmem_exit(m, int32(1204), int64(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				} else {
					v80 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						if v80 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
							F_on_shmem_exit(m, int32(1204), int64(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v85
							v88 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
							F_errmsg_internal(m, int32(_a_F_ProcSignalInit_4), v9)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(224), int32(_a_F_ProcSignalInit_2))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v26
									F_on_shmem_exit(m, int32(1204), int64(0))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v110 = m.ExcPending
		if v110 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ProcSignalInit_5), int32(0))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(177), int32(_a_F_ProcSignalInit_2))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
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
