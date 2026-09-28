package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_exec_stmt_execsql(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[0]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[1])))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_exec_prepare_plan(m, l0, v20, int32(2048))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 == v30 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v96 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v83)
	goto L8
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v36
	goto L14
L13:
	;
	v42 = v39
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v48 = int32(0)
	goto L15
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43+v48<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	switch v60 - int32(158) {
	case 0, 5:
		goto L18
	case 1, 2, 3, 4:
		goto L17
	default:
		goto L19
	}
L16:
	;
	goto L9
L17:
	;
	v70 = v48 + int32(1)
	if v70 != v42 {
		v48 = v70
		goto L15
	} else {
		goto L22
	}
L18:
	;
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v67)
	goto L9
L19:
	;
	if v60 == int32(192) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v60 != int32(103) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L16
L23:
	;
	m.G0 = v14 - int32(-64)
	return
L24:
	;
	v177 = int32(4)
	v178 = v17 & v177
	v180 = v19 & v177
	v182 = int32(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v184 != 0 {
		goto L44
	} else {
		goto L45
	}
L25:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v99 != int32(1) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+v104<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 != int32(1) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	if v112 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = int32(_a_F_exec_stmt_execsql_0)
	v117 = int32(_a_F_exec_stmt_execsql_1)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2])) = v12 + int32(-12)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v20
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v125 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_exec_prepare_plan(m, l0, v20, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v137 = F_exec_eval_expr(m, l0, v20, v12+int32(-13), v12+int32(-20), v12+int32(-24))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2])) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+36))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+v144<<(uint(int32(2))%32))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+51)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_exec_assign_value(m, l0, v148, v137, v149, v150, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v154 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_SPI_freetuptable(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	F_MemoryContextReset(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163+v164<<(uint(int32(2))%32))))
	v170 = int32(0)
	F_assign_simple_var(m, l0, v168, int64(1), v170, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(1)
	goto L23
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+20)) = v20
	v187 = v185
	goto L46
L45:
	;
	v187 = v182
	goto L46
L46:
	;
	v189 = base.B2i32(v178|v180 != int32(0))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v190 != int32(1) {
		v201 = v182
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	v204 = F_SPI_execute_plan_with_paramlist(m, v202, v187, v203, v201)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L63
	}
L48:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v194 != 0 {
		v201 = int32(2)
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v195 = int32(2)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v198 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v199 = v195
	goto L52
L51:
	;
	v199 = int32(1)
	goto L52
L52:
	;
	if v178|v180 != int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v200 = v195
	goto L55
L54:
	;
	v200 = v199
	goto L55
L55:
	;
	v201 = v200
	goto L47
L56:
	;
	v299 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[3]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[4]))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v303 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L57:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v283+v284<<(uint(int32(2))%32))))
	v290 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[3]))
	v294 = int32(0)
	F_assign_simple_var(m, l0, v288, base.I64_extend_i32_u(base.B2i32(v290 != int64(0))), v294, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L78
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L74
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L70
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L66
	}
L61:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223+v224<<(uint(int32(2))%32))))
	v230 = int32(0)
	F_assign_simple_var(m, l0, v228, int64(0), v230, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L65
	}
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208+v209<<(uint(int32(2))%32))))
	v215 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[3]))
	v219 = int32(0)
	F_assign_simple_var(m, l0, v213, base.I64_extend_i32_u(base.B2i32(v215 != int64(0))), v219, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L64
	}
L63:
	;
	switch v204 + int32(8) {
	case 0:
		goto L59
	default:
		goto L58
	case 6:
		goto L60
	case 12, 14:
		goto L56
	case 13:
		goto L57
	case 15, 16, 17, 19, 20, 21, 26, 27:
		goto L62
	case 22:
		goto L61
	}
L64:
	;
	goto L56
L65:
	;
	goto L56
L66:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_3), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_5), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_7), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_8), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v271 = F_SPI_result_code_string(m, v204)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v270
	F_errmsg_internal(m, int32(_a_F_exec_stmt_execsql_9), v14)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_10), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	goto L56
L79:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L153
	}
L80:
	;
	if v302 == int32(0) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v302 == int32(0) {
		goto L23
	} else {
		goto L144
	}
