package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CheckRecoveryConsistency(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int64
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	v10 = m.G0
	v12 = v10 - int32(1120)
	m.G0 = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[0]))
	if v15 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(1120)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[1]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[2]))
	if base.Ui64(v20) <= base.Ui64(v22-int64(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[3])))
	if v108 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v28 = *(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[4]))
	v31 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_errmsg_internal(m, int32(_a_F_CheckRecoveryConsistency_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[5]))
	v47 = F_LWLockAcquire(m, v43+int32(1152), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_1), int32(2226), int32(_a_F_CheckRecoveryConsistency_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[6]))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+136))
	if base.Ui64(v51) < base.Ui64(v20) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+144)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v50)+136)) = v20
	goto L15
L14:
	;
	goto L15
L15:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+168)) = uint8(v55)
	v57 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+160)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v50)+152)) = v57
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[7]))
	F_update_controlfile(m, v62, v50)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[5]))
	F_LWLockRelease(m, v66+int32(1152))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[2])) = v72
	*(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[4])) = v72
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[8])) = uint8(v78)
	v82 = F_errstart(m, int32(15), v78)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v82 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+76)) = uint32(v22)
	v87 = int64(32)
	v88 = int64(base.Ui64(v22) >> (uint(v87) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+72)) = uint32(v88)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v28)
	v92 = int64(base.Ui64(v28) >> (uint(v87) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v92)
	F_errmsg(m, int32(_a_F_CheckRecoveryConsistency_3), v12-int32(-64))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_1), int32(2240), int32(_a_F_CheckRecoveryConsistency_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[9]))
	if v439 != int32(3) {
		goto L1
	} else {
		goto L109
	}
L23:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[8])))
	if v110&int32(1) != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v114 = *(*int64)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[0]))
	if base.Ui64(v20) < base.Ui64(v114) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v116 = m.G0
	v118 = v116 - int32(112)
	m.G0 = v118
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[10]))
	if v121 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v123 = v118 + int32(20)
	F_hash_seq_init(m, v123, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	m.G0 = v118 + int32(112)
	v223 = F_AllocateDir(m, int32(_a_F_CheckRecoveryConsistency_4))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L58
	}
L29:
	;
	v126 = F_hash_seq_search(m, v123)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[10]))
	F_hash_destroy(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L57
	}
L31:
	;
	if v126 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v132 = v126
	goto L33
L33:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+20)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	v142 = v118 + int32(40)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	F_GetRelationPath(m, v142, v143, v144, v145, int32(-1), v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L35
	}
L34:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[11])))
	if v178 != 0 {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v152 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v152 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v142
	v159 = v139 & int32(1)
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
	v173 = F_hash_seq_search(m, v118+int32(20))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L48
	}
L40:
	;
	v160 = int32(_a_F_CheckRecoveryConsistency_5)
	goto L42
L41:
	;
	v160 = int32(_a_F_CheckRecoveryConsistency_6)
	goto L42
L42:
	;
	F_errmsg_internal(m, v160, v118)
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
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_7), v166, int32(_a_F_CheckRecoveryConsistency_8))
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
		v132 = v173
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
	F_errmsg_internal(m, int32(_a_F_CheckRecoveryConsistency_9), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_7), int32(258), int32(_a_F_CheckRecoveryConsistency_10))
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
	*(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[10])) = int32(0)
	goto L28
L58:
	;
	v226 = F_ReadDir(m, v223, int32(_a_F_CheckRecoveryConsistency_4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	if v226 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v228 = v226
	goto L63
L61:
	;
	goto L62
L62:
	;
	v393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[3])) = uint8(v393)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[12])))
	if v397 == v393 {
		goto L102
	} else {
		goto L103
	}
L63:
	;
	v238 = v228 + int32(19)
	v239 = int32(_a_F_CheckRecoveryConsistency_11)
	v243 = m.G0
	v245 = v243 - int32(32)
	v246 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v245)+24)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v245)+16)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v245)+8)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v245))) = v246
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[13])))
	if v254 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L62
L65:
	;
	v381 = F_ReadDir(m, v223, int32(_a_F_CheckRecoveryConsistency_4))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L99
	}
L66:
	;
	v323 = F_strlen(m, v238)
	mBase = m.M
	if v322 != v323 {
		goto L65
	} else {
		goto L85
	}
L67:
	;
	v322 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[14])))
	if v258 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v262 = v238
	goto L73
L71:
	;
	goto L72
L72:
	;
	v272 = v239
	v273 = v254
	goto L76
