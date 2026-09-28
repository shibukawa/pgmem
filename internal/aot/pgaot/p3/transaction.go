package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetNewTransactionId(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v279 int64
	_ = v279
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v371 int64
	_ = v371
	var v375 int64
	_ = v375
	var v378 int64
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v442 int64
	_ = v442
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v13 + int32(80)
	return v442
L2:
	;
	if v273&int32(_a_F_GetNewTransactionId_4) != 0 {
		goto L83
	} else {
		goto L84
	}
L3:
	;
	v199 = int32(3)
	if base.B2i32(base.Ui32(v70) < base.Ui32(v199))|base.B2i32(base.Ui32(v85) < base.Ui32(v199)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v136
	F_errmsg(m, int32(_a_F_GetNewTransactionId_12), v13+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L21
	} else {
		goto L59
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L21
	} else {
		goto L56
	}
L6:
	;
	if v21&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v21 = int32(1)
	goto L9
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+76)))
	v21 = v20
	goto L9
L9:
	;
	goto L6
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[1]))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L21
	} else {
		goto L53
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[2]))
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[3]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32)))) = v32
	v442 = int64(1)
	goto L1
L14:
	;
	goto L15
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[4])))
	if v46 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v56 != 0 {
		goto L5
	} else {
		goto L20
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[5]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+308))
	v54 = base.B2i32(v52 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[4])) = uint8(v54)
	v56 = v54
	goto L19
L18:
	;
	v56 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[6]))
	v62 = F_LWLockAcquire(m, v58+int32(384), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int64(0)
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[7]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	v70 = base.I32_wrap_i64(v69)
	v71 = int32(3)
	if base.B2i32(base.Ui32(v70) < base.Ui32(v71))|base.B2i32(base.Ui32(v68) < base.Ui32(v71)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[6]))
	F_LWLockRelease(m, v87+int32(384))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L21
	} else {
		goto L29
	}
L24:
	;
	if int32(0) <= v70-v68 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(v70) < base.Ui32(v68) {
		v273 = v70
		v279 = v69
		goto L2
	} else {
		goto L28
	}
L27:
	;
	v273 = v70
	v279 = v69
	goto L2
L28:
	;
	goto L23
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[8])))
	v96 = int32(0)
	if base.B2i32(v93&int32(1) == v96)|v70&int32(_a_F_GetNewTransactionId_3) == v96 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v125 = int32(3)
	if base.B2i32(base.Ui32(v70) < base.Ui32(v125))|base.B2i32(base.Ui32(v84) < base.Ui32(v125)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[8])))
	if v105 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	if v93&int32(1) == int32(0) {
		goto L3
	} else {
		goto L39
	}
L34:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[8])))
	if v120 != 0 {
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v109+int32(16)))) = int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[10]))
	v118 = F_pgmem_kill(m, v116, int32(10))
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L3
L39:
	;
	goto L30
L40:
	;
	v136 = F_get_database_name(m, v82)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L46
	}
L41:
	;
	if v70-v84 < int32(0) {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v70) < base.Ui32(v84) {
		goto L3
	} else {
		goto L45
	}
L44:
	;
	goto L40
L45:
	;
	goto L40
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	if v136 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v82
	F_errmsg(m, int32(_a_F_GetNewTransactionId_13), v13)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_GetNewTransactionId_14), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_GetNewTransactionId_1), int32(157), int32(_a_F_GetNewTransactionId_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_GetNewTransactionId_15), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_GetNewTransactionId_1), int32(78), int32(_a_F_GetNewTransactionId_2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_GetNewTransactionId_0), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_GetNewTransactionId_1), int32(94), int32(_a_F_GetNewTransactionId_2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errhint(m, int32(_a_F_GetNewTransactionId_14), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_GetNewTransactionId_1), int32(150), int32(_a_F_GetNewTransactionId_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[6]))
	v267 = F_LWLockAcquire(m, v263+int32(384), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L21
	} else {
		goto L82
	}
L63:
	;
	v210 = F_get_database_name(m, v82)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L21
	} else {
		goto L69
	}
