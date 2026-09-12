package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncReadBuffers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v587 int32
	_ = v587
	var v589 int64
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int64
	_ = v607
	var v619 int32
	_ = v619
	var v625 int64
	_ = v625
	var v631 int64
	_ = v631
	var v633 int32
	_ = v633
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v772 int32
	_ = v772
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int64
	_ = v801
	var v803 int64
	_ = v803
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v854 int64
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v871 int64
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v921 int64
	_ = v921
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v972 int64
	_ = v972
	var v973 int64
	_ = v973
	var v977 int64
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1007 int64
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1019 int64
	_ = v1019
	var v1020 int64
	_ = v1020
	var v1024 int64
	_ = v1024
	var v1050 int32
	_ = v1050
	var v1052 int64
	_ = v1052
	var v1054 int64
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int64
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1077 int64
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1094 int32
	_ = v1094
	var v1100 int64
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1116 int32
	_ = v1116
	var v1122 int64
	_ = v1122
	var v1128 int64
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1141 int64
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int64
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1151 int64
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	v27 = m.G0
	v29 = v27 - int32(512)
	m.G0 = v29
	v31 = int32(3)
	v32 = int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = int32(base.Ui32(v33)>>(uint(v31)%32)) & v32
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+34)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v44 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = v38<<(uint(int32(2))%32) + v41
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _consts[728])))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[729])))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	F_pgstat_prepare_report_checksum_failure(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v55 = v31
	v56 = v32
	v57 = v37 | int32(2)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = F_IOContextForStrategy(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v55 = v51
	v56 = int32(0)
	v57 = v37
	goto L1
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v70 = l0 + int32(48)
	v71 = F_pgaio_io_acquire_nb(m, v68, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v71 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v461 = v71
	goto L11
L11:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v487 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+22)))
	if v488 != 0 {
		goto L89
	} else {
		goto L90
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v79 = m.G0
	v81 = v79 - int32(80)
	m.G0 = v81
	v83 = F_pgaio_io_acquire_nb(m, v78, v70)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L16
	}
L13:
	;
	v461 = v395
	goto L11
L14:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L83
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L79
	}
L16:
	;
	if v83 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	goto L20
L18:
	;
	v395 = v83
	goto L19
L19:
	;
	m.G0 = v81 + int32(80)
	goto L13
L20:
	;
	v117 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	v395 = v389
	goto L19
L22:
	;
	if v117 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_errhidestmt(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v147 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	if int32(0) < v147 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	F_errhidecontext(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+160))
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+64)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v81)+68)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v81)+72)) = v125
	F_errmsg_internal(m, int32(165265), v81-int32(-64))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(474548), int32(768), int32(392292))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v389 = F_pgaio_io_acquire_nb(m, v78, v70)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L77
	}
L31:
	;
	v150 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v156 = v150
	v157 = v145
	v162 = v147
	v164 = v150
	v165 = v152
	goto L34
L32:
	;
	v213 = v145
	goto L33
L33:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+22)))
	if v236 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v182 = int32(7)
	v187 = v180 + v181<<(uint(v182)%32) + v156<<(uint(v182)%32)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v188 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if int32(0) < v203 {
		goto L30
	} else {
		goto L41
	}
L36:
	;
	F_pgaio_io_reclaim(m, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v201 = v157
	v202 = v162
	v203 = v164
	v204 = v165
	goto L38
L38:
	;
	v206 = v156 + int32(1)
	if v206 < v202 {
		v156 = v206
		v157 = v201
		v162 = v202
		v164 = v203
		v165 = v204
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v198 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v201 = v200
	v202 = v196
	v203 = v164 + int32(1)
	v204 = v198
	goto L38
L40:
	;
	goto L35
L41:
	;
	v213 = v201
	goto L33
L42:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v241 = v213
	goto L44
L44:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	if v242 != 0 {
		goto L30
	} else {
		goto L46
	}
L45:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v241 = v240
	goto L44
L46:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)+160))
	if v243 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)+156))
	v248 = v246 - int32(24)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if base.Ui32(int32(7)) < base.Ui32(v249) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	if v360 == int32(0) {
		goto L14
	} else {
		goto L76
	}
