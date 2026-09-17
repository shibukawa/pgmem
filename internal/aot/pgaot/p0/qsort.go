package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_partition_list_value_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = F_FunctionCall2Coll(m, v4, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_qsort_tuple_unsigned(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
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
	var v435 int64
	_ = v435
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v513 int64
	_ = v513
	var v516 int32
	_ = v516
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v565 int64
	_ = v565
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v577 int32
	_ = v577
	var v597 int32
	_ = v597
	var v609 int32
	_ = v609
	var v614 int64
	_ = v614
	var v616 int64
	_ = v616
	var v618 int64
	_ = v618
	var v620 int64
	_ = v620
	var v622 int64
	_ = v622
	var v624 int64
	_ = v624
	var v626 int32
	_ = v626
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = l0
	v20 = l1
	goto L1
L1:
	;
	v34 = v19 + int32(16)
	v36 = v20
	goto L3
L2:
	;
	m.G0 = v17 + int32(16)
	return
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple_unsigned[0]))
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v55 = v19 + v36<<(uint(int32(4))%32)
	if base.Ui32(v36) <= base.Ui32(int32(6)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v164 = v34
	goto L43
L14:
	;
	v71 = v34
	goto L15
L15:
	;
	if base.Ui32(v71) <= base.Ui32(v19) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v159 = v71 + int32(16)
	if base.Ui32(v159) < base.Ui32(v55) {
		v71 = v159
		goto L15
	} else {
		goto L42
	}
L18:
	;
	v79 = v71
	goto L19
L19:
	;
	v90 = v79 - int32(16)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79-int32(8)))))
	if v95 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L17
L21:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v90)+8)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = v141
	if base.Ui32(v19) < base.Ui32(v90) {
		v79 = v90
		goto L19
	} else {
		goto L41
	}
L22:
	;
	if v125 <= int32(0) {
		goto L17
	} else {
		goto L40
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v121 != 0 {
		goto L17
	} else {
		goto L38
	}
L24:
	;
	if v92&int32(1) != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v92&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
	if v100 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
	if v105 != 0 {
		goto L21
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79-int32(12))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v111 = base.B2i32(base.Ui32(v108) < base.Ui32(v109))
	v112 = base.B2i32(base.Ui32(v109) < base.Ui32(v108)) - v111
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
	if v113 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L17
L33:
	;
	if base.Ui32(v108) < base.Ui32(v109) {
		goto L21
	} else {
		goto L36
	}
L34:
	;
	v118 = v112
	goto L35
L35:
	;
	if v118 != 0 {
		v125 = v118
		goto L22
	} else {
		goto L37
	}
L36:
	;
	v118 = int32(0) - v112
	goto L35
L37:
	;
	goto L23
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v123 = m.T0[v122].(func(*base.Module, int32, int32, int32) int32)(m, v90, v79, l2)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v125 = v123
	goto L22
L40:
	;
	goto L21
L41:
	;
	goto L20
L42:
	;
	goto L16
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple_unsigned[0]))
	if v176 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v232 = v19 + v36<<(uint(int32(3))%32)&int32(-16)
	if v36 != int32(7) {
		goto L71
	} else {
		goto L72
	}
L45:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164-int32(8)))))
	if v183 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	goto L44
L50:
	;
	v224 = v164 + int32(16)
	if base.Ui32(v224) < base.Ui32(v55) {
		v164 = v224
		goto L43
	} else {
		goto L70
	}
L51:
	;
	if int32(0) < v217 {
		goto L49
	} else {
		goto L69
	}
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v211 != 0 {
		goto L50
	} else {
		goto L67
	}
L53:
	;
	if v180&int32(1) != 0 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v180&int32(1) != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+9)))
	if v188 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+9)))
	if v193 == int32(0) {
		goto L50
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(12))))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v201 = base.B2i32(base.Ui32(v198) < base.Ui32(v199))
	v202 = base.B2i32(base.Ui32(v199) < base.Ui32(v198)) - v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+8)))
	if v203 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L49
L62:
	;
	if base.Ui32(v198) < base.Ui32(v199) {
		goto L49
	} else {
		goto L65
	}
