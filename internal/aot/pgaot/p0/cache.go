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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v93 int32
	_ = v93
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[0]))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L3
L3:
	;
	if base.B2i32(base.Ui32(v15) < base.Ui32(int32(_a_F_CacheInvalidateHeapTupleCommon_0))) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	if v21 != int32(99) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v24 = F_isTempToastNamespace(m, v21)
	mBase = m.M
	v26 = v24
	goto L8
L7:
	;
	v26 = int32(1)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v27 = m.T0[l3].(func(*base.Module) int32)(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = int32(1)
	if v29 <= int32(2963) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	switch v29 - int32(1249) {
	case 0:
		goto L97
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L1
	case 10:
		goto L98
	default:
		goto L99
	}
L13:
	;
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	goto L13
L15:
	;
	v48 = int32(0)
	goto L14
L16:
	;
	if base.B2i32(v29 == int32(1214))|base.B2i32(base.Ui32(v29-int32(2608)) < base.Ui32(int32(2))) != 0 {
		v48 = v31
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch v29 - int32(3592) {
	case 0, 4:
		v48 = v31
		goto L14
	case 1, 2, 3:
		goto L15
	default:
		goto L21
	}
L19:
	;
	if v29 != int32(2396) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v48 = v31
	goto L14
L21:
	;
	if v29 == int32(2964) {
		v48 = v31
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v51 = int32(1)
	if v29 <= int32(3591) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	goto L25
L25:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[1]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v201 != 0 {
		goto L73
	} else {
		goto L74
	}
L26:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[3]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v127 < v128 {
		goto L51
	} else {
		goto L52
	}
L27:
	;
	goto L26
L28:
	;
	v122 = int32(0)
	goto L27
L29:
	;
	if base.B2i32(base.Ui32(v29-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v29-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v122 = v51
		goto L27
	} else {
		goto L50
	}
L30:
	;
	if v29 <= int32(2670) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v29 <= int32(_a_F_CacheInvalidateHeapTupleCommon_1) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	switch v29 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v122 = v51
		goto L27
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L28
	default:
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v63 = v29 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v63))|base.B2i32(int32(1)<<(uint(v63)%32)&int32(226492515) == int32(0)) != 0 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v29-int32(2396)) {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v122 = v51
	goto L27
L38:
	;
	v122 = v51
	goto L27
L39:
	;
	if base.Ui32(v29-int32(3592)) < base.Ui32(int32(2)) {
		v122 = v51
		goto L27
	} else {
		goto L48
	}
L40:
	;
	v76 = v29 - int32(_a_F_CacheInvalidateHeapTupleCommon_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v76))|base.B2i32(int32(1)<<(uint(v76)%32)&int32(963) == int32(0)) != 0 {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	switch v29 - int32(_a_F_CacheInvalidateHeapTupleCommon_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v122 = v51
		goto L27
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L28
	default:
		goto L44
	}
L43:
	;
	v122 = v51
	goto L27
L44:
	;
	if base.Ui32(v29-int32(_a_F_CacheInvalidateHeapTupleCommon_4)) < base.Ui32(int32(3)) {
		v122 = v51
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v93 = v29 - int32(_a_F_CacheInvalidateHeapTupleCommon_5)
	if base.Ui32(int32(15)) < base.Ui32(v93) {
		goto L28
	} else {
		goto L46
	}
L46:
	;
	if int32(1)<<(uint(v93)%32)&int32(_a_F_CacheInvalidateHeapTupleCommon_6) != 0 {
		v122 = v51
		goto L27
	} else {
		goto L47
	}
L47:
	;
	goto L28
L48:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v29-int32(4060)) {
		goto L28
	} else {
		goto L49
	}
L49:
	;
	v122 = v51
	goto L27
L50:
	;
	goto L28
L51:
	;
	v130 = v127
	goto L54
L52:
	;
	goto L53
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[4]))
	if v162 <= v128 {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	v142 = v126 + v130<<(uint(int32(4))%32)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v143 == int32(251) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	if v146 == v29 {
		goto L12
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v149 = v130 + int32(1)
	if v149 != v128 {
		v130 = v149
		goto L54
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L55
L61:
	;
	if v126 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v184 = v126
	goto L63
L63:
	;
	v188 = v184 + v128<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v29
	if v122 != 0 {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[4])) = v178
	*(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[3])) = v179
	v184 = v179
	goto L63
L65:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[5]))
	v170 = F_MemoryContextAlloc(m, v168, int32(512))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v176 = F_repalloc(m, v126, v162<<(uint(int32(5))%32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L10
	} else {
		goto L69
	}
L68:
	;
	v178 = int32(32)
	v179 = v170
	goto L64
L69:
	;
	v178 = v162 << (uint(int32(1)) % 32)
	v179 = v176
	goto L64
L70:
	;
	v191 = int32(0)
	goto L72
L71:
	;
	v191 = v124
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v191
	v193 = int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v195 + int32(1)
	goto L12
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v206 = v201
	goto L76
L74:
	;
	goto L75
L75:
	;
	goto L12
L76:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v206-int32(12))))
	if v215 != v202 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v254 != 0 {
		v206 = v254
		goto L76
	} else {
		goto L93
	}
