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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[790]))
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
	v41 = *(*int32)(unsafe.Add(mBase, _consts[792]))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	v44 = v42 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v44
	v47 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	v49 = v47 + int32(37)
	if int32(0) <= v49 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[792]))
	v23 = v18 + v12<<(uint(int32(7))%32) + int32(120)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24 | v25
	v29 = v12 + v25
	v31 = *(*int32)(unsafe.Add(mBase, _consts[790]))
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[792]))
	v61 = v58 + v52<<(uint(int32(7))%32)
	v63 = v61 + int32(8)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 == int32(0) {
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+96)) = int32(1)
	v71 = v61 + int32(104)
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v71, int32(521057), int32(402), int32(232749))
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v63)+96)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+56)) = int32(1)
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
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
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
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v681 int32
	_ = v681
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v761 int32
	_ = v761
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int64
	_ = v826
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v46 = *(*int32)(unsafe.Add(mBase, _consts[142]))
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
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v72)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+56)) = v74 + int64(1)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v80)) == int32(0) {
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
	v72 = v65
	goto L6
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	if v60 == int32(0) {
		v72 = v62
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
	if v92 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v92 = base.B2i32(base.Ui32(v80) < base.Ui32(v19))
	goto L15
L17:
	;
	goto L18
L18:
	;
	v92 = int32(base.Ui32(v80-v19) >> (uint(int32(31)) % 32))
	goto L15
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = int32(0)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v97 = m.G0
	v99 = v97 - int32(16)
	m.G0 = v99
	v102 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v105 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
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
	m.G0 = v99 + int32(16)
	v413 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v413+int32(512))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L76
	}
L24:
	;
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v105 != 0 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	F_errmsg_internal(m, int32(183772), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v102)+12)) = int64(0)
	goto L23
L30:
	;
	F_errfinish(m, int32(514532), int32(5045), int32(352997))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v19
	F_errmsg_internal(m, int32(48039), v99)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	if v131 < v132 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	F_errfinish(m, int32(514532), int32(5051), int32(352997))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = v234
	v247 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	v252 = v250 - v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	if v252 == v253 {
		goto L23
	} else {
		goto L58
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	v137 = v131
	v140 = v135
	v145 = v2
	goto L41
L39:
	;
	goto L40
L40:
	;
	v229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = v229
	v234 = v229
	goto L37
L41:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v140))))
	if v151 != int32(1) {
		v181 = v140
		v182 = v145
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+12)) = v191 - v189
	v195 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	v198 = v131
	goto L54
L43:
	;
	goto L42
L44:
	;
	v185 = v137 + int32(1)
	if v185 != v132 {
		v137 = v185
		v140 = v181
		v145 = v182
		goto L41
	} else {
		goto L53
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	v156 = int32(2)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v137<<(uint(v156)%32))))
	if base.B2i32(base.Ui32(v156) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v159)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v171 != 0 {
		v189 = v145
		goto L43
	} else {
		goto L50
	}
L47:
	;
	v171 = base.B2i32(base.Ui32(v19) <= base.Ui32(v159))
	goto L46
L48:
	;
	goto L49
L49:
	;
	v171 = base.B2i32(int32(0) <= v159-v19)
	goto L46
L50:
	;
	v172 = F_StandbyTransactionIdIsPrepared(m, v159)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v172 != 0 {
		v181 = v175
		v182 = v145
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v175))) = uint8(v177)
	v181 = v175
	v182 = v145 + int32(1)
	goto L44
L53:
	;
	v189 = v182
	goto L43
L54:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v198))))
	if v211 != 0 {
		v234 = v198
		goto L37
	} else {
		goto L56
	}
L55:
	;
	goto L40
L56:
	;
	v213 = v198 + int32(1)
	if v213 != v132 {
		v198 = v213
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v250 <= v251 {
		v360 = v247
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+20)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v249)+16)) = int32(0)
	v379 = m.G0
	v380 = int32(16)
	v381 = v379 - v380
	m.G0 = v381
	F___gettimeofday(m, v381)
	mBase = m.M
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v381)))
	v385 = int64(*(*int32)(unsafe.Add(mBase, uint32(v381)+8)))
	m.G0 = v381 + v380
	goto L75
L60:
	;
	v256 = int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	v261 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v251+v256 != v250 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v269 = v251
	v270 = v247
	v278 = int32(0)
	goto L64
L62:
	;
	v326 = v251
	v327 = v247
	goto L63
L63:
	;
	if v252&v256 == int32(0) {
		v360 = v327
		goto L59
	} else {
		goto L73
	}
L64:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v261))))
	if v283 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v326 = v321
	v327 = v319
	goto L63
