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
	var v89 int32
	_ = v89
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
	var v159 int32
	_ = v159
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	v3 = int32(0)
	if l0 == v3 {
		v504 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v509 + int32(1)
	v515 = F_query_tree_walker_impl(m, l0, int32(1047), l1, int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L30
	} else {
		goto L112
	}
L2:
	;
	return v504
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(58) {
	case 0:
		goto L9
	case 1, 2, 3, 4:
		v473 = v10
		goto L5
	case 5:
		goto L8
	case 6:
		goto L7
	default:
		goto L10
	}
L4:
	;
	v499 = F_expression_tree_walker_impl(m, l0, int32(1047), l1)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L30
	} else {
		goto L111
	}
L5:
	;
	if v473 == int32(67) {
		goto L1
	} else {
		goto L108
	}
L6:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v192 != v193 {
		goto L4
	} else {
		goto L49
	}
L7:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v185 == int32(0) {
		goto L4
	} else {
		goto L47
	}
L8:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v178 != 0 {
		v504 = v3
		goto L2
	} else {
		goto L46
	}
L9:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v171 != 0 {
		v504 = v3
		goto L2
	} else {
		goto L45
	}
L10:
	;
	if v10 == int32(319) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v10 != int32(6) {
		v473 = v10
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 != v18 {
		v504 = v3
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v25 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if int32(0) <= v82 {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v82 = base.I32_ctz(v68) | v69<<(uint(int32(5))%32)
	goto L14
L16:
	;
	v82 = int32(-2)
	goto L14
L17:
	;
	v35 = base.I32_div_s(int32(0), int32(32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v36 <= v35 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v39 = v25 + int32(8)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v35<<(uint(int32(2))%32))))
	v46 = v43 & int32(-1)
	if v46 != 0 {
		v68 = v46
		v69 = v35
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v48 = v35 + int32(1)
	if v48 == v36 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v51 = v48
	goto L21
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39+v51<<(uint(int32(2))%32))))
	if v58 != 0 {
		v68 = v58
		v69 = v51
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L16
L23:
	;
	v60 = v51 + int32(1)
	if v60 != v36 {
		v51 = v60
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v87 = v82
	v89 = v3
	goto L28
L26:
	;
	v159 = v3
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v159
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v163 == int32(0) {
		v504 = v3
		goto L2
	} else {
		goto L44
	}
L28:
	;
	v93 = F_bms_add_member(m, v89, v87+v24)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v159 = v93
	goto L27
L30:
	;
	return int32(0)
L31:
	;
	if v25 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if int32(0) <= v152 {
		v87 = v152
		v89 = v93
		goto L28
	} else {
		goto L43
	}
L33:
	;
	v152 = base.I32_ctz(v138) | v139<<(uint(int32(5))%32)
	goto L32
L34:
	;
	v152 = int32(-2)
	goto L32
L35:
	;
	v103 = v87 + int32(1)
	v105 = base.I32_div_s(v103, int32(32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v106 <= v105 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v109 = v25 + int32(8)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v105<<(uint(int32(2))%32))))
	v116 = v113 & (int32(-1) << (uint(v103) % 32))
	if v116 != 0 {
		v138 = v116
		v139 = v105
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v118 = v105 + int32(1)
	if v118 == v106 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v121 = v118
	goto L39
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v109+v121<<(uint(int32(2))%32))))
	if v128 != 0 {
		v138 = v128
		v139 = v121
		goto L33
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v130 = v121 + int32(1)
	if v130 != v106 {
		v121 = v130
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L29
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v166 + v163
	return int32(0)
L45:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172 + v173
	return int32(0)
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v179 + v180
	return int32(0)
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v188 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v189 + v185
	goto L4
L49:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v196 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v197 == v196 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if int32(0) <= v254 {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v254 = base.I32_ctz(v240) | v241<<(uint(int32(5))%32)
	goto L50
L52:
	;
	v254 = int32(-2)
	goto L50
L53:
	;
	v207 = base.I32_div_s(int32(0), int32(32))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v208 <= v207 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v211 = v197 + int32(8)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v218 = v215 & int32(-1)
	if v218 != 0 {
		v240 = v218
		v241 = v207
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v220 = v207 + int32(1)
	if v220 == v208 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v223 = v220
	goto L57
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211+v223<<(uint(int32(2))%32))))
	if v230 != 0 {
		v240 = v230
		v241 = v223
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v232 = v223 + int32(1)
	if v232 != v208 {
		v223 = v232
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v259 = v196
	v260 = v254
	goto L64
L62:
	;
	v327 = v196
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v327
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v334 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L64:
	;
	v265 = F_bms_add_member(m, v259, v260+v195)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L30
	} else {
		goto L66
	}
L65:
	;
	v327 = v265
	goto L63
L66:
	;
	if v197 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if int32(0) <= v322 {
		v259 = v265
		v260 = v322
		goto L64
	} else {
		goto L78
	}
L68:
	;
	v322 = base.I32_ctz(v308) | v309<<(uint(int32(5))%32)
	goto L67
L69:
	;
	v322 = int32(-2)
	goto L67
L70:
	;
	v273 = v260 + int32(1)
	v275 = base.I32_div_s(v273, int32(32))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v276 <= v275 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v279 = v197 + int32(8)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v275<<(uint(int32(2))%32))))
	v286 = v283 & (int32(-1) << (uint(v273) % 32))
	if v286 != 0 {
		v308 = v286
		v309 = v275
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v288 = v275 + int32(1)
	if v288 == v276 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v291 = v288
	goto L74
L74:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v279+v291<<(uint(int32(2))%32))))
	if v298 != 0 {
		v308 = v298
		v309 = v291
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v300 = v291 + int32(1)
	if v300 != v276 {
		v291 = v300
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L65
L79:
	;
	if int32(0) <= v391 {
		goto L90
	} else {
		goto L91
	}
L80:
	;
	v391 = base.I32_ctz(v377) | v378<<(uint(int32(5))%32)
	goto L79
L81:
	;
	v391 = int32(-2)
	goto L79
L82:
	;
	v344 = base.I32_div_s(int32(0), int32(32))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v345 <= v344 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v348 = v334 + int32(8)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v344<<(uint(int32(2))%32))))
	v355 = v352 & int32(-1)
	if v355 != 0 {
		v377 = v355
		v378 = v344
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v357 = v344 + int32(1)
	if v357 == v345 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v360 = v357
	goto L86
L86:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v348+v360<<(uint(int32(2))%32))))
	if v367 != 0 {
		v377 = v367
		v378 = v360
		goto L80
	} else {
		goto L88
	}
