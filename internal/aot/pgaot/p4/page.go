package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageBtreeConsolidate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v504 int32
	_ = v504
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(169)) < base.Ui32(v17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = l0 - v23 + int32(1)
	v29 = v20
	v30 = l1
	v31 = int32(0)
	goto L6
L4:
	;
	F_FreePageBtreeRemovePage(m, l0, v504)
	mBase = m.M
	goto L1
L5:
	;
	v279 = l1
	v281 = int32(0)
	v286 = v20
	goto L53
L6:
	;
	v43 = v29 + v26
	v47 = v43 - int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+3))
	v53 = int32(0)
	v54 = v50
	goto L8
L7:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v47+v83<<(uint(int32(3))%32))+24))
	if v31 <= int32(0) {
		v183 = v95
		goto L25
	} else {
		goto L26
	}
L8:
	;
	if base.Ui32(v54) <= base.Ui32(v53) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if base.Ui32(v50-int32(1)) <= base.Ui32(v83) {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	v83 = v53
	goto L10
L12:
	;
	goto L13
L13:
	;
	v69 = int32(1)
	v70 = int32(base.Ui32(v53+v54) >> (uint(v69) % 32))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v43+int32(11)+v70<<(uint(int32(3))%32))))
	v77 = base.B2i32(base.Ui32(v48) < base.Ui32(v76))
	if base.Ui32(v48) < base.Ui32(v76) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v78 = v53
	goto L16
L15:
	;
	v78 = v70 + v69
	goto L16
L16:
	;
	if base.Ui32(v48) < base.Ui32(v76) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v79 = v70
	goto L19
L18:
	;
	v79 = v54
	goto L19
L19:
	;
	if v76 != v48 {
		v53 = v78
		v54 = v79
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v83 = v70
	goto L10
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v43)+7))
	if v91 != 0 {
		v29 = v91
		v30 = v47
		v31 = v31 + int32(1)
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L7
L24:
	;
	goto L5
L25:
	;
	if v183 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L26:
	;
	v99 = v31 & int32(7)
	if v99 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if base.Ui32(v31) < base.Ui32(int32(8)) {
		v183 = v129
		goto L25
	} else {
		goto L34
	}
L28:
	;
	v128 = v31
	v129 = v95
	goto L27
L29:
	;
	goto L30
L30:
	;
	v105 = v31
	v106 = v95
	v108 = int32(0)
	goto L31
L31:
	;
	v119 = int32(1)
	v120 = v105 - v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v106+v26)+15))
	v124 = v108 + v119
	if v124 != v99 {
		v105 = v120
		v106 = v122
		v108 = v124
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v128 = v120
	v129 = v122
	goto L27
L33:
	;
	goto L32
L34:
	;
	v146 = v128
	v147 = v129
	goto L35
L35:
	;
	v160 = int32(8)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v147+v26)+15))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v26+v163)+15))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v26+v165)+15))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v26+v167)+15))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v26+v169)+15))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v26+v171)+15))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v26+v173)+15))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v26+v175)+15))
	if v160 < v146 {
		v146 = v146 - v160
		v147 = v177
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v183 = v177
	goto L25
L37:
	;
	goto L36
L38:
	;
	v198 = v183 + v26
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+3))
	if base.Ui32(int32(510)) < base.Ui32(v199+v17) {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v204 = v198 - int32(1)
	v206 = v199 << (uint(int32(3)) % 32)
	v208 = v198 + int32(11)
	v210 = l1 + int32(12)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v211 == int32(-1729435864) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v206 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v206 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	base.MemoryCopy(m, v210+v17<<(uint(int32(3))%32), v208, v206)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v198)+3))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v218 + v17
	v504 = v204
	goto L4
L46:
	;
	base.MemoryCopy(m, v210+v17<<(uint(int32(3))%32), v208, v206)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v198)+3))
	v226 = v225 + v17
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v226
	if v226 == int32(0) {
		v504 = v204
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v236 = int32(0)
	goto L50
L50:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1+v236<<(uint(int32(3))%32))+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v253)+7)) = l1 - v26 + int32(1)
	v257 = v236 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v257) < base.Ui32(v258) {
		v236 = v257
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v504 = v204
	goto L4
L52:
	;
	goto L51
L53:
	;
	v293 = v26 + v286
	v297 = v293 - int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v293)+3))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v303 = int32(0)
	v304 = v298
	goto L55
L54:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v297+v333<<(uint(int32(3))%32))+8))
	v345 = v26 + v344
	v347 = v345 - int32(1)
	if v344 != 0 {
		goto L72
	} else {
		goto L73
	}
L55:
	;
	if base.Ui32(v304) <= base.Ui32(v303) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v333 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	goto L56
L58:
	;
	v333 = v303
	goto L57
L59:
	;
	goto L60
L60:
	;
	v319 = int32(1)
	v320 = int32(base.Ui32(v303+v304) >> (uint(v319) % 32))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v293+int32(11)+v320<<(uint(int32(3))%32))))
	v327 = base.B2i32(base.Ui32(v299) < base.Ui32(v326))
	if base.Ui32(v299) < base.Ui32(v326) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v328 = v303
	goto L63
L62:
	;
	v328 = v320 + v319
	goto L63
L63:
	;
	if base.Ui32(v299) < base.Ui32(v326) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v329 = v320
	goto L66
L65:
	;
	v329 = v304
	goto L66
L66:
	;
	if v299 != v326 {
		v303 = v328
		v304 = v329
		goto L55
	} else {
		goto L67
	}
L67:
	;
	v333 = v320
	goto L57
L68:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v293)+7))
	if v340 != 0 {
		v279 = v297
		v281 = v281 + int32(1)
		v286 = v340
		goto L53
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L54
L71:
	;
	goto L1
L72:
	;
	v349 = v347
	goto L74
L73:
	;
	v349 = int32(0)
	goto L74
L74:
	;
	if v281 <= int32(0) {
		v408 = v349
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v408 == int32(0) {
		goto L1
	} else {
		goto L90
	}
L76:
	;
	if v281&int32(1) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v354 = int32(1)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v345)+3))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v347+v356<<(uint(int32(3))%32))+8))
	if v360 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v366 = v349
	v367 = v281
	goto L79
L79:
	;
	if v281 == int32(1) {
		v408 = v366
		goto L75
	} else {
		goto L83
	}
L80:
	;
	v365 = v26 + v360 - v354
	goto L82
L81:
	;
	v365 = int32(0)
	goto L82
L82:
	;
	v366 = v365
	v367 = v281 - v354
	goto L79
L83:
	;
	v372 = v366
	v373 = v367
	goto L84
L84:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v387 = int32(3)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v372+v386<<(uint(v387)%32))+8))
	v391 = v26 + v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+3))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391+v392<<(uint(v387)%32))+7))
	if v396 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v408 = v401
	goto L75
L86:
	;
	v401 = v26 + v396 - int32(1)
	goto L88
L87:
	;
	v401 = int32(0)
	goto L88
L88:
	;
	v402 = int32(2)
	if v402 < v373 {
		v372 = v401
		v373 = v373 - v402
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if base.Ui32(int32(510)) < base.Ui32(v424+v17) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v429 = v17 << (uint(int32(3)) % 32)
	v430 = int32(12)
	v431 = l1 + v430
	v433 = v408 + v430
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v434 == int32(-1729435864) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v504 = l1
	goto L4
L93:
	;
	if v429 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if v429 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	base.MemoryCopy(m, v433+v424<<(uint(int32(3))%32), v431, v429)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+4)) = v441 + v424
	goto L92
L99:
	;
	base.MemoryCopy(m, v433+v424<<(uint(int32(3))%32), v431, v429)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v449 = v448 + v424
	*(*int32)(unsafe.Add(mBase, uint32(v408)+4)) = v449
	if v449 == int32(0) {
		goto L92
	} else {
		goto L102
	}
L102:
	;
	v460 = int32(0)
	goto L103
L103:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v408+v460<<(uint(int32(3))%32))+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v476)+7)) = v408 - v26 + int32(1)
	v480 = v460 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if base.Ui32(v480) < base.Ui32(v481) {
		v460 = v480
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L92
L105:
	;
	goto L104
}
func F_FreePageBtreeRemovePage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v92 = v34 + int32(11)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v100 = v37
	v102 = int32(0)
	goto L18
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = l0 - v16 + int32(1)
	v21 = l1
	v24 = v15
	goto L5
L3:
	;
	goto L4
L4:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87
	return
L5:
	;
	v34 = v24 + v19
	v35 = int32(1)
	v36 = v34 - v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+3))
	if base.Ui32(v35) < base.Ui32(v37) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(1)
	v44 = l0 - v41 + v43
	v47 = (v21 - v19) & int32(-4096)
	v48 = v44 + v47
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(8225038576)
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v51
	v53 = v40 + v44
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v53 - v43
	goto L10
L9:
	;
	v57 = v51
	goto L10
L10:
	;
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v62 = v57 - v44 + int32(1)
	goto L13
L12:
	;
	v62 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v62
	v65 = v47 | int32(1)
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+7)) = v65
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v68 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+7))
	if v72 != 0 {
		v21 = v36
		v24 = v72
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	if base.Ui32(v100) <= base.Ui32(v102) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if base.Ui32(v37-int32(1)) <= base.Ui32(v127) {
		goto L31
	} else {
		goto L32
	}
L20:
	;
	goto L19
L21:
	;
	v127 = v102
	goto L20
L22:
	;
	goto L23
L23:
	;
	v114 = int32(1)
	v115 = int32(base.Ui32(v100+v102) >> (uint(v114) % 32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v92+v115<<(uint(int32(3))%32))))
	v122 = base.B2i32(base.Ui32(v93) < base.Ui32(v121))
	if base.Ui32(v93) < base.Ui32(v121) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v123 = v102
	goto L26
L25:
	;
	v123 = v115 + v114
	goto L26
L26:
	;
	if base.Ui32(v93) < base.Ui32(v121) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v124 = v115
	goto L29
L28:
	;
	v124 = v100
	goto L29
L29:
	;
	if v93 != v121 {
		v100 = v124
		v102 = v123
		goto L18
	} else {
		goto L30
	}
