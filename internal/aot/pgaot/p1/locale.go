package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_locale_messages(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	goto L3
L1:
	;
	if v144 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v12 + int32(48)
	goto L1
L3:
	;
	goto L6
L6:
	;
	goto L7
L7:
	;
	if l0 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v88 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v75 = F___get_locale(m, int32(5), l0)
	mBase = m.M
	if v75 == int32(-1) {
		v144 = int32(0)
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1100]))
	v88 = v87
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1100])) = v75
	v88 = v75
	goto L20
L25:
	;
	v92 = v88 + int32(8)
	goto L27
L26:
	;
	v92 = int32(541328)
	goto L27
L27:
	;
	v144 = v92
	goto L2
L37:
	;
	v152 = int32(520269)
	goto L45
L38:
	;
	goto L39
L39:
	;
	return
L40:
	;
	goto L39
L41:
	;
	v184 = F___memcpy(m, v179, v152, v162)
	mBase = m.M
	v185 = v179 + v162
	v186 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v186)
	v188 = int32(1)
	v192 = F___memcpy(m, v185+v188, v144, v175+v188)
	mBase = m.M
	v194 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v194 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	goto L40
L43:
	;
	goto L48
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	goto L42
L45:
	;
	v160 = F___strchrnul(m, v152, int32(61))
	mBase = m.M
	if v160 == v152 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v162 = v160 - v152
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_consts[1101]))))
	if v164 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v175 = F_strlen(m, v144)
	mBase = m.M
	v179 = F_emscripten_builtin_malloc(m, v162+v175+int32(2))
	mBase = m.M
	if v179 != 0 {
		goto L41
	} else {
		goto L51
	}
L51:
	;
	goto L42
L52:
	;
	goto L40
L53:
	;
	v232 = v228 << (uint(int32(2)) % 32)
	v234 = v232 + int32(8)
	v236 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	if v236 == v229 {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	v208 = int32(0)
	v209 = v194
	v210 = v198
	goto L60
L55:
	;
	v228 = int32(0)
	v229 = v199
	goto L53
L56:
	;
	v199 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v198 != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v199 = v194
	goto L55
L60:
	;
	v211 = F_strncmp(m, v179, v210, v162+int32(1))
	mBase = m.M
	if v211 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v228 = v219
	v229 = v224
	goto L53
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v179
	F___env_rm_add(m, v214, v179)
	mBase = m.M
	goto L52
L63:
	;
	goto L64
L64:
	;
	v219 = v208 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v220 != 0 {
		v208 = v219
		v209 = v209 + int32(4)
		v210 = v220
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	F_emscripten_builtin_free(m, v179)
	mBase = m.M
	goto L52
L67:
	;
	v251 = v248 + v228<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[474])) = v248
	*(*int32)(unsafe.Add(mBase, _consts[475])) = v248
	if v179 != 0 {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v238 = F_emscripten_builtin_realloc(m, v236, v234)
	mBase = m.M
	if v238 != 0 {
		v248 = v238
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v239 = F_emscripten_builtin_malloc(m, v234)
	mBase = m.M
	if v239 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L71:
	;
	goto L66
L72:
	;
	if v228 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v244 = F___memcpy(m, v239, v243, v232)
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	F_emscripten_builtin_free(m, v246)
	mBase = m.M
	v248 = v239
	goto L67
L76:
	;
	F___env_rm_add(m, int32(0), v179)
	mBase = m.M
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L52
}
func F_check_locale_monetary(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_check_locale(m, int32(4), v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