L64:
	;
	if int32(0) <= v70-v85 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(v70) < base.Ui32(v85) {
		goto L62
	} else {
		goto L68
	}
L67:
	;
	goto L62
L68:
	;
	goto L63
L69:
	;
	v214 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	if v210 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = base.F64_mul(base.F64_div(base.F64_convert_i32_u(v238), float64(2.147483647e+09)), float64(100))
	v249 = F_errdetail(m, int32(_a_F_GetNewTransactionId_9), v13+int32(32))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L79
	}
L72:
	;
	if v214 == int32(0) {
		goto L62
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v214 == int32(0) {
		goto L62
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v210
	v219 = v83 - v70
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v219
	F_errmsg(m, int32(_a_F_GetNewTransactionId_8), v13-int32(-64))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	v238 = v219
	v239 = int32(172)
	goto L71
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v82
	v230 = v83 - v70
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v230
	F_errmsg(m, int32(_a_F_GetNewTransactionId_11), v13+int32(48))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L21
	} else {
		goto L78
	}
L78:
	;
	v238 = v230
	v239 = int32(181)
	goto L71
L79:
	;
	F_errhint(m, int32(_a_F_GetNewTransactionId_10), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_GetNewTransactionId_1), v239, int32(_a_F_GetNewTransactionId_2))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	goto L62
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[7]))
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v270)+8))
	v273 = base.I32_wrap_i64(v271)
	v279 = v271
	goto L2
L83:
	;
	v285 = base.B2i32(v273 != int32(3))
	goto L85
L84:
	;
	v285 = int32(0)
	goto L85
L85:
	;
	if v285 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[11]))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	v292 = int32(base.Ui32(v273) >> (uint(int32(15)) % 32))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetNewTransactionId[12])))
	v295 = base.I32_rem_u_s(v292, v294)
	v298 = v290 + v295<<(uint(int32(7))%32)
	v300 = F_LWLockAcquire(m, v298, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L21
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[13]))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+24)))
	if v317 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v303 = base.I64_extend_i32_u(v292)
	v304 = F_SimpleLruZeroPage(m, int32(_a_F_GetNewTransactionId_5), v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L21
	} else {
		goto L90
	}
L90:
	;
	F_XLogSimpleInsertInt64(m, int32(3), int32(0), v303)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	F_LWLockRelease(m, v298)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L21
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	F_ExtendSUBTRANS(m, v273)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L21
	} else {
		goto L106
	}
L94:
	;
	v323 = int32(819)
	v324 = base.I32_div_u_s(v273, v323)
	if v273-v324*v323 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v328 = base.B2i32(v273 != int32(3))
	goto L97
L96:
	;
	v328 = int32(0)
	goto L97
L97:
	;
	if v328 != 0 {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[14]))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+28))
	v333 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetNewTransactionId[15])))
	v334 = base.I32_rem_u_s(v324, v333)
	v337 = v331 + v334<<(uint(int32(7))%32)
	v339 = F_LWLockAcquire(m, v337, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L21
	} else {
		goto L99
	}
L99:
	;
	v342 = base.I64_extend_i32_u(v324)
	v343 = F_SimpleLruZeroPage(m, int32(_a_F_GetNewTransactionId_7), v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L21
	} else {
		goto L100
	}
L100:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewTransactionId[16])))
	if v346 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_XLogSimpleInsertInt64(m, int32(18), int32(0), v342)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L21
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	F_LWLockRelease(m, v337)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L21
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	goto L93
L106:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[7]))
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v361)+8))
	v363 = int64(1)
	v371 = v362 + v363
	if base.Ui32(base.I32_wrap_i64(v371)) < base.Ui32(int32(3)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v375 = v362 + (v363-v362)&int64(4294967295) + int64(2)
	goto L109
L108:
	;
	v375 = v371
	goto L109
L109:
	;
	if base.Ui64(int64(2)) < base.Ui64(v371) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v378 = v375
	goto L112
L111:
	;
	v378 = v371
	goto L112
L112:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v361)+8)) = v378
	if l0 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[6]))
	F_LWLockRelease(m, v429+int32(384))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L21
	} else {
		goto L120
	}
