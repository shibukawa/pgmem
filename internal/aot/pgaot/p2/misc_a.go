package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbsoluteConfigLocation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v4 = m.G0
	v6 = v4 - int32(1024)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 == int32(47) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(1024)
	return v192
L2:
	;
	v11 = F_pstrdup(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	v192 = v11
	goto L1
L7:
	;
	F_join_path_components(m, v6, v186, l0)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L64
	}
L8:
	;
	goto L14
L9:
	;
	goto L10
L10:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_AbsoluteConfigLocation[0]))
	v186 = v185
	goto L7
L11:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v137 != 0 {
		goto L43
	} else {
		goto L44
	}
L12:
	;
	v131 = F_strlen(m, v120)
	mBase = m.M
	goto L11
L14:
	;
	goto L15
L15:
	;
	v21 = int32(1023)
	if (v6^l1)&int32(3) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
	goto L12
L17:
	;
	v105 = v100
	v106 = v101
	v107 = v102
	goto L38
L18:
	;
	if v95 == int32(0) {
		v120 = v93
		v121 = v94
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v93 = l1
	v94 = v6
	v95 = v21
	goto L18
L20:
	;
	goto L21
L21:
	;
	v25 = int32(0)
	if base.B2i32(l1&int32(3) == v25)|int32(0) == v25 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v61 == int32(0) {
		v120 = v58
		v121 = v59
		goto L16
	} else {
		goto L31
	}
L23:
	;
	v37 = l1
	v38 = v6
	v39 = v21
	goto L26
L24:
	;
	goto L25
L25:
	;
	v58 = l1
	v59 = v6
	v60 = v21
	v61 = int32(1)
	goto L22
L26:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v41)
	if v41 == int32(0) {
		v100 = v37
		v101 = v38
		v102 = v39
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v58 = v52
	v59 = v46
	v60 = v48
	v61 = v50
	goto L22
L28:
	;
	v45 = int32(1)
	v46 = v38 + v45
	v48 = v39 - v45
	v49 = int32(0)
	v50 = base.B2i32(v48 != v49)
	v52 = v37 + v45
	if v52&int32(3) == v49 {
		v58 = v52
		v59 = v46
		v60 = v48
		v61 = v50
		goto L22
	} else {
		goto L29
	}
L29:
	;
	if v48 != 0 {
		v37 = v52
		v38 = v46
		v39 = v48
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if base.B2i32(v64 == int32(0))|base.B2i32(base.Ui32(v60) < base.Ui32(int32(4))) != 0 {
		v93 = v58
		v94 = v59
		v95 = v60
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v71 = v58
	v72 = v59
	v73 = v60
	goto L33
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v79 = int32(-2139062144)
	if (int32(16843008)-v76|v76)&v79 != v79 {
		v100 = v71
		v101 = v72
		v102 = v73
		goto L17
	} else {
		goto L35
	}
L34:
	;
	v93 = v87
	v94 = v85
	v95 = v89
	goto L18
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v76
	v84 = int32(4)
	v85 = v72 + v84
	v87 = v71 + v84
	v89 = v73 - v84
	if base.Ui32(int32(3)) < base.Ui32(v89) {
		v71 = v87
		v72 = v85
		v73 = v89
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v100 = v93
	v101 = v94
	v102 = v95
	goto L17
L38:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v109)
	if v109 == int32(0) {
		v120 = v105
		v121 = v106
		goto L16
	} else {
		goto L40
	}
L39:
	;
	v120 = v116
	v121 = v114
	goto L16
L40:
	;
	v113 = int32(1)
	v114 = v106 + v113
	v116 = v105 + v113
	v118 = v107 - v113
	if v118 != 0 {
		v105 = v116
		v106 = v114
		v107 = v118
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v186 = v6
	goto L7
L43:
	;
	v138 = F_strlen(m, v6)
	mBase = m.M
	v141 = v138 + v6
	goto L46
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v145 = v141 - int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if base.B2i32(v146 == int32(47))&base.B2i32(base.Ui32(v6) < base.Ui32(v145)) != 0 {
		v141 = v145
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v152 = v145
	goto L49
L48:
	;
	goto L47
L49:
	;
	if base.Ui32(v6) < base.Ui32(v152) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v164 = v152
	goto L55
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v158 != int32(47) {
		v152 = v152 - int32(1)
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
	goto L53
L55:
	;
	if base.Ui32(v6) < base.Ui32(v164) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v6 == v164 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v168 = v164 - int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169 == int32(47) {
		v164 = v168
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L59
L61:
	;
	v177 = v6 + base.B2i32(v137 == int32(47))
	goto L63
L62:
	;
	v177 = v164
	goto L63
L63:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v178)
	goto L45
L64:
	;
	F_canonicalize_path_enc(m, v6)
	mBase = m.M
	v190 = F_pstrdup(m, v6)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v192 = v190
	goto L1
}
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v639 int32
	_ = v639
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	v2 = l1
	v4 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+15)) = uint8(v2)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v28 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v624 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = v4
	goto L4
L4:
	;
	v57 = v48 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v48<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	switch v63 {
	case 0:
		goto L10
	case 1:
		goto L8
	case 2:
		goto L9
	default:
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L14
	} else {
		goto L132
	}
L6:
	;
	goto L5
L7:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v57 < v586 {
		v48 = v57
		goto L4
	} else {
		goto L131
	}
L8:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v62)+36))
	if l2 != 0 {
		goto L114
	} else {
		goto L115
	}
L9:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	if v86 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if v2 == int32(0) {
		v76 = int32(1)
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v78 = F_table_open(m, v77, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if base.B2i32(l2 == int32(0))|base.B2i32(v69 != int32(1)) != 0 {
		v76 = v69
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v73 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = v73
	v76 = v73
	goto L11
L14:
	;
	return
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+21)) = uint8(v81)
	F_relation_close(m, v78, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = int32(0)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v91 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v91 < v95 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v105 = v91
	v106 = v91
	v115 = v91
	v117 = v91
	goto L23
L21:
	;
	v513 = v91
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = v513
	goto L7
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v106<<(uint(int32(2))%32))))
	if v124 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v513 = v490
	goto L22
L25:
	;
	v490 = F_lappend(m, v115, v478)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L14
	} else {
		goto L112
	}
L26:
	;
	if v163 == int32(0) {
		v475 = v105
		v478 = v124
		v487 = v117
		goto L25
	} else {
		goto L47
	}
L27:
	;
	goto L26
L28:
	;
	v125 = v124
	goto L31
L29:
	;
	goto L30
L30:
	;
	v163 = int32(0)
	goto L27
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	switch v126 - int32(15) {
	case 0:
		goto L39
	default:
		v163 = v125
		goto L27
	case 12:
		goto L38
	case 13:
		goto L37
	case 14:
		goto L36
	case 15:
		goto L35
	case 40:
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v160 != 0 {
		v125 = v160
		goto L31
	} else {
		goto L46
	}
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	if v154 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L45
	}
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v149 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L44
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	if v144 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L43
	}
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v139 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L42
	}
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	if v134 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L41
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v129 != int32(2) {
		v163 = v125
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v159 = v133
	goto L33
L41:
	;
	v159 = v125 + int32(4)
	goto L33
L42:
	;
	v159 = v125 + int32(4)
	goto L33
L43:
	;
	v159 = v125 + int32(4)
	goto L33
L44:
	;
	v159 = v125 + int32(4)
	goto L33
L45:
	;
	v159 = v125 + int32(4)
	goto L33
L46:
	;
	goto L32
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v166 != int32(6) {
		v475 = v105
		v478 = v124
		v487 = v117
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v117 != v169 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui32(v48) < base.Ui32(v169) {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	v180 = v105
	v181 = v117
	goto L51
L51:
	;
	v182 = int32(0)
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+8)))
	v187 = m.G0
	v189 = v187 - int32(96)
	m.G0 = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	switch v191 {
	case 0:
		goto L62
	case 1, 4, 5, 6, 9:
		v452 = v182
		goto L53
	case 2:
		goto L60
	case 3:
		goto L59
	case 7:
		goto L61
	case 8:
		goto L58
	default:
		goto L57
	}
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173+v169<<(uint(int32(2))%32)-int32(4))))
	v180 = v179
	v181 = v169
	goto L51
