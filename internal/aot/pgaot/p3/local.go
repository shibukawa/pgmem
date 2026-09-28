package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LocalProcessControlFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v4 = F_palloc(m, int32(312))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_LocalProcessControlFile[0])) = v4
		*(*int32)(unsafe.Add(mBase, _c_F_LocalProcessControlFile[1])) = v4
		F_ReadControlFile(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_StartLocalBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v43 int32
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l0 + int32(36)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v13 != int32(-1) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v18
		if l2 != 0 {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v20
			v43 = int32(1)
			m.G0 = v9 + int32(16)
			return v43
		} else {
			if l1 == int32(0) {
				v43 = int32(1)
				m.G0 = v9 + int32(16)
				return v43
			} else {
				F_pgaio_wref_wait(m, v9)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = int64(0)
					v34 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v31, v31)
					if v34&int64(16777216) != v31 {
						v43 = int32(0)
					} else {
						v43 = int32(2)
					}
					m.G0 = v9 + int32(16)
					return v43
				}
			}
		}
	} else {
		v31 = int64(0)
		v34 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v31, v31)
		if v34&int64(16777216) != v31 {
			v43 = int32(0)
		} else {
			v43 = int32(2)
		}
		m.G0 = v9 + int32(16)
		return v43
	}
}
func F_local_buffer_write_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v14 = v7 + int32(8)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_write_error_callback[0]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_GetRelationPath(m, v14, v15, v16, v17, v19, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14
				F_errcontext_msg(m, int32(_a_F_local_buffer_write_error_callback_0), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v7 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v7 + int32(80)
		return
	}
}
func F_read_local_xlog_page(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_read_local_xlog_page_guts(m, l0, l1, l2, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_update_local_synced_slot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v280 int64
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v429 int64
	_ = v429
	var v431 int64
	_ = v431
	var v433 int64
	_ = v433
	var v436 int32
	_ = v436
	var v438 int64
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v26 = F_GetWalRcvFlushRecPtr(m, v3, v21+int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = F_GetXLogReplayRecPtr(m, v21+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	m.G0 = v21 + int32(16)
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v32) < base.Ui64(v26) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v15 + int32(208)
	return v476
L5:
	;
	v41 = v26
	goto L7
L6:
	;
	v41 = v32
	goto L7
L7:
	;
	if v34 == v35 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = v41
	goto L10
L9:
	;
	v43 = v32
	goto L10
L10:
	;
	if base.Ui64(v43) < base.Ui64(v39) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	F_pgstat_report_replslotsync(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	if base.Ui64(v91) < base.Ui64(v92) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+288))
	if v49 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = base.AtomicRmwXchg32(m, v46, int32(0), int32(1))
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v65 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	F_s_lock(m, v46, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+288)) = int32(1)
	v60 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v46))), uint32(v60))
	goto L17
L21:
	;
	goto L20
L22:
	;
	if v65 == int32(0) {
		v476 = v3
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+16)) = uint32(v43)
	v75 = int64(32)
	v76 = int64(base.Ui64(v43) >> (uint(v75) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v73
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v72)
	v81 = int64(base.Ui64(v72) >> (uint(v75) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v81)
	F_errmsg(m, int32(_a_F_update_local_synced_slot_1), v15)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(250), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v476 = v3
	goto L4
L27:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v345)+288))
	if v346 != v353 {
		goto L97
	} else {
		goto L98
	}
L28:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	v345 = v344
	v346 = int32(0)
	v347 = v340
	goto L27
L29:
	;
	v186 = m.G0
	v188 = v186 - int32(1152)
	m.G0 = v188
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = int32(_a_F_update_local_synced_slot_4)
	*(*uint32)(unsafe.Add(mBase, uint32(v188)+24)) = uint32(v91)
	v194 = int64(base.Ui64(v91) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v188)+20)) = uint32(v194)
	v197 = v188 + int32(128)
	v201 = F_pg_sprintf(m, v197, int32(_a_F_update_local_synced_slot_5), v188+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L59
	}
L30:
	;
	v179 = v18 + int32(120)
	if base.Ui64(v92) < base.Ui64(v91) {
		v185 = v179
		goto L29
	} else {
		goto L56
	}
L31:
	;
	v117 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	F_pgstat_report_replslotsync(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L40
	}
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v95 = int32(3)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	if base.B2i32(base.Ui32(v94) < base.Ui32(v95))|base.B2i32(base.Ui32(v97) < base.Ui32(v95)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v103 = v94 - v97
	if v103 < int32(0) {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(v97) <= base.Ui32(v94) {
		goto L30
	} else {
		goto L39
	}
L36:
	;
	v107 = v18 + int32(120)
	if base.Ui64(v92) < base.Ui64(v91) {
		v185 = v107
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	if base.B2i32(int32(0) < v103)|base.B2i32(base.Ui64(v111) < base.Ui64(v39)) != 0 {
		v185 = v107
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v340 = v3
	goto L28
L39:
	;
	goto L31
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+288))
	if v122 != int32(2) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = base.AtomicRmwXchg32(m, v119, int32(0), int32(1))
	if v127 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)+92))
	if v138 == int32(2) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	F_s_lock(m, v119, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+288)) = int32(2)
	v133 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v119))), uint32(v133))
	goto L43
L47:
	;
	goto L46
L48:
	;
	v141 = int32(15)
	goto L50
L49:
	;
	v141 = int32(14)
	goto L50
