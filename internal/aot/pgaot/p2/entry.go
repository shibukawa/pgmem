package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryBeginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v361 int32
	_ = v361
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	v5 = l4
	v9 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16400)
	m.G0 = v20
	if l1 < v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v40 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v41 == v40 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(l1^int32(-1))<<(uint(int32(2))%32))))
	v39 = v31
	goto L1
L3:
	;
	goto L4
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v39 = v33 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(2))%32)+v39)+20))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v48&int32(32767))+6)))
	v61 = (v52&int32(8191)+int32(7))&int32(16376) | int32(4)
	goto L7
L6:
	;
	v61 = v9
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v64 = int32(4)
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+14)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
	v67 = v65 - v66
	if v67 <= v64 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L20
	} else {
		goto L70
	}
L9:
	;
	if base.Ui32(v70-int32(4)+v61) < base.Ui32((v63&int32(8191)+int32(7))&int32(16376)|int32(4)) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v70 = v64
	goto L12
L11:
	;
	v70 = v67
	goto L12
L12:
	;
	goto L9
L13:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	if l1 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v361 = v40
	goto L15
L15:
	;
	m.G0 = v20 + int32(16400)
	return v361
L16:
	;
	v119 = F_PageGetTempPageCopy(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L20
	} else {
		goto L23
	}
L17:
	;
	v89 = (l1 ^ int32(-1)) << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v91)))
	v94 = F_PageGetTempPageCopy(m, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v103 = l1 << (uint(int32(13)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v109 = F_PageGetTempPageCopy(m, v103+v105+int32(-8192))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L22
	}
L20:
	;
	return int32(0)
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99+v89)))
	v117 = v94
	v118 = v101
	goto L16
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v117 = v109
	v118 = v112 + v103 + int32(-8192)
	goto L16
L23:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+19)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v122 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_PageIndexTupleDelete(m, v117, v83)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v5 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v148 = int32(0)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v149) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+v129)+6)))
	if v131&int32(2) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v83<<(uint(int32(2))%32)+v117)+20))
	v140 = v117 + v137&int32(32767)
	v141 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v141)
	*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
	v145 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v145)
	goto L28
L31:
	;
	v250 = v121 << (uint(int32(8)) % 32)
	v252 = v157 + int32(1)
	if v252&int32(65535) == v83 {
		goto L52
	} else {
		goto L53
	}
L32:
	;
	v157 = int32(base.Ui32(v149+int32(262120)) >> (uint(int32(2)) % 32))
	goto L34
L33:
	;
	v157 = v148
	goto L34
L34:
	;
	v159 = v157 & int32(65535)
	if v159 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v233 = v148
	v241 = v20 + int32(16)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v169 = v148
	v177 = v20 + int32(16)
	v178 = v40
	goto L38
L38:
	;
	v186 = v178 & int32(65535)
	if v83 == v186 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v233 = v226
	v241 = v223
	goto L31
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+6)))
	v195 = (v189&int32(8191) + int32(7)) & int32(16376)
	if v195 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v202 = v169
	v204 = v177
	goto L42
L42:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v186<<(uint(int32(2))%32)+(v117+int32(24))-int32(4))))
	v213 = v117 + v210&int32(32767)
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+6)))
	v220 = (v214&int32(8191) + int32(7)) & int32(16376)
	if v220 != 0 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v202 = v169 + v195 + int32(4)
	v204 = v197 + v195
	goto L42
L44:
	;
	v196 = F__emscripten_memcpy_bulkmem(m, v177, v188, v195)
	mBase = m.M
	v197 = v196
	goto L46
L45:
	;
	v197 = v177
	goto L46
L46:
	;
	goto L43
L47:
	;
	v223 = v222 + v220
	v226 = v202 + v220 + int32(4)
	v228 = v178 + int32(1)
	if base.Ui32(v228&int32(65535)) <= base.Ui32(v159) {
		v169 = v226
		v177 = v223
		v178 = v228
		goto L38
	} else {
		goto L51
	}