L53:
	;
	m.G0 = v189 + int32(96)
	if v452&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L14
	} else {
		goto L106
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L14
	} else {
		goto L103
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L14
	} else {
		goto L100
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L14
	} else {
		goto L97
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L14
	} else {
		goto L93
	}
L59:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v180)+68))
	if v237 == int32(0) {
		v329 = int32(1)
		goto L72
	} else {
		goto L73
	}
L60:
	;
	if v183 <= int32(0) {
		goto L54
	} else {
		goto L69
	}
L61:
	;
	if v183 <= int32(0) {
		goto L55
	} else {
		goto L66
	}
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v194 = F_SearchSysCache2(m, int32(7), v193, v183)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	if v194 == int32(0) {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v199)+91)))
	F_ReleaseCatCache(m, v194)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v452 = v201
	goto L53
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
	if v206 == int32(0) {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v209 < v183 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211+v183<<(uint(int32(2))%32)-int32(4))))
	v452 = base.B2i32(v217 == int32(0))
	goto L53
L69:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v180)+52))
	if v222 == int32(0) {
		goto L54
	} else {
		goto L70
	}
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v225 < v183 {
		goto L54
	} else {
		goto L71
	}
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227+v183<<(uint(int32(2))%32)-int32(4))))
	v452 = base.B2i32(v233 == int32(0))
	goto L53
L72:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+72)))
	if base.B2i32(v330 == int32(1))&base.B2i32(v329 == v183) != 0 {
		v452 = v182
		goto L53
	} else {
		goto L88
	}
L73:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v241 <= int32(0) {
		v329 = int32(1)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v244 = int32(0)
	if v244 < v241 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v247 = v241
	goto L77
L76:
	;
	v247 = v244
	goto L77
L77:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v255 = v182
	v267 = v182
	goto L79
L78:
	;
	v329 = v276 + int32(1)
	goto L72
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v248+v267<<(uint(int32(2))%32))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	v276 = v255 + v275
	if base.B2i32(v183 <= v276)&base.B2i32(v255 < v183) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	if v285 != 0 {
		v452 = v182
		goto L53
	} else {
		goto L85
	}
L81:
	;
	v283 = v267 + int32(1)
	if v247 != v283 {
		v255 = v276
		v267 = v283
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	goto L78
L85:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v288 = F_get_expr_result_tupdesc(m, v286, int32(1))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	if v288 == int32(0) {
		v452 = v182
		goto L53
	} else {
		goto L87
	}
L87:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v298 = int32(100)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288+v292<<(uint(int32(4))%32)+(v255^int32(-1))*v298+v183*v298)+111)))
	v452 = v304
	goto L53
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L14
	} else {
		goto L90
	}
L90:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+68)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v189)+64)) = v183
	F_errmsg(m, int32(_a_F_AcquireRewriteLocks_0), v189-int32(-64))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L14
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3515), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L14
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+84)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v189)+80)) = v183
	F_errmsg(m, int32(_a_F_AcquireRewriteLocks_0), v189+int32(80))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3525), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v381
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_3), v189)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3529), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+20)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v189)+16)) = v183
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_4), v189+int32(16))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3415), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+32)) = v183
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_5), v189+int32(32))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3438), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+48)) = v183
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_5), v189+int32(48))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L14
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3455), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L14
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v467 = v182
	goto L111
L110:
	;
	v467 = v124
	goto L111
L111:
	;
	v475 = v180
	v478 = v467
	v487 = v181
	goto L25
L112:
	;
	v493 = v106 + int32(1)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v493 < v494 {
		v105 = v475
		v106 = v493
		v115 = v490
		v117 = v487
		goto L23
	} else {
		goto L113
	}
L113:
	;
	goto L24
L114:
	;
	v561 = int32(1)
	goto L116
L115:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v524 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	F_AcquireRewriteLocks(m, v519, v2, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L14
	} else {
		goto L130
	}
L117:
	;
	v561 = base.B2i32(v557 != int32(0))
	goto L116
L118:
	;
	goto L117
L119:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if v525 <= int32(0) {
		v557 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v557 = int32(0)
	goto L118
L122:
	;
	v528 = int32(0)
	if v528 < v525 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v531 = v525
	goto L125
L124:
	;
	v531 = v528
	goto L125
L125:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	v534 = int32(0)
	goto L126
L126:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v532+v534<<(uint(int32(2))%32))))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v543 == v57 {
		v557 = v542
		goto L118
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v546 = v534 + int32(1)
	if v546 != v531 {
		v534 = v546
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	goto L7
L131:
	;
	goto L1
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v169
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_6), v25)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_7), int32(255), int32(_a_F_AcquireRewriteLocks_8))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v688 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	v627 = int32(0)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v624)+4))
	if v628 <= v627 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v639 = v627
	goto L138
L138:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653+v639<<(uint(int32(2))%32))))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+16))
	F_AcquireRewriteLocks(m, v658, v2, int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L14
	} else {
		goto L140
	}
L139:
	;
	goto L135
L140:
	;
	v663 = v639 + int32(1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v624)+4))
	if v663 < v664 {
		v639 = v663
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v695 = F_query_tree_walker_impl(m, l0, int32(1041), v25+int32(15), int32(3))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L14
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	m.G0 = v25 + int32(16)
	return
L145:
	;
	goto L144
}
func F_ActiveSnapshotSet(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_ActiveSnapshotSet[0]))
	return base.B2i32(v2 != int32(0))
}
func F_AlterObjectOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	if l0 == int32(2613) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(2995)
	goto L3
L2:
	;
	v21 = l0
	goto L3
L3:
	;
	v22 = F_get_object_attnum_oid(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v24 = F_get_object_attnum_owner(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = F_get_object_attnum_namespace(m, v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v28 = F_get_object_attnum_acl(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v30 = F_get_object_attnum_name(m, v21)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v33 = F_table_open(m, v21, int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v36 = F_get_catalog_object_by_oid_extended(m, v33, v22, l1, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v40 = v16 + int32(111)
	v41 = F_heap_getattr_2(m, v36, v24, v38, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L67
	}
L15:
	;
	if v26 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v44 = F_heap_getattr_2(m, v36, v26, v43, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v47 = int32(0)
	goto L18
L18:
	;
	if l2 != v41 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v47 = v44
	goto L18
L20:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[0]))
	if v171 != 0 {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	v49 = F_superuser(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_UnlockTuple(m, v33, v36+int32(4), int32(7))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L61
	}
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+120)))
	v102 = F_palloc0(m, v99<<(uint(int32(2))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L45
	}
L25:
	;
	if v49 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[1]))
	v53 = F_has_privs_of_role(m, v52, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v53 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v30 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[1]))
	F_check_can_set_role(m, v80, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L39
	}
L31:
	;
	v74 = F_get_object_type(m, v21, l1)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L37
	}
L32:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v60 = F_heap_getattr_2(m, v36, v30, v57, v16+int32(111))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
	v64 = v16 + int32(32)
	v69 = F_pg_snprintf(m, v64, int32(64), int32(_a_F_AlterObjectOwner_internal_0), v16+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v72 = v60
	goto L31
L36:
	;
	v72 = v64
	goto L31
L37:
	;
	F_aclcheck_error(m, int32(2), v74, v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	if v47 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v87 = F_object_aclcheck(m, int32(2615), v47, l2, int64(512))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v87 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v92 = F_get_namespace_name(m, v47)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_aclcheck_error(m, v87, int32(36), v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L24
L45:
	;
	v104 = F_palloc0(m, v99)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v106 = F_palloc0(m, v99)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v108 = int32(1)
	v109 = v24 - v108
	*(*int32)(unsafe.Add(mBase, uint32(v102+v109<<(uint(int32(2))%32)))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v109))) = uint8(v108)
	if v28 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v141 = F_heap_modify_tuple(m, v36, v140, v102, v104, v106)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L54
	}
