package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LagTrackerRead(m *base.Module, l0 int32, l1 int64, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v113 int64
	_ = v113
	var v128 int32
	_ = v128
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v166 int64
	_ = v166
	var v172 float64
	_ = v172
	var v176 int64
	_ = v176
	var v181 int64
	_ = v181
	var v184 int64
	_ = v184
	var v192 int64
	_ = v192
	v17 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v20 = v17 + l0<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[708])))
	if v23 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v56 == v58 {
		v113 = v59
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[710])))
	v56 = v23
	v58 = v26
	v59 = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = v17 + l0<<(uint(int32(4))%32)
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[711])))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[712])))
	if base.Ui64(l1) < base.Ui64(v35) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 < v32 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[711])))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[713]))) = v44
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[712])))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[714]))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[710])))
	v54 = base.I32_rem_s(v50+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[708]))) = v54
	v56 = v54
	v58 = v50
	v59 = v32
	goto L1
L8:
	;
	v40 = int64(-1)
	goto L10
L9:
	;
	v40 = l2 - v32
	goto L10
L10:
	;
	return v40
L11:
	;
	v140 = int64(-1)
	if l2 < v135 {
		v192 = v140
		goto L19
	} else {
		goto L20
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17+l0<<(uint(int32(4))%32))+uint32(_consts[713]))) = int64(0)
	v128 = v58
	v135 = v113
	goto L11
L13:
	;
	v62 = v17 + int32(8)
	v65 = v62 + v56<<(uint(int32(4))%32)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	if base.Ui64(l1) < base.Ui64(v66) {
		v128 = v56
		v135 = v59
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v70 = v17 + l0<<(uint(int32(4))%32)
	v76 = v56
	v78 = v65
	goto L15
L15:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[713]))) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[714]))) = v90
	v95 = base.I32_rem_s(v76+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[708]))) = v95
	if v95 == v58 {
		v113 = v88
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v128 = v95
	v135 = v88
	goto L11
L17:
	;
	v100 = v62 + v95<<(uint(int32(4))%32)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
	if base.Ui64(v101) <= base.Ui64(l1) {
		v76 = v95
		v78 = v100
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	return v192
L20:
	;
	if v135 != int64(0) {
		v184 = v135
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v192 = l2 - v184
	goto L19
L22:
	;
	if v128 == v58 {
		v192 = v140
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v147 = v17 + l0<<(uint(int32(4))%32)
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[713])))
	if v150 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[714])))
	if base.Ui64(l1) < base.Ui64(v155) {
		v192 = v140
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v17+v128<<(uint(int32(4))%32))+16))
	v184 = v181
	goto L21
L27:
	;
	v159 = v17 + v128<<(uint(int32(4))%32)
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+16))
	if v160 < v150 {
		v192 = v140
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	v172 = base.F64_add(base.F64_mul(base.F64_convert_i64_s(v160-v150), base.F64_div(base.F64_convert_i64_u(l1-v155), base.F64_convert_i64_u(v166-v155))), base.F64_convert_i64_s(v150))
	if base.F64_lt(base.F64_abs(v172), float64(9.223372036854776e+18)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v176 = base.I64_trunc_f64_s(v172)
	v184 = v176
	goto L21
L30:
	;
	goto L31
L31:
	;
	v184 = int64(-9223372036854775807 - 1)
	goto L21
}
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v215 int64
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v234 int64
	_ = v234
	var v238 int64
	_ = v238
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v22-int32(3))&int32(255)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L97
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L11
	} else {
		goto L94
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L91
	}
L4:
	;
	if l1 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L11
	} else {
		goto L88
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_consts[100])))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 < l1 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+96)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = l1
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v50 = F_hash_search(m, v45, v20+int32(88), v43, v43)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	m.G0 = v20 + int32(112)
	return v460
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	return int32(0)
L12:
	;
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)+32))
	if int64(0) < v54 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v59 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v59 == int32(0) {
		v460 = v43
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v67
	F_errmsg_internal(m, int32(202680), v20+int32(32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(521220), int32(2111), int32(378753))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v460 = v43
	goto L9
L21:
	;
	v82 = int32(0)
	goto L23
L22:
	;
	v82 = v81
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v88 = v83
	goto L28
L24:
	;
	F_RemoveLocalLock(m, v50)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L11
	} else {
		goto L87
	}
L25:
	;
	v362 = int32(1) << (uint(l1) % 32)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	if v362&v363 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L26:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[805]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v337 = int32(0)
	v339 = F_hash_search_with_hash_value(m, v335, l0, v336, v337, v337)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L11
	} else {
		goto L67
	}
L27:
	;
	v311 = int32(0)
	v314 = F_errstart(m, int32(19), v311)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L11
	} else {
		goto L63
	}
L28:
	;
	v103 = v88 - int32(1)
	if v103 < int32(0) {
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v108)+8))
	v113 = v111 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v113
	if v113 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v108 = v84 + v103<<(uint(int32(4))%32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 != v82 {
		v88 = v103
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v50)+32))
	v135 = v133 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+32)) = v135
	if int64(0) < v135 {
		v460 = int32(1)
		goto L9
	} else {
		goto L39
	}
L33:
	;
	if v82 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ResourceOwnerForgetLock(m, v82, v50)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v121 = v119 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = v121
	if v121 <= v103 {
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v126 = v84 + v121<<(uint(int32(4))%32)
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v129
	goto L32
L39:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+53)) = uint8(v140)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v142 != int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v303 = v295 + v296&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v305 = F_LWLockAcquire(m, v303, int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L61
	}
L41:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v145 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(3) < l1 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v149 != v150 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	if v149 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32((v155-int32(1))&(v158*int32(49157))<<(uint(int32(2))%32))+uint32(_consts[807])))
	if v166 <= int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v174 = F_LWLockAcquire(m, v170+int32(584), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v177 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v180 = int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = (v179 - v180) & (v182 * int32(49157))
	v187 = v185 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[807]))) = v177
	v197 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v201 = v185 & int32(268435455) << (uint(int32(3)) % 32)
	v204 = v197
	v208 = v177
	v215 = int64(0)
	goto L48
L48:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v204)+604))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219+base.I32_wrap_i64(v215)<<(uint(int32(2))%32)+v185<<(uint(int32(6))%32))))
	if v182 != v225 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	F_LWLockRelease(m, v272+int32(584))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L59
	}
L50:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+600))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v255+v201)))
	if int64(base.Ui64(v257)>>(uint(v252)%64))&int64(7) != int64(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v249 = v204
	v250 = v208
	v252 = v215 * int64(3)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v204)+600))
	v230 = v229 + v201
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	v234 = v215 * int64(3)
	v238 = int64(1) << (uint(base.I64_extend_i32_u(l1-v180+base.I32_wrap_i64(v234))) % 64)
	if v231&v238 == int64(0) {
		v249 = v204
		v250 = v208
		v252 = v234
		goto L50
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v231 & (v238 ^ int64(-1))
	v247 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v249 = v247
	v250 = int32(1)
	v252 = v234
	goto L50
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[807])))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[807]))) = v263 + int32(1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v268 = v215 + int64(1)
	if v268 != int64(16) {
		v204 = v249
		v208 = v250
		v215 = v268
		goto L48
	} else {
		goto L58
	}
L58:
	;
	goto L49
L59:
	;
	if v250 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	goto L40
L61:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	if v307 == int32(0) {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	v359 = v310
	v360 = v307
	goto L25
L63:
	;
	if v314 == int32(0) {
		v460 = v311
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v322
	F_errmsg_internal(m, int32(202680), v20-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(521220), int32(2150), int32(378753))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v460 = v311
	goto L9
L67:
	;
	if v339 == int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v339
	v346 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v346
	v349 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v352 = int32(0)
	v354 = F_hash_search(m, v349, v20+int32(80), v352, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = v354
	if v354 == int32(0) {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v359 = v354
	v360 = v339
	goto L25
L71:
	;
	F_LWLockRelease(m, v303)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L11
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v360)+84))
	v393 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+84)) = v392 - v393
	v397 = l1 << (uint(int32(2)) % 32)
	v398 = v360 + v397
	v400 = v398 + int32(44)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v401 - v393
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v360)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+128)) = v405 - v393
	v410 = v398 + int32(88)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v413 = v411 - v393
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v413
	v416 = v362 ^ int32(-1)
	if v413 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v369 = int32(0)
	v372 = F_errstart(m, int32(19), v369)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	if v372 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v378
	F_errmsg_internal(m, int32(202680), v20+int32(48))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_RemoveLocalLock(m, v50)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L11
	} else {
		goto L81
	}
