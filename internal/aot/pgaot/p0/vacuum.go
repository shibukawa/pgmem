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
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v593 int32
	_ = v593
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v764 int32
	_ = v764
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int64
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
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
	v984 = m.ExcPending
	if v984 != 0 {
		goto L17
	} else {
		goto L129
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L17
	} else {
		goto L126
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L17
	} else {
		goto L119
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
		v237 = v80
		goto L13
	}
L12:
	;
	if v237 == int32(0) {
		goto L8
	} else {
		goto L46
	}
L13:
	;
	v246 = v79 + int32(1)
	v248 = v246 & int32(_a_F_vacuumLeafPage_2)
	if base.Ui32(v248) <= base.Ui32(v60) {
		v75 = v248
		v79 = v246
		v80 = v237
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v138))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v137)) == int32(0) {
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
		v237 = v123
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
	v237 = v123
	goto L13
L26:
	;
	if v150 == int32(0) {
		v237 = v80
		goto L13
	} else {
		goto L30
	}
L27:
	;
	v150 = base.B2i32(base.Ui32(v138) <= base.Ui32(v137))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v150 = base.B2i32(int32(0) <= v137-v138)
	goto L26
L30:
	;
	v154 = v94 + int32(6)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v155 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v160 = v155
	goto L34
L32:
	;
	v201 = v64
	goto L33
L33:
	;
	v215 = F_palloc(m, int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L17
	} else {
		goto L44
	}
L34:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+2)))
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
	v177 = int32(16)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+2)))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	if v175|v176<<(uint(v177)%32) == v180|v181<<(uint(v177)%32) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v201 = v160 + int32(8)
	goto L33
L36:
	;
	if v191 != 0 {
		v237 = v80
		goto L13
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)))
	if v187 == v188 {
		v191 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v191 = int32(0)
	goto L37
L41:
	;
	goto L40
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v192 != 0 {
		v160 = v192
		goto L34
	} else {
		goto L43
	}
L43:
	;
	goto L35
L44:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+4)) = uint16(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v219
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v221
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+6)) = uint8(v221)
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v215
	v237 = v80
	goto L13
L45:
	;
	goto L12
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))) = int64(0)
	v255 = v41 + int32(20)
	v256 = int32(1)
	v262 = v256
	v263 = v256
	v269 = int32(0)
	v271 = v5
	v273 = v5
	v275 = v5
	goto L47
L47:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v255+v263<<(uint(int32(2))%32))))
	v284 = v41 + v281&int32(_a_F_vacuumLeafPage_3)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v285&int32(3) != 0 {
		v462 = v269
		v464 = v271
		v466 = v273
		v468 = v275
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v478 = int32(_a_F_vacuumLeafPage_2)
	v479 = v468 & v478
	v481 = v462 & v478
	v483 = v466 & v478
	if v237 != v479+(v481+v483) {
		goto L5
	} else {
		goto L77
	}
L49:
	;
	v472 = v262 + int32(1)
	v473 = int32(_a_F_vacuumLeafPage_2)
	v474 = v472 & v473
	if base.Ui32(v474) <= base.Ui32(v58&v473) {
		v262 = v472
		v263 = v474
		v269 = v462
		v271 = v464
		v273 = v466
		v275 = v468
		goto L47
	} else {
		goto L76
	}
L50:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(448)+v263<<(uint(int32(1))%32)))))
	if v293 != 0 {
		v462 = v269
		v464 = v271
		v466 = v273
		v468 = v275
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v294 = int32(0)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v263))))
	if v299 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v300 = v294
	goto L54
L53:
	;
	v300 = v262
	goto L54
L54:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+4)))
	v303 = v301 & int32(_a_F_vacuumLeafPage_4)
	if v303 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v445 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(_a_F_vacuumLeafPage_5)+v273&int32(_a_F_vacuumLeafPage_2)<<(uint(v445)%32)))) = uint16(v262)
	v450 = v273 + v445
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))) = uint16(v450)
	v462 = v432
	v464 = v434
	v466 = v450
	v468 = v438
	goto L49
L56:
	;
	v308 = v303
	v310 = v294
	v311 = v300
	v314 = v269
	v316 = v271
	v320 = v275
	goto L59
L57:
	;
	goto L58
L58:
	;
	if v300&int32(_a_F_vacuumLeafPage_2) != 0 {
		v462 = v269
		v464 = v271
		v466 = v273
		v468 = v275
		goto L49
	} else {
		goto L75
	}
L59:
	;
	v324 = v308 & int32(_a_F_vacuumLeafPage_2)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v255+v324<<(uint(int32(2))%32))))
	v331 = v41 + v328&int32(_a_F_vacuumLeafPage_3)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v332&int32(3) != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	if v390&int32(_a_F_vacuumLeafPage_2) == int32(0) {
		v432 = v392
		v434 = v393
		v438 = v394
		goto L55
	} else {
		goto L73
	}
