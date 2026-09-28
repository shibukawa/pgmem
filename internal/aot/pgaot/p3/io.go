package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IoWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
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
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int64
	_ = v512
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int64
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v682 int64
	_ = v682
	var v685 int32
	_ = v685
	var v688 int64
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int64
	_ = v700
	var v708 int64
	_ = v708
	var v709 int64
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
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
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int64
	_ = v760
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v967 int32
	_ = v967
	var v968 int64
	_ = v968
	var v971 int64
	_ = v971
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
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
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int64
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1053 int32
	_ = v1053
	var v1058 int64
	_ = v1058
	var v1061 int64
	_ = v1061
	var v1064 int64
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1073 int64
	_ = v1073
	var v1082 int64
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int64
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1144 int64
	_ = v1144
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int64
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1245 int32
	_ = v1245
	var v1246 int64
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(368)
	m.G0 = v18
	v24 = int32(-1)
	v25 = v3
	v26 = v3
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
	if v24 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1245 = int32(m.ExcTag)
	v1246 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1245 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L7:
	;
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v18)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v569 = v26
	goto L9
L9:
	;
	if v569 != 0 {
		goto L120
	} else {
		goto L121
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v56 = m.G0
	v58 = v56 - int32(32)
	m.G0 = v58
	v61 = int32(967)
	switch v61 {
	case 0, 2:
		goto L12
	default:
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v99 = m.G0
	v101 = v99 - int32(32)
	m.G0 = v101
	v104 = int32(974)
	switch v104 {
	case 0, 2:
		goto L22
	default:
		goto L23
	}
L12:
	;
	F_sigemptyset(m, v58+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = int32(268435456)
	switch v61 {
	case 0:
		goto L17
	default:
		goto L15
	case 2:
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[0])) = int32(965)
	goto L12
L14:
	;
	goto L19
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = int32(0)
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = int32(-2)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v90 = F___sigaction(m, int32(1), v58+int32(12), int32(0))
	mBase = m.M
	m.G0 = v58 + int32(32)
	goto L11
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v140 = int32(0)
	v142 = m.G0
	v144 = v142 - int32(32)
	m.G0 = v144
	switch v140 {
	case 0, 2:
		goto L32
	default:
		goto L33
	}
L22:
	;
	F_sigemptyset(m, v101+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v101)+24)) = int32(268435456)
	switch v104 {
	case 0:
		goto L27
	default:
		goto L25
	case 2:
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[1])) = int32(972)
	goto L22
L24:
	;
	goto L29
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(0)
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(-2)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v133 = F___sigaction(m, int32(2), v101+int32(12), int32(0))
	mBase = m.M
	m.G0 = v101 + int32(32)
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v183 = int32(0)
	v185 = m.G0
	v187 = v185 - int32(32)
	m.G0 = v187
	switch v183 {
	case 0, 2:
		goto L42
	default:
		goto L43
	}
L32:
	;
	F_sigemptyset(m, v144+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(268435456)
	switch v140 {
	case 0:
		goto L37
	default:
		goto L35
	case 2:
		goto L36
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[2])) = int32(-2)
	goto L32
L34:
	;
	goto L39
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = int32(0)
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = int32(-2)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v176 = F___sigaction(m, int32(15), v144+int32(12), int32(0))
	mBase = m.M
	m.G0 = v144 + int32(32)
	goto L31
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v226 = int32(0)
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	switch v226 {
	case 0, 2:
		goto L52
	default:
		goto L53
	}
L42:
	;
	F_sigemptyset(m, v187+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v187)+24)) = int32(268435456)
	switch v183 {
	case 0:
		goto L47
	default:
		goto L45
	case 2:
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[3])) = int32(-2)
	goto L42
L44:
	;
	goto L49
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = int32(0)
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = int32(-2)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v219 = F___sigaction(m, int32(14), v187+int32(12), int32(0))
	mBase = m.M
	m.G0 = v187 + int32(32)
	goto L41
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v271 = m.G0
	v273 = v271 - int32(32)
	m.G0 = v273
	v276 = int32(970)
	switch v276 {
	case 0, 2:
		goto L62
	default:
		goto L63
	}
L52:
	;
	F_sigemptyset(m, v230+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(268435456)
	switch v226 {
	case 0:
		goto L57
	default:
		goto L55
	case 2:
		goto L56
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[4])) = int32(-2)
	goto L52
