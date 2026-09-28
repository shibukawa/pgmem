package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuumLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v105 float64
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v590 int32
	_ = v590
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v890 int64
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(_a_F_vacuumLeafPage_0)
	m.G0 = v22
	if l2 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	v45 = int32(0)
	base.MemoryFill(m, v22+int32(448), v45, int32(818))
	base.MemoryFill(m, v22+int32(32), v45, int32(409))
	if base.Ui32(v42) < base.Ui32(int32(25)) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(l2^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L1
L3:
	;
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[1]))
	v41 = v35 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L17
	} else {
		goto L130
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L17
	} else {
		goto L127
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L17
	} else {
		goto L120
	}
L8:
	;
	m.G0 = v22 + int32(_a_F_vacuumLeafPage_0)
	return
L9:
	;
	v58 = int32(base.Ui32(v42+int32(_a_F_vacuumLeafPage_1)) >> (uint(int32(2)) % 32))
	v60 = v58 & int32(_a_F_vacuumLeafPage_2)
	if v60 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v64 = l0 + int32(100)
	v67 = int32(1)
	v75 = v67
	v79 = v67
	v80 = v5
	goto L11
L11:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(20)+v75<<(uint(int32(2))%32))))
	v94 = v41 + v91&int32(_a_F_vacuumLeafPage_3)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	switch v95 & int32(3) {
	case 0:
		goto L15
	case 1:
		goto L14
	default:
		v234 = v80
		goto L13
	}
L12:
	;
	if v234 == int32(0) {
		goto L8
	} else {
		goto L47
	}
L13:
	;
	v243 = v79 + int32(1)
	v245 = v243 & int32(_a_F_vacuumLeafPage_2)
	if base.Ui32(v245) <= base.Ui32(v60) {
		v75 = v245
		v79 = v243
		v80 = v234
		goto L11
	} else {
		goto L46
	}
L14:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v138 = int32(3)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.B2i32(base.Ui32(v137) < base.Ui32(v138))|base.B2i32(base.Ui32(v140) < base.Ui32(v138)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, v94+int32(6), v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v126 = v124 & int32(_a_F_vacuumLeafPage_4)
	if v126 == int32(0) {
		v234 = v123
		goto L13
	} else {
		goto L23
	}
L17:
	;
	return
L18:
	;
	if v102 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v104)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v104)+16)) = base.F64_add(v105, float64(1))
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v75))) = uint8(v112)
	v123 = v80 + v112
	goto L16
L20:
	;
	goto L21
L21:
	;
	if l3 != 0 {
		v123 = v80
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v116)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v116)+8)) = base.F64_add(v117, float64(1))
	v123 = v80
	goto L16
L23:
	;
	if base.Ui32(v60) < base.Ui32(v126) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v134 = v22 + int32(448) + v126<<(uint(int32(1))%32)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134))))
	if v135 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v79)
	v234 = v123
	goto L13
L26:
	;
	v151 = v94 + int32(6)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v152 != 0 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	if int32(0) <= v137-v140 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v137) < base.Ui32(v140) {
		v234 = v80
		goto L13
	} else {
		goto L31
	}
L30:
	;
	v234 = v80
	goto L13
L31:
	;
	goto L26
L32:
	;
	v157 = v152
	goto L35
L33:
	;
	v198 = v64
	goto L34
L34:
	;
	v212 = F_palloc(m, int32(12))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L17
	} else {
		goto L45
	}
L35:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151))))
	v174 = int32(16)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+2)))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	if v172|v173<<(uint(v174)%32) == v177|v178<<(uint(v174)%32) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v198 = v157 + int32(8)
	goto L34
L37:
	;
	if v188 != 0 {
		v234 = v80
		goto L13
	} else {
		goto L43
	}
L38:
	;
	goto L37
L39:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+4)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+4)))
	if v184 == v185 {
		v188 = int32(1)
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v188 = int32(0)
	goto L38
L42:
	;
	goto L41
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	if v189 != 0 {
		v157 = v189
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v212)+4)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v216
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+8)) = v218
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)) = uint8(v218)
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v212
	v234 = v80
	goto L13
L46:
	;
	goto L12
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))) = int64(0)
	v252 = v41 + int32(20)
	v253 = int32(1)
	v259 = v253
	v260 = v253
	v266 = int32(0)
	v268 = v5
	v270 = v5
	v272 = v5
	goto L48
