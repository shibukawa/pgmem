package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BackgroundWriterMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int64
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v671 float32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v681 float32
	_ = v681
	var v686 float32
	_ = v686
	var v688 float32
	_ = v688
	var v689 float32
	_ = v689
	var v690 int32
	_ = v690
	var v692 float32
	_ = v692
	var v698 float32
	_ = v698
	var v701 float64
	_ = v701
	var v704 int32
	_ = v704
	var v705 float32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int64
	_ = v781
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v813 int32
	_ = v813
	var v815 int64
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v828 float32
	_ = v828
	var v839 int32
	_ = v839
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int64
	_ = v906
	var v907 int32
	_ = v907
	var v908 int64
	_ = v908
	var v911 int64
	_ = v911
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int64
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int64
	_ = v930
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int64
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1016 int64
	_ = v1016
	var v1018 int64
	_ = v1018
	var v1024 int64
	_ = v1024
	var v1025 int64
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int64
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1091 int32
	_ = v1091
	var v1092 int64
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(_a_F_BackgroundWriterMain_0)
	m.G0 = v22
	v25 = int32(-1)
	v27 = v3
	v28 = v3
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v25 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v1091 = int32(m.ExcTag)
	v1092 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1091 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v440 = v27
	v441 = v28
	goto L8
L8:
	;
	if v441 != 0 {
		goto L97
	} else {
		goto L98
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v54 = m.G0
	v56 = v54 - int32(32)
	m.G0 = v56
	v59 = int32(967)
	switch v59 {
	case 0, 2:
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v95 = int32(0)
	v97 = m.G0
	v99 = v97 - int32(32)
	m.G0 = v99
	switch v95 {
	case 0, 2:
		goto L21
	default:
		goto L22
	}
L11:
	;
	F_sigemptyset(m, v56+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = int32(268435456)
	switch v59 {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[1])) = int32(965)
	goto L11
L13:
	;
	goto L18
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(-2)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v88 = F___sigaction(m, int32(1), v56+int32(12), int32(0))
	mBase = m.M
	m.G0 = v56 + int32(32)
	goto L10
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v140 = m.G0
	v142 = v140 - int32(32)
	m.G0 = v142
	v145 = int32(969)
	switch v145 {
	case 0, 2:
		goto L31
	default:
		goto L32
	}
L21:
	;
	F_sigemptyset(m, v99+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = int32(268435456)
	switch v95 {
	case 0:
		goto L26
	default:
		goto L24
	case 2:
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[2])) = int32(-2)
	goto L21
L23:
	;
	goto L28
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(0)
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(-2)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v131 = F___sigaction(m, int32(2), v99+int32(12), int32(0))
	mBase = m.M
	m.G0 = v99 + int32(32)
	goto L20
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v181 = int32(0)
	v183 = m.G0
	v185 = v183 - int32(32)
	m.G0 = v185
	switch v181 {
	case 0, 2:
		goto L41
	default:
		goto L42
	}
L31:
	;
	F_sigemptyset(m, v142+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = int32(268435456)
	switch v145 {
	case 0:
		goto L36
	default:
		goto L34
	case 2:
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[3])) = int32(967)
	goto L31
L33:
	;
	goto L38
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = int32(0)
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = int32(-2)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v174 = F___sigaction(m, int32(15), v142+int32(12), int32(0))
	mBase = m.M
	m.G0 = v142 + int32(32)
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v224 = int32(0)
	v226 = m.G0
	v228 = v226 - int32(32)
	m.G0 = v228
	switch v224 {
	case 0, 2:
		goto L51
	default:
		goto L52
	}
L41:
	;
	F_sigemptyset(m, v185+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v185)+24)) = int32(268435456)
	switch v181 {
	case 0:
		goto L46
	default:
		goto L44
	case 2:
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[4])) = int32(-2)
	goto L41
L43:
	;
	goto L48
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = int32(0)
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = int32(-2)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v217 = F___sigaction(m, int32(14), v185+int32(12), int32(0))
	mBase = m.M
	m.G0 = v185 + int32(32)
	goto L40
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v269 = m.G0
	v271 = v269 - int32(32)
	m.G0 = v271
	v274 = int32(970)
	switch v274 {
	case 0, 2:
		goto L61
	default:
		goto L62
	}