L83:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308+v310<<(uint(int32(2))%32))))
	if base.Ui64(v299) <= base.Ui64(int64(1)) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	F_exec_move_row(m, l0, v314, v423, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L134
	}
L85:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v423 = v420
	goto L84
L86:
	;
	if base.I32_wrap_i64(v299) == int32(1) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v354 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L89:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v321 != int32(1) {
		v423 = int32(0)
		goto L84
	} else {
		goto L90
	}
L90:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+492)))
	if v326 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v329 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	v331 = int32(0)
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L95
	}
L94:
	;
	v331 = v329
	goto L93
L95:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_11), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	if v331 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v331
	F_errdetail_internal(m, int32(_a_F_exec_stmt_execsql_12), v12+int32(-32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_13), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v189|v357)&int32(1) == int32(0) {
		goto L85
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+492)))
	if v364 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L105
L107:
	;
	v391 = F_errstart(m, v389, int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L124
	}
L108:
	;
	v378 = int32(21)
	if int32(base.Ui32(v180)>>(uint(int32(2))%32)) != 0 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	v367 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v373 = int32(0)
	if v354 == v373 {
		v377 = v373
		goto L108
	} else {
		goto L114
	}
L112:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v369 == int32(0) {
		v377 = v367
		goto L108
	} else {
		goto L113
	}
L113:
	;
	v388 = v367
	v389 = int32(21)
	goto L107
L114:
	;
	v388 = v373
	v389 = int32(21)
	goto L107
L115:
	;
	v384 = int32(19)
	goto L117
L116:
	;
	v384 = int32(0)
	goto L117
L117:
	;
	if v178 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v385 = v378
	goto L120
L119:
	;
	v385 = v384
	goto L120
L120:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v386 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v387 = v378
	goto L123
L122:
	;
	v387 = v385
	goto L123
L123:
	;
	v388 = v377
	v389 = v387
	goto L107
L124:
	;
	if v391 == int32(0) {
		goto L85
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_14), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	if v388 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v388
	F_errdetail_internal(m, int32(_a_F_exec_stmt_execsql_12), v12+int32(-48))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_errhint(m, int32(_a_F_exec_stmt_execsql_15), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_16), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	goto L85
