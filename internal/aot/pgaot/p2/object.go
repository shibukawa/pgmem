package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RunObjectPostAlterHookStr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	v19 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	m.T0[v19].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(2), int32(6243), l0, l1, v7+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_get_object_attnum_namespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+24)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[290])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[291])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(724400)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[291])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(724400)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[291])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(724400)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[291])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(724400)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(56882), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(470610), int32(2777), int32(480420))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+26)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[290])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[291])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(724400)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[291])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(724400)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[291])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(724400)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[291])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(724400)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(56882), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(470610), int32(2777), int32(480420))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 != 0 {
		return int32(0)
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v5 != 0 {
			return int32(0)
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
			v9 = F_cstring_to_text_with_len(m, v6, v7-v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				return int32(0)
			}
		}
	}
}
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	v4 = int32(0)
	if l0&int32(3) == v4 {
		v33 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v66 = v58 - l0
	goto L1
L3:
	;
	v37 = v33
	goto L12
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v66 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v22 = l0
	goto L8
L8:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v58 = v26
	goto L2
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v52 = v37
	goto L15
L14:
	;
	goto L13
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v58 = v52
	goto L2
L17:
	;
	goto L16
L18:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v128 = v4
	v129 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if l1&int32(3) == int32(0) {
		v94 = l1
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v128 = int32(1)
	v129 = v127
	goto L18
L23:
	;
	v127 = v119 - l1
	goto L22
L24:
	;
	v98 = v94
	goto L33
L25:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v78 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v127 = int32(0)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v83 = l1
	goto L29
L29:
	;
	v87 = v83 + int32(1)
	if v87&int32(3) == int32(0) {
		v94 = v87
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v119 = v87
	goto L23
L31:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v92 != 0 {
		v83 = v87
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v107 = int32(-2139062144)
	if (int32(16843008)-v104|v104)&v107 == v107 {
		v98 = v98 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v113 = v98
	goto L36
L35:
	;
	goto L34
L36:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 != 0 {
		v113 = v113 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v119 = v113
	goto L23
L38:
	;
	goto L37
L39:
	;
	if l2&int32(3) == int32(0) {
		v153 = l2
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v190 = v128
	goto L41
L41:
	;
	v192 = int32(63) - v190
	if v129+v66 <= v192 {
		goto L60
	} else {
		goto L61
	}
L42:
	;
	v190 = v186 + v128 + int32(1)
	goto L41
L43:
	;
	v186 = v178 - l2
	goto L42
L44:
	;
	v157 = v153
	goto L53
L45:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v137 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v186 = int32(0)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v142 = l2
	goto L49
L49:
	;
	v146 = v142 + int32(1)
	if v146&int32(3) == int32(0) {
		v153 = v146
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v178 = v146
	goto L43
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v151 != 0 {
		v142 = v146
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v166 = int32(-2139062144)
	if (int32(16843008)-v163|v163)&v166 == v166 {
		v157 = v157 + int32(4)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v172 = v157
	goto L56
L55:
	;
	goto L54
L56:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v176 != 0 {
		v172 = v172 + int32(1)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v178 = v172
	goto L43
L58:
	;
	goto L57
L59:
	;
	v219 = F_pg_mbcliplen(m, l0, v215, v215)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v213 = v129
	v215 = v66
	goto L59
L61:
	;
	goto L62
L62:
	;
	v198 = v129
	v199 = v66
	goto L63
L63:
	;
	v206 = v199 - base.B2i32(v198 < v199)
	v207 = v198 - base.B2i32(v199 <= v198)
	if v192 < v206+v207 {
		v198 = v207
		v199 = v206
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v213 = v207
	v215 = v206
	goto L59
L65:
	;
	goto L64
L66:
	;
	return int32(0)
L67:
	;
	if l1 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v223 = F_pg_mbcliplen(m, l1, v213, v213)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	v225 = v213
	goto L70
L70:
	;
	v230 = F_palloc(m, v219+v190+v225+int32(1))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L66
	} else {
		goto L72
	}
L71:
	;
	v225 = v223
	goto L70
L72:
	;
	if v219 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if l1 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v232 = F__emscripten_memcpy_bulkmem(m, v230, l0, v219)
	mBase = m.M
	v233 = v232
	goto L76
L75:
	;
	v233 = v230
	goto L76
L76:
	;
	goto L73
L77:
	;
	v235 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v219+v233))) = uint8(v235)
	v238 = v219 + int32(1)
	if v225 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v244 = v219
	goto L79
L79:
	;
	v245 = v244 + v233
	if l2 != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v244 = v225 + v238
	goto L79
L81:
	;
	v240 = F__emscripten_memcpy_bulkmem(m, v233+v238, l1, v225)
	mBase = m.M
	goto L83
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v246 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v246)
	v249 = v245 + int32(1)
	if (l2^v249)&int32(3) != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	goto L86
L86:
	;
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v325)
	return v233
L87:
	;
	return v233
L88:
	;
	goto L87
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v303)
	if v303&int32(255) == int32(0) {
		goto L88
	} else {
		goto L104
	}
L90:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v302 = l2
	v303 = v255
	v304 = v249
	goto L89
L91:
	;
	goto L92
L92:
	;
	if l2&int32(3) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v259 = l2
	v261 = v249
	goto L96
L94:
	;
	v273 = l2
	v275 = v249
	goto L95
L95:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v280 = int32(-2139062144)
	if (int32(16843008)-v277|v277)&v280 != v280 {
		v302 = v273
		v303 = v277
		v304 = v275
		goto L89
	} else {
		goto L100
	}
L96:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v262)
	if v262 == int32(0) {
		goto L88
	} else {
		goto L98
	}
L97:
	;
	v273 = v269
	v275 = v267
	goto L95
L98:
	;
	v266 = int32(1)
	v267 = v261 + v266
	v269 = v259 + v266
	if v269&int32(3) != 0 {
		v259 = v269
		v261 = v267
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v285 = v273
	v286 = v277
	v287 = v275
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v286
	v289 = int32(4)
	v290 = v287 + v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v293 = v285 + v289
	v297 = int32(-2139062144)
	if (v291|(int32(16843008)-v291))&v297 == v297 {
		v285 = v293
		v286 = v291
		v287 = v290
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v302 = v293
	v303 = v291
	v304 = v290
	goto L89
L103:
	;
	goto L102
L104:
	;
	v311 = v302
	v313 = v304
	goto L105
L105:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)) = uint8(v314)
	v316 = int32(1)
	if v314 != 0 {
		v311 = v311 + v316
		v313 = v313 + v316
		goto L105
	} else {
		goto L107
	}
L106:
	;
	goto L88
L107:
	;
	goto L106
}
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_object_aclmask_ext(m, l0, l1, l2, l3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_object_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_object_aclmask_ext(m, l0, l1, l2, l3, l4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_object_address_present(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = v7 - int32(1)
	if v9 < v3 {
		v44 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v44
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = v9
	goto L3
L3:
	;
	v22 = v13 + v16*int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v12 != v23 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v44 = v38
	goto L1
L5:
	;
	v38 = int32(0)
	if v38 < v16 {
		v16 = v16 - int32(1)
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 != v26 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v29 == v30 {
		v44 = v28
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v29 == int32(0) {
		v44 = v28
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	goto L4
}
func F_record_object_address_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(2) <= v11 {
		F_pg_qsort(m, v10, v11, int32(12), int32(462))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v19 = int32(1)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if int32(2) <= v20 {
				v27 = v18
				v28 = v19
				v30 = int32(1)
				for {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v37 = v34 + v30*int32(12)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v33 != v38 {
						v50 = v27 + int32(12)
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
						v57 = v50
						v58 = v28 + int32(1)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						if v40 != v41 {
							v50 = v27 + int32(12)
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
							v57 = v50
							v58 = v28 + int32(1)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							if v43 == v44 {
								v57 = v27
								v58 = v28
							} else {
								if v43 != 0 {
									v50 = v27 + int32(12)
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
									*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
									v57 = v50
									v58 = v28 + int32(1)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v44
									v57 = v27
									v58 = v28
								}
							}
						}
					}
					v62 = v30 + int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v62 < v63 {
						v27 = v57
						v28 = v58
						v30 = v62
						continue
					} else {
						break
					}
					break
				}
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v69 = v65
				v70 = v58
			} else {
				v69 = v18
				v70 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v70
			v79 = v69
			v80 = v70
			F_recordMultipleDependencies(m, l0, v79, v80, l2)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v79 = v10
		v80 = v11
		F_recordMultipleDependencies(m, l0, v79, v80, l2)
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return
		} else {
			return
		}
	}
}
