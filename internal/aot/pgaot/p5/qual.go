package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQualBegin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
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
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_CreateExecutorState(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v312 != 0 {
		goto L54
	} else {
		goto L55
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v16
	v19 = int32(4562080)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = l0
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	if v58 == v48 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+64))
	if v164 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+152))
	if v62 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v65 = v62
	goto L10
L9:
	;
	v63 = F_MakePerTupleExprContext(m, v10)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	F_ExecSetParamPlanMulti(m, v61, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v65 = v63
	goto L10
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+84))
	if v70 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = v71
	goto L15
L14:
	;
	v72 = int32(0)
	goto L15
L15:
	;
	v75 = F_palloc0(m, v72*int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v75
	v79 = v72 - int32(1)
	if v79 < int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v72&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = v79 * int32(12)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88+v85)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v85+v86)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v85)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v85)+8)) = uint8(v96)
	v100 = v72 - int32(2)
	goto L20
L19:
	;
	v100 = v79
	goto L20
L20:
	;
	if v79 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v107 = v100
	goto L22
L22:
	;
	v112 = int32(12)
	v113 = v107 * v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v116+v113)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113+v114)+4)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v113)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120+v113)+8)) = uint8(v124)
	v127 = v113 - v112
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130+v127)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v128)+4)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v127)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v127)+8)) = uint8(v138)
	if v107 != int32(1) {
		v107 = v107 - int32(2)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L6
L24:
	;
	goto L23
L25:
	;
	v207 = F_palloc0(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v167 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v168 <= v167 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v172 = v167
	goto L28
L28:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v172<<(uint(int32(2))%32))))
	v186 = F_ExecInitNode(m, v184, v16, int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	v189 = F_lappend(m, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v189
	v193 = v172 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v193 < v194 {
		v172 = v193
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v210 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v253 = F_palloc(m, v14)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L40
	}
L35:
	;
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v214 <= v213 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v218 = v213
	goto L37
L37:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v228 = int32(2)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v218<<(uint(v228)%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v226+v233<<(uint(v228)%32)-int32(4)))) = v231
	v241 = v218 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v241 < v242 {
		v218 = v241
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L34
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v253
	v256 = F_palloc0(m, v14)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v259 == int32(0) {
		v295 = v256
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v14 != 0 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v262 = int32(0)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v263 <= v262 {
		v295 = v256
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v267 = v262
	goto L45
L45:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v267<<(uint(int32(2))%32))))
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v275+v280-v282))) = uint8(v282)
	v287 = v267 + v282
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v287 < v288 {
		v267 = v287
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v295 = v290
	goto L42
L47:
	;
	goto L46
L48:
	;
	v304 = F_ExecInitNode(m, v15, v16, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L52
	}
L49:
	;
	v301 = F__emscripten_memcpy_bulkmem(m, v300, v295, v14)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v304
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v20
	return
L53:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+84))
	if v316 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v313 = F__emscripten_memcpy_bulkmem(m, v310, v311, v312)
	mBase = m.M
	goto L56
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v309)+52))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v409 = F_bms_add_member(m, v407, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L73
	}
L58:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+64))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v10)+152))
	if v321 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v324 = v321
	goto L61
L60:
	;
	v322 = F_MakePerTupleExprContext(m, v10)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	F_ExecSetParamPlanMulti(m, v320, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v324 = v322
	goto L61
L63:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+84))
	if v328 == int32(0) {
		goto L57
	} else {
		goto L64
	}
L64:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v333 = v331 - int32(1)
	if v333 < int32(0) {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	if v331&int32(1) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v339 = v333 * int32(12)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v342+v339)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v339+v340)+4)) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348+v339)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v346+v339)+8)) = uint8(v350)
	v354 = v331 - int32(2)
	goto L68
L67:
	;
	v354 = v333
	goto L68
