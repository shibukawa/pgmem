package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v49
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v38 != v39 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v49 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 != 0 {
		v49 = v3
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v45 = F_equal(m, v43, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v49 = v45
	goto L1
}
func F_get_extension_control_directories(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	v9 = m.G0
	v11 = v9 - int32(1056)
	m.G0 = v11
	F_get_share_path(m, v11+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(32)
	v25 = F_psprintf(m, int32(271598), v11+int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[452]))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(1056)
	return v227
L5:
	;
	v33 = F_lappend(m, int32(0), v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v35 = F_pstrdup(m, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v227 = v33
	goto L4
L9:
	;
	v37 = v35
	v40 = int32(0)
	goto L10
L10:
	;
	v46 = v37
	goto L14
L11:
	;
	v227 = v220
	goto L4
L12:
	;
	v63 = v61 + int32(1)
	v64 = F_palloc(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L13:
	;
	if v56 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	goto L15
L17:
	;
	v56 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v48 == int32(58) {
		v56 = v46
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v46 = v46 + int32(1)
	goto L14
L21:
	;
	v59 = F_strlen(m, v37)
	mBase = m.M
	v61 = v59
	goto L12
L22:
	;
	goto L23
L23:
	;
	v61 = v56 - v37
	goto L12
L24:
	;
	if v63 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v180 = int32(289513)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _consts[453])))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v184 == int32(0) {
		v203 = v183
		v204 = v184
		goto L59
	} else {
		goto L60
	}
L26:
	;
	v177 = F_strlen(m, v173)
	mBase = m.M
	goto L25
L27:
	;
	v173 = v37
	goto L26
L28:
	;
	goto L29
L29:
	;
	v71 = v63 - int32(1)
	if (v64^v37)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v170)
	v173 = v166
	goto L26
L31:
	;
	v151 = v146
	v152 = v147
	v153 = v148
	goto L53
L32:
	;
	if v141 == int32(0) {
		v166 = v139
		v167 = v140
		goto L30
	} else {
		goto L52
	}
L33:
	;
	v139 = v37
	v140 = v64
	v141 = v71
	goto L32
L34:
	;
	goto L35
L35:
	;
	v75 = int32(0)
	if v37&int32(3) == v75 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v108 == int32(0) {
		v166 = v105
		v167 = v106
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v105 = v37
	v106 = v64
	v107 = v71
	v108 = base.B2i32(v71 != v75)
	goto L36
L38:
	;
	if v71 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v84 = v37
	v85 = v64
	v86 = v71
	goto L40
L40:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v88)
	if v88 == int32(0) {
		v146 = v84
		v147 = v85
		v148 = v86
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v105 = v99
	v106 = v93
	v107 = v95
	v108 = v97
	goto L36
L42:
	;
	v92 = int32(1)
	v93 = v85 + v92
	v95 = v86 - v92
	v96 = int32(0)
	v97 = base.B2i32(v95 != v96)
	v99 = v84 + v92
	if v99&int32(3) == v96 {
		v105 = v99
		v106 = v93
		v107 = v95
		v108 = v97
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v95 != 0 {
		v84 = v99
		v85 = v93
		v86 = v95
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v111 == int32(0) {
		v139 = v105
		v140 = v106
		v141 = v107
		goto L32
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(v107) < base.Ui32(int32(4)) {
		v139 = v105
		v140 = v106
		v141 = v107
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v117 = v105
	v118 = v106
	v119 = v107
	goto L48
L48:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v125 = int32(-2139062144)
	if (int32(16843008)-v122|v122)&v125 != v125 {
		v146 = v117
		v147 = v118
		v148 = v119
		goto L31
	} else {
		goto L50
	}
L49:
	;
	v139 = v133
	v140 = v131
	v141 = v135
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v122
	v130 = int32(4)
	v131 = v118 + v130
	v133 = v117 + v130
	v135 = v119 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v117 = v133
		v118 = v131
		v119 = v135
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v146 = v139
	v147 = v140
	v148 = v141
	goto L31
L53:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v155)
	if v155 == int32(0) {
		v166 = v151
		v167 = v152
		goto L30
	} else {
		goto L55
	}
L54:
	;
	v166 = v162
	v167 = v160
	goto L30
L55:
	;
	v159 = int32(1)
	v160 = v152 + v159
	v162 = v151 + v159
	v164 = v153 - v159
	if v164 != 0 {
		v151 = v162
		v152 = v160
		v153 = v164
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	F_pfree(m, v64)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L71
	}
L58:
	;
	if v204-v203 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	goto L58
L60:
	;
	if v183 != v184 {
		v203 = v183
		v204 = v184
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v188 = v64
	v189 = v180
	goto L62
L62:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v193 == int32(0) {
		v203 = v192
		v204 = v193
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v203 = v192
	v204 = v193
	goto L59
L64:
	;
	v196 = int32(1)
	if v192 == v193 {
		v188 = v188 + v196
		v189 = v189 + v196
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v209 = F_substitute_path_macro(m, v64, int32(289513), v25)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
	v213 = F_psprintf(m, int32(271598), v11)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v215 = v209
	goto L57
L70:
	;
	v215 = v213
	goto L57
L71:
	;
	F_canonicalize_path_enc(m, v215)
	mBase = m.M
	v220 = F_lappend(m, v40, v215)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v61))))
	if v223 != 0 {
		v37 = v37 + v63
		v40 = v220
		goto L10
	} else {
		goto L73
	}
L73:
	;
	goto L11
}
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
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
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v848 int32
	_ = v848
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(240)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v3
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L12
	} else {
		goto L270
	}
L2:
	;
	v201 = F_strlen(m, v192)
	mBase = m.M
	v204 = F_pnstrdup(m, v192, v201-int32(10))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L56
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v56
	v77 = F_psprintf(m, int32(301620), v14+int32(208))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L26
	}
L4:
	;
	if v69 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L8:
	;
	v42 = F_palloc(m, int32(1024))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L19
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(47) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v40 = v24
	goto L8
L14:
	;
	v29 = F_pstrdup(m, v20)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v31
	v37 = F_psprintf(m, int32(177111), v14+int32(192))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v40 = v29
	goto L8
L18:
	;
	v40 = v37
	goto L8
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+180)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v40
	v52 = F_pg_snprintf(m, v42, int32(1024), int32(301613), v14+int32(176))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v69 = v42
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v57
	v65 = F_psprintf(m, int32(301599), v14+int32(224))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v69 = v65
	goto L4
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = v72
	v195 = v69
	goto L2
L25:
	;
	if v132 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L26:
	;
	v79 = F_get_extension_control_directories(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v81 = int32(0)
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	if v79 == v81 {
		v132 = v81
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L44
	}
L29:
	;
	m.G0 = v84 + int32(32)
	goto L25
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v88 <= int32(0) {
		v132 = v81
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = v3
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v99<<(uint(int32(2))%32))))
	v107 = F_pstrdup(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v132 = v81
	goto L29
L34:
	;
	F_canonicalize_path_enc(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 != int32(47) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107
	v116 = F_psprintf(m, int32(177111), v84)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v118 = F_pg_file_exists(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v118 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = v116
	goto L29
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v107)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v125 = v99 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v125 < v126 {
		v99 = v125
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(321749)
	F_errmsg(m, int32(321994), v84+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(496737), int32(4047), int32(155037))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v166 = F_strlen(m, v132)
	mBase = m.M
	v173 = v166 + int32(1)
	goto L51
L49:
	;
	v187 = F_pnstrdup(m, v132, v185-v132)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L55
	}
L50:
	;
	goto L49
L51:
	;
	v175 = int32(0)
	if v173 == v175 {
		v185 = v175
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v185 = v180
	goto L50
L53:
	;
	v179 = v173 - int32(1)
	v180 = v132 + v179
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != int32(47) {
		v173 = v179
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v187
	v192 = v187
	v195 = v132
	goto L2
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204
	v208 = F_AllocateFile(m, v195, int32(230944))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v195)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L12
	} else {
		goto L269
	}
L58:
	;
	if v208 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if l1 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v239 = F_ParseConfigFp(m, v208, v195, int32(0), int32(21), v14+int32(236), v14+int32(232))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L70
	}
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v213 == int32(44) {
		goto L57
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v195
	F_errmsg(m, int32(298958), v14+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(496737), int32(696), int32(386819))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v241 = F_FreeFile(m, v208)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	if v243 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v256 = v243
	goto L75
L73:
	;
	v811 = int32(0)
	goto L74
L74:
	;
	F_FreeConfigVariables(m, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L12
	} else {
		goto L262
	}
L75:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v266 = int32(13593)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, _consts[441])))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v270 == int32(0) {
		v289 = v269
		v290 = v270
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	v811 = v798
	goto L74
L77:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	if v797 != 0 {
		v256 = v797
		goto L75
	} else {
		goto L261
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L12
	} else {
		goto L257
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L12
	} else {
		goto L253
	}
L80:
	;
	if v290-v289 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	goto L80
L82:
	;
	if v269 != v270 {
		v289 = v269
		v290 = v270
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v274 = v265
	v275 = v266
	goto L84
L84:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v279 == int32(0) {
		v289 = v278
		v290 = v279
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v289 = v278
	v290 = v279
	goto L81
L86:
	;
	v282 = int32(1)
	if v278 == v279 {
		v274 = v274 + v282
		v275 = v275 + v282
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	if l1 != 0 {
		goto L79
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v298 = int32(271193)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, _consts[442])))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v302 == int32(0) {
		v321 = v301
		v322 = v302
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v295 = F_pstrdup(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v295
	goto L77
L93:
	;
	if v322-v321 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	goto L93
L95:
	;
	if v301 != v302 {
		v321 = v301
		v322 = v302
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v306 = v265
	v307 = v298
	goto L97
L97:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	if v311 == int32(0) {
		v321 = v310
		v322 = v311
		goto L94
	} else {
		goto L99
	}
L98:
	;
	v321 = v310
	v322 = v311
	goto L94
L99:
	;
	v314 = int32(1)
	if v310 == v311 {
		v306 = v306 + v314
		v307 = v307 + v314
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	if l1 != 0 {
		goto L78
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v330 = int32(378058)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, _consts[443])))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v334 == int32(0) {
		v353 = v333
		v354 = v334
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v327 = F_pstrdup(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v327
	goto L77
L106:
	;
	if v354-v353 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L107:
	;
	goto L106
L108:
	;
	if v333 != v334 {
		v353 = v333
		v354 = v334
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v338 = v265
	v339 = v330
	goto L110
L110:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+1)))
	if v343 == int32(0) {
		v353 = v342
		v354 = v343
		goto L107
	} else {
		goto L112
	}
L111:
	;
	v353 = v342
	v354 = v343
	goto L107
L112:
	;
	v346 = int32(1)
	if v342 == v343 {
		v338 = v338 + v346
		v339 = v339 + v346
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v359 = F_pstrdup(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v362 = int32(95079)
	v365 = int32(*(*uint8)(unsafe.Add(mBase, _consts[444])))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v366 == int32(0) {
		v385 = v365
		v386 = v366
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v359
	goto L77
L118:
	;
	if v386-v385 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L119:
	;
	goto L118
L120:
	;
	if v365 != v366 {
		v385 = v365
		v386 = v366
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v370 = v265
	v371 = v362
	goto L122
L122:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v375 == int32(0) {
		v385 = v374
		v386 = v375
		goto L119
	} else {
		goto L124
	}
L123:
	;
	v385 = v374
	v386 = v375
	goto L119
L124:
	;
	v378 = int32(1)
	if v374 == v375 {
		v370 = v370 + v378
		v371 = v371 + v378
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v391 = F_pstrdup(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L12
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v394 = int32(506881)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, _consts[445])))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v398 == int32(0) {
		v417 = v397
		v418 = v398
		goto L131
	} else {
		goto L132
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v391
	goto L77
L130:
	;
	if v418-v417 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L131:
	;
	goto L130
L132:
	;
	if v397 != v398 {
		v417 = v397
		v418 = v398
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v402 = v265
	v403 = v394
	goto L134
L134:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+1)))
	if v407 == int32(0) {
		v417 = v406
		v418 = v407
		goto L131
	} else {
		goto L136
	}
L135:
	;
	v417 = v406
	v418 = v407
	goto L131
L136:
	;
	v410 = int32(1)
	if v406 == v407 {
		v402 = v402 + v410
		v403 = v403 + v410
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v423 = F_pstrdup(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L12
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v426 = int32(392325)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, _consts[446])))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v430 == int32(0) {
		v449 = v429
		v450 = v430
		goto L143
	} else {
		goto L144
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v423
	goto L77
L142:
	;
	if v450-v449 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	goto L142
L144:
	;
	if v429 != v430 {
		v449 = v429
		v450 = v430
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v434 = v265
	v435 = v426
	goto L146
L146:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+1)))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)))
	if v439 == int32(0) {
		v449 = v438
		v450 = v439
		goto L143
	} else {
		goto L148
	}
L147:
	;
	v449 = v438
	v450 = v439
	goto L143
L148:
	;
	v442 = int32(1)
	if v438 == v439 {
		v434 = v434 + v442
		v435 = v435 + v442
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v455 = F_strlen(m, v454)
	mBase = m.M
	v456 = F_parse_bool_with_len(m, v454, v455, l0+int32(32))
	mBase = m.M
	goto L153
L151:
	;
	goto L152
L152:
	;
	v476 = int32(217612)
	v479 = int32(*(*uint8)(unsafe.Add(mBase, _consts[447])))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v480 == int32(0) {
		v499 = v479
		v500 = v480
		goto L160
	} else {
		goto L161
	}
L153:
	;
	if v456 != 0 {
		goto L77
	} else {
		goto L154
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v464
	F_errmsg(m, int32(346074), v14-int32(-64))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(496737), int32(751), int32(386819))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L12
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	if v500-v499 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L160:
	;
	goto L159
L161:
	;
	if v479 != v480 {
		v499 = v479
		v500 = v480
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v484 = v265
	v485 = v476
	goto L163
L163:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+1)))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+1)))
	if v489 == int32(0) {
		v499 = v488
		v500 = v489
		goto L160
	} else {
		goto L165
	}
