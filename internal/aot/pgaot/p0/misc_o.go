package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OffsetVarNodes_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	v3 = int32(0)
	if l0 == v3 {
		v492 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510 + v511
	return int32(0)
L2:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v497 + int32(1)
	v503 = F_query_tree_walker_impl(m, l0, int32(1048), l1, int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L31
	} else {
		goto L113
	}
L3:
	;
	return v492
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(58) {
	case 0:
		goto L10
	case 1, 2, 3, 4:
		v461 = v10
		goto L6
	case 5:
		goto L9
	case 6:
		goto L8
	default:
		goto L11
	}
L5:
	;
	v487 = F_expression_tree_walker_impl(m, l0, int32(1048), l1)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L31
	} else {
		goto L112
	}
L6:
	;
	if v461 == int32(67) {
		goto L2
	} else {
		goto L109
	}
L7:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v180 != v181 {
		goto L5
	} else {
		goto L50
	}
L8:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v173 == int32(0) {
		goto L5
	} else {
		goto L48
	}
L9:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v172 != 0 {
		v492 = v3
		goto L3
	} else {
		goto L47
	}
L10:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v171 != 0 {
		v492 = v3
		goto L3
	} else {
		goto L46
	}
L11:
	;
	if v10 == int32(319) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v10 != int32(6) {
		v461 = v10
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 != v18 {
		v492 = v3
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v25 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if int32(0) <= v82 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v82 = base.I32_ctz(v68) | v69<<(uint(int32(5))%32)
	goto L15
L17:
	;
	v82 = int32(-2)
	goto L15
L18:
	;
	v35 = base.I32_div_s(int32(0), int32(32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v36 <= v35 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v39 = v25 + int32(8)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v35<<(uint(int32(2))%32))))
	v46 = v43 & int32(-1)
	if v46 != 0 {
		v68 = v46
		v69 = v35
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v48 = v35 + int32(1)
	if v48 == v36 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v51 = v48
	goto L22
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39+v51<<(uint(int32(2))%32))))
	if v58 != 0 {
		v68 = v58
		v69 = v51
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L17
L24:
	;
	v60 = v51 + int32(1)
	if v60 != v36 {
		v51 = v60
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v87 = v82
	v90 = v3
	goto L29
L27:
	;
	v160 = v3
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v163 == int32(0) {
		v492 = v3
		goto L3
	} else {
		goto L45
	}
L29:
	;
	v93 = F_bms_add_member(m, v90, v87+v24)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v160 = v93
	goto L28
L31:
	;
	return int32(0)
L32:
	;
	if v25 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if int32(0) <= v152 {
		v87 = v152
		v90 = v93
		goto L29
	} else {
		goto L44
	}
L34:
	;
	v152 = base.I32_ctz(v138) | v139<<(uint(int32(5))%32)
	goto L33
L35:
	;
	v152 = int32(-2)
	goto L33
L36:
	;
	v103 = v87 + int32(1)
	v105 = base.I32_div_s(v103, int32(32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v106 <= v105 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v109 = v25 + int32(8)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v105<<(uint(int32(2))%32))))
	v116 = v113 & (int32(-1) << (uint(v103) % 32))
	if v116 != 0 {
		v138 = v116
		v139 = v105
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v118 = v105 + int32(1)
	if v118 == v106 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v121 = v118
	goto L40
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v109+v121<<(uint(int32(2))%32))))
	if v128 != 0 {
		v138 = v128
		v139 = v121
		goto L34
	} else {
		goto L42
	}
L41:
	;
	goto L35
