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
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v316 int32
	_ = v316
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
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
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
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v20 + int32(80)
	return v23
L4:
	;
	v377 = F_ExecTypeFromTL(m, v367)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L70
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L67
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L64
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l1
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = F_ExecInitExpr(m, v42, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L61
	}
L10:
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
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v121 != int32(1) {
		goto L3
	} else {
		goto L19
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v65 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 == int32(7) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v69 <= int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v76 = int32(0)
	goto L16
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v76<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v95*int32(12)))) = v23
	v101 = v76 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v101 < v102 {
		v76 = v101
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L11
L18:
	;
	goto L17
L19:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v130 = F_AllocSetContextCreateInternal(m, v125, int32(67865), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v130
	v134 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v139 = F_AllocSetContextCreateInternal(m, v134, int32(67834), int32(0), int32(1024), int32(8192))
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
	F_errmsg_internal(m, int32(503567), v20+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(516749), int32(936), int32(294902))
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
		v367 = v3
		v369 = v3
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v207 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v367 = v3
	v369 = v3
	goto L4
L44:
	;
	goto L45
L45:
	;
	v213 = v179
	v218 = v3
	v220 = v3
	v221 = v3
	goto L46
L46:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v221<<(uint(int32(2))%32))))
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
		goto L48
	}
L47:
	;
	v367 = v240
	v369 = v249
	goto L4
L48:
	;
	v240 = F_lappend(m, v218, v238)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
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
		goto L50
	}
L50:
	;
	v249 = F_lappend(m, v220, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
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
		goto L52
	}
L52:
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
		goto L53
	}
L53:
	;
	if v270 == int32(0) {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v275 = F_get_opcode(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
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
		goto L56
	}
L56:
	;
	if v285 == int32(0) {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	F_fmgr_info(m, v289, v259+v198)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	F_fmgr_info(m, v293, v294+v259)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298+v254))) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v303 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v302+v252<<(uint(v303)%32)))) = uint16(v213)
	v310 = v221 + v303
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v310 < v311 {
		v213 = v213 + v303
		v218 = v240
		v220 = v249
		v221 = v310
		goto L46
	} else {
		goto L60
	}
L60:
	;
	goto L47
L61:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v317
	F_errmsg_internal(m, int32(455675), v20)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(516749), int32(827), int32(294902))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
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
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v331
	F_errmsg_internal(m, int32(46205), v20+int32(48))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(516749), int32(989), int32(294902))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v347
	F_errmsg_internal(m, int32(46362), v20+int32(32))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(516749), int32(996), int32(294902))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
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
	v380 = F_ExecInitExtraTupleSlot(m, v29, v377, int32(1646164))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v383 = F_ExecBuildProjectionInfo(m, v367, int32(0), v380, l1, int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v383
	v386 = F_ExecTypeFromTL(m, v369)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v386
	v390 = F_ExecInitExtraTupleSlot(m, v29, v386, int32(1646164))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v395 = F_ExecBuildProjectionInfo(m, v369, v392, v390, v393, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v395
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v403 = F_ExecBuildHash32FromAttrs(m, v377, int32(1646164), v198, v399, v400, v401, l1, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v403
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v410 = F_ExecBuildGroupingEqual(m, v377, v386, int32(1646164), int32(1646268), v177, v408, v203, v409, l1)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v410
	goto L3
}
func F_IsSubTransaction(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
		v11 = F_SimpleLruReadPage_ReadOnly(m, int32(4443928), base.I64_extend_i32_u(v9), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[139]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v18 = int32(2)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v11<<(uint(v18)%32))))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+l0&int32(2047)<<(uint(v18)%32))))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
			v30 = int32(*(*uint16)(unsafe.Add(mBase, _consts[140])))
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
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
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int64
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
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
	*(*int32)(unsafe.Add(mBase, uint32(v418+l2))) = v417
	return
L2:
	;
	v417 = v413
	v418 = int32(8)
	goto L1
