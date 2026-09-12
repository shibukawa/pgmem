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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	return v187
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
	v187 = v11
	goto L1
L7:
	;
	F_join_path_components(m, v6, v181, l0)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L68
	}
L8:
	;
	goto L14
L9:
	;
	goto L10
L10:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	v181 = v180
	goto L7
L11:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v133 != 0 {
		goto L44
	} else {
		goto L45
	}
L12:
	;
	v127 = F_strlen(m, v116)
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
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	goto L12
L17:
	;
	v101 = v96
	v102 = v97
	v103 = v98
	goto L39
L18:
	;
	if v91 == int32(0) {
		v116 = v89
		v117 = v90
		goto L16
	} else {
		goto L38
	}
L19:
	;
	v89 = l1
	v90 = v6
	v91 = v21
	goto L18
L20:
	;
	goto L21
L21:
	;
	if l1&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v58 == int32(0) {
		v116 = v55
		v117 = v56
		goto L16
	} else {
		goto L31
	}
L23:
	;
	v55 = l1
	v56 = v6
	v57 = v21
	v58 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v34 = l1
	v35 = v6
	v36 = v21
	goto L26
L26:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v38)
	if v38 == int32(0) {
		v96 = v34
		v97 = v35
		v98 = v36
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v55 = v49
	v56 = v43
	v57 = v45
	v58 = v47
	goto L22
L28:
	;
	v42 = int32(1)
	v43 = v35 + v42
	v45 = v36 - v42
	v46 = int32(0)
	v47 = base.B2i32(v45 != v46)
	v49 = v34 + v42
	if v49&int32(3) == v46 {
		v55 = v49
		v56 = v43
		v57 = v45
		v58 = v47
		goto L22
	} else {
		goto L29
	}
L29:
	;
	if v45 != 0 {
		v34 = v49
		v35 = v43
		v36 = v45
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v61 == int32(0) {
		v89 = v55
		v90 = v56
		v91 = v57
		goto L18
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(v57) < base.Ui32(int32(4)) {
		v89 = v55
		v90 = v56
		v91 = v57
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v67 = v55
	v68 = v56
	v69 = v57
	goto L34
L34:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v75 = int32(-2139062144)
	if (int32(16843008)-v72|v72)&v75 != v75 {
		v96 = v67
		v97 = v68
		v98 = v69
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v89 = v83
	v90 = v81
	v91 = v85
	goto L18
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v72
	v80 = int32(4)
	v81 = v68 + v80
	v83 = v67 + v80
	v85 = v69 - v80
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v67 = v83
		v68 = v81
		v69 = v85
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v96 = v89
	v97 = v90
	v98 = v91
	goto L17
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v105)
	if v105 == int32(0) {
		v116 = v101
		v117 = v102
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v116 = v112
	v117 = v110
	goto L16
L41:
	;
	v109 = int32(1)
	v110 = v102 + v109
	v112 = v101 + v109
	v114 = v103 - v109
	if v114 != 0 {
		v101 = v112
		v102 = v110
		v103 = v114
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v181 = v6
	goto L7
L44:
	;
	v134 = F_strlen(m, v6)
	mBase = m.M
	v137 = v134 + v6
	goto L47
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v141 = v137 - int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142 == int32(47) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v147 = v141
	goto L53
L49:
	;
	if base.Ui32(v6) < base.Ui32(v141) {
		v137 = v141
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
	goto L51
L53:
	;
	if base.Ui32(v6) < base.Ui32(v147) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v159 = v147
	goto L59
L55:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v153 != int32(47) {
		v147 = v147 - int32(1)
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	goto L57
L59:
	;
	if base.Ui32(v6) < base.Ui32(v159) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v6 == v159 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v163 = v159 - int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v164 == int32(47) {
		v159 = v163
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L63
L65:
	;
	v172 = v6 + base.B2i32(v133 == int32(47))
	goto L67
L66:
	;
	v172 = v159
	goto L67
L67:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v173)
	goto L46
L68:
	;
	F_canonicalize_path_enc(m, v6)
	mBase = m.M
	v185 = F_pstrdup(m, v6)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v187 = v185
	goto L1
}
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
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
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v578 int32
	_ = v578
	var v608 int32
	_ = v608
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	v2 = l1
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v2)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v29 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v633 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v50 = v4
	goto L4
L4:
	;
	v59 = v50 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v50<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	switch v65 {
	case 0:
		goto L11
	case 1:
		goto L9
	case 2:
		goto L10
	default:
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v59 < v608 {
		v50 = v59
		goto L4
	} else {
		goto L138
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+52)) = v578
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L16
	} else {
		goto L135
	}
L9:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v64)+36))
	if l2 != 0 {
		goto L118
	} else {
		goto L119
	}
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
	if v87 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	if v2 == int32(0) {
		v77 = int32(1)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v79 = F_table_open(m, v78, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	if l2 == int32(0) {
		v77 = v69
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v69 != int32(1) {
		v77 = v69
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v74 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v74
	v77 = v74
	goto L12
L16:
	;
	return
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+21)) = uint8(v82)
	F_sequence_close(m, v79, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	goto L6
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+52)) = int32(0)
	goto L6
L20:
	;
	goto L21
L21:
	;
	v92 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v96 <= v92 {
		v578 = v92
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v105 = v92
	v116 = v92
	v117 = v92
	v118 = v92
	goto L23
L23:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v105<<(uint(int32(2))%32))))
	if v126 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v578 = v496
	goto L7
L25:
	;
	v496 = F_lappend(m, v116, v480)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L16
	} else {
		goto L116
	}
L26:
	;
	if v165 == int32(0) {
		v480 = v126
		v491 = v117
		v492 = v118
		goto L25
	} else {
		goto L47
	}
L27:
	;
	goto L26
L28:
	;
	v127 = v126
	goto L31
L29:
	;
	goto L30
L30:
	;
	v165 = int32(0)
	goto L27
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	switch v128 - int32(15) {
	case 0:
		goto L39
	default:
		v165 = v127
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
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v162 != 0 {
		v127 = v162
		goto L31
	} else {
		goto L46
	}
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v156 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L45
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	if v151 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L44
	}
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
	if v146 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L43
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v141 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L42
	}
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v136 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L41
	}
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v131 != int32(2) {
		v165 = v127
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+28))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v161 = v135
	goto L33
L41:
	;
	v161 = v127 + int32(4)
	goto L33
L42:
	;
	v161 = v127 + int32(4)
	goto L33
