package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckRecoveryConsistency(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int64
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	v9 = m.G0
	v11 = v9 - int32(1120)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[171]))
	if v14 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(1120)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v21 = *(*int64)(unsafe.Add(mBase, _consts[174]))
	if base.Ui64(v19) <= base.Ui64(v21-int64(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[175])))
	if v107 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v27 = *(*int64)(unsafe.Add(mBase, _consts[176]))
	v30 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_errmsg_internal(m, int32(439727), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v46 = F_LWLockAcquire(m, v42+int32(1152), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	F_errfinish(m, int32(470996), int32(2226), int32(21597))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+136))
	if base.Ui64(v50) < base.Ui64(v19) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+144)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v49)+136)) = v19
	goto L15
L14:
	;
	goto L15
L15:
	;
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+152)) = v54
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+168)) = uint8(v56)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+160)) = v54
	v61 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	F_update_controlfile(m, v61, v49)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v65+int32(1152))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[174])) = v71
	*(*int64)(unsafe.Add(mBase, _consts[176])) = v71
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[177])) = uint8(v77)
	v81 = F_errstart(m, int32(15), v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v81 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+76)) = uint32(v21)
	v86 = int64(32)
	v87 = int64(base.Ui64(v21) >> (uint(v86) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+72)) = uint32(v87)
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+68)) = uint32(v27)
	v91 = int64(base.Ui64(v27) >> (uint(v86) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+64)) = uint32(v91)
	F_errmsg(m, int32(493545), v11-int32(-64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(470996), int32(2240), int32(21597))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	if v478 != int32(3) {
		goto L1
	} else {
		goto L125
	}
L23:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[177])))
	if v109 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v111 = *(*int64)(unsafe.Add(mBase, _consts[171]))
	if base.Ui64(v19) < base.Ui64(v111) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v113 = m.G0
	v115 = v113 - int32(112)
	m.G0 = v115
	v118 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	if v118 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_hash_seq_init(m, v115+int32(20), v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	m.G0 = v115 + int32(112)
	v221 = F_AllocateDir(m, int32(467736))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L58
	}
L29:
	;
	v125 = F_hash_seq_search(m, v115+int32(20))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	F_hash_destroy(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L57
	}
L31:
	;
	if v125 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v129 = v125
	goto L33
L33:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+20)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	F_GetRelationPath(m, v115+int32(40), v141, v142, v143, int32(-1), v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L35
	}
L34:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[179])))
	if v178 != 0 {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v150 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v150 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = v115 + int32(40)
	v159 = v137 & int32(1)
	if v159 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v173 = F_hash_seq_search(m, v115+int32(20))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L48
	}
L40:
	;
	v160 = int32(418913)
	goto L42
L41:
	;
	v160 = int32(66717)
	goto L42
L42:
	;
	F_errmsg_internal(m, v160, v115)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	if v159 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v166 = int32(93)
	goto L46
L45:
	;
	v166 = int32(96)
	goto L46
L46:
	;
	F_errfinish(m, int32(472461), v166, int32(389358))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	if v173 != 0 {
		v129 = v173
		goto L33
	} else {
		goto L49
	}
L49:
	;
	goto L34
L50:
	;
	v179 = int32(19)
	goto L52
L51:
	;
	v179 = int32(23)
	goto L52
L52:
	;
	v181 = F_errstart(m, v179, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	if v181 == int32(0) {
		goto L30
	} else {
		goto L54
	}
L54:
	;
	F_errmsg_internal(m, int32(161518), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(472461), int32(258), int32(161687))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	goto L30
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[178])) = int32(0)
	goto L28
L58:
	;
	v224 = F_ReadDir(m, v221, int32(467736))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	if v224 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v227 = v224
	goto L63
L61:
	;
	goto L62
L62:
	;
	v446 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[175])) = uint8(v446)
	F_SendPostmasterSignal(m, v446)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L120
	}
L63:
	;
	v235 = v227 + int32(19)
	v236 = int32(525078)
	v240 = m.G0
	v242 = v240 - int32(32)
	v243 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v242)+24)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242)+16)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242)+8)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = v243
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v251 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L62
L65:
	;
	v435 = F_ReadDir(m, v221, int32(467736))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L5
	} else {
		goto L118
	}
L66:
	;
	if v235&int32(3) == int32(0) {
		v343 = v235
		goto L89
	} else {
		goto L90
	}
L67:
	;
	v319 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v255 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v259 = v235
	goto L73
L71:
	;
	goto L72
L72:
	;
	v269 = v236
	v270 = v251
	goto L76
