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
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v2 = int32(235520)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[214])))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		v25 = v5
		v26 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v55
L2:
	;
	if v26-v25 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v5 != v6 {
		v25 = v5
		v26 = v6
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v10 = l0
	v11 = v2
	goto L6
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(0) {
		v25 = v14
		v26 = v15
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v25 = v14
	v26 = v15
	goto L3
L8:
	;
	v18 = int32(1)
	if v14 == v15 {
		v10 = v10 + v18
		v11 = v11 + v18
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v30 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v32 == v30 {
		v55 = v30
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v49 = int32(0)
	v52 = F_GetSysCacheOid(m, int32(37), l0, v49, v49, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L19
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v36 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return v32
L15:
	;
	goto L16
L16:
	;
	v41 = F_RunNamespaceSearchHook(m, v32, int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	return v46
L19:
	;
	v55 = v52
	goto L1
}
func F_get_namespace_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(38), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
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
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
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
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v494 int64
	_ = v494
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	v1 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[126])))
	if v18 == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v22 == v16 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	F_spcache_init(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return
L7:
	;
	v28 = F_spcache_insert(m, v25, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L12
	}
L8:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v449 != v450 {
		v461 = v438
		goto L140
	} else {
		goto L141
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v334
	v337 = int32(0)
	if v324 == v337 {
		goto L106
	} else {
		goto L107
	}
L10:
	;
	if v300 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L99
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v30 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = int32(4515392)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v37
	v39 = F_pstrdup(m, v25)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v210 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v211 == v210 {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	v44 = F_SplitIdentifierString(m, v39, int32(44), v13+int32(12))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v44 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v50 == v48 {
		v186 = v1
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v39)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L66
	}
L20:
	;
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v54 <= v53 {
		v186 = v1
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v58 = v53
	v61 = v1
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v58<<(uint(int32(2))%32))))
	v72 = int32(217582)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[213])))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v76 == int32(0) {
		v95 = v75
		v96 = v76
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v186 = v176
	goto L19
L24:
	;
	v179 = v58 + int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v179 < v180 {
		v58 = v179
		v61 = v176
		goto L22
	} else {
		goto L65
	}
L25:
	;
	if v96-v95 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	goto L25
L27:
	;
	if v75 != v76 {
		v95 = v75
		v96 = v76
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v80 = v71
	v81 = v72
	goto L29
L29:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v84
		v96 = v85
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v95 = v84
	v96 = v85
	goto L26
L31:
	;
	v88 = int32(1)
	if v84 == v85 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v101 = F_SearchSysCache1(m, int32(11), v16)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v126 = int32(235520)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[214])))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v130 == int32(0) {
		v149 = v129
		v150 = v130
		goto L45
	} else {
		goto L46
	}