L48:
	;
	v221 = F__emscripten_memcpy_bulkmem(m, v204, v213, v220)
	mBase = m.M
	v222 = v221
	goto L50
L49:
	;
	v222 = v204
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L39
L52:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+6)))
	v263 = (v257&int32(8191) + int32(7)) & int32(16376)
	if v263 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v269 = v233
	goto L54
L54:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+v271)+6)))
	F_PageInit(m, v119, v250, int32(8))
	mBase = m.M
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+16)))
	v277 = v119 + v276
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+6)) = uint16(v273)
	goto L59
L55:
	;
	v269 = v233 + v263 + int32(4)
	goto L54
L56:
	;
	v264 = F__emscripten_memcpy_bulkmem(m, v241, v256, v263)
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+16)))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119+v281)+6)))
	F_PageInit(m, v117, v250, int32(8))
	mBase = m.M
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v287 = v117 + v286
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v287)+6)) = uint16(v283)
	goto L60
L60:
	;
	v291 = int32(1)
	v298 = v20 + int32(16)
	v299 = v117
	v305 = int32(0)
	v307 = v291
	goto L61
L61:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+6)))
	v316 = v314 & int32(8191)
	if base.Ui32(int32(base.Ui32(v269)>>(uint(v291)%32))) < base.Ui32(v305) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v117
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v119
	v361 = int32(2)
	goto L15
L63:
	;
	v327 = int32(0)
	v329 = F_PageAddItemExtended(m, v325, v298, v316, v327, v327)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L20
	} else {
		goto L67
	}
L64:
	;
	v325 = v119
	v326 = v305
	goto L63
L65:
	;
	goto L66
L66:
	;
	v325 = v299
	v326 = v305 + (v316+int32(7))&int32(16376) + int32(4)
	goto L63
L67:
	;
	if v329 == int32(0) {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+6)))
	v342 = v307 + int32(1)
	v343 = int32(65535)
	if base.Ui32(v342&v343) <= base.Ui32(v252&v343) {
		v298 = v298 + (v333&int32(8191)+int32(7))&int32(16376)
		v299 = v325
		v305 = v326
		v307 = v342
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v377 + int32(4)
	F_errmsg_internal(m, int32(748991), v20)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(525626), int32(689), int32(430008))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_entry_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 float64
	_ = v228
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v236 float64
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 float64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int64
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
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
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v470 float64
	_ = v470
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int64
	_ = v485
	var v486 int64
	_ = v486
	var v494 int64
	_ = v494
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+412))
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[1460]))
	if v94 <= v91 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+376))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+364))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+352))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+340))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+328))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+316))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+304))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+292))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+280))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+268))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+256))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+244))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+232))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+220))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+208))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+196))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+184))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+160))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+148))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)+136))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+124))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+112))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(-64))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v91 = v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v55 + (v56 + (v57 + (v58 + (v59 + v25))))))))))))))))))))))))))))))
	goto L4
L3:
	;
	v91 = v25
	goto L4
L4:
	;
	goto L1
L5:
	;
	goto L8
L6:
	;
	goto L7
L7:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v455 = F_hash_search(m, v451, l0, int32(1), v17+int32(12))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L14
	} else {
		goto L58
	}
L8:
	;
	v110 = int32(4735020)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+412))
	if v115 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v183 = F_palloc(m, v180<<(uint(int32(2))%32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+376))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+364))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+352))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+340))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+328))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+316))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)+304))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v113)+292))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+280))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v113)+268))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v113)+256))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v113)+244))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v113)+232))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v113)+220))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+208))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v113)+196))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v113)+184))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v113)+172))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)+160))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v113)+148))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v113)+136))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v113)+124))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v113)+112))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v113)+100))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v113)+88))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(-64))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v113)+52))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v113)+40))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v180 = v116 + (v117 + (v118 + (v119 + (v120 + (v121 + (v122 + (v123 + (v124 + (v125 + (v126 + (v127 + (v128 + (v129 + (v130 + (v131 + (v132 + (v133 + (v134 + (v135 + (v136 + (v137 + (v138 + (v139 + (v140 + (v141 + (v144 + (v145 + (v146 + (v147 + (v148 + v114))))))))))))))))))))))))))))))
	goto L13
