package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackendMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int64
	_ = v111
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int64
	_ = v208
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int64
	_ = v468
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	v8 = *(*int32)(unsafe.Add(mBase, _consts[881]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = m.G0
	v12 = v10 - int32(432)
	m.G0 = v12
	F_ReserveExternalFD(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[882]))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_pg_usleep(m, v17*int32(1000000))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[172])) = uint8(v25)
	v27 = int32(4515120)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
	v33 = F_pq_init(m, v8)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	*(*int32)(unsafe.Add(mBase, _consts[580])) = v33
	*(*int32)(unsafe.Add(mBase, _consts[171])) = int32(2)
	v42 = int32(756936)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+292)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+276)) = v42
	v47 = int32(1148)
	v49 = m.G0
	v51 = v49 - int32(144)
	m.G0 = v51
	switch int32(1150) {
	case 0, 2:
		v61 = v47
		goto L9
	default:
		goto L10
	}
L8:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[623])) = v88
	*(*int32)(unsafe.Add(mBase, _consts[624])) = v88
	v98 = v88
	goto L22
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v61
	F_sigemptyset(m, v51+int32(8))
	mBase = m.M
	goto L12
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[622])) = v47
	v61 = int32(4730)
	goto L9
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+136)) = int32(268435456)
	v73 = v51 + int32(4)
	goto L16
L14:
	;
	m.G0 = v51 + int32(144)
	goto L8
L16:
	;
	goto L17
L17:
	;
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v84 = F___memcpy(m, int32(4682692), v73, int32(140))
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L14
L21:
	;
	F_sigprocmask(m, int32(4422600), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	v100 = int32(40)
	v101 = v98 * v100
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[626]))) = uint8(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[627]))) = v98
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[628]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[629]))) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[630]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[631]))) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[632]))) = uint8(v104)
	v130 = v98 | int32(1)
	v132 = v130 * v100
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[626]))) = uint8(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[627]))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[628]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[629]))) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[630]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[631]))) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[632]))) = uint8(v104)
	v161 = v98 | int32(2)
	v163 = v161 * v100
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[626]))) = uint8(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[627]))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[628]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[629]))) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[630]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[631]))) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[632]))) = uint8(v104)
	if base.B2i32(v98 == int32(20)) == v104 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[633])) = uint8(v229)
	F_pqsignal_be(m, int32(14), int32(1785))
	mBase = m.M
	goto L21
L24:
	;
	v196 = v98 | int32(3)
	v198 = v196 * int32(40)
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[626]))) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[627]))) = v196
	v208 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[628]))) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[629]))) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[630]))) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[631]))) = v201
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[632]))) = uint8(v201)
	v98 = v98 + int32(4)
	goto L22
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+144)) = uint8(v238)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+176)) = uint8(v238)
	v242 = int32(144)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v33)+272))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, _consts[883])))
	v255 = F_pg_getnameinfo_all(m, v33+v242, v244, v12+int32(176), int32(255), v12+v242, int32(32), v252^int32(3))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v310 = F_MemoryContextStrdup(m, v307, v12+int32(176))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L45
	}
L29:
	;
	if v255 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v261 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v261 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v267 = int32(4083728)
	v269 = v255 + int32(1)
	if v269 == int32(0) {
		v289 = v267
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v289 + base.B2i32(v291 == int32(0))
	F_errmsg_internal(m, int32(203993), v12+int32(112))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L43
	}
L34:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	goto L33
L35:
	;
	v273 = v267
	v274 = v269
	goto L36
L36:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v275 == int32(0) {
		v289 = v273
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v289 = v285
	goto L34
L38:
	;
	v279 = v273
	goto L39
L39:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v283 != 0 {
		v279 = v279 + int32(1)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v285 = v279 + int32(2)
	v287 = v274 + int32(1)
	if v287 != 0 {
		v273 = v285
		v274 = v287
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	F_errfinish(m, int32(495374), int32(220), int32(341473))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+276)) = v310
	v314 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v317 = F_MemoryContextStrdup(m, v314, v12+int32(144))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+292)) = v317
	v321 = int32(*(*uint8)(unsafe.Add(mBase, _consts[884])))
	if v321&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v255 != 0 {
		goto L59
	} else {
		goto L60
	}
L48:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+144)))
	v329 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v326 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_errfinish(m, int32(495374), v357, int32(341473))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	if v329 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v329 == int32(0) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v12 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v12 + int32(176)
	F_errmsg(m, int32(175963), v12+int32(96))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v357 = int32(236)
	goto L50
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v12 + int32(176)
	F_errmsg(m, int32(175934), v12+int32(80))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v357 = int32(240)
	goto L50
L58:
	;
	goto L47
L59:
	;
	F_RegisterTimeout(m, int32(0), int32(1149))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L107
	}
L60:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _consts[883])))
	if v364&int32(1) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v370 = v12 + int32(176)
	v371 = int32(659468)
	v375 = m.G0
	v377 = v375 - int32(32)
	v378 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+24)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v377)+16)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v377))) = v378
	v386 = int32(*(*uint8)(unsafe.Add(mBase, _consts[885])))
	if v386 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v457 = F_strlen(m, v12+int32(176))
	mBase = m.M
	if base.Ui32(v457) <= base.Ui32(v454) {
		goto L59
	} else {
		goto L83
	}
L63:
	;
	v454 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, _consts[886])))
	if v390 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v394 = v370
	goto L69
L67:
	;
	goto L68
L68:
	;
	v404 = v371
	v405 = v386
	goto L72
L69:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v400 == v386 {
		v394 = v394 + int32(1)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v454 = v394 - v370
	goto L62
L71:
	;
	goto L70
L72:
	;
	v412 = v377 + int32(base.Ui32(v405)>>(uint(int32(3))%32))&int32(28)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v414 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v413 | v414<<(uint(v405)%32)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	if v418 != 0 {
		v404 = v404 + v414
		v405 = v418
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if v421 == int32(0) {
		v446 = v370
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	v454 = v446 - v370
	goto L62
L76:
	;
	v425 = v370
	v426 = v421
	goto L77
L77:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v377+int32(base.Ui32(v426)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v434)>>(uint(v426)%32))&int32(1) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v446 = v442
	goto L75
L79:
	;
	v446 = v425
	goto L75
L80:
	;
	goto L81
L81:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+1)))
	v442 = v425 + int32(1)
	if v440 != 0 {
		v425 = v442
		v426 = v440
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	v460 = v12 + int32(176)
	v461 = int32(546984)
	v465 = m.G0
	v467 = v465 - int32(32)
	v468 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v467)+24)) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v467)+16)) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v467)+8)) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v467))) = v468
	v476 = int32(*(*uint8)(unsafe.Add(mBase, _consts[887])))
	if v476 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if base.Ui32(v457) <= base.Ui32(v544) {
		goto L59
	} else {
		goto L105
	}
L85:
	;
	v544 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, _consts[888])))
	if v480 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v484 = v460
	goto L91
L89:
	;
	goto L90
L90:
	;
	v494 = v461
	v495 = v476
	goto L94
L91:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v490 == v476 {
		v484 = v484 + int32(1)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v544 = v484 - v460
	goto L84
L93:
	;
	goto L92
L94:
	;
	v502 = v467 + int32(base.Ui32(v495)>>(uint(int32(3))%32))&int32(28)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v504 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v503 | v504<<(uint(v495)%32)
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	if v508 != 0 {
		v494 = v494 + v504
		v495 = v508
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v511 == int32(0) {
		v536 = v460
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v544 = v536 - v460
	goto L84
L98:
	;
	v515 = v460
	v516 = v511
	goto L99
L99:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v467+int32(base.Ui32(v516)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v524)>>(uint(v516)%32))&int32(1) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v536 = v532
	goto L97
L101:
	;
	v536 = v515
	goto L97
L102:
	;
	goto L103
L103:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	v532 = v515 + int32(1)
	if v530 != 0 {
		v515 = v532
		v516 = v530
		goto L99
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v550 = F_MemoryContextStrdup(m, v547, v12+int32(176))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+280)) = v550
	goto L59
L107:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[889]))
	F_enable_timeout_after(m, int32(0), v560*int32(1000))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L112
L110:
	;
	v588 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[890])) = uint8(v588)
	goto L117
L111:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+uint32(_consts[891]))))
	v586 = v585
	goto L110
L112:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	v576 = *(*int32)(unsafe.Add(mBase, _consts[578]))
	if v574 < v576 {
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v586 = int32(-1)
	goto L110
L114:
	;
	v578 = F_pq_recvbuf(m)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v578 == int32(0) {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	if v586 == int32(-1) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	F_InitProcess(m)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L201
	}
L119:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L197
	}
L120:
	;
	if v586 == int32(22) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, _consts[892])))
	if v595 != int32(1) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v613 = int32(0)
	v615 = F_ProcessStartupPacket(m, v33, v613, v613)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L129
	}
L124:
	;
	v600 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v600 == int32(0) {
		goto L119
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(447406), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(495374), int32(477), int32(231999))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L119
L129:
	;
	if v615 != 0 {
		goto L119
	} else {
		goto L130
	}
L130:
	;
	switch v9 - int32(1) {
	case 0:
		goto L138
	case 1:
		goto L136
	case 2:
		goto L135
	case 3:
		goto L137
	case 4:
		goto L134
	default:
		goto L133
	}
L131:
	;
	F_errdetail(m, int32(575641), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L194
	}
L132:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L189
	}
L133:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L162
	}
L134:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L158
	}
L135:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L154
	}
L136:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L150
	}
L137:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, _consts[233])))
	if v636 == int32(0) {
		goto L132
	} else {
		goto L143
	}
L138:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(233867), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(495374), int32(317), int32(341473))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, _consts[293])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(141694), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v640 == int32(1) {
		goto L131
	} else {
		goto L147
	}
L147:
	;
	F_errdetail(m, int32(650326), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(495374), int32(336), int32(341473))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(244086), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(495374), int32(341), int32(341473))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
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
	F_errcode(m, int32(50463173))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(411374), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(495374), int32(346), int32(341473))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(12485))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(22873), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(495374), int32(351), int32(341473))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_sigprocmask(m, int32(4422472), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_initStringInfo(m, v12+int32(128))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, _consts[707])))
	if v725 == int32(1) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	goto L170
L167:
	;
	goto L168
L168:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v33)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v747
	F_appendStringInfo(m, v12+int32(128), int32(737852), v12+int32(32))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L174
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v737
	F_appendStringInfo(m, v12+int32(128), int32(737852), v12+int32(48))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L173
	}
L170:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _consts[893]))
	goto L172
L172:
	;
	goto L169
L173:
	;
	goto L168
L174:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	if v757 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v756
	F_appendStringInfo(m, v12+int32(128), int32(737852), v12+int32(16))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v33)+276))
	F_appendStringInfoString(m, v12+int32(128), v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	goto L177
L179:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v33)+292))
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771))))
	if v772 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v771
	F_appendStringInfo(m, v12+int32(128), int32(675277), v12)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	if v779 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_pfree(m, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v784 = F_GetBackendTypeDesc(m, v783)
	mBase = m.M
	goto L187
L186:
	;
	goto L187
L187:
	;
	goto L184
L188:
	;
	m.G0 = v12 + int32(432)
	goto L118
L189:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errmsg(m, int32(141645), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errdetail(m, int32(649505), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(495374), int32(324), int32(341473))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(64)
	F_errhint(m, int32(607460), v12-int32(-64))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(495374), int32(331), int32(341473))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_sigprocmask(m, int32(4422472), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	v843 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v843
	v846 = *(*int32)(unsafe.Add(mBase, _consts[580]))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+360))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v846)+364))
	F_PostgresMain(m, v847, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
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
func F_BackgroundWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int64
	_ = v343
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int64
	_ = v440
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v539 int32
	_ = v539
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v594 int32
	_ = v594
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int64
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v16 = v3
	v17 = v3
	v18 = int32(-1)
	v19 = v3
	v20 = v11
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v18 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v865 = int32(m.ExcTag)
	v866 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v865 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L7:
	;
	v25 = v20 - int32(160)
	m.G0 = v25
	if l0 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v604 = v16
	v607 = v19
	v608 = v20
	v609 = v17
	goto L9
L9:
	;
	if v609 != 0 {
		goto L156
	} else {
		goto L157
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	v51 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v53 = F_MemoryContextAlloc(m, v51, int32(1460))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	F_errmsg_internal(m, int32(12370), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	F_errfinish(m, int32(495141), int32(725), int32(278479))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	goto L18
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v56 = F__emscripten_memcpy_bulkmem(m, v53, l0, int32(1460))
	mBase = m.M
	goto L20
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	F_MemoryContextDelete(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[198])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[617])) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	if v56 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[613])) = int32(0)
	goto L23
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[618]))
	if int32(0) < v80 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v78 = F_GetBackendTypeDesc(m, v77)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	F_pg_usleep(m, v80*int32(1000000))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		v863 = v25
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v56)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v96 = v89 & int32(2)
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v97 = int32(915)
	goto L35
L34:
	;
	v97 = int32(-2)
	goto L35