L49:
	;
	if int32(1)<<(uint(v249)%32)&int32(48) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_pgaio_io_reclaim(m, v248)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L75
	}
L51:
	;
	if v249 == int32(6) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v246)+24))
	v286 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L58
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = (v248 - v266) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(454514), v81+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(474548), int32(840), int32(392292))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	if v286 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errhidestmt(m)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_pgaio_io_wait(m, v248, v283)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L74
	}
L62:
	;
	F_errhidecontext(m)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v292 = int32(0)
	v294 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+24))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
	if base.Ui32(v300) <= base.Ui32(int32(2)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v309<<(uint(int32(2))%32))+uint32(_consts[718])))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	goto L68
L65:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300<<(uint(int32(2))%32))+uint32(_consts[717])))
	v308 = v307
	goto L67
L66:
	;
	v308 = v292
	goto L67
L67:
	;
	goto L64
L68:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if base.Ui32(v316) <= base.Ui32(int32(7)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(int32(2))%32))+uint32(_consts[719])))
	v324 = v323
	goto L71
L70:
	;
	v324 = v292
	goto L71
L71:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v81+int32(48)))) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = (v248 - v295) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+36)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v81)+40)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v81)+44)) = v324
	F_errmsg_internal(m, int32(97698), v81+int32(32))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(474548), int32(847), int32(392292))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	goto L48
L75:
	;
	goto L48
L76:
	;
	goto L30
L77:
	;
	if v389 == int32(0) {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	goto L21
L79:
	;
	F_errmsg_internal(m, int32(165228), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v431)+160))
	v434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v432
	F_errdetail_internal(m, int32(165290), v81)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(474548), int32(818), int32(392292))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errmsg_internal(m, int32(338627), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(474548), int32(881), int32(392292))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	m.G0 = v29 + int32(512)
	return v545
L87:
	;
	if v545 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L88:
	;
	v545 = v543
	goto L87
L89:
	;
	if v485 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if v485 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L100
	}
L93:
	;
	v491 = int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v500 = F_StartLocalBufferIO(m, v493+(v485^int32(-1))<<(uint(int32(6))%32), v491)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L5
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v504 = int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v514 = F_StartBufferIO(m, v506+v485<<(uint(int32(6))%32)+int32(-64), v504, v504)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	if v500 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v543 = v491
	goto L88
L98:
	;
	if v514 != 0 {
		v543 = v504
		goto L88
	} else {
		goto L99
	}
L99:
	;
	goto L92
L100:
	;
	goto L91
L101:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v530 = F_StartLocalBufferIO(m, v523+(v485^int32(-1))<<(uint(int32(6))%32), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L5
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v541 = F_StartBufferIO(m, v533+v485<<(uint(int32(6))%32)+int32(-64), int32(1), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L105
	}
L104:
	;
	v545 = v530
	goto L87
L105:
	;
	v543 = v541
	goto L88
L106:
	;
	v548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v549 = int32(1)
	v550 = v548 + v549
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)) = uint16(v550)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v549
	v555 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+16))
	if v556 == v461 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v653 = v33 | v62
	if v60 != 0 {
		goto L131
	} else {
		goto L132
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L117
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+16)) = int32(0)
	F_pgaio_io_reclaim(m, v461)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L5
	} else {
		goto L114
	}
L113:
	;
	goto L109
L114:
	;
	F_errmsg_internal(m, int32(337309), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(474548), int32(258), int32(345421))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	if v44 == int32(116) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v593 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v581 = int32(4342152)
	v583 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v583 + int64(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v587 = int32(4342120)
	v589 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v589 + int64(1)
	goto L118
L122:
	;
	v619 = v56*int32(320) + v55<<(uint(int32(6))%32)
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v619)+uint32(_consts[723])))
	*(*int64)(unsafe.Add(mBase, uint32(v619)+uint32(_consts[723]))) = v625 + int64(1)
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v619)+uint32(_consts[724])))
	*(*int64)(unsafe.Add(mBase, uint32(v619)+uint32(_consts[724]))) = v631
	v633 = int32(1)
	F_pgstat_count_backend_io_op(m, v56, v55, int32(2), v633, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v633)
	*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v633)
	goto L129