L43:
	;
	v161 = v127 + int32(4)
	goto L33
L44:
	;
	v161 = v127 + int32(4)
	goto L33
L45:
	;
	v161 = v127 + int32(4)
	goto L33
L46:
	;
	goto L32
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v168 != int32(6) {
		v480 = v126
		v491 = v117
		v492 = v118
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v117 != v171 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui32(v50) < base.Ui32(v171) {
		goto L8
	} else {
		goto L52
	}
L50:
	;
	v182 = v117
	v183 = v118
	goto L51
L51:
	;
	v184 = int32(0)
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165)+8)))
	v188 = m.G0
	v190 = v188 - int32(96)
	m.G0 = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	switch v192 {
	case 0:
		goto L62
	case 1, 4, 5, 6, 9:
		v456 = v184
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
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32)-int32(4))))
	v182 = v171
	v183 = v181
	goto L51
L53:
	;
	m.G0 = v190 + int32(96)
	if v456&int32(1) != 0 {
		goto L113
	} else {
		goto L114
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L16
	} else {
		goto L110
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L16
	} else {
		goto L107
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L16
	} else {
		goto L104
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L16
	} else {
		goto L101
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L97
	}
L59:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v183)+68))
	if v238 == int32(0) {
		v333 = int32(1)
		goto L76
	} else {
		goto L77
	}
L60:
	;
	if v185 <= int32(0) {
		goto L54
	} else {
		goto L71
	}
L61:
	;
	if v185 <= int32(0) {
		goto L55
	} else {
		goto L66
	}
L62:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	v195 = F_SearchSysCache2(m, int32(7), v194, v185)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L16
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
		goto L16
	} else {
		goto L65
	}
L65:
	;
	v456 = v202
	goto L53
L66:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)+96))
	if v207 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v210 = v208
	goto L69
L68:
	;
	v210 = int32(0)
	goto L69
L69:
	;
	if v210 < v185 {
		goto L55
	} else {
		goto L70
	}
L70:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212+v185<<(uint(int32(2))%32)-int32(4))))
	v456 = base.B2i32(v218 == int32(0))
	goto L53
L71:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v183)+52))
	if v223 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v226 = v224
	goto L74
L73:
	;
	v226 = int32(0)
	goto L74
L74:
	;
	if v226 < v185 {
		goto L54
	} else {
		goto L75
	}
L75:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228+v185<<(uint(int32(2))%32)-int32(4))))
	v456 = base.B2i32(v234 == int32(0))
	goto L53
L76:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+72)))
	if base.B2i32(v334 == int32(1))&base.B2i32(v185 == v333) != 0 {
		v456 = v184
		goto L53
	} else {
		goto L92
	}
L77:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v242 <= int32(0) {
		v333 = int32(1)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v245 = int32(0)
	if v245 < v242 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v248 = v242
	goto L81
L80:
	;
	v248 = v245
	goto L81
L81:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v258 = int32(0)
	v261 = v184
	goto L83
L82:
	;
	v333 = v279 + int32(1)
	goto L76
L83:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v249+v261<<(uint(int32(2))%32))))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v279 = v258 + v278
	if base.B2i32(v185 <= v279)&base.B2i32(v258 < v185) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	if v288 != 0 {
		v456 = v184
		goto L53
	} else {
		goto L89
	}
L85:
	;
	v286 = v261 + int32(1)
	if v248 != v286 {
		v258 = v279
		v261 = v286
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	goto L82
L89:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v291 = F_get_expr_result_tupdesc(m, v289, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L16
	} else {
		goto L90
	}
L90:
	;
	if v291 == int32(0) {
		v456 = v184
		goto L53
	} else {
		goto L91
	}
L91:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v301 = int32(100)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+v295<<(uint(int32(4))%32)+(v258^int32(-1))*v301+v185*v301)+111)))
	v456 = v307
	goto L53
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+68)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v190)+64)) = v185
	F_errmsg(m, int32(68786), v190-int32(-64))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(477156), int32(3515), int32(434266))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+84)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v190)+80)) = v185
	F_errmsg(m, int32(68786), v190+int32(80))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(477156), int32(3525), int32(434266))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v385
	F_errmsg_internal(m, int32(468804), v190)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(477156), int32(3529), int32(434266))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+20)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v185
	F_errmsg_internal(m, int32(44497), v190+int32(16))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L16
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(477156), int32(3415), int32(434266))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+32)) = v185
	F_errmsg_internal(m, int32(454174), v190+int32(32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(477156), int32(3438), int32(434266))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+48)) = v185
	F_errmsg_internal(m, int32(454174), v190+int32(48))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(477156), int32(3455), int32(434266))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L16
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v472 = v184
	goto L115
L114:
	;
	v472 = v126
	goto L115
L115:
	;
	v480 = v472
	v491 = v182
	v492 = v183
	goto L25
L116:
	;
	v499 = v105 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v499 < v500 {
		v105 = v499
		v116 = v496
		v117 = v491
		v118 = v492
		goto L23
	} else {
		goto L117
	}
L117:
	;
	goto L24
L118:
	;
	v544 = int32(1)
	goto L120
L119:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v507 != 0 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	F_AcquireRewriteLocks(m, v502, v2, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L16
	} else {
		goto L134
	}
L121:
	;
	v544 = base.B2i32(v540 != int32(0))
	goto L120
L122:
	;
	goto L121
L123:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if v508 <= int32(0) {
		v540 = int32(0)
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v540 = int32(0)
	goto L122
L126:
	;
	v511 = int32(0)
	if v511 < v508 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v514 = v508
	goto L129
L128:
	;
	v514 = v511
	goto L129
L129:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v507)+12))
	v517 = int32(0)
	goto L130
L130:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v515+v517<<(uint(int32(2))%32))))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v526 == v59 {
		v540 = v525
		goto L122
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v529 = v517 + int32(1)
	if v529 != v514 {
		v517 = v529
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L6
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v171
	F_errmsg_internal(m, int32(460972), v26)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(476270), int32(255), int32(146321))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	goto L5
L139:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v699 != 0 {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	v636 = int32(0)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v637 <= v636 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v646 = v636
	goto L142
L142:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v646<<(uint(int32(2))%32))))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	F_AcquireRewriteLocks(m, v668, v2, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L16
	} else {
		goto L144
	}
L143:
	;
	goto L139
