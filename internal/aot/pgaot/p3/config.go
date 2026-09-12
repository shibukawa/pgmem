package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetConfigOption(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v12 = F_find_option(m, l0, v3, l1, int32(21))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return v91
L2:
	;
	return int32(0)
L3:
	;
	if v12 == int32(0) {
		v91 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	switch v18 {
	case 0:
		goto L5
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v91 = v3
		goto L1
	}
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	if v45 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v29
	v31 = int32(4468816)
	v37 = F_pg_snprintf(m, v31, int32(256), int32(332991), v8+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
	v22 = int32(4468816)
	v26 = F_pg_snprintf(m, v22, int32(256), int32(480823), v8)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v91 = v22
	goto L1
L11:
	;
	v91 = v31
	goto L1
L12:
	;
	v42 = v40
	goto L14
L13:
	;
	v42 = int32(733277)
	goto L14
L14:
	;
	v91 = v42
	goto L1
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L22
	}
L16:
	;
	v48 = v45
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L18:
	;
	goto L15
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v56 == v44 {
		v91 = v53
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v59 = v48 + int32(12)
	if v59 != 0 {
		v48 = v59
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v44
	F_errmsg_internal(m, int32(179024), v8+int32(32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(492061), int32(3036), int32(339545))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v86 = int32(269260)
	goto L27
L26:
	;
	v86 = int32(333774)
	goto L27
L27:
	;
	v91 = v86
	goto L1
}
func F_ProcessConfigFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v13 = F_AllocSetContextCreateInternal(m, v8, int32(324469), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(4470752)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
		if v6 != 0 {
			v22 = int32(13)
		} else {
			v22 = int32(15)
		}
		v23 = F_ProcessConfigFileInternal(m, l0, int32(1), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
			F_MemoryContextDelete(m, v13)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_SelectConfigFiles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
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
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(240)
	m.G0 = v8
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v8 + int32(240)
	return v558
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = l1
	F_write_stderr(m, int32(731151), v8+int32(128))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L25
	} else {
		goto L159
	}
L3:
	;
	v208 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	v215 = F___fstatat(m, int32(-100), v210, v8+int32(144), v208)
	mBase = m.M
	goto L67
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = int32(333508)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v60
	v191 = F_pg_sprintf(m, v148, int32(174483), v8+int32(96))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L25
	} else {
		goto L64
	}
L5:
	;
	v173 = F_make_absolute_path(m, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L25
	} else {
		goto L62
	}
L6:
	;
	v59 = l0
	goto L8
L7:
	;
	v10 = int32(536679)
	v11 = int32(0)
	v16 = F___strchrnul(m, v10, int32(61))
	mBase = m.M
	if v10 == v16 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v60 = F_make_absolute_path(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v59 = v58
	goto L8
L10:
	;
	v58 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v19 = v16 - v10
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[1121]))))
	if v21 != 0 {
		v51 = v11
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v51
	goto L9
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
	if v23 == int32(0) {
		v51 = v11
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 == int32(0) {
		v51 = v11
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v30 = v23
	v31 = v26
	goto L17
L17:
	;
	v34 = F_strncmp(m, v10, v31, v19)
	mBase = m.M
	if v34 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v51 = v38 + int32(1)
	goto L13
L19:
	;
	goto L18
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v38 = v37 + v19
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 == int32(61) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v43 != 0 {
		v30 = v30 + int32(4)
		v31 = v43
		goto L17
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v51 = v11
	goto L13
L25:
	;
	return int32(0)
L26:
	;
	if v60 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = F___fstatat(m, int32(-100), v60, v8+int32(144), int32(0))
	mBase = m.M
	goto L30
L28:
	;
	goto L29
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	if v167 == int32(0) {
		goto L2
	} else {
		goto L61
	}
L30:
	;
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = l1
	F_write_stderr(m, int32(724828), v8+int32(112))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	if v85 != 0 {
		v170 = v85
		goto L5
	} else {
		goto L37
	}
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v77 != int32(44) {
		v558 = v3
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_write_stderr(m, int32(730253), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v558 = v3
	goto L1
L37:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	if v60&int32(3) == int32(0) {
		v111 = v60
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v148 = F_MemoryContextAllocExtended(m, v87, v144+int32(17), int32(2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L55
	}
L39:
	;
	v144 = v136 - v60
	goto L38
L40:
	;
	v115 = v111
	goto L49
L41:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v95 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v144 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	v100 = v60
	goto L45
L45:
	;
	v104 = v100 + int32(1)
	if v104&int32(3) == int32(0) {
		v111 = v104
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v136 = v104
	goto L39
L47:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v109 != 0 {
		v100 = v104
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v124 = int32(-2139062144)
	if (int32(16843008)-v121|v121)&v124 == v124 {
		v115 = v115 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v130 = v115
	goto L52
L51:
	;
	goto L50
L52:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v134 != 0 {
		v130 = v130 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v136 = v130
	goto L39
L54:
	;
	goto L53
L55:
	;
	if v148 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(492061), int32(647), int32(481401))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v170 = v167
	goto L5
L62:
	;
	v175 = int32(1)
	v176 = int32(10)
	v178 = int32(0)
	v182 = F_set_config_with_handle(m, int32(380772), int32(0), v173, v175, v176, v176, v178, v175, v178, v178)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	F_emscripten_builtin_free(m, v173)
	mBase = m.M
	goto L3
L64:
	;
	v194 = int32(0)
	v195 = int32(1)
	v196 = int32(10)
	v202 = F_set_config_with_handle(m, int32(380772), v194, v148, v195, v196, v196, v194, v195, v194, v194)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v148)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	goto L3
L67:
	;
	if v215 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	v218 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v218
	F_write_stderr(m, int32(724869), v8+int32(80))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L25
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L25
	} else {
		goto L72
	}
L71:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v558 = v208
	goto L1
L72:
	;
	v230 = int32(0)
	v233 = F_find_option(m, int32(13165), v230, v230, int32(23))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L25
	} else {
		goto L73
	}
L73:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+92))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v236|v60 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l1
	v242 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v242
	F_write_stderr(m, int32(730611), v8-int32(-64))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L25
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v236 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v558 = v208
	goto L1
L78:
	;
	v249 = v236
	goto L80
L79:
	;
	v249 = v60
	goto L80
L80:
	;
	v250 = F_make_absolute_path(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v252 = int32(4465420)
	v253 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	F_emscripten_builtin_free(m, v253)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[244])) = v250
	v258 = int32(0)
	v261 = int32(1)
	v262 = int32(10)
	v268 = F_set_config_with_handle(m, int32(13165), v258, v250, v261, v262, v262, v258, v261, v258, v258)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L25
	} else {
		goto L82
	}
L82:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	v274 = int32(0)
	v276 = int32(1)
	v283 = F_set_config_with_handle(m, int32(141263), v274, int32(97347), v276, v276, int32(10), v274, v276, v274, v274)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	if v286 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v558 = int32(1)
	goto L1
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(333473)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v60
	v534 = F_pg_sprintf(m, v502, int32(174483), v8)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L25
	} else {
		goto L156
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
	v523 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v523
	F_write_stderr(m, int32(730791), v8+int32(48))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L25
	} else {
		goto L155
	}
L88:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	if v60&int32(3) == int32(0) {
		v465 = v60
		goto L134
	} else {
		goto L135
	}
L89:
	;
	if v60 == int32(0) {
		goto L87
	} else {
		goto L131
	}
L90:
	;
	v425 = F_make_absolute_path(m, v422)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L25
	} else {
		goto L129
	}
L91:
	;
	if v60 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	v406 = F_make_absolute_path(m, v286)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L25
	} else {
		goto L126
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(333524)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v60
	v385 = F_pg_sprintf(m, v351, int32(174483), v8+int32(16))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L25
	} else {
		goto L122
	}
L95:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	if v60&int32(3) == int32(0) {
		v314 = v60
		goto L100
	} else {
		goto L101
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	v372 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v372
	F_write_stderr(m, int32(730973), v8+int32(32))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L121
	}
L98:
	;
	v351 = F_MemoryContextAllocExtended(m, v290, v347+int32(13), int32(2))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L25
	} else {
		goto L115
	}
L99:
	;
	v347 = v339 - v60
	goto L98
L100:
	;
	v318 = v314
	goto L109
L101:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v298 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v347 = int32(0)
	goto L98
L103:
	;
	goto L104
L104:
	;
	v303 = v60
	goto L105
L105:
	;
	v307 = v303 + int32(1)
	if v307&int32(3) == int32(0) {
		v314 = v307
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v339 = v307
	goto L99
L107:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v312 != 0 {
		v303 = v307
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v327 = int32(-2139062144)
	if (int32(16843008)-v324|v324)&v327 == v327 {
		v318 = v318 + int32(4)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v333 = v318
	goto L112
L111:
	;
	goto L110
L112:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v337 != 0 {
		v333 = v333 + int32(1)
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v339 = v333
	goto L99
L114:
	;
	goto L113
L115:
	;
	if v351 != 0 {
		goto L94
	} else {
		goto L116
	}
L116:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L25
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L25
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(492061), int32(647), int32(481401))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L25
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v558 = int32(0)
	goto L1
L122:
	;
	v388 = int32(0)
	v389 = int32(1)
	v390 = int32(10)
	v396 = F_set_config_with_handle(m, int32(380943), v388, v351, v389, v390, v390, v388, v389, v388, v388)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L25
	} else {
		goto L123
	}
L123:
	;
	F_pfree(m, v351)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L25
	} else {
		goto L124
	}
L124:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[1124]))
	if v401 == int32(0) {
		goto L88
	} else {
		goto L125
	}
