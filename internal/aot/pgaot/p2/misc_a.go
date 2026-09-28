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
	var v104 int32
	_ = v104
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
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v453 int32
	_ = v453
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v640 int32
	_ = v640
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
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
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v625 == int32(0) {
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
	v592 = m.ExcPending
	if v592 != 0 {
		goto L14
	} else {
		goto L132
	}
L6:
	;
	goto L5
L7:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v57 < v587 {
		v48 = v57
		goto L4
	} else {
		goto L131
	}
L8:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v62)+36))
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
	v78 = F_relation_open(m, v77, v76)
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
	v104 = v91
	v106 = v91
	v115 = v91
	v117 = v91
	goto L23
L21:
	;
	v514 = v91
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = v514
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
	v514 = v491
	goto L22
L25:
	;
	v491 = F_lappend(m, v115, v479)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L14
	} else {
		goto L112
	}
L26:
	;
	if v163 == int32(0) {
		v475 = v104
		v479 = v124
		v488 = v117
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
		v475 = v104
		v479 = v124
		v488 = v117
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
	v180 = v104
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
		v453 = v182
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
	if v453&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L14
	} else {
		goto L106
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L14
	} else {
		goto L103
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L14
	} else {
		goto L100
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L14
	} else {
		goto L97
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L14
	} else {
		goto L93
	}
L59:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v180)+68))
	if v238 == int32(0) {
		v330 = int32(1)
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
	v193 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v180)+16)))
	v195 = F_SearchSysCache2(m, int32(7), v193, base.I64_extend_i32_s(v183))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	if v195 == int32(0) {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+22)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v200)+91)))
	F_ReleaseCatCache(m, v195)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v453 = v202
	goto L53
L66:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v180)+96))
	if v207 == int32(0) {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v210 < v183 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212+v183<<(uint(int32(2))%32)-int32(4))))
	v453 = base.B2i32(v218 == int32(0))
	goto L53
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v180)+52))
	if v223 == int32(0) {
		goto L54
	} else {
		goto L70
	}
L70:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v226 < v183 {
		goto L54
	} else {
		goto L71
	}
L71:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228+v183<<(uint(int32(2))%32)-int32(4))))
	v453 = base.B2i32(v234 == int32(0))
	goto L53
L72:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+72)))
	if base.B2i32(v331 == int32(1))&base.B2i32(v330 == v183) != 0 {
		v453 = v182
		goto L53
	} else {
		goto L88
	}
L73:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v242 <= int32(0) {
		v330 = int32(1)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v245 = int32(0)
	if v245 < v242 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v248 = v242
	goto L77
L76:
	;
	v248 = v245
	goto L77
L77:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v257 = v182
	v268 = v182
	goto L79
L78:
	;
	v330 = v277 + int32(1)
	goto L72
L79:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v249+v268<<(uint(int32(2))%32))))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v277 = v257 + v276
	if base.B2i32(v183 <= v277)&base.B2i32(v257 < v183) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	if v286 != 0 {
		v453 = v182
		goto L53
	} else {
		goto L85
	}
L81:
	;
	v284 = v268 + int32(1)
	if v248 != v284 {
		v257 = v277
		v268 = v284
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
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v289 = F_get_expr_result_tupdesc(m, v287, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	if v289 == int32(0) {
		v453 = v182
		goto L53
	} else {
		goto L87
	}
L87:
	;
	v295 = int32(3)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+(v257^int32(-1))<<(uint(v295)%32)+v183<<(uint(v295)%32))+34)))
	v453 = int32(base.Ui32(v301&int32(4)) >> (uint(int32(2)) % 32))
	goto L53
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L14
	} else {
		goto L90
	}
L90:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+68)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v189)+64)) = v183
	F_errmsg(m, int32(_a_F_AcquireRewriteLocks_0), v189-int32(-64))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L14
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3537), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
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
	v363 = m.ExcPending
	if v363 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+84)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v189)+80)) = v183
	F_errmsg(m, int32(_a_F_AcquireRewriteLocks_0), v189+int32(80))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3547), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
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
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v382
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_3), v189)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3551), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
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
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+20)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v189)+16)) = v183
	F_errmsg_internal(m, int32(_a_F_AcquireRewriteLocks_4), v189+int32(16))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L14
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3437), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
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
	v419 = m.ExcPending
	if v419 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3460), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
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
	v435 = m.ExcPending
	if v435 != 0 {
		goto L14
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_1), int32(3477), int32(_a_F_AcquireRewriteLocks_2))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
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
	v468 = v182
	goto L111
L110:
	;
	v468 = v124
	goto L111
L111:
	;
	v475 = v180
	v479 = v468
	v488 = v181
	goto L25
L112:
	;
	v494 = v106 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v494 < v495 {
		v104 = v475
		v106 = v494
		v115 = v491
		v117 = v488
		goto L23
	} else {
		goto L113
	}
L113:
	;
	goto L24
L114:
	;
	v562 = int32(1)
	goto L116
L115:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v525 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	F_AcquireRewriteLocks(m, v520, v2, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L14
	} else {
		goto L130
	}
L117:
	;
	v562 = base.B2i32(v558 != int32(0))
	goto L116
L118:
	;
	goto L117
L119:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v526 <= int32(0) {
		v558 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v558 = int32(0)
	goto L118
L122:
	;
	v529 = int32(0)
	if v529 < v526 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v532 = v526
	goto L125
L124:
	;
	v532 = v529
	goto L125
L125:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v535 = int32(0)
	goto L126
L126:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v533+v535<<(uint(int32(2))%32))))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	if v544 == v57 {
		v558 = v543
		goto L118
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v547 = v535 + int32(1)
	if v547 != v532 {
		v535 = v547
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
	v597 = m.ExcPending
	if v597 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_AcquireRewriteLocks_7), int32(254), int32(_a_F_AcquireRewriteLocks_8))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
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
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v689 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	v628 = int32(0)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v629 <= v628 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v640 = v628
	goto L138
L138:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v640<<(uint(int32(2))%32))))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+16))
	F_AcquireRewriteLocks(m, v659, v2, int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L14
	} else {
		goto L140
	}
L139:
	;
	goto L135
