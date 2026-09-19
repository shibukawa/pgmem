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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
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
	var v375 int64
	_ = v375
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v413 int32
	_ = v413
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
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
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
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
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int64
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
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
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 float64
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
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
	v844 = m.ExcPending
	if v844 != 0 {
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
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[0]))
	if v695 != 0 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(0)
	if v611 != 0 {
		goto L103
	} else {
		goto L104
	}
L6:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v586 = v581
	v611 = v582
	goto L5
L7:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L12
	} else {
		goto L101
	}
L8:
	;
	v586 = v44
	v611 = int32(0)
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
	v99 = base.AtomicRmwXchg32(m, v68, int32(0), int32(1))
	if v99 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(1)
	v116 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+4)), uint32(v116))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
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
	v106 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v68))), uint32(v106))
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
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v122 = F_MultiExecProcNode(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v122
	if v122 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v127 != int32(478) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v133 = F_dsa_allocate_extended(m, v130, int32(56), int32(4))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v136 = F_dsa_get_address(m, v135, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	if v138 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+16)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v122)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = v497
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	v501 = F_dsa_get_address(m, v499, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L12
	} else {
		goto L83
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	if v139 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v146 = F_dsa_allocate_extended(m, v140, v139<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v155 = v2
	goto L42
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+104)) = v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v150 = F_dsa_get_address(m, v149, v146)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = int32(0)
	v155 = v150
	goto L42
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v163 = F_dsa_allocate_extended(m, v157, v156<<(uint(int32(2))%32)+int32(4), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L48
	}
L46:
	;
	v172 = v2
	goto L47
L47:
	;
	v173 = int32(-1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	switch v174 - int32(1) {
	case 0:
		goto L52
	case 1:
		goto L53
	default:
		goto L38
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+108)) = v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v167 = F_dsa_get_address(m, v166, v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = int32(0)
	v172 = v167
	goto L47
L50:
	;
	if int32(2) <= v425 {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = int32(0)
	v413 = v383
	v424 = v394
	v425 = v395
	goto L50
L52:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v359 = F_dsa_allocate_extended(m, v356, int32(52), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L75
	}
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	v179 = F_dsa_get_address(m, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	if v182 == int64(0) {
		v228 = v173
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v253 = int32(4)
	v262 = v228
	v267 = int32(0)
	v273 = v181
	v275 = v2
	v276 = v2
	goto L63
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	v189 = int32(0)
	goto L57
L57:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v189*int32(48))+4)))
	if v218 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v228 = v173
	goto L55
L59:
	;
	v228 = v189
	goto L55
L60:
	;
	goto L61
L61:
	;
	v222 = v189 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v222)) < base.Ui64(v182) {
		v189 = v222
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v290 = v262
	v295 = v267
	v298 = v267
	goto L67
L65:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v262 = v328
	v267 = v325
	v273 = v355
	v275 = v353
	v276 = v354
	goto L63
L66:
	;
	if v179 != 0 {
		v383 = v179
		v394 = v275
		v395 = v276
		goto L51
	} else {
		goto L74
	}
L67:
	;
	if v298&int32(1) != 0 {
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v336 = base.I32_div_s(v330-(v179+v253), int32(48))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+5)))
	if v337 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v319 = int32(1)
	v320 = v290 - v319
	v324 = base.B2i32(v318&(v320^v228) == int32(0))
	v325 = v324 | v295
	v328 = v318 & v320
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v330 = v290*int32(48) + v329
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+4)))
	if v331 != v319 {
		v290 = v328
		v295 = v325
		v298 = v324
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172+v253+v275<<(uint(int32(2))%32)))) = v336
	v353 = v275 + int32(1)
	v354 = v276
	goto L65
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155+v253+v276<<(uint(int32(2))%32)))) = v336
	v353 = v275
	v354 = v276 + int32(1)
	goto L65
L74:
	;
	v413 = int32(0)
	v424 = v275
	v425 = v276
	goto L50
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+96)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v363 = F_dsa_get_address(m, v362, v359)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v122)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+44)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v122)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+36)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v122)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+28)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v122)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+20)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v122)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+12)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v122)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+4)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = int32(0)
	v383 = v363
	v394 = v2
	v395 = v2
	goto L51
