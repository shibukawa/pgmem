package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_cascade_del(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	v16 = m.G0
	v18 = v16 - int32(624)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_cascade_del_0), int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = F_ri_FetchConstraintInfo(m, v27, v28, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	v34 = F_table_open(m, v32, int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+620)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+616)) = v41
	v46 = v18 + int32(616)
	v47 = F_ri_FetchPreparedPlan(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v304 = v47
	goto L9
L8:
	;
	F_initStringInfo(m, v18+int32(600))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v305 = int32(0)
	v309 = F_ri_PerformCheck(m, v30, v46, v304, v34, v37, v36, v305, v305, int32(1), int32(8))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L53
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v56 = F_get_namespace_name(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v58 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)) = uint8(v58)
	v62 = v18 + int32(336)
	v64 = v56
	goto L12
L12:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v77 != int32(34) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v94 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+1)) = uint16(v94)
	v97 = v18 + int32(336)
	v98 = F_strlen(m, v97)
	mBase = m.M
	v99 = v98 + v97
	v100 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)) = uint8(v94)
	v109 = v102 + int32(4)
	v111 = v99 + int32(1)
	goto L20
L14:
	;
	goto L13
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v89)
	v62 = v90
	v64 = v64 + int32(1)
	goto L12
L16:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v84 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v84)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v89 = v86
	v90 = v62 + int32(2)
	goto L15
L19:
	;
	v89 = v77
	v90 = v62 + int32(1)
	goto L15
L20:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v124 != int32(34) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v141 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v111)+1)) = uint16(v141)
	if v54 == int32(112) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v136)
	v109 = v109 + int32(1)
	v111 = v137
	goto L20
L24:
	;
	if v124 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v131 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)) = uint8(v131)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v136 = v133
	v137 = v111 + int32(2)
	goto L23
L27:
	;
	v136 = v124
	v137 = v111 + int32(1)
	goto L23
L28:
	;
	v147 = int32(_a_F_RI_FKey_cascade_del_1)
	goto L30
L29:
	;
	v147 = int32(_a_F_RI_FKey_cascade_del_2)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(336)
	F_appendStringInfo(m, v18+int32(600), int32(_a_F_RI_FKey_cascade_del_3), v18+int32(32))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if int32(0) < v159 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v174 = int32(0)
	v177 = int32(_a_F_RI_FKey_cascade_del_4)
	goto L35
L33:
	;
	v267 = v159
	goto L34
L34:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v18)+600))
	v287 = F_ri_PlanCheck(m, v282, v267, v18+int32(48), v18+int32(616), v34, v37)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L52
	}
L35:
	;
	v186 = v174 << (uint(int32(1)) % 32)
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(172)+v186))))
	v189 = F_attnumTypeId(m, v37, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v267 = v265
	goto L34
L37:
	;
	v191 = v186 + (v30 + int32(236))
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v191))))
	v193 = F_attnumTypeId(m, v34, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v191))))
	v196 = F_attnumAttName(m, v34, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v198 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+192)) = uint8(v198)
	v202 = v18 + int32(192)
	v204 = v196
	goto L40
L40:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v217 != int32(34) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v234 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v202)+1)) = uint16(v234)
	v237 = v174 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v237
	v240 = v18 + int32(176)
	v244 = F_pg_sprintf(m, v240, int32(_a_F_RI_FKey_cascade_del_5), v18+int32(16))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v229)
	v202 = v230
	v204 = v204 + int32(1)
	goto L40
L44:
	;
	if v217 == int32(0) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v224 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)) = uint8(v224)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v229 = v226
	v230 = v202 + int32(2)
	goto L43
L47:
	;
	v229 = v217
	v230 = v202 + int32(1)
	goto L43
L48:
	;
	v247 = v174 << (uint(int32(2)) % 32)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(300)+v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v177
	v252 = v18 + int32(600)
	F_appendStringInfo(m, v252, int32(_a_F_RI_FKey_cascade_del_6), v18)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_generate_operator_clause(m, v252, v240, v189, v249, v18+int32(192), v193)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(48)+v247))) = v189
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if v237 < v265 {
		v174 = v237
		v177 = int32(_a_F_RI_FKey_cascade_del_7)
		goto L35
	} else {
		goto L51
	}
L51:
	;
	goto L36
L52:
	;
	v304 = v287
	goto L9