L30:
	;
	v127 = v115
	goto L20
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v34)+3))
	v150 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+3)) = v149 - v150
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = l0 - v154 + v150
	v160 = (v21 - v19) & int32(-4096)
	v161 = v157 + v160
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = int64(8225038576)
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v164
	v166 = v153 + v157
	if v153 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v134 = int32(3)
	v136 = v92 + v127<<(uint(v134)%32)
	v141 = (v37 + (v127 ^ int32(-1))) << (uint(v134) % 32)
	if v141 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	base.MemoryCopy(m, v136, v136+int32(8), v141)
	goto L31
L34:
	;
	v170 = v166 - v150
	goto L36
L35:
	;
	v170 = v164
	goto L36
L36:
	;
	if v153 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v175 = v170 - v157 + int32(1)
	goto L39
L38:
	;
	v175 = int32(0)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+12)) = v175
	v178 = v160 | int32(1)
	if v153 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+7)) = v178
	goto L42
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v181 + int32(1)
	if v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_FreePageBtreeConsolidate(m, l0, v36)
	mBase = m.M
	return
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = l0 - v185 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v34)+11))
	v192 = v36
	goto L45
L45:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	if v204 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L46:
	;
	goto L43
L47:
	;
	v207 = v204 + v188
	v209 = v207 - int32(1)
	if v204 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v211 = v209
	goto L50
L49:
	;
	v211 = int32(0)
	goto L50
L50:
	;
	v213 = v207 + int32(11)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v207)+3))
	v220 = int32(0)
	v221 = v215
	goto L51
L51:
	;
	if base.Ui32(v221) <= base.Ui32(v220) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if base.Ui32(v244) < base.Ui32(v215) {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	goto L52
L54:
	;
	v244 = v220
	goto L53
L55:
	;
	goto L56
L56:
	;
	v232 = int32(1)
	v233 = int32(base.Ui32(v220+v221) >> (uint(v232) % 32))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v213+v233<<(uint(int32(3))%32))))
	v240 = base.B2i32(base.Ui32(v189) < base.Ui32(v239))
	if base.Ui32(v189) < base.Ui32(v239) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v241 = v220
	goto L59
L58:
	;
	v241 = v233 + v232
	goto L59
L59:
	;
	if base.Ui32(v189) < base.Ui32(v239) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v242 = v233
	goto L62
L61:
	;
	v242 = v221
	goto L62
L62:
	;
	if v189 != v239 {
		v220 = v241
		v221 = v242
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v244 = v233
	goto L53
L64:
	;
	v251 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v209+v244<<(uint(int32(3))%32))+16))
	if v255 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v265 = int32(-1)
	goto L66
L66:
	;
	v266 = v265 + v244
	*(*int32)(unsafe.Add(mBase, uint32(v213+v266<<(uint(int32(3))%32)))) = v189
	if v266 == int32(0) {
		v192 = v211
		goto L45
	} else {
		goto L73
	}
L67:
	;
	v260 = v188 + v255 - int32(1)
	goto L69
L68:
	;
	v260 = v251
	goto L69
L69:
	;
	if v192 != v260 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v262 = int32(-1)
	goto L72
L71:
	;
	v262 = v251
	goto L72
L72:
	;
	v265 = v262
	goto L66
L73:
	;
	goto L46
}
func F_FreePageManagerGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v7 = F_FreePageManagerGetInternal(m, l0, l1, l2)
	mBase = m.M
	v8 = F_FreePageBtreeCleanup(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v12) < base.Ui32(v8) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
	goto L5
L4:
	;
	goto L5
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v15 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+548))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	return v7
L9:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v77)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
	goto L8
L10:
	;
	v22 = l0 + int32(36)
	v25 = int32(128)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = v18
	v62 = int32(0)
	goto L22
L13:
	;
	v31 = v25 - int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31<<(uint(int32(2))%32))))
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v73 = int32(0)
	goto L9
L15:
	;
	v73 = v25
	goto L9
L16:
	;
	goto L17
L17:
	;
	v36 = int32(2)
	v37 = v25 - v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22+v37<<(uint(v36)%32))))
	if v41 != 0 {
		v73 = v31
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v43 = v25 - int32(3)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22+v43<<(uint(int32(2))%32))))
	if v47 != 0 {
		v73 = v37
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v49 = v25 - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22+v49<<(uint(int32(2))%32))))
	if v53 != 0 {
		v73 = v43
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v49 != 0 {
		v25 = v49
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v66 = v61 + (l0 - v55 + int32(1))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+3))
	if base.Ui32(v62) < base.Ui32(v67) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v73 = v69
	goto L9
L24:
	;
	v69 = v67
	goto L26
L25:
	;
	v69 = v62
	goto L26
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+11))
	if v70 != 0 {
		v61 = v70
		v62 = v69
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
}
func F_FreePageManagerPutInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v689 int32
	_ = v689
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1263 int32
	_ = v1263
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1536 int32
	_ = v1536
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1573 int32
	_ = v1573
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1745 int32
	_ = v1745
	var v1753 int32
	_ = v1753
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1869 int32
	_ = v1869
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1986 int32
	_ = v1986
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2157 int32
	_ = v2157
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2175 int32
	_ = v2175
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2338 int32
	_ = v2338
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2445 int32
	_ = v2445
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2519 int32
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2623 int32
	_ = v2623
	var v2630 int32
	_ = v2630
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = l0 - v29 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v27 + int32(32)
	return v2630
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583))) = int32(-364896016)
	v2609 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+8)) = v2609
	v2611 = v2586 + v2587
	if v2586 != 0 {
		goto L637
	} else {
		goto L638
	}
L3:
	;
	if v478 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L4:
	;
	v767 = v470
	v770 = v475 + v470 + int32(12)
	v783 = v473
	goto L3
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L189
	} else {
		goto L190
	}
L6:
	;
	v282 = v27 + int32(16)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v294 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v294
	if v292 != 0 {
		goto L90
	} else {
		goto L91
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	v39 = int32(129)
	if base.Ui32(v39) <= base.Ui32(l2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 == v73+v34 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v42 = v39
	goto L13
L12:
	;
	v42 = l2
	goto L13
L13:
	;
	v47 = l0 + v42<<(uint(int32(2))%32) + int32(32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = l1 << (uint(int32(12)) % 32)
	v51 = v32 + v50
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(-364896016)
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v55
	v57 = v48 + v32
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v57 - int32(1)
	goto L16
L15:
	;
	v61 = v55
	goto L16
L16:
	;
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = v61 - v32 + int32(1)
	goto L19
L18:
	;
	v66 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v66
	v69 = v50 | int32(1)
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+7)) = v69
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2630 = v72
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2 + v34
	v80 = v32 + v73<<(uint(int32(12))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v73 == l1+l2 {
		goto L48
	} else {
		goto L49
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v32)+7)) = v81
	goto L28
L27:
	;
	goto L28
L28:
	;
	if v81 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v99 = int32(129)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v99) <= base.Ui32(v100) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v81+v32)+11)) = v86
	goto L29
L31:
	;
	goto L32
L32:
	;
	v88 = int32(129)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if base.Ui32(v88) <= base.Ui32(v89) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = v88
	goto L35
L34:
	;
	v92 = v89
	goto L35
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(2))%32))+32)) = v96
	goto L29
L36:
	;
	v103 = v99
	goto L38
L37:
	;
	v103 = v100
	goto L38
L38:
	;
	v108 = l0 + v103<<(uint(int32(2))%32) + int32(32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = int32(1)
	v113 = l0 - v110 + v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = v114 << (uint(int32(12)) % 32)
	v117 = v113 + v116
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = int32(-364896016)
	v121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v121
	v123 = v109 + v113
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v127 = v123 - v112
	goto L41
L40:
	;
	v127 = v121
	goto L41
L41:
	;
	if v109 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v132 = v127 - v113 + int32(1)
	goto L44
L43:
	;
	v132 = int32(0)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v132
	v135 = v116 | int32(1)
	if v109 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+7)) = v135
	goto L47
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v135
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2630 = v138
	goto L1
L48:
	;
	v143 = v32 + v73<<(uint(int32(12))%32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	if v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v204 != 0 {
		goto L74
	} else {
		goto L75
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v32)+7)) = v144
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v144 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v164 = v163 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v164
	v166 = int32(129)
	if base.Ui32(v166) <= base.Ui32(v164) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v144+v32)+11)) = v149
	goto L54
L56:
	;
	goto L57
L57:
	;
	v151 = int32(129)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if base.Ui32(v151) <= base.Ui32(v152) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v155 = v151
	goto L60
L59:
	;
	v155 = v152
	goto L60
L60:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v155<<(uint(int32(2))%32))+32)) = v159
	goto L54
L61:
	;
	v169 = v166
	goto L63
L62:
	;
	v169 = v164
	goto L63
L63:
	;
	v174 = l0 + v169<<(uint(int32(2))%32) + int32(32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = int32(1)
	v179 = l0 - v176 + v178
	v181 = l1 << (uint(int32(12)) % 32)
	v182 = v179 + v181
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = int32(-364896016)
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v186
	v188 = v175 + v179
	if v175 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v192 = v188 - v178
	goto L66
L65:
	;
	v192 = v186
	goto L66
L66:
	;
	if v175 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v197 = v192 - v179 + int32(1)
	goto L69
L68:
	;
	v197 = int32(0)
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v197
	v200 = v181 | int32(1)
	if v175 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+7)) = v200
	goto L72
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2630 = v203
	goto L1
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v240))) = int64(6860498728)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = v247
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
	v252 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240 - v32 + v252
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v252
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	if v257 != 0 {
		goto L6
	} else {
		goto L85
	}
L74:
	;
	v206 = v204 + v32
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+11))
	if v207 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if l3 != 0 {
		v2630 = int32(0)
		goto L1
	} else {
		goto L83
	}
L77:
	;
	v208 = v207 + v32
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+7)) = v209
	v213 = v208 - int32(1)
	goto L79
L78:
	;
	v213 = int32(0)
	goto L79
L79:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v215 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v214 - v215
	if v207 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v222 = v213 - v32 + v215
	goto L82