L79:
	;
	F_errfinish(m, int32(521220), int32(2246), int32(378753))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v460 = v369
	goto L9
L82:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v419 & v416
	goto L84
L83:
	;
	goto L84
L84:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v422+v397)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v360)+20))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v426 & v416
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	F_CleanUpLock(m, v360, v359, v35, v429, base.B2i32(v425&v424 != int32(0)))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	F_LWLockRelease(m, v303)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	goto L24
L87:
	;
	v460 = int32(1)
	goto L9
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	F_errmsg_internal(m, int32(508766), v20)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(521220), int32(2082), int32(378753))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l1
	F_errmsg_internal(m, int32(508312), v20+int32(16))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(521220), int32(2085), int32(378753))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errmsg_internal(m, int32(118524), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(521220), int32(2221), int32(378753))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L11
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
	F_errmsg_internal(m, int32(118483), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(521220), int32(2231), int32(378753))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LruDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l0 * int32(48)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v13 = v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	F_pgaio_closing_fd(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v18 = F_close(m, v17)
		mBase = m.M
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
			v50 = int32(4470240)
			v52 = *(*int32)(unsafe.Add(mBase, _consts[755]))
			*(*int32)(unsafe.Add(mBase, _consts[755])) = v52 - int32(1)
			v57 = *(*int32)(unsafe.Add(mBase, _consts[265]))
			v58 = v57 + v10
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
			v60 = int32(48)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v57+v59*v60)+16)) = v63
			*(*int32)(unsafe.Add(mBase, uint32(v57+v63*v60)+20)) = v59
			m.G0 = v7 + int32(16)
			return
		} else {
			v21 = int32(15)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[762])))
			if v25&int32(1) != 0 {
				v28 = v21
			} else {
				v28 = int32(23)
			}
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
			if v29&int32(4) != 0 {
				v32 = v21
			} else {
				v32 = v28
			}
			v34 = F_errstart(m, v32, int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if v34 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
					v50 = int32(4470240)
					v52 = *(*int32)(unsafe.Add(mBase, _consts[755]))
					*(*int32)(unsafe.Add(mBase, _consts[755])) = v52 - int32(1)
					v57 = *(*int32)(unsafe.Add(mBase, _consts[265]))
					v58 = v57 + v10
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
					v60 = int32(48)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v57+v59*v60)+16)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v57+v63*v60)+20)) = v59
					m.G0 = v7 + int32(16)
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v38
					F_errmsg_internal(m, int32(313009), v7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, int32(523612), int32(1313), int32(368255))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
							v50 = int32(4470240)
							v52 = *(*int32)(unsafe.Add(mBase, _consts[755]))
							*(*int32)(unsafe.Add(mBase, _consts[755])) = v52 - int32(1)
							v57 = *(*int32)(unsafe.Add(mBase, _consts[265]))
							v58 = v57 + v10
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
							v60 = int32(48)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v57+v59*v60)+16)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(v57+v63*v60)+20)) = v59
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = m.Wasi_snapshot_preview1.Fd_seek(m, l0, l1, l2&int32(255), v7+int32(8))
	mBase = m.M
	if v13 == int32(0) {
		v20 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[158])) = v13
		v20 = int32(-1)
	}
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	m.G0 = v7 + int32(16)
	if v20 != 0 {
		v26 = int64(-1)
	} else {
		v26 = v21
	}
	return v26
}
func F__ltq_rregex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(5587), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_l2_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 float32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 float32
	_ = v57
	var v59 float32
	_ = v59
	var v60 float32
	_ = v60
	var v63 float32
	_ = v63
	var v65 float32
	_ = v65
	var v66 float32
	_ = v66
	var v69 float32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 float32
	_ = v84
	var v89 int32
	_ = v89
	var v91 float32
	_ = v91
	var v93 float32
	_ = v93
	var v94 float32
	_ = v94
	var v99 float32
	_ = v99
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v10 = float32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
			if v24 == v25 {
				v27 = base.I32_extend16_s(v24)
				if v27 <= int32(0) {
					v112 = float64(0)
				} else {
					v31 = int32(8)
					v32 = v22 + v31
					v34 = v17 + v31
					if v27 == int32(1) {
						v75 = int32(0)
						v84 = v10
					} else {
						v41 = int32(0)
						v47 = int32(0)
						v50 = v10
						for {
							v52 = int32(2)
							v53 = v41 << (uint(v52) % 32)
							v55 = v53 | int32(4)
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v34+v55)))
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v32+v55)))
							v60 = base.F32_sub(v57, v59)
							v63 = *(*float32)(unsafe.Add(mBase, uint32(v53+v34)))
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v53+v32)))
							v66 = base.F32_sub(v63, v65)
							v69 = base.F32_add(base.F32_mul(v60, v60), base.F32_add(base.F32_mul(v66, v66), v50))
							v71 = v41 + v52
							v73 = v47 + v52
							if v73 != v27&int32(32766) {
								v41 = v71
								v47 = v73
								v50 = v69
								continue
							} else {
								break
							}
							break
						}
						v75 = v71
						v84 = v69
					}
					if v27&int32(1) != 0 {
						v89 = v75 << (uint(int32(2)) % 32)
						v91 = *(*float32)(unsafe.Add(mBase, uint32(v34+v89)))
						v93 = *(*float32)(unsafe.Add(mBase, uint32(v89+v32)))
						v94 = base.F32_sub(v91, v93)
						v99 = base.F32_add(base.F32_mul(v94, v94), v84)
					} else {
						v99 = v84
					}
					v112 = base.F64_promote_f32(v99)
				}
				v114 = F_Float8GetDatum(m, base.F64_sqrt(v112))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(16)
					return v114
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v127
						F_errmsg(m, int32(499600), v14)
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518027), int32(76), int32(159834))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_latin4_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(11), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic(m, v6, v5, v10, int32(132), int32(11), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_launcher_determine_sleep(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v70 = int64(*(*int32)(unsafe.Add(mBase, _consts[518])))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v4
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = v70
		v73 = v4
		v75 = v70
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[519]))
		if v15 == int32(0) {
			v70 = int64(*(*int32)(unsafe.Add(mBase, _consts[518])))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v4
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v70
			v73 = v4
			v75 = v70
		} else {
			if v15 == int32(4159680) {
				v70 = int64(*(*int32)(unsafe.Add(mBase, _consts[518])))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v4
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v70
				v73 = v4
				v75 = v70
			} else {
				v23 = m.G0
				v24 = int32(16)
				v25 = v23 - v24
				m.G0 = v25
				F___gettimeofday(m, v25)
				mBase = m.M
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
				m.G0 = v25 + v24
				v39 = *(*int32)(unsafe.Add(mBase, _consts[520]))
				v40 = int32(12)
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v39-v40)))
				v49 = v42 - (v29 + v28*int64(1000000) - int64(946684800000000))
				if v49 <= int64(0) {
					v61 = int32(0)
					v62 = int32(0)
				} else {
					v53 = int64(1000000)
					v54 = base.I64_div_u_s(v49, v53)
					v61 = base.I32_wrap_i64(v54)
					v62 = base.I32_wrap_i64(v49 - v54*v53)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v10+v40))) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v62
				v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+12)))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v65
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v67
				v73 = v67
				v75 = v65
			}
		}
	}
	if v75 == int64(0) {
		if l1 != 0 {
			if int32(100000) < v73 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(100000)
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
			}
			m.G0 = v10 + int32(16)
			return
		} else {
			if v73 != 0 {
				if int32(100000) < v73 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(100000)
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
				}
				m.G0 = v10 + int32(16)
				return
			} else {
				F_rebuild_database_list(m, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_launcher_determine_sleep(m, l0, int32(1), l2)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				}
			}
		}
	} else {
		if int64(0) < v75 {
			if base.Ui64(v75) < base.Ui64(int64(301)) {
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(300)
			}
		} else {
			if int32(100000) < v73 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(100000)
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
			}
		}
		m.G0 = v10 + int32(16)
		return
	}
}
func F_lca(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v8 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v11 = F_palloc(m, v8<<(uint(int32(2))%32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(0)
	goto L6
L4:
	;
	v43 = v15
	goto L5
L5:
	;
	v48 = F_lca_inner(m, v11, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v21<<(uint(int32(3))%32))))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v43 = v39
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+v21<<(uint(int32(2))%32)))) = v34
	v38 = v21 + int32(1)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v38 < v39 {
		v21 = v38
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v50 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = int32(0)
	v58 = v50
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_pfree(m, v11)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11+v57<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v57<<(uint(int32(3))%32))))
	if v66 != v70 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_pfree(m, v66)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v75 = v58
	goto L18
