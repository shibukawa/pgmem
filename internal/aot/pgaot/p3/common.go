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
	var v77 int32
	_ = v77
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
	var v93 int32
	_ = v93
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
	var v158 int32
	_ = v158
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
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if int32(0) < v68 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v68 = base.I32_ctz(v54) | v55<<(uint(int32(5))%32)
	goto L1
L3:
	;
	v68 = int32(-2)
	goto L1
L4:
	;
	v21 = base.I32_div_s(int32(0), int32(32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 <= v21 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = l1 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v32 = v29 & int32(-1)
	if v32 != 0 {
		v54 = v32
		v55 = v21
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = v21 + int32(1)
	if v34 == v22 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v37 = v34
	goto L8
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37<<(uint(int32(2))%32))))
	if v44 != 0 {
		v54 = v44
		v55 = v37
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v46 = v37 + int32(1)
	if v46 != v22 {
		v37 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v74 = v68
	v77 = v4
	goto L15
L13:
	;
	v158 = v4
	goto L14
L14:
	;
	v161 = int32(0)
	if l2 == v161 {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v74 == v78 {
		v93 = v77
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v158 = v93
	goto L14
L17:
	;
	if l1 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v74<<(uint(int32(2))%32))))
	if v84 == int32(0) {
		v93 = v77
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+136))
	v88 = F_bms_add_members(m, v77, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v93 = v88
	goto L17
L22:
	;
	if int32(0) < v149 {
		v74 = v149
		v77 = v93
		goto L15
	} else {
		goto L33
	}
L23:
	;
	v149 = base.I32_ctz(v135) | v136<<(uint(int32(5))%32)
	goto L22
L24:
	;
	v149 = int32(-2)
	goto L22
L25:
	;
	v100 = v74 + int32(1)
	v102 = base.I32_div_s(v100, int32(32))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v103 <= v102 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v106 = l1 + int32(8)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v102<<(uint(int32(2))%32))))
	v113 = v110 & (int32(-1) << (uint(v100) % 32))
	if v113 != 0 {
		v135 = v113
		v136 = v102
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v115 = v102 + int32(1)
	if v115 == v103 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v118 = v115
	goto L29
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106+v118<<(uint(int32(2))%32))))
	if v125 != 0 {
		v135 = v125
		v136 = v118
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v127 = v118 + int32(1)
	if v127 != v103 {
		v118 = v127
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L16
L34:
	;
	v370 = F_bms_int_members(m, v158, v364)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L20
	} else {
		goto L85
	}
L35:
	;
	if v215 != 0 {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	v215 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v169 = int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v170 <= v169 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v173 = v169
	goto L41
L40:
	;
	v173 = v170
	goto L41
L41:
	;
	v178 = int32(0)
	v180 = int32(-1)
	goto L43
L42:
	;
	v215 = v207
	goto L35
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(8)+v178<<(uint(int32(2))%32))))
	if v188 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v199
	v207 = int32(1)
	goto L42
L45:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v188)))|base.B2i32(int32(0) <= v180) != 0 {
		v207 = v161
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v199 = v180
	goto L47
L47:
	;
	v201 = v178 + int32(1)
	if v201 != v173 {
		v178 = v201
		v180 = v199
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v199 = base.I32_ctz(v188) | v178<<(uint(int32(5))%32)
	goto L47
L49:
	;
	goto L44
L50:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216+v217<<(uint(int32(2))%32))))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+136))
	v364 = v222
	goto L34
L51:
	;
	goto L52
L52:
	;
	if l2 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v279 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	v279 = base.I32_ctz(v265) | v266<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v279 = int32(-2)
	goto L53
L56:
	;
	v232 = base.I32_div_s(int32(0), int32(32))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v233 <= v232 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v236 = l2 + int32(8)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236+v232<<(uint(int32(2))%32))))
	v243 = v240 & int32(-1)
	if v243 != 0 {
		v265 = v243
		v266 = v232
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v245 = v232 + int32(1)
	if v245 == v233 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v248 = v245
	goto L60
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v236+v248<<(uint(int32(2))%32))))
	if v255 != 0 {
		v265 = v255
		v266 = v248
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v257 = v248 + int32(1)
	if v257 != v233 {
		v248 = v257
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v364 = int32(0)
	goto L34
L65:
	;
	goto L66
L66:
	;
	v285 = int32(0)
	v287 = v279
	goto L67
L67:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v287 == v291 {
		v303 = v285
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v364 = v303
	goto L34
L69:
	;
	if l2 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v287<<(uint(int32(2))%32))))
	if v297 == int32(0) {
		v303 = v285
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v297)+136))
	v301 = F_bms_add_members(m, v285, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	v303 = v301
	goto L69
L73:
	;
	if int32(0) < v360 {
		v285 = v303
		v287 = v360
		goto L67
	} else {
		goto L84
	}
L74:
	;
	v360 = base.I32_ctz(v346) | v347<<(uint(int32(5))%32)
	goto L73
L75:
	;
	v360 = int32(-2)
	goto L73
L76:
	;
	v311 = v287 + int32(1)
	v313 = base.I32_div_s(v311, int32(32))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v314 <= v313 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v317 = l2 + int32(8)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v313<<(uint(int32(2))%32))))
	v324 = v321 & (int32(-1) << (uint(v311) % 32))
	if v324 != 0 {
		v346 = v324
		v347 = v313
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v326 = v313 + int32(1)
	if v326 == v314 {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v329 = v326
	goto L80
L80:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v317+v329<<(uint(int32(2))%32))))
	if v336 != 0 {
		v346 = v336
		v347 = v329
		goto L74
	} else {
		goto L82
	}
L81:
	;
	goto L75
L82:
	;
	v338 = v329 + int32(1)
	if v338 != v314 {
		v329 = v338
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	goto L68
L85:
	;
	m.G0 = v10 + int32(16)
	return v370
}