L87:
	;
	goto L81
L88:
	;
	v369 = v360 + int32(1)
	if v369 != v345 {
		v360 = v369
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v397 = v391
	v398 = v3
	goto L93
L91:
	;
	v466 = v3
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v466
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v473 = v470
	goto L5
L93:
	;
	v402 = F_bms_add_member(m, v398, v333+v397)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L30
	} else {
		goto L95
	}
L94:
	;
	v466 = v402
	goto L92
L95:
	;
	if v334 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if int32(0) <= v459 {
		v397 = v459
		v398 = v402
		goto L93
	} else {
		goto L107
	}
L97:
	;
	v459 = base.I32_ctz(v445) | v446<<(uint(int32(5))%32)
	goto L96
L98:
	;
	v459 = int32(-2)
	goto L96
L99:
	;
	v410 = v397 + int32(1)
	v412 = base.I32_div_s(v410, int32(32))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v413 <= v412 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v416 = v334 + int32(8)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416+v412<<(uint(int32(2))%32))))
	v423 = v420 & (int32(-1) << (uint(v410) % 32))
	if v423 != 0 {
		v445 = v423
		v446 = v412
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v425 = v412 + int32(1)
	if v425 == v413 {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v428 = v425
	goto L103
L103:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v416+v428<<(uint(int32(2))%32))))
	if v435 != 0 {
		v445 = v435
		v446 = v428
		goto L97
	} else {
		goto L105
	}
