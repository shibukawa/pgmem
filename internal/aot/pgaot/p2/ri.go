package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_cascade_del(m *base.Module, l0 int32) int32 {
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
	return int32(0)
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
	F_errfinish(m, int32(_a_F_RI_FKey_cascade_del_9), int32(1003), int32(_a_F_RI_FKey_cascade_del_0))
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
	return int32(0)
}
func F_RI_FKey_cascade_upd(m *base.Module, l0 int32) int32 {
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
	return int32(0)
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
	F_errfinish(m, int32(_a_F_RI_FKey_cascade_upd_11), int32(1120), int32(_a_F_RI_FKey_cascade_upd_0))
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
	return int32(0)
}
func F_RI_FKey_check(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	v16 = m.G0
	v18 = v16 - int32(688)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_ri_FetchConstraintInfo(m, v20, v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28&int32(3) == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v18 + int32(688)
	return
L4:
	;
	v33 = int32(28)
	goto L6
L5:
	;
	v33 = int32(24)
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+188))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, v25, v35, int32(_a_F_RI_FKey_check_0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v39 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v46 = F_table_open(m, v44, int32(2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v48 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L114
	}
L11:
	;
	F_relation_close(m, v46, int32(2))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L113
	}
L12:
	;
	v52 = v23 + int32(236)
	v54 = int32(1)
	v56 = int32(0)
	v58 = v48
	v59 = v54
	v61 = v54
	goto L13
L13:
	;
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52+v56<<(uint(int32(1))%32)))))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+6)))
	if v75 < v74 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v86&int32(1) != 0 {
		goto L11
	} else {
		goto L20
	}
L15:
	;
	F_slot_getsomeattrs_int(m, v35, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v80 = v58
	goto L17
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v83 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v74-v83))))
	v86 = v85 & v59
	v89 = (v85 ^ v83) & v61
	v91 = v56 + v83
	if v91 < v80 {
		v56 = v91
		v58 = v80
		v59 = v86
		v61 = v89
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v80 = v79
	goto L17
L19:
	;
	goto L14
L20:
	;
	if v89&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+164)))
	v99 = v97 - int32(102)
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	F_relation_close(m, v46, int32(2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	if v99 == int32(13) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L30
	}
L27:
	;
	goto L23
L28:
	;
	goto L21
L30:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v111 = v23 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v109 + int32(4)
	F_errmsg(m, int32(_a_F_RI_FKey_check_1), v18+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errdetail(m, int32(_a_F_RI_FKey_check_2), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errtableconstraint(m, v43, v111)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(319), int32(_a_F_RI_FKey_check_4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	goto L3
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+684)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+680)) = v139
	v144 = v18 + int32(680)
	v145 = F_ri_FetchPreparedPlan(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v145 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v521 = v145
	goto L41
L40:
	;
	F_initStringInfo(m, v18+int32(664))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v522 = int32(0)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+119)))
	v529 = F_ri_PerformCheck(m, v23, v144, v521, v43, v46, v522, v35, v522, base.B2i32(v525 == int32(112)), int32(5))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L109
	}
L42:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+119)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)+68))
	v154 = F_get_namespace_name(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v156 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+400)) = uint8(v156)
	v160 = v154
	v162 = v18 + int32(400)
	goto L44
L44:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v175 != int32(34) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v192 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+1)) = uint16(v192)
	v195 = v18 + int32(400)
	v196 = F_strlen(m, v195)
	mBase = m.M
	v197 = v196 + v195
	v198 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)) = uint8(v192)
	v207 = v197 + int32(1)
	v209 = v200 + int32(4)
	goto L52
L46:
	;
	goto L45
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v187)
	v160 = v160 + int32(1)
	v162 = v188
	goto L44
L48:
	;
	if v175 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v182 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)) = uint8(v182)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v187 = v184
	v188 = v162 + int32(2)
	goto L47
L51:
	;
	v187 = v175
	v188 = v162 + int32(1)
	goto L47
L52:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v222 != int32(34) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v239 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v207)+1)) = uint16(v239)
	if v152 == int32(112) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v234)
	v207 = v235
	v209 = v209 + int32(1)
	goto L52
L56:
	;
	if v222 == int32(0) {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v229 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)) = uint8(v229)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v234 = v231
	v235 = v207 + int32(2)
	goto L55
L59:
	;
	v234 = v222
	v235 = v207 + int32(1)
	goto L55
L60:
	;
	v245 = int32(_a_F_RI_FKey_check_5)
	goto L62
L61:
	;
	v245 = int32(_a_F_RI_FKey_check_6)
	goto L62
L62:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v246 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if int32(0) < v334 {
		goto L78
	} else {
		goto L79
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+v249<<(uint(int32(1))%32))+170)))
	v254 = F_attnumAttName(m, v46, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v18 + int32(400)
	F_appendStringInfo(m, v18+int32(664), int32(_a_F_RI_FKey_check_7), v18+int32(80))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L77
	}
L67:
	;
	v256 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v256)
	v260 = v254
	v262 = v18 + int32(256)
	goto L68
L68:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v275 != int32(34) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v292 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v262)+1)) = uint16(v292)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v18 + int32(400)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v18 + int32(256)
	F_appendStringInfo(m, v18+int32(664), int32(_a_F_RI_FKey_check_8), v18-int32(-64))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L76
	}
