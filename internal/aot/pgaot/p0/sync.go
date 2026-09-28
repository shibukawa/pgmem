package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_SyncRepReleaseWaiters(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v256 int64
	_ = v256
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int64
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v457 int64
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v548 int64
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v673 int32
	_ = v673
	var v674 int64
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v704 int32
	_ = v704
	var v705 int64
	_ = v705
	var v708 int64
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int64
	_ = v812
	var v813 int64
	_ = v813
	var v818 int64
	_ = v818
	var v824 int64
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v854 int32
	_ = v854
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	v1 = int32(0)
	v16 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	if v26 == v1 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v22 + int32(80)
	return
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[1]))
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[2]))
	v48 = F_LWLockAcquire(m, v44+int32(_a_F_SyncRepReleaseWaiters_0), v42)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[3])) = uint8(v38)
	goto L1
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(int32(1)) < base.Ui32(v29-int32(3)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v25)+32))
	if v34 != int64(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	return
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[4]))
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[2]))
	F_LWLockRelease(m, v896+int32(_a_F_SyncRepReleaseWaiters_0))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L7
	} else {
		goto L164
	}
L10:
	;
	v54 = F_SyncRepGetCandidateStandbys(m, v22+int32(76))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v894 = int32(1)
	goto L9
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	if v54 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_pfree(m, v56)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L7
	} else {
		goto L163
	}
L15:
	;
	v59 = v42
	goto L16
L16:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v59*int32(48))+40)))
	if v81 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[4]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = base.B2i32(v54 < v89)
	if v54 < v89 {
		v366 = v16
		v367 = v16
		v368 = v16
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v85 = v59 + int32(1)
	if v54 != v85 {
		v59 = v85
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L14
L22:
	;
	F_pfree(m, v56)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L79
	}
L23:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
	if v91 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v54 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v196 = int32(0)
	v198 = F_palloc_mul(m, int32(8), v54)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L62
	}
L27:
	;
	v180 = v56 + v163*int32(48)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v180)+24))
	if base.Ui64(v175-int64(1)) < base.Ui64(v181) {
		goto L53
	} else {
		goto L54
	}
L28:
	;
	v163 = v1
	v175 = v16
	v176 = v16
	v177 = v16
	goto L27
L29:
	;
	goto L30
L30:
	;
	v104 = v1
	v106 = v1
	v116 = v16
	v117 = v16
	v118 = v16
	goto L31
L31:
	;
	v121 = v56 + v104*int32(48)
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v121)+24))
	if base.Ui64(v116-int64(1)) < base.Ui64(v122) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v54&int32(1) == int32(0) {
		v366 = v131
		v367 = v141
		v368 = v151
		goto L22
	} else {
		goto L52
	}
L33:
	;
	v126 = v116
	goto L35
L34:
	;
	v126 = v122
	goto L35
L35:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v121)+72))
	if base.Ui64(v126-int64(1)) < base.Ui64(v127) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = v126
	goto L38
L37:
	;
	v131 = v127
	goto L38
L38:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v121)+16))
	if base.Ui64(v117-int64(1)) < base.Ui64(v132) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v136 = v117
	goto L41
L40:
	;
	v136 = v132
	goto L41
L41:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v121)+64))
	if base.Ui64(v136-int64(1)) < base.Ui64(v137) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v141 = v136
	goto L44
L43:
	;
	v141 = v137
	goto L44
L44:
	;
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v121)+8))
	if base.Ui64(v118-int64(1)) < base.Ui64(v142) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v146 = v118
	goto L47
L46:
	;
	v146 = v142
	goto L47
L47:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v121)+56))
	if base.Ui64(v146-int64(1)) < base.Ui64(v147) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v151 = v146
	goto L50
L49:
	;
	v151 = v147
	goto L50
L50:
	;
	v152 = int32(2)
	v153 = v104 + v152
	v155 = v106 + v152
	if v155 != v54&int32(2147483646) {
		v104 = v153
		v106 = v155
		v116 = v131
		v117 = v141
		v118 = v151
		goto L31
	} else {
		goto L51
	}
