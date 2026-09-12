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
	var v184 int32
	_ = v184
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
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
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int64
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 float64
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
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
	v852 = m.ExcPending
	if v852 != 0 {
		goto L12
	} else {
		goto L155
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
	v703 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v703 != 0 {
		goto L125
	} else {
		goto L126
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(0)
	if v617 != 0 {
		goto L105
	} else {
		goto L106
	}
L6:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v592 = v587
	v617 = v588
	goto L5
L7:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L12
	} else {
		goto L102
	}
L8:
	;
	v592 = v44
	v617 = int32(0)
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
	F_errmsg_internal(m, int32(283741), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(498279), int32(74), int32(233069))
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
	F_s_lock(m, v68, int32(498279), int32(426), int32(355237))
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
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v497
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = v499
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v118)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+20)) = v501
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v503
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	v507 = F_dsa_get_address(m, v505, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L84
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
	if int32(2) <= v428 {
		goto L78
	} else {
		goto L79
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = int32(0)
	v419 = v389
	v428 = v398
	v429 = v399
	goto L50
L52:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v363 = F_dsa_allocate_extended(m, v360, int32(52), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L76
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
		v230 = v169
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v249 = int32(4)
	v257 = v230
	v263 = int32(0)
	v269 = v2
	v270 = v2
	v274 = v177
	goto L63
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v184 = int32(0)
	goto L57
L57:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v184*int32(48))+4)))
	if v214 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v230 = v169
	goto L55
L59:
	;
	v230 = v184
	goto L55
L60:
	;
	goto L61
L61:
	;
	v218 = v184 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v218)) < base.Ui64(v178) {
		v184 = v218
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v285 = v257
	v291 = v263
	v295 = v263
	goto L67
L65:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v257 = v324
	v263 = v321
	v269 = v357
	v270 = v358
	v274 = v359
	goto L63
L66:
	;
	if v175 != 0 {
		v389 = v175
		v398 = v269
		v399 = v270
		goto L51
	} else {
		goto L75
	}
L67:
	;
	if v295&int32(1) != 0 {
		goto L66
	} else {
		goto L69
	}
L68:
	;
	if v326 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v315 = int32(1)
	v316 = v285 - v315
	v320 = base.B2i32(v314&(v316^v230) == int32(0))
	v321 = v320 | v291
	v324 = v314 & v316
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v326 = v285*int32(48) + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+4)))
	if v327 != v315 {
		v285 = v324
		v291 = v321
		v295 = v320
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v334 = base.I32_div_s(v326-(v175+v249), int32(48))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+5)))
	if v335 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168+v249+v270<<(uint(int32(2))%32)))) = v334
	v357 = v269
	v358 = v270 + int32(1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151+v249+v269<<(uint(int32(2))%32)))) = v334
	v357 = v269 + int32(1)
	v358 = v270
	goto L65
L75:
	;
	v419 = int32(0)
	v428 = v269
	v429 = v270
	goto L50
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+96)) = v363
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v367 = F_dsa_get_address(m, v366, v363)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+44)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v118)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+36)) = v371
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v118-int32(-64))))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+28)) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v118)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+20)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v118)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+12)) = v379
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v118)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v367)+4)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = int32(0)
	v389 = v367
	v398 = v2
	v399 = v2
	goto L51
L78:
	;
	v445 = int32(4)
	F_qsort_arg(m, v151+v445, v428, v445, int32(818), v419+v445)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L12
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v429 < int32(2) {
		goto L38
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v455 = int32(4)
	F_qsort_arg(m, v168+v455, v429, v455, int32(818), v419+v455)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L12
	} else {
		goto L83
	}
L83:
	;
	goto L38