L48:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v252+v260<<(uint(int32(2))%32))))
	v281 = v41 + v278&int32(_a_F_vacuumLeafPage_3)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v282&int32(3) != 0 {
		v459 = v266
		v461 = v268
		v463 = v270
		v465 = v272
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v475 = int32(_a_F_vacuumLeafPage_2)
	v476 = v465 & v475
	v478 = v459 & v475
	v480 = v463 & v475
	if v234 != v476+(v478+v480) {
		goto L5
	} else {
		goto L78
	}
L50:
	;
	v469 = v259 + int32(1)
	v470 = int32(_a_F_vacuumLeafPage_2)
	v471 = v469 & v470
	if base.Ui32(v471) <= base.Ui32(v58&v470) {
		v259 = v469
		v260 = v471
		v266 = v459
		v268 = v461
		v270 = v463
		v272 = v465
		goto L48
	} else {
		goto L77
	}
L51:
	;
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(448)+v260<<(uint(int32(1))%32)))))
	if v290 != 0 {
		v459 = v266
		v461 = v268
		v463 = v270
		v465 = v272
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v291 = int32(0)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v260))))
	if v296 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v297 = v291
	goto L55
L54:
	;
	v297 = v259
	goto L55
L55:
	;
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+4)))
	v300 = v298 & int32(_a_F_vacuumLeafPage_4)
	if v300 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v442 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(_a_F_vacuumLeafPage_5)+v270&int32(_a_F_vacuumLeafPage_2)<<(uint(v442)%32)))) = uint16(v259)
	v447 = v270 + v442
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))) = uint16(v447)
	v459 = v429
	v461 = v431
	v463 = v447
	v465 = v435
	goto L50
L57:
	;
	v305 = v300
	v307 = v291
	v308 = v297
	v311 = v266
	v313 = v268
	v317 = v272
	goto L60
L58:
	;
	goto L59
L59:
	;
	if v297&int32(_a_F_vacuumLeafPage_2) != 0 {
		v459 = v266
		v461 = v268
		v463 = v270
		v465 = v272
		goto L50
	} else {
		goto L76
	}
L60:
	;
	v321 = v305 & int32(_a_F_vacuumLeafPage_2)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v252+v321<<(uint(int32(2))%32))))
	v328 = v41 + v325&int32(_a_F_vacuumLeafPage_3)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v329&int32(3) != 0 {
		goto L6
	} else {
		goto L62
	}
L61:
	;
	if v387&int32(_a_F_vacuumLeafPage_2) == int32(0) {
		v429 = v389
		v431 = v390
		v435 = v391
		goto L56
	} else {
		goto L74
	}
L62:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v321))))
	if v335 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+4)))
	v394 = v392 & int32(_a_F_vacuumLeafPage_4)
	if v394 != 0 {
		v305 = v394
		v307 = v335
		v308 = v387
		v311 = v389
		v313 = v390
		v317 = v391
		goto L60
	} else {
		goto L73
	}
L64:
	;
	v342 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(_a_F_vacuumLeafPage_6)+v311&int32(_a_F_vacuumLeafPage_2)<<(uint(v342)%32)))) = uint16(v305)
	v347 = v311 + v342
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[3]))) = uint16(v347)
	v387 = v308
	v389 = v347
	v390 = v313
	v391 = v317
	goto L63
L65:
	;
	goto L66
L66:
	;
	if v308&int32(_a_F_vacuumLeafPage_2) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v353 = int32(1)
	v356 = v317 << (uint(v353) % 32) & int32(_a_F_vacuumLeafPage_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v356+(v22+int32(2912))))) = uint16(v259)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(3728)+v356))) = uint16(v305)
	v366 = v317 + v353
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))) = uint16(v366)
	v387 = v259
	v389 = v311
	v390 = v313
	v391 = v366
	goto L63
L68:
	;
	goto L69
L69:
	;
	if v307&int32(1) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v370 = int32(1)
	v373 = v313 << (uint(v370) % 32) & int32(_a_F_vacuumLeafPage_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v373+(v22+int32(1280))))) = uint16(v305)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2096)+v373))) = uint16(v308)
	v383 = v313 + v370
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))) = uint16(v383)
	v386 = v383
	goto L72
L71:
	;
	v386 = v313
	goto L72
L72:
	;
	v387 = v305
	v389 = v311
	v390 = v386
	v391 = v317
	goto L63
L73:
	;
	goto L61
L74:
	;
	if v335 == int32(0) {
		v459 = v389
		v461 = v390
		v463 = v270
		v465 = v391
		goto L50
	} else {
		goto L75
	}
