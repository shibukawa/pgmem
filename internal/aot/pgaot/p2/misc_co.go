package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ComputeXidHorizons(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int64
	_ = v534
	var v535 int32
	_ = v535
	var v539 int64
	_ = v539
	var v542 int64
	_ = v542
	var v543 int32
	_ = v543
	var v547 int64
	_ = v547
	var v550 int64
	_ = v550
	var v554 int64
	_ = v554
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v562 int64
	_ = v562
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v585 int64
	_ = v585
	var v590 int64
	_ = v590
	var v592 int64
	_ = v592
	var v594 int64
	_ = v594
	var v598 int32
	_ = v598
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])))
	if v21 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	v42 = F_LWLockAcquire(m, v38+int32(512), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[4]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	v29 = base.B2i32(v27 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])) = uint8(v29)
	v31 = v29
	goto L4
L3:
	;
	v31 = int32(0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[5]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v46
	v48 = int32(3)
	v51 = base.I32_wrap_i64(v46) + int32(1)
	if base.Ui32(v51) <= base.Ui32(v48) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v54 = v48
	goto L9
L8:
	;
	v54 = v51
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v54
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[6]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = v60
	goto L12
L11:
	;
	v61 = v54
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if int32(0) < v69 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v78 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if v31 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v78))))
	v97 = v78 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(36)+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v34)))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[7]))
	v106 = v103 + v99*int32(640)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+40))
	if v107 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L15
L18:
	;
	v195 = v78 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v195 < v196 {
		v78 = v195
		goto L16
	} else {
		goto L66
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v126 != 0 {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	if v101 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v101 == int32(0) {
		goto L18
	} else {
		goto L33
	}
L23:
	;
	v125 = v107
	goto L19
L24:
	;
	goto L25
L25:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v101))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v107)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v121 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v121 = base.B2i32(base.Ui32(v107) < base.Ui32(v101))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v121 = int32(base.Ui32(v107-v101) >> (uint(int32(31)) % 32))
	goto L26
L30:
	;
	v122 = v107
	goto L32
L31:
	;
	v122 = v101
	goto L32
L32:
	;
	v125 = v122
	goto L19
L33:
	;
	v125 = v101
	goto L19
L34:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v140 = v125
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v140
	if v95&int32(18) != 0 {
		goto L18
	} else {
		goto L44
	}
L37:
	;
	if v138 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v125))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v138 = int32(base.Ui32(v126-v125) >> (uint(int32(31)) % 32))
	goto L37
L41:
	;
	v139 = v126
	goto L43
L42:
	;
	v139 = v125
	goto L43
L43:
	;
	v140 = v139
	goto L36
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v144 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v144)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v158 = v125
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[8]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v106)+60))
	v164 = int32(0)
	if (base.B2i32(v161 == v162)|base.B2i32(v161 == v164)|int32(base.Ui32(v95)>>(uint(int32(5))%32))|v31)&int32(1) == v164 {
		goto L18
	} else {
		goto L55
	}
L48:
	;
	if v156 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v156 = base.B2i32(base.Ui32(v144) < base.Ui32(v125))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v156 = int32(base.Ui32(v144-v125) >> (uint(int32(31)) % 32))
	goto L48
L52:
	;
	v157 = v144
	goto L54
L53:
	;
	v157 = v125
	goto L54
L54:
	;
	v158 = v157
	goto L47
L55:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v175 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v175)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v189 = v125
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v189
	goto L18
L59:
	;
	if v187 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v187 = base.B2i32(base.Ui32(v175) < base.Ui32(v125))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v187 = int32(base.Ui32(v175-v125) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	v188 = v175
	goto L65
L64:
	;
	v188 = v125
	goto L65
L65:
	;
	v189 = v188
	goto L58
L66:
	;
	goto L17
L67:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v366 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v217+int32(512))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	if v226 < v225 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L67
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v293
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v309 != 0 {
		goto L99
	} else {
		goto L100
	}
L73:
	;
	if v261 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L74:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[9]))
	v231 = v226
	goto L78
L75:
	;
	goto L76
L76:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v270+int32(512))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L86
	}
L77:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v263+int32(512))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L84
	}
L78:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v229))))
	if v247 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v261 = int32(0)
	goto L77
L80:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[10]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v231<<(uint(int32(2))%32))))
	v261 = v255
	goto L77
L81:
	;
	goto L82
L82:
	;
	v257 = v231 + int32(1)
	if v257 != v225 {
		v231 = v257
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != 0 {
		goto L73
	} else {
		goto L85
	}
L85:
	;
	v293 = v261
	v294 = v261
	goto L72
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = v275
	v294 = int32(0)
	goto L72
L87:
	;
	v293 = v268
	v294 = int32(0)
	goto L72
L88:
	;
	goto L89
L89:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v261))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v268)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v290 != 0 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v290 = base.B2i32(base.Ui32(v268) < base.Ui32(v261))
	goto L90
L92:
	;
	goto L93
L93:
	;
	v290 = int32(base.Ui32(v268-v261) >> (uint(int32(31)) % 32))
	goto L90
L94:
	;
	v291 = v268
	goto L96
L95:
	;
	v291 = v261
	goto L96
L96:
	;
	v293 = v291
	v294 = v261
	goto L72
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v347
	goto L67
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v347 = v345
	goto L97
L99:
	;
	if v294 == int32(0) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	v325 = v294
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v327 == int32(0) {
		v347 = v294
		goto L97
	} else {
		goto L110
	}
L102:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v294))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v309)) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v323 != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v323 = base.B2i32(base.Ui32(v309) < base.Ui32(v294))
	goto L103
L105:
	;
	goto L106
L106:
	;
	v323 = int32(base.Ui32(v309-v294) >> (uint(int32(31)) % 32))
	goto L103
L107:
	;
	v324 = v309
	goto L109
L108:
	;
	v324 = v294
	goto L109
L109:
	;
	v325 = v324
	goto L101
L110:
	;
	if v294 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v347 = v327
	goto L97
L112:
	;
	goto L113
L113:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v294))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v327)) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v343 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v343 = base.B2i32(base.Ui32(v327) < base.Ui32(v294))
	goto L114
L116:
	;
	goto L117
L117:
	;
	v343 = int32(base.Ui32(v327-v294) >> (uint(int32(31)) % 32))
	goto L114
L118:
	;
	v344 = v327
	goto L120
L119:
	;
	v344 = v294
	goto L120
L120:
	;
	v347 = v344
	goto L97
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v461
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v466 != 0 {
		goto L175
	} else {
		goto L176
	}
L122:
	;
	if v439 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v419
	v461 = v417
	v463 = v419
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v441
	if v438 != 0 {
		goto L122
	} else {
		goto L160
	}
L125:
	;
	if v418 == int32(0) {
		goto L123
	} else {
		goto L152
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v366
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v417 = v416
	v418 = v415
	v419 = v366
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v388 == int32(0) {
		v407 = v385
		v408 = v386
		goto L139
	} else {
		goto L140
	}
L128:
	;
	v385 = v365
	v386 = v365
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v365 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v365))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v366)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v382 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v382 = base.B2i32(base.Ui32(v366) < base.Ui32(v365))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v382 = int32(base.Ui32(v366-v365) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v383 = v366
	goto L138
L137:
	;
	v383 = v365
	goto L138
L138:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v385 = v384
	v386 = v383
	goto L127
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v407
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v408 == int32(0) {
		v438 = v407
		v439 = v411
		v441 = v411
		goto L124
	} else {
		goto L151
	}
L140:
	;
	if v385 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v407 = v388
	v408 = v386
	goto L139
L142:
	;
	goto L143
L143:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v385))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v388)) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v404 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v404 = base.B2i32(base.Ui32(v388) < base.Ui32(v385))
	goto L144
