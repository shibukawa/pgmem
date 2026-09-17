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
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int64
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	v2 = l1
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 == int32(2) {
		v376 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L93
	}
L2:
	;
	m.G0 = v15 + int32(112)
	return v376
L3:
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(base.Ui32(v43^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		v376 = v3
		goto L2
	} else {
		goto L56
	}
L5:
	;
	return int32(0)
L6:
	;
	v43 = F_OpenTransientFile(m, v34, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L52
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L48
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L44
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L40
	}
L11:
	;
	m.G0 = v25 + int32(1120)
	goto L4
L12:
	;
	if v43 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildRestore[0]))
	if v48 == int32(44) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v69 = v25 + int32(96)
	F_fsync_fname(m, v69, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v34
	F_errmsg(m, int32(_a_F_SnapBuildRestore_8), v25-int32(-64))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1763), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_fsync_fname(m, int32(_a_F_SnapBuildRestore_0), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_SnapBuildRestoreContents(m, v43, v21, int32(16), v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v80 != int32(1369563137) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v84 = v15 + int32(16)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v85 != int32(6) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v90 = m.Env.Pgmem_crc32c(m, int32(-1), v84, int32(8))
	mBase = m.M
	v92 = v15 + int32(24)
	F_SnapBuildRestoreContents(m, v43, v92, int32(88), v69)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v97 = m.Env.Pgmem_crc32c(m, v90, v92, int32(88))
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v98 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v100 = v98 << (uint(int32(2)) % 32)
	v101 = F_MemoryContextAllocZero(m, v22, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	v108 = v97
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v101
	F_SnapBuildRestoreContents(m, v43, v101, v100, v69)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	v107 = m.Env.Pgmem_crc32c(m, v97, v106, v100)
	mBase = m.M
	v108 = v107
	goto L29
L32:
	;
	v113 = v111 << (uint(int32(2)) % 32)
	v114 = F_MemoryContextAllocZero(m, v22, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	v123 = v108
	goto L34
L34:
	;
	v126 = F_CloseTransientFile(m, v43)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v114
	F_SnapBuildRestoreContents(m, v43, v114, v113, v25+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	v122 = m.Env.Pgmem_crc32c(m, v108, v121, v113)
	mBase = m.M
	v123 = v122
	goto L34
L37:
	;
	if v126 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v129 = v123 ^ int32(-1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v129 != v130 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L11
L40:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = int32(1369563137)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_10), v25+int32(48))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1784), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_11), v25+int32(32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1790), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_12), v25+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1822), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25 + int32(96)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_13), v25)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1831), int32(_a_F_SnapBuildRestore_9))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v234 < int32(2) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	if v359 != 0 {
		goto L87
	} else {
		goto L88
	}
L58:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v238))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v237)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v250 != 0 {
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v250 = base.B2i32(base.Ui32(v237) < base.Ui32(v238))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v250 = int32(base.Ui32(v237-v238) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v259
	if v259 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v270 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v266
	goto L66
L68:
	;
	F_pfree(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v15)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v275
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v279 == v277 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v298 = F_MemoryContextAllocZero(m, v292, v293<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L77
	}
L73:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+30)))
	if v282 == int32(1) {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+44))
	v287 = v285 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+44)) = v287
	if v287 != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	F_pfree(m, v279)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = int32(5)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v306 = v298 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v298)+12)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v298)+8)) = v304
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+16)) = v309
	v312 = v309 << (uint(int32(2)) % 32)
	if v312 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	base.MemoryCopy(m, v306, v313, v312)
	goto L80
L79:
	;
	goto L80
L80:
	;
	F_pg_qsort(m, v306, v309, int32(4), int32(185))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v319 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v298)+64)) = v319
	*(*int64)(unsafe.Add(mBase, uint32(v298)+44)) = v319
	v323 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v298)+32)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v298)+20)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v298)+27)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v298
	v330 = int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v298)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+44)) = v331 + v330
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v335)+136)) = v2
	v339 = F_errstart(m, int32(15), v323)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if v339 == int32(0) {
		v376 = v330
		goto L2
	} else {
		goto L83
	}
L83:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v2)
	v345 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v345)
	F_errmsg(m, int32(_a_F_SnapBuildRestore_2), v15)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errdetail(m, int32(_a_F_SnapBuildRestore_3), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(1918), int32(_a_F_SnapBuildRestore_5))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v376 = v330
	goto L2
L87:
	;
	F_pfree(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v15)+108))
	if v362 == int32(0) {
		v376 = v3
		goto L2
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v362)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v376 = v3
	goto L2
L93:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildRestore_6), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_SnapBuildRestore_4), int32(344), int32(_a_F_SnapBuildRestore_7))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(167772214)
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
							F_errfinish(m, int32(_a_F_SnapBuildRestoreContents_2), int32(1951), int32(_a_F_SnapBuildRestoreContents_3))
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
							F_errfinish(m, int32(_a_F_SnapBuildRestoreContents_2), int32(1957), int32(_a_F_SnapBuildRestoreContents_3))
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
				F_errfinish(m, int32(_a_F_SnapBuildSnapDecRefcount_1), int32(344), int32(_a_F_SnapBuildSnapDecRefcount_2))
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
