package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsa_create_ext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = F_dsm_create(m, l1, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_dsm_pin_segment(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v14 = F_create_internal(m, v12, l1, l0, v13, v6, l1, l2)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
				F_on_dsm_detach(m, v6, int32(1770), v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v14
				}
			}
		}
	}
}
func F_dsa_free(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
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
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
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
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	v12 = l1
	goto L1
L1:
	;
	v21 = int32(0)
	v24 = base.AtomicRmwOr32(m, v21, int32(_a_F_dsa_free_0), v21)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+1468))
	if v25 != v27 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = F_LWLockAcquire(m, v201+v163<<(uint(int32(5))%32)+int32(224), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L60
	}
L3:
	;
	v32 = F_LWLockAcquire(m, v26+int32(1476), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = int32(0)
	v44 = int32(base.Ui32(v12) >> (uint(int32(27)) % 32))
	v45 = F_get_segment_by_index(m, l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L10
	}
L6:
	;
	return
L7:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v36+int32(1476))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(base.Ui32(v12)>>(uint(int32(10))%32))&int32(_a_F_dsa_free_1))))
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = int32(0)
	v57 = base.AtomicRmwOr32(m, v54, int32(_a_F_dsa_free_0), v54)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+1468))
	if v58 != v60 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v87 = v41
	goto L13
L13:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v91 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v65 = F_LWLockAcquire(m, v59+int32(1476), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v77 = int32(base.Ui32(v53) >> (uint(int32(27)) % 32))
	v80 = l0 + v77*int32(20)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v81 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v69+int32(1476))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v85 = v81
	goto L22
L21:
	;
	v82 = F_get_segment_by_index(m, l0, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L23
	}
L22:
	;
	v87 = v85 + v53&int32(134217727)
	goto L13
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v85 = v84
	goto L22
L24:
	;
	v92 = int32(0)
	v95 = base.AtomicRmwOr32(m, v92, int32(_a_F_dsa_free_0), v92)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+1468))
	if v96 != v98 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v125 = v41
	goto L26
L26:
	;
	if v12 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v103 = F_LWLockAcquire(m, v97+int32(1476), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v113 = int32(base.Ui32(v91) >> (uint(int32(27)) % 32))
	v116 = l0 + v113*int32(20)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if v117 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v107+int32(1476))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v121 = v117
	goto L35
L34:
	;
	v118 = F_get_segment_by_index(m, l0, v113)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v125 = v121 + v91&int32(134217727)
	goto L26
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v121 = v120
	goto L35
L37:
	;
	v128 = int32(0)
	v131 = base.AtomicRmwOr32(m, v128, int32(_a_F_dsa_free_0), v128)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+1468))
	if v132 != v134 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v162 = int32(0)
	goto L39
L39:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+20)))
	if v163 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v139 = F_LWLockAcquire(m, v133+int32(1476), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v150 = l0 + v44*int32(20)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	if v151 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v143+int32(1476))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v155 = v151
	goto L48
L47:
	;
	v152 = F_get_segment_by_index(m, l0, v44)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	v162 = v155 + v12&int32(134217727)
	goto L39
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v155 = v154
	goto L48
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v170 = F_LWLockAcquire(m, v166+int32(1476), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L2
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	F_FreePageManagerPut(m, v172, int32(base.Ui32(v173)>>(uint(int32(12))%32))&int32(_a_F_dsa_free_2), v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_rebin_segment(m, l0, v45)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v183+int32(1476))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = F_LWLockAcquire(m, v188+int32(256), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_unlink_span(m, l0, v87)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v196+int32(256))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v12 = v53
	goto L1
L60:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v162))) = uint16(v210)
	v213 = int32(1)
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163<<(uint(v213)%32))+uint32(_c_F_dsa_free[0]))))
	v216 = base.I32_div_u_s(v162-v125, v215)
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+26)) = uint16(v216)
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+24)))
	v220 = v218 + v213
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+24)) = uint16(v220)
	if v218 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v520+v163<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L6
	} else {
		goto L145
	}
L62:
	;
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+28)))
	if v312 != v220&int32(_a_F_dsa_free_3) {
		goto L61
	} else {
		goto L92
	}
L63:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+30)))
	if v222 != int32(3) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_unlink_span(m, l0, v87)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v227 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v228 = int32(0)
	v231 = base.AtomicRmwOr32(m, v228, int32(_a_F_dsa_free_0), v228)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+1468))
	if v232 != v234 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v266 = int32(0)
	goto L68
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	if v267 != 0 {
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v239 = F_LWLockAcquire(m, v233+int32(1476), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v251 = int32(base.Ui32(v227) >> (uint(int32(27)) % 32))
	v254 = l0 + v251*int32(20)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	if v255 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v243+int32(1476))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v259 = v255
	goto L77
L76:
	;
	v256 = F_get_segment_by_index(m, l0, v251)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L78
	}
L77:
	;
	v266 = v259 + v227&int32(134217727)
	goto L68
L78:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v259 = v258
	goto L77
