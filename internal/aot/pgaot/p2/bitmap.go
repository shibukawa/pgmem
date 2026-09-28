package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
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
	var v373 int64
	_ = v373
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
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
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int64
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
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
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int64
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 float64
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	v2 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v35 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L13
	} else {
		goto L158
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L13
	} else {
		goto L155
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+172))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L5
L5:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v708 = l0 + int32(212)
	v710 = l0 + int32(136)
	v712 = l0 + int32(128)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+188))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+168))
	v716 = m.T0[v715].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v706, v33, v708, v710, v712)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L13
	} else {
		goto L130
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(0)
	if v604 != 0 {
		goto L104
	} else {
		goto L105
	}
L7:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v578 = v574
	v604 = v575
	goto L6
L8:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L13
	} else {
		goto L102
	}
L9:
	;
	v578 = v44
	v604 = int32(0)
	goto L6
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v44 = F_MultiExecProcNode(m, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v66 = v40 + int32(12)
	v68 = v40 + int32(4)
	goto L22
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v44
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v49 == int32(486) {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_0), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_1), int32(113), int32(_a_F_BitmapHeapNext_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v99 = base.AtomicRmwXchg32(m, v68, int32(0), int32(1))
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(1)
	v114 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+4)), uint32(v114))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L33
	}
L24:
	;
	F_s_lock(m, v68, int32(_a_F_BitmapHeapNext_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v103 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v104 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v68))), uint32(v104))
	if v103 != int32(1) {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L23
L31:
	;
	F_ConditionVariableSleep(m, v66, int32(134217766))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v120 = F_MultiExecProcNode(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v120
	if v120 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v125 != int32(486) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v131 = F_dsa_allocate_extended(m, v128, int32(56), int32(4))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v134 = F_dsa_get_address(m, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+32))
	if v136 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v483
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v120)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v120)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v120)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+20)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v120)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v120)+96))
	v499 = F_dsa_get_address(m, v497, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L13
	} else {
		goto L84
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v120)+24))
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v144 = F_dsa_allocate_extended(m, v138, v137<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L44
	}
L42:
	;
	v153 = v2
	goto L43
L43:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	if v154 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+104)) = v144
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v148 = F_dsa_get_address(m, v147, v144)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(0)
	v153 = v148
	goto L43
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v161 = F_dsa_allocate_extended(m, v155, v154<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L49
	}
L47:
	;
	v170 = v2
	goto L48
L48:
	;
	v171 = int32(-1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	switch v172 - int32(1) {
	case 0:
		goto L53
	case 1:
		goto L54
	default:
		goto L39
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+108)) = v161
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v165 = F_dsa_get_address(m, v164, v161)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = int32(0)
	v170 = v165
	goto L48
L51:
	;
	if int32(2) <= v423 {
		goto L78
	} else {
		goto L79
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = int32(0)
	v411 = v381
	v422 = v392
	v423 = v393
	goto L51
L53:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v357 = F_dsa_allocate_extended(m, v354, int32(52), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L13
	} else {
		goto L76
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v120)+96))
	v177 = F_dsa_get_address(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v179)))
	if v180 == int64(0) {
		v225 = v171
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v251 = int32(4)
	v261 = v225
	v265 = int32(0)
	v271 = v179
	v273 = v2
	v274 = v2
	goto L64
L57:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	v188 = int32(0)
	goto L58
L58:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v188*int32(48))+4)))
	if v216 != int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v225 = v171
	goto L56
L60:
	;
	v225 = v188
	goto L56
L61:
	;
	goto L62
L62:
	;
	v220 = v188 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v220)) < base.Ui64(v180) {
		v188 = v220
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	v289 = v261
	v293 = v265
	v296 = v265
	goto L68
L66:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v261 = v326
	v265 = v323
	v271 = v353
	v273 = v351
	v274 = v352
	goto L64
L67:
	;
	if v177 != 0 {
		v381 = v177
		v392 = v273
		v393 = v274
		goto L52
	} else {
		goto L75
	}
L68:
	;
	if v296&int32(1) != 0 {
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v334 = base.I32_div_s(v328-(v177+v251), int32(48))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+5)))
	if v335 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v317 = int32(1)
	v318 = v289 - v317
	v322 = base.B2i32(v316&(v318^v225) == int32(0))
	v323 = v322 | v293
	v326 = v316 & v318
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v328 = v289*int32(48) + v327
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+4)))
	if v329 != v317 {
		v289 = v326
		v293 = v323
		v296 = v322
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170+v251+v273<<(uint(int32(2))%32)))) = v334
	v351 = v273 + int32(1)
	v352 = v274
	goto L66
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153+v251+v274<<(uint(int32(2))%32)))) = v334
	v351 = v273
	v352 = v274 + int32(1)
	goto L66
L75:
	;
	v411 = int32(0)
	v422 = v273
	v423 = v274
	goto L51
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+96)) = v357
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v361 = F_dsa_get_address(m, v360, v357)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v120)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+44)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v120)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+36)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v120)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+28)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v120)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+20)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+12)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v120)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v361)+4)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = int32(0)
	v381 = v361
	v392 = v2
	v393 = v2
	goto L52