L18:
	;
	v77 = v57 + int32(1)
	if v77 < base.I32_extend16_s(v75) {
		v57 = v77
		v58 = v75
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v75 = v74
	goto L18
L20:
	;
	goto L15
L21:
	;
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = v48
	goto L24
L23:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
	v92 = int32(0)
	goto L24
L24:
	;
	return v92
}
func F_lcons_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v7 = F_palloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(4294967768)
			v16 = v7 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
			return v7
		}
	} else {
		F_new_head_cell(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
			return l1
		}
	}
}
func F_left_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+140)) = l2
	v18 = F_make_oper_cache_key(m, l0, v12+int32(4), l1, v6, l2, l4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(192)
	return v155
L2:
	;
	v66 = F_OpernameGetOprid(m, l1, int32(0), l2)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L18
	}
L3:
	;
	return int32(0)
L4:
	;
	if v18 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v49 = v25
	goto L8
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = int64(601295421576)
	v34 = F_hash_create(m, int32(417488), int32(256), v12+int32(144), int32(40))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v52 = int32(0)
	v54 = F_hash_search(m, v49, v12+int32(4), v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[301])) = v34
	F_CacheRegisterSyscacheCallback(m, int32(39), int32(491), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(12), int32(491), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v49 = v48
	goto L8
L12:
	;
	if v54 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+136))
	if v58 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v62 = F_SearchSysCache1(m, int32(40), v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v62 != 0 {
		v155 = v62
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L2
L17:
	;
	if l3 != 0 {
		v155 = int32(0)
		goto L1
	} else {
		goto L38
	}
L18:
	;
	if v66 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = F_OpernameGetCandidates(m, l1, int32(108), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	v112 = v66
	v117 = v6
	goto L21
L21:
	;
	v120 = F_SearchSysCache1(m, int32(40), v112)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L34
	}
L22:
	;
	if v72 == int32(0) {
		v142 = v6
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v78 = v72
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v87 != 0 {
		v78 = v87
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v72
	v94 = F_func_match_argtypes(m, int32(1), v12+int32(140), v72, v12+int32(144))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L30
	}
L26:
	;
	goto L25
L27:
	;
	v106 = int32(2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v107 == int32(0) {
		v142 = v106
		goto L17
	} else {
		goto L33
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	v104 = v103
	goto L27
L29:
	;
	v96 = int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	v101 = F_func_select_candidate(m, v96, v12+int32(140), v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	switch v94 {
	case 0:
		v142 = v94
		goto L17
	case 1:
		goto L28
	default:
		goto L29
	}
L31:
	;
	if v101 != 0 {
		v104 = v101
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v142 = v96
	goto L17
L33:
	;
	v112 = v107
	v117 = v106
	goto L21
L34:
	;
	if v120 == int32(0) {
		v142 = v117
		goto L17
	} else {
		goto L35
	}
L35:
	;
	if v18 == int32(0) {
		v155 = v120
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v132 = F_hash_search(m, v127, v12+int32(4), int32(1), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+136)) = v112
	v155 = v120
	goto L1
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	F_op_error(m, l0, l1, int32(0), v146, v142, l4)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_levenshtein_less_equal_with_costs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = v12 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v30 = v28 & int32(1)
			if v30 != 0 {
				v31 = v17
			} else {
				v31 = v12 + int32(4)
			}
			if v28 == int32(1) {
				v34 = int32(4)
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v36&int32(254) == int32(2) {
					v45 = v34
				} else {
					v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
				}
				if v36 == int32(1) {
					v48 = v34
				} else {
					v48 = v45
				}
				v59 = v48
			} else {
				v49 = int32(1)
				if v30 != 0 {
					v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v60 = int32(1)
			v61 = v19 + v60
			if v21&v60 != 0 {
				v66 = v61
			} else {
				v66 = v19 + int32(4)
			}
			if v21 == int32(1) {
				v69 = int32(4)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
				if v71&int32(254) == int32(2) {
					v80 = v69
				} else {
					v80 = base.B2i32(v71 == int32(18)) << (uint(v69) % 32)
				}
				if v71 == int32(1) {
					v83 = v69
				} else {
					v83 = v80
				}
				v96 = v83
			} else {
				v84 = int32(1)
				if v21&v84 != 0 {
					v96 = int32(base.Ui32(v21)>>(uint(v84)%32)) - v84
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v96 = int32(base.Ui32(v90)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v98 = F_varstr_levenshtein_less_equal(m, v31, v59, v66, v96, v25, v24, v23, v22, int32(0))
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				return v98
			}
		}
	}
}
func F_ln_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v497 float64
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int64
	_ = v528
	var v530 int64
	_ = v530
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L7
	} else {
		goto L172
	}
L4:
	;
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v19
	v35 = F_palloc(m, v13<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L7
	} else {
		goto L168
	}
L7:
	;
	return
L8:
	;
	v37 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v37 < v39 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = v39 << (uint(int32(1)) % 32)
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v35 + int32(2)
	v58 = F_palloc(m, int32(4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L16
	}
L12:
	;
	goto L11
L13:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v35+int32(2), v44, v46)
	mBase = m.M
	goto L15
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(131072)
	v63 = *(*int64)(unsafe.Add(mBase, _consts[1125]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v63
	v66 = *(*int64)(unsafe.Add(mBase, _consts[1126]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v58 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v58
	v73 = l2 + int32(8)
	v79 = int32(0)
	goto L17
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v84 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v600 = v11 + int32(72)
	F_sqrt_var(m, v600, v600, v73-v83<<(uint(int32(2))%32)>>(uint(int32(1))%32))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L7
	} else {
		goto L166
	}
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v87 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v90 = int32(1)
	v91 = int32(-1)
	v92 = int32(0)
	if base.B2i32(v91 < v83)&base.B2i32(v92 < v84) == v92 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if v263 <= int32(0) {
		goto L19
	} else {
		goto L69
	}
L23:
	;
	v263 = v253
	goto L22
L24:
	;
	if v91 <= v123 {
		v158 = v91
		v160 = v92
		goto L33
	} else {
		goto L34
	}
L25:
	;
	v123 = v83
	v127 = v92
	goto L24
L26:
	;
	goto L27
L27:
	;
	v104 = v83
	v108 = v92
	goto L28
L28:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+v108<<(uint(int32(1))%32)))))
	if v114 != 0 {
		v253 = int32(1)
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v123 = v118
	v127 = v116
	goto L24
L30:
	;
	v115 = int32(1)
	v116 = v108 + v115
	v118 = v104 - v115
	if v118 <= v91 {
		v123 = v118
		v127 = v116
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if v116 < v84 {
		v104 = v118
		v108 = v116
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v123 != v158 {
		v199 = v127
		v200 = v160
		goto L41
	} else {
		goto L42
	}
L34:
	;
	goto L35
L35:
	;
	v139 = v91
	v141 = v92
	goto L36
L36:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141<<(uint(int32(1))%32))+uint32(_consts[1127]))))
	if v146 != 0 {
		v253 = int32(-1)
		goto L23
	} else {
		goto L38
	}
L37:
	;
	v158 = v150
	v160 = v148
	goto L33
L38:
	;
	v147 = int32(1)
	v148 = v141 + v147
	v150 = v139 - v147
	if v150 <= v123 {
		v158 = v150
		v160 = v148
		goto L33
	} else {
		goto L39
	}
L39:
	;
	if v148 < v90 {
		v139 = v150
		v141 = v148
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	if v84 < v199 {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v169 = v127
	v170 = v160
	goto L43
L43:
	;
	if v84 <= v169 {
		v199 = v169
		v200 = v170
		goto L41
	} else {
		goto L45
	}
L44:
	;
	if base.I32_extend16_s(v185) < base.I32_extend16_s(v183) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v90 <= v170 {
		v199 = v169
		v200 = v170
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v174 = int32(1)
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+v169<<(uint(v174)%32)))))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170<<(uint(v174)%32))+uint32(_consts[1127]))))
	if v183 == v185 {
		v169 = v169 + v174
		v170 = v170 + v174
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v192 = int32(1)
	goto L50
L49:
	;
	v192 = int32(-1)
	goto L50
L50:
	;
	v263 = v192
	goto L22
L51:
	;
	v203 = v199
	goto L53
L52:
	;
	v203 = v84
	goto L53
L53:
	;
	v210 = v199
	goto L54
L54:
	;
	if v203 == v210 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v253 = v236
	goto L23
L56:
	;
	if v90 < v200 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v236 = int32(1)
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+v210<<(uint(v236)%32)))))
	if v242 == int32(0) {
		v210 = v210 + v236
		goto L54
	} else {
		goto L68
	}