L61:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v324))))
	if v338 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+4)))
	v397 = v395 & int32(_a_F_vacuumLeafPage_4)
	if v397 != 0 {
		v308 = v397
		v310 = v338
		v311 = v390
		v314 = v392
		v316 = v393
		v320 = v394
		goto L59
	} else {
		goto L72
	}
L63:
	;
	v345 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(_a_F_vacuumLeafPage_6)+v314&int32(_a_F_vacuumLeafPage_2)<<(uint(v345)%32)))) = uint16(v308)
	v350 = v314 + v345
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[3]))) = uint16(v350)
	v390 = v311
	v392 = v350
	v393 = v316
	v394 = v320
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v311&int32(_a_F_vacuumLeafPage_2) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v356 = int32(1)
	v359 = v320 << (uint(v356) % 32) & int32(_a_F_vacuumLeafPage_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v359+(v22+int32(2912))))) = uint16(v262)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(3728)+v359))) = uint16(v308)
	v369 = v320 + v356
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))) = uint16(v369)
	v390 = v262
	v392 = v314
	v393 = v316
	v394 = v369
	goto L62
L67:
	;
	goto L68
L68:
	;
	if v310&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v373 = int32(1)
	v376 = v316 << (uint(v373) % 32) & int32(_a_F_vacuumLeafPage_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v376+(v22+int32(1280))))) = uint16(v308)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2096)+v376))) = uint16(v311)
	v386 = v316 + v373
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))) = uint16(v386)
	v389 = v386
	goto L71
L70:
	;
	v389 = v316
	goto L71
L71:
	;
	v390 = v308
	v392 = v314
	v393 = v389
	v394 = v320
	goto L62
L72:
	;
	goto L60
L73:
	;
	if v338 == int32(0) {
		v462 = v392
		v464 = v393
		v466 = v273
		v468 = v394
		goto L49
	} else {
		goto L74
	}
L74:
	;
	v404 = int32(1)
	v407 = v393 << (uint(v404) % 32) & int32(_a_F_vacuumLeafPage_7)
	v411 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v407+(v22+int32(1280))))) = uint16(v411)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2096)+v407))) = uint16(v390)
	v418 = v393 + v404
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))) = uint16(v418)
	v462 = v392
	v464 = v418
	v466 = v273
	v468 = v394
	goto L49
L75:
	;
	v432 = v269
	v434 = v271
	v438 = v275
	goto L55
L76:
	;
	goto L48
L77:
	;
	v487 = int32(_a_F_vacuumLeafPage_8)
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6])) = v489 + int32(1)
	v494 = l0 + int32(16)
	v497 = int32(2)
	F_spgPageIndexMultiDelete(m, v494, v41, v22+int32(_a_F_vacuumLeafPage_5), v483, v497, v497, int32(-1), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L17
	} else {
		goto L78
	}
L78:
	;
	v505 = int32(3)
	F_spgPageIndexMultiDelete(m, v494, v41, v22+int32(_a_F_vacuumLeafPage_6), v481, v505, v505, int32(-1), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	if v479 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v514 = v41 + int32(20)
	v515 = int32(0)
	if v479 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v667 = int32(0)
	goto L82
L82:
	;
	v668 = int32(3)
	F_spgPageIndexMultiDelete(m, v494, v41, v22+int32(3728), v667, v668, v668, int32(-1), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L17
	} else {
		goto L91
	}
L83:
	;
	v667 = v468 & int32(_a_F_vacuumLeafPage_2)
	goto L82
L84:
	;
	v529 = v515
	v533 = int32(0)
	goto L87
L85:
	;
	v593 = v515
	goto L86
L86:
	;
	v607 = v593 << (uint(int32(1)) % 32)
	v611 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607+(v22+int32(3728))))))
	v612 = int32(2)
	v614 = v514 + v611<<(uint(v612)%32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(2912)+v607))))
	v622 = v514 + v619<<(uint(v612)%32)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v615
	goto L83
L87:
	;
	v543 = v529 << (uint(int32(1)) % 32)
	v545 = v22 + int32(3728)
	v547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v543+v545))))
	v548 = int32(2)
	v550 = v514 + v547<<(uint(v548)%32)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v553 = v22 + int32(2912)
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v553+v543))))
	v558 = v514 + v555<<(uint(v548)%32)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = v551
	v563 = v543 | v548
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545+v563))))
	v568 = v514 + v565<<(uint(v548)%32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v553+v563))))
	v576 = v514 + v573<<(uint(v548)%32)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v569
	v581 = v529 + v548
	v583 = v533 + v548
	if v583 != v479&int32(_a_F_vacuumLeafPage_9) {
		v529 = v581
		v533 = v583
		goto L87
	} else {
		goto L89
	}
