package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendStringInfoStringQuoted(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var __phi111 int32
	_ = __phi111
	var v114 int32
	_ = v114
	var __phi114 int32
	_ = __phi114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1&int32(3) == v4 {
		v35 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v72 = base.B2i32(l2 < v68) & base.B2i32(int32(0) <= l2)
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v68 = v60 - l1
	goto L1
L3:
	;
	v39 = v35
	goto L12
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v68 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v24 = l1
	goto L8
L8:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v60 = v28
	goto L2
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v54 = v39
	goto L15
L14:
	;
	goto L13
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v60 = v54
	goto L2
L17:
	;
	goto L16
L18:
	;
	v73 = F_pg_mbcliplen(m, l1, v68, l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = l1
	v78 = v4
	goto L20
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v79 <= v80+int32(1) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	return
L22:
	;
	v75 = F_pnstrdup(m, l1, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = v75
	v78 = v75
	goto L20
L24:
	;
	v100 = int32(39)
	v101 = F___strchrnul(m, v77, v100)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v103 == v100 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v80))) = uint8(v89)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = v91 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v95+v93))) = uint8(v97)
	goto L24
L28:
	;
	goto L24
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v134
	if v72 != 0 {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	if v107 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v107 = v101
	goto L33
L32:
	;
	v107 = int32(0)
	goto L33
L33:
	;
	goto L30
L34:
	;
	v134 = v77
	goto L29
L35:
	;
	goto L36
L36:
	;
	__phi111 = v77
	__phi114 = v107
	v111 = __phi111
	v114 = __phi114
	goto L37
L37:
	;
	F_appendBinaryStringInfoNT(m, l0, v111, v114-v111+int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L39
	}
L38:
	;
	v134 = v114
	goto L29
L39:
	;
	v124 = int32(39)
	v125 = F___strchrnul(m, v114+int32(1), v124)
	mBase = m.M
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v127 == v124 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v131 != 0 {
		__phi111 = v114
		__phi114 = v131
		v111 = __phi111
		v114 = __phi114
		goto L37
	} else {
		goto L44
	}
L41:
	;
	v131 = v125
	goto L43
L42:
	;
	v131 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L38
L45:
	;
	v142 = int32(665947)
	goto L47
L46:
	;
	v142 = int32(663268)
	goto L47
L47:
	;
	F_appendStringInfo(m, l0, v142, v10)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	if v78 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_pfree(m, v78)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	m.G0 = v10 + int32(16)
	return
L52:
	;
	goto L51
}
func F_string_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = l1 - int32(1)
	if base.Ui32(v59) < base.Ui32(v61) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v59 = v51 - l0
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = l0
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	v63 = v59
	goto L20
L19:
	;
	v63 = v61
	goto L20
L20:
	;
	v69 = v63 - int32(1636608432)
	if l0&int32(3) != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	return v323 ^ v315 - base.I32_rotl(v323, int32(24))
L22:
	;
	v301 = int32(14)
	v303 = v297 ^ v298 - base.I32_rotl(v297, v301)
	v307 = v303 ^ v296 - base.I32_rotl(v303, int32(11))
	v311 = v307 ^ v297 - base.I32_rotl(v307, int32(25))
	v315 = v311 ^ v303 - base.I32_rotl(v311, int32(16))
	v319 = v315 ^ v307 - base.I32_rotl(v315, int32(4))
	v323 = v319 ^ v311 - base.I32_rotl(v319, v301)
	goto L21
L23:
	;
	switch v227 - int32(1) {
	case 0:
		v289 = v228
		v290 = v229
		v291 = v230
		goto L50
	case 1:
		v282 = v228
		v283 = v229
		v284 = v230
		goto L51
	case 2:
		v275 = v228
		v276 = v229
		v277 = v230
		goto L52
	case 3:
		v269 = v229
		v270 = v230
		goto L53
	case 4:
		v265 = v229
		v266 = v230
		goto L54
	case 5:
		v259 = v229
		v260 = v230
		goto L55
	case 6:
		v253 = v229
		v254 = v230
		goto L56
	case 7:
		v248 = v230
		goto L57
	case 8:
		v243 = v230
		goto L58
	case 9:
		v238 = v230
		goto L59
	case 10:
		goto L60
	default:
		v296 = v228
		v297 = v229
		v298 = v230
		goto L22
	}