L104:
	;
	goto L98
L105:
	;
	v437 = v428 + int32(1)
	if v437 != v413 {
		v428 = v437
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L94
L108:
	;
	if v473 != int32(322) {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v482 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483 + v484
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v487 + v488
	goto L4
L111:
	;
	v504 = v499
	goto L2
L112:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v517 - int32(1)
	return v515
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[587]))
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
			F_errmsg_internal(m, int32(471434), v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(489453), int32(135), int32(320424))
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
		v23 = *(*int32)(unsafe.Add(mBase, _consts[522]))
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
	v6 = F_socket(m, int32(1), int32(524290), v1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1122])) = v6
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
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
	if base.Ui32(v30-int32(9)) < base.Ui32(int32(5)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v93 = v24 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v93
	v24 = v93
	goto L3
L6:
	;
	if v30 == int32(32) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v11 + int32(16)
	return v85
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
	v85 = v23
	goto L8
L12:
	;
	v42 = F_repalloc(m, v23, v28<<(uint(int32(3))%32)+int32(24))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v47 = v23
	v48 = v24
	v49 = v28
	goto L14
L14:
	;
	v56 = F_uint32in_subr(m, v48, v11+int32(12), int32(428052), v15)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v47 = v42
	v48 = v46
	v49 = v28 << (uint(int32(1)) % 32)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47+v26<<(uint(int32(2))%32))+24)) = v56
	if v15 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v23 = v47
	v24 = v72
	v26 = v26 + int32(1)
	v28 = v49
	goto L3
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v61 != int32(447) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
	if v64 != int32(1) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v67)
	v85 = int32(0)
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
	var v70 int32
	_ = v70
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
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
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
	return v167
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[242]))
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
	v35 = F_hash_create(m, int32(392647), int32(256), v15+int32(140), int32(40))
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
	*(*int32)(unsafe.Add(mBase, _consts[242])) = v35
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
	v49 = *(*int32)(unsafe.Add(mBase, _consts[242]))
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
		v167 = v63
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
		v167 = int32(0)
		goto L1
	} else {
		goto L66
	}
L18:
	;
	v140 = F_SearchSysCache1(m, int32(40), v135)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L62
	}
L19:
	;
	v103 = F_OpernameGetCandidates(m, l1, int32(98), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L42
	}
L20:
	;
	v133 = v96
	v135 = v97
	v137 = v7
	v138 = l3
	goto L18
L21:
	;
	v89 = F_getBaseType(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L38
	}
L22:
	;
	if v79 != 0 {
		v96 = l2
		v97 = v79
		goto L20
	} else {
		goto L37
	}
L23:
	;
	v70 = base.B2i32(l2 == int32(705))
	goto L25
L24:
	;
	v70 = int32(0)
	goto L25
L25:
	;
	if v70 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = base.B2i32(l2 == int32(0)) | base.B2i32(l3 != int32(705))
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v81 = F_OpernameGetOprid(m, l1, l3, l3)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L35
	}
L29:
	;
	v78 = l3
	goto L31
L30:
	;
	v78 = l2
	goto L31
L31:
	;
	v79 = F_OpernameGetOprid(m, l1, l2, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v77 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	if v79 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v88 = l2
	goto L21
L35:
	;
	if v81 == int32(0) {
		v88 = l3
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v96 = int32(705)
	v97 = v81
	goto L20
L37:
	;
	goto L19
L38:
	;
	if v89 == v88 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v92 = F_OpernameGetOprid(m, l1, v89, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v92 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v96 = l2
	v97 = v92
	goto L20
L42:
	;
	if v103 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v155 = l2
	v159 = v7
	v160 = l3
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
	v107 = l3
	goto L48
L47:
	;
	v107 = l2
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v107
	if l2 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v109 = l2
	goto L51
L50:
	;
	v109 = l3
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
	v110 = v109
	goto L54
L53:
	;
	v110 = l2
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v15)+188)) = v103
	v118 = F_func_match_argtypes(m, int32(2), v15+int32(140), v103, v15+int32(188))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L58
	}
