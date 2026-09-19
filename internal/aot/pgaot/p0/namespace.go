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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	return v56
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
		v56 = v31
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v50 = int32(0)
	v53 = F_GetSysCacheOid(m, int32(37), l0, v50, v50, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
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
	v56 = v53
	goto L1
}
func F_get_namespace_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13917(m, l0, int32(38))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v481 int32
	_ = v481
	var v483 int64
	_ = v483
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	v1 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[0]))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])))
	if v18 == v1 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L7
	} else {
		goto L141
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2]))
	if v22 == v16 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[3]))
	F_spcache_init(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v28 = F_spcache_insert(m, v25, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = int32(_a_F_recomputeNamespacePath_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v37
	v39 = F_pstrdup(m, v25)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v211 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v212 == v211 {
		goto L64
	} else {
		goto L65
	}
L13:
	;
	v44 = F_SplitIdentifierString(m, v39, int32(44), v13+int32(12))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v50 == v48 {
		v186 = v1
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_pfree(m, v39)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L61
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 <= int32(0) {
		v186 = v1
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = v1
	v61 = v1
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v61<<(uint(int32(2))%32))))
	v71 = int32(_a_F_recomputeNamespacePath_1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[6])))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v186 = v177
	goto L16
L21:
	;
	v180 = v61 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v180 < v181 {
		v59 = v177
		v61 = v180
		goto L19
	} else {
		goto L60
	}
L22:
	;
	if v95-v96 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v80 = v70
	v81 = v71
	goto L25
L25:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v95 = v85
	v96 = v84
	goto L23
L27:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v101 = F_SearchSysCache1(m, int32(11), v16)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v126 = int32(_a_F_recomputeNamespacePath_2)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[7])))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v101 == int32(0) {
		v177 = v59
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
	v111 = int32(0)
	v114 = F_GetSysCacheOid(m, int32(37), v106+v107+int32(4), v111, v111, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_ReleaseCatCache(m, v101)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v114 == int32(0) {
		v177 = v59
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v122 = F_object_aclcheck(m, int32(2615), v114, v16, int64(256))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	if v122 != 0 {
		v177 = v59
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v124 = F_lappend_oid(m, v59, v114)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v177 = v124
	goto L21
L40:
	;
	if v150-v151 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v135 = v70
	v136 = v126
	goto L43
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v150 = v140
	v151 = v139
	goto L41
L45:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	if v156 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v163 = int32(0)
	v166 = F_GetSysCacheOid(m, int32(37), v70, v163, v163, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L55
	}
L50:
	;
	v157 = F_lappend_oid(m, v59, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v59 != 0 {
		v177 = v59
		goto L21
	} else {
		goto L54
	}
L53:
	;
	v177 = v157
	goto L21
L54:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)) = uint8(v159)
	v177 = int32(0)
	goto L21
L55:
	;
	if v166 == int32(0) {
		v177 = v59
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v172 = F_object_aclcheck(m, int32(2615), v166, v16, int64(256))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	if v172 != 0 {
		v177 = v59
		goto L21
	} else {
		goto L58
	}
L58:
	;
	v174 = F_lappend_oid(m, v59, v166)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v177 = v174
	goto L21
L60:
	;
	goto L20
L61:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_list_free(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v186
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v34
	goto L12
L63:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9]))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v438 != v439 {
		v450 = v428
		goto L134
	} else {
		goto L135
	}
L64:
	;
	F_list_free(m, v212)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L68
	}
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	if v216 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)))
	if v217 == int32(1) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v428 = v212
	goto L63
L68:
	;
	v222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v222
	v224 = int32(_a_F_recomputeNamespacePath_0)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v231 == v222 {
		v314 = v211
		v323 = v222
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v323
	v326 = int32(0)
	if v314 == v326 {
		goto L100
	} else {
		goto L101
	}
L70:
	;
	v234 = int32(0)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v235 <= v234 {
		v314 = v211
		v323 = v234
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v240 = v211
	v244 = int32(0)
	goto L72
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+v244<<(uint(int32(2))%32))))
	v254 = int32(0)
	if v240 == v254 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v302 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L74:
	;
	v304 = v244 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v304 < v305 {
		v240 = v302
		v244 = v304
		goto L72
	} else {
		goto L95
	}
L75:
	;
	if v292 != 0 {
		v302 = v240
		goto L74
	} else {
		goto L88
	}