L66:
	;
	v286 = int32(2)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v259+v269<<(uint(v286)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v259+v270<<(uint(v286)%32)))) = v292
	v295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v270+v261))) = uint8(v295)
	v299 = v270 + v295
	goto L68
L67:
	;
	v299 = v270
	goto L68
L68:
	;
	v300 = int32(1)
	v301 = v269 + v300
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v301))))
	if v303 == v300 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v306 = int32(2)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v259+v301<<(uint(v306)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v259+v299<<(uint(v306)%32)))) = v312
	v315 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v299+v261))) = uint8(v315)
	v319 = v299 + v315
	goto L71
L70:
	;
	v319 = v299
	goto L71
L71:
	;
	v320 = int32(2)
	v321 = v269 + v320
	v323 = v278 + v320
	if v323 != v252&int32(-2) {
		v269 = v321
		v270 = v319
		v278 = v323
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v261))))
	if v342 != int32(1) {
		v360 = v327
		goto L59
	} else {
		goto L74
	}
L74:
	;
	v345 = int32(2)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v259+v326<<(uint(v345)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v259+v327<<(uint(v345)%32)))) = v351
	v354 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v261))) = uint8(v354)
	v360 = v327 + v354
	goto L59
L75:
	;
	*(*int64)(unsafe.Add(mBase, _consts[776])) = v385 + v384*int64(1000000) - int64(946684800000000)
	goto L23
L76:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v420 = v418
	goto L77
L77:
	;
	v434 = v420 - int32(1)
	if base.Ui32(v434) < base.Ui32(int32(3)) {
		v420 = v434
		goto L77
	} else {
		goto L79
	}
L78:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v434)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v440 = m.G0
	v442 = v440 - int32(32)
	m.G0 = v442
	v447 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	F_hash_seq_init(m, v442+int32(12), v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v452 = F_hash_seq_search(m, v442+int32(12))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v452 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v455 = v452
	goto L86
L84:
	;
	goto L85
L85:
	;
	m.G0 = v442 + int32(32)
	v517 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	switch v517 - int32(2) {
	case 0:
		goto L104
	case 1:
		goto L100
	default:
		goto L103
	}
L86:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v469 = F_StandbyTransactionIdIsPrepared(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L85
L88:
	;
	v496 = F_hash_seq_search(m, v442+int32(12))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L98
	}
L89:
	;
	if v469 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v439))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v471)) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v483 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L92:
	;
	v483 = base.B2i32(base.Ui32(v471) < base.Ui32(v439))
	goto L91
L93:
	;
	goto L94
L94:
	;
	v483 = int32(base.Ui32(v471-v439) >> (uint(int32(31)) % 32))
	goto L91
L95:
	;
	F_StandbyReleaseXidEntryLocks(m, v455)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v492 = F_hash_search(m, v489, v455, int32(2), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L88
L98:
	;
	if v496 != 0 {
		v455 = v496
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
	F_errfinish(m, int32(514532), v964, int32(253861))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L212
	}
L102:
	;
	v905 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v906))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v905)) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L103:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v552 = F_LWLockAcquire(m, v548+int32(512), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L111
	}
L104:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v520 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v523 != 0 {
		goto L102
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v527 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v531 = F_LWLockAcquire(m, v527+int32(512), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v525)+12)) = int64(0)
	v538 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v538+int32(512))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[55])) = int32(1)
	goto L103
L111:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = F_palloc(m, (v554+v555)<<(uint(int32(2))%32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v561+v562 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v886+int32(512))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L195
	}
L114:
	;
	F_pfree(m, v559)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L144
	}
L115:
	;
	v568 = int32(0)
	v572 = v2
	goto L116
L116:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v581+v568<<(uint(int32(2))%32))))
	v586 = F_TransactionIdDidCommit(m, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	if v596 <= int32(0) {
		goto L114
	} else {
		goto L124
	}
L118:
	;
	v598 = v568 + int32(1)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v598 < v599+v600 {
		v568 = v598
		v572 = v596
		goto L116
	} else {
		goto L123
	}
L119:
	;
	if v586 != 0 {
		v596 = v572
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v588 = F_TransactionIdDidAbort(m, v585)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v588 != 0 {
		v596 = v572
		goto L118
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559+v572<<(uint(int32(2))%32)))) = v585
	v596 = v572 + int32(1)
	goto L118
L123:
	;
	goto L117
L124:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+12))
	if v607 != 0 {
		goto L113
	} else {
		goto L125
	}
L125:
	;
	F_pg_qsort(m, v559, v596, int32(4), int32(1102))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v612 = int32(1)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	F_KnownAssignedXidsAdd(m, v613, v613, v612)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v596 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v620 = v612
	goto L131
L129:
	;
	goto L130
