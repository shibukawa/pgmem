package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageManagerInitialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v3
	v5 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v3
	v11 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
	if l0 != 0 {
		v17 = l0 - l1 + v5
	} else {
		v17 = v11
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17
	v24 = F__emscripten_memset_bulkmem(m, l0+int32(36), base.I32_extend8_s(int32(0)), int32(516))
	mBase = m.M
	return
}
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v17) < base.Ui32(int32(24)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L20
	} else {
		goto L107
	}
L2:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v20) < base.Ui32(v17) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.Ui32(v22) < base.Ui32(v20) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.Ui32(int32(8193)) <= base.Ui32(v22) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v17 != int32(24) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = int32(base.Ui32(v17+int32(262120)) >> (uint(int32(2)) % 32))
	goto L8
L7:
	;
	v33 = int32(0)
	goto L8
L8:
	;
	v34 = int32(1)
	v35 = v33 + v34
	if base.Ui32((l3-v34)&int32(65535)) <= base.Ui32(int32(2047)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	m.G0 = v15 + int32(16)
	return v397 & int32(65535)
L10:
	;
	v188 = v178 & int32(65535)
	if l4&int32(2) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L11:
	;
	v154 = int32(65535)
	if base.Ui32(v145&v154) <= base.Ui32(v35&v154) {
		v178 = v145
		v182 = v149
		goto L10
	} else {
		goto L35
	}
L12:
	;
	if l4&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	if v74&int32(1) == int32(0) {
		v178 = v35
		v182 = v6
		goto L10
	} else {
		goto L25
	}
L15:
	;
	if base.Ui32(v33&int32(65535)) < base.Ui32(l3) {
		v145 = l3
		v149 = v6
		goto L11
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v145 = l3
	v149 = base.B2i32(base.Ui32(l3) <= base.Ui32(v33&int32(65535)))
	goto L11
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32)+l0)+20))
	if base.Ui32(v50) < base.Ui32(int32(32768)) {
		v145 = l3
		v149 = v6
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v53 = int32(0)
	v56 = F_errstart(m, int32(19), v53)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v56 == int32(0) {
		v397 = v53
		goto L9
	} else {
		goto L22
	}
L22:
	;
	F_errmsg_internal(m, int32(457667), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(490783), int32(235), int32(454017))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v397 = v53
	goto L9
L25:
	;
	v80 = v33 & int32(65535)
	if v80 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v145 = v132
	v149 = int32(0)
	goto L11
L27:
	;
	v127 = v74 & int32(65534)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v127)
	v132 = v117
	goto L26