L79:
	;
	v268 = int32(0)
	v271 = base.AtomicRmwOr32(m, v268, int32(_a_F_dsa_free_0), v268)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+1468))
	if v272 != v274 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v266)+24)) = v53
	v310 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+30)) = uint16(v310)
	goto L61
L82:
	;
	v279 = F_LWLockAcquire(m, v273+int32(1476), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v289 = int32(base.Ui32(v267) >> (uint(int32(27)) % 32))
	v292 = l0 + v289*int32(20)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	if v293 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v283+int32(1476))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v297 = v293
	goto L90
L89:
	;
	v294 = F_get_segment_by_index(m, l0, v289)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297+v267&int32(134217727))+4)) = v53
	goto L81
L91:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v297 = v296
	goto L90
L92:
	;
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+30)))
	if v316 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v319 == int32(0) {
		goto L61
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v53 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v323 = int32(0)
	v326 = base.AtomicRmwOr32(m, v323, int32(_a_F_dsa_free_0), v323)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+1468))
	if v327 != v329 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v357 = int32(0)
	goto L99
L99:
	;
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+20)))
	F_unlink_span(m, l0, v357)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L110
	}
L100:
	;
	v334 = F_LWLockAcquire(m, v328+int32(1476), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v344 = int32(base.Ui32(v53) >> (uint(int32(27)) % 32))
	v347 = l0 + v344*int32(20)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	if v348 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v338+int32(1476))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	v352 = v348
	goto L108
L107:
	;
	v349 = F_get_segment_by_index(m, l0, v344)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L109
	}
L108:
	;
	v357 = v352 + v53&int32(134217727)
	goto L99
L109:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v352 = v351
	goto L108
L110:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v366 = F_LWLockAcquire(m, v362+int32(1476), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+1468))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v369 != v370 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v376 = int32(0)
	goto L115
L113:
	;
	goto L114
L114:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v419 = F_get_segment_by_index(m, l0, int32(base.Ui32(v416)>>(uint(int32(27))%32)))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L122
	}
L115:
	;
	v387 = l0 + int32(8) + v376*int32(20)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	if v388 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v369
	goto L114
L117:
	;
	v402 = v376 + int32(1)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v402) <= base.Ui32(v403) {
		v376 = v402
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+24)))
	if v391 != int32(1) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	F_dsm_detach(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v387))) = int64(0)
	goto L117
L121:
	;
	goto L116
L122:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v357)+16))
	F_FreePageManagerPut(m, v421, int32(base.Ui32(v422)>>(uint(int32(12))%32))&int32(_a_F_dsa_free_2), v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+28))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v431 != v433 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v503+int32(1476))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L6
	} else {
		goto L140
	}
L125:
	;
	F_rebin_segment(m, l0, v419)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L139
	}
L126:
	;
	v436 = l0 + int32(8)
	if v419 == v436 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	if v438 != int32(-1) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+16))
	if v456 != int32(-1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v441 = F_get_segment_by_index(m, l0, v438)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v432)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v447+v448<<(uint(int32(2))%32))+160)) = v452
	goto L128
L132:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+16)) = v445
	goto L128
L133:
	;
	v459 = F_get_segment_by_index(m, l0, v456)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L136
	}
L134:
	;
	v466 = v455
	goto L135
L135:
	;
	v467 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+24)) = uint8(v467)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+1448))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+1448)) = v470 - v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	F_dsm_unpin_segment(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L137
	}
L136:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v461)+12)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v466 = v465
	goto L135
L137:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	F_dsm_detach(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v484 = base.I32_div_s(v419-v436, int32(5))
	v486 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v482+v484)+32)) = v486
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+1468))
	*(*int32)(unsafe.Add(mBase, uint32(v488)+1468)) = v489 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+8)) = v486
	*(*int64)(unsafe.Add(mBase, uint32(v419))) = int64(0)
	goto L124
L139:
	;
	goto L124
L140:
	;
	if v359 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_dsa_free(m, l0, v53)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	goto L61
L144:
	;
	goto L143
L145:
	;
	return
}
func F_dsa_on_dsm_detach_release_in_place(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v3 = int32(0)
	v8 = l1 + int32(1476)
	v10 = F_LWLockAcquire(m, v8, v3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1460))
	v14 = v12 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+1460)) = v14
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_LWLockRelease(m, v8)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(32)+v20<<(uint(int32(2))%32))))
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	F_dsm_unpin_segment(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v32 = v20 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1456))
	if base.Ui32(v32) <= base.Ui32(v33) {
		v20 = v32
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	goto L7
L13:
	;
	return
}
func F_dsa_pin_mapping(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	v11 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v11*int32(20))))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_dsm_pin_mapping(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v21 = v11 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v21) <= base.Ui32(v22) {
		v11 = v21
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	goto L5
}
func F_dsa_set_size_limit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_LWLockAcquire(m, v3+int32(1476), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+1452)) = l1
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_LWLockRelease(m, v11+int32(1476))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