L146:
	;
	goto L147
L147:
	;
	v404 = int32(base.Ui32(v388-v385) >> (uint(int32(31)) % 32))
	goto L144
L148:
	;
	v405 = v388
	goto L150
L149:
	;
	v405 = v385
	goto L150
L150:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v407 = v405
	v408 = v406
	goto L139
L151:
	;
	v417 = v407
	v418 = v411
	v419 = v408
	goto L125
L152:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v418))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v419)) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v434 != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v434 = base.B2i32(base.Ui32(v419) < base.Ui32(v418))
	goto L153
L155:
	;
	goto L156
L156:
	;
	v434 = int32(base.Ui32(v419-v418) >> (uint(int32(31)) % 32))
	goto L153
L157:
	;
	v435 = v419
	goto L159
L158:
	;
	v435 = v418
	goto L159
L159:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v438 = v437
	v439 = v436
	v441 = v435
	goto L124
L160:
	;
	v461 = v439
	v463 = v441
	goto L121
L161:
	;
	v461 = v438
	v463 = v441
	goto L121
L162:
	;
	goto L163
L163:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v439))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v438)) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v458 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v458 = base.B2i32(base.Ui32(v438) < base.Ui32(v439))
	goto L164
L166:
	;
	goto L167
L167:
	;
	v458 = int32(base.Ui32(v438-v439) >> (uint(int32(31)) % 32))
	goto L164
L168:
	;
	v459 = v438
	goto L170
L169:
	;
	v459 = v439
	goto L170
L170:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v461 = v459
	v463 = v460
	goto L121
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v530
	v534 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v539 = v534 + base.I64_extend_i32_s(v535-base.I32_wrap_i64(v534))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[11])) = v539
	v542 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v547 = v542 + base.I64_extend_i32_s(v543-base.I32_wrap_i64(v542))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[12])) = v547
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v554 = v550 + base.I64_extend_i32_s(v531-base.I32_wrap_i64(v550))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[13])) = v554
	v557 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v562 = v557 + base.I64_extend_i32_s(v558-base.I32_wrap_i64(v557))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[14])) = v562
	v564 = int32(_a_F_ComputeXidHorizons_0)
	v565 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15]))
	if base.Ui64(v565) < base.Ui64(v539) {
		goto L210
	} else {
		goto L211
	}
L172:
	;
	if v511 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v461
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v461 == int32(0) {
		v530 = v506
		v531 = v506
		goto L171
	} else {
		goto L199
	}
L174:
	;
	if v487 != 0 {
		goto L189
	} else {
		goto L190
	}
L175:
	;
	if v463 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v463
	if v463 == int32(0) {
		goto L173
	} else {
		goto L188
	}
L178:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v463))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v466)) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v481 = v461
	v482 = v466
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v482
	v487 = v481
	v488 = v482
	goto L174
L181:
	;
	if v478 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v478 = base.B2i32(base.Ui32(v466) < base.Ui32(v463))
	goto L181
L183:
	;
	goto L184
L184:
	;
	v478 = int32(base.Ui32(v466-v463) >> (uint(int32(31)) % 32))
	goto L181
L185:
	;
	v479 = v466
	goto L187
L186:
	;
	v479 = v463
	goto L187
L187:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v481 = v480
	v482 = v479
	goto L180
L188:
	;
	v487 = v461
	v488 = v463
	goto L174
L189:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v488)) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v502 = v488
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v502
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v509 = v502
	v511 = v504
	goto L172
L192:
	;
	if v500 != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v500 = base.B2i32(base.Ui32(v488) < base.Ui32(v487))
	goto L192
L194:
	;
	goto L195
L195:
	;
	v500 = int32(base.Ui32(v488-v487) >> (uint(int32(31)) % 32))
	goto L192
L196:
	;
	v501 = v488
	goto L198
L197:
	;
	v501 = v487
	goto L198
L198:
	;
	v502 = v501
	goto L191
L199:
	;
	v509 = v461
	v511 = v506
	goto L172
L200:
	;
	v530 = v509
	v531 = int32(0)
	goto L171
L201:
	;
	goto L202
L202:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v511))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v509)) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v526 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v526 = base.B2i32(base.Ui32(v509) < base.Ui32(v511))
	goto L203
L205:
	;
	goto L206
L206:
	;
	v526 = int32(base.Ui32(v509-v511) >> (uint(int32(31)) % 32))
	goto L203
L207:
	;
	v527 = v509
	goto L209
L208:
	;
	v527 = v511
	goto L209
L209:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v530 = v527
	v531 = v528
	goto L171
L210:
	;
	v568 = v539
	goto L212
L211:
	;
	v568 = v565
	goto L212
L212:
	;
	if base.I32_wrap_i64(v565) != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v570 = v568
	goto L215
L214:
	;
	v570 = v539
	goto L215
L215:
	;
	if base.I32_wrap_i64(v539) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v572 = v570
	goto L218
L217:
	;
	v572 = v565
	goto L218
L218:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15])) = v572
	v574 = int32(_a_F_ComputeXidHorizons_1)
	v575 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16]))
	if base.Ui64(v575) < base.Ui64(v547) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v578 = v547
	goto L221
L220:
	;
	v578 = v575
	goto L221
L221:
	;
	if base.I32_wrap_i64(v575) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v580 = v578
	goto L224
L223:
	;
	v580 = v547
	goto L224
L224:
	;
	if base.I32_wrap_i64(v547) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v582 = v580
	goto L227
L226:
	;
	v582 = v575
	goto L227
L227:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16])) = v582
	v584 = int32(_a_F_ComputeXidHorizons_2)
	v585 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[18])) = v562
	if base.Ui64(v585) < base.Ui64(v554) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v590 = v554
	goto L230
L229:
	;
	v590 = v585
	goto L230
L230:
	;
	if base.I32_wrap_i64(v585) != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v592 = v590
	goto L233
L232:
	;
	v592 = v554
	goto L233
L233:
	;
	if base.I32_wrap_i64(v554) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v594 = v592
	goto L236
L235:
	;
	v594 = v585
	goto L236
