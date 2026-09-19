package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_standby_decode(m *base.Module, l0 int32, l1 int32) {
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
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int64
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int64
	_ = v248
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int64
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v514 int64
	_ = v514
	var v516 int64
	_ = v516
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int64
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v581 int64
	_ = v581
	var v585 int32
	_ = v585
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v21, v22, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = v20 & int32(240)
	switch v27 - int32(16) {
	case 0:
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L3
	case 16:
		goto L4
	default:
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L191
	}
L4:
	;
	m.G0 = v15 + int32(16)
	return
L5:
	;
	if v27 != 0 {
		goto L3
	} else {
		goto L190
	}
L6:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v33 = m.G0
	v35 = v33 - int32(176)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v37 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v35 + int32(176)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v603 = m.G0
	v605 = v603 - int32(16)
	m.G0 = v605
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v601)+8))
	if v607 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L8:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v269
	if base.Ui32(v269) < base.Ui32(int32(3)) {
		goto L77
	} else {
		goto L78
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v45 = int32(0)
	if base.B2i32(base.Ui32(v40) < base.Ui32(int32(3)))|base.B2i32(v45 <= v43-v40) == v45 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_SnapBuildSerialize(m, v17, v30)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L76
	}
L12:
	;
	v52 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v81 == v43 {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+68)) = uint32(v30)
	v56 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+64)) = uint32(v56)
	F_errmsg_internal(m, int32(_a_F_standby_decode_0), v35-int32(-64))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	F_SnapBuildWaitSnapshot(m, v32, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v63
	F_errdetail_internal(m, int32(_a_F_standby_decode_1), v35+int32(48))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1277), int32(_a_F_standby_decode_3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L8
L23:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	if base.Ui64(v83) <= base.Ui64(v30) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+36)))
	if v121 != 0 {
		v126 = v37
		goto L34
	} else {
		goto L35
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v30 + int64(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v89 = v88
	goto L28
L27:
	;
	v89 = v43
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v91
	v99 = F_errstart(m, int32(15), v92)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v99 == int32(0) {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+84)) = uint32(v30)
	v105 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+80)) = uint32(v105)
	F_errmsg(m, int32(_a_F_standby_decode_4), v35+int32(80))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errdetail(m, int32(_a_F_standby_decode_5), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1315), int32(_a_F_standby_decode_3))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L7
L34:
	;
	switch v126 + int32(1) {
	case 0:
		goto L41
	case 1:
		goto L40
	default:
		v217 = v126
		goto L39
	}
L35:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)))
	if v122 != 0 {
		v126 = v37
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v123 = F_SnapBuildRestore(m, v17, v30)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v123 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v126 = v125
	goto L34
L39:
	;
	if v217 != int32(1) {
		goto L8
	} else {
		goto L65
	}
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v168))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v167)) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v134
	v138 = F_errstart(m, int32(15), v129)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v138 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+132)) = uint32(v30)
	v142 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+128)) = uint32(v142)
	F_errmsg(m, int32(_a_F_standby_decode_6), v35+int32(128))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_SnapBuildWaitSnapshot(m, v32, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+116)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v149
	F_errdetail(m, int32(_a_F_standby_decode_7), v35+int32(112))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1365), int32(_a_F_standby_decode_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L8
L50:
	;
	if v180 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v180 = base.B2i32(base.Ui32(v167) <= base.Ui32(v168))
	goto L50
L52:
	;
	goto L53
L53:
	;
	v180 = base.B2i32(v167-v168 <= int32(0))
	goto L50
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v183
	v187 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v217 = v216
	goto L39
L57:
	;
	if v187 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+164)) = uint32(v30)
	v191 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+160)) = uint32(v191)
	F_errmsg(m, int32(_a_F_standby_decode_8), v35+int32(160))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_SnapBuildWaitSnapshot(m, v32, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+148)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v35)+144)) = v198
	F_errdetail(m, int32(_a_F_standby_decode_7), v35+int32(144))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1389), int32(_a_F_standby_decode_3))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L8
