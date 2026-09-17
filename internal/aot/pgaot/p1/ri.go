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
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_check_ins_0), int32(1))
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
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_noaction_del_0), int32(3))
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
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setdefault_del_0), int32(3))
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L11
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L80
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L11
	} else {
		goto L77
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[0]))
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
	v305 = m.ExcPending
	if v305 != 0 {
		goto L11
	} else {
		goto L72
	}
L8:
	;
	v58 = v18
	goto L10
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = int64(3023656976388)
	v25 = v13 + int32(112)
	v27 = F_hash_create(m, int32(_a_F_ri_FetchConstraintInfo_0), int32(64), v25, int32(40))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v64 = F_hash_search(m, v58, v13+int32(108), int32(1), v13+int32(112))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L16
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[0])) = v27
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1485), int32(0))
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
	v43 = F_hash_create(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(256), v25, int32(40))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[1])) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = int64(292057776136)
	v52 = F_hash_create(m, int32(_a_F_ri_FetchConstraintInfo_2), int32(256), v25, int32(40))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[2])) = v52
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[0]))
	v58 = v56
	goto L10
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+112)))
	if v66 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v64)+88))
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v74 = F_SearchSysCache1(m, int32(19), v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L23
	}
L19:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)) = uint8(v69)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	if v71 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v74 == int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
	v80 = v78 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+72)))
	if v81 != int32(102) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+92))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v122 = F_GetSysCacheHashValue(m, int32(19), v115, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L36
	}
L27:
	;
	v90 = v84
	goto L30
L28:
	;
	goto L29
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v108
	v115 = v108
	goto L26
L30:
	;
	v96 = F_SearchSysCache1(m, int32(19), v90)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v90
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v115 = v107
	goto L26
L32:
	;
	if v96 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)+92))
	F_ReleaseCatCache(m, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	if v103 != 0 {
		v90 = v103
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v122
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v128 = F_GetSysCacheHashValue(m, int32(19), v126, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v128
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+20)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+28)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v80)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+36)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v80)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+44)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v80)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+52)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v80)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+60)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v80)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+68)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v80)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v64)+76)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v80)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+92)) = uint8(v151)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+93)) = uint8(v153)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+164)) = uint8(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+107)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+165)) = uint8(v157)
	F_DeconstructFkConstraintRow(m, v74, v64+int32(168), v64+int32(236), v64+int32(172), v64+int32(300), v64+int32(428), v64+int32(556), v64+int32(96), v64+int32(100))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+165)))
	if v177 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v64)+168))
	v182 = F_get_index_column_opclass(m, v180, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_ReleaseCatCache(m, v74)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	F_FindFKPeriodOpers(m, v182, v64+int32(684), v64+int32(688), v64+int32(692))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v195 = v64 + int32(696)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[3]))
	if v197 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+696)) = v207
	v209 = int32(_a_F_ri_FetchConstraintInfo_3)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+700)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v195
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[4])) = v195
	v214 = int32(_a_F_ri_FetchConstraintInfo_4)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[5]))
	v217 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[5])) = v216 + v217
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)) = uint8(v217)
	goto L17
L46:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[4]))
	v207 = v199
	goto L45
L47:
	;
	goto L48
L48:
	;
	v201 = int32(_a_F_ri_FetchConstraintInfo_3)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[3])) = v201
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchConstraintInfo[5])) = int32(0)
	v207 = v201
	goto L45
L49:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+164)))
	switch v263 - int32(102) {
	case 0, 13:
		goto L62
	default:
		goto L64
	case 10:
		goto L63
	}
L50:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v233 == v232 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v232 != v258 {
		goto L1
	} else {
		goto L60
	}
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v235 == v236 {
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
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v242 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_5), v13+int32(32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2240), int32(_a_F_ri_FetchConstraintInfo_7))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
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
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v260 != v261 {
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
	return v64
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L11
	} else {
		goto L68
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v270 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+164)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v270
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_8), v13+int32(16))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2254), int32(_a_F_ri_FetchConstraintInfo_7))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	v288 = m.ExcPending
	if v288 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_ri_FetchConstraintInfo_9), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2259), int32(_a_F_ri_FetchConstraintInfo_7))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
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
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v309 + int32(4)
	F_errmsg(m, int32(_a_F_ri_FetchConstraintInfo_10), v13)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_errhint(m, int32(_a_F_ri_FetchConstraintInfo_11), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2229), int32(_a_F_ri_FetchConstraintInfo_7))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
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
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v331
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_12), v13-int32(-64))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2297), int32(_a_F_ri_FetchConstraintInfo_13))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v347
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_14), v13+int32(96))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2302), int32(_a_F_ri_FetchConstraintInfo_13))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v90
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_12), v13+int32(80))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2375), int32(_a_F_ri_FetchConstraintInfo_15))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
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
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v378 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_5), v13+int32(48))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(2247), int32(_a_F_ri_FetchConstraintInfo_7))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
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