L236:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17])) = v594
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[20])) = v598
	return
}
func F_ConditionVariableCancelSleep(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0]))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		if v8 != 0 {
			F_s_lock(m, v7, int32(_a_F_ConditionVariableCancelSleep_0), int32(238), int32(_a_F_ConditionVariableCancelSleep_1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
				v23 = v18 + v20*int32(640)
				v25 = v23 + int32(84)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
				if v26 == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v29 != 0 {
						v35 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
						v40 = v26
						v41 = v35
						if v41 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
						}
						*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
					} else {
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v26 != int32(-1) {
						v35 = v30
						*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
						v40 = v26
						v41 = v35
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
						v40 = v34
						v41 = v30
					}
					if v41 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
				}
				v56 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v56
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
			v23 = v18 + v20*int32(640)
			v25 = v23 + int32(84)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				if v29 != 0 {
					v35 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
					v40 = v26
					v41 = v35
					if v41 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
				} else {
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				if v26 != int32(-1) {
					v35 = v30
					*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
					v40 = v26
					v41 = v35
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v40 = v34
					v41 = v30
				}
				if v41 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
				}
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
			}
			v56 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v56
			return
		}
	} else {
		return
	}
}
func F_ConditionalXactLockTableWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v18 = F_LockAcquireExtended(m, v7, int32(5), v3, int32(1), v3, l1)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v87
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		v87 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v87 = int32(1)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v33
	v40 = int32(0)
	v45 = F_LockAcquireExtended(m, v7, int32(5), v40, int32(1), v40, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v45 == int32(0) {
		v87 = v40
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = v33
	goto L11
L11:
	;
	v55 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v87 = v57 ^ int32(1)
	goto L1
L13:
	;
	goto L12
L14:
	;
	v57 = F_TransactionIdIsInProgress(m, v49)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v57 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalXactLockTableWait[0]))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v68 = F_SubTransGetTopmostTransaction(m, v49)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v68
	v76 = int32(0)
	v79 = F_LockAcquireExtended(m, v7, int32(5), v76, int32(1), v76, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v79 != 0 {
		v49 = v68
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L13
}
func F_CountOtherDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[0]))
	v31 = int32(0)
	goto L1
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[1]))
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v16 + int32(48)
	return v188
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	v50 = F_LWLockAcquire(m, v46+int32(512), int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L5
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v52 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v57+int32(512))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v62 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[3]))
	v70 = v62
	v71 = v62
	v72 = v62
	goto L14
L13:
	;
	v188 = int32(0)
	goto L9
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(36)+v70<<(uint(int32(2))%32))))
	v86 = v64 + v83*int32(640)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+60))
	if v87 != l0 {
		v123 = v71
		v124 = v72
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v131+int32(512))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L25
	}
L16:
	;
	v127 = v70 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v127 < v128 {
		v70 = v127
		v71 = v123
		v72 = v124
		goto L14
	} else {
		goto L24
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[4]))
	if v86 == v90 {
		v123 = v71
		v124 = v72
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v92 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v95 = int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96 + v95
	v123 = v95
	v124 = v72
	goto L16
L20:
	;
	goto L21
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[5]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v70))))
	v105 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v106 + v105
	if v104&v105 == int32(0) {
		v123 = v105
		v124 = v72
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if int32(9) < v72 {
		v123 = v105
		v124 = v72
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v72<<(uint(int32(2))%32)))) = v119
	v123 = v105
	v124 = v72 + int32(1)
	goto L16
L24:
	;
	goto L15
L25:
	;
	if v123 == int32(0) {
		v188 = v123
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v138 = int32(0)
	if v138 < v124 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v144 = v138
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pg_usleep(m, int32(_a_F_CountOtherDBBackends_0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L34
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v16+v144<<(uint(int32(2))%32))))
	v159 = F_kill(m, v157, int32(15))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v162 = v144 + int32(1)
	if v162 != v124 {
		v144 = v162
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v181 = v31 + int32(1)
	if v181 != int32(50) {
		v31 = v181
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v188 = v123
	goto L9
}
func F_coerce_null_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v15 = F_getBaseTypeAndTypmod(m, l0, v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v22 = F_makeConst(m, v15, v19, l2, l3, int32(0), int32(1), l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if l0 != v15 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v26 = int32(0)
				v30 = F_coerce_to_domain(m, v22, v15, v25, l0, v26, int32(2), int32(-1), v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = v30
					m.G0 = v10 + int32(16)
					return v32
				}
			} else {
				v32 = v22
				m.G0 = v10 + int32(16)
				return v32
			}
		}
	}
}
func F_colorTrgmInfoPenaltyCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F32_ne(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F32_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_combo_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt_len(m *base.Module, l0 int32, l1 int32) int32 {
	return l1 + int32(512)
}
func F_comp_ptrgm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_comp_ptrgm[0]))
	v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v17 = v6
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v17 = base.B2i32(v11 < v10) - base.B2i32(v10 < v11)
		}
		return v17
	}
}
func F_comparePairs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v85
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32(v6) < base.Ui32(v5) {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	if v71 != 0 {
		v85 = v71
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v71 = int32(0)
	goto L5
L7:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L17
L8:
	;
	if (v8|v9)&int32(3) != 0 {
		v40 = v8
		v41 = v9
		v42 = v5
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v33 = v8
	v34 = v9
	v35 = v5
	goto L10
L10:
	;
	if v35 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v17 = v8
	v18 = v9
	v19 = v5
	goto L12
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L10
L14:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L7
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = v50 - v51
	goto L5
L19:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L6
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v73 == v74 {
		v85 = int32(0)
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(1)
	goto L27
L26:
	;
	v78 = int32(-1)
	goto L27
L27:
	;
	return v78
L28:
	;
	v83 = int32(1)
	goto L30
L29:
	;
	v83 = int32(-1)
	goto L30
L30:
	;
	v85 = v83
	goto L1
}
func F_compare_distances(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F64_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_compare_values(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v22 = int32(-1)
			return v22
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_FunctionCall2Coll(m, v14, v15, v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = base.B2i32(v18 != int32(0))
				return v22
			}
		}
	}
}
func F_compatible_oper_opid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_oper(m, v4, l0, l1, l2, v4, int32(-1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			v58 = v4
			m.G0 = v10 + int32(16)
			if v58 != 0 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
				F_ReleaseCatCache(m, v58)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v68 = v65
					return v68
				}
			} else {
				v68 = v4
				return v68
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
			v23 = v21 + v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
			v25 = F_IsBinaryCoercible(m, l1, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = F_op_signature_string(m, l0, l1, l2)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
									F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, int32(0), int32(-1))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(475), int32(_a_F_compatible_oper_opid_2))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					v30 = F_IsBinaryCoercible(m, l2, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							F_ReleaseCatCache(m, v15)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = F_op_signature_string(m, l0, l1, l2)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
											F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, int32(0), int32(-1))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(475), int32(_a_F_compatible_oper_opid_2))
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
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
								}
							}
						} else {
							v58 = v15
							m.G0 = v10 + int32(16)
							if v58 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
								F_ReleaseCatCache(m, v58)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = v65
									return v68
								}
							} else {
								v68 = v4
								return v68
							}
						}
					}
				}
			}
		}
	}
}
func F_composite_to_json(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v21 = F_lookup_rowtype_tupdesc(m, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(base.Ui32(v23) >> (uint(int32(2)) % 32))
	F_appendStringInfoChar(m, l1, int32(123))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) < v31 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoChar(m, l1, int32(125))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v36 = int32(3)
	goto L10
L9:
	;
	v36 = int32(1)
	goto L10
L10:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(_a_F_composite_to_json_0)
	goto L13
L12:
	;
	v39 = int32(_a_F_composite_to_json_1)
	goto L13
L13:
	;
	v41 = v21 + int32(20)
	v43 = int32(0)
	v47 = v31
	v49 = int32(0)
	goto L14
L14:
	;
	v60 = v41 + v47<<(uint(int32(4))%32) + v43*int32(100)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+91)))
	if v61 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L7