L78:
	;
	v437 = int32(4)
	F_qsort_arg(m, v153+v437, v423, v437, int32(866), v411+v437)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L13
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v422 < int32(2) {
		goto L39
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v447 = int32(4)
	F_qsort_arg(m, v170+v447, v422, v447, int32(866), v411+v447)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	goto L39
L84:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v120)+104))
	v503 = F_dsa_get_address(m, v501, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v120)+112))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v120)+108))
	v507 = F_dsa_get_address(m, v505, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	if v499 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v511 = base.AtomicRmwAdd32(m, v499, int32(0), int32(1))
	goto L89
L88:
	;
	goto L89
L89:
	;
	if v503 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v514 = base.AtomicRmwAdd32(m, v503, int32(0), int32(1))
	goto L92
L91:
	;
	goto L92
L92:
	;
	if v507 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v517 = base.AtomicRmwAdd32(m, v507, int32(0), int32(1))
	goto L95
L94:
	;
	goto L95
L95:
	;
	F_LWLockInitialize(m, v134+int32(28), int32(80))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+32)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v131
	v532 = base.AtomicRmwXchg32(m, v40, int32(4), int32(1))
	if v532 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_s_lock(m, v68, int32(_a_F_BitmapHeapNext_3))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L13
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(2)
	v538 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+4)), uint32(v538))
	F_ConditionVariableBroadcast(m, v66)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L13
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L7
L102:
	;
	goto L7
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v636
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v639 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	v607 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v607)
	v610 = F_palloc0(m, int32(16))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L13
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v631 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v631)
	v633 = F_tbm_begin_private_iterate(m, v578)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L13
	} else {
		goto L116
	}
L107:
	;
	v612 = F_dsa_get_address(m, v39, v604)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v612
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+16))
	v616 = F_dsa_get_address(m, v39, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+4)) = v616
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v612)+8))
	if v619 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v612)+20))
	v621 = F_dsa_get_address(m, v39, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L13
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	if v624 == int32(0) {
		v636 = v610
		goto L103
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+8)) = v621
	goto L112
L114:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v612)+24))
	v628 = F_dsa_get_address(m, v39, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+12)) = v628
	v636 = v610
	goto L103
L116:
	;
	v636 = v633
	goto L103
L117:
	;
	v642 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L120
	}
L118:
	;
	v673 = v639
	goto L119
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v673)+16)) = v638
	v676 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)) = uint8(v676)
	goto L5
L120:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[0]))
	if v645 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BitmapHeapNext[1])))
	if v647&int32(1) == int32(0) {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v653)+8))
	v655 = int32(0)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v653)+132))
	if v642 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v665 = int32(1282)
	goto L127
L126:
	;
	v665 = int32(258)
	goto L127
L127:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v652)+188))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+8))
	v669 = m.T0[v668].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v652, v654, v655, v655, v655, v658<<(uint(int32(7))%32)&int32(2048)|v665)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L13
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v669
	v673 = v669
	goto L119
L129:
	;
	m.G0 = v31 + int32(16)
	return v33
L130:
	;
	if v716 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	goto L134
L132:
	;
	goto L133
L133:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	m.T0[v821].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L13
	} else {
		goto L154
	}
L134:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[2]))
	if v747 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L13
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	if v750 != int32(1) {
		goto L129
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v33
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v754 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L13
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v760 = int32(_a_F_BitmapHeapNext_4)
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3]))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v763
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v754)+24))
	v768 = m.T0[v767].(func(*base.Module, int32, int32, int32) int64)(m, v754, v34, v31+int32(15))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L13
	} else {
		goto L145
	}
L144:
	;
	goto L129
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v761
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	if v768 != int64(0) {
		goto L129
	} else {
		goto L147
	}
L147:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v777 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v778 = *(*float64)(unsafe.Add(mBase, uint32(v777)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v777)+432)) = base.F64_add(v778, float64(1))
	goto L150
L149:
	;
	goto L150
L150:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)+12))
	m.T0[v783].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)+188))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)+168))
	v790 = m.T0[v789].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v786, v33, v708, v710, v712)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	if v790 != 0 {
		goto L134
	} else {
		goto L153
	}
L153:
	;
	goto L135
L154:
	;
	goto L129
L155:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_0), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_1), int32(123), int32(_a_F_BitmapHeapNext_2))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errmsg_internal(m, int32(_a_F_BitmapHeapNext_5), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_6), int32(931), int32(_a_F_BitmapHeapNext_7))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BitmapHeapRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v12 != 0 {
		v13 = int32(_a_F_BitmapHeapRecheck_0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0])) = v16
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		v21 = m.T0[v20].(func(*base.Module, int32, int32, int32) int64)(m, v12, v10, v8+int32(15))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapRecheck[0])) = v14
			v31 = base.B2i32(v21 != int64(0))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			F_MemoryContextReset(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v31
			}
		}
	} else {
		v31 = int32(1)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		F_MemoryContextReset(m, v32)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v31
		}
	}
}
