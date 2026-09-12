package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateHeapTupleCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
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
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L3
L3:
	;
	if base.B2i32(base.Ui32(v20) < base.Ui32(int32(12000))) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v28 != int32(99) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v31 = F_isTempToastNamespace(m, v28)
	mBase = m.M
	v32 = v31
	goto L8
L7:
	;
	v32 = int32(1)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v33 = m.T0[l3].(func(*base.Module) int32)(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v37 = int32(1)
	if v35 <= int32(2963) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	switch v35 - int32(1249) {
	case 0:
		goto L101
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L1
	case 10:
		goto L102
	default:
		goto L103
	}
L13:
	;
	if v53 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	goto L13
L15:
	;
	v53 = int32(0)
	goto L14
L16:
	;
	if base.Ui32(v35-int32(2608)) < base.Ui32(int32(2)) {
		v53 = v37
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch v35 - int32(3592) {
	case 0, 4:
		v53 = v37
		goto L14
	case 1, 2, 3:
		goto L15
	default:
		goto L22
	}
L19:
	;
	if v35 == int32(1214) {
		v53 = v37
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v35 != int32(2396) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v53 = v37
	goto L14
L22:
	;
	if v35 == int32(2964) {
		v53 = v37
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v56 = int32(1)
	if v35 <= int32(3591) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	goto L26
L26:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[877]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v211 != 0 {
		goto L77
	} else {
		goto L78
	}
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[876]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v129 < v130 {
		goto L55
	} else {
		goto L56
	}
L28:
	;
	goto L27
L29:
	;
	v124 = int32(0)
	goto L28
L30:
	;
	if base.Ui32(v35-int32(2964)) < base.Ui32(int32(4)) {
		v124 = v56
		goto L28
	} else {
		goto L53
	}
L31:
	;
	if v35 <= int32(2670) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v35 <= int32(5999) {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	switch v35 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v124 = v56
		goto L28
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L29
	default:
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v68 = v35 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v68) {
		goto L30
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v35-int32(2396)) {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v124 = v56
	goto L28
L39:
	;
	if int32(1)<<(uint(v68)%32)&int32(226492515) == int32(0) {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v124 = v56
	goto L28
L41:
	;
	if base.Ui32(v35-int32(3592)) < base.Ui32(int32(2)) {
		v124 = v56
		goto L28
	} else {
		goto L51
	}
L42:
	;
	v80 = v35 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v80) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	switch v35 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v124 = v56
		goto L28
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L29
	default:
		goto L47
	}
L45:
	;
	if int32(1)<<(uint(v80)%32)&int32(963) == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v124 = v56
	goto L28
L47:
	;
	if base.Ui32(v35-int32(6000)) < base.Ui32(int32(3)) {
		v124 = v56
		goto L28
	} else {
		goto L48
	}
L48:
	;
	v96 = v35 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v96) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	if int32(1)<<(uint(v96)%32)&int32(49153) != 0 {
		v124 = v56
		goto L28
	} else {
		goto L50
	}
L50:
	;
	goto L29
L51:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v35-int32(4060)) {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v124 = v56
	goto L28
L53:
	;
	if base.Ui32(v35-int32(2846)) < base.Ui32(int32(2)) {
		v124 = v56
		goto L28
	} else {
		goto L54
	}
L54:
	;
	goto L29
L55:
	;
	v132 = v129
	goto L58
L56:
	;
	goto L57
L57:
	;
	if v124 != 0 {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v145 = v128 + v132<<(uint(int32(4))%32)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v146 == int32(251) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v149 == v35 {
		goto L12
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v152 = v132 + int32(1)
	if v152 != v130 {
		v132 = v152
		goto L58
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	goto L59
L65:
	;
	v166 = int32(0)
	goto L67
L66:
	;
	v166 = v126
	goto L67
L67:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[878]))
	if v168 <= v130 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v128 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v190 = v128
	goto L70
L70:
	;
	v194 = v190 + v130<<(uint(int32(4))%32)
	v195 = int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v195)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+13)))
	*(*uint16)(unsafe.Add(mBase, uint32(v194)+1)) = uint16(v197)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+3)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v166
	v204 = v33 + int32(12)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v205 + int32(1)
	goto L12
