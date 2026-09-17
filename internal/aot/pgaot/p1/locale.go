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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
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
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	goto L3
L1:
	;
	if v138 != 0 {
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
	if v84 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v75 = F___get_locale(m, int32(5), l0)
	mBase = m.M
	if v75 == int32(-1) {
		v138 = int32(0)
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[0]))
	v84 = v83
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[0])) = v75
	v84 = v75
	goto L20
L25:
	;
	v88 = v84 + int32(8)
	goto L27
L26:
	;
	v88 = int32(_a_F_assign_locale_messages_0)
	goto L27
L27:
	;
	v138 = v88
	goto L2
L37:
	;
	v146 = int32(_a_F_assign_locale_messages_1)
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
	v177 = F___memcpy(m, v172, v146, v155)
	mBase = m.M
	v178 = v172 + v155
	v179 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v179)
	v181 = int32(1)
	v185 = F___memcpy(m, v178+v181, v138, v168+v181)
	mBase = m.M
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[1]))
	if v187 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[2])) = int32(28)
	goto L42
L45:
	;
	v153 = F___strchrnul(m, v146, int32(61))
	mBase = m.M
	if v153 == v146 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v155 = v153 - v146
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+uint32(_c_F_assign_locale_messages[3]))))
	if v157 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v168 = F_strlen(m, v138)
	mBase = m.M
	v172 = F_emscripten_builtin_malloc(m, v155+v168+int32(2))
	mBase = m.M
	if v172 != 0 {
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
	v223 = v218 << (uint(int32(2)) % 32)
	v225 = v223 + int32(8)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[4]))
	if v217 == v227 {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	v198 = v187
	v199 = int32(0)
	v202 = v191
	goto L60
L55:
	;
	v217 = v192
	v218 = int32(0)
	goto L53
L56:
	;
	v192 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v191 != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v192 = v187
	goto L55
L60:
	;
	v203 = F_strncmp(m, v172, v202, v155+int32(1))
	mBase = m.M
	if v203 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[1]))
	v217 = v216
	v218 = v211
	goto L53
L62:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v172
	F___env_rm_add(m, v206, v172)
	mBase = m.M
	goto L52
L63:
	;
	goto L64
L64:
	;
	v211 = v199 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v212 != 0 {
		v198 = v198 + int32(4)
		v199 = v211
		v202 = v212
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	F_emscripten_builtin_free(m, v172)
	mBase = m.M
	goto L52
L67:
	;
	v242 = v239 + v218<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[1])) = v239
	*(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[4])) = v239
	if v172 != 0 {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v229 = F_emscripten_builtin_realloc(m, v227, v225)
	mBase = m.M
	if v229 != 0 {
		v239 = v229
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v230 = F_emscripten_builtin_malloc(m, v225)
	mBase = m.M
	if v230 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L71:
	;
	goto L66
L72:
	;
	if v218 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[1]))
	v235 = F___memcpy(m, v230, v234, v223)
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_assign_locale_messages[4]))
	F_emscripten_builtin_free(m, v237)
	mBase = m.M
	v239 = v230
	goto L67
L76:
	;
	F___env_rm_add(m, int32(0), v172)
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