L140:
	;
	v664 = v640 + int32(1)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v664 < v665 {
		v640 = v664
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v696 = F_query_tree_walker_impl(m, l0, int32(1119), v25+int32(15), int32(3))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
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
func F_AioShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
	if v9 != int32(-1) {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+92)) = int32(_a_F_AioShmemRequest_0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+84)) = int64(28)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(_a_F_AioShmemRequest_1)
		F_ShmemRequestStructWithOpts(m, v6+int32(80))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(_a_F_AioShmemRequest_2)
			v62 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
			v66 = F_mul_size(m, v62+int32(38), int32(164))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = int32(_a_F_AioShmemRequest_3)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v66
				F_ShmemRequestStructWithOpts(m, v6-int32(-64))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_AioShmemRequest_4)
					v80 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
					v84 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
					v86 = F_mul_size(m, v84, int32(128))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v88 = F_mul_size(m, v80+int32(38), v86)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(_a_F_AioShmemRequest_5)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v88
							F_ShmemRequestStructWithOpts(m, v6+int32(48))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_AioShmemRequest_6)
								v103 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
								v105 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
								v108 = F_mul_size(m, v103, v105+int32(38))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
									v112 = F_mul_size(m, v108, v111)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										v114 = F_mul_size(m, int32(8), v112)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_AioShmemRequest_7)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v114
											F_ShmemRequestStructWithOpts(m, v6+int32(32))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_AioShmemRequest_8)
												v129 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
												v131 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
												v134 = F_mul_size(m, v129, v131+int32(38))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
													v138 = F_mul_size(m, v134, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v140 = F_mul_size(m, int32(8), v138)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_AioShmemRequest_9)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v140
															F_ShmemRequestStructWithOpts(m, v6+int32(16))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return
															} else {
																v152 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[3]))
																v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
																if v153 != 0 {
																	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
																	m.T0[v153].(func(*base.Module, int32))(m, v154)
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(128)
																		return
																	}
																} else {
																	m.G0 = v6 + int32(128)
																	return
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
						}
					}
				}
			}
		}
	} else {
		v13 = int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[4]))
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
		v20 = base.I32_div_u_s(v15, v17+int32(38))
		if v20 <= v13 {
			v23 = v13
		} else {
			v23 = v20
		}
		if int32(64) <= v23 {
			v26 = int32(64)
		} else {
			v26 = v23
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v26
		v29 = v6 + int32(96)
		v32 = F_pg_snprintf(m, v29, int32(32), int32(_a_F_AioShmemRequest_10), v6)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v35 = int32(1)
			F_SetConfigOption(m, int32(_a_F_AioShmemRequest_11), v29, v35, v35)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
				if v40 != int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+92)) = int32(_a_F_AioShmemRequest_0)
					*(*int64)(unsafe.Add(mBase, uint32(v6)+84)) = int64(28)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(_a_F_AioShmemRequest_1)
					F_ShmemRequestStructWithOpts(m, v6+int32(80))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(_a_F_AioShmemRequest_2)
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
						v66 = F_mul_size(m, v62+int32(38), int32(164))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = int32(_a_F_AioShmemRequest_3)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v66
							F_ShmemRequestStructWithOpts(m, v6-int32(-64))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_AioShmemRequest_4)
								v80 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
								v86 = F_mul_size(m, v84, int32(128))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v88 = F_mul_size(m, v80+int32(38), v86)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(_a_F_AioShmemRequest_5)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v88
										F_ShmemRequestStructWithOpts(m, v6+int32(48))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_AioShmemRequest_6)
											v103 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
											v105 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
											v108 = F_mul_size(m, v103, v105+int32(38))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v111 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
												v112 = F_mul_size(m, v108, v111)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return
												} else {
													v114 = F_mul_size(m, int32(8), v112)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_AioShmemRequest_7)
														*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v114
														F_ShmemRequestStructWithOpts(m, v6+int32(32))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_AioShmemRequest_8)
															v129 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
															v131 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
															v134 = F_mul_size(m, v129, v131+int32(38))
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																v137 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
																v138 = F_mul_size(m, v134, v137)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	v140 = F_mul_size(m, int32(8), v138)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_AioShmemRequest_9)
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v140
																		F_ShmemRequestStructWithOpts(m, v6+int32(16))
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return
																		} else {
																			v152 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[3]))
																			v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
																			if v153 != 0 {
																				v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
																				m.T0[v153].(func(*base.Module, int32))(m, v154)
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					m.G0 = v6 + int32(128)
																					return
																				}
																			} else {
																				m.G0 = v6 + int32(128)
																				return
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
									}
								}
							}
						}
					}
				} else {
					F_SetConfigOption(m, int32(_a_F_AioShmemRequest_11), v29, int32(1), int32(10))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+92)) = int32(_a_F_AioShmemRequest_0)
						*(*int64)(unsafe.Add(mBase, uint32(v6)+84)) = int64(28)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(_a_F_AioShmemRequest_1)
						F_ShmemRequestStructWithOpts(m, v6+int32(80))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(_a_F_AioShmemRequest_2)
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
							v66 = F_mul_size(m, v62+int32(38), int32(164))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = int32(_a_F_AioShmemRequest_3)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v66
								F_ShmemRequestStructWithOpts(m, v6-int32(-64))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_AioShmemRequest_4)
									v80 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
									v84 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
									v86 = F_mul_size(m, v84, int32(128))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										v88 = F_mul_size(m, v80+int32(38), v86)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(_a_F_AioShmemRequest_5)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v88
											F_ShmemRequestStructWithOpts(m, v6+int32(48))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_AioShmemRequest_6)
												v103 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
												v105 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
												v108 = F_mul_size(m, v103, v105+int32(38))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
													v112 = F_mul_size(m, v108, v111)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return
													} else {
														v114 = F_mul_size(m, int32(8), v112)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_AioShmemRequest_7)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v114
															F_ShmemRequestStructWithOpts(m, v6+int32(32))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_AioShmemRequest_8)
																v129 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[2]))
																v131 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[1]))
																v134 = F_mul_size(m, v129, v131+int32(38))
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return
																} else {
																	v137 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[0]))
																	v138 = F_mul_size(m, v134, v137)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return
																	} else {
																		v140 = F_mul_size(m, int32(8), v138)
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_AioShmemRequest_9)
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v140
																			F_ShmemRequestStructWithOpts(m, v6+int32(16))
																			mBase = m.M
																			v150 = m.ExcPending
																			if v150 != 0 {
																				return
																			} else {
																				v152 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemRequest[3]))
																				v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
																				if v153 != 0 {
																					v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
																					m.T0[v153].(func(*base.Module, int32))(m, v154)
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						m.G0 = v6 + int32(128)
																						return
																					}
																				} else {
																					m.G0 = v6 + int32(128)
																					return
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
}
func F_AlterObjectOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	if l0 == int32(2613) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(2995)
	goto L3
L2:
	;
	v22 = l0
	goto L3