L68:
	;
	if v333 == int32(0) {
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v361 = v354
	goto L70
L70:
	;
	v366 = int32(12)
	v367 = v361 * v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v370+v367)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v367+v368)+4)) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v367)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v374+v367)+8)) = uint8(v378)
	v381 = v367 - v366
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384+v381)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v381+v382)+4)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v381)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v388+v381)+8)) = uint8(v392)
	if v361 != int32(1) {
		v361 = v361 - int32(2)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L57
L72:
	;
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+52)) = v409
	return
}
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 float64
	_ = v64
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 int32
	_ = v138
	var v142 float64
	_ = v142
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 float64
	_ = v202
	var v203 int32
	_ = v203
	var v205 float64
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 float64
	_ = v240
	var v241 float64
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 float64
	_ = v257
	var v258 float64
	_ = v258
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == v3 {
		v289 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v289
L2:
	;
	v15 = l0
	goto L4
L3:
	;
	v283 = F_expression_tree_walker_impl(m, v15, int32(820), l1)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L25
	} else {
		goto L60
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v23 - int32(9) {
	case 0, 2:
		v289 = v3
		goto L1
	case 1:
		goto L14
	case 3, 4, 5, 7, 12, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 29, 33, 34, 35, 36, 37, 38, 40, 41, 42, 43, 44, 45, 47, 48, 49:
		goto L3
	case 6:
		goto L6
	case 8, 9, 10:
		goto L16
	case 11:
		goto L15
	case 13:
		goto L9
	case 14:
		goto L8
	case 15:
		goto L7
	case 19:
		goto L13
	case 20:
		goto L12
	case 28:
		goto L11
	case 30, 31, 32, 39, 46, 50:
		goto L10
	default:
		goto L17
	}
L5:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_add_function_cost(m, v268, v269, v15, l1+int32(8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L25
	} else {
		goto L59
	}
L6:
	;
	goto L5
L7:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v267 != 0 {
		v15 = v267
		goto L4
	} else {
		goto L58
	}
L8:
	;
	v257 = *(*float64)(unsafe.Add(mBase, uint32(v15)+56))
	v258 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v257, v258)
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v15)+64))
	v262 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v261, v262)
	v289 = v3
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L25
	} else {
		goto L55
	}
L10:
	;
	v240 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v241 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v240, v241)
	goto L3
L11:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v208 == int32(0) {
		goto L3
	} else {
		goto L48
	}
L12:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v181 = v11 + int32(40)
	v182 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v181))) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v179
	v189 = F_cost_qual_eval_walker(m, v178, v11+int32(24))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L25
	} else {
		goto L45
	}
L13:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_getTypeInputInfo(m, v150, v11+int32(24), v11+int32(8))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L25
	} else {
		goto L40
	}
L14:
	;
	v146 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v146, v147)
	v289 = v3
	goto L1
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v89 = F_estimate_array_length(m, v85, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L25
	} else {
		goto L32
	}
L16:
	;
	F_set_opfuncid(m, v15)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L25
	} else {
		goto L30
	}
L17:
	;
	switch v23 - int32(318) {
	case 0:
		goto L18
	case 1:
		v289 = v3
		goto L1
	default:
		goto L3
	}
L18:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v15)+64))
	if base.F64_lt(v28, float64(0)) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31
	v38 = v15 - int32(-64)
	v40 = v11 + int32(32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v41 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v69 = v28
	goto L21
L21:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v69, v70)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v15)+72))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v73, v74)
	v289 = v3
	goto L1
L22:
	;
	v43 = v41
	goto L24
L23:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v43 = v42
	goto L24
L24:
	;
	v46 = F_cost_qual_eval_walker(m, v43, v11+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)))
	if v50 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = base.F64_add(v53, v56)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v69 = v64
	goto L21
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_add_function_cost(m, v79, v80, v15, l1+int32(8))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L3
L32:
	;
	F_set_opfuncid(m, v15)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v93 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v93
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_add_function_cost(m, v97, v98, int32(0), v11+int32(24))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v104 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v106 = v11 + int32(16)
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_function_cost(m, v111, v104, int32(0), v11+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v131, v132)
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v137 = F_estimate_array_length(m, v136, v88)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v106)))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_mul(v89, v117), base.F64_add(v119, base.F64_add(v120, v121)))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v126, base.F64_add(v117, v127))
	goto L3
