package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_cascade_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	v16 = m.G0
	v18 = v16 - int32(624)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(295468), int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = F_ri_FetchConstraintInfo(m, v27, v28, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	v34 = F_table_open(m, v32, int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+620)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+616)) = v41
	v46 = v18 + int32(616)
	v49 = F_ri_FetchPreparedPlan(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v49 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v370 = v49
	goto L9
L8:
	;
	F_initStringInfo(m, v18+int32(600))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v371 = int32(0)
	v375 = F_ri_PerformCheck(m, v30, v46, v370, v34, v37, v36, v371, v371, int32(1), int32(8))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L70
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	v58 = F_get_namespace_name(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v60 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)) = uint8(v60)
	v64 = v18 + int32(336)
	v65 = v58
	goto L12
L12:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v79 != int32(34) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v96 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+1)) = uint16(v96)
	v99 = v18 + int32(336)
	if v99&int32(3) == int32(0) {
		v123 = v99
		goto L22
	} else {
		goto L23
	}
L14:
	;
	goto L13
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v91)
	v64 = v92
	v65 = v65 + int32(1)
	goto L12
L16:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v91 = v88
	v92 = v64 + int32(2)
	goto L15
L19:
	;
	v91 = v79
	v92 = v64 + int32(1)
	goto L15
L20:
	;
	v159 = v156 + (v18 + int32(336))
	v160 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v164 = v159 + int32(1)
	v165 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
	v169 = v162 + int32(4)
	v170 = v164
	goto L37
L21:
	;
	v156 = v148 - v99
	goto L20
L22:
	;
	v127 = v123
	goto L31
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v107 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v156 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v112 = v99
	goto L27
L27:
	;
	v116 = v112 + int32(1)
	if v116&int32(3) == int32(0) {
		v123 = v116
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v148 = v116
	goto L21
L29:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v121 != 0 {
		v112 = v116
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v136 = int32(-2139062144)
	if (int32(16843008)-v133|v133)&v136 == v136 {
		v127 = v127 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v142 = v127
	goto L34
L33:
	;
	goto L32
L34:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 != 0 {
		v142 = v142 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v148 = v142
	goto L21
L36:
	;
	goto L35
L37:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v184 != int32(34) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v201 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v170)+1)) = uint16(v201)
	if v56&int32(255) == int32(112) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v196)
	v169 = v169 + int32(1)
	v170 = v197
	goto L37
L41:
	;
	if v184 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v191 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)) = uint8(v191)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v196 = v193
	v197 = v170 + int32(2)
	goto L40
L44:
	;
	v196 = v184
	v197 = v170 + int32(1)
	goto L40
L45:
	;
	v209 = int32(719562)
	goto L47
L46:
	;
	v209 = int32(706012)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(336)
	F_appendStringInfo(m, v18+int32(600), int32(167248), v18+int32(32))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if int32(0) < v221 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v238 = int32(0)
	v240 = int32(519526)
	goto L52
L50:
	;
	v333 = v221
	goto L51
L51:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v18)+600))
	v353 = F_ri_PlanCheck(m, v348, v333, v18+int32(48), v18+int32(616), v34, v37)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L69
	}
L52:
	;
	v248 = v238 << (uint(int32(1)) % 32)
	v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(172)+v248))))
	v251 = F_attnumTypeId(m, v37, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v333 = v331
	goto L51
L54:
	;
	v253 = v248 + (v30 + int32(236))
	v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253))))
	v255 = F_attnumTypeId(m, v34, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v257 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253))))
	v258 = F_attnumAttName(m, v34, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v260 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+192)) = uint8(v260)
	v264 = v18 + int32(192)
	v265 = v258
	goto L57
L57:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v279 != int32(34) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v296 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+1)) = uint16(v296)
	v299 = v238 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v299
	v306 = F_pg_sprintf(m, v18+int32(176), int32(449202), v18+int32(16))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v291)
	v264 = v292
	v265 = v265 + int32(1)
	goto L57
L61:
	;
	if v279 == int32(0) {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v286 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)) = uint8(v286)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v291 = v288
	v292 = v264 + int32(2)
	goto L60
L64:
	;
	v291 = v279
	v292 = v264 + int32(1)
	goto L60