L3:
	;
	v23 = F_get_object_attnum_oid(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v25 = F_get_object_attnum_owner(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = F_get_object_attnum_namespace(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v29 = F_get_object_attnum_acl(m, v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v31 = F_get_object_attnum_name(m, v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v34 = F_table_open(m, v22, int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v37 = F_get_catalog_object_by_oid_extended(m, v34, v23, l1, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v41 = v17 + int32(111)
	v42 = F_heap_getattr_2(m, v37, v25, v39, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L67
	}
L15:
	;
	v44 = base.I32_wrap_i64(v42)
	if v27 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v46 = F_heap_getattr_2(m, v37, v27, v45, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v50 = int32(0)
	goto L18
L18:
	;
	if l2 != v44 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v50 = base.I32_wrap_i64(v46)
	goto L18
L20:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[0]))
	if v179 != 0 {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	v52 = F_superuser(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_UnlockTuple(m, v34, v37+int32(4), int32(7))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L61
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+120)))
	v106 = F_palloc0(m, v103<<(uint(int32(3))%32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L45
	}
L25:
	;
	if v52 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[1]))
	v56 = F_has_privs_of_role(m, v55, v44)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v56 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v31 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectOwner_internal[1]))
	F_check_can_set_role(m, v84, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L39
	}
L31:
	;
	v78 = F_get_object_type(m, v22, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L37
	}
L32:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v63 = F_heap_getattr_2(m, v37, v31, v60, v17+int32(111))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
	v68 = v17 + int32(32)
	v73 = F_pg_snprintf(m, v68, int32(64), int32(_a_F_AlterObjectOwner_internal_0), v17+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v76 = base.I32_wrap_i64(v63)
	goto L31
L36:
	;
	v76 = v68
	goto L31
L37:
	;
	F_aclcheck_error(m, int32(2), v78, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	if v50 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v91 = F_object_aclcheck(m, int32(2615), v50, l2, int64(512))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v91 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v96 = F_get_namespace_name(m, v50)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_aclcheck_error(m, v91, int32(37), v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L24
L45:
	;
	v108 = F_palloc0(m, v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v110 = F_palloc0(m, v103)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v112 = int32(1)
	v113 = v25 - v112
	*(*int64)(unsafe.Add(mBase, uint32(v106+v113<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(l2)
	*(*uint8)(unsafe.Add(mBase, uint32(v110+v113))) = uint8(v112)
	if v29 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v148 = F_heap_modify_tuple(m, v37, v147, v106, v108, v110)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L54
	}
L49:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v127 = F_heap_getattr_2(m, v37, v29, v124, v17+int32(111))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+111)))
	if v129 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v131 = v29 - int32(1)
	v136 = F_pg_detoast_datum(m, base.I32_wrap_i64(v127))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v138 = F_aclnewowner(m, v136, v44, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106+v131<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v138)
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v110+v131))) = uint8(v143)
	goto L48
L54:
	;
	F_CatalogTupleUpdate(m, v34, v148+int32(4), v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_UnlockTuple(m, v34, v37+int32(4), int32(7))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_changeDependencyOnOwner(m, l0, l1, l2)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v106)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v108)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v110)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	v180 = int32(0)
	F_RunObjectPostAlterHook(m, l0, l1, v180, v180, v180)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v17 + int32(112)
	return
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v195 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterObjectOwner_internal_1), v17)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_AlterObjectOwner_internal_2), int32(943), int32(_a_F_AlterObjectOwner_internal_3))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
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
func F_accept_weak_input(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 == int64(0) {
		v8 = int32(_a_F_accept_weak_input_0)
	} else {
		v8 = int32(_a_F_accept_weak_input_1)
	}
	F_set_config_option(m, int32(_a_F_accept_weak_input_2), v8, int32(6), int32(13), int32(0), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v18 = int64(*(*uint8)(unsafe.Add(mBase, _c_F_accept_weak_input[0])))
		return v18
	}
}
func F_accumArrayResultArr(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
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
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v61
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v637 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0])) = v39
	if v27 != v26 {
		goto L157
	} else {
		goto L158
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v601)
	goto L1
L3:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	v537 = v407
	v538 = int32(1)
	v539 = v412
	v540 = v411
	v543 = v47
	v544 = v400
	v547 = v533
	goto L141
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L11
	} else {
		goto L137
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L133
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L11
	} else {
		goto L129
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L11
	} else {
		goto L125
	}
L8:
	;
	v26 = base.I32_wrap_i64(l1)
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L11
	} else {
		goto L121
	}
L11:
	;
	return int32(0)
L12:
	;
	if l0 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = F_initArrayResultArr(m, l3, int32(0), l4, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v37 = l0
	goto L15
L15:
	;
	v38 = int32(_a_F_accumArrayResultArr_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResultArr[0])) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v46 = v27 + int32(16)
	v47 = F_ArrayGetNItemsSafe(m, v44, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v37 = v35
	goto L15
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L20
L19:
	;
	v59 = v49
	goto L20
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v61 = v60 + v47
	if base.Ui32(int32(134217728)) <= base.Ui32(v61) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v64 = int32(2)
	v65 = v44 << (uint(v64) % 32)
	v66 = v46 + v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v70 = int32(base.Ui32(v67)>>(uint(v64)%32)) - v59
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v71 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v70 != 0 {
		goto L65
	} else {
		goto L66
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v211
	v215 = v211
	goto L22
L24:
	;
	if v44 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v71 != v44+int32(1) {
		goto L4
	} else {
		goto L42
	}
L27:
	;
	v77 = v44 + int32(1)
	if int32(7) <= v77 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v77
	v84 = base.B2i32(v65 == v80)
	if v84 == v80 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	base.MemoryCopy(m, v37+int32(36), v46, v65)
	goto L31
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = int32(1)
	if v84 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	base.MemoryCopy(m, v37+int32(60), v66, v65)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v97 = int32(1)
	v99 = int32(1024)
	v101 = v70 + v97
	if v101 <= v99 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = v99
	goto L37
L36:
	;
	v104 = v101
	goto L37
L37:
	;
	if v104&(v104-int32(1)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v111 = v97 << (uint(int32(32)-base.I32_clz(v104)) % 32)
	goto L40
L39:
	;
	v111 = v104
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v111
	v113 = F_palloc(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v211 = v113
	goto L23
L42:
	;
	v118 = int32(0)
	if v118 < v44 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = v44
	goto L45
L44:
	;
	v122 = v118
	goto L45
L45:
	;
	v129 = v118
	goto L47
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v180 = v179 + v70
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v180 <= v181 {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	if v129 == v122 {
		goto L46
	} else {
		goto L49
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L11
	} else {
		goto L54
	}
L49:
	;
	v147 = int32(2)
	v148 = v129 << (uint(v147) % 32)
	v150 = v129 + int32(1)
	v152 = v150 << (uint(v147) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(32)+v152)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148+v46)))
	if v154 == v156 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(56)+v152)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v148+v66)))
	if v159 == v161 {
		v129 = v150
		goto L47
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L48
L53:
	;
	goto L52
L54:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_1), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_3), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
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
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v215 = v183
	goto L22