L54:
	;
	goto L59
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(0)
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(-2)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v262 = F___sigaction(m, int32(13), v230+int32(12), int32(0))
	mBase = m.M
	m.G0 = v230 + int32(32)
	goto L51
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v314 = m.G0
	v316 = v314 - int32(32)
	m.G0 = v316
	v319 = int32(969)
	switch v319 {
	case 0, 2:
		goto L72
	default:
		goto L73
	}
L62:
	;
	F_sigemptyset(m, v273+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v273)+24)) = int32(268435456)
	switch v276 {
	case 0:
		goto L67
	default:
		goto L65
	case 2:
		goto L66
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[5])) = int32(968)
	goto L62
L64:
	;
	goto L69
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v273)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+12)) = int32(0)
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+12)) = int32(-2)
	goto L64
L69:
	;
	goto L70
L70:
	;
	v305 = F___sigaction(m, int32(10), v273+int32(12), int32(0))
	mBase = m.M
	m.G0 = v273 + int32(32)
	goto L61
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	v361 = F_LWLockAcquire(m, v357+int32(_a_F_IoWorkerMain_1), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L81
	}
L72:
	;
	F_sigemptyset(m, v316+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v316)+24)) = int32(268435456)
	switch v319 {
	case 0:
		goto L77
	default:
		goto L75
	case 2:
		goto L76
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8])) = int32(967)
	goto L72
L74:
	;
	goto L79
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = int32(_a_F_IoWorkerMain_0)
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = int32(0)
	goto L74
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = int32(-2)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v348 = F___sigaction(m, int32(12), v316+int32(12), int32(0))
	mBase = m.M
	m.G0 = v316 + int32(32)
	goto L71
L81:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v365 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v364)+16)))
	if v365 != int64(4294967295) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v364+v394<<(uint(int32(2))%32))+28)) = v399
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v402)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v402)+16)) = v403 | int64(1)<<(uint(base.I64_extend_i32_u(v394))%64)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v402)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+24)) = v409 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v415+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L90
	}
L83:
	;
	v372 = base.I32_wrap_i64(base.I64_ctz(v365 ^ int64(4294967295)))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6])) = v372
	v394 = v372
	goto L82
L84:
	;
	goto L85
L85:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	if v375 != int32(-1) {
		v394 = v375
		goto L82
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_errmsg_internal(m, int32(_a_F_IoWorkerMain_2), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_errfinish(m, int32(_a_F_IoWorkerMain_3), int32(590), int32(_a_F_IoWorkerMain_4))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	goto L3
L90:
	;
	if v403 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v427 = v423
	v437 = v403
	goto L94
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_on_shmem_exit(m, int32(1152), int64(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L114
	}
L94:
	;
	v439 = base.I64_ctz(v437)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v427+base.I32_wrap_i64(v439)<<(uint(int32(2))%32))+28))
	if v444 != int32(-1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[11]))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v455 = v449 + v444*int32(768) + int32(316)
	v456 = int32(0)
	v459 = base.AtomicRmwOr32(m, v456, int32(_a_F_IoWorkerMain_5), v456)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	if v460 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v509 = v427
	goto L98
L98:
	;
	v512 = v437 & base.I64_rotl(int64(-2), v439)
	if v512 != int64(0) {
		v427 = v509
		v437 = v512
		goto L94
	} else {
		goto L113
	}
L99:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v509 = v508
	goto L98
L100:
	;
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(1)
	v463 = int32(0)
	v466 = base.AtomicRmwOr32(m, v463, int32(_a_F_IoWorkerMain_5), v463)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v467 == v463 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	if v470 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12]))
	if v474 == v470 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v476 = m.G0
	v478 = v476 - int32(16)
	m.G0 = v478
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13]))
	if v481 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v504 = F_pgmem_kill(m, v470, int32(23))
	mBase = m.M
	goto L100
L107:
	;
	m.G0 = v478 + int32(16)
	goto L99
L108:
	;
	v484 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v478)+15)) = uint8(v484)
	goto L109
L109:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[14]))
	v492 = F_write(m, v488, v478+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v492 {
		goto L107
	} else {
		goto L111
	}
L110:
	;
	goto L107
L111:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[15]))
	if v496 == int32(27) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L95
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v537
	v540 = v18 + int32(48)
	v544 = F_pg_sprintf(m, v540, int32(_a_F_IoWorkerMain_6), v18+int32(32))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v547 = F_strlen(m, v540)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = int32(1153)
	v552 = int32(_a_F_IoWorkerMain_7)
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16])) = v18 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v553
	goto L116
L116:
	;
	v560 = v18 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v18 + int32(44)
	goto L119
L117:
	;
	v569 = int32(0)
	goto L9
