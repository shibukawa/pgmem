package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQualBegin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 == v2 {
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
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v308 != 0 {
		goto L52
	} else {
		goto L53
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v16
	v19 = int32(_a_F_EvalPlanQualBegin_0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualBegin[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualBegin[0])) = v22
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
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+64))
	if v163 == int32(0) {
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+84))
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = v70
	goto L15
L14:
	;
	v71 = v2
	goto L15
L15:
	;
	v74 = F_palloc0(m, v71*int32(12))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v74
	v78 = v71 - int32(1)
	if v78 < int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v71&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v84 = v78 * int32(12)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v85)+4)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v84)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v84)+8)) = uint8(v95)
	v99 = v71 - int32(2)
	goto L20
L19:
	;
	v99 = v78
	goto L20
L20:
	;
	if v78 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v103 = v99
	goto L22
L22:
	;
	v111 = int32(12)
	v112 = v103 * v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115+v112)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112+v113)+4)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+v112)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v112)+8)) = uint8(v123)
	v126 = v112 - v111
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129+v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v127)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v126)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v126)+8)) = uint8(v137)
	if v103 != int32(1) {
		v103 = v103 - int32(2)
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
	v206 = F_palloc0(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 <= int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v171 = int32(0)
	goto L28
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v171<<(uint(int32(2))%32))))
	v185 = F_ExecInitNode(m, v183, v16, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	v188 = F_lappend(m, v187, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v188
	v192 = v171 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v192 < v193 {
		v171 = v192
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v209 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v252 = F_palloc(m, v14)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L40
	}
L35:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v212 <= int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v217 = int32(0)
	goto L37
L37:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v227 = int32(2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226+v217<<(uint(v227)%32))))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v225+v232<<(uint(v227)%32)-int32(4)))) = v230
	v240 = v217 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v240 < v241 {
		v217 = v240
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v252
	v255 = F_palloc0(m, v14)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v255
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v258 == int32(0) {
		v291 = v255
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v14 != 0 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v261 <= int32(0) {
		v291 = v255
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v266 = int32(0)
	goto L45
L45:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v266<<(uint(int32(2))%32))))
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v274+v279-v281))) = uint8(v281)
	v286 = v266 + v281
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v286 < v287 {
		v266 = v286
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v291 = v289
	goto L42
L47:
	;
	goto L46
L48:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	base.MemoryCopy(m, v299, v291, v14)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v302 = F_ExecInitNode(m, v15, v16, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v302
	*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualBegin[0])) = v20
	return
L52:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	base.MemoryCopy(m, v309, v310, v308)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+84))
	if v313 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v307)+52))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v406 = F_bms_add_member(m, v404, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L71
	}
L56:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+64))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v10)+152))
	if v318 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v321 = v318
	goto L59
L58:
	;
	v319 = F_MakePerTupleExprContext(m, v10)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	F_ExecSetParamPlanMulti(m, v317, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	v321 = v319
	goto L59
L61:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+84))
	if v325 == int32(0) {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v330 = v328 - int32(1)
	if v330 < int32(0) {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	if v328&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v336 = v330 * int32(12)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v336+v337)+4)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+v336)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v343+v336)+8)) = uint8(v347)
	v351 = v328 - int32(2)
	goto L66
L65:
	;
	v351 = v330
	goto L66
L66:
	;
	if v330 == int32(0) {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	v355 = v351
	goto L68
L68:
	;
	v363 = int32(12)
	v364 = v355 * v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v367+v364)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v364+v365)+4)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v364)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v371+v364)+8)) = uint8(v375)
	v378 = v364 - v363
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v381+v378)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v378+v379)+4)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387+v378)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v385+v378)+8)) = uint8(v389)
	if v355 != int32(1) {
		v355 = v355 - int32(2)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L55
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+52)) = v406
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
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v140 float64
	_ = v140
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v197 int32
	_ = v197
	var v199 float64
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == v3 {
		v283 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v283
L2:
	;
	v15 = l0
	goto L4
L3:
	;
	v277 = F_expression_tree_walker_impl(m, v15, int32(820), l1)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L60
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v23 - int32(9) {
	case 0, 2:
		v283 = v3
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
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_add_function_cost(m, v262, v263, v15, l1+int32(8))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L59
	}
L6:
	;
	goto L5
L7:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v261 != 0 {
		v15 = v261
		goto L4
	} else {
		goto L58
	}
L8:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v15)+56))
	v252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v251, v252)
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v15)+64))
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v255, v256)
	v283 = v3
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L25
	} else {
		goto L55
	}
L10:
	;
	v234 = *(*float64)(unsafe.Add(mBase, _c_F_cost_qual_eval_walker[0]))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v234, v235)
	goto L3
L11:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v202 == int32(0) {
		goto L3
	} else {
		goto L48
	}
L12:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v176 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v176
	v183 = F_cost_qual_eval_walker(m, v174, v11+int32(24))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L25
	} else {
		goto L45
	}
L13:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v150 = v11 + int32(24)
	F_getTypeInputInfo(m, v148, v150, v11+int32(8))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L25
	} else {
		goto L40
	}
L14:
	;
	v144 = *(*float64)(unsafe.Add(mBase, _c_F_cost_qual_eval_walker[0]))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v144, v145)
	v283 = v3
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
		v283 = v3
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
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v32
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
	v283 = v3
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
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v62
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
	v105 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v105
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_function_cost(m, v109, v104, int32(0), v11+int32(8))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v129, v130)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v135 = F_estimate_array_length(m, v134, v88)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_mul(v89, v115), base.F64_add(v117, base.F64_add(v118, v119)))
	v124 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v124, base.F64_add(v115, v125))
	goto L3
L39:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(base.F64_mul(base.F64_mul(v133, v135), float64(0.5)), v140)
	goto L3
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v159 = l1 + int32(8)
	F_add_function_cost(m, v155, v156, int32(0), v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v163 = F_exprType(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	F_getTypeOutputInfo(m, v163, v150, v11+int32(7))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	F_add_function_cost(m, v169, v170, int32(0), v159)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	goto L3
L45:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v186, v187)
	if base.F64_gt(v185, float64(0)) == int32(0) {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v196 = F_estimate_array_length(m, v194, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(base.F64_mul(v185, v196), v199)
	goto L3
L48:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v205 <= int32(0) {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v214 = v3
	goto L50
L50:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v214<<(uint(int32(2))%32))))
	v224 = F_get_opcode(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	goto L3
L52:
	;
	F_add_function_cost(m, v218, v224, int32(0), l1+int32(8))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v230 = v214 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v230 < v231 {
		v214 = v230
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	F_errmsg_internal(m, int32(_a_F_cost_qual_eval_walker_0), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_cost_qual_eval_walker_1), int32(_a_F_cost_qual_eval_walker_2), int32(_a_F_cost_qual_eval_walker_3))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
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
	v283 = v3
	goto L1
L59:
	;
	goto L3
L60:
	;
	v283 = v277
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
	var v145 int32
	_ = v145
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
	return v145
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
		v145 = v3
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
	v130 = F_makeFuncExpr(m, int32(_a_F_get_qual_from_partbound_0), int32(16), v120, v128, v128)
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
	v145 = v135
	goto L1
L33:
	;
	v145 = v137
	goto L1
L34:
	;
	v145 = v140
	goto L1
}