L134:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v427 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	F_SPI_freetuptable(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v432 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	F_MemoryContextReset(m, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[4]))
	F_SPI_freetuptable(m, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	goto L23
L144:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_2))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_17), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v204 == int32(5) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_errhint(m, int32(_a_F_exec_stmt_execsql_18), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_19), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_20), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_4), int32(_a_F_exec_stmt_execsql_21), int32(_a_F_exec_stmt_execsql_6))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformSelectStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v575 int32
	_ = v575
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1188 int32
	_ = v1188
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1446 int32
	_ = v1446
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(48)
	m.G0 = v26
	v29 = F_palloc0(m, int32(168))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(4294967363)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v35 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+41)) = uint8(v36)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v39 = F_transformWithClause(m, l0, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v44 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v39
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+42)) = uint8(v42)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L1
	} else {
		goto L352
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformFromClause(m, l0, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L347
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v56 = F_transformTargetList(m, l0, v54, int32(14))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v56
	v60 = v29 + int32(76)
	if l2 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v161 = F_transformWhereClause(m, l0, v158, int32(6), int32(_a_F_transformSelectStmt_0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v64 = F_exprType(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_markTargetListOrigins(m, l0, v56)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L51
	}
L17:
	;
	v66 = F_exprTypmod(m, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v68 = F_exprCollation(m, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v56 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = F_exprType(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v89
	F_errmsg_plural(m, int32(_a_F_transformSelectStmt_1), int32(_a_F_transformSelectStmt_2), v89, v26)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L30
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v70 == int32(1) {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v81 = int32(0)
	F_errstart_cold(m, int32(21), v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v89 = v80
	goto L21
L28:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v89 = v81
	goto L21
L30:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_3), int32(2938), int32(_a_F_transformSelectStmt_4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(17)
	if v61 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v101
	v145 = F_list_make1_impl(m, int32(1), v26+int32(40))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L50
	}
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v112 = F_exprLocation(m, v63)
	mBase = m.M
	v113 = F_transformAssignmentIndirection(m, l0, v63, v107, int32(0), v64, v66, v68, v61, v109, v110, int32(2), v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v64 == v103 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v113
	goto L33
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v128 = int32(2)
	v131 = F_coerce_to_target_type(m, l0, v127, v103, v64, v66, v128, v128, int32(-1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L48
	}
L39:
	;
	if v64 != int32(2249) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v119 = F_typeOrDomainTypeRelid(m, v64)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v103 == int32(2249) {
		goto L33
	} else {
		goto L45
	}
L43:
	;
	if v119 == int32(0) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v125 = F_typeOrDomainTypeRelid(m, v103)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v125 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L38
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v131
	if v131 == int32(0) {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L33
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v145
	goto L13
L51:
	;
	goto L13
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v166 = F_transformWhereClause(m, l0, v163, int32(7), int32(_a_F_transformSelectStmt_5))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v171 = F_transformSortClause(m, l0, v169, v60, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v176 = v29 + int32(108)
	v179 = F_transformGroupClause(m, l0, v174, v176, v60, v171, int32(19), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v179
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+104)) = uint8(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v184 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v867 = F_transformLimitClause(m, l0, v863, int32(23), int32(_a_F_transformSelectStmt_6), v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L184
	}
L57:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+40)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v187
	goto L56
L58:
	;
	goto L59
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v193 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v197 = F_transformDistinctClause(m, l0, v60, v191, int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v184 == int32(0) {
		v453 = v4
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+40)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v197
	goto L56
L64:
	;
	v837 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+40)) = uint8(v837)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v627
	goto L56
L65:
	;
	if v191 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v204 <= int32(0) {
		v453 = v4
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v214 = v4
	v219 = v4
	goto L68
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230+v219<<(uint(int32(2))%32))))
	v236 = F_findTargetlistEntrySQL92(m, l0, v234, v60, int32(21))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v453 = v440
	goto L65
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
	if v238 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v242 == int32(0) {
		v415 = int32(1)
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v420 = v238
	goto L73
L73:
	;
	v440 = F_lappend_int(m, v214, v420)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L103
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+16)) = v415
	v420 = v415
	goto L73
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v246 <= int32(0) {
		v415 = int32(1)
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v250 = v246 & int32(3)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v252 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v246) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v415 = v370 + int32(1)
	goto L74
L78:
	;
	v263 = v252
	v264 = v252
	v268 = int32(0)
	goto L81
L79:
	;
	v312 = v252
	v313 = v252
	goto L80
L80:
	;
	v334 = v252
	v335 = v312
	v336 = v313
	goto L97
L81:
	;
	v285 = v251 + v264<<(uint(int32(2))%32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+16))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	if base.Ui32(v263) < base.Ui32(v293) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v250 == int32(0) {
		v370 = v301
		goto L77
	} else {
		goto L96
	}
L83:
	;
	v295 = v293
	goto L85
L84:
	;
	v295 = v263
	goto L85
L85:
	;
	if base.Ui32(v295) < base.Ui32(v291) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v297 = v291
	goto L88
L87:
	;
	v297 = v295
	goto L88
L88:
	;
	if base.Ui32(v297) < base.Ui32(v289) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v299 = v289
	goto L91
L90:
	;
	v299 = v297
	goto L91
L91:
	;
	if base.Ui32(v299) < base.Ui32(v287) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v301 = v287
	goto L94
L93:
	;
	v301 = v299
	goto L94
L94:
	;
	v302 = int32(4)
	v303 = v264 + v302
	v305 = v268 + v302
	if v305 != v246&int32(2147483644) {
		v263 = v301
		v264 = v303
		v268 = v305
		goto L81
	} else {
		goto L95
	}
L95:
	;
	goto L82
L96:
	;
	v312 = v301
	v313 = v303
	goto L80
L97:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v251+v336<<(uint(int32(2))%32))))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+16))
	if base.Ui32(v335) < base.Ui32(v359) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v370 = v361
	goto L77
L99:
	;
	v361 = v359
	goto L101
L100:
	;
	v361 = v335
	goto L101
L101:
	;
	v362 = int32(1)
	v365 = v334 + v362
	if v365 != v250 {
		v334 = v365
		v335 = v361
		v336 = v336 + v362
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	v443 = v219 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v443 < v444 {
		v214 = v440
		v219 = v443
		goto L68
	} else {
		goto L104
	}
L104:
	;
	goto L69
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L179
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L160
	}
L107:
	;
	v587 = int32(0)
	v591 = v561
	goto L136
L108:
	;
	v561 = int32(0)
	v575 = v4
	goto L107
L109:
	;
	goto L110
L110:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v472 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v561 = int32(0)
	v575 = v4
	goto L107
L112:
	;
	goto L113
L113:
	;
	v476 = int32(0)
	v480 = v476
	v481 = v476
	v494 = v4
	goto L114
L114:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v481<<(uint(int32(2))%32))))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	v507 = int32(0)
	if v453 == v507 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v561 = v552
	v575 = v554
	goto L107