L81:
	;
	v222 = int32(0)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v222
	v240 = v206 - int32(1)
	goto L73
L83:
	;
	v230 = F_FreePageManagerGetInternal(m, l0, int32(1), v27+int32(16))
	mBase = m.M
	if v230 == int32(0) {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v240 = v32 + v233<<(uint(int32(12))%32)
	goto L73
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = l1
	v260 = int32(129)
	if base.Ui32(v260) <= base.Ui32(l2) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v263 = v260
	goto L88
L87:
	;
	v263 = l2
	goto L88
L88:
	;
	v268 = l0 + v263<<(uint(int32(2))%32) + int32(32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v273 = l0 - v270 + int32(1)
	v275 = l1 << (uint(int32(12)) % 32)
	v276 = v273 + v275
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = l2
	v2583 = v276
	v2584 = v275
	v2586 = v269
	v2587 = v273
	v2588 = v268
	goto L2
L89:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v472 = v470 + int32(4)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v475 = v473 << (uint(int32(3)) % 32)
	v476 = v472 + v475
	if v473 != 0 {
		goto L140
	} else {
		goto L141
	}
L90:
	;
	v299 = int32(1)
	v300 = l0 - v293 + v299
	v303 = v300 + v292 - v299
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if v304 == int32(430584521) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v454 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+8)) = uint8(v454)
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v454
	goto L89
L93:
	;
	v312 = v303
	v313 = v294
	goto L96
L94:
	;
	v393 = int32(2)
	v394 = v303
	goto L95
L95:
	;
	v401 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if base.Ui32(int32(509)) < base.Ui32(v403) {
		goto L121
	} else {
		goto L122
	}
L96:
	;
	v320 = v312 + int32(12)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v323 = int32(0)
	v327 = v321
	goto L98
L97:
	;
	v393 = v367 + int32(1)
	v394 = v383
	goto L95
L98:
	;
	if base.Ui32(v327) <= base.Ui32(v323) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if base.Ui32(v350) < base.Ui32(v321) {
		goto L111
	} else {
		goto L112
	}
L100:
	;
	goto L99
L101:
	;
	v350 = v323
	goto L100
L102:
	;
	goto L103
L103:
	;
	v337 = int32(1)
	v338 = int32(base.Ui32(v323+v327) >> (uint(v337) % 32))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v320+v338<<(uint(int32(3))%32))))
	v345 = base.B2i32(base.Ui32(l1) < base.Ui32(v344))
	if base.Ui32(l1) < base.Ui32(v344) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v346 = v323
	goto L106
L105:
	;
	v346 = v338 + v337
	goto L106
L106:
	;
	if base.Ui32(l1) < base.Ui32(v344) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v347 = v338
	goto L109
L108:
	;
	v347 = v327
	goto L109
L109:
	;
	if l1 != v344 {
		v323 = v346
		v327 = v347
		goto L98
	} else {
		goto L110
	}
L110:
	;
	v350 = v338
	goto L100
L111:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v320+v350<<(uint(int32(3))%32))))
	v361 = base.B2i32(v359 != l1)
	goto L113
L112:
	;
	v361 = int32(1)
	goto L113
L113:
	;
	if base.Ui32(int32(509)) < base.Ui32(v321) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v367 = v313 + int32(1)
	goto L116
L115:
	;
	v367 = int32(0)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v367
	v369 = int32(3)
	v372 = int32(0)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v312+v350<<(uint(v369)%32)-base.B2i32(v350 != v372)&v361<<(uint(v369)%32))+16))
	v381 = v300 + v378 - int32(1)
	if v378 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v383 = v381
	goto L119
L118:
	;
	v383 = v372
	goto L119
L119:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v384 == int32(430584521) {
		v312 = v383
		v313 = v367
		goto L96
	} else {
		goto L120
	}
L120:
	;
	goto L97
L121:
	;
	v406 = v393
	goto L123
L122:
	;
	v406 = v401
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v406
	v409 = v394 + int32(12)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v411 = v401
	v415 = v410
	goto L124
L124:
	;
	if base.Ui32(v415) <= base.Ui32(v411) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v394
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if base.Ui32(v438) < base.Ui32(v444) {
		goto L137
	} else {
		goto L138
	}
L126:
	;
	goto L125
L127:
	;
	v438 = v411
	goto L126
L128:
	;
	goto L129
L129:
	;
	v425 = int32(1)
	v426 = int32(base.Ui32(v411+v415) >> (uint(v425) % 32))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v409+v426<<(uint(int32(3))%32))))
	v433 = base.B2i32(base.Ui32(l1) < base.Ui32(v432))
	if base.Ui32(l1) < base.Ui32(v432) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v434 = v411
	goto L132
L131:
	;
	v434 = v426 + v425
	goto L132
L132:
	;
	if base.Ui32(l1) < base.Ui32(v432) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v435 = v426
	goto L135
L134:
	;
	v435 = v415
	goto L135
L135:
	;
	if l1 != v432 {
		v411 = v434
		v415 = v435
		goto L124
	} else {
		goto L136
	}
L136:
	;
	v438 = v426
	goto L126
L137:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v409+v438<<(uint(int32(3))%32))))
	v452 = base.B2i32(l1 == v449)
	goto L139
L138:
	;
	v452 = int32(0)
	goto L139
L139:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+8)) = uint8(v452)
	goto L89
L140:
	;
	v478 = v476
	goto L142
L141:
	;
	v478 = int32(0)
	goto L142
L142:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if base.Ui32(v473) < base.Ui32(v479) {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	if v481 == int32(0) {
		v720 = v5
		goto L144
	} else {
		goto L145
	}
L144:
	;
	if v720 != 0 {
		goto L186
	} else {
		goto L187
	}
L145:
	;
	v488 = v481
	v490 = v470
	v492 = v5
	goto L147
L146:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v512+v555<<(uint(int32(3))%32))+24))
	v570 = int32(0)
	if v492 <= v570 {
		goto L164
	} else {
		goto L165
	}
L147:
	;
	v508 = v488 + v32
	v512 = v508 - int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v508)+3))
	v520 = int32(0)
	v522 = v515
	goto L149
L148:
	;
	v720 = int32(0)
	goto L144
L149:
	;
	if base.Ui32(v522) <= base.Ui32(v520) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	if base.Ui32(v555) < base.Ui32(v515-int32(1)) {
		goto L146
	} else {
		goto L162
	}
L151:
	;
	goto L150
L152:
	;
	v555 = v520
	goto L151
L153:
	;
	goto L154
L154:
	;
	v542 = int32(1)
	v543 = int32(base.Ui32(v520+v522) >> (uint(v542) % 32))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v508+int32(11)+v543<<(uint(int32(3))%32))))
	v550 = base.B2i32(base.Ui32(v513) < base.Ui32(v549))
	if base.Ui32(v513) < base.Ui32(v549) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v551 = v520
	goto L157
L156:
	;
	v551 = v543 + v542
	goto L157
L157:
	;
	if base.Ui32(v513) < base.Ui32(v549) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v552 = v543
	goto L160
L159:
	;
	v552 = v522
	goto L160
L160:
	;
	if v513 != v549 {
		v520 = v551
		v522 = v552
		goto L149
	} else {
		goto L161
	}
L161:
	;
	v555 = v543
	goto L151
L162:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v508)+7))
	if v564 != 0 {
		v488 = v564
		v490 = v512
		v492 = v492 + int32(1)
		goto L147
	} else {
		goto L163
	}
L163:
	;
	goto L148
L164:
	;
	if v569 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v579 = v492 & int32(7)
	if v579 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v577 = v569 + v32 - int32(1)
	goto L169
L168:
	;
	v577 = int32(0)
	goto L169
L169:
	;
	v720 = v577
	goto L144
L170:
	;
	if base.Ui32(int32(8)) <= base.Ui32(v492) {
		goto L177
	} else {
		goto L178
	}
L171:
	;
	v617 = v492
	v619 = v569
	goto L170
L172:
	;
	goto L173
L173:
	;
	v586 = v492
	v587 = v570
	v588 = v569
	goto L174
L174:
	;
	v606 = int32(1)
	v607 = v586 - v606
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v588+v32)+15))
	v611 = v587 + v606
	if v611 != v579 {
		v586 = v607
		v587 = v611
		v588 = v609
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v617 = v607
	v619 = v609
	goto L170
L176:
	;
	goto L175
L177:
	;
	v643 = v617
	v645 = v619
	goto L180
L178:
	;
	v689 = v619
	goto L179
L179:
	;
	if v689 != 0 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v663 = int32(8)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v645+v32)+15))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v32+v666)+15))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v32+v668)+15))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v32+v670)+15))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v32+v672)+15))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v32+v674)+15))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v32+v676)+15))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v32+v678)+15))
	if v663 < v643 {
		v643 = v643 - v663
		v645 = v680
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v689 = v680
	goto L179
L182:
	;
	goto L181
L183:
	;
	v711 = v689 + v32 - int32(1)
	goto L185
L184:
	;
	v711 = int32(0)
	goto L185
L185:
	;
	v720 = v711
	goto L144
L186:
	;
	v739 = v720 + int32(12)
	goto L188
L187:
	;
	v739 = int32(0)
	goto L188
L188:
	;
	v767 = v720
	v770 = v739
	v783 = int32(0)
	goto L3
L189:
	;
	return int32(0)
L190:
	;
	F_errmsg_internal(m, int32(_a_F_FreePageManagerPutInternal_0), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_FreePageManagerPutInternal_1), int32(1534), int32(_a_F_FreePageManagerPutInternal_2))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L189
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
	if v770 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L194:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if base.Ui32(v786+v787) < base.Ui32(l1) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v790 = l1 + l2
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = v790 - v786
	v793 = int32(0)
	if v770 == v793 {
		v831 = v793
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v836 = l0 - v833 + int32(1)
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	v840 = v836 + v837<<(uint(int32(12))%32)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+8))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	if v842 != 0 {
		goto L209
	} else {
		goto L210
	}