L65:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v221))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v220)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v233 == int32(0) {
		goto L8
	} else {
		goto L70
	}
L67:
	;
	v233 = base.B2i32(base.Ui32(v220) <= base.Ui32(v221))
	goto L66
L68:
	;
	goto L69
L69:
	;
	v233 = base.B2i32(v220-v221 <= int32(0))
	goto L66
L70:
	;
	v236 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v242 = F_errstart(m, int32(15), v236)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v242 == int32(0) {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+100)) = uint32(v30)
	v248 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+96)) = uint32(v248)
	F_errmsg(m, int32(_a_F_standby_decode_4), v35+int32(96))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errdetail(m, int32(_a_F_standby_decode_9), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1412), int32(_a_F_standby_decode_3))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L8
L76:
	;
	goto L8
L77:
	;
	v448 = int32(0)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+16))
	if base.B2i32(v450 == v448)|base.B2i32(v450 == v449+int32(12)) == v448 {
		goto L122
	} else {
		goto L123
	}
L78:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v277 = F_MemoryContextAlloc(m, v273, v274<<(uint(int32(2))%32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	if v279 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v331 = v322 << (uint(int32(2)) % 32)
	if v331 != 0 {
		goto L90
	} else {
		goto L91
	}
L81:
	;
	v322 = int32(0)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v283 = int32(0)
	v288 = v283
	v289 = v283
	v292 = v279
	goto L84
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v288<<(uint(int32(2))%32))))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(0) <= v301-v302 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v322 = v313
	goto L80
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277+v289<<(uint(int32(2))%32)))) = v301
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v313 = v289 + int32(1)
	v314 = v310
	goto L88
L87:
	;
	v313 = v289
	v314 = v292
	goto L88
L88:
	;
	v316 = v288 + int32(1)
	if base.Ui32(v316) < base.Ui32(v314) {
		v288 = v316
		v289 = v313
		v292 = v314
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	base.MemoryCopy(m, v332, v277, v331)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v336 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v336 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v322
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v338
	F_errmsg_internal(m, int32(_a_F_standby_decode_10), v35+int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v322
	F_pfree(m, v277)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(894), int32(_a_F_standby_decode_11))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v358 == int32(0) {
		goto L77
	} else {
		goto L100
	}
L100:
	;
	v365 = int32(0)
	goto L101
L101:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v375 = int32(2)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374+v365<<(uint(v375)%32))))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if base.B2i32(base.Ui32(v375) < base.Ui32(v379))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v378)) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v398 = v392 - v396
	if int32(0) < v398 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	goto L102
L104:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v391 != 0 {
		v396 = v365
		goto L103
	} else {
		goto L108
	}
L105:
	;
	v391 = base.B2i32(base.Ui32(v379) <= base.Ui32(v378))
	goto L104
L106:
	;
	goto L107
L107:
	;
	v391 = base.B2i32(int32(0) <= v378-v379)
	goto L104
L108:
	;
	v394 = v365 + int32(1)
	if base.Ui32(v394) < base.Ui32(v392) {
		v365 = v394
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v396 = v394
	goto L103
L110:
	;
	v416 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L116
	}
L111:
	;
	v402 = v398 << (uint(int32(2)) % 32)
	if v402 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_pfree(m, v397)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	base.MemoryCopy(m, v397, v397+v396<<(uint(int32(2))%32), v402)
	goto L110
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = int32(0)
	goto L110
L116:
	;
	if v416 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v398
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v418
	F_errmsg_internal(m, int32(_a_F_standby_decode_12), v35+int32(16))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v398
	goto L77
L120:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(931), int32(_a_F_standby_decode_11))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v450-int32(16))))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v463 = v462
	goto L124
L123:
	;
	v463 = v448
	goto L124
L124:
	;
	if v463 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v467 = v466
	goto L127
L126:
	;
	v467 = v463
	goto L127
