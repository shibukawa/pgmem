package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackgroundWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int64
	_ = v367
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int64
	_ = v828
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int64
	_ = v841
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(176)
	m.G0 = v10
	v15 = v3
	v16 = int32(-1)
	v17 = v3
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
	if v16 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v840 = int32(m.ExcTag)
	v841 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v840 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L7:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v522 = v15
	v524 = v17
	goto L9
L9:
	;
	if v524 != 0 {
		goto L127
	} else {
		goto L128
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v15
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v15
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[0]))
	v44 = F_MemoryContextAlloc(m, v42, int32(1472))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v15
	F_errmsg_internal(m, int32(_a_F_BackgroundWorkerMain_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v15
	F_errfinish(m, int32(_a_F_BackgroundWorkerMain_1), int32(739), int32(_a_F_BackgroundWorkerMain_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	base.MemoryCopy(m, v44, l0, int32(1472))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[1]))
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	F_MemoryContextDelete(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[2])) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	if v44 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[1])) = int32(0)
	goto L19
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[3]))
	if int32(0) < v65 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[4]))
	v63 = F_GetBackendTypeDesc(m, v62)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	F_pg_usleep(m, v65*int32(_a_F_BackgroundWorkerMain_3))
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v74 = int32(2)
	v78 = v72 & v74
	if v78 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = int32(966)
	goto L30
L29:
	;
	v79 = int32(-2)
	goto L30
L30:
	;
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	v87 = v79 + int32(2)
	switch v87 {
	case 0, 2:
		goto L32
	default:
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	if v78 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	F_sigemptyset(m, v84+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = int32(268435456)
	switch v87 {
	case 0:
		goto L37
	default:
		goto L35
	case 2:
		goto L36
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[5])) = v79
	goto L32
L34:
	;
	goto L39
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(0)
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(-2)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v116 = F___sigaction(m, v74, v84+int32(12), int32(0))
	mBase = m.M
	m.G0 = v84 + int32(32)
	goto L31
L41:
	;
	v124 = int32(968)
	goto L43
L42:
	;
	v124 = int32(-2)
	goto L43
L43:
	;
	v127 = m.G0
	v129 = v127 - int32(32)
	m.G0 = v129
	v132 = v124 + int32(2)
	switch v132 {
	case 0, 2:
		goto L45
	default:
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	if v78 != 0 {
		goto L54
	} else {
		goto L55
	}
L45:
	;
	F_sigemptyset(m, v129+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v129)+24)) = int32(268435456)
	switch v132 {
	case 0:
		goto L50
	default:
		goto L48
	case 2:
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[6])) = v124
	goto L45
L47:
	;
	goto L52
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = int32(0)
	goto L47
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = int32(-2)
	goto L47
L52:
	;
	goto L53
L53:
	;
	v161 = F___sigaction(m, int32(10), v129+int32(12), int32(0))
	mBase = m.M
	m.G0 = v129 + int32(32)
	goto L44
L54:
	;
	v169 = int32(970)
	goto L56
L55:
	;
	v169 = int32(-2)
	goto L56
L56:
	;
	v172 = m.G0
	v174 = v172 - int32(32)
	m.G0 = v174
	v177 = v169 + int32(2)
	switch v177 {
	case 0, 2:
		goto L58
	default:
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v215 = m.G0
	v217 = v215 - int32(32)
	m.G0 = v217
	v220 = int32(974)
	switch v220 {
	case 0, 2:
		goto L68
	default:
		goto L69
	}
L58:
	;
	F_sigemptyset(m, v174+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = int32(268435456)
	switch v177 {
	case 0:
		goto L63
	default:
		goto L61
	case 2:
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[7])) = v169
	goto L58
L60:
	;
	goto L65
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = int32(0)
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = int32(-2)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v206 = F___sigaction(m, int32(8), v174+int32(12), int32(0))
	mBase = m.M
	m.G0 = v174 + int32(32)
	goto L57
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v256 = int32(0)
	v258 = m.G0
	v260 = v258 - int32(32)
	m.G0 = v260
	switch v256 {
	case 0, 2:
		goto L78
	default:
		goto L79
	}
L68:
	;
	F_sigemptyset(m, v217+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = int32(268435456)
	switch v220 {
	case 0:
		goto L73
	default:
		goto L71
	case 2:
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[8])) = int32(972)
	goto L68
L70:
	;
	goto L75
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = int32(0)
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = int32(-2)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v249 = F___sigaction(m, int32(15), v217+int32(12), int32(0))
	mBase = m.M
	m.G0 = v217 + int32(32)
	goto L67
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[9])) = v297
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[10])) = v297
	v307 = v297
	goto L88