L63:
	;
	v208 = v202
	goto L64
L64:
	;
	if v208 != 0 {
		v217 = v208
		goto L51
	} else {
		goto L66
	}
L65:
	;
	v208 = int32(0) - v202
	goto L64
L66:
	;
	goto L52
L67:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v215 = m.T0[v214].(func(*base.Module, int32, int32, int32) int32)(m, v164-int32(16), v164, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v217 = v215
	goto L51
L69:
	;
	goto L50
L70:
	;
	goto L10
L71:
	;
	v236 = v55 - int32(16)
	if base.Ui32(v36) < base.Ui32(int32(41)) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v265 = v232
	goto L73
L73:
	;
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v265)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v273
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v265)))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v275
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v265)+8)) = v277
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v265))) = v279
	v282 = v55 - int32(16)
	v286 = v282
	v287 = v34
	v288 = v34
	v290 = v282
	goto L82
L74:
	;
	v262 = F_qsort_tuple_unsigned_med3(m, v257, v258, v259, l2)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L81
	}
L75:
	;
	v257 = v19
	v258 = v232
	v259 = v236
	goto L74
L76:
	;
	goto L77
L77:
	;
	v240 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v242 = v240 << (uint(int32(4)) % 32)
	v245 = v240 << (uint(int32(5)) % 32)
	v247 = F_qsort_tuple_unsigned_med3(m, v19, v19+v242, v19+v245, l2)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v251 = F_qsort_tuple_unsigned_med3(m, v232-v242, v232, v232+v242, l2)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v255 = F_qsort_tuple_unsigned_med3(m, v236-v245, v236-v242, v236, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v257 = v247
	v258 = v251
	v259 = v255
	goto L74
L81:
	;
	v265 = v262
	goto L73
L82:
	;
	if base.Ui32(v286) < base.Ui32(v287) {
		v376 = v287
		v377 = v288
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if base.Ui32(v376) <= base.Ui32(v286) {
		goto L116
	} else {
		goto L117
	}
L85:
	;
	v302 = v287
	v303 = v288
	goto L86
L86:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+8)))
	if v314 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v376 = v370
	v377 = v363
	goto L84
L88:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple_unsigned[0]))
	if v366 != 0 {
		goto L110
	} else {
		goto L111
	}
L89:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v303)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v348
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v350
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v302)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+8)) = v352
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v302)))
	*(*int64)(unsafe.Add(mBase, uint32(v303))) = v354
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v302)+8)) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v302))) = v358
	v363 = v303 + int32(16)
	goto L88
L90:
	;
	if int32(0) < v342 {
		v376 = v302
		v377 = v303
		goto L84
	} else {
		goto L108
	}
L91:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v338 != 0 {
		goto L89
	} else {
		goto L106
	}
L92:
	;
	if v313&int32(1) != 0 {
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v313&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+9)))
	if v319 != 0 {
		v363 = v303
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v376 = v302
	v377 = v303
	goto L84
L97:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+9)))
	if v322 == int32(0) {
		v363 = v303
		goto L88
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v328 = base.B2i32(base.Ui32(v325) < base.Ui32(v326))
	v329 = base.B2i32(base.Ui32(v326) < base.Ui32(v325)) - v328
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+8)))
	if v330 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v376 = v302
	v377 = v303
	goto L84
L101:
	;
	if base.Ui32(v325) < base.Ui32(v326) {
		v376 = v302
		v377 = v303
		goto L84
	} else {
		goto L104
	}
L102:
	;
	v335 = v329
	goto L103
L103:
	;
	if v335 != 0 {
		v342 = v335
		goto L90
	} else {
		goto L105
	}
L104:
	;
	v335 = int32(0) - v329
	goto L103
L105:
	;
	goto L91
L106:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v340 = m.T0[v339].(func(*base.Module, int32, int32, int32) int32)(m, v302, v19, l2)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	v342 = v340
	goto L90
L108:
	;
	if v342 != 0 {
		v363 = v303
		goto L88
	} else {
		goto L109
	}
