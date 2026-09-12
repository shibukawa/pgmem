package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v330 int64
	_ = v330
	var v333 int64
	_ = v333
	var v336 int64
	_ = v336
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v619 int64
	_ = v619
	var v628 int32
	_ = v628
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v637 int64
	_ = v637
	var v640 int64
	_ = v640
	var v643 int64
	_ = v643
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v47 = int32(1)
	if v16&v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v21&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v30 = v19
	goto L7
L6:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L7
L7:
	;
	if v21 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = v19
	goto L10
L9:
	;
	v33 = v30
	goto L10
L10:
	;
	v46 = v33
	goto L1
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	v51 = v47
	goto L14
L13:
	;
	v51 = int32(4)
	goto L14
L14:
	;
	v52 = l0 + v51
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[831]))
	if v53 < v55 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v663 = F_pg_regerror(m, v395, v14+int32(16))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L96
	} else {
		goto L180
	}
L16:
	;
	m.G0 = v14 + int32(176)
	return int32(4461044)
L17:
	;
	v58 = v53
	goto L20
L18:
	;
	goto L19
L19:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[832]))
	if v357 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L20:
	;
	v70 = v58 * int32(52)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[833])))
	if v73 != v46 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v343 = v58 + int32(1)
	if v343 != v55 {
		v58 = v343
		goto L20
	} else {
		goto L92
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[834])))
	if v77 != l1 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[835])))
	if v79 != l2 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[836])))
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v143 != 0 {
		goto L22
	} else {
		goto L44
	}
L27:
	;
	v143 = int32(0)
	goto L26
L28:
	;
	v117 = v112
	v118 = v113
	v119 = v114
	goto L38
L29:
	;
	if (v81|v52)&int32(3) != 0 {
		v112 = v81
		v113 = v52
		v114 = v46
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v105 = v81
	v106 = v52
	v107 = v46
	goto L31
L31:
	;
	if v107 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v89 = v81
	v90 = v52
	v91 = v46
	goto L33
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != v95 {
		v112 = v89
		v113 = v90
		v114 = v91
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v105 = v100
	v106 = v98
	v107 = v102
	goto L31
L35:
	;
	v97 = int32(4)
	v98 = v90 + v97
	v100 = v89 + v97
	v102 = v91 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v89 = v100
		v90 = v98
		v91 = v102
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v112 = v105
	v113 = v106
	v114 = v107
	goto L28
L38:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 == v123 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v143 = v122 - v123
	goto L26
L40:
	;
	v125 = int32(1)
	v130 = v119 - v125
	if v130 != 0 {
		v117 = v117 + v125
		v118 = v118 + v125
		v119 = v130
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	if v58 == int32(0) {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v147 = v14 + int32(168)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[837])))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v148
	v151 = v14 + int32(160)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[838])))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v152
	v155 = v14 + int32(152)
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[839])))
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = v156
	v159 = v14 + int32(144)
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[840])))
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = v160
	v163 = v14 + int32(136)
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[835])))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
	v167 = v14 + int32(128)
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[833])))
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[841])))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+120)) = v170
	v172 = int32(4461076)
	v173 = int32(4461024)
	v175 = v58 * int32(52)
	goto L48
L46:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, _consts[837])) = v321
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*int64)(unsafe.Add(mBase, _consts[838])) = v324
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
	*(*int64)(unsafe.Add(mBase, _consts[839])) = v327
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, _consts[840])) = v330
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
	*(*int64)(unsafe.Add(mBase, _consts[835])) = v333
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	*(*int64)(unsafe.Add(mBase, _consts[833])) = v336
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	*(*int64)(unsafe.Add(mBase, _consts[841])) = v339
	goto L16
L47:
	;
	goto L46
L48:
	;
	v179 = v172 + v175
	if base.Ui32(v173-v179) <= base.Ui32(int32(0)-v175<<(uint(int32(1))%32)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = F___memcpy(m, v172, v173, v175)
	mBase = m.M
	goto L46
L50:
	;
	goto L51
L51:
	;
	goto L55
L55:
	;
	goto L56
L56:
	;
	goto L68
L67:
	;
	if v251 == int32(0) {
		goto L47
	} else {
		goto L80
	}
L68:
	;
	if v179&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v216 = v175
	goto L72
L70:
	;
	v231 = v175
	goto L71
L71:
	;
	if base.Ui32(v231) <= base.Ui32(int32(3)) {
		v251 = v231
		goto L67
	} else {
		goto L76
	}
L72:
	;
	if v216 == int32(0) {
		goto L47
	} else {
		goto L74
	}
L73:
	;
	v231 = v222
	goto L71
L74:
	;
	v222 = v216 - int32(1)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[841]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[842]))) = uint8(v225)
	if (v172+v222)&int32(3) != 0 {
		v216 = v222
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v238 = v231
	goto L77
L77:
	;
	v242 = v238 - int32(4)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[841])))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[842]))) = v245
	if base.Ui32(int32(3)) < base.Ui32(v242) {
		v238 = v242
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v251 = v242
	goto L67
L79:
	;
	goto L78
L80:
	;
	v258 = v251
	goto L81
L81:
	;
	v262 = v258 - int32(1)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+uint32(_consts[841]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+uint32(_consts[842]))) = uint8(v265)
	if v262 != 0 {
		v258 = v262
		goto L81
	} else {
		goto L83
	}
L82:
	;
	goto L47
L83:
	;
	goto L82
L92:
	;
	goto L21