L144:
	;
	v673 = v646 + int32(1)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v673 < v674 {
		v646 = v673
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v704 = F_query_tree_walker_impl(m, l0, int32(1040), v26+int32(15), int32(3))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L16
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	m.G0 = v26 + int32(16)
	return
L149:
	;
	goto L148
}
func F_ActiveSnapshotSet(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	return base.B2i32(v2 != int32(0))
}
func F_AlterObjectOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	if l0 == int32(2613) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = int32(2995)
	goto L3
L2:
	;
	v20 = l0
	goto L3
L3:
	;
	v21 = F_get_object_attnum_oid(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v23 = F_get_object_attnum_owner(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = F_get_object_attnum_namespace(m, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v27 = F_get_object_attnum_acl(m, v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v29 = F_get_object_attnum_name(m, v20)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v32 = F_table_open(m, v20, int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v35 = F_get_catalog_object_by_oid_extended(m, v32, v21, l1, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v40 = F_heap_getattr_2(m, v35, v23, v37, v15+int32(111))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L67
	}
L15:
	;
	if v25 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v45 = F_heap_getattr_2(m, v35, v25, v42, v15+int32(111))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v48 = int32(0)
	goto L18
L18:
	;
	if l2 != v40 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v48 = v45
	goto L18
L20:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	if v173 != 0 {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	v50 = F_superuser(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_UnlockTuple(m, v32, v35+int32(4), int32(7))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L61
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+120)))
	v104 = F_palloc0(m, v101<<(uint(int32(2))%32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L45
	}
L25:
	;
	if v50 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v54 = F_has_privs_of_role(m, v53, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v54 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v29 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	F_check_can_set_role(m, v82, l2)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L39
	}
L31:
	;
	v76 = F_get_object_type(m, v20, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L37
	}
L32:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v61 = F_heap_getattr_2(m, v35, v29, v58, v15+int32(111))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	v70 = F_pg_snprintf(m, v15+int32(32), int32(64), int32(57422), v15+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v74 = v61
	goto L31
L36:
	;
	v74 = v15 + int32(32)
	goto L31
L37:
	;
	F_aclcheck_error(m, int32(2), v76, v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	if v48 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v89 = F_object_aclcheck(m, int32(2615), v48, l2, int64(512))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v89 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v94 = F_get_namespace_name(m, v48)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_aclcheck_error(m, v89, int32(36), v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L24
L45:
	;
	v106 = F_palloc0(m, v101)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v108 = F_palloc0(m, v101)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v110 = int32(1)
	v111 = v23 - v110
	*(*int32)(unsafe.Add(mBase, uint32(v104+v111<<(uint(int32(2))%32)))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v111))) = uint8(v110)
	if v27 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v143 = F_heap_modify_tuple(m, v35, v142, v104, v106, v108)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L54
	}
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v124 = F_heap_getattr_2(m, v35, v27, v121, v15+int32(111))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+111)))
	if v126 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v128 = v27 - int32(1)
	v132 = F_pg_detoast_datum(m, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v134 = F_aclnewowner(m, v132, v40, l2)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104+v128<<(uint(int32(2))%32)))) = v134
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v128))) = uint8(v138)
	goto L48
L54:
	;
	F_CatalogTupleUpdate(m, v32, v143+int32(4), v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_UnlockTuple(m, v32, v35+int32(4), int32(7))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_changeDependencyOnOwner(m, l0, l1, l2)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v104)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v106)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v108)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v174 = int32(0)
	F_RunObjectPostAlterHook(m, l0, l1, v174, v174, v174)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_sequence_close(m, v32, int32(3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v15 + int32(112)
	return
L67:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v189 + int32(4)
	F_errmsg_internal(m, int32(675046), v15)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(476138), int32(949), int32(298351))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
		v11 = *(*int32)(unsafe.Add(mBase, _consts[642]))
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
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v287 int32
	_ = v287
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
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
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v600 + v46
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v603 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
	if l1 != v24 {
		goto L153
	} else {
		goto L154
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v574))) = uint8(v569)
	goto L1
L3:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	v506 = v397
	v507 = int32(1)
	v508 = v402
	v511 = v390
	v512 = v46
	v513 = v401
	v517 = v502
	goto L137
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L10
	} else {
		goto L133
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L129
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L125
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
	v439 = m.ExcPending
	if v439 != 0 {
		goto L10
	} else {
		goto L121
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
	v35 = int32(4449520)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v42 = v40 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v45 = v24 + int32(16)
	v46 = F_ArrayGetNItems(m, v40, v45)
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
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v43 != 0 {
		goto L65
	} else {
		goto L66
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v194
	v198 = v194
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
		goto L42
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
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v68
	if v42 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = int32(1)
	if v42 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v76 = F__emscripten_memcpy_bulkmem(m, v34+int32(36), v45, v42)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v84 = int32(1)
	v86 = int32(1024)
	v88 = v61 + v84
	if v88 <= v86 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v82 = F__emscripten_memcpy_bulkmem(m, v34+int32(60), v48, v42)
	mBase = m.M
	goto L34
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	v91 = v86
	goto L37
L36:
	;
	v91 = v88
	goto L37
L37:
	;
	if v91&(v91-int32(1)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v98 = v84 << (uint(int32(32)-base.I32_clz(v91)) % 32)
	goto L40
L39:
	;
	v98 = v91
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v98
	v100 = F_palloc(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v194 = v100
	goto L21
L42:
	;
	v105 = int32(0)
	if v105 < v40 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v109 = v40
	goto L45
L44:
	;
	v109 = v105
	goto L45
L45:
	;
	v116 = v105
	goto L47
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v165 = v164 + v61
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v165 <= v166 {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	if v116 == v109 {
		goto L46
	} else {
		goto L49
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L54
	}
L49:
	;
	v132 = int32(2)
	v133 = v116 << (uint(v132) % 32)
	v135 = v116 + int32(1)
	v137 = v135 << (uint(v132) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(32)+v137)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133+v45)))
	if v139 == v141 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137+(v34+int32(56)))))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133+v48)))
	if v144 == v146 {
		v116 = v135
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
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(10805), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(475451), int32(5640), int32(198069))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
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
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v198 = v168
	goto L20
L59:
	;
	goto L60
L60:
	;
	v170 = v166 << (uint(int32(1)) % 32)
	if v165 < v170 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v172 = v170
	goto L63