L53:
	;
	v311 = F_SPI_finish(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v311 != int32(2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	F_errmsg_internal(m, int32(_a_F_RI_FKey_cascade_del_8), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_cascade_del_9), int32(1094), int32(_a_F_RI_FKey_cascade_del_0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	m.G0 = v18 + int32(624)
	return int64(0)
}
func F_RI_FKey_cascade_upd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	v20 = m.G0
	v22 = v20 - int32(784)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_cascade_upd_0), int32(2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v34 = F_ri_FetchConstraintInfo(m, v31, v32, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v38 = F_table_open(m, v36, int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+780)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+776)) = v46
	v51 = v22 + int32(776)
	v52 = F_ri_FetchPreparedPlan(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v52 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v364 = v52
	goto L9
L8:
	;
	F_initStringInfo(m, v22+int32(760))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v368 = F_ri_PerformCheck(m, v34, v51, v364, v38, v42, v40, v41, int32(0), int32(1), int32(9))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L56
	}
L10:
	;
	F_initStringInfo(m, v22+int32(744))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+119)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+68))
	v65 = F_get_namespace_name(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v67 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+480)) = uint8(v67)
	v71 = v22 + int32(480)
	v73 = v65
	goto L13
L13:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v90 != int32(34) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v107 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+1)) = uint16(v107)
	v110 = v22 + int32(480)
	v111 = F_strlen(m, v110)
	mBase = m.M
	v112 = v111 + v110
	v113 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)) = uint8(v107)
	v122 = v115 + int32(4)
	v124 = v112 + int32(1)
	goto L21
L15:
	;
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v102)
	v71 = v103
	v73 = v73 + int32(1)
	goto L13
L17:
	;
	if v90 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)) = uint8(v97)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v102 = v99
	v103 = v71 + int32(2)
	goto L16
L20:
	;
	v102 = v90
	v103 = v71 + int32(1)
	goto L16
L21:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v141 != int32(34) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v158 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v124)+1)) = uint16(v158)
	v160 = int32(_a_F_RI_FKey_cascade_upd_1)
	if v63 == int32(112) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v153)
	v122 = v122 + int32(1)
	v124 = v154
	goto L21
L25:
	;
	if v141 == int32(0) {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v148 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v153 = v150
	v154 = v124 + int32(2)
	goto L24
L28:
	;
	v153 = v141
	v154 = v124 + int32(1)
	goto L24
L29:
	;
	v165 = v160
	goto L31
L30:
	;
	v165 = int32(_a_F_RI_FKey_cascade_upd_2)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v22 + int32(480)
	F_appendStringInfo(m, v22+int32(760), int32(_a_F_RI_FKey_cascade_upd_3), v22+int32(48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v34)+168))
	if int32(0) < v177 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v191 = int32(0)
	v195 = v177
	v196 = v160
	v197 = int32(_a_F_RI_FKey_cascade_upd_4)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v22)+744))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v22)+748))
	F_appendBinaryStringInfo(m, v22+int32(760), v331, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L54
	}
L36:
	;
	v208 = v191 << (uint(int32(1)) % 32)
	v210 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34+int32(172)+v208))))
	v211 = F_attnumTypeId(m, v42, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v213 = v208 + (v34 + int32(236))
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v213))))
	v215 = F_attnumTypeId(m, v38, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v213))))
	v218 = F_attnumAttName(m, v38, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v220 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+336)) = uint8(v220)
	v224 = v22 + int32(336)
	v226 = v218
	goto L41
L41:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v243 != int32(34) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v260 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+1)) = uint16(v260)
	v263 = v191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v196
	v267 = v22 + int32(336)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v267
	F_appendStringInfo(m, v22+int32(760), int32(_a_F_RI_FKey_cascade_upd_5), v22+int32(32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L49
	}
L43:
	;
	goto L42
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v256))) = uint8(v255)
	v224 = v256
	v226 = v226 + int32(1)
	goto L41
L45:
	;
	if v243 == int32(0) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v250 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)) = uint8(v250)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v255 = v252
	v256 = v224 + int32(2)
	goto L44
L48:
	;
	v255 = v243
	v256 = v224 + int32(1)
	goto L44
L49:
	;
	v277 = v195 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v277
	v280 = v22 + int32(320)
	v284 = F_pg_sprintf(m, v280, int32(_a_F_RI_FKey_cascade_upd_6), v22+int32(16))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v287 = v191 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(300)+v287)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v197
	v292 = v22 + int32(744)
	F_appendStringInfo(m, v292, int32(_a_F_RI_FKey_cascade_upd_7), v22)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_generate_operator_clause(m, v292, v280, v211, v289, v267, v215)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v299 = v22 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v287+v299))) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v195<<(uint(int32(2))%32)+v299))) = v211
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v34)+168))
	if v263 < v308 {
		v191 = v263
		v195 = v277
		v196 = int32(_a_F_RI_FKey_cascade_upd_8)
		v197 = int32(_a_F_RI_FKey_cascade_upd_9)
		goto L36
	} else {
		goto L53
	}
L53:
	;
	goto L37
L54:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v22)+760))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v34)+168))
	v343 = F_ri_PlanCheck(m, v335, v336<<(uint(int32(1))%32), v22-int32(-64), v22+int32(776), v38, v42)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v364 = v343
	goto L9