L125:
	;
	v422 = v401
	goto L90
L126:
	;
	v408 = int32(1)
	v409 = int32(10)
	v411 = int32(0)
	v415 = F_set_config_with_handle(m, int32(380943), int32(0), v406, v408, v409, v409, v411, v408, v411, v411)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L25
	} else {
		goto L127
	}
L127:
	;
	F_emscripten_builtin_free(m, v406)
	mBase = m.M
	v419 = *(*int32)(unsafe.Add(mBase, _consts[1124]))
	if v419 == int32(0) {
		goto L89
	} else {
		goto L128
	}
L128:
	;
	v422 = v419
	goto L90
L129:
	;
	v427 = int32(1)
	v428 = int32(10)
	v430 = int32(0)
	v434 = F_set_config_with_handle(m, int32(380552), int32(0), v425, v427, v428, v428, v430, v427, v430, v430)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L25
	} else {
		goto L130
	}
L130:
	;
	F_emscripten_builtin_free(m, v425)
	mBase = m.M
	goto L85
L131:
	;
	goto L88
L132:
	;
	v502 = F_MemoryContextAllocExtended(m, v441, v498+int32(15), int32(2))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L25
	} else {
		goto L149
	}
L133:
	;
	v498 = v490 - v60
	goto L132
L134:
	;
	v469 = v465
	goto L143
