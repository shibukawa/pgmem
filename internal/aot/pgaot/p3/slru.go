package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruInternalWritePage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int64
	_ = v169
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v211 int64
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v229 int64
	_ = v229
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v257 int64
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v322 int64
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v672 int64
	_ = v672
	v18 = m.G0
	v20 = v18 - int32(1072)
	m.G0 = v20
	v22 = int32(3)
	v23 = l1 << (uint(v22) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v25)))
	v29 = l1 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v31 = v29 + v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(1072)
	return
L2:
	;
	goto L5
L3:
	;
	v66 = v32
	v67 = v31
	goto L4
L4:
	;
	if v66 != int32(2) {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v52+v23)))
	if v54 != v27 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v66 = v60
	v67 = v59
	goto L4
L7:
	;
	F_SimpleLruWaitIO(m, l0, l1)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v59 = v58 + v29
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 == int32(3) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+l1))))
	if v84&int32(1) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v89+l1<<(uint(int32(3))%32))))
	if v93 != v27 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(3)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97+l1))) = uint8(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v106 = F_LWLockAcquire(m, v101+l1<<(uint(int32(7))%32), v99)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v110 = l1 >> (uint(int32(4)) % 32)
	F_LWLockRelease(m, v108+v110<<(uint(int32(7))%32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+56))
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[0])) = uint8(v119)
	*(*uint8)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[1])) = uint8(v119)
	v125 = v117 << (uint(int32(6)) % 32)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_SlruInternalWritePage[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_SlruInternalWritePage[2]))) = v128 + int64(1)
	v133 = base.I64_div_s(v27, int64(32))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)+36))
	if v134 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l2 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L17:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)+40))
	v138 = v137 * l1
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v134+v138<<(uint(int32(3))%32))))
	if v137 < int32(2) {
		v257 = v142
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v257 == int64(0) {
		goto L16
	} else {
		goto L45
	}
L19:
	;
	v146 = v137 - int32(1)
	v147 = int32(3)
	v148 = v146 & v147
	if base.Ui32(v147) <= base.Ui32(v137-int32(2)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v159 = v138
	v160 = int32(0)
	v169 = v142
	goto L23
L21:
	;
	v201 = v138
	v211 = v142
	goto L22
L22:
	;
	v219 = v201
	v220 = int32(0)
	v229 = v211
	goto L39
L23:
	;
	v175 = v134 + v159<<(uint(int32(3))%32)
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v175)+8))
	if base.Ui64(v176) < base.Ui64(v169) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v148 == int32(0) {
		v257 = v192
		goto L18
	} else {
		goto L38
	}
L25:
	;
	v178 = v169
	goto L27
L26:
	;
	v178 = v176
	goto L27
L27:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v175)+16))
	if base.Ui64(v179) < base.Ui64(v178) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v181 = v178
	goto L30
L29:
	;
	v181 = v179
	goto L30
L30:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v175)+24))
	if base.Ui64(v182) < base.Ui64(v181) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v184 = v181
	goto L33
L32:
	;
	v184 = v182
	goto L33
L33:
	;
	v186 = v159 + int32(4)
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v134+v186<<(uint(int32(3))%32))))
	if base.Ui64(v190) < base.Ui64(v184) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v192 = v184
	goto L36
L35:
	;
	v192 = v190
	goto L36
L36:
	;
	v194 = v160 + int32(4)
	if v194 != v146&int32(-4) {
		v159 = v186
		v160 = v194
		v169 = v192
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L24
L38:
	;
	v201 = v186
	v211 = v192
	goto L22
L39:
	;
	v234 = v219 + int32(1)
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v134+v234<<(uint(int32(3))%32))))
	if base.Ui64(v238) < base.Ui64(v229) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v257 = v240
	goto L18
L41:
	;
	v240 = v229
	goto L43
L42:
	;
	v240 = v238
	goto L43
L43:
	;
	v242 = v220 + int32(1)
	if v242 != v148 {
		v219 = v234
		v220 = v242
		v229 = v240
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v263 = int32(_a_F_SlruInternalWritePage_0)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[3])) = v265 + int32(1)
	F_XLogFlush(m, v257)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v271 = int32(_a_F_SlruInternalWritePage_0)
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[3])) = v273 - int32(1)
	goto L16
L47:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L109
	}
L48:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v631 = F_LWLockAcquire(m, v626+v110<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L8
	} else {
		goto L107
	}
L49:
	;
	if l2 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L50:
	;
	v425 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4])) = v425
	v427 = int32(_a_F_SlruInternalWritePage_1)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = int32(167772215)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431+l1<<(uint(int32(2))%32))))
	v436 = int32(_a_F_SlruInternalWritePage_2)
	v444 = F_pwrite(m, v410, v435, v436, base.I64_extend_i32_s(base.I32_wrap_i64(v27-v133<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = v425
	if v444 != v436 {
		goto L73
	} else {
		goto L74
	}
L51:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v352 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v296 <= int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v305 = int32(0)
	goto L54
L54:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(72)+v305<<(uint(int32(3))%32))))
	if v133 != v322 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v327 = int32(0)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l2+v305<<(uint(int32(2))%32))+4))
	if v327 <= v331 {
		v410 = v331
		v411 = v327
		goto L50
	} else {
		goto L60
	}