L35:
	;
	v99 = m.G0
	v101 = v99 - int32(144)
	m.G0 = v101
	switch v97 + int32(2) {
	case 0, 2:
		v111 = v97
		goto L37
	default:
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	if v96 != 0 {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v111
	F_sigemptyset(m, v101+int32(8))
	mBase = m.M
	goto L40
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[619])) = v97
	v111 = int32(4730)
	goto L37
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+136)) = int32(268435456)
	v123 = v101 + int32(4)
	goto L44
L42:
	;
	m.G0 = v101 + int32(144)
	goto L36
L44:
	;
	goto L45
L45:
	;
	if v123 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v134 = F___memcpy(m, int32(4680872), v123, int32(140))
	mBase = m.M
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L42
L49:
	;
	v143 = int32(917)
	goto L51
L50:
	;
	v143 = int32(-2)
	goto L51
L51:
	;
	v145 = m.G0
	v147 = v145 - int32(144)
	m.G0 = v147
	switch v143 + int32(2) {
	case 0, 2:
		v157 = v143
		goto L53
	default:
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	if v96 != 0 {
		goto L65
	} else {
		goto L66
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v157
	F_sigemptyset(m, v147+int32(8))
	mBase = m.M
	goto L56
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[620])) = v143
	v157 = int32(4730)
	goto L53
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+136)) = int32(268435456)
	v169 = v147 + int32(4)
	goto L60
L58:
	;
	m.G0 = v147 + int32(144)
	goto L52
L60:
	;
	goto L61
L61:
	;
	if v169 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v180 = F___memcpy(m, int32(4681992), v169, int32(140))
	mBase = m.M
	goto L64
L63:
	;
	goto L64
L64:
	;
	goto L58
L65:
	;
	v189 = int32(919)
	goto L67
L66:
	;
	v189 = int32(-2)
	goto L67
L67:
	;
	v191 = m.G0
	v193 = v191 - int32(144)
	m.G0 = v193
	switch v189 + int32(2) {
	case 0, 2:
		v203 = v189
		goto L69
	default:
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v233 = int32(923)
	v235 = m.G0
	v237 = v235 - int32(144)
	m.G0 = v237
	switch int32(925) {
	case 0, 2:
		v247 = v233
		goto L82
	default:
		goto L83
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v203
	F_sigemptyset(m, v193+int32(8))
	mBase = m.M
	goto L72
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[621])) = v189
	v203 = int32(4730)
	goto L69
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+136)) = int32(268435456)
	v215 = v193 + int32(4)
	goto L76
L74:
	;
	m.G0 = v193 + int32(144)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if v215 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v226 = F___memcpy(m, int32(4681712), v215, int32(140))
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L74
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v277 = int32(-2)
	v279 = m.G0
	v281 = v279 - int32(144)
	m.G0 = v281
	switch int32(0) {
	case 0, 2:
		v291 = v277
		goto L95
	default:
		goto L96
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v247
	F_sigemptyset(m, v237+int32(8))
	mBase = m.M
	goto L85
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[622])) = v233
	v247 = int32(4730)
	goto L82
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+136)) = int32(268435456)
	v259 = v237 + int32(4)
	goto L89
L87:
	;
	m.G0 = v237 + int32(144)
	goto L81
L89:
	;
	goto L90
L90:
	;
	if v259 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v270 = F___memcpy(m, int32(4682692), v259, int32(140))
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L87
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v320 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[623])) = v320
	*(*int32)(unsafe.Add(mBase, _consts[624])) = v320
	v330 = v320
	goto L108
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v291
	F_sigemptyset(m, v281+int32(8))
	mBase = m.M
	goto L98
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[625])) = v277
	v291 = int32(4730)
	goto L95
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+136)) = int32(268435456)
	v303 = v281 + int32(4)
	goto L102
L100:
	;
	m.G0 = v281 + int32(144)
	goto L94
L102:
	;
	goto L103
L103:
	;
	if v303 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v314 = F___memcpy(m, int32(4680732), v303, int32(140))
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v469 = int32(-2)
	v471 = m.G0
	v473 = v471 - int32(144)
	m.G0 = v473
	switch int32(0) {
	case 0, 2:
		v483 = v469
		goto L114
	default:
		goto L115
	}
L108:
	;
	v332 = int32(40)
	v333 = v330 * v332
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[626]))) = uint8(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[627]))) = v330
	v343 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[628]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[629]))) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[630]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[631]))) = v336
	*(*uint8)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[632]))) = uint8(v336)
	v362 = v330 | int32(1)
	v364 = v362 * v332
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[626]))) = uint8(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[627]))) = v362
	*(*int64)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[628]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[629]))) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[630]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[631]))) = v336
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[632]))) = uint8(v336)
	v393 = v330 | int32(2)
	v395 = v393 * v332
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[626]))) = uint8(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[627]))) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[628]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[629]))) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[630]))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[631]))) = v336
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[632]))) = uint8(v336)
	if base.B2i32(v330 == int32(20)) == v336 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v461 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[633])) = uint8(v461)
	F_pqsignal_be(m, int32(14), int32(1785))
	mBase = m.M
	goto L107
L110:
	;
	v428 = v330 | int32(3)
	v430 = v428 * int32(40)
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[626]))) = uint8(v433)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[627]))) = v428
	v440 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[628]))) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[629]))) = v433
	*(*int64)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[630]))) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[631]))) = v433
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+uint32(_consts[632]))) = uint8(v433)
	v330 = v330 + int32(4)
	goto L108
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v513 = int32(-2)
	v515 = m.G0
	v517 = v515 - int32(144)
	m.G0 = v517
	switch int32(0) {
	case 0, 2:
		v527 = v513
		goto L127
	default:
		goto L128
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v483
	F_sigemptyset(m, v473+int32(8))
	mBase = m.M
	goto L117
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[634])) = v469
	v483 = int32(4730)
	goto L114
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+136)) = int32(268435456)
	v495 = v473 + int32(4)
	goto L121
L119:
	;
	m.G0 = v473 + int32(144)
	goto L113
L121:
	;
	goto L122
L122:
	;
	if v495 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v506 = F___memcpy(m, int32(4682412), v495, int32(140))
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L119
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v56
	v557 = int32(0)
	v559 = m.G0
	v561 = v559 - int32(144)
	m.G0 = v561
	switch int32(2) {
	case 0, 2:
		v571 = v557
		goto L140
	default:
		goto L141
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+4)) = v527
	F_sigemptyset(m, v517+int32(8))
	mBase = m.M
	goto L130
L128:
	;
	*(*int32)(unsafe.Add(mBase, _consts[635])) = v513
	v527 = int32(4730)
	goto L127
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+136)) = int32(268435456)
	v539 = v517 + int32(4)
	goto L134
L132:
	;
	m.G0 = v517 + int32(144)
	goto L126
L134:
	;
	goto L135
L135:
	;
	if v539 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v550 = F___memcpy(m, int32(4682272), v539, int32(140))
	mBase = m.M
	goto L138
L137:
	;
	goto L138
L138:
	;
	goto L132
L139:
	;
	goto L152
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+4)) = v571
	F_sigemptyset(m, v561+int32(8))
	mBase = m.M
	goto L142
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[636])) = v557
	v571 = int32(4730)
	goto L140
L142:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+136)) = int32(268435457)
	v583 = v561 + int32(4)
	goto L147
L145:
	;
	m.G0 = v561 + int32(144)
	goto L139
L147:
	;
	goto L148
L148:
	;
	if v583 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v594 = F___memcpy(m, int32(4682972), v583, int32(140))
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L145
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v11 + int32(4)
	goto L155
L153:
	;
	v604 = v25
	v607 = v53
	v608 = v25
	v609 = int32(0)
	goto L9
L155:
	;
	goto L153
L156:
	;
	v610 = int32(4509772)
	v612 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v612 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[49])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	F_InitProcess(m)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L162
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	F_EmitErrorReport(m)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	F_proc_exit(m, int32(1))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L161
	}
L161:
	;
	goto L3
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	F_BaseInit(m)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	v645 = v607 + int32(1228)
	v646 = m.G0
	v648 = v646 - int32(16)
	m.G0 = v648
	v651 = v607 + int32(204)
	v652 = int32(161591)
	v655 = int32(*(*uint8)(unsafe.Add(mBase, _consts[637])))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	if v656 == int32(0) {
		v675 = v655
		v676 = v656
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v607)+1324))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	m.T0[v832].(func(*base.Module, int32))(m, v849)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L235
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L232
	}
L166:
	;
	m.G0 = v648 + int32(16)
	goto L164
L167:
	;
	if v676-v675 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	goto L167
L169:
	;
	if v655 != v656 {
		v675 = v655
		v676 = v656
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v660 = v651
	v661 = v652
	goto L171
L171:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+1)))
	if v665 == int32(0) {
		v675 = v664
		v676 = v665
		goto L168
	} else {
		goto L173
	}
L172:
	;
	v675 = v664
	v676 = v665
	goto L168
L173:
	;
	v668 = int32(1)
	if v664 == v665 {
		v660 = v660 + v668
		v661 = v661 + v668
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v680 = int32(278460)
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v684 = int32(*(*uint8)(unsafe.Add(mBase, _consts[638])))
	if v684 == int32(0) {
		v703 = v683
		v704 = v684
		goto L179
	} else {
		goto L180
	}
L176:
	;
	goto L177
L177:
	;
	v830 = F_load_external_function(m, v651, v645, int32(1), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L231
	}
L178:
	;
	if v704-v703 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	goto L178
L180:
	;
	if v683 != v684 {
		v703 = v683
		v704 = v684
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v688 = v680
	v689 = v645
	goto L182
L182:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689)+1)))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+1)))
	if v693 == int32(0) {
		v703 = v692
		v704 = v693
		goto L179
	} else {
		goto L184
	}
L183:
	;
	v703 = v692
	v704 = v693
	goto L179
L184:
	;
	v696 = int32(1)
	if v692 == v693 {
		v688 = v688 + v696
		v689 = v689 + v696
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v709 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v832 = v709
	goto L166
L187:
	;
	goto L188
L188:
	;
	v710 = int32(278561)
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, _consts[640])))
	if v714 == int32(0) {
		v733 = v713
		v734 = v714
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if v734-v733 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L190:
	;
	goto L189
L191:
	;
	if v713 != v714 {
		v733 = v713
		v734 = v714
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v718 = v710
	v719 = v645
	goto L193
L193:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+1)))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718)+1)))
	if v723 == int32(0) {
		v733 = v722
		v734 = v723
		goto L190
	} else {
		goto L195
	}
L194:
	;
	v733 = v722
	v734 = v723
	goto L190
L195:
	;
	v726 = int32(1)
	if v722 == v723 {
		v718 = v718 + v726
		v719 = v719 + v726
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	v832 = v739
	goto L166
L198:
	;
	goto L199
L199:
	;
	v740 = int32(278431)
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v744 = int32(*(*uint8)(unsafe.Add(mBase, _consts[642])))
	if v744 == int32(0) {
		v763 = v743
		v764 = v744
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v764-v763 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L201:
	;
	goto L200
L202:
	;
	if v743 != v744 {
		v763 = v743
		v764 = v744
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v748 = v740
	v749 = v645
	goto L204
L204:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+1)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748)+1)))
	if v753 == int32(0) {
		v763 = v752
		v764 = v753
		goto L201
	} else {
		goto L206
	}
L205:
	;
	v763 = v752
	v764 = v753
	goto L201
L206:
	;
	v756 = int32(1)
	if v752 == v753 {
		v748 = v748 + v756
		v749 = v749 + v756
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v769 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v832 = v769
	goto L166
L209:
	;
	goto L210
L210:
	;
	v770 = int32(278423)
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v774 = int32(*(*uint8)(unsafe.Add(mBase, _consts[644])))
	if v774 == int32(0) {
		v793 = v773
		v794 = v774
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if v794-v793 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L212:
	;
	goto L211
L213:
	;
	if v773 != v774 {
		v793 = v773
		v794 = v774
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v778 = v770
	v779 = v645
	goto L215
L215:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+1)))
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+1)))
	if v783 == int32(0) {
		v793 = v782
		v794 = v783
		goto L212
	} else {
		goto L217
	}
L216:
	;
	v793 = v782
	v794 = v783
	goto L212
L217:
	;
	v786 = int32(1)
	if v782 == v783 {
		v778 = v778 + v786
		v779 = v779 + v786
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v832 = v799
	goto L166
L220:
	;
	goto L221
L221:
	;
	v800 = int32(278500)
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v804 = int32(*(*uint8)(unsafe.Add(mBase, _consts[646])))
	if v804 == int32(0) {
		v823 = v803
		v824 = v804
		goto L223
	} else {
		goto L224
	}
L222:
	;
	if v824-v823 != 0 {
		goto L165
	} else {
		goto L230
	}
L223:
	;
	goto L222
L224:
	;
	if v803 != v804 {
		v823 = v803
		v824 = v804
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v808 = v800
	v809 = v645
	goto L226
L226:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809)+1)))
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808)+1)))
	if v813 == int32(0) {
		v823 = v812
		v824 = v813
		goto L223
	} else {
		goto L228
	}
L227:
	;
	v823 = v812
	v824 = v813
	goto L223