L75:
	;
	v401 = int32(1)
	v404 = v390 << (uint(v401) % 32) & int32(_a_F_vacuumLeafPage_7)
	v408 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v404+(v22+int32(1280))))) = uint16(v408)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2096)+v404))) = uint16(v387)
	v415 = v390 + v401
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))) = uint16(v415)
	v459 = v389
	v461 = v415
	v463 = v270
	v465 = v391
	goto L50
L76:
	;
	v429 = v266
	v431 = v268
	v435 = v272
	goto L56
L77:
	;
	goto L49
L78:
	;
	v484 = int32(_a_F_vacuumLeafPage_8)
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6])) = v486 + int32(1)
	v491 = l0 + int32(16)
	v494 = int32(2)
	F_spgPageIndexMultiDelete(m, v491, v41, v22+int32(_a_F_vacuumLeafPage_5), v480, v494, v494, int32(-1), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	v502 = int32(3)
	F_spgPageIndexMultiDelete(m, v491, v41, v22+int32(_a_F_vacuumLeafPage_6), v478, v502, v502, int32(-1), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	if v476 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v511 = v41 + int32(20)
	v512 = int32(0)
	if v476 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v664 = int32(0)
	goto L83
L83:
	;
	v665 = int32(3)
	F_spgPageIndexMultiDelete(m, v491, v41, v22+int32(3728), v664, v665, v665, int32(-1), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L17
	} else {
		goto L92
	}
L84:
	;
	v664 = v465 & int32(_a_F_vacuumLeafPage_2)
	goto L83
L85:
	;
	v526 = v512
	v530 = int32(0)
	goto L88
L86:
	;
	v590 = v512
	goto L87
L87:
	;
	v604 = v590 << (uint(int32(1)) % 32)
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604+(v22+int32(3728))))))
	v609 = int32(2)
	v611 = v511 + v608<<(uint(v609)%32)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2912)+v604))))
	v619 = v511 + v616<<(uint(v609)%32)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v612
	goto L84
L88:
	;
	v540 = v526 << (uint(int32(1)) % 32)
	v542 = v22 + int32(3728)
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540+v542))))
	v545 = int32(2)
	v547 = v511 + v544<<(uint(v545)%32)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v550 = v22 + int32(2912)
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v550+v540))))
	v555 = v511 + v552<<(uint(v545)%32)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v548
	v560 = v540 | v545
	v562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542+v560))))
	v565 = v511 + v562<<(uint(v545)%32)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v550+v560))))
	v573 = v511 + v570<<(uint(v545)%32)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v566
	v578 = v526 + v545
	v580 = v530 + v545
	if v580 != v476&int32(_a_F_vacuumLeafPage_9) {
		v526 = v578
		v530 = v580
		goto L88
	} else {
		goto L90
	}
L89:
	;
	if v476&int32(1) == int32(0) {
		goto L84
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v590 = v578
	goto L87
L92:
	;
	v672 = v461 & int32(_a_F_vacuumLeafPage_2)
	if v672 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L17
	} else {
		goto L102
	}
L94:
	;
	v676 = v41 + int32(20)
	v677 = int32(0)
	if v672 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v689 = v677
	v695 = int32(0)
	goto L98
L96:
	;
	v761 = v677
	goto L97
L97:
	;
	v777 = v761 << (uint(int32(1)) % 32)
	v781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777+(v22+int32(2096))))))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v676+v781<<(uint(int32(2))%32))))
	v788 = v41 + v785&int32(_a_F_vacuumLeafPage_3)
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(1280)+v777))))
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v788)+4)))
	v798 = v792&int32(_a_F_vacuumLeafPage_4) | v795&int32(_a_F_vacuumLeafPage_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v788)+4)) = uint16(v798)
	goto L93
L98:
	;
	v705 = v689 << (uint(int32(1)) % 32)
	v707 = v22 + int32(2096)
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v705+v707))))
	v710 = int32(2)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v676+v709<<(uint(v710)%32))))
	v714 = int32(_a_F_vacuumLeafPage_3)
	v716 = v41 + v713&v714
	v718 = v22 + int32(1280)
	v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718+v705))))
	v721 = int32(_a_F_vacuumLeafPage_4)
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+4)))
	v724 = int32(_a_F_vacuumLeafPage_10)
	v726 = v720&v721 | v723&v724
	*(*uint16)(unsafe.Add(mBase, uint32(v716)+4)) = uint16(v726)
	v729 = v705 | v710
	v731 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v707+v729))))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v676+v731<<(uint(v710)%32))))
	v738 = v41 + v735&v714
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718+v729))))
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v738)+4)))
	v748 = v742&v721 | v745&v724
	*(*uint16)(unsafe.Add(mBase, uint32(v738)+4)) = uint16(v748)
	v751 = v689 + v710
	v753 = v695 + v710
	if v753 != v672&int32(_a_F_vacuumLeafPage_9) {
		v689 = v751
		v695 = v753
		goto L98
	} else {
		goto L100
	}
