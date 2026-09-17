package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BitmapHeapNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v379 int32
	_ = v379
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v409 int32
	_ = v409
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
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
	var v641 int64
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
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
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 float64
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	v2 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
	if v35 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L12
	} else {
		goto L153
	}
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+172))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v40 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L4
L4:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[0]))
	if v691 != 0 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(0)
	if v607 != 0 {
		goto L103
	} else {
		goto L104
	}
L6:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v582 = v577
	v607 = v578
	goto L5
L7:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L12
	} else {
		goto L101
	}
L8:
	;
	v582 = v44
	v607 = int32(0)
	goto L5
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v44 = F_MultiExecProcNode(m, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v66 = v40 + int32(12)
	v68 = v40 + int32(4)
	goto L21
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v44
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v49 == int32(478) {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_0), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_1), int32(74), int32(_a_F_BitmapHeapNext_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(1)
	if v97 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = int64(4294967296)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L12
	} else {
		goto L32
	}
L23:
	;
	F_s_lock(m, v68, int32(_a_F_BitmapHeapNext_1), int32(426), int32(_a_F_BitmapHeapNext_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(0)
	if v105 != int32(1) {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L22
L30:
	;
	F_ConditionVariableSleep(m, v66, int32(134217766))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v118 = F_MultiExecProcNode(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v118
	if v118 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v123 != int32(478) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v129 = F_dsa_allocate_extended(m, v126, int32(56), int32(4))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v132 = F_dsa_get_address(m, v131, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	if v134 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v481
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v483
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v118)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+20)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	v497 = F_dsa_get_address(m, v495, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L12
	} else {
		goto L83
	}
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	if v135 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v142 = F_dsa_allocate_extended(m, v136, v135<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v151 = v2
	goto L42
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	if v152 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+104)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v146 = F_dsa_get_address(m, v145, v142)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(0)
	v151 = v146
	goto L42
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v159 = F_dsa_allocate_extended(m, v153, v152<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L48
	}
L46:
	;
	v168 = v2
	goto L47
L47:
	;
	v169 = int32(-1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	switch v170 - int32(1) {
	case 0:
		goto L52
	case 1:
		goto L53
	default:
		goto L38
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+108)) = v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v163 = F_dsa_get_address(m, v162, v159)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	v168 = v163
	goto L47
L50:
	;
	if int32(2) <= v421 {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = int32(0)
	v409 = v379
	v420 = v390
	v421 = v391
	goto L50
L52:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v355 = F_dsa_allocate_extended(m, v352, int32(52), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L75
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	v175 = F_dsa_get_address(m, v173, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v177)))
	if v178 == int64(0) {
		v224 = v169
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v249 = int32(4)
	v258 = v224
	v261 = int32(0)
	v269 = v177
	v271 = v2
	v272 = v2
	goto L63
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v185 = int32(0)
	goto L57
L57:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v185*int32(48))+4)))
	if v214 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v224 = v169
	goto L55
L59:
	;
	v224 = v185
	goto L55
L60:
	;
	goto L61
L61:
	;
	v218 = v185 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v218)) < base.Ui64(v178) {
		v185 = v218
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v286 = v258
	v289 = v261
	v294 = v261
	goto L67
L65:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v258 = v324
	v261 = v321
	v269 = v351
	v271 = v349
	v272 = v350
	goto L63
L66:
	;
	if v175 != 0 {
		v379 = v175
		v390 = v271
		v391 = v272
		goto L51
	} else {
		goto L74
	}
L67:
	;
	if v294&int32(1) != 0 {
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v332 = base.I32_div_s(v326-(v175+v249), int32(48))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+5)))
	if v333 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v315 = int32(1)
	v316 = v286 - v315
	v320 = base.B2i32(v314&(v316^v224) == int32(0))
	v321 = v320 | v289
	v324 = v314 & v316
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v269)+20))
	v326 = v286*int32(48) + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+4)))
	if v327 != v315 {
		v286 = v324
		v289 = v321
		v294 = v320
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168+v249+v271<<(uint(int32(2))%32)))) = v332
	v349 = v271 + int32(1)
	v350 = v272
	goto L65
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151+v249+v272<<(uint(int32(2))%32)))) = v332
	v349 = v271
	v350 = v272 + int32(1)
	goto L65
L74:
	;
	v409 = int32(0)
	v420 = v271
	v421 = v272
	goto L50
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+96)) = v355
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v359 = F_dsa_get_address(m, v358, v355)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+44)) = v361
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v118)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+36)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v118)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+28)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v118)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+20)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v118)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+12)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v118)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+4)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = int32(0)
	v379 = v359
	v390 = v2
	v391 = v2
	goto L51