L70:
	;
	goto L69
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v288))) = uint8(v287)
	v260 = v260 + int32(1)
	v262 = v288
	goto L68
L72:
	;
	if v275 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v282 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)) = uint8(v282)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v287 = v284
	v288 = v262 + int32(2)
	goto L71
L75:
	;
	v287 = v275
	v288 = v262 + int32(1)
	goto L71
L76:
	;
	goto L63
L77:
	;
	goto L63
L78:
	;
	v348 = int32(0)
	v350 = int32(_a_F_RI_FKey_check_9)
	goto L81
L79:
	;
	goto L80
L80:
	;
	v458 = v18 + int32(664)
	F_appendStringInfoString(m, v458, int32(_a_F_RI_FKey_check_10))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L98
	}
L81:
	;
	v359 = v348 << (uint(int32(1)) % 32)
	v360 = v23 + int32(172) + v359
	v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v360))))
	v362 = F_attnumTypeId(m, v46, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	v365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v359+v52))))
	v366 = F_attnumTypeId(m, v43, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v360))))
	v369 = F_attnumAttName(m, v46, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v371 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v371)
	v375 = v369
	v377 = v18 + int32(256)
	goto L86
L86:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	if v390 != int32(34) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v407 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v377)+1)) = uint16(v407)
	v410 = v348 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v410
	v413 = v18 + int32(240)
	v417 = F_pg_sprintf(m, v413, int32(_a_F_RI_FKey_check_11), v18+int32(48))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v402)
	v375 = v375 + int32(1)
	v377 = v403
	goto L86
L90:
	;
	if v390 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v397 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)) = uint8(v397)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v402 = v399
	v403 = v377 + int32(2)
	goto L89
L93:
	;
	v402 = v390
	v403 = v377 + int32(1)
	goto L89
L94:
	;
	v420 = v348 << (uint(int32(2)) % 32)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(300)+v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v350
	v425 = v18 + int32(664)
	F_appendStringInfo(m, v425, int32(_a_F_RI_FKey_check_12), v18+int32(32))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_generate_operator_clause(m, v425, v18+int32(256), v362, v422, v413, v366)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(112)+v420))) = v366
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v410 < v440 {
		v348 = v410
		v350 = int32(_a_F_RI_FKey_check_13)
		goto L81
	} else {
		goto L97
	}
L97:
	;
	goto L82
L98:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v462 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v467 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+v463<<(uint(int32(1))%32))+234)))
	v468 = F_attnumTypeId(m, v43, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v18)+664))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v504 = F_ri_PlanCheck(m, v498, v499, v18+int32(112), v18+int32(680), v43, v46)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L108
	}
L102:
	;
	F_appendStringInfoString(m, v458, int32(_a_F_RI_FKey_check_14))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v473
	v476 = v18 + int32(240)
	v480 = F_pg_sprintf(m, v476, int32(_a_F_RI_FKey_check_11), v18+int32(16))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v23)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(_a_F_RI_FKey_check_5)
	F_appendStringInfo(m, v458, int32(_a_F_RI_FKey_check_12), v18)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_generate_operator_clause(m, v458, v476, v468, v482, int32(_a_F_RI_FKey_check_15), int32(_a_F_RI_FKey_check_16))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_appendStringInfoString(m, v458, int32(_a_F_RI_FKey_check_17))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L101
L108:
	;
	v521 = v504
	goto L41
L109:
	;
	v531 = F_SPI_finish(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v531 != int32(2) {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	F_relation_close(m, v46, int32(2))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L3
L113:
	;
	goto L3
L114:
	;
	F_errmsg_internal(m, int32(_a_F_RI_FKey_check_18), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_RI_FKey_check_3), int32(460), int32(_a_F_RI_FKey_check_4))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RI_FKey_restrict_del(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_restrict_upd(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
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
					F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2175), int32(_a_F_ri_CheckTrigger_2))
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
		if v12 != int32(442) {
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
						F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2175), int32(_a_F_ri_CheckTrigger_2))
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
							F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2184), int32(_a_F_ri_CheckTrigger_2))
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
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2198), int32(_a_F_ri_CheckTrigger_2))
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
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2204), int32(_a_F_ri_CheckTrigger_2))
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
									F_errfinish(m, int32(_a_F_ri_CheckTrigger_1), int32(2192), int32(_a_F_ri_CheckTrigger_2))
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
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
	return v103
L2:
	;
	v50 = v13
	goto L4
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(3023656976388)
	v20 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_0), int32(64), v10, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v51 = int32(0)
	v53 = F_hash_search(m, v50, l0, v51, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[1])) = v20
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1485), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(51539607560)
	v36 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_1), int32(256), v10, int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[0])) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(292057776136)
	v45 = F_hash_create(m, int32(_a_F_ri_FetchPreparedPlan_2), int32(256), v10, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[2])) = v45
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ri_FetchPreparedPlan[0]))
	v50 = v49
	goto L4
L10:
	;
	if v53 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v57 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v61 == int32(0) {
		v90 = v60
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v90 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v90 = v60
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v72 = v2
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v72<<(uint(int32(2))%32))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+95)))
	if v79 == int32(0) {
		v90 = v79
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v90 = v79
	goto L13
L18:
	;
	v83 = v72 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v83 < v84 {
		v72 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v103 = v57
	goto L1
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(0)
	F_SPI_freeplan(m, v57)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v103 = v2
	goto L1
}
