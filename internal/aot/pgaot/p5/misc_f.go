package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FindLockCycleRecurseMember(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v241 int32
	_ = v241
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int64
	_ = v670
	var v672 int64
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v709 int32
	_ = v709
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)))
	if v17 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_consts[696])))
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v34 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v258 = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v258 < v260 {
		goto L65
	} else {
		goto L66
	}
L6:
	;
	v38 = v16 + int32(24)
	if v34 == v38 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v42 = l2 + int32(1)
	v53 = v34
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53-int32(16))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+616))
	if v61 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v241 != v38 {
		v53 = v241
		goto L8
	} else {
		goto L61
	}
L11:
	;
	v62 = v61
	goto L13
L12:
	;
	v62 = v60
	goto L13
L13:
	;
	if v62 == l1 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	if v40 <= int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v53-int32(8))))
	v75 = int32(1)
	goto L16
L16:
	;
	v86 = int32(1) << (uint(v75) % 32)
	if v86&v33 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+616))
	if v99 != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v90 = v68 & v86
	goto L20
L19:
	;
	v90 = int32(0)
	goto L20
L20:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v94 = v75 + int32(1)
	if v94 <= v40 {
		v75 = v94
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L17
L24:
	;
	goto L10
L25:
	;
	if v200 != 0 {
		goto L56
	} else {
		goto L57
	}
L26:
	;
	v100 = v99
	goto L28
L27:
	;
	v100 = v60
	goto L28
L28:
	;
	v101 = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	v105 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v101 < v105 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = v101
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v105 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v103+v105<<(uint(int32(2))%32)))) = v100
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v142 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v103+v108<<(uint(int32(2))%32))))
	if v100 == v118 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	if v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v125 = v108 + int32(1)
	if v125 != v105 {
		v108 = v125
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v200 = int32(0)
	goto L25
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[700])) = v42
	v200 = int32(1)
	goto L25
L40:
	;
	goto L33
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v100)+624))
	if v152 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)+92))
	if v145 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v148 = F_FindLockCycleRecurseMember(m, v100, v100, v42, l3, l4)
	mBase = m.M
	if v148 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v200 = int32(1)
	goto L25
L45:
	;
	v200 = int32(0)
	goto L25
L46:
	;
	v156 = v100 + int32(620)
	if v152 == v156 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v158 = v152
	goto L48
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158-int32(624))))
	if v167 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L45
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v183 != v156 {
		v158 = v183
		goto L48
	} else {
		goto L55
	}
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158-int32(536))))
	if v172 == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v176 = v158 - int32(628)
	if v176 == v100 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v178 = F_FindLockCycleRecurseMember(m, v176, v100, v42, l3, l4)
	mBase = m.M
	if v178 == int32(0) {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v200 = int32(1)
	goto L25
L55:
	;
	goto L49
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v205 = v202 + l2*int32(24)
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+8)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+16)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+20)) = v212
	return int32(1)
L57:
	;
	goto L58
L58:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[291]))
	if l0 != v217 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+124)))
	if v219&int32(1) == int32(0) {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[702])) = v60
	goto L10
L61:
	;
	goto L9
L62:
	;
	return v709
L63:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v669 = v666 + l2*int32(24)
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v669))) = v670
	v672 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v669)+8)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+16)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+20)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v679 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v678*v679))) = l1
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v683*v679)+4)) = v377
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v688*v679)+8)) = v16
	v693 = int32(1)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v694 + v693
	v709 = v693
	goto L62
L64:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	if v487 <= int32(0) {
		goto L126
	} else {
		goto L127
	}
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v270 = v258
	goto L68
L66:
	;
	goto L67
L67:
	;
	v304 = v16 + int32(32)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+616))
	if v306 == int32(0) {
		v341 = l0
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v282 = v264 + v270*int32(12)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v283 == v16 {
		goto L64
	} else {
		goto L70
	}
L69:
	;
	goto L67
L70:
	;
	v286 = v270 + int32(1)
	if v286 != v260 {
		v270 = v286
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v348 = int32(0)
	if v305 == v348 {
		v709 = v348
		goto L62
	} else {
		goto L82
	}
L73:
	;
	v309 = int32(0)
	if v305 == v309 {
		v341 = v309
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if v305 == v304 {
		v341 = v309
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v319 = v305
	v321 = v309
	goto L76
L76:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v319)+616))
	if v328 == l1 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v341 = v330
	goto L72
L78:
	;
	v330 = v319
	goto L80
L79:
	;
	v330 = v321
	goto L80
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v331 != v304 {
		v319 = v331
		v321 = v330
		goto L76
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	if v305 == v304 {
		v709 = v348
		goto L62
	} else {
		goto L83
	}
L83:
	;
	v353 = l2 + int32(1)
	v359 = v305
	goto L84
L84:
	;
	if v359 == v341 {
		v709 = v348
		goto L62
	} else {
		goto L86
	}
L85:
	;
	v709 = v348
	goto L62
L86:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v359)+100))
	if int32(base.Ui32(v33)>>(uint(v370)%32))&int32(1) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v485 != v304 {
		v359 = v485
		goto L84
	} else {
		goto L125
	}
L88:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v359)+616))
	if v376 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v377 = v376
	goto L91
L90:
	;
	v377 = v359
	goto L91
L91:
	;
	if v377 == l1 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v359)+616))
	if v382 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v483 != 0 {
		goto L63
	} else {
		goto L124
	}
L94:
	;
	v383 = v382
	goto L96
L95:
	;
	v383 = v359
	goto L96
L96:
	;
	v384 = int32(0)
	v386 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	v388 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v384 < v388 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v391 = v384
	goto L100
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v388 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v386+v388<<(uint(int32(2))%32)))) = v383
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v425 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v386+v391<<(uint(int32(2))%32))))
	if v383 == v401 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L99
L102:
	;
	if v391 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v408 = v391 + int32(1)
	if v408 != v388 {
		v391 = v408
		goto L100
	} else {
		goto L108
	}
L105:
	;
	v483 = int32(0)
	goto L93
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[700])) = v353
	v483 = int32(1)
	goto L93
L108:
	;
	goto L101
L109:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v383)+624))
	if v435 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v383)+92))
	if v428 == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v431 = F_FindLockCycleRecurseMember(m, v383, v383, v353, l3, l4)
	mBase = m.M
	if v431 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v483 = int32(1)
	goto L93
L113:
	;
	v483 = int32(0)
	goto L93
L114:
	;
	v439 = v383 + int32(620)
	if v435 == v439 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v441 = v435
	goto L116
L116:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v441-int32(624))))
	if v450 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L113
L118:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v466 != v439 {
		v441 = v466
		goto L116
	} else {
		goto L123
	}
L119:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v441-int32(536))))
	if v455 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v459 = v441 - int32(628)
	if v459 == v383 {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v461 = F_FindLockCycleRecurseMember(m, v459, v383, v353, l3, l4)
	mBase = m.M
	if v461 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v483 = int32(1)
	goto L93
L123:
	;
	goto L117
L124:
	;
	goto L87
L125:
	;
	goto L85
L126:
	;
	return int32(0)
L127:
	;
	goto L128
L128:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v494 = l2 + int32(1)
	v495 = int32(0)
	v503 = v495
	goto L129
L129:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v492+v503<<(uint(int32(2))%32))))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+616))
	if v516 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v635 = v632 + l2*int32(24)
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v635))) = v636
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v635)+8)) = v638
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v635)+16)) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v635)+20)) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v645 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v644*v645))) = l1
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v649*v645)+4)) = v517
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v654*v645)+8)) = v16
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v660 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v659 + v660
	return v660
L131:
	;
	v517 = v516
	goto L133
L132:
	;
	v517 = v515
	goto L133
L133:
	;
	if v517 == l1 {
		v709 = v495
		goto L62
	} else {
		goto L134
	}
L134:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v515)+100))
	if int32(base.Ui32(v33)>>(uint(v519)%32))&int32(1) != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L130
L136:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v515)+616))
	if v526 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	v629 = v503 + int32(1)
	if v629 != v487 {
		v503 = v629
		goto L129
	} else {
		goto L171
	}
L139:
	;
	if v627 != 0 {
		goto L135
	} else {
		goto L170
	}
L140:
	;
	v527 = v526
	goto L142
L141:
	;
	v527 = v515
	goto L142
L142:
	;
	v528 = int32(0)
	v530 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	v532 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v528 < v532 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v535 = v528
	goto L146
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v532 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v530+v532<<(uint(int32(2))%32)))) = v527
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v569 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L146:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v530+v535<<(uint(int32(2))%32))))
	if v527 == v545 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L145
L148:
	;
	if v535 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v552 = v535 + int32(1)
	if v552 != v532 {
		v535 = v552
		goto L146
	} else {
		goto L154
	}
L151:
	;
	v627 = int32(0)
	goto L139
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[700])) = v494
	v627 = int32(1)
	goto L139
L154:
	;
	goto L147
L155:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v527)+624))
	if v579 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v527)+92))
	if v572 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v575 = F_FindLockCycleRecurseMember(m, v527, v527, v494, l3, l4)
	mBase = m.M
	if v575 == int32(0) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v627 = int32(1)
	goto L139
L159:
	;
	v627 = int32(0)
	goto L139
L160:
	;
	v583 = v527 + int32(620)
	if v579 == v583 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v585 = v579
	goto L162
L162:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v585-int32(624))))
	if v594 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L159
L164:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	if v610 != v583 {
		v585 = v610
		goto L162
	} else {
		goto L169
	}
L165:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v585-int32(536))))
	if v599 == int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v603 = v585 - int32(628)
	if v603 == v527 {
		goto L164
	} else {
		goto L167
	}
L167:
	;
	v605 = F_FindLockCycleRecurseMember(m, v603, v527, v494, l3, l4)
	mBase = m.M
	if v605 == int32(0) {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v627 = int32(1)
	goto L139
L169:
	;
	goto L163
L170:
	;
	goto L138
L171:
	;
	v709 = v495
	goto L62
}
func F_ForceSyncCommit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[105])) = uint8(v2)
	return
}
func F_ForgetBackgroundWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+192)))
	if v17&v6 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v20 + int32(1)
	} else {
	}
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v11*int32(1480)+v6))) = uint8(v24)
	v28 = F_errstart(m, int32(14), v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		if v28 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(737464), v7)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errfinish(m, int32(520757), int32(449), int32(232934))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1484))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
					*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
					F_pfree(m, l0)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1484))
			*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
			F_pfree(m, l0)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_ForgetDatabaseSyncRequests(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(4294967295)
	v14 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+10)) = uint16(v14)
	v20 = F_RegisterSyncRequest(m, v5+int32(8), int32(3), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v5 + int32(32)
		return
	}
}
func F_FreeDesc(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 {
	case 0:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v34 = F_fclose(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = v34
			v37 = int32(4478088)
			v39 = *(*int32)(unsafe.Add(mBase, _consts[644]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[644])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v48
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
			return v36
		}
	case 1:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = F_pgl_pclose(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v36 = v6
			v37 = int32(4478088)
			v39 = *(*int32)(unsafe.Add(mBase, _consts[644]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[644])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v48
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
			return v36
		}
	case 2:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v13 = F_close(m, v12)
		mBase = m.M
		F_emscripten_builtin_free(m, v10)
		mBase = m.M
		v36 = v13
		v37 = int32(4478088)
		v39 = *(*int32)(unsafe.Add(mBase, _consts[644]))
		v41 = v39 - int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[644])) = v41
		v44 = *(*int32)(unsafe.Add(mBase, _consts[645]))
		v47 = v44 + v41*int32(12)
		v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v48
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
		return v36
	case 3:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_pgaio_closing_fd(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = F_close(m, v18)
			mBase = m.M
			v36 = v19
			v37 = int32(4478088)
			v39 = *(*int32)(unsafe.Add(mBase, _consts[644]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[644])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v48
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
			return v36
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(459720), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525903), int32(2829), int32(512538))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_FreeSnapshotBuilder(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 != 0 {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+30)))
		if v7 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(93337), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(525719), int32(344), int32(93899))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
			v12 = v10 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v12
			if v12 == int32(0) {
				F_pfree(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					F_MemoryContextDelete(m, v5)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				F_MemoryContextDelete(m, v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int64)(unsafe.Add(mBase, _consts[650]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v8
	v14 = F_fsm_vacuum_page(m, l0, v5, int32(0), int32(-1), v5+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v456 int32
	_ = v456
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v681 int32
	_ = v681
	var v700 int32
	_ = v700
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v854 int32
	_ = v854
	var v867 int32
	_ = v867
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v931 int32
	_ = v931
	var v965 int32
	_ = v965
	var v977 int32
	_ = v977
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1102 int32
	_ = v1102
	var v1136 int32
	_ = v1136
	var v1161 int32
	_ = v1161
	var v1171 int32
	_ = v1171
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
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
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1471 int32
	_ = v1471
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1546 int32
	_ = v1546
	var v1591 int32
	_ = v1591
	var v1612 int32
	_ = v1612
	var v1623 int32
	_ = v1623
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1648 int32
	_ = v1648
	var v1673 int32
	_ = v1673
	var v1682 int32
	_ = v1682
	v8 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(128)
	m.G0 = v35
	F_DeconstructQualifiedName(m, l0, v35+int32(8), v35+int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v35 + int32(128)
	return v1682
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v54 = int32(0)
	v56 = F_SearchSysCacheList(m, int32(46), int32(1), v53, v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v46 = F_LookupExplicitNamespace(m, v45, l6)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	if v46 != 0 {
		v50 = v46
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v1682 = v8
	goto L3
L10:
	;
	v50 = v8
	goto L4
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	if int32(0) < v58 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v73 = v8
	v84 = v8
	v91 = v8
	goto L15
L13:
	;
	v1648 = v8
	goto L14
L14:
	;
	F_ReleaseCatCacheList(m, v56)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L1
	} else {
		goto L256
	}
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(48)+v91<<(uint(int32(2))%32))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+56))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+22)))
	v103 = v101 + v102
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+104)))
	if v50 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v1648 = v1612
	goto L14