L49:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v122 = F_heap_getattr_2(m, v36, v28, v119, v16+int32(111))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+111)))
	if v124 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v126 = v28 - int32(1)
	v130 = F_pg_detoast_datum(m, v122)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v132 = F_aclnewowner(m, v130, v41, l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102+v126<<(uint(int32(2))%32)))) = v132
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v126))) = uint8(v136)
	goto L48
L54:
	;
	F_CatalogTupleUpdate(m, v33, v141+int32(4), v141)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_UnlockTuple(m, v33, v36+int32(4), int32(7))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_changeDependencyOnOwner(m, l0, l1, l2)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v102)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v104)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v106)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	goto L20
L61:
	;
	goto L20
L62:
	;
	v172 = int32(0)
	F_RunObjectPostAlterHook(m, l0, l1, v172, v172, v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_relation_close(m, v33, int32(3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v16 + int32(112)
	return
L67:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v187 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterObjectOwner_internal_1), v16)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_AlterObjectOwner_internal_2), int32(949), int32(_a_F_AlterObjectOwner_internal_3))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ApplyLauncherForgetWorkerStartTime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherForgetWorkerStartTime[0]))
		v14 = F_dshash_delete_key(m, v11, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	if l3&int32(64) != 0 {
		v23 = int64(0)
		v24 = l1 << (uint(base.I64_extend_i32_u(l3+int32(-64))) % 64)
	} else {
		if l3 == int32(0) {
			v23 = l1
			v24 = l2
		} else {
			v15 = base.I64_extend_i32_u(l3)
			v23 = l1 << (uint(v15) % 64)
			v24 = l2<<(uint(v15)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-l3))%64))
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v24
	return
}
func F_abort(m *base.Module) {
	m.Env.X_abort_js(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_accept_weak_input(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 != 0 {
		v6 = int32(_a_F_accept_weak_input_0)
	} else {
		v6 = int32(_a_F_accept_weak_input_1)
	}
	F_set_config_option(m, int32(_a_F_accept_weak_input_2), v6, int32(6), int32(13), int32(0), int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_accept_weak_input[0])))
		return v16
	}
}
func F_accumArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v594 + v46
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v597 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0])) = v36
	if l1 != v24 {
		goto L150
	} else {
		goto L151
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v564))) = uint8(v563)
	goto L1
L3:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v501 = v392
	v502 = int32(1)
	v503 = v397
	v504 = v396
	v505 = v385
	v507 = v46
	v511 = v497
	goto L134
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L130
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L10
	} else {
		goto L126
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L10
	} else {
		goto L122
	}
L7:
	;
	v24 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L118
	}
L10:
	;
	return int32(0)
L11:
	;
	if l0 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = F_initArrayResultArr(m, l3, int32(0), l4, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v34 = l0
	goto L14
L14:
	;
	v35 = int32(_a_F_accumArrayResultArr_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0])) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v42 = v40 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v45 = v24 + int32(16)
	v46 = F_ArrayGetNItemsSafe(m, v40, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v34 = v32
	goto L14
L16:
	;
	v48 = v42 + v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v60 = v52
	goto L19
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v60 = (v53<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L19
L19:
	;
	v61 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v61 != 0 {
		goto L63
	} else {
		goto L64
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v198
	v202 = v198
	goto L20
L22:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v62 != v40+int32(1) {
		goto L4
	} else {
		goto L40
	}
L25:
	;
	v68 = v40 + int32(1)
	if int32(7) <= v68 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v68
	v75 = base.B2i32(v42 == v71)
	if v75 == v71 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v34+int32(36), v45, v42)
	goto L29
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = int32(1)
	if v75 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	base.MemoryCopy(m, v34+int32(60), v48, v42)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v88 = int32(1)
	v90 = int32(1024)
	v92 = v61 + v88
	if v92 <= v90 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v95 = v90
	goto L35
L34:
	;
	v95 = v92
	goto L35
L35:
	;
	if v95&(v95-int32(1)) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = v88 << (uint(int32(32)-base.I32_clz(v95)) % 32)
	goto L38
L37:
	;
	v102 = v95
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v102
	v104 = F_palloc(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v198 = v104
	goto L21
L40:
	;
	v109 = int32(0)
	if v109 < v40 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v113 = v40
	goto L43
L42:
	;
	v113 = v109
	goto L43
L43:
	;
	v120 = v109
	goto L45
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v169 = v168 + v61
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v169 <= v170 {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	if v120 == v113 {
		goto L44
	} else {
		goto L47
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L52
	}
L47:
	;
	v136 = int32(2)
	v137 = v120 << (uint(v136) % 32)
	v139 = v120 + int32(1)
	v141 = v139 << (uint(v136) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(32)+v141)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137+v45)))
	if v143 == v145 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(56)+v141)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v48+v137)))
	if v148 == v150 {
		v120 = v139
		goto L45
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L46
L51:
	;
	goto L50
L52:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_3), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v202 = v172
	goto L20
L57:
	;
	goto L58
L58:
	;
	v174 = v170 << (uint(int32(1)) % 32)
	if v169 < v174 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v176 = v174
	goto L61
L60:
	;
	v176 = v169
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v179 = F_repalloc(m, v178, v176)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v198 = v179
	goto L21
L63:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v43 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v228 + v61
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v231 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v225 = v43
	goto L68
L67:
	;
	v225 = (v40<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L68
L68:
	;
	base.MemoryCopy(m, v202+v217, v24+v225, v61)
	goto L65
L69:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v379 != 0 {
		goto L106
	} else {
		goto L107
	}
L70:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v234 == int32(0) {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v347 = v346 + v46
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v347 <= v348 {
		goto L69
	} else {
		goto L101
	}
L73:
	;
	v237 = int32(1)
	v239 = int32(256)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v243 = v46 + v240 + v237
	if v243 <= v239 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v246 = v239
	goto L76
L75:
	;
	v246 = v243
	goto L76
L76:
	;
	if v246&(v246-int32(1)) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v253 = v237 << (uint(int32(32)-base.I32_clz(v246)) % 32)
	goto L79
L78:
	;
	v253 = v246
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v253
	v258 = base.I32_div_s(v253+int32(7), int32(8))
	v259 = F_palloc(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v259
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v262 <= int32(0) {
		goto L69
	} else {
		goto L81
	}
L81:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v262 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v344 = v271 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v344)
	goto L69
L83:
	;
	v341 = v271 | int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v341)
	goto L69
L84:
	;
	v338 = v271 | int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v338)
	goto L69
L85:
	;
	v335 = v271 | int32(31)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v335)
	goto L69
L86:
	;
	v332 = v271 | int32(63)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v332)
	goto L69
L87:
	;
	v329 = v271 | int32(127)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v329)
	goto L69
L88:
	;
	v270 = v262
	v271 = v265
	v272 = v259
	goto L91
L89:
	;
	v311 = v265
	v312 = v259
	goto L90
L90:
	;
	v326 = v311 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v312))) = uint8(v326)
	goto L69
L91:
	;
	if v270 < int32(3) {
		goto L82
	} else {
		goto L93
	}
L92:
	;
	v311 = v303
	v312 = v305
	goto L90
