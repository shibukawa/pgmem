package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSubqueryScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(759), int32(760))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_convert_subquery_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
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
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	v5 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v26 = v25
	goto L3
L2:
	;
	v26 = v5
	goto L3
L3:
	;
	if l2 == int32(0) {
		v500 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v500
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L8
L8:
	;
	v41 = v5
	v46 = v5
	v49 = v5
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v49<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+41)))
	if v63 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v500 = v473
	goto L4
L11:
	;
	v490 = v49 + int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v490 < v491 {
		v41 = v473
		v46 = v478
		v49 = v490
		goto L9
	} else {
		goto L99
	}
L12:
	;
	v464 = F_lappend(m, v446, v368)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L19
	} else {
		goto L98
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L19
	} else {
		goto L95
	}
L14:
	;
	if v368 == int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L85
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	if v66 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v141 == int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L37
	}
L18:
	;
	v69 = F_get_sortgroupref_tle(m, v66, l3)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+26)))
	if v73 != 0 {
		v500 = v41
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v75 == int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v87 = int32(0)
	goto L24
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v81+v87<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v110 == int32(6) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v119 = F_copyObjectImpl(m, v109)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+8)))
	if v113 == v114 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v117 = v87 + int32(1)
	if v78 != v117 {
		v87 = v117
		goto L24
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v500 = v41
	goto L4
L32:
	;
	if v119 == int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v129 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v132 = F_get_eclass_for_sort_expr(m, l0, v119, v123, v127, v128, v129, v130, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	if v132 == int32(0) {
		v500 = v41
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)))
	v139 = F_make_canonical_pathkey(m, l0, v132, v136, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v368 = v139
	goto L14
L37:
	;
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v145 <= v144 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v368 = int32(0)
	goto L14
L39:
	;
	goto L40
L40:
	;
	v160 = int32(0)
	v168 = v144
	v169 = int32(-1)
	goto L41
L41:
	;
	if l3 == int32(0) {
		v341 = v160
		v350 = v169
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v368 = v341
	goto L14
L43:
	;
	v356 = v168 + int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v356 < v357 {
		v160 = v341
		v168 = v356
		v169 = v350
		goto L41
	} else {
		goto L84
	}
L44:
	;
	v176 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v177 <= v176 {
		v341 = v160
		v350 = v169
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v168<<(uint(int32(2))%32))))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v197 = v160
	v206 = v169
	v207 = v176
	goto L46
L46:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+26)))
	if v216 != 0 {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v341 = v314
	v350 = v323
	goto L43
L48:
	;
	v329 = v207 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v329 < v330 {
		v197 = v314
		v206 = v323
		v207 = v329
		goto L46
	} else {
		goto L83
	}
L49:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v218 == int32(0) {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v221 <= int32(0) {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v230 = int32(0)
	goto L52
L52:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v224+v230<<(uint(int32(2))%32))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v253 == int32(6) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v262 = F_copyObjectImpl(m, v252)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+8)))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+8)))
	if v256 == v257 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v260 = v230 + int32(1)
	if v221 != v260 {
		v230 = v260
		goto L52
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v314 = v197
	v323 = v206
	goto L48
L60:
	;
	if v262 == int32(0) {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L61
	}
L61:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v267 = F_canonicalize_ec_expression(m, v266, v185, v187)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	v269 = F_equal(m, v267, v186)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	if v269 == int32(0) {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L64
	}
L64:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v274 = int32(0)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v277 = F_get_eclass_for_sort_expr(m, l0, v262, v273, v185, v187, v274, v275, v274)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L65
	}
L65:
	;
	if v277 == int32(0) {
		v314 = v197
		v323 = v206
		goto L48
	} else {
		goto L66
	}
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)))
	v284 = F_make_canonical_pathkey(m, l0, v277, v281, v282, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v286 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v289 = v287
	goto L70
L69:
	;
	v289 = int32(0)
	goto L70
L70:
	;
	v291 = v289 - int32(1)
	if v46 < v26 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v46<<(uint(int32(2))%32))))
	if v298 == v284 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v301 = v291
	goto L73
L73:
	;
	v302 = base.B2i32(v206 < v301)
	if v206 < v301 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v300 = v289
	goto L76
L75:
	;
	v300 = v291
	goto L76
L76:
	;
	v301 = v300
	goto L73
L77:
	;
	v303 = v284
	goto L79
L78:
	;
	v303 = v197
	goto L79
L79:
	;
	if v206 < v301 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v304 = v301
	goto L82
L81:
	;
	v304 = v206
	goto L82
L82:
	;
	v314 = v303
	v323 = v304
	goto L48
L83:
	;
	goto L47
L84:
	;
	goto L42
L85:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+40)))
	if v385 != 0 {
		v473 = v41
		v478 = v46
		goto L11
	} else {
		goto L86
	}
L86:
	;
	if v41 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v446 = int32(0)
	goto L12
L88:
	;
	goto L89
L89:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v389 <= int32(0) {
		v446 = v41
		goto L12
	} else {
		goto L90
	}
L90:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v398 = int32(0)
	goto L91
L91:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v392+v398<<(uint(int32(2))%32))))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v384 == v421 {
		v473 = v41
		v478 = v46
		goto L11
	} else {
		goto L93
	}
L92:
	;
	v446 = v41
	goto L12