L51:
	;
	goto L32
L52:
	;
	v163 = v153
	v175 = v131
	v176 = v141
	v177 = v151
	goto L27
L53:
	;
	v185 = v175
	goto L55
L54:
	;
	v185 = v181
	goto L55
L55:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v180)+16))
	if base.Ui64(v176-int64(1)) < base.Ui64(v186) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v190 = v176
	goto L58
L57:
	;
	v190 = v186
	goto L58
L58:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	if base.Ui64(v177-int64(1)) < base.Ui64(v191) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v195 = v177
	goto L61
L60:
	;
	v195 = v191
	goto L61
L61:
	;
	v366 = v185
	v367 = v190
	v368 = v195
	goto L22
L62:
	;
	v201 = F_palloc_mul(m, int32(8), v54)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v204 = F_palloc_mul(m, int32(8), v54)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	if v54 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	F_pg_qsort(m, v198, v54, int32(8), int32(1101))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L73
	}
L66:
	;
	v212 = v196
	v226 = v1
	goto L69
L67:
	;
	v268 = v196
	goto L68
L68:
	;
	v288 = v268 << (uint(int32(3)) % 32)
	v292 = v56 + v268*int32(48)
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v292)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v198+v288))) = v293
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v292)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v201+v288))) = v296
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v292)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v204+v288))) = v299
	goto L65
L69:
	;
	v231 = int32(3)
	v232 = v212 << (uint(v231) % 32)
	v234 = int32(48)
	v236 = v56 + v212*v234
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v236)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v198+v232))) = v237
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v236)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v201+v232))) = v240
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v236)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v204+v232))) = v243
	v246 = v212 | int32(1)
	v248 = v246 << (uint(v231) % 32)
	v252 = v56 + v246*v234
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v198+v248))) = v253
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v201+v248))) = v256
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v252)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v204+v248))) = v259
	v261 = int32(2)
	v262 = v212 + v261
	v264 = v226 + v261
	if v264 != v54&int32(2147483646) {
		v212 = v262
		v226 = v264
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v54&int32(1) == int32(0) {
		goto L65
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v268 = v262
	goto L68
L73:
	;
	F_pg_qsort(m, v201, v54, int32(8), int32(1101))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	F_pg_qsort(m, v204, v54, int32(8), int32(1101))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v337 = v89&int32(255)<<(uint(int32(3))%32) - int32(8)
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v204+v337)))
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v337+v201)))
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v337+v198)))
	F_pfree(m, v198)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v201)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_pfree(m, v204)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	v366 = v339
	v367 = v341
	v368 = v343
	goto L22
L79:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[3])))
	if v372 != int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v54 < v89 {
		v894 = int32(0)
		goto L9
	} else {
		goto L92
	}
L81:
	;
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[3])) = uint8(v376)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[4]))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+8)))
	v383 = F_errstart(m, int32(15), v376)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	if v380 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_errfinish(m, int32(_a_F_SyncRepReleaseWaiters_1), v415, int32(_a_F_SyncRepReleaseWaiters_2))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L91
	}
L84:
	;
	if v383 == int32(0) {
		goto L80
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v383 == int32(0) {
		goto L80
	} else {
		goto L89
	}
L87:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[0]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+72))
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v392
	F_errmsg(m, int32(_a_F_SyncRepReleaseWaiters_3), v22+int32(48))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v415 = int32(539)
	goto L83
L89:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v406
	F_errmsg(m, int32(_a_F_SyncRepReleaseWaiters_4), v22-int32(-64))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v415 = int32(543)
	goto L83
L91:
	;
	goto L80
L92:
	;
	v422 = int32(0)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
	if base.Ui64(v368) <= base.Ui64(v424) {
		v535 = v422
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v41)+32))
	if base.Ui64(v367) <= base.Ui64(v548) {
		v658 = v422
		goto L114
	} else {
		goto L115
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v368
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[1]))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	if base.B2i32(v429 == int32(0))|base.B2i32(v429 == v428) != 0 {
		v535 = v422
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v434 = v429
	v440 = v422
	goto L96
L96:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v428)+24))
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v434-int32(12))))
	if base.Ui64(v454) < base.Ui64(v457) {
		v535 = v440
		goto L93
	} else {
		goto L98
	}