L93:
	;
	if v270 == int32(3) {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(v270) < base.Ui32(int32(5)) {
		goto L84
	} else {
		goto L95
	}
L95:
	;
	if v270 == int32(5) {
		goto L85
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(v270) < base.Ui32(int32(7)) {
		goto L86
	} else {
		goto L97
	}
L97:
	;
	if v270 == int32(7) {
		goto L87
	} else {
		goto L98
	}
L98:
	;
	v297 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v297)
	v300 = v270 - int32(8)
	if v300 == int32(0) {
		goto L69
	} else {
		goto L99
	}
L99:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v304 = int32(1)
	v305 = v272 + v304
	if v300 != v304 {
		v270 = v300
		v271 = v303
		v272 = v305
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L92
L101:
	;
	v351 = v348 << (uint(int32(1)) % 32)
	if v347 < v351 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v353 = v351
	goto L104
L103:
	;
	v353 = v347
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v353
	v358 = base.I32_div_s(v353+int32(7), int32(8))
	v359 = F_repalloc(m, v231, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v359
	goto L69
L106:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v385 = v45 + v380<<(uint(int32(3))%32)
	goto L108
L107:
	;
	v385 = int32(0)
	goto L108
L108:
	;
	if v46 <= int32(0) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v392 = int32(1) << (uint(v389&int32(7)) % 32)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v395 = base.I32_div_s(v389, int32(8))
	v396 = v393 + v395
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	if v385 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	v400 = v392
	v401 = v46
	v402 = v397
	v403 = v396
	goto L111
L111:
	;
	v415 = v400 | v402
	v416 = int32(1)
	v417 = v401 - v416
	v419 = v400 << (uint(v416) % 32)
	if v419 == int32(256) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v563 = v415
	v564 = v403
	goto L2
L113:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v415)
	if v417 == int32(0) {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(int32(1)) < base.Ui32(v401) {
		v400 = v419
		v401 = v417
		v402 = v415
		goto L111
	} else {
		goto L117
	}
L116:
	;
	v425 = int32(1)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	v400 = v425
	v401 = v417
	v402 = v426
	v403 = v403 + v425
	goto L111
L117:
	;
	goto L112
L118:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_5), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_6), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_7), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_8), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L10
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
	F_errcode(m, int32(261))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v68
	F_errmsg(m, int32(_a_F_accumArrayResultArr_9), v20)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_10), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L10
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_1), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_11), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	if v502&v511 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v535 == int32(1) {
		goto L1
	} else {
		goto L149
	}
L136:
	;
	v521 = v501 | v503
	goto L138
L137:
	;
	v521 = v503 & (v501 ^ int32(-1))
	goto L138
L138:
	;
	v522 = int32(1)
	v523 = v507 - v522
	v525 = v501 << (uint(v522) % 32)
	if v525 == int32(256) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v504))) = uint8(v521)
	if v523 == int32(0) {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v535 = v525
	v536 = v521
	v537 = v504
	goto L141
L141:
	;
	v539 = v502 << (uint(int32(1)) % 32)
	if v539 == int32(256) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+1)))
	v532 = int32(1)
	v535 = v532
	v536 = v531
	v537 = v504 + v532
	goto L141
L143:
	;
	goto L135
L144:
	;
	if v523 == int32(0) {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	v548 = v539
	v549 = v505
	v550 = v511
	goto L146
L146:
	;
	if base.Ui32(int32(1)) < base.Ui32(v507) {
		v501 = v535
		v502 = v548
		v503 = v536
		v504 = v537
		v505 = v549
		v507 = v523
		v511 = v550
		goto L134
	} else {
		goto L148
	}
L147:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+1)))
	v545 = int32(1)
	v548 = v545
	v549 = v505 + v545
	v550 = v544
	goto L146
L148:
	;
	goto L143
L149:
	;
	v563 = v536
	v564 = v537
	goto L2
L150:
	;
	F_pfree(m, v24)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L10
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	m.G0 = v20 + int32(16)
	return v34
L153:
	;
	goto L152
}
func F_addFkRecurseReferenced(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v17 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v17
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v37 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_createForeignKeyActionTriggers(m, v40, v41, l0, l4, l3, l13, l14, v31+int32(28), v31+int32(24))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+119)))
	if v49 != int32(112) {
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L37
	}
L7:
	;
	m.G0 = v31 + int32(32)
	return
L8:
	;
	v53 = F_RelationGetPartitionDesc(m, l2, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v55 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v60 = int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v87 = v17
	goto L11
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v87<<(uint(int32(2))%32))))
	v100 = F_table_open(m, v98, int32(6))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	v247 = F_index_get_partition(m, v100, l3)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L26
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v105 = F_build_attrmap_by_name_if_req(m, v102, v103, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v105 == int32(0) {
		v232 = l6
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v109 = F_palloc(m, l5<<(uint(v60)%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if l5 <= int32(0) {
		v232 = v109
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v113 = int32(0)
	if l5 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v131 = v113
	v139 = v113
	goto L22
L20:
	;
	v192 = v113
	goto L21
L21:
	;
	v206 = int32(1)
	v207 = v192 << (uint(v206) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v211 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v207))))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v211<<(uint(v206)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v207))) = uint16(v217)
	v232 = v109
	goto L13