L12:
	;
	v180 = v114
	goto L13
L13:
	;
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	F_hash_seq_init(m, v17+int32(12), v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v192 = int32(0)
	v197 = F_hash_seq_search(m, v17+int32(12))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v183)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L14
	} else {
		goto L48
	}
L18:
	;
	if v197 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pg_qsort(m, v183, int32(0), int32(4), int32(7702))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v215 = v197
	v216 = v192
	v219 = v192
	v220 = v192
	goto L23
L22:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+16)) = int32(1024)
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183+v216<<(uint(int32(2))%32)))) = v215
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v215)+256))
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v215)+24))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v215)+32))
	if v231 == int64(0)-v233 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_pg_qsort(m, v183, v240, int32(4), int32(7702))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L14
	} else {
		goto L33
	}
L25:
	;
	v236 = float64(0.5)
	goto L27
L26:
	;
	v236 = float64(0.99)
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v215)+256)) = base.F64_mul(v228, v236)
	v240 = v216 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+396))
	v242 = int32(-1)
	v246 = v219 + int32(base.Ui32(v241^v242)>>(uint(int32(31))%32))
	if v241 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v250 = v242
	goto L30
L29:
	;
	v250 = v241
	goto L30
L30:
	;
	v253 = v220 + v250 + int32(1)
	v256 = F_hash_seq_search(m, v17+int32(12))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	if v256 != 0 {
		v215 = v256
		v216 = v240
		v219 = v246
		v220 = v253
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v183+v240<<(uint(int32(1))%32)&int32(-4))))
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v269)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v263)+8)) = v270
	if v246 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v272 = base.I32_div_u_s(v253, v246)
	v274 = v272
	goto L36
L35:
	;
	v274 = int32(1024)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v274
	v277 = int32(10)
	if base.Ui32(v277) <= base.Ui32(v240) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v280 = v277
	goto L39
L38:
	;
	v280 = v240
	goto L39
L39:
	;
	v282 = base.I32_div_u_s(v240, int32(20))
	if base.Ui32(v216) < base.Ui32(int32(199)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v285 = v280
	goto L42
L41:
	;
	v285 = v282
	goto L42
L42:
	;
	if v285 == int32(0) {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v293 = int32(0)
	goto L44
L44:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v304 = int32(2)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v183+v293<<(uint(v304)%32))))
	v310 = F_hash_search(m, v303, v307, v304, int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L14
	} else {
		goto L46
	}
L45:
	;
	goto L17
L46:
	;
	v313 = v293 + int32(1)
	if v313 != v285 {
		v293 = v313
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+20)) = int32(1)
	if v333 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	F_s_lock(m, v339+int32(20), int32(518750), int32(2205), int32(513234))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L14
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+20)) = int32(0)
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v353)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v353)+40)) = v356 + int64(1)
	v362 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v364)+412))
	if v366 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _consts[1460]))
	if v434 <= v431 {
		goto L8
	} else {
		goto L57
	}