L84:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v118)+104))
	v511 = F_dsa_get_address(m, v509, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
	v515 = F_dsa_get_address(m, v513, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	if v507 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v517 + int32(1)
	goto L89
L88:
	;
	goto L89
L89:
	;
	if v511 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	*(*int32)(unsafe.Add(mBase, uint32(v511))) = v521 + int32(1)
	goto L92
L91:
	;
	goto L92
L92:
	;
	if v515 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v525 + int32(1)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v530 = v132 + int32(28)
	v531 = int32(76)
	*(*uint16)(unsafe.Add(mBase, uint32(v530))) = uint16(v531)
	*(*int32)(unsafe.Add(mBase, uint32(v530)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v530)+8)) = int64(-1)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v132)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+32)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v129
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1)
	if v544 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_s_lock(m, v68, int32(498279), int32(183), int32(355203))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L12
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = int64(8589934592)
	F_ConditionVariableBroadcast(m, v66)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L6
L102:
	;
	goto L6
L103:
	;
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v655 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v650
	goto L103
L105:
	;
	v620 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v620)
	v623 = F_palloc0(m, int32(16))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L12
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v645 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v645)
	v647 = F_tbm_begin_private_iterate(m, v592)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L12
	} else {
		goto L117
	}
L108:
	;
	v625 = F_dsa_get_address(m, v39, v617)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = v625
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v625)+16))
	v629 = F_dsa_get_address(m, v39, v628)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L12
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v629
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	if v632 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v625)+20))
	v634 = F_dsa_get_address(m, v39, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L12
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	if v637 == int32(0) {
		v650 = v623
		goto L104
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+8)) = v634
	goto L113
L115:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v625)+24))
	v641 = F_dsa_get_address(m, v39, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+12)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v623
	goto L103
L117:
	;
	v650 = v647
	goto L104
L118:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+8))
	v661 = int32(0)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v658)+188))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+8))
	v667 = m.T0[v666].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v658, v660, v661, v661, v661, int32(258))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L12
	} else {
		goto L121
	}
L119:
	;
	v670 = v655
	goto L120
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v670)+16)) = v654
	v672 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v672)
	goto L4
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v667
	v670 = v667
	goto L120
L122:
	;
	m.G0 = v31 + int32(16)
	return v33
L123:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)+12))
	m.T0[v840].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L12
	} else {
		goto L154
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L12
	} else {
		goto L151
	}
L125:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, _consts[125])))
	if v705&int32(1) == int32(0) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v711 = l0 + int32(156)
	goto L129
L128:
	;
	goto L127
L129:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+188))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+168))
	v748 = m.T0[v747].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v744, v33, v711, l0+int32(136), l0+int32(128))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L12
	} else {
		goto L131
	}
L130:
	;
	goto L124
L131:
	;
	if v748 == int32(0) {
		goto L123
	} else {
		goto L132
	}
L132:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v753 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L12
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	if v756 != int32(1) {
		goto L122
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v33
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v760 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v763)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L12
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v766 = int32(4520560)
	v767 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v769
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v760)+20))
	v774 = m.T0[v773].(func(*base.Module, int32, int32, int32) int32)(m, v760, v34, v31+int32(15))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L12
	} else {
		goto L142
	}
L141:
	;
	goto L122
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v767
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	if v774 != 0 {
		goto L122
	} else {
		goto L144
	}
L144:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v781 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v782 = *(*float64)(unsafe.Add(mBase, uint32(v781)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v781)+248)) = base.F64_add(v782, float64(1))
	goto L147
L146:
	;
	goto L147
L147:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	m.T0[v787].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v791 == int32(0) {
		goto L129
	} else {
		goto L149
	}
L149:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, _consts[125])))
	if v795&int32(1) != 0 {
		goto L129
	} else {
		goto L150
	}
L150:
	;
	goto L130
L151:
	;
	F_errmsg_internal(m, int32(337297), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(327602), int32(1944), int32(384576))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L12
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	goto L122
L155:
	;
	F_errmsg_internal(m, int32(283741), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(498279), int32(84), int32(233069))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L12
	} else {
		goto L157
	}
L157:
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v11 != 0 {
		v12 = int32(4520560)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, v11, v9, v7+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
			v31 = base.B2i32(v20 != int32(0))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			F_MemoryContextReset(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v31
			}
		}
	} else {
		v31 = int32(1)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		F_MemoryContextReset(m, v32)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v31
		}
	}
}