L55:
	;
	v129 = int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v130 == int32(0) {
		v155 = v110
		v159 = v129
		v160 = v107
		goto L17
	} else {
		goto L61
	}
L56:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v128 = v127
	goto L55
L57:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v124 = F_func_select_candidate(m, int32(2), v15+int32(140), v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	switch v118 {
	case 0:
		v155 = v110
		v159 = v118
		v160 = v107
		goto L17
	case 1:
		goto L56
	default:
		goto L57
	}
L59:
	;
	if v124 != 0 {
		v128 = v124
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v155 = v110
	v159 = int32(1)
	v160 = v107
	goto L17
L61:
	;
	v133 = v110
	v135 = v130
	v137 = v129
	v138 = v107
	goto L18
L62:
	;
	if v140 == int32(0) {
		v155 = v133
		v159 = v137
		v160 = v138
		goto L17
	} else {
		goto L63
	}
L63:
	;
	if v19 == int32(0) {
		v167 = v140
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[242]))
	v152 = F_hash_search(m, v147, v15+int32(4), int32(1), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+136)) = v135
	v167 = v140
	goto L1
L66:
	;
	F_op_error(m, l0, l1, v155, v160, v159, l5)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
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
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	var v132 int32
	_ = v132
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
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
	v484 = m.ExcPending
	if v484 != 0 {
		goto L37
	} else {
		goto L116
	}
L2:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v418
	v421 = F_palloc(m, int32(32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L37
	} else {
		goto L101
	}
L3:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v403)+16)) = v83
	v407 = v83
	v415 = v77
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v379
	v388 = F_MakeSingleTupleTableSlot(m, v379, int32(1592204))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L37
	} else {
		goto L100
	}
L5:
	;
	v372 = F_ExecTypeFromTL(m, v221)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L37
	} else {
		goto L99
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L37
	} else {
		goto L96
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L37
	} else {
		goto L90
	}
L25:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v407 = v52
	v415 = v54
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
	v76 = int32(4470560)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v80
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
		v113 = int32(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+12)) = uint8(v118)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v70)+36))
	if v120 != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v118 = v113 & int32(1)
	goto L39
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != int32(429) {
		v113 = int32(1)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)+172))
	if v97 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+152))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v99*int32(224))+4)))
	v118 = v103 & int32(1)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v90)+176))
	if v107 == int32(0) {
		v113 = int32(1)
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
	v113 = v110
	goto L40
L47:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v308 = F_get_sortgroupclause_tle(m, v306, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L37
	} else {
		goto L86
	}
L48:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+50)))
	v132 = base.B2i32(v130 == int32(104))
	v133 = v129 + v132
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v133
	v137 = F_palloc(m, v133<<(uint(int32(1))%32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L37
	} else {
		goto L56
	}
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if l1 != 0 {
		v129 = v121
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
	if v121 != int32(1) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+50)))
	if v124 != int32(104) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	goto L1
L55:
	;
	v129 = v3
	goto L48
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v137
	v141 = v133 << (uint(int32(2)) % 32)
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L37
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v142
	v145 = F_palloc(m, v141)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L37
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = v145
	v148 = F_palloc(m, v141)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+40)) = v148
	v151 = F_palloc(m, v133)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L37
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+44)) = v151
	if v120 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	if v132 == int32(0) {
		goto L5
	} else {
		goto L71
	}