L59:
	;
	v215 = v200
	goto L61
L60:
	;
	v215 = v90
	goto L61
L61:
	;
	v223 = v200
	goto L62
L62:
	;
	if v215 == v223 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v253 = int32(-1)
	goto L23
L64:
	;
	v263 = int32(0)
	goto L22
L65:
	;
	goto L66
L66:
	;
	v227 = int32(1)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223<<(uint(v227)%32))+uint32(_consts[1127]))))
	if v232 == int32(0) {
		v223 = v223 + v227
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	goto L55
L69:
	;
	v266 = v83
	v270 = v79
	v271 = v84
	goto L70
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v276 = int32(2)
	v277 = int32(0)
	if base.B2i32(v277 < v266)&base.B2i32(v277 < v271) == v277 {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	F_sub_var(m, v11+int32(72), int32(1770600), l1)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L7
	} else {
		goto L125
	}
L72:
	;
	goto L71
L73:
	;
	if v449 < int32(0) {
		v477 = v270
		goto L72
	} else {
		goto L120
	}
L74:
	;
	v449 = v439
	goto L73
L75:
	;
	if v277 <= v309 {
		v344 = v277
		v346 = v277
		goto L84
	} else {
		goto L85
	}
L76:
	;
	v309 = v266
	v313 = v277
	goto L75
L77:
	;
	goto L78
L78:
	;
	v290 = v266
	v294 = v277
	goto L79
L79:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+v294<<(uint(int32(1))%32)))))
	if v300 != 0 {
		v439 = int32(1)
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v309 = v304
	v313 = v302
	goto L75
L81:
	;
	v301 = int32(1)
	v302 = v294 + v301
	v304 = v290 - v301
	if v304 <= v277 {
		v309 = v304
		v313 = v302
		goto L75
	} else {
		goto L82
	}
L82:
	;
	if v302 < v271 {
		v290 = v304
		v294 = v302
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	if v309 != v344 {
		v385 = v313
		v386 = v346
		goto L92
	} else {
		goto L93
	}
L85:
	;
	goto L86
L86:
	;
	v325 = v277
	v327 = v277
	goto L87
L87:
	;
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327<<(uint(int32(1))%32))+uint32(_consts[1128]))))
	if v332 != 0 {
		v439 = int32(-1)
		goto L74
	} else {
		goto L89
	}
L88:
	;
	v344 = v336
	v346 = v334
	goto L84
L89:
	;
	v333 = int32(1)
	v334 = v327 + v333
	v336 = v325 - v333
	if v336 <= v309 {
		v344 = v336
		v346 = v334
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if v334 < v276 {
		v325 = v336
		v327 = v334
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	if v271 < v385 {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v355 = v313
	v356 = v346
	goto L94
L94:
	;
	if v271 <= v355 {
		v385 = v355
		v386 = v356
		goto L92
	} else {
		goto L96
	}
L95:
	;
	if base.I32_extend16_s(v371) < base.I32_extend16_s(v369) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	if v276 <= v356 {
		v385 = v355
		v386 = v356
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v360 = int32(1)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+v355<<(uint(v360)%32)))))
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v356<<(uint(v360)%32))+uint32(_consts[1128]))))
	if v369 == v371 {
		v355 = v355 + v360
		v356 = v356 + v360
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v378 = int32(1)
	goto L101
L100:
	;
	v378 = int32(-1)
	goto L101
L101:
	;
	v449 = v378
	goto L73
L102:
	;
	v389 = v385
	goto L104
L103:
	;
	v389 = v271
	goto L104
L104:
	;
	v396 = v385
	goto L105
L105:
	;
	if v389 == v396 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v439 = v422
	goto L74
L107:
	;
	if v276 < v386 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v422 = int32(1)
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+v396<<(uint(v422)%32)))))
	if v428 == int32(0) {
		v396 = v396 + v422
		goto L105
	} else {
		goto L119
	}
L110:
	;
	v401 = v386
	goto L112
L111:
	;
	v401 = v276
	goto L112
L112:
	;
	v409 = v386
	goto L113
L113:
	;
	if v401 == v409 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v439 = int32(-1)
	goto L74
L115:
	;
	v449 = int32(0)
	goto L73
L116:
	;
	goto L117
