package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_standby_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
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
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int64
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v517 int64
	_ = v517
	var v519 int64
	_ = v519
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int64
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int64
	_ = v581
	var v584 int64
	_ = v584
	var v588 int32
	_ = v588
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
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+48)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v22, v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = v21 & int32(240)
	switch v28 - int32(16) {
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
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L200
	}
L4:
	;
	m.G0 = v16 + int32(16)
	return
L5:
	;
	if v28 != 0 {
		goto L3
	} else {
		goto L199
	}
L6:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	v34 = m.G0
	v36 = v34 - int32(176)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v38 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v36 + int32(176)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v607 = m.G0
	v609 = v607 - int32(16)
	m.G0 = v609
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v605)+8))
	if v611 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L8:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v274
	if base.Ui32(v274) < base.Ui32(int32(3)) {
		goto L82
	} else {
		goto L83
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v46 = int32(0)
	if base.B2i32(base.Ui32(v41) < base.Ui32(int32(3)))|base.B2i32(v46 <= v44-v41) == v46 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_SnapBuildSerialize(m, v18, v31)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L81
	}
L12:
	;
	v53 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v82 == v44 {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+68)) = uint32(v31)
	v57 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+64)) = uint32(v57)
	F_errmsg_internal(m, int32(_a_F_standby_decode_0), v36-int32(-64))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	F_SnapBuildWaitSnapshot(m, v33, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v64
	F_errdetail_internal(m, int32(_a_F_standby_decode_1), v36+int32(48))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1281), int32(_a_F_standby_decode_3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
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
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	if base.Ui64(v84) <= base.Ui64(v31) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+36)))
	if v128 != 0 {
		v133 = v38
		goto L37
	} else {
		goto L38
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v31 + int64(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v90 = v89
	goto L28
L27:
	;
	v90 = v44
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v92
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_standby_decode[0]))
	if v101 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v104 = int32(14)
	goto L31
L30:
	;
	v104 = int32(15)
	goto L31
L31:
	;
	v106 = F_errstart(m, v104, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v106 == int32(0) {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+84)) = uint32(v31)
	v112 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+80)) = uint32(v112)
	F_errmsg(m, int32(_a_F_standby_decode_4), v36+int32(80))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v121 = F_errdetail(m, int32(_a_F_standby_decode_5), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1319), int32(_a_F_standby_decode_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L7
L37:
	;
	switch v133 + int32(1) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	default:
		goto L8
	}
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+37)))
	if v129 != 0 {
		v133 = v38
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v130 = F_SnapBuildRestore(m, v18, v31)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v130 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v133 = v132
	goto L37
L42:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v223 = int32(3)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if base.B2i32(base.Ui32(v222) < base.Ui32(v223))|base.B2i32(base.Ui32(v225) < base.Ui32(v223)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v175 = int32(3)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if base.B2i32(base.Ui32(v174) < base.Ui32(v175))|base.B2i32(base.Ui32(v177) < base.Ui32(v175)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v141
	v145 = F_errstart(m, int32(15), v136)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v145 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+116)) = uint32(v31)
	v149 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+112)) = uint32(v149)
	F_errmsg(m, int32(_a_F_standby_decode_6), v36+int32(112))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	F_SnapBuildWaitSnapshot(m, v33, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v156
	v163 = F_errdetail(m, int32(_a_F_standby_decode_7), v36+int32(96))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1369), int32(_a_F_standby_decode_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L8
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v189
	v193 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	if v174-v177 <= int32(0) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v177) < base.Ui32(v174) {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	goto L8
L58:
	;
	goto L53
L59:
	;
	if v193 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+148)) = uint32(v31)
	v197 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+144)) = uint32(v197)
	F_errmsg(m, int32(_a_F_standby_decode_8), v36+int32(144))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	F_SnapBuildWaitSnapshot(m, v33, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L66
	}
L63:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v204
	v211 = F_errdetail(m, int32(_a_F_standby_decode_7), v36+int32(128))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1393), int32(_a_F_standby_decode_3))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	goto L8
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_standby_decode[0]))
	if v242 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	if v222-v225 <= int32(0) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v225) < base.Ui32(v222) {
		goto L8
	} else {
		goto L72
	}