L228:
	;
	v816 = int32(1)
	if v812 == v813 {
		v808 = v808 + v816
		v809 = v809 + v816
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	v832 = v827
	goto L166
L231:
	;
	v832 = v830
	goto L166
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v648))) = v645
	F_errmsg_internal(m, int32(423724), v648)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(495141), int32(1355), int32(254320))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v607
	F_proc_exit(m, int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		v863 = v608
		goto L6
	} else {
		goto L236
	}
L236:
	;
	goto L5
L237:
	;
	v870 = int32(v866)
	m.G0 = v863
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	if v11+int32(4) == v877 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	m.ExcPending = 1
	goto L246
L239:
	;
	if v880 != 0 {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	v880 = v879
	goto L242
L241:
	;
	v880 = int32(0)
	goto L242
L242:
	;
	goto L239
L243:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v16 = v881
	v17 = v872
	v18 = v880
	v19 = v882
	v20 = v863
	goto L1
L244:
	;
	goto L245
L245:
	;
	F___wasm_longjmp(m, v873, v872)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	return
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BackgroundWorkerUnblockSignals(m *base.Module) {
	var v4 int32
	_ = v4
	F_sigprocmask(m, int32(4422344), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_BeginImplicitTransactionBlock(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	if v4 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(4)
	} else {
	}
	return
}
func F_BogusGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(669903), v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(492304), int32(315), int32(419241))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F___bswap_16(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(65535)
}
func F_basque_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_SN_close_env(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_basque_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bernoulli_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 float64
	_ = v26
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errmsg(m, int32(566466), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(497479), int32(148), int32(284388))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if base.F32_gt(v7, float32(100)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errmsg(m, int32(566466), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(497479), int32(148), int32(284388))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v12 = base.F64_promote_f32(v7)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errcode(m, int32(403177602))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errmsg(m, int32(566466), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(497479), int32(148), int32(284388))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v19 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v19)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l3
				v26 = base.F64_nearest(base.F64_div(base.F64_mul(v12, float64(4.294967296e+09)), float64(100)))
				if base.F64_lt(v26, float64(1.8446744073709552e+19))&base.F64_ge(v26, float64(0)) != 0 {
					v32 = base.I64_trunc_f64_u(v26)
					v34 = v32
				} else {
					v34 = int64(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v34
				v37 = base.F32_ge(v7, float32(25))
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v37)
				v39 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v39)
				return
			}
		}
	}
}
func F_big5_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
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
	var v172 int32
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v17, v18, v19, int32(36), int32(7))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v172)
	m.G0 = v12 + int32(16)
	return v163 - v16
L4:
	;
	v163 = v16
	v164 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v16
	v29 = v15
	v30 = v19
	goto L7
L7:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if int32(0) <= v37 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v163 = v155
	v164 = v156
	goto L3
L9:
	;
	if int32(0) < v157 {
		v28 = v155
		v29 = v156
		v30 = v157
		goto L7
	} else {
		goto L59
	}
L10:
	;
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v53 = F_pg_encoding_verifymbchar(m, int32(36), v28, v30)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v37)
	v46 = int32(1)
	v155 = v28 + v46
	v156 = v29 + v46
	v157 = v30 - v46
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v53 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v65 = (v60 | v37<<(uint(int32(8))%32)) & int32(65535)
	v67 = v12 + int32(15)
	if base.Ui32(v65) <= base.Ui32(int32(51519)) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v130 != 0 {
		goto L51
	} else {
		goto L52
	}
L25:
	;
	v129 = v127 & int32(65535)
	goto L24
L26:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v122)
	v127 = int32(63)
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v116)
	v127 = v115 | int32(-32640)
	goto L25
L28:
	;
	v110 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v110)
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+2)))
	v127 = v112 | int32(-32640)
	goto L25
L29:
	;
	v109 = int32(2233992)
	goto L28
L30:
	;
	v109 = int32(2233988)
	goto L28
L31:
	;
	v109 = int32(2233984)
	goto L28
L32:
	;
	v109 = int32(2233980)
	goto L28
L33:
	;
	v109 = int32(2233976)
	goto L28
L34:
	;
	v109 = int32(2233972)
	goto L28
L35:
	;
	v99 = F_BinarySearchRange(m, int32(2234000), int32(46), v65)
	mBase = m.M
	if v99 == int32(0) {
		goto L26
	} else {
		goto L50
	}
L36:
	;
	v115 = v94
	v116 = int32(149)
	goto L27
L37:
	;
	switch v65 - int32(51321) {
	case 0:
		v82 = int32(2233840)
		goto L40
	case 1, 3:
		goto L44
	case 2:
		goto L43
	case 4:
		goto L42
	default:
		goto L45
	}
L38:
	;
	goto L39
L39:
	;
	switch v65 - int32(63958) {
	case 0:
		v109 = int32(2233968)
		goto L28
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	case 5:
		goto L30
	case 6:
		goto L29
	default:
		goto L48
	}
L40:
	;
	v83 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v83)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+2)))
	v127 = v85 | int32(-32640)
	goto L25
L41:
	;
	v82 = int32(2233852)
	goto L40
L42:
	;
	v82 = int32(2233848)
	goto L40
L43:
	;
	v82 = int32(2233844)
	goto L40
L44:
	;
	v78 = F_BinarySearchRange(m, int32(2233856), int32(23), v65)
	mBase = m.M
	if v78 != 0 {
		v94 = v78
		goto L36
	} else {
		goto L47
	}
L45:
	;
	if v65 == int32(51362) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L26
L48:
	;
	if v65 != int32(51530) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v94 = int32(17474)
	goto L36
L50:
	;
	v115 = v99
	v116 = int32(150)
	goto L27
L51:
	;
	if v130&int32(254) == int32(246) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L57
	}
L54:
	;
	v135 = int32(157)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	v140 = v29 + int32(1)
	v141 = v137
	goto L56
L55:
	;
	v140 = v29
	v141 = v130
	goto L56
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)) = uint8(v129)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v141)
	v145 = int32(base.Ui32(v129) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)) = uint8(v145)
	v155 = v28 + v53
	v156 = v140 + int32(3)
	v157 = v30 - v53
	goto L9
L57:
	;
	F_report_untranslatable_char(m, int32(36), int32(7), v28, v30)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	goto L8
}
func F_bitlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
L38:
	;
	goto L37
}
func F_bitncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v8 = base.I32_div_s(l2, int32(8))
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	if v158 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	return v156
L3:
	;
	if v70 != 0 {
		v156 = v70
		goto L2
	} else {
		goto L21
	}
L4:
	;
	v70 = int32(0)
	goto L3
L5:
	;
	v44 = v39
	v45 = v40
	v46 = v41
	goto L15
L6:
	;
	if (l0|l1)&int32(3) != 0 {
		v39 = l0
		v40 = l1
		v41 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v32 = l0
	v33 = l1
	v34 = v8
	goto L8
L8:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v16 = l0
	v17 = l1
	v18 = v8
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v21 != v22 {
		v39 = v16
		v40 = v17
		v41 = v18
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v32 = v27
	v33 = v25
	v34 = v29
	goto L8
L12:
	;
	v24 = int32(4)
	v25 = v17 + v24
	v27 = v16 + v24
	v29 = v18 - v24
	if base.Ui32(int32(3)) < base.Ui32(v29) {
		v16 = v27
		v17 = v25
		v18 = v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v32
	v40 = v33
	v41 = v34
	goto L5
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 == v50 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v70 = v49 - v50
	goto L3
L17:
	;
	v52 = int32(1)
	v57 = v46 - v52
	if v57 != 0 {
		v44 = v44 + v52
		v45 = v45 + v52
		v46 = v57
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v71 = int32(0)
	v74 = l2 - v8<<(uint(int32(3))%32)
	if v74 <= v71 {
		v156 = v71
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v8))))
	v79 = int32(128)
	v80 = v78 & v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v8))))
	if v80 != v82&v79 {
		v158 = v80
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v74 == int32(1) {
		v156 = v71
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(1)
	v90 = int32(128)
	v91 = v78 << (uint(v88) % 32) & v90
	if v91 != v82<<(uint(v88)%32)&v90 {
		v158 = v91
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v74 < int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(2)
	v101 = int32(128)
	v102 = v78 << (uint(v99) % 32) & v101
	if v102 != v82<<(uint(v99)%32)&v101 {
		v158 = v102
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v74 == int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(3)
	v112 = int32(128)
	v113 = v78 << (uint(v110) % 32) & v112
	if v113 != v82<<(uint(v110)%32)&v112 {
		v158 = v113
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v74 < int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v121 = int32(4)
	v123 = int32(128)
	v124 = v78 << (uint(v121) % 32) & v123
	if v124 != v82<<(uint(v121)%32)&v123 {
		v158 = v124
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v74 == int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(5)
	v134 = int32(128)
	v135 = v78 << (uint(v132) % 32) & v134
	if v135 != v82<<(uint(v132)%32)&v134 {
		v158 = v135
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v74 < int32(7) {
		v156 = v71
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v143 = int32(6)
	v145 = int32(128)
	v146 = v78 << (uint(v143) % 32) & v145
	if v146 != v82<<(uint(v143)%32)&v145 {
		v158 = v146
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v156 = v71
	goto L2
L36:
	;
	v161 = int32(1)
	goto L38
L37:
	;
	v161 = int32(-1)
	goto L38
L38:
	;
	return v161
}
func F_bittoint4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if base.Ui32(v15) < base.Ui32(int32(33)) {
			v19 = v11 + int32(8)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v22 = int32(base.Ui32(v20) >> (uint(int32(2)) % 32))
			if base.Ui32(v11+v22) <= base.Ui32(v19) {
				v98 = v2
			} else {
				v25 = int32(7)
				v26 = v22 & v25
				if base.Ui32(v22-int32(9)) < base.Ui32(v25) {
					v67 = v19
					v76 = int32(0)
				} else {
					v37 = v19
					v38 = v2
					for {
						v46 = int32(8)
						v47 = v37 + v46
						if v38 != v22&int32(1073741816)-int32(16) {
							v37 = v47
							v38 = v38 + v46
							continue
						} else {
							break
						}
						break
					}
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
					v52 = int32(24)
					v54 = int32(65280)
					v56 = int32(8)
					v67 = v47
					v76 = v51<<(uint(v52)%32) | v51&v54<<(uint(v56)%32) | (int32(base.Ui32(v51)>>(uint(v56)%32))&v54 | int32(base.Ui32(v51)>>(uint(v52)%32)))
				}
				if v26 == int32(0) {
					v98 = v76
				} else {
					v79 = v67
					v80 = v76
					v81 = int32(0)
					for {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
						v91 = v88 | v80<<(uint(int32(8))%32)
						v92 = int32(1)
						v95 = v81 + v92
						if v95 != v26 {
							v79 = v79 + v92
							v80 = v91
							v81 = v95
							continue
						} else {
							break
						}
						break
					}
					v98 = v91
				}
			}
			return int32(base.Ui32(v98) >> (uint(v22<<(uint(int32(3))%32)-v15+int32(-64)) % 32))
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(402097), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492981), int32(1596), int32(557275))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
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
func F_bittoint8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v56 int64
	_ = v56
	var v63 int64
	_ = v63
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int64
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if base.Ui32(v14) < base.Ui32(int32(65)) {
			v18 = v10 + int32(8)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v21 = int32(base.Ui32(v19) >> (uint(int32(2)) % 32))
			if base.Ui32(v10+v21) <= base.Ui32(v18) {
				v115 = v8
			} else {
				v24 = int32(7)
				v25 = v21 & v24
				if base.Ui32(v21-int32(9)) < base.Ui32(v24) {
					v80 = v18
					v87 = v8
				} else {
					v35 = v18
					v36 = int32(0)
					v42 = v8
					for {
						v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+7)))
						v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+5)))
						v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+3)))
						v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
						v47 = int64(16)
						v49 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
						v50 = int64(8)
						v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
						v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
						v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+6)))
						v74 = v43 | ((v44|((v45|((v46|(v42<<(uint(v47)%64)|v49<<(uint(v50)%64)))<<(uint(v47)%64)|v56<<(uint(v50)%64)))<<(uint(v47)%64)|v63<<(uint(v50)%64)))<<(uint(v47)%64) | v70<<(uint(v50)%64))
						v75 = int32(8)
						v76 = v35 + v75
						if v36 != v21&int32(1073741816)-int32(16) {
							v35 = v76
							v36 = v36 + v75
							v42 = v74
							continue
						} else {
							break
						}
						break
					}
					v80 = v76
					v87 = v74
				}
				if v25 == int32(0) {
					v115 = v87
				} else {
					v91 = v80
					v92 = int32(0)
					v98 = v87
					for {
						v99 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
						v102 = v99 | v98<<(uint(int64(8))%64)
						v103 = int32(1)
						v106 = v92 + v103
						if v106 != v25 {
							v91 = v91 + v103
							v92 = v106
							v98 = v102
							continue
						} else {
							break
						}
						break
					}
					v115 = v102
				}
			}
			v123 = F_Int64GetDatum(m, int64(base.Ui64(v115)>>(uint(base.I64_extend_i32_u(v21<<(uint(int32(3))%32)-v14+int32(-64)))%64)))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				return v123
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(401807), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492981), int32(1676), int32(553789))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
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
func F_bittypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(103477))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_booland_statefunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		return base.B2i32(v7 != int32(0))
	}
}
func F_boolin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = v13
	goto L1