L24:
	;
	v178 = l0
	v179 = v63
	v180 = v69
	v181 = v69
	v182 = v69
	goto L47
L25:
	;
	if base.Ui32(int32(11)) < base.Ui32(v63) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(v63) < base.Ui32(int32(12)) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v226 = l0
	v227 = v63
	v228 = v69
	v229 = v69
	v230 = v69
	goto L23
L29:
	;
	switch v125 - int32(1) {
	case 0:
		v175 = v126
		goto L36
	case 1:
		v170 = v126
		goto L37
	case 2:
		goto L38
	case 3:
		v163 = v127
		goto L39
	case 4:
		v160 = v127
		goto L40
	case 5:
		v155 = v127
		goto L41
	case 6:
		goto L42
	case 7:
		v146 = v128
		goto L43
	case 8:
		v141 = v128
		goto L44
	case 9:
		v136 = v128
		goto L45
	case 10:
		goto L46
	default:
		v296 = v126
		v297 = v127
		v298 = v128
		goto L22
	}
L30:
	;
	v124 = l0
	v125 = v63
	v126 = v69
	v127 = v69
	v128 = v69
	goto L29
L31:
	;
	goto L32
L32:
	;
	v76 = l0
	v77 = v63
	v78 = v69
	v79 = v69
	v80 = v69
	goto L33
L33:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v83 = v82 + v79
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v87 = v86 + v80
	v89 = int32(4)
	v91 = v84 + v78 - v87 ^ base.I32_rotl(v87, v89)
	v95 = v83 - v91 ^ base.I32_rotl(v91, int32(6))
	v96 = v87 + v83
	v97 = v91 + v96
	v98 = v95 + v97
	v102 = v96 - v95 ^ base.I32_rotl(v95, int32(8))
	v106 = v97 - v102 ^ base.I32_rotl(v102, int32(16))
	v110 = v98 - v106 ^ base.I32_rotl(v106, int32(19))
	v111 = v102 + v98
	v112 = v106 + v111
	v113 = v110 + v112
	v117 = v111 - v110 ^ base.I32_rotl(v110, v89)
	v118 = int32(12)
	v119 = v76 + v118
	v121 = v77 - v118
	if base.Ui32(int32(11)) < base.Ui32(v121) {
		v76 = v119
		v77 = v121
		v78 = v112
		v79 = v113
		v80 = v117
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v124 = v119
	v125 = v121
	v126 = v112
	v127 = v113
	v128 = v117
	goto L29
L35:
	;
	goto L34
L36:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v296 = v175 + v176
	v297 = v127
	v298 = v128
	goto L22
L37:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	v175 = v171<<(uint(int32(8))%32) + v170
	goto L36
L38:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+2)))
	v170 = v166<<(uint(int32(16))%32) + v126
	goto L37
L39:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v296 = v164 + v126
	v297 = v163
	v298 = v128
	goto L22
L40:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+4)))
	v163 = v160 + v161
	goto L39
L41:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+5)))
	v160 = v156<<(uint(int32(8))%32) + v155
	goto L40
L42:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+6)))
	v155 = v151<<(uint(int32(16))%32) + v127
	goto L41
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v296 = v147 + v126
	v297 = v149 + v127
	v298 = v146
	goto L22
L44:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+8)))
	v146 = v142<<(uint(int32(8))%32) + v141
	goto L43
L45:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+9)))
	v141 = v137<<(uint(int32(16))%32) + v136
	goto L44
L46:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+10)))
	v136 = v132<<(uint(int32(24))%32) + v128
	goto L45