L62:
	;
	v172 = v165
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v175 = F_repalloc(m, v174, v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v194 = v175
	goto L21
L65:
	;
	v221 = v43
	goto L67
L66:
	;
	v221 = (v40<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L67
L67:
	;
	if v61 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v225 + v61
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v228 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v223 = F__emscripten_memcpy_bulkmem(m, v198+v213, v24+v221, v61)
	mBase = m.M
	goto L71
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v384 != 0 {
		goto L109
	} else {
		goto L110
	}
L73:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v231 == int32(0) {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v352 = v351 + v46
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v352 <= v353 {
		goto L72
	} else {
		goto L104
	}
L76:
	;
	v234 = int32(1)
	v236 = int32(256)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v240 = v46 + v237 + v234
	if v240 <= v236 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v243 = v236
	goto L79
L78:
	;
	v243 = v240
	goto L79
L79:
	;
	if v243&(v243-int32(1)) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v250 = v234 << (uint(int32(32)-base.I32_clz(v243)) % 32)
	goto L82
L81:
	;
	v250 = v243
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v250
	v255 = base.I32_div_s(v250+int32(7), int32(8))
	v256 = F_palloc(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v259 <= int32(0) {
		goto L72
	} else {
		goto L84
	}
L84:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	if v259 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	v349 = v270 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v349)
	goto L72
L86:
	;
	v346 = v270 | int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v346)
	goto L72
L87:
	;
	v343 = v270 | int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v343)
	goto L72
L88:
	;
	v340 = v270 | int32(31)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v340)
	goto L72
L89:
	;
	v337 = v270 | int32(63)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v337)
	goto L72
L90:
	;
	v334 = v270 | int32(127)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v334)
	goto L72
L91:
	;
	v267 = v259
	v269 = v256
	v270 = v262
	goto L94
L92:
	;
	v317 = v256
	v318 = v262
	goto L93
L93:
	;
	v331 = v318 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v331)
	goto L72
L94:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v267-int32(3)) {
		goto L85
	} else {
		goto L96
	}
L95:
	;
	v317 = v309
	v318 = v310
	goto L93
L96:
	;
	v287 = v267 & int32(-2)
	if v287 == int32(2) {
		goto L86
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v267-int32(5)) {
		goto L87
	} else {
		goto L98
	}
L98:
	;
	if v287 == int32(4) {
		goto L88
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v267-int32(7)) {
		goto L89
	} else {
		goto L100
	}
L100:
	;
	if v287 == int32(6) {
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v302 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v302)
	v305 = v267 - int32(8)
	if v305 == int32(0) {
		goto L72
	} else {
		goto L102
	}
L102:
	;
	v308 = int32(1)
	v309 = v269 + v308
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v305 != v308 {
		v267 = v305
		v269 = v309
		v270 = v310
		goto L94
	} else {
		goto L103
	}
L103:
	;
	goto L95
L104:
	;
	v356 = v353 << (uint(int32(1)) % 32)
	if v352 < v356 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v358 = v356
	goto L107
L106:
	;
	v358 = v352
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v358
	v363 = base.I32_div_s(v358+int32(7), int32(8))
	v364 = F_repalloc(m, v228, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v364
	goto L72
L109:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v390 = v45 + v385<<(uint(int32(3))%32)
	goto L111
L110:
	;
	v390 = int32(0)
	goto L111
L111:
	;
	if v46 <= int32(0) {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v397 = int32(1) << (uint(v394&int32(7)) % 32)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v400 = base.I32_div_s(v394, int32(8))
	v401 = v398 + v400
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v390 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	v405 = v397
	v406 = v46
	v407 = v402
	v412 = v401
	goto L114
L114:
	;
	v420 = v405 | v407
	v421 = int32(1)
	v422 = v406 - v421
	v424 = v405 << (uint(v421) % 32)
	if v424 == int32(256) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v569 = v420
	v574 = v412
	goto L2
L116:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v412))) = uint8(v420)
	if v422 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(int32(1)) < base.Ui32(v406) {
		v405 = v424
		v406 = v422
		v407 = v420
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v430 = int32(1)
	v432 = v412 + v430
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v405 = v430
	v406 = v422
	v407 = v433
	v412 = v432
	goto L114
L120:
	;
	goto L115
L121:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(106623), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(475451), int32(5579), int32(198069))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(106592), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(475451), int32(5607), int32(198069))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L10
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
	F_errcode(m, int32(261))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v68
	F_errmsg(m, int32(642424), v20)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(475451), int32(5612), int32(198069))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L10
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(10805), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L10
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(475451), int32(5634), int32(198069))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L10
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
	if v507&v517 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v540 == int32(1) {
		goto L1
	} else {
		goto L152
	}
L139:
	;
	v526 = v506 | v508
	goto L141
L140:
	;
	v526 = v508 & (v506 ^ int32(-1))
	goto L141
L141:
	;
	v527 = int32(1)
	v528 = v512 - v527
	v530 = v506 << (uint(v527) % 32)
	if v530 == int32(256) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v513))) = uint8(v526)
	if v528 == int32(0) {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	v540 = v530
	v541 = v526
	v542 = v513
	goto L144
L144:
	;
	v544 = v507 << (uint(int32(1)) % 32)
	if v544 == int32(256) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v536 = int32(1)
	v537 = v513 + v536
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	v540 = v536
	v541 = v538
	v542 = v537
	goto L144
L146:
	;
	goto L138
L147:
	;
	if v528 == int32(0) {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	v553 = v544
	v554 = v511
	v555 = v517
	goto L149
L149:
	;
	if base.Ui32(int32(1)) < base.Ui32(v512) {
		v506 = v540
		v507 = v553
		v508 = v541
		v511 = v554
		v512 = v528
		v513 = v542
		v517 = v555
		goto L137
	} else {
		goto L151
	}
L150:
	;
	v549 = int32(1)
	v550 = v511 + v549
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v553 = v549
	v554 = v550
	v555 = v551
	goto L149
L151:
	;
	goto L146
L152:
	;
	v569 = v541
	v574 = v542
	goto L2
L153:
	;
	F_pfree(m, v24)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L10
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	m.G0 = v20 + int32(16)
	return v34
L156:
	;
	goto L155
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
	var v86 int32
	_ = v86
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
	var v190 int32
	_ = v190
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
	var v236 int32
	_ = v236
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
	v86 = v17
	goto L11
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v86<<(uint(int32(2))%32))))
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
		v236 = l6
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
		v236 = v109
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
	v190 = v113
	goto L21
L21:
	;
	if l5&v60 == int32(0) {
		v236 = v109
		goto L13
	} else {
		goto L25
	}
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
	v190 = v172
	goto L21
