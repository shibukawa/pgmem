package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regc_ctype_get_cache(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[1]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v32 = F_emscripten_builtin_malloc(m, int32(40))
	mBase = m.M
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v17 != l0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v24 != 0 {
		v14 = v24
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 != v10 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	return v14 + int32(8)
L9:
	;
	goto L5
L10:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v305)
	mBase = m.M
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v307)
	mBase = m.M
	F_emscripten_builtin_free(m, v32)
	mBase = m.M
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = l0
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
	v39 = int32(512)
	v40 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+20)) = int64(274877906944)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v40
	v45 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v45
	v47 = int32(0)
	if base.B2i32(v40 == v47)|base.B2i32(v45 == v47) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v297 = int32(0)
	goto L13
L13:
	;
	return v297
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = l1
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v68 = v32 + int32(8)
	v69 = int32(0)
	v72 = v69
	v74 = v69
	goto L22
L16:
	;
	v63 = int32(128)
	goto L18
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[2]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	goto L19
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(-1)
	v66 = v63
	goto L15
L19:
	;
	if v58 == int32(6) {
		v66 = int32(2048)
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(256)
	goto L18
L21:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v243 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L22:
	;
	v77 = m.T0[l0].(func(*base.Module, int32) int32)(m, v74)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v162 <= int32(0) {
		goto L21
	} else {
		goto L52
	}
L24:
	;
	v164 = v74 + int32(1)
	if v164 != v66 {
		v72 = v162
		v74 = v164
		goto L22
	} else {
		goto L51
	}
L25:
	;
	return int32(0)
L26:
	;
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v162 = v72 + int32(1)
	goto L24
L28:
	;
	goto L29
L29:
	;
	if v72 <= int32(0) {
		v162 = v72
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v85 = v74 - v72
	if base.Ui32(int32(2)) <= base.Ui32(v72) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v155 == int32(0) {
		goto L10
	} else {
		goto L49
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v90 < v91 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v126 < v127 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v108 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v107+v106<<(uint(v108)%32)))) = v85
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v118 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v112+v113<<(uint(v108)%32))+4)) = v85 + v72 - v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v121 + v118
	v155 = v118
	goto L31
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v106 = v90
	v107 = v93
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v91 << (uint(int32(1)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v100 = F_emscripten_builtin_realloc(m, v97, v91<<(uint(int32(4))%32))
	mBase = m.M
	if v100 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v155 = int32(0)
	goto L31
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v100
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v106 = v105
	v107 = v100
	goto L35
L42:
	;
	v144 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v142 + v144
	*(*int32)(unsafe.Add(mBase, uint32(v143+v142<<(uint(int32(2))%32)))) = v85
	v155 = v144
	goto L31
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v142 = v126
	v143 = v129
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v127 << (uint(int32(1)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v136 = F_emscripten_builtin_realloc(m, v133, v127<<(uint(int32(3))%32))
	mBase = m.M
	if v136 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v155 = int32(0)
	goto L31
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v136
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v142 = v141
	v143 = v136
	goto L42
L49:
	;
	v160 = v74 + int32(1)
	if v160 != v66 {
		v72 = int32(0)
		v74 = v160
		goto L22
	} else {
		goto L50
	}
L50:
	;
	goto L21
L51:
	;
	goto L23
L52:
	;
	v168 = v66 - v162
	if base.Ui32(int32(2)) <= base.Ui32(v162) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v238 == int32(0) {
		goto L10
	} else {
		goto L71
	}
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v173 < v174 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v209 < v210 {
		goto L65
	} else {
		goto L66
	}
L57:
	;
	v191 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v190+v189<<(uint(v191)%32)))) = v168
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v201 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v195+v196<<(uint(v191)%32))+4)) = v168 + v162 - v201
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v204 + v201
	v238 = v201
	goto L53
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v189 = v173
	v190 = v176
	goto L57
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v174 << (uint(int32(1)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v183 = F_emscripten_builtin_realloc(m, v180, v174<<(uint(int32(4))%32))
	mBase = m.M
	if v183 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v238 = int32(0)
	goto L53
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v183
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v189 = v188
	v190 = v183
	goto L57
L64:
	;
	v227 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v225 + v227
	*(*int32)(unsafe.Add(mBase, uint32(v226+v225<<(uint(int32(2))%32)))) = v168
	v238 = v227
	goto L53
L65:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v225 = v209
	v226 = v212
	goto L64
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v210 << (uint(int32(1)) % 32)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v219 = F_emscripten_builtin_realloc(m, v216, v210<<(uint(int32(3))%32))
	mBase = m.M
	if v219 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v238 = int32(0)
	goto L53
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v219
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v225 = v224
	v226 = v219
	goto L64
L71:
	;
	goto L21
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v264 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v258
	goto L72
L74:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v246)
	mBase = m.M
	v248 = int32(0)
	v258 = v248
	v259 = v248
	goto L73
L75:
	;
	goto L76
L76:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v250 <= v243 {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v255 = F_emscripten_builtin_realloc(m, v252, v243<<(uint(int32(2))%32))
	mBase = m.M
	if v255 == int32(0) {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	v258 = v255
	v259 = v243
	goto L73
L79:
	;
	v285 = int32(_a_F_regc_ctype_get_cache_0)
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v286
	*(*int32)(unsafe.Add(mBase, _c_F_regc_ctype_get_cache[0])) = v32
	v297 = v68
	goto L13
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v279
	goto L79
L81:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v267)
	mBase = m.M
	v269 = int32(0)
	v279 = v269
	v280 = v269
	goto L80
L82:
	;
	goto L83
L83:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v271 <= v264 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v276 = F_emscripten_builtin_realloc(m, v273, v264<<(uint(int32(3))%32))
	mBase = m.M
	if v276 == int32(0) {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v279 = v276
	v280 = v264
	goto L80
}