L99:
	;
	if v461&int32(1) == int32(0) {
		goto L93
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v761 = v751
	goto L97
L102:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821)+118)))
	if v822 != int32(112) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v895 = int32(_a_F_vacuumLeafPage_8)
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6])) = v897 - int32(1)
	goto L8
L104:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[7]))
	if v826 <= int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v829 != 0 {
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L17
	} else {
		goto L110
	}
L108:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v830 != 0 {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[8]))) = v833
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[9]))) = uint8(v835)
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_11), int32(16))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))))
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_5), v844<<(uint(int32(1))%32))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[3]))))
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_6), v851<<(uint(int32(1))%32))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L17
	} else {
		goto L113
	}
L113:
	;
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))))
	F_XLogRegisterData(m, v22+int32(3728), v858<<(uint(int32(1))%32))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L17
	} else {
		goto L114
	}
L114:
	;
	v865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))))
	F_XLogRegisterData(m, v22+int32(2912), v865<<(uint(int32(1))%32))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L17
	} else {
		goto L115
	}
L115:
	;
	v872 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))))
	F_XLogRegisterData(m, v22+int32(2096), v872<<(uint(int32(1))%32))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L17
	} else {
		goto L116
	}
L116:
	;
	v879 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))))
	F_XLogRegisterData(m, v22+int32(1280), v879<<(uint(int32(1))%32))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L17
	} else {
		goto L117
	}
L117:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L17
	} else {
		goto L118
	}
L118:
	;
	v890 = F_XLogInsert(m, int32(16), int32(96))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotl(v890, int64(32))
	goto L103
L120:
	;
	if l2 < int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v947 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_12), v22+int32(16))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L17
	} else {
		goto L125
	}
L122:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[10]))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v931+(l2^int32(-1))*int32(56))+16))
	v946 = v937
	goto L121
L123:
	;
	goto L124
L124:
	;
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[11]))
	v940 = int32(56)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v939+l2*v940-v940)+16))
	v946 = v945
	goto L121
L125:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(179), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L17
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v966 & int32(3)
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_15), v22)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L17
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(266), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_16), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(326), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_delay_point(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v92 float64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 float64
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 float64
	_ = v114
	var v124 float64
	_ = v124
	var v132 float64
	_ = v132
	var v134 float64
	_ = v134
	var v136 float64
	_ = v136
	var v138 int32
	_ = v138
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v384 int32
	_ = v384
	v8 = float64(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L87
	}
L7:
	;
	m.G0 = v15 + int32(16)
	return
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[1]))
	if int32(0) <= v24 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_parallel_vacuum_update_shared_delay_params(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[2])))
	if v30 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[3]))
	if v34 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[3]))
	v39 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[4]))
	if base.B2i32(v38 == v39)|base.B2i32(v42 != int32(4)) == v39 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[5]))
	if v63 != 0 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[3])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v30 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L21:
	;
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_parallel_vacuum_propagate_shared_delay_params(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[2])))
	if v59 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	goto L17
L26:
	;
	if base.F64_gt(v124, float64(0)) == int32(0) {
		goto L7
	} else {
		goto L34
	}
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(_a_F_vacuum_delay_point_0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7]))
	v70 = base.AtomicRmwAdd32(m, v63, int32(0), v68)
	v71 = int32(_a_F_vacuum_delay_point_1)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8]))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7]))
	v76 = v73 + v75
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8])) = v76
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9]))
	if base.Ui32(v68+v70) < base.Ui32(v79) {
		v102 = v8
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7]))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9]))
	if v109 < v111 {
		goto L7
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7])) = int32(0)
	v124 = v102
	goto L26