L117:
	;
	v413 = int32(1)
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409<<(uint(v413)%32))+uint32(_consts[1128]))))
	if v418 == int32(0) {
		v409 = v409 + v413
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	goto L106
L120:
	;
	v453 = v11 + int32(72)
	F_sqrt_var(m, v453, v453, v73-v266<<(uint(int32(2))%32)>>(uint(int32(1))%32))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	F_mul_var(m, v11, int32(1770648), v11, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	v468 = v270 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v469 == int32(0) {
		v477 = v468
		goto L72
	} else {
		goto L123
	}
L123:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v473 == int32(0) {
		v266 = v472
		v270 = v468
		v271 = v469
		goto L70
	} else {
		goto L124
	}
L124:
	;
	v477 = v468
	goto L72
L125:
	;
	F_add_var(m, v11+int32(72), int32(1770600), v11+int32(24))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v497 = base.F64_mul(base.F64_convert_i32_s(v477+int32(1)), float64(0.301029995663981))
	if base.F64_lt(base.F64_abs(v497), float64(2.147483648e+09)) != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v504 = v503 + v73
	F_div_var(m, l1, v11+int32(24), l1, v504, int32(1), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L7
	} else {
		goto L131
	}
L128:
	;
	v501 = base.I32_trunc_f64_s(v497)
	v503 = v501
	goto L127
L129:
	;
	goto L130
L130:
	;
	v503 = int32(-2147483648)
	goto L127
L131:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v514 = F_palloc(m, v509<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v516 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v514))) = uint16(v516)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v516 < v518 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v525 = v518 << (uint(int32(1)) % 32)
	if v525 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	goto L135
L135:
	;
	v528 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v528
	v530 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v514 + int32(2)
	F_mul_var(m, l1, l1, v11+int32(72), v504)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L7
	} else {
		goto L140
	}
L136:
	;
	goto L135
L137:
	;
	v526 = F__emscripten_memcpy_bulkmem(m, v514+int32(2), v523, v525)
	mBase = m.M
	goto L139
L138:
	;
	goto L139
L139:
	;
	goto L136
L140:
	;
	v540 = int32(1)
	v543 = base.I32_div_s(v504<<(uint(v540)%32), int32(-4))
	v549 = v540
	goto L141
L141:
	;
	v554 = v11 + int32(48)
	F_mul_var(m, v554, v11+int32(72), v554, v504)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L7
	} else {
		goto L144
	}
L142:
	;
	F_mul_var(m, l1, v11, l1, l2)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L149
	}
L143:
	;
	goto L142
L144:
	;
	v564 = v549 + int32(2)
	F_div_var_int(m, v11+int32(48), v564, int32(0), v11+int32(24), v504, int32(1))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v571 == int32(0) {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	F_add_var(m, l1, v11+int32(24), l1)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v579+v543 <= v578 {
		v549 = v564
		goto L141
	} else {
		goto L148
	}
L148:
	;
	goto L143
L149:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
	if v584 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_pfree(m, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L7
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v587 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	F_pfree(m, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L7
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v590 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L156
L158:
	;
	F_pfree(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L7
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v593 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L160
L162:
	;
	F_pfree(m, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L7
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	m.G0 = v11 + int32(96)
	return
L165:
	;
	goto L164
L166:
	;
	F_mul_var(m, v11, int32(1770648), v11, int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	v79 = v79 + int32(1)
	goto L17
L168:
	;
	F_errcode(m, int32(352583810))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L7
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(239569), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(523892), int32(11130), int32(240684))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L7
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_errcode(m, int32(352583810))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(251854), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(523892), int32(11126), int32(240684))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L7
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_load_relcache_init_file(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v523 int64
	_ = v523
	var v557 int32
	_ = v557
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v654 int32
	_ = v654
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v768 int32
	_ = v768
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(1104)
	m.G0 = v19
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = F_AllocateFile(m, v19+int32(80), int32(242112))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L10
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(106923)
	v29 = F_pg_snprintf(m, v19+int32(80), int32(1024), int32(187351), v19+int32(48))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(106923)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1177]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v36
	v44 = F_pg_snprintf(m, v19+int32(80), int32(1024), int32(187320), v19-int32(-64))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L1
L7:
	;
	goto L1
L8:
	;
	m.G0 = v19 + int32(1104)
	return v768
L9:
	;
	if int32(0) < v77 {
		goto L150
	} else {
		goto L151
	}
L10:
	;
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = F_palloc(m, int32(400))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v768 = int32(0)
	goto L8
L14:
	;
	v58 = F_fread(m, v19+int32(76), int32(1), int32(4), v49)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v621)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L5
	} else {
		goto L148
	}
L16:
	;
	if v58 != int32(4) {
		v621 = v52
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v62 != int32(5714534) {
		v621 = v52
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v74 = v52
	v75 = v2
	v76 = v2
	v77 = v2
	v78 = int32(100)
	goto L19
L19:
	;
	v86 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v75 == int32(4))&base.B2i32(v76 == int32(7)) != 0 {
		goto L9
	} else {
		goto L143
	}
L21:
	;
	goto L20
L22:
	;
	if v86 != int32(4) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v86 != 0 {
		v621 = v74
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v115 != int32(276) {
		v621 = v74
		goto L15
	} else {
		goto L33
	}
L26:
	;
	if l0 == int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	if base.B2i32(v75 == int32(5))&base.B2i32(v76 == int32(6)) != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v99 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	if v99 == int32(0) {
		v621 = v74
		goto L15
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(25769803781)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v75
	F_errmsg_internal(m, int32(20374), v19)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(522810), int32(6529), int32(404494))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v621 = v74
	goto L15
L33:
	;
	if v78 <= v77 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v122 = F_repalloc(m, v74, v78<<(uint(int32(3))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v127 = int32(276)
	v128 = v74
	v129 = v78
	goto L36
L36:
	;
	v133 = F_palloc(m, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v127 = v126
	v128 = v122
	v129 = v78 << (uint(int32(1)) % 32)
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128+v77<<(uint(int32(2))%32)))) = v133
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v138 = F_fread(m, v133, int32(1), v137, v49)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v138 != v137 {
		v621 = v128
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v145 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	if v145 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L42
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v150 = F_palloc(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v153 = F_fread(m, v150, int32(1), v149, v49)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v153 != v149 {
		v621 = v128
		goto L15
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+48)) = v150
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	v158 = F_CreateTemplateTupleDesc(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+52)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v164
	goto L49
L48:
	;
	v166 = int32(2249)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = int32(-1)
	v171 = int32(0)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v171 < v173 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v178 = v171
	v181 = v171
	goto L53
L51:
	;
	v237 = v171
	goto L52
L52:
	;
	v252 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L62
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v198 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L55
	}
L54:
	;
	v237 = v227
	goto L52
L55:
	;
	if v198 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v202 != int32(100) {
		v621 = v128
		goto L15
	} else {
		goto L57
	}
L57:
	;
	v208 = int32(100)
	v212 = v192 + v193<<(uint(int32(4))%32) + v178*v208 + int32(20)
	v215 = F_fread(m, v212, int32(1), v208, v49)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v215 != int32(100) {
		v621 = v128
		goto L15
	} else {
		goto L59
	}
L59:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+86)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	F_populate_compact_attribute(m, v220, v178)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v223 = int32(1)
	v227 = base.B2i32(v219|v181&v223 != int32(0))
	v229 = v178 + v223
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v229 < v230 {
		v178 = v229
		v181 = v227
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	if v252 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L63
	}
L63:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v256 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v237 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	v257 = F_palloc(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+180)) = int32(0)
	goto L64
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+180)) = v257
	v261 = F_fread(m, v257, int32(1), v256, v49)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	if v261 != v256 {
		v621 = v128
		goto L15
	} else {
		goto L70
	}
L70:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v133)+180))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v256 == int32(base.Ui32(v265)>>(uint(int32(2))%32)) {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	v621 = v128
	goto L15
L72:
	;
	v273 = F_palloc0(m, int32(20))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+25)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+119)))
	if v282 == int32(105) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+16)) = uint8(v275)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v273
	goto L74