L59:
	;
	goto L60
L60:
	;
	v185 = v181 << (uint(int32(1)) % 32)
	if v180 < v185 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v187 = v185
	goto L63
L62:
	;
	v187 = v180
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v190 = F_repalloc(m, v189, v187)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v211 = v190
	goto L23
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v43 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v243 + v70
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v246 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v240 = v43
	goto L70
L69:
	;
	v240 = (v44<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L70
L70:
	;
	base.MemoryCopy(m, v215+v232, v27+v240, v70)
	goto L67
L71:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v394 != 0 {
		goto L109
	} else {
		goto L110
	}
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v249 == int32(0) {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v61 <= v361 {
		goto L71
	} else {
		goto L104
	}
L75:
	;
	v254 = int32(255)
	if v61 <= v254 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v257 = v254
	goto L78
L77:
	;
	v257 = v61
	goto L78
L78:
	;
	v259 = v257 + int32(1)
	if v259&v257 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v264 = int32(1) << (uint(int32(32)-base.I32_clz(v259)) % 32)
	goto L81
L80:
	;
	v264 = v259
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v264
	v270 = F_palloc(m, int32(base.Ui32(v264+int32(7))>>(uint(int32(3))%32)))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v270
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v273 <= int32(0) {
		goto L71
	} else {
		goto L83
	}
L83:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v273 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v359 = v282 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v359)
	goto L71
L85:
	;
	v356 = v282 | int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v356)
	goto L71
L86:
	;
	v353 = v282 | int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v353)
	goto L71
L87:
	;
	v350 = v282 | int32(31)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v350)
	goto L71
L88:
	;
	v347 = v282 | int32(63)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v347)
	goto L71
L89:
	;
	v344 = v282 | int32(127)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v344)
	goto L71
L90:
	;
	v341 = v324 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v326))) = uint8(v341)
	goto L71
L91:
	;
	v324 = v276
	v326 = v270
	goto L90
L92:
	;
	goto L93
L93:
	;
	v281 = v273
	v282 = v276
	v283 = v270
	goto L94
L94:
	;
	if v281 < int32(3) {
		goto L84
	} else {
		goto L96
	}
L95:
	;
	v324 = v316
	v326 = v318
	goto L90
L96:
	;
	if v281 == int32(3) {
		goto L85
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v281) < base.Ui32(int32(5)) {
		goto L86
	} else {
		goto L98
	}
L98:
	;
	if v281 == int32(5) {
		goto L87
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(v281) < base.Ui32(int32(7)) {
		goto L88
	} else {
		goto L100
	}
L100:
	;
	if v281 == int32(7) {
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v310 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v310)
	v313 = v281 - int32(8)
	if v313 == int32(0) {
		goto L71
	} else {
		goto L102
	}
L102:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v317 = int32(1)
	v318 = v283 + v317
	if v313 != v317 {
		v281 = v313
		v282 = v316
		v283 = v318
		goto L94
	} else {
		goto L103
	}
L103:
	;
	goto L95
L104:
	;
	v364 = v361 << (uint(int32(1)) % 32)
	if v61 < v364 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v366 = v364
	goto L107
L106:
	;
	v366 = v61
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v366
	v371 = base.I32_div_s(v366+int32(7), int32(8))
	v372 = F_repalloc(m, v246, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v372
	goto L71
L109:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v400 = v46 + v395<<(uint(int32(3))%32)
	goto L111
L110:
	;
	v400 = int32(0)
	goto L111
L111:
	;
	if v47 <= int32(0) {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v407 = int32(1) << (uint(v404&int32(7)) % 32)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v410 = base.I32_div_s(v404, int32(8))
	v411 = v408 + v410
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v400 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	v415 = v407
	v417 = v412
	v418 = v411
	v421 = v47
	goto L114
L114:
	;
	v432 = v415 | v417
	v433 = int32(1)
	v434 = v421 - v433
	v436 = v415 << (uint(v433) % 32)
	if v436 == int32(256) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v601 = v432
	v602 = v418
	goto L2
L116:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v432)
	if v434 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(int32(1)) < base.Ui32(v421) {
		v415 = v436
		v417 = v432
		v421 = v434
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v442 = int32(1)
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	v415 = v442
	v417 = v443
	v418 = v418 + v442
	v421 = v434
	goto L114
L120:
	;
	goto L115
L121:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L11
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_5), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_6), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(134217727)
	F_errmsg(m, int32(_a_F_accumArrayResultArr_7), v22)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_8), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L11
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_9), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_10), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v77
	F_errmsg(m, int32(_a_F_accumArrayResultArr_11), v22+int32(16))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_12), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_accumArrayResultArr_1), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_accumArrayResultArr_2), int32(_a_F_accumArrayResultArr_13), int32(_a_F_accumArrayResultArr_4))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	if v538&v547 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v573 == int32(1) {
		goto L1
	} else {
		goto L156
	}
L143:
	;
	v559 = v537 | v539
	goto L145
L144:
	;
	v559 = v539 & (v537 ^ int32(-1))
	goto L145
L145:
	;
	v560 = int32(1)
	v561 = v543 - v560
	v563 = v537 << (uint(v560) % 32)
	if v563 == int32(256) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v559)
	if v561 == int32(0) {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	v573 = v563
	v574 = v559
	v575 = v540
	goto L148
L148:
	;
	v577 = v538 << (uint(int32(1)) % 32)
	if v577 == int32(256) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	v570 = int32(1)
	v573 = v570
	v574 = v569
	v575 = v540 + v570
	goto L148
L150:
	;
	goto L142
L151:
	;
	if v561 == int32(0) {
		goto L150
	} else {
		goto L154
	}
L152:
	;
	v586 = v577
	v587 = v544
	v588 = v547
	goto L153
L153:
	;
	if base.Ui32(int32(1)) < base.Ui32(v543) {
		v537 = v573
		v538 = v586
		v539 = v574
		v540 = v575
		v543 = v561
		v544 = v587
		v547 = v588
		goto L141
	} else {
		goto L155
	}
L154:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+1)))
	v583 = int32(1)
	v586 = v583
	v587 = v544 + v583
	v588 = v582
	goto L153
L155:
	;
	goto L150
L156:
	;
	v601 = v574
	v602 = v575
	goto L2
L157:
	;
	F_pfree(m, v27)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L11
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	m.G0 = v22 + int32(32)
	return v37
L160:
	;
	goto L159
}
func F_aclmask_direct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v12 int32
	_ = v12
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v67 int64
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v5 = int64(0)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_check_acl(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L25
	}
L4:
	;
	return int64(0)
L5:
	;
	if l3 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int64(0)
L7:
	;
	goto L8