L109:
	;
	goto L89
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v370 = v302 + int32(16)
	if base.Ui32(v370) <= base.Ui32(v286) {
		v302 = v370
		v303 = v363
		goto L86
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	goto L87
L115:
	;
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v376)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v614
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v616
	v618 = *(*int64)(unsafe.Add(mBase, uint32(v390)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v376)+8)) = v618
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v390)))
	*(*int64)(unsafe.Add(mBase, uint32(v376))) = v620
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v390)+8)) = v622
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v390))) = v624
	v626 = int32(16)
	v286 = v390 - v626
	v287 = v376 + v626
	v288 = v377
	v290 = v394
	goto L82
L116:
	;
	v390 = v286
	v394 = v290
	goto L119
L117:
	;
	v462 = v286
	v466 = v290
	goto L118
L118:
	;
	v474 = int32(4)
	v475 = (v377 - v19) >> (uint(v474) % 32)
	v478 = (v376 - v377) >> (uint(v474) % 32)
	if v475 < v478 {
		goto L148
	} else {
		goto L149
	}
L119:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+8)))
	if v403 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v462 = v457
	v466 = v450
	goto L118
L121:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple_unsigned[0]))
	if v453 != 0 {
		goto L143
	} else {
		goto L144
	}
L122:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v390)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v435
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v390)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v437
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v394)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v390)+8)) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int64)(unsafe.Add(mBase, uint32(v390))) = v441
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v445
	v450 = v394 - int32(16)
	goto L121
L123:
	;
	if v429 < int32(0) {
		goto L115
	} else {
		goto L141
	}
L124:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v425 != 0 {
		goto L122
	} else {
		goto L139
	}
L125:
	;
	if v402&int32(1) != 0 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v402&int32(1) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+9)))
	if v408 != 0 {
		goto L115
	} else {
		goto L129
	}
L129:
	;
	v450 = v394
	goto L121
L130:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+9)))
	if v411 != 0 {
		v450 = v394
		goto L121
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v415 = base.B2i32(base.Ui32(v412) < base.Ui32(v413))
	v416 = base.B2i32(base.Ui32(v413) < base.Ui32(v412)) - v415
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+8)))
	if v417 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L115
L134:
	;
	if base.Ui32(v412) < base.Ui32(v413) {
		v450 = v394
		goto L121
	} else {
		goto L137
	}
L135:
	;
	v422 = v416
	goto L136
L136:
	;
	if v422 != 0 {
		v429 = v422
		goto L123
	} else {
		goto L138
	}
L137:
	;
	v422 = int32(0) - v416
	goto L136
L138:
	;
	goto L124
L139:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v427 = m.T0[v426].(func(*base.Module, int32, int32, int32) int32)(m, v390, v19, l2)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	v429 = v427
	goto L123
L141:
	;
	if v429 != 0 {
		v450 = v394
		goto L121
	} else {
		goto L142
	}
L142:
	;
	goto L122
L143:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L8
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v457 = v390 - int32(16)
	if base.Ui32(v376) <= base.Ui32(v457) {
		v390 = v457
		v394 = v450
		goto L119
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	goto L120
L148:
	;
	v480 = v475
	goto L150
L149:
	;
	v480 = v478
	goto L150
L150:
	;
	if v480 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v490 = int32(0)
	goto L154
L152:
	;
	goto L153
L153:
	;
	v533 = int32(4)
	v534 = (v466 - v462) >> (uint(v533) % 32)
	v539 = (v55-v466)>>(uint(v533)%32) - int32(1)
	if v534 < v539 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v500 = v490 << (uint(int32(4)) % 32)
	v501 = v19 + v500
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v501)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v502
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v501)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v504
	v506 = v500 + (v376 - v480<<(uint(int32(4))%32))
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v506)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v501)+8)) = v507
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v506)))
	*(*int64)(unsafe.Add(mBase, uint32(v501))) = v509
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+8)) = v511
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v506))) = v513
	v516 = v490 + int32(1)
	if v516 != v480 {
		v490 = v516
		goto L154
	} else {
		goto L156
	}
L155:
	;
	goto L153