L51:
	;
	F_sigemptyset(m, v228+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = int32(268435456)
	switch v224 {
	case 0:
		goto L56
	default:
		goto L54
	case 2:
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[5])) = int32(-2)
	goto L51
L53:
	;
	goto L58
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = int32(0)
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = int32(-2)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v260 = F___sigaction(m, int32(13), v228+int32(12), int32(0))
	mBase = m.M
	m.G0 = v228 + int32(32)
	goto L50
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v310 = int32(0)
	v312 = m.G0
	v314 = v312 - int32(32)
	m.G0 = v314
	switch v310 {
	case 0, 2:
		goto L71
	default:
		goto L72
	}
L61:
	;
	F_sigemptyset(m, v271+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(268435456)
	switch v274 {
	case 0:
		goto L66
	default:
		goto L64
	case 2:
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[6])) = int32(968)
	goto L61
L63:
	;
	goto L68
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(0)
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(-2)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v303 = F___sigaction(m, int32(10), v271+int32(12), int32(0))
	mBase = m.M
	m.G0 = v271 + int32(32)
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v355 = m.G0
	v357 = v355 - int32(32)
	m.G0 = v357
	v359 = int32(2)
	switch v359 {
	case 0, 2:
		goto L81
	default:
		goto L82
	}
L71:
	;
	F_sigemptyset(m, v314+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = int32(268435456)
	switch v310 {
	case 0:
		goto L76
	default:
		goto L74
	case 2:
		goto L75
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[7])) = int32(-2)
	goto L71
L73:
	;
	goto L78
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(_a_F_BackgroundWriterMain_1)
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(0)
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(-2)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v346 = F___sigaction(m, int32(12), v314+int32(12), int32(0))
	mBase = m.M
	m.G0 = v314 + int32(32)
	goto L70
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v398 = m.G0
	v399 = int32(16)
	v400 = v398 - v399
	m.G0 = v400
	F_gettimeofday(m, v400)
	mBase = m.M
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v400)))
	v404 = int64(*(*int32)(unsafe.Add(mBase, uint32(v400)+8)))
	m.G0 = v400 + v399
	goto L90
L81:
	;
	F_sigemptyset(m, v357+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v357)+24)) = int32(268435456)
	switch v359 {
	case 0:
		goto L86
	default:
		goto L84
	case 2:
		goto L85
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[8])) = int32(0)
	goto L81
L83:
	;
	goto L87
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+12)) = int32(_a_F_BackgroundWriterMain_1)
	v382 = int32(268435461)
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+12)) = int32(0)
	v382 = int32(268435457)
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+12)) = int32(-2)
	v382 = int32(268435457)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+24)) = v382
	goto L89
L89:
	;
	v389 = F___sigaction(m, int32(17), v357+int32(12), int32(0))
	mBase = m.M
	m.G0 = v357 + int32(32)
	goto L80
L90:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[9])) = v404 + v403*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v27
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[10]))
	v422 = F_AllocSetContextCreateInternal(m, v417, int32(_a_F_BackgroundWriterMain_2), int32(0), int32(_a_F_BackgroundWriterMain_3), int32(_a_F_BackgroundWriterMain_4))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[11])) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v422
	v427 = v22 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v427)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = int32(_a_F_BackgroundWriterMain_5)
	goto L92
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[12]))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[13]))) = v22 + int32(4)
	goto L96
L94:
	;
	v440 = v422
	v441 = int32(0)
	goto L8
L96:
	;
	goto L94
L97:
	;
	v442 = int32(_a_F_BackgroundWriterMain_6)
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[14])) = v444 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[15])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_EmitErrorReport(m)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[16])) = v22 + int32(_a_F_BackgroundWriterMain_7)
	F_pgmem_sigprocmask(m, int32(_a_F_BackgroundWriterMain_8), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L112
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_LWLockReleaseAll(m)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_UnlockBuffers(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_smgrdestroyall(m)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[11])) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_FlushErrorState(m)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_MemoryContextReset(m, v440)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v492 = v22 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v492)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v492))) = int32(_a_F_BackgroundWriterMain_5)
	goto L111