L116:
	;
	if v545 != 0 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v545 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v513 <= int32(0) {
		v539 = v507
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v545 = v539
	goto L116
L121:
	;
	v516 = int32(0)
	if v516 < v513 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v519 = v513
	goto L124
L123:
	;
	v519 = v516
	goto L124
L124:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v522 = int32(0)
	goto L125
L125:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v520+v522<<(uint(int32(2))%32))))
	v531 = base.B2i32(v530 == v506)
	if v530 == v506 {
		v539 = v531
		goto L120
	} else {
		goto L127
	}
L126:
	;
	v539 = v531
	goto L120
L127:
	;
	v533 = v522 + int32(1)
	if v533 != v519 {
		v522 = v533
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	if v494&int32(1) != 0 {
		goto L106
	} else {
		goto L132
	}
L130:
	;
	v552 = v480
	goto L131
L131:
	;
	v553 = int32(1)
	v554 = v545 ^ v553
	v556 = v481 + v553
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v556 < v557 {
		v480 = v552
		v481 = v556
		v494 = v554
		goto L114
	} else {
		goto L135
	}
L132:
	;
	v548 = F_copyObjectImpl(m, v505)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v550 = F_lappend(m, v480, v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v552 = v550
	goto L131
L135:
	;
	goto L115
L136:
	;
	v606 = int32(0)
	if v184 == v606 {
		v616 = v606
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v453 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v610 <= v587 {
		v616 = int32(0)
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v616 = v612 + v587<<(uint(int32(2))%32)
	goto L138
L141:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v624+v587<<(uint(int32(2))%32))))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v634 = F_get_sortgroupref_tle(m, v632, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L151
	}
L142:
	;
	v617 = int32(0)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if base.B2i32(v616 == v617)|base.B2i32(v619 <= v587) == v617 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v627 = v561
	goto L144
L144:
	;
	goto L64
L145:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	if v624 != 0 {
		goto L141
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v627 = v591
	goto L144
L148:
	;
	goto L147
L149:
	;
	v587 = v587 + int32(1)
	v591 = v714
	goto L136
L150:
	;
	if v575 != 0 {
		goto L105
	} else {
		goto L158
	}
L151:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v634)+16))
	v637 = int32(0)
	if base.B2i32(v636 == v637)|base.B2i32(v591 == v637) != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	if v642 <= int32(0) {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v591)+12))
	v650 = int32(0)
	goto L154
L154:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v645+v650<<(uint(int32(2))%32))))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	if v674 == v636 {
		v714 = v591
		goto L149
	} else {
		goto L156
	}
L155:
	;
	goto L150
L156:
	;
	v677 = v650 + int32(1)
	if v677 != v642 {
		v650 = v677
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v703 = F_exprLocation(m, v628)
	mBase = m.M
	v704 = F_addTargetToGroupList(m, l0, v634, v591, v702, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v714 = v704
	goto L149
L160:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_7))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_transformSelectStmt_8), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	v744 = int32(0)
	goto L164
L163:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v785+v744<<(uint(int32(2))%32))))
	v810 = F_exprLocation(m, v809)
	mBase = m.M
	F_parser_errposition(m, l0, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L177
	}
L164:
	;
	v767 = int32(0)
	if v453 == v767 {
		v777 = v767
		goto L166
	} else {
		goto L167
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L174
	}
L166:
	;
	if v184 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v771 <= v744 {
		v777 = int32(0)
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v777 = v773 + v744<<(uint(int32(2))%32)
	goto L166
L169:
	;
	goto L165
L170:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if base.B2i32(v777 == int32(0))|base.B2i32(v782 <= v744) != 0 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	if v785 == int32(0) {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v777)))
	if v788 == v742 {
		goto L163
	} else {
		goto L173
	}