L24:
	;
	goto L23
L25:
	;
	v206 = int32(1)
	v207 = v190 << (uint(v206) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v211 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+v207))))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v211<<(uint(v206)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v207))) = uint16(v217)
	v236 = v109
	goto L13
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
	F_addFkConstraint(m, v31+int32(12), int32(0), v254, l0, l1, v100, v247, l4, l5, v236, l7, l8, l9, l10, l11, l12, int32(1), l15)
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
	F_addFkRecurseReferenced(m, l0, l1, v100, v247, v258, l5, v236, l7, l8, l9, l10, l11, l12, v65, v64, l15)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_sequence_close(m, v100, int32(0))
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
	F_pfree(m, v236)
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
	v269 = v86 + int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v269 < v270 {
		v86 = v269
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
	F_errmsg_internal(m, int32(174925), v31)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(475396), int32(10968), int32(445575))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v67 int32
	_ = v67
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		if v20 != 0 {
			v25 = F__emscripten_memcpy_bulkmem(m, v21, v24, v20)
			mBase = m.M
			v26 = v25
		} else {
			v26 = v21
		}
		if v18 <= int32(0) {
		} else {
			v30 = v18 & int32(7)
			v31 = int32(0)
			if base.Ui32(int32(8)) <= base.Ui32(v18) {
				v40 = v31
				v45 = v4
				for {
					v49 = v26 + v40<<(uint(int32(5))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+116)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+148)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+180)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+212)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v49)+244)) = l2
					v58 = int32(8)
					v59 = v40 + v58
					v61 = v45 + v58
					if v61 != v18&int32(2147483640) {
						v40 = v59
						v45 = v61
						continue
					} else {
						break
					}
					break
				}
				v67 = v59
			} else {
				v67 = v31
			}
			if v30 == int32(0) {
			} else {
				v80 = v67
				v84 = v4
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v26+v80<<(uint(int32(5))%32))+20)) = l2
					v91 = int32(1)
					v94 = v84 + v91
					if v94 != v30 {
						v80 = v80 + v91
						v84 = v94
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
			v110 = F_makeAlias(m, l1, v15)
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
				*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v26
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v5 = l4
	if v5 != 0 {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v38 = v35 + l3<<(uint(int32(3))%32)
		v39 = int32(1)
		v40 = l5 - v39
		*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
		v43 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
		v45 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
		if v43 <= v45+v39 {
			if v43 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[1001])) = int32(2)
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
						v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
						v76 = v74 << (uint(int32(3)) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
						if v68 != 0 {
							v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
							mBase = m.M
						} else {
						}
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
						v89 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						if l6 != 0 {
							v98 = int32(4096)
						} else {
							v98 = v89
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v103 = v74 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1001])) = v43 << (uint(int32(1)) % 32)
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
						v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
						v76 = v74 << (uint(int32(3)) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
						if v68 != 0 {
							v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
							mBase = m.M
						} else {
						}
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
						v89 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						if l6 != 0 {
							v98 = int32(4096)
						} else {
							v98 = v89
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v103 = v74 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
				v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
				v76 = v74 << (uint(int32(3)) % 32)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
				if v68 != 0 {
					v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
					mBase = m.M
				} else {
				}
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
				v89 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				if l6 != 0 {
					v98 = int32(4096)
				} else {
					v98 = v89
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v103 = v74 + int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
				*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
				return
			}
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[1002])) = v10
		*(*int32)(unsafe.Add(mBase, _consts[1001])) = v10
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if base.Ui32(l3) < base.Ui32(v15) {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v38 = v35 + l3<<(uint(int32(3))%32)
			v39 = int32(1)
			v40 = l5 - v39
			*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v40)
			v43 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
			v45 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
			if v43 <= v45+v39 {
				if v43 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[1001])) = int32(2)
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
							v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							if v68 != 0 {
								v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
								mBase = m.M
							} else {
							}
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
							v89 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v98 = int32(4096)
							} else {
								v98 = v89
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v103 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[1001])) = v43 << (uint(int32(1)) % 32)
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
							v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							if v68 != 0 {
								v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
								mBase = m.M
							} else {
							}
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
							v89 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v98 = int32(4096)
							} else {
								v98 = v89
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v103 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
					v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
					v76 = v74 << (uint(int32(3)) % 32)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
					if v68 != 0 {
						v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
						mBase = m.M
					} else {
					}
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
					v89 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					if l6 != 0 {
						v98 = int32(4096)
					} else {
						v98 = v89
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v103 = v74 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
					*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
					v43 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
					v45 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
					if v43 <= v45+v39 {
						if v43 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _consts[1001])) = int32(2)
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
									v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
									if v68 != 0 {
										v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
										mBase = m.M
									} else {
									}
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
									v89 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v98 = int32(4096)
									} else {
										v98 = v89
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v103 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1001])) = v43 << (uint(int32(1)) % 32)
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
									v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
									if v68 != 0 {
										v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
										mBase = m.M
									} else {
									}
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
									v89 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v98 = int32(4096)
									} else {
										v98 = v89
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v103 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
							v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							if v68 != 0 {
								v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
								mBase = m.M
							} else {
							}
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
							v89 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v98 = int32(4096)
							} else {
								v98 = v89
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v103 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
					v43 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
					v45 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
					if v43 <= v45+v39 {
						if v43 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _consts[1001])) = int32(2)
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
									v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
									if v68 != 0 {
										v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
										mBase = m.M
									} else {
									}
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
									v89 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v98 = int32(4096)
									} else {
										v98 = v89
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v103 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1001])) = v43 << (uint(int32(1)) % 32)
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
									v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
									v76 = v74 << (uint(int32(3)) % 32)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
									if v68 != 0 {
										v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
										mBase = m.M
									} else {
									}
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
									v89 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									if l6 != 0 {
										v98 = int32(4096)
									} else {
										v98 = v89
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v103 = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
							v74 = *(*int32)(unsafe.Add(mBase, _consts[1002]))
							v76 = v74 << (uint(int32(3)) % 32)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v76+v77)+4)) = v71
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76)+4))
							if v68 != 0 {
								v83 = F__emscripten_memcpy_bulkmem(m, v82, l1, v68)
								mBase = m.M
							} else {
							}
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4))
							v89 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v87+v68))) = uint8(v89)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(v91+v76))) = uint16(v5)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							if l6 != 0 {
								v98 = int32(4096)
							} else {
								v98 = v89
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v94+v76)+2)) = uint16(v98)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
							v103 = v74 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[1002])) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v100+v103<<(uint(int32(3))%32))+4)) = int32(0)
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
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	v9 = int32(0)
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v110 = F_palloc0(m, int32(24))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L27
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v13 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v80 = F_palloc0(m, int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L23
	}