L111:
	;
	v497 = int32(_a_F_BackgroundWriterMain_6)
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[14])) = v499 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_pg_usleep(m, int32(_a_F_BackgroundWriterMain_9))
	mBase = m.M
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = int32(0)
	goto L99
L112:
	;
	v523 = int32(0)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[18]))
	v542 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v542
	v547 = base.AtomicRmwOr32(m, v542, int32(_a_F_BackgroundWriterMain_10), v542)
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v552 = int32(0)
	v553 = m.G0
	v555 = v553 - int32(16)
	m.G0 = v555
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[19]))
	v561 = base.AtomicRmwXchg32(m, v558, v552, int32(1))
	if v561 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_s_lock(m, v558, int32(_a_F_BackgroundWriterMain_11))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[19]))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[20]))
	v571 = v555 + int32(12)
	if v571 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L119
L121:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v572
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[20]))
	v576 = base.I32_div_u_s(v567, v575)
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v572 + v576
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[19]))
	v582 = v580
	goto L123
L122:
	;
	v582 = v566
	goto L123
L123:
	;
	v583 = int32(8)
	v586 = v555 + v583
	if v586 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v589 = base.AtomicRmwXchg32(m, v582, int32(12), int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = v589
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[19]))
	v593 = v592
	goto L126
L125:
	;
	v593 = v582
	goto L126
L126:
	;
	v594 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v593))), uint32(v594))
	v597 = base.I32_rem_u_s(v567, v569)
	v598 = int32(_a_F_BackgroundWriterMain_12)
	v600 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[21]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[21])) = v600 + base.I64_extend_i32_u(v601)
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[22]))
	if v606 <= v594 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	m.G0 = v555 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[23]))
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[24]))
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[25]))
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[21]))
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[26]))
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[27]))
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[28]))
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[29]))
	if v868|(v870|(v872|(v874|(v876|(v878|(v880|v882)))))) != 0 {
		goto L177
	} else {
		goto L178
	}
L128:
	;
	v610 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[30])) = uint8(v610)
	v862 = int32(1)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[30])))
	if v614 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[31])) = v658
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[32])) = v597
	v668 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[30])) = uint8(v668)
	v671 = *(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[33]))
	v672 = int32(0)
	if v601 != 0 {
		goto L143
	} else {
		goto L144
	}
L132:
	;
	v657 = v652
	v658 = v653
	v659 = v597
	v661 = v656
	v662 = v652
	goto L131
L133:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[32]))
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[20]))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[31]))
	v625 = v597 - v616 + v619*(v620-v622)
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[34]))
	if int32(0) < v627-v620 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35])) = v597
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[34])) = v648
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[20]))
	v652 = v651
	v653 = v648
	v656 = v552
	goto L132
L136:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35]))
	v657 = v619
	v658 = v620
	v659 = v632
	v661 = v625
	v662 = v597 - v632
	goto L131
L137:
	;
	goto L138
L138:
	;
	if v620 != v627 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[34])) = v620
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35])) = v597
	v652 = v619
	v653 = v620
	v656 = v625
	goto L132
L140:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35]))
	if v636 < v597 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v657 = v619
	v658 = v620
	v659 = v636
	v661 = v625
	v662 = v619 + v597 - v636
	goto L131
L142:
	;
	v690 = int32(_a_F_BackgroundWriterMain_13)
	v692 = *(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[36]))
	if base.F32_le(v692, v689) != 0 {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	v675 = base.B2i32(v672 < v661)
	goto L145
L144:
	;
	v675 = v672
	goto L145
L145:
	;
	if v675 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v688 = v671
	v689 = base.F32_convert_i32_u(v601)
	goto L142
L147:
	;
	goto L148
L148:
	;
	v681 = base.F32_convert_i32_u(v601)
	v686 = base.F32_add(v671, base.F32_mul(base.F32_sub(base.F32_div(base.F32_convert_i32_u(v661), v681), v671), float32(0.0625)))
	*(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[33])) = v686
	v688 = v686
	v689 = v681
	goto L142