L173:
	;
	v744 = v744 + int32(1)
	goto L164
L174:
	;
	F_errmsg_internal(m, int32(_a_F_transformSelectStmt_9), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(3189), int32(_a_F_transformSelectStmt_11))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(3122), int32(_a_F_transformSelectStmt_12))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_7))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_errmsg(m, int32(_a_F_transformSelectStmt_8), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v829 = F_exprLocation(m, v628)
	mBase = m.M
	F_parser_errposition(m, l0, v829)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(3151), int32(_a_F_transformSelectStmt_12))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v867
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v874 = F_transformLimitClause(m, l0, v870, int32(22), int32(_a_F_transformSelectStmt_13), v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+132)) = v874
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+136)) = v877
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v880 = int32(0)
	v882 = m.G0
	v884 = v882 - int32(112)
	m.G0 = v884
	if v879 == v880 {
		v1446 = v880
		goto L192
	} else {
		goto L193
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+116)) = v1446
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)))
	if v1560 == int32(1) {
		goto L327
	} else {
		goto L328
	}
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L322
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L319
	}
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L314
	}
L190:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L1
	} else {
		goto L309
	}
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L304
	}
L192:
	;
	m.G0 = v884 + int32(112)
	goto L186
L193:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v879)+4))
	if v888 <= int32(0) {
		v1446 = v880
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v899 = v880
	v900 = v880
	goto L195
L195:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v879)+12))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v914+v899<<(uint(int32(2))%32))))
	v919 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v884)+108)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v884)+104)) = v919
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v918)+4))
	if v923 == v919 {
		goto L203
	} else {
		goto L204
	}
L196:
	;
	v1446 = v1433
	goto L192
L197:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v918)+16))
	if v1199 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L198:
	;
	v1188 = int32(0)
	goto L197
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L244
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L239
	}
L201:
	;
	if v900 == int32(0) {
		goto L200
	} else {
		goto L223
	}
L202:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	if v1016 == int32(0) {
		goto L198
	} else {
		goto L222
	}
L203:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	if v1015 != 0 {
		goto L201
	} else {
		goto L221
	}
L204:
	;
	if v900 == int32(0) {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
	if v928 <= int32(0) {
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	v935 = int32(0)
	goto L207
L207:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v931+v935<<(uint(int32(2))%32))))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v960 != 0 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	goto L203
L209:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923))))
	if base.B2i32(v963 == int32(0))|base.B2i32(v963 != v966) != 0 {
		v984 = v963
		v985 = v966
		goto L213
	} else {
		goto L214
	}
L210:
	;
	goto L211
L211:
	;
	v990 = v935 + int32(1)
	if v928 != v990 {
		v935 = v990
		goto L207
	} else {
		goto L220
	}
L212:
	;
	if v984-v985 == int32(0) {
		goto L199
	} else {
		goto L219
	}
L213:
	;
	goto L212
L214:
	;
	v969 = v960
	v970 = v923
	goto L215
L215:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+1)))
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969)+1)))
	if v974 == int32(0) {
		v984 = v974
		v985 = v973
		goto L213
	} else {
		goto L217
	}
L216:
	;
	v984 = v974
	v985 = v973
	goto L213
L217:
	;
	v977 = int32(1)
	if v974 == v973 {
		v969 = v969 + v977
		v970 = v970 + v977
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	goto L211
L220:
	;
	goto L208
L221:
	;
	goto L198
L222:
	;
	goto L200
L223:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
	if v1021 <= int32(0) {
		goto L200
	} else {
		goto L224
	}
L224:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	v1028 = int32(0)
	goto L225
L225:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1024+v1028<<(uint(int32(2))%32))))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+4))
	if v1053 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L200
L227:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053))))
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015))))
	if base.B2i32(v1056 == int32(0))|base.B2i32(v1056 != v1059) != 0 {
		v1077 = v1056
		v1078 = v1059
		goto L231
	} else {
		goto L232
	}
L228:
	;
	goto L229
L229:
	;
	v1083 = v1028 + int32(1)
	if v1021 != v1083 {
		v1028 = v1083
		goto L225
	} else {
		goto L238
	}
L230:
	;
	if v1077-v1078 == int32(0) {
		v1188 = v1052
		goto L197
	} else {
		goto L237
	}