L123:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v593)+272))
	if v596 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+268)))
	if v599 != int32(1) {
		goto L122
	} else {
		goto L127
	}
L125:
	;
	v606 = v596
	goto L126
L126:
	;
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v606)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v606)+120)) = v607 + int64(1)
	goto L122
L127:
	;
	F_pgstat_assoc_relation(m, v593)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+272))
	v606 = v605
	goto L126
L129:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, _consts[409])))
	if v643 != int32(1) {
		goto L86
	} else {
		goto L130
	}
L130:
	;
	v646 = int32(4438576)
	v648 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	v650 = *(*int32)(unsafe.Add(mBase, _consts[727]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v648 + v650
	goto L86
L131:
	;
	v656 = v653 | int32(4)
	goto L133
L132:
	;
	v656 = v653
	goto L133
L133:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v657 < int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v675
	v677 = int32(1)
	v679 = v38 + v677
	v680 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v680 <= v679 {
		v772 = v677
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661+(v657^int32(-1))<<(uint(int32(2))%32))))
	v675 = v667
	goto L134
L136:
	;
	goto L137
L137:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v675 = v669 + v657<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L138:
	;
	v793 = l0 + int32(36)
	v795 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v793))) = (v461 - v796) >> (uint(int32(7)) % 32)
	v801 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v461)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v793)+4)) = uint32(v801)
	v803 = *(*int64)(unsafe.Add(mBase, uint32(v461)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v793)+8)) = uint32(v803)
	goto L155
L139:
	;
	v688 = v677
	v689 = v679
	goto L140
L140:
	;
	v710 = v41 + v689<<(uint(int32(2))%32)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	if v711 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v772 = v761
	goto L138
L142:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	if v740 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L143:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v722 = F_StartLocalBufferIO(m, v715+(v711^int32(-1))<<(uint(int32(6))%32), int32(1))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L5
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v731 = int32(1)
	v733 = F_StartBufferIO(m, v725+v711<<(uint(int32(6))%32)+int32(-64), v731, v731)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L5
	} else {
		goto L148
	}
L146:
	;
	if v722 != 0 {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	v772 = v688
	goto L138
L148:
	;
	if v733 == int32(0) {
		v772 = v688
		goto L138
	} else {
		goto L149
	}
L149:
	;
	goto L142
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v688<<(uint(int32(2))%32)))) = v758
	v760 = int32(1)
	v761 = v688 + v760
	v763 = v689 + v760
	v764 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v763 < v764 {
		v688 = v761
		v689 = v763
		goto L140
	} else {
		goto L154
	}
L151:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v744+(v740^int32(-1))<<(uint(int32(2))%32))))
	v758 = v750
	goto L150
L152:
	;
	goto L153
L153:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v758 = v752 + v740<<(uint(int32(13))%32) + int32(-8192)
	goto L150
L154:
	;
	goto L141
L155:
	;
	v805 = int32(0)
	v808 = v772 & int32(255)
	if v808 == v805 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v461)+13)) = uint8(v808)
	if v44 == int32(116) {
		goto L165
	} else {
		goto L166
	}
L157:
	;
	if v808 != int32(1) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v819 = v805
	v820 = v805
	goto L161
L159:
	;
	v882 = v805
	goto L160
L160:
	;
	if v808&int32(1) == int32(0) {
		goto L156
	} else {
		goto L164
	}
L161:
	;
	v841 = int32(4354468)
	v842 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)+16))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v461)+76))
	v845 = int32(3)
	v851 = int32(2)
	v854 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58+v819<<(uint(v851)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v843+v844<<(uint(v845)%32)+v819<<(uint(v845)%32)))) = v854
	v857 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)+16))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v461)+76))
	v864 = v819 | int32(1)
	v871 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58+v864<<(uint(v851)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v858+v859<<(uint(v845)%32)+v864<<(uint(v845)%32)))) = v871
	v874 = v819 + v851
	v876 = v820 + v851
	if v876 != v808&int32(254) {
		v819 = v874
		v820 = v876
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v882 = v874
	goto L160
L163:
	;
	goto L162
L164:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+16))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v461)+76))
	v912 = int32(3)
	v921 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58+v882<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v910+v911<<(uint(v912)%32)+v882<<(uint(v912)%32)))) = v921
	goto L156