L8:
	;
	if l1 != l2 {
		v24 = v5
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v19 = l3 & int64(-4294967296)
	if v19 == int64(0) {
		v24 = v5
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v19 != l3 {
		v24 = v19
		goto L9
	} else {
		goto L12
	}
L12:
	;
	return l3
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L15
L14:
	;
	v37 = v27
	goto L15
L15:
	;
	if v26 <= int32(0) {
		v67 = v24
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return v67
L17:
	;
	v42 = int32(0)
	v46 = v24
	goto L18
L18:
	;
	v52 = l0 + v37 + v42<<(uint(int32(4))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if l1 == v53 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v67 = v59
	goto L16
L20:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
	v57 = v55&l3 | v46
	if v57 == l3 {
		v67 = v57
		goto L16
	} else {
		goto L23
	}
L21:
	;
	v59 = v46
	goto L22
L22:
	;
	v61 = v42 + int32(1)
	if v61 != v26 {
		v42 = v61
		v46 = v59
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v59 = v57
	goto L22
L24:
	;
	goto L19
L25:
	;
	F_errmsg_internal(m, int32(_a_F_aclmask_direct_0), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_aclmask_direct_1), int32(1514), int32(_a_F_aclmask_direct_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addFkRecurseReferenced(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v227 int32
	_ = v227
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	v17 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v17
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v36 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_createForeignKeyActionTriggers(m, v39, v40, l0, l4, l3, l13, l14, v30+int32(28), v30+int32(24))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+119)))
	if v48 != int32(112) {
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
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L37
	}
L7:
	;
	m.G0 = v30 + int32(32)
	return
L8:
	;
	v52 = F_RelationGetPartitionDesc(m, l2, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v54 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v84 = v17
	goto L11
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v84<<(uint(int32(2))%32))))
	v96 = F_table_open(m, v94, int32(6))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	v241 = F_index_get_partition(m, v96, l3)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L26
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v101 = F_build_attrmap_by_name_if_req(m, v98, v99, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v101 == int32(0) {
		v227 = l6
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v106 = F_palloc_mul(m, int32(2), l5)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if l5 <= int32(0) {
		v227 = v106
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v110 = int32(0)
	if l5 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v128 = v110
	v136 = v110
	goto L22
L20:
	;
	v188 = v110
	goto L21
L21:
	;
	v201 = int32(1)
	v202 = v188 << (uint(v201) % 32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v202))))
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204+v206<<(uint(v201)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v106+v202))) = uint16(v212)
	v227 = v106
	goto L13