L16:
	;
	v186 = v47
	v187 = v49
	v190 = v43 + int32(1)
	goto L18
L17:
	;
	if v49&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v190 < v186 {
		v43 = v190
		v47 = v186
		v49 = v187
		goto L14
	} else {
		goto L55
	}
L19:
	;
	F_appendBinaryStringInfo(m, l1, v39, v36)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_escape_json(m, l1, v60+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	F_appendStringInfoChar(m, l1, int32(58))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v78 = v43 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+18)))
	if base.Ui32(v80&int32(2047)) <= base.Ui32(v43) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	if v153 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v86 = F_getmissingattr(m, v21, v78, v15+int32(27))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v88)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)))
	if v90&int32(1) == v88 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v152 = v86
	goto L25
L30:
	;
	v97 = v41 + v43<<(uint(int32(4))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if int32(0) <= v98 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+int32(base.Ui32(v43)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v134)>>(uint(v43&int32(7))%32))&int32(1) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	v103 = v79 + v101 + v98
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	if v104 != int32(1) {
		v152 = v103
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v129 = F_nocachegetattr(m, v15+int32(28), v78, v21)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L36:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	switch v107 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L39
	default:
		goto L37
	case 3:
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v152 = v112
	goto L25
L39:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103))))
	v152 = v111
	goto L25
L40:
	;
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v103))))
	v152 = v110
	goto L25
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = base.I32_extend16_s(v107)
	F_errmsg_internal(m, int32(_a_F_composite_to_json_2), v15)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_composite_to_json_3), int32(70), int32(_a_F_composite_to_json_4))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v152 = v129
	goto L25
L45:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v142)
	v152 = int32(0)
	goto L25
L46:
	;
	goto L47
L47:
	;
	v147 = F_nocachegetattr(m, v15+int32(28), v78, v21)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v152 = v147
	goto L25
L49:
	;
	v177 = int32(1)
	F_datum_to_json_internal(m, v152, v176&v177, l1, v175, v174, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L50:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v156
	v174 = v156
	v175 = v156
	v176 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v60)+68))
	F_json_categorize_type(m, v163, int32(0), v15+int32(20), v15+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	v174 = v171
	v175 = v172
	v176 = v173
	goto L49
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v186 = v183
	v187 = v177
	v190 = v78
	goto L18
L55:
	;
	goto L15
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if int32(0) <= v207 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_DecrTupleDescRefCount(m, v21)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	m.G0 = v15 + int32(48)
	return
L60:
	;
	goto L59
}
func F_compress_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.Env.Pgmem_deflate_write(m, v7, l2, l3)
	mBase = m.M
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-105)
L2:
	;
	goto L3
L3:
	;
	v14 = l1 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = m.Env.Pgmem_zstream_read(m, v22, v14, v15)
	mBase = m.M
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v32
L6:
	;
	return int32(-105)
L7:
	;
	goto L8
L8:
	;
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = F_pushf_write(m, l0, v14, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if int32(0) <= v32 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_computeDistance(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v27 float64
	_ = v27
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v84 float64
	_ = v84
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v133 float64
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v167 int32
	_ = v167
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 != 0 {
		v17 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return float64(0)
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
			v156 = v21
			m.G0 = v11 + int32(16)
			return v156
		}
	} else {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		if base.F64_le(v22, v23) == int32(0) {
			v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v75 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_le(v74, v75) == int32(0) {
				v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return float64(0)
				} else {
					v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
					v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return float64(0)
					} else {
						v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
						v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
						v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
						v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return float64(0)
						} else {
							v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
							v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
							v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
							v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return float64(0)
							} else {
								v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
								if base.F64_lt(v130, v125) != 0 {
									v150 = v130
								} else {
									v150 = v125
								}
								if base.F64_gt(v150, v139) != 0 {
									v152 = v139
								} else {
									v152 = v150
								}
								if base.F64_gt(v152, v148) != 0 {
									v154 = v148
								} else {
									v154 = v152
								}
								v156 = v154
								m.G0 = v11 + int32(16)
								return v156
							}
						}
					}
				}
			} else {
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.F64_ge(v74, v79) == int32(0) {
					v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return float64(0)
					} else {
						v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
						v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return float64(0)
						} else {
							v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
							v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
							v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
							v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return float64(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
								v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
								v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return float64(0)
								} else {
									v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
									if base.F64_lt(v130, v125) != 0 {
										v150 = v130
									} else {
										v150 = v125
									}
									if base.F64_gt(v150, v139) != 0 {
										v152 = v139
									} else {
										v152 = v150
									}
									if base.F64_gt(v152, v148) != 0 {
										v154 = v148
									} else {
										v154 = v152
									}
									v156 = v154
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					}
				} else {
					if base.F64_gt(v22, v23) != 0 {
						v84 = base.F64_sub(v22, v23)
						if base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v156 = v84
							m.G0 = v11 + int32(16)
							return v156
						} else {
							if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v84
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v156 = v84
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					} else {
						v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						if base.F64_gt(v94, v22) != 0 {
							v96 = base.F64_sub(v94, v22)
							if base.F64_ne(base.F64_abs(v96), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v96
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v96
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v94), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v96
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1256), int32(_a_F_computeDistance_2))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return float64(0)
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
			}
		} else {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
			if base.F64_ge(v22, v27) == int32(0) {
				v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v75 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v74, v75) == int32(0) {
					v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return float64(0)
					} else {
						v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
						v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return float64(0)
						} else {
							v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
							v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
							v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
							v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return float64(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
								v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
								v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return float64(0)
								} else {
									v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
									if base.F64_lt(v130, v125) != 0 {
										v150 = v130
									} else {
										v150 = v125
									}
									if base.F64_gt(v150, v139) != 0 {
										v152 = v139
									} else {
										v152 = v150
									}
									if base.F64_gt(v152, v148) != 0 {
										v154 = v148
									} else {
										v154 = v152
									}
									v156 = v154
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					}
				} else {
					v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v74, v79) == int32(0) {
						v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return float64(0)
						} else {
							v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
							v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return float64(0)
							} else {
								v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
								v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
								v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
								v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return float64(0)
								} else {
									v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
									*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
									v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
									v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return float64(0)
									} else {
										v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
										if base.F64_lt(v130, v125) != 0 {
											v150 = v130
										} else {
											v150 = v125
										}
										if base.F64_gt(v150, v139) != 0 {
											v152 = v139
										} else {
											v152 = v150
										}
										if base.F64_gt(v152, v148) != 0 {
											v154 = v148
										} else {
											v154 = v152
										}
										v156 = v154
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						}
					} else {
						if base.F64_gt(v22, v23) != 0 {
							v84 = base.F64_sub(v22, v23)
							if base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v84
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v84
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v84
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							if base.F64_gt(v94, v22) != 0 {
								v96 = base.F64_sub(v94, v22)
								if base.F64_ne(base.F64_abs(v96), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v96
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v96
										m.G0 = v11 + int32(16)
										return v156
									} else {
										if base.F64_ne(base.F64_abs(v94), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v156 = v96
											m.G0 = v11 + int32(16)
											return v156
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1256), int32(_a_F_computeDistance_2))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return float64(0)
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
				}
			} else {
				v31 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v32 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v31, v32) != 0 {
					v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v31, v34) != 0 {
						v156 = float64(0)
						m.G0 = v11 + int32(16)
						return v156
					} else {
						v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
						v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						if base.F64_gt(v36, v37) != 0 {
							v39 = base.F64_sub(v36, v37)
							if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v39
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v39
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v39
										m.G0 = v11 + int32(16)
										return v156
									} else {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							if base.F64_gt(v49, v36) != 0 {
								v51 = base.F64_sub(v49, v36)
								if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v51
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v51
										m.G0 = v11 + int32(16)
										return v156
									} else {
										if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v156 = v51
											m.G0 = v11 + int32(16)
											return v156
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1245), int32(_a_F_computeDistance_2))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return float64(0)
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
				} else {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
					v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					if base.F64_gt(v36, v37) != 0 {
						v39 = base.F64_sub(v36, v37)
						if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v156 = v39
							m.G0 = v11 + int32(16)
							return v156
						} else {
							if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v39
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v39
									m.G0 = v11 + int32(16)
									return v156
								} else {
									F_float_overflow_error(m)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						if base.F64_gt(v49, v36) != 0 {
							v51 = base.F64_sub(v49, v36)
							if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v51
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v51
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v51
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1245), int32(_a_F_computeDistance_2))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return float64(0)
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
			}
		}
	}
}
func F_computeLeafRecompressWALData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v2 {
		v34 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = F_palloc(m, v34<<(uint(int32(1))%32)+int32(_a_F_computeLeafRecompressWALData_0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	if l0 == v14 {
		v34 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = v2
	v20 = v14
	goto L4
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
	v30 = v19 + base.B2i32(v27 != int32(0))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v31 != l0 {
		v19 = v30
		v20 = v31
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v34 = v30
	goto L1
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v34)
	v50 = v46 + int32(2)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 == int32(0) {
		v151 = v50
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v151 - v46
	m.G0 = v12 + int32(16)
	return
L10:
	;
	if l0 == v51 {
		v151 = v50
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v56 = v50
	v57 = v51
	v59 = v2
	goto L12
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+8)))
	switch v64 {
	case 0:
		goto L18
	case 1:
		goto L16
	default:
		goto L17
	}
L13:
	;
	v151 = v143
	goto L9
L14:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v148 != l0 {
		v56 = v143
		v57 = v148
		v59 = v147
		goto L12
	} else {
		goto L39
	}
L15:
	;
	v143 = v138 + v56 + int32(2)
	v147 = v135 + v59
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v131)
	v135 = v131
	v138 = int32(0)
	goto L15
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)))
	v72 = (v68 + int32(1)) & int32(_a_F_computeLeafRecompressWALData_1)
	v74 = v72 + int32(8)
	if v64 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v143 = v56
	v147 = v59 + int32(1)
	goto L14
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L36
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v74 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	if base.Ui32(v77*int32(6)) <= base.Ui32(v74) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v64)
	if v64&int32(254) != int32(2) {
		goto L19
	} else {
		goto L31
	}