L73:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v265 == v251 {
		v259 = v259 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v319 = v259 - v235
	goto L66
L75:
	;
	goto L74
L76:
	;
	v277 = v242 + int32(base.Ui32(v270)>>(uint(int32(3))%32))&int32(28)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v278 | v279<<(uint(v270)%32)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v283 != 0 {
		v269 = v269 + v279
		v270 = v283
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v286 == int32(0) {
		v311 = v235
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v319 = v311 - v235
	goto L66
L80:
	;
	v290 = v235
	v291 = v286
	goto L81
L81:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v242+int32(base.Ui32(v291)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v299)>>(uint(v291)%32))&int32(1) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v311 = v307
	goto L79
L83:
	;
	v311 = v290
	goto L79
L84:
	;
	goto L85
L85:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	v307 = v290 + int32(1)
	if v305 != 0 {
		v290 = v307
		v291 = v305
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	if v319 != v376 {
		goto L65
	} else {
		goto L104
	}
L88:
	;
	v376 = v368 - v235
	goto L87
L89:
	;
	v347 = v343
	goto L98
L90:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v327 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v376 = int32(0)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v332 = v235
	goto L94
L94:
	;
	v336 = v332 + int32(1)
	if v336&int32(3) == int32(0) {
		v343 = v336
		goto L89
	} else {
		goto L96
	}
L95:
	;
	v368 = v336
	goto L88
L96:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	if v341 != 0 {
		v332 = v336
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v356 = int32(-2139062144)
	if (int32(16843008)-v353|v353)&v356 == v356 {
		v347 = v347 + int32(4)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v362 = v347
	goto L101
L100:
	;
	goto L99
L101:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v366 != 0 {
		v362 = v362 + int32(1)
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v368 = v362
	goto L88
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(467736)
	v387 = F_pg_snprintf(m, v11+int32(80), int32(1034), int32(167509), v11+int32(48))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v393 = F_get_dirent_type(m, v11+int32(80), v227, int32(0), int32(21))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	if v393 == int32(4) {
		goto L65
	} else {
		goto L107
	}
L107:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v400 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v401 = int32(19)
	goto L110
L109:
	;
	v401 = int32(23)
	goto L110
L110:
	;
	v403 = F_errstart(m, v401, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	if v403 == int32(0) {
		goto L65
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(467736)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v235
	F_errmsg(m, int32(175533), v11+int32(32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(467736)
	F_errdetail(m, int32(555926), v11+int32(16))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errhint(m, int32(591079), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(470996), int32(2186), int32(12607))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	goto L65
L118:
	;
	if v435 != 0 {
		v227 = v435
		goto L63
	} else {
		goto L119
	}
L119:
	;
	goto L64
L120:
	;
	v453 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	if v453 == int32(0) {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v19)
	v459 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v459)
	F_errmsg(m, int32(491519), v11)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(470996), int32(2270), int32(21597))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	goto L22
L125:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, _consts[183])))
	if v482 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, _consts[175])))
	if v484 != int32(1) {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v488 != int32(1) {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+96)) = int32(1)
	if v493 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	F_s_lock(m, v497+int32(96), int32(470996), int32(2283), int32(21597))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L5
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+96)) = int32(0)
	v509 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v506))) = uint8(v509)
	*(*uint8)(unsafe.Add(mBase, _consts[183])) = uint8(v509)
	F_SendPostmasterSignal(m, int32(2))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L5
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	goto L1
}
func F_RecoveryInProgress(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v3 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[113]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+316))
		v11 = base.B2i32(v9 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v11)
		v13 = v11
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_SetRecoveryPause(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[173]))
		F_s_lock(m, v10+int32(96), int32(470996), int32(3114), int32(343996))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[173]))
			if l0 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
				if v20 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(1)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = int32(0)
				return
			} else {
				v27 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v27
				F_ConditionVariableBroadcast(m, v19+int32(84))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[173]))
		if l0 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
			if v20 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(1)
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = int32(0)
			return
		} else {
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v27
			F_ConditionVariableBroadcast(m, v19+int32(84))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_assign_recovery_prefetch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, _consts[155])) = l0
	v6 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	if v6 == int32(13) {
		v9 = int32(4339728)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[156]))
		*(*int32)(unsafe.Add(mBase, _consts[156])) = v11 + int32(1)
	} else {
	}
	return
}
func F_assign_recovery_target_name(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	switch v4 {
	case 0, 3:
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[186])) = int32(0)
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[186])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[187])) = l0
				*(*int32)(unsafe.Add(mBase, _consts[186])) = int32(3)
				return
			}
		}
	default:
		F_error_multiple_recovery_targets(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