L56:
	;
	v370 = F_SPI_finish(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v370 != int32(2) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_relation_close(m, v38, int32(3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F_RI_FKey_cascade_upd_10), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_cascade_upd_11), int32(1211), int32(_a_F_RI_FKey_cascade_upd_0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	m.G0 = v22 + int32(784)
	return int64(0)
}
func F_RI_FKey_check(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v269 int32
	_ = v269
	var v284 int32
	_ = v284
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v752 int32
	_ = v752
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int64
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1326 int32
	_ = v1326
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int64
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1447 int64
	_ = v1447
	var v1449 int64
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1462 int64
	_ = v1462
	var v1464 int64
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1602 int64
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1673 int32
	_ = v1673
	var v1675 int64
	_ = v1675
	var v1678 int64
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1698 int32
	_ = v1698
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1757 int64
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1860 int32
	_ = v1860
	var v1865 int32
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1955 int64
	_ = v1955
	var v1956 int64
	_ = v1956
	var v1957 int64
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2105 int32
	_ = v2105
	v2 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(2528)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v37 = F_ri_FetchConstraintInfo(m, v34, v35, v2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42&int32(3) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L289
	}
L4:
	;
	m.G0 = v32 + int32(2528)
	return
L5:
	;
	v47 = int32(28)
	goto L7
L6:
	;
	v47 = int32(24)
	goto L7
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v39)+188))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+72))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v39, v49, int32(_a_F_RI_FKey_check_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v53 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if v57 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = v37 + int32(236)
	v64 = int32(1)
	v66 = int32(0)
	v69 = v64
	v71 = v57
	v72 = v64
	goto L11
L11:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62+v66<<(uint(int32(1))%32)))))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+6)))
	if v99 < v98 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v112&int32(1) != 0 {
		goto L4
	} else {
		goto L18
	}
L13:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	m.T0[v102].(func(*base.Module, int32, int32))(m, v49, v98)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v106 = v71
	goto L15
L15:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	v109 = int32(1)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v98-v109))))
	v112 = v111 & v72
	v115 = (v111 ^ v109) & v69
	v117 = v66 + v109
	if v117 < v106 {
		v66 = v117
		v69 = v115
		v71 = v106
		v72 = v112
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	v106 = v105
	goto L15
L17:
	;
	goto L12
L18:
	;
	if v115&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+708)))
	if v159 != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+164)))
	v125 = v123 - int32(102)
	if v125 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v125 == int32(13) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	goto L4
L25:
	;
	goto L19
L27:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	v137 = v37 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+132)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v135 + int32(4)
	F_errmsg(m, int32(_a_F_RI_FKey_check_1), v32+int32(128))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v149 = F_errdetail(m, int32(_a_F_RI_FKey_check_2), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errtableconstraint(m, v60, v137)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(390), int32(_a_F_RI_FKey_check_4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v1085 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L155
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+712)) = int32(1)
	goto L33
L35:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L73
	}
L36:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+165)))
	if v160 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v37)+712))
	if v161 == int32(2) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v37)+84))
	v166 = F_table_open(m, v164, int32(2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v171 = F_ri_LoadConstraintInfo(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)+704))
	v175 = F_index_open(m, v173, int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+4)))
	if v177 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v181 = F_ri_LoadConstraintInfo(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v183 = v171
	goto L45
L45:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+712))
	switch v184 {
	case 0:
		goto L48
	case 1:
		goto L33
	default:
		goto L47
	}
L46:
	;
	v183 = v181
	goto L45
L47:
	;
	F_relation_close(m, v175, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L71
	}
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v175)+48))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+84))
	if v186 != int32(403) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+712)) = int32(2)
	goto L47
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v175)+192))
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v189)+10)))
	if int32(0) < v190 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v194 = int32(0)
	v196 = v189
	goto L54
L52:
	;
	goto L53
L53:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v183)+168))
	if v269 <= int32(0) {
		goto L34
	} else {
		goto L59
	}
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v175)+248))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223+v194<<(uint(int32(2))%32))))
	v231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196+v194<<(uint(int32(1))%32))+48)))
	v232 = F_attnumCollationId(m, v166, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	if v227 != v232 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v236 = v194 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v175)+192))
	v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v237)+10)))
	if v236 < v238 {
		v194 = v236
		v196 = v237
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v284 = v2
	goto L60
L60:
	;
	v305 = int32(0)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v175)+192))
	v307 = int32(*(*int16)(unsafe.Add(mBase, uint32(v306)+10)))
	if v307 <= v305 {
		v353 = v305
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L34
L62:
	;
	v382 = int32(2)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(300)+v284<<(uint(v382)%32))))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v175)+208))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386+v353<<(uint(v382)%32))))
	v391 = F_get_op_opfamily_strategy(m, v385, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L68
	}
