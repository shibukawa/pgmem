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
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int64
	_ = v471
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v551 int32
	_ = v551
	var v560 int32
	_ = v560
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v673 int64
	_ = v673
	v18 = m.G0
	v20 = v18 - int32(1072)
	m.G0 = v20
	v22 = int32(3)
	v23 = l1 << (uint(v22) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
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
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
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
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v632 = F_LWLockAcquire(m, v627+v110<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
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
	v426 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4])) = v426
	v428 = int32(_a_F_SlruInternalWritePage_1)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(167772213)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+l1<<(uint(int32(2))%32))))
	v437 = int32(_a_F_SlruInternalWritePage_2)
	v445 = F_pwrite(m, v411, v436, v437, base.I64_extend_i32_s(base.I32_wrap_i64(v27-v133<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v426
	if v445 != v437 {
		goto L73
	} else {
		goto L74
	}
L51:
	;
	v352 = l0 + int32(16)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v353 == int32(1) {
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
		v411 = v331
		v412 = v327
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
	v377 = F_OpenTransientFile(m, v20+int32(48), int32(66))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L8
	} else {
		goto L67
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v352
	v362 = F_pg_snprintf(m, v20+int32(48), int32(1024), int32(_a_F_SlruInternalWritePage_3), v20)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v352
	v372 = F_pg_snprintf(m, v20+int32(48), int32(1024), int32(_a_F_SlruInternalWritePage_4), v20+int32(16))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
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
	if v377 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(0)
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v386
	goto L49
L69:
	;
	goto L70
L70:
	;
	v388 = int32(1)
	if l2 == int32(0) {
		v411 = v377
		v412 = v388
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(15) < v391 {
		v411 = v377
		v412 = v388
		goto L50
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v391<<(uint(int32(2))%32))+4)) = v377
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v398<<(uint(int32(3))%32))+72)) = v133
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v403 + int32(1)
	v411 = v377
	v412 = int32(0)
	goto L50
L73:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	if v454 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v468 == int32(5) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v459 = v454
	goto L78
L77:
	;
	v456 = int32(51)
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4])) = v456
	v459 = v456
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v459
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(3)
	if v412 == int32(0) {
		goto L49
	} else {
		goto L79
	}
L79:
	;
	v466 = F_CloseTransientFile(m, v411)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	goto L49
L81:
	;
	if v412 == int32(0) {
		goto L48
	} else {
		goto L94
	}
L82:
	;
	v471 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v471
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v471
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v133
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v468)
	v479 = int32(0)
	v481 = F_RegisterSyncRequest(m, v20+int32(24), v479, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	if v481 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = int32(167772212)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[8])))
	if v489 != int32(1) {
		v503 = int32(0)
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[5]))
	v506 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v505))) = v506
	if v503 == v506 {
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
	v494 = F_fsync(m, v411)
	mBase = m.M
	if v494 != int32(-1) {
		v503 = v494
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v503 = int32(-1)
	goto L86
L90:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	if v498 == int32(27) {
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
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v515
	v517 = F_CloseTransientFile(m, v411)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v521 = F_CloseTransientFile(m, v411)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	if v521 == int32(0) {
		goto L48
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[6])) = int32(5)
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[7])) = v530
	goto L49
L97:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v606 = F_LWLockAcquire(m, v601+v110<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L8
	} else {
		goto L104
	}
L98:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v551 <= int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v560 = int32(0)
	goto L100
L100:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(4)+v560<<(uint(int32(2))%32))))
	v578 = F_CloseTransientFile(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L8
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	v581 = v560 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v581 < v582 {
		v560 = v581
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v610 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v608+l1))) = uint8(v610)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v613 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v612+l1<<(uint(v613)%32)))) = v613
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F_LWLockRelease(m, v618+l1<<(uint(int32(7))%32))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_SlruReportIOError(m, l0, v27, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	goto L47
L107:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v635 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v634+l1<<(uint(v635)%32)))) = v635
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F_LWLockRelease(m, v640+l1<<(uint(int32(7))%32))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	goto L47
L109:
	;
	v665 = int32(_a_F_SlruInternalWritePage_5)
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[9])) = v667 + int32(1)
	v671 = int32(_a_F_SlruInternalWritePage_6)
	v673 = *(*int64)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_SlruInternalWritePage[10])) = v673 + int64(1)
	goto L1
}
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = l0 + int32(16)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v14 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = F_OpenTransientFile(m, l2, int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L9
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	v21 = F_pg_snprintf(m, l2, int32(1024), int32(_a_F_SlruSyncFileTag_0), v9)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v13)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	v31 = F_pg_snprintf(m, l2, int32(1024), int32(_a_F_SlruSyncFileTag_1), v9+int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	return v71
L9:
	;
	if v34 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v71 = int32(-1)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(167772210)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[1])))
	if v45 != int32(1) {
		v59 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2]))
	v66 = F_CloseTransientFile(m, v34)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v50 = F_fsync(m, v34)
	mBase = m.M
	if v50 != int32(-1) {
		v59 = v50
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(-1)
	goto L14
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2]))
	if v54 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SlruSyncFileTag[2])) = v65
	v71 = v59
	goto L8
}