L119:
	;
	goto L117
L120:
	;
	v581 = int32(_a_F_IoWorkerMain_8)
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17])) = v583 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_EmitErrorReport(m)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[18])) = v18 + int32(208)
	F_pgmem_sigprocmask(m, int32(_a_F_IoWorkerMain_9), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L6
	} else {
		goto L130
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_LWLockReleaseAll(m)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	if v596 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[15])) = v599
	v601 = int32(_a_F_IoWorkerMain_10)
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[19])) = v603 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	F_pgaio_io_process_completion(m, v608, int32(0)-v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v25
	F_proc_exit(m, int32(1))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L6
	} else {
		goto L129
	}
L128:
	;
	v614 = int32(_a_F_IoWorkerMain_10)
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[19])) = v616 - int32(1)
	goto L127
L129:
	;
	goto L3
L130:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[20]))
	if v634 != 0 {
		v1210 = v25
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16])) = v1224
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1210
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L6
	} else {
		goto L259
	}
L132:
	;
	v635 = int32(0)
	v640 = v635
	v641 = v25
	v647 = v635
	v649 = v635
	v652 = int64(0)
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	v660 = F_LWLockAcquire(m, v656+int32(_a_F_IoWorkerMain_11), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L6
	} else {
		goto L135
	}
L134:
	;
	v1210 = v1174
	goto L131
L135:
	;
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[21]))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+8))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v663)+4))
	if v664 == v665 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[22]))
	if v1186 != 0 {
		goto L249
	} else {
		goto L250
	}
L137:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v967)+8))
	v971 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v967)+8)) = v968 | int64(1)<<(uint(v971)%64)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v977 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v977+int32(_a_F_IoWorkerMain_11))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L6
	} else {
		goto L205
	}
L138:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v663+v664<<(uint(int32(2))%32))+12))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v672 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v663)+8)) = (v671 - v672) & (v664 + v672)
	if v670 == int32(-1) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v681)+8))
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	v688 = v682 & base.I64_rotl(int64(-2), base.I64_extend_i32_u(v685))
	*(*int64)(unsafe.Add(mBase, uint32(v681)+8)) = v688
	if v647 < v640 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[23]))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v859 = v855 + v670<<(uint(int32(7))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v859
	v864 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L6
	} else {
		goto L179
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v844+int32(_a_F_IoWorkerMain_11))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L6
	} else {
		goto L178
	}
L142:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v663)+8))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v663)+4))
	if base.Ui32(v692) < base.Ui32(v691) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v696 = v694 + v692
	goto L145
L144:
	;
	v696 = v692
	goto L145
L145:
	;
	v697 = v696 - v691
	if v697 <= int32(0) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v700 = int64(-1)
	if v685 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v708 = v700
	goto L149
L148:
	;
	v708 = v700 << (uint(base.I64_extend_i32_u(v685+int32(1))) % 64)
	goto L149
L149:
	;
	v709 = v708 & v688
	if v709 == int64(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v714+int32(_a_F_IoWorkerMain_11))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v760 = base.I64_ctz(v709)
	*(*int64)(unsafe.Add(mBase, uint32(v681)+8)) = v688 & base.I64_rotl(int64(-2), v760)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v766+int32(_a_F_IoWorkerMain_11))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L6
	} else {
		goto L162
	}
L153:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+24))
	if v697 <= v721 {
		goto L140
	} else {
		goto L154
	}
L154:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[24]))
	if v724 <= v721 {
		goto L140
	} else {
		goto L155
	}
L155:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	if v726 != 0 {
		goto L140
	} else {
		goto L156
	}
L156:
	;
	v727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v720))) = uint8(v727)
	v729 = int32(0)
	v732 = base.AtomicRmwOr32(m, v729, int32(_a_F_IoWorkerMain_12), v729)
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+1)))
	if v735 != 0 {
		goto L140
	} else {
		goto L157
	}
L157:
	;
	v736 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v734)+1)) = uint8(v736)
	v738 = int32(0)
	v741 = base.AtomicRmwOr32(m, v738, int32(_a_F_IoWorkerMain_12), v738)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v745 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_IoWorkerMain[25])))
	if v745 == v736 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L140
L159:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v749+int32(24)))) = int32(1)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[27]))
	v758 = F_pgmem_kill(m, v756, int32(10))
	mBase = m.M
	goto L161
L160:
	;
	goto L161
L161:
	;
	goto L158
L162:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v772+base.I32_wrap_i64(v760)<<(uint(int32(2))%32))+28))
	if v777 == int32(-1) {
		goto L140
	} else {
		goto L163
	}