L164:
	;
	v499 = v488
	v500 = v489
	goto L160
L165:
	;
	v492 = int32(1)
	if v488 == v489 {
		v484 = v484 + v492
		v485 = v485 + v492
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v505 = F_strlen(m, v504)
	mBase = m.M
	v506 = F_parse_bool_with_len(m, v504, v505, l0+int32(33))
	mBase = m.M
	goto L170
L168:
	;
	goto L169
L169:
	;
	v526 = int32(441162)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, _consts[448])))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v530 == int32(0) {
		v549 = v529
		v550 = v530
		goto L177
	} else {
		goto L178
	}
L170:
	;
	if v506 != 0 {
		goto L77
	} else {
		goto L171
	}
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v514
	F_errmsg(m, int32(346074), v14+int32(80))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L12
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(496737), int32(759), int32(386819))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	if v550-v549 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L177:
	;
	goto L176
L178:
	;
	if v529 != v530 {
		v549 = v529
		v550 = v530
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v534 = v265
	v535 = v526
	goto L180
L180:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+1)))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	if v539 == int32(0) {
		v549 = v538
		v550 = v539
		goto L177
	} else {
		goto L182
	}
L181:
	;
	v549 = v538
	v550 = v539
	goto L177
L182:
	;
	v542 = int32(1)
	if v538 == v539 {
		v534 = v534 + v542
		v535 = v535 + v542
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v555 = F_strlen(m, v554)
	mBase = m.M
	v556 = F_parse_bool_with_len(m, v554, v555, l0+int32(34))
	mBase = m.M
	goto L187
L185:
	;
	goto L186
L186:
	;
	v576 = int32(336102)
	v579 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v580 == int32(0) {
		v599 = v579
		v600 = v580
		goto L194
	} else {
		goto L195
	}
L187:
	;
	if v556 != 0 {
		goto L77
	} else {
		goto L188
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v564
	F_errmsg(m, int32(346074), v14+int32(96))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(496737), int32(767), int32(386819))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	if v600-v599 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L194:
	;
	goto L193
L195:
	;
	if v579 != v580 {
		v599 = v579
		v600 = v580
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v584 = v265
	v585 = v576
	goto L197
L197:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+1)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+1)))
	if v589 == int32(0) {
		v599 = v588
		v600 = v589
		goto L194
	} else {
		goto L199
	}