L39:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(base.F64_mul(base.F64_mul(v135, v137), float64(0.5)), v142)
	goto L3
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v161 = l1 + int32(8)
	F_add_function_cost(m, v157, v158, int32(0), v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v165 = F_exprType(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	F_getTypeOutputInfo(m, v165, v11+int32(24), v11+int32(7))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	F_add_function_cost(m, v173, v174, int32(0), v161)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	goto L3
L45:
	;
	v191 = *(*float64)(unsafe.Add(mBase, uint32(v181)))
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v192, v193)
	if base.F64_gt(v191, float64(0)) == int32(0) {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v202 = F_estimate_array_length(m, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(base.F64_mul(v191, v202), v205)
	goto L3
L48:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v211 <= int32(0) {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v220 = v3
	goto L50
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v220<<(uint(int32(2))%32))))
	v230 = F_get_opcode(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	goto L3
L52:
	;
	F_add_function_cost(m, v224, v230, int32(0), l1+int32(8))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v236 = v220 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v236 < v237 {
		v220 = v236
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	F_errmsg_internal(m, int32(117589), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(523335), int32(5006), int32(233094))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L25
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
	v289 = v3
	goto L1
L59:
	;
	goto L3
L60:
	;
	v289 = v283
	goto L1
}
func F_get_qual_from_partbound(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v144
L2:
	;
	v140 = F_get_qual_for_range(m, l0, l1, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L34
	}
L3:
	;
	v137 = F_get_qual_for_list(m, l0, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L33
	}
L4:
	;
	v19 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch v16 - int32(104) {
	case 0:
		goto L4
	default:
		v144 = v3
		goto L1
	case 4:
		goto L3
	case 10:
		goto L2
	}
L7:
	;
	v23 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = F_makeConst(m, int32(26), int32(-1), v23, int32(4), v25, v23, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v37 = F_makeConst(m, int32(23), int32(-1), v32, int32(4), v34, v32, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = F_makeConst(m, int32(23), int32(-1), v41, int32(4), v43, v41, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v46
	v60 = F_list_make3_impl(m, v10+int32(12), v10+int32(8), v10+int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = v63
	goto L14
L13:
	;
	v64 = v3
	goto L14
L14:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
	if int32(0) < v65 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v69 = int32(0)
	v70 = v60
	v74 = v64
	goto L18
L16:
	;
	v120 = v60
	goto L17
L17:
	;
	v128 = int32(0)
	v130 = F_makeFuncExpr(m, int32(5028), int32(16), v120, v128, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L31
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76+v69<<(uint(int32(1))%32)))))
	if v80 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v120 = v113
	goto L17
L20:
	;
	v113 = F_lappend(m, v70, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L29
	}
L21:
	;
	v83 = v69 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83+v84)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v83)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v83)))
	v94 = F_makeVar(m, int32(1), v80, v86, v89, v92, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v97 = F_copyObjectImpl(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	v110 = v94
	v111 = v74
	goto L20
L25:
	;
	v100 = v74 + int32(4)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if base.Ui32(v100) < base.Ui32(v103+v104<<(uint(int32(2))%32)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v109 = v100
	goto L28
L27:
	;
	v109 = int32(0)
	goto L28
L28:
	;
	v110 = v97
	v111 = v109
	goto L20
L29:
	;
	v116 = v69 + int32(1)
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v116 < v117 {
		v69 = v116
		v70 = v113
		v74 = v111
		goto L18
	} else {
		goto L30
	}
L30:
	;
	goto L19
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v130
	v135 = F_list_make1_impl(m, int32(1), v10)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v144 = v135
	goto L1
L33:
	;
	v144 = v137
	goto L1
L34:
	;
	v144 = v140
	goto L1
}
