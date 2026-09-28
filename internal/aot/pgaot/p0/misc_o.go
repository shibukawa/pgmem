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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
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
	v503 = F_query_tree_walker_impl(m, l0, int32(1126), l1, int32(0))
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
	v487 = F_expression_tree_walker_impl(m, l0, int32(1126), l1)
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
	if v10 == int32(321) {
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
	v33 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v36 <= v33 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v39 = v25 + int32(8)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v46 = v43 & int32(-1)
	if v46 != 0 {
		v68 = v46
		v69 = v33
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v47 = int32(1)
	if v47 == v36 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v51 = v47
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
	v105 = int32(base.Ui32(v103) >> (uint(int32(5)) % 32))
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
	v193 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v196 <= v193 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v199 = v185 + int32(8)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v206 = v203 & int32(-1)
	if v206 != 0 {
		v228 = v206
		v229 = v193
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v207 = int32(1)
	if v207 == v196 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v211 = v207
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
	v263 = int32(base.Ui32(v261) >> (uint(int32(5)) % 32))
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
	v330 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v333 <= v330 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v336 = v322 + int32(8)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v343 = v340 & int32(-1)
	if v343 != 0 {
		v365 = v343
		v366 = v330
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v344 = int32(1)
	if v344 == v333 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v348 = v344
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
	v400 = int32(base.Ui32(v398) >> (uint(int32(5)) % 32))
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
	if v461 != int32(324) {
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
		F_errstart_cold(m, int32(24), int32(0))
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
	Fn14356(m, l0, int32(_a_F_offsethash_stat_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_oid8out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 == int64(0) {
		v18 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v18)
		v189 = int32(1)
	} else {
		v25 = int32(1233)
		v30 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v8))^int32(63))*v25+v25) >> (uint(int32(12)) % 32))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v30<<(uint(int32(3))%32))+uint32(_c_F_oid8out[0])))
		v35 = v30 + base.B2i32(base.Ui64(v33) <= base.Ui64(v8))
		if base.Ui64(int64(100000000)) <= base.Ui64(v8) {
			v39 = v8
			v42 = v2
			for {
				v48 = v6 + v35 - v42
				v49 = int32(8)
				v52 = base.I64_div_u_s(v39, int64(100000000))
				v56 = base.I32_wrap_i64(v39 + v52*int64(4194967296))
				v58 = base.I32_div_u_s(v56, int32(_a_F_oid8out_0))
				v59 = int32(1)
				v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58<<(uint(v59)%32))+uint32(_c_F_oid8out[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v48-v49))) = uint16(v61)
				v65 = int32(_a_F_oid8out_1)
				v66 = base.I32_div_u_s(v56, v65)
				v67 = int32(100)
				v68 = base.I32_rem_u_s(v66, v67)
				v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68<<(uint(v59)%32))+uint32(_c_F_oid8out[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v48-int32(6)))) = uint16(v71)
				v77 = v56 - v66*v65
				v78 = int32(_a_F_oid8out_2)
				v81 = base.I32_div_u_s(v77&v78, v67)
				v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81<<(uint(v59)%32))+uint32(_c_F_oid8out[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v48-int32(4)))) = uint16(v84)
				v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v77-v81*v67)&v78<<(uint(v59)%32))+uint32(_c_F_oid8out[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v48-int32(2)))) = uint16(v95)
				v98 = v42 + v49
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v39) {
					v39 = v52
					v42 = v98
					continue
				} else {
					break
				}
				break
			}
			v101 = v52
			v104 = v98
		} else {
			v101 = v8
			v104 = v2
		}
		v110 = base.I32_wrap_i64(v101)
		if base.Ui64(int64(10000)) <= base.Ui64(v101) {
			v114 = v6 + v35 - v104
			v115 = int32(4)
			v118 = base.I32_div_u_s(v110, int32(_a_F_oid8out_1))
			v121 = v110 + v118*int32(-10000)
			v122 = int32(100)
			v123 = base.I32_div_u_s(v121, v122)
			v124 = int32(1)
			v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123<<(uint(v124)%32))+uint32(_c_F_oid8out[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v114-v115))) = uint16(v126)
			v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v121-v123*v122)<<(uint(v124)%32))+uint32(_c_F_oid8out[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v114-int32(2)))) = uint16(v135)
			v139 = v118
			v140 = v104 | v115
		} else {
			v139 = v110
			v140 = v104
		}
		if base.Ui32(int32(100)) <= base.Ui32(v139) {
			v148 = int32(2)
			v150 = int32(_a_F_oid8out_2)
			v152 = int32(100)
			v153 = base.I32_div_u_s(v139&v150, v152)
			v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v139-v153*v152)&v150<<(uint(int32(1))%32))+uint32(_c_F_oid8out[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v6+v35-v140-v148))) = uint16(v161)
			v165 = v153
			v166 = v140 + v148
		} else {
			v165 = v139
			v166 = v140
		}
		if base.Ui32(int32(10)) <= base.Ui32(v165) {
			v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165<<(uint(int32(1))%32))+uint32(_c_F_oid8out[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v6+v35-v166-int32(2)))) = uint16(v175)
			v189 = v35
		} else {
			v178 = v165 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v178)
			v189 = v35
		}
	}
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6+v189))) = uint8(v191)
	v194 = v189 + int32(1)
	v195 = F_palloc(m, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		return int64(0)
	} else {
		if v194 != 0 {
			base.MemoryCopy(m, v195, v6, v194)
		} else {
		}
		m.G0 = v6 + int32(32)
		return base.I64_extend_i32_u(v195)
	}
}
func F_oidvectorgt(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < base.I32_wrap_i64(v2)))
	}
}
func F_oidvectorin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v95 int64
	_ = v95
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = F_palloc0(m, int32(152))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = v18
	v25 = v14
	v27 = int32(0)
	v29 = int32(32)
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if base.B2i32(base.Ui32(v32-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v32 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v95
L5:
	;
	v41 = v25 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v41
	v25 = v41
	goto L3
L6:
	;
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	m.G0 = v12 + int32(16)
	goto L7
L9:
	;
	if v29 <= v27 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(26)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27<<(uint(int32(4))%32) + int32(96)
	v95 = base.I64_extend_i32_u(v24)
	goto L8
L12:
	;
	v48 = F_repalloc(m, v24, v29<<(uint(int32(3))%32)+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v53 = v24
	v54 = v25
	v55 = v29
	goto L14
L14:
	;
	v62 = F_uint32in_subr(m, v54, v12+int32(12), int32(_a_F_oidvectorin_0), v16)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v53 = v48
	v54 = v52
	v55 = v29 << (uint(int32(1)) % 32)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+v27<<(uint(int32(2))%32))+24)) = v62
	if v16 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v24 = v53
	v25 = v78
	v27 = v27 + int32(1)
	v29 = v55
	goto L3
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v67 != int32(453) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
	if v70 != int32(1) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v73)
	v95 = int64(0)
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v7
	v21 = F_make_oper_cache_key(m, l0, v15+int32(16), l1, l2, l3, l5)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(208)
	return v172
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
	if v21 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v52 = v28
	goto L8
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+160)) = int64(601295421576)
	v37 = F_hash_create(m, int32(_a_F_oper_0), int64(256), v15+int32(152), int32(40))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v55 = int32(0)
	v57 = F_hash_search(m, v52, v15+int32(16), v55, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_oper[0])) = v37
	F_CacheRegisterSyscacheCallback(m, int32(39), int32(526), int64(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(12), int32(526), int64(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	v52 = v51
	goto L8
L12:
	;
	if v57 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+136))
	if v61 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v66 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v61))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v66 != 0 {
		v172 = v66
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
		v172 = int32(0)
		goto L1
	} else {
		goto L66
	}