L17:
	;
	v1637 = v91 + int32(1)
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	if v1637 < v1638 {
		v73 = v1612
		v84 = v1623
		v91 = v1637
		goto L15
	} else {
		goto L255
	}
L18:
	;
	v1612 = v73
	v1623 = v1591
	goto L17
L19:
	;
	v200 = v103 + int32(136)
	v202 = v100 + int32(40)
	if l5 == int32(0) {
		v227 = v104
		v228 = v200
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+68))
	if v106 != v50 {
		v1591 = v84
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v109 == int32(0) {
		v1591 = v84
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v173 = int32(0)
	goto L19
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		v1591 = v84
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v115 = int32(0)
	if v115 < v112 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v118 = v112
	goto L28
L27:
	;
	v118 = v115
	goto L28
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v103)+68))
	v123 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	v131 = int32(0)
	goto L29
L29:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v119+v131<<(uint(int32(2))%32))))
	if v121 != v123 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v1591 = v84
	goto L18
L31:
	;
	v163 = base.B2i32(v160 == v121)
	goto L33
L32:
	;
	v163 = int32(0)
	goto L33
L33:
	;
	if v163 != 0 {
		v173 = v131
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v165 = v131 + int32(1)
	if v165 != v118 {
		v131 = v165
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	if l1 < v227 {
		goto L133
	} else {
		goto L134
	}
L37:
	;
	if v227 <= l1 {
		goto L123
	} else {
		goto L124
	}
L38:
	;
	v650 = base.B2i32(v227 <= l1) | (l4 ^ int32(1))
	if v650 != 0 {
		goto L113
	} else {
		goto L114
	}
L39:
	;
	v646 = int32(0)
	v647 = v623
	v648 = v84
	goto L38
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L109
	}
L41:
	;
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v209 = F_SysCacheGetAttr(m, int32(46), v202, int32(21), v35+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	if v211 != 0 {
		v227 = v104
		v228 = v200
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v212 = F_pg_detoast_datum(m, v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v214 != int32(1) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	if v217 < int32(0) {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	if v220 != 0 {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	if v221 != int32(26) {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v227 = v217
	v228 = v212 + int32(24)
	goto L41
L50:
	;
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v623 = int32(0)
	if l3 == v623 {
		goto L39
	} else {
		goto L107
	}
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v103)+88))
	if v229 != 0 {
		v1591 = v84
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v231 = l4 & base.B2i32(l1 < v227)
	if v231 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L55
L57:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v100)+56))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+22)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v243 = F_SysCacheGetAttr(m, int32(47), v202, int32(23), v35+int32(15))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L63
	}
L58:
	;
	v232 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+106)))
	if v227 <= l1+v232 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if l1 != v227 {
		v1591 = v84
		goto L18
	} else {
		goto L62
	}
L61:
	;
	v1591 = v84
	goto L18
L62:
	;
	goto L57
L63:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+15)))
	if v245 != 0 {
		v1591 = v84
		goto L18
	} else {
		goto L64
	}
L64:
	;
	v253 = F_get_func_arg_info(m, v202, v35+int32(124), v35+int32(120), v35+int32(116))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v257 = F_palloc(m, v227<<(uint(int32(2))%32))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v263 = F__emscripten_memset_bulkmem(m, v35+int32(16), base.I32_extend8_s(int32(0)), v227)
	mBase = m.M
	goto L67
L67:
	;
	v264 = l1 - v238
	if v264 <= int32(0) {
		v456 = int32(0)
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v472 <= int32(0) {
		v681 = v456
		goto L37
	} else {
		goto L81
	}
L69:
	;
	v271 = F__emscripten_memset_bulkmem(m, v35+int32(16), base.I32_extend8_s(int32(1)), v264)
	mBase = m.M
	goto L70
L70:
	;
	v273 = v264 & int32(7)
	v274 = int32(0)
	if base.Ui32(v238-l1) <= base.Ui32(int32(-8)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v282 = v274
	v294 = int32(0)
	goto L74
L72:
	;
	v365 = v274
	goto L73
L73:
	;
	if v273 == int32(0) {
		v456 = v264
		goto L68
	} else {
		goto L77
	}
L74:
	;
	v314 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v282<<(uint(v314)%32)))) = v282
	v319 = v282 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v319<<(uint(v314)%32)))) = v319
	v325 = v282 | v314
	*(*int32)(unsafe.Add(mBase, uint32(v257+v325<<(uint(v314)%32)))) = v325
	v331 = v282 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v331<<(uint(v314)%32)))) = v331
	v337 = v282 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v337<<(uint(v314)%32)))) = v337
	v343 = v282 | int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v343<<(uint(v314)%32)))) = v343
	v349 = v282 | int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v349<<(uint(v314)%32)))) = v349
	v355 = v282 | int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v355<<(uint(v314)%32)))) = v355
	v360 = int32(8)
	v361 = v282 + v360
	v363 = v294 + v360
	if v363 != v264&int32(2147483640) {
		v282 = v361
		v294 = v363
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v365 = v361
	goto L73
L76:
	;
	goto L75
L77:
	;
	v399 = v365
	v410 = v274
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257+v399<<(uint(int32(2))%32)))) = v399
	v435 = int32(1)
	v438 = v410 + v435
	if v438 != v273 {
		v399 = v399 + v435
		v410 = v438
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v456 = v264
	goto L68
L80:
	;
	goto L79
L81:
	;
	v475 = int32(0)
	if v253 <= v475 {
		v1591 = v84
		goto L18
	} else {
		goto L82
	}
L82:
	;
	v494 = v456
	v507 = v475
	goto L83
L83:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v35)+116))
	v511 = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v514+v507<<(uint(int32(2))%32))))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v35)+120))
	v522 = v511
	v533 = v511
	goto L85
L84:
	;
	v1591 = v84
	goto L18
L85:
	;
	if l5|base.B2i32(v510 == v511) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L84
L87:
	;
	v621 = v522 + int32(1)
	if v621 != v253 {
		v522 = v621
		v533 = v618
		goto L85
	} else {
		goto L106
	}
L88:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522+v510))))
	v559 = v557 - int32(98)
	if base.Ui32(int32(20)) < base.Ui32(v559) {
		v618 = v533
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v519+v522<<(uint(int32(2))%32))))
	if v572 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if int32(1)<<(uint(v559)%32)&int32(1048705) == int32(0) {
		v618 = v533
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v618 = v533 + int32(1)
	goto L87
L94:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v578 == int32(0) {
		v597 = v577
		v598 = v578
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v598-v597 != 0 {
		goto L93
	} else {
		goto L103
	}
L96:
	;
	goto L95
L97:
	;
	if v577 != v578 {
		v597 = v577
		v598 = v578
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v582 = v572
	v583 = v518
	goto L99
L99:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+1)))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+1)))
	if v587 == int32(0) {
		v597 = v586
		v598 = v587
		goto L96
	} else {
		goto L101
	}
L100:
	;
	v597 = v586
	v598 = v587
	goto L96
L101:
	;
	v590 = int32(1)
	if v586 == v587 {
		v582 = v582 + v590
		v583 = v583 + v590
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v602 = v35 + int32(16) + v533
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	if v603 != 0 {
		v1591 = v84
		goto L18
	} else {
		goto L104
	}
L104:
	;
	v604 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v604)
	*(*int32)(unsafe.Add(mBase, uint32(v257+v494<<(uint(int32(2))%32)))) = v533
	v611 = v494 + v604
	v613 = v507 + v604
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v613 < v614 {
		v494 = v611
		v507 = v613
		goto L83
	} else {
		goto L105
	}
L105:
	;
	v681 = v611
	goto L37
L106:
	;
	goto L86
L107:
	;
	if l1 < v227 {
		goto L39
	} else {
		goto L108
	}
L108:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v103)+88))
	v629 = base.B2i32(v627 != int32(0))
	v646 = v627
	v647 = v629
	v648 = v84 | v629
	goto L38
L109:
	;
	F_errmsg_internal(m, int32(162159), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(525427), int32(1289), int32(170933))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
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
	v657 = v650 ^ int32(1)
	v658 = int32(0)
	if l1 < v658 {
		v786 = v646
		v793 = v658
		v804 = v647
		v805 = v655
		v813 = v657
		goto L36
	} else {
		goto L117
	}
L113:
	;
	v655 = v648
	goto L112
L114:
	;
	goto L115
L115:
	;
	v652 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+106)))
	if v227 <= l1+v652 {
		v655 = int32(1)
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v1591 = v648
	goto L18
L117:
	;
	if l1 == v227 {
		v786 = v646
		v793 = v658
		v804 = v647
		v805 = v655
		v813 = v657
		goto L36
	} else {
		goto L118
	}
L118:
	;
	if v647 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v786 = v646
	v793 = v658
	v804 = int32(1)
	v805 = v655
	v813 = v657
	goto L36
L120:
	;
	goto L121
L121:
	;
	if v650 != 0 {
		v1591 = v655
		goto L18
	} else {
		goto L122
	}
L122:
	;
	v786 = v646
	v793 = v658
	v804 = int32(0)
	v805 = v655
	v813 = int32(1)
	goto L36
L123:
	;
	v784 = int32(0)
	v786 = v784
	v793 = v257
	v804 = v784
	v805 = int32(1)
	v813 = v231
	goto L36
L124:
	;
	if v227 <= v264 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v700 = int32(*(*int16)(unsafe.Add(mBase, uint32(v236+v237)+106)))
	v716 = v264
	v718 = v681
	goto L126
L126:
	;
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(16)+v716))))
	if v737 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L123
L128:
	;
	if v716 < v227-v700 {
		v1591 = v84
		goto L18
	} else {
		goto L131
	}
L129:
	;
	v747 = v718
	goto L130
L130:
	;
	v749 = v716 + int32(1)
	if v749 != v227 {
		v716 = v749
		v718 = v747
		goto L126
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257+v718<<(uint(int32(2))%32)))) = v716
	v747 = v718 + int32(1)
	goto L130
L132:
	;
	goto L127
L133:
	;
	v819 = v227
	goto L135
L134:
	;
	v819 = l1
	goto L135
L135:
	;
	v821 = v819 << (uint(int32(2)) % 32)
	v824 = F_palloc(m, v821+int32(32))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v173
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v824)+28)) = v793
	*(*int32)(unsafe.Add(mBase, uint32(v824)+16)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v824)+12)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v824)+8)) = v827
	if v793 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v804 != 0 {
		goto L157
	} else {
		goto L158
	}
