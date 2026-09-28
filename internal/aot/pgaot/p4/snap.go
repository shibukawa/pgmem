package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildRestore(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int64
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	v2 = l1
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 == int32(2) {
		v394 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(112)
	return v394
L2:
	;
	v21 = v15 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = m.G0
	v25 = v23 - int32(1120)
	m.G0 = v25
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = int32(_a_F_SnapBuildRestore_0)
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+88)) = uint32(v2)
	v31 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+84)) = uint32(v31)
	v34 = v25 + int32(96)
	v38 = F_pg_sprintf(m, v34, int32(_a_F_SnapBuildRestore_1), v25+int32(80))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if int32(base.Ui32(v43^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		v394 = v3
		goto L1
	} else {
		goto L55
	}
L4:
	;
	return int32(0)
L5:
	;
	v43 = F_OpenTransientFile(m, v34, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L11
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L51
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L47
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L43
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L39
	}
L10:
	;
	m.G0 = v25 + int32(1120)
	goto L3
L11:
	;
	if v43 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestore[0]))
	if v48 == int32(44) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v69 = v25 + int32(96)
	F_fsync_fname(m, v69, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L20
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v34
	F_errmsg(m, int32(_a_F_SnapBuildRestore_8), v25-int32(-64))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1767), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildRestore_0), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_SnapBuildRestoreContents(m, v43, v21, int32(16), v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v80 != int32(1369563137) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v84 = v15 + int32(16)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v85 != int32(6) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v90 = m.Env.Pgmem_crc32c(m, int32(-1), v84, int32(8))
	mBase = m.M
	v92 = v15 + int32(24)
	F_SnapBuildRestoreContents(m, v43, v92, int32(88), v69)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v97 = m.Env.Pgmem_crc32c(m, v90, v92, int32(88))
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v98 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = v98 << (uint(int32(2)) % 32)
	v101 = F_MemoryContextAllocZero(m, v22, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	v108 = v97
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v101
	F_SnapBuildRestoreContents(m, v43, v101, v100, v69)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	v107 = m.Env.Pgmem_crc32c(m, v97, v106, v100)
	mBase = m.M
	v108 = v107
	goto L28
L31:
	;
	v113 = v111 << (uint(int32(2)) % 32)
	v114 = F_MemoryContextAllocZero(m, v22, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v123 = v108
	goto L33
L33:
	;
	v126 = F_CloseTransientFile(m, v43)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v114
	F_SnapBuildRestoreContents(m, v43, v114, v113, v25+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	v122 = m.Env.Pgmem_crc32c(m, v108, v121, v113)
	mBase = m.M
	v123 = v122
	goto L33
L36:
	;
	if v126 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v129 = v123 ^ int32(-1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v129 != v130 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = int32(1369563137)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_10), v25+int32(48))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1788), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_11), v25+int32(32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1794), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_12), v25+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1826), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_13), v25)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1835), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v234 < int32(2) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	if v377 != 0 {
		goto L94
	} else {
		goto L95
	}
L57:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v238 = int32(3)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(base.Ui32(v237) < base.Ui32(v238))|base.B2i32(base.Ui32(v240) < base.Ui32(v238)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v253
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v256
	if v256 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	if int32(0) <= v237-v240 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(v237) < base.Ui32(v240) {
		goto L56
	} else {
		goto L63
	}
L62:
	;
	goto L56
L63:
	;
	goto L58
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v267 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v263
	goto L66
L68:
	;
	F_pfree(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v15)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v272
	v274 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v276 == v274 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L70
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L91
	}
L73:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v295 = F_MemoryContextAllocZero(m, v289, v290<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L78
	}
L74:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+30)))
	if v279 == int32(1) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)+44))
	v284 = v282 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+44)) = v284
	if v284 != 0 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v276)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = int32(5)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v303 = v295 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v306
	v309 = v306 << (uint(int32(2)) % 32)
	if v309 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	base.MemoryCopy(m, v303, v310, v309)
	goto L81
L80:
	;
	goto L81
L81:
	;
	F_pg_qsort(m, v303, v306, int32(4), int32(187))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v316 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v295)+64)) = v316
	*(*int64)(unsafe.Add(mBase, uint32(v295)+44)) = v316
	v320 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+32)) = v320
	*(*int64)(unsafe.Add(mBase, uint32(v295)+20)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v295)+27)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v295
	v327 = int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v295)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+44)) = v328 + v327
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v332)+136)) = v2
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestore[1]))
	if v337 == v327 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v340 = int32(14)
	goto L85
L84:
	;
	v340 = int32(15)
	goto L85
L85:
	;
	v342 = F_errstart(m, v340, int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v342 == int32(0) {
		v394 = v327
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v2)
	v348 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v348)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_2), v15)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v355 = F_errdetail(m, int32(_a_F_SnapBuildRestore_3), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1922), int32(_a_F_SnapBuildRestore_5))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v394 = v327
	goto L1
L91:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildRestore_6), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(348), int32(_a_F_SnapBuildRestore_7))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
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
	F_pfree(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v15)+108))
	if v380 == int32(0) {
		v394 = v3
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	F_pfree(m, v380)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v394 = v3
	goto L1
}
func F_SnapBuildRestoreContents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a_F_SnapBuildRestoreContents_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestoreContents[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(167772216)
	v15 = F_read(m, l0, l1, l2)
	mBase = m.M
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestoreContents[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
	if v15 != l2 {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestoreContents[1]))
		v23 = F_CloseTransientFile(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v15 < int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestoreContents[1])) = v22
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
						F_errmsg(m, int32(_a_F_SnapBuildRestoreContents_1), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SnapBuildRestoreContents_2), int32(1955), int32(_a_F_SnapBuildRestoreContents_3))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
						F_errmsg(m, int32(_a_F_SnapBuildRestoreContents_4), v9+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SnapBuildRestoreContents_2), int32(1961), int32(_a_F_SnapBuildRestoreContents_3))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
	} else {
		m.G0 = v9 + int32(32)
		return
	}
}
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v3 != int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v8 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v8
		if v8 == int32(0) {
			F_pfree(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_SnapBuildSnapDecRefcount_0), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_SnapBuildSnapDecRefcount_1), int32(348), int32(_a_F_SnapBuildSnapDecRefcount_2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
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