L56:
	;
	v325 = v305 + int32(1)
	if v325 != v296 {
		v305 = v325
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L51
L60:
	;
	goto L51
L61:
	;
	v376 = F_OpenTransientFile(m, v20+int32(48), int32(66))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L67
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v351
	v361 = F_pg_snprintf(m, v20+int32(48), int32(1024), int32(_a_F_SlruInternalWritePage_3), v20)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+20)) = uint32(v133)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v351
	v371 = F_pg_snprintf(m, v20+int32(48), int32(1024), int32(_a_F_SlruInternalWritePage_4), v20+int32(16))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L8
	} else {
		goto L66
	}
L65:
	;
	goto L61
L66:
	;
	goto L61
L67:
	;
	if v376 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(0)
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v385
	goto L49
L69:
	;
	goto L70
L70:
	;
	v387 = int32(1)
	if l2 == int32(0) {
		v410 = v376
		v411 = v387
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(15) < v390 {
		v410 = v376
		v411 = v387
		goto L50
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v390<<(uint(int32(2))%32))+4)) = v376
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v397<<(uint(int32(3))%32))+72)) = v133
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v402 + int32(1)
	v410 = v376
	v411 = int32(0)
	goto L50
L73:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	if v453 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v467 == int32(5) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v458 = v453
	goto L78
L77:
	;
	v455 = int32(51)
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4])) = v455
	v458 = v455
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v458
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(3)
	if v411 == int32(0) {
		goto L49
	} else {
		goto L79
	}
L79:
	;
	v465 = F_CloseTransientFile(m, v410)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	goto L49
L81:
	;
	if v411 == int32(0) {
		goto L48
	} else {
		goto L94
	}
L82:
	;
	v470 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v470
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v470
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v133
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v467)
	v478 = int32(0)
	v480 = F_RegisterSyncRequest(m, v20+int32(24), v478, v478)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	if v480 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = int32(167772214)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[8])))
	if v488 != int32(1) {
		v502 = int32(0)
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	v505 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v505
	if v502 == v505 {
		goto L81
	} else {
		goto L92
	}
L86:
	;
	goto L85
L87:
	;
	goto L88
L88:
	;
	v493 = F_fsync(m, v410)
	mBase = m.M
	if v493 != int32(-1) {
		v502 = v493
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v502 = int32(-1)
	goto L86
L90:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	if v497 == int32(27) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(4)
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v514
	v516 = F_CloseTransientFile(m, v410)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v520 = F_CloseTransientFile(m, v410)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	if v520 == int32(0) {
		goto L48
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(5)
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v529
	goto L49
L97:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v605 = F_LWLockAcquire(m, v600+v110<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L8
	} else {
		goto L104
	}
L98:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v550 <= int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v559 = int32(0)
	goto L100
L100:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(4)+v559<<(uint(int32(2))%32))))
	v577 = F_CloseTransientFile(m, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L8
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	v580 = v559 + int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v580 < v581 {
		v559 = v580
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v609 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v607+l1))) = uint8(v609)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v612 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v611+l1<<(uint(v612)%32)))) = v612
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F_LWLockRelease(m, v617+l1<<(uint(int32(7))%32))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_SlruReportIOError(m, l0, v27, int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	goto L47
L107:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v634 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v633+l1<<(uint(v634)%32)))) = v634
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F_LWLockRelease(m, v639+l1<<(uint(int32(7))%32))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	goto L47
L109:
	;
	v664 = int32(_a_F_SlruInternalWritePage_5)
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[9])) = v666 + int32(1)
	v670 = int32(_a_F_SlruInternalWritePage_6)
	v672 = *(*int64)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[10])) = v672 + int64(1)
	goto L1
}
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = F_OpenTransientFile(m, l2, int32(2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L9
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
	v20 = F_pg_snprintf(m, l2, int32(1024), int32(_a_F_SlruSyncFileTag_0), v9)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v12)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v30 = F_pg_snprintf(m, l2, int32(1024), int32(_a_F_SlruSyncFileTag_1), v9+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L1
L7:
	;
	goto L1
L8:
	;
	m.G0 = v9 + int32(32)
	return v70
L9:
	;
	if v33 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = int32(-1)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(167772212)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[1])))
	if v44 != int32(1) {
		v58 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2]))
	v65 = F_CloseTransientFile(m, v33)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	goto L16
L16:
	;
	v49 = F_fsync(m, v33)
	mBase = m.M
	if v49 != int32(-1) {
		v58 = v49
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v58 = int32(-1)
	goto L14
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2]))
	if v53 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2])) = v64
	v70 = v58
	goto L8
}