L138:
	;
	if v227 <= int32(0) {
		goto L137
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v1006 = v227 << (uint(int32(2)) % 32)
	if v1006 != 0 {
		goto L153
	} else {
		goto L154
	}
L141:
	;
	v835 = v227 & int32(3)
	v837 = v824 + int32(32)
	v838 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v227) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v854 = v838
	v867 = int32(0)
	goto L145
L143:
	;
	v931 = v838
	goto L144
L144:
	;
	if v835 == int32(0) {
		goto L137
	} else {
		goto L148
	}
L145:
	;
	v877 = int32(2)
	v878 = v854 << (uint(v877) % 32)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v793+v878)))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v228+v881<<(uint(v877)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v837+v878))) = v885
	v887 = int32(4)
	v888 = v878 | v887
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v793+v888)))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v228+v891<<(uint(v877)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v837+v888))) = v895
	v898 = v878 | int32(8)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v793+v898)))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v228+v901<<(uint(v877)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v837+v898))) = v905
	v908 = v878 | int32(12)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v793+v908)))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v228+v911<<(uint(v877)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v837+v908))) = v915
	v918 = v854 + v887
	v920 = v867 + v887
	if v920 != v227&int32(2147483644) {
		v854 = v918
		v867 = v920
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v931 = v918
	goto L144
L147:
	;
	goto L146
L148:
	;
	v965 = v931
	v977 = v838
	goto L149
L149:
	;
	v988 = int32(2)
	v989 = v965 << (uint(v988) % 32)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v793+v989)))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v228+v992<<(uint(v988)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v837+v989))) = v996
	v998 = int32(1)
	v1001 = v977 + v998
	if v1001 != v835 {
		v965 = v965 + v998
		v977 = v1001
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L137
L151:
	;
	goto L150
L152:
	;
	goto L137
L153:
	;
	v1007 = F__emscripten_memcpy_bulkmem(m, v824+int32(32), v228, v1006)
	mBase = m.M
	goto L155
L154:
	;
	goto L155
L155:
	;
	goto L152
L156:
	;
	v1207 = int32(0)
	if v813 != 0 {
		goto L170
	} else {
		goto L171
	}
L157:
	;
	v1041 = v819 - v227
	v1042 = int32(1)
	v1043 = v1041 + v1042
	*(*int32)(unsafe.Add(mBase, uint32(v824)+20)) = v1043
	v1046 = v824 + int32(32)
	v1048 = v227 - v1042
	v1051 = v1043 & int32(7)
	if v1051 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824)+20)) = int32(0)
	goto L156
L160:
	;
	v1059 = int32(0)
	v1061 = v1048
	goto L163
L161:
	;
	v1102 = v1048
	goto L162
L162:
	;
	if base.Ui32(v1041) < base.Ui32(int32(7)) {
		goto L156
	} else {
		goto L166
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1046+v1061<<(uint(int32(2))%32)))) = v786
	v1088 = int32(1)
	v1089 = v1061 + v1088
	v1091 = v1059 + v1088
	if v1091 != v1051 {
		v1059 = v1091
		v1061 = v1089
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v1102 = v1089
	goto L162
L165:
	;
	goto L164
L166:
	;
	v1136 = v1102
	goto L167
L167:
	;
	v1161 = v1046 + v1136<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1161))) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+4)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+8)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+12)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+16)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+20)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+24)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v1161)+28)) = v786
	v1171 = v1136 + int32(8)
	if v1171 != v819 {
		v1136 = v1171
		goto L167
	} else {
		goto L169
	}
L168:
	;
	goto L156
L169:
	;
	goto L168
L170:
	;
	v1210 = v227 - l1
	goto L172
L171:
	;
	v1210 = v1207
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824)+24)) = v1210
	if v73 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if (v805^int32(1))&base.B2i32(v50 != int32(0)) != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v1546 = v1207
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v1546
	v1612 = v824
	v1623 = v805
	goto L17
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v73
	v1612 = v824
	v1623 = v805
	goto L17
L177:
	;
	goto L178
L178:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+37)))
	if (v1218^int32(-1)|v805)&int32(1) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+4))
	if v1440 == v173 {
		goto L235
	} else {
		goto L236
	}
L180:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v1226 != v819 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v1299 = v824 + int32(32)
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v824)+16))
	v1301 = v1300 - v1210
	v1303 = v1301 << (uint(int32(2)) % 32)
	v1304 = v73
	goto L205
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v73
	v1612 = v824
	v1623 = v805
	goto L17
L184:
	;
	goto L185
L185:
	;
	v1229 = int32(32)
	v1230 = v824 + v1229
	v1232 = v73 + v1229
	if base.Ui32(int32(4)) <= base.Ui32(v821) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	if v1294 == int32(0) {
		v1408 = v73
		goto L179
	} else {
		goto L204
	}
L187:
	;
	v1294 = int32(0)
	goto L186
L188:
	;
	v1268 = v1263
	v1269 = v1264
	v1270 = v1265
	goto L198
L189:
	;
	if (v1230|v1232)&int32(3) != 0 {
		v1263 = v1230
		v1264 = v1232
		v1265 = v821
		goto L188
	} else {
		goto L192
	}
L190:
	;
	v1256 = v1230
	v1257 = v1232
	v1258 = v821
	goto L191
L191:
	;
	if v1258 == int32(0) {
		goto L187
	} else {
		goto L197
	}
L192:
	;
	v1240 = v1230
	v1241 = v1232
	v1242 = v821
	goto L193
L193:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1241)))
	if v1245 != v1246 {
		v1263 = v1240
		v1264 = v1241
		v1265 = v1242
		goto L188
	} else {
		goto L195
	}
L194:
	;
	v1256 = v1251
	v1257 = v1249
	v1258 = v1253
	goto L191
L195:
	;
	v1248 = int32(4)
	v1249 = v1241 + v1248
	v1251 = v1240 + v1248
	v1253 = v1242 - v1248
	if base.Ui32(int32(3)) < base.Ui32(v1253) {
		v1240 = v1251
		v1241 = v1249
		v1242 = v1253
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v1263 = v1256
	v1264 = v1257
	v1265 = v1258
	goto L188
L198:
	;
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268))))
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269))))
	if v1273 == v1274 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v1294 = v1273 - v1274
	goto L186
L200:
	;
	v1276 = int32(1)
	v1281 = v1270 - v1276
	if v1281 != 0 {
		v1268 = v1268 + v1276
		v1269 = v1269 + v1276
		v1270 = v1281
		goto L198
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	goto L199
L203:
	;
	goto L187
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v73
	v1612 = v824
	v1623 = v805
	goto L17
L205:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+16))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+24))
	if v1336-v1337 == v1301 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v73
	v1612 = v824
	v1623 = v805
	goto L17
L207:
	;
	v1341 = v1304 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v1303) {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	goto L209
L209:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1304)))
	if v1406 != 0 {
		v1304 = v1406
		goto L205
	} else {
		goto L229
	}
L210:
	;
	if v1403 == int32(0) {
		v1408 = v1304
		goto L179
	} else {
		goto L228
	}
L211:
	;
	v1403 = int32(0)
	goto L210
L212:
	;
	v1377 = v1372
	v1378 = v1373
	v1379 = v1374
	goto L222
L213:
	;
	if (v1299|v1341)&int32(3) != 0 {
		v1372 = v1299
		v1373 = v1341
		v1374 = v1303
		goto L212
	} else {
		goto L216
	}
L214:
	;
	v1365 = v1299
	v1366 = v1341
	v1367 = v1303
	goto L215
L215:
	;
	if v1367 == int32(0) {
		goto L211
	} else {
		goto L221
	}
L216:
	;
	v1349 = v1299
	v1350 = v1341
	v1351 = v1303
	goto L217
L217:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	if v1354 != v1355 {
		v1372 = v1349
		v1373 = v1350
		v1374 = v1351
		goto L212
	} else {
		goto L219
	}
L218:
	;
	v1365 = v1360
	v1366 = v1358
	v1367 = v1362
	goto L215
L219:
	;
	v1357 = int32(4)
	v1358 = v1350 + v1357
	v1360 = v1349 + v1357
	v1362 = v1351 - v1357
	if base.Ui32(int32(3)) < base.Ui32(v1362) {
		v1349 = v1360
		v1350 = v1358
		v1351 = v1362
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v1372 = v1365
	v1373 = v1366
	v1374 = v1367
	goto L212
L222:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377))))
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	if v1382 == v1383 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1403 = v1382 - v1383
	goto L210
L224:
	;
	v1385 = int32(1)
	v1390 = v1379 - v1385
	if v1390 != 0 {
		v1377 = v1377 + v1385
		v1378 = v1378 + v1385
		v1379 = v1390
		goto L222
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	goto L223
L227:
	;
	goto L211
L228:
	;
	goto L209
L229:
	;
	goto L206
L230:
	;
	if v73 == v1408 {
		goto L246
	} else {
		goto L247
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1408)+8)) = int32(0)
	F_pfree(m, v824)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L1
	} else {
		goto L244
	}
L232:
	;
	if int32(0) < v1442 {
		goto L230
	} else {
		goto L243
	}
L233:
	;
	if int32(0) <= v1447 {
		goto L231
	} else {
		goto L242
	}
L234:
	;
	F_pfree(m, v824)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L241
	}
L235:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+20))
	if v804 == int32(0) {
		goto L232
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1447 = v173 - v1440
	if v1447 <= int32(0) {
		goto L233
	} else {
		goto L240
	}
L238:
	;
	if v1442 == int32(0) {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	goto L231
L240:
	;
	goto L234
L241:
	;
	v1591 = v805
	goto L18
L242:
	;
	goto L230
L243:
	;
	goto L231
L244:
	;
	v1591 = v805
	goto L18
L245:
	;
	F_pfree(m, v1408)
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L254
	}
L246:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1408)))
	v1536 = v1464
	goto L245
L247:
	;
	goto L248
L248:
	;
	v1471 = v73
	goto L250
L249:
	;
	v1536 = v73
	goto L245
L250:
	;
	if v1471 == int32(0) {
		goto L249
	} else {
		goto L252
	}
L251:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1408)))
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = v1501
	goto L249
L252:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1471)))
	if v1408 != v1499 {
		v1471 = v1499
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	v1546 = v1536
	goto L175
L255:
	;
	goto L16
L256:
	;
	v1682 = v1648
	goto L3
}
func F___fseeko(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v4 < int32(0) {
		v7 = F___fseeko_unlocked(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v12 = F___fseeko_unlocked(m, l0, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	if l0 < int32(0) {
		if l0 != int32(-100) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if base.B2i32(l3 == int32(0))&base.B2i32(v16 == int32(47)) != 0 {
				v27 = m.Env.X__syscall_stat64(m, l1, l2)
				mBase = m.M
				v33 = v27
			} else {
				if l3 != int32(256) {
					v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v33 = v29
				} else {
					if v16 != int32(47) {
						v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v29
					} else {
						v31 = m.Env.X__syscall_lstat64(m, l1, l2)
						mBase = m.M
						v33 = v31
					}
				}
			}
		} else {
			if l3 == int32(256) {
				v31 = m.Env.X__syscall_lstat64(m, l1, l2)
				mBase = m.M
				v33 = v31
			} else {
				if l3 != 0 {
					v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v33 = v29
				} else {
					v27 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v33 = v27
				}
			}
		}
	} else {
		if l3 != int32(4096) {
			if l0 != int32(-100) {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if base.B2i32(l3 == int32(0))&base.B2i32(v16 == int32(47)) != 0 {
					v27 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v33 = v27
				} else {
					if l3 != int32(256) {
						v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v29
					} else {
						if v16 != int32(47) {
							v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
							mBase = m.M
							v33 = v29
						} else {
							v31 = m.Env.X__syscall_lstat64(m, l1, l2)
							mBase = m.M
							v33 = v31
						}
					}
				}
			} else {
				if l3 == int32(256) {
					v31 = m.Env.X__syscall_lstat64(m, l1, l2)
					mBase = m.M
					v33 = v31
				} else {
					if l3 != 0 {
						v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v29
					} else {
						v27 = m.Env.X__syscall_stat64(m, l1, l2)
						mBase = m.M
						v33 = v27
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v10 != 0 {
				v29 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
				mBase = m.M
				v33 = v29
			} else {
				v11 = m.Env.X__syscall_fstat64(m, l0, l2)
				mBase = m.M
				v33 = v11
			}
		}
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v33) {
		*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0) - v33
		v41 = int32(-1)
	} else {
		v41 = v33
	}
	return v41
}
func F_fastgetattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = int32(1)
	v15 = l1 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v17&v14 == v5 {
		v24 = l2 + v15<<(uint(int32(4))%32)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
		if v25 < int32(0) {
			v72 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v78 = v72
				m.G0 = v10 + int32(16)
				return v78
			}
		} else {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v30 = v16 + v28 + v25
			v32 = v24 + int32(20)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)))
			if v33 != int32(1) {
				v78 = v30
				m.G0 = v10 + int32(16)
				return v78
			} else {
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
				switch v36&int32(65535) - int32(1) {
				case 0:
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30))))
					v78 = v41
					m.G0 = v10 + int32(16)
					return v78
				case 1:
					v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30))))
					v78 = v42
					m.G0 = v10 + int32(16)
					return v78
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v36
						F_errmsg_internal(m, int32(507173), v10)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(343885), int32(70), int32(73857))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					v78 = v43
					m.G0 = v10 + int32(16)
					return v78
				}
			}
		}
	} else {
		v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+23)))
		if int32(base.Ui32(v63)>>(uint(v15)%32))&int32(1) != 0 {
			v72 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v78 = v72
				m.G0 = v10 + int32(16)
				return v78
			}
		} else {
			v67 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v67)
			v78 = int32(0)
			m.G0 = v10 + int32(16)
			return v78
		}
	}
}
func F_fdw_handler_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(230880)
			F_errmsg(m, int32(203466), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519257), int32(369), int32(72907))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_finalize_primnode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(8) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L4
	case 15:
		goto L3
	}