L42:
	;
	v130 = v121 + int32(1)
	if v130 != v106 {
		v121 = v130
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L30
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v166 + v163
	return int32(0)
L46:
	;
	goto L1
L47:
	;
	goto L1
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v176 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v177 + v173
	goto L5
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v185 == v184 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if int32(0) <= v242 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v242 = base.I32_ctz(v228) | v229<<(uint(int32(5))%32)
	goto L51
L53:
	;
	v242 = int32(-2)
	goto L51
L54:
	;
	v195 = base.I32_div_s(int32(0), int32(32))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v196 <= v195 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v199 = v185 + int32(8)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v195<<(uint(int32(2))%32))))
	v206 = v203 & int32(-1)
	if v206 != 0 {
		v228 = v206
		v229 = v195
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v208 = v195 + int32(1)
	if v208 == v196 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v211 = v208
	goto L58
L58:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199+v211<<(uint(int32(2))%32))))
	if v218 != 0 {
		v228 = v218
		v229 = v211
		goto L52
	} else {
		goto L60
	}
L59:
	;
	goto L53
L60:
	;
	v220 = v211 + int32(1)
	if v220 != v196 {
		v211 = v220
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v247 = v184
	v248 = v242
	goto L65
L63:
	;
	v315 = v184
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v315
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v322 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	v253 = F_bms_add_member(m, v247, v248+v183)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L31
	} else {
		goto L67
	}
L66:
	;
	v315 = v253
	goto L64
L67:
	;
	if v185 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if int32(0) <= v310 {
		v247 = v253
		v248 = v310
		goto L65
	} else {
		goto L79
	}
L69:
	;
	v310 = base.I32_ctz(v296) | v297<<(uint(int32(5))%32)
	goto L68
L70:
	;
	v310 = int32(-2)
	goto L68
L71:
	;
	v261 = v248 + int32(1)
	v263 = base.I32_div_s(v261, int32(32))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v264 <= v263 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v267 = v185 + int32(8)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v263<<(uint(int32(2))%32))))
	v274 = v271 & (int32(-1) << (uint(v261) % 32))
	if v274 != 0 {
		v296 = v274
		v297 = v263
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v276 = v263 + int32(1)
	if v276 == v264 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v279 = v276
	goto L75
L75:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v267+v279<<(uint(int32(2))%32))))
	if v286 != 0 {
		v296 = v286
		v297 = v279
		goto L69
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v288 = v279 + int32(1)
	if v288 != v264 {
		v279 = v288
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L66
L80:
	;
	if int32(0) <= v379 {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v379 = base.I32_ctz(v365) | v366<<(uint(int32(5))%32)
	goto L80
L82:
	;
	v379 = int32(-2)
	goto L80
L83:
	;
	v332 = base.I32_div_s(int32(0), int32(32))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v333 <= v332 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v336 = v322 + int32(8)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336+v332<<(uint(int32(2))%32))))
	v343 = v340 & int32(-1)
	if v343 != 0 {
		v365 = v343
		v366 = v332
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v345 = v332 + int32(1)
	if v345 == v333 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v348 = v345
	goto L87
L87:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v336+v348<<(uint(int32(2))%32))))
	if v355 != 0 {
		v365 = v355
		v366 = v348
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L82
L89:
	;
	v357 = v348 + int32(1)
	if v357 != v333 {
		v348 = v357
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v385 = v379
	v387 = v3
	goto L94
L92:
	;
	v455 = v3
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v455
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = v458
	goto L6
L94:
	;
	v390 = F_bms_add_member(m, v387, v321+v385)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L31
	} else {
		goto L96
	}
L95:
	;
	v455 = v390
	goto L93
L96:
	;
	if v322 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if int32(0) <= v447 {
		v385 = v447
		v387 = v390
		goto L94
	} else {
		goto L108
	}
L98:
	;
	v447 = base.I32_ctz(v433) | v434<<(uint(int32(5))%32)
	goto L97
L99:
	;
	v447 = int32(-2)
	goto L97
L100:
	;
	v398 = v385 + int32(1)
	v400 = base.I32_div_s(v398, int32(32))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v401 <= v400 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v404 = v322 + int32(8)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v404+v400<<(uint(int32(2))%32))))
	v411 = v408 & (int32(-1) << (uint(v398) % 32))
	if v411 != 0 {
		v433 = v411
		v434 = v400
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v413 = v400 + int32(1)
	if v413 == v401 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v416 = v413
	goto L104
L104:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v404+v416<<(uint(int32(2))%32))))
	if v423 != 0 {
		v433 = v423
		v434 = v416
		goto L98
	} else {
		goto L106
	}
