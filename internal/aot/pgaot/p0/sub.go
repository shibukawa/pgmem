package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitSubPlan(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v23 = F_palloc0(m, int32(96))
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
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(392)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l0
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+144))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v39
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L74
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L71
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l1
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = F_ExecInitExpr(m, v42, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L68
	}
L8:
	;
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v45
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+52)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v23)+20)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v23)+68)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v23)+76)) = v49
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v62 == v45 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v121 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v65 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 == int32(7) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v69 <= int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v76 = int32(0)
	goto L14
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v76<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v95*int32(12)))) = v23
	v101 = v76 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v101 < v102 {
		v76 = v101
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	goto L15
L17:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitSubPlan[0]))
	v130 = F_AllocSetContextCreateInternal(m, v125, int32(_a_F_ExecInitSubPlan_0), int32(0), int32(_a_F_ExecInitSubPlan_1), int32(_a_F_ExecInitSubPlan_2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	m.G0 = v20 + int32(80)
	return v23
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v130
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitSubPlan[0]))
	v139 = F_AllocSetContextCreateInternal(m, v134, int32(_a_F_ExecInitSubPlan_3), int32(0), int32(1024), int32(_a_F_ExecInitSubPlan_1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v139
	v142 = F_CreateExprContext(m, v29)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	switch v146 - int32(17) {
	case 0:
		goto L24
	default:
		goto L25
	case 4:
		goto L26
	}
L23:
	;
	if v175 != 0 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v145
	v173 = F_list_make1_impl(m, int32(1), v20+int32(60))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L31
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v149 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v175 = v150
	goto L23
L28:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v156
	F_errmsg_internal(m, int32(_a_F_ExecInitSubPlan_4), v20+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ExecInitSubPlan_5), int32(936), int32(_a_F_ExecInitSubPlan_6))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v175 = v173
	goto L23
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v177 = v176
	goto L34
L33:
	;
	v177 = v3
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v177
	v179 = int32(1)
	v182 = F_palloc(m, v177<<(uint(v179)%32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v182
	v186 = v177 << (uint(int32(2)) % 32)
	v187 = F_palloc(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v187
	v190 = F_palloc(m, v186)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v190
	v194 = v177 * int32(28)
	v195 = F_palloc(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v195
	v198 = F_palloc(m, v194)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v200 = F_palloc(m, v194)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v200
	v203 = F_palloc(m, v186)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v175 == int32(0) {
		v322 = v3
		v323 = v3
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v331 = F_ExecTypeFromTL(m, v322)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L60
	}
L43:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v207 <= int32(0) {
		v322 = v3
		v323 = v3
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v213 = v179
	v219 = v3
	v220 = v3
	v223 = v3
	goto L45
L45:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v223<<(uint(int32(2))%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = base.I32_extend16_s(v213)
	v236 = int32(0)
	v238 = F_makeTargetEntry(m, v234, v235, v236, v236)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v322 = v240
	v323 = v249
	goto L42
L47:
	;
	v240 = F_lappend(m, v219, v238)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v245 = int32(0)
	v247 = F_makeTargetEntry(m, v244, v235, v245, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v249 = F_lappend(m, v220, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v252 = v213 - int32(1)
	v254 = v252 << (uint(int32(2)) % 32)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v203+v254))) = v256
	v259 = v252 * int32(28)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	F_fmgr_info(m, v256, v259+v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v264+v259)+24)) = v231
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v270 = F_get_compatible_hash_operators(m, v267, v20+int32(72))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v270 == int32(0) {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v275 = F_get_opcode(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v277+v254))) = v275
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v285 = F_get_op_hash_functions(m, v280, v20+int32(68), v20-int32(-64))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v285 == int32(0) {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	F_fmgr_info(m, v289, v259+v198)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	F_fmgr_info(m, v293, v294+v259)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298+v254))) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v303 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v302+v252<<(uint(v303)%32)))) = uint16(v213)
	v310 = v223 + v303
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v310 < v311 {
		v213 = v213 + v303
		v219 = v240
		v220 = v249
		v223 = v310
		goto L45
	} else {
		goto L59
	}
L59:
	;
	goto L46