L198:
	;
	v599 = v588
	v600 = v589
	goto L194
L199:
	;
	v592 = int32(1)
	if v588 == v589 {
		v584 = v584 + v592
		v585 = v585 + v592
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v606 = F_pg_char_to_encoding_private(m, v604)
	mBase = m.M
	if base.Ui32(int32(35)) <= base.Ui32(v606) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	goto L203
L203:
	;
	v632 = int32(161769)
	v635 = int32(*(*uint8)(unsafe.Add(mBase, _consts[449])))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v636 == int32(0) {
		v655 = v635
		v656 = v636
		goto L214
	} else {
		goto L215
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v609
	if int32(0) <= v609 {
		goto L77
	} else {
		goto L208
	}
L205:
	;
	v609 = int32(-1)
	goto L207
L206:
	;
	v609 = v606
	goto L207
L207:
	;
	goto L204
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L12
	} else {
		goto L210
	}
L210:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v620
	F_errmsg(m, int32(381921), v14+int32(112))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L12
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(496737), int32(776), int32(386819))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L12
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	if v656-v655 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L214:
	;
	goto L213
L215:
	;
	if v635 != v636 {
		v655 = v635
		v656 = v636
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v640 = v265
	v641 = v632
	goto L217
L217:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+1)))
	if v645 == int32(0) {
		v655 = v644
		v656 = v645
		goto L214
	} else {
		goto L219
	}