L127:
	;
	v470 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v470 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v472
	F_errmsg_internal(m, int32(_a_F_standby_decode_13), v35)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v487 = m.G0
	v489 = v487 - int32(16)
	m.G0 = v489
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_standby_decode[0]))
	v495 = base.AtomicRmwXchg32(m, v492, int32(0), int32(1))
	if v495 != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1185), int32(_a_F_standby_decode_14))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	F_s_lock(m, v492, int32(_a_F_standby_decode_15), int32(1688), int32(_a_F_standby_decode_16))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v492)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v501))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v467)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	goto L136
L138:
	;
	m.G0 = v489 + int32(16)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v558 < int32(2) {
		goto L7
	} else {
		goto L153
	}
L139:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v492)+240)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v492)+236)) = v467
	v549 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v492))), uint32(v549))
	F_LogicalConfirmReceivedLocation(m, v514)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L152
	}
L140:
	;
	v544 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v492))), uint32(v544))
	goto L138
L141:
	;
	if v513 != 0 {
		goto L140
	} else {
		goto L145
	}
L142:
	;
	v513 = base.B2i32(base.Ui32(v467) <= base.Ui32(v501))
	goto L141
L143:
	;
	goto L144
L144:
	;
	v513 = base.B2i32(v467-v501 <= int32(0))
	goto L141
L145:
	;
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v492)+120))
	if base.Ui64(v30) <= base.Ui64(v514) {
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v492)+240))
	if v516 != int64(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v492)+240)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v492)+236)) = v467
	v521 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v492))), uint32(v521))
	v526 = F_errstart(m, int32(14), v521)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v526 == int32(0) {
		goto L138
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v467
	*(*uint32)(unsafe.Add(mBase, uint32(v489)+8)) = uint32(v30)
	v533 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v489)+4)) = uint32(v533)
	F_errmsg_internal(m, int32(_a_F_standby_decode_17), v489)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_standby_decode_15), int32(1731), int32(_a_F_standby_decode_16))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	goto L138
L152:
	;
	goto L138
L153:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	if v562 != v561+int32(4) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v569 = v562 - int32(188)
	goto L156
L155:
	;
	v569 = int32(0)
	goto L156
L156:
	;
	if v562 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v571 = v569
	goto L159
L158:
	;
	v571 = int32(0)
	goto L159
L159:
	;
	if v571 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v571)+48))
	if v572 == int64(0) {
		goto L7
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v577)+136))
	if v578 == int64(0) {
		goto L7
	} else {
		goto L165
	}
L163:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v30, v572)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L7
L165:
	;
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	if v581 == int64(0) {
		goto L7
	} else {
		goto L166
	}
L166:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v30, v581)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	goto L7
L168:
	;
	m.G0 = v605 + int32(16)
	goto L4
L169:
	;
	v611 = v601 + int32(4)
	if v607 == v611 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v614 = v607
	goto L171
L171:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v627 = v614 - int32(184)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v602))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v628)) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L168
L173:
	;
	if v640 == int32(0) {
		goto L168
	} else {
		goto L177
	}
L174:
	;
	v640 = base.B2i32(base.Ui32(v628) < base.Ui32(v602))
	goto L173
L175:
	;
	goto L176
L176:
	;
	v640 = int32(base.Ui32(v628-v602) >> (uint(int32(31)) % 32))
	goto L173
L177:
	;
	v645 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	if v645 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	*(*int32)(unsafe.Add(mBase, uint32(v605))) = v647
	F_errmsg_internal(m, int32(_a_F_standby_decode_18), v605)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v658 = v614 - int32(188)
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658))))
	if v659&int32(16) != 0 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	F_errfinish(m, int32(_a_F_standby_decode_19), int32(3142), int32(_a_F_standby_decode_20))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v601)+84))
	m.T0[v663].(func(*base.Module, int32, int32, int64))(m, v601, v658, int64(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	F_ReorderBufferCleanupTXN(m, v601, v658)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	if v625 != v611 {
		v614 = v625
		goto L171
	} else {
		goto L189
	}
L189:
	;
	goto L172
L190:
	;
	goto L4
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v27
	F_errmsg_internal(m, int32(_a_F_standby_decode_21), v15)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_standby_decode_22), int32(397), int32(_a_F_standby_decode_23))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
