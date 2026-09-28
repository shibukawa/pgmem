package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LookupNamespaceNoError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v2 = int32(_a_F_LookupNamespaceNoError_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LookupNamespaceNoError[0])))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v57
L2:
	;
	if v26-v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v11 = l0
	v12 = v2
	goto L5
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v26 = v16
	v27 = v15
	goto L3
L7:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_LookupNamespaceNoError[1]))
	if v33 == v31 {
		v57 = v31
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = int64(0)
	v54 = F_GetSysCacheOid(m, int32(37), base.I64_extend_i32_u(l0), v51, v51, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L18
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_LookupNamespaceNoError[2]))
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return v33
L14:
	;
	goto L15
L15:
	;
	v42 = F_RunNamespaceSearchHook(m, v33, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_LookupNamespaceNoError[1]))
	return v47
L18:
	;
	v57 = v54
	goto L1
}
func F_get_namespace_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14286(m, l0, int32(38))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_recomputeNamespacePath(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v493 int64
	_ = v493
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	v1 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[0]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])))
	if v19 == v1 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L7
	} else {
		goto L141
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2]))
	if v23 == v17 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[3]))
	F_spcache_init(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	v29 = F_spcache_insert(m, v26, v17)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = int32(_a_F_recomputeNamespacePath_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v38
	v40 = F_pstrdup(m, v26)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v219 == v218 {
		goto L64
	} else {
		goto L65
	}
L13:
	;
	v45 = F_SplitIdentifierString(m, v40, int32(44), v14+int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)) = uint8(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v51 == v49 {
		v191 = v1
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_pfree(m, v40)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L61
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 <= int32(0) {
		v191 = v1
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v61 = v1
	v63 = v1
	goto L19
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63<<(uint(int32(2))%32))))
	v74 = int32(_a_F_recomputeNamespacePath_1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[6])))
	if base.B2i32(v77 == int32(0))|base.B2i32(v77 != v80) != 0 {
		v98 = v77
		v99 = v80
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v191 = v182
	goto L16
L21:
	;
	v185 = v63 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v185 < v186 {
		v61 = v182
		v63 = v185
		goto L19
	} else {
		goto L60
	}
L22:
	;
	if v98-v99 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v83 = v73
	v84 = v74
	goto L25
L25:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v88
		v99 = v87
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v98 = v88
	v99 = v87
	goto L23
L27:
	;
	v91 = int32(1)
	if v88 == v87 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v104 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(v17))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v130 = int32(_a_F_recomputeNamespacePath_2)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[7])))
	if base.B2i32(v133 == int32(0))|base.B2i32(v133 != v136) != 0 {
		v154 = v133
		v155 = v136
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v104 == int32(0) {
		v182 = v61
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
	v115 = int64(0)
	v118 = F_GetSysCacheOid(m, int32(37), base.I64_extend_i32_u(v109+v110+int32(4)), v115, v115, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_ReleaseCatCache(m, v104)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v118 == int32(0) {
		v182 = v61
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v126 = F_object_aclcheck(m, int32(2615), v118, v17, int64(256))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	if v126 != 0 {
		v182 = v61
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v128 = F_lappend_oid(m, v61, v118)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v182 = v128
	goto L21
L40:
	;
	if v154-v155 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v139 = v73
	v140 = v130
	goto L43
L43:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	if v144 == int32(0) {
		v154 = v144
		v155 = v143
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v154 = v144
	v155 = v143
	goto L41
L45:
	;
	v147 = int32(1)
	if v144 == v143 {
		v139 = v139 + v147
		v140 = v140 + v147
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	if v160 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v168 = int64(0)
	v171 = F_GetSysCacheOid(m, int32(37), base.I64_extend_i32_u(v73), v168, v168, v168)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L55
	}
L50:
	;
	v161 = F_lappend_oid(m, v61, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v61 != 0 {
		v182 = v61
		goto L21
	} else {
		goto L54
	}
L53:
	;
	v182 = v161
	goto L21
L54:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)) = uint8(v163)
	v182 = int32(0)
	goto L21
L55:
	;
	if v171 == int32(0) {
		v182 = v61
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v177 = F_object_aclcheck(m, int32(2615), v171, v17, int64(256))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	if v177 != 0 {
		v182 = v61
		goto L21
	} else {
		goto L58
	}
L58:
	;
	v179 = F_lappend_oid(m, v61, v171)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v182 = v179
	goto L21
L60:
	;
	goto L20
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_list_free(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v191
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v35
	goto L12
L63:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9]))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v448 != v449 {
		v460 = v437
		goto L134
	} else {
		goto L135
	}
L64:
	;
	F_list_free(m, v219)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L68
	}
L65:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	if v223 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	if v224 == int32(1) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v437 = v219
	goto L63
L68:
	;
	v229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v229
	v231 = int32(_a_F_recomputeNamespacePath_0)
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v235
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v238 == v229 {
		v322 = v218
		v332 = v229
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v332
	v335 = int32(0)
	if v322 == v335 {
		goto L100
	} else {
		goto L101
	}
L70:
	;
	v241 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v242 <= v241 {
		v322 = v218
		v332 = v241
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v247 = v218
	v251 = int32(0)
	goto L72
L72:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v251<<(uint(int32(2))%32))))
	v262 = int32(0)
	if v247 == v262 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v310 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L74:
	;
	v312 = v251 + int32(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v312 < v313 {
		v247 = v310
		v251 = v312
		goto L72
	} else {
		goto L95
	}
L75:
	;
	if v300 != 0 {
		v310 = v247
		goto L74
	} else {
		goto L88
	}
L76:
	;
	v300 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v268 <= int32(0) {
		v294 = v262
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v300 = v294
	goto L75
L80:
	;
	v271 = int32(0)
	if v271 < v268 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v274 = v268
	goto L83
L82:
	;
	v274 = v271
	goto L83
L83:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v277 = int32(0)
	goto L84
L84:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275+v277<<(uint(int32(2))%32))))
	v286 = base.B2i32(v285 == v261)
	if v285 == v261 {
		v294 = v286
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v294 = v286
	goto L79
L86:
	;
	v288 = v277 + int32(1)
	if v288 != v274 {
		v277 = v288
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	if v302 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v304 = F_RunNamespaceSearchHook(m, v261, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v308 = F_lappend_oid(m, v247, v261)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	if v304 == int32(0) {
		v310 = v247
		goto L74
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v310 = v308
	goto L74
L95:
	;
	goto L73
L96:
	;
	v317 = int32(0)
	v322 = v317
	v332 = v317
	goto L69
L97:
	;
	goto L98
L98:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v322 = v310
	v332 = v320
	goto L69
L99:
	;
	if v373 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L100:
	;
	v373 = int32(0)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v341 <= int32(0) {
		v367 = v335
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v373 = v367
	goto L99
L104:
	;
	v344 = int32(0)
	if v344 < v341 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v347 = v341
	goto L107
L106:
	;
	v347 = v344
	goto L107
L107:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v350 = int32(0)
	goto L108
L108:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v348+v350<<(uint(int32(2))%32))))
	v359 = base.B2i32(v358 == int32(11))
	if v358 == int32(11) {
		v367 = v359
		goto L103
	} else {
		goto L110
	}
L109:
	;
	v367 = v359
	goto L103
L110:
	;
	v361 = v350 + int32(1)
	if v361 != v347 {
		v350 = v361
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v377 = F_lcons_oid(m, int32(11), v322)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L115
	}
L113:
	;
	v379 = v322
	goto L114
L114:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	if v381 == int32(0) {
		v427 = v379
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v379 = v377
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v427
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v232
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)) = uint8(base.B2i32(v432 != int32(0)))
	v437 = v427
	goto L63
L117:
	;
	v384 = int32(0)
	if v379 == v384 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v422 != 0 {
		v427 = v379
		goto L116
	} else {
		goto L131
	}
L119:
	;
	v422 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v390 <= int32(0) {
		v416 = v384
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v422 = v416
	goto L118
L123:
	;
	v393 = int32(0)
	if v393 < v390 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v396 = v390
	goto L126
L125:
	;
	v396 = v393
	goto L126
L126:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v399 = int32(0)
	goto L127
L127:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v397+v399<<(uint(int32(2))%32))))
	v408 = base.B2i32(v407 == v381)
	if v407 == v381 {
		v416 = v408
		goto L122
	} else {
		goto L129
	}