L149:
	;
	v698 = v689
	goto L151
L150:
	;
	v698 = base.F32_add(v692, base.F32_mul(base.F32_sub(v689, v692), float32(0.0625)))
	goto L151
L151:
	;
	v701 = *(*float64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[37]))
	v704 = base.I32_trunc_sat_f64_s(base.F64_mul(v701, base.F64_promote_f32(v698)))
	if v704 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v705 = v698
	goto L154
L153:
	;
	v705 = float32(0)
	goto L154
L154:
	;
	*(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[36])) = v705
	v707 = int32(0)
	v711 = base.I32_trunc_sat_f32_s(base.F32_div(base.F32_convert_i32_s(v657-v662), v688))
	if v662 <= v707 {
		v797 = v662
		v798 = v711
		v802 = v707
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v813 = int32(_a_F_BackgroundWriterMain_14)
	v815 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[29]))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[29])) = v815 + base.I64_extend_i32_s(v802)
	v820 = v662 - v797
	v821 = int32(0)
	if base.B2i32(v798 == v711)|base.B2i32(v820 <= v821) == v821 {
		goto L174
	} else {
		goto L175
	}
L156:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[38]))
	v722 = base.I32_trunc_sat_f32_s(base.F32_div(base.F32_convert_i32_s(v657), base.F32_div(float32(120000), base.F32_convert_i32_s(v717)))) + v711
	if v704 < v722 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v724 = v722
	goto L159
L158:
	;
	v724 = v704
	goto L159
L159:
	;
	if v724 <= v711 {
		v797 = v662
		v798 = v711
		v802 = v707
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v726 = v662
	v730 = v711
	v732 = v659
	v734 = v707
	goto L161
L161:
	;
	v746 = F_SyncOneBuffer(m, v732, int32(1), v22+v583)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L163
	}
L162:
	;
	v797 = v769
	v798 = v790
	v802 = v789
	goto L155
L163:
	;
	v748 = int32(_a_F_BackgroundWriterMain_15)
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35]))
	v752 = v750 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35])) = v752
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[20]))
	if v755 <= v752 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v757 = int32(_a_F_BackgroundWriterMain_16)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[34]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[34])) = v759 + int32(1)
	v764 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[35])) = v764
	v767 = v764
	goto L166
L165:
	;
	v767 = v752
	goto L166
L166:
	;
	v768 = int32(1)
	v769 = v726 - v768
	if v746&v768 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if base.Ui32(v726) < base.Ui32(int32(2)) {
		v797 = v769
		v798 = v790
		v802 = v789
		goto L155
	} else {
		goto L172
	}
L168:
	;
	v772 = int32(1)
	v773 = v730 + v772
	v775 = v734 + v772
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[22]))
	if v775 < v777 {
		v789 = v775
		v790 = v773
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v789 = v734
	v790 = int32(base.Ui32(v746)>>(uint(int32(1))%32)) + v730
	goto L167
L171:
	;
	v779 = int32(_a_F_BackgroundWriterMain_17)
	v781 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[27]))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[27])) = v781 + int64(1)
	v797 = v769
	v798 = v773
	v802 = v775
	goto L155
L172:
	;
	if v790 < v724 {
		v726 = v769
		v730 = v790
		v732 = v767
		v734 = v789
		goto L161
	} else {
		goto L173
	}
L173:
	;
	goto L162
L174:
	;
	v826 = int32(_a_F_BackgroundWriterMain_18)
	v828 = *(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[33]))
	*(*float32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[33])) = base.F32_add(v828, base.F32_mul(base.F32_sub(base.F32_div(base.F32_convert_i32_u(v820), base.F32_convert_i32_u(v798-v711)), v828), float32(0.0625)))
	goto L176
L175:
	;
	goto L176
L176:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	v862 = base.B2i32(v662|v839 == int32(0))
	goto L127
