package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EmitProcSignalBarrier(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	if v1 < v7+int32(38) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v1
	goto L4
L2:
	;
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	v44 = v42 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	v49 = v47 + int32(37)
	if int32(0) <= v49 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v23 = v18 + v12<<(uint(int32(7))%32) + int32(120)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24 | v25
	v29 = v12 + v25
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[0]))
	if v29 < v31+int32(38) {
		v12 = v29
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
	v52 = v49
	goto L10
L8:
	;
	goto L9
L9:
	;
	return v44
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_EmitProcSignalBarrier[1]))
	v61 = v58 + v52<<(uint(int32(7))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v62 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	if int32(0) < v52 {
		v52 = v52 - int32(1)
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v66 = v61 + int32(8)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+96)) = int32(1)
	v71 = v61 + int32(104)
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v71, int32(_a_F_EmitProcSignalBarrier_0), int32(402), int32(_a_F_EmitProcSignalBarrier_1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v79 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v66)+96)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+56)) = int32(1)
	v85 = F_kill(m, v79, int32(10))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(0)
	goto L12
L22:
	;
	goto L12
L23:
	;
	goto L11
}
func F_ProcArrayApplyRecoveryInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v677 int32
	_ = v677
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v716 int32
	_ = v716
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v757 int32
	_ = v757
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int64
	_ = v822
	var v823 int32
	_ = v823
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v25 = F_LWLockAcquire(m, v21+int32(512), v2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = v19
	goto L3
L3:
	;
	v42 = v28 - int32(1)
	if base.Ui32(v42) < base.Ui32(int32(3)) {
		v28 = v42
		goto L3
	} else {
		goto L5
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v48 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L4
L6:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+56)) = v72 + int64(1)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v78)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v42))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v48)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v65 = v46
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v65)+48)) = v47 + base.I64_extend_i32_s(v42-base.I32_wrap_i64(v47))
	v71 = v65
	goto L6
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	if v60 == int32(0) {
		v71 = v62
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v60 = base.B2i32(base.Ui32(v48) < base.Ui32(v42))
	goto L10
L12:
	;
	goto L13
L13:
	;
	v60 = int32(base.Ui32(v48-v42) >> (uint(int32(31)) % 32))
	goto L10
L14:
	;
	v65 = v62
	goto L9
L15:
	;
	if v90 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v90 = base.B2i32(base.Ui32(v78) < base.Ui32(v19))
	goto L15
L17:
	;
	goto L18
L18:
	;
	v90 = int32(base.Ui32(v78-v19) >> (uint(int32(31)) % 32))
	goto L15
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = int32(0)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v95 = m.G0
	v97 = v95 - int32(16)
	m.G0 = v97
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v103 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v19 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	m.G0 = v97 + int32(16)
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v411+int32(512))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L76
	}
L24:
	;
	if v103 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v103 != 0 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_0), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v100)+12)) = int64(0)
	goto L23
L30:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(_a_F_ProcArrayApplyRecoveryInfo_2), int32(_a_F_ProcArrayApplyRecoveryInfo_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v19
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_4), v97)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	if v129 < v130 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(_a_F_ProcArrayApplyRecoveryInfo_5), int32(_a_F_ProcArrayApplyRecoveryInfo_3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = v233
	v245 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	v250 = v248 - v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v250 == v251 {
		goto L23
	} else {
		goto L58
	}
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v135 = v129
	v141 = v133
	v145 = v2
	goto L41
L39:
	;
	goto L40
L40:
	;
	v227 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = v227
	v233 = v227
	goto L37
L41:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v141))))
	if v149 != int32(1) {
		v180 = v141
		v181 = v145
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v189 - v188
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	v197 = v129
	goto L54
L43:
	;
	goto L42
L44:
	;
	v183 = v135 + int32(1)
	if v183 != v130 {
		v135 = v183
		v141 = v180
		v145 = v181
		goto L41
	} else {
		goto L53
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[4]))
	v154 = int32(2)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v135<<(uint(v154)%32))))
	if base.B2i32(base.Ui32(v154) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v157)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v169 != 0 {
		v188 = v145
		goto L43
	} else {
		goto L50
	}