L93:
	;
	v424 = v398 + int32(1)
	if v424 != v389 {
		v398 = v424
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	F_errmsg_internal(m, int32(335796), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L19
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(488601), int32(1080), int32(111613))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v473 = v464
	v478 = v46 + int32(1)
	goto L11
L99:
	;
	goto L10
}
func F_subquery_is_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
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
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v11 != 0 {
		v350 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L20
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L20
	} else {
		goto L109
	}
L3:
	;
	return v350
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v12 != 0 {
		v350 = v4
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v14 != 0 {
		v350 = v4
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v15 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v22 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v20)
	goto L10
L12:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v16 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v17 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if l0 == l1 {
		goto L81
	} else {
		goto L82
	}
L16:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+45)))
	if v24 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v28 = F_flatten_group_exprs(m, int32(0), l0, v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v32 = v23
	goto L19
L19:
	;
	if v32 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	return int32(0)
L21:
	;
	v32 = v28
	goto L19
L22:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v45 = v4
	goto L24
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v45<<(uint(int32(2))%32))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+26)))
	if v53 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L15
L26:
	;
	v243 = v45 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v243 < v244 {
		v45 = v243
		goto L24
	} else {
		goto L79
	}
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v55 != int32(1) {
		v72 = v54
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225))))
	v228 = v221 + v227
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v230 = v229 | v220
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v230)
	goto L26
L29:
	;
	v75 = v52 + int32(8)
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+8)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v76))))
	if v78&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+8)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v58))))
	if v60&int32(2) != 0 {
		v72 = v54
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v64 = F_expression_returns_set(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v64 == int32(0) {
		v72 = v66
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v220 = int32(2)
	v221 = v66
	v225 = v52 + int32(8)
	goto L28
L34:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v220 = v215
	v221 = v216
	v225 = v75
	goto L28
L35:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v87 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v82 = F_contain_volatile_functions(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	if v82 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v215 = int32(1)
	goto L34
L39:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v133 != int32(1) {
		goto L26
	} else {
		goto L56
	}
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v91))))
	if v93&int32(4) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v97 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v99 == v97 {
		v130 = v97
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v130 != 0 {
		goto L39
	} else {
		goto L55
	}
L43:
	;
	goto L42
L44:
	;
	if v96 == int32(0) {
		v130 = v97
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if int32(0) < v104 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v108 = int32(0)
	goto L49
L47:
	;
	goto L48
L48:
	;
	v130 = v97
	goto L43
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v108<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v99 == v117 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	v130 = int32(1)
	goto L43
L52:
	;
	goto L53
L53:
	;
	v121 = v108 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v121 < v122 {
		v108 = v121
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v215 = int32(4)
	goto L34
L56:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75))))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
	if v139&int32(4) != 0 {
		goto L26
	} else {
		goto L57
	}
L57:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v142 == int32(0) {
		goto L26
	} else {
		goto L58
	}
L58:
	;
	v145 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v146 <= v145 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v153 = v145
	goto L60
L60:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v153<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v167 == v165 {
		v198 = v165
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v215 = int32(8)
	goto L34
L62:
	;
	if v198 != 0 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	goto L62
L64:
	;
	if v164 == int32(0) {
		v198 = v165
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if int32(0) < v172 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v176 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v198 = v165
	goto L63
L69:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v176<<(uint(int32(2))%32))))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v167 == v185 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L68
L71:
	;
	v198 = int32(1)
	goto L63
L72:
	;
	goto L73
L73:
	;
	v189 = v176 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v189 < v190 {
		v176 = v189
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v201 = v153 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v201 < v202 {
		v153 = v201
		goto L60
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L61
L78:
	;
	goto L26
L79:
	;
	goto L25
L80:
	;
	v350 = int32(1)
	goto L3
L81:
	;
	if v256 == int32(0) {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v256 != 0 {
		v350 = v4
		goto L3
	} else {
		goto L87
	}
L84:
	;
	v260 = F_recurse_pushdown_safe(m, v256, l1, l2)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	if v260 != 0 {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v350 = v4
	goto L3
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	if v264 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v267 = v265
	goto L90
L89:
	;
	v267 = int32(0)
	goto L90
L90:
	;
	if v262 == int32(0) {
		v329 = v267
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v329 != 0 {
		goto L1
	} else {
		goto L108
	}
L92:
	;
	v270 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v271 <= v270 {
		v329 = v267
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v274 = v270
	v275 = v271
	v279 = v267
	goto L94
L94:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284+v274<<(uint(int32(2))%32))))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+26)))
	if v289 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v329 = v320
	goto L91
L96:
	;
	if v279 == int32(0) {
		goto L2
	} else {
		goto L99
	}
L97:
	;
	v318 = v275
	v320 = v279
	goto L98
L98:
	;
	v322 = v274 + int32(1)
	if v322 < v318 {
		v274 = v322
		v275 = v318
		v279 = v320
		goto L94
	} else {
		goto L107
	}
L99:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v295 = F_exprType(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L20
	} else {
		goto L100
	}
L100:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v295 != v297 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288)+8)))
	v301 = v299 + v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v304 = v302 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v304)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v308 = v279 + int32(4)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if base.Ui32(v308) < base.Ui32(v310+v311<<(uint(int32(2))%32)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v316 = v308
	goto L106
L105:
	;
	v316 = int32(0)
	goto L106
L106:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v318 = v317
	v320 = v316
	goto L98
L107:
	;
	goto L95
L108:
	;
	goto L80
L109:
	;
	F_errmsg_internal(m, int32(166211), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(489093), int32(3862), int32(160502))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errmsg_internal(m, int32(166211), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L20
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(489093), int32(3868), int32(160502))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