L218:
	;
	v655 = v644
	v656 = v645
	goto L214
L219:
	;
	v648 = int32(1)
	if v644 == v645 {
		v640 = v640 + v648
		v641 = v641 + v648
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v661 = F_pstrdup(m, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L12
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v685 = int32(357624)
	v688 = int32(*(*uint8)(unsafe.Add(mBase, _consts[450])))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v689 == int32(0) {
		v708 = v688
		v709 = v689
		goto L232
	} else {
		goto L233
	}
L224:
	;
	v664 = F_SplitIdentifierString(m, v661, int32(44), l0+int32(40))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L12
	} else {
		goto L225
	}
L225:
	;
	if v664 != 0 {
		goto L77
	} else {
		goto L226
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L12
	} else {
		goto L227
	}
L227:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L12
	} else {
		goto L228
	}
L228:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v673
	F_errmsg(m, int32(163490), v14+int32(128))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L12
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(496737), int32(790), int32(386819))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L12
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	if v709-v708 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L232:
	;
	goto L231
L233:
	;
	if v688 != v689 {
		v708 = v688
		v709 = v689
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v693 = v265
	v694 = v685
	goto L235
L235:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693)+1)))
	if v698 == int32(0) {
		v708 = v697
		v709 = v698
		goto L232
	} else {
		goto L237
	}