L22:
	;
	v141 = int32(1)
	v142 = v128 << (uint(v141) % 32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v142))))
	v150 = int32(2)
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+v146<<(uint(v141)%32)-v150))))
	*(*uint16)(unsafe.Add(mBase, uint32(v106+v142))) = uint16(v152)
	v155 = v142 | v150
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v155))))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157+v159<<(uint(v141)%32)-v150))))
	*(*uint16)(unsafe.Add(mBase, uint32(v106+v155))) = uint16(v165)
	v168 = v128 + v150
	v170 = v136 + v150
	if v170 != l5&int32(2147483646) {
		v128 = v168
		v136 = v170
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if l5&int32(1) == int32(0) {
		v227 = v106
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v188 = v168
	goto L21
L26:
	;
	if v241 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_addFkConstraint(m, v30+int32(12), int32(0), v248, l0, l1, v96, v241, l4, l5, v227, l7, l8, l9, l10, l11, l12, int32(1), l15)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	F_addFkRecurseReferenced(m, l0, l1, v96, v241, v252, l5, v227, l7, l8, l9, l10, l11, l12, v62, v61, l15)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_relation_close(m, v96, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v227)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v263 = v84 + int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v263 < v264 {
		v84 = v263
		goto L11
	} else {
		goto L36
	}
L34:
	;
	F_free_attrmap(m, v101)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v300 + int32(4)
	F_errmsg_internal(m, int32(_a_F_addFkRecurseReferenced_0), v30)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_addFkRecurseReferenced_1), int32(_a_F_addFkRecurseReferenced_2), int32(_a_F_addFkRecurseReferenced_3))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v16 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v19 = v17
	} else {
		v19 = int32(0)
	}
	v20 = F_palloc_mul(m, int32(32), v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v23 = v19 << (uint(int32(5)) % 32)
		if v23 != 0 {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			base.MemoryCopy(m, v20, v25, v23)
		} else {
		}
		if v19 <= int32(0) {
		} else {
			v30 = v19 & int32(7)
			v31 = int32(0)
			if base.Ui32(int32(8)) <= base.Ui32(v19) {
				v40 = v31
				v45 = v4
				for {
					v49 = v20 + v40<<(uint(int32(5))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v49)+244)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+212)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+180)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+148)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+116)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = l2
					v58 = int32(8)
					v59 = v40 + v58
					v61 = v45 + v58
					if v61 != v19&int32(2147483640) {
						v40 = v59
						v45 = v61
						continue
					} else {
						break
					}
					break
				}
				if v30 == int32(0) {
				} else {
					v69 = v59
					v80 = v69
					v86 = v4
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v20+v80<<(uint(int32(5))%32))+20)) = l2
						v91 = int32(1)
						v94 = v86 + v91
						if v94 != v30 {
							v80 = v80 + v91
							v86 = v94
							continue
						} else {
							break
						}
						break
					}
				}
			} else {
				v69 = v31
				v80 = v69
				v86 = v4
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v20+v80<<(uint(int32(5))%32))+20)) = l2
					v91 = int32(1)
					v94 = v86 + v91
					if v94 != v30 {
						v80 = v80 + v91
						v86 = v94
						continue
					} else {
						break
					}
					break
				}
			}
		}
		v108 = F_palloc(m, int32(28))
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return
		} else {
			v110 = F_makeAlias(m, l1, v16)
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v114
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v117
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v120
				v124 = int32(0)
				F_addNSItemToQuery(m, l0, v108, v124, int32(1), v124)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v5 = l4
	if v5 != 0 {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v39 = v36 + l3<<(uint(int32(3))%32)
		v40 = int32(1)
		v41 = l5 - v40
		*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v41)
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
		v46 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
		if v44 <= v46+v40 {
			if v44 == int32(0) {
				v53 = int32(2)
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v53
				v57 = F_palloc_mul(m, int32(8), v53)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v68 = v57
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
					v71 = l2 - l1
					v74 = F_palloc(m, v71+int32(1))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
						v79 = v77 << (uint(int32(3)) % 32)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
						if v71 != 0 {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
							base.MemoryCopy(m, v85, l1, v71)
						} else {
						}
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
						v91 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						if l6 != 0 {
							v100 = int32(_a_F_addWrd_0)
						} else {
							v100 = v91
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v105 = v77 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
						*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
						return
					}
				}
			} else {
				v61 = v44 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v61
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v65 = F_repalloc_mul(m, v63, int32(8), v61)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					v68 = v65
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
					v71 = l2 - l1
					v74 = F_palloc(m, v71+int32(1))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
						v79 = v77 << (uint(int32(3)) % 32)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
						if v71 != 0 {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
							base.MemoryCopy(m, v85, l1, v71)
						} else {
						}
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
						v91 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						if l6 != 0 {
							v100 = int32(_a_F_addWrd_0)
						} else {
							v100 = v91
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v105 = v77 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
						*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
						return
					}
				}
			}
		} else {
			v71 = l2 - l1
			v74 = F_palloc(m, v71+int32(1))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
				v79 = v77 << (uint(int32(3)) % 32)
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
				if v71 != 0 {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
					base.MemoryCopy(m, v85, l1, v71)
				} else {
				}
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
				v91 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				if l6 != 0 {
					v100 = int32(_a_F_addWrd_0)
				} else {
					v100 = v91
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v105 = v77 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
				*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
				return
			}
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v10
		*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v10
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if base.Ui32(l3) < base.Ui32(v15) {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v39 = v36 + l3<<(uint(int32(3))%32)
			v40 = int32(1)
			v41 = l5 - v40
			*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v41)
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
			v46 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
			if v44 <= v46+v40 {
				if v44 == int32(0) {
					v53 = int32(2)
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v53
					v57 = F_palloc_mul(m, int32(8), v53)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v68 = v57
						*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
						v71 = l2 - l1
						v74 = F_palloc(m, v71+int32(1))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v79 = v77 << (uint(int32(3)) % 32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
							if v71 != 0 {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
								base.MemoryCopy(m, v85, l1, v71)
							} else {
							}
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
							v91 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if l6 != 0 {
								v100 = int32(_a_F_addWrd_0)
							} else {
								v100 = v91
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v105 = v77 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
							*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				} else {
					v61 = v44 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v61
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v65 = F_repalloc_mul(m, v63, int32(8), v61)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v68 = v65
						*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
						v71 = l2 - l1
						v74 = F_palloc(m, v71+int32(1))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v79 = v77 << (uint(int32(3)) % 32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
							if v71 != 0 {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
								base.MemoryCopy(m, v85, l1, v71)
							} else {
							}
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
							v91 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if l6 != 0 {
								v100 = int32(_a_F_addWrd_0)
							} else {
								v100 = v91
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v105 = v77 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
							*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				}
			} else {
				v71 = l2 - l1
				v74 = F_palloc(m, v71+int32(1))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					v79 = v77 << (uint(int32(3)) % 32)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
					if v71 != 0 {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
						base.MemoryCopy(m, v85, l1, v71)
					} else {
					}
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
					v91 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					if l6 != 0 {
						v100 = int32(_a_F_addWrd_0)
					} else {
						v100 = v91
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v105 = v77 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
					*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
					return
				}
			}
		} else {
			if v15 == int32(0) {
				v19 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v19
				v23 = F_palloc_mul(m, int32(8), v19)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v33 = v23
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v33
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v39 = v36 + l3<<(uint(int32(3))%32)
					v40 = int32(1)
					v41 = l5 - v40
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v41)
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					if v44 <= v46+v40 {
						if v44 == int32(0) {
							v53 = int32(2)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v53
							v57 = F_palloc_mul(m, int32(8), v53)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v68 = v57
								*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
								v71 = l2 - l1
								v74 = F_palloc(m, v71+int32(1))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v79 = v77 << (uint(int32(3)) % 32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
									if v71 != 0 {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
										base.MemoryCopy(m, v85, l1, v71)
									} else {
									}
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									if l6 != 0 {
										v100 = int32(_a_F_addWrd_0)
									} else {
										v100 = v91
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v105 = v77 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							v61 = v44 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v65 = F_repalloc_mul(m, v63, int32(8), v61)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v68 = v65
								*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
								v71 = l2 - l1
								v74 = F_palloc(m, v71+int32(1))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v79 = v77 << (uint(int32(3)) % 32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
									if v71 != 0 {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
										base.MemoryCopy(m, v85, l1, v71)
									} else {
									}
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									if l6 != 0 {
										v100 = int32(_a_F_addWrd_0)
									} else {
										v100 = v91
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v105 = v77 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						}
					} else {
						v71 = l2 - l1
						v74 = F_palloc(m, v71+int32(1))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v79 = v77 << (uint(int32(3)) % 32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
							if v71 != 0 {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
								base.MemoryCopy(m, v85, l1, v71)
							} else {
							}
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
							v91 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if l6 != 0 {
								v100 = int32(_a_F_addWrd_0)
							} else {
								v100 = v91
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v105 = v77 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
							*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				}
			} else {
				v26 = v15 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = F_repalloc_mul(m, v28, int32(8), v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v33 = v30
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v33
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v39 = v36 + l3<<(uint(int32(3))%32)
					v40 = int32(1)
					v41 = l5 - v40
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v41)
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[0]))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
					if v44 <= v46+v40 {
						if v44 == int32(0) {
							v53 = int32(2)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v53
							v57 = F_palloc_mul(m, int32(8), v53)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v68 = v57
								*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
								v71 = l2 - l1
								v74 = F_palloc(m, v71+int32(1))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v79 = v77 << (uint(int32(3)) % 32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
									if v71 != 0 {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
										base.MemoryCopy(m, v85, l1, v71)
									} else {
									}
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									if l6 != 0 {
										v100 = int32(_a_F_addWrd_0)
									} else {
										v100 = v91
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v105 = v77 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							v61 = v44 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[0])) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v65 = F_repalloc_mul(m, v63, int32(8), v61)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v68 = v65
								*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v68
								v71 = l2 - l1
								v74 = F_palloc(m, v71+int32(1))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
									v79 = v77 << (uint(int32(3)) % 32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
									if v71 != 0 {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
										base.MemoryCopy(m, v85, l1, v71)
									} else {
									}
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									if l6 != 0 {
										v100 = int32(_a_F_addWrd_0)
									} else {
										v100 = v91
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									v105 = v77 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						}
					} else {
						v71 = l2 - l1
						v74 = F_palloc(m, v71+int32(1))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_addWrd[1]))
							v79 = v77 << (uint(int32(3)) % 32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v79+v80)+4)) = v74
							if v71 != 0 {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)+4))
								base.MemoryCopy(m, v85, l1, v71)
							} else {
							}
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v79)+4))
							v91 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v89+v71))) = uint8(v91)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v93+v79))) = uint16(v5)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if l6 != 0 {
								v100 = int32(_a_F_addWrd_0)
							} else {
								v100 = v91
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v96+v79)+2)) = uint16(v100)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v105 = v77 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_addWrd[1])) = v105
							*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(3))%32))+4)) = int32(0)
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 < v14 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v16 == int32(0) {
			v20 = F_palloc0_mul(m, int32(4), v14)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v31 = v20
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v33
				v37 = F_palloc0(m, int32(28))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = l7
					v42 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+13)) = uint8(base.B2i32(l6 != v42))
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v42)
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(277)
					if l4 == v42 {
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v53)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v53)
					} else {
					}
					v58 = l8 << (uint(int32(2)) % 32)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59)))
					v62 = F_lappend(m, v61, v37)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v64+v58))) = v62
						if int32(0) <= l2 {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v58)))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+144))
							v73 = F_bms_add_member(m, v72, l2)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v71)+144)) = v73
								return
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			v23 = F_mul_size(m, int32(4), v13)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v27 = F_mul_size(m, int32(4), v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = F_repalloc0(m, v16, v23, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = v29
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v31
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v33
						v37 = F_palloc0(m, int32(28))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l6
							*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = l7
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v37)+13)) = uint8(base.B2i32(l6 != v42))
							*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v42)
							*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(277)
							if l4 == v42 {
								v53 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v53)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v53)
							} else {
							}
							v58 = l8 << (uint(int32(2)) % 32)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59)))
							v62 = F_lappend(m, v61, v37)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v64+v58))) = v62
								if int32(0) <= l2 {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v58)))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+144))
									v73 = F_bms_add_member(m, v72, l2)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v71)+144)) = v73
										return
									}
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v37 = F_palloc0(m, int32(28))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = l7
			v42 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v37)+13)) = uint8(base.B2i32(l6 != v42))
			*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v42)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(277)
			if l4 == v42 {
				v53 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v53)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v53)
			} else {
			}
			v58 = l8 << (uint(int32(2)) % 32)
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59)))
			v62 = F_lappend(m, v61, v37)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v64+v58))) = v62
				if int32(0) <= l2 {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v58)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+144))
					v73 = F_bms_add_member(m, v72, l2)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v71)+144)) = v73
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
	v208 = F_makeConst(m, int32(16), int32(-1), v198, v204, int64(0), v198, v204)
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
func F_align_fetch_then_add(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	if int32(0) < l3 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v14 = (l4 + v8 - int32(1)) & (int32(0) - l4)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14 + l3
		v17 = l0 + v14
		if l2 != 0 {
			switch l3 - int32(1) {
			case 0:
				v20 = int64(*(*int8)(unsafe.Add(mBase, uint32(v17))))
				return v20
			case 1:
				v22 = int64(*(*int16)(unsafe.Add(mBase, uint32(v17))))
				return v22
			default:
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				return v26
			case 3:
				v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17))))
				return v24
			}
		} else {
			return base.I64_extend_i32_u(v17)
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if l3 == int32(-1) {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v30))))
			if v34&int32(1) == int32(0) {
				v44 = (v30 + l4 - int32(1)) & (int32(0) - l4)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v44))))
				v48 = v44
				v49 = v47
			} else {
				v48 = v30
				v49 = v34
			}
			v50 = l0 + v48
			if v49 == int32(1) {
				v54 = int32(18)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
				if v56 == v54 {
					v59 = v54
				} else {
					v59 = int32(2)
				}
				if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v66 = int32(6)
				} else {
					v66 = v59
				}
				v75 = v66
			} else {
				v67 = int32(1)
				if v49&v67 != 0 {
					v75 = int32(base.Ui32(v49) >> (uint(v67) % 32))
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v75 = int32(base.Ui32(v71) >> (uint(int32(2)) % 32))
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v75 + v48
			return base.I64_extend_i32_u(v50)
		} else {
			v81 = int32(1)
			v85 = (v30 + l4 - v81) & (int32(0) - l4)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v85
			v87 = l0 + v85
			v88 = F_strlen(m, v87)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v88 + v85 + v81
			return base.I64_extend_i32_u(v87)
		}
	}
}
func F_amvalidate(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = base.I32_wrap_i64(v10)
	v15 = F_SearchSysCache1(m, int32(14), v10&int64(4294967295))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)+4))
			F_ReleaseCatCache(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v26 = F_GetIndexAmRoutineByAmId(m, v22, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v22
							F_errmsg_internal(m, int32(_a_F_amvalidate_0), v8+int32(16))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_amvalidate_1), int32(209), int32(_a_F_amvalidate_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = m.T0[v28].(func(*base.Module, int32) int32)(m, v11)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							m.G0 = v8 + int32(32)
							return base.I64_extend_i32_u(v31)
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
				F_errmsg_internal(m, int32(_a_F_amvalidate_3), v8)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_amvalidate_1), int32(198), int32(_a_F_amvalidate_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
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
func F_anycompatible_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatible_in_0), int32(376), int32(_a_F_anycompatible_in_1), int32(_a_F_anycompatible_in_2), int32(_a_F_anycompatible_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anycompatiblenonarray_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblenonarray_in_0), int32(377), int32(_a_F_anycompatiblenonarray_in_1), int32(_a_F_anycompatiblenonarray_in_2), int32(_a_F_anycompatiblenonarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anyelement_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyelement_in_0), int32(374), int32(_a_F_anyelement_in_1), int32(_a_F_anyelement_in_2), int32(_a_F_anyelement_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anymultirange_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anymultirange_in_0), int32(233), int32(_a_F_anymultirange_in_1), int32(_a_F_anymultirange_in_2), int32(_a_F_anymultirange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anynonarray_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anynonarray_in_0), int32(375), int32(_a_F_anynonarray_in_1), int32(_a_F_anynonarray_in_2), int32(_a_F_anynonarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	if v18 == int32(4) {
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
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
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
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2312), int32(_a_F_apply_spooled_messages_7))
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
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v174)+40))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v174)+32))
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
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v253)+40))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v253)+32))
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
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2407), int32(_a_F_apply_spooled_messages_7))
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
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2400), int32(_a_F_apply_spooled_messages_7))
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
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2363), int32(_a_F_apply_spooled_messages_7))
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
	F_errfinish(m, int32(_a_F_apply_spooled_messages_1), int32(2279), int32(_a_F_apply_spooled_messages_14))
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
func F_arabic_UTF_8_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_SN_new_env(m, int32(32))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v7 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v3)+30)) = uint8(v7)
			*(*uint16)(unsafe.Add(mBase, uint32(v3)+28)) = uint16(v7)
		} else {
		}
		return v3
	}
}
func F_armenian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v9
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v504 < v506 {
		v572 = int32(0)
		goto L104
	} else {
		goto L105
	}