L54:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+376))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)+364))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v364)+352))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+340))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v364)+328))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v364)+316))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v364)+304))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v364)+292))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v364)+280))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v364)+268))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v364)+256))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v364)+244))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v364)+232))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v364)+220))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v364)+208))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v364)+196))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v364)+184))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v364)+172))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v364)+160))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v364)+148))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v364)+136))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v364)+124))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v364)+112))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v364)+100))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v364)+88))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v364)+76))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v364-int32(-64))))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v364)+52))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v364)+40))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v364)+28))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v364)+16))
	v431 = v367 + (v368 + (v369 + (v370 + (v371 + (v372 + (v373 + (v374 + (v375 + (v376 + (v377 + (v378 + (v379 + (v380 + (v381 + (v382 + (v383 + (v384 + (v385 + (v386 + (v387 + (v388 + (v389 + (v390 + (v391 + (v392 + (v395 + (v396 + (v397 + (v398 + (v399 + v365))))))))))))))))))))))))))))))
	goto L56
L55:
	;
	v431 = v365
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L9
L58:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v457 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v465 = F__emscripten_memset_bulkmem(m, v455+int32(24), base.I32_extend8_s(int32(0)), int32(368))
	mBase = m.M
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v17 + int32(32)
	return v455
L62:
	;
	if l4 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v469 = *(*float64)(unsafe.Add(mBase, uint32(v468)+8))
	v470 = v469
	goto L65
L64:
	;
	v470 = float64(1)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+424)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v455)+256)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v455)+400)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v455)+396)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v455)+392)) = l1
	v480 = m.G0
	v481 = int32(16)
	v482 = v480 - v481
	m.G0 = v482
	F___gettimeofday(m, v482)
	mBase = m.M
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
	v486 = int64(*(*int32)(unsafe.Add(mBase, uint32(v482)+8)))
	m.G0 = v482 + v481
	v494 = v486 + v485*int64(1000000) - int64(946684800000000)
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v455)+416)) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v455)+408)) = v494
	goto L61
}
func F_entry_reset(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v119 int64
	_ = v119
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int64
	_ = v266
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L104
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = F_LWLockAcquire(m, v25, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+412))
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v105 = m.G0
	v106 = int32(16)
	v107 = v105 - v106
	m.G0 = v107
	F___gettimeofday(m, v107)
	mBase = m.M
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
	v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v107)+8)))
	m.G0 = v107 + v106
	v119 = v111 + v110*int64(1000000) - int64(946684800000000)
	goto L10
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+376))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+364))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+352))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+340))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+328))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+316))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+304))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+292))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+280))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+268))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+256))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)+244))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+232))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+220))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+196))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+184))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+172))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+160))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+148))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+136))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+124))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v34)+112))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34-int32(-64))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v101 = v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + (v65 + (v66 + (v67 + (v68 + (v69 + v35))))))))))))))))))))))))))))))
	goto L9
L8:
	;
	v101 = v35
	goto L9
L9:
	;
	goto L6
L10:
	;
	if l0 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	if v296 == v101 {
		goto L66
	} else {
		goto L67
	}
L12:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	F_hash_seq_init(m, v12+int32(24), v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L29
	}
L13:
	;
	if l1 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if l2 == int64(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v138 = F_hash_search(m, v133, v12+int32(24), v131, v131)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v161)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v167 = int32(0)
	v169 = F_hash_search(m, v164, v12+int32(24), v167, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	if v138 == int32(0) {
		v160 = v131
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v142 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v138)+56)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v138)+416)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v138)+80)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v138)+72)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v138-int32(-64)))) = v142
	v160 = v131
	goto L16
L20:
	;
	goto L21
L21:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v157 = F_hash_search(m, v154, v138, int32(2), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v160 = int32(1)
	goto L16
L23:
	;
	if v169 == int32(0) {
		v296 = v160
		goto L11
	} else {
		goto L24
	}
L24:
	;
	if l3 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v173 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v169)+56)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v169)+416)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v169)+80)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v169)+72)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v169-int32(-64)))) = v173
	v296 = v160
	goto L11
L26:
	;
	goto L27
