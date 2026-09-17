package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAllSchemaPublicationRelations(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	v4 = F_GetPublicationSchemas(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v13 <= v12 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L8
L8:
	;
	v18 = v12
	v20 = int32(0)
	goto L9
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v22 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v18<<(uint(v22)%32))))
	v27 = F_GetSchemaPublicationRelations(m, v25, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	return v29
L11:
	;
	v29 = F_list_concat(m, v20, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v32 < v33 {
		v18 = v32
		v20 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
}
func F_ReleaseAllPlanCacheRefsInOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v12 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v15 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L13
	} else {
		goto L26
	}
L5:
	;
	v19 = l0 + int32(24)
	v21 = v17
	v22 = v17
	v23 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v29 = v19 + v23<<(uint(int32(3))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v30 == int32(_a_F_ReleaseAllPlanCacheRefsInOwner_0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v19+v21<<(uint(int32(3))%32)-int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v43 = v41 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v43)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[0]))
	m.T0[v46].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v52 = v22
	v53 = v23
	goto L12
L12:
	;
	v55 = v53 + int32(1)
	v57 = v52 & int32(255)
	if v55 < v57 {
		v21 = v57
		v22 = v52
		v23 = v55
		goto L8
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v52 = v51
	v53 = v23 - int32(1)
	goto L12
L15:
	;
	goto L9
L16:
	;
	v69 = v66
	v71 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v109)
	m.G0 = v10 + int32(16)
	goto L1
L19:
	;
	v76 = v71 << (uint(int32(3)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v78 = v76 + v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 == int32(_a_F_ReleaseAllPlanCacheRefsInOwner_0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, uint32(v85+v76)+4)) = v83
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v89 - int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[0]))
	m.T0[v94].(func(*base.Module, int32))(m, v82)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	v98 = v69
	goto L23
L23:
	;
	v100 = v71 + int32(1)
	if base.Ui32(v100) < base.Ui32(v98) {
		v69 = v98
		v71 = v100
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v98 = v97
	goto L23
L25:
	;
	goto L20
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v119
	F_errmsg_internal(m, int32(_a_F_ReleaseAllPlanCacheRefsInOwner_1), v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ReleaseAllPlanCacheRefsInOwner_2), int32(822), int32(_a_F_ReleaseAllPlanCacheRefsInOwner_3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_disable_all_timeouts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_disable_all_timeouts[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_disable_all_timeouts[1])) = v2
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[2])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[3])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[4])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[5])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[6])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[7])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[8])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[9])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[10])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[11])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[12])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[13])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[14])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[15])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[16])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[17])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[18])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[19])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[20])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[21])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[22])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[23])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[24])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[25])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[26])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[27])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[28])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[29])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[30])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[31])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[32])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[33])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[34])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[35])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[36])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[37])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[38])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[39])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[40])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[41])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[42])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[43])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[44])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[45])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[46])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[47])) = uint8(v2)
	return
}
func F_show_all_settings(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
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
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
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
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 float64
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 float64
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 float64
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 float64
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
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
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int64
	_ = v409
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v10 = m.G0
	v12 = v10 - int32(480)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	goto L28
L4:
	;
	return int32(0)
L5:
	;
	v22 = int32(_a_F_show_all_settings_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0])) = v25
	v28 = F_CreateTemplateTupleDesc(m, int32(17))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v28, int32(1), int32(_a_F_show_all_settings_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v28, int32(2), int32(_a_F_show_all_settings_2), int32(25), int32(-1), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v28, int32(3), int32(_a_F_show_all_settings_3), int32(25), int32(-1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v28, int32(4), int32(_a_F_show_all_settings_4), int32(25), int32(-1), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v28, int32(5), int32(_a_F_show_all_settings_5), int32(25), int32(-1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v28, int32(6), int32(_a_F_show_all_settings_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v28, int32(7), int32(_a_F_show_all_settings_7), int32(25), int32(-1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v28, int32(8), int32(_a_F_show_all_settings_8), int32(25), int32(-1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v28, int32(9), int32(_a_F_show_all_settings_9), int32(25), int32(-1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v28, int32(10), int32(_a_F_show_all_settings_10), int32(25), int32(-1), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v28, int32(11), int32(_a_F_show_all_settings_11), int32(25), int32(-1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v28, int32(12), int32(_a_F_show_all_settings_12), int32(1009), int32(-1), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v28, int32(13), int32(_a_F_show_all_settings_13), int32(25), int32(-1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v28, int32(14), int32(_a_F_show_all_settings_14), int32(25), int32(-1), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v28, int32(15), int32(_a_F_show_all_settings_15), int32(25), int32(-1), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v28, int32(16), int32(_a_F_show_all_settings_16), int32(23), int32(-1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_TupleDescInitEntry(m, v28, int32(17), int32(_a_F_show_all_settings_17), int32(16), int32(-1), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v149 = F_TupleDescGetAttInMetadata(m, v28)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v149
	v154 = F_get_guc_variables(m, v12+int32(220))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v154
	v157 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+220)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v157
	*(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0])) = v23
	goto L3
L26:
	;
	m.G0 = v12 + int32(480)
	return v466
L27:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v221
	v224 = F_ShowGUCOption(m, v184, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L42
	}
L28:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	v167 = base.I32_wrap_i64(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v167 < v168 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v174 = v167
	v180 = v166
	goto L32
L30:
	;
	goto L31
L31:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L41
	}
L32:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v171+v174<<(uint(int32(2))%32))))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+20))
	if v185&int32(4) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v198 = v180
	goto L36
L35:
	;
	if v185&int32(1024) == int32(0) {
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v200 = v198 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v200
	v202 = base.I32_wrap_i64(v200)
	if v202 < v168 {
		v174 = v202
		v180 = v200
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[1]))
	v195 = F_has_privs_of_role(m, v193, int32(3374))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v195 != 0 {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	v198 = v197
	goto L36
L40:
	;
	goto L33
L41:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = int32(2)
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v218)
	v466 = int32(0)
	goto L26
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+148)) = v224
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v184)+20))
	v228 = F_get_config_unit_name(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	v232 = int32(2)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231<<(uint(v232)%32))+uint32(_c_F_show_all_settings[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240<<(uint(v232)%32))+uint32(_c_F_show_all_settings[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245<<(uint(v232)%32))+uint32(_c_F_show_all_settings[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v184)+32))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250<<(uint(v232)%32))+uint32(_c_F_show_all_settings[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v253
	switch v245 {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		goto L45
	}
L44:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v184)+32))
	if v415 != int32(3) {
		goto L89
	} else {
		goto L90
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	v409 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+188)) = v409
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = v409
	goto L44
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v392 = F_config_enum_get_options(m, v184, int32(_a_F_show_all_settings_18), int32(_a_F_show_all_settings_19), int32(_a_F_show_all_settings_20))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L83
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	if v373 != 0 {
		goto L75
	} else {
		goto L76
	}