L18:
	;
	v145 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v140))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L62
	}
L19:
	;
	v107 = F_OpernameGetCandidates(m, l1, int32(98), int32(0), v15+int32(12))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L42
	}
L20:
	;
	v137 = v98
	v140 = v100
	v141 = l3
	v142 = v7
	goto L18
L21:
	;
	v91 = F_getBaseType(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L38
	}
L22:
	;
	if v81 != 0 {
		v98 = l2
		v100 = v81
		goto L20
	} else {
		goto L37
	}
L23:
	;
	v72 = base.B2i32(l2 == int32(705))
	goto L25
L24:
	;
	v72 = int32(0)
	goto L25
L25:
	;
	if v72 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = base.B2i32(l2 == int32(0)) | base.B2i32(l3 != int32(705))
	if v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v83 = F_OpernameGetOprid(m, l1, l3, l3)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L35
	}
L29:
	;
	v80 = l3
	goto L31
L30:
	;
	v80 = l2
	goto L31
L31:
	;
	v81 = F_OpernameGetOprid(m, l1, l2, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v79 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	if v81 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v89 = l2
	goto L21
L35:
	;
	if v83 == int32(0) {
		v89 = l3
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v98 = int32(705)
	v100 = v83
	goto L20
L37:
	;
	goto L19
L38:
	;
	if v89 == v91 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v94 = F_OpernameGetOprid(m, l1, v91, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v94 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v98 = l2
	v100 = v94
	goto L20
L42:
	;
	if v107 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v160 = l2
	v164 = l3
	v165 = v7
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
	v111 = l3
	goto L48
L47:
	;
	v111 = l2
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v111
	if l2 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v113 = l2
	goto L51
L50:
	;
	v113 = l3
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
	v114 = v113
	goto L54
L53:
	;
	v114 = l2
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v107
	v122 = F_func_match_argtypes(m, int32(2), v15+int32(152), v107, v15+int32(204))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L58
	}
L55:
	;
	v133 = int32(2)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	if v134 == int32(0) {
		v160 = v114
		v164 = v111
		v165 = v133
		goto L17
	} else {
		goto L61
	}
L56:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	v132 = v131
	goto L55
L57:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	v128 = F_func_select_candidate(m, int32(2), v15+int32(152), v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	switch v122 {
	case 0:
		v160 = v114
		v164 = v111
		v165 = v122
		goto L17
	case 1:
		goto L56
	default:
		goto L57
	}
L59:
	;
	if v128 != 0 {
		v132 = v128
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v160 = v114
	v164 = v111
	v165 = int32(1)
	goto L17
L61:
	;
	v137 = v114
	v140 = v134
	v141 = v111
	v142 = v133
	goto L18
L62:
	;
	if v145 == int32(0) {
		v160 = v137
		v164 = v141
		v165 = v142
		goto L17
	} else {
		goto L63
	}
L63:
	;
	if v21 == int32(0) {
		v172 = v145
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_oper[0]))
	v157 = F_hash_search(m, v152, v15+int32(16), int32(1), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+136)) = v140
	v172 = v145
	goto L1
L66:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_op_error(m, l0, l1, v160, v164, v165, v167, l5)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = v16 + int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 == v3 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L38
	} else {
		goto L136
	}
L2:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v512
	v515 = F_palloc(m, int32(32))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L38
	} else {
		goto L121
	}