L47:
	;
	v169 = base.B2i32(base.Ui32(v19) <= base.Ui32(v157))
	goto L46
L48:
	;
	goto L49
L49:
	;
	v169 = base.B2i32(int32(0) <= v157-v19)
	goto L46
L50:
	;
	v170 = F_StandbyTransactionIdIsPrepared(m, v157)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	if v170 != 0 {
		v180 = v173
		v181 = v145
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v173))) = uint8(v175)
	v180 = v173
	v181 = v145 + int32(1)
	goto L44
L53:
	;
	v188 = v181
	goto L43
L54:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v197))))
	if v209 != 0 {
		v233 = v197
		goto L37
	} else {
		goto L56
	}
L55:
	;
	goto L40
L56:
	;
	v211 = v197 + int32(1)
	if v211 != v130 {
		v197 = v211
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v248 <= v249 {
		v359 = v245
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247)+20)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v247)+16)) = int32(0)
	v377 = m.G0
	v378 = int32(16)
	v379 = v377 - v378
	m.G0 = v379
	F_gettimeofday(m, v379)
	mBase = m.M
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v379)))
	v383 = int64(*(*int32)(unsafe.Add(mBase, uint32(v379)+8)))
	m.G0 = v379 + v378
	goto L75
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[4]))
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[3]))
	if v249+int32(1) != v248 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v267 = v249
	v269 = v245
	v276 = int32(0)
	goto L64
L62:
	;
	v326 = v249
	v328 = v245
	goto L63
L63:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v257))))
	if v340 != int32(1) {
		v359 = v328
		goto L59
	} else {
		goto L74
	}
L64:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v257))))
	if v281 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v250&int32(1) == int32(0) {
		v359 = v317
		goto L59
	} else {
		goto L73
	}
L66:
	;
	v284 = int32(2)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v255+v267<<(uint(v284)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v269<<(uint(v284)%32)))) = v290
	v293 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v257))) = uint8(v293)
	v297 = v269 + v293
	goto L68
L67:
	;
	v297 = v269
	goto L68
L68:
	;
	v298 = int32(1)
	v299 = v267 + v298
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257+v299))))
	if v301 == v298 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v304 = int32(2)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v255+v299<<(uint(v304)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v297<<(uint(v304)%32)))) = v310
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v297+v257))) = uint8(v313)
	v317 = v297 + v313
	goto L71
L70:
	;
	v317 = v297
	goto L71
L71:
	;
	v318 = int32(2)
	v319 = v267 + v318
	v321 = v276 + v318
	if v321 != v250&int32(-2) {
		v267 = v319
		v269 = v317
		v276 = v321
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v326 = v319
	v328 = v317
	goto L63
L74:
	;
	v343 = int32(2)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v255+v326<<(uint(v343)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v328<<(uint(v343)%32)))) = v349
	v352 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v328+v257))) = uint8(v352)
	v359 = v328 + v352
	goto L59
L75:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[5])) = v383 + v382*int64(1000000) - int64(946684800000000)
	goto L23
L76:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v418 = v416
	goto L77
L77:
	;
	v432 = v418 - int32(1)
	if base.Ui32(v432) < base.Ui32(int32(3)) {
		v418 = v432
		goto L77
	} else {
		goto L79
	}
L78:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v432)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v438 = m.G0
	v440 = v438 - int32(32)
	m.G0 = v440
	v443 = v440 + int32(12)
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[6]))
	F_hash_seq_init(m, v443, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v448 = F_hash_seq_search(m, v443)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v448 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v451 = v448
	goto L86
L84:
	;
	goto L85
L85:
	;
	m.G0 = v440 + int32(32)
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7]))
	switch v513 - int32(2) {
	case 0:
		goto L104
	case 1:
		goto L100
	default:
		goto L103
	}
L86:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v465 = F_StandbyTransactionIdIsPrepared(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L85
L88:
	;
	v492 = F_hash_seq_search(m, v440+int32(12))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L98
	}