L197:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	if base.Ui32(v790) < base.Ui32(v796) {
		v831 = v793
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = v798 + (v796 - v786)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v805 = l0 - v802 + int32(1)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	v809 = v805 + v806<<(uint(int32(12))%32)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+8))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v809)+12))
	if v811 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v805+v811)+7)) = v810
	goto L201
L200:
	;
	goto L201
L201:
	;
	if v810 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v831 = int32(1)
	goto L196
L203:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v809)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v810+v805)+11)) = v815
	goto L202
L204:
	;
	goto L205
L205:
	;
	v817 = int32(129)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	if base.Ui32(v817) <= base.Ui32(v818) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v821 = v817
	goto L208
L207:
	;
	v821 = v818
	goto L208
L208:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v809)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v821<<(uint(int32(2))%32))+32)) = v825
	goto L202
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v836+v842)+7)) = v841
	goto L211
L210:
	;
	goto L211
L211:
	;
	if v841 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v859 = int32(129)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if base.Ui32(v859) <= base.Ui32(v860) {
		goto L219
	} else {
		goto L220
	}
L213:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v841+v836)+11)) = v846
	goto L212
L214:
	;
	goto L215
L215:
	;
	v848 = int32(129)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if base.Ui32(v848) <= base.Ui32(v849) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v852 = v848
	goto L218
L217:
	;
	v852 = v849
	goto L218
L218:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v852<<(uint(int32(2))%32))+32)) = v856
	goto L212
L219:
	;
	v863 = v859
	goto L221
L220:
	;
	v863 = v860
	goto L221
L221:
	;
	v868 = l0 + v863<<(uint(int32(2))%32) + int32(32)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v872 = int32(1)
	v873 = l0 - v870 + v872
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	v876 = v874 << (uint(int32(12)) % 32)
	v877 = v873 + v876
	*(*int32)(unsafe.Add(mBase, uint32(v877)+4)) = v860
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = int32(-364896016)
	v881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v877)+8)) = v881
	v883 = v869 + v873
	if v869 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v887 = v883 - v872
	goto L224
L223:
	;
	v887 = v881
	goto L224
L224:
	;
	if v869 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v892 = v887 - v873 + int32(1)
	goto L227
L226:
	;
	v892 = int32(0)
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877)+12)) = v892
	v895 = v876 | int32(1)
	if v869 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v883)+7)) = v895
	goto L230
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v868))) = v895
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if v831 == int32(0) {
		v2630 = v898
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	if v912 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v2630 = v898
	goto L1
L233:
	;
	F_FreePageBtreeRemovePage(m, l0, v767)
	mBase = m.M
	goto L232
L234:
	;
	goto L235
L235:
	;
	v917 = v912 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v767)+4)) = v917
	if base.Ui32(v917) <= base.Ui32(v783) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	F_FreePageBtreeConsolidate(m, l0, v767)
	mBase = m.M
	goto L232
L237:
	;
	v921 = v767 + int32(12)
	v924 = (v917 - v783) << (uint(int32(3)) % 32)
	if v924 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v927 = v921 + v783<<(uint(int32(3))%32)
	base.MemoryCopy(m, v927, v927+int32(8), v924)
	goto L240
L239:
	;
	goto L240
L240:
	;
	if v783 != 0 {
		goto L236
	} else {
		goto L241
	}
L241:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v935 = l0 - v932 + int32(1)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	v941 = v767
	goto L242
L242:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v941)+8))
	if v951 == int32(0) {
		goto L236
	} else {
		goto L244
	}
L243:
	;
	goto L236
L244:
	;
	v954 = v951 + v935
	v956 = v954 - int32(1)
	if v951 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v958 = v956
	goto L247
L246:
	;
	v958 = int32(0)
	goto L247
L247:
	;
	v960 = v954 + int32(11)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v954)+3))
	v965 = int32(0)
	v968 = v962
	goto L248
L248:
	;
	if base.Ui32(v968) <= base.Ui32(v965) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	if base.Ui32(v992) < base.Ui32(v962) {
		goto L261
	} else {
		goto L262
	}
L250:
	;
	goto L249
L251:
	;
	v992 = v965
	goto L250
L252:
	;
	goto L253
L253:
	;
	v979 = int32(1)
	v980 = int32(base.Ui32(v965+v968) >> (uint(v979) % 32))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v960+v980<<(uint(int32(3))%32))))
	v987 = base.B2i32(base.Ui32(v936) < base.Ui32(v986))
	if base.Ui32(v936) < base.Ui32(v986) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v988 = v965
	goto L256
L255:
	;
	v988 = v980 + v979
	goto L256
L256:
	;
	if base.Ui32(v936) < base.Ui32(v986) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v989 = v980
	goto L259
L258:
	;
	v989 = v968
	goto L259
L259:
	;
	if v936 != v986 {
		v965 = v988
		v968 = v989
		goto L248
	} else {
		goto L260
	}
L260:
	;
	v992 = v980
	goto L250
L261:
	;
	v998 = int32(0)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v956+v992<<(uint(int32(3))%32))+16))
	if v1002 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1012 = int32(-1)
	goto L263
L263:
	;
	v1013 = v1012 + v992
	*(*int32)(unsafe.Add(mBase, uint32(v960+v1013<<(uint(int32(3))%32)))) = v936
	if v1013 == int32(0) {
		v941 = v958
		goto L242
	} else {
		goto L270
	}
L264:
	;
	v1007 = v935 + v1002 - int32(1)
	goto L266
L265:
	;
	v1007 = v998
	goto L266
L266:
	;
	if v1007 != v941 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1009 = int32(-1)
	goto L269
L268:
	;
	v1009 = v998
	goto L269
L269:
	;
	v1012 = v1009
	goto L263
L270:
	;
	goto L243
L271:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v1254 != 0 {
		goto L330
	} else {
		goto L331
	}
L272:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	if base.Ui32(l1+l2) < base.Ui32(v1052) {
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1058 = l0 - v1055 + int32(1)
	v1061 = v1058 + v1052<<(uint(int32(12))%32)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+8))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+12))
	if v1064 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1058+v1064)+7)) = v1062
	goto L276
L275:
	;
	goto L276
L276:
	;
	v1068 = v1063 + (v1052 - l1)
	if v1062 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1083 = int32(129)
	if base.Ui32(v1083) <= base.Ui32(v1068) {
		goto L284
	} else {
		goto L285
	}
L278:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1062+v1058)+11)) = v1070
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1072 = int32(129)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+4))
	if base.Ui32(v1072) <= base.Ui32(v1073) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1076 = v1072
	goto L283
L282:
	;
	v1076 = v1073
	goto L283
L283:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v1076<<(uint(int32(2))%32))+32)) = v1080
	goto L277
L284:
	;
	v1086 = v1083
	goto L286
L285:
	;
	v1086 = v1068
	goto L286
L286:
	;
	v1091 = l0 + v1086<<(uint(int32(2))%32) + int32(32)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1095 = int32(1)
	v1096 = l0 - v1093 + v1095
	v1098 = l1 << (uint(int32(12)) % 32)
	v1099 = v1096 + v1098
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+4)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v1099))) = int32(-364896016)
	v1103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+8)) = v1103
	v1105 = v1092 + v1096
	if v1092 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1109 = v1105 - v1095
	goto L289
L288:
	;
	v1109 = v1103
	goto L289
L289:
	;
	if v1092 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1114 = v1109 - v1096 + int32(1)
	goto L292
L291:
	;
	v1114 = int32(0)
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+12)) = v1114
	v1117 = v1098 | int32(1)
	if v1092 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+7)) = v1117
	goto L295
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(v770)+4)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v770))) = l1
	if v783 != 0 {
		v2630 = v1068
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1125 = l0 - v1122 + int32(1)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v767)+12))
	v1135 = v767
	goto L297
L297:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+8))
	if v1151 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	v2630 = v1252
	goto L1
L299:
	;
	v1152 = v1151 + v1125
	v1154 = v1152 - int32(1)
	if v1151 != 0 {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	goto L301
L301:
	;
	goto L298
L302:
	;
	v1156 = v1154
	goto L304
L303:
	;
	v1156 = int32(0)
	goto L304
L304:
	;
	v1158 = v1152 + int32(11)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+3))
	v1165 = int32(0)
	v1167 = v1160
	goto L305
L305:
	;
	if base.Ui32(v1167) <= base.Ui32(v1165) {
		goto L308
	} else {
		goto L309
	}
L306:
	;
	if base.Ui32(v1200) < base.Ui32(v1160) {
		goto L318
	} else {
		goto L319
	}
L307:
	;
	goto L306
L308:
	;
	v1200 = v1165
	goto L307
L309:
	;
	goto L310
L310:
	;
	v1187 = int32(1)
	v1188 = int32(base.Ui32(v1165+v1167) >> (uint(v1187) % 32))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1158+v1188<<(uint(int32(3))%32))))
	v1195 = base.B2i32(base.Ui32(v1126) < base.Ui32(v1194))
	if base.Ui32(v1126) < base.Ui32(v1194) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1196 = v1165
	goto L313
L312:
	;
	v1196 = v1188 + v1187
	goto L313
L313:
	;
	if base.Ui32(v1126) < base.Ui32(v1194) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1197 = v1188
	goto L316
L315:
	;
	v1197 = v1167
	goto L316
L316:
	;
	if v1126 != v1194 {
		v1165 = v1196
		v1167 = v1197
		goto L305
	} else {
		goto L317
	}
L317:
	;
	v1200 = v1188
	goto L307
L318:
	;
	v1206 = int32(0)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1154+v1200<<(uint(int32(3))%32))+16))
	if v1210 != 0 {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	v1220 = int32(-1)
	goto L320
L320:
	;
	v1221 = v1220 + v1200
	*(*int32)(unsafe.Add(mBase, uint32(v1158+v1221<<(uint(int32(3))%32)))) = v1126
	if v1221 == int32(0) {
		v1135 = v1156
		goto L297
	} else {
		goto L327
	}
L321:
	;
	v1215 = v1125 + v1210 - int32(1)
	goto L323
L322:
	;
	v1215 = v1206
	goto L323
L323:
	;
	if v1215 != v1135 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1217 = int32(-1)
	goto L326
L325:
	;
	v1217 = v1206
	goto L326
L326:
	;
	v1220 = v1217
	goto L320
L327:
	;
	goto L301
L328:
	;
	v1776 = l1
	v1778 = v1753
	v1779 = int32(0)
	goto L444
L329:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L189
	} else {
		goto L440
	}