L28:
	;
	v117 = int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v90 = int32(1)
	goto L31
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v90&int32(65535)<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
	if base.Ui32(v106) < base.Ui32(int32(32768)) {
		v132 = v90
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v117 = v110
	goto L27
L33:
	;
	v110 = v90 + int32(1)
	if base.Ui32(v110&int32(65535)) <= base.Ui32(v80) {
		v90 = v110
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v159 = int32(0)
	v162 = F_errstart(m, int32(19), v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	if v162 == int32(0) {
		v397 = v159
		goto L9
	} else {
		goto L37
	}
L37:
	;
	F_errmsg_internal(m, int32(393715), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(490783), int32(292), int32(454017))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	v397 = v159
	goto L9
L40:
	;
	v213 = v17 + int32(4)
	v215 = v35 & int32(65535)
	if v188 == v215 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	if base.Ui32(v188) < base.Ui32(int32(292)) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v195 = int32(0)
	v198 = F_errstart(m, int32(19), v195)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	if v198 == int32(0) {
		v397 = v195
		goto L9
	} else {
		goto L44
	}
L44:
	;
	F_errmsg_internal(m, int32(401133), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(490783), int32(299), int32(454017))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v397 = v195
	goto L9
L47:
	;
	v217 = v213
	goto L49
L48:
	;
	v217 = v17
	goto L49
L49:
	;
	if v182 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v218 = v213
	goto L52
L51:
	;
	v218 = v217
	goto L52
L52:
	;
	v223 = v20 - (l2+int32(7))&int32(-8)
	if v223 < v218 {
		v397 = int32(0)
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v227 = v188<<(uint(int32(2))%32) + l0
	v229 = v227 + int32(20)
	if v182 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v231 = v227 + int32(24)
	v234 = (v215 - v188) << (uint(int32(2)) % 32)
	if v231 == v229 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v223&int32(32767) | l2<<(uint(int32(17))%32) | int32(32768)
	if l2 != 0 {
		goto L104
	} else {
		goto L105
	}
L57:
	;
	goto L56
L58:
	;
	goto L57
L59:
	;
	v238 = v231 + v234
	if base.Ui32(v229-v238) <= base.Ui32(int32(0)-v234<<(uint(int32(1))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v245 = F___memcpy(m, v231, v229, v234)
	mBase = m.M
	goto L57
L61:
	;
	goto L62
L62:
	;
	v248 = (v231 ^ v229) & int32(3)
	if base.Ui32(v231) < base.Ui32(v229) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v350 == int32(0) {
		goto L58
	} else {
		goto L99
	}
L64:
	;
	if base.Ui32(v328) <= base.Ui32(int32(3)) {
		v349 = v327
		v350 = v328
		v351 = v329
		goto L63
	} else {
		goto L95
	}
L65:
	;
	if v248 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v248 != 0 {
		v310 = v234
		goto L78
	} else {
		goto L79
	}
L68:
	;
	v349 = v229
	v350 = v234
	v351 = v231
	goto L63
L69:
	;
	goto L70
L70:
	;
	if v231&int32(3) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v327 = v229
	v328 = v234
	v329 = v231
	goto L64
L72:
	;
	goto L73
L73:
	;
	v255 = v229
	v256 = v234
	v257 = v231
	goto L74
L74:
	;
	if v256 == int32(0) {
		goto L58
	} else {
		goto L76
	}
L75:
	;
	v327 = v264
	v328 = v266
	v329 = v268
	goto L64
L76:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	*(*uint8)(unsafe.Add(mBase, uint32(v257))) = uint8(v261)
	v263 = int32(1)
	v264 = v255 + v263
	v266 = v256 - v263
	v268 = v257 + v263
	if v268&int32(3) != 0 {
		v255 = v264
		v256 = v266
		v257 = v268
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	if v310 == int32(0) {
		goto L58
	} else {
		goto L91
	}
L79:
	;
	if v238&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v275 = v234
	goto L83
L81:
	;
	v290 = v234
	goto L82
L82:
	;
	if base.Ui32(v290) <= base.Ui32(int32(3)) {
		v310 = v290
		goto L78
	} else {
		goto L87
	}
L83:
	;
	if v275 == int32(0) {
		goto L58
	} else {
		goto L85
	}
L84:
	;
	v290 = v281
	goto L82
L85:
	;
	v281 = v275 - int32(1)
	v282 = v231 + v281
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v281))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v284)
	if v282&int32(3) != 0 {
		v275 = v281
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v297 = v290
	goto L88
L88:
	;
	v301 = v297 - int32(4)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v229+v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v301))) = v304
	if base.Ui32(int32(3)) < base.Ui32(v301) {
		v297 = v301
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v310 = v301
	goto L78
L90:
	;
	goto L89
L91:
	;
	v317 = v310
	goto L92
L92:
	;
	v321 = v317 - int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v231+v321))) = uint8(v324)
	if v321 != 0 {
		v317 = v321
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L58
L94:
	;
	goto L93
L95:
	;
	v334 = v327
	v335 = v328
	v336 = v329
	goto L96
L96:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = v338
	v340 = int32(4)
	v341 = v334 + v340
	v343 = v336 + v340
	v345 = v335 - v340
	if base.Ui32(int32(3)) < base.Ui32(v345) {
		v334 = v341
		v335 = v345
		v336 = v343
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v349 = v341
	v350 = v345
	v351 = v343
	goto L63
L98:
	;
	goto L97
L99:
	;
	v356 = v349
	v357 = v350
	v358 = v351
	goto L100
L100:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	*(*uint8)(unsafe.Add(mBase, uint32(v358))) = uint8(v360)
	v362 = int32(1)
	v367 = v357 - v362
	if v367 != 0 {
		v356 = v356 + v362
		v357 = v367
		v358 = v358 + v362
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L58
L102:
	;
	goto L101
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v223)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v218)
	v397 = v178
	goto L9
L104:
	;
	v388 = F__emscripten_memcpy_bulkmem(m, l0+v223, l1, l2)
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v419
	F_errmsg(m, int32(56420), v15)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L20
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(490783), int32(217), int32(454017))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PageRestoreTempPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v5 = v3 << (uint(int32(8)) % 32)
	if v5 != 0 {
		v6 = F__emscripten_memcpy_bulkmem(m, l1, l0, v5)
		mBase = m.M
	} else {
	}
	F_pfree(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_PageSetChecksumInplace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v3 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[199]))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+252))
		if base.B2i32(v8 != int32(0)) == int32(0) {
		} else {
			v13 = F_pg_checksum_page(m, l0, l1)
			mBase = m.M
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v13)
		}
	}
	return
}
func F_ReadPageInternal(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v105 int32
	_ = v105
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v14 = base.I32_wrap_i64(l1) & (v11 - int32(1))
	v16 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v11))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	if v16 != v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v105
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	if v16 == v17 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1192))
	if v14 != v19 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if base.Ui32(l2) <= base.Ui32(v21) {
		v105 = v21
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v94
	v105 = v95
	goto L1
L7:
	;
	v91 = int32(0)
	v94 = v91
	v95 = int32(-1)
	v96 = v91
	v99 = int64(0)
	goto L6
L8:
	;
	v52 = int32(-2)
	v53 = int32(24)
	if base.Ui32(l2) <= base.Ui32(v53) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	if v14 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = l1 - base.I64_extend_i32_u(v14)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = m.T0[v35].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, v31, int32(8192), v33, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v36 == int32(-2) {
		v105 = int32(-2)
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v36 < int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v45 = F_XLogReaderValidatePageHeader(m, l0, v31, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	if v45 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	v56 = v53
	goto L19
L18:
	;
	v56 = l2
	goto L19
L19:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = m.T0[v59].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v56, v57, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v60 == int32(-2) {
		v105 = v52
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v60 < int32(25) {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
	if v69&int32(2) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = int32(40)
	goto L25
L24:
	;
	v72 = int32(24)
	goto L25
L25:
	;
	if base.Ui32(v60) < base.Ui32(v72) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = m.T0[v75].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v72, v74, v68)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	v82 = v60
	goto L28
L28:
	;
	v83 = F_XLogReaderValidatePageHeader(m, l0, l1, v68)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L32
	}
L29:
	;
	if v76 == int32(-2) {
		v105 = v52
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v76 < int32(0) {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v82 = v76
	goto L28
L32:
	;
	if v83 != 0 {
		v94 = v82
		v95 = v82
		v96 = v14
		v99 = v16
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L7
}