L97:
	;
	v535 = v527
	goto L93
L98:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v453
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = int64(0)
	v465 = int32(0)
	v468 = base.AtomicRmwOr32(m, v465, int32(_a_F_SyncRepReleaseWaiters_5), v465)
	*(*int32)(unsafe.Add(mBase, uint32(v434-int32(4)))) = int32(2)
	v474 = v434 - int32(280)
	v478 = base.AtomicRmwOr32(m, v465, int32(_a_F_SyncRepReleaseWaiters_6), v465)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	if v479 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v527 = v440 + int32(1)
	if v428 != v453 {
		v434 = v453
		v440 = v527
		goto L96
	} else {
		goto L113
	}
L100:
	;
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = int32(1)
	v482 = int32(0)
	v485 = base.AtomicRmwOr32(m, v482, int32(_a_F_SyncRepReleaseWaiters_6), v482)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v474)+4))
	if v486 == v482 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v474)+12))
	if v489 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[6]))
	if v493 == v489 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v495 = m.G0
	v497 = v495 - int32(16)
	m.G0 = v497
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[7]))
	if v500 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v523 = F_pgmem_kill(m, v489, int32(23))
	mBase = m.M
	goto L100
L107:
	;
	m.G0 = v497 + int32(16)
	goto L99
L108:
	;
	v503 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v497)+15)) = uint8(v503)
	goto L109
L109:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[8]))
	v511 = F_write(m, v507, v497+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v511 {
		goto L107
	} else {
		goto L111
	}
L110:
	;
	goto L107
L111:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[9]))
	if v515 == int32(27) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L97
L114:
	;
	v673 = int32(0)
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v41)+40))
	if base.Ui64(v366) <= base.Ui64(v674) {
		v782 = v673
		goto L136
	} else {
		goto L137
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v367
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[1]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+12))
	if v553 == int32(0) {
		v658 = v422
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v557 = v552 + int32(8)
	if v553 == v557 {
		v658 = v422
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v559 = v553
	v563 = v422
	goto L118
L118:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v552)+32))
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v559-int32(12))))
	if base.Ui64(v579) < base.Ui64(v582) {
		v658 = v563
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v658 = v652
	goto L114
L120:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+4)) = v578
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	*(*int32)(unsafe.Add(mBase, uint32(v578))) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v559))) = int64(0)
	v590 = int32(0)
	v593 = base.AtomicRmwOr32(m, v590, int32(_a_F_SyncRepReleaseWaiters_5), v590)
	*(*int32)(unsafe.Add(mBase, uint32(v559-int32(4)))) = int32(2)
	v599 = v559 - int32(280)
	v603 = base.AtomicRmwOr32(m, v590, int32(_a_F_SyncRepReleaseWaiters_6), v590)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	if v604 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v652 = v563 + int32(1)
	if v557 != v578 {
		v559 = v578
		v563 = v652
		goto L118
	} else {
		goto L135
	}
L122:
	;
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v599))) = int32(1)
	v607 = int32(0)
	v610 = base.AtomicRmwOr32(m, v607, int32(_a_F_SyncRepReleaseWaiters_6), v607)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v611 == v607 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v599)+12))
	if v614 == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[6]))
	if v618 == v614 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v620 = m.G0
	v622 = v620 - int32(16)
	m.G0 = v622
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[7]))
	if v625 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v648 = F_pgmem_kill(m, v614, int32(23))
	mBase = m.M
	goto L122
L129:
	;
	m.G0 = v622 + int32(16)
	goto L121
L130:
	;
	v628 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v622)+15)) = uint8(v628)
	goto L131
L131:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[8]))
	v636 = F_write(m, v632, v622+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v636 {
		goto L129
	} else {
		goto L133
	}
L132:
	;
	goto L129
L133:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[9]))
	if v640 == int32(27) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	goto L119