L48:
	;
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v184)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+128)) = v321
	v324 = v12 + int32(224)
	v329 = F_pg_snprintf(m, v324, int32(256), int32(_a_F_show_all_settings_21), v12+int32(128))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L67
	}
L49:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v184)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v273
	v276 = v12 + int32(224)
	v281 = F_pg_snprintf(m, v276, int32(256), int32(_a_F_show_all_settings_22), v12-int32(-64))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L59
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+96)))
	if v261 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v262 = int32(_a_F_show_all_settings_23)
	goto L53
L52:
	;
	v262 = int32(_a_F_show_all_settings_24)
	goto L53
L53:
	;
	v263 = F_pstrdup(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v263
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+112)))
	if v268 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v269 = int32(_a_F_show_all_settings_23)
	goto L57
L56:
	;
	v269 = int32(_a_F_show_all_settings_24)
	goto L57
L57:
	;
	v270 = F_pstrdup(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v270
	goto L44
L59:
	;
	v283 = F_pstrdup(m, v276)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v184)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v286
	v292 = F_pg_snprintf(m, v276, int32(256), int32(_a_F_show_all_settings_22), v12+int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v294 = F_pstrdup(m, v276)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v294
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v299
	v305 = F_pg_snprintf(m, v276, int32(256), int32(_a_F_show_all_settings_22), v12+int32(32))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v307 = F_pstrdup(m, v276)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v307
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v184)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v310
	v316 = F_pg_snprintf(m, v276, int32(256), int32(_a_F_show_all_settings_22), v12+int32(16))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v318 = F_pstrdup(m, v276)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v318
	goto L44
L67:
	;
	v331 = F_pstrdup(m, v324)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v331
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v184)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+112)) = v334
	v340 = F_pg_snprintf(m, v324, int32(256), int32(_a_F_show_all_settings_21), v12+int32(112))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v342 = F_pstrdup(m, v324)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v342
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v184)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+96)) = v347
	v353 = F_pg_snprintf(m, v324, int32(256), int32(_a_F_show_all_settings_21), v12+int32(96))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v355 = F_pstrdup(m, v324)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v355
	v358 = *(*float64)(unsafe.Add(mBase, uint32(v184)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+80)) = v358
	v364 = F_pg_snprintf(m, v324, int32(256), int32(_a_F_show_all_settings_21), v12+int32(80))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v366 = F_pstrdup(m, v324)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v366
	goto L44
L75:
	;
	v374 = F_pstrdup(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	v377 = int32(0)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v184)+112))
	if v379 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v377 = v374
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	goto L44
L80:
	;
	goto L81
L81:
	;
	v384 = F_pstrdup(m, v379)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v384
	goto L44
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = v392
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	v396 = F_config_enum_lookup_by_value(m, v184, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v398 = F_pstrdup(m, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v184)+116))
	v402 = F_config_enum_lookup_by_value(m, v184, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v404 = F_pstrdup(m, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v404
	goto L44
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v441
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v184)+28))
	if v445&int32(2) != 0 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v437 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v437
	v441 = v437
	goto L88
L90:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[1]))
	v421 = F_has_privs_of_role(m, v419, int32(3374))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v421 == int32(0) {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v184)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v184)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v427
	v430 = v12 + int32(224)
	v433 = F_pg_snprintf(m, v430, int32(256), int32(_a_F_show_all_settings_22), v12)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v435 = F_pstrdup(m, v430)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v441 = v435
	goto L88
L95:
	;
	v448 = int32(_a_F_show_all_settings_25)
	goto L97
L96:
	;
	v448 = int32(_a_F_show_all_settings_26)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v448
	v452 = F_BuildTupleFromCStrings(m, v170, v12+int32(144))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v452)+16))
	v455 = F_HeapTupleHeaderGetDatum(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v457 + int64(1)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v461)+20)) = int32(1)
	v466 = v455
	goto L26
}