L76:
	;
	v292 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v260 <= int32(0) {
		v286 = v254
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v292 = v286
	goto L75
L80:
	;
	v263 = int32(0)
	if v263 < v260 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v266 = v260
	goto L83
L82:
	;
	v266 = v263
	goto L83
L83:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v269 = int32(0)
	goto L84
L84:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v267+v269<<(uint(int32(2))%32))))
	v278 = base.B2i32(v277 == v253)
	if v277 == v253 {
		v286 = v278
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v286 = v278
	goto L79
L86:
	;
	v280 = v269 + int32(1)
	if v280 != v266 {
		v269 = v280
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	if v294 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v296 = F_RunNamespaceSearchHook(m, v253, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v300 = F_lappend_oid(m, v240, v253)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	if v296 == int32(0) {
		v302 = v240
		goto L74
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v302 = v300
	goto L74
L95:
	;
	goto L73
L96:
	;
	v309 = int32(0)
	v314 = v309
	v323 = v309
	goto L69
L97:
	;
	goto L98
L98:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v314 = v302
	v323 = v312
	goto L69
L99:
	;
	if v364 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L100:
	;
	v364 = int32(0)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v332 <= int32(0) {
		v358 = v326
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v364 = v358
	goto L99
L104:
	;
	v335 = int32(0)
	if v335 < v332 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v338 = v332
	goto L107
L106:
	;
	v338 = v335
	goto L107
L107:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	v341 = int32(0)
	goto L108
L108:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v339+v341<<(uint(int32(2))%32))))
	v350 = base.B2i32(v349 == int32(11))
	if v349 == int32(11) {
		v358 = v350
		goto L103
	} else {
		goto L110
	}
L109:
	;
	v358 = v350
	goto L103
L110:
	;
	v352 = v341 + int32(1)
	if v352 != v338 {
		v341 = v352
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v368 = F_lcons_oid(m, int32(11), v314)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L115
	}
L113:
	;
	v370 = v314
	goto L114
L114:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	if v372 == int32(0) {
		v418 = v370
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v370 = v368
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v418
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v225
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[10]))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)) = uint8(base.B2i32(v423 != int32(0)))
	v428 = v418
	goto L63
L117:
	;
	v375 = int32(0)
	if v370 == v375 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v413 != 0 {
		v418 = v370
		goto L116
	} else {
		goto L131
	}
L119:
	;
	v413 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v381 <= int32(0) {
		v407 = v375
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v413 = v407
	goto L118
L123:
	;
	v384 = int32(0)
	if v384 < v381 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v387 = v381
	goto L126
L125:
	;
	v387 = v384
	goto L126
L126:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v390 = int32(0)
	goto L127
L127:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v388+v390<<(uint(int32(2))%32))))
	v399 = base.B2i32(v398 == v372)
	if v398 == v372 {
		v407 = v399
		goto L122
	} else {
		goto L129
	}
L128:
	;
	v407 = v399
	goto L122
L129:
	;
	v401 = v390 + int32(1)
	if v401 != v387 {
		v390 = v401
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[8]))
	v416 = F_lcons_oid(m, v415, v370)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v418 = v416
	goto L116
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2])) = v16
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[12])) = v494
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[13])) = v498
	v502 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])))
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[15])) = uint8(v502)
	v505 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])) = uint8(v505)
	goto L2
L134:
	;
	v451 = int32(_a_F_recomputeNamespacePath_0)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4]))
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v455
	v457 = F_list_copy(m, v450)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L139
	}
L135:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)))
	if v442 != v443 {
		v450 = v428
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	v447 = F_equal(m, v428, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	if v447 != 0 {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v450 = v449
	goto L134
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[4])) = v452
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11]))
	F_list_free(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[11])) = v457
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[9])) = v468
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)))
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[14])) = uint8(v471)
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[2])) = v16
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[12])) = v457
	*(*int32)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[13])) = v468
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[15])) = uint8(v471)
	v481 = int32(_a_F_recomputeNamespacePath_3)
	v483 = *(*int64)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[17])) = v483 + int64(1)
	v488 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_recomputeNamespacePath[1])) = uint8(v488)
	goto L2
L141:
	;
	F_errmsg_internal(m, int32(_a_F_recomputeNamespacePath_4), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_recomputeNamespacePath_5), int32(_a_F_recomputeNamespacePath_6), int32(_a_F_recomputeNamespacePath_7))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
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