L22:
	;
	v145 = int32(1)
	v146 = v131 << (uint(v145) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v146))))
	v154 = int32(2)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+v150<<(uint(v145)%32)-v154))))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v146))) = uint16(v156)
	v159 = v146 | v154
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v163 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v159))))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+v163<<(uint(v145)%32)-v154))))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v159))) = uint16(v169)
	v172 = v131 + v154
	v174 = v139 + v154
	if v174 != l5&int32(2147483646) {
		v131 = v172
		v139 = v174
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if l5&v60 == int32(0) {
		v232 = v109
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v192 = v172
	goto L21
L26:
	;
	if v247 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_addFkConstraint(m, v31+int32(12), int32(0), v254, l0, l1, v100, v247, l4, l5, v232, l7, l8, l9, l10, l11, l12, int32(1), l15)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	F_addFkRecurseReferenced(m, l0, l1, v100, v247, v258, l5, v232, l7, l8, l9, l10, l11, l12, v65, v64, l15)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_relation_close(m, v100, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v105 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v232)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v269 = v87 + int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v269 < v270 {
		v87 = v269
		goto L11
	} else {
		goto L36
	}
L34:
	;
	F_free_attrmap(m, v105)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L12
L37:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v307 + int32(4)
	F_errmsg_internal(m, int32(_a_F_addFkRecurseReferenced_0), v31)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_addFkRecurseReferenced_1), int32(_a_F_addFkRecurseReferenced_2), int32(_a_F_addFkRecurseReferenced_3))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addNSItemForReturning(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v15 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v18 = v16
	} else {
		v18 = int32(0)
	}
	v20 = v18 << (uint(int32(5)) % 32)
	v21 = F_palloc(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v20 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			base.MemoryCopy(m, v21, v24, v20)
		} else {
		}
		if v18 <= int32(0) {
		} else {
			v29 = v18 & int32(7)
			v30 = int32(0)
			if base.Ui32(int32(8)) <= base.Ui32(v18) {
				v39 = v30
				v44 = v4
				for {
					v48 = v21 + v39<<(uint(int32(5))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+244)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+212)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+180)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+148)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+116)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+84)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+52)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = l2
					v57 = int32(8)
					v58 = v39 + v57
					v60 = v44 + v57
					if v60 != v18&int32(2147483640) {
						v39 = v58
						v44 = v60
						continue
					} else {
						break
					}
					break
				}
				if v29 == int32(0) {
				} else {
					v68 = v58
					v79 = v68
					v85 = v4
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v21+v79<<(uint(int32(5))%32))+20)) = l2
						v90 = int32(1)
						v93 = v85 + v90
						if v93 != v29 {
							v79 = v79 + v90
							v85 = v93
							continue
						} else {
							break
						}
						break
					}
				}
			} else {
				v68 = v30
				v79 = v68
				v85 = v4
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v21+v79<<(uint(int32(5))%32))+20)) = l2
					v90 = int32(1)
					v93 = v85 + v90
					if v93 != v29 {
						v79 = v79 + v90
						v85 = v93
						continue
					} else {
						break
					}
					break
				}
			}
		}
		v107 = F_palloc(m, int32(28))
		mBase = m.M
		v108 = m.ExcPending
		if v108 != 0 {
			return
		} else {
			v109 = F_makeAlias(m, l1, v15)
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v107))) = v109
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v113
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v116
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v107)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v119
				v123 = int32(0)
				F_addNSItemToQuery(m, l0, v107, v123, int32(1), v123)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_addWrd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = l4
	if v5 != 0 {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v38 = v35 + l3<<(uint(int32(3))%32)
		v39 = int32(1)
		v40 = l5 - v39
		*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
		v45 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
		if v43 <= v45+v39 {
			if v43 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = int32(2)
				v55 = F_palloc(m, int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v66 = v55
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
					v68 = l2 - l1
					v71 = F_palloc(m, v68+int32(1))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
						v76 = v74 << (uint(int32(3)) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
						if v68 != 0 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							base.MemoryCopy(m, v82, l1, v68)
						} else {
						}
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
						v88 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						if l6 != 0 {
							v97 = int32(_a_F_addWrd_0)
						} else {
							v97 = v88
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v102 = v74 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
						*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v43 << (uint(int32(1)) % 32)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v64 = F_repalloc(m, v61, v43<<(uint(int32(4))%32))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = v64
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
					v68 = l2 - l1
					v71 = F_palloc(m, v68+int32(1))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
						v76 = v74 << (uint(int32(3)) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
						if v68 != 0 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							base.MemoryCopy(m, v82, l1, v68)
						} else {
						}
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
						v88 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						if l6 != 0 {
							v97 = int32(_a_F_addWrd_0)
						} else {
							v97 = v88
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v102 = v74 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
						*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
						return
					}
				}
			}
		} else {
			v68 = l2 - l1
			v71 = F_palloc(m, v68+int32(1))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
				v76 = v74 << (uint(int32(3)) % 32)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
				if v68 != 0 {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
					base.MemoryCopy(m, v82, l1, v68)
				} else {
				}
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
				v88 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				if l6 != 0 {
					v97 = int32(_a_F_addWrd_0)
				} else {
					v97 = v88
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v102 = v74 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
				*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
				return
			}
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v10
		*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v10
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if base.Ui32(l3) < base.Ui32(v15) {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v38 = v35 + l3<<(uint(int32(3))%32)
			v39 = int32(1)
			v40 = l5 - v39
			*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
			if v43 <= v45+v39 {
				if v43 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = int32(2)
					v55 = F_palloc(m, int32(16))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v66 = v55
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
						v68 = l2 - l1
						v71 = F_palloc(m, v68+int32(1))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							if v68 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
								base.MemoryCopy(m, v82, l1, v68)
							} else {
							}
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
							v88 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v97 = int32(_a_F_addWrd_0)
							} else {
								v97 = v88
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v102 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
							*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v43 << (uint(int32(1)) % 32)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v64 = F_repalloc(m, v61, v43<<(uint(int32(4))%32))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = v64
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
						v68 = l2 - l1
						v71 = F_palloc(m, v68+int32(1))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							if v68 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
								base.MemoryCopy(m, v82, l1, v68)
							} else {
							}
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
							v88 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v97 = int32(_a_F_addWrd_0)
							} else {
								v97 = v88
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v102 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
							*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				}
			} else {
				v68 = l2 - l1
				v71 = F_palloc(m, v68+int32(1))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					v76 = v74 << (uint(int32(3)) % 32)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
					if v68 != 0 {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
						base.MemoryCopy(m, v82, l1, v68)
					} else {
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
					v88 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					if l6 != 0 {
						v97 = int32(_a_F_addWrd_0)
					} else {
						v97 = v88
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v102 = v74 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
					*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
					return
				}
			}
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(16)
				v22 = F_palloc(m, int32(128))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v32 = v22
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v38 = v35 + l3<<(uint(int32(3))%32)
					v39 = int32(1)
					v40 = l5 - v39
					*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					if v43 <= v45+v39 {
						if v43 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = int32(2)
							v55 = F_palloc(m, int32(16))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v66 = v55
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
								v68 = l2 - l1
								v71 = F_palloc(m, v68+int32(1))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									if v68 != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
										base.MemoryCopy(m, v82, l1, v68)
									} else {
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
									v88 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v97 = int32(_a_F_addWrd_0)
									} else {
										v97 = v88
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v102 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v43 << (uint(int32(1)) % 32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v64 = F_repalloc(m, v61, v43<<(uint(int32(4))%32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = v64
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
								v68 = l2 - l1
								v71 = F_palloc(m, v68+int32(1))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									if v68 != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
										base.MemoryCopy(m, v82, l1, v68)
									} else {
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
									v88 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v97 = int32(_a_F_addWrd_0)
									} else {
										v97 = v88
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v102 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						}
					} else {
						v68 = l2 - l1
						v71 = F_palloc(m, v68+int32(1))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							if v68 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
								base.MemoryCopy(m, v82, l1, v68)
							} else {
							}
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
							v88 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v97 = int32(_a_F_addWrd_0)
							} else {
								v97 = v88
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v102 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
							*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v15 << (uint(int32(1)) % 32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = F_repalloc(m, v27, v15<<(uint(int32(4))%32))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = v30
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v38 = v35 + l3<<(uint(int32(3))%32)
					v39 = int32(1)
					v40 = l5 - v39
					*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					if v43 <= v45+v39 {
						if v43 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = int32(2)
							v55 = F_palloc(m, int32(16))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v66 = v55
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
								v68 = l2 - l1
								v71 = F_palloc(m, v68+int32(1))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									if v68 != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
										base.MemoryCopy(m, v82, l1, v68)
									} else {
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
									v88 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v97 = int32(_a_F_addWrd_0)
									} else {
										v97 = v88
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v102 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v43 << (uint(int32(1)) % 32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v64 = F_repalloc(m, v61, v43<<(uint(int32(4))%32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = v64
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66
								v68 = l2 - l1
								v71 = F_palloc(m, v68+int32(1))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									if v68 != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
										base.MemoryCopy(m, v82, l1, v68)
									} else {
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
									v88 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v97 = int32(_a_F_addWrd_0)
									} else {
										v97 = v88
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v102 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						}
					} else {
						v68 = l2 - l1
						v71 = F_palloc(m, v68+int32(1))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							if v68 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
								base.MemoryCopy(m, v82, l1, v68)
							} else {
							}
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v76)+4))
							v88 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v86+v68))) = uint8(v88)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v90+v76))) = uint16(v5)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v97 = int32(_a_F_addWrd_0)
							} else {
								v97 = v88
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v76)+2)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v102 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v102
							*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				}
			}
		}
	}
}
func F_add_child_eq_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 < v14 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v16 == int32(0) {
			v21 = F_palloc0(m, v14<<(uint(int32(2))%32))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v29 = v21
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31
				v35 = F_palloc0(m, int32(28))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = l7
					v40 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+13)) = uint8(base.B2i32(l6 != v40))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v40)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(274)
					if l4 == v40 {
						v51 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v51)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v51)
					} else {
					}
					v56 = l8 << (uint(int32(2)) % 32)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56+v57)))
					v60 = F_lappend(m, v59, v35)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v62+v56))) = v60
						if int32(0) <= l2 {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v56)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
							v71 = F_bms_add_member(m, v70, l2)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v71
								return
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			v23 = int32(2)
			v27 = F_repalloc0(m, v16, v13<<(uint(v23)%32), v14<<(uint(v23)%32))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = v27
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31
				v35 = F_palloc0(m, int32(28))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = l7
					v40 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+13)) = uint8(base.B2i32(l6 != v40))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v40)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(274)
					if l4 == v40 {
						v51 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v51)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v51)
					} else {
					}
					v56 = l8 << (uint(int32(2)) % 32)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56+v57)))
					v60 = F_lappend(m, v59, v35)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v62+v56))) = v60
						if int32(0) <= l2 {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v56)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
							v71 = F_bms_add_member(m, v70, l2)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v71
								return
							}
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v35 = F_palloc0(m, int32(28))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = l7
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v35)+13)) = uint8(base.B2i32(l6 != v40))
			*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v40)
			*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(274)
			if l4 == v40 {
				v51 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)) = uint8(v51)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v51)
			} else {
			}
			v56 = l8 << (uint(int32(2)) % 32)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v56+v57)))
			v60 = F_lappend(m, v59, v35)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v62+v56))) = v60
				if int32(0) <= l2 {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v56)))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
					v71 = F_bms_add_member(m, v70, l2)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v71
						return
					}
				} else {
					return
				}
			}
		}
	}
}
func F_add_with_check_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	v9 = int32(0)
	if l3 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v188 = F_palloc0(m, int32(24))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L48
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v9
	v27 = v9
	goto L7