L231:
	;
	goto L230
L232:
	;
	v1062 = v1053
	v1063 = v1015
	goto L233
L233:
	;
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+1)))
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+1)))
	if v1067 == int32(0) {
		v1077 = v1067
		v1078 = v1066
		goto L231
	} else {
		goto L235
	}
L234:
	;
	v1077 = v1067
	v1078 = v1066
	goto L231
L235:
	;
	v1070 = int32(1)
	if v1067 == v1066 {
		v1062 = v1062 + v1070
		v1063 = v1063 + v1070
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	goto L229
L238:
	;
	goto L226
L239:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+80)) = v1115
	F_errmsg(m, int32(_a_F_transformSelectStmt_14), v884+int32(80))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1122)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2808), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v918)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+96)) = v1137
	F_errmsg(m, int32(_a_F_transformSelectStmt_17), v884+int32(96))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2795), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v918)+12))
	v1274 = F_transformGroupClause(m, l0, v1270, int32(0), v60, v1250, int32(9), int32(1))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L259
	}
L250:
	;
	v1250 = int32(0)
	goto L249
L251:
	;
	goto L252
L252:
	;
	v1203 = int32(0)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	if v1205 <= v1203 {
		v1250 = v1203
		goto L249
	} else {
		goto L253
	}
L253:
	;
	v1210 = v1203
	v1211 = v1203
	goto L254
L254:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+12))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1231+v1210<<(uint(int32(2))%32))))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1235)+4))
	v1238 = F_findTargetlistEntrySQL99(m, l0, v1236, v60, int32(10))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L256
	}
L255:
	;
	v1250 = v1241
	goto L249
L256:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v1241 = F_addTargetToSortList(m, l0, v1238, v1211, v1240, v1235)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1244 = v1210 + int32(1)
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	if v1244 < v1245 {
		v1210 = v1244
		v1211 = v1241
		goto L254
	} else {
		goto L258
	}
L258:
	;
	goto L255
L259:
	;
	v1277 = F_palloc0(m, int32(56))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1277))) = int32(108)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v918)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+4)) = v1281
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+8)) = v1283
	if v1188 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v918)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+20)) = v1362
	v1366 = int32(0)
	if base.B2i32(v1362&int32(2) == v1366)|base.B2i32(v1362&int32(_a_F_transformSelectStmt_18) == v1366) == v1366 {
		goto L287
	} else {
		goto L288
	}
L262:
	;
	if v1274 != 0 {
		goto L191
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v1356 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1277)+52)) = uint8(v1356)
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+16)) = v1250
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+12)) = v1274
	v1360 = v1250
	goto L261
L265:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+12))
	v1286 = F_copyObjectImpl(m, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+12)) = v1286
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+16))
	if v1250 != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1277)+52)) = uint8(v1318)
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+16)) = v1319
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+20))
	if v1322 == int32(1058) {
		v1360 = v1319
		goto L261
	} else {
		goto L278
	}
L268:
	;
	v1290 = int32(0)
	if v1289 == v1290 {
		v1318 = v1290
		v1319 = v1250
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1316 = F_copyObjectImpl(m, v1289)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L277
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+48)) = v1300
	F_errmsg(m, int32(_a_F_transformSelectStmt_19), v884+int32(48))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1307)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2869), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	v1318 = int32(1)
	v1319 = v1316
	goto L267
L278:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v918)+4))
	if v1250|v1325 != 0 {
		goto L190
	} else {
		goto L279
	}
L279:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v918)+20))
	if v1327 != int32(1058) {
		goto L190
	} else {
		goto L280
	}
L280:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+32)) = v1337
	F_errmsg(m, int32(_a_F_transformSelectStmt_20), v884+int32(32))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_errhint(m, int32(_a_F_transformSelectStmt_21), int32(0))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1348)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2906), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	if v1360 == int32(0) {
		goto L189
	} else {
		goto L290
	}
L288:
	;
	v1407 = v1362
	goto L289
L289:
	;
	if v1407&int32(8) != 0 {
		goto L296
	} else {
		goto L297
	}
L290:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+4))
	if v1377 != int32(1) {
		goto L189
	} else {
		goto L291
	}
