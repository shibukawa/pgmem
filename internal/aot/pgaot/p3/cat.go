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
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L73
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
	if v65 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	m.G0 = v18 + int32(32)
	return v304
L16:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v78 != v262 {
		goto L64
	} else {
		goto L65
	}
L17:
	;
	v158 = m.G0
	v160 = v158 - int32(208)
	m.G0 = v160
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = l2
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v168 = F_table_open(m, v166, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L32
	}
L18:
	;
	if v65 == v64 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v78 = v65
	goto L20
L20:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+12)))
	if v86 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v141 != v64 {
		v78 = v141
		goto L20
	} else {
		goto L30
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78-int32(20))))
	if v89 != v57 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v104 = int32(0)
	goto L25
L25:
	;
	v110 = v104 << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v78-int32(16)+v110)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)+v110)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110+(l0+int32(32)))))
	v119 = m.T0[v118].(func(*base.Module, int32, int32) int32)(m, v112, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L16
L27:
	;
	if v119 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v124 = v104 + int32(1)
	if l1 != v124 {
		v104 = v124
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
	v304 = v256
	goto L15
L32:
	;
	v175 = l1 * int32(48)
	if v175 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+204)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v160)+156)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v160)+108)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v160)+60)) = l2
	goto L37
L34:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, v160+int32(16), l0+int32(104), v175)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v198 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v199 - int32(1) {
	case 0, 1:
		v209 = v198
		goto L39
	default:
		goto L40
	case 7, 9, 10, 20:
		goto L41
	case 33:
		goto L42
	}
L39:
	;
	v210 = int32(0)
	v214 = F_systable_beginscan(m, v168, v197, v209, v210, l1, v160+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L46
	}
L40:
	;
	v209 = int32(1)
	goto L39
L41:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1072])))
	if v205 != int32(1) {
		v209 = v198
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1073])))
	if v203 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v209 = v198
	goto L39
L44:
	;
	goto L40
L45:
	;
	F_systable_endscan(m, v214)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L63
	}
L46:
	;
	v216 = F_systable_getnext(m, v214)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v216 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v219 = F_CatalogCacheCreateEntry(m, l0, v216, int32(0), v57, v61)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v238 = v210
	goto L50
L50:
	;
	F_systable_endscan(m, v214)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L55
	}
L51:
	;
	if v219 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_ResourceOwnerEnlarge(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+32)) = v227 + int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_ResourceOwnerRemember(m, v232, v219+int32(40), int32(1771908))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v238 = v219
	goto L50
L55:
	;
	F_sequence_close(m, v168, int32(1))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if v238 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v160 + int32(208)
	goto L31
L58:
	;
	v246 = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v248 == v246 {
		v256 = v246
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v256 = v238 + int32(40)
	goto L57
L61:
	;
	v252 = F_CatalogCacheCreateEntry(m, l0, int32(0), v160, v57, v61)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v256 = v246
	goto L57
L63:
	;
	goto L37
L64:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v269 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+13)))
	if v280 != 0 {
		v304 = int32(0)
		goto L15
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
	v273 = v64
	goto L69
L68:
	;
	v273 = v269
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v78
	goto L66
L70:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_ResourceOwnerEnlarge(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v286 = v78 + int32(8)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v287 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	v294 = v78 + int32(16)
	F_ResourceOwnerRemember(m, v292, v294, int32(1771908))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v304 = v294
	goto L15
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	F_errmsg_internal(m, int32(503315), v18)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(522694), int32(373), int32(364197))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