L2:
	;
	if v127 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v127 = v99
	goto L2
L4:
	;
	if v6 <= v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v127 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(1413) < v98 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v45 = v32 + int32(1)
	if v45 == v6 {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v54 = v32 + int32(2)
	if v54 != v6 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L9
L18:
	;
	v70 = v32 + int32(3)
	if v70 != v6 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v116 = v99 + v32
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v32 = v116
	goto L4
L23:
	;
	v103 = v98 - int32(1377)
	if v103 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = v130 + v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = v131
	goto L30
L28:
	;
	if v251 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v251 = v222
	goto L28
L30:
	;
	if v146 <= v155 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v251 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v162 = int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v147))))
	if base.Ui32(v164) < base.Ui32(int32(192)) {
		v221 = v164
		v222 = v162
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(1413) < v221 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v168 = v155 + int32(1)
	if v168 == v146 {
		v221 = v164
		v222 = v162
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v147))))
	v173 = v171 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v164) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v147))))
	v189 = v187 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v164) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v177 = v155 + int32(2)
	if v177 != v146 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v221 = v164<<(uint(int32(6))%32)&int32(1984) | v173
	v222 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v193))))
	v221 = v206&int32(63) | (v164<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v173<<(uint(int32(12))%32) | v189<<(uint(int32(6))%32))
	v222 = int32(4)
	goto L35
