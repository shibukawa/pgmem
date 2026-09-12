package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruInternalWritePage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v187 int64
	_ = v187
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v232 int64
	_ = v232
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v255 int64
	_ = v255
	var v260 int32
	_ = v260
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int32
	_ = v268
	var v286 int64
	_ = v286
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v357 int64
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int64
	_ = v512
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v596 int32
	_ = v596
	var v605 int32
	_ = v605
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int64
	_ = v727
	v21 = m.G0
	v23 = v21 - int32(1072)
	m.G0 = v23
	v25 = int32(3)
	v26 = l1 << (uint(v25) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v26+v28)))
	v32 = l1 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v34 = v32 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 == v25 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(1072)
	return
L2:
	;
	goto L5
L3:
	;
	v72 = v35
	v73 = v34
	goto L4
L4:
	;
	if v72 != int32(2) {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v58+v26)))
	if v60 != v30 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v72 = v66
	v73 = v65
	goto L4
L7:
	;
	F_SimpleLruWaitIO(m, l0, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v65 = v64 + v32
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 == int32(3) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+l1))))
	if v93&int32(1) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v98+l1<<(uint(int32(3))%32))))
	if v102 != v30 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(3)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+l1))) = uint8(v108)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	v115 = F_LWLockAcquire(m, v110+l1<<(uint(int32(7))%32), v108)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v119 = l1 >> (uint(int32(4)) % 32)
	F_LWLockRelease(m, v117+v119<<(uint(int32(7))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+56))
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v128)
	*(*uint8)(unsafe.Add(mBase, _consts[155])) = uint8(v128)
	v134 = v126 << (uint(int32(6)) % 32)
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[156])))
	*(*int64)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[156]))) = v137 + int64(1)
	v142 = base.I64_div_s(v30, int64(32))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
	if v143 == int32(0) {
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
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v125)+40))
	v147 = v146 * l1
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v143+v147<<(uint(int32(3))%32))))
	if v146 < int32(2) {
		v286 = v151
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v286 == int64(0) {
		goto L16
	} else {
		goto L45
	}
L19:
	;
	v155 = v146 - int32(1)
	v156 = int32(3)
	v157 = v155 & v156
	if base.Ui32(v156) <= base.Ui32(v146-int32(2)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v174 = v147
	v179 = int32(0)
	v187 = v151
	goto L23
L21:
	;
	v219 = v147
	v232 = v151
	goto L22
L22:
	;
	if v157 == int32(0) {
		v286 = v232
		goto L18
	} else {
		goto L38
	}
L23:
	;
	v192 = v174 << (uint(int32(3)) % 32)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v143+int32(8)+v192)))
	if base.Ui64(v194) < base.Ui64(v187) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v219 = v206
	v232 = v212
	goto L22
L25:
	;
	v196 = v187
	goto L27
L26:
	;
	v196 = v194
	goto L27
L27:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v192+(v143+int32(16)))))
	if base.Ui64(v198) < base.Ui64(v196) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v200 = v196
	goto L30
L29:
	;
	v200 = v198
	goto L30
L30:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v192+(v143+int32(24)))))
	if base.Ui64(v202) < base.Ui64(v200) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v204 = v200
	goto L33
L32:
	;
	v204 = v202
	goto L33
L33:
	;
	v206 = v174 + int32(4)
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v143+v206<<(uint(int32(3))%32))))
	if base.Ui64(v210) < base.Ui64(v204) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v212 = v204
	goto L36
L35:
	;
	v212 = v210
	goto L36
L36:
	;
	v214 = v179 + int32(4)
	if v214 != v155&int32(-4) {
		v174 = v206
		v179 = v214
		v187 = v212
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L24
L38:
	;
	v242 = v219
	v244 = int32(0)
	v255 = v232
	goto L39
L39:
	;
	v260 = v242 + int32(1)
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v143+v260<<(uint(int32(3))%32))))
	if base.Ui64(v264) < base.Ui64(v255) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v286 = v266
	goto L18
L41:
	;
	v266 = v255
	goto L43
L42:
	;
	v266 = v264
	goto L43