L135:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v449 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v498 = int32(0)
	goto L132
L137:
	;
	goto L138
L138:
	;
	v454 = v60
	goto L139
L139:
	;
	v458 = v454 + int32(1)
	if v458&int32(3) == int32(0) {
		v465 = v458
		goto L134
	} else {
		goto L141
	}
L140:
	;
	v490 = v458
	goto L133
L141:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v463 != 0 {
		v454 = v458
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v478 = int32(-2139062144)
	if (int32(16843008)-v475|v475)&v478 == v478 {
		v469 = v469 + int32(4)
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v484 = v469
	goto L146
L145:
	;
	goto L144
L146:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v488 != 0 {
		v484 = v484 + int32(1)
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v490 = v484
	goto L133
L148:
	;
	goto L147
L149:
	;
	if v502 != 0 {
		goto L86
	} else {
		goto L150
	}
L150:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L25
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L25
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L25
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(492061), int32(647), int32(481401))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L25
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v558 = int32(0)
	goto L1
L156:
	;
	v537 = int32(0)
	v538 = int32(1)
	v539 = int32(10)
	v545 = F_set_config_with_handle(m, int32(380552), v537, v502, v538, v539, v539, v537, v538, v537, v537)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L25
	} else {
		goto L157
	}
L157:
	;
	F_pfree(m, v502)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L25
	} else {
		goto L158
	}
L158:
	;
	goto L85
L159:
	;
	v558 = v3
	goto L1
}
func F_set_config_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
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
	if base.B2i32(l3 != int32(9))&base.B2i32(base.Ui32(l3) <= base.Ui32(int32(10))) != 0 {
		v16 = int32(10)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v16 = v15
	}
	v17 = int32(0)
	v19 = F_set_config_with_handle(m, l0, int32(0), l1, l2, l3, v16, l4, l5, v17, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		return
	}
}