L65:
	;
	v309 = v238 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(300)+v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v240
	F_appendStringInfo(m, v18+int32(600), int32(700503), v18)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_generate_operator_clause(m, v18+int32(600), v18+int32(176), v251, v311, v18+int32(192), v255)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(48)+v309))) = v251
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if v299 < v331 {
		v238 = v299
		v240 = int32(522738)
		goto L52
	} else {
		goto L68
	}
L68:
	;
	goto L53
L69:
	;
	v370 = v353
	goto L9
L70:
	;
	v377 = F_SPI_finish(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v377 != int32(2) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_sequence_close(m, v34, int32(3))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L78
	}
L75:
	;
	F_errmsg_internal(m, int32(436920), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(474668), int32(1003), int32(295468))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	m.G0 = v18 + int32(624)
	return int32(0)
}
func F_RI_FKey_cascade_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
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
	var v284 int32
	_ = v284
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v427 int32
	_ = v427
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
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	v19 = m.G0
	v21 = v19 - int32(784)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(405307), int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v33 = F_ri_FetchConstraintInfo(m, v30, v31, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
	v37 = F_table_open(m, v35, int32(3))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+780)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = v45
	v50 = v21 + int32(776)
	v53 = F_ri_FetchPreparedPlan(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v427 = v53
	goto L9
L8:
	;
	F_initStringInfo(m, v21+int32(760))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v431 = F_ri_PerformCheck(m, v33, v50, v427, v37, v41, v39, v40, int32(0), int32(1), int32(9))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L73
	}
L10:
	;
	F_initStringInfo(m, v21+int32(744))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+119)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v66 = F_get_namespace_name(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v68 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+480)) = uint8(v68)
	v72 = v21 + int32(480)
	v74 = v66
	goto L13
L13:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v90 != int32(34) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v107 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+1)) = uint16(v107)
	v110 = v21 + int32(480)
	if v110&int32(3) == int32(0) {
		v134 = v110
		goto L23
	} else {
		goto L24
	}
L15:
	;
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v102)
	v72 = v103
	v74 = v74 + int32(1)
	goto L13
L17:
	;
	if v90 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)) = uint8(v97)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v102 = v99
	v103 = v72 + int32(2)
	goto L16
L20:
	;
	v102 = v90
	v103 = v72 + int32(1)
	goto L16
L21:
	;
	v170 = v167 + (v21 + int32(480))
	v171 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v171)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v175 = v170 + int32(1)
	v176 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v176)
	v180 = v173 + int32(4)
	v182 = v175
	goto L38
L22:
	;
	v167 = v159 - v110
	goto L21
L23:
	;
	v138 = v134
	goto L32
L24:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v118 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v167 = int32(0)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v123 = v110
	goto L28
L28:
	;
	v127 = v123 + int32(1)
	if v127&int32(3) == int32(0) {
		v134 = v127
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v159 = v127
	goto L22
L30:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		v123 = v127
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v147 = int32(-2139062144)
	if (int32(16843008)-v144|v144)&v147 == v147 {
		v138 = v138 + int32(4)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v153 = v138
	goto L35
L34:
	;
	goto L33
L35:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 != 0 {
		v153 = v153 + int32(1)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v159 = v153
	goto L22
L37:
	;
	goto L36
L38:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v198 != int32(34) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v215 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v182)+1)) = uint16(v215)
	v217 = int32(719562)
	if v64&int32(255) == int32(112) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v210)
	v180 = v180 + int32(1)
	v182 = v211
	goto L38
L42:
	;
	if v198 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v205 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)) = uint8(v205)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v210 = v207
	v211 = v182 + int32(2)
	goto L41
L45:
	;
	v210 = v198
	v211 = v182 + int32(1)
	goto L41
L46:
	;
	v224 = v217
	goto L48
L47:
	;
	v224 = int32(706012)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v21 + int32(480)
	F_appendStringInfo(m, v21+int32(760), int32(501346), v21+int32(48))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	if int32(0) < v236 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v252 = int32(0)
	v255 = v236
	v257 = v217
	v258 = int32(519526)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v21)+744))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v21)+748))
	F_appendBinaryStringInfo(m, v21+int32(760), v395, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L71
	}
