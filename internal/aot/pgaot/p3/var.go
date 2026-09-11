package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418+l2))) = v416
	return
L2:
	;
	v416 = v413
	v418 = int32(8)
	goto L1
L3:
	;
	v413 = int32(_a_F_add_var_0)
	goto L2
L4:
	;
	if v5 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if v5 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L7:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = int32(0)
	if base.B2i32(v19 < v16)&base.B2i32(v20 < v15) == v20 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	return
L11:
	;
	v413 = int32(0)
	goto L2
L12:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L70
	}
L13:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L69
	}
L14:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v192 != 0 {
		goto L62
	} else {
		goto L63
	}
L15:
	;
	switch v191 {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L12
	}
L16:
	;
	v191 = v181
	goto L15
L17:
	;
	if v19 <= v51 {
		v86 = v19
		v88 = v20
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v51 = v16
	v55 = v20
	goto L17
L19:
	;
	goto L20
L20:
	;
	v32 = v16
	v36 = v20
	goto L21
L21:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v181 = int32(1)
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v51 = v46
	v55 = v44
	goto L17
L23:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v44 < v15 {
		v32 = v46
		v36 = v44
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	if v51 != v86 {
		v127 = v55
		v128 = v88
		goto L34
	} else {
		goto L35
	}
L27:
	;
	if v18 <= int32(0) {
		v86 = v19
		v88 = v20
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v67 = v19
	v69 = v20
	goto L29
L29:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v69<<(uint(int32(1))%32)))))
	if v74 != 0 {
		v181 = int32(-1)
		goto L16
	} else {
		goto L31
	}
L30:
	;
	v86 = v78
	v88 = v76
	goto L26
L31:
	;
	v75 = int32(1)
	v76 = v69 + v75
	v78 = v67 - v75
	if v78 <= v51 {
		v86 = v78
		v88 = v76
		goto L26
	} else {
		goto L32
	}
L32:
	;
	if v76 < v18 {
		v67 = v78
		v69 = v76
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if v15 < v127 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v97 = v55
	v98 = v88
	goto L36
L36:
	;
	if v15 <= v97 {
		v127 = v97
		v128 = v98
		goto L34
	} else {
		goto L38
	}
L37:
	;
	if base.I32_extend16_s(v113) < base.I32_extend16_s(v111) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v18 <= v98 {
		v127 = v97
		v128 = v98
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v102 = int32(1)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v97<<(uint(v102)%32)))))
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98<<(uint(v102)%32)+v17))))
	if v111 == v113 {
		v97 = v97 + v102
		v98 = v98 + v102
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v120 = int32(1)
	goto L43
L42:
	;
	v120 = int32(-1)
	goto L43
L43:
	;
	v191 = v120
	goto L15
L44:
	;
	v131 = v127
	goto L46
L45:
	;
	v131 = v15
	goto L46
L46:
	;
	v138 = v127
	goto L47
L47:
	;
	if v131 == v138 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v181 = v164
	goto L16
L49:
	;
	if v18 < v128 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v164 = int32(1)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v138<<(uint(v164)%32)))))
	if v170 == int32(0) {
		v138 = v138 + v164
		goto L47
	} else {
		goto L61
	}
L52:
	;
	v143 = v128
	goto L54
L53:
	;
	v143 = v18
	goto L54
L54:
	;
	v151 = v128
	goto L55
L55:
	;
	if v143 == v151 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v181 = int32(-1)
	goto L16
L57:
	;
	v191 = int32(0)
	goto L15
L58:
	;
	goto L59
L59:
	;
	v155 = int32(1)
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v151<<(uint(v155)%32)))))
	if v160 == int32(0) {
		v151 = v151 + v155
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	goto L48
L62:
	;
	F_pfree(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L10
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v197 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v197
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v202 < v201 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v204 = v201
	goto L68
L67:
	;
	v204 = v202
	goto L68
L68:
	;
	v416 = v204
	v418 = int32(12)
	goto L1
L69:
	;
	v413 = int32(0)
	goto L2
L70:
	;
	goto L3
L71:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v219 = int32(0)
	if base.B2i32(v218 < v215)&base.B2i32(v219 < v214) == v219 {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	goto L73
L73:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L133
	}
L74:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L10
	} else {
		goto L132
	}
L75:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L10
	} else {
		goto L131
	}