L24:
	;
	v83 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v83)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+2)) = uint16(v85)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v91 = v85 * int32(6)
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v96 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v96)
	v105 = v96
	goto L20
L27:
	;
	v135 = int32(1)
	v138 = v91 + int32(2)
	goto L15
L28:
	;
	v92 = F__emscripten_memcpy_bulkmem(m, v56+v83, v89, v91)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v105 = v64
	goto L20
L32:
	;
	v135 = base.B2i32(v105 != int32(2))
	v138 = (v72 + int32(9)) & int32(_a_F_computeLeafRecompressWALData_2)
	goto L15
L33:
	;
	v109 = F__emscripten_memcpy_bulkmem(m, v56+int32(2), v108, v74)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v64
	F_errmsg_internal(m, int32(_a_F_computeLeafRecompressWALData_3), v12)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_computeLeafRecompressWALData_4), int32(955), int32(_a_F_computeLeafRecompressWALData_5))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	goto L13
}
func F_connect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_connect(m, l0, int32(_a_F_connect_0), int32(12), v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_connect[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_consider_index_join_outer_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v329 int32
	_ = v329
	var v351 int32
	_ = v351
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	v11 = int32(0)
	if l7 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v11
	goto L4
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v37<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v55 = F_list_member(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v55 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v59 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v372 = v37 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v372 < v373 {
		v37 = v372
		goto L4
	} else {
		goto L86
	}
L11:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v54, l9)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L85
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v62 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v78 = int32(0)
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v78<<(uint(int32(2))%32))))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L11
L16:
	;
	v329 = v78 + int32(1)
	if v329 != v62 {
		v78 = v329
		goto L14
	} else {
		goto L84
	}
L17:
	;
	if v185 != int32(3) {
		goto L16
	} else {
		goto L53
	}
L18:
	;
	v185 = base.B2i32(v90 != int32(0))
	goto L17
L19:
	;
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
	v185 = int32(2)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v107 < v108 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v110 = v107
	goto L26
L25:
	;
	v110 = v108
	goto L26
L26:
	;
	if v110 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = int32(1)
	goto L29
L28:
	;
	v113 = v110
	goto L29
L29:
	;
	v114 = int32(8)
	v118 = int32(0)
	v120 = v118
	v121 = v118
	goto L32
L30:
	;
	v185 = int32(3)
	goto L17
L31:
	;
	v185 = v171
	goto L17
L32:
	;
	v131 = v121 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v54+v114+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v90+v114))))
	if v133&(v135^int32(-1)) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v108 < v107 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v157 = v121 + int32(1)
	if v157 != v113 {
		v120 = v154
		v121 = v157
		goto L32
	} else {
		goto L42
	}
L35:
	;
	v141 = int32(3)
	if v120 == int32(1) {
		v171 = v141
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v135&(v133^int32(-1)) == int32(0) {
		v154 = v120
		goto L34
	} else {
		goto L40
	}
L38:
	;
	if v135&(v133^int32(-1)) != 0 {
		v171 = v141
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v154 = int32(2)
	goto L34
L40:
	;
	if v120 == int32(2) {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	v154 = int32(1)
	goto L34
L42:
	;
	goto L33
L43:
	;
	if v154 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v108 <= v107 {
		v171 = v154
		goto L31
	} else {
		goto L49
	}
L46:
	;
	v164 = int32(3)
	goto L48
L47:
	;
	v164 = int32(2)
	goto L48
L48:
	;
	v185 = v164
	goto L17
L49:
	;
	if v154 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = int32(3)
	goto L52
L51:
	;
	v170 = int32(1)
	goto L52
L52:
	;
	v171 = v170
	goto L31
L53:
	;
	if v52 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v300 != 0 {
		goto L78
	} else {
		goto L79
	}
L55:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v191 <= v190 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v202 = v190
	v209 = v191
	goto L57
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v202<<(uint(int32(2))%32))))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+60))
	if v52 == v219 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L54
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	v222 = int32(0)
	if v221 == v222 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v277 = v209
	goto L61
L61:
	;
	v279 = v202 + int32(1)
	if v279 < v277 {
		v202 = v279
		v209 = v277
		goto L57
	} else {
		goto L77
	}