L36:
	;
	if v101 == int32(0) {
		v176 = v61
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
	v111 = int32(0)
	v114 = F_GetSysCacheOid(m, int32(37), v106+v107+int32(4), v111, v111, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_ReleaseCatCache(m, v101)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v114 == int32(0) {
		v176 = v61
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v122 = F_object_aclcheck(m, int32(2615), v114, v16, int64(256))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v122 != 0 {
		v176 = v61
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v124 = F_lappend_oid(m, v61, v114)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v176 = v124
	goto L24
L44:
	;
	if v150-v149 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	if v129 != v130 {
		v149 = v129
		v150 = v130
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v134 = v71
	v135 = v126
	goto L48
L48:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v139 == int32(0) {
		v149 = v138
		v150 = v139
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v149 = v138
	v150 = v139
	goto L45
L50:
	;
	v142 = int32(1)
	if v138 == v139 {
		v134 = v134 + v142
		v135 = v135 + v142
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v155 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v162 = int32(0)
	v165 = F_GetSysCacheOid(m, int32(37), v71, v162, v162, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L60
	}
L55:
	;
	v156 = F_lappend_oid(m, v61, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v61 != 0 {
		v176 = v61
		goto L24
	} else {
		goto L59
	}
L58:
	;
	v176 = v156
	goto L24
L59:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)) = uint8(v158)
	v176 = int32(0)
	goto L24
L60:
	;
	if v165 == int32(0) {
		v176 = v61
		goto L24
	} else {
		goto L61
	}
L61:
	;
	v171 = F_object_aclcheck(m, int32(2615), v165, v16, int64(256))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	if v171 != 0 {
		v176 = v61
		goto L24
	} else {
		goto L63
	}
L63:
	;
	v173 = F_lappend_oid(m, v61, v165)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v176 = v173
	goto L24
L65:
	;
	goto L23
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_list_free(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v186
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
	goto L15
L68:
	;
	F_list_free(m, v211)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L72
	}
L69:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v215 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)))
	if v216 == int32(1) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v438 = v211
	goto L8
L72:
	;
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v221
	v223 = int32(4515392)
	v224 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v227 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v227
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v230 == v221 {
		v324 = v210
		v334 = v221
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v233 = int32(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v234 <= v233 {
		v324 = v210
		v334 = v233
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v237 = v210
	v238 = v233
	goto L75
L75:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v238<<(uint(int32(2))%32))))
	v252 = int32(0)
	if v237 == v252 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L10
L77:
	;
	v302 = v238 + int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v302 < v303 {
		v237 = v300
		v238 = v302
		goto L75
	} else {
		goto L98
	}
L78:
	;
	if v290 != 0 {
		v300 = v237
		goto L77
	} else {
		goto L91
	}
L79:
	;
	v290 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v258 <= int32(0) {
		v283 = v252
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v290 = v283
	goto L78
L83:
	;
	v261 = int32(0)
	if v261 < v258 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v264 = v258
	goto L86
L85:
	;
	v264 = v261
	goto L86
L86:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v267 = int32(0)
	goto L87
L87:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v265+v267<<(uint(int32(2))%32))))
	v276 = base.B2i32(v275 == v251)
	if v275 == v251 {
		v283 = v276
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v283 = v276
	goto L82
L89:
	;
	v278 = v267 + int32(1)
	if v278 != v264 {
		v267 = v278
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v292 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v294 = F_RunNamespaceSearchHook(m, v251, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v298 = F_lappend_oid(m, v237, v251)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L97
	}
L95:
	;
	if v294 == int32(0) {
		v300 = v237
		goto L77
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v300 = v298
	goto L77
L98:
	;
	goto L76
L99:
	;
	F_errmsg_internal(m, int32(29507), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(499480), int32(4126), int32(321984))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v320 = int32(0)
	v324 = v320
	v334 = v320
	goto L9
L103:
	;
	goto L104
L104:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v324 = v300
	v334 = v323
	goto L9
L105:
	;
	if v375 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L106:
	;
	v375 = int32(0)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v343 <= int32(0) {
		v368 = v337
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v375 = v368
	goto L105
L110:
	;
	v346 = int32(0)
	if v346 < v343 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v349 = v343
	goto L113
L112:
	;
	v349 = v346
	goto L113
L113:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v352 = int32(0)
	goto L114
L114:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v350+v352<<(uint(int32(2))%32))))
	v361 = base.B2i32(v360 == int32(11))
	if v360 == int32(11) {
		v368 = v361
		goto L109
	} else {
		goto L116
	}
L115:
	;
	v368 = v361
	goto L109
L116:
	;
	v363 = v352 + int32(1)
	if v363 != v349 {
		v352 = v363
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v379 = F_lcons_oid(m, int32(11), v324)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	v381 = v324
	goto L120
L120:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v383 == int32(0) {
		v429 = v381
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v381 = v379
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v429
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v224
	v434 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)) = uint8(base.B2i32(v434 != int32(0)))
	v438 = v429
	goto L8
L123:
	;
	v386 = int32(0)
	if v381 == v386 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v424 != 0 {
		v429 = v381
		goto L122
	} else {
		goto L137
	}
L125:
	;
	v424 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v392 <= int32(0) {
		v417 = v386
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v424 = v417
	goto L124
L129:
	;
	v395 = int32(0)
	if v395 < v392 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v398 = v392
	goto L132
L131:
	;
	v398 = v395
	goto L132
L132:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v401 = int32(0)
	goto L133
L133:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v399+v401<<(uint(int32(2))%32))))
	v410 = base.B2i32(v409 == v383)
	if v409 == v383 {
		v417 = v410
		goto L128
	} else {
		goto L135
	}
L134:
	;
	v417 = v410
	goto L128
L135:
	;
	v412 = v401 + int32(1)
	if v412 != v398 {
		v401 = v412
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v427 = F_lcons_oid(m, v426, v381)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v429 = v427
	goto L122
L139:
	;
	*(*int32)(unsafe.Add(mBase, _consts[209])) = v16
	v505 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	*(*int32)(unsafe.Add(mBase, _consts[216])) = v505
	v509 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	*(*int32)(unsafe.Add(mBase, _consts[217])) = v509
	v513 = int32(*(*uint8)(unsafe.Add(mBase, _consts[218])))
	*(*uint8)(unsafe.Add(mBase, _consts[219])) = uint8(v513)
	v516 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v516)
	goto L1
L140:
	;
	v462 = int32(4515392)
	v463 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v466 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v466
	v468 = F_list_copy(m, v461)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L6
	} else {
		goto L145
	}
L141:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, _consts[218])))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)))
	if v453 != v454 {
		v461 = v438
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	v458 = F_equal(m, v438, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v458 != 0 {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v461 = v460
	goto L140
L145:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v463
	v473 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	F_list_free(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[215])) = v468
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int32)(unsafe.Add(mBase, _consts[211])) = v479
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)))
	*(*uint8)(unsafe.Add(mBase, _consts[218])) = uint8(v482)
	*(*int32)(unsafe.Add(mBase, _consts[209])) = v16
	*(*int32)(unsafe.Add(mBase, _consts[216])) = v468
	*(*int32)(unsafe.Add(mBase, _consts[217])) = v479
	*(*uint8)(unsafe.Add(mBase, _consts[219])) = uint8(v482)
	v492 = int32(4121288)
	v494 = *(*int64)(unsafe.Add(mBase, _consts[220]))
	*(*int64)(unsafe.Add(mBase, _consts[220])) = v494 + int64(1)
	v499 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v499)
	goto L1
}