L60:
	;
	v334 = F_ExecInitExtraTupleSlot(m, v29, v331, int32(_a_F_ExecInitSubPlan_7))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v337 = F_ExecBuildProjectionInfo(m, v322, int32(0), v334, l1, int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v337
	v340 = F_ExecTypeFromTL(m, v323)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v340
	v344 = F_ExecInitExtraTupleSlot(m, v29, v340, int32(_a_F_ExecInitSubPlan_7))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v349 = F_ExecBuildProjectionInfo(m, v323, v346, v344, v347, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v349
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v357 = F_ExecBuildHash32FromAttrs(m, v331, int32(_a_F_ExecInitSubPlan_7), v198, v353, v354, v355, l1, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v357
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v364 = F_ExecBuildGroupingEqual(m, v331, v340, int32(_a_F_ExecInitSubPlan_7), int32(_a_F_ExecInitSubPlan_8), v177, v362, v203, v363, l1)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v364
	goto L19
L68:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v392
	F_errmsg_internal(m, int32(_a_F_ExecInitSubPlan_9), v20)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecInitSubPlan_5), int32(827), int32(_a_F_ExecInitSubPlan_6))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v406
	F_errmsg_internal(m, int32(_a_F_ExecInitSubPlan_10), v20+int32(48))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ExecInitSubPlan_5), int32(989), int32(_a_F_ExecInitSubPlan_6))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
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
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v422
	F_errmsg_internal(m, int32(_a_F_ExecInitSubPlan_11), v20+int32(32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ExecInitSubPlan_5), int32(996), int32(_a_F_ExecInitSubPlan_6))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IsSubTransaction(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsSubTransaction[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	return base.B2i32(int32(1) < v3)
}
func F_SubTransGetParent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	if base.Ui32(int32(3)) <= base.Ui32(l0) {
		v9 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
		v11 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_SubTransGetParent_0), base.I64_extend_i32_u(v9), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetParent[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v18 = int32(2)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v11<<(uint(v18)%32))))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+l0&int32(2047)<<(uint(v18)%32))))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
			v30 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_SubTransGetParent[1])))
			v31 = base.I32_rem_u_s(v9, v30)
			F_LWLockRelease(m, v28+v31<<(uint(int32(7))%32))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = v27
				return v37
			}
		}
	} else {
		v37 = int32(0)
		return v37
	}
}
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422+l2))) = v421
	return
L2:
	;
	v421 = v417
	v422 = int32(8)
	goto L1
L3:
	;
	if v6 == int32(_a_F_sub_var_0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v213 = int32(_a_F_sub_var_0)
	if v6 == v213 {
		goto L66
	} else {
		goto L67
	}
L6:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = int32(0)
	if base.B2i32(v19 < v16)&base.B2i32(v20 < v15) == v20 {
		v51 = v16
		v55 = v20
		goto L16
	} else {
		goto L17
	}
L9:
	;
	return
L10:
	;
	v417 = v4
	goto L2
L11:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L65
	}
L12:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L64
	}
L13:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v194 != 0 {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	switch v193 {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L15:
	;
	v193 = v183
	goto L14
L16:
	;
	if base.B2i32(v18 <= int32(0))|base.B2i32(v19 <= v51) != 0 {
		v87 = v19
		v89 = v20
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v32 = v16
	v36 = v20
	goto L18
L18:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v183 = int32(1)
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v51 = v46
	v55 = v44
	goto L16
L20:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if v44 < v15 {
		v32 = v46
		v36 = v44
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	if v51 != v87 {
		v129 = v55
		v130 = v89
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v68 = v19
	v70 = v20
	goto L25
L25:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v70<<(uint(int32(1))%32)))))
	if v75 != 0 {
		v183 = int32(-1)
		goto L15
	} else {
		goto L27
	}
L26:
	;
	v87 = v79
	v89 = v77
	goto L23
L27:
	;
	v76 = int32(1)
	v77 = v70 + v76
	v79 = v68 - v76
	if v79 <= v51 {
		v87 = v79
		v89 = v77
		goto L23
	} else {
		goto L28
	}
L28:
	;
	if v77 < v18 {
		v68 = v79
		v70 = v77
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	if v15 < v129 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v98 = v55
	v99 = v89
	goto L32
L32:
	;
	if base.B2i32(v15 <= v98)|base.B2i32(v18 <= v99) != 0 {
		v129 = v98
		v130 = v99
		goto L30
	} else {
		goto L34
	}
L33:
	;
	if base.I32_extend16_s(v115) < base.I32_extend16_s(v113) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v104 = int32(1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v98<<(uint(v104)%32)))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(v104)%32)+v17))))
	if v113 == v115 {
		v98 = v98 + v104
		v99 = v99 + v104
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v122 = int32(1)
	goto L38
L37:
	;
	v122 = int32(-1)
	goto L38
L38:
	;
	v193 = v122
	goto L14
L39:
	;
	v133 = v129
	goto L41
L40:
	;
	v133 = v15
	goto L41
L41:
	;
	v140 = v129
	goto L42
L42:
	;
	if v133 == v140 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v183 = v166
	goto L15
L44:
	;
	if v18 < v130 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v166 = int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v140<<(uint(v166)%32)))))
	if v172 == int32(0) {
		v140 = v140 + v166
		goto L42
	} else {
		goto L56
	}
L47:
	;
	v145 = v130
	goto L49
L48:
	;
	v145 = v18
	goto L49
L49:
	;
	v153 = v130
	goto L50
L50:
	;
	if v145 == v153 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v183 = int32(-1)
	goto L15
L52:
	;
	v193 = int32(0)
	goto L14