L76:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v391 != 0 {
		goto L124
	} else {
		goto L125
	}
L77:
	;
	switch v390 {
	case 0:
		goto L76
	case 1:
		goto L75
	default:
		goto L74
	}
L78:
	;
	v390 = v380
	goto L77
L79:
	;
	if v218 <= v250 {
		v285 = v218
		v287 = v219
		goto L88
	} else {
		goto L89
	}
L80:
	;
	v250 = v215
	v254 = v219
	goto L79
L81:
	;
	goto L82
L82:
	;
	v231 = v215
	v235 = v219
	goto L83
L83:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213+v235<<(uint(int32(1))%32)))))
	if v241 != 0 {
		v380 = int32(1)
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v250 = v245
	v254 = v243
	goto L79
L85:
	;
	v242 = int32(1)
	v243 = v235 + v242
	v245 = v231 - v242
	if v245 <= v218 {
		v250 = v245
		v254 = v243
		goto L79
	} else {
		goto L86
	}
L86:
	;
	if v243 < v214 {
		v231 = v245
		v235 = v243
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	if v250 != v285 {
		v326 = v254
		v327 = v287
		goto L96
	} else {
		goto L97
	}
L89:
	;
	if v217 <= int32(0) {
		v285 = v218
		v287 = v219
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v266 = v218
	v268 = v219
	goto L91
L91:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v268<<(uint(int32(1))%32)))))
	if v273 != 0 {
		v380 = int32(-1)
		goto L78
	} else {
		goto L93
	}
L92:
	;
	v285 = v277
	v287 = v275
	goto L88
L93:
	;
	v274 = int32(1)
	v275 = v268 + v274
	v277 = v266 - v274
	if v277 <= v250 {
		v285 = v277
		v287 = v275
		goto L88
	} else {
		goto L94
	}
L94:
	;
	if v275 < v217 {
		v266 = v277
		v268 = v275
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	if v214 < v326 {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v296 = v254
	v297 = v287
	goto L98
L98:
	;
	if v214 <= v296 {
		v326 = v296
		v327 = v297
		goto L96
	} else {
		goto L100
	}
L99:
	;
	if base.I32_extend16_s(v312) < base.I32_extend16_s(v310) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	if v217 <= v297 {
		v326 = v296
		v327 = v297
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v301 = int32(1)
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213+v296<<(uint(v301)%32)))))
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297<<(uint(v301)%32)+v216))))
	if v310 == v312 {
		v296 = v296 + v301
		v297 = v297 + v301
		goto L98
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v319 = int32(1)
	goto L105
L104:
	;
	v319 = int32(-1)
	goto L105
L105:
	;
	v390 = v319
	goto L77
L106:
	;
	v330 = v326
	goto L108
L107:
	;
	v330 = v214
	goto L108
L108:
	;
	v337 = v326
	goto L109
L109:
	;
	if v330 == v337 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v380 = v363
	goto L78
L111:
	;
	if v217 < v327 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v363 = int32(1)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213+v337<<(uint(v363)%32)))))
	if v369 == int32(0) {
		v337 = v337 + v363
		goto L109
	} else {
		goto L123
	}
L114:
	;
	v342 = v327
	goto L116
L115:
	;
	v342 = v217
	goto L116
L116:
	;
	v350 = v327
	goto L117
L117:
	;
	if v342 == v350 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v380 = int32(-1)
	goto L78
L119:
	;
	v390 = int32(0)
	goto L77
L120:
	;
	goto L121
L121:
	;
	v354 = int32(1)
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v350<<(uint(v354)%32)))))
	if v359 == int32(0) {
		v350 = v350 + v354
		goto L117
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	goto L110
L124:
	;
	F_pfree(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v396 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v396
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v396
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v401 < v400 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L126
L128:
	;
	v403 = v400
	goto L130
L129:
	;
	v403 = v401
	goto L130
L130:
	;
	v416 = v403
	v418 = int32(12)
	goto L1
L131:
	;
	goto L3
L132:
	;
	v413 = int32(0)
	goto L2
L133:
	;
	goto L3
}
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
	v13 = F_pull_var_clause_walker(m, l0, v6+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(16)
		return v17
	}
}