L330:
	;
	if l3 != 0 {
		v2630 = int32(0)
		goto L1
	} else {
		goto L333
	}
L331:
	;
	v1525 = v470
	v1528 = v473
	v1536 = v479
	goto L332
L332:
	;
	v1543 = v1525 + int32(12)
	v1544 = int32(3)
	v1546 = v1543 + v1528<<(uint(v1544)%32)
	v1549 = (v1536 - v1528) << (uint(v1544) % 32)
	if v1549 != 0 {
		goto L400
	} else {
		goto L401
	}
L333:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v1254) <= base.Ui32(v1256) {
		v1753 = v470
		goto L328
	} else {
		goto L334
	}
L334:
	;
	v1263 = int32(0)
	goto L335
L335:
	;
	v1287 = F_FreePageManagerGetInternal(m, l0, int32(1), v27+int32(12))
	mBase = m.M
	if v1287 == int32(0) {
		goto L329
	} else {
		goto L337
	}
L336:
	;
	v1326 = v27 + int32(16)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1338 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1326)+12)) = v1338
	if v1336 != 0 {
		goto L349
	} else {
		goto L350
	}
L337:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1293 = int32(1)
	v1294 = l0 - v1291 + v1293
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1297 = v1295 << (uint(int32(12)) % 32)
	v1298 = v1294 + v1297
	*(*int64)(unsafe.Add(mBase, uint32(v1298))) = int64(8225038576)
	v1301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1298)+8)) = v1301
	v1303 = v1290 + v1294
	if v1290 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1307 = v1303 - v1293
	goto L340
L339:
	;
	v1307 = v1301
	goto L340
L340:
	;
	if v1290 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1312 = v1307 - v1294 + int32(1)
	goto L343
L342:
	;
	v1312 = int32(0)
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1298)+12)) = v1312
	v1315 = v1297 | int32(1)
	if v1290 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+7)) = v1315
	goto L346
L345:
	;
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1315
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1319 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1318 + v1319
	v1323 = v1263 + v1319
	if v1323 != v1254-v1256 {
		v1263 = v1323
		goto L335
	} else {
		goto L347
	}
L347:
	;
	goto L336
L348:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v1515 != 0 {
		v1753 = v1514
		goto L328
	} else {
		goto L399
	}
L349:
	;
	v1343 = int32(1)
	v1344 = l0 - v1337 + v1343
	v1347 = v1344 + v1336 - v1343
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)))
	if v1348 == int32(430584521) {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	goto L351
L351:
	;
	v1498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1326)+8)) = uint8(v1498)
	*(*int32)(unsafe.Add(mBase, uint32(v1326))) = v1498
	goto L348
L352:
	;
	v1356 = v1347
	v1357 = v1338
	goto L355
L353:
	;
	v1437 = int32(2)
	v1438 = v1347
	goto L354
L354:
	;
	v1445 = int32(0)
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	if base.Ui32(int32(509)) < base.Ui32(v1447) {
		goto L380
	} else {
		goto L381
	}
L355:
	;
	v1364 = v1356 + int32(12)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+4))
	v1367 = int32(0)
	v1371 = v1365
	goto L357
L356:
	;
	v1437 = v1411 + int32(1)
	v1438 = v1427
	goto L354
L357:
	;
	if base.Ui32(v1371) <= base.Ui32(v1367) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	if base.Ui32(v1394) < base.Ui32(v1365) {
		goto L370
	} else {
		goto L371
	}
L359:
	;
	goto L358
L360:
	;
	v1394 = v1367
	goto L359
L361:
	;
	goto L362
L362:
	;
	v1381 = int32(1)
	v1382 = int32(base.Ui32(v1367+v1371) >> (uint(v1381) % 32))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1364+v1382<<(uint(int32(3))%32))))
	v1389 = base.B2i32(base.Ui32(l1) < base.Ui32(v1388))
	if base.Ui32(l1) < base.Ui32(v1388) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1390 = v1367
	goto L365
L364:
	;
	v1390 = v1382 + v1381
	goto L365
L365:
	;
	if base.Ui32(l1) < base.Ui32(v1388) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1391 = v1382
	goto L368
L367:
	;
	v1391 = v1371
	goto L368
L368:
	;
	if l1 != v1388 {
		v1367 = v1390
		v1371 = v1391
		goto L357
	} else {
		goto L369
	}
L369:
	;
	v1394 = v1382
	goto L359
L370:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1364+v1394<<(uint(int32(3))%32))))
	v1405 = base.B2i32(v1403 != l1)
	goto L372
L371:
	;
	v1405 = int32(1)
	goto L372
L372:
	;
	if base.Ui32(int32(509)) < base.Ui32(v1365) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1411 = v1357 + int32(1)
	goto L375
L374:
	;
	v1411 = int32(0)
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1326)+12)) = v1411
	v1413 = int32(3)
	v1416 = int32(0)
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1356+v1394<<(uint(v1413)%32)-base.B2i32(v1394 != v1416)&v1405<<(uint(v1413)%32))+16))
	v1425 = v1344 + v1422 - int32(1)
	if v1422 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1427 = v1425
	goto L378
L377:
	;
	v1427 = v1416
	goto L378
L378:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1425)))
	if v1428 == int32(430584521) {
		v1356 = v1427
		v1357 = v1411
		goto L355
	} else {
		goto L379
	}
L379:
	;
	goto L356
L380:
	;
	v1450 = v1437
	goto L382
L381:
	;
	v1450 = v1445
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1326)+12)) = v1450
	v1453 = v1438 + int32(12)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	v1455 = v1445
	v1459 = v1454
	goto L383
L383:
	;
	if base.Ui32(v1459) <= base.Ui32(v1455) {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1326)+4)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v1326))) = v1438
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	if base.Ui32(v1482) < base.Ui32(v1488) {
		goto L396
	} else {
		goto L397
	}
L385:
	;
	goto L384
L386:
	;
	v1482 = v1455
	goto L385
L387:
	;
	goto L388
L388:
	;
	v1469 = int32(1)
	v1470 = int32(base.Ui32(v1455+v1459) >> (uint(v1469) % 32))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1453+v1470<<(uint(int32(3))%32))))
	v1477 = base.B2i32(base.Ui32(l1) < base.Ui32(v1476))
	if base.Ui32(l1) < base.Ui32(v1476) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1478 = v1455
	goto L391
L390:
	;
	v1478 = v1470 + v1469
	goto L391
L391:
	;
	if base.Ui32(l1) < base.Ui32(v1476) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1479 = v1470
	goto L394
L393:
	;
	v1479 = v1459
	goto L394
L394:
	;
	if l1 != v1476 {
		v1455 = v1478
		v1459 = v1479
		goto L383
	} else {
		goto L395
	}
L395:
	;
	v1482 = v1470
	goto L385
L396:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1453+v1482<<(uint(int32(3))%32))))
	v1496 = base.B2i32(l1 == v1493)
	goto L398
L397:
	;
	v1496 = int32(0)
	goto L398
L398:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1326)+8)) = uint8(v1496)
	goto L348
L399:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+4))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v1525 = v1514
	v1528 = v1517
	v1536 = v1516
	goto L332
L400:
	;
	base.MemoryCopy(m, v1546+int32(8), v1546, v1549)
	goto L402
L401:
	;
	goto L402
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1546)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1546))) = l1
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1525)+4)) = v1555 + int32(1)
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1528 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1564 = l0 - v1559 + int32(1)
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1543)))
	v1573 = v1525
	goto L406
L404:
	;
	v1698 = v1559
	goto L405
L405:
	;
	v1716 = int32(129)
	if base.Ui32(v1716) <= base.Ui32(l2) {
		goto L437
	} else {
		goto L438
	}
L406:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1573)+8))
	if v1590 != 0 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1698 = v1691
	goto L405
L408:
	;
	v1591 = v1590 + v1564
	v1593 = v1591 - int32(1)
	if v1590 != 0 {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	goto L410
L410:
	;
	goto L407
L411:
	;
	v1595 = v1593
	goto L413
L412:
	;
	v1595 = int32(0)
	goto L413
L413:
	;
	v1597 = v1591 + int32(11)
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1591)+3))
	v1604 = int32(0)
	v1606 = v1599
	goto L414
L414:
	;
	if base.Ui32(v1606) <= base.Ui32(v1604) {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	if base.Ui32(v1639) < base.Ui32(v1599) {
		goto L427
	} else {
		goto L428
	}
L416:
	;
	goto L415
L417:
	;
	v1639 = v1604
	goto L416
L418:
	;
	goto L419
L419:
	;
	v1626 = int32(1)
	v1627 = int32(base.Ui32(v1604+v1606) >> (uint(v1626) % 32))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1597+v1627<<(uint(int32(3))%32))))
	v1634 = base.B2i32(base.Ui32(v1565) < base.Ui32(v1633))
	if base.Ui32(v1565) < base.Ui32(v1633) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1635 = v1604
	goto L422
L421:
	;
	v1635 = v1627 + v1626
	goto L422
L422:
	;
	if base.Ui32(v1565) < base.Ui32(v1633) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1636 = v1627
	goto L425
L424:
	;
	v1636 = v1606
	goto L425
L425:
	;
	if v1565 != v1633 {
		v1604 = v1635
		v1606 = v1636
		goto L414
	} else {
		goto L426
	}
L426:
	;
	v1639 = v1627
	goto L416
L427:
	;
	v1645 = int32(0)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1593+v1639<<(uint(int32(3))%32))+16))
	if v1649 != 0 {
		goto L430
	} else {
		goto L431
	}
L428:
	;
	v1659 = int32(-1)
	goto L429
L429:
	;
	v1660 = v1659 + v1639
	*(*int32)(unsafe.Add(mBase, uint32(v1597+v1660<<(uint(int32(3))%32)))) = v1565
	if v1660 == int32(0) {
		v1573 = v1595
		goto L406
	} else {
		goto L436
	}
