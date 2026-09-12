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
				F_on_dsm_detach(m, v6, int32(1786), v17)
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
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
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
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
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	v13 = l0 + int32(12)
	v15 = l1
	goto L1
L1:
	;
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = F_LWLockAcquire(m, v190+v152<<(uint(int32(5))%32)+int32(224), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
	v44 = int32(base.Ui32(v15) >> (uint(int32(27)) % 32))
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(base.Ui32(v15)>>(uint(int32(10))%32))&int32(131068))))
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
	if v54 != v56 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v83 = v41
	goto L13
L13:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v87 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v73 = int32(base.Ui32(v53) >> (uint(int32(27)) % 32))
	v76 = v13 + v73*int32(20)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v65+int32(1476))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v81 = v77
	goto L22
L21:
	;
	v78 = F_get_segment_by_index(m, l0, v73)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L23
	}
L22:
	;
	v83 = v81 + v53&int32(134217727)
	goto L13
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v81 = v80
	goto L22
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+1468))
	if v88 != v90 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v118 = v41
	goto L26
L26:
	;
	if v15 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v95 = F_LWLockAcquire(m, v89+int32(1476), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v107 = int32(base.Ui32(v87) >> (uint(int32(27)) % 32))
	v110 = v13 + v107*int32(20)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v99+int32(1476))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v115 = v111
	goto L35
L34:
	;
	v112 = F_get_segment_by_index(m, l0, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v118 = v87&int32(134217727) + v115
	goto L26
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v115 = v114
	goto L35
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+1468))
	if v121 != v123 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v151 = int32(0)
	goto L39
L39:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+20)))
	if v152 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v128 = F_LWLockAcquire(m, v122+int32(1476), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v139 = v13 + v44*int32(20)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v132+int32(1476))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v144 = v140
	goto L48
L47:
	;
	v141 = F_get_segment_by_index(m, l0, v44)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	v151 = v144 + v15&int32(134217727)
	goto L39
L49:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v144 = v143
	goto L48
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = F_LWLockAcquire(m, v155+int32(1476), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	F_FreePageManagerPut(m, v161, int32(base.Ui32(v162)>>(uint(int32(12))%32))&int32(32767), v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_rebin_segment(m, l0, v45)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v172+int32(1476))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v181 = F_LWLockAcquire(m, v177+int32(256), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_unlink_span(m, l0, v83)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v185+int32(256))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v15 = v53
	goto L1
L60:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v199)
	v202 = int32(1)
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152<<(uint(v202)%32))+uint32(_consts[1183]))))
	v207 = base.I32_div_u_s(v151-v118, v206)
	*(*uint16)(unsafe.Add(mBase, uint32(v83)+26)) = uint16(v207)
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+24)))
	v211 = v209 + v202
	*(*uint16)(unsafe.Add(mBase, uint32(v83)+24)) = uint16(v211)
	if v209 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v510+v152<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L145
	}
L62:
	;
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+28)))
	if v298 != v211&int32(65535) {
		goto L61
	} else {
		goto L92
	}
L63:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+30)))
	if v213 != int32(3) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_unlink_span(m, l0, v83)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v218 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+1468))
	if v219 != v221 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v253 = int32(0)
	goto L68
L68:
	;
	v255 = v253 + int32(24)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v256 != 0 {
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v226 = F_LWLockAcquire(m, v220+int32(1476), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v238 = int32(base.Ui32(v218) >> (uint(int32(27)) % 32))
	v241 = v13 + v238*int32(20)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v242 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v230+int32(1476))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v246 = v242
	goto L77
L76:
	;
	v243 = F_get_segment_by_index(m, l0, v238)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L78
	}
L77:
	;
	v253 = v246 + v218&int32(134217727)
	goto L68
L78:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v246 = v245
	goto L77
L79:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+1468))
	if v257 != v259 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v53
	v296 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v83)+30)) = uint16(v296)
	goto L61
L82:
	;
	v264 = F_LWLockAcquire(m, v258+int32(1476), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v276 = int32(base.Ui32(v256) >> (uint(int32(27)) % 32))
	v279 = v13 + v276*int32(20)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v280 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v268+int32(1476))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v284 = v280
	goto L90
L89:
	;
	v281 = F_get_segment_by_index(m, l0, v276)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256&int32(134217727)+v284)+4)) = v53
	goto L81
L91:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v284 = v283
	goto L90
L92:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+30)))
	if v302 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v305 == int32(0) {
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
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+1468))
	if v309 != v311 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v341 = int32(0)
	goto L99
L99:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+20)))
	F_unlink_span(m, l0, v341)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L110
	}
L100:
	;
	v316 = F_LWLockAcquire(m, v310+int32(1476), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v326 = int32(base.Ui32(v53) >> (uint(int32(27)) % 32))
	v331 = l0 + v326*int32(20) + int32(12)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v332 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v320+int32(1476))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	v336 = v332
	goto L108
L107:
	;
	v333 = F_get_segment_by_index(m, l0, v326)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L109
	}
L108:
	;
	v341 = v336 + v53&int32(134217727)
	goto L99
L109:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v336 = v335
	goto L108
L110:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v351 = F_LWLockAcquire(m, v347+int32(1476), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+1468))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v354 != v355 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v364 = int32(0)
	goto L115
L113:
	;
	goto L114
L114:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v406 = F_get_segment_by_index(m, l0, int32(base.Ui32(v403)>>(uint(int32(27))%32)))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L122
	}
L115:
	;
	v373 = l0 + int32(8) + v364*int32(20)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	if v374 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v354
	goto L114
L117:
	;
	v388 = v364 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v388) <= base.Ui32(v389) {
		v364 = v388
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+24)))
	if v377 != int32(1) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	F_dsm_detach(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v373))) = int64(0)
	goto L117
L121:
	;
	goto L116
L122:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v341)+16))
	F_FreePageManagerPut(m, v408, int32(base.Ui32(v409)>>(uint(int32(12))%32))&int32(32767), v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v418 != v420 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v492+int32(1476))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L140
	}
L125:
	;
	F_rebin_segment(m, l0, v406)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L139
	}
L126:
	;
	v423 = l0 + int32(8)
	if v406 == v423 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v427 = base.I32_div_s(v406-v423, int32(20))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	if v428 != int32(-1) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	if v445 != int32(-1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v431 = F_get_segment_by_index(m, l0, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v419)+20))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v438<<(uint(int32(2))%32))+160)) = v442
	goto L128
L132:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+16)) = v435
	goto L128
L133:
	;
	v448 = F_get_segment_by_index(m, l0, v445)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L136
	}
L134:
	;
	v455 = v444
	goto L135
L135:
	;
	v456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+24)) = uint8(v456)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+1448))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v458)+1448)) = v459 - v461
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	F_dsm_unpin_segment(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L137
	}
L136:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+12)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v455 = v454
	goto L135
L137:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	F_dsm_detach(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v475 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v471+v427<<(uint(int32(2))%32))+32)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+1468))
	*(*int32)(unsafe.Add(mBase, uint32(v477)+1468)) = v478 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v406)+8)) = v475
	*(*int64)(unsafe.Add(mBase, uint32(v406))) = int64(0)
	goto L124
L139:
	;
	goto L124
L140:
	;
	if v344 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_dsa_free(m, l0, v53)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
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