L136:
	;
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[2]))
	F_LWLockRelease(m, v800+int32(_a_F_SyncRepReleaseWaiters_0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L7
	} else {
		goto L158
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+40)) = v366
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[1]))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+20))
	if v679 == int32(0) {
		v782 = v673
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v683 = v678 + int32(16)
	if v679 == v683 {
		v782 = v673
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v685 = v679
	v687 = v673
	goto L140
L140:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v678)+40))
	v708 = *(*int64)(unsafe.Add(mBase, uint32(v685-int32(12))))
	if base.Ui64(v705) < base.Ui64(v708) {
		v782 = v687
		goto L136
	} else {
		goto L142
	}
L141:
	;
	v782 = v778
	goto L136
L142:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v704
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	*(*int32)(unsafe.Add(mBase, uint32(v704))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v685))) = int64(0)
	v716 = int32(0)
	v719 = base.AtomicRmwOr32(m, v716, int32(_a_F_SyncRepReleaseWaiters_5), v716)
	*(*int32)(unsafe.Add(mBase, uint32(v685-int32(4)))) = int32(2)
	v725 = v685 - int32(280)
	v729 = base.AtomicRmwOr32(m, v716, int32(_a_F_SyncRepReleaseWaiters_6), v716)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	if v730 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v778 = v687 + int32(1)
	if v683 != v704 {
		v685 = v704
		v687 = v778
		goto L140
	} else {
		goto L157
	}
L144:
	;
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = int32(1)
	v733 = int32(0)
	v736 = base.AtomicRmwOr32(m, v733, int32(_a_F_SyncRepReleaseWaiters_6), v733)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	if v737 == v733 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v725)+12))
	if v740 == int32(0) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[6]))
	if v744 == v740 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v746 = m.G0
	v748 = v746 - int32(16)
	m.G0 = v748
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[7]))
	if v751 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v774 = F_pgmem_kill(m, v740, int32(23))
	mBase = m.M
	goto L144
L151:
	;
	m.G0 = v748 + int32(16)
	goto L143
L152:
	;
	v754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v748)+15)) = uint8(v754)
	goto L153
L153:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[8]))
	v762 = F_write(m, v758, v748+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v762 {
		goto L151
	} else {
		goto L155
	}
L154:
	;
	goto L151
L155:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[9]))
	if v766 == int32(27) {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	goto L141
L158:
	;
	v807 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L7
	} else {
		goto L159
	}
L159:
	;
	if v807 == int32(0) {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+32)) = uint32(v366)
	v812 = int64(32)
	v813 = int64(base.Ui64(v366) >> (uint(v812) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+28)) = uint32(v813)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v782
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+20)) = uint32(v367)
	v818 = int64(base.Ui64(v367) >> (uint(v812) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+16)) = uint32(v818)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v535
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+8)) = uint32(v368)
	v824 = int64(base.Ui64(v368) >> (uint(v812) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+4)) = uint32(v824)
	F_errmsg_internal(m, int32(_a_F_SyncRepReleaseWaiters_7), v22)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_SyncRepReleaseWaiters_1), int32(582), int32(_a_F_SyncRepReleaseWaiters_2))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L7
	} else {
		goto L162
	}
L162:
	;
	goto L1
L163:
	;
	goto L12
L164:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepReleaseWaiters[3])) = uint8(v894)
	goto L1
}
func F_SyncReplicationSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(192)
	m.G0 = v15
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = int32(-1)
	v24 = v2
	v29 = int64(0)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v23 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v345 = int32(m.ExcTag)
	v346 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v345 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v21
	v35 = base.I64_extend_i32_u(l0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v35
	F_before_shmem_exit(m, int32(1092), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v52 = v20
	v53 = v21
	v54 = v24
	v55 = v29
	goto L8
L8:
	;
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1]))
	goto L10
L10:
	;
	v46 = v15 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v15 + int32(8)
	goto L13
