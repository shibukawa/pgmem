package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v20 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CatalogCacheInitializeCache(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l2
	switch l1 - int32(1) {
	case 0:
		v52 = v7
		goto L7
	case 1:
		v45 = v7
		goto L8
	case 2:
		v38 = v7
		goto L9
	case 3:
		goto L10
	default:
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L71
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, l3)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v34 = m.T0[v33].(func(*base.Module, int32) int32)(m, l5)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v38 = base.I32_rotl(v34, int32(24))
	goto L9
L12:
	;
	v45 = base.I32_rotl(v40, int32(16)) ^ v38
	goto L8
L13:
	;
	v52 = base.I32_rotl(v47, int32(8)) ^ v45
	goto L7
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v57 = v52 ^ v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = v57 & (v58 - int32(1))
	v64 = v56 + v61<<(uint(int32(3))%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v66 = int32(0)
	if base.B2i32(v65 == v66)|base.B2i32(v65 == v64) == v66 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	m.G0 = v18 + int32(32)
	return v307
L16:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v81 != v266 {
		goto L62
	} else {
		goto L63
	}
L17:
	;
	v81 = v65
	goto L20
L18:
	;
	goto L19
L19:
	;
	v161 = m.G0
	v163 = v161 - int32(208)
	m.G0 = v163
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = l2
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v171 = F_table_open(m, v169, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L31
	}
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+12)))
	if v89 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v144 != v64 {
		v81 = v144
		goto L20
	} else {
		goto L30
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81-int32(20))))
	if v92 != v57 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v103 = int32(0)
	goto L25
L25:
	;
	v113 = v103 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81-int32(16)+v113)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)+v113)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113+(l0+int32(32)))))
	v122 = m.T0[v121].(func(*base.Module, int32, int32) int32)(m, v115, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L16
L27:
	;
	if v122 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v127 = v103 + int32(1)
	if l1 != v127 {
		v103 = v127
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L21
L31:
	;
	v174 = l1 * int32(48)
	if v174 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	base.MemoryCopy(m, v163+int32(16), l0+int32(104), v174)
	goto L34
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+204)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v163)+156)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v163)+108)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v163)+60)) = l2
	goto L36
L35:
	;
	F_systable_endscan(m, v216)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L54
	}
L36:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v201 - int32(1) {
	case 0, 1:
		v211 = v200
		goto L38
	default:
		goto L39
	case 7, 9, 10, 20:
		goto L40
	case 33:
		goto L41
	}
L37:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0]))
	F_ResourceOwnerEnlarge(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L52
	}
L38:
	;
	v212 = int32(0)
	v216 = F_systable_beginscan(m, v171, v199, v211, v212, l1, v163+int32(16))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L44
	}
L39:
	;
	v211 = int32(1)
	goto L38
L40:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[1])))
	if v207 != int32(1) {
		v211 = v200
		goto L38
	} else {
		goto L43
	}
L41:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[2])))
	if v205 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v211 = v200
	goto L38
L43:
	;
	goto L39
L44:
	;
	v218 = F_systable_getnext(m, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v218 == int32(0) {
		v244 = v212
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v223 = F_CatalogCacheCreateEntry(m, l0, v218, int32(0), v57, v61)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v223 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_systable_endscan(m, v216)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L37
L51:
	;
	goto L36
L52:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+32)) = v233 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0]))
	F_ResourceOwnerRemember(m, v238, v223+int32(40), int32(_a_F_SearchCatCacheInternal_0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v244 = v223
	goto L35
L54:
	;
	F_relation_close(m, v171, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v244 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	m.G0 = v163 + int32(208)
	v307 = v262
	goto L15
L57:
	;
	v252 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[3]))
	if v254 == v252 {
		v262 = v252
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v262 = v244 + int32(40)
	goto L56
L60:
	;
	v258 = F_CatalogCacheCreateEntry(m, l0, int32(0), v163, v57, v61)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v262 = v252
	goto L56
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v273 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+13)))
	if v285 != 0 {
		v307 = int32(0)
		goto L15
	} else {
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
	v277 = v64
	goto L67
L66:
	;
	v277 = v273
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v81
	goto L64
L68:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0]))
	F_ResourceOwnerEnlarge(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v290 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0]))
	v297 = v81 + int32(16)
	F_ResourceOwnerRemember(m, v295, v297, int32(_a_F_SearchCatCacheInternal_0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v307 = v297
	goto L15
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheInternal_1), v18)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_SearchCatCacheInternal_2), int32(373), int32(_a_F_SearchCatCacheInternal_3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