L156:
	;
	goto L155
L157:
	;
	v541 = v534
	goto L159
L158:
	;
	v541 = v539
	goto L159
L159:
	;
	if v541 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v551 = int32(0)
	goto L163
L161:
	;
	goto L162
L162:
	;
	if base.Ui32(v478) <= base.Ui32(v534) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v561 = v551 << (uint(int32(4)) % 32)
	v562 = v376 + v561
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v562)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v563
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v562)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v565
	v567 = v561 + (v55 - v541<<(uint(int32(4))%32))
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+8)) = v568
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v567)))
	*(*int64)(unsafe.Add(mBase, uint32(v562))) = v570
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v567)+8)) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v567))) = v574
	v577 = v551 + int32(1)
	if v577 != v541 {
		v551 = v577
		goto L163
	} else {
		goto L165
	}
L164:
	;
	goto L162
L165:
	;
	goto L164
L166:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v478) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v534) {
		goto L174
	} else {
		goto L175
	}
L169:
	;
	F_qsort_tuple_unsigned(m, v19, v478, l2)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L8
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(v534) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v19 = v55 - v534<<(uint(int32(4))%32)
	v20 = v534
	goto L1
L174:
	;
	F_qsort_tuple_unsigned(m, v55-v534<<(uint(int32(4))%32), v534, l2)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	if base.Ui32(int32(1)) < base.Ui32(v478) {
		v36 = v478
		goto L3
	} else {
		goto L178
	}
L177:
	;
	goto L176
L178:
	;
	goto L10
}
func F_qsort_tuple_unsigned_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
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
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v413 int32
	_ = v413
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v16 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L1:
	;
	return v413
L2:
	;
	v413 = l0
	goto L1
L3:
	;
	return l2
L4:
	;
	if int32(0) <= v385 {
		v413 = l0
		goto L1
	} else {
		goto L115
	}
L5:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v380 != 0 {
		goto L2
	} else {
		goto L113
	}
L6:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v363 = base.B2i32(base.Ui32(v361) < base.Ui32(v335))
	v364 = base.B2i32(base.Ui32(v335) < base.Ui32(v361)) - v363
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
	if v365 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L7:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+9)))
	if v358 == int32(0) {
		goto L3
	} else {
		goto L107
	}
L8:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v342 != 0 {
		goto L101
	} else {
		goto L102
	}
L9:
	;
	if v325 < int32(0) {
		v413 = l1
		goto L1
	} else {
		goto L100
	}
L10:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v317 != 0 {
		v334 = v309
		v335 = v310
		v336 = v311
		goto L8
	} else {
		goto L98
	}
L11:
	;
	v298 = base.B2i32(base.Ui32(v293) < base.Ui32(v290))
	v299 = base.B2i32(base.Ui32(v290) < base.Ui32(v293)) - v298
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	if v300 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L12:
	;
	if int32(0) <= v279 {
		v413 = l2
		goto L1
	} else {
		goto L92
	}
L13:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v273 != 0 {
		goto L88
	} else {
		goto L89
	}
L14:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v255 = base.B2i32(base.Ui32(v253) < base.Ui32(v246))
	v256 = base.B2i32(base.Ui32(v246) < base.Ui32(v253)) - v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+8)))
	if v257 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L15:
	;
	if v221&int32(1) == int32(0) {
		v246 = v217
		v247 = v218
		goto L14
	} else {
		goto L79
	}
L16:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+9)))
	if v237 != 0 {
		goto L2
	} else {
		goto L78
	}
L17:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v224 == int32(0) {
		goto L15
	} else {
		goto L76
	}
L18:
	;
	if int32(0) < v203 {
		v413 = l1
		goto L1
	} else {
		goto L75
	}
L19:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v199 != 0 {
		v217 = v192
		v218 = v193
		v221 = v196
		goto L17
	} else {
		goto L73
	}
L20:
	;
	v182 = base.B2i32(base.Ui32(v177) < base.Ui32(v174))
	v183 = base.B2i32(base.Ui32(v174) < base.Ui32(v177)) - v182
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
	if v184 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L21:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
	if v172 != 0 {
		v413 = l1
		goto L1
	} else {
		goto L67
	}
