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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v488 int32
	_ = v488
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
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
		v488 = v5
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L15
	} else {
		goto L91
	}
L5:
	;
	return v488
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= int32(0) {
		v488 = v5
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v44 = v5
	v47 = v5
	v49 = v5
	goto L8
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v49<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+41)))
	if v61 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v488 = v461
	goto L5
L10:
	;
	if v362 == int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L79
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
	if v64 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v139 == int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L33
	}
L14:
	;
	v67 = F_get_sortgroupref_tle(m, v64, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+26)))
	if v71 != 0 {
		v488 = v44
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v73 == int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v76 <= int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v85 = int32(0)
	goto L20
L20:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v79+v85<<(uint(int32(2))%32))))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v108 == int32(6) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v117 = F_copyObjectImpl(m, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+8)))
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	if v111 == v112 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v115 = v85 + int32(1)
	if v76 != v115 {
		v85 = v115
		goto L20
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v488 = v44
	goto L5
L28:
	;
	if v117 == int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v127 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v130 = F_get_eclass_for_sort_expr(m, l0, v117, v121, v125, v126, v127, v128, v127)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	if v130 == int32(0) {
		v488 = v44
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	v137 = F_make_canonical_pathkey(m, l0, v130, v134, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v362 = v137
	goto L10
L33:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v143 <= v142 {
		v362 = v142
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v154 = v142
	v158 = int32(-1)
	v159 = int32(0)
	goto L35
L35:
	;
	if l3 == int32(0) {
		v335 = v154
		v339 = v158
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v362 = v335
	goto L10
L37:
	;
	v353 = v159 + int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v353 < v354 {
		v154 = v335
		v158 = v339
		v159 = v353
		goto L35
	} else {
		goto L78
	}
L38:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v173 <= int32(0) {
		v335 = v154
		v339 = v158
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v159<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v191 = v154
	v195 = v158
	v199 = int32(0)
	goto L40
L40:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v199<<(uint(int32(2))%32))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+26)))
	if v213 != 0 {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v335 = v308
	v339 = v312
	goto L37
L42:
	;
	v326 = v199 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v326 < v327 {
		v191 = v308
		v195 = v312
		v199 = v326
		goto L40
	} else {
		goto L77
	}
L43:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v215 == int32(0) {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v218 <= int32(0) {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v227 = int32(0)
	goto L46
L46:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v221+v227<<(uint(int32(2))%32))))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v250 == int32(6) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v259 = F_copyObjectImpl(m, v249)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+8)))
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+8)))
	if v253 == v254 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v257 = v227 + int32(1)
	if v218 != v257 {
		v227 = v257
		goto L46
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v308 = v191
	v312 = v195
	goto L42
L54:
	;
	if v259 == int32(0) {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L55
	}
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v264 = F_canonicalize_ec_expression(m, v263, v181, v183)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v266 = F_equal(m, v264, v182)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	if v266 == int32(0) {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v274 = F_get_eclass_for_sort_expr(m, l0, v259, v270, v181, v183, v271, v272, v271)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	if v274 == int32(0) {
		v308 = v191
		v312 = v195
		goto L42
	} else {
		goto L60
	}
L60:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	v281 = F_make_canonical_pathkey(m, l0, v274, v278, v279, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	if v283 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v286 = v284
	goto L64
L63:
	;
	v286 = int32(0)
	goto L64
L64:
	;
	v288 = v286 - int32(1)
	if v47 < v26 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v47<<(uint(int32(2))%32))))
	if v295 == v281 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v298 = v288
	goto L67
L67:
	;
	v299 = base.B2i32(v195 < v298)
	if v195 < v298 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v297 = v286
	goto L70
L69:
	;
	v297 = v288
	goto L70
L70:
	;
	v298 = v297
	goto L67
L71:
	;
	v300 = v281
	goto L73
L72:
	;
	v300 = v191
	goto L73
L73:
	;
	if v195 < v298 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v301 = v298
	goto L76
L75:
	;
	v301 = v195
	goto L76
L76:
	;
	v308 = v300
	v312 = v301
	goto L42
L77:
	;
	goto L41
L78:
	;
	goto L36
L79:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+40)))
	if v382 != 0 {
		v461 = v44
		v464 = v47
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v473 = v49 + int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v473 < v474 {
		v44 = v461
		v47 = v464
		v49 = v473
		goto L8
	} else {
		goto L90
	}
L81:
	;
	if v44 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v447 = F_lappend(m, v44, v362)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L15
	} else {
		goto L89
	}
L83:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v385 <= int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v394 = int32(0)
	goto L85
L85:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v388+v394<<(uint(int32(2))%32))))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v381 == v417 {
		v461 = v44
		v464 = v47
		goto L80
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	v420 = v394 + int32(1)
	if v385 != v420 {
		v394 = v420
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v461 = v447
	v464 = v47 + int32(1)
	goto L80
L90:
	;
	goto L9
L91:
	;
	F_errmsg_internal(m, int32(_a_F_convert_subquery_pathkeys_0), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L15
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_convert_subquery_pathkeys_1), int32(1080), int32(_a_F_convert_subquery_pathkeys_2))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L15
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v47 int32
	_ = v47
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v131 int32
	_ = v131
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
	var v152 int32
	_ = v152
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
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
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
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
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
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v351 int32
	_ = v351
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
		v351 = v4
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
		goto L110
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L20
	} else {
		goto L107
	}
