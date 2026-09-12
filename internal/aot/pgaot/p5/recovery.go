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
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int64
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	v9 = m.G0
	v11 = v9 - int32(1120)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[172]))
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v21 = *(*int64)(unsafe.Add(mBase, _consts[175]))
	if base.Ui64(v19) <= base.Ui64(v21-int64(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[176])))
	if v107 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v27 = *(*int64)(unsafe.Add(mBase, _consts[177]))
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
	F_errmsg_internal(m, int32(476770), int32(0))
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	F_errfinish(m, int32(511127), int32(2226), int32(23379))
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
	v49 = *(*int32)(unsafe.Add(mBase, _consts[121]))
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
	v61 = *(*int32)(unsafe.Add(mBase, _consts[136]))
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
	v65 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	*(*int64)(unsafe.Add(mBase, _consts[175])) = v71
	*(*int64)(unsafe.Add(mBase, _consts[177])) = v71
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[178])) = uint8(v77)
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
	F_errmsg(m, int32(536614), v11-int32(-64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(511127), int32(2240), int32(23379))
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
	v422 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if v422 != int32(3) {
		goto L1
	} else {
		goto L108
	}
L23:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[178])))
	if v109 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v111 = *(*int64)(unsafe.Add(mBase, _consts[172]))
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
	v118 = *(*int32)(unsafe.Add(mBase, _consts[179]))
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
	v221 = F_AllocateDir(m, int32(507063))
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
	v203 = *(*int32)(unsafe.Add(mBase, _consts[179]))
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
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
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
	v160 = int32(455275)
	goto L42
L41:
	;
	v160 = int32(75095)
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
	F_errfinish(m, int32(513291), v166, int32(423011))
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
	F_errmsg_internal(m, int32(179576), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(513291), int32(258), int32(179778))
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
	*(*int32)(unsafe.Add(mBase, _consts[179])) = int32(0)
	goto L28
L58:
	;
	v224 = F_ReadDir(m, v221, int32(507063))
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
	v390 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[176])) = uint8(v390)
	F_SendPostmasterSignal(m, v390)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L103
	}
L63:
	;
	v235 = v227 + int32(19)
	v236 = int32(573285)
	v240 = m.G0
	v242 = v240 - int32(32)
	v243 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v242)+24)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242)+16)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242)+8)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = v243
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
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
	v379 = F_ReadDir(m, v221, int32(507063))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L101
	}
L66:
	;
	v320 = F_strlen(m, v235)
	mBase = m.M
	if v319 != v320 {
		goto L65
	} else {
		goto L87
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
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(507063)
	v331 = F_pg_snprintf(m, v11+int32(80), int32(1034), int32(185927), v11+int32(48))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	v337 = F_get_dirent_type(m, v11+int32(80), v227, int32(0), int32(21))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	if v337 == int32(4) {
		goto L65
	} else {
		goto L90
	}
L90:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _consts[183])))
	if v344 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v345 = int32(19)
	goto L93
L92:
	;
	v345 = int32(23)
	goto L93
L93:
	;
	v347 = F_errstart(m, v345, int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	if v347 == int32(0) {
		goto L65
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(507063)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v235
	F_errmsg(m, int32(194025), v11+int32(32))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(507063)
	F_errdetail(m, int32(616997), v11+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	F_errhint(m, int32(653361), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(511127), int32(2186), int32(13663))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	goto L65
L101:
	;
	if v379 != 0 {
		v227 = v379
		goto L63
	} else {
		goto L102
	}
L102:
	;
	goto L64
L103:
	;
	v397 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	if v397 == int32(0) {
		goto L22
	} else {
		goto L105
	}
L105:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v19)
	v403 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v403)
	F_errmsg(m, int32(534588), v11)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(511127), int32(2270), int32(23379))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	goto L22
L108:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v426 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, _consts[176])))
	if v428 != int32(1) {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, _consts[185])))
	if v432 != int32(1) {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v436)+96)) = int32(1)
	if v437 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	F_s_lock(m, v441+int32(96), int32(511127), int32(2283), int32(23379))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+96)) = int32(0)
	v453 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v453)
	*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v453)
	F_SendPostmasterSignal(m, int32(2))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v3 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[114]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+316))
		v11 = base.B2i32(v9 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[113])) = uint8(v11)
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[174]))
		F_s_lock(m, v10+int32(96), int32(511127), int32(3114), int32(374767))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[174]))
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
		v19 = *(*int32)(unsafe.Add(mBase, _consts[174]))
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
	*(*int32)(unsafe.Add(mBase, _consts[156])) = l0
	v6 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if v6 == int32(13) {
		v9 = int32(4444640)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[157]))
		*(*int32)(unsafe.Add(mBase, _consts[157])) = v11 + int32(1)
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[187]))
	switch v4 {
	case 0, 3:
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[187])) = int32(0)
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[187])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[188])) = l0
				*(*int32)(unsafe.Add(mBase, _consts[187])) = int32(3)
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