L44:
	;
	v193 = v155 + int32(3)
	if v193 != v146 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v221 = v164<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v173<<(uint(int32(6))%32) | v189
	v222 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v226 = v221 - int32(1377)
	if v226 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v226)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v232)>>(uint(v226&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v240 = v222 + v155
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240
	v155 = v240
	goto L30
L52:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v255 = v254 + v251
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v255
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v278 = v255
	goto L55
L53:
	;
	if v373 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v373 = v345
	goto L53
L55:
	;
	if v269 <= v278 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v373 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v285 = int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v270))))
	if base.Ui32(v287) < base.Ui32(int32(192)) {
		v344 = v287
		v345 = v285
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(1413) < v344 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v291 = v278 + int32(1)
	if v291 == v269 {
		v344 = v287
		v345 = v285
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+v270))))
	v296 = v294 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v287) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300+v270))))
	v312 = v310 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v287) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v300 = v278 + int32(2)
	if v300 != v269 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v344 = v287<<(uint(int32(6))%32)&int32(1984) | v296
	v345 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v316))))
	v344 = v329&int32(63) | (v287<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v296<<(uint(int32(12))%32) | v312<<(uint(int32(6))%32))
	v345 = int32(4)
	goto L60
L69:
	;
	v316 = v278 + int32(3)
	if v316 != v269 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v344 = v287<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v296<<(uint(int32(6))%32) | v312
	v345 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v362 = v345 + v278
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v362
	v278 = v362
	goto L55
L74:
	;
	v349 = v344 - int32(1377)
	if v349 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v349)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v355)>>(uint(v349&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v377 = v376 + v373
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v377
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v400 = v377
	goto L81
L79:
	;
	if v496 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v496 = v467
	goto L79
L81:
	;
	if v391 <= v400 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v496 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v407 = int32(1)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v392))))
	if base.Ui32(v409) < base.Ui32(int32(192)) {
		v466 = v409
		v467 = v407
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(1413) < v466 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v413 = v400 + int32(1)
	if v413 == v391 {
		v466 = v409
		v467 = v407
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v392))))
	v418 = v416 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v409) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+v392))))
	v434 = v432 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v409) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v422 = v400 + int32(2)
	if v422 != v391 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v466 = v409<<(uint(int32(6))%32)&int32(1984) | v418
	v467 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392+v438))))
	v466 = v451&int32(63) | (v409<<(uint(int32(18))%32)&int32(_a_F_armenian_UTF_8_stem_0) | v418<<(uint(int32(12))%32) | v434<<(uint(int32(6))%32))
	v467 = int32(4)
	goto L86
L95:
	;
	v438 = v400 + int32(3)
	if v438 != v391 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v466 = v409<<(uint(int32(12))%32)&int32(_a_F_armenian_UTF_8_stem_1) | v418<<(uint(int32(6))%32) | v434
	v467 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v471 = v466 - int32(1377)
	if v471 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v471)>>(uint(int32(3))%32)))+uint32(_c_F_armenian_UTF_8_stem[0]))))
	if int32(base.Ui32(v477)>>(uint(v471&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v485 = v467 + v400
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v485
	v400 = v485
	goto L81
L103:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v499 + v496
	goto L1
L104:
	;
	return v572
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v506
	v513 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_2), int32(57), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v527
	v533 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_3), int32(71), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L107
	} else {
		goto L112
	}
L107:
	;
	return int32(0)
L108:
	;
	if v513 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v519
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v519 < v521 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v523 = F_slice_del(m, l0)
	mBase = m.M
	if v523 < int32(0) {
		v572 = v523
		goto L104
	} else {
		goto L111
	}
L111:
	;
	goto L106
L112:
	;
	if v533 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v535
	v537 = F_slice_del(m, l0)
	mBase = m.M
	if v537 < int32(0) {
		v572 = v537
		goto L104
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	v547 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_4), int32(23), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L107
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	if v547 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v549
	v551 = F_slice_del(m, l0)
	mBase = m.M
	if v551 < int32(0) {
		v572 = v551
		goto L104
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v555
	v561 = F_find_among_b(m, l0, int32(_a_F_armenian_UTF_8_stem_5), int32(40), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L107
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v561 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v563
	v565 = F_slice_del(m, l0)
	mBase = m.M
	if v565 < int32(0) {
		v572 = v565
		goto L104
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v572 = int32(1)
	goto L104
L126:
	;
	goto L125
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+88))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+92))
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
func F_arraycontained(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14229(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v212 int32
	_ = v212
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
	var v311 int32
	_ = v311
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
	v48 = int32(_a_F_assignable_custom_variable_name_2)
	v49 = base.I32_extend8_s(v35)
	v50 = int32(54)
	goto L24
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
	goto L31
L31:
	;
	v89 = v49 & int32(255)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assignable_custom_variable_name[1])))
	if base.B2i32(v89 == v90)|int32(0) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v48
	v101 = v50
	goto L35
L33:
	;
	v119 = v48
	v121 = v50
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
	v212 = F_errdetail(m, int32(_a_F_assignable_custom_variable_name_4), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v341 = int32(1020)
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
	v311 = F_errdetail(m, int32(_a_F_assignable_custom_variable_name_5), v12+int32(-32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L51
	} else {
		goto L87
	}
L87:
	;
	v341 = int32(1037)
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
	v341 = int32(1050)
	goto L3
L94:
	;
	goto L2
}
func F_avl_sigusr2_handler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_avl_sigusr2_handler_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_avl_sigusr2_handler_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_avl_sigusr2_handler[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
