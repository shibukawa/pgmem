package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_check_ins(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	F_ri_CheckTrigger(m, l0, int32(149815), int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_RI_FKey_check(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_noaction_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(307929), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_setdefault_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(307885), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(0), int32(3))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L11
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L11
	} else {
		goto L80
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L11
	} else {
		goto L77
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L72
	}
L8:
	;
	v61 = v18
	goto L10
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = int64(3023656976388)
	v27 = F_hash_create(m, int32(399119), int32(64), v13+int32(112), int32(40))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v67 = F_hash_search(m, v61, v13+int32(108), int32(1), v13+int32(112))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L16
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1126])) = v27
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1501), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = int64(51539607560)
	v45 = F_hash_create(m, int32(399020), int32(256), v13+int32(112), int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1127])) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = int64(292057776136)
	v56 = F_hash_create(m, int32(399433), int32(256), v13+int32(112), int32(40))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1128])) = v56
	v60 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	v61 = v60
	goto L10
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+112)))
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v67)+88))
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v77 = F_SearchSysCache1(m, int32(19), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L23
	}
L19:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)) = uint8(v72)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v74 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v77 == int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+22)))
	v83 = v81 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+72)))
	if v84 != int32(102) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+92))
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v125 = F_GetSysCacheHashValue(m, int32(19), v118, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L36
	}
L27:
	;
	v93 = v87
	goto L30
L28:
	;
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v111
	v118 = v111
	goto L26
L30:
	;
	v99 = F_SearchSysCache1(m, int32(19), v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v93
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v118 = v110
	goto L26
L32:
	;
	if v99 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+22)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104)+92))
	F_ReleaseCatCache(m, v99)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	if v106 != 0 {
		v93 = v106
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v125
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v131 = F_GetSysCacheHashValue(m, int32(19), v129, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v131
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+20)) = v134
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v83)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+28)) = v136
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+36)) = v138
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v83)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+44)) = v140
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v83)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+52)) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v83)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+60)) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v83)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+68)) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v83)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+76)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v83)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+84)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+88)) = v152
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+92)) = uint8(v154)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+93)) = uint8(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+164)) = uint8(v158)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+107)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+165)) = uint8(v160)
	F_DeconstructFkConstraintRow(m, v77, v67+int32(168), v67+int32(236), v67+int32(172), v67+int32(300), v67+int32(428), v67+int32(556), v67+int32(96), v67+int32(100))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+165)))
	if v180 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v83)+88))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v67)+168))
	v185 = F_get_index_column_opclass(m, v183, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_ReleaseCatCache(m, v77)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	F_FindFKPeriodOpers(m, v185, v67+int32(684), v67+int32(688), v67+int32(692))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v198 = v67 + int32(696)
	v200 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	if v200 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+696)) = v210
	v212 = int32(4501944)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+700)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v198
	*(*int32)(unsafe.Add(mBase, _consts[1130])) = v198
	v217 = int32(4501952)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[1131]))
	v220 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1131])) = v219 + v220
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)) = uint8(v220)
	goto L17
L46:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[1130]))
	v210 = v202
	goto L45
L47:
	;
	goto L48
L48:
	;
	v204 = int32(4501944)
	*(*int32)(unsafe.Add(mBase, _consts[1129])) = v204
	*(*int32)(unsafe.Add(mBase, _consts[1131])) = int32(0)
	v210 = v204
	goto L45
L49:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+164)))
	switch v266 - int32(102) {
	case 0, 13:
		goto L62
	default:
		goto L64
	case 10:
		goto L63
	}
L50:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v236 == v235 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v235 != v261 {
		goto L1
	} else {
		goto L60
	}
L53:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v238 == v239 {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v245 + int32(4)
	F_errmsg_internal(m, int32(718769), v13+int32(32))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(493696), int32(2240), int32(242185))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v263 != v264 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L49
L62:
	;
	m.G0 = v13 + int32(160)
	return v67
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L11
	} else {
		goto L68
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+164)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v273
	F_errmsg_internal(m, int32(484127), v13+int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(493696), int32(2254), int32(242185))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(446394), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(493696), int32(2259), int32(242185))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v312 + int32(4)
	F_errmsg(m, int32(718715), v13)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_errhint(m, int32(656948), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(493696), int32(2229), int32(242185))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v334
	F_errmsg_internal(m, int32(41029), v13-int32(-64))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(493696), int32(2297), int32(242208))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v350
	F_errmsg_internal(m, int32(89386), v13+int32(96))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(493696), int32(2302), int32(242208))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v93
	F_errmsg_internal(m, int32(41029), v13+int32(80))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(493696), int32(2375), int32(84358))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v381 + int32(4)
	F_errmsg_internal(m, int32(718769), v13+int32(48))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(493696), int32(2247), int32(242185))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