L1:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v22-int32(9)))&base.B2i32(v22 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v34 = F_strlen(m, v15)
	mBase = m.M
	if v34 == int32(0) {
		v61 = v2
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v15 = v15 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v66 = v11 + int32(15)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	switch v68 - int32(48) {
	case 0:
		goto L22
	case 1:
		goto L23
	default:
		goto L20
	case 22, 54:
		goto L27
	case 30, 62:
		goto L25
	case 31, 63:
		goto L24
	case 36, 68:
		goto L28
	case 41, 73:
		goto L26
	}
L7:
	;
	v41 = v34
	goto L8
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+(v15-int32(1))))))
	if base.Ui32(v48-int32(9)) < base.Ui32(int32(5)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v2
	goto L6
L10:
	;
	v56 = v41 - int32(1)
	if v56 != 0 {
		v41 = v56
		goto L8
	} else {
		goto L13
	}
L11:
	;
	if v48 == int32(32) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v61 = v41
	goto L6
L13:
	;
	goto L9
L14:
	;
	m.G0 = v11 + int32(16)
	return v153
L15:
	;
	if v129 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	v129 = v123
	goto L15
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v119)
	v123 = v121
	goto L16
L18:
	;
	if v66 == int32(0) {
		v123 = v114
		goto L16
	} else {
		goto L45
	}
L19:
	;
	v114 = int32(1)
	goto L18
L20:
	;
	v110 = int32(0)
	if v66 != 0 {
		v119 = v110
		v121 = v110
		goto L17
	} else {
		goto L44
	}
L21:
	;
	v119 = int32(0)
	v121 = v105
	goto L17
L22:
	;
	v100 = int32(1)
	if v61 != v100 {
		goto L20
	} else {
		goto L42
	}
L23:
	;
	v97 = int32(1)
	if v61 != v97 {
		goto L20
	} else {
		goto L41
	}
L24:
	;
	v86 = int32(2)
	if base.Ui32(v61) <= base.Ui32(v86) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v82 = F_pg_strncasecmp(m, v15, int32(240853), v61)
	mBase = m.M
	if v82 != 0 {
		goto L20
	} else {
		goto L33
	}
L26:
	;
	v78 = F_pg_strncasecmp(m, v15, int32(157159), v61)
	mBase = m.M
	if v78 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L27:
	;
	v74 = F_pg_strncasecmp(m, v15, int32(361107), v61)
	mBase = m.M
	if v74 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v72 = F_pg_strncasecmp(m, v15, int32(344091), v61)
	mBase = m.M
	if v72 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	if v66 != 0 {
		v105 = int32(1)
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v129 = int32(1)
	goto L15
L32:
	;
	goto L20
L33:
	;
	if v66 != 0 {
		v105 = int32(1)
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v129 = int32(1)
	goto L15
L35:
	;
	v89 = v86
	goto L37
L36:
	;
	v89 = v61
	goto L37
L37:
	;
	v90 = F_pg_strncasecmp(m, v15, int32(272934), v89)
	mBase = m.M
	if v90 == int32(0) {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v94 = F_pg_strncasecmp(m, v15, int32(338762), v89)
	mBase = m.M
	if v94 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v66 != 0 {
		v105 = int32(1)
		goto L21
	} else {
		goto L40
	}
L40:
	;
	v129 = int32(1)
	goto L15
L41:
	;
	v114 = v97
	goto L18
L42:
	;
	if v66 != 0 {
		v105 = v100
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v129 = int32(1)
	goto L15
L44:
	;
	v123 = v110
	goto L16
L45:
	;
	v119 = v114
	v121 = int32(1)
	goto L17
L46:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	v153 = v130
	goto L14
L47:
	;
	goto L48
L48:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = F_errsave_start(m, v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	if v133 == int32(0) {
		v153 = v131
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(283974)
	F_errmsg(m, int32(724727), v11)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, v132, int32(497063), int32(151), int32(276169))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v153 = v131
	goto L14
}
func F_boolle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 == v3) | base.B2i32(v5 != v3)
}
func F_boolne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 != v3) ^ base.B2i32(v5 != v3)
}
func F_boolop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 != 0 {
				v20 = F_array_contains_nulls(m, v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(152490), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497058), int32(424), int32(234708))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v25 = F_ArrayGetNItems(m, v22, v12+int32(16))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v29 != 0 {
								v37 = v29
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v37+v12, v25, v9+int32(7))
							mBase = m.M
							v42 = F__int_unique(m, v12)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								if v44 != 0 {
									v52 = v44
								} else {
									v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v53 = v42 + v52
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
								v57 = F_ArrayGetNItems(m, v45, v42+int32(16))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(7097))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v42)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v76 != v17 {
												F_pfree(m, v17)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(16)
													return v72
												}
											} else {
												m.G0 = v9 + int32(16)
												return v72
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v25 = F_ArrayGetNItems(m, v22, v12+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v29 != 0 {
						v37 = v29
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v37+v12, v25, v9+int32(7))
					mBase = m.M
					v42 = F__int_unique(m, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						if v44 != 0 {
							v52 = v44
						} else {
							v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						v53 = v42 + v52
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
						v57 = F_ArrayGetNItems(m, v45, v42+int32(16))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(7097))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v42)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v76 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v72
										}
									} else {
										m.G0 = v9 + int32(16)
										return v72
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
func F_booltext(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 != 0 {
		v5 = int32(344091)
	} else {
		v5 = int32(361107)
	}
	v6 = F_cstring_to_text(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_boot_yylex_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
		return int32(1)
	} else {
		v11 = F_palloc(m, int32(96))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
			if v11 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(48)
				return int32(1)
			} else {
				v26 = F__emscripten_memset_bulkmem(m, v11, base.I32_extend8_s(int32(0)), int32(96))
				mBase = m.M
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v28
				v30 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v27)+52)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(v27)+36)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v27)+4)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v28
				return v28
			}
		}
	}
}
func F_bpchargt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v54 = v17
	goto L17
L16:
	;
	v54 = v12 + int32(4)
	goto L17
L17:
	;
	v59 = v51
	goto L18
L18:
	;
	if v59 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(1)
	v79 = v19 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v82 = v80 & v78
	if v80 == v78 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v77 = v51 >> (uint(int32(31)) % 32) & v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = v59 - int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v71))))
	if v73 == int32(32) {
		v59 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v77 = v59
	goto L20
L25:
	;
	if v82 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v110 = v99
	goto L25
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v113 = v79
	goto L38
L37:
	;
	v113 = v19 + int32(4)
	goto L38
L38:
	;
	v118 = v110
	goto L39
L39:
	;
	if v118 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v137 = int32(1)
	if v21&v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v110 >> (uint(int32(31)) % 32) & v110
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = v118 - int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v130))))
	if v132 == int32(32) {
		v118 = v130
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v118
	goto L41
L46:
	;
	v141 = v137
	goto L48
L47:
	;
	v141 = int32(4)
	goto L48
L48:
	;
	v143 = int32(1)
	if v80&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_varstr_cmp(m, v12+v141, v77, v19+v147, v136, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.B2i32(int32(0) < v150)
L60:
	;
	goto L59
}
func F_bpcharlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(1)
	v12 = v7 + v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = v13 & v11
	if v13 == v11 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v15 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v18 = int32(4)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v20&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v33 = int32(1)
	if v15 != 0 {
		v43 = int32(base.Ui32(v13)>>(uint(v33)%32)) - v33
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v29 = v18
	goto L9
L8:
	;
	v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
	goto L9
L9:
	;
	if v20 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = v18
	goto L12
L11:
	;
	v32 = v29
	goto L12
L12:
	;
	v43 = v32
	goto L3
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v46 = v12
	goto L16
L15:
	;
	v46 = v7 + int32(4)
	goto L16
L16:
	;
	v50 = v43
	goto L17
L17:
	;
	if v50 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67*int32(28))+uint32(_consts[979])))
	goto L24
L19:
	;
	goto L18
L20:
	;
	v64 = v43 >> (uint(int32(31)) % 32) & v43
	goto L19
L21:
	;
	goto L22
L22:
	;
	v58 = v50 - int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v58))))
	if v60 == int32(32) {
		v50 = v58
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v64 = v50
	goto L19
L24:
	;
	if v72 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v77&v75 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v84 = v64
	goto L27
L27:
	;
	return v84
L28:
	;
	v80 = v75
	goto L30
L29:
	;
	v80 = int32(4)
	goto L30
L30:
	;
	v82 = F_pg_mbstrlen_with_len(m, v7+v80, v64)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v84 = v82
	goto L27
}
func F_bqarr_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != 0 {
			v15 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + v14<<(uint(int32(3))%32)
			v22 = F_palloc(m, v15)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v22
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v26)
				F_infix_3(m, v7, int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v31 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							m.G0 = v7 + int32(16)
							return v35
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						m.G0 = v7 + int32(16)
						return v35
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(16023), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497058), int32(650), int32(66863))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
func F_brinbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(12))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_brinRevmapInitialize(m, l0, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
				v14 = F_brin_build_desc(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
					return v4
				}
			}
		}
	}
}
func F_brinendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	F_brinRevmapTerminate(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		F_MemoryContextDelete(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, v2)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_brinhandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = v7
		v9 = int32(5)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(4222124650660278)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(3)
		v21 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = v9
		v27 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+76)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(6)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = int32(10)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(11)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(13)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v7
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+25)) = int32(16777217)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+17)) = int64(4311744769)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+13)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v27
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v27
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v27
		return v3
	}
}
func F_brinoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(1024), int32(12), int32(756944), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_btbpchar_pattern_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v3 = int32(4515120)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v7
	F_varstr_sortsupport(m, v6, int32(1042), int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v4
		return int32(0)
	}
}
func F_btbuildempty(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = F_smgr_bulk_start_rel(m, l0, int32(3))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = F_smgr_bulk_get_buf(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = int32(0)
				F_PageInit(m, v10, int32(8192), int32(16))
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = uint8(v5)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(-4616189618054758400)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v12
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(17180209506)
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
				v32 = int32(8)
				*(*uint16)(unsafe.Add(mBase, uint32(v10+v30)+12)) = uint16(v32)
				v34 = int32(72)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v34)
				F_smgr_bulk_write(m, v8, int32(0), v10, int32(1))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_smgr_bulk_finish(m, v8)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_btbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	v5 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v33 = int32(-1)
	v34 = v5
	v35 = v5
	v36 = v5
	v37 = v5
	v38 = v5
	v39 = v5
	v40 = v5
	v41 = v5
	v42 = v25
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
	if v33 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L2
L5:
	;
	v432 = int32(m.ExcTag)
	v433 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v432 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L6:
	;
	if v93 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v86 = v37
	v87 = v34
	v88 = v35
	v89 = v36
	v90 = v38
	v91 = v39
	v92 = v40
	v93 = v41
	v94 = v42
	goto L6
L8:
	;
	goto L9
L9:
	;
	v53 = v42 - int32(160)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v66 = v40
	v67 = l1
	goto L12
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v53
	v64 = F_palloc0(m, int32(40))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		v424 = v53
		goto L5
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v53
	F_before_shmem_exit(m, int32(237), v55)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		v424 = v53
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v66 = v64
	v67 = v64
	goto L12
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v82 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v25
	goto L18
L16:
	;
	v86 = v55
	v87 = v53
	v88 = v82
	v89 = v80
	v90 = v67
	v91 = v55
	v92 = v66
	v93 = int32(0)
	v94 = v53
	goto L6
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	v106 = int32(0)
	v107 = m.G0
	v109 = v107 - int32(16)
	m.G0 = v109
	v112 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v116 = F_LWLockAcquire(m, v112+int32(2560), v106)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v89
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	F_cancel_before_shmem_exit(m, int32(237), v86)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L61
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	F_btvacuumscan(m, l0, v90, l2, l3, v126&int32(65535))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L49
	}
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	v120 = int32(1)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119))))
	if base.Ui32(int32(65406)) < base.Ui32(v121) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v126 = v120
	goto L26
L25:
	;
	v126 = v121 + v120
	goto L26
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v119))) = uint16(v126)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if int32(0) < v128 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v235+int32(2560))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L45
	}
L28:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v213+int32(2560))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L41
	}
L29:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+60))
	v147 = v106
	goto L32
L30:
	;
	goto L31
L31:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	if v189 <= v128 {
		goto L27
	} else {
		goto L39
	}
L32:
	;
	v158 = v119 + int32(12) + v147*int32(12)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v133 == v159 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v91)+64))
	if v161 == v162 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v165 = v147 + int32(1)
	if v165 != v128 {
		v147 = v165
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v91)+60))
	v194 = v119 + v128*int32(12)
	*(*uint16)(unsafe.Add(mBase, uint32(v194)+20)) = uint16(v126)
	*(*int64)(unsafe.Add(mBase, uint32(v194)+12)) = v191
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v197 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v202+int32(2560))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L40
	}