L78:
	;
	F_sigemptyset(m, v260+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v260)+24)) = int32(268435456)
	switch v256 {
	case 0:
		goto L83
	default:
		goto L81
	case 2:
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[11])) = int32(-2)
	goto L78
L80:
	;
	goto L85
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = int32(0)
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = int32(-2)
	goto L80
L85:
	;
	goto L86
L86:
	;
	v292 = F___sigaction(m, int32(1), v260+int32(12), int32(0))
	mBase = m.M
	m.G0 = v260 + int32(32)
	goto L77
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v388 = int32(0)
	v390 = m.G0
	v392 = v390 - int32(32)
	m.G0 = v392
	switch v388 {
	case 0, 2:
		goto L94
	default:
		goto L95
	}
L88:
	;
	v309 = int32(40)
	v310 = v307 * v309
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[13]))) = v307
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[14]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[15]))) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[16]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[17]))) = v311
	*(*uint8)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v311)
	v325 = v307 | int32(1)
	v327 = v325 * v309
	*(*uint8)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[13]))) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[14]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[15]))) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[16]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[17]))) = v311
	*(*uint8)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v311)
	v342 = v307 | int32(2)
	v344 = v342 * v309
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[13]))) = v342
	*(*int64)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[14]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[15]))) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[16]))) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[17]))) = v311
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v311)
	if v307 != int32(20) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v380 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[19])) = uint8(v380)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L87
L90:
	;
	v361 = v307 | int32(3)
	v363 = v361 * int32(40)
	v364 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v364)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[13]))) = v361
	v367 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[14]))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[15]))) = v364
	*(*int64)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[16]))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[17]))) = v364
	*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v364)
	v307 = v307 + int32(4)
	goto L88
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v431 = int32(0)
	v433 = m.G0
	v435 = v433 - int32(32)
	m.G0 = v435
	switch v431 {
	case 0, 2:
		goto L104
	default:
		goto L105
	}
L94:
	;
	F_sigemptyset(m, v392+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v392)+24)) = int32(268435456)
	switch v388 {
	case 0:
		goto L99
	default:
		goto L97
	case 2:
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[20])) = int32(-2)
	goto L94
L96:
	;
	goto L101
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = int32(0)
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = int32(-2)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v424 = F___sigaction(m, int32(13), v392+int32(12), int32(0))
	mBase = m.M
	m.G0 = v392 + int32(32)
	goto L93
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v44
	v476 = m.G0
	v478 = v476 - int32(32)
	m.G0 = v478
	v480 = int32(2)
	switch v480 {
	case 0, 2:
		goto L114
	default:
		goto L115
	}
L104:
	;
	F_sigemptyset(m, v435+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v435)+24)) = int32(268435456)
	switch v431 {
	case 0:
		goto L109
	default:
		goto L107
	case 2:
		goto L108
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[21])) = int32(-2)
	goto L104
L106:
	;
	goto L111
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v435)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+12)) = int32(0)
	goto L106
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+12)) = int32(-2)
	goto L106
L111:
	;
	goto L112
L112:
	;
	v467 = F___sigaction(m, int32(12), v435+int32(12), int32(0))
	mBase = m.M
	m.G0 = v435 + int32(32)
	goto L103
L113:
	;
	goto L123
L114:
	;
	F_sigemptyset(m, v478+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = int32(268435456)
	switch v480 {
	case 0:
		goto L119
	default:
		goto L117
	case 2:
		goto L118
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[22])) = int32(0)
	goto L114
L116:
	;
	goto L120
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(_a_F_BackgroundWorkerMain_4)
	v503 = int32(268435461)
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(0)
	v503 = int32(268435457)
	goto L116
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(-2)
	v503 = int32(268435457)
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = v503
	goto L122
L122:
	;
	v510 = F___sigaction(m, int32(17), v478+int32(12), int32(0))
	mBase = m.M
	m.G0 = v478 + int32(32)
	goto L113
L123:
	;
	v515 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v515)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v10 + int32(12)
	goto L126
L124:
	;
	v522 = v44
	v524 = int32(0)
	goto L9
L126:
	;
	goto L124