L3:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104<<(uint(int32(2))%32)-int32(4))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_finalize_primnode(m, v111, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L28
	}
L4:
	;
	v97 = F_expression_tree_walker_impl(m, l0, int32(848), l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+276))
	if v26 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_bms_add_member(m, v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18
	return int32(0)
L10:
	;
	if v81 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v81 = v70
	goto L10
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v29 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 != int32(1) {
		v70 = v3
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v35 <= int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(0)
	v45 = v35
	goto L17
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v70 = v66
	goto L11
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v43<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 == v54 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v81 = int32(0)
	goto L10
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v58 = F_equal(m, v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v61 = v45
	goto L21
L21:
	;
	v63 = v43 + int32(1)
	if v63 < v61 {
		v43 = v63
		v45 = v61
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v58 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v61 = v60
	goto L21
L24:
	;
	goto L18
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v86 = F_bms_add_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v86
	goto L4
L27:
	;
	return v97
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = F_finalize_primnode(m, v147, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L36
	}
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v122 = v3
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v122<<(uint(int32(2))%32))))
	v133 = F_bms_del_member(m, v127, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v133
	v137 = v122 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v137 < v138 {
		v122 = v137
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)+64))
	v151 = F_bms_copy(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v153 == int32(0) {
		v181 = v151
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v186 = F_bms_join(m, v185, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L45
	}
L39:
	;
	v156 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v157 <= v156 {
		v181 = v151
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v162 = v156
	v163 = v151
	goto L41
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v162<<(uint(int32(2))%32))))
	v172 = F_bms_del_member(m, v163, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	v181 = v172
	goto L38
L43:
	;
	v175 = v162 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v175 < v176 {
		v162 = v175
		v163 = v172
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v186
	goto L1
}
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	v4 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = l2
	v28 = v4
	v30 = v4
	v31 = v4
	v33 = v4
	goto L1
L1:
	;
	v44 = v27
	v47 = v30
	v48 = v31
	v50 = v33
	goto L3
L2:
	;
	v160 = v131
	goto L28
L3:
	;
	v62 = (v44-v47)>>(uint(int32(1))%32) + v47
	v65 = l1 + v62*int32(20)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v48 < v50 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v149 = base.B2i32(v128 != v131) & base.B2i32(v131 <= int32(0))
	if v149 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if int32(1) < v128-v131 {
		v44 = v128
		v47 = v131
		v48 = v132
		v50 = v134
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v128 = v44
	v131 = v62
	v132 = v112
	v134 = v50
	goto L5
L7:
	;
	v68 = v48
	goto L9
L8:
	;
	v68 = v50
	goto L9
L9:
	;
	v71 = v66 + (v68 ^ int32(-1))
	if v71 < int32(0) {
		v112 = v68
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v77 = v68
	v78 = v71
	goto L13
L11:
	;
	v128 = v62
	v131 = v47
	v132 = v48
	v134 = v108
	goto L5
L12:
	;
	if int32(0) <= v98 {
		v112 = v77
		goto L6
	} else {
		goto L18
	}
L13:
	;
	if v19 == v18-v77 {
		v108 = v18 - v19
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v128 = v44
	v131 = v62
	v132 = v66
	v134 = v50
	goto L5
L15:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v18-int32(1)-v77))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v78))))
	v98 = v94 - v97
	if v98 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v99 = int32(1)
	if int32(0) < v78 {
		v77 = v77 + v99
		v78 = v78 - v99
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v108 = v77
	goto L11
L19:
	;
	goto L4
L20:
	;
	if v149 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L2
L23:
	;
	v151 = int32(1)
	goto L25
L24:
	;
	v151 = v28
	goto L25
L25:
	;
	if v28 == int32(0) {
		v27 = v128
		v28 = v151
		v30 = v131
		v31 = v132
		v33 = v134
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	return v196
L28:
	;
	v174 = l1 + v160*int32(20)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 <= v132 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v196 = v195
	goto L27
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v175
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	if v179 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	if v190 <= v191 {
		v160 = v191
		goto L28
	} else {
		goto L38
	}
L34:
	;
	v182 = m.T0[v179].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v186
	if v182 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v196 = v190
	goto L27
}
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = F_getBaseType(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = v5
	goto L3
L3:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = v14
	goto L3
L6:
	;
	v19 = F_getBaseType(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v21 = v5
	goto L8
L8:
	;
	if v18 != v21 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v21 = v19
	goto L8
L10:
	;
	v25 = F_SearchSysCache2(m, int32(12), v18, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v141 = int32(2)
	goto L12
L12:
	;
	m.G0 = v10 + int32(32)
	return v141
L13:
	;
	if v132 != 0 {
		goto L51
	} else {
		goto L52
	}
L14:
	;
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+16)))
	switch v31 - int32(97) {
	case 0:
		v49 = int32(1)
		goto L19
	default:
		goto L21
	case 4:
		goto L20
	case 8:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v21&int32(-9) == int32(22) {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+17)))
	switch v56 - int32(98) {
	case 0:
		goto L30
	default:
		goto L29
	case 4:
		goto L28
	case 7:
		v81 = int32(4)
		goto L27
	}
L19:
	;
	if base.Ui32(v49) <= base.Ui32(l2) {
		goto L18
	} else {
		goto L25
	}
L20:
	;
	v49 = int32(3)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	F_errmsg_internal(m, int32(504563), v10)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(525398), int32(3196), int32(24567))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v132 = int32(0)
	goto L13
L27:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L35
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v78
	v81 = int32(1)
	goto L27
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v132 = int32(2)
	goto L13
L32:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+17)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v66
	F_errmsg_internal(m, int32(510917), v10+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(525398), int32(3218), int32(24567))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v132 = v81
	goto L13
L36:
	;
	v105 = int32(0)
	if l2 == v105 {
		v132 = v105
		goto L13
	} else {
		goto L44
	}
L37:
	;
	v88 = F_get_element_type(m, v21)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v88 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v92 = F_get_element_type(m, v18)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v92 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v98 = F_find_coercion_pathway(m, v88, v92, l2, v10+int32(24))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v98 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v132 = int32(3)
	goto L13
L44:
	;
	F_get_type_category_preferred(m, v21, v10+int32(24), v10+int32(31))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v114 == int32(83) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v132 = int32(4)
	goto L13
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(l2) < base.Ui32(int32(3)) {
		v132 = v105
		goto L13
	} else {
		goto L49
	}
L49:
	;
	F_get_type_category_preferred(m, v18, v10+int32(24), v10+int32(31))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	v132 = base.B2i32(v126 == int32(83)) << (uint(int32(2)) % 32)
	goto L13
L51:
	;
	v136 = v132
	goto L53
L52:
	;
	v136 = int32(4)
	goto L53
L53:
	;
	if l2 == int32(2) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v139 = v136
	goto L56
L55:
	;
	v139 = v132
	goto L56
L56:
	;
	v141 = v139
	goto L12
}
func F_find_nonnullable_rels(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_find_nonnullable_rels_walker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
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
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1273]))
	v24 = F_hash_search(m, v19, v13+int32(8), v5, v5)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v481
L2:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v463<<(uint(int32(2))%32))+uint32(_consts[1274])))
	v479 = F_find_option(m, v477, int32(0), l2, l3)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L8
	} else {
		goto L167
	}
L3:
	;
	v175 = int32(306171)
	v180 = v30
	goto L72
L4:
	;
	if v170 != 0 {
		goto L3
	} else {
		goto L69
	}
L5:
	;
	v170 = v165 + base.I32_extend8_s(v164)
	goto L4
L6:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v162 != 0 {
		goto L66
	} else {
		goto L67
	}