L63:
	;
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+int32(172)+v284<<(uint(int32(1))%32)))))
	v316 = v305
	goto L64
L64:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306+int32(48)+v316<<(uint(int32(1))%32)))))
	if v348 == v315 {
		v353 = v316
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v353 = v307
	goto L62
L66:
	;
	v351 = v316 + int32(1)
	if v351 != v307 {
		v316 = v351
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v391 != int32(3) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	v396 = v284 + int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v183)+168))
	if v396 < v397 {
		v284 = v396
		goto L60
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	F_relation_close(m, v166, int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L35
L73:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v37)+84))
	v499 = F_table_open(m, v497, int32(2))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+152)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = v501
	v506 = v32 + int32(148)
	v507 = F_ri_FetchPreparedPlan(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v507 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v995 = v507
	goto L78
L77:
	;
	F_initStringInfo(m, v32+int32(2240))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v996 = int32(0)
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v499)+48))
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998)+119)))
	v1003 = F_ri_PerformCheck(m, v37, v506, v995, v60, v499, v996, v49, v996, base.B2i32(v999 == int32(112)), int32(5))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L146
	}
L79:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v499)+48))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+119)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v513)+68))
	v516 = F_get_namespace_name(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v518 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+160)) = uint8(v518)
	v522 = v516
	v524 = v32 + int32(160)
	goto L81
L81:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v551 != int32(34) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v568 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v524)+1)) = uint16(v568)
	v571 = v32 + int32(160)
	v572 = F_strlen(m, v571)
	mBase = m.M
	v573 = v572 + v571
	v574 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v573))) = uint8(v574)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v499)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)) = uint8(v568)
	v583 = v573 + int32(1)
	v585 = v576 + int32(4)
	goto L89
L83:
	;
	goto L82
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v564))) = uint8(v563)
	v522 = v522 + int32(1)
	v524 = v564
	goto L81
L85:
	;
	if v551 == int32(0) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v558 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v524)+1)) = uint8(v558)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	v563 = v560
	v564 = v524 + int32(2)
	goto L84
L88:
	;
	v563 = v551
	v564 = v524 + int32(1)
	goto L84
L89:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	if v612 != int32(34) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v629 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v583)+1)) = uint16(v629)
	if v514 == int32(112) {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	goto L90
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v625))) = uint8(v624)
	v583 = v625
	v585 = v585 + int32(1)
	goto L89
L93:
	;
	if v612 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v619 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+1)) = uint8(v619)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v624 = v621
	v625 = v583 + int32(2)
	goto L92
L96:
	;
	v624 = v612
	v625 = v583 + int32(1)
	goto L92
L97:
	;
	v635 = int32(_a_F_RI_FKey_check_5)
	goto L99
L98:
	;
	v635 = int32(_a_F_RI_FKey_check_6)
	goto L99
L99:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+165)))
	if v636 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if int32(0) < v752 {
		goto L115
	} else {
		goto L116
	}
L101:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v639<<(uint(int32(1))%32))+170)))
	v644 = F_attnumAttName(m, v499, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v32 + int32(160)
	F_appendStringInfo(m, v32+int32(2240), int32(_a_F_RI_FKey_check_7), v32+int32(80))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L114
	}
L104:
	;
	v646 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+2272)) = uint8(v646)
	v650 = v644
	v652 = v32 + int32(2272)
	goto L105
L105:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650))))
	if v679 != int32(34) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v696 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v652)+1)) = uint16(v696)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v32 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v32 + int32(2272)
	F_appendStringInfo(m, v32+int32(2240), int32(_a_F_RI_FKey_check_8), v32-int32(-64))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L113
	}
L107:
	;
	goto L106
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v692))) = uint8(v691)
	v650 = v650 + int32(1)
	v652 = v692
	goto L105
L109:
	;
	if v679 == int32(0) {
		goto L107
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v686 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v652)+1)) = uint8(v686)
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650))))
	v691 = v688
	v692 = v652 + int32(2)
	goto L108
L112:
	;
	v691 = v679
	v692 = v652 + int32(1)
	goto L108
L113:
	;
	goto L100
L114:
	;
	goto L100
L115:
	;
	v764 = int32(0)
	v769 = int32(_a_F_RI_FKey_check_9)
	goto L118
L116:
	;
	goto L117
L117:
	;
	v918 = v32 + int32(2240)
	F_appendStringInfoString(m, v918, int32(_a_F_RI_FKey_check_10))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L135
	}