L127:
	;
	v525 = int32(_a_F_BackgroundWorkerMain_5)
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[23])) = v527 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[24])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L6
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[25])) = v10 + int32(16)
	F_InitProcess(m)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L133
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	F_EmitErrorReport(m)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	F_proc_exit(m, int32(1))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	goto L3
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	F_BaseInit(m)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	v556 = v522 + int32(1228)
	v557 = m.G0
	v559 = v557 - int32(16)
	m.G0 = v559
	v562 = v522 + int32(204)
	v563 = int32(_a_F_BackgroundWorkerMain_6)
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[26])))
	if base.B2i32(v566 == int32(0))|base.B2i32(v566 != v569) != 0 {
		v587 = v566
		v588 = v569
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v522)+1328))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	m.T0[v811].(func(*base.Module, int64))(m, v828)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L6
	} else {
		goto L220
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L6
	} else {
		goto L217
	}
L137:
	;
	m.G0 = v559 + int32(16)
	goto L135
L138:
	;
	if v587-v588 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	goto L138
L140:
	;
	v572 = v562
	v573 = v563
	goto L141
L141:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+1)))
	if v577 == int32(0) {
		v587 = v577
		v588 = v576
		goto L139
	} else {
		goto L143
	}
L142:
	;
	v587 = v577
	v588 = v576
	goto L139
L143:
	;
	v580 = int32(1)
	if v577 == v576 {
		v572 = v572 + v580
		v573 = v573 + v580
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v592 = int32(_a_F_BackgroundWorkerMain_7)
	v595 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[27])))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v595 == int32(0))|base.B2i32(v595 != v598) != 0 {
		v616 = v595
		v617 = v598
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	v809 = F_load_external_function(m, v562, v556, int32(1), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L216
	}
L148:
	;
	if v616-v617 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	v601 = v592
	v602 = v556
	goto L151
L151:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+1)))
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	if v606 == int32(0) {
		v616 = v606
		v617 = v605
		goto L149
	} else {
		goto L153
	}
L152:
	;
	v616 = v606
	v617 = v605
	goto L149
L153:
	;
	v609 = int32(1)
	if v606 == v605 {
		v601 = v601 + v609
		v602 = v602 + v609
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[28]))
	v811 = v622
	goto L137
L156:
	;
	goto L157
L157:
	;
	v623 = int32(_a_F_BackgroundWorkerMain_8)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[29])))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v626 == int32(0))|base.B2i32(v626 != v629) != 0 {
		v647 = v626
		v648 = v629
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v647-v648 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	goto L158
L160:
	;
	v632 = v623
	v633 = v556
	goto L161
L161:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+1)))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+1)))
	if v637 == int32(0) {
		v647 = v637
		v648 = v636
		goto L159
	} else {
		goto L163
	}
L162:
	;
	v647 = v637
	v648 = v636
	goto L159
L163:
	;
	v640 = int32(1)
	if v637 == v636 {
		v632 = v632 + v640
		v633 = v633 + v640
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v653 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[30]))
	v811 = v653
	goto L137
L166:
	;
	goto L167
L167:
	;
	v654 = int32(_a_F_BackgroundWorkerMain_9)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[31])))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v657 == int32(0))|base.B2i32(v657 != v660) != 0 {
		v678 = v657
		v679 = v660
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v678-v679 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	goto L168
L170:
	;
	v663 = v654
	v664 = v556
	goto L171
L171:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+1)))
	if v668 == int32(0) {
		v678 = v668
		v679 = v667
		goto L169
	} else {
		goto L173
	}
L172:
	;
	v678 = v668
	v679 = v667
	goto L169
L173:
	;
	v671 = int32(1)
	if v668 == v667 {
		v663 = v663 + v671
		v664 = v664 + v671
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[32]))
	v811 = v684
	goto L137
L176:
	;
	goto L177
L177:
	;
	v685 = int32(_a_F_BackgroundWorkerMain_10)
	v688 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[33])))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v688 == int32(0))|base.B2i32(v688 != v691) != 0 {
		v709 = v688
		v710 = v691
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v709-v710 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	goto L178
L180:
	;
	v694 = v685
	v695 = v556
	goto L181
L181:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695)+1)))
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	if v699 == int32(0) {
		v709 = v699
		v710 = v698
		goto L179
	} else {
		goto L183
	}
L182:
	;
	v709 = v699
	v710 = v698
	goto L179
L183:
	;
	v702 = int32(1)
	if v699 == v698 {
		v694 = v694 + v702
		v695 = v695 + v702
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[34]))
	v811 = v715
	goto L137