L114:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+48)) = v273
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[3]))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v383)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v387+v388<<(uint(int32(2))%32)))) = v273
	goto L113
L115:
	;
	goto L116
L116:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[3]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[2]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+32))
	v401 = v395 + v398<<(uint(int32(1))%32)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+56)))
	if base.Ui32(v402) <= base.Ui32(int32(63)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397+v402<<(uint(int32(2))%32))+60)) = v273
	v409 = int32(0)
	v412 = base.AtomicRmwOr32(m, v409, int32(_a_F_GetNewTransactionId_6), v409)
	v414 = v402 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v414)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+56)) = uint8(v414)
	goto L113
L118:
	;
	goto L119
L119:
	;
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+1)) = uint8(v419)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewTransactionId[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(v422)+57)) = uint8(v419)
	goto L113
L120:
	;
	v442 = v279
	goto L1
}
func F_TransactionIdGetCommitLSN(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitLSN[0]))
	if v9 == l0 {
		v12 = *(*int64)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitLSN[1]))
		v23 = v12
		m.G0 = v6 + int32(16)
		return v23
	} else {
		if base.Ui32(l0) < base.Ui32(int32(3)) {
			v23 = int64(0)
			m.G0 = v6 + int32(16)
			return v23
		} else {
			v18 = F_TransactionIdGetStatus(m, l0, v6+int32(8))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
				v23 = v22
				m.G0 = v6 + int32(16)
				return v23
			}
		}
	}
}
func F_TransactionIdGetStatus(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l0
	v15 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v19 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_TransactionIdGetStatus_0), base.I64_extend_i32_u(v15), v10+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(_a_F_TransactionIdGetStatus_1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[0]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		v26 = int32(2)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v19<<(uint(v26)%32))))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(l0)>>(uint(v26)%32))&int32(_a_F_TransactionIdGetStatus_2)))))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v46 = *(*int64)(unsafe.Add(mBase, uint32(v36+v19<<(uint(int32(13))%32)+int32(base.Ui32(v40)>>(uint(v26)%32))&int32(_a_F_TransactionIdGetStatus_3))))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v46
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[0]))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
		v52 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[1])))
		v53 = base.I32_rem_u_s(v15, v52)
		F_LWLockRelease(m, v50+v53<<(uint(int32(7))%32))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return int32(base.Ui32(v35)>>(uint(l0<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
		}
	}
}
func F_TransactionIdIsCurrentTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[0]))
	if v14 == l0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[1]))
	if v19 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v125
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[2]))
	if v23 == int32(0) {
		v125 = int32(0)
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[3]))
	v95 = int32(0)
	v98 = v19 - int32(1)
	goto L30
L11:
	;
	v28 = v23
	goto L12
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v34 == int32(4) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v125 = int32(0)
	goto L7
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	if v88 != 0 {
		v28 = v88
		goto L12
	} else {
		goto L29
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v37 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v40 = int32(1)
	if l0 == v37 {
		v125 = v40
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	v44 = v42 - int32(1)
	if v44 < int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v50 = int32(0)
	v53 = v44
	goto L19
L19:
	;
	v58 = int32(2)
	v59 = base.I32_div_s(v53-v50, v58)
	v60 = v59 + v50
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v47+v60<<(uint(v58)%32))))
	if v64 == l0 {
		v125 = v40
		goto L7
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v73 = base.B2i32(v64-l0 < int32(0)) | base.B2i32(base.Ui32(v64) < base.Ui32(int32(3)))
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v74 = v60 + int32(1)
	goto L24
L23:
	;
	v74 = v50
	goto L24
L24:
	;
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v77 = v53
	goto L27
L26:
	;
	v77 = v60 - int32(1)
	goto L27
L27:
	;
	if v74 <= v77 {
		v50 = v74
		v53 = v77
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	goto L13
L30:
	;
	v103 = int32(2)
	v104 = base.I32_div_s(v98-v95, v103)
	v105 = v104 + v95
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v93+v105<<(uint(v103)%32))))
	v110 = base.B2i32(v109 == l0)
	if v109 == l0 {
		v125 = v110
		goto L7
	} else {
		goto L32
	}