L53:
	;
	v266 = v252 << (uint(int32(1)) % 32)
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33+int32(172)+v266))))
	v269 = F_attnumTypeId(m, v41, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	v271 = v266 + (v33 + int32(236))
	v272 = int32(*(*int16)(unsafe.Add(mBase, uint32(v271))))
	v273 = F_attnumTypeId(m, v37, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v271))))
	v276 = F_attnumAttName(m, v37, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v278 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+336)) = uint8(v278)
	v282 = v21 + int32(336)
	v284 = v276
	goto L58
L58:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v300 != int32(34) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v317 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+1)) = uint16(v317)
	v320 = v252 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(336)
	F_appendStringInfo(m, v21+int32(760), int32(449194), v21+int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v312)
	v282 = v313
	v284 = v284 + int32(1)
	goto L58
L62:
	;
	if v300 == int32(0) {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v307 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)) = uint8(v307)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v312 = v309
	v313 = v282 + int32(2)
	goto L61
L65:
	;
	v312 = v300
	v313 = v282 + int32(1)
	goto L61
L66:
	;
	v334 = v255 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v334
	v341 = F_pg_sprintf(m, v21+int32(320), int32(449202), v21+int32(16))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v344 = v252 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(300)+v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v258
	F_appendStringInfo(m, v21+int32(744), int32(700503), v21)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_generate_operator_clause(m, v21+int32(744), v21+int32(320), v269, v346, v21+int32(336), v273)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v362 = v21 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v362+v344))) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v362+v255<<(uint(int32(2))%32)))) = v269
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	if v320 < v373 {
		v252 = v320
		v255 = v334
		v257 = int32(632705)
		v258 = int32(522738)
		goto L53
	} else {
		goto L70
	}
L70:
	;
	goto L54
L71:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v21)+760))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	v407 = F_ri_PlanCheck(m, v399, v400<<(uint(int32(1))%32), v21-int32(-64), v21+int32(776), v37, v41)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v427 = v407
	goto L9
L73:
	;
	v433 = F_SPI_finish(m)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v433 != int32(2) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_sequence_close(m, v37, int32(3))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	F_errmsg_internal(m, int32(436920), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(474668), int32(1120), int32(405307))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	m.G0 = v21 + int32(784)
	return int32(0)
}
func F_RI_FKey_check(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
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
	var v611 int32
	_ = v611
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	v16 = m.G0
	v18 = v16 - int32(688)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_ri_FetchConstraintInfo(m, v20, v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28&int32(3) == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v18 + int32(688)
	return
L4:
	;
	v33 = int32(28)
	goto L6
L5:
	;
	v33 = int32(24)
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+188))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, v25, v35, int32(4120968))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v39 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v46 = F_table_open(m, v44, int32(2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v48 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L126
	}
L11:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L125
	}
L12:
	;
	v52 = v23 + int32(236)
	v54 = int32(1)
	v56 = int32(0)
	v59 = v54
	v61 = v54
	v64 = v48
	goto L13
L13:
	;
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52+v56<<(uint(int32(1))%32)))))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+6)))
	if v75 < v74 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v86&int32(1) != 0 {
		goto L11
	} else {
		goto L20
	}
L15:
	;
	F_slot_getsomeattrs_int(m, v35, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v80 = v64
	goto L17
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v83 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v74-v83))))
	v86 = v85 & v59
	v89 = (v85 ^ v83) & v61
	v91 = v56 + v83
	if v91 < v80 {
		v56 = v91
		v59 = v86
		v61 = v89
		v64 = v80
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v80 = v79
	goto L17
L19:
	;
	goto L14
L20:
	;
	if v89&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+164)))
	switch v97 - int32(102) {
	case 0:
		goto L24
	default:
		goto L21
	case 13:
		goto L23
	}
L23:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v109 = v23 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v107 + int32(4)
	F_errmsg(m, int32(661172), v18+int32(96))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errdetail(m, int32(559663), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errtableconstraint(m, v43, v109)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(474668), int32(319), int32(305282))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	goto L3
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+684)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+680)) = v136
	v141 = v18 + int32(680)
	v144 = F_ri_FetchPreparedPlan(m, v141)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v144 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v595 = v144
	goto L36
L35:
	;
	F_initStringInfo(m, v18+int32(664))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v596 = int32(0)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+119)))
	v603 = F_ri_PerformCheck(m, v23, v141, v595, v43, v46, v596, v35, v596, base.B2i32(v599 == int32(112)), int32(5))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L121
	}
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+119)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+68))
	v153 = F_get_namespace_name(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v155 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+400)) = uint8(v155)
	v159 = v153
	v161 = v18 + int32(400)
	goto L39
