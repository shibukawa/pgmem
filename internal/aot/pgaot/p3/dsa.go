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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
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
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
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
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
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
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
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
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
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
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	v12 = l1
	goto L1
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1468))
	if v21 != v23 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = F_LWLockAcquire(m, v185+v147<<(uint(int32(5))%32)+int32(224), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L60
	}
L3:
	;
	v28 = F_LWLockAcquire(m, v22+int32(1476), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v37 = int32(0)
	v40 = int32(base.Ui32(v12) >> (uint(int32(27)) % 32))
	v41 = F_get_segment_by_index(m, l0, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v32+int32(1476))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+int32(base.Ui32(v12)>>(uint(int32(10))%32))&int32(_a_F_dsa_free_0))))
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+1468))
	if v50 != v52 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v79 = v37
	goto L13
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v83 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v57 = F_LWLockAcquire(m, v51+int32(1476), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v69 = int32(base.Ui32(v49) >> (uint(int32(27)) % 32))
	v72 = l0 + v69*int32(20)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v61+int32(1476))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v77 = v73
	goto L22
L21:
	;
	v74 = F_get_segment_by_index(m, l0, v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L23
	}
L22:
	;
	v79 = v77 + v49&int32(134217727)
	goto L13
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v77 = v76
	goto L22
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+1468))
	if v84 != v86 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v113 = v37
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
	v91 = F_LWLockAcquire(m, v85+int32(1476), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v101 = int32(base.Ui32(v83) >> (uint(int32(27)) % 32))
	v104 = l0 + v101*int32(20)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if v105 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v95+int32(1476))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v109 = v105
	goto L35
L34:
	;
	v106 = F_get_segment_by_index(m, l0, v101)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v113 = v109 + v83&int32(134217727)
	goto L26
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v109 = v108
	goto L35
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+1468))
	if v116 != v118 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v146 = int32(0)
	goto L39
L39:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+20)))
	if v147 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v123 = F_LWLockAcquire(m, v117+int32(1476), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v134 = l0 + v40*int32(20)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	if v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v127+int32(1476))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v136 = F_get_segment_by_index(m, l0, v40)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	v146 = v139 + v12&int32(134217727)
	goto L39
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v139 = v138
	goto L48
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = F_LWLockAcquire(m, v150+int32(1476), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	F_FreePageManagerPut(m, v156, int32(base.Ui32(v157)>>(uint(int32(12))%32))&int32(_a_F_dsa_free_1), v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_rebin_segment(m, l0, v41)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v167+int32(1476))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = F_LWLockAcquire(m, v172+int32(256), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_unlink_span(m, l0, v79)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v180+int32(256))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v12 = v49
	goto L1
L60:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v146))) = uint16(v194)
	v197 = int32(1)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147<<(uint(v197)%32))+uint32(_c_F_dsa_free[0]))))
	v200 = base.I32_div_u_s(v146-v113, v199)
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+26)) = uint16(v200)
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+24)))
	v204 = v202 + v197
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+24)) = uint16(v204)
	if v202 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v492+v147<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L145
	}
L62:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+28)))
	if v288 != v204&int32(_a_F_dsa_free_2) {
		goto L61
	} else {
		goto L92
	}
L63:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+30)))
	if v206 != int32(3) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_unlink_span(m, l0, v79)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v211 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+1468))
	if v212 != v214 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v246 = int32(0)
	goto L68
L68:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	if v247 != 0 {
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v219 = F_LWLockAcquire(m, v213+int32(1476), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v231 = int32(base.Ui32(v211) >> (uint(int32(27)) % 32))
	v234 = l0 + v231*int32(20)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	if v235 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v223+int32(1476))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v239 = v235
	goto L77
L76:
	;
	v236 = F_get_segment_by_index(m, l0, v231)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L78
	}
L77:
	;
	v246 = v239 + v211&int32(134217727)
	goto L68
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v239 = v238
	goto L77
L79:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+1468))
	if v248 != v250 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v49
	v286 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+30)) = uint16(v286)
	goto L61
L82:
	;
	v255 = F_LWLockAcquire(m, v249+int32(1476), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v265 = int32(base.Ui32(v247) >> (uint(int32(27)) % 32))
	v268 = l0 + v265*int32(20)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	if v269 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v259+int32(1476))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v273 = v269
	goto L90
L89:
	;
	v270 = F_get_segment_by_index(m, l0, v265)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273+v247&int32(134217727))+4)) = v49
	goto L81
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v273 = v272
	goto L90
L92:
	;
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+30)))
	if v292 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v295 == int32(0) {
		goto L61
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v49 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+1468))
	if v299 != v301 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v329 = int32(0)
	goto L99
L99:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+20)))
	F_unlink_span(m, l0, v329)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L110
	}
L100:
	;
	v306 = F_LWLockAcquire(m, v300+int32(1476), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v316 = int32(base.Ui32(v49) >> (uint(int32(27)) % 32))
	v319 = l0 + v316*int32(20)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	if v320 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v310+int32(1476))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	v324 = v320
	goto L108
L107:
	;
	v321 = F_get_segment_by_index(m, l0, v316)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L109
	}
L108:
	;
	v329 = v324 + v49&int32(134217727)
	goto L99
L109:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v324 = v323
	goto L108
L110:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v338 = F_LWLockAcquire(m, v334+int32(1476), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+1468))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v341 != v342 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v348 = int32(0)
	goto L115
L113:
	;
	goto L114
L114:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v391 = F_get_segment_by_index(m, l0, int32(base.Ui32(v388)>>(uint(int32(27))%32)))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L122
	}
L115:
	;
	v359 = l0 + int32(8) + v348*int32(20)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	if v360 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v341
	goto L114
L117:
	;
	v374 = v348 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v374) <= base.Ui32(v375) {
		v348 = v374
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+24)))
	if v363 != int32(1) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	F_dsm_detach(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v359))) = int64(0)
	goto L117
L121:
	;
	goto L116
L122:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v329)+16))
	F_FreePageManagerPut(m, v393, int32(base.Ui32(v394)>>(uint(int32(12))%32))&int32(_a_F_dsa_free_1), v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if v403 != v405 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v475+int32(1476))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L140
	}
L125:
	;
	F_rebin_segment(m, l0, v391)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L139
	}
L126:
	;
	v408 = l0 + int32(8)
	if v391 == v408 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	if v410 != int32(-1) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+16))
	if v428 != int32(-1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v413 = F_get_segment_by_index(m, l0, v410)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v404)+20))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v404)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v419+v420<<(uint(int32(2))%32))+160)) = v424
	goto L128
L132:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v415)+16)) = v417
	goto L128
L133:
	;
	v431 = F_get_segment_by_index(m, l0, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L136
	}
L134:
	;
	v438 = v427
	goto L135
L135:
	;
	v439 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v438)+24)) = uint8(v439)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+1448))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v441)+1448)) = v442 - v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+12))
	F_dsm_unpin_segment(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L137
	}
L136:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+12)) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v438 = v437
	goto L135
L137:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_dsm_detach(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v456 = base.I32_div_s(v391-v408, int32(5))
	v458 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v454+v456)+32)) = v458
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+1468))
	*(*int32)(unsafe.Add(mBase, uint32(v460)+1468)) = v461 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = v458
	*(*int64)(unsafe.Add(mBase, uint32(v391))) = int64(0)
	goto L124
L139:
	;
	goto L124
L140:
	;
	if v331 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_dsa_free(m, l0, v49)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
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