L40:
	;
	m.G0 = v109 + int32(16)
	goto L22
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v222 + int32(4)
	F_errmsg_internal(m, int32(692736), v109)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(493612), int32(3578), int32(285799))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errmsg_internal(m, int32(118743), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(493612), int32(3586), int32(285799))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	F_cancel_before_shmem_exit(m, int32(237), v86)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v89
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	v283 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v289 = F_LWLockAcquire(m, v285+int32(2560), v283)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v293 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v368+int32(2560))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L60
	}
L53:
	;
	v297 = v292 + int32(12)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v91)+60))
	v304 = v283
	goto L54
L54:
	;
	v323 = v297 + v304*int32(12)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v324 != v298 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L52
L56:
	;
	v343 = v304 + int32(1)
	if v343 != v293 {
		v304 = v343
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v91)+64))
	if v326 != v327 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v329 = int32(12)
	v333 = v293*v329 + v297 - v329
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v333)))
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v338 - int32(1)
	goto L52
L59:
	;
	goto L55
L60:
	;
	m.G0 = v25 + int32(32)
	return v90
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	F__bt_end_vacuum_callback(m, int32(0), v86)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v87
	F_pg_re_throw(m)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		v424 = v94
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L4
L64:
	;
	v437 = int32(v433)
	m.G0 = v424
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v25 == v442 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	m.ExcPending = 1
	goto L73
L66:
	;
	if v445 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v445 = v444
	goto L69
L68:
	;
	v445 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v33 = v445
	v34 = v446
	v35 = v451
	v36 = v452
	v37 = v450
	v38 = v449
	v39 = v447
	v40 = v448
	v41 = v439
	v42 = v424
	goto L1
L71:
	;
	goto L72
L72:
	;
	F___wasm_longjmp(m, v440, v439)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	return int32(0)
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = int64(0)
	goto L1
L1:
	;
	v21 = F__bt_first(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v62
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	return int64(0)
L5:
	;
	if v21 == int32(0) {
		v62 = v19
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v30 = l0 + int32(60)
	v33 = v19
	goto L7
L7:
	;
	F_tbm_add_tuples(m, l1, v30, int32(1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v40 = v38 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v40
	v43 = v33 + int64(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	if v44 < v40 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = F__bt_next(m, l0, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v52 = v40
	goto L12
L12:
	;
	v30 = v10 + int32(104) + v52*int32(10)
	v33 = v43
	goto L7
L13:
	;
	if v47 == int32(0) {
		v62 = v43
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v52 = v51
	goto L12
L15:
	;
	v64 = F__bt_start_prim_scan(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L2
L18:
	;
	if v64 != 0 {
		v19 = v62
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func F_btgettreeheight(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F__bt_getrootheight(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_btint82cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v6 < v5) - base.B2i32(v5 < v6)
}
func F_btint84cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v6 < v5) - base.B2i32(v5 < v6)
}
func F_btint8cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
}
func F_btint8skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(201)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(202)
	v8 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v2))) = v8
		v14 = F_Int64GetDatum(m, int64(9223372036854775807))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v2)+4)) = v14
			return int32(0)
		}
	}
}
func F_btmarkpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[48])))
	if v4 != 0 {
		F_ReleaseBuffer(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[48]))) = int32(0)
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
			if v9 != int32(-1) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[48]))) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
		if v9 != int32(-1) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[48]))) = int64(-4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
			return
		}
	}
}
func F_btnamecmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	goto L3
L3:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	return v45 - v46
L6:
	;
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L12
L9:
	;
	v41 = v4
	v45 = int32(0)
	goto L10
L10:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L4
L11:
	;
	v41 = v36
	v45 = v38
	goto L10
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v36 = v30
	v38 = int32(0)
	goto L11
L14:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	return int32(0)
L19:
	;
	return v57
}
func F_btparallelrescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	v5 = v3 + v4
	v7 = v5 + int32(12)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(-1)
		F_LWLockRelease(m, v7)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_btvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 float64
	_ = v220
	var v221 float64
	_ = v221
	var v226 int32
	_ = v226
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v10 != 0 {
		v226 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v226
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_ReadBuffer(m, v14, v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v83 = l1
	goto L5
L5:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	v90 = v88 - v89
	v91 = m.G0
	v93 = v91 - int32(32)
	m.G0 = v93
	v96 = F_ReadBuffer(m, v87, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L31
	}
L6:
	;
	if v70 == int32(0) {
		v226 = v13
		goto L1
	} else {
		goto L28
	}
L7:
	;
	return int32(0)
L8:
	;
	F_LockBuffer(m, v16, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F__bt_checkpage(m, v14, v16)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v16 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	if base.Ui32(v43) <= base.Ui32(int32(2)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v16^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L11
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v42 = v36 + v16<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	F_LockBuffer(m, v16, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	F_LockBuffer(m, v16, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	F_ReleaseBuffer(m, v16)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v70 = int32(1)
	goto L6
L20:
	;
	F_ReleaseBuffer(m, v16)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v52 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v70 = v67
	goto L6
L23:
	;
	v60 = F_RelationGetNumberOfBlocksInFork(m, v14, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v67 = int32(0)
	goto L22
L26:
	;
	v63 = base.I32_div_u_s(v60, int32(20))
	if base.Ui32(v63) < base.Ui32(v52) {
		v67 = int32(1)
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v74 = F_palloc0(m, int32(40))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v76 = int32(0)
	F_btvacuumscan(m, l0, v74, v76, v76, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)) = uint8(v81)
	v83 = v74
	goto L5
L31:
	;
	F_LockBuffer(m, v96, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F__bt_checkpage(m, v87, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	if v96 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_ReleaseBuffer(m, v96)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L61
	}
L35:
	;
	F_LockBuffer(m, v96, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L43
	}
L36:
	;
	v122 = v120 + int32(28)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if base.Ui32(v123) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L40
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106+(v96^int32(-1))<<(uint(int32(2))%32))))
	v120 = v112
	goto L36
L38:
	;
	goto L39
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v120 = v114 + v96<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+48))
	if v126 != v90 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_LockBuffer(m, v96, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	F_LockBuffer(m, v96, int32(2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v137 = int32(4509780)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v139 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if base.Ui32(v143) <= base.Ui32(int32(2)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v146 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+12)) = uint16(v146)
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120-int32(-64)))) = uint8(v150)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = int32(3)
	goto L47
L46:
	;
	goto L47
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+48)) = v90
	F_MarkBufferDirty(m, v96)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v87)+48))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+118)))
	if v160 != int32(112) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v204 = int32(4509780)
	v206 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v206 - int32(1)
	F_LockBuffer(m, v96, int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L60
	}
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v164 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	if v167 != 0 {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L56
	}
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v87)+40))
	if v168 != 0 {
		goto L49
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_XLogRegisterBuffer(m, int32(0), v96, int32(14))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v120)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v120)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v183
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120-int32(-64)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+28)) = uint8(v188)
	F_XLogRegisterBufData(m, int32(0), v93+int32(4), int32(28))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v198 = F_XLogInsert(m, int32(11), int32(224))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_rotr(v198, int64(32))
	goto L49
L60:
	;
	goto L34
L61:
	;
	m.G0 = v93 + int32(32)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v219 != 0 {
		v226 = v83
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v83)+8))
	if base.F64_lt(v220, v221) == int32(0) {
		v226 = v83
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v83)+8)) = v220
	v226 = v83
	goto L1
}
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v10 = F_build_attrmap_by_name(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 != v15 {
		v76 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v76
L4:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v17 < v18 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L6:
	;
	v21 = int32(20)
	v27 = v17
	goto L9
L7:
	;
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v63 = v58
	goto L5
L9:
	;
	v35 = v27 << (uint(int32(4)) % 32)
	v36 = l0 + v21 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != 0 {
		v76 = v10
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v63 = v40
	goto L5
L11:
	;
	v38 = int32(1)
	v39 = v27 + v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40+v27<<(uint(v38)%32)))))
	if v39 != v44 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v44 != 0 {
		v76 = v10
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v18 != v39 {
		v27 = v39
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+9)))
	if v46 != int32(1) {
		v76 = v10
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	v50 = l1 + v21 + v35
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	if v49 != v51 {
		v76 = v10
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+12)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+12)))
	if v53 != v54 {
		v76 = v10
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	goto L10
L20:
	;
	F_pfree(m, v10)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(0)
	goto L3
}
func F_build_colinfo_names_hash(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 < int32(32) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(274877907008)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_hash_create(m, int32(163215), v20+v11, v9, int32(1048))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v26 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = int32(0)
	v32 = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v62 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v31<<(uint(int32(2))%32))))
	if v40 == int32(0) {
		v51 = v32
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v54 = v31 + int32(1)
	if v54 < v51 {
		v31 = v54
		v32 = v51
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v43 == int32(0) {
		v51 = v32
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v48 = F_hash_search(m, v43, v40, int32(1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = v50
	goto L10
L14:
	;
	goto L9
L15:
	;
	v67 = int32(0)
	v68 = v62
	goto L18
L16:
	;
	goto L17
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v98 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v67<<(uint(int32(2))%32))))
	if v76 == int32(0) {
		v87 = v68
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v90 = v67 + int32(1)
	if v90 < v87 {
		v67 = v90
		v68 = v87
		goto L18
	} else {
		goto L24
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v79 == int32(0) {
		v87 = v68
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v84 = F_hash_search(m, v79, v76, int32(1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = v86
	goto L20
L24:
	;
	goto L19
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v101 <= int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v106 = int32(0)
	goto L27
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L1
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v106<<(uint(int32(2))%32))))
	v119 = F_hash_search(m, v111, v116, int32(1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v122 = v106 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v122 < v123 {
		v106 = v122
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
}
func F_build_joinrel_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
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
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int64
	_ = v625
	var v640 int64
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v659 int64
	_ = v659
	var v661 int64
	_ = v661
	var v664 int64
	_ = v664
	var v666 int32
	_ = v666
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+32)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == v7 {
		v659 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v661 = int64(1073741823)
	if v661 <= v659 {
		goto L158
	} else {
		goto L159
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		v659 = v21
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = v7
	v43 = v21
	goto L4
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(6) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v659 = v640
	goto L1
L6:
	;
	v643 = v41 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v643 < v644 {
		v41 = v643
		v43 = v640
		goto L4
	} else {
		goto L156
	}
L7:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	v621 = F_lappend(m, v620, v606)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L15
	} else {
		goto L155
	}
L8:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v599 = F_bms_intersect(m, v598, v29)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L15
	} else {
		goto L153
	}
L9:
	;
	if v50 == int32(319) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v355 == int32(-4) {
		goto L102
	} else {
		goto L103
	}
L12:
	;
	v55 = F_find_placeholder_info(m, l0, v49)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L15
	} else {
		goto L95
	}
L15:
	;
	return
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v57 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v112 == int32(0) {
		v640 = v43
		goto L6
	} else {
		goto L31
	}
L18:
	;
	v112 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v65 = int32(1)
	if v29 == int32(0) {
		v103 = v65
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v112 = v103
	goto L17
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v69 < v68 {
		v103 = v65
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v71 = int32(1)
	if v68 <= v71 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = v71
	goto L26
L25:
	;
	v74 = v68
	goto L26
L26:
	;
	v75 = int32(8)
	v80 = int32(0)
	goto L27
L27:
	;
	v87 = v80 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57+v75+v87)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v29+v75))))
	v94 = v89 & (v91 ^ int32(-1))
	v96 = base.B2i32(v94 != int32(0))
	if v94 != 0 {
		v103 = v96
		goto L21
	} else {
		goto L29
	}
L28:
	;
	v103 = v96
	goto L21
L29:
	;
	v98 = v80 + int32(1)
	if v98 != v74 {
		v80 = v98
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if l5 == int32(0) {
		v606 = v49
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v117 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v119 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if l4 == int32(0) {
		goto L8
	} else {
		goto L72
	}
L35:
	;
	v122 = F_bms_is_member(m, v119, v29)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	if v122 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v128 = int32(0)
	if v126 == v128 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v181 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v181 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v127 == int32(0) {
		v172 = v128
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v181 = v172
	goto L38
L43:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v138 < v137 {
		v172 = v128
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v140 = int32(1)
	if v137 <= v140 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v143 = v140
	goto L47
L46:
	;
	v143 = v137
	goto L47
L47:
	;
	v144 = int32(8)
	v149 = int32(0)
	goto L48
L48:
	;
	v156 = v149 << (uint(int32(2)) % 32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v126+v144+v156)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+(v127+v144))))
	v163 = v158 & (v160 ^ int32(-1))
	v165 = base.B2i32(v163 == int32(0))
	if v163 != 0 {
		v172 = v165
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v172 = v165
	goto L42
L50:
	;
	v167 = v149 + int32(1)
	if v167 != v143 {
		v149 = v167
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v184 != int32(2) {
		goto L34
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v247 = F_bms_add_member(m, v245, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L71
	}
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v189 = int32(0)
	if v187 == v189 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v242 == int32(0) {
		goto L34
	} else {
		goto L70
	}
L57:
	;
	v242 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v188 == int32(0) {
		v233 = v189
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v242 = v233
	goto L56
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v199 < v198 {
		v233 = v189
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v201 = int32(1)
	if v198 <= v201 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v204 = v201
	goto L65
L64:
	;
	v204 = v198
	goto L65
L65:
	;
	v205 = int32(8)
	v210 = int32(0)
	goto L66
L66:
	;
	v217 = v210 << (uint(int32(2)) % 32)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v187+v205+v217)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217+(v188+v205))))
	v224 = v219 & (v221 ^ int32(-1))
	v226 = base.B2i32(v224 == int32(0))
	if v224 != 0 {
		v233 = v226
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v233 = v226
	goto L60
L68:
	;
	v228 = v210 + int32(1)
	if v228 != v204 {
		v210 = v228
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L54
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v247
	goto L34
L72:
	;
	v252 = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v253 <= v252 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v262 = v252
	goto L74
L74:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272+v262<<(uint(int32(2))%32))))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+16))
	v278 = int32(0)
	if v271 == v278 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L8