L7:
	;
	if base.Ui32((v31-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	return int32(0)
L9:
	;
	if v24 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v31 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v481 = v33
	goto L1
L13:
	;
	v161 = int32(306162)
	goto L6
L14:
	;
	v42 = v31 | int32(32)
	goto L16
L15:
	;
	v42 = v31
	goto L16
L16:
	;
	if v42 != int32(115) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v164 = v42
	v165 = int32(-115)
	goto L5
L18:
	;
	goto L19
L19:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v47 == int32(0) {
		v161 = int32(306163)
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v47-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v58 = v47 | int32(32)
	goto L23
L22:
	;
	v58 = v47
	goto L23
L23:
	;
	if v58 != int32(111) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v164 = v58
	v165 = int32(-111)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+2)))
	if v63 == int32(0) {
		v161 = int32(306164)
		goto L6
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v63-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = v63 | int32(32)
	goto L30
L29:
	;
	v74 = v63
	goto L30
L30:
	;
	if v74 != int32(114) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v164 = v74
	v165 = int32(-114)
	goto L5
L32:
	;
	goto L33
L33:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+3)))
	if v79 == int32(0) {
		v161 = int32(306165)
		goto L6
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32((v79-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v79 | int32(32)
	goto L37
L36:
	;
	v90 = v79
	goto L37
L37:
	;
	if v90 != int32(116) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v164 = v90
	v165 = int32(-116)
	goto L5
L39:
	;
	goto L40
L40:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v95 == int32(0) {
		v161 = int32(306166)
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32((v95-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = v95 | int32(32)
	goto L44
L43:
	;
	v106 = v95
	goto L44
L44:
	;
	if v106 != int32(95) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v164 = v106
	v165 = int32(-95)
	goto L5
L46:
	;
	goto L47
L47:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+5)))
	if v111 == int32(0) {
		v161 = int32(306167)
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v114 = int32(-109)
	if base.Ui32((v111-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v123 = v111 | int32(32)
	goto L51
L50:
	;
	v123 = v111
	goto L51
L51:
	;
	if v123 != int32(109) {
		v164 = v123
		v165 = v114
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+6)))
	if v127 == int32(0) {
		v161 = int32(306168)
		goto L6
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32((v127-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v138 = v127 | int32(32)
	goto L56
L55:
	;
	v138 = v127
	goto L56
L56:
	;
	if v138 != int32(101) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v164 = v138
	v165 = int32(-101)
	goto L5
L58:
	;
	goto L59
L59:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
	if v143 == int32(0) {
		v161 = int32(306169)
		goto L6
	} else {
		goto L60
	}
L60:
	;
	if base.Ui32((v143-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v154 = v143 | int32(32)
	goto L63
L62:
	;
	v154 = v143
	goto L63
L63:
	;
	if v154 != int32(109) {
		v164 = v154
		v165 = v114
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+8)))
	if v157 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v161 = int32(306170)
	goto L6
L66:
	;
	v163 = int32(-1)
	goto L68
L67:
	;
	v163 = v5
	goto L68
L68:
	;
	v170 = v163
	goto L4
L69:
	;
	v463 = int32(1)
	goto L2
L70:
	;
	v269 = int32(360652)
	v274 = v30
	goto L101
L71:
	;
	if v261 != 0 {
		goto L70
	} else {
		goto L98
	}
L72:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v187 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v261 = base.I32_extend8_s(v251) - base.I32_extend8_s(v249)
	goto L71
L74:
	;
	goto L73
L75:
	;
	v222 = int32(2)
	if base.Ui32((v212-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v175 == int32(306181) {
		goto L70
	} else {
		goto L79
	}
L77:
	;
	v217 = v175
	goto L78
L78:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v218 != 0 {
		goto L88
	} else {
		goto L89
	}
L79:
	;
	if base.Ui32((v187-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v198 = v187 | int32(32)
	goto L82
L81:
	;
	v198 = v187
	goto L82
L82:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v208 = v199 | int32(32)
	goto L85
L84:
	;
	v208 = v199
	goto L85
L85:
	;
	if v198 != v208&int32(255) {
		v249 = v208
		v251 = v198
		goto L74
	} else {
		goto L86
	}
L86:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v212 != 0 {
		goto L75
	} else {
		goto L87
	}
L87:
	;
	v217 = v175 + int32(1)
	goto L78
L88:
	;
	v219 = int32(-1)
	goto L90
L89:
	;
	v219 = int32(0)
	goto L90
L90:
	;
	v261 = v219
	goto L71
L91:
	;
	v234 = v212 | int32(32)
	goto L93
L92:
	;
	v234 = v212
	goto L93
L93:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(1)))))
	if base.Ui32((v235-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v244 = v235 | int32(32)
	goto L96
L95:
	;
	v244 = v235
	goto L96
L96:
	;
	if v234 == v244&int32(255) {
		v175 = v175 + v222
		v180 = v180 + v222
		goto L72
	} else {
		goto L97
	}
L97:
	;
	v249 = v244
	v251 = v234
	goto L74
L98:
	;
	v463 = int32(3)
	goto L2
L99:
	;
	if l1 != 0 {
		goto L130
	} else {
		goto L131
	}
L100:
	;
	if v355 != 0 {
		goto L99
	} else {
		goto L127
	}
L101:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v281 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v355 = base.I32_extend8_s(v345) - base.I32_extend8_s(v343)
	goto L100
L103:
	;
	goto L102
L104:
	;
	v316 = int32(2)
	if base.Ui32((v306-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L105:
	;
	if v269 == int32(360666) {
		goto L99
	} else {
		goto L108
	}
L106:
	;
	v311 = v269
	goto L107
L107:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v312 != 0 {
		goto L117
	} else {
		goto L118
	}
L108:
	;
	if base.Ui32((v281-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v292 = v281 | int32(32)
	goto L111
L110:
	;
	v292 = v281
	goto L111
L111:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if base.Ui32((v293-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v302 = v293 | int32(32)
	goto L114
L113:
	;
	v302 = v293
	goto L114
L114:
	;
	if v292 != v302&int32(255) {
		v343 = v302
		v345 = v292
		goto L103
	} else {
		goto L115
	}
L115:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v306 != 0 {
		goto L104
	} else {
		goto L116
	}
L116:
	;
	v311 = v269 + int32(1)
	goto L107
L117:
	;
	v313 = int32(-1)
	goto L119
L118:
	;
	v313 = int32(0)
	goto L119
L119:
	;
	v355 = v313
	goto L100
L120:
	;
	v328 = v306 | int32(32)
	goto L122
L121:
	;
	v328 = v306
	goto L122
L122:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+int32(1)))))
	if base.Ui32((v329-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v338 = v329 | int32(32)
	goto L125
L124:
	;
	v338 = v329
	goto L125
L125:
	;
	if v328 == v338&int32(255) {
		v269 = v269 + v316
		v274 = v274 + v316
		goto L101
	} else {
		goto L126
	}
L126:
	;
	v343 = v338
	v345 = v328
	goto L103
L127:
	;
	v463 = int32(5)
	goto L2
L128:
	;
	F_pfree(m, v397)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L166
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+4)) = v397
	v481 = v397
	goto L1
L130:
	;
	v362 = F_assignable_custom_variable_name(m, v30, l2, l3)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L8
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v439 = int32(0)
	if l2 != 0 {
		v481 = v439
		goto L1
	} else {
		goto L160
	}
L133:
	;
	if v362 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v481 = int32(0)
	goto L1
L135:
	;
	goto L136
L136:
	;
	v367 = int32(0)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v370 = *(*int32)(unsafe.Add(mBase, _consts[1275]))
	v373 = F_MemoryContextAllocExtended(m, v370, int32(124), int32(2))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	if v373 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v378 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v397 = F__emscripten_memset_bulkmem(m, v373, base.I32_extend8_s(int32(0)), int32(124))
	mBase = m.M
	goto L146
L141:
	;
	if v378 == int32(0) {
		v481 = v367
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(14086), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L8
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(525932), int32(647), int32(512753))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	v481 = v367
	goto L1
L146:
	;
	v398 = F_guc_strdup(m, l3, v368)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = v398
	if v398 == int32(0) {
		goto L128
	} else {
		goto L148
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v397)+20)) = int64(12884902532)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(416348)
	*(*int64)(unsafe.Add(mBase, uint32(v397)+4)) = int64(197568495622)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+92)) = v397 + int32(120)
	v413 = *(*int32)(unsafe.Add(mBase, _consts[1273]))
	v417 = F_hash_search(m, v413, v397, int32(3), v13+int32(15))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L149
	}
L149:
	;
	if v417 != 0 {
		goto L129
	} else {
		goto L150
	}
L150:
	;
	v420 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	if v420 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L8
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if v434 == int32(0) {
		goto L128
	} else {
		goto L158
	}
L155:
	;
	F_errmsg(m, int32(14086), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(525932), int32(1060), int32(416061))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	F_pfree(m, v434)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	goto L128
L160:
	;
	v441 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L8
	} else {
		goto L161
	}
L161:
	;
	if v441 == int32(0) {
		v481 = v439
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v448
	F_errmsg(m, int32(736353), v13)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(525932), int32(1279), int32(259787))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L8
	} else {
		goto L165
	}
L165:
	;
	v481 = v439
	goto L1
L166:
	;
	v481 = int32(0)
	goto L1
L167:
	;
	v481 = v479
	goto L1
}
func F_find_typed_table_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v14 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v8+int32(-48), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = F_table_beginscan_catalog(m, v14, int32(1), v8+int32(-48))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	m.T0[v62].(func(*base.Module, int32))(m, v28)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v30 = F_heap_getnext(m, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = int32(0)
	goto L5
L9:
	;
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v38 = int32(0)
	v42 = v30
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)))
	v49 = F_lappend_oid(m, v38, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v53 = v49
	goto L5
L14:
	;
	v51 = F_heap_getnext(m, v28)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v51 != 0 {
		v38 = v49
		v42 = v51
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_sequence_close(m, v14, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 - int32(-64)
	return v53
L19:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg(m, int32(414206), v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(642517), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(519627), int32(7120), int32(179240))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fix_opfuncids(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_fix_opfuncids_walker(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_flatCopyTargetEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	v4 = F_palloc0(m, int32(28))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(62)
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v16
		return v4
	}
}
func F_float48le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float32
	_ = v13
	var v14 float64
	_ = v14
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v13 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = base.F64_promote_f32(v13)
		v23 = base.F64_ge(v7, v14) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v23 = int32(1)
	}
	return v23
}
func F_float48ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = base.F64_promote_f32(v10)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313)))
	} else {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9)) | base.F64_ne(v6, v11)
	}
}
func F_float4_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v8 float32
	_ = v8
	var v20 int32
	_ = v20
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = base.F32_abs(base.F32_sub(v5, v6))
	if base.F32_ne(v8, math.Float32frombits(uint32(0x7f800000))) != 0 {
		return base.I32_reinterpret_f32(v8)
	} else {
		if base.F32_eq(base.F32_abs(v5), math.Float32frombits(uint32(0x7f800000))) != 0 {
			return base.I32_reinterpret_f32(v8)
		} else {
			if base.F32_eq(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000))) != 0 {
				return base.I32_reinterpret_f32(v8)
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
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
func F_float4_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 float32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 float64
	_ = v158
	var v162 float32
	_ = v162
	var v163 float64
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v22 == int32(1) {
			v25 = int32(4)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v27&int32(254) == int32(2) {
				v36 = v25
			} else {
				v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
			}
			if v27 == int32(1) {
				v39 = v25
			} else {
				v39 = v36
			}
			v52 = v39
		} else {
			v40 = int32(1)
			if v22&v40 != 0 {
				v52 = int32(base.Ui32(v22)>>(uint(v40)%32)) - v40
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v52-int32(1)) {
			v58 = F_cstring_to_text(m, int32(793540))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v260 = v58
				m.G0 = v14 + int32(96)
				return v260
			}
		} else {
			v60 = base.F32_reinterpret_i32(v16)
			v65 = F_palloc0(m, v52<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v71 = F_NUM_cache(m, v52, v14+int32(60), v18, v14+int32(59))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v73&int32(1024) != 0 {
						v76 = base.F32_nearest(v60)
						if base.F32_lt(base.F32_abs(v76), float32(2.1474836e+09)) != 0 {
							v86 = base.I32_trunc_f32_s(v76)
							v88 = v86
						} else {
							v88 = int32(-2147483648)
						}
						if base.Ui32(base.I32_reinterpret_f32(v76)&int32(2147483647)) < base.Ui32(int32(2139095041)) {
							v90 = v88
						} else {
							v90 = int32(2147483647)
						}
						if base.F32_ge(v76, float32(-2.1474836e+09)) != 0 {
							v94 = v90
						} else {
							v94 = int32(2147483647)
						}
						if base.F32_lt(v76, float32(2.1474836e+09)) != 0 {
							v98 = v94
						} else {
							v98 = int32(2147483647)
						}
						v99 = F_int_to_roman(m, v98)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v227 = v99
							v229 = v2
							v230 = int32(0)
							v238 = v65 + int32(4)
							F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
							mBase = m.M
							v242 = m.ExcPending
							if v242 != 0 {
								return int32(0)
							} else {
								v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v243 == int32(1) {
									F_pfree(m, v71)
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return int32(0)
									} else {
										v248 = F_strlen(m, v238)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
										v260 = v65
										m.G0 = v14 + int32(96)
										return v260
									}
								} else {
									v248 = F_strlen(m, v238)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
									v260 = v65
									m.G0 = v14 + int32(96)
									return v260
								}
							}
						}
					} else {
						if v73&int32(16384) != 0 {
							if base.B2i32(base.Ui32(v16&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_ne(base.F32_abs(v60), math.Float32frombits(uint32(0x7f800000))) == int32(0) {
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v116 = v114 + v115
								v119 = F_palloc(m, v116+int32(7))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v123 = v116 + int32(6)
									v125 = F__emscripten_memset_bulkmem(m, v119, base.I32_extend8_s(int32(35)), v123)
									mBase = m.M
									v127 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v125+v123))) = uint8(v127)
									v129 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v129)
									v132 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v114+v125)+1)) = uint8(v132)
									v227 = v119
									v229 = v127
									v230 = int32(0)
									v238 = v65 + int32(4)
									F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int32(0)
									} else {
										v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v243 == int32(1) {
											F_pfree(m, v71)
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return int32(0)
											} else {
												v248 = F_strlen(m, v238)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
												v260 = v65
												m.G0 = v14 + int32(96)
												return v260
											}
										} else {
											v248 = F_strlen(m, v238)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
											v260 = v65
											m.G0 = v14 + int32(96)
											return v260
										}
									}
								}
							} else {
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v135
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = base.F64_promote_f32(v60)
								v139 = int32(0)
								v143 = F_psprintf(m, int32(441295), v14+int32(32))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
									if v145 != int32(43) {
										v227 = v143
										v229 = v2
										v230 = v139
									} else {
										v148 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v148)
										v227 = v143
										v229 = v2
										v230 = v139
									}
									v238 = v65 + int32(4)
									F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int32(0)
									} else {
										v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v243 == int32(1) {
											F_pfree(m, v71)
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return int32(0)
											} else {
												v248 = F_strlen(m, v238)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
												v260 = v65
												m.G0 = v14 + int32(96)
												return v260
											}
										} else {
											v248 = F_strlen(m, v238)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
											v260 = v65
											m.G0 = v14 + int32(96)
											return v260
										}
									}
								}
							}
						} else {
							if v73&int32(2048) != 0 {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v152 + v153
								v158 = F_pow(m, float64(10), base.F64_convert_i32_s(v152))
								mBase = m.M
								v162 = base.F32_mul(v60, base.F32_demote_f64(v158))
							} else {
								v162 = v60
							}
							v163 = base.F64_promote_f32(v162)
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v163)
							v169 = F_psprintf(m, int32(357879), v14+int32(16))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								v171 = F_strlen(m, v169)
								mBase = m.M
								if v171 <= int32(5) {
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if v174+v171 < int32(7) {
										v182 = v174
									} else {
										v180 = int32(6) - v171
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v180
										v182 = v180
									}
								} else {
									v180 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v180
									v182 = v180
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v163
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v182
								v186 = F_psprintf(m, int32(357884), v14)
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int32(0)
								} else {
									v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
									v190 = base.B2i32(v188 == int32(45))
									v191 = v186 + v190
									v192 = int32(46)
									v193 = F___strchrnul(m, v191, v192)
									mBase = m.M
									v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
									if v195 == v192 {
										v199 = v193
									} else {
										v199 = int32(0)
									}
									if v199 != 0 {
										v202 = v199 - v191
									} else {
										v201 = F_strlen(m, v191)
										mBase = m.M
										v202 = v201
									}
									if v188 == int32(45) {
										v205 = int32(45)
									} else {
										v205 = int32(43)
									}
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if v202 < v206 {
										v227 = v191
										v229 = v205
										v230 = v206 - v202
										v238 = v65 + int32(4)
										F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
											return int32(0)
										} else {
											v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v243 == int32(1) {
												F_pfree(m, v71)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													v248 = F_strlen(m, v238)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
													v260 = v65
													m.G0 = v14 + int32(96)
													return v260
												}
											} else {
												v248 = F_strlen(m, v238)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
												v260 = v65
												m.G0 = v14 + int32(96)
												return v260
											}
										}
									} else {
										v209 = int32(0)
										if v202 <= v206 {
											v227 = v191
											v229 = v205
											v230 = v209
											v238 = v65 + int32(4)
											F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
											mBase = m.M
											v242 = m.ExcPending
											if v242 != 0 {
												return int32(0)
											} else {
												v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v243 == int32(1) {
													F_pfree(m, v71)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														v248 = F_strlen(m, v238)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
														v260 = v65
														m.G0 = v14 + int32(96)
														return v260
													}
												} else {
													v248 = F_strlen(m, v238)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
													v260 = v65
													m.G0 = v14 + int32(96)
													return v260
												}
											}
										} else {
											v211 = v182 + v206
											v214 = F_palloc(m, v211+int32(2))
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return int32(0)
											} else {
												v218 = v211 + int32(1)
												v220 = F__emscripten_memset_bulkmem(m, v214, base.I32_extend8_s(int32(35)), v218)
												mBase = m.M
												v222 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v220+v218))) = uint8(v222)
												v225 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v220+v206))) = uint8(v225)
												v227 = v214
												v229 = v205
												v230 = v209
												v238 = v65 + int32(4)
												F_NUM_processor(m, v71, v14+int32(60), v238, v227, int32(0), v230, v229, int32(1))
												mBase = m.M
												v242 = m.ExcPending
												if v242 != 0 {
													return int32(0)
												} else {
													v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v243 == int32(1) {
														F_pfree(m, v71)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															v248 = F_strlen(m, v238)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
															v260 = v65
															m.G0 = v14 + int32(96)
															return v260
														}
													} else {
														v248 = F_strlen(m, v238)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v65))) = v248<<(uint(int32(2))%32) + int32(16)
														v260 = v65
														m.G0 = v14 + int32(96)
														return v260
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
func F_float4div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v14 float32
	_ = v14
	var v16 float32
	_ = v16
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.F32_eq(v6, float32(0)) != 0 {
		if base.Ui32(base.I32_reinterpret_f32(v5)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			F_float_zero_divide_error(m)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v14 = base.F32_div(v5, v6)
			v16 = math.Float32frombits(uint32(0x7f800000))
			if base.F32_eq(base.F32_abs(v14), v16)&base.F32_ne(base.F32_abs(v5), v16) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F32_ne(v14, float32(0)) != 0 {
					return base.I32_reinterpret_f32(v14)
				} else {
					if base.F32_eq(v5, float32(0)) != 0 {
						return base.I32_reinterpret_f32(v14)
					} else {
						if base.F32_ne(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000))) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return base.I32_reinterpret_f32(v14)
						}
					}
				}
			}
		}
	} else {
		v14 = base.F32_div(v5, v6)
		v16 = math.Float32frombits(uint32(0x7f800000))
		if base.F32_eq(base.F32_abs(v14), v16)&base.F32_ne(base.F32_abs(v5), v16) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if base.F32_ne(v14, float32(0)) != 0 {
				return base.I32_reinterpret_f32(v14)
			} else {
				if base.F32_eq(v5, float32(0)) != 0 {
					return base.I32_reinterpret_f32(v14)
				} else {
					if base.F32_ne(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000))) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return base.I32_reinterpret_f32(v14)
					}
				}
			}
		}
	}
}
func F_float4in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 float32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v122 float32
	_ = v122
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
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 float32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 float32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 float32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 float32
	_ = v519
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v539 float32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v565 float32
	_ = v565
	v9 = float32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = l0
	goto L3