L105:
	;
	goto L99
L106:
	;
	v425 = v416 + int32(1)
	if v425 != v401 {
		v416 = v425
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	goto L95
L109:
	;
	if v461 != int32(322) {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v470 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v471 + v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v475 + v476
	goto L5
L112:
	;
	v492 = v487
	goto L3
L113:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v505 - int32(1)
	return v503
}
func F_OpenTransientFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFile[0]))
	v5 = F_OpenTransientFilePerm(m, l0, l1, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_OwnLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
			F_errmsg_internal(m, int32(_a_F_OwnLatch_0), v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_OwnLatch_1), int32(135), int32(_a_F_OwnLatch_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_OwnLatch[0]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
		m.G0 = v6 + int32(16)
		return
	}
}
func F___openlog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v1 = int32(0)
	v6 = F_socket(m, int32(1), int32(_a_F___openlog_0), v1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F___openlog[0])) = v6
	if v1 <= v6 {
		v10 = F_connect(m, v6)
		mBase = m.M
	} else {
	}
	return
}
func F___overflow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v2 = l1
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 != 0 {
		v36 = v10
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v36 == v37 {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, v7+int32(15), int32(1))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				if v51 != int32(1) {
				} else {
				}
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v39 == v2&int32(255) {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, v7+int32(15), int32(1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					if v51 != int32(1) {
					} else {
					}
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v2)
				m.G0 = v7 + int32(16)
				return
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v12 - int32(1) | v12
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v17&int32(8) != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17 | int32(32)
			v34 = int32(-1)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26 + v29
			v34 = int32(0)
		}
		if v34 != 0 {
			m.G0 = v7 + int32(16)
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v36 = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v36 == v37 {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, v7+int32(15), int32(1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					if v51 != int32(1) {
					} else {
					}
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v39 == v2&int32(255) {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, v7+int32(15), int32(1))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						if v51 != int32(1) {
						} else {
						}
						m.G0 = v7 + int32(16)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v2)
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_offsethash_insert_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_offsethash_insert_hash_internal(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_offsethash_stat(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	Fn13976(m, l0, int32(_a_F_offsethash_stat_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_oidvectorgt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_oidvectorin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = F_palloc0(m, int32(152))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = v17
	v24 = v13
	v26 = int32(0)
	v28 = int32(32)
	goto L3
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(base.Ui32(v30-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v30 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v89
L5:
	;
	v39 = v24 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v39
	v24 = v39
	goto L3
L6:
	;
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	m.G0 = v11 + int32(16)
	goto L7
L9:
	;
	if v28 <= v26 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(26)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26<<(uint(int32(4))%32) + int32(96)
	v89 = v23
	goto L8
L12:
	;
	v46 = F_repalloc(m, v23, v28<<(uint(int32(3))%32)+int32(24))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v51 = v23
	v52 = v24
	v53 = v28
	goto L14
L14:
	;
	v60 = F_uint32in_subr(m, v52, v11+int32(12), int32(_a_F_oidvectorin_0), v15)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v51 = v46
	v52 = v50
	v53 = v28 << (uint(int32(1)) % 32)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v26<<(uint(int32(2))%32))+24)) = v60
	if v15 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v23 = v51
	v24 = v76
	v26 = v26 + int32(1)
	v28 = v53
	goto L3
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v65 != int32(447) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
	if v68 != int32(1) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
	v89 = int32(0)
	goto L8
}
func F_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(192)
	m.G0 = v15
	v19 = F_make_oper_cache_key(m, l0, v15+int32(4), l1, l2, l3, l5)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(192)
	return v165
L2:
	;
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	return int32(0)
L4:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v50 = v26
	goto L8
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+156)) = int64(601295421576)
	v35 = F_hash_create(m, int32(_a_F_oper_0), int32(256), v15+int32(140), int32(40))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v53 = int32(0)
	v55 = F_hash_search(m, v50, v15+int32(4), v53, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_oper[0])) = v35
	F_CacheRegisterSyscacheCallback(m, int32(39), int32(491), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(12), int32(491), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	v50 = v49
	goto L8
L12:
	;
	if v55 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+136))
	if v59 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v63 = F_SearchSysCache1(m, int32(40), v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v63 != 0 {
		v165 = v63
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L2
L17:
	;
	if l4 != 0 {
		v165 = int32(0)
		goto L1
	} else {
		goto L66
	}
L18:
	;
	v139 = F_SearchSysCache1(m, int32(40), v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L62
	}
L19:
	;
	v102 = F_OpernameGetCandidates(m, l1, int32(98), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L42
	}
L20:
	;
	v132 = v95
	v135 = v97
	v136 = l3
	v137 = v7
	goto L18
L21:
	;
	v88 = F_getBaseType(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L38
	}
L22:
	;
	if v78 != 0 {
		v95 = l2
		v97 = v78
		goto L20
	} else {
		goto L37
	}
L23:
	;
	v69 = base.B2i32(l2 == int32(705))
	goto L25
L24:
	;
	v69 = int32(0)
	goto L25
L25:
	;
	if v69 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v76 = base.B2i32(l2 == int32(0)) | base.B2i32(l3 != int32(705))
	if v76 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v80 = F_OpernameGetOprid(m, l1, l3, l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L35
	}
L29:
	;
	v77 = l3
	goto L31
L30:
	;
	v77 = l2
	goto L31
L31:
	;
	v78 = F_OpernameGetOprid(m, l1, l2, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v76 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	if v78 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v86 = l2
	goto L21
L35:
	;
	if v80 == int32(0) {
		v86 = l3
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v95 = int32(705)
	v97 = v80
	goto L20
L37:
	;
	goto L19
L38:
	;
	if v86 == v88 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v91 = F_OpernameGetOprid(m, l1, v88, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v91 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v95 = l2
	v97 = v91
	goto L20
L42:
	;
	if v102 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = l2
	v158 = l3
	v159 = v7
	goto L17
L44:
	;
	goto L45
L45:
	;
	if l3 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v106 = l3
	goto L48
L47:
	;
	v106 = l2
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v106
	if l2 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v108 = l2
	goto L51
L50:
	;
	v108 = l3
	goto L51
L51:
	;
	if l3 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v109 = v108
	goto L54
L53:
	;
	v109 = l2
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v15)+188)) = v102
	v117 = F_func_match_argtypes(m, int32(2), v15+int32(140), v102, v15+int32(188))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L58
	}
L55:
	;
	v128 = int32(2)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	if v129 == int32(0) {
		v154 = v109
		v158 = v106
		v159 = v128
		goto L17
	} else {
		goto L61
	}
L56:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v127 = v126
	goto L55
L57:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v123 = F_func_select_candidate(m, int32(2), v15+int32(140), v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	switch v117 {
	case 0:
		v154 = v109
		v158 = v106
		v159 = v117
		goto L17
	case 1:
		goto L56
	default:
		goto L57
	}
L59:
	;
	if v123 != 0 {
		v127 = v123
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v154 = v109
	v158 = v106
	v159 = int32(1)
	goto L17
L61:
	;
	v132 = v109
	v135 = v129
	v136 = v106
	v137 = v128
	goto L18
L62:
	;
	if v139 == int32(0) {
		v154 = v132
		v158 = v136
		v159 = v137
		goto L17
	} else {
		goto L63
	}
L63:
	;
	if v19 == int32(0) {
		v165 = v139
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	v151 = F_hash_search(m, v146, v15+int32(4), int32(1), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+136)) = v135
	v165 = v139
	goto L1
L66:
	;
	F_op_error(m, l0, l1, v154, v158, v159, l5)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ordered_set_startup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
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
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = v15 + int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v3 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L37
	} else {
		goto L116
	}
L2:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v413
	v416 = F_palloc(m, int32(32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L37
	} else {
		goto L101
	}
L3:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v398)+16)) = v83
	v402 = v83
	v410 = v77
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v372
	v383 = F_MakeTupleTableSlot(m, v372, int32(_a_F_ordered_set_startup_0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L37
	} else {
		goto L100
	}
L5:
	;
	v367 = F_ExecTypeFromTL(m, v216)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L37
	} else {
		goto L99
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L37
	} else {
		goto L96
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L37
	} else {
		goto L93
	}
L8:
	;
	if v48 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v48 = v45
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
	v45 = v41
	goto L9
L11:
	;
	v37 = int32(0)
	if v18 == v37 {
		v45 = v37
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	switch v23 - int32(429) {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L11
	}
L13:
	;
	if v18 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v18 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = int32(1)
	goto L8
L16:
	;
	goto L17
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v40 = v30
	v41 = int32(1)
	goto L10
L18:
	;
	v48 = int32(2)
	goto L8
L19:
	;
	goto L20
L20:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
	v40 = v35
	v41 = int32(2)
	goto L10
L21:
	;
	v40 = v37
	v41 = v3
	goto L10
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v52 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L37
	} else {
		goto L90
	}
L25:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0]))
	v402 = v52
	v410 = v54
	goto L2