L236:
	;
	v708 = v697
	v709 = v698
	goto L232
L237:
	;
	v701 = int32(1)
	if v697 == v698 {
		v693 = v693 + v701
		v694 = v694 + v701
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v714 = F_pstrdup(m, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L12
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L12
	} else {
		goto L249
	}
L242:
	;
	v717 = F_SplitIdentifierString(m, v714, int32(44), l0+int32(44))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L12
	} else {
		goto L243
	}
L243:
	;
	if v717 != 0 {
		goto L77
	} else {
		goto L244
	}
L244:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v726
	F_errmsg(m, int32(163490), v14+int32(144))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(496737), int32(805), int32(386819))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L12
	} else {
		goto L250
	}
L250:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v745
	F_errmsg(m, int32(716719), v14+int32(160))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(496737), int32(812), int32(386819))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L12
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L12
	} else {
		goto L254
	}
L254:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v765
	F_errmsg(m, int32(388202), v14+int32(32))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(496737), int32(719), int32(386819))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L12
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L12
	} else {
		goto L258
	}
L258:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v784
	F_errmsg(m, int32(388202), v14+int32(48))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L12
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(496737), int32(729), int32(386819))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L12
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	goto L76
L262:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v814 != int32(1) {
		goto L57
	} else {
		goto L263
	}
L263:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v817 == int32(0) {
		goto L57
	} else {
		goto L264
	}
L264:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L12
	} else {
		goto L265
	}
L265:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L12
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(344482), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L12
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(496737), int32(820), int32(386819))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L12
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	m.G0 = v14 + int32(240)
	return
L270:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L12
	} else {
		goto L271
	}
L271:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v870
	F_errmsg(m, int32(395891), v14)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L12
	} else {
		goto L272
	}
L272:
	;
	F_errhint(m, int32(627193), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L12
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(496737), int32(673), int32(386819))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L12
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