L5:
	;
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v65 = v9
	goto L5
L7:
	;
	goto L8
L8:
	;
	v24 = v9
	v26 = v9
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	if l7 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v65 = v49
	goto L5
L11:
	;
	v52 = v24 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v52 < v53 {
		v24 = v52
		v26 = v49
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v41 = F_copyObjectImpl(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v35 != 0 {
		v40 = v35
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v37 == int32(0) {
		v49 = v26
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v40 = v37
	goto L12
L18:
	;
	return
L19:
	;
	v43 = F_lappend(m, v26, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
	v47 = v45 | v46
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v47)
	v49 = v43
	goto L11
L21:
	;
	goto L10
L22:
	;
	goto L4
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(105)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v88 = F_pstrdup(m, v85+int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v88
	v96 = int32(1)
	v100 = F_makeConst(m, int32(16), int32(-1), v90, v96, v90, v90, v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+20)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v100
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v106 = F_lappend(m, v105, v80)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v106
	return
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = int32(105)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v118 = F_pstrdup(m, v115+int32(4))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+20)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v118
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v125 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+16)) = v134
	F_ChangeVarNodes(m, v134, int32(1), l1)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L34
	}
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v134 = v129
	goto L29
L31:
	;
	goto L32
L32:
	;
	v132 = F_makeBoolExpr(m, int32(1), v65, int32(-1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v134 = v132
	goto L29
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v141 = F_list_append_unique(m, v140, v110)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v141
	if l4 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v146 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v152 = int32(0)
	goto L39
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v152<<(uint(int32(2))%32))))
	if l7 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L36
L41:
	;
	v210 = v152 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v210 < v211 {
		v152 = v210
		goto L39
	} else {
		goto L54
	}
L42:
	;
	v174 = F_copyObjectImpl(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L18
	} else {
		goto L48
	}
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	if v168 != 0 {
		v173 = v168
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v170 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v173 = v170
	goto L42
L48:
	;
	F_ChangeVarNodes(m, v174, int32(1), l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v180 = F_palloc0(m, int32(24))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = int32(105)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v188 = F_pstrdup(m, v185+int32(4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v188
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v192 = F_pstrdup(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+20)) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v192
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v199 = F_list_append_unique(m, v198, v180)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v199
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+24)))
	v204 = v202 | v203
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v204)
	goto L41
L54:
	;
	goto L40
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
							F_errmsg_internal(m, int32(51158), v8+int32(16))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(478485), int32(196), int32(342440))
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
				F_errmsg_internal(m, int32(40426), v8)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(478485), int32(185), int32(342440))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(374989)
			F_errmsg(m, int32(183897), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(475049), int32(376), int32(268037))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anycompatiblenonarray_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(22724)
			F_errmsg(m, int32(183897), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(475049), int32(377), int32(267518))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anyelement_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(90936)
			F_errmsg(m, int32(183897), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(475049), int32(374), int32(267611))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anymultirange_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(384839)
			F_errmsg(m, int32(183897), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(475049), int32(233), int32(268066))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anynonarray_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(22712)
			F_errmsg(m, int32(183897), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(475049), int32(375), int32(267503))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int64
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v9 = m.G0
	v11 = v9 - int32(2224)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[626]))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v53 < int32(0) {
		goto L14
	} else {
		goto L15
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[636]))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	if v23 == int64(0) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L4
L6:
	;
	if l2 != v23 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, _consts[660])) = l2
	v31 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v31 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int64)(unsafe.Add(mBase, _consts[660]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+116)) = uint32(v36)
	v39 = int64(base.Ui64(v36) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+112)) = uint32(v39)
	F_errmsg(m, int32(495750), v11+int32(112))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(476327), int32(4924), int32(161052))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	goto L17
L14:
	;
	v57 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[661])) = v57
	goto L16
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	if base.B2i32(v61 == int32(2)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_PushActiveSnapshot(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v76
	v79 = *(*int32)(unsafe.Add(mBase, _consts[626]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l1
	v84 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	v91 = F_pg_snprintf(m, v11+int32(160), int32(1024), int32(161103), v11+int32(96))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v95 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v95 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(160)
	F_errmsg_internal(m, int32(679173), v11+int32(80))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v110 = int32(4449572)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	v114 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v114
	v118 = int32(0)
	v120 = F_BufFileOpenFileSet(m, l0, v11+int32(160), v118, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(476327), int32(2044), int32(161749))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v111
	*(*int32)(unsafe.Add(mBase, _consts[662])) = v120
	v127 = F_palloc(m, int32(8192))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, _consts[663])) = l2
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v84
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[664])) = uint8(v134)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	F_PopActiveSnapshot(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v148 = v127
	v149 = int32(0)
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L8
	} else {
		goto L95
	}
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v153 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L92
	}
L39:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	v162 = F_BufFileReadMaybeEOF(m, v157, v11+int32(124), int32(4), int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L45
	}
L42:
	;
	goto L41
L43:
	;
	goto L38
L44:
	;
	v314 = base.I32_rem_s(v208, int32(1000))
	if v314 != 0 {
		v148 = v167
		v149 = v208
		goto L37
	} else {
		goto L87
	}
L45:
	;
	if v162 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	if v164 <= int32(0) {
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v285 = v149
	goto L48
L48:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	if v288 != 0 {
		goto L77
	} else {
		goto L78
	}
L49:
	;
	v167 = F_repalloc(m, v148, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	F_BufFileReadExact(m, v170, v167, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(156)))) = v180
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v175)+32))
	v183 = int64(*(*int32)(unsafe.Add(mBase, uint32(v175)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(144)))) = v182 + v183
	goto L52
L52:
	;
	v186 = int32(4449520)
	v187 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v190 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v167
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v195
	F_apply_dispatch(m, v11+int32(128))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	F_MemoryContextReset(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v187
	v208 = v149 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	if v210 != 0 {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v11)+144))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v214 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v214 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	goto L60
L57:
	;
	v218 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[661])) = v218
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	if base.B2i32(v222 == int32(2)) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v231 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L66
	}