L11:
	;
	v52 = v42
	v53 = v44
	v54 = int32(0)
	v55 = v35
	goto L8
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v15 + int32(16)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[2]))
	F_check_and_set_sync_info(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0])) = v52
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_cancel_before_shmem_exit(m, int32(1092), v55)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L5
	} else {
		goto L68
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_validate_remote_info(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[3]))
	v86 = F_AllocSetContextCreateInternal(m, v81, int32(_a_F_SyncReplicationSlots_0), int32(0), int32(_a_F_SyncReplicationSlots_1), int32(_a_F_SyncReplicationSlots_2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v93 = v22
	v94 = int32(0)
	goto L20
L20:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)) = uint8(v101)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[4]))
	if v104 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_MemoryContextDelete(m, v86)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L57
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_ProcessInterrupts(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[5]))
	if v112 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_slotsync_reread_config(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_MemoryContextReset(m, v86)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v125 = int32(_a_F_SyncReplicationSlots_3)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[3])) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v133 = F_fetch_remote_slots(m, l0, v94)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v141 = F_synchronize_slots(m, v133, v15+int32(15))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[3])) = v126
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	v148 = int32(0)
	if v94|base.B2i32(v145&int32(1) == v148) == v148 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L21
L34:
	;
	v211 = int32(_a_F_SyncReplicationSlots_4)
	v213 = int32(_a_F_SyncReplicationSlots_5)
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[6]))
	v217 = v215 << (uint(int32(1)) % 32)
	if v213 <= v217 {
		goto L48
	} else {
		goto L49
	}
L35:
	;
	if v133 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v145&int32(1) == int32(0) {
		goto L33
	} else {
		goto L47
	}
L38:
	;
	v203 = v93
	v204 = int32(0)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v156 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v158 <= v156 {
		v203 = v93
		v204 = v156
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v165 = v93
	v166 = v156
	v167 = v156
	goto L42
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v167<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v183 = F_pstrdup(m, v178)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v203 = v189
	v204 = v189
	goto L34
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v189 = F_lappend(m, v166, v183)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v192 = v167 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v192 < v193 {
		v165 = v189
		v166 = v189
		v167 = v192
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v203 = v93
	v204 = v94
	goto L34
L48:
	;
	v220 = v213
	goto L50
L49:
	;
	v220 = v217
	goto L50
L50:
	;
	if v141 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v221 = int32(200)
	goto L53
L52:
	;
	v221 = v220
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[6])) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[7]))
	v231 = F_WaitLatch(m, v228, int32(41), v221, int32(83886091))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	if v231&int32(1) == int32(0) {
		v93 = v203
		v94 = v204
		goto L20
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[7]))
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v243
	v248 = base.AtomicRmwOr32(m, v243, int32(_a_F_SyncReplicationSlots_6), v243)
	goto L56
L56:
	;
	v93 = v203
	v94 = v204
	goto L20
L57:
	;
	if v94 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_list_free_deep(m, v94)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_ReplicationSlotCleanup(m, int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[8]))
	v272 = base.AtomicRmwXchg32(m, v269, int32(16), int32(1))
	if v272 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_s_lock(m, v269+int32(16), int32(_a_F_SyncReplicationSlots_7))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v282 = int32(_a_F_SyncReplicationSlots_8)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = int32(-1)
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+5)) = uint8(v286)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[8]))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v289)+16)), uint32(v286))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v93
	*(*uint8)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[9])) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_cancel_before_shmem_exit(m, int32(1092), v55)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L5
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0])) = v52
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v53
	m.G0 = v15 + int32(192)
	return
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_slotsync_failure_callback(m, v15, v55)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v55
	F_pg_re_throw(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	goto L4
L71:
	;
	v350 = int32(v346)
	m.G0 = v15
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v15+int32(8) == v356 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	m.ExcPending = 1
	goto L80
L73:
	;
	if v360 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v360 = v358
	goto L76
L75:
	;
	v360 = int32(0)
	goto L76
L76:
	;
	goto L73
L77:
	;
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v15)+184))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v15)+172))
	v20 = v363
	v21 = v362
	v22 = v364
	v23 = v360
	v24 = v352
	v29 = v361
	goto L1
L78:
	;
	goto L79
L79:
	;
	F___wasm_longjmp(m, v353, v352)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	return
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