L118:
	;
	v791 = v764 << (uint(int32(1)) % 32)
	v792 = v37 + int32(172) + v791
	v793 = int32(*(*int16)(unsafe.Add(mBase, uint32(v792))))
	v794 = F_attnumTypeId(m, v499, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v797 = int32(*(*int16)(unsafe.Add(mBase, uint32(v791+v62))))
	v798 = F_attnumTypeId(m, v60, v797)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v800 = int32(*(*int16)(unsafe.Add(mBase, uint32(v792))))
	v801 = F_attnumAttName(m, v499, v800)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v803 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+2272)) = uint8(v803)
	v807 = v801
	v809 = v32 + int32(2272)
	goto L123
L123:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v836 != int32(34) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v853 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v809)+1)) = uint16(v853)
	v856 = v764 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v856
	v859 = v32 + int32(1952)
	v863 = F_pg_sprintf(m, v859, int32(_a_F_RI_FKey_check_11), v32+int32(48))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L131
	}
L125:
	;
	goto L124
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v849))) = uint8(v848)
	v807 = v807 + int32(1)
	v809 = v849
	goto L123
L127:
	;
	if v836 == int32(0) {
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v843 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v809)+1)) = uint8(v843)
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	v848 = v845
	v849 = v809 + int32(2)
	goto L126
L130:
	;
	v848 = v836
	v849 = v809 + int32(1)
	goto L126
L131:
	;
	v866 = v764 << (uint(int32(2)) % 32)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(300)+v866)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v769
	v871 = v32 + int32(2240)
	F_appendStringInfo(m, v871, int32(_a_F_RI_FKey_check_12), v32+int32(32))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_generate_operator_clause(m, v871, v32+int32(2272), v794, v868, v859, v798)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(1984)+v866))) = v798
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if v856 < v886 {
		v764 = v856
		v769 = int32(_a_F_RI_FKey_check_13)
		goto L118
	} else {
		goto L134
	}
L134:
	;
	goto L119
L135:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+165)))
	if v922 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v923<<(uint(int32(1))%32))+234)))
	v928 = F_attnumTypeId(m, v60, v927)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v32)+2240))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	v964 = F_ri_PlanCheck(m, v958, v959, v32+int32(1984), v32+int32(148), v60, v499)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L145
	}
L139:
	;
	F_appendStringInfoString(m, v918, int32(_a_F_RI_FKey_check_14))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v933
	v936 = v32 + int32(1952)
	v940 = F_pg_sprintf(m, v936, int32(_a_F_RI_FKey_check_11), v32+int32(16))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v37)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(_a_F_RI_FKey_check_5)
	F_appendStringInfo(m, v918, int32(_a_F_RI_FKey_check_12), v32)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_generate_operator_clause(m, v918, v936, v928, v942, int32(_a_F_RI_FKey_check_15), int32(_a_F_RI_FKey_check_16))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_appendStringInfoString(m, v918, int32(_a_F_RI_FKey_check_17))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	v995 = v964
	goto L78
L146:
	;
	v1005 = F_SPI_finish(m)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v1005 == int32(2) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_relation_close(m, v499, int32(2))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L152
	}
L151:
	;
	goto L4
L152:
	;
	F_errmsg_internal(m, int32(_a_F_RI_FKey_check_18), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(551), int32(_a_F_RI_FKey_check_4))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v1087 = F_RegisterSnapshot(m, v1085)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_PushActiveSnapshot(m, v1087)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v1092 = F_table_slot_create(m, v166, int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(148)))) = v1099
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(156)))) = v1102
	goto L159
L159:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+80))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[1])) = v1106 | int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0])) = v1105
	goto L160
L160:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+68))
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0]))
	v1119 = F_object_aclcheck(m, int32(2615), v1115, v1117, int64(256))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v1119 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+68))
	v1124 = F_get_namespace_name(m, v1123)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v166)+56))
	v1130 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0]))
	v1132 = F_pg_class_aclmask(m, v1128, v1130, int64(6))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L168
	}
L165:
	;
	F_aclcheck_error(m, v1119, int32(37), v1124)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v1266 = int32(0)
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v171)+168))
	v1270 = F_index_beginscan(m, v166, v175, v1087, v1266, v1267, v1266, v1266)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L183
	}
L168:
	;
	if v1132 == int64(6) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1137 = F_palloc0(m, int32(40))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1137))) = int32(102)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v166)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1137)+16)) = int64(6)
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+4)) = v1141
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v171)+168))
	if int32(0) < v1145 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1137)+28))
	v1152 = int32(0)
	v1154 = v1150
	goto L174
L172:
	;
	goto L173
L173:
	;
	v1223 = F_ExecCheckOneRelPerms(m, v1137)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L178
	}
L174:
	;
	v1184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v171+int32(172)+v1152<<(uint(int32(1))%32)))))
	v1187 = F_bms_add_member(m, v1154, v1184+int32(7))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+28)) = v1187
	v1191 = v1152 + int32(1)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v171)+168))
	if v1191 < v1192 {
		v1152 = v1191
		v1154 = v1187
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1137)+28))
	F_bms_free(m, v1225)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_pfree(m, v1137)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	if v1223 != 0 {
		goto L167
	} else {
		goto L181
	}