L5:
	;
	v66 = v9
	goto L6
L6:
	;
	if v66 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
	if l7 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v66 = v51
	goto L6
L9:
	;
	v54 = v26 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v54 < v55 {
		v26 = v54
		v27 = v51
		goto L7
	} else {
		goto L19
	}
L10:
	;
	v43 = F_copyObjectImpl(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v37 != 0 {
		v42 = v37
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v39 == int32(0) {
		v51 = v27
		goto L9
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v42 = v39
	goto L10
L16:
	;
	return
L17:
	;
	v45 = F_lappend(m, v27, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	v49 = v47 | v48
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v49)
	v51 = v45
	goto L9
L19:
	;
	goto L8
L20:
	;
	v72 = F_palloc0(m, int32(24))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(105)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = F_pstrdup(m, v77+int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+20)) = uint8(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v80
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v87 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v96
	F_ChangeVarNodes(m, v96, int32(1), l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L28
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v96 = v91
	goto L23
L25:
	;
	goto L26
L26:
	;
	v94 = F_makeBoolExpr(m, int32(1), v66, int32(-1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v96 = v94
	goto L23
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v103 = F_list_append_unique(m, v102, v72)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v103
	if l4 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v108 <= int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v114 = int32(0)
	goto L32
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v114<<(uint(int32(2))%32))))
	if l7 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L1
L34:
	;
	v172 = v114 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v172 < v173 {
		v114 = v172
		goto L32
	} else {
		goto L47
	}
L35:
	;
	v136 = F_copyObjectImpl(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L16
	} else {
		goto L41
	}
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v130 != 0 {
		v135 = v130
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v132 == int32(0) {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v135 = v132
	goto L35
L41:
	;
	F_ChangeVarNodes(m, v136, int32(1), l1)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	v142 = F_palloc0(m, int32(24))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(105)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v150 = F_pstrdup(m, v147+int32(4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v150
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v154 = F_pstrdup(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v142)+20)) = uint8(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v154
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v161 = F_list_append_unique(m, v160, v142)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v161
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+24)))
	v166 = v164 | v165
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v166)
	goto L34
L47:
	;
	goto L33
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(105)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v196 = F_pstrdup(m, v193+int32(4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	v198 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v196
	v204 = int32(1)
	v208 = F_makeConst(m, int32(16), int32(-1), v198, v204, v198, v198, v204)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+20)) = uint8(v210)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v208
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v214 = F_lappend(m, v213, v188)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v214
	goto L1
}
func F_adjust_appendrel_attrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v15 = F_adjust_appendrel_attrs_mutator(m, l1, v8+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v15
	}
}
func F_amvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_SearchSysCache1(m, int32(14), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17)+4))
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_GetIndexAmRoutineByAmId(m, v19, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v19
							F_errmsg_internal(m, int32(_a_F_amvalidate_0), v8+int32(16))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_amvalidate_1), int32(196), int32(_a_F_amvalidate_2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = m.T0[v25].(func(*base.Module, int32) int32)(m, v11)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return v28
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
				F_errmsg_internal(m, int32(_a_F_amvalidate_3), v8)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_amvalidate_1), int32(185), int32(_a_F_amvalidate_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_anycompatible_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatible_in_0), int32(376), int32(_a_F_anycompatible_in_1), int32(_a_F_anycompatible_in_2), int32(_a_F_anycompatible_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anycompatiblenonarray_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatiblenonarray_in_0), int32(377), int32(_a_F_anycompatiblenonarray_in_1), int32(_a_F_anycompatiblenonarray_in_2), int32(_a_F_anycompatiblenonarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anyelement_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anyelement_in_0), int32(374), int32(_a_F_anyelement_in_1), int32(_a_F_anyelement_in_2), int32(_a_F_anyelement_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anymultirange_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anymultirange_in_0), int32(233), int32(_a_F_anymultirange_in_1), int32(_a_F_anymultirange_in_2), int32(_a_F_anymultirange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anynonarray_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anynonarray_in_0), int32(375), int32(_a_F_anynonarray_in_1), int32(_a_F_anynonarray_in_2), int32(_a_F_anynonarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_apply_spooled_messages(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	v9 = m.G0
	v11 = v9 - int32(2224)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[0]))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[1]))
	if v54 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 == int32(3) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[2]))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	if base.B2i32(v23 == int64(0))|base.B2i32(l2 != v23) != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L4
L6:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_spooled_messages[3])) = l2
	v32 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v32 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_apply_spooled_messages[3]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+116)) = uint32(v37)
	v40 = int64(base.Ui64(v37) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+112)) = uint32(v40)
	F_errmsg(m, int32(_a_F_apply_spooled_messages_0), v11+int32(112))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(_a_F_apply_spooled_messages_2), int32(_a_F_apply_spooled_messages_3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[4]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	goto L16
L13:
	;
	v58 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_spooled_messages[5])) = v58
	goto L15
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	if base.B2i32(v62 == int32(2)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v71 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_PushActiveSnapshot(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7])) = v77
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l1
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[8]))
	v87 = v11 + int32(160)
	v92 = F_pg_snprintf(m, v87, int32(1024), int32(_a_F_apply_spooled_messages_4), v11+int32(96))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v96 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if v96 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v87
	F_errmsg_internal(m, int32(_a_F_apply_spooled_messages_5), v11+int32(80))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v109 = int32(_a_F_apply_spooled_messages_6)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[9]))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[9])) = v113
	v117 = int32(0)
	v119 = F_BufFileOpenFileSet(m, l0, v11+int32(160), v117, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2044), int32(_a_F_apply_spooled_messages_7))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[9])) = v110
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11])) = v119
	v126 = F_palloc(m, int32(_a_F_apply_spooled_messages_8))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_spooled_messages[12])) = l2
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7])) = v85
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_spooled_messages[13])) = uint8(v133)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	F_PopActiveSnapshot(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v147 = v126
	v148 = int32(0)
	goto L36
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L94
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[14]))
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L91
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11]))
	v161 = F_BufFileReadCommon(m, v156, v11+int32(124), int32(4), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	v313 = base.I32_rem_s(v207, int32(1000))
	if v313 != 0 {
		v147 = v166
		v148 = v207
		goto L36
	} else {
		goto L86
	}
L44:
	;
	if v161 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	if v163 <= int32(0) {
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v284 = v148
	goto L47
L47:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11]))
	if v287 != 0 {
		goto L76
	} else {
		goto L77
	}
L48:
	;
	v166 = F_repalloc(m, v147, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	F_BufFileReadExact(m, v169, v166, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(156)))) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v174)+32))
	v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v174)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(144)))) = v181 + v182
	goto L51
L51:
	;
	v185 = int32(_a_F_apply_spooled_messages_9)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7]))
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7])) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v166
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v194
	F_apply_dispatch(m, v11+int32(128))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[8]))
	F_MemoryContextReset(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7])) = v186
	v207 = v148 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11]))
	if v209 != 0 {
		goto L43
	} else {
		goto L54
	}
L54:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v11)+144))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[1]))
	if v213 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[4]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+20))
	goto L59
L56:
	;
	v217 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_spooled_messages[5])) = v217
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	if base.B2i32(v221 == int32(2)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v230 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L65
	}