L165:
	;
	v954 = int32(3)
	goto L167
L166:
	;
	v954 = int32(2)
	goto L167
L167:
	;
	F_pgaio_io_register_callbacks(m, v461, v954, v656&int32(255))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+3)))
	v960 = v959 | v57
	*(*uint8)(unsafe.Add(mBase, uint32(v461)+3)) = uint8(v960)
	goto L169
L169:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, _consts[731])))
	v966 = m.G0
	v968 = v966 - int32(16)
	m.G0 = v968
	if v963 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v982 = int32(4438508)
	v984 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v984 + int32(1)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v981)+36))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v989*int32(80))+uint32(_consts[732])))
	m.T0[v994].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v461, v981, v42, v38+v43, v29, v772)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L5
	} else {
		goto L174
	}
L171:
	;
	F___clock_gettime(m, int32(1), v968)
	mBase = m.M
	v972 = int64(*(*int32)(unsafe.Add(mBase, uint32(v968)+8)))
	v973 = *(*int64)(unsafe.Add(mBase, uint32(v968)))
	v977 = v972 + v973*int64(1000000000)
	goto L173
L172:
	;
	v977 = int64(0)
	goto L173
L173:
	;
	m.G0 = v968 + int32(16)
	goto L170
L174:
	;
	v997 = int32(4438508)
	v999 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v1000 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v999 - v1000
	v1007 = base.I64_extend_i32_s(v772 << (uint(int32(13)) % 32))
	v1011 = m.G0
	v1013 = v1011 - int32(16)
	m.G0 = v1013
	if v977 != int64(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v1141 = base.I64_extend_i32_s(v772)
	if v44 == int32(116) {
		goto L193
	} else {
		goto L194
	}
L176:
	;
	F___clock_gettime(m, int32(1), v1013)
	mBase = m.M
	v1019 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1013)+8)))
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(v1013)))
	v1024 = v1019 + (v1020*int64(1000000000) - v977)
	if v56 == int32(2) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	v1116 = v56*int32(320) + v55<<(uint(int32(6))%32)
	v1122 = *(*int64)(unsafe.Add(mBase, uint32(v1116)+uint32(_consts[733])))
	*(*int64)(unsafe.Add(mBase, uint32(v1116)+uint32(_consts[733]))) = v1122 + base.I64_extend_i32_u(v1000)
	v1128 = *(*int64)(unsafe.Add(mBase, uint32(v1116)+uint32(_consts[734])))
	*(*int64)(unsafe.Add(mBase, uint32(v1116)+uint32(_consts[734]))) = v1128 + v1007
	F_pgstat_count_backend_io_op(m, v56, v55, int32(6), v1000, v1007)
	mBase = m.M
	v1133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v1133)
	*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v1133)
	m.G0 = v1013 + int32(16)
	goto L175
L179:
	;
	v1071 = v56*int32(320) + v55<<(uint(int32(6))%32)
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v1071)+uint32(_consts[735])))
	*(*int64)(unsafe.Add(mBase, uint32(v1071)+uint32(_consts[735]))) = v1077 + v1024
	v1081 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if base.Ui32(int32(16)) < base.Ui32(v1081) {
		goto L189
	} else {
		goto L190
	}
L180:
	;
	goto L182
L182:
	;
	goto L183
L183:
	;
	goto L186
L186:
	;
	v1050 = int32(4423576)
	v1052 = *(*int64)(unsafe.Add(mBase, _consts[736]))
	v1054 = base.I64_div_s(v1024, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[736])) = v1052 + v1054
	switch v56 {
	case 0:
		goto L188
	case 1:
		goto L187
	default:
		goto L179
	}