L64:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	F_PushActiveSnapshot(m, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v237
	v240 = *(*int32)(unsafe.Add(mBase, _consts[626]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = l1
	v250 = F_pg_snprintf(m, v11+int32(1200), int32(1024), int32(161103), v11+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v254 = int32(0)
	v256 = F_BufFileOpenFileSet(m, l0, v11+int32(1200), v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v261 = F_BufFileSeek(m, v256, int32(0), int64(0), int32(2))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(1196)))) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v256)+32))
	v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(1184)))) = v269 + v270
	goto L71
L71:
	;
	F_BufFileClose(m, v256)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1196))
	if v279 != v212 {
		goto L36
	} else {
		goto L75
	}
L75:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v11)+1184))
	if v281 != v211 {
		goto L36
	} else {
		goto L76
	}
L76:
	;
	v285 = v208
	goto L48
L77:
	;
	F_BufFileClose(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v296 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[662])) = int32(0)
	goto L79
L81:
	;
	if v296 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v11 + int32(160)
	F_errmsg_internal(m, int32(679241), v11)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	m.G0 = v11 + int32(2224)
	return
L85:
	;
	F_errfinish(m, int32(476327), int32(2139), int32(161749))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v317 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	if v317 == int32(0) {
		v148 = v167
		v149 = v208
		goto L37
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v11 + int32(160)
	F_errmsg_internal(m, int32(679206), v11-int32(-64))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(476327), int32(2132), int32(161749))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v148 = v167
	v149 = v208
	goto L37
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v11 + int32(160)
	F_errmsg_internal(m, int32(678309), v11+int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(476327), int32(2095), int32(161749))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(1200)
	F_errmsg_internal(m, int32(678240), v11+int32(32))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(476327), int32(2011), int32(388174))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
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
	v99 = v84&int32(63) | (v42<<(uint(int32(18))%32)&int32(1835008) | v51<<(uint(int32(12))%32) | v67<<(uint(int32(6))%32))
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
	v99 = v42<<(uint(int32(12))%32)&int32(61440) | v51<<(uint(int32(6))%32) | v67
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
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(3))%32)))+uint32(_consts[1318]))))
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
	v223 = v208&int32(63) | (v166<<(uint(int32(18))%32)&int32(1835008) | v175<<(uint(int32(12))%32) | v191<<(uint(int32(6))%32))
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
	v223 = v166<<(uint(int32(12))%32)&int32(61440) | v175<<(uint(int32(6))%32) | v191
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
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v228)>>(uint(int32(3))%32)))+uint32(_consts[1318]))))
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
	v346 = v331&int32(63) | (v289<<(uint(int32(18))%32)&int32(1835008) | v298<<(uint(int32(12))%32) | v314<<(uint(int32(6))%32))
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
	v346 = v289<<(uint(int32(12))%32)&int32(61440) | v298<<(uint(int32(6))%32) | v314
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
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_consts[1318]))))
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
	v468 = v453&int32(63) | (v411<<(uint(int32(18))%32)&int32(1835008) | v420<<(uint(int32(12))%32) | v436<<(uint(int32(6))%32))
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
	v468 = v411<<(uint(int32(12))%32)&int32(61440) | v420<<(uint(int32(6))%32) | v436
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
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v473)>>(uint(int32(3))%32)))+uint32(_consts[1318]))))
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
	v516 = F_find_among_b(m, l0, int32(4182048), int32(57))
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
	v537 = F_find_among_b(m, l0, int32(4183200), int32(71))
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
	v551 = F_find_among_b(m, l0, int32(4184624), int32(23))
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
	v565 = F_find_among_b(m, l0, int32(4185088), int32(40))
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DatumGetAnyArrayP(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_DatumGetAnyArrayP(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v6, v11, v13, int32(1), v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					if v27 == int32(-1) {
						return v18
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v11 == v30 {
							return v18
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v18
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v6 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						if v27 == int32(-1) {
							return v18
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v11 == v30 {
								return v18
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									return v18
								}
							}
						}
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							if v27 == int32(-1) {
								return v18
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v11 == v30 {
									return v18
								} else {
									F_pfree(m, v11)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										return v18
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
	*(*int32)(unsafe.Add(mBase, _consts[418])) = int32(7)
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[419])) = uint8(base.B2i32(v3 != v8))
	v14 = int32(1)
	v15 = int32(base.Ui32(v3)>>(uint(int32(2))%32)) & v14
	*(*uint8)(unsafe.Add(mBase, _consts[420])) = uint8(v15)
	v21 = int32(base.Ui32(v3)>>(uint(v14)%32)) & v14
	*(*uint8)(unsafe.Add(mBase, _consts[421])) = uint8(v21)
	*(*uint8)(unsafe.Add(mBase, _consts[422])) = uint8(v8)
	return
}
func F_assign_io_combine_limit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	if v5 < l0 {
		v7 = v5
	} else {
		v7 = l0
	}
	*(*int32)(unsafe.Add(mBase, _consts[428])) = v7
	return
}
func F_assign_synchronous_standby_names(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[687])) = l1
	return
}
func F_assignable_custom_variable_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var __phi34 int32
	_ = __phi34
	var v35 int32
	_ = v35
	var __phi35 int32
	_ = __phi35
	var v37 int32
	_ = v37
	var __phi37 int32
	_ = __phi37
	var v41 int32
	_ = v41
	var __phi41 int32
	_ = __phi41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
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
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v419 int32
	_ = v419
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = int32(46)
	v19 = F___strchrnul(m, l0, v18)
	mBase = m.M
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == v18 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v15 - int32(-64)
	return v419
L2:
	;
	v419 = int32(0)
	goto L1
L3:
	;
	F_errfinish(m, int32(480467), v399, int32(364481))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L55
	} else {
		goto L116
	}
L4:
	;
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v25 = v19
	goto L7
L6:
	;
	v25 = v4
	goto L7
L7:
	;
	goto L4
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
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
		goto L111
	}
L11:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
	if v217 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L12:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L54
	}
L13:
	;
	v29 = v25 - l0
	__phi34 = v26
	__phi35 = l0
	__phi37 = int32(1)
	__phi41 = v4
	v34 = __phi34
	v35 = __phi35
	v37 = __phi37
	v41 = __phi41
	goto L14
L14:
	;
	if v34 == int32(46) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if base.B2i32(v34&int32(255) != int32(46))&v170 != 0 {
		goto L11
	} else {
		goto L53
	}