L128:
	;
	v416 = v408
	goto L122
L129:
	;
	v410 = v399 + int32(1)
	if v410 != v396 {
		v399 = v410
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	v425 = F_lcons_oid(m, v424, v379)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v427 = v425
	goto L116
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2])) = v17
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[12])) = v504
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[13])) = v508
	v512 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])))
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[15])) = uint8(v512)
	v515 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])) = uint8(v515)
	goto L2
L134:
	;
	v461 = int32(_a_F_recomputeNamespacePath_0)
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v465
	v467 = F_list_copy(m, v460)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L7
	} else {
		goto L139
	}
L135:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)))
	if v452 != v453 {
		v460 = v437
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	v457 = F_equal(m, v437, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	if v457 != 0 {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v460 = v459
	goto L134
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v462
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	F_list_free(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11])) = v467
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9])) = v478
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)))
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])) = uint8(v481)
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2])) = v17
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[12])) = v467
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[13])) = v478
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[15])) = uint8(v481)
	v491 = int32(_a_F_recomputeNamespacePath_3)
	v493 = *(*int64)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[17])) = v493 + int64(1)
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])) = uint8(v498)
	goto L2
L141:
	;
	F_errmsg_internal(m, int32(_a_F_recomputeNamespacePath_4), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_recomputeNamespacePath_5), int32(_a_F_recomputeNamespacePath_6), int32(_a_F_recomputeNamespacePath_7))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