L63:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	F_PushActiveSnapshot(m, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[7])) = v236
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[0]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = l1
	v244 = v11 + int32(1200)
	v249 = F_pg_snprintf(m, v244, int32(1024), int32(_a_F_apply_spooled_messages_4), v11+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	v251 = int32(0)
	v253 = F_BufFileOpenFileSet(m, l0, v244, v251, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v258 = F_BufFileSeek(m, v253, int32(0), int64(0), int32(2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v253)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(1196)))) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v253)+32))
	v267 = int64(*(*int32)(unsafe.Add(mBase, uint32(v253)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(1184)))) = v266 + v267
	goto L70
L70:
	;
	F_BufFileClose(m, v253)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1196))
	if v276 != v211 {
		goto L35
	} else {
		goto L74
	}
L74:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v11)+1184))
	if v278 != v210 {
		goto L35
	} else {
		goto L75
	}
L75:
	;
	v284 = v207
	goto L47
L76:
	;
	F_BufFileClose(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L7
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v295 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_spooled_messages[11])) = int32(0)
	goto L78
L80:
	;
	if v295 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v11 + int32(160)
	F_errmsg_internal(m, int32(_a_F_apply_spooled_messages_10), v11)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	m.G0 = v11 + int32(2224)
	return
L84:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2139), int32(_a_F_apply_spooled_messages_7))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v316 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	if v316 == int32(0) {
		v147 = v166
		v148 = v207
		goto L36
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v11 + int32(160)
	F_errmsg_internal(m, int32(_a_F_apply_spooled_messages_11), v11-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2132), int32(_a_F_apply_spooled_messages_7))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v147 = v166
	v148 = v207
	goto L36
L91:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v11 + int32(160)
	F_errmsg_internal(m, int32(_a_F_apply_spooled_messages_12), v11+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2095), int32(_a_F_apply_spooled_messages_7))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(1200)
	F_errmsg_internal(m, int32(_a_F_apply_spooled_messages_13), v11+int32(32))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2011), int32(_a_F_apply_spooled_messages_14))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_armenian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v453 int32
	_ = v453
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = v10
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v507 < v510 {
		v577 = int32(0)
		goto L104
	} else {
		goto L105
	}
L2:
	;
	if v128 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v128 = v100
	goto L2
L4:
	;
	if v24 <= v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v128 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v40 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v25))))
	if base.Ui32(v42) < base.Ui32(int32(192)) {
		v99 = v42
		v100 = v40
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(1413) < v99 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v46 = v33 + int32(1)
	if v46 == v24 {
		v99 = v42
		v100 = v40
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v25))))
	v51 = v49 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v42) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v25))))
	v67 = v65 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v42) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v55 = v33 + int32(2)
	if v55 != v24 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v99 = v42<<(uint(int32(6))%32)&int32(1984) | v51
	v100 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v71))))
	v99 = v84&int32(63) | (v42<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v51<<(uint(int32(12))%32) | v67<<(uint(int32(6))%32))
	v100 = int32(4)
	goto L9
L18:
	;
	v71 = v33 + int32(3)
	if v71 != v24 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v99 = v42<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v51<<(uint(int32(6))%32) | v67
	v100 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v117 = v100 + v33
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v33 = v117
	goto L4
L23:
	;
	v104 = v99 - int32(1377)
	if v104 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v110)>>(uint(v104&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = v131 + v128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v132
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = v147
	goto L30
L28:
	;
	if v253 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v253 = v224
	goto L28
L30:
	;
	if v148 <= v157 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v253 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v164 = int32(1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v149))))
	if base.Ui32(v166) < base.Ui32(int32(192)) {
		v223 = v166
		v224 = v164
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(1413) < v223 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v170 = v157 + int32(1)
	if v170 == v148 {
		v223 = v166
		v224 = v164
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v149))))
	v175 = v173 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v166) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v149))))
	v191 = v189 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v166) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v179 = v157 + int32(2)
	if v179 != v148 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v223 = v166<<(uint(int32(6))%32)&int32(1984) | v175
	v224 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v195))))
	v223 = v208&int32(63) | (v166<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v175<<(uint(int32(12))%32) | v191<<(uint(int32(6))%32))
	v224 = int32(4)
	goto L35
L44:
	;
	v195 = v157 + int32(3)
	if v195 != v148 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v223 = v166<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v175<<(uint(int32(6))%32) | v191
	v224 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v228 = v223 - int32(1377)
	if v228 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v228)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v234)>>(uint(v228&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v242 = v224 + v157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v242
	v157 = v242
	goto L30
L52:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = v256 + v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = v257
	goto L55
L53:
	;
	if v375 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v375 = v347
	goto L53
L55:
	;
	if v271 <= v280 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v375 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v287 = int32(1)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v272))))
	if base.Ui32(v289) < base.Ui32(int32(192)) {
		v346 = v289
		v347 = v287
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(1413) < v346 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v293 = v280 + int32(1)
	if v293 == v271 {
		v346 = v289
		v347 = v287
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v272))))
	v298 = v296 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v289) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v272))))
	v314 = v312 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v289) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v302 = v280 + int32(2)
	if v302 != v271 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v346 = v289<<(uint(int32(6))%32)&int32(1984) | v298
	v347 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v318))))
	v346 = v331&int32(63) | (v289<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v298<<(uint(int32(12))%32) | v314<<(uint(int32(6))%32))
	v347 = int32(4)
	goto L60
L69:
	;
	v318 = v280 + int32(3)
	if v318 != v271 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v346 = v289<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v298<<(uint(int32(6))%32) | v314
	v347 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v364 = v347 + v280
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v364
	v280 = v364
	goto L55
L74:
	;
	v351 = v346 - int32(1377)
	if v351 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v357)>>(uint(v351&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v378 + v375
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v402 = v379
	goto L81
L79:
	;
	if v498 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v498 = v469
	goto L79
L81:
	;
	if v393 <= v402 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v498 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v409 = int32(1)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v394))))
	if base.Ui32(v411) < base.Ui32(int32(192)) {
		v468 = v411
		v469 = v409
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(1413) < v468 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v415 = v402 + int32(1)
	if v415 == v393 {
		v468 = v411
		v469 = v409
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v394))))
	v420 = v418 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v411) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v394))))
	v436 = v434 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v411) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v424 = v402 + int32(2)
	if v424 != v393 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v468 = v411<<(uint(int32(6))%32)&int32(1984) | v420
	v469 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+v440))))
	v468 = v453&int32(63) | (v411<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v420<<(uint(int32(12))%32) | v436<<(uint(int32(6))%32))
	v469 = int32(4)
	goto L86
L95:
	;
	v440 = v402 + int32(3)
	if v440 != v393 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v468 = v411<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v420<<(uint(int32(6))%32) | v436
	v469 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v473 = v468 - int32(1377)
	if v473 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v473)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v479)>>(uint(v473&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v487 = v469 + v402
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v487
	v402 = v487
	goto L81
L103:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v502 + v498
	goto L1
L104:
	;
	return v577
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v510
	v516 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_2), int32(57))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v532
	v537 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_3), int32(71))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L107
	} else {
		goto L113
	}
L107:
	;
	return int32(0)
L108:
	;
	if v516 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v522
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if v522 < v525 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v527 = F_slice_del(m, l0)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	if v527 < int32(0) {
		v577 = v527
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L106
L113:
	;
	if v537 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v539
	v541 = F_slice_del(m, l0)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L107
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v546
	v551 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_4), int32(23))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L107
	} else {
		goto L119
	}
L117:
	;
	if v541 < int32(0) {
		v577 = v541
		goto L104
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	if v551 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v553
	v555 = F_slice_del(m, l0)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L107
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v560
	v565 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_5), int32(40))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L107
	} else {
		goto L125
	}