L77:
	;
	v435 = int32(4)
	F_qsort_arg(m, v151+v435, v421, v435, int32(818), v409+v435)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v420 < int32(2) {
		goto L38
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v445 = int32(4)
	F_qsort_arg(m, v168+v445, v420, v445, int32(818), v409+v445)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	goto L38
L83:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v118)+104))
	v501 = F_dsa_get_address(m, v499, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
	v505 = F_dsa_get_address(m, v503, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	if v497 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v507 + int32(1)
	goto L88
L87:
	;
	goto L88
L88:
	;
	if v501 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v511 + int32(1)
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v505 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	*(*int32)(unsafe.Add(mBase, uint32(v505))) = v515 + int32(1)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v520 = v132 + int32(28)
	v521 = int32(76)
	*(*uint16)(unsafe.Add(mBase, uint32(v520))) = uint16(v521)
	*(*int32)(unsafe.Add(mBase, uint32(v520)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v520)+8)) = int64(-1)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v132)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+32)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v129
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1)
	if v534 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_s_lock(m, v68, int32(_a_F_BitmapHeapNext_1), int32(183), int32(_a_F_BitmapHeapNext_4))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L12
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = int64(8589934592)
	F_ConditionVariableBroadcast(m, v66)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L12
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	goto L6
L101:
	;
	goto L6
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v639
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v642 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v610 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v610)
	v613 = F_palloc0(m, int32(16))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L12
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v634 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v634)
	v636 = F_tbm_begin_private_iterate(m, v582)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L12
	} else {
		goto L115
	}
L106:
	;
	v615 = F_dsa_get_address(m, v39, v607)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613))) = v615
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v615)+16))
	v619 = F_dsa_get_address(m, v39, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613)+4)) = v619
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	if v622 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v615)+20))
	v624 = F_dsa_get_address(m, v39, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v615)+12))
	if v627 == int32(0) {
		v639 = v613
		goto L102
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613)+8)) = v624
	goto L111
L113:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v615)+24))
	v631 = F_dsa_get_address(m, v39, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613)+12)) = v631
	v639 = v613
	goto L102
L115:
	;
	v639 = v636
	goto L102
L116:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	v648 = int32(0)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v645)+188))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+8))
	v654 = m.T0[v653].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v645, v647, v648, v648, v648, int32(258))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L12
	} else {
		goto L119
	}
L117:
	;
	v658 = v642
	goto L118
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v658)+16)) = v641
	v660 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v660)
	goto L4
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v654
	v658 = v654
	goto L118
L120:
	;
	m.G0 = v31 + int32(16)
	return v33
L121:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	m.T0[v828].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L12
	} else {
		goto L152
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L12
	} else {
		goto L149
	}
L123:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BitmapHeapNext[1])))
	if v693&int32(1) == int32(0) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v699 = l0 + int32(156)
	goto L127
L126:
	;
	goto L125
L127:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+188))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)+168))
	v736 = m.T0[v735].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v732, v33, v699, l0+int32(136), l0+int32(128))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L12
	} else {
		goto L129
	}
L128:
	;
	goto L122
L129:
	;
	if v736 == int32(0) {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[2]))
	if v741 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v744 != int32(1) {
		goto L120
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v33
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v748 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L12
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v754 = int32(_a_F_BitmapHeapNext_5)
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3]))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v757
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v748)+20))
	v762 = m.T0[v761].(func(*base.Module, int32, int32, int32) int32)(m, v748, v34, v31+int32(15))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L12
	} else {
		goto L140
	}
L139:
	;
	goto L120
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v755
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	if v762 != 0 {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v769 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v769)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v769)+248)) = base.F64_add(v770, float64(1))
	goto L145
L144:
	;
	goto L145
L145:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+12))
	m.T0[v775].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[0]))
	if v779 == int32(0) {
		goto L127
	} else {
		goto L147
	}
L147:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BitmapHeapNext[1])))
	if v783&int32(1) != 0 {
		goto L127
	} else {
		goto L148
	}
L148:
	;
	goto L128
L149:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_6), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_7), int32(1944), int32(_a_F_BitmapHeapNext_8))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	goto L120
L153:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_0), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_1), int32(84), int32(_a_F_BitmapHeapNext_2))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BitmapHeapRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v11 != 0 {
		v12 = int32(_a_F_BitmapHeapRecheck_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0])) = v15
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, v11, v9, v7+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0])) = v13
			v30 = base.B2i32(v20 != int32(0))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			F_MemoryContextReset(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v30
			}
		}
	} else {
		v30 = int32(1)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		F_MemoryContextReset(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v30
		}
	}
}