L3:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v496)+16)) = v88
	v500 = v88
	v508 = v82
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v468
	v480 = F_MakeSingleTupleTableSlot(m, v468, int32(_a_F_ordered_set_startup_0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L38
	} else {
		goto L120
	}
L5:
	;
	v463 = F_ExecTypeFromTL(m, v223)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L38
	} else {
		goto L119
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L38
	} else {
		goto L116
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L38
	} else {
		goto L113
	}
L8:
	;
	if v49 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v49 = v46
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
	v46 = v42
	goto L9
L11:
	;
	v38 = int32(0)
	if v19 == v38 {
		v46 = v38
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	switch v24 - int32(435) {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L11
	}
L13:
	;
	if v19 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v19 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = int32(1)
	goto L8
L16:
	;
	goto L17
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v41 = v31
	v42 = int32(1)
	goto L10
L18:
	;
	v49 = int32(2)
	goto L8
L19:
	;
	goto L20
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+376))
	v41 = v36
	v42 = int32(2)
	goto L10
L21:
	;
	v41 = v38
	v42 = v3
	goto L10
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v53 != 0 {
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
	v427 = m.ExcPending
	if v427 != 0 {
		goto L38
	} else {
		goto L110
	}
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0]))
	v500 = v53
	v508 = v55
	goto L2