L26:
	;
	goto L27
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v55 == int32(0) {
		v70 = v3
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v70 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L29:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v58 != int32(429) {
		v70 = v3
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+172))
	if v61 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+176))
	if v64 == int32(0) {
		v70 = v3
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v67 = v61
	goto L33
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v70 = v68
	goto L28
L34:
	;
	v67 = v64
	goto L33
L35:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+50)))
	if v73 == int32(110) {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v76 = int32(_a_F_ordered_set_startup_1)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v80
	v83 = F_palloc0(m, int32(104))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v70
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v90 == int32(0) {
		v111 = int32(1)
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v113 = v111 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+12)) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v70)+36))
	if v115 != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != int32(429) {
		v111 = int32(1)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)+172))
	if v97 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
	v111 = v108
	goto L39
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+152))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v107 = v98 + v99*int32(224)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v90)+176))
	if v104 == int32(0) {
		v111 = int32(1)
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v107 = v104
	goto L42
L47:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v303 = F_get_sortgroupclause_tle(m, v301, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L37
	} else {
		goto L86
	}
L48:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+50)))
	v128 = v124 + base.B2i32(v125 == int32(104))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v128
	v132 = F_palloc(m, v128<<(uint(int32(1))%32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L37
	} else {
		goto L56
	}
L49:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if l1 != 0 {
		v124 = v116
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	if v116 != int32(1) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+50)))
	if v119 != int32(104) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	goto L1
