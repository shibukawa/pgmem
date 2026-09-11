package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_common_eclass_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == v4 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v161 = int32(0)
	if l2 == v161 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	if v68 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v68 = base.I32_ctz(v54) | v55<<(uint(int32(5))%32)
	goto L2
L4:
	;
	v68 = int32(-2)
	goto L2
L5:
	;
	v21 = base.I32_div_s(int32(0), int32(32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 <= v21 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = l1 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v32 = v29 & int32(-1)
	if v32 != 0 {
		v54 = v32
		v55 = v21
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v34 = v21 + int32(1)
	if v34 == v22 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v34
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37<<(uint(int32(2))%32))))
	if v44 != 0 {
		v54 = v44
		v55 = v37
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v46 = v37 + int32(1)
	if v46 != v22 {
		v37 = v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v156 = v4
	goto L1
L14:
	;
	goto L15
L15:
	;
	v74 = v68
	v75 = v4
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v74 == v78 {
		v92 = v75
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v156 = v92
	goto L1
L18:
	;
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v74<<(uint(int32(2))%32))))
	if v84 == int32(0) {
		v92 = v75
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+136))
	v88 = F_bms_add_members(m, v75, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v92 = v88
	goto L18
L23:
	;
	if int32(0) < v149 {
		v74 = v149
		v75 = v92
		goto L16
	} else {
		goto L34
	}
L24:
	;
	v149 = base.I32_ctz(v135) | v136<<(uint(int32(5))%32)
	goto L23
L25:
	;
	v149 = int32(-2)
	goto L23
L26:
	;
	v100 = v74 + int32(1)
	v102 = base.I32_div_s(v100, int32(32))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v103 <= v102 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v106 = l1 + int32(8)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v102<<(uint(int32(2))%32))))
	v113 = v110 & (int32(-1) << (uint(v100) % 32))
	if v113 != 0 {
		v135 = v113
		v136 = v102
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v115 = v102 + int32(1)
	if v115 == v103 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v118 = v115
	goto L30
L30:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106+v118<<(uint(int32(2))%32))))
	if v125 != 0 {
		v135 = v125
		v136 = v118
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	v127 = v118 + int32(1)
	if v127 != v103 {
		v118 = v127
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L17
L35:
	;
	v369 = F_bms_int_members(m, v156, v363)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L21
	} else {
		goto L87
	}
L36:
	;
	if v214 != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	v214 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v169 = int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v170 <= v169 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v173 = v169
	goto L42
L41:
	;
	v173 = v170
	goto L42
L42:
	;
	v178 = int32(0)
	v181 = int32(-1)
	goto L44
L43:
	;
	v214 = v206
	goto L36
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(8)+v178<<(uint(int32(2))%32))))
	if v188 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v198
	v206 = int32(1)
	goto L43
L46:
	;
	if int32(0) <= v181 {
		v206 = v161
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v198 = v181
	goto L48
L48:
	;
	v200 = v178 + int32(1)
	if v200 != v173 {
		v178 = v200
		v181 = v198
		goto L44
	} else {
		goto L51
	}
L49:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v188)) {
		v206 = v161
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v198 = base.I32_ctz(v188) | v178<<(uint(int32(5))%32)
	goto L48
L51:
	;
	goto L45
L52:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215+v216<<(uint(int32(2))%32))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+136))
	v363 = v221
	goto L35
L53:
	;
	goto L54
L54:
	;
	if l2 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if v278 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	v278 = base.I32_ctz(v264) | v265<<(uint(int32(5))%32)
	goto L55
L57:
	;
	v278 = int32(-2)
	goto L55
L58:
	;
	v231 = base.I32_div_s(int32(0), int32(32))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v232 <= v231 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v235 = l2 + int32(8)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v231<<(uint(int32(2))%32))))
	v242 = v239 & int32(-1)
	if v242 != 0 {
		v264 = v242
		v265 = v231
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v244 = v231 + int32(1)
	if v244 == v232 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v247 = v244
	goto L62
L62:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v235+v247<<(uint(int32(2))%32))))
	if v254 != 0 {
		v264 = v254
		v265 = v247
		goto L56
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v256 = v247 + int32(1)
	if v256 != v232 {
		v247 = v256
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v363 = int32(0)
	goto L35
L67:
	;
	goto L68
L68:
	;
	v284 = int32(0)
	v286 = v278
	goto L69
L69:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v286 == v290 {
		v302 = v284
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v363 = v302
	goto L35
L71:
	;
	if l2 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v286<<(uint(int32(2))%32))))
	if v296 == int32(0) {
		v302 = v284
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v296)+136))
	v300 = F_bms_add_members(m, v284, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	v302 = v300
	goto L71
L75:
	;
	if int32(0) < v359 {
		v284 = v302
		v286 = v359
		goto L69
	} else {
		goto L86
	}
L76:
	;
	v359 = base.I32_ctz(v345) | v346<<(uint(int32(5))%32)
	goto L75
L77:
	;
	v359 = int32(-2)
	goto L75
L78:
	;
	v310 = v286 + int32(1)
	v312 = base.I32_div_s(v310, int32(32))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v313 <= v312 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v316 = l2 + int32(8)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v316+v312<<(uint(int32(2))%32))))
	v323 = v320 & (int32(-1) << (uint(v310) % 32))
	if v323 != 0 {
		v345 = v323
		v346 = v312
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v325 = v312 + int32(1)
	if v325 == v313 {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v328 = v325
	goto L82
L82:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v316+v328<<(uint(int32(2))%32))))
	if v335 != 0 {
		v345 = v335
		v346 = v328
		goto L76
	} else {
		goto L84
	}
L83:
	;
	goto L77
L84:
	;
	v337 = v328 + int32(1)
	if v337 != v313 {
		v328 = v337
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L70
L87:
	;
	m.G0 = v10 + int32(16)
	return v369
}