L31:
	;
	v125 = v110
	goto L7
L32:
	;
	v113 = base.B2i32(base.Ui32(v109) < base.Ui32(l0))
	if base.Ui32(v109) < base.Ui32(l0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v105 + int32(1)
	goto L35
L34:
	;
	v114 = v95
	goto L35
L35:
	;
	if base.Ui32(v109) < base.Ui32(l0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = v98
	goto L38
L37:
	;
	v117 = v105 - int32(1)
	goto L38
L38:
	;
	if v114 <= v117 {
		v95 = v114
		v98 = v117
		goto L30
	} else {
		goto L39
	}
L39:
	;
	goto L31
}
func F_TransactionIdSetPageStatusInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v257 int32
	_ = v257
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l0
	v22 = base.B2i32(l4 == int64(0))
	v25 = F_SimpleLruReadPage(m, int32(_a_F_TransactionIdSetPageStatusInternal_0), l5, v22, v17+int32(12))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		if v27 == int32(0) {
		} else {
			v32 = int32(0)
			if base.B2i32(l3 != int32(1))|base.B2i32(l1 <= v32) == v32 {
				v40 = int32(0)
				for {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
					v56 = int32(1)
					v58 = int32(2)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l2+v40<<(uint(v58)%32))))
					v65 = int32(base.Ui32(v61&int32(_a_F_TransactionIdSetPageStatusInternal_1)) >> (uint(v58) % 32))
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v25<<(uint(v58)%32))))
					v73 = v65 + v72
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
					v78 = v61 << (uint(v56) % 32) & int32(6)
					if base.B2i32(v55 == v56)&base.B2i32(int32(base.Ui32(v74)>>(uint(v78)%32))&int32(3) == v56) != 0 {
					} else {
						v87 = v74 | int32(3)<<(uint(v78)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v87)
						if l4 == int64(0) {
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
							v95 = v91 + v25<<(uint(int32(13))%32) + v65&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
							if base.Ui64(l4) <= base.Ui64(v96) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v95))) = l4
							}
						}
					}
					v101 = v40 + int32(1)
					if v101 != l1 {
						v40 = v101
						continue
					} else {
						break
					}
					break
				}
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				v104 = v103
			} else {
				v104 = v27
			}
			v119 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
			v121 = int32(2)
			v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v25<<(uint(v121)%32))))
			v126 = v104 & int32(_a_F_TransactionIdSetPageStatusInternal_1)
			v129 = v124 + int32(base.Ui32(v126)>>(uint(v121)%32))
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
			v134 = v104 << (uint(int32(1)) % 32) & int32(6)
			if l3 != int32(3) {
				v154 = v130&(int32(3)<<(uint(v134)%32)^int32(-1)) | l3<<(uint(v134)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v154)
				if l4 == int64(0) {
				} else {
					v157 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+36))
					v166 = v158 + v25<<(uint(int32(13))%32) + int32(base.Ui32(v126)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
					v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)))
					if base.Ui64(l4) <= base.Ui64(v167) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v166))) = l4
					}
				}
			} else {
				v138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
				if v138&int32(1) == int32(0) {
					v154 = v130&(int32(3)<<(uint(v134)%32)^int32(-1)) | l3<<(uint(v134)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v154)
					if l4 == int64(0) {
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+36))
						v166 = v158 + v25<<(uint(int32(13))%32) + int32(base.Ui32(v126)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
						v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)))
						if base.Ui64(l4) <= base.Ui64(v167) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v166))) = l4
						}
					}
				} else {
					if int32(base.Ui32(v130)>>(uint(v134)%32))&int32(3) == int32(1) {
					} else {
						v154 = v130&(int32(3)<<(uint(v134)%32)^int32(-1)) | l3<<(uint(v134)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v154)
						if l4 == int64(0) {
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+36))
							v166 = v158 + v25<<(uint(int32(13))%32) + int32(base.Ui32(v126)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)))
							if base.Ui64(l4) <= base.Ui64(v167) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v166))) = l4
							}
						}
					}
				}
			}
		}
		if int32(0) < l1 {
			v189 = int32(0)
			for {
				v203 = int32(2)
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l2+v189<<(uint(v203)%32))))
				v210 = int32(base.Ui32(v206&int32(_a_F_TransactionIdSetPageStatusInternal_1)) >> (uint(v203) % 32))
				v212 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
				v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v25<<(uint(v203)%32))))
				v218 = v210 + v217
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
				v223 = v206 << (uint(int32(1)) % 32) & int32(6)
				if l3 != int32(3) {
					v243 = v219&(int32(3)<<(uint(v223)%32)^int32(-1)) | l3<<(uint(v223)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v243)
					if l4 == int64(0) {
					} else {
						v246 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
						v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
						v251 = v247 + v25<<(uint(int32(13))%32) + v210&int32(_a_F_TransactionIdSetPageStatusInternal_2)
						v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
						if base.Ui64(l4) <= base.Ui64(v252) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v251))) = l4
						}
					}
				} else {
					v227 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
					if v227&int32(1) == int32(0) {
						v243 = v219&(int32(3)<<(uint(v223)%32)^int32(-1)) | l3<<(uint(v223)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v243)
						if l4 == int64(0) {
						} else {
							v246 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
							v251 = v247 + v25<<(uint(int32(13))%32) + v210&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
							if base.Ui64(l4) <= base.Ui64(v252) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v251))) = l4
							}
						}
					} else {
						if int32(base.Ui32(v219)>>(uint(v223)%32))&int32(3) == int32(1) {
						} else {
							v243 = v219&(int32(3)<<(uint(v223)%32)^int32(-1)) | l3<<(uint(v223)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v243)
							if l4 == int64(0) {
							} else {
								v246 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
								v251 = v247 + v25<<(uint(int32(13))%32) + v210&int32(_a_F_TransactionIdSetPageStatusInternal_2)
								v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
								if base.Ui64(l4) <= base.Ui64(v252) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v251))) = l4
								}
							}
						}
					}
				}
				v257 = v189 + int32(1)
				if v257 != l1 {
					v189 = v257
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v274 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
		v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
		v277 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v275+v25))) = uint8(v277)
		m.G0 = v17 + int32(16)
		return
	}
}
func F_TransactionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	return
}
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = l2
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(267)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v10)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockApplyTransactionForSession[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	v19 = F_LockRelease(m, v8, l3, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_check_transaction_read_only(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		v70 = v4
		return v70
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[0])))
		if v7&int32(1) == int32(0) {
			v70 = v4
			return v70
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
			if base.B2i32(v14 == int32(2)) == int32(0) {
				return int32(1)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[2])))
				if v22&int32(1) != 0 {
					v70 = v4
					return v70
				} else {
					v25 = int32(16777538)
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[1]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
					if int32(1) < v29 {
						v53 = v25
						v54 = int32(_a_F_check_transaction_read_only_0)
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
						v64 = F_format_elog_string(m, v54, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
							v70 = int32(0)
							return v70
						}
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[7])))
						if v34 != 0 {
							v53 = v25
							v54 = int32(_a_F_check_transaction_read_only_1)
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
							v64 = F_format_elog_string(m, v54, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
								v70 = int32(0)
								return v70
							}
						} else {
							v35 = int32(1)
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[8])))
							if v38 == v35 {
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[9]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+308))
								v46 = base.B2i32(v44 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[8])) = uint8(v46)
								v48 = v46
							} else {
								v48 = int32(0)
							}
							if v48 == int32(0) {
								v70 = v35
								return v70
							} else {
								v53 = int32(1088)
								v54 = int32(_a_F_check_transaction_read_only_2)
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
								v64 = F_format_elog_string(m, v54, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
									v70 = int32(0)
									return v70
								}
							}
						}
					}
				}
			}
		}
	}
}