L39:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v174 != int32(34) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v191 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+1)) = uint16(v191)
	v194 = v18 + int32(400)
	if v194&int32(3) == int32(0) {
		v218 = v194
		goto L49
	} else {
		goto L50
	}
L41:
	;
	goto L40
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v186)
	v159 = v159 + int32(1)
	v161 = v187
	goto L39
L43:
	;
	if v174 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v181 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)) = uint8(v181)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v186 = v183
	v187 = v161 + int32(2)
	goto L42
L46:
	;
	v186 = v174
	v187 = v161 + int32(1)
	goto L42
L47:
	;
	v254 = v251 + (v18 + int32(400))
	v255 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v255)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v259 = v254 + int32(1)
	v260 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v260)
	v264 = v259
	v266 = v257 + int32(4)
	goto L64
L48:
	;
	v251 = v243 - v194
	goto L47
L49:
	;
	v222 = v218
	goto L58
L50:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v202 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v251 = int32(0)
	goto L47
L52:
	;
	goto L53
L53:
	;
	v207 = v194
	goto L54
L54:
	;
	v211 = v207 + int32(1)
	if v211&int32(3) == int32(0) {
		v218 = v211
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v243 = v211
	goto L48
L56:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v216 != 0 {
		v207 = v211
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v231 = int32(-2139062144)
	if (int32(16843008)-v228|v228)&v231 == v231 {
		v222 = v222 + int32(4)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v237 = v222
	goto L61
L60:
	;
	goto L59
L61:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v241 != 0 {
		v237 = v237 + int32(1)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v243 = v237
	goto L48
L63:
	;
	goto L62
L64:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v279 != int32(34) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v296 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+1)) = uint16(v296)
	if v151&int32(255) == int32(112) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v291)
	v264 = v292
	v266 = v266 + int32(1)
	goto L64
L68:
	;
	if v279 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v286 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)) = uint8(v286)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v291 = v288
	v292 = v264 + int32(2)
	goto L67
L71:
	;
	v291 = v279
	v292 = v264 + int32(1)
	goto L67
L72:
	;
	v304 = int32(719562)
	goto L74
L73:
	;
	v304 = int32(706012)
	goto L74
L74:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v305 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if int32(0) < v393 {
		goto L90
	} else {
		goto L91
	}
L76:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v312 = int32(*(*int16)(unsafe.Add(mBase, uint32(v308<<(uint(int32(1))%32)+v23)+170)))
	v313 = F_attnumAttName(m, v46, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v18 + int32(400)
	F_appendStringInfo(m, v18+int32(664), int32(28180), v18+int32(80))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L89
	}
L79:
	;
	v315 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v315)
	v319 = v313
	v321 = v18 + int32(256)
	goto L80
L80:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v334 != int32(34) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v351 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v321)+1)) = uint16(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v18 + int32(400)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v18 + int32(256)
	F_appendStringInfo(m, v18+int32(664), int32(28138), v18-int32(-64))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L88
	}
L82:
	;
	goto L81
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v347))) = uint8(v346)
	v319 = v319 + int32(1)
	v321 = v347
	goto L80
L84:
	;
	if v334 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v341 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)) = uint8(v341)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v346 = v343
	v347 = v321 + int32(2)
	goto L83
L87:
	;
	v346 = v334
	v347 = v321 + int32(1)
	goto L83
L88:
	;
	goto L75
L89:
	;
	goto L75
L90:
	;
	v407 = int32(0)
	v413 = int32(519526)
	goto L93
L91:
	;
	goto L92
L92:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(28201))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L110
	}
L93:
	;
	v418 = v407 << (uint(int32(1)) % 32)
	v419 = v23 + int32(172) + v418
	v420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v419))))
	v421 = F_attnumTypeId(m, v46, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	v424 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418+v52))))
	v425 = F_attnumTypeId(m, v43, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v427 = int32(*(*int16)(unsafe.Add(mBase, uint32(v419))))
	v428 = F_attnumAttName(m, v46, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v430 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v430)
	v434 = v428
	v436 = v18 + int32(256)
	goto L98
L98:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v449 != int32(34) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v466 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v436)+1)) = uint16(v466)
	v469 = v407 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v469
	v476 = F_pg_sprintf(m, v18+int32(240), int32(449202), v18+int32(48))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L106
	}