L55:
	;
	v124 = v3
	goto L48
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v132
	v136 = v128 << (uint(int32(2)) % 32)
	v137 = F_palloc(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L37
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v137
	v140 = F_palloc(m, v136)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L37
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = v140
	v143 = F_palloc(m, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+40)) = v143
	v146 = F_palloc(m, v128)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L37
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+44)) = v146
	if v115 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	if v125 != int32(104) {
		goto L5
	} else {
		goto L71
	}
L62:
	;
	v208 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v152 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v153 <= v152 {
		v208 = v152
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v160 = v152
	goto L66
L66:
	;
	v169 = v160 << (uint(int32(2)) % 32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169+v170)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v174 = F_get_sortgroupclause_tle(m, v172, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L37
	} else {
		goto L68
	}
L67:
	;
	v208 = v201
	goto L61
L68:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v176+v160<<(uint(int32(1))%32)))) = uint16(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v169))) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v186+v169))) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v191 = F_exprCollation(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L37
	} else {
		goto L69
	}
L69:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v169))) = v191
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v196+v160))) = uint8(v198)
	v201 = v160 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v201 < v202 {
		v160 = v201
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v219 = int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	if v216 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
	v228 = v224 + int32(1)
	goto L74
L73:
	;
	v228 = int32(1)
	goto L74
L74:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v220+v208<<(uint(v219)%32)))) = uint16(v228)
	v231 = v208 << (uint(int32(2)) % 32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v232))) = int32(97)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v231))) = int32(96)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	v242 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v240+v231))) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v244+v208))) = uint8(v242)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v249 = F_ExecTypeFromTL(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L37
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v249
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v254 = v252 + int32(1)
	v255 = F_CreateTemplateTupleDesc(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L37
	} else {
		goto L76
	}
L76:
	;
	if int32(0) < v252 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v264 = v219
	goto L80
L78:
	;
	goto L79
L79:
	;
	F_TupleDescInitEntry(m, v255, base.I32_extend16_s(v254), int32(_a_F_ordered_set_startup_2), int32(23), int32(-1), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L37
	} else {
		goto L84
	}
L80:
	;
	v271 = base.I32_extend16_s(v264)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	F_TupleDescCopyEntry(m, v255, v271, v272, v271)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L37
	} else {
		goto L82
	}