L88:
	;
	if v479&int32(1) == int32(0) {
		goto L83
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v593 = v581
	goto L86
L91:
	;
	v675 = v464 & int32(_a_F_vacuumLeafPage_2)
	if v675 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L17
	} else {
		goto L101
	}
L93:
	;
	v679 = v41 + int32(20)
	v680 = int32(0)
	if v675 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v692 = v680
	v698 = int32(0)
	goto L97
L95:
	;
	v764 = v680
	goto L96
L96:
	;
	v780 = v764 << (uint(int32(1)) % 32)
	v784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v780+(v22+int32(2096))))))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v679+v784<<(uint(int32(2))%32))))
	v791 = v41 + v788&int32(_a_F_vacuumLeafPage_3)
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(1280)+v780))))
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791)+4)))
	v801 = v795&int32(_a_F_vacuumLeafPage_4) | v798&int32(_a_F_vacuumLeafPage_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v791)+4)) = uint16(v801)
	goto L92
L97:
	;
	v708 = v692 << (uint(int32(1)) % 32)
	v710 = v22 + int32(2096)
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708+v710))))
	v713 = int32(2)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v679+v712<<(uint(v713)%32))))
	v717 = int32(_a_F_vacuumLeafPage_3)
	v719 = v41 + v716&v717
	v721 = v22 + int32(1280)
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v721+v708))))
	v724 = int32(_a_F_vacuumLeafPage_4)
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v719)+4)))
	v727 = int32(_a_F_vacuumLeafPage_10)
	v729 = v723&v724 | v726&v727
	*(*uint16)(unsafe.Add(mBase, uint32(v719)+4)) = uint16(v729)
	v732 = v708 | v713
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v710+v732))))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v679+v734<<(uint(v713)%32))))
	v741 = v41 + v738&v717
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v721+v732))))
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v741)+4)))
	v751 = v745&v724 | v748&v727
	*(*uint16)(unsafe.Add(mBase, uint32(v741)+4)) = uint16(v751)
	v754 = v692 + v713
	v756 = v698 + v713
	if v756 != v675&int32(_a_F_vacuumLeafPage_9) {
		v692 = v754
		v698 = v756
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v464&int32(1) == int32(0) {
		goto L92
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v764 = v754
	goto L96
L101:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824)+118)))
	if v825 != int32(112) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v898 = int32(_a_F_vacuumLeafPage_8)
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[6])) = v900 - int32(1)
	goto L8
L103:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[7]))
	if v829 <= int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v832 != 0 {
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L17
	} else {
		goto L109
	}
L107:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v833 != 0 {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[8]))) = v836
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[9]))) = uint8(v838)
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_11), int32(16))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	v847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[2]))))
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_5), v847<<(uint(int32(1))%32))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	v854 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[3]))))
	F_XLogRegisterData(m, v22+int32(_a_F_vacuumLeafPage_6), v854<<(uint(int32(1))%32))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v861 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))))
	F_XLogRegisterData(m, v22+int32(3728), v861<<(uint(int32(1))%32))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L17
	} else {
		goto L113
	}
L113:
	;
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[4]))))
	F_XLogRegisterData(m, v22+int32(2912), v868<<(uint(int32(1))%32))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L17
	} else {
		goto L114
	}
L114:
	;
	v875 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))))
	F_XLogRegisterData(m, v22+int32(2096), v875<<(uint(int32(1))%32))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L17
	} else {
		goto L115
	}
L115:
	;
	v882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_vacuumLeafPage[5]))))
	F_XLogRegisterData(m, v22+int32(1280), v882<<(uint(int32(1))%32))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L17
	} else {
		goto L116
	}
L116:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L17
	} else {
		goto L117
	}
L117:
	;
	v893 = F_XLogInsert(m, int32(16), int32(96))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L17
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v893, int64(32))
	goto L102
L119:
	;
	if l2 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v950 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_12), v22+int32(16))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L17
	} else {
		goto L124
	}
L121:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[10]))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v934+(l2^int32(-1))<<(uint(int32(6))%32))+16))
	v949 = v940
	goto L120
L122:
	;
	goto L123
L123:
	;
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumLeafPage[11]))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v942+l2<<(uint(int32(6))%32)+int32(-64))+16))
	v949 = v948
	goto L120
L124:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(179), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L17
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v969 & int32(3)
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_15), v22)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(266), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L17
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errmsg_internal(m, int32(_a_F_vacuumLeafPage_16), int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_vacuumLeafPage_13), int32(326), int32(_a_F_vacuumLeafPage_14))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_delay_point(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 float64
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 float64
	_ = v107
	var v116 float64
	_ = v116
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 int32
	_ = v130
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v354 int32
	_ = v354
	v7 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v21 != 0 {
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
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L79
	}
L7:
	;
	m.G0 = v14 + int32(16)
	return
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[1])))
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[2]))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[2]))
	v32 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[3]))
	if base.B2i32(v31 == v32)|base.B2i32(v35 != int32(4)) == v32 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[4]))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[2])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v23 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L17:
	;
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[1])))
	if v50 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	goto L13