L31:
	;
	v82 = base.F64_convert_i32_s(v76)
	v83 = base.F64_convert_i32_s(v79)
	if base.F64_gt(v82, base.F64_mul(base.F64_div(v83, base.F64_convert_i32_s(v66)), float64(0.5))) == int32(0) {
		v102 = v8
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v92 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[10]))
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[5]))
	v95 = int32(0)
	v96 = base.AtomicRmwSub32(m, v94, v95, v76)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8])) = v95
	v102 = base.F64_div(base.F64_mul(v92, v82), v83)
	goto L30
L33:
	;
	v114 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[10]))
	v124 = base.F64_div(base.F64_mul(v114, base.F64_convert_i32_s(v109)), base.F64_convert_i32_s(v111))
	goto L26
L34:
	;
	v132 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[10]))
	v134 = base.F64_mul(v132, float64(4))
	if base.F64_gt(v124, v134) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v136 = v134
	goto L37
L36:
	;
	v136 = v124
	goto L37
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[11])))
	if v138 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F___clock_gettime(m, int32(1), v15)
	mBase = m.M
	v143 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v148 = v143 + v144*int64(1000000000)
	goto L40
L39:
	;
	v148 = int64(0)
	goto L40
L40:
	;
	v149 = int32(_a_F_vacuum_delay_point_2)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = int32(150994952)
	F_pg_usleep(m, base.I32_trunc_sat_f64_s(base.F64_mul(v136, float64(1000))))
	mBase = m.M
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = int32(0)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[11])))
	if v162 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[13])))
	if v296 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L42:
	;
	F___clock_gettime(m, int32(1), v15)
	mBase = m.M
	v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v171 = v167 + v168*int64(1000000000)
	v172 = v171 - v148
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[1]))
	if int32(0) <= v174 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v177 = int32(_a_F_vacuum_delay_point_6)
	v179 = *(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20]))
	v180 = v179 + v172
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20])) = v180
	v183 = *(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[21]))
	if v171-v183 < int64(1000000000) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if l0 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	F_pgstat_progress_parallel_incr_param(m, int32(10), v180)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[21])) = v171
	goto L41
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[22]))
	if v198 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[22]))
	if v247 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	goto L41
L52:
	;
	goto L51
L53:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[23])))
	if v202&int32(1) == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v207 = int32(_a_F_vacuum_delay_point_7)
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	v210 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v209 + v210
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v213 + v210
	v217 = int32(0)
	v219 = int32(_a_F_vacuum_delay_point_8)
	v220 = base.AtomicRmwOr32(m, v217, v219, v217)
	v225 = v198 + int32(296)
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v225)))
	*(*int64)(unsafe.Add(mBase, uint32(v225))) = v226 + v172
	v232 = base.AtomicRmwOr32(m, v217, v219, v217)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v233 + v210
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v239 - v210
	goto L52
L55:
	;
	goto L41
L56:
	;
	goto L55
L57:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[23])))
	if v251&int32(1) == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v256 = int32(_a_F_vacuum_delay_point_7)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	v259 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v258 + v259
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v262 + v259
	v266 = int32(0)
	v268 = int32(_a_F_vacuum_delay_point_8)
	v269 = base.AtomicRmwOr32(m, v266, v268, v266)
	v274 = v247 + int32(312)
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = v275 + v172
	v281 = base.AtomicRmwOr32(m, v266, v268, v266)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v282 + v259
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v288 - v259
	goto L56
L59:
	;
	v299 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v304 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7])) = v304
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[14]))
	if v307 == v304 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v299 == int32(0) {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[4]))
	if v357 == int32(4) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[15]))
	if v312 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L78
	}
L67:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[16]))
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[17]))
	if int32(0) < v317 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v339 = v312
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9])) = v339
	goto L64
L70:
	;
	v322 = v317
	goto L72
L71:
	;
	v322 = v319
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9])) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v307)+32))
	if v324 == int32(0) {
		goto L64
	} else {
		goto L73
	}
L73:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[18]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_vacuum_delay_point[19])))
	if v329 <= int32(0) {
		goto L66
	} else {
		goto L74
	}
L74:
	;
	v332 = int32(1)
	v333 = base.I32_div_s(v322, v329)
	if v333 <= v332 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v336 = v332
	goto L77
L76:
	;
	v336 = v333
	goto L77
L77:
	;
	v339 = v336
	goto L69
L78:
	;
	F_errmsg_internal(m, int32(_a_F_vacuum_delay_point_3), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_vacuum_delay_point_4), int32(1782), int32(_a_F_vacuum_delay_point_5))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_parallel_vacuum_propagate_shared_delay_params(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v363 == int32(0) {
		goto L7
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	goto L7
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