L187:
	;
	v1062 = int32(4342216)
	v1064 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, _consts[378])) = v1064 + v1024
	goto L179
L188:
	;
	v1057 = int32(4342200)
	v1059 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	*(*int64)(unsafe.Add(mBase, _consts[376])) = v1059 + v1024
	goto L179
L189:
	;
	goto L178
L190:
	;
	if int32(1)<<(uint(v1081)%32)&int32(115186) == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v1094 = v56*int32(320) + v55<<(uint(int32(6))%32)
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1094)+uint32(_consts[737])))
	*(*int64)(unsafe.Add(mBase, uint32(v1094)+uint32(_consts[737]))) = v1100 + v1024
	v1104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v1104)
	*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v1104)
	goto L189
L192:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[409])))
	if v1155 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v1144 = int32(4342160)
	v1146 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	*(*int64)(unsafe.Add(mBase, _consts[371])) = v1146 + v1141
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1149 = int32(4342128)
	v1151 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	*(*int64)(unsafe.Add(mBase, _consts[368])) = v1151 + v1141
	goto L192
L196:
	;
	v1158 = int32(4438576)
	v1160 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	v1162 = *(*int32)(unsafe.Add(mBase, _consts[739]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v1160 + v1162*v772
	goto L198
L197:
	;
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v772
	goto L86
}
func F_asyncQueueAdvanceTail(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v12 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v16 = F_LWLockAcquire(m, v12+int32(6016), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[24]))
		v23 = F_LWLockAcquire(m, v19+int32(3456), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[157]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
			if v30 != int32(-1) {
				v38 = v30
				v39 = v28
				v40 = v27
				v45 = v29
				for {
					v48 = v38 << (uint(int32(5)) % 32)
					v49 = v26 + int32(72) + v48
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
					if v45 < v50 {
						v57 = v39
						v58 = v40
						v60 = v45
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
						if v45 == v50 {
							if v39 < v52 {
								v57 = v39
								v58 = v40
								v60 = v45
							} else {
								v55 = v45
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
								v57 = v52
								v58 = v56
								v60 = v55
							}
						} else {
							v55 = v50
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
							v57 = v52
							v58 = v56
							v60 = v55
						}
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v48+(v26-int32(-64)))))
					if v63 != int32(-1) {
						v38 = v63
						v39 = v57
						v40 = v58
						v45 = v60
						continue
					} else {
						break
					}
					break
				}
				v68 = v57
				v69 = v58
				v74 = v60
			} else {
				v68 = v28
				v69 = v27
				v74 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v69
			*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v68
			*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v74
			v79 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
			v81 = *(*int32)(unsafe.Add(mBase, _consts[24]))
			F_LWLockRelease(m, v81+int32(3456))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				v87 = base.I64_rem_s(v74, int64(32))
				if v79 < v74-v87 {
					F_SimpleLruTruncate(m, int32(4340728), v74)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, _consts[24]))
						v98 = F_LWLockAcquire(m, v94+int32(3456), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _consts[157]))
							*(*int64)(unsafe.Add(mBase, uint32(v101)+32)) = v74
							v104 = *(*int32)(unsafe.Add(mBase, _consts[24]))
							F_LWLockRelease(m, v104+int32(3456))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, _consts[24]))
								F_LWLockRelease(m, v110+int32(6016))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, _consts[24]))
					F_LWLockRelease(m, v110+int32(6016))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_asyncQueuePagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	return base.B2i32(l0 < l1)
}
func F_mark_async_capable_plan(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v5 = l1
	goto L1
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v7 != int32(301) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	goto L2
L4:
	;
	switch v7 - int32(287) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v42 == int32(331) {
		goto L3
	} else {
		goto L20
	}
L7:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)) = uint8(v38)
	return v38
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == int32(331) {
		goto L3
	} else {
		goto L16
	}
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(331) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v15 = F_trivial_subqueryscan(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v15 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	v23 = F_mark_async_capable_plan(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v23 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+168))
	if v30 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v33 = m.T0[v30].(func(*base.Module, int32) int32)(m, v5)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	v5 = v45
	goto L1
}
