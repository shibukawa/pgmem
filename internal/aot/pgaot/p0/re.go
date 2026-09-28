package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
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
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v329 int64
	_ = v329
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	v11 = m.G0
	v13 = v11 - int32(176)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v45 = int32(1)
	if v15&v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v21 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v24 = int32(16)
	goto L7
L6:
	;
	v24 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = int32(4)
	goto L10
L9:
	;
	v31 = v24
	goto L10
L10:
	;
	v44 = v31
	goto L1
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	v49 = v45
	goto L14
L13:
	;
	v49 = int32(4)
	goto L14
L14:
	;
	v50 = l0 + v49
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[0]))
	if v51 < v53 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v353 = v13 + int32(16)
	F_pg_regerror(m, v231, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L53
	} else {
		goto L89
	}
L16:
	;
	m.G0 = v13 + int32(176)
	return int32(_a_F_RE_compile_and_cache_0)
L17:
	;
	v56 = v51
	goto L20
L18:
	;
	goto L19
L19:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[1]))
	if v194 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L20:
	;
	v67 = v56 * int32(52)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[2])))
	if v68 != v44 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v181 = v56 + int32(1)
	if v181 != v53 {
		v56 = v181
		goto L20
	} else {
		goto L49
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[3])))
	if v72 != l1 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[4])))
	if v74 != l2 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[5])))
	if base.Ui32(int32(4)) <= base.Ui32(v44) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v138 != 0 {
		goto L22
	} else {
		goto L44
	}
L27:
	;
	v138 = int32(0)
	goto L26
L28:
	;
	v112 = v107
	v113 = v108
	v114 = v109
	goto L38
L29:
	;
	if (v76|v50)&int32(3) != 0 {
		v107 = v76
		v108 = v50
		v109 = v44
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v100 = v76
	v101 = v50
	v102 = v44
	goto L31
L31:
	;
	if v102 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v84 = v76
	v85 = v50
	v86 = v44
	goto L33
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v89 != v90 {
		v107 = v84
		v108 = v85
		v109 = v86
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v100 = v95
	v101 = v93
	v102 = v97
	goto L31
L35:
	;
	v92 = int32(4)
	v93 = v85 + v92
	v95 = v84 + v92
	v97 = v86 - v92
	if base.Ui32(int32(3)) < base.Ui32(v97) {
		v84 = v95
		v85 = v93
		v86 = v97
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v107 = v100
	v108 = v101
	v109 = v102
	goto L28
L38:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 == v118 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v138 = v117 - v118
	goto L26
L40:
	;
	v120 = int32(1)
	v125 = v114 - v120
	if v125 != 0 {
		v112 = v112 + v120
		v113 = v113 + v120
		v114 = v125
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	if v56 == int32(0) {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+168)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[7])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+152)) = v145
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+136)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_RE_compile_and_cache[10])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+120)) = v153
	if v67 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, int32(_a_F_RE_compile_and_cache_1), int32(_a_F_RE_compile_and_cache_2), v67)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[6])) = v159
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v13)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[7])) = v162
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[8])) = v165
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v13)+144))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[9])) = v168
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v13)+136))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[4])) = v171
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[2])) = v174
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[10])) = v177
	goto L16
L49:
	;
	goto L21
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[11]))
	v204 = F_AllocSetContextCreateInternal(m, v199, int32(_a_F_RE_compile_and_cache_3), int32(0), int32(1024), int32(_a_F_RE_compile_and_cache_4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v211 = v44 + int32(1)
	v212 = F_palloc_mul(m, int32(4), v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L53
	} else {
		goto L55
	}
L53:
	;
	return int32(0)
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[1])) = v204
	goto L52
L55:
	;
	v214 = F_pg_mb2wchar_with_len(m, v50, v212, v44)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[12]))
	v222 = F_AllocSetContextCreateInternal(m, v217, int32(_a_F_RE_compile_and_cache_5), int32(0), int32(1024), int32(_a_F_RE_compile_and_cache_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v224 = int32(_a_F_RE_compile_and_cache_6)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[12])) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v222
	v231 = F_pg_regcomp(m, v13+int32(140), v212, v214, l1, l2)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v212)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	if v231 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	v235 = F_palloc(m, v211)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v235
	if v44 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v235, v50, v44)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v239+v44))) = uint8(v241)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+36)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v44
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[0]))
	if int32(32) <= v250 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v255 = v250 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[0])) = v255
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255*int32(52))+uint32(_c_F_RE_compile_and_cache[10])))
	F_MemoryContextDelete(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L53
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[1]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v269 != v265 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L67
L69:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[0]))
	if v299 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L70:
	;
	if v269 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	if v265 != 0 {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v263)+28))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	if v274 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v273 == int32(0) {
		goto L73
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = v273
	goto L75
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+20)) = v273
	goto L75
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v273)+24)) = v279
	goto L73
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v265
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+28)) = v286
	if v286 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v263)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = int32(0)
	goto L72
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+24)) = v263
	goto L85
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+20)) = v263
	goto L69
L86:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[6])) = v311
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v13)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[7])) = v314
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[8])) = v317
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v13)+144))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[9])) = v320
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v13)+136))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[4])) = v323
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[2])) = v326
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[10])) = v329
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[12])) = v225
	*(*int32)(unsafe.Add(mBase, _c_F_RE_compile_and_cache[0])) = v299 + int32(1)
	goto L16
L87:
	;
	v303 = v299 * int32(52)
	if v303 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	base.MemoryCopy(m, int32(_a_F_RE_compile_and_cache_1), int32(_a_F_RE_compile_and_cache_2), v303)
	goto L86
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L53
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L53
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v353
	F_errmsg(m, int32(_a_F_RE_compile_and_cache_7), v13)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L53
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_RE_compile_and_cache_8), int32(223), int32(_a_F_RE_compile_and_cache_9))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L53
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReScanExprContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = int32(_a_F_ReScanExprContext_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReScanExprContext[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ReScanExprContext[0])) = v8
	v11 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_MemoryContextReset(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L10
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	m.T0[v16].(func(*base.Module, int64))(m, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReScanExprContext[0])) = v6
	goto L3
L6:
	;
	return
L7:
	;
	F_pfree(m, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v21 != 0 {
		v11 = v21
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	return
}