L89:
	;
	if v465 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v437))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v467)) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v479 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L92:
	;
	v479 = base.B2i32(base.Ui32(v467) < base.Ui32(v437))
	goto L91
L93:
	;
	goto L94
L94:
	;
	v479 = int32(base.Ui32(v467-v437) >> (uint(int32(31)) % 32))
	goto L91
L95:
	;
	F_StandbyReleaseXidEntryLocks(m, v451)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[6]))
	v488 = F_hash_search(m, v485, v451, int32(2), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L88
L98:
	;
	if v492 != 0 {
		v451 = v492
		goto L86
	} else {
		goto L99
	}
L99:
	;
	goto L87
L100:
	;
	m.G0 = v17 + int32(48)
	return
L101:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), v960, int32(_a_F_ProcArrayApplyRecoveryInfo_6))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L212
	}
L102:
	;
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v902))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v901)) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L103:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v548 = F_LWLockAcquire(m, v544+int32(512), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L111
	}
L104:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v516 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v519 != 0 {
		goto L102
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	v527 = F_LWLockAcquire(m, v523+int32(512), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v521)+12)) = int64(0)
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v534+int32(512))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(1)
	goto L103
L111:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v555 = F_palloc(m, (v550+v551)<<(uint(int32(2))%32))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v557+v558 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v882+int32(512))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L195
	}
L114:
	;
	F_pfree(m, v555)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L144
	}
L115:
	;
	v564 = int32(0)
	v572 = v2
	goto L116
L116:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v577+v564<<(uint(int32(2))%32))))
	v582 = F_TransactionIdDidCommit(m, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	if v592 <= int32(0) {
		goto L114
	} else {
		goto L124
	}
L118:
	;
	v594 = v564 + int32(1)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v594 < v595+v596 {
		v564 = v594
		v572 = v592
		goto L116
	} else {
		goto L123
	}
L119:
	;
	if v582 != 0 {
		v592 = v572
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v584 = F_TransactionIdDidAbort(m, v581)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v584 != 0 {
		v592 = v572
		goto L118
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555+v572<<(uint(int32(2))%32)))) = v581
	v592 = v572 + int32(1)
	goto L118
L123:
	;
	goto L117
L124:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	if v603 != 0 {
		goto L113
	} else {
		goto L125
	}
L125:
	;
	F_pg_qsort(m, v555, v592, int32(4), int32(1105))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v608 = int32(1)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	F_KnownAssignedXidsAdd(m, v609, v609, v608)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v592 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v616 = v608
	goto L131
L129:
	;
	goto L130
L130:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L143
	}
L131:
	;
	v631 = v555 + v616<<(uint(int32(2))%32)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v631-int32(4))))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	if v634 == v635 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L130
L133:
	;
	v659 = v616 + int32(1)
	if v659 != v592 {
		v616 = v659
		goto L131
	} else {
		goto L142
	}
L134:
	;
	v639 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	F_KnownAssignedXidsAdd(m, v635, v635, int32(1))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L141
	}
L137:
	;
	if v639 == int32(0) {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v643
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_7), v17+int32(16))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(1217), int32(_a_F_ProcArrayApplyRecoveryInfo_6))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L133
L141:
	;
	goto L133
L142:
	;
	goto L132
L143:
	;
	goto L114
L144:
	;
	v694 = int32(_a_F_ProcArrayApplyRecoveryInfo_8)
	v695 = int32(3)
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	v699 = v697 + int32(1)
	if base.Ui32(v699) <= base.Ui32(v695) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v702 = v695
	goto L147
L146:
	;
	v702 = v699
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v702
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v704))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v702)) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v716 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v716 = base.B2i32(base.Ui32(v702) < base.Ui32(v704))
	goto L148
L150:
	;
	goto L151
L151:
	;
	v716 = int32(base.Ui32(v702-v704) >> (uint(int32(31)) % 32))
	goto L148
L152:
	;
	goto L155
L153:
	;
	goto L154
L154:
	;
	v773 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	v775 = v773
	goto L166
