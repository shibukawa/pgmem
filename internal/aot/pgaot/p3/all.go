package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAllSchemaPublicationRelations(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = F_GetPublicationSchemas(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v32
L2:
	;
	v13 = v8
	v15 = v2
	goto L9
L3:
	;
	return int32(0)
L4:
	;
	if v4 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v8 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v8 < v9 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v32 = v2
	goto L1
L8:
	;
	goto L7
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v17 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v13<<(uint(v17)%32))))
	v22 = F_GetSchemaPublicationRelations(m, v20, v17)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v32 = v24
	goto L1
L11:
	;
	v24 = F_list_concat(m, v15, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v27 = v13 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v27 < v28 {
		v13 = v27
		v15 = v24
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
	var v26 int32
	_ = v26
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
	v22 = int32(0)
	v26 = v17
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
	v29 = v19 + v22<<(uint(int32(3))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v30 == int32(1776044) {
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
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v26<<(uint(int32(3))%32)+v19-int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v43 = v41 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v43)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
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
	v52 = v21
	v53 = v22
	goto L12
L12:
	;
	v55 = v53 + int32(1)
	v57 = v52 & int32(255)
	if v55 < v57 {
		v21 = v52
		v22 = v55
		v26 = v57
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
	v53 = v22 - int32(1)
	goto L12
L15:
	;
	goto L9
L16:
	;
	v69 = v66
	v70 = int32(0)
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
	v76 = v70 << (uint(int32(3)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v78 = v76 + v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 == int32(1776044) {
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
	v94 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
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
	v100 = v70 + int32(1)
	if base.Ui32(v100) < base.Ui32(v98) {
		v69 = v98
		v70 = v100
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
	v119 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v119
	F_errmsg_internal(m, int32(467246), v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(521095), int32(822), int32(447451))
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
	*(*int32)(unsafe.Add(mBase, _consts[628])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[629])) = v2
	*(*uint8)(unsafe.Add(mBase, _consts[631])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[637])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1164])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1165])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1166])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1167])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[912])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1168])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1169])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1170])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1171])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1172])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1173])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1174])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1175])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1176])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1177])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1178])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1179])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1180])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1181])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1182])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[914])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1183])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1184])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1185])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1186])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1187])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1188])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1189])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1190])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1191])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1192])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1193])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1194])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1195])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1196])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1197])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1198])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1199])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1200])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1201])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1202])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1203])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1204])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _consts[1205])) = uint8(v2)
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
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 float64
	_ = v343
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 float64
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 float64
	_ = v375
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 float64
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int64
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
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
	v22 = int32(4562096)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
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
	F_TupleDescInitEntry(m, v28, int32(1), int32(401664), int32(25), int32(-1), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v28, int32(2), int32(346631), int32(25), int32(-1), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v28, int32(3), int32(106627), int32(25), int32(-1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v28, int32(4), int32(14351), int32(25), int32(-1), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v28, int32(5), int32(512819), int32(25), int32(-1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v28, int32(6), int32(512830), int32(25), int32(-1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v28, int32(7), int32(67887), int32(25), int32(-1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v28, int32(8), int32(385351), int32(25), int32(-1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v28, int32(9), int32(435750), int32(25), int32(-1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v28, int32(10), int32(325620), int32(25), int32(-1), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v28, int32(11), int32(325593), int32(25), int32(-1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v28, int32(12), int32(163502), int32(1009), int32(-1), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v28, int32(13), int32(325601), int32(25), int32(-1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v28, int32(14), int32(325610), int32(25), int32(-1), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v28, int32(15), int32(406402), int32(25), int32(-1), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v28, int32(16), int32(393032), int32(23), int32(-1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_TupleDescInitEntry(m, v28, int32(17), int32(88431), int32(16), int32(-1), int32(0))
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	goto L3
L26:
	;
	m.G0 = v12 + int32(480)
	return v502
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
	v193 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
	v502 = int32(0)
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
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231<<(uint(v232)%32))+uint32(_consts[1159])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242<<(uint(v232)%32))+uint32(_consts[1160])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249<<(uint(v232)%32))+uint32(_consts[1161])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v184)+32))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256<<(uint(v232)%32))+uint32(_consts[1162])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v261
	switch v249 {
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
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v184)+32))
	if v450 != int32(3) {
		goto L89
	} else {
		goto L90
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	v445 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+188)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = v445
	goto L44
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v428 = F_config_enum_get_options(m, v184, int32(724216), int32(6887), int32(767671))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L83
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	if v409 != 0 {
		goto L75
	} else {
		goto L76
	}
L48:
	;
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v184)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+128)) = v343
	v351 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(356304), v12+int32(128))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L67
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v184)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v281
	v289 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(512558), v12-int32(-64))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L59
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+96)))
	if v269 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v270 = int32(287119)
	goto L53
L52:
	;
	v270 = int32(357125)
	goto L53
L53:
	;
	v271 = F_pstrdup(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v271
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+112)))
	if v276 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v277 = int32(287119)
	goto L57
L56:
	;
	v277 = int32(357125)
	goto L57
L57:
	;
	v278 = F_pstrdup(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v278
	goto L44
L59:
	;
	v293 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v293
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v184)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v296
	v304 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(512558), v12+int32(48))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v308 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v308
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v313
	v321 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(512558), v12+int32(32))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v325 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v325
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v184)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v328
	v336 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(512558), v12+int32(16))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v340 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v340
	goto L44
L67:
	;
	v355 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v355
	v358 = *(*float64)(unsafe.Add(mBase, uint32(v184)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+112)) = v358
	v366 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(356304), v12+int32(112))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v370 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v370
	v375 = *(*float64)(unsafe.Add(mBase, uint32(v184)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+96)) = v375
	v383 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(356304), v12+int32(96))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v387 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v387
	v390 = *(*float64)(unsafe.Add(mBase, uint32(v184)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+80)) = v390
	v398 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(356304), v12+int32(80))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v402 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v402
	goto L44
L75:
	;
	v410 = F_pstrdup(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	v413 = int32(0)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v184)+112))
	if v415 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v413 = v410
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
	v420 = F_pstrdup(m, v415)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v420
	goto L44
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = v428
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v184)+96))
	v432 = F_config_enum_lookup_by_value(m, v184, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v434 = F_pstrdup(m, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v434
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v184)+116))
	v438 = F_config_enum_lookup_by_value(m, v184, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v440 = F_pstrdup(m, v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v440
	goto L44
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v477
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v184)+28))
	if v481&int32(2) != 0 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v474 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v474
	v477 = v474
	goto L88
L90:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v456 = F_has_privs_of_role(m, v454, int32(3374))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v456 == int32(0) {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v184)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v184)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v462
	v468 = F_pg_snprintf(m, v12+int32(224), int32(256), int32(512558), v12)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v472 = F_pstrdup(m, v12+int32(224))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v477 = v472
	goto L88
L95:
	;
	v484 = int32(120589)
	goto L97
L96:
	;
	v484 = int32(358482)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v484
	v488 = F_BuildTupleFromCStrings(m, v170, v12+int32(144))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488)+16))
	v491 = F_HeapTupleHeaderGetDatum(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v493 + int64(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v497)+20)) = int32(1)
	v502 = v491
	goto L26
}