L43:
	;
	v268 = v244 + int32(1)
	if v268 != v157 {
		v242 = v260
		v244 = v268
		v255 = v266
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v292 = int32(4556756)
	v294 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v294 + int32(1)
	F_XLogFlush(m, v286)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v300 = int32(4556756)
	v302 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v302 - int32(1)
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
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v683 = F_LWLockAcquire(m, v678+v119<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
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
	v467 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v467
	v469 = int32(4164780)
	v470 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = int32(167772213)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473+l1<<(uint(int32(2))%32))))
	v478 = int32(8192)
	v486 = F_pwrite(m, v449, v477, v478, base.I64_extend_i32_s(base.I32_wrap_i64(v30-v142<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v488 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v467
	if v486 != v478 {
		goto L73
	} else {
		goto L74
	}
L51:
	;
	v390 = l0 + int32(16)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v391 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v328 <= int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v337 = int32(0)
	goto L54
L54:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(72)+v337<<(uint(int32(3))%32))))
	if v142 != v357 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v362 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l2+v337<<(uint(int32(2))%32))+4))
	if v362 <= v366 {
		v449 = v366
		v451 = v362
		goto L50
	} else {
		goto L60
	}
L56:
	;
	v360 = v337 + int32(1)
	if v360 != v328 {
		v337 = v360
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
	v415 = F_OpenTransientFile(m, v23+int32(48), int32(66))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L8
	} else {
		goto L67
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v390
	v400 = F_pg_snprintf(m, v23+int32(48), int32(1024), int32(536846), v23)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+20)) = uint32(v142)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v390
	v410 = F_pg_snprintf(m, v23+int32(48), int32(1024), int32(537240), v23+int32(16))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	if v415 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _consts[157])) = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v424
	goto L49
L69:
	;
	goto L70
L70:
	;
	v426 = int32(1)
	if l2 == int32(0) {
		v449 = v415
		v451 = v426
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(15) < v429 {
		v449 = v415
		v451 = v426
		goto L50
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v429<<(uint(int32(2))%32))+4)) = v415
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v436<<(uint(int32(3))%32))+72)) = v142
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v441 + int32(1)
	v449 = v415
	v451 = int32(0)
	goto L50
L73:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v495 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v509 == int32(5) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v500 = v495
	goto L78
L77:
	;
	v497 = int32(51)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v497
	v500 = v497
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v500
	*(*int32)(unsafe.Add(mBase, _consts[157])) = int32(3)
	if v451 == int32(0) {
		goto L49
	} else {
		goto L79
	}
L79:
	;
	v507 = F_CloseTransientFile(m, v449)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	goto L49
L81:
	;
	if v451 == int32(0) {
		goto L48
	} else {
		goto L94
	}
L82:
	;
	v512 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v512
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v512
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v142
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+24)) = uint16(v509)
	v520 = int32(0)
	v522 = F_RegisterSyncRequest(m, v23+int32(24), v520, v520)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	if v522 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = int32(167772212)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v530 != int32(1) {
		v544 = int32(0)
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	v547 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v547
	if v544 == v547 {
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
	v535 = F_fsync(m, v449)
	mBase = m.M
	if v535 != int32(-1) {
		v544 = v535
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v544 = int32(-1)
	goto L86
L90:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v539 == int32(27) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[157])) = int32(4)
	v556 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v556
	v558 = F_CloseTransientFile(m, v449)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L49
L94:
	;
	v563 = F_CloseTransientFile(m, v449)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	if v563 == int32(0) {
		goto L48
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[157])) = int32(5)
	v572 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v572
	goto L49
L97:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v657 = F_LWLockAcquire(m, v652+v119<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L8
	} else {
		goto L104
	}
L98:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v596 <= int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v605 = int32(0)
	goto L100
L100:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(4)+v605<<(uint(int32(2))%32))))
	v626 = F_CloseTransientFile(m, v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L8
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	v629 = v605 + int32(1)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v629 < v630 {
		v605 = v629
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v661 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v659+l1))) = uint8(v661)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v664 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v663+l1<<(uint(v664)%32)))) = v664
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	F_LWLockRelease(m, v669+l1<<(uint(int32(7))%32))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_SlruReportIOError(m, l0, v30, int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	goto L47
L107:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v686 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v685+l1<<(uint(v686)%32)))) = v686
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	F_LWLockRelease(m, v691+l1<<(uint(int32(7))%32))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	goto L47
L109:
	;
	v719 = int32(4457780)
	v721 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v721 + int32(1)
	v725 = int32(4541800)
	v727 = *(*int64)(unsafe.Add(mBase, _consts[160]))
	*(*int64)(unsafe.Add(mBase, _consts[160])) = v727 + int64(1)
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
	var v70 int32
	_ = v70
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
	v21 = F_pg_snprintf(m, l2, int32(1024), int32(536846), v9)
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
	v31 = F_pg_snprintf(m, l2, int32(1024), int32(537240), v9+int32(16))
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
	return v70
L9:
	;
	if v34 < int32(0) {
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(167772210)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v45 != int32(1) {
		v59 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[43]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _consts[43]))
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
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v65
	v70 = v59
	goto L8
}