L3:
	;
	return v351
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v12 != 0 {
		v351 = v4
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
		v351 = v4
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
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if l0 == l1 {
		goto L79
	} else {
		goto L80
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
	v47 = v4
	goto L24
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v47<<(uint(int32(2))%32))))
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
	v244 = v47 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v244 < v245 {
		v47 = v244
		goto L24
	} else {
		goto L77
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
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225))))
	v229 = v221 + v228
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v231 = v230 | v222
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v231)
	goto L26
L29:
	;
	v74 = v52 + int32(8)
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+8)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v75))))
	if v77&int32(1) != 0 {
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
	v221 = v66
	v222 = int32(2)
	v225 = v52 + int32(8)
	goto L28
L34:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v221 = v217
	v222 = v216
	v225 = v74
	goto L28
L35:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v86 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v81 = F_contain_volatile_functions(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	if v81 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v216 = int32(1)
	goto L34
L39:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v133 != int32(1) {
		goto L26
	} else {
		goto L55
	}
L40:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v90))))
	if v92&int32(4) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if base.B2i32(v98 == v96)|base.B2i32(v95 == v96) != 0 {
		v131 = v96
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v131 != 0 {
		goto L39
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if int32(0) < v104 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v108 = int32(0)
	goto L48
L46:
	;
	goto L47
L47:
	;
	v131 = v96
	goto L43
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v108<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v98 == v117 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v131 = int32(1)
	goto L43
L51:
	;
	goto L52
L52:
	;
	v121 = v108 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v121 < v122 {
		v108 = v121
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v216 = int32(4)
	goto L34
L55:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74))))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
	if v139&int32(4) != 0 {
		goto L26
	} else {
		goto L56
	}
L56:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v142 == int32(0) {
		goto L26
	} else {
		goto L57
	}
L57:
	;
	v145 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v146 <= v145 {
		goto L26
	} else {
		goto L58
	}
L58:
	;
	v152 = v145
	goto L59
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v152<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if base.B2i32(v167 == v165)|base.B2i32(v164 == v165) != 0 {
		v200 = v165
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v216 = int32(8)
	goto L34
L61:
	;
	if v200 != 0 {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	goto L61
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if int32(0) < v173 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v177 = int32(0)
	goto L67
L65:
	;
	goto L66
L66:
	;
	v200 = v165
	goto L62
L67:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v177<<(uint(int32(2))%32))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v167 == v186 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L66
L69:
	;
	v200 = int32(1)
	goto L62
L70:
	;
	goto L71
L71:
	;
	v190 = v177 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v190 < v191 {
		v177 = v190
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v202 = v152 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v202 < v203 {
		v152 = v202
		goto L59
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L60
L76:
	;
	goto L26
L77:
	;
	goto L25
L78:
	;
	v351 = int32(1)
	goto L3
L79:
	;
	if v257 == int32(0) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v257 != 0 {
		v351 = v4
		goto L3
	} else {
		goto L85
	}
L82:
	;
	v261 = F_recurse_pushdown_safe(m, v257, l1, l2)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	if v261 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	v351 = v4
	goto L3
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	if v265 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v268 = v266
	goto L88
L87:
	;
	v268 = int32(0)
	goto L88
L88:
	;
	if v263 == int32(0) {
		v330 = v268
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if v330 != 0 {
		goto L1
	} else {
		goto L106
	}
L90:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v271 <= int32(0) {
		v330 = v268
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v275 = int32(0)
	v276 = v271
	v281 = v268
	goto L92
L92:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v275<<(uint(int32(2))%32))))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+26)))
	if v290 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v330 = v320
	goto L89
L94:
	;
	if v281 == int32(0) {
		goto L2
	} else {
		goto L97
	}
L95:
	;
	v319 = v276
	v320 = v281
	goto L96
L96:
	;
	v322 = v275 + int32(1)
	if v322 < v319 {
		v275 = v322
		v276 = v319
		v281 = v320
		goto L92
	} else {
		goto L105
	}
L97:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v296 = F_exprType(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v296 != v298 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(v289)+8)))
	v302 = v300 + v301
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v305 = v303 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v302))) = uint8(v305)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v309 = v281 + int32(4)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if base.Ui32(v309) < base.Ui32(v311+v312<<(uint(int32(2))%32)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v317 = v309
	goto L104
L103:
	;
	v317 = int32(0)
	goto L104
L104:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v319 = v318
	v320 = v317
	goto L96
L105:
	;
	goto L93
L106:
	;
	goto L78
L107:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_is_pushdown_safe_0), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_subquery_is_pushdown_safe_1), int32(3862), int32(_a_F_subquery_is_pushdown_safe_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L20
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_is_pushdown_safe_0), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_subquery_is_pushdown_safe_1), int32(3868), int32(_a_F_subquery_is_pushdown_safe_2))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L20
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