L71:
	;
	*(*int32)(unsafe.Add(mBase, _consts[878])) = v184
	*(*int32)(unsafe.Add(mBase, _consts[876])) = v185
	v190 = v185
	goto L70
L72:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[73]))
	v176 = F_MemoryContextAlloc(m, v174, int32(512))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L10
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v182 = F_repalloc(m, v128, v168<<(uint(int32(5))%32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L76
	}
L75:
	;
	v184 = int32(32)
	v185 = v176
	goto L71
L76:
	;
	v184 = v168 << (uint(int32(1)) % 32)
	v185 = v182
	goto L71
L77:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v216 = v211
	goto L80
L78:
	;
	goto L79
L79:
	;
	goto L12
L80:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(12))))
	if v226 != v212 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L79
L82:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v265 != 0 {
		v216 = v265
		goto L80
	} else {
		goto L97
	}
L83:
	;
	v229 = v216 - int32(100)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(92))))
	if v232 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_CatalogCacheInitializeCache(m, v229)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v238 = v216 - int32(36)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v240 = F_CatalogCacheComputeTupleHashValue(m, v229, v239, l1)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v245 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216-int32(4)))))
	if v248 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v249 = int32(0)
	goto L91
L90:
	;
	v249 = v245
	goto L91
L91:
	;
	F_RegisterCatcacheInvalidation(m, v242, v240, v249, v33)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	if l2 == int32(0) {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v255 = F_CatalogCacheComputeTupleHashValue(m, v229, v254, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	if v255 == v240 {
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	F_RegisterCatcacheInvalidation(m, v258, v255, v249, v33)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	goto L82
L97:
	;
	goto L81
L98:
	;
	F_RegisterRelcacheInvalidation(m, v33, v325, v324)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L109
	}
L99:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+22)))
	v315 = v313 + v314
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+72)))
	if v316 != int32(102) {
		goto L1
	} else {
		goto L107
	}
L100:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+22)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v307+v308)))
	v312 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v324 = v310
	v325 = v312
	goto L98
L101:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+22)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301+v302)))
	v306 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v324 = v304
	v325 = v306
	goto L98
L102:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+22)))
	v297 = v295 + v296
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+117)))
	if v298 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	switch v35 - int32(2606) {
	case 0:
		goto L99
	default:
		goto L1
	case 4:
		goto L100
	}
L104:
	;
	v299 = int32(0)
	goto L106
L105:
	;
	v299 = v294
	goto L106
L106:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v324 = v300
	v325 = v299
	goto L98
L107:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315)+80))
	if v319 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v324 = v319
	v325 = v323
	goto L98
L109:
	;
	goto L1
}
func F_CacheInvalidateRelcache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = F_PrepareInvalidationState(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v7 != 0 {
			v12 = int32(0)
		} else {
			v12 = v5
		}
		F_RegisterRelcacheInvalidation(m, v9, v12, v8)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_CreateCacheMemoryContext(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v2 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[84]))
		v12 = F_AllocSetContextCreateInternal(m, v7, int32(67938), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[204])) = v12
			return
		}
	} else {
		return
	}
}
func F_cache_multirange_element_properties(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v8 == int32(0) {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		if v11 != int32(109) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
			m.G0 = v6 + int32(16)
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v15 = F_get_multirange_range(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v57
						F_errmsg_internal(m, int32(55417), v6)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(525524), int32(1068), int32(254860))
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
				} else {
					v20 = F_lookup_type_cache(m, v15, int32(2048))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v20
						if v20 == int32(0) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
							m.G0 = v6 + int32(16)
							return
						} else {
							v25 = v20
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
							if v26 == int32(0) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
								m.G0 = v6 + int32(16)
								return
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								v31 = F_lookup_type_cache(m, v29, int32(16400))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
									if v33 != 0 {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v34 | int32(4096)
									} else {
									}
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
									if v38 == int32(0) {
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v41 | int32(8192)
									}
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v25 = v8
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
		if v26 == int32(0) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
			m.G0 = v6 + int32(16)
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v31 = F_lookup_type_cache(m, v29, int32(16400))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
				if v33 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v34 | int32(4096)
				} else {
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
				if v38 == int32(0) {
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v41 | int32(8192)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