L26:
	;
	goto L27
L27:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v57 == v56 {
		v69 = v56
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v75 == int32(0) {
		goto L7
	} else {
		goto L36
	}
L29:
	;
	v75 = v69
	goto L28
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v60 != int32(435) {
		v69 = v56
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+172))
	if v63 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v64
	goto L28
L33:
	;
	goto L34
L34:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+176))
	if v65 == int32(0) {
		v69 = v56
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v69 = v68
	goto L29
L36:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+50)))
	if v78 == int32(110) {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v81 = int32(_a_F_ordered_set_startup_1)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v85
	v88 = F_palloc0(m, int32(104))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v75
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v95 == int32(0) {
		v116 = int32(1)
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v118 = v116 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+12)) = uint8(v118)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
	if v120 != 0 {
		goto L50
	} else {
		goto L51
	}
L41:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v99 != int32(435) {
		v116 = int32(1)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+172))
	if v102 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+4)))
	v116 = v113
	goto L40
L44:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95)+152))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v112 = v103 + v104*int32(240)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+176))
	if v109 == int32(0) {
		v116 = int32(1)
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v112 = v109
	goto L43
L48:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	v399 = F_get_sortgroupclause_tle(m, v397, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L38
	} else {
		goto L106
	}
L49:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+50)))
	v133 = v129 + base.B2i32(v130 == int32(104))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v133
	v137 = F_palloc(m, v133<<(uint(int32(1))%32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L38
	} else {
		goto L57
	}
L50:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if l1 != 0 {
		v129 = v121
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	if v121 != int32(1) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+50)))
	if v124 != int32(104) {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	goto L1
L56:
	;
	v129 = v3
	goto L49
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v137
	v141 = v133 << (uint(int32(2)) % 32)
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L38
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v142
	v145 = F_palloc(m, v141)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = v145
	v148 = F_palloc(m, v141)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L38
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = v148
	v151 = F_palloc(m, v133)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v151
	if v120 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	if v130 != int32(104) {
		goto L5
	} else {
		goto L72
	}
L63:
	;
	v213 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v157 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v158 <= v157 {
		v213 = v157
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v164 = v157
	goto L67
L67:
	;
	v175 = v164 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175+v176)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	v180 = F_get_sortgroupclause_tle(m, v178, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L38
	} else {
		goto L69
	}
L68:
	;
	v213 = v207
	goto L62
L69:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v182+v164<<(uint(int32(1))%32)))) = uint16(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v88)+32))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v188+v175))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v192+v175))) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v197 = F_exprCollation(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L38
	} else {
		goto L70
	}
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v88)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v199+v175))) = v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v202+v164))) = uint8(v204)
	v207 = v164 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v207 < v208 {
		v164 = v207
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v226 = int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
	if v223 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223)+4)))
	v235 = v231 + int32(1)
	goto L75
L74:
	;
	v235 = int32(1)
	goto L75
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v227+v213<<(uint(v226)%32)))) = uint16(v235)
	v238 = v213 << (uint(int32(2)) % 32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v88)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v238+v239))) = int32(97)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v243+v238))) = int32(96)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v88)+40))
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v247+v238))) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v251+v213))) = uint8(v249)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	v256 = F_ExecTypeFromTL(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L38
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v261 = v259 + int32(1)
	v262 = F_CreateTemplateTupleDesc(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	if int32(0) < v259 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v271 = v226
	goto L81
L79:
	;
	goto L80
L80:
	;
	F_TupleDescInitEntry(m, v262, base.I32_extend16_s(v261), int32(_a_F_ordered_set_startup_2), int32(23), int32(-1), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L38
	} else {
		goto L85
	}
L81:
	;
	v279 = base.I32_extend16_s(v271)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	F_TupleDescCopyEntry(m, v262, v279, v280, v279)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	v284 = v271 + int32(1)
	if v284 <= v259 {
		v271 = v284
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v306 = int32(0)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v306 < v315 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	F_FreeTupleDesc(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L38
	} else {
		goto L105
	}
L87:
	;
	v319 = v262 + int32(28)
	v326 = v306
	v327 = v315
	v329 = v306
	goto L91
L88:
	;
	v383 = v306
	v390 = v315
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v262)+16)) = v383
	goto L86