L430:
	;
	v1654 = v1564 + v1649 - int32(1)
	goto L432
L431:
	;
	v1654 = v1645
	goto L432
L432:
	;
	if v1654 != v1573 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1656 = int32(-1)
	goto L435
L434:
	;
	v1656 = v1645
	goto L435
L435:
	;
	v1659 = v1656
	goto L429
L436:
	;
	goto L410
L437:
	;
	v1719 = v1716
	goto L439
L438:
	;
	v1719 = l2
	goto L439
L439:
	;
	v1724 = l0 + v1719<<(uint(int32(2))%32) + int32(32)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1724)))
	v1728 = l0 - v1698 + int32(1)
	v1730 = l1 << (uint(int32(12)) % 32)
	v1731 = v1728 + v1730
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+4)) = l2
	v2583 = v1731
	v2584 = v1730
	v2586 = v1725
	v2587 = v1728
	v2588 = v1724
	goto L2
L440:
	;
	F_errmsg_internal(m, int32(_a_F_FreePageManagerPutInternal_0), int32(0))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L189
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(_a_F_FreePageManagerPutInternal_1), int32(1689), int32(_a_F_FreePageManagerPutInternal_2))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L189
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L443:
	;
	v2565 = int32(129)
	if base.Ui32(v2565) <= base.Ui32(l2) {
		goto L634
	} else {
		goto L635
	}
L444:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+8))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1800 = l0 - v1797 + int32(1)
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1802 = v1800 + v1801
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+11))
	if v1803 != 0 {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v2367 = v2359 + int32(11)
	v2372 = v2363
	v2375 = int32(0)
	goto L585
L446:
	;
	v1804 = v1800 + v1803
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v1804)+7)) = v1805
	v1809 = v1804 - int32(1)
	goto L448
L447:
	;
	v1809 = int32(0)
	goto L448
L448:
	;
	v1811 = v1802 + int32(11)
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1813 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1812 - v1813
	if v1803 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1820 = v1809 - v1800 + v1813
	goto L451
L450:
	;
	v1820 = int32(0)
	goto L451
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1820
	v1823 = v1802 - int32(1)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1778)))
	*(*int32)(unsafe.Add(mBase, uint32(v1823))) = v1824
	if v1801 != 0 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v1827 = v1823
	goto L454
L453:
	;
	v1827 = int32(0)
	goto L454
L454:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+4))
	v1830 = int32(base.Ui32(v1828) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v1827)+4)) = v1830
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+7)) = v1832
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+4))
	v1835 = v1834 - v1830
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+4)) = v1835
	v1838 = v1778 + int32(12)
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+4))
	v1841 = v1839 << (uint(int32(3)) % 32)
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1778)))
	if v1842 == int32(-1729435864) {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	if base.Ui32(v1776) < base.Ui32(v1923) {
		goto L467
	} else {
		goto L468
	}
L456:
	;
	if v1841 == int32(0) {
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	if v1841 != 0 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	base.MemoryCopy(m, v1811, v1838+v1835<<(uint(int32(3))%32), v1841)
	goto L455
L460:
	;
	base.MemoryCopy(m, v1811, v1838+v1835<<(uint(int32(3))%32), v1841)
	goto L462
L461:
	;
	goto L462
L462:
	;
	if v1839 == int32(0) {
		goto L455
	} else {
		goto L463
	}
L463:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1859 = int32(1)
	v1860 = l0 - v1857 + v1859
	v1869 = int32(0)
	goto L464
L464:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1823+v1869<<(uint(int32(3))%32))+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1860+v1892)+7)) = v1827 - v1860 + v1859
	v1896 = v1869 + int32(1)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+4))
	if base.Ui32(v1896) < base.Ui32(v1897) {
		v1869 = v1896
		goto L464
	} else {
		goto L466
	}
L465:
	;
	goto L455
L466:
	;
	goto L465
L467:
	;
	v1925 = v1778
	goto L469
L468:
	;
	v1925 = v1827
	goto L469
L469:
	;
	v1927 = v1925 + int32(12)
	v1928 = int32(0)
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	if v1779 == v1928 {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	if v1795 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L471:
	;
	v1935 = v1929
	v1936 = v1928
	goto L474
L472:
	;
	goto L473
L473:
	;
	v2104 = v1928
	v2106 = v1929
	goto L520
L474:
	;
	if base.Ui32(v1935) <= base.Ui32(v1936) {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	v1975 = int32(3)
	v1977 = v1927 + v1972<<(uint(v1975)%32)
	v1980 = (v1929 - v1972) << (uint(v1975) % 32)
	if v1980 != 0 {
		goto L487
	} else {
		goto L488
	}
L476:
	;
	goto L475
L477:
	;
	v1972 = v1936
	goto L476
L478:
	;
	goto L479
L479:
	;
	v1958 = int32(1)
	v1959 = int32(base.Ui32(v1935+v1936) >> (uint(v1958) % 32))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1927+v1959<<(uint(int32(3))%32))))
	v1966 = base.B2i32(base.Ui32(v1776) < base.Ui32(v1965))
	if base.Ui32(v1776) < base.Ui32(v1965) {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v1967 = v1936
	goto L482
L481:
	;
	v1967 = v1959 + v1958
	goto L482
L482:
	;
	if base.Ui32(v1776) < base.Ui32(v1965) {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v1968 = v1959
	goto L485
L484:
	;
	v1968 = v1935
	goto L485
L485:
	;
	if v1776 != v1965 {
		v1935 = v1968
		v1936 = v1967
		goto L474
	} else {
		goto L486
	}
L486:
	;
	v1972 = v1959
	goto L476
L487:
	;
	base.MemoryCopy(m, v1977+int32(8), v1977, v1980)
	goto L489
L488:
	;
	goto L489
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1977)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1977))) = v1776
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1925)+4)) = v1986 + int32(1)
	if base.B2i32(v1778 != v1925)|v1972 != 0 {
		goto L470
	} else {
		goto L490
	}
L490:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1995 = l0 - v1992 + int32(1)
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1838)))
	v2000 = v1778
	goto L491
L491:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+8))
	if v2021 == int32(0) {
		goto L470
	} else {
		goto L493
	}
L492:
	;
	goto L470
L493:
	;
	v2024 = v2021 + v1995
	v2026 = v2024 - int32(1)
	if v2021 != 0 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2028 = v2026
	goto L496
L495:
	;
	v2028 = int32(0)
	goto L496
L496:
	;
	v2030 = v2024 + int32(11)
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+3))
	v2037 = int32(0)
	v2039 = v2032
	goto L497
L497:
	;
	if base.Ui32(v2039) <= base.Ui32(v2037) {
		goto L500
	} else {
		goto L501
	}
L498:
	;
	if base.Ui32(v2072) < base.Ui32(v2032) {
		goto L510
	} else {
		goto L511
	}
L499:
	;
	goto L498
L500:
	;
	v2072 = v2037
	goto L499
L501:
	;
	goto L502
L502:
	;
	v2059 = int32(1)
	v2060 = int32(base.Ui32(v2037+v2039) >> (uint(v2059) % 32))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2030+v2060<<(uint(int32(3))%32))))
	v2067 = base.B2i32(base.Ui32(v1996) < base.Ui32(v2066))
	if base.Ui32(v1996) < base.Ui32(v2066) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v2068 = v2037
	goto L505
L504:
	;
	v2068 = v2060 + v2059
	goto L505
L505:
	;
	if base.Ui32(v1996) < base.Ui32(v2066) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2069 = v2060
	goto L508
L507:
	;
	v2069 = v2039
	goto L508
L508:
	;
	if v1996 != v2066 {
		v2037 = v2068
		v2039 = v2069
		goto L497
	} else {
		goto L509
	}
L509:
	;
	v2072 = v2060
	goto L499
L510:
	;
	v2078 = int32(0)
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2026+v2072<<(uint(int32(3))%32))+16))
	if v2082 != 0 {
		goto L513
	} else {
		goto L514
	}
L511:
	;
	v2092 = int32(-1)
	goto L512
L512:
	;
	v2093 = v2092 + v2072
	*(*int32)(unsafe.Add(mBase, uint32(v2030+v2093<<(uint(int32(3))%32)))) = v1996
	if v2093 == int32(0) {
		v2000 = v2028
		goto L491
	} else {
		goto L519
	}
L513:
	;
	v2087 = v1995 + v2082 - int32(1)
	goto L515
L514:
	;
	v2087 = v2078
	goto L515
L515:
	;
	if v2000 != v2087 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v2089 = int32(-1)
	goto L518
L517:
	;
	v2089 = v2078
	goto L518
L518:
	;
	v2092 = v2089
	goto L512
L519:
	;
	goto L492
L520:
	;
	if base.Ui32(v2106) <= base.Ui32(v2104) {
		goto L523
	} else {
		goto L524
	}
L521:
	;
	v2143 = int32(3)
	v2145 = v1927 + v2138<<(uint(v2143)%32)
	v2148 = (v1929 - v2138) << (uint(v2143) % 32)
	if v2148 != 0 {
		goto L533
	} else {
		goto L534
	}
L522:
	;
	goto L521
L523:
	;
	v2138 = v2104
	goto L522
L524:
	;
	goto L525
L525:
	;
	v2126 = int32(1)
	v2127 = int32(base.Ui32(v2104+v2106) >> (uint(v2126) % 32))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v1927+v2127<<(uint(int32(3))%32))))
	v2134 = base.B2i32(base.Ui32(v1776) < base.Ui32(v2133))
	if base.Ui32(v1776) < base.Ui32(v2133) {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2135 = v2104
	goto L528
L527:
	;
	v2135 = v2127 + v2126
	goto L528
L528:
	;
	if base.Ui32(v1776) < base.Ui32(v2133) {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2136 = v2127
	goto L531
L530:
	;
	v2136 = v2106
	goto L531
L531:
	;
	if v1776 != v2133 {
		v2104 = v2135
		v2106 = v2136
		goto L520
	} else {
		goto L532
	}
L532:
	;
	v2138 = v2127
	goto L522
L533:
	;
	base.MemoryCopy(m, v2145+int32(8), v2145, v2148)
	goto L535
L534:
	;
	goto L535
L535:
	;
	v2153 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2145)+4)) = v1779 - v32 + v2153
	*(*int32)(unsafe.Add(mBase, uint32(v2145))) = v1776
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1925)+4)) = v2157 + v2153
	*(*int32)(unsafe.Add(mBase, uint32(v1779)+8)) = v1925 - v32 + v2153
	if base.B2i32(v1778 != v1925)|v2138 != 0 {
		goto L470
	} else {
		goto L536
	}