L100:
	;
	goto L99
L101:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v461)
	v434 = v434 + int32(1)
	v436 = v462
	goto L98
L102:
	;
	if v449 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v456 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)) = uint8(v456)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v461 = v458
	v462 = v436 + int32(2)
	goto L101
L105:
	;
	v461 = v449
	v462 = v436 + int32(1)
	goto L101
L106:
	;
	v479 = v407 << (uint(int32(2)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(300)+v479)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v413
	F_appendStringInfo(m, v18+int32(664), int32(700503), v18+int32(32))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_generate_operator_clause(m, v18+int32(664), v18+int32(256), v421, v481, v18+int32(240), v425)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(112)+v479))) = v425
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v469 < v503 {
		v407 = v469
		v413 = int32(522738)
		goto L93
	} else {
		goto L109
	}
L109:
	;
	goto L94
L110:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v525 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v532 = int32(*(*int16)(unsafe.Add(mBase, uint32(v526<<(uint(int32(1))%32)+v52-int32(2)))))
	v533 = F_attnumTypeId(m, v43, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v18)+664))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v578 = F_ri_PlanCheck(m, v572, v573, v18+int32(112), v18+int32(680), v43, v46)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L120
	}
L114:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(706907))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v540
	v547 = F_pg_sprintf(m, v18+int32(240), int32(449202), v18+int32(16))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v23)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(719562)
	F_appendStringInfo(m, v18+int32(664), int32(700503), v18)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_generate_operator_clause(m, v18+int32(664), v18+int32(240), v533, v549, int32(323872), int32(4537))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(638467))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	goto L113
L120:
	;
	v595 = v578
	goto L36
L121:
	;
	v605 = F_SPI_finish(m)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	if v605 != int32(2) {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L3
L125:
	;
	goto L3
L126:
	;
	F_errmsg_internal(m, int32(436920), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(474668), int32(460), int32(305282))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RI_FKey_restrict_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(295406), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_restrict_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(405227), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			F_errcode(m, int32(16908867))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(216107), v7)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					F_errfinish(m, int32(474668), int32(2175), int32(215421))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v12 != int32(442) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				F_errcode(m, int32(16908867))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(216107), v7)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errfinish(m, int32(474668), int32(2175), int32(215421))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v15&int32(28) != int32(4) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_errcode(m, int32(16908867))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = l1
						F_errmsg(m, int32(496356), v7-int32(-64))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errfinish(m, int32(474668), int32(2184), int32(215421))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v21 = v15 & int32(3)
				switch l2 - int32(2) {
				case 0:
					if v21 == int32(2) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
								F_errmsg(m, int32(518607), v7+int32(32))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(474668), int32(2198), int32(215421))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 1:
					if v21 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
								F_errmsg(m, int32(518076), v7+int32(48))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									F_errfinish(m, int32(474668), int32(2204), int32(215421))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v7 + int32(80)
						return
					}
				default:
					if v21 == int32(0) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
								F_errmsg(m, int32(497680), v7+int32(16))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									F_errfinish(m, int32(474668), int32(2192), int32(215421))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
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
			}
		}
	}
}
func F_ri_FetchPreparedPlan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1140]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v103
L2:
	;
	v50 = v13
	goto L4
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(3023656976388)
	v20 = F_hash_create(m, int32(383396), int32(64), v10, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v51 = int32(0)
	v53 = F_hash_search(m, v50, l0, v51, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1141])) = v20
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1500), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(51539607560)
	v36 = F_hash_create(m, int32(383297), int32(256), v10, int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1140])) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(292057776136)
	v45 = F_hash_create(m, int32(383710), int32(256), v10, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1142])) = v45
	v49 = *(*int32)(unsafe.Add(mBase, _consts[1140]))
	v50 = v49
	goto L4
L10:
	;
	if v53 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v57 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v61 == int32(0) {
		v91 = v60
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v91 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v91 = v60
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v71 = v2
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v71<<(uint(int32(2))%32))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+95)))
	if v79 == int32(0) {
		v91 = v79
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v91 = v79
	goto L13
L18:
	;
	v83 = v71 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v83 < v84 {
		v71 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v103 = v57
	goto L1
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(0)
	F_SPI_freeplan(m, v57)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v103 = v2
	goto L1
}