L76:
	;
	v521 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+260)) = v521
	v523 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+128)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+68)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v133)+92)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+228)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v133)+176)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v133)+164)) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v133)+156)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v133)+136)) = v521
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+27)) = uint8(v521)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+76)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+100)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+108)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+116)) = v523
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+124)) = uint8(v521)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+236)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v133)+244)) = v521
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+25)))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v557
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+88)) = uint8(v521)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+28)) = uint8(v521)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+84)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v133)+256)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v133)+272)) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v133)+32)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+40)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v133)+144)) = v523
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+152)) = uint8(v521)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v133)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+60)) = v577
	v581 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+117)))
	if v583 != 0 {
		goto L139
	} else {
		goto L140
	}
L77:
	;
	v289 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v498 = v75 + v280
	switch v282 - int32(83) {
	case 0, 26, 31, 33:
		goto L136
	default:
		v512 = v498
		v513 = v76
		goto L76
	}
L80:
	;
	if v289 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v294 = F_palloc(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+196)) = v294
	v298 = F_fread(m, v294, int32(1), v293, v49)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v298 != v293 {
		v621 = v128
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v133)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v301)+16)) = v301 + int32(24)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v133)+196))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+16))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+192)) = v306 + v307
	v311 = *(*int32)(unsafe.Add(mBase, _consts[404]))
	v316 = F_AllocSetContextCreateInternal(m, v311, int32(253880), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+200)) = v316
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v322 = F_MemoryContextStrdup(m, v316, v319+int32(4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+36)) = v322
	goto L87
L87:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v133)+184))
	v326 = F_GetIndexAmRoutine(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v133)+200))
	v330 = F_MemoryContextAlloc(m, v328, int32(140))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	goto L91
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+204)) = v333
	F_pfree(m, v326)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L94
	}
L91:
	;
	v333 = F__emscripten_memcpy_bulkmem(m, v330, v326, int32(140))
	mBase = m.M
	goto L93
L93:
	;
	goto L90
L94:
	;
	v342 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	if v342 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L96
	}
L96:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v347 = F_MemoryContextAlloc(m, v316, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v350 = F_fread(m, v347, int32(1), v346, v49)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if v350 != v346 {
		v621 = v128
		goto L15
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+208)) = v347
	v358 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	if v358 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L101
	}
L101:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v363 = F_MemoryContextAlloc(m, v316, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v366 = F_fread(m, v363, int32(1), v362, v49)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	if v366 != v362 {
		v621 = v128
		goto L15
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+212)) = v363
	v374 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v374 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L106
	}
L106:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v379 = F_MemoryContextAlloc(m, v316, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v382 = F_fread(m, v379, int32(1), v378, v49)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	if v382 != v378 {
		v621 = v128
		goto L15
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+216)) = v379
	v390 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	if v390 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L111
	}
L111:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v395 = F_MemoryContextAlloc(m, v316, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v398 = F_fread(m, v395, int32(1), v394, v49)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	if v398 != v394 {
		v621 = v128
		goto L15
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+248)) = v395
	v406 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v406 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v411 = F_MemoryContextAlloc(m, v316, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v414 = F_fread(m, v411, int32(1), v410, v49)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if v414 != v410 {
		v621 = v128
		goto L15
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+224)) = v411
	v418 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	v421 = F_MemoryContextAllocZero(m, v316, v418<<(uint(int32(2))%32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+252)) = v421
	v424 = int32(0)
	v425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v424 < v425 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v433 = v424
	goto L124
L122:
	;
	v475 = v425
	goto L123
L123:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v133)+204))
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v490)+6)))
	v495 = F_MemoryContextAllocZero(m, v316, v475*v491*int32(28))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L135
	}
L124:
	;
	v448 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L126
	}
L125:
	;
	v475 = v471
	goto L123
L126:
	;
	if v448 != int32(4) {
		v621 = v128
		goto L15
	} else {
		goto L127
	}
L127:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v452 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v453 = F_MemoryContextAlloc(m, v316, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v470 = v433 + int32(1)
	v471 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v470 < v471 {
		v433 = v470
		goto L124
	} else {
		goto L134
	}
L131:
	;
	v456 = v433 << (uint(int32(2)) % 32)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v133)+252))
	*(*int32)(unsafe.Add(mBase, uint32(v456+v457))) = v453
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v133)+252))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v460+v456)))
	v464 = F_fread(m, v462, int32(1), v452, v49)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	if v464 != v452 {
		v621 = v128
		goto L15
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	goto L125
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+220)) = v495
	v512 = v75
	v513 = v76 + v280
	goto L76
L136:
	;
	F_RelationInitTableAccessMethod(m, v133)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	v512 = v498
	v513 = v76
	goto L76
L138:
	;
	F_RelationInitPhysicalAddr(m, v133)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L142
	}
L139:
	;
	v584 = v521
	goto L141
L140:
	;
	v584 = v581
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+64)) = v584
	goto L138
L142:
	;
	v74 = v128
	v75 = v512
	v76 = v513
	v77 = v77 + int32(1)
	v78 = v129
	goto L19
L143:
	;
	v595 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	if v595 == int32(0) {
		v621 = v74
		goto L15
	} else {
		goto L145
	}
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = int64(30064771076)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v75
	F_errmsg_internal(m, int32(20481), v19+int32(32))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L5
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(522810), int32(6543), int32(404494))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L5
	} else {
		goto L147
	}
L147:
	;
	v621 = v74
	goto L15
L148:
	;
	v631 = F_FreeFile(m, v49)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L149
	}
L149:
	;
	goto L13
L150:
	;
	v654 = int32(0)
	goto L153
L151:
	;
	goto L152
L152:
	;
	F_pfree(m, v74)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L5
	} else {
		goto L170
	}
L153:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _consts[1174]))
	v673 = v74 + v654<<(uint(int32(2))%32)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	v680 = F_hash_search(m, v670, v674+int32(56), int32(1), v19+int32(72))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L155
	}
L154:
	;
	goto L152
L155:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	if v682 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v722 = v654 + int32(1)
	if v722 != v77 {
		v654 = v722
		goto L153
	} else {
		goto L169
	}
L157:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+4)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	if v688 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+4)) = v718
	goto L156
L160:
	;
	F_RelationDestroyRelation(m, v685, int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	if v695 == int32(0) {
		goto L156
	} else {
		goto L164
	}
L163:
	;
	goto L156
L164:
	;
	v700 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if v700 == int32(0) {
		goto L156
	} else {
		goto L166
	}
L166:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v685)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v704 + int32(4)
	F_errmsg_internal(m, int32(732504), v19+int32(16))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(522810), int32(6556), int32(404494))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	goto L156
L169:
	;
	goto L154
L170:
	;
	v742 = F_FreeFile(m, v49)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L5
	} else {
		goto L171
	}
L171:
	;
	if l0 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v745 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1178])) = uint8(v745)
	v768 = v745
	goto L8
L173:
	;
	goto L174