L79:
	;
	v218 = v206 - int32(100)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v206-int32(92))))
	if v221 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_CatalogCacheInitializeCache(m, v218)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v227 = v206 - int32(36)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v229 = F_CatalogCacheComputeTupleHashValue(m, v218, v228, l1)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L10
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(4)))))
	if v237 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v238 = int32(0)
	goto L87
L86:
	;
	v238 = v234
	goto L87
L87:
	;
	F_RegisterCatcacheInvalidation(m, v231, v229, v238, v27)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	if l2 == int32(0) {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v244 = F_CatalogCacheComputeTupleHashValue(m, v218, v243, l2)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	if v229 == v244 {
		goto L78
	} else {
		goto L91
	}
L91:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	F_RegisterCatcacheInvalidation(m, v247, v244, v238, v27)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	goto L78
L93:
	;
	goto L77
L94:
	;
	F_RegisterRelcacheInvalidation(m, v27, v312, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L10
	} else {
		goto L105
	}
L95:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+22)))
	v302 = v300 + v301
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+72)))
	if v303 != int32(102) {
		goto L1
	} else {
		goto L103
	}
L96:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+22)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294+v295)))
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v311 = v297
	v312 = v299
	goto L94
L97:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+22)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288+v289)))
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v311 = v291
	v312 = v293
	goto L94
L98:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+22)))
	v284 = v282 + v283
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+117)))
	if v285 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	switch v29 - int32(2606) {
	case 0:
		goto L95
	default:
		goto L1
	case 4:
		goto L96
	}
L100:
	;
	v286 = int32(0)
	goto L102
L101:
	;
	v286 = v281
	goto L102
L102:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v311 = v287
	v312 = v286
	goto L94
L103:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)+80))
	if v306 == int32(0) {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateHeapTupleCommon[2]))
	v311 = v306
	v312 = v310
	goto L94
L105:
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelcache[0]))
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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCacheMemoryContext[0]))
	if v2 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCacheMemoryContext[1]))
		v12 = F_AllocSetContextCreateInternal(m, v7, int32(_a_F_CreateCacheMemoryContext_0), int32(0), int32(_a_F_CreateCacheMemoryContext_1), int32(_a_F_CreateCacheMemoryContext_2))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_CreateCacheMemoryContext[0])) = v12
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
						F_errmsg_internal(m, int32(_a_F_cache_multirange_element_properties_0), v6)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_cache_multirange_element_properties_1), int32(1068), int32(_a_F_cache_multirange_element_properties_2))
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
								v31 = F_lookup_type_cache(m, v29, int32(_a_F_cache_multirange_element_properties_3))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
									if v33 != 0 {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v34 | int32(_a_F_cache_multirange_element_properties_4)
									} else {
									}
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
									if v38 == int32(0) {
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v41 | int32(_a_F_cache_multirange_element_properties_5)
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
			v31 = F_lookup_type_cache(m, v29, int32(_a_F_cache_multirange_element_properties_3))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
				if v33 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v34 | int32(_a_F_cache_multirange_element_properties_4)
				} else {
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
				if v38 == int32(0) {
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v41 | int32(_a_F_cache_multirange_element_properties_5)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v46 | int32(512)
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