L27:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v188 = F_hash_search(m, v185, v169, int32(2), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v296 = v160 + int32(1)
	goto L11
L29:
	;
	v200 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v203 = int32(0)
	if base.B2i32(l0|l1 == v203)&base.B2i32(l2 == int64(0)) == v203 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v210 = int32(0)
	if v200 == v210 {
		v296 = v210
		goto L11
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v254 = int32(0)
	if v200 == v254 {
		v296 = v254
		goto L11
	} else {
		goto L56
	}
L34:
	;
	v217 = v200
	v219 = v210
	goto L35
L35:
	;
	if l0 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v296 = v249
	goto L11
L37:
	;
	v252 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L54
	}
L38:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v222 != l0 {
		v249 = v219
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if l1 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v224 != l1 {
		v249 = v219
		goto L37
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if l2 != int64(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v217)+8))
	if v228 != l2 {
		v249 = v219
		goto L37
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if l3 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v230 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v217)+56)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v217)+416)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v217)+80)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v217)+72)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v217-int32(-64)))) = v230
	v249 = v219
	goto L37
L51:
	;
	goto L52
L52:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v245 = F_hash_search(m, v242, v217, int32(2), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v249 = v219 + int32(1)
	goto L37
L54:
	;
	if v252 != 0 {
		v217 = v252
		v219 = v249
		goto L35
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	v261 = v200
	v263 = v254
	goto L57
L57:
	;
	if l3 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v296 = v285
	goto L11
L59:
	;
	v288 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L64
	}
L60:
	;
	v266 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v261)+56)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v261)+416)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v261)+80)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v261)+72)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v261-int32(-64)))) = v266
	v285 = v263
	goto L59
L61:
	;
	goto L62
L62:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v281 = F_hash_search(m, v278, v261, int32(2), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v285 = v263 + int32(1)
	goto L59
L64:
	;
	if v288 != 0 {
		v261 = v288
		v263 = v285
		goto L57
	} else {
		goto L65
	}
L65:
	;
	goto L58
L66:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v300)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+20)) = int32(1)
	if v302 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v436 = v300
	goto L68
L68:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	F_LWLockRelease(m, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L103
	}
L69:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	F_s_lock(m, v308+int32(20), int32(518750), int32(2749), int32(112913))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int64)(unsafe.Add(mBase, uint32(v322)+48)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v322)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v322)+20)) = int32(0)
	v332 = F_AllocateFile(m, int32(119677), int32(34101))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+24)) = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v407)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+20)) = int32(1)
	if v410 != 0 {
		goto L99
	} else {
		goto L100
	}
L74:
	;
	if v332 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v338 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v332)+76))
	if v359 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	if v338 == int32(0) {
		goto L73
	} else {
		goto L79
	}
L79:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(119677)
	F_errmsg(m, int32(314401), v12)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(518750), int32(2764), int32(112913))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	goto L73
L83:
	;
	v402 = F_FreeFile(m, v332)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L98
	}
L84:
	;
	v373 = F_ftruncate(m, v371, int64(0))
	mBase = m.M
	if v373 == int32(0) {
		goto L83
	} else {
		goto L92
	}
L85:
	;
	if v364 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v332)+60))
	v364 = v362
	goto L85
L87:
	;
	goto L88
L88:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v332)+60))
	v364 = v363
	goto L85
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(8)
	v371 = int32(-1)
	goto L91
L90:
	;
	v371 = v364
	goto L91
L91:
	;
	goto L84
L92:
	;
	v378 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	if v378 == int32(0) {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(119677)
	F_errmsg(m, int32(314432), v12+int32(16))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(518750), int32(2773), int32(112913))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	goto L83
L98:
	;
	goto L73
L99:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	F_s_lock(m, v416+int32(20), int32(518750), int32(2780), int32(112913))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+20)) = int32(0)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+32)) = v431 + int32(1)
	v436 = v428
	goto L68
L102:
	;
	goto L101
L103:
	;
	m.G0 = v12 + int32(48)
	return v119
L104:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(724373), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(518750), int32(2687), int32(112913))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