L181:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	F_aclcheck_error(m, int32(1), int32(42), v1232+int32(4))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L167
L183:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v171)+716))
	if v1272 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v1523 <= int32(0) {
		v1766 = v1523
		goto L209
	} else {
		goto L210
	}
L185:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v171)+168))
	v1523 = v1273
	v1526 = v171
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v1275 = F_ri_LoadConstraintInfo(m, v1274)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v1277 = int32(_a_F_RI_FKey_check_19)
	v1278 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[2]))
	v1281 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[2])) = v1281
	v1284 = F_palloc(m, int32(2248))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v1286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+2244)) = v1286
	v1289 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[3]))
	v1294 = F_AllocSetContextCreateInternal(m, v1289, int32(_a_F_RI_FKey_check_20), v1286, int32(1024), int32(_a_F_RI_FKey_check_21))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+2240)) = v1294
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+168))
	if int32(0) < v1297 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1326 = int32(0)
	goto L194
L192:
	;
	v1491 = v1297
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+716)) = v1284
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[2])) = v1278
	v1523 = v1491
	v1526 = v1275
	goto L184
L194:
	;
	v1347 = v1326 << (uint(int32(2)) % 32)
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1275+int32(300)+v1347)))
	v1351 = v1326 << (uint(int32(1)) % 32)
	v1353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1275+int32(236)+v1351))))
	v1354 = F_attnumTypeId(m, v60, v1353)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	v1491 = v1487
	goto L193
L196:
	;
	v1356 = F_ri_HashCompareOp(m, v1349, v1354)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v1358 = int32(0)
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+168))
	if v1359 <= v1358 {
		v1404 = v1358
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1435 = v1404 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1351+(v1284+int32(2176))))) = uint16(v1435)
	v1438 = v1326 * int32(28)
	v1439 = v1284 + int32(896) + v1438
	v1441 = v1356 + int32(40)
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+2240))
	v1443 = *(*int64)(unsafe.Add(mBase, uint32(v1441)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1439)+16)) = v1443
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+24)) = v1445
	v1447 = *(*int64)(unsafe.Add(mBase, uint32(v1441)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1439)+8)) = v1447
	v1449 = *(*int64)(unsafe.Add(mBase, uint32(v1441)))
	*(*int64)(unsafe.Add(mBase, uint32(v1439))) = v1449
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+20)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+16)) = int32(0)
	goto L204
L199:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v175)+192))
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1351+(v1275+int32(172))))))
	v1367 = v1358
	goto L200
L200:
	;
	v1399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1362+int32(48)+v1367<<(uint(int32(1))%32)))))
	if v1399 == v1366 {
		v1404 = v1367
		goto L198
	} else {
		goto L202
	}
L201:
	;
	v1404 = v1359
	goto L198
L202:
	;
	v1402 = v1367 + int32(1)
	if v1402 != v1359 {
		v1367 = v1402
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v1454 = v1438 + v1284
	v1456 = v1356 + int32(12)
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+2240))
	v1458 = *(*int64)(unsafe.Add(mBase, uint32(v1456)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1454)+16)) = v1458
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1456)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+24)) = v1460
	v1462 = *(*int64)(unsafe.Add(mBase, uint32(v1456)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1454)+8)) = v1462
	v1464 = *(*int64)(unsafe.Add(mBase, uint32(v1456)))
	*(*int64)(unsafe.Add(mBase, uint32(v1454))) = v1464
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+20)) = v1457
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+16)) = int32(0)
	goto L205
L205:
	;
	v1470 = F_get_opcode(m, v1349)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1347+(v1284+int32(1792))))) = v1470
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v175)+208))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1404<<(uint(int32(2))%32))))
	F_get_op_opfamily_properties(m, v1349, v1477, int32(0), v1347+(v1284+int32(2048)), v32+int32(2272), v1347+(v1284+int32(1920)))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v1486 = v1326 + int32(1)
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+168))
	if v1486 < v1487 {
		v1326 = v1486
		goto L194
	} else {
		goto L208
	}
L208:
	;
	goto L195
L209:
	;
	v1795 = int32(0)
	F_index_rescan(m, v1270, v32+int32(160), v1766, v1795, v1795)
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L1
	} else {
		goto L234
	}
L210:
	;
	v1555 = int32(0)
	v1557 = v1523
	goto L211
L211:
	;
	v1587 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1526+int32(236)+v1555<<(uint(int32(1))%32)))))
	v1588 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+6)))
	if v1588 < v1587 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	if v1595 <= int32(0) {
		v1766 = v1595
		goto L209
	} else {
		goto L221
	}