L62:
	;
	if v275 != 0 {
		goto L16
	} else {
		goto L76
	}
L63:
	;
	v275 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v90 == int32(0) {
		v266 = v222
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v275 = v266
	goto L62
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v232 < v231 {
		v266 = v222
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v237 = v234
	goto L71
L70:
	;
	v237 = v231
	goto L71
L71:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L72
L72:
	;
	v250 = v243 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v221+v238+v250)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+(v90+v238))))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v266 = v259
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v266 = v259
	goto L66
L74:
	;
	v261 = v243 + int32(1)
	if v261 != v237 {
		v243 = v261
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v277 = v276
	goto L61
L77:
	;
	goto L58
L78:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v303 = v301
	goto L80
L79:
	;
	v303 = int32(0)
	goto L80
L80:
	;
	if l8*int32(10) <= v303 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v305 = F_bms_union(m, v54, v90)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v305, l9)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	goto L16
L84:
	;
	goto L15
L85:
	;
	goto L10
L86:
	;
	goto L5
}
func F_conv_utf8_to(m *base.Module, l0 int32) int32 {
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		v45 = l0
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_conv_utf8_to_0)) {
			v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(1984) | l0&int32(63)
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(16777215)) {
				v45 = int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_1) | (int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | l0&int32(63))
			} else {
				v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | (int32(base.Ui32(l0)>>(uint(int32(6))%32))&int32(_a_F_conv_utf8_to_2) | (int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_3) | l0&int32(63)))
			}
		}
	}
	if base.Ui32(v45-int32(1106)) <= base.Ui32(int32(_a_F_conv_utf8_to_4)) {
		v51 = v45 - int32(286)
		v52 = int32(_a_F_conv_utf8_to_0)
		v53 = v51 & v52
		v55 = base.I32_div_u_s(v53, int32(1260))
		v58 = int32(10)
		v59 = base.I32_div_u_s(v53, v58)
		v61 = base.I32_rem_u_s(v59, int32(126))
		return v55<<(uint(int32(16))%32) | (v61<<(uint(int32(8))%32) + (v51-v59*v58)&v52 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
	} else {
		if base.Ui32(v45-int32(_a_F_conv_utf8_to_6)) <= base.Ui32(int32(2109)) {
			v81 = v45 - int32(576)
			v82 = int32(_a_F_conv_utf8_to_0)
			v83 = v81 & v82
			v85 = base.I32_div_u_s(v83, int32(1260))
			v88 = int32(10)
			v89 = base.I32_div_u_s(v83, v88)
			v91 = base.I32_rem_u_s(v89, int32(126))
			return v85<<(uint(int32(16))%32) | (v91<<(uint(int32(8))%32) + (v81-v89*v88)&v82 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
		} else {
			if base.Ui32(v45-int32(_a_F_conv_utf8_to_7)) <= base.Ui32(int32(764)) {
				v111 = v45 - int32(878)
				v112 = int32(_a_F_conv_utf8_to_0)
				v113 = v111 & v112
				v115 = base.I32_div_u_s(v113, int32(1260))
				v120 = int32(10)
				v121 = base.I32_div_u_s(v113, v120)
				v123 = base.I32_rem_u_s(v121, int32(126))
				return v115<<(uint(int32(16))%32) + int32(2146828288) | (v123<<(uint(int32(8))%32) + (v111-v121*v120)&v112 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
			} else {
				if base.Ui32(v45-int32(_a_F_conv_utf8_to_8)) <= base.Ui32(int32(884)) {
					v143 = v45 - int32(887)
					v144 = int32(_a_F_conv_utf8_to_0)
					v145 = v143 & v144
					v147 = base.I32_div_u_s(v145, int32(1260))
					v152 = int32(10)
					v153 = base.I32_div_u_s(v145, v152)
					v155 = base.I32_rem_u_s(v153, int32(126))
					return v147<<(uint(int32(16))%32) + int32(2146828288) | (v155<<(uint(int32(8))%32) + (v143-v153*v152)&v144 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
				} else {
					if base.Ui32(v45-int32(_a_F_conv_utf8_to_9)) <= base.Ui32(int32(470)) {
						v175 = v45 - int32(889)
						v176 = int32(_a_F_conv_utf8_to_0)
						v177 = v175 & v176
						v179 = base.I32_div_u_s(v177, int32(1260))
						v184 = int32(10)
						v185 = base.I32_div_u_s(v177, v184)
						v187 = base.I32_rem_u_s(v185, int32(126))
						return v179<<(uint(int32(16))%32) + int32(2146828288) | (v187<<(uint(int32(8))%32) + (v175-v185*v184)&v176 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
					} else {
						if base.Ui32(v45-int32(_a_F_conv_utf8_to_10)) <= base.Ui32(int32(372)) {
							v207 = v45 - int32(894)
							v208 = int32(_a_F_conv_utf8_to_0)
							v209 = v207 & v208
							v211 = base.I32_div_u_s(v209, int32(1260))
							v216 = int32(10)
							v217 = base.I32_div_u_s(v209, v216)
							v219 = base.I32_rem_u_s(v217, int32(126))
							return v211<<(uint(int32(16))%32) + int32(2146828288) | (v219<<(uint(int32(8))%32) + (v207-v217*v216)&v208 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
						} else {
							if base.Ui32(v45-int32(_a_F_conv_utf8_to_11)) <= base.Ui32(int32(440)) {
								v239 = v45 - int32(900)
								v240 = int32(_a_F_conv_utf8_to_0)
								v241 = v239 & v240
								v243 = base.I32_div_u_s(v241, int32(1260))
								v248 = int32(10)
								v249 = base.I32_div_u_s(v241, v248)
								v251 = base.I32_rem_u_s(v249, int32(126))
								return v243<<(uint(int32(16))%32) + int32(2146828288) | (v251<<(uint(int32(8))%32) + (v239-v249*v248)&v240 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
							} else {
								if base.Ui32(v45-int32(_a_F_conv_utf8_to_12)) <= base.Ui32(int32(702)) {
									v271 = v45 - int32(911)
									v272 = int32(_a_F_conv_utf8_to_0)
									v273 = v271 & v272
									v275 = base.I32_div_u_s(v273, int32(1260))
									v280 = int32(10)
									v281 = base.I32_div_u_s(v273, v280)
									v283 = base.I32_rem_u_s(v281, int32(126))
									return v275<<(uint(int32(16))%32) + int32(2146828288) | (v283<<(uint(int32(8))%32) + (v271-v281*v280)&v272 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
								} else {
									if base.Ui32(v45-int32(_a_F_conv_utf8_to_13)) <= base.Ui32(int32(_a_F_conv_utf8_to_14)) {
										v303 = v45 - int32(_a_F_conv_utf8_to_15)
										v304 = int32(_a_F_conv_utf8_to_0)
										v305 = v303 & v304
										v307 = base.I32_div_u_s(v305, int32(_a_F_conv_utf8_to_16))
										v311 = base.I32_div_u_s(v305, int32(1260))
										v312 = int32(10)
										v313 = base.I32_rem_u_s(v311, v312)
										v320 = base.I32_div_u_s(v305, v312)
										v322 = base.I32_rem_u_s(v320, int32(126))
										return v307<<(uint(int32(24))%32) | v313<<(uint(int32(16))%32) - int32(2130706432) | (v322<<(uint(int32(8))%32) + (v303-v320*v312)&v304 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
									} else {
										if base.Ui32(v45-int32(_a_F_conv_utf8_to_18)) <= base.Ui32(int32(_a_F_conv_utf8_to_19)) {
											v342 = v45 - int32(_a_F_conv_utf8_to_20)
											v343 = int32(_a_F_conv_utf8_to_0)
											v344 = v342 & v343
											v346 = base.I32_div_u_s(v344, int32(_a_F_conv_utf8_to_16))
											v350 = base.I32_div_u_s(v344, int32(1260))
											v351 = int32(10)
											v352 = base.I32_rem_u_s(v350, v351)
											v359 = base.I32_div_u_s(v344, v351)
											v361 = base.I32_rem_u_s(v359, int32(126))
											return v346<<(uint(int32(24))%32) | v352<<(uint(int32(16))%32) - int32(2130706432) | (v361<<(uint(int32(8))%32) + (v342-v359*v351)&v343 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
										} else {
											if base.Ui32(v45-int32(_a_F_conv_utf8_to_21)) <= base.Ui32(int32(1029)) {
												v381 = v45 - int32(_a_F_conv_utf8_to_22)
												v382 = int32(_a_F_conv_utf8_to_0)
												v383 = v381 & v382
												v385 = base.I32_div_u_s(v383, int32(_a_F_conv_utf8_to_16))
												v389 = base.I32_div_u_s(v383, int32(1260))
												v390 = int32(10)
												v391 = base.I32_rem_u_s(v389, v390)
												v398 = base.I32_div_u_s(v383, v390)
												v400 = base.I32_rem_u_s(v398, int32(126))
												return v385<<(uint(int32(24))%32) | v391<<(uint(int32(16))%32) - int32(2130706432) | (v400<<(uint(int32(8))%32) + (v381-v398*v390)&v382 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
											} else {
												if base.Ui32(v45-int32(_a_F_conv_utf8_to_23)) <= base.Ui32(int32(25)) {
													v420 = v45 - int32(_a_F_conv_utf8_to_24)
													v421 = int32(_a_F_conv_utf8_to_0)
													v422 = v420 & v421
													v424 = base.I32_div_u_s(v422, int32(_a_F_conv_utf8_to_16))
													v428 = base.I32_div_u_s(v422, int32(1260))
													v429 = int32(10)
													v430 = base.I32_rem_u_s(v428, v429)
													v437 = base.I32_div_u_s(v422, v429)
													v439 = base.I32_rem_u_s(v437, int32(126))
													return v424<<(uint(int32(24))%32) | v430<<(uint(int32(16))%32) - int32(2130706432) | (v439<<(uint(int32(8))%32) + (v420-v437*v429)&v421 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
												} else {
													if base.Ui32(v45-int32(_a_F_conv_utf8_to_25)) <= base.Ui32(int32(_a_F_conv_utf8_to_26)) {
														v459 = v45 + int32(_a_F_conv_utf8_to_27)
														v460 = int32(10)
														v461 = base.I32_div_u_s(v459, v460)
														v463 = base.I32_rem_u_s(v461, int32(126))
														v473 = base.I32_div_u_s(v459, int32(_a_F_conv_utf8_to_16))
														v477 = base.I32_div_u_s(v459, int32(1260))
														v481 = base.I32_rem_u_s(v477&int32(_a_F_conv_utf8_to_0), v460)
														v492 = v463<<(uint(int32(8))%32) + (v459 - v461*v460) + int32(_a_F_conv_utf8_to_5) | (v473<<(uint(int32(24))%32) | v481<<(uint(int32(16))%32) - int32(2130706432)) | int32(_a_F_conv_utf8_to_17)
													} else {
														v492 = int32(0)
													}
													return v492
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
func F_convert_EXISTS_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v294 int32
	_ = v294
	v5 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	if v11 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v294
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_copyObjectImpl(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v17 = F_simplify_EXISTS_query(m, l0, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v17 == int32(0) {
		v294 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
	v26 = F_contain_vars_of_level(m, v13, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v26 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = F_contain_vars_of_level(m, v22, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v29 == int32(0) {
		v294 = v5
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v33 = F_contain_volatile_functions(m, v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v33 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_replace_empty_jointree(m, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = v39
	goto L16
L15:
	;
	v40 = int32(0)
	goto L16
L16:
	;
	F_OffsetVarNodes(m, v13, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	F_OffsetVarNodes(m, v22, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_IncrementVarSublevelsUp(m, v13, int32(-1), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_IncrementVarSublevelsUp(m, v22, int32(-1), int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v53 = F_pull_varnos(m, l0, v22)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v53 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if int32(0) <= v111 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v111 = base.I32_ctz(v97) | v98<<(uint(int32(5))%32)
	goto L22
L24:
	;
	v111 = int32(-2)
	goto L22
L25:
	;
	v64 = base.I32_div_s(int32(0), int32(32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v65 <= v64 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v68 = v53 + int32(8)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
	v75 = v72 & int32(-1)
	if v75 != 0 {
		v97 = v75
		v98 = v64
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v77 = v64 + int32(1)
	if v77 == v65 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v80 = v77
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v68+v80<<(uint(int32(2))%32))))
	if v87 != 0 {
		v97 = v87
		v98 = v80
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v89 = v80 + int32(1)
	if v89 != v65 {
		v80 = v89
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v118 = v111
	v121 = v5
	goto L36
L34:
	;
	v192 = v5
	goto L35
L35:
	;
	F_bms_free(m, v53)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L54
	}
L36:
	;
	if v118 <= v40 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v192 = v126
	goto L35
L38:
	;
	v124 = F_bms_add_member(m, v121, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L41
	}
L39:
	;
	v126 = v121
	goto L40
L40:
	;
	if v53 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v126 = v124
	goto L40
L42:
	;
	if int32(0) <= v182 {
		v118 = v182
		v121 = v126
		goto L36
	} else {
		goto L53
	}
L43:
	;
	v182 = base.I32_ctz(v168) | v169<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v182 = int32(-2)
	goto L42
L45:
	;
	v133 = v118 + int32(1)
	v135 = base.I32_div_s(v133, int32(32))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v136 <= v135 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v139 = v53 + int32(8)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v135<<(uint(int32(2))%32))))
	v146 = v143 & (int32(-1) << (uint(v133) % 32))
	if v146 != 0 {
		v168 = v146
		v169 = v135
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v148 = v135 + int32(1)
	if v148 == v136 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v151 = v148
	goto L49
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v139+v151<<(uint(int32(2))%32))))
	if v158 != 0 {
		v168 = v158
		v169 = v151
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v160 = v151 + int32(1)
	if v160 != v136 {
		v151 = v160
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L37
L54:
	;
	v196 = int32(0)
	if v192 == v196 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v250 == int32(0) {
		v294 = v196
		goto L1
	} else {
		goto L69
	}
L56:
	;
	v250 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if l3 == int32(0) {
		v241 = v196
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v250 = v241
	goto L55
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v207 < v206 {
		v241 = v196
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v209 = int32(1)
	if v206 <= v209 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v212 = v209
	goto L64
L63:
	;
	v212 = v206
	goto L64
L64:
	;
	v213 = int32(8)
	v218 = int32(0)
	goto L65
L65:
	;
	v225 = v218 << (uint(int32(2)) % 32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v192+v213+v225)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+(l3+v213))))
	v232 = v227 & (v229 ^ int32(-1))
	v234 = base.B2i32(v232 == int32(0))
	if v232 != 0 {
		v241 = v234
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v241 = v234
	goto L59
L67:
	;
	v236 = v218 + int32(1)
	if v236 != v212 {
		v218 = v236
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_CombineRangeTables(m, v12+int32(52), v12+int32(56), v257, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v262 = F_palloc0(m, int32(40))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = v264
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+8)) = uint8(v264)
	if l2 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v270 = int32(5)
	goto L74
L73:
	;
	v270 = int32(4)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = int32(64)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 == int32(0) {
		v283 = v274
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v284 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v262)+32)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v262)+28)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v262)+20)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v262)+16)) = v283
	v294 = v262
	goto L1
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 != int32(1) {
		v283 = v274
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = v282
	goto L75
}
func F_convert_combining_aggrefs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v80 = int32(0)
		m.G0 = v7 + int32(16)
		return v80
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 == int32(9) {
			v16 = F_palloc0(m, int32(72))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(9)
				v23 = F__emscripten_memcpy_bulkmem(m, v16, l0, int32(72))
				mBase = m.M
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v25
				v29 = F_copyObjectImpl(m, v23)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v31
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = int32(6)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					if v40 == int32(2281) {
						v43 = int32(17)
					} else {
						v43 = v40
					}
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v43
					v50 = int32(0)
					v52 = F_makeTargetEntry(m, v23, int32(1), v50, v50)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v52
						v59 = F_list_make1_impl(m, int32(1), v7+int32(8))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(9)
							v80 = v29
							m.G0 = v7 + int32(16)
							return v80
						}
					}
				}
			}
		} else {
			v77 = F_expression_tree_mutator_impl(m, l0, int32(836), l1)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				v80 = v77
				m.G0 = v7 + int32(16)
				return v80
			}
		}
	}
}
func F_copy_dest_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	return
}
func F_copydir(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v8 = m.G0
	v10 = v8 - int32(_a_F_copydir_0)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[0]))
	v14 = F_mkdir(m, l1, v13)
	mBase = m.M
	goto L1
L1:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v17 = F_AllocateDir(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L58
	}
L5:
	;
	return
L6:
	;
	v19 = F_ReadDir(m, v17, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v19
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_FreeDir(m, v17)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L33
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[1]))
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+19)))
	if v32 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v90 = F_ReadDir(m, v17, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L31
	}
L18:
	;
	v45 = v25 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v54 = F_pg_snprintf(m, v10+int32(2112), int32(2048), int32(_a_F_copydir_1), v10+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
	if v35 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
	if v38 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+21)))
	if v41 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v64 = F_pg_snprintf(m, v10-int32(-64), int32(2048), int32(_a_F_copydir_1), v10+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v70 = F_get_dirent_type(m, v10+int32(2112), v25, int32(0), int32(21))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_copy_file(m, v10+int32(2112), v10-int32(-64))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L30
	}
L26:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	switch v70 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L26
	default:
		goto L17
	}
L28:
	;
	F_copydir(m, v10+int32(2112), v10-int32(-64), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	goto L17
L31:
	;
	if v90 != 0 {
		v25 = v90
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copydir[2])))
	if v102 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = F_AllocateDir(m, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v10 + int32(_a_F_copydir_0)
	return
L37:
	;
	v107 = F_ReadDir(m, v105, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v107 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v113 = v107
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_FreeDir(m, v105)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L56
	}
L42:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+19)))
	if v116 != int32(46) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v151 = F_ReadDir(m, v105, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v113 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	v136 = F_pg_snprintf(m, v10-int32(-64), int32(2048), int32(_a_F_copydir_1), v10)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L50
	}
L46:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+20)))
	if v119 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+20)))
	if v122 != int32(46) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+21)))
	if v125 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v142 = F_get_dirent_type(m, v10-int32(-64), v113, int32(0), int32(21))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v142 != int32(2) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_fsync_fname(m, v10-int32(-64), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	if v151 != 0 {
		v113 = v151
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	F_fsync_fname(m, l1, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	goto L36
L58:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
	F_errmsg(m, int32(_a_F_copydir_2), v10+int32(48))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_copydir_3), int32(58), int32(_a_F_copydir_4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_core_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_core_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_core_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_cost_sort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 float64, l6 int32, l7 float64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v45 int64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v73 float64
	_ = v73
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v114 float64
	_ = v114
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = v12 + int32(8)
	v20 = float64(2)
	if base.F64_lt(l3, v20) != 0 {
		v23 = v20
	} else {
		v23 = l3
	}
	v25 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	v28 = base.F64_mul(v23, base.F64_add(base.F64_add(v25, v25), l5))
	v35 = base.F64_convert_i32_u((l4+int32(7))&int32(-8) + int32(24))
	v37 = base.F64_mul(l3, v35)
	v41 = base.F64_lt(l7, v23) & base.F64_gt(l7, float64(0))
	if v41 != 0 {
		v42 = base.F64_mul(l7, v35)
	} else {
		v42 = v37
	}
	v45 = base.I64_extend_i32_s(l6) << (uint(int64(10)) % 64)
	v46 = base.F64_convert_i64_s(v45)
	if base.F64_gt(v42, v46) != 0 {
		v48 = F_log(m, v23)
		mBase = m.M
		v49 = F_tuplesort_merge_order(m, v45)
		mBase = m.M
		v52 = base.F64_mul(base.F64_div(v48, float64(0.693147180559945)), v28)
		*(*float64)(unsafe.Add(mBase, uint32(v15))) = v52
		v57 = base.F64_ceil(base.F64_mul(v37, float64(0.0001220703125)))
		v59 = base.F64_div(v37, v46)
		v60 = base.F64_convert_i32_s(v49)
		if base.F64_gt(v59, v60) != 0 {
			v62 = F_log(m, v59)
			mBase = m.M
			v63 = F_log(m, v60)
			mBase = m.M
			v66 = base.F64_ceil(base.F64_div(v62, v63))
		} else {
			v66 = float64(1)
		}
		v69 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[1]))
		v73 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[2]))
		v98 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v57, v57), v66), base.F64_add(base.F64_mul(v69, float64(0.75)), base.F64_mul(v73, float64(0.25)))), v52)
	} else {
		if v41 != 0 {
			v80 = l7
		} else {
			v80 = v23
		}
		v81 = base.F64_add(v80, v80)
		if base.F64_gt(v37, v46)|base.F64_gt(v23, v81) != 0 {
			v84 = F_log(m, v81)
			mBase = m.M
			v98 = base.F64_mul(base.F64_div(v84, float64(0.693147180559945)), v28)
		} else {
			v88 = F_log(m, v23)
			mBase = m.M
			v98 = base.F64_mul(base.F64_div(v88, float64(0.693147180559945)), v28)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(v15))) = v98
	v101 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_mul(v23, v101)
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_sort[3])))
	v108 = base.F64_add(l2, v104)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v107 ^ int32(1))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v108, v114)
	m.G0 = v12 + int32(16)
	return
}