L155:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	F_ExtendSUBTRANS(m, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	v735 = int32(_a_F_ProcArrayApplyRecoveryInfo_8)
	v736 = int32(3)
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9]))
	v740 = v738 + int32(1)
	if base.Ui32(v740) <= base.Ui32(v736) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v743 = v736
	goto L160
L159:
	;
	v743 = v740
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v743
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v745))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v743)) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	if v757 != 0 {
		goto L155
	} else {
		goto L165
	}
L162:
	;
	v757 = base.B2i32(base.Ui32(v743) < base.Ui32(v745))
	goto L161
L163:
	;
	goto L164
L164:
	;
	v757 = int32(base.Ui32(v743-v745) >> (uint(int32(31)) % 32))
	goto L161
L165:
	;
	goto L156
L166:
	;
	v789 = v775 - int32(1)
	if base.Ui32(v789) < base.Ui32(int32(3)) {
		v775 = v789
		goto L166
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[9])) = v789
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v794 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v821)+48))
	if v823 != 0 {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8])) = v789
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(2)
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+24)) = v789
	goto L169
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8])) = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[2]))
	if v794 == int32(2) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v812)+24)) = v789
	goto L169
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v812)+24)) = int32(0)
	goto L169
L176:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[0]))
	F_LWLockRelease(m, v847+int32(512))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L185
	}
L177:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v819))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v823)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v840 = v821
	goto L179
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v840)+48)) = v822 + base.I64_extend_i32_s(v819-base.I32_wrap_i64(v822))
	goto L176
L180:
	;
	if v835 == int32(0) {
		goto L176
	} else {
		goto L184
	}
L181:
	;
	v835 = base.B2i32(base.Ui32(v823) < base.Ui32(v819))
	goto L180
L182:
	;
	goto L183
L183:
	;
	v835 = int32(base.Ui32(v823-v819) >> (uint(int32(31)) % 32))
	goto L180
L184:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[1]))
	v840 = v839
	goto L179
L185:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v856 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7]))
	v859 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v856 == int32(3) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v859 == int32(0) {
		goto L100
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	if v859 == int32(0) {
		goto L100
	} else {
		goto L193
	}
L191:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_9), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v960 = int32(1304)
	goto L101
L193:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v872
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_10), v17)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v960 = int32(1310)
	goto L101
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_11), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_ProcArrayApplyRecoveryInfo_1), int32(1192), int32(_a_F_ProcArrayApplyRecoveryInfo_6))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
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
	if v914 != 0 {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v914 = base.B2i32(base.Ui32(v901) < base.Ui32(v902))
	goto L199
L201:
	;
	goto L202
L202:
	;
	v914 = int32(base.Ui32(v901-v902) >> (uint(int32(31)) % 32))
	goto L199
L203:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[7])) = int32(3)
	v920 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v931 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L209
	}
L206:
	;
	if v920 == int32(0) {
		goto L100
	} else {
		goto L207
	}
L207:
	;
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_9), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v960 = int32(1125)
	goto L101
L209:
	;
	if v931 == int32(0) {
		goto L100
	} else {
		goto L210
	}
L210:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v937 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayApplyRecoveryInfo[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v937
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v935
	F_errmsg_internal(m, int32(_a_F_ProcArrayApplyRecoveryInfo_10), v17+int32(32))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v960 = int32(1132)
	goto L101
L212:
	;
	goto L100
}
func F_ProcArrayInstallImportedXmin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[1]))
	v17 = F_LWLockAcquire(m, v13+int32(512), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v114 = v3
	goto L3
L3:
	;
	return v114
L4:
	;
	return int32(0)
L5:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v22 <= v21 {
		v105 = v21
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[1]))
	F_LWLockRelease(m, v107+int32(512))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L23
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[2]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[3]))
	v33 = v3
	v36 = v22
	v37 = v28
	v38 = v30
	goto L9
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[4])) = l0
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+40)) = l0
	v105 = int32(1)
	goto L6
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v33))))
	if v42&int32(2) != 0 {
		v83 = v36
		v84 = v37
		v85 = v38
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v105 = int32(0)
	goto L6
L11:
	;
	v87 = v33 + int32(1)
	if v87 < v83 {
		v33 = v87
		v36 = v83
		v37 = v84
		v38 = v85
		goto L9
	} else {
		goto L22
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v33<<(uint(int32(2))%32))))
	v51 = v38 + v48*int32(640)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v52 != v53 {
		v83 = v36
		v84 = v37
		v85 = v38
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v55 != v56 {
		v83 = v36
		v84 = v37
		v85 = v38
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[6]))
	if v58 != v60 {
		v83 = v36
		v84 = v37
		v85 = v38
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
	if base.Ui32(v62) < base.Ui32(int32(3)) {
		v83 = v36
		v84 = v37
		v85 = v38
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v76 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	v76 = base.B2i32(base.Ui32(v62) <= base.Ui32(l0))
	goto L17
L19:
	;
	goto L20
L20:
	;
	v76 = base.B2i32(v62-l0 <= int32(0))
	goto L17
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[3]))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayInstallImportedXmin[2]))
	v83 = v77
	v84 = v81
	v85 = v79
	goto L11