L291:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+12))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v1383 = F_get_sortgroupclause_expr(m, v1381, v1382)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+12))
	v1392 = F_get_ordering_op_properties(m, v1385, v884+int32(108), v884+int32(104), v884+int32(100))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	if v1392 == int32(0) {
		goto L188
	} else {
		goto L294
	}
L294:
	;
	v1396 = F_exprCollation(m, v1383)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+40)) = v1396
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+16)))
	v1401 = v1399 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1277)+44)) = uint8(v1401)
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1277)+45)) = uint8(v1403)
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+20))
	v1407 = v1405
	goto L289
L296:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+16))
	if v1410 == int32(0) {
		goto L187
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v884)+108))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v884)+104))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v918)+24))
	v1418 = F_transformFrameOffset(m, l0, v1407, v1413, v1414, v1277+int32(32), v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L300
	}
L299:
	;
	goto L298
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+24)) = v1418
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+20))
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v884)+108))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v884)+104))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v918)+28))
	v1427 = F_transformFrameOffset(m, l0, v1421, v1422, v1423, v1277+int32(36), v1426)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1430 = v899 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+48)) = v1430
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+28)) = v1427
	v1433 = F_lappend(m, v900, v1277)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v879)+4))
	if v1430 < v1435 {
		v899 = v1430
		v900 = v1433
		goto L195
	} else {
		goto L303
	}
L303:
	;
	goto L196
L304:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+64)) = v1470
	F_errmsg(m, int32(_a_F_transformSelectStmt_22), v884-int32(-64))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1477)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2857), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+16)) = v1492
	F_errmsg(m, int32(_a_F_transformSelectStmt_20), v884+int32(16))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1499)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2899), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errmsg(m, int32(_a_F_transformSelectStmt_23), int32(0))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1518)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2926), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = v1530
	F_errmsg_internal(m, int32(_a_F_transformSelectStmt_24), v884)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2935), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	F_errcode(m, int32(_a_F_transformSelectStmt_16))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	F_errmsg(m, int32(_a_F_transformSelectStmt_25), int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	F_parser_errposition(m, l0, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_10), int32(2949), int32(_a_F_transformSelectStmt_15))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	F_resolveTargetListUnknowns(m, l0, v1563)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1566
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1568
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1571 = F_makeFromExpr(m, v1570, v161)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L331
	}
L330:
	;
	goto L329
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1571
	v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+39)) = uint8(v1574)
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+37)) = uint8(v1576)
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+38)) = uint8(v1578)
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+36)) = uint8(v1580)
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v1582 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	F_assign_query_collations(m, l0, v29)
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L1
	} else {
		goto L339
	}
L333:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	if v1585 <= int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1590 = int32(0)
	goto L335
L335:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+12))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1612+v1590<<(uint(int32(2))%32))))
	F_transformLockingClause(m, l0, v29, v1616, int32(0))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L337
	}
L336:
	;
	goto L332
L337:
	;
	v1621 = v1590 + int32(1)
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	if v1621 < v1622 {
		v1590 = v1621
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v1649 != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	m.G0 = v26 + int32(48)
	return v29
L341:
	;
	F_parseCheckAggregates(m, l0, v29)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L346
	}
L342:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
	if v1650 != 0 {
		goto L341
	} else {
		goto L343
	}
L343:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v1651 != 0 {
		goto L341
	} else {
		goto L344
	}
L344:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v1652 == int32(0) {
		goto L340
	} else {
		goto L345
	}
L345:
	;
	goto L341
L346:
	;
	goto L340
L347:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	F_errmsg(m, int32(_a_F_transformSelectStmt_26), int32(0))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1673 = F_exprLocation(m, v1672)
	mBase = m.M
	F_parser_errposition(m, l0, v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_3), int32(1444), int32(_a_F_transformSelectStmt_27))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v1689 = F_format_type_be(m, v64)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v1691 = F_format_type_be(m, v103)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v1691
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v1689
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v1688
	F_errmsg(m, int32(_a_F_transformSelectStmt_28), v26+int32(16))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	F_errhint(m, int32(_a_F_transformSelectStmt_29), int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1705 = F_exprLocation(m, v127)
	mBase = m.M
	F_parser_errposition(m, l0, v1705)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	F_errfinish(m, int32(_a_F_transformSelectStmt_3), int32(3002), int32(_a_F_transformSelectStmt_4))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