L73:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v268 == v254 {
		v262 = v262 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v322 = v262 - v238
	goto L66
L75:
	;
	goto L74
L76:
	;
	v280 = v245 + int32(base.Ui32(v273)>>(uint(int32(3))%32))&int32(28)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v282 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v281 | v282<<(uint(v273)%32)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	if v286 != 0 {
		v272 = v272 + v282
		v273 = v286
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v289 == int32(0) {
		v312 = v238
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v322 = v312 - v238
	goto L66
L80:
	;
	v293 = v238
	v294 = v289
	goto L81
L81:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(base.Ui32(v294)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v302)>>(uint(v294)%32))&int32(1) == int32(0) {
		v312 = v293
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v312 = v310
	goto L79
L83:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	v310 = v293 + int32(1)
	if v308 != 0 {
		v293 = v310
		v294 = v308
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(_a_F_CheckRecoveryConsistency_4)
	v329 = v12 + int32(80)
	v334 = F_pg_snprintf(m, v329, int32(1034), int32(_a_F_CheckRecoveryConsistency_12), v12+int32(48))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v338 = F_get_dirent_type(m, v329, v228, int32(0), int32(21))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	if v338 == int32(4) {
		goto L65
	} else {
		goto L88
	}
L88:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[15])))
	if v345 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v346 = int32(19)
	goto L91
L90:
	;
	v346 = int32(23)
	goto L91
L91:
	;
	v348 = F_errstart(m, v346, int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	if v348 == int32(0) {
		goto L65
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(_a_F_CheckRecoveryConsistency_4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v238
	F_errmsg(m, int32(_a_F_CheckRecoveryConsistency_13), v12+int32(32))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_CheckRecoveryConsistency_4)
	F_errdetail(m, int32(_a_F_CheckRecoveryConsistency_14), v12+int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_CheckRecoveryConsistency_15), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_1), int32(2186), int32(_a_F_CheckRecoveryConsistency_16))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	goto L65
L99:
	;
	if v381 != 0 {
		v228 = v381
		goto L63
	} else {
		goto L100
	}
L100:
	;
	goto L64
L101:
	;
	v413 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L105
	}
L102:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v401+int32(4)))) = int32(1)
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[17]))
	v410 = F_pgmem_kill(m, v408, int32(10))
	mBase = m.M
	goto L104
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	if v413 == int32(0) {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v20)
	v419 = int64(base.Ui64(v20) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v419)
	F_errmsg(m, int32(_a_F_CheckRecoveryConsistency_17), v12)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_CheckRecoveryConsistency_1), int32(2270), int32(_a_F_CheckRecoveryConsistency_2))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	goto L22
L109:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[18])))
	if v443&int32(1) != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[3])))
	if v447&int32(1) == int32(0) {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[12])))
	if v453&int32(1) == int32(0) {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[1]))
	v462 = base.AtomicRmwXchg32(m, v459, int32(96), int32(1))
	if v462 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[1]))
	F_s_lock(m, v464+int32(96), int32(_a_F_CheckRecoveryConsistency_1), int32(2283), int32(_a_F_CheckRecoveryConsistency_2))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L5
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[1]))
	v474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v474)
	v476 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v473)+96)), uint32(v476))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[18])) = uint8(v474)
	v484 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[12])))
	if v484 == v474 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	goto L1
L118:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v488+int32(8)))) = int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRecoveryConsistency[17]))
	v497 = F_pgmem_kill(m, v495, int32(10))
	mBase = m.M
	goto L120
L119:
	;
	goto L120
L120:
	;
	goto L117
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RecoveryInProgress[0])))
	if v3 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryInProgress[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+316))
		v11 = base.B2i32(v9 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_RecoveryInProgress[0])) = uint8(v11)
		v13 = v11
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_SetRecoveryPause(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SetRecoveryPause[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(96), int32(1))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_SetRecoveryPause[0]))
		F_s_lock(m, v9+int32(96), int32(_a_F_SetRecoveryPause_0), int32(3114), int32(_a_F_SetRecoveryPause_1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_SetRecoveryPause[0]))
			if l0 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
				if v19 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = int32(1)
				} else {
				}
				v24 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v24))
				return
			} else {
				v27 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v27
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v27))
				F_ConditionVariableBroadcast(m, v18+int32(84))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_SetRecoveryPause[0]))
		if l0 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
			if v19 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = int32(1)
			} else {
			}
			v24 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v24))
			return
		} else {
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v27
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v27))
			F_ConditionVariableBroadcast(m, v18+int32(84))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_prefetch[0])) = l0
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_prefetch[1]))
	if v6 == int32(13) {
		v9 = int32(_a_F_assign_recovery_prefetch_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_prefetch[2]))
		*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_prefetch[2])) = v11 + int32(1)
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_name[0]))
	switch v4 {
	case 0, 3:
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_name[0])) = int32(0)
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_name[0])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_name[1])) = l0
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_name[0])) = int32(3)
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