L22:
	;
	goto L10
L23:
	;
	v114 = v105
	goto L3
}
func F_ProcSignalInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
	if int32(0) <= v13 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[1]))
		if v17+int32(38) <= v13 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v127
				v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v130 + int32(38)
				F_errmsg_internal(m, int32(_a_F_ProcSignalInit_0), v10+int32(16))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(176), int32(_a_F_ProcSignalInit_2))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
			v25 = v22 + v13<<(uint(int32(7))%32)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = int32(1)
			if v26 != 0 {
				F_s_lock(m, v25+int32(104), int32(_a_F_ProcSignalInit_1), int32(179), int32(_a_F_ProcSignalInit_2))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = v25 + int32(8)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					v39 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v25)+96)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25)+88)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25)+80)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25-int32(-64)))) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v39
					v55 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v37)+112)) = v55
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
					*(*int64)(unsafe.Add(mBase, uint32(v58))) = v59
					*(*int64)(unsafe.Add(mBase, uint32(v37)+104)) = v59
					if base.B2i32(l1 == v55)|base.B2i32(l1 <= v55) == v55 {
						base.MemoryCopy(m, v25+int32(16), l0, l1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l1
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
					v75 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = v75
					*(*int32)(unsafe.Add(mBase, uint32(v37))) = v74
					if v38 == v75 {
						*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
						F_on_shmem_exit(m, int32(1106), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					} else {
						v82 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							if v82 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
								F_on_shmem_exit(m, int32(1106), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									m.G0 = v10 + int32(32)
									return
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v87
								v90 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v90
								F_errmsg_internal(m, int32(_a_F_ProcSignalInit_3), v10)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(213), int32(_a_F_ProcSignalInit_2))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
										F_on_shmem_exit(m, int32(1106), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v37 = v25 + int32(8)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+96)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25)+88)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25)+80)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25-int32(-64)))) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v39
				v55 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+112)) = v55
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[2]))
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
				*(*int64)(unsafe.Add(mBase, uint32(v58))) = v59
				*(*int64)(unsafe.Add(mBase, uint32(v37)+104)) = v59
				if base.B2i32(l1 == v55)|base.B2i32(l1 <= v55) == v55 {
					base.MemoryCopy(m, v25+int32(16), l0, l1)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l1
				v74 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
				v75 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = v75
				*(*int32)(unsafe.Add(mBase, uint32(v37))) = v74
				if v38 == v75 {
					*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
					F_on_shmem_exit(m, int32(1106), int32(0))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				} else {
					v82 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						if v82 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
							F_on_shmem_exit(m, int32(1106), int32(0))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								m.G0 = v10 + int32(32)
								return
							}
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v87
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v90
							F_errmsg_internal(m, int32(_a_F_ProcSignalInit_3), v10)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(213), int32(_a_F_ProcSignalInit_2))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcSignalInit[4])) = v37
									F_on_shmem_exit(m, int32(1106), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
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
		v112 = m.ExcPending
		if v112 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ProcSignalInit_4), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ProcSignalInit_1), int32(174), int32(_a_F_ProcSignalInit_2))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