L71:
	;
	goto L8
L72:
	;
	goto L67
L73:
	;
	v245 = int32(14)
	goto L75
L74:
	;
	v245 = int32(15)
	goto L75
L75:
	;
	v247 = F_errstart(m, v245, int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v247 == int32(0) {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+164)) = uint32(v31)
	v253 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+160)) = uint32(v253)
	F_errmsg(m, int32(_a_F_standby_decode_4), v36+int32(160))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v262 = F_errdetail(m, int32(_a_F_standby_decode_9), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1416), int32(_a_F_standby_decode_3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L8
L81:
	;
	goto L8
L82:
	;
	v454 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+16))
	if base.B2i32(v456 == v454)|base.B2i32(v456 == v455+int32(12)) == v454 {
		goto L128
	} else {
		goto L129
	}
L83:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v282 = F_MemoryContextAlloc(m, v278, v279<<(uint(int32(2))%32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	if v284 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v338 = v326 << (uint(int32(2)) % 32)
	if v338 != 0 {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v326 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v288 = int32(0)
	v291 = v288
	v292 = v288
	v297 = v284
	goto L89
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v291<<(uint(int32(2))%32))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if int32(0) <= v307-v308 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v326 = v319
	goto L85
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282+v292<<(uint(int32(2))%32)))) = v307
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v319 = v292 + int32(1)
	v320 = v316
	goto L93
L92:
	;
	v319 = v292
	v320 = v297
	goto L93
L93:
	;
	v322 = v291 + int32(1)
	if base.Ui32(v322) < base.Ui32(v320) {
		v291 = v322
		v292 = v319
		v297 = v320
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	base.MemoryCopy(m, v339, v282, v338)
	goto L97
L96:
	;
	goto L97
L97:
	;
	v343 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v343 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v326
	*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v345
	F_errmsg_internal(m, int32(_a_F_standby_decode_10), v36+int32(32))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v326
	F_pfree(m, v282)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(898), int32(_a_F_standby_decode_11))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	if v365 == int32(0) {
		goto L82
	} else {
		goto L105
	}
L105:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v374 = int32(0)
	goto L107
L106:
	;
	v403 = v365 - v402
	if int32(0) < v403 {
		goto L117
	} else {
		goto L118
	}
L107:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v368+v374<<(uint(int32(2))%32))))
	if base.B2i32(base.Ui32(v370) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v389) < base.Ui32(int32(3))) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v402 = v365
	goto L106
L109:
	;
	v400 = v374 + int32(1)
	if v400 != v365 {
		v374 = v400
		goto L107
	} else {
		goto L115
	}
L110:
	;
	if v389-v370 < int32(0) {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v370) <= base.Ui32(v389) {
		v402 = v374
		goto L106
	} else {
		goto L114
	}
L113:
	;
	v402 = v374
	goto L106
L114:
	;
	goto L109
L115:
	;
	goto L108
L116:
	;
	v421 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L122
	}
L117:
	;
	v407 = v403 << (uint(int32(2)) % 32)
	if v407 == int32(0) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_pfree(m, v368)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	base.MemoryCopy(m, v368, v368+v402<<(uint(int32(2))%32), v407)
	goto L116
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(0)
	goto L116
L122:
	;
	if v421 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v423
	F_errmsg_internal(m, int32(_a_F_standby_decode_12), v36+int32(16))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v403
	goto L82
L126:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(935), int32(_a_F_standby_decode_11))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v456-int32(16))))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v469 = v468
	goto L130
L129:
	;
	v469 = v454
	goto L130
L130:
	;
	if v469 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v473 = v472
	goto L133
L132:
	;
	v473 = v469
	goto L133