L81:
	;
	goto L79
L82:
	;
	v276 = v264 + int32(1)
	if v276 <= v252 {
		v264 = v276
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	F_FreeTupleDesc(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L37
	} else {
		goto L85
	}
L85:
	;
	v372 = v255
	goto L4
L86:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v306 = F_exprType(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L37
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+52)) = v306
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+60)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+64)) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v314 = F_exprCollation(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L37
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+68)) = v314
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+72)) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v83)+52))
	F_get_typlenbyvalalign(m, v319, v83+int32(56), v83+int32(58), v83+int32(59))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L37
	} else {
		goto L89
	}
L89:
	;
	goto L3
L90:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_3), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L37
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_4), int32(127), int32(_a_F_ordered_set_startup_5))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L37
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_3), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L37
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_4), int32(144), int32(_a_F_ordered_set_startup_5))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L37
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_6), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L37
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_4), int32(146), int32(_a_F_ordered_set_startup_5))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L37
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v372 = v367
	goto L4
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v383
	goto L3
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v402
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = v419
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+12)))
	if l1 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+24)) = uint8(v442)
	*(*int64)(unsafe.Add(mBase, uint32(v416)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+8)) = v441
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v447 == v442 {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v402)+16))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v402)+24))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v402)+32))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v402)+40))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v402)+44))
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[1]))
	v431 = F_tuplesort_begin_heap(m, v422, v423, v424, v425, v426, v427, v429, int32(0), v421)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L37
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v402)+52))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v402)+60))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v402)+68))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+72)))
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[1]))
	v439 = F_tuplesort_begin_datum(m, v433, v434, v435, v436, v438, v421)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L37
	} else {
		goto L107
	}
L106:
	;
	v441 = v431
	goto L102
L107:
	;
	v441 = v439
	goto L102
L108:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v410
	m.G0 = v15 + int32(16)
	return v416
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L37
	} else {
		goto L113
	}
L110:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if v450 != int32(429) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447)+168))
	F_RegisterExprContextCallback(m, v453, int32(1455), v416)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L37
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_7), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L37
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_8), int32(_a_F_ordered_set_startup_9), int32(_a_F_ordered_set_startup_10))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L37
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_11), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L37
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_4), int32(251), int32(_a_F_ordered_set_startup_5))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L37
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_output_plugin_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v18 = v16 + int32(137)
		v20 = v16 + int32(24)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v11 != int64(0) {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*uint32)(unsafe.Add(mBase, uint32(v9)+32)) = uint32(v24)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v20
			v30 = int64(base.Ui64(v24) >> (uint(int64(32)) % 64))
			*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)) = uint32(v30)
			F_errcontext_msg(m, int32(_a_F_output_plugin_error_callback_0), v9+int32(16))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				m.G0 = v9 + int32(48)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
			F_errcontext_msg(m, int32(_a_F_output_plugin_error_callback_1), v9)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				m.G0 = v9 + int32(48)
				return
			}
		}
	}
}