L76:
	;
	if v331 != 0 {
		goto L90
	} else {
		goto L91
	}
L77:
	;
	v331 = int32(1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v277 == int32(0) {
		v322 = v278
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v331 = v322
	goto L76
L81:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v288 < v287 {
		v322 = v278
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v290 = int32(1)
	if v287 <= v290 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v293 = v290
	goto L85
L84:
	;
	v293 = v287
	goto L85
L85:
	;
	v294 = int32(8)
	v299 = int32(0)
	goto L86
L86:
	;
	v306 = v299 << (uint(int32(2)) % 32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v271+v294+v306)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v277+v294))))
	v313 = v308 & (v310 ^ int32(-1))
	v315 = base.B2i32(v313 == int32(0))
	if v313 != 0 {
		v322 = v315
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v322 = v315
	goto L80
L88:
	;
	v317 = v299 + int32(1)
	if v317 != v293 {
		v299 = v317
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v334 = F_bms_add_member(m, v332, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L15
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v338 = v262 + int32(1)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v338 < v339 {
		v262 = v338
		goto L74
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v334
	goto L92
L94:
	;
	goto L75
L95:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v345
	F_errmsg_internal(m, int32(480492), v18)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L15
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(499150), int32(1172), int32(73417))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L15
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v578 = F_lappend(m, v577, v562)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L15
	} else {
		goto L152
	}
L99:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v454)+24))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v555 = F_bms_intersect(m, v554, v29)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L15
	} else {
		goto L150
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L15
	} else {
		goto L147
	}
L101:
	;
	v448 = int64(*(*int32)(unsafe.Add(mBase, uint32(v447))))
	if l5 == int32(0) {
		v562 = v49
		goto L98
	} else {
		goto L122
	}
L102:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v359+v360<<(uint(int32(2))%32)-int32(4))))
	v447 = v366 + int32(8)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v369) <= base.Ui32(v355) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v371+v355<<(uint(int32(2))%32))))
	if v375 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v378 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v375)+80)))
	v382 = (v378 - v379) << (uint(int32(2)) % 32)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)+84))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382+v383)))
	if v385 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v440 == int32(0) {
		v640 = v43
		goto L6
	} else {
		goto L121
	}
L108:
	;
	v440 = int32(0)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v393 = int32(1)
	if v29 == int32(0) {
		v431 = v393
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v440 = v431
	goto L107
L112:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v397 < v396 {
		v431 = v393
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v399 = int32(1)
	if v396 <= v399 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v402 = v399
	goto L116
L115:
	;
	v402 = v396
	goto L116
L116:
	;
	v403 = int32(8)
	v408 = int32(0)
	goto L117
L117:
	;
	v415 = v408 << (uint(int32(2)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v385+v403+v415)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+(v29+v403))))
	v422 = v417 & (v419 ^ int32(-1))
	v424 = base.B2i32(v422 != int32(0))
	if v422 != 0 {
		v431 = v424
		goto L111
	} else {
		goto L119
	}
L118:
	;
	v431 = v424
	goto L111
L119:
	;
	v426 = v408 + int32(1)
	if v426 != v402 {
		v408 = v426
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v375)+88))
	v447 = v443 + v382
	goto L101
L122:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v451 == int32(-4) {
		v562 = v49
		goto L98
	} else {
		goto L123
	}
L123:
	;
	v454 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v456 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if l4 == int32(0) {
		goto L99
	} else {
		goto L137
	}
L126:
	;
	v459 = F_bms_is_member(m, v456, v29)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	if v459 == int32(0) {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v465 = F_bms_is_member(m, v463, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L15
	} else {
		goto L129
	}
L129:
	;
	if v465 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v469 != int32(2) {
		goto L125
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v454)+24))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v480 = F_bms_add_member(m, v478, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L15
	} else {
		goto L136
	}
L133:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v474 = F_bms_is_member(m, v472, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L15
	} else {
		goto L134
	}
L134:
	;
	if v474 == int32(0) {
		goto L125
	} else {
		goto L135
	}
L135:
	;
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+24)) = v480
	goto L125
L137:
	;
	v485 = int32(0)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v486 <= v485 {
		goto L99
	} else {
		goto L138
	}
L138:
	;
	v495 = v485
	goto L139
L139:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v505+v495<<(uint(int32(2))%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
	v511 = F_bms_is_member(m, v504, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L15
	} else {
		goto L141
	}
L140:
	;
	goto L99
L141:
	;
	if v511 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v454)+24))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v509)+24))
	v515 = F_bms_add_member(m, v513, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L15
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v519 = v495 + int32(1)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v519 < v520 {
		v495 = v519
		goto L139
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+24)) = v515
	goto L144
L146:
	;
	goto L140
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v355
	F_errmsg_internal(m, int32(478509), v18+int32(16))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L15
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(499150), int32(426), int32(307472))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L15
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	v557 = F_bms_join(m, v553, v555)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L15
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+24)) = v557
	v562 = v454
	goto L98
L152:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+4)) = v578
	v640 = v43 + v448
	goto L6
L153:
	;
	v601 = F_bms_join(m, v597, v599)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v601
	v606 = v117
	goto L7
L155:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v621
	v625 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+24)))
	v640 = v43 + v625
	goto L6
L156:
	;
	goto L5
L157:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v666)+32)) = base.I32_wrap_i64(v664)
	m.G0 = v18 + int32(32)
	return
L158:
	;
	v664 = v661
	goto L160
L159:
	;
	v664 = v659
	goto L160
L160:
	;
	goto L157
}
func F_build_paths_for_OR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(144)
	return v196
L2:
	;
	v28 = v5
	v31 = v5
	v32 = v5
	goto L7
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if int32(0) < v18 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v196 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v31<<(uint(int32(2))%32))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+110)))
	if v38 != int32(1) {
		v180 = v28
		v184 = v32
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v196 = v180
	goto L1
L9:
	;
	v186 = v31 + int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v186 < v187 {
		v28 = v180
		v31 = v186
		v32 = v184
		goto L7
	} else {
		goto L40
	}
L10:
	;
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+88))
	if v43 == v41 {
		v66 = v32
		v67 = v41
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v73 = F__emscripten_memset_bulkmem(m, v15+int32(12), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L22
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+100)))
	if v47 != 0 {
		v66 = v32
		v67 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v32
	v54 = v43
	goto L16
L15:
	;
	v48 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v56 = F_predicate_implied_by(m, v54, v53, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+88))
	v53 = v48
	v54 = v52
	goto L16
L19:
	;
	if v56 == int32(0) {
		v180 = v28
		v184 = v53
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+88))
	v62 = F_predicate_implied_by(m, v60, l3, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v66 = v53
	v67 = v62 ^ int32(1)
	goto L11
L22:
	;
	if l2 == int32(0) {
		v109 = v41
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if (v109|v67)&int32(1) == int32(0) {
		v180 = v28
		v184 = v66
		goto L9
	} else {
		goto L30
	}
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v76 <= int32(0) {
		v109 = v41
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v83 = v41
	goto L26
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v83<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v95, v37, v15+int32(12))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)))
	v109 = v104
	goto L23
L28:
	;
	v101 = v83 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v101 < v102 {
		v83 = v101
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if l3 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v169 = F_build_index_paths(m, l0, l1, v37, v15+int32(12), v67, int32(1), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L17
	} else {
		goto L38
	}
L32:
	;
	v124 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v125 <= v124 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v132 = v124
	goto L34
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v132<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v144, v37, v15+int32(12))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	goto L31
L36:
	;
	v150 = v132 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v150 < v151 {
		v132 = v150
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v171 = F_list_concat(m, v28, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v180 = v171
	v184 = v66
	goto L9
L40:
	;
	goto L8
}
func F_build_sorted_items(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	v6 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = v17 * l3
	v24 = F_palloc0(m, v18*int32(5)+v17*int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v28 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v24 + v28*int32(12)
	v43 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = F_palloc(m, v84<<(uint(int32(2))%32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v56 = v24 + v43*int32(12)
	v57 = l3 * v43
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v33 + v18<<(uint(int32(2))%32) + v57
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v33 + v57<<(uint(int32(2))%32)
	v65 = v43 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v65 < v66 {
		v43 = v65
		goto L6
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	goto L7
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v89 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v139 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v110 = v98 << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112+v110)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v116 = F_get_typlen(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87+v110))) = v116
	v120 = v98 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 < v121 {
		v98 = v120
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_qsort_interruptible(m, v24, v296, int32(12), int32(1062), l2)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L46
	}
L18:
	;
	F_pfree(m, v24)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L45
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v155 = v6
	v157 = v6
	goto L22
L22:
	;
	if base.B2i32(l3 <= int32(0)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v296
	if v296 != 0 {
		goto L17
	} else {
		goto L44
	}
L24:
	;
	v304 = v157 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v304 < v305 {
		v155 = v296
		v157 = v304
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v166 = v24 + v155*int32(12)
	v175 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v296 = v155 + int32(1)
	goto L24
L28:
	;
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v185 <= v184 {
		v222 = v184
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v233 = int32(2)
	v234 = v222 << (uint(v233) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234+v235)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v157<<(uint(v233)%32))))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242+v234)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v157))))
	if v246 != 0 {
		v257 = v241
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4+v175<<(uint(int32(1))%32)))))
	v198 = v184
	goto L32
L32:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v198<<(uint(int32(1))%32)))))
	if v192 == v212 {
		v222 = v198
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v222 = v185
	goto L30
L34:
	;
	v215 = v198 + int32(1)
	if v215 != v185 {
		v198 = v215
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v258+v175<<(uint(int32(2))%32)))) = v257
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v263+v175))) = uint8(v246)
	v267 = v175 + int32(1)
	if v267 != l3 {
		v175 = v267
		goto L28
	} else {
		goto L42
	}