L163:
	;
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[11]))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v788 = v782 + v777*int32(768) + int32(316)
	v789 = int32(0)
	v792 = base.AtomicRmwOr32(m, v789, int32(_a_F_IoWorkerMain_5), v789)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	if v793 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L140
L165:
	;
	goto L164
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = int32(1)
	v796 = int32(0)
	v799 = base.AtomicRmwOr32(m, v796, int32(_a_F_IoWorkerMain_5), v796)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v800 == v796 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v788)+12))
	if v803 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12]))
	if v807 == v803 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v809 = m.G0
	v811 = v809 - int32(16)
	m.G0 = v811
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13]))
	if v814 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v837 = F_pgmem_kill(m, v803, int32(23))
	mBase = m.M
	goto L165
L172:
	;
	m.G0 = v811 + int32(16)
	goto L164
L173:
	;
	v817 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v811)+15)) = uint8(v817)
	goto L174
L174:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[14]))
	v825 = F_write(m, v821, v811+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v825 {
		goto L172
	} else {
		goto L176
	}
L175:
	;
	goto L172
L176:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[15]))
	if v829 == int32(27) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	goto L140
L179:
	;
	if v864 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	F_errhidestmt(m)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L6
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v924 = int32(2)
	v925 = base.I32_div_s(v640, v924)
	v926 = int32(_a_F_IoWorkerMain_8)
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17])) = v928 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = int32(44)
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+1)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v935<<(uint(v924)%32))+uint32(_c_F_IoWorkerMain[28])))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v938)))
	m.T0[v939].(func(*base.Module, int32))(m, v859)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L6
	} else {
		goto L197
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	F_errhidecontext(m)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[23]))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+24))
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+2)))
	if base.Ui32(v880) <= base.Ui32(int32(2)) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+1)))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v889<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[28])))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v892)+8))
	goto L190
L187:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v880<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[29])))
	v887 = v885
	goto L189
L188:
	;
	v887 = int32(0)
	goto L189
L189:
	;
	goto L186
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	if base.Ui32(v895) <= base.Ui32(int32(7)) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v905 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v893
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v887
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v859 - v875) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_IoWorkerMain_13), v18)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L6
	} else {
		goto L195
	}
L192:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v895<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[30])))
	v902 = v900
	goto L194
L193:
	;
	v902 = int32(0)
	goto L194
L194:
	;
	goto L191
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	F_errfinish(m, int32(_a_F_IoWorkerMain_3), int32(865), int32(_a_F_IoWorkerMain_14))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	goto L182
L197:
	;
	v942 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v942
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v942
	F_pgaio_io_perform_synchronously(m, v859)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v950 = v647 + int32(1)
	v952 = base.B2i32(v950 == int32(4))
	if v950 == int32(4) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v953 = v925
	goto L201
L200:
	;
	v953 = v640
	goto L201
L201:
	;
	if v950 == int32(4) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v955 = int32(2)
	goto L204
L203:
	;
	v955 = v950
	goto L204
L204:
	;
	v956 = int32(_a_F_IoWorkerMain_8)
	v958 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17])) = v958 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = int32(0)
	v1173 = v953
	v1174 = v641
	v1179 = v955
	v1181 = v649
	v1183 = int64(0)
	goto L136
L205:
	;
	v983 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983))))
	if v984 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v987 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v983))) = uint8(v987)
	v992 = base.AtomicRmwOr32(m, v987, int32(_a_F_IoWorkerMain_12), v987)
	goto L208
L207:
	;
	goto L208
L208:
	;
	v993 = int32(-1)
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[31]))
	if v995 == v993 {
		v1087 = v641
		v1088 = v993
		v1091 = v649
		v1093 = v652
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1087
	v1097 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[32]))
	v1100 = F_WaitLatch(m, v1097, int32(41), v1088, int32(83886086))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L6
	} else {
		goto L237
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
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
	goto L211
L211:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[33]))
	if v1018 < v1020 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1087 = v641
	v1088 = v993
	v1091 = v649
	v1093 = int64(0)
	goto L209
L213:
	;
	goto L214
L214:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	v1031 = F_LWLockAcquire(m, v1027+int32(_a_F_IoWorkerMain_1), int32(1))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v1035 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v1038 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v1038+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	if v1044 != int32(63)-base.I32_wrap_i64(base.I64_clz(v1035)) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1087 = v641
	v1088 = v993
	v1091 = v649
	v1093 = int64(0)
	goto L209
L218:
	;
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v641
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[31]))
	v1058 = int64(0)
	if v1024 == v649 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1061 = v652
	goto L222