L47:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v185 = v184 + v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v189 = v188 + v182
	v191 = int32(4)
	v193 = v186 + v180 - v189 ^ base.I32_rotl(v189, v191)
	v197 = v185 - v193 ^ base.I32_rotl(v193, int32(6))
	v198 = v189 + v185
	v199 = v193 + v198
	v200 = v197 + v199
	v204 = v198 - v197 ^ base.I32_rotl(v197, int32(8))
	v208 = v199 - v204 ^ base.I32_rotl(v204, int32(16))
	v212 = v200 - v208 ^ base.I32_rotl(v208, int32(19))
	v213 = v204 + v200
	v214 = v208 + v213
	v215 = v212 + v214
	v219 = v213 - v212 ^ base.I32_rotl(v212, v191)
	v220 = int32(12)
	v221 = v178 + v220
	v223 = v179 - v220
	if base.Ui32(int32(11)) < base.Ui32(v223) {
		v178 = v221
		v179 = v223
		v180 = v214
		v181 = v215
		v182 = v219
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v226 = v221
	v227 = v223
	v228 = v214
	v229 = v215
	v230 = v219
	goto L23
L49:
	;
	goto L48
L50:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v296 = v289 + v292
	v297 = v290
	v298 = v291
	goto L22
L51:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	v289 = v285<<(uint(int32(8))%32) + v282
	v290 = v283
	v291 = v284
	goto L50
L52:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+2)))
	v282 = v278<<(uint(int32(16))%32) + v275
	v283 = v276
	v284 = v277
	goto L51
L53:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+3)))
	v275 = v271<<(uint(int32(24))%32) + v228
	v276 = v269
	v277 = v270
	goto L52
L54:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+4)))
	v269 = v265 + v267
	v270 = v266
	goto L53
L55:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+5)))
	v265 = v261<<(uint(int32(8))%32) + v259
	v266 = v260
	goto L54
L56:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+6)))
	v259 = v255<<(uint(int32(16))%32) + v253
	v260 = v254
	goto L55
L57:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+7)))
	v253 = v249<<(uint(int32(24))%32) + v229
	v254 = v248
	goto L56
L58:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+8)))
	v248 = v244<<(uint(int32(8))%32) + v243
	goto L57
L59:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+9)))
	v243 = v239<<(uint(int32(16))%32) + v238
	goto L58
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+10)))
	v238 = v234<<(uint(int32(24))%32) + v230
	goto L59
}
func F_transform_string_values_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	if l2 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if l1&int32(3) == int32(0) {
		v32 = l1
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_appendStringInfoString(m, v118, l1)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L21
	} else {
		goto L43
	}
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = m.T0[v66].(func(*base.Module, int32, int32, int32) int32)(m, v8, l1, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v65 = v57 - l1
	goto L4
L6:
	;
	v36 = v32
	goto L15
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v65 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v21 = l1
	goto L11
L11:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v57 = v25
	goto L5
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v36
	goto L18
L17:
	;
	goto L16
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v57 = v51
	goto L5
L20:
	;
	goto L19
L21:
	;
	return int32(0)
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = F_pg_detoast_datum_packed(m, v67)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v105 = int32(1)
	if v74&v105 != 0 {
		goto L35
	} else {
		goto L36
	}
L24:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v77 = int32(4)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v79&int32(254) == int32(2) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v92 = int32(1)
	if v74&v92 != 0 {
		v104 = int32(base.Ui32(v74)>>(uint(v92)%32)) - v92
		goto L23
	} else {
		goto L34
	}
L28:
	;
	v88 = v77
	goto L30
L29:
	;
	v88 = base.B2i32(v79 == int32(18)) << (uint(v77) % 32)
	goto L30
L30:
	;
	if v79 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v91 = v77
	goto L33
L32:
	;
	v91 = v88
	goto L33
L33:
	;
	v104 = v91
	goto L23
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v104 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L35:
	;
	v109 = v105
	goto L37
L36:
	;
	v109 = int32(4)
	goto L37
L37:
	;
	F_escape_json_with_len(m, v71, v72+v109, v104)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	if v72 != v67 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pfree(m, v72)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	return int32(0)
L42:
	;
	goto L41
L43:
	;
	return int32(0)
}
