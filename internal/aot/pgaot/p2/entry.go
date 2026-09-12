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
	F_errmsg_internal(m, int32(743898), v20)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(522150), int32(689), int32(427267))
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
