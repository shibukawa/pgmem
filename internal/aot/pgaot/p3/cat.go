package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int32 {
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
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
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
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = l5
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = l2
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
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L82
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = m.T0[v53].(func(*base.Module, int64) int32)(m, l2)
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
	v47 = m.T0[v46].(func(*base.Module, int64) int32)(m, l3)
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
	v40 = m.T0[v39].(func(*base.Module, int64) int32)(m, l4)
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
	v34 = m.T0[v33].(func(*base.Module, int64) int32)(m, l5)
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
	m.G0 = v18 + int32(48)
	return v345
L16:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v80 != v302 {
		goto L73
	} else {
		goto L74
	}
L17:
	;
	v80 = v65
	goto L20
L18:
	;
	goto L19
L19:
	;
	v161 = m.G0
	v163 = v161 - int32(256)
	m.G0 = v163
	*(*int64)(unsafe.Add(mBase, uint32(v163)+24)) = l5
	*(*int64)(unsafe.Add(mBase, uint32(v163)+16)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = l2
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
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+52)))
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
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v144 != v64 {
		v80 = v144
		goto L20
	} else {
		goto L30
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v90 != v57 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v102 = int32(0)
	goto L25
L25:
	;
	v111 = v102 << (uint(int32(3)) % 32)
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v80+int32(16)+v111)))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(16)+v111)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(32)+v102<<(uint(int32(2))%32))))
	v122 = m.T0[v121].(func(*base.Module, int64, int64) int32)(m, v113, v117)
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
	v127 = v102 + int32(1)
	if l1 != v127 {
		v102 = v127
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
	v174 = l1 * int32(56)
	if v174 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	base.MemoryCopy(m, v163+int32(32), l0+int32(104), v174)
	goto L34
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v163)+248)) = l5
	*(*int64)(unsafe.Add(mBase, uint32(v163)+192)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v163)+136)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v163)+80)) = l2
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v187 - int32(1) {
	case 0, 1:
		v197 = v184
		goto L37
	default:
		goto L38
	case 7, 9, 10, 20, 42, 43:
		goto L39
	case 33:
		goto L40
	}
L35:
	;
	F_systable_endscan(m, v273)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L65
	}
L36:
	;
	v201 = F_systable_beginscan(m, v171, v185, v197, int32(0), l1, v163+int32(32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v197 = int32(1)
	goto L37
L39:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0])))
	if v193 != int32(1) {
		v197 = v184
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[1])))
	if v191 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v197 = v184
	goto L37
L42:
	;
	goto L38
L43:
	;
	v203 = F_systable_getnext(m, v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	if v203 == int32(0) {
		v273 = v201
		v274 = v184
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v214 = v201
	v218 = v203
	goto L46
L46:
	;
	v223 = F_CatalogCacheCreateEntry(m, l0, v218, int32(0), v57, v61)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[2]))
	F_ResourceOwnerEnlarge(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L63
	}
L48:
	;
	if v223 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_systable_endscan(m, v214)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L47
L52:
	;
	v229 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v232 - int32(1) {
	case 0, 1:
		v242 = v229
		goto L54
	default:
		goto L55
	case 7, 9, 10, 20, 42, 43:
		goto L56
	case 33:
		goto L57
	}
L53:
	;
	v246 = F_systable_beginscan(m, v171, v230, v242, int32(0), l1, v163+int32(32))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v242 = int32(1)
	goto L54
L56:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[0])))
	if v238 != int32(1) {
		v242 = v229
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[1])))
	if v236 != 0 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v242 = v229
	goto L54
L59:
	;
	goto L55
L60:
	;
	v248 = F_systable_getnext(m, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v248 != 0 {
		v214 = v246
		v218 = v248
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v273 = v246
	v274 = v229
	goto L35
L63:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+48)) = v254 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[2]))
	F_ResourceOwnerRemember(m, v259, base.I64_extend_i32_u(v223+int32(56)), int32(_a_F_SearchCatCacheInternal_0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v273 = v214
	v274 = v223
	goto L35
L65:
	;
	F_relation_close(m, v171, int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	if v274 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	m.G0 = v163 + int32(256)
	v345 = v298
	goto L15
L68:
	;
	v288 = int32(0)
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[3]))
	if v290 == v288 {
		v298 = v288
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v298 = v274 + int32(56)
	goto L67
L71:
	;
	v294 = F_CatalogCacheCreateEntry(m, l0, int32(0), v163, v57, v61)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v298 = v288
	goto L67
L73:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v305
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v309 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+53)))
	if v321 != 0 {
		v345 = int32(0)
		goto L15
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
	v313 = v64
	goto L78
L77:
	;
	v313 = v309
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v80
	goto L75
L79:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[2]))
	F_ResourceOwnerEnlarge(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+48)) = v326 + int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheInternal[2]))
	v333 = v80 + int32(56)
	F_ResourceOwnerRemember(m, v331, base.I64_extend_i32_u(v333), int32(_a_F_SearchCatCacheInternal_0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v345 = v333
	goto L15
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheInternal_1), v18)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_SearchCatCacheInternal_2), int32(385), int32(_a_F_SearchCatCacheInternal_3))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