L90:
	;
	v383 = v377
	v390 = v356
	goto L89
L91:
	;
	v335 = v319 + v315<<(uint(int32(3))%32) + v326*int32(100)
	v338 = v319 + v326<<(uint(int32(3))%32)
	if v315 != v327 {
		v356 = v327
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v377 = v315
	goto L90
L93:
	;
	v357 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338)+2)))
	if v357 <= int32(0) {
		v377 = v326
		goto L90
	} else {
		goto L101
	}
L94:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+7)))
	if v340 != int32(118) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v356 = v326
	goto L93
L96:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+4)))
	if v343 != int32(1) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+6)))
	if v346&int32(6) != 0 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v349 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338)+2)))
	if v349 <= int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+90)))
	if v352 != int32(118) {
		v356 = v315
		goto L93
	} else {
		goto L100
	}
L100:
	;
	goto L95
L101:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+90)))
	if v360 == int32(118) {
		v377 = v326
		goto L90
	} else {
		goto L102
	}
L102:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+5)))
	v369 = (v329 + v363 - int32(1)) & (int32(0) - v363)
	if int32(_a_F_ordered_set_startup_3) < v369 {
		v377 = v326
		goto L90
	} else {
		goto L103
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v338))) = uint16(v369)
	v375 = v326 + int32(1)
	if v375 != v315 {
		v326 = v375
		v327 = v356
		v329 = v369 + v357
		goto L91
	} else {
		goto L104
	}
L104:
	;
	goto L92
L105:
	;
	v468 = v262
	goto L4
L106:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v402 = F_exprType(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L38
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+52)) = v402
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+60)) = v405
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+64)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v410 = F_exprCollation(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L38
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+68)) = v410
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+72)) = uint8(v413)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v88)+52))
	F_get_typlenbyvalalign(m, v415, v88+int32(56), v88+int32(58), v88+int32(59))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L38
	} else {
		goto L109
	}
L109:
	;
	goto L3
L110:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_4), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L38
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_5), int32(127), int32(_a_F_ordered_set_startup_6))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L38
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_4), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L38
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_5), int32(144), int32(_a_F_ordered_set_startup_6))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L38
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
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_7), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L38
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_5), int32(146), int32(_a_F_ordered_set_startup_6))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L38
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v468 = v463
	goto L4
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = v480
	goto L3
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v500
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+4)) = v518
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+12)))
	if l1 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v541 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v515)+24)) = uint8(v541)
	*(*int64)(unsafe.Add(mBase, uint32(v515)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v515)+8)) = v540
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v547 == v541 {
		goto L129
	} else {
		goto L130
	}
L123:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v500)+28))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v500)+32))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v500)+40))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v500)+44))
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[1]))
	v530 = F_tuplesort_begin_heap(m, v521, v522, v523, v524, v525, v526, v528, int32(0), v520)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L38
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v500)+52))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v500)+60))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v500)+68))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+72)))
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[1]))
	v538 = F_tuplesort_begin_datum(m, v532, v533, v534, v535, v537, v520)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L38
	} else {
		goto L127
	}
L126:
	;
	v540 = v530
	goto L122
L127:
	;
	v540 = v538
	goto L122
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ordered_set_startup[0])) = v508
	m.G0 = v16 + int32(16)
	return v515
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L38
	} else {
		goto L133
	}
L130:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	if v550 != int32(435) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v547)+168))
	F_RegisterExprContextCallback(m, v553, int32(1601), base.I64_extend_i32_u(v515))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L38
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_8), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L38
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_9), int32(_a_F_ordered_set_startup_10), int32(_a_F_ordered_set_startup_11))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L38
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(_a_F_ordered_set_startup_12), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L38
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_ordered_set_startup_5), int32(252), int32(_a_F_ordered_set_startup_6))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L38
	} else {
		goto L138
	}
L138:
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