L3:
	;
	if v6 == int32(16384) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v211 = int32(16384)
	if v6 == v211 {
		goto L70
	} else {
		goto L71
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
		goto L17
	} else {
		goto L18
	}
L9:
	;
	return
L10:
	;
	v413 = v4
	goto L2
L11:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L69
	}
L12:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L68
	}
L13:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v192 != 0 {
		goto L61
	} else {
		goto L62
	}
L14:
	;
	switch v191 {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L15:
	;
	v191 = v181
	goto L14
L16:
	;
	if v19 <= v51 {
		v86 = v19
		v88 = v20
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v51 = v16
	v55 = v20
	goto L16
L18:
	;
	goto L19
L19:
	;
	v32 = v16
	v36 = v20
	goto L20
L20:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v181 = int32(1)
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v51 = v46
	v55 = v44
	goto L16
L22:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L16
	} else {
		goto L23
	}
L23:
	;
	if v44 < v15 {
		v32 = v46
		v36 = v44
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	if v51 != v86 {
		v127 = v55
		v128 = v88
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v18 <= int32(0) {
		v86 = v19
		v88 = v20
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v67 = v19
	v69 = v20
	goto L28
L28:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v69<<(uint(int32(1))%32)))))
	if v74 != 0 {
		v181 = int32(-1)
		goto L15
	} else {
		goto L30
	}
L29:
	;
	v86 = v78
	v88 = v76
	goto L25
L30:
	;
	v75 = int32(1)
	v76 = v69 + v75
	v78 = v67 - v75
	if v78 <= v51 {
		v86 = v78
		v88 = v76
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v76 < v18 {
		v67 = v78
		v69 = v76
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v15 < v127 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v97 = v55
	v98 = v88
	goto L35
L35:
	;
	if v15 <= v97 {
		v127 = v97
		v128 = v98
		goto L33
	} else {
		goto L37
	}
L36:
	;
	if base.I32_extend16_s(v113) < base.I32_extend16_s(v111) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v18 <= v98 {
		v127 = v97
		v128 = v98
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v102 = int32(1)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v97<<(uint(v102)%32)))))
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98<<(uint(v102)%32)+v17))))
	if v111 == v113 {
		v97 = v97 + v102
		v98 = v98 + v102
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v120 = int32(1)
	goto L42
L41:
	;
	v120 = int32(-1)
	goto L42
L42:
	;
	v191 = v120
	goto L14
L43:
	;
	v131 = v127
	goto L45
L44:
	;
	v131 = v15
	goto L45
L45:
	;
	v138 = v127
	goto L46
L46:
	;
	if v131 == v138 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v181 = v164
	goto L15
L48:
	;
	if v18 < v128 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v164 = int32(1)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v138<<(uint(v164)%32)))))
	if v170 == int32(0) {
		v138 = v138 + v164
		goto L46
	} else {
		goto L60
	}
L51:
	;
	v143 = v128
	goto L53
L52:
	;
	v143 = v18
	goto L53
L53:
	;
	v151 = v128
	goto L54
L54:
	;
	if v143 == v151 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v181 = int32(-1)
	goto L15
L56:
	;
	v191 = int32(0)
	goto L14
L57:
	;
	goto L58
L58:
	;
	v155 = int32(1)
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v151<<(uint(v155)%32)))))
	if v160 == int32(0) {
		v151 = v151 + v155
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	goto L47
L61:
	;
	F_pfree(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v197 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v197
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v202 < v201 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v204 = v201
	goto L67
L66:
	;
	v204 = v202
	goto L67
L67:
	;
	v417 = v204
	v418 = int32(12)
	goto L1
L68:
	;
	v413 = v4
	goto L2
L69:
	;
	v413 = int32(16384)
	goto L2
L70:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v220 = int32(0)
	if base.B2i32(v219 < v216)&base.B2i32(v220 < v215) == v220 {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	goto L72
L72:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L132
	}
L73:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L9
	} else {
		goto L131
	}
L74:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L130
	}
L75:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v392 != 0 {
		goto L123
	} else {
		goto L124
	}