L174:
	;
	v749 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1179])) = uint8(v749)
	v768 = v749
	goto L8
}
func F_locate_var_of_level_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = int32(0)
	if l0 == v3 {
		v40 = v3
		return v40
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(58) {
		case 0:
			v40 = v3
			return v40
		case 1, 2, 3, 4, 5, 6, 7, 8:
			v37 = F_expression_tree_walker_impl(m, l0, int32(902), l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v40 = v37
				return v40
			}
		case 9:
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v21 + int32(1)
			v27 = F_query_tree_walker_impl(m, l0, int32(902), l1, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v31 - int32(1)
				return v27
			}
		default:
			if v7 != int32(6) {
				v37 = F_expression_tree_walker_impl(m, l0, int32(902), l1)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = v37
					return v40
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v12 != v13 {
					v40 = v3
					return v40
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v15 < int32(0) {
						v40 = v3
						return v40
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
						return int32(1)
					}
				}
			}
		}
	}
}
func F_locate_windowfunc_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(11) {
			v19 = F_expression_tree_walker_impl(m, l0, int32(1046), l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v11 < int32(0) {
				v19 = F_expression_tree_walker_impl(m, l0, int32(1046), l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
				return int32(1)
			}
		}
	}
}
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = F_lookup_rowtype_tupdesc_internal(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		if int32(0) <= v7 {
			F_IncrTupleDescRefCount(m, v3)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v3
			}
		} else {
			return v3
		}
	}
}
func F_ltsGetBlock(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v70 int64
	_ = v70
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v85 int64
	_ = v85
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v192 int64
	_ = v192
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v209 int64
	_ = v209
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v245 int64
	_ = v245
	var v250 int64
	_ = v250
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v12 == int32(1) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		if int32(0) < v16 {
			v168 = v16
			v169 = v15
			v178 = v168 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v178
			v183 = *(*int64)(unsafe.Add(mBase, uint32(v169+v178<<(uint(int32(3))%32))))
			return v183
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = int32(8)
				v24 = F_palloc(m, int32(64))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v44 = v24
					*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v46
					if v46 <= int32(0) {
						v168 = v46
						v169 = v44
					} else {
						v50 = v46
						v54 = v50
						for {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							if base.Ui64(v64) <= base.Ui64(int64(1)) {
								if base.I32_wrap_i64(v64) != int32(1) {
									v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v70 + int64(1)
									v151 = v70
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
									v151 = v76
								}
							} else {
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								v79 = v64 - int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v79
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v79)<<(uint(int32(3))%32))))
								v95 = int64(0)
								for {
									v98 = int64(1)
									v99 = v95 << (uint(v98) % 64)
									v101 = v99 + int64(2)
									v103 = v99 | v98
									if base.Ui64(v79) <= base.Ui64(v103) {
										v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
										if base.Ui64(v103) < base.Ui64(v79) {
											v119 = v103
										} else {
											v119 = v101
										}
										if base.Ui64(v103) < base.Ui64(v79) {
											v122 = v119
											v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
											if v85 <= v127 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
												v95 = v122
												continue
											}
											break
										} else {
											if base.Ui64(v79) <= base.Ui64(v101) {
												break
											} else {
												v122 = v119
												v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
												if v85 <= v127 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
													v95 = v122
													continue
												}
												break
											}
											break
										}
										break
									} else {
										if base.Ui64(v79) <= base.Ui64(v101) {
											v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
											if base.Ui64(v103) < base.Ui64(v79) {
												v119 = v103
											} else {
												v119 = v101
											}
											if base.Ui64(v103) < base.Ui64(v79) {
												v122 = v119
												v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
												if v85 <= v127 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
													v95 = v122
													continue
												}
												break
											} else {
												if base.Ui64(v79) <= base.Ui64(v101) {
													break
												} else {
													v122 = v119
													v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
													if v85 <= v127 {
														break
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
														v95 = v122
														continue
													}
													break
												}
												break
											}
											break
										} else {
											v107 = int32(3)
											v110 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v103)<<(uint(v107)%32))))
											v115 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v101)<<(uint(v107)%32))))
											if v110 < v115 {
												v117 = v103
											} else {
												v117 = v101
											}
											v122 = v117
											v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
											if v85 <= v127 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
												v95 = v122
												continue
											}
											break
										}
										break
									}
									break
								}
								*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v85
								v151 = v77
							}
							v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
							*(*int64)(unsafe.Add(mBase, uint32(v153+v54<<(uint(int32(3))%32)-int32(8)))) = v151
							v160 = int32(1)
							if v160 < v54 {
								v54 = v54 - v160
								continue
							} else {
								break
							}
							break
						}
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
						v168 = v164
						v169 = v165
					}
					v178 = v168 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v178
					v183 = *(*int64)(unsafe.Add(mBase, uint32(v169+v178<<(uint(int32(3))%32))))
					return v183
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if int32(128) <= v28 {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v28
					v50 = v28
					v54 = v50
					for {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						if base.Ui64(v64) <= base.Ui64(int64(1)) {
							if base.I32_wrap_i64(v64) != int32(1) {
								v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v70 + int64(1)
								v151 = v70
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								v151 = v76
							}
						} else {
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
							v79 = v64 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v79
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v79)<<(uint(int32(3))%32))))
							v95 = int64(0)
							for {
								v98 = int64(1)
								v99 = v95 << (uint(v98) % 64)
								v101 = v99 + int64(2)
								v103 = v99 | v98
								if base.Ui64(v79) <= base.Ui64(v103) {
									v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
									if base.Ui64(v103) < base.Ui64(v79) {
										v119 = v103
									} else {
										v119 = v101
									}
									if base.Ui64(v103) < base.Ui64(v79) {
										v122 = v119
										v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
										if v85 <= v127 {
											break
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
											v95 = v122
											continue
										}
										break
									} else {
										if base.Ui64(v79) <= base.Ui64(v101) {
											break
										} else {
											v122 = v119
											v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
											if v85 <= v127 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
												v95 = v122
												continue
											}
											break
										}
										break
									}
									break
								} else {
									if base.Ui64(v79) <= base.Ui64(v101) {
										v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
										if base.Ui64(v103) < base.Ui64(v79) {
											v119 = v103
										} else {
											v119 = v101
										}
										if base.Ui64(v103) < base.Ui64(v79) {
											v122 = v119
											v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
											if v85 <= v127 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
												v95 = v122
												continue
											}
											break
										} else {
											if base.Ui64(v79) <= base.Ui64(v101) {
												break
											} else {
												v122 = v119
												v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
												if v85 <= v127 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
													v95 = v122
													continue
												}
												break
											}
											break
										}
										break
									} else {
										v107 = int32(3)
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v103)<<(uint(v107)%32))))
										v115 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v101)<<(uint(v107)%32))))
										if v110 < v115 {
											v117 = v103
										} else {
											v117 = v101
										}
										v122 = v117
										v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
										if v85 <= v127 {
											break
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
											v95 = v122
											continue
										}
										break
									}
									break
								}
								break
							}
							*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v85
							v151 = v77
						}
						v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
						*(*int64)(unsafe.Add(mBase, uint32(v153+v54<<(uint(int32(3))%32)-int32(8)))) = v151
						v160 = int32(1)
						if v160 < v54 {
							v54 = v54 - v160
							continue
						} else {
							break
						}
						break
					}
					v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
					v168 = v164
					v169 = v165
					v178 = v168 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v178
					v183 = *(*int64)(unsafe.Add(mBase, uint32(v169+v178<<(uint(int32(3))%32))))
					return v183
				} else {
					v32 = int32(128)
					v34 = v28 << (uint(int32(1)) % 32)
					if v32 <= v34 {
						v37 = v32
					} else {
						v37 = v34
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v37
					v41 = F_repalloc(m, v15, v37<<(uint(int32(3))%32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v44 = v41
						*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v46
						if v46 <= int32(0) {
							v168 = v46
							v169 = v44
						} else {
							v50 = v46
							v54 = v50
							for {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								if base.Ui64(v64) <= base.Ui64(int64(1)) {
									if base.I32_wrap_i64(v64) != int32(1) {
										v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v70 + int64(1)
										v151 = v70
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
										v76 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
										v151 = v76
									}
								} else {
									v77 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
									v79 = v64 - int64(1)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v79
									v85 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v79)<<(uint(int32(3))%32))))
									v95 = int64(0)
									for {
										v98 = int64(1)
										v99 = v95 << (uint(v98) % 64)
										v101 = v99 + int64(2)
										v103 = v99 | v98
										if base.Ui64(v79) <= base.Ui64(v103) {
											v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
											if base.Ui64(v103) < base.Ui64(v79) {
												v119 = v103
											} else {
												v119 = v101
											}
											if base.Ui64(v103) < base.Ui64(v79) {
												v122 = v119
												v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
												if v85 <= v127 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
													v95 = v122
													continue
												}
												break
											} else {
												if base.Ui64(v79) <= base.Ui64(v101) {
													break
												} else {
													v122 = v119
													v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
													if v85 <= v127 {
														break
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
														v95 = v122
														continue
													}
													break
												}
												break
											}
											break
										} else {
											if base.Ui64(v79) <= base.Ui64(v101) {
												v118 = base.B2i32(base.Ui64(v103) < base.Ui64(v79))
												if base.Ui64(v103) < base.Ui64(v79) {
													v119 = v103
												} else {
													v119 = v101
												}
												if base.Ui64(v103) < base.Ui64(v79) {
													v122 = v119
													v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
													if v85 <= v127 {
														break
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
														v95 = v122
														continue
													}
													break
												} else {
													if base.Ui64(v79) <= base.Ui64(v101) {
														break
													} else {
														v122 = v119
														v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
														if v85 <= v127 {
															break
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
															v95 = v122
															continue
														}
														break
													}
													break
												}
												break
											} else {
												v107 = int32(3)
												v110 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v103)<<(uint(v107)%32))))
												v115 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v101)<<(uint(v107)%32))))
												if v110 < v115 {
													v117 = v103
												} else {
													v117 = v101
												}
												v122 = v117
												v127 = *(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v122)<<(uint(int32(3))%32))))
												if v85 <= v127 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v127
													v95 = v122
													continue
												}
												break
											}
											break
										}
										break
									}
									*(*int64)(unsafe.Add(mBase, uint32(v63+base.I32_wrap_i64(v95)<<(uint(int32(3))%32)))) = v85
									v151 = v77
								}
								v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
								*(*int64)(unsafe.Add(mBase, uint32(v153+v54<<(uint(int32(3))%32)-int32(8)))) = v151
								v160 = int32(1)
								if v160 < v54 {
									v54 = v54 - v160
									continue
								} else {
									break
								}
								break
							}
							v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
							v168 = v164
							v169 = v165
						}
						v178 = v168 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v178
						v183 = *(*int64)(unsafe.Add(mBase, uint32(v169+v178<<(uint(int32(3))%32))))
						return v183
					}
				}
			}
		}
	} else {
		v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v186 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
		if base.Ui64(v186) <= base.Ui64(int64(1)) {
			if base.I32_wrap_i64(v186) != int32(1) {
				v192 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v192 + int64(1)
				return v192
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
				v199 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
				return v199
			}
		} else {
			v201 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
			v203 = v186 - int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v203
			v209 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v203)<<(uint(int32(3))%32))))
			v218 = int64(0)
			for {
				v221 = int64(1)
				v222 = v218 << (uint(v221) % 64)
				v224 = v222 + int64(2)
				v226 = v222 | v221
				if base.Ui64(v203) <= base.Ui64(v226) {
					v241 = base.B2i32(base.Ui64(v226) < base.Ui64(v203))
					if base.Ui64(v226) < base.Ui64(v203) {
						v242 = v226
					} else {
						v242 = v224
					}
					if base.Ui64(v226) < base.Ui64(v203) {
						v245 = v242
						v250 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v245)<<(uint(int32(3))%32))))
						if v209 <= v250 {
							break
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v250
							v218 = v245
							continue
						}
						break
					} else {
						if base.Ui64(v203) <= base.Ui64(v224) {
							break
						} else {
							v245 = v242
							v250 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v245)<<(uint(int32(3))%32))))
							if v209 <= v250 {
								break
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v250
								v218 = v245
								continue
							}
							break
						}
						break
					}
					break
				} else {
					if base.Ui64(v203) <= base.Ui64(v224) {
						v241 = base.B2i32(base.Ui64(v226) < base.Ui64(v203))
						if base.Ui64(v226) < base.Ui64(v203) {
							v242 = v226
						} else {
							v242 = v224
						}
						if base.Ui64(v226) < base.Ui64(v203) {
							v245 = v242
							v250 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v245)<<(uint(int32(3))%32))))
							if v209 <= v250 {
								break
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v250
								v218 = v245
								continue
							}
							break
						} else {
							if base.Ui64(v203) <= base.Ui64(v224) {
								break
							} else {
								v245 = v242
								v250 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v245)<<(uint(int32(3))%32))))
								if v209 <= v250 {
									break
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v250
									v218 = v245
									continue
								}
								break
							}
							break
						}
						break
					} else {
						v230 = int32(3)
						v233 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v226)<<(uint(v230)%32))))
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v224)<<(uint(v230)%32))))
						if v233 < v238 {
							v240 = v226
						} else {
							v240 = v224
						}
						v245 = v240
						v250 = *(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v245)<<(uint(int32(3))%32))))
						if v209 <= v250 {
							break
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v250
							v218 = v245
							continue
						}
						break
					}
					break
				}
				break
			}
			*(*int64)(unsafe.Add(mBase, uint32(v185+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v209
			return v201
		}
	}
}
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v92 int64
	_ = v92
	var v103 int32
	_ = v103
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = int32(0)
	v24 = v17
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	if v24 == int64(-1) {
		v151 = v20
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return base.B2i32(int32(0) < v151)
L4:
	;
	goto L3
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v34 = v33 + v24
	v35 = F_BufFileSeekBlock(m, v32, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = v30 + v20
	F_BufFileReadExact(m, v39, v40, int32(8192))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1244])))
	if int64(0) <= v131 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+40)))
	if v46 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v47 < base.I64_extend_i32_u(v48) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+48)) = v66 + int64(1)
	if v66 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v64 = v51
	v66 = v47
	goto L13