L130:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L143
	}
L131:
	;
	v635 = v559 + v620<<(uint(int32(2))%32)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v635-int32(4))))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	if v638 == v639 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L130
L133:
	;
	v663 = v620 + int32(1)
	if v663 != v596 {
		v620 = v663
		goto L131
	} else {
		goto L142
	}
L134:
	;
	v643 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	F_KnownAssignedXidsAdd(m, v639, v639, int32(1))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L141
	}
L137:
	;
	if v643 == int32(0) {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v647
	F_errmsg_internal(m, int32(258523), v17+int32(16))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(514532), int32(1217), int32(253861))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
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
	v698 = int32(4470772)
	v699 = int32(3)
	v701 = *(*int32)(unsafe.Add(mBase, _consts[779]))
	v703 = v701 + int32(1)
	if base.Ui32(v703) <= base.Ui32(v699) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v706 = v699
	goto L147
L146:
	;
	v706 = v703
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, _consts[779])) = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v708))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v706)) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v720 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v720 = base.B2i32(base.Ui32(v706) < base.Ui32(v708))
	goto L148
L150:
	;
	goto L151
L151:
	;
	v720 = int32(base.Ui32(v706-v708) >> (uint(int32(31)) % 32))
	goto L148
L152:
	;
	goto L155
L153:
	;
	goto L154
L154:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _consts[779]))
	v779 = v777
	goto L166
L155:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _consts[779]))
	F_ExtendSUBTRANS(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	v739 = int32(4470772)
	v740 = int32(3)
	v742 = *(*int32)(unsafe.Add(mBase, _consts[779]))
	v744 = v742 + int32(1)
	if base.Ui32(v744) <= base.Ui32(v740) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v747 = v740
	goto L160
L159:
	;
	v747 = v744
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, _consts[779])) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v749))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v747)) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	if v761 != 0 {
		goto L155
	} else {
		goto L165
	}
L162:
	;
	v761 = base.B2i32(base.Ui32(v747) < base.Ui32(v749))
	goto L161
L163:
	;
	goto L164
L164:
	;
	v761 = int32(base.Ui32(v747-v749) >> (uint(int32(31)) % 32))
	goto L161
L165:
	;
	goto L156
L166:
	;
	v793 = v779 - int32(1)
	if base.Ui32(v793) < base.Ui32(int32(3)) {
		v779 = v793
		goto L166
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _consts[779])) = v793
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v798 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v825 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v825)+8))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v825)+48))
	if v827 != 0 {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[778])) = v793
	*(*int32)(unsafe.Add(mBase, _consts[55])) = int32(2)
	v807 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+24)) = v793
	goto L169
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, _consts[55])) = int32(3)
	*(*int32)(unsafe.Add(mBase, _consts[778])) = int32(0)
	v816 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	if v798 == int32(2) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+24)) = v793
	goto L169
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+24)) = int32(0)
	goto L169
L176:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v851+int32(512))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L185
	}
L177:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v823))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v827)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v844 = v825
	goto L179
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v844)+48)) = v826 + base.I64_extend_i32_s(v823-base.I32_wrap_i64(v826))
	goto L176
L180:
	;
	if v839 == int32(0) {
		goto L176
	} else {
		goto L184
	}
L181:
	;
	v839 = base.B2i32(base.Ui32(v827) < base.Ui32(v823))
	goto L180
L182:
	;
	goto L183
L183:
	;
	v839 = int32(base.Ui32(v827-v823) >> (uint(int32(31)) % 32))
	goto L180
L184:
	;
	v843 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v844 = v843
	goto L179