L123:
	;
	if v555 < int32(0) {
		v577 = v555
		goto L104
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v565 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v567
	v569 = F_slice_del(m, l0)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L107
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v577 = int32(1)
	goto L104
L129:
	;
	if v569 < int32(0) {
		v577 = v569
		goto L104
	} else {
		goto L130
	}
L130:
	;
	goto L128
}
func F_arrayconst_cleanup_fn(m *base.Module, l0 int32) {
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+76))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+80))
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
			F_list_free(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_pfree(m, v2)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_arraycontained(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13868(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_assign_createrole_self_grant(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_createrole_self_grant[0])) = int32(7)
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_createrole_self_grant[1])) = uint8(base.B2i32(v3 != v8))
	v14 = int32(1)
	v15 = int32(base.Ui32(v3)>>(uint(int32(2))%32)) & v14
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_createrole_self_grant[2])) = uint8(v15)
	v21 = int32(base.Ui32(v3)>>(uint(v14)%32)) & v14
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_createrole_self_grant[3])) = uint8(v21)
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_createrole_self_grant[4])) = uint8(v8)
	return
}
func F_assign_io_combine_limit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_assign_io_combine_limit[0]))
	if v5 < l0 {
		v7 = v5
	} else {
		v7 = l0
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_io_combine_limit[1])) = v7
	return
}
func F_assign_synchronous_standby_names(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_assign_synchronous_standby_names[0])) = l1
	return
}
func F_assignable_custom_variable_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
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
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v364 int32
	_ = v364
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v17 = int32(46)
	v18 = F___strchrnul(m, l0, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 == v17 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v14 - int32(-64)
	return v364
L2:
	;
	v364 = int32(0)
	goto L1
L3:
	;
	F_errfinish(m, int32(_a_F_assignable_custom_variable_name_0), v341, int32(_a_F_assignable_custom_variable_name_1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L51
	} else {
		goto L94
	}
L4:
	;
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v24 = v18
	goto L7
L6:
	;
	v24 = v4
	goto L7
L7:
	;
	goto L4
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L89
	}
L11:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_assignable_custom_variable_name[0]))
	if v216 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L12:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L50
	}
L13:
	;
	v28 = v24 - l0
	v33 = l0
	v35 = v25
	v37 = v4
	v39 = int32(1)
	goto L14
L14:
	;
	v42 = v35 & int32(255)
	if v42 == int32(46) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if base.B2i32(v35&int32(255) != int32(46))&v171 != 0 {
		goto L11
	} else {
		goto L49
	}
L16:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v176 != 0 {
		v33 = v33 + int32(1)
		v35 = v176
		v37 = v171
		v39 = base.B2i32(v42 == int32(46))
		goto L14
	} else {
		goto L48
	}
L17:
	;
	if v39 == int32(0) {
		v171 = int32(1)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v49 = base.I32_extend8_s(v35)
	goto L25
L20:
	;
	goto L12
L21:
	;
	if v155|base.B2i32(v49 < int32(0)) != 0 {
		v171 = v37
		goto L16
	} else {
		goto L46
	}
L22:
	;
	v155 = int32(0)
	goto L21
L23:
	;
	v133 = v126
	v135 = v128
	goto L40
L24:
	;
	if base.B2i32(v72 != v73) == int32(0) {
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v64 = int32(_a_F_assignable_custom_variable_name_2)
	v66 = int32(54)
	goto L26
L26:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v69 == v49&int32(255) {
		v126 = v64
		v128 = v66
		goto L23
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	v71 = int32(1)
	v72 = v66 - v71
	v73 = int32(0)
	v76 = v64 + v71
	if v76&int32(3) == v73 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if v72 != 0 {
		v64 = v76
		v66 = v72
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v89 = v49 & int32(255)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if base.B2i32(v89 == v90)|base.B2i32(base.Ui32(v72) < base.Ui32(int32(4))) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v76
	v101 = v72
	goto L35
L33:
	;
	v119 = v76
	v121 = v72
	goto L34
L34:
	;
	if v121 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v106 = v105 ^ v89*int32(16843009)
	v109 = int32(-2139062144)
	if (int32(16843008)-v106|v106)&v109 != v109 {
		v126 = v99
		v128 = v101
		goto L23
	} else {
		goto L37
	}
L36:
	;
	v119 = v114
	v121 = v116
	goto L34
L37:
	;
	v113 = int32(4)
	v114 = v99 + v113
	v116 = v101 - v113
	if base.Ui32(int32(3)) < base.Ui32(v116) {
		v99 = v114
		v101 = v116
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v126 = v119
	v128 = v121
	goto L23
L40:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v49&int32(255) == v138 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L22
L42:
	;
	v155 = v133
	goto L21
L43:
	;
	goto L44
L44:
	;
	v140 = int32(1)
	v143 = v135 - v140
	if v143 != 0 {
		v133 = v133 + v140
		v135 = v143
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v49))%64)&int64(287948969894477825) == int64(0))|(v39|base.B2i32(base.Ui32(int32(63)) < base.Ui32(v42))) != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v171 = v37
	goto L16
L48:
	;
	goto L15
L49:
	;
	goto L12
L50:
	;
	v193 = int32(0)
	v195 = F_errstart(m, l2, v193)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return int32(0)
L52:
	;
	if v195 == int32(0) {
		v364 = v193
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(_a_F_assignable_custom_variable_name_3), v12+int32(-48))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	F_errdetail(m, int32(_a_F_assignable_custom_variable_name_4), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v341 = int32(1139)
	goto L3
L57:
	;
	v364 = int32(1)
	goto L1
L58:
	;
	goto L59
L59:
	;
	v220 = int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v221 <= int32(0) {
		v364 = v220
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v224 = int32(0)
	if v224 < v221 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v227 = v221
	goto L63
L62:
	;
	v227 = v224
	goto L63
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v235 = int32(0)
	goto L64
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v228+v235<<(uint(int32(2))%32))))
	v245 = F_strlen(m, v244)
	mBase = m.M
	if v245 != v28 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v364 = v220
	goto L1
L66:
	;
	v315 = v235 + int32(1)
	if v315 != v227 {
		v235 = v315
		goto L64
	} else {
		goto L88
	}
L67:
	;
	if v28 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v291 != 0 {
		goto L66
	} else {
		goto L81
	}
L69:
	;
	v291 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v252 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v253 = l0
	v254 = v244
	v255 = v28
	v256 = v252
	goto L76
L73:
	;
	v279 = v244
	v283 = int32(0)
	goto L74
L74:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v291 = v283 - v284
	goto L68
L75:
	;
	v279 = v274
	v283 = v276
	goto L74
L76:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if base.B2i32(v256 != v258)|base.B2i32(v258 == int32(0)) != 0 {
		v274 = v254
		v276 = v256
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v274 = v268
	v276 = int32(0)
	goto L75
L78:
	;
	v264 = v255 - int32(1)
	if v264 == int32(0) {
		v274 = v254
		v276 = v256
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v267 = int32(1)
	v268 = v254 + v267
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
	if v269 != 0 {
		v253 = v253 + v267
		v254 = v268
		v255 = v264
		v256 = v269
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v292 = int32(0)
	v294 = F_errstart(m, l2, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L51
	} else {
		goto L83
	}
L83:
	;
	if v294 == int32(0) {
		v364 = v292
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L51
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l0
	F_errmsg(m, int32(_a_F_assignable_custom_variable_name_3), v12+int32(-16))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L51
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v244
	F_errdetail(m, int32(_a_F_assignable_custom_variable_name_5), v12+int32(-32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L51
	} else {
		goto L87
	}
L87:
	;
	v341 = int32(1156)
	goto L3
L88:
	;
	goto L65
L89:
	;
	v318 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L51
	} else {
		goto L90
	}
L90:
	;
	if v318 == int32(0) {
		v364 = v4
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L51
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_assignable_custom_variable_name_6), v14)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L51
	} else {
		goto L93
	}
L93:
	;
	v341 = int32(1169)
	goto L3
L94:
	;
	goto L2
}
func F_avl_sigusr2_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