L93:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v367 = F_AllocSetContextCreateInternal(m, v362, int32(61589), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v373 = v46 + int32(1)
	v376 = F_palloc(m, v373<<(uint(int32(2))%32))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L96
	} else {
		goto L98
	}
L96:
	;
	return int32(0)
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[832])) = v367
	goto L95
L98:
	;
	v378 = F_pg_mb2wchar_with_len(m, v52, v376, v46)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v386 = F_AllocSetContextCreateInternal(m, v381, int32(61552), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v388 = int32(4476144)
	v389 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v386
	v395 = F_pg_regcomp(m, v14+int32(140), v376, v378, l1, l2)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v376)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L96
	} else {
		goto L102
	}
L102:
	;
	if v395 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	v399 = F_palloc(m, v373)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L96
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v399
	if v46 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v404+v46))) = uint8(v406)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v14)+120))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+36)) = v409
	goto L109
L106:
	;
	v402 = F__emscripten_memcpy_bulkmem(m, v399, v52, v46)
	mBase = m.M
	goto L108
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v46
	v415 = *(*int32)(unsafe.Add(mBase, _consts[831]))
	if int32(32) <= v415 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v420 = v415 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[831])) = v420
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v420*int32(52))+uint32(_consts[841])))
	F_MemoryContextDelete(m, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L96
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v14)+120))
	v432 = *(*int32)(unsafe.Add(mBase, _consts[832]))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)+16))
	if v436 != v432 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L112
L114:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[831]))
	if int32(0) < v467 {
		goto L131
	} else {
		goto L132
	}
L115:
	;
	if v436 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	goto L114
L118:
	;
	if v432 != 0 {
		goto L125
	} else {
		goto L126
	}
L119:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v430)+28))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v430)+24))
	if v441 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v440 == int32(0) {
		goto L118
	} else {
		goto L124
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+28)) = v440
	goto L120
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+20)) = v440
	goto L120
L124:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v430)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v446
	goto L118
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+16)) = v432
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v430)+28)) = v453
	if v453 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v430)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+16)) = int32(0)
	goto L117
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = v430
	goto L130
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+20)) = v430
	goto L114
L131:
	;
	v470 = int32(4461076)
	v471 = int32(4461024)
	v473 = v467 * int32(52)
	goto L136
L132:
	;
	goto L133
L133:
	;
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	*(*int64)(unsafe.Add(mBase, _consts[841])) = v619
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v389
	*(*int32)(unsafe.Add(mBase, _consts[831])) = v467 + int32(1)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
	*(*int32)(unsafe.Add(mBase, _consts[837])) = v628
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v14)+160))
	*(*int64)(unsafe.Add(mBase, _consts[838])) = v631
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v14)+152))
	*(*int64)(unsafe.Add(mBase, _consts[839])) = v634
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v14)+144))
	*(*int64)(unsafe.Add(mBase, _consts[840])) = v637
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	*(*int64)(unsafe.Add(mBase, _consts[835])) = v640
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
	*(*int64)(unsafe.Add(mBase, _consts[833])) = v643
	goto L16
L134:
	;
	goto L133
L135:
	;
	goto L134
L136:
	;
	v477 = v470 + v473
	if base.Ui32(v471-v477) <= base.Ui32(int32(0)-v473<<(uint(int32(1))%32)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v484 = F___memcpy(m, v470, v471, v473)
	mBase = m.M
	goto L134
L138:
	;
	goto L139
L139:
	;
	goto L143
L143:
	;
	goto L144
L144:
	;
	goto L156
L155:
	;
	if v549 == int32(0) {
		goto L135
	} else {
		goto L168
	}
L156:
	;
	if v477&int32(3) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v514 = v473
	goto L160
L158:
	;
	v529 = v473
	goto L159
L159:
	;
	if base.Ui32(v529) <= base.Ui32(int32(3)) {
		v549 = v529
		goto L155
	} else {
		goto L164
	}
L160:
	;
	if v514 == int32(0) {
		goto L135
	} else {
		goto L162
	}
L161:
	;
	v529 = v520
	goto L159
L162:
	;
	v520 = v514 - int32(1)
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+uint32(_consts[841]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+uint32(_consts[842]))) = uint8(v523)
	if (v470+v520)&int32(3) != 0 {
		v514 = v520
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v536 = v529
	goto L165
L165:
	;
	v540 = v536 - int32(4)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v540)+uint32(_consts[841])))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+uint32(_consts[842]))) = v543
	if base.Ui32(int32(3)) < base.Ui32(v540) {
		v536 = v540
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v549 = v540
	goto L155
L167:
	;
	goto L166
L168:
	;
	v556 = v549
	goto L169
L169:
	;
	v560 = v556 - int32(1)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+uint32(_consts[841]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v560)+uint32(_consts[842]))) = uint8(v563)
	if v560 != 0 {
		v556 = v560
		goto L169
	} else {
		goto L171
	}
L170:
	;
	goto L135
L171:
	;
	goto L170
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L96
	} else {
		goto L181
	}
L181:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L96
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v14 + int32(16)
	F_errmsg(m, int32(200258), v14)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L96
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(490200), int32(223), int32(394795))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L96
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReScanExprContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = int32(4476144)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v8
	v11 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_MemoryContextReset(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L10
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	m.T0[v16].(func(*base.Module, int32))(m, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v6
	goto L3
L6:
	;
	return
L7:
	;
	F_pfree(m, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v21 != 0 {
		v11 = v21
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	return
}