L213:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+16))
	m.T0[v1591].(func(*base.Module, int32, int32))(m, v49, v1587)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	v1595 = v1557
	goto L215
L215:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v1598 = v1587 - int32(1)
	v1602 = *(*int64)(unsafe.Add(mBase, uint32(v1596+v1598<<(uint(int32(3))%32))))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1608+v1598))))
	if v1610 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+168))
	v1595 = v1594
	goto L215
L217:
	;
	v1611 = int32(110)
	goto L219
L218:
	;
	v1611 = int32(32)
	goto L219
L219:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(1952)+v1555))) = uint8(v1611)
	*(*int64)(unsafe.Add(mBase, uint32(v32+int32(1984)+v1555<<(uint(int32(3))%32)))) = v1602
	v1620 = v1555 + int32(1)
	if v1620 < v1595 {
		v1555 = v1620
		v1557 = v1595
		goto L211
	} else {
		goto L220
	}
L220:
	;
	goto L212
L221:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+716))
	v1628 = int32(0)
	v1630 = v1595
	goto L222
L222:
	;
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(1952)+v1628))))
	if v1660 == int32(110) {
		v1682 = v1630
		goto L224
	} else {
		goto L225
	}
L223:
	;
	if v1682 <= int32(0) {
		v1766 = v1682
		goto L209
	} else {
		goto L229
	}
L224:
	;
	v1685 = v1628 + int32(1)
	if v1685 < v1682 {
		v1628 = v1685
		v1630 = v1682
		goto L222
	} else {
		goto L228
	}
L225:
	;
	v1665 = v1624 + int32(896) + v1628*int32(28)
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+4))
	if v1666 == int32(0) {
		v1682 = v1630
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v1673 = v32 + int32(1984) + v1628<<(uint(int32(3))%32)
	v1675 = *(*int64)(unsafe.Add(mBase, uint32(v1673)))
	v1678 = F_FunctionCall3Coll(m, v1665, int32(0), v1675, int64(-1), int64(0))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1673))) = v1678
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+168))
	v1682 = v1681
	goto L224
L228:
	;
	goto L223
L229:
	;
	v1698 = int32(0)
	goto L230
L230:
	;
	v1729 = int32(1)
	v1732 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1624+int32(2176)+v1698<<(uint(v1729)%32)))))
	v1734 = v1732 - v1729
	v1739 = int32(2)
	v1740 = v1698 << (uint(v1739) % 32)
	v1742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1624+int32(2048)+v1740))))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1740+(v1624+int32(1920)))))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v175)+248))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1745+v1734<<(uint(v1739)%32))))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1740+(v1624+int32(1792)))))
	v1757 = *(*int64)(unsafe.Add(mBase, uint32(v32+int32(1984)+v1698<<(uint(int32(3))%32))))
	F_ScanKeyEntryInitialize(m, v32+int32(160)+v1734*int32(56), int32(0), v1732, v1742, v1744, v1749, v1751, v1757)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L1
	} else {
		goto L232
	}
L231:
	;
	v1766 = v1762
	goto L209
L232:
	;
	v1761 = v1698 + int32(1)
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+168))
	if v1761 < v1762 {
		v1698 = v1761
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v1800 = F_index_getnext_slot(m, v1270, int32(1), v1092)
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[1])) = v2043
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0])) = v2042
	goto L282
L236:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[1])) = v1994
	*(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[0])) = v1993
	goto L276
L237:
	;
	if v1800 == int32(0) {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[4]))
	v1809 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1811 = int32(0)
	v1813 = int32(1)
	if v1813 < v1805 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1817 = v1813
	goto L242
L241:
	;
	v1817 = int32(3)
	goto L242
L242:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v166)+188))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+104))
	v1822 = m.T0[v1821].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v166, v1092+int32(32), v1087, v1092, v1809, v1811, v1811, v1817, v32+int32(2272))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	if v1822 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	switch v1822 - int32(1) {
	case 0:
		goto L248
	case 1:
		goto L236
	case 2:
		goto L249
	case 3:
		goto L250
	default:
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2288)))
	if v1894 == int32(0) {
		goto L235
	} else {
		goto L266
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L263
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L1
	} else {
		goto L260
	}
L249:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[4]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L256
	}
L250:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, _c_F_RI_FKey_check[4]))
	if v1827 < int32(2) {
		goto L236
	} else {
		goto L251
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_RI_FKey_check_22), int32(0))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(2950), int32(_a_F_RI_FKey_check_23))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	if int32(2) <= v1847 {
		goto L3
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_RI_FKey_check_24), v32+int32(112))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(2964), int32(_a_F_RI_FKey_check_23))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errmsg_internal(m, int32(_a_F_RI_FKey_check_25), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(2976), int32(_a_F_RI_FKey_check_23))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v1822
	F_errmsg_internal(m, int32(_a_F_RI_FKey_check_26), v32+int32(96))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(2980), int32(_a_F_RI_FKey_check_23))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	v1898 = F_BuildIndexInfo(m, v175)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_FormIndexDatum(m, v1898, v1092, int32(0), v32+int32(2272), v32+int32(2240))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v1766 <= int32(0) {
		goto L235
	} else {
		goto L269
	}