L53:
	;
	goto L54
L54:
	;
	v157 = int32(1)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32)+v17))))
	if v162 == int32(0) {
		v153 = v153 + v157
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	goto L43
L57:
	;
	F_pfree(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v199 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v199
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v204 < v203 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v206 = v203
	goto L63
L62:
	;
	v206 = v204
	goto L63
L63:
	;
	v421 = v206
	v422 = int32(12)
	goto L1
L64:
	;
	v417 = v4
	goto L2
L65:
	;
	v417 = int32(_a_F_sub_var_0)
	goto L2
L66:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v222 = int32(0)
	if base.B2i32(v221 < v218)&base.B2i32(v222 < v217) == v222 {
		v253 = v218
		v257 = v222
		goto L74
	} else {
		goto L75
	}
L67:
	;
	goto L68
L68:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L124
	}
L69:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L123
	}
L70:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L9
	} else {
		goto L122
	}
L71:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v396 != 0 {
		goto L115
	} else {
		goto L116
	}
L72:
	;
	switch v395 {
	case 0:
		goto L71
	case 1:
		goto L70
	default:
		goto L69
	}
L73:
	;
	v395 = v385
	goto L72
L74:
	;
	if base.B2i32(v220 <= int32(0))|base.B2i32(v221 <= v253) != 0 {
		v289 = v221
		v291 = v222
		goto L81
	} else {
		goto L82
	}
L75:
	;
	v234 = v218
	v238 = v222
	goto L76
L76:
	;
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v238<<(uint(int32(1))%32)))))
	if v244 != 0 {
		v385 = int32(1)
		goto L73
	} else {
		goto L78
	}
L77:
	;
	v253 = v248
	v257 = v246
	goto L74
L78:
	;
	v245 = int32(1)
	v246 = v238 + v245
	v248 = v234 - v245
	if v248 <= v221 {
		v253 = v248
		v257 = v246
		goto L74
	} else {
		goto L79
	}
L79:
	;
	if v246 < v217 {
		v234 = v248
		v238 = v246
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v253 != v289 {
		v331 = v257
		v332 = v291
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v270 = v221
	v272 = v222
	goto L83
L83:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219+v272<<(uint(int32(1))%32)))))
	if v277 != 0 {
		v385 = int32(-1)
		goto L73
	} else {
		goto L85
	}
L84:
	;
	v289 = v281
	v291 = v279
	goto L81
L85:
	;
	v278 = int32(1)
	v279 = v272 + v278
	v281 = v270 - v278
	if v281 <= v253 {
		v289 = v281
		v291 = v279
		goto L81
	} else {
		goto L86
	}
L86:
	;
	if v279 < v220 {
		v270 = v281
		v272 = v279
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	if v217 < v331 {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v300 = v257
	v301 = v291
	goto L90
L90:
	;
	if base.B2i32(v217 <= v300)|base.B2i32(v220 <= v301) != 0 {
		v331 = v300
		v332 = v301
		goto L88
	} else {
		goto L92
	}
L91:
	;
	if base.I32_extend16_s(v317) < base.I32_extend16_s(v315) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v306 = int32(1)
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v300<<(uint(v306)%32)))))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301<<(uint(v306)%32)+v219))))
	if v315 == v317 {
		v300 = v300 + v306
		v301 = v301 + v306
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v324 = int32(1)
	goto L96
L95:
	;
	v324 = int32(-1)
	goto L96
L96:
	;
	v395 = v324
	goto L72
L97:
	;
	v335 = v331
	goto L99
L98:
	;
	v335 = v217
	goto L99
L99:
	;
	v342 = v331
	goto L100
L100:
	;
	if v335 == v342 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v385 = v368
	goto L73
L102:
	;
	if v220 < v332 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v368 = int32(1)
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v342<<(uint(v368)%32)))))
	if v374 == int32(0) {
		v342 = v342 + v368
		goto L100
	} else {
		goto L114
	}
L105:
	;
	v347 = v332
	goto L107
L106:
	;
	v347 = v220
	goto L107
L107:
	;
	v355 = v332
	goto L108
L108:
	;
	if v347 == v355 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v385 = int32(-1)
	goto L73
L110:
	;
	v395 = int32(0)
	goto L72
L111:
	;
	goto L112
L112:
	;
	v359 = int32(1)
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355<<(uint(v359)%32)+v219))))
	if v364 == int32(0) {
		v355 = v355 + v359
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	goto L101
L115:
	;
	F_pfree(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v401 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v401
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v401
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v406 < v405 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v408 = v405
	goto L121
L120:
	;
	v408 = v406
	goto L121
L121:
	;
	v421 = v408
	v422 = int32(12)
	goto L1
L122:
	;
	v417 = v213
	goto L2
L123:
	;
	v417 = int32(0)
	goto L2
L124:
	;
	v417 = v213
	goto L2
}