L62:
	;
	v212 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v157 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v158 <= v157 {
		v212 = v157
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v164 = v157
	goto L66
L66:
	;
	v174 = v164 << (uint(int32(2)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174+v175)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v179 = F_get_sortgroupclause_tle(m, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L37
	} else {
		goto L68
	}
L67:
	;
	v212 = v206
	goto L61
L68:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v181+v164<<(uint(int32(1))%32)))) = uint16(v185)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v187+v174))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v174))) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v196 = F_exprCollation(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L37
	} else {
		goto L69
	}
L69:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v198+v174))) = v196
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201+v164))) = uint8(v203)
	v206 = v164 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v206 < v207 {
		v164 = v206
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v224 = int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	if v221 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	v233 = v229 + int32(1)
	goto L74
L73:
	;
	v233 = int32(1)
	goto L74
L74:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v225+v212<<(uint(v224)%32)))) = uint16(v233)
	v236 = v212 << (uint(int32(2)) % 32)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v237))) = int32(97)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v236))) = int32(96)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v245+v236))) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v249+v212))) = uint8(v247)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v254 = F_ExecTypeFromTL(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L37
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v254
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v259 = v257 + int32(1)
	v260 = F_CreateTemplateTupleDesc(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L37
	} else {
		goto L76
	}
L76:
	;
	if int32(0) < v257 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v268 = v224
	goto L80
L78:
	;
	goto L79
L79:
	;
	F_TupleDescInitEntry(m, v260, base.I32_extend16_s(v259), int32(332158), int32(23), int32(-1), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L37
	} else {
		goto L84
	}
L80:
	;
	v276 = base.I32_extend16_s(v268)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	F_TupleDescCopyEntry(m, v260, v276, v277, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L37
	} else {
		goto L82
	}
L81:
	;
	goto L79
L82:
	;
	v281 = v268 + int32(1)
	if v281 <= v257 {
		v268 = v281
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	F_FreeTupleDesc(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L37
	} else {
		goto L85
	}
L85:
	;
	v379 = v260
	goto L4
L86:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v311 = F_exprType(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L37
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+52)) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+60)) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+64)) = v316
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v319 = F_exprCollation(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L37
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+68)) = v319
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+72)) = uint8(v322)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v83)+52))
	F_get_typlenbyvalalign(m, v324, v83+int32(56), v83+int32(58), v83+int32(59))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L37
	} else {
		goto L89
	}
L89:
	;
	goto L3
L90:
	;
	F_errmsg_internal(m, int32(60931), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L37
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(485628), int32(127), int32(228064))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
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
	F_errmsg_internal(m, int32(60931), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L37
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(485628), int32(144), int32(228064))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
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
	F_errmsg_internal(m, int32(348902), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L37
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(485628), int32(146), int32(228064))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
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
	v379 = v372
	goto L4
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v388
	goto L3
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v407
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+4)) = v424
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+12)))
	if l1 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v421)+24)) = uint8(v446)
	*(*int64)(unsafe.Add(mBase, uint32(v421)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+8)) = v445
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v451 == v446 {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v407)+32))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v407)+40))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v407)+44))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v435 = F_tuplesort_begin_heap(m, v427, v428, v429, v430, v431, v432, v434, v426)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L37
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v407)+52))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v407)+60))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v407)+68))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+72)))
	v442 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v443 = F_tuplesort_begin_datum(m, v437, v438, v439, v440, v442, v426)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L37
	} else {
		goto L107
	}
L106:
	;
	v445 = v435
	goto L102
L107:
	;
	v445 = v443
	goto L102
L108:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v415
	m.G0 = v15 + int32(16)
	return v421
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L37
	} else {
		goto L113
	}
L110:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if v454 != int32(429) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+168))
	F_RegisterExprContextCallback(m, v457, int32(1470), v421)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L37
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	F_errmsg_internal(m, int32(59523), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L37
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(489725), int32(4770), int32(313932))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
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
	F_errmsg_internal(m, int32(145928), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L37
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(485628), int32(251), int32(228064))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
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
			F_errcontext_msg(m, int32(506937), v9+int32(16))
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
			F_errcontext_msg(m, int32(313486), v9)
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