L185:
	;
	F_KnownAssignedXidsDisplay(m, int32(12))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v860 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	v863 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v860 == int32(3) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v863 == int32(0) {
		goto L100
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	if v863 == int32(0) {
		goto L100
	} else {
		goto L193
	}
L191:
	;
	F_errmsg_internal(m, int32(475814), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v964 = int32(1304)
	goto L101
L193:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v878 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v878
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v876
	F_errmsg_internal(m, int32(700520), v17)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v964 = int32(1310)
	goto L101
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errmsg_internal(m, int32(8677), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(514532), int32(1192), int32(253861))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
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
	if v918 != 0 {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v918 = base.B2i32(base.Ui32(v905) < base.Ui32(v906))
	goto L199
L201:
	;
	goto L202
L202:
	;
	v918 = int32(base.Ui32(v905-v906) >> (uint(int32(31)) % 32))
	goto L199
L203:
	;
	*(*int32)(unsafe.Add(mBase, _consts[55])) = int32(3)
	v924 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v935 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L209
	}
L206:
	;
	if v924 == int32(0) {
		goto L100
	} else {
		goto L207
	}
L207:
	;
	F_errmsg_internal(m, int32(475814), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v964 = int32(1125)
	goto L101
L209:
	;
	if v935 == int32(0) {
		goto L100
	} else {
		goto L210
	}
L210:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v941 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v939
	F_errmsg_internal(m, int32(700520), v17+int32(32))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v964 = int32(1132)
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v107 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v33 = v3
	v36 = v22
	v37 = v28
	v38 = v30
	goto L9
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[791])) = l0
	v93 = *(*int32)(unsafe.Add(mBase, _consts[188]))
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
	v60 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v79 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v81 = *(*int32)(unsafe.Add(mBase, _consts[156]))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	if int32(0) <= v15 {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[790]))
		if v19+int32(38) <= v15 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
				return
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, _consts[746]))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v151
				v154 = *(*int32)(unsafe.Add(mBase, _consts[790]))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v154 + int32(38)
				F_errmsg_internal(m, int32(710046), v12+int32(16))
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return
				} else {
					F_errfinish(m, int32(521057), int32(176), int32(107093))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[792]))
			v27 = v24 + v15<<(uint(int32(7))%32)
			v29 = v27 + int32(104)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(1)
			v34 = v27 + int32(8)
			if v30 != 0 {
				F_s_lock(m, v29, int32(521057), int32(179), int32(107093))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v42 = v27 + int32(48)
					if v42&int32(3) == int32(0) {
						v50 = v15<<(uint(int32(7))%32) + v24
						v52 = v50 + int32(104)
						v54 = v50 + int32(52)
						if base.Ui32(v54) < base.Ui32(v52) {
							v56 = v52
						} else {
							v56 = v54
						}
						v65 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), (v56-v50-int32(49))&int32(-4)+int32(4))
						mBase = m.M
					} else {
						v66 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v42))) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+48)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+40)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+32)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v66
					}
					v83 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v83
					v86 = *(*int32)(unsafe.Add(mBase, _consts[792]))
					v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
					*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87
					*(*int64)(unsafe.Add(mBase, uint32(v34)+104)) = v87
					if v83 < l1 {
						if l1 != 0 {
							v94 = F__emscripten_memcpy_bulkmem(m, v27+int32(16), l0, l1)
							mBase = m.M
						} else {
						}
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = l1
					v98 = *(*int32)(unsafe.Add(mBase, _consts[353]))
					v99 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v98
					if v40 == v99 {
						*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
						F_on_shmem_exit(m, int32(1103), int32(0))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v106 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							if v106 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
								F_on_shmem_exit(m, int32(1103), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, _consts[353]))
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
								v114 = *(*int32)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v114
								F_errmsg_internal(m, int32(8737), v12)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_errfinish(m, int32(521057), int32(213), int32(107093))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
										F_on_shmem_exit(m, int32(1103), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return
										} else {
											m.G0 = v12 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v42 = v27 + int32(48)
				if v42&int32(3) == int32(0) {
					v50 = v15<<(uint(int32(7))%32) + v24
					v52 = v50 + int32(104)
					v54 = v50 + int32(52)
					if base.Ui32(v54) < base.Ui32(v52) {
						v56 = v52
					} else {
						v56 = v54
					}
					v65 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), (v56-v50-int32(49))&int32(-4)+int32(4))
					mBase = m.M
				} else {
					v66 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+48)) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+40)) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+32)) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v66
					*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v66
				}
				v83 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v83
				v86 = *(*int32)(unsafe.Add(mBase, _consts[792]))
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
				*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87
				*(*int64)(unsafe.Add(mBase, uint32(v34)+104)) = v87
				if v83 < l1 {
					if l1 != 0 {
						v94 = F__emscripten_memcpy_bulkmem(m, v27+int32(16), l0, l1)
						mBase = m.M
					} else {
					}
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = l1
				v98 = *(*int32)(unsafe.Add(mBase, _consts[353]))
				v99 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v99
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = v98
				if v40 == v99 {
					*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
					F_on_shmem_exit(m, int32(1103), int32(0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						m.G0 = v12 + int32(32)
						return
					}
				} else {
					v106 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						if v106 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
							F_on_shmem_exit(m, int32(1103), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, _consts[353]))
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
							v114 = *(*int32)(unsafe.Add(mBase, _consts[746]))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v114
							F_errmsg_internal(m, int32(8737), v12)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_errfinish(m, int32(521057), int32(213), int32(107093))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[793])) = v34
									F_on_shmem_exit(m, int32(1103), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return
									} else {
										m.G0 = v12 + int32(32)
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
		v136 = m.ExcPending
		if v136 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(113188), int32(0))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return
			} else {
				F_errfinish(m, int32(521057), int32(174), int32(107093))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
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