L21:
	;
	if base.F64_gt(v116, float64(0)) == int32(0) {
		goto L7
	} else {
		goto L29
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[5]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = int32(_a_F_vacuum_delay_point_0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6]))
	v61 = v58 + v60
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v61
	v63 = int32(_a_F_vacuum_delay_point_1)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6]))
	v68 = v65 + v67
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7])) = v68
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8]))
	if base.Ui32(v61) < base.Ui32(v71) {
		v95 = v7
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8]))
	if v102 < v104 {
		goto L7
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6])) = int32(0)
	v116 = v95
	goto L21
L26:
	;
	v73 = base.F64_convert_i32_s(v68)
	v74 = base.F64_convert_i32_s(v71)
	if base.F64_gt(v73, base.F64_mul(base.F64_div(v74, base.F64_convert_i32_s(v57)), float64(0.5))) == int32(0) {
		v95 = v7
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v83 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[4]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86 - v68
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[7])) = int32(0)
	v95 = base.F64_div(base.F64_mul(v83, v73), v74)
	goto L25
L28:
	;
	v107 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9]))
	v116 = base.F64_div(base.F64_mul(v107, base.F64_convert_i32_s(v102)), base.F64_convert_i32_s(v104))
	goto L21
L29:
	;
	v124 = *(*float64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[9]))
	v126 = base.F64_mul(v124, float64(4))
	if base.F64_gt(v116, v126) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v128 = v126
	goto L32
L31:
	;
	v128 = v116
	goto L32
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[10])))
	if v130 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v140 = v135 + v136*int64(1000000000)
	goto L35
L34:
	;
	v140 = int64(0)
	goto L35
L35:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(150994951)
	F_pg_usleep(m, base.I32_trunc_sat_f64_s(base.F64_mul(v128, float64(1000))))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(0)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[10])))
	if v155 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[12])))
	if v273 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L38:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v160 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v164 = v160 + v161*int64(1000000000)
	v165 = v164 - v140
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[19]))
	if int32(0) <= v167 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v170 = int32(_a_F_vacuum_delay_point_5)
	v172 = *(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20]))
	v173 = v172 + v165
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20])) = v173
	v176 = *(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[21]))
	if v164-v176 < int64(1000000000) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l0 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_pgstat_progress_parallel_incr_param(m, int32(10), v173)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[20])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_vacuum_delay_point[21])) = v164
	goto L37
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[22]))
	if v191 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[22]))
	if v232 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	goto L37
L48:
	;
	goto L47
L49:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[23])))
	if v195&int32(1) == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v200 = int32(_a_F_vacuum_delay_point_6)
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	v203 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v202 + v203
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v206 + v203
	v214 = v191 + int32(296)
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	*(*int64)(unsafe.Add(mBase, uint32(v214))) = v215 + v165
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v218 + v203
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v224 - v203
	goto L48
L51:
	;
	goto L37
L52:
	;
	goto L51
L53:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_delay_point[23])))
	if v236&int32(1) == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v241 = int32(_a_F_vacuum_delay_point_6)
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	v244 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v243 + v244
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v247 + v244
	v255 = v232 + int32(312)
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
	*(*int64)(unsafe.Add(mBase, uint32(v255))) = v256 + v165
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v259 + v244
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[24])) = v265 - v244
	goto L52
L55:
	;
	v276 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v281 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[6])) = v281
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[13]))
	if v284 == v281 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v276 == int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[0]))
	if v334 == int32(0) {
		goto L7
	} else {
		goto L77
	}
L61:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[14]))
	if v289 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L74
	}
L63:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[15]))
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[16]))
	if int32(0) < v294 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v316 = v289
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8])) = v316
	goto L60
L66:
	;
	v299 = v294
	goto L68
L67:
	;
	v299 = v296
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[8])) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v284)+32))
	if v301 == int32(0) {
		goto L60
	} else {
		goto L69
	}
L69:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_delay_point[17]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+uint32(_c_F_vacuum_delay_point[18])))
	if v306 <= int32(0) {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v309 = int32(1)
	v310 = base.I32_div_s(v299, v306)
	if v310 <= v309 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v313 = v309
	goto L73
L72:
	;
	v313 = v310
	goto L73
L73:
	;
	v316 = v313
	goto L65
L74:
	;
	F_errmsg_internal(m, int32(_a_F_vacuum_delay_point_2), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_vacuum_delay_point_3), int32(1754), int32(_a_F_vacuum_delay_point_4))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L7
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