L536:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2170 = l0 - v2167 + int32(1)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v1838)))
	v2175 = v1778
	goto L537
L537:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2175)+8))
	if v2196 == int32(0) {
		goto L470
	} else {
		goto L539
	}
L538:
	;
	goto L470
L539:
	;
	v2199 = v2196 + v2170
	v2201 = v2199 - int32(1)
	if v2196 != 0 {
		goto L540
	} else {
		goto L541
	}
L540:
	;
	v2203 = v2201
	goto L542
L541:
	;
	v2203 = int32(0)
	goto L542
L542:
	;
	v2205 = v2199 + int32(11)
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2199)+3))
	v2212 = int32(0)
	v2214 = v2207
	goto L543
L543:
	;
	if base.Ui32(v2214) <= base.Ui32(v2212) {
		goto L546
	} else {
		goto L547
	}
L544:
	;
	if base.Ui32(v2247) < base.Ui32(v2207) {
		goto L556
	} else {
		goto L557
	}
L545:
	;
	goto L544
L546:
	;
	v2247 = v2212
	goto L545
L547:
	;
	goto L548
L548:
	;
	v2234 = int32(1)
	v2235 = int32(base.Ui32(v2212+v2214) >> (uint(v2234) % 32))
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2205+v2235<<(uint(int32(3))%32))))
	v2242 = base.B2i32(base.Ui32(v2171) < base.Ui32(v2241))
	if base.Ui32(v2171) < base.Ui32(v2241) {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v2243 = v2212
	goto L551
L550:
	;
	v2243 = v2235 + v2234
	goto L551
L551:
	;
	if base.Ui32(v2171) < base.Ui32(v2241) {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v2244 = v2235
	goto L554
L553:
	;
	v2244 = v2214
	goto L554
L554:
	;
	if v2171 != v2241 {
		v2212 = v2243
		v2214 = v2244
		goto L543
	} else {
		goto L555
	}
L555:
	;
	v2247 = v2235
	goto L545
L556:
	;
	v2253 = int32(0)
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2201+v2247<<(uint(int32(3))%32))+16))
	if v2257 != 0 {
		goto L559
	} else {
		goto L560
	}
L557:
	;
	v2267 = int32(-1)
	goto L558
L558:
	;
	v2268 = v2267 + v2247
	*(*int32)(unsafe.Add(mBase, uint32(v2205+v2268<<(uint(int32(3))%32)))) = v2171
	if v2268 == int32(0) {
		v2175 = v2203
		goto L537
	} else {
		goto L565
	}
L559:
	;
	v2262 = v2170 + v2257 - int32(1)
	goto L561
L560:
	;
	v2262 = v2253
	goto L561
L561:
	;
	if v2175 != v2262 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v2264 = int32(-1)
	goto L564
L563:
	;
	v2264 = v2253
	goto L564
L564:
	;
	v2267 = v2264
	goto L558
L565:
	;
	goto L538
L566:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2305 = l0 - v2302 + int32(1)
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2307 = v2305 + v2306
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+11))
	if v2308 != 0 {
		goto L569
	} else {
		goto L570
	}
L567:
	;
	goto L568
L568:
	;
	v2359 = v32 + v1795
	v2361 = v2359 - int32(1)
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+3))
	if base.Ui32(int32(509)) < base.Ui32(v2363) {
		v1776 = v2362
		v1778 = v2361
		v1779 = v1823
		goto L444
	} else {
		goto L584
	}
L569:
	;
	v2309 = v2305 + v2308
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v2309)+7)) = v2310
	v2314 = v2309 - int32(1)
	goto L571
L570:
	;
	v2314 = int32(0)
	goto L571
L571:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2316 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2315 - v2316
	if v2308 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v2323 = v2314 - v2305 + v2316
	goto L574
L573:
	;
	v2323 = int32(0)
	goto L574
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2323
	v2325 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+7)) = v2325
	v2327 = int32(1)
	v2328 = v2307 - v2327
	*(*int64)(unsafe.Add(mBase, uint32(v2328))) = int64(9020519113)
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+15)) = v1778 - v32 + v2327
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+11)) = v2331
	if v2306 != 0 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v2338 = v2328
	goto L577
L576:
	;
	v2338 = v2325
	goto L577
L577:
	;
	if v2306 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v2343 = v2338 - v32 + int32(1)
	goto L580
L579:
	;
	v2343 = int32(0)
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+8)) = v2343
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+11))
	if v1801 != 0 {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v2350 = v1827 - v32 + int32(1)
	goto L583
L582:
	;
	v2350 = int32(0)
	goto L583
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+23)) = v2350
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+19)) = v2345
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+7)) = v2343
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2343
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2355 + int32(1)
	goto L443
L584:
	;
	goto L445
L585:
	;
	if base.Ui32(v2372) <= base.Ui32(v2375) {
		goto L588
	} else {
		goto L589
	}
L586:
	;
	v2412 = int32(3)
	v2414 = v2367 + v2408<<(uint(v2412)%32)
	v2417 = (v2363 - v2408) << (uint(v2412) % 32)
	if v2417 != 0 {
		goto L598
	} else {
		goto L599
	}
L587:
	;
	goto L586
L588:
	;
	v2408 = v2375
	goto L587
L589:
	;
	goto L590
L590:
	;
	v2395 = int32(1)
	v2396 = int32(base.Ui32(v2372+v2375) >> (uint(v2395) % 32))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2367+v2396<<(uint(int32(3))%32))))
	v2403 = base.B2i32(base.Ui32(v2362) < base.Ui32(v2402))
	if base.Ui32(v2362) < base.Ui32(v2402) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v2404 = v2375
	goto L593
L592:
	;
	v2404 = v2396 + v2395
	goto L593
L593:
	;
	if base.Ui32(v2362) < base.Ui32(v2402) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v2405 = v2396
	goto L596
L595:
	;
	v2405 = v2372
	goto L596
L596:
	;
	if v2362 != v2402 {
		v2372 = v2405
		v2375 = v2404
		goto L585
	} else {
		goto L597
	}
L597:
	;
	v2408 = v2396
	goto L587
L598:
	;
	base.MemoryCopy(m, v2414+int32(8), v2414, v2417)
	goto L600
L599:
	;
	goto L600
L600:
	;
	if v1801 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v2425 = v1827 - v32 + int32(1)
	goto L603
L602:
	;
	v2425 = int32(0)
	goto L603
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2414)+4)) = v2425
	*(*int32)(unsafe.Add(mBase, uint32(v2414))) = v2362
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+3))
	*(*int32)(unsafe.Add(mBase, uint32(v2359)+3)) = v2428 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+7)) = v1795
	if v2408 != 0 {
		goto L443
	} else {
		goto L604
	}
L604:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2436 = l0 - v2433 + int32(1)
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2367)))
	v2445 = v2361
	goto L605
L605:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2445)+8))
	if v2462 == int32(0) {
		goto L443
	} else {
		goto L607
	}
L606:
	;
	goto L443
L607:
	;
	v2465 = v2462 + v2436
	v2467 = v2465 - int32(1)
	if v2462 != 0 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v2469 = v2467
	goto L610
L609:
	;
	v2469 = int32(0)
	goto L610
L610:
	;
	v2471 = v2465 + int32(11)
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2465)+3))
	v2478 = int32(0)
	v2480 = v2473
	goto L611
L611:
	;
	if base.Ui32(v2480) <= base.Ui32(v2478) {
		goto L614
	} else {
		goto L615
	}
L612:
	;
	if base.Ui32(v2513) < base.Ui32(v2473) {
		goto L624
	} else {
		goto L625
	}
L613:
	;
	goto L612
L614:
	;
	v2513 = v2478
	goto L613
L615:
	;
	goto L616
L616:
	;
	v2500 = int32(1)
	v2501 = int32(base.Ui32(v2478+v2480) >> (uint(v2500) % 32))
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2471+v2501<<(uint(int32(3))%32))))
	v2508 = base.B2i32(base.Ui32(v2437) < base.Ui32(v2507))
	if base.Ui32(v2437) < base.Ui32(v2507) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2509 = v2478
	goto L619
L618:
	;
	v2509 = v2501 + v2500
	goto L619
L619:
	;
	if base.Ui32(v2437) < base.Ui32(v2507) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v2510 = v2501
	goto L622
L621:
	;
	v2510 = v2480
	goto L622
L622:
	;
	if v2437 != v2507 {
		v2478 = v2509
		v2480 = v2510
		goto L611
	} else {
		goto L623
	}
L623:
	;
	v2513 = v2501
	goto L613
L624:
	;
	v2519 = int32(0)
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2513<<(uint(int32(3))%32))+16))
	if v2523 != 0 {
		goto L627
	} else {
		goto L628
	}
L625:
	;
	v2533 = int32(-1)
	goto L626
L626:
	;
	v2534 = v2533 + v2513
	*(*int32)(unsafe.Add(mBase, uint32(v2471+v2534<<(uint(int32(3))%32)))) = v2437
	if v2534 == int32(0) {
		v2445 = v2469
		goto L605
	} else {
		goto L633
	}
L627:
	;
	v2528 = v2436 + v2523 - int32(1)
	goto L629
L628:
	;
	v2528 = v2519
	goto L629
L629:
	;
	if v2528 != v2445 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v2530 = int32(-1)
	goto L632
L631:
	;
	v2530 = v2519
	goto L632
L632:
	;
	v2533 = v2530
	goto L626
L633:
	;
	goto L606
L634:
	;
	v2568 = v2565
	goto L636
L635:
	;
	v2568 = l2
	goto L636
