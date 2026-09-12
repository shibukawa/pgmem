package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 float64
	_ = v286
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 float64
	_ = v350
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 float64
	_ = v406
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v430 int32
	_ = v430
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(base.B2i32(v487 == int32(0)))
	return v12
L5:
	;
	v487 = v465
	goto L4
L6:
	;
	v465 = int32(1)
	goto L5
L7:
	;
	v36 = v31
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(8)
	v38 = v10 + v37
	v40 = v5 + v37
	v48 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L47
	} else {
		goto L48
	}
L13:
	;
	v60 = v48 << (uint(int32(3)) % 32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v40+v60)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v72 = v62
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = int32(8)
	v108 = v10 + v107
	v110 = v5 + v107
	v119 = int32(0)
	goto L30
L15:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38+v60)))
	v76 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v83 = v74
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v69) != 0 {
		v72 = v62
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if base.F64_gt(v72, v83) != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v81) != 0 {
		v83 = v74
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v83 = v81
	goto L18
L21:
	;
	if v29 < int32(0) {
		v92 = v62
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 < int32(0) {
		v100 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v90) != 0 {
		v92 = v62
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v90
	goto L22
L25:
	;
	v102 = int32(-1)
	if base.F64_lt(v92, v100) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v98) != 0 {
		v100 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	goto L25
L28:
	;
	v105 = v48 + int32(1)
	if v105 != v36 {
		v48 = v105
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v131 = v119 << (uint(int32(3)) % 32)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v110+v131)))
	v135 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v143 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v108+v131)))
	v147 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v154 = v145
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v140) != 0 {
		v143 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v143 = v140
	goto L32
L35:
	;
	if base.F64_gt(v143, v154) != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v152) != 0 {
		v154 = v145
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = v152
	goto L35
L38:
	;
	if v29 < int32(0) {
		v163 = v133
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v32 < int32(0) {
		v171 = v145
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v161) != 0 {
		v163 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L39
L42:
	;
	if base.F64_lt(v163, v171) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v169) != 0 {
		v171 = v145
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v171 = v169
	goto L42
L45:
	;
	v175 = v119 + int32(1)
	if v175 != v36 {
		v119 = v175
		goto L30
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	v197 = v5 + int32(8)
	v207 = v36
	goto L50
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v220 = v197 + v207<<(uint(int32(3))%32)
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v266 = v36
	goto L68
L52:
	;
	if base.F64_lt(v243, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v225 = int32(3)
	v227 = v197 + (v207+v29)<<(uint(v225)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v207+v31)<<(uint(v225)%32))))
	if base.F64_lt(v221, v232) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_gt(v221, float64(0)) != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v234 = v220
	goto L58
L57:
	;
	v234 = v227
	goto L58
L58:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v234)))
	if base.F64_gt(v235, float64(0)) != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.F64_lt(v221, v238) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v240 = v221
	goto L62
L61:
	;
	v240 = v238
	goto L62
L62:
	;
	v243 = v240
	goto L52
L63:
	;
	v243 = v221
	goto L52
L64:
	;
	v487 = int32(-1)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v250 = v207 + int32(1)
	if v250 != v31 {
		v207 = v250
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	v274 = v197 + v266<<(uint(int32(3))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v487 = int32(-1)
	goto L4
L70:
	;
	if base.F64_lt(v297, float64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v279 = int32(3)
	v281 = v197 + (v29+v266)<<(uint(v279)%32)
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v266+v31)<<(uint(v279)%32))))
	if base.F64_gt(v275, v286) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if base.F64_gt(v275, float64(0)) != 0 {
		goto L6
	} else {
		goto L81
	}
L74:
	;
	v288 = v274
	goto L76
L75:
	;
	v288 = v281
	goto L76
L76:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.F64_gt(v289, float64(0)) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_gt(v275, v292) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v275
	goto L80
L79:
	;
	v294 = v292
	goto L80
L80:
	;
	v297 = v294
	goto L70
L81:
	;
	v297 = v275
	goto L70
L82:
	;
	v304 = int32(1)
	v306 = v266 + v304
	if v306 == v31 {
		v465 = v304
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L69
L85:
	;
	v266 = v306
	goto L68
L86:
	;
	v487 = int32(0)
	goto L4
L87:
	;
	goto L88
L88:
	;
	v312 = v10 + int32(8)
	v328 = v31
	goto L89
L89:
	;
	v335 = v312 + v328<<(uint(int32(3))%32)
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v386 = v36
	goto L106
L91:
	;
	if base.F64_lt(v364, float64(0)) != 0 {
		goto L6
	} else {
		goto L104
	}
L92:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_lt(v336, v361) != 0 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v340 = int32(3)
	v342 = v312 + (v32+v328)<<(uint(v340)%32)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v328+v34)<<(uint(v340)%32))))
	if base.F64_lt(v336, v347) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.F64_gt(v336, float64(0)) == int32(0) {
		v364 = v336
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v349 = v335
	goto L98
L97:
	;
	v349 = v342
	goto L98
L98:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v349)))
	if base.F64_gt(v350, float64(0)) == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v487 = int32(-1)
	goto L4
L100:
	;
	v487 = int32(-1)
	goto L4
L101:
	;
	v363 = v336
	goto L103
L102:
	;
	v363 = v361
	goto L103
L103:
	;
	v364 = v363
	goto L91
L104:
	;
	v370 = v328 + int32(1)
	if v370 != v34 {
		v328 = v370
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v394 = v312 + v386<<(uint(int32(3))%32)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v465 = int32(-1)
	goto L5
L108:
	;
	if base.F64_lt(v423, float64(0)) != 0 {
		goto L6
	} else {
		goto L121
	}
L109:
	;
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v401)))
	if base.F64_gt(v395, v420) != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v399 = int32(3)
	v401 = v312 + (v32+v386)<<(uint(v399)%32)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v386+v34)<<(uint(v399)%32))))
	if base.F64_gt(v395, v406) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.F64_gt(v395, float64(0)) == int32(0) {
		v423 = v395
		goto L108
	} else {
		goto L117
	}
L113:
	;
	v408 = v394
	goto L115
L114:
	;
	v408 = v401
	goto L115
L115:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)))
	if base.F64_gt(v409, float64(0)) == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v487 = int32(-1)
	goto L4
L117:
	;
	v487 = int32(-1)
	goto L4
L118:
	;
	v422 = v395
	goto L120
L119:
	;
	v422 = v420
	goto L120
L120:
	;
	v423 = v422
	goto L108
L121:
	;
	v430 = v386 + int32(1)
	if v34 != v430 {
		v386 = v430
		goto L106
	} else {
		goto L122
	}
L122:
	;
	goto L107
}