L77:
	;
	v439 = int32(4)
	F_qsort_arg(m, v155+v439, v425, v439, int32(818), v413+v439)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L12
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v424 < int32(2) {
		goto L38
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v449 = int32(4)
	F_qsort_arg(m, v172+v449, v424, v449, int32(818), v413+v449)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	goto L38
L83:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v122)+104))
	v505 = F_dsa_get_address(m, v503, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	v509 = F_dsa_get_address(m, v507, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	if v501 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v513 = base.AtomicRmwAdd32(m, v501, int32(0), int32(1))
	goto L88
L87:
	;
	goto L88
L88:
	;
	if v505 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v516 = base.AtomicRmwAdd32(m, v505, int32(0), int32(1))
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v509 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v519 = base.AtomicRmwAdd32(m, v509, int32(0), int32(1))
	goto L94
L93:
	;
	goto L94
L94:
	;
	v521 = v136 + int32(28)
	v522 = int32(76)
	*(*uint16)(unsafe.Add(mBase, uint32(v521))) = uint16(v522)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v521)+8)) = int64(-1)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v136)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v133
	v537 = base.AtomicRmwXchg32(m, v40, int32(4), int32(1))
	if v537 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_s_lock(m, v68, int32(_a_F_BitmapHeapNext_1), int32(183), int32(_a_F_BitmapHeapNext_4))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L12
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(2)
	v545 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+4)), uint32(v545))
	F_ConditionVariableBroadcast(m, v66)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v643
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v646 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v614 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v614)
	v617 = F_palloc0(m, int32(16))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L12
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v638 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v638)
	v640 = F_tbm_begin_private_iterate(m, v586)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L12
	} else {
		goto L115
	}
L106:
	;
	v619 = F_dsa_get_address(m, v39, v611)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v619
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	v623 = F_dsa_get_address(m, v39, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+4)) = v623
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	if v626 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v619)+20))
	v628 = F_dsa_get_address(m, v39, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v619)+12))
	if v631 == int32(0) {
		v643 = v617
		goto L102
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+8)) = v628
	goto L111
L113:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v619)+24))
	v635 = F_dsa_get_address(m, v39, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+12)) = v635
	v643 = v617
	goto L102
L115:
	;
	v643 = v640
	goto L102
L116:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v650)+8))
	v652 = int32(0)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v649)+188))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+8))
	v658 = m.T0[v657].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v649, v651, v652, v652, v652, int32(258))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L12
	} else {
		goto L119
	}
L117:
	;
	v662 = v646
	goto L118
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v662)+16)) = v645
	v664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v664)
	goto L4
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v658
	v662 = v658
	goto L118
L120:
	;
	m.G0 = v31 + int32(16)
	return v33
L121:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	m.T0[v832].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L12
	} else {
		goto L152
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L12
	} else {
		goto L149
	}
L123:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BitmapHeapNext[1])))
	if v697&int32(1) == int32(0) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v703 = l0 + int32(156)
	goto L127
L126:
	;
	goto L125
L127:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+188))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)+168))
	v740 = m.T0[v739].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v736, v33, v703, l0+int32(136), l0+int32(128))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L12
	} else {
		goto L129
	}
L128:
	;
	goto L122
L129:
	;
	if v740 == int32(0) {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[2]))
	if v745 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
	if v748 != int32(1) {
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
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v752 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v755)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L12
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v758 = int32(_a_F_BitmapHeapNext_5)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3]))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v761
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v752)+20))
	v766 = m.T0[v765].(func(*base.Module, int32, int32, int32) int32)(m, v752, v34, v31+int32(15))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L12
	} else {
		goto L140
	}
L139:
	;
	goto L120
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[3])) = v759
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	if v766 != 0 {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v773 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v774 = *(*float64)(unsafe.Add(mBase, uint32(v773)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v773)+248)) = base.F64_add(v774, float64(1))
	goto L145
L144:
	;
	goto L145
L145:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	m.T0[v779].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_BitmapHeapNext[0]))
	if v783 == int32(0) {
		goto L127
	} else {
		goto L147
	}
L147:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BitmapHeapNext[1])))
	if v787&int32(1) != 0 {
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
	v825 = m.ExcPending
	if v825 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_7), int32(1944), int32(_a_F_BitmapHeapNext_8))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
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
	v848 = m.ExcPending
	if v848 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_BitmapHeapNext_1), int32(84), int32(_a_F_BitmapHeapNext_2))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
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