L636:
	;
	v2573 = l0 + v2568<<(uint(int32(2))%32) + int32(32)
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2573)))
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2578 = l0 - v2575 + int32(1)
	v2580 = l1 << (uint(int32(12)) % 32)
	v2581 = v2578 + v2580
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+4)) = l2
	v2583 = v2581
	v2584 = v2580
	v2586 = v2574
	v2587 = v2578
	v2588 = v2573
	goto L2
L637:
	;
	v2615 = v2611 - int32(1)
	goto L639
L638:
	;
	v2615 = v2609
	goto L639
L639:
	;
	if v2586 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v2620 = v2615 - v2587 + int32(1)
	goto L642
L641:
	;
	v2620 = int32(0)
	goto L642
L642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+12)) = v2620
	v2623 = v2584 | int32(1)
	if v2586 != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2611)+7)) = v2623
	goto L645
L644:
	;
	goto L645
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2588))) = v2623
	v2630 = l2
	goto L1
}
func F_LockPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(16973824)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v18 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_PageGetExactFreeSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v4 = v2 - v3
	v5 = int32(0)
	if v5 < v4 {
		v8 = v4
	} else {
		v8 = v5
	}
	return v8
}
func F_PageIndexTupleDeleteNoCompact(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v17) < base.Ui32(int32(24)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v216 = m.ExcPending
		if v216 != 0 {
			return
		} else {
			F_errcode(m, int32(16779816))
			mBase = m.M
			v219 = m.ExcPending
			if v219 != 0 {
				return
			} else {
				v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
				v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v222
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v221
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v220
				F_errmsg(m, int32(_a_F_PageIndexTupleDeleteNoCompact_0), v15)
				mBase = m.M
				v228 = m.ExcPending
				if v228 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PageIndexTupleDeleteNoCompact_1), int32(1314), int32(_a_F_PageIndexTupleDeleteNoCompact_2))
					mBase = m.M
					v233 = m.ExcPending
					if v233 != 0 {
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
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
		if base.Ui32(v20) < base.Ui32(v17) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v216 = m.ExcPending
			if v216 != 0 {
				return
			} else {
				F_errcode(m, int32(16779816))
				mBase = m.M
				v219 = m.ExcPending
				if v219 != 0 {
					return
				} else {
					v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
					v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v222
					*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v221
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v220
					F_errmsg(m, int32(_a_F_PageIndexTupleDeleteNoCompact_0), v15)
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PageIndexTupleDeleteNoCompact_1), int32(1314), int32(_a_F_PageIndexTupleDeleteNoCompact_2))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
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
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			if base.B2i32(base.Ui32(v22) < base.Ui32(v20))|base.B2i32(base.Ui32(int32(_a_F_PageIndexTupleDeleteNoCompact_3)) < base.Ui32(v22))|base.B2i32((v22+int32(7))&int32(_a_F_PageIndexTupleDeleteNoCompact_4) != v22) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v216 = m.ExcPending
				if v216 != 0 {
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return
					} else {
						v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
						v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v222
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v221
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v220
						F_errmsg(m, int32(_a_F_PageIndexTupleDeleteNoCompact_0), v15)
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleDeleteNoCompact_1), int32(1314), int32(_a_F_PageIndexTupleDeleteNoCompact_2))
							mBase = m.M
							v233 = m.ExcPending
							if v233 != 0 {
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
				if v17 != int32(24) {
					v40 = int32(base.Ui32(v17+int32(_a_F_PageIndexTupleDeleteNoCompact_5)) >> (uint(int32(2)) % 32))
				} else {
					v40 = int32(0)
				}
				v41 = int32(_a_F_PageIndexTupleDeleteNoCompact_6)
				v42 = v40 & v41
				if base.Ui32(v42) <= base.Ui32((l1-int32(1))&v41) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
						F_errmsg_internal(m, int32(_a_F_PageIndexTupleDeleteNoCompact_7), v15+int32(32))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleDeleteNoCompact_1), int32(1318), int32(_a_F_PageIndexTupleDeleteNoCompact_2))
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v49 = l0 + int32(20)
					v52 = v49 + l1<<(uint(int32(2))%32)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v55 = int32(base.Ui32(v53) >> (uint(int32(17)) % 32))
					v57 = v53 & int32(_a_F_PageIndexTupleDeleteNoCompact_8)
					if base.B2i32(base.Ui32(v57) < base.Ui32(v20))|base.B2i32(base.Ui32(v22) < base.Ui32(v57+v55))|base.B2i32(v57 != (v57+int32(7))&int32(_a_F_PageIndexTupleDeleteNoCompact_9)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							F_errcode(m, int32(16779816))
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v55
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v57
								F_errmsg(m, int32(_a_F_PageIndexTupleDeleteNoCompact_10), v15+int32(16))
								mBase = m.M
								v262 = m.ExcPending
								if v262 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_PageIndexTupleDeleteNoCompact_1), int32(1330), int32(_a_F_PageIndexTupleDeleteNoCompact_2))
									mBase = m.M
									v267 = m.ExcPending
									if v267 != 0 {
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
						if base.Ui32(l1) < base.Ui32(v40&int32(_a_F_PageIndexTupleDeleteNoCompact_6)) {
							*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
							v80 = v17
							v81 = v42
						} else {
							v76 = v17 - int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v76)
							v80 = v76
							v81 = v42 - int32(1)
						}
						v83 = (v55 + int32(7)) & int32(_a_F_PageIndexTupleDeleteNoCompact_9)
						if base.Ui32(v20) < base.Ui32(v57) {
							v85 = v57 - v20
							if v85 != 0 {
								v86 = l0 + v20
								base.MemoryCopy(m, v86+v83, v86, v85)
							} else {
							}
							v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
							v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							v92 = v91
							v93 = v90
						} else {
							v92 = v80
							v93 = v20
						}
						v94 = v83 + v93
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v94)
						if base.B2i32(v81 == int32(0))|base.B2i32(base.Ui32(v92&int32(_a_F_PageIndexTupleDeleteNoCompact_6)) < base.Ui32(int32(25))) != 0 {
						} else {
							v103 = int32(1)
							if v81 != v103 {
								v111 = v103
								v115 = int32(0)
								for {
									v125 = v49 + v111<<(uint(int32(2))%32)
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
									if base.B2i32(base.Ui32(v126) < base.Ui32(int32(_a_F_PageIndexTupleDeleteNoCompact_11)))|base.B2i32(base.Ui32(v57) < base.Ui32(v126&int32(_a_F_PageIndexTupleDeleteNoCompact_8))) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v125))) = (v83+v126)&int32(_a_F_PageIndexTupleDeleteNoCompact_8) | v126&int32(-32768)
									} else {
									}
									v143 = v125 + int32(4)
									v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
									if base.B2i32(base.Ui32(v144) < base.Ui32(int32(_a_F_PageIndexTupleDeleteNoCompact_11)))|base.B2i32(base.Ui32(v57) < base.Ui32(v144&int32(_a_F_PageIndexTupleDeleteNoCompact_8))) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v143))) = (v83+v144)&int32(_a_F_PageIndexTupleDeleteNoCompact_8) | v144&int32(-32768)
									} else {
									}
									v160 = int32(2)
									v161 = v111 + v160
									v163 = v115 + v160
									if v163 != v81&int32(-2) {
										v111 = v161
										v115 = v163
										continue
									} else {
										break
									}
									break
								}
								if v81&int32(1) == int32(0) {
								} else {
									v167 = v161
									v181 = v49 + v167<<(uint(int32(2))%32)
									v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
									if base.B2i32(base.Ui32(v182) < base.Ui32(int32(_a_F_PageIndexTupleDeleteNoCompact_11)))|base.B2i32(base.Ui32(v57) < base.Ui32(v182&int32(_a_F_PageIndexTupleDeleteNoCompact_8))) != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v181))) = (v182+v83)&int32(_a_F_PageIndexTupleDeleteNoCompact_8) | v182&int32(-32768)
									}
								}
							} else {
								v167 = v103
								v181 = v49 + v167<<(uint(int32(2))%32)
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
								if base.B2i32(base.Ui32(v182) < base.Ui32(int32(_a_F_PageIndexTupleDeleteNoCompact_11)))|base.B2i32(base.Ui32(v57) < base.Ui32(v182&int32(_a_F_PageIndexTupleDeleteNoCompact_8))) != 0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v181))) = (v182+v83)&int32(_a_F_PageIndexTupleDeleteNoCompact_8) | v182&int32(-32768)
								}
							}
						}
						m.G0 = v15 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_get_page_from_raw(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 == int32(1) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v14 == int32(18) {
			v17 = int32(16)
		} else {
			v17 = int32(0)
		}
		if base.Ui32((v14-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v24 = int32(4)
		} else {
			v24 = v17
		}
		v39 = v24
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_get_page_from_raw_0), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_get_page_from_raw_1)
					F_errdetail(m, int32(_a_F_get_page_from_raw_2), v6)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_page_from_raw_3), int32(230), int32(_a_F_get_page_from_raw_4))
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
			}
		}
	} else {
		if v8&int32(1) != 0 {
			v27 = int32(1)
			v39 = int32(base.Ui32(v8)>>(uint(v27)%32)) - v27
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_get_page_from_raw_0), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_get_page_from_raw_1)
						F_errdetail(m, int32(_a_F_get_page_from_raw_2), v6)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_page_from_raw_3), int32(230), int32(_a_F_get_page_from_raw_4))
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
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v35 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - int32(4)
			if v35 == int32(_a_F_get_page_from_raw_1) {
				v65 = F_palloc(m, int32(_a_F_get_page_from_raw_1))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = int32(1)
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v69&v67 != 0 {
						v72 = v67
					} else {
						v72 = int32(4)
					}
					base.MemoryCopy(m, v65, l0+v72, int32(_a_F_get_page_from_raw_1))
					m.G0 = v6 + int32(16)
					return v65
				}
			} else {
				v39 = v35
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_get_page_from_raw_0), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v39
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_get_page_from_raw_1)
							F_errdetail(m, int32(_a_F_get_page_from_raw_2), v6)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_page_from_raw_3), int32(230), int32(_a_F_get_page_from_raw_4))
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
					}
				}
			}
		}
	}
}