L22:
	;
	if v90&int32(1) == int32(0) {
		v174 = v91
		v175 = v82
		v177 = v84
		v179 = v86
		v180 = v88
		goto L20
	} else {
		goto L66
	}
L23:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v154 != 0 {
		v164 = v153
		v165 = v13
		goto L21
	} else {
		goto L65
	}
L24:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+9)))
	if v145 == int32(0) {
		v413 = l1
		goto L1
	} else {
		goto L63
	}
L25:
	;
	if v115&int32(1) == int32(0) {
		v290 = v116
		v291 = v74
		v293 = v76
		v295 = v111
		v296 = v113
		goto L11
	} else {
		goto L62
	}
L26:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+9)))
	if v132 != 0 {
		v413 = l1
		goto L1
	} else {
		goto L61
	}
L27:
	;
	v111 = l2 + int32(8)
	v113 = l2 + int32(4)
	v114 = int32(1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v77&v114 == int32(0) {
		goto L25
	} else {
		goto L59
	}
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+9)))
	if v106 == int32(0) {
		v413 = l1
		goto L1
	} else {
		goto L57
	}
L29:
	;
	v86 = l2 + int32(8)
	v88 = l2 + int32(4)
	v89 = int32(1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v80&v89 == int32(0) {
		goto L22
	} else {
		goto L55
	}
L30:
	;
	if v73 < int32(0) {
		goto L27
	} else {
		goto L54
	}
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v64 != 0 {
		v80 = v14
		v82 = v13
		v84 = v15
		goto L29
	} else {
		goto L51
	}
L32:
	;
	if v14&int32(1) != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v14&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)))
	if v21 != int32(1) {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v24 != 0 {
		v139 = v13
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v290 = v29
	v291 = v13
	v293 = v15
	v295 = l2 + int32(8)
	v296 = l2 + int32(4)
	goto L11
L38:
	;
	v33 = l2 + int32(8)
	v35 = l2 + int32(4)
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)))
	if v39 == v36 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = base.B2i32(base.Ui32(v51) < base.Ui32(v15))
	v54 = base.B2i32(base.Ui32(v15) < base.Ui32(v51)) - v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
	if v55 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v42 = int32(1)
	if v37&v42 == int32(0) {
		v99 = v38
		v100 = v13
		goto L28
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v37&int32(1) == int32(0) {
		v124 = v38
		v125 = v13
		goto L26
	} else {
		goto L45
	}
L44:
	;
	v192 = v38
	v193 = v13
	v196 = v42
	v197 = v33
	v198 = v35
	goto L19
L45:
	;
	v309 = v36
	v310 = v38
	v311 = v13
	v315 = v33
	v316 = v35
	goto L10
L46:
	;
	if base.Ui32(v51) < base.Ui32(v15) {
		goto L23
	} else {
		goto L49
	}
L47:
	;
	v60 = v54
	goto L48
L48:
	;
	if v60 != 0 {
		v73 = v60
		v74 = v13
		v76 = v15
		v77 = int32(0)
		goto L30
	} else {
		goto L50
	}
L49:
	;
	v60 = int32(0) - v54
	goto L48
L50:
	;
	goto L31
L51:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v66 = m.T0[v65].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v73 = v66
	v74 = v70
	v76 = v71
	v77 = v72
	goto L30
L54:
	;
	v80 = v77
	v82 = v74
	v84 = v76
	goto L29
L55:
	;
	if v90&int32(1) != 0 {
		v192 = v91
		v193 = v82
		v196 = v89
		v197 = v86
		v198 = v88
		goto L19
	} else {
		goto L56
	}
L56:
	;
	v99 = v91
	v100 = v82
	goto L28
L57:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v109 != 0 {
		v231 = v100
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v246 = v99
	v247 = v100
	goto L14
L59:
	;
	if v115&int32(1) != 0 {
		v309 = v114
		v310 = v116
		v311 = v74
		v315 = v111
		v316 = v113
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v124 = v116
	v125 = v74
	goto L26
L61:
	;
	v334 = int32(0)
	v335 = v124
	v336 = v125
	goto L8
L62:
	;
	v139 = v74
	goto L24
L63:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v148 == int32(0) {
		v352 = v139
		goto L7
	} else {
		goto L64
	}
L64:
	;
	goto L5
L65:
	;
	v174 = v153
	v175 = v13
	v177 = v15
	v179 = l2 + int32(8)
	v180 = l2 + int32(4)
	goto L20
L66:
	;
	v164 = v91
	v165 = v82
	goto L21
L67:
	;
	v217 = v164
	v218 = v165
	v221 = int32(1)
	goto L17
L68:
	;
	if base.Ui32(v177) < base.Ui32(v174) {
		v413 = l1
		goto L1
	} else {
		goto L71
	}
L69:
	;
	v189 = v183
	goto L70
L70:
	;
	if v189 != 0 {
		v203 = v189
		v209 = v179
		v210 = v180
		goto L18
	} else {
		goto L72
	}
L71:
	;
	v189 = int32(0) - v183
	goto L70
L72:
	;
	v192 = v174
	v193 = v175
	v196 = int32(0)
	v197 = v179
	v198 = v180
	goto L19
L73:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v201 = m.T0[v200].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L52
	} else {
		goto L74
	}
L74:
	;
	v203 = v201
	v209 = v197
	v210 = v198
	goto L18
L75:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = v215
	v218 = v214
	v221 = v213
	goto L17
L76:
	;
	if v221&int32(1) != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v231 = v218
	goto L16
L78:
	;
	v413 = l2
	goto L1
L79:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
	if v242 == int32(0) {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v413 = l2
	goto L1
L81:
	;
	if base.Ui32(v253) < base.Ui32(v246) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v263 = v256
	goto L83
L83:
	;
	if v263 != 0 {
		v279 = v263
		goto L12
	} else {
		goto L87
	}
L84:
	;
	return l2
L85:
	;
	goto L86
L86:
	;
	v263 = int32(0) - v256
	goto L83
L87:
	;
	goto L13
L88:
	;
	return l2
L89:
	;
	goto L90
L90:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v276 = m.T0[v275].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l3)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L52
	} else {
		goto L91
	}
L91:
	;
	v279 = v276
	goto L12
L92:
	;
	goto L2
L93:
	;
	if base.Ui32(v293) < base.Ui32(v290) {
		v334 = int32(0)
		v335 = v290
		v336 = v291
		goto L8
	} else {
		goto L96
	}
L94:
	;
	v307 = v299
	goto L95
L95:
	;
	if v307 != 0 {
		v325 = v307
		v327 = v295
		v328 = v296
		goto L9
	} else {
		goto L97
	}
L96:
	;
	v307 = int32(0) - v299
	goto L95
L97:
	;
	v309 = int32(0)
	v310 = v290
	v311 = v291
	v315 = v295
	v316 = v296
	goto L10
L98:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v319 = m.T0[v318].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L52
	} else {
		goto L99
	}
L99:
	;
	v325 = v319
	v327 = v315
	v328 = v316
	goto L9
L100:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v334 = v331
	v335 = v333
	v336 = v332
	goto L8
L101:
	;
	if v334&int32(1) != 0 {
		goto L5
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v334&int32(1) == int32(0) {
		goto L6
	} else {
		goto L106
	}
L104:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
	if v345 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	v413 = l0
	goto L1
L106:
	;
	v352 = v336
	goto L7
L107:
	;
	v413 = l0
	goto L1
L108:
	;
	if base.Ui32(v361) < base.Ui32(v335) {
		goto L2
	} else {
		goto L111
	}
L109:
	;
	v370 = v364
	goto L110
L110:
	;
	if v370 != 0 {
		v385 = v370
		goto L4
	} else {
		goto L112
	}
L111:
	;
	v370 = int32(0) - v364
	goto L110
L112:
	;
	goto L5
L113:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v382 = m.T0[v381].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l3)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L52
	} else {
		goto L114
	}
L114:
	;
	v385 = v382
	goto L4
L115:
	;
	goto L3
}