L269:
	;
	v1909 = int32(0)
	goto L270
L270:
	;
	v1941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(2240)+v1909))))
	if v1941 != 0 {
		goto L236
	} else {
		goto L272
	}
L271:
	;
	goto L235
L272:
	;
	v1946 = v32 + int32(160) + v1909*int32(56)
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+12))
	v1955 = *(*int64)(unsafe.Add(mBase, uint32(v32+int32(2272)+v1909<<(uint(int32(3))%32))))
	v1956 = *(*int64)(unsafe.Add(mBase, uint32(v1946)+48))
	v1957 = F_FunctionCall2Coll(m, v1946+int32(16), v1949, v1955, v1956)
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	if v1957 == int64(0) {
		goto L236
	} else {
		goto L274
	}
L274:
	;
	v1962 = v1909 + int32(1)
	if v1766 != v1962 {
		v1909 = v1962
		goto L270
	} else {
		goto L275
	}
L275:
	;
	goto L271
L276:
	;
	F_index_endscan(m, v1270)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_ExecDropSingleTupleTableSlot(m, v1092)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	F_UnregisterSnapshot(m, v1087)
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v2007 = int32(0)
	F_ri_ReportViolation(m, v1526, v166, v60, v49, v2007, int32(1), v2007, v2007)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_index_endscan(m, v1270)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_ExecDropSingleTupleTableSlot(m, v1092)
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_UnregisterSnapshot(m, v1087)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_relation_close(m, v175, int32(0))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_relation_close(m, v166, int32(0))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	goto L4
L289:
	;
	F_errmsg(m, int32(_a_F_RI_FKey_check_27), int32(0))
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(2957), int32(_a_F_RI_FKey_check_23))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RI_FKey_restrict_del(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_restrict_del_0), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_restrict_upd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_restrict_upd_0), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			F_errcode(m, int32(16908867))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(_a_F_ri_CheckTrigger_0), v7)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2266), int32(_a_F_ri_CheckTrigger_2))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v12 != int32(448) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				F_errcode(m, int32(16908867))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(_a_F_ri_CheckTrigger_0), v7)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2266), int32(_a_F_ri_CheckTrigger_2))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v15&int32(28) != int32(4) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_errcode(m, int32(16908867))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = l1
						F_errmsg(m, int32(_a_F_ri_CheckTrigger_3), v7-int32(-64))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2275), int32(_a_F_ri_CheckTrigger_2))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v21 = v15 & int32(3)
				switch l2 - int32(2) {
				case 0:
					if v21 == int32(2) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
								F_errmsg(m, int32(_a_F_ri_CheckTrigger_4), v7+int32(32))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2289), int32(_a_F_ri_CheckTrigger_2))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 1:
					if v21 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
								F_errmsg(m, int32(_a_F_ri_CheckTrigger_5), v7+int32(48))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2295), int32(_a_F_ri_CheckTrigger_2))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v7 + int32(80)
						return
					}
				default:
					if v21 == int32(0) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
								F_errmsg(m, int32(_a_F_ri_CheckTrigger_6), v7+int32(16))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2283), int32(_a_F_ri_CheckTrigger_2))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
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
	}
}
func F_ri_FetchPreparedPlan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[0]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v108
L2:
	;
	v55 = v13
	goto L4
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(3092376453124)
	v20 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_0), int64(64), v10, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v56 = int32(0)
	v58 = F_hash_search(m, v55, l0, v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L11
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[1])) = v20
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1698), int64(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(1698), int64(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(51539607560)
	v41 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_1), int64(256), v10, int32(40))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[0])) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(292057776136)
	v50 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_2), int64(256), v10, int32(40))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[2])) = v50
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[0]))
	v55 = v54
	goto L4
L11:
	;
	if v58 == int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if v62 == int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v65 = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v66 == int32(0) {
		v95 = v65
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v95 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v69 <= int32(0) {
		v95 = v65
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v77 = v2
	goto L17
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77<<(uint(int32(2))%32))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+95)))
	if v84 == int32(0) {
		v95 = v84
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v95 = v84
	goto L14
L19:
	;
	v88 = v77 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v88 < v89 {
		v77 = v88
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v108 = v62
	goto L1
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = int32(0)
	F_SPI_freeplan(m, v62)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v108 = v2
	goto L1
}