L133:
	;
	v476 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v476 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v479
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v478
	F_errmsg_internal(m, int32(_a_F_standby_decode_13), v36)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v493 = m.G0
	v495 = v493 - int32(16)
	m.G0 = v495
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_standby_decode[1]))
	v501 = base.AtomicRmwXchg32(m, v498, int32(0), int32(1))
	if v501 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	F_errfinish(m, int32(_a_F_standby_decode_2), int32(1189), int32(_a_F_standby_decode_14))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	F_s_lock(m, v498, int32(_a_F_standby_decode_15))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v505 = int32(3)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v498)+100))
	if base.B2i32(base.Ui32(v473) < base.Ui32(v505))|base.B2i32(base.Ui32(v507) < base.Ui32(v505)) == int32(0) {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	goto L142
L144:
	;
	m.G0 = v495 + int32(16)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v561 < int32(2) {
		goto L7
	} else {
		goto L161
	}
L145:
	;
	v554 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v498))), uint32(v554))
	goto L144
L146:
	;
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v498)+120))
	if base.Ui64(v517) < base.Ui64(v31) {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	if int32(0) < v473-v507 {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	if base.Ui32(v473) <= base.Ui32(v507) {
		goto L145
	} else {
		goto L151
	}
L150:
	;
	goto L145
L151:
	;
	goto L146
L152:
	;
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v498)+240))
	if v519 != int64(0) {
		goto L145
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v498)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v498)+236)) = v473
	v548 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v498))), uint32(v548))
	F_LogicalConfirmReceivedLocation(m, v517)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L160
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v498)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v498)+236)) = v473
	v524 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v498))), uint32(v524))
	v529 = F_errstart(m, int32(14), v524)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v529 == int32(0) {
		goto L144
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v473
	*(*uint32)(unsafe.Add(mBase, uint32(v495)+8)) = uint32(v31)
	v536 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v495)+4)) = uint32(v536)
	F_errmsg_internal(m, int32(_a_F_standby_decode_16), v495)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_standby_decode_17), int32(1811), int32(_a_F_standby_decode_18))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	goto L144
L160:
	;
	goto L144
L161:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	if v565 != v564+int32(4) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v572 = v565 - int32(188)
	goto L164
L163:
	;
	v572 = int32(0)
	goto L164
L164:
	;
	if v565 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v574 = v572
	goto L167
L166:
	;
	v574 = int32(0)
	goto L167
L167:
	;
	if v574 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v574)+48))
	if v575 == int64(0) {
		goto L7
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v580)+136))
	if v581 == int64(0) {
		goto L7
	} else {
		goto L173
	}
L171:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v31, v575)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	goto L7
L173:
	;
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	if v584 == int64(0) {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v31, v584)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L7
L176:
	;
	m.G0 = v609 + int32(16)
	goto L4
L177:
	;
	v615 = v605 + int32(4)
	if v611 == v615 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v619 = v611
	goto L179
L179:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	v634 = v619 - int32(184)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	if base.B2i32(base.Ui32(v606) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v635) < base.Ui32(int32(3))) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L176
L181:
	;
	v647 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L187
	}
L182:
	;
	if v635-v606 < int32(0) {
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	if base.Ui32(v606) <= base.Ui32(v635) {
		goto L176
	} else {
		goto L186
	}
L185:
	;
	goto L176
L186:
	;
	goto L181
L187:
	;
	if v647 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	*(*int32)(unsafe.Add(mBase, uint32(v609))) = v649
	F_errmsg_internal(m, int32(_a_F_standby_decode_19), v609)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v660 = v619 - int32(188)
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	if v661&int32(16) != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	F_errfinish(m, int32(_a_F_standby_decode_20), int32(3187), int32(_a_F_standby_decode_21))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v605)+84))
	m.T0[v665].(func(*base.Module, int32, int32, int64))(m, v605, v660, int64(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	F_ReorderBufferCleanupTXN(m, v605, v660)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L197
	}
L196:
	;
	goto L195
L197:
	;
	if v632 != v615 {
		v619 = v632
		goto L179
	} else {
		goto L198
	}
L198:
	;
	goto L180
L199:
	;
	goto L4
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v28
	F_errmsg_internal(m, int32(_a_F_standby_decode_22), v16)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_standby_decode_23), int32(398), int32(_a_F_standby_decode_24))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