L37:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v234+v87)))
	if v248 != int32(-1) {
		v257 = v241
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v251 = F_toast_raw_datum_size(m, v241)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v251) {
		v296 = v155
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v255 = F_pg_detoast_datum(m, v241)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v257 = v255
	goto L36
L42:
	;
	goto L29
L43:
	;
	goto L23
L44:
	;
	goto L18
L45:
	;
	return int32(0)
L46:
	;
	return v24
}
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v7 = l1 << (uint(int32(2)) % 32)
	v9 = v7 + int32(24)
	v10 = F_palloc0(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
		} else {
			if l1 <= int32(0) {
			} else {
				if v7 != 0 {
					v20 = F__emscripten_memcpy_bulkmem(m, v10+int32(24), l0, v7)
					mBase = m.M
				} else {
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(26)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v9 << (uint(int32(2)) % 32)
		return v10
	}
}
func F_byteacmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v78 = int32(1)
	if v16&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v50 = int32(4)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = int32(1)
	if v47&v65 != 0 {
		v77 = int32(base.Ui32(v47)>>(uint(v65)%32)) - v65
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v61 = v50
	goto L21
L20:
	;
	v61 = base.B2i32(v52 == int32(18)) << (uint(v50) % 32)
	goto L21
L21:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = v50
	goto L24
L23:
	;
	v64 = v61
	goto L24
L24:
	;
	v77 = v64
	goto L15
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = v9 + v82
	v84 = int32(1)
	if v47&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v84
	goto L31
L30:
	;
	v88 = int32(4)
	goto L31
L31:
	;
	v89 = v14 + v88
	v90 = base.B2i32(v46 < v77)
	if v46 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v46
	goto L34
L33:
	;
	v91 = v77
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v153 = int32(0)
	goto L35
L37:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L47
L38:
	;
	if (v83|v89)&int32(3) != 0 {
		v122 = v83
		v123 = v89
		v124 = v91
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v115 = v83
	v116 = v89
	v117 = v91
	goto L40
L40:
	;
	if v117 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v99 = v83
	v100 = v89
	v101 = v91
	goto L42
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L40
L44:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L37
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v153 = v132 - v133
	goto L35
L49:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v158 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v153 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v164 = v153
	goto L63
L62:
	;
	v164 = base.B2i32(v77 < v46) - v90
	goto L63
L63:
	;
	return v164
}
func F_byteaeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_toast_raw_datum_size(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v108
L2:
	;
	return int32(0)
L3:
	;
	v12 = F_toast_raw_datum_size(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v8 != v12 {
		v108 = int32(0)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v15 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v17 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v19 = int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21&v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = v19
	goto L10
L9:
	;
	v24 = int32(4)
	goto L10
L10:
	;
	v25 = v15 + v24
	v26 = int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v26
	goto L13
L12:
	;
	v31 = int32(4)
	goto L13
L13:
	;
	v32 = v17 + v31
	v33 = int32(4)
	v34 = v8 - v33
	if base.Ui32(v33) <= base.Ui32(v34) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v97 != v15 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v96 = int32(0)
	goto L14
L16:
	;
	v70 = v65
	v71 = v66
	v72 = v67
	goto L26
L17:
	;
	if (v25|v32)&int32(3) != 0 {
		v65 = v25
		v66 = v32
		v67 = v34
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v58 = v25
	v59 = v32
	v60 = v34
	goto L19
L19:
	;
	if v60 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v42 = v25
	v43 = v32
	v44 = v34
	goto L21
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v47 != v48 {
		v65 = v42
		v66 = v43
		v67 = v44
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v58 = v53
	v59 = v51
	v60 = v55
	goto L19
L23:
	;
	v50 = int32(4)
	v51 = v43 + v50
	v53 = v42 + v50
	v55 = v44 - v50
	if base.Ui32(int32(3)) < base.Ui32(v55) {
		v42 = v53
		v43 = v51
		v44 = v55
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v65 = v58
	v66 = v59
	v67 = v60
	goto L16
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v75 == v76 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v96 = v75 - v76
	goto L14
L28:
	;
	v78 = int32(1)
	v83 = v72 - v78
	if v83 != 0 {
		v70 = v70 + v78
		v71 = v71 + v78
		v72 = v83
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	F_pfree(m, v15)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v102 = base.B2i32(v96 == int32(0))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 == v103 {
		v108 = v102
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v17)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v108 = v102
	goto L1
}
func F_bytealt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v78 = int32(1)
	if v16&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v50 = int32(4)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = int32(1)
	if v47&v65 != 0 {
		v77 = int32(base.Ui32(v47)>>(uint(v65)%32)) - v65
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v61 = v50
	goto L21
L20:
	;
	v61 = base.B2i32(v52 == int32(18)) << (uint(v50) % 32)
	goto L21
L21:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = v50
	goto L24
L23:
	;
	v64 = v61
	goto L24
L24:
	;
	v77 = v64
	goto L15
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = v9 + v82
	v84 = int32(1)
	if v47&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v84
	goto L31
L30:
	;
	v88 = int32(4)
	goto L31
L31:
	;
	v89 = v14 + v88
	v90 = base.B2i32(v46 < v77)
	if v46 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v46
	goto L34
L33:
	;
	v91 = v77
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v153 = int32(0)
	goto L35
L37:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L47
L38:
	;
	if (v83|v89)&int32(3) != 0 {
		v122 = v83
		v123 = v89
		v124 = v91
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v115 = v83
	v116 = v89
	v117 = v91
	goto L40
L40:
	;
	if v117 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v99 = v83
	v100 = v89
	v101 = v91
	goto L42
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L40
L44:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L37
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v153 = v132 - v133
	goto L35
L49:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v158 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v162 = int32(0)
	return base.B2i32(v153 == v162)&v90 | base.B2i32(v153 < v162)
L60:
	;
	goto L59
}
func F_byteaout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int64
	_ = v240
	var v244 int32
	_ = v244
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v259 int32
	_ = v259
	var v266 int64
	_ = v266
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[1058]))
		switch v21 {
		case 0:
			v186 = int32(4)
			v188 = v16 + v186
			v189 = int32(1)
			v190 = v16 + v189
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v193 = v191 & v189
			if v191 == v189 {
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
				if base.Ui32((v196-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v220 = v186
					if v193 != 0 {
						v221 = v190
					} else {
						v221 = v188
					}
					v222 = int32(1)
					if v220 == v222 {
						v276 = v221
						v285 = int64(1)
					} else {
						v231 = v221
						v232 = int32(0)
						v240 = int64(1)
						for {
							v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
							if base.Ui32((v244-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v251 = int64(4)
							} else {
								v251 = int64(1)
							}
							if v244 == int32(92) {
								v254 = int64(2)
							} else {
								v254 = v251
							}
							v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
							if base.Ui32((v259-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v266 = int64(4)
							} else {
								v266 = int64(1)
							}
							if v259 == int32(92) {
								v269 = int64(2)
							} else {
								v269 = v266
							}
							v270 = v240 + v254 + v269
							v271 = int32(2)
							v272 = v231 + v271
							v274 = v232 + v271
							if v274 != v220&int32(-2) {
								v231 = v272
								v232 = v274
								v240 = v270
								continue
							} else {
								break
							}
							break
						}
						v276 = v272
						v285 = v270
					}
					if v220&v222 != 0 {
						v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
						if base.Ui32((v289-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v296 = int64(4)
						} else {
							v296 = int64(1)
						}
						if v289 == int32(92) {
							v299 = int64(2)
						} else {
							v299 = v296
						}
						v302 = v285 + v299
					} else {
						v302 = v285
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v302) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v421 = m.ExcPending
						if v421 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v424 = m.ExcPending
							if v424 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(400095), int32(0))
								mBase = m.M
								v428 = m.ExcPending
								if v428 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499834), int32(438), int32(66423))
									mBase = m.M
									v433 = m.ExcPending
									if v433 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v314 = v302
						v316 = F_palloc(m, base.I32_wrap_i64(v314))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v319 = int32(1)
							v320 = v318 & v319
							if v318 == v319 {
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
								if base.Ui32((v324-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v348 = int32(4)
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v346 = base.B2i32(v324 == int32(18)) << (uint(int32(4)) % 32)
									if v346 != 0 {
										v348 = v346
										if v320 != 0 {
											v349 = v190
										} else {
											v349 = v188
										}
										v350 = v316
										v351 = v349
										v355 = v348
										for {
											v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
											if v360 == int32(92) {
												v363 = int32(23644)
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
												v397 = v350 + int32(2)
											} else {
												if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v373 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
													v375 = int32(7)
													v377 = int32(48)
													v378 = v360&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
													v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
													v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
													v397 = v350 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
													v397 = v350 + int32(1)
												}
											}
											v398 = int32(1)
											v401 = v355 - v398
											if v401 != 0 {
												v350 = v397
												v351 = v351 + v398
												v355 = v401
												continue
											} else {
												break
											}
											break
										}
										v402 = v397
										v406 = v316
									} else {
										v402 = v316
										v406 = v316
									}
								}
							} else {
								v335 = int32(1)
								if v320 != 0 {
									v346 = int32(base.Ui32(v318)>>(uint(v335)%32)) - v335
								} else {
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v346 = int32(base.Ui32(v339)>>(uint(int32(2))%32)) - int32(4)
								}
								if v346 != 0 {
									v348 = v346
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v402 = v316
									v406 = v316
								}
							}
							v412 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
							m.G0 = v13 + int32(16)
							return v406
						}
					}
				} else {
					v217 = base.B2i32(v196 == int32(18)) << (uint(int32(4)) % 32)
					if v217 != 0 {
						v220 = v217
						if v193 != 0 {
							v221 = v190
						} else {
							v221 = v188
						}
						v222 = int32(1)
						if v220 == v222 {
							v276 = v221
							v285 = int64(1)
						} else {
							v231 = v221
							v232 = int32(0)
							v240 = int64(1)
							for {
								v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
								if base.Ui32((v244-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v251 = int64(4)
								} else {
									v251 = int64(1)
								}
								if v244 == int32(92) {
									v254 = int64(2)
								} else {
									v254 = v251
								}
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
								if base.Ui32((v259-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v266 = int64(4)
								} else {
									v266 = int64(1)
								}
								if v259 == int32(92) {
									v269 = int64(2)
								} else {
									v269 = v266
								}
								v270 = v240 + v254 + v269
								v271 = int32(2)
								v272 = v231 + v271
								v274 = v232 + v271
								if v274 != v220&int32(-2) {
									v231 = v272
									v232 = v274
									v240 = v270
									continue
								} else {
									break
								}
								break
							}
							v276 = v272
							v285 = v270
						}
						if v220&v222 != 0 {
							v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
							if base.Ui32((v289-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v296 = int64(4)
							} else {
								v296 = int64(1)
							}
							if v289 == int32(92) {
								v299 = int64(2)
							} else {
								v299 = v296
							}
							v302 = v285 + v299
						} else {
							v302 = v285
						}
						if base.Ui64(int64(1073741824)) <= base.Ui64(v302) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v421 = m.ExcPending
							if v421 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v424 = m.ExcPending
								if v424 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(400095), int32(0))
									mBase = m.M
									v428 = m.ExcPending
									if v428 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499834), int32(438), int32(66423))
										mBase = m.M
										v433 = m.ExcPending
										if v433 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v314 = v302
							v316 = F_palloc(m, base.I32_wrap_i64(v314))
							mBase = m.M
							v317 = m.ExcPending
							if v317 != 0 {
								return int32(0)
							} else {
								v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
								v319 = int32(1)
								v320 = v318 & v319
								if v318 == v319 {
									v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
									if base.Ui32((v324-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v348 = int32(4)
										if v320 != 0 {
											v349 = v190
										} else {
											v349 = v188
										}
										v350 = v316
										v351 = v349
										v355 = v348
										for {
											v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
											if v360 == int32(92) {
												v363 = int32(23644)
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
												v397 = v350 + int32(2)
											} else {
												if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v373 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
													v375 = int32(7)
													v377 = int32(48)
													v378 = v360&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
													v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
													v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
													v397 = v350 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
													v397 = v350 + int32(1)
												}
											}
											v398 = int32(1)
											v401 = v355 - v398
											if v401 != 0 {
												v350 = v397
												v351 = v351 + v398
												v355 = v401
												continue
											} else {
												break
											}
											break
										}
										v402 = v397
										v406 = v316
									} else {
										v346 = base.B2i32(v324 == int32(18)) << (uint(int32(4)) % 32)
										if v346 != 0 {
											v348 = v346
											if v320 != 0 {
												v349 = v190
											} else {
												v349 = v188
											}
											v350 = v316
											v351 = v349
											v355 = v348
											for {
												v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
												if v360 == int32(92) {
													v363 = int32(23644)
													*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
													v397 = v350 + int32(2)
												} else {
													if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
														v373 = int32(92)
														*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
														v375 = int32(7)
														v377 = int32(48)
														v378 = v360&v375 | v377
														*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
														v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
														*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
														v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
														*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
														v397 = v350 + int32(4)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
														v397 = v350 + int32(1)
													}
												}
												v398 = int32(1)
												v401 = v355 - v398
												if v401 != 0 {
													v350 = v397
													v351 = v351 + v398
													v355 = v401
													continue
												} else {
													break
												}
												break
											}
											v402 = v397
											v406 = v316
										} else {
											v402 = v316
											v406 = v316
										}
									}
								} else {
									v335 = int32(1)
									if v320 != 0 {
										v346 = int32(base.Ui32(v318)>>(uint(v335)%32)) - v335
									} else {
										v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
										v346 = int32(base.Ui32(v339)>>(uint(int32(2))%32)) - int32(4)
									}
									if v346 != 0 {
										v348 = v346
										if v320 != 0 {
											v349 = v190
										} else {
											v349 = v188
										}
										v350 = v316
										v351 = v349
										v355 = v348
										for {
											v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
											if v360 == int32(92) {
												v363 = int32(23644)
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
												v397 = v350 + int32(2)
											} else {
												if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v373 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
													v375 = int32(7)
													v377 = int32(48)
													v378 = v360&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
													v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
													v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
													v397 = v350 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
													v397 = v350 + int32(1)
												}
											}
											v398 = int32(1)
											v401 = v355 - v398
											if v401 != 0 {
												v350 = v397
												v351 = v351 + v398
												v355 = v401
												continue
											} else {
												break
											}
											break
										}
										v402 = v397
										v406 = v316
									} else {
										v402 = v316
										v406 = v316
									}
								}
								v412 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
								m.G0 = v13 + int32(16)
								return v406
							}
						}
					} else {
						v314 = int64(1)
						v316 = F_palloc(m, base.I32_wrap_i64(v314))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v319 = int32(1)
							v320 = v318 & v319
							if v318 == v319 {
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
								if base.Ui32((v324-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v348 = int32(4)
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v346 = base.B2i32(v324 == int32(18)) << (uint(int32(4)) % 32)
									if v346 != 0 {
										v348 = v346
										if v320 != 0 {
											v349 = v190
										} else {
											v349 = v188
										}
										v350 = v316
										v351 = v349
										v355 = v348
										for {
											v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
											if v360 == int32(92) {
												v363 = int32(23644)
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
												v397 = v350 + int32(2)
											} else {
												if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v373 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
													v375 = int32(7)
													v377 = int32(48)
													v378 = v360&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
													v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
													v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
													v397 = v350 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
													v397 = v350 + int32(1)
												}
											}
											v398 = int32(1)
											v401 = v355 - v398
											if v401 != 0 {
												v350 = v397
												v351 = v351 + v398
												v355 = v401
												continue
											} else {
												break
											}
											break
										}
										v402 = v397
										v406 = v316
									} else {
										v402 = v316
										v406 = v316
									}
								}
							} else {
								v335 = int32(1)
								if v320 != 0 {
									v346 = int32(base.Ui32(v318)>>(uint(v335)%32)) - v335
								} else {
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v346 = int32(base.Ui32(v339)>>(uint(int32(2))%32)) - int32(4)
								}
								if v346 != 0 {
									v348 = v346
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v402 = v316
									v406 = v316
								}
							}
							v412 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
							m.G0 = v13 + int32(16)
							return v406
						}
					}
				}
			} else {
				v207 = int32(1)
				if v193 != 0 {
					v217 = int32(base.Ui32(v191)>>(uint(v207)%32)) - v207
				} else {
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v217 = int32(base.Ui32(v211)>>(uint(int32(2))%32)) - int32(4)
				}
				if v217 != 0 {
					v220 = v217
					if v193 != 0 {
						v221 = v190
					} else {
						v221 = v188
					}
					v222 = int32(1)
					if v220 == v222 {
						v276 = v221
						v285 = int64(1)
					} else {
						v231 = v221
						v232 = int32(0)
						v240 = int64(1)
						for {
							v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
							if base.Ui32((v244-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v251 = int64(4)
							} else {
								v251 = int64(1)
							}
							if v244 == int32(92) {
								v254 = int64(2)
							} else {
								v254 = v251
							}
							v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
							if base.Ui32((v259-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v266 = int64(4)
							} else {
								v266 = int64(1)
							}
							if v259 == int32(92) {
								v269 = int64(2)
							} else {
								v269 = v266
							}
							v270 = v240 + v254 + v269
							v271 = int32(2)
							v272 = v231 + v271
							v274 = v232 + v271
							if v274 != v220&int32(-2) {
								v231 = v272
								v232 = v274
								v240 = v270
								continue
							} else {
								break
							}
							break
						}
						v276 = v272
						v285 = v270
					}
					if v220&v222 != 0 {
						v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
						if base.Ui32((v289-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v296 = int64(4)
						} else {
							v296 = int64(1)
						}
						if v289 == int32(92) {
							v299 = int64(2)
						} else {
							v299 = v296
						}
						v302 = v285 + v299
					} else {
						v302 = v285
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v302) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v421 = m.ExcPending
						if v421 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v424 = m.ExcPending
							if v424 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(400095), int32(0))
								mBase = m.M
								v428 = m.ExcPending
								if v428 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499834), int32(438), int32(66423))
									mBase = m.M
									v433 = m.ExcPending
									if v433 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v314 = v302
						v316 = F_palloc(m, base.I32_wrap_i64(v314))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v319 = int32(1)
							v320 = v318 & v319
							if v318 == v319 {
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
								if base.Ui32((v324-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v348 = int32(4)
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v346 = base.B2i32(v324 == int32(18)) << (uint(int32(4)) % 32)
									if v346 != 0 {
										v348 = v346
										if v320 != 0 {
											v349 = v190
										} else {
											v349 = v188
										}
										v350 = v316
										v351 = v349
										v355 = v348
										for {
											v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
											if v360 == int32(92) {
												v363 = int32(23644)
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
												v397 = v350 + int32(2)
											} else {
												if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v373 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
													v375 = int32(7)
													v377 = int32(48)
													v378 = v360&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
													v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
													v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
													*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
													v397 = v350 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
													v397 = v350 + int32(1)
												}
											}
											v398 = int32(1)
											v401 = v355 - v398
											if v401 != 0 {
												v350 = v397
												v351 = v351 + v398
												v355 = v401
												continue
											} else {
												break
											}
											break
										}
										v402 = v397
										v406 = v316
									} else {
										v402 = v316
										v406 = v316
									}
								}
							} else {
								v335 = int32(1)
								if v320 != 0 {
									v346 = int32(base.Ui32(v318)>>(uint(v335)%32)) - v335
								} else {
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v346 = int32(base.Ui32(v339)>>(uint(int32(2))%32)) - int32(4)
								}
								if v346 != 0 {
									v348 = v346
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v402 = v316
									v406 = v316
								}
							}
							v412 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
							m.G0 = v13 + int32(16)
							return v406
						}
					}
				} else {
					v314 = int64(1)
					v316 = F_palloc(m, base.I32_wrap_i64(v314))
					mBase = m.M
					v317 = m.ExcPending
					if v317 != 0 {
						return int32(0)
					} else {
						v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
						v319 = int32(1)
						v320 = v318 & v319
						if v318 == v319 {
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
							if base.Ui32((v324-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v348 = int32(4)
								if v320 != 0 {
									v349 = v190
								} else {
									v349 = v188
								}
								v350 = v316
								v351 = v349
								v355 = v348
								for {
									v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
									if v360 == int32(92) {
										v363 = int32(23644)
										*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
										v397 = v350 + int32(2)
									} else {
										if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v373 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
											v375 = int32(7)
											v377 = int32(48)
											v378 = v360&v375 | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
											v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
											v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
											v397 = v350 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
											v397 = v350 + int32(1)
										}
									}
									v398 = int32(1)
									v401 = v355 - v398
									if v401 != 0 {
										v350 = v397
										v351 = v351 + v398
										v355 = v401
										continue
									} else {
										break
									}
									break
								}
								v402 = v397
								v406 = v316
							} else {
								v346 = base.B2i32(v324 == int32(18)) << (uint(int32(4)) % 32)
								if v346 != 0 {
									v348 = v346
									if v320 != 0 {
										v349 = v190
									} else {
										v349 = v188
									}
									v350 = v316
									v351 = v349
									v355 = v348
									for {
										v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
										if v360 == int32(92) {
											v363 = int32(23644)
											*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
											v397 = v350 + int32(2)
										} else {
											if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v373 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
												v375 = int32(7)
												v377 = int32(48)
												v378 = v360&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
												v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
												v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
												*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
												v397 = v350 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
												v397 = v350 + int32(1)
											}
										}
										v398 = int32(1)
										v401 = v355 - v398
										if v401 != 0 {
											v350 = v397
											v351 = v351 + v398
											v355 = v401
											continue
										} else {
											break
										}
										break
									}
									v402 = v397
									v406 = v316
								} else {
									v402 = v316
									v406 = v316
								}
							}
						} else {
							v335 = int32(1)
							if v320 != 0 {
								v346 = int32(base.Ui32(v318)>>(uint(v335)%32)) - v335
							} else {
								v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v346 = int32(base.Ui32(v339)>>(uint(int32(2))%32)) - int32(4)
							}
							if v346 != 0 {
								v348 = v346
								if v320 != 0 {
									v349 = v190
								} else {
									v349 = v188
								}
								v350 = v316
								v351 = v349
								v355 = v348
								for {
									v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
									if v360 == int32(92) {
										v363 = int32(23644)
										*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v363)
										v397 = v350 + int32(2)
									} else {
										if base.Ui32((v360-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v373 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v373)
											v375 = int32(7)
											v377 = int32(48)
											v378 = v360&v375 | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+3)) = uint8(v378)
											v383 = int32(base.Ui32(v360)>>(uint(int32(6))%32)) | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)) = uint8(v383)
											v390 = int32(base.Ui32(v360)>>(uint(int32(3))%32))&v375 | v377
											*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)) = uint8(v390)
											v397 = v350 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
											v397 = v350 + int32(1)
										}
									}
									v398 = int32(1)
									v401 = v355 - v398
									if v401 != 0 {
										v350 = v397
										v351 = v351 + v398
										v355 = v401
										continue
									} else {
										break
									}
									break
								}
								v402 = v397
								v406 = v316
							} else {
								v402 = v316
								v406 = v316
							}
						}
						v412 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
						m.G0 = v13 + int32(16)
						return v406
					}
				}
			}
		case 1:
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v22 == int32(1) {
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
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
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v57 = F_palloc(m, v52<<(uint(int32(1))%32)+int32(3))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v59 = int32(30812)
				*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v59)
				v62 = v57 + int32(2)
				v63 = int32(1)
				v64 = v16 + v63
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				v69 = v67 & v63
				if v69 != 0 {
					v70 = v64
				} else {
					v70 = v16 + int32(4)
				}
				if v67 == int32(1) {
					v73 = int32(4)
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
					if v75&int32(254) == int32(2) {
						v84 = v73
					} else {
						v84 = base.B2i32(v75 == int32(18)) << (uint(v73) % 32)
					}
					if v75 == int32(1) {
						v87 = v73
					} else {
						v87 = v84
					}
					v98 = v87
				} else {
					v88 = int32(1)
					if v69 != 0 {
						v98 = int32(base.Ui32(v67)>>(uint(v88)%32)) - v88
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						v98 = int32(base.Ui32(v92)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v102 = v70 + v98
				if base.Ui32(v102) <= base.Ui32(v70) {
				} else {
					v105 = v98 & int32(3)
					if v105 != 0 {
						v106 = v70
						v108 = v62
						v109 = int32(0)
						for {
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
							v113 = int32(1)
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(v113)%32))+uint32(_consts[408]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v117)
							v120 = v108 + int32(2)
							v122 = v106 + v113
							v124 = v109 + v113
							if v124 != v105 {
								v106 = v122
								v108 = v120
								v109 = v124
								continue
							} else {
								break
							}
							break
						}
						v126 = v122
						v128 = v120
					} else {
						v126 = v70
						v128 = v62
					}
					if base.Ui32(v98-int32(1)) < base.Ui32(int32(3)) {
					} else {
						v136 = v126
						v138 = v128
						for {
							v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
							v143 = int32(1)
							v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142<<(uint(v143)%32))+uint32(_consts[408]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v147)
							v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
							v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149<<(uint(v143)%32))+uint32(_consts[408]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)) = uint16(v154)
							v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2)))
							v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156<<(uint(v143)%32))+uint32(_consts[408]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)) = uint16(v161)
							v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+3)))
							v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163<<(uint(v143)%32))+uint32(_consts[408]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v138)+6)) = uint16(v168)
							v173 = v136 + int32(4)
							if v173 != v102 {
								v136 = v173
								v138 = v138 + int32(8)
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v402 = v62 + base.I32_wrap_i64(base.I64_extend_i32_u(v98)<<(uint(int64(1))%64))
				v406 = v57
				v412 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v412)
				m.G0 = v13 + int32(16)
				return v406
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v437 = m.ExcPending
			if v437 != 0 {
				return int32(0)
			} else {
				v439 = *(*int32)(unsafe.Add(mBase, _consts[1058]))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v439
				F_errmsg_internal(m, int32(483120), v13)
				mBase = m.M
				v443 = m.ExcPending
				if v443 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499834), int32(469), int32(66423))
					mBase = m.M
					v448 = m.ExcPending
					if v448 != 0 {
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
func F_byteapos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v193 int32
	_ = v193
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	return v193
L16:
	;
	if v80 <= v46 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	if v75 <= int32(0) {
		v193 = int32(1)
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v80 = int32(4)
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v62 = int32(1)
	if v47&v62 != 0 {
		v75 = int32(base.Ui32(v47)>>(uint(v62)%32)) - v62
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v75 = base.B2i32(v51 == int32(18)) << (uint(int32(4)) % 32)
	goto L17
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L23:
	;
	v80 = v75
	goto L16
L24:
	;
	v83 = int32(1)
	if v16&v83 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v193 = int32(0)
	goto L15
L27:
	;
	v89 = v83
	goto L29
L28:
	;
	v89 = int32(4)
	goto L29
L29:
	;
	v92 = int32(1)
	if v47&v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = v92
	goto L32
L31:
	;
	v96 = int32(4)
	goto L32
L32:
	;
	v97 = v14 + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v99 = int32(0)
	v100 = v9 + v89
	goto L33
L33:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v98 != v106 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L26
L35:
	;
	v173 = int32(1)
	v176 = v99 + v173
	if v176 != v46-v80+v83 {
		v99 = v176
		v100 = v100 + v173
		goto L33
	} else {
		goto L56
	}
L36:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v80) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v169 != 0 {
		goto L35
	} else {
		goto L55
	}
L38:
	;
	v169 = int32(0)
	goto L37
L39:
	;
	v143 = v138
	v144 = v139
	v145 = v140
	goto L49
L40:
	;
	if (v100|v97)&int32(3) != 0 {
		v138 = v100
		v139 = v97
		v140 = v80
		goto L39
	} else {
		goto L43
	}
L41:
	;
	v131 = v100
	v132 = v97
	v133 = v80
	goto L42
L42:
	;
	if v133 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L43:
	;
	v115 = v100
	v116 = v97
	v117 = v80
	goto L44
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v120 != v121 {
		v138 = v115
		v139 = v116
		v140 = v117
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v131 = v126
	v132 = v124
	v133 = v128
	goto L42
L46:
	;
	v123 = int32(4)
	v124 = v116 + v123
	v126 = v115 + v123
	v128 = v117 - v123
	if base.Ui32(int32(3)) < base.Ui32(v128) {
		v115 = v126
		v116 = v124
		v117 = v128
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v138 = v131
	v139 = v132
	v140 = v133
	goto L39
L49:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 == v149 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v169 = v148 - v149
	goto L37
L51:
	;
	v151 = int32(1)
	v156 = v145 - v151
	if v156 != 0 {
		v143 = v143 + v151
		v144 = v144 + v151
		v145 = v156
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	goto L38
L55:
	;
	return v99 + int32(1)
L56:
	;
	goto L34
}
func F_byteartrim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v12 = F_dobyteatrim(m, v3, v8, int32(0), int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