L177:
	;
	v890 = int32(_a_F_BackgroundWriterMain_19)
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[39]))
	v893 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[39])) = v892 + v893
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[40]))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+592))
	*(*int32)(unsafe.Add(mBase, uint32(v897)+592)) = v898 + v893
	v902 = int32(0)
	v904 = int32(_a_F_BackgroundWriterMain_20)
	v905 = base.AtomicRmwOr32(m, v902, v904, v902)
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v897)+600))
	v907 = int32(_a_F_BackgroundWriterMain_14)
	v908 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[29]))
	*(*int64)(unsafe.Add(mBase, uint32(v897)+600)) = v906 + v908
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v897)+608))
	v912 = int32(_a_F_BackgroundWriterMain_17)
	v913 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v897)+608)) = v911 + v913
	v916 = *(*int64)(unsafe.Add(mBase, uint32(v897)+616))
	v917 = int32(_a_F_BackgroundWriterMain_12)
	v918 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[21]))
	*(*int64)(unsafe.Add(mBase, uint32(v897)+616)) = v916 + v918
	v924 = base.AtomicRmwOr32(m, v902, v904, v902)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v897)+592))
	*(*int32)(unsafe.Add(mBase, uint32(v897)+592)) = v925 + v893
	v930 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[29])) = v930
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[27])) = v930
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[21])) = v930
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[24])) = v930
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[39]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[39])) = v943 - v893
	F_pgstat_flush_io(m, v902)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L5
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L5
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v957 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[41]))
	v960 = base.AtomicRmwXchg32(m, v957, int32(4), int32(1))
	if v960 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	F_s_lock(m, v957+int32(4), int32(_a_F_BackgroundWriterMain_21))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L5
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[41]))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v967)+12))
	v969 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v967)+4)), uint32(v969))
	v972 = int32(_a_F_BackgroundWriterMain_22)
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[42]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[42])) = v968
	if v973 != v968 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L184
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_smgrdestroyall(m)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[43]))
	if v981 <= int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L188
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1042 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[18]))
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[38]))
	v1047 = F_WaitLatch(m, v1042, int32(41), v1045, int32(83886083))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L5
	} else {
		goto L202
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v987 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[44])))
	if v987 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v997 != 0 {
		goto L190
	} else {
		goto L196
	}
L193:
	;
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[45]))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)+308))
	v995 = base.B2i32(v993 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[44])) = uint8(v995)
	v997 = v995
	goto L195
L194:
	;
	v997 = int32(0)
	goto L195
L195:
	;
	goto L192
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1002 = m.G0
	v1003 = int32(16)
	v1004 = v1002 - v1003
	m.G0 = v1004
	F_gettimeofday(m, v1004)
	mBase = m.M
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(v1004)))
	v1008 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1004)+8)))
	m.G0 = v1004 + v1003
	v1016 = v1008 + v1007*int64(1000000) - int64(946684800000000)
	goto L197
L197:
	;
	v1018 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[9]))
	if v1016 < v1018+int64(15000000) {
		goto L190
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1024 = *(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[46]))
	v1025 = F_GetLastImportantRecPtr(m)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	if base.Ui64(v1025) < base.Ui64(v1024) {
		goto L190
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1029 = F_LogStandbySnapshot(m)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[9])) = v1016
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[46])) = v1029
	goto L190
L202:
	;
	if base.B2i32(v862&v523 == int32(0))|base.B2i32(v1047 != int32(8)) != 0 {
		v523 = v862
		goto L113
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[47]))
	F_StrategyNotifyBgWriter(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[18]))
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWriterMain[38]))
	v1066 = F_WaitLatch(m, v1059, int32(41), v1062*int32(50), int32(83886082))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0]))) = v440
	F_StrategyNotifyBgWriter(m, int32(-1))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v523 = v862
	goto L113
L207:
	;
	v1096 = int32(v1092)
	m.G0 = v22
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	if v22+int32(4) == v1102 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	m.ExcPending = 1
	goto L216
L209:
	;
	if v1106 != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+4))
	v1106 = v1104
	goto L212
L211:
	;
	v1106 = int32(0)
	goto L212
L212:
	;
	goto L209
L213:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_BackgroundWriterMain[0])))
	v25 = v1106
	v27 = v1107
	v28 = v1098
	goto L1
L214:
	;
	goto L215
L215:
	;
	F___wasm_longjmp(m, v1099, v1098)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	return
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