L15:
	;
	goto L16
L16:
	;
	v53 = v48 << (uint(int32(4)) % 32)
	if base.Ui32(int32(1073741823)) < base.Ui32(v53) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+56)) = v48 << (uint(int32(1)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v60 = F_repalloc(m, v59, v53)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+44)) = v60
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v64 = v60
	v66 = v63
	goto L13
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64+v103<<(uint(int32(3))%32)))) = v34
	goto L10
L20:
	;
	v103 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v79 = v66
	goto L23
L23:
	;
	v83 = base.I32_wrap_i64(v79)
	v84 = int64(1)
	v85 = v79 - v84
	v87 = int64(base.Ui64(v85) >> (uint(v84) % 64))
	v88 = base.I32_wrap_i64(v87)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v64+v88<<(uint(int32(3))%32))))
	if v92 < v34 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v103 = v88
	goto L19
L25:
	;
	v103 = v83
	goto L19
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64+v83<<(uint(int32(3))%32)))) = v92
	if base.Ui64(int64(1)) < base.Ui64(v85) {
		v79 = v87
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v136 = int32(8176)
	goto L31
L30:
	;
	v136 = int32(0) - base.I32_wrap_i64(v131)
	goto L31
L31:
	;
	v137 = v126 + v136
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1244])))
	if v139 < int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
	v151 = v137
	goto L4
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v139
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(8192) < v145-v137 {
		v20 = v137
		v24 = v139
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v151 = v137
	goto L4
L36:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v34
	F_errmsg(m, int32(405080), v13)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(522391), int32(288), int32(331999))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