L221:
	;
	v1061 = v1058
	goto L222
L222:
	;
	if v652 == int64(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1064 = v1058
	goto L225
L224:
	;
	v1064 = v1061
	goto L225
L225:
	;
	v1066 = base.B2i32(v1064 == int64(0))
	if v1064 == int64(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1067 = base.I64_extend_i32_s(v1053)*int64(1000) + v1016
	goto L228
L227:
	;
	v1067 = v1064
	goto L228
L228:
	;
	if v1067 <= v1016 {
		v1085 = int32(0)
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v1064 == int64(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	goto L229
L231:
	;
	v1073 = v1067 - v1016
	if base.B2i32(int64(0) < v1016)^base.B2i32(v1073 < v1067)|base.B2i32(int64(2147483646000) < v1073) != 0 {
		v1085 = int32(2147483647)
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1082 = base.I64_div_s(v1073+int64(999), int64(1000))
	v1085 = base.I32_wrap_i64(v1082)
	goto L230
L233:
	;
	v1086 = v1053
	goto L235
L234:
	;
	v1086 = v649
	goto L235
L235:
	;
	v1087 = v1085
	v1088 = v1085
	v1091 = v1086
	v1093 = v1067
	goto L209
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1087
	v1166 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[32]))
	v1167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1166))) = v1167
	v1172 = base.AtomicRmwOr32(m, v1167, int32(_a_F_IoWorkerMain_5), v1167)
	goto L248
L237:
	;
	if v1100 == int32(8) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[33]))
	if v1105 < v1107 {
		v1161 = v640
		v1162 = v647
		goto L236
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1155 = v640 + int32(1)
	if v1155 != int32(4) {
		v1161 = v1155
		v1162 = v647
		goto L236
	} else {
		goto L247
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1087
	v1111 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	v1115 = F_LWLockAcquire(m, v1111+int32(_a_F_IoWorkerMain_1), int32(1))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L6
	} else {
		goto L242
	}
L242:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9]))
	v1119 = *(*int64)(unsafe.Add(mBase, uint32(v1118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1087
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	F_LWLockRelease(m, v1122+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L6
	} else {
		goto L243
	}
L243:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	if v1128 != int32(63)-base.I32_wrap_i64(base.I64_clz(v1119)) {
		v1161 = v640
		v1162 = v647
		goto L236
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1087
	v1138 = m.G0
	v1139 = int32(16)
	v1140 = v1138 - v1139
	m.G0 = v1140
	F_gettimeofday(m, v1140)
	mBase = m.M
	v1143 = *(*int64)(unsafe.Add(mBase, uint32(v1140)))
	v1144 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1140)+8)))
	m.G0 = v1140 + v1139
	goto L245
L245:
	;
	if v1144+v1143*int64(1000000)-int64(946684800000000) < v1093 {
		v1161 = v640
		v1162 = v647
		goto L236
	} else {
		goto L246
	}
L246:
	;
	v1210 = v1087
	goto L131
L247:
	;
	v1158 = int32(2)
	v1160 = base.I32_div_s(v647, v1158)
	v1161 = v1158
	v1162 = v1160
	goto L236
L248:
	;
	v1173 = v1161
	v1174 = v1087
	v1179 = v1162
	v1181 = v1091
	v1183 = v1093
	goto L136
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1174
	F_ProcessInterrupts(m)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L6
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[34]))
	if v1191 != 0 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+364)) = v1174
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[34])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L6
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[20]))
	if v1205 == int32(0) {
		v640 = v1173
		v641 = v1174
		v647 = v1179
		v649 = v1181
		v652 = v1183
		goto L133
	} else {
		goto L258
	}
L256:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6]))
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[24]))
	if v1202 <= v1200 {
		v1210 = v1174
		goto L131
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	goto L134
L259:
	;
	goto L5
L260:
	;
	v1250 = int32(v1246)
	m.G0 = v18
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+4))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1253)))
	if v18+int32(44) == v1256 {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	m.ExcPending = 1
	goto L269
L262:
	;
	if v1260 != 0 {
		goto L266
	} else {
		goto L267
	}
L263:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+4))
	v1260 = v1258
	goto L265
L264:
	;
	v1260 = int32(0)
	goto L265
L265:
	;
	goto L262
L266:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v18)+364))
	v24 = v1260
	v25 = v1261
	v26 = v1252
	goto L1
L267:
	;
	goto L268
L268:
	;
	F___wasm_longjmp(m, v1253, v1252)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	return
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