L76:
	;
	switch v391 {
	case 0:
		goto L75
	case 1:
		goto L74
	default:
		goto L73
	}
L77:
	;
	v391 = v381
	goto L76
L78:
	;
	if v219 <= v251 {
		v286 = v219
		v288 = v220
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v251 = v216
	v255 = v220
	goto L78
L80:
	;
	goto L81
L81:
	;
	v232 = v216
	v236 = v220
	goto L82
L82:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214+v236<<(uint(int32(1))%32)))))
	if v242 != 0 {
		v381 = int32(1)
		goto L77
	} else {
		goto L84
	}
L83:
	;
	v251 = v246
	v255 = v244
	goto L78
L84:
	;
	v243 = int32(1)
	v244 = v236 + v243
	v246 = v232 - v243
	if v246 <= v219 {
		v251 = v246
		v255 = v244
		goto L78
	} else {
		goto L85
	}
L85:
	;
	if v244 < v215 {
		v232 = v246
		v236 = v244
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	if v251 != v286 {
		v327 = v255
		v328 = v288
		goto L95
	} else {
		goto L96
	}
L88:
	;
	if v218 <= int32(0) {
		v286 = v219
		v288 = v220
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v267 = v219
	v269 = v220
	goto L90
L90:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217+v269<<(uint(int32(1))%32)))))
	if v274 != 0 {
		v381 = int32(-1)
		goto L77
	} else {
		goto L92
	}
L91:
	;
	v286 = v278
	v288 = v276
	goto L87
L92:
	;
	v275 = int32(1)
	v276 = v269 + v275
	v278 = v267 - v275
	if v278 <= v251 {
		v286 = v278
		v288 = v276
		goto L87
	} else {
		goto L93
	}
L93:
	;
	if v276 < v218 {
		v267 = v278
		v269 = v276
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	if v215 < v327 {
		goto L105
	} else {
		goto L106
	}
L96:
	;
	v297 = v255
	v298 = v288
	goto L97
L97:
	;
	if v215 <= v297 {
		v327 = v297
		v328 = v298
		goto L95
	} else {
		goto L99
	}
L98:
	;
	if base.I32_extend16_s(v313) < base.I32_extend16_s(v311) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if v218 <= v298 {
		v327 = v297
		v328 = v298
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v302 = int32(1)
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214+v297<<(uint(v302)%32)))))
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298<<(uint(v302)%32)+v217))))
	if v311 == v313 {
		v297 = v297 + v302
		v298 = v298 + v302
		goto L97
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v320 = int32(1)
	goto L104
L103:
	;
	v320 = int32(-1)
	goto L104
L104:
	;
	v391 = v320
	goto L76
L105:
	;
	v331 = v327
	goto L107
L106:
	;
	v331 = v215
	goto L107
L107:
	;
	v338 = v327
	goto L108
L108:
	;
	if v331 == v338 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v381 = v364
	goto L77
L110:
	;
	if v218 < v328 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v364 = int32(1)
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214+v338<<(uint(v364)%32)))))
	if v370 == int32(0) {
		v338 = v338 + v364
		goto L108
	} else {
		goto L122
	}
L113:
	;
	v343 = v328
	goto L115
L114:
	;
	v343 = v218
	goto L115
L115:
	;
	v351 = v328
	goto L116
L116:
	;
	if v343 == v351 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v381 = int32(-1)
	goto L77
L118:
	;
	v391 = int32(0)
	goto L76
L119:
	;
	goto L120
L120:
	;
	v355 = int32(1)
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217+v351<<(uint(v355)%32)))))
	if v360 == int32(0) {
		v351 = v351 + v355
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	goto L109
L123:
	;
	F_pfree(m, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v397 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v397
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v397
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v402 < v401 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	v404 = v401
	goto L129
L128:
	;
	v404 = v402
	goto L129
L129:
	;
	v417 = v404
	v418 = int32(12)
	goto L1
L130:
	;
	v413 = v211
	goto L2
L131:
	;
	v413 = int32(0)
	goto L2
L132:
	;
	v413 = v211
	goto L2
}