L186:
	;
	goto L187
L187:
	;
	v716 = int32(_a_F_BackgroundWorkerMain_11)
	v719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[35])))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v719 == int32(0))|base.B2i32(v719 != v722) != 0 {
		v740 = v719
		v741 = v722
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if v740-v741 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L189:
	;
	goto L188
L190:
	;
	v725 = v716
	v726 = v556
	goto L191
L191:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+1)))
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+1)))
	if v730 == int32(0) {
		v740 = v730
		v741 = v729
		goto L189
	} else {
		goto L193
	}
L192:
	;
	v740 = v730
	v741 = v729
	goto L189
L193:
	;
	v733 = int32(1)
	if v730 == v729 {
		v725 = v725 + v733
		v726 = v726 + v733
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[36]))
	v811 = v746
	goto L137
L196:
	;
	goto L197
L197:
	;
	v747 = int32(_a_F_BackgroundWorkerMain_12)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[37])))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v750 == int32(0))|base.B2i32(v750 != v753) != 0 {
		v771 = v750
		v772 = v753
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v771-v772 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L199:
	;
	goto L198
L200:
	;
	v756 = v747
	v757 = v556
	goto L201
L201:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757)+1)))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756)+1)))
	if v761 == int32(0) {
		v771 = v761
		v772 = v760
		goto L199
	} else {
		goto L203
	}
L202:
	;
	v771 = v761
	v772 = v760
	goto L199
L203:
	;
	v764 = int32(1)
	if v761 == v760 {
		v756 = v756 + v764
		v757 = v757 + v764
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[38]))
	v811 = v777
	goto L137
L206:
	;
	goto L207
L207:
	;
	v778 = int32(_a_F_BackgroundWorkerMain_13)
	v781 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[39])))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if base.B2i32(v781 == int32(0))|base.B2i32(v781 != v784) != 0 {
		v802 = v781
		v803 = v784
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v802-v803 != 0 {
		goto L136
	} else {
		goto L215
	}
L209:
	;
	goto L208
L210:
	;
	v787 = v778
	v788 = v556
	goto L211
L211:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+1)))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787)+1)))
	if v792 == int32(0) {
		v802 = v792
		v803 = v791
		goto L209
	} else {
		goto L213
	}
L212:
	;
	v802 = v792
	v803 = v791
	goto L209
L213:
	;
	v795 = int32(1)
	if v792 == v791 {
		v787 = v787 + v795
		v788 = v788 + v795
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[40]))
	v811 = v806
	goto L137
L216:
	;
	v811 = v809
	goto L137
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v556
	F_errmsg_internal(m, int32(_a_F_BackgroundWorkerMain_14), v559)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_BackgroundWorkerMain_1), int32(1368), int32(_a_F_BackgroundWorkerMain_15))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L6
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v522
	F_proc_exit(m, int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L6
	} else {
		goto L221
	}
L221:
	;
	goto L5
L222:
	;
	v845 = int32(v841)
	m.G0 = v10
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	if v10+int32(12) == v851 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	m.ExcPending = 1
	goto L231
L224:
	;
	if v855 != 0 {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v855 = v853
	goto L227
L226:
	;
	v855 = int32(0)
	goto L227
L227:
	;
	goto L224
L228:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	v15 = v856
	v16 = v855
	v17 = v847
	goto L1
L229:
	;
	goto L230
L230:
	;
	F___wasm_longjmp(m, v848, v847)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	return
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BackgroundWorkerUnblockSignals(m *base.Module) {
	var v4 int32
	_ = v4
	F_pgmem_sigprocmask(m, int32(_a_F_BackgroundWorkerUnblockSignals_0), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
	v13 = F_LWLockAcquire(m, v9+int32(_a_F_GetBackgroundWorkerPid_0), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = v6 + v7*int32(1488)
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
		if v17 == v21 {
			v24 = v20 + int32(16)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v25 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
				F_LWLockRelease(m, v37+int32(_a_F_GetBackgroundWorkerPid_0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					switch v35 + int32(1) {
					case 0:
						return int32(1)
					case 1:
						v49 = int32(2)
						return v49
					default:
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
						v49 = int32(0)
						return v49
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
				F_LWLockRelease(m, v28+int32(_a_F_GetBackgroundWorkerPid_0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return int32(2)
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
			F_LWLockRelease(m, v28+int32(_a_F_GetBackgroundWorkerPid_0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return int32(2)
			}
		}
	}
}