L50:
	;
	v143 = F_errstart(m, v141, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v143 == int32(0) {
		v476 = v117
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v147
	F_errmsg(m, int32(_a_F_update_local_synced_slot_6), v15-int32(-64))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v157
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+48)) = uint32(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v155
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v154)
	v162 = int64(32)
	v163 = int64(base.Ui64(v156) >> (uint(v162) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v163)
	v166 = int64(base.Ui64(v154) >> (uint(v162) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v166)
	v171 = F_errdetail(m, int32(_a_F_update_local_synced_slot_7), v15+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(295), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v476 = v117
	goto L4
L56:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	if base.Ui64(v181) < base.Ui64(v39) {
		v185 = v179
		goto L29
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v94) <= base.Ui32(v97) {
		v340 = v3
		goto L28
	} else {
		goto L58
	}
L58:
	;
	v185 = v179
	goto L29
L59:
	;
	v207 = F___fstatat(m, int32(-100), v197, v188+int32(32), int32(0))
	mBase = m.M
	goto L61
L60:
	;
	m.G0 = v188 + int32(1152)
	if v207 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	if v207 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[1]))
	if v211 == int32(44) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v197
	F_errmsg(m, int32(_a_F_update_local_synced_slot_8), v188)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_9), int32(2077), int32(_a_F_update_local_synced_slot_10))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v234 = int32(1)
	v237 = base.AtomicRmwXchg32(m, v18, int32(0), v234)
	if v237 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v256 = F_LogicalSlotAdvanceAndCheckSnapState(m, v253, v15+int32(144))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	F_s_lock(m, v18, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v241
	v243 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v245
	v247 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18))), uint32(v247))
	v340 = v234
	goto L28
L74:
	;
	goto L73
L75:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v259 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v258 == v259 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	if v262 != 0 {
		v292 = int32(0)
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L93
	}
L79:
	;
	v294 = int32(1)
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
	if v252 != v295 {
		v301 = v294
		goto L86
	} else {
		goto L87
	}
L80:
	;
	v263 = int32(3)
	v266 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v266 == int32(0) {
		v292 = v263
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v270
	F_errmsg(m, int32(_a_F_update_local_synced_slot_6), v15+int32(96))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+84)) = uint32(v277)
	v280 = int64(base.Ui64(v277) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+80)) = uint32(v280)
	v285 = F_errdetail(m, int32(_a_F_update_local_synced_slot_11), v15+int32(80))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(370), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v292 = v263
	goto L79
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	if v262 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	if v251 != v297 {
		v301 = v294
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	v301 = base.B2i32(v250 != v299)
	goto L86
L89:
	;
	v345 = v303
	v346 = int32(0)
	v347 = v301
	goto L27
L90:
	;
	goto L91
L91:
	;
	F_pgstat_report_replslotsync(m, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v345 = v303
	v346 = v292
	v347 = v301
	goto L27
L93:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v311
	F_errmsg_internal(m, int32(_a_F_update_local_synced_slot_12), v15+int32(128))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+124)) = uint32(v319)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+116)) = uint32(v318)
	v322 = int64(32)
	v323 = int64(base.Ui64(v319) >> (uint(v322) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+120)) = uint32(v323)
	v326 = int64(base.Ui64(v318) >> (uint(v322) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+112)) = uint32(v326)
	F_errdetail_internal(m, int32(_a_F_update_local_synced_slot_13), v15+int32(112))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(356), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v357 = base.AtomicRmwXchg32(m, v345, int32(0), int32(1))
	if v357 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if l1 != v365 {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	F_s_lock(m, v345, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+288)) = v346
	v362 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v345))), uint32(v362))
	goto L99
L103:
	;
	goto L102
L104:
	;
	v457 = int32(1)
	v460 = base.AtomicRmwXchg32(m, v18, int32(0), v457)
	if v460 != 0 {
		goto L130
	} else {
		goto L131
	}
L105:
	;
	v449 = int32(0)
	if v347 == v449 {
		v476 = v449
		goto L4
	} else {
		goto L127
	}
L106:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v409 = F_strncpy(m, v15+int32(144), v407, int32(64))
	mBase = m.M
	v410 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409)+63)) = uint8(v410)
	goto L119
L107:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v367 != v368 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+202)))
	if v370 != v371 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = v18 + int32(137)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	if base.B2i32(v378 == int32(0))|base.B2i32(v378 != v381) != 0 {
		v399 = v378
		v400 = v381
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v399-v400 != 0 {
		goto L106
	} else {
		goto L117
	}
L111:
	;
	goto L110
L112:
	;
	v384 = v373
	v385 = v375
	goto L113
L113:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+1)))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	if v389 == int32(0) {
		v399 = v389
		v400 = v388
		goto L111
	} else {
		goto L115
	}
L114:
	;
	v399 = v389
	v400 = v388
	goto L111
L115:
	;
	v392 = int32(1)
	if v389 == v388 {
		v384 = v384 + v392
		v385 = v385 + v392
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v402 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v18)+128))
	if v402 == v403 {
		goto L105
	} else {
		goto L118
	}
L118:
	;
	goto L106
L119:
	;
	v412 = int32(1)
	v415 = base.AtomicRmwXchg32(m, v18, int32(0), v412)
	if v415 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_s_lock(m, v18, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+193)) = v419
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v15)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+185)) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v15)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+177)) = v423
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+169)) = v425
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v15)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+161)) = v427
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v15)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+153)) = v429
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v15)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+145)) = v431
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v15)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+137)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = l1
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v436)
	v438 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v438
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+202)) = uint8(v440)
	v442 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18))), uint32(v442))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v347 != 0 {
		goto L104
	} else {
		goto L126
	}
L126:
	;
	v476 = v412
	goto L4
L127:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	goto L104
L130:
	;
	F_s_lock(m, v18, int32(_a_F_update_local_synced_slot_0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v464
	v466 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18))), uint32(v466))
	F_ReplicationSlotsComputeRequiredXmin(m, v466)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v476 = v457
	goto L4
}