L1:
	;
	m.G0 = v13 - int32(-64)
	return v565
L2:
	;
	v525 = v516
	goto L162
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.Ui32(v25-int32(9)) < base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v514 = v15 + v512
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v514
	v516 = v514
	v519 = v513
	goto L2
L5:
	;
	goto L4
L6:
	;
	v15 = v15 + int32(1)
	goto L3
L7:
	;
	if v25 == int32(32) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = F_errsave_start(m, l3)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v60 = F_strtof(m, v15, v11+int32(-4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float32(0)
L13:
	;
	if v34 == int32(0) {
		v565 = v9
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
	F_errmsg(m, int32(761048), v11+int32(-16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l3, int32(518199), int32(208), int32(327116))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v565 = v9
	goto L1
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v64 == v15 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = int32(3)
	v71 = v15
	v72 = int32(554746)
	v73 = v66
	goto L23
L20:
	;
	if v63 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v516 = v64
	v519 = v60
	goto L2
L22:
	;
	if v118 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v118 = int32(0)
	goto L22
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 == v77 {
		v99 = v76
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v101 = int32(1)
	if v99 != 0 {
		v71 = v71 + v101
		v72 = v72 + v101
		v73 = v73 - v101
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v76-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = v76 | int32(32)
	goto L32
L31:
	;
	v87 = v76
	goto L32
L32:
	;
	if base.Ui32((v77-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v77 | int32(32)
	goto L35
L34:
	;
	v96 = v77
	goto L35
L35:
	;
	if v87 == v96 {
		v99 = v87
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v118 = v87 - v96
	goto L22
L37:
	;
	goto L27
L38:
	;
	v512 = v66
	v513 = math.Float32frombits(uint32(0x7fc00000))
	goto L5
L39:
	;
	goto L40
L40:
	;
	v122 = math.Float32frombits(uint32(0x7f800000))
	v123 = int32(8)
	v128 = v15
	v129 = int32(11540)
	v130 = v123
	goto L42
L41:
	;
	if v175 == int32(0) {
		v512 = v123
		v513 = v122
		goto L5
	} else {
		goto L57
	}
L42:
	;
	if v130 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v175 = int32(0)
	goto L41
L44:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v133 == v134 {
		v156 = v133
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v158 = int32(1)
	if v156 != 0 {
		v128 = v128 + v158
		v129 = v129 + v158
		v130 = v130 - v158
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v133-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v144 = v133 | int32(32)
	goto L51
L50:
	;
	v144 = v133
	goto L51
L51:
	;
	if base.Ui32((v134-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v153 = v134 | int32(32)
	goto L54
L53:
	;
	v153 = v134
	goto L54
L54:
	;
	if v144 == v153 {
		v156 = v144
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v175 = v144 - v153
	goto L41
L56:
	;
	goto L46
L57:
	;
	v178 = int32(9)
	v183 = v15
	v184 = int32(11539)
	v185 = v178
	goto L59
L58:
	;
	if v230 == int32(0) {
		v512 = v178
		v513 = v122
		goto L5
	} else {
		goto L74
	}
L59:
	;
	if v185 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v230 = int32(0)
	goto L58
L61:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 == v189 {
		v211 = v188
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v213 = int32(1)
	if v211 != 0 {
		v183 = v183 + v213
		v184 = v184 + v213
		v185 = v185 - v213
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v188-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = v188 | int32(32)
	goto L68
L67:
	;
	v199 = v188
	goto L68
L68:
	;
	if base.Ui32((v189-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v208 = v189 | int32(32)
	goto L71
L70:
	;
	v208 = v189
	goto L71
L71:
	;
	if v199 == v208 {
		v211 = v199
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v230 = v199 - v208
	goto L58
L73:
	;
	goto L63
L74:
	;
	v237 = v15
	v238 = int32(11529)
	v239 = int32(9)
	goto L76
L75:
	;
	if v284 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v239 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v284 = int32(0)
	goto L75
L78:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v242 == v243 {
		v265 = v242
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v267 = int32(1)
	if v265 != 0 {
		v237 = v237 + v267
		v238 = v238 + v267
		v239 = v239 - v267
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v242-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v253 = v242 | int32(32)
	goto L85
L84:
	;
	v253 = v242
	goto L85
L85:
	;
	if base.Ui32((v243-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v262 = v243 | int32(32)
	goto L88
L87:
	;
	v262 = v243
	goto L88
L88:
	;
	if v253 == v262 {
		v265 = v253
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v284 = v253 - v262
	goto L75
L90:
	;
	goto L80
L91:
	;
	v512 = v178
	v513 = math.Float32frombits(uint32(0xff800000))
	goto L5
L92:
	;
	goto L93
L93:
	;
	v288 = int32(3)
	v293 = v15
	v294 = int32(356551)
	v295 = v288
	goto L95
L94:
	;
	if v340 == int32(0) {
		v512 = v288
		v513 = v122
		goto L5
	} else {
		goto L110
	}
L95:
	;
	if v295 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v340 = int32(0)
	goto L94
L97:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v298 == v299 {
		v321 = v298
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v323 = int32(1)
	if v321 != 0 {
		v293 = v293 + v323
		v294 = v294 + v323
		v295 = v295 - v323
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v309 = v298 | int32(32)
	goto L104
L103:
	;
	v309 = v298
	goto L104
L104:
	;
	if base.Ui32((v299-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v318 = v299 | int32(32)
	goto L107
L106:
	;
	v318 = v299
	goto L107
L107:
	;
	if v309 == v318 {
		v321 = v309
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v340 = v309 - v318
	goto L94
L109:
	;
	goto L99
L110:
	;
	v343 = int32(4)
	v348 = v15
	v349 = int32(356550)
	v350 = v343
	goto L112
L111:
	;
	if v395 == int32(0) {
		v512 = v343
		v513 = v122
		goto L5
	} else {
		goto L127
	}
L112:
	;
	if v350 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v395 = int32(0)
	goto L111
L114:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v353 == v354 {
		v376 = v353
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v378 = int32(1)
	if v376 != 0 {
		v348 = v348 + v378
		v349 = v349 + v378
		v350 = v350 - v378
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v353-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v364 = v353 | int32(32)
	goto L121
L120:
	;
	v364 = v353
	goto L121
L121:
	;
	if base.Ui32((v354-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v373 = v354 | int32(32)
	goto L124
L123:
	;
	v373 = v354
	goto L124
L124:
	;
	if v364 == v373 {
		v376 = v364
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v395 = v364 - v373
	goto L111
L126:
	;
	goto L116
L127:
	;
	v403 = v15
	v404 = int32(356545)
	v405 = int32(4)
	goto L129
L128:
	;
	if v450 == int32(0) {
		v512 = v343
		v513 = math.Float32frombits(uint32(0xff800000))
		goto L5
	} else {
		goto L144
	}
L129:
	;
	if v405 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v450 = int32(0)
	goto L128
L131:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v408 == v409 {
		v431 = v408
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v433 = int32(1)
	if v431 != 0 {
		v403 = v403 + v433
		v404 = v404 + v433
		v405 = v405 - v433
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v408-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v419 = v408 | int32(32)
	goto L138
L137:
	;
	v419 = v408
	goto L138
L138:
	;
	if base.Ui32((v409-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v428 = v409 | int32(32)
	goto L141
L140:
	;
	v428 = v409
	goto L141
L141:
	;
	if v419 == v428 {
		v431 = v419
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v450 = v419 - v428
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v63 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v455 = base.I32_reinterpret_f32(v60)
	if base.B2i32(v455&int32(2147483647) == int32(0))|base.B2i32(v455 == int32(2139095040)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v490 = float32(0)
	v491 = F_errsave_start(m, l3)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v465 = F_pstrdup(m, v15)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L12
	} else {
		goto L151
	}
L149:
	;
	if base.F32_eq(v60, math.Float32frombits(uint32(0xff800000))) != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v516 = v64
	v519 = v60
	goto L2
L151:
	;
	v469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465+(v64-v15)))) = uint8(v469)
	v471 = float32(0)
	v472 = F_errsave_start(m, l3)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v472 == int32(0) {
		v565 = v471
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v465
	F_errmsg(m, int32(329885), v11+int32(-48))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l3, int32(518199), int32(288), int32(327116))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v565 = v471
	goto L1
L157:
	;
	if v491 == int32(0) {
		v565 = v490
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(761048), v11+int32(-32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l3, int32(518199), int32(295), int32(327116))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v565 = v490
	goto L1
L162:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	if base.Ui32(v530-int32(9)) < base.Ui32(int32(5)) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v539 = float32(0)
	v540 = F_errsave_start(m, l3)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L12
	} else {
		goto L169
	}
L164:
	;
	goto L163
L165:
	;
	v525 = v525 + int32(1)
	goto L162
L166:
	;
	if v530 == int32(32) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	if v530 != 0 {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v565 = v519
	goto L1
L169:
	;
	if v540 == int32(0) {
		v565 = v539
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L12
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg(m, int32(761048), v13)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	F_errsave_finish(m, l3, int32(518199), int32(309), int32(327116))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	v565 = v539
	goto L1
}
func F_float4up(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2
}
func F_float84gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v21 int32
	_ = v21
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_float8div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v16 float64
	_ = v16
	var v18 float64
	_ = v18
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.F64_eq(v8, float64(0)) != 0 {
		if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			F_float_zero_divide_error(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v16 = base.F64_div(v6, v8)
			v18 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v16), v18)&base.F64_ne(base.F64_abs(v6), v18) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_ne(v16, float64(0)) != 0 {
					v31 = F_Float8GetDatum(m, v16)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						return v31
					}
				} else {
					if base.F64_eq(v6, float64(0)) != 0 {
						v31 = F_Float8GetDatum(m, v16)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							return v31
						}
					} else {
						if base.F64_ne(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v31 = F_Float8GetDatum(m, v16)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								return v31
							}
						}
					}
				}
			}
		}
	} else {
		v16 = base.F64_div(v6, v8)
		v18 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(base.F64_abs(v16), v18)&base.F64_ne(base.F64_abs(v6), v18) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if base.F64_ne(v16, float64(0)) != 0 {
				v31 = F_Float8GetDatum(m, v16)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					return v31
				}
			} else {
				if base.F64_eq(v6, float64(0)) != 0 {
					v31 = F_Float8GetDatum(m, v16)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						return v31
					}
				} else {
					if base.F64_ne(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v31 = F_Float8GetDatum(m, v16)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							return v31
						}
					}
				}
			}
		}
	}
}
func F_float8gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v122 int32
	_ = v122
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
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
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v454 int32
	_ = v454
	var v459 int64
	_ = v459
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 float64
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 float64
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 float64
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 float64
	_ = v522
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v548 float64
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v575 float64
	_ = v575
	v10 = float64(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = l0
	goto L3
L1:
	;
	m.G0 = v15 - int32(-64)
	return v575
L2:
	;
	v530 = v520
	goto L164
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.Ui32(v29-int32(9)) < base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v518 = v17 + v516
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v518
	v520 = v518
	v522 = v517
	goto L2
L5:
	;
	goto L4
L6:
	;
	v17 = v17 + int32(1)
	goto L3
L7:
	;
	if v29 == int32(32) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = F_errsave_start(m, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v64 = F_strtod(m, v17, v13+int32(-4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float64(0)
L13:
	;
	if v38 == int32(0) {
		v575 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	F_errmsg(m, int32(761048), v13+int32(-16))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l4, int32(518199), int32(414), int32(327098))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v575 = v10
	goto L1
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v68 == v17 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v70 = int32(3)
	v75 = v17
	v76 = int32(554746)
	v77 = v70
	goto L23
L20:
	;
	if v67 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v520 = v68
	v522 = v64
	goto L2
L22:
	;
	if v122 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v77 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v122 = int32(0)
	goto L22
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		v103 = v80
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v105 = int32(1)
	if v103 != 0 {
		v75 = v75 + v105
		v76 = v76 + v105
		v77 = v77 - v105
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v80-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v91 = v80 | int32(32)
	goto L32
L31:
	;
	v91 = v80
	goto L32
L32:
	;
	if base.Ui32((v81-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v100 = v81 | int32(32)
	goto L35
L34:
	;
	v100 = v81
	goto L35
L35:
	;
	if v91 == v100 {
		v103 = v91
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v122 = v91 - v100
	goto L22
L37:
	;
	goto L27
L38:
	;
	v516 = v70
	v517 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L5
L39:
	;
	goto L40
L40:
	;
	v126 = math.Float64frombits(uint64(0x7ff0000000000000))
	v127 = int32(8)
	v132 = v17
	v133 = int32(11540)
	v134 = v127
	goto L42
L41:
	;
	if v179 == int32(0) {
		v516 = v127
		v517 = v126
		goto L5
	} else {
		goto L57
	}
L42:
	;
	if v134 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v179 = int32(0)
	goto L41
L44:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v137 == v138 {
		v160 = v137
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v162 = int32(1)
	if v160 != 0 {
		v132 = v132 + v162
		v133 = v133 + v162
		v134 = v134 - v162
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v137-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v148 = v137 | int32(32)
	goto L51
L50:
	;
	v148 = v137
	goto L51
L51:
	;
	if base.Ui32((v138-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v157 = v138 | int32(32)
	goto L54
L53:
	;
	v157 = v138
	goto L54
L54:
	;
	if v148 == v157 {
		v160 = v148
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v179 = v148 - v157
	goto L41
L56:
	;
	goto L46
L57:
	;
	v182 = int32(9)
	v187 = v17
	v188 = int32(11539)
	v189 = v182
	goto L59
L58:
	;
	if v234 == int32(0) {
		v516 = v182
		v517 = v126
		goto L5
	} else {
		goto L74
	}
L59:
	;
	if v189 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v234 = int32(0)
	goto L58
L61:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == v193 {
		v215 = v192
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v217 = int32(1)
	if v215 != 0 {
		v187 = v187 + v217
		v188 = v188 + v217
		v189 = v189 - v217
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v192-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v203 = v192 | int32(32)
	goto L68
L67:
	;
	v203 = v192
	goto L68
L68:
	;
	if base.Ui32((v193-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v212 = v193 | int32(32)
	goto L71
L70:
	;
	v212 = v193
	goto L71
L71:
	;
	if v203 == v212 {
		v215 = v203
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v234 = v203 - v212
	goto L58
L73:
	;
	goto L63
L74:
	;
	v241 = v17
	v242 = int32(11529)
	v243 = int32(9)
	goto L76
L75:
	;
	if v288 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v243 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v288 = int32(0)
	goto L75
L78:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v246 == v247 {
		v269 = v246
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v271 = int32(1)
	if v269 != 0 {
		v241 = v241 + v271
		v242 = v242 + v271
		v243 = v243 - v271
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v246-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v257 = v246 | int32(32)
	goto L85
L84:
	;
	v257 = v246
	goto L85
L85:
	;
	if base.Ui32((v247-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v266 = v247 | int32(32)
	goto L88
L87:
	;
	v266 = v247
	goto L88
L88:
	;
	if v257 == v266 {
		v269 = v257
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v288 = v257 - v266
	goto L75
L90:
	;
	goto L80
L91:
	;
	v516 = v182
	v517 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L5
L92:
	;
	goto L93
L93:
	;
	v292 = int32(3)
	v297 = v17
	v298 = int32(356551)
	v299 = v292
	goto L95
L94:
	;
	if v344 == int32(0) {
		v516 = v292
		v517 = v126
		goto L5
	} else {
		goto L110
	}
L95:
	;
	if v299 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v344 = int32(0)
	goto L94
L97:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v302 == v303 {
		v325 = v302
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v327 = int32(1)
	if v325 != 0 {
		v297 = v297 + v327
		v298 = v298 + v327
		v299 = v299 - v327
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v302-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v313 = v302 | int32(32)
	goto L104
L103:
	;
	v313 = v302
	goto L104
L104:
	;
	if base.Ui32((v303-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v322 = v303 | int32(32)
	goto L107
L106:
	;
	v322 = v303
	goto L107
L107:
	;
	if v313 == v322 {
		v325 = v313
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v344 = v313 - v322
	goto L94
L109:
	;
	goto L99
L110:
	;
	v347 = int32(4)
	v352 = v17
	v353 = int32(356550)
	v354 = v347
	goto L112
L111:
	;
	if v399 == int32(0) {
		v516 = v347
		v517 = v126
		goto L5
	} else {
		goto L127
	}
L112:
	;
	if v354 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v399 = int32(0)
	goto L111
L114:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v357 == v358 {
		v380 = v357
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v382 = int32(1)
	if v380 != 0 {
		v352 = v352 + v382
		v353 = v353 + v382
		v354 = v354 - v382
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v357-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v368 = v357 | int32(32)
	goto L121
L120:
	;
	v368 = v357
	goto L121
L121:
	;
	if base.Ui32((v358-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v377 = v358 | int32(32)
	goto L124
L123:
	;
	v377 = v358
	goto L124
L124:
	;
	if v368 == v377 {
		v380 = v368
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v399 = v368 - v377
	goto L111
L126:
	;
	goto L116
L127:
	;
	v407 = v17
	v408 = int32(356545)
	v409 = int32(4)
	goto L129
L128:
	;
	if v454 == int32(0) {
		v516 = v347
		v517 = math.Float64frombits(uint64(0xfff0000000000000))
		goto L5
	} else {
		goto L144
	}
L129:
	;
	if v409 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v454 = int32(0)
	goto L128
L131:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if v412 == v413 {
		v435 = v412
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v437 = int32(1)
	if v435 != 0 {
		v407 = v407 + v437
		v408 = v408 + v437
		v409 = v409 - v437
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v412-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v423 = v412 | int32(32)
	goto L138
L137:
	;
	v423 = v412
	goto L138
L138:
	;
	if base.Ui32((v413-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v432 = v413 | int32(32)
	goto L141
L140:
	;
	v432 = v413
	goto L141
L141:
	;
	if v423 == v432 {
		v435 = v423
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v454 = v423 - v432
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v67 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v459 = base.I64_reinterpret_f64(v64)
	if base.B2i32(v459&int64(9223372036854775807) == int64(0))|base.B2i32(v459 == int64(9218868437227405312)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v494 = float64(0)
	v495 = F_errsave_start(m, l4)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v469 = F_pstrdup(m, v17)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L151
	}
L149:
	;
	if base.F64_eq(v64, math.Float64frombits(uint64(0xfff0000000000000))) != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v520 = v68
	v522 = v64
	goto L2
L151:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v469+(v68-v17)))) = uint8(v473)
	v475 = float64(0)
	v476 = F_errsave_start(m, l4)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v476 == int32(0) {
		v575 = v475
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v469
	F_errmsg(m, int32(285575), v13+int32(-48))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l4, int32(518199), int32(490), int32(327098))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v575 = v475
	goto L1
L157:
	;
	if v495 == int32(0) {
		v575 = v494
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l2
	F_errmsg(m, int32(761048), v13+int32(-32))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l4, int32(518199), int32(497), int32(327098))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v575 = v494
	goto L1
L162:
	;
	v548 = float64(0)
	v549 = F_errsave_start(m, l4)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L12
	} else {
		goto L173
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v530
	v575 = v522
	goto L1
L164:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if base.Ui32(v536-int32(9)) < base.Ui32(int32(5)) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	if l1 == int32(0) {
		goto L162
	} else {
		goto L172
	}
L166:
	;
	goto L165
L167:
	;
	v530 = v530 + int32(1)
	goto L164
L168:
	;
	if v536 == int32(32) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	if v536 != 0 {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	if l1 != 0 {
		goto L163
	} else {
		goto L171
	}
L171:
	;
	v575 = v522
	goto L1
L172:
	;
	goto L163
L173:
	;
	if v549 == int32(0) {
		v575 = v548
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	F_errmsg(m, int32(761048), v15)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	F_errsave_finish(m, l4, int32(518199), int32(511), int32(327098))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L177
	}
L177:
	;
	v575 = v548
	goto L1
}
func F_float8le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		v22 = base.F64_ge(v6, v13) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_flt4_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 float32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_cash_mul_float8(m, v3, base.F64_promote_f32(v4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_flt8_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_cash_mul_float8(m, v3, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_fmt_u(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	if base.Ui64(l0) < base.Ui64(int64(4294967296)) {
		v29 = l1
		v33 = l0
	} else {
		v9 = l0
		v10 = l1
		for {
			v16 = v10 - int32(1)
			v17 = int64(10)
			v18 = base.I64_div_u_s(v9, v17)
			v24 = base.I32_wrap_i64(v9-v18*v17) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v24)
			if base.Ui64(int64(42949672959)) < base.Ui64(v9) {
				v9 = v18
				v10 = v16
				continue
			} else {
				break
			}
			break
		}
		v29 = v16
		v33 = v18
	}
	if v33 != int64(0) {
		v38 = v29
		v39 = base.I32_wrap_i64(v33)
		for {
			v44 = v38 - int32(1)
			v45 = int32(10)
			v46 = base.I32_div_u_s(v39, v45)
			v51 = v39 - v46*v45 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v51)
			if base.Ui32(int32(9)) < base.Ui32(v39) {
				v38 = v44
				v39 = v46
				continue
			} else {
				break
			}
			break
		}
		v56 = v44
	} else {
		v56 = v29
	}
	return v56
}
func F_forkname_chars(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	v5 = int32(3)
	v6 = int32(302010)
	goto L6
L1:
	;
	return v166
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v164
	v166 = v163
	goto L1
L3:
	;
	if l1 == int32(0) {
		v166 = v159
		goto L1
	} else {
		goto L55
	}
L4:
	;
	if v43-v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L7
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1320])))
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v14 = v6
	v15 = l0
	v16 = v5
	v17 = v13
	goto L12
L9:
	;
	v39 = l0
	v43 = int32(0)
	goto L10
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	goto L4
L11:
	;
	v39 = v34
	v43 = v36
	goto L10
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v17 != v19 {
		v34 = v15
		v36 = v17
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v34 = v28
	v36 = int32(0)
	goto L11
L14:
	;
	if v19 == int32(0) {
		v34 = v15
		v36 = v17
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v24 = v16 - int32(1)
	if v24 == int32(0) {
		v34 = v15
		v36 = v17
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v27 = int32(1)
	v28 = v15 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v29 != 0 {
		v14 = v14 + v27
		v15 = v28
		v16 = v24
		v17 = v29
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v159 = v5
	v160 = int32(1)
	goto L3
L19:
	;
	goto L20
L20:
	;
	v55 = int32(2)
	v57 = int32(299944)
	goto L23
L21:
	;
	if v94-v95 == int32(0) {
		v159 = v55
		v160 = v55
		goto L3
	} else {
		goto L35
	}
L23:
	;
	goto L24
L24:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1321])))
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = v57
	v66 = l0
	v67 = v55
	v68 = v64
	goto L29
L26:
	;
	v90 = l0
	v94 = int32(0)
	goto L27
L27:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	goto L21
L28:
	;
	v90 = v85
	v94 = v87
	goto L27
L29:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v68 != v70 {
		v85 = v66
		v87 = v68
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v85 = v79
	v87 = int32(0)
	goto L28
L31:
	;
	if v70 == int32(0) {
		v85 = v66
		v87 = v68
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v75 = v67 - int32(1)
	if v75 == int32(0) {
		v85 = v66
		v87 = v68
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v78 = int32(1)
	v79 = v66 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v80 != 0 {
		v65 = v65 + v78
		v66 = v79
		v67 = v75
		v68 = v80
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v105 = int32(4)
	v106 = int32(107261)
	goto L38
L36:
	;
	if v143-v144 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	goto L39
L39:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1322])))
	if v113 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v114 = v106
	v115 = l0
	v116 = v105
	v117 = v113
	goto L44
L41:
	;
	v139 = l0
	v143 = int32(0)
	goto L42
L42:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	goto L36
L43:
	;
	v139 = v134
	v143 = v136
	goto L42
L44:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v117 != v119 {
		v134 = v115
		v136 = v117
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v134 = v128
	v136 = int32(0)
	goto L43
L46:
	;
	if v119 == int32(0) {
		v134 = v115
		v136 = v117
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v124 = v116 - int32(1)
	if v124 == int32(0) {
		v134 = v115
		v136 = v117
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v127 = int32(1)
	v128 = v115 + v127
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v129 != 0 {
		v114 = v114 + v127
		v115 = v128
		v116 = v124
		v117 = v129
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	if l1 != 0 {
		v163 = v105
		v164 = int32(3)
		goto L2
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v156 = int32(0)
	if l1 == v156 {
		v166 = v156
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v166 = v105
	goto L1
L54:
	;
	v163 = v156
	v164 = int32(-1)
	goto L2
L55:
	;
	v163 = v159
	v164 = v160
	goto L2
}
func F_free_parsestate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v8-int32(1) < int32(1665) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v13 != 0 {
			F_sequence_close(m, v13, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(17039621))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1664)
				F_errmsg(m, int32(178387), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(525367), int32(83), int32(369938))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
}
func F_free_struct_lconv(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_emscripten_builtin_free(m, v2)
	mBase = m.M
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_emscripten_builtin_free(m, v4)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_emscripten_builtin_free(m, v6)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_emscripten_builtin_free(m, v8)
	mBase = m.M
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_emscripten_builtin_free(m, v10)
	mBase = m.M
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_emscripten_builtin_free(m, v12)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_emscripten_builtin_free(m, v14)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_emscripten_builtin_free(m, v16)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_emscripten_builtin_free(m, v20)
	mBase = m.M
	return
}
func F_freeifaddrs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l0 != 0 {
		v3 = l0
		for {
			v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			F_emscripten_builtin_free(m, v3)
			mBase = m.M
			if v5 != 0 {
				v3 = v5
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_frexp(m *base.Module, l0 float64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 float64
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v70 float64
	_ = v70
	v5 = base.I64_reinterpret_f64(l0)
	v9 = int32(2047)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & v9
	if v10 != v9 {
		if v10 == int32(0) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v58 = l0
				v59 = int32(0)
			} else {
				v19 = base.F64_mul(l0, float64(1.8446744073709552e+19))
				v22 = base.I64_reinterpret_f64(v19)
				v26 = int32(2047)
				v27 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(52))%64))) & v26
				if v27 != v26 {
					if v27 == int32(0) {
						if base.F64_eq(v19, float64(0)) != 0 {
							v41 = v19
							v42 = int32(0)
						} else {
							v37 = F_frexp(m, base.F64_mul(v19, float64(1.8446744073709552e+19)), l1)
							mBase = m.M
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v41 = v37
							v42 = v38 + int32(-64)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
						v54 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27 - int32(1022)
						v52 = base.F64_reinterpret_i64(v22&int64(-9218868437227405313) | int64(4602678819172646912))
						v54 = v52
					}
				} else {
					v52 = v19
					v54 = v52
				}
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v58 = v54
				v59 = v55 + int32(-64)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
			return v58
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10 - int32(1022)
			v70 = base.F64_reinterpret_i64(v5&int64(-9218868437227405313) | int64(4602678819172646912))
			return v70
		}
	} else {
		v70 = l0
		return v70
	}
}
func F_fscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_vfscanf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_fsm_get_avail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1)+uint32(_consts[651]))))
	return v6
}
func F_fsm_search_avail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v263
L2:
	;
	v35 = v33 + int32(28)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.Ui32(v36) < base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+(l0^int32(-1))<<(uint(int32(2))%32))))
	v33 = v25
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v33 = v27 + l0<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v263 = int32(-1)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v44 = l3
	goto L9
L9:
	;
	v52 = int32(4095)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(24))))
	v55 = v53 + v52
	if base.Ui32(int32(4068)) < base.Ui32(v53) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v263 = int32(-1)
	goto L1
L11:
	;
	if v93 <= int32(4094) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v58 = v52
	goto L14
L13:
	;
	v58 = v55
	goto L14
L14:
	;
	if v58 <= int32(0) {
		v93 = v55
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v66 = v58
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v35))))
	if base.Ui32(l1) <= base.Ui32(v73) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v93 = v86
	goto L11
L18:
	;
	v93 = v66
	goto L11
L19:
	;
	goto L20
L20:
	;
	v75 = int32(1)
	if (v66+int32(2))&(v66+v75) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v66
	goto L23
L22:
	;
	v84 = int32(base.Ui32(v66)>>(uint(v75)%32)) - v75
	goto L23
L23:
	;
	v86 = base.I32_div_s(v84, int32(2))
	if int32(1) < v84 {
		v66 = v86
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	F_MarkBufferDirtyHint(m, l0, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L45
	} else {
		goto L70
	}
L26:
	;
	v106 = v93
	goto L29
L27:
	;
	v239 = v93
	goto L28
L28:
	;
	v249 = (v239 + int32(61441)) & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v249 + l2
	v263 = v249
	goto L1
L29:
	;
	v114 = v106 << (uint(int32(1)) % 32)
	if base.Ui32(v114) <= base.Ui32(int32(8163)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v239 = v232
	goto L28
L31:
	;
	if v232 < int32(4095) {
		v106 = v232
		goto L29
	} else {
		goto L69
	}
L32:
	;
	v118 = v114 | int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v118))))
	if base.Ui32(l1) <= base.Ui32(v120) {
		v232 = v118
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v124 = v114 + int32(2)
	if base.Ui32(v124) <= base.Ui32(int32(8163)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v35))))
	if base.Ui32(l1) <= base.Ui32(v128) {
		v232 = v124
		goto L31
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v131 = v14 + int32(36)
	if l0 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v164 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v155
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v154
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28)))) = v160
	goto L40
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v153 = v140 + (l0^int32(-1))<<(uint(int32(6))%32)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v153 = v147 + l0<<(uint(int32(6))%32) + int32(-64)
	goto L41
L45:
	;
	return int32(0)
L46:
	;
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v172
	F_errmsg_internal(m, int32(41617), v14)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v44&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_errfinish(m, int32(525215), int32(277), int32(320741))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	F_LockBuffer(m, l0, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L45
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v200 = int32(4094)
	goto L57
L55:
	;
	F_LockBuffer(m, l0, int32(2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v200) {
		v223 = int32(0)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L25
L59:
	;
	v224 = v200 + v35
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v225 != v223&int32(255) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v208 = v200 << (uint(int32(1)) % 32)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v208)+1)))
	if v200 == int32(4081) {
		v223 = v210
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v214 = v210 & int32(255)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+(v208+int32(2))))))
	if base.Ui32(v218) < base.Ui32(v214) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v220 = v214
	goto L64
L63:
	;
	v220 = v218
	goto L64
L64:
	;
	v223 = v220
	goto L59
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v223)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v200 != 0 {
		v200 = v200 - int32(1)
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L58
L69:
	;
	goto L30
L70:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.Ui32(l1) <= base.Ui32(v256) {
		v44 = int32(1)
		goto L9
	} else {
		goto L71
	}
L71:
	;
	goto L10
}
func F_fsync_fname(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[643])))
	if v7 != 0 {
		v8 = int32(21)
	} else {
		v8 = int32(23)
	}
	v9 = F_fsync_fname_ext(m, l0, l1, int32(0), v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