L16:
	;
	v175 = v35 + int32(1)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v176 != 0 {
		__phi34 = v176
		__phi35 = v175
		__phi37 = base.B2i32(v34 == int32(46))
		__phi41 = v170
		v34 = __phi34
		v35 = __phi35
		v37 = __phi37
		v41 = __phi41
		goto L14
	} else {
		goto L52
	}
L17:
	;
	v45 = int32(1)
	if v37&v45 == int32(0) {
		v170 = v45
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v51 = base.I32_extend8_s(v34)
	goto L25
L20:
	;
	goto L12
L21:
	;
	if v51 < int32(0) {
		v170 = v41
		goto L16
	} else {
		goto L47
	}
L22:
	;
	v155 = int32(0)
	goto L21
L23:
	;
	v133 = v126
	v135 = v128
	goto L41
L24:
	;
	if base.B2i32(v73 != v74) == int32(0) {
		goto L22
	} else {
		goto L32
	}
L25:
	;
	goto L26
L26:
	;
	v65 = int32(487762)
	v67 = int32(54)
	goto L27
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v70 == v51&int32(255) {
		v126 = v65
		v128 = v67
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v72 = int32(1)
	v73 = v67 - v72
	v74 = int32(0)
	v77 = v65 + v72
	if v77&int32(3) == v74 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	if v73 != 0 {
		v65 = v77
		v67 = v73
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v89 == v51&int32(255) {
		v119 = v77
		v121 = v73
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v121 == int32(0) {
		goto L22
	} else {
		goto L40
	}
L34:
	;
	if base.Ui32(v73) < base.Ui32(int32(4)) {
		v119 = v77
		v121 = v73
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v99 = v77
	v101 = v73
	goto L36
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v106 = v105 ^ v51&int32(255)*int32(16843009)
	v109 = int32(-2139062144)
	if (int32(16843008)-v106|v106)&v109 != v109 {
		v126 = v99
		v128 = v101
		goto L23
	} else {
		goto L38
	}
L37:
	;
	v119 = v114
	v121 = v116
	goto L33
L38:
	;
	v113 = int32(4)
	v114 = v99 + v113
	v116 = v101 - v113
	if base.Ui32(int32(3)) < base.Ui32(v116) {
		v99 = v114
		v101 = v116
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v126 = v119
	v128 = v121
	goto L23
L41:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v51&int32(255) == v138 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L22
L43:
	;
	v155 = v133
	goto L21
L44:
	;
	goto L45
L45:
	;
	v140 = int32(1)
	v143 = v135 - v140
	if v143 != 0 {
		v133 = v133 + v140
		v135 = v143
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	if v155 != 0 {
		v170 = v41
		goto L16
	} else {
		goto L48
	}
L48:
	;
	if v37&int32(1) != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(63)) < base.Ui32(v34) {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v51))%64)&int64(287948969894477825) == int64(0) {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v170 = v41
	goto L16
L52:
	;
	goto L15
L53:
	;
	goto L12
L54:
	;
	v194 = int32(0)
	v196 = F_errstart(m, l2, v194)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	if v196 == int32(0) {
		v419 = v194
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(677120), v13+int32(-48))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	F_errdetail(m, int32(549168), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v399 = int32(1139)
	goto L3
L61:
	;
	v419 = int32(1)
	goto L1
L62:
	;
	goto L63
L63:
	;
	v221 = int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v222 <= int32(0) {
		v419 = v221
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v225 = int32(0)
	if v225 < v222 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v228 = v222
	goto L67
L66:
	;
	v228 = v225
	goto L67
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v235 = int32(0)
	goto L68
L68:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v229+v235<<(uint(int32(2))%32))))
	if v246&int32(3) == int32(0) {
		v270 = v246
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v419 = v221
	goto L1
L70:
	;
	v372 = v235 + int32(1)
	if v372 != v228 {
		v235 = v372
		goto L68
	} else {
		goto L110
	}
L71:
	;
	if v303 != v29 {
		goto L70
	} else {
		goto L88
	}
L72:
	;
	v303 = v295 - v246
	goto L71
L73:
	;
	v274 = v270
	goto L82
L74:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v254 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v303 = int32(0)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v259 = v246
	goto L78
L78:
	;
	v263 = v259 + int32(1)
	if v263&int32(3) == int32(0) {
		v270 = v263
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v295 = v263
	goto L72
L80:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v268 != 0 {
		v259 = v263
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v283 = int32(-2139062144)
	if (int32(16843008)-v280|v280)&v283 == v283 {
		v274 = v274 + int32(4)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v289 = v274
	goto L85
L84:
	;
	goto L83
L85:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v293 != 0 {
		v289 = v289 + int32(1)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v295 = v289
	goto L72
L87:
	;
	goto L86
L88:
	;
	if v29 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v348 != 0 {
		goto L70
	} else {
		goto L103
	}
L90:
	;
	v348 = int32(0)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v310 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v311 = l0
	v312 = v246
	v313 = v29
	v314 = v310
	goto L97
L94:
	;
	v336 = v246
	v340 = int32(0)
	goto L95
L95:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	v348 = v340 - v341
	goto L89
L96:
	;
	v336 = v331
	v340 = v333
	goto L95
L97:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	if v314 != v316 {
		v331 = v312
		v333 = v314
		goto L96
	} else {
		goto L99
	}
L98:
	;
	v331 = v325
	v333 = int32(0)
	goto L96
L99:
	;
	if v316 == int32(0) {
		v331 = v312
		v333 = v314
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v321 = v313 - int32(1)
	if v321 == int32(0) {
		v331 = v312
		v333 = v314
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v324 = int32(1)
	v325 = v312 + v324
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v326 != 0 {
		v311 = v311 + v324
		v312 = v325
		v313 = v321
		v314 = v326
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v349 = int32(0)
	v351 = F_errstart(m, l2, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L55
	} else {
		goto L105
	}
L105:
	;
	if v351 == int32(0) {
		v419 = v349
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L55
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
	F_errmsg(m, int32(677120), v13+int32(-16))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L55
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v246
	F_errdetail(m, int32(540459), v13+int32(-32))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L55
	} else {
		goto L109
	}
L109:
	;
	v399 = int32(1156)
	goto L3
L110:
	;
	goto L69
L111:
	;
	v375 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L55
	} else {
		goto L112
	}
L112:
	;
	if v375 == int32(0) {
		v419 = v4
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L55
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(663274), v15)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L55
	} else {
		goto L115
	}
L115:
	;
	v399 = int32(1169)
	goto L3
L116:
	;
	goto L2
}
func F_avl_sigusr2_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[514])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
