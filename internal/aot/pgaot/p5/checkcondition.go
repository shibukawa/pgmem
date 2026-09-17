package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_gin_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+(l1-v5)>>(uint(int32(3))%32)))))
	return v10
}
func F_checkcondition_str_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14) <= base.Ui32(v13) {
		v149 = int32(0)
		v150 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v156 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v157 != v156 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = v16 & int32(4095)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v19 + int32(base.Ui32(v16)>>(uint(int32(12))%32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = v13
	v30 = v14
	goto L3
L3:
	;
	v37 = int32(2)
	v40 = base.I32_div_s((v30-v28)>>(uint(v37)%32), v37)
	v43 = v28 + v40<<(uint(v37)%32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v48 = int32(base.Ui32(v44)>>(uint(int32(1))%32)) & int32(2047)
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if base.Ui32(v28) < base.Ui32(v30) {
		goto L47
	} else {
		goto L48
	}
L5:
	;
	goto L4
L6:
	;
	v128 = int32(0)
	v132 = base.B2i32(v128 < v126)
	if v128 < v126 {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	if v48 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v48 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L10:
	;
	v126 = int32(1)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v54 = v23 + int32(base.Ui32(v44)>>(uint(int32(12))%32))
	v55 = base.B2i32(base.Ui32(v18) < base.Ui32(v48))
	if base.Ui32(v18) < base.Ui32(v48) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = v18
	goto L15
L14:
	;
	v56 = v48
	goto L15
L15:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v56) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v118 != 0 {
		v126 = v118
		goto L6
	} else {
		goto L34
	}
L17:
	;
	v118 = int32(0)
	goto L16
L18:
	;
	v92 = v87
	v93 = v88
	v94 = v89
	goto L28
L19:
	;
	if (v22|v54)&int32(3) != 0 {
		v87 = v22
		v88 = v54
		v89 = v56
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v80 = v22
	v81 = v54
	v82 = v56
	goto L21
L21:
	;
	if v82 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v64 = v22
	v65 = v54
	v66 = v56
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v69 != v70 {
		v87 = v64
		v88 = v65
		v89 = v66
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v80 = v75
	v81 = v73
	v82 = v77
	goto L21
L25:
	;
	v72 = int32(4)
	v73 = v65 + v72
	v75 = v64 + v72
	v77 = v66 - v72
	if base.Ui32(int32(3)) < base.Ui32(v77) {
		v64 = v75
		v65 = v73
		v66 = v77
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v87 = v80
	v88 = v81
	v89 = v82
	goto L18
L28:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 == v98 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = v97 - v98
	goto L16
L30:
	;
	v100 = int32(1)
	v105 = v94 - v100
	if v105 != 0 {
		v92 = v92 + v100
		v93 = v93 + v100
		v94 = v105
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	if v48 == v18 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v18) < base.Ui32(v48) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v122 = int32(-1)
	goto L38
L37:
	;
	v122 = int32(1)
	goto L38
L38:
	;
	v126 = v122
	goto L6
L39:
	;
	v126 = int32(-1)
	goto L6
L40:
	;
	v133 = v43 + int32(4)
	goto L42
L41:
	;
	v133 = v28
	goto L42
L42:
	;
	if v128 < v126 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v134 = v30
	goto L45
L44:
	;
	v134 = v43
	goto L45
L45:
	;
	if base.Ui32(v133) < base.Ui32(v134) {
		v28 = v133
		v30 = v134
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v149 = v128
	v150 = v134
	goto L1
L47:
	;
	v139 = v43
	goto L49
L48:
	;
	v139 = v30
	goto L49
L49:
	;
	v140 = F_checkclass_str(m, l0, v43, l1, l2)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	return int32(0)
L51:
	;
	v149 = v140
	v150 = v139
	goto L1
L52:
	;
	return v149
L53:
	;
	goto L54
L54:
	;
	if base.B2i32(l2 == int32(0))&base.B2i32(v149 == int32(1)) != 0 {
		v468 = v156
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return v468
L56:
	;
	if l2 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v166 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v178 = int32(0)
	v186 = v178
	v187 = v178
	v188 = v150
	v189 = v178
	v190 = v178
	goto L66
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_pfree(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L50
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v172)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v172
	goto L59
L63:
	;
	goto L62
L64:
	;
	if v190 == int32(0) {
		v468 = int32(2)
		goto L55
	} else {
		goto L141
	}
L65:
	;
	if l2 == int32(0) {
		v468 = v189
		goto L55
	} else {
		goto L128
	}
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v194) <= base.Ui32(v188) {
		goto L65
	} else {
		goto L68
	}
L67:
	;
	return int32(1)
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v198 = v196 & int32(4095)
	if v198 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v203 = int32(base.Ui32(v199)>>(uint(int32(1))%32)) & int32(2047)
	if v203 == int32(0) {
		goto L65
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v282 = F_checkclass_str(m, l0, v188, l1, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L50
	} else {
		goto L96
	}
L72:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v207 = int32(12)
	v209 = v206 + int32(base.Ui32(v196)>>(uint(v207)%32))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v213 = v210 + int32(base.Ui32(v199)>>(uint(v207)%32))
	if base.Ui32(v198) < base.Ui32(v203) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v215 = v198
	goto L75
L74:
	;
	v215 = v203
	goto L75
L75:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v215) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v277|base.B2i32(base.Ui32(v203) < base.Ui32(v198)) != 0 {
		goto L65
	} else {
		goto L94
	}
L77:
	;
	v277 = int32(0)
	goto L76
L78:
	;
	v251 = v246
	v252 = v247
	v253 = v248
	goto L88
L79:
	;
	if (v209|v213)&int32(3) != 0 {
		v246 = v209
		v247 = v213
		v248 = v215
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v239 = v209
	v240 = v213
	v241 = v215
	goto L81
L81:
	;
	if v241 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v223 = v209
	v224 = v213
	v225 = v215
	goto L83
L83:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v228 != v229 {
		v246 = v223
		v247 = v224
		v248 = v225
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v239 = v234
	v240 = v232
	v241 = v236
	goto L81
L85:
	;
	v231 = int32(4)
	v232 = v224 + v231
	v234 = v223 + v231
	v236 = v225 - v231
	if base.Ui32(int32(3)) < base.Ui32(v236) {
		v223 = v234
		v224 = v232
		v225 = v236
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v246 = v239
	v247 = v240
	v248 = v241
	goto L78
L88:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v256 == v257 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v277 = v256 - v257
	goto L76
L90:
	;
	v259 = int32(1)
	v264 = v253 - v259
	if v264 != 0 {
		v251 = v251 + v259
		v252 = v252 + v259
		v253 = v264
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	goto L71
L95:
	;
	if l2|base.B2i32(v364 != int32(1)) != 0 {
		v186 = v361
		v187 = v362
		v188 = v188 + int32(4)
		v189 = v364
		v190 = v365
		goto L66
	} else {
		goto L127
	}
L96:
	;
	if v282 == int32(0) {
		v361 = v186
		v362 = v187
		v364 = v189
		v365 = v190
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if l2 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v282 == int32(2) {
		goto L64
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v282 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L101:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v186 < v288+v187 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v295 = v186
	v299 = v190
	goto L105
L103:
	;
	v323 = v288
	v324 = v186
	v328 = v190
	goto L104
L104:
	;
	v333 = v323 << (uint(int32(1)) % 32)
	if v333 != 0 {
		goto L114
	} else {
		goto L115
	}
L105:
	;
	if v295 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v323 = v317
	v324 = v315
	v328 = v316
	goto L104
L107:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v315 < v317+v187 {
		v295 = v315
		v299 = v316
		goto L105
	} else {
		goto L113
	}
L108:
	;
	v307 = F_palloc(m, int32(512))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L50
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v311 = F_repalloc(m, v299, v295<<(uint(int32(2))%32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L50
	} else {
		goto L112
	}
L111:
	;
	v315 = int32(256)
	v316 = v307
	goto L107
L112:
	;
	v315 = v295 << (uint(int32(1)) % 32)
	v316 = v311
	goto L107
L113:
	;
	goto L106
L114:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	base.MemoryCopy(m, v328+v187<<(uint(int32(1))%32), v337, v333)
	goto L116
L115:
	;
	goto L116
L116:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v340 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_pfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L50
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v347 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v347)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v347
	v361 = v324
	v362 = v339 + v187
	v364 = v189
	v365 = v328
	goto L95
L120:
	;
	goto L119
L121:
	;
	v355 = v282
	goto L123
L122:
	;
	v355 = v189
	goto L123
L123:
	;
	if v189 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v356 = v355
	goto L126
L125:
	;
	v356 = v282
	goto L126
L126:
	;
	v361 = v186
	v362 = v187
	v364 = v356
	v365 = v190
	goto L95
L127:
	;
	goto L67
L128:
	;
	v382 = int32(1)
	if v187 <= int32(0) {
		v468 = v189
		goto L55
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v190
	F_pg_qsort(m, v190, v187, int32(2), int32(1520))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L50
	} else {
		goto L130
	}
L130:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v187) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v398 = v382
	v400 = int32(0)
	goto L134
L132:
	;
	v442 = v382
	goto L133
L133:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v450)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v442
	v468 = v450
	goto L55
L134:
	;
	v406 = int32(1)
	v408 = v392 + v398<<(uint(v406)%32)
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408))))
	v413 = int32(_a_F_checkcondition_str_1_0)
	v414 = v412 & v413
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392+v400<<(uint(v406)%32)))))
	v417 = v415 & v413
	goto L137
L135:
	;
	v442 = v432 + int32(1)
	goto L133
L136:
	;
	v434 = v398 + int32(1)
	if v434 != v187 {
		v398 = v434
		v400 = v432
		goto L134
	} else {
		goto L140
	}
L137:
	;
	if base.B2i32(base.Ui32(v417) < base.Ui32(v414))-base.B2i32(base.Ui32(v414) < base.Ui32(v417)) == int32(0) {
		v432 = v400
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v424 = v400 + int32(1)
	if v424 == v398 {
		v432 = v398
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408))))
	*(*uint16)(unsafe.Add(mBase, uint32(v392+v424<<(uint(int32(1))%32)))) = uint16(v429)
	v432 = v424
	goto L136
L140:
	;
	goto L135
L141:
	;
	F_pfree(m, v190)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L50
	} else {
		goto L142
	}
L142:
	;
	return int32(2)
}
